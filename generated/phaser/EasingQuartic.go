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
func (self *EasingQuartic) InI(args ...interface{}) float64{
    return self.Call("In", args).Float()
}

// Quartic ease-out.
func (self *EasingQuartic) OutI(args ...interface{}) float64{
    return self.Call("Out", args).Float()
}

// Quartic ease-in/out.
func (self *EasingQuartic) InOutI(args ...interface{}) float64{
    return self.Call("InOut", args).Float()
}
