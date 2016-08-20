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
func (self *RoundedRectangle) GetX() float64{
    return self.Get("x").Float()
}

// The x coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) SetX(member float64) {
    self.Set("x", member)
}

// The y coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) GetY() float64{
    return self.Get("y").Float()
}

// The y coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) SetY(member float64) {
    self.Set("y", member)
}

// The width of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) SetHeight(member float64) {
    self.Set("height", member)
}

// The radius of the rounded corners.
func (self *RoundedRectangle) GetRadius() float64{
    return self.Get("radius").Float()
}

// The radius of the rounded corners.
func (self *RoundedRectangle) SetRadius(member float64) {
    self.Set("radius", member)
}

// The const type of this object.
func (self *RoundedRectangle) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *RoundedRectangle) SetType(member float64) {
    self.Set("type", member)
}



// Returns a new RoundedRectangle object with the same values for the x, y, width, height and
// radius properties as this RoundedRectangle object.
func (self *RoundedRectangle) CloneI(args ...interface{}) RoundedRectangle{
    return RoundedRectangle{self.Call("clone", args)}
}

// Determines whether the specified coordinates are contained within the region defined by this Rounded Rectangle object.
func (self *RoundedRectangle) ContainsI(args ...interface{}) bool{
    return self.Call("contains", args).Bool()
}
