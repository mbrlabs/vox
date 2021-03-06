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

	// these are the offset in world coordinates of the chunk
	xOffset := float32(chunk.Position.X) * ChunkWidth
	yOffset := float32(chunk.Position.Y) * ChunkHeight
	zOffset := float32(chunk.Position.Z) * ChunkDepth

	hasFace := false
	for x := 0; x < ChunkWidth; x++ {
		for z := 0; z < ChunkDepth; z++ {
			for y := 0; y < ChunkHeight; y++ {
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

				// left face
				hasFace = left == BlockNil && (chunk.left == nil || !chunk.left.Get(ChunkWidth-1, y, z).Active()) // check if adjacient chunk has occluding block
				hasFace = hasFace || left != BlockNil && !left.Active()                                           // check if there is a adjacient block in the same chunk
				if hasFace {
					cm.addLeftFace(xx, yy, zz, data, blockType)
				}

				// right face
				hasFace = right == BlockNil && (chunk.right == nil || !chunk.right.Get(0, y, z).Active())
				hasFace = hasFace || right != BlockNil && !right.Active()
				if hasFace {
					cm.addRightFace(xx, yy, zz, data, blockType)
				}

				// top face
				hasFace = top == BlockNil && (chunk.top == nil || !chunk.top.Get(x, 0, z).Active())
				hasFace = hasFace || top != BlockNil && !top.Active()
				if hasFace {
					cm.addTopFace(xx, yy, zz, data, blockType)
				}

				// bottom face
				hasFace = bottom == BlockNil && (chunk.bottom == nil || !chunk.bottom.Get(x, ChunkHeight-1, z).Active())
				hasFace = hasFace || bottom != BlockNil && !bottom.Active()
				if hasFace {
					cm.addBottomFace(xx, yy, zz, data, blockType)
				}

				// front face
				hasFace = front == BlockNil && (chunk.front == nil || !chunk.front.Get(x, y, 0).Active())
				hasFace = hasFace || front != BlockNil && !front.Active()
				if hasFace {
					cm.addFrontFace(xx, yy, zz, data, blockType)
				}

				// back face
				hasFace = back == BlockNil && (chunk.back == nil || !chunk.back.Get(x, y, ChunkDepth-1).Active())
				hasFace = hasFace || back != BlockNil && !back.Active()
				if hasFace {
					cm.addBackFace(xx, yy, zz, data, blockType)
				}
			}
		}
	}

	if len(data.Positions) == 0 {
		return nil
	}
	return data
}

func (cm *CulledMesher) addUvs(data *MeshData, region *TextureRegion) {
	uvs := &region.Uvs
	data.Uvs = append(data.Uvs,
		uvs[0].X, uvs[0].Y,
		uvs[1].X, uvs[1].Y,
		uvs[2].X, uvs[2].Y,
		uvs[3].X, uvs[3].Y,
	)
}

func (cm *CulledMesher) addLeftFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x, y, z-CubeSize,
		x, y, z,
		x, y+CubeSize, z,
		x, y+CubeSize, z-CubeSize,
	)
	data.Normals = append(data.Normals,
		-1, 0, 0,
		-1, 0, 0,
		-1, 0, 0,
		-1, 0, 0,
	)
	cm.addUvs(data, blockType.Side)
	data.IndexCount += 6
}

func (cm *CulledMesher) addRightFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x+CubeSize, y, z,
		x+CubeSize, y, z-CubeSize,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x+CubeSize, y+CubeSize, z,
	)
	data.Normals = append(data.Normals,
		1, 0, 0,
		1, 0, 0,
		1, 0, 0,
		1, 0, 0,
	)
	cm.addUvs(data, blockType.Side)
	data.IndexCount += 6
}

func (cm *CulledMesher) addTopFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x, y+CubeSize, z,
		x+CubeSize, y+CubeSize, z,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x, y+CubeSize, z-CubeSize,
	)
	data.Normals = append(data.Normals,
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
		0, 1, 0,
	)
	cm.addUvs(data, blockType.Top)
	data.IndexCount += 6
}

func (cm *CulledMesher) addBottomFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x, y, z,
		x+CubeSize, y, z,
		x+CubeSize, y, z-CubeSize,
		x, y, z-CubeSize,
	)
	data.Normals = append(data.Normals,
		0, -1, 0,
		0, -1, 0,
		0, -1, 0,
		0, -1, 0,
	)
	cm.addUvs(data, blockType.Bottom)
	data.IndexCount += 6
}

func (cm *CulledMesher) addFrontFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x, y, z,
		x+CubeSize, y, z,
		x+CubeSize, y+CubeSize, z,
		x, y+CubeSize, z,
	)
	data.Normals = append(data.Normals,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
		0, 0, 1,
	)
	cm.addUvs(data, blockType.Side)
	data.IndexCount += 6
}

func (cm *CulledMesher) addBackFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x, y, z-CubeSize,
		x+CubeSize, y, z-CubeSize,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x, y+CubeSize, z-CubeSize,
	)
	data.Normals = append(data.Normals,
		0, 0, -1,
		0, 0, -1,
		0, 0, -1,
		0, 0, -1,
	)
	cm.addUvs(data, blockType.Side)
	data.IndexCount += 6
}
