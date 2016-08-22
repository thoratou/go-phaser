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
func (self *PhysicsArcadeBody) GetSprite() *Sprite{
    return &Sprite{self.Get("sprite")}
}

// Reference to the parent Sprite.
func (self *PhysicsArcadeBody) SetSprite(member *Sprite) {
    self.Set("sprite", member)
}

// Local reference to game.
func (self *PhysicsArcadeBody) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsArcadeBody) SetGame(member *Game) {
    self.Set("game", member)
}

// The type of physics system this body belongs to.
func (self *PhysicsArcadeBody) GetType() int{
    return self.Get("type").Int()
}

// The type of physics system this body belongs to.
func (self *PhysicsArcadeBody) SetType(member int) {
    self.Set("type", member)
}

// A disabled body won't be checked for any form of collision or overlap or have its pre/post updates run.
func (self *PhysicsArcadeBody) GetEnable() bool{
    return self.Get("enable").Bool()
}

// A disabled body won't be checked for any form of collision or overlap or have its pre/post updates run.
func (self *PhysicsArcadeBody) SetEnable(member bool) {
    self.Set("enable", member)
}

// If `true` this Body is using circular collision detection. If `false` it is using rectangular.
// Use `Body.setCircle` to control the collision shape this Body uses.
func (self *PhysicsArcadeBody) GetIsCircle() bool{
    return self.Get("isCircle").Bool()
}

// If `true` this Body is using circular collision detection. If `false` it is using rectangular.
// Use `Body.setCircle` to control the collision shape this Body uses.
func (self *PhysicsArcadeBody) SetIsCircle(member bool) {
    self.Set("isCircle", member)
}

// The radius of the circular collision shape this Body is using if Body.setCircle has been enabled.
// If you wish to change the radius then call `setCircle` again with the new value.
// If you wish to stop the Body using a circle then call `setCircle` with a radius of zero (or undefined).
func (self *PhysicsArcadeBody) GetRadius() int{
    return self.Get("radius").Int()
}

// The radius of the circular collision shape this Body is using if Body.setCircle has been enabled.
// If you wish to change the radius then call `setCircle` again with the new value.
// If you wish to stop the Body using a circle then call `setCircle` with a radius of zero (or undefined).
func (self *PhysicsArcadeBody) SetRadius(member int) {
    self.Set("radius", member)
}

// The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsArcadeBody) GetOffset() *Point{
    return &Point{self.Get("offset")}
}

// The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsArcadeBody) SetOffset(member *Point) {
    self.Set("offset", member)
}

// The position of the physics body.
func (self *PhysicsArcadeBody) GetPosition() *Point{
    return &Point{self.Get("position")}
}

// The position of the physics body.
func (self *PhysicsArcadeBody) SetPosition(member *Point) {
    self.Set("position", member)
}

// The previous position of the physics body.
func (self *PhysicsArcadeBody) GetPrev() *Point{
    return &Point{self.Get("prev")}
}

// The previous position of the physics body.
func (self *PhysicsArcadeBody) SetPrev(member *Point) {
    self.Set("prev", member)
}

// Allow this Body to be rotated? (via angularVelocity, etc)
func (self *PhysicsArcadeBody) GetAllowRotation() bool{
    return self.Get("allowRotation").Bool()
}

// Allow this Body to be rotated? (via angularVelocity, etc)
func (self *PhysicsArcadeBody) SetAllowRotation(member bool) {
    self.Set("allowRotation", member)
}

// An Arcade Physics Body can have angularVelocity and angularAcceleration. Please understand that the collision Body
// itself never rotates, it is always axis-aligned. However these values are passed up to the parent Sprite and updates its rotation.
func (self *PhysicsArcadeBody) GetRotation() int{
    return self.Get("rotation").Int()
}

// An Arcade Physics Body can have angularVelocity and angularAcceleration. Please understand that the collision Body
// itself never rotates, it is always axis-aligned. However these values are passed up to the parent Sprite and updates its rotation.
func (self *PhysicsArcadeBody) SetRotation(member int) {
    self.Set("rotation", member)
}

// The previous rotation of the physics body.
func (self *PhysicsArcadeBody) GetPreRotation() int{
    return self.Get("preRotation").Int()
}

// The previous rotation of the physics body.
func (self *PhysicsArcadeBody) SetPreRotation(member int) {
    self.Set("preRotation", member)
}

// The calculated width of the physics body.
func (self *PhysicsArcadeBody) GetWidth() int{
    return self.Get("width").Int()
}

// The calculated width of the physics body.
func (self *PhysicsArcadeBody) SetWidth(member int) {
    self.Set("width", member)
}

// The calculated height of the physics body.
func (self *PhysicsArcadeBody) GetHeight() int{
    return self.Get("height").Int()
}

// The calculated height of the physics body.
func (self *PhysicsArcadeBody) SetHeight(member int) {
    self.Set("height", member)
}

// The un-scaled original size.
func (self *PhysicsArcadeBody) GetSourceWidth() int{
    return self.Get("sourceWidth").Int()
}

// The un-scaled original size.
func (self *PhysicsArcadeBody) SetSourceWidth(member int) {
    self.Set("sourceWidth", member)
}

// The un-scaled original size.
func (self *PhysicsArcadeBody) GetSourceHeight() int{
    return self.Get("sourceHeight").Int()
}

// The un-scaled original size.
func (self *PhysicsArcadeBody) SetSourceHeight(member int) {
    self.Set("sourceHeight", member)
}

// The calculated width / 2 of the physics body.
func (self *PhysicsArcadeBody) GetHalfWidth() int{
    return self.Get("halfWidth").Int()
}

// The calculated width / 2 of the physics body.
func (self *PhysicsArcadeBody) SetHalfWidth(member int) {
    self.Set("halfWidth", member)
}

// The calculated height / 2 of the physics body.
func (self *PhysicsArcadeBody) GetHalfHeight() int{
    return self.Get("halfHeight").Int()
}

// The calculated height / 2 of the physics body.
func (self *PhysicsArcadeBody) SetHalfHeight(member int) {
    self.Set("halfHeight", member)
}

// The center coordinate of the Physics Body.
func (self *PhysicsArcadeBody) GetCenter() *Point{
    return &Point{self.Get("center")}
}

// The center coordinate of the Physics Body.
func (self *PhysicsArcadeBody) SetCenter(member *Point) {
    self.Set("center", member)
}

// The velocity, or rate of change in speed of the Body. Measured in pixels per second.
func (self *PhysicsArcadeBody) GetVelocity() *Point{
    return &Point{self.Get("velocity")}
}

// The velocity, or rate of change in speed of the Body. Measured in pixels per second.
func (self *PhysicsArcadeBody) SetVelocity(member *Point) {
    self.Set("velocity", member)
}

// The new velocity. Calculated during the Body.preUpdate and applied to its position.
func (self *PhysicsArcadeBody) GetNewVelocity() *Point{
    return &Point{self.Get("newVelocity")}
}

// The new velocity. Calculated during the Body.preUpdate and applied to its position.
func (self *PhysicsArcadeBody) SetNewVelocity(member *Point) {
    self.Set("newVelocity", member)
}

// The Sprite position is updated based on the delta x/y values. You can set a cap on those (both +-) using deltaMax.
func (self *PhysicsArcadeBody) GetDeltaMax() *Point{
    return &Point{self.Get("deltaMax")}
}

// The Sprite position is updated based on the delta x/y values. You can set a cap on those (both +-) using deltaMax.
func (self *PhysicsArcadeBody) SetDeltaMax(member *Point) {
    self.Set("deltaMax", member)
}

// The acceleration is the rate of change of the velocity. Measured in pixels per second squared.
func (self *PhysicsArcadeBody) GetAcceleration() *Point{
    return &Point{self.Get("acceleration")}
}

// The acceleration is the rate of change of the velocity. Measured in pixels per second squared.
func (self *PhysicsArcadeBody) SetAcceleration(member *Point) {
    self.Set("acceleration", member)
}

// The drag applied to the motion of the Body.
func (self *PhysicsArcadeBody) GetDrag() *Point{
    return &Point{self.Get("drag")}
}

// The drag applied to the motion of the Body.
func (self *PhysicsArcadeBody) SetDrag(member *Point) {
    self.Set("drag", member)
}

// Allow this Body to be influenced by gravity? Either world or local.
func (self *PhysicsArcadeBody) GetAllowGravity() bool{
    return self.Get("allowGravity").Bool()
}

// Allow this Body to be influenced by gravity? Either world or local.
func (self *PhysicsArcadeBody) SetAllowGravity(member bool) {
    self.Set("allowGravity", member)
}

// A local gravity applied to this Body. If non-zero this over rides any world gravity, unless Body.allowGravity is set to false.
func (self *PhysicsArcadeBody) GetGravity() *Point{
    return &Point{self.Get("gravity")}
}

// A local gravity applied to this Body. If non-zero this over rides any world gravity, unless Body.allowGravity is set to false.
func (self *PhysicsArcadeBody) SetGravity(member *Point) {
    self.Set("gravity", member)
}

// The elasticity of the Body when colliding. bounce.x/y = 1 means full rebound, bounce.x/y = 0.5 means 50% rebound velocity.
func (self *PhysicsArcadeBody) GetBounce() *Point{
    return &Point{self.Get("bounce")}
}

// The elasticity of the Body when colliding. bounce.x/y = 1 means full rebound, bounce.x/y = 0.5 means 50% rebound velocity.
func (self *PhysicsArcadeBody) SetBounce(member *Point) {
    self.Set("bounce", member)
}

// The elasticity of the Body when colliding with the World bounds.
// By default this property is `null`, in which case `Body.bounce` is used instead. Set this property
// to a Phaser.Point object in order to enable a World bounds specific bounce value.
func (self *PhysicsArcadeBody) GetWorldBounce() *Point{
    return &Point{self.Get("worldBounce")}
}

// The elasticity of the Body when colliding with the World bounds.
// By default this property is `null`, in which case `Body.bounce` is used instead. Set this property
// to a Phaser.Point object in order to enable a World bounds specific bounce value.
func (self *PhysicsArcadeBody) SetWorldBounce(member *Point) {
    self.Set("worldBounce", member)
}

// A Signal that is dispatched when this Body collides with the world bounds.
// Due to the potentially high volume of signals this could create it is disabled by default.
// To use this feature set this property to a Phaser.Signal: `sprite.body.onWorldBounds = new Phaser.Signal()`
// and it will be called when a collision happens, passing five arguments:
// `onWorldBounds(sprite, up, down, left, right)`
// where the Sprite is a reference to the Sprite that owns this Body, and the other arguments are booleans
// indicating on which side of the world the Body collided.
func (self *PhysicsArcadeBody) GetOnWorldBounds() *Signal{
    return &Signal{self.Get("onWorldBounds")}
}

// A Signal that is dispatched when this Body collides with the world bounds.
// Due to the potentially high volume of signals this could create it is disabled by default.
// To use this feature set this property to a Phaser.Signal: `sprite.body.onWorldBounds = new Phaser.Signal()`
// and it will be called when a collision happens, passing five arguments:
// `onWorldBounds(sprite, up, down, left, right)`
// where the Sprite is a reference to the Sprite that owns this Body, and the other arguments are booleans
// indicating on which side of the world the Body collided.
func (self *PhysicsArcadeBody) SetOnWorldBounds(member *Signal) {
    self.Set("onWorldBounds", member)
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
func (self *PhysicsArcadeBody) GetOnCollide() *Signal{
    return &Signal{self.Get("onCollide")}
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
func (self *PhysicsArcadeBody) SetOnCollide(member *Signal) {
    self.Set("onCollide", member)
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
func (self *PhysicsArcadeBody) GetOnOverlap() *Signal{
    return &Signal{self.Get("onOverlap")}
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
func (self *PhysicsArcadeBody) SetOnOverlap(member *Signal) {
    self.Set("onOverlap", member)
}

// The maximum velocity in pixels per second sq. that the Body can reach.
func (self *PhysicsArcadeBody) GetMaxVelocity() *Point{
    return &Point{self.Get("maxVelocity")}
}

// The maximum velocity in pixels per second sq. that the Body can reach.
func (self *PhysicsArcadeBody) SetMaxVelocity(member *Point) {
    self.Set("maxVelocity", member)
}

// The amount of movement that will occur if another object 'rides' this one.
func (self *PhysicsArcadeBody) GetFriction() *Point{
    return &Point{self.Get("friction")}
}

// The amount of movement that will occur if another object 'rides' this one.
func (self *PhysicsArcadeBody) SetFriction(member *Point) {
    self.Set("friction", member)
}

// The angular velocity controls the rotation speed of the Body. It is measured in radians per second.
func (self *PhysicsArcadeBody) GetAngularVelocity() int{
    return self.Get("angularVelocity").Int()
}

// The angular velocity controls the rotation speed of the Body. It is measured in radians per second.
func (self *PhysicsArcadeBody) SetAngularVelocity(member int) {
    self.Set("angularVelocity", member)
}

// The angular acceleration is the rate of change of the angular velocity. Measured in radians per second squared.
func (self *PhysicsArcadeBody) GetAngularAcceleration() int{
    return self.Get("angularAcceleration").Int()
}

// The angular acceleration is the rate of change of the angular velocity. Measured in radians per second squared.
func (self *PhysicsArcadeBody) SetAngularAcceleration(member int) {
    self.Set("angularAcceleration", member)
}

// The drag applied during the rotation of the Body.
func (self *PhysicsArcadeBody) GetAngularDrag() int{
    return self.Get("angularDrag").Int()
}

// The drag applied during the rotation of the Body.
func (self *PhysicsArcadeBody) SetAngularDrag(member int) {
    self.Set("angularDrag", member)
}

// The maximum angular velocity in radians per second that the Body can reach.
func (self *PhysicsArcadeBody) GetMaxAngular() int{
    return self.Get("maxAngular").Int()
}

// The maximum angular velocity in radians per second that the Body can reach.
func (self *PhysicsArcadeBody) SetMaxAngular(member int) {
    self.Set("maxAngular", member)
}

// The mass of the Body. When two bodies collide their mass is used in the calculation to determine the exchange of velocity.
func (self *PhysicsArcadeBody) GetMass() int{
    return self.Get("mass").Int()
}

// The mass of the Body. When two bodies collide their mass is used in the calculation to determine the exchange of velocity.
func (self *PhysicsArcadeBody) SetMass(member int) {
    self.Set("mass", member)
}

// The angle of the Body in radians, as calculated by its angularVelocity.
func (self *PhysicsArcadeBody) GetAngle() int{
    return self.Get("angle").Int()
}

// The angle of the Body in radians, as calculated by its angularVelocity.
func (self *PhysicsArcadeBody) SetAngle(member int) {
    self.Set("angle", member)
}

// The speed of the Body as calculated by its velocity.
func (self *PhysicsArcadeBody) GetSpeed() int{
    return self.Get("speed").Int()
}

// The speed of the Body as calculated by its velocity.
func (self *PhysicsArcadeBody) SetSpeed(member int) {
    self.Set("speed", member)
}

// A const reference to the direction the Body is traveling or facing.
func (self *PhysicsArcadeBody) GetFacing() int{
    return self.Get("facing").Int()
}

// A const reference to the direction the Body is traveling or facing.
func (self *PhysicsArcadeBody) SetFacing(member int) {
    self.Set("facing", member)
}

// An immovable Body will not receive any impacts from other bodies.
func (self *PhysicsArcadeBody) GetImmovable() bool{
    return self.Get("immovable").Bool()
}

// An immovable Body will not receive any impacts from other bodies.
func (self *PhysicsArcadeBody) SetImmovable(member bool) {
    self.Set("immovable", member)
}

// If you have a Body that is being moved around the world via a tween or a Group motion, but its local x/y position never
// actually changes, then you should set Body.moves = false. Otherwise it will most likely fly off the screen.
// If you want the physics system to move the body around, then set moves to true. Set to true to allow the Physics system to move this Body, otherwise false to move it manually.
func (self *PhysicsArcadeBody) GetMoves() bool{
    return self.Get("moves").Bool()
}

// If you have a Body that is being moved around the world via a tween or a Group motion, but its local x/y position never
// actually changes, then you should set Body.moves = false. Otherwise it will most likely fly off the screen.
// If you want the physics system to move the body around, then set moves to true. Set to true to allow the Physics system to move this Body, otherwise false to move it manually.
func (self *PhysicsArcadeBody) SetMoves(member bool) {
    self.Set("moves", member)
}

// This flag allows you to disable the custom x separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) GetCustomSeparateX() bool{
    return self.Get("customSeparateX").Bool()
}

// This flag allows you to disable the custom x separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) SetCustomSeparateX(member bool) {
    self.Set("customSeparateX", member)
}

// This flag allows you to disable the custom y separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) GetCustomSeparateY() bool{
    return self.Get("customSeparateY").Bool()
}

// This flag allows you to disable the custom y separation that takes place by Physics.Arcade.separate.
// Used in combination with your own collision processHandler you can create whatever type of collision response you need. Use a custom separation system or the built-in one?
func (self *PhysicsArcadeBody) SetCustomSeparateY(member bool) {
    self.Set("customSeparateY", member)
}

// When this body collides with another, the amount of overlap is stored here. The amount of horizontal overlap during the collision.
func (self *PhysicsArcadeBody) GetOverlapX() int{
    return self.Get("overlapX").Int()
}

// When this body collides with another, the amount of overlap is stored here. The amount of horizontal overlap during the collision.
func (self *PhysicsArcadeBody) SetOverlapX(member int) {
    self.Set("overlapX", member)
}

// When this body collides with another, the amount of overlap is stored here. The amount of vertical overlap during the collision.
func (self *PhysicsArcadeBody) GetOverlapY() int{
    return self.Get("overlapY").Int()
}

// When this body collides with another, the amount of overlap is stored here. The amount of vertical overlap during the collision.
func (self *PhysicsArcadeBody) SetOverlapY(member int) {
    self.Set("overlapY", member)
}

// If `Body.isCircle` is true, and this body collides with another circular body, the amount of overlap is stored here. The amount of overlap during the collision.
func (self *PhysicsArcadeBody) GetOverlapR() int{
    return self.Get("overlapR").Int()
}

// If `Body.isCircle` is true, and this body collides with another circular body, the amount of overlap is stored here. The amount of overlap during the collision.
func (self *PhysicsArcadeBody) SetOverlapR(member int) {
    self.Set("overlapR", member)
}

// If a body is overlapping with another body, but neither of them are moving (maybe they spawned on-top of each other?) this is set to true. Body embed value.
func (self *PhysicsArcadeBody) GetEmbedded() bool{
    return self.Get("embedded").Bool()
}

// If a body is overlapping with another body, but neither of them are moving (maybe they spawned on-top of each other?) this is set to true. Body embed value.
func (self *PhysicsArcadeBody) SetEmbedded(member bool) {
    self.Set("embedded", member)
}

// A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsArcadeBody) GetCollideWorldBounds() bool{
    return self.Get("collideWorldBounds").Bool()
}

// A Body can be set to collide against the World bounds automatically and rebound back into the World if this is set to true. Otherwise it will leave the World. Should the Body collide with the World bounds?
func (self *PhysicsArcadeBody) SetCollideWorldBounds(member bool) {
    self.Set("collideWorldBounds", member)
}

// Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up. An object containing allowed collision.
func (self *PhysicsArcadeBody) GetCheckCollision() interface{}{
    return self.Get("checkCollision")
}

// Set the checkCollision properties to control which directions collision is processed for this Body.
// For example checkCollision.up = false means it won't collide when the collision happened while moving up. An object containing allowed collision.
func (self *PhysicsArcadeBody) SetCheckCollision(member interface{}) {
    self.Set("checkCollision", member)
}

// This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsArcadeBody) GetTouching() interface{}{
    return self.Get("touching")
}

// This object is populated with boolean values when the Body collides with another.
// touching.up = true means the collision happened to the top of this Body for example. An object containing touching results.
func (self *PhysicsArcadeBody) SetTouching(member interface{}) {
    self.Set("touching", member)
}

// This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsArcadeBody) GetWasTouching() interface{}{
    return self.Get("wasTouching")
}

// This object is populated with previous touching values from the bodies previous collision. An object containing previous touching results.
func (self *PhysicsArcadeBody) SetWasTouching(member interface{}) {
    self.Set("wasTouching", member)
}

// This object is populated with boolean values when the Body collides with the World bounds or a Tile.
// For example if blocked.up is true then the Body cannot move up. An object containing on which faces this Body is blocked from moving, if any.
func (self *PhysicsArcadeBody) GetBlocked() interface{}{
    return self.Get("blocked")
}

// This object is populated with boolean values when the Body collides with the World bounds or a Tile.
// For example if blocked.up is true then the Body cannot move up. An object containing on which faces this Body is blocked from moving, if any.
func (self *PhysicsArcadeBody) SetBlocked(member interface{}) {
    self.Set("blocked", member)
}

// If this is an especially small or fast moving object then it can sometimes skip over tilemap collisions if it moves through a tile in a step.
// Set this padding value to add extra padding to its bounds. tilePadding.x applied to its width, y to its height. Extra padding to be added to this sprite's dimensions when checking for tile collision.
func (self *PhysicsArcadeBody) GetTilePadding() *Point{
    return &Point{self.Get("tilePadding")}
}

// If this is an especially small or fast moving object then it can sometimes skip over tilemap collisions if it moves through a tile in a step.
// Set this padding value to add extra padding to its bounds. tilePadding.x applied to its width, y to its height. Extra padding to be added to this sprite's dimensions when checking for tile collision.
func (self *PhysicsArcadeBody) SetTilePadding(member *Point) {
    self.Set("tilePadding", member)
}

// If this Body in a preUpdate (true) or postUpdate (false) state?
func (self *PhysicsArcadeBody) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// If this Body in a preUpdate (true) or postUpdate (false) state?
func (self *PhysicsArcadeBody) SetDirty(member bool) {
    self.Set("dirty", member)
}

// If true and you collide this Sprite against a Group, it will disable the collision check from using a QuadTree.
func (self *PhysicsArcadeBody) GetSkipQuadTree() bool{
    return self.Get("skipQuadTree").Bool()
}

// If true and you collide this Sprite against a Group, it will disable the collision check from using a QuadTree.
func (self *PhysicsArcadeBody) SetSkipQuadTree(member bool) {
    self.Set("skipQuadTree", member)
}

// If true the Body will check itself against the Sprite.getBounds() dimensions and adjust its width and height accordingly.
// If false it will compare its dimensions against the Sprite scale instead, and adjust its width height if the scale has changed.
// Typically you would need to enable syncBounds if your sprite is the child of a responsive display object such as a FlexLayer, 
// or in any situation where the Sprite scale doesn't change, but its parents scale is effecting the dimensions regardless.
func (self *PhysicsArcadeBody) GetSyncBounds() bool{
    return self.Get("syncBounds").Bool()
}

// If true the Body will check itself against the Sprite.getBounds() dimensions and adjust its width and height accordingly.
// If false it will compare its dimensions against the Sprite scale instead, and adjust its width height if the scale has changed.
// Typically you would need to enable syncBounds if your sprite is the child of a responsive display object such as a FlexLayer, 
// or in any situation where the Sprite scale doesn't change, but its parents scale is effecting the dimensions regardless.
func (self *PhysicsArcadeBody) SetSyncBounds(member bool) {
    self.Set("syncBounds", member)
}

// Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) GetIsMoving() bool{
    return self.Get("isMoving").Bool()
}

// Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) SetIsMoving(member bool) {
    self.Set("isMoving", member)
}

// Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) GetStopVelocityOnCollide() bool{
    return self.Get("stopVelocityOnCollide").Bool()
}

// Set by the `moveTo` and `moveFrom` methods.
func (self *PhysicsArcadeBody) SetStopVelocityOnCollide(member bool) {
    self.Set("stopVelocityOnCollide", member)
}

// Listen for the completion of `moveTo` or `moveFrom` events.
func (self *PhysicsArcadeBody) GetOnMoveComplete() *Signal{
    return &Signal{self.Get("onMoveComplete")}
}

// Listen for the completion of `moveTo` or `moveFrom` events.
func (self *PhysicsArcadeBody) SetOnMoveComplete(member *Signal) {
    self.Set("onMoveComplete", member)
}

// Optional callback. If set, invoked during the running of `moveTo` or `moveFrom` events.
func (self *PhysicsArcadeBody) SetMovementCallback(member func(...interface{})) {
    self.Set("movementCallback", member)
}

// Context in which to call the movementCallback.
func (self *PhysicsArcadeBody) GetMovementCallbackContext() interface{}{
    return self.Get("movementCallbackContext")
}

// Context in which to call the movementCallback.
func (self *PhysicsArcadeBody) SetMovementCallbackContext(member interface{}) {
    self.Set("movementCallbackContext", member)
}

// The x position of the Body. The same as `Body.x`.
func (self *PhysicsArcadeBody) GetLeft() int{
    return self.Get("left").Int()
}

// The x position of the Body. The same as `Body.x`.
func (self *PhysicsArcadeBody) SetLeft(member int) {
    self.Set("left", member)
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsArcadeBody) GetRight() int{
    return self.Get("right").Int()
}

// The right value of this Body (same as Body.x + Body.width)
func (self *PhysicsArcadeBody) SetRight(member int) {
    self.Set("right", member)
}

// The y position of the Body. The same as `Body.y`.
func (self *PhysicsArcadeBody) GetTop() int{
    return self.Get("top").Int()
}

// The y position of the Body. The same as `Body.y`.
func (self *PhysicsArcadeBody) SetTop(member int) {
    self.Set("top", member)
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsArcadeBody) GetBottom() int{
    return self.Get("bottom").Int()
}

// The bottom value of this Body (same as Body.y + Body.height)
func (self *PhysicsArcadeBody) SetBottom(member int) {
    self.Set("bottom", member)
}

// The x position.
func (self *PhysicsArcadeBody) GetX() int{
    return self.Get("x").Int()
}

// The x position.
func (self *PhysicsArcadeBody) SetX(member int) {
    self.Set("x", member)
}

// The y position.
func (self *PhysicsArcadeBody) GetY() int{
    return self.Get("y").Int()
}

// The y position.
func (self *PhysicsArcadeBody) SetY(member int) {
    self.Set("y", member)
}



// Internal method.
func (self *PhysicsArcadeBody) UpdateBoundsI(args ...interface{}) {
    self.Call("updateBounds", args)
}

// Internal method.
func (self *PhysicsArcadeBody) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Internal method.
func (self *PhysicsArcadeBody) UpdateMovementI(args ...interface{}) {
    self.Call("updateMovement", args)
}

// If this Body is moving as a result of a call to `moveTo` or `moveFrom` (i.e. it
// has Body.isMoving true), then calling this method will stop the movement before
// either the duration or distance counters expire.
// 
// The `onMoveComplete` signal is dispatched.
func (self *PhysicsArcadeBody) StopMovementI(args ...interface{}) {
    self.Call("stopMovement", args)
}

// Internal method.
func (self *PhysicsArcadeBody) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}

// Internal method.
func (self *PhysicsArcadeBody) CheckWorldBoundsI(args ...interface{}) bool{
    return self.Call("checkWorldBounds", args).Bool()
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
    return self.Call("moveFrom", args).Bool()
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
    return self.Call("moveTo", args).Bool()
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
    self.Call("setSize", args)
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
    self.Call("setCircle", args)
}

// Resets all Body values (velocity, acceleration, rotation, etc)
func (self *PhysicsArcadeBody) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Returns the bounds of this physics body.
// 
// Only used internally by the World collision methods.
func (self *PhysicsArcadeBody) GetBoundsI(args ...interface{}) interface{}{
    return self.Call("getBounds", args)
}

// Tests if a world point lies within this Body.
func (self *PhysicsArcadeBody) HitTestI(args ...interface{}) bool{
    return self.Call("hitTest", args).Bool()
}

// Returns true if the bottom of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnFloorI(args ...interface{}) bool{
    return self.Call("onFloor", args).Bool()
}

// Returns true if the top of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnTopI(args ...interface{}) bool{
    return self.Call("onTop", args).Bool()
}

// Returns true if either side of this Body is in contact with either the world bounds or a tile.
func (self *PhysicsArcadeBody) OnWallI(args ...interface{}) bool{
    return self.Call("onWall", args).Bool()
}

// Returns the absolute delta x value.
func (self *PhysicsArcadeBody) DeltaAbsXI(args ...interface{}) int{
    return self.Call("deltaAbsX", args).Int()
}

// Returns the absolute delta y value.
func (self *PhysicsArcadeBody) DeltaAbsYI(args ...interface{}) int{
    return self.Call("deltaAbsY", args).Int()
}

// Returns the delta x value. The difference between Body.x now and in the previous step.
func (self *PhysicsArcadeBody) DeltaXI(args ...interface{}) int{
    return self.Call("deltaX", args).Int()
}

// Returns the delta y value. The difference between Body.y now and in the previous step.
func (self *PhysicsArcadeBody) DeltaYI(args ...interface{}) int{
    return self.Call("deltaY", args).Int()
}

// Returns the delta z value. The difference between Body.rotation now and in the previous step.
func (self *PhysicsArcadeBody) DeltaZI(args ...interface{}) int{
    return self.Call("deltaZ", args).Int()
}

// Destroys this Body.
// 
// First it calls Group.removeFromHash if the Game Object this Body belongs to is part of a Group.
// Then it nulls the Game Objects body reference, and nulls this Body.sprite reference.
func (self *PhysicsArcadeBody) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Render Sprite Body.
func (self *PhysicsArcadeBody) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// Render Sprite Body Physics Data as text.
func (self *PhysicsArcadeBody) RenderBodyInfoI(args ...interface{}) {
    self.Call("renderBodyInfo", args)
}
