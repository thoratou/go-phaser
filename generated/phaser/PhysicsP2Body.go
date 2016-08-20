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
func (self *PhysicsP2Body) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2Body) SetGame(member *Game) {
    self.Set("game", member)
}

// Local reference to the P2 World.
func (self *PhysicsP2Body) GetWorld() *PhysicsP2{
    return &PhysicsP2{self.Get("world")}
}

// Local reference to the P2 World.
func (self *PhysicsP2Body) SetWorld(member *PhysicsP2) {
    self.Set("world", member)
}

// Reference to the parent Sprite.
func (self *PhysicsP2Body) GetSprite() *Sprite{
    return &Sprite{self.Get("sprite")}
}

// Reference to the parent Sprite.
func (self *PhysicsP2Body) SetSprite(member *Sprite) {
    self.Set("sprite", member)
}

// The type of physics system this body belongs to.
func (self *PhysicsP2Body) GetType() float64{
    return self.Get("type").Float()
}

// The type of physics system this body belongs to.
func (self *PhysicsP2Body) SetType(member float64) {
    self.Set("type", member)
}

// The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsP2Body) GetOffset() *Point{
    return &Point{self.Get("offset")}
}

// The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsP2Body) SetOffset(member *Point) {
    self.Set("offset", member)
}

// The p2 Body data.
func (self *PhysicsP2Body) GetData() *P2Body{
    return &P2Body{self.Get("data")}
}

// The p2 Body data.
func (self *PhysicsP2Body) SetData(member *P2Body) {
    self.Set("data", member)
}

// The velocity of the body. Set velocity.x to a negative value to move to the left, position to the right. velocity.y negative values move up, positive move down.
func (self *PhysicsP2Body) GetVelocity() *PhysicsP2InversePointProxy{
    return &PhysicsP2InversePointProxy{self.Get("velocity")}
}

// The velocity of the body. Set velocity.x to a negative value to move to the left, position to the right. velocity.y negative values move up, positive move down.
func (self *PhysicsP2Body) SetVelocity(member *PhysicsP2InversePointProxy) {
    self.Set("velocity", member)
}

// The force applied to the body.
func (self *PhysicsP2Body) GetForce() *PhysicsP2InversePointProxy{
    return &PhysicsP2InversePointProxy{self.Get("force")}
}

// The force applied to the body.
func (self *PhysicsP2Body) SetForce(member *PhysicsP2InversePointProxy) {
    self.Set("force", member)
}

// A locally applied gravity force to the Body. Applied directly before the world step. NOTE: Not currently implemented.
func (self *PhysicsP2Body) GetGravity() *Point{
    return &Point{self.Get("gravity")}
}

// A locally applied gravity force to the Body. Applied directly before the world step. NOTE: Not currently implemented.
func (self *PhysicsP2Body) SetGravity(member *Point) {
    self.Set("gravity", member)
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
func (self *PhysicsP2Body) GetOnBeginContact() *Signal{
    return &Signal{self.Get("onBeginContact")}
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
func (self *PhysicsP2Body) SetOnBeginContact(member *Signal) {
    self.Set("onBeginContact", member)
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
func (self *PhysicsP2Body) GetOnEndContact() *Signal{
    return &Signal{self.Get("onEndContact")}
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
func (self *PhysicsP2Body) SetOnEndContact(member *Signal) {
    self.Set("onEndContact", member)
}

// Array of CollisionGroups that this Bodies shapes collide with.
func (self *PhysicsP2Body) GetCollidesWith() []interface{}{
	array := self.Get("collidesWith")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Array of CollisionGroups that this Bodies shapes collide with.
func (self *PhysicsP2Body) SetCollidesWith(member []interface{}) {
    self.Set("collidesWith", member)
}

// To avoid deleting this body during a physics step, and causing all kinds of problems, set removeNextStep to true to have it removed in the next preUpdate.
func (self *PhysicsP2Body) GetRemoveNextStep() bool{
    return self.Get("removeNextStep").Bool()
}

// To avoid deleting this body during a physics step, and causing all kinds of problems, set removeNextStep to true to have it removed in the next preUpdate.
func (self *PhysicsP2Body) SetRemoveNextStep(member bool) {
    self.Set("removeNextStep", member)
}

// Reference to the debug body.
func (self *PhysicsP2Body) GetDebugBody() *PhysicsP2BodyDebug{
    return &PhysicsP2BodyDebug{self.Get("debugBody")}
}

// Reference to the debug body.
func (self *PhysicsP2Body) SetDebugBody(member *PhysicsP2BodyDebug) {
    self.Set("debugBody", member)
}

// Internally used by Sprite.x/y
func (self *PhysicsP2Body) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// Internally used by Sprite.x/y
func (self *PhysicsP2Body) SetDirty(member bool) {
    self.Set("dirty", member)
}

// Dynamic body. Dynamic bodies body can move and respond to collisions and forces.
func (self *PhysicsP2Body) GetDYNAMIC() float64{
    return self.Get("DYNAMIC").Float()
}

// Dynamic body. Dynamic bodies body can move and respond to collisions and forces.
func (self *PhysicsP2Body) SetDYNAMIC(member float64) {
    self.Set("DYNAMIC", member)
}

// Static body. Static bodies do not move, and they do not respond to forces or collision.
func (self *PhysicsP2Body) GetSTATIC() float64{
    return self.Get("STATIC").Float()
}

// Static body. Static bodies do not move, and they do not respond to forces or collision.
func (self *PhysicsP2Body) SetSTATIC(member float64) {
    self.Set("STATIC", member)
}

// Kinematic body. Kinematic bodies only moves according to its .velocity, and does not respond to collisions or force.
func (self *PhysicsP2Body) GetKINEMATIC() float64{
    return self.Get("KINEMATIC").Float()
}

// Kinematic body. Kinematic bodies only moves according to its .velocity, and does not respond to collisions or force.
func (self *PhysicsP2Body) SetKINEMATIC(member float64) {
    self.Set("KINEMATIC", member)
}

// Returns true if the Body is static. Setting Body.static to 'false' will make it dynamic.
func (self *PhysicsP2Body) GetStatic() bool{
    return self.Get("static").Bool()
}

// Returns true if the Body is static. Setting Body.static to 'false' will make it dynamic.
func (self *PhysicsP2Body) SetStatic(member bool) {
    self.Set("static", member)
}

// Returns true if the Body is dynamic. Setting Body.dynamic to 'false' will make it static.
func (self *PhysicsP2Body) GetDynamic() bool{
    return self.Get("dynamic").Bool()
}

// Returns true if the Body is dynamic. Setting Body.dynamic to 'false' will make it static.
func (self *PhysicsP2Body) SetDynamic(member bool) {
    self.Set("dynamic", member)
}

// Returns true if the Body is kinematic. Setting Body.kinematic to 'false' will make it static.
func (self *PhysicsP2Body) GetKinematic() bool{
    return self.Get("kinematic").Bool()
}

// Returns true if the Body is kinematic. Setting Body.kinematic to 'false' will make it static.
func (self *PhysicsP2Body) SetKinematic(member bool) {
    self.Set("kinematic", member)
}

// -
func (self *PhysicsP2Body) GetAllowSleep() bool{
    return self.Get("allowSleep").Bool()
}

// -
func (self *PhysicsP2Body) SetAllowSleep(member bool) {
    self.Set("allowSleep", member)
}

// The angle of the Body in degrees from its original orientation. Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. For example, the statement Body.angle = 450 is the same as Body.angle = 90.
// If you wish to work in radians instead of degrees use the property Body.rotation instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in degrees.
func (self *PhysicsP2Body) GetAngle() float64{
    return self.Get("angle").Float()
}

// The angle of the Body in degrees from its original orientation. Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. For example, the statement Body.angle = 450 is the same as Body.angle = 90.
// If you wish to work in radians instead of degrees use the property Body.rotation instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in degrees.
func (self *PhysicsP2Body) SetAngle(member float64) {
    self.Set("angle", member)
}

// Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The angular damping acting acting on the body.
func (self *PhysicsP2Body) GetAngularDamping() float64{
    return self.Get("angularDamping").Float()
}

// Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The angular damping acting acting on the body.
func (self *PhysicsP2Body) SetAngularDamping(member float64) {
    self.Set("angularDamping", member)
}

// The angular force acting on the body.
func (self *PhysicsP2Body) GetAngularForce() float64{
    return self.Get("angularForce").Float()
}

// The angular force acting on the body.
func (self *PhysicsP2Body) SetAngularForce(member float64) {
    self.Set("angularForce", member)
}

// The angular velocity of the body.
func (self *PhysicsP2Body) GetAngularVelocity() float64{
    return self.Get("angularVelocity").Float()
}

// The angular velocity of the body.
func (self *PhysicsP2Body) SetAngularVelocity(member float64) {
    self.Set("angularVelocity", member)
}

// Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The linear damping acting on the body in the velocity direction.
func (self *PhysicsP2Body) GetDamping() float64{
    return self.Get("damping").Float()
}

// Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The linear damping acting on the body in the velocity direction.
func (self *PhysicsP2Body) SetDamping(member float64) {
    self.Set("damping", member)
}

// -
func (self *PhysicsP2Body) GetFixedRotation() bool{
    return self.Get("fixedRotation").Bool()
}

// -
func (self *PhysicsP2Body) SetFixedRotation(member bool) {
    self.Set("fixedRotation", member)
}

// The inertia of the body around the Z axis..
func (self *PhysicsP2Body) GetInertia() float64{
    return self.Get("inertia").Float()
}

// The inertia of the body around the Z axis..
func (self *PhysicsP2Body) SetInertia(member float64) {
    self.Set("inertia", member)
}

// The mass of the body.
func (self *PhysicsP2Body) GetMass() float64{
    return self.Get("mass").Float()
}

// The mass of the body.
func (self *PhysicsP2Body) SetMass(member float64) {
    self.Set("mass", member)
}

// The type of motion this body has. Should be one of: Body.STATIC (the body does not move), Body.DYNAMIC (body can move and respond to collisions) and Body.KINEMATIC (only moves according to its .velocity).
func (self *PhysicsP2Body) GetMotionState() float64{
    return self.Get("motionState").Float()
}

// The type of motion this body has. Should be one of: Body.STATIC (the body does not move), Body.DYNAMIC (body can move and respond to collisions) and Body.KINEMATIC (only moves according to its .velocity).
func (self *PhysicsP2Body) SetMotionState(member float64) {
    self.Set("motionState", member)
}

// The angle of the Body in radians.
// If you wish to work in degrees instead of radians use the Body.angle property instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in radians.
func (self *PhysicsP2Body) GetRotation() float64{
    return self.Get("rotation").Float()
}

// The angle of the Body in radians.
// If you wish to work in degrees instead of radians use the Body.angle property instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in radians.
func (self *PhysicsP2Body) SetRotation(member float64) {
    self.Set("rotation", member)
}

// .
func (self *PhysicsP2Body) GetSleepSpeedLimit() float64{
    return self.Get("sleepSpeedLimit").Float()
}

// .
func (self *PhysicsP2Body) SetSleepSpeedLimit(member float64) {
    self.Set("sleepSpeedLimit", member)
}

// The x coordinate of this Body.
func (self *PhysicsP2Body) GetX() float64{
    return self.Get("x").Float()
}

// The x coordinate of this Body.
func (self *PhysicsP2Body) SetX(member float64) {
    self.Set("x", member)
}

// The y coordinate of this Body.
func (self *PhysicsP2Body) GetY() float64{
    return self.Get("y").Float()
}

// The y coordinate of this Body.
func (self *PhysicsP2Body) SetY(member float64) {
    self.Set("y", member)
}

// The Body ID. Each Body that has been added to the World has a unique ID.
func (self *PhysicsP2Body) GetId() float64{
    return self.Get("id").Float()
}

// The Body ID. Each Body that has been added to the World has a unique ID.
func (self *PhysicsP2Body) SetId(member float64) {
    self.Set("id", member)
}

// Enable or disable debug drawing of this body
func (self *PhysicsP2Body) GetDebug() bool{
    return self.Get("debug").Bool()
}

// Enable or disable debug drawing of this body
func (self *PhysicsP2Body) SetDebug(member bool) {
    self.Set("debug", member)
}

// A Body can be set to collide against the World bounds automatically if this is set to true. Otherwise it will leave the World.
// Note that this only applies if your World has bounds! The response to the collision should be managed via CollisionMaterials.
// Also note that when you set this it will only effect Body shapes that already exist. If you then add further shapes to your Body
// after setting this it will *not* proactively set them to collide with the bounds. Should the Body collide with the World bounds?
func (self *PhysicsP2Body) GetCollideWorldBounds() bool{
    return self.Get("collideWorldBounds").Bool()
}

// A Body can be set to collide against the World bounds automatically if this is set to true. Otherwise it will leave the World.
// Note that this only applies if your World has bounds! The response to the collision should be managed via CollisionMaterials.
// Also note that when you set this it will only effect Body shapes that already exist. If you then add further shapes to your Body
// after setting this it will *not* proactively set them to collide with the bounds. Should the Body collide with the World bounds?
func (self *PhysicsP2Body) SetCollideWorldBounds(member bool) {
    self.Set("collideWorldBounds", member)
}



// Sets a callback to be fired any time a shape in this Body impacts with a shape in the given Body. The impact test is performed against body.id values.
// The callback will be sent 4 parameters: This body, the body that impacted, the Shape in this body and the shape in the impacting body.
// Note that the impact event happens after collision resolution, so it cannot be used to prevent a collision from happening.
// It also happens mid-step. So do not destroy a Body during this callback, instead set safeDestroy to true so it will be killed on the next preUpdate.
func (self *PhysicsP2Body) CreateBodyCallbackI(args ...interface{}) {
    self.Call("createBodyCallback", args)
}

// Sets a callback to be fired any time this Body impacts with the given Group. The impact test is performed against shape.collisionGroup values.
// The callback will be sent 4 parameters: This body, the body that impacted, the Shape in this body and the shape in the impacting body.
// This callback will only fire if this Body has been assigned a collision group.
// Note that the impact event happens after collision resolution, so it cannot be used to prevent a collision from happening.
// It also happens mid-step. So do not destroy a Body during this callback, instead set safeDestroy to true so it will be killed on the next preUpdate.
func (self *PhysicsP2Body) CreateGroupCallbackI(args ...interface{}) {
    self.Call("createGroupCallback", args)
}

// Gets the collision bitmask from the groups this body collides with.
func (self *PhysicsP2Body) GetCollisionMaskI(args ...interface{}) float64{
    return self.Call("getCollisionMask", args).Float()
}

// Updates the collisionMask.
func (self *PhysicsP2Body) UpdateCollisionMaskI(args ...interface{}) {
    self.Call("updateCollisionMask", args)
}

// Sets the given CollisionGroup to be the collision group for all shapes in this Body, unless a shape is specified.
// This also resets the collisionMask.
func (self *PhysicsP2Body) SetCollisionGroupI(args ...interface{}) {
    self.Call("setCollisionGroup", args)
}

// Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollisionI(args ...interface{}) {
    self.Call("clearCollision", args)
}

// Removes the given CollisionGroup, or array of CollisionGroups, from the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) RemoveCollisionGroupI(args ...interface{}) {
    self.Call("removeCollisionGroup", args)
}

// Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) CollidesI(args ...interface{}) {
    self.Call("collides", args)
}

// Moves the shape offsets so their center of mass becomes the body center of mass.
func (self *PhysicsP2Body) AdjustCenterOfMassI(args ...interface{}) {
    self.Call("adjustCenterOfMass", args)
}

// Gets the velocity of a point in the body.
func (self *PhysicsP2Body) GetVelocityAtPointI(args ...interface{}) []interface{}{
	array := self.Call("getVelocityAtPoint", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Apply damping, see http://code.google.com/p/bullet/issues/detail?id=74 for details.
func (self *PhysicsP2Body) ApplyDampingI(args ...interface{}) {
    self.Call("applyDamping", args)
}

// Apply impulse to a point relative to the body.
// This could for example be a point on the Body surface. An impulse is a force added to a body during a short 
// period of time (impulse = force * time). Impulses will be added to Body.velocity and Body.angularVelocity.
func (self *PhysicsP2Body) ApplyImpulseI(args ...interface{}) {
    self.Call("applyImpulse", args)
}

// Apply impulse to a point local to the body.
// 
// This could for example be a point on the Body surface. An impulse is a force added to a body during a short 
// period of time (impulse = force * time). Impulses will be added to Body.velocity and Body.angularVelocity.
func (self *PhysicsP2Body) ApplyImpulseLocalI(args ...interface{}) {
    self.Call("applyImpulseLocal", args)
}

// Apply force to a world point.
// 
// This could for example be a point on the RigidBody surface. Applying force 
// this way will add to Body.force and Body.angularForce.
func (self *PhysicsP2Body) ApplyForceI(args ...interface{}) {
    self.Call("applyForce", args)
}

// Sets the force on the body to zero.
func (self *PhysicsP2Body) SetZeroForceI(args ...interface{}) {
    self.Call("setZeroForce", args)
}

// If this Body is dynamic then this will zero its angular velocity.
func (self *PhysicsP2Body) SetZeroRotationI(args ...interface{}) {
    self.Call("setZeroRotation", args)
}

// If this Body is dynamic then this will zero its velocity on both axis.
func (self *PhysicsP2Body) SetZeroVelocityI(args ...interface{}) {
    self.Call("setZeroVelocity", args)
}

// Sets the Body damping and angularDamping to zero.
func (self *PhysicsP2Body) SetZeroDampingI(args ...interface{}) {
    self.Call("setZeroDamping", args)
}

// Transform a world point to local body frame.
func (self *PhysicsP2Body) ToLocalFrameI(args ...interface{}) {
    self.Call("toLocalFrame", args)
}

// Transform a local point to world frame.
func (self *PhysicsP2Body) ToWorldFrameI(args ...interface{}) {
    self.Call("toWorldFrame", args)
}

// This will rotate the Body by the given speed to the left (counter-clockwise).
func (self *PhysicsP2Body) RotateLeftI(args ...interface{}) {
    self.Call("rotateLeft", args)
}

// This will rotate the Body by the given speed to the left (clockwise).
func (self *PhysicsP2Body) RotateRightI(args ...interface{}) {
    self.Call("rotateRight", args)
}

// Moves the Body forwards based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveForwardI(args ...interface{}) {
    self.Call("moveForward", args)
}

// Moves the Body backwards based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveBackwardI(args ...interface{}) {
    self.Call("moveBackward", args)
}

// Applies a force to the Body that causes it to 'thrust' forwards, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustI(args ...interface{}) {
    self.Call("thrust", args)
}

// Applies a force to the Body that causes it to 'thrust' to the left, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustLeftI(args ...interface{}) {
    self.Call("thrustLeft", args)
}

// Applies a force to the Body that causes it to 'thrust' to the right, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustRightI(args ...interface{}) {
    self.Call("thrustRight", args)
}

// Applies a force to the Body that causes it to 'thrust' backwards (in reverse), based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ReverseI(args ...interface{}) {
    self.Call("reverse", args)
}

// If this Body is dynamic then this will move it to the left by setting its x velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveLeftI(args ...interface{}) {
    self.Call("moveLeft", args)
}

// If this Body is dynamic then this will move it to the right by setting its x velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveRightI(args ...interface{}) {
    self.Call("moveRight", args)
}

// If this Body is dynamic then this will move it up by setting its y velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveUpI(args ...interface{}) {
    self.Call("moveUp", args)
}

// If this Body is dynamic then this will move it down by setting its y velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveDownI(args ...interface{}) {
    self.Call("moveDown", args)
}

// Internal method. This is called directly before the sprites are sent to the renderer and after the update function has finished.
func (self *PhysicsP2Body) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Internal method. This is called directly before the sprites are sent to the renderer and after the update function has finished.
func (self *PhysicsP2Body) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}

// Resets the Body force, velocity (linear and angular) and rotation. Optionally resets damping and mass.
func (self *PhysicsP2Body) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Adds this physics body to the world.
func (self *PhysicsP2Body) AddToWorldI(args ...interface{}) {
    self.Call("addToWorld", args)
}

// Removes this physics body from the world.
func (self *PhysicsP2Body) RemoveFromWorldI(args ...interface{}) {
    self.Call("removeFromWorld", args)
}

// Destroys this Body and all references it holds to other objects.
func (self *PhysicsP2Body) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Removes all Shapes from this Body.
func (self *PhysicsP2Body) ClearShapesI(args ...interface{}) {
    self.Call("clearShapes", args)
}

// Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShapeI(args ...interface{}) *P2Shape{
    return &P2Shape{self.Call("addShape", args)}
}

// Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircleI(args ...interface{}) *P2Circle{
    return &P2Circle{self.Call("addCircle", args)}
}

// Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangleI(args ...interface{}) *P2Box{
    return &P2Box{self.Call("addRectangle", args)}
}

// Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlaneI(args ...interface{}) *P2Plane{
    return &P2Plane{self.Call("addPlane", args)}
}

// Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticleI(args ...interface{}) *P2Particle{
    return &P2Particle{self.Call("addParticle", args)}
}

// Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLineI(args ...interface{}) *P2Line{
    return &P2Line{self.Call("addLine", args)}
}

// Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsuleI(args ...interface{}) *P2Capsule{
    return &P2Capsule{self.Call("addCapsule", args)}
}

// Reads a polygon shape path, and assembles convex shapes from that and puts them at proper offset points. The shape must be simple and without holes.
// This function expects the x.y values to be given in pixels. If you want to provide them at p2 world scales then call Body.data.fromPolygon directly.
func (self *PhysicsP2Body) AddPolygonI(args ...interface{}) bool{
    return self.Call("addPolygon", args).Bool()
}

// Remove a shape from the body. Will automatically update the mass properties and bounding radius.
func (self *PhysicsP2Body) RemoveShapeI(args ...interface{}) bool{
    return self.Call("removeShape", args).Bool()
}

// Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircleI(args ...interface{}) {
    self.Call("setCircle", args)
}

// Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangleI(args ...interface{}) *P2Rectangle{
    return &P2Rectangle{self.Call("setRectangle", args)}
}

// Clears any previously set shapes.
// Then creates a Rectangle shape sized to match the dimensions and orientation of the Sprite given.
// If no Sprite is given it defaults to using the parent of this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangleFromSpriteI(args ...interface{}) *P2Rectangle{
    return &P2Rectangle{self.Call("setRectangleFromSprite", args)}
}

// Adds the given Material to all Shapes that belong to this Body.
// If you only wish to apply it to a specific Shape in this Body then provide that as the 2nd parameter.
func (self *PhysicsP2Body) SetMaterialI(args ...interface{}) {
    self.Call("setMaterial", args)
}

// Updates the debug draw if any body shapes change.
func (self *PhysicsP2Body) ShapeChangedI(args ...interface{}) {
    self.Call("shapeChanged", args)
}

// Reads the shape data from a physics data file stored in the Game.Cache and adds it as a polygon to this Body.
// The shape data format is based on the output of the
// {@link https://github.com/photonstorm/phaser/tree/master/resources/PhysicsEditor%20Exporter|custom phaser exporter} for
// {@link https://www.codeandweb.com/physicseditor|PhysicsEditor}
func (self *PhysicsP2Body) AddPhaserPolygonI(args ...interface{}) []interface{}{
	array := self.Call("addPhaserPolygon", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Add a polygon fixture. This is used during #loadPolygon.
func (self *PhysicsP2Body) AddFixtureI(args ...interface{}) []interface{}{
	array := self.Call("addFixture", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Reads the shape data from a physics data file stored in the Game.Cache and adds it as a polygon to this Body.
// 
// As well as reading the data from the Cache you can also pass `null` as the first argument and a
// physics data object as the second. When doing this you must ensure the structure of the object is correct in advance.
// 
// For more details see the format of the Lime / Corona Physics Editor export.
func (self *PhysicsP2Body) LoadPolygonI(args ...interface{}) bool{
    return self.Call("loadPolygon", args).Bool()
}
