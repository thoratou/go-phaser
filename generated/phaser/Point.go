// Automatic generation for Phaser.Point
// generated file Point.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A Point object represents a location in a two-dimensional coordinate system, where x represents the horizontal axis and y represents the vertical axis.
// The following code creates a point at (0,0):
// `var myPoint = new Phaser.Point();`
// You can also use them as 2D Vectors and you'll find different vector related methods in this class.
type Point struct {
    *js.Object
}


// The x value of the point.
func (self *Point) GetX() float64{
    return self.Get("x").Float()
}

// The x value of the point.
func (self *Point) SetX(member float64) {
    self.Set("x", member)
}

// The y value of the point.
func (self *Point) GetY() float64{
    return self.Get("y").Float()
}

// The y value of the point.
func (self *Point) SetY(member float64) {
    self.Set("y", member)
}

// The const type of this object.
func (self *Point) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *Point) SetType(member float64) {
    self.Set("type", member)
}



// Copies the x and y properties from any given object to this Point.
func (self *Point) CopyFromI(args ...interface{}) *Point{
    return &Point{self.Call("copyFrom", args)}
}

// Inverts the x and y values of this Point
func (self *Point) InvertI(args ...interface{}) *Point{
    return &Point{self.Call("invert", args)}
}

// Sets the `x` and `y` values of this Point object to the given values.
// If you omit the `y` value then the `x` value will be applied to both, for example:
// `Point.setTo(2)` is the same as `Point.setTo(2, 2)`
func (self *Point) SetToI(args ...interface{}) *Point{
    return &Point{self.Call("setTo", args)}
}

// Sets the `x` and `y` values of this Point object to the given values.
// If you omit the `y` value then the `x` value will be applied to both, for example:
// `Point.set(2)` is the same as `Point.set(2, 2)`
func (self *Point) SetI(args ...interface{}) *Point{
    return &Point{self.Call("set", args)}
}

// Adds the given x and y values to this Point.
func (self *Point) AddI(args ...interface{}) *Point{
    return &Point{self.Call("add", args)}
}

// Subtracts the given x and y values from this Point.
func (self *Point) SubtractI(args ...interface{}) *Point{
    return &Point{self.Call("subtract", args)}
}

// Multiplies Point.x and Point.y by the given x and y values. Sometimes known as `Scale`.
func (self *Point) MultiplyI(args ...interface{}) *Point{
    return &Point{self.Call("multiply", args)}
}

// Divides Point.x and Point.y by the given x and y values.
func (self *Point) DivideI(args ...interface{}) *Point{
    return &Point{self.Call("divide", args)}
}

// Clamps the x value of this Point to be between the given min and max.
func (self *Point) ClampXI(args ...interface{}) *Point{
    return &Point{self.Call("clampX", args)}
}

// Clamps the y value of this Point to be between the given min and max
func (self *Point) ClampYI(args ...interface{}) *Point{
    return &Point{self.Call("clampY", args)}
}

// Clamps this Point object values to be between the given min and max.
func (self *Point) ClampI(args ...interface{}) *Point{
    return &Point{self.Call("clamp", args)}
}

// Creates a copy of the given Point.
func (self *Point) CloneI(args ...interface{}) *Point{
    return &Point{self.Call("clone", args)}
}

// Copies the x and y properties from this Point to any given object.
func (self *Point) CopyToI(args ...interface{}) interface{}{
    return self.Call("copyTo", args)
}

// Returns the distance of this Point object to the given object (can be a Circle, Point or anything with x/y properties)
func (self *Point) DistanceI(args ...interface{}) float64{
    return self.Call("distance", args).Float()
}

// Determines whether the given objects x/y values are equal to this Point object.
func (self *Point) EqualsI(args ...interface{}) bool{
    return self.Call("equals", args).Bool()
}

// Returns the angle between this Point object and another object with public x and y properties.
func (self *Point) AngleI(args ...interface{}) float64{
    return self.Call("angle", args).Float()
}

// Rotates this Point around the x/y coordinates given to the desired angle.
func (self *Point) RotateI(args ...interface{}) *Point{
    return &Point{self.Call("rotate", args)}
}

// Calculates the length of the Point object.
func (self *Point) GetMagnitudeI(args ...interface{}) float64{
    return self.Call("getMagnitude", args).Float()
}

// Calculates the length squared of the Point object.
func (self *Point) GetMagnitudeSqI(args ...interface{}) float64{
    return self.Call("getMagnitudeSq", args).Float()
}

// Alters the length of the Point without changing the direction.
func (self *Point) SetMagnitudeI(args ...interface{}) *Point{
    return &Point{self.Call("setMagnitude", args)}
}

// Alters the Point object so that its length is 1, but it retains the same direction.
func (self *Point) NormalizeI(args ...interface{}) *Point{
    return &Point{self.Call("normalize", args)}
}

// Determine if this point is at 0,0.
func (self *Point) IsZeroI(args ...interface{}) bool{
    return self.Call("isZero", args).Bool()
}

// The dot product of this and another Point object.
func (self *Point) DotI(args ...interface{}) float64{
    return self.Call("dot", args).Float()
}

// The cross product of this and another Point object.
func (self *Point) CrossI(args ...interface{}) float64{
    return self.Call("cross", args).Float()
}

// Make this Point perpendicular (90 degrees rotation)
func (self *Point) PerpI(args ...interface{}) *Point{
    return &Point{self.Call("perp", args)}
}

// Make this Point perpendicular (-90 degrees rotation)
func (self *Point) RperpI(args ...interface{}) *Point{
    return &Point{self.Call("rperp", args)}
}

// Right-hand normalize (make unit length) this Point.
func (self *Point) NormalRightHandI(args ...interface{}) *Point{
    return &Point{self.Call("normalRightHand", args)}
}

// Math.floor() both the x and y properties of this Point.
func (self *Point) FloorI(args ...interface{}) *Point{
    return &Point{self.Call("floor", args)}
}

// Math.ceil() both the x and y properties of this Point.
func (self *Point) CeilI(args ...interface{}) *Point{
    return &Point{self.Call("ceil", args)}
}

// Returns a string representation of this object.
func (self *Point) ToStringI(args ...interface{}) string{
    return self.Call("toString", args).String()
}

// Creates a negative Point.
func (self *Point) NegativeI(args ...interface{}) *Point{
    return &Point{self.Call("negative", args)}
}

// Adds two 2D Points together and multiplies the result by the given scalar.
func (self *Point) MultiplyAddI(args ...interface{}) *Point{
    return &Point{self.Call("multiplyAdd", args)}
}

// Interpolates the two given Points, based on the `f` value (between 0 and 1) and returns a new Point.
func (self *Point) InterpolateI(args ...interface{}) *Point{
    return &Point{self.Call("interpolate", args)}
}

// Project two Points onto another Point.
func (self *Point) ProjectI(args ...interface{}) *Point{
    return &Point{self.Call("project", args)}
}

// Project two Points onto a Point of unit length.
func (self *Point) ProjectUnitI(args ...interface{}) *Point{
    return &Point{self.Call("projectUnit", args)}
}

// Calculates centroid (or midpoint) from an array of points. If only one point is provided, that point is returned.
func (self *Point) CentroidI(args ...interface{}) *Point{
    return &Point{self.Call("centroid", args)}
}

// Parses an object for x and/or y properties and returns a new Phaser.Point with matching values.
// If the object doesn't contain those properties a Point with x/y of zero will be returned.
func (self *Point) ParseI(args ...interface{}) *Point{
    return &Point{self.Call("parse", args)}
}
