// Automatic generation for Phaser.Ellipse
// generated file Ellipse.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Creates a Ellipse object. A curve on a plane surrounding two focal points.
type Ellipse struct {
    *js.Object
}


// The X coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) GetX() float64{
    return self.Get("x").Float()
}

// The X coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) SetX(member float64) {
    self.Set("x", member)
}

// The Y coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) GetY() float64{
    return self.Get("y").Float()
}

// The Y coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) SetY(member float64) {
    self.Set("y", member)
}

// The overall width of this ellipse.
func (self *Ellipse) GetWidth() float64{
    return self.Get("width").Float()
}

// The overall width of this ellipse.
func (self *Ellipse) SetWidth(member float64) {
    self.Set("width", member)
}

// The overall height of this ellipse.
func (self *Ellipse) GetHeight() float64{
    return self.Get("height").Float()
}

// The overall height of this ellipse.
func (self *Ellipse) SetHeight(member float64) {
    self.Set("height", member)
}

// The const type of this object.
func (self *Ellipse) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *Ellipse) SetType(member float64) {
    self.Set("type", member)
}

// The left coordinate of the Ellipse. The same as the X coordinate.
func (self *Ellipse) GetLeft() interface{}{
    return self.Get("left")
}

// The left coordinate of the Ellipse. The same as the X coordinate.
func (self *Ellipse) SetLeft(member interface{}) {
    self.Set("left", member)
}

// The x coordinate of the rightmost point of the Ellipse. Changing the right property of an Ellipse object has no effect on the x property, but does adjust the width. Gets or sets the value of the rightmost point of the ellipse.
func (self *Ellipse) GetRight() float64{
    return self.Get("right").Float()
}

// The x coordinate of the rightmost point of the Ellipse. Changing the right property of an Ellipse object has no effect on the x property, but does adjust the width. Gets or sets the value of the rightmost point of the ellipse.
func (self *Ellipse) SetRight(member float64) {
    self.Set("right", member)
}

// The top of the Ellipse. The same as its y property. Gets or sets the top of the ellipse.
func (self *Ellipse) GetTop() float64{
    return self.Get("top").Float()
}

// The top of the Ellipse. The same as its y property. Gets or sets the top of the ellipse.
func (self *Ellipse) SetTop(member float64) {
    self.Set("top", member)
}

// The sum of the y and height properties. Changing the bottom property of an Ellipse doesn't adjust the y property, but does change the height. Gets or sets the bottom of the ellipse.
func (self *Ellipse) GetBottom() float64{
    return self.Get("bottom").Float()
}

// The sum of the y and height properties. Changing the bottom property of an Ellipse doesn't adjust the y property, but does change the height. Gets or sets the bottom of the ellipse.
func (self *Ellipse) SetBottom(member float64) {
    self.Set("bottom", member)
}

// Determines whether or not this Ellipse object is empty. Will return a value of true if the Ellipse objects dimensions are less than or equal to 0; otherwise false.
// If set to true it will reset all of the Ellipse objects properties to 0. An Ellipse object is empty if its width or height is less than or equal to 0. Gets or sets the empty state of the ellipse.
func (self *Ellipse) GetEmpty() bool{
    return self.Get("empty").Bool()
}

// Determines whether or not this Ellipse object is empty. Will return a value of true if the Ellipse objects dimensions are less than or equal to 0; otherwise false.
// If set to true it will reset all of the Ellipse objects properties to 0. An Ellipse object is empty if its width or height is less than or equal to 0. Gets or sets the empty state of the ellipse.
func (self *Ellipse) SetEmpty(member bool) {
    self.Set("empty", member)
}



// Sets the members of the Ellipse to the specified values.
func (self *Ellipse) SetToI(args ...interface{}) *Ellipse{
    return &Ellipse{self.Call("setTo", args)}
}

// Returns the framing rectangle of the ellipse as a Phaser.Rectangle object.
func (self *Ellipse) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getBounds", args)}
}

// Copies the x, y, width and height properties from any given object to this Ellipse.
func (self *Ellipse) CopyFromI(args ...interface{}) *Ellipse{
    return &Ellipse{self.Call("copyFrom", args)}
}

// Copies the x, y, width and height properties from this Ellipse to any given object.
func (self *Ellipse) CopyToI(args ...interface{}) interface{}{
    return self.Call("copyTo", args)
}

// Returns a new Ellipse object with the same values for the x, y, width, and height properties as this Ellipse object.
func (self *Ellipse) CloneI(args ...interface{}) *Ellipse{
    return &Ellipse{self.Call("clone", args)}
}

// Return true if the given x/y coordinates are within this Ellipse object.
func (self *Ellipse) ContainsI(args ...interface{}) bool{
    return self.Call("contains", args).Bool()
}

// Returns a uniformly distributed random point from anywhere within this Ellipse.
func (self *Ellipse) RandomI(args ...interface{}) *Point{
    return &Point{self.Call("random", args)}
}

// Returns a string representation of this object.
func (self *Ellipse) ToStringI(args ...interface{}) string{
    return self.Call("toString", args).String()
}
