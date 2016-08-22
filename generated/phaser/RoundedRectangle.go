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


// The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle() *RoundedRectangle {
    return &RoundedRectangle{js.Global.Call("Phaser.RoundedRectangle")}
}

// The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle1O(x int) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Call("Phaser.RoundedRectangle", x)}
}

// The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle2O(x int, y int) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Call("Phaser.RoundedRectangle", x, y)}
}

// The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle3O(x int, y int, width int) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Call("Phaser.RoundedRectangle", x, y, width)}
}

// The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle4O(x int, y int, width int, height int) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Call("Phaser.RoundedRectangle", x, y, width, height)}
}

// The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangle5O(x int, y int, width int, height int, radius int) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Call("Phaser.RoundedRectangle", x, y, width, height, radius)}
}

// The Rounded Rectangle object is an area defined by its position and has nice rounded corners, 
// as indicated by its top-left corner point (x, y) and by its width and its height.
func NewRoundedRectangleI(args ...interface{}) *RoundedRectangle {
    return &RoundedRectangle{js.Global.Call("Phaser.RoundedRectangle", args)}
}



// The x coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) GetXA() int{
    return self.Object.Get("x").Int()
}

// The x coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) SetXA(member int) {
    self.Object.Set("x", member)
}

// The y coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) GetYA() int{
    return self.Object.Get("y").Int()
}

// The y coordinate of the top-left corner of the Rectangle.
func (self *RoundedRectangle) SetYA(member int) {
    self.Object.Set("y", member)
}

// The width of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the Rectangle. This value should never be set to a negative.
func (self *RoundedRectangle) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The radius of the rounded corners.
func (self *RoundedRectangle) GetRadiusA() int{
    return self.Object.Get("radius").Int()
}

// The radius of the rounded corners.
func (self *RoundedRectangle) SetRadiusA(member int) {
    self.Object.Set("radius", member)
}

// The const type of this object.
func (self *RoundedRectangle) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *RoundedRectangle) SetTypeA(member int) {
    self.Object.Set("type", member)
}



// Returns a new RoundedRectangle object with the same values for the x, y, width, height and
// radius properties as this RoundedRectangle object.
func (self *RoundedRectangle) Clone() *RoundedRectangle{
    return &RoundedRectangle{self.Object.Call("clone")}
}

// Returns a new RoundedRectangle object with the same values for the x, y, width, height and
// radius properties as this RoundedRectangle object.
func (self *RoundedRectangle) CloneI(args ...interface{}) *RoundedRectangle{
    return &RoundedRectangle{self.Object.Call("clone", args)}
}

// Determines whether the specified coordinates are contained within the region defined by this Rounded Rectangle object.
func (self *RoundedRectangle) Contains(x int, y int) bool{
    return self.Object.Call("contains", x, y).Bool()
}

// Determines whether the specified coordinates are contained within the region defined by this Rounded Rectangle object.
func (self *RoundedRectangle) ContainsI(args ...interface{}) bool{
    return self.Object.Call("contains", args).Bool()
}
