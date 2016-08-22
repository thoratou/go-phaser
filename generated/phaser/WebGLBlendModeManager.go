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
func (self *WebGLBlendModeManager) GetCurrentBlendModeA() int{
    return self.Object.Get("currentBlendMode").Int()
}

// 
func (self *WebGLBlendModeManager) SetCurrentBlendModeA(member int) {
    self.Object.Set("currentBlendMode", member)
}



// Sets the WebGL Context.
func (self *WebGLBlendModeManager) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// Sets the WebGL Context.
func (self *WebGLBlendModeManager) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// Sets-up the given blendMode from WebGL's point of view.
func (self *WebGLBlendModeManager) SetBlendMode(blendMode int) {
    self.Object.Call("setBlendMode", blendMode)
}

// Sets-up the given blendMode from WebGL's point of view.
func (self *WebGLBlendModeManager) SetBlendModeI(args ...interface{}) {
    self.Object.Call("setBlendMode", args)
}

// Destroys this object.
func (self *WebGLBlendModeManager) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this object.
func (self *WebGLBlendModeManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
