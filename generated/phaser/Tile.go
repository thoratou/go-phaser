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
func (self *Tile) GetLayer() interface{}{
    return self.Get("layer")
}

// The layer in the Tilemap data that this tile belongs to.
func (self *Tile) SetLayer(member interface{}) {
    self.Set("layer", member)
}

// The index of this tile within the map data corresponding to the tileset, or -1 if this represents a blank/null tile.
func (self *Tile) GetIndex() int{
    return self.Get("index").Int()
}

// The index of this tile within the map data corresponding to the tileset, or -1 if this represents a blank/null tile.
func (self *Tile) SetIndex(member int) {
    self.Set("index", member)
}

// The x map coordinate of this tile.
func (self *Tile) GetX() int{
    return self.Get("x").Int()
}

// The x map coordinate of this tile.
func (self *Tile) SetX(member int) {
    self.Set("x", member)
}

// The y map coordinate of this tile.
func (self *Tile) GetY() int{
    return self.Get("y").Int()
}

// The y map coordinate of this tile.
func (self *Tile) SetY(member int) {
    self.Set("y", member)
}

// The rotation angle of this tile.
func (self *Tile) GetRotation() int{
    return self.Get("rotation").Int()
}

// The rotation angle of this tile.
func (self *Tile) SetRotation(member int) {
    self.Set("rotation", member)
}

// Whether this tile is flipped (mirrored) or not.
func (self *Tile) GetFlipped() bool{
    return self.Get("flipped").Bool()
}

// Whether this tile is flipped (mirrored) or not.
func (self *Tile) SetFlipped(member bool) {
    self.Set("flipped", member)
}

// The x map coordinate of this tile.
func (self *Tile) GetWorldX() interface{}{
    return self.Get("worldX")
}

// The x map coordinate of this tile.
func (self *Tile) SetWorldX(member interface{}) {
    self.Set("worldX", member)
}

// The y map coordinate of this tile.
func (self *Tile) GetWorldY() interface{}{
    return self.Get("worldY")
}

// The y map coordinate of this tile.
func (self *Tile) SetWorldY(member interface{}) {
    self.Set("worldY", member)
}

// The width of the tile in pixels.
func (self *Tile) GetWidth() int{
    return self.Get("width").Int()
}

// The width of the tile in pixels.
func (self *Tile) SetWidth(member int) {
    self.Set("width", member)
}

// The height of the tile in pixels.
func (self *Tile) GetHeight() int{
    return self.Get("height").Int()
}

// The height of the tile in pixels.
func (self *Tile) SetHeight(member int) {
    self.Set("height", member)
}

// The width of the tile in pixels.
func (self *Tile) GetCenterX() interface{}{
    return self.Get("centerX")
}

// The width of the tile in pixels.
func (self *Tile) SetCenterX(member interface{}) {
    self.Set("centerX", member)
}

// The height of the tile in pixels.
func (self *Tile) GetCenterY() interface{}{
    return self.Get("centerY")
}

// The height of the tile in pixels.
func (self *Tile) SetCenterY(member interface{}) {
    self.Set("centerY", member)
}

// The alpha value at which this tile is drawn to the canvas.
func (self *Tile) GetAlpha() int{
    return self.Get("alpha").Int()
}

// The alpha value at which this tile is drawn to the canvas.
func (self *Tile) SetAlpha(member int) {
    self.Set("alpha", member)
}

// Tile specific properties.
func (self *Tile) GetProperties() interface{}{
    return self.Get("properties")
}

// Tile specific properties.
func (self *Tile) SetProperties(member interface{}) {
    self.Set("properties", member)
}

// Has this tile been walked / turned into a poly?
func (self *Tile) GetScanned() bool{
    return self.Get("scanned").Bool()
}

// Has this tile been walked / turned into a poly?
func (self *Tile) SetScanned(member bool) {
    self.Set("scanned", member)
}

// Is the top of this tile an interesting edge?
func (self *Tile) GetFaceTop() bool{
    return self.Get("faceTop").Bool()
}

// Is the top of this tile an interesting edge?
func (self *Tile) SetFaceTop(member bool) {
    self.Set("faceTop", member)
}

// Is the bottom of this tile an interesting edge?
func (self *Tile) GetFaceBottom() bool{
    return self.Get("faceBottom").Bool()
}

// Is the bottom of this tile an interesting edge?
func (self *Tile) SetFaceBottom(member bool) {
    self.Set("faceBottom", member)
}

// Is the left of this tile an interesting edge?
func (self *Tile) GetFaceLeft() bool{
    return self.Get("faceLeft").Bool()
}

// Is the left of this tile an interesting edge?
func (self *Tile) SetFaceLeft(member bool) {
    self.Set("faceLeft", member)
}

// Is the right of this tile an interesting edge?
func (self *Tile) GetFaceRight() bool{
    return self.Get("faceRight").Bool()
}

// Is the right of this tile an interesting edge?
func (self *Tile) SetFaceRight(member bool) {
    self.Set("faceRight", member)
}

// Indicating collide with any object on the left.
func (self *Tile) GetCollideLeft() bool{
    return self.Get("collideLeft").Bool()
}

// Indicating collide with any object on the left.
func (self *Tile) SetCollideLeft(member bool) {
    self.Set("collideLeft", member)
}

// Indicating collide with any object on the right.
func (self *Tile) GetCollideRight() bool{
    return self.Get("collideRight").Bool()
}

// Indicating collide with any object on the right.
func (self *Tile) SetCollideRight(member bool) {
    self.Set("collideRight", member)
}

// Indicating collide with any object on the top.
func (self *Tile) GetCollideUp() bool{
    return self.Get("collideUp").Bool()
}

// Indicating collide with any object on the top.
func (self *Tile) SetCollideUp(member bool) {
    self.Set("collideUp", member)
}

// Indicating collide with any object on the bottom.
func (self *Tile) GetCollideDown() bool{
    return self.Get("collideDown").Bool()
}

// Indicating collide with any object on the bottom.
func (self *Tile) SetCollideDown(member bool) {
    self.Set("collideDown", member)
}

// Tile collision callback.
func (self *Tile) SetCollisionCallback(member func(...interface{})) {
    self.Set("collisionCallback", member)
}

// The context in which the collision callback will be called.
func (self *Tile) GetCollisionCallbackContext() interface{}{
    return self.Get("collisionCallbackContext")
}

// The context in which the collision callback will be called.
func (self *Tile) SetCollisionCallbackContext(member interface{}) {
    self.Set("collisionCallbackContext", member)
}

// True if this tile can collide on any of its faces.
func (self *Tile) GetCollides() bool{
    return self.Get("collides").Bool()
}

// True if this tile can collide on any of its faces.
func (self *Tile) SetCollides(member bool) {
    self.Set("collides", member)
}

// True if this tile can collide on any of its faces or has a collision callback set.
func (self *Tile) GetCanCollide() bool{
    return self.Get("canCollide").Bool()
}

// True if this tile can collide on any of its faces or has a collision callback set.
func (self *Tile) SetCanCollide(member bool) {
    self.Set("canCollide", member)
}

// The x value in pixels.
func (self *Tile) GetLeft() int{
    return self.Get("left").Int()
}

// The x value in pixels.
func (self *Tile) SetLeft(member int) {
    self.Set("left", member)
}

// The sum of the x and width properties.
func (self *Tile) GetRight() int{
    return self.Get("right").Int()
}

// The sum of the x and width properties.
func (self *Tile) SetRight(member int) {
    self.Set("right", member)
}

// The y value.
func (self *Tile) GetTop() int{
    return self.Get("top").Int()
}

// The y value.
func (self *Tile) SetTop(member int) {
    self.Set("top", member)
}

// The sum of the y and height properties.
func (self *Tile) GetBottom() int{
    return self.Get("bottom").Int()
}

// The sum of the y and height properties.
func (self *Tile) SetBottom(member int) {
    self.Set("bottom", member)
}



// Check if the given x and y world coordinates are within this Tile.
func (self *Tile) ContainsPointI(args ...interface{}) bool{
    return self.Call("containsPoint", args).Bool()
}

// Check for intersection with this tile.
func (self *Tile) IntersectsI(args ...interface{}) {
    self.Call("intersects", args)
}

// Set a callback to be called when this tile is hit by an object.
// The callback must true true for collision processing to take place.
func (self *Tile) SetCollisionCallbackI(args ...interface{}) {
    self.Call("setCollisionCallback", args)
}

// Clean up memory.
func (self *Tile) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Sets the collision flags for each side of this tile and updates the interesting faces list.
func (self *Tile) SetCollisionI(args ...interface{}) {
    self.Call("setCollision", args)
}

// Reset collision status flags.
func (self *Tile) ResetCollisionI(args ...interface{}) {
    self.Call("resetCollision", args)
}

// Is this tile interesting?
func (self *Tile) IsInterestingI(args ...interface{}) bool{
    return self.Call("isInteresting", args).Bool()
}

// Copies the tile data and properties from the given tile to this tile.
func (self *Tile) CopyI(args ...interface{}) {
    self.Call("copy", args)
}
