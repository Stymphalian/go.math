package matrix

import (
	"fmt"
	"math"
)

// Right-Hand coordinate system
// // axis angle
// // quaterion
// // euler angles
// // matrix

// @param angle - In Radians
// @param v Vec3 - unit vector of the direction
func AxisAngleToQ8n(angle float64, x, y, z float64) *q8n {
	return &q8n{
		math.Cos(angle / 2),
		math.Sin(angle/2) * x,
		math.Sin(angle/2) * y,
		math.Sin(angle/2) * z,
	}
}

// @param x,y,z euler angles (radians) around the x,y and z axis
// Perform the operation in the order
// yaw =>  pitch => roll
func EulerToQ8n(yaw, pitch, roll float64) *q8n {
	// A := math.Cos(yaw/2)
	// B := math.Sin(yaw/2)
	// C := math.Cos(pitch/2)
	// D := math.Sin(pitch/2)
	// E := math.Cos(roll/2)
	// F := math.Sin(roll/2)

	// out := &q8n
	// out.w = A*C*E - B*D*F
	// out.x = B*C*E + A*D*F
	// out.y = A*D*E - B*C*F
	// out.z = A*C*F + B*D*E
	// return

	yawQ := &q8n{math.Cos(yaw / 2), 0, math.Sin(yaw / 2), 0}
	pitchQ := &q8n{math.Cos(pitch / 2), math.Sin(pitch / 2), 0, 0}
	rollQ := &q8n{math.Cos(roll / 2), 0, 0, math.Sin(roll / 2)}

	// note must be applied in reverse order
	return rollQ.MultIn(pitchQ).MultIn(yawQ)
}

func Mat4ToQ8n(mat *Mat4) *q8n {
    // Reference : http://www.flipcode.com/documents/matrfaq.html#Q55
	// 0  1  2  3
	// 4  5  6  7
	// 8  9  10 11
	// 12 13 14 15
	trace := mat.Get(0, 0) + mat.Get(1, 1) + mat.Get(2, 2) + 1

	if trace > 0 {
		s := 0.5 / math.Sqrt(trace)
		return &q8n{
			0.25 / s,
			(mat.GetAt(9) - mat.GetAt(6)) * s,
			(mat.GetAt(2) - mat.GetAt(8)) * s,
			(mat.GetAt(4) - mat.GetAt(1)) * s,
		}
	}

	// Find the column which has the maximum diagonal value
	max_col := 0
	champ := mat.Get(0, 0)
	for col := 1; col < 3; col += 1 {
		cand := mat.Get(col, col)
		if cand > champ {
			champ = cand
			max_col = col
		}
	}

    // TODO : UNTESTED!!!
	var w, x, y, z, s float64
	switch max_col {
	case 0:
		s = 2 * math.Sqrt(1.0+mat.GetAt(0)-mat.GetAt(5)-mat.GetAt(10))
		x = 0.5 / 2
		y = (mat.GetAt(4) + mat.GetAt(1)) / s
		z = (mat.GetAt(8) + mat.GetAt(2)) / s
		w = (mat.GetAt(9) + mat.GetAt(6)) / s
	case 1:
		s = 2 * math.Sqrt(1.0+mat.GetAt(5)-mat.GetAt(0)-mat.GetAt(10))
		x = (mat.GetAt(4) + mat.GetAt(1)) / s
		y = 0.5 / 2
		z = (mat.GetAt(9) + mat.GetAt(6)) / s
		w = (mat.GetAt(8) + mat.GetAt(2)) / s
	case 2:
		s = 2 * math.Sqrt(1.0+mat.GetAt(10)-mat.GetAt(0)-mat.GetAt(5))
		x = (mat.GetAt(8) + mat.GetAt(2)) / s
		y = (mat.GetAt(9) + mat.GetAt(6)) / s
		z = 0.5 / 2
		w = (mat.GetAt(4) + mat.GetAt(1)) / s
	}

	return &q8n{w, x, y, z}
}

// Takes the provided quaternion and returns the angle axis components
// The returned values are ambiguous. There is no way to know if the original
// angle and axis were specified using -ve angle or +ve angle
// Take the case such as 90 around the axis [-1,0,0]
// compared to the case of -90 around the axis [1,0,0]
// There is no ways to tell which one the user specified.
// Therefore by convention, this will always return  +ve angle case.
func Q8nToAxisAngle(q *q8n) (angle, x, y, z float64) {
	angle = 2 * math.Acos(q.w)
	s := math.Sin(angle / 2)
	x = q.x / s
	y = q.y / s
	z = q.z / s
	return
}

// func Q8nToEuler(q *q8n) (x, y, z float64){
//  	return
// }

// Calculates a Mat4 from the provided quaternion
// func Q8nToMat44(q *q8n) (mat *Mat4){
//     //      1 - 2y² - 2z²       2xy + 2wz           2xz - 2wy
//     // M =  2xy - 2wz           1 - 2x² - 2z²       2yz + 2wx
//     //      2xz + 2wy           2yz - 2wx           1 - 2x² - 2y²
// 	w,x,y,z := q.w,q.x,q.y,q.z

// 	mat.mat[0] = 1 - 2*y*y - 2*z*z
// 	mat.mat[1] = 2*x*y + 2*w*z
// 	mat.mat[2] = 2*x*z - 2*w*y
// 	mat.mat[3] = 0

// 	mat.mat[4] = 2*x*y - 2*w*z
// 	mat.mat[5] = 1 - 2*x*x - 2*z*z
// 	mat.mat[6] = 2*y*z + 2*w*x
// 	mat.mat[7] = 0

// 	mat.mat[8] = 2*x*z + 2*w*y
// 	mat.mat[9] = 2*y*z - 2*w*x
// 	mat.mat[10] = 1 - 2*x*x - 2*y*y
// 	mat.mat[11] = 0

// 	mat.mat[12] = 0
// 	mat.mat[13] = 0
// 	mat.mat[14] = 0
// 	mat.mat[15] = 1
// 	return
// }

// func AxisAngleToMat44(angle float64, x, y, z float64) (mat *Mat4) {

// 	return
// }


// Return a matrix representing the specified rotations in euler angles
// Rotations are applied in the order yaw => pitch => roll
func EulerToMat4(yaw,pitch,roll float64) (mat *Mat4) {
    mat = &Mat4{}
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
    //   roll          pitch          yaw
    // | cz  -sz  0 | | 1    0    0  | | cy   0   sy |
    // | sz   cz  0 |x| 0    cx  -sx |x| 0    1   0  |
    // | 0    0   1 | | 0    sx   cx | | -sy  0   cy |

    // first row
    mat.mat[0] = cz*cy - sx*sz*sy
    mat.mat[1] = -sz*cx
    mat.mat[2] = cz*sy + sx*sz*cy
    // second row
    mat.mat[4] = sz*cy + cz*sx*sy
    mat.mat[5] = cz*cx
    mat.mat[6] = sz*sy -cz*sx*cy
    // third row
    mat.mat[8] = -sy*cx
    mat.mat[9] = sx
    mat.mat[10] = cx*cy

	mat.mat[3] = 0
	mat.mat[7] = 0
	mat.mat[11] = 0
	mat.mat[12] = 0
	mat.mat[13] = 0
	mat.mat[14] = 0
	mat.mat[15] = 1
	return
}

// func Mat44ToAxisAngle(mat *Mat4) (angle, x, y, z float64) {
// 	return
// }

// func Mat44toEuler(mat *Mat4) (angle_x, angle_y, angle_z float64) {
// 	//Calculate Y-axis angle
// 	var trx, try, C float64
// 	//var D float64
// 	angle_y = -math.Asin(mat.mat[2])
// 	//D = angle_y
// 	C = math.Cos(angle_y)

// 	// Gimbal lock?
// 	if math.Abs(C) > 0.005 {
// 		// No, so get X-axis angle
// 		trx = mat.mat[10] / C
// 		try = -mat.mat[6] / C

// 		angle_x = math.Atan2(try, trx)

// 		// get the z-axis angle
// 		trx =  mat.mat[0] / C /* Get Z-axis angle */
// 		try = -mat.mat[1] / C

// 		angle_z = math.Atan2(try, trx)
// 	} else {
// 		// gimball lock has occured
// 		// set the x-axis angle to zero
// 		angle_x = 0
// 		// And calculate Z-axis angle
// 		trx = mat.mat[5]
// 		try = mat.mat[4]
// 		angle_z = math.Atan2(try, trx)
// 	}

// 	// clamp all the angles into the proper ranges
// 	angle_x = clamp(angle_x, 0, 2*math.Pi)
// 	angle_y = clamp(angle_y, 0, 2*math.Pi)
// 	angle_z = clamp(angle_z, 0, 2*math.Pi)
// 	return
// }

// // I need to test this...
// func LookAtMat44(eye,center,up *Vec3) (mat *Mat4){
// 	forward := center.Sub(eye).NormalizeIn()
// 	right:= forward.Cross(up).NormalizeIn()
// 	up := right.Cross(forward).NormalizeIn()

// 	mat.SetRow(0,-forward.X,-forward.Y,-forward.Z,-eye.X)
// 	mat.SetRow(1,up.X,up.Y,up.Z,-eye.Y)
// 	mat.SetRow(2,right.X,right.Y,right.Z,-eye.Z)
// 	mat.SetRow(3,0,0,0,1)
// 	return
// }

// Apply the matrix against the Vector
// Return a new vector with the result m*v
func MultMat4Vec3(m *Mat4, v *Vec3) *Vec3 {
	// 0   1   2   3
	// 4   5   6   7
	// 8   9   10  11
	// 12  13  14  15
	return &Vec3{
		m.mat[0]*v.X + m.mat[1]*v.Y + m.mat[2]*v.Z + m.mat[3],
		m.mat[4]*v.X + m.mat[5]*v.Y + m.mat[6]*v.Z + m.mat[7],
		m.mat[8]*v.X + m.mat[9]*v.Y + m.mat[10]*v.Z + m.mat[11],
	}
}

// Apply the matrix against the Vector
// Return a new vector with the result v*m
func MultVec3Mat4(m *Mat4, v *Vec3) *Vec3 {
	// 0   1   2   3
	// 4   5   6   7
	// 8   9   10  11
	// 12  13  14  15
	return &Vec3{
		v.X*m.mat[0] + v.Y*m.mat[4] + v.Z*m.mat[8] + m.mat[12],
		v.X*m.mat[1] + v.Y*m.mat[5] + v.Z*m.mat[9] + m.mat[13],
		v.X*m.mat[2] + v.Y*m.mat[6] + v.Z*m.mat[10] + m.mat[14],
	}
}

// TODO this can be made faster by assuming that
// the quaternion is a unit quaternion
// This might be a vaild assumption that we can make
// because the X,Y,Z,W parameters are private and read-only
func RotateVecQ8n(q *q8n, v *Vec3) *Vec3 {
	vq := &q8n{0.0, v.X, v.Y, v.Z}
	//inv_q := q.Conjugate()
	inv_q := q.Inverse()
	rs := q.Mult(vq).MultIn(inv_q)

	if rs.w != 0 {
		fmt.Println("What not zero!")
	}
	return &Vec3{rs.x, rs.y, rs.z}
}
