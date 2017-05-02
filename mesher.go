package gocraft

type Mesher interface {
	Generate(chunk *Chunk, vao *Vao)
}

type RawMesh struct {
	Positions []float32
	Indices   []uint16
	Normals   []float32
	Colors    []float32
}

type StupidMesher struct {
}

func (sm *StupidMesher) Generate(chunk *Chunk) *RawMesh {
	var pos []float32
	var ind []uint16
	var norm []float32
	var col []float32

	for x := 0; x < ChunkWidth; x++ {
		for z := 0; z < ChunkDepth; z++ {
			for y := 0; y < ChunkHeight; y++ {
				//block := chunk.GetBlock(x, y, z)
				//if block.
			}
		}
	}

	return &RawMesh{
		Positions: pos,
		Indices:   ind,
		Normals:   norm,
		Colors:    col,
	}
}
