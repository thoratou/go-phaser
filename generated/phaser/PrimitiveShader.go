// Automatic generation for PIXI.PrimitiveShader
// generated file PrimitiveShader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type PrimitiveShader struct {
    *js.Object
}


// 
func (self *PrimitiveShader) GetGl() WebGLContext{
    return WrapWebGLContext(self.Get("gl"))
}

// 
func (self *PrimitiveShader) SetGl(member WebGLContext) {
    self.Set("gl", member)
}

// The WebGL program.
func (self *PrimitiveShader) GetProgram() interface{}{
    return self.Get("program")
}

// The WebGL program.
func (self *PrimitiveShader) SetProgram(member interface{}) {
    self.Set("program", member)
}

// The fragment shader.
func (self *PrimitiveShader) GetFragmentSrc() []interface{}{
	array := self.Get("fragmentSrc")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The fragment shader.
func (self *PrimitiveShader) SetFragmentSrc(member []interface{}) {
    self.Set("fragmentSrc", member)
}

// The vertex shader.
func (self *PrimitiveShader) GetVertexSrc() []interface{}{
	array := self.Get("vertexSrc")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The vertex shader.
func (self *PrimitiveShader) SetVertexSrc(member []interface{}) {
    self.Set("vertexSrc", member)
}



// Initialises the shader.
func (self *PrimitiveShader) InitI(args ...interface{}) {
    self.Call("init", args)
}

// Destroys the shader.
func (self *PrimitiveShader) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
