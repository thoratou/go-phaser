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
func (self *CanvasBuffer) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the Canvas in pixels.
func (self *CanvasBuffer) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the Canvas in pixels.
func (self *CanvasBuffer) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the Canvas in pixels.
func (self *CanvasBuffer) SetHeight(member float64) {
    self.Set("height", member)
}

// The Canvas object that belongs to this CanvasBuffer.
func (self *CanvasBuffer) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("canvas"))
}

// The Canvas object that belongs to this CanvasBuffer.
func (self *CanvasBuffer) SetCanvas(member dom.HTMLCanvasElement) {
    self.Set("canvas", member)
}

// A CanvasRenderingContext2D object representing a two-dimensional rendering context.
func (self *CanvasBuffer) GetContext() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Get("context"))
}

// A CanvasRenderingContext2D object representing a two-dimensional rendering context.
func (self *CanvasBuffer) SetContext(member dom.CanvasRenderingContext2D) {
    self.Set("context", member)
}



// Clears the canvas that was created by the CanvasBuffer class.
func (self *CanvasBuffer) ClearI(args ...interface{}) {
    self.Call("clear", args)
}

// Resizes the canvas to the specified width and height.
func (self *CanvasBuffer) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// Frees the canvas up for use again.
func (self *CanvasBuffer) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
