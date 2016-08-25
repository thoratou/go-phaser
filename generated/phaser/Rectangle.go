// Package phaser Automatic generation for Phaser.Rectangle
// generated file Rectangle.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Rectangle Creates a new Rectangle object with the top-left corner specified by the x and y parameters and with the specified width and height parameters.
// If you call this function without parameters, a Rectangle with x, y, width, and height properties set to 0 is created.
type Rectangle struct {
    *js.Object
}

// NewRectangle Creates a new Rectangle object with the top-left corner specified by the x and y parameters and with the specified width and height parameters.
// If you call this function without parameters, a Rectangle with x, y, width, and height properties set to 0 is created.
func NewRectangle(x int, y int, width int, height int) *Rectangle {
    return &Rectangle{js.Global.Get("Phaser").Get("Rectangle").New(x, y, width, height)}
}
// NewRectangleI Creates a new Rectangle object with the top-left corner specified by the x and y parameters and with the specified width and height parameters.
// If you call this function without parameters, a Rectangle with x, y, width, and height properties set to 0 is created.
func NewRectangleI(args ...interface{}) *Rectangle {
    return &Rectangle{js.Global.Get("Phaser").Get("Rectangle").New(args)}
}



// Rectangle Binding conversion method to Rectangle point 
func ToRectangle(jsStruct interface{}) *Rectangle {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Rectangle{Object: object}
	}
	return nil
}



// X The x coordinate of the top-left corner of the Rectangle.
func (self *Rectangle) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The x coordinate of the top-left corner of the Rectangle.
func (self *Rectangle) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The y coordinate of the top-left corner of the Rectangle.
func (self *Rectangle) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The y coordinate of the top-left corner of the Rectangle.
func (self *Rectangle) SetYA(member int) {
    self.Object.Set("y", member)
}

// Width The width of the Rectangle. This value should never be set to a negative.
func (self *Rectangle) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the Rectangle. This value should never be set to a negative.
func (self *Rectangle) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the Rectangle. This value should never be set to a negative.
func (self *Rectangle) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the Rectangle. This value should never be set to a negative.
func (self *Rectangle) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Type The const type of this object.
func (self *Rectangle) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *Rectangle) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// HalfWidth Half of the width of the Rectangle.
func (self *Rectangle) HalfWidth() int{
    return self.Object.Get("halfWidth").Int()
}

// SetHalfWidthA Half of the width of the Rectangle.
func (self *Rectangle) SetHalfWidthA(member int) {
    self.Object.Set("halfWidth", member)
}

// HalfHeight Half of the height of the Rectangle.
func (self *Rectangle) HalfHeight() int{
    return self.Object.Get("halfHeight").Int()
}

// SetHalfHeightA Half of the height of the Rectangle.
func (self *Rectangle) SetHalfHeightA(member int) {
    self.Object.Set("halfHeight", member)
}

// Bottom The sum of the y and height properties. Changing the bottom property of a Rectangle object has no effect on the x, y and width properties, but does change the height property.
func (self *Rectangle) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA The sum of the y and height properties. Changing the bottom property of a Rectangle object has no effect on the x, y and width properties, but does change the height property.
func (self *Rectangle) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// BottomLeft The location of the Rectangles bottom left corner as a Point object. Gets or sets the location of the Rectangles bottom left corner as a Point object.
func (self *Rectangle) BottomLeft() *Point{
    return &Point{self.Object.Get("bottomLeft")}
}

// SetBottomLeftA The location of the Rectangles bottom left corner as a Point object. Gets or sets the location of the Rectangles bottom left corner as a Point object.
func (self *Rectangle) SetBottomLeftA(member *Point) {
    self.Object.Set("bottomLeft", member)
}

// BottomRight The location of the Rectangles bottom right corner as a Point object. Gets or sets the location of the Rectangles bottom right corner as a Point object.
func (self *Rectangle) BottomRight() *Point{
    return &Point{self.Object.Get("bottomRight")}
}

// SetBottomRightA The location of the Rectangles bottom right corner as a Point object. Gets or sets the location of the Rectangles bottom right corner as a Point object.
func (self *Rectangle) SetBottomRightA(member *Point) {
    self.Object.Set("bottomRight", member)
}

// Left The x coordinate of the left of the Rectangle. Changing the left property of a Rectangle object has no effect on the y and height properties. However it does affect the width property, whereas changing the x value does not affect the width property.
func (self *Rectangle) Left() int{
    return self.Object.Get("left").Int()
}

// SetLeftA The x coordinate of the left of the Rectangle. Changing the left property of a Rectangle object has no effect on the y and height properties. However it does affect the width property, whereas changing the x value does not affect the width property.
func (self *Rectangle) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// Right The sum of the x and width properties. Changing the right property of a Rectangle object has no effect on the x, y and height properties, however it does affect the width property.
func (self *Rectangle) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA The sum of the x and width properties. Changing the right property of a Rectangle object has no effect on the x, y and height properties, however it does affect the width property.
func (self *Rectangle) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Volume The volume of the Rectangle derived from width * height.
func (self *Rectangle) Volume() int{
    return self.Object.Get("volume").Int()
}

// SetVolumeA The volume of the Rectangle derived from width * height.
func (self *Rectangle) SetVolumeA(member int) {
    self.Object.Set("volume", member)
}

// Perimeter The perimeter size of the Rectangle. This is the sum of all 4 sides.
func (self *Rectangle) Perimeter() int{
    return self.Object.Get("perimeter").Int()
}

// SetPerimeterA The perimeter size of the Rectangle. This is the sum of all 4 sides.
func (self *Rectangle) SetPerimeterA(member int) {
    self.Object.Set("perimeter", member)
}

// CenterX The x coordinate of the center of the Rectangle.
func (self *Rectangle) CenterX() int{
    return self.Object.Get("centerX").Int()
}

// SetCenterXA The x coordinate of the center of the Rectangle.
func (self *Rectangle) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// CenterY The y coordinate of the center of the Rectangle.
func (self *Rectangle) CenterY() int{
    return self.Object.Get("centerY").Int()
}

// SetCenterYA The y coordinate of the center of the Rectangle.
func (self *Rectangle) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// RandomX A random value between the left and right values (inclusive) of the Rectangle.
func (self *Rectangle) RandomX() int{
    return self.Object.Get("randomX").Int()
}

// SetRandomXA A random value between the left and right values (inclusive) of the Rectangle.
func (self *Rectangle) SetRandomXA(member int) {
    self.Object.Set("randomX", member)
}

// RandomY A random value between the top and bottom values (inclusive) of the Rectangle.
func (self *Rectangle) RandomY() int{
    return self.Object.Get("randomY").Int()
}

// SetRandomYA A random value between the top and bottom values (inclusive) of the Rectangle.
func (self *Rectangle) SetRandomYA(member int) {
    self.Object.Set("randomY", member)
}

// Top The y coordinate of the top of the Rectangle. Changing the top property of a Rectangle object has no effect on the x and width properties.
// However it does affect the height property, whereas changing the y value does not affect the height property.
func (self *Rectangle) Top() int{
    return self.Object.Get("top").Int()
}

// SetTopA The y coordinate of the top of the Rectangle. Changing the top property of a Rectangle object has no effect on the x and width properties.
// However it does affect the height property, whereas changing the y value does not affect the height property.
func (self *Rectangle) SetTopA(member int) {
    self.Object.Set("top", member)
}

// TopLeft The location of the Rectangles top left corner as a Point object.
func (self *Rectangle) TopLeft() *Point{
    return &Point{self.Object.Get("topLeft")}
}

// SetTopLeftA The location of the Rectangles top left corner as a Point object.
func (self *Rectangle) SetTopLeftA(member *Point) {
    self.Object.Set("topLeft", member)
}

// TopRight The location of the Rectangles top right corner as a Point object. The location of the Rectangles top left corner as a Point object.
func (self *Rectangle) TopRight() *Point{
    return &Point{self.Object.Get("topRight")}
}

// SetTopRightA The location of the Rectangles top right corner as a Point object. The location of the Rectangles top left corner as a Point object.
func (self *Rectangle) SetTopRightA(member *Point) {
    self.Object.Set("topRight", member)
}

// Empty Determines whether or not this Rectangle object is empty. A Rectangle object is empty if its width or height is less than or equal to 0.
// If set to true then all of the Rectangle properties are set to 0. Gets or sets the Rectangles empty state.
func (self *Rectangle) Empty() bool{
    return self.Object.Get("empty").Bool()
}

// SetEmptyA Determines whether or not this Rectangle object is empty. A Rectangle object is empty if its width or height is less than or equal to 0.
// If set to true then all of the Rectangle properties are set to 0. Gets or sets the Rectangles empty state.
func (self *Rectangle) SetEmptyA(member bool) {
    self.Object.Set("empty", member)
}


// Offset Adjusts the location of the Rectangle object, as determined by its top-left corner, by the specified amounts.
func (self *Rectangle) Offset(dx int, dy int) *Rectangle{
    return &Rectangle{self.Object.Call("offset", dx, dy)}
}

// OffsetI Adjusts the location of the Rectangle object, as determined by its top-left corner, by the specified amounts.
func (self *Rectangle) OffsetI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("offset", args)}
}

// OffsetPoint Adjusts the location of the Rectangle object using a Point object as a parameter. This method is similar to the Rectangle.offset() method, except that it takes a Point object as a parameter.
func (self *Rectangle) OffsetPoint(point *Point) *Rectangle{
    return &Rectangle{self.Object.Call("offsetPoint", point)}
}

// OffsetPointI Adjusts the location of the Rectangle object using a Point object as a parameter. This method is similar to the Rectangle.offset() method, except that it takes a Point object as a parameter.
func (self *Rectangle) OffsetPointI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("offsetPoint", args)}
}

// SetTo Sets the members of Rectangle to the specified values.
func (self *Rectangle) SetTo(x int, y int, width int, height int) *Rectangle{
    return &Rectangle{self.Object.Call("setTo", x, y, width, height)}
}

// SetToI Sets the members of Rectangle to the specified values.
func (self *Rectangle) SetToI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("setTo", args)}
}

// Scale Scales the width and height of this Rectangle by the given amounts.
func (self *Rectangle) Scale(x int) *Rectangle{
    return &Rectangle{self.Object.Call("scale", x)}
}

// Scale1O Scales the width and height of this Rectangle by the given amounts.
func (self *Rectangle) Scale1O(x int, y int) *Rectangle{
    return &Rectangle{self.Object.Call("scale", x, y)}
}

// ScaleI Scales the width and height of this Rectangle by the given amounts.
func (self *Rectangle) ScaleI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("scale", args)}
}

// CenterOn Centers this Rectangle so that the center coordinates match the given x and y values.
func (self *Rectangle) CenterOn(x int, y int) *Rectangle{
    return &Rectangle{self.Object.Call("centerOn", x, y)}
}

// CenterOnI Centers this Rectangle so that the center coordinates match the given x and y values.
func (self *Rectangle) CenterOnI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("centerOn", args)}
}

// Floor Runs Math.floor() on both the x and y values of this Rectangle.
func (self *Rectangle) Floor() {
    self.Object.Call("floor")
}

// FloorI Runs Math.floor() on both the x and y values of this Rectangle.
func (self *Rectangle) FloorI(args ...interface{}) {
    self.Object.Call("floor", args)
}

// FloorAll Runs Math.floor() on the x, y, width and height values of this Rectangle.
func (self *Rectangle) FloorAll() {
    self.Object.Call("floorAll")
}

// FloorAllI Runs Math.floor() on the x, y, width and height values of this Rectangle.
func (self *Rectangle) FloorAllI(args ...interface{}) {
    self.Object.Call("floorAll", args)
}

// Ceil Runs Math.ceil() on both the x and y values of this Rectangle.
func (self *Rectangle) Ceil() {
    self.Object.Call("ceil")
}

// CeilI Runs Math.ceil() on both the x and y values of this Rectangle.
func (self *Rectangle) CeilI(args ...interface{}) {
    self.Object.Call("ceil", args)
}

// CeilAll Runs Math.ceil() on the x, y, width and height values of this Rectangle.
func (self *Rectangle) CeilAll() {
    self.Object.Call("ceilAll")
}

// CeilAllI Runs Math.ceil() on the x, y, width and height values of this Rectangle.
func (self *Rectangle) CeilAllI(args ...interface{}) {
    self.Object.Call("ceilAll", args)
}

// CopyFrom Copies the x, y, width and height properties from any given object to this Rectangle.
func (self *Rectangle) CopyFrom(source interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("copyFrom", source)}
}

// CopyFromI Copies the x, y, width and height properties from any given object to this Rectangle.
func (self *Rectangle) CopyFromI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("copyFrom", args)}
}

// CopyTo Copies the x, y, width and height properties from this Rectangle to any given object.
func (self *Rectangle) CopyTo(source interface{}) interface{}{
    return self.Object.Call("copyTo", source)
}

// CopyToI Copies the x, y, width and height properties from this Rectangle to any given object.
func (self *Rectangle) CopyToI(args ...interface{}) interface{}{
    return self.Object.Call("copyTo", args)
}

// Inflate Increases the size of the Rectangle object by the specified amounts. The center point of the Rectangle object stays the same, and its size increases to the left and right by the dx value, and to the top and the bottom by the dy value.
func (self *Rectangle) Inflate(dx int, dy int) *Rectangle{
    return &Rectangle{self.Object.Call("inflate", dx, dy)}
}

// InflateI Increases the size of the Rectangle object by the specified amounts. The center point of the Rectangle object stays the same, and its size increases to the left and right by the dx value, and to the top and the bottom by the dy value.
func (self *Rectangle) InflateI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("inflate", args)}
}

// Size The size of the Rectangle object, expressed as a Point object with the values of the width and height properties.
func (self *Rectangle) Size() *Point{
    return &Point{self.Object.Call("size")}
}

// Size1O The size of the Rectangle object, expressed as a Point object with the values of the width and height properties.
func (self *Rectangle) Size1O(output *Point) *Point{
    return &Point{self.Object.Call("size", output)}
}

// SizeI The size of the Rectangle object, expressed as a Point object with the values of the width and height properties.
func (self *Rectangle) SizeI(args ...interface{}) *Point{
    return &Point{self.Object.Call("size", args)}
}

// Resize Resize the Rectangle by providing a new width and height.
// The x and y positions remain unchanged.
func (self *Rectangle) Resize(width int, height int) *Rectangle{
    return &Rectangle{self.Object.Call("resize", width, height)}
}

// ResizeI Resize the Rectangle by providing a new width and height.
// The x and y positions remain unchanged.
func (self *Rectangle) ResizeI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("resize", args)}
}

// Clone Returns a new Rectangle object with the same values for the x, y, width, and height properties as the original Rectangle object.
func (self *Rectangle) Clone() *Rectangle{
    return &Rectangle{self.Object.Call("clone")}
}

// Clone1O Returns a new Rectangle object with the same values for the x, y, width, and height properties as the original Rectangle object.
func (self *Rectangle) Clone1O(output *Rectangle) *Rectangle{
    return &Rectangle{self.Object.Call("clone", output)}
}

// CloneI Returns a new Rectangle object with the same values for the x, y, width, and height properties as the original Rectangle object.
func (self *Rectangle) CloneI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("clone", args)}
}

// Contains Determines whether the specified coordinates are contained within the region defined by this Rectangle object.
func (self *Rectangle) Contains(x int, y int) bool{
    return self.Object.Call("contains", x, y).Bool()
}

// ContainsI Determines whether the specified coordinates are contained within the region defined by this Rectangle object.
func (self *Rectangle) ContainsI(args ...interface{}) bool{
    return self.Object.Call("contains", args).Bool()
}

// ContainsRect Determines whether the first Rectangle object is fully contained within the second Rectangle object.
// A Rectangle object is said to contain another if the second Rectangle object falls entirely within the boundaries of the first.
func (self *Rectangle) ContainsRect(b *Rectangle) bool{
    return self.Object.Call("containsRect", b).Bool()
}

// ContainsRectI Determines whether the first Rectangle object is fully contained within the second Rectangle object.
// A Rectangle object is said to contain another if the second Rectangle object falls entirely within the boundaries of the first.
func (self *Rectangle) ContainsRectI(args ...interface{}) bool{
    return self.Object.Call("containsRect", args).Bool()
}

// Equals Determines whether the two Rectangles are equal.
// This method compares the x, y, width and height properties of each Rectangle.
func (self *Rectangle) Equals(b *Rectangle) bool{
    return self.Object.Call("equals", b).Bool()
}

// EqualsI Determines whether the two Rectangles are equal.
// This method compares the x, y, width and height properties of each Rectangle.
func (self *Rectangle) EqualsI(args ...interface{}) bool{
    return self.Object.Call("equals", args).Bool()
}

// Intersection If the Rectangle object specified in the toIntersect parameter intersects with this Rectangle object, returns the area of intersection as a Rectangle object. If the Rectangles do not intersect, this method returns an empty Rectangle object with its properties set to 0.
func (self *Rectangle) Intersection(b *Rectangle, out *Rectangle) *Rectangle{
    return &Rectangle{self.Object.Call("intersection", b, out)}
}

// IntersectionI If the Rectangle object specified in the toIntersect parameter intersects with this Rectangle object, returns the area of intersection as a Rectangle object. If the Rectangles do not intersect, this method returns an empty Rectangle object with its properties set to 0.
func (self *Rectangle) IntersectionI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("intersection", args)}
}

// Intersects Determines whether this Rectangle and another given Rectangle intersect with each other.
// This method checks the x, y, width, and height properties of the two Rectangles.
func (self *Rectangle) Intersects(b *Rectangle) bool{
    return self.Object.Call("intersects", b).Bool()
}

// IntersectsI Determines whether this Rectangle and another given Rectangle intersect with each other.
// This method checks the x, y, width, and height properties of the two Rectangles.
func (self *Rectangle) IntersectsI(args ...interface{}) bool{
    return self.Object.Call("intersects", args).Bool()
}

// IntersectsRaw Determines whether the coordinates given intersects (overlaps) with this Rectangle.
func (self *Rectangle) IntersectsRaw(left int, right int, top int, bottom int, tolerance int) bool{
    return self.Object.Call("intersectsRaw", left, right, top, bottom, tolerance).Bool()
}

// IntersectsRawI Determines whether the coordinates given intersects (overlaps) with this Rectangle.
func (self *Rectangle) IntersectsRawI(args ...interface{}) bool{
    return self.Object.Call("intersectsRaw", args).Bool()
}

// Union Adds two Rectangles together to create a new Rectangle object, by filling in the horizontal and vertical space between the two Rectangles.
func (self *Rectangle) Union(b *Rectangle) *Rectangle{
    return &Rectangle{self.Object.Call("union", b)}
}

// Union1O Adds two Rectangles together to create a new Rectangle object, by filling in the horizontal and vertical space between the two Rectangles.
func (self *Rectangle) Union1O(b *Rectangle, out *Rectangle) *Rectangle{
    return &Rectangle{self.Object.Call("union", b, out)}
}

// UnionI Adds two Rectangles together to create a new Rectangle object, by filling in the horizontal and vertical space between the two Rectangles.
func (self *Rectangle) UnionI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("union", args)}
}

// Random Returns a uniformly distributed random point from anywhere within this Rectangle.
func (self *Rectangle) Random() *Point{
    return &Point{self.Object.Call("random")}
}

// Random1O Returns a uniformly distributed random point from anywhere within this Rectangle.
func (self *Rectangle) Random1O(out interface{}) *Point{
    return &Point{self.Object.Call("random", out)}
}

// RandomI Returns a uniformly distributed random point from anywhere within this Rectangle.
func (self *Rectangle) RandomI(args ...interface{}) *Point{
    return &Point{self.Object.Call("random", args)}
}

// GetPoint Returns a point based on the given position constant, which can be one of:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// This method returns the same values as calling Rectangle.bottomLeft, etc, but those
// calls always create a new Point object, where-as this one allows you to use your own.
func (self *Rectangle) GetPoint() *Point{
    return &Point{self.Object.Call("getPoint")}
}

// GetPoint1O Returns a point based on the given position constant, which can be one of:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// This method returns the same values as calling Rectangle.bottomLeft, etc, but those
// calls always create a new Point object, where-as this one allows you to use your own.
func (self *Rectangle) GetPoint1O(position int) *Point{
    return &Point{self.Object.Call("getPoint", position)}
}

// GetPoint2O Returns a point based on the given position constant, which can be one of:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// This method returns the same values as calling Rectangle.bottomLeft, etc, but those
// calls always create a new Point object, where-as this one allows you to use your own.
func (self *Rectangle) GetPoint2O(position int, out *Point) *Point{
    return &Point{self.Object.Call("getPoint", position, out)}
}

// GetPointI Returns a point based on the given position constant, which can be one of:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`,
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// This method returns the same values as calling Rectangle.bottomLeft, etc, but those
// calls always create a new Point object, where-as this one allows you to use your own.
func (self *Rectangle) GetPointI(args ...interface{}) *Point{
    return &Point{self.Object.Call("getPoint", args)}
}

// ToString Returns a string representation of this object.
func (self *Rectangle) ToString() string{
    return self.Object.Call("toString").String()
}

// ToStringI Returns a string representation of this object.
func (self *Rectangle) ToStringI(args ...interface{}) string{
    return self.Object.Call("toString", args).String()
}

// InflatePoint Increases the size of the Rectangle object. This method is similar to the Rectangle.inflate() method except it takes a Point object as a parameter.
func (self *Rectangle) InflatePoint(a *Rectangle, point *Point) *Rectangle{
    return &Rectangle{self.Object.Call("inflatePoint", a, point)}
}

// InflatePointI Increases the size of the Rectangle object. This method is similar to the Rectangle.inflate() method except it takes a Point object as a parameter.
func (self *Rectangle) InflatePointI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("inflatePoint", args)}
}

// ContainsRaw Determines whether the specified coordinates are contained within the region defined by the given raw values.
func (self *Rectangle) ContainsRaw(rx int, ry int, rw int, rh int, x int, y int) bool{
    return self.Object.Call("containsRaw", rx, ry, rw, rh, x, y).Bool()
}

// ContainsRawI Determines whether the specified coordinates are contained within the region defined by the given raw values.
func (self *Rectangle) ContainsRawI(args ...interface{}) bool{
    return self.Object.Call("containsRaw", args).Bool()
}

// ContainsPoint Determines whether the specified point is contained within the rectangular region defined by this Rectangle object. This method is similar to the Rectangle.contains() method, except that it takes a Point object as a parameter.
func (self *Rectangle) ContainsPoint(a *Rectangle, point *Point) bool{
    return self.Object.Call("containsPoint", a, point).Bool()
}

// ContainsPointI Determines whether the specified point is contained within the rectangular region defined by this Rectangle object. This method is similar to the Rectangle.contains() method, except that it takes a Point object as a parameter.
func (self *Rectangle) ContainsPointI(args ...interface{}) bool{
    return self.Object.Call("containsPoint", args).Bool()
}

// SameDimensions Determines if the two objects (either Rectangles or Rectangle-like) have the same width and height values under strict equality.
func (self *Rectangle) SameDimensions(a interface{}, b interface{}) bool{
    return self.Object.Call("sameDimensions", a, b).Bool()
}

// SameDimensionsI Determines if the two objects (either Rectangles or Rectangle-like) have the same width and height values under strict equality.
func (self *Rectangle) SameDimensionsI(args ...interface{}) bool{
    return self.Object.Call("sameDimensions", args).Bool()
}

// Aabb Calculates the Axis Aligned Bounding Box (or aabb) from an array of points.
func (self *Rectangle) Aabb(points []Point) *Rectangle{
    return &Rectangle{self.Object.Call("aabb", points)}
}

// Aabb1O Calculates the Axis Aligned Bounding Box (or aabb) from an array of points.
func (self *Rectangle) Aabb1O(points []Point, out *Rectangle) *Rectangle{
    return &Rectangle{self.Object.Call("aabb", points, out)}
}

// AabbI Calculates the Axis Aligned Bounding Box (or aabb) from an array of points.
func (self *Rectangle) AabbI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("aabb", args)}
}

