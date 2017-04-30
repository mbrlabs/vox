package glm

import (
	"testing"
)

func TestMat4Identity(t *testing.T) {
	m := NewMat4(false)
	for i := 0; i < 16; i++ {
		m.Data[i] = 123
	}
	m.Identity()

	if m.Data[m00] != 1 || m.Data[m11] != 1 || m.Data[m22] != 1 || m.Data[m33] != 1 {
		t.Error()
	}

	if m.Data[m01] != 0 || m.Data[m02] != 0 || m.Data[m03] != 0 ||
		m.Data[m10] != 0 || m.Data[m12] != 0 || m.Data[m13] != 0 ||
		m.Data[m20] != 0 || m.Data[m21] != 0 || m.Data[m23] != 0 ||
		m.Data[m30] != 0 || m.Data[m31] != 0 || m.Data[m32] != 0 {
		t.Error()
	}
}
