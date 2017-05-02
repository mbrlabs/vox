package gocraft

type Mesher interface {
	Generate(chunk *Chunk) *RawMesh
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
	mesh := &RawMesh{}

	for x := 0; x < ChunkWidth; x++ {
		for z := 0; z < ChunkDepth; z++ {
			for y := 0; y < ChunkHeight; y++ {
				//block := chunk.GetBlock(x, y, z)
				//if block.
				sm.addCube(float32(x), float32(y), float32(z), mesh)
			}
		}
	}

	return mesh
}

func (sm *StupidMesher) addCube(x, y, z float32, mesh *RawMesh) {
	var CubeSize float32 = 1.0

	// TODO check for overflow
	idxOffset := uint16(len(mesh.Positions) / 3)

	mesh.Positions = append(mesh.Positions,
		// front positions
		x, y, z,
		x+CubeSize, y, z,
		x+CubeSize, y+CubeSize, z,
		x, y+CubeSize, z,

		// back positions
		x, y, z-CubeSize,
		x+CubeSize, y, z-CubeSize,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x, y+CubeSize, z-CubeSize,
	)

	mesh.Indices = append(mesh.Indices,
		// front
		idxOffset, idxOffset+1, idxOffset+2,
		idxOffset+2, idxOffset+3, idxOffset,
		// back
		idxOffset+5, idxOffset+4, idxOffset+7,
		idxOffset+7, idxOffset+6, idxOffset+5,
		// top
		idxOffset+3, idxOffset+2, idxOffset+6,
		idxOffset+6, idxOffset+7, idxOffset+3,
		// bottom
		idxOffset, idxOffset+1, idxOffset+5,
		idxOffset+5, idxOffset+4, idxOffset,
		// left
		idxOffset+4, idxOffset, idxOffset+3,
		idxOffset+3, idxOffset+7, idxOffset+4,
		// right
		idxOffset+1, idxOffset+5, idxOffset+6,
		idxOffset+6, idxOffset+2, idxOffset+1,
	)
}
