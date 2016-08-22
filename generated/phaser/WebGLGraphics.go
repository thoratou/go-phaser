// Automatic generation for PIXI.WebGLGraphics
// generated file WebGLGraphics.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A set of functions used by the webGL renderer to draw the primitive graphics data
type WebGLGraphics struct {
    *js.Object
}




// Renders the graphics object
func (self *WebGLGraphics) RenderGraphicsI(args ...interface{}) {
    self.Call("renderGraphics", args)
}

// Updates the graphics object
func (self *WebGLGraphics) UpdateGraphicsI(args ...interface{}) {
    self.Call("updateGraphics", args)
}

// 
func (self *WebGLGraphics) SwitchModeI(args ...interface{}) {
    self.Call("switchMode", args)
}

// Builds a rectangle to draw
func (self *WebGLGraphics) BuildRectangleI(args ...interface{}) {
    self.Call("buildRectangle", args)
}

// Builds a rounded rectangle to draw
func (self *WebGLGraphics) BuildRoundedRectangleI(args ...interface{}) {
    self.Call("buildRoundedRectangle", args)
}

// Calculate the points for a quadratic bezier curve. (helper function..)
// Based on: https://stackoverflow.com/questions/785097/how-do-i-implement-a-bezier-curve-in-c
func (self *WebGLGraphics) QuadraticBezierCurveI(args ...interface{}) []int{
	array := self.Call("quadraticBezierCurve", args)
	length := array.Length()
	out := make([]int, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Int()
		
	}
	return out
}

// Builds a circle to draw
func (self *WebGLGraphics) BuildCircleI(args ...interface{}) {
    self.Call("buildCircle", args)
}

// Builds a line to draw
func (self *WebGLGraphics) BuildLineI(args ...interface{}) {
    self.Call("buildLine", args)
}

// Builds a complex polygon to draw
func (self *WebGLGraphics) BuildComplexPolyI(args ...interface{}) {
    self.Call("buildComplexPoly", args)
}

// Builds a polygon to draw
func (self *WebGLGraphics) BuildPolyI(args ...interface{}) {
    self.Call("buildPoly", args)
}
