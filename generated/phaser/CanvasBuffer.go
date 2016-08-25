// Package phaser Automatic generation for PIXI.CanvasBuffer
// generated file CanvasBuffer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// CanvasBuffer Creates a Canvas element of the given size.
type CanvasBuffer struct {
    *js.Object
}

// NewCanvasBuffer Creates a Canvas element of the given size.
func NewCanvasBuffer(width int, height int) *CanvasBuffer {
    return &CanvasBuffer{js.Global.Get("PIXI").Get("CanvasBuffer").New(width, height)}
}
// NewCanvasBufferI Creates a Canvas element of the given size.
func NewCanvasBufferI(args ...interface{}) *CanvasBuffer {
    return &CanvasBuffer{js.Global.Get("PIXI").Get("CanvasBuffer").New(args)}
}



// CanvasBuffer Binding conversion method to CanvasBuffer point 
func ToCanvasBuffer(jsStruct interface{}) *CanvasBuffer {
    if object, ok := jsStruct.(*js.Object); ok {
		return &CanvasBuffer{Object: object}
	}
	return nil
}



// Width The width of the Canvas in pixels.
func (self *CanvasBuffer) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the Canvas in pixels.
func (self *CanvasBuffer) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the Canvas in pixels.
func (self *CanvasBuffer) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the Canvas in pixels.
func (self *CanvasBuffer) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Canvas The Canvas object that belongs to this CanvasBuffer.
func (self *CanvasBuffer) Canvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("canvas"))
}

// SetCanvasA The Canvas object that belongs to this CanvasBuffer.
func (self *CanvasBuffer) SetCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("canvas", member)
}

// Context A CanvasRenderingContext2D object representing a two-dimensional rendering context.
func (self *CanvasBuffer) Context() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("context"))
}

// SetContextA A CanvasRenderingContext2D object representing a two-dimensional rendering context.
func (self *CanvasBuffer) SetContextA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("context", member)
}


// Clear Clears the canvas that was created by the CanvasBuffer class.
func (self *CanvasBuffer) Clear() {
    self.Object.Call("clear")
}

// ClearI Clears the canvas that was created by the CanvasBuffer class.
func (self *CanvasBuffer) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// Resize Resizes the canvas to the specified width and height.
func (self *CanvasBuffer) Resize(width int, height int) {
    self.Object.Call("resize", width, height)
}

// ResizeI Resizes the canvas to the specified width and height.
func (self *CanvasBuffer) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Destroy Frees the canvas up for use again.
func (self *CanvasBuffer) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Frees the canvas up for use again.
func (self *CanvasBuffer) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

