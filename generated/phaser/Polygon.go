// Automatic generation for Phaser.Polygon
// generated file Polygon.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Creates a new Polygon.
// 
// The points can be set from a variety of formats:
// 
// - An array of Point objects: `[new Phaser.Point(x1, y1), ...]`
// - An array of objects with public x/y properties: `[obj1, obj2, ...]`
// - An array of paired numbers that represent point coordinates: `[x1,y1, x2,y2, ...]`
// - As separate Point arguments: `setTo(new Phaser.Point(x1, y1), ...)`
// - As separate objects with public x/y properties arguments: `setTo(obj1, obj2, ...)`
// - As separate arguments representing point coordinates: `setTo(x1,y1, x2,y2, ...)`
type Polygon struct {
    *js.Object
}


// The area of this Polygon.
func (self *Polygon) GetAreaA() int{
    return self.Object.Get("area").Int()
}

// The area of this Polygon.
func (self *Polygon) SetAreaA(member int) {
    self.Object.Set("area", member)
}

// Is the Polygon closed or not?
func (self *Polygon) GetClosedA() bool{
    return self.Object.Get("closed").Bool()
}

// Is the Polygon closed or not?
func (self *Polygon) SetClosedA(member bool) {
    self.Object.Set("closed", member)
}

// Has this Polygon been flattened by a call to `Polygon.flatten` ?
func (self *Polygon) GetFlattenedA() bool{
    return self.Object.Get("flattened").Bool()
}

// Has this Polygon been flattened by a call to `Polygon.flatten` ?
func (self *Polygon) SetFlattenedA(member bool) {
    self.Object.Set("flattened", member)
}

// The base object type.
func (self *Polygon) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The base object type.
func (self *Polygon) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Sets and modifies the points of this polygon.
// 
// See {@link Phaser.Polygon#setTo setTo} for the different kinds of arrays formats that can be assigned. The array of vertex points.
func (self *Polygon) GetPointsA() []Point{
	array00 := self.Object.Get("points")
	length00 := array00.Length()
	out00 := make([]Point, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = Point{array00.Index(i00)}
	}
	return out00
}

// Sets and modifies the points of this polygon.
// 
// See {@link Phaser.Polygon#setTo setTo} for the different kinds of arrays formats that can be assigned. The array of vertex points.
func (self *Polygon) SetPointsA(member []Point) {
    self.Object.Set("points", member)
}



// Export the points as an array of flat numbers, following the sequence [ x,y, x,y, x,y ]
func (self *Polygon) ToNumberArray(output []interface{}) []interface{}{
	array00 := self.Object.Call("toNumberArray", output)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Export the points as an array of flat numbers, following the sequence [ x,y, x,y, x,y ]
func (self *Polygon) ToNumberArrayI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("toNumberArray", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Flattens this Polygon so the points are a sequence of numbers.
// Any Point objects found are removed and replaced with two numbers.
// Also sets the Polygon.flattened property to `true`.
func (self *Polygon) Flatten() *Polygon{
    return &Polygon{self.Object.Call("flatten")}
}

// Flattens this Polygon so the points are a sequence of numbers.
// Any Point objects found are removed and replaced with two numbers.
// Also sets the Polygon.flattened property to `true`.
func (self *Polygon) FlattenI(args ...interface{}) *Polygon{
    return &Polygon{self.Object.Call("flatten", args)}
}

// Creates a copy of the given Polygon.
// This is a deep clone, the resulting copy contains new Phaser.Point objects
func (self *Polygon) Clone(output *Polygon) *Polygon{
    return &Polygon{self.Object.Call("clone", output)}
}

// Creates a copy of the given Polygon.
// This is a deep clone, the resulting copy contains new Phaser.Point objects
func (self *Polygon) CloneI(args ...interface{}) *Polygon{
    return &Polygon{self.Object.Call("clone", args)}
}

// Checks whether the x and y coordinates are contained within this polygon.
func (self *Polygon) Contains(x int, y int) bool{
    return self.Object.Call("contains", x, y).Bool()
}

// Checks whether the x and y coordinates are contained within this polygon.
func (self *Polygon) ContainsI(args ...interface{}) bool{
    return self.Object.Call("contains", args).Bool()
}

// Sets this Polygon to the given points.
// 
// The points can be set from a variety of formats:
// 
// - An array of Point objects: `[new Phaser.Point(x1, y1), ...]`
// - An array of objects with public x/y properties: `[obj1, obj2, ...]`
// - An array of paired numbers that represent point coordinates: `[x1,y1, x2,y2, ...]`
// - An array of arrays with two elements representing x/y coordinates: `[[x1, y1], [x2, y2], ...]`
// - As separate Point arguments: `setTo(new Phaser.Point(x1, y1), ...)`
// - As separate objects with public x/y properties arguments: `setTo(obj1, obj2, ...)`
// - As separate arguments representing point coordinates: `setTo(x1,y1, x2,y2, ...)`
// 
// `setTo` may also be called without any arguments to remove all points.
func (self *Polygon) SetTo(points interface{}) *Polygon{
    return &Polygon{self.Object.Call("setTo", points)}
}

// Sets this Polygon to the given points.
// 
// The points can be set from a variety of formats:
// 
// - An array of Point objects: `[new Phaser.Point(x1, y1), ...]`
// - An array of objects with public x/y properties: `[obj1, obj2, ...]`
// - An array of paired numbers that represent point coordinates: `[x1,y1, x2,y2, ...]`
// - An array of arrays with two elements representing x/y coordinates: `[[x1, y1], [x2, y2], ...]`
// - As separate Point arguments: `setTo(new Phaser.Point(x1, y1), ...)`
// - As separate objects with public x/y properties arguments: `setTo(obj1, obj2, ...)`
// - As separate arguments representing point coordinates: `setTo(x1,y1, x2,y2, ...)`
// 
// `setTo` may also be called without any arguments to remove all points.
func (self *Polygon) SetToI(args ...interface{}) *Polygon{
    return &Polygon{self.Object.Call("setTo", args)}
}

// Calcuates the area of the Polygon. This is available in the property Polygon.area
func (self *Polygon) CalculateArea(y0 int) int{
    return self.Object.Call("calculateArea", y0).Int()
}

// Calcuates the area of the Polygon. This is available in the property Polygon.area
func (self *Polygon) CalculateAreaI(args ...interface{}) int{
    return self.Object.Call("calculateArea", args).Int()
}
