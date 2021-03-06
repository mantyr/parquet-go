package parquet

import "testing"

func TestUnpack8int32(t *testing.T) {
	for _, test := range unpack8int32Tests {
		unpacker := unpack8Int32FuncByWidth[test.width]
		if got := unpacker(test.data); got != test.values {
			t.Errorf("unpack for width %d: got %v, want %v", test.width, got, test.values)
		}
	}
}

var unpack8int32Tests = []struct {
	width  int
	data   []byte
	values [8]int32
}{
	// bit width = 0
	{0, []byte{}, [8]int32{0, 0, 0, 0, 0, 0, 0, 0}},

	// bit width = 1
	{1, []byte{0x00}, [8]int32{0, 0, 0, 0, 0, 0, 0, 0}},
	{1, []byte{0xFF}, [8]int32{1, 1, 1, 1, 1, 1, 1, 1}},
	{1, []byte{0x4D}, [8]int32{1, 0, 1, 1, 0, 0, 1, 0}},

	// bit width = 2
	{2, []byte{0x55, 0x55}, [8]int32{1, 1, 1, 1, 1, 1, 1, 1}},
	{2, []byte{0xAA, 0xAA}, [8]int32{2, 2, 2, 2, 2, 2, 2, 2}},
	{2, []byte{0xA4, 0x41}, [8]int32{0, 1, 2, 2, 1, 0, 0, 1}},

	// bit width = 3
	{3, []byte{0x00, 0x00, 0x00}, [8]int32{0, 0, 0, 0, 0, 0, 0, 0}},
	{3, []byte{0x88, 0xC6, 0xFA}, [8]int32{0, 1, 2, 3, 4, 5, 6, 7}},
	{3, []byte{0x77, 0x39, 0x05}, [8]int32{7, 6, 5, 4, 3, 2, 1, 0}},
	{3, []byte{0x23, 0xA2, 0x11}, [8]int32{3, 4, 0, 1, 2, 3, 4, 0}},

	// bit width = 4
	{4, []byte{0x00, 0x00, 0x00, 0x00}, [8]int32{0, 0, 0, 0, 0, 0, 0, 0}},
	{4, []byte{0x10, 0x32, 0x54, 0x76}, [8]int32{0, 1, 2, 3, 4, 5, 6, 7}},
	{4, []byte{0x67, 0x45, 0x23, 0x01}, [8]int32{7, 6, 5, 4, 3, 2, 1, 0}},
	{4, []byte{0xEF, 0xCD, 0xAB, 0x89}, [8]int32{15, 14, 13, 12, 11, 10, 9, 8}},

	// bit width = 5
	{5, []byte{0x00, 0x00, 0x00, 0x00, 0x00}, [8]int32{0, 0, 0, 0, 0, 0, 0, 0}},
	{5, []byte{0x20, 0x88, 0x41, 0x8A, 0x39}, [8]int32{0, 1, 2, 3, 4, 5, 6, 7}},
	{5, []byte{0xC7, 0x14, 0x32, 0x44, 0x00}, [8]int32{7, 6, 5, 4, 3, 2, 1, 0}},
	{5, []byte{0xDF, 0x77, 0xBE, 0x75, 0xC6}, [8]int32{31, 30, 29, 28, 27, 26, 25, 24}},

	// bit width = 6
	{
		6,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		6,
		[]byte{0x40, 0x20, 0x0C, 0x44, 0x61, 0x1C},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		6,
		[]byte{0x87, 0x51, 0x10, 0x83, 0x10, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		6,
		[]byte{0xBF, 0xDF, 0xF3, 0xBB, 0x9E, 0xE3},
		[8]int32{63, 62, 61, 60, 59, 58, 57, 56},
	},

	// bit width = 7
	{
		7,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		7,
		[]byte{0x80, 0x80, 0x60, 0x40, 0x28, 0x18, 0x0E},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		7,
		[]byte{0x07, 0x43, 0x81, 0x30, 0x10, 0x04, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		7,
		[]byte{0x7F, 0x7F, 0x9F, 0xBF, 0xD7, 0xE7, 0xF1},
		[8]int32{127, 126, 125, 124, 123, 122, 121, 120},
	},

	// bit width = 8
	{
		8,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		8,
		[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		8,
		[]byte{0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		8,
		[]byte{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
		[8]int32{255, 254, 253, 252, 251, 250, 249, 248},
	},

	// bit width = 9
	{
		9,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		9,
		[]byte{0x00, 0x02, 0x08, 0x18, 0x40, 0xA0, 0x80, 0x81, 0x03},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		9,
		[]byte{0x07, 0x0C, 0x14, 0x20, 0x30, 0x40, 0x40, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		9,
		[]byte{0xFF, 0xFD, 0xF7, 0xE7, 0xBF, 0x5F, 0x7F, 0x7E, 0xFC},
		[8]int32{511, 510, 509, 508, 507, 506, 505, 504},
	},

	// bit width = 10
	{
		10,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		10,
		[]byte{0x00, 0x04, 0x20, 0xC0, 0x00, 0x04, 0x14, 0x60, 0xC0, 0x01},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		10,
		[]byte{0x07, 0x18, 0x50, 0x00, 0x01, 0x03, 0x08, 0x10, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		10,
		[]byte{0xFF, 0xFB, 0xDF, 0x3F, 0xFF, 0xFB, 0xEB, 0x9F, 0x3F, 0xFE},
		[8]int32{1023, 1022, 1021, 1020, 1019, 1018, 1017, 1016},
	},

	// bit width = 11
	{
		11,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		11,
		[]byte{0x00, 0x08, 0x80, 0x00, 0x06, 0x40, 0x80, 0x02, 0x18, 0xE0, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		11,
		[]byte{0x07, 0x30, 0x40, 0x01, 0x08, 0x30, 0x00, 0x01, 0x04, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		11,
		[]byte{0xFF, 0xF7, 0x7F, 0xFF, 0xF9, 0xBF, 0x7F, 0xFD, 0xE7, 0x1F, 0xFF},
		[8]int32{2047, 2046, 2045, 2044, 2043, 2042, 2041, 2040},
	},

	// bit width = 12
	{
		12,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		12,
		[]byte{0x00, 0x10, 0x00, 0x02, 0x30, 0x00, 0x04, 0x50, 0x00, 0x06, 0x70, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		12,
		[]byte{0x07, 0x60, 0x00, 0x05, 0x40, 0x00, 0x03, 0x20, 0x00, 0x01, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		12,
		[]byte{0xFF, 0xEF, 0xFF, 0xFD, 0xCF, 0xFF, 0xFB, 0xAF, 0xFF, 0xF9, 0x8F, 0xFF},
		[8]int32{4095, 4094, 4093, 4092, 4091, 4090, 4089, 4088},
	},

	// bit width = 13
	{
		13,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		13,
		[]byte{0x00, 0x20, 0x00, 0x08, 0x80, 0x01, 0x40, 0x00, 0x0A, 0x80, 0x01, 0x38, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		13,
		[]byte{0x07, 0xC0, 0x00, 0x14, 0x00, 0x02, 0x30, 0x00, 0x04, 0x40, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		13,
		[]byte{0xFF, 0xDF, 0xFF, 0xF7, 0x7F, 0xFE, 0xBF, 0xFF, 0xF5, 0x7F, 0xFE, 0xC7, 0xFF},
		[8]int32{8191, 8190, 8189, 8188, 8187, 8186, 8185, 8184},
	},

	// bit width = 14
	{
		14,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		14,
		[]byte{0x00, 0x40, 0x00, 0x20, 0x00, 0x0C, 0x00, 0x04, 0x40, 0x01, 0x60, 0x00, 0x1C, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		14,
		[]byte{0x07, 0x80, 0x01, 0x50, 0x00, 0x10, 0x00, 0x03, 0x80, 0x00, 0x10, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		14,
		[]byte{0xFF, 0xBF, 0xFF, 0xDF, 0xFF, 0xF3, 0xFF, 0xFB, 0xBF, 0xFE, 0x9F, 0xFF, 0xE3, 0xFF},
		[8]int32{16383, 16382, 16381, 16380, 16379, 16378, 16377, 16376},
	},

	// bit width = 15
	{
		15,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		15,
		[]byte{0x00, 0x80, 0x00, 0x80, 0x00, 0x60, 0x00, 0x40, 0x00, 0x28, 0x00, 0x18, 0x00, 0x0E, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		15,
		[]byte{0x07, 0x00, 0x03, 0x40, 0x01, 0x80, 0x00, 0x30, 0x00, 0x10, 0x00, 0x04, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		15,
		[]byte{0xFF, 0x7F, 0xFF, 0x7F, 0xFF, 0x9F, 0xFF, 0xBF, 0xFF, 0xD7, 0xFF, 0xE7, 0xFF, 0xF1, 0xFF},
		[8]int32{32767, 32766, 32765, 32764, 32763, 32762, 32761, 32760},
	},

	// bit width = 16
	{
		16,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		16,
		[]byte{0x00, 0x00, 0x01, 0x00, 0x02, 0x00, 0x03, 0x00, 0x04, 0x00, 0x05, 0x00, 0x06, 0x00, 0x07, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		16,
		[]byte{0x07, 0x00, 0x06, 0x00, 0x05, 0x00, 0x04, 0x00, 0x03, 0x00, 0x02, 0x00, 0x01, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		16,
		[]byte{0xFF, 0xFF, 0xFE, 0xFF, 0xFD, 0xFF, 0xFC, 0xFF, 0xFB, 0xFF, 0xFA, 0xFF, 0xF9, 0xFF, 0xF8, 0xFF},
		[8]int32{65535, 65534, 65533, 65532, 65531, 65530, 65529, 65528},
	},

	// bit width = 17
	{
		17,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		17,
		[]byte{0x00, 0x00, 0x02, 0x00, 0x08, 0x00, 0x18, 0x00, 0x40, 0x00, 0xA0, 0x00, 0x80, 0x01, 0x80, 0x03, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		17,
		[]byte{0x07, 0x00, 0x0C, 0x00, 0x14, 0x00, 0x20, 0x00, 0x30, 0x00, 0x40, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		17,
		[]byte{0xFF, 0xFF, 0xFD, 0xFF, 0xF7, 0xFF, 0xE7, 0xFF, 0xBF, 0xFF, 0x5F, 0xFF, 0x7F, 0xFE, 0x7F, 0xFC, 0xFF},
		[8]int32{131071, 131070, 131069, 131068, 131067, 131066, 131065, 131064},
	},

	// bit width = 18
	{
		18,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		18,
		[]byte{0x00, 0x00, 0x04, 0x00, 0x20, 0x00, 0xC0, 0x00, 0x00, 0x04, 0x00, 0x14, 0x00, 0x60, 0x00, 0xC0, 0x01, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		18,
		[]byte{0x07, 0x00, 0x18, 0x00, 0x50, 0x00, 0x00, 0x01, 0x00, 0x03, 0x00, 0x08, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		18,
		[]byte{0xFF, 0xFF, 0xFB, 0xFF, 0xDF, 0xFF, 0x3F, 0xFF, 0xFF, 0xFB, 0xFF, 0xEB, 0xFF, 0x9F, 0xFF, 0x3F, 0xFE, 0xFF},
		[8]int32{262143, 262142, 262141, 262140, 262139, 262138, 262137, 262136},
	},

	// bit width = 19
	{
		19,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		19,
		[]byte{0x00, 0x00, 0x08, 0x00, 0x80, 0x00, 0x00, 0x06, 0x00, 0x40, 0x00, 0x80, 0x02, 0x00, 0x18, 0x00, 0xE0, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		19,
		[]byte{0x07, 0x00, 0x30, 0x00, 0x40, 0x01, 0x00, 0x08, 0x00, 0x30, 0x00, 0x00, 0x01, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		19,
		[]byte{0xFF, 0xFF, 0xF7, 0xFF, 0x7F, 0xFF, 0xFF, 0xF9, 0xFF, 0xBF, 0xFF, 0x7F, 0xFD, 0xFF, 0xE7, 0xFF, 0x1F, 0xFF, 0xFF},
		[8]int32{524287, 524286, 524285, 524284, 524283, 524282, 524281, 524280},
	},

	// bit width = 20
	{
		20,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		20,
		[]byte{0x00, 0x00, 0x10, 0x00, 0x00, 0x02, 0x00, 0x30, 0x00, 0x00, 0x04, 0x00, 0x50, 0x00, 0x00, 0x06, 0x00, 0x70, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		20,
		[]byte{0x07, 0x00, 0x60, 0x00, 0x00, 0x05, 0x00, 0x40, 0x00, 0x00, 0x03, 0x00, 0x20, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		20,
		[]byte{0xFF, 0xFF, 0xEF, 0xFF, 0xFF, 0xFD, 0xFF, 0xCF, 0xFF, 0xFF, 0xFB, 0xFF, 0xAF, 0xFF, 0xFF, 0xF9, 0xFF, 0x8F, 0xFF, 0xFF},
		[8]int32{1048575, 1048574, 1048573, 1048572, 1048571, 1048570, 1048569, 1048568},
	},

	// bit width = 21
	{
		21,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		21,
		[]byte{0x00, 0x00, 0x20, 0x00, 0x00, 0x08, 0x00, 0x80, 0x01, 0x00, 0x40, 0x00, 0x00, 0x0A, 0x00, 0x80, 0x01, 0x00, 0x38, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		21,
		[]byte{0x07, 0x00, 0xC0, 0x00, 0x00, 0x14, 0x00, 0x00, 0x02, 0x00, 0x30, 0x00, 0x00, 0x04, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		21,
		[]byte{0xFF, 0xFF, 0xDF, 0xFF, 0xFF, 0xF7, 0xFF, 0x7F, 0xFE, 0xFF, 0xBF, 0xFF, 0xFF, 0xF5, 0xFF, 0x7F, 0xFE, 0xFF, 0xC7, 0xFF, 0xFF},
		[8]int32{2097151, 2097150, 2097149, 2097148, 2097147, 2097146, 2097145, 2097144},
	},

	// bit width = 22
	{
		22,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		22,
		[]byte{0x00, 0x00, 0x40, 0x00, 0x00, 0x20, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x04, 0x00, 0x40, 0x01, 0x00, 0x60, 0x00, 0x00, 0x1C, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		22,
		[]byte{0x07, 0x00, 0x80, 0x01, 0x00, 0x50, 0x00, 0x00, 0x10, 0x00, 0x00, 0x03, 0x00, 0x80, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		22,
		[]byte{0xFF, 0xFF, 0xBF, 0xFF, 0xFF, 0xDF, 0xFF, 0xFF, 0xF3, 0xFF, 0xFF, 0xFB, 0xFF, 0xBF, 0xFE, 0xFF, 0x9F, 0xFF, 0xFF, 0xE3, 0xFF, 0xFF},
		[8]int32{4194303, 4194302, 4194301, 4194300, 4194299, 4194298, 4194297, 4194296},
	},

	// bit width = 23
	{
		23,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		23,
		[]byte{0x00, 0x00, 0x80, 0x00, 0x00, 0x80, 0x00, 0x00, 0x60, 0x00, 0x00, 0x40, 0x00, 0x00, 0x28, 0x00, 0x00, 0x18, 0x00, 0x00, 0x0E, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		23,
		[]byte{0x07, 0x00, 0x00, 0x03, 0x00, 0x40, 0x01, 0x00, 0x80, 0x00, 0x00, 0x30, 0x00, 0x00, 0x10, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		23,
		[]byte{0xFF, 0xFF, 0x7F, 0xFF, 0xFF, 0x7F, 0xFF, 0xFF, 0x9F, 0xFF, 0xFF, 0xBF, 0xFF, 0xFF, 0xD7, 0xFF, 0xFF, 0xE7, 0xFF, 0xFF, 0xF1, 0xFF, 0xFF},
		[8]int32{8388607, 8388606, 8388605, 8388604, 8388603, 8388602, 8388601, 8388600},
	},

	// bit width = 24
	{
		24,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		24,
		[]byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x02, 0x00, 0x00, 0x03, 0x00, 0x00, 0x04, 0x00, 0x00, 0x05, 0x00, 0x00, 0x06, 0x00, 0x00, 0x07, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		24,
		[]byte{0x07, 0x00, 0x00, 0x06, 0x00, 0x00, 0x05, 0x00, 0x00, 0x04, 0x00, 0x00, 0x03, 0x00, 0x00, 0x02, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		24,
		[]byte{0xFF, 0xFF, 0xFF, 0xFE, 0xFF, 0xFF, 0xFD, 0xFF, 0xFF, 0xFC, 0xFF, 0xFF, 0xFB, 0xFF, 0xFF, 0xFA, 0xFF, 0xFF, 0xF9, 0xFF, 0xFF, 0xF8, 0xFF, 0xFF},
		[8]int32{16777215, 16777214, 16777213, 16777212, 16777211, 16777210, 16777209, 16777208},
	},

	// bit width = 25
	{
		25,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		25,
		[]byte{0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x08, 0x00, 0x00, 0x18, 0x00, 0x00, 0x40, 0x00, 0x00, 0xA0, 0x00, 0x00, 0x80, 0x01, 0x00, 0x80, 0x03, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		25,
		[]byte{0x07, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x14, 0x00, 0x00, 0x20, 0x00, 0x00, 0x30, 0x00, 0x00, 0x40, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		25,
		[]byte{0xFF, 0xFF, 0xFF, 0xFD, 0xFF, 0xFF, 0xF7, 0xFF, 0xFF, 0xE7, 0xFF, 0xFF, 0xBF, 0xFF, 0xFF, 0x5F, 0xFF, 0xFF, 0x7F, 0xFE, 0xFF, 0x7F, 0xFC, 0xFF, 0xFF},
		[8]int32{33554431, 33554430, 33554429, 33554428, 33554427, 33554426, 33554425, 33554424},
	},

	// bit width = 26
	{
		26,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		26,
		[]byte{0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x20, 0x00, 0x00, 0xC0, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x14, 0x00, 0x00, 0x60, 0x00, 0x00, 0xC0, 0x01, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		26,
		[]byte{0x07, 0x00, 0x00, 0x18, 0x00, 0x00, 0x50, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x03, 0x00, 0x00, 0x08, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		26,
		[]byte{0xFF, 0xFF, 0xFF, 0xFB, 0xFF, 0xFF, 0xDF, 0xFF, 0xFF, 0x3F, 0xFF, 0xFF, 0xFF, 0xFB, 0xFF, 0xFF, 0xEB, 0xFF, 0xFF, 0x9F, 0xFF, 0xFF, 0x3F, 0xFE, 0xFF, 0xFF},
		[8]int32{67108863, 67108862, 67108861, 67108860, 67108859, 67108858, 67108857, 67108856},
	},

	// bit width = 27
	{
		27,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		27,
		[]byte{0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x40, 0x00, 0x00, 0x80, 0x02, 0x00, 0x00, 0x18, 0x00, 0x00, 0xE0, 0x00, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		27,
		[]byte{0x07, 0x00, 0x00, 0x30, 0x00, 0x00, 0x40, 0x01, 0x00, 0x00, 0x08, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		27,
		[]byte{0xFF, 0xFF, 0xFF, 0xF7, 0xFF, 0xFF, 0x7F, 0xFF, 0xFF, 0xFF, 0xF9, 0xFF, 0xFF, 0xBF, 0xFF, 0xFF, 0x7F, 0xFD, 0xFF, 0xFF, 0xE7, 0xFF, 0xFF, 0x1F, 0xFF, 0xFF, 0xFF},
		[8]int32{134217727, 134217726, 134217725, 134217724, 134217723, 134217722, 134217721, 134217720},
	},

	// bit width = 28
	{
		28,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		28,
		[]byte{0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x50, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x70, 0x00, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		28,
		[]byte{0x07, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		28,
		[]byte{0xFF, 0xFF, 0xFF, 0xEF, 0xFF, 0xFF, 0xFF, 0xFD, 0xFF, 0xFF, 0xCF, 0xFF, 0xFF, 0xFF, 0xFB, 0xFF, 0xFF, 0xAF, 0xFF, 0xFF, 0xFF, 0xF9, 0xFF, 0xFF, 0x8F, 0xFF, 0xFF, 0xFF},
		[8]int32{268435455, 268435454, 268435453, 268435452, 268435451, 268435450, 268435449, 268435448},
	},

	// bit width = 29
	{
		29,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		29,
		[]byte{0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x0A, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		29,
		[]byte{0x07, 0x00, 0x00, 0xC0, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		29,
		[]byte{0xFF, 0xFF, 0xFF, 0xDF, 0xFF, 0xFF, 0xFF, 0xF7, 0xFF, 0xFF, 0x7F, 0xFE, 0xFF, 0xFF, 0xBF, 0xFF, 0xFF, 0xFF, 0xF5, 0xFF, 0xFF, 0x7F, 0xFE, 0xFF, 0xFF, 0xC7, 0xFF, 0xFF, 0xFF},
		[8]int32{536870911, 536870910, 536870909, 536870908, 536870907, 536870906, 536870905, 536870904},
	},

	// bit width = 30
	{
		30,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		30,
		[]byte{0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x0C, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x40, 0x01, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x1C, 0x00, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		30,
		[]byte{0x07, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x50, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		30,
		[]byte{0xFF, 0xFF, 0xFF, 0xBF, 0xFF, 0xFF, 0xFF, 0xDF, 0xFF, 0xFF, 0xFF, 0xF3, 0xFF, 0xFF, 0xFF, 0xFB, 0xFF, 0xFF, 0xBF, 0xFE, 0xFF, 0xFF, 0x9F, 0xFF, 0xFF, 0xFF, 0xE3, 0xFF, 0xFF, 0xFF},
		[8]int32{1073741823, 1073741822, 1073741821, 1073741820, 1073741819, 1073741818, 1073741817, 1073741816},
	},

	// bit width = 31
	{
		31,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		31,
		[]byte{0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x18, 0x00, 0x00, 0x00, 0x0E, 0x00, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		31,
		[]byte{0x07, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x40, 0x01, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		31,
		[]byte{0xFF, 0xFF, 0xFF, 0x7F, 0xFF, 0xFF, 0xFF, 0x7F, 0xFF, 0xFF, 0xFF, 0x9F, 0xFF, 0xFF, 0xFF, 0xBF, 0xFF, 0xFF, 0xFF, 0xD7, 0xFF, 0xFF, 0xFF, 0xE7, 0xFF, 0xFF, 0xFF, 0xF1, 0xFF, 0xFF, 0xFF},
		[8]int32{2147483647, 2147483646, 2147483645, 2147483644, 2147483643, 2147483642, 2147483641, 2147483640},
	},

	// bit width = 32
	{
		32,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		32,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00},
		[8]int32{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		32,
		[]byte{0x07, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		[8]int32{7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		32,
		[]byte{0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0xFF, 0xFF, 0xFF, 0xFD, 0xFF, 0xFF, 0xFF, 0xFC, 0xFF, 0xFF, 0xFF, 0xFB, 0xFF, 0xFF, 0xFF, 0xFA, 0xFF, 0xFF, 0xFF, 0xF9, 0xFF, 0xFF, 0xFF},
		[8]int32{0, -1, -2, -3, -4, -5, -6, -7},
	},
}
