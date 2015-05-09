package lmath

import (
	"fmt"
	"math"
)

const (
	mat4Dim = 4
)

var (
	Mat4Identity = &Mat4{[16]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}}
)

type Mat4 struct {
	mat [16]float64
}

// New Mat4 with the given values
// Row-Order
func NewMat4(
	v11, v12, v13, v14,
	v21, v22, v23, v24,
	v31, v32, v33, v34,
	v41, v42, v43, v44 float64) *Mat4 {

	// 0   1   2   3
	// 4   5   6   7
	// 8   9   10  11
	// 12  13  14  15
	out := Mat4{}
	out.mat[0] = v11
	out.mat[1] = v12
	out.mat[2] = v13
	out.mat[3] = v14
	out.mat[4] = v21
	out.mat[5] = v22
	out.mat[6] = v23
	out.mat[7] = v24
	out.mat[8] = v31
	out.mat[9] = v32
	out.mat[10] = v33
	out.mat[11] = v34
	out.mat[12] = v41
	out.mat[13] = v42
	out.mat[14] = v43
	out.mat[15] = v44

	return &out
}

// Load the matrix with 16 floats
// Specified in Row-Major order
func (this *Mat4) Load(mat [16]float64) *Mat4 {
	this.mat = mat
	return this
}

// Retrieve a 16 float array of all the values of the matrix
// Given in Row-Major order
func (this *Mat4) Dump() (mat [16]float64) {
	mat = this.mat
	return
}

// Retrieve a 16 float array of all the values of the matrix
// Given in Col-Major order
func (this *Mat4) DumpOpenGL() (mat [16]float64) {
	mat[0], mat[1], mat[2], mat[3] = this.Col(0)
	mat[4], mat[5], mat[6], mat[7] = this.Col(1)
	mat[8], mat[9], mat[10], mat[11] = this.Col(2)
	mat[12], mat[13], mat[14], mat[15] = this.Col(3)
	return
}

// Compare this matrix to the other.
// Return true if all elements between them are the same.
// Equality is measured using an epsilon (< 0.0000001)
func (this *Mat4) Eq(other *Mat4) bool {
	for k, _ := range this.mat {
		if closeEq(this.mat[k], other.mat[k], epsilon) == false {
			return false
		}
	}
	return true
}

// Retrieve the element at column/x and row/y
// 0 indexed
// Does not do any bounds checking
func (this *Mat4) Get(row, col int) float64 {
	return this.mat[row*mat4Dim+col]
}

// Set the value at the specified columnd and row.
// 0 indexed
// Does not do any bounds checking
func (this *Mat4) Set(row, col int, value float64) *Mat4 {
	this.mat[row*mat4Dim+col] = value
	return this
}

// Retrieve the element at the given index
// 0 indexed
func (this *Mat4) At(index int) float64 {
	return this.mat[index]
}

// Set the element of the matrix specified at the index to the given value
// 0 indexed
// Return a pointer to the 'this'
func (this *Mat4) SetAt(index int, value float64) *Mat4 {
	this.mat[index] = value
	return this
}

// Set the specified row of the matrix to the given x,y,z,w values
// 0 indexed
// Does not do bounds checking of the row
func (this *Mat4) SetRow(row int, x, y, z, w float64) *Mat4 {
	this.mat[row*mat4Dim] = x
	this.mat[row*mat4Dim+1] = y
	this.mat[row*mat4Dim+2] = z
	this.mat[row*mat4Dim+3] = w
	return this
}

// Set the specified column of the matrix to the given x,y,z,w values
// 0 indexed
// Does not do bounds checking on the col
func (this *Mat4) SetCol(col int, x, y, z, w float64) *Mat4 {
	this.mat[mat4Dim*0+col] = x
	this.mat[mat4Dim*1+col] = y
	this.mat[mat4Dim*2+col] = z
	this.mat[mat4Dim*3+col] = w
	return this
}

// Retrieve the x,y,z,w elements from the specivied row
// 0 indexed
// Does not bounds check the row
func (this *Mat4) Row(row int) (x, y, z, w float64) {
	x = this.mat[row*mat4Dim]
	y = this.mat[row*mat4Dim+1]
	z = this.mat[row*mat4Dim+2]
	w = this.mat[row*mat4Dim+3]
	return
}

// Retrieve the x,y,z,w elements from the specified column
// 0 indexed
// Does not bounds check the column
func (this *Mat4) Col(col int) (x, y, z, w float64) {
	x = this.mat[mat4Dim*0+col]
	y = this.mat[mat4Dim*1+col]
	z = this.mat[mat4Dim*2+col]
	w = this.mat[mat4Dim*3+col]
	return
}

// Add in a constant value to all the terms fo the matrix
// Return a new matrix with the result
func (this *Mat4) AddScalar(val float64) *Mat4 {
	out := *this
	return out.AddInScalar(val)
}

// Add in a constant value to all the terms fo the matrix
// Returns a pointer to 'this'
func (this *Mat4) AddInScalar(val float64) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] += val
	}
	return this
}

// Subtract in a constant value to all the terms fo the matrix
// Return a new matrix with the result
func (this *Mat4) SubScalar(val float64) *Mat4 {
	out := *this
	return out.SubInScalar(val)
}

// Subtract in a constant value to all the terms fo the matrix
// Returns a pointer to 'this'
func (this *Mat4) SubInScalar(val float64) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] -= val
	}
	return this
}

// Multiplies in a constant value to all the terms fo the matrix
// Return a new matrix with the result
func (this *Mat4) MultScalar(val float64) *Mat4 {
	out := *this
	return out.MultInScalar(val)
}

// Multiplies in a constant value to all the terms fo the matrix
// Returns a pointer to 'this'
func (this *Mat4) MultInScalar(val float64) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] *= val
	}
	return this
}

// Divides in a constant value to all the terms fo the matrix
// Does not check of a division by 0
// Return a new matrix with the result
func (this *Mat4) DivScalar(val float64) *Mat4 {
	out := *this
	return out.DivInScalar(val)
}

// Divides in a constant value to all the terms fo the matrix
// Does not check of a division by 0
// Returns a pointer to 'this'
func (this *Mat4) DivInScalar(val float64) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] /= val
	}
	return this
}

// Add the 'other' matrix to 'this' and store the result in 'this'
// Returns a pointer to 'this'
func (this *Mat4) Add(other *Mat4) *Mat4 {
	out := *this
	return out.AddIn(other)
}

// Add the 'other' matrix to 'this' and store the result in 'this'
// Returns a pointer to 'this'
func (this *Mat4) AddIn(other *Mat4) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] += other.mat[k]
	}
	return this
}

// Subtract the 'other' matrix to 'this' and store the result in 'this'
// Returns a pointer to 'this'
func (this *Mat4) Sub(other *Mat4) *Mat4 {
	out := *this
	return out.SubIn(other)
}

// Subtract the 'other' matrix to 'this' and store the result in 'this'
// Returns a pointer to 'this'
func (this *Mat4) SubIn(other *Mat4) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] -= other.mat[k]
	}
	return this
}

// Multiply 'this' matrix with the other matrix.
// Return a new matrix with the result
func (this *Mat4) Mult(other *Mat4) *Mat4 {
	out := *this
	out.MultIn(other)
	return &out
}

// Multiply 'this' matrix with the other matrix.
// Store the result into 'this'
// Returns a pointer to 'this'
func (this *Mat4) MultIn(o *Mat4) *Mat4 {
	// 0   1   2   3
	// 4   5   6   7
	// 8   9   10  11
	// 12  13  14  15
	m := *this
	this.mat[0] = m.mat[0]*o.mat[0] + m.mat[1]*o.mat[4] + m.mat[2]*o.mat[8] + m.mat[3]*o.mat[12]
	this.mat[1] = m.mat[0]*o.mat[1] + m.mat[1]*o.mat[5] + m.mat[2]*o.mat[9] + m.mat[3]*o.mat[13]
	this.mat[2] = m.mat[0]*o.mat[2] + m.mat[1]*o.mat[6] + m.mat[2]*o.mat[10] + m.mat[3]*o.mat[14]
	this.mat[3] = m.mat[0]*o.mat[3] + m.mat[1]*o.mat[7] + m.mat[2]*o.mat[11] + m.mat[3]*o.mat[15]

	this.mat[4] = m.mat[4]*o.mat[0] + m.mat[5]*o.mat[4] + m.mat[6]*o.mat[8] + m.mat[7]*o.mat[12]
	this.mat[5] = m.mat[4]*o.mat[1] + m.mat[5]*o.mat[5] + m.mat[6]*o.mat[9] + m.mat[7]*o.mat[13]
	this.mat[6] = m.mat[4]*o.mat[2] + m.mat[5]*o.mat[6] + m.mat[6]*o.mat[10] + m.mat[7]*o.mat[14]
	this.mat[7] = m.mat[4]*o.mat[3] + m.mat[5]*o.mat[7] + m.mat[6]*o.mat[11] + m.mat[7]*o.mat[15]

	this.mat[8] = m.mat[8]*o.mat[0] + m.mat[9]*o.mat[4] + m.mat[10]*o.mat[8] + m.mat[11]*o.mat[12]
	this.mat[9] = m.mat[8]*o.mat[1] + m.mat[9]*o.mat[5] + m.mat[10]*o.mat[9] + m.mat[11]*o.mat[13]
	this.mat[10] = m.mat[8]*o.mat[2] + m.mat[9]*o.mat[6] + m.mat[10]*o.mat[10] + m.mat[11]*o.mat[14]
	this.mat[11] = m.mat[8]*o.mat[3] + m.mat[9]*o.mat[7] + m.mat[10]*o.mat[11] + m.mat[11]*o.mat[15]

	this.mat[12] = m.mat[12]*o.mat[0] + m.mat[13]*o.mat[4] + m.mat[14]*o.mat[8] + m.mat[15]*o.mat[12]
	this.mat[13] = m.mat[12]*o.mat[1] + m.mat[13]*o.mat[5] + m.mat[14]*o.mat[9] + m.mat[15]*o.mat[13]
	this.mat[14] = m.mat[12]*o.mat[2] + m.mat[13]*o.mat[6] + m.mat[14]*o.mat[10] + m.mat[15]*o.mat[14]
	this.mat[15] = m.mat[12]*o.mat[3] + m.mat[13]*o.mat[7] + m.mat[14]*o.mat[11] + m.mat[15]*o.mat[15]

	return this
}

// Returns a new matrix which is transpose to 'this'
func (this *Mat4) Transpose() *Mat4 {
	out := *this
	out.TransposeIn()
	return &out
}

// Take the transpose of this matrix.
// can definitely be way more efficient
func (this *Mat4) TransposeIn() *Mat4 {
	m00, m01, m02, m03 := this.Row(0)
	m10, m11, m12, m13 := this.Row(1)
	m20, m21, m22, m23 := this.Row(2)
	m30, m31, m32, m33 := this.Row(3)

	this.SetCol(0, m00, m01, m02, m03)
	this.SetCol(1, m10, m11, m12, m13)
	this.SetCol(2, m20, m21, m22, m23)
	this.SetCol(3, m30, m31, m32, m33)

	return this
}

// Calculate the determinant of a 2x2 matrix
// Values are givein in Row-Major order
func det2x2(x, y, z, w float64) float64 {
	return x*w - y*z
}

// Calculate the determinant of a 3x3 matrix
// Values are given in Row-Major order
func det3x3(a1, a2, a3, b1, b2, b3, c1, c2, c3 float64) float64 {
	// a1 a2 a3
	// b1 b2 b3
	// c1 c2 c3
	return (a1*det2x2(b2, b3, c2, c3) -
		b1*det2x2(a2, a3, c2, c3) +
		c1*det2x2(a2, a3, b2, b3))
}

// Get the determinant of the matrix
// Uses a straight-up Cramers-Rule implementation.
func (this *Mat4) Determinant() float64 {
	// 0   1   2   3
	// 4   5   6   7
	// 8   9   10  11
	// 12  13  14  15

	// Use Cramer's rule to calculate the determinant
	return (this.mat[0]*det3x3(this.mat[5], this.mat[6], this.mat[7],
		this.mat[9], this.mat[10], this.mat[11],
		this.mat[13], this.mat[14], this.mat[15]) -

		this.mat[4]*det3x3(this.mat[1], this.mat[2], this.mat[3],
			this.mat[9], this.mat[10], this.mat[11],
			this.mat[13], this.mat[14], this.mat[15]) +

		this.mat[8]*det3x3(this.mat[1], this.mat[2], this.mat[3],
			this.mat[5], this.mat[6], this.mat[7],
			this.mat[13], this.mat[14], this.mat[15]) -

		this.mat[12]*det3x3(this.mat[1], this.mat[2], this.mat[3],
			this.mat[5], this.mat[6], this.mat[7],
			this.mat[9], this.mat[10], this.mat[11]))
}

// Returns a new matrix which is the Adjoint matrix of this
func (this *Mat4) Adjoint() *Mat4 {
	a1, a2, a3, a4 := this.mat[0], this.mat[1], this.mat[2], this.mat[3]
	b1, b2, b3, b4 := this.mat[4], this.mat[5], this.mat[6], this.mat[7]
	c1, c2, c3, c4 := this.mat[8], this.mat[9], this.mat[10], this.mat[11]
	d1, d2, d3, d4 := this.mat[12], this.mat[13], this.mat[14], this.mat[15]

	// 0 1 2 3 			a1 a2 a3 a4
	// 4 5 6 7 			b1 b2 b3 b4
	// 8 9 10 11 		c1 c2 c3 c4
	// 12 13 14 15 		d1 d2 d3 d4
	m := Mat4{}
	m.mat[0] = det3x3(b2, b3, b4, c2, c3, c4, d2, d3, d4)
	m.mat[1] = -det3x3(b1, b3, b4, c1, c3, c4, d1, d3, d4)
	m.mat[2] = det3x3(b1, b2, b4, c1, c2, c4, d1, d2, d4)
	m.mat[3] = -det3x3(b1, b2, b3, c1, c2, c3, d1, d2, d3)

	m.mat[4] = -det3x3(a2, a3, a4, c2, c3, c4, d2, d3, d4)
	m.mat[5] = det3x3(a1, a3, a4, c1, c3, c4, d1, d3, d4)
	m.mat[6] = -det3x3(a1, a2, a4, c1, c2, c4, d1, d2, d4)
	m.mat[7] = det3x3(a1, a2, a3, c1, c2, c3, d1, d2, d3)

	m.mat[8] = det3x3(a2, a3, a4, b2, b3, b4, d2, d3, d4)
	m.mat[9] = -det3x3(a1, a3, a4, b1, b3, b4, d1, d3, d4)
	m.mat[10] = det3x3(a1, a2, a4, b1, b2, b4, d1, d2, d4)
	m.mat[11] = -det3x3(a1, a2, a3, b1, b2, b3, d1, d2, d3)

	m.mat[12] = -det3x3(a2, a3, a4, b2, b3, b4, c2, c3, c4)
	m.mat[13] = det3x3(a1, a3, a4, b1, b3, b4, c1, c3, c4)
	m.mat[14] = -det3x3(a1, a2, a4, b1, b2, b4, c1, c2, c4)
	m.mat[15] = det3x3(a1, a2, a3, b1, b2, b3, c1, c2, c3)

	return m.Transpose()
}

// Returns a new matrix which is the inverse matrix of this
// The bool flag is false if an inverse does not exist
// TODO: Needs further testing
// Try out with rotation matrices.
// The inverse of a valid rotation matrix should just be the transpose
func (this *Mat4) Inverse() *Mat4 {
	out := *this
	det := out.Determinant()
	return out.Adjoint().DivInScalar(det)
}

// Returns true if the inverse of this matrix exists
// False otherwise
// Internally it checks to see if the determinant is zero.
func (this *Mat4) HasInverse() bool {
	return !closeEq(this.Determinant(), 0, epsilon)
}

// Sets the matrix to the identity matrix
func (this *Mat4) ToIdentity() *Mat4 {
	this.mat = [16]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	return this
}

//==============================================================================

func (this *Mat4) IsIdentity() bool {
	iden := [16]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	for k, _ := range iden {
		if !closeEq(this.mat[k], iden[k], 0) {
			return false
		}
	}
	return true
}

// Check to see if the matrix is a valid rotation matrix
// The two properties it checks are
// 1) Determinant() == 1
// 2) m*m.Transpose  == Identity
func IsRotationMatrix(m *Mat4) bool {
	return closeEq(m.Determinant(), 1, epsilon) && m.Mult(m.Transpose()).IsIdentity()
}

// Implement the Stringer interface
// Prints out each row of the matrix on its own line
func (this *Mat4) String() string {
	return fmt.Sprintf("%f %f %f %f\n%f %f %f %f\n%f %f %f %f\n%f %f %f %f",
		this.mat[0], this.mat[1], this.mat[2], this.mat[3],
		this.mat[4], this.mat[5], this.mat[6], this.mat[7],
		this.mat[8], this.mat[9], this.mat[10], this.mat[11],
		this.mat[12], this.mat[13], this.mat[14], this.mat[15])
}

//==============================================================================

func (this *Mat4) ToTranslate(x, y, z float64) *Mat4 {
	this.ToIdentity()
	this.mat[3] = x
	this.mat[7] = y
	this.mat[11] = z
	return this
}

func (this *Mat4) ToScale(x, y, z float64) *Mat4 {
	this.ToIdentity()
	this.mat[0] = x
	this.mat[5] = y
	this.mat[10] = z
	return this
}

func (this *Mat4) ToSkew(x, y, z float64) *Mat4 {
	// 0   -z     y    0
	// z    0    -x    0
	// -y   x     0    0
	// 0    0     0    1

	this.ToIdentity()
	this.mat[0], this.mat[5], this.mat[10] = 0, 0, 0

	this.mat[6] = -x
	this.mat[9] = x

	this.mat[2] = y
	this.mat[8] = -y

	this.mat[1] = -z
	this.mat[4] = z

	return this
}

// Make this matrix into a rotation matrix about the x-axis
func (this *Mat4) ToRotateX(x float64) *Mat4 {
	this.ToIdentity()
	this.mat[5] = math.Cos(x)
	this.mat[6] = -math.Sin(x)
	this.mat[9] = math.Sin(x)
	this.mat[10] = math.Cos(x)
	return this
}

// Make this matrix into a rotation matrix about the y-axis
func (this *Mat4) ToRotateY(y float64) *Mat4 {
	this.ToIdentity()
	this.mat[0] = math.Cos(y)
	this.mat[2] = math.Sin(y)
	this.mat[8] = -math.Sin(y)
	this.mat[10] = math.Cos(y)
	return this
}

// Make this matrix into a rotation matrix about the z-axis
func (this *Mat4) ToRotateZ(z float64) *Mat4 {
	this.ToIdentity()
	this.mat[0] = math.Cos(z)
	this.mat[1] = -math.Sin(z)
	this.mat[4] = math.Sin(z)
	this.mat[5] = math.Cos(z)
	return this
}

// =============================================================================

// Return a rotation matrix which rotates a vector about the axis [x,y,z] with
// the given angle (radians).
// Set this matrix as a rotation matrix from the give angle(radians) and axis
func (this *Mat4) FromAxisAngle(angle, x, y, z float64) *Mat4 {
	//Reference http://en.wikipedia.org/wiki/Rotation_matrix
	c := math.Cos(angle)
	s := math.Sin(angle)
	t := (1 - c)

	return this.Load([16]float64{c + x*x*t, x*y*t - z*s, x*z*t + y*s, 0,
		y*x*t + z*s, c + y*y*t, y*z*t - x*s, 0,
		z*x*t - y*s, z*y*t + x*s, c + z*z*t, 0,
		0, 0, 0, 1})
}

func (this *Mat4) FromEuler(pitch, yaw, roll float64) *Mat4 {
	cx := math.Cos(pitch)
	sx := math.Sin(pitch)
	cy := math.Cos(yaw)
	sy := math.Sin(yaw)
	cz := math.Cos(roll)
	sz := math.Sin(roll)

	// This matrix was derived by multiplying each indiviudual rotation matrix
	// together into a single matrix.
	// note the matrices are applied in reverse order compared to the application
	// of the rotations.
	//   roll         yaw             pitch
	// | cz  -sz  0 | | cy   0   sy | | 1    0    0  |
	// | sz   cz  0 |x| 0    1   0  |x| 0    cx  -sx |
	// | 0    0   1 | | -sy  0   cy | | 0    sx   cx |

	// first row
	this.mat[0] = cz * cy
	this.mat[1] = cz*sy*sx - sz*cx
	this.mat[2] = sz*sx + cz*cx*sy
	this.mat[3] = 0
	// second row
	this.mat[4] = sz * cy
	this.mat[5] = cz*cx + sx*sy*sz
	this.mat[6] = sz*sy*cx - cz*sx
	this.mat[7] = 0
	// third row
	this.mat[8] = -sy
	this.mat[9] = sx * cy
	this.mat[10] = cy * cx
	this.mat[11] = 0

	this.mat[12] = 0
	this.mat[13] = 0
	this.mat[14] = 0
	this.mat[15] = 1
	return this
}

func (this *Mat4) AxisAngle() (angle, x, y, z float64) {
	// Reference
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToAngle/
	m00, m01, m02 := this.Get(0, 0), this.Get(0, 1), this.Get(0, 2)
	m10, m11, m12 := this.Get(1, 0), this.Get(1, 1), this.Get(1, 2)
	m20, m21, m22 := this.Get(2, 0), this.Get(2, 1), this.Get(2, 2)

	if closeEq(math.Abs(m01-m10), 0, epsilon) &&
		closeEq(math.Abs(m02-m20), 0, epsilon) &&
		closeEq(math.Abs(m12-m21), 0, epsilon) {
		// singularity check
		// Checking for cases in which the angle is either 0 or 180

		if this.IsIdentity() {
			// If the angle is 0, then the rotation matrix will be the identity matrix
			// A 0 angle means that there is an arbitrary axis.
			angle, x, y, z = 0, 1, 0, 0
			return
		}

		// Angle is 180, we need to find the axis it rotates around
		angle = math.Pi

		xx := (m00 + 1) / 2
		yy := (m11 + 1) / 2
		zz := (m22 + 1) / 2
		xy := (m01 + m10) / 4
		xz := (m02 + m20) / 4
		yz := (m12 + m21) / 4

		if (xx > yy) && (xx > zz) { // m[0][0] is the largest diagonal term
			if xx < epsilon {
				x = 0
				y = math.Sqrt(2) / 2
				z = math.Sqrt(2) / 2
			} else {
				x = math.Sqrt(xx)
				y = xy / x
				z = xz / x
			}
		} else if yy > zz { // m[1][1] is the largest diagonal term
			if yy < epsilon {
				x = math.Sqrt(2) / 2
				y = 0
				z = math.Sqrt(2) / 2
			} else {
				y = math.Sqrt(yy)
				x = xy / y
				z = yz / y
			}
		} else { // m[2][2] is the largest diagonal term so base result on this
			if zz < epsilon {
				x = math.Sqrt(2) / 2
				y = math.Sqrt(2) / 2
				z = 0
			} else {
				z = math.Sqrt(zz)
				x = xz / z
				y = yz / z
			}
		}
		return
	}

	// no singularity; therefore calculate as normal
	angle = math.Acos((m00 + m11 + m22 - 1) / 2)
	A := (m21 - m12)
	B := (m02 - m20)
	C := (m10 - m01)

	x = A / math.Sqrt(A*A+B*B+C*C)
	y = B / math.Sqrt(A*A+B*B+C*C)
	z = C / math.Sqrt(A*A+B*B+C*C)
	return
}

// Return the pitch,yaw and roll values for the given rotation matrix
// Assumption is that mat is a valid rotation matrix
// following the conventions of this package
// (x-y-z rotation order,row-major order, right-handed)
// The returned euler angle may not be the exact angle in which you supplied
// but they can be used to make an equilvalent rotation matrix.
func (this *Mat4) Euler() (pitch, yaw, roll float64) {
	// The method for calculating the euler angles from a rotation matrix
	// uses the method described in this document
	// http://staff.city.ac.uk/~sbbh653/publications/euler.pdf

	// The rotation matrix we are using will be of the following form
	// cos(x) is abbreviated as cx ( similarily sin(x) = sx)
	// This corresponds to the pitch => yaw => roll rotation matrix
	// cz*cy       cz*sy*sx - sz*cx         sz*sx + cz*cx*sy   | r11 r12 r13
	// sz*cy       cz*cx + sx*sy*sz         sz*sy*cx - cz*sx   | r21 r22 r23
	// -sy         sx*cy                    cx*cy              | r31 r32 r33

	// We want to determine the x,y,z angles
	// 1) Find the 'y' angle
	//      This is easily accomplished because term r31 is simply '-sin(y)'
	// 2) There are two possible angles for y because
	//      sin(y) == sin(pi - y)
	// 3) To find the value of x, we observe the following
	//      r32/r33 = tan(x)
	//      (sin(x)cos(y)) / (cos(x)cos(y))
	//      (sin(x)/cos(x)) == tan(x) by defn.
	// 4) We can also calculate x and z by.
	//      x = atan2(r32,r33) == atan2( (sin(x)cos(y)) / (cos(x)cos(y)) )
	// 		z = atan2(r21,r11) == atan2( (sin(z)cos(y)) / (cos(z)cos(y)) )

	var x, y, z float64
	r31 := this.Get(2, 0)
	if closeEq(r31, 1, epsilon) {
		// we are in gimbal lock
		z = 0
		y = -math.Pi / 2
		x = -z + math.Atan2(-this.Get(0, 1), -this.Get(0, 2))
	} else if closeEq(r31, -1, epsilon) {
		// we are in gimbal lock
		z = 0
		y = math.Pi / 2
		x = z + math.Atan2(this.Get(0, 1), this.Get(0, 2))
	} else {
		y = -math.Asin(r31)
		cos_y := math.Cos(y)
		x = math.Atan2(this.Get(2, 1)/cos_y, this.Get(2, 2)/cos_y)
		z = math.Atan2(this.Get(1, 0)/cos_y, this.Get(0, 0)/cos_y)

		// There are two alternative values for y,here is the second option
		// y2 := math.Pi - y
		// cos_y2 := math.Cos(y2)
		// x2 := math.Atan2(this.Get(2, 1)/cos_y2, this.Get(2, 2)/cos_y2)
		// z2 := math.Atan2(this.Get(1, 0)/cos_y2, this.Get(0, 0)/cos_y2)
	}

	pitch = x
	yaw = y
	roll = z
	return
}

// Creates a rotation matrix from the given quaternion
// returns this
func (this *Mat4) FromQuat(q *Quat) *Mat4 {
	*this = *q.Mat4()
	return this
}

// Returns the quaternion represented by this rotation matrix.
func (this *Mat4) Quat() *Quat {
	q := &Quat{}
	return q.FromMat4(this)
}

func (this *Mat4) MultVec3(v *Vec3) *Vec3 {
	// 0   1   2   3
	// 4   5   6   7
	// 8   9   10  11
	// 12  13  14  15
	return &Vec3{
		this.mat[0]*v.X + this.mat[1]*v.Y + this.mat[2]*v.Z + this.mat[3],
		this.mat[4]*v.X + this.mat[5]*v.Y + this.mat[6]*v.Z + this.mat[7],
		this.mat[8]*v.X + this.mat[9]*v.Y + this.mat[10]*v.Z + this.mat[11],
	}
}
