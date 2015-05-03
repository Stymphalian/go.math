package matrix

import (
	"math"
)

// References
//http://mathworld.wolfram.com/Quaternion.html
//https://luckytoilet.wordpress.com/2014/11/24/visualizing-quaternions-with-unity/#comments
//http://www.gamasutra.com/view/feature/131686/rotating_objects_using_quaternions.php?page=1

type q8n struct {
	w, x, y, z float64
}

// w is the angle component
// x,y,z is the vector
func (this *q8n) Set(w, x, y, z float64) *q8n {
	this.x = x
	this.y = y
	this.z = z
	this.w = w
	return this
}

// Multiply a constant value against every element of the Quaternion
// Returns a new Quaternion
func (this *q8n) MultConst(scale float64) *q8n {
	out := *this
	return out.MultInConst(scale)
}

// Multiply a constant value against every element of the Quaternion
// Place the result into the 'this'
// Return a pointer to 'this'
func (this *q8n) MultInConst(scale float64) *q8n {
	this.x *= scale
	this.y *= scale
	this.z *= scale
	this.w *= scale
	return this
}

// Returns a new quaternion with 'this' added with the other quaternion
func (this *q8n) Add(q *q8n) *q8n {
	out := *this
	return out.AddIn(q)
}

// Add the other quaternion.
// Store the result into 'this'.
// Return a pointer to 'this'
func (this *q8n) AddIn(q *q8n) *q8n {
	this.x += q.x
	this.y += q.y
	this.z += q.z
	this.w += q.w
	return this
}

// Multiply 'this' with the other quaternion.
// Return a new quaterion with the result.
// Multiplying quaternions is NOT commutative (order matters!)
func (this *q8n) Mult(q *q8n) *q8n {
	out := *this
	return out.MultIn(q)
}

// Multiply 'this' with the other quaternion.
// Store the result into 'this'
// Return a pointer to 'this'
// Multiplying quaternions is NOT commutative (order matters!)
func (this *q8n) MultIn(q *q8n) *q8n {
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
func (this *q8n) Conjugate() *q8n {
	out := *this
	out.x = -this.x
	out.y = -this.y
	out.z = -this.z
	return &out
}

// Return a new quaternion which is the inverse of 'this'
func (this *q8n) Inverse() *q8n {
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
func (this *q8n) Equals(q *q8n) bool {
	return closeEquals(this.x, q.x, epsilon) &&
		closeEquals(this.y, q.y, epsilon) &&
		closeEquals(this.z, q.z, epsilon) &&
		closeEquals(this.w, q.w, epsilon)
}

// Make this quaterion into a unit length quaternion
// Returns a pointer to 'this'
func (this *q8n) ToUnit() *q8n {
	m := this.Norm()
	if closeEquals(m, 0, epsilon) {
		return this
	}
	return this.MultInConst(1 / m)
}


// Returns the norm of this quaternion
// sqrt(x^2 + y^2 + z^2 + w^2)
func (this *q8n) Norm() float64 {
	return math.Sqrt(this.SqrdNorm())
}

// Returns the squared norm of this quaternion
// So that we can save a math.Sqrt
// x^2 + y^2 + z^2 + w^2
func (this *q8n) SqrdNorm() float64 {
	return this.x*this.x +
		this.y*this.y +
		this.z*this.z +
		this.w*this.w
}
