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

const (
	ChunkWidth  = 16
	ChunkDepth  = 16
	ChunkHeight = 16
	ChunkXZ     = ChunkWidth * ChunkDepth
	ChunkXYZ    = ChunkXZ * ChunkHeight
)

type Chunk struct {
	Blocks [ChunkWidth * ChunkDepth * ChunkHeight]Block
	Mesh   *Vao
}

func NewChunk() *Chunk {
	return &Chunk{}
}

func (c *Chunk) Get(x, y, z int) Block {
	return c.Blocks[x+z*ChunkDepth+y*ChunkXZ]
}

func (c *Chunk) Set(x, y, z int, block Block) {
	c.Blocks[x+z*ChunkDepth+y*ChunkXZ] = block
}
