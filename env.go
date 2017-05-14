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
	"github.com/mbrlabs/vox/glm"
)

type SunLight struct {
	Color     *Color
	Direction *glm.Vector3
	Intensity float32
}

type Fog struct {
	Color   *Color
	Density float32
}

type Environment struct {
	Sun *SunLight
	Fog *Fog
}

func NewEnvironment() *Environment {
	dir := &glm.Vector3{-1.5, 1, 1}
	sun := &SunLight{
		Color:     ColorWhite.Copy(),
		Direction: dir.Norm(),
		Intensity: 1,
	}

	fog := &Fog{
		Color:   ColorWhite.Copy(),
		Density: 0.1,
	}

	return &Environment{sun, fog}
}
