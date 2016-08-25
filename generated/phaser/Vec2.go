// Package phaser Automatic generation for vec2
// generated file Vec2.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Vec2 empty description
type Vec2 struct {
    *js.Object
}

// NewVec2 empty description
func NewVec2() *Vec2 {
    return &Vec2{js.Global.Get("vec2").New()}
}
// NewVec2I empty description
func NewVec2I(args ...interface{}) *Vec2 {
    return &Vec2{js.Global.Get("vec2").New(args)}
}



// Vec2 Binding conversion method to Vec2 point 
func ToVec2(jsStruct interface{}) *Vec2 {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Vec2{Object: object}
	}
	return nil
}




