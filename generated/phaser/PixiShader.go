// Automatic generation for PIXI.PixiShader
// generated file PixiShader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type PixiShader struct {
    *js.Object
}


// 
func NewPixiShader(gl *WebGLContext) *PixiShader {
    return &PixiShader{js.Global.Get("PIXI").Get("PixiShader").New(gl)}
}

// 
func NewPixiShaderI(args ...interface{}) *PixiShader {
    return &PixiShader{js.Global.Get("PIXI").Get("PixiShader").New(args)}
}



// 
func (self *PixiShader) Gl() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// 
func (self *PixiShader) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// The WebGL program.
func (self *PixiShader) Program() interface{}{
    return self.Object.Get("program")
}

// The WebGL program.
func (self *PixiShader) SetProgramA(member interface{}) {
    self.Object.Set("program", member)
}

// The fragment shader.
func (self *PixiShader) FragmentSrc() []interface{}{
	array00 := self.Object.Get("fragmentSrc")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// The fragment shader.
func (self *PixiShader) SetFragmentSrcA(member []interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// A local texture counter for multi-texture shaders.
func (self *PixiShader) TextureCount() int{
    return self.Object.Get("textureCount").Int()
}

// A local texture counter for multi-texture shaders.
func (self *PixiShader) SetTextureCountA(member int) {
    self.Object.Set("textureCount", member)
}

// A dirty flag
func (self *PixiShader) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// A dirty flag
func (self *PixiShader) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// The Default Vertex shader source.
func (self *PixiShader) DefaultVertexSrc() string{
    return self.Object.Get("defaultVertexSrc").String()
}

// The Default Vertex shader source.
func (self *PixiShader) SetDefaultVertexSrcA(member string) {
    self.Object.Set("defaultVertexSrc", member)
}



// Initialises the shader.
func (self *PixiShader) Init() {
    self.Object.Call("init")
}

// Initialises the shader.
func (self *PixiShader) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Initialises the shader uniform values.
// 
// Uniforms are specified in the GLSL_ES Specification: http://www.khronos.org/registry/webgl/specs/latest/1.0/
// http://www.khronos.org/registry/gles/specs/2.0/GLSL_ES_Specification_1.0.17.pdf
func (self *PixiShader) InitUniforms() {
    self.Object.Call("initUniforms")
}

// Initialises the shader uniform values.
// 
// Uniforms are specified in the GLSL_ES Specification: http://www.khronos.org/registry/webgl/specs/latest/1.0/
// http://www.khronos.org/registry/gles/specs/2.0/GLSL_ES_Specification_1.0.17.pdf
func (self *PixiShader) InitUniformsI(args ...interface{}) {
    self.Object.Call("initUniforms", args)
}

// Initialises a Sampler2D uniform (which may only be available later on after initUniforms once the texture has loaded)
func (self *PixiShader) InitSampler2D() {
    self.Object.Call("initSampler2D")
}

// Initialises a Sampler2D uniform (which may only be available later on after initUniforms once the texture has loaded)
func (self *PixiShader) InitSampler2DI(args ...interface{}) {
    self.Object.Call("initSampler2D", args)
}

// Updates the shader uniform values.
func (self *PixiShader) SyncUniforms() {
    self.Object.Call("syncUniforms")
}

// Updates the shader uniform values.
func (self *PixiShader) SyncUniformsI(args ...interface{}) {
    self.Object.Call("syncUniforms", args)
}

// Destroys the shader.
func (self *PixiShader) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the shader.
func (self *PixiShader) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
