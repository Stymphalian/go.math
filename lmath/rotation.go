package lmath

import (
	// "fmt"
	"math"
)

// Rotation functions used to create
// rotaiton matrices, quaternions from/to euler, and axis/angle
//
// Right-Hand coordinate system
// All angles are specified in Radians
// rotations are applied in Yaw => Pitch => Roll order


// Return a rotation matrix which rotates a vector about the axis [x,y,z] with
// the given angle (radians).
func AxisAngleToMat4(angle float64, x, y, z float64) *Mat4 {
	//Reference http://en.wikipedia.org/wiki/Rotation_matrix
	c := math.Cos(angle)
	s := math.Sin(angle)
	t := (1 - c)

	mat := NewMat4(c+x*x*t, x*y*t-z*s, x*z*t+y*s, 0,
		y*x*t+z*s, c+y*y*t, y*z*t-x*s, 0,
		z*x*t-y*s, z*y*t+x*s, c+z*z*t, 0,
		0, 0, 0, 1)
	return mat
}

// Return a rotation matrix which rotates a vector about the axis [x,y,z] with
// the given angle.
func AxisAngleToMat4_2(angle float64, x, y, z float64) *Mat4 {
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
	x, y, z = x/L, y/L, z/L

	if closeEq(x, 1, epsilon) &&
		closeEq(y, 0, epsilon) &&
		closeEq(z, 0, epsilon) {
		// specified axis is aligned with the x-axis; therefore
		// perform the R_xz rotation about the z-axis instead.
		v := math.Sqrt(x*x + y*y)
		A := c*z*s - s*y
		B := (-c*y - s*x*z)
		C := s*x + c*y*s
		D := c*x - s*y*s

		mat := &Mat4{}
		mat.ToIdentity()
		mat.Set(0, 0, v*v*x*x+A*z*x-B*y)
		mat.Set(0, 1, v*v*y*x+B*x+A*y*z)
		mat.Set(0, 2, x*z-A)

		mat.Set(1, 0, x*y*v*v-D*y+C*x*z)
		mat.Set(1, 1, v*v*y*y+C*z*y+D*x)
		mat.Set(1, 2, y*z-C)

		mat.Set(2, 0, x*z*(1-c)-s*y)
		mat.Set(2, 1, s*x+y*z*(1-c))
		mat.Set(2, 2, c*v*v+z*z)
		return mat
	} else {
		// rotate about the x-axis
		v := math.Sqrt(y*y + z*z)
		A := (s*z - x*y*c)
		B := (c*z + x*y*s)
		C := (-x*z*c - y*s)
		D := (x*z*s - y*c)

		mat := &Mat4{}
		mat.ToIdentity()
		mat.Set(0, 0, c*v*v+x*x)
		mat.Set(0, 1, x*y*(1-c)-s*z)
		mat.Set(0, 2, s*y+x*z*(1-c))

		mat.Set(1, 0, A+x*y)
		mat.Set(1, 1, v*v*y*y-A*x*y+B*z)
		mat.Set(1, 2, v*v*z*y-B*y-A*x*z)

		mat.Set(2, 0, C+x*z)
		mat.Set(2, 1, z*y*v*v-C*x*y+D*z)
		mat.Set(2, 2, z*z*v*v-C*x*x-D*y)
		return mat
	}
}

// Alternate implementation of creating a rotation matrix which rotates a vector
// about the axis [a,b,c] with the given angle.
func AxisAngleToMat4_3(angle float64, a, b, c float64) *Mat4 {
	eye := &Vec3{0, 0, 0}
	at := &Vec3{a, b, c}
	up := &Vec3{0, 1, 0}

	if at.Equals(&Vec3{0, 1, 0}) {
		up = &Vec3{0, 0, -1}
	} else if at.Equals(&Vec3{0, -1, 0}) {
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

// Return a matrix representing the specified rotations in euler angles (radians)
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

func Mat4ToAxisAngle(mat *Mat4) (angle, x, y, z float64) {
	// Reference
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToAngle/
	m00, m01, m02 := mat.Get(0, 0), mat.Get(0, 1), mat.Get(0, 2)
	m10, m11, m12 := mat.Get(1, 0), mat.Get(1, 1), mat.Get(1, 2)
	m20, m21, m22 := mat.Get(2, 0), mat.Get(2, 1), mat.Get(2, 2)

	if closeEq(math.Abs(m01-m10), 0, epsilon) &&
		closeEq(math.Abs(m02-m20), 0, epsilon) &&
		closeEq(math.Abs(m12-m21), 0, epsilon) {
		// singularity check
		// Checking for cases in which the angle is either 0 or 180

		if mat.IsIdentity() {
			// If the angle is 0, then the rotation matrix will be the identity matrix
			// A 0 angle means that there is an arbitrary axis.
			angle, x, y, z = 0, 1, 0, 0
			return
		}

		// Angle is 180, we need to find the axis it rotates around
		angle = math.Pi

		xx := (m00 + 1) / 2
		yy := (m11 + 1) / 2
		zz := (m22 + 1) / 2
		xy := (m01 + m10) / 4
		xz := (m02 + m20) / 4
		yz := (m12 + m21) / 4

		if (xx > yy) && (xx > zz) { // m[0][0] is the largest diagonal term
			if xx < epsilon {
				x = 0
				y = math.Sqrt(2) / 2
				z = math.Sqrt(2) / 2
			} else {
				x = math.Sqrt(xx)
				y = xy / x
				z = xz / x
			}
		} else if yy > zz { // m[1][1] is the largest diagonal term
			if yy < epsilon {
				x = math.Sqrt(2) / 2
				y = 0
				z = math.Sqrt(2) / 2
			} else {
				y = math.Sqrt(yy)
				x = xy / y
				z = yz / y
			}
		} else { // m[2][2] is the largest diagonal term so base result on this
			if zz < epsilon {
				x = math.Sqrt(2) / 2
				y = math.Sqrt(2) / 2
				z = 0
			} else {
				z = math.Sqrt(zz)
				x = xz / z
				y = yz / z
			}
		}
		return
	}

	// no singularity; therefore calculate as normal
	angle = math.Acos((m00 + m11 + m22 - 1) / 2)
	A := (m21 - m12)
	B := (m02 - m20)
	C := (m10 - m01)

	x = A / math.Sqrt(A*A+B*B+C*C)
	y = B / math.Sqrt(A*A+B*B+C*C)
	z = C / math.Sqrt(A*A+B*B+C*C)
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
	// This corresponds to the pitch => yaw => roll rotation matrix
	// cz*cy       cz*sy*sx - sz*cx         sz*sx + cz*cx*sy   | r11 r12 r13
	// sz*cy       cz*cx + sx*sy*sz         sz*sy*cx - cz*sx   | r21 r22 r23
	// -sy         sx*cy                    cx*cy              | r31 r32 r33

	// We want to determine the x,y,z angles
	// 1) Find the 'y' angle
	//      This is easily accomplished because term r31 is simply '-sin(y)'
	// 2) There are two possible angles for y because
	//      sin(y) == sin(pi - y)
	// 3) To find the value of x, we observe the following
	//      r32/r33 = tan(x)
	//      (sin(x)cos(y)) / (cos(x)cos(y))
	//      (sin(x)/cos(x)) == tan(x) by defn.
	// 4) Therefore we can calculate x by.
	//      x = atan2(r32,r33)

	var x, y, z float64
	r31 := mat.Get(2, 0)
	if closeEq(r31, 1, epsilon) {
		// we are in gimbal lock
		z = 0
		y = -math.Pi / 2
		x = -z + math.Atan2(-mat.Get(0, 1), -mat.Get(0, 2))
	} else if closeEq(r31, -1, epsilon) {
		// we are in gimbal lock
		z = 0
		y = math.Pi / 2
		x = z + math.Atan2(mat.Get(0, 1), mat.Get(0, 2))
	} else {
		y = -math.Asin(r31)
		cos_y := math.Cos(y)
		x = math.Atan2(mat.Get(2, 1)/cos_y, mat.Get(2, 2)/cos_y)
		z = math.Atan2(mat.Get(1, 0)/cos_y, mat.Get(0, 0)/cos_y)

		m01, m10 := mat.Get(0, 1), mat.Get(1, 0)
		m02, m20 := mat.Get(0, 2), mat.Get(2, 0)
		m12, m21 := mat.Get(1, 2), mat.Get(2, 1)
		if closeEq(math.Abs(m01-m10), 0, epsilon) &&
			closeEq(math.Abs(m02-m20), 0, epsilon) &&
			closeEq(math.Abs(m12-m21), 0, epsilon) {
			// singularity check
			// Checking for cases in which the angle is either 0 or 180

			if !mat.IsIdentity() {
				// If the angle is 0, then the rotation matrix will be the identity matrix
				// A 0 angle means that there is an arbitrary axis.

				y = math.Pi - y
				cos_y := math.Cos(y)
				x = math.Atan2(mat.Get(2, 1)/cos_y, mat.Get(2, 2)/cos_y)
				z = math.Atan2(mat.Get(1, 0)/cos_y, mat.Get(0, 0)/cos_y)
			}
		}
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

func IsRotationMatrix(m *Mat4) bool {
	return closeEq(m.Determinant(), 1, epsilon) && m.Mult(m.Transpose()).IsIdentity()
}

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
func MultVec3Mat4(v *Vec3, m *Mat4) *Vec3 {
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
