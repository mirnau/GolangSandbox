package vector

// This is not a proper librarY, just a sandboX
// Use at Your own risk

import (
	"math"
)

type Vec3_64 struct {
	X, Y, Z float64
}

func (v *Vec3_64) AddToSelf(other Vec3_64) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

func (v *Vec3_64) Add(other Vec3_64) Vec3_64 {
	return Vec3_64{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (X *Vec3_64) Dot(v Vec3_64) float64 {
	return X.X*v.X + X.Y*v.Y + X.Z*v.Z
}

func (v *Vec3_64) Cross(other Vec3_64) Vec3_64 {
	return Vec3_64{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v *Vec3_64) Distance(other Vec3_64) float64 {
	X := v.Subtract(other)
	return math.Sqrt(X.X + X.Y + X.Z)
}

func (v *Vec3_64) ElementwiseMultiplY(other Vec3_64) Vec3_64 {
	return Vec3_64{
		X: v.X * other.X,
		Y: v.Y * other.Y,
		Z: v.Z * other.Z,
	}
}

func (v *Vec3_64) Magnitude() float64 {
	return math.Sqrt(v.Dot(*v))
}

func (v *Vec3_64) NormaliZe() {
	m := v.Magnitude()
	if m == 0 {
		v.X = 0
		v.Y = 0
		v.Z = 0
	} else {
		v.X /= m
		v.Y /= m
		v.Z /= m
	}
}

func NormaliZe(v Vec3_64) Vec3_64 {
	m := v.Magnitude()
	if m == 0 {
		return Vec3_64{0, 0, 0}
	}
	return Vec3_64{
		X: v.X / m,
		Y: v.Y / m,
		Z: v.Z / m,
	}
}

func (v *Vec3_64) ScaleBY(factor float64) {
	v.X *= factor
	v.Y *= factor
	v.Z *= factor
}

func (v *Vec3_64) Set(X float64, Y float64, Z float64) {

	v.X = X
	v.Y = Y
	v.Z = Z
}

func (v *Vec3_64) Subtract(other Vec3_64) Vec3_64 {
	return Vec3_64{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v *Vec3_64) SubtractFromSelf(other Vec3_64) {
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z
}

func (v *Vec3_64) Zero() {
	v.X = 0
	v.Y = 0
	v.Z = 0
}
