// Package phaser Automatic generation for Phaser.Physics.P2.Body
// generated file PhysicsP2Body.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsP2Body The Physics Body is typically linked to a single Sprite and defines properties that determine how the physics body is simulated.
// These properties affect how the body reacts to forces, what forces it generates on itself (to simulate friction), and how it reacts to collisions in the scene.
// In most cases, the properties are used to simulate physical effects. Each body also has its own property values that determine exactly how it reacts to forces and collisions in the scene.
// By default a single Rectangle shape is added to the Body that matches the dimensions of the parent Sprite. See addShape, removeShape, clearShapes to add extra shapes around the Body.
// Note: When bound to a Sprite to avoid single-pixel jitters on mobile devices we strongly recommend using Sprite sizes that are even on both axis, i.e. 128x128 not 127x127.
// Note: When a game object is given a P2 body it has its anchor x/y set to 0.5, so it becomes centered.
type PhysicsP2Body struct {
    *js.Object
}

// NewPhysicsP2Body The Physics Body is typically linked to a single Sprite and defines properties that determine how the physics body is simulated.
// These properties affect how the body reacts to forces, what forces it generates on itself (to simulate friction), and how it reacts to collisions in the scene.
// In most cases, the properties are used to simulate physical effects. Each body also has its own property values that determine exactly how it reacts to forces and collisions in the scene.
// By default a single Rectangle shape is added to the Body that matches the dimensions of the parent Sprite. See addShape, removeShape, clearShapes to add extra shapes around the Body.
// Note: When bound to a Sprite to avoid single-pixel jitters on mobile devices we strongly recommend using Sprite sizes that are even on both axis, i.e. 128x128 not 127x127.
// Note: When a game object is given a P2 body it has its anchor x/y set to 0.5, so it becomes centered.
func NewPhysicsP2Body(game *Game) *PhysicsP2Body {
    return &PhysicsP2Body{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Body").New(game)}
}
// NewPhysicsP2Body1O The Physics Body is typically linked to a single Sprite and defines properties that determine how the physics body is simulated.
// These properties affect how the body reacts to forces, what forces it generates on itself (to simulate friction), and how it reacts to collisions in the scene.
// In most cases, the properties are used to simulate physical effects. Each body also has its own property values that determine exactly how it reacts to forces and collisions in the scene.
// By default a single Rectangle shape is added to the Body that matches the dimensions of the parent Sprite. See addShape, removeShape, clearShapes to add extra shapes around the Body.
// Note: When bound to a Sprite to avoid single-pixel jitters on mobile devices we strongly recommend using Sprite sizes that are even on both axis, i.e. 128x128 not 127x127.
// Note: When a game object is given a P2 body it has its anchor x/y set to 0.5, so it becomes centered.
func NewPhysicsP2Body1O(game *Game, sprite *Sprite) *PhysicsP2Body {
    return &PhysicsP2Body{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Body").New(game, sprite)}
}
// NewPhysicsP2Body2O The Physics Body is typically linked to a single Sprite and defines properties that determine how the physics body is simulated.
// These properties affect how the body reacts to forces, what forces it generates on itself (to simulate friction), and how it reacts to collisions in the scene.
// In most cases, the properties are used to simulate physical effects. Each body also has its own property values that determine exactly how it reacts to forces and collisions in the scene.
// By default a single Rectangle shape is added to the Body that matches the dimensions of the parent Sprite. See addShape, removeShape, clearShapes to add extra shapes around the Body.
// Note: When bound to a Sprite to avoid single-pixel jitters on mobile devices we strongly recommend using Sprite sizes that are even on both axis, i.e. 128x128 not 127x127.
// Note: When a game object is given a P2 body it has its anchor x/y set to 0.5, so it becomes centered.
func NewPhysicsP2Body2O(game *Game, sprite *Sprite, x int) *PhysicsP2Body {
    return &PhysicsP2Body{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Body").New(game, sprite, x)}
}
// NewPhysicsP2Body3O The Physics Body is typically linked to a single Sprite and defines properties that determine how the physics body is simulated.
// These properties affect how the body reacts to forces, what forces it generates on itself (to simulate friction), and how it reacts to collisions in the scene.
// In most cases, the properties are used to simulate physical effects. Each body also has its own property values that determine exactly how it reacts to forces and collisions in the scene.
// By default a single Rectangle shape is added to the Body that matches the dimensions of the parent Sprite. See addShape, removeShape, clearShapes to add extra shapes around the Body.
// Note: When bound to a Sprite to avoid single-pixel jitters on mobile devices we strongly recommend using Sprite sizes that are even on both axis, i.e. 128x128 not 127x127.
// Note: When a game object is given a P2 body it has its anchor x/y set to 0.5, so it becomes centered.
func NewPhysicsP2Body3O(game *Game, sprite *Sprite, x int, y int) *PhysicsP2Body {
    return &PhysicsP2Body{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Body").New(game, sprite, x, y)}
}
// NewPhysicsP2Body4O The Physics Body is typically linked to a single Sprite and defines properties that determine how the physics body is simulated.
// These properties affect how the body reacts to forces, what forces it generates on itself (to simulate friction), and how it reacts to collisions in the scene.
// In most cases, the properties are used to simulate physical effects. Each body also has its own property values that determine exactly how it reacts to forces and collisions in the scene.
// By default a single Rectangle shape is added to the Body that matches the dimensions of the parent Sprite. See addShape, removeShape, clearShapes to add extra shapes around the Body.
// Note: When bound to a Sprite to avoid single-pixel jitters on mobile devices we strongly recommend using Sprite sizes that are even on both axis, i.e. 128x128 not 127x127.
// Note: When a game object is given a P2 body it has its anchor x/y set to 0.5, so it becomes centered.
func NewPhysicsP2Body4O(game *Game, sprite *Sprite, x int, y int, mass int) *PhysicsP2Body {
    return &PhysicsP2Body{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Body").New(game, sprite, x, y, mass)}
}
// NewPhysicsP2BodyI The Physics Body is typically linked to a single Sprite and defines properties that determine how the physics body is simulated.
// These properties affect how the body reacts to forces, what forces it generates on itself (to simulate friction), and how it reacts to collisions in the scene.
// In most cases, the properties are used to simulate physical effects. Each body also has its own property values that determine exactly how it reacts to forces and collisions in the scene.
// By default a single Rectangle shape is added to the Body that matches the dimensions of the parent Sprite. See addShape, removeShape, clearShapes to add extra shapes around the Body.
// Note: When bound to a Sprite to avoid single-pixel jitters on mobile devices we strongly recommend using Sprite sizes that are even on both axis, i.e. 128x128 not 127x127.
// Note: When a game object is given a P2 body it has its anchor x/y set to 0.5, so it becomes centered.
func NewPhysicsP2BodyI(args ...interface{}) *PhysicsP2Body {
    return &PhysicsP2Body{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Body").New(args)}
}



// PhysicsP2Body Binding conversion method to PhysicsP2Body point 
func ToPhysicsP2Body(jsStruct interface{}) *PhysicsP2Body {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsP2Body{Object: object}
	}
	return nil
}



// Game Local reference to game.
func (self *PhysicsP2Body) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *PhysicsP2Body) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// World Local reference to the P2 World.
func (self *PhysicsP2Body) World() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// SetWorldA Local reference to the P2 World.
func (self *PhysicsP2Body) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}

// Sprite Reference to the parent Sprite.
func (self *PhysicsP2Body) Sprite() *Sprite{
    return &Sprite{self.Object.Get("sprite")}
}

// SetSpriteA Reference to the parent Sprite.
func (self *PhysicsP2Body) SetSpriteA(member *Sprite) {
    self.Object.Set("sprite", member)
}

// Type The type of physics system this body belongs to.
func (self *PhysicsP2Body) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The type of physics system this body belongs to.
func (self *PhysicsP2Body) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Offset The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsP2Body) Offset() *Point{
    return &Point{self.Object.Get("offset")}
}

// SetOffsetA The offset of the Physics Body from the Sprite x/y position.
func (self *PhysicsP2Body) SetOffsetA(member *Point) {
    self.Object.Set("offset", member)
}

// Data The p2 Body data.
func (self *PhysicsP2Body) Data() *P2Body{
    return &P2Body{self.Object.Get("data")}
}

// SetDataA The p2 Body data.
func (self *PhysicsP2Body) SetDataA(member *P2Body) {
    self.Object.Set("data", member)
}

// Velocity The velocity of the body. Set velocity.x to a negative value to move to the left, position to the right. velocity.y negative values move up, positive move down.
func (self *PhysicsP2Body) Velocity() *PhysicsP2InversePointProxy{
    return &PhysicsP2InversePointProxy{self.Object.Get("velocity")}
}

// SetVelocityA The velocity of the body. Set velocity.x to a negative value to move to the left, position to the right. velocity.y negative values move up, positive move down.
func (self *PhysicsP2Body) SetVelocityA(member *PhysicsP2InversePointProxy) {
    self.Object.Set("velocity", member)
}

// Force The force applied to the body.
func (self *PhysicsP2Body) Force() *PhysicsP2InversePointProxy{
    return &PhysicsP2InversePointProxy{self.Object.Get("force")}
}

// SetForceA The force applied to the body.
func (self *PhysicsP2Body) SetForceA(member *PhysicsP2InversePointProxy) {
    self.Object.Set("force", member)
}

// Gravity A locally applied gravity force to the Body. Applied directly before the world step. NOTE: Not currently implemented.
func (self *PhysicsP2Body) Gravity() *Point{
    return &Point{self.Object.Get("gravity")}
}

// SetGravityA A locally applied gravity force to the Body. Applied directly before the world step. NOTE: Not currently implemented.
func (self *PhysicsP2Body) SetGravityA(member *Point) {
    self.Object.Set("gravity", member)
}

// OnBeginContact Dispatched when a first contact is created between shapes in two bodies. 
// This event is fired during the step, so collision has already taken place.
// 
// The event will be sent 5 arguments in this order:
// 
// The Phaser.Physics.P2.Body it is in contact with. *This might be null* if the Body was created directly in the p2 world.
// The p2.Body this Body is in contact with.
// The Shape from this body that caused the contact.
// The Shape from the contact body.
// The Contact Equation data array.
func (self *PhysicsP2Body) OnBeginContact() *Signal{
    return &Signal{self.Object.Get("onBeginContact")}
}

// SetOnBeginContactA Dispatched when a first contact is created between shapes in two bodies. 
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

// OnEndContact Dispatched when contact ends between shapes in two bodies.
// This event is fired during the step, so collision has already taken place.
// 
// The event will be sent 4 arguments in this order:
// 
// The Phaser.Physics.P2.Body it is in contact with. *This might be null* if the Body was created directly in the p2 world.
// The p2.Body this Body has ended contact with.
// The Shape from this body that caused the original contact.
// The Shape from the contact body.
func (self *PhysicsP2Body) OnEndContact() *Signal{
    return &Signal{self.Object.Get("onEndContact")}
}

// SetOnEndContactA Dispatched when contact ends between shapes in two bodies.
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

// CollidesWith Array of CollisionGroups that this Bodies shapes collide with.
func (self *PhysicsP2Body) CollidesWith() []interface{}{
	array00 := self.Object.Get("collidesWith")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetCollidesWithA Array of CollisionGroups that this Bodies shapes collide with.
func (self *PhysicsP2Body) SetCollidesWithA(member []interface{}) {
    self.Object.Set("collidesWith", member)
}

// RemoveNextStep To avoid deleting this body during a physics step, and causing all kinds of problems, set removeNextStep to true to have it removed in the next preUpdate.
func (self *PhysicsP2Body) RemoveNextStep() bool{
    return self.Object.Get("removeNextStep").Bool()
}

// SetRemoveNextStepA To avoid deleting this body during a physics step, and causing all kinds of problems, set removeNextStep to true to have it removed in the next preUpdate.
func (self *PhysicsP2Body) SetRemoveNextStepA(member bool) {
    self.Object.Set("removeNextStep", member)
}

// DebugBody Reference to the debug body.
func (self *PhysicsP2Body) DebugBody() *PhysicsP2BodyDebug{
    return &PhysicsP2BodyDebug{self.Object.Get("debugBody")}
}

// SetDebugBodyA Reference to the debug body.
func (self *PhysicsP2Body) SetDebugBodyA(member *PhysicsP2BodyDebug) {
    self.Object.Set("debugBody", member)
}

// Dirty Internally used by Sprite.x/y
func (self *PhysicsP2Body) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// SetDirtyA Internally used by Sprite.x/y
func (self *PhysicsP2Body) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// DYNAMIC Dynamic body. Dynamic bodies body can move and respond to collisions and forces.
func (self *PhysicsP2Body) DYNAMIC() int{
    return self.Object.Get("DYNAMIC").Int()
}

// SetDYNAMICA Dynamic body. Dynamic bodies body can move and respond to collisions and forces.
func (self *PhysicsP2Body) SetDYNAMICA(member int) {
    self.Object.Set("DYNAMIC", member)
}

// STATIC Static body. Static bodies do not move, and they do not respond to forces or collision.
func (self *PhysicsP2Body) STATIC() int{
    return self.Object.Get("STATIC").Int()
}

// SetSTATICA Static body. Static bodies do not move, and they do not respond to forces or collision.
func (self *PhysicsP2Body) SetSTATICA(member int) {
    self.Object.Set("STATIC", member)
}

// KINEMATIC Kinematic body. Kinematic bodies only moves according to its .velocity, and does not respond to collisions or force.
func (self *PhysicsP2Body) KINEMATIC() int{
    return self.Object.Get("KINEMATIC").Int()
}

// SetKINEMATICA Kinematic body. Kinematic bodies only moves according to its .velocity, and does not respond to collisions or force.
func (self *PhysicsP2Body) SetKINEMATICA(member int) {
    self.Object.Set("KINEMATIC", member)
}

// Static Returns true if the Body is static. Setting Body.static to 'false' will make it dynamic.
func (self *PhysicsP2Body) Static() bool{
    return self.Object.Get("static").Bool()
}

// SetStaticA Returns true if the Body is static. Setting Body.static to 'false' will make it dynamic.
func (self *PhysicsP2Body) SetStaticA(member bool) {
    self.Object.Set("static", member)
}

// Dynamic Returns true if the Body is dynamic. Setting Body.dynamic to 'false' will make it static.
func (self *PhysicsP2Body) Dynamic() bool{
    return self.Object.Get("dynamic").Bool()
}

// SetDynamicA Returns true if the Body is dynamic. Setting Body.dynamic to 'false' will make it static.
func (self *PhysicsP2Body) SetDynamicA(member bool) {
    self.Object.Set("dynamic", member)
}

// Kinematic Returns true if the Body is kinematic. Setting Body.kinematic to 'false' will make it static.
func (self *PhysicsP2Body) Kinematic() bool{
    return self.Object.Get("kinematic").Bool()
}

// SetKinematicA Returns true if the Body is kinematic. Setting Body.kinematic to 'false' will make it static.
func (self *PhysicsP2Body) SetKinematicA(member bool) {
    self.Object.Set("kinematic", member)
}

// AllowSleep -
func (self *PhysicsP2Body) AllowSleep() bool{
    return self.Object.Get("allowSleep").Bool()
}

// SetAllowSleepA -
func (self *PhysicsP2Body) SetAllowSleepA(member bool) {
    self.Object.Set("allowSleep", member)
}

// Angle The angle of the Body in degrees from its original orientation. Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. For example, the statement Body.angle = 450 is the same as Body.angle = 90.
// If you wish to work in radians instead of degrees use the property Body.rotation instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in degrees.
func (self *PhysicsP2Body) Angle() int{
    return self.Object.Get("angle").Int()
}

// SetAngleA The angle of the Body in degrees from its original orientation. Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. For example, the statement Body.angle = 450 is the same as Body.angle = 90.
// If you wish to work in radians instead of degrees use the property Body.rotation instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in degrees.
func (self *PhysicsP2Body) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// AngularDamping Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The angular damping acting acting on the body.
func (self *PhysicsP2Body) AngularDamping() int{
    return self.Object.Get("angularDamping").Int()
}

// SetAngularDampingA Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The angular damping acting acting on the body.
func (self *PhysicsP2Body) SetAngularDampingA(member int) {
    self.Object.Set("angularDamping", member)
}

// AngularForce The angular force acting on the body.
func (self *PhysicsP2Body) AngularForce() int{
    return self.Object.Get("angularForce").Int()
}

// SetAngularForceA The angular force acting on the body.
func (self *PhysicsP2Body) SetAngularForceA(member int) {
    self.Object.Set("angularForce", member)
}

// AngularVelocity The angular velocity of the body.
func (self *PhysicsP2Body) AngularVelocity() int{
    return self.Object.Get("angularVelocity").Int()
}

// SetAngularVelocityA The angular velocity of the body.
func (self *PhysicsP2Body) SetAngularVelocityA(member int) {
    self.Object.Set("angularVelocity", member)
}

// Damping Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The linear damping acting on the body in the velocity direction.
func (self *PhysicsP2Body) Damping() int{
    return self.Object.Get("damping").Int()
}

// SetDampingA Damping is specified as a value between 0 and 1, which is the proportion of velocity lost per second. The linear damping acting on the body in the velocity direction.
func (self *PhysicsP2Body) SetDampingA(member int) {
    self.Object.Set("damping", member)
}

// FixedRotation -
func (self *PhysicsP2Body) FixedRotation() bool{
    return self.Object.Get("fixedRotation").Bool()
}

// SetFixedRotationA -
func (self *PhysicsP2Body) SetFixedRotationA(member bool) {
    self.Object.Set("fixedRotation", member)
}

// Inertia The inertia of the body around the Z axis..
func (self *PhysicsP2Body) Inertia() int{
    return self.Object.Get("inertia").Int()
}

// SetInertiaA The inertia of the body around the Z axis..
func (self *PhysicsP2Body) SetInertiaA(member int) {
    self.Object.Set("inertia", member)
}

// Mass The mass of the body.
func (self *PhysicsP2Body) Mass() int{
    return self.Object.Get("mass").Int()
}

// SetMassA The mass of the body.
func (self *PhysicsP2Body) SetMassA(member int) {
    self.Object.Set("mass", member)
}

// MotionState The type of motion this body has. Should be one of: Body.STATIC (the body does not move), Body.DYNAMIC (body can move and respond to collisions) and Body.KINEMATIC (only moves according to its .velocity).
func (self *PhysicsP2Body) MotionState() int{
    return self.Object.Get("motionState").Int()
}

// SetMotionStateA The type of motion this body has. Should be one of: Body.STATIC (the body does not move), Body.DYNAMIC (body can move and respond to collisions) and Body.KINEMATIC (only moves according to its .velocity).
func (self *PhysicsP2Body) SetMotionStateA(member int) {
    self.Object.Set("motionState", member)
}

// Rotation The angle of the Body in radians.
// If you wish to work in degrees instead of radians use the Body.angle property instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in radians.
func (self *PhysicsP2Body) Rotation() int{
    return self.Object.Get("rotation").Int()
}

// SetRotationA The angle of the Body in radians.
// If you wish to work in degrees instead of radians use the Body.angle property instead. Working in radians is faster as it doesn't have to convert values. The angle of this Body in radians.
func (self *PhysicsP2Body) SetRotationA(member int) {
    self.Object.Set("rotation", member)
}

// SleepSpeedLimit .
func (self *PhysicsP2Body) SleepSpeedLimit() int{
    return self.Object.Get("sleepSpeedLimit").Int()
}

// SetSleepSpeedLimitA .
func (self *PhysicsP2Body) SetSleepSpeedLimitA(member int) {
    self.Object.Set("sleepSpeedLimit", member)
}

// X The x coordinate of this Body.
func (self *PhysicsP2Body) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The x coordinate of this Body.
func (self *PhysicsP2Body) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The y coordinate of this Body.
func (self *PhysicsP2Body) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The y coordinate of this Body.
func (self *PhysicsP2Body) SetYA(member int) {
    self.Object.Set("y", member)
}

// Id The Body ID. Each Body that has been added to the World has a unique ID.
func (self *PhysicsP2Body) Id() int{
    return self.Object.Get("id").Int()
}

// SetIdA The Body ID. Each Body that has been added to the World has a unique ID.
func (self *PhysicsP2Body) SetIdA(member int) {
    self.Object.Set("id", member)
}

// Debug Enable or disable debug drawing of this body
func (self *PhysicsP2Body) Debug() bool{
    return self.Object.Get("debug").Bool()
}

// SetDebugA Enable or disable debug drawing of this body
func (self *PhysicsP2Body) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// CollideWorldBounds A Body can be set to collide against the World bounds automatically if this is set to true. Otherwise it will leave the World.
// Note that this only applies if your World has bounds! The response to the collision should be managed via CollisionMaterials.
// Also note that when you set this it will only effect Body shapes that already exist. If you then add further shapes to your Body
// after setting this it will *not* proactively set them to collide with the bounds. Should the Body collide with the World bounds?
func (self *PhysicsP2Body) CollideWorldBounds() bool{
    return self.Object.Get("collideWorldBounds").Bool()
}

// SetCollideWorldBoundsA A Body can be set to collide against the World bounds automatically if this is set to true. Otherwise it will leave the World.
// Note that this only applies if your World has bounds! The response to the collision should be managed via CollisionMaterials.
// Also note that when you set this it will only effect Body shapes that already exist. If you then add further shapes to your Body
// after setting this it will *not* proactively set them to collide with the bounds. Should the Body collide with the World bounds?
func (self *PhysicsP2Body) SetCollideWorldBoundsA(member bool) {
    self.Object.Set("collideWorldBounds", member)
}


// CreateBodyCallback Sets a callback to be fired any time a shape in this Body impacts with a shape in the given Body. The impact test is performed against body.id values.
// The callback will be sent 4 parameters: This body, the body that impacted, the Shape in this body and the shape in the impacting body.
// Note that the impact event happens after collision resolution, so it cannot be used to prevent a collision from happening.
// It also happens mid-step. So do not destroy a Body during this callback, instead set safeDestroy to true so it will be killed on the next preUpdate.
func (self *PhysicsP2Body) CreateBodyCallback(object interface{}, callback interface{}, callbackContext interface{}) {
    self.Object.Call("createBodyCallback", object, callback, callbackContext)
}

// CreateBodyCallbackI Sets a callback to be fired any time a shape in this Body impacts with a shape in the given Body. The impact test is performed against body.id values.
// The callback will be sent 4 parameters: This body, the body that impacted, the Shape in this body and the shape in the impacting body.
// Note that the impact event happens after collision resolution, so it cannot be used to prevent a collision from happening.
// It also happens mid-step. So do not destroy a Body during this callback, instead set safeDestroy to true so it will be killed on the next preUpdate.
func (self *PhysicsP2Body) CreateBodyCallbackI(args ...interface{}) {
    self.Object.Call("createBodyCallback", args)
}

// CreateGroupCallback Sets a callback to be fired any time this Body impacts with the given Group. The impact test is performed against shape.collisionGroup values.
// The callback will be sent 4 parameters: This body, the body that impacted, the Shape in this body and the shape in the impacting body.
// This callback will only fire if this Body has been assigned a collision group.
// Note that the impact event happens after collision resolution, so it cannot be used to prevent a collision from happening.
// It also happens mid-step. So do not destroy a Body during this callback, instead set safeDestroy to true so it will be killed on the next preUpdate.
func (self *PhysicsP2Body) CreateGroupCallback(group *PhysicsCollisionGroup, callback interface{}, callbackContext interface{}) {
    self.Object.Call("createGroupCallback", group, callback, callbackContext)
}

// CreateGroupCallbackI Sets a callback to be fired any time this Body impacts with the given Group. The impact test is performed against shape.collisionGroup values.
// The callback will be sent 4 parameters: This body, the body that impacted, the Shape in this body and the shape in the impacting body.
// This callback will only fire if this Body has been assigned a collision group.
// Note that the impact event happens after collision resolution, so it cannot be used to prevent a collision from happening.
// It also happens mid-step. So do not destroy a Body during this callback, instead set safeDestroy to true so it will be killed on the next preUpdate.
func (self *PhysicsP2Body) CreateGroupCallbackI(args ...interface{}) {
    self.Object.Call("createGroupCallback", args)
}

// GetCollisionMask Gets the collision bitmask from the groups this body collides with.
func (self *PhysicsP2Body) GetCollisionMask() int{
    return self.Object.Call("getCollisionMask").Int()
}

// GetCollisionMaskI Gets the collision bitmask from the groups this body collides with.
func (self *PhysicsP2Body) GetCollisionMaskI(args ...interface{}) int{
    return self.Object.Call("getCollisionMask", args).Int()
}

// UpdateCollisionMask Updates the collisionMask.
func (self *PhysicsP2Body) UpdateCollisionMask() {
    self.Object.Call("updateCollisionMask")
}

// UpdateCollisionMask1O Updates the collisionMask.
func (self *PhysicsP2Body) UpdateCollisionMask1O(shape *P2Shape) {
    self.Object.Call("updateCollisionMask", shape)
}

// UpdateCollisionMaskI Updates the collisionMask.
func (self *PhysicsP2Body) UpdateCollisionMaskI(args ...interface{}) {
    self.Object.Call("updateCollisionMask", args)
}

// SetCollisionGroup Sets the given CollisionGroup to be the collision group for all shapes in this Body, unless a shape is specified.
// This also resets the collisionMask.
func (self *PhysicsP2Body) SetCollisionGroup(group *PhysicsCollisionGroup) {
    self.Object.Call("setCollisionGroup", group)
}

// SetCollisionGroup1O Sets the given CollisionGroup to be the collision group for all shapes in this Body, unless a shape is specified.
// This also resets the collisionMask.
func (self *PhysicsP2Body) SetCollisionGroup1O(group *PhysicsCollisionGroup, shape *P2Shape) {
    self.Object.Call("setCollisionGroup", group, shape)
}

// SetCollisionGroupI Sets the given CollisionGroup to be the collision group for all shapes in this Body, unless a shape is specified.
// This also resets the collisionMask.
func (self *PhysicsP2Body) SetCollisionGroupI(args ...interface{}) {
    self.Object.Call("setCollisionGroup", args)
}

// ClearCollision Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollision() {
    self.Object.Call("clearCollision")
}

// ClearCollision1O Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollision1O(clearGroup bool) {
    self.Object.Call("clearCollision", clearGroup)
}

// ClearCollision2O Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollision2O(clearGroup bool, clearMask bool) {
    self.Object.Call("clearCollision", clearGroup, clearMask)
}

// ClearCollision3O Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollision3O(clearGroup bool, clearMask bool, shape *P2Shape) {
    self.Object.Call("clearCollision", clearGroup, clearMask, shape)
}

// ClearCollisionI Clears the collision data from the shapes in this Body. Optionally clears Group and/or Mask.
func (self *PhysicsP2Body) ClearCollisionI(args ...interface{}) {
    self.Object.Call("clearCollision", args)
}

// RemoveCollisionGroup Removes the given CollisionGroup, or array of CollisionGroups, from the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) RemoveCollisionGroup(group interface{}) {
    self.Object.Call("removeCollisionGroup", group)
}

// RemoveCollisionGroup1O Removes the given CollisionGroup, or array of CollisionGroups, from the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) RemoveCollisionGroup1O(group interface{}, clearCallback bool) {
    self.Object.Call("removeCollisionGroup", group, clearCallback)
}

// RemoveCollisionGroup2O Removes the given CollisionGroup, or array of CollisionGroups, from the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) RemoveCollisionGroup2O(group interface{}, clearCallback bool, shape *P2Shape) {
    self.Object.Call("removeCollisionGroup", group, clearCallback, shape)
}

// RemoveCollisionGroupI Removes the given CollisionGroup, or array of CollisionGroups, from the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) RemoveCollisionGroupI(args ...interface{}) {
    self.Object.Call("removeCollisionGroup", args)
}

// Collides Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) Collides(group interface{}) {
    self.Object.Call("collides", group)
}

// Collides1O Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) Collides1O(group interface{}, callback interface{}) {
    self.Object.Call("collides", group, callback)
}

// Collides2O Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) Collides2O(group interface{}, callback interface{}, callbackContext interface{}) {
    self.Object.Call("collides", group, callback, callbackContext)
}

// Collides3O Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) Collides3O(group interface{}, callback interface{}, callbackContext interface{}, shape *P2Shape) {
    self.Object.Call("collides", group, callback, callbackContext, shape)
}

// CollidesI Adds the given CollisionGroup, or array of CollisionGroups, to the list of groups that this body will collide with and updates the collision masks.
func (self *PhysicsP2Body) CollidesI(args ...interface{}) {
    self.Object.Call("collides", args)
}

// AdjustCenterOfMass Moves the shape offsets so their center of mass becomes the body center of mass.
func (self *PhysicsP2Body) AdjustCenterOfMass() {
    self.Object.Call("adjustCenterOfMass")
}

// AdjustCenterOfMassI Moves the shape offsets so their center of mass becomes the body center of mass.
func (self *PhysicsP2Body) AdjustCenterOfMassI(args ...interface{}) {
    self.Object.Call("adjustCenterOfMass", args)
}

// GetVelocityAtPoint Gets the velocity of a point in the body.
func (self *PhysicsP2Body) GetVelocityAtPoint(result []interface{}, relativePoint []interface{}) []interface{}{
	array00 := self.Object.Call("getVelocityAtPoint", result, relativePoint)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetVelocityAtPointI Gets the velocity of a point in the body.
func (self *PhysicsP2Body) GetVelocityAtPointI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("getVelocityAtPoint", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ApplyDamping Apply damping, see http://code.google.com/p/bullet/issues/detail?id=74 for details.
func (self *PhysicsP2Body) ApplyDamping(dt int) {
    self.Object.Call("applyDamping", dt)
}

// ApplyDampingI Apply damping, see http://code.google.com/p/bullet/issues/detail?id=74 for details.
func (self *PhysicsP2Body) ApplyDampingI(args ...interface{}) {
    self.Object.Call("applyDamping", args)
}

// ApplyImpulse Apply impulse to a point relative to the body.
// This could for example be a point on the Body surface. An impulse is a force added to a body during a short 
// period of time (impulse = force * time). Impulses will be added to Body.velocity and Body.angularVelocity.
func (self *PhysicsP2Body) ApplyImpulse(impulse interface{}, worldX int, worldY int) {
    self.Object.Call("applyImpulse", impulse, worldX, worldY)
}

// ApplyImpulseI Apply impulse to a point relative to the body.
// This could for example be a point on the Body surface. An impulse is a force added to a body during a short 
// period of time (impulse = force * time). Impulses will be added to Body.velocity and Body.angularVelocity.
func (self *PhysicsP2Body) ApplyImpulseI(args ...interface{}) {
    self.Object.Call("applyImpulse", args)
}

// ApplyImpulseLocal Apply impulse to a point local to the body.
// 
// This could for example be a point on the Body surface. An impulse is a force added to a body during a short 
// period of time (impulse = force * time). Impulses will be added to Body.velocity and Body.angularVelocity.
func (self *PhysicsP2Body) ApplyImpulseLocal(impulse interface{}, localX int, localY int) {
    self.Object.Call("applyImpulseLocal", impulse, localX, localY)
}

// ApplyImpulseLocalI Apply impulse to a point local to the body.
// 
// This could for example be a point on the Body surface. An impulse is a force added to a body during a short 
// period of time (impulse = force * time). Impulses will be added to Body.velocity and Body.angularVelocity.
func (self *PhysicsP2Body) ApplyImpulseLocalI(args ...interface{}) {
    self.Object.Call("applyImpulseLocal", args)
}

// ApplyForce Apply force to a world point.
// 
// This could for example be a point on the RigidBody surface. Applying force 
// this way will add to Body.force and Body.angularForce.
func (self *PhysicsP2Body) ApplyForce(force interface{}, worldX int, worldY int) {
    self.Object.Call("applyForce", force, worldX, worldY)
}

// ApplyForceI Apply force to a world point.
// 
// This could for example be a point on the RigidBody surface. Applying force 
// this way will add to Body.force and Body.angularForce.
func (self *PhysicsP2Body) ApplyForceI(args ...interface{}) {
    self.Object.Call("applyForce", args)
}

// SetZeroForce Sets the force on the body to zero.
func (self *PhysicsP2Body) SetZeroForce() {
    self.Object.Call("setZeroForce")
}

// SetZeroForceI Sets the force on the body to zero.
func (self *PhysicsP2Body) SetZeroForceI(args ...interface{}) {
    self.Object.Call("setZeroForce", args)
}

// SetZeroRotation If this Body is dynamic then this will zero its angular velocity.
func (self *PhysicsP2Body) SetZeroRotation() {
    self.Object.Call("setZeroRotation")
}

// SetZeroRotationI If this Body is dynamic then this will zero its angular velocity.
func (self *PhysicsP2Body) SetZeroRotationI(args ...interface{}) {
    self.Object.Call("setZeroRotation", args)
}

// SetZeroVelocity If this Body is dynamic then this will zero its velocity on both axis.
func (self *PhysicsP2Body) SetZeroVelocity() {
    self.Object.Call("setZeroVelocity")
}

// SetZeroVelocityI If this Body is dynamic then this will zero its velocity on both axis.
func (self *PhysicsP2Body) SetZeroVelocityI(args ...interface{}) {
    self.Object.Call("setZeroVelocity", args)
}

// SetZeroDamping Sets the Body damping and angularDamping to zero.
func (self *PhysicsP2Body) SetZeroDamping() {
    self.Object.Call("setZeroDamping")
}

// SetZeroDampingI Sets the Body damping and angularDamping to zero.
func (self *PhysicsP2Body) SetZeroDampingI(args ...interface{}) {
    self.Object.Call("setZeroDamping", args)
}

// ToLocalFrame Transform a world point to local body frame.
func (self *PhysicsP2Body) ToLocalFrame(out interface{}, worldPoint interface{}) {
    self.Object.Call("toLocalFrame", out, worldPoint)
}

// ToLocalFrameI Transform a world point to local body frame.
func (self *PhysicsP2Body) ToLocalFrameI(args ...interface{}) {
    self.Object.Call("toLocalFrame", args)
}

// ToWorldFrame Transform a local point to world frame.
func (self *PhysicsP2Body) ToWorldFrame(out []interface{}, localPoint []interface{}) {
    self.Object.Call("toWorldFrame", out, localPoint)
}

// ToWorldFrameI Transform a local point to world frame.
func (self *PhysicsP2Body) ToWorldFrameI(args ...interface{}) {
    self.Object.Call("toWorldFrame", args)
}

// RotateLeft This will rotate the Body by the given speed to the left (counter-clockwise).
func (self *PhysicsP2Body) RotateLeft(speed int) {
    self.Object.Call("rotateLeft", speed)
}

// RotateLeftI This will rotate the Body by the given speed to the left (counter-clockwise).
func (self *PhysicsP2Body) RotateLeftI(args ...interface{}) {
    self.Object.Call("rotateLeft", args)
}

// RotateRight This will rotate the Body by the given speed to the left (clockwise).
func (self *PhysicsP2Body) RotateRight(speed int) {
    self.Object.Call("rotateRight", speed)
}

// RotateRightI This will rotate the Body by the given speed to the left (clockwise).
func (self *PhysicsP2Body) RotateRightI(args ...interface{}) {
    self.Object.Call("rotateRight", args)
}

// MoveForward Moves the Body forwards based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveForward(speed int) {
    self.Object.Call("moveForward", speed)
}

// MoveForwardI Moves the Body forwards based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveForwardI(args ...interface{}) {
    self.Object.Call("moveForward", args)
}

// MoveBackward Moves the Body backwards based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveBackward(speed int) {
    self.Object.Call("moveBackward", speed)
}

// MoveBackwardI Moves the Body backwards based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveBackwardI(args ...interface{}) {
    self.Object.Call("moveBackward", args)
}

// Thrust Applies a force to the Body that causes it to 'thrust' forwards, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) Thrust(speed int) {
    self.Object.Call("thrust", speed)
}

// ThrustI Applies a force to the Body that causes it to 'thrust' forwards, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustI(args ...interface{}) {
    self.Object.Call("thrust", args)
}

// ThrustLeft Applies a force to the Body that causes it to 'thrust' to the left, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustLeft(speed int) {
    self.Object.Call("thrustLeft", speed)
}

// ThrustLeftI Applies a force to the Body that causes it to 'thrust' to the left, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustLeftI(args ...interface{}) {
    self.Object.Call("thrustLeft", args)
}

// ThrustRight Applies a force to the Body that causes it to 'thrust' to the right, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustRight(speed int) {
    self.Object.Call("thrustRight", speed)
}

// ThrustRightI Applies a force to the Body that causes it to 'thrust' to the right, based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ThrustRightI(args ...interface{}) {
    self.Object.Call("thrustRight", args)
}

// Reverse Applies a force to the Body that causes it to 'thrust' backwards (in reverse), based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) Reverse(speed int) {
    self.Object.Call("reverse", speed)
}

// ReverseI Applies a force to the Body that causes it to 'thrust' backwards (in reverse), based on its current angle and the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) ReverseI(args ...interface{}) {
    self.Object.Call("reverse", args)
}

// MoveLeft If this Body is dynamic then this will move it to the left by setting its x velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveLeft(speed int) {
    self.Object.Call("moveLeft", speed)
}

// MoveLeftI If this Body is dynamic then this will move it to the left by setting its x velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveLeftI(args ...interface{}) {
    self.Object.Call("moveLeft", args)
}

// MoveRight If this Body is dynamic then this will move it to the right by setting its x velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveRight(speed int) {
    self.Object.Call("moveRight", speed)
}

// MoveRightI If this Body is dynamic then this will move it to the right by setting its x velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveRightI(args ...interface{}) {
    self.Object.Call("moveRight", args)
}

// MoveUp If this Body is dynamic then this will move it up by setting its y velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveUp(speed int) {
    self.Object.Call("moveUp", speed)
}

// MoveUpI If this Body is dynamic then this will move it up by setting its y velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveUpI(args ...interface{}) {
    self.Object.Call("moveUp", args)
}

// MoveDown If this Body is dynamic then this will move it down by setting its y velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveDown(speed int) {
    self.Object.Call("moveDown", speed)
}

// MoveDownI If this Body is dynamic then this will move it down by setting its y velocity to the given speed.
// The speed is represented in pixels per second. So a value of 100 would move 100 pixels in 1 second (1000ms).
func (self *PhysicsP2Body) MoveDownI(args ...interface{}) {
    self.Object.Call("moveDown", args)
}

// PreUpdate Internal method. This is called directly before the sprites are sent to the renderer and after the update function has finished.
func (self *PhysicsP2Body) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI Internal method. This is called directly before the sprites are sent to the renderer and after the update function has finished.
func (self *PhysicsP2Body) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// PostUpdate Internal method. This is called directly before the sprites are sent to the renderer and after the update function has finished.
func (self *PhysicsP2Body) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI Internal method. This is called directly before the sprites are sent to the renderer and after the update function has finished.
func (self *PhysicsP2Body) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Reset Resets the Body force, velocity (linear and angular) and rotation. Optionally resets damping and mass.
func (self *PhysicsP2Body) Reset(x int, y int) {
    self.Object.Call("reset", x, y)
}

// Reset1O Resets the Body force, velocity (linear and angular) and rotation. Optionally resets damping and mass.
func (self *PhysicsP2Body) Reset1O(x int, y int, resetDamping bool) {
    self.Object.Call("reset", x, y, resetDamping)
}

// Reset2O Resets the Body force, velocity (linear and angular) and rotation. Optionally resets damping and mass.
func (self *PhysicsP2Body) Reset2O(x int, y int, resetDamping bool, resetMass bool) {
    self.Object.Call("reset", x, y, resetDamping, resetMass)
}

// ResetI Resets the Body force, velocity (linear and angular) and rotation. Optionally resets damping and mass.
func (self *PhysicsP2Body) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// AddToWorld Adds this physics body to the world.
func (self *PhysicsP2Body) AddToWorld() {
    self.Object.Call("addToWorld")
}

// AddToWorldI Adds this physics body to the world.
func (self *PhysicsP2Body) AddToWorldI(args ...interface{}) {
    self.Object.Call("addToWorld", args)
}

// RemoveFromWorld Removes this physics body from the world.
func (self *PhysicsP2Body) RemoveFromWorld() {
    self.Object.Call("removeFromWorld")
}

// RemoveFromWorldI Removes this physics body from the world.
func (self *PhysicsP2Body) RemoveFromWorldI(args ...interface{}) {
    self.Object.Call("removeFromWorld", args)
}

// Destroy Destroys this Body and all references it holds to other objects.
func (self *PhysicsP2Body) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this Body and all references it holds to other objects.
func (self *PhysicsP2Body) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// ClearShapes Removes all Shapes from this Body.
func (self *PhysicsP2Body) ClearShapes() {
    self.Object.Call("clearShapes")
}

// ClearShapesI Removes all Shapes from this Body.
func (self *PhysicsP2Body) ClearShapesI(args ...interface{}) {
    self.Object.Call("clearShapes", args)
}

// AddShape Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShape(shape *P2Shape) *P2Shape{
    return &P2Shape{self.Object.Call("addShape", shape)}
}

// AddShape1O Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShape1O(shape *P2Shape, offsetX int) *P2Shape{
    return &P2Shape{self.Object.Call("addShape", shape, offsetX)}
}

// AddShape2O Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShape2O(shape *P2Shape, offsetX int, offsetY int) *P2Shape{
    return &P2Shape{self.Object.Call("addShape", shape, offsetX, offsetY)}
}

// AddShape3O Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShape3O(shape *P2Shape, offsetX int, offsetY int, rotation int) *P2Shape{
    return &P2Shape{self.Object.Call("addShape", shape, offsetX, offsetY, rotation)}
}

// AddShapeI Add a shape to the body. You can pass a local transform when adding a shape, so that the shape gets an offset and an angle relative to the body center of mass.
// Will automatically update the mass properties and bounding radius.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) AddShapeI(args ...interface{}) *P2Shape{
    return &P2Shape{self.Object.Call("addShape", args)}
}

// AddCircle Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircle(radius int) *P2Circle{
    return &P2Circle{self.Object.Call("addCircle", radius)}
}

// AddCircle1O Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircle1O(radius int, offsetX int) *P2Circle{
    return &P2Circle{self.Object.Call("addCircle", radius, offsetX)}
}

// AddCircle2O Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircle2O(radius int, offsetX int, offsetY int) *P2Circle{
    return &P2Circle{self.Object.Call("addCircle", radius, offsetX, offsetY)}
}

// AddCircle3O Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircle3O(radius int, offsetX int, offsetY int, rotation int) *P2Circle{
    return &P2Circle{self.Object.Call("addCircle", radius, offsetX, offsetY, rotation)}
}

// AddCircleI Adds a Circle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCircleI(args ...interface{}) *P2Circle{
    return &P2Circle{self.Object.Call("addCircle", args)}
}

// AddRectangle Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangle(width int, height int) *P2Box{
    return &P2Box{self.Object.Call("addRectangle", width, height)}
}

// AddRectangle1O Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangle1O(width int, height int, offsetX int) *P2Box{
    return &P2Box{self.Object.Call("addRectangle", width, height, offsetX)}
}

// AddRectangle2O Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangle2O(width int, height int, offsetX int, offsetY int) *P2Box{
    return &P2Box{self.Object.Call("addRectangle", width, height, offsetX, offsetY)}
}

// AddRectangle3O Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangle3O(width int, height int, offsetX int, offsetY int, rotation int) *P2Box{
    return &P2Box{self.Object.Call("addRectangle", width, height, offsetX, offsetY, rotation)}
}

// AddRectangleI Adds a Rectangle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddRectangleI(args ...interface{}) *P2Box{
    return &P2Box{self.Object.Call("addRectangle", args)}
}

// AddPlane Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlane() *P2Plane{
    return &P2Plane{self.Object.Call("addPlane")}
}

// AddPlane1O Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlane1O(offsetX int) *P2Plane{
    return &P2Plane{self.Object.Call("addPlane", offsetX)}
}

// AddPlane2O Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlane2O(offsetX int, offsetY int) *P2Plane{
    return &P2Plane{self.Object.Call("addPlane", offsetX, offsetY)}
}

// AddPlane3O Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlane3O(offsetX int, offsetY int, rotation int) *P2Plane{
    return &P2Plane{self.Object.Call("addPlane", offsetX, offsetY, rotation)}
}

// AddPlaneI Adds a Plane shape to this Body. The plane is facing in the Y direction. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddPlaneI(args ...interface{}) *P2Plane{
    return &P2Plane{self.Object.Call("addPlane", args)}
}

// AddParticle Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticle() *P2Particle{
    return &P2Particle{self.Object.Call("addParticle")}
}

// AddParticle1O Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticle1O(offsetX int) *P2Particle{
    return &P2Particle{self.Object.Call("addParticle", offsetX)}
}

// AddParticle2O Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticle2O(offsetX int, offsetY int) *P2Particle{
    return &P2Particle{self.Object.Call("addParticle", offsetX, offsetY)}
}

// AddParticle3O Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticle3O(offsetX int, offsetY int, rotation int) *P2Particle{
    return &P2Particle{self.Object.Call("addParticle", offsetX, offsetY, rotation)}
}

// AddParticleI Adds a Particle shape to this Body. You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddParticleI(args ...interface{}) *P2Particle{
    return &P2Particle{self.Object.Call("addParticle", args)}
}

// AddLine Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLine(length int) *P2Line{
    return &P2Line{self.Object.Call("addLine", length)}
}

// AddLine1O Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLine1O(length int, offsetX int) *P2Line{
    return &P2Line{self.Object.Call("addLine", length, offsetX)}
}

// AddLine2O Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLine2O(length int, offsetX int, offsetY int) *P2Line{
    return &P2Line{self.Object.Call("addLine", length, offsetX, offsetY)}
}

// AddLine3O Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLine3O(length int, offsetX int, offsetY int, rotation int) *P2Line{
    return &P2Line{self.Object.Call("addLine", length, offsetX, offsetY, rotation)}
}

// AddLineI Adds a Line shape to this Body.
// The line shape is along the x direction, and stretches from [-length/2, 0] to [length/2,0].
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddLineI(args ...interface{}) *P2Line{
    return &P2Line{self.Object.Call("addLine", args)}
}

// AddCapsule Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsule(length int, radius int) *P2Capsule{
    return &P2Capsule{self.Object.Call("addCapsule", length, radius)}
}

// AddCapsule1O Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsule1O(length int, radius int, offsetX int) *P2Capsule{
    return &P2Capsule{self.Object.Call("addCapsule", length, radius, offsetX)}
}

// AddCapsule2O Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsule2O(length int, radius int, offsetX int, offsetY int) *P2Capsule{
    return &P2Capsule{self.Object.Call("addCapsule", length, radius, offsetX, offsetY)}
}

// AddCapsule3O Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsule3O(length int, radius int, offsetX int, offsetY int, rotation int) *P2Capsule{
    return &P2Capsule{self.Object.Call("addCapsule", length, radius, offsetX, offsetY, rotation)}
}

// AddCapsuleI Adds a Capsule shape to this Body.
// You can control the offset from the center of the body and the rotation.
func (self *PhysicsP2Body) AddCapsuleI(args ...interface{}) *P2Capsule{
    return &P2Capsule{self.Object.Call("addCapsule", args)}
}

// AddPolygon Reads a polygon shape path, and assembles convex shapes from that and puts them at proper offset points. The shape must be simple and without holes.
// This function expects the x.y values to be given in pixels. If you want to provide them at p2 world scales then call Body.data.fromPolygon directly.
func (self *PhysicsP2Body) AddPolygon(options interface{}, points interface{}) bool{
    return self.Object.Call("addPolygon", options, points).Bool()
}

// AddPolygonI Reads a polygon shape path, and assembles convex shapes from that and puts them at proper offset points. The shape must be simple and without holes.
// This function expects the x.y values to be given in pixels. If you want to provide them at p2 world scales then call Body.data.fromPolygon directly.
func (self *PhysicsP2Body) AddPolygonI(args ...interface{}) bool{
    return self.Object.Call("addPolygon", args).Bool()
}

// RemoveShape Remove a shape from the body. Will automatically update the mass properties and bounding radius.
func (self *PhysicsP2Body) RemoveShape(shape interface{}) bool{
    return self.Object.Call("removeShape", shape).Bool()
}

// RemoveShapeI Remove a shape from the body. Will automatically update the mass properties and bounding radius.
func (self *PhysicsP2Body) RemoveShapeI(args ...interface{}) bool{
    return self.Object.Call("removeShape", args).Bool()
}

// SetCircle Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircle(radius int) {
    self.Object.Call("setCircle", radius)
}

// SetCircle1O Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircle1O(radius int, offsetX int) {
    self.Object.Call("setCircle", radius, offsetX)
}

// SetCircle2O Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircle2O(radius int, offsetX int, offsetY int) {
    self.Object.Call("setCircle", radius, offsetX, offsetY)
}

// SetCircle3O Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircle3O(radius int, offsetX int, offsetY int, rotation int) {
    self.Object.Call("setCircle", radius, offsetX, offsetY, rotation)
}

// SetCircleI Clears any previously set shapes. Then creates a new Circle shape and adds it to this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetCircleI(args ...interface{}) {
    self.Object.Call("setCircle", args)
}

// SetRectangle Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle() *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle")}
}

// SetRectangle1O Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle1O(width int) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", width)}
}

// SetRectangle2O Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle2O(width int, height int) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", width, height)}
}

// SetRectangle3O Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle3O(width int, height int, offsetX int) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", width, height, offsetX)}
}

// SetRectangle4O Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle4O(width int, height int, offsetX int, offsetY int) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", width, height, offsetX, offsetY)}
}

// SetRectangle5O Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangle5O(width int, height int, offsetX int, offsetY int, rotation int) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", width, height, offsetX, offsetY, rotation)}
}

// SetRectangleI Clears any previously set shapes. The creates a new Rectangle shape at the given size and offset, and adds it to this Body.
// If you wish to create a Rectangle to match the size of a Sprite or Image see Body.setRectangleFromSprite.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangleI(args ...interface{}) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangle", args)}
}

// SetRectangleFromSprite Clears any previously set shapes.
// Then creates a Rectangle shape sized to match the dimensions and orientation of the Sprite given.
// If no Sprite is given it defaults to using the parent of this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangleFromSprite() *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangleFromSprite")}
}

// SetRectangleFromSprite1O Clears any previously set shapes.
// Then creates a Rectangle shape sized to match the dimensions and orientation of the Sprite given.
// If no Sprite is given it defaults to using the parent of this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangleFromSprite1O(sprite interface{}) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangleFromSprite", sprite)}
}

// SetRectangleFromSpriteI Clears any previously set shapes.
// Then creates a Rectangle shape sized to match the dimensions and orientation of the Sprite given.
// If no Sprite is given it defaults to using the parent of this Body.
// If this Body had a previously set Collision Group you will need to re-apply it to the new Shape this creates.
func (self *PhysicsP2Body) SetRectangleFromSpriteI(args ...interface{}) *P2Rectangle{
    return &P2Rectangle{self.Object.Call("setRectangleFromSprite", args)}
}

// SetMaterial Adds the given Material to all Shapes that belong to this Body.
// If you only wish to apply it to a specific Shape in this Body then provide that as the 2nd parameter.
func (self *PhysicsP2Body) SetMaterial(material *PhysicsP2Material) {
    self.Object.Call("setMaterial", material)
}

// SetMaterial1O Adds the given Material to all Shapes that belong to this Body.
// If you only wish to apply it to a specific Shape in this Body then provide that as the 2nd parameter.
func (self *PhysicsP2Body) SetMaterial1O(material *PhysicsP2Material, shape *P2Shape) {
    self.Object.Call("setMaterial", material, shape)
}

// SetMaterialI Adds the given Material to all Shapes that belong to this Body.
// If you only wish to apply it to a specific Shape in this Body then provide that as the 2nd parameter.
func (self *PhysicsP2Body) SetMaterialI(args ...interface{}) {
    self.Object.Call("setMaterial", args)
}

// ShapeChanged Updates the debug draw if any body shapes change.
func (self *PhysicsP2Body) ShapeChanged() {
    self.Object.Call("shapeChanged")
}

// ShapeChangedI Updates the debug draw if any body shapes change.
func (self *PhysicsP2Body) ShapeChangedI(args ...interface{}) {
    self.Object.Call("shapeChanged", args)
}

// AddPhaserPolygon Reads the shape data from a physics data file stored in the Game.Cache and adds it as a polygon to this Body.
// The shape data format is based on the output of the
// {@link https://github.com/photonstorm/phaser/tree/master/resources/PhysicsEditor%20Exporter|custom phaser exporter} for
// {@link https://www.codeandweb.com/physicseditor|PhysicsEditor}
func (self *PhysicsP2Body) AddPhaserPolygon(key string, object string) []interface{}{
	array00 := self.Object.Call("addPhaserPolygon", key, object)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// AddPhaserPolygonI Reads the shape data from a physics data file stored in the Game.Cache and adds it as a polygon to this Body.
// The shape data format is based on the output of the
// {@link https://github.com/photonstorm/phaser/tree/master/resources/PhysicsEditor%20Exporter|custom phaser exporter} for
// {@link https://www.codeandweb.com/physicseditor|PhysicsEditor}
func (self *PhysicsP2Body) AddPhaserPolygonI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("addPhaserPolygon", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// AddFixture Add a polygon fixture. This is used during #loadPolygon.
func (self *PhysicsP2Body) AddFixture(fixtureData string) []interface{}{
	array00 := self.Object.Call("addFixture", fixtureData)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// AddFixtureI Add a polygon fixture. This is used during #loadPolygon.
func (self *PhysicsP2Body) AddFixtureI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("addFixture", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// LoadPolygon Reads the shape data from a physics data file stored in the Game.Cache and adds it as a polygon to this Body.
// 
// As well as reading the data from the Cache you can also pass `null` as the first argument and a
// physics data object as the second. When doing this you must ensure the structure of the object is correct in advance.
// 
// For more details see the format of the Lime / Corona Physics Editor export.
func (self *PhysicsP2Body) LoadPolygon(key string, object interface{}) bool{
    return self.Object.Call("loadPolygon", key, object).Bool()
}

// LoadPolygonI Reads the shape data from a physics data file stored in the Game.Cache and adds it as a polygon to this Body.
// 
// As well as reading the data from the Cache you can also pass `null` as the first argument and a
// physics data object as the second. When doing this you must ensure the structure of the object is correct in advance.
// 
// For more details see the format of the Lime / Corona Physics Editor export.
func (self *PhysicsP2Body) LoadPolygonI(args ...interface{}) bool{
    return self.Object.Call("loadPolygon", args).Bool()
}

