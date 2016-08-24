// Package phaser Automatic generation for PIXI.WebGLSpriteBatch
// generated file WebGLSpriteBatch.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// WebGLSpriteBatch empty description
type WebGLSpriteBatch struct {
    *js.Object
}

// NewWebGLSpriteBatch empty description
func NewWebGLSpriteBatch() *WebGLSpriteBatch {
    return &WebGLSpriteBatch{js.Global.Get("PIXI").Get("WebGLSpriteBatch").New()}
}
// NewWebGLSpriteBatchI empty description
func NewWebGLSpriteBatchI(args ...interface{}) *WebGLSpriteBatch {
    return &WebGLSpriteBatch{js.Global.Get("PIXI").Get("WebGLSpriteBatch").New(args)}
}



// VertSize empty description
func (self *WebGLSpriteBatch) VertSize() int{
    return self.Object.Get("vertSize").Int()
}

// SetVertSizeA empty description
func (self *WebGLSpriteBatch) SetVertSizeA(member int) {
    self.Object.Set("vertSize", member)
}

// Size The number of images in the SpriteBatch before it flushes
func (self *WebGLSpriteBatch) Size() int{
    return self.Object.Get("size").Int()
}

// SetSizeA The number of images in the SpriteBatch before it flushes
func (self *WebGLSpriteBatch) SetSizeA(member int) {
    self.Object.Set("size", member)
}

// Vertices Holds the vertices
func (self *WebGLSpriteBatch) Vertices() *ArrayBuffer{
    return &ArrayBuffer{self.Object.Get("vertices")}
}

// SetVerticesA Holds the vertices
func (self *WebGLSpriteBatch) SetVerticesA(member *ArrayBuffer) {
    self.Object.Set("vertices", member)
}

// Positions View on the vertices as a Float32Array
func (self *WebGLSpriteBatch) Positions() *Float32Array{
    return &Float32Array{self.Object.Get("positions")}
}

// SetPositionsA View on the vertices as a Float32Array
func (self *WebGLSpriteBatch) SetPositionsA(member *Float32Array) {
    self.Object.Set("positions", member)
}

// Colors View on the vertices as a Uint32Array
func (self *WebGLSpriteBatch) Colors() *Uint32Array{
    return &Uint32Array{self.Object.Get("colors")}
}

// SetColorsA View on the vertices as a Uint32Array
func (self *WebGLSpriteBatch) SetColorsA(member *Uint32Array) {
    self.Object.Set("colors", member)
}

// Indices Holds the indices
func (self *WebGLSpriteBatch) Indices() *Uint16Array{
    return &Uint16Array{self.Object.Get("indices")}
}

// SetIndicesA Holds the indices
func (self *WebGLSpriteBatch) SetIndicesA(member *Uint16Array) {
    self.Object.Set("indices", member)
}

// LastIndexCount empty description
func (self *WebGLSpriteBatch) LastIndexCount() int{
    return self.Object.Get("lastIndexCount").Int()
}

// SetLastIndexCountA empty description
func (self *WebGLSpriteBatch) SetLastIndexCountA(member int) {
    self.Object.Set("lastIndexCount", member)
}

// Drawing empty description
func (self *WebGLSpriteBatch) Drawing() bool{
    return self.Object.Get("drawing").Bool()
}

// SetDrawingA empty description
func (self *WebGLSpriteBatch) SetDrawingA(member bool) {
    self.Object.Set("drawing", member)
}

// CurrentBatchSize empty description
func (self *WebGLSpriteBatch) CurrentBatchSize() int{
    return self.Object.Get("currentBatchSize").Int()
}

// SetCurrentBatchSizeA empty description
func (self *WebGLSpriteBatch) SetCurrentBatchSizeA(member int) {
    self.Object.Set("currentBatchSize", member)
}

// CurrentBaseTexture empty description
func (self *WebGLSpriteBatch) CurrentBaseTexture() *BaseTexture{
    return &BaseTexture{self.Object.Get("currentBaseTexture")}
}

// SetCurrentBaseTextureA empty description
func (self *WebGLSpriteBatch) SetCurrentBaseTextureA(member *BaseTexture) {
    self.Object.Set("currentBaseTexture", member)
}

// Dirty empty description
func (self *WebGLSpriteBatch) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// SetDirtyA empty description
func (self *WebGLSpriteBatch) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// Textures empty description
func (self *WebGLSpriteBatch) Textures() []interface{}{
	array00 := self.Object.Get("textures")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetTexturesA empty description
func (self *WebGLSpriteBatch) SetTexturesA(member []interface{}) {
    self.Object.Set("textures", member)
}

// BlendModes empty description
func (self *WebGLSpriteBatch) BlendModes() []interface{}{
	array00 := self.Object.Get("blendModes")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetBlendModesA empty description
func (self *WebGLSpriteBatch) SetBlendModesA(member []interface{}) {
    self.Object.Set("blendModes", member)
}

// Shaders empty description
func (self *WebGLSpriteBatch) Shaders() []interface{}{
	array00 := self.Object.Get("shaders")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetShadersA empty description
func (self *WebGLSpriteBatch) SetShadersA(member []interface{}) {
    self.Object.Set("shaders", member)
}

// Sprites empty description
func (self *WebGLSpriteBatch) Sprites() []interface{}{
	array00 := self.Object.Get("sprites")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetSpritesA empty description
func (self *WebGLSpriteBatch) SetSpritesA(member []interface{}) {
    self.Object.Set("sprites", member)
}

// DefaultShader empty description
func (self *WebGLSpriteBatch) DefaultShader() *AbstractFilter{
    return &AbstractFilter{self.Object.Get("defaultShader")}
}

// SetDefaultShaderA empty description
func (self *WebGLSpriteBatch) SetDefaultShaderA(member *AbstractFilter) {
    self.Object.Set("defaultShader", member)
}


// InitDefaultShaders empty description
func (self *WebGLSpriteBatch) InitDefaultShaders() {
    self.Object.Call("initDefaultShaders")
}

// InitDefaultShadersI empty description
func (self *WebGLSpriteBatch) InitDefaultShadersI(args ...interface{}) {
    self.Object.Call("initDefaultShaders", args)
}

// CompileVertexShader empty description
func (self *WebGLSpriteBatch) CompileVertexShader(gl *WebGLContext, shaderSrc []interface{}) {
    self.Object.Call("CompileVertexShader", gl, shaderSrc)
}

// CompileVertexShaderI empty description
func (self *WebGLSpriteBatch) CompileVertexShaderI(args ...interface{}) {
    self.Object.Call("CompileVertexShader", args)
}

// CompileFragmentShader empty description
func (self *WebGLSpriteBatch) CompileFragmentShader(gl *WebGLContext, shaderSrc []interface{}) {
    self.Object.Call("CompileFragmentShader", gl, shaderSrc)
}

// CompileFragmentShaderI empty description
func (self *WebGLSpriteBatch) CompileFragmentShaderI(args ...interface{}) {
    self.Object.Call("CompileFragmentShader", args)
}

// _CompileShader empty description
func (self *WebGLSpriteBatch) _CompileShader(gl *WebGLContext, shaderSrc []interface{}, shaderType int) {
    self.Object.Call("_CompileShader", gl, shaderSrc, shaderType)
}

// _CompileShaderI empty description
func (self *WebGLSpriteBatch) _CompileShaderI(args ...interface{}) {
    self.Object.Call("_CompileShader", args)
}

// CompileProgram empty description
func (self *WebGLSpriteBatch) CompileProgram(gl *WebGLContext, vertexSrc []interface{}, fragmentSrc []interface{}) {
    self.Object.Call("compileProgram", gl, vertexSrc, fragmentSrc)
}

// CompileProgramI empty description
func (self *WebGLSpriteBatch) CompileProgramI(args ...interface{}) {
    self.Object.Call("compileProgram", args)
}

// SetContext empty description
func (self *WebGLSpriteBatch) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// SetContextI empty description
func (self *WebGLSpriteBatch) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// Begin empty description
func (self *WebGLSpriteBatch) Begin(renderSession interface{}) {
    self.Object.Call("begin", renderSession)
}

// BeginI empty description
func (self *WebGLSpriteBatch) BeginI(args ...interface{}) {
    self.Object.Call("begin", args)
}

// End empty description
func (self *WebGLSpriteBatch) End() {
    self.Object.Call("end")
}

// EndI empty description
func (self *WebGLSpriteBatch) EndI(args ...interface{}) {
    self.Object.Call("end", args)
}

// Render empty description
func (self *WebGLSpriteBatch) Render(sprite *Sprite) {
    self.Object.Call("render", sprite)
}

// Render1O empty description
func (self *WebGLSpriteBatch) Render1O(sprite *Sprite, matrix *Matrix) {
    self.Object.Call("render", sprite, matrix)
}

// RenderI empty description
func (self *WebGLSpriteBatch) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// RenderTilingSprite Renders a TilingSprite using the spriteBatch.
func (self *WebGLSpriteBatch) RenderTilingSprite(sprite *TilingSprite) {
    self.Object.Call("renderTilingSprite", sprite)
}

// RenderTilingSpriteI Renders a TilingSprite using the spriteBatch.
func (self *WebGLSpriteBatch) RenderTilingSpriteI(args ...interface{}) {
    self.Object.Call("renderTilingSprite", args)
}

// Flush Renders the content and empties the current batch.
func (self *WebGLSpriteBatch) Flush() {
    self.Object.Call("flush")
}

// FlushI Renders the content and empties the current batch.
func (self *WebGLSpriteBatch) FlushI(args ...interface{}) {
    self.Object.Call("flush", args)
}

// RenderBatch empty description
func (self *WebGLSpriteBatch) RenderBatch(texture *Texture, size int, startIndex int) {
    self.Object.Call("renderBatch", texture, size, startIndex)
}

// RenderBatchI empty description
func (self *WebGLSpriteBatch) RenderBatchI(args ...interface{}) {
    self.Object.Call("renderBatch", args)
}

// Stop empty description
func (self *WebGLSpriteBatch) Stop() {
    self.Object.Call("stop")
}

// StopI empty description
func (self *WebGLSpriteBatch) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Start empty description
func (self *WebGLSpriteBatch) Start() {
    self.Object.Call("start")
}

// StartI empty description
func (self *WebGLSpriteBatch) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Destroy Destroys the SpriteBatch.
func (self *WebGLSpriteBatch) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the SpriteBatch.
func (self *WebGLSpriteBatch) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

