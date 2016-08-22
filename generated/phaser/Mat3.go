// Automatic generation for mat3
// generated file Mat3.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type Mat3 struct {
    *js.Object
}


// 
func NewMat3() *Mat3 {
    return &Mat3{js.Global.Call("mat3")}
}

// 
func NewMat3I(args ...interface{}) *Mat3 {
    return &Mat3{js.Global.Call("mat3", args)}
}




