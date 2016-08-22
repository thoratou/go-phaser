// Automatic generation for Phaser.Physics.P2.Body
// generated file PhysicsP2Body.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Physics Body is typically linked to a single Sprite and defines properties that determine how the physics body is simulated.
// These properties affect how the body reacts to forces, what forces it generates on itself (to simulate friction), and how it reacts to collisions in the scene.
// In most cases, the properties are used to simulate physical effects. Each body also has its own property values that determine exactly how it reacts to forces and collisions in the scene.
// By default a single Rectangle shape is added to the Body that matches the dimensions of the parent Sprite. See addShape, removeShape, clearShapes to add extra shapes around the Body.
// Note: When bound to a Sprite to avoid single-pixel jitters on mobile devices we strongly recommend using Sprite sizes that are even on both axis, i.e. 128x128 not 127x127.
// Note: When a game object is given a P2 body it has its anchor x/y set to 0.5, so it becomes centered.
type PhysicsP2Body struct {
    *js.Object
}


// Local reference to game.
func (self *PhysicsP2Body) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2Body) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to the P2 World.
func (self *PhysicsP2Body) GetWorldA() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// Local reference to the P2 World.
func (self *PhysicsP2Body) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}

// Reference to the parent Sprite.
func (self *PhysicsP2Body) GetSpriteA() *Sprite{
    return &Sprite{self.Object.Get("sprite")}
}

// Reference to the parent Sprite.
func (self *PhysicsP2Body) SetSpriteA(member *Sprite) {
    self.Object.Set("sprite", member)
}

// The type of physics system this body belongs to.
func (self *PhysicsP2Body) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The type of physics system this body belongs to.
func (self *PhysicsP2Body) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsP2Body) GetOffsetA() *Point{
    return &Point{self.Object.Get("offset")}
}

// The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsP2Body) SetOffsetA(member *Point) {
    self.Object.Set("offset", member)
}

// The p2 Body data.
func (self *PhysicsP2Body) GetDataA() *P2Body{
    return &P2Body{self.Object.Get("data")}
}

// The p2 Body data.
func (self *PhysicsP2Body) SetDataA(member *P2Body) {
    self.Object.Set("data", member)
}

// The velocity of the body. Set velocity.x to a negative value to move to the left, position to the right. velocity.y negative values move up, positive move down.
func (self *PhysicsP2Body) GetVelocityA() *PhysicsP2InversePointProxy{
    return &PhysicsP2InversePointProxy{self.Object.Get("velocity")}
}

// The velocity of the body. Set velocity.x to a negative value to move to the left, position to the right. velocity.y negative values move up, positive move down.
func (self *PhysicsP2Body) SetVelocityA(member *PhysicsP2InversePointProxy) {
    self.Object.Set("velocity", member)
}

// The force applied to the body.
func (self *PhysicsP2Body) GetForceA() *PhysicsP2InversePointProxy{
    return &PhysicsP2InversePointProxy{self.Object.Get("force")}
}

// The force applied to the body.
func (self *PhysicsP2Body) SetForceA(member *PhysicsP2InversePointProxy) {
    self.Object.Set("force", member)
}

// A locally applied gravity force to the Body. Applied directly before the world step. NOTE: Not currently implemented.
func (self *PhysicsP2Body) GetGravityA() *Point{
    return &Point{self.Object.Get("gravity")}
}

// A locally applied gravity force to the Body. Applied directly before the world step. NOTE: Not currently implemented.
func (self *PhysicsP2Body) SetGravityA(member *Point) {
    self.Object.Set("gravity", member)
}

// Dispatched when a first contact is created between shapes in two bodies. 
// This event is fired during the step, so collision has already taken place.
// 
// The event will be sent 5 arguments in this order:
// 
// The Phaser.Physics.P2.Body it is in contact with. *This might be null* if the Body was created directly in the p2 world.
// The p2.Body this Body is in contact with.
// The Shape from this body that caused the contact.
// The Shape from the contact body.
// The Contact Equation data array.
func (self *PhysicsP2Body) GetOnBeginContactA() *Signal{
    return &Signal{self.Object.Get("onBeginContact")}
}

// Dispatched when a first contact is created between shapes in two bodies. 
// This event is fired during the step, so collision has already taken place.
// 
// The event will be sent 5 arguments in this order:
// 
// The Phaser.Physics.P2.Body it is in contact with. *This might be null* if the Body was created directly in the p2 world.
// The p2.Body this Body is in contact with.
// The Shape from this body that caused the contact.
// The Shape from the contact body.
// The Contact Equation data array.
func (self *PhysicsP2Body) SetOnBeginContactA(member *Signal) {
    self.Object.Set("onBeginContact", member)
}

// Dispatched when contact ends between shapes in two bodies.
// This event is fired during the step, so collision has already taken place.
// 
// The event will be sent 4 arguments in this order:
// 
// The Phaser.Physics.P2.Body it is in contact with. *This might be null* if the Body was created directly in the p2 world.
// The p2.Body this Body has ended contact with.
// The Shape from this body that caused the original contact.
// The Shape from the contact body.
func (self *PhysicsP2Body) GetOnEndContactA() *Signal{
    return &Signal{self.Object.Get("onEndContact")}
}

// Dispatched when contact ends between shapes in two bodies.
// This event is fired during the step, so collision has already taken place.
// 
// The event will be sent 4 arguments in this order:
// 
// The Phaser.Physics.P2.Body it is in contact with. *This might be null* if the Body was created directly in the p2 world.
// The p2.Body this Body has ended contact with.
// The Shape from this body that caused the original contact.
// The Shape from the contact body.
func (self *PhysicsP2Body) SetOnEndContactA(member *Signal) {
    self.Object.Set("onEndContact", member)
}

// Array of CollisionGroups that this Bodies shapes collide with.
func (self *PhysicsP2Body) GetCollidesWithA() []interface{}{
	array00 := self.Object.Get("collidesWith")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Array of CollisionGroups that this Bodies shapes collide with.
func (self *PhysicsP2Body) SetCollidesWithA(member []interface{}) {
    self.Object.Set("collidesWith", member)
}

// To avoid deleting this body during a physics step, and causing all kinds of problems, set removeNextStep to true to have it removed in the next preUpdate.
func (self *PhysicsP2Body) GetRemoveNextStepA() bool{
    return self.Object.Get("removeNextStep").Bool()
}

// To avoid deleting this body during a physics step, and causing all kinds of problems, set removeNextStep to true to have it removed in the next preUpdate.
func (self *PhysicsP2Body) SetRemoveNextStepA(member bool) {
    self.Object.Set("removeNextStep", member)
}

// Reference to the debug body.
func (self *PhysicsP2Body) GetDebugBodyA() *PhysicsP2BodyDebug{
    return &PhysicsP2BodyDebug{self.Object.Get("debugBody")}
}

// Reference to the debug body.
func (self *PhysicsP2Body) SetDebugBodyA(member *PhysicsP2BodyDebug) {
    self.Object.Set("debugBody", member)
}

// Internally used by Sprite.x/y
func (self *PhysicsP2Body) GetDirtyA() bool{
    return self.Object.Get("dirty").Bool()
}

// Internally used by Sprite.x/y
func (self *PhysicsP2Body) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// Dynamic body. Dynamic bodies body can move and respond to collisions and forces.
func (self *PhysicsP2Body) GetDYNAMICA() int{
    return self.Object.Get("DYNAMIC").Int()
}

// Dynamic body. Dynamic bodies body can move and respond to collisions and forces.
func (self *PhysicsP2Body) SetDYNAMICA(member int) {
    self.Object.Set("DYNAMIC", member)
}

// Static body. Static bodies do not move, and they do not respond to forces or collision.
func (self *PhysicsP2Body) GetSTATICA() int{
    return self.Object.Get("STATIC").Int()
}

// Static body. Static bodies do not move, and they do not respond to forces or collision.
func (self *PhysicsP2Body) SetSTATICA(member int) {
    self.Object.Set("STATIC", member)
}

// Kinematic body. Kinematic bodies only moves according to its .velocity, and does not respond to collisions or force.
func (self *PhysicsP2Body) GetKINEMATICA() int{
    return self.Object.Get("KINEMATIC").Int()
}

// Kinematic body. Kinematic bodies only moves according to its .velocity, and does not respond to collisions or force.
func (self *PhysicsP2Body) SetKINEMATICA(member int) {
    self.Object.Set("KINEMATIC", member)
}

// Returns true if the Body is static. Setting Body.static to 'false' will make it dynamic.
func (self *PhysicsP2Body) GetStaticA() bool{
    return self.Object.Get("static").Bool()
}

// Returns true if the Body is static. Setting Body.static to 'false' will make it dynamic.
func (self *PhysicsP2Body) SetStaticA(member bool) {
    self.Object.Set("static", member)
}

// Returns true if the Body is dynamic. Setting Body.dynamic to 'false' will make it static.
func (self *PhysicsP2Body) GetDynamicA() bool{
    return self.Object.Get("dynamic").Bool()
}

// Returns true if the Body is dynamic. Setting Body.dynamic to 'false' will make it static.
func (self *PhysicsP2Body) SetDynamicA(member bool) {
    self.Object.Set("dynamic", member)
}

// Returns true if the Body is kinematic. Setting Body.kinematic to 'false' will make it static.
func (self *PhysicsP2Body) GetKinematicA() bool{
    return self.Object.Get("kinematic").Bool()
}

// Returns true if the Body is kinematic. Setting Body.kinematic to 'false' will make it static.
func (self *PhysicsP2Body) SetKinematicA(member bool) {
    self.Object.Set("kinematic", member)
}

// -
func (self *PhysicsP2Body) GetAllowSleepA() bool{
    return self.Object.Get("allowSleep").Bool()
}

// -
func (self *PhysicsP2Body) SetAllowSleepA(member bool) {
    self.Object.Set("allowSleep", member)
}

// The angle of the Body in degrees from its original orientation. Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. For example, the statement Body.angle = 450 is the same as Body.angle = 90.
// If you wish to work in radians instead of degrees use the property Body.rotation instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in degrees.
func (self *PhysicsP2Body) GetAngleA() int{
    return self.Object.Get("angle").Int()
}

// The angle of the Body in degrees from its original orientation. Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. For example, the statement Body.angle = 450 is the same as Body.angle = 90.
// If you wish to work in radians instead of degrees use the property Body.rotation instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in degrees.
func (self *PhysicsP2Body) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The angular damping acting acting on the body.
func (self *PhysicsP2Body) GetAngularDampingA() int{
    return self.Object.Get("angularDamping").Int()
}

// Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The angular damping acting acting on the body.
func (self *PhysicsP2Body) SetAngularDampingA(member int) {
    self.Object.Set("angularDamping", member)
}

// The angular force acting on the body.
func (self *PhysicsP2Body) GetAngularForceA() int{
    return self.Object.Get("angularForce").Int()
}

// The angular force acting on the body.
func (self *PhysicsP2Body) SetAngularForceA(member int) {
    self.Object.Set("angularForce", member)
}

// The angular velocity of the body.
func (self *PhysicsP2Body) GetAngularVelocityA() int{
    return self.Object.Get("angularVelocity").Int()
}

// The angular velocity of the body.
func (self *PhysicsP2Body) SetAngularVelocityA(member int) {
    self.Object.Set("angularVelocity", member)
}

// Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The linear damping acting on the body in the velocity direction.
func (self *PhysicsP2Body) GetDampingA() int{
    return self.Object.Get("damping").Int()
}

// Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The linear damping acting on the body in the velocity direction.
func (self *PhysicsP2Body) SetDampingA(member int) {
    self.Object.Set("damping", member)
}

// -
func (self *PhysicsP2Body) GetFixedRotationA() bool{
    return self.Object.Get("fixedRotation").Bool()
}

// -
func (self *PhysicsP2Body) SetFixedRotationA(member bool) {
    self.Object.Set("fixedRotation", member)
}

// The inertia of the body around the Z axis..
func (self *PhysicsP2Body) GetInertiaA() int{
    return self.Object.Get("inertia").Int()
}

// The inertia of the body around the Z axis..
func (self *PhysicsP2Body) SetInertiaA(member int) {
    self.Object.Set("inertia", member)
}

// The mass of the body.
func (self *PhysicsP2Body) GetMassA() int{
    return self.Object.Get("mass").Int()
}

// The mass of the body.
func (self *PhysicsP2Body) SetMassA(member int) {
    self.Object.Set("mass", member)
}

// The type of motion this body has. Should be one of: Body.STATIC (the body does not move), Body.DYNAMIC (body can move and respond to collisions) and Body.KINEMATIC (only moves according to its .velocity).
func (self *PhysicsP2Body) GetMotionStateA() int{
    return self.Object.Get("motionState").Int()
}

// The type of motion this body has. Should be one of: Body.STATIC (the body does not move), Body.DYNAMIC (body can move and respond to collisions) and Body.KINEMATIC (only moves according to its .velocity).
func (self *PhysicsP2Body) SetMotionStateA(member int) {
    self.Object.Set("motionState", member)
}

// The angle of the Body in radians.
// If you wish to work in degrees instead of radians use the Body.angle property instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in radians.
func (self *PhysicsP2Body) GetRotationA() int{
    return self.Object.Get("rotation").Int()
}

// The angle of the Body in radians.
// If you wish to work in degrees instead of radians use the Body.angle property instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in radians.
func (self *PhysicsP2Body) SetRotationA(member int) {
    self.Object.Set("rotation", member)
}

// .
func (self *PhysicsP2Body) GetSleepSpeedLimitA() int{
    return self.Object.Get("sleepSpeedLimit").Int()
}

// .
func (self *PhysicsP2Body) SetSleepSpeedLimitA(member int) {
    self.Object.Set("sleepSpeedLimit", member)
}

// The x coordinate of this Body.
func (self *PhysicsP2Body) GetXA() int{
    return self.Object.Get("x").Int()
}

// The x coordinate of this Body.
func (self *PhysicsP2Body) SetXA(member int) {
    self.Object.Set("x", member)
}

// The y coordinate of this Body.
func (self *PhysicsP2Body) GetYA() int{
    return self.Object.Get("y").Int()
}

// The y coordinate of this Body.
func (self *PhysicsP2Body) SetYA(member int) {
    self.Object.Set("y", member)
}

// The Body ID. Each Body that has been added to the World has a unique ID.
func (self *PhysicsP2Body) GetIdA() int{
    return self.Object.Get("id").Int()
}

// The Body ID. Each Body that has been added to the World has a unique ID.
func (self *PhysicsP2Body) SetIdA(member int) {
    self.Object.Set("id", member)
}

// Enable or disable debug drawing of this body
func (self *PhysicsP2Body) GetDebugA() bool{
    return self.Object.Get("debug").Bool()
}

// Enable or disable debug drawing of this body
func (self *PhysicsP2Body) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// A Body can be set to collide against the World bounds automatically if this is set to true. Otherwise it will leave the World.
// Note that this only applies if your World has bounds! The response to the collision should be managed via CollisionMaterials.
// Also note that when you set this it will only effect Body shapes that already exist. If you then add further shapes to your Body
// after setting this it will *not* proactively set them to collide with the bounds. Should the Body collide with the World bounds?
func (self *PhysicsP2Body) GetCollideWorldBoundsA() bool{
    return self.Object.Get("collideWorldBounds").Bool()
}

// A Body can be set to collide against the World bounds automatically if this is set to true. Otherwise it will leave the World.
// Note that this only applies if your World has bounds! The response to the collision should be managed via CollisionMaterials.
// Also note that when you set this it will only effect Body shapes that already exist. If you then add further shapes to your Body
// after setting this it will *not* proactively set them to collide with the bounds. Should the Body collide with the World bounds?
func (self *PhysicsP2Body) SetCollideWorldBoundsA(member bool) {
    self.Object.Set("collideWorldBounds", member)
}



// Sets a callback to be fired any time a shape in this Body impacts with a shape in the given Body. The impact test is performed against body.id values.
// The callback will be sent 4 parameters: This body, the body that impacted, the Shape in this body and the shape in the impacting body.
// Note that the impact event happens after collision resolution, so it cannot be used to prevent a collision from happening.
// It also happens mid-step. So do not destroy a Body during this callback, instead set safeDestroy to true so it will be killed on the next preUpdate.
func (self *PhysicsP2Body) CreateBodyCallback(object interface{}, callback func(...interface{}), callbackContext interface{}) {
    self.Object.Call("createBodyCallback", object, callback, callbackContext)
}

// Sets a callback to be fired any time a shape in this Body impacts with a shape in the given Body. The impact test is performed against body.id values.
// The callback will be sent 4 parameters: This body, the body that impacted, the Shape in this body and the shape in the impacting body.
// Note that the impact event happens after collision resolution, so it cannot be used to prevent a collision from happening.
// It also happens mid-step. So do not destroy a Body during this callback, instead set safeDestroy to true so it will be killed on the next preUpdate.
func (self *PhysicsP2Body) CreateBodyCallbackI(args ...interface{}) {
    self.Object.Call("createBodyCallback", args)
}

// Sets a callback to be fired any time this Body impacts with the given Group. The impact test is performed against shape.collisionGroup values.
// The callback will be sent 4 parameters: This body, the body that impacted, the Shape in this body and the shape in the impacting body.
// This callback will only fire if this Body has been assigned a collision group.
// Note that the impact event happens after collision resolution, so it cannot be used to prevent a collision from happening.
// It also happens mid-step. So do not destroy a Body during this callback, instead set safeDestroy to true so it will be killed on the next preUpdate.
func (self *PhysicsP2Body) CreateGroupCallback(group *PhysicsCollisionGroup, callback func(...interface{}), callbackContext interface{}) {
    self.Object.Call("createGroupCallback", group, callback, callbackContext)
}

// Sets a callback to be fired any time this Body impacts with the given Group. The impact test is performed against shape.collisionGroup values.
// The callback will be sent 4 parameters: This body, the body that impacted, the Shape in this body and the shape in the impacting body.
// This callback will only fire if this Body has been assigned a collision group.
// Note that the impact event happens after collision resolution, so it cannot be used to prevent a collision from happening.
// It also happens mid-step. So do not destroy a Body during this callback, instead set safeDestroy to true so it will be killed on the next preUpdate.
func (self *PhysicsP2Body) CreateGroupCallbackI(args ...interface{}) {
    self.Object.Call("createGroupCallback", args)
}

// Gets the collision bitmask from the groups this body collides with.
func (self *PhysicsP2Body) GetCollisionMask() int{
    return self.Object.Call("getCollisionMask").Int()
}

// Gets the collision bitmask from the groups this body collides with.
func (self *PhysicsP2Body) GetCollisionMaskI(args ...interface{}) int{
    return self.Object.Call("getCollisionMask", args).Int()
}

// Updates the collisionMask.
func (self *PhysicsP2Body) UpdateCollisionMask() {
    self.Object.Call("updateCollisionMask")
}

// Updates the collisionMask.
func (self *PhysicsP2Body) UpdateCollisionMask1O(shape *P2Shape) {
    self.Object.Call("updateCollisionMask", shape)
}

// Updates the collisionMask.
func (self *PhysicsP2Body) UpdateCollisionMaskI(args ...interface{}) {
    self.Object.Call("updateCollisionMask", args)
}

// Sets the given CollisionGroup to be the collision group for all shapes in this Body, unless a shape is specified.
// This also resets the collisionMask.
func (self *PhysicsP2Body) SetCollisionGroup(group *PhysicsCollisionGroup) {
    self.Object.Call("setCollisionGroup", group)
}

// Sets the given CollisionGroup to be the collision group for all shapes in this Body, unless a shape is specified.
// This also resets the collisionMask.
func (self *PhysicsP2Body) SetCollisionGroup1O(group *PhysicsCollisionGroup, shape *P2Shape) {
    self.Object.Call("setCollisionGroup", group, shape)
}

// Sets the given CollisionGroup to be the collision group for all shapes in this Body, unless a shape is specified.
// This also resets the collisionMask.
func (self *PhysicsP2Body) SetCollisionGroupI(args ...interface{}) {
    self.Object.Call("setCollisionGroup", args)
}

// Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollision() {
    self.Object.Call("clearCollision")
}

// Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollision1O(clearGroup bool) {
    self.Object.Call("clearCollision", clearGroup)
}

// Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollision2O(clearGroup bool, clearMask bool) {
    self.Object.Call("clearCollision", clearGroup, clearMask)
}

// Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollision3O(clearGroup bool, clearMask bool, shape *P2Shape) {
    self.Object.Call("clearCollision", clearGroup, clearMask, shape)
}

// Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollisionI(args ...interface{}) {
    self.Object.Call("clearCollision", args)
}

// Removes the given CollisionGroup, or array of CollisionGroups, from the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) RemoveCollisionGroup(group interface{}) {
    self.Object.Call("removeCollisionGroup", group)
}

// Removes the given CollisionGroup, or array of CollisionGroups, from the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) RemoveCollisionGroup1O(group interface{}, clearCallback bool) {
    self.Object.Call("removeCollisionGroup", group, clearCallback)
}

// Removes the given CollisionGroup, or array of CollisionGroups, from the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) RemoveCollisionGroup2O(group interface{}, clearCallback bool, shape *P2Shape) {
    self.Object.Call("removeCollisionGroup", group, clearCallback, shape)
}

// Removes the given CollisionGroup, or array of CollisionGroups, from the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) RemoveCollisionGroupI(args ...interface{}) {
    self.Object.Call("removeCollisionGroup", args)
}

// Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) Collides(group interface{}) {
    self.Object.Call("collides", group)
}

// Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) Collides1O(group interface{}, callback func(...interface{})) {
    self.Object.Call("collides", group, callback)
}

// Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) Collides2O(group interface{}, callback func(...interface{}), callbackContext interface{}) {
    self.Object.Call("collides", group, callback, callbackContext)
}

// Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) Collides3O(group interface{}, callback func(...interface{}), callbackContext interface{}, shape *P2Shape) {
    self.Object.Call("collides", group, callback, callbackContext, shape)
}

// Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) CollidesI(args ...interface{}) {
    self.Object.Call("collides", args)
}

// Moves the shape offsets so their center of mass becomes the body center of mass.
func (self *PhysicsP2Body) AdjustCenterOfMass() {
    self.Object.Call("adjustCenterOfMass")
}

// Moves the shape offsets so their center of mass becomes the body center of mass.
func (self *PhysicsP2Body) AdjustCenterOfMassI(args ...interface{}) {
    self.Object.Call("adjustCenterOfMass", args)
}

// Gets the velocity of a point in the body.
func (self *PhysicsP2Body) GetVelocityAtPoint(result []interface{}, relativePoint []interface{}) []interface{}{
	array00 := self.Object.Call("getVelocityAtPoint", result, relativePoint)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Gets the velocity of a point in the body.
func (self *PhysicsP2Body) GetVelocityAtPointI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("getVelocityAtPoint", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Apply damping, see http://code.google.com/p/bullet/issues/detail?id=74 for details.
func (self *PhysicsP2Body) ApplyDamping(dt int) {
    self.Object.Call("applyDamping", dt)
}

// Apply damping, see http://code.google.com/p/bullet/issues/detail?id=74 for details.
func (self *PhysicsP2Body) ApplyDampingI(args ...interface{}) {
    self.Object.Call("applyDamping", args)
}

// Apply impulse to a point relative to the body.
// This could for example be a point on the Body surface. An impulse is a force added to a body during a short 
// period of time (impulse = force * time). Impulses will be added to Body.velocity and Body.angularVelocity.
func (self *PhysicsP2Body) ApplyImpulse(impulse interface{}, worldX int, worldY int) {
    self.Object.Call("applyImpulse", impulse, worldX, worldY)
}

// Apply impulse to a point relative to the body.
// This could for example be a point on the Body surface. An impulse is a force added to a body during a short 
// period of time (impulse = force * time). Impulses will be added to Body.velocity and Body.angularVelocity.
func (self *PhysicsP2Body) ApplyImpulseI(args ...interface{}) {
    self.Object.Call("applyImpulse", args)
}

// Apply impulse to a point local to the body.
// 
// This could for example be a point on the Body surface. An impulse is a force added to a body during a short 
// period of time (impulse = force * time). Impulses will be added to Body.velocity and Body.angularVelocity.
func (self *PhysicsP2Body) ApplyImpulseLocal(impulse interface{}, localX int, localY int) {
    self.Object.Call("applyImpulseLocal", impulse, localX, localY)
}

// Apply impulse to a point local to the body.
// 
// This could for example be a point on the Body surface. An impulse is a force added to a body during a short 
// period of time (impulse = force * time). Impulses will be added to Body.velocity and Body.angularVelocity.
func (self *PhysicsP2Body) ApplyImpulseLocalI(args ...interface{}) {
    self.Object.Call("applyImpulseLocal", args)
}

// Apply force to a world point.
// 
// This could for example be a point on the RigidBody surface. Applying force 
// this way will add to Body.force and Body.angularForce.
func (self *PhysicsP2Body) ApplyForce(force interface{}, worldX int, worldY int) {
    self.Object.Call("applyForce", force, worldX, worldY)
}

// Apply force to a world point.
// 
// This could for example be a point on the RigidBody surface. Applying force 
// this way will add to Body.force and Body.angularForce.
func (self *PhysicsP2Body) ApplyForceI(args ...interface{}) {
    self.Object.Call("applyForce", args)
}

// Sets the force on the body to zero.
func (self *PhysicsP2Body) SetZeroForce() {
    self.Object.Call("setZeroForce")
}

// Sets the force on the body to zero.
func (self *PhysicsP2Body) SetZeroForceI(args ...interface{}) {
    self.Object.Call("setZeroForce", args)
}

// If this Body is dynamic then this will zero its angular velocity.
func (self *PhysicsP2Body) SetZeroRotation() {
    self.Object.Call("setZeroRotation")
}

// If this Body is dynamic then this will zero its angular velocity.
func (self *PhysicsP2Body) SetZeroRotationI(args ...interface{}) {
    self.Object.Call("setZeroRotation", args)
}

// If this Body is dynamic then this will zero its velocity on both axis.
func (self *PhysicsP2Body) SetZeroVelocity() {
    self.Object.Call("setZeroVelocity")
}

// If this Body is dynamic then this will zero its velocity on both axis.
func (self *PhysicsP2Body) SetZeroVelocityI(args ...interface{}) {
    self.Object.Call("setZeroVelocity", args)
}

// Sets the Body damping and angularDamping to zero.
func (self *PhysicsP2Body) SetZeroDamping() {
    self.Object.Call("setZeroDamping")
}

// Sets the Body damping and angularDamping to zero.
func (self *PhysicsP2Body) SetZeroDampingI(args ...interface{}) {
    self.Object.Call("setZeroDamping", args)
}

// Transform a world point to local body frame.
func (self *PhysicsP2Body) ToLocalFrame(out interface{}, worldPoint interface{}) {
    self.Object.Call("toLocalFrame", out, worldPoint)
}

// Transform a world point to local body frame.
func (self *PhysicsP2Body) ToLocalFrameI(args ...interface{}) {
    self.Object.Call("toLocalFrame", args)
}

// Transform a local point to world frame.
func (self *PhysicsP2Body) ToWorldFrame(out []interface{}, localPoint []interface{}) {
    self.Object.Call("toWorldFrame", out, localPoint)
}

// Transform a local point to world frame.
func (self *PhysicsP2Body) ToWorldFrameI(args ...interface{}) {
    self.Object.Call("toWorldFrame", args)
}

// This will rotate the Body by the given speed to the left (counter-clockwise).
func (self *PhysicsP2Body) RotateLeft(speed int) {
    self.Object.Call("rotateLeft", speed)
}

// This will rotate the Body by the given speed to the left (counter-clockwise).
func (self *PhysicsP2Body) RotateLeftI(args ...interface{}) {
    self.Object.Call("rotateLeft", args)
}

// This will rotate the Body by the given speed to the left (clockwise).
func (self *PhysicsP2Body) RotateRight(speed int) {
    self.Object.Call("rotateRight", speed)
}

// This will rotate the Body by the given speed to the left (clockwise).
func (self *PhysicsP2Body) RotateRightI(args ...interface{}) {
    self.Object.Call("rotateRight", args)
}

// Moves the Body forwards based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveForward(speed int) {
    self.Object.Call("moveForward", speed)
}

// Moves the Body forwards based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveForwardI(args ...interface{}) {
    self.Object.Call("moveForward", args)
}

// Moves the Body backwards based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveBackward(speed int) {
    self.Object.Call("moveBackward", speed)
}

// Moves the Body backwards based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveBackwardI(args ...interface{}) {
    self.Object.Call("moveBackward", args)
}

// Applies a force to the Body that causes it to 'thrust' forwards, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) Thrust(speed int) {
    self.Object.Call("thrust", speed)
}

// Applies a force to the Body that causes it to 'thrust' forwards, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustI(args ...interface{}) {
    self.Object.Call("thrust", args)
}

// Applies a force to the Body that causes it to 'thrust' to the left, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustLeft(speed int) {
    self.Object.Call("thrustLeft", speed)
}

// Applies a force to the Body that causes it to 'thrust' to the left, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustLeftI(args ...interface{}) {
    self.Object.Call("thrustLeft", args)
}

// Applies a force to the Body that causes it to 'thrust' to the right, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustRight(speed int) {
    self.Object.Call("thrustRight", speed)
}

// Applies a force to the Body that causes it to 'thrust' to the right, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustRightI(args ...interface{}) {
    self.Object.Call("thrustRight", args)
}

// Applies a force to the Body that causes it to 'thrust' backwards (in reverse), based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) Reverse(speed int) {
    self.Object.Call("reverse", speed)
}

// Applies a force to the Body that causes it to 'thrust' backwards (in reverse), based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ReverseI(args ...interface{}) {
    self.Object.Call("reverse", args)
}

// If this Body is dynamic then this will move it to the left by setting its x velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveLeft(speed int) {
    self.Object.Call("moveLeft", speed)
}

// If this Body is dynamic then this will move it to the left by setting its x velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveLeftI(args ...interface{}) {
    self.Object.Call("moveLeft", args)
}

// If this Body is dynamic then this will move it to the right by setting its x velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveRight(speed int) {
    self.Object.Call("moveRight", speed)
}

// If this Body is dynamic then this will move it to the right by setting its x velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveRightI(args ...interface{}) {
    self.Object.Call("moveRight", args)
}

// If this Body is dynamic then this will move it up by setting its y velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveUp(speed int) {
    self.Object.Call("moveUp", speed)
}

// If this Body is dynamic then this will move it up by setting its y velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveUpI(args ...interface{}) {
    self.Object.Call("moveUp", args)
}

// If this Body is dynamic then this will move it down by setting its y velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveDown(speed int) {
    self.Object.Call("moveDown", speed)
}

// If this Body is dynamic then this will move it down by setting its y velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveDownI(args ...interface{}) {
    self.Object.Call("moveDown", args)
}

// Internal method. This is called directly before the sprites are sent to the renderer and after the update function has finished.
func (self *PhysicsP2Body) PreUpdate() {
    self.Object.Call("preUpdate")
}

// Internal method. This is called directly before the sprites are sent to the renderer and after the update function has finished.
func (self *PhysicsP2Body) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Internal method. This is called directly before the sprites are sent to the renderer and after the update function has finished.
func (self *PhysicsP2Body) PostUpdate() {
    self.Object.Call("postUpdate")
}

// Internal method. This is called directly before the sprites are sent to the renderer and after the update function has finished.
func (self *PhysicsP2Body) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Resets the Body force, velocity (linear and angular) and rotation. Optionally resets damping and mass.
func (self *PhysicsP2Body) Reset(x int, y int) {
    self.Object.Call("reset", x, y)
}

// Resets the Body force, velocity (linear and angular) and rotation. Optionally resets damping and mass.
func (self *PhysicsP2Body) Reset1O(x int, y int, resetDamping bool) {
    self.Object.Call("reset", x, y, resetDamping)
}

// Resets the Body force, velocity (linear and angular) and rotation. Optionally resets damping and mass.
func (self *PhysicsP2Body) Reset2O(x int, y int, resetDamping bool, resetMass bool) {
    self.Object.Call("reset", x, y, resetDamping, resetMass)
}

// Resets the Body force, velocity (linear and angular) and rotation. Optionally resets damping and mass.
func (self *PhysicsP2Body) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Adds this physics body to the world.
func (self *PhysicsP2Body) AddToWorld() {
    self.Object.Call("addToWorld")
}

// Adds this physics body to the world.
func (self *PhysicsP2Body) AddToWorldI(args ...interface{}) {
    self.Object.Call("addToWorld", args)
}

// Removes this physics body from the world.
func (self *PhysicsP2Body) RemoveFromWorld() {
    self.Object.Call("removeFromWorld")
}

// Removes this physics body from the world.
func (self *PhysicsP2Body) RemoveFromWorldI(args ...interface{}) {
    self.Object.Call("removeFromWorld", args)
}

// Destroys this Body and all references it holds to other objects.
func (self *PhysicsP2Body) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this Body and all references it holds to other objects.
func (self *PhysicsP2Body) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Removes all Shapes from this Body.
func (self *PhysicsP2Body) ClearShapes() {
    self.Object.Call("clearShapes")
}

// Removes all Shapes from this Body.
func (self *PhysicsP2Body) ClearShapesI(args ...interface{}) {
    self.Object.Call("clearShapes", args)
}

// Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShape(shape *P2Shape) *P2Shape{
    return &P2Shape{self.Object.Call("addShape", shape)}
}

// Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShape1O(shape *P2Shape, offsetX int) *P2Shape{
    return &P2Shape{self.Object.Call("addShape", shape, offsetX)}
}

// Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShape2O(shape *P2Shape, offsetX int, offsetY int) *P2Shape{
    return &P2Shape{self.Object.Call("addShape", shape, offsetX, offsetY)}
}

// Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShape3O(shape *P2Shape, offsetX int, offsetY int, rotation int) *P2Shape{
    return &P2Shape{self.Object.Call("addShape", shape, offsetX, offsetY, rotation)}
}

// Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShapeI(args ...interface{}) *P2Shape{
    return &P2Shape{self.Object.Call("addShape", args)}
}

// Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircle(radius int) *P2Circle{
    return &P2Circle{self.Object.Call("addCircle", radius)}
}

// Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircle1O(radius int, offsetX int) *P2Circle{
    return &P2Circle{self.Object.Call("addCircle", radius, offsetX)}
}

// Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircle2O(radius int, offsetX int, offsetY int) *P2Circle{
    return &P2Circle{self.Object.Call("addCircle", radius, offsetX, offsetY)}
}

// Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircle3O(radius int, offsetX int, offsetY int, rotation int) *P2Circle{
    return &P2Circle{self.Object.Call("addCircle", radius, offsetX, offsetY, rotation)}
}

// Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircleI(args ...interface{}) *P2Circle{
    return &P2Circle{self.Object.Call("addCircle", args)}
}

// Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangle(width int, height int) *P2Box{
    return &P2Box{self.Object.Call("addRectangle", width, height)}
}

// Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangle1O(width int, height int, offsetX int) *P2Box{
    return &P2Box{self.Object.Call("addRectangle", width, height, offsetX)}
}

// Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangle2O(width int, height int, offsetX int, offsetY int) *P2Box{
    return &P2Box{self.Object.Call("addRectangle", width, height, offsetX, offsetY)}
}

// Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangle3O(width int, height int, offsetX int, offsetY int, rotation int) *P2Box{
    return &P2Box{self.Object.Call("addRectangle", width, height, offsetX, offsetY, rotation)}
}

// Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangleI(args ...interface{}) *P2Box{
    return &P2Box{self.Object.Call("addRectangle", args)}
}

// Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlane() *P2Plane{
    return &P2Plane{self.Object.Call("addPlane")}
}

// Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlane1O(offsetX int) *P2Plane{
    return &P2Plane{self.Object.Call("addPlane", offsetX)}
}

// Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlane2O(offsetX int, offsetY int) *P2Plane{
    return &P2Plane{self.Object.Call("addPlane", offsetX, offsetY)}
}

// Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlane3O(offsetX int, offsetY int, rotation int) *P2Plane{
    return &P2Plane{self.Object.Call("addPlane", offsetX, offsetY, rotation)}
}

// Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlaneI(args ...interface{}) *P2Plane{
    return &P2Plane{self.Object.Call("addPlane", args)}
}

// Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticle() *P2Particle{
    return &P2Particle{self.Object.Call("addParticle")}
}

// Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticle1O(offsetX int) *P2Particle{
    return &P2Particle{self.Object.Call("addParticle", offsetX)}
}

// Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticle2O(offsetX int, offsetY int) *P2Particle{
    return &P2Particle{self.Object.Call("addParticle", offsetX, offsetY)}
}

// Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticle3O(offsetX int, offsetY int, rotation int) *P2Particle{
    return &P2Particle{self.Object.Call("addParticle", offsetX, offsetY, rotation)}
}

// Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticleI(args ...interface{}) *P2Particle{
    return &P2Particle{self.Object.Call("addParticle", args)}
}

// Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLine(length int) *P2Line{
    return &P2Line{self.Object.Call("addLine", length)}
}

// Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLine1O(length int, offsetX int) *P2Line{
    return &P2Line{self.Object.Call("addLine", length, offsetX)}
}

// Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLine2O(length int, offsetX int, offsetY int) *P2Line{
    return &P2Line{self.Object.Call("addLine", length, offsetX, offsetY)}
}

// Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLine3O(length int, offsetX int, offsetY int, rotation int) *P2Line{
    return &P2Line{self.Object.Call("addLine", length, offsetX, offsetY, rotation)}
}

// Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLineI(args ...interface{}) *P2Line{
    return &P2Line{self.Object.Call("addLine", args)}
}

// Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsule(length int, radius int) *P2Capsule{
    return &P2Capsule{self.Object.Call("addCapsule", length, radius)}
}

// Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsule1O(length int, radius int, offsetX int) *P2Capsule{
    return &P2Capsule{self.Object.Call("addCapsule", length, radius, offsetX)}
}

// Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsule2O(length int, radius int, offsetX int, offsetY int) *P2Capsule{
    return &P2Capsule{self.Object.Call("addCapsule", length, radius, offsetX, offsetY)}
}

// Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsule3O(length int, radius int, offsetX int, offsetY int, rotation int) *P2Capsule{
    return &P2Capsule{self.Object.Call("addCapsule", length, radius, offsetX, offsetY, rotation)}
}

// Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsuleI(args ...interface{}) *P2Capsule{
    return &P2Capsule{self.Object.Call("addCapsule", args)}
}

// Reads a polygon shape path, and assembles convex shapes from that and puts them at proper offset points. The shape must be simple and without holes.
// This function expects the x.y values to be given in pixels. If you want to provide them at p2 world scales then call Body.data.fromPolygon directly.
func (self *PhysicsP2Body) AddPolygon(options interface{}, points interface{}) bool{
    return self.Object.Call("addPolygon", options, points).Bool()
}

// Reads a polygon shape path, and assembles convex shapes from that and puts them at proper offset points. The shape must be simple and without holes.
// This function expects the x.y values to be given in pixels. If you want to provide them at p2 world scales then call Body.data.fromPolygon directly.
func (self *PhysicsP2Body) AddPolygonI(args ...interface{}) bool{
    return self.Object.Call("addPolygon", args).Bool()
}

// Remove a shape from the body. Will automatically update the mass properties and bounding radius.
func (self *PhysicsP2Body) RemoveShape(shape interface{}) bool{
    return self.Object.Call("removeShape", shape).Bool()
}

// Remove a shape from the body. Will automatically update the mass properties and bounding radius.
func (self *PhysicsP2Body) RemoveShapeI(args ...interface{}) bool{
    return self.Object.Call("removeShape", args).Bool()
}

// Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircle(radius int) {
    self.Object.Call("setCircle", radius)
}

// Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircle1O(radius int, offsetX int) {
    self.Object.Call("setCircle", radius, offsetX)
}

// Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircle2O(radius int, offsetX int, offsetY int) {
    self.Object.Call("setCircle", radius, offsetX, offsetY)
}

// Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircle3O(radius int, offsetX int, offsetY int, rotation int) {
    self.Object.Call("setCircle", radius, offsetX, offsetY, rotation)
}

// Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircleI(args ...interface{}) {
    self.Object.Call("setCircle", args)
}

// Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle() *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle")}
}

// Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle1O(width int) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", width)}
}

// Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle2O(width int, height int) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", width, height)}
}

// Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle3O(width int, height int, offsetX int) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", width, height, offsetX)}
}

// Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle4O(width int, height int, offsetX int, offsetY int) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", width, height, offsetX, offsetY)}
}

// Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle5O(width int, height int, offsetX int, offsetY int, rotation int) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", width, height, offsetX, offsetY, rotation)}
}

// Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangleI(args ...interface{}) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", args)}
}

// Clears any previously set shapes.
// Then creates a Rectangle shape sized to match the dimensions and orientation of the Sprite given.
// If no Sprite is given it defaults to using the parent of this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangleFromSprite() *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangleFromSprite")}
}

// Clears any previously set shapes.
// Then creates a Rectangle shape sized to match the dimensions and orientation of the Sprite given.
// If no Sprite is given it defaults to using the parent of this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangleFromSprite1O(sprite interface{}) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangleFromSprite", sprite)}
}

// Clears any previously set shapes.
// Then creates a Rectangle shape sized to match the dimensions and orientation of the Sprite given.
// If no Sprite is given it defaults to using the parent of this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangleFromSpriteI(args ...interface{}) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangleFromSprite", args)}
}

// Adds the given Material to all Shapes that belong to this Body.
// If you only wish to apply it to a specific Shape in this Body then provide that as the 2nd parameter.
func (self *PhysicsP2Body) SetMaterial(material *PhysicsP2Material) {
    self.Object.Call("setMaterial", material)
}

// Adds the given Material to all Shapes that belong to this Body.
// If you only wish to apply it to a specific Shape in this Body then provide that as the 2nd parameter.
func (self *PhysicsP2Body) SetMaterial1O(material *PhysicsP2Material, shape *P2Shape) {
    self.Object.Call("setMaterial", material, shape)
}

// Adds the given Material to all Shapes that belong to this Body.
// If you only wish to apply it to a specific Shape in this Body then provide that as the 2nd parameter.
func (self *PhysicsP2Body) SetMaterialI(args ...interface{}) {
    self.Object.Call("setMaterial", args)
}

// Updates the debug draw if any body shapes change.
func (self *PhysicsP2Body) ShapeChanged() {
    self.Object.Call("shapeChanged")
}

// Updates the debug draw if any body shapes change.
func (self *PhysicsP2Body) ShapeChangedI(args ...interface{}) {
    self.Object.Call("shapeChanged", args)
}

// Reads the shape data from a physics data file stored in the Game.Cache and adds it as a polygon to this Body.
// The shape data format is based on the output of the
// {@link https://github.com/photonstorm/phaser/tree/master/resources/PhysicsEditor%20Exporter|custom phaser exporter} for
// {@link https://www.codeandweb.com/physicseditor|PhysicsEditor}
func (self *PhysicsP2Body) AddPhaserPolygon(key string, object string) []interface{}{
	array00 := self.Object.Call("addPhaserPolygon", key, object)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Reads the shape data from a physics data file stored in the Game.Cache and adds it as a polygon to this Body.
// The shape data format is based on the output of the
// {@link https://github.com/photonstorm/phaser/tree/master/resources/PhysicsEditor%20Exporter|custom phaser exporter} for
// {@link https://www.codeandweb.com/physicseditor|PhysicsEditor}
func (self *PhysicsP2Body) AddPhaserPolygonI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("addPhaserPolygon", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Add a polygon fixture. This is used during #loadPolygon.
func (self *PhysicsP2Body) AddFixture(fixtureData string) []interface{}{
	array00 := self.Object.Call("addFixture", fixtureData)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Add a polygon fixture. This is used during #loadPolygon.
func (self *PhysicsP2Body) AddFixtureI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("addFixture", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Reads the shape data from a physics data file stored in the Game.Cache and adds it as a polygon to this Body.
// 
// As well as reading the data from the Cache you can also pass `null` as the first argument and a
// physics data object as the second. When doing this you must ensure the structure of the object is correct in advance.
// 
// For more details see the format of the Lime / Corona Physics Editor export.
func (self *PhysicsP2Body) LoadPolygon(key string, object interface{}) bool{
    return self.Object.Call("loadPolygon", key, object).Bool()
}

// Reads the shape data from a physics data file stored in the Game.Cache and adds it as a polygon to this Body.
// 
// As well as reading the data from the Cache you can also pass `null` as the first argument and a
// physics data object as the second. When doing this you must ensure the structure of the object is correct in advance.
// 
// For more details see the format of the Lime / Corona Physics Editor export.
func (self *PhysicsP2Body) LoadPolygonI(args ...interface{}) bool{
    return self.Object.Call("loadPolygon", args).Bool()
}
