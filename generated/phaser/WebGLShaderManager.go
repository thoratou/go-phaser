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
func (self *WebGLShaderManager) GetMaxAttibs() float64{
    return self.Get("maxAttibs").Float()
}

// 
func (self *WebGLShaderManager) SetMaxAttibs(member float64) {
    self.Set("maxAttibs", member)
}

// 
func (self *WebGLShaderManager) GetAttribState() []interface{}{
	array := self.Get("attribState")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// 
func (self *WebGLShaderManager) SetAttribState(member []interface{}) {
    self.Set("attribState", member)
}

// 
func (self *WebGLShaderManager) GetTempAttribState() []interface{}{
	array := self.Get("tempAttribState")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// 
func (self *WebGLShaderManager) SetTempAttribState(member []interface{}) {
    self.Set("tempAttribState", member)
}

// 
func (self *WebGLShaderManager) GetStack() []interface{}{
	array := self.Get("stack")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// 
func (self *WebGLShaderManager) SetStack(member []interface{}) {
    self.Set("stack", member)
}



// Initialises the context and the properties.
func (self *WebGLShaderManager) SetContextI(args ...interface{}) {
    self.Call("setContext", args)
}

// Takes the attributes given in parameters.
func (self *WebGLShaderManager) SetAttribsI(args ...interface{}) {
    self.Call("setAttribs", args)
}

// Sets the current shader.
func (self *WebGLShaderManager) SetShaderI(args ...interface{}) {
    self.Call("setShader", args)
}

// Destroys this object.
func (self *WebGLShaderManager) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
