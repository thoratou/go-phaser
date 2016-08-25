// Package phaser Automatic generation for Phaser.Easing.Bounce
// generated file EasingBounce.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingBounce Bounce easing.
type EasingBounce struct {
    *js.Object
}

// NewEasingBounce Bounce easing.
func NewEasingBounce() *EasingBounce {
    return &EasingBounce{js.Global.Get("Phaser").Get("Easing").Get("Bounce").New()}
}
// NewEasingBounceI Bounce easing.
func NewEasingBounceI(args ...interface{}) *EasingBounce {
    return &EasingBounce{js.Global.Get("Phaser").Get("Easing").Get("Bounce").New(args)}
}



// EasingBounce Binding conversion method to EasingBounce point 
func ToEasingBounce(jsStruct interface{}) *EasingBounce {
    if object, ok := jsStruct.(*js.Object); ok {
		return &EasingBounce{Object: object}
	}
	return nil
}




// In Bounce ease-in.
func (self *EasingBounce) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// InI Bounce ease-in.
func (self *EasingBounce) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Out Bounce ease-out.
func (self *EasingBounce) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// OutI Bounce ease-out.
func (self *EasingBounce) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// InOut Bounce ease-in/out.
func (self *EasingBounce) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// InOutI Bounce ease-in/out.
func (self *EasingBounce) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}

