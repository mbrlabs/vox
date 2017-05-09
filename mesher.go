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
	Generate(chunk *Chunk, chunks map[ChunkPosition]*Chunk, bank *BlockBank) *MeshData
}

// ----------------------------------------------------------------------------

type CulledMesher struct {
}

func (cm *CulledMesher) Generate(chunk *Chunk, chunks map[ChunkPosition]*Chunk, bank *BlockBank) *MeshData {
	data := &MeshData{}

	// these are the offset in world coordinates of the chunk
	xOffset := float32(chunk.Position.X) * ChunkWidth
	yOffset := float32(chunk.Position.Y) * ChunkHeight
	zOffset := float32(chunk.Position.Z) * ChunkDepth

	// get sourounding chunks
	chunkPos := &ChunkPosition{}
	leftChunk := chunks[*chunkPos.Set(chunk.Position.X-1, chunk.Position.Y, chunk.Position.Z)]
	rightChunk := chunks[*chunkPos.Set(chunk.Position.X+1, chunk.Position.Y, chunk.Position.Z)]
	topChunk := chunks[*chunkPos.Set(chunk.Position.X, chunk.Position.Y+1, chunk.Position.Z)]
	bottomChunk := chunks[*chunkPos.Set(chunk.Position.X, chunk.Position.Y-1, chunk.Position.Z)]
	frontChunk := chunks[*chunkPos.Set(chunk.Position.X, chunk.Position.Y, chunk.Position.Z+1)]
	backChunk := chunks[*chunkPos.Set(chunk.Position.X, chunk.Position.Y, chunk.Position.Z-1)]

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
				hasFace = left == BlockNil && (leftChunk == nil || !leftChunk.Get(ChunkWidth-1, y, z).Active()) // check if adjacient chunk has occluding block
				hasFace = hasFace || left != BlockNil && !left.Active()                                         // check if there is a adjacient block in the same chunk
				if hasFace {
					cm.addLeftFace(xx, yy, zz, data, blockType)
				}

				// right face
				hasFace = right == BlockNil && (rightChunk == nil || !rightChunk.Get(0, y, z).Active())
				hasFace = hasFace || right != BlockNil && !right.Active()
				if hasFace {
					cm.addRightFace(xx, yy, zz, data, blockType)
				}

				// top face
				hasFace = top == BlockNil && (topChunk == nil || !topChunk.Get(x, 0, z).Active())
				hasFace = hasFace || top != BlockNil && !top.Active()
				if hasFace {
					cm.addTopFace(xx, yy, zz, data, blockType)
				}

				// bottom face
				hasFace = bottom == BlockNil && (bottomChunk == nil || !bottomChunk.Get(x, ChunkHeight-1, z).Active())
				hasFace = hasFace || bottom != BlockNil && !bottom.Active()
				if hasFace {
					cm.addBottomFace(xx, yy, zz, data, blockType)
				}

				// front face
				hasFace = front == BlockNil && (frontChunk == nil || !frontChunk.Get(x, y, ChunkDepth-1).Active())
				hasFace = hasFace || front != BlockNil && !front.Active()
				if hasFace {
					cm.addFrontFace(xx, yy, zz, data, blockType)
				}

				// back face
				hasFace = back == BlockNil && (backChunk == nil || !backChunk.Get(x, y, 0).Active())
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

func (cm *CulledMesher) addLeftFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x, y, z-CubeSize,
		x, y, z,
		x, y+CubeSize, z,
		x, y+CubeSize, z-CubeSize,
	)
	data.Uvs = append(data.Uvs, 0, 0, 1, 0, 1, 1, 0, 1)
	data.IndexCount += 6
}

func (cm *CulledMesher) addRightFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x+CubeSize, y, z,
		x+CubeSize, y, z-CubeSize,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x+CubeSize, y+CubeSize, z,
	)
	data.Uvs = append(data.Uvs, 0, 0, 1, 0, 1, 1, 0, 1)
	data.IndexCount += 6
}

func (cm *CulledMesher) addTopFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x, y+CubeSize, z,
		x+CubeSize, y+CubeSize, z,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x, y+CubeSize, z-CubeSize,
	)
	data.Uvs = append(data.Uvs, 0, 0, 1, 0, 1, 1, 0, 1)
	data.IndexCount += 6
}

func (cm *CulledMesher) addBottomFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x, y, z,
		x+CubeSize, y, z,
		x+CubeSize, y, z-CubeSize,
		x, y, z-CubeSize,
	)
	data.Uvs = append(data.Uvs, 0, 0, 1, 0, 1, 1, 0, 1)
	data.IndexCount += 6
}

func (cm *CulledMesher) addFrontFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x, y, z,
		x+CubeSize, y, z,
		x+CubeSize, y+CubeSize, z,
		x, y+CubeSize, z,
	)
	data.Uvs = append(data.Uvs, 0, 0, 1, 0, 1, 1, 0, 1)
	data.IndexCount += 6
}

func (cm *CulledMesher) addBackFace(x, y, z float32, data *MeshData, blockType *BlockType) {
	data.Positions = append(data.Positions,
		x+CubeSize, y, z-CubeSize,
		x+CubeSize, y+CubeSize, z-CubeSize,
		x, y+CubeSize, z-CubeSize,
		x, y, z-CubeSize,
	)
	data.Uvs = append(data.Uvs, 0, 0, 1, 0, 1, 1, 0, 1)
	data.IndexCount += 6
}
