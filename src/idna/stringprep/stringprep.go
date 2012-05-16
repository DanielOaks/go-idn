// Copyright 2010 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

// Package stringprep provides data and functions to support
// the preparation of internationalized strings. See RFC 3454.
package stringprep

import (
	"unicode"
	"code.google.com/p/go-idn/src/norm"
)


type MapTable []CaseRange

value = data[index [cp>>LOWER_WIDTH] + (cp&LOWER_MASK)]

type Table struct {
}

// map maps ...
func (t *Table) map(b []byte) []byte {
	r := bytes.NewReader(b)
	out := bytes.Buffer
	for ch, size, err := r.ReadRune(); err != nil {
		if ch
	}
	
}

func Prep(profile string, b []byte) (result []byte, err error) {
	
	
	// Step 1. Map -- For each character in the input, check if it has a mapping
	// and, if so, replace it with its mapping
	for table := range p.Map {
		b = table.map(b)
	}
	
	// Step 2. Normalize -- Possibly normalizez the result of step 1 
	// using Unicode normalization
	if p.Normalize {
		norm.NFKC.Bytes(b)
	}
	
	// Step 3. Prohibit -- Check if any characters that are not allowed are 
	// in the output. If any are found, return an error.
	
	// Step 4. Check bidi -- Possibly check for right-to-left characters, 
	// and if any are found, make sure that the whole string satisfies the
	// requirements for bidirectional strings. If the string does not 
	// satisfy the requirements for bidirectional strings, return an error.

Error:
	return nil, err
}



func (p *Profile) PrepString(s string) (string, error) {
	res, err := p.Prep([]byte(s))
	return string(res), err
}