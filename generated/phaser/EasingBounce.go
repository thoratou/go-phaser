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
func (self *EasingBounce) InI(args ...interface{}) float64{
    return self.Call("In", args).Float()
}

// Bounce ease-out.
func (self *EasingBounce) OutI(args ...interface{}) float64{
    return self.Call("Out", args).Float()
}

// Bounce ease-in/out.
func (self *EasingBounce) InOutI(args ...interface{}) float64{
    return self.Call("InOut", args).Float()
}
