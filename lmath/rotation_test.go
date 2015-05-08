package lmath

import (
	"fmt"
	"math"
	"os"
	"testing"
)

type rotation_test_struct struct {
	pitch, yaw, roll float64
	start_vec        *Vec3
	want             *Vec3
}

var common_cases []rotation_test_struct

func TestMat4ToAxisAngle(t *testing.T) {
	cases := []struct {
		angle, x, y, z float64
	}{
		//test basic rotations using a [1,0,0] vector
		{90, 1, 0, 0},
		{90, 0, 1, 0},
		{90, 0, 0, 1},
		{45, 1, 0, 0},
		{45, 0, 1, 0},
		{45, 0, 0, 1}, //5
		{180, 1, 0, 0},
		{180, 0, 1, 0},
		{180, 0, 0, 1},
		{90, 1, 1, 0},
		{90, 1, 1, 0}, //10
		{90, 0, -1, 1},
		{45, 1, 0, 1},
		{45, 0, 1, 0},
		{45, 1, 0, 1},
		{180, 1, -2, 0}, //15
		{180, 0, 1, 20},
		{180, 0, 20, 1},
		{180, -4, 4, 1},
	}

	var m *Mat4
	for testIndex, c := range cases {
		v := &Vec3{c.x, c.y, c.z}
		v.NormalizeIn()
		m = AxisAngleToMat4(Radians(c.angle), v.X, v.Y, v.Z)
		get_angle, get_x, get_y, get_z := Mat4ToAxisAngle(m)

		m2 := AxisAngleToMat4(get_angle, get_x, get_y, get_z)
		v2 := MultMat4Vec3(m, &Vec3{1, 0, 0})
		v3 := MultMat4Vec3(m2, &Vec3{1, 0, 0})
		if v2.Eq(v3) == false {
			fmt.Printf("Not good %d %2.5f %2.5f %2.5f %2.5f %2.5f %2.5f\n", testIndex, v2.X, v2.Y, v2.Z, v3.X, v3.Y, v3.Z)
		}

		if !closeEq(Degrees(get_angle), c.angle, epsilon) ||
			!closeEq(get_x, v.X, epsilon) ||
			!closeEq(get_y, v.Y, epsilon) ||
			!closeEq(get_z, v.Z, epsilon) {

			if closeEq(get_angle, math.Pi, epsilon) &&
				closeEq(math.Abs(get_x)-math.Abs(v.X), 0, epsilon) &&
				closeEq(math.Abs(get_y)-math.Abs(v.Y), 0, epsilon) &&
				closeEq(math.Abs(get_z)-math.Abs(v.Z), 0, epsilon) {
				continue
			} else {
				t.Errorf("TestMat4ToAxisAngle %d %v \n%f %f %f %f\n%f %f %f %f\n", testIndex, v, Degrees(get_angle), get_x, get_y, get_z, c.angle, v.X, v.Y, v.Z)
			}
		}
	}
}

func TestMultMatVec(t *testing.T) {
	cases := []struct {
		mat          [16]float64
		orig_v, want *Vec3
	}{
		{[16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{[16]float64{2, 0, 0, 0, 0, 2, 0, 0, 0, 0, 2, 0, 0, 0, 0, 1}, &Vec3{1, 0, 0}, &Vec3{2, 0, 0}},
		{[16]float64{2, 0, 0, 0, 0, 2, 0, 0, 0, 0, 2, 0, 0, 0, 0, 1}, &Vec3{1, 1, 1}, &Vec3{2, 2, 2}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, &Vec3{1, 0, 0}, &Vec3{5, 13, 21}},
		{[16]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, &Vec3{1, 2, 3}, &Vec3{18, 46, 74}},
	}

	m := &Mat4{}
	for testIndex, c := range cases {
		m.Load(c.mat)
		get := MultMat4Vec3(m, c.orig_v)
		if get.Eq(c.want) == false {
			t.Errorf("TestMultMat4Vec3 %d \n%v\n%v\n\n", testIndex, m, get)
		}
	}
}

func TestEulerToMat4(t *testing.T) {
	var m *Mat4
	for testIndex, c := range common_cases {
		// m = EulerToMat4(Radians(c.yaw), Radians(c.pitch), Radians(c.roll))
		m = EulerToMat4(Radians(c.pitch), Radians(c.yaw), Radians(c.roll))
		get := MultMat4Vec3(m, c.start_vec)
		if get.Eq(c.want) == false {
			t.Errorf("TestEulerToMat4 %d \n%v\n%v\n\n", testIndex, m, get)
		}
	}
}

func TestMat4ToEuler(t *testing.T) {
	var m *Mat4
	for testIndex, c := range common_cases {
		m = EulerToMat4(Radians(c.pitch), Radians(c.yaw), Radians(c.roll))
		x, y, z := Mat4ToEuler(m)

		if x != Radians(c.pitch) || y != Radians(c.yaw) || z != Radians(c.roll) {
			t.Errorf("TestMat4ToEuler %d %f %f %f", testIndex, x, y, z)
		}
	}
}



func TestAxisAngleToMat4(t *testing.T) {
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
		{180, &Vec3{0, 0, 1}, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}}, //7

		//test basic rotations using a [0,1,0] vector
		{90, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{90, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}, &Vec3{0, 0, 1}},
		{90, &Vec3{0, 0, 1}, &Vec3{0, 1, 0}, &Vec3{-1, 0, 0}},
		{-90, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{-90, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}, &Vec3{0, 0, -1}},
		{-90, &Vec3{0, 0, 1}, &Vec3{0, 1, 0}, &Vec3{1, 0, 0}},
		{360, &Vec3{0, 0, 1}, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
		{180, &Vec3{0, 0, 1}, &Vec3{0, 1, 0}, &Vec3{0, -1, 0}}, //15

		// test negative axes
		{90, &Vec3{0, -1, 0}, &Vec3{1, 0, 0}, &Vec3{0, 0, 1}},
		{90, &Vec3{-1, 0, 0}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{90, &Vec3{0, 0, -1}, &Vec3{1, 0, 0}, &Vec3{0, -1, 0}},
		{-90, &Vec3{0, -1, 0}, &Vec3{1, 0, 0}, &Vec3{0, 0, -1}},
		{-90, &Vec3{-1, 0, 0}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{-90, &Vec3{0, 0, -1}, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}},
		{360, &Vec3{0, 0, -1}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{180, &Vec3{0, 0, -1}, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}}, //23

		// test arbitraty axis
		{360, &Vec3{1, 1, 0}, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
		{90, &Vec3{1, 1, 0}, &Vec3{1, 0, 0}, &Vec3{0.5, 0.5, -0.7071067811}},
		{45, &Vec3{1, 1, 0}, &Vec3{1, 0, 0}, &Vec3{0.85355339059, 0.1464466094067, -0.5}}, //26
	}

	var m *Mat4
	for testIndex, c := range cases {
		c.axis.NormalizeIn()
		m = AxisAngleToMat4_2(Radians(c.angle), c.axis.X, c.axis.Y, c.axis.Z)
		get := MultMat4Vec3(m, c.start_vec)
		if get.Eq(c.want) == false {
			t.Errorf("TestAxisAngleToMat4 %d \n%v\n%v\n\n", testIndex, m, get)
		}

		c.axis.NormalizeIn()
		m = AxisAngleToMat4(Radians(c.angle), c.axis.X, c.axis.Y, c.axis.Z)
		get = MultMat4Vec3(m, c.start_vec)
		if get.Eq(c.want) == false {
			t.Errorf("TestAxisAngleToMat4 %d \n%v\n%v\n\n", testIndex, m, get)
		}
	}
}

func TestMain(m *testing.M) {
	common_cases = []rotation_test_struct{

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

	os.Exit(m.Run())
}
