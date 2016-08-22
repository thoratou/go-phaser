// Automatic generation for Phaser.Easing.Sinusoidal
// generated file EasingSinusoidal.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Sinusoidal easing.
type EasingSinusoidal struct {
    *js.Object
}




// Sinusoidal ease-in.
func (self *EasingSinusoidal) InI(args ...interface{}) int{
    return self.Call("In", args).Int()
}

// Sinusoidal ease-out.
func (self *EasingSinusoidal) OutI(args ...interface{}) int{
    return self.Call("Out", args).Int()
}

// Sinusoidal ease-in/out.
func (self *EasingSinusoidal) InOutI(args ...interface{}) int{
    return self.Call("InOut", args).Int()
}
