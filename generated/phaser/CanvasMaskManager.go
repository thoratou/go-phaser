// Automatic generation for PIXI.CanvasMaskManager
// generated file CanvasMaskManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A set of functions used to handle masking.
type CanvasMaskManager struct {
    *js.Object
}




// This method adds it to the current stack of masks.
func (self *CanvasMaskManager) PushMaskI(args ...interface{}) {
    self.Call("pushMask", args)
}

// Restores the current drawing context to the state it was before the mask was applied.
func (self *CanvasMaskManager) PopMaskI(args ...interface{}) {
    self.Call("popMask", args)
}
