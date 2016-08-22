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




// Elastic ease-in.
func (self *EasingElastic) InI(args ...interface{}) int{
    return self.Call("In", args).Int()
}

// Elastic ease-out.
func (self *EasingElastic) OutI(args ...interface{}) int{
    return self.Call("Out", args).Int()
}

// Elastic ease-in/out.
func (self *EasingElastic) InOutI(args ...interface{}) int{
    return self.Call("InOut", args).Int()
}
