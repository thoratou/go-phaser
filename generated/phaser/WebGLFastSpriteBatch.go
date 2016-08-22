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
func (self *WebGLFastSpriteBatch) GetVertSize() int{
    return self.Get("vertSize").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetVertSize(member int) {
    self.Set("vertSize", member)
}

// 
func (self *WebGLFastSpriteBatch) GetMaxSize() int{
    return self.Get("maxSize").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetMaxSize(member int) {
    self.Set("maxSize", member)
}

// 
func (self *WebGLFastSpriteBatch) GetSize() int{
    return self.Get("size").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetSize(member int) {
    self.Set("size", member)
}

// Vertex data
func (self *WebGLFastSpriteBatch) GetVertices() *Float32Array{
    return &Float32Array{self.Get("vertices")}
}

// Vertex data
func (self *WebGLFastSpriteBatch) SetVertices(member *Float32Array) {
    self.Set("vertices", member)
}

// Index data
func (self *WebGLFastSpriteBatch) GetIndices() *Uint16Array{
    return &Uint16Array{self.Get("indices")}
}

// Index data
func (self *WebGLFastSpriteBatch) SetIndices(member *Uint16Array) {
    self.Set("indices", member)
}

// 
func (self *WebGLFastSpriteBatch) GetVertexBuffer() interface{}{
    return self.Get("vertexBuffer")
}

// 
func (self *WebGLFastSpriteBatch) SetVertexBuffer(member interface{}) {
    self.Set("vertexBuffer", member)
}

// 
func (self *WebGLFastSpriteBatch) GetIndexBuffer() interface{}{
    return self.Get("indexBuffer")
}

// 
func (self *WebGLFastSpriteBatch) SetIndexBuffer(member interface{}) {
    self.Set("indexBuffer", member)
}

// 
func (self *WebGLFastSpriteBatch) GetLastIndexCount() int{
    return self.Get("lastIndexCount").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetLastIndexCount(member int) {
    self.Set("lastIndexCount", member)
}

// 
func (self *WebGLFastSpriteBatch) GetDrawing() bool{
    return self.Get("drawing").Bool()
}

// 
func (self *WebGLFastSpriteBatch) SetDrawing(member bool) {
    self.Set("drawing", member)
}

// 
func (self *WebGLFastSpriteBatch) GetCurrentBatchSize() int{
    return self.Get("currentBatchSize").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetCurrentBatchSize(member int) {
    self.Set("currentBatchSize", member)
}

// 
func (self *WebGLFastSpriteBatch) GetCurrentBaseTexture() *BaseTexture{
    return &BaseTexture{self.Get("currentBaseTexture")}
}

// 
func (self *WebGLFastSpriteBatch) SetCurrentBaseTexture(member *BaseTexture) {
    self.Set("currentBaseTexture", member)
}

// 
func (self *WebGLFastSpriteBatch) GetCurrentBlendMode() int{
    return self.Get("currentBlendMode").Int()
}

// 
func (self *WebGLFastSpriteBatch) SetCurrentBlendMode(member int) {
    self.Set("currentBlendMode", member)
}

// 
func (self *WebGLFastSpriteBatch) GetRenderSession() interface{}{
    return self.Get("renderSession")
}

// 
func (self *WebGLFastSpriteBatch) SetRenderSession(member interface{}) {
    self.Set("renderSession", member)
}

// 
func (self *WebGLFastSpriteBatch) GetShader() interface{}{
    return self.Get("shader")
}

// 
func (self *WebGLFastSpriteBatch) SetShader(member interface{}) {
    self.Set("shader", member)
}

// 
func (self *WebGLFastSpriteBatch) GetMatrix() *Matrix{
    return &Matrix{self.Get("matrix")}
}

// 
func (self *WebGLFastSpriteBatch) SetMatrix(member *Matrix) {
    self.Set("matrix", member)
}



// Sets the WebGL Context.
func (self *WebGLFastSpriteBatch) SetContextI(args ...interface{}) {
    self.Call("setContext", args)
}

// 
func (self *WebGLFastSpriteBatch) BeginI(args ...interface{}) {
    self.Call("begin", args)
}

// 
func (self *WebGLFastSpriteBatch) EndI(args ...interface{}) {
    self.Call("end", args)
}

// 
func (self *WebGLFastSpriteBatch) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// 
func (self *WebGLFastSpriteBatch) RenderSpriteI(args ...interface{}) {
    self.Call("renderSprite", args)
}

// 
func (self *WebGLFastSpriteBatch) FlushI(args ...interface{}) {
    self.Call("flush", args)
}

// 
func (self *WebGLFastSpriteBatch) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// 
func (self *WebGLFastSpriteBatch) StartI(args ...interface{}) {
    self.Call("start", args)
}
