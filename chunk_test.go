package vox

import "testing"

// this suceeeds if it does not panic
func TestChunkGet(t *testing.T) {
	var chunk Chunk
	for x := 0; x < ChunkWidth; x++ {
		for z := 0; z < ChunkDepth; z++ {
			for y := 0; y < ChunkHeight; y++ {
				chunk.Get(x, y, z)
			}
		}
	}
}
