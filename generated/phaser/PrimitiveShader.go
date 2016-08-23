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
func NewPrimitiveShader(gl *WebGLContext) *PrimitiveShader {
    return &PrimitiveShader{js.Global.Get("PIXI").Get("PrimitiveShader").New(gl)}
}

// 
func NewPrimitiveShaderI(args ...interface{}) *PrimitiveShader {
    return &PrimitiveShader{js.Global.Get("PIXI").Get("PrimitiveShader").New(args)}
}



// 
func (self *PrimitiveShader) Gl() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// 
func (self *PrimitiveShader) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// The WebGL program.
func (self *PrimitiveShader) Program() interface{}{
    return self.Object.Get("program")
}

// The WebGL program.
func (self *PrimitiveShader) SetProgramA(member interface{}) {
    self.Object.Set("program", member)
}

// The fragment shader.
func (self *PrimitiveShader) FragmentSrc() []interface{}{
	array00 := self.Object.Get("fragmentSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// The fragment shader.
func (self *PrimitiveShader) SetFragmentSrcA(member []interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// The vertex shader.
func (self *PrimitiveShader) VertexSrc() []interface{}{
	array00 := self.Object.Get("vertexSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// The vertex shader.
func (self *PrimitiveShader) SetVertexSrcA(member []interface{}) {
    self.Object.Set("vertexSrc", member)
}



// Initialises the shader.
func (self *PrimitiveShader) Init() {
    self.Object.Call("init")
}

// Initialises the shader.
func (self *PrimitiveShader) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Destroys the shader.
func (self *PrimitiveShader) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the shader.
func (self *PrimitiveShader) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
