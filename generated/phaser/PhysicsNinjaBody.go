// Automatic generation for Phaser.Physics.Ninja.Body
// generated file PhysicsNinjaBody.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
type PhysicsNinjaBody struct {
    *js.Object
}


// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody(system *PhysicsNinja, sprite *Sprite) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Call("Phaser.Physics.Ninja.Body", system, sprite)}
}

// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody1O(system *PhysicsNinja, sprite *Sprite, type_ int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Call("Phaser.Physics.Ninja.Body", system, sprite, type_)}
}

// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody2O(system *PhysicsNinja, sprite *Sprite, type_ int, id int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Call("Phaser.Physics.Ninja.Body", system, sprite, type_, id)}
}

// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody3O(system *PhysicsNinja, sprite *Sprite, type_ int, id int, radius int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Call("Phaser.Physics.Ninja.Body", system, sprite, type_, id, radius)}
}

// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody4O(system *PhysicsNinja, sprite *Sprite, type_ int, id int, radius int, x int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Call("Phaser.Physics.Ninja.Body", system, sprite, type_, id, radius, x)}
}

// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody5O(system *PhysicsNinja, sprite *Sprite, type_ int, id int, radius int, x int, y int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Call("Phaser.Physics.Ninja.Body", system, sprite, type_, id, radius, x, y)}
}

// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody6O(system *PhysicsNinja, sprite *Sprite, type_ int, id int, radius int, x int, y int, width int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Call("Phaser.Physics.Ninja.Body", system, sprite, type_, id, radius, x, y, width)}
}

// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBody7O(system *PhysicsNinja, sprite *Sprite, type_ int, id int, radius int, x int, y int, width int, height int) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Call("Phaser.Physics.Ninja.Body", system, sprite, type_, id, radius, x, y, width, height)}
}

// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, bounce values etc all on the Body.
func NewPhysicsNinjaBodyI(args ...interface{}) *PhysicsNinjaBody {
    return &PhysicsNinjaBody{js.Global.Call("Phaser.Physics.Ninja.Body", args)}
}



// Reference to the parent Sprite.
func (self *PhysicsNinjaBody) GetSpriteA() *Sprite{
    return &Sprite{self.Object.Get("sprite")}
}

// Reference to the parent Sprite.
func (self *PhysicsNinjaBody) SetSpriteA(member *Sprite) {
    self.Object.Set("sprite", member)
}

// Local reference to game.
func (self *PhysicsNinjaBody) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsNinjaBody) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The type of physics system this body belongs to.
func (self *PhysicsNinjaBody) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The type of physics system this body belongs to.
func (self *PhysicsNinjaBody) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The parent physics system.
func (self *PhysicsNinjaBody) GetSystemA() *PhysicsNinja{
    return &PhysicsNinja{self.Object.Get("system")}
}

// The parent physics system.
func (self *PhysicsNinjaBody) SetSystemA(member *PhysicsNinja) {
    self.Object.Set("system", member)
}

// The AABB object this body is using for collision.
func (self *PhysicsNinjaBody) GetAabbA() *PhysicsNinjaAABB{
    return &PhysicsNinjaAABB{self.Object.Get("aabb")}
}

// The AABB object this body is using for collision.
func (self *PhysicsNinjaBody) SetAabbA(member *PhysicsNinjaAABB) {
    self.Object.Set("aabb", member)
}

// The Tile object this body is using for collision.
func (self *PhysicsNinjaBody) GetTileA() *PhysicsNinjaTile{
    return &PhysicsNinjaTile{self.Object.Get("tile")}
}

// The Tile object this body is using for collision.
func (self *PhysicsNinjaBody) SetTileA(member *PhysicsNinjaTile) {
    self.Object.Set("tile", member)
}

// The Circle object this body is using for collision.
func (self *PhysicsNinjaBody) GetCircleA() *PhysicsNinjaCircle{
    return &PhysicsNinjaCircle{self.Object.Get("circle")}
}

// The Circle object this body is using for collision.
func (self *PhysicsNinjaBody) SetCircleA(member *PhysicsNinjaCircle) {
    self.Object.Set("circle", member)
}

// A local reference to the body shape.
func (self *PhysicsNinjaBody) GetShapeA() interface{}{
    return self.Object.Get("shape")
}

// A local reference to the body shape.
func (self *PhysicsNinjaBody) SetShapeA(member interface{}) {
    self.Object.Set("shape", member)
}

// The drag applied to this object as it moves.
func (self *PhysicsNinjaBody) GetDragA() int{
    return self.Object.Get("drag").Int()
}

// The drag applied to this object as it moves.
func (self *PhysicsNinjaBody) SetDragA(member int) {
    self.Object.Set("drag", member)
}

// The friction applied to this object as it moves.
func (self *PhysicsNinjaBody) GetFrictionA() int{
    return self.Object.Get("friction").Int()
}

// The friction applied to this object as it moves.
func (self *PhysicsNinjaBody) SetFrictionA(member int) {
    self.Object.Set("friction", member)
}

// How much of the world gravity should be applied to this object? 1 = all of it, 0.5 = 50%, etc.
func (self *PhysicsNinjaBody) GetGravityScaleA() int{
    return self.Object.Get("gravityScale").Int()
}

// How much of the world gravity should be applied to this object? 1 = all of it, 0.5 = 50%, etc.
func (self *PhysicsNinjaBody) SetGravityScaleA(member int) {
    self.Object.Set("gravityScale", member)
}

// The bounciness of this object when it collides. A value between 0 and 1. We recommend setting it to 0.999 to avoid jittering.
func (self *PhysicsNinjaBody) GetBounceA() int{
    return self.Object.Get("bounce").Int()
}

// The bounciness of this object when it collides. A value between 0 and 1. We recommend setting it to 0.999 to avoid jittering.
func (self *PhysicsNinjaBody) SetBounceA(member int) {
    self.Object.Set("bounce", member)
}

// The velocity in pixels per second sq. of the Body.
func (self *PhysicsNinjaBody) GetVelocityA() *Point{
    return &Point{self.Object.Get("velocity")}
}

// The velocity in pixels per second sq. of the Body.
func (self *PhysicsNinjaBody) SetVelocityA(member *Point) {
    self.Object.Set("velocity", member)
}

// A const reference to the direction the Body is traveling or facing.
func (self *PhysicsNinjaBody) GetFacingA() int{
    return self.Object.Get("facing").Int()
}

// A const reference to the direction the Body is traveling or facing.
func (self *PhysicsNinjaBody) SetFacingA(member int) {
    self.Object.Set("facing", member)
}

// An immovable Body will not receive any impacts from other bodies. Not fully implemented.
func (self *PhysicsNinjaBody) GetImmovableA() bool{
    return self.Object.Get("immovable").Bool()
}

// An immovable Body will not receive any impacts from other bodies. Not fully implemented.
func (self *PhysicsNinjaBody) SetImmovableA(member bool) {
    self.Object.Set("immovable", member)
}

// A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsNinjaBody) GetCollideWorldBoundsA() bool{
    return self.Object.Get("collideWorldBounds").Bool()
}

// A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsNinjaBody) SetCollideWorldBoundsA(member bool) {
    self.Object.Set("collideWorldBounds", member)
}

// Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up. An object containing allowed collision.
func (self *PhysicsNinjaBody) GetCheckCollisionA() interface{}{
    return self.Object.Get("checkCollision")
}

// Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up. An object containing allowed collision.
func (self *PhysicsNinjaBody) SetCheckCollisionA(member interface{}) {
    self.Object.Set("checkCollision", member)
}

// This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsNinjaBody) GetTouchingA() interface{}{
    return self.Object.Get("touching")
}

// This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsNinjaBody) SetTouchingA(member interface{}) {
    self.Object.Set("touching", member)
}

// This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsNinjaBody) GetWasTouchingA() interface{}{
    return self.Object.Get("wasTouching")
}

// This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsNinjaBody) SetWasTouchingA(member interface{}) {
    self.Object.Set("wasTouching", member)
}

// The maximum speed this body can travel at (taking drag and friction into account)
func (self *PhysicsNinjaBody) GetMaxSpeedA() int{
    return self.Object.Get("maxSpeed").Int()
}

// The maximum speed this body can travel at (taking drag and friction into account)
func (self *PhysicsNinjaBody) SetMaxSpeedA(member int) {
    self.Object.Set("maxSpeed", member)
}

// The x position.
func (self *PhysicsNinjaBody) GetXA() int{
    return self.Object.Get("x").Int()
}

// The x position.
func (self *PhysicsNinjaBody) SetXA(member int) {
    self.Object.Set("x", member)
}

// The y position.
func (self *PhysicsNinjaBody) GetYA() int{
    return self.Object.Get("y").Int()
}

// The y position.
func (self *PhysicsNinjaBody) SetYA(member int) {
    self.Object.Set("y", member)
}

// The width of this Body
func (self *PhysicsNinjaBody) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of this Body
func (self *PhysicsNinjaBody) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of this Body
func (self *PhysicsNinjaBody) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of this Body
func (self *PhysicsNinjaBody) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaBody) GetBottomA() int{
    return self.Object.Get("bottom").Int()
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaBody) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaBody) GetRightA() int{
    return self.Object.Get("right").Int()
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaBody) SetRightA(member int) {
    self.Object.Set("right", member)
}

// The speed of this Body
func (self *PhysicsNinjaBody) GetSpeedA() int{
    return self.Object.Get("speed").Int()
}

// The speed of this Body
func (self *PhysicsNinjaBody) SetSpeedA(member int) {
    self.Object.Set("speed", member)
}

// The angle of this Body
func (self *PhysicsNinjaBody) GetAngleA() int{
    return self.Object.Get("angle").Int()
}

// The angle of this Body
func (self *PhysicsNinjaBody) SetAngleA(member int) {
    self.Object.Set("angle", member)
}



// Internal method.
func (self *PhysicsNinjaBody) PreUpdate() {
    self.Object.Call("preUpdate")
}

// Internal method.
func (self *PhysicsNinjaBody) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Internal method.
func (self *PhysicsNinjaBody) PostUpdate() {
    self.Object.Call("postUpdate")
}

// Internal method.
func (self *PhysicsNinjaBody) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Stops all movement of this body.
func (self *PhysicsNinjaBody) SetZeroVelocity() {
    self.Object.Call("setZeroVelocity")
}

// Stops all movement of this body.
func (self *PhysicsNinjaBody) SetZeroVelocityI(args ...interface{}) {
    self.Object.Call("setZeroVelocity", args)
}

// Resets all Body values and repositions on the Sprite.
func (self *PhysicsNinjaBody) Reset() {
    self.Object.Call("reset")
}

// Resets all Body values and repositions on the Sprite.
func (self *PhysicsNinjaBody) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Returns the absolute delta x value.
func (self *PhysicsNinjaBody) DeltaAbsX() int{
    return self.Object.Call("deltaAbsX").Int()
}

// Returns the absolute delta x value.
func (self *PhysicsNinjaBody) DeltaAbsXI(args ...interface{}) int{
    return self.Object.Call("deltaAbsX", args).Int()
}

// Returns the absolute delta y value.
func (self *PhysicsNinjaBody) DeltaAbsY() int{
    return self.Object.Call("deltaAbsY").Int()
}

// Returns the absolute delta y value.
func (self *PhysicsNinjaBody) DeltaAbsYI(args ...interface{}) int{
    return self.Object.Call("deltaAbsY", args).Int()
}

// Returns the delta x value. The difference between Body.x now and in the previous step.
func (self *PhysicsNinjaBody) DeltaX() int{
    return self.Object.Call("deltaX").Int()
}

// Returns the delta x value. The difference between Body.x now and in the previous step.
func (self *PhysicsNinjaBody) DeltaXI(args ...interface{}) int{
    return self.Object.Call("deltaX", args).Int()
}

// Returns the delta y value. The difference between Body.y now and in the previous step.
func (self *PhysicsNinjaBody) DeltaY() int{
    return self.Object.Call("deltaY").Int()
}

// Returns the delta y value. The difference between Body.y now and in the previous step.
func (self *PhysicsNinjaBody) DeltaYI(args ...interface{}) int{
    return self.Object.Call("deltaY", args).Int()
}

// Destroys this body's reference to the sprite and system, and destroys its shape.
func (self *PhysicsNinjaBody) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this body's reference to the sprite and system, and destroys its shape.
func (self *PhysicsNinjaBody) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Render Sprite's Body.
func (self *PhysicsNinjaBody) Render(context interface{}, body *PhysicsNinjaBody) {
    self.Object.Call("render", context, body)
}

// Render Sprite's Body.
func (self *PhysicsNinjaBody) Render1O(context interface{}, body *PhysicsNinjaBody, color string) {
    self.Object.Call("render", context, body, color)
}

// Render Sprite's Body.
func (self *PhysicsNinjaBody) Render2O(context interface{}, body *PhysicsNinjaBody, color string, filled bool) {
    self.Object.Call("render", context, body, color, filled)
}

// Render Sprite's Body.
func (self *PhysicsNinjaBody) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}
