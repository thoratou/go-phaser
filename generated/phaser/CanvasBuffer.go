// Automatic generation for PIXI.CanvasBuffer
// generated file CanvasBuffer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// Creates a Canvas element of the given size.
type CanvasBuffer struct {
    *js.Object
}


// The width of the Canvas in pixels.
func (self *CanvasBuffer) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the Canvas in pixels.
func (self *CanvasBuffer) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the Canvas in pixels.
func (self *CanvasBuffer) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the Canvas in pixels.
func (self *CanvasBuffer) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The Canvas object that belongs to this CanvasBuffer.
func (self *CanvasBuffer) GetCanvasA() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("canvas"))
}

// The Canvas object that belongs to this CanvasBuffer.
func (self *CanvasBuffer) SetCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("canvas", member)
}

// A CanvasRenderingContext2D object representing a two-dimensional rendering context.
func (self *CanvasBuffer) GetContextA() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("context"))
}

// A CanvasRenderingContext2D object representing a two-dimensional rendering context.
func (self *CanvasBuffer) SetContextA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("context", member)
}



// Clears the canvas that was created by the CanvasBuffer class.
func (self *CanvasBuffer) Clear() {
    self.Object.Call("clear")
}

// Clears the canvas that was created by the CanvasBuffer class.
func (self *CanvasBuffer) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// Resizes the canvas to the specified width and height.
func (self *CanvasBuffer) Resize(width int, height int) {
    self.Object.Call("resize", width, height)
}

// Resizes the canvas to the specified width and height.
func (self *CanvasBuffer) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Frees the canvas up for use again.
func (self *CanvasBuffer) Destroy() {
    self.Object.Call("destroy")
}

// Frees the canvas up for use again.
func (self *CanvasBuffer) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
