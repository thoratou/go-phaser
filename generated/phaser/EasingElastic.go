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
func (self *EasingElastic) InI(args ...interface{}) float64{
    return self.Call("In", args).Float()
}

// Elastic ease-out.
func (self *EasingElastic) OutI(args ...interface{}) float64{
    return self.Call("Out", args).Float()
}

// Elastic ease-in/out.
func (self *EasingElastic) InOutI(args ...interface{}) float64{
    return self.Call("InOut", args).Float()
}
