// Package phaser Automatic generation for PIXI.WebGLBlendModeManager
// generated file WebGLBlendModeManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// WebGLBlendModeManager empty description
type WebGLBlendModeManager struct {
    *js.Object
}

// NewWebGLBlendModeManager empty description
func NewWebGLBlendModeManager(gl *WebGLContext) *WebGLBlendModeManager {
    return &WebGLBlendModeManager{js.Global.Get("PIXI").Get("WebGLBlendModeManager").New(gl)}
}
// NewWebGLBlendModeManagerI empty description
func NewWebGLBlendModeManagerI(args ...interface{}) *WebGLBlendModeManager {
    return &WebGLBlendModeManager{js.Global.Get("PIXI").Get("WebGLBlendModeManager").New(args)}
}



// CurrentBlendMode empty description
func (self *WebGLBlendModeManager) CurrentBlendMode() int{
    return self.Object.Get("currentBlendMode").Int()
}

// SetCurrentBlendModeA empty description
func (self *WebGLBlendModeManager) SetCurrentBlendModeA(member int) {
    self.Object.Set("currentBlendMode", member)
}


// SetContext Sets the WebGL Context.
func (self *WebGLBlendModeManager) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// SetContextI Sets the WebGL Context.
func (self *WebGLBlendModeManager) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// SetBlendMode Sets-up the given blendMode from WebGL's point of view.
func (self *WebGLBlendModeManager) SetBlendMode(blendMode int) {
    self.Object.Call("setBlendMode", blendMode)
}

// SetBlendModeI Sets-up the given blendMode from WebGL's point of view.
func (self *WebGLBlendModeManager) SetBlendModeI(args ...interface{}) {
    self.Object.Call("setBlendMode", args)
}

// Destroy Destroys this object.
func (self *WebGLBlendModeManager) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this object.
func (self *WebGLBlendModeManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

