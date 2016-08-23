// Automatic generation for PIXI.WebGLFilterManager
// generated file WebGLFilterManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type WebGLFilterManager struct {
    *js.Object
}


// 
func NewWebGLFilterManager() *WebGLFilterManager {
    return &WebGLFilterManager{js.Global.Get("PIXI").Get("WebGLFilterManager").New()}
}

// 
func NewWebGLFilterManagerI(args ...interface{}) *WebGLFilterManager {
    return &WebGLFilterManager{js.Global.Get("PIXI").Get("WebGLFilterManager").New(args)}
}



// 
func (self *WebGLFilterManager) FilterStack() []interface{}{
	array00 := self.Object.Get("filterStack")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// 
func (self *WebGLFilterManager) SetFilterStackA(member []interface{}) {
    self.Object.Set("filterStack", member)
}

// 
func (self *WebGLFilterManager) OffsetX() int{
    return self.Object.Get("offsetX").Int()
}

// 
func (self *WebGLFilterManager) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// 
func (self *WebGLFilterManager) OffsetY() int{
    return self.Object.Get("offsetY").Int()
}

// 
func (self *WebGLFilterManager) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}



// Initialises the context and the properties.
func (self *WebGLFilterManager) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// Initialises the context and the properties.
func (self *WebGLFilterManager) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// 
func (self *WebGLFilterManager) Begin(renderSession *RenderSession, buffer *ArrayBuffer) {
    self.Object.Call("begin", renderSession, buffer)
}

// 
func (self *WebGLFilterManager) BeginI(args ...interface{}) {
    self.Object.Call("begin", args)
}

// Applies the filter and adds it to the current filter stack.
func (self *WebGLFilterManager) PushFilter(filterBlock interface{}) {
    self.Object.Call("pushFilter", filterBlock)
}

// Applies the filter and adds it to the current filter stack.
func (self *WebGLFilterManager) PushFilterI(args ...interface{}) {
    self.Object.Call("pushFilter", args)
}

// Removes the last filter from the filter stack and doesn't return it.
func (self *WebGLFilterManager) PopFilter() {
    self.Object.Call("popFilter")
}

// Removes the last filter from the filter stack and doesn't return it.
func (self *WebGLFilterManager) PopFilterI(args ...interface{}) {
    self.Object.Call("popFilter", args)
}

// Applies the filter to the specified area.
func (self *WebGLFilterManager) ApplyFilterPass(filter *AbstractFilter, filterArea *Texture, width int, height int) {
    self.Object.Call("applyFilterPass", filter, filterArea, width, height)
}

// Applies the filter to the specified area.
func (self *WebGLFilterManager) ApplyFilterPassI(args ...interface{}) {
    self.Object.Call("applyFilterPass", args)
}

// Initialises the shader buffers.
func (self *WebGLFilterManager) InitShaderBuffers() {
    self.Object.Call("initShaderBuffers")
}

// Initialises the shader buffers.
func (self *WebGLFilterManager) InitShaderBuffersI(args ...interface{}) {
    self.Object.Call("initShaderBuffers", args)
}

// Destroys the filter and removes it from the filter stack.
func (self *WebGLFilterManager) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the filter and removes it from the filter stack.
func (self *WebGLFilterManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
