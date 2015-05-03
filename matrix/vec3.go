package matrix

import (
	//"fmt"
	"math"
)


// A Vector 3 containing the three components
// X, Y, Z
type Vec3 struct {
	X, Y, Z float64
}

// Returns a new vector which is the result of adding 'this' with the
// other vector
func (this *Vec3) Add(other *Vec3) *Vec3 {
	out := *this
	return out.AddIn(other)
}

// Adds 'this' with the other vector.
// Store the result into 'this'
// Return a pointer to 'this'
func (this *Vec3) AddIn(other *Vec3) *Vec3 {
	this.X += other.X
	this.Y += other.Y
	this.Z += other.Z
	return this
}

// Returns a new vector which is the result of subtracting 'this' with the
// other vector
func (this *Vec3) Sub(other *Vec3) *Vec3 {
	out := *this
	return out.SubIn(other)
}

// Subtracts'this' with the other vector.
// Store the result into 'this'
// Return a pointer to 'this'
func (this *Vec3) SubIn(other *Vec3) *Vec3 {
	this.X -= other.X
	this.Y -= other.Y
	this.Z -= other.Z
	return this
}

// Returns a new vector where every element is multiplied by the scale
func (this *Vec3) Mult(scale float64) *Vec3 {
	out := *this
	return out.MultIn(scale)
}

// Multiply the each element of this vector with the scale value.
// Return a pointer to 'this'
func (this *Vec3) MultIn(scale float64) *Vec3 {
	this.X *= scale
	this.Y *= scale
	this.Z *= scale
	return this
}

// Returns a new vector where every element is division by the scale
func (this *Vec3) Div(scale float64) *Vec3 {
	out := *this
	return out.DivIn(scale)
}

// Divide the each element of this vector with the scale value.
// Return a pointer to 'this'
func (this *Vec3) DivIn(scale float64) *Vec3 {
	this.X /= scale
	this.Y /= scale
	this.Z /= scale
	return this
}

// Do a pair-wise element multiplication with the provided vector
// Returns a new vector with the result
func (this *Vec3) Outer(other *Vec3) *Vec3 {
	out := *this
	return out.OuterIn(other)
}

// Do a element-wise multiplication with the provided vector
// Store the result into 'this'
// Return a pointer to 'this'
func (this *Vec3) OuterIn(other *Vec3) *Vec3 {
	this.X = this.X * other.X
	this.Y = this.Y * other.Y
	this.Z = this.Z * other.Z
	return this
}


// Returns the Dot product between 'this' and the other vector
func (this *Vec3) Dot(other *Vec3) float64 {
	return this.X*other.X +
		this.Y*other.Y +
		this.Z*other.Z
}

// Returns a new vector which is the Cross product with 'this' X 'other'
func (this *Vec3) Cross(other *Vec3) *Vec3 {
	out := *this
	return out.CrossIn(other)
}

// Take the cross product between 'this' X 'other'
// Store the result into 'this'
// Return a pointer to 'this'
func (this *Vec3) CrossIn(other *Vec3) *Vec3 {
	var out Vec3
	out.X = (this.Y*other.Z - other.Y*this.Z)
	out.Y = (this.X*other.Z - other.X*this.Z)
	out.Z = (this.X*other.Y - other.X*this.Y)
	*this = out
	return this
}


// Checks for equality between the vectors.
// Equal is all elemnets are equal within an epsilon ( < 0.0000001)
func (this *Vec3) Equals(other *Vec3) bool {
	return closeEquals(this.X, other.X, epsilon) &&
		closeEquals(this.Y, other.Y, epsilon) &&
		closeEquals(this.Z, other.Z, epsilon)
}


// Return the length of the vector
// sqrt(x^2 + y^2 + z^2)
func (this *Vec3) Length() float64 {
	return math.Sqrt(this.X*this.X + this.Y*this.Y + this.Z*this.Z)
}

// Normalize the vector
// Return a pointer to 'this'
func (this *Vec3) NormalizeIn() *Vec3 {
	mag := this.Length()
	return this.DivIn(mag)
}

// Return a new vector which is the normalized version of 'this'
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
