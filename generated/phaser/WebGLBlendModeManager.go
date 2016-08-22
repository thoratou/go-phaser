// Automatic generation for PIXI.WebGLBlendModeManager
// generated file WebGLBlendModeManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type WebGLBlendModeManager struct {
    *js.Object
}


// 
func (self *WebGLBlendModeManager) GetCurrentBlendMode() int{
    return self.Get("currentBlendMode").Int()
}

// 
func (self *WebGLBlendModeManager) SetCurrentBlendMode(member int) {
    self.Set("currentBlendMode", member)
}



// Sets the WebGL Context.
func (self *WebGLBlendModeManager) SetContextI(args ...interface{}) {
    self.Call("setContext", args)
}

// Sets-up the given blendMode from WebGL's point of view.
func (self *WebGLBlendModeManager) SetBlendModeI(args ...interface{}) {
    self.Call("setBlendMode", args)
}

// Destroys this object.
func (self *WebGLBlendModeManager) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
