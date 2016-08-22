// Automatic generation for PIXI.WebGLStencilManager
// generated file WebGLStencilManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type WebGLStencilManager struct {
    *js.Object
}


// 
func NewWebGLStencilManager() *WebGLStencilManager {
    return &WebGLStencilManager{js.Global.Call("PIXI.WebGLStencilManager")}
}

// 
func NewWebGLStencilManagerI(args ...interface{}) *WebGLStencilManager {
    return &WebGLStencilManager{js.Global.Call("PIXI.WebGLStencilManager", args)}
}





// Sets the drawing context to the one given in parameter.
func (self *WebGLStencilManager) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// Sets the drawing context to the one given in parameter.
func (self *WebGLStencilManager) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// Applies the Mask and adds it to the current filter stack.
func (self *WebGLStencilManager) PushMask(graphics *Graphics, webGLData []interface{}, renderSession interface{}) {
    self.Object.Call("pushMask", graphics, webGLData, renderSession)
}

// Applies the Mask and adds it to the current filter stack.
func (self *WebGLStencilManager) PushMaskI(args ...interface{}) {
    self.Object.Call("pushMask", args)
}

// TODO this does not belong here!
func (self *WebGLStencilManager) BindGraphics(graphics *Graphics, webGLData []interface{}, renderSession interface{}) {
    self.Object.Call("bindGraphics", graphics, webGLData, renderSession)
}

// TODO this does not belong here!
func (self *WebGLStencilManager) BindGraphicsI(args ...interface{}) {
    self.Object.Call("bindGraphics", args)
}

// 
func (self *WebGLStencilManager) PopStencil(graphics *Graphics, webGLData []interface{}, renderSession interface{}) {
    self.Object.Call("popStencil", graphics, webGLData, renderSession)
}

// 
func (self *WebGLStencilManager) PopStencilI(args ...interface{}) {
    self.Object.Call("popStencil", args)
}

// Destroys the mask stack.
func (self *WebGLStencilManager) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the mask stack.
func (self *WebGLStencilManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
