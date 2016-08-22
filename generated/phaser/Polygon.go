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
func (self *Polygon) GetArea() int{
    return self.Get("area").Int()
}

// The area of this Polygon.
func (self *Polygon) SetArea(member int) {
    self.Set("area", member)
}

// Is the Polygon closed or not?
func (self *Polygon) GetClosed() bool{
    return self.Get("closed").Bool()
}

// Is the Polygon closed or not?
func (self *Polygon) SetClosed(member bool) {
    self.Set("closed", member)
}

// Has this Polygon been flattened by a call to `Polygon.flatten` ?
func (self *Polygon) GetFlattened() bool{
    return self.Get("flattened").Bool()
}

// Has this Polygon been flattened by a call to `Polygon.flatten` ?
func (self *Polygon) SetFlattened(member bool) {
    self.Set("flattened", member)
}

// The base object type.
func (self *Polygon) GetType() int{
    return self.Get("type").Int()
}

// The base object type.
func (self *Polygon) SetType(member int) {
    self.Set("type", member)
}

// Sets and modifies the points of this polygon.
// 
// See {@link Phaser.Polygon#setTo setTo} for the different kinds of arrays formats that can be assigned. The array of vertex points.
func (self *Polygon) GetPoints() []Point{
	array := self.Get("points")
	length := array.Length()
	out := make([]Point, length, length)
	for i := 0; i < length; i++ {
		out[i] = Point{array.Index(i)}
	}
	return out
}

// Sets and modifies the points of this polygon.
// 
// See {@link Phaser.Polygon#setTo setTo} for the different kinds of arrays formats that can be assigned. The array of vertex points.
func (self *Polygon) SetPoints(member []Point) {
    self.Set("points", member)
}



// Export the points as an array of flat numbers, following the sequence [ x,y, x,y, x,y ]
func (self *Polygon) ToNumberArrayI(args ...interface{}) []interface{}{
	array := self.Call("toNumberArray", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Flattens this Polygon so the points are a sequence of numbers.
// Any Point objects found are removed and replaced with two numbers.
// Also sets the Polygon.flattened property to `true`.
func (self *Polygon) FlattenI(args ...interface{}) *Polygon{
    return &Polygon{self.Call("flatten", args)}
}

// Creates a copy of the given Polygon.
// This is a deep clone, the resulting copy contains new Phaser.Point objects
func (self *Polygon) CloneI(args ...interface{}) *Polygon{
    return &Polygon{self.Call("clone", args)}
}

// Checks whether the x and y coordinates are contained within this polygon.
func (self *Polygon) ContainsI(args ...interface{}) bool{
    return self.Call("contains", args).Bool()
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
    return &Polygon{self.Call("setTo", args)}
}

// Calcuates the area of the Polygon. This is available in the property Polygon.area
func (self *Polygon) CalculateAreaI(args ...interface{}) int{
    return self.Call("calculateArea", args).Int()
}
