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




// Quartic ease-in.
func (self *EasingQuartic) InI(args ...interface{}) int{
    return self.Call("In", args).Int()
}

// Quartic ease-out.
func (self *EasingQuartic) OutI(args ...interface{}) int{
    return self.Call("Out", args).Int()
}

// Quartic ease-in/out.
func (self *EasingQuartic) InOutI(args ...interface{}) int{
    return self.Call("InOut", args).Int()
}
