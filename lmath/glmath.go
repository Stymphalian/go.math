package lmath

import (
	"math"
)

// The below matrices use the derivations from here:
//http://www.songho.ca/opengl/gl_projectionmatrix.html

// Creates and orthographic projection matrix from the given paramters
// precondition : far > near > 0
// near and far must be positive, and near < far
// The matrix is given in a right-handed system, but will transform
// a vector into a left-handed system (to match the NDC space of OpenGL)
func OrthoMat4(left, right, bottom, top, near, far float64) *Mat4 {
	return NewMat4(
		2/(right-left), 0, 0, -(right+left)/(right-left),
		0, 2/(top-bottom), 0, -(top+bottom)/(top-bottom),
		//0,0,2/(n-f), -(n+f)/(n-f)
		0, 0, -2/(far-near), -(far+near)/(far-near),
		0, 0, 0, 1)
}

// Create a perspective frustum matrix from the provided parameters.
// precondition : far > near > 0
// The matrix is given in a right-handed system, but will transform
// a vector into a left-handed sytsem ( to match the NDC space of OpenGL)
func FrustumMat4(left, right, bottom, top, near, far float64) *Mat4 {
	return NewMat4(
		2*near/(right-left), 0, (right+left)/(right-left), 0,
		0, 2*near/(top-bottom), (top+bottom)/(top-bottom), 0,
		// TODO: I dont' use negative hear in my glwidget impl...
		// I also take the transpose
		0, 0, -(far+near)/(far-near), -2*far*near/(far-near),
		0, 0, -1, 0)
}

// Creates a normalized viewing frustum using the given perspective parameters
// fov ( y-direction)  angle in radians
// aspect - ratio between the width and the height
// precondition: far > near > 0
func PerspectiveMat4(fov_y, aspect, near, far float64) *Mat4 {
	top := math.Atan(fov_y/2) * near
	right := top * aspect
	return FrustumMat4(-right, right, -top, top, near, far)
}

// Create a LookAt rotation matrix.
// eye is the position of the camera
// at is the position in which the camera "looksAt"
// up is the direction which is considered up. It is up to the user
// to ensure that the forward dir is not parallel to up.
func LookAtMat4(eye, at, up *Vec3) *Mat4 {
	// up := &Vec3{0, 1, 0}
	// if at.Eq(&Vec3{0, 1, 0}) {
	//  up = &Vec3{0, 0, -1}
	// } else if at.Eq(&Vec3{0, -1, 0}) {
	//  up = &Vec3{0, 0, 1}
	// }

	forward := at.Sub(eye).NormalizeIn()
	right := up.Cross(forward).NormalizeIn()
	up = forward.Cross(right).NormalizeIn()

	translate := &Mat4{}
	translate.ToTranslate(-eye.X, -eye.Y, -eye.Z)
	mat := NewMat4(right.X, right.Y, right.Z, 0,
		up.X, up.Y, up.Z, 0,
		forward.X, forward.Y, forward.Z, 0,
		0, 0, 0, 1)
	return translate.MultIn(mat)
}
