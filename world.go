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
	bank      *BlockBank

	Chunks       map[ChunkPosition]*Chunk
	allChunks    map[ChunkPosition]*Chunk
	remeshNeeded []*Chunk
	uploadNeeded []*Chunk
	unloadNeeded []*Chunk
}

func NewWorld(bank *BlockBank, mesher Mesher, generator Generator) *World {
	return &World{
		mesher:    mesher,
		generator: generator,
		bank:      bank,

		Chunks:       make(map[ChunkPosition]*Chunk),
		allChunks:    make(map[ChunkPosition]*Chunk),
		remeshNeeded: make([]*Chunk, 0),
		uploadNeeded: make([]*Chunk, 0),
		unloadNeeded: make([]*Chunk, 0),
	}
}

// GenerateNewChunk generates a new chunk at the given chunk-coordinates.
// This does not perform, any OpenGL calls.
func (w *World) GenerateNewChunk(x, y, z int) {
	chunk := w.generator.GenerateChunkAt(x, y, z, w.bank)

	w.allChunks[chunk.Position] = chunk
	w.remeshNeeded = append(w.remeshNeeded, chunk)
}

func (w *World) RemoveChunk(x, y, z int) {
	chunk := w.allChunks[ChunkPosition{x, y, z}]
	if chunk != nil {
		delete(w.allChunks, chunk.Position)
		delete(w.Chunks, chunk.Position)
		w.unloadNeeded = append(w.unloadNeeded, chunk)
	}
}

func (w *World) Update() {

	// NOTE: later this must run in a worker goroutine
	// generate meshes
	if len(w.remeshNeeded) > 0 {
		for _, c := range w.remeshNeeded {
			c.meshData = w.mesher.Generate(c, w.allChunks, w.bank)
			if c.meshData != nil {
				w.uploadNeeded = append(w.uploadNeeded, c)
			}
		}

		w.clearRemeshSlice()
	}

	// NOTE: later this must run on the main goroutine
	// upload
	if len(w.uploadNeeded) > 0 {
		for _, c := range w.uploadNeeded {
			// upload to gpu
			if c.Mesh != nil {
				c.Mesh.Dispose()
			}
			c.Mesh = NewMesh()
			c.Mesh.Load(c.meshData)
			c.meshData = nil

			// add to list
			w.Chunks[c.Position] = c
		}
		w.clearUploadSlice()
	}

	// NOTE: later this must run on the main goroutine
	// unload
	if len(w.unloadNeeded) > 0 {
		for _, c := range w.unloadNeeded {
			// upload to gpu
			if c.Mesh != nil {
				c.Mesh.Dispose()
			}
		}
		w.clearUnloadSlice()
	}
}

func (w *World) clearRemeshSlice() {
	// TODO properly clear this (see: https://github.com/golang/go/wiki/SliceTricks)
	// also clearing slices in go really sucks...
	w.remeshNeeded = make([]*Chunk, 0)
}

func (w *World) clearUploadSlice() {
	// TODO properly clear this (see: https://github.com/golang/go/wiki/SliceTricks)
	// also clearing slices in go really sucks...
	w.uploadNeeded = make([]*Chunk, 0)
}

func (w *World) clearUnloadSlice() {
	// TODO properly clear this (see: https://github.com/golang/go/wiki/SliceTricks)
	// also clearing slices in go really sucks...
	w.unloadNeeded = make([]*Chunk, 0)
}
