// Package phaser Automatic generation for Phaser.Physics
// generated file Physics.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Physics The Physics Manager is responsible for looking after all of the running physics systems.
// Phaser supports 4 physics systems: Arcade Physics, P2, Ninja Physics and Box2D via a commercial plugin.
// 
// Game Objects (such as Sprites) can only belong to 1 physics system, but you can have multiple systems active in a single game.
// 
// For example you could have P2 managing a polygon-built terrain landscape that an vehicle drives over, while it could be firing bullets that use the
// faster (due to being much simpler) Arcade Physics system.
type Physics struct {
    *js.Object
}

// NewPhysics The Physics Manager is responsible for looking after all of the running physics systems.
// Phaser supports 4 physics systems: Arcade Physics, P2, Ninja Physics and Box2D via a commercial plugin.
// 
// Game Objects (such as Sprites) can only belong to 1 physics system, but you can have multiple systems active in a single game.
// 
// For example you could have P2 managing a polygon-built terrain landscape that an vehicle drives over, while it could be firing bullets that use the
// faster (due to being much simpler) Arcade Physics system.
func NewPhysics(game *Game) *Physics {
    return &Physics{js.Global.Get("Phaser").Get("Physics").New(game)}
}
// NewPhysics1O The Physics Manager is responsible for looking after all of the running physics systems.
// Phaser supports 4 physics systems: Arcade Physics, P2, Ninja Physics and Box2D via a commercial plugin.
// 
// Game Objects (such as Sprites) can only belong to 1 physics system, but you can have multiple systems active in a single game.
// 
// For example you could have P2 managing a polygon-built terrain landscape that an vehicle drives over, while it could be firing bullets that use the
// faster (due to being much simpler) Arcade Physics system.
func NewPhysics1O(game *Game, physicsConfig interface{}) *Physics {
    return &Physics{js.Global.Get("Phaser").Get("Physics").New(game, physicsConfig)}
}
// NewPhysicsI The Physics Manager is responsible for looking after all of the running physics systems.
// Phaser supports 4 physics systems: Arcade Physics, P2, Ninja Physics and Box2D via a commercial plugin.
// 
// Game Objects (such as Sprites) can only belong to 1 physics system, but you can have multiple systems active in a single game.
// 
// For example you could have P2 managing a polygon-built terrain landscape that an vehicle drives over, while it could be firing bullets that use the
// faster (due to being much simpler) Arcade Physics system.
func NewPhysicsI(args ...interface{}) *Physics {
    return &Physics{js.Global.Get("Phaser").Get("Physics").New(args)}
}



// Physics Binding conversion method to Physics point 
func ToPhysics(jsStruct interface{}) *Physics {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Physics{Object: object}
	}
	return nil
}



// Game Local reference to game.
func (self *Physics) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *Physics) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Config The physics configuration object as passed to the game on creation.
func (self *Physics) Config() interface{}{
    return self.Object.Get("config")
}

// SetConfigA The physics configuration object as passed to the game on creation.
func (self *Physics) SetConfigA(member interface{}) {
    self.Object.Set("config", member)
}

// Arcade The Arcade Physics system.
func (self *Physics) Arcade() *PhysicsArcade{
    return &PhysicsArcade{self.Object.Get("arcade")}
}

// SetArcadeA The Arcade Physics system.
func (self *Physics) SetArcadeA(member *PhysicsArcade) {
    self.Object.Set("arcade", member)
}

// P2 The P2.JS Physics system.
func (self *Physics) P2() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("p2")}
}

// SetP2A The P2.JS Physics system.
func (self *Physics) SetP2A(member *PhysicsP2) {
    self.Object.Set("p2", member)
}

// Ninja The N+ Ninja Physics system.
func (self *Physics) Ninja() *PhysicsNinja{
    return &PhysicsNinja{self.Object.Get("ninja")}
}

// SetNinjaA The N+ Ninja Physics system.
func (self *Physics) SetNinjaA(member *PhysicsNinja) {
    self.Object.Set("ninja", member)
}

// Box2d The Box2D Physics system.
func (self *Physics) Box2d() *PhysicsBox2D{
    return &PhysicsBox2D{self.Object.Get("box2d")}
}

// SetBox2dA The Box2D Physics system.
func (self *Physics) SetBox2dA(member *PhysicsBox2D) {
    self.Object.Set("box2d", member)
}

// Chipmunk The Chipmunk Physics system (to be done).
func (self *Physics) Chipmunk() *PhysicsChipmunk{
    return &PhysicsChipmunk{self.Object.Get("chipmunk")}
}

// SetChipmunkA The Chipmunk Physics system (to be done).
func (self *Physics) SetChipmunkA(member *PhysicsChipmunk) {
    self.Object.Set("chipmunk", member)
}

// Matter The MatterJS Physics system (coming soon).
func (self *Physics) Matter() *PhysicsMatter{
    return &PhysicsMatter{self.Object.Get("matter")}
}

// SetMatterA The MatterJS Physics system (coming soon).
func (self *Physics) SetMatterA(member *PhysicsMatter) {
    self.Object.Set("matter", member)
}

// ARCADE empty description
func (self *Physics) ARCADE() int{
    return self.Object.Get("ARCADE").Int()
}

// SetARCADEA empty description
func (self *Physics) SetARCADEA(member int) {
    self.Object.Set("ARCADE", member)
}

// P2JS empty description
func (self *Physics) P2JS() int{
    return self.Object.Get("P2JS").Int()
}

// SetP2JSA empty description
func (self *Physics) SetP2JSA(member int) {
    self.Object.Set("P2JS", member)
}

// NINJA empty description
func (self *Physics) NINJA() int{
    return self.Object.Get("NINJA").Int()
}

// SetNINJAA empty description
func (self *Physics) SetNINJAA(member int) {
    self.Object.Set("NINJA", member)
}

// BOX2D empty description
func (self *Physics) BOX2D() int{
    return self.Object.Get("BOX2D").Int()
}

// SetBOX2DA empty description
func (self *Physics) SetBOX2DA(member int) {
    self.Object.Set("BOX2D", member)
}

// CHIPMUNK empty description
func (self *Physics) CHIPMUNK() int{
    return self.Object.Get("CHIPMUNK").Int()
}

// SetCHIPMUNKA empty description
func (self *Physics) SetCHIPMUNKA(member int) {
    self.Object.Set("CHIPMUNK", member)
}

// MATTERJS empty description
func (self *Physics) MATTERJS() int{
    return self.Object.Get("MATTERJS").Int()
}

// SetMATTERJSA empty description
func (self *Physics) SetMATTERJSA(member int) {
    self.Object.Set("MATTERJS", member)
}


// ParseConfig Parses the Physics Configuration object passed to the Game constructor and starts any physics systems specified within.
func (self *Physics) ParseConfig() {
    self.Object.Call("parseConfig")
}

// ParseConfigI Parses the Physics Configuration object passed to the Game constructor and starts any physics systems specified within.
func (self *Physics) ParseConfigI(args ...interface{}) {
    self.Object.Call("parseConfig", args)
}

// StartSystem This will create an instance of the requested physics simulation.
// Phaser.Physics.Arcade is running by default, but all others need activating directly.
// 
// You can start the following physics systems:
// 
// Phaser.Physics.P2JS - A full-body advanced physics system by Stefan Hedman.
// Phaser.Physics.NINJA - A port of Metanet Softwares N+ physics system.
// Phaser.Physics.BOX2D - A commercial Phaser Plugin (see http://phaser.io)
// 
// Both Ninja Physics and Box2D require their respective plugins to be loaded before you can start them.
// They are not bundled into the core Phaser library.
// 
// If the physics world has already been created (i.e. in another state in your game) then 
// calling startSystem will reset the physics world, not re-create it. If you need to start them again from their constructors 
// then set Phaser.Physics.p2 (or whichever system you want to recreate) to `null` before calling `startSystem`.
func (self *Physics) StartSystem(system int) {
    self.Object.Call("startSystem", system)
}

// StartSystemI This will create an instance of the requested physics simulation.
// Phaser.Physics.Arcade is running by default, but all others need activating directly.
// 
// You can start the following physics systems:
// 
// Phaser.Physics.P2JS - A full-body advanced physics system by Stefan Hedman.
// Phaser.Physics.NINJA - A port of Metanet Softwares N+ physics system.
// Phaser.Physics.BOX2D - A commercial Phaser Plugin (see http://phaser.io)
// 
// Both Ninja Physics and Box2D require their respective plugins to be loaded before you can start them.
// They are not bundled into the core Phaser library.
// 
// If the physics world has already been created (i.e. in another state in your game) then 
// calling startSystem will reset the physics world, not re-create it. If you need to start them again from their constructors 
// then set Phaser.Physics.p2 (or whichever system you want to recreate) to `null` before calling `startSystem`.
func (self *Physics) StartSystemI(args ...interface{}) {
    self.Object.Call("startSystem", args)
}

// Enable This will create a default physics body on the given game object or array of objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// It can be for any of the physics systems that have been started:
// 
// Phaser.Physics.Arcade - A light weight AABB based collision system with basic separation.
// Phaser.Physics.P2JS - A full-body advanced physics system supporting multiple object shapes, polygon loading, contact materials, springs and constraints.
// Phaser.Physics.NINJA - A port of Metanet Softwares N+ physics system. Advanced AABB and Circle vs. Tile collision.
// Phaser.Physics.BOX2D - A port of https://code.google.com/p/box2d-html5
// Phaser.Physics.MATTER - A full-body and light-weight advanced physics system (still in development)
// Phaser.Physics.CHIPMUNK is still in development.
// 
// If you require more control over what type of body is created, for example to create a Ninja Physics Circle instead of the default AABB, then see the
// individual physics systems `enable` methods instead of using this generic one.
func (self *Physics) Enable(object interface{}) {
    self.Object.Call("enable", object)
}

// Enable1O This will create a default physics body on the given game object or array of objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// It can be for any of the physics systems that have been started:
// 
// Phaser.Physics.Arcade - A light weight AABB based collision system with basic separation.
// Phaser.Physics.P2JS - A full-body advanced physics system supporting multiple object shapes, polygon loading, contact materials, springs and constraints.
// Phaser.Physics.NINJA - A port of Metanet Softwares N+ physics system. Advanced AABB and Circle vs. Tile collision.
// Phaser.Physics.BOX2D - A port of https://code.google.com/p/box2d-html5
// Phaser.Physics.MATTER - A full-body and light-weight advanced physics system (still in development)
// Phaser.Physics.CHIPMUNK is still in development.
// 
// If you require more control over what type of body is created, for example to create a Ninja Physics Circle instead of the default AABB, then see the
// individual physics systems `enable` methods instead of using this generic one.
func (self *Physics) Enable1O(object interface{}, system int) {
    self.Object.Call("enable", object, system)
}

// Enable2O This will create a default physics body on the given game object or array of objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// It can be for any of the physics systems that have been started:
// 
// Phaser.Physics.Arcade - A light weight AABB based collision system with basic separation.
// Phaser.Physics.P2JS - A full-body advanced physics system supporting multiple object shapes, polygon loading, contact materials, springs and constraints.
// Phaser.Physics.NINJA - A port of Metanet Softwares N+ physics system. Advanced AABB and Circle vs. Tile collision.
// Phaser.Physics.BOX2D - A port of https://code.google.com/p/box2d-html5
// Phaser.Physics.MATTER - A full-body and light-weight advanced physics system (still in development)
// Phaser.Physics.CHIPMUNK is still in development.
// 
// If you require more control over what type of body is created, for example to create a Ninja Physics Circle instead of the default AABB, then see the
// individual physics systems `enable` methods instead of using this generic one.
func (self *Physics) Enable2O(object interface{}, system int, debug bool) {
    self.Object.Call("enable", object, system, debug)
}

// EnableI This will create a default physics body on the given game object or array of objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
// It can be for any of the physics systems that have been started:
// 
// Phaser.Physics.Arcade - A light weight AABB based collision system with basic separation.
// Phaser.Physics.P2JS - A full-body advanced physics system supporting multiple object shapes, polygon loading, contact materials, springs and constraints.
// Phaser.Physics.NINJA - A port of Metanet Softwares N+ physics system. Advanced AABB and Circle vs. Tile collision.
// Phaser.Physics.BOX2D - A port of https://code.google.com/p/box2d-html5
// Phaser.Physics.MATTER - A full-body and light-weight advanced physics system (still in development)
// Phaser.Physics.CHIPMUNK is still in development.
// 
// If you require more control over what type of body is created, for example to create a Ninja Physics Circle instead of the default AABB, then see the
// individual physics systems `enable` methods instead of using this generic one.
func (self *Physics) EnableI(args ...interface{}) {
    self.Object.Call("enable", args)
}

// PreUpdate preUpdate checks.
func (self *Physics) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI preUpdate checks.
func (self *Physics) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Update Updates all running physics systems.
func (self *Physics) Update() {
    self.Object.Call("update")
}

// UpdateI Updates all running physics systems.
func (self *Physics) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// SetBoundsToWorld Updates the physics bounds to match the world dimensions.
func (self *Physics) SetBoundsToWorld() {
    self.Object.Call("setBoundsToWorld")
}

// SetBoundsToWorldI Updates the physics bounds to match the world dimensions.
func (self *Physics) SetBoundsToWorldI(args ...interface{}) {
    self.Object.Call("setBoundsToWorld", args)
}

// Clear Clears down all active physics systems. This doesn't destroy them, it just clears them of objects and is called when the State changes.
func (self *Physics) Clear() {
    self.Object.Call("clear")
}

// ClearI Clears down all active physics systems. This doesn't destroy them, it just clears them of objects and is called when the State changes.
func (self *Physics) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// Reset Resets the active physics system. Called automatically on a Phaser.State swap.
func (self *Physics) Reset() {
    self.Object.Call("reset")
}

// ResetI Resets the active physics system. Called automatically on a Phaser.State swap.
func (self *Physics) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Destroy Destroys all active physics systems. Usually only called on a Game Shutdown, not on a State swap.
func (self *Physics) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys all active physics systems. Usually only called on a Game Shutdown, not on a State swap.
func (self *Physics) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

