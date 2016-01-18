package stringprep

var (
	Profiles = map[string]Profile{
		"nameprep": Nameprep,
	}
)

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
