package matrix

import (
	// "fmt"
    "os"
	"math"
	"testing"
)

type rotation_test_struct struct{
    pitch,yaw,roll float64
    start_vec        *Vec3
    want             *Vec3
}
var common_cases []rotation_test_struct

func TestAxisAngleToQ8n(t *testing.T) {
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

	var q *q8n
	for testIndex, c := range cases {
		c.axis.NormalizeIn()
		q = AxisAngleToQ8n(degToRad(c.angle), c.axis.X, c.axis.Y, c.axis.Z)
		get := RotateVecQ8n(q, c.start_vec)
		if get.Equals(c.want) == false {
			t.Errorf("TestAxisAngleToQ8n %d \n %v\n%v\n\n", testIndex, q, get)
		}
	}
}

func TestEulerToQ8n(t *testing.T) {
	var q *q8n
	for testIndex, c := range common_cases {
		q = EulerToQ8n(degToRad(c.pitch), degToRad(c.yaw), degToRad(c.roll))
		get := RotateVecQ8n(q, c.start_vec)
		if get.Equals(c.want) == false {
			t.Errorf("TestEulerToQ8n %d \n %v\n%v\n\n", testIndex, q, get)
		}
	}
}

func TestQ8nToAxisAngle(t *testing.T) {
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

	var q *q8n
	for testIndex, c := range cases {
		v := &Vec3{c.x, c.y, c.z}
		v.NormalizeIn()
		q = AxisAngleToQ8n(degToRad(c.angle), v.X, v.Y, v.Z)
		get_angle, get_x, get_y, get_z := Q8nToAxisAngle(q)
		if !closeEquals(radToDeg(get_angle), c.angle, epsilon) ||
			!closeEquals(get_x, v.X, epsilon) ||
			!closeEquals(get_y, v.Y, epsilon) ||
			!closeEquals(get_z, v.Z, epsilon) {
			t.Errorf("TestQ8nToAxisAngle %d %v %f %f %f %f\n%f %f %f %f\n", testIndex, v, radToDeg(get_angle), get_x, get_y, get_z, c.angle, v.X, v.Y, v.Z)
		}
	}
}

func TestMultMatVec(t *testing.T) {
    cases := []struct {
        mat [16]float64
        orig_v, want *Vec3
    }{
        {[16]float64{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1},&Vec3{1,0,0},&Vec3{1,0,0}},
        {[16]float64{2,0,0,0,0,2,0,0,0,0,2,0,0,0,0,1},&Vec3{1,0,0},&Vec3{2,0,0}},
        {[16]float64{2,0,0,0,0,2,0,0,0,0,2,0,0,0,0,1},&Vec3{1,1,1},&Vec3{2,2,2}},
        {[16]float64{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16},&Vec3{1,0,0},&Vec3{5,13,21}},
        {[16]float64{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16},&Vec3{1,2,3},&Vec3{18,46,74}},
    }

    m := &Mat4{}
    for testIndex, c := range cases {
        m.Load(c.mat)
        get := MultMat4Vec3(m,c.orig_v)
        if get.Equals(c.want) == false {
            t.Errorf("TestMultMat4Vec3 %d \n%v\n%v\n\n", testIndex, m, get)
        }
    }
}



func TestEulerToMat4(t *testing.T) {
    var m *Mat4
    for testIndex, c := range common_cases {
        // m = EulerToMat4(degToRad(c.yaw), degToRad(c.pitch), degToRad(c.roll))
        m = EulerToMat4(degToRad(c.pitch),degToRad(c.yaw),degToRad(c.roll))
        get := MultMat4Vec3(m,c.start_vec)
        if get.Equals(c.want) == false {
            t.Errorf("TestEulerToMat4 %d \n%v\n%v\n\n", testIndex, m, get)
        }
    }
}


func TestMat4ToQ8n(t *testing.T) {
    var m *Mat4
    var q *q8n
    for testIndex, c := range common_cases {
        m = EulerToMat4(degToRad(c.pitch),degToRad(c.yaw),degToRad(c.roll))
        q = Mat4ToQ8n(m)

        get := RotateVecQ8n(q, c.start_vec)
        if get.Equals(c.want) == false {
            t.Errorf("TestMat4ToQ8n %d \n %v\n%v\n\n", testIndex, q, get)
        }
    }
}

func TestMat4ToEuler(t *testing.T) {
    var m *Mat4
    for testIndex, c := range common_cases {
        m = EulerToMat4(degToRad(c.pitch),degToRad(c.yaw),degToRad(c.roll))
        x,y,z := Mat4ToEuler(m)

        if x != degToRad(c.pitch) || y != degToRad(c.yaw) || z != degToRad(c.roll){
            t.Errorf("TestMat4ToEuler %d %f %f %f",testIndex,x,y,z)
        }
    }
}

func TestQ8nToMat4(t *testing.T) {
    var m *Mat4
    var q *q8n
    for testIndex, c := range common_cases {
        q =  EulerToQ8n(degToRad(c.pitch),degToRad(c.yaw),degToRad(c.roll))
        m = Q8nToMat4(q)

        get := MultMat4Vec3(m,c.start_vec)
        if get.Equals(c.want) == false {
            t.Errorf("TestQ8nToMat4 %d \n%v\n%v\n\n", testIndex, m, get)
        }
    }
}

// func TestQ8nToEuler(t *testing.T) {
//     var q *q8n
//     for testIndex, c := range common_cases {
//         q =  EulerToQ8n(degToRad(c.yaw),degToRad(c.pitch),degToRad(c.roll))
//         yaw,pitch,roll :=  Q8nToEuler(q)

//         if(!closeEquals(yaw,   degToRad(c.yaw),epsilon) ||
//             !closeEquals(pitch,degToRad(c.pitch),epsilon) ||
//             !closeEquals(roll, degToRad(c.roll),epsilon)){
//             t.Errorf("TestQ8nToEuler %d %f %f %f ",testIndex,yaw,pitch,roll)
//         }
//     }
// }

func TestMain(m *testing.M){
    common_cases = []rotation_test_struct {
        //test basic rotations using a [0,1,0] vector
        // pitch,yaw,roll
        {0, 0, 90, &Vec3{0, 1, 0}, &Vec3{-1, 0, 0}},
        {0, 90, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
        {90, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 0, 1}},
        {0, 0, -90, &Vec3{0, 1, 0}, &Vec3{1, 0, 0}},
        {0, -90, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},
        {-90, 0, 0, &Vec3{0, 1, 0}, &Vec3{0, 0, -1}},
        {0, 180, 0, &Vec3{0, 1, 0}, &Vec3{0, 1, 0}},//6

        // test basic rotation using a [1,0,0] vector
        {0, 0, 90, &Vec3{1, 0, 0}, &Vec3{0, 1, 0}},
        {0, 90, 0, &Vec3{1, 0, 0}, &Vec3{0, 0, -1}},
        {90, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
        {0, 0, -90, &Vec3{1, 0, 0}, &Vec3{0, -1, 0}},
        {0, -90, 0, &Vec3{1, 0, 0}, &Vec3{0, 0, 1}},
        {-90, 0, 0, &Vec3{1, 0, 0}, &Vec3{1, 0, 0}},
        {0, 0, 180, &Vec3{1, 0, 0}, &Vec3{-1, 0, 0}},//13

        // basic rotation using a non major axis vector
        {0, 0, 90, &Vec3{1, 1, 0}, &Vec3{-1, 1, 0}},
        {0, 90, 0, &Vec3{1, -1, 0}, &Vec3{0, -1, -1}},
        {90, 0, 0, &Vec3{-1, -1, 0}, &Vec3{-1, 0, -1}},//16

        // two rotations
        {90, 0, 45, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, -math.Sqrt(2) / 2, 0}},
        {90, 45, 0, &Vec3{0, 0, 1}, &Vec3{0, -1, 0}},
        {45, 90, 0, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, -math.Sqrt(2) / 2, 0}},
        {45, 90, 90, &Vec3{0, 0, 1}, &Vec3{math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0}},
    }

    os.Exit(m.Run())
}