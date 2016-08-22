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


// Linear easing.
func NewEasingLinear() *EasingLinear {
    return &EasingLinear{js.Global.Call("Phaser.Easing.Linear")}
}

// Linear easing.
func NewEasingLinearI(args ...interface{}) *EasingLinear {
    return &EasingLinear{js.Global.Call("Phaser.Easing.Linear", args)}
}





// Linear Easing (no variation).
func (self *EasingLinear) None(k int) int{
    return self.Object.Call("None", k).Int()
}

// Linear Easing (no variation).
func (self *EasingLinear) NoneI(args ...interface{}) int{
    return self.Object.Call("None", args).Int()
}
