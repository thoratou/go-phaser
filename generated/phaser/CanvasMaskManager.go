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


// A set of functions used to handle masking.
func NewCanvasMaskManager() *CanvasMaskManager {
    return &CanvasMaskManager{js.Global.Call("PIXI.CanvasMaskManager")}
}

// A set of functions used to handle masking.
func NewCanvasMaskManagerI(args ...interface{}) *CanvasMaskManager {
    return &CanvasMaskManager{js.Global.Call("PIXI.CanvasMaskManager", args)}
}





// This method adds it to the current stack of masks.
func (self *CanvasMaskManager) PushMask(maskData interface{}, renderSession interface{}) {
    self.Object.Call("pushMask", maskData, renderSession)
}

// This method adds it to the current stack of masks.
func (self *CanvasMaskManager) PushMaskI(args ...interface{}) {
    self.Object.Call("pushMask", args)
}

// Restores the current drawing context to the state it was before the mask was applied.
func (self *CanvasMaskManager) PopMask(renderSession interface{}) {
    self.Object.Call("popMask", renderSession)
}

// Restores the current drawing context to the state it was before the mask was applied.
func (self *CanvasMaskManager) PopMaskI(args ...interface{}) {
    self.Object.Call("popMask", args)
}
