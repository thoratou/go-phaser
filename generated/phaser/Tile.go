// Package phaser Automatic generation for Phaser.Tile
// generated file Tile.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Tile A Tile is a representation of a single tile within the Tilemap.
type Tile struct {
    *js.Object
}

// NewTile A Tile is a representation of a single tile within the Tilemap.
func NewTile(layer interface{}, index int, x int, y int, width int, height int) *Tile {
    return &Tile{js.Global.Get("Phaser").Get("Tile").New(layer, index, x, y, width, height)}
}
// NewTileI A Tile is a representation of a single tile within the Tilemap.
func NewTileI(args ...interface{}) *Tile {
    return &Tile{js.Global.Get("Phaser").Get("Tile").New(args)}
}



// Tile Binding conversion method to Tile point 
func ToTile(jsStruct interface{}) *Tile {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Tile{Object: object}
	}
	return nil
}



// Layer The layer in the Tilemap data that this tile belongs to.
func (self *Tile) Layer() interface{}{
    return self.Object.Get("layer")
}

// SetLayerA The layer in the Tilemap data that this tile belongs to.
func (self *Tile) SetLayerA(member interface{}) {
    self.Object.Set("layer", member)
}

// Index The index of this tile within the map data corresponding to the tileset, or -1 if this represents a blank/null tile.
func (self *Tile) Index() int{
    return self.Object.Get("index").Int()
}

// SetIndexA The index of this tile within the map data corresponding to the tileset, or -1 if this represents a blank/null tile.
func (self *Tile) SetIndexA(member int) {
    self.Object.Set("index", member)
}

// X The x map coordinate of this tile.
func (self *Tile) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The x map coordinate of this tile.
func (self *Tile) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The y map coordinate of this tile.
func (self *Tile) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The y map coordinate of this tile.
func (self *Tile) SetYA(member int) {
    self.Object.Set("y", member)
}

// Rotation The rotation angle of this tile.
func (self *Tile) Rotation() int{
    return self.Object.Get("rotation").Int()
}

// SetRotationA The rotation angle of this tile.
func (self *Tile) SetRotationA(member int) {
    self.Object.Set("rotation", member)
}

// Flipped Whether this tile is flipped (mirrored) or not.
func (self *Tile) Flipped() bool{
    return self.Object.Get("flipped").Bool()
}

// SetFlippedA Whether this tile is flipped (mirrored) or not.
func (self *Tile) SetFlippedA(member bool) {
    self.Object.Set("flipped", member)
}

// WorldX The x map coordinate of this tile.
func (self *Tile) WorldX() interface{}{
    return self.Object.Get("worldX")
}

// SetWorldXA The x map coordinate of this tile.
func (self *Tile) SetWorldXA(member interface{}) {
    self.Object.Set("worldX", member)
}

// WorldY The y map coordinate of this tile.
func (self *Tile) WorldY() interface{}{
    return self.Object.Get("worldY")
}

// SetWorldYA The y map coordinate of this tile.
func (self *Tile) SetWorldYA(member interface{}) {
    self.Object.Set("worldY", member)
}

// Width The width of the tile in pixels.
func (self *Tile) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the tile in pixels.
func (self *Tile) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the tile in pixels.
func (self *Tile) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the tile in pixels.
func (self *Tile) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// CenterX The width of the tile in pixels.
func (self *Tile) CenterX() interface{}{
    return self.Object.Get("centerX")
}

// SetCenterXA The width of the tile in pixels.
func (self *Tile) SetCenterXA(member interface{}) {
    self.Object.Set("centerX", member)
}

// CenterY The height of the tile in pixels.
func (self *Tile) CenterY() interface{}{
    return self.Object.Get("centerY")
}

// SetCenterYA The height of the tile in pixels.
func (self *Tile) SetCenterYA(member interface{}) {
    self.Object.Set("centerY", member)
}

// Alpha The alpha value at which this tile is drawn to the canvas.
func (self *Tile) Alpha() int{
    return self.Object.Get("alpha").Int()
}

// SetAlphaA The alpha value at which this tile is drawn to the canvas.
func (self *Tile) SetAlphaA(member int) {
    self.Object.Set("alpha", member)
}

// Properties Tile specific properties.
func (self *Tile) Properties() interface{}{
    return self.Object.Get("properties")
}

// SetPropertiesA Tile specific properties.
func (self *Tile) SetPropertiesA(member interface{}) {
    self.Object.Set("properties", member)
}

// Scanned Has this tile been walked / turned into a poly?
func (self *Tile) Scanned() bool{
    return self.Object.Get("scanned").Bool()
}

// SetScannedA Has this tile been walked / turned into a poly?
func (self *Tile) SetScannedA(member bool) {
    self.Object.Set("scanned", member)
}

// FaceTop Is the top of this tile an interesting edge?
func (self *Tile) FaceTop() bool{
    return self.Object.Get("faceTop").Bool()
}

// SetFaceTopA Is the top of this tile an interesting edge?
func (self *Tile) SetFaceTopA(member bool) {
    self.Object.Set("faceTop", member)
}

// FaceBottom Is the bottom of this tile an interesting edge?
func (self *Tile) FaceBottom() bool{
    return self.Object.Get("faceBottom").Bool()
}

// SetFaceBottomA Is the bottom of this tile an interesting edge?
func (self *Tile) SetFaceBottomA(member bool) {
    self.Object.Set("faceBottom", member)
}

// FaceLeft Is the left of this tile an interesting edge?
func (self *Tile) FaceLeft() bool{
    return self.Object.Get("faceLeft").Bool()
}

// SetFaceLeftA Is the left of this tile an interesting edge?
func (self *Tile) SetFaceLeftA(member bool) {
    self.Object.Set("faceLeft", member)
}

// FaceRight Is the right of this tile an interesting edge?
func (self *Tile) FaceRight() bool{
    return self.Object.Get("faceRight").Bool()
}

// SetFaceRightA Is the right of this tile an interesting edge?
func (self *Tile) SetFaceRightA(member bool) {
    self.Object.Set("faceRight", member)
}

// CollideLeft Indicating collide with any object on the left.
func (self *Tile) CollideLeft() bool{
    return self.Object.Get("collideLeft").Bool()
}

// SetCollideLeftA Indicating collide with any object on the left.
func (self *Tile) SetCollideLeftA(member bool) {
    self.Object.Set("collideLeft", member)
}

// CollideRight Indicating collide with any object on the right.
func (self *Tile) CollideRight() bool{
    return self.Object.Get("collideRight").Bool()
}

// SetCollideRightA Indicating collide with any object on the right.
func (self *Tile) SetCollideRightA(member bool) {
    self.Object.Set("collideRight", member)
}

// CollideUp Indicating collide with any object on the top.
func (self *Tile) CollideUp() bool{
    return self.Object.Get("collideUp").Bool()
}

// SetCollideUpA Indicating collide with any object on the top.
func (self *Tile) SetCollideUpA(member bool) {
    self.Object.Set("collideUp", member)
}

// CollideDown Indicating collide with any object on the bottom.
func (self *Tile) CollideDown() bool{
    return self.Object.Get("collideDown").Bool()
}

// SetCollideDownA Indicating collide with any object on the bottom.
func (self *Tile) SetCollideDownA(member bool) {
    self.Object.Set("collideDown", member)
}

// CollisionCallback Tile collision callback.
func (self *Tile) CollisionCallback() interface{}{
    return self.Object.Get("collisionCallback")
}

// SetCollisionCallbackA Tile collision callback.
func (self *Tile) SetCollisionCallbackA(member interface{}) {
    self.Object.Set("collisionCallback", member)
}

// CollisionCallbackContext The context in which the collision callback will be called.
func (self *Tile) CollisionCallbackContext() interface{}{
    return self.Object.Get("collisionCallbackContext")
}

// SetCollisionCallbackContextA The context in which the collision callback will be called.
func (self *Tile) SetCollisionCallbackContextA(member interface{}) {
    self.Object.Set("collisionCallbackContext", member)
}

// Collides True if this tile can collide on any of its faces.
func (self *Tile) Collides() bool{
    return self.Object.Get("collides").Bool()
}

// SetCollidesA True if this tile can collide on any of its faces.
func (self *Tile) SetCollidesA(member bool) {
    self.Object.Set("collides", member)
}

// CanCollide True if this tile can collide on any of its faces or has a collision callback set.
func (self *Tile) CanCollide() bool{
    return self.Object.Get("canCollide").Bool()
}

// SetCanCollideA True if this tile can collide on any of its faces or has a collision callback set.
func (self *Tile) SetCanCollideA(member bool) {
    self.Object.Set("canCollide", member)
}

// Left The x value in pixels.
func (self *Tile) Left() int{
    return self.Object.Get("left").Int()
}

// SetLeftA The x value in pixels.
func (self *Tile) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// Right The sum of the x and width properties.
func (self *Tile) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA The sum of the x and width properties.
func (self *Tile) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Top The y value.
func (self *Tile) Top() int{
    return self.Object.Get("top").Int()
}

// SetTopA The y value.
func (self *Tile) SetTopA(member int) {
    self.Object.Set("top", member)
}

// Bottom The sum of the y and height properties.
func (self *Tile) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA The sum of the y and height properties.
func (self *Tile) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}


// ContainsPoint Check if the given x and y world coordinates are within this Tile.
func (self *Tile) ContainsPoint(x int, y int) bool{
    return self.Object.Call("containsPoint", x, y).Bool()
}

// ContainsPointI Check if the given x and y world coordinates are within this Tile.
func (self *Tile) ContainsPointI(args ...interface{}) bool{
    return self.Object.Call("containsPoint", args).Bool()
}

// Intersects Check for intersection with this tile.
func (self *Tile) Intersects(x int, y int, right int, bottom int) {
    self.Object.Call("intersects", x, y, right, bottom)
}

// IntersectsI Check for intersection with this tile.
func (self *Tile) IntersectsI(args ...interface{}) {
    self.Object.Call("intersects", args)
}

// SetCollisionCallback Set a callback to be called when this tile is hit by an object.
// The callback must true true for collision processing to take place.
func (self *Tile) SetCollisionCallback(callback interface{}, context interface{}) {
    self.Object.Call("setCollisionCallback", callback, context)
}

// SetCollisionCallbackI Set a callback to be called when this tile is hit by an object.
// The callback must true true for collision processing to take place.
func (self *Tile) SetCollisionCallbackI(args ...interface{}) {
    self.Object.Call("setCollisionCallback", args)
}

// Destroy Clean up memory.
func (self *Tile) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Clean up memory.
func (self *Tile) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// SetCollision Sets the collision flags for each side of this tile and updates the interesting faces list.
func (self *Tile) SetCollision(left bool, right bool, up bool, down bool) {
    self.Object.Call("setCollision", left, right, up, down)
}

// SetCollisionI Sets the collision flags for each side of this tile and updates the interesting faces list.
func (self *Tile) SetCollisionI(args ...interface{}) {
    self.Object.Call("setCollision", args)
}

// ResetCollision Reset collision status flags.
func (self *Tile) ResetCollision() {
    self.Object.Call("resetCollision")
}

// ResetCollisionI Reset collision status flags.
func (self *Tile) ResetCollisionI(args ...interface{}) {
    self.Object.Call("resetCollision", args)
}

// IsInteresting Is this tile interesting?
func (self *Tile) IsInteresting(collides bool, faces bool) bool{
    return self.Object.Call("isInteresting", collides, faces).Bool()
}

// IsInterestingI Is this tile interesting?
func (self *Tile) IsInterestingI(args ...interface{}) bool{
    return self.Object.Call("isInteresting", args).Bool()
}

// Copy Copies the tile data and properties from the given tile to this tile.
func (self *Tile) Copy(tile *Tile) {
    self.Object.Call("copy", tile)
}

// CopyI Copies the tile data and properties from the given tile to this tile.
func (self *Tile) CopyI(args ...interface{}) {
    self.Object.Call("copy", args)
}

