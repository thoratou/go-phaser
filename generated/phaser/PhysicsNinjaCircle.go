// Automatic generation for Phaser.Physics.Ninja.Circle
// generated file PhysicsNinjaCircle.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Ninja Physics Circle constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
type PhysicsNinjaCircle struct {
    *js.Object
}


// A reference to the body that owns this shape.
func (self *PhysicsNinjaCircle) GetBody() interface{}{
    return self.Get("body")
}

// A reference to the body that owns this shape.
func (self *PhysicsNinjaCircle) SetBody(member interface{}) {
    self.Set("body", member)
}

// A reference to the physics system.
func (self *PhysicsNinjaCircle) GetSystem() *PhysicsNinja{
    return &PhysicsNinja{self.Get("system")}
}

// A reference to the physics system.
func (self *PhysicsNinjaCircle) SetSystem(member *PhysicsNinja) {
    self.Set("system", member)
}

// The position of this object.
func (self *PhysicsNinjaCircle) GetPos() *Point{
    return &Point{self.Get("pos")}
}

// The position of this object.
func (self *PhysicsNinjaCircle) SetPos(member *Point) {
    self.Set("pos", member)
}

// The position of this object in the previous update.
func (self *PhysicsNinjaCircle) GetOldpos() *Point{
    return &Point{self.Get("oldpos")}
}

// The position of this object in the previous update.
func (self *PhysicsNinjaCircle) SetOldpos(member *Point) {
    self.Set("oldpos", member)
}

// The radius of this circle shape.
func (self *PhysicsNinjaCircle) GetRadius() float64{
    return self.Get("radius").Float()
}

// The radius of this circle shape.
func (self *PhysicsNinjaCircle) SetRadius(member float64) {
    self.Set("radius", member)
}

// Half the width.
func (self *PhysicsNinjaCircle) GetXw() float64{
    return self.Get("xw").Float()
}

// Half the width.
func (self *PhysicsNinjaCircle) SetXw(member float64) {
    self.Set("xw", member)
}

// Half the height.
func (self *PhysicsNinjaCircle) GetYw() interface{}{
    return self.Get("yw")
}

// Half the height.
func (self *PhysicsNinjaCircle) SetYw(member interface{}) {
    self.Set("yw", member)
}

// The width.
func (self *PhysicsNinjaCircle) GetWidth() float64{
    return self.Get("width").Float()
}

// The width.
func (self *PhysicsNinjaCircle) SetWidth(member float64) {
    self.Set("width", member)
}

// The height.
func (self *PhysicsNinjaCircle) GetHeight() float64{
    return self.Get("height").Float()
}

// The height.
func (self *PhysicsNinjaCircle) SetHeight(member float64) {
    self.Set("height", member)
}

// The velocity of this object.
func (self *PhysicsNinjaCircle) GetVelocity() *Point{
    return &Point{self.Get("velocity")}
}

// The velocity of this object.
func (self *PhysicsNinjaCircle) SetVelocity(member *Point) {
    self.Set("velocity", member)
}

// All of the collision response handlers.
func (self *PhysicsNinjaCircle) GetCircleTileProjections() interface{}{
    return self.Get("circleTileProjections")
}

// All of the collision response handlers.
func (self *PhysicsNinjaCircle) SetCircleTileProjections(member interface{}) {
    self.Set("circleTileProjections", member)
}



// Updates this Circles position.
func (self *PhysicsNinjaCircle) IntegrateI(args ...interface{}) {
    self.Call("integrate", args)
}

// Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaCircle) ReportCollisionVsWorldI(args ...interface{}) {
    self.Call("reportCollisionVsWorld", args)
}

// Collides this Circle against the world bounds.
func (self *PhysicsNinjaCircle) CollideWorldBoundsI(args ...interface{}) {
    self.Call("collideWorldBounds", args)
}

// Collides this Circle with a Tile.
func (self *PhysicsNinjaCircle) CollideCircleVsTileI(args ...interface{}) bool{
    return self.Call("collideCircleVsTile", args).Bool()
}

// Resolves tile collision.
func (self *PhysicsNinjaCircle) ResolveCircleTileI(args ...interface{}) float64{
    return self.Call("resolveCircleTile", args).Float()
}

// Resolves Full tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_FullI(args ...interface{}) float64{
    return self.Call("projCircle_Full", args).Float()
}

// Resolves 45 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_45DegI(args ...interface{}) float64{
    return self.Call("projCircle_45Deg", args).Float()
}

// Resolves Concave tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_ConcaveI(args ...interface{}) float64{
    return self.Call("projCircle_Concave", args).Float()
}

// Resolves Convex tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_ConvexI(args ...interface{}) float64{
    return self.Call("projCircle_Convex", args).Float()
}

// Resolves Half tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_HalfI(args ...interface{}) float64{
    return self.Call("projCircle_Half", args).Float()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_22DegSI(args ...interface{}) float64{
    return self.Call("projCircle_22DegS", args).Float()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_22DegBI(args ...interface{}) float64{
    return self.Call("projCircle_22DegB", args).Float()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_67DegSI(args ...interface{}) float64{
    return self.Call("projCircle_67DegS", args).Float()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_67DegBI(args ...interface{}) float64{
    return self.Call("projCircle_67DegB", args).Float()
}

// Destroys this Circle's reference to Body and System
func (self *PhysicsNinjaCircle) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Render this circle for debugging purposes.
func (self *PhysicsNinjaCircle) RenderI(args ...interface{}) {
    self.Call("render", args)
}
