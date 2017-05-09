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

type World struct {
	mesher    Mesher
	generator Generator

	Chunks map[ChunkPosition]*Chunk
}

func NewWorld() *World {
	return &World{
		Chunks:    make(map[ChunkPosition]*Chunk),
		mesher:    &CulledMesher{},
		generator: &FlatGenerator{},
	}
}

func (w *World) CreateChunk(x, y, z int, bank *BlockBank) {
	chunk := w.generator.GenerateChunkAt(x, y, z, bank)
	w.Chunks[chunk.Position] = chunk
}

func (w *World) LoadChunks(bank *BlockBank) {
	for _, chunk := range w.Chunks {
		meshData := w.mesher.Generate(chunk, w.Chunks, bank)
		if meshData != nil {
			mesh := NewMesh()
			mesh.Load(meshData)
			chunk.Mesh = mesh
		}
	}
}
