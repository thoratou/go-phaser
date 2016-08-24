// Package phaser Automatic generation for mat3
// generated file Mat3.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Mat3 empty description
type Mat3 struct {
    *js.Object
}

// NewMat3 empty description
func NewMat3() *Mat3 {
    return &Mat3{js.Global.Get("mat3").New()}
}
// NewMat3I empty description
func NewMat3I(args ...interface{}) *Mat3 {
    return &Mat3{js.Global.Get("mat3").New(args)}
}




