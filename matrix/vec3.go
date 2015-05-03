package matrix

import (
	//"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func (this *Vec3) Add(other *Vec3) *Vec3 {
	out := *this
	return out.AddIn(other)
}

func (this *Vec3) AddIn(other *Vec3) *Vec3 {
	this.X += other.X
	this.Y += other.Y
	this.Z += other.Z
	return this
}

func (this *Vec3) Sub(other *Vec3) *Vec3 {
	out := *this
	return out.SubIn(other)
}

func (this *Vec3) SubIn(other *Vec3) *Vec3 {
	this.X -= other.X
	this.Y -= other.Y
	this.Z -= other.Z
	return this
}

func (this *Vec3) Mult(scale float64) *Vec3 {
	out := *this
	return out.MultIn(scale)
}

func (this *Vec3) MultIn(scale float64) *Vec3 {
	this.X *= scale
	this.Y *= scale
	this.Z *= scale
	return this
}

func (this *Vec3) Div(scale float64) *Vec3 {
	out := *this
	return out.DivIn(scale)
}

func (this *Vec3) DivIn(scale float64) *Vec3 {
	this.X /= scale
	this.Y /= scale
	this.Z /= scale
	return this
}

func (this *Vec3) Outer(other *Vec3) *Vec3 {
	out := *this
	return out.OuterIn(other)
}
func (this *Vec3) OuterIn(other *Vec3) *Vec3 {
	this.X = this.X * other.X
	this.Y = this.Y * other.Y
	this.Z = this.Z * other.Z
	return this
}

func (this *Vec3) Dot(other *Vec3) float64 {
	return this.X*other.X +
		this.Y*other.Y +
		this.Z*other.Z
}

func (this *Vec3) Cross(other *Vec3) *Vec3 {
	out := *this
	return out.CrossIn(other)
}

func (this *Vec3) CrossIn(other *Vec3) *Vec3 {
	var out Vec3
	out.X = (this.Y*other.Z - other.Y*this.Z)
	out.Y = (this.X*other.Z - other.X*this.Z)
	out.Z = (this.X*other.Y - other.X*this.Y)
	*this = out
	return this
}

func (this *Vec3) Equals(other *Vec3) bool {
	return closeEquals(this.X, other.X, epsilon) &&
		closeEquals(this.Y, other.Y, epsilon) &&
		closeEquals(this.Z, other.Z, epsilon)
}

func (this *Vec3) Length() float64 {
	return math.Sqrt(this.X*this.X + this.Y*this.Y + this.Z*this.Z)
}

func (this *Vec3) NormalizeIn() *Vec3 {
	mag := this.Length()
	return this.DivIn(mag)
}

func (this *Vec3) Normalize() *Vec3 {
	out := *this
	(&out).NormalizeIn()
	return &out
}

func (this *Vec3) Set(x, y, z float64) *Vec3 {
	this.X = x
	this.Y = y
	this.Z = z
	return this
}

// Add(other *Vec) Vec
// AddIn(other *Vec)
// Sub(other *Vec) Vec
// SubIn(other *Vec)
// Mult(mag float64) Vec
// MultIn(mag float64 )
// Div(mag float64) Vec
// DivIn(mag float64 )
// Dot(other *Vec) float64
// Cross(other *Vec) Vec
// Equal(other *Vec)
// Length(other *Vec) float64
// NormalizeIn()
// Normalized() Vec
