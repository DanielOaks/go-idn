/*
Test cases extracted from RFC 3492

6.4 Overflow handling:
   For IDNA, 26-bit unsigned integers are sufficient to handle all valid
   IDNA labels without overflow, because any string that needed a 27-bit
   delta would have to exceed either the code point limit (0..10FFFF) or
   the label length limit (63 characters).  However, overflow handling
   is necessary because the inputs are not necessarily valid IDNA
   labels.

   The last example is an ASCII string that breaks the existing rules
   for host name labels.  (It is not a realistic example for IDNA,
   because IDNA never encodes pure ASCII labels.)

*/

package punycode

import (
	"fmt"
	"reflect"
	"testing"
	//   "os"
	//      "bytes"
	"unicode/utf8"
)

type testCase struct {
	unicode  []rune
	punycode string
}

var punyTests = []testCase{
	testCase{ // (A) Arabic (Egyptian)
		[]rune{0x0644, 0x064A, 0x0647, 0x0645, 0x0627, 0x0628, 0x062A, 0x0643, 0x0644, 0x0645, 0x0648, 0x0634, 0x0639, 0x0631, 0x0628, 0x064A, 0x061F},
		"egbpdaj6bu4bxfgehfvwxn"},
	testCase{ // (B) Chinese (simplified)
		[]rune{0x4ED6, 0x4EEC, 0x4E3A, 0x4EC0, 0x4E48, 0x4E0D, 0x8BF4, 0x4E2D, 0x6587},
		"ihqwcrb4cv8a8dqg056pqjye"},
	testCase{ //(C) Chinese (traditional)
		[]rune{0x4ED6, 0x5011, 0x7232, 0x4EC0, 0x9EBD, 0x4E0D, 0x8AAA, 0x4E2D, 0x6587},
		"ihqwctvzc91f659drss3x8bo0yb"},
	testCase{ // (D) Czech: Pro<ccaron>prost<ecaron>nemluv<iacute><ccaron>esky
		[]rune{0x0050, 0x0072, 0x006F, 0x010D, 0x0070, 0x0072, 0x006F, 0x0073, 0x0074, 0x011B, 0x006E, 0x0065, 0x006D, 0x006C, 0x0075, 0x0076, 0x00ED, 0x010D, 0x0065, 0x0073, 0x006B, 0x0079},
		"Proprostnemluvesky-uyb24dma41a"},
	testCase{ // (E) Hebrew:
		[]rune{0x05DC, 0x05DE, 0x05D4, 0x05D4, 0x05DD, 0x05E4, 0x05E9, 0x05D5, 0x05D8, 0x05DC, 0x05D0, 0x05DE, 0x05D3, 0x05D1, 0x05E8, 0x05D9, 0x05DD, 0x05E2, 0x05D1, 0x05E8, 0x05D9, 0x05EA},
		"4dbcagdahymbxekheh6e0a7fei0b"},
	testCase{ // (F) Hindi (Devanagari):
		[]rune{0x092F, 0x0939, 0x0932, 0x094B, 0x0917, 0x0939, 0x093F, 0x0928, 0x094D, 0x0926, 0x0940, 0x0915, 0x094D, 0x092F, 0x094B, 0x0902, 0x0928, 0x0939, 0x0940, 0x0902, 0x092C, 0x094B, 0x0932, 0x0938, 0x0915, 0x0924, 0x0947, 0x0939, 0x0948, 0x0902},
		"i1baa7eci9glrd9b2ae1bj0hfcgg6iyaf8o0a1dig0cd"},
	testCase{ // (G) Japanese (kanji and hiragana):
		[]rune{0x306A, 0x305C, 0x307F, 0x3093, 0x306A, 0x65E5, 0x672C, 0x8A9E, 0x3092, 0x8A71, 0x3057, 0x3066, 0x304F, 0x308C, 0x306A, 0x3044, 0x306E, 0x304B},
		"n8jok5ay5dzabd5bym9f0cm5685rrjetr6pdxa"},
	testCase{ // (H) Korean (Hangul syllables):
		[]rune{0xC138, 0xACC4, 0xC758, 0xBAA8, 0xB4E0, 0xC0AC, 0xB78C, 0xB4E4, 0xC774, 0xD55C, 0xAD6D, 0xC5B4, 0xB97C, 0xC774, 0xD574, 0xD55C, 0xB2E4, 0xBA74, 0xC5BC, 0xB9C8, 0xB098, 0xC88B, 0xC744, 0xAE4C},
		"989aomsvi5e83db1d2a355cv1e0vak1dwrv93d5xbh15a0dt30a5jpsd879ccm6fea98c"},
	testCase{ // (I) Russian (Cyrillic):
		[]rune{0x043F, 0x043E, 0x0447, 0x0435, 0x043C, 0x0443, 0x0436, 0x0435, 0x043E, 0x043D, 0x0438, 0x043D, 0x0435, 0x0433, 0x043E, 0x0432, 0x043E, 0x0440, 0x044F, 0x0442, 0x043F, 0x043E, 0x0440, 0x0443, 0x0441, 0x0441, 0x043A, 0x0438},
		"b1abfaaepdrnnbgefbadotcwatmq2g4l"}, // RFC says b1abfaaepdrnnbgefbaDotcwatmq2g4l but is mentioned in Errata ID: 3026
	testCase{ // (J) Spanish: Porqu<eacute>nopuedensimplementehablarenEspa<ntilde>ol
		[]rune{0x0050, 0x006F, 0x0072, 0x0071, 0x0075, 0x00E9, 0x006E, 0x006F, 0x0070, 0x0075, 0x0065, 0x0064, 0x0065, 0x006E, 0x0073, 0x0069, 0x006D, 0x0070, 0x006C, 0x0065, 0x006D, 0x0065, 0x006E, 0x0074, 0x0065, 0x0068, 0x0061, 0x0062, 0x006C, 0x0061, 0x0072, 0x0065, 0x006E, 0x0045, 0x0073, 0x0070, 0x0061, 0x00F1, 0x006F, 0x006C},
		"PorqunopuedensimplementehablarenEspaol-fmd56a"},
	testCase{ // (K) Vietnamese:  T<adotbelow>isaoh<odotbelow>kh<ocirc>ngth<ecirchookabove>ch<ihookabove>n<oacute>iti<ecircacute>ngVi<ecircdotbelow>t
		[]rune{0x0054, 0x1EA1, 0x0069, 0x0073, 0x0061, 0x006F, 0x0068, 0x1ECD, 0x006B, 0x0068, 0x00F4, 0x006E, 0x0067, 0x0074, 0x0068, 0x1EC3, 0x0063, 0x0068, 0x1EC9, 0x006E, 0x00F3, 0x0069, 0x0074, 0x0069, 0x1EBF, 0x006E, 0x0067, 0x0056, 0x0069, 0x1EC7, 0x0074},
		"TisaohkhngthchnitingVit-kjcr8268qyxafd2f1b9g"},
	testCase{ // (L) 3<nen>B<gumi><kinpachi><sensei>
		[]rune{0x0033, 0x5E74, 0x0042, 0x7D44, 0x91D1, 0x516B, 0x5148, 0x751F},
		"3B-ww4c5e180e575a65lsy2b"},
	testCase{ // (M) <amuro><namie>-with-SUPER-MONKEYS
		[]rune{0x5B89, 0x5BA4, 0x5948, 0x7F8E, 0x6075, 0x002D, 0x0077, 0x0069, 0x0074, 0x0068, 0x002D, 0x0053, 0x0055, 0x0050, 0x0045, 0x0052, 0x002D, 0x004D, 0x004F, 0x004E, 0x004B, 0x0045, 0x0059, 0x0053},
		"-with-SUPER-MONKEYS-pc58ag80a8qai00g7n9n"},
	testCase{ // (N) Hello-Another-Way-<sorezore><no><basho>
		[]rune{0x0048, 0x0065, 0x006C, 0x006C, 0x006F, 0x002D, 0x0041, 0x006E, 0x006F, 0x0074, 0x0068, 0x0065, 0x0072, 0x002D, 0x0057, 0x0061, 0x0079, 0x002D, 0x305D, 0x308C, 0x305E, 0x308C, 0x306E, 0x5834, 0x6240},
		"Hello-Another-Way--fc4qua05auwb3674vfr0b"},
	testCase{ // (O) <hitotsu><yane><no><shita>2
		[]rune{0x3072, 0x3068, 0x3064, 0x5C4B, 0x6839, 0x306E, 0x4E0B, 0x0032},
		"2-u9tlzr9756bt3uc0v"},
	testCase{ // (P) Maji<de>Koi<suru>5<byou><mae>
		[]rune{0x004D, 0x0061, 0x006A, 0x0069, 0x3067, 0x004B, 0x006F, 0x0069, 0x3059, 0x308B, 0x0035, 0x79D2, 0x524D},
		"MajiKoi5-783gue6qz075azm5e"},
	testCase{ // (Q) <pafii>de<runba>
		[]rune{0x30D1, 0x30D5, 0x30A3, 0x30FC, 0x0064, 0x0065, 0x30EB, 0x30F3, 0x30D0},
		"de-jg4avhby1noc0d"},
	testCase{ // (R) <sono><supiido><de>
		[]rune{0x305D, 0x306E, 0x30B9, 0x30D4, 0x30FC, 0x30C9, 0x3067},
		"d9juau41awczczp"},
	testCase{ // (S) -> $1.00 <-
		[]rune{0x002D, 0x003E, 0x0020, 0x0024, 0x0031, 0x002E, 0x0030, 0x0030, 0x0020, 0x003C, 0x002D},
		"-> $1.00 <--"},
}

func TestEncode(t *testing.T) {
	for _, tt := range punyTests {

		var out string
		var err error = nil

		// Encode the test case to an utf8 string
		in := make([]byte, 0, 4*len(tt.unicode))
		in_len := 0
		for i := 0; i < len(tt.unicode); i++ {
			rune := tt.unicode[i]
			str := make([]byte, utf8.RuneLen(rune))
			in_len += utf8.EncodeRune(str, rune)
			in = append(in, str...)
		}

		var o []byte
		o, err = Encode(in)
		out = string(o)
		if out != tt.punycode {
			if err == nil {
				t.Errorf("Encode(%v) = %v; want %v", hex8(in), out, tt.punycode)
			} else {
				t.Errorf("Encode(%v) results in %v error", hex8(in), err.Error())
			}
		}
	}
}

func TestDecode(t *testing.T) {
	for _, tt := range punyTests {

		// Encode the test case to an utf8 string
		ut := make([]byte, 0, 4*len(tt.unicode))
		ut_len := 0
		for i := 0; i < len(tt.unicode); i++ {
			rune := tt.unicode[i]
			str := make([]byte, utf8.RuneLen(rune))
			ut_len += utf8.EncodeRune(str, rune)
			ut = append(ut, str...)
		}

		var out []byte
		var err error = nil
		out, err = Decode([]byte(tt.punycode))
		if !reflect.DeepEqual(out, ut) {
			if err == nil {
				t.Errorf("Decode(%v) = %v; want %v", string(tt.punycode), hex8([]byte(out)), hex8([]byte(ut)))
			} else {
				t.Errorf("Decode(%v) results in %v error", string(tt.punycode), err.Error())
			}
		}
	}
}

type hex32 []int
type hex8 []uint8

func (h hex32) Format(f fmt.State, c int) {
	fmt.Fprint(f, "[")
	for i, v := range h {
		if i > 0 {
			fmt.Fprint(f, " ")
		}
		fmt.Fprintf(f, "%x", v)
	}
	fmt.Fprint(f, "]")
}

func (h hex8) Format(f fmt.State, c uint8) {
	fmt.Fprint(f, "[")
	for i, v := range h {
		if i > 0 {
			fmt.Fprint(f, " ")
		}
		fmt.Fprintf(f, "%x", v)
	}
	fmt.Fprint(f, "]")
}