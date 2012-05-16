// Package iscsi contains the profile data required to implement the iSCSI stringprep profile. 
// The package is typically only imported for the side effect of registering its tables.
//
// To use iscsi, link this package into your program:
//     import _ "idna/stringprep/iscsi"
// Then use stringprep.Prep("iscsi", b) 
package iscsi

import (
	"idna/stringprep"
)

func init() {
	stringprep.RegisterProfile(Alias, prof)
}

// Alias is the name that iSCSI is registered to with stringprep
const Alias "iscsi"

var prof = &stringprep.Profile{
	
}