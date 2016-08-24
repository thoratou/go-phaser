// Package phaser Automatic generation for PIXI.WebGLStencilManager
// generated file WebGLStencilManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// WebGLStencilManager empty description
type WebGLStencilManager struct {
    *js.Object
}

// NewWebGLStencilManager empty description
func NewWebGLStencilManager() *WebGLStencilManager {
    return &WebGLStencilManager{js.Global.Get("PIXI").Get("WebGLStencilManager").New()}
}
// NewWebGLStencilManagerI empty description
func NewWebGLStencilManagerI(args ...interface{}) *WebGLStencilManager {
    return &WebGLStencilManager{js.Global.Get("PIXI").Get("WebGLStencilManager").New(args)}
}




// SetContext Sets the drawing context to the one given in parameter.
func (self *WebGLStencilManager) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// SetContextI Sets the drawing context to the one given in parameter.
func (self *WebGLStencilManager) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// PushMask Applies the Mask and adds it to the current filter stack.
func (self *WebGLStencilManager) PushMask(graphics *Graphics, webGLData []interface{}, renderSession interface{}) {
    self.Object.Call("pushMask", graphics, webGLData, renderSession)
}

// PushMaskI Applies the Mask and adds it to the current filter stack.
func (self *WebGLStencilManager) PushMaskI(args ...interface{}) {
    self.Object.Call("pushMask", args)
}

// BindGraphics TODO this does not belong here!
func (self *WebGLStencilManager) BindGraphics(graphics *Graphics, webGLData []interface{}, renderSession interface{}) {
    self.Object.Call("bindGraphics", graphics, webGLData, renderSession)
}

// BindGraphicsI TODO this does not belong here!
func (self *WebGLStencilManager) BindGraphicsI(args ...interface{}) {
    self.Object.Call("bindGraphics", args)
}

// PopStencil empty description
func (self *WebGLStencilManager) PopStencil(graphics *Graphics, webGLData []interface{}, renderSession interface{}) {
    self.Object.Call("popStencil", graphics, webGLData, renderSession)
}

// PopStencilI empty description
func (self *WebGLStencilManager) PopStencilI(args ...interface{}) {
    self.Object.Call("popStencil", args)
}

// Destroy Destroys the mask stack.
func (self *WebGLStencilManager) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the mask stack.
func (self *WebGLStencilManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

