package matrix

import (
	"fmt"
)

const (
	mat4Dim = 4
)

type Mat4 struct {
	mat [16]float64
}

// Row-Oder
// 0   1   2   3
// 4   5   6   7
// 8   9   10  11
// 12  13  14  15
func NewMat4(
	v11, v12, v13, v14,
	v21, v22, v23, v24,
	v31, v32, v33, v34,
	v41, v42, v43, v44 float64) *Mat4 {

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

func (this *Mat4) Equals(other *Mat4) bool {
	for k, _ := range this.mat {
		if closeEquals(this.mat[k], other.mat[k], epsilon) == false {
			return false
		}
	}
	return true
}

func (this *Mat4) Get(col, row int) float64 {
	return this.mat[row*mat4Dim+col]
}
func (this *Mat4) Set(col, row int, value float64) *Mat4 {
	this.mat[row*mat4Dim+col] = value
	return this
}

func (this *Mat4) Load(mat [16]float64) *Mat4 {
	this.mat = mat
	return this
}

func (this *Mat4) Dump() (mat [16]float64) {
	mat = this.mat
	return
}
func (this *Mat4) DumpOpenGL() (mat [16]float64) {
	mat[0], mat[1], mat[2], mat[3] = this.GetCol(0)
	mat[4], mat[5], mat[6], mat[7] = this.GetCol(1)
	mat[8], mat[9], mat[10], mat[11] = this.GetCol(2)
	mat[12], mat[13], mat[14], mat[15] = this.GetCol(3)
	return
}

func (this *Mat4) SetRow(row int, x, y, z, w float64) *Mat4 {
	this.mat[row*mat4Dim] = x
	this.mat[row*mat4Dim+1] = y
	this.mat[row*mat4Dim+2] = z
	this.mat[row*mat4Dim+3] = w
	return this
}
func (this *Mat4) SetCol(col int, x, y, z, w float64) *Mat4 {
	this.mat[mat4Dim*0+col] = x
	this.mat[mat4Dim*1+col] = y
	this.mat[mat4Dim*2+col] = z
	this.mat[mat4Dim*3+col] = w
	return this
}

func (this *Mat4) GetRow(row int) (x, y, z, w float64) {
	x = this.mat[row*mat4Dim]
	y = this.mat[row*mat4Dim+1]
	z = this.mat[row*mat4Dim+2]
	w = this.mat[row*mat4Dim+3]
	return
}
func (this *Mat4) GetCol(col int) (x, y, z, w float64) {
	x = this.mat[mat4Dim*0+col]
	y = this.mat[mat4Dim*1+col]
	z = this.mat[mat4Dim*2+col]
	w = this.mat[mat4Dim*3+col]
	return
}

func (this *Mat4) AddInConst(val float64) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] += val
	}
	return this
}
func (this *Mat4) SubInConst(val float64) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] -= val
	}
	return this
}
func (this *Mat4) MultInConst(val float64) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] *= val
	}
	return this
}
func (this *Mat4) DivInConst(val float64) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] /= val
	}
	return this
}
func (this *Mat4) AddIn(other *Mat4) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] += other.mat[k]
	}
	return this
}
func (this *Mat4) SubIn(other *Mat4) *Mat4 {
	for k, _ := range this.mat {
		this.mat[k] -= other.mat[k]
	}
	return this
}

// 0   1   2   3
// 4   5   6   7
// 8   9   10  11
// 12  13  14  15
func (this *Mat4) MultIn(o *Mat4) *Mat4 {
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

func (this *Mat4) Transpose() *Mat4 {
	var out Mat4
	var x, y, z, w float64
	x, y, z, w = this.GetRow(0)
	out.SetCol(0, x, y, z, w)

	x, y, z, w = this.GetRow(1)
	out.SetCol(1, x, y, z, w)

	x, y, z, w = this.GetRow(2)
	out.SetCol(2, x, y, z, w)

	x, y, z, w = this.GetRow(3)
	out.SetCol(3, x, y, z, w)

	return &out
}

func det2x2(x, y, z, w float64) float64 {
	return x*w - y*z
}

// a1 a2 a3
// b1 b2 b3
// c1 c2 c3
func det3x3(a1, a2, a3, b1, b2, b3, c1, c2, c3 float64) float64 {
	return (a1*det2x2(b2, b3, c2, c3) -
		b1*det2x2(a2, a3, c2, c3) +
		c1*det2x2(a2, a3, b2, b3))
}

// 0   1   2   3
// 4   5   6   7
// 8   9   10  11
// 12  13  14  15
func (this *Mat4) Determinant() float64 {
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

func (this *Mat4) Adjoint() *Mat4 {
	a1, a2, a3, a4 := this.mat[0], this.mat[1], this.mat[2], this.mat[3]
	b1, b2, b3, b4 := this.mat[4], this.mat[5], this.mat[6], this.mat[7]
	c1, c2, c3, c4 := this.mat[8], this.mat[9], this.mat[10], this.mat[11]
	d1, d2, d3, d4 := this.mat[12], this.mat[13], this.mat[14], this.mat[15]

	m := Mat4{}

	// 0 1 2 3 			a1 a2 a3 a4
	// 4 5 6 7 			b1 b2 b3 b4
	// 8 9 10 11 		c1 c2 c3 c4
	// 12 13 14 15 		d1 d2 d3 d4
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

func (this *Mat4) Inverse() (*Mat4, bool) {
	out := *this

	det := out.Determinant()
	if closeEquals(det, 0, epsilon) {
		return nil, false
	}

	adj := out.Adjoint()
	return adj.DivInConst(det), true
}

func (this *Mat4) ToIdentity() *Mat4 {
	this.mat = [16]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	return this
}

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


// 0   -z     y    0
// z    0    -x    0
// -y   x     0    0
// 0    0     0    1
func (this *Mat4) ToSkew(x, y, z float64) *Mat4 {
	this.ToIdentity()
	this.mat[0],this.mat[5],this.mat[10] = 0,0,0

	this.mat[6] = -x
	this.mat[9] = x

	this.mat[2] = y
	this.mat[8] = -y

	this.mat[1] = -z
	this.mat[4] = z

	return this
}

// rotations can be done in a variety of ways.
// We leave that to the rotation package.
// angles is rad
func (this *Mat4) ToRotateAll(x, y, z float64) *Mat4 {
	return this
}
func (this *Mat4) ToRotateAxis(delta float64, axis *Vec3) *Mat4 {
	return this
}


func (this *Mat4) String() string {
	return fmt.Sprintf("%f %f %f %f\n%f %f %f %f\n%f %f %f %f\n%f %f %f %f",
		this.mat[0], this.mat[1], this.mat[2], this.mat[3],
		this.mat[4], this.mat[5], this.mat[6], this.mat[7],
		this.mat[8], this.mat[9], this.mat[10], this.mat[11],
		this.mat[12], this.mat[13], this.mat[14], this.mat[15])
}
