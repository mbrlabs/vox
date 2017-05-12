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

import (
	"sync"
)

type World struct {
	mesher    Mesher
	generator Generator
	bank      *BlockBank

	Chunks       map[ChunkPosition]*Chunk
	allChunks    map[ChunkPosition]*Chunk
	remeshNeeded []*Chunk
	uploadNeeded []*Chunk
	unloadNeeded []*Chunk

	remeshMutex sync.Mutex

	MaxUploadsPerFrame int
}

// NewWorld creates a new world
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

		MaxUploadsPerFrame: 8,
	}
}

// GenerateNewChunk generates a new chunk at the given chunk-coordinates.
// This does not perform, any OpenGL calls.
func (w *World) GenerateNewChunk(x, y, z int) {
	// generate chunk
	chunk := w.generator.GenerateChunkAt(x, y, z, w.bank)
	chunk.setNeighbors(w.allChunks)
	// TODO: re-mesh neighbors

	// add to world
	w.allChunks[chunk.Position] = chunk
	w.remeshNeeded = append(w.remeshNeeded, chunk)
}

// RemoveChunk schedules the cunk for removal.
func (w *World) RemoveChunk(x, y, z int) {
	chunk := w.allChunks[ChunkPosition{x, y, z}]
	if chunk != nil {
		// TODO: re-mesh neighbors
		chunk.unsetNeighbors()
		delete(w.allChunks, chunk.Position)
		delete(w.Chunks, chunk.Position)
		w.unloadNeeded = append(w.unloadNeeded, chunk)
	}
}

// Update updates the chunk meshes
func (w *World) Update() {
	w.processUnloading()
	w.processMeshing()
	w.processUploading()
}

func (w *World) processMeshing() {
	if len(w.remeshNeeded) > 0 {
		for _, c := range w.remeshNeeded {
			c.meshData = w.mesher.Generate(c, w.bank)
			if c.meshData != nil {
				w.uploadNeeded = append(w.uploadNeeded, c)
			}
		}

		w.remeshNeeded = make([]*Chunk, 0)
	}
}

func (w *World) processUploading() {
	if len(w.uploadNeeded) == 0 {
		return
	}

	for i := 0; i < w.MaxUploadsPerFrame; i++ {
		// pop chunk
		chunk := w.uploadNeeded[0]
		w.uploadNeeded = w.uploadNeeded[1:]

		// dispose old mesh
		if chunk.Mesh != nil {
			chunk.Mesh.Dispose()
		}

		// upload new mesh
		chunk.Mesh = NewMesh()
		chunk.Mesh.Load(chunk.meshData)
		chunk.meshData = nil

		// add to list
		w.Chunks[chunk.Position] = chunk

		// done?
		if len(w.uploadNeeded) == 0 {
			return
		}
	}
}

func (w *World) processUnloading() {
	if len(w.unloadNeeded) > 0 {
		for _, c := range w.unloadNeeded {
			// upload to gpu
			if c.Mesh != nil {
				c.Mesh.Dispose()
			}
		}
		w.unloadNeeded = make([]*Chunk, 0)
	}
}

func (w *World) processGenerating() {
	// TODO
}
