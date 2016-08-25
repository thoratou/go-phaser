// Package phaser Automatic generation for Phaser.Easing.Linear
// generated file EasingLinear.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingLinear Linear easing.
type EasingLinear struct {
    *js.Object
}

// NewEasingLinear Linear easing.
func NewEasingLinear() *EasingLinear {
    return &EasingLinear{js.Global.Get("Phaser").Get("Easing").Get("Linear").New()}
}
// NewEasingLinearI Linear easing.
func NewEasingLinearI(args ...interface{}) *EasingLinear {
    return &EasingLinear{js.Global.Get("Phaser").Get("Easing").Get("Linear").New(args)}
}



// EasingLinear Binding conversion method to EasingLinear point 
func ToEasingLinear(jsStruct interface{}) *EasingLinear {
    if object, ok := jsStruct.(*js.Object); ok {
		return &EasingLinear{Object: object}
	}
	return nil
}




// None Linear Easing (no variation).
func (self *EasingLinear) None(k int) int{
    return self.Object.Call("None", k).Int()
}

// NoneI Linear Easing (no variation).
func (self *EasingLinear) NoneI(args ...interface{}) int{
    return self.Object.Call("None", args).Int()
}

