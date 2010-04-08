// Copyright 2010 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

// This package contains methods for the preparation of internationalized strings ("stringprep") as described in RFC 3454
package stringprep

import (
	"unicode/normalization"
	"utf8"
	"bytes"
)

// Flags used when calling stringprep
const (
	NO_NFKC       = 1
	NO_BIDI       = 2
	NO_UNASSIGNED = 4
)

/* Steps in a stringprep profile. */
const (
	NFKC                = 1
	BIDI                = 2
	MAP_TABLE           = 3
	UNASSIGNED_TABLE    = 4
	PROHIBIT_TABLE      = 5
	BIDI_PROHIBIT_TABLE = 6
	BIDI_RAL_TABLE      = 7
	BIDI_L_TABLE        = 8
)

const (
	MAX_MAP_CHARS = 4
)


type TableElement struct {
	Lo  int
	Hi  int
	Map d // can be empty
}

type Table []TableElement
type Profile []ProfileElement
type d [MAX_MAP_CHARS]int

type ProfileElement struct {
	Step  int // see Step const's
	Table Table
}

func Nameprep(label string) string {
	return Stringprep(label, Profiles["nameprep"])
}

func Nodeprep(label string) string {
	return Stringprep(label, Profiles["nodeprep"])
}

func Resourceprep(label string) string {
	return Stringprep(label, Profiles["resourceprep"])
}

func NameprepRunes(label []int) []int {
	return StringprepRunes(label, Profiles["nameprep"])
}

func NodeprepRunes(label []int) []int {
	return StringprepRunes(label, Profiles["nodeprep"])
}

func ResourceprepRunes(label []int) []int {
	return StringprepRunes(label, Profiles["resourceprep"])
}


func Stringprep(input string, profile Profile) string {
	input_runes := bytes.Runes([]byte(input))
	return stringify(StringprepRunes(input_runes, profile))
}

// Prepare the input UCS-4 string according to the stringprep profile,
// and return the results.
func StringprepRunes(input []int, profile Profile) []int {
	output := make([]int, len(input))
	copy(output[0:], input[0:])

	for i := 0; i < len(profile); i++ {
		switch profile[i].Step {
		case NFKC:
			output = normalization.NFKC(output)
			break
		case BIDI:
			done_prohibited := 0
			done_ral := 0
			done_l := 0
			contains_ral := -1
			contains_l := -1

			for j := 0; i < len(profile); j++ {
				switch profile[j].Step {
				case BIDI_PROHIBIT_TABLE:
					done_prohibited = 1
					for k := 0; k < len(output); k++ {
						if in_table(output[k], profile[j].Table) {
							return nil
						}
					}

				case BIDI_RAL_TABLE:
					done_ral = 1
					for k := 0; k < len(output); k++ {
						if in_table(output[k], profile[j].Table) {
							contains_ral = j
						}
					}

				case BIDI_L_TABLE:
					done_l = 1
					for k := 0; k < len(output); k++ {
						if in_table(output[k], profile[j].Table) {
							contains_l = j
						}
					}
				}
			}

			if done_prohibited != 1 || done_ral != 1 || done_l != 1 {
				return nil // PROFILE ERROR
			}

			if contains_ral != -1 && contains_l != -1 {
				return nil // BIDI BOTH L AND RAL
			}

			if contains_ral != -1 {
				return nil // Error?
			}

			break
		case MAP_TABLE:
			output = map_table(output, profile[i].Table)
			break
		case UNASSIGNED_TABLE:
			break
			//switch profile[i].Table
		case PROHIBIT_TABLE:
			for k := 0; k < len(output); k++ {
				if in_table(output[k], profile[i].Table) {
					return nil
				}
			}
			break
		case BIDI_PROHIBIT_TABLE:
			break
		case BIDI_RAL_TABLE:
			break
		case BIDI_L_TABLE:
			break
		default:
			return nil // PROFILE ERROR
		}
	}

	return output
}


// Returns true if code point is in table 
func in_table(c int, table Table) bool {
	for i := 0; i < len(table); i++ {
		if table[i].Lo <= c && c <= table[i].Hi {
			return true
		}
	}
	return false
}

// Returns a filtered rune sequence 
func filter(input []int, table Table) []int {
	output := make([]int, len(input))
	c := 0 // count

	for i := 0; i < len(input); i++ {
		if !in_table(input[i], table) {
			output[c] = input[i]
			c++
		}
	}

	return output[0:len(output)]
}

func map_table(input []int, table Table) []int {

	output := make([]int, len(input))
	c := 0 // count

	for i := 0; i < len(input); i++ {
		// If rune is in table, replace it with its map
		if in_table(input[i], table) {
			for k := 0; k < len(table); k++ {
				if input[i] == table[k].Lo {
					copy(output[c:], table[k].Map[0:len(table[k].Map)])
					c += len(table[k].Map)
					break
				}
			}
		} else {
			output[c] = input[i]
			c++
		}

	}
	return output[0:len(output)]
}


// turn a slice of runes into an equivalent string 
func stringify(runes []int) string {
	t := make([]byte, len(runes)*4) // kludge! 
	i := 0
	for _, r := range runes {
		i += utf8.EncodeRune(r, t[i:])
	}
	return string(t)
}
