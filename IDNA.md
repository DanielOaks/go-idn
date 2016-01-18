IDNs use characters drawn from a large repertoire (Unicode), but IDNA allows the non-ASCII characters to be represented using only the ASCII characters already allowed in so-called host names today.  This backward-compatible representation is required in existing protocols  like DNS, so that IDNs can be introduced with no changes to the existing infrastructure.  IDNA is only meant for processing domain names, not free text.

The IDNA methodology encodes only select label components of domain names with procedures known as ToASCII and ToUnicode.

```
import "idn/idna"
```

This package implements a mechanism called IDNA for handling International Domain Names (IDN) in applications in a standard fashion as described [RFC 3490](http://tools.ietf.org/html/rfc3490).

# API #
## func ToASCII ##
```
func ToASCII(label string) (string, os.Error)
```
Converts a Unicode string to ASCII using the procedure in RFC 3490 section 4.1. Unassigned characters are not allowed and STD3 ASCII rules are enforced. The input string may be a domain name containing dots.

## func ToUnicode ##
```
func ToUnicode(label string) (string, os.Error)
```
Converts a Punycode string to Unicode using the procedure in RFC 3490 section 4.2. Unassigned characters are not allowed and STD3 ASCII rules are enforced. The input string may be a domain name containing dots.

# Current status #
IDNA has been implemented fully but hasn't been tested. It is waiting for unit-tests as well as [Nameprep](Stringprep.md).. The two exported function names are standard as described in [RFC 3490](http://tools.ietf.org/html/rfc3490) and are unlikely to change.