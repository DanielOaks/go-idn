// Copyright 2010 Hannes Baldursson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is part of go-idn

package stringprep

var Profiles = map[string]Profile{
	"Nameprep":     _nameprep,
	"Nodeprep":     _xmpp_nodeprep,
	"Resourceprep": _xmpp_resourceprep,
	"plain":        _plain, /* sasl-anon-00. */
	"trace":        _trace, /* sasl-anon-01,02,03. */
	"SASLprep":     _saslprep,
	"iSCSI":        _iscsi, /* IANA. */
}

// Nameprep - As descrirbed in RFC 3491: http://tools.ietf.org/html/rfc3491
var _nameprep = Profile{
	ProfileElement{MAP_TABLE, _B1},
	ProfileElement{MAP_TABLE, _B2},
	ProfileElement{NFKC, nil},
	ProfileElement{PROHIBIT_TABLE, _C12},
	ProfileElement{PROHIBIT_TABLE, _C22},
	ProfileElement{PROHIBIT_TABLE, _C3},
	ProfileElement{PROHIBIT_TABLE, _C4},
	ProfileElement{PROHIBIT_TABLE, _C5},
	ProfileElement{PROHIBIT_TABLE, _C6},
	ProfileElement{PROHIBIT_TABLE, _C7},
	ProfileElement{PROHIBIT_TABLE, _C8},
	ProfileElement{PROHIBIT_TABLE, _C9},
	ProfileElement{BIDI, nil},
	ProfileElement{BIDI_PROHIBIT_TABLE, _C8},
	ProfileElement{BIDI_RAL_TABLE, _D1},
	ProfileElement{BIDI_L_TABLE, _D2},
	ProfileElement{UNASSIGNED_TABLE, _A1},
}

// XMPP Nodeprep - As described in RFC 3920: http://tools.ietf.org/html/rfc3920#appendix-A
var _xmpp_nodeprep = Profile{
	ProfileElement{MAP_TABLE, _B1},
	ProfileElement{MAP_TABLE, _B2},
	ProfileElement{NFKC, nil},
	ProfileElement{PROHIBIT_TABLE, _C11},
	ProfileElement{PROHIBIT_TABLE, _C12},
	ProfileElement{PROHIBIT_TABLE, _C21},
	ProfileElement{PROHIBIT_TABLE, _C22},
	ProfileElement{PROHIBIT_TABLE, _C3},
	ProfileElement{PROHIBIT_TABLE, _C4},
	ProfileElement{PROHIBIT_TABLE, _C5},
	ProfileElement{PROHIBIT_TABLE, _C6},
	ProfileElement{PROHIBIT_TABLE, _C7},
	ProfileElement{PROHIBIT_TABLE, _C8},
	ProfileElement{PROHIBIT_TABLE, _C9},
	ProfileElement{PROHIBIT_TABLE, _nodeprep_prohibit},
	ProfileElement{BIDI, nil},
	ProfileElement{BIDI_PROHIBIT_TABLE, _C8},
	ProfileElement{BIDI_RAL_TABLE, _D1},
	ProfileElement{BIDI_L_TABLE, _D2},
	ProfileElement{UNASSIGNED_TABLE, _A1},
}

// XMPP Resourceprep - As described in RFC 3920: http://tools.ietf.org/html/rfc3920#appendix-B
var _xmpp_resourceprep = Profile{
	ProfileElement{MAP_TABLE, _B1},
	ProfileElement{NFKC, nil},
	ProfileElement{PROHIBIT_TABLE, _C12},
	ProfileElement{PROHIBIT_TABLE, _C21},
	ProfileElement{PROHIBIT_TABLE, _C22},
	ProfileElement{PROHIBIT_TABLE, _C3},
	ProfileElement{PROHIBIT_TABLE, _C4},
	ProfileElement{PROHIBIT_TABLE, _C5},
	ProfileElement{PROHIBIT_TABLE, _C6},
	ProfileElement{PROHIBIT_TABLE, _C7},
	ProfileElement{PROHIBIT_TABLE, _C8},
	ProfileElement{PROHIBIT_TABLE, _C9},
	ProfileElement{BIDI, nil},
	ProfileElement{BIDI_PROHIBIT_TABLE, _C8},
	ProfileElement{BIDI_RAL_TABLE, _D1},
	ProfileElement{BIDI_L_TABLE, _D2},
	ProfileElement{UNASSIGNED_TABLE, _A1},
}

var _plain = Profile{
	ProfileElement{PROHIBIT_TABLE, _C21},
	ProfileElement{PROHIBIT_TABLE, _C22},
	ProfileElement{PROHIBIT_TABLE, _C3},
	ProfileElement{PROHIBIT_TABLE, _C4},
	ProfileElement{PROHIBIT_TABLE, _C5},
	ProfileElement{PROHIBIT_TABLE, _C6},
	ProfileElement{PROHIBIT_TABLE, _C8},
	ProfileElement{PROHIBIT_TABLE, _C9},
	ProfileElement{BIDI, nil},
	ProfileElement{BIDI_PROHIBIT_TABLE, _C8},
	ProfileElement{BIDI_RAL_TABLE, _D1},
	ProfileElement{BIDI_L_TABLE, _D2},
}

// Trace - As described in RFC 4505: http://tools.ietf.org/html/rfc4505
var _trace = Profile{
	ProfileElement{PROHIBIT_TABLE, _C21},
	ProfileElement{PROHIBIT_TABLE, _C22},
	ProfileElement{PROHIBIT_TABLE, _C3},
	ProfileElement{PROHIBIT_TABLE, _C4},
	ProfileElement{PROHIBIT_TABLE, _C5},
	ProfileElement{PROHIBIT_TABLE, _C6},
	ProfileElement{PROHIBIT_TABLE, _C8},
	ProfileElement{PROHIBIT_TABLE, _C9},
	ProfileElement{BIDI, nil},
	ProfileElement{BIDI_PROHIBIT_TABLE, _C8},
	ProfileElement{BIDI_RAL_TABLE, _D1},
	ProfileElement{BIDI_L_TABLE, _D2},
}

// SASLprep - As described in RFC 4013: http://tools.ietf.org/html/rfc4013
var _saslprep = Profile{
	ProfileElement{MAP_TABLE, _saslprep_space_map},
	ProfileElement{MAP_TABLE, _B1},
	ProfileElement{NFKC, nil},
	ProfileElement{PROHIBIT_TABLE, _C12},
	ProfileElement{PROHIBIT_TABLE, _C21},
	ProfileElement{PROHIBIT_TABLE, _C22},
	ProfileElement{PROHIBIT_TABLE, _C3},
	ProfileElement{PROHIBIT_TABLE, _C4},
	ProfileElement{PROHIBIT_TABLE, _C5},
	ProfileElement{PROHIBIT_TABLE, _C6},
	ProfileElement{PROHIBIT_TABLE, _C7},
	ProfileElement{PROHIBIT_TABLE, _C8},
	ProfileElement{PROHIBIT_TABLE, _C9},
	ProfileElement{BIDI, nil},
	ProfileElement{BIDI_PROHIBIT_TABLE, _C8},
	ProfileElement{BIDI_RAL_TABLE, _D1},
	ProfileElement{BIDI_L_TABLE, _D2},
	ProfileElement{UNASSIGNED_TABLE, _A1},
}

// iSCSI - As described in RFC 3722: http://tools.ietf.org/html/rfc3722
var _iscsi = Profile{
	ProfileElement{MAP_TABLE, _B1},
	ProfileElement{MAP_TABLE, _B2},
	ProfileElement{NFKC, nil},
	ProfileElement{PROHIBIT_TABLE, _C11},
	ProfileElement{PROHIBIT_TABLE, _C12},
	ProfileElement{PROHIBIT_TABLE, _C21},
	ProfileElement{PROHIBIT_TABLE, _C22},
	ProfileElement{PROHIBIT_TABLE, _C3},
	ProfileElement{PROHIBIT_TABLE, _C4},
	ProfileElement{PROHIBIT_TABLE, _C5},
	ProfileElement{PROHIBIT_TABLE, _C6},
	ProfileElement{PROHIBIT_TABLE, _C7},
	ProfileElement{PROHIBIT_TABLE, _C8},
	ProfileElement{PROHIBIT_TABLE, _C9},
	ProfileElement{PROHIBIT_TABLE, _iscsi_prohibit},
	ProfileElement{BIDI, nil},
	ProfileElement{BIDI_PROHIBIT_TABLE, _C8},
	ProfileElement{BIDI_RAL_TABLE, _D1},
	ProfileElement{BIDI_L_TABLE, _D2},
	ProfileElement{UNASSIGNED_TABLE, _A1},
}

var _iscsi_prohibit = Table{
	TableElement{0x0000, 0x002C, d{}}, /* [ASCII CONTROL CHARACTERS and SPACE through ,] */
	TableElement{0x002F, 0x002F, d{}}, /* [ASCII /] */
	TableElement{0x003B, 0x0040, d{}}, /* [ASCII ; through @] */
	TableElement{0x005B, 0x0060, d{}}, /* [ASCII [ through `] */
	TableElement{0x007B, 0x007F, d{}}, /* [ASCII { through DEL] */
}

var _saslprep_space_map = Table{
	TableElement{0x0000A0, 0x0000A0, d{0x0020}}, /* 00A0; NO-BREAK SPACE */
	TableElement{0x001680, 0x001680, d{0x0020}}, /* 1680; OGHAM SPACE MARK */
	TableElement{0x002000, 0x002000, d{0x0020}}, /* 2000; EN QUAD */
	TableElement{0x002001, 0x002001, d{0x0020}}, /* 2001; EM QUAD */
	TableElement{0x002002, 0x002002, d{0x0020}}, /* 2002; EN SPACE */
	TableElement{0x002003, 0x002003, d{0x0020}}, /* 2003; EM SPACE */
	TableElement{0x002004, 0x002004, d{0x0020}}, /* 2004; THREE-PER-EM SPACE */
	TableElement{0x002005, 0x002005, d{0x0020}}, /* 2005; FOUR-PER-EM SPACE */
	TableElement{0x002006, 0x002006, d{0x0020}}, /* 2006; SIX-PER-EM SPACE */
	TableElement{0x002007, 0x002007, d{0x0020}}, /* 2007; FIGURE SPACE */
	TableElement{0x002008, 0x002008, d{0x0020}}, /* 2008; PUNCTUATION SPACE */
	TableElement{0x002009, 0x002009, d{0x0020}}, /* 2009; THIN SPACE */
	TableElement{0x00200A, 0x00200A, d{0x0020}}, /* 200A; HAIR SPACE */
	TableElement{0x00200B, 0x00200B, d{0x0020}}, /* 200B; ZERO WIDTH SPACE */
	TableElement{0x00202F, 0x00202F, d{0x0020}}, /* 202F; NARROW NO-BREAK SPACE */
	TableElement{0x00205F, 0x00205F, d{0x0020}}, /* 205F; MEDIUM MATHEMATICAL SPACE */
	TableElement{0x003000, 0x003000, d{0x0020}}, /* 3000; IDEOGRAPHIC SPACE */
}

var _nodeprep_prohibit = Table{
	TableElement{0x000022, 0x000022, d{}}, /* #x22 (") */
	TableElement{0x000026, 0x000026, d{}}, /* #x26 (&) */
	TableElement{0x000027, 0x000027, d{}}, /* #x27 (') */
	TableElement{0x00002F, 0x00002F, d{}}, /* #x2F (/) */
	TableElement{0x00003A, 0x00003A, d{}}, /* #x3A (:) */
	TableElement{0x00003C, 0x00003C, d{}}, /* #x3C (<) */
	TableElement{0x00003E, 0x00003E, d{}}, /* #x3E (>) */
	TableElement{0x000040, 0x000040, d{}}, /* #x40 (@) */
}
