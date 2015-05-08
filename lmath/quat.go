package lmath

import (
	"math"
)

// References
//http://mathworld.wolfram.com/Quaternion.html
//https://luckytoilet.wordpress.com/2014/11/24/visualizing-quaternions-with-unity/#comments
//http://www.gamasutra.com/view/feature/131686/rotating_objects_using_quaternions.php?page=1

type Quat struct {
	w, x, y, z float64
}

// w is the angle component
// x,y,z is the vector
func (this *Quat) Set(w, x, y, z float64) *Quat {
	this.x = x
	this.y = y
	this.z = z
	this.w = w
	return this
}

// Multiply a constant value against every element of the Quaternion
// Returns a new Quaternion
func (this *Quat) MultConst(scale float64) *Quat {
	out := *this
	return out.MultInConst(scale)
}

// Multiply a constant value against every element of the Quaternion
// Place the result into the 'this'
// Return a pointer to 'this'
func (this *Quat) MultInConst(scale float64) *Quat {
	this.x *= scale
	this.y *= scale
	this.z *= scale
	this.w *= scale
	return this
}

// Returns a new quaternion with 'this' added with the other quaternion
func (this *Quat) Add(q *Quat) *Quat {
	out := *this
	return out.AddIn(q)
}

// Add the other quaternion.
// Store the result into 'this'.
// Return a pointer to 'this'
func (this *Quat) AddIn(q *Quat) *Quat {
	this.x += q.x
	this.y += q.y
	this.z += q.z
	this.w += q.w
	return this
}

// Multiply 'this' with the other quaternion.
// Return a new quaterion with the result.
// Multiplying quaternions is NOT commutative (order matters!)
func (this *Quat) Mult(q *Quat) *Quat {
	out := *this
	return out.MultIn(q)
}

// Multiply 'this' with the other quaternion.
// Store the result into 'this'
// Return a pointer to 'this'
// Multiplying quaternions is NOT commutative (order matters!)
func (this *Quat) MultIn(q *Quat) *Quat {
	w := this.w*q.w - this.x*q.x - this.y*q.y - this.z*q.z
	x := this.w*q.x + this.x*q.w + this.y*q.z - this.z*q.y
	y := this.w*q.y + this.y*q.w - this.x*q.z + this.z*q.x
	z := this.w*q.z + this.z*q.w + this.x*q.y - this.y*q.x

	this.x = x
	this.y = y
	this.z = z
	this.w = w
	return this
}

// Return a new quaternion which is the conjugate of 'this'
// Conjugate is defined as [w,-x,-y,-z]
func (this *Quat) Conjugate() *Quat {
	out := *this
	return out.ConjugateIn()
}

// Calc the conjugate of the quaternion and store the result in 'this'
func (this *Quat) ConjugateIn() *Quat {
	this.x = -this.x
	this.y = -this.y
	this.z = -this.z
	return this
}

func (this *Quat) InverseIn() *Quat {
	this.ConjugateIn()
	n := this.SqrdNorm()

	if closeEquals(n, 0, epsilon) {
		return this
	}

	return this.MultInConst(1 / n)
}

// Return a new quaternion which is the inverse of 'this'
func (this *Quat) Inverse() *Quat {
	out := *this
	conj := out.Conjugate()
	n := out.SqrdNorm()

	if closeEquals(n, 0, epsilon) {
		return &out
	}

	return conj.MultInConst(1 / n)
}

// Return true if all elements are equal between both quaternions
// comparisons are done using an epsilon (< 0.0000001)
func (this *Quat) Equals(q *Quat) bool {
	return closeEquals(this.x, q.x, epsilon) &&
		closeEquals(this.y, q.y, epsilon) &&
		closeEquals(this.z, q.z, epsilon) &&
		closeEquals(this.w, q.w, epsilon)
}

// Make this quaterion into a unit length quaternion
// Returns a pointer to 'this'
func (this *Quat) ToUnit() *Quat {
	m := this.Norm()
	if closeEquals(m, 0, epsilon) {
		return this
	}
	return this.MultInConst(1 / m)
}

// Returns the norm of this quaternion
// sqrt(x^2 + y^2 + z^2 + w^2)
func (this *Quat) Norm() float64 {
	return math.Sqrt(this.SqrdNorm())
}

// Returns the squared norm of this quaternion
// So that we can save a math.Sqrt
// x^2 + y^2 + z^2 + w^2
func (this *Quat) SqrdNorm() float64 {
	return this.x*this.x +
		this.y*this.y +
		this.z*this.z +
		this.w*this.w
}
