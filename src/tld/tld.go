// Copyright 2010 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

// This package provides Top-Level Domain (TLD) specific validation tables, and a mechanism to compare strings against those tables
package tld

// The representation of a range of Unicode code points.  The range runs from Lo to Hi inclusive.
type Range struct {
	Lo int
	Hi int
}
