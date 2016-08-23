// Automatic generation for PIXI.WebGLShaderManager
// generated file WebGLShaderManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type WebGLShaderManager struct {
    *js.Object
}


// 
func NewWebGLShaderManager() *WebGLShaderManager {
    return &WebGLShaderManager{js.Global.Get("PIXI").Get("WebGLShaderManager").New()}
}

// 
func NewWebGLShaderManagerI(args ...interface{}) *WebGLShaderManager {
    return &WebGLShaderManager{js.Global.Get("PIXI").Get("WebGLShaderManager").New(args)}
}



// 
func (self *WebGLShaderManager) MaxAttibs() int{
    return self.Object.Get("maxAttibs").Int()
}

// 
func (self *WebGLShaderManager) SetMaxAttibsA(member int) {
    self.Object.Set("maxAttibs", member)
}

// 
func (self *WebGLShaderManager) AttribState() []interface{}{
	array00 := self.Object.Get("attribState")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// 
func (self *WebGLShaderManager) SetAttribStateA(member []interface{}) {
    self.Object.Set("attribState", member)
}

// 
func (self *WebGLShaderManager) TempAttribState() []interface{}{
	array00 := self.Object.Get("tempAttribState")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// 
func (self *WebGLShaderManager) SetTempAttribStateA(member []interface{}) {
    self.Object.Set("tempAttribState", member)
}

// 
func (self *WebGLShaderManager) Stack() []interface{}{
	array00 := self.Object.Get("stack")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// 
func (self *WebGLShaderManager) SetStackA(member []interface{}) {
    self.Object.Set("stack", member)
}



// Initialises the context and the properties.
func (self *WebGLShaderManager) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// Initialises the context and the properties.
func (self *WebGLShaderManager) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// Takes the attributes given in parameters.
func (self *WebGLShaderManager) SetAttribs(attribs []interface{}) {
    self.Object.Call("setAttribs", attribs)
}

// Takes the attributes given in parameters.
func (self *WebGLShaderManager) SetAttribsI(args ...interface{}) {
    self.Object.Call("setAttribs", args)
}

// Sets the current shader.
func (self *WebGLShaderManager) SetShader(shader interface{}) {
    self.Object.Call("setShader", shader)
}

// Sets the current shader.
func (self *WebGLShaderManager) SetShaderI(args ...interface{}) {
    self.Object.Call("setShader", args)
}

// Destroys this object.
func (self *WebGLShaderManager) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this object.
func (self *WebGLShaderManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
