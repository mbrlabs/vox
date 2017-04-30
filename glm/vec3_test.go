package glm

import "testing"
import "github.com/mbrlabs/gocraft/assert"

func TestVector3Add(t *testing.T) {
	v := &Vector3{0, 0, 0}
	v.Add(1, 2, 3).Add(2, 3, 4).Add(-3, -5, -7)

	v2 := &Vector3{1, -1, 1}
	v.AddVector3(v2)

	assert.ApproxEquals(t, v.X, 1)
	assert.ApproxEquals(t, v.Y, -1)
	assert.ApproxEquals(t, v.Z, 1)
}

func TestVector3Sub(t *testing.T) {
	v := &Vector3{0, 0, 0}
	v.Sub(1, 2, 3).Sub(2, 3, 4).Sub(-3, -5, -7)

	v2 := &Vector3{1, -1, 1}
	v.SubVector3(v2)

	assert.ApproxEquals(t, v.X, -1)
	assert.ApproxEquals(t, v.Y, 1)
	assert.ApproxEquals(t, v.Z, -1)
}

func TestVector3Mul(t *testing.T) {
	v := &Vector3{1, 1, 1}
	v.Mul(1, 2, 3).Mul(2, 3, 4).Mul(-3, -5, -7)

	v2 := &Vector3{2, 2, 2}
	v.MulVector3(v2)

	assert.ApproxEquals(t, v.X, -12)
	assert.ApproxEquals(t, v.Y, -60)
	assert.ApproxEquals(t, v.Z, -168)
}

func TestVector3Div(t *testing.T) {
	v := &Vector3{64, 32, 16}
	v.Div(1, 1, 1).Div(4, 2, 2)

	v2 := &Vector3{1, 2, 2}
	v.DivVector3(v2)

	assert.ApproxEquals(t, v.X, 16)
	assert.ApproxEquals(t, v.Y, 8)
	assert.ApproxEquals(t, v.Z, 4)
}
