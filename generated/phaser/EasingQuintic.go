// Package phaser Automatic generation for Phaser.Easing.Quintic
// generated file EasingQuintic.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingQuintic Quintic easing.
type EasingQuintic struct {
    *js.Object
}

// NewEasingQuintic Quintic easing.
func NewEasingQuintic() *EasingQuintic {
    return &EasingQuintic{js.Global.Get("Phaser").Get("Easing").Get("Quintic").New()}
}
// NewEasingQuinticI Quintic easing.
func NewEasingQuinticI(args ...interface{}) *EasingQuintic {
    return &EasingQuintic{js.Global.Get("Phaser").Get("Easing").Get("Quintic").New(args)}
}



// EasingQuintic Binding conversion method to EasingQuintic point 
func ToEasingQuintic(jsStruct interface{}) *EasingQuintic {
    if object, ok := jsStruct.(*js.Object); ok {
		return &EasingQuintic{Object: object}
	}
	return nil
}




// In Quintic ease-in.
func (self *EasingQuintic) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// InI Quintic ease-in.
func (self *EasingQuintic) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Out Quintic ease-out.
func (self *EasingQuintic) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// OutI Quintic ease-out.
func (self *EasingQuintic) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// InOut Quintic ease-in/out.
func (self *EasingQuintic) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// InOutI Quintic ease-in/out.
func (self *EasingQuintic) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}

