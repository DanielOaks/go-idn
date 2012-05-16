package stringprep

var _nameprep = &Profile{
	Map: []Table{_B1, _B2}
	Normalize: true
	Prohibit: []Table{_C12, _C22, _C3, _C4, _C5, _C6, _C7, _C8, _C9}
	
	ProfileElement{BIDI, nil},
	ProfileElement{BIDI_PROHIBIT_TABLE, _C8},
	ProfileElement{BIDI_RAL_TABLE, _D1},
	ProfileElement{BIDI_L_TABLE, _D2},
	ProfileElement{UNASSIGNED_TABLE, _A1},
}
