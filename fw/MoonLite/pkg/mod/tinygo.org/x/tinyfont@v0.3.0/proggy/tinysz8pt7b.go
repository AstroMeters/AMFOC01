package proggy

import (
	"tinygo.org/x/tinyfont"
)

var TinySZ8pt7b = tinyfont.Font{
	BBox: [4]int8{6, 10, 0, -7},
	Glyphs: []tinyfont.Glyph{
		/*   */ tinyfont.Glyph{Rune: 32, Width: 0x1, Height: 0x1, XAdvance: 0x6, XOffset: 0, YOffset: 0, Bitmaps: []uint8{0x0}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x1, Height: 0x7, XAdvance: 0x6, XOffset: 2, YOffset: -6, Bitmaps: []uint8{0xfa}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0x3, Height: 0x3, XAdvance: 0x6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xb6, 0x80}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0x5, Height: 0x6, XAdvance: 0x6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x57, 0xd4, 0xaf, 0xa8}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x23, 0xe8, 0xe2, 0xf8, 0x80}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0x6, Height: 0x6, XAdvance: 0x6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x46, 0xa5, 0xa, 0x56, 0x20}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0x6, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x62, 0x49, 0x18, 0x96, 0x27, 0x40}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x1, Height: 0x3, XAdvance: 0x6, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0xe0}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x3, Height: 0x9, XAdvance: 0x6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x2a, 0x49, 0x22, 0x20}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x3, Height: 0x9, XAdvance: 0x6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x88, 0x92, 0x4a, 0x80}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x25, 0x5d, 0x52, 0x0}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x21, 0x3e, 0x42, 0x0}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x2, Height: 0x2, XAdvance: 0x6, XOffset: 1, YOffset: 0, Bitmaps: []uint8{0x60}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x5, Height: 0x1, XAdvance: 0x6, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0xf8}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x1, Height: 0x1, XAdvance: 0x6, XOffset: 2, YOffset: 0, Bitmaps: []uint8{0x80}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0x5, Height: 0x9, XAdvance: 0x6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x8, 0x44, 0x22, 0x11, 0x8, 0x80}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x67, 0x5c, 0xc5, 0xc0}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0x3, Height: 0x7, XAdvance: 0x6, XOffset: 1, YOffset: -6, Bitmaps: []uint8{0x59, 0x24, 0xb8}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x42, 0x64, 0x43, 0xe0}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x42, 0x60, 0xc5, 0xc0}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x11, 0x95, 0x2f, 0x88, 0x40}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xfc, 0x21, 0xe0, 0xc5, 0xc0}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x32, 0x21, 0xe8, 0xc5, 0xc0}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xf8, 0x44, 0x22, 0x11, 0x0}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x62, 0xe8, 0xc5, 0xc0}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x62, 0xf0, 0x89, 0x80}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x1, Height: 0x4, XAdvance: 0x6, XOffset: 2, YOffset: -4, Bitmaps: []uint8{0x90}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x2, Height: 0x5, XAdvance: 0x6, XOffset: 1, YOffset: -4, Bitmaps: []uint8{0x41, 0x80}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x9, 0xb0, 0x60, 0x80}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0x5, Height: 0x3, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xf8, 0x3e}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -5, Bitmaps: []uint8{0x83, 0x6, 0xc8, 0x0}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x42, 0x22, 0x0, 0x80}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0x6, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x7a, 0x1b, 0x6d, 0xbe, 0x7, 0xc0}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x21, 0x14, 0xa7, 0x46, 0x20}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xf4, 0x63, 0xe8, 0xc7, 0xc0}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x61, 0x8, 0x45, 0xc0}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xe4, 0xa3, 0x18, 0xcb, 0x80}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xfc, 0x21, 0xe8, 0x43, 0xe0}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xfc, 0x21, 0xe8, 0x42, 0x0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x61, 0x38, 0xc5, 0xe0}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x8c, 0x63, 0xf8, 0xc6, 0x20}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x3, Height: 0x7, XAdvance: 0x6, XOffset: 1, YOffset: -6, Bitmaps: []uint8{0xe9, 0x24, 0xb8}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x38, 0x42, 0x10, 0x87, 0xc0}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x8c, 0xa9, 0x8a, 0x4a, 0x20}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x84, 0x21, 0x8, 0x43, 0xe0}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x8c, 0x77, 0xba, 0xd6, 0x20}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xce, 0x6b, 0x5a, 0xce, 0x60}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x63, 0x18, 0xc5, 0xc0}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xf4, 0x63, 0xe8, 0x42, 0x0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x63, 0x18, 0xcd, 0xe0}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xf4, 0x63, 0xe8, 0xc6, 0x20}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x74, 0x60, 0xe0, 0xc5, 0xc0}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xf9, 0x8, 0x42, 0x10, 0x80}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x8c, 0x63, 0x18, 0xc5, 0xc0}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x8c, 0x62, 0xa5, 0x10, 0x80}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x8d, 0x6b, 0x55, 0x29, 0x40}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x8c, 0x54, 0x45, 0x46, 0x20}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x8c, 0x62, 0xe2, 0x10, 0x80}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0xf8, 0x44, 0x44, 0x43, 0xe0}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0x3, Height: 0x9, XAdvance: 0x6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xf2, 0x49, 0x24, 0xe0}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x5, Height: 0x9, XAdvance: 0x6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x84, 0x10, 0x82, 0x10, 0x42, 0x8}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0x3, Height: 0x9, XAdvance: 0x6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xe4, 0x92, 0x49, 0xe0}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0x5, Height: 0x3, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x22, 0xa2}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0x6, Height: 0x1, XAdvance: 0x6, XOffset: 0, YOffset: 1, Bitmaps: []uint8{0xfc}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x2, Height: 0x2, XAdvance: 0x6, XOffset: 2, YOffset: -6, Bitmaps: []uint8{0x90}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x70, 0x5f, 0x17, 0x80}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0x5, Height: 0x8, XAdvance: 0x6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x84, 0x21, 0xe8, 0xc6, 0x3e}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x7c, 0x21, 0x7, 0x80}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0x5, Height: 0x8, XAdvance: 0x6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x8, 0x42, 0xf8, 0xc6, 0x2f}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x74, 0x7f, 0x7, 0x80}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0x5, Height: 0x8, XAdvance: 0x6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x3a, 0x11, 0xe4, 0x21, 0x8}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x74, 0x63, 0x17, 0x85, 0xc0}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0x5, Height: 0x8, XAdvance: 0x6, XOffset: 0, YOffset: -7, Bitmaps: []uint8{0x84, 0x21, 0xe8, 0xc6, 0x31}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x2, Height: 0x7, XAdvance: 0x6, XOffset: 1, YOffset: -6, Bitmaps: []uint8{0x4d, 0x54}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0x4, Height: 0x9, XAdvance: 0x6, XOffset: 0, YOffset: -6, Bitmaps: []uint8{0x10, 0x31, 0x11, 0x11, 0xe0}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0x4, Height: 0x8, XAdvance: 0x6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x88, 0x89, 0xac, 0xa9}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x2, Height: 0x8, XAdvance: 0x6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xd5, 0x55}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xd5, 0x6b, 0x5a, 0x80}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xf4, 0x63, 0x18, 0x80}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x74, 0x63, 0x17, 0x0}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xf4, 0x63, 0x1f, 0x42, 0x0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x7c, 0x63, 0x17, 0x84, 0x20}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xb6, 0x61, 0x8, 0x0}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x7c, 0x1c, 0x1f, 0x0}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0x4, Height: 0x7, XAdvance: 0x6, XOffset: 1, YOffset: -6, Bitmaps: []uint8{0x88, 0xe8, 0x88, 0x70}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x8c, 0x63, 0x17, 0x80}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x8c, 0x54, 0xa2, 0x0}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x8d, 0x6a, 0xa5, 0x0}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x8a, 0x88, 0xa8, 0x80}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0x5, Height: 0x7, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0x8c, 0x63, 0x17, 0x85, 0xc0}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0x5, Height: 0x5, XAdvance: 0x6, XOffset: 0, YOffset: -4, Bitmaps: []uint8{0xf8, 0x88, 0x8f, 0x80}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0x3, Height: 0x9, XAdvance: 0x6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x29, 0x28, 0x92, 0x20}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x1, Height: 0x9, XAdvance: 0x6, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0xff, 0x80}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0x3, Height: 0x9, XAdvance: 0x6, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x89, 0x22, 0x92, 0x80}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0x6, Height: 0x2, XAdvance: 0x6, XOffset: 0, YOffset: -3, Bitmaps: []uint8{0x66, 0x60}},
	},

	YAdvance: 0xa,
}
