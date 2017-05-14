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

import opensimplex "github.com/ojrac/opensimplex-go"

type SimplexNoise struct {
	seed  int64
	noise *opensimplex.Noise
}

func NewSimplex(seed int64) *SimplexNoise {
	return &SimplexNoise{
		seed:  seed,
		noise: opensimplex.NewWithSeed(seed),
	}
}

func (s *SimplexNoise) Simplex2(x, y float64, octaves int, persistence, lacunarity float64) float64 {
	var freq float64 = 1
	var amp float64 = 1
	var max float64 = 1
	total := s.noise.Eval2(x, y)
	for i := 1; i < octaves; i++ {
		freq *= lacunarity
		amp *= persistence
		max += amp
		total += s.noise.Eval2(x*freq, y*freq) * amp
	}

	return (1 + total/max) / 2
}
