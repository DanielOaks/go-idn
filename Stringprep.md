Go-idn includes a generic stringprep implementation which includes a number of standard stringprep profiles and an API for user-defined profiles.

**THIS PAGE IS A WORK IN PROGRESS!**

Stringprep is a framework for preparing Unicode text strings in order to increase the likelihood that string input and string comparison work in ways that make sense for typical users throughout the world. The stringprep protocol is useful for protocol identifier values, company and personal names, internationalized domain names, and other text strings.

Stringprep is defined in [RFC 3454](http://tools.ietf.org/html/rfc3454).



# API #


## func StringprepRunes ##

```
func StringprepRunes(input []int, profile Profile) []int
```

Prepare the input rune array according to the stringprep profile, and return the results as rune array.


---

## func Stringprep ##
```
func Stringprep(input string, profile Profile) string
```

Prepare the input string according to the stringprep profile, and return the results as an UTF-8 string


---

## func Nameprep ##
```
func Nameprep(label string) string
```

Same as `Stringprep(label, Profiles["nameprep"])`


---

## func Nodeprep ##
```
func Nodeprep(label string) string
```

Same as `Stringprep(label, Profiles["nodeprep"])`


---

## func Resourceprep ##
```
func Resourceprep(label string) string
```

Same as `Stringprep(label, Profiles["resourceprep"])`


---




# Profiles #


The standard profiles which are included with Go-idn are:
  * [Nameprep](http://tools.ietf.org/html/rfc3491)
  * [XMPP](http://tools.ietf.org/html/rfc3920)
    * Nodeprep
    * Resourceprep
  * SASL
    * [SASLprep SASLprep](http://tools.ietf.org/html/rfc4013)
  * iSCSI

The profiles are exported through the Profiles variable
```
var Profiles = map[string]Profile{
	"Nameprep":     _nameprep,
	"Nodeprep":     _xmpp_nodeprep,
	"Resourceprep": _xmpp_resourceprep,
	"plain":        _plain, /* sasl-anon-00. */
	"trace":        _trace, /* sasl-anon-01,02,03. */
	"SASLprep":     _saslprep,
	"iSCSI":        _iscsi, /* IANA. */
}
```








## Profile API ##
A profile of stringprep can create tables different from those in the Tables map, but it will be an exception when they do.  The intention of stringprep is to define the tables and have the profiles of stringprep select among those defined tables.


```
/* Steps in a stringprep profile. */
const (
	NFKC                = 1
	BIDI                = 2
	MAP_TABLE           = 3
	UNASSIGNED_TABLE    = 4
	PROHIBIT_TABLE      = 5
	BIDI_PROHIBIT_TABLE = 6
	BIDI_RAL_TABLE      = 7
	BIDI_L_TABLE        = 8
)
```

The stringprep package defines a profile as an array of ProfileElements
```
type Profile []ProfileElement
type ProfileElement struct {
	Step  int // see Step const's
	Table Table
}
```


**A profile example**
```
// Nameprep - As descrirbed in RFC 3491: http://tools.ietf.org/html/rfc3491
var Nameprep = Profile{
	ProfileElement{MAP_TABLE, Tables["B1"]},
	ProfileElement{MAP_TABLE, Tables["B2"]},
	ProfileElement{NFKC, nil},
	ProfileElement{PROHIBIT_TABLE, Tables["C12"]},
	ProfileElement{PROHIBIT_TABLE, Tables["C22"]},
	ProfileElement{PROHIBIT_TABLE, Tables["C3"]},
	ProfileElement{PROHIBIT_TABLE, Tables["C4"]},
	ProfileElement{PROHIBIT_TABLE, Tables["C5"]},
	ProfileElement{PROHIBIT_TABLE, Tables["C6"]},
	ProfileElement{PROHIBIT_TABLE, Tables["C7"]},
	ProfileElement{PROHIBIT_TABLE, Tables["C8"]},
	ProfileElement{PROHIBIT_TABLE, Tables["C9"]},
	ProfileElement{BIDI, nil},
	ProfileElement{BIDI_PROHIBIT_TABLE, Tables["C8"]},
	ProfileElement{BIDI_RAL_TABLE, Tables["D1"]},
	ProfileElement{BIDI_L_TABLE, Tables["D2"]},
	ProfileElement{UNASSIGNED_TABLE, Tables["A1"]},
}
```


## Tables ##
The stringprep package exports the tables defined in RFC 3454 through the Tables map. As noted in the Profiles chapter these tables are used to create profiles

```
type TableElement struct {
	Lo  int
	Hi  int
	Map d // can be empty
}

type Table []TableElement
type d [MAX_MAP_CHARS]int


var Tables = map[string]Table {
	"A1":_A1,
	"B1":_B1,
	"B2":_B2,
	"B3":_B3,
	"C11":_C11,
	"C12":_C12,
	"C21":_C21,
	"C22":_C22,
	"C3":_C3,
	"C4":_C4,
	"C5":_C5,
	"C6":_C6,
	"C7":_C7,
	"C8":_C8,
	"C9":_C9,
	"D1":_D1,
	"D2":_D2,
}
```

# Current status #
The stringprep framework has been implemented 85% and is mostly hanging on Go's lack of support for Unicode normalization (NFKC). The API is _not_ stable and is likely to change at some point. All the standard profiles have been implemented.

The implementation is still lacking proper unit tests, which will eventually be implemented once we assemble some test cases.


# Related Resources #

  * [Stringprep specification](http://tools.ietf.org/html/rfc3454)
  * [Nameprep: Stringprep profile for IDN](http://tools.ietf.org/html/rfc3491)
  * [iSCSI: Stringprep profile for iSCSI](http://tools.ietf.org/html/rfc3722)
  * [Nodeprep and Resourceprep: Stringprep profiles for XMPP](http://tools.ietf.org/html/rfc3920)
  * [SASLprep: Stringprep profile for usernames and passwords](http://tools.ietf.org/html/rfc4013)
  * [Trace: Stringprep profile for SASL ANONYMOUS tokens](http://tools.ietf.org/html/rfc4505)