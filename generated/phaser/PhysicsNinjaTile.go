// Package phaser Automatic generation for Phaser.Physics.Ninja.Tile
// generated file PhysicsNinjaTile.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsNinjaTile Ninja Physics Tile constructor.
// A Tile is defined by its width, height and type. It's type can include slope data, such as 45 degree slopes, or convex slopes.
// Understand that for any type including a slope (types 2 to 29) the Tile must be SQUARE, i.e. have an equal width and height.
// Also note that as Tiles are primarily used for levels they have gravity disabled and world bounds collision disabled by default.
// 
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
type PhysicsNinjaTile struct {
    *js.Object
}

// NewPhysicsNinjaTile Ninja Physics Tile constructor.
// A Tile is defined by its width, height and type. It's type can include slope data, such as 45 degree slopes, or convex slopes.
// Understand that for any type including a slope (types 2 to 29) the Tile must be SQUARE, i.e. have an equal width and height.
// Also note that as Tiles are primarily used for levels they have gravity disabled and world bounds collision disabled by default.
// 
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaTile(body *PhysicsNinjaBody, x int, y int, width int, height int) *PhysicsNinjaTile {
    return &PhysicsNinjaTile{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Tile").New(body, x, y, width, height)}
}
// NewPhysicsNinjaTile1O Ninja Physics Tile constructor.
// A Tile is defined by its width, height and type. It's type can include slope data, such as 45 degree slopes, or convex slopes.
// Understand that for any type including a slope (types 2 to 29) the Tile must be SQUARE, i.e. have an equal width and height.
// Also note that as Tiles are primarily used for levels they have gravity disabled and world bounds collision disabled by default.
// 
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaTile1O(body *PhysicsNinjaBody, x int, y int, width int, height int, type_ int) *PhysicsNinjaTile {
    return &PhysicsNinjaTile{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Tile").New(body, x, y, width, height, type_)}
}
// NewPhysicsNinjaTileI Ninja Physics Tile constructor.
// A Tile is defined by its width, height and type. It's type can include slope data, such as 45 degree slopes, or convex slopes.
// Understand that for any type including a slope (types 2 to 29) the Tile must be SQUARE, i.e. have an equal width and height.
// Also note that as Tiles are primarily used for levels they have gravity disabled and world bounds collision disabled by default.
// 
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaTileI(args ...interface{}) *PhysicsNinjaTile {
    return &PhysicsNinjaTile{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Tile").New(args)}
}



// PhysicsNinjaTile Binding conversion method to PhysicsNinjaTile point 
func ToPhysicsNinjaTile(jsStruct interface{}) *PhysicsNinjaTile {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsNinjaTile{Object: object}
	}
	return nil
}



// Body A reference to the body that owns this shape.
func (self *PhysicsNinjaTile) Body() interface{}{
    return self.Object.Get("body")
}

// SetBodyA A reference to the body that owns this shape.
func (self *PhysicsNinjaTile) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// System A reference to the physics system.
func (self *PhysicsNinjaTile) System() *PhysicsNinja{
    return &PhysicsNinja{self.Object.Get("system")}
}

// SetSystemA A reference to the physics system.
func (self *PhysicsNinjaTile) SetSystemA(member *PhysicsNinja) {
    self.Object.Set("system", member)
}

// Id The ID of this Tile.
func (self *PhysicsNinjaTile) Id() int{
    return self.Object.Get("id").Int()
}

// SetIdA The ID of this Tile.
func (self *PhysicsNinjaTile) SetIdA(member int) {
    self.Object.Set("id", member)
}

// Type The type of this Tile.
func (self *PhysicsNinjaTile) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The type of this Tile.
func (self *PhysicsNinjaTile) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Pos The position of this object.
func (self *PhysicsNinjaTile) Pos() *Point{
    return &Point{self.Object.Get("pos")}
}

// SetPosA The position of this object.
func (self *PhysicsNinjaTile) SetPosA(member *Point) {
    self.Object.Set("pos", member)
}

// Oldpos The position of this object in the previous update.
func (self *PhysicsNinjaTile) Oldpos() *Point{
    return &Point{self.Object.Get("oldpos")}
}

// SetOldposA The position of this object in the previous update.
func (self *PhysicsNinjaTile) SetOldposA(member *Point) {
    self.Object.Set("oldpos", member)
}

// Xw Half the width.
func (self *PhysicsNinjaTile) Xw() int{
    return self.Object.Get("xw").Int()
}

// SetXwA Half the width.
func (self *PhysicsNinjaTile) SetXwA(member int) {
    self.Object.Set("xw", member)
}

// Yw Half the height.
func (self *PhysicsNinjaTile) Yw() interface{}{
    return self.Object.Get("yw")
}

// SetYwA Half the height.
func (self *PhysicsNinjaTile) SetYwA(member interface{}) {
    self.Object.Set("yw", member)
}

// Width The width.
func (self *PhysicsNinjaTile) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width.
func (self *PhysicsNinjaTile) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height.
func (self *PhysicsNinjaTile) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height.
func (self *PhysicsNinjaTile) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Velocity The velocity of this object.
func (self *PhysicsNinjaTile) Velocity() *Point{
    return &Point{self.Object.Get("velocity")}
}

// SetVelocityA The velocity of this object.
func (self *PhysicsNinjaTile) SetVelocityA(member *Point) {
    self.Object.Set("velocity", member)
}

// X The x position.
func (self *PhysicsNinjaTile) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The x position.
func (self *PhysicsNinjaTile) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The y position.
func (self *PhysicsNinjaTile) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The y position.
func (self *PhysicsNinjaTile) SetYA(member int) {
    self.Object.Set("y", member)
}

// Bottom The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaTile) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaTile) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// Right The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaTile) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaTile) SetRightA(member int) {
    self.Object.Set("right", member)
}


// Integrate Updates this objects position.
func (self *PhysicsNinjaTile) Integrate() {
    self.Object.Call("integrate")
}

// IntegrateI Updates this objects position.
func (self *PhysicsNinjaTile) IntegrateI(args ...interface{}) {
    self.Object.Call("integrate", args)
}

// CollideWorldBounds Tiles cannot collide with the world bounds, it's up to you to keep them where you want them. But we need this API stub to satisfy the Body.
func (self *PhysicsNinjaTile) CollideWorldBounds() {
    self.Object.Call("collideWorldBounds")
}

// CollideWorldBoundsI Tiles cannot collide with the world bounds, it's up to you to keep them where you want them. But we need this API stub to satisfy the Body.
func (self *PhysicsNinjaTile) CollideWorldBoundsI(args ...interface{}) {
    self.Object.Call("collideWorldBounds", args)
}

// ReportCollisionVsWorld Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaTile) ReportCollisionVsWorld(px int, py int, dx int, dy int, obj int) {
    self.Object.Call("reportCollisionVsWorld", px, py, dx, dy, obj)
}

// ReportCollisionVsWorldI Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaTile) ReportCollisionVsWorldI(args ...interface{}) {
    self.Object.Call("reportCollisionVsWorld", args)
}

// SetType Tiles cannot collide with the world bounds, it's up to you to keep them where you want them. But we need this API stub to satisfy the Body.
func (self *PhysicsNinjaTile) SetType(id int) {
    self.Object.Call("setType", id)
}

// SetTypeI Tiles cannot collide with the world bounds, it's up to you to keep them where you want them. But we need this API stub to satisfy the Body.
func (self *PhysicsNinjaTile) SetTypeI(args ...interface{}) {
    self.Object.Call("setType", args)
}

// Clear Sets this tile to be empty.
func (self *PhysicsNinjaTile) Clear() {
    self.Object.Call("clear")
}

// ClearI Sets this tile to be empty.
func (self *PhysicsNinjaTile) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// Destroy Destroys this Tiles reference to Body and System.
func (self *PhysicsNinjaTile) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this Tiles reference to Body and System.
func (self *PhysicsNinjaTile) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// UpdateType This converts a tile from implicitly-defined (via id), to explicit (via properties).
// Don't call directly, instead of setType.
func (self *PhysicsNinjaTile) UpdateType() {
    self.Object.Call("updateType")
}

// UpdateTypeI This converts a tile from implicitly-defined (via id), to explicit (via properties).
// Don't call directly, instead of setType.
func (self *PhysicsNinjaTile) UpdateTypeI(args ...interface{}) {
    self.Object.Call("updateType", args)
}

