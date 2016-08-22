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
func (self *EasingCubic) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// Cubic ease-in.
func (self *EasingCubic) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Cubic ease-out.
func (self *EasingCubic) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// Cubic ease-out.
func (self *EasingCubic) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// Cubic ease-in/out.
func (self *EasingCubic) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// Cubic ease-in/out.
func (self *EasingCubic) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}
