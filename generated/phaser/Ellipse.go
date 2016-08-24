// Package phaser Automatic generation for Phaser.Ellipse
// generated file Ellipse.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Ellipse Creates a Ellipse object. A curve on a plane surrounding two focal points.
type Ellipse struct {
    *js.Object
}

// NewEllipse Creates a Ellipse object. A curve on a plane surrounding two focal points.
func NewEllipse() *Ellipse {
    return &Ellipse{js.Global.Get("Phaser").Get("Ellipse").New()}
}
// NewEllipse1O Creates a Ellipse object. A curve on a plane surrounding two focal points.
func NewEllipse1O(x int) *Ellipse {
    return &Ellipse{js.Global.Get("Phaser").Get("Ellipse").New(x)}
}
// NewEllipse2O Creates a Ellipse object. A curve on a plane surrounding two focal points.
func NewEllipse2O(x int, y int) *Ellipse {
    return &Ellipse{js.Global.Get("Phaser").Get("Ellipse").New(x, y)}
}
// NewEllipse3O Creates a Ellipse object. A curve on a plane surrounding two focal points.
func NewEllipse3O(x int, y int, width int) *Ellipse {
    return &Ellipse{js.Global.Get("Phaser").Get("Ellipse").New(x, y, width)}
}
// NewEllipse4O Creates a Ellipse object. A curve on a plane surrounding two focal points.
func NewEllipse4O(x int, y int, width int, height int) *Ellipse {
    return &Ellipse{js.Global.Get("Phaser").Get("Ellipse").New(x, y, width, height)}
}
// NewEllipseI Creates a Ellipse object. A curve on a plane surrounding two focal points.
func NewEllipseI(args ...interface{}) *Ellipse {
    return &Ellipse{js.Global.Get("Phaser").Get("Ellipse").New(args)}
}



// X The X coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The X coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The Y coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The Y coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) SetYA(member int) {
    self.Object.Set("y", member)
}

// Width The overall width of this ellipse.
func (self *Ellipse) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The overall width of this ellipse.
func (self *Ellipse) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The overall height of this ellipse.
func (self *Ellipse) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The overall height of this ellipse.
func (self *Ellipse) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Type The const type of this object.
func (self *Ellipse) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *Ellipse) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Left The left coordinate of the Ellipse. The same as the X coordinate.
func (self *Ellipse) Left() interface{}{
    return self.Object.Get("left")
}

// SetLeftA The left coordinate of the Ellipse. The same as the X coordinate.
func (self *Ellipse) SetLeftA(member interface{}) {
    self.Object.Set("left", member)
}

// Right The x coordinate of the rightmost point of the Ellipse. Changing the right property of an Ellipse object has no effect on the x property, but does adjust the width. Gets or sets the value of the rightmost point of the ellipse.
func (self *Ellipse) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA The x coordinate of the rightmost point of the Ellipse. Changing the right property of an Ellipse object has no effect on the x property, but does adjust the width. Gets or sets the value of the rightmost point of the ellipse.
func (self *Ellipse) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Top The top of the Ellipse. The same as its y property. Gets or sets the top of the ellipse.
func (self *Ellipse) Top() int{
    return self.Object.Get("top").Int()
}

// SetTopA The top of the Ellipse. The same as its y property. Gets or sets the top of the ellipse.
func (self *Ellipse) SetTopA(member int) {
    self.Object.Set("top", member)
}

// Bottom The sum of the y and height properties. Changing the bottom property of an Ellipse doesn't adjust the y property, but does change the height. Gets or sets the bottom of the ellipse.
func (self *Ellipse) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA The sum of the y and height properties. Changing the bottom property of an Ellipse doesn't adjust the y property, but does change the height. Gets or sets the bottom of the ellipse.
func (self *Ellipse) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// Empty Determines whether or not this Ellipse object is empty. Will return a value of true if the Ellipse objects dimensions are less than or equal to 0; otherwise false.
// If set to true it will reset all of the Ellipse objects properties to 0. An Ellipse object is empty if its width or height is less than or equal to 0. Gets or sets the empty state of the ellipse.
func (self *Ellipse) Empty() bool{
    return self.Object.Get("empty").Bool()
}

// SetEmptyA Determines whether or not this Ellipse object is empty. Will return a value of true if the Ellipse objects dimensions are less than or equal to 0; otherwise false.
// If set to true it will reset all of the Ellipse objects properties to 0. An Ellipse object is empty if its width or height is less than or equal to 0. Gets or sets the empty state of the ellipse.
func (self *Ellipse) SetEmptyA(member bool) {
    self.Object.Set("empty", member)
}


// SetTo Sets the members of the Ellipse to the specified values.
func (self *Ellipse) SetTo(x int, y int, width int, height int) *Ellipse{
    return &Ellipse{self.Object.Call("setTo", x, y, width, height)}
}

// SetToI Sets the members of the Ellipse to the specified values.
func (self *Ellipse) SetToI(args ...interface{}) *Ellipse{
    return &Ellipse{self.Object.Call("setTo", args)}
}

// GetBounds Returns the framing rectangle of the ellipse as a Phaser.Rectangle object.
func (self *Ellipse) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// GetBoundsI Returns the framing rectangle of the ellipse as a Phaser.Rectangle object.
func (self *Ellipse) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// CopyFrom Copies the x, y, width and height properties from any given object to this Ellipse.
func (self *Ellipse) CopyFrom(source interface{}) *Ellipse{
    return &Ellipse{self.Object.Call("copyFrom", source)}
}

// CopyFromI Copies the x, y, width and height properties from any given object to this Ellipse.
func (self *Ellipse) CopyFromI(args ...interface{}) *Ellipse{
    return &Ellipse{self.Object.Call("copyFrom", args)}
}

// CopyTo Copies the x, y, width and height properties from this Ellipse to any given object.
func (self *Ellipse) CopyTo(dest interface{}) interface{}{
    return self.Object.Call("copyTo", dest)
}

// CopyToI Copies the x, y, width and height properties from this Ellipse to any given object.
func (self *Ellipse) CopyToI(args ...interface{}) interface{}{
    return self.Object.Call("copyTo", args)
}

// Clone Returns a new Ellipse object with the same values for the x, y, width, and height properties as this Ellipse object.
func (self *Ellipse) Clone(output *Ellipse) *Ellipse{
    return &Ellipse{self.Object.Call("clone", output)}
}

// CloneI Returns a new Ellipse object with the same values for the x, y, width, and height properties as this Ellipse object.
func (self *Ellipse) CloneI(args ...interface{}) *Ellipse{
    return &Ellipse{self.Object.Call("clone", args)}
}

// Contains Return true if the given x/y coordinates are within this Ellipse object.
func (self *Ellipse) Contains(x int, y int) bool{
    return self.Object.Call("contains", x, y).Bool()
}

// ContainsI Return true if the given x/y coordinates are within this Ellipse object.
func (self *Ellipse) ContainsI(args ...interface{}) bool{
    return self.Object.Call("contains", args).Bool()
}

// Random Returns a uniformly distributed random point from anywhere within this Ellipse.
func (self *Ellipse) Random() *Point{
    return &Point{self.Object.Call("random")}
}

// Random1O Returns a uniformly distributed random point from anywhere within this Ellipse.
func (self *Ellipse) Random1O(out interface{}) *Point{
    return &Point{self.Object.Call("random", out)}
}

// RandomI Returns a uniformly distributed random point from anywhere within this Ellipse.
func (self *Ellipse) RandomI(args ...interface{}) *Point{
    return &Point{self.Object.Call("random", args)}
}

// ToString Returns a string representation of this object.
func (self *Ellipse) ToString() string{
    return self.Object.Call("toString").String()
}

// ToStringI Returns a string representation of this object.
func (self *Ellipse) ToStringI(args ...interface{}) string{
    return self.Object.Call("toString", args).String()
}

