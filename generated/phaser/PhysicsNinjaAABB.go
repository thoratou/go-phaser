// Automatic generation for Phaser.Physics.Ninja.AABB
// generated file PhysicsNinjaAABB.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Ninja Physics AABB constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
type PhysicsNinjaAABB struct {
    *js.Object
}


// Ninja Physics AABB constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaAABB(body *PhysicsNinjaBody, x int, y int, width int, height int) *PhysicsNinjaAABB {
    return &PhysicsNinjaAABB{js.Global.Call("Phaser.Physics.Ninja.AABB", body, x, y, width, height)}
}

// Ninja Physics AABB constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaAABBI(args ...interface{}) *PhysicsNinjaAABB {
    return &PhysicsNinjaAABB{js.Global.Call("Phaser.Physics.Ninja.AABB", args)}
}



// A reference to the body that owns this shape.
func (self *PhysicsNinjaAABB) GetBodyA() interface{}{
    return self.Object.Get("body")
}

// A reference to the body that owns this shape.
func (self *PhysicsNinjaAABB) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// A reference to the physics system.
func (self *PhysicsNinjaAABB) GetSystemA() *PhysicsNinja{
    return &PhysicsNinja{self.Object.Get("system")}
}

// A reference to the physics system.
func (self *PhysicsNinjaAABB) SetSystemA(member *PhysicsNinja) {
    self.Object.Set("system", member)
}

// The position of this object.
func (self *PhysicsNinjaAABB) GetPosA() *Point{
    return &Point{self.Object.Get("pos")}
}

// The position of this object.
func (self *PhysicsNinjaAABB) SetPosA(member *Point) {
    self.Object.Set("pos", member)
}

// The position of this object in the previous update.
func (self *PhysicsNinjaAABB) GetOldposA() *Point{
    return &Point{self.Object.Get("oldpos")}
}

// The position of this object in the previous update.
func (self *PhysicsNinjaAABB) SetOldposA(member *Point) {
    self.Object.Set("oldpos", member)
}

// Half the width.
func (self *PhysicsNinjaAABB) GetXwA() int{
    return self.Object.Get("xw").Int()
}

// Half the width.
func (self *PhysicsNinjaAABB) SetXwA(member int) {
    self.Object.Set("xw", member)
}

// Half the height.
func (self *PhysicsNinjaAABB) GetYwA() interface{}{
    return self.Object.Get("yw")
}

// Half the height.
func (self *PhysicsNinjaAABB) SetYwA(member interface{}) {
    self.Object.Set("yw", member)
}

// The width.
func (self *PhysicsNinjaAABB) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width.
func (self *PhysicsNinjaAABB) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height.
func (self *PhysicsNinjaAABB) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height.
func (self *PhysicsNinjaAABB) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The velocity of this object.
func (self *PhysicsNinjaAABB) GetVelocityA() *Point{
    return &Point{self.Object.Get("velocity")}
}

// The velocity of this object.
func (self *PhysicsNinjaAABB) SetVelocityA(member *Point) {
    self.Object.Set("velocity", member)
}

// All of the collision response handlers.
func (self *PhysicsNinjaAABB) GetAabbTileProjectionsA() interface{}{
    return self.Object.Get("aabbTileProjections")
}

// All of the collision response handlers.
func (self *PhysicsNinjaAABB) SetAabbTileProjectionsA(member interface{}) {
    self.Object.Set("aabbTileProjections", member)
}



// Updates this AABBs position.
func (self *PhysicsNinjaAABB) Integrate() {
    self.Object.Call("integrate")
}

// Updates this AABBs position.
func (self *PhysicsNinjaAABB) IntegrateI(args ...interface{}) {
    self.Object.Call("integrate", args)
}

// Process a collision partner-agnostic collision response and apply the resulting forces.
func (self *PhysicsNinjaAABB) ReportCollision(px int, py int, dx int, dy int) {
    self.Object.Call("reportCollision", px, py, dx, dy)
}

// Process a collision partner-agnostic collision response and apply the resulting forces.
func (self *PhysicsNinjaAABB) ReportCollisionI(args ...interface{}) {
    self.Object.Call("reportCollision", args)
}

// Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaAABB) ReportCollisionVsWorld(px int, py int, dx int, dy int) {
    self.Object.Call("reportCollisionVsWorld", px, py, dx, dy)
}

// Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaAABB) ReportCollisionVsWorldI(args ...interface{}) {
    self.Object.Call("reportCollisionVsWorld", args)
}

// 
func (self *PhysicsNinjaAABB) Reverse() {
    self.Object.Call("reverse")
}

// 
func (self *PhysicsNinjaAABB) ReverseI(args ...interface{}) {
    self.Object.Call("reverse", args)
}

// Process a body collision and apply the resulting forces. Still very much WIP and doesn't work fully. Feel free to fix!
func (self *PhysicsNinjaAABB) ReportCollisionVsBody(px int, py int, dx int, dy int, obj int) {
    self.Object.Call("reportCollisionVsBody", px, py, dx, dy, obj)
}

// Process a body collision and apply the resulting forces. Still very much WIP and doesn't work fully. Feel free to fix!
func (self *PhysicsNinjaAABB) ReportCollisionVsBodyI(args ...interface{}) {
    self.Object.Call("reportCollisionVsBody", args)
}

// Collides this AABB against the world bounds.
func (self *PhysicsNinjaAABB) CollideWorldBounds() {
    self.Object.Call("collideWorldBounds")
}

// Collides this AABB against the world bounds.
func (self *PhysicsNinjaAABB) CollideWorldBoundsI(args ...interface{}) {
    self.Object.Call("collideWorldBounds", args)
}

// Collides this AABB against a AABB.
func (self *PhysicsNinjaAABB) CollideAABBVsAABB(aabb *PhysicsNinjaAABB) {
    self.Object.Call("collideAABBVsAABB", aabb)
}

// Collides this AABB against a AABB.
func (self *PhysicsNinjaAABB) CollideAABBVsAABBI(args ...interface{}) {
    self.Object.Call("collideAABBVsAABB", args)
}

// Collides this AABB against a Tile.
func (self *PhysicsNinjaAABB) CollideAABBVsTile(tile *PhysicsNinjaTile) {
    self.Object.Call("collideAABBVsTile", tile)
}

// Collides this AABB against a Tile.
func (self *PhysicsNinjaAABB) CollideAABBVsTileI(args ...interface{}) {
    self.Object.Call("collideAABBVsTile", args)
}

// Resolves tile collision.
func (self *PhysicsNinjaAABB) ResolveTile(x int, y int, body *PhysicsNinjaAABB, tile *PhysicsNinjaTile) bool{
    return self.Object.Call("resolveTile", x, y, body, tile).Bool()
}

// Resolves tile collision.
func (self *PhysicsNinjaAABB) ResolveTileI(args ...interface{}) bool{
    return self.Object.Call("resolveTile", args).Bool()
}

// Resolves Full tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_Full(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_Full", x, y, obj, t).Int()
}

// Resolves Full tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_FullI(args ...interface{}) int{
    return self.Object.Call("projAABB_Full", args).Int()
}

// Resolves Half tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_Half(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_Half", x, y, obj, t).Int()
}

// Resolves Half tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_HalfI(args ...interface{}) int{
    return self.Object.Call("projAABB_Half", args).Int()
}

// Resolves 45 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_45Deg(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_45Deg", x, y, obj, t).Int()
}

// Resolves 45 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_45DegI(args ...interface{}) int{
    return self.Object.Call("projAABB_45Deg", args).Int()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_22DegS(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_22DegS", x, y, obj, t).Int()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_22DegSI(args ...interface{}) int{
    return self.Object.Call("projAABB_22DegS", args).Int()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_22DegB(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_22DegB", x, y, obj, t).Int()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_22DegBI(args ...interface{}) int{
    return self.Object.Call("projAABB_22DegB", args).Int()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_67DegS(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_67DegS", x, y, obj, t).Int()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_67DegSI(args ...interface{}) int{
    return self.Object.Call("projAABB_67DegS", args).Int()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_67DegB(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_67DegB", x, y, obj, t).Int()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_67DegBI(args ...interface{}) int{
    return self.Object.Call("projAABB_67DegB", args).Int()
}

// Resolves Convex tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_Convex(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_Convex", x, y, obj, t).Int()
}

// Resolves Convex tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_ConvexI(args ...interface{}) int{
    return self.Object.Call("projAABB_Convex", args).Int()
}

// Resolves Concave tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_Concave(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_Concave", x, y, obj, t).Int()
}

// Resolves Concave tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_ConcaveI(args ...interface{}) int{
    return self.Object.Call("projAABB_Concave", args).Int()
}

// Destroys this AABB's reference to Body and System
func (self *PhysicsNinjaAABB) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this AABB's reference to Body and System
func (self *PhysicsNinjaAABB) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Render this AABB for debugging purposes.
func (self *PhysicsNinjaAABB) Render(context interface{}, xOffset int, yOffset int, color string, filled bool) {
    self.Object.Call("render", context, xOffset, yOffset, color, filled)
}

// Render this AABB for debugging purposes.
func (self *PhysicsNinjaAABB) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}
