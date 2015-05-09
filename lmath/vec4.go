package lmath

import (
    //"fmt"
    "math"
)

// A Vector 4 containing the three components
// X, Y, Z,W
type Vec4 struct {
    X, Y, Z,W float64
}

var (
    Vec4Right   = &Vec4{1, 0, 0,1}
    Vec4Up      = &Vec4{0, 1, 0,1}
    Vec4Forward = &Vec4{0, 0, 1,1}
    Vec4Zero    = &Vec4{0, 0, 0,0}
)

// Returns a new vector which is the result of adding 'this' with the
// other vector
func (this *Vec4) Add(other *Vec4) *Vec4 {
    out := *this
    return out.AddIn(other)
}

// Adds 'this' with the other vector.
// Store the result into 'this'
// Return a pointer to 'this'
func (this *Vec4) AddIn(other *Vec4) *Vec4 {
    this.X += other.X
    this.Y += other.Y
    this.Z += other.Z
    this.W += other.W
    return this
}

// Returns a new vector which is the result of subtracting 'this' with the
// other vector
func (this *Vec4) Sub(other *Vec4) *Vec4 {
    out := *this
    return out.SubIn(other)
}

// Subtracts'this' with the other vector.
// Store the result into 'this'
// Return a pointer to 'this'
func (this *Vec4) SubIn(other *Vec4) *Vec4 {
    this.X -= other.X
    this.Y -= other.Y
    this.Z -= other.Z
    this.W -= other.W
    return this
}

// Returns a new vector with the scalar added to every element
func (this *Vec4) AddScalar(scale float64) *Vec4 {
    out := *this
    return out.AddInScalar(scale)
}

// Add the scale to every element in the vector
// Return this
func (this *Vec4) AddInScalar(scale float64) *Vec4 {
    this.X += scale
    this.Y += scale
    this.Z += scale
    this.W += scale
    return this
}

// Returns a new vector with the scalar subtracted to every element
func (this *Vec4) SubScalar(scale float64) *Vec4 {
    out := *this
    return out.SubInScalar(scale)
}

// Subtract the scale from every element in the vector
// Return a pointer to 'this'
func (this *Vec4) SubInScalar(scale float64) *Vec4 {
    this.X -= scale
    this.Y -= scale
    this.Z -= scale
    this.W -= scale
    return this
}

// Returns a new vector where every element is multiplied by the scale
func (this *Vec4) MultScalar(scale float64) *Vec4 {
    out := *this
    return out.MultInScalar(scale)
}

// Multiply the each element of this vector with the scale value.
// Return a pointer to 'this'
func (this *Vec4) MultInScalar(scale float64) *Vec4 {
    this.X *= scale
    this.Y *= scale
    this.Z *= scale
    this.W *= scale
    return this
}

// Returns a new vector where every element is division by the scale
func (this *Vec4) DivScalar(scale float64) *Vec4 {
    out := *this
    return out.DivInScalar(scale)
}

// Divide the each element of this vector with the scale value.
// Return a pointer to 'this'
func (this *Vec4) DivInScalar(scale float64) *Vec4 {
    this.X /= scale
    this.Y /= scale
    this.Z /= scale
    this.W /= scale
    return this
}

// Do a pair-wise element multiplication with the provided vector
// Returns a new vector with the result
func (this *Vec4) Outer(other *Vec4) *Vec4 {
    out := *this
    return out.OuterIn(other)
}

// Do a element-wise multiplication with the provided vector
// Store the result into 'this'
// Return a pointer to 'this'
func (this *Vec4) OuterIn(other *Vec4) *Vec4 {
    this.X = this.X * other.X
    this.Y = this.Y * other.Y
    this.Z = this.Z * other.Z
    this.W = this.W * other.W
    return this
}

// Returns the Dot product between 'this' and the other vector
func (this *Vec4) Dot(other *Vec4) float64 {
    return this.X*other.X +
        this.Y*other.Y +
        this.Z*other.Z +
        this.W*other.W
}

// Return the length of the vector
// sqrt(x^2 + y^2 + z^2)
func (this *Vec4) Length() float64 {
    return math.Sqrt(this.X*this.X + this.Y*this.Y + this.Z*this.Z + this.W*this.W)
}

// Return the squared length of the vector
// x^2 + y^2 + z^2
func (this *Vec4) LengthSq() float64 {
    return this.X*this.X + this.Y*this.Y + this.Z*this.Z + this.W*this.W
}

// Checks for equality between the vectors.
// Equal is all elemnets are equal within an epsilon ( < 0.0000001)
func (this *Vec4) Eq(other *Vec4) bool {
    return closeEq(this.X, other.X, epsilon) &&
        closeEq(this.Y, other.Y, epsilon) &&
        closeEq(this.Z, other.Z, epsilon) &&
        closeEq(this.W, other.W, epsilon)
}

// Normalize the vector
// Return a pointer to 'this'
func (this *Vec4) NormalizeIn() *Vec4 {
    mag := this.Length()
    return this.DivInScalar(mag)
}

// Return a new vector which is the normalized version of 'this'
func (this *Vec4) Normalize() *Vec4 {
    out := *this
    (&out).NormalizeIn()
    return &out
}

// Set X,Y,Z,W parameters of the vector.
func (this *Vec4) Set(x, y, z, w float64) *Vec4 {
    this.X = x
    this.Y = y
    this.Z = z
    this.W = w
    return this
}

// Make a vector which is the projection of this onto other
func (this *Vec4) Proj(other *Vec4) *Vec4 {
    n := this.Length() * other.Length()
    return other.Normalize().MultScalar(this.Dot(other) / n)
}

// Return a copy of this vector
func (this *Vec4) Copy() *Vec4 {
    out := *this
    return &out
}

//==============================================================================
// Vector 4 specific methods


// Apply the matrix against the Vector
// Return a new vector with the result v*m
func (this *Vec4) MultMat4(right *Mat4) *Vec4 {
    // 0   1   2   3
    // 4   5   6   7
    // 8   9   10  11
    // 12  13  14  15
    return &Vec4{
        this.X*right.At(0) + this.Y*right.At(4) + this.Z*right.At(8)  +  this.W*right.At(12),
        this.X*right.At(1) + this.Y*right.At(5) + this.Z*right.At(9)  +  this.W*right.At(13),
        this.X*right.At(2) + this.Y*right.At(6) + this.Z*right.At(10) +  this.W*right.At(14),
        this.X*right.At(3) + this.Y*right.At(7) + this.Z*right.At(11) +  this.W*right.At(15),
    }
}
