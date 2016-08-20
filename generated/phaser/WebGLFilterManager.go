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
func (self *WebGLFilterManager) GetFilterStack() []interface{}{
	array := self.Get("filterStack")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// 
func (self *WebGLFilterManager) SetFilterStack(member []interface{}) {
    self.Set("filterStack", member)
}

// 
func (self *WebGLFilterManager) GetOffsetX() float64{
    return self.Get("offsetX").Float()
}

// 
func (self *WebGLFilterManager) SetOffsetX(member float64) {
    self.Set("offsetX", member)
}

// 
func (self *WebGLFilterManager) GetOffsetY() float64{
    return self.Get("offsetY").Float()
}

// 
func (self *WebGLFilterManager) SetOffsetY(member float64) {
    self.Set("offsetY", member)
}



// Initialises the context and the properties.
func (self *WebGLFilterManager) SetContextI(args ...interface{}) {
    self.Call("setContext", args)
}

// 
func (self *WebGLFilterManager) BeginI(args ...interface{}) {
    self.Call("begin", args)
}

// Applies the filter and adds it to the current filter stack.
func (self *WebGLFilterManager) PushFilterI(args ...interface{}) {
    self.Call("pushFilter", args)
}

// Removes the last filter from the filter stack and doesn't return it.
func (self *WebGLFilterManager) PopFilterI(args ...interface{}) {
    self.Call("popFilter", args)
}

// Applies the filter to the specified area.
func (self *WebGLFilterManager) ApplyFilterPassI(args ...interface{}) {
    self.Call("applyFilterPass", args)
}

// Initialises the shader buffers.
func (self *WebGLFilterManager) InitShaderBuffersI(args ...interface{}) {
    self.Call("initShaderBuffers", args)
}

// Destroys the filter and removes it from the filter stack.
func (self *WebGLFilterManager) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
