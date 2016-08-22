// Automatic generation for Phaser.Tile
// generated file Tile.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A Tile is a representation of a single tile within the Tilemap.
type Tile struct {
    *js.Object
}


// The layer in the Tilemap data that this tile belongs to.
func (self *Tile) GetLayerA() interface{}{
    return self.Object.Get("layer")
}

// The layer in the Tilemap data that this tile belongs to.
func (self *Tile) SetLayerA(member interface{}) {
    self.Object.Set("layer", member)
}

// The index of this tile within the map data corresponding to the tileset, or -1 if this represents a blank/null tile.
func (self *Tile) GetIndexA() int{
    return self.Object.Get("index").Int()
}

// The index of this tile within the map data corresponding to the tileset, or -1 if this represents a blank/null tile.
func (self *Tile) SetIndexA(member int) {
    self.Object.Set("index", member)
}

// The x map coordinate of this tile.
func (self *Tile) GetXA() int{
    return self.Object.Get("x").Int()
}

// The x map coordinate of this tile.
func (self *Tile) SetXA(member int) {
    self.Object.Set("x", member)
}

// The y map coordinate of this tile.
func (self *Tile) GetYA() int{
    return self.Object.Get("y").Int()
}

// The y map coordinate of this tile.
func (self *Tile) SetYA(member int) {
    self.Object.Set("y", member)
}

// The rotation angle of this tile.
func (self *Tile) GetRotationA() int{
    return self.Object.Get("rotation").Int()
}

// The rotation angle of this tile.
func (self *Tile) SetRotationA(member int) {
    self.Object.Set("rotation", member)
}

// Whether this tile is flipped (mirrored) or not.
func (self *Tile) GetFlippedA() bool{
    return self.Object.Get("flipped").Bool()
}

// Whether this tile is flipped (mirrored) or not.
func (self *Tile) SetFlippedA(member bool) {
    self.Object.Set("flipped", member)
}

// The x map coordinate of this tile.
func (self *Tile) GetWorldXA() interface{}{
    return self.Object.Get("worldX")
}

// The x map coordinate of this tile.
func (self *Tile) SetWorldXA(member interface{}) {
    self.Object.Set("worldX", member)
}

// The y map coordinate of this tile.
func (self *Tile) GetWorldYA() interface{}{
    return self.Object.Get("worldY")
}

// The y map coordinate of this tile.
func (self *Tile) SetWorldYA(member interface{}) {
    self.Object.Set("worldY", member)
}

// The width of the tile in pixels.
func (self *Tile) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the tile in pixels.
func (self *Tile) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the tile in pixels.
func (self *Tile) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the tile in pixels.
func (self *Tile) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The width of the tile in pixels.
func (self *Tile) GetCenterXA() interface{}{
    return self.Object.Get("centerX")
}

// The width of the tile in pixels.
func (self *Tile) SetCenterXA(member interface{}) {
    self.Object.Set("centerX", member)
}

// The height of the tile in pixels.
func (self *Tile) GetCenterYA() interface{}{
    return self.Object.Get("centerY")
}

// The height of the tile in pixels.
func (self *Tile) SetCenterYA(member interface{}) {
    self.Object.Set("centerY", member)
}

// The alpha value at which this tile is drawn to the canvas.
func (self *Tile) GetAlphaA() int{
    return self.Object.Get("alpha").Int()
}

// The alpha value at which this tile is drawn to the canvas.
func (self *Tile) SetAlphaA(member int) {
    self.Object.Set("alpha", member)
}

// Tile specific properties.
func (self *Tile) GetPropertiesA() interface{}{
    return self.Object.Get("properties")
}

// Tile specific properties.
func (self *Tile) SetPropertiesA(member interface{}) {
    self.Object.Set("properties", member)
}

// Has this tile been walked / turned into a poly?
func (self *Tile) GetScannedA() bool{
    return self.Object.Get("scanned").Bool()
}

// Has this tile been walked / turned into a poly?
func (self *Tile) SetScannedA(member bool) {
    self.Object.Set("scanned", member)
}

// Is the top of this tile an interesting edge?
func (self *Tile) GetFaceTopA() bool{
    return self.Object.Get("faceTop").Bool()
}

// Is the top of this tile an interesting edge?
func (self *Tile) SetFaceTopA(member bool) {
    self.Object.Set("faceTop", member)
}

// Is the bottom of this tile an interesting edge?
func (self *Tile) GetFaceBottomA() bool{
    return self.Object.Get("faceBottom").Bool()
}

// Is the bottom of this tile an interesting edge?
func (self *Tile) SetFaceBottomA(member bool) {
    self.Object.Set("faceBottom", member)
}

// Is the left of this tile an interesting edge?
func (self *Tile) GetFaceLeftA() bool{
    return self.Object.Get("faceLeft").Bool()
}

// Is the left of this tile an interesting edge?
func (self *Tile) SetFaceLeftA(member bool) {
    self.Object.Set("faceLeft", member)
}

// Is the right of this tile an interesting edge?
func (self *Tile) GetFaceRightA() bool{
    return self.Object.Get("faceRight").Bool()
}

// Is the right of this tile an interesting edge?
func (self *Tile) SetFaceRightA(member bool) {
    self.Object.Set("faceRight", member)
}

// Indicating collide with any object on the left.
func (self *Tile) GetCollideLeftA() bool{
    return self.Object.Get("collideLeft").Bool()
}

// Indicating collide with any object on the left.
func (self *Tile) SetCollideLeftA(member bool) {
    self.Object.Set("collideLeft", member)
}

// Indicating collide with any object on the right.
func (self *Tile) GetCollideRightA() bool{
    return self.Object.Get("collideRight").Bool()
}

// Indicating collide with any object on the right.
func (self *Tile) SetCollideRightA(member bool) {
    self.Object.Set("collideRight", member)
}

// Indicating collide with any object on the top.
func (self *Tile) GetCollideUpA() bool{
    return self.Object.Get("collideUp").Bool()
}

// Indicating collide with any object on the top.
func (self *Tile) SetCollideUpA(member bool) {
    self.Object.Set("collideUp", member)
}

// Indicating collide with any object on the bottom.
func (self *Tile) GetCollideDownA() bool{
    return self.Object.Get("collideDown").Bool()
}

// Indicating collide with any object on the bottom.
func (self *Tile) SetCollideDownA(member bool) {
    self.Object.Set("collideDown", member)
}

// Tile collision callback.
func (self *Tile) SetCollisionCallbackA(member func(...interface{})) {
    self.Object.Set("collisionCallback", member)
}

// The context in which the collision callback will be called.
func (self *Tile) GetCollisionCallbackContextA() interface{}{
    return self.Object.Get("collisionCallbackContext")
}

// The context in which the collision callback will be called.
func (self *Tile) SetCollisionCallbackContextA(member interface{}) {
    self.Object.Set("collisionCallbackContext", member)
}

// True if this tile can collide on any of its faces.
func (self *Tile) GetCollidesA() bool{
    return self.Object.Get("collides").Bool()
}

// True if this tile can collide on any of its faces.
func (self *Tile) SetCollidesA(member bool) {
    self.Object.Set("collides", member)
}

// True if this tile can collide on any of its faces or has a collision callback set.
func (self *Tile) GetCanCollideA() bool{
    return self.Object.Get("canCollide").Bool()
}

// True if this tile can collide on any of its faces or has a collision callback set.
func (self *Tile) SetCanCollideA(member bool) {
    self.Object.Set("canCollide", member)
}

// The x value in pixels.
func (self *Tile) GetLeftA() int{
    return self.Object.Get("left").Int()
}

// The x value in pixels.
func (self *Tile) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// The sum of the x and width properties.
func (self *Tile) GetRightA() int{
    return self.Object.Get("right").Int()
}

// The sum of the x and width properties.
func (self *Tile) SetRightA(member int) {
    self.Object.Set("right", member)
}

// The y value.
func (self *Tile) GetTopA() int{
    return self.Object.Get("top").Int()
}

// The y value.
func (self *Tile) SetTopA(member int) {
    self.Object.Set("top", member)
}

// The sum of the y and height properties.
func (self *Tile) GetBottomA() int{
    return self.Object.Get("bottom").Int()
}

// The sum of the y and height properties.
func (self *Tile) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}



// Check if the given x and y world coordinates are within this Tile.
func (self *Tile) ContainsPoint(x int, y int) bool{
    return self.Object.Call("containsPoint", x, y).Bool()
}

// Check if the given x and y world coordinates are within this Tile.
func (self *Tile) ContainsPointI(args ...interface{}) bool{
    return self.Object.Call("containsPoint", args).Bool()
}

// Check for intersection with this tile.
func (self *Tile) Intersects(x int, y int, right int, bottom int) {
    self.Object.Call("intersects", x, y, right, bottom)
}

// Check for intersection with this tile.
func (self *Tile) IntersectsI(args ...interface{}) {
    self.Object.Call("intersects", args)
}

// Set a callback to be called when this tile is hit by an object.
// The callback must true true for collision processing to take place.
func (self *Tile) SetCollisionCallback(callback func(...interface{}), context interface{}) {
    self.Object.Call("setCollisionCallback", callback, context)
}

// Set a callback to be called when this tile is hit by an object.
// The callback must true true for collision processing to take place.
func (self *Tile) SetCollisionCallbackI(args ...interface{}) {
    self.Object.Call("setCollisionCallback", args)
}

// Clean up memory.
func (self *Tile) Destroy() {
    self.Object.Call("destroy")
}

// Clean up memory.
func (self *Tile) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Sets the collision flags for each side of this tile and updates the interesting faces list.
func (self *Tile) SetCollision(left bool, right bool, up bool, down bool) {
    self.Object.Call("setCollision", left, right, up, down)
}

// Sets the collision flags for each side of this tile and updates the interesting faces list.
func (self *Tile) SetCollisionI(args ...interface{}) {
    self.Object.Call("setCollision", args)
}

// Reset collision status flags.
func (self *Tile) ResetCollision() {
    self.Object.Call("resetCollision")
}

// Reset collision status flags.
func (self *Tile) ResetCollisionI(args ...interface{}) {
    self.Object.Call("resetCollision", args)
}

// Is this tile interesting?
func (self *Tile) IsInteresting(collides bool, faces bool) bool{
    return self.Object.Call("isInteresting", collides, faces).Bool()
}

// Is this tile interesting?
func (self *Tile) IsInterestingI(args ...interface{}) bool{
    return self.Object.Call("isInteresting", args).Bool()
}

// Copies the tile data and properties from the given tile to this tile.
func (self *Tile) Copy(tile *Tile) {
    self.Object.Call("copy", tile)
}

// Copies the tile data and properties from the given tile to this tile.
func (self *Tile) CopyI(args ...interface{}) {
    self.Object.Call("copy", args)
}
