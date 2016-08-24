// Package phaser Automatic generation for Phaser.RoundedRectangle
// generated file RoundedRectangle.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// RoundedRectangle The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
type RoundedRectangle struct {
    *js.Object
}

// NewRoundedRectangle The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle() *RoundedRectangle {
    return &RoundedRectangle{js.Global.Get("Phaser").Get("RoundedRectangle").New()}
}
// NewRoundedRectangle1O The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle1O(x int) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Get("Phaser").Get("RoundedRectangle").New(x)}
}
// NewRoundedRectangle2O The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle2O(x int, y int) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Get("Phaser").Get("RoundedRectangle").New(x, y)}
}
// NewRoundedRectangle3O The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle3O(x int, y int, width int) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Get("Phaser").Get("RoundedRectangle").New(x, y, width)}
}
// NewRoundedRectangle4O The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle4O(x int, y int, width int, height int) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Get("Phaser").Get("RoundedRectangle").New(x, y, width, height)}
}
// NewRoundedRectangle5O The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle5O(x int, y int, width int, height int, radius int) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Get("Phaser").Get("RoundedRectangle").New(x, y, width, height, radius)}
}
// NewRoundedRectangleI The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangleI(args ...interface{}) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Get("Phaser").Get("RoundedRectangle").New(args)}
}



// X The x coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The x coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The y coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The y coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) SetYA(member int) {
    self.Object.Set("y", member)
}

// Width The width of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Radius The radius of the rounded corners.
func (self *RoundedRectangle) Radius() int{
    return self.Object.Get("radius").Int()
}

// SetRadiusA The radius of the rounded corners.
func (self *RoundedRectangle) SetRadiusA(member int) {
    self.Object.Set("radius", member)
}

// Type The const type of this object.
func (self *RoundedRectangle) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *RoundedRectangle) SetTypeA(member int) {
    self.Object.Set("type", member)
}


// Clone Returns a new RoundedRectangle object with the same values for the x, y, width, height and
// radius properties as this RoundedRectangle object.
func (self *RoundedRectangle) Clone() *RoundedRectangle{
    return &RoundedRectangle{self.Object.Call("clone")}
}

// CloneI Returns a new RoundedRectangle object with the same values for the x, y, width, height and
// radius properties as this RoundedRectangle object.
func (self *RoundedRectangle) CloneI(args ...interface{}) *RoundedRectangle{
    return &RoundedRectangle{self.Object.Call("clone", args)}
}

// Contains Determines whether the specified coordinates are contained within the region defined by this Rounded Rectangle object.
func (self *RoundedRectangle) Contains(x int, y int) bool{
    return self.Object.Call("contains", x, y).Bool()
}

// ContainsI Determines whether the specified coordinates are contained within the region defined by this Rounded Rectangle object.
func (self *RoundedRectangle) ContainsI(args ...interface{}) bool{
    return self.Object.Call("contains", args).Bool()
}

