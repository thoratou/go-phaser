// Package phaser Automatic generation for Phaser.Physics.Ninja.Body
// generated file PhysicsNinjaBody.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsNinjaBody The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
type PhysicsNinjaBody struct {
    *js.Object
}

// NewPhysicsNinjaBody The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody(system *PhysicsNinja, sprite *Sprite) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Body").New(system, sprite)}
}
// NewPhysicsNinjaBody1O The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody1O(system *PhysicsNinja, sprite *Sprite, type_ int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Body").New(system, sprite, type_)}
}
// NewPhysicsNinjaBody2O The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody2O(system *PhysicsNinja, sprite *Sprite, type_ int, id int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Body").New(system, sprite, type_, id)}
}
// NewPhysicsNinjaBody3O The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody3O(system *PhysicsNinja, sprite *Sprite, type_ int, id int, radius int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Body").New(system, sprite, type_, id, radius)}
}
// NewPhysicsNinjaBody4O The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody4O(system *PhysicsNinja, sprite *Sprite, type_ int, id int, radius int, x int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Body").New(system, sprite, type_, id, radius, x)}
}
// NewPhysicsNinjaBody5O The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody5O(system *PhysicsNinja, sprite *Sprite, type_ int, id int, radius int, x int, y int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Body").New(system, sprite, type_, id, radius, x, y)}
}
// NewPhysicsNinjaBody6O The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody6O(system *PhysicsNinja, sprite *Sprite, type_ int, id int, radius int, x int, y int, width int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Body").New(system, sprite, type_, id, radius, x, y, width)}
}
// NewPhysicsNinjaBody7O The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody7O(system *PhysicsNinja, sprite *Sprite, type_ int, id int, radius int, x int, y int, width int, height int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Body").New(system, sprite, type_, id, radius, x, y, width, height)}
}
// NewPhysicsNinjaBodyI The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBodyI(args ...interface{}) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Get("Phaser").Get("Physics").Get("Ninja").Get("Body").New(args)}
}



// Sprite Reference to the parent Sprite.
func (self *PhysicsNinjaBody) Sprite() *Sprite{
    return &Sprite{self.Object.Get("sprite")}
}

// SetSpriteA Reference to the parent Sprite.
func (self *PhysicsNinjaBody) SetSpriteA(member *Sprite) {
    self.Object.Set("sprite", member)
}

// Game Local reference to game.
func (self *PhysicsNinjaBody) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *PhysicsNinjaBody) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Type The type of physics system this body belongs to.
func (self *PhysicsNinjaBody) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The type of physics system this body belongs to.
func (self *PhysicsNinjaBody) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// System The parent physics system.
func (self *PhysicsNinjaBody) System() *PhysicsNinja{
    return &PhysicsNinja{self.Object.Get("system")}
}

// SetSystemA The parent physics system.
func (self *PhysicsNinjaBody) SetSystemA(member *PhysicsNinja) {
    self.Object.Set("system", member)
}

// Aabb The AABB object this body is using for collision.
func (self *PhysicsNinjaBody) Aabb() *PhysicsNinjaAABB{
    return &PhysicsNinjaAABB{self.Object.Get("aabb")}
}

// SetAabbA The AABB object this body is using for collision.
func (self *PhysicsNinjaBody) SetAabbA(member *PhysicsNinjaAABB) {
    self.Object.Set("aabb", member)
}

// Tile The Tile object this body is using for collision.
func (self *PhysicsNinjaBody) Tile() *PhysicsNinjaTile{
    return &PhysicsNinjaTile{self.Object.Get("tile")}
}

// SetTileA The Tile object this body is using for collision.
func (self *PhysicsNinjaBody) SetTileA(member *PhysicsNinjaTile) {
    self.Object.Set("tile", member)
}

// Circle The Circle object this body is using for collision.
func (self *PhysicsNinjaBody) Circle() *PhysicsNinjaCircle{
    return &PhysicsNinjaCircle{self.Object.Get("circle")}
}

// SetCircleA The Circle object this body is using for collision.
func (self *PhysicsNinjaBody) SetCircleA(member *PhysicsNinjaCircle) {
    self.Object.Set("circle", member)
}

// Shape A local reference to the body shape.
func (self *PhysicsNinjaBody) Shape() interface{}{
    return self.Object.Get("shape")
}

// SetShapeA A local reference to the body shape.
func (self *PhysicsNinjaBody) SetShapeA(member interface{}) {
    self.Object.Set("shape", member)
}

// Drag The drag applied to this object as it moves.
func (self *PhysicsNinjaBody) Drag() int{
    return self.Object.Get("drag").Int()
}

// SetDragA The drag applied to this object as it moves.
func (self *PhysicsNinjaBody) SetDragA(member int) {
    self.Object.Set("drag", member)
}

// Friction The friction applied to this object as it moves.
func (self *PhysicsNinjaBody) Friction() int{
    return self.Object.Get("friction").Int()
}

// SetFrictionA The friction applied to this object as it moves.
func (self *PhysicsNinjaBody) SetFrictionA(member int) {
    self.Object.Set("friction", member)
}

// GravityScale How much of the world gravity should be applied to this object? 1 = all of it, 0.5 = 50%, etc.
func (self *PhysicsNinjaBody) GravityScale() int{
    return self.Object.Get("gravityScale").Int()
}

// SetGravityScaleA How much of the world gravity should be applied to this object? 1 = all of it, 0.5 = 50%, etc.
func (self *PhysicsNinjaBody) SetGravityScaleA(member int) {
    self.Object.Set("gravityScale", member)
}

// Bounce The bounciness of this object when it collides. A value between 0 and 1. We recommend setting it to 0.999 to avoid jittering.
func (self *PhysicsNinjaBody) Bounce() int{
    return self.Object.Get("bounce").Int()
}

// SetBounceA The bounciness of this object when it collides. A value between 0 and 1. We recommend setting it to 0.999 to avoid jittering.
func (self *PhysicsNinjaBody) SetBounceA(member int) {
    self.Object.Set("bounce", member)
}

// Velocity The velocity in pixels per second sq. of the Body.
func (self *PhysicsNinjaBody) Velocity() *Point{
    return &Point{self.Object.Get("velocity")}
}

// SetVelocityA The velocity in pixels per second sq. of the Body.
func (self *PhysicsNinjaBody) SetVelocityA(member *Point) {
    self.Object.Set("velocity", member)
}

// Facing A const reference to the direction the Body is traveling or facing.
func (self *PhysicsNinjaBody) Facing() int{
    return self.Object.Get("facing").Int()
}

// SetFacingA A const reference to the direction the Body is traveling or facing.
func (self *PhysicsNinjaBody) SetFacingA(member int) {
    self.Object.Set("facing", member)
}

// Immovable An immovable Body will not receive any impacts from other bodies. Not fully implemented.
func (self *PhysicsNinjaBody) Immovable() bool{
    return self.Object.Get("immovable").Bool()
}

// SetImmovableA An immovable Body will not receive any impacts from other bodies. Not fully implemented.
func (self *PhysicsNinjaBody) SetImmovableA(member bool) {
    self.Object.Set("immovable", member)
}

// CollideWorldBounds A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsNinjaBody) CollideWorldBounds() bool{
    return self.Object.Get("collideWorldBounds").Bool()
}

// SetCollideWorldBoundsA A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsNinjaBody) SetCollideWorldBoundsA(member bool) {
    self.Object.Set("collideWorldBounds", member)
}

// CheckCollision Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up. An object containing allowed collision.
func (self *PhysicsNinjaBody) CheckCollision() interface{}{
    return self.Object.Get("checkCollision")
}

// SetCheckCollisionA Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up. An object containing allowed collision.
func (self *PhysicsNinjaBody) SetCheckCollisionA(member interface{}) {
    self.Object.Set("checkCollision", member)
}

// Touching This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsNinjaBody) Touching() interface{}{
    return self.Object.Get("touching")
}

// SetTouchingA This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsNinjaBody) SetTouchingA(member interface{}) {
    self.Object.Set("touching", member)
}

// WasTouching This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsNinjaBody) WasTouching() interface{}{
    return self.Object.Get("wasTouching")
}

// SetWasTouchingA This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsNinjaBody) SetWasTouchingA(member interface{}) {
    self.Object.Set("wasTouching", member)
}

// MaxSpeed The maximum speed this body can travel at (taking drag and friction into account)
func (self *PhysicsNinjaBody) MaxSpeed() int{
    return self.Object.Get("maxSpeed").Int()
}

// SetMaxSpeedA The maximum speed this body can travel at (taking drag and friction into account)
func (self *PhysicsNinjaBody) SetMaxSpeedA(member int) {
    self.Object.Set("maxSpeed", member)
}

// X The x position.
func (self *PhysicsNinjaBody) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The x position.
func (self *PhysicsNinjaBody) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The y position.
func (self *PhysicsNinjaBody) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The y position.
func (self *PhysicsNinjaBody) SetYA(member int) {
    self.Object.Set("y", member)
}

// Width The width of this Body
func (self *PhysicsNinjaBody) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of this Body
func (self *PhysicsNinjaBody) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of this Body
func (self *PhysicsNinjaBody) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of this Body
func (self *PhysicsNinjaBody) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Bottom The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaBody) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaBody) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// Right The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaBody) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaBody) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Speed The speed of this Body
func (self *PhysicsNinjaBody) Speed() int{
    return self.Object.Get("speed").Int()
}

// SetSpeedA The speed of this Body
func (self *PhysicsNinjaBody) SetSpeedA(member int) {
    self.Object.Set("speed", member)
}

// Angle The angle of this Body
func (self *PhysicsNinjaBody) Angle() int{
    return self.Object.Get("angle").Int()
}

// SetAngleA The angle of this Body
func (self *PhysicsNinjaBody) SetAngleA(member int) {
    self.Object.Set("angle", member)
}


// PreUpdate Internal method.
func (self *PhysicsNinjaBody) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI Internal method.
func (self *PhysicsNinjaBody) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// PostUpdate Internal method.
func (self *PhysicsNinjaBody) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI Internal method.
func (self *PhysicsNinjaBody) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// SetZeroVelocity Stops all movement of this body.
func (self *PhysicsNinjaBody) SetZeroVelocity() {
    self.Object.Call("setZeroVelocity")
}

// SetZeroVelocityI Stops all movement of this body.
func (self *PhysicsNinjaBody) SetZeroVelocityI(args ...interface{}) {
    self.Object.Call("setZeroVelocity", args)
}

// Reset Resets all Body values and repositions on the Sprite.
func (self *PhysicsNinjaBody) Reset() {
    self.Object.Call("reset")
}

// ResetI Resets all Body values and repositions on the Sprite.
func (self *PhysicsNinjaBody) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// DeltaAbsX Returns the absolute delta x value.
func (self *PhysicsNinjaBody) DeltaAbsX() int{
    return self.Object.Call("deltaAbsX").Int()
}

// DeltaAbsXI Returns the absolute delta x value.
func (self *PhysicsNinjaBody) DeltaAbsXI(args ...interface{}) int{
    return self.Object.Call("deltaAbsX", args).Int()
}

// DeltaAbsY Returns the absolute delta y value.
func (self *PhysicsNinjaBody) DeltaAbsY() int{
    return self.Object.Call("deltaAbsY").Int()
}

// DeltaAbsYI Returns the absolute delta y value.
func (self *PhysicsNinjaBody) DeltaAbsYI(args ...interface{}) int{
    return self.Object.Call("deltaAbsY", args).Int()
}

// DeltaX Returns the delta x value. The difference between Body.x now and in the previous step.
func (self *PhysicsNinjaBody) DeltaX() int{
    return self.Object.Call("deltaX").Int()
}

// DeltaXI Returns the delta x value. The difference between Body.x now and in the previous step.
func (self *PhysicsNinjaBody) DeltaXI(args ...interface{}) int{
    return self.Object.Call("deltaX", args).Int()
}

// DeltaY Returns the delta y value. The difference between Body.y now and in the previous step.
func (self *PhysicsNinjaBody) DeltaY() int{
    return self.Object.Call("deltaY").Int()
}

// DeltaYI Returns the delta y value. The difference between Body.y now and in the previous step.
func (self *PhysicsNinjaBody) DeltaYI(args ...interface{}) int{
    return self.Object.Call("deltaY", args).Int()
}

// Destroy Destroys this body's reference to the sprite and system, and destroys its shape.
func (self *PhysicsNinjaBody) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this body's reference to the sprite and system, and destroys its shape.
func (self *PhysicsNinjaBody) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Render Render Sprite's Body.
func (self *PhysicsNinjaBody) Render(context interface{}, body *PhysicsNinjaBody) {
    self.Object.Call("render", context, body)
}

// Render1O Render Sprite's Body.
func (self *PhysicsNinjaBody) Render1O(context interface{}, body *PhysicsNinjaBody, color string) {
    self.Object.Call("render", context, body, color)
}

// Render2O Render Sprite's Body.
func (self *PhysicsNinjaBody) Render2O(context interface{}, body *PhysicsNinjaBody, color string, filled bool) {
    self.Object.Call("render", context, body, color, filled)
}

// RenderI Render Sprite's Body.
func (self *PhysicsNinjaBody) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

