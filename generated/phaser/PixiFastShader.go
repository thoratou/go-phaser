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
func NewPixiFastShader(gl *WebGLContext) *PixiFastShader {
    return &PixiFastShader{js.Global.Get("PIXI").Get("PixiFastShader").New(gl)}
}

// 
func NewPixiFastShaderI(args ...interface{}) *PixiFastShader {
    return &PixiFastShader{js.Global.Get("PIXI").Get("PixiFastShader").New(args)}
}



// 
func (self *PixiFastShader) Gl() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// 
func (self *PixiFastShader) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// The WebGL program.
func (self *PixiFastShader) Program() interface{}{
    return self.Object.Get("program")
}

// The WebGL program.
func (self *PixiFastShader) SetProgramA(member interface{}) {
    self.Object.Set("program", member)
}

// The fragment shader.
func (self *PixiFastShader) FragmentSrc() []interface{}{
	array00 := self.Object.Get("fragmentSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// The fragment shader.
func (self *PixiFastShader) SetFragmentSrcA(member []interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// The vertex shader.
func (self *PixiFastShader) VertexSrc() []interface{}{
	array00 := self.Object.Get("vertexSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// The vertex shader.
func (self *PixiFastShader) SetVertexSrcA(member []interface{}) {
    self.Object.Set("vertexSrc", member)
}

// A local texture counter for multi-texture shaders.
func (self *PixiFastShader) TextureCount() int{
    return self.Object.Get("textureCount").Int()
}

// A local texture counter for multi-texture shaders.
func (self *PixiFastShader) SetTextureCountA(member int) {
    self.Object.Set("textureCount", member)
}



// Initialises the shader.
func (self *PixiFastShader) Init() {
    self.Object.Call("init")
}

// Initialises the shader.
func (self *PixiFastShader) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Destroys the shader.
func (self *PixiFastShader) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the shader.
func (self *PixiFastShader) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
