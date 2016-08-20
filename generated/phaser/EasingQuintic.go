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




// Quintic ease-in.
func (self *EasingQuintic) InI(args ...interface{}) float64{
    return self.Call("In", args).Float()
}

// Quintic ease-out.
func (self *EasingQuintic) OutI(args ...interface{}) float64{
    return self.Call("Out", args).Float()
}

// Quintic ease-in/out.
func (self *EasingQuintic) InOutI(args ...interface{}) float64{
    return self.Call("InOut", args).Float()
}
