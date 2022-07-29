package table

var x1d4 = []string{
	"A",      // 0x00
	"B",      // 0x01
	"C",      // 0x02
	"D",      // 0x03
	"E",      // 0x04
	"F",      // 0x05
	"G",      // 0x06
	"H",      // 0x07
	"I",      // 0x08
	"J",      // 0x09
	"K",      // 0x0a
	"L",      // 0x0b
	"M",      // 0x0c
	"N",      // 0x0d
	"O",      // 0x0e
	"P",      // 0x0f
	"Q",      // 0x10
	"R",      // 0x11
	"S",      // 0x12
	"T",      // 0x13
	"U",      // 0x14
	"V",      // 0x15
	"W",      // 0x16
	"X",      // 0x17
	"Y",      // 0x18
	"Z",      // 0x19
	"a",      // 0x1a
	"b",      // 0x1b
	"c",      // 0x1c
	"d",      // 0x1d
	"e",      // 0x1e
	"f",      // 0x1f
	"g",      // 0x20
	"h",      // 0x21
	"i",      // 0x22
	"j",      // 0x23
	"k",      // 0x24
	"l",      // 0x25
	"m",      // 0x26
	"n",      // 0x27
	"o",      // 0x28
	"p",      // 0x29
	"q",      // 0x2a
	"r",      // 0x2b
	"s",      // 0x2c
	"t",      // 0x2d
	"u",      // 0x2e
	"v",      // 0x2f
	"w",      // 0x30
	"x",      // 0x31
	"y",      // 0x32
	"z",      // 0x33
	"A",      // 0x34
	"B",      // 0x35
	"C",      // 0x36
	"D",      // 0x37
	"E",      // 0x38
	"F",      // 0x39
	"G",      // 0x3a
	"H",      // 0x3b
	"I",      // 0x3c
	"J",      // 0x3d
	"K",      // 0x3e
	"L",      // 0x3f
	"M",      // 0x40
	"N",      // 0x41
	"O",      // 0x42
	"P",      // 0x43
	"Q",      // 0x44
	"R",      // 0x45
	"S",      // 0x46
	"T",      // 0x47
	"U",      // 0x48
	"V",      // 0x49
	"W",      // 0x4a
	"X",      // 0x4b
	"Y",      // 0x4c
	"Z",      // 0x4d
	"a",      // 0x4e
	"b",      // 0x4f
	"c",      // 0x50
	"d",      // 0x51
	"e",      // 0x52
	"f",      // 0x53
	"g",      // 0x54
	"\uffff", // 0x55
	"i",      // 0x56
	"j",      // 0x57
	"k",      // 0x58
	"l",      // 0x59
	"m",      // 0x5a
	"n",      // 0x5b
	"o",      // 0x5c
	"p",      // 0x5d
	"q",      // 0x5e
	"r",      // 0x5f
	"s",      // 0x60
	"t",      // 0x61
	"u",      // 0x62
	"v",      // 0x63
	"w",      // 0x64
	"x",      // 0x65
	"y",      // 0x66
	"z",      // 0x67
	"A",      // 0x68
	"B",      // 0x69
	"C",      // 0x6a
	"D",      // 0x6b
	"E",      // 0x6c
	"F",      // 0x6d
	"G",      // 0x6e
	"H",      // 0x6f
	"I",      // 0x70
	"J",      // 0x71
	"K",      // 0x72
	"L",      // 0x73
	"M",      // 0x74
	"N",      // 0x75
	"O",      // 0x76
	"P",      // 0x77
	"Q",      // 0x78
	"R",      // 0x79
	"S",      // 0x7a
	"T",      // 0x7b
	"U",      // 0x7c
	"V",      // 0x7d
	"W",      // 0x7e
	"X",      // 0x7f
	"Y",      // 0x80
	"Z",      // 0x81
	"a",      // 0x82
	"b",      // 0x83
	"c",      // 0x84
	"d",      // 0x85
	"e",      // 0x86
	"f",      // 0x87
	"g",      // 0x88
	"h",      // 0x89
	"i",      // 0x8a
	"j",      // 0x8b
	"k",      // 0x8c
	"l",      // 0x8d
	"m",      // 0x8e
	"n",      // 0x8f
	"o",      // 0x90
	"p",      // 0x91
	"q",      // 0x92
	"r",      // 0x93
	"s",      // 0x94
	"t",      // 0x95
	"u",      // 0x96
	"v",      // 0x97
	"w",      // 0x98
	"x",      // 0x99
	"y",      // 0x9a
	"z",      // 0x9b
	"A",      // 0x9c
	"\uffff", // 0x9d
	"C",      // 0x9e
	"D",      // 0x9f
	"\uffff", // 0xa0
	"\uffff", // 0xa1
	"G",      // 0xa2
	"\uffff", // 0xa3
	"\uffff", // 0xa4
	"J",      // 0xa5
	"K",      // 0xa6
	"\uffff", // 0xa7
	"\uffff", // 0xa8
	"N",      // 0xa9
	"O",      // 0xaa
	"P",      // 0xab
	"Q",      // 0xac
	"\uffff", // 0xad
	"S",      // 0xae
	"T",      // 0xaf
	"U",      // 0xb0
	"V",      // 0xb1
	"W",      // 0xb2
	"X",      // 0xb3
	"Y",      // 0xb4
	"Z",      // 0xb5
	"a",      // 0xb6
	"b",      // 0xb7
	"c",      // 0xb8
	"d",      // 0xb9
	"\uffff", // 0xba
	"f",      // 0xbb
	"\uffff", // 0xbc
	"h",      // 0xbd
	"i",      // 0xbe
	"j",      // 0xbf
	"k",      // 0xc0
	"l",      // 0xc1
	"m",      // 0xc2
	"n",      // 0xc3
	"\uffff", // 0xc4
	"p",      // 0xc5
	"q",      // 0xc6
	"r",      // 0xc7
	"s",      // 0xc8
	"t",      // 0xc9
	"u",      // 0xca
	"v",      // 0xcb
	"w",      // 0xcc
	"x",      // 0xcd
	"y",      // 0xce
	"z",      // 0xcf
	"A",      // 0xd0
	"B",      // 0xd1
	"C",      // 0xd2
	"D",      // 0xd3
	"E",      // 0xd4
	"F",      // 0xd5
	"G",      // 0xd6
	"H",      // 0xd7
	"I",      // 0xd8
	"J",      // 0xd9
	"K",      // 0xda
	"L",      // 0xdb
	"M",      // 0xdc
	"N",      // 0xdd
	"O",      // 0xde
	"P",      // 0xdf
	"Q",      // 0xe0
	"R",      // 0xe1
	"S",      // 0xe2
	"T",      // 0xe3
	"U",      // 0xe4
	"V",      // 0xe5
	"W",      // 0xe6
	"X",      // 0xe7
	"Y",      // 0xe8
	"Z",      // 0xe9
	"a",      // 0xea
	"b",      // 0xeb
	"c",      // 0xec
	"d",      // 0xed
	"e",      // 0xee
	"f",      // 0xef
	"g",      // 0xf0
	"h",      // 0xf1
	"i",      // 0xf2
	"j",      // 0xf3
	"k",      // 0xf4
	"l",      // 0xf5
	"m",      // 0xf6
	"n",      // 0xf7
	"o",      // 0xf8
	"p",      // 0xf9
	"q",      // 0xfa
	"r",      // 0xfb
	"s",      // 0xfc
	"t",      // 0xfd
	"u",      // 0xfe
	"v",      // 0xff
}
