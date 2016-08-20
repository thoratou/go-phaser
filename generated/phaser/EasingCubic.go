// Automatic generation for Phaser.Easing.Cubic
// generated file EasingCubic.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Cubic easing.
type EasingCubic struct {
    *js.Object
}




// Cubic ease-in.
func (self *EasingCubic) InI(args ...interface{}) float64{
    return self.Call("In", args).Float()
}

// Cubic ease-out.
func (self *EasingCubic) OutI(args ...interface{}) float64{
    return self.Call("Out", args).Float()
}

// Cubic ease-in/out.
func (self *EasingCubic) InOutI(args ...interface{}) float64{
    return self.Call("InOut", args).Float()
}
