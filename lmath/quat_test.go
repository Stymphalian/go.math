package lmath

import (
	"math"
	"testing"
)

func TestEquals(t *testing.T) {
	cases := []struct {
		orig, other *Quat
		want        bool
	}{
		{&Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}, true},
		{&Quat{1, 0, 0, 0}, &Quat{0, 0, 0, 0}, false},
		{&Quat{1, 2, 3, 4}, &Quat{1.000, 2.000, 3.00, 4.0}, true},
		{&Quat{-1, -2, 3, 4}, &Quat{-1, -2, 3, 4}, true},
		{&Quat{-1, 2, 3, 4}, &Quat{-1, -2, 3, 4}, false},
	}

	for testIndex, test := range cases {
		get := test.orig.Eq(test.other)
		if get != test.want {
			t.Errorf("TestEq %d", testIndex)
		}
	}
}

func TestSet(t *testing.T) {
	cases := []struct {
		x, y, z, w float64
		want       *Quat
	}{
		{0, 0, 0, 0, &Quat{0, 0, 0, 0}},
		{1, 2, 3, 4, &Quat{1, 2, 3, 4}},
		{0.1, 0.2, 0.3, 0.4, &Quat{0.1, 0.2, 0.3, 0.4}},
		{-0.1, 0.2, -0.3, 0.4, &Quat{-0.1, 0.2, -0.3, 0.4}},
	}

	for testIndex, test := range cases {
		orig := &Quat{-1, -1, -1, -1}
		get := orig.Set(test.x, test.y, test.z, test.w)
		if get.Eq(test.want) == false {
			t.Errorf("TestSet %d", testIndex)
		}
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		orig, other, want *Quat
	}{
		{&Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}},
		{&Quat{0, 0, 0, 0}, &Quat{1, 2, 3, 4}, &Quat{1, 2, 3, 4}},
		{&Quat{1, 0, 0, 0}, &Quat{0, 1, 0, 0}, &Quat{1, 1, 0, 0}},
		{&Quat{0, 1, 0, 0}, &Quat{1, 0, 0, 0}, &Quat{1, 1, 0, 0}},
		{&Quat{1, 2, 3, 4}, &Quat{5, 6, 7, 8}, &Quat{6, 8, 10, 12}},
		{&Quat{1, -2, 3, -4}, &Quat{1, -2, 3, -4}, &Quat{2, -4, 6, -8}},
		{&Quat{1, 2, 3, 4}, &Quat{1, -2, 3, -4}, &Quat{2, 0, 6, 0}},
	}

	for testIndex, test := range cases {
		get := test.orig.Add(test.other)
		if get.Eq(test.want) == false {
			t.Errorf("TestAdd %d", testIndex)
		}

		get = test.orig.AddIn(test.other)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestAddIn %d", testIndex)
		}
	}
}

func TestSub(t *testing.T) {
	cases := []struct {
		orig, other, want *Quat
	}{
		{&Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}},
		{&Quat{0, 0, 0, 0}, &Quat{1, 2, 3, 4}, &Quat{-1, -2, -3, -4}},
		{&Quat{1, 0, 0, 0}, &Quat{0, 1, 0, 0}, &Quat{1, -1, 0, 0}},
		{&Quat{0, 1, 0, 0}, &Quat{1, 0, 0, 0}, &Quat{-1, 1, 0, 0}},
		{&Quat{1, 2, 3, 4}, &Quat{5, 6, 7, 8}, &Quat{-4,-4,-4,-4}},
		{&Quat{1, -2, 3, -4}, &Quat{1, -2, 3, -4}, &Quat{0,0,0,0}},
		{&Quat{1, 2, 3, 4}, &Quat{1, -2, 3, -4}, &Quat{0,4,0,8}},
	}

	for testIndex, test := range cases {
		get := test.orig.Sub(test.other)
		if get.Eq(test.want) == false {
			t.Errorf("TestSub %d", testIndex)
		}

		get = test.orig.SubIn(test.other)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestSubIn %d", testIndex)
		}
	}
}

func TestMult(t *testing.T) {
	cases := []struct {
		orig, other, want *Quat
	}{
		{&Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}},
		{&Quat{0, 0, 0, 0}, &Quat{1, 2, 3, 4}, &Quat{0, 0, 0, 0}},
		{&Quat{1, 0, 0, 0}, &Quat{0, 1, 0, 0}, &Quat{0, 1, 0, 0}},
		{&Quat{0, 1, 0, 0}, &Quat{1, 0, 0, 0}, &Quat{0, 1, 0, 0}},
		{&Quat{1, 2, 3, 4}, &Quat{5, 6, 7, 8}, &Quat{-60, 12, 30, 24}},
		{&Quat{1, -2, 3, -4}, &Quat{1, -2, 3, -4}, &Quat{-28, -4, 6, -8}},
	}

	for testIndex, test := range cases {
		get := test.orig.Mult(test.other)
		if get.Eq(test.want) == false {
			t.Errorf("TestMult %d %v", testIndex, get)
		}

		get = test.orig.MultIn(test.other)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestMultIn %d %v", testIndex, get)
		}
	}
}

func TestAddScalar(t *testing.T) {
	cases := []struct {
		orig, want *Quat
		scale      float64
	}{
		{&Quat{0, 0, 0, 0}, &Quat{2, 2, 2, 2}, 2},
		{&Quat{0, 0, 0, 0}, &Quat{-1,-1,-1,-1}, -1},
		{&Quat{1, 2, 3, 4}, &Quat{0,1,2,3}, -1},
		{&Quat{1, 2, 3, 4}, &Quat{3,4,5,6}, 2},
		{&Quat{1, 2, 3, 4}, &Quat{1,2,3,4}, 0},
	}

	for testIndex, test := range cases {
		get := test.orig.AddScalar(test.scale)
		if get.Eq(test.want) == false {
			t.Errorf("TestAddScalar %d", testIndex)
		}

		get = test.orig.AddInScalar(test.scale)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestAddInScalar %d", testIndex)
		}
	}
}

func TestSubScalar(t *testing.T) {
	cases := []struct {
		orig, want *Quat
		scale      float64
	}{
		{&Quat{0, 0, 0, 0}, &Quat{-2,-2,-2,-2}, 2},
		{&Quat{0, 0, 0, 0}, &Quat{1,1,1,1}, -1},
		{&Quat{1, 2, 3, 4}, &Quat{2,3,4,5}, -1},
		{&Quat{1, 2, 3, 4}, &Quat{-1,0,1,2}, 2},
		{&Quat{1, 2, 3, 4}, &Quat{1,2,3,4}, 0},
	}

	for testIndex, test := range cases {
		get := test.orig.SubScalar(test.scale)
		if get.Eq(test.want) == false {
			t.Errorf("TestSubScalar %d", testIndex)
		}

		get = test.orig.SubInScalar(test.scale)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestSubInScalar %d", testIndex)
		}
	}
}

func TestMultScalar(t *testing.T) {
	cases := []struct {
		orig, want *Quat
		scale      float64
	}{
		{&Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}, 2},
		{&Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}, -1},
		{&Quat{1, 2, 3, 4}, &Quat{-1, -2, -3, -4}, -1},
		{&Quat{1, 2, 3, 4}, &Quat{2, 4, 6, 8}, 2},
		{&Quat{1, 2, 3, 4}, &Quat{0, 0, 0, 0}, 0},
	}

	for testIndex, test := range cases {
		get := test.orig.MultScalar(test.scale)
		if get.Eq(test.want) == false {
			t.Errorf("TestMultScalar %d", testIndex)
		}

		get = test.orig.MultInScalar(test.scale)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestMultInScalar %d", testIndex)
		}
	}
}

func TestDivScalar(t *testing.T) {
	cases := []struct {
		orig, want *Quat
		scale      float64
	}{
		{&Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}, 2},
		{&Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}, -1},
		{&Quat{1, 2, 3, 4}, &Quat{-1, -2, -3, -4}, -1},
		{&Quat{1, 2, 3, 4}, &Quat{0.5,1,1.5,2}, 2},
	}

	for testIndex, test := range cases {
		get := test.orig.DivScalar(test.scale)
		if get.Eq(test.want) == false {
			t.Errorf("TestDivScalar %d", testIndex)
		}

		get = test.orig.DivInScalar(test.scale)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestDivInScalar %d", testIndex)
		}
	}
}


func TestToUnit(t *testing.T) {
	mag := math.Sqrt(30)
	cases := []struct {
		orig, want *Quat
	}{
		{&Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}},
		{&Quat{1, 0, 0, 0}, &Quat{1, 0, 0, 0}},
		{&Quat{0, 1, 0, 0}, &Quat{0, 1, 0, 0}},
		{&Quat{0, -1, 0, 0}, &Quat{0, -1, 0, 0}},
		{&Quat{1, 2, 3, 4}, &Quat{1 / mag, 2 / mag, 3 / mag, 4 / mag}},
		{&Quat{1 / mag, 2 / mag, 3 / mag, 4 / mag}, &Quat{1 / mag, 2 / mag, 3 / mag, 4 / mag}},
		{&Quat{1, -2, 3, -4}, &Quat{1 / mag, -2 / mag, 3 / mag, -4 / mag}},
	}

	for testIndex, test := range cases {
		get := test.orig.ToUnit()
		if get.Eq(test.want) == false {
			t.Errorf("TestToUnit %d", testIndex)
		}
	}
}

func TestNorm(t *testing.T) {
	mag := math.Sqrt(30)

	cases := []struct {
		orig *Quat
		want float64
	}{
		{&Quat{0, 0, 0, 0}, 0},
		{&Quat{1, 0, 0, 0}, 1},
		{&Quat{0, 1, 0, 0}, 1},
		{&Quat{0, -1, 0, 0}, 1},
		{&Quat{1, 2, 3, 4}, mag},
		{&Quat{1 / mag, 2 / mag, 3 / mag, 4 / mag}, 1},
		{&Quat{1, -2, 3, -4}, mag},
	}

	for testIndex, test := range cases {
		get := test.orig.Norm()
		if closeEq(get, test.want, epsilon) == false {
			t.Errorf("TestNorm %d", testIndex)
		}

		get = test.orig.NormSq()
		if closeEq(get, test.want*test.want, epsilon) == false {
			t.Errorf("TestNormSq %d", testIndex)
		}
	}
}

func TestConjugate(t *testing.T) {
	mag := math.Sqrt(30)
	cases := []struct {
		orig, want *Quat
	}{
		{&Quat{0, 0, 0, 0}, &Quat{0, 0, 0, 0}},
		{&Quat{1, 0, 0, 0}, &Quat{1, 0, 0, 0}},
		{&Quat{1, 1, 0, 0}, &Quat{1, -1, 0, 0}},
		{&Quat{1, -1, 0, 0}, &Quat{1, 1, 0, 0}},
		{&Quat{1, 2, 3, 4}, &Quat{1, -2, -3, -4}},
		{&Quat{1 / mag, 2 / mag, 3 / mag, 4 / mag}, &Quat{1 / mag, -2 / mag, -3 / mag, -4 / mag}},
		{&Quat{1, -2, 3, -4}, &Quat{1, 2, -3, 4}},
	}

	for testIndex, test := range cases {
		get := test.orig.Conjugate()
		if get.Eq(test.want) == false {
			t.Errorf("TestConjugate %d", testIndex)
		}

		get = test.orig.ConjugateIn()
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestConjugateIn %d", testIndex)
		}

	}
}

func TestInverse(t *testing.T) {
	cases := []struct {
		orig *Quat
	}{
		{&Quat{1, 0, 0, 0}},
		{&Quat{1, 1, 0, 0}},
		{&Quat{1, -1, 0, 0}},
		{&Quat{1, 2, 3, 4}},
		{&Quat{1, -2, 3, -4}},
	}

	for testIndex, test := range cases {
		get := test.orig.Inverse().Mult(test.orig)
		if get.Eq(QuatIdentity) == false {
			t.Errorf("TestInverse %d", testIndex)
		}

		orig := *test.orig
		get = test.orig.InverseIn().MultIn(&orig)
		if get.Eq(QuatIdentity) == false {
			t.Errorf("TestInverseIn %d", testIndex)
		}
	}
}



// =============================================================================

func TestFromAxisAngle(t *testing.T) {
	cases := []struct {
		angle     float64
		axis      *Vec3
		start_vec *Vec3
		want      *Vec3
	}{
		//test basic rotations using a [1,0,0] vector
		{90, &Vec3{0, 1, 0}, &Vec3{1, 0, 0}, &Vec3{0, 0, -1}},
		{90, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{90, &Vec3{0, 0, 1}, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}},
		{-90, &Vec3{0, 1, 0}, &Vec3{1, 0, 0}, &Vec3{0, 0, 1}},
		{-90, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{-90, &Vec3{0, 0, 1}, &Vec3{1, 0, 0}, &Vec3{0, -1, 0}},
		{360, &Vec3{0, 0, 1}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{180, &Vec3{0, 0, 1}, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}},

		//test basic rotations using a [0,1,0] vector
		{90, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{90, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}, &Vec3{0, 0, 1}},
		{90, &Vec3{0, 0, 1}, &Vec3{0, 1, 0}, &Vec3{-1, 0, 0}},
		{-90, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{-90, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}, &Vec3{0, 0, -1}},
		{-90, &Vec3{0, 0, 1}, &Vec3{0, 1, 0}, &Vec3{1, 0, 0}},
		{360, &Vec3{0, 0, 1}, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{180, &Vec3{0, 0, 1}, &Vec3{0, 1, 0}, &Vec3{0, -1, 0}},

		// test negative axes
		{90, &Vec3{0, -1, 0}, &Vec3{1, 0, 0}, &Vec3{0, 0, 1}},
		{90, &Vec3{-1, 0, 0}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{90, &Vec3{0, 0, -1}, &Vec3{1, 0, 0}, &Vec3{0, -1, 0}},
		{-90, &Vec3{0, -1, 0}, &Vec3{1, 0, 0}, &Vec3{0, 0, -1}},
		{-90, &Vec3{-1, 0, 0}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{-90, &Vec3{0, 0, -1}, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}},
		{360, &Vec3{0, 0, -1}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{180, &Vec3{0, 0, -1}, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}},

		// test arbitraty axis
		{360, &Vec3{1, 1, 0}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{90, &Vec3{1, 1, 0}, &Vec3{1, 0, 0}, &Vec3{0.5, 0.5, -0.7071067811}},
		{45, &Vec3{1, 1, 0}, &Vec3{1, 0, 0}, &Vec3{0.85355339059, 0.1464466094067, -0.5}},
	}

	q := &Quat{}
	for testIndex, c := range cases {
		c.axis.NormalizeIn()
		q.FromAxisAngle(Radians(c.angle), c.axis.X, c.axis.Y, c.axis.Z)
		get := q.RotateVec3(c.start_vec)
		if get.Equals(c.want) == false {
			t.Errorf("TestFromAxisAngle %d \n %v\n%v\n\n", testIndex, q, get)
		}
	}
}

func TestFromEuler(t *testing.T) {
	common_cases := []struct {
		pitch, yaw, roll float64
		start_vec        *Vec3
		want             *Vec3
	}{
		//test basic rotations using a [0,1,0] vector
		// pitch,yaw,roll
		{0, 0, 90, &Vec3{0, 1, 0}, &Vec3{-1, 0, 0}},
		{0, 90, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{90, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 0, 1}},
		{0, 0, -90, &Vec3{0, 1, 0}, &Vec3{1, 0, 0}},
		{0, -90, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{-90, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 0, -1}},
		{0, 180, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}}, //6

		// test basic rotation using a [1,0,0] vector
		{0, 0, 90, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}},
		{0, 90, 0, &Vec3{1, 0, 0}, &Vec3{0, 0, -1}},
		{90, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 0, -90, &Vec3{1, 0, 0}, &Vec3{0, -1, 0}},
		{0, -90, 0, &Vec3{1, 0, 0}, &Vec3{0, 0, 1}},
		{-90, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 0, 180, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}}, //13

		// basic rotation using a non major axis vector
		{0, 0, 90, &Vec3{1, 1, 0}, &Vec3{-1, 1, 0}},
		{0, 90, 0, &Vec3{1, -1, 0}, &Vec3{0, -1, -1}},
		{90, 0, 0, &Vec3{-1, -1, 0}, &Vec3{-1, 0, -1}}, //16

		// two rotations
		{90, 0, 45, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, -math.Sqrt(2) / 2, 0}},
		{90, 45, 0, &Vec3{0, 0, 1}, &Vec3{0, -1, 0}},
		{45, 90, 0, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, -math.Sqrt(2) / 2, 0}},
		{45, 90, 90, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0}}, //20
	}

	q := &Quat{}
	for testIndex, c := range common_cases {
		q.FromEuler(Radians(c.pitch), Radians(c.yaw), Radians(c.roll))
		get := q.RotateVec3(c.start_vec)
		if get.Equals(c.want) == false {
			t.Errorf("TestFromEuler %d \n %v\n%v\n\n", testIndex, q, get)
		}
	}
}

func TestFromMat4(t *testing.T) {
	common_cases := []struct {
		pitch, yaw, roll float64
		start_vec        *Vec3
		want             *Vec3
	}{
		//test basic rotations using a [0,1,0] vector
		// pitch,yaw,roll
		{0, 0, 90, &Vec3{0, 1, 0}, &Vec3{-1, 0, 0}},
		{0, 90, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{90, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 0, 1}},
		{0, 0, -90, &Vec3{0, 1, 0}, &Vec3{1, 0, 0}},
		{0, -90, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{-90, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 0, -1}},
		{0, 180, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}}, //6

		// test basic rotation using a [1,0,0] vector
		{0, 0, 90, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}},
		{0, 90, 0, &Vec3{1, 0, 0}, &Vec3{0, 0, -1}},
		{90, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 0, -90, &Vec3{1, 0, 0}, &Vec3{0, -1, 0}},
		{0, -90, 0, &Vec3{1, 0, 0}, &Vec3{0, 0, 1}},
		{-90, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 0, 180, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}}, //13

		// basic rotation using a non major axis vector
		{0, 0, 90, &Vec3{1, 1, 0}, &Vec3{-1, 1, 0}},
		{0, 90, 0, &Vec3{1, -1, 0}, &Vec3{0, -1, -1}},
		{90, 0, 0, &Vec3{-1, -1, 0}, &Vec3{-1, 0, -1}}, //16

		// two rotations
		{90, 0, 45, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, -math.Sqrt(2) / 2, 0}},
		{90, 45, 0, &Vec3{0, 0, 1}, &Vec3{0, -1, 0}},
		{45, 90, 0, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, -math.Sqrt(2) / 2, 0}},
		{45, 90, 90, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0}}, //20
	}

	var m *Mat4
	q := &Quat{}
	for testIndex, c := range common_cases {
		m = EulerToMat4(Radians(c.pitch), Radians(c.yaw), Radians(c.roll))
		q.FromMat4(m)
		get := q.RotateVec3(c.start_vec)
		if get.Equals(c.want) == false {
			t.Errorf("TestFromMat4 %d \n %v\n%v\n\n", testIndex, q, get)
		}
	}
}

func TestAxisAngle(t *testing.T) {
	cases := []struct {
		angle, x, y, z float64
	}{
		//test basic rotations using a [1,0,0] vector
		{90, 1, 0, 0},
		{90, 0, 1, 0},
		{90, 0, 0, 1},
		{45, 1, 0, 0},
		{45, 0, 1, 0},
		{45, 0, 0, 1},
		{180, 1, 0, 0},
		{180, 0, 1, 0},
		{180, 0, 0, 1},
		{90, 1, 1, 0},
		{90, 1, 1, 0},
		{90, 0, -1, 1},
		{45, 1, 0, 1},
		{45, 0, 1, 0},
		{45, 1, 0, 1},
		{180, 1, -2, 0},
		{180, 0, 1, 20},
		{180, -4, 4, 1},
	}

	//var q *Quat
	q := &Quat{}
	for testIndex, c := range cases {
		v := &Vec3{c.x, c.y, c.z}
		v.NormalizeIn()
		q.FromAxisAngle(Radians(c.angle), v.X, v.Y, v.Z)
		get_angle, get_x, get_y, get_z := q.AxisAngle()
		if !closeEq(Degrees(get_angle), c.angle, epsilon) ||
			!closeEq(get_x, v.X, epsilon) ||
			!closeEq(get_y, v.Y, epsilon) ||
			!closeEq(get_z, v.Z, epsilon) {
			t.Errorf("TestQuatToAxisAngle %d %v %f %f %f %f\n%f %f %f %f\n",
				testIndex, v, Degrees(get_angle), get_x, get_y, get_z, c.angle, v.X, v.Y, v.Z)
		}
	}
}

func TestMat4(t *testing.T) {
	common_cases := []struct {
		pitch, yaw, roll float64
		start_vec        *Vec3
		want             *Vec3
	}{
		//test basic rotations using a [0,1,0] vector
		// pitch,yaw,roll
		{0, 0, 90, &Vec3{0, 1, 0}, &Vec3{-1, 0, 0}},
		{0, 90, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{90, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 0, 1}},
		{0, 0, -90, &Vec3{0, 1, 0}, &Vec3{1, 0, 0}},
		{0, -90, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{-90, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 0, -1}},
		{0, 180, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}}, //6

		// test basic rotation using a [1,0,0] vector
		{0, 0, 90, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}},
		{0, 90, 0, &Vec3{1, 0, 0}, &Vec3{0, 0, -1}},
		{90, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 0, -90, &Vec3{1, 0, 0}, &Vec3{0, -1, 0}},
		{0, -90, 0, &Vec3{1, 0, 0}, &Vec3{0, 0, 1}},
		{-90, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 0, 180, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}}, //13

		// basic rotation using a non major axis vector
		{0, 0, 90, &Vec3{1, 1, 0}, &Vec3{-1, 1, 0}},
		{0, 90, 0, &Vec3{1, -1, 0}, &Vec3{0, -1, -1}},
		{90, 0, 0, &Vec3{-1, -1, 0}, &Vec3{-1, 0, -1}}, //16

		// two rotations
		{90, 0, 45, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, -math.Sqrt(2) / 2, 0}},
		{90, 45, 0, &Vec3{0, 0, 1}, &Vec3{0, -1, 0}},
		{45, 90, 0, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, -math.Sqrt(2) / 2, 0}},
		{45, 90, 90, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0}}, //20
	}
	var m *Mat4
	q := &Quat{}
	for testIndex, c := range common_cases {
		q.FromEuler(Radians(c.pitch), Radians(c.yaw), Radians(c.roll))
		m = q.Mat4()

		get := MultMat4Vec3(m, c.start_vec)
		if get.Equals(c.want) == false {
			t.Errorf("TestMat4 %d \n%v\n%v\n\n", testIndex, m, get)
		}
	}
}

func TestQuatToEuler(t *testing.T) {
	common_cases2 := []rotation_test_struct{
		{180, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 180, 0, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}},
		{0, 0, 180, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}},//2
		{180, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, -1, 0}},
		{0, 180, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{0, 0, 180, &Vec3{0, 1, 0}, &Vec3{0, -1, 0}},//5
		{180, 0, 0, &Vec3{0, 0, 1}, &Vec3{0, 0, -1}},
		{0, 180, 0, &Vec3{0, 0, 1}, &Vec3{0, 0, -1}},
		{0, 0, 180, &Vec3{0, 0, 1}, &Vec3{0, 0, 1}}, //8

		{180, 0, 0, &Vec3{-1, 0, 0}, &Vec3{-1, 0, 0}},
		{0, 180, 0, &Vec3{-1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 0, 180, &Vec3{-1, 0, 0}, &Vec3{1, 0, 0}},//11
		{180, 0, 0, &Vec3{0, -1, 0}, &Vec3{0, 1, 0}},
		{0, 180, 0, &Vec3{0, -1, 0}, &Vec3{0, -1, 0}},
		{0, 0, 180, &Vec3{0, -1, 0}, &Vec3{0, 1, 0}},//14
		{180, 0, 0, &Vec3{0, 0, -1}, &Vec3{0, 0, 1}},
		{0, 180, 0, &Vec3{0, 0, -1}, &Vec3{0, 0, 1}},
		{0, 0, 180, &Vec3{0, 0, -1}, &Vec3{0, 0, -1}},//17

		{0, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{0, 0, 0, &Vec3{0, 0, 1}, &Vec3{0, 0, 1}},//2
	 	{45, 90, 90, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0}},

	 	//test basic rotations using a [0,1,0] vector
		// pitch,yaw,roll
		{0, 0, 90, &Vec3{0, 1, 0}, &Vec3{-1, 0, 0}},
		{0, 90, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{90, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 0, 1}},
		{0, 0, -90, &Vec3{0, 1, 0}, &Vec3{1, 0, 0}},
		{0, -90, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{-90, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 0, -1}},
		{0, 180, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}}, //6

		// test basic rotation using a [1,0,0] vector
		{0, 0, 90, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}},
		{0, 90, 0, &Vec3{1, 0, 0}, &Vec3{0, 0, -1}},
		{90, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 0, -90, &Vec3{1, 0, 0}, &Vec3{0, -1, 0}},
		{0, -90, 0, &Vec3{1, 0, 0}, &Vec3{0, 0, 1}},
		{-90, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{0, 0, 180, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}}, //13

		// basic rotation using a non major axis vector
		{0, 0, 90, &Vec3{1, 1, 0}, &Vec3{-1, 1, 0}},
		{0, 90, 0, &Vec3{1, -1, 0}, &Vec3{0, -1, -1}},
		{90, 0, 0, &Vec3{-1, -1, 0}, &Vec3{-1, 0, -1}}, //16

		// two rotations
		{90, 0, 45, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, -math.Sqrt(2) / 2, 0}},
		{90, 45, 0, &Vec3{0, 0, 1}, &Vec3{0, -1, 0}},
		{45, 90, 0, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, -math.Sqrt(2) / 2, 0}},
		{45, 90, 90, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0}}, //20
	}

	q := &Quat{}
	for testIndex, c := range common_cases2 {
		q.FromEuler(Radians(c.pitch), Radians(c.yaw), Radians(c.roll))
		pitch, yaw, roll := q.Euler()

		if !closeEq(yaw, Radians(c.yaw), epsilon) ||
			!closeEq(pitch, Radians(c.pitch), epsilon) ||
			!closeEq(roll, Radians(c.roll), epsilon) {
			t.Errorf("TestEuler %d %f %f %f ", testIndex, pitch, yaw, roll)
		}
	}
}
