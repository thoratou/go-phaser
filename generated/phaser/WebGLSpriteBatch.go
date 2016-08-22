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
func (self *WebGLSpriteBatch) GetVertSizeA() int{
    return self.Object.Get("vertSize").Int()
}

// 
func (self *WebGLSpriteBatch) SetVertSizeA(member int) {
    self.Object.Set("vertSize", member)
}

// The number of images in the SpriteBatch before it flushes
func (self *WebGLSpriteBatch) GetSizeA() int{
    return self.Object.Get("size").Int()
}

// The number of images in the SpriteBatch before it flushes
func (self *WebGLSpriteBatch) SetSizeA(member int) {
    self.Object.Set("size", member)
}

// Holds the vertices
func (self *WebGLSpriteBatch) GetVerticesA() *ArrayBuffer{
    return &ArrayBuffer{self.Object.Get("vertices")}
}

// Holds the vertices
func (self *WebGLSpriteBatch) SetVerticesA(member *ArrayBuffer) {
    self.Object.Set("vertices", member)
}

// View on the vertices as a Float32Array
func (self *WebGLSpriteBatch) GetPositionsA() *Float32Array{
    return &Float32Array{self.Object.Get("positions")}
}

// View on the vertices as a Float32Array
func (self *WebGLSpriteBatch) SetPositionsA(member *Float32Array) {
    self.Object.Set("positions", member)
}

// View on the vertices as a Uint32Array
func (self *WebGLSpriteBatch) GetColorsA() *Uint32Array{
    return &Uint32Array{self.Object.Get("colors")}
}

// View on the vertices as a Uint32Array
func (self *WebGLSpriteBatch) SetColorsA(member *Uint32Array) {
    self.Object.Set("colors", member)
}

// Holds the indices
func (self *WebGLSpriteBatch) GetIndicesA() *Uint16Array{
    return &Uint16Array{self.Object.Get("indices")}
}

// Holds the indices
func (self *WebGLSpriteBatch) SetIndicesA(member *Uint16Array) {
    self.Object.Set("indices", member)
}

// 
func (self *WebGLSpriteBatch) GetLastIndexCountA() int{
    return self.Object.Get("lastIndexCount").Int()
}

// 
func (self *WebGLSpriteBatch) SetLastIndexCountA(member int) {
    self.Object.Set("lastIndexCount", member)
}

// 
func (self *WebGLSpriteBatch) GetDrawingA() bool{
    return self.Object.Get("drawing").Bool()
}

// 
func (self *WebGLSpriteBatch) SetDrawingA(member bool) {
    self.Object.Set("drawing", member)
}

// 
func (self *WebGLSpriteBatch) GetCurrentBatchSizeA() int{
    return self.Object.Get("currentBatchSize").Int()
}

// 
func (self *WebGLSpriteBatch) SetCurrentBatchSizeA(member int) {
    self.Object.Set("currentBatchSize", member)
}

// 
func (self *WebGLSpriteBatch) GetCurrentBaseTextureA() *BaseTexture{
    return &BaseTexture{self.Object.Get("currentBaseTexture")}
}

// 
func (self *WebGLSpriteBatch) SetCurrentBaseTextureA(member *BaseTexture) {
    self.Object.Set("currentBaseTexture", member)
}

// 
func (self *WebGLSpriteBatch) GetDirtyA() bool{
    return self.Object.Get("dirty").Bool()
}

// 
func (self *WebGLSpriteBatch) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// 
func (self *WebGLSpriteBatch) GetTexturesA() []interface{}{
	array00 := self.Object.Get("textures")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// 
func (self *WebGLSpriteBatch) SetTexturesA(member []interface{}) {
    self.Object.Set("textures", member)
}

// 
func (self *WebGLSpriteBatch) GetBlendModesA() []interface{}{
	array00 := self.Object.Get("blendModes")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// 
func (self *WebGLSpriteBatch) SetBlendModesA(member []interface{}) {
    self.Object.Set("blendModes", member)
}

// 
func (self *WebGLSpriteBatch) GetShadersA() []interface{}{
	array00 := self.Object.Get("shaders")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// 
func (self *WebGLSpriteBatch) SetShadersA(member []interface{}) {
    self.Object.Set("shaders", member)
}

// 
func (self *WebGLSpriteBatch) GetSpritesA() []interface{}{
	array00 := self.Object.Get("sprites")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// 
func (self *WebGLSpriteBatch) SetSpritesA(member []interface{}) {
    self.Object.Set("sprites", member)
}

// 
func (self *WebGLSpriteBatch) GetDefaultShaderA() *AbstractFilter{
    return &AbstractFilter{self.Object.Get("defaultShader")}
}

// 
func (self *WebGLSpriteBatch) SetDefaultShaderA(member *AbstractFilter) {
    self.Object.Set("defaultShader", member)
}



// 
func (self *WebGLSpriteBatch) InitDefaultShaders() {
    self.Object.Call("initDefaultShaders")
}

// 
func (self *WebGLSpriteBatch) InitDefaultShadersI(args ...interface{}) {
    self.Object.Call("initDefaultShaders", args)
}

// 
func (self *WebGLSpriteBatch) CompileVertexShader(gl *WebGLContext, shaderSrc []interface{}) {
    self.Object.Call("CompileVertexShader", gl, shaderSrc)
}

// 
func (self *WebGLSpriteBatch) CompileVertexShaderI(args ...interface{}) {
    self.Object.Call("CompileVertexShader", args)
}

// 
func (self *WebGLSpriteBatch) CompileFragmentShader(gl *WebGLContext, shaderSrc []interface{}) {
    self.Object.Call("CompileFragmentShader", gl, shaderSrc)
}

// 
func (self *WebGLSpriteBatch) CompileFragmentShaderI(args ...interface{}) {
    self.Object.Call("CompileFragmentShader", args)
}

// 
func (self *WebGLSpriteBatch) _CompileShader(gl *WebGLContext, shaderSrc []interface{}, shaderType int) {
    self.Object.Call("_CompileShader", gl, shaderSrc, shaderType)
}

// 
func (self *WebGLSpriteBatch) _CompileShaderI(args ...interface{}) {
    self.Object.Call("_CompileShader", args)
}

// 
func (self *WebGLSpriteBatch) CompileProgram(gl *WebGLContext, vertexSrc []interface{}, fragmentSrc []interface{}) {
    self.Object.Call("compileProgram", gl, vertexSrc, fragmentSrc)
}

// 
func (self *WebGLSpriteBatch) CompileProgramI(args ...interface{}) {
    self.Object.Call("compileProgram", args)
}

// 
func (self *WebGLSpriteBatch) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// 
func (self *WebGLSpriteBatch) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// 
func (self *WebGLSpriteBatch) Begin(renderSession interface{}) {
    self.Object.Call("begin", renderSession)
}

// 
func (self *WebGLSpriteBatch) BeginI(args ...interface{}) {
    self.Object.Call("begin", args)
}

// 
func (self *WebGLSpriteBatch) End() {
    self.Object.Call("end")
}

// 
func (self *WebGLSpriteBatch) EndI(args ...interface{}) {
    self.Object.Call("end", args)
}

// 
func (self *WebGLSpriteBatch) Render(sprite *Sprite) {
    self.Object.Call("render", sprite)
}

// 
func (self *WebGLSpriteBatch) Render1O(sprite *Sprite, matrix *Matrix) {
    self.Object.Call("render", sprite, matrix)
}

// 
func (self *WebGLSpriteBatch) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// Renders a TilingSprite using the spriteBatch.
func (self *WebGLSpriteBatch) RenderTilingSprite(sprite *TilingSprite) {
    self.Object.Call("renderTilingSprite", sprite)
}

// Renders a TilingSprite using the spriteBatch.
func (self *WebGLSpriteBatch) RenderTilingSpriteI(args ...interface{}) {
    self.Object.Call("renderTilingSprite", args)
}

// Renders the content and empties the current batch.
func (self *WebGLSpriteBatch) Flush() {
    self.Object.Call("flush")
}

// Renders the content and empties the current batch.
func (self *WebGLSpriteBatch) FlushI(args ...interface{}) {
    self.Object.Call("flush", args)
}

// 
func (self *WebGLSpriteBatch) RenderBatch(texture *Texture, size int, startIndex int) {
    self.Object.Call("renderBatch", texture, size, startIndex)
}

// 
func (self *WebGLSpriteBatch) RenderBatchI(args ...interface{}) {
    self.Object.Call("renderBatch", args)
}

// 
func (self *WebGLSpriteBatch) Stop() {
    self.Object.Call("stop")
}

// 
func (self *WebGLSpriteBatch) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// 
func (self *WebGLSpriteBatch) Start() {
    self.Object.Call("start")
}

// 
func (self *WebGLSpriteBatch) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Destroys the SpriteBatch.
func (self *WebGLSpriteBatch) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the SpriteBatch.
func (self *WebGLSpriteBatch) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
