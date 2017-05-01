package gocraft

type Block uint8

const (
	blockActiveMask = 0x80 // 0b10000000
	blockTypeMask   = 0x7F // 0b01111111
)

func (b Block) Active() bool {
	return (blockActiveMask & b) == blockActiveMask
}

func (b Block) Activate(active bool) Block {
	if active {
		return b | blockActiveMask
	}
	return b & blockTypeMask
}

func (b Block) BlockType() uint8 {
	return uint8(b & blockTypeMask)
}

type BlockType struct {
	Color *Color
}
