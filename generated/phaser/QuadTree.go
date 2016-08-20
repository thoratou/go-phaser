// Automatic generation for Phaser.QuadTree
// generated file QuadTree.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A QuadTree implementation. The original code was a conversion of the Java code posted to GameDevTuts.
// However I've tweaked it massively to add node indexing, removed lots of temp. var creation and significantly increased performance as a result.
// Original version at https://github.com/timohausmann/quadtree-js/
type QuadTree struct {
    *js.Object
}


// The maximum number of objects per node.
func (self *QuadTree) GetMaxObjects() float64{
    return self.Get("maxObjects").Float()
}

// The maximum number of objects per node.
func (self *QuadTree) SetMaxObjects(member float64) {
    self.Set("maxObjects", member)
}

// The maximum number of levels to break down to.
func (self *QuadTree) GetMaxLevels() float64{
    return self.Get("maxLevels").Float()
}

// The maximum number of levels to break down to.
func (self *QuadTree) SetMaxLevels(member float64) {
    self.Set("maxLevels", member)
}

// The current level.
func (self *QuadTree) GetLevel() float64{
    return self.Get("level").Float()
}

// The current level.
func (self *QuadTree) SetLevel(member float64) {
    self.Set("level", member)
}

// Object that contains the quadtree bounds.
func (self *QuadTree) GetBounds() interface{}{
    return self.Get("bounds")
}

// Object that contains the quadtree bounds.
func (self *QuadTree) SetBounds(member interface{}) {
    self.Set("bounds", member)
}

// Array of quadtree children.
func (self *QuadTree) GetObjects() []interface{}{
	array := self.Get("objects")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Array of quadtree children.
func (self *QuadTree) SetObjects(member []interface{}) {
    self.Set("objects", member)
}

// Array of associated child nodes.
func (self *QuadTree) GetNodes() []interface{}{
	array := self.Get("nodes")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Array of associated child nodes.
func (self *QuadTree) SetNodes(member []interface{}) {
    self.Set("nodes", member)
}



// Resets the QuadTree.
func (self *QuadTree) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Populates this quadtree with the children of the given Group. In order to be added the child must exist and have a body property.
func (self *QuadTree) PopulateI(args ...interface{}) {
    self.Call("populate", args)
}

// Handler for the populate method.
func (self *QuadTree) PopulateHandlerI(args ...interface{}) {
    self.Call("populateHandler", args)
}

// Split the node into 4 subnodes
func (self *QuadTree) SplitI(args ...interface{}) {
    self.Call("split", args)
}

// Insert the object into the node. If the node exceeds the capacity, it will split and add all objects to their corresponding subnodes.
func (self *QuadTree) InsertI(args ...interface{}) {
    self.Call("insert", args)
}

// Determine which node the object belongs to.
func (self *QuadTree) GetIndexI(args ...interface{}) float64{
    return self.Call("getIndex", args).Float()
}

// Return all objects that could collide with the given Sprite or Rectangle.
func (self *QuadTree) RetrieveI(args ...interface{}) []interface{}{
	array := self.Call("retrieve", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Clear the quadtree.
func (self *QuadTree) ClearI(args ...interface{}) {
    self.Call("clear", args)
}
