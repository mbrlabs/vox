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

import (
	"testing"

	"github.com/mbrlabs/vox/assert"
)

func TestMat4Identity(t *testing.T) {
	m := NewMat4(false)
	for i := 0; i < 16; i++ {
		m.Data[i] = 123
	}
	m.Identity()

	assert.ApproxEquals(t, m.Data[m00], 1)
	assert.ApproxEquals(t, m.Data[m11], 1)
	assert.ApproxEquals(t, m.Data[m22], 1)
	assert.ApproxEquals(t, m.Data[m33], 1)

	assert.ApproxEquals(t, m.Data[m01], 0)
	assert.ApproxEquals(t, m.Data[m02], 0)
	assert.ApproxEquals(t, m.Data[m03], 0)
	assert.ApproxEquals(t, m.Data[m10], 0)
	assert.ApproxEquals(t, m.Data[m12], 0)
	assert.ApproxEquals(t, m.Data[m13], 0)
	assert.ApproxEquals(t, m.Data[m20], 0)
	assert.ApproxEquals(t, m.Data[m21], 0)
	assert.ApproxEquals(t, m.Data[m23], 0)
	assert.ApproxEquals(t, m.Data[m30], 0)
	assert.ApproxEquals(t, m.Data[m31], 0)
	assert.ApproxEquals(t, m.Data[m32], 0)
}

func TestMat4Mul(t *testing.T) {
	left := NewMat4(false)
	left.Data[m00] = 1
	left.Data[m01] = 0
	left.Data[m02] = 2
	left.Data[m03] = 1
	left.Data[m10] = 0
	left.Data[m11] = 1
	left.Data[m12] = 6
	left.Data[m13] = -10
	left.Data[m20] = 3
	left.Data[m21] = 4
	left.Data[m22] = 7
	left.Data[m23] = 0
	left.Data[m30] = 7
	left.Data[m31] = -38
	left.Data[m32] = 0
	left.Data[m33] = 1

	right := NewMat4(false)
	right.Data[m00] = 1
	right.Data[m01] = 4
	right.Data[m02] = 12
	right.Data[m03] = 3
	right.Data[m10] = 0
	right.Data[m11] = 1
	right.Data[m12] = 0
	right.Data[m13] = 1
	right.Data[m20] = 6
	right.Data[m21] = 7
	right.Data[m22] = 0
	right.Data[m23] = 0
	right.Data[m30] = 12
	right.Data[m31] = 0
	right.Data[m32] = 6
	right.Data[m33] = 11

	left.Mul(right)

	assert.ApproxEquals(t, left.Data[m00], 25)
	assert.ApproxEquals(t, left.Data[m10], -84)
	assert.ApproxEquals(t, left.Data[m20], 45)
	assert.ApproxEquals(t, left.Data[m30], 19)

	assert.ApproxEquals(t, left.Data[m01], 18)
	assert.ApproxEquals(t, left.Data[m11], 43)
	assert.ApproxEquals(t, left.Data[m21], 65)
	assert.ApproxEquals(t, left.Data[m31], -10)

	assert.ApproxEquals(t, left.Data[m02], 18)
	assert.ApproxEquals(t, left.Data[m12], -60)
	assert.ApproxEquals(t, left.Data[m22], 36)
	assert.ApproxEquals(t, left.Data[m32], 90)

	assert.ApproxEquals(t, left.Data[m03], 14)
	assert.ApproxEquals(t, left.Data[m13], -109)
	assert.ApproxEquals(t, left.Data[m23], 13)
	assert.ApproxEquals(t, left.Data[m33], -6)
}
