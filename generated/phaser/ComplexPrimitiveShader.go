// Automatic generation for PIXI.ComplexPrimitiveShader
// generated file ComplexPrimitiveShader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type ComplexPrimitiveShader struct {
    *js.Object
}


// 
func (self *ComplexPrimitiveShader) GetGl() WebGLContext{
    return WrapWebGLContext(self.Get("gl"))
}

// 
func (self *ComplexPrimitiveShader) SetGl(member WebGLContext) {
    self.Set("gl", member)
}

// The WebGL program.
func (self *ComplexPrimitiveShader) GetProgram() interface{}{
    return self.Get("program")
}

// The WebGL program.
func (self *ComplexPrimitiveShader) SetProgram(member interface{}) {
    self.Set("program", member)
}

// The fragment shader.
func (self *ComplexPrimitiveShader) GetFragmentSrc() []interface{}{
	array := self.Get("fragmentSrc")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The fragment shader.
func (self *ComplexPrimitiveShader) SetFragmentSrc(member []interface{}) {
    self.Set("fragmentSrc", member)
}

// The vertex shader.
func (self *ComplexPrimitiveShader) GetVertexSrc() []interface{}{
	array := self.Get("vertexSrc")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The vertex shader.
func (self *ComplexPrimitiveShader) SetVertexSrc(member []interface{}) {
    self.Set("vertexSrc", member)
}



// Initialises the shader.
func (self *ComplexPrimitiveShader) InitI(args ...interface{}) {
    self.Call("init", args)
}

// Destroys the shader.
func (self *ComplexPrimitiveShader) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
