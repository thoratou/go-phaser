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


// Bounce easing.
func NewEasingBounce() *EasingBounce {
    return &EasingBounce{js.Global.Call("Phaser.Easing.Bounce")}
}

// Bounce easing.
func NewEasingBounceI(args ...interface{}) *EasingBounce {
    return &EasingBounce{js.Global.Call("Phaser.Easing.Bounce", args)}
}





// Bounce ease-in.
func (self *EasingBounce) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// Bounce ease-in.
func (self *EasingBounce) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Bounce ease-out.
func (self *EasingBounce) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// Bounce ease-out.
func (self *EasingBounce) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// Bounce ease-in/out.
func (self *EasingBounce) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// Bounce ease-in/out.
func (self *EasingBounce) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}
