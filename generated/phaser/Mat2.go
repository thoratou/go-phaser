// Package phaser Automatic generation for mat2
// generated file Mat2.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Mat2 empty description
type Mat2 struct {
    *js.Object
}

// NewMat2 empty description
func NewMat2() *Mat2 {
    return &Mat2{js.Global.Get("mat2").New()}
}
// NewMat2I empty description
func NewMat2I(args ...interface{}) *Mat2 {
    return &Mat2{js.Global.Get("mat2").New(args)}
}



// Mat2 Binding conversion method to Mat2 point 
func ToMat2(jsStruct interface{}) *Mat2 {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Mat2{Object: object}
	}
	return nil
}




