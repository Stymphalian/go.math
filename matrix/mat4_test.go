package matrix

import (
	// "fmt"
	// "math"
	"testing"
)

//Test creation with 16 values
//Test Equals with matrices
func TestNewMat4(t *testing.T) {
	m := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)
	m2 := NewMat4(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)
	if m.Equals(m2) == false {
		t.Errorf("TestNewMat ")
	}
}

// Test Get
// Test Set
func TestMatGettersSetter(t *testing.T) {
	cases := []struct {
		orig          *Mat4
		x, y          int
		wantBeforeSet float64
		wantAfterSet  float64
	}{
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 0, 0, 1, 10},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 1, 0, 2, 20},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 2, 0, 3, 30},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 3, 0, 4, 40},

		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 0, 1, 5, 50},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 1, 1, 6, 60},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 2, 1, 7, 70},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 3, 1, 8, 80},

		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 0, 2, 9, 90},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 1, 2, 10, 100},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 2, 2, 11, 110},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 3, 2, 12, 120},

		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 0, 3, 13, 130},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}, 1, 3, 14, 130},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, -15, 16}}, 2, 3, -15, -150},
		{&Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, -16}}, 3, 3, -16, -160},
	}

	for testIndex, c := range cases {
		get := c.orig.Get(c.y, c.x)
		if get != c.wantBeforeSet {
			t.Errorf("TestMatGetttersSetter wantBeforeSet %d %v", testIndex, get)
		}

		get = c.orig.Set(c.y, c.x, c.wantAfterSet).Get(c.y, c.x)
		if get != c.wantAfterSet {
			t.Errorf("TestMatGetttersSetter wantAfterSet %d %v", testIndex, get)
		}
	}

	orig := &Mat4{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}}
	for k, _ := range orig.Dump() {
		get := orig.GetAt(k)
		if get != float64(k+1) {
			t.Errorf("TestMatGetterSetter GetAt %d %v", k, get)
		}

		orig.SetAt(k, float64((k+1)*10))
		get = orig.GetAt(k)
		if get != float64((k+1)*10) {
			t.Errorf("TestMatGetterSetter SetAt %d", k)
		}
	}
}

// Test Load Array
// Test Dump
// Test Dump
func TestLoadDump(t *testing.T) {
	cases := []struct {
		loadArray [16]float64
	}{
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[16]float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13, -14, -15, -16}},
	}

	m := &Mat4{}
	for testIndex, c := range cases {
		m.Load(c.loadArray)
		get := m.Dump()

		for k, _ := range get {
			if get[k] != c.loadArray[k] {
				t.Errorf("TestLoadDump %d", testIndex)
				break
			}
		}
	}

	m.Load([16]float64{1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15, 4, 8, 12, 16})
	get := m.DumpOpenGL()
	for k, _ := range get {
		if get[k] != cases[0].loadArray[k] {
			t.Errorf("TestDumpOpenGL")
			break
		}
	}
}

func TestGetRow(t *testing.T) {
	cases := []struct {
		orig     [16]float64
		rowIndex int
		want     [4]float64
	}{
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 0, [4]float64{1, 2, 3, 4}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1, [4]float64{5, 6, 7, 8}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 2, [4]float64{9, 10, 11, 12}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 3, [4]float64{13, 14, 15, 16}},
	}

	m := &Mat4{}
	var x, y, z, w float64
	for testIndex, c := range cases {
		m.Load(c.orig)
		x, y, z, w = m.GetRow(c.rowIndex)
		if x != c.want[0] || y != c.want[1] || z != c.want[2] || w != c.want[3] {
			t.Errorf("TestGetRow %d", testIndex)
		}
	}
}

func TestSetRow(t *testing.T) {
	cases := []struct {
		orig       [16]float64
		rowIndex   int
		x, y, z, w float64
	}{
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 0, -1, -2, -3, -4},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1, -5, -6, -7, -8},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 2, -9, -10, -11, -12},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 3, -13, -14, -15, -16},
	}

	m := &Mat4{}
	var x, y, z, w float64
	for testIndex, c := range cases {
		m.Load(c.orig)
		m.SetRow(c.rowIndex, c.x, c.y, c.z, c.w)
		x, y, z, w = m.GetRow(c.rowIndex)
		if x != c.x || y != c.y || z != c.z || w != c.w {
			t.Errorf("TestSetRow %d", testIndex)
		}
	}
}

func TestGetCol(t *testing.T) {
	cases := []struct {
		orig     [16]float64
		colIndex int
		want     [4]float64
	}{
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 0, [4]float64{1, 5, 9, 13}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1, [4]float64{2, 6, 10, 14}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 2, [4]float64{3, 7, 11, 15}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 3, [4]float64{4, 8, 12, 16}},
	}

	m := &Mat4{}
	var x, y, z, w float64
	for testIndex, c := range cases {
		m.Load(c.orig)
		x, y, z, w = m.GetCol(c.colIndex)
		if x != c.want[0] || y != c.want[1] || z != c.want[2] || w != c.want[3] {
			t.Errorf("TestGetRow %d", testIndex)
		}
	}
}

func TestSetCol(t *testing.T) {
	cases := []struct {
		orig       [16]float64
		colIndex   int
		x, y, z, w float64
	}{
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 0, -1, -2, -3, -4},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1, -5, -6, -7, -8},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 2, -9, -10, -11, -12},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 3, -13, -14, -15, -16},
	}

	m := &Mat4{}
	var x, y, z, w float64
	for testIndex, c := range cases {
		m.Load(c.orig)
		m.SetCol(c.colIndex, c.x, c.y, c.z, c.w)
		x, y, z, w = m.GetCol(c.colIndex)
		if x != c.x || y != c.y || z != c.z || w != c.w {
			t.Errorf("TestSetRow %d", testIndex)
		}
	}
}

func TestAddInConst(t *testing.T) {
	cases := []struct {
		orig  [16]float64
		value float64
	}{
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1},
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, -1},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 0},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, -1},
	}

	m := &Mat4{}
	for testIndex, c := range cases {
		m.Load(c.orig)
		m.AddInConst(c.value)

		get := m.Dump()
		for k, _ := range c.orig {
			if get[k] != c.orig[k]+c.value {
				t.Errorf("TestAddInConst %d %d", testIndex, k)
				break
			}
		}
	}
}

func TestSubInConst(t *testing.T) {
	cases := []struct {
		orig  [16]float64
		value float64
	}{
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1},
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, -1},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 0},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, -1},
	}

	m := &Mat4{}
	for testIndex, c := range cases {
		m.Load(c.orig)
		m.SubInConst(c.value)

		get := m.Dump()
		for k, _ := range c.orig {
			if get[k] != c.orig[k]-c.value {
				t.Errorf("TestSubInConst %d %d", testIndex, k)
				break
			}
		}
	}
}

func TestMultInConst(t *testing.T) {
	cases := []struct {
		orig  [16]float64
		value float64
	}{
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1},
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, -1},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 0},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, -1},
	}

	m := &Mat4{}
	for testIndex, c := range cases {
		m.Load(c.orig)
		m.MultInConst(c.value)

		get := m.Dump()
		for k, _ := range c.orig {
			if get[k] != c.orig[k]*c.value {
				t.Errorf("TestMultInConst %d %d", testIndex, k)
				break
			}
		}
	}
}

func TestDivInConst(t *testing.T) {
	cases := []struct {
		orig  [16]float64
		value float64
	}{
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1},
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, -1},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, -1},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 2},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, -2},
	}

	m := &Mat4{}
	for testIndex, c := range cases {
		m.Load(c.orig)
		m.DivInConst(c.value)

		get := m.Dump()
		for k, _ := range c.orig {
			if closeEquals(get[k], c.orig[k]/c.value, epsilon) == false {
				t.Errorf("TestDivInConst %d %d", testIndex, k)
				break
			}
		}
	}
}

func TestAddIn(t *testing.T) {
	cases := []struct {
		orig, other [16]float64
	}{
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13, -14, -15, -16}},
	}

	m := &Mat4{}
	m2 := &Mat4{}
	for testIndex, c := range cases {
		m.Load(c.orig)
		m2.Load(c.other)
		m.AddIn(m2)

		get := m.Dump()
		for k, _ := range c.orig {
			if closeEquals(get[k], c.orig[k]+c.other[k], epsilon) == false {
				t.Errorf("TestAddIn %d %d", testIndex, k)
				break
			}
		}
	}
}

func TestSubIn(t *testing.T) {
	cases := []struct {
		orig, other [16]float64
	}{
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -12, -13, -14, -15, -16}},
	}

	m := &Mat4{}
	m2 := &Mat4{}
	for testIndex, c := range cases {
		m.Load(c.orig)
		m2.Load(c.other)
		m.SubIn(m2)

		get := m.Dump()
		for k, _ := range c.orig {
			if closeEquals(get[k], c.orig[k]-c.other[k], epsilon) == false {
				t.Errorf("TestSubIn %d %d", testIndex, k)
				break
			}
		}
	}
}

func TestMultIn(t *testing.T) {
	cases := []struct {
		orig, other, want [16]float64
	}{
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},

		{[16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1},
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},

		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1},
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},

		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{90, 100, 110, 120, 202, 228, 254, 280, 314, 356, 398, 440, 426, 484, 542, 600}},

		{[16]float64{5, 6, 7, 8, 1, 2, 3, 4, 13, 14, 15, 16, 9, 10, 11, 12},
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{202, 228, 254, 280, 90, 100, 110, 120, 426, 484, 542, 600, 314, 356, 398, 440}},

		{
			[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{5, 6, 7, 8, 1, 2, 3, 4, 13, 14, 15, 16, 9, 10, 11, 12},
			[16]float64{82, 92, 102, 112, 194, 220, 246, 272, 306, 348, 390, 432, 418, 476, 534, 592}},
	}

	orig := &Mat4{}
	other := &Mat4{}
	for testIndex, c := range cases {
		orig.Load(c.orig)
		other.Load(c.other)
		orig.MultIn(other)

		get := orig.Dump()
		for k, _ := range c.orig {
			if closeEquals(get[k], c.want[k], epsilon) == false {
				t.Errorf("TestMultIn %d %d", testIndex, k)
				// fmt.Println(orig, "\n")
				break
			}
		}
	}
}

func TestIdentity(t *testing.T) {
	m := &Mat4{}
	m.Load([16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
	m.ToIdentity()

	get := m.Dump()
	want := [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	for k, _ := range get {
		if want[k] != get[k] {
			t.Errorf("TestIdentity %d", k)
			break
		}
	}
}

func TestTranspose(t *testing.T) {
	cases := []struct {
		orig, want [16]float64
	}{
		{[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},

		{[16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1},
			[16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}},

		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[16]float64{1, 5, 9, 13, 2, 6, 10, 14, 3, 7, 11, 15, 4, 8, 12, 16}},
	}

	orig := &Mat4{}
	for testIndex, c := range cases {
		orig.Load(c.orig)
		get := orig.Transpose().Dump()

		for k, _ := range c.want {
			if closeEquals(get[k], c.want[k], epsilon) == false {
				t.Errorf("TestTranspose %d %d", testIndex, k)
				break
			}
		}
	}
}

func TestDeterminant(t *testing.T) {
	want := 0.554020973727016224
	m := &Mat4{}
	m.Load([16]float64{
		0.5 * 0.5,
		-0.5 * 0.866,
		0.866,
		0,

		0.866*0.866*0.866 + 0.5*0.5,
		-0.866*0.866*0.866 + 0.5*0.5,
		0.866 * 0.5,
		0,

		-0.5*0.866*0.5 + 0.866*0.866,
		0.5*0.866*0.866 + 0.866*0.5,
		-0.5 * 0.5,
		0,

		0, 0, 0, 1})

	get := m.Determinant()
	if closeEquals(get, want, epsilon) == false {
		t.Errorf("TestDeterminant %v", get)
	}
}

func TestAdjoint(t *testing.T) {
	m := &Mat4{}
	m.Load([16]float64{
		0.5 * 0.5,
		-0.5 * 0.866,
		0.866,
		0,

		0.866*0.866*0.866 + 0.5*0.5,
		-0.866*0.866*0.866 + 0.5*0.5,
		0.866 * 0.5,
		0,

		-0.5*0.866*0.5 + 0.866*0.866,
		0.5*0.866*0.866 + 0.866*0.5,
		-0.5 * 0.5,
		0,

		0, 0, 0, 1})

	mat := m.Adjoint()
	get := mat.Dump()
	want := [16]float64{
		-0.249989, 0.591459, 0.158445, 0,
		0.455852, -0.524473, 0.670684, 0,
		0.939841, -0.432981, 0.289602, 0,
		0, 0, 0, 0.554021}

	for k, _ := range want {
		// NOTE: the check uses lower precision because
		// the adjoint values I got from wolframalpha
		// only went up to 6 places
		if closeEquals(get[k], want[k], 0.0001) == false {
			t.Errorf("TestAdjoint %d %v %v", k, get[k], want[k])
			break
		}
	}
}

func TestInverseMatrix(t *testing.T) {
	m := &Mat4{}
	cases := []struct{
		orig,want [16]float64
		want_inverse_flag bool
	}{
		{[16]float64{
		0.5 * 0.5,
		-0.5 * 0.866,
		0.866,
		0,

		0.866*0.866*0.866 + 0.5*0.5,
		-0.866*0.866*0.866 + 0.5*0.5,
		0.866 * 0.5,
		0,

		-0.5*0.866*0.5 + 0.866*0.866,
		0.5*0.866*0.866 + 0.866*0.5,
		-0.5 * 0.5,
		0,

		0, 0, 0, 1},[16]float64{
		-0.451227, 1.06758, 0.285991, 0,
		0.822806, -0.946666, 1.21058, 0,
		1.6964, -0.781524, 0.522727, 0,
		0, 0, 0, 1},true},
		{[16]float64{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
			[16]float64{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},false},
		{[16]float64{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1},
			[16]float64{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1},true},
	}

	for testIndex,c := range cases {
		m.Load(c.orig)
		get_inverse_flag := m.HasInverse()
		if get_inverse_flag != c.want_inverse_flag {
				t.Errorf("TestInverse %d %v",testIndex,get_inverse_flag)
				continue
		}

		mat := m.Inverse()
		get := mat.Dump()

		for k, _ := range c.want {
			if closeEquals(get[k], c.want[k], 0.0001) == false {
				t.Errorf("TestInverse %d %d %v %v",testIndex, k, get[k], c.want[k])
				break
			}
		}
	}
}

func TestToTranslate(t *testing.T) {
	cases := []struct {
		x, y, z float64
		want    [16]float64
	}{
		{0, 0, 0, [16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}},
		{1, 2, 3, [16]float64{1, 0, 0, 1, 0, 1, 0, 2, 0, 0, 1, 3, 0, 0, 0, 1}},
		{-1, -2, -3, [16]float64{1, 0, 0, -1, 0, 1, 0, -2, 0, 0, 1, -3, 0, 0, 0, 1}},
	}

	orig := &Mat4{}
	for testIndex, c := range cases {

		orig.ToTranslate(c.x, c.y, c.z)
		get := orig.Dump()
		for k, _ := range c.want {
			if c.want[k] != get[k] {
				t.Errorf("TestToTranslate %d", testIndex)
				break
			}
		}
	}
}

func TestToScale(t *testing.T) {
	cases := []struct {
		x, y, z float64
		want    [16]float64
	}{
		{0, 0, 0, [16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{1, 2, 3, [16]float64{1, 0, 0, 0, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 0, 1}},
		{-1, -2, -3, [16]float64{-1, 0, 0, 0, 0, -2, 0, 0, 0, 0, -3, 0, 0, 0, 0, 1}},
	}

	orig := &Mat4{}
	for testIndex, c := range cases {

		orig.ToScale(c.x, c.y, c.z)
		get := orig.Dump()
		for k, _ := range c.want {
			if c.want[k] != get[k] {
				t.Errorf("TestToScale %d", testIndex)
				break
			}
		}
	}
}

func TestToSkew(t *testing.T) {
	cases := []struct {
		x, y, z float64
		want    [16]float64
	}{
		{0, 0, 0, [16]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}},
		{1, 2, 3, [16]float64{
			0, -3, 2, 0,
			3, 0, -1, 0,
			-2, 1, 0, 0,
			0, 0, 0, 1,
		}},
		{-1, -2, -3, [16]float64{
			0, 3, -2, 0,
			-3, 0, 1, 0,
			2, -1, 0, 0,
			0, 0, 0, 1,
		}},
	}

	orig := &Mat4{}
	for testIndex, c := range cases {

		orig.ToSkew(c.x, c.y, c.z)
		get := orig.Dump()
		for k, _ := range c.want {
			if c.want[k] != get[k] {
				t.Errorf("TestToScale %d", testIndex)
				break
			}
		}
	}
}
