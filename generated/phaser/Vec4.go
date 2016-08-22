// Automatic generation for vec4
// generated file Vec4.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type Vec4 struct {
    *js.Object
}


// 
func NewVec4() *Vec4 {
    return &Vec4{js.Global.Call("vec4")}
}

// 
func NewVec4I(args ...interface{}) *Vec4 {
    return &Vec4{js.Global.Call("vec4", args)}
}




