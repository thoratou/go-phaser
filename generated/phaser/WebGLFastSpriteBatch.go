// Automatic generation for PIXI.WebGLFastSpriteBatch
// generated file WebGLFastSpriteBatch.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type WebGLFastSpriteBatch struct {
    *js.Object
}


// 
func NewWebGLFastSpriteBatch() *WebGLFastSpriteBatch {
    return &WebGLFastSpriteBatch{js.Global.Get("PIXI").Get("WebGLFastSpriteBatch").New()}
}

// 
func NewWebGLFastSpriteBatchI(args ...interface{}) *WebGLFastSpriteBatch {
    return &WebGLFastSpriteBatch{js.Global.Get("PIXI").Get("WebGLFastSpriteBatch").New(args)}
}



// 
func (self *WebGLFastSpriteBatch) VertSize() int{
    return self.Object.Get("vertSize").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetVertSizeA(member int) {
    self.Object.Set("vertSize", member)
}

// 
func (self *WebGLFastSpriteBatch) MaxSize() int{
    return self.Object.Get("maxSize").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetMaxSizeA(member int) {
    self.Object.Set("maxSize", member)
}

// 
func (self *WebGLFastSpriteBatch) Size() int{
    return self.Object.Get("size").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetSizeA(member int) {
    self.Object.Set("size", member)
}

// Vertex data
func (self *WebGLFastSpriteBatch) Vertices() *Float32Array{
    return &Float32Array{self.Object.Get("vertices")}
}

// Vertex data
func (self *WebGLFastSpriteBatch) SetVerticesA(member *Float32Array) {
    self.Object.Set("vertices", member)
}

// Index data
func (self *WebGLFastSpriteBatch) Indices() *Uint16Array{
    return &Uint16Array{self.Object.Get("indices")}
}

// Index data
func (self *WebGLFastSpriteBatch) SetIndicesA(member *Uint16Array) {
    self.Object.Set("indices", member)
}

// 
func (self *WebGLFastSpriteBatch) VertexBuffer() interface{}{
    return self.Object.Get("vertexBuffer")
}

// 
func (self *WebGLFastSpriteBatch) SetVertexBufferA(member interface{}) {
    self.Object.Set("vertexBuffer", member)
}

// 
func (self *WebGLFastSpriteBatch) IndexBuffer() interface{}{
    return self.Object.Get("indexBuffer")
}

// 
func (self *WebGLFastSpriteBatch) SetIndexBufferA(member interface{}) {
    self.Object.Set("indexBuffer", member)
}

// 
func (self *WebGLFastSpriteBatch) LastIndexCount() int{
    return self.Object.Get("lastIndexCount").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetLastIndexCountA(member int) {
    self.Object.Set("lastIndexCount", member)
}

// 
func (self *WebGLFastSpriteBatch) Drawing() bool{
    return self.Object.Get("drawing").Bool()
}

// 
func (self *WebGLFastSpriteBatch) SetDrawingA(member bool) {
    self.Object.Set("drawing", member)
}

// 
func (self *WebGLFastSpriteBatch) CurrentBatchSize() int{
    return self.Object.Get("currentBatchSize").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetCurrentBatchSizeA(member int) {
    self.Object.Set("currentBatchSize", member)
}

// 
func (self *WebGLFastSpriteBatch) CurrentBaseTexture() *BaseTexture{
    return &BaseTexture{self.Object.Get("currentBaseTexture")}
}

// 
func (self *WebGLFastSpriteBatch) SetCurrentBaseTextureA(member *BaseTexture) {
    self.Object.Set("currentBaseTexture", member)
}

// 
func (self *WebGLFastSpriteBatch) CurrentBlendMode() int{
    return self.Object.Get("currentBlendMode").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetCurrentBlendModeA(member int) {
    self.Object.Set("currentBlendMode", member)
}

// 
func (self *WebGLFastSpriteBatch) RenderSession() interface{}{
    return self.Object.Get("renderSession")
}

// 
func (self *WebGLFastSpriteBatch) SetRenderSessionA(member interface{}) {
    self.Object.Set("renderSession", member)
}

// 
func (self *WebGLFastSpriteBatch) Shader() interface{}{
    return self.Object.Get("shader")
}

// 
func (self *WebGLFastSpriteBatch) SetShaderA(member interface{}) {
    self.Object.Set("shader", member)
}

// 
func (self *WebGLFastSpriteBatch) Matrix() *Matrix{
    return &Matrix{self.Object.Get("matrix")}
}

// 
func (self *WebGLFastSpriteBatch) SetMatrixA(member *Matrix) {
    self.Object.Set("matrix", member)
}



// Sets the WebGL Context.
func (self *WebGLFastSpriteBatch) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// Sets the WebGL Context.
func (self *WebGLFastSpriteBatch) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// 
func (self *WebGLFastSpriteBatch) Begin(spriteBatch *WebGLSpriteBatch, renderSession interface{}) {
    self.Object.Call("begin", spriteBatch, renderSession)
}

// 
func (self *WebGLFastSpriteBatch) BeginI(args ...interface{}) {
    self.Object.Call("begin", args)
}

// 
func (self *WebGLFastSpriteBatch) End() {
    self.Object.Call("end")
}

// 
func (self *WebGLFastSpriteBatch) EndI(args ...interface{}) {
    self.Object.Call("end", args)
}

// 
func (self *WebGLFastSpriteBatch) Render(spriteBatch *WebGLSpriteBatch) {
    self.Object.Call("render", spriteBatch)
}

// 
func (self *WebGLFastSpriteBatch) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// 
func (self *WebGLFastSpriteBatch) RenderSprite(sprite *Sprite) {
    self.Object.Call("renderSprite", sprite)
}

// 
func (self *WebGLFastSpriteBatch) RenderSpriteI(args ...interface{}) {
    self.Object.Call("renderSprite", args)
}

// 
func (self *WebGLFastSpriteBatch) Flush() {
    self.Object.Call("flush")
}

// 
func (self *WebGLFastSpriteBatch) FlushI(args ...interface{}) {
    self.Object.Call("flush", args)
}

// 
func (self *WebGLFastSpriteBatch) Stop() {
    self.Object.Call("stop")
}

// 
func (self *WebGLFastSpriteBatch) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// 
func (self *WebGLFastSpriteBatch) Start() {
    self.Object.Call("start")
}

// 
func (self *WebGLFastSpriteBatch) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}
