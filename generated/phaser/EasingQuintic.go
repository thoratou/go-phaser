// Automatic generation for Phaser.Easing.Quintic
// generated file EasingQuintic.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Quintic easing.
type EasingQuintic struct {
    *js.Object
}


// Quintic easing.
func NewEasingQuintic() *EasingQuintic {
    return &EasingQuintic{js.Global.Call("Phaser.Easing.Quintic")}
}

// Quintic easing.
func NewEasingQuinticI(args ...interface{}) *EasingQuintic {
    return &EasingQuintic{js.Global.Call("Phaser.Easing.Quintic", args)}
}





// Quintic ease-in.
func (self *EasingQuintic) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// Quintic ease-in.
func (self *EasingQuintic) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Quintic ease-out.
func (self *EasingQuintic) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// Quintic ease-out.
func (self *EasingQuintic) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// Quintic ease-in/out.
func (self *EasingQuintic) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// Quintic ease-in/out.
func (self *EasingQuintic) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}
