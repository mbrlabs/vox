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

var (
	ColorWhite = NewColor(1, 1, 1, 1)
	ColorBlack = NewColor(0, 0, 0, 1)
	ColorRed   = NewColor(1, 0, 0, 1)
	ColorGreen = NewColor(0, 1, 0, 1)
	ColorBlue  = NewColor(0, 0, 1, 1)
	ColorTeal  = NewColor(0.5, 1, 1, 1)
)

type Color struct {
	R, G, B, A float32
}

func NewColor(r, g, b, a float32) *Color {
	return &Color{R: r, G: g, B: b, A: a}
}

func (c *Color) Copy() *Color {
	return NewColor(c.R, c.G, c.B, c.A)
}
