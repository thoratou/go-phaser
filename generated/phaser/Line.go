// Package phaser Automatic generation for Phaser.Line
// generated file Line.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Line Creates a new Line object with a start and an end point.
type Line struct {
    *js.Object
}

// NewLine Creates a new Line object with a start and an end point.
func NewLine() *Line {
    return &Line{js.Global.Get("Phaser").Get("Line").New()}
}
// NewLine1O Creates a new Line object with a start and an end point.
func NewLine1O(x1 int) *Line {
    return &Line{js.Global.Get("Phaser").Get("Line").New(x1)}
}
// NewLine2O Creates a new Line object with a start and an end point.
func NewLine2O(x1 int, y1 int) *Line {
    return &Line{js.Global.Get("Phaser").Get("Line").New(x1, y1)}
}
// NewLine3O Creates a new Line object with a start and an end point.
func NewLine3O(x1 int, y1 int, x2 int) *Line {
    return &Line{js.Global.Get("Phaser").Get("Line").New(x1, y1, x2)}
}
// NewLine4O Creates a new Line object with a start and an end point.
func NewLine4O(x1 int, y1 int, x2 int, y2 int) *Line {
    return &Line{js.Global.Get("Phaser").Get("Line").New(x1, y1, x2, y2)}
}
// NewLineI Creates a new Line object with a start and an end point.
func NewLineI(args ...interface{}) *Line {
    return &Line{js.Global.Get("Phaser").Get("Line").New(args)}
}



// Start The start point of the line.
func (self *Line) Start() *Point{
    return &Point{self.Object.Get("start")}
}

// SetStartA The start point of the line.
func (self *Line) SetStartA(member *Point) {
    self.Object.Set("start", member)
}

// End The end point of the line.
func (self *Line) End() *Point{
    return &Point{self.Object.Get("end")}
}

// SetEndA The end point of the line.
func (self *Line) SetEndA(member *Point) {
    self.Object.Set("end", member)
}

// Type The const type of this object.
func (self *Line) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *Line) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Length Gets the length of the line segment.
func (self *Line) Length() int{
    return self.Object.Get("length").Int()
}

// SetLengthA Gets the length of the line segment.
func (self *Line) SetLengthA(member int) {
    self.Object.Set("length", member)
}

// Angle Gets the angle of the line in radians.
func (self *Line) Angle() int{
    return self.Object.Get("angle").Int()
}

// SetAngleA Gets the angle of the line in radians.
func (self *Line) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// Slope Gets the slope of the line (y/x).
func (self *Line) Slope() int{
    return self.Object.Get("slope").Int()
}

// SetSlopeA Gets the slope of the line (y/x).
func (self *Line) SetSlopeA(member int) {
    self.Object.Set("slope", member)
}

// PerpSlope Gets the perpendicular slope of the line (x/y).
func (self *Line) PerpSlope() int{
    return self.Object.Get("perpSlope").Int()
}

// SetPerpSlopeA Gets the perpendicular slope of the line (x/y).
func (self *Line) SetPerpSlopeA(member int) {
    self.Object.Set("perpSlope", member)
}

// X Gets the x coordinate of the top left of the bounds around this line.
func (self *Line) X() int{
    return self.Object.Get("x").Int()
}

// SetXA Gets the x coordinate of the top left of the bounds around this line.
func (self *Line) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y Gets the y coordinate of the top left of the bounds around this line.
func (self *Line) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA Gets the y coordinate of the top left of the bounds around this line.
func (self *Line) SetYA(member int) {
    self.Object.Set("y", member)
}

// Left Gets the left-most point of this line.
func (self *Line) Left() int{
    return self.Object.Get("left").Int()
}

// SetLeftA Gets the left-most point of this line.
func (self *Line) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// Right Gets the right-most point of this line.
func (self *Line) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA Gets the right-most point of this line.
func (self *Line) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Top Gets the top-most point of this line.
func (self *Line) Top() int{
    return self.Object.Get("top").Int()
}

// SetTopA Gets the top-most point of this line.
func (self *Line) SetTopA(member int) {
    self.Object.Set("top", member)
}

// Bottom Gets the bottom-most point of this line.
func (self *Line) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA Gets the bottom-most point of this line.
func (self *Line) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// Width Gets the width of this bounds of this line.
func (self *Line) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA Gets the width of this bounds of this line.
func (self *Line) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height Gets the height of this bounds of this line.
func (self *Line) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA Gets the height of this bounds of this line.
func (self *Line) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// NormalX Gets the x component of the left-hand normal of this line.
func (self *Line) NormalX() int{
    return self.Object.Get("normalX").Int()
}

// SetNormalXA Gets the x component of the left-hand normal of this line.
func (self *Line) SetNormalXA(member int) {
    self.Object.Set("normalX", member)
}

// NormalY Gets the y component of the left-hand normal of this line.
func (self *Line) NormalY() int{
    return self.Object.Get("normalY").Int()
}

// SetNormalYA Gets the y component of the left-hand normal of this line.
func (self *Line) SetNormalYA(member int) {
    self.Object.Set("normalY", member)
}

// NormalAngle Gets the angle in radians of the normal of this line (line.angle - 90 degrees.)
func (self *Line) NormalAngle() int{
    return self.Object.Get("normalAngle").Int()
}

// SetNormalAngleA Gets the angle in radians of the normal of this line (line.angle - 90 degrees.)
func (self *Line) SetNormalAngleA(member int) {
    self.Object.Set("normalAngle", member)
}


// SetTo Sets the components of the Line to the specified values.
func (self *Line) SetTo() *Line{
    return &Line{self.Object.Call("setTo")}
}

// SetTo1O Sets the components of the Line to the specified values.
func (self *Line) SetTo1O(x1 int) *Line{
    return &Line{self.Object.Call("setTo", x1)}
}

// SetTo2O Sets the components of the Line to the specified values.
func (self *Line) SetTo2O(x1 int, y1 int) *Line{
    return &Line{self.Object.Call("setTo", x1, y1)}
}

// SetTo3O Sets the components of the Line to the specified values.
func (self *Line) SetTo3O(x1 int, y1 int, x2 int) *Line{
    return &Line{self.Object.Call("setTo", x1, y1, x2)}
}

// SetTo4O Sets the components of the Line to the specified values.
func (self *Line) SetTo4O(x1 int, y1 int, x2 int, y2 int) *Line{
    return &Line{self.Object.Call("setTo", x1, y1, x2, y2)}
}

// SetToI Sets the components of the Line to the specified values.
func (self *Line) SetToI(args ...interface{}) *Line{
    return &Line{self.Object.Call("setTo", args)}
}

// FromSprite Sets the line to match the x/y coordinates of the two given sprites.
// Can optionally be calculated from their center coordinates.
func (self *Line) FromSprite(startSprite *Sprite, endSprite *Sprite) *Line{
    return &Line{self.Object.Call("fromSprite", startSprite, endSprite)}
}

// FromSprite1O Sets the line to match the x/y coordinates of the two given sprites.
// Can optionally be calculated from their center coordinates.
func (self *Line) FromSprite1O(startSprite *Sprite, endSprite *Sprite, useCenter bool) *Line{
    return &Line{self.Object.Call("fromSprite", startSprite, endSprite, useCenter)}
}

// FromSpriteI Sets the line to match the x/y coordinates of the two given sprites.
// Can optionally be calculated from their center coordinates.
func (self *Line) FromSpriteI(args ...interface{}) *Line{
    return &Line{self.Object.Call("fromSprite", args)}
}

// FromAngle Sets this line to start at the given `x` and `y` coordinates and for the segment to extend at `angle` for the given `length`.
func (self *Line) FromAngle(x int, y int, angle int, length int) *Line{
    return &Line{self.Object.Call("fromAngle", x, y, angle, length)}
}

// FromAngleI Sets this line to start at the given `x` and `y` coordinates and for the segment to extend at `angle` for the given `length`.
func (self *Line) FromAngleI(args ...interface{}) *Line{
    return &Line{self.Object.Call("fromAngle", args)}
}

// Rotate Rotates the line by the amount specified in `angle`.
// 
// Rotation takes place from the center of the line.
// If you wish to rotate around a different point see Line.rotateAround.
// 
// If you wish to rotate the ends of the Line then see Line.start.rotate or Line.end.rotate.
func (self *Line) Rotate(angle int) *Line{
    return &Line{self.Object.Call("rotate", angle)}
}

// Rotate1O Rotates the line by the amount specified in `angle`.
// 
// Rotation takes place from the center of the line.
// If you wish to rotate around a different point see Line.rotateAround.
// 
// If you wish to rotate the ends of the Line then see Line.start.rotate or Line.end.rotate.
func (self *Line) Rotate1O(angle int, asDegrees bool) *Line{
    return &Line{self.Object.Call("rotate", angle, asDegrees)}
}

// RotateI Rotates the line by the amount specified in `angle`.
// 
// Rotation takes place from the center of the line.
// If you wish to rotate around a different point see Line.rotateAround.
// 
// If you wish to rotate the ends of the Line then see Line.start.rotate or Line.end.rotate.
func (self *Line) RotateI(args ...interface{}) *Line{
    return &Line{self.Object.Call("rotate", args)}
}

// RotateAround Rotates the line by the amount specified in `angle`.
// 
// Rotation takes place around the coordinates given.
func (self *Line) RotateAround(x int, y int, angle int) *Line{
    return &Line{self.Object.Call("rotateAround", x, y, angle)}
}

// RotateAround1O Rotates the line by the amount specified in `angle`.
// 
// Rotation takes place around the coordinates given.
func (self *Line) RotateAround1O(x int, y int, angle int, asDegrees bool) *Line{
    return &Line{self.Object.Call("rotateAround", x, y, angle, asDegrees)}
}

// RotateAroundI Rotates the line by the amount specified in `angle`.
// 
// Rotation takes place around the coordinates given.
func (self *Line) RotateAroundI(args ...interface{}) *Line{
    return &Line{self.Object.Call("rotateAround", args)}
}

// Intersects Checks for intersection between this line and another Line.
// If asSegment is true it will check for segment intersection. If asSegment is false it will check for line intersection.
// Returns the intersection segment of AB and EF as a Point, or null if there is no intersection.
func (self *Line) Intersects(line *Line) *Point{
    return &Point{self.Object.Call("intersects", line)}
}

// Intersects1O Checks for intersection between this line and another Line.
// If asSegment is true it will check for segment intersection. If asSegment is false it will check for line intersection.
// Returns the intersection segment of AB and EF as a Point, or null if there is no intersection.
func (self *Line) Intersects1O(line *Line, asSegment bool) *Point{
    return &Point{self.Object.Call("intersects", line, asSegment)}
}

// Intersects2O Checks for intersection between this line and another Line.
// If asSegment is true it will check for segment intersection. If asSegment is false it will check for line intersection.
// Returns the intersection segment of AB and EF as a Point, or null if there is no intersection.
func (self *Line) Intersects2O(line *Line, asSegment bool, result *Point) *Point{
    return &Point{self.Object.Call("intersects", line, asSegment, result)}
}

// IntersectsI Checks for intersection between this line and another Line.
// If asSegment is true it will check for segment intersection. If asSegment is false it will check for line intersection.
// Returns the intersection segment of AB and EF as a Point, or null if there is no intersection.
func (self *Line) IntersectsI(args ...interface{}) *Point{
    return &Point{self.Object.Call("intersects", args)}
}

// Reflect Returns the reflected angle between two lines.
// This is the outgoing angle based on the angle of this line and the normalAngle of the given line.
func (self *Line) Reflect(line *Line) int{
    return self.Object.Call("reflect", line).Int()
}

// ReflectI Returns the reflected angle between two lines.
// This is the outgoing angle based on the angle of this line and the normalAngle of the given line.
func (self *Line) ReflectI(args ...interface{}) int{
    return self.Object.Call("reflect", args).Int()
}

// MidPoint Returns a Point object where the x and y values correspond to the center (or midpoint) of the Line segment.
func (self *Line) MidPoint() *Point{
    return &Point{self.Object.Call("midPoint")}
}

// MidPoint1O Returns a Point object where the x and y values correspond to the center (or midpoint) of the Line segment.
func (self *Line) MidPoint1O(out *Point) *Point{
    return &Point{self.Object.Call("midPoint", out)}
}

// MidPointI Returns a Point object where the x and y values correspond to the center (or midpoint) of the Line segment.
func (self *Line) MidPointI(args ...interface{}) *Point{
    return &Point{self.Object.Call("midPoint", args)}
}

// CenterOn Centers this Line on the given coordinates.
// 
// The line is centered by positioning the start and end points so that the lines midpoint matches
// the coordinates given.
func (self *Line) CenterOn(x int, y int) *Line{
    return &Line{self.Object.Call("centerOn", x, y)}
}

// CenterOnI Centers this Line on the given coordinates.
// 
// The line is centered by positioning the start and end points so that the lines midpoint matches
// the coordinates given.
func (self *Line) CenterOnI(args ...interface{}) *Line{
    return &Line{self.Object.Call("centerOn", args)}
}

// PointOnLine Tests if the given coordinates fall on this line. See pointOnSegment to test against just the line segment.
func (self *Line) PointOnLine(x int, y int) bool{
    return self.Object.Call("pointOnLine", x, y).Bool()
}

// PointOnLineI Tests if the given coordinates fall on this line. See pointOnSegment to test against just the line segment.
func (self *Line) PointOnLineI(args ...interface{}) bool{
    return self.Object.Call("pointOnLine", args).Bool()
}

// PointOnSegment Tests if the given coordinates fall on this line and within the segment. See pointOnLine to test against just the line.
func (self *Line) PointOnSegment(x int, y int) bool{
    return self.Object.Call("pointOnSegment", x, y).Bool()
}

// PointOnSegmentI Tests if the given coordinates fall on this line and within the segment. See pointOnLine to test against just the line.
func (self *Line) PointOnSegmentI(args ...interface{}) bool{
    return self.Object.Call("pointOnSegment", args).Bool()
}

// Random Picks a random point from anywhere on the Line segment and returns it.
func (self *Line) Random() *Point{
    return &Point{self.Object.Call("random")}
}

// Random1O Picks a random point from anywhere on the Line segment and returns it.
func (self *Line) Random1O(out interface{}) *Point{
    return &Point{self.Object.Call("random", out)}
}

// RandomI Picks a random point from anywhere on the Line segment and returns it.
func (self *Line) RandomI(args ...interface{}) *Point{
    return &Point{self.Object.Call("random", args)}
}

// CoordinatesOnLine Using Bresenham's line algorithm this will return an array of all coordinates on this line.
// The start and end points are rounded before this runs as the algorithm works on integers.
func (self *Line) CoordinatesOnLine() []interface{}{
	array00 := self.Object.Call("coordinatesOnLine")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// CoordinatesOnLine1O Using Bresenham's line algorithm this will return an array of all coordinates on this line.
// The start and end points are rounded before this runs as the algorithm works on integers.
func (self *Line) CoordinatesOnLine1O(stepRate int) []interface{}{
	array00 := self.Object.Call("coordinatesOnLine", stepRate)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// CoordinatesOnLine2O Using Bresenham's line algorithm this will return an array of all coordinates on this line.
// The start and end points are rounded before this runs as the algorithm works on integers.
func (self *Line) CoordinatesOnLine2O(stepRate int, results []interface{}) []interface{}{
	array00 := self.Object.Call("coordinatesOnLine", stepRate, results)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// CoordinatesOnLineI Using Bresenham's line algorithm this will return an array of all coordinates on this line.
// The start and end points are rounded before this runs as the algorithm works on integers.
func (self *Line) CoordinatesOnLineI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("coordinatesOnLine", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Clone Returns a new Line object with the same values for the start and end properties as this Line object.
func (self *Line) Clone(output *Line) *Line{
    return &Line{self.Object.Call("clone", output)}
}

// CloneI Returns a new Line object with the same values for the start and end properties as this Line object.
func (self *Line) CloneI(args ...interface{}) *Line{
    return &Line{self.Object.Call("clone", args)}
}

// IntersectsPoints Checks for intersection between two lines as defined by the given start and end points.
// If asSegment is true it will check for line segment intersection. If asSegment is false it will check for line intersection.
// Returns the intersection segment of AB and EF as a Point, or null if there is no intersection.
// Adapted from code by Keith Hair
func (self *Line) IntersectsPoints(a *Point, b *Point, e *Point, f *Point) *Point{
    return &Point{self.Object.Call("intersectsPoints", a, b, e, f)}
}

// IntersectsPoints1O Checks for intersection between two lines as defined by the given start and end points.
// If asSegment is true it will check for line segment intersection. If asSegment is false it will check for line intersection.
// Returns the intersection segment of AB and EF as a Point, or null if there is no intersection.
// Adapted from code by Keith Hair
func (self *Line) IntersectsPoints1O(a *Point, b *Point, e *Point, f *Point, asSegment bool) *Point{
    return &Point{self.Object.Call("intersectsPoints", a, b, e, f, asSegment)}
}

// IntersectsPoints2O Checks for intersection between two lines as defined by the given start and end points.
// If asSegment is true it will check for line segment intersection. If asSegment is false it will check for line intersection.
// Returns the intersection segment of AB and EF as a Point, or null if there is no intersection.
// Adapted from code by Keith Hair
func (self *Line) IntersectsPoints2O(a *Point, b *Point, e *Point, f *Point, asSegment bool, result interface{}) *Point{
    return &Point{self.Object.Call("intersectsPoints", a, b, e, f, asSegment, result)}
}

// IntersectsPointsI Checks for intersection between two lines as defined by the given start and end points.
// If asSegment is true it will check for line segment intersection. If asSegment is false it will check for line intersection.
// Returns the intersection segment of AB and EF as a Point, or null if there is no intersection.
// Adapted from code by Keith Hair
func (self *Line) IntersectsPointsI(args ...interface{}) *Point{
    return &Point{self.Object.Call("intersectsPoints", args)}
}

