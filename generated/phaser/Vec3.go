// Package phaser Automatic generation for vec3
// generated file Vec3.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Vec3 empty description
type Vec3 struct {
    *js.Object
}

// NewVec3 empty description
func NewVec3() *Vec3 {
    return &Vec3{js.Global.Get("vec3").New()}
}
// NewVec3I empty description
func NewVec3I(args ...interface{}) *Vec3 {
    return &Vec3{js.Global.Get("vec3").New(args)}
}




