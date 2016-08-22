// Automatic generation for Phaser.Line
// generated file Line.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Creates a new Line object with a start and an end point.
type Line struct {
    *js.Object
}


// The start point of the line.
func (self *Line) GetStart() *Point{
    return &Point{self.Get("start")}
}

// The start point of the line.
func (self *Line) SetStart(member *Point) {
    self.Set("start", member)
}

// The end point of the line.
func (self *Line) GetEnd() *Point{
    return &Point{self.Get("end")}
}

// The end point of the line.
func (self *Line) SetEnd(member *Point) {
    self.Set("end", member)
}

// The const type of this object.
func (self *Line) GetType() int{
    return self.Get("type").Int()
}

// The const type of this object.
func (self *Line) SetType(member int) {
    self.Set("type", member)
}

// Gets the length of the line segment.
func (self *Line) GetLength() int{
    return self.Get("length").Int()
}

// Gets the length of the line segment.
func (self *Line) SetLength(member int) {
    self.Set("length", member)
}

// Gets the angle of the line in radians.
func (self *Line) GetAngle() int{
    return self.Get("angle").Int()
}

// Gets the angle of the line in radians.
func (self *Line) SetAngle(member int) {
    self.Set("angle", member)
}

// Gets the slope of the line (y/x).
func (self *Line) GetSlope() int{
    return self.Get("slope").Int()
}

// Gets the slope of the line (y/x).
func (self *Line) SetSlope(member int) {
    self.Set("slope", member)
}

// Gets the perpendicular slope of the line (x/y).
func (self *Line) GetPerpSlope() int{
    return self.Get("perpSlope").Int()
}

// Gets the perpendicular slope of the line (x/y).
func (self *Line) SetPerpSlope(member int) {
    self.Set("perpSlope", member)
}

// Gets the x coordinate of the top left of the bounds around this line.
func (self *Line) GetX() int{
    return self.Get("x").Int()
}

// Gets the x coordinate of the top left of the bounds around this line.
func (self *Line) SetX(member int) {
    self.Set("x", member)
}

// Gets the y coordinate of the top left of the bounds around this line.
func (self *Line) GetY() int{
    return self.Get("y").Int()
}

// Gets the y coordinate of the top left of the bounds around this line.
func (self *Line) SetY(member int) {
    self.Set("y", member)
}

// Gets the left-most point of this line.
func (self *Line) GetLeft() int{
    return self.Get("left").Int()
}

// Gets the left-most point of this line.
func (self *Line) SetLeft(member int) {
    self.Set("left", member)
}

// Gets the right-most point of this line.
func (self *Line) GetRight() int{
    return self.Get("right").Int()
}

// Gets the right-most point of this line.
func (self *Line) SetRight(member int) {
    self.Set("right", member)
}

// Gets the top-most point of this line.
func (self *Line) GetTop() int{
    return self.Get("top").Int()
}

// Gets the top-most point of this line.
func (self *Line) SetTop(member int) {
    self.Set("top", member)
}

// Gets the bottom-most point of this line.
func (self *Line) GetBottom() int{
    return self.Get("bottom").Int()
}

// Gets the bottom-most point of this line.
func (self *Line) SetBottom(member int) {
    self.Set("bottom", member)
}

// Gets the width of this bounds of this line.
func (self *Line) GetWidth() int{
    return self.Get("width").Int()
}

// Gets the width of this bounds of this line.
func (self *Line) SetWidth(member int) {
    self.Set("width", member)
}

// Gets the height of this bounds of this line.
func (self *Line) GetHeight() int{
    return self.Get("height").Int()
}

// Gets the height of this bounds of this line.
func (self *Line) SetHeight(member int) {
    self.Set("height", member)
}

// Gets the x component of the left-hand normal of this line.
func (self *Line) GetNormalX() int{
    return self.Get("normalX").Int()
}

// Gets the x component of the left-hand normal of this line.
func (self *Line) SetNormalX(member int) {
    self.Set("normalX", member)
}

// Gets the y component of the left-hand normal of this line.
func (self *Line) GetNormalY() int{
    return self.Get("normalY").Int()
}

// Gets the y component of the left-hand normal of this line.
func (self *Line) SetNormalY(member int) {
    self.Set("normalY", member)
}

// Gets the angle in radians of the normal of this line (line.angle - 90 degrees.)
func (self *Line) GetNormalAngle() int{
    return self.Get("normalAngle").Int()
}

// Gets the angle in radians of the normal of this line (line.angle - 90 degrees.)
func (self *Line) SetNormalAngle(member int) {
    self.Set("normalAngle", member)
}



// Sets the components of the Line to the specified values.
func (self *Line) SetToI(args ...interface{}) *Line{
    return &Line{self.Call("setTo", args)}
}

// Sets the line to match the x/y coordinates of the two given sprites.
// Can optionally be calculated from their center coordinates.
func (self *Line) FromSpriteI(args ...interface{}) *Line{
    return &Line{self.Call("fromSprite", args)}
}

// Sets this line to start at the given `x` and `y` coordinates and for the segment to extend at `angle` for the given `length`.
func (self *Line) FromAngleI(args ...interface{}) *Line{
    return &Line{self.Call("fromAngle", args)}
}

// Rotates the line by the amount specified in `angle`.
// 
// Rotation takes place from the center of the line.
// If you wish to rotate around a different point see Line.rotateAround.
// 
// If you wish to rotate the ends of the Line then see Line.start.rotate or Line.end.rotate.
func (self *Line) RotateI(args ...interface{}) *Line{
    return &Line{self.Call("rotate", args)}
}

// Rotates the line by the amount specified in `angle`.
// 
// Rotation takes place around the coordinates given.
func (self *Line) RotateAroundI(args ...interface{}) *Line{
    return &Line{self.Call("rotateAround", args)}
}

// Checks for intersection between this line and another Line.
// If asSegment is true it will check for segment intersection. If asSegment is false it will check for line intersection.
// Returns the intersection segment of AB and EF as a Point, or null if there is no intersection.
func (self *Line) IntersectsI(args ...interface{}) *Point{
    return &Point{self.Call("intersects", args)}
}

// Returns the reflected angle between two lines.
// This is the outgoing angle based on the angle of this line and the normalAngle of the given line.
func (self *Line) ReflectI(args ...interface{}) int{
    return self.Call("reflect", args).Int()
}

// Returns a Point object where the x and y values correspond to the center (or midpoint) of the Line segment.
func (self *Line) MidPointI(args ...interface{}) *Point{
    return &Point{self.Call("midPoint", args)}
}

// Centers this Line on the given coordinates.
// 
// The line is centered by positioning the start and end points so that the lines midpoint matches
// the coordinates given.
func (self *Line) CenterOnI(args ...interface{}) *Line{
    return &Line{self.Call("centerOn", args)}
}

// Tests if the given coordinates fall on this line. See pointOnSegment to test against just the line segment.
func (self *Line) PointOnLineI(args ...interface{}) bool{
    return self.Call("pointOnLine", args).Bool()
}

// Tests if the given coordinates fall on this line and within the segment. See pointOnLine to test against just the line.
func (self *Line) PointOnSegmentI(args ...interface{}) bool{
    return self.Call("pointOnSegment", args).Bool()
}

// Picks a random point from anywhere on the Line segment and returns it.
func (self *Line) RandomI(args ...interface{}) *Point{
    return &Point{self.Call("random", args)}
}

// Using Bresenham's line algorithm this will return an array of all coordinates on this line.
// The start and end points are rounded before this runs as the algorithm works on integers.
func (self *Line) CoordinatesOnLineI(args ...interface{}) []interface{}{
	array := self.Call("coordinatesOnLine", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Returns a new Line object with the same values for the start and end properties as this Line object.
func (self *Line) CloneI(args ...interface{}) *Line{
    return &Line{self.Call("clone", args)}
}

// Checks for intersection between two lines as defined by the given start and end points.
// If asSegment is true it will check for line segment intersection. If asSegment is false it will check for line intersection.
// Returns the intersection segment of AB and EF as a Point, or null if there is no intersection.
// Adapted from code by Keith Hair
func (self *Line) IntersectsPointsI(args ...interface{}) *Point{
    return &Point{self.Call("intersectsPoints", args)}
}
