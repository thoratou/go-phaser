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
func (self *PixiShader) GetGl() WebGLContext{
    return WrapWebGLContext(self.Get("gl"))
}

// 
func (self *PixiShader) SetGl(member WebGLContext) {
    self.Set("gl", member)
}

// The WebGL program.
func (self *PixiShader) GetProgram() interface{}{
    return self.Get("program")
}

// The WebGL program.
func (self *PixiShader) SetProgram(member interface{}) {
    self.Set("program", member)
}

// The fragment shader.
func (self *PixiShader) GetFragmentSrc() []interface{}{
	array := self.Get("fragmentSrc")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The fragment shader.
func (self *PixiShader) SetFragmentSrc(member []interface{}) {
    self.Set("fragmentSrc", member)
}

// A local texture counter for multi-texture shaders.
func (self *PixiShader) GetTextureCount() float64{
    return self.Get("textureCount").Float()
}

// A local texture counter for multi-texture shaders.
func (self *PixiShader) SetTextureCount(member float64) {
    self.Set("textureCount", member)
}

// A dirty flag
func (self *PixiShader) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// A dirty flag
func (self *PixiShader) SetDirty(member bool) {
    self.Set("dirty", member)
}

// The Default Vertex shader source.
func (self *PixiShader) GetDefaultVertexSrc() string{
    return self.Get("defaultVertexSrc").String()
}

// The Default Vertex shader source.
func (self *PixiShader) SetDefaultVertexSrc(member string) {
    self.Set("defaultVertexSrc", member)
}



// Initialises the shader.
func (self *PixiShader) InitI(args ...interface{}) {
    self.Call("init", args)
}

// Initialises the shader uniform values.
// 
// Uniforms are specified in the GLSL_ES Specification: http://www.khronos.org/registry/webgl/specs/latest/1.0/
// http://www.khronos.org/registry/gles/specs/2.0/GLSL_ES_Specification_1.0.17.pdf
func (self *PixiShader) InitUniformsI(args ...interface{}) {
    self.Call("initUniforms", args)
}

// Initialises a Sampler2D uniform (which may only be available later on after initUniforms once the texture has loaded)
func (self *PixiShader) InitSampler2DI(args ...interface{}) {
    self.Call("initSampler2D", args)
}

// Updates the shader uniform values.
func (self *PixiShader) SyncUniformsI(args ...interface{}) {
    self.Call("syncUniforms", args)
}

// Destroys the shader.
func (self *PixiShader) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
