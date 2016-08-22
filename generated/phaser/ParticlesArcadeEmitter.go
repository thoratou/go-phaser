// Automatic generation for Phaser.Particles.Arcade.Emitter
// generated file ParticlesArcadeEmitter.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Emitter is a lightweight particle emitter that uses Arcade Physics.
// It can be used for one-time explosions or for continuous effects like rain and fire.
// All it really does is launch Particle objects out at set intervals, and fixes their positions and velocities accordingly.
type ParticlesArcadeEmitter struct {
    *js.Object
}


// The total number of particles in this emitter.
func (self *ParticlesArcadeEmitter) GetMaxParticles() int{
    return self.Get("maxParticles").Int()
}

// The total number of particles in this emitter.
func (self *ParticlesArcadeEmitter) SetMaxParticles(member int) {
    self.Set("maxParticles", member)
}

// A handy string name for this emitter. Can be set to anything.
func (self *ParticlesArcadeEmitter) GetName() string{
    return self.Get("name").String()
}

// A handy string name for this emitter. Can be set to anything.
func (self *ParticlesArcadeEmitter) SetName(member string) {
    self.Set("name", member)
}

// Internal Phaser Type value.
func (self *ParticlesArcadeEmitter) GetType() int{
    return self.Get("type").Int()
}

// Internal Phaser Type value.
func (self *ParticlesArcadeEmitter) SetType(member int) {
    self.Set("type", member)
}

// The const physics body type of this object.
func (self *ParticlesArcadeEmitter) GetPhysicsType() int{
    return self.Get("physicsType").Int()
}

// The const physics body type of this object.
func (self *ParticlesArcadeEmitter) SetPhysicsType(member int) {
    self.Set("physicsType", member)
}

// The area of the emitter. Particles can be randomly generated from anywhere within this rectangle.
func (self *ParticlesArcadeEmitter) GetArea() *Rectangle{
    return &Rectangle{self.Get("area")}
}

// The area of the emitter. Particles can be randomly generated from anywhere within this rectangle.
func (self *ParticlesArcadeEmitter) SetArea(member *Rectangle) {
    self.Set("area", member)
}

// The minimum possible velocity of a particle.
func (self *ParticlesArcadeEmitter) GetMinParticleSpeed() *Point{
    return &Point{self.Get("minParticleSpeed")}
}

// The minimum possible velocity of a particle.
func (self *ParticlesArcadeEmitter) SetMinParticleSpeed(member *Point) {
    self.Set("minParticleSpeed", member)
}

// The maximum possible velocity of a particle.
func (self *ParticlesArcadeEmitter) GetMaxParticleSpeed() *Point{
    return &Point{self.Get("maxParticleSpeed")}
}

// The maximum possible velocity of a particle.
func (self *ParticlesArcadeEmitter) SetMaxParticleSpeed(member *Point) {
    self.Set("maxParticleSpeed", member)
}

// The minimum possible scale of a particle. This is applied to the X and Y axis. If you need to control each axis see minParticleScaleX.
func (self *ParticlesArcadeEmitter) GetMinParticleScale() int{
    return self.Get("minParticleScale").Int()
}

// The minimum possible scale of a particle. This is applied to the X and Y axis. If you need to control each axis see minParticleScaleX.
func (self *ParticlesArcadeEmitter) SetMinParticleScale(member int) {
    self.Set("minParticleScale", member)
}

// The maximum possible scale of a particle. This is applied to the X and Y axis. If you need to control each axis see maxParticleScaleX.
func (self *ParticlesArcadeEmitter) GetMaxParticleScale() int{
    return self.Get("maxParticleScale").Int()
}

// The maximum possible scale of a particle. This is applied to the X and Y axis. If you need to control each axis see maxParticleScaleX.
func (self *ParticlesArcadeEmitter) SetMaxParticleScale(member int) {
    self.Set("maxParticleScale", member)
}

// An array of the calculated scale easing data applied to particles with scaleRates > 0.
func (self *ParticlesArcadeEmitter) GetScaleData() []interface{}{
	array := self.Get("scaleData")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of the calculated scale easing data applied to particles with scaleRates > 0.
func (self *ParticlesArcadeEmitter) SetScaleData(member []interface{}) {
    self.Set("scaleData", member)
}

// The minimum possible angular velocity of a particle.
func (self *ParticlesArcadeEmitter) GetMinRotation() int{
    return self.Get("minRotation").Int()
}

// The minimum possible angular velocity of a particle.
func (self *ParticlesArcadeEmitter) SetMinRotation(member int) {
    self.Set("minRotation", member)
}

// The maximum possible angular velocity of a particle.
func (self *ParticlesArcadeEmitter) GetMaxRotation() int{
    return self.Get("maxRotation").Int()
}

// The maximum possible angular velocity of a particle.
func (self *ParticlesArcadeEmitter) SetMaxRotation(member int) {
    self.Set("maxRotation", member)
}

// The minimum possible alpha value of a particle.
func (self *ParticlesArcadeEmitter) GetMinParticleAlpha() int{
    return self.Get("minParticleAlpha").Int()
}

// The minimum possible alpha value of a particle.
func (self *ParticlesArcadeEmitter) SetMinParticleAlpha(member int) {
    self.Set("minParticleAlpha", member)
}

// The maximum possible alpha value of a particle.
func (self *ParticlesArcadeEmitter) GetMaxParticleAlpha() int{
    return self.Get("maxParticleAlpha").Int()
}

// The maximum possible alpha value of a particle.
func (self *ParticlesArcadeEmitter) SetMaxParticleAlpha(member int) {
    self.Set("maxParticleAlpha", member)
}

// An array of the calculated alpha easing data applied to particles with alphaRates > 0.
func (self *ParticlesArcadeEmitter) GetAlphaData() []interface{}{
	array := self.Get("alphaData")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of the calculated alpha easing data applied to particles with alphaRates > 0.
func (self *ParticlesArcadeEmitter) SetAlphaData(member []interface{}) {
    self.Set("alphaData", member)
}

// Sets the `body.gravity.y` of each particle sprite to this value on launch.
func (self *ParticlesArcadeEmitter) GetGravity() int{
    return self.Get("gravity").Int()
}

// Sets the `body.gravity.y` of each particle sprite to this value on launch.
func (self *ParticlesArcadeEmitter) SetGravity(member int) {
    self.Set("gravity", member)
}

// For emitting your own particle class types. They must extend Phaser.Particle.
func (self *ParticlesArcadeEmitter) GetParticleClass() interface{}{
    return self.Get("particleClass")
}

// For emitting your own particle class types. They must extend Phaser.Particle.
func (self *ParticlesArcadeEmitter) SetParticleClass(member interface{}) {
    self.Set("particleClass", member)
}

// The X and Y drag component of particles launched from the emitter.
func (self *ParticlesArcadeEmitter) GetParticleDrag() *Point{
    return &Point{self.Get("particleDrag")}
}

// The X and Y drag component of particles launched from the emitter.
func (self *ParticlesArcadeEmitter) SetParticleDrag(member *Point) {
    self.Set("particleDrag", member)
}

// The angular drag component of particles launched from the emitter if they are rotating.
func (self *ParticlesArcadeEmitter) GetAngularDrag() int{
    return self.Get("angularDrag").Int()
}

// The angular drag component of particles launched from the emitter if they are rotating.
func (self *ParticlesArcadeEmitter) SetAngularDrag(member int) {
    self.Set("angularDrag", member)
}

// How often a particle is emitted in ms (if emitter is started with Explode === false).
func (self *ParticlesArcadeEmitter) GetFrequency() int{
    return self.Get("frequency").Int()
}

// How often a particle is emitted in ms (if emitter is started with Explode === false).
func (self *ParticlesArcadeEmitter) SetFrequency(member int) {
    self.Set("frequency", member)
}

// How long each particle lives once it is emitted in ms. Default is 2 seconds. Set lifespan to 'zero' for particles to live forever.
func (self *ParticlesArcadeEmitter) GetLifespan() int{
    return self.Get("lifespan").Int()
}

// How long each particle lives once it is emitted in ms. Default is 2 seconds. Set lifespan to 'zero' for particles to live forever.
func (self *ParticlesArcadeEmitter) SetLifespan(member int) {
    self.Set("lifespan", member)
}

// How much each particle should bounce on each axis. 1 = full bounce, 0 = no bounce.
func (self *ParticlesArcadeEmitter) GetBounce() *Point{
    return &Point{self.Get("bounce")}
}

// How much each particle should bounce on each axis. 1 = full bounce, 0 = no bounce.
func (self *ParticlesArcadeEmitter) SetBounce(member *Point) {
    self.Set("bounce", member)
}

// Determines whether the emitter is currently emitting particles. It is totally safe to directly toggle this.
func (self *ParticlesArcadeEmitter) GetOn() bool{
    return self.Get("on").Bool()
}

// Determines whether the emitter is currently emitting particles. It is totally safe to directly toggle this.
func (self *ParticlesArcadeEmitter) SetOn(member bool) {
    self.Set("on", member)
}

// When a particle is created its anchor will be set to match this Point object (defaults to x/y: 0.5 to aid in rotation)
func (self *ParticlesArcadeEmitter) GetParticleAnchor() *Point{
    return &Point{self.Get("particleAnchor")}
}

// When a particle is created its anchor will be set to match this Point object (defaults to x/y: 0.5 to aid in rotation)
func (self *ParticlesArcadeEmitter) SetParticleAnchor(member *Point) {
    self.Set("particleAnchor", member)
}

// The blendMode as set on the particle when emitted from the Emitter. Defaults to NORMAL. Needs browser capable of supporting canvas blend-modes (most not available in WebGL)
func (self *ParticlesArcadeEmitter) GetBlendMode() int{
    return self.Get("blendMode").Int()
}

// The blendMode as set on the particle when emitted from the Emitter. Defaults to NORMAL. Needs browser capable of supporting canvas blend-modes (most not available in WebGL)
func (self *ParticlesArcadeEmitter) SetBlendMode(member int) {
    self.Set("blendMode", member)
}

// The point the particles are emitted from.
// Emitter.x and Emitter.y control the containers location, which updates all current particles
// Emitter.emitX and Emitter.emitY control the emission location relative to the x/y position.
func (self *ParticlesArcadeEmitter) GetEmitX() int{
    return self.Get("emitX").Int()
}

// The point the particles are emitted from.
// Emitter.x and Emitter.y control the containers location, which updates all current particles
// Emitter.emitX and Emitter.emitY control the emission location relative to the x/y position.
func (self *ParticlesArcadeEmitter) SetEmitX(member int) {
    self.Set("emitX", member)
}

// The point the particles are emitted from.
// Emitter.x and Emitter.y control the containers location, which updates all current particles
// Emitter.emitX and Emitter.emitY control the emission location relative to the x/y position.
func (self *ParticlesArcadeEmitter) GetEmitY() int{
    return self.Get("emitY").Int()
}

// The point the particles are emitted from.
// Emitter.x and Emitter.y control the containers location, which updates all current particles
// Emitter.emitX and Emitter.emitY control the emission location relative to the x/y position.
func (self *ParticlesArcadeEmitter) SetEmitY(member int) {
    self.Set("emitY", member)
}

// When a new Particle is emitted this controls if it will automatically scale in size. Use Emitter.setScale to configure.
func (self *ParticlesArcadeEmitter) GetAutoScale() bool{
    return self.Get("autoScale").Bool()
}

// When a new Particle is emitted this controls if it will automatically scale in size. Use Emitter.setScale to configure.
func (self *ParticlesArcadeEmitter) SetAutoScale(member bool) {
    self.Set("autoScale", member)
}

// When a new Particle is emitted this controls if it will automatically change alpha. Use Emitter.setAlpha to configure.
func (self *ParticlesArcadeEmitter) GetAutoAlpha() bool{
    return self.Get("autoAlpha").Bool()
}

// When a new Particle is emitted this controls if it will automatically change alpha. Use Emitter.setAlpha to configure.
func (self *ParticlesArcadeEmitter) SetAutoAlpha(member bool) {
    self.Set("autoAlpha", member)
}

// If this is `true` then when the Particle is emitted it will be bought to the top of the Emitters display list.
func (self *ParticlesArcadeEmitter) GetParticleBringToTop() bool{
    return self.Get("particleBringToTop").Bool()
}

// If this is `true` then when the Particle is emitted it will be bought to the top of the Emitters display list.
func (self *ParticlesArcadeEmitter) SetParticleBringToTop(member bool) {
    self.Set("particleBringToTop", member)
}

// If this is `true` then when the Particle is emitted it will be sent to the back of the Emitters display list.
func (self *ParticlesArcadeEmitter) GetParticleSendToBack() bool{
    return self.Get("particleSendToBack").Bool()
}

// If this is `true` then when the Particle is emitted it will be sent to the back of the Emitters display list.
func (self *ParticlesArcadeEmitter) SetParticleSendToBack(member bool) {
    self.Set("particleSendToBack", member)
}

// Gets or sets the width of the Emitter. This is the region in which a particle can be emitted.
func (self *ParticlesArcadeEmitter) GetWidth() int{
    return self.Get("width").Int()
}

// Gets or sets the width of the Emitter. This is the region in which a particle can be emitted.
func (self *ParticlesArcadeEmitter) SetWidth(member int) {
    self.Set("width", member)
}

// Gets or sets the height of the Emitter. This is the region in which a particle can be emitted.
func (self *ParticlesArcadeEmitter) GetHeight() int{
    return self.Get("height").Int()
}

// Gets or sets the height of the Emitter. This is the region in which a particle can be emitted.
func (self *ParticlesArcadeEmitter) SetHeight(member int) {
    self.Set("height", member)
}

// Gets or sets the x position of the Emitter.
func (self *ParticlesArcadeEmitter) GetX() int{
    return self.Get("x").Int()
}

// Gets or sets the x position of the Emitter.
func (self *ParticlesArcadeEmitter) SetX(member int) {
    self.Set("x", member)
}

// Gets or sets the y position of the Emitter.
func (self *ParticlesArcadeEmitter) GetY() int{
    return self.Get("y").Int()
}

// Gets or sets the y position of the Emitter.
func (self *ParticlesArcadeEmitter) SetY(member int) {
    self.Set("y", member)
}

// Gets the left position of the Emitter.
func (self *ParticlesArcadeEmitter) GetLeft() int{
    return self.Get("left").Int()
}

// Gets the left position of the Emitter.
func (self *ParticlesArcadeEmitter) SetLeft(member int) {
    self.Set("left", member)
}

// Gets the right position of the Emitter.
func (self *ParticlesArcadeEmitter) GetRight() int{
    return self.Get("right").Int()
}

// Gets the right position of the Emitter.
func (self *ParticlesArcadeEmitter) SetRight(member int) {
    self.Set("right", member)
}

// Gets the top position of the Emitter.
func (self *ParticlesArcadeEmitter) GetTop() int{
    return self.Get("top").Int()
}

// Gets the top position of the Emitter.
func (self *ParticlesArcadeEmitter) SetTop(member int) {
    self.Set("top", member)
}

// Gets the bottom position of the Emitter.
func (self *ParticlesArcadeEmitter) GetBottom() int{
    return self.Get("bottom").Int()
}

// Gets the bottom position of the Emitter.
func (self *ParticlesArcadeEmitter) SetBottom(member int) {
    self.Set("bottom", member)
}

// A reference to the currently running Game.
func (self *ParticlesArcadeEmitter) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *ParticlesArcadeEmitter) SetGame(member *Game) {
    self.Set("game", member)
}

// The z-depth value of this object within its parent container/Group - the World is a Group as well.
// This value must be unique for each child in a Group.
func (self *ParticlesArcadeEmitter) GetZ() int{
    return self.Get("z").Int()
}

// The z-depth value of this object within its parent container/Group - the World is a Group as well.
// This value must be unique for each child in a Group.
func (self *ParticlesArcadeEmitter) SetZ(member int) {
    self.Set("z", member)
}

// The alive property is useful for Groups that are children of other Groups and need to be included/excluded in checks like forEachAlive.
func (self *ParticlesArcadeEmitter) GetAlive() bool{
    return self.Get("alive").Bool()
}

// The alive property is useful for Groups that are children of other Groups and need to be included/excluded in checks like forEachAlive.
func (self *ParticlesArcadeEmitter) SetAlive(member bool) {
    self.Set("alive", member)
}

// If exists is true the group is updated, otherwise it is skipped.
func (self *ParticlesArcadeEmitter) GetExists() bool{
    return self.Get("exists").Bool()
}

// If exists is true the group is updated, otherwise it is skipped.
func (self *ParticlesArcadeEmitter) SetExists(member bool) {
    self.Set("exists", member)
}

// A group with `ignoreDestroy` set to `true` ignores all calls to its `destroy` method.
func (self *ParticlesArcadeEmitter) GetIgnoreDestroy() bool{
    return self.Get("ignoreDestroy").Bool()
}

// A group with `ignoreDestroy` set to `true` ignores all calls to its `destroy` method.
func (self *ParticlesArcadeEmitter) SetIgnoreDestroy(member bool) {
    self.Set("ignoreDestroy", member)
}

// A Group is that has `pendingDestroy` set to `true` is flagged to have its destroy method
// called on the next logic update.
// You can set it directly to flag the Group to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy a Group from within one of its own callbacks
// or a callback of one of its children.
func (self *ParticlesArcadeEmitter) GetPendingDestroy() bool{
    return self.Get("pendingDestroy").Bool()
}

// A Group is that has `pendingDestroy` set to `true` is flagged to have its destroy method
// called on the next logic update.
// You can set it directly to flag the Group to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy a Group from within one of its own callbacks
// or a callback of one of its children.
func (self *ParticlesArcadeEmitter) SetPendingDestroy(member bool) {
    self.Set("pendingDestroy", member)
}

// The type of objects that will be created when using {@link Phaser.Group#create create} or {@link Phaser.Group#createMultiple createMultiple}.
// 
// Any object may be used but it should extend either Sprite or Image and accept the same constructor arguments:
// when a new object is created it is passed the following parameters to its constructor: `(game, x, y, key, frame)`.
func (self *ParticlesArcadeEmitter) GetClassType() interface{}{
    return self.Get("classType")
}

// The type of objects that will be created when using {@link Phaser.Group#create create} or {@link Phaser.Group#createMultiple createMultiple}.
// 
// Any object may be used but it should extend either Sprite or Image and accept the same constructor arguments:
// when a new object is created it is passed the following parameters to its constructor: `(game, x, y, key, frame)`.
func (self *ParticlesArcadeEmitter) SetClassType(member interface{}) {
    self.Set("classType", member)
}

// The current display object that the group cursor is pointing to, if any. (Can be set manually.)
// 
// The cursor is a way to iterate through the children in a Group using {@link Phaser.Group#next next} and {@link Phaser.Group#previous previous}.
func (self *ParticlesArcadeEmitter) GetCursor() *DisplayObject{
    return &DisplayObject{self.Get("cursor")}
}

// The current display object that the group cursor is pointing to, if any. (Can be set manually.)
// 
// The cursor is a way to iterate through the children in a Group using {@link Phaser.Group#next next} and {@link Phaser.Group#previous previous}.
func (self *ParticlesArcadeEmitter) SetCursor(member *DisplayObject) {
    self.Set("cursor", member)
}

// A Group with `inputEnableChildren` set to `true` will automatically call `inputEnabled = true` 
// on any children _added_ to, or _created by_, this Group.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
func (self *ParticlesArcadeEmitter) GetInputEnableChildren() bool{
    return self.Get("inputEnableChildren").Bool()
}

// A Group with `inputEnableChildren` set to `true` will automatically call `inputEnabled = true` 
// on any children _added_ to, or _created by_, this Group.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
func (self *ParticlesArcadeEmitter) SetInputEnableChildren(member bool) {
    self.Set("inputEnableChildren", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputDown signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) GetOnChildInputDown() *Signal{
    return &Signal{self.Get("onChildInputDown")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputDown signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) SetOnChildInputDown(member *Signal) {
    self.Set("onChildInputDown", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputUp signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 3 arguments: A reference to the Sprite that triggered the signal, 
// a reference to the Pointer that caused it, and a boolean value `isOver` that tells you if the Pointer
// is still over the Sprite or not.
func (self *ParticlesArcadeEmitter) GetOnChildInputUp() *Signal{
    return &Signal{self.Get("onChildInputUp")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputUp signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 3 arguments: A reference to the Sprite that triggered the signal, 
// a reference to the Pointer that caused it, and a boolean value `isOver` that tells you if the Pointer
// is still over the Sprite or not.
func (self *ParticlesArcadeEmitter) SetOnChildInputUp(member *Signal) {
    self.Set("onChildInputUp", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputOver signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) GetOnChildInputOver() *Signal{
    return &Signal{self.Get("onChildInputOver")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputOver signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) SetOnChildInputOver(member *Signal) {
    self.Set("onChildInputOver", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputOut signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) GetOnChildInputOut() *Signal{
    return &Signal{self.Get("onChildInputOut")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputOut signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) SetOnChildInputOut(member *Signal) {
    self.Set("onChildInputOut", member)
}

// If true all Sprites created by, or added to this group, will have a physics body enabled on them.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
// 
// The default body type is controlled with {@link Phaser.Group#physicsBodyType physicsBodyType}.
func (self *ParticlesArcadeEmitter) GetEnableBody() bool{
    return self.Get("enableBody").Bool()
}

// If true all Sprites created by, or added to this group, will have a physics body enabled on them.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
// 
// The default body type is controlled with {@link Phaser.Group#physicsBodyType physicsBodyType}.
func (self *ParticlesArcadeEmitter) SetEnableBody(member bool) {
    self.Set("enableBody", member)
}

// If true when a physics body is created (via {@link Phaser.Group#enableBody enableBody}) it will create a physics debug object as well.
// 
// This only works for P2 bodies.
func (self *ParticlesArcadeEmitter) GetEnableBodyDebug() bool{
    return self.Get("enableBodyDebug").Bool()
}

// If true when a physics body is created (via {@link Phaser.Group#enableBody enableBody}) it will create a physics debug object as well.
// 
// This only works for P2 bodies.
func (self *ParticlesArcadeEmitter) SetEnableBodyDebug(member bool) {
    self.Set("enableBodyDebug", member)
}

// If {@link Phaser.Group#enableBody enableBody} is true this is the type of physics body that is created on new Sprites.
// 
// The valid values are {@link Phaser.Physics.ARCADE}, {@link Phaser.Physics.P2JS}, {@link Phaser.Physics.NINJA}, etc.
func (self *ParticlesArcadeEmitter) GetPhysicsBodyType() int{
    return self.Get("physicsBodyType").Int()
}

// If {@link Phaser.Group#enableBody enableBody} is true this is the type of physics body that is created on new Sprites.
// 
// The valid values are {@link Phaser.Physics.ARCADE}, {@link Phaser.Physics.P2JS}, {@link Phaser.Physics.NINJA}, etc.
func (self *ParticlesArcadeEmitter) SetPhysicsBodyType(member int) {
    self.Set("physicsBodyType", member)
}

// If this Group contains Arcade Physics Sprites you can set a custom sort direction via this property.
// 
// It should be set to one of the Phaser.Physics.Arcade sort direction constants: 
// 
// Phaser.Physics.Arcade.SORT_NONE
// Phaser.Physics.Arcade.LEFT_RIGHT
// Phaser.Physics.Arcade.RIGHT_LEFT
// Phaser.Physics.Arcade.TOP_BOTTOM
// Phaser.Physics.Arcade.BOTTOM_TOP
// 
// If set to `null` the Group will use whatever Phaser.Physics.Arcade.sortDirection is set to. This is the default behavior.
func (self *ParticlesArcadeEmitter) GetPhysicsSortDirection() int{
    return self.Get("physicsSortDirection").Int()
}

// If this Group contains Arcade Physics Sprites you can set a custom sort direction via this property.
// 
// It should be set to one of the Phaser.Physics.Arcade sort direction constants: 
// 
// Phaser.Physics.Arcade.SORT_NONE
// Phaser.Physics.Arcade.LEFT_RIGHT
// Phaser.Physics.Arcade.RIGHT_LEFT
// Phaser.Physics.Arcade.TOP_BOTTOM
// Phaser.Physics.Arcade.BOTTOM_TOP
// 
// If set to `null` the Group will use whatever Phaser.Physics.Arcade.sortDirection is set to. This is the default behavior.
func (self *ParticlesArcadeEmitter) SetPhysicsSortDirection(member int) {
    self.Set("physicsSortDirection", member)
}

// This signal is dispatched when the group is destroyed.
func (self *ParticlesArcadeEmitter) GetOnDestroy() *Signal{
    return &Signal{self.Get("onDestroy")}
}

// This signal is dispatched when the group is destroyed.
func (self *ParticlesArcadeEmitter) SetOnDestroy(member *Signal) {
    self.Set("onDestroy", member)
}

// The current index of the Group cursor. Advance it with Group.next.
func (self *ParticlesArcadeEmitter) GetCursorIndex() int{
    return self.Get("cursorIndex").Int()
}

// The current index of the Group cursor. Advance it with Group.next.
func (self *ParticlesArcadeEmitter) SetCursorIndex(member int) {
    self.Set("cursorIndex", member)
}

// A Group that is fixed to the camera uses its x/y coordinates as offsets from the top left of the camera. These are stored in Group.cameraOffset.
// 
// Note that the cameraOffset values are in addition to any parent in the display list.
// So if this Group was in a Group that has x: 200, then this will be added to the cameraOffset.x
func (self *ParticlesArcadeEmitter) GetFixedToCamera() bool{
    return self.Get("fixedToCamera").Bool()
}

// A Group that is fixed to the camera uses its x/y coordinates as offsets from the top left of the camera. These are stored in Group.cameraOffset.
// 
// Note that the cameraOffset values are in addition to any parent in the display list.
// So if this Group was in a Group that has x: 200, then this will be added to the cameraOffset.x
func (self *ParticlesArcadeEmitter) SetFixedToCamera(member bool) {
    self.Set("fixedToCamera", member)
}

// If this object is {@link Phaser.Group#fixedToCamera fixedToCamera} then this stores the x/y position offset relative to the top-left of the camera view.
// If the parent of this Group is also `fixedToCamera` then the offset here is in addition to that and should typically be disabled.
func (self *ParticlesArcadeEmitter) GetCameraOffset() *Point{
    return &Point{self.Get("cameraOffset")}
}

// If this object is {@link Phaser.Group#fixedToCamera fixedToCamera} then this stores the x/y position offset relative to the top-left of the camera view.
// If the parent of this Group is also `fixedToCamera` then the offset here is in addition to that and should typically be disabled.
func (self *ParticlesArcadeEmitter) SetCameraOffset(member *Point) {
    self.Set("cameraOffset", member)
}

// The hash array is an array belonging to this Group into which you can add any of its children via Group.addToHash and Group.removeFromHash.
// 
// Only children of this Group can be added to and removed from the hash.
// 
// This hash is used automatically by Phaser Arcade Physics in order to perform non z-index based destructive sorting.
// However if you don't use Arcade Physics, or this isn't a physics enabled Group, then you can use the hash to perform your own
// sorting and filtering of Group children without touching their z-index (and therefore display draw order)
func (self *ParticlesArcadeEmitter) GetHash() []interface{}{
	array := self.Get("hash")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The hash array is an array belonging to this Group into which you can add any of its children via Group.addToHash and Group.removeFromHash.
// 
// Only children of this Group can be added to and removed from the hash.
// 
// This hash is used automatically by Phaser Arcade Physics in order to perform non z-index based destructive sorting.
// However if you don't use Arcade Physics, or this isn't a physics enabled Group, then you can use the hash to perform your own
// sorting and filtering of Group children without touching their z-index (and therefore display draw order)
func (self *ParticlesArcadeEmitter) SetHash(member []interface{}) {
    self.Set("hash", member)
}

// Total number of existing children in the group.
func (self *ParticlesArcadeEmitter) GetTotal() int{
    return self.Get("total").Int()
}

// Total number of existing children in the group.
func (self *ParticlesArcadeEmitter) SetTotal(member int) {
    self.Set("total", member)
}

// Total number of children in this group, regardless of exists/alive status.
func (self *ParticlesArcadeEmitter) GetLength() int{
    return self.Get("length").Int()
}

// Total number of children in this group, regardless of exists/alive status.
func (self *ParticlesArcadeEmitter) SetLength(member int) {
    self.Set("length", member)
}

// The angle of rotation of the group container, in degrees.
// 
// This adjusts the group itself by modifying its local rotation transform.
// 
// This has no impact on the rotation/angle properties of the children, but it will update their worldTransform
// and on-screen orientation and position.
func (self *ParticlesArcadeEmitter) GetAngle() int{
    return self.Get("angle").Int()
}

// The angle of rotation of the group container, in degrees.
// 
// This adjusts the group itself by modifying its local rotation transform.
// 
// This has no impact on the rotation/angle properties of the children, but it will update their worldTransform
// and on-screen orientation and position.
func (self *ParticlesArcadeEmitter) SetAngle(member int) {
    self.Set("angle", member)
}

// The center x coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *ParticlesArcadeEmitter) GetCenterX() int{
    return self.Get("centerX").Int()
}

// The center x coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *ParticlesArcadeEmitter) SetCenterX(member int) {
    self.Set("centerX", member)
}

// The center y coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *ParticlesArcadeEmitter) GetCenterY() int{
    return self.Get("centerY").Int()
}

// The center y coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *ParticlesArcadeEmitter) SetCenterY(member int) {
    self.Set("centerY", member)
}

// The angle of rotation of the group container, in radians.
// 
// This will adjust the group container itself by modifying its rotation.
// This will have no impact on the rotation value of its children, but it will update their worldTransform and on-screen position.
func (self *ParticlesArcadeEmitter) GetRotation() int{
    return self.Get("rotation").Int()
}

// The angle of rotation of the group container, in radians.
// 
// This will adjust the group container itself by modifying its rotation.
// This will have no impact on the rotation value of its children, but it will update their worldTransform and on-screen position.
func (self *ParticlesArcadeEmitter) SetRotation(member int) {
    self.Set("rotation", member)
}

// The visible state of the group. Non-visible Groups and all of their children are not rendered.
func (self *ParticlesArcadeEmitter) GetVisible() bool{
    return self.Get("visible").Bool()
}

// The visible state of the group. Non-visible Groups and all of their children are not rendered.
func (self *ParticlesArcadeEmitter) SetVisible(member bool) {
    self.Set("visible", member)
}

// The alpha value of the group container.
func (self *ParticlesArcadeEmitter) GetAlpha() int{
    return self.Get("alpha").Int()
}

// The alpha value of the group container.
func (self *ParticlesArcadeEmitter) SetAlpha(member int) {
    self.Set("alpha", member)
}

// [read-only] The array of children of this container.
func (self *ParticlesArcadeEmitter) GetChildren() []DisplayObject{
	array := self.Get("children")
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// [read-only] The array of children of this container.
func (self *ParticlesArcadeEmitter) SetChildren(member []DisplayObject) {
    self.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *ParticlesArcadeEmitter) GetIgnoreChildInput() bool{
    return self.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *ParticlesArcadeEmitter) SetIgnoreChildInput(member bool) {
    self.Set("ignoreChildInput", member)
}



// Called automatically by the game loop, decides when to launch particles and when to "die".
func (self *ParticlesArcadeEmitter) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// This function generates a new set of particles for use by this emitter.
// The particles are stored internally waiting to be emitted via Emitter.start.
func (self *ParticlesArcadeEmitter) MakeParticlesI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("makeParticles", args)}
}

// Call this function to turn off all the particles and the emitter.
func (self *ParticlesArcadeEmitter) KillI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("kill", args)}
}

// Handy for bringing game objects "back to life". Just sets alive and exists back to true.
func (self *ParticlesArcadeEmitter) ReviveI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("revive", args)}
}

// Call this function to emit the given quantity of particles at all once (an explosion)
func (self *ParticlesArcadeEmitter) ExplodeI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("explode", args)}
}

// Call this function to start emitting a flow of particles at the given frequency.
// It will carry on going until the total given is reached.
// Each time the flow is run the quantity number of particles will be emitted together.
// If you set the total to be 20 and quantity to be 5 then flow will emit 4 times in total (4 x 5 = 20 total)
// If you set the total to be -1 then no quantity cap is used and it will keep emitting.
func (self *ParticlesArcadeEmitter) FlowI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("flow", args)}
}

// Call this function to start emitting particles.
func (self *ParticlesArcadeEmitter) StartI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("start", args)}
}

// This function is used internally to emit the next particle in the queue.
// 
// However it can also be called externally to emit a particle.
// 
// When called externally you can use the arguments to override any defaults the Emitter has set.
func (self *ParticlesArcadeEmitter) EmitParticleI(args ...interface{}) bool{
    return self.Call("emitParticle", args).Bool()
}

// Destroys this Emitter, all associated child Particles and then removes itself from the Particle Manager.
func (self *ParticlesArcadeEmitter) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// A more compact way of setting the width and height of the emitter.
func (self *ParticlesArcadeEmitter) SetSizeI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("setSize", args)}
}

// A more compact way of setting the X velocity range of the emitter.
func (self *ParticlesArcadeEmitter) SetXSpeedI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("setXSpeed", args)}
}

// A more compact way of setting the Y velocity range of the emitter.
func (self *ParticlesArcadeEmitter) SetYSpeedI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("setYSpeed", args)}
}

// A more compact way of setting the angular velocity constraints of the particles.
func (self *ParticlesArcadeEmitter) SetRotationI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("setRotation", args)}
}

// A more compact way of setting the alpha constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed at which the Particle change in alpha from min to max.
// If rate is zero, which is the default, the particle won't change alpha - instead it will pick a random alpha between min and max on emit.
func (self *ParticlesArcadeEmitter) SetAlphaI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("setAlpha", args)}
}

// A more compact way of setting the scale constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed and ease which the Particle uses to change in scale from min to max across both axis.
// If rate is zero, which is the default, the particle won't change scale during update, instead it will pick a random scale between min and max on emit.
func (self *ParticlesArcadeEmitter) SetScaleI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("setScale", args)}
}

// Change the emitters center to match the center of any object with a `center` property, such as a Sprite.
// If the object doesn't have a center property it will be set to object.x + object.width / 2
func (self *ParticlesArcadeEmitter) AtI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Call("at", args)}
}

// Adds an existing object as the top child in this group.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If the child was already in this Group, it is simply returned, and nothing else happens to it.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
// 
// Use {@link Phaser.Group#addAt addAt} to control where a child is added. Use {@link Phaser.Group#create create} to create and add a new child.
func (self *ParticlesArcadeEmitter) AddI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("add", args)}
}

// Adds an existing object to this group.
// 
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) AddAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addAt", args)}
}

// Adds a child of this Group into the hash array.
// This call will return false if the child is not a child of this Group, or is already in the hash.
func (self *ParticlesArcadeEmitter) AddToHashI(args ...interface{}) bool{
    return self.Call("addToHash", args).Bool()
}

// Removes a child of this Group from the hash array.
// This call will return false if the child is not in the hash.
func (self *ParticlesArcadeEmitter) RemoveFromHashI(args ...interface{}) bool{
    return self.Call("removeFromHash", args).Bool()
}

// Adds an array of existing Display Objects to this Group.
// 
// The Display Objects are automatically added to the top of this Group, and will render on-top of everything already in this Group.
// 
// As well as an array you can also pass another Group as the first argument. In this case all of the children from that
// Group will be removed from it and added into this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) AddMultipleI(args ...interface{}) interface{}{
    return self.Call("addMultiple", args)
}

// Returns the child found at the given index within this group.
func (self *ParticlesArcadeEmitter) GetAtI(args ...interface{}) interface{}{
    return self.Call("getAt", args)
}

// Creates a new Phaser.Sprite object and adds it to the top of this group.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) CreateI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("create", args)}
}

// Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
// 
// This method is useful if you need to quickly generate a pool of sprites, such as bullets.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// You can provide an array as the `key` and / or `frame` arguments. When you do this
// it will create `quantity` Sprites for every key (and frame) in the arrays.
// 
// For example:
// 
// `createMultiple(25, ['ball', 'carrot'])`
// 
// In the above code there are 2 keys (ball and carrot) which means that 50 sprites will be
// created in total, 25 of each. You can also have the `frame` as an array:
// 
// `createMultiple(5, 'bricks', [0, 1, 2, 3])`
// 
// In the above there is one key (bricks), which is a sprite sheet. The frames array tells
// this method to use frames 0, 1, 2 and 3. So in total it will create 20 sprites, because
// the quantity was set to 5, so that is 5 brick sprites of frame 0, 5 brick sprites with
// frame 1, and so on.
// 
// If you set both the key and frame arguments to be arrays then understand it will create
// a total quantity of sprites equal to the size of both arrays times each other. I.e.:
// 
// `createMultiple(20, ['diamonds', 'balls'], [0, 1, 2])`
// 
// The above will create 20 'diamonds' of frame 0, 20 with frame 1 and 20 with frame 2.
// It will then create 20 'balls' of frame 0, 20 with frame 1 and 20 with frame 2.
// In total it will have created 120 sprites.
// 
// By default the Sprites will have their `exists` property set to `false`, and they will be 
// positioned at 0x0, relative to the `Group.x / y` values.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) CreateMultipleI(args ...interface{}) []interface{}{
	array := self.Call("createMultiple", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Internal method that re-applies all of the children's Z values.
// 
// This must be called whenever children ordering is altered so that their `z` indices are correctly updated.
func (self *ParticlesArcadeEmitter) UpdateZI(args ...interface{}) {
    self.Call("updateZ", args)
}

// This method iterates through all children in the Group (regardless if they are visible or exist)
// and then changes their position so they are arranged in a Grid formation. Children must have
// the `alignTo` method in order to be positioned by this call. All default Phaser Game Objects have
// this.
// 
// The grid dimensions are determined by the first four arguments. The `rows` and `columns` arguments
// relate to the width and height of the grid respectively.
// 
// For example if the Group had 100 children in it:
// 
// `Group.align(10, 10, 32, 32)`
// 
// This will align all of the children into a grid formation of 10x10, using 32 pixels per
// grid cell. If you want a wider grid, you could do:
// 
// `Group.align(25, 4, 32, 32)`
// 
// This will align the children into a grid of 25x4, again using 32 pixels per grid cell.
// 
// You can choose to set _either_ the `rows` or `columns` value to -1. Doing so tells the method
// to keep on aligning children until there are no children left. For example if this Group had
// 48 children in it, the following:
// 
// `Group.align(-1, 8, 32, 32)`
// 
// ... will align the children so that there are 8 columns vertically (the second argument), 
// and each row will contain 6 sprites, except the last one, which will contain 5 (totaling 48)
// 
// You can also do:
// 
// `Group.align(10, -1, 32, 32)`
// 
// In this case it will create a grid 10 wide, and as tall as it needs to be in order to fit
// all of the children in.
// 
// The `position` property allows you to control where in each grid cell the child is positioned.
// This is a constant and can be one of `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, 
// `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, 
// `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` or `Phaser.BOTTOM_RIGHT`.
// 
// The final argument; `offset` lets you start the alignment from a specific child index.
func (self *ParticlesArcadeEmitter) AlignI(args ...interface{}) {
    self.Call("align", args)
}

// Sets the group cursor to the first child in the group.
// 
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *ParticlesArcadeEmitter) ResetCursorI(args ...interface{}) interface{}{
    return self.Call("resetCursor", args)
}

// Advances the group cursor to the next (higher) object in the group.
// 
// If the cursor is at the end of the group (top child) it is moved the start of the group (bottom child).
func (self *ParticlesArcadeEmitter) NextI(args ...interface{}) interface{}{
    return self.Call("next", args)
}

// Moves the group cursor to the previous (lower) child in the group.
// 
// If the cursor is at the start of the group (bottom child) it is moved to the end (top child).
func (self *ParticlesArcadeEmitter) PreviousI(args ...interface{}) interface{}{
    return self.Call("previous", args)
}

// Swaps the position of two children in this group.
// 
// Both children must be in this group, a child cannot be swapped with itself, and unparented children cannot be swapped.
func (self *ParticlesArcadeEmitter) SwapI(args ...interface{}) {
    self.Call("swap", args)
}

// Brings the given child to the top of this group so it renders above all other children.
func (self *ParticlesArcadeEmitter) BringToTopI(args ...interface{}) interface{}{
    return self.Call("bringToTop", args)
}

// Sends the given child to the bottom of this group so it renders below all other children.
func (self *ParticlesArcadeEmitter) SendToBackI(args ...interface{}) interface{}{
    return self.Call("sendToBack", args)
}

// Moves the given child up one place in this group unless it's already at the top.
func (self *ParticlesArcadeEmitter) MoveUpI(args ...interface{}) interface{}{
    return self.Call("moveUp", args)
}

// Moves the given child down one place in this group unless it's already at the bottom.
func (self *ParticlesArcadeEmitter) MoveDownI(args ...interface{}) interface{}{
    return self.Call("moveDown", args)
}

// Positions the child found at the given index within this group to the given x and y coordinates.
func (self *ParticlesArcadeEmitter) XyI(args ...interface{}) {
    self.Call("xy", args)
}

// Reverses all children in this group.
// 
// This operation applies only to immediate children and does not propagate to subgroups.
func (self *ParticlesArcadeEmitter) ReverseI(args ...interface{}) {
    self.Call("reverse", args)
}

// Get the index position of the given child in this group, which should match the child's `z` property.
func (self *ParticlesArcadeEmitter) GetIndexI(args ...interface{}) int{
    return self.Call("getIndex", args).Int()
}

// Searches the Group for the first instance of a child with the `name`
// property matching the given argument. Should more than one child have
// the same name only the first instance is returned.
func (self *ParticlesArcadeEmitter) GetByNameI(args ...interface{}) interface{}{
    return self.Call("getByName", args)
}

// Replaces a child of this Group with the given newChild. The newChild cannot be a member of this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) ReplaceI(args ...interface{}) interface{}{
    return self.Call("replace", args)
}

// Checks if the child has the given property.
// 
// Will scan up to 4 levels deep only.
func (self *ParticlesArcadeEmitter) HasPropertyI(args ...interface{}) bool{
    return self.Call("hasProperty", args).Bool()
}

// Sets a property to the given value on the child. The operation parameter controls how the value is set.
// 
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *ParticlesArcadeEmitter) SetPropertyI(args ...interface{}) bool{
    return self.Call("setProperty", args).Bool()
}

// Checks a property for the given value on the child.
func (self *ParticlesArcadeEmitter) CheckPropertyI(args ...interface{}) bool{
    return self.Call("checkProperty", args).Bool()
}

// Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetI(args ...interface{}) bool{
    return self.Call("set", args).Bool()
}

// Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAllI(args ...interface{}) {
    self.Call("setAll", args)
}

// Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAllChildrenI(args ...interface{}) {
    self.Call("setAllChildren", args)
}

// Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *ParticlesArcadeEmitter) CheckAllI(args ...interface{}) {
    self.Call("checkAll", args)
}

// Adds the amount to the given property on all children in this group.
// 
// `Group.addAll('x', 10)` will add 10 to the child.x value for each child.
func (self *ParticlesArcadeEmitter) AddAllI(args ...interface{}) {
    self.Call("addAll", args)
}

// Subtracts the amount from the given property on all children in this group.
// 
// `Group.subAll('x', 10)` will minus 10 from the child.x value for each child.
func (self *ParticlesArcadeEmitter) SubAllI(args ...interface{}) {
    self.Call("subAll", args)
}

// Multiplies the given property by the amount on all children in this group.
// 
// `Group.multiplyAll('x', 2)` will x2 the child.x value for each child.
func (self *ParticlesArcadeEmitter) MultiplyAllI(args ...interface{}) {
    self.Call("multiplyAll", args)
}

// Divides the given property by the amount on all children in this group.
// 
// `Group.divideAll('x', 2)` will half the child.x value for each child.
func (self *ParticlesArcadeEmitter) DivideAllI(args ...interface{}) {
    self.Call("divideAll", args)
}

// Calls a function, specified by name, on all children in the group who exist (or do not exist).
// 
// After the existsValue parameter you can add as many parameters as you like, which will all be passed to the child callback.
func (self *ParticlesArcadeEmitter) CallAllExistsI(args ...interface{}) {
    self.Call("callAllExists", args)
}

// Returns a reference to a function that exists on a child of the group based on the given callback array.
func (self *ParticlesArcadeEmitter) CallbackFromArrayI(args ...interface{}) {
    self.Call("callbackFromArray", args)
}

// Calls a function, specified by name, on all on children.
// 
// The function is called for all children regardless if they are dead or alive (see callAllExists for different options).
// After the method parameter and context you can add as many extra parameters as you like, which will all be passed to the child.
func (self *ParticlesArcadeEmitter) CallAllI(args ...interface{}) {
    self.Call("callAll", args)
}

// The core preUpdate - as called by World.
func (self *ParticlesArcadeEmitter) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// The core postUpdate - as called by World.
func (self *ParticlesArcadeEmitter) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}

// Find children matching a certain predicate.
// 
// For example:
// 
//     var healthyList = Group.filter(function(child, index, children) {
//         return child.health > 10 ? true : false;
//     }, true);
//     healthyList.callAll('attack');
// 
// Note: Currently this will skip any children which are Groups themselves.
func (self *ParticlesArcadeEmitter) FilterI(args ...interface{}) *ArraySet{
    return &ArraySet{self.Call("filter", args)}
}

// Call a function on each child in this group.
// 
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
// 
//     Group.forEach(awardBonusGold, this, true, 100, 500)
// 
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
// 
// Note: This check will skip any children which are Groups themselves.
func (self *ParticlesArcadeEmitter) ForEachI(args ...interface{}) {
    self.Call("forEach", args)
}

// Call a function on each existing child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachExistsI(args ...interface{}) {
    self.Call("forEachExists", args)
}

// Call a function on each alive child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachAliveI(args ...interface{}) {
    self.Call("forEachAlive", args)
}

// Call a function on each dead child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachDeadI(args ...interface{}) {
    self.Call("forEachDead", args)
}

// Sort the children in the group according to a particular key and ordering.
// 
// Call this function to sort the group according to a particular key value and order.
// 
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
// 
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *ParticlesArcadeEmitter) SortI(args ...interface{}) {
    self.Call("sort", args)
}

// Sort the children in the group according to custom sort function.
// 
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *ParticlesArcadeEmitter) CustomSortI(args ...interface{}) {
    self.Call("customSort", args)
}

// An internal helper function for the sort process.
func (self *ParticlesArcadeEmitter) AscendingSortHandlerI(args ...interface{}) {
    self.Call("ascendingSortHandler", args)
}

// An internal helper function for the sort process.
func (self *ParticlesArcadeEmitter) DescendingSortHandlerI(args ...interface{}) {
    self.Call("descendingSortHandler", args)
}

// Iterates over the children of the group performing one of several actions for matched children.
// 
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
// 
// The result depends on the `returnType`:
// 
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
// 
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *ParticlesArcadeEmitter) IterateI(args ...interface{}) interface{}{
    return self.Call("iterate", args)
}

// Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstExistsI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getFirstExists", args)}
}

// Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstAliveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getFirstAlive", args)}
}

// Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstDeadI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getFirstDead", args)}
}

// Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *ParticlesArcadeEmitter) ResetChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("resetChild", args)}
}

// Return the child at the top of this group.
// 
// The top child is the child displayed (rendered) above every other child.
func (self *ParticlesArcadeEmitter) GetTopI(args ...interface{}) interface{}{
    return self.Call("getTop", args)
}

// Returns the child at the bottom of this group.
// 
// The bottom child the child being displayed (rendered) below every other child.
func (self *ParticlesArcadeEmitter) GetBottomI(args ...interface{}) interface{}{
    return self.Call("getBottom", args)
}

// Get the closest child to given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *ParticlesArcadeEmitter) GetClosestToI(args ...interface{}) interface{}{
    return self.Call("getClosestTo", args)
}

// Get the child furthest away from the given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *ParticlesArcadeEmitter) GetFurthestFromI(args ...interface{}) interface{}{
    return self.Call("getFurthestFrom", args)
}

// Get the number of living children in this group.
func (self *ParticlesArcadeEmitter) CountLivingI(args ...interface{}) int{
    return self.Call("countLiving", args).Int()
}

// Get the number of dead children in this group.
func (self *ParticlesArcadeEmitter) CountDeadI(args ...interface{}) int{
    return self.Call("countDead", args).Int()
}

// Returns a random child from the group.
func (self *ParticlesArcadeEmitter) GetRandomI(args ...interface{}) interface{}{
    return self.Call("getRandom", args)
}

// Removes the given child from this group.
// 
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
// 
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *ParticlesArcadeEmitter) RemoveI(args ...interface{}) bool{
    return self.Call("remove", args).Bool()
}

// Moves all children from this Group to the Group given.
func (self *ParticlesArcadeEmitter) MoveAllI(args ...interface{}) *Group{
    return &Group{self.Call("moveAll", args)}
}

// Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *ParticlesArcadeEmitter) RemoveAllI(args ...interface{}) {
    self.Call("removeAll", args)
}

// Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *ParticlesArcadeEmitter) RemoveBetweenI(args ...interface{}) {
    self.Call("removeBetween", args)
}

// Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *ParticlesArcadeEmitter) AlignInI(args ...interface{}) *Group{
    return &Group{self.Call("alignIn", args)}
}

// Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *ParticlesArcadeEmitter) AlignToI(args ...interface{}) *Group{
    return &Group{self.Call("alignTo", args)}
}

// Adds a child to the container.
func (self *ParticlesArcadeEmitter) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *ParticlesArcadeEmitter) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *ParticlesArcadeEmitter) SwapChildrenI(args ...interface{}) {
    self.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *ParticlesArcadeEmitter) GetChildIndexI(args ...interface{}) int{
    return self.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *ParticlesArcadeEmitter) SetChildIndexI(args ...interface{}) {
    self.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *ParticlesArcadeEmitter) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *ParticlesArcadeEmitter) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *ParticlesArcadeEmitter) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *ParticlesArcadeEmitter) RemoveChildrenI(args ...interface{}) {
    self.Call("removeChildren", args)
}

// Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *ParticlesArcadeEmitter) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getBounds", args)}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *ParticlesArcadeEmitter) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *ParticlesArcadeEmitter) SetStageReferenceI(args ...interface{}) {
    self.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *ParticlesArcadeEmitter) RemoveStageReferenceI(args ...interface{}) {
    self.Call("removeStageReference", args)
}

// Renders the object using the WebGL renderer
func (self *ParticlesArcadeEmitter) _renderWebGLI(args ...interface{}) {
    self.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *ParticlesArcadeEmitter) _renderCanvasI(args ...interface{}) {
    self.Call("_renderCanvas", args)
}
