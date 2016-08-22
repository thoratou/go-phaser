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
func (self *StripShader) GetGlA() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// 
func (self *StripShader) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// The WebGL program.
func (self *StripShader) GetProgramA() interface{}{
    return self.Object.Get("program")
}

// The WebGL program.
func (self *StripShader) SetProgramA(member interface{}) {
    self.Object.Set("program", member)
}

// The fragment shader.
func (self *StripShader) GetFragmentSrcA() []interface{}{
	array00 := self.Object.Get("fragmentSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// The fragment shader.
func (self *StripShader) SetFragmentSrcA(member []interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// The vertex shader.
func (self *StripShader) GetVertexSrcA() []interface{}{
	array00 := self.Object.Get("vertexSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// The vertex shader.
func (self *StripShader) SetVertexSrcA(member []interface{}) {
    self.Object.Set("vertexSrc", member)
}



// Initialises the shader.
func (self *StripShader) Init() {
    self.Object.Call("init")
}

// Initialises the shader.
func (self *StripShader) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Destroys the shader.
func (self *StripShader) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the shader.
func (self *StripShader) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
