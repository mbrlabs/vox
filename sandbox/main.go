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
	"runtime"

	"github.com/mbrlabs/vox"
)

const (
	windowTitle  = "Vox Sandbox"
	windowWidth  = 1024
	windowHeight = 768
)

// ----------------------------------------------------------------------------

func init() {
	runtime.LockOSThread()
}

func main() {
	window := vox.NewWindow(&vox.WindowConfig{
		Height:       windowHeight,
		Width:        windowWidth,
		Title:        windowTitle,
		Resizable:    false,
		Fullscreen:   false,
		Vsync:        true,
		HiddenCursor: true,
	})

	window.Start(&Sandbox{})
}
