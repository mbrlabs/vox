package gocraft

type Block uint8

const (
	blockActiveMask = 0x80 // 0b10000000
)

func (b Block) Active() bool {
	return (blockActiveMask & b) == blockActiveMask
}

func (b Block) Activate(active bool) Block {
	if active {
		return b | blockActiveMask
	}
	return 0
}
