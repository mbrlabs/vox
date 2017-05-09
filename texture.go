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
	"os"

	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Pixmap struct {
	Data   []uint8
	Width  int32
	Height int32
}

func NewPixmap(path string) *Pixmap {
	// decode image
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// extract pixels
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	pixels := make([]uint8, 0)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			pixels = append(pixels, uint8(r/257), uint8(g/257), uint8(b/257))
		}
	}

	return &Pixmap{
		Data:   pixels,
		Width:  int32(width),
		Height: int32(height),
	}
}

type Texture struct {
	id     uint32
	width  int32
	height int32
}

func NewTexture(path string, genMipmaps bool) *Texture {
	pixmap := NewPixmap(path)

	// generate texture
	tex := &Texture{
		width:  pixmap.Width,
		height: pixmap.Height,
	}
	gl.GenTextures(1, &tex.id)

	// upload to gpu & generate mipmaps
	gl.BindTexture(gl.TEXTURE_2D, tex.id)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, tex.width, tex.height, 0, gl.RGB, gl.UNSIGNED_BYTE, gl.Ptr(pixmap.Data))
	if genMipmaps {
		gl.GenerateMipmap(gl.TEXTURE_2D)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
		gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_LOD_BIAS, -1)
	} else {
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	}
	gl.BindTexture(gl.TEXTURE_2D, 0)

	return tex
}

func (t *Texture) Width() int32 {
	return t.width
}

func (t *Texture) Height() int32 {
	return t.height
}

func (t *Texture) Bind() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, t.id)
}

func (t *Texture) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}
