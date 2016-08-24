// Package phaser Automatic generation for PIXI.WebGLMaskManager
// generated file WebGLMaskManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// WebGLMaskManager empty description
type WebGLMaskManager struct {
    *js.Object
}

// NewWebGLMaskManager empty description
func NewWebGLMaskManager() *WebGLMaskManager {
    return &WebGLMaskManager{js.Global.Get("PIXI").Get("WebGLMaskManager").New()}
}
// NewWebGLMaskManagerI empty description
func NewWebGLMaskManagerI(args ...interface{}) *WebGLMaskManager {
    return &WebGLMaskManager{js.Global.Get("PIXI").Get("WebGLMaskManager").New(args)}
}




// SetContext Sets the drawing context to the one given in parameter.
func (self *WebGLMaskManager) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// SetContextI Sets the drawing context to the one given in parameter.
func (self *WebGLMaskManager) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// PushMask Applies the Mask and adds it to the current filter stack.
func (self *WebGLMaskManager) PushMask(maskData []interface{}, renderSession interface{}) {
    self.Object.Call("pushMask", maskData, renderSession)
}

// PushMaskI Applies the Mask and adds it to the current filter stack.
func (self *WebGLMaskManager) PushMaskI(args ...interface{}) {
    self.Object.Call("pushMask", args)
}

// PopMask Removes the last filter from the filter stack and doesn't return it.
func (self *WebGLMaskManager) PopMask(maskData []interface{}, renderSession interface{}) {
    self.Object.Call("popMask", maskData, renderSession)
}

// PopMaskI Removes the last filter from the filter stack and doesn't return it.
func (self *WebGLMaskManager) PopMaskI(args ...interface{}) {
    self.Object.Call("popMask", args)
}

// Destroy Destroys the mask stack.
func (self *WebGLMaskManager) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the mask stack.
func (self *WebGLMaskManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

