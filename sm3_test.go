package sm3_test

import "testing"

import "github.com/sammyne/sm3"

func TestSum(t *testing.T) {
	testVector := []struct {
		data   []byte
		expect [sm3.Size]byte
	}{
		{
			[]byte("abc"),
			[sm3.Size]byte{
				0x66, 0xc7, 0xf0, 0xf4, 0x62, 0xee, 0xed, 0xd9,
				0xd1, 0xf2, 0xd4, 0x6b, 0xdc, 0x10, 0xe4, 0xe2,
				0x41, 0x67, 0xc4, 0x87, 0x5c, 0xf2, 0xf7, 0xa2,
				0x29, 0x7d, 0xa0, 0x2b, 0x8f, 0x4b, 0xa8, 0xe0,
			},
		},
		{
			[]byte("abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd"),
			[sm3.Size]byte{
				0xde, 0xbe, 0x9f, 0xf9, 0x22, 0x75, 0xb8, 0xa1,
				0x38, 0x60, 0x48, 0x89, 0xc1, 0x8e, 0x5a, 0x4d,
				0x6f, 0xdb, 0x70, 0xe5, 0x38, 0x7e, 0x57, 0x65,
				0x29, 0x3d, 0xcb, 0xa3, 0x9c, 0x0c, 0x57, 0x32,
			},
		},
	}

	for i, c := range testVector {
		if got := sm3.Sum(c.data); got != c.expect {
			t.Fatalf("#%d failed: expect %v, got %v", i, c.expect, got)
		}
	}
}
