// Copyright 2010 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

// Package stringprep implements Stringprep as described in RFC 3454
package stringprep

import (
	"bytes"
	"code.google.com/p/go-idn/idna/norm32"
	"unicode/utf8"
)

// InvalidStringError represents an invalid string error in the input stream.
type InvalidStringError struct {
	// TODO: What should we report?
}

func (e InvalidString) Error() string {
	panic("TODO")
	return "Invalid string"
}

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

// Exports tables
var Tables = map[string]Table{
	"A1":  _A1,
	"B1":  _B1,
	"B2":  _B2,
	"B3":  _B3,
	"C11": _C11,
	"C12": _C12,
	"C21": _C21,
	"C22": _C22,
	"C3":  _C3,
	"C4":  _C4,
	"C5":  _C5,
	"C6":  _C6,
	"C7":  _C7,
	"C8":  _C8,
	"C9":  _C9,
	"D1":  _D1,
	"D2":  _D2,
}

func Stringprep(input string, profile Profile) string {
	input_runes := bytes.Runes([]byte(input))
	return stringify(StringprepRunes(input_runes, profile))
}

// Prepare the input rune array according to the stringprep profile,
// and return the results as a rune array.
func StringprepRunes(input []rune, profile Profile) []int {
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

// Returns true if the rune is in table 
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

// Iterates over the input rune array and replaces runes with their maps
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
		i += utf8.EncodeRune(t[i:], r)
	}
	return string(t)
}
