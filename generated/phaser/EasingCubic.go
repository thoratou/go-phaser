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
func (self *EasingCubic) InI(args ...interface{}) int{
    return self.Call("In", args).Int()
}

// Cubic ease-out.
func (self *EasingCubic) OutI(args ...interface{}) int{
    return self.Call("Out", args).Int()
}

// Cubic ease-in/out.
func (self *EasingCubic) InOutI(args ...interface{}) int{
    return self.Call("InOut", args).Int()
}
