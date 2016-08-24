// Package phaser Automatic generation for Phaser.Easing.Circular
// generated file EasingCircular.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingCircular Circular easing.
type EasingCircular struct {
    *js.Object
}

// NewEasingCircular Circular easing.
func NewEasingCircular() *EasingCircular {
    return &EasingCircular{js.Global.Get("Phaser").Get("Easing").Get("Circular").New()}
}
// NewEasingCircularI Circular easing.
func NewEasingCircularI(args ...interface{}) *EasingCircular {
    return &EasingCircular{js.Global.Get("Phaser").Get("Easing").Get("Circular").New(args)}
}




// In Circular ease-in.
func (self *EasingCircular) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// InI Circular ease-in.
func (self *EasingCircular) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Out Circular ease-out.
func (self *EasingCircular) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// OutI Circular ease-out.
func (self *EasingCircular) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// InOut Circular ease-in/out.
func (self *EasingCircular) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// InOutI Circular ease-in/out.
func (self *EasingCircular) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}

