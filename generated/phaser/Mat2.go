// Automatic generation for mat2
// generated file Mat2.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type Mat2 struct {
    *js.Object
}


// 
func NewMat2() *Mat2 {
    return &Mat2{js.Global.Call("mat2")}
}

// 
func NewMat2I(args ...interface{}) *Mat2 {
    return &Mat2{js.Global.Call("mat2", args)}
}




