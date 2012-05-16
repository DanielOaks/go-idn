// Package resourceprep contains the profile data required to implement SASLprep. 
// The package is typically only imported for the side effect of registering its tables.
//
// To use resourceprep, link this package into your program:
//     import _ "idna/stringprep/resourceprep"
// Then use stringprep.Prep("resourceprep", b) 
package resourceprep

import (
	"idna/stringprep"
)

func init() {
	stringprep.RegisterProfile(Alias, prof)
}

// Alias is the name that resourceprep is registered to with stringprep
const Alias "resourceprep"

var prof = &stringprep.Profile{
	
}