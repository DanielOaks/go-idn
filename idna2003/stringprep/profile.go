// Copyright 2012 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

package stringprep

type Profile struct {
	normalize bool
	checkBidi bool
	key       Trie
	val       []byte
}

func MkProfile(key Trie, val []byte, norm bool, bidi bool) *Profile {
	return &Profile{key, val}
}

func (p *Profile) Value(r rune) string {

}

const (
	typeMask  = 0xC000 // 11000000 00000000
	valueMask = 0x3FFF // 00111111 11111111
)

type valueType int

const (
	Unassigned valueType = iota
	Map
	Prohibited
	Delete
)

func getType(u uint16) valueType {
	return (u & typeMask) >> 14
}

func getVal(u uint16) uint16 {
	return u & valueMask
}
