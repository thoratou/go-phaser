// Package phaser Automatic generation for Phaser.QuadTree
// generated file QuadTree.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// QuadTree A QuadTree implementation. The original code was a conversion of the Java code posted to GameDevTuts.
// However I've tweaked it massively to add node indexing, removed lots of temp. var creation and significantly increased performance as a result.
// Original version at https://github.com/timohausmann/quadtree-js/
type QuadTree struct {
    *js.Object
}

// NewQuadTree A QuadTree implementation. The original code was a conversion of the Java code posted to GameDevTuts.
// However I've tweaked it massively to add node indexing, removed lots of temp. var creation and significantly increased performance as a result.
// Original version at https://github.com/timohausmann/quadtree-js/
func NewQuadTree(x int, y int, width int, height int) *QuadTree {
    return &QuadTree{js.Global.Get("Phaser").Get("QuadTree").New(x, y, width, height)}
}
// NewQuadTree1O A QuadTree implementation. The original code was a conversion of the Java code posted to GameDevTuts.
// However I've tweaked it massively to add node indexing, removed lots of temp. var creation and significantly increased performance as a result.
// Original version at https://github.com/timohausmann/quadtree-js/
func NewQuadTree1O(x int, y int, width int, height int, maxObjects int) *QuadTree {
    return &QuadTree{js.Global.Get("Phaser").Get("QuadTree").New(x, y, width, height, maxObjects)}
}
// NewQuadTree2O A QuadTree implementation. The original code was a conversion of the Java code posted to GameDevTuts.
// However I've tweaked it massively to add node indexing, removed lots of temp. var creation and significantly increased performance as a result.
// Original version at https://github.com/timohausmann/quadtree-js/
func NewQuadTree2O(x int, y int, width int, height int, maxObjects int, maxLevels int) *QuadTree {
    return &QuadTree{js.Global.Get("Phaser").Get("QuadTree").New(x, y, width, height, maxObjects, maxLevels)}
}
// NewQuadTree3O A QuadTree implementation. The original code was a conversion of the Java code posted to GameDevTuts.
// However I've tweaked it massively to add node indexing, removed lots of temp. var creation and significantly increased performance as a result.
// Original version at https://github.com/timohausmann/quadtree-js/
func NewQuadTree3O(x int, y int, width int, height int, maxObjects int, maxLevels int, level int) *QuadTree {
    return &QuadTree{js.Global.Get("Phaser").Get("QuadTree").New(x, y, width, height, maxObjects, maxLevels, level)}
}
// NewQuadTreeI A QuadTree implementation. The original code was a conversion of the Java code posted to GameDevTuts.
// However I've tweaked it massively to add node indexing, removed lots of temp. var creation and significantly increased performance as a result.
// Original version at https://github.com/timohausmann/quadtree-js/
func NewQuadTreeI(args ...interface{}) *QuadTree {
    return &QuadTree{js.Global.Get("Phaser").Get("QuadTree").New(args)}
}



// MaxObjects The maximum number of objects per node.
func (self *QuadTree) MaxObjects() int{
    return self.Object.Get("maxObjects").Int()
}

// SetMaxObjectsA The maximum number of objects per node.
func (self *QuadTree) SetMaxObjectsA(member int) {
    self.Object.Set("maxObjects", member)
}

// MaxLevels The maximum number of levels to break down to.
func (self *QuadTree) MaxLevels() int{
    return self.Object.Get("maxLevels").Int()
}

// SetMaxLevelsA The maximum number of levels to break down to.
func (self *QuadTree) SetMaxLevelsA(member int) {
    self.Object.Set("maxLevels", member)
}

// Level The current level.
func (self *QuadTree) Level() int{
    return self.Object.Get("level").Int()
}

// SetLevelA The current level.
func (self *QuadTree) SetLevelA(member int) {
    self.Object.Set("level", member)
}

// Bounds Object that contains the quadtree bounds.
func (self *QuadTree) Bounds() interface{}{
    return self.Object.Get("bounds")
}

// SetBoundsA Object that contains the quadtree bounds.
func (self *QuadTree) SetBoundsA(member interface{}) {
    self.Object.Set("bounds", member)
}

// Objects Array of quadtree children.
func (self *QuadTree) Objects() []interface{}{
	array00 := self.Object.Get("objects")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetObjectsA Array of quadtree children.
func (self *QuadTree) SetObjectsA(member []interface{}) {
    self.Object.Set("objects", member)
}

// Nodes Array of associated child nodes.
func (self *QuadTree) Nodes() []interface{}{
	array00 := self.Object.Get("nodes")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetNodesA Array of associated child nodes.
func (self *QuadTree) SetNodesA(member []interface{}) {
    self.Object.Set("nodes", member)
}


// Reset Resets the QuadTree.
func (self *QuadTree) Reset(x int, y int, width int, height int) {
    self.Object.Call("reset", x, y, width, height)
}

// Reset1O Resets the QuadTree.
func (self *QuadTree) Reset1O(x int, y int, width int, height int, maxObjects int) {
    self.Object.Call("reset", x, y, width, height, maxObjects)
}

// Reset2O Resets the QuadTree.
func (self *QuadTree) Reset2O(x int, y int, width int, height int, maxObjects int, maxLevels int) {
    self.Object.Call("reset", x, y, width, height, maxObjects, maxLevels)
}

// Reset3O Resets the QuadTree.
func (self *QuadTree) Reset3O(x int, y int, width int, height int, maxObjects int, maxLevels int, level int) {
    self.Object.Call("reset", x, y, width, height, maxObjects, maxLevels, level)
}

// ResetI Resets the QuadTree.
func (self *QuadTree) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Populate Populates this quadtree with the children of the given Group. In order to be added the child must exist and have a body property.
func (self *QuadTree) Populate(group *Group) {
    self.Object.Call("populate", group)
}

// PopulateI Populates this quadtree with the children of the given Group. In order to be added the child must exist and have a body property.
func (self *QuadTree) PopulateI(args ...interface{}) {
    self.Object.Call("populate", args)
}

// PopulateHandler Handler for the populate method.
func (self *QuadTree) PopulateHandler(sprite interface{}) {
    self.Object.Call("populateHandler", sprite)
}

// PopulateHandlerI Handler for the populate method.
func (self *QuadTree) PopulateHandlerI(args ...interface{}) {
    self.Object.Call("populateHandler", args)
}

// Split Split the node into 4 subnodes
func (self *QuadTree) Split() {
    self.Object.Call("split")
}

// SplitI Split the node into 4 subnodes
func (self *QuadTree) SplitI(args ...interface{}) {
    self.Object.Call("split", args)
}

// Insert Insert the object into the node. If the node exceeds the capacity, it will split and add all objects to their corresponding subnodes.
func (self *QuadTree) Insert(body interface{}) {
    self.Object.Call("insert", body)
}

// InsertI Insert the object into the node. If the node exceeds the capacity, it will split and add all objects to their corresponding subnodes.
func (self *QuadTree) InsertI(args ...interface{}) {
    self.Object.Call("insert", args)
}

// GetIndex Determine which node the object belongs to.
func (self *QuadTree) GetIndex(rect interface{}) int{
    return self.Object.Call("getIndex", rect).Int()
}

// GetIndexI Determine which node the object belongs to.
func (self *QuadTree) GetIndexI(args ...interface{}) int{
    return self.Object.Call("getIndex", args).Int()
}

// Retrieve Return all objects that could collide with the given Sprite or Rectangle.
func (self *QuadTree) Retrieve(source interface{}) []interface{}{
	array00 := self.Object.Call("retrieve", source)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// RetrieveI Return all objects that could collide with the given Sprite or Rectangle.
func (self *QuadTree) RetrieveI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("retrieve", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Clear Clear the quadtree.
func (self *QuadTree) Clear() {
    self.Object.Call("clear")
}

// ClearI Clear the quadtree.
func (self *QuadTree) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

