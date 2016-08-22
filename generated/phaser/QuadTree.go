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
func (self *QuadTree) GetMaxObjectsA() int{
    return self.Object.Get("maxObjects").Int()
}

// The maximum number of objects per node.
func (self *QuadTree) SetMaxObjectsA(member int) {
    self.Object.Set("maxObjects", member)
}

// The maximum number of levels to break down to.
func (self *QuadTree) GetMaxLevelsA() int{
    return self.Object.Get("maxLevels").Int()
}

// The maximum number of levels to break down to.
func (self *QuadTree) SetMaxLevelsA(member int) {
    self.Object.Set("maxLevels", member)
}

// The current level.
func (self *QuadTree) GetLevelA() int{
    return self.Object.Get("level").Int()
}

// The current level.
func (self *QuadTree) SetLevelA(member int) {
    self.Object.Set("level", member)
}

// Object that contains the quadtree bounds.
func (self *QuadTree) GetBoundsA() interface{}{
    return self.Object.Get("bounds")
}

// Object that contains the quadtree bounds.
func (self *QuadTree) SetBoundsA(member interface{}) {
    self.Object.Set("bounds", member)
}

// Array of quadtree children.
func (self *QuadTree) GetObjectsA() []interface{}{
	array00 := self.Object.Get("objects")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Array of quadtree children.
func (self *QuadTree) SetObjectsA(member []interface{}) {
    self.Object.Set("objects", member)
}

// Array of associated child nodes.
func (self *QuadTree) GetNodesA() []interface{}{
	array00 := self.Object.Get("nodes")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Array of associated child nodes.
func (self *QuadTree) SetNodesA(member []interface{}) {
    self.Object.Set("nodes", member)
}



// Resets the QuadTree.
func (self *QuadTree) Reset(x int, y int, width int, height int, maxObjects int, maxLevels int, level int) {
    self.Object.Call("reset", x, y, width, height, maxObjects, maxLevels, level)
}

// Resets the QuadTree.
func (self *QuadTree) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Populates this quadtree with the children of the given Group. In order to be added the child must exist and have a body property.
func (self *QuadTree) Populate(group *Group) {
    self.Object.Call("populate", group)
}

// Populates this quadtree with the children of the given Group. In order to be added the child must exist and have a body property.
func (self *QuadTree) PopulateI(args ...interface{}) {
    self.Object.Call("populate", args)
}

// Handler for the populate method.
func (self *QuadTree) PopulateHandler(sprite interface{}) {
    self.Object.Call("populateHandler", sprite)
}

// Handler for the populate method.
func (self *QuadTree) PopulateHandlerI(args ...interface{}) {
    self.Object.Call("populateHandler", args)
}

// Split the node into 4 subnodes
func (self *QuadTree) Split() {
    self.Object.Call("split")
}

// Split the node into 4 subnodes
func (self *QuadTree) SplitI(args ...interface{}) {
    self.Object.Call("split", args)
}

// Insert the object into the node. If the node exceeds the capacity, it will split and add all objects to their corresponding subnodes.
func (self *QuadTree) Insert(body interface{}) {
    self.Object.Call("insert", body)
}

// Insert the object into the node. If the node exceeds the capacity, it will split and add all objects to their corresponding subnodes.
func (self *QuadTree) InsertI(args ...interface{}) {
    self.Object.Call("insert", args)
}

// Determine which node the object belongs to.
func (self *QuadTree) GetIndex(rect interface{}) int{
    return self.Object.Call("getIndex", rect).Int()
}

// Determine which node the object belongs to.
func (self *QuadTree) GetIndexI(args ...interface{}) int{
    return self.Object.Call("getIndex", args).Int()
}

// Return all objects that could collide with the given Sprite or Rectangle.
func (self *QuadTree) Retrieve(source interface{}) []interface{}{
	array00 := self.Object.Call("retrieve", source)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Return all objects that could collide with the given Sprite or Rectangle.
func (self *QuadTree) RetrieveI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("retrieve", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Clear the quadtree.
func (self *QuadTree) Clear() {
    self.Object.Call("clear")
}

// Clear the quadtree.
func (self *QuadTree) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}
