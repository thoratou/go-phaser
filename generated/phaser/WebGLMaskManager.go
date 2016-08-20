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
func (self *WebGLMaskManager) SetContextI(args ...interface{}) {
    self.Call("setContext", args)
}

// Applies the Mask and adds it to the current filter stack.
func (self *WebGLMaskManager) PushMaskI(args ...interface{}) {
    self.Call("pushMask", args)
}

// Removes the last filter from the filter stack and doesn't return it.
func (self *WebGLMaskManager) PopMaskI(args ...interface{}) {
    self.Call("popMask", args)
}

// Destroys the mask stack.
func (self *WebGLMaskManager) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
