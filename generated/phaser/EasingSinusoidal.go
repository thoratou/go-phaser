// Package phaser Automatic generation for Phaser.Easing.Sinusoidal
// generated file EasingSinusoidal.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingSinusoidal Sinusoidal easing.
type EasingSinusoidal struct {
    *js.Object
}

// NewEasingSinusoidal Sinusoidal easing.
func NewEasingSinusoidal() *EasingSinusoidal {
    return &EasingSinusoidal{js.Global.Get("Phaser").Get("Easing").Get("Sinusoidal").New()}
}
// NewEasingSinusoidalI Sinusoidal easing.
func NewEasingSinusoidalI(args ...interface{}) *EasingSinusoidal {
    return &EasingSinusoidal{js.Global.Get("Phaser").Get("Easing").Get("Sinusoidal").New(args)}
}




// In Sinusoidal ease-in.
func (self *EasingSinusoidal) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// InI Sinusoidal ease-in.
func (self *EasingSinusoidal) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Out Sinusoidal ease-out.
func (self *EasingSinusoidal) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// OutI Sinusoidal ease-out.
func (self *EasingSinusoidal) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// InOut Sinusoidal ease-in/out.
func (self *EasingSinusoidal) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// InOutI Sinusoidal ease-in/out.
func (self *EasingSinusoidal) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}

