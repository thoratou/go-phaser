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
func (self *EasingQuadratic) InI(args ...interface{}) int{
    return self.Call("In", args).Int()
}

// Ease-out.
func (self *EasingQuadratic) OutI(args ...interface{}) int{
    return self.Call("Out", args).Int()
}

// Ease-in/out.
func (self *EasingQuadratic) InOutI(args ...interface{}) int{
    return self.Call("InOut", args).Int()
}
