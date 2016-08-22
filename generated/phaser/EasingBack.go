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
func (self *EasingBack) InI(args ...interface{}) int{
    return self.Call("In", args).Int()
}

// Back ease-out.
func (self *EasingBack) OutI(args ...interface{}) int{
    return self.Call("Out", args).Int()
}

// Back ease-in/out.
func (self *EasingBack) InOutI(args ...interface{}) int{
    return self.Call("InOut", args).Int()
}
