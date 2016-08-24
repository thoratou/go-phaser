// Package phaser Automatic generation for mat4
// generated file Mat4.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Mat4 empty description
type Mat4 struct {
    *js.Object
}

// NewMat4 empty description
func NewMat4() *Mat4 {
    return &Mat4{js.Global.Get("mat4").New()}
}
// NewMat4I empty description
func NewMat4I(args ...interface{}) *Mat4 {
    return &Mat4{js.Global.Get("mat4").New(args)}
}




