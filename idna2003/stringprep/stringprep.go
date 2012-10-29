// Copyright 2012 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

// Package stringprep implements Stringprep as described in RFC 3454
package stringprep

import (
	//"bytes"
	"code.google.com/p/go-idn/idna/norm32"
	"io"
	"unicode/utf8"
)

// InvalidStringError represents an invalid string error in the input stream.
type InvalidStringError struct {
	// TODO: What should we report?
}

func (e InvalidStringError) Error() string {
	panic("TODO")
	return "Invalid string"
}

const (
	MAX_MAP_CHARS = 4
)

// A Profile...
//
// For a Profile p, this documentation uses the notation p(x) to mean the bytes
// or string x prepared with the given profile.
type Profile struct {
	// Defaults to true
	Normalize bool

	// Defaults to false
	AllowUnassigned bool

	// Defined by profile
	CheckBidi bool

	f          norm32.Form
	mappings   Table
	prohibited Table
}

/*

// Append returns p(append(out, b...)). The buffer out must be nil, empty or
// equal to p(out).
func (p *Profile) Append(out []byte, src ...byte) []byte { return nil }

// AppendString returns p(append(out, []byte(s))). The buffer out must be nil, 
//empty, or equal to p(out).
func (p *Profile) AppendString(out []byte, src string) []byte { return nil }

// Bytes returns p(b). May return b if p(b) = b.
func (p *Profile) Bytes(b []byte) []byte { return nil }

// Reader returns a new reader that implements Read by reading data from r and
// returning p(data).
func (p *Profile) Reader(r io.Reader) io.Reader { return nil }

// String returns p(s).
func (p *Profile) String(s string) string { return "" }

// Writer returns a new writer that implements Write(b) by writing p(b) to w.
// The returned writer may use an an internal buffer to maintain state across
// Write calls. Calling its Close method writes any buffered data to w.
func (p *Profile) Writer(w io.Writer) io.WriteCloser { return nil }
*/
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
