// Package phaser Automatic generation for Phaser.Physics.Arcade.Body
// generated file PhysicsArcadeBody.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsArcadeBody The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, acceleration, bounce values etc all on the Body.
type PhysicsArcadeBody struct {
    *js.Object
}

// NewPhysicsArcadeBody The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, acceleration, bounce values etc all on the Body.
func NewPhysicsArcadeBody(sprite *Sprite) *PhysicsArcadeBody {
    return &PhysicsArcadeBody{js.Global.Get("Phaser").Get("Physics").Get("Arcade").Get("Body").New(sprite)}
}
// NewPhysicsArcadeBodyI The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, acceleration, bounce values etc all on the Body.
func NewPhysicsArcadeBodyI(args ...interface{}) *PhysicsArcadeBody {
    return &PhysicsArcadeBody{js.Global.Get("Phaser").Get("Physics").Get("Arcade").Get("Body").New(args)}
}



// PhysicsArcadeBody Binding conversion method to PhysicsArcadeBody point 
func ToPhysicsArcadeBody(jsStruct interface{}) *PhysicsArcadeBody {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsArcadeBody{Object: object}
	}
	return nil
}



// Sprite Reference to the parent Sprite.
func (self *PhysicsArcadeBody) Sprite() *Sprite{
    return &Sprite{self.Object.Get("sprite")}
}

// SetSpriteA Reference to the parent Sprite.
func (self *PhysicsArcadeBody) SetSpriteA(member *Sprite) {
    self.Object.Set("sprite", member)
}

// Game Local reference to game.
func (self *PhysicsArcadeBody) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *PhysicsArcadeBody) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Type The type of physics system this body belongs to.
func (self *PhysicsArcadeBody) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The type of physics system this body belongs to.
func (self *PhysicsArcadeBody) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Enable A disabled body won't be checked for any form of collision or overlap or have its pre/post updates run.
func (self *PhysicsArcadeBody) Enable() bool{
    return self.Object.Get("enable").Bool()
}

// SetEnableA A disabled body won't be checked for any form of collision or overlap or have its pre/post updates run.
func (self *PhysicsArcadeBody) SetEnableA(member bool) {
    self.Object.Set("enable", member)
}

// IsCircle If `true` this Body is using circular collision detection. If `false` it is using rectangular.
// Use `Body.setCircle` to control the collision shape this Body uses.
func (self *PhysicsArcadeBody) IsCircle() bool{
    return self.Object.Get("isCircle").Bool()
}

// SetIsCircleA If `true` this Body is using circular collision detection. If `false` it is using rectangular.
// Use `Body.setCircle` to control the collision shape this Body uses.
func (self *PhysicsArcadeBody) SetIsCircleA(member bool) {
    self.Object.Set("isCircle", member)
}

// Radius The radius of the circular collision shape this Body is using if Body.setCircle has been enabled.
// If you wish to change the radius then call `setCircle` again with the new value.
// If you wish to stop the Body using a circle then call `setCircle` with a radius of zero (or undefined).
func (self *PhysicsArcadeBody) Radius() int{
    return self.Object.Get("radius").Int()
}

// SetRadiusA The radius of the circular collision shape this Body is using if Body.setCircle has been enabled.
// If you wish to change the radius then call `setCircle` again with the new value.
// If you wish to stop the Body using a circle then call `setCircle` with a radius of zero (or undefined).
func (self *PhysicsArcadeBody) SetRadiusA(member int) {
    self.Object.Set("radius", member)
}

// Offset The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsArcadeBody) Offset() *Point{
    return &Point{self.Object.Get("offset")}
}

// SetOffsetA The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsArcadeBody) SetOffsetA(member *Point) {
    self.Object.Set("offset", member)
}

// Position The position of the physics body.
func (self *PhysicsArcadeBody) Position() *Point{
    return &Point{self.Object.Get("position")}
}

// SetPositionA The position of the physics body.
func (self *PhysicsArcadeBody) SetPositionA(member *Point) {
    self.Object.Set("position", member)
}

// Prev The previous position of the physics body.
func (self *PhysicsArcadeBody) Prev() *Point{
    return &Point{self.Object.Get("prev")}
}

// SetPrevA The previous position of the physics body.
func (self *PhysicsArcadeBody) SetPrevA(member *Point) {
    self.Object.Set("prev", member)
}

// AllowRotation Allow this Body to be rotated? (via angularVelocity, etc)
func (self *PhysicsArcadeBody) AllowRotation() bool{
    return self.Object.Get("allowRotation").Bool()
}

// SetAllowRotationA Allow this Body to be rotated? (via angularVelocity, etc)
func (self *PhysicsArcadeBody) SetAllowRotationA(member bool) {
    self.Object.Set("allowRotation", member)
}

// Rotation The Body's rotation in degrees, as calculated by its angularVelocity and angularAcceleration. Please understand that the collision Body
// itself never rotates, it is always axis-aligned. However these values are passed up to the parent Sprite and updates its rotation.
func (self *PhysicsArcadeBody) Rotation() int{
    return self.Object.Get("rotation").Int()
}

// SetRotationA The Body's rotation in degrees, as calculated by its angularVelocity and angularAcceleration. Please understand that the collision Body
// itself never rotates, it is always axis-aligned. However these values are passed up to the parent Sprite and updates its rotation.
func (self *PhysicsArcadeBody) SetRotationA(member int) {
    self.Object.Set("rotation", member)
}

// PreRotation The previous rotation of the physics body.
func (self *PhysicsArcadeBody) PreRotation() int{
    return self.Object.Get("preRotation").Int()
}

// SetPreRotationA The previous rotation of the physics body.
func (self *PhysicsArcadeBody) SetPreRotationA(member int) {
    self.Object.Set("preRotation", member)
}

// Width The calculated width of the physics body.
func (self *PhysicsArcadeBody) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The calculated width of the physics body.
func (self *PhysicsArcadeBody) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The calculated height of the physics body.
func (self *PhysicsArcadeBody) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The calculated height of the physics body.
func (self *PhysicsArcadeBody) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// SourceWidth The un-scaled original size.
func (self *PhysicsArcadeBody) SourceWidth() int{
    return self.Object.Get("sourceWidth").Int()
}

// SetSourceWidthA The un-scaled original size.
func (self *PhysicsArcadeBody) SetSourceWidthA(member int) {
    self.Object.Set("sourceWidth", member)
}

// SourceHeight The un-scaled original size.
func (self *PhysicsArcadeBody) SourceHeight() int{
    return self.Object.Get("sourceHeight").Int()
}

// SetSourceHeightA The un-scaled original size.
func (self *PhysicsArcadeBody) SetSourceHeightA(member int) {
    self.Object.Set("sourceHeight", member)
}

// HalfWidth The calculated width / 2 of the physics body.
func (self *PhysicsArcadeBody) HalfWidth() int{
    return self.Object.Get("halfWidth").Int()
}

// SetHalfWidthA The calculated width / 2 of the physics body.
func (self *PhysicsArcadeBody) SetHalfWidthA(member int) {
    self.Object.Set("halfWidth", member)
}

// HalfHeight The calculated height / 2 of the physics body.
func (self *PhysicsArcadeBody) HalfHeight() int{
    return self.Object.Get("halfHeight").Int()
}

// SetHalfHeightA The calculated height / 2 of the physics body.
func (self *PhysicsArcadeBody) SetHalfHeightA(member int) {
    self.Object.Set("halfHeight", member)
}

// Center The center coordinate of the Physics Body.
func (self *PhysicsArcadeBody) Center() *Point{
    return &Point{self.Object.Get("center")}
}

// SetCenterA The center coordinate of the Physics Body.
func (self *PhysicsArcadeBody) SetCenterA(member *Point) {
    self.Object.Set("center", member)
}

// Velocity The velocity, or rate of change in speed of the Body. Measured in pixels per second.
func (self *PhysicsArcadeBody) Velocity() *Point{
    return &Point{self.Object.Get("velocity")}
}

// SetVelocityA The velocity, or rate of change in speed of the Body. Measured in pixels per second.
func (self *PhysicsArcadeBody) SetVelocityA(member *Point) {
    self.Object.Set("velocity", member)
}

// NewVelocity The new velocity. Calculated during the Body.preUpdate and applied to its position.
func (self *PhysicsArcadeBody) NewVelocity() *Point{
    return &Point{self.Object.Get("newVelocity")}
}

// SetNewVelocityA The new velocity. Calculated during the Body.preUpdate and applied to its position.
func (self *PhysicsArcadeBody) SetNewVelocityA(member *Point) {
    self.Object.Set("newVelocity", member)
}

// DeltaMax The Sprite position is updated based on the delta x/y values. You can set a cap on those (both +-) using deltaMax.
func (self *PhysicsArcadeBody) DeltaMax() *Point{
    return &Point{self.Object.Get("deltaMax")}
}

// SetDeltaMaxA The Sprite position is updated based on the delta x/y values. You can set a cap on those (both +-) using deltaMax.
func (self *PhysicsArcadeBody) SetDeltaMaxA(member *Point) {
    self.Object.Set("deltaMax", member)
}

// Acceleration The acceleration is the rate of change of the velocity. Measured in pixels per second squared.
func (self *PhysicsArcadeBody) Acceleration() *Point{
    return &Point{self.Object.Get("acceleration")}
}

// SetAccelerationA The acceleration is the rate of change of the velocity. Measured in pixels per second squared.
func (self *PhysicsArcadeBody) SetAccelerationA(member *Point) {
    self.Object.Set("acceleration", member)
}

// Drag The drag applied to the motion of the Body.
func (self *PhysicsArcadeBody) Drag() *Point{
    return &Point{self.Object.Get("drag")}
}

// SetDragA The drag applied to the motion of the Body.
func (self *PhysicsArcadeBody) SetDragA(member *Point) {
    self.Object.Set("drag", member)
}

// AllowGravity Allow this Body to be influenced by gravity? Either world or local.
func (self *PhysicsArcadeBody) AllowGravity() bool{
    return self.Object.Get("allowGravity").Bool()
}

// SetAllowGravityA Allow this Body to be influenced by gravity? Either world or local.
func (self *PhysicsArcadeBody) SetAllowGravityA(member bool) {
    self.Object.Set("allowGravity", member)
}

// Gravity A local gravity applied to this Body. If non-zero this over rides any world gravity, unless Body.allowGravity is set to false.
func (self *PhysicsArcadeBody) Gravity() *Point{
    return &Point{self.Object.Get("gravity")}
}

// SetGravityA A local gravity applied to this Body. If non-zero this over rides any world gravity, unless Body.allowGravity is set to false.
func (self *PhysicsArcadeBody) SetGravityA(member *Point) {
    self.Object.Set("gravity", member)
}

// Bounce The elasticity of the Body when colliding. bounce.x/y = 1 means full rebound, bounce.x/y = 0.5 means 50% rebound velocity.
func (self *PhysicsArcadeBody) Bounce() *Point{
    return &Point{self.Object.Get("bounce")}
}

// SetBounceA The elasticity of the Body when colliding. bounce.x/y = 1 means full rebound, bounce.x/y = 0.5 means 50% rebound velocity.
func (self *PhysicsArcadeBody) SetBounceA(member *Point) {
    self.Object.Set("bounce", member)
}

// WorldBounce The elasticity of the Body when colliding with the World bounds.
// By default this property is `null`, in which case `Body.bounce` is used instead. Set this property
// to a Phaser.Point object in order to enable a World bounds specific bounce value.
func (self *PhysicsArcadeBody) WorldBounce() *Point{
    return &Point{self.Object.Get("worldBounce")}
}

// SetWorldBounceA The elasticity of the Body when colliding with the World bounds.
// By default this property is `null`, in which case `Body.bounce` is used instead. Set this property
// to a Phaser.Point object in order to enable a World bounds specific bounce value.
func (self *PhysicsArcadeBody) SetWorldBounceA(member *Point) {
    self.Object.Set("worldBounce", member)
}

// OnWorldBounds A Signal that is dispatched when this Body collides with the world bounds.
// Due to the potentially high volume of signals this could create it is disabled by default.
// To use this feature set this property to a Phaser.Signal: `sprite.body.onWorldBounds = new Phaser.Signal()`
// and it will be called when a collision happens, passing five arguments:
// `onWorldBounds(sprite, up, down, left, right)`
// where the Sprite is a reference to the Sprite that owns this Body, and the other arguments are booleans
// indicating on which side of the world the Body collided.
func (self *PhysicsArcadeBody) OnWorldBounds() *Signal{
    return &Signal{self.Object.Get("onWorldBounds")}
}

// SetOnWorldBoundsA A Signal that is dispatched when this Body collides with the world bounds.
// Due to the potentially high volume of signals this could create it is disabled by default.
// To use this feature set this property to a Phaser.Signal: `sprite.body.onWorldBounds = new Phaser.Signal()`
// and it will be called when a collision happens, passing five arguments:
// `onWorldBounds(sprite, up, down, left, right)`
// where the Sprite is a reference to the Sprite that owns this Body, and the other arguments are booleans
// indicating on which side of the world the Body collided.
func (self *PhysicsArcadeBody) SetOnWorldBoundsA(member *Signal) {
    self.Object.Set("onWorldBounds", member)
}

// OnCollide A Signal that is dispatched when this Body collides with another Body.
// 
// You still need to call `game.physics.arcade.collide` in your `update` method in order
// for this signal to be dispatched.
// 
// Usually you'd pass a callback to the `collide` method, but this signal provides for
// a different level of notification.
// 
// Due to the potentially high volume of signals this could create it is disabled by default.
// 
// To use this feature set this property to a Phaser.Signal: `sprite.body.onCollide = new Phaser.Signal()`
// and it will be called when a collision happens, passing two arguments: the sprites which collided.
// The first sprite in the argument is always the owner of this Body.
// 
// If two Bodies with this Signal set collide, both will dispatch the Signal.
func (self *PhysicsArcadeBody) OnCollide() *Signal{
    return &Signal{self.Object.Get("onCollide")}
}

// SetOnCollideA A Signal that is dispatched when this Body collides with another Body.
// 
// You still need to call `game.physics.arcade.collide` in your `update` method in order
// for this signal to be dispatched.
// 
// Usually you'd pass a callback to the `collide` method, but this signal provides for
// a different level of notification.
// 
// Due to the potentially high volume of signals this could create it is disabled by default.
// 
// To use this feature set this property to a Phaser.Signal: `sprite.body.onCollide = new Phaser.Signal()`
// and it will be called when a collision happens, passing two arguments: the sprites which collided.
// The first sprite in the argument is always the owner of this Body.
// 
// If two Bodies with this Signal set collide, both will dispatch the Signal.
func (self *PhysicsArcadeBody) SetOnCollideA(member *Signal) {
    self.Object.Set("onCollide", member)
}

// OnOverlap A Signal that is dispatched when this Body overlaps with another Body.
// 
// You still need to call `game.physics.arcade.overlap` in your `update` method in order
// for this signal to be dispatched.
// 
// Usually you'd pass a callback to the `overlap` method, but this signal provides for
// a different level of notification.
// 
// Due to the potentially high volume of signals this could create it is disabled by default.
// 
// To use this feature set this property to a Phaser.Signal: `sprite.body.onOverlap = new Phaser.Signal()`
// and it will be called when a collision happens, passing two arguments: the sprites which collided.
// The first sprite in the argument is always the owner of this Body.
// 
// If two Bodies with this Signal set collide, both will dispatch the Signal.
func (self *PhysicsArcadeBody) OnOverlap() *Signal{
    return &Signal{self.Object.Get("onOverlap")}
}

// SetOnOverlapA A Signal that is dispatched when this Body overlaps with another Body.
// 
// You still need to call `game.physics.arcade.overlap` in your `update` method in order
// for this signal to be dispatched.
// 
// Usually you'd pass a callback to the `overlap` method, but this signal provides for
// a different level of notification.
// 
// Due to the potentially high volume of signals this could create it is disabled by default.
// 
// To use this feature set this property to a Phaser.Signal: `sprite.body.onOverlap = new Phaser.Signal()`
// and it will be called when a collision happens, passing two arguments: the sprites which collided.
// The first sprite in the argument is always the owner of this Body.
// 
// If two Bodies with this Signal set collide, both will dispatch the Signal.
func (self *PhysicsArcadeBody) SetOnOverlapA(member *Signal) {
    self.Object.Set("onOverlap", member)
}

// MaxVelocity The maximum velocity in pixels per second sq. that the Body can reach.
func (self *PhysicsArcadeBody) MaxVelocity() *Point{
    return &Point{self.Object.Get("maxVelocity")}
}

// SetMaxVelocityA The maximum velocity in pixels per second sq. that the Body can reach.
func (self *PhysicsArcadeBody) SetMaxVelocityA(member *Point) {
    self.Object.Set("maxVelocity", member)
}

// Friction The amount of movement that will occur if another object 'rides' this one.
func (self *PhysicsArcadeBody) Friction() *Point{
    return &Point{self.Object.Get("friction")}
}

// SetFrictionA The amount of movement that will occur if another object 'rides' this one.
func (self *PhysicsArcadeBody) SetFrictionA(member *Point) {
    self.Object.Set("friction", member)
}

// AngularVelocity The angular velocity controls the rotation speed of the Body. It is measured in degrees per second.
func (self *PhysicsArcadeBody) AngularVelocity() int{
    return self.Object.Get("angularVelocity").Int()
}

// SetAngularVelocityA The angular velocity controls the rotation speed of the Body. It is measured in degrees per second.
func (self *PhysicsArcadeBody) SetAngularVelocityA(member int) {
    self.Object.Set("angularVelocity", member)
}

// AngularAcceleration The angular acceleration is the rate of change of the angular velocity. Measured in degrees per second squared.
func (self *PhysicsArcadeBody) AngularAcceleration() int{
    return self.Object.Get("angularAcceleration").Int()
}

// SetAngularAccelerationA The angular acceleration is the rate of change of the angular velocity. Measured in degrees per second squared.
func (self *PhysicsArcadeBody) SetAngularAccelerationA(member int) {
    self.Object.Set("angularAcceleration", member)
}

// AngularDrag The drag applied during the rotation of the Body. Measured in degrees per second squared.
func (self *PhysicsArcadeBody) AngularDrag() int{
    return self.Object.Get("angularDrag").Int()
}

// SetAngularDragA The drag applied during the rotation of the Body. Measured in degrees per second squared.
func (self *PhysicsArcadeBody) SetAngularDragA(member int) {
    self.Object.Set("angularDrag", member)
}

// MaxAngular The maximum angular velocity in degrees per second that the Body can reach.
func (self *PhysicsArcadeBody) MaxAngular() int{
    return self.Object.Get("maxAngular").Int()
}

// SetMaxAngularA The maximum angular velocity in degrees per second that the Body can reach.
func (self *PhysicsArcadeBody) SetMaxAngularA(member int) {
    self.Object.Set("maxAngular", member)
}

// Mass The mass of the Body. When two bodies collide their mass is used in the calculation to determine the exchange of velocity.
func (self *PhysicsArcadeBody) Mass() int{
    return self.Object.Get("mass").Int()
}

// SetMassA The mass of the Body. When two bodies collide their mass is used in the calculation to determine the exchange of velocity.
func (self *PhysicsArcadeBody) SetMassA(member int) {
    self.Object.Set("mass", member)
}

// Angle The angle of the Body's velocity in radians.
func (self *PhysicsArcadeBody) Angle() int{
    return self.Object.Get("angle").Int()
}

// SetAngleA The angle of the Body's velocity in radians.
func (self *PhysicsArcadeBody) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// Speed The speed of the Body as calculated by its velocity.
func (self *PhysicsArcadeBody) Speed() int{
    return self.Object.Get("speed").Int()
}

// SetSpeedA The speed of the Body as calculated by its velocity.
func (self *PhysicsArcadeBody) SetSpeedA(member int) {
    self.Object.Set("speed", member)
}

// Facing A const reference to the direction the Body is traveling or facing.
func (self *PhysicsArcadeBody) Facing() int{
    return self.Object.Get("facing").Int()
}

// SetFacingA A const reference to the direction the Body is traveling or facing.
func (self *PhysicsArcadeBody) SetFacingA(member int) {
    self.Object.Set("facing", member)
}

// Immovable An immovable Body will not receive any impacts from other bodies.
func (self *PhysicsArcadeBody) Immovable() bool{
    return self.Object.Get("immovable").Bool()
}

// SetImmovableA An immovable Body will not receive any impacts from other bodies.
func (self *PhysicsArcadeBody) SetImmovableA(member bool) {
    self.Object.Set("immovable", member)
}

// Moves If you have a Body that is being moved around the world via a tween or a Group motion, but its local x/y position never
// actually changes, then you should set Body.moves = false. Otherwise it will most likely fly off the screen.
// If you want the physics system to move the body around, then set moves to true. Set to true to allow the Physics system to move this Body, otherwise false to move it manually.
func (self *PhysicsArcadeBody) Moves() bool{
    return self.Object.Get("moves").Bool()
}

// SetMovesA If you have a Body that is being moved around the world via a tween or a Group motion, but its local x/y position never
// actually changes, then you should set Body.moves = false. Otherwise it will most likely fly off the screen.
// If you want the physics system to move the body around, then set moves to true. Set to true to allow the Physics system to move this Body, otherwise false to move it manually.
func (self *PhysicsArcadeBody) SetMovesA(member bool) {
    self.Object.Set("moves", member)
}

// CustomSeparateX This flag allows you to disable the custom x separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) CustomSeparateX() bool{
    return self.Object.Get("customSeparateX").Bool()
}

// SetCustomSeparateXA This flag allows you to disable the custom x separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) SetCustomSeparateXA(member bool) {
    self.Object.Set("customSeparateX", member)
}

// CustomSeparateY This flag allows you to disable the custom y separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) CustomSeparateY() bool{
    return self.Object.Get("customSeparateY").Bool()
}

// SetCustomSeparateYA This flag allows you to disable the custom y separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) SetCustomSeparateYA(member bool) {
    self.Object.Set("customSeparateY", member)
}

// OverlapX When this body collides with another, the amount of overlap is stored here. The amount of horizontal overlap during the collision.
func (self *PhysicsArcadeBody) OverlapX() int{
    return self.Object.Get("overlapX").Int()
}

// SetOverlapXA When this body collides with another, the amount of overlap is stored here. The amount of horizontal overlap during the collision.
func (self *PhysicsArcadeBody) SetOverlapXA(member int) {
    self.Object.Set("overlapX", member)
}

// OverlapY When this body collides with another, the amount of overlap is stored here. The amount of vertical overlap during the collision.
func (self *PhysicsArcadeBody) OverlapY() int{
    return self.Object.Get("overlapY").Int()
}

// SetOverlapYA When this body collides with another, the amount of overlap is stored here. The amount of vertical overlap during the collision.
func (self *PhysicsArcadeBody) SetOverlapYA(member int) {
    self.Object.Set("overlapY", member)
}

// OverlapR If `Body.isCircle` is true, and this body collides with another circular body, the amount of overlap is stored here. The amount of overlap during the collision.
func (self *PhysicsArcadeBody) OverlapR() int{
    return self.Object.Get("overlapR").Int()
}

// SetOverlapRA If `Body.isCircle` is true, and this body collides with another circular body, the amount of overlap is stored here. The amount of overlap during the collision.
func (self *PhysicsArcadeBody) SetOverlapRA(member int) {
    self.Object.Set("overlapR", member)
}

// Embedded If a body is overlapping with another body, but neither of them are moving (maybe they spawned on-top of each other?) this is set to true. Body embed value.
func (self *PhysicsArcadeBody) Embedded() bool{
    return self.Object.Get("embedded").Bool()
}

// SetEmbeddedA If a body is overlapping with another body, but neither of them are moving (maybe they spawned on-top of each other?) this is set to true. Body embed value.
func (self *PhysicsArcadeBody) SetEmbeddedA(member bool) {
    self.Object.Set("embedded", member)
}

// CollideWorldBounds A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsArcadeBody) CollideWorldBounds() bool{
    return self.Object.Get("collideWorldBounds").Bool()
}

// SetCollideWorldBoundsA A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsArcadeBody) SetCollideWorldBoundsA(member bool) {
    self.Object.Set("collideWorldBounds", member)
}

// CheckCollision Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up.
// If you need to disable a Body entirely, use `body.enable = false`, this will also disable motion.
// If you need to disable just collision and/or overlap checks, but retain motion, set `checkCollision.none = true`. An object containing allowed collision.
func (self *PhysicsArcadeBody) CheckCollision() interface{}{
    return self.Object.Get("checkCollision")
}

// SetCheckCollisionA Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up.
// If you need to disable a Body entirely, use `body.enable = false`, this will also disable motion.
// If you need to disable just collision and/or overlap checks, but retain motion, set `checkCollision.none = true`. An object containing allowed collision.
func (self *PhysicsArcadeBody) SetCheckCollisionA(member interface{}) {
    self.Object.Set("checkCollision", member)
}

// Touching This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsArcadeBody) Touching() interface{}{
    return self.Object.Get("touching")
}

// SetTouchingA This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsArcadeBody) SetTouchingA(member interface{}) {
    self.Object.Set("touching", member)
}

// WasTouching This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsArcadeBody) WasTouching() interface{}{
    return self.Object.Get("wasTouching")
}

// SetWasTouchingA This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsArcadeBody) SetWasTouchingA(member interface{}) {
    self.Object.Set("wasTouching", member)
}

// Blocked This object is populated with boolean values when the Body collides with the World bounds or a Tile.
// For example if blocked.up is true then the Body cannot move up. An object containing on which faces this Body is blocked from moving, if any.
func (self *PhysicsArcadeBody) Blocked() interface{}{
    return self.Object.Get("blocked")
}

// SetBlockedA This object is populated with boolean values when the Body collides with the World bounds or a Tile.
// For example if blocked.up is true then the Body cannot move up. An object containing on which faces this Body is blocked from moving, if any.
func (self *PhysicsArcadeBody) SetBlockedA(member interface{}) {
    self.Object.Set("blocked", member)
}

// TilePadding If this is an especially small or fast moving object then it can sometimes skip over tilemap collisions if it moves through a tile in a step.
// Set this padding value to add extra padding to its bounds. tilePadding.x applied to its width, y to its height. Extra padding to be added to this sprite's dimensions when checking for tile collision.
func (self *PhysicsArcadeBody) TilePadding() *Point{
    return &Point{self.Object.Get("tilePadding")}
}

// SetTilePaddingA If this is an especially small or fast moving object then it can sometimes skip over tilemap collisions if it moves through a tile in a step.
// Set this padding value to add extra padding to its bounds. tilePadding.x applied to its width, y to its height. Extra padding to be added to this sprite's dimensions when checking for tile collision.
func (self *PhysicsArcadeBody) SetTilePaddingA(member *Point) {
    self.Object.Set("tilePadding", member)
}

// Dirty If this Body in a preUpdate (true) or postUpdate (false) state?
func (self *PhysicsArcadeBody) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// SetDirtyA If this Body in a preUpdate (true) or postUpdate (false) state?
func (self *PhysicsArcadeBody) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// SkipQuadTree If true and you collide this Sprite against a Group, it will disable the collision check from using a QuadTree.
func (self *PhysicsArcadeBody) SkipQuadTree() bool{
    return self.Object.Get("skipQuadTree").Bool()
}

// SetSkipQuadTreeA If true and you collide this Sprite against a Group, it will disable the collision check from using a QuadTree.
func (self *PhysicsArcadeBody) SetSkipQuadTreeA(member bool) {
    self.Object.Set("skipQuadTree", member)
}

// SyncBounds If true the Body will check itself against the Sprite.getBounds() dimensions and adjust its width and height accordingly.
// If false it will compare its dimensions against the Sprite scale instead, and adjust its width height if the scale has changed.
// Typically you would need to enable syncBounds if your sprite is the child of a responsive display object such as a FlexLayer, 
// or in any situation where the Sprite scale doesn't change, but its parents scale is effecting the dimensions regardless.
func (self *PhysicsArcadeBody) SyncBounds() bool{
    return self.Object.Get("syncBounds").Bool()
}

// SetSyncBoundsA If true the Body will check itself against the Sprite.getBounds() dimensions and adjust its width and height accordingly.
// If false it will compare its dimensions against the Sprite scale instead, and adjust its width height if the scale has changed.
// Typically you would need to enable syncBounds if your sprite is the child of a responsive display object such as a FlexLayer, 
// or in any situation where the Sprite scale doesn't change, but its parents scale is effecting the dimensions regardless.
func (self *PhysicsArcadeBody) SetSyncBoundsA(member bool) {
    self.Object.Set("syncBounds", member)
}

// IsMoving Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) IsMoving() bool{
    return self.Object.Get("isMoving").Bool()
}

// SetIsMovingA Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) SetIsMovingA(member bool) {
    self.Object.Set("isMoving", member)
}

// StopVelocityOnCollide Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) StopVelocityOnCollide() bool{
    return self.Object.Get("stopVelocityOnCollide").Bool()
}

// SetStopVelocityOnCollideA Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) SetStopVelocityOnCollideA(member bool) {
    self.Object.Set("stopVelocityOnCollide", member)
}

// OnMoveComplete Listen for the completion of `moveTo` or `moveFrom` events.
func (self *PhysicsArcadeBody) OnMoveComplete() *Signal{
    return &Signal{self.Object.Get("onMoveComplete")}
}

// SetOnMoveCompleteA Listen for the completion of `moveTo` or `moveFrom` events.
func (self *PhysicsArcadeBody) SetOnMoveCompleteA(member *Signal) {
    self.Object.Set("onMoveComplete", member)
}

// MovementCallback Optional callback. If set, invoked during the running of `moveTo` or `moveFrom` events.
func (self *PhysicsArcadeBody) MovementCallback() interface{}{
    return self.Object.Get("movementCallback")
}

// SetMovementCallbackA Optional callback. If set, invoked during the running of `moveTo` or `moveFrom` events.
func (self *PhysicsArcadeBody) SetMovementCallbackA(member interface{}) {
    self.Object.Set("movementCallback", member)
}

// MovementCallbackContext Context in which to call the movementCallback.
func (self *PhysicsArcadeBody) MovementCallbackContext() interface{}{
    return self.Object.Get("movementCallbackContext")
}

// SetMovementCallbackContextA Context in which to call the movementCallback.
func (self *PhysicsArcadeBody) SetMovementCallbackContextA(member interface{}) {
    self.Object.Set("movementCallbackContext", member)
}

// Left The x position of the Body. The same as `Body.x`.
func (self *PhysicsArcadeBody) Left() int{
    return self.Object.Get("left").Int()
}

// SetLeftA The x position of the Body. The same as `Body.x`.
func (self *PhysicsArcadeBody) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// Right The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsArcadeBody) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsArcadeBody) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Top The y position of the Body. The same as `Body.y`.
func (self *PhysicsArcadeBody) Top() int{
    return self.Object.Get("top").Int()
}

// SetTopA The y position of the Body. The same as `Body.y`.
func (self *PhysicsArcadeBody) SetTopA(member int) {
    self.Object.Set("top", member)
}

// Bottom The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsArcadeBody) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsArcadeBody) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// X The x position.
func (self *PhysicsArcadeBody) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The x position.
func (self *PhysicsArcadeBody) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The y position.
func (self *PhysicsArcadeBody) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The y position.
func (self *PhysicsArcadeBody) SetYA(member int) {
    self.Object.Set("y", member)
}


// UpdateBounds Internal method.
func (self *PhysicsArcadeBody) UpdateBounds() {
    self.Object.Call("updateBounds")
}

// UpdateBoundsI Internal method.
func (self *PhysicsArcadeBody) UpdateBoundsI(args ...interface{}) {
    self.Object.Call("updateBounds", args)
}

// PreUpdate Internal method.
func (self *PhysicsArcadeBody) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI Internal method.
func (self *PhysicsArcadeBody) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// UpdateMovement Internal method.
func (self *PhysicsArcadeBody) UpdateMovement() {
    self.Object.Call("updateMovement")
}

// UpdateMovementI Internal method.
func (self *PhysicsArcadeBody) UpdateMovementI(args ...interface{}) {
    self.Object.Call("updateMovement", args)
}

// StopMovement If this Body is moving as a result of a call to `moveTo` or `moveFrom` (i.e. it
// has Body.isMoving true), then calling this method will stop the movement before
// either the duration or distance counters expire.
// 
// The `onMoveComplete` signal is dispatched.
func (self *PhysicsArcadeBody) StopMovement() {
    self.Object.Call("stopMovement")
}

// StopMovement1O If this Body is moving as a result of a call to `moveTo` or `moveFrom` (i.e. it
// has Body.isMoving true), then calling this method will stop the movement before
// either the duration or distance counters expire.
// 
// The `onMoveComplete` signal is dispatched.
func (self *PhysicsArcadeBody) StopMovement1O(stopVelocity bool) {
    self.Object.Call("stopMovement", stopVelocity)
}

// StopMovementI If this Body is moving as a result of a call to `moveTo` or `moveFrom` (i.e. it
// has Body.isMoving true), then calling this method will stop the movement before
// either the duration or distance counters expire.
// 
// The `onMoveComplete` signal is dispatched.
func (self *PhysicsArcadeBody) StopMovementI(args ...interface{}) {
    self.Object.Call("stopMovement", args)
}

// PostUpdate Internal method.
func (self *PhysicsArcadeBody) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI Internal method.
func (self *PhysicsArcadeBody) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// CheckWorldBounds Internal method.
func (self *PhysicsArcadeBody) CheckWorldBounds() bool{
    return self.Object.Call("checkWorldBounds").Bool()
}

// CheckWorldBoundsI Internal method.
func (self *PhysicsArcadeBody) CheckWorldBoundsI(args ...interface{}) bool{
    return self.Object.Call("checkWorldBounds", args).Bool()
}

// MoveFrom Note: This method is experimental, and may be changed or removed in a future release.
// 
// This method moves the Body in the given direction, for the duration specified.
// It works by setting the velocity on the Body, and an internal timer, and then
// monitoring the duration each frame. When the duration is up the movement is
// stopped and the `Body.onMoveComplete` signal is dispatched.
// 
// Movement also stops if the Body collides or overlaps with any other Body.
// 
// You can control if the velocity should be reset to zero on collision, by using
// the property `Body.stopVelocityOnCollide`.
// 
// Stop the movement at any time by calling `Body.stopMovement`.
// 
// You can optionally set a speed in pixels per second. If not specified it
// will use the current `Body.speed` value. If this is zero, the function will return false.
// 
// Please note that due to browser timings you should allow for a variance in 
// when the duration will actually expire. Depending on system it may be as much as
// +- 50ms. Also this method doesn't take into consideration any other forces acting
// on the Body, such as Gravity, drag or maxVelocity, all of which may impact the
// movement.
func (self *PhysicsArcadeBody) MoveFrom(duration int) bool{
    return self.Object.Call("moveFrom", duration).Bool()
}

// MoveFrom1O Note: This method is experimental, and may be changed or removed in a future release.
// 
// This method moves the Body in the given direction, for the duration specified.
// It works by setting the velocity on the Body, and an internal timer, and then
// monitoring the duration each frame. When the duration is up the movement is
// stopped and the `Body.onMoveComplete` signal is dispatched.
// 
// Movement also stops if the Body collides or overlaps with any other Body.
// 
// You can control if the velocity should be reset to zero on collision, by using
// the property `Body.stopVelocityOnCollide`.
// 
// Stop the movement at any time by calling `Body.stopMovement`.
// 
// You can optionally set a speed in pixels per second. If not specified it
// will use the current `Body.speed` value. If this is zero, the function will return false.
// 
// Please note that due to browser timings you should allow for a variance in 
// when the duration will actually expire. Depending on system it may be as much as
// +- 50ms. Also this method doesn't take into consideration any other forces acting
// on the Body, such as Gravity, drag or maxVelocity, all of which may impact the
// movement.
func (self *PhysicsArcadeBody) MoveFrom1O(duration int, speed int) bool{
    return self.Object.Call("moveFrom", duration, speed).Bool()
}

// MoveFrom2O Note: This method is experimental, and may be changed or removed in a future release.
// 
// This method moves the Body in the given direction, for the duration specified.
// It works by setting the velocity on the Body, and an internal timer, and then
// monitoring the duration each frame. When the duration is up the movement is
// stopped and the `Body.onMoveComplete` signal is dispatched.
// 
// Movement also stops if the Body collides or overlaps with any other Body.
// 
// You can control if the velocity should be reset to zero on collision, by using
// the property `Body.stopVelocityOnCollide`.
// 
// Stop the movement at any time by calling `Body.stopMovement`.
// 
// You can optionally set a speed in pixels per second. If not specified it
// will use the current `Body.speed` value. If this is zero, the function will return false.
// 
// Please note that due to browser timings you should allow for a variance in 
// when the duration will actually expire. Depending on system it may be as much as
// +- 50ms. Also this method doesn't take into consideration any other forces acting
// on the Body, such as Gravity, drag or maxVelocity, all of which may impact the
// movement.
func (self *PhysicsArcadeBody) MoveFrom2O(duration int, speed int, direction int) bool{
    return self.Object.Call("moveFrom", duration, speed, direction).Bool()
}

// MoveFromI Note: This method is experimental, and may be changed or removed in a future release.
// 
// This method moves the Body in the given direction, for the duration specified.
// It works by setting the velocity on the Body, and an internal timer, and then
// monitoring the duration each frame. When the duration is up the movement is
// stopped and the `Body.onMoveComplete` signal is dispatched.
// 
// Movement also stops if the Body collides or overlaps with any other Body.
// 
// You can control if the velocity should be reset to zero on collision, by using
// the property `Body.stopVelocityOnCollide`.
// 
// Stop the movement at any time by calling `Body.stopMovement`.
// 
// You can optionally set a speed in pixels per second. If not specified it
// will use the current `Body.speed` value. If this is zero, the function will return false.
// 
// Please note that due to browser timings you should allow for a variance in 
// when the duration will actually expire. Depending on system it may be as much as
// +- 50ms. Also this method doesn't take into consideration any other forces acting
// on the Body, such as Gravity, drag or maxVelocity, all of which may impact the
// movement.
func (self *PhysicsArcadeBody) MoveFromI(args ...interface{}) bool{
    return self.Object.Call("moveFrom", args).Bool()
}

// MoveTo Note: This method is experimental, and may be changed or removed in a future release.
// 
// This method moves the Body in the given direction, for the duration specified.
// It works by setting the velocity on the Body, and an internal distance counter.
// The distance is monitored each frame. When the distance equals the distance
// specified in this call, the movement is stopped, and the `Body.onMoveComplete` 
// signal is dispatched.
// 
// Movement also stops if the Body collides or overlaps with any other Body.
// 
// You can control if the velocity should be reset to zero on collision, by using
// the property `Body.stopVelocityOnCollide`.
// 
// Stop the movement at any time by calling `Body.stopMovement`.
// 
// Please note that due to browser timings you should allow for a variance in 
// when the distance will actually expire.
// 
// Note: This method doesn't take into consideration any other forces acting
// on the Body, such as Gravity, drag or maxVelocity, all of which may impact the
// movement.
func (self *PhysicsArcadeBody) MoveTo(duration int, distance int) bool{
    return self.Object.Call("moveTo", duration, distance).Bool()
}

// MoveTo1O Note: This method is experimental, and may be changed or removed in a future release.
// 
// This method moves the Body in the given direction, for the duration specified.
// It works by setting the velocity on the Body, and an internal distance counter.
// The distance is monitored each frame. When the distance equals the distance
// specified in this call, the movement is stopped, and the `Body.onMoveComplete` 
// signal is dispatched.
// 
// Movement also stops if the Body collides or overlaps with any other Body.
// 
// You can control if the velocity should be reset to zero on collision, by using
// the property `Body.stopVelocityOnCollide`.
// 
// Stop the movement at any time by calling `Body.stopMovement`.
// 
// Please note that due to browser timings you should allow for a variance in 
// when the distance will actually expire.
// 
// Note: This method doesn't take into consideration any other forces acting
// on the Body, such as Gravity, drag or maxVelocity, all of which may impact the
// movement.
func (self *PhysicsArcadeBody) MoveTo1O(duration int, distance int, direction int) bool{
    return self.Object.Call("moveTo", duration, distance, direction).Bool()
}

// MoveToI Note: This method is experimental, and may be changed or removed in a future release.
// 
// This method moves the Body in the given direction, for the duration specified.
// It works by setting the velocity on the Body, and an internal distance counter.
// The distance is monitored each frame. When the distance equals the distance
// specified in this call, the movement is stopped, and the `Body.onMoveComplete` 
// signal is dispatched.
// 
// Movement also stops if the Body collides or overlaps with any other Body.
// 
// You can control if the velocity should be reset to zero on collision, by using
// the property `Body.stopVelocityOnCollide`.
// 
// Stop the movement at any time by calling `Body.stopMovement`.
// 
// Please note that due to browser timings you should allow for a variance in 
// when the distance will actually expire.
// 
// Note: This method doesn't take into consideration any other forces acting
// on the Body, such as Gravity, drag or maxVelocity, all of which may impact the
// movement.
func (self *PhysicsArcadeBody) MoveToI(args ...interface{}) bool{
    return self.Object.Call("moveTo", args).Bool()
}

// SetSize You can modify the size of the physics Body to be any dimension you need.
// This allows you to make it smaller, or larger, than the parent Sprite.
// You can also control the x and y offset of the Body. This is the position of the
// Body relative to the top-left of the Sprite _texture_.
// 
// For example: If you have a Sprite with a texture that is 80x100 in size,
// and you want the physics body to be 32x32 pixels in the middle of the texture, you would do:
// 
// `setSize(32, 32, 24, 34)`
// 
// Where the first two parameters is the new Body size (32x32 pixels).
// 24 is the horizontal offset of the Body from the top-left of the Sprites texture, and 34
// is the vertical offset.
// 
// Calling `setSize` on a Body that has already had `setCircle` will reset all of the Circle
// properties, making this Body rectangular again.
func (self *PhysicsArcadeBody) SetSize(width int, height int) {
    self.Object.Call("setSize", width, height)
}

// SetSize1O You can modify the size of the physics Body to be any dimension you need.
// This allows you to make it smaller, or larger, than the parent Sprite.
// You can also control the x and y offset of the Body. This is the position of the
// Body relative to the top-left of the Sprite _texture_.
// 
// For example: If you have a Sprite with a texture that is 80x100 in size,
// and you want the physics body to be 32x32 pixels in the middle of the texture, you would do:
// 
// `setSize(32, 32, 24, 34)`
// 
// Where the first two parameters is the new Body size (32x32 pixels).
// 24 is the horizontal offset of the Body from the top-left of the Sprites texture, and 34
// is the vertical offset.
// 
// Calling `setSize` on a Body that has already had `setCircle` will reset all of the Circle
// properties, making this Body rectangular again.
func (self *PhysicsArcadeBody) SetSize1O(width int, height int, offsetX int) {
    self.Object.Call("setSize", width, height, offsetX)
}

// SetSize2O You can modify the size of the physics Body to be any dimension you need.
// This allows you to make it smaller, or larger, than the parent Sprite.
// You can also control the x and y offset of the Body. This is the position of the
// Body relative to the top-left of the Sprite _texture_.
// 
// For example: If you have a Sprite with a texture that is 80x100 in size,
// and you want the physics body to be 32x32 pixels in the middle of the texture, you would do:
// 
// `setSize(32, 32, 24, 34)`
// 
// Where the first two parameters is the new Body size (32x32 pixels).
// 24 is the horizontal offset of the Body from the top-left of the Sprites texture, and 34
// is the vertical offset.
// 
// Calling `setSize` on a Body that has already had `setCircle` will reset all of the Circle
// properties, making this Body rectangular again.
func (self *PhysicsArcadeBody) SetSize2O(width int, height int, offsetX int, offsetY int) {
    self.Object.Call("setSize", width, height, offsetX, offsetY)
}

// SetSizeI You can modify the size of the physics Body to be any dimension you need.
// This allows you to make it smaller, or larger, than the parent Sprite.
// You can also control the x and y offset of the Body. This is the position of the
// Body relative to the top-left of the Sprite _texture_.
// 
// For example: If you have a Sprite with a texture that is 80x100 in size,
// and you want the physics body to be 32x32 pixels in the middle of the texture, you would do:
// 
// `setSize(32, 32, 24, 34)`
// 
// Where the first two parameters is the new Body size (32x32 pixels).
// 24 is the horizontal offset of the Body from the top-left of the Sprites texture, and 34
// is the vertical offset.
// 
// Calling `setSize` on a Body that has already had `setCircle` will reset all of the Circle
// properties, making this Body rectangular again.
func (self *PhysicsArcadeBody) SetSizeI(args ...interface{}) {
    self.Object.Call("setSize", args)
}

// SetCircle Sets this Body as using a circle, of the given radius, for all collision detection instead of a rectangle.
// The radius is given in pixels and is the distance from the center of the circle to the edge.
// 
// You can also control the x and y offset, which is the position of the Body relative to the top-left of the Sprite.
// 
// To change a Body back to being rectangular again call `Body.setSize`.
// 
// Note: Circular collision only happens with other Arcade Physics bodies, it does not
// work against tile maps, where rectangular collision is the only method supported.
func (self *PhysicsArcadeBody) SetCircle() {
    self.Object.Call("setCircle")
}

// SetCircle1O Sets this Body as using a circle, of the given radius, for all collision detection instead of a rectangle.
// The radius is given in pixels and is the distance from the center of the circle to the edge.
// 
// You can also control the x and y offset, which is the position of the Body relative to the top-left of the Sprite.
// 
// To change a Body back to being rectangular again call `Body.setSize`.
// 
// Note: Circular collision only happens with other Arcade Physics bodies, it does not
// work against tile maps, where rectangular collision is the only method supported.
func (self *PhysicsArcadeBody) SetCircle1O(radius int) {
    self.Object.Call("setCircle", radius)
}

// SetCircle2O Sets this Body as using a circle, of the given radius, for all collision detection instead of a rectangle.
// The radius is given in pixels and is the distance from the center of the circle to the edge.
// 
// You can also control the x and y offset, which is the position of the Body relative to the top-left of the Sprite.
// 
// To change a Body back to being rectangular again call `Body.setSize`.
// 
// Note: Circular collision only happens with other Arcade Physics bodies, it does not
// work against tile maps, where rectangular collision is the only method supported.
func (self *PhysicsArcadeBody) SetCircle2O(radius int, offsetX int) {
    self.Object.Call("setCircle", radius, offsetX)
}

// SetCircle3O Sets this Body as using a circle, of the given radius, for all collision detection instead of a rectangle.
// The radius is given in pixels and is the distance from the center of the circle to the edge.
// 
// You can also control the x and y offset, which is the position of the Body relative to the top-left of the Sprite.
// 
// To change a Body back to being rectangular again call `Body.setSize`.
// 
// Note: Circular collision only happens with other Arcade Physics bodies, it does not
// work against tile maps, where rectangular collision is the only method supported.
func (self *PhysicsArcadeBody) SetCircle3O(radius int, offsetX int, offsetY int) {
    self.Object.Call("setCircle", radius, offsetX, offsetY)
}

// SetCircleI Sets this Body as using a circle, of the given radius, for all collision detection instead of a rectangle.
// The radius is given in pixels and is the distance from the center of the circle to the edge.
// 
// You can also control the x and y offset, which is the position of the Body relative to the top-left of the Sprite.
// 
// To change a Body back to being rectangular again call `Body.setSize`.
// 
// Note: Circular collision only happens with other Arcade Physics bodies, it does not
// work against tile maps, where rectangular collision is the only method supported.
func (self *PhysicsArcadeBody) SetCircleI(args ...interface{}) {
    self.Object.Call("setCircle", args)
}

// Reset Resets all Body values (velocity, acceleration, rotation, etc)
func (self *PhysicsArcadeBody) Reset(x int, y int) {
    self.Object.Call("reset", x, y)
}

// ResetI Resets all Body values (velocity, acceleration, rotation, etc)
func (self *PhysicsArcadeBody) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// GetBounds Returns the bounds of this physics body.
// 
// Only used internally by the World collision methods.
func (self *PhysicsArcadeBody) GetBounds(obj interface{}) interface{}{
    return self.Object.Call("getBounds", obj)
}

// GetBoundsI Returns the bounds of this physics body.
// 
// Only used internally by the World collision methods.
func (self *PhysicsArcadeBody) GetBoundsI(args ...interface{}) interface{}{
    return self.Object.Call("getBounds", args)
}

// HitTest Tests if a world point lies within this Body.
func (self *PhysicsArcadeBody) HitTest(x int, y int) bool{
    return self.Object.Call("hitTest", x, y).Bool()
}

// HitTestI Tests if a world point lies within this Body.
func (self *PhysicsArcadeBody) HitTestI(args ...interface{}) bool{
    return self.Object.Call("hitTest", args).Bool()
}

// OnFloor Returns true if the bottom of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnFloor() bool{
    return self.Object.Call("onFloor").Bool()
}

// OnFloorI Returns true if the bottom of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnFloorI(args ...interface{}) bool{
    return self.Object.Call("onFloor", args).Bool()
}

// OnCeiling Returns true if the top of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnCeiling() bool{
    return self.Object.Call("onCeiling").Bool()
}

// OnCeilingI Returns true if the top of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnCeilingI(args ...interface{}) bool{
    return self.Object.Call("onCeiling", args).Bool()
}

// OnWall Returns true if either side of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnWall() bool{
    return self.Object.Call("onWall").Bool()
}

// OnWallI Returns true if either side of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnWallI(args ...interface{}) bool{
    return self.Object.Call("onWall", args).Bool()
}

// DeltaAbsX Returns the absolute delta x value.
func (self *PhysicsArcadeBody) DeltaAbsX() int{
    return self.Object.Call("deltaAbsX").Int()
}

// DeltaAbsXI Returns the absolute delta x value.
func (self *PhysicsArcadeBody) DeltaAbsXI(args ...interface{}) int{
    return self.Object.Call("deltaAbsX", args).Int()
}

// DeltaAbsY Returns the absolute delta y value.
func (self *PhysicsArcadeBody) DeltaAbsY() int{
    return self.Object.Call("deltaAbsY").Int()
}

// DeltaAbsYI Returns the absolute delta y value.
func (self *PhysicsArcadeBody) DeltaAbsYI(args ...interface{}) int{
    return self.Object.Call("deltaAbsY", args).Int()
}

// DeltaX Returns the delta x value. The difference between Body.x now and in the previous step.
func (self *PhysicsArcadeBody) DeltaX() int{
    return self.Object.Call("deltaX").Int()
}

// DeltaXI Returns the delta x value. The difference between Body.x now and in the previous step.
func (self *PhysicsArcadeBody) DeltaXI(args ...interface{}) int{
    return self.Object.Call("deltaX", args).Int()
}

// DeltaY Returns the delta y value. The difference between Body.y now and in the previous step.
func (self *PhysicsArcadeBody) DeltaY() int{
    return self.Object.Call("deltaY").Int()
}

// DeltaYI Returns the delta y value. The difference between Body.y now and in the previous step.
func (self *PhysicsArcadeBody) DeltaYI(args ...interface{}) int{
    return self.Object.Call("deltaY", args).Int()
}

// DeltaZ Returns the delta z value. The difference between Body.rotation now and in the previous step.
func (self *PhysicsArcadeBody) DeltaZ() int{
    return self.Object.Call("deltaZ").Int()
}

// DeltaZI Returns the delta z value. The difference between Body.rotation now and in the previous step.
func (self *PhysicsArcadeBody) DeltaZI(args ...interface{}) int{
    return self.Object.Call("deltaZ", args).Int()
}

// Destroy Destroys this Body.
// 
// First it calls Group.removeFromHash if the Game Object this Body belongs to is part of a Group.
// Then it nulls the Game Objects body reference, and nulls this Body.sprite reference.
func (self *PhysicsArcadeBody) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this Body.
// 
// First it calls Group.removeFromHash if the Game Object this Body belongs to is part of a Group.
// Then it nulls the Game Objects body reference, and nulls this Body.sprite reference.
func (self *PhysicsArcadeBody) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Render Render Sprite Body.
func (self *PhysicsArcadeBody) Render(context interface{}, body *PhysicsArcadeBody) {
    self.Object.Call("render", context, body)
}

// Render1O Render Sprite Body.
func (self *PhysicsArcadeBody) Render1O(context interface{}, body *PhysicsArcadeBody, color string) {
    self.Object.Call("render", context, body, color)
}

// Render2O Render Sprite Body.
func (self *PhysicsArcadeBody) Render2O(context interface{}, body *PhysicsArcadeBody, color string, filled bool) {
    self.Object.Call("render", context, body, color, filled)
}

// RenderI Render Sprite Body.
func (self *PhysicsArcadeBody) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// RenderBodyInfo Render Sprite Body Physics Data as text.
func (self *PhysicsArcadeBody) RenderBodyInfo(body *PhysicsArcadeBody, x int, y int) {
    self.Object.Call("renderBodyInfo", body, x, y)
}

// RenderBodyInfo1O Render Sprite Body Physics Data as text.
func (self *PhysicsArcadeBody) RenderBodyInfo1O(body *PhysicsArcadeBody, x int, y int, color string) {
    self.Object.Call("renderBodyInfo", body, x, y, color)
}

// RenderBodyInfoI Render Sprite Body Physics Data as text.
func (self *PhysicsArcadeBody) RenderBodyInfoI(args ...interface{}) {
    self.Object.Call("renderBodyInfo", args)
}

