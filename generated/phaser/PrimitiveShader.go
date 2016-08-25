// Package phaser Automatic generation for PIXI.PrimitiveShader
// generated file PrimitiveShader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PrimitiveShader empty description
type PrimitiveShader struct {
    *js.Object
}

// NewPrimitiveShader empty description
func NewPrimitiveShader(gl *WebGLContext) *PrimitiveShader {
    return &PrimitiveShader{js.Global.Get("PIXI").Get("PrimitiveShader").New(gl)}
}
// NewPrimitiveShaderI empty description
func NewPrimitiveShaderI(args ...interface{}) *PrimitiveShader {
    return &PrimitiveShader{js.Global.Get("PIXI").Get("PrimitiveShader").New(args)}
}



// PrimitiveShader Binding conversion method to PrimitiveShader point 
func ToPrimitiveShader(jsStruct interface{}) *PrimitiveShader {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PrimitiveShader{Object: object}
	}
	return nil
}



// Gl empty description
func (self *PrimitiveShader) Gl() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// SetGlA empty description
func (self *PrimitiveShader) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// Program The WebGL program.
func (self *PrimitiveShader) Program() interface{}{
    return self.Object.Get("program")
}

// SetProgramA The WebGL program.
func (self *PrimitiveShader) SetProgramA(member interface{}) {
    self.Object.Set("program", member)
}

// FragmentSrc The fragment shader.
func (self *PrimitiveShader) FragmentSrc() []interface{}{
	array00 := self.Object.Get("fragmentSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetFragmentSrcA The fragment shader.
func (self *PrimitiveShader) SetFragmentSrcA(member []interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// VertexSrc The vertex shader.
func (self *PrimitiveShader) VertexSrc() []interface{}{
	array00 := self.Object.Get("vertexSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetVertexSrcA The vertex shader.
func (self *PrimitiveShader) SetVertexSrcA(member []interface{}) {
    self.Object.Set("vertexSrc", member)
}


// Init Initialises the shader.
func (self *PrimitiveShader) Init() {
    self.Object.Call("init")
}

// InitI Initialises the shader.
func (self *PrimitiveShader) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Destroy Destroys the shader.
func (self *PrimitiveShader) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the shader.
func (self *PrimitiveShader) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

