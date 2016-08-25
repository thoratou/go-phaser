// Package phaser Automatic generation for PIXI.CanvasMaskManager
// generated file CanvasMaskManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// CanvasMaskManager A set of functions used to handle masking.
type CanvasMaskManager struct {
    *js.Object
}

// NewCanvasMaskManager A set of functions used to handle masking.
func NewCanvasMaskManager() *CanvasMaskManager {
    return &CanvasMaskManager{js.Global.Get("PIXI").Get("CanvasMaskManager").New()}
}
// NewCanvasMaskManagerI A set of functions used to handle masking.
func NewCanvasMaskManagerI(args ...interface{}) *CanvasMaskManager {
    return &CanvasMaskManager{js.Global.Get("PIXI").Get("CanvasMaskManager").New(args)}
}



// CanvasMaskManager Binding conversion method to CanvasMaskManager point 
func ToCanvasMaskManager(jsStruct interface{}) *CanvasMaskManager {
    if object, ok := jsStruct.(*js.Object); ok {
		return &CanvasMaskManager{Object: object}
	}
	return nil
}




// PushMask This method adds it to the current stack of masks.
func (self *CanvasMaskManager) PushMask(maskData interface{}, renderSession interface{}) {
    self.Object.Call("pushMask", maskData, renderSession)
}

// PushMaskI This method adds it to the current stack of masks.
func (self *CanvasMaskManager) PushMaskI(args ...interface{}) {
    self.Object.Call("pushMask", args)
}

// PopMask Restores the current drawing context to the state it was before the mask was applied.
func (self *CanvasMaskManager) PopMask(renderSession interface{}) {
    self.Object.Call("popMask", renderSession)
}

// PopMaskI Restores the current drawing context to the state it was before the mask was applied.
func (self *CanvasMaskManager) PopMaskI(args ...interface{}) {
    self.Object.Call("popMask", args)
}

