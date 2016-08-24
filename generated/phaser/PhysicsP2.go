// Package phaser Automatic generation for Phaser.Physics.P2
// generated file PhysicsP2.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsP2 This is your main access to the P2 Physics World.
// From here you can create materials, listen for events and add bodies into the physics simulation.
type PhysicsP2 struct {
    *js.Object
}

// NewPhysicsP2 This is your main access to the P2 Physics World.
// From here you can create materials, listen for events and add bodies into the physics simulation.
func NewPhysicsP2(game *Game) *PhysicsP2 {
    return &PhysicsP2{js.Global.Get("Phaser").Get("Physics").Get("P2").New(game)}
}
// NewPhysicsP21O This is your main access to the P2 Physics World.
// From here you can create materials, listen for events and add bodies into the physics simulation.
func NewPhysicsP21O(game *Game, config interface{}) *PhysicsP2 {
    return &PhysicsP2{js.Global.Get("Phaser").Get("Physics").Get("P2").New(game, config)}
}
// NewPhysicsP2I This is your main access to the P2 Physics World.
// From here you can create materials, listen for events and add bodies into the physics simulation.
func NewPhysicsP2I(args ...interface{}) *PhysicsP2 {
    return &PhysicsP2{js.Global.Get("Phaser").Get("Physics").Get("P2").New(args)}
}



// Game Local reference to game.
func (self *PhysicsP2) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *PhysicsP2) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Config The p2 World configuration object.
func (self *PhysicsP2) Config() interface{}{
    return self.Object.Get("config")
}

// SetConfigA The p2 World configuration object.
func (self *PhysicsP2) SetConfigA(member interface{}) {
    self.Object.Set("config", member)
}

// World The p2 World in which the simulation is run.
func (self *PhysicsP2) World() *P2World{
    return &P2World{self.Object.Get("world")}
}

// SetWorldA The p2 World in which the simulation is run.
func (self *PhysicsP2) SetWorldA(member *P2World) {
    self.Object.Set("world", member)
}

// FrameRate The frame rate the world will be stepped at. Defaults to 1 / 60, but you can change here. Also see useElapsedTime property.
func (self *PhysicsP2) FrameRate() int{
    return self.Object.Get("frameRate").Int()
}

// SetFrameRateA The frame rate the world will be stepped at. Defaults to 1 / 60, but you can change here. Also see useElapsedTime property.
func (self *PhysicsP2) SetFrameRateA(member int) {
    self.Object.Set("frameRate", member)
}

// UseElapsedTime If true the frameRate value will be ignored and instead p2 will step with the value of Game.Time.physicsElapsed, which is a delta time value.
func (self *PhysicsP2) UseElapsedTime() bool{
    return self.Object.Get("useElapsedTime").Bool()
}

// SetUseElapsedTimeA If true the frameRate value will be ignored and instead p2 will step with the value of Game.Time.physicsElapsed, which is a delta time value.
func (self *PhysicsP2) SetUseElapsedTimeA(member bool) {
    self.Object.Set("useElapsedTime", member)
}

// Paused The paused state of the P2 World.
func (self *PhysicsP2) Paused() bool{
    return self.Object.Get("paused").Bool()
}

// SetPausedA The paused state of the P2 World.
func (self *PhysicsP2) SetPausedA(member bool) {
    self.Object.Set("paused", member)
}

// Materials A local array of all created Materials.
func (self *PhysicsP2) Materials() []PhysicsP2Material{
	array00 := self.Object.Get("materials")
	length00 := array00.Length()
	out00 := make([]PhysicsP2Material, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = PhysicsP2Material{array00.Index(i00)}
	}
	return out00
}

// SetMaterialsA A local array of all created Materials.
func (self *PhysicsP2) SetMaterialsA(member []PhysicsP2Material) {
    self.Object.Set("materials", member)
}

// Gravity The gravity applied to all bodies each step.
func (self *PhysicsP2) Gravity() *PhysicsP2InversePointProxy{
    return &PhysicsP2InversePointProxy{self.Object.Get("gravity")}
}

// SetGravityA The gravity applied to all bodies each step.
func (self *PhysicsP2) SetGravityA(member *PhysicsP2InversePointProxy) {
    self.Object.Set("gravity", member)
}

// Walls An object containing the 4 wall bodies that bound the physics world.
func (self *PhysicsP2) Walls() interface{}{
    return self.Object.Get("walls")
}

// SetWallsA An object containing the 4 wall bodies that bound the physics world.
func (self *PhysicsP2) SetWallsA(member interface{}) {
    self.Object.Set("walls", member)
}

// OnBodyAdded This signal is dispatched when a new Body is added to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was added to the world.
func (self *PhysicsP2) OnBodyAdded() *Signal{
    return &Signal{self.Object.Get("onBodyAdded")}
}

// SetOnBodyAddedA This signal is dispatched when a new Body is added to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was added to the world.
func (self *PhysicsP2) SetOnBodyAddedA(member *Signal) {
    self.Object.Set("onBodyAdded", member)
}

// OnBodyRemoved This signal is dispatched when a Body is removed to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was removed from the world.
func (self *PhysicsP2) OnBodyRemoved() *Signal{
    return &Signal{self.Object.Get("onBodyRemoved")}
}

// SetOnBodyRemovedA This signal is dispatched when a Body is removed to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was removed from the world.
func (self *PhysicsP2) SetOnBodyRemovedA(member *Signal) {
    self.Object.Set("onBodyRemoved", member)
}

// OnSpringAdded This signal is dispatched when a Spring is added to the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was added to the world.
func (self *PhysicsP2) OnSpringAdded() *Signal{
    return &Signal{self.Object.Get("onSpringAdded")}
}

// SetOnSpringAddedA This signal is dispatched when a Spring is added to the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was added to the world.
func (self *PhysicsP2) SetOnSpringAddedA(member *Signal) {
    self.Object.Set("onSpringAdded", member)
}

// OnSpringRemoved This signal is dispatched when a Spring is removed from the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was removed from the world.
func (self *PhysicsP2) OnSpringRemoved() *Signal{
    return &Signal{self.Object.Get("onSpringRemoved")}
}

// SetOnSpringRemovedA This signal is dispatched when a Spring is removed from the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was removed from the world.
func (self *PhysicsP2) SetOnSpringRemovedA(member *Signal) {
    self.Object.Set("onSpringRemoved", member)
}

// OnConstraintAdded This signal is dispatched when a Constraint is added to the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was added to the world.
func (self *PhysicsP2) OnConstraintAdded() *Signal{
    return &Signal{self.Object.Get("onConstraintAdded")}
}

// SetOnConstraintAddedA This signal is dispatched when a Constraint is added to the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was added to the world.
func (self *PhysicsP2) SetOnConstraintAddedA(member *Signal) {
    self.Object.Set("onConstraintAdded", member)
}

// OnConstraintRemoved This signal is dispatched when a Constraint is removed from the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was removed from the world.
func (self *PhysicsP2) OnConstraintRemoved() *Signal{
    return &Signal{self.Object.Get("onConstraintRemoved")}
}

// SetOnConstraintRemovedA This signal is dispatched when a Constraint is removed from the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was removed from the world.
func (self *PhysicsP2) SetOnConstraintRemovedA(member *Signal) {
    self.Object.Set("onConstraintRemoved", member)
}

// OnContactMaterialAdded This signal is dispatched when a Contact Material is added to the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was added to the world.
func (self *PhysicsP2) OnContactMaterialAdded() *Signal{
    return &Signal{self.Object.Get("onContactMaterialAdded")}
}

// SetOnContactMaterialAddedA This signal is dispatched when a Contact Material is added to the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was added to the world.
func (self *PhysicsP2) SetOnContactMaterialAddedA(member *Signal) {
    self.Object.Set("onContactMaterialAdded", member)
}

// OnContactMaterialRemoved This signal is dispatched when a Contact Material is removed from the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was removed from the world.
func (self *PhysicsP2) OnContactMaterialRemoved() *Signal{
    return &Signal{self.Object.Get("onContactMaterialRemoved")}
}

// SetOnContactMaterialRemovedA This signal is dispatched when a Contact Material is removed from the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was removed from the world.
func (self *PhysicsP2) SetOnContactMaterialRemovedA(member *Signal) {
    self.Object.Set("onContactMaterialRemoved", member)
}

// PostBroadphaseCallback A postBroadphase callback.
func (self *PhysicsP2) PostBroadphaseCallback() interface{}{
    return self.Object.Get("postBroadphaseCallback")
}

// SetPostBroadphaseCallbackA A postBroadphase callback.
func (self *PhysicsP2) SetPostBroadphaseCallbackA(member interface{}) {
    self.Object.Set("postBroadphaseCallback", member)
}

// CallbackContext The context under which the callbacks are fired.
func (self *PhysicsP2) CallbackContext() interface{}{
    return self.Object.Get("callbackContext")
}

// SetCallbackContextA The context under which the callbacks are fired.
func (self *PhysicsP2) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// OnBeginContact This Signal is dispatched when a first contact is created between two bodies. This happens *before* the step has been done.
// 
// It sends 5 arguments: `bodyA`, `bodyB`, `shapeA`, `shapeB` and `contactEquations`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) OnBeginContact() *Signal{
    return &Signal{self.Object.Get("onBeginContact")}
}

// SetOnBeginContactA This Signal is dispatched when a first contact is created between two bodies. This happens *before* the step has been done.
// 
// It sends 5 arguments: `bodyA`, `bodyB`, `shapeA`, `shapeB` and `contactEquations`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) SetOnBeginContactA(member *Signal) {
    self.Object.Set("onBeginContact", member)
}

// OnEndContact This Signal is dispatched when final contact occurs between two bodies. This happens *before* the step has been done.
// 
// It sends 4 arguments: `bodyA`, `bodyB`, `shapeA` and `shapeB`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) OnEndContact() *Signal{
    return &Signal{self.Object.Get("onEndContact")}
}

// SetOnEndContactA This Signal is dispatched when final contact occurs between two bodies. This happens *before* the step has been done.
// 
// It sends 4 arguments: `bodyA`, `bodyB`, `shapeA` and `shapeB`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) SetOnEndContactA(member *Signal) {
    self.Object.Set("onEndContact", member)
}

// CollisionGroups An array containing the collision groups that have been defined in the World.
func (self *PhysicsP2) CollisionGroups() []interface{}{
	array00 := self.Object.Get("collisionGroups")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetCollisionGroupsA An array containing the collision groups that have been defined in the World.
func (self *PhysicsP2) SetCollisionGroupsA(member []interface{}) {
    self.Object.Set("collisionGroups", member)
}

// NothingCollisionGroup A default collision group.
func (self *PhysicsP2) NothingCollisionGroup() *PhysicsP2CollisionGroup{
    return &PhysicsP2CollisionGroup{self.Object.Get("nothingCollisionGroup")}
}

// SetNothingCollisionGroupA A default collision group.
func (self *PhysicsP2) SetNothingCollisionGroupA(member *PhysicsP2CollisionGroup) {
    self.Object.Set("nothingCollisionGroup", member)
}

// BoundsCollisionGroup A default collision group.
func (self *PhysicsP2) BoundsCollisionGroup() *PhysicsP2CollisionGroup{
    return &PhysicsP2CollisionGroup{self.Object.Get("boundsCollisionGroup")}
}

// SetBoundsCollisionGroupA A default collision group.
func (self *PhysicsP2) SetBoundsCollisionGroupA(member *PhysicsP2CollisionGroup) {
    self.Object.Set("boundsCollisionGroup", member)
}

// EverythingCollisionGroup A default collision group.
func (self *PhysicsP2) EverythingCollisionGroup() *PhysicsP2CollisionGroup{
    return &PhysicsP2CollisionGroup{self.Object.Get("everythingCollisionGroup")}
}

// SetEverythingCollisionGroupA A default collision group.
func (self *PhysicsP2) SetEverythingCollisionGroupA(member *PhysicsP2CollisionGroup) {
    self.Object.Set("everythingCollisionGroup", member)
}

// BoundsCollidesWith An array of the bodies the world bounds collides with.
func (self *PhysicsP2) BoundsCollidesWith() []interface{}{
	array00 := self.Object.Get("boundsCollidesWith")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetBoundsCollidesWithA An array of the bodies the world bounds collides with.
func (self *PhysicsP2) SetBoundsCollidesWithA(member []interface{}) {
    self.Object.Set("boundsCollidesWith", member)
}

// Friction Friction between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) Friction() int{
    return self.Object.Get("friction").Int()
}

// SetFrictionA Friction between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) SetFrictionA(member int) {
    self.Object.Set("friction", member)
}

// Restitution Default coefficient of restitution between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) Restitution() int{
    return self.Object.Get("restitution").Int()
}

// SetRestitutionA Default coefficient of restitution between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) SetRestitutionA(member int) {
    self.Object.Set("restitution", member)
}

// ContactMaterial The default Contact Material being used by the World.
func (self *PhysicsP2) ContactMaterial() *P2ContactMaterial{
    return &P2ContactMaterial{self.Object.Get("contactMaterial")}
}

// SetContactMaterialA The default Contact Material being used by the World.
func (self *PhysicsP2) SetContactMaterialA(member *P2ContactMaterial) {
    self.Object.Set("contactMaterial", member)
}

// ApplySpringForces Enable to automatically apply spring forces each step.
func (self *PhysicsP2) ApplySpringForces() bool{
    return self.Object.Get("applySpringForces").Bool()
}

// SetApplySpringForcesA Enable to automatically apply spring forces each step.
func (self *PhysicsP2) SetApplySpringForcesA(member bool) {
    self.Object.Set("applySpringForces", member)
}

// ApplyDamping Enable to automatically apply body damping each step.
func (self *PhysicsP2) ApplyDamping() bool{
    return self.Object.Get("applyDamping").Bool()
}

// SetApplyDampingA Enable to automatically apply body damping each step.
func (self *PhysicsP2) SetApplyDampingA(member bool) {
    self.Object.Set("applyDamping", member)
}

// ApplyGravity Enable to automatically apply gravity each step.
func (self *PhysicsP2) ApplyGravity() bool{
    return self.Object.Get("applyGravity").Bool()
}

// SetApplyGravityA Enable to automatically apply gravity each step.
func (self *PhysicsP2) SetApplyGravityA(member bool) {
    self.Object.Set("applyGravity", member)
}

// SolveConstraints Enable/disable constraint solving in each step.
func (self *PhysicsP2) SolveConstraints() bool{
    return self.Object.Get("solveConstraints").Bool()
}

// SetSolveConstraintsA Enable/disable constraint solving in each step.
func (self *PhysicsP2) SetSolveConstraintsA(member bool) {
    self.Object.Set("solveConstraints", member)
}

// Time The World time.
func (self *PhysicsP2) Time() bool{
    return self.Object.Get("time").Bool()
}

// SetTimeA The World time.
func (self *PhysicsP2) SetTimeA(member bool) {
    self.Object.Set("time", member)
}

// EmitImpactEvent Set to true if you want to the world to emit the "impact" event. Turning this off could improve performance.
func (self *PhysicsP2) EmitImpactEvent() bool{
    return self.Object.Get("emitImpactEvent").Bool()
}

// SetEmitImpactEventA Set to true if you want to the world to emit the "impact" event. Turning this off could improve performance.
func (self *PhysicsP2) SetEmitImpactEventA(member bool) {
    self.Object.Set("emitImpactEvent", member)
}

// SleepMode How to deactivate bodies during simulation. Possible modes are: World.NO_SLEEPING, World.BODY_SLEEPING and World.ISLAND_SLEEPING.
// If sleeping is enabled, you might need to wake up the bodies if they fall asleep when they shouldn't. If you want to enable sleeping in the world, but want to disable it for a particular body, see Body.allowSleep.
func (self *PhysicsP2) SleepMode() int{
    return self.Object.Get("sleepMode").Int()
}

// SetSleepModeA How to deactivate bodies during simulation. Possible modes are: World.NO_SLEEPING, World.BODY_SLEEPING and World.ISLAND_SLEEPING.
// If sleeping is enabled, you might need to wake up the bodies if they fall asleep when they shouldn't. If you want to enable sleeping in the world, but want to disable it for a particular body, see Body.allowSleep.
func (self *PhysicsP2) SetSleepModeA(member int) {
    self.Object.Set("sleepMode", member)
}

// Total The total number of bodies in the world.
func (self *PhysicsP2) Total() int{
    return self.Object.Get("total").Int()
}

// SetTotalA The total number of bodies in the world.
func (self *PhysicsP2) SetTotalA(member int) {
    self.Object.Set("total", member)
}


// RemoveBodyNextStep This will add a P2 Physics body into the removal list for the next step.
func (self *PhysicsP2) RemoveBodyNextStep(body *PhysicsP2Body) {
    self.Object.Call("removeBodyNextStep", body)
}

// RemoveBodyNextStepI This will add a P2 Physics body into the removal list for the next step.
func (self *PhysicsP2) RemoveBodyNextStepI(args ...interface{}) {
    self.Object.Call("removeBodyNextStep", args)
}

// PreUpdate Called at the start of the core update loop. Purges flagged bodies from the world.
func (self *PhysicsP2) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI Called at the start of the core update loop. Purges flagged bodies from the world.
func (self *PhysicsP2) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Enable This will create a P2 Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// Note: When the game object is enabled for P2 physics it has its anchor x/y set to 0.5 so it becomes centered.
func (self *PhysicsP2) Enable(object interface{}) {
    self.Object.Call("enable", object)
}

// Enable1O This will create a P2 Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// Note: When the game object is enabled for P2 physics it has its anchor x/y set to 0.5 so it becomes centered.
func (self *PhysicsP2) Enable1O(object interface{}, debug bool) {
    self.Object.Call("enable", object, debug)
}

// Enable2O This will create a P2 Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// Note: When the game object is enabled for P2 physics it has its anchor x/y set to 0.5 so it becomes centered.
func (self *PhysicsP2) Enable2O(object interface{}, debug bool, children bool) {
    self.Object.Call("enable", object, debug, children)
}

// EnableI This will create a P2 Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// Note: When the game object is enabled for P2 physics it has its anchor x/y set to 0.5 so it becomes centered.
func (self *PhysicsP2) EnableI(args ...interface{}) {
    self.Object.Call("enable", args)
}

// EnableBody Creates a P2 Physics body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
func (self *PhysicsP2) EnableBody(object interface{}, debug bool) {
    self.Object.Call("enableBody", object, debug)
}

// EnableBodyI Creates a P2 Physics body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
func (self *PhysicsP2) EnableBodyI(args ...interface{}) {
    self.Object.Call("enableBody", args)
}

// SetImpactEvents Impact event handling is disabled by default. Enable it before any impact events will be dispatched.
// In a busy world hundreds of impact events can be generated every step, so only enable this if you cannot do what you need via beginContact or collision masks.
func (self *PhysicsP2) SetImpactEvents(state bool) {
    self.Object.Call("setImpactEvents", state)
}

// SetImpactEventsI Impact event handling is disabled by default. Enable it before any impact events will be dispatched.
// In a busy world hundreds of impact events can be generated every step, so only enable this if you cannot do what you need via beginContact or collision masks.
func (self *PhysicsP2) SetImpactEventsI(args ...interface{}) {
    self.Object.Call("setImpactEvents", args)
}

// SetPostBroadphaseCallback Sets a callback to be fired after the Broadphase has collected collision pairs in the world.
// Just because a pair exists it doesn't mean they *will* collide, just that they potentially could do.
// If your calback returns `false` the pair will be removed from the narrowphase. This will stop them testing for collision this step.
// Returning `true` from the callback will ensure they are checked in the narrowphase.
func (self *PhysicsP2) SetPostBroadphaseCallback(callback interface{}, context interface{}) {
    self.Object.Call("setPostBroadphaseCallback", callback, context)
}

// SetPostBroadphaseCallbackI Sets a callback to be fired after the Broadphase has collected collision pairs in the world.
// Just because a pair exists it doesn't mean they *will* collide, just that they potentially could do.
// If your calback returns `false` the pair will be removed from the narrowphase. This will stop them testing for collision this step.
// Returning `true` from the callback will ensure they are checked in the narrowphase.
func (self *PhysicsP2) SetPostBroadphaseCallbackI(args ...interface{}) {
    self.Object.Call("setPostBroadphaseCallback", args)
}

// PostBroadphaseHandler Internal handler for the postBroadphase event.
func (self *PhysicsP2) PostBroadphaseHandler(event interface{}) {
    self.Object.Call("postBroadphaseHandler", event)
}

// PostBroadphaseHandlerI Internal handler for the postBroadphase event.
func (self *PhysicsP2) PostBroadphaseHandlerI(args ...interface{}) {
    self.Object.Call("postBroadphaseHandler", args)
}

// ImpactHandler Handles a p2 impact event.
func (self *PhysicsP2) ImpactHandler(event interface{}) {
    self.Object.Call("impactHandler", event)
}

// ImpactHandlerI Handles a p2 impact event.
func (self *PhysicsP2) ImpactHandlerI(args ...interface{}) {
    self.Object.Call("impactHandler", args)
}

// BeginContactHandler Handles a p2 begin contact event.
func (self *PhysicsP2) BeginContactHandler(event interface{}) {
    self.Object.Call("beginContactHandler", event)
}

// BeginContactHandlerI Handles a p2 begin contact event.
func (self *PhysicsP2) BeginContactHandlerI(args ...interface{}) {
    self.Object.Call("beginContactHandler", args)
}

// EndContactHandler Handles a p2 end contact event.
func (self *PhysicsP2) EndContactHandler(event interface{}) {
    self.Object.Call("endContactHandler", event)
}

// EndContactHandlerI Handles a p2 end contact event.
func (self *PhysicsP2) EndContactHandlerI(args ...interface{}) {
    self.Object.Call("endContactHandler", args)
}

// UpdateBoundsCollisionGroup By default the World will be set to collide everything with everything. The bounds of the world is a Body with 4 shapes, one for each face.
// If you start to use your own collision groups then your objects will no longer collide with the bounds.
// To fix this you need to adjust the bounds to use its own collision group first BEFORE changing your Sprites collision group.
func (self *PhysicsP2) UpdateBoundsCollisionGroup() {
    self.Object.Call("updateBoundsCollisionGroup")
}

// UpdateBoundsCollisionGroup1O By default the World will be set to collide everything with everything. The bounds of the world is a Body with 4 shapes, one for each face.
// If you start to use your own collision groups then your objects will no longer collide with the bounds.
// To fix this you need to adjust the bounds to use its own collision group first BEFORE changing your Sprites collision group.
func (self *PhysicsP2) UpdateBoundsCollisionGroup1O(setCollisionGroup bool) {
    self.Object.Call("updateBoundsCollisionGroup", setCollisionGroup)
}

// UpdateBoundsCollisionGroupI By default the World will be set to collide everything with everything. The bounds of the world is a Body with 4 shapes, one for each face.
// If you start to use your own collision groups then your objects will no longer collide with the bounds.
// To fix this you need to adjust the bounds to use its own collision group first BEFORE changing your Sprites collision group.
func (self *PhysicsP2) UpdateBoundsCollisionGroupI(args ...interface{}) {
    self.Object.Call("updateBoundsCollisionGroup", args)
}

// SetBounds Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds(x int, y int, width int, height int) {
    self.Object.Call("setBounds", x, y, width, height)
}

// SetBounds1O Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds1O(x int, y int, width int, height int, left bool) {
    self.Object.Call("setBounds", x, y, width, height, left)
}

// SetBounds2O Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds2O(x int, y int, width int, height int, left bool, right bool) {
    self.Object.Call("setBounds", x, y, width, height, left, right)
}

// SetBounds3O Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds3O(x int, y int, width int, height int, left bool, right bool, top bool) {
    self.Object.Call("setBounds", x, y, width, height, left, right, top)
}

// SetBounds4O Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds4O(x int, y int, width int, height int, left bool, right bool, top bool, bottom bool) {
    self.Object.Call("setBounds", x, y, width, height, left, right, top, bottom)
}

// SetBounds5O Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds5O(x int, y int, width int, height int, left bool, right bool, top bool, bottom bool, setCollisionGroup bool) {
    self.Object.Call("setBounds", x, y, width, height, left, right, top, bottom, setCollisionGroup)
}

// SetBoundsI Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBoundsI(args ...interface{}) {
    self.Object.Call("setBounds", args)
}

// SetupWall Internal method called by setBounds. Responsible for creating, updating or 
// removing the wall body shapes.
func (self *PhysicsP2) SetupWall(create bool, wall string, x int, y int, angle float64) {
    self.Object.Call("setupWall", create, wall, x, y, angle)
}

// SetupWall1O Internal method called by setBounds. Responsible for creating, updating or 
// removing the wall body shapes.
func (self *PhysicsP2) SetupWall1O(create bool, wall string, x int, y int, angle float64, setCollisionGroup bool) {
    self.Object.Call("setupWall", create, wall, x, y, angle, setCollisionGroup)
}

// SetupWallI Internal method called by setBounds. Responsible for creating, updating or 
// removing the wall body shapes.
func (self *PhysicsP2) SetupWallI(args ...interface{}) {
    self.Object.Call("setupWall", args)
}

// Pause Pauses the P2 World independent of the game pause state.
func (self *PhysicsP2) Pause() {
    self.Object.Call("pause")
}

// PauseI Pauses the P2 World independent of the game pause state.
func (self *PhysicsP2) PauseI(args ...interface{}) {
    self.Object.Call("pause", args)
}

// Resume Resumes a paused P2 World.
func (self *PhysicsP2) Resume() {
    self.Object.Call("resume")
}

// ResumeI Resumes a paused P2 World.
func (self *PhysicsP2) ResumeI(args ...interface{}) {
    self.Object.Call("resume", args)
}

// Update Internal P2 update loop.
func (self *PhysicsP2) Update() {
    self.Object.Call("update")
}

// UpdateI Internal P2 update loop.
func (self *PhysicsP2) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Reset Called by Phaser.Physics when a State swap occurs.
// Starts the begin and end Contact listeners again.
func (self *PhysicsP2) Reset() {
    self.Object.Call("reset")
}

// ResetI Called by Phaser.Physics when a State swap occurs.
// Starts the begin and end Contact listeners again.
func (self *PhysicsP2) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Clear Clears all bodies from the simulation, resets callbacks and resets the collision bitmask.
// 
// The P2 world is also cleared:
// 
// * Removes all solver equations
// * Removes all constraints
// * Removes all bodies
// * Removes all springs
// * Removes all contact materials
// 
// This is called automatically when you switch state.
func (self *PhysicsP2) Clear() {
    self.Object.Call("clear")
}

// ClearI Clears all bodies from the simulation, resets callbacks and resets the collision bitmask.
// 
// The P2 world is also cleared:
// 
// * Removes all solver equations
// * Removes all constraints
// * Removes all bodies
// * Removes all springs
// * Removes all contact materials
// 
// This is called automatically when you switch state.
func (self *PhysicsP2) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// Destroy Clears all bodies from the simulation and unlinks World from Game. Should only be called on game shutdown. Call `clear` on a State change.
func (self *PhysicsP2) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Clears all bodies from the simulation and unlinks World from Game. Should only be called on game shutdown. Call `clear` on a State change.
func (self *PhysicsP2) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// AddBody Add a body to the world.
func (self *PhysicsP2) AddBody(body *PhysicsP2Body) bool{
    return self.Object.Call("addBody", body).Bool()
}

// AddBodyI Add a body to the world.
func (self *PhysicsP2) AddBodyI(args ...interface{}) bool{
    return self.Object.Call("addBody", args).Bool()
}

// RemoveBody Removes a body from the world. This will silently fail if the body wasn't part of the world to begin with.
func (self *PhysicsP2) RemoveBody(body *PhysicsP2Body) *PhysicsP2Body{
    return &PhysicsP2Body{self.Object.Call("removeBody", body)}
}

// RemoveBodyI Removes a body from the world. This will silently fail if the body wasn't part of the world to begin with.
func (self *PhysicsP2) RemoveBodyI(args ...interface{}) *PhysicsP2Body{
    return &PhysicsP2Body{self.Object.Call("removeBody", args)}
}

// AddSpring Adds a Spring to the world.
func (self *PhysicsP2) AddSpring(spring interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("addSpring", spring)}
}

// AddSpringI Adds a Spring to the world.
func (self *PhysicsP2) AddSpringI(args ...interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("addSpring", args)}
}

// RemoveSpring Removes a Spring from the world.
func (self *PhysicsP2) RemoveSpring(spring *PhysicsP2Spring) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("removeSpring", spring)}
}

// RemoveSpringI Removes a Spring from the world.
func (self *PhysicsP2) RemoveSpringI(args ...interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("removeSpring", args)}
}

// CreateDistanceConstraint Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraint(bodyA interface{}, bodyB interface{}, distance int) *PhysicsP2DistanceConstraint{
    return &PhysicsP2DistanceConstraint{self.Object.Call("createDistanceConstraint", bodyA, bodyB, distance)}
}

// CreateDistanceConstraint1O Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraint1O(bodyA interface{}, bodyB interface{}, distance int, localAnchorA []interface{}) *PhysicsP2DistanceConstraint{
    return &PhysicsP2DistanceConstraint{self.Object.Call("createDistanceConstraint", bodyA, bodyB, distance, localAnchorA)}
}

// CreateDistanceConstraint2O Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraint2O(bodyA interface{}, bodyB interface{}, distance int, localAnchorA []interface{}, localAnchorB []interface{}) *PhysicsP2DistanceConstraint{
    return &PhysicsP2DistanceConstraint{self.Object.Call("createDistanceConstraint", bodyA, bodyB, distance, localAnchorA, localAnchorB)}
}

// CreateDistanceConstraint3O Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraint3O(bodyA interface{}, bodyB interface{}, distance int, localAnchorA []interface{}, localAnchorB []interface{}, maxForce int) *PhysicsP2DistanceConstraint{
    return &PhysicsP2DistanceConstraint{self.Object.Call("createDistanceConstraint", bodyA, bodyB, distance, localAnchorA, localAnchorB, maxForce)}
}

// CreateDistanceConstraintI Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraintI(args ...interface{}) *PhysicsP2DistanceConstraint{
    return &PhysicsP2DistanceConstraint{self.Object.Call("createDistanceConstraint", args)}
}

// CreateGearConstraint Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateGearConstraint(bodyA interface{}, bodyB interface{}) *PhysicsP2GearConstraint{
    return &PhysicsP2GearConstraint{self.Object.Call("createGearConstraint", bodyA, bodyB)}
}

// CreateGearConstraint1O Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateGearConstraint1O(bodyA interface{}, bodyB interface{}, angle int) *PhysicsP2GearConstraint{
    return &PhysicsP2GearConstraint{self.Object.Call("createGearConstraint", bodyA, bodyB, angle)}
}

// CreateGearConstraint2O Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateGearConstraint2O(bodyA interface{}, bodyB interface{}, angle int, ratio int) *PhysicsP2GearConstraint{
    return &PhysicsP2GearConstraint{self.Object.Call("createGearConstraint", bodyA, bodyB, angle, ratio)}
}

// CreateGearConstraintI Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateGearConstraintI(args ...interface{}) *PhysicsP2GearConstraint{
    return &PhysicsP2GearConstraint{self.Object.Call("createGearConstraint", args)}
}

// CreateRevoluteConstraint Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func (self *PhysicsP2) CreateRevoluteConstraint(bodyA interface{}, pivotA []interface{}, bodyB interface{}, pivotB []interface{}) *PhysicsP2RevoluteConstraint{
    return &PhysicsP2RevoluteConstraint{self.Object.Call("createRevoluteConstraint", bodyA, pivotA, bodyB, pivotB)}
}

// CreateRevoluteConstraint1O Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func (self *PhysicsP2) CreateRevoluteConstraint1O(bodyA interface{}, pivotA []interface{}, bodyB interface{}, pivotB []interface{}, maxForce int) *PhysicsP2RevoluteConstraint{
    return &PhysicsP2RevoluteConstraint{self.Object.Call("createRevoluteConstraint", bodyA, pivotA, bodyB, pivotB, maxForce)}
}

// CreateRevoluteConstraint2O Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func (self *PhysicsP2) CreateRevoluteConstraint2O(bodyA interface{}, pivotA []interface{}, bodyB interface{}, pivotB []interface{}, maxForce int, worldPivot *Float32Array) *PhysicsP2RevoluteConstraint{
    return &PhysicsP2RevoluteConstraint{self.Object.Call("createRevoluteConstraint", bodyA, pivotA, bodyB, pivotB, maxForce, worldPivot)}
}

// CreateRevoluteConstraintI Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func (self *PhysicsP2) CreateRevoluteConstraintI(args ...interface{}) *PhysicsP2RevoluteConstraint{
    return &PhysicsP2RevoluteConstraint{self.Object.Call("createRevoluteConstraint", args)}
}

// CreateLockConstraint Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraint(bodyA interface{}, bodyB interface{}) *PhysicsP2LockConstraint{
    return &PhysicsP2LockConstraint{self.Object.Call("createLockConstraint", bodyA, bodyB)}
}

// CreateLockConstraint1O Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraint1O(bodyA interface{}, bodyB interface{}, offset []interface{}) *PhysicsP2LockConstraint{
    return &PhysicsP2LockConstraint{self.Object.Call("createLockConstraint", bodyA, bodyB, offset)}
}

// CreateLockConstraint2O Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraint2O(bodyA interface{}, bodyB interface{}, offset []interface{}, angle int) *PhysicsP2LockConstraint{
    return &PhysicsP2LockConstraint{self.Object.Call("createLockConstraint", bodyA, bodyB, offset, angle)}
}

// CreateLockConstraint3O Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraint3O(bodyA interface{}, bodyB interface{}, offset []interface{}, angle int, maxForce int) *PhysicsP2LockConstraint{
    return &PhysicsP2LockConstraint{self.Object.Call("createLockConstraint", bodyA, bodyB, offset, angle, maxForce)}
}

// CreateLockConstraintI Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraintI(args ...interface{}) *PhysicsP2LockConstraint{
    return &PhysicsP2LockConstraint{self.Object.Call("createLockConstraint", args)}
}

// CreatePrismaticConstraint Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint(bodyA interface{}, bodyB interface{}) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB)}
}

// CreatePrismaticConstraint1O Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint1O(bodyA interface{}, bodyB interface{}, lockRotation bool) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB, lockRotation)}
}

// CreatePrismaticConstraint2O Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint2O(bodyA interface{}, bodyB interface{}, lockRotation bool, anchorA []interface{}) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB, lockRotation, anchorA)}
}

// CreatePrismaticConstraint3O Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint3O(bodyA interface{}, bodyB interface{}, lockRotation bool, anchorA []interface{}, anchorB []interface{}) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB, lockRotation, anchorA, anchorB)}
}

// CreatePrismaticConstraint4O Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint4O(bodyA interface{}, bodyB interface{}, lockRotation bool, anchorA []interface{}, anchorB []interface{}, axis []interface{}) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB, lockRotation, anchorA, anchorB, axis)}
}

// CreatePrismaticConstraint5O Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint5O(bodyA interface{}, bodyB interface{}, lockRotation bool, anchorA []interface{}, anchorB []interface{}, axis []interface{}, maxForce int) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB, lockRotation, anchorA, anchorB, axis, maxForce)}
}

// CreatePrismaticConstraintI Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraintI(args ...interface{}) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", args)}
}

// AddConstraint Adds a Constraint to the world.
func (self *PhysicsP2) AddConstraint(constraint *PhysicsP2Constraint) *PhysicsP2Constraint{
    return &PhysicsP2Constraint{self.Object.Call("addConstraint", constraint)}
}

// AddConstraintI Adds a Constraint to the world.
func (self *PhysicsP2) AddConstraintI(args ...interface{}) *PhysicsP2Constraint{
    return &PhysicsP2Constraint{self.Object.Call("addConstraint", args)}
}

// RemoveConstraint Removes a Constraint from the world.
func (self *PhysicsP2) RemoveConstraint(constraint *PhysicsP2Constraint) *PhysicsP2Constraint{
    return &PhysicsP2Constraint{self.Object.Call("removeConstraint", constraint)}
}

// RemoveConstraintI Removes a Constraint from the world.
func (self *PhysicsP2) RemoveConstraintI(args ...interface{}) *PhysicsP2Constraint{
    return &PhysicsP2Constraint{self.Object.Call("removeConstraint", args)}
}

// AddContactMaterial Adds a Contact Material to the world.
func (self *PhysicsP2) AddContactMaterial(material *PhysicsP2ContactMaterial) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("addContactMaterial", material)}
}

// AddContactMaterialI Adds a Contact Material to the world.
func (self *PhysicsP2) AddContactMaterialI(args ...interface{}) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("addContactMaterial", args)}
}

// RemoveContactMaterial Removes a Contact Material from the world.
func (self *PhysicsP2) RemoveContactMaterial(material *PhysicsP2ContactMaterial) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("removeContactMaterial", material)}
}

// RemoveContactMaterialI Removes a Contact Material from the world.
func (self *PhysicsP2) RemoveContactMaterialI(args ...interface{}) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("removeContactMaterial", args)}
}

// GetContactMaterial Gets a Contact Material based on the two given Materials.
func (self *PhysicsP2) GetContactMaterial(materialA *PhysicsP2Material, materialB *PhysicsP2Material) interface{}{
    return self.Object.Call("getContactMaterial", materialA, materialB)
}

// GetContactMaterialI Gets a Contact Material based on the two given Materials.
func (self *PhysicsP2) GetContactMaterialI(args ...interface{}) interface{}{
    return self.Object.Call("getContactMaterial", args)
}

// SetMaterial Sets the given Material against all Shapes owned by all the Bodies in the given array.
func (self *PhysicsP2) SetMaterial(material *PhysicsP2Material, bodies []PhysicsP2Body) {
    self.Object.Call("setMaterial", material, bodies)
}

// SetMaterialI Sets the given Material against all Shapes owned by all the Bodies in the given array.
func (self *PhysicsP2) SetMaterialI(args ...interface{}) {
    self.Object.Call("setMaterial", args)
}

// CreateMaterial Creates a Material. Materials are applied to Shapes owned by a Body and can be set with Body.setMaterial().
// Materials are a way to control what happens when Shapes collide. Combine unique Materials together to create Contact Materials.
// Contact Materials have properties such as friction and restitution that allow for fine-grained collision control between different Materials.
func (self *PhysicsP2) CreateMaterial() *PhysicsP2Material{
    return &PhysicsP2Material{self.Object.Call("createMaterial")}
}

// CreateMaterial1O Creates a Material. Materials are applied to Shapes owned by a Body and can be set with Body.setMaterial().
// Materials are a way to control what happens when Shapes collide. Combine unique Materials together to create Contact Materials.
// Contact Materials have properties such as friction and restitution that allow for fine-grained collision control between different Materials.
func (self *PhysicsP2) CreateMaterial1O(name string) *PhysicsP2Material{
    return &PhysicsP2Material{self.Object.Call("createMaterial", name)}
}

// CreateMaterial2O Creates a Material. Materials are applied to Shapes owned by a Body and can be set with Body.setMaterial().
// Materials are a way to control what happens when Shapes collide. Combine unique Materials together to create Contact Materials.
// Contact Materials have properties such as friction and restitution that allow for fine-grained collision control between different Materials.
func (self *PhysicsP2) CreateMaterial2O(name string, body *PhysicsP2Body) *PhysicsP2Material{
    return &PhysicsP2Material{self.Object.Call("createMaterial", name, body)}
}

// CreateMaterialI Creates a Material. Materials are applied to Shapes owned by a Body and can be set with Body.setMaterial().
// Materials are a way to control what happens when Shapes collide. Combine unique Materials together to create Contact Materials.
// Contact Materials have properties such as friction and restitution that allow for fine-grained collision control between different Materials.
func (self *PhysicsP2) CreateMaterialI(args ...interface{}) *PhysicsP2Material{
    return &PhysicsP2Material{self.Object.Call("createMaterial", args)}
}

// CreateContactMaterial Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterial() *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("createContactMaterial")}
}

// CreateContactMaterial1O Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterial1O(materialA *PhysicsP2Material) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("createContactMaterial", materialA)}
}

// CreateContactMaterial2O Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterial2O(materialA *PhysicsP2Material, materialB *PhysicsP2Material) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("createContactMaterial", materialA, materialB)}
}

// CreateContactMaterial3O Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterial3O(materialA *PhysicsP2Material, materialB *PhysicsP2Material, options interface{}) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("createContactMaterial", materialA, materialB, options)}
}

// CreateContactMaterialI Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterialI(args ...interface{}) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("createContactMaterial", args)}
}

// GetBodies Populates and returns an array with references to of all current Bodies in the world.
func (self *PhysicsP2) GetBodies() []PhysicsP2Body{
	array00 := self.Object.Call("getBodies")
	length00 := array00.Length()
	out00 := make([]PhysicsP2Body, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = PhysicsP2Body{array00.Index(i00)}
	}
	return out00
}

// GetBodiesI Populates and returns an array with references to of all current Bodies in the world.
func (self *PhysicsP2) GetBodiesI(args ...interface{}) []PhysicsP2Body{
	array00 := self.Object.Call("getBodies", args)
	length00 := array00.Length()
	out00 := make([]PhysicsP2Body, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = PhysicsP2Body{array00.Index(i00)}
	}
	return out00
}

// GetBody Checks the given object to see if it has a p2.Body and if so returns it.
func (self *PhysicsP2) GetBody(object interface{}) *P2Body{
    return &P2Body{self.Object.Call("getBody", object)}
}

// GetBodyI Checks the given object to see if it has a p2.Body and if so returns it.
func (self *PhysicsP2) GetBodyI(args ...interface{}) *P2Body{
    return &P2Body{self.Object.Call("getBody", args)}
}

// GetSprings Populates and returns an array of all current Springs in the world.
func (self *PhysicsP2) GetSprings() []PhysicsP2Spring{
	array00 := self.Object.Call("getSprings")
	length00 := array00.Length()
	out00 := make([]PhysicsP2Spring, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = PhysicsP2Spring{array00.Index(i00)}
	}
	return out00
}

// GetSpringsI Populates and returns an array of all current Springs in the world.
func (self *PhysicsP2) GetSpringsI(args ...interface{}) []PhysicsP2Spring{
	array00 := self.Object.Call("getSprings", args)
	length00 := array00.Length()
	out00 := make([]PhysicsP2Spring, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = PhysicsP2Spring{array00.Index(i00)}
	}
	return out00
}

// GetConstraints Populates and returns an array of all current Constraints in the world.
// You will get an array of p2 constraints back. This can be of mixed types, for example the array may contain
// PrismaticConstraints, RevoluteConstraints or any other valid p2 constraint type.
func (self *PhysicsP2) GetConstraints() []PhysicsP2Constraint{
	array00 := self.Object.Call("getConstraints")
	length00 := array00.Length()
	out00 := make([]PhysicsP2Constraint, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = PhysicsP2Constraint{array00.Index(i00)}
	}
	return out00
}

// GetConstraintsI Populates and returns an array of all current Constraints in the world.
// You will get an array of p2 constraints back. This can be of mixed types, for example the array may contain
// PrismaticConstraints, RevoluteConstraints or any other valid p2 constraint type.
func (self *PhysicsP2) GetConstraintsI(args ...interface{}) []PhysicsP2Constraint{
	array00 := self.Object.Call("getConstraints", args)
	length00 := array00.Length()
	out00 := make([]PhysicsP2Constraint, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = PhysicsP2Constraint{array00.Index(i00)}
	}
	return out00
}

// HitTest Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
// (if any) by checking the Body.parent.sprite property. Body.parent is a Phaser.Physics.P2.Body property.
func (self *PhysicsP2) HitTest(worldPoint *Point) []interface{}{
	array00 := self.Object.Call("hitTest", worldPoint)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// HitTest1O Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
// (if any) by checking the Body.parent.sprite property. Body.parent is a Phaser.Physics.P2.Body property.
func (self *PhysicsP2) HitTest1O(worldPoint *Point, bodies []interface{}) []interface{}{
	array00 := self.Object.Call("hitTest", worldPoint, bodies)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// HitTest2O Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
// (if any) by checking the Body.parent.sprite property. Body.parent is a Phaser.Physics.P2.Body property.
func (self *PhysicsP2) HitTest2O(worldPoint *Point, bodies []interface{}, precision int) []interface{}{
	array00 := self.Object.Call("hitTest", worldPoint, bodies, precision)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// HitTest3O Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
// (if any) by checking the Body.parent.sprite property. Body.parent is a Phaser.Physics.P2.Body property.
func (self *PhysicsP2) HitTest3O(worldPoint *Point, bodies []interface{}, precision int, filterStatic bool) []interface{}{
	array00 := self.Object.Call("hitTest", worldPoint, bodies, precision, filterStatic)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// HitTestI Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
// (if any) by checking the Body.parent.sprite property. Body.parent is a Phaser.Physics.P2.Body property.
func (self *PhysicsP2) HitTestI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("hitTest", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ToJSON Converts the current world into a JSON object.
func (self *PhysicsP2) ToJSON() interface{}{
    return self.Object.Call("toJSON")
}

// ToJSONI Converts the current world into a JSON object.
func (self *PhysicsP2) ToJSONI(args ...interface{}) interface{}{
    return self.Object.Call("toJSON", args)
}

// CreateCollisionGroup Creates a new Collision Group and optionally applies it to the given object.
// Collision Groups are handled using bitmasks, therefore you have a fixed limit you can create before you need to re-use older groups.
func (self *PhysicsP2) CreateCollisionGroup() {
    self.Object.Call("createCollisionGroup")
}

// CreateCollisionGroup1O Creates a new Collision Group and optionally applies it to the given object.
// Collision Groups are handled using bitmasks, therefore you have a fixed limit you can create before you need to re-use older groups.
func (self *PhysicsP2) CreateCollisionGroup1O(object interface{}) {
    self.Object.Call("createCollisionGroup", object)
}

// CreateCollisionGroupI Creates a new Collision Group and optionally applies it to the given object.
// Collision Groups are handled using bitmasks, therefore you have a fixed limit you can create before you need to re-use older groups.
func (self *PhysicsP2) CreateCollisionGroupI(args ...interface{}) {
    self.Object.Call("createCollisionGroup", args)
}

// CreateSpring Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring(bodyA interface{}, bodyB interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB)}
}

// CreateSpring1O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring1O(bodyA interface{}, bodyB interface{}, restLength int) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength)}
}

// CreateSpring2O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring2O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness)}
}

// CreateSpring3O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring3O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int, damping int) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness, damping)}
}

// CreateSpring4O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring4O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int, damping int, worldA []interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness, damping, worldA)}
}

// CreateSpring5O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring5O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int, damping int, worldA []interface{}, worldB []interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness, damping, worldA, worldB)}
}

// CreateSpring6O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring6O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int, damping int, worldA []interface{}, worldB []interface{}, localA []interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness, damping, worldA, worldB, localA)}
}

// CreateSpring7O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring7O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int, damping int, worldA []interface{}, worldB []interface{}, localA []interface{}, localB []interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness, damping, worldA, worldB, localA, localB)}
}

// CreateSpringI Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpringI(args ...interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", args)}
}

// CreateRotationalSpring Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpring(bodyA interface{}, bodyB interface{}) *PhysicsP2RotationalSpring{
    return &PhysicsP2RotationalSpring{self.Object.Call("createRotationalSpring", bodyA, bodyB)}
}

// CreateRotationalSpring1O Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpring1O(bodyA interface{}, bodyB interface{}, restAngle int) *PhysicsP2RotationalSpring{
    return &PhysicsP2RotationalSpring{self.Object.Call("createRotationalSpring", bodyA, bodyB, restAngle)}
}

// CreateRotationalSpring2O Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpring2O(bodyA interface{}, bodyB interface{}, restAngle int, stiffness int) *PhysicsP2RotationalSpring{
    return &PhysicsP2RotationalSpring{self.Object.Call("createRotationalSpring", bodyA, bodyB, restAngle, stiffness)}
}

// CreateRotationalSpring3O Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpring3O(bodyA interface{}, bodyB interface{}, restAngle int, stiffness int, damping int) *PhysicsP2RotationalSpring{
    return &PhysicsP2RotationalSpring{self.Object.Call("createRotationalSpring", bodyA, bodyB, restAngle, stiffness, damping)}
}

// CreateRotationalSpringI Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpringI(args ...interface{}) *PhysicsP2RotationalSpring{
    return &PhysicsP2RotationalSpring{self.Object.Call("createRotationalSpring", args)}
}

// CreateBody Creates a new Body and adds it to the World.
func (self *PhysicsP2) CreateBody(x int, y int, mass int, addToWorld bool, options interface{}, points interface{}) *PhysicsP2Body{
    return &PhysicsP2Body{self.Object.Call("createBody", x, y, mass, addToWorld, options, points)}
}

// CreateBodyI Creates a new Body and adds it to the World.
func (self *PhysicsP2) CreateBodyI(args ...interface{}) *PhysicsP2Body{
    return &PhysicsP2Body{self.Object.Call("createBody", args)}
}

// CreateParticle Creates a new Particle and adds it to the World.
func (self *PhysicsP2) CreateParticle(x int, y int, mass int, addToWorld bool, options interface{}, points interface{}) {
    self.Object.Call("createParticle", x, y, mass, addToWorld, options, points)
}

// CreateParticleI Creates a new Particle and adds it to the World.
func (self *PhysicsP2) CreateParticleI(args ...interface{}) {
    self.Object.Call("createParticle", args)
}

// ConvertCollisionObjects Converts all of the polylines objects inside a Tiled ObjectGroup into physics bodies that are added to the world.
// Note that the polylines must be created in such a way that they can withstand polygon decomposition.
func (self *PhysicsP2) ConvertCollisionObjects(map_ *Tilemap) []interface{}{
	array00 := self.Object.Call("convertCollisionObjects", map_)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ConvertCollisionObjects1O Converts all of the polylines objects inside a Tiled ObjectGroup into physics bodies that are added to the world.
// Note that the polylines must be created in such a way that they can withstand polygon decomposition.
func (self *PhysicsP2) ConvertCollisionObjects1O(map_ *Tilemap, layer interface{}) []interface{}{
	array00 := self.Object.Call("convertCollisionObjects", map_, layer)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ConvertCollisionObjects2O Converts all of the polylines objects inside a Tiled ObjectGroup into physics bodies that are added to the world.
// Note that the polylines must be created in such a way that they can withstand polygon decomposition.
func (self *PhysicsP2) ConvertCollisionObjects2O(map_ *Tilemap, layer interface{}, addToWorld bool) []interface{}{
	array00 := self.Object.Call("convertCollisionObjects", map_, layer, addToWorld)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ConvertCollisionObjectsI Converts all of the polylines objects inside a Tiled ObjectGroup into physics bodies that are added to the world.
// Note that the polylines must be created in such a way that they can withstand polygon decomposition.
func (self *PhysicsP2) ConvertCollisionObjectsI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("convertCollisionObjects", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ClearTilemapLayerBodies Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsP2) ClearTilemapLayerBodies(map_ *Tilemap) {
    self.Object.Call("clearTilemapLayerBodies", map_)
}

// ClearTilemapLayerBodies1O Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsP2) ClearTilemapLayerBodies1O(map_ *Tilemap, layer interface{}) {
    self.Object.Call("clearTilemapLayerBodies", map_, layer)
}

// ClearTilemapLayerBodiesI Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsP2) ClearTilemapLayerBodiesI(args ...interface{}) {
    self.Object.Call("clearTilemapLayerBodies", args)
}

// ConvertTilemap Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
// Only call this *after* you have specified all of the tiles you wish to collide with calls like Tilemap.setCollisionBetween, etc.
// Every time you call this method it will destroy any previously created bodies and remove them from the world.
// Therefore understand it's a very expensive operation and not to be done in a core game update loop.
func (self *PhysicsP2) ConvertTilemap(map_ *Tilemap) []interface{}{
	array00 := self.Object.Call("convertTilemap", map_)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ConvertTilemap1O Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
// Only call this *after* you have specified all of the tiles you wish to collide with calls like Tilemap.setCollisionBetween, etc.
// Every time you call this method it will destroy any previously created bodies and remove them from the world.
// Therefore understand it's a very expensive operation and not to be done in a core game update loop.
func (self *PhysicsP2) ConvertTilemap1O(map_ *Tilemap, layer interface{}) []interface{}{
	array00 := self.Object.Call("convertTilemap", map_, layer)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ConvertTilemap2O Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
// Only call this *after* you have specified all of the tiles you wish to collide with calls like Tilemap.setCollisionBetween, etc.
// Every time you call this method it will destroy any previously created bodies and remove them from the world.
// Therefore understand it's a very expensive operation and not to be done in a core game update loop.
func (self *PhysicsP2) ConvertTilemap2O(map_ *Tilemap, layer interface{}, addToWorld bool) []interface{}{
	array00 := self.Object.Call("convertTilemap", map_, layer, addToWorld)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ConvertTilemap3O Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
// Only call this *after* you have specified all of the tiles you wish to collide with calls like Tilemap.setCollisionBetween, etc.
// Every time you call this method it will destroy any previously created bodies and remove them from the world.
// Therefore understand it's a very expensive operation and not to be done in a core game update loop.
func (self *PhysicsP2) ConvertTilemap3O(map_ *Tilemap, layer interface{}, addToWorld bool, optimize bool) []interface{}{
	array00 := self.Object.Call("convertTilemap", map_, layer, addToWorld, optimize)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ConvertTilemapI Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
// Only call this *after* you have specified all of the tiles you wish to collide with calls like Tilemap.setCollisionBetween, etc.
// Every time you call this method it will destroy any previously created bodies and remove them from the world.
// Therefore understand it's a very expensive operation and not to be done in a core game update loop.
func (self *PhysicsP2) ConvertTilemapI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("convertTilemap", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Mpx Convert p2 physics value (meters) to pixel scale.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) Mpx(v int) int{
    return self.Object.Call("mpx", v).Int()
}

// MpxI Convert p2 physics value (meters) to pixel scale.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) MpxI(args ...interface{}) int{
    return self.Object.Call("mpx", args).Int()
}

// Pxm Convert pixel value to p2 physics scale (meters).
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) Pxm(v int) int{
    return self.Object.Call("pxm", v).Int()
}

// PxmI Convert pixel value to p2 physics scale (meters).
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) PxmI(args ...interface{}) int{
    return self.Object.Call("pxm", args).Int()
}

// Mpxi Convert p2 physics value (meters) to pixel scale and inverses it.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) Mpxi(v int) int{
    return self.Object.Call("mpxi", v).Int()
}

// MpxiI Convert p2 physics value (meters) to pixel scale and inverses it.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) MpxiI(args ...interface{}) int{
    return self.Object.Call("mpxi", args).Int()
}

// Pxmi Convert pixel value to p2 physics scale (meters) and inverses it.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) Pxmi(v int) int{
    return self.Object.Call("pxmi", v).Int()
}

// PxmiI Convert pixel value to p2 physics scale (meters) and inverses it.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) PxmiI(args ...interface{}) int{
    return self.Object.Call("pxmi", args).Int()
}

