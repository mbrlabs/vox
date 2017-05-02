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

type Block uint8

const (
	blockActiveMask = 0x80 // 0b10000000
	blockTypeMask   = 0x7F // 0b01111111
)

func (b Block) Active() bool {
	return (blockActiveMask & b) == blockActiveMask
}

func (b Block) Activate(active bool) Block {
	if active {
		return b | blockActiveMask
	}
	return b & blockTypeMask
}

func (b Block) BlockType() uint8 {
	return uint8(b & blockTypeMask)
}

type BlockType struct {
	Color *Color
}
