
package idn 

import (
 "go-idn.googlecode.com/hg/src/punycode"
 "go-idn.googlecode.com/hg/src/normalization"
 "go-idn.googlecode.com/hg/src/stringprep" 
 "go-idn.googlecode.com/hg/src/idna"
)

// Hack to make goinstall fetch all packages at once.
func NotImportant() bool {
	dep := idna.ACE_PREFIX
	dep2 := stringprep.NO_NFKC
	dep3  punycode.TMIN
	dep4 := NKFD
	
	return true
}


