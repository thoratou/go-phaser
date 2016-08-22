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


// Back easing.
func NewEasingBack() *EasingBack {
    return &EasingBack{js.Global.Call("Phaser.Easing.Back")}
}

// Back easing.
func NewEasingBackI(args ...interface{}) *EasingBack {
    return &EasingBack{js.Global.Call("Phaser.Easing.Back", args)}
}





// Back ease-in.
func (self *EasingBack) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// Back ease-in.
func (self *EasingBack) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Back ease-out.
func (self *EasingBack) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// Back ease-out.
func (self *EasingBack) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// Back ease-in/out.
func (self *EasingBack) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// Back ease-in/out.
func (self *EasingBack) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}
