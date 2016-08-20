// Automatic generation for Phaser.Easing.Exponential
// generated file EasingExponential.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Exponential easing.
type EasingExponential struct {
    *js.Object
}




// Exponential ease-in.
func (self *EasingExponential) InI(args ...interface{}) float64{
    return self.Call("In", args).Float()
}

// Exponential ease-out.
func (self *EasingExponential) OutI(args ...interface{}) float64{
    return self.Call("Out", args).Float()
}

// Exponential ease-in/out.
func (self *EasingExponential) InOutI(args ...interface{}) float64{
    return self.Call("InOut", args).Float()
}
