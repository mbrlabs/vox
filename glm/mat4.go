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

var (
	tmpMat4 *Mat4    = NewMat4(false)
	tmpVec3 *Vector3 = &Vector3{0, 0, 0}
)

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

func (m *Mat4) Set(data [16]float32) *Mat4 {
	for i := 0; i < 16; i++ {
		m.Data[i] = data[i]
	}
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

func (m *Mat4) Mul(o *Mat4) *Mat4 {
	var tmp [16]float32
	tmp[m00] = m.Data[m00]*o.Data[m00] + m.Data[m01]*o.Data[m10] + m.Data[m02]*o.Data[m20] + m.Data[m03]*o.Data[m30]
	tmp[m10] = m.Data[m10]*o.Data[m00] + m.Data[m11]*o.Data[m10] + m.Data[m12]*o.Data[m20] + m.Data[m13]*o.Data[m30]
	tmp[m20] = m.Data[m20]*o.Data[m00] + m.Data[m21]*o.Data[m10] + m.Data[m22]*o.Data[m20] + m.Data[m23]*o.Data[m30]
	tmp[m30] = m.Data[m30]*o.Data[m00] + m.Data[m31]*o.Data[m10] + m.Data[m32]*o.Data[m20] + m.Data[m33]*o.Data[m30]

	tmp[m01] = m.Data[m00]*o.Data[m01] + m.Data[m01]*o.Data[m11] + m.Data[m02]*o.Data[m21] + m.Data[m03]*o.Data[m31]
	tmp[m11] = m.Data[m10]*o.Data[m01] + m.Data[m11]*o.Data[m11] + m.Data[m12]*o.Data[m21] + m.Data[m13]*o.Data[m31]
	tmp[m21] = m.Data[m20]*o.Data[m01] + m.Data[m21]*o.Data[m11] + m.Data[m22]*o.Data[m21] + m.Data[m23]*o.Data[m31]
	tmp[m31] = m.Data[m30]*o.Data[m01] + m.Data[m31]*o.Data[m11] + m.Data[m32]*o.Data[m21] + m.Data[m33]*o.Data[m31]

	tmp[m02] = m.Data[m00]*o.Data[m02] + m.Data[m01]*o.Data[m12] + m.Data[m02]*o.Data[m22] + m.Data[m03]*o.Data[m32]
	tmp[m12] = m.Data[m10]*o.Data[m02] + m.Data[m11]*o.Data[m12] + m.Data[m12]*o.Data[m22] + m.Data[m13]*o.Data[m32]
	tmp[m22] = m.Data[m20]*o.Data[m02] + m.Data[m21]*o.Data[m12] + m.Data[m22]*o.Data[m22] + m.Data[m23]*o.Data[m32]
	tmp[m32] = m.Data[m30]*o.Data[m02] + m.Data[m31]*o.Data[m12] + m.Data[m32]*o.Data[m22] + m.Data[m33]*o.Data[m32]

	tmp[m03] = m.Data[m00]*o.Data[m03] + m.Data[m01]*o.Data[m13] + m.Data[m02]*o.Data[m23] + m.Data[m03]*o.Data[m33]
	tmp[m13] = m.Data[m10]*o.Data[m03] + m.Data[m11]*o.Data[m13] + m.Data[m12]*o.Data[m23] + m.Data[m13]*o.Data[m33]
	tmp[m23] = m.Data[m20]*o.Data[m03] + m.Data[m21]*o.Data[m13] + m.Data[m22]*o.Data[m23] + m.Data[m23]*o.Data[m33]
	tmp[m33] = m.Data[m30]*o.Data[m03] + m.Data[m31]*o.Data[m13] + m.Data[m32]*o.Data[m23] + m.Data[m33]*o.Data[m33]

	return m.Set(tmp)
}

func (m *Mat4) Translation(x, y, z float32) *Mat4 {
	m.Identity()
	m.Data[m03] = x
	m.Data[m13] = y
	m.Data[m23] = z
	return m
}

func (m *Mat4) Translate(x, y, z float32) *Mat4 {
	tmpMat4.Translation(x, y, z)
	return m.Mul(tmpMat4)
}

func (m *Mat4) Scaling(x, y, z float32) *Mat4 {
	m.Reset()
	m.Data[m00] = x
	m.Data[m11] = y
	m.Data[m22] = z
	return m
}

func (m *Mat4) Scale(x, y, z float32) *Mat4 {
	m.Data[m00] *= x
	m.Data[m11] *= y
	m.Data[m22] *= z
	return m
}

func (m *Mat4) Rotation(angle, xAxis, yAxis, zAxis float32) *Mat4 {
	if angle == 0 {
		return m
	}

	m.Identity()
	rad := float64(ToRadians * angle)
	c := float32(math.Cos(rad))
	s := float32(math.Sin(rad))
	omc := 1.0 - c

	m.Data[m00] = xAxis*xAxis*omc + c
	m.Data[m10] = yAxis*xAxis*omc - zAxis*s
	m.Data[m20] = xAxis*zAxis*omc + yAxis*s

	m.Data[m01] = yAxis*xAxis*omc + zAxis*s
	m.Data[m11] = yAxis*yAxis*omc + c
	m.Data[m21] = yAxis*zAxis*omc - xAxis*s

	m.Data[m02] = xAxis*zAxis*omc - yAxis*s
	m.Data[m12] = yAxis*zAxis*omc + xAxis*s
	m.Data[m22] = zAxis*zAxis*omc + c

	return m
}

func (m *Mat4) Rotate(angle, xAxis, yAxis, zAxis float32) *Mat4 {
	tmpMat4.Rotation(angle, xAxis, yAxis, zAxis)
	return m.Mul(tmpMat4)
}

func (m *Mat4) String() string {
	return fmt.Sprintf("Matrix4 {\n  %v %v %v %v\n  %v %v %v %v\n  %v %v %v %v\n  %v %v %v %v\n}\n",
		m.Data[m00], m.Data[m01], m.Data[m02], m.Data[m03],
		m.Data[m10], m.Data[m11], m.Data[m12], m.Data[m13],
		m.Data[m20], m.Data[m21], m.Data[m22], m.Data[m23],
		m.Data[m30], m.Data[m31], m.Data[m32], m.Data[m33])
}
