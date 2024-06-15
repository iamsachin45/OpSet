package main

type BitSet struct {
	bits []uint64
}

// NewBitSet initializes a new BitSet
func NewBitSet(size uint32) *BitSet {
	numUint64s := (size + 63) / 64
	//fmt.Println(numUint64s, " ", (size + 63))
	return &BitSet{
		bits: make([]uint64, numUint64s),
	}
}

// Add inserts a number into the BitSet
func (bs *BitSet) Add(num uint32) {
	index := num / 64
	bitPosition := num % 64
	if index >= uint32(len(bs.bits)) {
		newBits := make([]uint64, index+1)
		copy(newBits, bs.bits)
		bs.bits = newBits
	}
	bs.bits[index] |= 1 << bitPosition
}

// Contains checks if a number is present in the BitSet
func (bs *BitSet) Contains(num uint32) bool {
	index := num / 64
	bitPosition := num % 64
	if index >= uint32(len(bs.bits)) {
		return false
	}
	return bs.bits[index]&(1<<bitPosition) != 0
}

// Remove deletes a number from the BitSet
func (bs *BitSet) Remove(num uint32) {
	index := num / 64
	bitPosition := num % 64
	if index < uint32(len(bs.bits)) {
		bs.bits[index] &^= 1 << bitPosition
	}
}

func main() {
	size := uint32(65) // Total numbers we want to support
	bitSet := NewBitSet(size)

}
