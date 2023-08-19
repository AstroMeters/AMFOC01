package freesans

import (
	"tinygo.org/x/tinyfont"
)

var BoldOblique12pt7b = tinyfont.Font{
	BBox: [4]int8{27, 28, -2, -22},
	Glyphs: []tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x0, Height: 0x0, XAdvance: 0x7, XOffset: 0, YOffset: 1, Bitmaps: []uint8{}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x7, Height: 0x11, XAdvance: 0x8, XOffset: 3, YOffset: -16, Bitmaps: []uint8{0x1c, 0x3c, 0x78, 0xe1, 0xc3, 0x8f, 0x1c, 0x38, 0x70, 0xc1, 0x83, 0x0, 0x1c, 0x78, 0xf0}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0xa, Height: 0x6, XAdvance: 0xb, XOffset: 4, YOffset: -17, Bitmaps: []uint8{0x71, 0xfc, 0xfe, 0x3b, 0x8e, 0xc3, 0x30, 0xc0}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0xf, Height: 0x10, XAdvance: 0xd, XOffset: 1, YOffset: -15, Bitmaps: []uint8{0x1, 0x8c, 0x7, 0x38, 0xc, 0x61, 0xff, 0xf3, 0xff, 0xe7, 0xff, 0x83, 0x9c, 0xe, 0x70, 0x1c, 0xe1, 0xff, 0xf3, 0xff, 0xc7, 0xff, 0x83, 0x18, 0xe, 0x70, 0x18, 0xc0, 0x73, 0x80}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0xf, Height: 0x15, XAdvance: 0xd, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x0, 0x40, 0x7, 0xf0, 0x3f, 0xf0, 0xff, 0xf3, 0xc9, 0xe7, 0xb3, 0xcf, 0x60, 0x1f, 0xc0, 0x3f, 0xc0, 0x3f, 0xe0, 0x1f, 0xe0, 0x1b, 0xe0, 0x33, 0xde, 0x47, 0xbc, 0x8f, 0x7f, 0x7c, 0x7f, 0xf0, 0x7f, 0x80, 0x18, 0x0, 0x20, 0x0, 0xc0, 0x0}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0x12, Height: 0x12, XAdvance: 0x15, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0x0, 0x1, 0x87, 0x80, 0xc3, 0xf0, 0x61, 0xfe, 0x10, 0xe1, 0x8c, 0x30, 0x66, 0xc, 0x3b, 0x3, 0xfc, 0x80, 0x7e, 0x60, 0xf, 0x30, 0x0, 0x18, 0x70, 0xc, 0x7e, 0x3, 0x1f, 0xc1, 0x8e, 0x30, 0xc3, 0x1c, 0x60, 0xfe, 0x18, 0x1f, 0x8c, 0x7, 0x80}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0xf, Height: 0x11, XAdvance: 0x11, XOffset: 2, YOffset: -16, Bitmaps: []uint8{0x1, 0xe0, 0x7, 0xf0, 0x1f, 0xe0, 0x79, 0xc0, 0xf3, 0x81, 0xee, 0x1, 0xf8, 0x1, 0xe0, 0x1f, 0xc6, 0x7b, 0xdd, 0xe3, 0xf7, 0x87, 0xef, 0x7, 0x9f, 0x1f, 0x3f, 0xff, 0x3f, 0xde, 0x3f, 0x1c}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x4, Height: 0x6, XAdvance: 0x6, XOffset: 4, YOffset: -17, Bitmaps: []uint8{0x7f, 0xee, 0xcc}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x9, Height: 0x16, XAdvance: 0x8, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x3, 0x83, 0x81, 0x81, 0xc1, 0xc0, 0xe0, 0xe0, 0x70, 0x70, 0x38, 0x3c, 0x1c, 0xe, 0x7, 0x3, 0x81, 0xc0, 0xe0, 0x70, 0x18, 0xe, 0x7, 0x1, 0x80}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x9, Height: 0x16, XAdvance: 0x8, XOffset: -1, YOffset: -16, Bitmaps: []uint8{0x6, 0x3, 0x81, 0xc0, 0x60, 0x38, 0x1c, 0xe, 0x7, 0x3, 0x81, 0xc0, 0xe0, 0xe0, 0x70, 0x38, 0x38, 0x1c, 0x1c, 0xe, 0xe, 0x6, 0x7, 0x7, 0x0}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0x8, Height: 0x8, XAdvance: 0x9, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0xc, 0xc, 0x4f, 0xff, 0x1c, 0x3c, 0x6c, 0x44}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0xc, Height: 0xb, XAdvance: 0xe, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x3, 0x80, 0x38, 0x7, 0x0, 0x70, 0x7f, 0xff, 0xff, 0xff, 0xf0, 0xe0, 0xe, 0x0, 0xe0, 0xc, 0x0}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x5, Height: 0x7, XAdvance: 0x7, XOffset: 1, YOffset: -2, Bitmaps: []uint8{0x7b, 0xdc, 0x23, 0x33, 0x0}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x7, Height: 0x3, XAdvance: 0x8, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0x7f, 0xff, 0xf0}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x4, Height: 0x3, XAdvance: 0x7, XOffset: 2, YOffset: -2, Bitmaps: []uint8{0x7f, 0xe0}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0xa, Height: 0x11, XAdvance: 0x7, XOffset: 0, YOffset: -16, Bitmaps: []uint8{0x0, 0xc0, 0x30, 0x18, 0x4, 0x3, 0x0, 0x80, 0x60, 0x10, 0xc, 0x2, 0x1, 0x80, 0x40, 0x30, 0x8, 0x6, 0x1, 0x0, 0xc0, 0x0}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0xd, Height: 0x11, XAdvance: 0xd, XOffset: 2, YOffset: -16, Bitmaps: []uint8{0x3, 0xc0, 0x7f, 0x87, 0xfc, 0x78, 0xf3, 0xc7, 0xbc, 0x3d, 0xe1, 0xef, 0xf, 0xf0, 0x7f, 0x87, 0xbc, 0x3d, 0xe1, 0xef, 0x1e, 0x78, 0xf3, 0xff, 0xf, 0xf0, 0x3e, 0x0}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0x9, Height: 0x11, XAdvance: 0xd, XOffset: 4, YOffset: -16, Bitmaps: []uint8{0x3, 0x83, 0x83, 0xcf, 0xef, 0xf0, 0x78, 0x38, 0x1c, 0xe, 0xf, 0x7, 0x3, 0x81, 0xc1, 0xe0, 0xf0, 0x70, 0x38, 0x0}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0xf, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x3, 0xf0, 0xf, 0xf8, 0x7f, 0xf8, 0xf1, 0xf3, 0xc1, 0xe7, 0x83, 0xc0, 0x7, 0x80, 0x1e, 0x0, 0x78, 0x3, 0xe0, 0xf, 0x0, 0x7c, 0x1, 0xe0, 0x7, 0x0, 0x1f, 0xfc, 0x3f, 0xf8, 0xff, 0xf0}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0xd, Height: 0x11, XAdvance: 0xd, XOffset: 2, YOffset: -16, Bitmaps: []uint8{0x7, 0xe0, 0xff, 0x8f, 0xfe, 0xf8, 0xf7, 0x87, 0x80, 0x78, 0xf, 0x80, 0xfc, 0x7, 0xe0, 0xf, 0x80, 0x3c, 0x1, 0xef, 0xf, 0x78, 0xf3, 0xff, 0x8f, 0xf8, 0x3f, 0x0}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0xd, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x0, 0x78, 0x7, 0xc0, 0x7e, 0x3, 0xf0, 0x37, 0x3, 0x38, 0x31, 0xc3, 0x9e, 0x38, 0xf1, 0x87, 0x1f, 0xfe, 0xff, 0xf7, 0xff, 0x80, 0xf0, 0x7, 0x0, 0x38, 0x3, 0xc0}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0xe, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x7, 0xfc, 0x1f, 0xf0, 0xff, 0xc3, 0x0, 0x1c, 0x0, 0x7f, 0x81, 0xff, 0xf, 0xfe, 0x38, 0xf8, 0x1, 0xe0, 0x7, 0x80, 0x1e, 0xf0, 0xf3, 0xc7, 0xcf, 0xfe, 0x1f, 0xf0, 0x3f, 0x0}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0xd, Height: 0x11, XAdvance: 0xd, XOffset: 2, YOffset: -16, Bitmaps: []uint8{0x3, 0xe0, 0x7f, 0x87, 0xfe, 0x78, 0xf3, 0xc0, 0x3d, 0xe1, 0xff, 0x8f, 0xfe, 0xf8, 0xf7, 0xc7, 0xbc, 0x3d, 0xe1, 0xef, 0x1e, 0x7c, 0xf3, 0xff, 0xf, 0xf0, 0x1f, 0x0}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0xd, Height: 0x11, XAdvance: 0xd, XOffset: 3, YOffset: -16, Bitmaps: []uint8{0x7f, 0xfb, 0xff, 0xdf, 0xfe, 0x0, 0xe0, 0xe, 0x0, 0xe0, 0xe, 0x0, 0xe0, 0xf, 0x0, 0x70, 0x7, 0x0, 0x78, 0x3, 0x80, 0x3c, 0x1, 0xc0, 0xe, 0x0, 0xf0, 0x0}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0xe, Height: 0x11, XAdvance: 0xd, XOffset: 1, YOffset: -16, Bitmaps: []uint8{0x3, 0xf0, 0x1f, 0xe0, 0xff, 0xc7, 0x8f, 0x1c, 0x3c, 0x71, 0xe0, 0xff, 0x3, 0xf8, 0x3f, 0xf1, 0xf1, 0xe7, 0x87, 0xbc, 0x1e, 0xf0, 0x7b, 0xe3, 0xcf, 0xff, 0x1f, 0xf8, 0x1f, 0x80}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0xe, Height: 0x11, XAdvance: 0xd, XOffset: 2, YOffset: -16, Bitmaps: []uint8{0x3, 0xe0, 0x3f, 0xe1, 0xff, 0x8f, 0x9f, 0x3c, 0x3d, 0xe0, 0xf7, 0x83, 0xde, 0x1f, 0x78, 0xfd, 0xff, 0xe3, 0xff, 0x87, 0xde, 0x0, 0xf3, 0xc7, 0x8f, 0xfe, 0x1f, 0xf0, 0x3f, 0x0}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x6, Height: 0xc, XAdvance: 0x8, XOffset: 3, YOffset: -11, Bitmaps: []uint8{0x1c, 0xf3, 0x80, 0x0, 0x0, 0x0, 0x1, 0xcf, 0x38}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x7, Height: 0x10, XAdvance: 0x8, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0xe, 0x3c, 0x70, 0x0, 0x0, 0x0, 0x0, 0x0, 0xf1, 0xe3, 0x81, 0x6, 0x18, 0x60}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0xd, Height: 0xc, XAdvance: 0xe, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0x0, 0x0, 0x1, 0xc0, 0x7e, 0x1f, 0xe7, 0xf8, 0x7e, 0x3, 0xe0, 0x1f, 0xe0, 0x3f, 0xc0, 0x7f, 0x0, 0x78, 0x0, 0xc0}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0xe, Height: 0x9, XAdvance: 0xe, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x3f, 0xfc, 0xff, 0xf3, 0xff, 0x80, 0x0, 0x0, 0x0, 0x0, 0x7, 0xff, 0x9f, 0xfc, 0x7f, 0xf0}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0xd, Height: 0xc, XAdvance: 0xe, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0x30, 0x1, 0xe0, 0xf, 0xe0, 0x3f, 0xc0, 0x7f, 0x80, 0x7c, 0x7, 0xe1, 0xfe, 0x7f, 0x87, 0xe0, 0x38, 0x0, 0x0, 0x0}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0xd, Height: 0x12, XAdvance: 0xf, XOffset: 4, YOffset: -17, Bitmaps: []uint8{0xf, 0xc1, 0xff, 0x8f, 0xfc, 0xf1, 0xff, 0x7, 0xf0, 0x3c, 0x1, 0xe0, 0x1e, 0x1, 0xe0, 0x3e, 0x3, 0xe0, 0x1c, 0x1, 0xc0, 0xe, 0x0, 0x0, 0x7, 0x80, 0x3c, 0x1, 0xc0, 0x0}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0x16, Height: 0x15, XAdvance: 0x17, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x0, 0x3f, 0x80, 0x3, 0xff, 0x80, 0x3c, 0xf, 0x1, 0xc0, 0xe, 0xe, 0x0, 0x1c, 0x70, 0xf7, 0x73, 0x87, 0xf8, 0xcc, 0x31, 0xe3, 0x61, 0x87, 0xd, 0x8c, 0x1c, 0x3c, 0x30, 0x61, 0xb1, 0x81, 0x86, 0xc6, 0xc, 0x3b, 0x18, 0x71, 0xcc, 0x63, 0xce, 0x31, 0xfb, 0xf0, 0xe3, 0xcf, 0x1, 0xc0, 0x0, 0x3, 0xc0, 0xc0, 0x7, 0xff, 0x0, 0x7, 0xf0, 0x0}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0x11, Height: 0x12, XAdvance: 0x11, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0x0, 0x3e, 0x0, 0x3f, 0x0, 0x1f, 0x80, 0x1f, 0xc0, 0xf, 0xe0, 0xf, 0xf0, 0x7, 0x7c, 0x7, 0x1e, 0x3, 0x8f, 0x3, 0x87, 0x83, 0xc3, 0xc1, 0xff, 0xe1, 0xff, 0xf0, 0xff, 0xfc, 0xf0, 0x1e, 0x70, 0xf, 0x78, 0x7, 0xb8, 0x3, 0xc0}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0x11, Height: 0x12, XAdvance: 0x11, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xf, 0xfe, 0xf, 0xff, 0x87, 0xff, 0xe3, 0xc0, 0xf1, 0xc0, 0x78, 0xe0, 0x3c, 0xf0, 0x3c, 0x7f, 0xfc, 0x3f, 0xfc, 0x1f, 0xff, 0xe, 0x7, 0xcf, 0x1, 0xe7, 0x80, 0xf3, 0x80, 0x79, 0xc0, 0x79, 0xff, 0xf8, 0xff, 0xfc, 0x7f, 0xf8, 0x0}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0x11, Height: 0x12, XAdvance: 0x11, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0x1, 0xf8, 0x3, 0xff, 0x3, 0xff, 0xc3, 0xe1, 0xf3, 0xc0, 0x79, 0xe0, 0x3d, 0xe0, 0x0, 0xf0, 0x0, 0xf0, 0x0, 0x78, 0x0, 0x3c, 0x0, 0x1e, 0x0, 0xf, 0x0, 0xe7, 0x80, 0xf3, 0xe0, 0xf0, 0xff, 0xf8, 0x3f, 0xf0, 0x7, 0xe0, 0x0}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0x11, Height: 0x12, XAdvance: 0x11, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x1f, 0xfc, 0xf, 0xff, 0x87, 0xff, 0xc3, 0x81, 0xf1, 0xc0, 0x79, 0xe0, 0x3c, 0xf0, 0x1e, 0x78, 0xf, 0x38, 0x7, 0x9c, 0x3, 0xde, 0x3, 0xcf, 0x1, 0xe7, 0x81, 0xf3, 0x80, 0xf1, 0xc1, 0xf1, 0xff, 0xf0, 0xff, 0xf0, 0x7f, 0xe0, 0x0}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0x10, Height: 0x12, XAdvance: 0x10, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xf, 0xff, 0x1f, 0xff, 0x1f, 0xff, 0x1c, 0x0, 0x1c, 0x0, 0x3c, 0x0, 0x3c, 0x0, 0x3f, 0xfc, 0x3f, 0xfc, 0x3f, 0xfc, 0x78, 0x0, 0x78, 0x0, 0x78, 0x0, 0x70, 0x0, 0x70, 0x0, 0xff, 0xf8, 0xff, 0xf8, 0xff, 0xf8}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0x10, Height: 0x12, XAdvance: 0xf, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x1f, 0xff, 0x1f, 0xfe, 0x1f, 0xfe, 0x1c, 0x0, 0x1c, 0x0, 0x3c, 0x0, 0x3c, 0x0, 0x3f, 0xf8, 0x3f, 0xf8, 0x3f, 0xf8, 0x78, 0x0, 0x78, 0x0, 0x78, 0x0, 0x70, 0x0, 0xf0, 0x0, 0xf0, 0x0, 0xf0, 0x0, 0xe0, 0x0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0x11, Height: 0x12, XAdvance: 0x13, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0x1, 0xfc, 0x3, 0xff, 0x3, 0xff, 0xc3, 0xe0, 0xf3, 0xc0, 0x39, 0xc0, 0x1, 0xe0, 0x0, 0xf0, 0x0, 0xf0, 0x7f, 0x78, 0x3f, 0xbc, 0x1f, 0xde, 0x1, 0xcf, 0x0, 0xe7, 0xc0, 0xf1, 0xf0, 0xf8, 0xff, 0xfc, 0x3f, 0xec, 0x7, 0xe6, 0x0}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0x11, Height: 0x12, XAdvance: 0x11, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x1e, 0x3, 0x8f, 0x1, 0xc7, 0x1, 0xe3, 0x80, 0xf3, 0xc0, 0x79, 0xe0, 0x38, 0xf0, 0x1c, 0x7f, 0xfe, 0x3f, 0xff, 0x3f, 0xff, 0x9e, 0x3, 0x8f, 0x1, 0xc7, 0x1, 0xe3, 0x80, 0xf3, 0xc0, 0x71, 0xe0, 0x38, 0xf0, 0x3c, 0x70, 0x1e, 0x0}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x7, Height: 0x12, XAdvance: 0x7, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x1e, 0x3c, 0x78, 0xe1, 0xc7, 0x8f, 0x1e, 0x38, 0x71, 0xe3, 0xc7, 0x8e, 0x1c, 0x78, 0xf1, 0xe0}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0xe, Height: 0x12, XAdvance: 0xd, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x0, 0x1c, 0x0, 0xf0, 0x3, 0xc0, 0xf, 0x0, 0x38, 0x0, 0xe0, 0x7, 0x80, 0x1e, 0x0, 0x78, 0x1, 0xc0, 0x7, 0x3c, 0x3c, 0xf0, 0xf3, 0xc3, 0x8f, 0x1e, 0x3f, 0xf8, 0x7f, 0xc0, 0xfc, 0x0}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0x12, Height: 0x12, XAdvance: 0x11, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x1e, 0x7, 0xc7, 0x83, 0xe1, 0xe1, 0xe0, 0x70, 0xf0, 0x1c, 0x78, 0xf, 0x3c, 0x3, 0xde, 0x0, 0xff, 0x0, 0x3f, 0xc0, 0xf, 0xf0, 0x7, 0xde, 0x1, 0xe7, 0xc0, 0x78, 0xf0, 0x1c, 0x3e, 0xf, 0x7, 0x83, 0xc0, 0xf0, 0xf0, 0x3c, 0x38, 0x7, 0x80}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0xd, Height: 0x12, XAdvance: 0xf, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xe, 0x0, 0xf0, 0x7, 0x80, 0x3c, 0x1, 0xc0, 0xe, 0x0, 0xf0, 0x7, 0x80, 0x38, 0x1, 0xc0, 0x1e, 0x0, 0xf0, 0x7, 0x80, 0x38, 0x1, 0xc0, 0x1f, 0xfe, 0xff, 0xf7, 0xff, 0x80}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0x14, Height: 0x12, XAdvance: 0x14, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x1f, 0x3, 0xf1, 0xf0, 0x3f, 0x1f, 0x7, 0xf1, 0xf0, 0x7f, 0x3f, 0xf, 0xe3, 0xf0, 0xee, 0x3b, 0x1e, 0xe3, 0xb1, 0xde, 0x3b, 0x1d, 0xe7, 0xb3, 0x9c, 0x7b, 0x39, 0xc7, 0x37, 0x9c, 0x73, 0x73, 0xcf, 0x3f, 0x3c, 0xf3, 0xe3, 0x8f, 0x3e, 0x38, 0xe3, 0xc3, 0x8e, 0x3c, 0x78}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0x12, Height: 0x12, XAdvance: 0x11, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x1e, 0x3, 0x87, 0xc0, 0xe1, 0xf0, 0x38, 0x7c, 0x1e, 0x1f, 0x87, 0x8f, 0xe1, 0xc3, 0xb8, 0x70, 0xef, 0x1c, 0x39, 0xcf, 0x1e, 0x73, 0xc7, 0x8e, 0xe1, 0xc3, 0xb8, 0x70, 0xee, 0x1c, 0x1f, 0x8f, 0x7, 0xe3, 0xc1, 0xf0, 0xe0, 0x3c, 0x38, 0xf, 0x0}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0x11, Height: 0x12, XAdvance: 0x13, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0x1, 0xf8, 0x3, 0xff, 0x3, 0xff, 0xc3, 0xe3, 0xe3, 0xc0, 0xf9, 0xe0, 0x3d, 0xe0, 0x1e, 0xf0, 0xf, 0xf0, 0x7, 0xf8, 0x3, 0xfc, 0x3, 0xde, 0x1, 0xef, 0x0, 0xf7, 0xc0, 0xf1, 0xf0, 0xf0, 0xff, 0xf0, 0x3f, 0xf0, 0x7, 0xe0, 0x0}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0x10, Height: 0x12, XAdvance: 0x10, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x1f, 0xfc, 0x1f, 0xfe, 0x1f, 0xff, 0x1c, 0x1f, 0x1c, 0xf, 0x3c, 0xf, 0x3c, 0xf, 0x3c, 0x1e, 0x3f, 0xfc, 0x3f, 0xfc, 0x7f, 0xf0, 0x78, 0x0, 0x78, 0x0, 0x70, 0x0, 0x70, 0x0, 0xf0, 0x0, 0xf0, 0x0, 0xf0, 0x0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0x11, Height: 0x13, XAdvance: 0x13, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0x1, 0xf8, 0x3, 0xff, 0x3, 0xff, 0xc3, 0xe3, 0xe3, 0xc0, 0xf9, 0xc0, 0x3d, 0xe0, 0x1e, 0xf0, 0xf, 0xf0, 0x7, 0xf8, 0x3, 0xfc, 0x3, 0xde, 0x9, 0xef, 0xe, 0xe7, 0xc7, 0xf1, 0xf1, 0xf0, 0xff, 0xf8, 0x3f, 0xfe, 0x7, 0xe6, 0x0, 0x2, 0x0}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0x11, Height: 0x12, XAdvance: 0x11, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xf, 0xfe, 0xf, 0xff, 0x87, 0xff, 0xe3, 0x81, 0xf1, 0xc0, 0x78, 0xe0, 0x3c, 0xf0, 0x1c, 0x78, 0x1e, 0x3f, 0xfc, 0x1f, 0xfc, 0x1f, 0xff, 0x8f, 0x3, 0xc7, 0x81, 0xe3, 0x80, 0xf1, 0xc0, 0xf1, 0xe0, 0x78, 0xf0, 0x3c, 0x78, 0x1f, 0x0}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0x10, Height: 0x12, XAdvance: 0x10, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x3, 0xf8, 0xf, 0xfe, 0x1f, 0xff, 0x1e, 0x1f, 0x3c, 0xf, 0x3c, 0xf, 0x3c, 0x0, 0x3f, 0x0, 0x1f, 0xf0, 0xf, 0xfc, 0x1, 0xfe, 0x0, 0x3e, 0xf0, 0x1e, 0xf0, 0x1e, 0xf8, 0x3c, 0x7f, 0xf8, 0x7f, 0xf0, 0x1f, 0xc0}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0xf, Height: 0x12, XAdvance: 0xf, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0x7f, 0xfe, 0xff, 0xfd, 0xff, 0xf8, 0x1c, 0x0, 0x78, 0x0, 0xf0, 0x1, 0xe0, 0x3, 0x80, 0x7, 0x0, 0x1e, 0x0, 0x3c, 0x0, 0x78, 0x0, 0xe0, 0x1, 0xc0, 0x7, 0x80, 0xf, 0x0, 0x1e, 0x0, 0x38, 0x0}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0x10, Height: 0x12, XAdvance: 0x11, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0x1e, 0x7, 0x1c, 0xf, 0x3c, 0xf, 0x3c, 0xf, 0x3c, 0xe, 0x38, 0xe, 0x78, 0x1e, 0x78, 0x1e, 0x78, 0x1e, 0x78, 0x1c, 0x70, 0x1c, 0xf0, 0x3c, 0xf0, 0x3c, 0xf0, 0x38, 0xf8, 0x78, 0xff, 0xf0, 0x7f, 0xe0, 0x1f, 0x80}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0xf, Height: 0x12, XAdvance: 0x10, XOffset: 4, YOffset: -17, Bitmaps: []uint8{0xf0, 0x1f, 0xe0, 0x39, 0xc0, 0xf3, 0x81, 0xc7, 0x7, 0x8e, 0xe, 0x1c, 0x3c, 0x3c, 0x70, 0x79, 0xe0, 0xf3, 0x80, 0xef, 0x1, 0xdc, 0x3, 0xb8, 0x7, 0xe0, 0xf, 0x80, 0x1f, 0x0, 0x3c, 0x0, 0x78, 0x0}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0x15, Height: 0x12, XAdvance: 0x17, XOffset: 4, YOffset: -17, Bitmaps: []uint8{0xf0, 0x70, 0x7f, 0x87, 0x83, 0xfc, 0x3c, 0x3d, 0xe1, 0xe1, 0xef, 0x1f, 0xe, 0x78, 0xd8, 0xf3, 0xc6, 0xc7, 0xe, 0x76, 0x78, 0x73, 0x33, 0x83, 0xb9, 0x9c, 0x1d, 0xcd, 0xc0, 0xec, 0x6e, 0x7, 0xe3, 0xe0, 0x3e, 0x1f, 0x1, 0xf0, 0xf0, 0xf, 0x87, 0x80, 0x78, 0x38, 0x3, 0xc1, 0xc0, 0x0}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0x12, Height: 0x12, XAdvance: 0x10, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xf, 0x3, 0xc3, 0xc1, 0xe0, 0xf8, 0xf0, 0x1e, 0x78, 0x7, 0x9e, 0x0, 0xff, 0x0, 0x3f, 0x80, 0xf, 0xc0, 0x1, 0xe0, 0x0, 0xf8, 0x0, 0x3f, 0x0, 0x1f, 0xc0, 0xf, 0xf0, 0x7, 0x9e, 0x3, 0xc7, 0x80, 0xf0, 0xf0, 0x78, 0x3c, 0x3c, 0xf, 0x80}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0xf, Height: 0x12, XAdvance: 0x10, XOffset: 4, YOffset: -17, Bitmaps: []uint8{0x78, 0x1e, 0xf0, 0x79, 0xe0, 0xf3, 0xc3, 0xc3, 0xcf, 0x7, 0x9e, 0xf, 0x78, 0xf, 0xe0, 0x1f, 0x80, 0x3f, 0x0, 0x3c, 0x0, 0x70, 0x0, 0xe0, 0x3, 0xc0, 0x7, 0x80, 0xf, 0x0, 0x1c, 0x0, 0x38, 0x0}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0x11, Height: 0x12, XAdvance: 0xf, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0x1f, 0xff, 0xf, 0xff, 0x87, 0xff, 0xc0, 0x3, 0xc0, 0x3, 0xe0, 0x3, 0xe0, 0x3, 0xe0, 0x3, 0xe0, 0x1, 0xe0, 0x1, 0xe0, 0x1, 0xe0, 0x1, 0xe0, 0x1, 0xe0, 0x1, 0xe0, 0x1, 0xe0, 0x1, 0xff, 0xf0, 0xff, 0xf8, 0x7f, 0xfc, 0x0}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0xa, Height: 0x17, XAdvance: 0x8, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xf, 0xc3, 0xf0, 0xfc, 0x38, 0x1e, 0x7, 0x1, 0xc0, 0x70, 0x1c, 0xf, 0x3, 0x80, 0xe0, 0x38, 0xe, 0x7, 0x1, 0xc0, 0x70, 0x1c, 0xf, 0x3, 0x80, 0xfc, 0x3f, 0xf, 0xc0}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x4, Height: 0x17, XAdvance: 0x7, XOffset: 3, YOffset: -22, Bitmaps: []uint8{0x8, 0x88, 0xc4, 0x44, 0x66, 0x66, 0x66, 0x62, 0x22, 0x33, 0x33, 0x30}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0xa, Height: 0x17, XAdvance: 0x8, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0xf, 0xc3, 0xf0, 0xfc, 0x7, 0x3, 0xc0, 0xe0, 0x38, 0xe, 0x3, 0x81, 0xc0, 0x70, 0x1c, 0x7, 0x3, 0xc0, 0xe0, 0x38, 0xe, 0x3, 0x81, 0xe0, 0x70, 0xfc, 0x3f, 0xf, 0xc0}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0xb, Height: 0xb, XAdvance: 0xe, XOffset: 3, YOffset: -16, Bitmaps: []uint8{0x3, 0x80, 0xf0, 0x1e, 0x7, 0xe1, 0xdc, 0x3b, 0x8e, 0x71, 0x86, 0x70, 0xfc, 0x1f, 0x83, 0x80}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0xf, Height: 0x2, XAdvance: 0xd, XOffset: -2, YOffset: 4, Bitmaps: []uint8{0x7f, 0xfe, 0xff, 0xfc}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x4, Height: 0x3, XAdvance: 0x8, XOffset: 4, YOffset: -17, Bitmaps: []uint8{0xe6, 0x30}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0xd, Height: 0xd, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x7, 0xe0, 0xff, 0x8f, 0xfe, 0x70, 0xe0, 0x7, 0x3, 0xf8, 0xff, 0xcf, 0x9e, 0xf0, 0xf7, 0x8f, 0x3f, 0xf8, 0xff, 0xc3, 0xdf, 0x0}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0xf, Height: 0x12, XAdvance: 0xf, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xe, 0x0, 0x1c, 0x0, 0x38, 0x0, 0xf0, 0x1, 0xe0, 0x3, 0x9f, 0x7, 0xff, 0xf, 0xff, 0x3e, 0x3e, 0x78, 0x3c, 0xf0, 0x79, 0xc0, 0xf3, 0x81, 0xef, 0x7, 0x9f, 0x1f, 0x3f, 0xfc, 0x7f, 0xf0, 0xef, 0x80}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0xd, Height: 0xd, XAdvance: 0xd, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x7, 0xc0, 0xff, 0x8f, 0xfe, 0xf8, 0xf7, 0x87, 0xb8, 0x3, 0xc0, 0x1e, 0x0, 0xf0, 0xf7, 0x8f, 0x1f, 0xf8, 0xff, 0x81, 0xf0, 0x0}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0xf, Height: 0x12, XAdvance: 0xf, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x0, 0x1e, 0x0, 0x38, 0x0, 0x70, 0x0, 0xe0, 0x3, 0xc0, 0xf7, 0x87, 0xfe, 0x1f, 0xfc, 0x7c, 0x78, 0xf0, 0x73, 0xc0, 0xe7, 0x81, 0x8f, 0x7, 0x1e, 0xe, 0x3e, 0x3c, 0x7f, 0xf8, 0x7f, 0xe0, 0x7d, 0xc0}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0xd, Height: 0xd, XAdvance: 0xd, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x7, 0xc0, 0xff, 0x8f, 0xfe, 0xf0, 0xf7, 0x87, 0xff, 0xff, 0xff, 0xfe, 0x0, 0xf0, 0x7, 0xc7, 0x9f, 0xf8, 0xff, 0x81, 0xf0, 0x0}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0x9, Height: 0x12, XAdvance: 0x8, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x7, 0x87, 0xc7, 0xe3, 0xc1, 0xc3, 0xf9, 0xfc, 0x78, 0x3c, 0x1c, 0xe, 0x7, 0x7, 0x83, 0x81, 0xc0, 0xe0, 0xf0, 0x78, 0x0}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0xf, Height: 0x12, XAdvance: 0xf, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x3, 0xde, 0x1f, 0xf8, 0x7f, 0xf1, 0xf1, 0xe3, 0xc1, 0xcf, 0x3, 0x9e, 0x6, 0x3c, 0xc, 0x78, 0x38, 0xf8, 0xf1, 0xff, 0xc1, 0xff, 0x81, 0xf7, 0x0, 0xe, 0x3c, 0x3c, 0x78, 0xf0, 0x7f, 0xc0, 0x7e, 0x0}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0xe, Height: 0x12, XAdvance: 0xf, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0x1e, 0x0, 0x70, 0x1, 0xc0, 0x7, 0x0, 0x3c, 0x0, 0xf7, 0xc3, 0xbf, 0x8f, 0xff, 0x3c, 0x3d, 0xe0, 0xe7, 0x83, 0x9c, 0xe, 0x70, 0x79, 0xc1, 0xef, 0x7, 0x3c, 0x1c, 0xe0, 0x73, 0x83, 0xc0}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x7, Height: 0x12, XAdvance: 0x7, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xe, 0x3c, 0x70, 0x0, 0x3, 0x8f, 0x1e, 0x38, 0x71, 0xe3, 0xc7, 0xe, 0x1c, 0x78, 0xf1, 0xc0}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0xa, Height: 0x17, XAdvance: 0x7, XOffset: -1, YOffset: -17, Bitmaps: []uint8{0x3, 0xc0, 0xe0, 0x38, 0x0, 0x0, 0x1, 0xe0, 0x70, 0x1c, 0x7, 0x3, 0xc0, 0xf0, 0x38, 0xe, 0x3, 0x81, 0xe0, 0x70, 0x1c, 0x7, 0x3, 0xc0, 0xf0, 0xf8, 0x3e, 0xf, 0x0}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0xf, Height: 0x12, XAdvance: 0xd, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xe, 0x0, 0x1c, 0x0, 0x38, 0x0, 0xf0, 0x1, 0xe0, 0x3, 0x87, 0x87, 0x1e, 0xe, 0x78, 0x3d, 0xe0, 0x7f, 0x80, 0xfe, 0x1, 0xfe, 0x3, 0xfc, 0xf, 0x38, 0x1e, 0x78, 0x38, 0xf0, 0x70, 0xf0, 0xe1, 0xe0}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x7, Height: 0x12, XAdvance: 0x7, XOffset: 2, YOffset: -17, Bitmaps: []uint8{0xe, 0x3c, 0x78, 0xe1, 0xc3, 0x8f, 0x1e, 0x38, 0x71, 0xe3, 0xc7, 0xe, 0x1c, 0x78, 0xf1, 0xc0}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0x15, Height: 0xd, XAdvance: 0x15, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1c, 0xf1, 0xe0, 0xef, 0xdf, 0x87, 0xff, 0xfe, 0x7c, 0x78, 0xf3, 0xc3, 0x87, 0x9c, 0x1c, 0x38, 0xe1, 0xe1, 0xc7, 0xe, 0xe, 0x78, 0x70, 0xf3, 0xc3, 0x87, 0x9c, 0x3c, 0x38, 0xe1, 0xe1, 0xc7, 0xe, 0xe, 0x0}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0xe, Height: 0xd, XAdvance: 0xf, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x3d, 0xf0, 0xef, 0xe3, 0xff, 0xcf, 0xf, 0x78, 0x39, 0xc0, 0xe7, 0x3, 0x9c, 0x1e, 0xf0, 0x7b, 0xc1, 0xce, 0x7, 0x38, 0x1c, 0xe0, 0xf0}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0xe, Height: 0xd, XAdvance: 0xf, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x7, 0xe0, 0x7f, 0xe3, 0xff, 0x9f, 0x1f, 0x78, 0x3f, 0xc0, 0xff, 0x3, 0xfc, 0x1f, 0xf0, 0x7b, 0xe3, 0xe7, 0xff, 0x1f, 0xf8, 0x1f, 0x80}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0x10, Height: 0x12, XAdvance: 0xf, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0xe, 0x7c, 0xf, 0xfe, 0xf, 0xff, 0x1f, 0x1f, 0x1e, 0xf, 0x1e, 0xf, 0x1c, 0xf, 0x1c, 0xf, 0x3c, 0x1e, 0x3e, 0x3e, 0x3f, 0xfc, 0x3f, 0xf8, 0x7b, 0xe0, 0x78, 0x0, 0x70, 0x0, 0x70, 0x0, 0x70, 0x0, 0xf0, 0x0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0xe, Height: 0x12, XAdvance: 0xf, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x7, 0xbc, 0x7f, 0xf3, 0xff, 0x9f, 0x1e, 0x78, 0x3b, 0xc0, 0xef, 0x3, 0x3c, 0xc, 0xf0, 0x73, 0xe3, 0xcf, 0xff, 0x1f, 0xf8, 0x3c, 0xe0, 0x3, 0x80, 0x1e, 0x0, 0x78, 0x1, 0xc0, 0x7, 0x0}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0xa, Height: 0xd, XAdvance: 0x9, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x3d, 0xce, 0xe3, 0xf8, 0xf0, 0x78, 0x1e, 0x7, 0x1, 0xc0, 0xf0, 0x3c, 0xe, 0x3, 0x80, 0xe0, 0x0}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0xc, Height: 0xd, XAdvance: 0xd, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0x1f, 0xc3, 0xfe, 0x7f, 0xff, 0xf, 0xf0, 0xf, 0xe0, 0x7f, 0xc1, 0xfe, 0x3, 0xee, 0x1e, 0xff, 0xc7, 0xfc, 0x3f, 0x0}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0x8, Height: 0xf, XAdvance: 0x8, XOffset: 2, YOffset: -14, Bitmaps: []uint8{0x1e, 0x1e, 0x1c, 0x7f, 0xff, 0x3c, 0x38, 0x38, 0x38, 0x78, 0x78, 0x70, 0x7c, 0xf8, 0x78}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0xe, Height: 0xd, XAdvance: 0xf, XOffset: 2, YOffset: -12, Bitmaps: []uint8{0x38, 0x3c, 0xe0, 0xe3, 0x83, 0x9e, 0xe, 0x70, 0x79, 0xc1, 0xe7, 0x7, 0x3c, 0x1c, 0xf0, 0xf3, 0xe7, 0xcf, 0xff, 0x1f, 0xf8, 0x3c, 0xe0}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0xd, Height: 0xd, XAdvance: 0xd, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0xf0, 0x77, 0x87, 0xbc, 0x38, 0xe3, 0xc7, 0x1c, 0x39, 0xe1, 0xce, 0xe, 0xe0, 0x77, 0x3, 0xf0, 0xf, 0x80, 0x78, 0x3, 0xc0, 0x0}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0x11, Height: 0xd, XAdvance: 0x13, XOffset: 3, YOffset: -12, Bitmaps: []uint8{0xf1, 0xc3, 0xf8, 0xe3, 0xfc, 0xf1, 0xde, 0x79, 0xef, 0x3c, 0xe7, 0xb6, 0x73, 0xdb, 0x70, 0xed, 0xb8, 0x7c, 0xf8, 0x3e, 0x7c, 0x1f, 0x3c, 0xf, 0x1e, 0x7, 0x8e, 0x0}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0x10, Height: 0xd, XAdvance: 0xd, XOffset: 0, YOffset: -12, Bitmaps: []uint8{0xf, 0x1e, 0xf, 0x3c, 0xf, 0x38, 0x7, 0x70, 0x7, 0xf0, 0x3, 0xe0, 0x3, 0xc0, 0x7, 0xc0, 0xf, 0xe0, 0x1e, 0xe0, 0x3c, 0xf0, 0x3c, 0xf0, 0x78, 0x78}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0xf, Height: 0x12, XAdvance: 0xd, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x3c, 0x1c, 0x78, 0x78, 0xf0, 0xe1, 0xe3, 0xc1, 0xc7, 0x3, 0x9e, 0x7, 0x38, 0xe, 0xe0, 0x1d, 0xc0, 0x3f, 0x0, 0x7e, 0x0, 0x78, 0x0, 0xf0, 0x1, 0xc0, 0x7, 0x0, 0x7e, 0x0, 0xf8, 0x1, 0xe0, 0x0}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0xd, Height: 0xd, XAdvance: 0xc, XOffset: 1, YOffset: -12, Bitmaps: []uint8{0x1f, 0xf9, 0xff, 0xcf, 0xfc, 0x1, 0xe0, 0x3e, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3c, 0x3, 0xc0, 0x3f, 0xf9, 0xff, 0xcf, 0xfc, 0x0}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0x9, Height: 0x17, XAdvance: 0x9, XOffset: 3, YOffset: -17, Bitmaps: []uint8{0x7, 0x87, 0xc3, 0xe3, 0xc1, 0xc0, 0xe0, 0x70, 0x38, 0x3c, 0x1c, 0xe, 0x1e, 0xf, 0x3, 0x81, 0xc0, 0xe0, 0x70, 0x78, 0x38, 0x1c, 0xf, 0x87, 0xc1, 0xc0}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x6, Height: 0x17, XAdvance: 0x7, XOffset: 1, YOffset: -17, Bitmaps: []uint8{0xc, 0x30, 0x86, 0x18, 0x61, 0x8c, 0x30, 0xc3, 0xc, 0x61, 0x86, 0x18, 0x63, 0xc, 0x30, 0xc2, 0x0, 0x0}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0x9, Height: 0x17, XAdvance: 0x9, XOffset: 0, YOffset: -17, Bitmaps: []uint8{0x7, 0x7, 0xc3, 0xe0, 0x70, 0x38, 0x3c, 0x1c, 0xe, 0x7, 0x3, 0x81, 0xe0, 0xf0, 0xe0, 0x70, 0x78, 0x38, 0x1c, 0xe, 0x7, 0x7, 0x8f, 0x87, 0xc3, 0xc0}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0xc, Height: 0x5, XAdvance: 0xe, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0x3c, 0x7, 0xe0, 0xc7, 0x30, 0x7e, 0x1, 0xc0}},
	},

	YAdvance: 0x1d,
}