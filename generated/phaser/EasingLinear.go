// Automatic generation for Phaser.Easing.Linear
// generated file EasingLinear.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Linear easing.
type EasingLinear struct {
    *js.Object
}




// Linear Easing (no variation).
func (self *EasingLinear) NoneI(args ...interface{}) float64{
    return self.Call("None", args).Float()
}
