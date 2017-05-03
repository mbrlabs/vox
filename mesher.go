// Copyright (c) 2017 Marcus Brummer.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vox

type Mesher interface {
	Generate(chunk *Chunk, bank *BlockBank) *RawMesh
}

type RawMesh struct {
	Positions []float32
	Indices   []uint16
	Normals   []float32
	Colors    []float32
}

type StupidMesher struct {
}

func (sm *StupidMesher) Generate(chunk *Chunk, bank *BlockBank) *RawMesh {
	mesh := &RawMesh{}

	// positions
	for x := 0; x < ChunkWidth; x++ {
		for z := 0; z < ChunkDepth; z++ {
			for y := 0; y < ChunkHeight; y++ {
				block := chunk.Get(x, y, z)
				if block.Active() {
					//fmt.Println(x, y, z)
					sm.addCube(float32(x), float32(y), float32(z), block, bank, mesh)
				}
			}
		}
	}

	// TODO remove and generate real normals
	mesh.Normals = append(mesh.Normals, 7)

	return mesh
}

func (sm *StupidMesher) addCube(x, y, z float32, block Block, bank *BlockBank, mesh *RawMesh) {
	var CubeSize float32 = 1.0

	// TODO check for overflow
	idxOffset := uint16(len(mesh.Positions) / 3)

	// positions
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

	// indices
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

	// colors
	blockType := bank.TypeOf(block)
	for i := 0; i < 8; i++ {
		mesh.Colors = append(mesh.Colors, blockType.Color.R, blockType.Color.G, blockType.Color.B)
	}

}
