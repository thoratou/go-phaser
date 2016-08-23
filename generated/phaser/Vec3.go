// Automatic generation for vec3
// generated file Vec3.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type Vec3 struct {
    *js.Object
}


// 
func NewVec3() *Vec3 {
    return &Vec3{js.Global.Get("vec3").New()}
}

// 
func NewVec3I(args ...interface{}) *Vec3 {
    return &Vec3{js.Global.Get("vec3").New(args)}
}




