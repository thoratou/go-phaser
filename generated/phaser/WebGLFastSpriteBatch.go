// Package phaser Automatic generation for PIXI.WebGLFastSpriteBatch
// generated file WebGLFastSpriteBatch.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// WebGLFastSpriteBatch empty description
type WebGLFastSpriteBatch struct {
    *js.Object
}

// NewWebGLFastSpriteBatch empty description
func NewWebGLFastSpriteBatch() *WebGLFastSpriteBatch {
    return &WebGLFastSpriteBatch{js.Global.Get("PIXI").Get("WebGLFastSpriteBatch").New()}
}
// NewWebGLFastSpriteBatchI empty description
func NewWebGLFastSpriteBatchI(args ...interface{}) *WebGLFastSpriteBatch {
    return &WebGLFastSpriteBatch{js.Global.Get("PIXI").Get("WebGLFastSpriteBatch").New(args)}
}



// WebGLFastSpriteBatch Binding conversion method to WebGLFastSpriteBatch point 
func ToWebGLFastSpriteBatch(jsStruct interface{}) *WebGLFastSpriteBatch {
    if object, ok := jsStruct.(*js.Object); ok {
		return &WebGLFastSpriteBatch{Object: object}
	}
	return nil
}



// VertSize empty description
func (self *WebGLFastSpriteBatch) VertSize() int{
    return self.Object.Get("vertSize").Int()
}

// SetVertSizeA empty description
func (self *WebGLFastSpriteBatch) SetVertSizeA(member int) {
    self.Object.Set("vertSize", member)
}

// MaxSize empty description
func (self *WebGLFastSpriteBatch) MaxSize() int{
    return self.Object.Get("maxSize").Int()
}

// SetMaxSizeA empty description
func (self *WebGLFastSpriteBatch) SetMaxSizeA(member int) {
    self.Object.Set("maxSize", member)
}

// Size empty description
func (self *WebGLFastSpriteBatch) Size() int{
    return self.Object.Get("size").Int()
}

// SetSizeA empty description
func (self *WebGLFastSpriteBatch) SetSizeA(member int) {
    self.Object.Set("size", member)
}

// Vertices Vertex data
func (self *WebGLFastSpriteBatch) Vertices() *Float32Array{
    return &Float32Array{self.Object.Get("vertices")}
}

// SetVerticesA Vertex data
func (self *WebGLFastSpriteBatch) SetVerticesA(member *Float32Array) {
    self.Object.Set("vertices", member)
}

// Indices Index data
func (self *WebGLFastSpriteBatch) Indices() *Uint16Array{
    return &Uint16Array{self.Object.Get("indices")}
}

// SetIndicesA Index data
func (self *WebGLFastSpriteBatch) SetIndicesA(member *Uint16Array) {
    self.Object.Set("indices", member)
}

// VertexBuffer empty description
func (self *WebGLFastSpriteBatch) VertexBuffer() interface{}{
    return self.Object.Get("vertexBuffer")
}

// SetVertexBufferA empty description
func (self *WebGLFastSpriteBatch) SetVertexBufferA(member interface{}) {
    self.Object.Set("vertexBuffer", member)
}

// IndexBuffer empty description
func (self *WebGLFastSpriteBatch) IndexBuffer() interface{}{
    return self.Object.Get("indexBuffer")
}

// SetIndexBufferA empty description
func (self *WebGLFastSpriteBatch) SetIndexBufferA(member interface{}) {
    self.Object.Set("indexBuffer", member)
}

// LastIndexCount empty description
func (self *WebGLFastSpriteBatch) LastIndexCount() int{
    return self.Object.Get("lastIndexCount").Int()
}

// SetLastIndexCountA empty description
func (self *WebGLFastSpriteBatch) SetLastIndexCountA(member int) {
    self.Object.Set("lastIndexCount", member)
}

// Drawing empty description
func (self *WebGLFastSpriteBatch) Drawing() bool{
    return self.Object.Get("drawing").Bool()
}

// SetDrawingA empty description
func (self *WebGLFastSpriteBatch) SetDrawingA(member bool) {
    self.Object.Set("drawing", member)
}

// CurrentBatchSize empty description
func (self *WebGLFastSpriteBatch) CurrentBatchSize() int{
    return self.Object.Get("currentBatchSize").Int()
}

// SetCurrentBatchSizeA empty description
func (self *WebGLFastSpriteBatch) SetCurrentBatchSizeA(member int) {
    self.Object.Set("currentBatchSize", member)
}

// CurrentBaseTexture empty description
func (self *WebGLFastSpriteBatch) CurrentBaseTexture() *BaseTexture{
    return &BaseTexture{self.Object.Get("currentBaseTexture")}
}

// SetCurrentBaseTextureA empty description
func (self *WebGLFastSpriteBatch) SetCurrentBaseTextureA(member *BaseTexture) {
    self.Object.Set("currentBaseTexture", member)
}

// CurrentBlendMode empty description
func (self *WebGLFastSpriteBatch) CurrentBlendMode() int{
    return self.Object.Get("currentBlendMode").Int()
}

// SetCurrentBlendModeA empty description
func (self *WebGLFastSpriteBatch) SetCurrentBlendModeA(member int) {
    self.Object.Set("currentBlendMode", member)
}

// RenderSession empty description
func (self *WebGLFastSpriteBatch) RenderSession() interface{}{
    return self.Object.Get("renderSession")
}

// SetRenderSessionA empty description
func (self *WebGLFastSpriteBatch) SetRenderSessionA(member interface{}) {
    self.Object.Set("renderSession", member)
}

// Shader empty description
func (self *WebGLFastSpriteBatch) Shader() interface{}{
    return self.Object.Get("shader")
}

// SetShaderA empty description
func (self *WebGLFastSpriteBatch) SetShaderA(member interface{}) {
    self.Object.Set("shader", member)
}

// Matrix empty description
func (self *WebGLFastSpriteBatch) Matrix() *Matrix{
    return &Matrix{self.Object.Get("matrix")}
}

// SetMatrixA empty description
func (self *WebGLFastSpriteBatch) SetMatrixA(member *Matrix) {
    self.Object.Set("matrix", member)
}


// SetContext Sets the WebGL Context.
func (self *WebGLFastSpriteBatch) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// SetContextI Sets the WebGL Context.
func (self *WebGLFastSpriteBatch) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// Begin empty description
func (self *WebGLFastSpriteBatch) Begin(spriteBatch *WebGLSpriteBatch, renderSession interface{}) {
    self.Object.Call("begin", spriteBatch, renderSession)
}

// BeginI empty description
func (self *WebGLFastSpriteBatch) BeginI(args ...interface{}) {
    self.Object.Call("begin", args)
}

// End empty description
func (self *WebGLFastSpriteBatch) End() {
    self.Object.Call("end")
}

// EndI empty description
func (self *WebGLFastSpriteBatch) EndI(args ...interface{}) {
    self.Object.Call("end", args)
}

// Render empty description
func (self *WebGLFastSpriteBatch) Render(spriteBatch *WebGLSpriteBatch) {
    self.Object.Call("render", spriteBatch)
}

// RenderI empty description
func (self *WebGLFastSpriteBatch) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// RenderSprite empty description
func (self *WebGLFastSpriteBatch) RenderSprite(sprite *Sprite) {
    self.Object.Call("renderSprite", sprite)
}

// RenderSpriteI empty description
func (self *WebGLFastSpriteBatch) RenderSpriteI(args ...interface{}) {
    self.Object.Call("renderSprite", args)
}

// Flush empty description
func (self *WebGLFastSpriteBatch) Flush() {
    self.Object.Call("flush")
}

// FlushI empty description
func (self *WebGLFastSpriteBatch) FlushI(args ...interface{}) {
    self.Object.Call("flush", args)
}

// Stop empty description
func (self *WebGLFastSpriteBatch) Stop() {
    self.Object.Call("stop")
}

// StopI empty description
func (self *WebGLFastSpriteBatch) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Start empty description
func (self *WebGLFastSpriteBatch) Start() {
    self.Object.Call("start")
}

// StartI empty description
func (self *WebGLFastSpriteBatch) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

