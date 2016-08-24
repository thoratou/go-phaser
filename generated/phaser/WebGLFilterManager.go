// Package phaser Automatic generation for PIXI.WebGLFilterManager
// generated file WebGLFilterManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// WebGLFilterManager empty description
type WebGLFilterManager struct {
    *js.Object
}

// NewWebGLFilterManager empty description
func NewWebGLFilterManager() *WebGLFilterManager {
    return &WebGLFilterManager{js.Global.Get("PIXI").Get("WebGLFilterManager").New()}
}
// NewWebGLFilterManagerI empty description
func NewWebGLFilterManagerI(args ...interface{}) *WebGLFilterManager {
    return &WebGLFilterManager{js.Global.Get("PIXI").Get("WebGLFilterManager").New(args)}
}



// FilterStack empty description
func (self *WebGLFilterManager) FilterStack() []interface{}{
	array00 := self.Object.Get("filterStack")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetFilterStackA empty description
func (self *WebGLFilterManager) SetFilterStackA(member []interface{}) {
    self.Object.Set("filterStack", member)
}

// OffsetX empty description
func (self *WebGLFilterManager) OffsetX() int{
    return self.Object.Get("offsetX").Int()
}

// SetOffsetXA empty description
func (self *WebGLFilterManager) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// OffsetY empty description
func (self *WebGLFilterManager) OffsetY() int{
    return self.Object.Get("offsetY").Int()
}

// SetOffsetYA empty description
func (self *WebGLFilterManager) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}


// SetContext Initialises the context and the properties.
func (self *WebGLFilterManager) SetContext(gl *WebGLContext) {
    self.Object.Call("setContext", gl)
}

// SetContextI Initialises the context and the properties.
func (self *WebGLFilterManager) SetContextI(args ...interface{}) {
    self.Object.Call("setContext", args)
}

// Begin empty description
func (self *WebGLFilterManager) Begin(renderSession *RenderSession, buffer *ArrayBuffer) {
    self.Object.Call("begin", renderSession, buffer)
}

// BeginI empty description
func (self *WebGLFilterManager) BeginI(args ...interface{}) {
    self.Object.Call("begin", args)
}

// PushFilter Applies the filter and adds it to the current filter stack.
func (self *WebGLFilterManager) PushFilter(filterBlock interface{}) {
    self.Object.Call("pushFilter", filterBlock)
}

// PushFilterI Applies the filter and adds it to the current filter stack.
func (self *WebGLFilterManager) PushFilterI(args ...interface{}) {
    self.Object.Call("pushFilter", args)
}

// PopFilter Removes the last filter from the filter stack and doesn't return it.
func (self *WebGLFilterManager) PopFilter() {
    self.Object.Call("popFilter")
}

// PopFilterI Removes the last filter from the filter stack and doesn't return it.
func (self *WebGLFilterManager) PopFilterI(args ...interface{}) {
    self.Object.Call("popFilter", args)
}

// ApplyFilterPass Applies the filter to the specified area.
func (self *WebGLFilterManager) ApplyFilterPass(filter *AbstractFilter, filterArea *Texture, width int, height int) {
    self.Object.Call("applyFilterPass", filter, filterArea, width, height)
}

// ApplyFilterPassI Applies the filter to the specified area.
func (self *WebGLFilterManager) ApplyFilterPassI(args ...interface{}) {
    self.Object.Call("applyFilterPass", args)
}

// InitShaderBuffers Initialises the shader buffers.
func (self *WebGLFilterManager) InitShaderBuffers() {
    self.Object.Call("initShaderBuffers")
}

// InitShaderBuffersI Initialises the shader buffers.
func (self *WebGLFilterManager) InitShaderBuffersI(args ...interface{}) {
    self.Object.Call("initShaderBuffers", args)
}

// Destroy Destroys the filter and removes it from the filter stack.
func (self *WebGLFilterManager) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the filter and removes it from the filter stack.
func (self *WebGLFilterManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

