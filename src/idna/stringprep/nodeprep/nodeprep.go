// Package nodeprep contains the profile data required to implement Nodeprep. 
// The package is typically only imported for the side effect of registering its tables.
//
// To use nodeprep, link this package into your program:
//     import _ "idna/stringprep/nodeprep"
// Then use stringprep.Prep("nodeprep", b) 
package nodeprep

import (
	"idna/stringprep"
)

func init() {
	stringprep.RegisterProfile(Alias, prof)
}

// Alias is the name that nodeprep is registered to with stringprep
const Alias "nodeprep"

var prof = &stringprep.Profile{
	
}