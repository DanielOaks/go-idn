// Copyright 2010 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

// This package implements a mechanism called  IDNA for handling
// International Domain Names (IDN) in applications in a standard fashion 
// as described RFC 3490
package idna

import (
	"idn/punycode" 
	"idn/stringprep" 
	"os"
	"strings"
)

// IDNA section 5
const (
	ACE_PREFIX = "xn--"
)


//
// Converts a Unicode string to ASCII using the procedure in RFC 3490
// section 4.1. Unassigned characters are not allowed and STD3 ASCII
// rules are enforced. The input string may be a domain name
// containing dots.
//
func ToASCII(label string) (string, os.Error) {


	// Step 1: If the sequence contains any code points outside the ASCII range
	// (0..7F) then proceed to step 2, otherwise skip to step 3.
	for i := 0; i < len(label); i++ {
		if label[i] > 127 {
			// Step 2: Perform the steps specified in [NAMEPREP] and fail if there is an error. 
			// The AllowUnassigned flag is used in [NAMEPREP].
			label = stringprep.Nameprep(label)
			break
		}
	}

	// Step 3: UseSTD3ASCIIRules is false
	// Step 4: If the sequence contains any code points outside the ASCII range 
	// (0..7F) then proceed to step 5, otherwise skip to step 8.

	isASCII := true
	for i := 0; i < len(label); i++ {
		if label[i] > 127 {
			isASCII = false
			break
		}

	}

	if !isASCII {

		// Step 5 Verify that the sequence does NOT begin with the ACE prefix.
		if strings.HasPrefix(label, ACE_PREFIX) {
			return label, os.ErrorString("Label starts with ACE prefix")
		}

		var err os.Error

		// Step 6: Encode with punycode
		label, err = punycode.ToASCII(label)
		if err != nil {
			return "", err // delegate err
		}
		// Step 7: Prepend ACE prefix
		label = ACE_PREFIX + label
	}
	
	//8. Verify that the number of code points is in the range 1 to 63 inclusive.
	if 0 < len(label) && len(label) < 64 {
		return label, nil
	}

	return "", os.ErrorString("label empty or too long")
}

//
// Converts a Punycode string to Unicode using the procedure in RFC 3490
// section 4.2. Unassigned characters are not allowed and STD3 ASCII
// rules are enforced. The input string may be a domain name
// containing dots.
//
func ToUnicode(label string) (string, os.Error) {

	// Step 1: If all code points in the sequence are in the ASCII range (0..7F) then skip to step 3.
	for i := 0; i < len(label); i++ {
		if label[i] > 127 {
			// Step 2: Perform the steps specified in [NAMEPREP] and fail if there is an error.
			label = stringprep.Nameprep(label)
			break
		}
	}
	
	original := label
	
	// Step 3: Verify that the sequence begins with the ACE prefix, and save a copy of the sequence.
	if !strings.HasPrefix(label, ACE_PREFIX) {
		return label, os.ErrorString("Label doesn't begin with the ACE prefix")
	}  // else
	
	// 4. Remove the ACE prefix.
	label = strings.Split(label, ACE_PREFIX, 0)[1]
	
	// 5. Decode the sequence using the decoding algorithm in [PUNYCODE] and fail if there is an error. 
	results, err := punycode.ToUnicode(label)
	
	if err != nil {
		return original, os.ErrorString("Failed punycode decoding: "+ err.String())
	}
	
	// 6. Apply ToASCII.
	verification, err := ToASCII(label)
	
	if err != nil {
		return original, os.ErrorString("Failed ToASCII on the decoded sequence: " + err.String())
	}
	
	// 7. Verify that the result of step 6 matches the saved copy from step 3, 
	// 	  using a case-insensitive ASCII comparison.
	if strings.ToLower(verification) == strings.ToLower(original) {
		return results, nil
	}
	
	return original, os.ErrorString("Failed verification step")
}
