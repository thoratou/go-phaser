// Automatic generation for PIXI.PixiFastShader
// generated file PixiFastShader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type PixiFastShader struct {
    *js.Object
}


// 
func (self *PixiFastShader) GetGl() WebGLContext{
    return WrapWebGLContext(self.Get("gl"))
}

// 
func (self *PixiFastShader) SetGl(member WebGLContext) {
    self.Set("gl", member)
}

// The WebGL program.
func (self *PixiFastShader) GetProgram() interface{}{
    return self.Get("program")
}

// The WebGL program.
func (self *PixiFastShader) SetProgram(member interface{}) {
    self.Set("program", member)
}

// The fragment shader.
func (self *PixiFastShader) GetFragmentSrc() []interface{}{
	array := self.Get("fragmentSrc")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The fragment shader.
func (self *PixiFastShader) SetFragmentSrc(member []interface{}) {
    self.Set("fragmentSrc", member)
}

// The vertex shader.
func (self *PixiFastShader) GetVertexSrc() []interface{}{
	array := self.Get("vertexSrc")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The vertex shader.
func (self *PixiFastShader) SetVertexSrc(member []interface{}) {
    self.Set("vertexSrc", member)
}

// A local texture counter for multi-texture shaders.
func (self *PixiFastShader) GetTextureCount() float64{
    return self.Get("textureCount").Float()
}

// A local texture counter for multi-texture shaders.
func (self *PixiFastShader) SetTextureCount(member float64) {
    self.Set("textureCount", member)
}



// Initialises the shader.
func (self *PixiFastShader) InitI(args ...interface{}) {
    self.Call("init", args)
}

// Destroys the shader.
func (self *PixiFastShader) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
