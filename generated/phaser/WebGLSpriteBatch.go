// Automatic generation for PIXI.WebGLSpriteBatch
// generated file WebGLSpriteBatch.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type WebGLSpriteBatch struct {
    *js.Object
}


// 
func (self *WebGLSpriteBatch) GetVertSize() float64{
    return self.Get("vertSize").Float()
}

// 
func (self *WebGLSpriteBatch) SetVertSize(member float64) {
    self.Set("vertSize", member)
}

// The number of images in the SpriteBatch before it flushes
func (self *WebGLSpriteBatch) GetSize() float64{
    return self.Get("size").Float()
}

// The number of images in the SpriteBatch before it flushes
func (self *WebGLSpriteBatch) SetSize(member float64) {
    self.Set("size", member)
}

// Holds the vertices
func (self *WebGLSpriteBatch) GetVertices() ArrayBuffer{
    return ArrayBuffer{self.Get("vertices")}
}

// Holds the vertices
func (self *WebGLSpriteBatch) SetVertices(member ArrayBuffer) {
    self.Set("vertices", member)
}

// View on the vertices as a Float32Array
func (self *WebGLSpriteBatch) GetPositions() Float32Array{
    return Float32Array{self.Get("positions")}
}

// View on the vertices as a Float32Array
func (self *WebGLSpriteBatch) SetPositions(member Float32Array) {
    self.Set("positions", member)
}

// View on the vertices as a Uint32Array
func (self *WebGLSpriteBatch) GetColors() Uint32Array{
    return Uint32Array{self.Get("colors")}
}

// View on the vertices as a Uint32Array
func (self *WebGLSpriteBatch) SetColors(member Uint32Array) {
    self.Set("colors", member)
}

// Holds the indices
func (self *WebGLSpriteBatch) GetIndices() Uint16Array{
    return Uint16Array{self.Get("indices")}
}

// Holds the indices
func (self *WebGLSpriteBatch) SetIndices(member Uint16Array) {
    self.Set("indices", member)
}

// 
func (self *WebGLSpriteBatch) GetLastIndexCount() float64{
    return self.Get("lastIndexCount").Float()
}

// 
func (self *WebGLSpriteBatch) SetLastIndexCount(member float64) {
    self.Set("lastIndexCount", member)
}

// 
func (self *WebGLSpriteBatch) GetDrawing() bool{
    return self.Get("drawing").Bool()
}

// 
func (self *WebGLSpriteBatch) SetDrawing(member bool) {
    self.Set("drawing", member)
}

// 
func (self *WebGLSpriteBatch) GetCurrentBatchSize() float64{
    return self.Get("currentBatchSize").Float()
}

// 
func (self *WebGLSpriteBatch) SetCurrentBatchSize(member float64) {
    self.Set("currentBatchSize", member)
}

// 
func (self *WebGLSpriteBatch) GetCurrentBaseTexture() BaseTexture{
    return BaseTexture{self.Get("currentBaseTexture")}
}

// 
func (self *WebGLSpriteBatch) SetCurrentBaseTexture(member BaseTexture) {
    self.Set("currentBaseTexture", member)
}

// 
func (self *WebGLSpriteBatch) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// 
func (self *WebGLSpriteBatch) SetDirty(member bool) {
    self.Set("dirty", member)
}

// 
func (self *WebGLSpriteBatch) GetTextures() []interface{}{
	array := self.Get("textures")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// 
func (self *WebGLSpriteBatch) SetTextures(member []interface{}) {
    self.Set("textures", member)
}

// 
func (self *WebGLSpriteBatch) GetBlendModes() []interface{}{
	array := self.Get("blendModes")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// 
func (self *WebGLSpriteBatch) SetBlendModes(member []interface{}) {
    self.Set("blendModes", member)
}

// 
func (self *WebGLSpriteBatch) GetShaders() []interface{}{
	array := self.Get("shaders")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// 
func (self *WebGLSpriteBatch) SetShaders(member []interface{}) {
    self.Set("shaders", member)
}

// 
func (self *WebGLSpriteBatch) GetSprites() []interface{}{
	array := self.Get("sprites")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// 
func (self *WebGLSpriteBatch) SetSprites(member []interface{}) {
    self.Set("sprites", member)
}

// 
func (self *WebGLSpriteBatch) GetDefaultShader() AbstractFilter{
    return AbstractFilter{self.Get("defaultShader")}
}

// 
func (self *WebGLSpriteBatch) SetDefaultShader(member AbstractFilter) {
    self.Set("defaultShader", member)
}



// 
func (self *WebGLSpriteBatch) InitDefaultShadersI(args ...interface{}) {
    self.Call("initDefaultShaders", args)
}

// 
func (self *WebGLSpriteBatch) CompileVertexShaderI(args ...interface{}) {
    self.Call("CompileVertexShader", args)
}

// 
func (self *WebGLSpriteBatch) CompileFragmentShaderI(args ...interface{}) {
    self.Call("CompileFragmentShader", args)
}

// 
func (self *WebGLSpriteBatch) _CompileShaderI(args ...interface{}) {
    self.Call("_CompileShader", args)
}

// 
func (self *WebGLSpriteBatch) CompileProgramI(args ...interface{}) {
    self.Call("compileProgram", args)
}

// 
func (self *WebGLSpriteBatch) SetContextI(args ...interface{}) {
    self.Call("setContext", args)
}

// 
func (self *WebGLSpriteBatch) BeginI(args ...interface{}) {
    self.Call("begin", args)
}

// 
func (self *WebGLSpriteBatch) EndI(args ...interface{}) {
    self.Call("end", args)
}

// 
func (self *WebGLSpriteBatch) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// Renders a TilingSprite using the spriteBatch.
func (self *WebGLSpriteBatch) RenderTilingSpriteI(args ...interface{}) {
    self.Call("renderTilingSprite", args)
}

// Renders the content and empties the current batch.
func (self *WebGLSpriteBatch) FlushI(args ...interface{}) {
    self.Call("flush", args)
}

// 
func (self *WebGLSpriteBatch) RenderBatchI(args ...interface{}) {
    self.Call("renderBatch", args)
}

// 
func (self *WebGLSpriteBatch) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// 
func (self *WebGLSpriteBatch) StartI(args ...interface{}) {
    self.Call("start", args)
}

// Destroys the SpriteBatch.
func (self *WebGLSpriteBatch) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
