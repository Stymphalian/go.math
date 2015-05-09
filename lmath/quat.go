package lmath

import (
	"math"
)

// References
//http://mathworld.wolfram.com/Quaternion.html
//https://luckytoilet.wordpress.com/2014/11/24/visualizing-quaternions-with-unity/#comments
//http://www.gamasutra.com/view/feature/131686/rotating_objects_using_quaternions.php?page=1

// Strutcure to hold the quaternion
// When used as a rotation quaternion ensure that the quaternion is unit length.
// W is the angle in radians.
// X,Y,Z is the axis of rotation.
type Quat struct {
	W, X, Y, Z float64
}

var (
	QuatIdentity = &Quat{1, 0, 0, 0}
	QuatZero     = &Quat{0, 0, 0, 0}
)

// Return true if all elements are equal between both quaternions.
// Comparisons are done using an epsilon (< 0.0000001)
func (this *Quat) Eq(other *Quat) bool {
	return closeEq(this.X, other.X, epsilon) &&
		closeEq(this.Y, other.Y, epsilon) &&
		closeEq(this.Z, other.Z, epsilon) &&
		closeEq(this.W, other.W, epsilon)
}

// Set the component of the quaternion. Return this
func (this *Quat) Set(w, x, y, z float64) *Quat {
	this.W = w
	this.X = x
	this.Y = y
	this.Z = z
	return this
}

// Add the two quaternions (q + other).
// Returns a new quaternion with the result.
func (this *Quat) Add(other *Quat) *Quat {
	out := *this
	out.W += other.W
	out.X += other.X
	out.Y += other.Y
	out.Z += other.Z
	return &out
}

// Add two quaternionns (q + other) storing the result in this. Returns this.
func (this *Quat) AddIn(other *Quat) *Quat {
	this.W += other.W
	this.X += other.X
	this.Y += other.Y
	this.Z += other.Z
	return this
}

// Subtract the two quaternions (q - other).
// Returns a new quaternion with the result.
func (this *Quat) Sub(other *Quat) *Quat {
	out := *this
	return out.SubIn(other)
}

// Sub two quaternionns (q - other). Result is stored in this. Return this.
func (this *Quat) SubIn(other *Quat) *Quat {
	this.W -= other.W
	this.X -= other.X
	this.Y -= other.Y
	this.Z -= other.Z
	return this
}

// Multiply the two quaterions.
// Return a new quaterion with the result.
// Multiplying quaternions is NOT commutative (order matters!).
func (this *Quat) Mult(q *Quat) *Quat {
	out := *this
	return out.MultIn(q)
}

// Multiply 'this' with the other quaternion.
// Store the result into 'this'.
// Return a pointer to 'this'.
// Multiplying quaternions is NOT commutative (order matters!).
func (this *Quat) MultIn(other *Quat) *Quat {
	w := this.W*other.W - this.X*other.X - this.Y*other.Y - this.Z*other.Z
	x := this.W*other.X + this.X*other.W + this.Y*other.Z - this.Z*other.Y
	y := this.W*other.Y + this.Y*other.W - this.X*other.Z + this.Z*other.X
	z := this.W*other.Z + this.Z*other.W + this.X*other.Y - this.Y*other.X

	this.X = x
	this.Y = y
	this.Z = z
	this.W = w
	return this
}

// Add a scalar quantity to the quaternion.
// Returns a new quaternion with the result.
func (this *Quat) AddScalar(val float64) *Quat {
	out := *this
	return out.AddInScalar(val)
}

// Add a scalar quantity to the quaternion. Returns this.
func (this *Quat) AddInScalar(val float64) *Quat {
	this.W += val
	this.X += val
	this.Y += val
	this.Z += val
	return this
}

// Subtract a scalar quantity to the quaternion.
// Returns a new quaternion with the result.
func (this *Quat) SubScalar(val float64) *Quat {
	out := *this
	return out.SubInScalar(val)
}

// Add a scalar quantity to the quaternion. Returns this.
func (this *Quat) SubInScalar(val float64) *Quat {
	this.W -= val
	this.X -= val
	this.Y -= val
	this.Z -= val
	return this
}

// Multiply a scalar quantity to the quaternion.
// Returns a new quaternion with the result.
func (this *Quat) MultScalar(val float64) *Quat {
	out := *this
	return out.MultInScalar(val)
}

// Add a scalar quantity to the quaternion. Returns this.
func (this *Quat) MultInScalar(val float64) *Quat {
	this.W *= val
	this.X *= val
	this.Y *= val
	this.Z *= val
	return this
}

// Divide a scalar quantity to the quaternion.
// Returns a new quaternion with the result.
func (this *Quat) DivScalar(val float64) *Quat {
	out := *this
	return out.DivInScalar(val)
}

// Add a scalar quantity to the quaternion.
// Returns this.
func (this *Quat) DivInScalar(val float64) *Quat {
	this.W /= val
	this.X /= val
	this.Y /= val
	this.Z /= val
	return this
}

// Make this quaterion into a unit length quaternion.
// Returns a pointer to this.
func (this *Quat) ToUnit() *Quat {
	n := this.Norm()
	if closeEq(n, 0, epsilon) {
		return this
	}
	return this.DivInScalar(n)
}

// Returns the norm of this quaternion.
// sqrt(x^2 + y^2 + z^2 + w^2)
func (this *Quat) Norm() float64 {
	return math.Sqrt(this.NormSq())
}

// Returns the squared norm of this quaternion.
// So that we can save a math.Sqrt.
// x^2 + y^2 + z^2 + w^2
func (this *Quat) NormSq() float64 {
	return this.X*this.X +
		this.Y*this.Y +
		this.Z*this.Z +
		this.W*this.W
}

// Return a new quaternion which is the conjugate of this
// Conjugate is defined as [w,-x,-y,-z]
func (this *Quat) Conjugate() *Quat {
	out := *this
	return out.ConjugateIn()
}

// Set the quaternion to the conjugate
// Conjugate is defined as [w,-x,-y,-z]
func (this *Quat) ConjugateIn() *Quat {
	this.X = -this.X
	this.Y = -this.Y
	this.Z = -this.Z
	return this
}

// Return a new quaternion which is the inverse of this
func (this *Quat) Inverse() *Quat {
	out := *this
	return out.InverseIn()
}

// Set this quaternion as the inverse
func (this *Quat) InverseIn() *Quat {
	this.ConjugateIn()
	n := this.NormSq()
	if closeEq(n, 0, epsilon) {
		return this
	}
	return this.DivInScalar(n)
}

// Create a copy of this Quaternion and return the result.
func (this *Quat) Copy() *Quat {
	out := *this
	return &out
}

//==============================================================================

// Return the axis component of the quaternion
func (this *Quat) Axis() *Vec3 {
	angle := 2 * math.Acos(this.W)
	s := math.Sin(angle / 2)
	return &Vec3{this.X / s, this.Y / s, this.Z / s}
}

// return the angle component of the quaternion
func (this *Quat) Angle() float64 {
	return 2 * math.Acos(this.W)
}

//==============================================================================

// Apply this quaternion as a rotation to the vec3.
// Perform the operation q * v * q^-1
// Return a new vector with result
func (this *Quat) RotateVec3(v *Vec3) *Vec3 {
	vq := &Quat{0.0, v.X, v.Y, v.Z}
	rs := this.Mult(vq).MultIn(this.Inverse())
	return &Vec3{rs.X, rs.Y, rs.Z}
}

// Create a quaternion from the specified euler angles.
// Return this
// 	Perform the operation in the order.
// 	pitch(x) => yaw(y) => row(z)
func (this *Quat) FromEuler(pitch, yaw, roll float64) *Quat {
	yawQ := &Quat{math.Cos(yaw / 2), 0, math.Sin(yaw / 2), 0}
	pitchQ := &Quat{math.Cos(pitch / 2), math.Sin(pitch / 2), 0, 0}
	rollQ := &Quat{math.Cos(roll / 2), 0, 0, math.Sin(roll / 2)}

	// note must be applied in reverse order
	// pitch => yaw => roll
	*this = *(rollQ.MultIn(yawQ).MultIn(pitchQ))
	return this
}

// Set the quaternion as a rotation with with specified angle (radians) and axis.
// The axis should be normalized!
// Return this
func (this *Quat) FromAxisAngle(angle float64, x, y, z float64) *Quat {
	return this.Set(math.Cos(angle/2),
		math.Sin(angle/2)*x,
		math.Sin(angle/2)*y,
		math.Sin(angle/2)*z)
}

// Set the quaternion from the specified rotation matrix.
// Assumption is that the matrix is a valid rotation matrix.
// Matrix should be in right-hand coordinate system
// Pitch-Yaw-Roll euler angle formation
// Return this
func (this *Quat) FromMat4(mat *Mat4) *Quat {
	// Reference : http://www.flipcode.com/documents/matrfaq.html#Q55
	// 0  1  2  3
	// 4  5  6  7
	// 8  9  10 11
	// 12 13 14 15
	trace := mat.Get(0, 0) + mat.Get(1, 1) + mat.Get(2, 2) + 1

	if trace > 0 {
		s := 0.5 / math.Sqrt(trace)
		return this.Set(
			0.25/s,
			(mat.At(9)-mat.At(6))*s,
			(mat.At(2)-mat.At(8))*s,
			(mat.At(4)-mat.At(1))*s,
		)
	}

	// Find the column which has the maximum diagonal value
	max_col := 0
	champ := mat.Get(0, 0)
	for col := 1; col < 3; col += 1 {
		cand := mat.Get(col, col)
		if cand > champ {
			champ = cand
			max_col = col
		}
	}

	// TODO : UNTESTED!!!
	var w, x, y, z, s float64
	switch max_col {
	case 0:
		s = 2 * math.Sqrt(1.0+mat.At(0)-mat.At(5)-mat.At(10))
		x = 0.5 / 2
		y = (mat.At(4) + mat.At(1)) / s
		z = (mat.At(8) + mat.At(2)) / s
		w = (mat.At(9) + mat.At(6)) / s
	case 1:
		s = 2 * math.Sqrt(1.0+mat.At(5)-mat.At(0)-mat.At(10))
		x = (mat.At(4) + mat.At(1)) / s
		y = 0.5 / 2
		z = (mat.At(9) + mat.At(6)) / s
		w = (mat.At(8) + mat.At(2)) / s
	case 2:
		s = 2 * math.Sqrt(1.0+mat.At(10)-mat.At(0)-mat.At(5))
		x = (mat.At(8) + mat.At(2)) / s
		y = (mat.At(9) + mat.At(6)) / s
		z = 0.5 / 2
		w = (mat.At(4) + mat.At(1)) / s
	}

	return this.Set(w, x, y, z)
}

// Extract out the euler angles from the quaternion
// Extract out the angles assuming the quaterion is encoded
// as pitch -> yaw -> roll
// The returned euler angle may not match the same angle passed in using the
// FromEuler(),but it is guaranteed to be form an equivlent rotation quaternion
func (this *Quat) Euler() (pitch, yaw, roll float64) {
	// Reference http://www.euclideanspace.com/maths/geometry/rotations/conversions/quaternionToEuler/

	test := this.X*this.Y + this.Z*this.W
	if test > 0.499 { // singularity at north pole
		yaw = 2 * math.Atan2(this.X, this.W)
		roll = math.Pi / 2
		pitch = 0
		return
	}
	if test < -0.499 { // singularity at south pole
		yaw = -2 * math.Atan2(this.X, this.W)
		roll = -math.Pi / 2
		pitch = 0
		return
	}

	sqx := this.X * this.X
	sqy := this.Y * this.Y
	sqz := this.Z * this.Z
	yaw = math.Atan2(2*this.Y*this.W-2*this.X*this.Z, 1-2*sqy-2*sqz)
	roll = math.Asin(2 * test)
	pitch = math.Atan2(2*this.X*this.W-2*this.Y*this.Z, 1-2*sqx-2*sqz)
	return
}

// Return the axis and angle of this rotation quaternion
func (this *Quat) AxisAngle() (angle, x, y, z float64) {
	axis := this.Axis()
	return this.Angle(), axis.X, axis.Y, axis.Z
}

// Return a mat4 from the provided quaternion
func (this *Quat) Mat4() *Mat4 {
	// Reference
	// Derivation of the below matrix can be found here
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/quaternionToMatrix/index.htm
	//     1 - 2y² - 2z²    2yx - 2wz        2xz + 2wy
	// M=  2xy + 2wz        1 - 2x² - 2z²    2yz - 2wx
	//     2xz - 2wy        2yz + 2wx        1 - 2x² - 2y²

	w, x, y, z := this.W, this.X, this.Y, this.Z

	mat := &Mat4{}
	// 0 1 2 3
	// 4 5 6 7
	// 8 9 10 11
	// 12 13 14 15
	mat.mat[0] = 1 - 2*y*y - 2*z*z
	mat.mat[1] = 2*y*z - 2*w*z
	mat.mat[2] = 2*x*z + 2*w*y
	mat.mat[3] = 0

	mat.mat[4] = 2*x*y + 2*w*z
	mat.mat[5] = 1 - 2*x*x - 2*z*z
	mat.mat[6] = 2*y*z - 2*w*x
	mat.mat[7] = 0

	mat.mat[8] = 2*x*z - 2*w*y
	mat.mat[9] = 2*y*z + 2*w*x
	mat.mat[10] = 1 - 2*x*x - 2*y*y
	mat.mat[11] = 0

	mat.mat[12] = 0
	mat.mat[13] = 0
	mat.mat[14] = 0
	mat.mat[15] = 1
	return mat
}
