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
func (self *FilterTexture) GetGl() WebGLContext{
    return WrapWebGLContext(self.Get("gl"))
}

// 
func (self *FilterTexture) SetGl(member WebGLContext) {
    self.Set("gl", member)
}

// 
func (self *FilterTexture) GetFrameBuffer() interface{}{
    return self.Get("frameBuffer")
}

// 
func (self *FilterTexture) SetFrameBuffer(member interface{}) {
    self.Set("frameBuffer", member)
}

// 
func (self *FilterTexture) GetTexture() interface{}{
    return self.Get("texture")
}

// 
func (self *FilterTexture) SetTexture(member interface{}) {
    self.Set("texture", member)
}

// 
func (self *FilterTexture) GetScaleMode() int{
    return self.Get("scaleMode").Int()
}

// 
func (self *FilterTexture) SetScaleMode(member int) {
    self.Set("scaleMode", member)
}



// Clears the filter texture.
func (self *FilterTexture) ClearI(args ...interface{}) {
    self.Call("clear", args)
}

// Resizes the texture to the specified width and height
func (self *FilterTexture) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// Destroys the filter texture.
func (self *FilterTexture) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
