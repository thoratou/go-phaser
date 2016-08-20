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




// Sets the drawing context to the one given in parameter.
func (self *WebGLStencilManager) SetContextI(args ...interface{}) {
    self.Call("setContext", args)
}

// Applies the Mask and adds it to the current filter stack.
func (self *WebGLStencilManager) PushMaskI(args ...interface{}) {
    self.Call("pushMask", args)
}

// TODO this does not belong here!
func (self *WebGLStencilManager) BindGraphicsI(args ...interface{}) {
    self.Call("bindGraphics", args)
}

// 
func (self *WebGLStencilManager) PopStencilI(args ...interface{}) {
    self.Call("popStencil", args)
}

// Destroys the mask stack.
func (self *WebGLStencilManager) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
