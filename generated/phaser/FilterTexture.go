// Package phaser Automatic generation for PIXI.FilterTexture
// generated file FilterTexture.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// FilterTexture empty description
type FilterTexture struct {
    *js.Object
}

// NewFilterTexture empty description
func NewFilterTexture(gl *WebGLContext, width int, height int, scaleMode int) *FilterTexture {
    return &FilterTexture{js.Global.Get("PIXI").Get("FilterTexture").New(gl, width, height, scaleMode)}
}
// NewFilterTextureI empty description
func NewFilterTextureI(args ...interface{}) *FilterTexture {
    return &FilterTexture{js.Global.Get("PIXI").Get("FilterTexture").New(args)}
}



// FilterTexture Binding conversion method to FilterTexture point 
func ToFilterTexture(jsStruct interface{}) *FilterTexture {
    if object, ok := jsStruct.(*js.Object); ok {
		return &FilterTexture{Object: object}
	}
	return nil
}



// Gl empty description
func (self *FilterTexture) Gl() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// SetGlA empty description
func (self *FilterTexture) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// FrameBuffer empty description
func (self *FilterTexture) FrameBuffer() interface{}{
    return self.Object.Get("frameBuffer")
}

// SetFrameBufferA empty description
func (self *FilterTexture) SetFrameBufferA(member interface{}) {
    self.Object.Set("frameBuffer", member)
}

// Texture empty description
func (self *FilterTexture) Texture() interface{}{
    return self.Object.Get("texture")
}

// SetTextureA empty description
func (self *FilterTexture) SetTextureA(member interface{}) {
    self.Object.Set("texture", member)
}

// ScaleMode empty description
func (self *FilterTexture) ScaleMode() int{
    return self.Object.Get("scaleMode").Int()
}

// SetScaleModeA empty description
func (self *FilterTexture) SetScaleModeA(member int) {
    self.Object.Set("scaleMode", member)
}


// Clear Clears the filter texture.
func (self *FilterTexture) Clear() {
    self.Object.Call("clear")
}

// ClearI Clears the filter texture.
func (self *FilterTexture) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// Resize Resizes the texture to the specified width and height
func (self *FilterTexture) Resize(width int, height int) {
    self.Object.Call("resize", width, height)
}

// ResizeI Resizes the texture to the specified width and height
func (self *FilterTexture) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Destroy Destroys the filter texture.
func (self *FilterTexture) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the filter texture.
func (self *FilterTexture) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

