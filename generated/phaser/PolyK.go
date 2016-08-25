// Package phaser Automatic generation for PIXI.PolyK
// generated file PolyK.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PolyK Based on the Polyk library http://polyk.ivank.net released under MIT licence.
// This is an amazing lib!
// Slightly modified by Mat Groves (matgroves.com);
type PolyK struct {
    *js.Object
}

// NewPolyK Based on the Polyk library http://polyk.ivank.net released under MIT licence.
// This is an amazing lib!
// Slightly modified by Mat Groves (matgroves.com);
func NewPolyK() *PolyK {
    return &PolyK{js.Global.Get("PIXI").Get("PolyK").New()}
}
// NewPolyKI Based on the Polyk library http://polyk.ivank.net released under MIT licence.
// This is an amazing lib!
// Slightly modified by Mat Groves (matgroves.com);
func NewPolyKI(args ...interface{}) *PolyK {
    return &PolyK{js.Global.Get("PIXI").Get("PolyK").New(args)}
}



// PolyK Binding conversion method to PolyK point 
func ToPolyK(jsStruct interface{}) *PolyK {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PolyK{Object: object}
	}
	return nil
}




// Triangulate Triangulates shapes for webGL graphic fills.
func (self *PolyK) Triangulate() {
    self.Object.Call("Triangulate")
}

// TriangulateI Triangulates shapes for webGL graphic fills.
func (self *PolyK) TriangulateI(args ...interface{}) {
    self.Object.Call("Triangulate", args)
}

// _PointInTriangle Checks whether a point is within a triangle
func (self *PolyK) _PointInTriangle(px int, py int, ax int, ay int, bx int, by int, cx int, cy int) bool{
    return self.Object.Call("_PointInTriangle", px, py, ax, ay, bx, by, cx, cy).Bool()
}

// _PointInTriangleI Checks whether a point is within a triangle
func (self *PolyK) _PointInTriangleI(args ...interface{}) bool{
    return self.Object.Call("_PointInTriangle", args).Bool()
}

// _convex Checks whether a shape is convex
func (self *PolyK) _convex() bool{
    return self.Object.Call("_convex").Bool()
}

// _convexI Checks whether a shape is convex
func (self *PolyK) _convexI(args ...interface{}) bool{
    return self.Object.Call("_convex", args).Bool()
}

