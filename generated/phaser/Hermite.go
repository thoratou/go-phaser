// Package phaser Automatic generation for Phaser.Hermite
// generated file Hermite.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Hermite A data representation of a Hermite Curve (see http://en.wikipedia.org/wiki/Cubic_Hermite_spline)
// 
// A Hermite curve has a start and end point and tangent vectors for both of them.
// The curve will always pass through the two control points and the shape of it is controlled
// by the length and direction of the tangent vectors.  At the control points the curve will
// be facing exactly in the vector direction.
// 
// As these curves change speed (speed = distance between points separated by an equal change in
// 't' value - see Hermite.getPoint) this class attempts to reduce the variation by pre-calculating
// the `accuracy` number of points on the curve. The straight-line distances to these points are stored
// in the private 'points' array, and this information is used by Hermite.findT() to convert a pixel
// distance along the curve into a 'time' value.
// 
// Higher `accuracy` values will result in more even movement, but require more memory for the points
// list. 5 works, but 10 seems to be an ideal value for the length of curves found in most games on
// a desktop screen. If you use very long curves (more than 400 pixels) you may need to increase
// this value further.
type Hermite struct {
    *js.Object
}

// NewHermite A data representation of a Hermite Curve (see http://en.wikipedia.org/wiki/Cubic_Hermite_spline)
// 
// A Hermite curve has a start and end point and tangent vectors for both of them.
// The curve will always pass through the two control points and the shape of it is controlled
// by the length and direction of the tangent vectors.  At the control points the curve will
// be facing exactly in the vector direction.
// 
// As these curves change speed (speed = distance between points separated by an equal change in
// 't' value - see Hermite.getPoint) this class attempts to reduce the variation by pre-calculating
// the `accuracy` number of points on the curve. The straight-line distances to these points are stored
// in the private 'points' array, and this information is used by Hermite.findT() to convert a pixel
// distance along the curve into a 'time' value.
// 
// Higher `accuracy` values will result in more even movement, but require more memory for the points
// list. 5 works, but 10 seems to be an ideal value for the length of curves found in most games on
// a desktop screen. If you use very long curves (more than 400 pixels) you may need to increase
// this value further.
func NewHermite(p1x int, p1y int, p2x int, p2y int, v1x int, v1y int, v2x int, v2y int) *Hermite {
    return &Hermite{js.Global.Get("Phaser").Get("Hermite").New(p1x, p1y, p2x, p2y, v1x, v1y, v2x, v2y)}
}
// NewHermite1O A data representation of a Hermite Curve (see http://en.wikipedia.org/wiki/Cubic_Hermite_spline)
// 
// A Hermite curve has a start and end point and tangent vectors for both of them.
// The curve will always pass through the two control points and the shape of it is controlled
// by the length and direction of the tangent vectors.  At the control points the curve will
// be facing exactly in the vector direction.
// 
// As these curves change speed (speed = distance between points separated by an equal change in
// 't' value - see Hermite.getPoint) this class attempts to reduce the variation by pre-calculating
// the `accuracy` number of points on the curve. The straight-line distances to these points are stored
// in the private 'points' array, and this information is used by Hermite.findT() to convert a pixel
// distance along the curve into a 'time' value.
// 
// Higher `accuracy` values will result in more even movement, but require more memory for the points
// list. 5 works, but 10 seems to be an ideal value for the length of curves found in most games on
// a desktop screen. If you use very long curves (more than 400 pixels) you may need to increase
// this value further.
func NewHermite1O(p1x int, p1y int, p2x int, p2y int, v1x int, v1y int, v2x int, v2y int, accuracy int) *Hermite {
    return &Hermite{js.Global.Get("Phaser").Get("Hermite").New(p1x, p1y, p2x, p2y, v1x, v1y, v2x, v2y, accuracy)}
}
// NewHermiteI A data representation of a Hermite Curve (see http://en.wikipedia.org/wiki/Cubic_Hermite_spline)
// 
// A Hermite curve has a start and end point and tangent vectors for both of them.
// The curve will always pass through the two control points and the shape of it is controlled
// by the length and direction of the tangent vectors.  At the control points the curve will
// be facing exactly in the vector direction.
// 
// As these curves change speed (speed = distance between points separated by an equal change in
// 't' value - see Hermite.getPoint) this class attempts to reduce the variation by pre-calculating
// the `accuracy` number of points on the curve. The straight-line distances to these points are stored
// in the private 'points' array, and this information is used by Hermite.findT() to convert a pixel
// distance along the curve into a 'time' value.
// 
// Higher `accuracy` values will result in more even movement, but require more memory for the points
// list. 5 works, but 10 seems to be an ideal value for the length of curves found in most games on
// a desktop screen. If you use very long curves (more than 400 pixels) you may need to increase
// this value further.
func NewHermiteI(args ...interface{}) *Hermite {
    return &Hermite{js.Global.Get("Phaser").Get("Hermite").New(args)}
}



// Hermite Binding conversion method to Hermite point 
func ToHermite(jsStruct interface{}) *Hermite {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Hermite{Object: object}
	}
	return nil
}



// Accuracy The amount of points to pre-calculate on the curve.
func (self *Hermite) Accuracy() int{
    return self.Object.Get("accuracy").Int()
}

// SetAccuracyA The amount of points to pre-calculate on the curve.
func (self *Hermite) SetAccuracyA(member int) {
    self.Object.Set("accuracy", member)
}

// P1x The x coordinate of the start of the curve. Setting this value will recalculate the curve.
func (self *Hermite) P1x() int{
    return self.Object.Get("p1x").Int()
}

// SetP1xA The x coordinate of the start of the curve. Setting this value will recalculate the curve.
func (self *Hermite) SetP1xA(member int) {
    self.Object.Set("p1x", member)
}

// P1y The y coordinate of the start of the curve. Setting this value will recalculate the curve.
func (self *Hermite) P1y() int{
    return self.Object.Get("p1y").Int()
}

// SetP1yA The y coordinate of the start of the curve. Setting this value will recalculate the curve.
func (self *Hermite) SetP1yA(member int) {
    self.Object.Set("p1y", member)
}

// P2x The x coordinate of the end of the curve. Setting this value will recalculate the curve.
func (self *Hermite) P2x() int{
    return self.Object.Get("p2x").Int()
}

// SetP2xA The x coordinate of the end of the curve. Setting this value will recalculate the curve.
func (self *Hermite) SetP2xA(member int) {
    self.Object.Set("p2x", member)
}

// P2y The y coordinate of the end of the curve. Setting this value will recalculate the curve.
func (self *Hermite) P2y() int{
    return self.Object.Get("p2y").Int()
}

// SetP2yA The y coordinate of the end of the curve. Setting this value will recalculate the curve.
func (self *Hermite) SetP2yA(member int) {
    self.Object.Set("p2y", member)
}

// V1x The x component of the tangent vector for the start of the curve. Setting this value will recalculate the curve.
func (self *Hermite) V1x() int{
    return self.Object.Get("v1x").Int()
}

// SetV1xA The x component of the tangent vector for the start of the curve. Setting this value will recalculate the curve.
func (self *Hermite) SetV1xA(member int) {
    self.Object.Set("v1x", member)
}

// V1y The y component of the tangent vector for the start of the curve. Setting this value will recalculate the curve.
func (self *Hermite) V1y() int{
    return self.Object.Get("v1y").Int()
}

// SetV1yA The y component of the tangent vector for the start of the curve. Setting this value will recalculate the curve.
func (self *Hermite) SetV1yA(member int) {
    self.Object.Set("v1y", member)
}

// V2x The x component of the tangent vector for the end of the curve. Setting this value will recalculate the curve.
func (self *Hermite) V2x() int{
    return self.Object.Get("v2x").Int()
}

// SetV2xA The x component of the tangent vector for the end of the curve. Setting this value will recalculate the curve.
func (self *Hermite) SetV2xA(member int) {
    self.Object.Set("v2x", member)
}

// V2y The y component of the tangent vector for the end of the curve. Setting this value will recalculate the curve.
func (self *Hermite) V2y() int{
    return self.Object.Get("v2y").Int()
}

// SetV2yA The y component of the tangent vector for the end of the curve. Setting this value will recalculate the curve.
func (self *Hermite) SetV2yA(member int) {
    self.Object.Set("v2y", member)
}


// Recalculate Performs the curve calculations.
// 
// This is called automatically if you change any of the curves public properties, such as `Hermite.p1x` or `Hermite.v2y`.
// 
// If you adjust any of the internal private values, then call this to update the points.
func (self *Hermite) Recalculate() *Hermite{
    return &Hermite{self.Object.Call("recalculate")}
}

// RecalculateI Performs the curve calculations.
// 
// This is called automatically if you change any of the curves public properties, such as `Hermite.p1x` or `Hermite.v2y`.
// 
// If you adjust any of the internal private values, then call this to update the points.
func (self *Hermite) RecalculateI(args ...interface{}) *Hermite{
    return &Hermite{self.Object.Call("recalculate", args)}
}

// CalculateEvenPoints Calculate a number of points along the curve, based on `Hermite.accuracy`, and stores them in the private `_points` array.
func (self *Hermite) CalculateEvenPoints() int{
    return self.Object.Call("calculateEvenPoints").Int()
}

// CalculateEvenPointsI Calculate a number of points along the curve, based on `Hermite.accuracy`, and stores them in the private `_points` array.
func (self *Hermite) CalculateEvenPointsI(args ...interface{}) int{
    return self.Object.Call("calculateEvenPoints", args).Int()
}

// FindT Convert a distance along this curve into a `time` value which will be between 0 and 1.
// 
// For example if this curve has a length of 100 pixels then `findT(50)` would return `0.5`.
func (self *Hermite) FindT(distance int) int{
    return self.Object.Call("findT", distance).Int()
}

// FindTI Convert a distance along this curve into a `time` value which will be between 0 and 1.
// 
// For example if this curve has a length of 100 pixels then `findT(50)` would return `0.5`.
func (self *Hermite) FindTI(args ...interface{}) int{
    return self.Object.Call("findT", args).Int()
}

// GetX Get the X component of a point on the curve based on the `t` (time) value, which must be between 0 and 1.
func (self *Hermite) GetX() int{
    return self.Object.Call("getX").Int()
}

// GetX1O Get the X component of a point on the curve based on the `t` (time) value, which must be between 0 and 1.
func (self *Hermite) GetX1O(t int) int{
    return self.Object.Call("getX", t).Int()
}

// GetXI Get the X component of a point on the curve based on the `t` (time) value, which must be between 0 and 1.
func (self *Hermite) GetXI(args ...interface{}) int{
    return self.Object.Call("getX", args).Int()
}

// GetY Get the Y component of a point on the curve based on the `t` (time) value, which must be between 0 and 1.
func (self *Hermite) GetY() int{
    return self.Object.Call("getY").Int()
}

// GetY1O Get the Y component of a point on the curve based on the `t` (time) value, which must be between 0 and 1.
func (self *Hermite) GetY1O(t int) int{
    return self.Object.Call("getY", t).Int()
}

// GetYI Get the Y component of a point on the curve based on the `t` (time) value, which must be between 0 and 1.
func (self *Hermite) GetYI(args ...interface{}) int{
    return self.Object.Call("getY", args).Int()
}

// GetPoint Get a point on the curve using the `t` (time) value, which must be between 0 and 1.
func (self *Hermite) GetPoint() *Point{
    return &Point{self.Object.Call("getPoint")}
}

// GetPoint1O Get a point on the curve using the `t` (time) value, which must be between 0 and 1.
func (self *Hermite) GetPoint1O(t int) *Point{
    return &Point{self.Object.Call("getPoint", t)}
}

// GetPoint2O Get a point on the curve using the `t` (time) value, which must be between 0 and 1.
func (self *Hermite) GetPoint2O(t int, point interface{}) *Point{
    return &Point{self.Object.Call("getPoint", t, point)}
}

// GetPointI Get a point on the curve using the `t` (time) value, which must be between 0 and 1.
func (self *Hermite) GetPointI(args ...interface{}) *Point{
    return &Point{self.Object.Call("getPoint", args)}
}

// GetPointWithDistance Get a point on the curve using the distance, in pixels, along the curve.
func (self *Hermite) GetPointWithDistance() *Point{
    return &Point{self.Object.Call("getPointWithDistance")}
}

// GetPointWithDistance1O Get a point on the curve using the distance, in pixels, along the curve.
func (self *Hermite) GetPointWithDistance1O(distance int) *Point{
    return &Point{self.Object.Call("getPointWithDistance", distance)}
}

// GetPointWithDistance2O Get a point on the curve using the distance, in pixels, along the curve.
func (self *Hermite) GetPointWithDistance2O(distance int, point interface{}) *Point{
    return &Point{self.Object.Call("getPointWithDistance", distance, point)}
}

// GetPointWithDistanceI Get a point on the curve using the distance, in pixels, along the curve.
func (self *Hermite) GetPointWithDistanceI(args ...interface{}) *Point{
    return &Point{self.Object.Call("getPointWithDistance", args)}
}

// GetAngle Calculate and return the angle, in radians, of the curves tangent based on time.
func (self *Hermite) GetAngle() int{
    return self.Object.Call("getAngle").Int()
}

// GetAngle1O Calculate and return the angle, in radians, of the curves tangent based on time.
func (self *Hermite) GetAngle1O(t int) int{
    return self.Object.Call("getAngle", t).Int()
}

// GetAngleI Calculate and return the angle, in radians, of the curves tangent based on time.
func (self *Hermite) GetAngleI(args ...interface{}) int{
    return self.Object.Call("getAngle", args).Int()
}

// GetAngleWithDistance Calculate and return the angle, in radians, of the curves tangent at the given pixel distance along the curves length.
func (self *Hermite) GetAngleWithDistance() int{
    return self.Object.Call("getAngleWithDistance").Int()
}

// GetAngleWithDistance1O Calculate and return the angle, in radians, of the curves tangent at the given pixel distance along the curves length.
func (self *Hermite) GetAngleWithDistance1O(distance int) int{
    return self.Object.Call("getAngleWithDistance", distance).Int()
}

// GetAngleWithDistanceI Calculate and return the angle, in radians, of the curves tangent at the given pixel distance along the curves length.
func (self *Hermite) GetAngleWithDistanceI(args ...interface{}) int{
    return self.Object.Call("getAngleWithDistance", args).Int()
}

// GetEntryTangent Get the angle of the curves entry point.
func (self *Hermite) GetEntryTangent(point interface{}) *Point{
    return &Point{self.Object.Call("getEntryTangent", point)}
}

// GetEntryTangentI Get the angle of the curves entry point.
func (self *Hermite) GetEntryTangentI(args ...interface{}) *Point{
    return &Point{self.Object.Call("getEntryTangent", args)}
}

