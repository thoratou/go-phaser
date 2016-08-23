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
    return &Mat2{js.Global.Get("mat2").New()}
}

// 
func NewMat2I(args ...interface{}) *Mat2 {
    return &Mat2{js.Global.Get("mat2").New(args)}
}




