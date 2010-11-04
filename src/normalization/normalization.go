// Copyright 2010 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

// Unicode Normalization. 
// Because Unicode contains a large number of precomposed characters
// there are multiple ways a character can be represented.
package normalization


// Normalization Forms. At the moment only NFKC is supported.
const (
	NFD  = iota // Canonical Decomposition
	NFC         // Canonicaal Decomposition followed by Canonical Composition
	NFKD        // Compatibility Decomposition
	//NFKC // Compatibility Decomposition, followed by Canonical Composition
)

/*
// TODO: Create a function all the normalization forms. 
//       The function can then be called by NFKC() with 
//       the  correct form parameter. 
func Normalize(input []int, form int) []int {
	var do_compat bool = (form == NFKC || form == NFKD)
	var do_compose bool = (mode == NFC || mode == NFKC)
}
//*/


// Applies NFKC normalization to an array of runes and returns a normalized rune array
func NFKC(input []int) []int {
	if len(input) == 0 {
		return input
	}
	output := make([]int, 0, len(input))

	for i := 0; i < len(input); i++ {
		var code int = input[i]

		// In Unicode 3.0, Hangul was defined as the block from U+AC00
		// to U+D7A3, however, since Unicode 3.2 the block extends until
		// U+D7AF. The decomposeHangul function only decomposes until
		// U+D7A3. Should this be changed?
		if code >= 0xAC00 && code <= 0xD7AF {

			// Append to output
			hang := decomposeHangul(code)
			for i := 0; i < len(hang); i++ {
				output = addCP(output, hang[i])

			}

		} else {
			// Decompose:
			// Here we replace 'code' with the value in _Decomposition[code] if it exist
			if mappings, ok := _Decomposition[code]; ok {
				for i := 0; i < len(mappings); i++ {
					output = addCP(output, mappings[i])
				}
			} else {
				output = addCP(output, code)
			}
		}
	}

	// Bring the string in to canonical order
	output = canonicalOrdering(output)

	// Do the canonical composition
	last_cc := 0
	last_start := 0

	for i := 0; i < len(output); i++ {
		cc := combiningClass(output[i])

		if i > 0 && (last_cc == 0 || last_cc < cc) {
			// try to combine characters 
			a := output[last_start]
			b := output[i]

			c := compose(a, b)

			if c != -1 {
				output[last_start] = c
				output = remove(output, i)

				i--

				if i == last_start {
					last_cc = 0
				} else {
					last_cc = combiningClass(output[i-1])
				}
				continue
			}
		}

		if cc == 0 {
			last_start = i
		}

		last_cc = cc

	}

	return output
}


func canonicalOrdering(input []int) []int {
	if len(input) <= 1 {
		return input
	}

	temp := 0
	ccHere := 0
	ccPrev := 0

	for i := 1; i < len(input); i++ {
		ccHere = combiningClass(input[i])
		ccPrev = combiningClass(input[i-1])

		if ccHere != 0 && ccPrev > ccHere {
			temp = input[i]
			input[i] = input[i-1]
			input[i-1] = temp
			if i > 1 {
				i -= 2
			}
		}
	}
	return input
}


/// Tries to compose two characters canonically and returns the composed character or -1 if no composition could be found.
func compose(a, b int) int {

	h := composeHangul(a, b)
	if h != -1 {
		return h
	}

	ai := composeIndex(a)
	if ai >= _Composition_singleFirstStart && ai < _Composition_singleSecondStart {
		if b == _Composition_singleFirst[ai-_Composition_singleFirstStart][0] {
			return _Composition_singleFirst[ai-_Composition_singleFirstStart][1]
		} else {
			return -1
		}
	}

	bi := composeIndex(b)

	if bi >= _Composition_singleSecondStart {
		if a == _Composition_singleSecond[bi-_Composition_singleSecondStart][0] {
			return _Composition_singleSecond[bi-_Composition_singleSecondStart][1]
		} else {
			return -1
		}
	}

	if ai >= 0 && ai < _Composition_multiSecondStart && bi >= _Composition_multiSecondStart && bi < _Composition_singleFirstStart {
		var f []int = _Composition_multiFirst[ai]

		if bi-_Composition_multiSecondStart < len(f) {
			var r int = f[bi-_Composition_multiSecondStart]
			if r == 0 {
				return -1
			} else {
				return r
			}
		}
	}
	return -1
}

// Returns the index of a rune inside the composition table if found, -1 otherwise
func composeIndex(a int) int {
	if a>>8 >= len(_Composition_composePage) {
		return -1
	}
	var ap int = _Composition_composePage[a>>8]
	if ap == -1 {
		return -1
	}

	return _Composition_composeData[ap][a&0xff]

}


// Returns the combining class of a given character.
func combiningClass(c int) int {

	var h int = c >> 8
	var l int = c & 0xff

	// fmt.Printf("combiningClass():\t"+strconv.Itob(c,  10)+"\t"+strconv.Itob(h,  10)+"\t"+strconv.Itob(l,  10)+"\n")


	var i int = _CombiningClass_i[h]
	if i > -1 {
		return _CombiningClass_c[i][l]
	} // else {

	return 0
}


// Entire hangul code copied from:
// http://www.unicode.org/unicode/reports/tr15/
// Several hangul specific constants
const (
	SBase  = 0xAC00
	LBase  = 0x1100
	VBase  = 0x1161
	TBase  = 0x11A7
	LCount = 19
	VCount = 21
	TCount = 28
	NCount = VCount * TCount
	SCount = LCount * NCount
)


// Composes two hangul characters/runes and returns the composed rune or -1 if the two runes cannot be composed.
func composeHangul(a, b int) int {

	// 1. check to see if two current characters are L and V
	var LIndex int = a - LBase
	if 0 <= LIndex && LIndex < LCount {
		var VIndex int = b - VBase
		if 0 <= VIndex && VIndex < VCount {
			// make syllable of form LV
			return SBase + (LIndex*VCount+VIndex)*TCount
		}
	}

	// 2. check to see if two current characters are LV and T
	var SIndex int = a - SBase
	if 0 <= SIndex && SIndex < SCount && (SIndex%TCount) == 0 {
		var TIndex int = b - TBase
		if 0 <= TIndex && TIndex <= TCount {
			// make syllable of form LVT
			return a + TIndex
		}
	}
	return -1
}

// 
// Decomposes a hangul character.
// 
// Returns a rune array containing the hangul decomposition of the input
// rune. If no hangul decomposition can be found, a rune array
// containing the rune itself is returned.</returns>
func decomposeHangul(s int) []int {

	var SIndex int = s - SBase

	if SIndex < 0 || SIndex >= SCount {
		out := make([]int, 1)
		out[0] = s
		return out
	}

	var L int = LBase + SIndex/NCount
	var V int = VBase + (SIndex%NCount)/TCount
	var T int = TBase + SIndex%TCount

	out := make([]int, 2, 3)

	out[0] = L
	out[1] = V
	if T != TBase {
		out = out[0:3]
		out[2] = T
		return out
	}

	return out
}


// addCP appends rune b to the end of s and returns the result.
// If s has enough capacity, it is extended in place; otherwise a
// new array is allocated and returned.
//
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
	lens := len(s)
	a := s[0:pos]
	b := s[pos:]

	news := make([]int, lens+1, resize(lens+1))
	copy(news[0:], a)
	news[pos] = cp
	copy(news[pos+1:], b)

	s = news

	return s
}

func remove(s []int, pos int) []int {
	lens := len(s)
	a := s[0:pos]
	b := s[pos+1:]

	news := make([]int, lens-1)
	copy(news[0:], a)
	copy(news[pos:], b)
	s = news
	return s

}
