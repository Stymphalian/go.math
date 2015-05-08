package lmath

import (
	// "fmt"
	// "math"
)

// Rotation functions used to create
// rotaiton matrices, quaternions from/to euler, and axis/angle
//


// // I need to test this...
func LookAtMat44(eye,center,up *Vec3) (mat *Mat4){
	forward := center.Sub(eye).NormalizeIn()
	right:= forward.Cross(up).NormalizeIn()
	up = right.Cross(forward).NormalizeIn()

	mat.SetRow(0,-forward.X,-forward.Y,-forward.Z,-eye.X)
	mat.SetRow(1,up.X,up.Y,up.Z,-eye.Y)
	mat.SetRow(2,right.X,right.Y,right.Z,-eye.Z)
	mat.SetRow(3,0,0,0,1)
	return
	// eye := &Vec3{0, 0, 0}
	// at := &Vec3{a, b, c}
	// up := &Vec3{0, 1, 0}

	// if at.Eq(&Vec3{0, 1, 0}) {
	// 	up = &Vec3{0, 0, -1}
	// } else if at.Eq(&Vec3{0, -1, 0}) {
	// 	up = &Vec3{0, 0, 1}
	// }

	// forward := at.Sub(eye).NormalizeIn()
	// right := up.Cross(forward).NormalizeIn()
	// up = forward.Cross(right).NormalizeIn()

	// mat := NewMat4(right.X, right.Y, right.Z, 0,
	// 	up.X, up.Y, up.Z, 0,
	// 	forward.X, forward.Y, forward.Z, 0,
	// 	0, 0, 0, 1)

	// rotMat := &Mat4{}
	// rotMat.ToIdentity()
	// rotMat.Set(0, 0, math.Cos(angle))
	// rotMat.Set(0, 1, -math.Sin(angle))
	// rotMat.Set(1, 0, math.Sin(angle))
	// rotMat.Set(1, 1, math.Cos(angle))

	// return mat.Transpose().MultIn(rotMat).MultIn(mat)
}
