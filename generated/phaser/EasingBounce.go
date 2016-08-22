// Automatic generation for Phaser.Easing.Bounce
// generated file EasingBounce.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Bounce easing.
type EasingBounce struct {
    *js.Object
}




// Bounce ease-in.
func (self *EasingBounce) InI(args ...interface{}) int{
    return self.Call("In", args).Int()
}

// Bounce ease-out.
func (self *EasingBounce) OutI(args ...interface{}) int{
    return self.Call("Out", args).Int()
}

// Bounce ease-in/out.
func (self *EasingBounce) InOutI(args ...interface{}) int{
    return self.Call("InOut", args).Int()
}
