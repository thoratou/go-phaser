// Package phaser Automatic generation for Phaser.Physics.Ninja.Circle
// generated file PhysicsNinjaCircle.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsNinjaCircle Ninja Physics Circle constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
type PhysicsNinjaCircle struct {
    *js.Object
}

// NewPhysicsNinjaCircle Ninja Physics Circle constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaCircle(body *PhysicsNinjaBody, x int, y int, radius int) *PhysicsNinjaCircle {
    return &PhysicsNinjaCircle{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Circle").New(body, x, y, radius)}
}
// NewPhysicsNinjaCircleI Ninja Physics Circle constructor.
// Note: This class could be massively optimised and reduced in size. I leave that challenge up to you.
func NewPhysicsNinjaCircleI(args ...interface{}) *PhysicsNinjaCircle {
    return &PhysicsNinjaCircle{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Circle").New(args)}
}



// PhysicsNinjaCircle Binding conversion method to PhysicsNinjaCircle point 
func ToPhysicsNinjaCircle(jsStruct interface{}) *PhysicsNinjaCircle {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsNinjaCircle{Object: object}
	}
	return nil
}



// Body A reference to the body that owns this shape.
func (self *PhysicsNinjaCircle) Body() interface{}{
    return self.Object.Get("body")
}

// SetBodyA A reference to the body that owns this shape.
func (self *PhysicsNinjaCircle) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// System A reference to the physics system.
func (self *PhysicsNinjaCircle) System() *PhysicsNinja{
    return &PhysicsNinja{self.Object.Get("system")}
}

// SetSystemA A reference to the physics system.
func (self *PhysicsNinjaCircle) SetSystemA(member *PhysicsNinja) {
    self.Object.Set("system", member)
}

// Pos The position of this object.
func (self *PhysicsNinjaCircle) Pos() *Point{
    return &Point{self.Object.Get("pos")}
}

// SetPosA The position of this object.
func (self *PhysicsNinjaCircle) SetPosA(member *Point) {
    self.Object.Set("pos", member)
}

// Oldpos The position of this object in the previous update.
func (self *PhysicsNinjaCircle) Oldpos() *Point{
    return &Point{self.Object.Get("oldpos")}
}

// SetOldposA The position of this object in the previous update.
func (self *PhysicsNinjaCircle) SetOldposA(member *Point) {
    self.Object.Set("oldpos", member)
}

// Radius The radius of this circle shape.
func (self *PhysicsNinjaCircle) Radius() int{
    return self.Object.Get("radius").Int()
}

// SetRadiusA The radius of this circle shape.
func (self *PhysicsNinjaCircle) SetRadiusA(member int) {
    self.Object.Set("radius", member)
}

// Xw Half the width.
func (self *PhysicsNinjaCircle) Xw() int{
    return self.Object.Get("xw").Int()
}

// SetXwA Half the width.
func (self *PhysicsNinjaCircle) SetXwA(member int) {
    self.Object.Set("xw", member)
}

// Yw Half the height.
func (self *PhysicsNinjaCircle) Yw() interface{}{
    return self.Object.Get("yw")
}

// SetYwA Half the height.
func (self *PhysicsNinjaCircle) SetYwA(member interface{}) {
    self.Object.Set("yw", member)
}

// Width The width.
func (self *PhysicsNinjaCircle) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width.
func (self *PhysicsNinjaCircle) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height.
func (self *PhysicsNinjaCircle) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height.
func (self *PhysicsNinjaCircle) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Velocity The velocity of this object.
func (self *PhysicsNinjaCircle) Velocity() *Point{
    return &Point{self.Object.Get("velocity")}
}

// SetVelocityA The velocity of this object.
func (self *PhysicsNinjaCircle) SetVelocityA(member *Point) {
    self.Object.Set("velocity", member)
}

// CircleTileProjections All of the collision response handlers.
func (self *PhysicsNinjaCircle) CircleTileProjections() interface{}{
    return self.Object.Get("circleTileProjections")
}

// SetCircleTileProjectionsA All of the collision response handlers.
func (self *PhysicsNinjaCircle) SetCircleTileProjectionsA(member interface{}) {
    self.Object.Set("circleTileProjections", member)
}


// Integrate Updates this Circles position.
func (self *PhysicsNinjaCircle) Integrate() {
    self.Object.Call("integrate")
}

// IntegrateI Updates this Circles position.
func (self *PhysicsNinjaCircle) IntegrateI(args ...interface{}) {
    self.Object.Call("integrate", args)
}

// ReportCollisionVsWorld Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaCircle) ReportCollisionVsWorld(px int, py int, dx int, dy int, obj int) {
    self.Object.Call("reportCollisionVsWorld", px, py, dx, dy, obj)
}

// ReportCollisionVsWorldI Process a world collision and apply the resulting forces.
func (self *PhysicsNinjaCircle) ReportCollisionVsWorldI(args ...interface{}) {
    self.Object.Call("reportCollisionVsWorld", args)
}

// CollideWorldBounds Collides this Circle against the world bounds.
func (self *PhysicsNinjaCircle) CollideWorldBounds() {
    self.Object.Call("collideWorldBounds")
}

// CollideWorldBoundsI Collides this Circle against the world bounds.
func (self *PhysicsNinjaCircle) CollideWorldBoundsI(args ...interface{}) {
    self.Object.Call("collideWorldBounds", args)
}

// CollideCircleVsTile Collides this Circle with a Tile.
func (self *PhysicsNinjaCircle) CollideCircleVsTile(t *PhysicsNinjaTile) bool{
    return self.Object.Call("collideCircleVsTile", t).Bool()
}

// CollideCircleVsTileI Collides this Circle with a Tile.
func (self *PhysicsNinjaCircle) CollideCircleVsTileI(args ...interface{}) bool{
    return self.Object.Call("collideCircleVsTile", args).Bool()
}

// ResolveCircleTile Resolves tile collision.
func (self *PhysicsNinjaCircle) ResolveCircleTile(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("resolveCircleTile", x, y, oH, oV, obj, t).Int()
}

// ResolveCircleTileI Resolves tile collision.
func (self *PhysicsNinjaCircle) ResolveCircleTileI(args ...interface{}) int{
    return self.Object.Call("resolveCircleTile", args).Int()
}

// ProjCircle_Full Resolves Full tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_Full(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_Full", x, y, oH, oV, obj, t).Int()
}

// ProjCircle_FullI Resolves Full tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_FullI(args ...interface{}) int{
    return self.Object.Call("projCircle_Full", args).Int()
}

// ProjCircle_45Deg Resolves 45 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_45Deg(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_45Deg", x, y, oH, oV, obj, t).Int()
}

// ProjCircle_45DegI Resolves 45 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_45DegI(args ...interface{}) int{
    return self.Object.Call("projCircle_45Deg", args).Int()
}

// ProjCircle_Concave Resolves Concave tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_Concave(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_Concave", x, y, oH, oV, obj, t).Int()
}

// ProjCircle_ConcaveI Resolves Concave tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_ConcaveI(args ...interface{}) int{
    return self.Object.Call("projCircle_Concave", args).Int()
}

// ProjCircle_Convex Resolves Convex tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_Convex(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_Convex", x, y, oH, oV, obj, t).Int()
}

// ProjCircle_ConvexI Resolves Convex tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_ConvexI(args ...interface{}) int{
    return self.Object.Call("projCircle_Convex", args).Int()
}

// ProjCircle_Half Resolves Half tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_Half(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_Half", x, y, oH, oV, obj, t).Int()
}

// ProjCircle_HalfI Resolves Half tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_HalfI(args ...interface{}) int{
    return self.Object.Call("projCircle_Half", args).Int()
}

// ProjCircle_22DegS Resolves 22 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_22DegS(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_22DegS", x, y, oH, oV, obj, t).Int()
}

// ProjCircle_22DegSI Resolves 22 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_22DegSI(args ...interface{}) int{
    return self.Object.Call("projCircle_22DegS", args).Int()
}

// ProjCircle_22DegB Resolves 22 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_22DegB(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_22DegB", x, y, oH, oV, obj, t).Int()
}

// ProjCircle_22DegBI Resolves 22 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_22DegBI(args ...interface{}) int{
    return self.Object.Call("projCircle_22DegB", args).Int()
}

// ProjCircle_67DegS Resolves 67 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_67DegS(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_67DegS", x, y, oH, oV, obj, t).Int()
}

// ProjCircle_67DegSI Resolves 67 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_67DegSI(args ...interface{}) int{
    return self.Object.Call("projCircle_67DegS", args).Int()
}

// ProjCircle_67DegB Resolves 67 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_67DegB(x int, y int, oH int, oV int, obj *PhysicsNinjaCircle, t *PhysicsNinjaTile) int{
    return self.Object.Call("projCircle_67DegB", x, y, oH, oV, obj, t).Int()
}

// ProjCircle_67DegBI Resolves 67 Degree tile collision.
func (self *PhysicsNinjaCircle) ProjCircle_67DegBI(args ...interface{}) int{
    return self.Object.Call("projCircle_67DegB", args).Int()
}

// Destroy Destroys this Circle's reference to Body and System
func (self *PhysicsNinjaCircle) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this Circle's reference to Body and System
func (self *PhysicsNinjaCircle) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Render Render this circle for debugging purposes.
func (self *PhysicsNinjaCircle) Render(context interface{}, xOffset int, yOffset int, color string, filled bool) {
    self.Object.Call("render", context, xOffset, yOffset, color, filled)
}

// RenderI Render this circle for debugging purposes.
func (self *PhysicsNinjaCircle) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

