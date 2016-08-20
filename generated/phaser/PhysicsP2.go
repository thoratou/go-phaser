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


// Local reference to game.
func (self *PhysicsP2) GetGame() Game{
    return Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2) SetGame(member Game) {
    self.Set("game", member)
}

// The p2 World configuration object.
func (self *PhysicsP2) GetConfig() interface{}{
    return self.Get("config")
}

// The p2 World configuration object.
func (self *PhysicsP2) SetConfig(member interface{}) {
    self.Set("config", member)
}

// The p2 World in which the simulation is run.
func (self *PhysicsP2) GetWorld() P2World{
    return P2World{self.Get("world")}
}

// The p2 World in which the simulation is run.
func (self *PhysicsP2) SetWorld(member P2World) {
    self.Set("world", member)
}

// The frame rate the world will be stepped at. Defaults to 1 / 60, but you can change here. Also see useElapsedTime property.
func (self *PhysicsP2) GetFrameRate() float64{
    return self.Get("frameRate").Float()
}

// The frame rate the world will be stepped at. Defaults to 1 / 60, but you can change here. Also see useElapsedTime property.
func (self *PhysicsP2) SetFrameRate(member float64) {
    self.Set("frameRate", member)
}

// If true the frameRate value will be ignored and instead p2 will step with the value of Game.Time.physicsElapsed, which is a delta time value.
func (self *PhysicsP2) GetUseElapsedTime() bool{
    return self.Get("useElapsedTime").Bool()
}

// If true the frameRate value will be ignored and instead p2 will step with the value of Game.Time.physicsElapsed, which is a delta time value.
func (self *PhysicsP2) SetUseElapsedTime(member bool) {
    self.Set("useElapsedTime", member)
}

// The paused state of the P2 World.
func (self *PhysicsP2) GetPaused() bool{
    return self.Get("paused").Bool()
}

// The paused state of the P2 World.
func (self *PhysicsP2) SetPaused(member bool) {
    self.Set("paused", member)
}

// A local array of all created Materials.
func (self *PhysicsP2) GetMaterials() []PhysicsP2Material{
	array := self.Get("materials")
	length := array.Length()
	out := make([]PhysicsP2Material, length, length)
	for i := 0; i < length; i++ {
		out[i] = PhysicsP2Material{array.Index(i)}
	}
	return out
}

// A local array of all created Materials.
func (self *PhysicsP2) SetMaterials(member []PhysicsP2Material) {
    self.Set("materials", member)
}

// The gravity applied to all bodies each step.
func (self *PhysicsP2) GetGravity() PhysicsP2InversePointProxy{
    return PhysicsP2InversePointProxy{self.Get("gravity")}
}

// The gravity applied to all bodies each step.
func (self *PhysicsP2) SetGravity(member PhysicsP2InversePointProxy) {
    self.Set("gravity", member)
}

// An object containing the 4 wall bodies that bound the physics world.
func (self *PhysicsP2) GetWalls() interface{}{
    return self.Get("walls")
}

// An object containing the 4 wall bodies that bound the physics world.
func (self *PhysicsP2) SetWalls(member interface{}) {
    self.Set("walls", member)
}

// This signal is dispatched when a new Body is added to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was added to the world.
func (self *PhysicsP2) GetOnBodyAdded() Signal{
    return Signal{self.Get("onBodyAdded")}
}

// This signal is dispatched when a new Body is added to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was added to the world.
func (self *PhysicsP2) SetOnBodyAdded(member Signal) {
    self.Set("onBodyAdded", member)
}

// This signal is dispatched when a Body is removed to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was removed from the world.
func (self *PhysicsP2) GetOnBodyRemoved() Signal{
    return Signal{self.Get("onBodyRemoved")}
}

// This signal is dispatched when a Body is removed to the World.
// 
// It sends 1 argument: `body` which is the `Phaser.Physics.P2.Body` that was removed from the world.
func (self *PhysicsP2) SetOnBodyRemoved(member Signal) {
    self.Set("onBodyRemoved", member)
}

// This signal is dispatched when a Spring is added to the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was added to the world.
func (self *PhysicsP2) GetOnSpringAdded() Signal{
    return Signal{self.Get("onSpringAdded")}
}

// This signal is dispatched when a Spring is added to the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was added to the world.
func (self *PhysicsP2) SetOnSpringAdded(member Signal) {
    self.Set("onSpringAdded", member)
}

// This signal is dispatched when a Spring is removed from the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was removed from the world.
func (self *PhysicsP2) GetOnSpringRemoved() Signal{
    return Signal{self.Get("onSpringRemoved")}
}

// This signal is dispatched when a Spring is removed from the World.
// 
// It sends 1 argument: `spring` which is either a `Phaser.Physics.P2.Spring`, `p2.LinearSpring` or `p2.RotationalSpring` that was removed from the world.
func (self *PhysicsP2) SetOnSpringRemoved(member Signal) {
    self.Set("onSpringRemoved", member)
}

// This signal is dispatched when a Constraint is added to the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was added to the world.
func (self *PhysicsP2) GetOnConstraintAdded() Signal{
    return Signal{self.Get("onConstraintAdded")}
}

// This signal is dispatched when a Constraint is added to the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was added to the world.
func (self *PhysicsP2) SetOnConstraintAdded(member Signal) {
    self.Set("onConstraintAdded", member)
}

// This signal is dispatched when a Constraint is removed from the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was removed from the world.
func (self *PhysicsP2) GetOnConstraintRemoved() Signal{
    return Signal{self.Get("onConstraintRemoved")}
}

// This signal is dispatched when a Constraint is removed from the World.
// 
// It sends 1 argument: `constraint` which is the `Phaser.Physics.P2.Constraint` that was removed from the world.
func (self *PhysicsP2) SetOnConstraintRemoved(member Signal) {
    self.Set("onConstraintRemoved", member)
}

// This signal is dispatched when a Contact Material is added to the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was added to the world.
func (self *PhysicsP2) GetOnContactMaterialAdded() Signal{
    return Signal{self.Get("onContactMaterialAdded")}
}

// This signal is dispatched when a Contact Material is added to the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was added to the world.
func (self *PhysicsP2) SetOnContactMaterialAdded(member Signal) {
    self.Set("onContactMaterialAdded", member)
}

// This signal is dispatched when a Contact Material is removed from the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was removed from the world.
func (self *PhysicsP2) GetOnContactMaterialRemoved() Signal{
    return Signal{self.Get("onContactMaterialRemoved")}
}

// This signal is dispatched when a Contact Material is removed from the World.
// 
// It sends 1 argument: `material` which is the `Phaser.Physics.P2.ContactMaterial` that was removed from the world.
func (self *PhysicsP2) SetOnContactMaterialRemoved(member Signal) {
    self.Set("onContactMaterialRemoved", member)
}

// A postBroadphase callback.
func (self *PhysicsP2) SetPostBroadphaseCallback(member func(...interface{})) {
    self.Set("postBroadphaseCallback", member)
}

// The context under which the callbacks are fired.
func (self *PhysicsP2) GetCallbackContext() interface{}{
    return self.Get("callbackContext")
}

// The context under which the callbacks are fired.
func (self *PhysicsP2) SetCallbackContext(member interface{}) {
    self.Set("callbackContext", member)
}

// This Signal is dispatched when a first contact is created between two bodies. This happens *before* the step has been done.
// 
// It sends 5 arguments: `bodyA`, `bodyB`, `shapeA`, `shapeB` and `contactEquations`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) GetOnBeginContact() Signal{
    return Signal{self.Get("onBeginContact")}
}

// This Signal is dispatched when a first contact is created between two bodies. This happens *before* the step has been done.
// 
// It sends 5 arguments: `bodyA`, `bodyB`, `shapeA`, `shapeB` and `contactEquations`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) SetOnBeginContact(member Signal) {
    self.Set("onBeginContact", member)
}

// This Signal is dispatched when final contact occurs between two bodies. This happens *before* the step has been done.
// 
// It sends 4 arguments: `bodyA`, `bodyB`, `shapeA` and `shapeB`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) GetOnEndContact() Signal{
    return Signal{self.Get("onEndContact")}
}

// This Signal is dispatched when final contact occurs between two bodies. This happens *before* the step has been done.
// 
// It sends 4 arguments: `bodyA`, `bodyB`, `shapeA` and `shapeB`.
// 
// It is possible that in certain situations the `bodyA` or `bodyB` values are `null`. You should check for this
// in your own code to avoid processing potentially null physics bodies.
func (self *PhysicsP2) SetOnEndContact(member Signal) {
    self.Set("onEndContact", member)
}

// An array containing the collision groups that have been defined in the World.
func (self *PhysicsP2) GetCollisionGroups() []interface{}{
	array := self.Get("collisionGroups")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array containing the collision groups that have been defined in the World.
func (self *PhysicsP2) SetCollisionGroups(member []interface{}) {
    self.Set("collisionGroups", member)
}

// A default collision group.
func (self *PhysicsP2) GetNothingCollisionGroup() PhysicsP2CollisionGroup{
    return PhysicsP2CollisionGroup{self.Get("nothingCollisionGroup")}
}

// A default collision group.
func (self *PhysicsP2) SetNothingCollisionGroup(member PhysicsP2CollisionGroup) {
    self.Set("nothingCollisionGroup", member)
}

// A default collision group.
func (self *PhysicsP2) GetBoundsCollisionGroup() PhysicsP2CollisionGroup{
    return PhysicsP2CollisionGroup{self.Get("boundsCollisionGroup")}
}

// A default collision group.
func (self *PhysicsP2) SetBoundsCollisionGroup(member PhysicsP2CollisionGroup) {
    self.Set("boundsCollisionGroup", member)
}

// A default collision group.
func (self *PhysicsP2) GetEverythingCollisionGroup() PhysicsP2CollisionGroup{
    return PhysicsP2CollisionGroup{self.Get("everythingCollisionGroup")}
}

// A default collision group.
func (self *PhysicsP2) SetEverythingCollisionGroup(member PhysicsP2CollisionGroup) {
    self.Set("everythingCollisionGroup", member)
}

// An array of the bodies the world bounds collides with.
func (self *PhysicsP2) GetBoundsCollidesWith() []interface{}{
	array := self.Get("boundsCollidesWith")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of the bodies the world bounds collides with.
func (self *PhysicsP2) SetBoundsCollidesWith(member []interface{}) {
    self.Set("boundsCollidesWith", member)
}

// Friction between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) GetFriction() float64{
    return self.Get("friction").Float()
}

// Friction between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) SetFriction(member float64) {
    self.Set("friction", member)
}

// Default coefficient of restitution between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) GetRestitution() float64{
    return self.Get("restitution").Float()
}

// Default coefficient of restitution between colliding bodies. This value is used if no matching ContactMaterial is found for a Material pair.
func (self *PhysicsP2) SetRestitution(member float64) {
    self.Set("restitution", member)
}

// The default Contact Material being used by the World.
func (self *PhysicsP2) GetContactMaterial() P2ContactMaterial{
    return P2ContactMaterial{self.Get("contactMaterial")}
}

// The default Contact Material being used by the World.
func (self *PhysicsP2) SetContactMaterial(member P2ContactMaterial) {
    self.Set("contactMaterial", member)
}

// Enable to automatically apply spring forces each step.
func (self *PhysicsP2) GetApplySpringForces() bool{
    return self.Get("applySpringForces").Bool()
}

// Enable to automatically apply spring forces each step.
func (self *PhysicsP2) SetApplySpringForces(member bool) {
    self.Set("applySpringForces", member)
}

// Enable to automatically apply body damping each step.
func (self *PhysicsP2) GetApplyDamping() bool{
    return self.Get("applyDamping").Bool()
}

// Enable to automatically apply body damping each step.
func (self *PhysicsP2) SetApplyDamping(member bool) {
    self.Set("applyDamping", member)
}

// Enable to automatically apply gravity each step.
func (self *PhysicsP2) GetApplyGravity() bool{
    return self.Get("applyGravity").Bool()
}

// Enable to automatically apply gravity each step.
func (self *PhysicsP2) SetApplyGravity(member bool) {
    self.Set("applyGravity", member)
}

// Enable/disable constraint solving in each step.
func (self *PhysicsP2) GetSolveConstraints() bool{
    return self.Get("solveConstraints").Bool()
}

// Enable/disable constraint solving in each step.
func (self *PhysicsP2) SetSolveConstraints(member bool) {
    self.Set("solveConstraints", member)
}

// The World time.
func (self *PhysicsP2) GetTime() bool{
    return self.Get("time").Bool()
}

// The World time.
func (self *PhysicsP2) SetTime(member bool) {
    self.Set("time", member)
}

// Set to true if you want to the world to emit the "impact" event. Turning this off could improve performance.
func (self *PhysicsP2) GetEmitImpactEvent() bool{
    return self.Get("emitImpactEvent").Bool()
}

// Set to true if you want to the world to emit the "impact" event. Turning this off could improve performance.
func (self *PhysicsP2) SetEmitImpactEvent(member bool) {
    self.Set("emitImpactEvent", member)
}

// How to deactivate bodies during simulation. Possible modes are: World.NO_SLEEPING, World.BODY_SLEEPING and World.ISLAND_SLEEPING.
// If sleeping is enabled, you might need to wake up the bodies if they fall asleep when they shouldn't. If you want to enable sleeping in the world, but want to disable it for a particular body, see Body.allowSleep.
func (self *PhysicsP2) GetSleepMode() float64{
    return self.Get("sleepMode").Float()
}

// How to deactivate bodies during simulation. Possible modes are: World.NO_SLEEPING, World.BODY_SLEEPING and World.ISLAND_SLEEPING.
// If sleeping is enabled, you might need to wake up the bodies if they fall asleep when they shouldn't. If you want to enable sleeping in the world, but want to disable it for a particular body, see Body.allowSleep.
func (self *PhysicsP2) SetSleepMode(member float64) {
    self.Set("sleepMode", member)
}

// The total number of bodies in the world.
func (self *PhysicsP2) GetTotal() float64{
    return self.Get("total").Float()
}

// The total number of bodies in the world.
func (self *PhysicsP2) SetTotal(member float64) {
    self.Set("total", member)
}



// This will add a P2 Physics body into the removal list for the next step.
func (self *PhysicsP2) RemoveBodyNextStepI(args ...interface{}) {
    self.Call("removeBodyNextStep", args)
}

// Called at the start of the core update loop. Purges flagged bodies from the world.
func (self *PhysicsP2) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// This will create a P2 Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// Note: When the game object is enabled for P2 physics it has its anchor x/y set to 0.5 so it becomes centered.
func (self *PhysicsP2) EnableI(args ...interface{}) {
    self.Call("enable", args)
}

// Creates a P2 Physics body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
func (self *PhysicsP2) EnableBodyI(args ...interface{}) {
    self.Call("enableBody", args)
}

// Impact event handling is disabled by default. Enable it before any impact events will be dispatched.
// In a busy world hundreds of impact events can be generated every step, so only enable this if you cannot do what you need via beginContact or collision masks.
func (self *PhysicsP2) SetImpactEventsI(args ...interface{}) {
    self.Call("setImpactEvents", args)
}

// Sets a callback to be fired after the Broadphase has collected collision pairs in the world.
// Just because a pair exists it doesn't mean they *will* collide, just that they potentially could do.
// If your calback returns `false` the pair will be removed from the narrowphase. This will stop them testing for collision this step.
// Returning `true` from the callback will ensure they are checked in the narrowphase.
func (self *PhysicsP2) SetPostBroadphaseCallbackI(args ...interface{}) {
    self.Call("setPostBroadphaseCallback", args)
}

// Internal handler for the postBroadphase event.
func (self *PhysicsP2) PostBroadphaseHandlerI(args ...interface{}) {
    self.Call("postBroadphaseHandler", args)
}

// Handles a p2 impact event.
func (self *PhysicsP2) ImpactHandlerI(args ...interface{}) {
    self.Call("impactHandler", args)
}

// Handles a p2 begin contact event.
func (self *PhysicsP2) BeginContactHandlerI(args ...interface{}) {
    self.Call("beginContactHandler", args)
}

// Handles a p2 end contact event.
func (self *PhysicsP2) EndContactHandlerI(args ...interface{}) {
    self.Call("endContactHandler", args)
}

// By default the World will be set to collide everything with everything. The bounds of the world is a Body with 4 shapes, one for each face.
// If you start to use your own collision groups then your objects will no longer collide with the bounds.
// To fix this you need to adjust the bounds to use its own collision group first BEFORE changing your Sprites collision group.
func (self *PhysicsP2) UpdateBoundsCollisionGroupI(args ...interface{}) {
    self.Call("updateBoundsCollisionGroup", args)
}

// Sets the bounds of the Physics world to match the given world pixel dimensions.
// You can optionally set which 'walls' to create: left, right, top or bottom.
// If none of the walls are given it will default to use the walls settings it had previously.
// I.e. if you previously told it to not have the left or right walls, and you then adjust the world size
// the newly created bounds will also not have the left and right walls.
// Explicitly state them in the parameters to override this.
func (self *PhysicsP2) SetBoundsI(args ...interface{}) {
    self.Call("setBounds", args)
}

// Internal method called by setBounds. Responsible for creating, updating or 
// removing the wall body shapes.
func (self *PhysicsP2) SetupWallI(args ...interface{}) {
    self.Call("setupWall", args)
}

// Pauses the P2 World independent of the game pause state.
func (self *PhysicsP2) PauseI(args ...interface{}) {
    self.Call("pause", args)
}

// Resumes a paused P2 World.
func (self *PhysicsP2) ResumeI(args ...interface{}) {
    self.Call("resume", args)
}

// Internal P2 update loop.
func (self *PhysicsP2) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Called by Phaser.Physics when a State swap occurs.
// Starts the begin and end Contact listeners again.
func (self *PhysicsP2) ResetI(args ...interface{}) {
    self.Call("reset", args)
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
    self.Call("clear", args)
}

// Clears all bodies from the simulation and unlinks World from Game. Should only be called on game shutdown. Call `clear` on a State change.
func (self *PhysicsP2) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Add a body to the world.
func (self *PhysicsP2) AddBodyI(args ...interface{}) bool{
    return self.Call("addBody", args).Bool()
}

// Removes a body from the world. This will silently fail if the body wasn't part of the world to begin with.
func (self *PhysicsP2) RemoveBodyI(args ...interface{}) PhysicsP2Body{
    return PhysicsP2Body{self.Call("removeBody", args)}
}

// Adds a Spring to the world.
func (self *PhysicsP2) AddSpringI(args ...interface{}) PhysicsP2Spring{
    return PhysicsP2Spring{self.Call("addSpring", args)}
}

// Removes a Spring from the world.
func (self *PhysicsP2) RemoveSpringI(args ...interface{}) PhysicsP2Spring{
    return PhysicsP2Spring{self.Call("removeSpring", args)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateDistanceConstraintI(args ...interface{}) PhysicsP2DistanceConstraint{
    return PhysicsP2DistanceConstraint{self.Call("createDistanceConstraint", args)}
}

// Creates a constraint that tries to keep the distance between two bodies constant.
func (self *PhysicsP2) CreateGearConstraintI(args ...interface{}) PhysicsP2GearConstraint{
    return PhysicsP2GearConstraint{self.Call("createGearConstraint", args)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
func (self *PhysicsP2) CreateRevoluteConstraintI(args ...interface{}) PhysicsP2RevoluteConstraint{
    return PhysicsP2RevoluteConstraint{self.Call("createRevoluteConstraint", args)}
}

// Locks the relative position between two bodies.
func (self *PhysicsP2) CreateLockConstraintI(args ...interface{}) PhysicsP2LockConstraint{
    return PhysicsP2LockConstraint{self.Call("createLockConstraint", args)}
}

// Constraint that only allows bodies to move along a line, relative to each other.
// See http://www.iforce2d.net/b2dtut/joints-prismatic
func (self *PhysicsP2) CreatePrismaticConstraintI(args ...interface{}) PhysicsP2PrismaticConstraint{
    return PhysicsP2PrismaticConstraint{self.Call("createPrismaticConstraint", args)}
}

// Adds a Constraint to the world.
func (self *PhysicsP2) AddConstraintI(args ...interface{}) PhysicsP2Constraint{
    return PhysicsP2Constraint{self.Call("addConstraint", args)}
}

// Removes a Constraint from the world.
func (self *PhysicsP2) RemoveConstraintI(args ...interface{}) PhysicsP2Constraint{
    return PhysicsP2Constraint{self.Call("removeConstraint", args)}
}

// Adds a Contact Material to the world.
func (self *PhysicsP2) AddContactMaterialI(args ...interface{}) PhysicsP2ContactMaterial{
    return PhysicsP2ContactMaterial{self.Call("addContactMaterial", args)}
}

// Removes a Contact Material from the world.
func (self *PhysicsP2) RemoveContactMaterialI(args ...interface{}) PhysicsP2ContactMaterial{
    return PhysicsP2ContactMaterial{self.Call("removeContactMaterial", args)}
}

// Gets a Contact Material based on the two given Materials.
func (self *PhysicsP2) GetContactMaterialI(args ...interface{}) interface{}{
    return self.Call("getContactMaterial", args)
}

// Sets the given Material against all Shapes owned by all the Bodies in the given array.
func (self *PhysicsP2) SetMaterialI(args ...interface{}) {
    self.Call("setMaterial", args)
}

// Creates a Material. Materials are applied to Shapes owned by a Body and can be set with Body.setMaterial().
// Materials are a way to control what happens when Shapes collide. Combine unique Materials together to create Contact Materials.
// Contact Materials have properties such as friction and restitution that allow for fine-grained collision control between different Materials.
func (self *PhysicsP2) CreateMaterialI(args ...interface{}) PhysicsP2Material{
    return PhysicsP2Material{self.Call("createMaterial", args)}
}

// Creates a Contact Material from the two given Materials. You can then edit the properties of the Contact Material directly.
func (self *PhysicsP2) CreateContactMaterialI(args ...interface{}) PhysicsP2ContactMaterial{
    return PhysicsP2ContactMaterial{self.Call("createContactMaterial", args)}
}

// Populates and returns an array with references to of all current Bodies in the world.
func (self *PhysicsP2) GetBodiesI(args ...interface{}) []PhysicsP2Body{
	array := self.Call("getBodies", args)
	length := array.Length()
	out := make([]PhysicsP2Body, length, length)
	for i := 0; i < length; i++ {
		
			out[i] = PhysicsP2Body{array.Index(i)}
	}
	return out
}

// Checks the given object to see if it has a p2.Body and if so returns it.
func (self *PhysicsP2) GetBodyI(args ...interface{}) P2Body{
    return P2Body{self.Call("getBody", args)}
}

// Populates and returns an array of all current Springs in the world.
func (self *PhysicsP2) GetSpringsI(args ...interface{}) []PhysicsP2Spring{
	array := self.Call("getSprings", args)
	length := array.Length()
	out := make([]PhysicsP2Spring, length, length)
	for i := 0; i < length; i++ {
		
			out[i] = PhysicsP2Spring{array.Index(i)}
	}
	return out
}

// Populates and returns an array of all current Constraints in the world.
// You will get an array of p2 constraints back. This can be of mixed types, for example the array may contain
// PrismaticConstraints, RevoluteConstraints or any other valid p2 constraint type.
func (self *PhysicsP2) GetConstraintsI(args ...interface{}) []PhysicsP2Constraint{
	array := self.Call("getConstraints", args)
	length := array.Length()
	out := make([]PhysicsP2Constraint, length, length)
	for i := 0; i < length; i++ {
		
			out[i] = PhysicsP2Constraint{array.Index(i)}
	}
	return out
}

// Test if a world point overlaps bodies. You will get an array of actual P2 bodies back. You can find out which Sprite a Body belongs to
// (if any) by checking the Body.parent.sprite property. Body.parent is a Phaser.Physics.P2.Body property.
func (self *PhysicsP2) HitTestI(args ...interface{}) []interface{}{
	array := self.Call("hitTest", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Converts the current world into a JSON object.
func (self *PhysicsP2) ToJSONI(args ...interface{}) interface{}{
    return self.Call("toJSON", args)
}

// Creates a new Collision Group and optionally applies it to the given object.
// Collision Groups are handled using bitmasks, therefore you have a fixed limit you can create before you need to re-use older groups.
func (self *PhysicsP2) CreateCollisionGroupI(args ...interface{}) {
    self.Call("createCollisionGroup", args)
}

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateSpringI(args ...interface{}) PhysicsP2Spring{
    return PhysicsP2Spring{self.Call("createSpring", args)}
}

// Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func (self *PhysicsP2) CreateRotationalSpringI(args ...interface{}) PhysicsP2RotationalSpring{
    return PhysicsP2RotationalSpring{self.Call("createRotationalSpring", args)}
}

// Creates a new Body and adds it to the World.
func (self *PhysicsP2) CreateBodyI(args ...interface{}) PhysicsP2Body{
    return PhysicsP2Body{self.Call("createBody", args)}
}

// Creates a new Particle and adds it to the World.
func (self *PhysicsP2) CreateParticleI(args ...interface{}) {
    self.Call("createParticle", args)
}

// Converts all of the polylines objects inside a Tiled ObjectGroup into physics bodies that are added to the world.
// Note that the polylines must be created in such a way that they can withstand polygon decomposition.
func (self *PhysicsP2) ConvertCollisionObjectsI(args ...interface{}) []interface{}{
	array := self.Call("convertCollisionObjects", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsP2) ClearTilemapLayerBodiesI(args ...interface{}) {
    self.Call("clearTilemapLayerBodies", args)
}

// Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics bodies.
// Only call this *after* you have specified all of the tiles you wish to collide with calls like Tilemap.setCollisionBetween, etc.
// Every time you call this method it will destroy any previously created bodies and remove them from the world.
// Therefore understand it's a very expensive operation and not to be done in a core game update loop.
func (self *PhysicsP2) ConvertTilemapI(args ...interface{}) []interface{}{
	array := self.Call("convertTilemap", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Convert p2 physics value (meters) to pixel scale.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) MpxI(args ...interface{}) float64{
    return self.Call("mpx", args).Float()
}

// Convert pixel value to p2 physics scale (meters).
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) PxmI(args ...interface{}) float64{
    return self.Call("pxm", args).Float()
}

// Convert p2 physics value (meters) to pixel scale and inverses it.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) MpxiI(args ...interface{}) float64{
    return self.Call("mpxi", args).Float()
}

// Convert pixel value to p2 physics scale (meters) and inverses it.
// By default Phaser uses a scale of 20px per meter.
// If you need to modify this you can over-ride these functions via the Physics Configuration object.
func (self *PhysicsP2) PxmiI(args ...interface{}) float64{
    return self.Call("pxmi", args).Float()
}
