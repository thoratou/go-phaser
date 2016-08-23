// Automatic generation for vec2
// generated file Vec2.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type Vec2 struct {
    *js.Object
}


// 
func NewVec2() *Vec2 {
    return &Vec2{js.Global.Get("vec2").New()}
}

// 
func NewVec2I(args ...interface{}) *Vec2 {
    return &Vec2{js.Global.Get("vec2").New(args)}
}




