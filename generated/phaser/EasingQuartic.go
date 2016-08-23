// Automatic generation for Phaser.Easing.Quartic
// generated file EasingQuartic.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Quartic easing.
type EasingQuartic struct {
    *js.Object
}


// Quartic easing.
func NewEasingQuartic() *EasingQuartic {
    return &EasingQuartic{js.Global.Get("Phaser").Get("Easing").Get("Quartic").New()}
}

// Quartic easing.
func NewEasingQuarticI(args ...interface{}) *EasingQuartic {
    return &EasingQuartic{js.Global.Get("Phaser").Get("Easing").Get("Quartic").New(args)}
}





// Quartic ease-in.
func (self *EasingQuartic) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// Quartic ease-in.
func (self *EasingQuartic) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Quartic ease-out.
func (self *EasingQuartic) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// Quartic ease-out.
func (self *EasingQuartic) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// Quartic ease-in/out.
func (self *EasingQuartic) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// Quartic ease-in/out.
func (self *EasingQuartic) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}
