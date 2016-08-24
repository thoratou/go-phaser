// Package phaser Automatic generation for Phaser.Easing.Quadratic
// generated file EasingQuadratic.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingQuadratic Quadratic easing.
type EasingQuadratic struct {
    *js.Object
}

// NewEasingQuadratic Quadratic easing.
func NewEasingQuadratic() *EasingQuadratic {
    return &EasingQuadratic{js.Global.Get("Phaser").Get("Easing").Get("Quadratic").New()}
}
// NewEasingQuadraticI Quadratic easing.
func NewEasingQuadraticI(args ...interface{}) *EasingQuadratic {
    return &EasingQuadratic{js.Global.Get("Phaser").Get("Easing").Get("Quadratic").New(args)}
}




// In Ease-in.
func (self *EasingQuadratic) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// InI Ease-in.
func (self *EasingQuadratic) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Out Ease-out.
func (self *EasingQuadratic) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// OutI Ease-out.
func (self *EasingQuadratic) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// InOut Ease-in/out.
func (self *EasingQuadratic) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// InOutI Ease-in/out.
func (self *EasingQuadratic) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}

