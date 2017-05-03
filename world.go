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

	Chunks    []*Chunk
	BlockBank *BlockBank
}

func NewWorld() *World {
	return &World{
		mesher:    &StupidMesher{},
		generator: &RandomGenerator{},
		BlockBank: NewBlockBank(),
	}
}

func (w *World) GenerateDebugWorld() {
	c := w.generator.GenerateChunkAt(0, 0, 0, w.BlockBank)
	mesh := w.mesher.Generate(c, w.BlockBank)
	vao := NewVao()
	vao.Load(mesh.Positions, mesh.Indices, mesh.Colors, mesh.Normals)
	c.Mesh = vao

	w.Chunks = append(w.Chunks, c)
}
