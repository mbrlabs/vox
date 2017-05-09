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

package main

import (
	"github.com/mbrlabs/vox"
)

const (
	TypeBrick   = 0x01
	TypeGrass   = 0x02
	TypeBedrock = 0x03
)

func createBlockTypes(atlas *vox.TextureAtlas) []*vox.BlockType {
	types := make([]*vox.BlockType, 0)

	brick, _ := atlas.Regions["brick"]
	bedrock, _ := atlas.Regions["bedrock"]
	grassTop, _ := atlas.Regions["grass_top"]
	grassSide, _ := atlas.Regions["grass_side"]

	types = append(types,
		&vox.BlockType{TypeBrick, brick, brick, brick},
		&vox.BlockType{TypeBedrock, bedrock, bedrock, bedrock},
		&vox.BlockType{TypeGrass, grassTop, grassTop, grassSide},
	)

	return types
}
