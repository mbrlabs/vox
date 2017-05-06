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

type RandomGenerator struct {
}

func (g *RandomGenerator) GenerateChunkAt(x, y, z int, bank *BlockBank) *Chunk {
	c := NewChunk()

	typeIdx := 0

	for i := 0; i < ChunkXYZ; i++ {
		// active or inactive
		if rand.Int()%2 == 0 {
			c.Blocks[i] = c.Blocks[i].Activate(true)
		} else {
			continue
		}

		// block type (color)
		if typeIdx >= len(bank.Types) {
			typeIdx = 0
		}
		c.Blocks[i] = c.Blocks[i].ChangeType(bank.Types[typeIdx])
		typeIdx++
	}

	return c
}
