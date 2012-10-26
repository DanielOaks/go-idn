package profile

type Profile struct {
	normalize bool
	checkBidi bool
	key       Trie
	val       []byte
}

func MkProfile(key Trie, val []byte, norm bool, bidi bool) *Profile {
	return &Profile{key, val}
}

func (p *Profile) Val(r rune) string {

}
