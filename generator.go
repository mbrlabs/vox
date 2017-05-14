// Copyright (c) 2017 Marcus Brummer.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License iss distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vox

import "math/rand"

type Generator interface {
	GenerateChunkAt(x, y, z int, bank *BlockBank) *Chunk
}

type FlatGenerator struct {
}

func (g *FlatGenerator) GenerateChunkAt(x, y, z int, bank *BlockBank) *Chunk {
	c := NewChunk(x, y, z)

	typeIdx := 0
	for i := 0; i < ChunkXYZ; i++ {
		c.Blocks[i] = c.Blocks[i].Activate(true)

		// block type
		if typeIdx >= len(bank.Types) {
			typeIdx = 0
		}
		c.Blocks[i] = c.Blocks[i].ChangeType(bank.Types[typeIdx])
		typeIdx++
	}

	return c
}

type SimplexGenerator struct {
	seed  int64
	noise *SimplexNoise
}

func NewSimplexGenerator(seed int64) Generator {
	return &SimplexGenerator{
		seed:  seed,
		noise: NewSimplex(seed),
	}
}

func (g *SimplexGenerator) GenerateChunkAt(xx, yy, zz int, bank *BlockBank) *Chunk {
	c := NewChunk(xx, yy, zz)

	t := uint8(1 + rand.Int()%3)

	worldX := float64(xx)
	worldZ := float64(zz)
	scaleX := 1.0 / float64(ChunkWidth)
	scaleZ := 1.0 / float64(ChunkDepth)

	typeIdx := 0
	for z := 0; z < ChunkDepth; z++ {
		for x := 0; x < ChunkWidth; x++ {
			simplex := g.noise.Simplex2(worldX+float64(x)*scaleX, worldZ+float64(z)*scaleZ, 3, 0.5, 2)
			height := int(simplex * ChunkHeight)
			for y := 0; y < ChunkHeight; y++ {
				i := c.IndexAt(x, y, z)
				c.Blocks[i] = c.Blocks[i].Activate(y < height)

				// block type
				if typeIdx >= len(bank.Types) {
					typeIdx = 0
				}
				c.Blocks[i] = c.Blocks[i].ChangeType(bank.typeMap[t])
				typeIdx++
			}
		}
	}

	return c
}
