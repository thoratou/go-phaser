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
func NewComplexPrimitiveShader(gl *WebGLContext) *ComplexPrimitiveShader {
    return &ComplexPrimitiveShader{js.Global.Call("PIXI.ComplexPrimitiveShader", gl)}
}

// 
func NewComplexPrimitiveShaderI(args ...interface{}) *ComplexPrimitiveShader {
    return &ComplexPrimitiveShader{js.Global.Call("PIXI.ComplexPrimitiveShader", args)}
}



// 
func (self *ComplexPrimitiveShader) GetGlA() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// 
func (self *ComplexPrimitiveShader) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// The WebGL program.
func (self *ComplexPrimitiveShader) GetProgramA() interface{}{
    return self.Object.Get("program")
}

// The WebGL program.
func (self *ComplexPrimitiveShader) SetProgramA(member interface{}) {
    self.Object.Set("program", member)
}

// The fragment shader.
func (self *ComplexPrimitiveShader) GetFragmentSrcA() []interface{}{
	array00 := self.Object.Get("fragmentSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// The fragment shader.
func (self *ComplexPrimitiveShader) SetFragmentSrcA(member []interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// The vertex shader.
func (self *ComplexPrimitiveShader) GetVertexSrcA() []interface{}{
	array00 := self.Object.Get("vertexSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// The vertex shader.
func (self *ComplexPrimitiveShader) SetVertexSrcA(member []interface{}) {
    self.Object.Set("vertexSrc", member)
}



// Initialises the shader.
func (self *ComplexPrimitiveShader) Init() {
    self.Object.Call("init")
}

// Initialises the shader.
func (self *ComplexPrimitiveShader) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Destroys the shader.
func (self *ComplexPrimitiveShader) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the shader.
func (self *ComplexPrimitiveShader) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
