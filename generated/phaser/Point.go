// Package phaser Automatic generation for Phaser.Point
// generated file Point.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Point A Point object represents a location in a two-dimensional coordinate system, where x represents the horizontal axis and y represents the vertical axis.
// The following code creates a point at (0,0):
// `var myPoint = new Phaser.Point();`
// You can also use them as 2D Vectors and you'll find different vector related methods in this class.
type Point struct {
    *js.Object
}

// NewPoint A Point object represents a location in a two-dimensional coordinate system, where x represents the horizontal axis and y represents the vertical axis.
// The following code creates a point at (0,0):
// `var myPoint = new Phaser.Point();`
// You can also use them as 2D Vectors and you'll find different vector related methods in this class.
func NewPoint() *Point {
    return &Point{js.Global.Get("Phaser").Get("Point").New()}
}
// NewPoint1O A Point object represents a location in a two-dimensional coordinate system, where x represents the horizontal axis and y represents the vertical axis.
// The following code creates a point at (0,0):
// `var myPoint = new Phaser.Point();`
// You can also use them as 2D Vectors and you'll find different vector related methods in this class.
func NewPoint1O(x int) *Point {
    return &Point{js.Global.Get("Phaser").Get("Point").New(x)}
}
// NewPoint2O A Point object represents a location in a two-dimensional coordinate system, where x represents the horizontal axis and y represents the vertical axis.
// The following code creates a point at (0,0):
// `var myPoint = new Phaser.Point();`
// You can also use them as 2D Vectors and you'll find different vector related methods in this class.
func NewPoint2O(x int, y int) *Point {
    return &Point{js.Global.Get("Phaser").Get("Point").New(x, y)}
}
// NewPointI A Point object represents a location in a two-dimensional coordinate system, where x represents the horizontal axis and y represents the vertical axis.
// The following code creates a point at (0,0):
// `var myPoint = new Phaser.Point();`
// You can also use them as 2D Vectors and you'll find different vector related methods in this class.
func NewPointI(args ...interface{}) *Point {
    return &Point{js.Global.Get("Phaser").Get("Point").New(args)}
}



// X The x value of the point.
func (self *Point) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The x value of the point.
func (self *Point) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The y value of the point.
func (self *Point) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The y value of the point.
func (self *Point) SetYA(member int) {
    self.Object.Set("y", member)
}

// Type The const type of this object.
func (self *Point) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *Point) SetTypeA(member int) {
    self.Object.Set("type", member)
}


// CopyFrom Copies the x and y properties from any given object to this Point.
func (self *Point) CopyFrom(source interface{}) *Point{
    return &Point{self.Object.Call("copyFrom", source)}
}

// CopyFromI Copies the x and y properties from any given object to this Point.
func (self *Point) CopyFromI(args ...interface{}) *Point{
    return &Point{self.Object.Call("copyFrom", args)}
}

// Invert Inverts the x and y values of this Point
func (self *Point) Invert() *Point{
    return &Point{self.Object.Call("invert")}
}

// InvertI Inverts the x and y values of this Point
func (self *Point) InvertI(args ...interface{}) *Point{
    return &Point{self.Object.Call("invert", args)}
}

// SetTo Sets the `x` and `y` values of this Point object to the given values.
// If you omit the `y` value then the `x` value will be applied to both, for example:
// `Point.setTo(2)` is the same as `Point.setTo(2, 2)`
func (self *Point) SetTo(x int) *Point{
    return &Point{self.Object.Call("setTo", x)}
}

// SetTo1O Sets the `x` and `y` values of this Point object to the given values.
// If you omit the `y` value then the `x` value will be applied to both, for example:
// `Point.setTo(2)` is the same as `Point.setTo(2, 2)`
func (self *Point) SetTo1O(x int, y int) *Point{
    return &Point{self.Object.Call("setTo", x, y)}
}

// SetToI Sets the `x` and `y` values of this Point object to the given values.
// If you omit the `y` value then the `x` value will be applied to both, for example:
// `Point.setTo(2)` is the same as `Point.setTo(2, 2)`
func (self *Point) SetToI(args ...interface{}) *Point{
    return &Point{self.Object.Call("setTo", args)}
}

// Set Sets the `x` and `y` values of this Point object to the given values.
// If you omit the `y` value then the `x` value will be applied to both, for example:
// `Point.set(2)` is the same as `Point.set(2, 2)`
func (self *Point) Set(x int) *Point{
    return &Point{self.Object.Call("set", x)}
}

// Set1O Sets the `x` and `y` values of this Point object to the given values.
// If you omit the `y` value then the `x` value will be applied to both, for example:
// `Point.set(2)` is the same as `Point.set(2, 2)`
func (self *Point) Set1O(x int, y int) *Point{
    return &Point{self.Object.Call("set", x, y)}
}

// SetI Sets the `x` and `y` values of this Point object to the given values.
// If you omit the `y` value then the `x` value will be applied to both, for example:
// `Point.set(2)` is the same as `Point.set(2, 2)`
func (self *Point) SetI(args ...interface{}) *Point{
    return &Point{self.Object.Call("set", args)}
}

// Add Adds the given x and y values to this Point.
func (self *Point) Add(x int, y int) *Point{
    return &Point{self.Object.Call("add", x, y)}
}

// AddI Adds the given x and y values to this Point.
func (self *Point) AddI(args ...interface{}) *Point{
    return &Point{self.Object.Call("add", args)}
}

// Subtract Subtracts the given x and y values from this Point.
func (self *Point) Subtract(x int, y int) *Point{
    return &Point{self.Object.Call("subtract", x, y)}
}

// SubtractI Subtracts the given x and y values from this Point.
func (self *Point) SubtractI(args ...interface{}) *Point{
    return &Point{self.Object.Call("subtract", args)}
}

// Multiply Multiplies Point.x and Point.y by the given x and y values. Sometimes known as `Scale`.
func (self *Point) Multiply(x int, y int) *Point{
    return &Point{self.Object.Call("multiply", x, y)}
}

// MultiplyI Multiplies Point.x and Point.y by the given x and y values. Sometimes known as `Scale`.
func (self *Point) MultiplyI(args ...interface{}) *Point{
    return &Point{self.Object.Call("multiply", args)}
}

// Divide Divides Point.x and Point.y by the given x and y values.
func (self *Point) Divide(x int, y int) *Point{
    return &Point{self.Object.Call("divide", x, y)}
}

// DivideI Divides Point.x and Point.y by the given x and y values.
func (self *Point) DivideI(args ...interface{}) *Point{
    return &Point{self.Object.Call("divide", args)}
}

// ClampX Clamps the x value of this Point to be between the given min and max.
func (self *Point) ClampX(min int, max int) *Point{
    return &Point{self.Object.Call("clampX", min, max)}
}

// ClampXI Clamps the x value of this Point to be between the given min and max.
func (self *Point) ClampXI(args ...interface{}) *Point{
    return &Point{self.Object.Call("clampX", args)}
}

// ClampY Clamps the y value of this Point to be between the given min and max
func (self *Point) ClampY(min int, max int) *Point{
    return &Point{self.Object.Call("clampY", min, max)}
}

// ClampYI Clamps the y value of this Point to be between the given min and max
func (self *Point) ClampYI(args ...interface{}) *Point{
    return &Point{self.Object.Call("clampY", args)}
}

// Clamp Clamps this Point object values to be between the given min and max.
func (self *Point) Clamp(min int, max int) *Point{
    return &Point{self.Object.Call("clamp", min, max)}
}

// ClampI Clamps this Point object values to be between the given min and max.
func (self *Point) ClampI(args ...interface{}) *Point{
    return &Point{self.Object.Call("clamp", args)}
}

// Clone Creates a copy of the given Point.
func (self *Point) Clone() *Point{
    return &Point{self.Object.Call("clone")}
}

// Clone1O Creates a copy of the given Point.
func (self *Point) Clone1O(output *Point) *Point{
    return &Point{self.Object.Call("clone", output)}
}

// CloneI Creates a copy of the given Point.
func (self *Point) CloneI(args ...interface{}) *Point{
    return &Point{self.Object.Call("clone", args)}
}

// CopyTo Copies the x and y properties from this Point to any given object.
func (self *Point) CopyTo(dest interface{}) interface{}{
    return self.Object.Call("copyTo", dest)
}

// CopyToI Copies the x and y properties from this Point to any given object.
func (self *Point) CopyToI(args ...interface{}) interface{}{
    return self.Object.Call("copyTo", args)
}

// Distance Returns the distance of this Point object to the given object (can be a Circle, Point or anything with x/y properties)
func (self *Point) Distance(dest interface{}) int{
    return self.Object.Call("distance", dest).Int()
}

// Distance1O Returns the distance of this Point object to the given object (can be a Circle, Point or anything with x/y properties)
func (self *Point) Distance1O(dest interface{}, round bool) int{
    return self.Object.Call("distance", dest, round).Int()
}

// DistanceI Returns the distance of this Point object to the given object (can be a Circle, Point or anything with x/y properties)
func (self *Point) DistanceI(args ...interface{}) int{
    return self.Object.Call("distance", args).Int()
}

// Equals Determines whether the given objects x/y values are equal to this Point object.
func (self *Point) Equals(a interface{}) bool{
    return self.Object.Call("equals", a).Bool()
}

// EqualsI Determines whether the given objects x/y values are equal to this Point object.
func (self *Point) EqualsI(args ...interface{}) bool{
    return self.Object.Call("equals", args).Bool()
}

// Angle Returns the angle between this Point object and another object with public x and y properties.
func (self *Point) Angle(a interface{}) int{
    return self.Object.Call("angle", a).Int()
}

// Angle1O Returns the angle between this Point object and another object with public x and y properties.
func (self *Point) Angle1O(a interface{}, asDegrees bool) int{
    return self.Object.Call("angle", a, asDegrees).Int()
}

// AngleI Returns the angle between this Point object and another object with public x and y properties.
func (self *Point) AngleI(args ...interface{}) int{
    return self.Object.Call("angle", args).Int()
}

// Rotate Rotates this Point around the x/y coordinates given to the desired angle.
func (self *Point) Rotate(x int, y int, angle int) *Point{
    return &Point{self.Object.Call("rotate", x, y, angle)}
}

// Rotate1O Rotates this Point around the x/y coordinates given to the desired angle.
func (self *Point) Rotate1O(x int, y int, angle int, asDegrees bool) *Point{
    return &Point{self.Object.Call("rotate", x, y, angle, asDegrees)}
}

// Rotate2O Rotates this Point around the x/y coordinates given to the desired angle.
func (self *Point) Rotate2O(x int, y int, angle int, asDegrees bool, distance int) *Point{
    return &Point{self.Object.Call("rotate", x, y, angle, asDegrees, distance)}
}

// RotateI Rotates this Point around the x/y coordinates given to the desired angle.
func (self *Point) RotateI(args ...interface{}) *Point{
    return &Point{self.Object.Call("rotate", args)}
}

// GetMagnitude Calculates the length of the Point object.
func (self *Point) GetMagnitude() int{
    return self.Object.Call("getMagnitude").Int()
}

// GetMagnitudeI Calculates the length of the Point object.
func (self *Point) GetMagnitudeI(args ...interface{}) int{
    return self.Object.Call("getMagnitude", args).Int()
}

// GetMagnitudeSq Calculates the length squared of the Point object.
func (self *Point) GetMagnitudeSq() int{
    return self.Object.Call("getMagnitudeSq").Int()
}

// GetMagnitudeSqI Calculates the length squared of the Point object.
func (self *Point) GetMagnitudeSqI(args ...interface{}) int{
    return self.Object.Call("getMagnitudeSq", args).Int()
}

// SetMagnitude Alters the length of the Point without changing the direction.
func (self *Point) SetMagnitude(magnitude int) *Point{
    return &Point{self.Object.Call("setMagnitude", magnitude)}
}

// SetMagnitudeI Alters the length of the Point without changing the direction.
func (self *Point) SetMagnitudeI(args ...interface{}) *Point{
    return &Point{self.Object.Call("setMagnitude", args)}
}

// Normalize Alters the Point object so that its length is 1, but it retains the same direction.
func (self *Point) Normalize() *Point{
    return &Point{self.Object.Call("normalize")}
}

// NormalizeI Alters the Point object so that its length is 1, but it retains the same direction.
func (self *Point) NormalizeI(args ...interface{}) *Point{
    return &Point{self.Object.Call("normalize", args)}
}

// IsZero Determine if this point is at 0,0.
func (self *Point) IsZero() bool{
    return self.Object.Call("isZero").Bool()
}

// IsZeroI Determine if this point is at 0,0.
func (self *Point) IsZeroI(args ...interface{}) bool{
    return self.Object.Call("isZero", args).Bool()
}

// Dot The dot product of this and another Point object.
func (self *Point) Dot(a *Point) int{
    return self.Object.Call("dot", a).Int()
}

// DotI The dot product of this and another Point object.
func (self *Point) DotI(args ...interface{}) int{
    return self.Object.Call("dot", args).Int()
}

// Cross The cross product of this and another Point object.
func (self *Point) Cross(a *Point) int{
    return self.Object.Call("cross", a).Int()
}

// CrossI The cross product of this and another Point object.
func (self *Point) CrossI(args ...interface{}) int{
    return self.Object.Call("cross", args).Int()
}

// Perp Make this Point perpendicular (90 degrees rotation)
func (self *Point) Perp() *Point{
    return &Point{self.Object.Call("perp")}
}

// PerpI Make this Point perpendicular (90 degrees rotation)
func (self *Point) PerpI(args ...interface{}) *Point{
    return &Point{self.Object.Call("perp", args)}
}

// Rperp Make this Point perpendicular (-90 degrees rotation)
func (self *Point) Rperp() *Point{
    return &Point{self.Object.Call("rperp")}
}

// RperpI Make this Point perpendicular (-90 degrees rotation)
func (self *Point) RperpI(args ...interface{}) *Point{
    return &Point{self.Object.Call("rperp", args)}
}

// NormalRightHand Right-hand normalize (make unit length) this Point.
func (self *Point) NormalRightHand() *Point{
    return &Point{self.Object.Call("normalRightHand")}
}

// NormalRightHandI Right-hand normalize (make unit length) this Point.
func (self *Point) NormalRightHandI(args ...interface{}) *Point{
    return &Point{self.Object.Call("normalRightHand", args)}
}

// Floor Math.floor() both the x and y properties of this Point.
func (self *Point) Floor() *Point{
    return &Point{self.Object.Call("floor")}
}

// FloorI Math.floor() both the x and y properties of this Point.
func (self *Point) FloorI(args ...interface{}) *Point{
    return &Point{self.Object.Call("floor", args)}
}

// Ceil Math.ceil() both the x and y properties of this Point.
func (self *Point) Ceil() *Point{
    return &Point{self.Object.Call("ceil")}
}

// CeilI Math.ceil() both the x and y properties of this Point.
func (self *Point) CeilI(args ...interface{}) *Point{
    return &Point{self.Object.Call("ceil", args)}
}

// ToString Returns a string representation of this object.
func (self *Point) ToString() string{
    return self.Object.Call("toString").String()
}

// ToStringI Returns a string representation of this object.
func (self *Point) ToStringI(args ...interface{}) string{
    return self.Object.Call("toString", args).String()
}

// Negative Creates a negative Point.
func (self *Point) Negative(a *Point) *Point{
    return &Point{self.Object.Call("negative", a)}
}

// Negative1O Creates a negative Point.
func (self *Point) Negative1O(a *Point, out *Point) *Point{
    return &Point{self.Object.Call("negative", a, out)}
}

// NegativeI Creates a negative Point.
func (self *Point) NegativeI(args ...interface{}) *Point{
    return &Point{self.Object.Call("negative", args)}
}

// MultiplyAdd Adds two 2D Points together and multiplies the result by the given scalar.
func (self *Point) MultiplyAdd(a *Point, b *Point, s int) *Point{
    return &Point{self.Object.Call("multiplyAdd", a, b, s)}
}

// MultiplyAdd1O Adds two 2D Points together and multiplies the result by the given scalar.
func (self *Point) MultiplyAdd1O(a *Point, b *Point, s int, out *Point) *Point{
    return &Point{self.Object.Call("multiplyAdd", a, b, s, out)}
}

// MultiplyAddI Adds two 2D Points together and multiplies the result by the given scalar.
func (self *Point) MultiplyAddI(args ...interface{}) *Point{
    return &Point{self.Object.Call("multiplyAdd", args)}
}

// Interpolate Interpolates the two given Points, based on the `f` value (between 0 and 1) and returns a new Point.
func (self *Point) Interpolate(a *Point, b *Point, f int) *Point{
    return &Point{self.Object.Call("interpolate", a, b, f)}
}

// Interpolate1O Interpolates the two given Points, based on the `f` value (between 0 and 1) and returns a new Point.
func (self *Point) Interpolate1O(a *Point, b *Point, f int, out *Point) *Point{
    return &Point{self.Object.Call("interpolate", a, b, f, out)}
}

// InterpolateI Interpolates the two given Points, based on the `f` value (between 0 and 1) and returns a new Point.
func (self *Point) InterpolateI(args ...interface{}) *Point{
    return &Point{self.Object.Call("interpolate", args)}
}

// Project Project two Points onto another Point.
func (self *Point) Project(a *Point, b *Point) *Point{
    return &Point{self.Object.Call("project", a, b)}
}

// Project1O Project two Points onto another Point.
func (self *Point) Project1O(a *Point, b *Point, out *Point) *Point{
    return &Point{self.Object.Call("project", a, b, out)}
}

// ProjectI Project two Points onto another Point.
func (self *Point) ProjectI(args ...interface{}) *Point{
    return &Point{self.Object.Call("project", args)}
}

// ProjectUnit Project two Points onto a Point of unit length.
func (self *Point) ProjectUnit(a *Point, b *Point) *Point{
    return &Point{self.Object.Call("projectUnit", a, b)}
}

// ProjectUnit1O Project two Points onto a Point of unit length.
func (self *Point) ProjectUnit1O(a *Point, b *Point, out *Point) *Point{
    return &Point{self.Object.Call("projectUnit", a, b, out)}
}

// ProjectUnitI Project two Points onto a Point of unit length.
func (self *Point) ProjectUnitI(args ...interface{}) *Point{
    return &Point{self.Object.Call("projectUnit", args)}
}

// Centroid Calculates centroid (or midpoint) from an array of points. If only one point is provided, that point is returned.
func (self *Point) Centroid(points []Point) *Point{
    return &Point{self.Object.Call("centroid", points)}
}

// Centroid1O Calculates centroid (or midpoint) from an array of points. If only one point is provided, that point is returned.
func (self *Point) Centroid1O(points []Point, out *Point) *Point{
    return &Point{self.Object.Call("centroid", points, out)}
}

// CentroidI Calculates centroid (or midpoint) from an array of points. If only one point is provided, that point is returned.
func (self *Point) CentroidI(args ...interface{}) *Point{
    return &Point{self.Object.Call("centroid", args)}
}

// Parse Parses an object for x and/or y properties and returns a new Phaser.Point with matching values.
// If the object doesn't contain those properties a Point with x/y of zero will be returned.
func (self *Point) Parse(obj interface{}) *Point{
    return &Point{self.Object.Call("parse", obj)}
}

// Parse1O Parses an object for x and/or y properties and returns a new Phaser.Point with matching values.
// If the object doesn't contain those properties a Point with x/y of zero will be returned.
func (self *Point) Parse1O(obj interface{}, xProp string) *Point{
    return &Point{self.Object.Call("parse", obj, xProp)}
}

// Parse2O Parses an object for x and/or y properties and returns a new Phaser.Point with matching values.
// If the object doesn't contain those properties a Point with x/y of zero will be returned.
func (self *Point) Parse2O(obj interface{}, xProp string, yProp string) *Point{
    return &Point{self.Object.Call("parse", obj, xProp, yProp)}
}

// ParseI Parses an object for x and/or y properties and returns a new Phaser.Point with matching values.
// If the object doesn't contain those properties a Point with x/y of zero will be returned.
func (self *Point) ParseI(args ...interface{}) *Point{
    return &Point{self.Object.Call("parse", args)}
}

