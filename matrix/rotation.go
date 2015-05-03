package matrix

import (
// "math"
)

// // axis angle
// // quaterion
// // euler angles
// // matrix

func A() {

}

// // @param angle - In Radians
// // @param v Vec3 - unit vector of the direction
// func AxisAngleToQ8n(angle float64, x, y, z float64) *q8n {
// 	return &q8n{
// 		math.Cos(angle/2),
// 		math.Sin(angle/2) * x,
// 		math.Sin(angle/2) * y,
// 		math.Sin(angle/2) * z,
// 	}
// }

// // @param x,y,z euler angles (radians) around the x,y and z axis
// func EulerToQ8n(pitch, yaw, roll float64) *q8n {
// 	// A := math.Cos(yaw/2)
// 	// B := math.Sin(yaw/2)
// 	// C := math.Cos(pitch/2)
// 	// D := math.Sin(pitch/2)
// 	// E := math.Cos(roll/2)
// 	// F := math.Sin(roll/2)

// 	// out := &q8n
// 	// out.w = A*C*E - B*D*F
// 	// out.x = B*C*E + A*D*F
// 	// out.y = A*D*E - B*C*F
// 	// out.z = A*C*F + B*D*E
// 	// return

// 	pitchQ := &q8n{math.Cos(pitch / 2), math.Sin(pitch / 2), 0, 0}
// 	yawQ := &q8n{math.Cos(yaw / 2), 0, math.Sin(yaw / 2), 0}
// 	rollQ := &q8n{math.Cos(roll / 2), 0, 0, math.Sin(roll / 2)}
// 	return yawQ.MultIn(pitchQ).MultIn(rollQ)
// }

// func Mat44ToQ8n(v [16]float64) *q8n {

// 	return &q8n{}
// }

// func Q8nToAxisAngle(q *q8n) (angle, x, y, z float64) {
// 	angle := math.Acos(2*q.w)
// 	s = math.Sin(angle/2)
// 	x = q.x/s
// 	y = q.y/s
// 	z = q.z/s
// 	return
// }
// func Q8nToEuler(q *q8n) (x, y, z float64){

//  	return
// }

// // 	    1 - 2y² - 2z² 		2xy + 2wz 			2xz - 2wy
// // M =  2xy - 2wz 			1 - 2x² - 2z² 		2yz + 2wx
// //  	2xz + 2wy 			2yz - 2wx 			1 - 2x² - 2y²
// // xy := q.x*q.y  wz := q.w*q.z  xz := q.x*q.z
// // wy := q.w*q.y  yz := q.y*q.z  wx := q.w*q.x
// // xx := q.x*q.x  zz := q.z*q.z  yy := q.y*q.y
// func Q8nToMat44(q *q8n) (mat *Mat4){
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

// func AxisAngleToMat44(angle float64, x, y, z float64) (mat [16]float64) {

// 	return
// }
// func EulerToMat44(x, y, z float64) (mat *Mat4) {
// 	a := math.Cos(x)
// 	b := math.Sin(x)
// 	c := math.Cos(y)
// 	d := math.Sin(y)
// 	e := math.Cos(z)
// 	f := math.Sin(z)

// 	ad := a * d
// 	bd := b * d

// 	mat.mat[0] = c * e
// 	mat.mat[1] = -c * f
// 	mat.mat[2] = -d
// 	mat.mat[4] = -bd*e + a*f
// 	mat.mat[5] = bd*f + a*e
// 	mat.mat[6] = -b * c
// 	mat.mat[8] = ad*e + b*f
// 	mat.mat[9] = -ad*f + b*e
// 	mat.mat[10] = a * c

// 	mat.mat[3] = 0
// 	mat.mat[7] = 0
// 	mat.mat[11] = 0
// 	mat.mat[12] = 0
// 	mat.mat[13] = 0
// 	mat.mat[14] = 0
// 	mat.mat[15] = 1
// 	return
// }
// func Mat44ToAxisAngle(arr [16]float64) (angle, x, y, z float64) {
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
