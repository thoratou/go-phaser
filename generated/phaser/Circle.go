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
func (self *Circle) GetXA() int{
    return self.Object.Get("x").Int()
}

// The x coordinate of the center of the circle.
func (self *Circle) SetXA(member int) {
    self.Object.Set("x", member)
}

// The y coordinate of the center of the circle.
func (self *Circle) GetYA() int{
    return self.Object.Get("y").Int()
}

// The y coordinate of the center of the circle.
func (self *Circle) SetYA(member int) {
    self.Object.Set("y", member)
}

// The const type of this object.
func (self *Circle) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *Circle) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The largest distance between any two points on the circle. The same as the radius * 2. Gets or sets the diameter of the circle.
func (self *Circle) GetDiameterA() int{
    return self.Object.Get("diameter").Int()
}

// The largest distance between any two points on the circle. The same as the radius * 2. Gets or sets the diameter of the circle.
func (self *Circle) SetDiameterA(member int) {
    self.Object.Set("diameter", member)
}

// The length of a line extending from the center of the circle to any point on the circle itself. The same as half the diameter. Gets or sets the radius of the circle.
func (self *Circle) GetRadiusA() int{
    return self.Object.Get("radius").Int()
}

// The length of a line extending from the center of the circle to any point on the circle itself. The same as half the diameter. Gets or sets the radius of the circle.
func (self *Circle) SetRadiusA(member int) {
    self.Object.Set("radius", member)
}

// The x coordinate of the leftmost point of the circle. Changing the left property of a Circle object has no effect on the x and y properties. However it does affect the diameter, whereas changing the x value does not affect the diameter property.
func (self *Circle) GetLeftA() interface{}{
    return self.Object.Get("left")
}

// The x coordinate of the leftmost point of the circle. Changing the left property of a Circle object has no effect on the x and y properties. However it does affect the diameter, whereas changing the x value does not affect the diameter property.
func (self *Circle) SetLeftA(member interface{}) {
    self.Object.Set("left", member)
}

// The x coordinate of the rightmost point of the circle. Changing the right property of a Circle object has no effect on the x and y properties. However it does affect the diameter, whereas changing the x value does not affect the diameter property. Gets or sets the value of the rightmost point of the circle.
func (self *Circle) GetRightA() int{
    return self.Object.Get("right").Int()
}

// The x coordinate of the rightmost point of the circle. Changing the right property of a Circle object has no effect on the x and y properties. However it does affect the diameter, whereas changing the x value does not affect the diameter property. Gets or sets the value of the rightmost point of the circle.
func (self *Circle) SetRightA(member int) {
    self.Object.Set("right", member)
}

// The sum of the y minus the radius property. Changing the top property of a Circle object has no effect on the x and y properties, but does change the diameter. Gets or sets the top of the circle.
func (self *Circle) GetTopA() int{
    return self.Object.Get("top").Int()
}

// The sum of the y minus the radius property. Changing the top property of a Circle object has no effect on the x and y properties, but does change the diameter. Gets or sets the top of the circle.
func (self *Circle) SetTopA(member int) {
    self.Object.Set("top", member)
}

// The sum of the y and radius properties. Changing the bottom property of a Circle object has no effect on the x and y properties, but does change the diameter. Gets or sets the bottom of the circle.
func (self *Circle) GetBottomA() int{
    return self.Object.Get("bottom").Int()
}

// The sum of the y and radius properties. Changing the bottom property of a Circle object has no effect on the x and y properties, but does change the diameter. Gets or sets the bottom of the circle.
func (self *Circle) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// The area of this Circle.
func (self *Circle) GetAreaA() int{
    return self.Object.Get("area").Int()
}

// The area of this Circle.
func (self *Circle) SetAreaA(member int) {
    self.Object.Set("area", member)
}

// Determines whether or not this Circle object is empty. Will return a value of true if the Circle objects diameter is less than or equal to 0; otherwise false.
// If set to true it will reset all of the Circle objects properties to 0. A Circle object is empty if its diameter is less than or equal to 0. Gets or sets the empty state of the circle.
func (self *Circle) GetEmptyA() bool{
    return self.Object.Get("empty").Bool()
}

// Determines whether or not this Circle object is empty. Will return a value of true if the Circle objects diameter is less than or equal to 0; otherwise false.
// If set to true it will reset all of the Circle objects properties to 0. A Circle object is empty if its diameter is less than or equal to 0. Gets or sets the empty state of the circle.
func (self *Circle) SetEmptyA(member bool) {
    self.Object.Set("empty", member)
}



// The circumference of the circle.
func (self *Circle) Circumference() int{
    return self.Object.Call("circumference").Int()
}

// The circumference of the circle.
func (self *Circle) CircumferenceI(args ...interface{}) int{
    return self.Object.Call("circumference", args).Int()
}

// Returns a uniformly distributed random point from anywhere within this Circle.
func (self *Circle) Random() *Point{
    return &Point{self.Object.Call("random")}
}

// Returns a uniformly distributed random point from anywhere within this Circle.
func (self *Circle) Random1O(out interface{}) *Point{
    return &Point{self.Object.Call("random", out)}
}

// Returns a uniformly distributed random point from anywhere within this Circle.
func (self *Circle) RandomI(args ...interface{}) *Point{
    return &Point{self.Object.Call("random", args)}
}

// Returns the framing rectangle of the circle as a Phaser.Rectangle object.
func (self *Circle) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// Returns the framing rectangle of the circle as a Phaser.Rectangle object.
func (self *Circle) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Sets the members of Circle to the specified values.
func (self *Circle) SetTo(x int, y int, diameter int) *Circle{
    return &Circle{self.Object.Call("setTo", x, y, diameter)}
}

// Sets the members of Circle to the specified values.
func (self *Circle) SetToI(args ...interface{}) *Circle{
    return &Circle{self.Object.Call("setTo", args)}
}

// Copies the x, y and diameter properties from any given object to this Circle.
func (self *Circle) CopyFrom(source interface{}) *Circle{
    return &Circle{self.Object.Call("copyFrom", source)}
}

// Copies the x, y and diameter properties from any given object to this Circle.
func (self *Circle) CopyFromI(args ...interface{}) *Circle{
    return &Circle{self.Object.Call("copyFrom", args)}
}

// Copies the x, y and diameter properties from this Circle to any given object.
func (self *Circle) CopyTo(dest interface{}) interface{}{
    return self.Object.Call("copyTo", dest)
}

// Copies the x, y and diameter properties from this Circle to any given object.
func (self *Circle) CopyToI(args ...interface{}) interface{}{
    return self.Object.Call("copyTo", args)
}

// Returns the distance from the center of the Circle object to the given object
// (can be Circle, Point or anything with x/y properties)
func (self *Circle) Distance(dest interface{}) int{
    return self.Object.Call("distance", dest).Int()
}

// Returns the distance from the center of the Circle object to the given object
// (can be Circle, Point or anything with x/y properties)
func (self *Circle) Distance1O(dest interface{}, round bool) int{
    return self.Object.Call("distance", dest, round).Int()
}

// Returns the distance from the center of the Circle object to the given object
// (can be Circle, Point or anything with x/y properties)
func (self *Circle) DistanceI(args ...interface{}) int{
    return self.Object.Call("distance", args).Int()
}

// Returns a new Circle object with the same values for the x, y, width, and height properties as this Circle object.
func (self *Circle) Clone(output *Circle) *Circle{
    return &Circle{self.Object.Call("clone", output)}
}

// Returns a new Circle object with the same values for the x, y, width, and height properties as this Circle object.
func (self *Circle) CloneI(args ...interface{}) *Circle{
    return &Circle{self.Object.Call("clone", args)}
}

// Return true if the given x/y coordinates are within this Circle object.
func (self *Circle) Contains(x int, y int) bool{
    return self.Object.Call("contains", x, y).Bool()
}

// Return true if the given x/y coordinates are within this Circle object.
func (self *Circle) ContainsI(args ...interface{}) bool{
    return self.Object.Call("contains", args).Bool()
}

// Returns a Point object containing the coordinates of a point on the circumference of the Circle based on the given angle.
func (self *Circle) CircumferencePoint(angle int) *Point{
    return &Point{self.Object.Call("circumferencePoint", angle)}
}

// Returns a Point object containing the coordinates of a point on the circumference of the Circle based on the given angle.
func (self *Circle) CircumferencePoint1O(angle int, asDegrees bool) *Point{
    return &Point{self.Object.Call("circumferencePoint", angle, asDegrees)}
}

// Returns a Point object containing the coordinates of a point on the circumference of the Circle based on the given angle.
func (self *Circle) CircumferencePoint2O(angle int, asDegrees bool, out *Point) *Point{
    return &Point{self.Object.Call("circumferencePoint", angle, asDegrees, out)}
}

// Returns a Point object containing the coordinates of a point on the circumference of the Circle based on the given angle.
func (self *Circle) CircumferencePointI(args ...interface{}) *Point{
    return &Point{self.Object.Call("circumferencePoint", args)}
}

// Adjusts the location of the Circle object, as determined by its center coordinate, by the specified amounts.
func (self *Circle) Offset(dx int, dy int) *Circle{
    return &Circle{self.Object.Call("offset", dx, dy)}
}

// Adjusts the location of the Circle object, as determined by its center coordinate, by the specified amounts.
func (self *Circle) OffsetI(args ...interface{}) *Circle{
    return &Circle{self.Object.Call("offset", args)}
}

// Adjusts the location of the Circle object using a Point object as a parameter. This method is similar to the Circle.offset() method, except that it takes a Point object as a parameter.
func (self *Circle) OffsetPoint(point *Point) *Circle{
    return &Circle{self.Object.Call("offsetPoint", point)}
}

// Adjusts the location of the Circle object using a Point object as a parameter. This method is similar to the Circle.offset() method, except that it takes a Point object as a parameter.
func (self *Circle) OffsetPointI(args ...interface{}) *Circle{
    return &Circle{self.Object.Call("offsetPoint", args)}
}

// Returns a string representation of this object.
func (self *Circle) ToString() string{
    return self.Object.Call("toString").String()
}

// Returns a string representation of this object.
func (self *Circle) ToStringI(args ...interface{}) string{
    return self.Object.Call("toString", args).String()
}

// Determines whether the two Circle objects match. This method compares the x, y and diameter properties.
func (self *Circle) Equals(a *Circle, b *Circle) bool{
    return self.Object.Call("equals", a, b).Bool()
}

// Determines whether the two Circle objects match. This method compares the x, y and diameter properties.
func (self *Circle) EqualsI(args ...interface{}) bool{
    return self.Object.Call("equals", args).Bool()
}

// Determines whether the two Circle objects intersect.
// This method checks the radius distances between the two Circle objects to see if they intersect.
func (self *Circle) Intersects(a *Circle, b *Circle) bool{
    return self.Object.Call("intersects", a, b).Bool()
}

// Determines whether the two Circle objects intersect.
// This method checks the radius distances between the two Circle objects to see if they intersect.
func (self *Circle) IntersectsI(args ...interface{}) bool{
    return self.Object.Call("intersects", args).Bool()
}

// Checks if the given Circle and Rectangle objects intersect.
func (self *Circle) IntersectsRectangle(c *Circle, r *Rectangle) bool{
    return self.Object.Call("intersectsRectangle", c, r).Bool()
}

// Checks if the given Circle and Rectangle objects intersect.
func (self *Circle) IntersectsRectangleI(args ...interface{}) bool{
    return self.Object.Call("intersectsRectangle", args).Bool()
}
