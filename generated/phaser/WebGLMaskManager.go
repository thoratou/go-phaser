// Automatic generation for PIXI.WebGLMaskManager
// generated file WebGLMaskManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type WebGLMaskManager struct {
    *js.Object
}




// Sets the drawing context to the one given in parameter.
func (self *WebGLMaskManager) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// Sets the drawing context to the one given in parameter.
func (self *WebGLMaskManager) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// Applies the Mask and adds it to the current filter stack.
func (self *WebGLMaskManager) PushMask(maskData []interface{}, renderSession interface{}) {
    self.Object.Call("pushMask", maskData, renderSession)
}

// Applies the Mask and adds it to the current filter stack.
func (self *WebGLMaskManager) PushMaskI(args ...interface{}) {
    self.Object.Call("pushMask", args)
}

// Removes the last filter from the filter stack and doesn't return it.
func (self *WebGLMaskManager) PopMask(maskData []interface{}, renderSession interface{}) {
    self.Object.Call("popMask", maskData, renderSession)
}

// Removes the last filter from the filter stack and doesn't return it.
func (self *WebGLMaskManager) PopMaskI(args ...interface{}) {
    self.Object.Call("popMask", args)
}

// Destroys the mask stack.
func (self *WebGLMaskManager) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the mask stack.
func (self *WebGLMaskManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
