package glm

import (
	"fmt"
	"math"
)

/*
   M00   M01   M02   M03
   M10   M11   M12   M13
   M20   M21   M22   M23
   M30   M31   M32   M33
*/

// Column 0
const m00 = 0
const m10 = 1
const m20 = 2
const m30 = 3

// Column 1
const m01 = 4
const m11 = 5
const m21 = 6
const m31 = 7

// Column 2
const m02 = 8
const m12 = 9
const m22 = 10
const m32 = 11

// Column 3
const m03 = 12
const m13 = 13
const m23 = 14
const m33 = 15

type Mat4 struct {
	Data [16]float32
}

func NewMat4(idt bool) *Mat4 {
	mat := &Mat4{}
	if idt {
		mat.Identity()
	}
	return mat
}

func (m *Mat4) Reset() *Mat4 {
	for i := 0; i < 16; i++ {
		m.Data[i] = 0
	}

	return m
}

func (m *Mat4) Identity() *Mat4 {
	m.Reset()
	m.Data[m00] = 1
	m.Data[m11] = 1
	m.Data[m22] = 1
	m.Data[m33] = 1
	return m
}

func (m *Mat4) Perspective(fov, aspectRatio, near, far float32) *Mat4 {
	m.Identity()

	q := float32(1.0 / math.Tan(ToRadians*0.5*float64(fov)))
	a := q / aspectRatio
	b := (near + far) / (near - far)
	c := (2.0 / near * far) / (near - far)

	m.Data[m00] = a
	m.Data[m11] = q
	m.Data[m22] = b
	m.Data[m32] = -1
	m.Data[m23] = c

	return m
}

func (m *Mat4) String() string {
	return fmt.Sprintf("Matrix4 {\n  %v %v %v %v\n  %v %v %v %v\n  %v %v %v %v\n  %v %v %v %v\n}\n",
		m.Data[m00], m.Data[m01], m.Data[m02], m.Data[m03],
		m.Data[m10], m.Data[m11], m.Data[m12], m.Data[m13],
		m.Data[m20], m.Data[m21], m.Data[m22], m.Data[m23],
		m.Data[m30], m.Data[m31], m.Data[m32], m.Data[m33])
}
