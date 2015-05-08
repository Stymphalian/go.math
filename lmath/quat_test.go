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
		get := test.orig.Equals(test.other)
		if get != test.want {
			t.Errorf("TestEquals %d", testIndex)
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
		if get.Equals(test.want) == false {
			t.Errorf("TestSet %d", testIndex)
		}
	}
}

func TestMultConst(t *testing.T) {
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
		get := test.orig.MultConst(test.scale)
		if get.Equals(test.want) == false {
			t.Errorf("TestMultConst %d", testIndex)
		}

		get2 := test.orig.MultInConst(test.scale)
		if get2.Equals(test.want) == false {
			t.Errorf("TestMultInConst %d", testIndex)
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
		if get.Equals(test.want) == false {
			t.Errorf("TestAdd %d", testIndex)
		}

		get2 := test.orig.AddIn(test.other)
		if get2.Equals(test.want) == false {
			t.Errorf("TestAddIn %d", testIndex)
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
		if get.Equals(test.want) == false {
			t.Errorf("TestMult %d %v", testIndex, get)
		}

		get2 := test.orig.MultIn(test.other)
		if get2.Equals(test.want) == false {
			t.Errorf("TestMultIn %d %v", testIndex, get)
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
		if get.Equals(test.want) == false {
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
		if closeEquals(get, test.want, epsilon) == false {
			t.Errorf("TestNorm %d", testIndex)
		}

		get2 := test.orig.SqrdNorm()
		if closeEquals(get2, test.want*test.want, epsilon) == false {
			t.Errorf("TestSqrdNorm %d", testIndex)
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
		if get.Equals(test.want) == false {
			t.Errorf("TestConjugate %d", testIndex)
		}

		get = test.orig.ConjugateIn()
		if get.Equals(test.want) == false {
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

	idenQuat := &Quat{1, 0, 0, 0}
	for testIndex, test := range cases {
		get := test.orig.Inverse().Mult(test.orig)
		if get.Equals(idenQuat) == false {
			t.Errorf("TestInverse %d", testIndex)
		}

		orig := *test.orig
		get = test.orig.InverseIn().MultIn(&orig)
		if get.Equals(idenQuat) == false {
			t.Errorf("TestInverseIn %d", testIndex)
		}
	}
}
