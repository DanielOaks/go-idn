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

type Profile []ProfileElement
type d [MAX_MAP_CHARS]int

type ProfileElement struct {
	Step  int // see Step const's
	Table Table
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

// turn a slice of runes into an equivalent string 
func stringify(runes []int) string {
	t := make([]byte, len(runes)*4) // kludge! 
	i := 0
	for _, r := range runes {
		i += utf8.EncodeRune(t[i:], r)
	}
	return string(t)
}
