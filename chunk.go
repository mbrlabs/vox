package gocraft

const (
	ChunkWidth  = 16
	ChunkDepth  = 16
	ChunkHeight = 16
	ChunkXZSize = ChunkWidth * ChunkDepth
)

type Chunk struct {
	Blocks [ChunkWidth * ChunkDepth * ChunkHeight]Block
}

func (c *Chunk) Get(x, y, z int) Block {
	return c.Blocks[x+z*ChunkDepth+y*ChunkXZSize]
}

func (c *Chunk) Set(x, y, z int, block Block) {
	c.Blocks[x+z*ChunkDepth+y*ChunkXZSize] = block
}
