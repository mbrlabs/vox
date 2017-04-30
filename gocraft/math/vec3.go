package math

import "fmt"

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
