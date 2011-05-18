// Copyright 2010 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

// Package punycode implements the punycode data encoding as used for encoding
// of labels in the IDNA framework, as described in RFC 3492. Punycode is used
// by the IDNA protocol [IDNA] for converting domain labels into ASCII; it is
// not designed for any other purpose.
// It is explicitly not designed for processing arbitrary free text.
package punycode


import (
	"os"
	"strings"
	"utf8"
)


// Punycode uses the following Bootstring parameter values:
const (
	TMIN         = 1
	TMAX         = 26
	BASE         = 36
	INITIAL_N    = 0x80
	INITIAL_BIAS = 72
	DAMP         = 700
	SKEW         = 38
	DELIMITER    = 0x2D // hyphen '-'
)

// The maximum value of a signed int32.
// Used for overflow detection
const (
	MAXINT_S = 2147483647
)

// Error strings
const (
	BAD_INPUT = "Bad Input"
	OVERFLOW  = "Overflow"
)


// ToASCII returns the Punycode encoding of the input string and a nil os.Error when successful. 
// Input is assumed to be an utf8 encoded string
func ToASCII(input string) (string, os.Error) {
	var n int = INITIAL_N
	var delta int = 0
	var bias int = INITIAL_BIAS

	// Create a byte array for the output.
	output := make([]byte, 0, len(input))
	input_s := utf8.NewString(input) 

	
	// Copy all basic code points to the output
	var b int = 0
	for i := 0; i < input_s.RuneCount(); i++ {
		if isBasic(input_s.At(i)) {
			// input[i] is guranteed to be less than 128
			output = append(output, byte(input_s.At(i)))
			b++
		}
	}


	// Append DELIMITER 
	if b > 0 {
		output = append(output, DELIMITER)
	}

	var h int = b

	for h < input_s.RuneCount() {

		var m int = MAXINT_S

		// Find the minimum code point >= n
		for i := 0; i < input_s.RuneCount(); i++ {
			c  := input_s.At(i)
			if c >= n && c < m {
				m = c
			}
		}

		if (m - n) > ((MAXINT_S - delta) / (h + 1)) {
			// overflow
			return "", os.ErrorString(string(OVERFLOW))
		}

		delta = delta + (m-n)*(h+1)
		n = m

		for j := 0; j < input_s.RuneCount(); j++ {
			var c int = input_s.At(j)
			if c < n {
				delta++
				if 0 == delta {
					return "", os.ErrorString(string(OVERFLOW))
				}
			}

			if c == n {
				var q int = delta

				var k int
				for k = BASE; true; k += BASE {
					var t int

					if k <= bias {
						t = TMIN
					} else if k >= (bias + TMAX) {
						t = TMAX
					} else {
						t = k - bias
					}

					if q < t {
						break
					}

					var err os.Error
					var nbyte int
					nbyte, err = digit2codepoint(t + (q-t)%(BASE-t))

					if err != nil {
						return "", err
					}

					output = append(output, byte(nbyte))
					q = (q - t) / (BASE - t)

				}

				var err os.Error
				var nbyte int
				nbyte, err = digit2codepoint(q)

				if err != nil {
					return "", err
				}

				output = append(output, byte(nbyte))
				bias = adapt(delta, h == b, h+1)
				delta = 0
				h++

			}
		}

		delta++
		n++
	}
	

	return string(output[0:(len(output))]), nil
}


// ToUnicode takes a punycoded string and returns the decoded string and a nil os.Error when successful.
func ToUnicode(input string) (string, os.Error) {
	var n int = INITIAL_N
	var i int = 0
	var bias int = INITIAL_BIAS

	input_s := utf8.NewString(input) 
	var output []int  = make([]int, 0, len(input))


	
	var d int = strings.LastIndex(input, string(DELIMITER))
	if d > 0 {
		// Assert that characters before DELIMITER are ASCII
		// TODO: Clarify
		for j := 0; j < d; j++ {
			if !isBasic(input_s.At(j)) {
				return "", os.ErrorString(BAD_INPUT)
			}
			output = addCP(output, input_s.At(j))
		}
		d++

	} else {
		d = 0
	}

	for d < input_s.RuneCount() {
		
		var (
			oldi int = i
			w int = 1
			k int
		)
		
		for k = BASE; true; k += BASE {
			if d == input_s.RuneCount() {
				return "", os.ErrorString(BAD_INPUT)
			}

			var c int = input_s.At(d)
			d++

			var (
				err os.Error
				digit int
			)
			digit, err = codepoint2digit(c)

			if err != nil {
				return "", err
			}

			if digit > ((MAXINT_S - i) / w) {
				return "", os.ErrorString(OVERFLOW + " line 202")
			}

			i = i + digit*w

			var t int
			if k <= bias {
				t = TMIN
			} else if k >= bias+TMAX {
				t = TMAX
			} else {
				t = k - bias
			}

			if digit < t {
				break
			}
			w = w * (BASE - t)

		}

		bias = adapt(i-oldi, oldi == 0, len(output)+1)

		if i/(len(output)+1) > (MAXINT_S - n) {
			return "", os.ErrorString(OVERFLOW + " line 226")
		}

		n = n + i/(len(output)+1)
		i = i % (len(output) + 1)

		output = insert(output, i, n)
		i++
	}
	

	
	return string(output), nil
}



// Bias adaption function as described in RFC 3492 - 6.1
func adapt(delta int, first bool, numchars int) int {
	if first {
		delta = delta / DAMP
	} else {
		delta = delta / 2
	}

	delta = delta + (delta / numchars)

	var k int = 0
	for delta > ((BASE-TMIN)*TMAX)/2 {
		delta = delta / (BASE - TMIN)
		k = k + BASE
	}
	var bias int = k + ((BASE-TMIN+1)*delta)/(delta+SKEW)
	return bias
}

// Returns true if c < 128 (is a basic ASCII code point)
func isBasic(c int) bool {
	return c < 0x80
}


// codepoint2digit(cp) returns the numeric value of a basic rune
// (for use in representing integers) in the range 0 to 
// base-1, or base if cp does not represent a value.          
func codepoint2digit(cp int) (int, os.Error) {
	if cp-48 < 10 {
		// '0'..'9' : 26..35
		return cp - 22, nil
	} else if cp-65 < 26 {
		// 'a'..'z' : 0..25
		return cp - 65, nil
	} else if cp-97 < 26 {
		return cp - 97, nil
	} else {
		return BASE, nil
	}
	// else Bad input
	return -1, os.ErrorString(BAD_INPUT)
}

// Returns the rune and a non-nil Error when d < 36.
// Else it returns (unicode.MaxRune + 1) and a BadInputError
func digit2codepoint(d int) (int, os.Error) {

	if d < 26 {
		// 0..25 : 'a'..'z'
		return d + 'a', nil
	} else if d < 36 {
		// 26..35 : '0'..'9';
		return d - 26 + '0', nil
	}
	// else Bad input!
	return -1, os.ErrorString(BAD_INPUT)
}


// addCP appends  unicode cp  b to the end of s and returns the result.
// If s has enough capacity, it is extended in place; otherwise a
// new array is allocated and returned.
func addCP(s []int, t int) []int {
	lens := len(s)
	if lens+1 <= cap(s) {
		s = s[0 : lens+1]
	} else {
		news := make([]int, lens+1, resize(lens+1))
		copy(news, s)
		s = news
	}
	s[lens] = t
	return s
}

// How big to make a byte array when growing.
// Heuristic: Scale by 50% to give n log n time.
func resize(n int) int {
	if n < 16 {
		n = 16
	}
	return n + n/2
}


func insert(s []int, pos int, cp int) []int {
	lens:=len(s)
	a := s[0:pos]
	b := s[pos:]
	
	news := make([]int, lens+1, resize(lens+1))
	copy(news[0:], a)
	news[pos] = cp
	copy(news[pos+1:], b)
	
	s = news
	
	return s
}
/*
func insert(s []int, pos int, cp int) []int {
	lens:=len(s)
	a := s[0:pos]
	b := s[pos:]
	
	news := make([]int, lens+1, resize(lens+1))
	copy(news[0:], a)
	news[pos] = cp
	copy(news[pos+1:], b)
	
	s=news
	
	return s
}*/

func lastIndex(s []int, sep int) int {
	last := -1
	for i:=0;i<len(s);i++ {
		if s[i] == sep {
			last = i
		}
	}
	return last
}

// turn a slice of runes into an equivalent string
func stringify(runes []int) string {
	t := make([]byte, len(runes)*4) // kludge!
	i := 0
	for _, r := range runes {
		i += utf8.EncodeRune(t[i:], r)
	}
	return string(t)
}

func runify(str string) []int {
	t := make([]int, len(str))
	//i:=0
	for i, s:= range str {
		t[i]=s
	}
	return t
}
