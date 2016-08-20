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
func (self *EasingSinusoidal) InI(args ...interface{}) float64{
    return self.Call("In", args).Float()
}

// Sinusoidal ease-out.
func (self *EasingSinusoidal) OutI(args ...interface{}) float64{
    return self.Call("Out", args).Float()
}

// Sinusoidal ease-in/out.
func (self *EasingSinusoidal) InOutI(args ...interface{}) float64{
    return self.Call("InOut", args).Float()
}
