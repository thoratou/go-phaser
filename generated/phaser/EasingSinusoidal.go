// Automatic generation for Phaser.Easing.Sinusoidal
// generated file EasingSinusoidal.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Sinusoidal easing.
type EasingSinusoidal struct {
    *js.Object
}


// Sinusoidal easing.
func NewEasingSinusoidal() *EasingSinusoidal {
    return &EasingSinusoidal{js.Global.Get("Phaser").Get("Easing").Get("Sinusoidal").New()}
}

// Sinusoidal easing.
func NewEasingSinusoidalI(args ...interface{}) *EasingSinusoidal {
    return &EasingSinusoidal{js.Global.Get("Phaser").Get("Easing").Get("Sinusoidal").New(args)}
}





// Sinusoidal ease-in.
func (self *EasingSinusoidal) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// Sinusoidal ease-in.
func (self *EasingSinusoidal) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Sinusoidal ease-out.
func (self *EasingSinusoidal) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// Sinusoidal ease-out.
func (self *EasingSinusoidal) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// Sinusoidal ease-in/out.
func (self *EasingSinusoidal) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// Sinusoidal ease-in/out.
func (self *EasingSinusoidal) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}
