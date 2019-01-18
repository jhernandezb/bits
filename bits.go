package bits

// GetBit extracts a specific bit in a byte.
func GetBit(b byte, n uint8) uint8 {
	return b & (1 << n) >> n
}

// CreateByteMask creates mask of consecutive bits from start to end index.
func CreateByteMask(start, end uint8) (r byte) {
	for i := start; i <= end; i++ {
		r = r | (1 << i)
	}
	return
}
