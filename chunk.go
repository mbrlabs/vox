package vox

const (
	ChunkWidth  = 16
	ChunkDepth  = 16
	ChunkHeight = 16
	ChunkXZ     = ChunkWidth * ChunkDepth
	ChunkXYZ    = ChunkXZ * ChunkHeight
)

type Chunk struct {
	Blocks [ChunkWidth * ChunkDepth * ChunkHeight]Block
}

func NewChunk() *Chunk {
	c := &Chunk{}
	for i := 0; i < ChunkXYZ; i++ {
		c.Blocks[i] = c.Blocks[i].Activate(true)
	}

	return c
}

func (c *Chunk) Get(x, y, z int) Block {
	return c.Blocks[x+z*ChunkDepth+y*ChunkXZ]
}

func (c *Chunk) Set(x, y, z int, block Block) {
	c.Blocks[x+z*ChunkDepth+y*ChunkXZ] = block
}
