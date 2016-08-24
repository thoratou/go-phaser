// Package phaser Automatic generation for Phaser.Easing.Back
// generated file EasingBack.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EasingBack Back easing.
type EasingBack struct {
    *js.Object
}

// NewEasingBack Back easing.
func NewEasingBack() *EasingBack {
    return &EasingBack{js.Global.Get("Phaser").Get("Easing").Get("Back").New()}
}
// NewEasingBackI Back easing.
func NewEasingBackI(args ...interface{}) *EasingBack {
    return &EasingBack{js.Global.Get("Phaser").Get("Easing").Get("Back").New(args)}
}




// In Back ease-in.
func (self *EasingBack) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// InI Back ease-in.
func (self *EasingBack) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Out Back ease-out.
func (self *EasingBack) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// OutI Back ease-out.
func (self *EasingBack) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// InOut Back ease-in/out.
func (self *EasingBack) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// InOutI Back ease-in/out.
func (self *EasingBack) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}

