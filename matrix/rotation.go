package matrix

import (
	"fmt"
	"math"
)

// Right-Hand coordinate system
// All angles are specified in Radians
// rotations are applied in Yaw => Pitch => Roll order

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
// pitch => yaw => row
// x => y => z
func EulerToQ8n(pitch, yaw, roll float64) *q8n {
	yawQ := &q8n{math.Cos(yaw / 2), 0, math.Sin(yaw / 2), 0}
	pitchQ := &q8n{math.Cos(pitch / 2), math.Sin(pitch / 2), 0, 0}
	rollQ := &q8n{math.Cos(roll / 2), 0, 0, math.Sin(roll / 2)}

	// note must be applied in reverse order
	// pitch => yaw => roll
	return rollQ.MultIn(yawQ).MultIn(pitchQ)
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

// NOT DONE!
func Q8nToEuler(q *q8n) (pitch, yaw, roll float64) {
	// Reference http://www.euclideanspace.com/maths/geometry/rotations/conversions/quaternionToEuler/

	test := q.x*q.y + q.z*q.w
	fmt.Println(test)
	if test > 0.499 { // singularity at north pole
		yaw = 2 * math.Atan2(q.x, q.w)
		roll = math.Pi / 2
		pitch = 0
		return
	}
	if test < -0.499 { // singularity at south pole
		yaw = -2 * math.Atan2(q.x, q.w)
		roll = -math.Pi / 2
		pitch = 0
		return
	}
	sqx := q.x * q.x
	sqy := q.y * q.y
	sqz := q.z * q.z
	yaw = math.Atan2(2*q.y*q.w-2*q.x*q.z, 1-2*sqy-2*sqz)
	roll = math.Asin(2 * test)
	pitch = math.Atan2(2*q.x*q.w-2*q.y*q.z, 1-2*sqx-2*sqz)
	return
}

//Calculates a Mat4 from the provided quaternion
func Q8nToMat4(q *q8n) (mat *Mat4) {
	// Reference
	// Derivation of the below matrix can be found here
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/quaternionToMatrix/index.htm
	//     1 - 2y² - 2z²    2yx - 2wz        2xz + 2wy
	// M=  2xy + 2wz        1 - 2x² - 2z²    2yz - 2wx
	//     2xz - 2wy        2yz + 2wx        1 - 2x² - 2y²

	w, x, y, z := q.w, q.x, q.y, q.z

	mat = &Mat4{}
	// 0 1 2 3
	// 4 5 6 7
	// 8 9 10 11
	// 12 13 14 15
	mat.mat[0] = 1 - 2*y*y - 2*z*z
	mat.mat[1] = 2*y*z - 2*w*z
	mat.mat[2] = 2*x*z + 2*w*y
	mat.mat[3] = 0

	mat.mat[4] = 2*x*y + 2*w*z
	mat.mat[5] = 1 - 2*x*x - 2*z*z
	mat.mat[6] = 2*y*z - 2*w*x
	mat.mat[7] = 0

	mat.mat[8] = 2*x*z - 2*w*y
	mat.mat[9] = 2*y*z + 2*w*x
	mat.mat[10] = 1 - 2*x*x - 2*y*y
	mat.mat[11] = 0

	mat.mat[12] = 0
	mat.mat[13] = 0
	mat.mat[14] = 0
	mat.mat[15] = 1
	return
}

// {{x/v,-y/v,0},{y/v,x/v,0},{0,0,1}}
// {{z,0,v},{0,1,0},{-v,0,z}}
// -- {{(x z)/v, -(y/v), x}, {(y z)/v, x/v, y}, {-v, 0, z}}
// {{c,-s,0},{s,c,0},{0,0,1}}
// ---{{-((s y)/v) + (c x z)/v, -((c y)/v) - (s x z)/v, x}, {(s x)/v + (c y z)/v, (c x)/v - (s y z)/v, y}, {-(c v), s v, z}}


// {{A/v, B/v, x}, {C/v,D/v, y}, {-(c v), s v, z}}
// {{(x z)/v,(y z)/v,-v}, { -(y/v), x/v,0}, {x,y, z}}

func AxisAngleToMat4_2(angle float64, x,y,z float64) *Mat4{
	// References:
	// http://inside.mines.edu/fs_home/gmurray/ArbitraryAxisRotation/
	// http://www.engr.uvic.ca/~mech410/lectures/4_2_RotateArbi.pdf
	// Basic idea behind this method.
	// 1. Rotate the axis into the xz plane ( rotate around the x-axis or z-axis) R_xz
	// 2. Rotate the axis onto the z-axis ( rotate about y-axis) R_y
	// 3. Apply the rotation around the z-axis R_theta
	// 4. Apply the inverse rotation matrix R_y
	// 5. Apply the inverse rotation matrix R_xz

	c := math.Cos(angle)
	s := math.Sin(angle)
	// normalize the x,y,z axis components to make our life easier
	L := math.Sqrt(x*x + y*y + z*z)
	x,y,z = x/L,y/L,z/L

	if( closeEquals(x,1,epsilon) &&
		closeEquals(y,0,epsilon) &&
		closeEquals(z,0,epsilon)){
		// specified axis is aligned with the x-axis; therefore
		// perform the R_xz rotation about the z-axis instead.
		v := math.Sqrt(x*x + y*y)
		A := c*z*s -s*y
		B := (-c*y -s*x*z)
		C := s*x + c*y*s
		D := c*x -s*y*s

		mat := &Mat4{}
		mat.ToIdentity()
		mat.Set(0,0,v*v*x*x + A*z*x - B*y)
		mat.Set(0,1,v*v*y*x + B*x + A*y*z)
		mat.Set(0,2,x*z -A)

		mat.Set(1,0,x*y*v*v - D*y + C*x*z)
		mat.Set(1,1,v*v*y*y + C*z*y + D*x)
		mat.Set(1,2,y*z -C)

		mat.Set(2,0,x*z*(1- c)-s*y)
		mat.Set(2,1,s*x + y*z*(1-c))
		mat.Set(2,2,c*v*v +z*z)
		return mat
	}else{
		// rotate about the x-axis
		v := math.Sqrt(y*y + z*z)
		A:= (s*z - x*y*c)
		B:= (c*z + x*y*s)
		C:= (-x*z*c-y*s)
		D:= (x*z*s -y*c)

		mat := &Mat4{}
		mat.ToIdentity()
		mat.Set(0,0, c*v*v + x*x)
		mat.Set(0,1, x*y*(1-c) - s*z)
		mat.Set(0,2, s*y + x*z*(1-c))

		mat.Set(1,0, A + x*y)
		mat.Set(1,1, v*v*y*y - A*x*y + B*z)
		mat.Set(1,2, v*v*z*y - B*y - A*x*z)

		mat.Set(2,0, C + x*z)
		mat.Set(2,1, z*y*v*v - C*x*y + D*z)
		mat.Set(2,2, z*z*v*v - C*x*x - D*y)
		return mat
	}
}

func AxisAngleToMat4(angle float64, a, b, c float64) *Mat4 {
	eye := &Vec3{0, 0, 0}
	at := &Vec3{a, b, c}
	up := &Vec3{0,1,0}

	if at.Equals(&Vec3{0,1,0}) {
		up = &Vec3{0, 0, -1}
	} else if at.Equals(&Vec3{0,-1,0}) {
		up = &Vec3{0, 0, 1}
	}

	forward := at.Sub(eye).NormalizeIn()
	right := up.Cross(forward).NormalizeIn()
	up = forward.Cross(right).NormalizeIn()

	mat := NewMat4(right.X, right.Y, right.Z, 0,
		up.X, up.Y, up.Z, 0,
		forward.X, forward.Y, forward.Z, 0,
		0, 0, 0, 1)

	rotMat := &Mat4{}
	rotMat.ToIdentity()
	rotMat.Set(0, 0, math.Cos(angle))
	rotMat.Set(0, 1, -math.Sin(angle))
	rotMat.Set(1, 0, math.Sin(angle))
	rotMat.Set(1, 1, math.Cos(angle))

	return mat.Transpose().MultIn(rotMat).MultIn(mat)
}

// Return a matrix representing the specified rotations in euler angles
// Rotations are applied in the order pitch => yaw => roll
func EulerToMat4(pitch, yaw, roll float64) (mat *Mat4) {
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
	//   roll         yaw             pitch
	// | cz  -sz  0 | | cy   0   sy | | 1    0    0  |
	// | sz   cz  0 |x| 0    1   0  |x| 0    cx  -sx |
	// | 0    0   1 | | -sy  0   cy | | 0    sx   cx |

	// first row
	mat.mat[0] = cz * cy
	mat.mat[1] = cz*sy*sx - sz*cx
	mat.mat[2] = sz*sx + cz*cx*sy
	// second row
	mat.mat[4] = sz * cy
	mat.mat[5] = cz*cx + sx*sy*sz
	mat.mat[6] = sz*sy*cx - cz*sx
	// third row
	mat.mat[8] = -sy
	mat.mat[9] = sx * cy
	mat.mat[10] = cy * cx

	mat.mat[3] = 0
	mat.mat[7] = 0
	mat.mat[11] = 0
	mat.mat[12] = 0
	mat.mat[13] = 0
	mat.mat[14] = 0
	mat.mat[15] = 1
	return
}

func Mat44ToAxisAngle(mat *Mat4) (angle, x, y, z float64) {
	return
}

// Return the pitch,yaw and roll values for the given rotation matrix
// Assumption is that mat is a valid rotation matrix
// following the conventions of this package
// (x-y-z rotation order,row-major order, right-handed)
func Mat4ToEuler(mat *Mat4) (pitch, yaw, roll float64) {
	// The method for calculating the euler angles from a rotation matrix
	// uses the method described in this document
	// http://staff.city.ac.uk/~sbbh653/publications/euler.pdf

	// The rotation matrix we are using will be of the following form
	// cos(x) is abbreviated as cx ( similarily sin(x) = sx)
	//
	// cz*cy       cz*sy*sx - sz*cx         sz*sx + cz*cx*sy   | r11 r12 r13
	// sz*cy       cz*cx + sx*sy*sz         sz*sy*cx - cz*sx   | r21 r22 r23
	// -sy         sx*cy                    cx*cy              | r31 r32 r33

	// We want to determine the x,y,z angles
	// 1) Find the 'y' angle
	//      This is easily accomplished because term r31 is simply '-sin(y)'
	// 2) There are two possible angles for y because
	//      sin(y) == sin( pi - y)
	// 3) To find the value of x, we observe the following
	//      r32/r33 = tan(x)
	//      (sin(x)cos(y)) / (cos(x)cos(y))
	//      (sin(x)/cos(x)) == tan(x) by defn.
	// 4) Therefore we can calculate x by.
	//      x = atan2(r32,r33)

	var x, y, z float64
	r31 := mat.Get(2, 0)
	if closeEquals(r31, 1, epsilon) {
		// we are in gimbal lock
		fmt.Println("1")
		z = 0
		y = -math.Pi / 2
		x = -z + math.Atan2(-mat.Get(0, 1), -mat.Get(0, 2))
	} else if closeEquals(r31, -1, epsilon) {
		// we are in gimbal lock
		fmt.Println("-1")
		z = 0
		y = math.Pi / 2
		x = z + math.Atan2(mat.Get(0, 1), mat.Get(0, 2))
	} else {
		fmt.Println("normal")
		y = -math.Asin(r31)
		// y = math.Pi + math.Asin(r31) // alt-solution
		cos_y := math.Cos(y)
		x = math.Atan2(mat.Get(2, 1)/cos_y, mat.Get(2, 2)/cos_y)
		z = math.Atan2(mat.Get(1, 0)/cos_y, mat.Get(0, 0)/cos_y)

		// y2 := math.Pi + math.Asin(r31) // alt-solution
		// cos_y2 := math.Cos(y2)
		// x2 := math.Atan2(mat.Get(2,1)/cos_y2,mat.Get(2,2)/cos_y2)
		// z2 := math.Atan2(mat.Get(1,0)/cos_y2,mat.Get(0,0)/cos_y2)

		// fmt.Printf("%2.4f %2.5f %2.5f %2.5f %2.5f %2.5f \n",x,y,z,x2,y2,z2)
	}

	pitch = x
	yaw = y
	roll = z
	return
}

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

	if !closeEquals(rs.w, 0, epsilon) {
		fmt.Println("What not zero!", rs)
	}
	return &Vec3{rs.x, rs.y, rs.z}
}
