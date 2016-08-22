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
func (self *EasingExponential) InI(args ...interface{}) int{
    return self.Call("In", args).Int()
}

// Exponential ease-out.
func (self *EasingExponential) OutI(args ...interface{}) int{
    return self.Call("Out", args).Int()
}

// Exponential ease-in/out.
func (self *EasingExponential) InOutI(args ...interface{}) int{
    return self.Call("InOut", args).Int()
}
