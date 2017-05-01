package gocraft

import "testing"

func TestBlock(t *testing.T) {
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
