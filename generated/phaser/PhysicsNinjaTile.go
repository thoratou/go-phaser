// Automatic generation for Phaser.Physics.Ninja.Tile
// generated file PhysicsNinjaTile.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Ninja Physics Tile constructor.
// A Tile is defined by its width, height and type. It's type can include slope data, such as 45 degree slopes, or convex slopes.
// Understand that for any type including a slope (types 2 to 29) the Tile must be SQUARE, i.e. have an equal width and height.
// Also note that as Tiles are primarily used for levels they have gravity disabled and world bounds collision disabled by default.
// 
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
type PhysicsNinjaTile struct {
    *js.Object
}


// A reference to the body that owns this shape.
func (self *PhysicsNinjaTile) GetBody() interface{}{
    return self.Get("body")
}

// A reference to the body that owns this shape.
func (self *PhysicsNinjaTile) SetBody(member interface{}) {
    self.Set("body", member)
}

// A reference to the physics system.
func (self *PhysicsNinjaTile) GetSystem() PhysicsNinja{
    return PhysicsNinja{self.Get("system")}
}

// A reference to the physics system.
func (self *PhysicsNinjaTile) SetSystem(member PhysicsNinja) {
    self.Set("system", member)
}

// The ID of this Tile.
func (self *PhysicsNinjaTile) GetId() float64{
    return self.Get("id").Float()
}

// The ID of this Tile.
func (self *PhysicsNinjaTile) SetId(member float64) {
    self.Set("id", member)
}

// The type of this Tile.
func (self *PhysicsNinjaTile) GetType() float64{
    return self.Get("type").Float()
}

// The type of this Tile.
func (self *PhysicsNinjaTile) SetType(member float64) {
    self.Set("type", member)
}

// The position of this object.
func (self *PhysicsNinjaTile) GetPos() Point{
    return Point{self.Get("pos")}
}

// The position of this object.
func (self *PhysicsNinjaTile) SetPos(member Point) {
    self.Set("pos", member)
}

// The position of this object in the previous update.
func (self *PhysicsNinjaTile) GetOldpos() Point{
    return Point{self.Get("oldpos")}
}

// The position of this object in the previous update.
func (self *PhysicsNinjaTile) SetOldpos(member Point) {
    self.Set("oldpos", member)
}

// Half the width.
func (self *PhysicsNinjaTile) GetXw() float64{
    return self.Get("xw").Float()
}

// Half the width.
func (self *PhysicsNinjaTile) SetXw(member float64) {
    self.Set("xw", member)
}

// Half the height.
func (self *PhysicsNinjaTile) GetYw() interface{}{
    return self.Get("yw")
}

// Half the height.
func (self *PhysicsNinjaTile) SetYw(member interface{}) {
    self.Set("yw", member)
}

// The width.
func (self *PhysicsNinjaTile) GetWidth() float64{
    return self.Get("width").Float()
}

// The width.
func (self *PhysicsNinjaTile) SetWidth(member float64) {
    self.Set("width", member)
}

// The height.
func (self *PhysicsNinjaTile) GetHeight() float64{
    return self.Get("height").Float()
}

// The height.
func (self *PhysicsNinjaTile) SetHeight(member float64) {
    self.Set("height", member)
}

// The velocity of this object.
func (self *PhysicsNinjaTile) GetVelocity() Point{
    return Point{self.Get("velocity")}
}

// The velocity of this object.
func (self *PhysicsNinjaTile) SetVelocity(member Point) {
    self.Set("velocity", member)
}

// The x position.
func (self *PhysicsNinjaTile) GetX() float64{
    return self.Get("x").Float()
}

// The x position.
func (self *PhysicsNinjaTile) SetX(member float64) {
    self.Set("x", member)
}

// The y position.
func (self *PhysicsNinjaTile) GetY() float64{
    return self.Get("y").Float()
}

// The y position.
func (self *PhysicsNinjaTile) SetY(member float64) {
    self.Set("y", member)
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaTile) GetBottom() float64{
    return self.Get("bottom").Float()
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaTile) SetBottom(member float64) {
    self.Set("bottom", member)
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaTile) GetRight() float64{
    return self.Get("right").Float()
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaTile) SetRight(member float64) {
    self.Set("right", member)
}



// Updates this objects position.
func (self *PhysicsNinjaTile) IntegrateI(args ...interface{}) {
    self.Call("integrate", args)
}

// Tiles cannot collide with the world bounds, it's up to you to keep them where you want them. But we need this API stub to satisfy the Body.
func (self *PhysicsNinjaTile) CollideWorldBoundsI(args ...interface{}) {
    self.Call("collideWorldBounds", args)
}

// Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaTile) ReportCollisionVsWorldI(args ...interface{}) {
    self.Call("reportCollisionVsWorld", args)
}

// Tiles cannot collide with the world bounds, it's up to you to keep them where you want them. But we need this API stub to satisfy the Body.
func (self *PhysicsNinjaTile) SetTypeI(args ...interface{}) {
    self.Call("setType", args)
}

// Sets this tile to be empty.
func (self *PhysicsNinjaTile) ClearI(args ...interface{}) {
    self.Call("clear", args)
}

// Destroys this Tiles reference to Body and System.
func (self *PhysicsNinjaTile) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// This converts a tile from implicitly-defined (via id), to explicit (via properties).
// Don't call directly, instead of setType.
func (self *PhysicsNinjaTile) UpdateTypeI(args ...interface{}) {
    self.Call("updateType", args)
}
