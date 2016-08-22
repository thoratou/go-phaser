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


// Reference to the parent Sprite.
func (self *PhysicsNinjaBody) GetSprite() *Sprite{
    return &Sprite{self.Get("sprite")}
}

// Reference to the parent Sprite.
func (self *PhysicsNinjaBody) SetSprite(member *Sprite) {
    self.Set("sprite", member)
}

// Local reference to game.
func (self *PhysicsNinjaBody) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsNinjaBody) SetGame(member *Game) {
    self.Set("game", member)
}

// The type of physics system this body belongs to.
func (self *PhysicsNinjaBody) GetType() int{
    return self.Get("type").Int()
}

// The type of physics system this body belongs to.
func (self *PhysicsNinjaBody) SetType(member int) {
    self.Set("type", member)
}

// The parent physics system.
func (self *PhysicsNinjaBody) GetSystem() *PhysicsNinja{
    return &PhysicsNinja{self.Get("system")}
}

// The parent physics system.
func (self *PhysicsNinjaBody) SetSystem(member *PhysicsNinja) {
    self.Set("system", member)
}

// The AABB object this body is using for collision.
func (self *PhysicsNinjaBody) GetAabb() *PhysicsNinjaAABB{
    return &PhysicsNinjaAABB{self.Get("aabb")}
}

// The AABB object this body is using for collision.
func (self *PhysicsNinjaBody) SetAabb(member *PhysicsNinjaAABB) {
    self.Set("aabb", member)
}

// The Tile object this body is using for collision.
func (self *PhysicsNinjaBody) GetTile() *PhysicsNinjaTile{
    return &PhysicsNinjaTile{self.Get("tile")}
}

// The Tile object this body is using for collision.
func (self *PhysicsNinjaBody) SetTile(member *PhysicsNinjaTile) {
    self.Set("tile", member)
}

// The Circle object this body is using for collision.
func (self *PhysicsNinjaBody) GetCircle() *PhysicsNinjaCircle{
    return &PhysicsNinjaCircle{self.Get("circle")}
}

// The Circle object this body is using for collision.
func (self *PhysicsNinjaBody) SetCircle(member *PhysicsNinjaCircle) {
    self.Set("circle", member)
}

// A local reference to the body shape.
func (self *PhysicsNinjaBody) GetShape() interface{}{
    return self.Get("shape")
}

// A local reference to the body shape.
func (self *PhysicsNinjaBody) SetShape(member interface{}) {
    self.Set("shape", member)
}

// The drag applied to this object as it moves.
func (self *PhysicsNinjaBody) GetDrag() int{
    return self.Get("drag").Int()
}

// The drag applied to this object as it moves.
func (self *PhysicsNinjaBody) SetDrag(member int) {
    self.Set("drag", member)
}

// The friction applied to this object as it moves.
func (self *PhysicsNinjaBody) GetFriction() int{
    return self.Get("friction").Int()
}

// The friction applied to this object as it moves.
func (self *PhysicsNinjaBody) SetFriction(member int) {
    self.Set("friction", member)
}

// How much of the world gravity should be applied to this object? 1 = all of it, 0.5 = 50%, etc.
func (self *PhysicsNinjaBody) GetGravityScale() int{
    return self.Get("gravityScale").Int()
}

// How much of the world gravity should be applied to this object? 1 = all of it, 0.5 = 50%, etc.
func (self *PhysicsNinjaBody) SetGravityScale(member int) {
    self.Set("gravityScale", member)
}

// The bounciness of this object when it collides. A value between 0 and 1. We recommend setting it to 0.999 to avoid jittering.
func (self *PhysicsNinjaBody) GetBounce() int{
    return self.Get("bounce").Int()
}

// The bounciness of this object when it collides. A value between 0 and 1. We recommend setting it to 0.999 to avoid jittering.
func (self *PhysicsNinjaBody) SetBounce(member int) {
    self.Set("bounce", member)
}

// The velocity in pixels per second sq. of the Body.
func (self *PhysicsNinjaBody) GetVelocity() *Point{
    return &Point{self.Get("velocity")}
}

// The velocity in pixels per second sq. of the Body.
func (self *PhysicsNinjaBody) SetVelocity(member *Point) {
    self.Set("velocity", member)
}

// A const reference to the direction the Body is traveling or facing.
func (self *PhysicsNinjaBody) GetFacing() int{
    return self.Get("facing").Int()
}

// A const reference to the direction the Body is traveling or facing.
func (self *PhysicsNinjaBody) SetFacing(member int) {
    self.Set("facing", member)
}

// An immovable Body will not receive any impacts from other bodies. Not fully implemented.
func (self *PhysicsNinjaBody) GetImmovable() bool{
    return self.Get("immovable").Bool()
}

// An immovable Body will not receive any impacts from other bodies. Not fully implemented.
func (self *PhysicsNinjaBody) SetImmovable(member bool) {
    self.Set("immovable", member)
}

// A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsNinjaBody) GetCollideWorldBounds() bool{
    return self.Get("collideWorldBounds").Bool()
}

// A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsNinjaBody) SetCollideWorldBounds(member bool) {
    self.Set("collideWorldBounds", member)
}

// Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up. An object containing allowed collision.
func (self *PhysicsNinjaBody) GetCheckCollision() interface{}{
    return self.Get("checkCollision")
}

// Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up. An object containing allowed collision.
func (self *PhysicsNinjaBody) SetCheckCollision(member interface{}) {
    self.Set("checkCollision", member)
}

// This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsNinjaBody) GetTouching() interface{}{
    return self.Get("touching")
}

// This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsNinjaBody) SetTouching(member interface{}) {
    self.Set("touching", member)
}

// This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsNinjaBody) GetWasTouching() interface{}{
    return self.Get("wasTouching")
}

// This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsNinjaBody) SetWasTouching(member interface{}) {
    self.Set("wasTouching", member)
}

// The maximum speed this body can travel at (taking drag and friction into account)
func (self *PhysicsNinjaBody) GetMaxSpeed() int{
    return self.Get("maxSpeed").Int()
}

// The maximum speed this body can travel at (taking drag and friction into account)
func (self *PhysicsNinjaBody) SetMaxSpeed(member int) {
    self.Set("maxSpeed", member)
}

// The x position.
func (self *PhysicsNinjaBody) GetX() int{
    return self.Get("x").Int()
}

// The x position.
func (self *PhysicsNinjaBody) SetX(member int) {
    self.Set("x", member)
}

// The y position.
func (self *PhysicsNinjaBody) GetY() int{
    return self.Get("y").Int()
}

// The y position.
func (self *PhysicsNinjaBody) SetY(member int) {
    self.Set("y", member)
}

// The width of this Body
func (self *PhysicsNinjaBody) GetWidth() int{
    return self.Get("width").Int()
}

// The width of this Body
func (self *PhysicsNinjaBody) SetWidth(member int) {
    self.Set("width", member)
}

// The height of this Body
func (self *PhysicsNinjaBody) GetHeight() int{
    return self.Get("height").Int()
}

// The height of this Body
func (self *PhysicsNinjaBody) SetHeight(member int) {
    self.Set("height", member)
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaBody) GetBottom() int{
    return self.Get("bottom").Int()
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsNinjaBody) SetBottom(member int) {
    self.Set("bottom", member)
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaBody) GetRight() int{
    return self.Get("right").Int()
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsNinjaBody) SetRight(member int) {
    self.Set("right", member)
}

// The speed of this Body
func (self *PhysicsNinjaBody) GetSpeed() int{
    return self.Get("speed").Int()
}

// The speed of this Body
func (self *PhysicsNinjaBody) SetSpeed(member int) {
    self.Set("speed", member)
}

// The angle of this Body
func (self *PhysicsNinjaBody) GetAngle() int{
    return self.Get("angle").Int()
}

// The angle of this Body
func (self *PhysicsNinjaBody) SetAngle(member int) {
    self.Set("angle", member)
}



// Internal method.
func (self *PhysicsNinjaBody) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Internal method.
func (self *PhysicsNinjaBody) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}

// Stops all movement of this body.
func (self *PhysicsNinjaBody) SetZeroVelocityI(args ...interface{}) {
    self.Call("setZeroVelocity", args)
}

// Resets all Body values and repositions on the Sprite.
func (self *PhysicsNinjaBody) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Returns the absolute delta x value.
func (self *PhysicsNinjaBody) DeltaAbsXI(args ...interface{}) int{
    return self.Call("deltaAbsX", args).Int()
}

// Returns the absolute delta y value.
func (self *PhysicsNinjaBody) DeltaAbsYI(args ...interface{}) int{
    return self.Call("deltaAbsY", args).Int()
}

// Returns the delta x value. The difference between Body.x now and in the previous step.
func (self *PhysicsNinjaBody) DeltaXI(args ...interface{}) int{
    return self.Call("deltaX", args).Int()
}

// Returns the delta y value. The difference between Body.y now and in the previous step.
func (self *PhysicsNinjaBody) DeltaYI(args ...interface{}) int{
    return self.Call("deltaY", args).Int()
}

// Destroys this body's reference to the sprite and system, and destroys its shape.
func (self *PhysicsNinjaBody) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Render Sprite's Body.
func (self *PhysicsNinjaBody) RenderI(args ...interface{}) {
    self.Call("render", args)
}
