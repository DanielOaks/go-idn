// Copyright 2012 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

package stringprep

import (
	// "fmt"
	// "reflect"
	"testing"
	// "os"

)

func TestInTable(t *testing.T) {
	c := 0x000221
	if !in_table(c, _A1) {
		t.Errorf("in_table(0x000221, _A1) = false; want true")
	}

	d := 0x000220
	if in_table(d, _A1) {
		t.Errorf("in_table(0x000220, _A1) = true; want false")
	}
}

func TestNameprep(t *testing.T) {
	input := []int{0x0644, 0x064A, 0x0647, 0x0645, 0x0627, 0x0628, 0x062A, 0x0643, 0x0644, 0x0645, 0x0648, 0x0634, 0x0639, 0x0631, 0x0628, 0x064A, 0x061F}
	output := StringprepRunes(input, Profiles["nameprep"])
	if output == nil {

		t.Errorf("Stringprep(\"asdf\", Profiles[\"nameprep\"]) = %s; want \"asdf\"", string(output))
	}

}
