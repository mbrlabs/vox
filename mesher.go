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
	Generate(chunk *Chunk, bank *BlockBank) *MeshData
}

// ----------------------------------------------------------------------------

type CulledMesher struct {
}

func (cm *CulledMesher) Generate(chunk *Chunk, bank *BlockBank) *MeshData {
	data := &MeshData{}

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
					cm.addLeftFace(xx, yy, zz, data)
					cm.addFaceColors(blockType, data)
					data.IndexCount += 6
				}
				if right == BlockNil || !right.Active() {
					cm.addRightFace(xx, yy, zz, data)
					cm.addFaceColors(blockType, data)
					data.IndexCount += 6
				}
				if top == BlockNil || !top.Active() {
					cm.addTopFace(xx, yy, zz, data)
					cm.addFaceColors(blockType, data)
					data.IndexCount += 6
				}
				if bottom == BlockNil || !bottom.Active() {
					cm.addBottomFace(xx, yy, zz, data)
					cm.addFaceColors(blockType, data)
					data.IndexCount += 6
				}
				if front == BlockNil || !front.Active() {
					cm.addFrontFace(xx, yy, zz, data)
					cm.addFaceColors(blockType, data)
					data.IndexCount += 6
				}
				if back == BlockNil || !back.Active() {
					cm.addBackFace(xx, yy, zz, data)
					cm.addFaceColors(blockType, data)
					data.IndexCount += 6
				}
			}
		}
	}
	return data
}

func (cm *CulledMesher) addFaceColors(blockType *BlockType, data *MeshData) {
	for i := 0; i < 4; i++ {
		data.Colors = append(data.Colors, blockType.Color.R, blockType.Color.G, blockType.Color.B)
	}
}

func (cm *CulledMesher) addLeftFace(x, y, z float32, data *MeshData) {
	data.Positions = append(data.Positions,
		x, y, z-CubeSize,
		x, y, z,
		x, y+CubeSize, z,
		x, y+CubeSize, z-CubeSize,
	)
}

func (cm *CulledMesher) addRightFace(x, y, z float32, data *MeshData) {
	data.Positions = append(data.Positions,
		x+CubeSize, y, z,
		x+CubeSize, y, z-CubeSize,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x+CubeSize, y+CubeSize, z,
	)
}

func (cm *CulledMesher) addTopFace(x, y, z float32, data *MeshData) {
	data.Positions = append(data.Positions,
		x, y+CubeSize, z,
		x+CubeSize, y+CubeSize, z,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x, y+CubeSize, z-CubeSize,
	)
}

func (cm *CulledMesher) addBottomFace(x, y, z float32, data *MeshData) {
	data.Positions = append(data.Positions,
		x, y, z,
		x+CubeSize, y, z,
		x+CubeSize, y, z-CubeSize,
		x, y, z-CubeSize,
	)
}

func (cm *CulledMesher) addFrontFace(x, y, z float32, data *MeshData) {
	data.Positions = append(data.Positions,
		x, y, z,
		x+CubeSize, y, z,
		x+CubeSize, y+CubeSize, z,
		x, y+CubeSize, z,
	)
}

func (cm *CulledMesher) addBackFace(x, y, z float32, data *MeshData) {
	data.Positions = append(data.Positions,
		x+CubeSize, y, z-CubeSize,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x, y+CubeSize, z-CubeSize,
		x, y, z-CubeSize,
	)
}
