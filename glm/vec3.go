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

package glm

import "fmt"
import "math"

type Vector3 struct {
	X, Y, Z float32
}

func (v *Vector3) Set(x, y, z float32) *Vector3 {
	v.X = x
	v.Y = y
	v.Z = z
	return v
}

func (v *Vector3) SetVector3(other *Vector3) *Vector3 {
	return v.Set(other.X, other.Y, other.Z)
}

func (v *Vector3) Add(x, y, z float32) *Vector3 {
	v.X += x
	v.Y += y
	v.Z += z
	return v
}

func (v *Vector3) AddVector3(other *Vector3) *Vector3 {
	return v.Add(other.X, other.Y, other.Z)
}

func (v *Vector3) Sub(x, y, z float32) *Vector3 {
	v.X -= x
	v.Y -= y
	v.Z -= z
	return v
}

func (v *Vector3) SubVector3(other *Vector3) *Vector3 {
	return v.Sub(other.X, other.Y, other.Z)
}

func (v *Vector3) Mul(x, y, z float32) *Vector3 {
	v.X *= x
	v.Y *= y
	v.Z *= z
	return v
}

func (v *Vector3) Scale(s float32) *Vector3 {
	v.X *= s
	v.Y *= s
	v.Z *= s
	return v
}

func (v *Vector3) Len2() float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vector3) Norm() *Vector3 {
	len2 := float64(v.Len2())
	if len2 == 0 || len2 == 1 {
		return v
	}
	return v.Scale(1.0 / float32(math.Sqrt(len2)))
}

func (v *Vector3) Cross(other *Vector3) *Vector3 {
	return v.Set(v.Y*other.Z-v.Z*other.Y, v.Z*other.X-v.X*other.Z, v.X*other.Y-v.Y*other.X)
}

func (v *Vector3) MulVector3(other *Vector3) *Vector3 {
	return v.Mul(other.X, other.Y, other.Z)
}

func (v *Vector3) Div(x, y, z float32) *Vector3 {
	v.X /= x
	v.Y /= y
	v.Z /= z
	return v
}

func (v *Vector3) DivVector3(other *Vector3) *Vector3 {
	return v.Div(other.X, other.Y, other.Z)
}

func (v *Vector3) String() string {
	return fmt.Sprintf("Vector3{%v, %v, %v}\n", v.X, v.Y, v.Z)
}

func (v *Vector3) Equals(x, y, z float32) bool {
	return v.X == x && v.Y == y && v.Z == z
}

func (v *Vector3) EqualsVector(other *Vector3) bool {
	return v.Equals(other.X, other.Y, other.Z)
}
