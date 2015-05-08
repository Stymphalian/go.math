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
		get := test.orig.Eq(test.other)
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
		if get.Eq(test.want) == false {
			t.Errorf("TestAddVec3 %d", testIndex)
		}

		get = test.orig.AddIn(test.other)
		if get != test.orig || get.Eq(test.want) == false {
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
		if get.Eq(test.want) == false {
			t.Errorf("TestSubVec3 %d", testIndex)
		}

		get = test.orig.SubIn(test.other)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestSubVec3 SubIn %d", testIndex)
		}
	}
}

func TestAddScalarVec3(t *testing.T) {
	cases := []struct{
		orig,want *Vec3
		scale float64
	}{
		{&Vec3{0,0,0},&Vec3{0,0,0},0},
		{&Vec3{0,0,0},&Vec3{1,1,1},1},
		{&Vec3{0,0,0},&Vec3{-1,-1,-1},-1},
		{&Vec3{1,2,3},&Vec3{5,6,7},4},
		{&Vec3{1,2,3},&Vec3{-3,-2,-1},-4},
		{&Vec3{1,-2,3},&Vec3{5,2,7},4},
	}

	for testIndex, test := range cases{
		get := test.orig.AddScalar(test.scale)
		if get.Eq(test.want) == false {
			t.Errorf("TestAddScalarVec3 %d",testIndex)
		}

		get = test.orig.AddInScalar(test.scale)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestAddInScalarVec3 %d",testIndex)
		}
	}
}

func TestSubScalarVec3(t *testing.T) {
	cases := []struct{
		orig,want *Vec3
		scale float64
	}{
		{&Vec3{0,0,0},&Vec3{0,0,0},0},
		{&Vec3{0,0,0},&Vec3{-1,-1,-1},1},
		{&Vec3{0,0,0},&Vec3{1,1,1},-1},
		{&Vec3{1,2,3},&Vec3{-3,-2,-1},4},
		{&Vec3{1,2,3},&Vec3{5,6,7},-4},
		{&Vec3{1,-2,3},&Vec3{-3,-6,-1},4},
	}

	for testIndex, test := range cases{
		get := test.orig.SubScalar(test.scale)
		if get.Eq(test.want) == false {
			t.Errorf("TestSubScalarVec3 %d",testIndex)
		}

		get = test.orig.SubInScalar(test.scale)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestSubInScalarVec3 %d",testIndex)
		}
	}
}

func TestMultScalarVec3(t *testing.T) {
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
		get := test.orig.MultScalar(test.scale)
		if get.Eq(test.want) == false {
			t.Errorf("TestMultVec3 %d", testIndex)
		}

		get = test.orig.MultInScalar(test.scale)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestMultVec3 MultIn %d", testIndex)
		}
	}
}

func TestDivScalarVec3(t *testing.T) {
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
		get := test.orig.DivScalar(test.scale)
		if get.Eq(test.want) == false {
			t.Errorf("TestDivVec3 %d", testIndex)
		}

		get = test.orig.DivInScalar(test.scale)
		if get != test.orig || get.Eq(test.want) == false {
			t.Errorf("TestDivInVec3 %d", testIndex)
		}
	}
}

func TestOuterVec3(t *testing.T) {
	cases := []struct {
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
		if get.Eq(test.want) == false {
			t.Errorf("TestOuterVec3 %d", testIndex)
		}

		get = test.orig.OuterIn(test.other)
		if get != test.orig || get.Eq(test.want) == false {
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
		if get.Eq(test.want) == false {
			t.Errorf("TestCrossVec3 %d", testIndex)
		}

		get = test.orig.CrossIn(test.other)
		if get!= test.orig || get.Eq(test.want) == false {
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
		if get.Eq(test.want) == false {
			t.Errorf("TestNormalizeVec3 %d", testIndex)
		}

		get = test.orig.NormalizeIn()
		if get != test.orig  || get.Eq(test.want) == false {
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
		if get.Eq(test.want) == false {
			t.Errorf("TestSetVec3 %d", testIndex)
		}
	}
}
