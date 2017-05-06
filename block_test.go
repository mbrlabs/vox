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
	grass := &BlockType{ID: 1, Color: nil}
	water := &BlockType{ID: 2, Color: nil}

	var block Block
	block = block.ChangeType(grass)

	if block.TypeID() != grass.ID {
		t.Error()
	}

	block = block.Activate(true)

	if block.TypeID() != grass.ID {
		t.Error()
	}

	if !block.Active() {
		t.Error()
	}

	block = block.Activate(false)

	if block.TypeID() != grass.ID {
		t.Error()
	}

	if block.Active() {
		t.Error()
	}

	block = block.ChangeType(water)

	if block.TypeID() != water.ID {
		t.Error()
	}

}
