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


// A reference to the body that owns this shape.
func (self *PhysicsNinjaAABB) GetBody() interface{}{
    return self.Get("body")
}

// A reference to the body that owns this shape.
func (self *PhysicsNinjaAABB) SetBody(member interface{}) {
    self.Set("body", member)
}

// A reference to the physics system.
func (self *PhysicsNinjaAABB) GetSystem() PhysicsNinja{
    return PhysicsNinja{self.Get("system")}
}

// A reference to the physics system.
func (self *PhysicsNinjaAABB) SetSystem(member PhysicsNinja) {
    self.Set("system", member)
}

// The position of this object.
func (self *PhysicsNinjaAABB) GetPos() Point{
    return Point{self.Get("pos")}
}

// The position of this object.
func (self *PhysicsNinjaAABB) SetPos(member Point) {
    self.Set("pos", member)
}

// The position of this object in the previous update.
func (self *PhysicsNinjaAABB) GetOldpos() Point{
    return Point{self.Get("oldpos")}
}

// The position of this object in the previous update.
func (self *PhysicsNinjaAABB) SetOldpos(member Point) {
    self.Set("oldpos", member)
}

// Half the width.
func (self *PhysicsNinjaAABB) GetXw() float64{
    return self.Get("xw").Float()
}

// Half the width.
func (self *PhysicsNinjaAABB) SetXw(member float64) {
    self.Set("xw", member)
}

// Half the height.
func (self *PhysicsNinjaAABB) GetYw() interface{}{
    return self.Get("yw")
}

// Half the height.
func (self *PhysicsNinjaAABB) SetYw(member interface{}) {
    self.Set("yw", member)
}

// The width.
func (self *PhysicsNinjaAABB) GetWidth() float64{
    return self.Get("width").Float()
}

// The width.
func (self *PhysicsNinjaAABB) SetWidth(member float64) {
    self.Set("width", member)
}

// The height.
func (self *PhysicsNinjaAABB) GetHeight() float64{
    return self.Get("height").Float()
}

// The height.
func (self *PhysicsNinjaAABB) SetHeight(member float64) {
    self.Set("height", member)
}

// The velocity of this object.
func (self *PhysicsNinjaAABB) GetVelocity() Point{
    return Point{self.Get("velocity")}
}

// The velocity of this object.
func (self *PhysicsNinjaAABB) SetVelocity(member Point) {
    self.Set("velocity", member)
}

// All of the collision response handlers.
func (self *PhysicsNinjaAABB) GetAabbTileProjections() interface{}{
    return self.Get("aabbTileProjections")
}

// All of the collision response handlers.
func (self *PhysicsNinjaAABB) SetAabbTileProjections(member interface{}) {
    self.Set("aabbTileProjections", member)
}



// Updates this AABBs position.
func (self *PhysicsNinjaAABB) IntegrateI(args ...interface{}) {
    self.Call("integrate", args)
}

// Process a collision partner-agnostic collision response and apply the resulting forces.
func (self *PhysicsNinjaAABB) ReportCollisionI(args ...interface{}) {
    self.Call("reportCollision", args)
}

// Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaAABB) ReportCollisionVsWorldI(args ...interface{}) {
    self.Call("reportCollisionVsWorld", args)
}

// 
func (self *PhysicsNinjaAABB) ReverseI(args ...interface{}) {
    self.Call("reverse", args)
}

// Process a body collision and apply the resulting forces. Still very much WIP and doesn't work fully. Feel free to fix!
func (self *PhysicsNinjaAABB) ReportCollisionVsBodyI(args ...interface{}) {
    self.Call("reportCollisionVsBody", args)
}

// Collides this AABB against the world bounds.
func (self *PhysicsNinjaAABB) CollideWorldBoundsI(args ...interface{}) {
    self.Call("collideWorldBounds", args)
}

// Collides this AABB against a AABB.
func (self *PhysicsNinjaAABB) CollideAABBVsAABBI(args ...interface{}) {
    self.Call("collideAABBVsAABB", args)
}

// Collides this AABB against a Tile.
func (self *PhysicsNinjaAABB) CollideAABBVsTileI(args ...interface{}) {
    self.Call("collideAABBVsTile", args)
}

// Resolves tile collision.
func (self *PhysicsNinjaAABB) ResolveTileI(args ...interface{}) bool{
    return self.Call("resolveTile", args).Bool()
}

// Resolves Full tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_FullI(args ...interface{}) float64{
    return self.Call("projAABB_Full", args).Float()
}

// Resolves Half tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_HalfI(args ...interface{}) float64{
    return self.Call("projAABB_Half", args).Float()
}

// Resolves 45 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_45DegI(args ...interface{}) float64{
    return self.Call("projAABB_45Deg", args).Float()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_22DegSI(args ...interface{}) float64{
    return self.Call("projAABB_22DegS", args).Float()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_22DegBI(args ...interface{}) float64{
    return self.Call("projAABB_22DegB", args).Float()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_67DegSI(args ...interface{}) float64{
    return self.Call("projAABB_67DegS", args).Float()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_67DegBI(args ...interface{}) float64{
    return self.Call("projAABB_67DegB", args).Float()
}

// Resolves Convex tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_ConvexI(args ...interface{}) float64{
    return self.Call("projAABB_Convex", args).Float()
}

// Resolves Concave tile collision.
func (self *PhysicsNinjaAABB) ProjAABB_ConcaveI(args ...interface{}) float64{
    return self.Call("projAABB_Concave", args).Float()
}

// Destroys this AABB's reference to Body and System
func (self *PhysicsNinjaAABB) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Render this AABB for debugging purposes.
func (self *PhysicsNinjaAABB) RenderI(args ...interface{}) {
    self.Call("render", args)
}
