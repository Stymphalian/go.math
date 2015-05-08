package lmath

import (
	// "fmt"
	"math"
	"testing"
)

func TestEqualVec3(t *testing.T) {
	var cases = []struct {
		orig, other *Vec3
		want        bool
	}{
		{&Vec3{0, 0, 0}, &Vec3{1, 2, 3}, false},
		{&Vec3{1, 2, 3}, &Vec3{0, 0, 0}, false},
		{&Vec3{1, 2, 3}, &Vec3{-1, -2, -3}, false},
		{&Vec3{0, 0, 0}, &Vec3{0, 0, 0}, true},
		{&Vec3{1.0, 2.0, 3.0}, &Vec3{1.0, 2.0, 3.0}, true},
	}

	for testIndex, test := range cases {
		get := test.orig.Equals(test.other)
		if get != test.want {
			t.Errorf("TestEqualVec3 %d", testIndex)
		}
	}
}

func TestAddVec3(t *testing.T) {

	var cases = []struct {
		orig, other, want *Vec3
	}{
		{&Vec3{0, 0, 0}, &Vec3{1, 2, 3}, &Vec3{1, 2, 3}},
		{&Vec3{1, 2, 3}, &Vec3{0, 0, 0}, &Vec3{1, 2, 3}},
		{&Vec3{1, 2, 3}, &Vec3{-1, -2, -3}, &Vec3{0, 0, 0}},
		{&Vec3{0, 0, 0}, &Vec3{0, 0, 0}, &Vec3{0, 0, 0}},
	}

	for testIndex, test := range cases {
		get := test.orig.Add(test.other)
		if get.Equals(test.want) == false {
			t.Errorf("TestAddVec3 %d", testIndex)
		}

		get2 := test.orig.AddIn(test.other)
		if get2.Equals(test.want) == false {
			t.Errorf("TestAddVec3 AddIn %d", testIndex)
		}
	}
}

func TestSubVec3(t *testing.T) {

	var cases = []struct {
		orig, other, want *Vec3
	}{
		{&Vec3{0, 0, 0}, &Vec3{1, 2, 3}, &Vec3{-1, -2, -3}},
		{&Vec3{1, 2, 3}, &Vec3{0, 0, 0}, &Vec3{1, 2, 3}},
		{&Vec3{1, 2, 3}, &Vec3{-1, -2, -3}, &Vec3{2, 4, 6}},
		{&Vec3{0, 0, 0}, &Vec3{0, 0, 0}, &Vec3{0, 0, 0}},
	}

	for testIndex, test := range cases {
		get := test.orig.Sub(test.other)
		if get.Equals(test.want) == false {
			t.Errorf("TestSubVec3 %d", testIndex)
		}

		get2 := test.orig.SubIn(test.other)
		if get2.Equals(test.want) == false {
			t.Errorf("TestSubVec3 SubIn %d", testIndex)
		}
	}
}

func TestMultVec3(t *testing.T) {
	var cases = []struct {
		orig  *Vec3
		scale float64
		want  *Vec3
	}{
		{&Vec3{0, 0, 0}, 2.0, &Vec3{0, 0, 0}},
		{&Vec3{0, 0, 0}, -2, &Vec3{0, 0, 0}},
		{&Vec3{1, 2, 3}, 2, &Vec3{2, 4, 6}},
		{&Vec3{1, 2, 3}, 0.5, &Vec3{0.5, 1, 1.5}},
		{&Vec3{1, 2, 3}, -1, &Vec3{-1, -2, -3}},
		{&Vec3{1, 2, 3}, -0.5, &Vec3{-0.5, -1, -1.5}},
		{&Vec3{1, 2, 3}, 0, &Vec3{0, 0, 0}},
	}

	for testIndex, test := range cases {
		get := test.orig.Mult(test.scale)
		if get.Equals(test.want) == false {
			t.Errorf("TestMultVec3 %d", testIndex)
		}

		get2 := test.orig.MultIn(test.scale)
		if get2.Equals(test.want) == false {
			t.Errorf("TestMultVec3 MultIn %d", testIndex)
		}
	}
}

func TestDivVec3(t *testing.T) {
	var cases = []struct {
		orig  *Vec3
		scale float64
		want  *Vec3
	}{
		{&Vec3{0, 0, 0}, 2.0, &Vec3{0, 0, 0}},
		{&Vec3{0, 0, 0}, -2.0, &Vec3{0, 0, 0}},
		{&Vec3{1, 2, 3}, 2, &Vec3{0.5, 1, 1.5}},
		{&Vec3{1, 2, 3}, 0.5, &Vec3{2, 4, 6}},
		{&Vec3{1, 2, 3}, -1, &Vec3{-1, -2, -3}},
		{&Vec3{1, 2, 3}, -0.5, &Vec3{-2, -4, -6}},
	}

	for testIndex, test := range cases {
		get := test.orig.Div(test.scale)
		if get.Equals(test.want) == false {
			t.Errorf("TestDivVec3 %d", testIndex)
		}

		get2 := test.orig.DivIn(test.scale)
		if get2.Equals(test.want) == false {
			t.Errorf("TestDivInVec3 %d", testIndex)
		}
	}
}

func TestOuterVec3(t *testing.T) {
	var cases = []struct {
		orig, other, want *Vec3
	}{
		{&Vec3{0, 0, 0}, &Vec3{1, 2, 3}, &Vec3{0, 0, 0}},
		{&Vec3{1, 2, 3}, &Vec3{0, 0, 0}, &Vec3{0, 0, 0}},
		{&Vec3{1, 2, 3}, &Vec3{-1, -2, -3}, &Vec3{-1, -4, -9}},
		{&Vec3{1, 2, 3}, &Vec3{1, 2, 3}, &Vec3{1, 4, 9}},
		{&Vec3{0, 0, 0}, &Vec3{0, 0, 0}, &Vec3{0, 0, 0}},
	}

	for testIndex, test := range cases {
		get := test.orig.Outer(test.other)
		if get.Equals(test.want) == false {
			t.Errorf("TestOuterVec3 %d", testIndex)
		}

		get2 := test.orig.OuterIn(test.other)
		if get2.Equals(test.want) == false {
			t.Errorf("TestOuterInVec3 %d", testIndex)
		}
	}
}

func TestDotVec3(t *testing.T) {
	var cases = []struct {
		orig  *Vec3
		other *Vec3
		want  float64
	}{
		{&Vec3{0, 0, 0}, &Vec3{1, 2, 3}, 0},
		{&Vec3{1, 2, 3}, &Vec3{0, 0, 0}, 0},
		{&Vec3{1, 2, 3}, &Vec3{-1, -2, -3}, -14},
		{&Vec3{1, 2, 3}, &Vec3{1, 2, 3}, 14},
		{&Vec3{0, 0, 0}, &Vec3{0, 0, 0}, 0},
	}

	for testIndex, test := range cases {
		get := test.orig.Dot(test.other)
		if get != test.want {
			t.Errorf("TestDotVec3 %d", testIndex)
		}
	}
}

func TestCrossVec3(t *testing.T) {
	var cases = []struct {
		orig, other, want *Vec3
	}{
		{&Vec3{0, 0, 0}, &Vec3{1, 2, 3}, &Vec3{0, 0, 0}},
		{&Vec3{1, 2, 3}, &Vec3{0, 0, 0}, &Vec3{0, 0, 0}},
		{&Vec3{1, 0, 0}, &Vec3{0, 1, 0}, &Vec3{0, 0, 1}},
		{&Vec3{1, 2, 3}, &Vec3{-1, -2, -3}, &Vec3{0, 0, 0}},
		{&Vec3{math.Sqrt(2), math.Sqrt(2), 0}, &Vec3{0, 0, -1}, &Vec3{-math.Sqrt(2), math.Sqrt(2), 0}},
		{&Vec3{0, 0, -1}, &Vec3{math.Sqrt(2), math.Sqrt(2), 0}, &Vec3{math.Sqrt(2), -math.Sqrt(2), 0}},
	}

	for testIndex, test := range cases {
		get := test.orig.Cross(test.other)
		if get.Equals(test.want) == false {
			t.Errorf("TestCrossVec3 %d", testIndex)
		}

		get2 := test.orig.CrossIn(test.other)
		if get2.Equals(test.want) == false {
			t.Errorf("TestCrossVec3 %d", testIndex)
		}
	}
}

func TestLengthVec3(t *testing.T) {
	var cases = []struct {
		orig *Vec3
		want float64
	}{
		{&Vec3{0, 0, 0}, 0},
		{&Vec3{1, 2, 3}, math.Sqrt(14)},
		{&Vec3{1, 0, 0}, 1},
		{&Vec3{1 / math.Sqrt(14), 2 / math.Sqrt(14), 3 / math.Sqrt(14)}, 1},
	}

	for testIndex, test := range cases {
		get := test.orig.Length()
		if get != test.want {
			t.Errorf("TestLengthVec3 %d", testIndex)
		}
	}
}

func TestNormalizeVec3(t *testing.T) {
	var cases = []struct {
		orig, want *Vec3
	}{
		{&Vec3{1, 2, 3}, &Vec3{1 / math.Sqrt(14), 2 / math.Sqrt(14), 3 / math.Sqrt(14)}},
		{&Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{&Vec3{1 / math.Sqrt(14), 2 / math.Sqrt(14), 3 / math.Sqrt(14)}, &Vec3{1 / math.Sqrt(14), 2 / math.Sqrt(14), 3 / math.Sqrt(14)}},
	}

	for testIndex, test := range cases {
		get := test.orig.Normalize()
		if get.Equals(test.want) == false {
			t.Errorf("TestNormalizeVec3 %d", testIndex)
		}

		get2 := test.orig.NormalizeIn()
		if get2.Equals(test.want) == false {
			t.Errorf("TestNormalizeInVec3 %d", testIndex)
		}
	}
}

func TestSetVec3(t *testing.T) {
	var cases = []struct {
		x, y, z    float64
		orig, want *Vec3
	}{
		{1, 2, 3, &Vec3{0, 0, 0}, &Vec3{1, 2, 3}},
		{0, 0, 1, &Vec3{0, 0, 0}, &Vec3{0, 0, 1}},
		{0, 0, 1, &Vec3{1, 2, 3}, &Vec3{0, 0, 1}},
		{-1, 2, 4, &Vec3{1, -1, 3}, &Vec3{-1, 2, 4}},
	}

	for testIndex, test := range cases {
		get := test.orig.Set(test.x, test.y, test.z)
		if get.Equals(test.want) == false {
			t.Errorf("TestSetVec3 %d", testIndex)
		}

	}
}

func TestVec3InFuncs(t *testing.T) {
	a := &Vec3{1, 2, 3}
	b := &Vec3{4, 5, 6}
	a.Add(b)
	if (a.Equals(&Vec3{1, 2, 3}) == false) {
		t.Errorf("TestVec3InFuncs 1")
	}

	a.AddIn(b)
	if (a.Equals(&Vec3{5, 7, 9}) == false) {
		t.Errorf("TestVec3InFuncs 2")
	}

}
