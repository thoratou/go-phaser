// Automatic generation for Phaser.Rectangle
// generated file Rectangle.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Creates a new Rectangle object with the top-left corner specified by the x and y parameters and with the specified width and height parameters.
// If you call this function without parameters, a Rectangle with x, y, width, and height properties set to 0 is created.
type Rectangle struct {
    *js.Object
}


// The x coordinate of the top-left corner of the Rectangle.
func (self *Rectangle) GetX() float64{
    return self.Get("x").Float()
}

// The x coordinate of the top-left corner of the Rectangle.
func (self *Rectangle) SetX(member float64) {
    self.Set("x", member)
}

// The y coordinate of the top-left corner of the Rectangle.
func (self *Rectangle) GetY() float64{
    return self.Get("y").Float()
}

// The y coordinate of the top-left corner of the Rectangle.
func (self *Rectangle) SetY(member float64) {
    self.Set("y", member)
}

// The width of the Rectangle. This value should never be set to a negative.
func (self *Rectangle) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the Rectangle. This value should never be set to a negative.
func (self *Rectangle) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the Rectangle. This value should never be set to a negative.
func (self *Rectangle) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the Rectangle. This value should never be set to a negative.
func (self *Rectangle) SetHeight(member float64) {
    self.Set("height", member)
}

// The const type of this object.
func (self *Rectangle) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *Rectangle) SetType(member float64) {
    self.Set("type", member)
}

// Half of the width of the Rectangle.
func (self *Rectangle) GetHalfWidth() float64{
    return self.Get("halfWidth").Float()
}

// Half of the width of the Rectangle.
func (self *Rectangle) SetHalfWidth(member float64) {
    self.Set("halfWidth", member)
}

// Half of the height of the Rectangle.
func (self *Rectangle) GetHalfHeight() float64{
    return self.Get("halfHeight").Float()
}

// Half of the height of the Rectangle.
func (self *Rectangle) SetHalfHeight(member float64) {
    self.Set("halfHeight", member)
}

// The sum of the y and height properties. Changing the bottom property of a Rectangle object has no effect on the x, y and width properties, but does change the height property.
func (self *Rectangle) GetBottom() float64{
    return self.Get("bottom").Float()
}

// The sum of the y and height properties. Changing the bottom property of a Rectangle object has no effect on the x, y and width properties, but does change the height property.
func (self *Rectangle) SetBottom(member float64) {
    self.Set("bottom", member)
}

// The location of the Rectangles bottom left corner as a Point object. Gets or sets the location of the Rectangles bottom left corner as a Point object.
func (self *Rectangle) GetBottomLeft() *Point{
    return &Point{self.Get("bottomLeft")}
}

// The location of the Rectangles bottom left corner as a Point object. Gets or sets the location of the Rectangles bottom left corner as a Point object.
func (self *Rectangle) SetBottomLeft(member *Point) {
    self.Set("bottomLeft", member)
}

// The location of the Rectangles bottom right corner as a Point object. Gets or sets the location of the Rectangles bottom right corner as a Point object.
func (self *Rectangle) GetBottomRight() *Point{
    return &Point{self.Get("bottomRight")}
}

// The location of the Rectangles bottom right corner as a Point object. Gets or sets the location of the Rectangles bottom right corner as a Point object.
func (self *Rectangle) SetBottomRight(member *Point) {
    self.Set("bottomRight", member)
}

// The x coordinate of the left of the Rectangle. Changing the left property of a Rectangle object has no effect on the y and height properties. However it does affect the width property, whereas changing the x value does not affect the width property.
func (self *Rectangle) GetLeft() float64{
    return self.Get("left").Float()
}

// The x coordinate of the left of the Rectangle. Changing the left property of a Rectangle object has no effect on the y and height properties. However it does affect the width property, whereas changing the x value does not affect the width property.
func (self *Rectangle) SetLeft(member float64) {
    self.Set("left", member)
}

// The sum of the x and width properties. Changing the right property of a Rectangle object has no effect on the x, y and height properties, however it does affect the width property.
func (self *Rectangle) GetRight() float64{
    return self.Get("right").Float()
}

// The sum of the x and width properties. Changing the right property of a Rectangle object has no effect on the x, y and height properties, however it does affect the width property.
func (self *Rectangle) SetRight(member float64) {
    self.Set("right", member)
}

// The volume of the Rectangle derived from width * height.
func (self *Rectangle) GetVolume() float64{
    return self.Get("volume").Float()
}

// The volume of the Rectangle derived from width * height.
func (self *Rectangle) SetVolume(member float64) {
    self.Set("volume", member)
}

// The perimeter size of the Rectangle. This is the sum of all 4 sides.
func (self *Rectangle) GetPerimeter() float64{
    return self.Get("perimeter").Float()
}

// The perimeter size of the Rectangle. This is the sum of all 4 sides.
func (self *Rectangle) SetPerimeter(member float64) {
    self.Set("perimeter", member)
}

// The x coordinate of the center of the Rectangle.
func (self *Rectangle) GetCenterX() float64{
    return self.Get("centerX").Float()
}

// The x coordinate of the center of the Rectangle.
func (self *Rectangle) SetCenterX(member float64) {
    self.Set("centerX", member)
}

// The y coordinate of the center of the Rectangle.
func (self *Rectangle) GetCenterY() float64{
    return self.Get("centerY").Float()
}

// The y coordinate of the center of the Rectangle.
func (self *Rectangle) SetCenterY(member float64) {
    self.Set("centerY", member)
}

// A random value between the left and right values (inclusive) of the Rectangle.
func (self *Rectangle) GetRandomX() float64{
    return self.Get("randomX").Float()
}

// A random value between the left and right values (inclusive) of the Rectangle.
func (self *Rectangle) SetRandomX(member float64) {
    self.Set("randomX", member)
}

// A random value between the top and bottom values (inclusive) of the Rectangle.
func (self *Rectangle) GetRandomY() float64{
    return self.Get("randomY").Float()
}

// A random value between the top and bottom values (inclusive) of the Rectangle.
func (self *Rectangle) SetRandomY(member float64) {
    self.Set("randomY", member)
}

// The y coordinate of the top of the Rectangle. Changing the top property of a Rectangle object has no effect on the x and width properties.
// However it does affect the height property, whereas changing the y value does not affect the height property.
func (self *Rectangle) GetTop() float64{
    return self.Get("top").Float()
}

// The y coordinate of the top of the Rectangle. Changing the top property of a Rectangle object has no effect on the x and width properties.
// However it does affect the height property, whereas changing the y value does not affect the height property.
func (self *Rectangle) SetTop(member float64) {
    self.Set("top", member)
}

// The location of the Rectangles top left corner as a Point object.
func (self *Rectangle) GetTopLeft() *Point{
    return &Point{self.Get("topLeft")}
}

// The location of the Rectangles top left corner as a Point object.
func (self *Rectangle) SetTopLeft(member *Point) {
    self.Set("topLeft", member)
}

// The location of the Rectangles top right corner as a Point object. The location of the Rectangles top left corner as a Point object.
func (self *Rectangle) GetTopRight() *Point{
    return &Point{self.Get("topRight")}
}

// The location of the Rectangles top right corner as a Point object. The location of the Rectangles top left corner as a Point object.
func (self *Rectangle) SetTopRight(member *Point) {
    self.Set("topRight", member)
}

// Determines whether or not this Rectangle object is empty. A Rectangle object is empty if its width or height is less than or equal to 0.
// If set to true then all of the Rectangle properties are set to 0. Gets or sets the Rectangles empty state.
func (self *Rectangle) GetEmpty() bool{
    return self.Get("empty").Bool()
}

// Determines whether or not this Rectangle object is empty. A Rectangle object is empty if its width or height is less than or equal to 0.
// If set to true then all of the Rectangle properties are set to 0. Gets or sets the Rectangles empty state.
func (self *Rectangle) SetEmpty(member bool) {
    self.Set("empty", member)
}



// Adjusts the location of the Rectangle object, as determined by its top-left corner, by the specified amounts.
func (self *Rectangle) OffsetI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("offset", args)}
}

// Adjusts the location of the Rectangle object using a Point object as a parameter. This method is similar to the Rectangle.offset() method, except that it takes a Point object as a parameter.
func (self *Rectangle) OffsetPointI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("offsetPoint", args)}
}

// Sets the members of Rectangle to the specified values.
func (self *Rectangle) SetToI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("setTo", args)}
}

// Scales the width and height of this Rectangle by the given amounts.
func (self *Rectangle) ScaleI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("scale", args)}
}

// Centers this Rectangle so that the center coordinates match the given x and y values.
func (self *Rectangle) CenterOnI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("centerOn", args)}
}

// Runs Math.floor() on both the x and y values of this Rectangle.
func (self *Rectangle) FloorI(args ...interface{}) {
    self.Call("floor", args)
}

// Runs Math.floor() on the x, y, width and height values of this Rectangle.
func (self *Rectangle) FloorAllI(args ...interface{}) {
    self.Call("floorAll", args)
}

// Runs Math.ceil() on both the x and y values of this Rectangle.
func (self *Rectangle) CeilI(args ...interface{}) {
    self.Call("ceil", args)
}

// Runs Math.ceil() on the x, y, width and height values of this Rectangle.
func (self *Rectangle) CeilAllI(args ...interface{}) {
    self.Call("ceilAll", args)
}

// Copies the x, y, width and height properties from any given object to this Rectangle.
func (self *Rectangle) CopyFromI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("copyFrom", args)}
}

// Copies the x, y, width and height properties from this Rectangle to any given object.
func (self *Rectangle) CopyToI(args ...interface{}) interface{}{
    return self.Call("copyTo", args)
}

// Increases the size of the Rectangle object by the specified amounts. The center point of the Rectangle object stays the same, and its size increases to the left and right by the dx value, and to the top and the bottom by the dy value.
func (self *Rectangle) InflateI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("inflate", args)}
}

// The size of the Rectangle object, expressed as a Point object with the values of the width and height properties.
func (self *Rectangle) SizeI(args ...interface{}) *Point{
    return &Point{self.Call("size", args)}
}

// Resize the Rectangle by providing a new width and height.
// The x and y positions remain unchanged.
func (self *Rectangle) ResizeI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("resize", args)}
}

// Returns a new Rectangle object with the same values for the x, y, width, and height properties as the original Rectangle object.
func (self *Rectangle) CloneI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("clone", args)}
}

// Determines whether the specified coordinates are contained within the region defined by this Rectangle object.
func (self *Rectangle) ContainsI(args ...interface{}) bool{
    return self.Call("contains", args).Bool()
}

// Determines whether the first Rectangle object is fully contained within the second Rectangle object.
// A Rectangle object is said to contain another if the second Rectangle object falls entirely within the boundaries of the first.
func (self *Rectangle) ContainsRectI(args ...interface{}) bool{
    return self.Call("containsRect", args).Bool()
}

// Determines whether the two Rectangles are equal.
// This method compares the x, y, width and height properties of each Rectangle.
func (self *Rectangle) EqualsI(args ...interface{}) bool{
    return self.Call("equals", args).Bool()
}

// If the Rectangle object specified in the toIntersect parameter intersects with this Rectangle object, returns the area of intersection as a Rectangle object. If the Rectangles do not intersect, this method returns an empty Rectangle object with its properties set to 0.
func (self *Rectangle) IntersectionI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("intersection", args)}
}

// Determines whether this Rectangle and another given Rectangle intersect with each other.
// This method checks the x, y, width, and height properties of the two Rectangles.
func (self *Rectangle) IntersectsI(args ...interface{}) bool{
    return self.Call("intersects", args).Bool()
}

// Determines whether the coordinates given intersects (overlaps) with this Rectangle.
func (self *Rectangle) IntersectsRawI(args ...interface{}) bool{
    return self.Call("intersectsRaw", args).Bool()
}

// Adds two Rectangles together to create a new Rectangle object, by filling in the horizontal and vertical space between the two Rectangles.
func (self *Rectangle) UnionI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("union", args)}
}

// Returns a uniformly distributed random point from anywhere within this Rectangle.
func (self *Rectangle) RandomI(args ...interface{}) *Point{
    return &Point{self.Call("random", args)}
}

// Returns a point based on the given position constant, which can be one of:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// This method returns the same values as calling Rectangle.bottomLeft, etc, but those
// calls always create a new Point object, where-as this one allows you to use your own.
func (self *Rectangle) GetPointI(args ...interface{}) *Point{
    return &Point{self.Call("getPoint", args)}
}

// Returns a string representation of this object.
func (self *Rectangle) ToStringI(args ...interface{}) string{
    return self.Call("toString", args).String()
}

// Increases the size of the Rectangle object. This method is similar to the Rectangle.inflate() method except it takes a Point object as a parameter.
func (self *Rectangle) InflatePointI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("inflatePoint", args)}
}

// Determines whether the specified coordinates are contained within the region defined by the given raw values.
func (self *Rectangle) ContainsRawI(args ...interface{}) bool{
    return self.Call("containsRaw", args).Bool()
}

// Determines whether the specified point is contained within the rectangular region defined by this Rectangle object. This method is similar to the Rectangle.contains() method, except that it takes a Point object as a parameter.
func (self *Rectangle) ContainsPointI(args ...interface{}) bool{
    return self.Call("containsPoint", args).Bool()
}

// Determines if the two objects (either Rectangles or Rectangle-like) have the same width and height values under strict equality.
func (self *Rectangle) SameDimensionsI(args ...interface{}) bool{
    return self.Call("sameDimensions", args).Bool()
}

// Calculates the Axis Aligned Bounding Box (or aabb) from an array of points.
func (self *Rectangle) AabbI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("aabb", args)}
}
