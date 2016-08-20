// Automatic generation for Phaser.Easing.Circular
// generated file EasingCircular.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Circular easing.
type EasingCircular struct {
    *js.Object
}




// Circular ease-in.
func (self *EasingCircular) InI(args ...interface{}) float64{
    return self.Call("In", args).Float()
}

// Circular ease-out.
func (self *EasingCircular) OutI(args ...interface{}) float64{
    return self.Call("Out", args).Float()
}

// Circular ease-in/out.
func (self *EasingCircular) InOutI(args ...interface{}) float64{
    return self.Call("InOut", args).Float()
}
