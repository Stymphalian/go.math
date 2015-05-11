package lmath

import (
	"fmt"
	"testing"
)

func TestEmptyGlmath(t *testing.T) {
	//fmt.Println("hello")
}

type v3 struct {
	X, Y, Z float64
}

func (this *v3) A() {
	fmt.Println("A")
}

func (this v3) B() {
	fmt.Println("B")
}

// func TestCrossPointer(t *testing.T){
//     v := &v3{1,2,3};
//     b := v3{1,2,3}
//     b.A()
//     b.B()
//     v.A()
//     v.B()
// }

func (this *v3) mult_1(other v3) *v3 {
	this.X *= other.X
	this.Y *= other.Y
	this.Z *= other.Z
	return this
}

func (this v3) mult0(other v3) v3 {
	this.mult_1(other)
	return this
}

func (this *v3) mult1(other *v3) *v3 {
	out := &v3{0, 0, 0}
	out.X = this.X * other.X
	out.Y = this.Y * other.Y
	out.Z = this.Z * other.Z
	return out
}

func (this *v3) mult2(other *v3) (out v3) {
	out.X = this.X * other.X
	out.Y = this.Y * other.Y
	out.Z = this.Z * other.Z
	return
}

func (this *v3) mult22(other v3) (out v3) {
	out.X = this.X * other.X
	out.Y = this.Y * other.Y
	out.Z = this.Z * other.Z
	return
}

func (this *v3) mult3(other v3) *v3 {
	out := &v3{0, 0, 0}
	out.X = this.X * other.X
	out.Y = this.Y * other.Y
	out.Z = this.Z * other.Z
	return out
}

func (this *v3) mult4(other v3) (out v3) {
	out.X = this.X * other.X
	out.Y = this.Y * other.Y
	out.Z = this.Z * other.Z
	return
}

func (this v3) mult5(other v3) *v3 {
	out := &v3{0, 0, 0}
	out.X = this.X * other.X
	out.Y = this.Y * other.Y
	out.Z = this.Z * other.Z
	return out
}

func (this v3) mult6(other v3) (out v3) {
	out.X = this.X * other.X
	out.Y = this.Y * other.Y
	out.Z = this.Z * other.Z
	return
}

func BenchmarkMult0(b *testing.B) {
	v := v3{1, 2, 3}
	o := v3{4, 5, 6}
	var rs v3
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = v.mult0(o)
		c += rs.X
	}
	fmt.Println(rs)
}

func BenchmarkMult1(b *testing.B) {
	v := &v3{1, 2, 3}
	o := &v3{4, 5, 6}
	var rs *v3
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = v.mult1(o)
		c += rs.X
	}
	fmt.Println(rs)
}

func BenchmarkMult2(b *testing.B) {
	v := &v3{1, 2, 3}
	o := &v3{4, 5, 6}
	var rs v3
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = v.mult2(o)
		c += rs.X
	}
	fmt.Println(rs)
}
func BenchmarkMult22(b *testing.B) {
	v := &v3{1, 2, 3}
	o := v3{4, 5, 6}
	var rs v3
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = v.mult22(o)
		c += rs.X
	}
	fmt.Println(rs)
}

func BenchmarkMult3(b *testing.B) {
	v := &v3{1, 2, 3}
	o := v3{4, 5, 6}
	var rs *v3
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = v.mult3(o)
		c += rs.X
	}
	fmt.Println(rs)
}

func BenchmarkMult4(b *testing.B) {
	v := &v3{1, 2, 3}
	o := v3{4, 5, 6}
	var rs v3
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = v.mult4(o)
		c += rs.X
	}
	fmt.Println(rs)
}

func BenchmarkMult5(b *testing.B) {
	v := v3{1, 2, 3}
	o := v3{4, 5, 6}
	var rs *v3
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = v.mult5(o)
		c += rs.X
	}
	fmt.Println(rs)
}

func BenchmarkMult6(b *testing.B) {
	v := v3{1, 2, 3}
	o := v3{4, 5, 6}
	var rs v3
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = v.mult6(o)
		c += rs.X
	}
	fmt.Println(rs)
}

func (this *Mat4) MultStuff1(o *Mat4) (m *Mat4) {
	// 0   1   2   3
	// 4   5   6   7
	// 8   9   10  11
	// 12  13  14  15
	m = &Mat4{}
	m.mat[0] = this.mat[0]*o.mat[0] + this.mat[1]*o.mat[4] + this.mat[2]*o.mat[8] + this.mat[3]*o.mat[12]
	m.mat[1] = this.mat[0]*o.mat[1] + this.mat[1]*o.mat[5] + this.mat[2]*o.mat[9] + this.mat[3]*o.mat[13]
	m.mat[2] = this.mat[0]*o.mat[2] + this.mat[1]*o.mat[6] + this.mat[2]*o.mat[10] + this.mat[3]*o.mat[14]
	m.mat[3] = this.mat[0]*o.mat[3] + this.mat[1]*o.mat[7] + this.mat[2]*o.mat[11] + this.mat[3]*o.mat[15]

	m.mat[4] = this.mat[4]*o.mat[0] + this.mat[5]*o.mat[4] + this.mat[6]*o.mat[8] + this.mat[7]*o.mat[12]
	m.mat[5] = this.mat[4]*o.mat[1] + this.mat[5]*o.mat[5] + this.mat[6]*o.mat[9] + this.mat[7]*o.mat[13]
	m.mat[6] = this.mat[4]*o.mat[2] + this.mat[5]*o.mat[6] + this.mat[6]*o.mat[10] + this.mat[7]*o.mat[14]
	m.mat[7] = this.mat[4]*o.mat[3] + this.mat[5]*o.mat[7] + this.mat[6]*o.mat[11] + this.mat[7]*o.mat[15]

	m.mat[8] = this.mat[8]*o.mat[0] + this.mat[9]*o.mat[4] + this.mat[10]*o.mat[8] + this.mat[11]*o.mat[12]
	m.mat[9] = this.mat[8]*o.mat[1] + this.mat[9]*o.mat[5] + this.mat[10]*o.mat[9] + this.mat[11]*o.mat[13]
	m.mat[10] = this.mat[8]*o.mat[2] + this.mat[9]*o.mat[6] + this.mat[10]*o.mat[10] + this.mat[11]*o.mat[14]
	m.mat[11] = this.mat[8]*o.mat[3] + this.mat[9]*o.mat[7] + this.mat[10]*o.mat[11] + this.mat[11]*o.mat[15]

	m.mat[12] = this.mat[12]*o.mat[0] + this.mat[13]*o.mat[4] + this.mat[14]*o.mat[8] + this.mat[15]*o.mat[12]
	m.mat[13] = this.mat[12]*o.mat[1] + this.mat[13]*o.mat[5] + this.mat[14]*o.mat[9] + this.mat[15]*o.mat[13]
	m.mat[14] = this.mat[12]*o.mat[2] + this.mat[13]*o.mat[6] + this.mat[14]*o.mat[10] + this.mat[15]*o.mat[14]
	m.mat[15] = this.mat[12]*o.mat[3] + this.mat[13]*o.mat[7] + this.mat[14]*o.mat[11] + this.mat[15]*o.mat[15]
	return
}

func (this *Mat4) MultStuff2(o *Mat4) (m Mat4) {
	// 0   1   2   3
	// 4   5   6   7
	// 8   9   10  11
	// 12  13  14  15
	m.mat[0] = this.mat[0]*o.mat[0] + this.mat[1]*o.mat[4] + this.mat[2]*o.mat[8] + this.mat[3]*o.mat[12]
	m.mat[1] = this.mat[0]*o.mat[1] + this.mat[1]*o.mat[5] + this.mat[2]*o.mat[9] + this.mat[3]*o.mat[13]
	m.mat[2] = this.mat[0]*o.mat[2] + this.mat[1]*o.mat[6] + this.mat[2]*o.mat[10] + this.mat[3]*o.mat[14]
	m.mat[3] = this.mat[0]*o.mat[3] + this.mat[1]*o.mat[7] + this.mat[2]*o.mat[11] + this.mat[3]*o.mat[15]

	m.mat[4] = this.mat[4]*o.mat[0] + this.mat[5]*o.mat[4] + this.mat[6]*o.mat[8] + this.mat[7]*o.mat[12]
	m.mat[5] = this.mat[4]*o.mat[1] + this.mat[5]*o.mat[5] + this.mat[6]*o.mat[9] + this.mat[7]*o.mat[13]
	m.mat[6] = this.mat[4]*o.mat[2] + this.mat[5]*o.mat[6] + this.mat[6]*o.mat[10] + this.mat[7]*o.mat[14]
	m.mat[7] = this.mat[4]*o.mat[3] + this.mat[5]*o.mat[7] + this.mat[6]*o.mat[11] + this.mat[7]*o.mat[15]

	m.mat[8] = this.mat[8]*o.mat[0] + this.mat[9]*o.mat[4] + this.mat[10]*o.mat[8] + this.mat[11]*o.mat[12]
	m.mat[9] = this.mat[8]*o.mat[1] + this.mat[9]*o.mat[5] + this.mat[10]*o.mat[9] + this.mat[11]*o.mat[13]
	m.mat[10] = this.mat[8]*o.mat[2] + this.mat[9]*o.mat[6] + this.mat[10]*o.mat[10] + this.mat[11]*o.mat[14]
	m.mat[11] = this.mat[8]*o.mat[3] + this.mat[9]*o.mat[7] + this.mat[10]*o.mat[11] + this.mat[11]*o.mat[15]

	m.mat[12] = this.mat[12]*o.mat[0] + this.mat[13]*o.mat[4] + this.mat[14]*o.mat[8] + this.mat[15]*o.mat[12]
	m.mat[13] = this.mat[12]*o.mat[1] + this.mat[13]*o.mat[5] + this.mat[14]*o.mat[9] + this.mat[15]*o.mat[13]
	m.mat[14] = this.mat[12]*o.mat[2] + this.mat[13]*o.mat[6] + this.mat[14]*o.mat[10] + this.mat[15]*o.mat[14]
	m.mat[15] = this.mat[12]*o.mat[3] + this.mat[13]*o.mat[7] + this.mat[14]*o.mat[11] + this.mat[15]*o.mat[15]
	return
}

func (this Mat4) MultStuff3(o Mat4) (m Mat4) {
	// 0   1   2   3
	// 4   5   6   7
	// 8   9   10  11
	// 12  13  14  15
	m.mat[0] = this.mat[0]*o.mat[0] + this.mat[1]*o.mat[4] + this.mat[2]*o.mat[8] + this.mat[3]*o.mat[12]
	m.mat[1] = this.mat[0]*o.mat[1] + this.mat[1]*o.mat[5] + this.mat[2]*o.mat[9] + this.mat[3]*o.mat[13]
	m.mat[2] = this.mat[0]*o.mat[2] + this.mat[1]*o.mat[6] + this.mat[2]*o.mat[10] + this.mat[3]*o.mat[14]
	m.mat[3] = this.mat[0]*o.mat[3] + this.mat[1]*o.mat[7] + this.mat[2]*o.mat[11] + this.mat[3]*o.mat[15]

	m.mat[4] = this.mat[4]*o.mat[0] + this.mat[5]*o.mat[4] + this.mat[6]*o.mat[8] + this.mat[7]*o.mat[12]
	m.mat[5] = this.mat[4]*o.mat[1] + this.mat[5]*o.mat[5] + this.mat[6]*o.mat[9] + this.mat[7]*o.mat[13]
	m.mat[6] = this.mat[4]*o.mat[2] + this.mat[5]*o.mat[6] + this.mat[6]*o.mat[10] + this.mat[7]*o.mat[14]
	m.mat[7] = this.mat[4]*o.mat[3] + this.mat[5]*o.mat[7] + this.mat[6]*o.mat[11] + this.mat[7]*o.mat[15]

	m.mat[8] = this.mat[8]*o.mat[0] + this.mat[9]*o.mat[4] + this.mat[10]*o.mat[8] + this.mat[11]*o.mat[12]
	m.mat[9] = this.mat[8]*o.mat[1] + this.mat[9]*o.mat[5] + this.mat[10]*o.mat[9] + this.mat[11]*o.mat[13]
	m.mat[10] = this.mat[8]*o.mat[2] + this.mat[9]*o.mat[6] + this.mat[10]*o.mat[10] + this.mat[11]*o.mat[14]
	m.mat[11] = this.mat[8]*o.mat[3] + this.mat[9]*o.mat[7] + this.mat[10]*o.mat[11] + this.mat[11]*o.mat[15]

	m.mat[12] = this.mat[12]*o.mat[0] + this.mat[13]*o.mat[4] + this.mat[14]*o.mat[8] + this.mat[15]*o.mat[12]
	m.mat[13] = this.mat[12]*o.mat[1] + this.mat[13]*o.mat[5] + this.mat[14]*o.mat[9] + this.mat[15]*o.mat[13]
	m.mat[14] = this.mat[12]*o.mat[2] + this.mat[13]*o.mat[6] + this.mat[14]*o.mat[10] + this.mat[15]*o.mat[14]
	m.mat[15] = this.mat[12]*o.mat[3] + this.mat[13]*o.mat[7] + this.mat[14]*o.mat[11] + this.mat[15]*o.mat[15]
	return
}

func BenchmarkMultStuff1(b *testing.B) {
	m1 := &Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}
	m2 := &Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}
	var rs *Mat4
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = m1.MultStuff1(m2)
		c += rs.At(5)
	}
	fmt.Println(c)
}

func BenchmarkMultStuff2(b *testing.B) {
	m1 := &Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}
	m2 := &Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}
	var rs Mat4
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = m1.MultStuff2(m2)
		c += rs.At(5)
	}
	fmt.Println(c)
}

func BenchmarkMultStuff3(b *testing.B) {
	m1 := &Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}
	m2 := Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}
	var rs Mat4
	c := 0.0
	for i := 0; i < b.N; i++ {
		rs = m1.MultStuff3(m2)
		c += rs.At(5)
	}
	fmt.Println(c)
}
