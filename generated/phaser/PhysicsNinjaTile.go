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
func (self *PhysicsNinjaTile) GetBodyA() interface{}{
    return self.Object.Get("body")
}

// A reference to the body that owns this shape.
func (self *PhysicsNinjaTile) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// A reference to the physics system.
func (self *PhysicsNinjaTile) GetSystemA() *PhysicsNinja{
    return &PhysicsNinja{self.Object.Get("system")}
}

// A reference to the physics system.
func (self *PhysicsNinjaTile) SetSystemA(member *PhysicsNinja) {
    self.Object.Set("system", member)
}

// The ID of this Tile.
func (self *PhysicsNinjaTile) GetIdA() int{
    return self.Object.Get("id").Int()
}

// The ID of this Tile.
func (self *PhysicsNinjaTile) SetIdA(member int) {
    self.Object.Set("id", member)
}

// The type of this Tile.
func (self *PhysicsNinjaTile) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The type of this Tile.
func (self *PhysicsNinjaTile) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The position of this object.
func (self *PhysicsNinjaTile) GetPosA() *Point{
    return &Point{self.Object.Get("pos")}
}

// The position of this object.
func (self *PhysicsNinjaTile) SetPosA(member *Point) {
    self.Object.Set("pos", member)
}

// The position of this object in the previous update.
func (self *PhysicsNinjaTile) GetOldposA() *Point{
    return &Point{self.Object.Get("oldpos")}
}

// The position of this object in the previous update.
func (self *PhysicsNinjaTile) SetOldposA(member *Point) {
    self.Object.Set("oldpos", member)
}

// Half the width.
func (self *PhysicsNinjaTile) GetXwA() int{
    return self.Object.Get("xw").Int()
}

// Half the width.
func (self *PhysicsNinjaTile) SetXwA(member int) {
    self.Object.Set("xw", member)
}

// Half the height.
func (self *PhysicsNinjaTile) GetYwA() interface{}{
    return self.Object.Get("yw")
}

// Half the height.
func (self *PhysicsNinjaTile) SetYwA(member interface{}) {
    self.Object.Set("yw", member)
}

// The width.
func (self *PhysicsNinjaTile) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width.
func (self *PhysicsNinjaTile) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height.
func (self *PhysicsNinjaTile) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height.
func (self *PhysicsNinjaTile) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The velocity of this object.
func (self *PhysicsNinjaTile) GetVelocityA() *Point{
    return &Point{self.Object.Get("velocity")}
}

// The velocity of this object.
func (self *PhysicsNinjaTile) SetVelocityA(member *Point) {
    self.Object.Set("velocity", member)
}

// The x position.
func (self *PhysicsNinjaTile) GetXA() int{
    return self.Object.Get("x").Int()
}

// The x position.
func (self *PhysicsNinjaTile) SetXA(member int) {
    self.Object.Set("x", member)
}

// The y position.
func (self *PhysicsNinjaTile) GetYA() int{
    return self.Object.Get("y").Int()
}

// The y position.
func (self *PhysicsNinjaTile) SetYA(member int) {
    self.Object.Set("y", member)
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaTile) GetBottomA() int{
    return self.Object.Get("bottom").Int()
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaTile) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaTile) GetRightA() int{
    return self.Object.Get("right").Int()
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaTile) SetRightA(member int) {
    self.Object.Set("right", member)
}



// Updates this objects position.
func (self *PhysicsNinjaTile) Integrate() {
    self.Object.Call("integrate")
}

// Updates this objects position.
func (self *PhysicsNinjaTile) IntegrateI(args ...interface{}) {
    self.Object.Call("integrate", args)
}

// Tiles cannot collide with the world bounds, it's up to you to keep them where you want them. But we need this API stub to satisfy the Body.
func (self *PhysicsNinjaTile) CollideWorldBounds() {
    self.Object.Call("collideWorldBounds")
}

// Tiles cannot collide with the world bounds, it's up to you to keep them where you want them. But we need this API stub to satisfy the Body.
func (self *PhysicsNinjaTile) CollideWorldBoundsI(args ...interface{}) {
    self.Object.Call("collideWorldBounds", args)
}

// Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaTile) ReportCollisionVsWorld(px int, py int, dx int, dy int, obj int) {
    self.Object.Call("reportCollisionVsWorld", px, py, dx, dy, obj)
}

// Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaTile) ReportCollisionVsWorldI(args ...interface{}) {
    self.Object.Call("reportCollisionVsWorld", args)
}

// Tiles cannot collide with the world bounds, it's up to you to keep them where you want them. But we need this API stub to satisfy the Body.
func (self *PhysicsNinjaTile) SetType(id int) {
    self.Object.Call("setType", id)
}

// Tiles cannot collide with the world bounds, it's up to you to keep them where you want them. But we need this API stub to satisfy the Body.
func (self *PhysicsNinjaTile) SetTypeI(args ...interface{}) {
    self.Object.Call("setType", args)
}

// Sets this tile to be empty.
func (self *PhysicsNinjaTile) Clear() {
    self.Object.Call("clear")
}

// Sets this tile to be empty.
func (self *PhysicsNinjaTile) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// Destroys this Tiles reference to Body and System.
func (self *PhysicsNinjaTile) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this Tiles reference to Body and System.
func (self *PhysicsNinjaTile) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// This converts a tile from implicitly-defined (via id), to explicit (via properties).
// Don't call directly, instead of setType.
func (self *PhysicsNinjaTile) UpdateType() {
    self.Object.Call("updateType")
}

// This converts a tile from implicitly-defined (via id), to explicit (via properties).
// Don't call directly, instead of setType.
func (self *PhysicsNinjaTile) UpdateTypeI(args ...interface{}) {
    self.Object.Call("updateType", args)
}
