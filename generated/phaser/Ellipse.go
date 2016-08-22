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
func (self *Ellipse) GetXA() int{
    return self.Object.Get("x").Int()
}

// The X coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) SetXA(member int) {
    self.Object.Set("x", member)
}

// The Y coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) GetYA() int{
    return self.Object.Get("y").Int()
}

// The Y coordinate of the upper-left corner of the framing rectangle of this ellipse.
func (self *Ellipse) SetYA(member int) {
    self.Object.Set("y", member)
}

// The overall width of this ellipse.
func (self *Ellipse) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The overall width of this ellipse.
func (self *Ellipse) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The overall height of this ellipse.
func (self *Ellipse) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The overall height of this ellipse.
func (self *Ellipse) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The const type of this object.
func (self *Ellipse) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *Ellipse) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The left coordinate of the Ellipse. The same as the X coordinate.
func (self *Ellipse) GetLeftA() interface{}{
    return self.Object.Get("left")
}

// The left coordinate of the Ellipse. The same as the X coordinate.
func (self *Ellipse) SetLeftA(member interface{}) {
    self.Object.Set("left", member)
}

// The x coordinate of the rightmost point of the Ellipse. Changing the right property of an Ellipse object has no effect on the x property, but does adjust the width. Gets or sets the value of the rightmost point of the ellipse.
func (self *Ellipse) GetRightA() int{
    return self.Object.Get("right").Int()
}

// The x coordinate of the rightmost point of the Ellipse. Changing the right property of an Ellipse object has no effect on the x property, but does adjust the width. Gets or sets the value of the rightmost point of the ellipse.
func (self *Ellipse) SetRightA(member int) {
    self.Object.Set("right", member)
}

// The top of the Ellipse. The same as its y property. Gets or sets the top of the ellipse.
func (self *Ellipse) GetTopA() int{
    return self.Object.Get("top").Int()
}

// The top of the Ellipse. The same as its y property. Gets or sets the top of the ellipse.
func (self *Ellipse) SetTopA(member int) {
    self.Object.Set("top", member)
}

// The sum of the y and height properties. Changing the bottom property of an Ellipse doesn't adjust the y property, but does change the height. Gets or sets the bottom of the ellipse.
func (self *Ellipse) GetBottomA() int{
    return self.Object.Get("bottom").Int()
}

// The sum of the y and height properties. Changing the bottom property of an Ellipse doesn't adjust the y property, but does change the height. Gets or sets the bottom of the ellipse.
func (self *Ellipse) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// Determines whether or not this Ellipse object is empty. Will return a value of true if the Ellipse objects dimensions are less than or equal to 0; otherwise false.
// If set to true it will reset all of the Ellipse objects properties to 0. An Ellipse object is empty if its width or height is less than or equal to 0. Gets or sets the empty state of the ellipse.
func (self *Ellipse) GetEmptyA() bool{
    return self.Object.Get("empty").Bool()
}

// Determines whether or not this Ellipse object is empty. Will return a value of true if the Ellipse objects dimensions are less than or equal to 0; otherwise false.
// If set to true it will reset all of the Ellipse objects properties to 0. An Ellipse object is empty if its width or height is less than or equal to 0. Gets or sets the empty state of the ellipse.
func (self *Ellipse) SetEmptyA(member bool) {
    self.Object.Set("empty", member)
}



// Sets the members of the Ellipse to the specified values.
func (self *Ellipse) SetTo(x int, y int, width int, height int) *Ellipse{
    return &Ellipse{self.Object.Call("setTo", x, y, width, height)}
}

// Sets the members of the Ellipse to the specified values.
func (self *Ellipse) SetToI(args ...interface{}) *Ellipse{
    return &Ellipse{self.Object.Call("setTo", args)}
}

// Returns the framing rectangle of the ellipse as a Phaser.Rectangle object.
func (self *Ellipse) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// Returns the framing rectangle of the ellipse as a Phaser.Rectangle object.
func (self *Ellipse) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Copies the x, y, width and height properties from any given object to this Ellipse.
func (self *Ellipse) CopyFrom(source interface{}) *Ellipse{
    return &Ellipse{self.Object.Call("copyFrom", source)}
}

// Copies the x, y, width and height properties from any given object to this Ellipse.
func (self *Ellipse) CopyFromI(args ...interface{}) *Ellipse{
    return &Ellipse{self.Object.Call("copyFrom", args)}
}

// Copies the x, y, width and height properties from this Ellipse to any given object.
func (self *Ellipse) CopyTo(dest interface{}) interface{}{
    return self.Object.Call("copyTo", dest)
}

// Copies the x, y, width and height properties from this Ellipse to any given object.
func (self *Ellipse) CopyToI(args ...interface{}) interface{}{
    return self.Object.Call("copyTo", args)
}

// Returns a new Ellipse object with the same values for the x, y, width, and height properties as this Ellipse object.
func (self *Ellipse) Clone(output *Ellipse) *Ellipse{
    return &Ellipse{self.Object.Call("clone", output)}
}

// Returns a new Ellipse object with the same values for the x, y, width, and height properties as this Ellipse object.
func (self *Ellipse) CloneI(args ...interface{}) *Ellipse{
    return &Ellipse{self.Object.Call("clone", args)}
}

// Return true if the given x/y coordinates are within this Ellipse object.
func (self *Ellipse) Contains(x int, y int) bool{
    return self.Object.Call("contains", x, y).Bool()
}

// Return true if the given x/y coordinates are within this Ellipse object.
func (self *Ellipse) ContainsI(args ...interface{}) bool{
    return self.Object.Call("contains", args).Bool()
}

// Returns a uniformly distributed random point from anywhere within this Ellipse.
func (self *Ellipse) Random() *Point{
    return &Point{self.Object.Call("random")}
}

// Returns a uniformly distributed random point from anywhere within this Ellipse.
func (self *Ellipse) Random1O(out interface{}) *Point{
    return &Point{self.Object.Call("random", out)}
}

// Returns a uniformly distributed random point from anywhere within this Ellipse.
func (self *Ellipse) RandomI(args ...interface{}) *Point{
    return &Point{self.Object.Call("random", args)}
}

// Returns a string representation of this object.
func (self *Ellipse) ToString() string{
    return self.Object.Call("toString").String()
}

// Returns a string representation of this object.
func (self *Ellipse) ToStringI(args ...interface{}) string{
    return self.Object.Call("toString", args).String()
}
