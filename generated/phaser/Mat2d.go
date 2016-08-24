// Package phaser Automatic generation for mat2d
// generated file Mat2d.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Mat2d A mat2d contains six elements defined as:
// <pre>
// [a, c, tx,
//  b, d, ty]
// </pre>
// This is a short form for the 3x3 matrix:
// <pre>
// [a, c, tx,
//  b, d, ty,
//  0, 0, 1]
// </pre>
// The last row is ignored so the array is shorter and operations are faster.
type Mat2d struct {
    *js.Object
}

// NewMat2d A mat2d contains six elements defined as:
// <pre>
// [a, c, tx,
//  b, d, ty]
// </pre>
// This is a short form for the 3x3 matrix:
// <pre>
// [a, c, tx,
//  b, d, ty,
//  0, 0, 1]
// </pre>
// The last row is ignored so the array is shorter and operations are faster.
func NewMat2d() *Mat2d {
    return &Mat2d{js.Global.Get("mat2d").New()}
}
// NewMat2dI A mat2d contains six elements defined as:
// <pre>
// [a, c, tx,
//  b, d, ty]
// </pre>
// This is a short form for the 3x3 matrix:
// <pre>
// [a, c, tx,
//  b, d, ty,
//  0, 0, 1]
// </pre>
// The last row is ignored so the array is shorter and operations are faster.
func NewMat2dI(args ...interface{}) *Mat2d {
    return &Mat2d{js.Global.Get("mat2d").New(args)}
}




