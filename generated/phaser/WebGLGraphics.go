// Package phaser Automatic generation for PIXI.WebGLGraphics
// generated file WebGLGraphics.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// WebGLGraphics A set of functions used by the webGL renderer to draw the primitive graphics data
type WebGLGraphics struct {
    *js.Object
}

// NewWebGLGraphics A set of functions used by the webGL renderer to draw the primitive graphics data
func NewWebGLGraphics() *WebGLGraphics {
    return &WebGLGraphics{js.Global.Get("PIXI").Get("WebGLGraphics").New()}
}
// NewWebGLGraphicsI A set of functions used by the webGL renderer to draw the primitive graphics data
func NewWebGLGraphicsI(args ...interface{}) *WebGLGraphics {
    return &WebGLGraphics{js.Global.Get("PIXI").Get("WebGLGraphics").New(args)}
}



// WebGLGraphics Binding conversion method to WebGLGraphics point 
func ToWebGLGraphics(jsStruct interface{}) *WebGLGraphics {
    if object, ok := jsStruct.(*js.Object); ok {
		return &WebGLGraphics{Object: object}
	}
	return nil
}




// RenderGraphics Renders the graphics object
func (self *WebGLGraphics) RenderGraphics(graphics *Graphics, renderSession interface{}) {
    self.Object.Call("renderGraphics", graphics, renderSession)
}

// RenderGraphicsI Renders the graphics object
func (self *WebGLGraphics) RenderGraphicsI(args ...interface{}) {
    self.Object.Call("renderGraphics", args)
}

// UpdateGraphics Updates the graphics object
func (self *WebGLGraphics) UpdateGraphics(graphicsData *Graphics, gl *WebGLContext) {
    self.Object.Call("updateGraphics", graphicsData, gl)
}

// UpdateGraphicsI Updates the graphics object
func (self *WebGLGraphics) UpdateGraphicsI(args ...interface{}) {
    self.Object.Call("updateGraphics", args)
}

// SwitchMode empty description
func (self *WebGLGraphics) SwitchMode(webGL *WebGLContext, type_ int) {
    self.Object.Call("switchMode", webGL, type_)
}

// SwitchModeI empty description
func (self *WebGLGraphics) SwitchModeI(args ...interface{}) {
    self.Object.Call("switchMode", args)
}

// BuildRectangle Builds a rectangle to draw
func (self *WebGLGraphics) BuildRectangle(graphicsData *Graphics, webGLData interface{}) {
    self.Object.Call("buildRectangle", graphicsData, webGLData)
}

// BuildRectangleI Builds a rectangle to draw
func (self *WebGLGraphics) BuildRectangleI(args ...interface{}) {
    self.Object.Call("buildRectangle", args)
}

// BuildRoundedRectangle Builds a rounded rectangle to draw
func (self *WebGLGraphics) BuildRoundedRectangle(graphicsData *Graphics, webGLData interface{}) {
    self.Object.Call("buildRoundedRectangle", graphicsData, webGLData)
}

// BuildRoundedRectangleI Builds a rounded rectangle to draw
func (self *WebGLGraphics) BuildRoundedRectangleI(args ...interface{}) {
    self.Object.Call("buildRoundedRectangle", args)
}

// QuadraticBezierCurve Calculate the points for a quadratic bezier curve. (helper function..)
// Based on: https://stackoverflow.com/questions/785097/how-do-i-implement-a-bezier-curve-in-c
func (self *WebGLGraphics) QuadraticBezierCurve(fromX int, fromY int, cpX int, cpY int, toX int, toY int) []int{
	array00 := self.Object.Call("quadraticBezierCurve", fromX, fromY, cpX, cpY, toX, toY)
	length00 := array00.Length()
	out00 := make([]int, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Int()
		
	}
	return out00
}

// QuadraticBezierCurveI Calculate the points for a quadratic bezier curve. (helper function..)
// Based on: https://stackoverflow.com/questions/785097/how-do-i-implement-a-bezier-curve-in-c
func (self *WebGLGraphics) QuadraticBezierCurveI(args ...interface{}) []int{
	array00 := self.Object.Call("quadraticBezierCurve", args)
	length00 := array00.Length()
	out00 := make([]int, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Int()
		
	}
	return out00
}

// BuildCircle Builds a circle to draw
func (self *WebGLGraphics) BuildCircle(graphicsData *Graphics, webGLData interface{}) {
    self.Object.Call("buildCircle", graphicsData, webGLData)
}

// BuildCircleI Builds a circle to draw
func (self *WebGLGraphics) BuildCircleI(args ...interface{}) {
    self.Object.Call("buildCircle", args)
}

// BuildLine Builds a line to draw
func (self *WebGLGraphics) BuildLine(graphicsData *Graphics, webGLData interface{}) {
    self.Object.Call("buildLine", graphicsData, webGLData)
}

// BuildLineI Builds a line to draw
func (self *WebGLGraphics) BuildLineI(args ...interface{}) {
    self.Object.Call("buildLine", args)
}

// BuildComplexPoly Builds a complex polygon to draw
func (self *WebGLGraphics) BuildComplexPoly(graphicsData *Graphics, webGLData interface{}) {
    self.Object.Call("buildComplexPoly", graphicsData, webGLData)
}

// BuildComplexPolyI Builds a complex polygon to draw
func (self *WebGLGraphics) BuildComplexPolyI(args ...interface{}) {
    self.Object.Call("buildComplexPoly", args)
}

// BuildPoly Builds a polygon to draw
func (self *WebGLGraphics) BuildPoly(graphicsData *Graphics, webGLData interface{}) {
    self.Object.Call("buildPoly", graphicsData, webGLData)
}

// BuildPolyI Builds a polygon to draw
func (self *WebGLGraphics) BuildPolyI(args ...interface{}) {
    self.Object.Call("buildPoly", args)
}

