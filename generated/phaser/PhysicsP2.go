// Automatic generation for Phaser.Physics.P2
// generated file PhysicsP2.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// This is your main access to the P2 Physics World.
// From here you can create materials, listen for events and add bodies into the physics simulation.
type PhysicsP2 struct {
    *js.Object
}


// This is your main access to the P2 Physics World.
// From here you can create materials, listen for events and add bodies into the physics simulation.
func NewPhysicsP2(game *Game) *PhysicsP2 {
    return &PhysicsP2{js.Global.Get("Phaser").Get("Physics").Get("P2").New(game)}
}

// This is your main access to the P2 Physics World.
// From here you can create materials, listen for events and add bodies into the physics simulation.
func NewPhysicsP21O(game *Game, config interface{}) *PhysicsP2 {
    return &PhysicsP2{js.Global.Get("Phaser").Get("Physics").Get("P2").New(game, config)}
}

// This is your main access to the P2 Physics World.
// From here you can create materials, listen for events and add bodies into the physics simulation.
func NewPhysicsP2I(args ...interface{}) *PhysicsP2 {
    return &PhysicsP2{js.Global.Get("Phaser").Get("Physics").Get("P2").New(args)}
}



// Local reference to game.
func (self *PhysicsP2) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The p2 World configuration object.
func (self *PhysicsP2) Config() interface{}{
    return self.Object.Get("config")
}

// The p2 World configuration object.
func (self *PhysicsP2) SetConfigA(member interface{}) {
    self.Object.Set("config", member)
}

// The p2 World in which the simulation is run.
func (self *PhysicsP2) World() *P2World{
    return &P2World{self.Object.Get("world")}
}

// The p2 World in which the simulation is run.
func (self *PhysicsP2) SetWorldA(member *P2World) {
    self.Object.Set("world", member)
}

// The frame rate the world will be stepped at. Defaults to 1 / 60, but you can change here. Also see useElapsedTime property.
func (self *PhysicsP2) FrameRate() int{
    return self.Object.Get("frameRate").Int()
}

// The frame rate the world will be stepped at. Defaults to 1 / 60, but you can change here. Also see useElapsedTime property.
func (self *PhysicsP2) SetFrameRateA(member int) {
    self.Object.Set("frameRate", member)
}

// If true the frameRate value will be ignored and instead p2 will step with the value of Game.Time.physicsElapsed, which is a delta time value.
func (self *PhysicsP2) UseElapsedTime() bool{
    return self.Object.Get("useElapsedTime").Bool()
}

// If true the frameRate value will be ignored and instead p2 will step with the value of Game.Time.physicsElapsed, which is a delta time value.
func (self *PhysicsP2) SetUseElapsedTimeA(member bool) {
    self.Object.Set("useElapsedTime", member)
}

// The paused state of the P2 World.
func (self *PhysicsP2) Paused() bool{
    return self.Object.Get("paused").Bool()
}

// The paused state of the P2 World.
func (self *PhysicsP2) SetPausedA(member bool) {
    self.Object.Set("paused", member)
}

// A local array of all created Materials.
func (self *PhysicsP2) Materials() []PhysicsP2Material{
	array00 := self.Object.Get("materials")
	length00 := array00.Length()
	out00 := make([]PhysicsP2Material, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = PhysicsP2Material{array00.Index(i00)}
	}
	return out00
}

// A local array of all created Materials.
func (self *PhysicsP2) SetMaterialsA(member []PhysicsP2Material) {
    self.Object.Set("materials", member)
}

// The gravity applied to all bodies each step.
func (self *PhysicsP2) Gravity() *PhysicsP2InversePointProxy{
    return &PhysicsP2InversePointProxy{self.Object.Get("gravity")}
}

// The gravity applied to all bodies each step.
func (self *PhysicsP2) SetGravityA(member *PhysicsP2InversePointProxy) {
    self.Object.Set("gravity", member)
}

// An object containing the 4 wall bodies that bound the physics world.
func (self *PhysicsP2) Walls() interface{}{
    return self.Object.Get("walls")
}

// An object containing the 4 wall bodies that bound the physics world.
func (self *PhysicsP2) SetWallsA(member interface{}) {
    self.Object.Set("walls", member)
}

// This signal is dispatched when a new Body is added to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was added to the world.
func (self *PhysicsP2) OnBodyAdded() *Signal{
    return &Signal{self.Object.Get("onBodyAdded")}
}

// This signal is dispatched when a new Body is added to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was added to the world.
func (self *PhysicsP2) SetOnBodyAddedA(member *Signal) {
    self.Object.Set("onBodyAdded", member)
}

// This signal is dispatched when a Body is removed to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was removed from the world.
func (self *PhysicsP2) OnBodyRemoved() *Signal{
    return &Signal{self.Object.Get("onBodyRemoved")}
}

// This signal is dispatched when a Body is removed to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was removed from the world.
func (self *PhysicsP2) SetOnBodyRemovedA(member *Signal) {
    self.Object.Set("onBodyRemoved", member)
}

// This signal is dispatched when a Spring is added to the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was added to the world.
func (self *PhysicsP2) OnSpringAdded() *Signal{
    return &Signal{self.Object.Get("onSpringAdded")}
}

// This signal is dispatched when a Spring is added to the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was added to the world.
func (self *PhysicsP2) SetOnSpringAddedA(member *Signal) {
    self.Object.Set("onSpringAdded", member)
}

// This signal is dispatched when a Spring is removed from the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was removed from the world.
func (self *PhysicsP2) OnSpringRemoved() *Signal{
    return &Signal{self.Object.Get("onSpringRemoved")}
}

// This signal is dispatched when a Spring is removed from the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was removed from the world.
func (self *PhysicsP2) SetOnSpringRemovedA(member *Signal) {
    self.Object.Set("onSpringRemoved", member)
}

// This signal is dispatched when a Constraint is added to the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was added to the world.
func (self *PhysicsP2) OnConstraintAdded() *Signal{
    return &Signal{self.Object.Get("onConstraintAdded")}
}

// This signal is dispatched when a Constraint is added to the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was added to the world.
func (self *PhysicsP2) SetOnConstraintAddedA(member *Signal) {
    self.Object.Set("onConstraintAdded", member)
}

// This signal is dispatched when a Constraint is removed from the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was removed from the world.
func (self *PhysicsP2) OnConstraintRemoved() *Signal{
    return &Signal{self.Object.Get("onConstraintRemoved")}
}

// This signal is dispatched when a Constraint is removed from the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was removed from the world.
func (self *PhysicsP2) SetOnConstraintRemovedA(member *Signal) {
    self.Object.Set("onConstraintRemoved", member)
}

// This signal is dispatched when a Contact Material is added to the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was added to the world.
func (self *PhysicsP2) OnContactMaterialAdded() *Signal{
    return &Signal{self.Object.Get("onContactMaterialAdded")}
}

// This signal is dispatched when a Contact Material is added to the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was added to the world.
func (self *PhysicsP2) SetOnContactMaterialAddedA(member *Signal) {
    self.Object.Set("onContactMaterialAdded", member)
}

// This signal is dispatched when a Contact Material is removed from the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was removed from the world.
func (self *PhysicsP2) OnContactMaterialRemoved() *Signal{
    return &Signal{self.Object.Get("onContactMaterialRemoved")}
}

// This signal is dispatched when a Contact Material is removed from the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was removed from the world.
func (self *PhysicsP2) SetOnContactMaterialRemovedA(member *Signal) {
    self.Object.Set("onContactMaterialRemoved", member)
}

// A postBroadphase callback.
func (self *PhysicsP2) PostBroadphaseCallback() interface{}{
    return self.Object.Get("postBroadphaseCallback")
}

// A postBroadphase callback.
func (self *PhysicsP2) SetPostBroadphaseCallbackA(member interface{}) {
    self.Object.Set("postBroadphaseCallback", member)
}

// The context under which the callbacks are fired.
func (self *PhysicsP2) CallbackContext() interface{}{
    return self.Object.Get("callbackContext")
}

// The context under which the callbacks are fired.
func (self *PhysicsP2) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// This Signal is dispatched when a first contact is created between two bodies. This happens *before* the step has been done.
// 
// It sends 5 arguments: `bodyA`, `bodyB`, `shapeA`, `shapeB` and `contactEquations`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) OnBeginContact() *Signal{
    return &Signal{self.Object.Get("onBeginContact")}
}

// This Signal is dispatched when a first contact is created between two bodies. This happens *before* the step has been done.
// 
// It sends 5 arguments: `bodyA`, `bodyB`, `shapeA`, `shapeB` and `contactEquations`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) SetOnBeginContactA(member *Signal) {
    self.Object.Set("onBeginContact", member)
}

// This Signal is dispatched when final contact occurs between two bodies. This happens *before* the step has been done.
// 
// It sends 4 arguments: `bodyA`, `bodyB`, `shapeA` and `shapeB`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) OnEndContact() *Signal{
    return &Signal{self.Object.Get("onEndContact")}
}

// This Signal is dispatched when final contact occurs between two bodies. This happens *before* the step has been done.
// 
// It sends 4 arguments: `bodyA`, `bodyB`, `shapeA` and `shapeB`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) SetOnEndContactA(member *Signal) {
    self.Object.Set("onEndContact", member)
}

// An array containing the collision groups that have been defined in the World.
func (self *PhysicsP2) CollisionGroups() []interface{}{
	array00 := self.Object.Get("collisionGroups")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// An array containing the collision groups that have been defined in the World.
func (self *PhysicsP2) SetCollisionGroupsA(member []interface{}) {
    self.Object.Set("collisionGroups", member)
}

// A default collision group.
func (self *PhysicsP2) NothingCollisionGroup() *PhysicsP2CollisionGroup{
    return &PhysicsP2CollisionGroup{self.Object.Get("nothingCollisionGroup")}
}

// A default collision group.
func (self *PhysicsP2) SetNothingCollisionGroupA(member *PhysicsP2CollisionGroup) {
    self.Object.Set("nothingCollisionGroup", member)
}

// A default collision group.
func (self *PhysicsP2) BoundsCollisionGroup() *PhysicsP2CollisionGroup{
    return &PhysicsP2CollisionGroup{self.Object.Get("boundsCollisionGroup")}
}

// A default collision group.
func (self *PhysicsP2) SetBoundsCollisionGroupA(member *PhysicsP2CollisionGroup) {
    self.Object.Set("boundsCollisionGroup", member)
}

// A default collision group.
func (self *PhysicsP2) EverythingCollisionGroup() *PhysicsP2CollisionGroup{
    return &PhysicsP2CollisionGroup{self.Object.Get("everythingCollisionGroup")}
}

// A default collision group.
func (self *PhysicsP2) SetEverythingCollisionGroupA(member *PhysicsP2CollisionGroup) {
    self.Object.Set("everythingCollisionGroup", member)
}

// An array of the bodies the world bounds collides with.
func (self *PhysicsP2) BoundsCollidesWith() []interface{}{
	array00 := self.Object.Get("boundsCollidesWith")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// An array of the bodies the world bounds collides with.
func (self *PhysicsP2) SetBoundsCollidesWithA(member []interface{}) {
    self.Object.Set("boundsCollidesWith", member)
}

// Friction between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) Friction() int{
    return self.Object.Get("friction").Int()
}

// Friction between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) SetFrictionA(member int) {
    self.Object.Set("friction", member)
}

// Default coefficient of restitution between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) Restitution() int{
    return self.Object.Get("restitution").Int()
}

// Default coefficient of restitution between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) SetRestitutionA(member int) {
    self.Object.Set("restitution", member)
}

// The default Contact Material being used by the World.
func (self *PhysicsP2) ContactMaterial() *P2ContactMaterial{
    return &P2ContactMaterial{self.Object.Get("contactMaterial")}
}

// The default Contact Material being used by the World.
func (self *PhysicsP2) SetContactMaterialA(member *P2ContactMaterial) {
    self.Object.Set("contactMaterial", member)
}

// Enable to automatically apply spring forces each step.
func (self *PhysicsP2) ApplySpringForces() bool{
    return self.Object.Get("applySpringForces").Bool()
}

// Enable to automatically apply spring forces each step.
func (self *PhysicsP2) SetApplySpringForcesA(member bool) {
    self.Object.Set("applySpringForces", member)
}

// Enable to automatically apply body damping each step.
func (self *PhysicsP2) ApplyDamping() bool{
    return self.Object.Get("applyDamping").Bool()
}

// Enable to automatically apply body damping each step.
func (self *PhysicsP2) SetApplyDampingA(member bool) {
    self.Object.Set("applyDamping", member)
}

// Enable to automatically apply gravity each step.
func (self *PhysicsP2) ApplyGravity() bool{
    return self.Object.Get("applyGravity").Bool()
}

// Enable to automatically apply gravity each step.
func (self *PhysicsP2) SetApplyGravityA(member bool) {
    self.Object.Set("applyGravity", member)
}

// Enable/disable constraint solving in each step.
func (self *PhysicsP2) SolveConstraints() bool{
    return self.Object.Get("solveConstraints").Bool()
}

// Enable/disable constraint solving in each step.
func (self *PhysicsP2) SetSolveConstraintsA(member bool) {
    self.Object.Set("solveConstraints", member)
}

// The World time.
func (self *PhysicsP2) Time() bool{
    return self.Object.Get("time").Bool()
}

// The World time.
func (self *PhysicsP2) SetTimeA(member bool) {
    self.Object.Set("time", member)
}

// Set to true if you want to the world to emit the "impact" event. Turning this off could improve performance.
func (self *PhysicsP2) EmitImpactEvent() bool{
    return self.Object.Get("emitImpactEvent").Bool()
}

// Set to true if you want to the world to emit the "impact" event. Turning this off could improve performance.
func (self *PhysicsP2) SetEmitImpactEventA(member bool) {
    self.Object.Set("emitImpactEvent", member)
}

// How to deactivate bodies during simulation. Possible modes are: World.NO_SLEEPING, World.BODY_SLEEPING and World.ISLAND_SLEEPING.
// If sleeping is enabled, you might need to wake up the bodies if they fall asleep when they shouldn't. If you want to enable sleeping in the world, but want to disable it for a particular body, see Body.allowSleep.
func (self *PhysicsP2) SleepMode() int{
    return self.Object.Get("sleepMode").Int()
}

// How to deactivate bodies during simulation. Possible modes are: World.NO_SLEEPING, World.BODY_SLEEPING and World.ISLAND_SLEEPING.
// If sleeping is enabled, you might need to wake up the bodies if they fall asleep when they shouldn't. If you want to enable sleeping in the world, but want to disable it for a particular body, see Body.allowSleep.
func (self *PhysicsP2) SetSleepModeA(member int) {
    self.Object.Set("sleepMode", member)
}

// The total number of bodies in the world.
func (self *PhysicsP2) Total() int{
    return self.Object.Get("total").Int()
}

// The total number of bodies in the world.
func (self *PhysicsP2) SetTotalA(member int) {
    self.Object.Set("total", member)
}



// This will add a P2 Physics body into the removal list for the next step.
func (self *PhysicsP2) RemoveBodyNextStep(body *PhysicsP2Body) {
    self.Object.Call("removeBodyNextStep", body)
}

// This will add a P2 Physics body into the removal list for the next step.
func (self *PhysicsP2) RemoveBodyNextStepI(args ...interface{}) {
    self.Object.Call("removeBodyNextStep", args)
}

// Called at the start of the core update loop. Purges flagged bodies from the world.
func (self *PhysicsP2) PreUpdate() {
    self.Object.Call("preUpdate")
}

// Called at the start of the core update loop. Purges flagged bodies from the world.
func (self *PhysicsP2) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// This will create a P2 Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// Note: When the game object is enabled for P2 physics it has its anchor x/y set to 0.5 so it becomes centered.
func (self *PhysicsP2) Enable(object interface{}) {
    self.Object.Call("enable", object)
}

// This will create a P2 Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// Note: When the game object is enabled for P2 physics it has its anchor x/y set to 0.5 so it becomes centered.
func (self *PhysicsP2) Enable1O(object interface{}, debug bool) {
    self.Object.Call("enable", object, debug)
}

// This will create a P2 Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// Note: When the game object is enabled for P2 physics it has its anchor x/y set to 0.5 so it becomes centered.
func (self *PhysicsP2) Enable2O(object interface{}, debug bool, children bool) {
    self.Object.Call("enable", object, debug, children)
}

// This will create a P2 Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// Note: When the game object is enabled for P2 physics it has its anchor x/y set to 0.5 so it becomes centered.
func (self *PhysicsP2) EnableI(args ...interface{}) {
    self.Object.Call("enable", args)
}

// Creates a P2 Physics body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
func (self *PhysicsP2) EnableBody(object interface{}, debug bool) {
    self.Object.Call("enableBody", object, debug)
}

// Creates a P2 Physics body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
func (self *PhysicsP2) EnableBodyI(args ...interface{}) {
    self.Object.Call("enableBody", args)
}

// Impact event handling is disabled by default. Enable it before any impact events will be dispatched.
// In a busy world hundreds of impact events can be generated every step, so only enable this if you cannot do what you need via beginContact or collision masks.
func (self *PhysicsP2) SetImpactEvents(state bool) {
    self.Object.Call("setImpactEvents", state)
}

// Impact event handling is disabled by default. Enable it before any impact events will be dispatched.
// In a busy world hundreds of impact events can be generated every step, so only enable this if you cannot do what you need via beginContact or collision masks.
func (self *PhysicsP2) SetImpactEventsI(args ...interface{}) {
    self.Object.Call("setImpactEvents", args)
}

// Sets a callback to be fired after the Broadphase has collected collision pairs in the world.
// Just because a pair exists it doesn't mean they *will* collide, just that they potentially could do.
// If your calback returns `false` the pair will be removed from the narrowphase. This will stop them testing for collision this step.
// Returning `true` from the callback will ensure they are checked in the narrowphase.
func (self *PhysicsP2) SetPostBroadphaseCallback(callback interface{}, context interface{}) {
    self.Object.Call("setPostBroadphaseCallback", callback, context)
}

// Sets a callback to be fired after the Broadphase has collected collision pairs in the world.
// Just because a pair exists it doesn't mean they *will* collide, just that they potentially could do.
// If your calback returns `false` the pair will be removed from the narrowphase. This will stop them testing for collision this step.
// Returning `true` from the callback will ensure they are checked in the narrowphase.
func (self *PhysicsP2) SetPostBroadphaseCallbackI(args ...interface{}) {
    self.Object.Call("setPostBroadphaseCallback", args)
}

// Internal handler for the postBroadphase event.
func (self *PhysicsP2) PostBroadphaseHandler(event interface{}) {
    self.Object.Call("postBroadphaseHandler", event)
}

// Internal handler for the postBroadphase event.
func (self *PhysicsP2) PostBroadphaseHandlerI(args ...interface{}) {
    self.Object.Call("postBroadphaseHandler", args)
}

// Handles a p2 impact event.
func (self *PhysicsP2) ImpactHandler(event interface{}) {
    self.Object.Call("impactHandler", event)
}

// Handles a p2 impact event.
func (self *PhysicsP2) ImpactHandlerI(args ...interface{}) {
    self.Object.Call("impactHandler", args)
}

// Handles a p2 begin contact event.
func (self *PhysicsP2) BeginContactHandler(event interface{}) {
    self.Object.Call("beginContactHandler", event)
}

// Handles a p2 begin contact event.
func (self *PhysicsP2) BeginContactHandlerI(args ...interface{}) {
    self.Object.Call("beginContactHandler", args)
}

// Handles a p2 end contact event.
func (self *PhysicsP2) EndContactHandler(event interface{}) {
    self.Object.Call("endContactHandler", event)
}

// Handles a p2 end contact event.
func (self *PhysicsP2) EndContactHandlerI(args ...interface{}) {
    self.Object.Call("endContactHandler", args)
}

// By default the World will be set to collide everything with everything. The bounds of the world is a Body with 4 shapes, one for each face.
// If you start to use your own collision groups then your objects will no longer collide with the bounds.
// To fix this you need to adjust the bounds to use its own collision group first BEFORE changing your Sprites collision group.
func (self *PhysicsP2) UpdateBoundsCollisionGroup() {
    self.Object.Call("updateBoundsCollisionGroup")
}

// By default the World will be set to collide everything with everything. The bounds of the world is a Body with 4 shapes, one for each face.
// If you start to use your own collision groups then your objects will no longer collide with the bounds.
// To fix this you need to adjust the bounds to use its own collision group first BEFORE changing your Sprites collision group.
func (self *PhysicsP2) UpdateBoundsCollisionGroup1O(setCollisionGroup bool) {
    self.Object.Call("updateBoundsCollisionGroup", setCollisionGroup)
}

// By default the World will be set to collide everything with everything. The bounds of the world is a Body with 4 shapes, one for each face.
// If you start to use your own collision groups then your objects will no longer collide with the bounds.
// To fix this you need to adjust the bounds to use its own collision group first BEFORE changing your Sprites collision group.
func (self *PhysicsP2) UpdateBoundsCollisionGroupI(args ...interface{}) {
    self.Object.Call("updateBoundsCollisionGroup", args)
}

// Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds(x int, y int, width int, height int) {
    self.Object.Call("setBounds", x, y, width, height)
}

// Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds1O(x int, y int, width int, height int, left bool) {
    self.Object.Call("setBounds", x, y, width, height, left)
}

// Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds2O(x int, y int, width int, height int, left bool, right bool) {
    self.Object.Call("setBounds", x, y, width, height, left, right)
}

// Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds3O(x int, y int, width int, height int, left bool, right bool, top bool) {
    self.Object.Call("setBounds", x, y, width, height, left, right, top)
}

// Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds4O(x int, y int, width int, height int, left bool, right bool, top bool, bottom bool) {
    self.Object.Call("setBounds", x, y, width, height, left, right, top, bottom)
}

// Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBounds5O(x int, y int, width int, height int, left bool, right bool, top bool, bottom bool, setCollisionGroup bool) {
    self.Object.Call("setBounds", x, y, width, height, left, right, top, bottom, setCollisionGroup)
}

// Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBoundsI(args ...interface{}) {
    self.Object.Call("setBounds", args)
}

// Internal method called by setBounds. Responsible for creating, updating or 
// removing the wall body shapes.
func (self *PhysicsP2) SetupWall(create bool, wall string, x int, y int, angle float64) {
    self.Object.Call("setupWall", create, wall, x, y, angle)
}

// Internal method called by setBounds. Responsible for creating, updating or 
// removing the wall body shapes.
func (self *PhysicsP2) SetupWall1O(create bool, wall string, x int, y int, angle float64, setCollisionGroup bool) {
    self.Object.Call("setupWall", create, wall, x, y, angle, setCollisionGroup)
}

// Internal method called by setBounds. Responsible for creating, updating or 
// removing the wall body shapes.
func (self *PhysicsP2) SetupWallI(args ...interface{}) {
    self.Object.Call("setupWall", args)
}

// Pauses the P2 World independent of the game pause state.
func (self *PhysicsP2) Pause() {
    self.Object.Call("pause")
}

// Pauses the P2 World independent of the game pause state.
func (self *PhysicsP2) PauseI(args ...interface{}) {
    self.Object.Call("pause", args)
}

// Resumes a paused P2 World.
func (self *PhysicsP2) Resume() {
    self.Object.Call("resume")
}

// Resumes a paused P2 World.
func (self *PhysicsP2) ResumeI(args ...interface{}) {
    self.Object.Call("resume", args)
}

// Internal P2 update loop.
func (self *PhysicsP2) Update() {
    self.Object.Call("update")
}

// Internal P2 update loop.
func (self *PhysicsP2) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Called by Phaser.Physics when a State swap occurs.
// Starts the begin and end Contact listeners again.
func (self *PhysicsP2) Reset() {
    self.Object.Call("reset")
}

// Called by Phaser.Physics when a State swap occurs.
// Starts the begin and end Contact listeners again.
func (self *PhysicsP2) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Clears all bodies from the simulation, resets callbacks and resets the collision bitmask.
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

// Clears all bodies from the simulation, resets callbacks and resets the collision bitmask.
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

// Clears all bodies from the simulation and unlinks World from Game. Should only be called on game shutdown. Call `clear` on a State change.
func (self *PhysicsP2) Destroy() {
    self.Object.Call("destroy")
}

// Clears all bodies from the simulation and unlinks World from Game. Should only be called on game shutdown. Call `clear` on a State change.
func (self *PhysicsP2) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Add a body to the world.
func (self *PhysicsP2) AddBody(body *PhysicsP2Body) bool{
    return self.Object.Call("addBody", body).Bool()
}

// Add a body to the world.
func (self *PhysicsP2) AddBodyI(args ...interface{}) bool{
    return self.Object.Call("addBody", args).Bool()
}

// Removes a body from the world. This will silently fail if the body wasn't part of the world to begin with.
func (self *PhysicsP2) RemoveBody(body *PhysicsP2Body) *PhysicsP2Body{
    return &PhysicsP2Body{self.Object.Call("removeBody", body)}
}

// Removes a body from the world. This will silently fail if the body wasn't part of the world to begin with.
func (self *PhysicsP2) RemoveBodyI(args ...interface{}) *PhysicsP2Body{
    return &PhysicsP2Body{self.Object.Call("removeBody", args)}
}

// Adds a Spring to the world.
func (self *PhysicsP2) AddSpring(spring interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("addSpring", spring)}
}

// Adds a Spring to the world.
func (self *PhysicsP2) AddSpringI(args ...interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("addSpring", args)}
}

// Removes a Spring from the world.
func (self *PhysicsP2) RemoveSpring(spring *PhysicsP2Spring) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("removeSpring", spring)}
}

// Removes a Spring from the world.
func (self *PhysicsP2) RemoveSpringI(args ...interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("removeSpring", args)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraint(bodyA interface{}, bodyB interface{}, distance int) *PhysicsP2DistanceConstraint{
    return &PhysicsP2DistanceConstraint{self.Object.Call("createDistanceConstraint", bodyA, bodyB, distance)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraint1O(bodyA interface{}, bodyB interface{}, distance int, localAnchorA []interface{}) *PhysicsP2DistanceConstraint{
    return &PhysicsP2DistanceConstraint{self.Object.Call("createDistanceConstraint", bodyA, bodyB, distance, localAnchorA)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraint2O(bodyA interface{}, bodyB interface{}, distance int, localAnchorA []interface{}, localAnchorB []interface{}) *PhysicsP2DistanceConstraint{
    return &PhysicsP2DistanceConstraint{self.Object.Call("createDistanceConstraint", bodyA, bodyB, distance, localAnchorA, localAnchorB)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraint3O(bodyA interface{}, bodyB interface{}, distance int, localAnchorA []interface{}, localAnchorB []interface{}, maxForce int) *PhysicsP2DistanceConstraint{
    return &PhysicsP2DistanceConstraint{self.Object.Call("createDistanceConstraint", bodyA, bodyB, distance, localAnchorA, localAnchorB, maxForce)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraintI(args ...interface{}) *PhysicsP2DistanceConstraint{
    return &PhysicsP2DistanceConstraint{self.Object.Call("createDistanceConstraint", args)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateGearConstraint(bodyA interface{}, bodyB interface{}) *PhysicsP2GearConstraint{
    return &PhysicsP2GearConstraint{self.Object.Call("createGearConstraint", bodyA, bodyB)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateGearConstraint1O(bodyA interface{}, bodyB interface{}, angle int) *PhysicsP2GearConstraint{
    return &PhysicsP2GearConstraint{self.Object.Call("createGearConstraint", bodyA, bodyB, angle)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateGearConstraint2O(bodyA interface{}, bodyB interface{}, angle int, ratio int) *PhysicsP2GearConstraint{
    return &PhysicsP2GearConstraint{self.Object.Call("createGearConstraint", bodyA, bodyB, angle, ratio)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateGearConstraintI(args ...interface{}) *PhysicsP2GearConstraint{
    return &PhysicsP2GearConstraint{self.Object.Call("createGearConstraint", args)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func (self *PhysicsP2) CreateRevoluteConstraint(bodyA interface{}, pivotA []interface{}, bodyB interface{}, pivotB []interface{}) *PhysicsP2RevoluteConstraint{
    return &PhysicsP2RevoluteConstraint{self.Object.Call("createRevoluteConstraint", bodyA, pivotA, bodyB, pivotB)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func (self *PhysicsP2) CreateRevoluteConstraint1O(bodyA interface{}, pivotA []interface{}, bodyB interface{}, pivotB []interface{}, maxForce int) *PhysicsP2RevoluteConstraint{
    return &PhysicsP2RevoluteConstraint{self.Object.Call("createRevoluteConstraint", bodyA, pivotA, bodyB, pivotB, maxForce)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func (self *PhysicsP2) CreateRevoluteConstraint2O(bodyA interface{}, pivotA []interface{}, bodyB interface{}, pivotB []interface{}, maxForce int, worldPivot *Float32Array) *PhysicsP2RevoluteConstraint{
    return &PhysicsP2RevoluteConstraint{self.Object.Call("createRevoluteConstraint", bodyA, pivotA, bodyB, pivotB, maxForce, worldPivot)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func (self *PhysicsP2) CreateRevoluteConstraintI(args ...interface{}) *PhysicsP2RevoluteConstraint{
    return &PhysicsP2RevoluteConstraint{self.Object.Call("createRevoluteConstraint", args)}
}

// Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraint(bodyA interface{}, bodyB interface{}) *PhysicsP2LockConstraint{
    return &PhysicsP2LockConstraint{self.Object.Call("createLockConstraint", bodyA, bodyB)}
}

// Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraint1O(bodyA interface{}, bodyB interface{}, offset []interface{}) *PhysicsP2LockConstraint{
    return &PhysicsP2LockConstraint{self.Object.Call("createLockConstraint", bodyA, bodyB, offset)}
}

// Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraint2O(bodyA interface{}, bodyB interface{}, offset []interface{}, angle int) *PhysicsP2LockConstraint{
    return &PhysicsP2LockConstraint{self.Object.Call("createLockConstraint", bodyA, bodyB, offset, angle)}
}

// Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraint3O(bodyA interface{}, bodyB interface{}, offset []interface{}, angle int, maxForce int) *PhysicsP2LockConstraint{
    return &PhysicsP2LockConstraint{self.Object.Call("createLockConstraint", bodyA, bodyB, offset, angle, maxForce)}
}

// Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraintI(args ...interface{}) *PhysicsP2LockConstraint{
    return &PhysicsP2LockConstraint{self.Object.Call("createLockConstraint", args)}
}

// Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint(bodyA interface{}, bodyB interface{}) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB)}
}

// Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint1O(bodyA interface{}, bodyB interface{}, lockRotation bool) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB, lockRotation)}
}

// Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint2O(bodyA interface{}, bodyB interface{}, lockRotation bool, anchorA []interface{}) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB, lockRotation, anchorA)}
}

// Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint3O(bodyA interface{}, bodyB interface{}, lockRotation bool, anchorA []interface{}, anchorB []interface{}) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB, lockRotation, anchorA, anchorB)}
}

// Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint4O(bodyA interface{}, bodyB interface{}, lockRotation bool, anchorA []interface{}, anchorB []interface{}, axis []interface{}) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB, lockRotation, anchorA, anchorB, axis)}
}

// Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraint5O(bodyA interface{}, bodyB interface{}, lockRotation bool, anchorA []interface{}, anchorB []interface{}, axis []interface{}, maxForce int) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", bodyA, bodyB, lockRotation, anchorA, anchorB, axis, maxForce)}
}

// Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraintI(args ...interface{}) *PhysicsP2PrismaticConstraint{
    return &PhysicsP2PrismaticConstraint{self.Object.Call("createPrismaticConstraint", args)}
}

// Adds a Constraint to the world.
func (self *PhysicsP2) AddConstraint(constraint *PhysicsP2Constraint) *PhysicsP2Constraint{
    return &PhysicsP2Constraint{self.Object.Call("addConstraint", constraint)}
}

// Adds a Constraint to the world.
func (self *PhysicsP2) AddConstraintI(args ...interface{}) *PhysicsP2Constraint{
    return &PhysicsP2Constraint{self.Object.Call("addConstraint", args)}
}

// Removes a Constraint from the world.
func (self *PhysicsP2) RemoveConstraint(constraint *PhysicsP2Constraint) *PhysicsP2Constraint{
    return &PhysicsP2Constraint{self.Object.Call("removeConstraint", constraint)}
}

// Removes a Constraint from the world.
func (self *PhysicsP2) RemoveConstraintI(args ...interface{}) *PhysicsP2Constraint{
    return &PhysicsP2Constraint{self.Object.Call("removeConstraint", args)}
}

// Adds a Contact Material to the world.
func (self *PhysicsP2) AddContactMaterial(material *PhysicsP2ContactMaterial) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("addContactMaterial", material)}
}

// Adds a Contact Material to the world.
func (self *PhysicsP2) AddContactMaterialI(args ...interface{}) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("addContactMaterial", args)}
}

// Removes a Contact Material from the world.
func (self *PhysicsP2) RemoveContactMaterial(material *PhysicsP2ContactMaterial) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("removeContactMaterial", material)}
}

// Removes a Contact Material from the world.
func (self *PhysicsP2) RemoveContactMaterialI(args ...interface{}) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("removeContactMaterial", args)}
}

// Gets a Contact Material based on the two given Materials.
func (self *PhysicsP2) GetContactMaterial(materialA *PhysicsP2Material, materialB *PhysicsP2Material) interface{}{
    return self.Object.Call("getContactMaterial", materialA, materialB)
}

// Gets a Contact Material based on the two given Materials.
func (self *PhysicsP2) GetContactMaterialI(args ...interface{}) interface{}{
    return self.Object.Call("getContactMaterial", args)
}

// Sets the given Material against all Shapes owned by all the Bodies in the given array.
func (self *PhysicsP2) SetMaterial(material *PhysicsP2Material, bodies []PhysicsP2Body) {
    self.Object.Call("setMaterial", material, bodies)
}

// Sets the given Material against all Shapes owned by all the Bodies in the given array.
func (self *PhysicsP2) SetMaterialI(args ...interface{}) {
    self.Object.Call("setMaterial", args)
}

// Creates a Material. Materials are applied to Shapes owned by a Body and can be set with Body.setMaterial().
// Materials are a way to control what happens when Shapes collide. Combine unique Materials together to create Contact Materials.
// Contact Materials have properties such as friction and restitution that allow for fine-grained collision control between different Materials.
func (self *PhysicsP2) CreateMaterial() *PhysicsP2Material{
    return &PhysicsP2Material{self.Object.Call("createMaterial")}
}

// Creates a Material. Materials are applied to Shapes owned by a Body and can be set with Body.setMaterial().
// Materials are a way to control what happens when Shapes collide. Combine unique Materials together to create Contact Materials.
// Contact Materials have properties such as friction and restitution that allow for fine-grained collision control between different Materials.
func (self *PhysicsP2) CreateMaterial1O(name string) *PhysicsP2Material{
    return &PhysicsP2Material{self.Object.Call("createMaterial", name)}
}

// Creates a Material. Materials are applied to Shapes owned by a Body and can be set with Body.setMaterial().
// Materials are a way to control what happens when Shapes collide. Combine unique Materials together to create Contact Materials.
// Contact Materials have properties such as friction and restitution that allow for fine-grained collision control between different Materials.
func (self *PhysicsP2) CreateMaterial2O(name string, body *PhysicsP2Body) *PhysicsP2Material{
    return &PhysicsP2Material{self.Object.Call("createMaterial", name, body)}
}

// Creates a Material. Materials are applied to Shapes owned by a Body and can be set with Body.setMaterial().
// Materials are a way to control what happens when Shapes collide. Combine unique Materials together to create Contact Materials.
// Contact Materials have properties such as friction and restitution that allow for fine-grained collision control between different Materials.
func (self *PhysicsP2) CreateMaterialI(args ...interface{}) *PhysicsP2Material{
    return &PhysicsP2Material{self.Object.Call("createMaterial", args)}
}

// Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterial() *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("createContactMaterial")}
}

// Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterial1O(materialA *PhysicsP2Material) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("createContactMaterial", materialA)}
}

// Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterial2O(materialA *PhysicsP2Material, materialB *PhysicsP2Material) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("createContactMaterial", materialA, materialB)}
}

// Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterial3O(materialA *PhysicsP2Material, materialB *PhysicsP2Material, options interface{}) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("createContactMaterial", materialA, materialB, options)}
}

// Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterialI(args ...interface{}) *PhysicsP2ContactMaterial{
    return &PhysicsP2ContactMaterial{self.Object.Call("createContactMaterial", args)}
}

// Populates and returns an array with references to of all current Bodies in the world.
func (self *PhysicsP2) GetBodies() []PhysicsP2Body{
	array00 := self.Object.Call("getBodies")
	length00 := array00.Length()
	out00 := make([]PhysicsP2Body, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = PhysicsP2Body{array00.Index(i00)}
	}
	return out00
}

// Populates and returns an array with references to of all current Bodies in the world.
func (self *PhysicsP2) GetBodiesI(args ...interface{}) []PhysicsP2Body{
	array00 := self.Object.Call("getBodies", args)
	length00 := array00.Length()
	out00 := make([]PhysicsP2Body, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = PhysicsP2Body{array00.Index(i00)}
	}
	return out00
}

// Checks the given object to see if it has a p2.Body and if so returns it.
func (self *PhysicsP2) GetBody(object interface{}) *P2Body{
    return &P2Body{self.Object.Call("getBody", object)}
}

// Checks the given object to see if it has a p2.Body and if so returns it.
func (self *PhysicsP2) GetBodyI(args ...interface{}) *P2Body{
    return &P2Body{self.Object.Call("getBody", args)}
}

// Populates and returns an array of all current Springs in the world.
func (self *PhysicsP2) GetSprings() []PhysicsP2Spring{
	array00 := self.Object.Call("getSprings")
	length00 := array00.Length()
	out00 := make([]PhysicsP2Spring, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = PhysicsP2Spring{array00.Index(i00)}
	}
	return out00
}

// Populates and returns an array of all current Springs in the world.
func (self *PhysicsP2) GetSpringsI(args ...interface{}) []PhysicsP2Spring{
	array00 := self.Object.Call("getSprings", args)
	length00 := array00.Length()
	out00 := make([]PhysicsP2Spring, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = PhysicsP2Spring{array00.Index(i00)}
	}
	return out00
}

// Populates and returns an array of all current Constraints in the world.
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

// Populates and returns an array of all current Constraints in the world.
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

// Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
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

// Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
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

// Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
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

// Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
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

// Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
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

// Converts the current world into a JSON object.
func (self *PhysicsP2) ToJSON() interface{}{
    return self.Object.Call("toJSON")
}

// Converts the current world into a JSON object.
func (self *PhysicsP2) ToJSONI(args ...interface{}) interface{}{
    return self.Object.Call("toJSON", args)
}

// Creates a new Collision Group and optionally applies it to the given object.
// Collision Groups are handled using bitmasks, therefore you have a fixed limit you can create before you need to re-use older groups.
func (self *PhysicsP2) CreateCollisionGroup() {
    self.Object.Call("createCollisionGroup")
}

// Creates a new Collision Group and optionally applies it to the given object.
// Collision Groups are handled using bitmasks, therefore you have a fixed limit you can create before you need to re-use older groups.
func (self *PhysicsP2) CreateCollisionGroup1O(object interface{}) {
    self.Object.Call("createCollisionGroup", object)
}

// Creates a new Collision Group and optionally applies it to the given object.
// Collision Groups are handled using bitmasks, therefore you have a fixed limit you can create before you need to re-use older groups.
func (self *PhysicsP2) CreateCollisionGroupI(args ...interface{}) {
    self.Object.Call("createCollisionGroup", args)
}

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring(bodyA interface{}, bodyB interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB)}
}

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring1O(bodyA interface{}, bodyB interface{}, restLength int) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength)}
}

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring2O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness)}
}

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring3O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int, damping int) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness, damping)}
}

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring4O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int, damping int, worldA []interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness, damping, worldA)}
}

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring5O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int, damping int, worldA []interface{}, worldB []interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness, damping, worldA, worldB)}
}

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring6O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int, damping int, worldA []interface{}, worldB []interface{}, localA []interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness, damping, worldA, worldB, localA)}
}

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpring7O(bodyA interface{}, bodyB interface{}, restLength int, stiffness int, damping int, worldA []interface{}, worldB []interface{}, localA []interface{}, localB []interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", bodyA, bodyB, restLength, stiffness, damping, worldA, worldB, localA, localB)}
}

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpringI(args ...interface{}) *PhysicsP2Spring{
    return &PhysicsP2Spring{self.Object.Call("createSpring", args)}
}

// Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpring(bodyA interface{}, bodyB interface{}) *PhysicsP2RotationalSpring{
    return &PhysicsP2RotationalSpring{self.Object.Call("createRotationalSpring", bodyA, bodyB)}
}

// Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpring1O(bodyA interface{}, bodyB interface{}, restAngle int) *PhysicsP2RotationalSpring{
    return &PhysicsP2RotationalSpring{self.Object.Call("createRotationalSpring", bodyA, bodyB, restAngle)}
}

// Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpring2O(bodyA interface{}, bodyB interface{}, restAngle int, stiffness int) *PhysicsP2RotationalSpring{
    return &PhysicsP2RotationalSpring{self.Object.Call("createRotationalSpring", bodyA, bodyB, restAngle, stiffness)}
}

// Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpring3O(bodyA interface{}, bodyB interface{}, restAngle int, stiffness int, damping int) *PhysicsP2RotationalSpring{
    return &PhysicsP2RotationalSpring{self.Object.Call("createRotationalSpring", bodyA, bodyB, restAngle, stiffness, damping)}
}

// Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpringI(args ...interface{}) *PhysicsP2RotationalSpring{
    return &PhysicsP2RotationalSpring{self.Object.Call("createRotationalSpring", args)}
}

// Creates a new Body and adds it to the World.
func (self *PhysicsP2) CreateBody(x int, y int, mass int, addToWorld bool, options interface{}, points interface{}) *PhysicsP2Body{
    return &PhysicsP2Body{self.Object.Call("createBody", x, y, mass, addToWorld, options, points)}
}

// Creates a new Body and adds it to the World.
func (self *PhysicsP2) CreateBodyI(args ...interface{}) *PhysicsP2Body{
    return &PhysicsP2Body{self.Object.Call("createBody", args)}
}

// Creates a new Particle and adds it to the World.
func (self *PhysicsP2) CreateParticle(x int, y int, mass int, addToWorld bool, options interface{}, points interface{}) {
    self.Object.Call("createParticle", x, y, mass, addToWorld, options, points)
}

// Creates a new Particle and adds it to the World.
func (self *PhysicsP2) CreateParticleI(args ...interface{}) {
    self.Object.Call("createParticle", args)
}

// Converts all of the polylines objects inside a Tiled ObjectGroup into physics bodies that are added to the world.
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

// Converts all of the polylines objects inside a Tiled ObjectGroup into physics bodies that are added to the world.
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

// Converts all of the polylines objects inside a Tiled ObjectGroup into physics bodies that are added to the world.
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

// Converts all of the polylines objects inside a Tiled ObjectGroup into physics bodies that are added to the world.
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

// Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsP2) ClearTilemapLayerBodies(map_ *Tilemap) {
    self.Object.Call("clearTilemapLayerBodies", map_)
}

// Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsP2) ClearTilemapLayerBodies1O(map_ *Tilemap, layer interface{}) {
    self.Object.Call("clearTilemapLayerBodies", map_, layer)
}

// Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsP2) ClearTilemapLayerBodiesI(args ...interface{}) {
    self.Object.Call("clearTilemapLayerBodies", args)
}

// Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
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

// Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
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

// Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
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

// Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
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

// Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
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

// Convert p2 physics value (meters) to pixel scale.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) Mpx(v int) int{
    return self.Object.Call("mpx", v).Int()
}

// Convert p2 physics value (meters) to pixel scale.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) MpxI(args ...interface{}) int{
    return self.Object.Call("mpx", args).Int()
}

// Convert pixel value to p2 physics scale (meters).
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) Pxm(v int) int{
    return self.Object.Call("pxm", v).Int()
}

// Convert pixel value to p2 physics scale (meters).
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) PxmI(args ...interface{}) int{
    return self.Object.Call("pxm", args).Int()
}

// Convert p2 physics value (meters) to pixel scale and inverses it.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) Mpxi(v int) int{
    return self.Object.Call("mpxi", v).Int()
}

// Convert p2 physics value (meters) to pixel scale and inverses it.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) MpxiI(args ...interface{}) int{
    return self.Object.Call("mpxi", args).Int()
}

// Convert pixel value to p2 physics scale (meters) and inverses it.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) Pxmi(v int) int{
    return self.Object.Call("pxmi", v).Int()
}

// Convert pixel value to p2 physics scale (meters) and inverses it.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) PxmiI(args ...interface{}) int{
    return self.Object.Call("pxmi", args).Int()
}
