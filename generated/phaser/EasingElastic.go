// Automatic generation for Phaser.Easing.Elastic
// generated file EasingElastic.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Elastic easing.
type EasingElastic struct {
    *js.Object
}


// Elastic easing.
func NewEasingElastic() *EasingElastic {
    return &EasingElastic{js.Global.Call("Phaser.Easing.Elastic")}
}

// Elastic easing.
func NewEasingElasticI(args ...interface{}) *EasingElastic {
    return &EasingElastic{js.Global.Call("Phaser.Easing.Elastic", args)}
}





// Elastic ease-in.
func (self *EasingElastic) In(k int) int{
    return self.Object.Call("In", k).Int()
}

// Elastic ease-in.
func (self *EasingElastic) InI(args ...interface{}) int{
    return self.Object.Call("In", args).Int()
}

// Elastic ease-out.
func (self *EasingElastic) Out(k int) int{
    return self.Object.Call("Out", k).Int()
}

// Elastic ease-out.
func (self *EasingElastic) OutI(args ...interface{}) int{
    return self.Object.Call("Out", args).Int()
}

// Elastic ease-in/out.
func (self *EasingElastic) InOut(k int) int{
    return self.Object.Call("InOut", k).Int()
}

// Elastic ease-in/out.
func (self *EasingElastic) InOutI(args ...interface{}) int{
    return self.Object.Call("InOut", args).Int()
}
