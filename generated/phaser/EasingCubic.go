// Package phaser Automatic generation for Phaser.Easing.Cubic
// generated file EasingCubic.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingCubic Cubic easing.
type EasingCubic struct {
    *js.Object
}

// NewEasingCubic Cubic easing.
func NewEasingCubic() *EasingCubic {
    return &EasingCubic{js.Global.Get("Phaser").Get("Easing").Get("Cubic").New()}
}
// NewEasingCubicI Cubic easing.
func NewEasingCubicI(args ...interface{}) *EasingCubic {
    return &EasingCubic{js.Global.Get("Phaser").Get("Easing").Get("Cubic").New(args)}
}



// EasingCubic Binding conversion method to EasingCubic point 
func ToEasingCubic(jsStruct interface{}) *EasingCubic {
    if object, ok := jsStruct.(*js.Object); ok {
		return &EasingCubic{Object: object}
	}
	return nil
}




// In Cubic ease-in.
func (self *EasingCubic) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// InI Cubic ease-in.
func (self *EasingCubic) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Out Cubic ease-out.
func (self *EasingCubic) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// OutI Cubic ease-out.
func (self *EasingCubic) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// InOut Cubic ease-in/out.
func (self *EasingCubic) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// InOutI Cubic ease-in/out.
func (self *EasingCubic) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}

