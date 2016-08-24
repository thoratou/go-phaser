// Package phaser Automatic generation for PIXI.WebGLShaderManager
// generated file WebGLShaderManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// WebGLShaderManager empty description
type WebGLShaderManager struct {
    *js.Object
}

// NewWebGLShaderManager empty description
func NewWebGLShaderManager() *WebGLShaderManager {
    return &WebGLShaderManager{js.Global.Get("PIXI").Get("WebGLShaderManager").New()}
}
// NewWebGLShaderManagerI empty description
func NewWebGLShaderManagerI(args ...interface{}) *WebGLShaderManager {
    return &WebGLShaderManager{js.Global.Get("PIXI").Get("WebGLShaderManager").New(args)}
}



// MaxAttibs empty description
func (self *WebGLShaderManager) MaxAttibs() int{
    return self.Object.Get("maxAttibs").Int()
}

// SetMaxAttibsA empty description
func (self *WebGLShaderManager) SetMaxAttibsA(member int) {
    self.Object.Set("maxAttibs", member)
}

// AttribState empty description
func (self *WebGLShaderManager) AttribState() []interface{}{
	array00 := self.Object.Get("attribState")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetAttribStateA empty description
func (self *WebGLShaderManager) SetAttribStateA(member []interface{}) {
    self.Object.Set("attribState", member)
}

// TempAttribState empty description
func (self *WebGLShaderManager) TempAttribState() []interface{}{
	array00 := self.Object.Get("tempAttribState")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetTempAttribStateA empty description
func (self *WebGLShaderManager) SetTempAttribStateA(member []interface{}) {
    self.Object.Set("tempAttribState", member)
}

// Stack empty description
func (self *WebGLShaderManager) Stack() []interface{}{
	array00 := self.Object.Get("stack")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetStackA empty description
func (self *WebGLShaderManager) SetStackA(member []interface{}) {
    self.Object.Set("stack", member)
}


// SetContext Initialises the context and the properties.
func (self *WebGLShaderManager) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// SetContextI Initialises the context and the properties.
func (self *WebGLShaderManager) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// SetAttribs Takes the attributes given in parameters.
func (self *WebGLShaderManager) SetAttribs(attribs []interface{}) {
    self.Object.Call("setAttribs", attribs)
}

// SetAttribsI Takes the attributes given in parameters.
func (self *WebGLShaderManager) SetAttribsI(args ...interface{}) {
    self.Object.Call("setAttribs", args)
}

// SetShader Sets the current shader.
func (self *WebGLShaderManager) SetShader(shader interface{}) {
    self.Object.Call("setShader", shader)
}

// SetShaderI Sets the current shader.
func (self *WebGLShaderManager) SetShaderI(args ...interface{}) {
    self.Object.Call("setShader", args)
}

// Destroy Destroys this object.
func (self *WebGLShaderManager) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this object.
func (self *WebGLShaderManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

