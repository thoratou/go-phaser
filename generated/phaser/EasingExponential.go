// Package phaser Automatic generation for Phaser.Easing.Exponential
// generated file EasingExponential.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingExponential Exponential easing.
type EasingExponential struct {
    *js.Object
}

// NewEasingExponential Exponential easing.
func NewEasingExponential() *EasingExponential {
    return &EasingExponential{js.Global.Get("Phaser").Get("Easing").Get("Exponential").New()}
}
// NewEasingExponentialI Exponential easing.
func NewEasingExponentialI(args ...interface{}) *EasingExponential {
    return &EasingExponential{js.Global.Get("Phaser").Get("Easing").Get("Exponential").New(args)}
}



// EasingExponential Binding conversion method to EasingExponential point 
func ToEasingExponential(jsStruct interface{}) *EasingExponential {
    if object, ok := jsStruct.(*js.Object); ok {
		return &EasingExponential{Object: object}
	}
	return nil
}




// In Exponential ease-in.
func (self *EasingExponential) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// InI Exponential ease-in.
func (self *EasingExponential) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Out Exponential ease-out.
func (self *EasingExponential) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// OutI Exponential ease-out.
func (self *EasingExponential) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// InOut Exponential ease-in/out.
func (self *EasingExponential) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// InOutI Exponential ease-in/out.
func (self *EasingExponential) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}

