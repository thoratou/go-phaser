// Package phaser Automatic generation for Phaser.Easing.Quartic
// generated file EasingQuartic.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingQuartic Quartic easing.
type EasingQuartic struct {
    *js.Object
}

// NewEasingQuartic Quartic easing.
func NewEasingQuartic() *EasingQuartic {
    return &EasingQuartic{js.Global.Get("Phaser").Get("Easing").Get("Quartic").New()}
}
// NewEasingQuarticI Quartic easing.
func NewEasingQuarticI(args ...interface{}) *EasingQuartic {
    return &EasingQuartic{js.Global.Get("Phaser").Get("Easing").Get("Quartic").New(args)}
}




// In Quartic ease-in.
func (self *EasingQuartic) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// InI Quartic ease-in.
func (self *EasingQuartic) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Out Quartic ease-out.
func (self *EasingQuartic) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// OutI Quartic ease-out.
func (self *EasingQuartic) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// InOut Quartic ease-in/out.
func (self *EasingQuartic) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// InOutI Quartic ease-in/out.
func (self *EasingQuartic) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}

