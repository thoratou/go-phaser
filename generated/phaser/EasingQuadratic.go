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
func (self *EasingQuadratic) InI(args ...interface{}) float64{
    return self.Call("In", args).Float()
}

// Ease-out.
func (self *EasingQuadratic) OutI(args ...interface{}) float64{
    return self.Call("Out", args).Float()
}

// Ease-in/out.
func (self *EasingQuadratic) InOutI(args ...interface{}) float64{
    return self.Call("InOut", args).Float()
}
