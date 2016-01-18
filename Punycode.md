Punycode is an instance of a general encoding syntax (Bootstring) by which a string of Unicode characters can be transformed uniquely and reversibly into a smaller, restricted character set. Punycode is intended for the encoding of labels in the Internationalized Domain Names in Applications (IDNA) framework, such that these domain names may be represented in the ASCII character set allowed in the Domain Name System of the Internet. The encoding syntax is defined in IETF document [RFC 3492](http://tools.ietf.org/html/rfc3492).

```
import "idn/punycode"
```

Package punycode implements the punycode data encoding as used for encoding of labels in the IDNA framework, as described in RFC 3492. Punycode is used by the IDNA protocol for converting domain labels into ASCII; it is not designed for any other purpose. It is explicitly not designed for processing arbitrary free text.


## API ##


---

## func ToASCII ##
```
func ToASCII(input string) (string, os.Error)
```

ToASCII returns the Punycode encoding of the string input and a nil os.Error when successful.


---

## func ToASCIIRunes ##
```
func ToASCIIRunes(runes []int) (string, os.Error)
```

ToASCII returns the Punycode encoding of the Rune sequence "runes" and a nil os.Error when successful.


---

## func ToUnicode ##
```
func ToUnicode(input string) (string, os.Error)
```

ToUnicode returns the UTF-8 encoded string of the Punycode string input and a nil os.Error when successful.


---

## funct ToUnicodeRunes ##
```
func ToUnicodeRunes(input string) ([]int, os.Error)
```
ToUnicode returns the Unicode code point sequence of the Punycode sequence and a nil os.Error when successful.


---

# Current status #
The specification has been implemented 100% and unit tests have been written. The code passes all unit tests (fails one, only when case-sensitive) and should be ready for review.

The API is not 100% stable and may change.