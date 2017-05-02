package vox

import (
	"testing"
)

func TestBlockActive(t *testing.T) {
	var blocks [12]Block

	for _, b := range blocks {
		if b.Active() {
			t.Error()
		}
	}

	blocks[7] = blocks[7].Activate(true)
	if !blocks[7].Active() {
		t.Error()
	}
	blocks[7] = blocks[7].Activate(false)
	if blocks[7].Active() {
		t.Error()
	}
}

func TestBlockType(t *testing.T) {
	var block Block = 7

	if block.BlockType() != 7 {
		t.Error()
	}

	block = block.Activate(true)

	if block.BlockType() != 7 {
		t.Error()
	}

	if !block.Active() {
		t.Error()
	}

	block = block.Activate(false)

	if block.BlockType() != 7 {
		t.Error()
	}

	if block.Active() {
		t.Error()
	}
}
