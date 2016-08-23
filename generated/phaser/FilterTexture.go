// Automatic generation for PIXI.FilterTexture
// generated file FilterTexture.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type FilterTexture struct {
    *js.Object
}


// 
func NewFilterTexture(gl *WebGLContext, width int, height int, scaleMode int) *FilterTexture {
    return &FilterTexture{js.Global.Get("PIXI").Get("FilterTexture").New(gl, width, height, scaleMode)}
}

// 
func NewFilterTextureI(args ...interface{}) *FilterTexture {
    return &FilterTexture{js.Global.Get("PIXI").Get("FilterTexture").New(args)}
}



// 
func (self *FilterTexture) Gl() WebGLContext{
    return WrapWebGLContext(self.Object.Get("gl"))
}

// 
func (self *FilterTexture) SetGlA(member WebGLContext) {
    self.Object.Set("gl", member)
}

// 
func (self *FilterTexture) FrameBuffer() interface{}{
    return self.Object.Get("frameBuffer")
}

// 
func (self *FilterTexture) SetFrameBufferA(member interface{}) {
    self.Object.Set("frameBuffer", member)
}

// 
func (self *FilterTexture) Texture() interface{}{
    return self.Object.Get("texture")
}

// 
func (self *FilterTexture) SetTextureA(member interface{}) {
    self.Object.Set("texture", member)
}

// 
func (self *FilterTexture) ScaleMode() int{
    return self.Object.Get("scaleMode").Int()
}

// 
func (self *FilterTexture) SetScaleModeA(member int) {
    self.Object.Set("scaleMode", member)
}



// Clears the filter texture.
func (self *FilterTexture) Clear() {
    self.Object.Call("clear")
}

// Clears the filter texture.
func (self *FilterTexture) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// Resizes the texture to the specified width and height
func (self *FilterTexture) Resize(width int, height int) {
    self.Object.Call("resize", width, height)
}

// Resizes the texture to the specified width and height
func (self *FilterTexture) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Destroys the filter texture.
func (self *FilterTexture) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the filter texture.
func (self *FilterTexture) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
