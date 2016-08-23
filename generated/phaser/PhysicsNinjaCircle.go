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


// Ninja Physics Circle constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaCircle(body *PhysicsNinjaBody, x int, y int, radius int) *PhysicsNinjaCircle {
    return &PhysicsNinjaCircle{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Circle").New(body, x, y, radius)}
}

// Ninja Physics Circle constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaCircleI(args ...interface{}) *PhysicsNinjaCircle {
    return &PhysicsNinjaCircle{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Circle").New(args)}
}



// A reference to the body that owns this shape.
func (self *PhysicsNinjaCircle) Body() interface{}{
    return self.Object.Get("body")
}

// A reference to the body that owns this shape.
func (self *PhysicsNinjaCircle) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// A reference to the physics system.
func (self *PhysicsNinjaCircle) System() *PhysicsNinja{
    return &PhysicsNinja{self.Object.Get("system")}
}

// A reference to the physics system.
func (self *PhysicsNinjaCircle) SetSystemA(member *PhysicsNinja) {
    self.Object.Set("system", member)
}

// The position of this object.
func (self *PhysicsNinjaCircle) Pos() *Point{
    return &Point{self.Object.Get("pos")}
}

// The position of this object.
func (self *PhysicsNinjaCircle) SetPosA(member *Point) {
    self.Object.Set("pos", member)
}

// The position of this object in the previous update.
func (self *PhysicsNinjaCircle) Oldpos() *Point{
    return &Point{self.Object.Get("oldpos")}
}

// The position of this object in the previous update.
func (self *PhysicsNinjaCircle) SetOldposA(member *Point) {
    self.Object.Set("oldpos", member)
}

// The radius of this circle shape.
func (self *PhysicsNinjaCircle) Radius() int{
    return self.Object.Get("radius").Int()
}

// The radius of this circle shape.
func (self *PhysicsNinjaCircle) SetRadiusA(member int) {
    self.Object.Set("radius", member)
}

// Half the width.
func (self *PhysicsNinjaCircle) Xw() int{
    return self.Object.Get("xw").Int()
}

// Half the width.
func (self *PhysicsNinjaCircle) SetXwA(member int) {
    self.Object.Set("xw", member)
}

// Half the height.
func (self *PhysicsNinjaCircle) Yw() interface{}{
    return self.Object.Get("yw")
}

// Half the height.
func (self *PhysicsNinjaCircle) SetYwA(member interface{}) {
    self.Object.Set("yw", member)
}

// The width.
func (self *PhysicsNinjaCircle) Width() int{
    return self.Object.Get("width").Int()
}

// The width.
func (self *PhysicsNinjaCircle) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height.
func (self *PhysicsNinjaCircle) Height() int{
    return self.Object.Get("height").Int()
}

// The height.
func (self *PhysicsNinjaCircle) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The velocity of this object.
func (self *PhysicsNinjaCircle) Velocity() *Point{
    return &Point{self.Object.Get("velocity")}
}

// The velocity of this object.
func (self *PhysicsNinjaCircle) SetVelocityA(member *Point) {
    self.Object.Set("velocity", member)
}

// All of the collision response handlers.
func (self *PhysicsNinjaCircle) CircleTileProjections() interface{}{
    return self.Object.Get("circleTileProjections")
}

// All of the collision response handlers.
func (self *PhysicsNinjaCircle) SetCircleTileProjectionsA(member interface{}) {
    self.Object.Set("circleTileProjections", member)
}



// Updates this Circles position.
func (self *PhysicsNinjaCircle) Integrate() {
    self.Object.Call("integrate")
}

// Updates this Circles position.
func (self *PhysicsNinjaCircle) IntegrateI(args ...interface{}) {
    self.Object.Call("integrate", args)
}

// Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaCircle) ReportCollisionVsWorld(px int, py int, dx int, dy int, obj int) {
    self.Object.Call("reportCollisionVsWorld", px, py, dx, dy, obj)
}

// Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaCircle) ReportCollisionVsWorldI(args ...interface{}) {
    self.Object.Call("reportCollisionVsWorld", args)
}

// Collides this Circle against the world bounds.
func (self *PhysicsNinjaCircle) CollideWorldBounds() {
    self.Object.Call("collideWorldBounds")
}

// Collides this Circle against the world bounds.
func (self *PhysicsNinjaCircle) CollideWorldBoundsI(args ...interface{}) {
    self.Object.Call("collideWorldBounds", args)
}

// Collides this Circle with a Tile.
func (self *PhysicsNinjaCircle) CollideCircleVsTile(t *PhysicsNinjaTile) bool{
    return self.Object.Call("collideCircleVsTile", t).Bool()
}

// Collides this Circle with a Tile.
func (self *PhysicsNinjaCircle) CollideCircleVsTileI(args ...interface{}) bool{
    return self.Object.Call("collideCircleVsTile", args).Bool()
}

// Resolves tile collision.
func (self *PhysicsNinjaCircle) ResolveCircleTile(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("resolveCircleTile", x, y, oH, oV, obj, t).Int()
}

// Resolves tile collision.
func (self *PhysicsNinjaCircle) ResolveCircleTileI(args ...interface{}) int{
    return self.Object.Call("resolveCircleTile", args).Int()
}

// Resolves Full tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_Full(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_Full", x, y, oH, oV, obj, t).Int()
}

// Resolves Full tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_FullI(args ...interface{}) int{
    return self.Object.Call("projCircle_Full", args).Int()
}

// Resolves 45 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_45Deg(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_45Deg", x, y, oH, oV, obj, t).Int()
}

// Resolves 45 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_45DegI(args ...interface{}) int{
    return self.Object.Call("projCircle_45Deg", args).Int()
}

// Resolves Concave tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_Concave(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_Concave", x, y, oH, oV, obj, t).Int()
}

// Resolves Concave tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_ConcaveI(args ...interface{}) int{
    return self.Object.Call("projCircle_Concave", args).Int()
}

// Resolves Convex tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_Convex(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_Convex", x, y, oH, oV, obj, t).Int()
}

// Resolves Convex tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_ConvexI(args ...interface{}) int{
    return self.Object.Call("projCircle_Convex", args).Int()
}

// Resolves Half tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_Half(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_Half", x, y, oH, oV, obj, t).Int()
}

// Resolves Half tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_HalfI(args ...interface{}) int{
    return self.Object.Call("projCircle_Half", args).Int()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_22DegS(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_22DegS", x, y, oH, oV, obj, t).Int()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_22DegSI(args ...interface{}) int{
    return self.Object.Call("projCircle_22DegS", args).Int()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_22DegB(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_22DegB", x, y, oH, oV, obj, t).Int()
}

// Resolves 22 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_22DegBI(args ...interface{}) int{
    return self.Object.Call("projCircle_22DegB", args).Int()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_67DegS(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_67DegS", x, y, oH, oV, obj, t).Int()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_67DegSI(args ...interface{}) int{
    return self.Object.Call("projCircle_67DegS", args).Int()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_67DegB(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_67DegB", x, y, oH, oV, obj, t).Int()
}

// Resolves 67 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_67DegBI(args ...interface{}) int{
    return self.Object.Call("projCircle_67DegB", args).Int()
}

// Destroys this Circle's reference to Body and System
func (self *PhysicsNinjaCircle) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this Circle's reference to Body and System
func (self *PhysicsNinjaCircle) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Render this circle for debugging purposes.
func (self *PhysicsNinjaCircle) Render(context interface{}, xOffset int, yOffset int, color string, filled bool) {
    self.Object.Call("render", context, xOffset, yOffset, color, filled)
}

// Render this circle for debugging purposes.
func (self *PhysicsNinjaCircle) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}
