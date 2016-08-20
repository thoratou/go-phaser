// Automatic generation for PIXI.PolyK
// generated file PolyK.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Based on the Polyk library http://polyk.ivank.net released under MIT licence.
// This is an amazing lib!
// Slightly modified by Mat Groves (matgroves.com);
type PolyK struct {
    *js.Object
}




// Triangulates shapes for webGL graphic fills.
func (self *PolyK) TriangulateI(args ...interface{}) {
    self.Call("Triangulate", args)
}

// Checks whether a point is within a triangle
func (self *PolyK) _PointInTriangleI(args ...interface{}) bool{
    return self.Call("_PointInTriangle", args).Bool()
}

// Checks whether a shape is convex
func (self *PolyK) _convexI(args ...interface{}) bool{
    return self.Call("_convex", args).Bool()
}
