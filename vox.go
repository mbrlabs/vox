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

var (
	voxOnce sync.Once
	Vox     *vox
)

// ----------------------------------------------------------------------------

// Game todo
type Game interface {
	Disposable
	Create()
	Resize(width, height int)
	Render(delta float32)
	Update(delta float32)
}

// ----------------------------------------------------------------------------

type vox struct {
	win *Window
}

func setupVox(win *Window) {
	voxOnce.Do(func() {
		Vox = &vox{win: win}
	})
}

func (v *vox) DeltaTime() float32 {
	return v.win.deltaTime
}

func (v *vox) DeltaMouseX() float32 {
	return v.win.deltaX
}

func (v *vox) DeltaMouseY() float32 {
	return v.win.deltaY
}

func (v *vox) Exit() {
	v.win.exitRequested = true
}

// AddKeyListener todo
func (v *vox) AddKeyListener(listener KeyListener) {
	v.win.addKeyListener(listener)
}

// AddMouseListener todo
func (v *vox) AddMouseListener(listener MouseListener) {
	v.win.addMouseListener(listener)
}
