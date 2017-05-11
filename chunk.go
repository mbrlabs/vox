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

import "fmt"

const (
	ChunkWidth  = 16
	ChunkDepth  = 16
	ChunkHeight = 16
	ChunkXZ     = ChunkWidth * ChunkDepth
	ChunkXYZ    = ChunkXZ * ChunkHeight
)

type ChunkPosition struct {
	X, Y, Z int
}

func (p *ChunkPosition) Set(x, y, z int) *ChunkPosition {
	p.X = x
	p.Y = y
	p.Z = z
	return p
}

func (p *ChunkPosition) String() string {
	return fmt.Sprintf("ChunkPosition{%v, %v, %v}", p.X, p.Y, p.Z)
}

type Chunk struct {
	Position ChunkPosition
	Blocks   [ChunkXYZ]Block
	Mesh     *Mesh
	meshData *MeshData

	left   *Chunk
	right  *Chunk
	front  *Chunk
	bottom *Chunk
	top    *Chunk
	back   *Chunk
}

func NewChunk(x, y, z int) *Chunk {
	return &Chunk{
		Position: ChunkPosition{x, y, z},
	}
}

func (c *Chunk) Get(x, y, z int) Block {
	idx := x + z*ChunkDepth + y*ChunkXZ
	if x < 0 || y < 0 || z < 0 || x >= ChunkWidth || y >= ChunkHeight || z >= ChunkDepth {
		return BlockNil
	}

	return c.Blocks[idx]
}

func (c *Chunk) Set(x, y, z int, block Block) {
	c.Blocks[x+z*ChunkDepth+y*ChunkXZ] = block
}

func (c *Chunk) IndexAt(x, y, z int) int {
	return x + z*ChunkDepth + y*ChunkXZ
}

func (c *Chunk) setNeighbors(chunks map[ChunkPosition]*Chunk) {
	chunkPos := &ChunkPosition{}
	c.left = chunks[*chunkPos.Set(c.Position.X-1, c.Position.Y, c.Position.Z)]
	if c.left != nil {
		c.left.right = c
	}
	c.right = chunks[*chunkPos.Set(c.Position.X+1, c.Position.Y, c.Position.Z)]
	if c.right != nil {
		c.right.left = c
	}
	c.top = chunks[*chunkPos.Set(c.Position.X, c.Position.Y+1, c.Position.Z)]
	if c.top != nil {
		c.top.bottom = c
	}
	c.bottom = chunks[*chunkPos.Set(c.Position.X, c.Position.Y-1, c.Position.Z)]
	if c.bottom != nil {
		c.bottom.top = c
	}
	c.front = chunks[*chunkPos.Set(c.Position.X, c.Position.Y, c.Position.Z+1)]
	if c.front != nil {
		c.front.back = c
	}
	c.back = chunks[*chunkPos.Set(c.Position.X, c.Position.Y, c.Position.Z-1)]
	if c.back != nil {
		c.back.front = c
	}
}

func (c *Chunk) unsetNeighbors() {
	if c.left != nil {
		c.left.right = nil
	}
	if c.right != nil {
		c.right.left = nil
	}
	if c.top != nil {
		c.top.bottom = nil
	}
	if c.bottom != nil {
		c.bottom.top = nil
	}
	if c.front != nil {
		c.front.back = nil
	}
	if c.back != nil {
		c.back.front = nil
	}
}
