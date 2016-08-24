// Package phaser Automatic generation for Phaser.Particles.Arcade.Emitter
// generated file ParticlesArcadeEmitter.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ParticlesArcadeEmitter Emitter is a lightweight particle emitter that uses Arcade Physics.
// It can be used for one-time explosions or for continuous effects like rain and fire.
// All it really does is launch Particle objects out at set intervals, and fixes their positions and velocities accordingly.
type ParticlesArcadeEmitter struct {
    *js.Object
}

// NewParticlesArcadeEmitter Emitter is a lightweight particle emitter that uses Arcade Physics.
// It can be used for one-time explosions or for continuous effects like rain and fire.
// All it really does is launch Particle objects out at set intervals, and fixes their positions and velocities accordingly.
func NewParticlesArcadeEmitter(game *Game) *ParticlesArcadeEmitter {
    return &ParticlesArcadeEmitter{js.Global.Get("Phaser").Get("Particles").Get("Arcade").Get("Emitter").New(game)}
}
// NewParticlesArcadeEmitter1O Emitter is a lightweight particle emitter that uses Arcade Physics.
// It can be used for one-time explosions or for continuous effects like rain and fire.
// All it really does is launch Particle objects out at set intervals, and fixes their positions and velocities accordingly.
func NewParticlesArcadeEmitter1O(game *Game, x int) *ParticlesArcadeEmitter {
    return &ParticlesArcadeEmitter{js.Global.Get("Phaser").Get("Particles").Get("Arcade").Get("Emitter").New(game, x)}
}
// NewParticlesArcadeEmitter2O Emitter is a lightweight particle emitter that uses Arcade Physics.
// It can be used for one-time explosions or for continuous effects like rain and fire.
// All it really does is launch Particle objects out at set intervals, and fixes their positions and velocities accordingly.
func NewParticlesArcadeEmitter2O(game *Game, x int, y int) *ParticlesArcadeEmitter {
    return &ParticlesArcadeEmitter{js.Global.Get("Phaser").Get("Particles").Get("Arcade").Get("Emitter").New(game, x, y)}
}
// NewParticlesArcadeEmitter3O Emitter is a lightweight particle emitter that uses Arcade Physics.
// It can be used for one-time explosions or for continuous effects like rain and fire.
// All it really does is launch Particle objects out at set intervals, and fixes their positions and velocities accordingly.
func NewParticlesArcadeEmitter3O(game *Game, x int, y int, maxParticles int) *ParticlesArcadeEmitter {
    return &ParticlesArcadeEmitter{js.Global.Get("Phaser").Get("Particles").Get("Arcade").Get("Emitter").New(game, x, y, maxParticles)}
}
// NewParticlesArcadeEmitterI Emitter is a lightweight particle emitter that uses Arcade Physics.
// It can be used for one-time explosions or for continuous effects like rain and fire.
// All it really does is launch Particle objects out at set intervals, and fixes their positions and velocities accordingly.
func NewParticlesArcadeEmitterI(args ...interface{}) *ParticlesArcadeEmitter {
    return &ParticlesArcadeEmitter{js.Global.Get("Phaser").Get("Particles").Get("Arcade").Get("Emitter").New(args)}
}



// MaxParticles The total number of particles in this emitter.
func (self *ParticlesArcadeEmitter) MaxParticles() int{
    return self.Object.Get("maxParticles").Int()
}

// SetMaxParticlesA The total number of particles in this emitter.
func (self *ParticlesArcadeEmitter) SetMaxParticlesA(member int) {
    self.Object.Set("maxParticles", member)
}

// Name A handy string name for this emitter. Can be set to anything.
func (self *ParticlesArcadeEmitter) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA A handy string name for this emitter. Can be set to anything.
func (self *ParticlesArcadeEmitter) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Type Internal Phaser Type value.
func (self *ParticlesArcadeEmitter) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA Internal Phaser Type value.
func (self *ParticlesArcadeEmitter) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// PhysicsType The const physics body type of this object.
func (self *ParticlesArcadeEmitter) PhysicsType() int{
    return self.Object.Get("physicsType").Int()
}

// SetPhysicsTypeA The const physics body type of this object.
func (self *ParticlesArcadeEmitter) SetPhysicsTypeA(member int) {
    self.Object.Set("physicsType", member)
}

// Area The area of the emitter. Particles can be randomly generated from anywhere within this rectangle.
func (self *ParticlesArcadeEmitter) Area() *Rectangle{
    return &Rectangle{self.Object.Get("area")}
}

// SetAreaA The area of the emitter. Particles can be randomly generated from anywhere within this rectangle.
func (self *ParticlesArcadeEmitter) SetAreaA(member *Rectangle) {
    self.Object.Set("area", member)
}

// MinParticleSpeed The minimum possible velocity of a particle.
func (self *ParticlesArcadeEmitter) MinParticleSpeed() *Point{
    return &Point{self.Object.Get("minParticleSpeed")}
}

// SetMinParticleSpeedA The minimum possible velocity of a particle.
func (self *ParticlesArcadeEmitter) SetMinParticleSpeedA(member *Point) {
    self.Object.Set("minParticleSpeed", member)
}

// MaxParticleSpeed The maximum possible velocity of a particle.
func (self *ParticlesArcadeEmitter) MaxParticleSpeed() *Point{
    return &Point{self.Object.Get("maxParticleSpeed")}
}

// SetMaxParticleSpeedA The maximum possible velocity of a particle.
func (self *ParticlesArcadeEmitter) SetMaxParticleSpeedA(member *Point) {
    self.Object.Set("maxParticleSpeed", member)
}

// MinParticleScale The minimum possible scale of a particle. This is applied to the X and Y axis. If you need to control each axis see minParticleScaleX.
func (self *ParticlesArcadeEmitter) MinParticleScale() int{
    return self.Object.Get("minParticleScale").Int()
}

// SetMinParticleScaleA The minimum possible scale of a particle. This is applied to the X and Y axis. If you need to control each axis see minParticleScaleX.
func (self *ParticlesArcadeEmitter) SetMinParticleScaleA(member int) {
    self.Object.Set("minParticleScale", member)
}

// MaxParticleScale The maximum possible scale of a particle. This is applied to the X and Y axis. If you need to control each axis see maxParticleScaleX.
func (self *ParticlesArcadeEmitter) MaxParticleScale() int{
    return self.Object.Get("maxParticleScale").Int()
}

// SetMaxParticleScaleA The maximum possible scale of a particle. This is applied to the X and Y axis. If you need to control each axis see maxParticleScaleX.
func (self *ParticlesArcadeEmitter) SetMaxParticleScaleA(member int) {
    self.Object.Set("maxParticleScale", member)
}

// ScaleData An array of the calculated scale easing data applied to particles with scaleRates > 0.
func (self *ParticlesArcadeEmitter) ScaleData() []interface{}{
	array00 := self.Object.Get("scaleData")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetScaleDataA An array of the calculated scale easing data applied to particles with scaleRates > 0.
func (self *ParticlesArcadeEmitter) SetScaleDataA(member []interface{}) {
    self.Object.Set("scaleData", member)
}

// MinRotation The minimum possible angular velocity of a particle.
func (self *ParticlesArcadeEmitter) MinRotation() int{
    return self.Object.Get("minRotation").Int()
}

// SetMinRotationA The minimum possible angular velocity of a particle.
func (self *ParticlesArcadeEmitter) SetMinRotationA(member int) {
    self.Object.Set("minRotation", member)
}

// MaxRotation The maximum possible angular velocity of a particle.
func (self *ParticlesArcadeEmitter) MaxRotation() int{
    return self.Object.Get("maxRotation").Int()
}

// SetMaxRotationA The maximum possible angular velocity of a particle.
func (self *ParticlesArcadeEmitter) SetMaxRotationA(member int) {
    self.Object.Set("maxRotation", member)
}

// MinParticleAlpha The minimum possible alpha value of a particle.
func (self *ParticlesArcadeEmitter) MinParticleAlpha() int{
    return self.Object.Get("minParticleAlpha").Int()
}

// SetMinParticleAlphaA The minimum possible alpha value of a particle.
func (self *ParticlesArcadeEmitter) SetMinParticleAlphaA(member int) {
    self.Object.Set("minParticleAlpha", member)
}

// MaxParticleAlpha The maximum possible alpha value of a particle.
func (self *ParticlesArcadeEmitter) MaxParticleAlpha() int{
    return self.Object.Get("maxParticleAlpha").Int()
}

// SetMaxParticleAlphaA The maximum possible alpha value of a particle.
func (self *ParticlesArcadeEmitter) SetMaxParticleAlphaA(member int) {
    self.Object.Set("maxParticleAlpha", member)
}

// AlphaData An array of the calculated alpha easing data applied to particles with alphaRates > 0.
func (self *ParticlesArcadeEmitter) AlphaData() []interface{}{
	array00 := self.Object.Get("alphaData")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetAlphaDataA An array of the calculated alpha easing data applied to particles with alphaRates > 0.
func (self *ParticlesArcadeEmitter) SetAlphaDataA(member []interface{}) {
    self.Object.Set("alphaData", member)
}

// Gravity Sets the `body.gravity.y` of each particle sprite to this value on launch.
func (self *ParticlesArcadeEmitter) Gravity() int{
    return self.Object.Get("gravity").Int()
}

// SetGravityA Sets the `body.gravity.y` of each particle sprite to this value on launch.
func (self *ParticlesArcadeEmitter) SetGravityA(member int) {
    self.Object.Set("gravity", member)
}

// ParticleClass For emitting your own particle class types. They must extend Phaser.Particle.
func (self *ParticlesArcadeEmitter) ParticleClass() interface{}{
    return self.Object.Get("particleClass")
}

// SetParticleClassA For emitting your own particle class types. They must extend Phaser.Particle.
func (self *ParticlesArcadeEmitter) SetParticleClassA(member interface{}) {
    self.Object.Set("particleClass", member)
}

// ParticleDrag The X and Y drag component of particles launched from the emitter.
func (self *ParticlesArcadeEmitter) ParticleDrag() *Point{
    return &Point{self.Object.Get("particleDrag")}
}

// SetParticleDragA The X and Y drag component of particles launched from the emitter.
func (self *ParticlesArcadeEmitter) SetParticleDragA(member *Point) {
    self.Object.Set("particleDrag", member)
}

// AngularDrag The angular drag component of particles launched from the emitter if they are rotating.
func (self *ParticlesArcadeEmitter) AngularDrag() int{
    return self.Object.Get("angularDrag").Int()
}

// SetAngularDragA The angular drag component of particles launched from the emitter if they are rotating.
func (self *ParticlesArcadeEmitter) SetAngularDragA(member int) {
    self.Object.Set("angularDrag", member)
}

// Frequency How often a particle is emitted in ms (if emitter is started with Explode === false).
func (self *ParticlesArcadeEmitter) Frequency() int{
    return self.Object.Get("frequency").Int()
}

// SetFrequencyA How often a particle is emitted in ms (if emitter is started with Explode === false).
func (self *ParticlesArcadeEmitter) SetFrequencyA(member int) {
    self.Object.Set("frequency", member)
}

// Lifespan How long each particle lives once it is emitted in ms. Default is 2 seconds. Set lifespan to 'zero' for particles to live forever.
func (self *ParticlesArcadeEmitter) Lifespan() int{
    return self.Object.Get("lifespan").Int()
}

// SetLifespanA How long each particle lives once it is emitted in ms. Default is 2 seconds. Set lifespan to 'zero' for particles to live forever.
func (self *ParticlesArcadeEmitter) SetLifespanA(member int) {
    self.Object.Set("lifespan", member)
}

// Bounce How much each particle should bounce on each axis. 1 = full bounce, 0 = no bounce.
func (self *ParticlesArcadeEmitter) Bounce() *Point{
    return &Point{self.Object.Get("bounce")}
}

// SetBounceA How much each particle should bounce on each axis. 1 = full bounce, 0 = no bounce.
func (self *ParticlesArcadeEmitter) SetBounceA(member *Point) {
    self.Object.Set("bounce", member)
}

// On Determines whether the emitter is currently emitting particles. It is totally safe to directly toggle this.
func (self *ParticlesArcadeEmitter) On() bool{
    return self.Object.Get("on").Bool()
}

// SetOnA Determines whether the emitter is currently emitting particles. It is totally safe to directly toggle this.
func (self *ParticlesArcadeEmitter) SetOnA(member bool) {
    self.Object.Set("on", member)
}

// ParticleAnchor When a particle is created its anchor will be set to match this Point object (defaults to x/y: 0.5 to aid in rotation)
func (self *ParticlesArcadeEmitter) ParticleAnchor() *Point{
    return &Point{self.Object.Get("particleAnchor")}
}

// SetParticleAnchorA When a particle is created its anchor will be set to match this Point object (defaults to x/y: 0.5 to aid in rotation)
func (self *ParticlesArcadeEmitter) SetParticleAnchorA(member *Point) {
    self.Object.Set("particleAnchor", member)
}

// BlendMode The blendMode as set on the particle when emitted from the Emitter. Defaults to NORMAL. Needs browser capable of supporting canvas blend-modes (most not available in WebGL)
func (self *ParticlesArcadeEmitter) BlendMode() int{
    return self.Object.Get("blendMode").Int()
}

// SetBlendModeA The blendMode as set on the particle when emitted from the Emitter. Defaults to NORMAL. Needs browser capable of supporting canvas blend-modes (most not available in WebGL)
func (self *ParticlesArcadeEmitter) SetBlendModeA(member int) {
    self.Object.Set("blendMode", member)
}

// EmitX The point the particles are emitted from.
// Emitter.x and Emitter.y control the containers location, which updates all current particles
// Emitter.emitX and Emitter.emitY control the emission location relative to the x/y position.
func (self *ParticlesArcadeEmitter) EmitX() int{
    return self.Object.Get("emitX").Int()
}

// SetEmitXA The point the particles are emitted from.
// Emitter.x and Emitter.y control the containers location, which updates all current particles
// Emitter.emitX and Emitter.emitY control the emission location relative to the x/y position.
func (self *ParticlesArcadeEmitter) SetEmitXA(member int) {
    self.Object.Set("emitX", member)
}

// EmitY The point the particles are emitted from.
// Emitter.x and Emitter.y control the containers location, which updates all current particles
// Emitter.emitX and Emitter.emitY control the emission location relative to the x/y position.
func (self *ParticlesArcadeEmitter) EmitY() int{
    return self.Object.Get("emitY").Int()
}

// SetEmitYA The point the particles are emitted from.
// Emitter.x and Emitter.y control the containers location, which updates all current particles
// Emitter.emitX and Emitter.emitY control the emission location relative to the x/y position.
func (self *ParticlesArcadeEmitter) SetEmitYA(member int) {
    self.Object.Set("emitY", member)
}

// AutoScale When a new Particle is emitted this controls if it will automatically scale in size. Use Emitter.setScale to configure.
func (self *ParticlesArcadeEmitter) AutoScale() bool{
    return self.Object.Get("autoScale").Bool()
}

// SetAutoScaleA When a new Particle is emitted this controls if it will automatically scale in size. Use Emitter.setScale to configure.
func (self *ParticlesArcadeEmitter) SetAutoScaleA(member bool) {
    self.Object.Set("autoScale", member)
}

// AutoAlpha When a new Particle is emitted this controls if it will automatically change alpha. Use Emitter.setAlpha to configure.
func (self *ParticlesArcadeEmitter) AutoAlpha() bool{
    return self.Object.Get("autoAlpha").Bool()
}

// SetAutoAlphaA When a new Particle is emitted this controls if it will automatically change alpha. Use Emitter.setAlpha to configure.
func (self *ParticlesArcadeEmitter) SetAutoAlphaA(member bool) {
    self.Object.Set("autoAlpha", member)
}

// ParticleBringToTop If this is `true` then when the Particle is emitted it will be bought to the top of the Emitters display list.
func (self *ParticlesArcadeEmitter) ParticleBringToTop() bool{
    return self.Object.Get("particleBringToTop").Bool()
}

// SetParticleBringToTopA If this is `true` then when the Particle is emitted it will be bought to the top of the Emitters display list.
func (self *ParticlesArcadeEmitter) SetParticleBringToTopA(member bool) {
    self.Object.Set("particleBringToTop", member)
}

// ParticleSendToBack If this is `true` then when the Particle is emitted it will be sent to the back of the Emitters display list.
func (self *ParticlesArcadeEmitter) ParticleSendToBack() bool{
    return self.Object.Get("particleSendToBack").Bool()
}

// SetParticleSendToBackA If this is `true` then when the Particle is emitted it will be sent to the back of the Emitters display list.
func (self *ParticlesArcadeEmitter) SetParticleSendToBackA(member bool) {
    self.Object.Set("particleSendToBack", member)
}

// Width Gets or sets the width of the Emitter. This is the region in which a particle can be emitted.
func (self *ParticlesArcadeEmitter) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA Gets or sets the width of the Emitter. This is the region in which a particle can be emitted.
func (self *ParticlesArcadeEmitter) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height Gets or sets the height of the Emitter. This is the region in which a particle can be emitted.
func (self *ParticlesArcadeEmitter) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA Gets or sets the height of the Emitter. This is the region in which a particle can be emitted.
func (self *ParticlesArcadeEmitter) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// X Gets or sets the x position of the Emitter.
func (self *ParticlesArcadeEmitter) X() int{
    return self.Object.Get("x").Int()
}

// SetXA Gets or sets the x position of the Emitter.
func (self *ParticlesArcadeEmitter) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y Gets or sets the y position of the Emitter.
func (self *ParticlesArcadeEmitter) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA Gets or sets the y position of the Emitter.
func (self *ParticlesArcadeEmitter) SetYA(member int) {
    self.Object.Set("y", member)
}

// Left Gets the left position of the Emitter.
func (self *ParticlesArcadeEmitter) Left() int{
    return self.Object.Get("left").Int()
}

// SetLeftA Gets the left position of the Emitter.
func (self *ParticlesArcadeEmitter) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// Right Gets the right position of the Emitter.
func (self *ParticlesArcadeEmitter) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA Gets the right position of the Emitter.
func (self *ParticlesArcadeEmitter) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Top Gets the top position of the Emitter.
func (self *ParticlesArcadeEmitter) Top() int{
    return self.Object.Get("top").Int()
}

// SetTopA Gets the top position of the Emitter.
func (self *ParticlesArcadeEmitter) SetTopA(member int) {
    self.Object.Set("top", member)
}

// Bottom Gets the bottom position of the Emitter.
func (self *ParticlesArcadeEmitter) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA Gets the bottom position of the Emitter.
func (self *ParticlesArcadeEmitter) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// Game A reference to the currently running Game.
func (self *ParticlesArcadeEmitter) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *ParticlesArcadeEmitter) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Z The z-depth value of this object within its parent container/Group - the World is a Group as well.
// This value must be unique for each child in a Group.
func (self *ParticlesArcadeEmitter) Z() int{
    return self.Object.Get("z").Int()
}

// SetZA The z-depth value of this object within its parent container/Group - the World is a Group as well.
// This value must be unique for each child in a Group.
func (self *ParticlesArcadeEmitter) SetZA(member int) {
    self.Object.Set("z", member)
}

// Alive The alive property is useful for Groups that are children of other Groups and need to be included/excluded in checks like forEachAlive.
func (self *ParticlesArcadeEmitter) Alive() bool{
    return self.Object.Get("alive").Bool()
}

// SetAliveA The alive property is useful for Groups that are children of other Groups and need to be included/excluded in checks like forEachAlive.
func (self *ParticlesArcadeEmitter) SetAliveA(member bool) {
    self.Object.Set("alive", member)
}

// Exists If exists is true the group is updated, otherwise it is skipped.
func (self *ParticlesArcadeEmitter) Exists() bool{
    return self.Object.Get("exists").Bool()
}

// SetExistsA If exists is true the group is updated, otherwise it is skipped.
func (self *ParticlesArcadeEmitter) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// IgnoreDestroy A group with `ignoreDestroy` set to `true` ignores all calls to its `destroy` method.
func (self *ParticlesArcadeEmitter) IgnoreDestroy() bool{
    return self.Object.Get("ignoreDestroy").Bool()
}

// SetIgnoreDestroyA A group with `ignoreDestroy` set to `true` ignores all calls to its `destroy` method.
func (self *ParticlesArcadeEmitter) SetIgnoreDestroyA(member bool) {
    self.Object.Set("ignoreDestroy", member)
}

// PendingDestroy A Group is that has `pendingDestroy` set to `true` is flagged to have its destroy method
// called on the next logic update.
// You can set it directly to flag the Group to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy a Group from within one of its own callbacks
// or a callback of one of its children.
func (self *ParticlesArcadeEmitter) PendingDestroy() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// SetPendingDestroyA A Group is that has `pendingDestroy` set to `true` is flagged to have its destroy method
// called on the next logic update.
// You can set it directly to flag the Group to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy a Group from within one of its own callbacks
// or a callback of one of its children.
func (self *ParticlesArcadeEmitter) SetPendingDestroyA(member bool) {
    self.Object.Set("pendingDestroy", member)
}

// ClassType The type of objects that will be created when using {@link Phaser.Group#create create} or {@link Phaser.Group#createMultiple createMultiple}.
// 
// Any object may be used but it should extend either Sprite or Image and accept the same constructor arguments:
// when a new object is created it is passed the following parameters to its constructor: `(game, x, y, key, frame)`.
func (self *ParticlesArcadeEmitter) ClassType() interface{}{
    return self.Object.Get("classType")
}

// SetClassTypeA The type of objects that will be created when using {@link Phaser.Group#create create} or {@link Phaser.Group#createMultiple createMultiple}.
// 
// Any object may be used but it should extend either Sprite or Image and accept the same constructor arguments:
// when a new object is created it is passed the following parameters to its constructor: `(game, x, y, key, frame)`.
func (self *ParticlesArcadeEmitter) SetClassTypeA(member interface{}) {
    self.Object.Set("classType", member)
}

// Cursor The current display object that the group cursor is pointing to, if any. (Can be set manually.)
// 
// The cursor is a way to iterate through the children in a Group using {@link Phaser.Group#next next} and {@link Phaser.Group#previous previous}.
func (self *ParticlesArcadeEmitter) Cursor() *DisplayObject{
    return &DisplayObject{self.Object.Get("cursor")}
}

// SetCursorA The current display object that the group cursor is pointing to, if any. (Can be set manually.)
// 
// The cursor is a way to iterate through the children in a Group using {@link Phaser.Group#next next} and {@link Phaser.Group#previous previous}.
func (self *ParticlesArcadeEmitter) SetCursorA(member *DisplayObject) {
    self.Object.Set("cursor", member)
}

// InputEnableChildren A Group with `inputEnableChildren` set to `true` will automatically call `inputEnabled = true` 
// on any children _added_ to, or _created by_, this Group.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
func (self *ParticlesArcadeEmitter) InputEnableChildren() bool{
    return self.Object.Get("inputEnableChildren").Bool()
}

// SetInputEnableChildrenA A Group with `inputEnableChildren` set to `true` will automatically call `inputEnabled = true` 
// on any children _added_ to, or _created by_, this Group.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
func (self *ParticlesArcadeEmitter) SetInputEnableChildrenA(member bool) {
    self.Object.Set("inputEnableChildren", member)
}

// OnChildInputDown This Signal is dispatched whenever a child of this Group emits an onInputDown signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) OnChildInputDown() *Signal{
    return &Signal{self.Object.Get("onChildInputDown")}
}

// SetOnChildInputDownA This Signal is dispatched whenever a child of this Group emits an onInputDown signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) SetOnChildInputDownA(member *Signal) {
    self.Object.Set("onChildInputDown", member)
}

// OnChildInputUp This Signal is dispatched whenever a child of this Group emits an onInputUp signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 3 arguments: A reference to the Sprite that triggered the signal, 
// a reference to the Pointer that caused it, and a boolean value `isOver` that tells you if the Pointer
// is still over the Sprite or not.
func (self *ParticlesArcadeEmitter) OnChildInputUp() *Signal{
    return &Signal{self.Object.Get("onChildInputUp")}
}

// SetOnChildInputUpA This Signal is dispatched whenever a child of this Group emits an onInputUp signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 3 arguments: A reference to the Sprite that triggered the signal, 
// a reference to the Pointer that caused it, and a boolean value `isOver` that tells you if the Pointer
// is still over the Sprite or not.
func (self *ParticlesArcadeEmitter) SetOnChildInputUpA(member *Signal) {
    self.Object.Set("onChildInputUp", member)
}

// OnChildInputOver This Signal is dispatched whenever a child of this Group emits an onInputOver signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) OnChildInputOver() *Signal{
    return &Signal{self.Object.Get("onChildInputOver")}
}

// SetOnChildInputOverA This Signal is dispatched whenever a child of this Group emits an onInputOver signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) SetOnChildInputOverA(member *Signal) {
    self.Object.Set("onChildInputOver", member)
}

// OnChildInputOut This Signal is dispatched whenever a child of this Group emits an onInputOut signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) OnChildInputOut() *Signal{
    return &Signal{self.Object.Get("onChildInputOut")}
}

// SetOnChildInputOutA This Signal is dispatched whenever a child of this Group emits an onInputOut signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *ParticlesArcadeEmitter) SetOnChildInputOutA(member *Signal) {
    self.Object.Set("onChildInputOut", member)
}

// EnableBody If true all Sprites created by, or added to this group, will have a physics body enabled on them.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
// 
// The default body type is controlled with {@link Phaser.Group#physicsBodyType physicsBodyType}.
func (self *ParticlesArcadeEmitter) EnableBody() bool{
    return self.Object.Get("enableBody").Bool()
}

// SetEnableBodyA If true all Sprites created by, or added to this group, will have a physics body enabled on them.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
// 
// The default body type is controlled with {@link Phaser.Group#physicsBodyType physicsBodyType}.
func (self *ParticlesArcadeEmitter) SetEnableBodyA(member bool) {
    self.Object.Set("enableBody", member)
}

// EnableBodyDebug If true when a physics body is created (via {@link Phaser.Group#enableBody enableBody}) it will create a physics debug object as well.
// 
// This only works for P2 bodies.
func (self *ParticlesArcadeEmitter) EnableBodyDebug() bool{
    return self.Object.Get("enableBodyDebug").Bool()
}

// SetEnableBodyDebugA If true when a physics body is created (via {@link Phaser.Group#enableBody enableBody}) it will create a physics debug object as well.
// 
// This only works for P2 bodies.
func (self *ParticlesArcadeEmitter) SetEnableBodyDebugA(member bool) {
    self.Object.Set("enableBodyDebug", member)
}

// PhysicsBodyType If {@link Phaser.Group#enableBody enableBody} is true this is the type of physics body that is created on new Sprites.
// 
// The valid values are {@link Phaser.Physics.ARCADE}, {@link Phaser.Physics.P2JS}, {@link Phaser.Physics.NINJA}, etc.
func (self *ParticlesArcadeEmitter) PhysicsBodyType() int{
    return self.Object.Get("physicsBodyType").Int()
}

// SetPhysicsBodyTypeA If {@link Phaser.Group#enableBody enableBody} is true this is the type of physics body that is created on new Sprites.
// 
// The valid values are {@link Phaser.Physics.ARCADE}, {@link Phaser.Physics.P2JS}, {@link Phaser.Physics.NINJA}, etc.
func (self *ParticlesArcadeEmitter) SetPhysicsBodyTypeA(member int) {
    self.Object.Set("physicsBodyType", member)
}

// PhysicsSortDirection If this Group contains Arcade Physics Sprites you can set a custom sort direction via this property.
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
func (self *ParticlesArcadeEmitter) PhysicsSortDirection() int{
    return self.Object.Get("physicsSortDirection").Int()
}

// SetPhysicsSortDirectionA If this Group contains Arcade Physics Sprites you can set a custom sort direction via this property.
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
func (self *ParticlesArcadeEmitter) SetPhysicsSortDirectionA(member int) {
    self.Object.Set("physicsSortDirection", member)
}

// OnDestroy This signal is dispatched when the group is destroyed.
func (self *ParticlesArcadeEmitter) OnDestroy() *Signal{
    return &Signal{self.Object.Get("onDestroy")}
}

// SetOnDestroyA This signal is dispatched when the group is destroyed.
func (self *ParticlesArcadeEmitter) SetOnDestroyA(member *Signal) {
    self.Object.Set("onDestroy", member)
}

// CursorIndex The current index of the Group cursor. Advance it with Group.next.
func (self *ParticlesArcadeEmitter) CursorIndex() int{
    return self.Object.Get("cursorIndex").Int()
}

// SetCursorIndexA The current index of the Group cursor. Advance it with Group.next.
func (self *ParticlesArcadeEmitter) SetCursorIndexA(member int) {
    self.Object.Set("cursorIndex", member)
}

// FixedToCamera A Group that is fixed to the camera uses its x/y coordinates as offsets from the top left of the camera. These are stored in Group.cameraOffset.
// 
// Note that the cameraOffset values are in addition to any parent in the display list.
// So if this Group was in a Group that has x: 200, then this will be added to the cameraOffset.x
func (self *ParticlesArcadeEmitter) FixedToCamera() bool{
    return self.Object.Get("fixedToCamera").Bool()
}

// SetFixedToCameraA A Group that is fixed to the camera uses its x/y coordinates as offsets from the top left of the camera. These are stored in Group.cameraOffset.
// 
// Note that the cameraOffset values are in addition to any parent in the display list.
// So if this Group was in a Group that has x: 200, then this will be added to the cameraOffset.x
func (self *ParticlesArcadeEmitter) SetFixedToCameraA(member bool) {
    self.Object.Set("fixedToCamera", member)
}

// CameraOffset If this object is {@link Phaser.Group#fixedToCamera fixedToCamera} then this stores the x/y position offset relative to the top-left of the camera view.
// If the parent of this Group is also `fixedToCamera` then the offset here is in addition to that and should typically be disabled.
func (self *ParticlesArcadeEmitter) CameraOffset() *Point{
    return &Point{self.Object.Get("cameraOffset")}
}

// SetCameraOffsetA If this object is {@link Phaser.Group#fixedToCamera fixedToCamera} then this stores the x/y position offset relative to the top-left of the camera view.
// If the parent of this Group is also `fixedToCamera` then the offset here is in addition to that and should typically be disabled.
func (self *ParticlesArcadeEmitter) SetCameraOffsetA(member *Point) {
    self.Object.Set("cameraOffset", member)
}

// Hash The hash array is an array belonging to this Group into which you can add any of its children via Group.addToHash and Group.removeFromHash.
// 
// Only children of this Group can be added to and removed from the hash.
// 
// This hash is used automatically by Phaser Arcade Physics in order to perform non z-index based destructive sorting.
// However if you don't use Arcade Physics, or this isn't a physics enabled Group, then you can use the hash to perform your own
// sorting and filtering of Group children without touching their z-index (and therefore display draw order)
func (self *ParticlesArcadeEmitter) Hash() []interface{}{
	array00 := self.Object.Get("hash")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetHashA The hash array is an array belonging to this Group into which you can add any of its children via Group.addToHash and Group.removeFromHash.
// 
// Only children of this Group can be added to and removed from the hash.
// 
// This hash is used automatically by Phaser Arcade Physics in order to perform non z-index based destructive sorting.
// However if you don't use Arcade Physics, or this isn't a physics enabled Group, then you can use the hash to perform your own
// sorting and filtering of Group children without touching their z-index (and therefore display draw order)
func (self *ParticlesArcadeEmitter) SetHashA(member []interface{}) {
    self.Object.Set("hash", member)
}

// Total Total number of existing children in the group.
func (self *ParticlesArcadeEmitter) Total() int{
    return self.Object.Get("total").Int()
}

// SetTotalA Total number of existing children in the group.
func (self *ParticlesArcadeEmitter) SetTotalA(member int) {
    self.Object.Set("total", member)
}

// Length Total number of children in this group, regardless of exists/alive status.
func (self *ParticlesArcadeEmitter) Length() int{
    return self.Object.Get("length").Int()
}

// SetLengthA Total number of children in this group, regardless of exists/alive status.
func (self *ParticlesArcadeEmitter) SetLengthA(member int) {
    self.Object.Set("length", member)
}

// Angle The angle of rotation of the group container, in degrees.
// 
// This adjusts the group itself by modifying its local rotation transform.
// 
// This has no impact on the rotation/angle properties of the children, but it will update their worldTransform
// and on-screen orientation and position.
func (self *ParticlesArcadeEmitter) Angle() int{
    return self.Object.Get("angle").Int()
}

// SetAngleA The angle of rotation of the group container, in degrees.
// 
// This adjusts the group itself by modifying its local rotation transform.
// 
// This has no impact on the rotation/angle properties of the children, but it will update their worldTransform
// and on-screen orientation and position.
func (self *ParticlesArcadeEmitter) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// CenterX The center x coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *ParticlesArcadeEmitter) CenterX() int{
    return self.Object.Get("centerX").Int()
}

// SetCenterXA The center x coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *ParticlesArcadeEmitter) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// CenterY The center y coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *ParticlesArcadeEmitter) CenterY() int{
    return self.Object.Get("centerY").Int()
}

// SetCenterYA The center y coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *ParticlesArcadeEmitter) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// Rotation The angle of rotation of the group container, in radians.
// 
// This will adjust the group container itself by modifying its rotation.
// This will have no impact on the rotation value of its children, but it will update their worldTransform and on-screen position.
func (self *ParticlesArcadeEmitter) Rotation() int{
    return self.Object.Get("rotation").Int()
}

// SetRotationA The angle of rotation of the group container, in radians.
// 
// This will adjust the group container itself by modifying its rotation.
// This will have no impact on the rotation value of its children, but it will update their worldTransform and on-screen position.
func (self *ParticlesArcadeEmitter) SetRotationA(member int) {
    self.Object.Set("rotation", member)
}

// Visible The visible state of the group. Non-visible Groups and all of their children are not rendered.
func (self *ParticlesArcadeEmitter) Visible() bool{
    return self.Object.Get("visible").Bool()
}

// SetVisibleA The visible state of the group. Non-visible Groups and all of their children are not rendered.
func (self *ParticlesArcadeEmitter) SetVisibleA(member bool) {
    self.Object.Set("visible", member)
}

// Alpha The alpha value of the group container.
func (self *ParticlesArcadeEmitter) Alpha() int{
    return self.Object.Get("alpha").Int()
}

// SetAlphaA The alpha value of the group container.
func (self *ParticlesArcadeEmitter) SetAlphaA(member int) {
    self.Object.Set("alpha", member)
}

// Children [read-only] The array of children of this container.
func (self *ParticlesArcadeEmitter) Children() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// SetChildrenA [read-only] The array of children of this container.
func (self *ParticlesArcadeEmitter) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// IgnoreChildInput If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *ParticlesArcadeEmitter) IgnoreChildInput() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// SetIgnoreChildInputA If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *ParticlesArcadeEmitter) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}


// Update Called automatically by the game loop, decides when to launch particles and when to "die".
func (self *ParticlesArcadeEmitter) Update() {
    self.Object.Call("update")
}

// UpdateI Called automatically by the game loop, decides when to launch particles and when to "die".
func (self *ParticlesArcadeEmitter) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// MakeParticles This function generates a new set of particles for use by this emitter.
// The particles are stored internally waiting to be emitted via Emitter.start.
func (self *ParticlesArcadeEmitter) MakeParticles(keys interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("makeParticles", keys)}
}

// MakeParticles1O This function generates a new set of particles for use by this emitter.
// The particles are stored internally waiting to be emitted via Emitter.start.
func (self *ParticlesArcadeEmitter) MakeParticles1O(keys interface{}, frames interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("makeParticles", keys, frames)}
}

// MakeParticles2O This function generates a new set of particles for use by this emitter.
// The particles are stored internally waiting to be emitted via Emitter.start.
func (self *ParticlesArcadeEmitter) MakeParticles2O(keys interface{}, frames interface{}, quantity int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("makeParticles", keys, frames, quantity)}
}

// MakeParticles3O This function generates a new set of particles for use by this emitter.
// The particles are stored internally waiting to be emitted via Emitter.start.
func (self *ParticlesArcadeEmitter) MakeParticles3O(keys interface{}, frames interface{}, quantity int, collide bool) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("makeParticles", keys, frames, quantity, collide)}
}

// MakeParticles4O This function generates a new set of particles for use by this emitter.
// The particles are stored internally waiting to be emitted via Emitter.start.
func (self *ParticlesArcadeEmitter) MakeParticles4O(keys interface{}, frames interface{}, quantity int, collide bool, collideWorldBounds bool) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("makeParticles", keys, frames, quantity, collide, collideWorldBounds)}
}

// MakeParticlesI This function generates a new set of particles for use by this emitter.
// The particles are stored internally waiting to be emitted via Emitter.start.
func (self *ParticlesArcadeEmitter) MakeParticlesI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("makeParticles", args)}
}

// Kill Call this function to turn off all the particles and the emitter.
func (self *ParticlesArcadeEmitter) Kill() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("kill")}
}

// KillI Call this function to turn off all the particles and the emitter.
func (self *ParticlesArcadeEmitter) KillI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("kill", args)}
}

// Revive Handy for bringing game objects "back to life". Just sets alive and exists back to true.
func (self *ParticlesArcadeEmitter) Revive() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("revive")}
}

// ReviveI Handy for bringing game objects "back to life". Just sets alive and exists back to true.
func (self *ParticlesArcadeEmitter) ReviveI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("revive", args)}
}

// Explode Call this function to emit the given quantity of particles at all once (an explosion)
func (self *ParticlesArcadeEmitter) Explode() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("explode")}
}

// Explode1O Call this function to emit the given quantity of particles at all once (an explosion)
func (self *ParticlesArcadeEmitter) Explode1O(lifespan int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("explode", lifespan)}
}

// Explode2O Call this function to emit the given quantity of particles at all once (an explosion)
func (self *ParticlesArcadeEmitter) Explode2O(lifespan int, quantity int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("explode", lifespan, quantity)}
}

// ExplodeI Call this function to emit the given quantity of particles at all once (an explosion)
func (self *ParticlesArcadeEmitter) ExplodeI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("explode", args)}
}

// Flow Call this function to start emitting a flow of particles at the given frequency.
// It will carry on going until the total given is reached.
// Each time the flow is run the quantity number of particles will be emitted together.
// If you set the total to be 20 and quantity to be 5 then flow will emit 4 times in total (4 x 5 = 20 total)
// If you set the total to be -1 then no quantity cap is used and it will keep emitting.
func (self *ParticlesArcadeEmitter) Flow() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("flow")}
}

// Flow1O Call this function to start emitting a flow of particles at the given frequency.
// It will carry on going until the total given is reached.
// Each time the flow is run the quantity number of particles will be emitted together.
// If you set the total to be 20 and quantity to be 5 then flow will emit 4 times in total (4 x 5 = 20 total)
// If you set the total to be -1 then no quantity cap is used and it will keep emitting.
func (self *ParticlesArcadeEmitter) Flow1O(lifespan int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("flow", lifespan)}
}

// Flow2O Call this function to start emitting a flow of particles at the given frequency.
// It will carry on going until the total given is reached.
// Each time the flow is run the quantity number of particles will be emitted together.
// If you set the total to be 20 and quantity to be 5 then flow will emit 4 times in total (4 x 5 = 20 total)
// If you set the total to be -1 then no quantity cap is used and it will keep emitting.
func (self *ParticlesArcadeEmitter) Flow2O(lifespan int, frequency int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("flow", lifespan, frequency)}
}

// Flow3O Call this function to start emitting a flow of particles at the given frequency.
// It will carry on going until the total given is reached.
// Each time the flow is run the quantity number of particles will be emitted together.
// If you set the total to be 20 and quantity to be 5 then flow will emit 4 times in total (4 x 5 = 20 total)
// If you set the total to be -1 then no quantity cap is used and it will keep emitting.
func (self *ParticlesArcadeEmitter) Flow3O(lifespan int, frequency int, quantity int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("flow", lifespan, frequency, quantity)}
}

// Flow4O Call this function to start emitting a flow of particles at the given frequency.
// It will carry on going until the total given is reached.
// Each time the flow is run the quantity number of particles will be emitted together.
// If you set the total to be 20 and quantity to be 5 then flow will emit 4 times in total (4 x 5 = 20 total)
// If you set the total to be -1 then no quantity cap is used and it will keep emitting.
func (self *ParticlesArcadeEmitter) Flow4O(lifespan int, frequency int, quantity int, total int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("flow", lifespan, frequency, quantity, total)}
}

// Flow5O Call this function to start emitting a flow of particles at the given frequency.
// It will carry on going until the total given is reached.
// Each time the flow is run the quantity number of particles will be emitted together.
// If you set the total to be 20 and quantity to be 5 then flow will emit 4 times in total (4 x 5 = 20 total)
// If you set the total to be -1 then no quantity cap is used and it will keep emitting.
func (self *ParticlesArcadeEmitter) Flow5O(lifespan int, frequency int, quantity int, total int, immediate bool) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("flow", lifespan, frequency, quantity, total, immediate)}
}

// FlowI Call this function to start emitting a flow of particles at the given frequency.
// It will carry on going until the total given is reached.
// Each time the flow is run the quantity number of particles will be emitted together.
// If you set the total to be 20 and quantity to be 5 then flow will emit 4 times in total (4 x 5 = 20 total)
// If you set the total to be -1 then no quantity cap is used and it will keep emitting.
func (self *ParticlesArcadeEmitter) FlowI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("flow", args)}
}

// Start Call this function to start emitting particles.
func (self *ParticlesArcadeEmitter) Start() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("start")}
}

// Start1O Call this function to start emitting particles.
func (self *ParticlesArcadeEmitter) Start1O(explode bool) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("start", explode)}
}

// Start2O Call this function to start emitting particles.
func (self *ParticlesArcadeEmitter) Start2O(explode bool, lifespan int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("start", explode, lifespan)}
}

// Start3O Call this function to start emitting particles.
func (self *ParticlesArcadeEmitter) Start3O(explode bool, lifespan int, frequency int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("start", explode, lifespan, frequency)}
}

// Start4O Call this function to start emitting particles.
func (self *ParticlesArcadeEmitter) Start4O(explode bool, lifespan int, frequency int, quantity int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("start", explode, lifespan, frequency, quantity)}
}

// Start5O Call this function to start emitting particles.
func (self *ParticlesArcadeEmitter) Start5O(explode bool, lifespan int, frequency int, quantity int, forceQuantity int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("start", explode, lifespan, frequency, quantity, forceQuantity)}
}

// StartI Call this function to start emitting particles.
func (self *ParticlesArcadeEmitter) StartI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("start", args)}
}

// EmitParticle This function is used internally to emit the next particle in the queue.
// 
// However it can also be called externally to emit a particle.
// 
// When called externally you can use the arguments to override any defaults the Emitter has set.
func (self *ParticlesArcadeEmitter) EmitParticle() bool{
    return self.Object.Call("emitParticle").Bool()
}

// EmitParticle1O This function is used internally to emit the next particle in the queue.
// 
// However it can also be called externally to emit a particle.
// 
// When called externally you can use the arguments to override any defaults the Emitter has set.
func (self *ParticlesArcadeEmitter) EmitParticle1O(x int) bool{
    return self.Object.Call("emitParticle", x).Bool()
}

// EmitParticle2O This function is used internally to emit the next particle in the queue.
// 
// However it can also be called externally to emit a particle.
// 
// When called externally you can use the arguments to override any defaults the Emitter has set.
func (self *ParticlesArcadeEmitter) EmitParticle2O(x int, y int) bool{
    return self.Object.Call("emitParticle", x, y).Bool()
}

// EmitParticle3O This function is used internally to emit the next particle in the queue.
// 
// However it can also be called externally to emit a particle.
// 
// When called externally you can use the arguments to override any defaults the Emitter has set.
func (self *ParticlesArcadeEmitter) EmitParticle3O(x int, y int, key interface{}) bool{
    return self.Object.Call("emitParticle", x, y, key).Bool()
}

// EmitParticle4O This function is used internally to emit the next particle in the queue.
// 
// However it can also be called externally to emit a particle.
// 
// When called externally you can use the arguments to override any defaults the Emitter has set.
func (self *ParticlesArcadeEmitter) EmitParticle4O(x int, y int, key interface{}, frame interface{}) bool{
    return self.Object.Call("emitParticle", x, y, key, frame).Bool()
}

// EmitParticleI This function is used internally to emit the next particle in the queue.
// 
// However it can also be called externally to emit a particle.
// 
// When called externally you can use the arguments to override any defaults the Emitter has set.
func (self *ParticlesArcadeEmitter) EmitParticleI(args ...interface{}) bool{
    return self.Object.Call("emitParticle", args).Bool()
}

// Destroy Destroys this Emitter, all associated child Particles and then removes itself from the Particle Manager.
func (self *ParticlesArcadeEmitter) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this Emitter, all associated child Particles and then removes itself from the Particle Manager.
func (self *ParticlesArcadeEmitter) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// SetSize A more compact way of setting the width and height of the emitter.
func (self *ParticlesArcadeEmitter) SetSize(width int, height int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setSize", width, height)}
}

// SetSizeI A more compact way of setting the width and height of the emitter.
func (self *ParticlesArcadeEmitter) SetSizeI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setSize", args)}
}

// SetXSpeed A more compact way of setting the X velocity range of the emitter.
func (self *ParticlesArcadeEmitter) SetXSpeed() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setXSpeed")}
}

// SetXSpeed1O A more compact way of setting the X velocity range of the emitter.
func (self *ParticlesArcadeEmitter) SetXSpeed1O(min int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setXSpeed", min)}
}

// SetXSpeed2O A more compact way of setting the X velocity range of the emitter.
func (self *ParticlesArcadeEmitter) SetXSpeed2O(min int, max int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setXSpeed", min, max)}
}

// SetXSpeedI A more compact way of setting the X velocity range of the emitter.
func (self *ParticlesArcadeEmitter) SetXSpeedI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setXSpeed", args)}
}

// SetYSpeed A more compact way of setting the Y velocity range of the emitter.
func (self *ParticlesArcadeEmitter) SetYSpeed() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setYSpeed")}
}

// SetYSpeed1O A more compact way of setting the Y velocity range of the emitter.
func (self *ParticlesArcadeEmitter) SetYSpeed1O(min int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setYSpeed", min)}
}

// SetYSpeed2O A more compact way of setting the Y velocity range of the emitter.
func (self *ParticlesArcadeEmitter) SetYSpeed2O(min int, max int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setYSpeed", min, max)}
}

// SetYSpeedI A more compact way of setting the Y velocity range of the emitter.
func (self *ParticlesArcadeEmitter) SetYSpeedI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setYSpeed", args)}
}

// SetRotation A more compact way of setting the angular velocity constraints of the particles.
func (self *ParticlesArcadeEmitter) SetRotation() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setRotation")}
}

// SetRotation1O A more compact way of setting the angular velocity constraints of the particles.
func (self *ParticlesArcadeEmitter) SetRotation1O(min int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setRotation", min)}
}

// SetRotation2O A more compact way of setting the angular velocity constraints of the particles.
func (self *ParticlesArcadeEmitter) SetRotation2O(min int, max int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setRotation", min, max)}
}

// SetRotationI A more compact way of setting the angular velocity constraints of the particles.
func (self *ParticlesArcadeEmitter) SetRotationI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setRotation", args)}
}

// SetAlpha A more compact way of setting the alpha constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed at which the Particle change in alpha from min to max.
// If rate is zero, which is the default, the particle won't change alpha - instead it will pick a random alpha between min and max on emit.
func (self *ParticlesArcadeEmitter) SetAlpha() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setAlpha")}
}

// SetAlpha1O A more compact way of setting the alpha constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed at which the Particle change in alpha from min to max.
// If rate is zero, which is the default, the particle won't change alpha - instead it will pick a random alpha between min and max on emit.
func (self *ParticlesArcadeEmitter) SetAlpha1O(min int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setAlpha", min)}
}

// SetAlpha2O A more compact way of setting the alpha constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed at which the Particle change in alpha from min to max.
// If rate is zero, which is the default, the particle won't change alpha - instead it will pick a random alpha between min and max on emit.
func (self *ParticlesArcadeEmitter) SetAlpha2O(min int, max int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setAlpha", min, max)}
}

// SetAlpha3O A more compact way of setting the alpha constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed at which the Particle change in alpha from min to max.
// If rate is zero, which is the default, the particle won't change alpha - instead it will pick a random alpha between min and max on emit.
func (self *ParticlesArcadeEmitter) SetAlpha3O(min int, max int, rate int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setAlpha", min, max, rate)}
}

// SetAlpha4O A more compact way of setting the alpha constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed at which the Particle change in alpha from min to max.
// If rate is zero, which is the default, the particle won't change alpha - instead it will pick a random alpha between min and max on emit.
func (self *ParticlesArcadeEmitter) SetAlpha4O(min int, max int, rate int, ease interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setAlpha", min, max, rate, ease)}
}

// SetAlpha5O A more compact way of setting the alpha constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed at which the Particle change in alpha from min to max.
// If rate is zero, which is the default, the particle won't change alpha - instead it will pick a random alpha between min and max on emit.
func (self *ParticlesArcadeEmitter) SetAlpha5O(min int, max int, rate int, ease interface{}, yoyo bool) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setAlpha", min, max, rate, ease, yoyo)}
}

// SetAlphaI A more compact way of setting the alpha constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed at which the Particle change in alpha from min to max.
// If rate is zero, which is the default, the particle won't change alpha - instead it will pick a random alpha between min and max on emit.
func (self *ParticlesArcadeEmitter) SetAlphaI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setAlpha", args)}
}

// SetScale A more compact way of setting the scale constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed and ease which the Particle uses to change in scale from min to max across both axis.
// If rate is zero, which is the default, the particle won't change scale during update, instead it will pick a random scale between min and max on emit.
func (self *ParticlesArcadeEmitter) SetScale() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setScale")}
}

// SetScale1O A more compact way of setting the scale constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed and ease which the Particle uses to change in scale from min to max across both axis.
// If rate is zero, which is the default, the particle won't change scale during update, instead it will pick a random scale between min and max on emit.
func (self *ParticlesArcadeEmitter) SetScale1O(minX int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setScale", minX)}
}

// SetScale2O A more compact way of setting the scale constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed and ease which the Particle uses to change in scale from min to max across both axis.
// If rate is zero, which is the default, the particle won't change scale during update, instead it will pick a random scale between min and max on emit.
func (self *ParticlesArcadeEmitter) SetScale2O(minX int, maxX int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setScale", minX, maxX)}
}

// SetScale3O A more compact way of setting the scale constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed and ease which the Particle uses to change in scale from min to max across both axis.
// If rate is zero, which is the default, the particle won't change scale during update, instead it will pick a random scale between min and max on emit.
func (self *ParticlesArcadeEmitter) SetScale3O(minX int, maxX int, minY int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setScale", minX, maxX, minY)}
}

// SetScale4O A more compact way of setting the scale constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed and ease which the Particle uses to change in scale from min to max across both axis.
// If rate is zero, which is the default, the particle won't change scale during update, instead it will pick a random scale between min and max on emit.
func (self *ParticlesArcadeEmitter) SetScale4O(minX int, maxX int, minY int, maxY int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setScale", minX, maxX, minY, maxY)}
}

// SetScale5O A more compact way of setting the scale constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed and ease which the Particle uses to change in scale from min to max across both axis.
// If rate is zero, which is the default, the particle won't change scale during update, instead it will pick a random scale between min and max on emit.
func (self *ParticlesArcadeEmitter) SetScale5O(minX int, maxX int, minY int, maxY int, rate int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setScale", minX, maxX, minY, maxY, rate)}
}

// SetScale6O A more compact way of setting the scale constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed and ease which the Particle uses to change in scale from min to max across both axis.
// If rate is zero, which is the default, the particle won't change scale during update, instead it will pick a random scale between min and max on emit.
func (self *ParticlesArcadeEmitter) SetScale6O(minX int, maxX int, minY int, maxY int, rate int, ease interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setScale", minX, maxX, minY, maxY, rate, ease)}
}

// SetScale7O A more compact way of setting the scale constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed and ease which the Particle uses to change in scale from min to max across both axis.
// If rate is zero, which is the default, the particle won't change scale during update, instead it will pick a random scale between min and max on emit.
func (self *ParticlesArcadeEmitter) SetScale7O(minX int, maxX int, minY int, maxY int, rate int, ease interface{}, yoyo bool) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setScale", minX, maxX, minY, maxY, rate, ease, yoyo)}
}

// SetScaleI A more compact way of setting the scale constraints of the particles.
// The rate parameter, if set to a value above zero, lets you set the speed and ease which the Particle uses to change in scale from min to max across both axis.
// If rate is zero, which is the default, the particle won't change scale during update, instead it will pick a random scale between min and max on emit.
func (self *ParticlesArcadeEmitter) SetScaleI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("setScale", args)}
}

// At Change the emitters center to match the center of any object with a `center` property, such as a Sprite.
// If the object doesn't have a center property it will be set to object.x + object.width / 2
func (self *ParticlesArcadeEmitter) At(object interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("at", object)}
}

// AtI Change the emitters center to match the center of any object with a `center` property, such as a Sprite.
// If the object doesn't have a center property it will be set to object.x + object.width / 2
func (self *ParticlesArcadeEmitter) AtI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("at", args)}
}

// Add Adds an existing object as the top child in this group.
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
func (self *ParticlesArcadeEmitter) Add(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("add", child)}
}

// Add1O Adds an existing object as the top child in this group.
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
func (self *ParticlesArcadeEmitter) Add1O(child *DisplayObject, silent bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("add", child, silent)}
}

// Add2O Adds an existing object as the top child in this group.
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
func (self *ParticlesArcadeEmitter) Add2O(child *DisplayObject, silent bool, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("add", child, silent, index)}
}

// AddI Adds an existing object as the top child in this group.
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
    return &DisplayObject{self.Object.Call("add", args)}
}

// AddAt Adds an existing object to this group.
// 
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) AddAt(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addAt", child)}
}

// AddAt1O Adds an existing object to this group.
// 
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) AddAt1O(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addAt", child, index)}
}

// AddAt2O Adds an existing object to this group.
// 
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) AddAt2O(child *DisplayObject, index int, silent bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("addAt", child, index, silent)}
}

// AddAtI Adds an existing object to this group.
// 
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) AddAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addAt", args)}
}

// AddToHash Adds a child of this Group into the hash array.
// This call will return false if the child is not a child of this Group, or is already in the hash.
func (self *ParticlesArcadeEmitter) AddToHash(child *DisplayObject) bool{
    return self.Object.Call("addToHash", child).Bool()
}

// AddToHashI Adds a child of this Group into the hash array.
// This call will return false if the child is not a child of this Group, or is already in the hash.
func (self *ParticlesArcadeEmitter) AddToHashI(args ...interface{}) bool{
    return self.Object.Call("addToHash", args).Bool()
}

// RemoveFromHash Removes a child of this Group from the hash array.
// This call will return false if the child is not in the hash.
func (self *ParticlesArcadeEmitter) RemoveFromHash(child *DisplayObject) bool{
    return self.Object.Call("removeFromHash", child).Bool()
}

// RemoveFromHashI Removes a child of this Group from the hash array.
// This call will return false if the child is not in the hash.
func (self *ParticlesArcadeEmitter) RemoveFromHashI(args ...interface{}) bool{
    return self.Object.Call("removeFromHash", args).Bool()
}

// AddMultiple Adds an array of existing Display Objects to this Group.
// 
// The Display Objects are automatically added to the top of this Group, and will render on-top of everything already in this Group.
// 
// As well as an array you can also pass another Group as the first argument. In this case all of the children from that
// Group will be removed from it and added into this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) AddMultiple(children interface{}) interface{}{
    return self.Object.Call("addMultiple", children)
}

// AddMultiple1O Adds an array of existing Display Objects to this Group.
// 
// The Display Objects are automatically added to the top of this Group, and will render on-top of everything already in this Group.
// 
// As well as an array you can also pass another Group as the first argument. In this case all of the children from that
// Group will be removed from it and added into this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) AddMultiple1O(children interface{}, silent bool) interface{}{
    return self.Object.Call("addMultiple", children, silent)
}

// AddMultipleI Adds an array of existing Display Objects to this Group.
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
    return self.Object.Call("addMultiple", args)
}

// GetAt Returns the child found at the given index within this group.
func (self *ParticlesArcadeEmitter) GetAt(index int) interface{}{
    return self.Object.Call("getAt", index)
}

// GetAtI Returns the child found at the given index within this group.
func (self *ParticlesArcadeEmitter) GetAtI(args ...interface{}) interface{}{
    return self.Object.Call("getAt", args)
}

// Create Creates a new Phaser.Sprite object and adds it to the top of this group.
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
func (self *ParticlesArcadeEmitter) Create(x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", x, y)}
}

// Create1O Creates a new Phaser.Sprite object and adds it to the top of this group.
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
func (self *ParticlesArcadeEmitter) Create1O(x int, y int, key interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", x, y, key)}
}

// Create2O Creates a new Phaser.Sprite object and adds it to the top of this group.
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
func (self *ParticlesArcadeEmitter) Create2O(x int, y int, key interface{}, frame interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", x, y, key, frame)}
}

// Create3O Creates a new Phaser.Sprite object and adds it to the top of this group.
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
func (self *ParticlesArcadeEmitter) Create3O(x int, y int, key interface{}, frame interface{}, exists bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", x, y, key, frame, exists)}
}

// Create4O Creates a new Phaser.Sprite object and adds it to the top of this group.
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
func (self *ParticlesArcadeEmitter) Create4O(x int, y int, key interface{}, frame interface{}, exists bool, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", x, y, key, frame, exists, index)}
}

// CreateI Creates a new Phaser.Sprite object and adds it to the top of this group.
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
    return &DisplayObject{self.Object.Call("create", args)}
}

// CreateMultiple Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
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
func (self *ParticlesArcadeEmitter) CreateMultiple(quantity int, key interface{}) []interface{}{
	array00 := self.Object.Call("createMultiple", quantity, key)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// CreateMultiple1O Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
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
func (self *ParticlesArcadeEmitter) CreateMultiple1O(quantity int, key interface{}, frame interface{}) []interface{}{
	array00 := self.Object.Call("createMultiple", quantity, key, frame)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// CreateMultiple2O Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
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
func (self *ParticlesArcadeEmitter) CreateMultiple2O(quantity int, key interface{}, frame interface{}, exists bool) []interface{}{
	array00 := self.Object.Call("createMultiple", quantity, key, frame, exists)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// CreateMultipleI Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
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
	array00 := self.Object.Call("createMultiple", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// UpdateZ Internal method that re-applies all of the children's Z values.
// 
// This must be called whenever children ordering is altered so that their `z` indices are correctly updated.
func (self *ParticlesArcadeEmitter) UpdateZ() {
    self.Object.Call("updateZ")
}

// UpdateZI Internal method that re-applies all of the children's Z values.
// 
// This must be called whenever children ordering is altered so that their `z` indices are correctly updated.
func (self *ParticlesArcadeEmitter) UpdateZI(args ...interface{}) {
    self.Object.Call("updateZ", args)
}

// Align This method iterates through all children in the Group (regardless if they are visible or exist)
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
func (self *ParticlesArcadeEmitter) Align(rows int, columns int, cellWidth int, cellHeight int) {
    self.Object.Call("align", rows, columns, cellWidth, cellHeight)
}

// Align1O This method iterates through all children in the Group (regardless if they are visible or exist)
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
func (self *ParticlesArcadeEmitter) Align1O(rows int, columns int, cellWidth int, cellHeight int, position int) {
    self.Object.Call("align", rows, columns, cellWidth, cellHeight, position)
}

// Align2O This method iterates through all children in the Group (regardless if they are visible or exist)
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
func (self *ParticlesArcadeEmitter) Align2O(rows int, columns int, cellWidth int, cellHeight int, position int, offset int) {
    self.Object.Call("align", rows, columns, cellWidth, cellHeight, position, offset)
}

// AlignI This method iterates through all children in the Group (regardless if they are visible or exist)
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
    self.Object.Call("align", args)
}

// ResetCursor Sets the group cursor to the first child in the group.
// 
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *ParticlesArcadeEmitter) ResetCursor() interface{}{
    return self.Object.Call("resetCursor")
}

// ResetCursor1O Sets the group cursor to the first child in the group.
// 
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *ParticlesArcadeEmitter) ResetCursor1O(index int) interface{}{
    return self.Object.Call("resetCursor", index)
}

// ResetCursorI Sets the group cursor to the first child in the group.
// 
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *ParticlesArcadeEmitter) ResetCursorI(args ...interface{}) interface{}{
    return self.Object.Call("resetCursor", args)
}

// Next Advances the group cursor to the next (higher) object in the group.
// 
// If the cursor is at the end of the group (top child) it is moved the start of the group (bottom child).
func (self *ParticlesArcadeEmitter) Next() interface{}{
    return self.Object.Call("next")
}

// NextI Advances the group cursor to the next (higher) object in the group.
// 
// If the cursor is at the end of the group (top child) it is moved the start of the group (bottom child).
func (self *ParticlesArcadeEmitter) NextI(args ...interface{}) interface{}{
    return self.Object.Call("next", args)
}

// Previous Moves the group cursor to the previous (lower) child in the group.
// 
// If the cursor is at the start of the group (bottom child) it is moved to the end (top child).
func (self *ParticlesArcadeEmitter) Previous() interface{}{
    return self.Object.Call("previous")
}

// PreviousI Moves the group cursor to the previous (lower) child in the group.
// 
// If the cursor is at the start of the group (bottom child) it is moved to the end (top child).
func (self *ParticlesArcadeEmitter) PreviousI(args ...interface{}) interface{}{
    return self.Object.Call("previous", args)
}

// Swap Swaps the position of two children in this group.
// 
// Both children must be in this group, a child cannot be swapped with itself, and unparented children cannot be swapped.
func (self *ParticlesArcadeEmitter) Swap(child1 interface{}, child2 interface{}) {
    self.Object.Call("swap", child1, child2)
}

// SwapI Swaps the position of two children in this group.
// 
// Both children must be in this group, a child cannot be swapped with itself, and unparented children cannot be swapped.
func (self *ParticlesArcadeEmitter) SwapI(args ...interface{}) {
    self.Object.Call("swap", args)
}

// BringToTop Brings the given child to the top of this group so it renders above all other children.
func (self *ParticlesArcadeEmitter) BringToTop(child interface{}) interface{}{
    return self.Object.Call("bringToTop", child)
}

// BringToTopI Brings the given child to the top of this group so it renders above all other children.
func (self *ParticlesArcadeEmitter) BringToTopI(args ...interface{}) interface{}{
    return self.Object.Call("bringToTop", args)
}

// SendToBack Sends the given child to the bottom of this group so it renders below all other children.
func (self *ParticlesArcadeEmitter) SendToBack(child interface{}) interface{}{
    return self.Object.Call("sendToBack", child)
}

// SendToBackI Sends the given child to the bottom of this group so it renders below all other children.
func (self *ParticlesArcadeEmitter) SendToBackI(args ...interface{}) interface{}{
    return self.Object.Call("sendToBack", args)
}

// MoveUp Moves the given child up one place in this group unless it's already at the top.
func (self *ParticlesArcadeEmitter) MoveUp(child interface{}) interface{}{
    return self.Object.Call("moveUp", child)
}

// MoveUpI Moves the given child up one place in this group unless it's already at the top.
func (self *ParticlesArcadeEmitter) MoveUpI(args ...interface{}) interface{}{
    return self.Object.Call("moveUp", args)
}

// MoveDown Moves the given child down one place in this group unless it's already at the bottom.
func (self *ParticlesArcadeEmitter) MoveDown(child interface{}) interface{}{
    return self.Object.Call("moveDown", child)
}

// MoveDownI Moves the given child down one place in this group unless it's already at the bottom.
func (self *ParticlesArcadeEmitter) MoveDownI(args ...interface{}) interface{}{
    return self.Object.Call("moveDown", args)
}

// Xy Positions the child found at the given index within this group to the given x and y coordinates.
func (self *ParticlesArcadeEmitter) Xy(index int, x int, y int) {
    self.Object.Call("xy", index, x, y)
}

// XyI Positions the child found at the given index within this group to the given x and y coordinates.
func (self *ParticlesArcadeEmitter) XyI(args ...interface{}) {
    self.Object.Call("xy", args)
}

// Reverse Reverses all children in this group.
// 
// This operation applies only to immediate children and does not propagate to subgroups.
func (self *ParticlesArcadeEmitter) Reverse() {
    self.Object.Call("reverse")
}

// ReverseI Reverses all children in this group.
// 
// This operation applies only to immediate children and does not propagate to subgroups.
func (self *ParticlesArcadeEmitter) ReverseI(args ...interface{}) {
    self.Object.Call("reverse", args)
}

// GetIndex Get the index position of the given child in this group, which should match the child's `z` property.
func (self *ParticlesArcadeEmitter) GetIndex(child interface{}) int{
    return self.Object.Call("getIndex", child).Int()
}

// GetIndexI Get the index position of the given child in this group, which should match the child's `z` property.
func (self *ParticlesArcadeEmitter) GetIndexI(args ...interface{}) int{
    return self.Object.Call("getIndex", args).Int()
}

// GetByName Searches the Group for the first instance of a child with the `name`
// property matching the given argument. Should more than one child have
// the same name only the first instance is returned.
func (self *ParticlesArcadeEmitter) GetByName(name string) interface{}{
    return self.Object.Call("getByName", name)
}

// GetByNameI Searches the Group for the first instance of a child with the `name`
// property matching the given argument. Should more than one child have
// the same name only the first instance is returned.
func (self *ParticlesArcadeEmitter) GetByNameI(args ...interface{}) interface{}{
    return self.Object.Call("getByName", args)
}

// Replace Replaces a child of this Group with the given newChild. The newChild cannot be a member of this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) Replace(oldChild interface{}, newChild interface{}) interface{}{
    return self.Object.Call("replace", oldChild, newChild)
}

// ReplaceI Replaces a child of this Group with the given newChild. The newChild cannot be a member of this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *ParticlesArcadeEmitter) ReplaceI(args ...interface{}) interface{}{
    return self.Object.Call("replace", args)
}

// HasProperty Checks if the child has the given property.
// 
// Will scan up to 4 levels deep only.
func (self *ParticlesArcadeEmitter) HasProperty(child interface{}, key []string) bool{
    return self.Object.Call("hasProperty", child, key).Bool()
}

// HasPropertyI Checks if the child has the given property.
// 
// Will scan up to 4 levels deep only.
func (self *ParticlesArcadeEmitter) HasPropertyI(args ...interface{}) bool{
    return self.Object.Call("hasProperty", args).Bool()
}

// SetProperty Sets a property to the given value on the child. The operation parameter controls how the value is set.
// 
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *ParticlesArcadeEmitter) SetProperty(child interface{}, key []interface{}, value interface{}) bool{
    return self.Object.Call("setProperty", child, key, value).Bool()
}

// SetProperty1O Sets a property to the given value on the child. The operation parameter controls how the value is set.
// 
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *ParticlesArcadeEmitter) SetProperty1O(child interface{}, key []interface{}, value interface{}, operation int) bool{
    return self.Object.Call("setProperty", child, key, value, operation).Bool()
}

// SetProperty2O Sets a property to the given value on the child. The operation parameter controls how the value is set.
// 
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *ParticlesArcadeEmitter) SetProperty2O(child interface{}, key []interface{}, value interface{}, operation int, force bool) bool{
    return self.Object.Call("setProperty", child, key, value, operation, force).Bool()
}

// SetPropertyI Sets a property to the given value on the child. The operation parameter controls how the value is set.
// 
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *ParticlesArcadeEmitter) SetPropertyI(args ...interface{}) bool{
    return self.Object.Call("setProperty", args).Bool()
}

// CheckProperty Checks a property for the given value on the child.
func (self *ParticlesArcadeEmitter) CheckProperty(child interface{}, key []interface{}, value interface{}) bool{
    return self.Object.Call("checkProperty", child, key, value).Bool()
}

// CheckProperty1O Checks a property for the given value on the child.
func (self *ParticlesArcadeEmitter) CheckProperty1O(child interface{}, key []interface{}, value interface{}, force bool) bool{
    return self.Object.Call("checkProperty", child, key, value, force).Bool()
}

// CheckPropertyI Checks a property for the given value on the child.
func (self *ParticlesArcadeEmitter) CheckPropertyI(args ...interface{}) bool{
    return self.Object.Call("checkProperty", args).Bool()
}

// Set Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) Set(child *Sprite, key string, value interface{}) bool{
    return self.Object.Call("set", child, key, value).Bool()
}

// Set1O Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) Set1O(child *Sprite, key string, value interface{}, checkAlive bool) bool{
    return self.Object.Call("set", child, key, value, checkAlive).Bool()
}

// Set2O Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) Set2O(child *Sprite, key string, value interface{}, checkAlive bool, checkVisible bool) bool{
    return self.Object.Call("set", child, key, value, checkAlive, checkVisible).Bool()
}

// Set3O Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) Set3O(child *Sprite, key string, value interface{}, checkAlive bool, checkVisible bool, operation int) bool{
    return self.Object.Call("set", child, key, value, checkAlive, checkVisible, operation).Bool()
}

// Set4O Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) Set4O(child *Sprite, key string, value interface{}, checkAlive bool, checkVisible bool, operation int, force bool) bool{
    return self.Object.Call("set", child, key, value, checkAlive, checkVisible, operation, force).Bool()
}

// SetI Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetI(args ...interface{}) bool{
    return self.Object.Call("set", args).Bool()
}

// SetAll Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAll(key string, value interface{}) {
    self.Object.Call("setAll", key, value)
}

// SetAll1O Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAll1O(key string, value interface{}, checkAlive bool) {
    self.Object.Call("setAll", key, value, checkAlive)
}

// SetAll2O Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAll2O(key string, value interface{}, checkAlive bool, checkVisible bool) {
    self.Object.Call("setAll", key, value, checkAlive, checkVisible)
}

// SetAll3O Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAll3O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int) {
    self.Object.Call("setAll", key, value, checkAlive, checkVisible, operation)
}

// SetAll4O Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAll4O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int, force bool) {
    self.Object.Call("setAll", key, value, checkAlive, checkVisible, operation, force)
}

// SetAllI Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAllI(args ...interface{}) {
    self.Object.Call("setAll", args)
}

// SetAllChildren Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAllChildren(key string, value interface{}) {
    self.Object.Call("setAllChildren", key, value)
}

// SetAllChildren1O Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAllChildren1O(key string, value interface{}, checkAlive bool) {
    self.Object.Call("setAllChildren", key, value, checkAlive)
}

// SetAllChildren2O Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAllChildren2O(key string, value interface{}, checkAlive bool, checkVisible bool) {
    self.Object.Call("setAllChildren", key, value, checkAlive, checkVisible)
}

// SetAllChildren3O Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAllChildren3O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int) {
    self.Object.Call("setAllChildren", key, value, checkAlive, checkVisible, operation)
}

// SetAllChildren4O Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAllChildren4O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int, force bool) {
    self.Object.Call("setAllChildren", key, value, checkAlive, checkVisible, operation, force)
}

// SetAllChildrenI Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *ParticlesArcadeEmitter) SetAllChildrenI(args ...interface{}) {
    self.Object.Call("setAllChildren", args)
}

// CheckAll Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *ParticlesArcadeEmitter) CheckAll(key string, value interface{}) {
    self.Object.Call("checkAll", key, value)
}

// CheckAll1O Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *ParticlesArcadeEmitter) CheckAll1O(key string, value interface{}, checkAlive bool) {
    self.Object.Call("checkAll", key, value, checkAlive)
}

// CheckAll2O Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *ParticlesArcadeEmitter) CheckAll2O(key string, value interface{}, checkAlive bool, checkVisible bool) {
    self.Object.Call("checkAll", key, value, checkAlive, checkVisible)
}

// CheckAll3O Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *ParticlesArcadeEmitter) CheckAll3O(key string, value interface{}, checkAlive bool, checkVisible bool, force bool) {
    self.Object.Call("checkAll", key, value, checkAlive, checkVisible, force)
}

// CheckAllI Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *ParticlesArcadeEmitter) CheckAllI(args ...interface{}) {
    self.Object.Call("checkAll", args)
}

// AddAll Adds the amount to the given property on all children in this group.
// 
// `Group.addAll('x', 10)` will add 10 to the child.x value for each child.
func (self *ParticlesArcadeEmitter) AddAll(property string, amount int, checkAlive bool, checkVisible bool) {
    self.Object.Call("addAll", property, amount, checkAlive, checkVisible)
}

// AddAllI Adds the amount to the given property on all children in this group.
// 
// `Group.addAll('x', 10)` will add 10 to the child.x value for each child.
func (self *ParticlesArcadeEmitter) AddAllI(args ...interface{}) {
    self.Object.Call("addAll", args)
}

// SubAll Subtracts the amount from the given property on all children in this group.
// 
// `Group.subAll('x', 10)` will minus 10 from the child.x value for each child.
func (self *ParticlesArcadeEmitter) SubAll(property string, amount int, checkAlive bool, checkVisible bool) {
    self.Object.Call("subAll", property, amount, checkAlive, checkVisible)
}

// SubAllI Subtracts the amount from the given property on all children in this group.
// 
// `Group.subAll('x', 10)` will minus 10 from the child.x value for each child.
func (self *ParticlesArcadeEmitter) SubAllI(args ...interface{}) {
    self.Object.Call("subAll", args)
}

// MultiplyAll Multiplies the given property by the amount on all children in this group.
// 
// `Group.multiplyAll('x', 2)` will x2 the child.x value for each child.
func (self *ParticlesArcadeEmitter) MultiplyAll(property string, amount int, checkAlive bool, checkVisible bool) {
    self.Object.Call("multiplyAll", property, amount, checkAlive, checkVisible)
}

// MultiplyAllI Multiplies the given property by the amount on all children in this group.
// 
// `Group.multiplyAll('x', 2)` will x2 the child.x value for each child.
func (self *ParticlesArcadeEmitter) MultiplyAllI(args ...interface{}) {
    self.Object.Call("multiplyAll", args)
}

// DivideAll Divides the given property by the amount on all children in this group.
// 
// `Group.divideAll('x', 2)` will half the child.x value for each child.
func (self *ParticlesArcadeEmitter) DivideAll(property string, amount int, checkAlive bool, checkVisible bool) {
    self.Object.Call("divideAll", property, amount, checkAlive, checkVisible)
}

// DivideAllI Divides the given property by the amount on all children in this group.
// 
// `Group.divideAll('x', 2)` will half the child.x value for each child.
func (self *ParticlesArcadeEmitter) DivideAllI(args ...interface{}) {
    self.Object.Call("divideAll", args)
}

// CallAllExists Calls a function, specified by name, on all children in the group who exist (or do not exist).
// 
// After the existsValue parameter you can add as many parameters as you like, which will all be passed to the child callback.
func (self *ParticlesArcadeEmitter) CallAllExists(callback string, existsValue bool, parameter interface{}) {
    self.Object.Call("callAllExists", callback, existsValue, parameter)
}

// CallAllExistsI Calls a function, specified by name, on all children in the group who exist (or do not exist).
// 
// After the existsValue parameter you can add as many parameters as you like, which will all be passed to the child callback.
func (self *ParticlesArcadeEmitter) CallAllExistsI(args ...interface{}) {
    self.Object.Call("callAllExists", args)
}

// CallbackFromArray Returns a reference to a function that exists on a child of the group based on the given callback array.
func (self *ParticlesArcadeEmitter) CallbackFromArray(child interface{}, callback []interface{}, length int) {
    self.Object.Call("callbackFromArray", child, callback, length)
}

// CallbackFromArrayI Returns a reference to a function that exists on a child of the group based on the given callback array.
func (self *ParticlesArcadeEmitter) CallbackFromArrayI(args ...interface{}) {
    self.Object.Call("callbackFromArray", args)
}

// CallAll Calls a function, specified by name, on all on children.
// 
// The function is called for all children regardless if they are dead or alive (see callAllExists for different options).
// After the method parameter and context you can add as many extra parameters as you like, which will all be passed to the child.
func (self *ParticlesArcadeEmitter) CallAll(method string, context string, args interface{}) {
    self.Object.Call("callAll", method, context, args)
}

// CallAllI Calls a function, specified by name, on all on children.
// 
// The function is called for all children regardless if they are dead or alive (see callAllExists for different options).
// After the method parameter and context you can add as many extra parameters as you like, which will all be passed to the child.
func (self *ParticlesArcadeEmitter) CallAllI(args ...interface{}) {
    self.Object.Call("callAll", args)
}

// PreUpdate The core preUpdate - as called by World.
func (self *ParticlesArcadeEmitter) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI The core preUpdate - as called by World.
func (self *ParticlesArcadeEmitter) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// PostUpdate The core postUpdate - as called by World.
func (self *ParticlesArcadeEmitter) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI The core postUpdate - as called by World.
func (self *ParticlesArcadeEmitter) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Filter Find children matching a certain predicate.
// 
// For example:
// 
//     var healthyList = Group.filter(function(child, index, children) {
//         return child.health > 10 ? true : false;
//     }, true);
//     healthyList.callAll('attack');
// 
// Note: Currently this will skip any children which are Groups themselves.
func (self *ParticlesArcadeEmitter) Filter(predicate interface{}) *ArraySet{
    return &ArraySet{self.Object.Call("filter", predicate)}
}

// Filter1O Find children matching a certain predicate.
// 
// For example:
// 
//     var healthyList = Group.filter(function(child, index, children) {
//         return child.health > 10 ? true : false;
//     }, true);
//     healthyList.callAll('attack');
// 
// Note: Currently this will skip any children which are Groups themselves.
func (self *ParticlesArcadeEmitter) Filter1O(predicate interface{}, checkExists bool) *ArraySet{
    return &ArraySet{self.Object.Call("filter", predicate, checkExists)}
}

// FilterI Find children matching a certain predicate.
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
    return &ArraySet{self.Object.Call("filter", args)}
}

// ForEach Call a function on each child in this group.
// 
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
// 
//     Group.forEach(awardBonusGold, this, true, 100, 500)
// 
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
// 
// Note: This check will skip any children which are Groups themselves.
func (self *ParticlesArcadeEmitter) ForEach(callback interface{}, callbackContext interface{}) {
    self.Object.Call("forEach", callback, callbackContext)
}

// ForEach1O Call a function on each child in this group.
// 
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
// 
//     Group.forEach(awardBonusGold, this, true, 100, 500)
// 
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
// 
// Note: This check will skip any children which are Groups themselves.
func (self *ParticlesArcadeEmitter) ForEach1O(callback interface{}, callbackContext interface{}, checkExists bool) {
    self.Object.Call("forEach", callback, callbackContext, checkExists)
}

// ForEach2O Call a function on each child in this group.
// 
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
// 
//     Group.forEach(awardBonusGold, this, true, 100, 500)
// 
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
// 
// Note: This check will skip any children which are Groups themselves.
func (self *ParticlesArcadeEmitter) ForEach2O(callback interface{}, callbackContext interface{}, checkExists bool, args interface{}) {
    self.Object.Call("forEach", callback, callbackContext, checkExists, args)
}

// ForEachI Call a function on each child in this group.
// 
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
// 
//     Group.forEach(awardBonusGold, this, true, 100, 500)
// 
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
// 
// Note: This check will skip any children which are Groups themselves.
func (self *ParticlesArcadeEmitter) ForEachI(args ...interface{}) {
    self.Object.Call("forEach", args)
}

// ForEachExists Call a function on each existing child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachExists(callback interface{}, callbackContext interface{}) {
    self.Object.Call("forEachExists", callback, callbackContext)
}

// ForEachExists1O Call a function on each existing child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachExists1O(callback interface{}, callbackContext interface{}, args interface{}) {
    self.Object.Call("forEachExists", callback, callbackContext, args)
}

// ForEachExistsI Call a function on each existing child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachExistsI(args ...interface{}) {
    self.Object.Call("forEachExists", args)
}

// ForEachAlive Call a function on each alive child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachAlive(callback interface{}, callbackContext interface{}) {
    self.Object.Call("forEachAlive", callback, callbackContext)
}

// ForEachAlive1O Call a function on each alive child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachAlive1O(callback interface{}, callbackContext interface{}, args interface{}) {
    self.Object.Call("forEachAlive", callback, callbackContext, args)
}

// ForEachAliveI Call a function on each alive child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachAliveI(args ...interface{}) {
    self.Object.Call("forEachAlive", args)
}

// ForEachDead Call a function on each dead child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachDead(callback interface{}, callbackContext interface{}) {
    self.Object.Call("forEachDead", callback, callbackContext)
}

// ForEachDead1O Call a function on each dead child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachDead1O(callback interface{}, callbackContext interface{}, args interface{}) {
    self.Object.Call("forEachDead", callback, callbackContext, args)
}

// ForEachDeadI Call a function on each dead child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *ParticlesArcadeEmitter) ForEachDeadI(args ...interface{}) {
    self.Object.Call("forEachDead", args)
}

// Sort Sort the children in the group according to a particular key and ordering.
// 
// Call this function to sort the group according to a particular key value and order.
// 
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
// 
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *ParticlesArcadeEmitter) Sort() {
    self.Object.Call("sort")
}

// Sort1O Sort the children in the group according to a particular key and ordering.
// 
// Call this function to sort the group according to a particular key value and order.
// 
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
// 
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *ParticlesArcadeEmitter) Sort1O(key string) {
    self.Object.Call("sort", key)
}

// Sort2O Sort the children in the group according to a particular key and ordering.
// 
// Call this function to sort the group according to a particular key value and order.
// 
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
// 
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *ParticlesArcadeEmitter) Sort2O(key string, order int) {
    self.Object.Call("sort", key, order)
}

// SortI Sort the children in the group according to a particular key and ordering.
// 
// Call this function to sort the group according to a particular key value and order.
// 
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
// 
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *ParticlesArcadeEmitter) SortI(args ...interface{}) {
    self.Object.Call("sort", args)
}

// CustomSort Sort the children in the group according to custom sort function.
// 
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *ParticlesArcadeEmitter) CustomSort(sortHandler interface{}) {
    self.Object.Call("customSort", sortHandler)
}

// CustomSort1O Sort the children in the group according to custom sort function.
// 
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *ParticlesArcadeEmitter) CustomSort1O(sortHandler interface{}, context interface{}) {
    self.Object.Call("customSort", sortHandler, context)
}

// CustomSortI Sort the children in the group according to custom sort function.
// 
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *ParticlesArcadeEmitter) CustomSortI(args ...interface{}) {
    self.Object.Call("customSort", args)
}

// AscendingSortHandler An internal helper function for the sort process.
func (self *ParticlesArcadeEmitter) AscendingSortHandler(a interface{}, b interface{}) {
    self.Object.Call("ascendingSortHandler", a, b)
}

// AscendingSortHandlerI An internal helper function for the sort process.
func (self *ParticlesArcadeEmitter) AscendingSortHandlerI(args ...interface{}) {
    self.Object.Call("ascendingSortHandler", args)
}

// DescendingSortHandler An internal helper function for the sort process.
func (self *ParticlesArcadeEmitter) DescendingSortHandler(a interface{}, b interface{}) {
    self.Object.Call("descendingSortHandler", a, b)
}

// DescendingSortHandlerI An internal helper function for the sort process.
func (self *ParticlesArcadeEmitter) DescendingSortHandlerI(args ...interface{}) {
    self.Object.Call("descendingSortHandler", args)
}

// Iterate Iterates over the children of the group performing one of several actions for matched children.
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
func (self *ParticlesArcadeEmitter) Iterate(key string, value interface{}, returnType int) interface{}{
    return self.Object.Call("iterate", key, value, returnType)
}

// Iterate1O Iterates over the children of the group performing one of several actions for matched children.
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
func (self *ParticlesArcadeEmitter) Iterate1O(key string, value interface{}, returnType int, callback interface{}) interface{}{
    return self.Object.Call("iterate", key, value, returnType, callback)
}

// Iterate2O Iterates over the children of the group performing one of several actions for matched children.
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
func (self *ParticlesArcadeEmitter) Iterate2O(key string, value interface{}, returnType int, callback interface{}, callbackContext interface{}) interface{}{
    return self.Object.Call("iterate", key, value, returnType, callback, callbackContext)
}

// Iterate3O Iterates over the children of the group performing one of several actions for matched children.
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
func (self *ParticlesArcadeEmitter) Iterate3O(key string, value interface{}, returnType int, callback interface{}, callbackContext interface{}, args []interface{}) interface{}{
    return self.Object.Call("iterate", key, value, returnType, callback, callbackContext, args)
}

// IterateI Iterates over the children of the group performing one of several actions for matched children.
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
    return self.Object.Call("iterate", args)
}

// GetFirstExists Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstExists() *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists")}
}

// GetFirstExists1O Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstExists1O(exists bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists)}
}

// GetFirstExists2O Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstExists2O(exists bool, createIfNull bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull)}
}

// GetFirstExists3O Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstExists3O(exists bool, createIfNull bool, x int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x)}
}

// GetFirstExists4O Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstExists4O(exists bool, createIfNull bool, x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x, y)}
}

// GetFirstExists5O Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstExists5O(exists bool, createIfNull bool, x int, y int, key interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x, y, key)}
}

// GetFirstExists6O Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstExists6O(exists bool, createIfNull bool, x int, y int, key interface{}, frame interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x, y, key, frame)}
}

// GetFirstExistsI Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstExistsI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", args)}
}

// GetFirstAlive Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstAlive() *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive")}
}

// GetFirstAlive1O Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstAlive1O(createIfNull bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull)}
}

// GetFirstAlive2O Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstAlive2O(createIfNull bool, x int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x)}
}

// GetFirstAlive3O Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstAlive3O(createIfNull bool, x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x, y)}
}

// GetFirstAlive4O Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstAlive4O(createIfNull bool, x int, y int, key interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x, y, key)}
}

// GetFirstAlive5O Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstAlive5O(createIfNull bool, x int, y int, key interface{}, frame interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x, y, key, frame)}
}

// GetFirstAliveI Get the first child that is alive (`child.alive === true`).
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
    return &DisplayObject{self.Object.Call("getFirstAlive", args)}
}

// GetFirstDead Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstDead() *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead")}
}

// GetFirstDead1O Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstDead1O(createIfNull bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", createIfNull)}
}

// GetFirstDead2O Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstDead2O(createIfNull bool, x int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x)}
}

// GetFirstDead3O Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstDead3O(createIfNull bool, x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x, y)}
}

// GetFirstDead4O Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstDead4O(createIfNull bool, x int, y int, key interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x, y, key)}
}

// GetFirstDead5O Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *ParticlesArcadeEmitter) GetFirstDead5O(createIfNull bool, x int, y int, key interface{}, frame interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x, y, key, frame)}
}

// GetFirstDeadI Get the first child that is dead (`child.alive === false`).
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
    return &DisplayObject{self.Object.Call("getFirstDead", args)}
}

// ResetChild Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *ParticlesArcadeEmitter) ResetChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", child)}
}

// ResetChild1O Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *ParticlesArcadeEmitter) ResetChild1O(child *DisplayObject, x int) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", child, x)}
}

// ResetChild2O Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *ParticlesArcadeEmitter) ResetChild2O(child *DisplayObject, x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", child, x, y)}
}

// ResetChild3O Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *ParticlesArcadeEmitter) ResetChild3O(child *DisplayObject, x int, y int, key interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", child, x, y, key)}
}

// ResetChild4O Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *ParticlesArcadeEmitter) ResetChild4O(child *DisplayObject, x int, y int, key interface{}, frame interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", child, x, y, key, frame)}
}

// ResetChildI Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *ParticlesArcadeEmitter) ResetChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", args)}
}

// GetTop Return the child at the top of this group.
// 
// The top child is the child displayed (rendered) above every other child.
func (self *ParticlesArcadeEmitter) GetTop() interface{}{
    return self.Object.Call("getTop")
}

// GetTopI Return the child at the top of this group.
// 
// The top child is the child displayed (rendered) above every other child.
func (self *ParticlesArcadeEmitter) GetTopI(args ...interface{}) interface{}{
    return self.Object.Call("getTop", args)
}

// GetBottom Returns the child at the bottom of this group.
// 
// The bottom child the child being displayed (rendered) below every other child.
func (self *ParticlesArcadeEmitter) GetBottom() interface{}{
    return self.Object.Call("getBottom")
}

// GetBottomI Returns the child at the bottom of this group.
// 
// The bottom child the child being displayed (rendered) below every other child.
func (self *ParticlesArcadeEmitter) GetBottomI(args ...interface{}) interface{}{
    return self.Object.Call("getBottom", args)
}

// GetClosestTo Get the closest child to given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *ParticlesArcadeEmitter) GetClosestTo(object interface{}) interface{}{
    return self.Object.Call("getClosestTo", object)
}

// GetClosestTo1O Get the closest child to given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *ParticlesArcadeEmitter) GetClosestTo1O(object interface{}, callback interface{}) interface{}{
    return self.Object.Call("getClosestTo", object, callback)
}

// GetClosestTo2O Get the closest child to given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *ParticlesArcadeEmitter) GetClosestTo2O(object interface{}, callback interface{}, callbackContext interface{}) interface{}{
    return self.Object.Call("getClosestTo", object, callback, callbackContext)
}

// GetClosestToI Get the closest child to given Object, with optional callback to filter children.
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
    return self.Object.Call("getClosestTo", args)
}

// GetFurthestFrom Get the child furthest away from the given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *ParticlesArcadeEmitter) GetFurthestFrom(object interface{}) interface{}{
    return self.Object.Call("getFurthestFrom", object)
}

// GetFurthestFrom1O Get the child furthest away from the given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *ParticlesArcadeEmitter) GetFurthestFrom1O(object interface{}, callback interface{}) interface{}{
    return self.Object.Call("getFurthestFrom", object, callback)
}

// GetFurthestFrom2O Get the child furthest away from the given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *ParticlesArcadeEmitter) GetFurthestFrom2O(object interface{}, callback interface{}, callbackContext interface{}) interface{}{
    return self.Object.Call("getFurthestFrom", object, callback, callbackContext)
}

// GetFurthestFromI Get the child furthest away from the given Object, with optional callback to filter children.
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
    return self.Object.Call("getFurthestFrom", args)
}

// CountLiving Get the number of living children in this group.
func (self *ParticlesArcadeEmitter) CountLiving() int{
    return self.Object.Call("countLiving").Int()
}

// CountLivingI Get the number of living children in this group.
func (self *ParticlesArcadeEmitter) CountLivingI(args ...interface{}) int{
    return self.Object.Call("countLiving", args).Int()
}

// CountDead Get the number of dead children in this group.
func (self *ParticlesArcadeEmitter) CountDead() int{
    return self.Object.Call("countDead").Int()
}

// CountDeadI Get the number of dead children in this group.
func (self *ParticlesArcadeEmitter) CountDeadI(args ...interface{}) int{
    return self.Object.Call("countDead", args).Int()
}

// GetRandom Returns a random child from the group.
func (self *ParticlesArcadeEmitter) GetRandom() interface{}{
    return self.Object.Call("getRandom")
}

// GetRandom1O Returns a random child from the group.
func (self *ParticlesArcadeEmitter) GetRandom1O(startIndex int) interface{}{
    return self.Object.Call("getRandom", startIndex)
}

// GetRandom2O Returns a random child from the group.
func (self *ParticlesArcadeEmitter) GetRandom2O(startIndex int, length int) interface{}{
    return self.Object.Call("getRandom", startIndex, length)
}

// GetRandomI Returns a random child from the group.
func (self *ParticlesArcadeEmitter) GetRandomI(args ...interface{}) interface{}{
    return self.Object.Call("getRandom", args)
}

// Remove Removes the given child from this group.
// 
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
// 
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *ParticlesArcadeEmitter) Remove(child interface{}) bool{
    return self.Object.Call("remove", child).Bool()
}

// Remove1O Removes the given child from this group.
// 
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
// 
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *ParticlesArcadeEmitter) Remove1O(child interface{}, destroy bool) bool{
    return self.Object.Call("remove", child, destroy).Bool()
}

// Remove2O Removes the given child from this group.
// 
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
// 
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *ParticlesArcadeEmitter) Remove2O(child interface{}, destroy bool, silent bool) bool{
    return self.Object.Call("remove", child, destroy, silent).Bool()
}

// RemoveI Removes the given child from this group.
// 
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
// 
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *ParticlesArcadeEmitter) RemoveI(args ...interface{}) bool{
    return self.Object.Call("remove", args).Bool()
}

// MoveAll Moves all children from this Group to the Group given.
func (self *ParticlesArcadeEmitter) MoveAll(group *Group) *Group{
    return &Group{self.Object.Call("moveAll", group)}
}

// MoveAll1O Moves all children from this Group to the Group given.
func (self *ParticlesArcadeEmitter) MoveAll1O(group *Group, silent bool) *Group{
    return &Group{self.Object.Call("moveAll", group, silent)}
}

// MoveAllI Moves all children from this Group to the Group given.
func (self *ParticlesArcadeEmitter) MoveAllI(args ...interface{}) *Group{
    return &Group{self.Object.Call("moveAll", args)}
}

// RemoveAll Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *ParticlesArcadeEmitter) RemoveAll() {
    self.Object.Call("removeAll")
}

// RemoveAll1O Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *ParticlesArcadeEmitter) RemoveAll1O(destroy bool) {
    self.Object.Call("removeAll", destroy)
}

// RemoveAll2O Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *ParticlesArcadeEmitter) RemoveAll2O(destroy bool, silent bool) {
    self.Object.Call("removeAll", destroy, silent)
}

// RemoveAll3O Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *ParticlesArcadeEmitter) RemoveAll3O(destroy bool, silent bool, destroyTexture bool) {
    self.Object.Call("removeAll", destroy, silent, destroyTexture)
}

// RemoveAllI Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *ParticlesArcadeEmitter) RemoveAllI(args ...interface{}) {
    self.Object.Call("removeAll", args)
}

// RemoveBetween Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *ParticlesArcadeEmitter) RemoveBetween(startIndex int) {
    self.Object.Call("removeBetween", startIndex)
}

// RemoveBetween1O Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *ParticlesArcadeEmitter) RemoveBetween1O(startIndex int, endIndex int) {
    self.Object.Call("removeBetween", startIndex, endIndex)
}

// RemoveBetween2O Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *ParticlesArcadeEmitter) RemoveBetween2O(startIndex int, endIndex int, destroy bool) {
    self.Object.Call("removeBetween", startIndex, endIndex, destroy)
}

// RemoveBetween3O Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *ParticlesArcadeEmitter) RemoveBetween3O(startIndex int, endIndex int, destroy bool, silent bool) {
    self.Object.Call("removeBetween", startIndex, endIndex, destroy, silent)
}

// RemoveBetweenI Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *ParticlesArcadeEmitter) RemoveBetweenI(args ...interface{}) {
    self.Object.Call("removeBetween", args)
}

// AlignIn Aligns this Group within another Game Object, or Rectangle, known as the
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
func (self *ParticlesArcadeEmitter) AlignIn(container interface{}) *Group{
    return &Group{self.Object.Call("alignIn", container)}
}

// AlignIn1O Aligns this Group within another Game Object, or Rectangle, known as the
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
func (self *ParticlesArcadeEmitter) AlignIn1O(container interface{}, position int) *Group{
    return &Group{self.Object.Call("alignIn", container, position)}
}

// AlignIn2O Aligns this Group within another Game Object, or Rectangle, known as the
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
func (self *ParticlesArcadeEmitter) AlignIn2O(container interface{}, position int, offsetX int) *Group{
    return &Group{self.Object.Call("alignIn", container, position, offsetX)}
}

// AlignIn3O Aligns this Group within another Game Object, or Rectangle, known as the
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
func (self *ParticlesArcadeEmitter) AlignIn3O(container interface{}, position int, offsetX int, offsetY int) *Group{
    return &Group{self.Object.Call("alignIn", container, position, offsetX, offsetY)}
}

// AlignInI Aligns this Group within another Game Object, or Rectangle, known as the
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
    return &Group{self.Object.Call("alignIn", args)}
}

// AlignTo Aligns this Group to the side of another Game Object, or Rectangle, known as the
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
func (self *ParticlesArcadeEmitter) AlignTo(parent interface{}) *Group{
    return &Group{self.Object.Call("alignTo", parent)}
}

// AlignTo1O Aligns this Group to the side of another Game Object, or Rectangle, known as the
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
func (self *ParticlesArcadeEmitter) AlignTo1O(parent interface{}, position int) *Group{
    return &Group{self.Object.Call("alignTo", parent, position)}
}

// AlignTo2O Aligns this Group to the side of another Game Object, or Rectangle, known as the
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
func (self *ParticlesArcadeEmitter) AlignTo2O(parent interface{}, position int, offsetX int) *Group{
    return &Group{self.Object.Call("alignTo", parent, position, offsetX)}
}

// AlignTo3O Aligns this Group to the side of another Game Object, or Rectangle, known as the
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
func (self *ParticlesArcadeEmitter) AlignTo3O(parent interface{}, position int, offsetX int, offsetY int) *Group{
    return &Group{self.Object.Call("alignTo", parent, position, offsetX, offsetY)}
}

// AlignToI Aligns this Group to the side of another Game Object, or Rectangle, known as the
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
    return &Group{self.Object.Call("alignTo", args)}
}

// AddChild Adds a child to the container.
func (self *ParticlesArcadeEmitter) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// AddChildI Adds a child to the container.
func (self *ParticlesArcadeEmitter) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// AddChildAt Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *ParticlesArcadeEmitter) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// AddChildAtI Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *ParticlesArcadeEmitter) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// SwapChildren Swaps the position of 2 Display Objects within this container.
func (self *ParticlesArcadeEmitter) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// SwapChildrenI Swaps the position of 2 Display Objects within this container.
func (self *ParticlesArcadeEmitter) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// GetChildIndex Returns the index position of a child DisplayObject instance
func (self *ParticlesArcadeEmitter) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// GetChildIndexI Returns the index position of a child DisplayObject instance
func (self *ParticlesArcadeEmitter) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// SetChildIndex Changes the position of an existing child in the display object container
func (self *ParticlesArcadeEmitter) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// SetChildIndexI Changes the position of an existing child in the display object container
func (self *ParticlesArcadeEmitter) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// GetChildAt Returns the child at the specified index
func (self *ParticlesArcadeEmitter) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// GetChildAtI Returns the child at the specified index
func (self *ParticlesArcadeEmitter) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// RemoveChild Removes a child from the container.
func (self *ParticlesArcadeEmitter) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// RemoveChildI Removes a child from the container.
func (self *ParticlesArcadeEmitter) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// RemoveChildAt Removes a child from the specified index position.
func (self *ParticlesArcadeEmitter) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// RemoveChildAtI Removes a child from the specified index position.
func (self *ParticlesArcadeEmitter) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// RemoveChildren Removes all children from this container that are within the begin and end indexes.
func (self *ParticlesArcadeEmitter) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// RemoveChildrenI Removes all children from this container that are within the begin and end indexes.
func (self *ParticlesArcadeEmitter) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// GetBounds Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *ParticlesArcadeEmitter) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// GetBoundsI Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *ParticlesArcadeEmitter) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// GetLocalBounds Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *ParticlesArcadeEmitter) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// GetLocalBoundsI Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *ParticlesArcadeEmitter) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// SetStageReference Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *ParticlesArcadeEmitter) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// SetStageReferenceI Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *ParticlesArcadeEmitter) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// RemoveStageReference Removes the current stage reference from the container and all of its children.
func (self *ParticlesArcadeEmitter) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// RemoveStageReferenceI Removes the current stage reference from the container and all of its children.
func (self *ParticlesArcadeEmitter) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// _renderWebGL Renders the object using the WebGL renderer
func (self *ParticlesArcadeEmitter) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// _renderWebGLI Renders the object using the WebGL renderer
func (self *ParticlesArcadeEmitter) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// _renderCanvas Renders the object using the Canvas renderer
func (self *ParticlesArcadeEmitter) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// _renderCanvasI Renders the object using the Canvas renderer
func (self *ParticlesArcadeEmitter) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

