package parquet

import (
	"reflect"
	"testing"
)

func rleDecodeAll(w int, data []byte) (a []int32, err error) {
	d := newRLEDecoder(data, w)
	for d.hasNext() {
		a = append(a, d.nextInt32())
	}
	return a, d.err()
}

func repeatInt32(count int, value int32) (a []int32) {
	for i := 0; i < count; i++ {
		a = append(a, value)
	}
	return
}

var rle32Tests = []struct {
	width  int
	data   []byte
	values []int32
}{
	// Single RLE run: 1-bit per value, 10 x 0
	{1, []byte{0x14, 0x00}, repeatInt32(10, 0)},

	// Single RLE run: 20-bits per value, 300x1
	{20, []byte{0xD8, 0x04, 0x01, 0x00, 0x00}, repeatInt32(300, 1)},

	// 2 RLE runs: 1-bit per value, 10x0, 9x1
	{1, []byte{0x14, 0x00, 0x12, 0x01}, append(repeatInt32(10, 0), repeatInt32(9, 1)...)},
}

func TestRLEDecoder(t *testing.T) {
	for i, test := range rle32Tests {
		values, err := rleDecodeAll(test.width, test.data)
		if err != nil {
			t.Errorf("test %d. unexpected error: %s", i, err)
		}
		if !reflect.DeepEqual(values, test.values) {
			t.Errorf("test %d. got %v, want %v", i, values, test.values)
		}
	}
}