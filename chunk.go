package gocraft

const (
	ChunkWidth  = 16
	ChunkDepth  = 16
	ChunkHeight = 16
)

type Chunk struct {
	Blocks [ChunkWidth * ChunkHeight * ChunkDepth]Block
}

func (c *Chunk) GetBlock(x, y, z int) Block {
	// FIXME not sure about this
	return c.Blocks[x+z*ChunkDepth+y*ChunkHeight]
}
