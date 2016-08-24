// Package phaser Automatic generation for Phaser.Physics.Ninja.AABB
// generated file PhysicsNinjaAABB.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsNinjaAABB Ninja Physics AABB constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
type PhysicsNinjaAABB struct {
    *js.Object
}

// NewPhysicsNinjaAABB Ninja Physics AABB constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaAABB(body *PhysicsNinjaBody, x int, y int, width int, height int) *PhysicsNinjaAABB {
    return &PhysicsNinjaAABB{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("AABB").New(body, x, y, width, height)}
}
// NewPhysicsNinjaAABBI Ninja Physics AABB constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaAABBI(args ...interface{}) *PhysicsNinjaAABB {
    return &PhysicsNinjaAABB{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("AABB").New(args)}
}



// Body A reference to the body that owns this shape.
func (self *PhysicsNinjaAABB) Body() interface{}{
    return self.Object.Get("body")
}

// SetBodyA A reference to the body that owns this shape.
func (self *PhysicsNinjaAABB) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// System A reference to the physics system.
func (self *PhysicsNinjaAABB) System() *PhysicsNinja{
    return &PhysicsNinja{self.Object.Get("system")}
}

// SetSystemA A reference to the physics system.
func (self *PhysicsNinjaAABB) SetSystemA(member *PhysicsNinja) {
    self.Object.Set("system", member)
}

// Pos The position of this object.
func (self *PhysicsNinjaAABB) Pos() *Point{
    return &Point{self.Object.Get("pos")}
}

// SetPosA The position of this object.
func (self *PhysicsNinjaAABB) SetPosA(member *Point) {
    self.Object.Set("pos", member)
}

// Oldpos The position of this object in the previous update.
func (self *PhysicsNinjaAABB) Oldpos() *Point{
    return &Point{self.Object.Get("oldpos")}
}

// SetOldposA The position of this object in the previous update.
func (self *PhysicsNinjaAABB) SetOldposA(member *Point) {
    self.Object.Set("oldpos", member)
}

// Xw Half the width.
func (self *PhysicsNinjaAABB) Xw() int{
    return self.Object.Get("xw").Int()
}

// SetXwA Half the width.
func (self *PhysicsNinjaAABB) SetXwA(member int) {
    self.Object.Set("xw", member)
}

// Yw Half the height.
func (self *PhysicsNinjaAABB) Yw() interface{}{
    return self.Object.Get("yw")
}

// SetYwA Half the height.
func (self *PhysicsNinjaAABB) SetYwA(member interface{}) {
    self.Object.Set("yw", member)
}

// Width The width.
func (self *PhysicsNinjaAABB) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width.
func (self *PhysicsNinjaAABB) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height.
func (self *PhysicsNinjaAABB) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height.
func (self *PhysicsNinjaAABB) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Velocity The velocity of this object.
func (self *PhysicsNinjaAABB) Velocity() *Point{
    return &Point{self.Object.Get("velocity")}
}

// SetVelocityA The velocity of this object.
func (self *PhysicsNinjaAABB) SetVelocityA(member *Point) {
    self.Object.Set("velocity", member)
}

// AabbTileProjections All of the collision response handlers.
func (self *PhysicsNinjaAABB) AabbTileProjections() interface{}{
    return self.Object.Get("aabbTileProjections")
}

// SetAabbTileProjectionsA All of the collision response handlers.
func (self *PhysicsNinjaAABB) SetAabbTileProjectionsA(member interface{}) {
    self.Object.Set("aabbTileProjections", member)
}


// Integrate Updates this AABBs position.
func (self *PhysicsNinjaAABB) Integrate() {
    self.Object.Call("integrate")
}

// IntegrateI Updates this AABBs position.
func (self *PhysicsNinjaAABB) IntegrateI(args ...interface{}) {
    self.Object.Call("integrate", args)
}

// ReportCollision Process a collision partner-agnostic collision response and apply the resulting forces.
func (self *PhysicsNinjaAABB) ReportCollision(px int, py int, dx int, dy int) {
    self.Object.Call("reportCollision", px, py, dx, dy)
}

// ReportCollisionI Process a collision partner-agnostic collision response and apply the resulting forces.
func (self *PhysicsNinjaAABB) ReportCollisionI(args ...interface{}) {
    self.Object.Call("reportCollision", args)
}

// ReportCollisionVsWorld Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaAABB) ReportCollisionVsWorld(px int, py int, dx int, dy int) {
    self.Object.Call("reportCollisionVsWorld", px, py, dx, dy)
}

// ReportCollisionVsWorldI Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaAABB) ReportCollisionVsWorldI(args ...interface{}) {
    self.Object.Call("reportCollisionVsWorld", args)
}

// Reverse empty description
func (self *PhysicsNinjaAABB) Reverse() {
    self.Object.Call("reverse")
}

// ReverseI empty description
func (self *PhysicsNinjaAABB) ReverseI(args ...interface{}) {
    self.Object.Call("reverse", args)
}

// ReportCollisionVsBody Process a body collision and apply the resulting forces. Still very much WIP and doesn't work fully. Feel free to fix!
func (self *PhysicsNinjaAABB) ReportCollisionVsBody(px int, py int, dx int, dy int, obj int) {
    self.Object.Call("reportCollisionVsBody", px, py, dx, dy, obj)
}

// ReportCollisionVsBodyI Process a body collision and apply the resulting forces. Still very much WIP and doesn't work fully. Feel free to fix!
func (self *PhysicsNinjaAABB) ReportCollisionVsBodyI(args ...interface{}) {
    self.Object.Call("reportCollisionVsBody", args)
}

// CollideWorldBounds Collides this AABB against the world bounds.
func (self *PhysicsNinjaAABB) CollideWorldBounds() {
    self.Object.Call("collideWorldBounds")
}

// CollideWorldBoundsI Collides this AABB against the world bounds.
func (self *PhysicsNinjaAABB) CollideWorldBoundsI(args ...interface{}) {
    self.Object.Call("collideWorldBounds", args)
}

// CollideAABBVsAABB Collides this AABB against a AABB.
func (self *PhysicsNinjaAABB) CollideAABBVsAABB(aabb *PhysicsNinjaAABB) {
    self.Object.Call("collideAABBVsAABB", aabb)
}

// CollideAABBVsAABBI Collides this AABB against a AABB.
func (self *PhysicsNinjaAABB) CollideAABBVsAABBI(args ...interface{}) {
    self.Object.Call("collideAABBVsAABB", args)
}

// CollideAABBVsTile Collides this AABB against a Tile.
func (self *PhysicsNinjaAABB) CollideAABBVsTile(tile *PhysicsNinjaTile) {
    self.Object.Call("collideAABBVsTile", tile)
}

// CollideAABBVsTileI Collides this AABB against a Tile.
func (self *PhysicsNinjaAABB) CollideAABBVsTileI(args ...interface{}) {
    self.Object.Call("collideAABBVsTile", args)
}

// ResolveTile Resolves tile collision.
func (self *PhysicsNinjaAABB) ResolveTile(x int, y int, body *PhysicsNinjaAABB, tile *PhysicsNinjaTile) bool{
    return self.Object.Call("resolveTile", x, y, body, tile).Bool()
}

// ResolveTileI Resolves tile collision.
func (self *PhysicsNinjaAABB) ResolveTileI(args ...interface{}) bool{
    return self.Object.Call("resolveTile", args).Bool()
}

// ProjAABB_Full Resolves Full tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_Full(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_Full", x, y, obj, t).Int()
}

// ProjAABB_FullI Resolves Full tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_FullI(args ...interface{}) int{
    return self.Object.Call("projAABB_Full", args).Int()
}

// ProjAABB_Half Resolves Half tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_Half(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_Half", x, y, obj, t).Int()
}

// ProjAABB_HalfI Resolves Half tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_HalfI(args ...interface{}) int{
    return self.Object.Call("projAABB_Half", args).Int()
}

// ProjAABB_45Deg Resolves 45 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_45Deg(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_45Deg", x, y, obj, t).Int()
}

// ProjAABB_45DegI Resolves 45 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_45DegI(args ...interface{}) int{
    return self.Object.Call("projAABB_45Deg", args).Int()
}

// ProjAABB_22DegS Resolves 22 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_22DegS(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_22DegS", x, y, obj, t).Int()
}

// ProjAABB_22DegSI Resolves 22 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_22DegSI(args ...interface{}) int{
    return self.Object.Call("projAABB_22DegS", args).Int()
}

// ProjAABB_22DegB Resolves 22 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_22DegB(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_22DegB", x, y, obj, t).Int()
}

// ProjAABB_22DegBI Resolves 22 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_22DegBI(args ...interface{}) int{
    return self.Object.Call("projAABB_22DegB", args).Int()
}

// ProjAABB_67DegS Resolves 67 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_67DegS(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_67DegS", x, y, obj, t).Int()
}

// ProjAABB_67DegSI Resolves 67 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_67DegSI(args ...interface{}) int{
    return self.Object.Call("projAABB_67DegS", args).Int()
}

// ProjAABB_67DegB Resolves 67 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_67DegB(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_67DegB", x, y, obj, t).Int()
}

// ProjAABB_67DegBI Resolves 67 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_67DegBI(args ...interface{}) int{
    return self.Object.Call("projAABB_67DegB", args).Int()
}

// ProjAABB_Convex Resolves Convex tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_Convex(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_Convex", x, y, obj, t).Int()
}

// ProjAABB_ConvexI Resolves Convex tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_ConvexI(args ...interface{}) int{
    return self.Object.Call("projAABB_Convex", args).Int()
}

// ProjAABB_Concave Resolves Concave tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_Concave(x int, y int, obj *PhysicsNinjaAABB, t *PhysicsNinjaTile) int{
    return self.Object.Call("projAABB_Concave", x, y, obj, t).Int()
}

// ProjAABB_ConcaveI Resolves Concave tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_ConcaveI(args ...interface{}) int{
    return self.Object.Call("projAABB_Concave", args).Int()
}

// Destroy Destroys this AABB's reference to Body and System
func (self *PhysicsNinjaAABB) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this AABB's reference to Body and System
func (self *PhysicsNinjaAABB) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Render Render this AABB for debugging purposes.
func (self *PhysicsNinjaAABB) Render(context interface{}, xOffset int, yOffset int, color string, filled bool) {
    self.Object.Call("render", context, xOffset, yOffset, color, filled)
}

// RenderI Render this AABB for debugging purposes.
func (self *PhysicsNinjaAABB) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

