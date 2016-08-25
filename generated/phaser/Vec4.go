// Package phaser Automatic generation for vec4
// generated file Vec4.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Vec4 empty description
type Vec4 struct {
    *js.Object
}

// NewVec4 empty description
func NewVec4() *Vec4 {
    return &Vec4{js.Global.Get("vec4").New()}
}
// NewVec4I empty description
func NewVec4I(args ...interface{}) *Vec4 {
    return &Vec4{js.Global.Get("vec4").New(args)}
}



// Vec4 Binding conversion method to Vec4 point 
func ToVec4(jsStruct interface{}) *Vec4 {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Vec4{Object: object}
	}
	return nil
}




