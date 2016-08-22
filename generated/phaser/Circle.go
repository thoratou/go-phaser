// Automatic generation for Phaser.Circle
// generated file Circle.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Creates a new Circle object with the center coordinate specified by the x and y parameters and the diameter specified by the diameter parameter.
// If you call this function without parameters, a circle with x, y, diameter and radius properties set to 0 is created.
type Circle struct {
    *js.Object
}


// The x coordinate of the center of the circle.
func (self *Circle) GetX() int{
    return self.Get("x").Int()
}

// The x coordinate of the center of the circle.
func (self *Circle) SetX(member int) {
    self.Set("x", member)
}

// The y coordinate of the center of the circle.
func (self *Circle) GetY() int{
    return self.Get("y").Int()
}

// The y coordinate of the center of the circle.
func (self *Circle) SetY(member int) {
    self.Set("y", member)
}

// The const type of this object.
func (self *Circle) GetType() int{
    return self.Get("type").Int()
}

// The const type of this object.
func (self *Circle) SetType(member int) {
    self.Set("type", member)
}

// The largest distance between any two points on the circle. The same as the radius * 2. Gets or sets the diameter of the circle.
func (self *Circle) GetDiameter() int{
    return self.Get("diameter").Int()
}

// The largest distance between any two points on the circle. The same as the radius * 2. Gets or sets the diameter of the circle.
func (self *Circle) SetDiameter(member int) {
    self.Set("diameter", member)
}

// The length of a line extending from the center of the circle to any point on the circle itself. The same as half the diameter. Gets or sets the radius of the circle.
func (self *Circle) GetRadius() int{
    return self.Get("radius").Int()
}

// The length of a line extending from the center of the circle to any point on the circle itself. The same as half the diameter. Gets or sets the radius of the circle.
func (self *Circle) SetRadius(member int) {
    self.Set("radius", member)
}

// The x coordinate of the leftmost point of the circle. Changing the left property of a Circle object has no effect on the x and y properties. However it does affect the diameter, whereas changing the x value does not affect the diameter property.
func (self *Circle) GetLeft() interface{}{
    return self.Get("left")
}

// The x coordinate of the leftmost point of the circle. Changing the left property of a Circle object has no effect on the x and y properties. However it does affect the diameter, whereas changing the x value does not affect the diameter property.
func (self *Circle) SetLeft(member interface{}) {
    self.Set("left", member)
}

// The x coordinate of the rightmost point of the circle. Changing the right property of a Circle object has no effect on the x and y properties. However it does affect the diameter, whereas changing the x value does not affect the diameter property. Gets or sets the value of the rightmost point of the circle.
func (self *Circle) GetRight() int{
    return self.Get("right").Int()
}

// The x coordinate of the rightmost point of the circle. Changing the right property of a Circle object has no effect on the x and y properties. However it does affect the diameter, whereas changing the x value does not affect the diameter property. Gets or sets the value of the rightmost point of the circle.
func (self *Circle) SetRight(member int) {
    self.Set("right", member)
}

// The sum of the y minus the radius property. Changing the top property of a Circle object has no effect on the x and y properties, but does change the diameter. Gets or sets the top of the circle.
func (self *Circle) GetTop() int{
    return self.Get("top").Int()
}

// The sum of the y minus the radius property. Changing the top property of a Circle object has no effect on the x and y properties, but does change the diameter. Gets or sets the top of the circle.
func (self *Circle) SetTop(member int) {
    self.Set("top", member)
}

// The sum of the y and radius properties. Changing the bottom property of a Circle object has no effect on the x and y properties, but does change the diameter. Gets or sets the bottom of the circle.
func (self *Circle) GetBottom() int{
    return self.Get("bottom").Int()
}

// The sum of the y and radius properties. Changing the bottom property of a Circle object has no effect on the x and y properties, but does change the diameter. Gets or sets the bottom of the circle.
func (self *Circle) SetBottom(member int) {
    self.Set("bottom", member)
}

// The area of this Circle.
func (self *Circle) GetArea() int{
    return self.Get("area").Int()
}

// The area of this Circle.
func (self *Circle) SetArea(member int) {
    self.Set("area", member)
}

// Determines whether or not this Circle object is empty. Will return a value of true if the Circle objects diameter is less than or equal to 0; otherwise false.
// If set to true it will reset all of the Circle objects properties to 0. A Circle object is empty if its diameter is less than or equal to 0. Gets or sets the empty state of the circle.
func (self *Circle) GetEmpty() bool{
    return self.Get("empty").Bool()
}

// Determines whether or not this Circle object is empty. Will return a value of true if the Circle objects diameter is less than or equal to 0; otherwise false.
// If set to true it will reset all of the Circle objects properties to 0. A Circle object is empty if its diameter is less than or equal to 0. Gets or sets the empty state of the circle.
func (self *Circle) SetEmpty(member bool) {
    self.Set("empty", member)
}



// The circumference of the circle.
func (self *Circle) CircumferenceI(args ...interface{}) int{
    return self.Call("circumference", args).Int()
}

// Returns a uniformly distributed random point from anywhere within this Circle.
func (self *Circle) RandomI(args ...interface{}) *Point{
    return &Point{self.Call("random", args)}
}

// Returns the framing rectangle of the circle as a Phaser.Rectangle object.
func (self *Circle) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getBounds", args)}
}

// Sets the members of Circle to the specified values.
func (self *Circle) SetToI(args ...interface{}) *Circle{
    return &Circle{self.Call("setTo", args)}
}

// Copies the x, y and diameter properties from any given object to this Circle.
func (self *Circle) CopyFromI(args ...interface{}) *Circle{
    return &Circle{self.Call("copyFrom", args)}
}

// Copies the x, y and diameter properties from this Circle to any given object.
func (self *Circle) CopyToI(args ...interface{}) interface{}{
    return self.Call("copyTo", args)
}

// Returns the distance from the center of the Circle object to the given object
// (can be Circle, Point or anything with x/y properties)
func (self *Circle) DistanceI(args ...interface{}) int{
    return self.Call("distance", args).Int()
}

// Returns a new Circle object with the same values for the x, y, width, and height properties as this Circle object.
func (self *Circle) CloneI(args ...interface{}) *Circle{
    return &Circle{self.Call("clone", args)}
}

// Return true if the given x/y coordinates are within this Circle object.
func (self *Circle) ContainsI(args ...interface{}) bool{
    return self.Call("contains", args).Bool()
}

// Returns a Point object containing the coordinates of a point on the circumference of the Circle based on the given angle.
func (self *Circle) CircumferencePointI(args ...interface{}) *Point{
    return &Point{self.Call("circumferencePoint", args)}
}

// Adjusts the location of the Circle object, as determined by its center coordinate, by the specified amounts.
func (self *Circle) OffsetI(args ...interface{}) *Circle{
    return &Circle{self.Call("offset", args)}
}

// Adjusts the location of the Circle object using a Point object as a parameter. This method is similar to the Circle.offset() method, except that it takes a Point object as a parameter.
func (self *Circle) OffsetPointI(args ...interface{}) *Circle{
    return &Circle{self.Call("offsetPoint", args)}
}

// Returns a string representation of this object.
func (self *Circle) ToStringI(args ...interface{}) string{
    return self.Call("toString", args).String()
}

// Determines whether the two Circle objects match. This method compares the x, y and diameter properties.
func (self *Circle) EqualsI(args ...interface{}) bool{
    return self.Call("equals", args).Bool()
}

// Determines whether the two Circle objects intersect.
// This method checks the radius distances between the two Circle objects to see if they intersect.
func (self *Circle) IntersectsI(args ...interface{}) bool{
    return self.Call("intersects", args).Bool()
}

// Checks if the given Circle and Rectangle objects intersect.
func (self *Circle) IntersectsRectangleI(args ...interface{}) bool{
    return self.Call("intersectsRectangle", args).Bool()
}
