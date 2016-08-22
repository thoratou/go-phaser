// Automatic generation for Phaser.Easing.Quadratic
// generated file EasingQuadratic.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Quadratic easing.
type EasingQuadratic struct {
    *js.Object
}




// Ease-in.
func (self *EasingQuadratic) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// Ease-in.
func (self *EasingQuadratic) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Ease-out.
func (self *EasingQuadratic) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// Ease-out.
func (self *EasingQuadratic) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// Ease-in/out.
func (self *EasingQuadratic) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// Ease-in/out.
func (self *EasingQuadratic) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}
