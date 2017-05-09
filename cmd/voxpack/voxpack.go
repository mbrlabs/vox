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
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	AtlasSize  = 512
	Padding    = 2
	AtlasJson  = "atlas.json"
	AtlasImage = "atlas.png"
)

func main() {
	dir := pwd()
	files := getImageFiles(dir)

	// exit if no images found
	if len(files) > 0 {
		fmt.Printf("Found %v images\n", len(files))
	} else {
		fmt.Println("No images")
		os.Exit(1)
	}

	packer := NewTexturePacker(loadImages(files), files)
	packer.pack()
}

func pwd() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func getImageFiles(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var filter []os.FileInfo
	for _, file := range files {
		name := file.Name()
		if !file.IsDir() && strings.HasSuffix(name, "png") && name != AtlasImage {
			filter = append(filter, file)
		}
	}

	return filter
}

func loadImages(files []os.FileInfo) []image.Image {
	fmt.Printf("Loading %v images...\n", len(files))

	images := make([]image.Image, 0)
	for _, info := range files {
		file, err := os.Open(info.Name())
		if err != nil {
			panic(err)
		}

		img, _, err := image.Decode(file)
		file.Close()
		if err != nil {
			panic(err)
		}

		images = append(images, img)
	}

	fmt.Println("Finished loading")
	return images
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type TextureRegion struct {
	Name   string  `json:"name"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	X      int     `json:"x"`
	Y      int     `json:"y"`
	U      float32 `json:"u"`
	V      float32 `json:"v"`
}

type TexturePacker struct {
	images        []image.Image
	fileInfo      []os.FileInfo
	atlas         *image.RGBA
	regions       []TextureRegion
	cursorX       int
	cursorY       int
	maxLineHeight int
}

func NewTexturePacker(images []image.Image, files []os.FileInfo) *TexturePacker {
	return &TexturePacker{
		images:   images,
		fileInfo: files,
		cursorX:  Padding,
		cursorY:  Padding,
	}
}

func (p *TexturePacker) pack() {
	bounds := image.Rect(0, 0, AtlasSize, AtlasSize)
	p.atlas = image.NewRGBA(bounds)
	p.regions = make([]TextureRegion, 0)

	// merge images
	for i, image := range p.images {
		p.addImage(image, p.fileInfo[i])
	}

	// write atlas
	out, err := os.OpenFile(AtlasImage, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	err = png.Encode(out, p.atlas)
	if err != nil {
		panic(err)
	}
	out.Close()

	// write atlas.json
	json, _ := json.Marshal(p.regions)
	err = ioutil.WriteFile(AtlasJson, json, 0644)
	if err != nil {
		panic(err)
	}

}

func (p *TexturePacker) addImage(src image.Image, info os.FileInfo) {
	srcWidth := src.Bounds().Max.X
	srcHeight := src.Bounds().Max.Y

	if p.cursorX+srcWidth+Padding >= AtlasSize {
		p.cursorY += p.maxLineHeight + Padding
		p.cursorX = Padding
		p.maxLineHeight = 0
	}

	if p.cursorY >= AtlasSize {
		panic("Output image not big enough")
	}

	for x := 0; x < srcWidth; x++ {
		for y := 0; y < srcHeight; y++ {
			p.putPixel(p.cursorX+x, p.cursorY+y, src.At(x, y))
		}
	}

	p.regions = append(p.regions,
		TextureRegion{
			Name:   strings.TrimSuffix(info.Name(), ".png"),
			Width:  srcWidth,
			Height: srcHeight,
			X:      p.cursorX,
			Y:      p.cursorY,
			U:      float32(float64(p.cursorX) / float64(AtlasSize)),
			V:      float32(float64(p.cursorY) / float64(AtlasSize)),
		},
	)

	p.cursorX += srcWidth + Padding
	p.maxLineHeight = max(p.maxLineHeight, srcHeight)
}

func (p *TexturePacker) putPixel(x, y int, pixel color.Color) {
	p.atlas.Set(x, AtlasSize-y, pixel)
}
