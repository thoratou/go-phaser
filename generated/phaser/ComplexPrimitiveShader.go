// Package phaser Automatic generation for PIXI.ComplexPrimitiveShader
// generated file ComplexPrimitiveShader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ComplexPrimitiveShader empty description
type ComplexPrimitiveShader struct {
    *js.Object
}

// NewComplexPrimitiveShader empty description
func NewComplexPrimitiveShader(gl *WebGLContext) *ComplexPrimitiveShader {
    return &ComplexPrimitiveShader{js.Global.Get("PIXI").Get("ComplexPrimitiveShader").New(gl)}
}
// NewComplexPrimitiveShaderI empty description
func NewComplexPrimitiveShaderI(args ...interface{}) *ComplexPrimitiveShader {
    return &ComplexPrimitiveShader{js.Global.Get("PIXI").Get("ComplexPrimitiveShader").New(args)}
}



// ComplexPrimitiveShader Binding conversion method to ComplexPrimitiveShader point 
func ToComplexPrimitiveShader(jsStruct interface{}) *ComplexPrimitiveShader {
    if object, ok := jsStruct.(*js.Object); ok {
		return &ComplexPrimitiveShader{Object: object}
	}
	return nil
}



// Gl empty description
func (self *ComplexPrimitiveShader) Gl() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// SetGlA empty description
func (self *ComplexPrimitiveShader) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// Program The WebGL program.
func (self *ComplexPrimitiveShader) Program() interface{}{
    return self.Object.Get("program")
}

// SetProgramA The WebGL program.
func (self *ComplexPrimitiveShader) SetProgramA(member interface{}) {
    self.Object.Set("program", member)
}

// FragmentSrc The fragment shader.
func (self *ComplexPrimitiveShader) FragmentSrc() []interface{}{
	array00 := self.Object.Get("fragmentSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetFragmentSrcA The fragment shader.
func (self *ComplexPrimitiveShader) SetFragmentSrcA(member []interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// VertexSrc The vertex shader.
func (self *ComplexPrimitiveShader) VertexSrc() []interface{}{
	array00 := self.Object.Get("vertexSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetVertexSrcA The vertex shader.
func (self *ComplexPrimitiveShader) SetVertexSrcA(member []interface{}) {
    self.Object.Set("vertexSrc", member)
}


// Init Initialises the shader.
func (self *ComplexPrimitiveShader) Init() {
    self.Object.Call("init")
}

// InitI Initialises the shader.
func (self *ComplexPrimitiveShader) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Destroy Destroys the shader.
func (self *ComplexPrimitiveShader) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the shader.
func (self *ComplexPrimitiveShader) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

