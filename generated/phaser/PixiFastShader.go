// Package phaser Automatic generation for PIXI.PixiFastShader
// generated file PixiFastShader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PixiFastShader empty description
type PixiFastShader struct {
    *js.Object
}

// NewPixiFastShader empty description
func NewPixiFastShader(gl *WebGLContext) *PixiFastShader {
    return &PixiFastShader{js.Global.Get("PIXI").Get("PixiFastShader").New(gl)}
}
// NewPixiFastShaderI empty description
func NewPixiFastShaderI(args ...interface{}) *PixiFastShader {
    return &PixiFastShader{js.Global.Get("PIXI").Get("PixiFastShader").New(args)}
}



// Gl empty description
func (self *PixiFastShader) Gl() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// SetGlA empty description
func (self *PixiFastShader) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// Program The WebGL program.
func (self *PixiFastShader) Program() interface{}{
    return self.Object.Get("program")
}

// SetProgramA The WebGL program.
func (self *PixiFastShader) SetProgramA(member interface{}) {
    self.Object.Set("program", member)
}

// FragmentSrc The fragment shader.
func (self *PixiFastShader) FragmentSrc() []interface{}{
	array00 := self.Object.Get("fragmentSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetFragmentSrcA The fragment shader.
func (self *PixiFastShader) SetFragmentSrcA(member []interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// VertexSrc The vertex shader.
func (self *PixiFastShader) VertexSrc() []interface{}{
	array00 := self.Object.Get("vertexSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetVertexSrcA The vertex shader.
func (self *PixiFastShader) SetVertexSrcA(member []interface{}) {
    self.Object.Set("vertexSrc", member)
}

// TextureCount A local texture counter for multi-texture shaders.
func (self *PixiFastShader) TextureCount() int{
    return self.Object.Get("textureCount").Int()
}

// SetTextureCountA A local texture counter for multi-texture shaders.
func (self *PixiFastShader) SetTextureCountA(member int) {
    self.Object.Set("textureCount", member)
}


// Init Initialises the shader.
func (self *PixiFastShader) Init() {
    self.Object.Call("init")
}

// InitI Initialises the shader.
func (self *PixiFastShader) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Destroy Destroys the shader.
func (self *PixiFastShader) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the shader.
func (self *PixiFastShader) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

