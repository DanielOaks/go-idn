Go-idn is a project that hopes to bring IDN to Go and aims to become feature compatible with libidn.

Go-idn is a mostly documented implementation of the Stringprep, Punycode and IDNA specifications. Go-idn's purpose is to encode and decode internationalized domain names using pure Go code.

The library contains a generic Stringprep implementation. Profiles for Nameprep, iSCSI, SASL and XMPP are included. Punycode and ASCII Compatible Encoding (ACE) via IDNA are supported. A mechanism to define Top-Level Domain (TLD) specific validation tables, and to compare strings against those tables, is included. Default tables for some TLDs are also included.

# Installing #
Installing go-idn is fairly simple. Just run `go get code.google.com/p/go-idn/idna` and all the packages will be included.

# Help is required! #
If there is anyone capable willing to help, I'm all ears (hannson@gmail.)

If anyone requests a Google Group for go-idn I will create one and participate.

# What's new? #
**2. October 2012:**
> I'm back! I've made a lot of backwards incompatible changes to the source code (sorry, it had to be done!) the most noticeable are the package names and the folder hierarchy. The normalizaton package has been completely replaced with norm32 (which is generously stolen from Golang's exp/norm and modified for Unicode 3.2.0 as defined in IDNA 2003.)
Once I've revisited the stringprep package and written proper unit tests I will probably make a stable release but until then you may find major and minor changes to the API.
All help is appreciated.

**7. September 2010:**
> The Unicode normalization is now complete. It succeeds all the test cases in [NormalizationTests.txt](http://unicode.org/Public/5.2.0/ucd/NormalizationTest.txt) including the PRI #29 Tests. Next up is code cleanup, more unit-tests and performance improvements.
**20. April 2010:**
> I'll be very busy doing other things until 15. May. I'm also looking for another set of eyes to bounce off ideas (for the API in particular).

> Feel free to send me an email at [hannson@gmail.com](mailto:hannson@gmail.com) if you're interested in this project and have questions or ideas.

**11. April 2010:**
> The Unicode normalization is making good progress. As of this writing the NFKC unit test passes on the first 13916 lines in [NormalizationTest.txt](http://www.unicode.org/Public/5.2.0/ucd/NormalizationTest.txt) (warning: 2.2MB).  The implementation of tables.go is still limited to 16bit runes which is unfortunate, but once we finish maketables.go it should pass more tests.


---



---


# Components #
![https://docs.google.com/drawings/pub?id=1SQkpxU6ewTgDJALcAHl69n5zuK58Hag3Zjw8oGB7ZnU&w=599&h=553&nonsense=something_that_ends_with.png](https://docs.google.com/drawings/pub?id=1SQkpxU6ewTgDJALcAHl69n5zuK58Hag3Zjw8oGB7ZnU&w=599&h=553&nonsense=something_that_ends_with.png)


# Status #
The current status of the project is 97%. It's mostly hanging on unit tests that need to be written and and nothing has been written for TLD (which is a low priority long term goal) but anything else is mostly working.

| **Component**                              | **Status** | **Notes**                                                                                                                           |
|:-------------------------------------------|:-----------|:------------------------------------------------------------------------------------------------------------------------------------|
| [Unicode normalization](Normalization.md)    | 100%       | Passes all the test cases in NormalizationTests.txt                                                                                 |
| [Punycode](Punycode.md)                               | 100%       | Passes all the test cases.                                                                                                          |
| [Stringprep](Stringprep.md)                             | 95%        | Unit tests need to be written to determine the correctness of the implementation.                                                   |
| [IDNA](IDNA.md)                                   | 95%        | Unit tests need to be written to determine the correctness of the implementation.                                                   |
| [TLD](TLD.md)                                    | _N/A_      | Work hasn't started.                                                                                                                |



---

# Punycode #
[Punycode](Punycode.md) is an instance of a general encoding syntax (Bootstring) by which a string of Unicode characters can be transformed uniquely and reversibly into a smaller, restricted character set. Punycode is intended for the encoding of labels in the Internationalized Domain Names in Applications (IDNA) framework, such that these domain names may be represented in the ASCII character set allowed in the Domain Name System of the Internet. The encoding syntax is defined in IETF document [RFC 3492](http://tools.ietf.org/html/rfc3492).

```
import "go-idn.googlecode.com/hg/src/punycode"
```

Package punycode implements the punycode data encoding as used for encoding of labels in the IDNA framework, as described in RFC 3492. Punycode is used by the IDNA protocol [IDNA](IDNA.md) for converting domain labels into ASCII; it is not designed for any other purpose. It is explicitly not designed for processing arbitrary free text.

## Current status ##
The specification has been implemented 100% and unit tests have been written. The code passes all unit tests (fails one, only when case-sensitive) and should be ready for review.



---

# Stringprep #
[Stringprep](Stringprep.md) is a framework for preparing Unicode text strings in order to increase the likelihood that string input and string comparison work in ways that make sense for typical users throughout the world. The stringprep protocol is useful for protocol identifier values, company and personal names, internationalized domain names, and other text strings.

```
import "go-idn.googlecode.com/hg/src/stringprep"
```
This package contains methods for the preparation of internationalized strings ("stringprep") as described in [RFC 3454](http://tools.ietf.org/html/rfc3492).

## Profiles ##

The following standard profiles are included.
  * Nameprep
  * XMPP
    * Nodeprep
    * Resourceprep
  * SASL
    * SASLprep
  * iSCSI


## Current status ##
All the standard profiles have been implemented. Needs some unit-test cases.



---

# IDNA #
The [IDNA](IDNA.md) methodology encodes only select label components of domain names with procedures known as ToASCII and ToUnicode.

```
import "go-idn.googlecode.com/hg/src/idna"
```

This package implements a mechanism called IDNA for handling International Domain Names (IDN) in applications in a standard fashion as described RFC 3490.


## Current status ##
IDNA specification has been implemented 100% and the API is stable because the function names are standard as described in [RFC 3490](http://tools.ietf.org/html/rfc3490), thus unlikely to change. It depends on [punycode](Punycode.md) (which fails a single case-sensitive unit test) but IDNA is specifically case-insensitive so it should be bug-free. It requires some unit-tests cases.




---

# TLD #

---

# Related resources #
  * [IDNA specification](http://tools.ietf.org/html/rfc3490)
  * [Punycode specification](http://tools.ietf.org/html/rfc3492)
  * [Stringprep specification](http://tools.ietf.org/html/rfc3492)
    * Standard profiles
      * [Nameprep: Stringprep profile for IDN](http://tools.ietf.org/html/rfc3491)
      * [iSCSI: Stringprep profile for iSCSI](http://tools.ietf.org/html/rfc3722)
      * [Nodeprep and Resourceprep: Stringprep profiles for XMPP](http://tools.ietf.org/html/rfc3920)
      * [SASLprep: Stringprep profile for usernames and passwords](http://tools.ietf.org/html/rfc4013)
      * [Trace: Stringprep profile for SASL ANONYMOUS tokens](http://tools.ietf.org/html/rfc4505)
  * [TLD specification](http://tools.ietf.org/html/draft-hoffman-idn-reg-02)
