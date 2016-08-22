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
func (self *EasingQuintic) InI(args ...interface{}) int{
    return self.Call("In", args).Int()
}

// Quintic ease-out.
func (self *EasingQuintic) OutI(args ...interface{}) int{
    return self.Call("Out", args).Int()
}

// Quintic ease-in/out.
func (self *EasingQuintic) InOutI(args ...interface{}) int{
    return self.Call("InOut", args).Int()
}
