// Automatic generation for Phaser.Easing.Back
// generated file EasingBack.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Back easing.
type EasingBack struct {
    *js.Object
}




// Back ease-in.
func (self *EasingBack) InI(args ...interface{}) float64{
    return self.Call("In", args).Float()
}

// Back ease-out.
func (self *EasingBack) OutI(args ...interface{}) float64{
    return self.Call("Out", args).Float()
}

// Back ease-in/out.
func (self *EasingBack) InOutI(args ...interface{}) float64{
    return self.Call("InOut", args).Float()
}
