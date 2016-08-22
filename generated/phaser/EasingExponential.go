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


// Exponential easing.
func NewEasingExponential() *EasingExponential {
    return &EasingExponential{js.Global.Call("Phaser.Easing.Exponential")}
}

// Exponential easing.
func NewEasingExponentialI(args ...interface{}) *EasingExponential {
    return &EasingExponential{js.Global.Call("Phaser.Easing.Exponential", args)}
}





// Exponential ease-in.
func (self *EasingExponential) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// Exponential ease-in.
func (self *EasingExponential) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Exponential ease-out.
func (self *EasingExponential) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// Exponential ease-out.
func (self *EasingExponential) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// Exponential ease-in/out.
func (self *EasingExponential) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// Exponential ease-in/out.
func (self *EasingExponential) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}
