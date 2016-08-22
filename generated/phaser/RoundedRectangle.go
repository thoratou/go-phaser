// Automatic generation for Phaser.RoundedRectangle
// generated file RoundedRectangle.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
type RoundedRectangle struct {
    *js.Object
}


// The x coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) GetX() int{
    return self.Get("x").Int()
}

// The x coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) SetX(member int) {
    self.Set("x", member)
}

// The y coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) GetY() int{
    return self.Get("y").Int()
}

// The y coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) SetY(member int) {
    self.Set("y", member)
}

// The width of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) GetWidth() int{
    return self.Get("width").Int()
}

// The width of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) SetWidth(member int) {
    self.Set("width", member)
}

// The height of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) GetHeight() int{
    return self.Get("height").Int()
}

// The height of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) SetHeight(member int) {
    self.Set("height", member)
}

// The radius of the rounded corners.
func (self *RoundedRectangle) GetRadius() int{
    return self.Get("radius").Int()
}

// The radius of the rounded corners.
func (self *RoundedRectangle) SetRadius(member int) {
    self.Set("radius", member)
}

// The const type of this object.
func (self *RoundedRectangle) GetType() int{
    return self.Get("type").Int()
}

// The const type of this object.
func (self *RoundedRectangle) SetType(member int) {
    self.Set("type", member)
}



// Returns a new RoundedRectangle object with the same values for the x, y, width, height and
// radius properties as this RoundedRectangle object.
func (self *RoundedRectangle) CloneI(args ...interface{}) *RoundedRectangle{
    return &RoundedRectangle{self.Call("clone", args)}
}

// Determines whether the specified coordinates are contained within the region defined by this Rounded Rectangle object.
func (self *RoundedRectangle) ContainsI(args ...interface{}) bool{
    return self.Call("contains", args).Bool()
}
