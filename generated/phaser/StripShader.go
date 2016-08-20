// Automatic generation for PIXI.StripShader
// generated file StripShader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type StripShader struct {
    *js.Object
}


// 
func (self *StripShader) GetGl() WebGLContext{
    return WrapWebGLContext(self.Get("gl"))
}

// 
func (self *StripShader) SetGl(member WebGLContext) {
    self.Set("gl", member)
}

// The WebGL program.
func (self *StripShader) GetProgram() interface{}{
    return self.Get("program")
}

// The WebGL program.
func (self *StripShader) SetProgram(member interface{}) {
    self.Set("program", member)
}

// The fragment shader.
func (self *StripShader) GetFragmentSrc() []interface{}{
	array := self.Get("fragmentSrc")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The fragment shader.
func (self *StripShader) SetFragmentSrc(member []interface{}) {
    self.Set("fragmentSrc", member)
}

// The vertex shader.
func (self *StripShader) GetVertexSrc() []interface{}{
	array := self.Get("vertexSrc")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The vertex shader.
func (self *StripShader) SetVertexSrc(member []interface{}) {
    self.Set("vertexSrc", member)
}



// Initialises the shader.
func (self *StripShader) InitI(args ...interface{}) {
    self.Call("init", args)
}

// Destroys the shader.
func (self *StripShader) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
