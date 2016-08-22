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
func (self *EasingCircular) InI(args ...interface{}) int{
    return self.Call("In", args).Int()
}

// Circular ease-out.
func (self *EasingCircular) OutI(args ...interface{}) int{
    return self.Call("Out", args).Int()
}

// Circular ease-in/out.
func (self *EasingCircular) InOutI(args ...interface{}) int{
    return self.Call("InOut", args).Int()
}
