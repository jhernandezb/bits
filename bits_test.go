package bits_test

import (
	"testing"

	"github.com/jhernandezb/bits"
)

func TestGetBit(t *testing.T) {
	const testByte = 0x05 // 101
	Equals(t, bits.GetBit(testByte, 0), uint8(1))
	Equals(t, bits.GetBit(testByte, 1), uint8(0))
	Equals(t, bits.GetBit(testByte, 2), uint8(1))
}

func TestCreateMask(t *testing.T) {
	Equals(t, bits.CreateByteMask(0, 0), byte(1))
	// 1100
	Equals(t, bits.CreateByteMask(2, 3), byte(0xC))

	// Value: 1111 1111 Mask: 0000 1100  => 0xC
	Equals(t, uint8(bits.CreateByteMask(2, 3)&0xFF), byte(0xc))
}
