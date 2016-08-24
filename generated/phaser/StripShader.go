// Package phaser Automatic generation for PIXI.StripShader
// generated file StripShader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// StripShader empty description
type StripShader struct {
    *js.Object
}

// NewStripShader empty description
func NewStripShader(gl *WebGLContext) *StripShader {
    return &StripShader{js.Global.Get("PIXI").Get("StripShader").New(gl)}
}
// NewStripShaderI empty description
func NewStripShaderI(args ...interface{}) *StripShader {
    return &StripShader{js.Global.Get("PIXI").Get("StripShader").New(args)}
}



// Gl empty description
func (self *StripShader) Gl() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// SetGlA empty description
func (self *StripShader) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// Program The WebGL program.
func (self *StripShader) Program() interface{}{
    return self.Object.Get("program")
}

// SetProgramA The WebGL program.
func (self *StripShader) SetProgramA(member interface{}) {
    self.Object.Set("program", member)
}

// FragmentSrc The fragment shader.
func (self *StripShader) FragmentSrc() []interface{}{
	array00 := self.Object.Get("fragmentSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetFragmentSrcA The fragment shader.
func (self *StripShader) SetFragmentSrcA(member []interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// VertexSrc The vertex shader.
func (self *StripShader) VertexSrc() []interface{}{
	array00 := self.Object.Get("vertexSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetVertexSrcA The vertex shader.
func (self *StripShader) SetVertexSrcA(member []interface{}) {
    self.Object.Set("vertexSrc", member)
}


// Init Initialises the shader.
func (self *StripShader) Init() {
    self.Object.Call("init")
}

// InitI Initialises the shader.
func (self *StripShader) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Destroy Destroys the shader.
func (self *StripShader) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the shader.
func (self *StripShader) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

