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
	"math"
)

const Radius = 10

type World struct {
	mesher    Mesher
	generator Generator
	bank      *BlockBank

	// Chunks are all chunks that are ready to be rendered
	Chunks map[ChunkPosition]*Chunk
	// These are ALL chunks that are currently loaded (not included chunks that are scheduled for removal)
	allChunks map[ChunkPosition]*Chunk
	// These are chunks that need their mesh to be regenerated
	meshingNeeded map[ChunkPosition]*Chunk
	// These are chunks that need thier mesh to be uploaded to OpenGL
	uploadNeeded map[ChunkPosition]*Chunk
	// These are chunks that need to be disposed
	disposeNeeded map[ChunkPosition]*Chunk

	MaxUploadsPerFrame int
}

// NewWorld creates a new world
func NewWorld(bank *BlockBank, mesher Mesher, generator Generator) *World {
	return &World{
		mesher:    mesher,
		generator: generator,
		bank:      bank,

		Chunks:        make(map[ChunkPosition]*Chunk),
		allChunks:     make(map[ChunkPosition]*Chunk),
		meshingNeeded: make(map[ChunkPosition]*Chunk),
		uploadNeeded:  make(map[ChunkPosition]*Chunk),
		disposeNeeded: make(map[ChunkPosition]*Chunk),

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
	w.meshingNeeded[chunk.Position] = chunk
}

// RemoveChunk schedules the cunk for removal.
func (w *World) RemoveChunk(x, y, z int) {
	chunk := w.allChunks[ChunkPosition{x, y, z}]
	if chunk != nil {
		// TODO: re-mesh neighbors
		chunk.unsetNeighbors()
		delete(w.allChunks, chunk.Position)
		delete(w.Chunks, chunk.Position)
		delete(w.meshingNeeded, chunk.Position)

		w.disposeNeeded[chunk.Position] = chunk
	}
}

// Update updates the chunk meshes
func (w *World) Update() {
	w.processDispose()
	w.processMeshing()
	w.processUploading()
}

func (w *World) processMeshing() {
	if len(w.meshingNeeded) > 0 {
		for _, c := range w.meshingNeeded {
			c.meshData = w.mesher.Generate(c, w.bank)
			if c.meshData != nil {
				w.uploadNeeded[c.Position] = c
			}
		}

		w.meshingNeeded = make(map[ChunkPosition]*Chunk, 0)
	}
}

func (w *World) processUploading() {
	if len(w.uploadNeeded) == 0 {
		return
	}

	count := w.MaxUploadsPerFrame
	for _, chunk := range w.uploadNeeded {
		delete(w.uploadNeeded, chunk.Position)

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
		count--
		if count == 0 {
			return
		}
	}
}

func (w *World) processDispose() {
	if len(w.disposeNeeded) > 0 {
		for _, c := range w.disposeNeeded {
			delete(w.disposeNeeded, c.Position)
			if c.Mesh != nil {
				c.Mesh.Dispose()
			}
		}
	}
}

func (w *World) processGenerating() {
	// TODO
}

type InfiniteWorldController struct {
	oldPos *ChunkPosition
	pos    *ChunkPosition

	cam   *Camera
	world *World

	initialized bool
}

func NewInifinteWorldController(cam *Camera, world *World) *InfiniteWorldController {
	c := &InfiniteWorldController{&ChunkPosition{}, &ChunkPosition{}, cam, world, false}

	return c
}

func (c *InfiniteWorldController) Update() {
	chunkX := int(c.cam.position.X) / ChunkWidth
	chunkY := int(c.cam.position.Y) / ChunkHeight
	chunkZ := int(c.cam.position.Z) / ChunkDepth
	c.pos.Set(chunkX, chunkY, chunkZ)

	if !c.initialized {
		c.init()
		c.initialized = true
		return
	}

	if !c.pos.Equals(c.oldPos) {
		// remove
		for _, chunk := range c.world.allChunks {
			if math.Abs(float64(c.pos.X-chunk.Position.X)) > Radius || math.Abs(float64(c.pos.Z-chunk.Position.Z)) > Radius {
				c.world.RemoveChunk(chunk.Position.X, chunk.Position.Y, chunk.Position.Z)
			}
		}

		doGenerate := make(map[ChunkPosition]bool)

		// x direction
		dx := c.pos.X - c.oldPos.X
		if dx > 0 {
			for x := chunkX + Radius - dx; x < chunkX+Radius; x++ {
				for z := chunkZ - Radius; z < chunkZ+Radius; z++ {
					doGenerate[ChunkPosition{x, 0, z}] = true
				}
			}
		} else if dx < 0 {
			for x := chunkX - Radius; x < chunkX-Radius-dx; x++ {
				for z := chunkZ - Radius; z < chunkZ+Radius; z++ {
					doGenerate[ChunkPosition{x, 0, z}] = true
				}
			}
		}

		// z direction
		dz := c.pos.Z - c.oldPos.Z
		if dz > 0 {
			for z := chunkZ + Radius; z >= chunkZ+Radius-dz; z-- {
				for x := chunkX - Radius; x < chunkX+Radius; x++ {
					doGenerate[ChunkPosition{x, 0, z}] = true
				}
			}
		} else if dz < 0 {
			for z := chunkZ - Radius - dz; z > chunkZ-Radius; z-- {
				for x := chunkX - Radius; x < chunkX+Radius; x++ {
					doGenerate[ChunkPosition{x, 0, z}] = true
				}
			}
		}

		for pos := range doGenerate {
			c.world.GenerateNewChunk(pos.X, pos.Y, pos.Z)
		}

	}

	c.oldPos.Set(c.pos.X, c.pos.Y, c.pos.Z)
}

func (c *InfiniteWorldController) init() {
	for x := -Radius; x < Radius; x++ {
		for z := -Radius; z < Radius; z++ {
			c.world.GenerateNewChunk(x, 0, z)
		}
	}
}
