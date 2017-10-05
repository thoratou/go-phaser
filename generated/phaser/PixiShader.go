// Package phaser Automatic generation for PIXI.PixiShader
// generated file PixiShader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PixiShader empty description
type PixiShader struct {
    *js.Object
}

// NewPixiShader empty description
func NewPixiShader(gl *WebGLContext) *PixiShader {
    return &PixiShader{js.Global.Get("PIXI").Get("PixiShader").New(gl)}
}
// NewPixiShaderI empty description
func NewPixiShaderI(args ...interface{}) *PixiShader {
    return &PixiShader{js.Global.Get("PIXI").Get("PixiShader").New(args)}
}



// PixiShader Binding conversion method to PixiShader point 
func ToPixiShader(jsStruct interface{}) *PixiShader {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PixiShader{Object: object}
	}
	return nil
}



// Gl empty description
func (self *PixiShader) Gl() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// SetGlA empty description
func (self *PixiShader) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// Program The WebGL program.
func (self *PixiShader) Program() interface{}{
    return self.Object.Get("program")
}

// SetProgramA The WebGL program.
func (self *PixiShader) SetProgramA(member interface{}) {
    self.Object.Set("program", member)
}

// FragmentSrc The fragment shader.
func (self *PixiShader) FragmentSrc() []interface{}{
	array00 := self.Object.Get("fragmentSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetFragmentSrcA The fragment shader.
func (self *PixiShader) SetFragmentSrcA(member []interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// TextureCount A local texture counter for multi-texture shaders.
func (self *PixiShader) TextureCount() int{
    return self.Object.Get("textureCount").Int()
}

// SetTextureCountA A local texture counter for multi-texture shaders.
func (self *PixiShader) SetTextureCountA(member int) {
    self.Object.Set("textureCount", member)
}

// Dirty A dirty flag
func (self *PixiShader) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// SetDirtyA A dirty flag
func (self *PixiShader) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// DefaultVertexSrc The Default Vertex shader source.
func (self *PixiShader) DefaultVertexSrc() string{
    return self.Object.Get("defaultVertexSrc").String()
}

// SetDefaultVertexSrcA The Default Vertex shader source.
func (self *PixiShader) SetDefaultVertexSrcA(member string) {
    self.Object.Set("defaultVertexSrc", member)
}


// Init Initialises the shader.
func (self *PixiShader) Init() {
    self.Object.Call("init")
}

// InitI Initialises the shader.
func (self *PixiShader) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// InitUniforms Initialises the shader uniform values.
// 
// 
// 
// Uniforms are specified in the GLSL_ES Specification: http://www.khronos.org/registry/webgl/specs/latest/1.0/
// 
// http://www.khronos.org/registry/gles/specs/2.0/GLSL_ES_Specification_1.0.17.pdf
func (self *PixiShader) InitUniforms() {
    self.Object.Call("initUniforms")
}

// InitUniformsI Initialises the shader uniform values.
// 
// 
// 
// Uniforms are specified in the GLSL_ES Specification: http://www.khronos.org/registry/webgl/specs/latest/1.0/
// 
// http://www.khronos.org/registry/gles/specs/2.0/GLSL_ES_Specification_1.0.17.pdf
func (self *PixiShader) InitUniformsI(args ...interface{}) {
    self.Object.Call("initUniforms", args)
}

// InitSampler2D Initialises a Sampler2D uniform (which may only be available later on after initUniforms once the texture has loaded)
func (self *PixiShader) InitSampler2D() {
    self.Object.Call("initSampler2D")
}

// InitSampler2DI Initialises a Sampler2D uniform (which may only be available later on after initUniforms once the texture has loaded)
func (self *PixiShader) InitSampler2DI(args ...interface{}) {
    self.Object.Call("initSampler2D", args)
}

// SyncUniforms Updates the shader uniform values.
func (self *PixiShader) SyncUniforms() {
    self.Object.Call("syncUniforms")
}

// SyncUniformsI Updates the shader uniform values.
func (self *PixiShader) SyncUniformsI(args ...interface{}) {
    self.Object.Call("syncUniforms", args)
}

// Destroy Destroys the shader.
func (self *PixiShader) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the shader.
func (self *PixiShader) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

