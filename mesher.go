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

const CubeSize = 1.0

type Mesher interface {
	Generate(chunk *Chunk, bank *BlockBank) *RawMesh
}

type RawMesh struct {
	Positions []float32
	Indices   []uint16
	Colors    []float32
}

// ----------------------------------------------------------------------------
type StupidMesher struct {
}

func (sm *StupidMesher) Generate(chunk *Chunk, bank *BlockBank) *RawMesh {
	mesh := &RawMesh{}

	xOffset := float32(chunk.Position.X) * ChunkWidth
	yOffset := float32(chunk.Position.Y) * ChunkHeight
	zOffset := float32(chunk.Position.Z) * ChunkDepth

	// positions
	for x := 0; x < ChunkWidth; x++ {
		for z := 0; z < ChunkDepth; z++ {
			for y := 0; y < ChunkHeight; y++ {
				block := chunk.Get(x, y, z)
				if block.Active() {
					xx := xOffset + float32(x)
					yy := yOffset + float32(y)
					zz := zOffset + float32(z)
					sm.addCube(xx, yy, zz, block, bank, mesh)
				}
			}
		}
	}

	return mesh
}

func (sm *StupidMesher) addCube(x, y, z float32, block Block, bank *BlockBank, mesh *RawMesh) {
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

// ----------------------------------------------------------------------------

type CulledMesher struct {
}

func (cm *CulledMesher) Generate(chunk *Chunk, bank *BlockBank) *RawMesh {
	mesh := &RawMesh{}

	xOffset := float32(chunk.Position.X) * ChunkWidth
	yOffset := float32(chunk.Position.Y) * ChunkHeight
	zOffset := float32(chunk.Position.Z) * ChunkDepth

	for x := 0; x < ChunkWidth; x++ {
		for z := 0; z < ChunkDepth; z++ {
			for y := 0; y < ChunkHeight; y++ {
				//fmt.Println(x, y, z)

				// skip block if inactive
				block := chunk.Get(x, y, z)
				if !block.Active() {
					continue
				}

				blockType := bank.TypeOf(block)

				// get offsets
				xx := xOffset + float32(x)
				yy := yOffset + float32(y)
				zz := zOffset + float32(z)

				// get sourrounding neighbors
				left := chunk.Get(x-1, y, z)
				right := chunk.Get(x+1, y, z)
				top := chunk.Get(x, y+1, z)
				bottom := chunk.Get(x, y-1, z)
				front := chunk.Get(x, y, z+1)
				back := chunk.Get(x, y, z-1)

				// add new faces if adjaciant neighbor is inactive
				if left == BlockNil || !left.Active() {
					cm.addLeftFace(xx, yy, zz, mesh)
					cm.addFaceColors(blockType, mesh)
					cm.addFaceIndices(mesh)
				}
				if right == BlockNil || !right.Active() {
					cm.addRightFace(xx, yy, zz, mesh)
					cm.addFaceColors(blockType, mesh)
					cm.addFaceIndices(mesh)
				}
				if top == BlockNil || !top.Active() {
					cm.addTopFace(xx, yy, zz, mesh)
					cm.addFaceColors(blockType, mesh)
					cm.addFaceIndices(mesh)
				}
				if bottom == BlockNil || !bottom.Active() {
					cm.addBottomFace(xx, yy, zz, mesh)
					cm.addFaceColors(blockType, mesh)
					cm.addFaceIndices(mesh)
				}
				if front == BlockNil || !front.Active() {
					cm.addFrontFace(xx, yy, zz, mesh)
					cm.addFaceColors(blockType, mesh)
					cm.addFaceIndices(mesh)
				}
				if back == BlockNil || !back.Active() {
					cm.addBackFace(xx, yy, zz, mesh)
					cm.addFaceColors(blockType, mesh)
					cm.addFaceIndices(mesh)
				}
			}
		}
	}

	return mesh
}

func (cm *CulledMesher) addFaceIndices(mesh *RawMesh) {
	verts := uint16(len(mesh.Positions) / 3)
	mesh.Indices = append(mesh.Indices,
		verts-4, verts-3, verts-2,
		verts-2, verts-1, verts-4,
	)
}

func (cm *CulledMesher) addFaceColors(blockType *BlockType, mesh *RawMesh) {
	for i := 0; i < 4; i++ {
		mesh.Colors = append(mesh.Colors, blockType.Color.R, blockType.Color.G, blockType.Color.B)
	}
}

func (cm *CulledMesher) addLeftFace(x, y, z float32, mesh *RawMesh) {
	mesh.Positions = append(mesh.Positions,
		x, y, z-CubeSize,
		x, y, z,
		x, y+CubeSize, z,
		x, y+CubeSize, z-CubeSize,
	)
}

func (cm *CulledMesher) addRightFace(x, y, z float32, mesh *RawMesh) {
	mesh.Positions = append(mesh.Positions,
		x+CubeSize, y, z,
		x+CubeSize, y, z-CubeSize,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x+CubeSize, y+CubeSize, z,
	)
}

func (cm *CulledMesher) addTopFace(x, y, z float32, mesh *RawMesh) {
	mesh.Positions = append(mesh.Positions,
		x, y+CubeSize, z,
		x+CubeSize, y+CubeSize, z,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x, y+CubeSize, z-CubeSize,
	)
}

func (cm *CulledMesher) addBottomFace(x, y, z float32, mesh *RawMesh) {
	mesh.Positions = append(mesh.Positions,
		x, y, z,
		x+CubeSize, y, z,
		x+CubeSize, y, z-CubeSize,
		x, y, z-CubeSize,
	)
}

func (cm *CulledMesher) addFrontFace(x, y, z float32, mesh *RawMesh) {
	mesh.Positions = append(mesh.Positions,
		x, y, z,
		x+CubeSize, y, z,
		x+CubeSize, y+CubeSize, z,
		x, y+CubeSize, z,
	)
}

func (cm *CulledMesher) addBackFace(x, y, z float32, mesh *RawMesh) {
	mesh.Positions = append(mesh.Positions,
		x+CubeSize, y, z-CubeSize,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x, y+CubeSize, z-CubeSize,
		x, y, z-CubeSize,
	)
}
