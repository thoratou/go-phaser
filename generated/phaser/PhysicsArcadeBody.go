// Automatic generation for Phaser.Physics.Arcade.Body
// generated file PhysicsArcadeBody.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Physics Body is linked to a single Sprite. All physics operations should be performed against the body rather than
// the Sprite itself. For example you can set the velocity, acceleration, bounce values etc all on the Body.
type PhysicsArcadeBody struct {
    *js.Object
}


// Reference to the parent Sprite.
func (self *PhysicsArcadeBody) GetSpriteA() *Sprite{
    return &Sprite{self.Object.Get("sprite")}
}

// Reference to the parent Sprite.
func (self *PhysicsArcadeBody) SetSpriteA(member *Sprite) {
    self.Object.Set("sprite", member)
}

// Local reference to game.
func (self *PhysicsArcadeBody) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsArcadeBody) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The type of physics system this body belongs to.
func (self *PhysicsArcadeBody) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The type of physics system this body belongs to.
func (self *PhysicsArcadeBody) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// A disabled body won't be checked for any form of collision or overlap or have its pre/post updates run.
func (self *PhysicsArcadeBody) GetEnableA() bool{
    return self.Object.Get("enable").Bool()
}

// A disabled body won't be checked for any form of collision or overlap or have its pre/post updates run.
func (self *PhysicsArcadeBody) SetEnableA(member bool) {
    self.Object.Set("enable", member)
}

// If `true` this Body is using circular collision detection. If `false` it is using rectangular.
// Use `Body.setCircle` to control the collision shape this Body uses.
func (self *PhysicsArcadeBody) GetIsCircleA() bool{
    return self.Object.Get("isCircle").Bool()
}

// If `true` this Body is using circular collision detection. If `false` it is using rectangular.
// Use `Body.setCircle` to control the collision shape this Body uses.
func (self *PhysicsArcadeBody) SetIsCircleA(member bool) {
    self.Object.Set("isCircle", member)
}

// The radius of the circular collision shape this Body is using if Body.setCircle has been enabled.
// If you wish to change the radius then call `setCircle` again with the new value.
// If you wish to stop the Body using a circle then call `setCircle` with a radius of zero (or undefined).
func (self *PhysicsArcadeBody) GetRadiusA() int{
    return self.Object.Get("radius").Int()
}

// The radius of the circular collision shape this Body is using if Body.setCircle has been enabled.
// If you wish to change the radius then call `setCircle` again with the new value.
// If you wish to stop the Body using a circle then call `setCircle` with a radius of zero (or undefined).
func (self *PhysicsArcadeBody) SetRadiusA(member int) {
    self.Object.Set("radius", member)
}

// The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsArcadeBody) GetOffsetA() *Point{
    return &Point{self.Object.Get("offset")}
}

// The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsArcadeBody) SetOffsetA(member *Point) {
    self.Object.Set("offset", member)
}

// The position of the physics body.
func (self *PhysicsArcadeBody) GetPositionA() *Point{
    return &Point{self.Object.Get("position")}
}

// The position of the physics body.
func (self *PhysicsArcadeBody) SetPositionA(member *Point) {
    self.Object.Set("position", member)
}

// The previous position of the physics body.
func (self *PhysicsArcadeBody) GetPrevA() *Point{
    return &Point{self.Object.Get("prev")}
}

// The previous position of the physics body.
func (self *PhysicsArcadeBody) SetPrevA(member *Point) {
    self.Object.Set("prev", member)
}

// Allow this Body to be rotated? (via angularVelocity, etc)
func (self *PhysicsArcadeBody) GetAllowRotationA() bool{
    return self.Object.Get("allowRotation").Bool()
}

// Allow this Body to be rotated? (via angularVelocity, etc)
func (self *PhysicsArcadeBody) SetAllowRotationA(member bool) {
    self.Object.Set("allowRotation", member)
}

// An Arcade Physics Body can have angularVelocity and angularAcceleration. Please understand that the collision Body
// itself never rotates, it is always axis-aligned. However these values are passed up to the parent Sprite and updates its rotation.
func (self *PhysicsArcadeBody) GetRotationA() int{
    return self.Object.Get("rotation").Int()
}

// An Arcade Physics Body can have angularVelocity and angularAcceleration. Please understand that the collision Body
// itself never rotates, it is always axis-aligned. However these values are passed up to the parent Sprite and updates its rotation.
func (self *PhysicsArcadeBody) SetRotationA(member int) {
    self.Object.Set("rotation", member)
}

// The previous rotation of the physics body.
func (self *PhysicsArcadeBody) GetPreRotationA() int{
    return self.Object.Get("preRotation").Int()
}

// The previous rotation of the physics body.
func (self *PhysicsArcadeBody) SetPreRotationA(member int) {
    self.Object.Set("preRotation", member)
}

// The calculated width of the physics body.
func (self *PhysicsArcadeBody) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The calculated width of the physics body.
func (self *PhysicsArcadeBody) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The calculated height of the physics body.
func (self *PhysicsArcadeBody) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The calculated height of the physics body.
func (self *PhysicsArcadeBody) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The un-scaled original size.
func (self *PhysicsArcadeBody) GetSourceWidthA() int{
    return self.Object.Get("sourceWidth").Int()
}

// The un-scaled original size.
func (self *PhysicsArcadeBody) SetSourceWidthA(member int) {
    self.Object.Set("sourceWidth", member)
}

// The un-scaled original size.
func (self *PhysicsArcadeBody) GetSourceHeightA() int{
    return self.Object.Get("sourceHeight").Int()
}

// The un-scaled original size.
func (self *PhysicsArcadeBody) SetSourceHeightA(member int) {
    self.Object.Set("sourceHeight", member)
}

// The calculated width / 2 of the physics body.
func (self *PhysicsArcadeBody) GetHalfWidthA() int{
    return self.Object.Get("halfWidth").Int()
}

// The calculated width / 2 of the physics body.
func (self *PhysicsArcadeBody) SetHalfWidthA(member int) {
    self.Object.Set("halfWidth", member)
}

// The calculated height / 2 of the physics body.
func (self *PhysicsArcadeBody) GetHalfHeightA() int{
    return self.Object.Get("halfHeight").Int()
}

// The calculated height / 2 of the physics body.
func (self *PhysicsArcadeBody) SetHalfHeightA(member int) {
    self.Object.Set("halfHeight", member)
}

// The center coordinate of the Physics Body.
func (self *PhysicsArcadeBody) GetCenterA() *Point{
    return &Point{self.Object.Get("center")}
}

// The center coordinate of the Physics Body.
func (self *PhysicsArcadeBody) SetCenterA(member *Point) {
    self.Object.Set("center", member)
}

// The velocity, or rate of change in speed of the Body. Measured in pixels per second.
func (self *PhysicsArcadeBody) GetVelocityA() *Point{
    return &Point{self.Object.Get("velocity")}
}

// The velocity, or rate of change in speed of the Body. Measured in pixels per second.
func (self *PhysicsArcadeBody) SetVelocityA(member *Point) {
    self.Object.Set("velocity", member)
}

// The new velocity. Calculated during the Body.preUpdate and applied to its position.
func (self *PhysicsArcadeBody) GetNewVelocityA() *Point{
    return &Point{self.Object.Get("newVelocity")}
}

// The new velocity. Calculated during the Body.preUpdate and applied to its position.
func (self *PhysicsArcadeBody) SetNewVelocityA(member *Point) {
    self.Object.Set("newVelocity", member)
}

// The Sprite position is updated based on the delta x/y values. You can set a cap on those (both +-) using deltaMax.
func (self *PhysicsArcadeBody) GetDeltaMaxA() *Point{
    return &Point{self.Object.Get("deltaMax")}
}

// The Sprite position is updated based on the delta x/y values. You can set a cap on those (both +-) using deltaMax.
func (self *PhysicsArcadeBody) SetDeltaMaxA(member *Point) {
    self.Object.Set("deltaMax", member)
}

// The acceleration is the rate of change of the velocity. Measured in pixels per second squared.
func (self *PhysicsArcadeBody) GetAccelerationA() *Point{
    return &Point{self.Object.Get("acceleration")}
}

// The acceleration is the rate of change of the velocity. Measured in pixels per second squared.
func (self *PhysicsArcadeBody) SetAccelerationA(member *Point) {
    self.Object.Set("acceleration", member)
}

// The drag applied to the motion of the Body.
func (self *PhysicsArcadeBody) GetDragA() *Point{
    return &Point{self.Object.Get("drag")}
}

// The drag applied to the motion of the Body.
func (self *PhysicsArcadeBody) SetDragA(member *Point) {
    self.Object.Set("drag", member)
}

// Allow this Body to be influenced by gravity? Either world or local.
func (self *PhysicsArcadeBody) GetAllowGravityA() bool{
    return self.Object.Get("allowGravity").Bool()
}

// Allow this Body to be influenced by gravity? Either world or local.
func (self *PhysicsArcadeBody) SetAllowGravityA(member bool) {
    self.Object.Set("allowGravity", member)
}

// A local gravity applied to this Body. If non-zero this over rides any world gravity, unless Body.allowGravity is set to false.
func (self *PhysicsArcadeBody) GetGravityA() *Point{
    return &Point{self.Object.Get("gravity")}
}

// A local gravity applied to this Body. If non-zero this over rides any world gravity, unless Body.allowGravity is set to false.
func (self *PhysicsArcadeBody) SetGravityA(member *Point) {
    self.Object.Set("gravity", member)
}

// The elasticity of the Body when colliding. bounce.x/y = 1 means full rebound, bounce.x/y = 0.5 means 50% rebound velocity.
func (self *PhysicsArcadeBody) GetBounceA() *Point{
    return &Point{self.Object.Get("bounce")}
}

// The elasticity of the Body when colliding. bounce.x/y = 1 means full rebound, bounce.x/y = 0.5 means 50% rebound velocity.
func (self *PhysicsArcadeBody) SetBounceA(member *Point) {
    self.Object.Set("bounce", member)
}

// The elasticity of the Body when colliding with the World bounds.
// By default this property is `null`, in which case `Body.bounce` is used instead. Set this property
// to a Phaser.Point object in order to enable a World bounds specific bounce value.
func (self *PhysicsArcadeBody) GetWorldBounceA() *Point{
    return &Point{self.Object.Get("worldBounce")}
}

// The elasticity of the Body when colliding with the World bounds.
// By default this property is `null`, in which case `Body.bounce` is used instead. Set this property
// to a Phaser.Point object in order to enable a World bounds specific bounce value.
func (self *PhysicsArcadeBody) SetWorldBounceA(member *Point) {
    self.Object.Set("worldBounce", member)
}

// A Signal that is dispatched when this Body collides with the world bounds.
// Due to the potentially high volume of signals this could create it is disabled by default.
// To use this feature set this property to a Phaser.Signal: `sprite.body.onWorldBounds = new Phaser.Signal()`
// and it will be called when a collision happens, passing five arguments:
// `onWorldBounds(sprite, up, down, left, right)`
// where the Sprite is a reference to the Sprite that owns this Body, and the other arguments are booleans
// indicating on which side of the world the Body collided.
func (self *PhysicsArcadeBody) GetOnWorldBoundsA() *Signal{
    return &Signal{self.Object.Get("onWorldBounds")}
}

// A Signal that is dispatched when this Body collides with the world bounds.
// Due to the potentially high volume of signals this could create it is disabled by default.
// To use this feature set this property to a Phaser.Signal: `sprite.body.onWorldBounds = new Phaser.Signal()`
// and it will be called when a collision happens, passing five arguments:
// `onWorldBounds(sprite, up, down, left, right)`
// where the Sprite is a reference to the Sprite that owns this Body, and the other arguments are booleans
// indicating on which side of the world the Body collided.
func (self *PhysicsArcadeBody) SetOnWorldBoundsA(member *Signal) {
    self.Object.Set("onWorldBounds", member)
}

// A Signal that is dispatched when this Body collides with another Body.
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
func (self *PhysicsArcadeBody) GetOnCollideA() *Signal{
    return &Signal{self.Object.Get("onCollide")}
}

// A Signal that is dispatched when this Body collides with another Body.
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

// A Signal that is dispatched when this Body overlaps with another Body.
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
func (self *PhysicsArcadeBody) GetOnOverlapA() *Signal{
    return &Signal{self.Object.Get("onOverlap")}
}

// A Signal that is dispatched when this Body overlaps with another Body.
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

// The maximum velocity in pixels per second sq. that the Body can reach.
func (self *PhysicsArcadeBody) GetMaxVelocityA() *Point{
    return &Point{self.Object.Get("maxVelocity")}
}

// The maximum velocity in pixels per second sq. that the Body can reach.
func (self *PhysicsArcadeBody) SetMaxVelocityA(member *Point) {
    self.Object.Set("maxVelocity", member)
}

// The amount of movement that will occur if another object 'rides' this one.
func (self *PhysicsArcadeBody) GetFrictionA() *Point{
    return &Point{self.Object.Get("friction")}
}

// The amount of movement that will occur if another object 'rides' this one.
func (self *PhysicsArcadeBody) SetFrictionA(member *Point) {
    self.Object.Set("friction", member)
}

// The angular velocity controls the rotation speed of the Body. It is measured in radians per second.
func (self *PhysicsArcadeBody) GetAngularVelocityA() int{
    return self.Object.Get("angularVelocity").Int()
}

// The angular velocity controls the rotation speed of the Body. It is measured in radians per second.
func (self *PhysicsArcadeBody) SetAngularVelocityA(member int) {
    self.Object.Set("angularVelocity", member)
}

// The angular acceleration is the rate of change of the angular velocity. Measured in radians per second squared.
func (self *PhysicsArcadeBody) GetAngularAccelerationA() int{
    return self.Object.Get("angularAcceleration").Int()
}

// The angular acceleration is the rate of change of the angular velocity. Measured in radians per second squared.
func (self *PhysicsArcadeBody) SetAngularAccelerationA(member int) {
    self.Object.Set("angularAcceleration", member)
}

// The drag applied during the rotation of the Body.
func (self *PhysicsArcadeBody) GetAngularDragA() int{
    return self.Object.Get("angularDrag").Int()
}

// The drag applied during the rotation of the Body.
func (self *PhysicsArcadeBody) SetAngularDragA(member int) {
    self.Object.Set("angularDrag", member)
}

// The maximum angular velocity in radians per second that the Body can reach.
func (self *PhysicsArcadeBody) GetMaxAngularA() int{
    return self.Object.Get("maxAngular").Int()
}

// The maximum angular velocity in radians per second that the Body can reach.
func (self *PhysicsArcadeBody) SetMaxAngularA(member int) {
    self.Object.Set("maxAngular", member)
}

// The mass of the Body. When two bodies collide their mass is used in the calculation to determine the exchange of velocity.
func (self *PhysicsArcadeBody) GetMassA() int{
    return self.Object.Get("mass").Int()
}

// The mass of the Body. When two bodies collide their mass is used in the calculation to determine the exchange of velocity.
func (self *PhysicsArcadeBody) SetMassA(member int) {
    self.Object.Set("mass", member)
}

// The angle of the Body in radians, as calculated by its angularVelocity.
func (self *PhysicsArcadeBody) GetAngleA() int{
    return self.Object.Get("angle").Int()
}

// The angle of the Body in radians, as calculated by its angularVelocity.
func (self *PhysicsArcadeBody) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// The speed of the Body as calculated by its velocity.
func (self *PhysicsArcadeBody) GetSpeedA() int{
    return self.Object.Get("speed").Int()
}

// The speed of the Body as calculated by its velocity.
func (self *PhysicsArcadeBody) SetSpeedA(member int) {
    self.Object.Set("speed", member)
}

// A const reference to the direction the Body is traveling or facing.
func (self *PhysicsArcadeBody) GetFacingA() int{
    return self.Object.Get("facing").Int()
}

// A const reference to the direction the Body is traveling or facing.
func (self *PhysicsArcadeBody) SetFacingA(member int) {
    self.Object.Set("facing", member)
}

// An immovable Body will not receive any impacts from other bodies.
func (self *PhysicsArcadeBody) GetImmovableA() bool{
    return self.Object.Get("immovable").Bool()
}

// An immovable Body will not receive any impacts from other bodies.
func (self *PhysicsArcadeBody) SetImmovableA(member bool) {
    self.Object.Set("immovable", member)
}

// If you have a Body that is being moved around the world via a tween or a Group motion, but its local x/y position never
// actually changes, then you should set Body.moves = false. Otherwise it will most likely fly off the screen.
// If you want the physics system to move the body around, then set moves to true. Set to true to allow the Physics system to move this Body, otherwise false to move it manually.
func (self *PhysicsArcadeBody) GetMovesA() bool{
    return self.Object.Get("moves").Bool()
}

// If you have a Body that is being moved around the world via a tween or a Group motion, but its local x/y position never
// actually changes, then you should set Body.moves = false. Otherwise it will most likely fly off the screen.
// If you want the physics system to move the body around, then set moves to true. Set to true to allow the Physics system to move this Body, otherwise false to move it manually.
func (self *PhysicsArcadeBody) SetMovesA(member bool) {
    self.Object.Set("moves", member)
}

// This flag allows you to disable the custom x separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) GetCustomSeparateXA() bool{
    return self.Object.Get("customSeparateX").Bool()
}

// This flag allows you to disable the custom x separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) SetCustomSeparateXA(member bool) {
    self.Object.Set("customSeparateX", member)
}

// This flag allows you to disable the custom y separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) GetCustomSeparateYA() bool{
    return self.Object.Get("customSeparateY").Bool()
}

// This flag allows you to disable the custom y separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) SetCustomSeparateYA(member bool) {
    self.Object.Set("customSeparateY", member)
}

// When this body collides with another, the amount of overlap is stored here. The amount of horizontal overlap during the collision.
func (self *PhysicsArcadeBody) GetOverlapXA() int{
    return self.Object.Get("overlapX").Int()
}

// When this body collides with another, the amount of overlap is stored here. The amount of horizontal overlap during the collision.
func (self *PhysicsArcadeBody) SetOverlapXA(member int) {
    self.Object.Set("overlapX", member)
}

// When this body collides with another, the amount of overlap is stored here. The amount of vertical overlap during the collision.
func (self *PhysicsArcadeBody) GetOverlapYA() int{
    return self.Object.Get("overlapY").Int()
}

// When this body collides with another, the amount of overlap is stored here. The amount of vertical overlap during the collision.
func (self *PhysicsArcadeBody) SetOverlapYA(member int) {
    self.Object.Set("overlapY", member)
}

// If `Body.isCircle` is true, and this body collides with another circular body, the amount of overlap is stored here. The amount of overlap during the collision.
func (self *PhysicsArcadeBody) GetOverlapRA() int{
    return self.Object.Get("overlapR").Int()
}

// If `Body.isCircle` is true, and this body collides with another circular body, the amount of overlap is stored here. The amount of overlap during the collision.
func (self *PhysicsArcadeBody) SetOverlapRA(member int) {
    self.Object.Set("overlapR", member)
}

// If a body is overlapping with another body, but neither of them are moving (maybe they spawned on-top of each other?) this is set to true. Body embed value.
func (self *PhysicsArcadeBody) GetEmbeddedA() bool{
    return self.Object.Get("embedded").Bool()
}

// If a body is overlapping with another body, but neither of them are moving (maybe they spawned on-top of each other?) this is set to true. Body embed value.
func (self *PhysicsArcadeBody) SetEmbeddedA(member bool) {
    self.Object.Set("embedded", member)
}

// A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsArcadeBody) GetCollideWorldBoundsA() bool{
    return self.Object.Get("collideWorldBounds").Bool()
}

// A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsArcadeBody) SetCollideWorldBoundsA(member bool) {
    self.Object.Set("collideWorldBounds", member)
}

// Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up. An object containing allowed collision.
func (self *PhysicsArcadeBody) GetCheckCollisionA() interface{}{
    return self.Object.Get("checkCollision")
}

// Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up. An object containing allowed collision.
func (self *PhysicsArcadeBody) SetCheckCollisionA(member interface{}) {
    self.Object.Set("checkCollision", member)
}

// This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsArcadeBody) GetTouchingA() interface{}{
    return self.Object.Get("touching")
}

// This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsArcadeBody) SetTouchingA(member interface{}) {
    self.Object.Set("touching", member)
}

// This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsArcadeBody) GetWasTouchingA() interface{}{
    return self.Object.Get("wasTouching")
}

// This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsArcadeBody) SetWasTouchingA(member interface{}) {
    self.Object.Set("wasTouching", member)
}

// This object is populated with boolean values when the Body collides with the World bounds or a Tile.
// For example if blocked.up is true then the Body cannot move up. An object containing on which faces this Body is blocked from moving, if any.
func (self *PhysicsArcadeBody) GetBlockedA() interface{}{
    return self.Object.Get("blocked")
}

// This object is populated with boolean values when the Body collides with the World bounds or a Tile.
// For example if blocked.up is true then the Body cannot move up. An object containing on which faces this Body is blocked from moving, if any.
func (self *PhysicsArcadeBody) SetBlockedA(member interface{}) {
    self.Object.Set("blocked", member)
}

// If this is an especially small or fast moving object then it can sometimes skip over tilemap collisions if it moves through a tile in a step.
// Set this padding value to add extra padding to its bounds. tilePadding.x applied to its width, y to its height. Extra padding to be added to this sprite's dimensions when checking for tile collision.
func (self *PhysicsArcadeBody) GetTilePaddingA() *Point{
    return &Point{self.Object.Get("tilePadding")}
}

// If this is an especially small or fast moving object then it can sometimes skip over tilemap collisions if it moves through a tile in a step.
// Set this padding value to add extra padding to its bounds. tilePadding.x applied to its width, y to its height. Extra padding to be added to this sprite's dimensions when checking for tile collision.
func (self *PhysicsArcadeBody) SetTilePaddingA(member *Point) {
    self.Object.Set("tilePadding", member)
}

// If this Body in a preUpdate (true) or postUpdate (false) state?
func (self *PhysicsArcadeBody) GetDirtyA() bool{
    return self.Object.Get("dirty").Bool()
}

// If this Body in a preUpdate (true) or postUpdate (false) state?
func (self *PhysicsArcadeBody) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// If true and you collide this Sprite against a Group, it will disable the collision check from using a QuadTree.
func (self *PhysicsArcadeBody) GetSkipQuadTreeA() bool{
    return self.Object.Get("skipQuadTree").Bool()
}

// If true and you collide this Sprite against a Group, it will disable the collision check from using a QuadTree.
func (self *PhysicsArcadeBody) SetSkipQuadTreeA(member bool) {
    self.Object.Set("skipQuadTree", member)
}

// If true the Body will check itself against the Sprite.getBounds() dimensions and adjust its width and height accordingly.
// If false it will compare its dimensions against the Sprite scale instead, and adjust its width height if the scale has changed.
// Typically you would need to enable syncBounds if your sprite is the child of a responsive display object such as a FlexLayer, 
// or in any situation where the Sprite scale doesn't change, but its parents scale is effecting the dimensions regardless.
func (self *PhysicsArcadeBody) GetSyncBoundsA() bool{
    return self.Object.Get("syncBounds").Bool()
}

// If true the Body will check itself against the Sprite.getBounds() dimensions and adjust its width and height accordingly.
// If false it will compare its dimensions against the Sprite scale instead, and adjust its width height if the scale has changed.
// Typically you would need to enable syncBounds if your sprite is the child of a responsive display object such as a FlexLayer, 
// or in any situation where the Sprite scale doesn't change, but its parents scale is effecting the dimensions regardless.
func (self *PhysicsArcadeBody) SetSyncBoundsA(member bool) {
    self.Object.Set("syncBounds", member)
}

// Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) GetIsMovingA() bool{
    return self.Object.Get("isMoving").Bool()
}

// Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) SetIsMovingA(member bool) {
    self.Object.Set("isMoving", member)
}

// Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) GetStopVelocityOnCollideA() bool{
    return self.Object.Get("stopVelocityOnCollide").Bool()
}

// Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) SetStopVelocityOnCollideA(member bool) {
    self.Object.Set("stopVelocityOnCollide", member)
}

// Listen for the completion of `moveTo` or `moveFrom` events.
func (self *PhysicsArcadeBody) GetOnMoveCompleteA() *Signal{
    return &Signal{self.Object.Get("onMoveComplete")}
}

// Listen for the completion of `moveTo` or `moveFrom` events.
func (self *PhysicsArcadeBody) SetOnMoveCompleteA(member *Signal) {
    self.Object.Set("onMoveComplete", member)
}

// Optional callback. If set, invoked during the running of `moveTo` or `moveFrom` events.
func (self *PhysicsArcadeBody) SetMovementCallbackA(member func(...interface{})) {
    self.Object.Set("movementCallback", member)
}

// Context in which to call the movementCallback.
func (self *PhysicsArcadeBody) GetMovementCallbackContextA() interface{}{
    return self.Object.Get("movementCallbackContext")
}

// Context in which to call the movementCallback.
func (self *PhysicsArcadeBody) SetMovementCallbackContextA(member interface{}) {
    self.Object.Set("movementCallbackContext", member)
}

// The x position of the Body. The same as `Body.x`.
func (self *PhysicsArcadeBody) GetLeftA() int{
    return self.Object.Get("left").Int()
}

// The x position of the Body. The same as `Body.x`.
func (self *PhysicsArcadeBody) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsArcadeBody) GetRightA() int{
    return self.Object.Get("right").Int()
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsArcadeBody) SetRightA(member int) {
    self.Object.Set("right", member)
}

// The y position of the Body. The same as `Body.y`.
func (self *PhysicsArcadeBody) GetTopA() int{
    return self.Object.Get("top").Int()
}

// The y position of the Body. The same as `Body.y`.
func (self *PhysicsArcadeBody) SetTopA(member int) {
    self.Object.Set("top", member)
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsArcadeBody) GetBottomA() int{
    return self.Object.Get("bottom").Int()
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsArcadeBody) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// The x position.
func (self *PhysicsArcadeBody) GetXA() int{
    return self.Object.Get("x").Int()
}

// The x position.
func (self *PhysicsArcadeBody) SetXA(member int) {
    self.Object.Set("x", member)
}

// The y position.
func (self *PhysicsArcadeBody) GetYA() int{
    return self.Object.Get("y").Int()
}

// The y position.
func (self *PhysicsArcadeBody) SetYA(member int) {
    self.Object.Set("y", member)
}



// Internal method.
func (self *PhysicsArcadeBody) UpdateBounds() {
    self.Object.Call("updateBounds")
}

// Internal method.
func (self *PhysicsArcadeBody) UpdateBoundsI(args ...interface{}) {
    self.Object.Call("updateBounds", args)
}

// Internal method.
func (self *PhysicsArcadeBody) PreUpdate() {
    self.Object.Call("preUpdate")
}

// Internal method.
func (self *PhysicsArcadeBody) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Internal method.
func (self *PhysicsArcadeBody) UpdateMovement() {
    self.Object.Call("updateMovement")
}

// Internal method.
func (self *PhysicsArcadeBody) UpdateMovementI(args ...interface{}) {
    self.Object.Call("updateMovement", args)
}

// If this Body is moving as a result of a call to `moveTo` or `moveFrom` (i.e. it
// has Body.isMoving true), then calling this method will stop the movement before
// either the duration or distance counters expire.
// 
// The `onMoveComplete` signal is dispatched.
func (self *PhysicsArcadeBody) StopMovement() {
    self.Object.Call("stopMovement")
}

// If this Body is moving as a result of a call to `moveTo` or `moveFrom` (i.e. it
// has Body.isMoving true), then calling this method will stop the movement before
// either the duration or distance counters expire.
// 
// The `onMoveComplete` signal is dispatched.
func (self *PhysicsArcadeBody) StopMovement1O(stopVelocity bool) {
    self.Object.Call("stopMovement", stopVelocity)
}

// If this Body is moving as a result of a call to `moveTo` or `moveFrom` (i.e. it
// has Body.isMoving true), then calling this method will stop the movement before
// either the duration or distance counters expire.
// 
// The `onMoveComplete` signal is dispatched.
func (self *PhysicsArcadeBody) StopMovementI(args ...interface{}) {
    self.Object.Call("stopMovement", args)
}

// Internal method.
func (self *PhysicsArcadeBody) PostUpdate() {
    self.Object.Call("postUpdate")
}

// Internal method.
func (self *PhysicsArcadeBody) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Internal method.
func (self *PhysicsArcadeBody) CheckWorldBounds() bool{
    return self.Object.Call("checkWorldBounds").Bool()
}

// Internal method.
func (self *PhysicsArcadeBody) CheckWorldBoundsI(args ...interface{}) bool{
    return self.Object.Call("checkWorldBounds", args).Bool()
}

// Note: This method is experimental, and may be changed or removed in a future release.
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

// Note: This method is experimental, and may be changed or removed in a future release.
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

// Note: This method is experimental, and may be changed or removed in a future release.
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

// Note: This method is experimental, and may be changed or removed in a future release.
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

// Note: This method is experimental, and may be changed or removed in a future release.
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

// Note: This method is experimental, and may be changed or removed in a future release.
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

// Note: This method is experimental, and may be changed or removed in a future release.
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

// You can modify the size of the physics Body to be any dimension you need.
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

// You can modify the size of the physics Body to be any dimension you need.
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

// You can modify the size of the physics Body to be any dimension you need.
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

// You can modify the size of the physics Body to be any dimension you need.
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

// Sets this Body as using a circle, of the given radius, for all collision detection instead of a rectangle.
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

// Sets this Body as using a circle, of the given radius, for all collision detection instead of a rectangle.
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

// Sets this Body as using a circle, of the given radius, for all collision detection instead of a rectangle.
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

// Sets this Body as using a circle, of the given radius, for all collision detection instead of a rectangle.
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

// Sets this Body as using a circle, of the given radius, for all collision detection instead of a rectangle.
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

// Resets all Body values (velocity, acceleration, rotation, etc)
func (self *PhysicsArcadeBody) Reset(x int, y int) {
    self.Object.Call("reset", x, y)
}

// Resets all Body values (velocity, acceleration, rotation, etc)
func (self *PhysicsArcadeBody) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Returns the bounds of this physics body.
// 
// Only used internally by the World collision methods.
func (self *PhysicsArcadeBody) GetBounds(obj interface{}) interface{}{
    return self.Object.Call("getBounds", obj)
}

// Returns the bounds of this physics body.
// 
// Only used internally by the World collision methods.
func (self *PhysicsArcadeBody) GetBoundsI(args ...interface{}) interface{}{
    return self.Object.Call("getBounds", args)
}

// Tests if a world point lies within this Body.
func (self *PhysicsArcadeBody) HitTest(x int, y int) bool{
    return self.Object.Call("hitTest", x, y).Bool()
}

// Tests if a world point lies within this Body.
func (self *PhysicsArcadeBody) HitTestI(args ...interface{}) bool{
    return self.Object.Call("hitTest", args).Bool()
}

// Returns true if the bottom of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnFloor() bool{
    return self.Object.Call("onFloor").Bool()
}

// Returns true if the bottom of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnFloorI(args ...interface{}) bool{
    return self.Object.Call("onFloor", args).Bool()
}

// Returns true if the top of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnTop() bool{
    return self.Object.Call("onTop").Bool()
}

// Returns true if the top of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnTopI(args ...interface{}) bool{
    return self.Object.Call("onTop", args).Bool()
}

// Returns true if either side of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnWall() bool{
    return self.Object.Call("onWall").Bool()
}

// Returns true if either side of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnWallI(args ...interface{}) bool{
    return self.Object.Call("onWall", args).Bool()
}

// Returns the absolute delta x value.
func (self *PhysicsArcadeBody) DeltaAbsX() int{
    return self.Object.Call("deltaAbsX").Int()
}

// Returns the absolute delta x value.
func (self *PhysicsArcadeBody) DeltaAbsXI(args ...interface{}) int{
    return self.Object.Call("deltaAbsX", args).Int()
}

// Returns the absolute delta y value.
func (self *PhysicsArcadeBody) DeltaAbsY() int{
    return self.Object.Call("deltaAbsY").Int()
}

// Returns the absolute delta y value.
func (self *PhysicsArcadeBody) DeltaAbsYI(args ...interface{}) int{
    return self.Object.Call("deltaAbsY", args).Int()
}

// Returns the delta x value. The difference between Body.x now and in the previous step.
func (self *PhysicsArcadeBody) DeltaX() int{
    return self.Object.Call("deltaX").Int()
}

// Returns the delta x value. The difference between Body.x now and in the previous step.
func (self *PhysicsArcadeBody) DeltaXI(args ...interface{}) int{
    return self.Object.Call("deltaX", args).Int()
}

// Returns the delta y value. The difference between Body.y now and in the previous step.
func (self *PhysicsArcadeBody) DeltaY() int{
    return self.Object.Call("deltaY").Int()
}

// Returns the delta y value. The difference between Body.y now and in the previous step.
func (self *PhysicsArcadeBody) DeltaYI(args ...interface{}) int{
    return self.Object.Call("deltaY", args).Int()
}

// Returns the delta z value. The difference between Body.rotation now and in the previous step.
func (self *PhysicsArcadeBody) DeltaZ() int{
    return self.Object.Call("deltaZ").Int()
}

// Returns the delta z value. The difference between Body.rotation now and in the previous step.
func (self *PhysicsArcadeBody) DeltaZI(args ...interface{}) int{
    return self.Object.Call("deltaZ", args).Int()
}

// Destroys this Body.
// 
// First it calls Group.removeFromHash if the Game Object this Body belongs to is part of a Group.
// Then it nulls the Game Objects body reference, and nulls this Body.sprite reference.
func (self *PhysicsArcadeBody) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this Body.
// 
// First it calls Group.removeFromHash if the Game Object this Body belongs to is part of a Group.
// Then it nulls the Game Objects body reference, and nulls this Body.sprite reference.
func (self *PhysicsArcadeBody) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Render Sprite Body.
func (self *PhysicsArcadeBody) Render(context interface{}, body *PhysicsArcadeBody) {
    self.Object.Call("render", context, body)
}

// Render Sprite Body.
func (self *PhysicsArcadeBody) Render1O(context interface{}, body *PhysicsArcadeBody, color string) {
    self.Object.Call("render", context, body, color)
}

// Render Sprite Body.
func (self *PhysicsArcadeBody) Render2O(context interface{}, body *PhysicsArcadeBody, color string, filled bool) {
    self.Object.Call("render", context, body, color, filled)
}

// Render Sprite Body.
func (self *PhysicsArcadeBody) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// Render Sprite Body Physics Data as text.
func (self *PhysicsArcadeBody) RenderBodyInfo(body *PhysicsArcadeBody, x int, y int) {
    self.Object.Call("renderBodyInfo", body, x, y)
}

// Render Sprite Body Physics Data as text.
func (self *PhysicsArcadeBody) RenderBodyInfo1O(body *PhysicsArcadeBody, x int, y int, color string) {
    self.Object.Call("renderBodyInfo", body, x, y, color)
}

// Render Sprite Body Physics Data as text.
func (self *PhysicsArcadeBody) RenderBodyInfoI(args ...interface{}) {
    self.Object.Call("renderBodyInfo", args)
}
