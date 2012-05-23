// Package nameprep contains the profile data required to implement Nameprep. 
// The package is typically only imported for the side effect of registering its tables.
//
// To use nameprep, link this package into your program:
//     import _ "idna/stringprep/nameprep"
// Then use stringprep.Prep("nameprep", b) 
package nameprep

import (
	"idna/stringprep"
)

func init() {
	stringprep.RegisterProfile(Alias, prof)
}

// Alias is the name that nameprep is registered to with stringprep
const Alias = "nameprep"

var prof = &stringprep.Profile{}
