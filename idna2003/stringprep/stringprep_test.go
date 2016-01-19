// Copyright 2012 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

package stringprep

import "testing"

type testcase struct {
	Mapping string
	Input   []rune
	Output  []rune
}

// from http://tools.ietf.org/html/draft-josefsson-idn-test-vectors-00#section-4
var mappingTests = []testcase{
	{"nameprep", []rune{0x0066, 0x006f, 0x006f, 0x00ad, 0x034f, 0x1806, 0x180b, 0x0062, 0x0061, 0x0072, 0x200b, 0x2060, 0x0062, 0x0061, 0x007a, 0xfe00, 0xfe08, 0xfe0f, 0xfeff}, []rune{0x0066, 0x006f, 0x006f, 0x0062, 0x0061, 0x0072, 0x0062, 0x0061, 0x007a}},
	{"nameprep", []rune{0x0043, 0x0041, 0x0046, 0x0045}, []rune{0x0063, 0x0061, 0x0066, 0x0065}},
	{"nameprep", []rune{0x00df}, []rune{0x0073, 0x0073}},
}

func TestInTable(t *testing.T) {
	c := rune(0x000221)
	if !in_table(c, _A1) {
		t.Errorf("in_table(0x000221, _A1) = false; want true")
	}

	d := rune(0x000220)
	if in_table(d, _A1) {
		t.Errorf("in_table(0x000220, _A1) = true; want false")
	}
}

func TestNameprep(t *testing.T) {
	input := []rune{0x0644, 0x064A, 0x0647, 0x0645, 0x0627, 0x0628, 0x062A, 0x0643, 0x0644, 0x0645, 0x0648, 0x0634, 0x0639, 0x0631, 0x0628, 0x064A, 0x061F}
	_, err := StringprepRunes(input, Profiles["nameprep"])
	if err == nil {
		t.Error("stringprep: Incorrect rune array nameprep'd successfully, did not get any errors")
	}
}

func TestMapping(t *testing.T) {
	for i, test := range mappingTests {
		output, _ := StringprepRunes(test.Input, Profiles[test.Mapping])
		if string(output) != string(test.Output) {
			t.Error(
				"For test", i, test.Input,
				"expected", test.Output,
				"got", output,
			)
		}
	}
}
