/*
lmath is a small 3D linear algebra library which provides support for
Vec3/4, Mat3/4 and Quaternions. Supports rotations using euler angle,axis-angle,
rotation matrices, quaternions as well as conversions between the representations.
*/
package lmath

// Conventions
//
//  Right-Hand coordinate system
//  All angles are specified in Radians
//  Rotations are applied in Pitch => Yaw => Roll order
var (
    // I don't like that this is var and not a const.
    Version = struct {Major, Minor,Patch int}{0,0,0}
)

// TODO:
// Test glmath functions
// Add additional tests
//  multiply Vec4 tests
