// Automatic generation for Phaser.Physics
// generated file Physics.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Physics Manager is responsible for looking after all of the running physics systems.
// Phaser supports 4 physics systems: Arcade Physics, P2, Ninja Physics and Box2D via a commercial plugin.
// 
// Game Objects (such as Sprites) can only belong to 1 physics system, but you can have multiple systems active in a single game.
// 
// For example you could have P2 managing a polygon-built terrain landscape that an vehicle drives over, while it could be firing bullets that use the
// faster (due to being much simpler) Arcade Physics system.
type Physics struct {
    *js.Object
}


// Local reference to game.
func (self *Physics) GetGame() Game{
    return Game{self.Get("game")}
}

// Local reference to game.
func (self *Physics) SetGame(member Game) {
    self.Set("game", member)
}

// The physics configuration object as passed to the game on creation.
func (self *Physics) GetConfig() interface{}{
    return self.Get("config")
}

// The physics configuration object as passed to the game on creation.
func (self *Physics) SetConfig(member interface{}) {
    self.Set("config", member)
}

// The Arcade Physics system.
func (self *Physics) GetArcade() PhysicsArcade{
    return PhysicsArcade{self.Get("arcade")}
}

// The Arcade Physics system.
func (self *Physics) SetArcade(member PhysicsArcade) {
    self.Set("arcade", member)
}

// The P2.JS Physics system.
func (self *Physics) GetP2() PhysicsP2{
    return PhysicsP2{self.Get("p2")}
}

// The P2.JS Physics system.
func (self *Physics) SetP2(member PhysicsP2) {
    self.Set("p2", member)
}

// The N+ Ninja Physics system.
func (self *Physics) GetNinja() PhysicsNinja{
    return PhysicsNinja{self.Get("ninja")}
}

// The N+ Ninja Physics system.
func (self *Physics) SetNinja(member PhysicsNinja) {
    self.Set("ninja", member)
}

// The Box2D Physics system.
func (self *Physics) GetBox2d() PhysicsBox2D{
    return PhysicsBox2D{self.Get("box2d")}
}

// The Box2D Physics system.
func (self *Physics) SetBox2d(member PhysicsBox2D) {
    self.Set("box2d", member)
}

// The Chipmunk Physics system (to be done).
func (self *Physics) GetChipmunk() PhysicsChipmunk{
    return PhysicsChipmunk{self.Get("chipmunk")}
}

// The Chipmunk Physics system (to be done).
func (self *Physics) SetChipmunk(member PhysicsChipmunk) {
    self.Set("chipmunk", member)
}

// The MatterJS Physics system (coming soon).
func (self *Physics) GetMatter() PhysicsMatter{
    return PhysicsMatter{self.Get("matter")}
}

// The MatterJS Physics system (coming soon).
func (self *Physics) SetMatter(member PhysicsMatter) {
    self.Set("matter", member)
}

// 
func (self *Physics) GetARCADE() float64{
    return self.Get("ARCADE").Float()
}

// 
func (self *Physics) SetARCADE(member float64) {
    self.Set("ARCADE", member)
}

// 
func (self *Physics) GetP2JS() float64{
    return self.Get("P2JS").Float()
}

// 
func (self *Physics) SetP2JS(member float64) {
    self.Set("P2JS", member)
}

// 
func (self *Physics) GetNINJA() float64{
    return self.Get("NINJA").Float()
}

// 
func (self *Physics) SetNINJA(member float64) {
    self.Set("NINJA", member)
}

// 
func (self *Physics) GetBOX2D() float64{
    return self.Get("BOX2D").Float()
}

// 
func (self *Physics) SetBOX2D(member float64) {
    self.Set("BOX2D", member)
}

// 
func (self *Physics) GetCHIPMUNK() float64{
    return self.Get("CHIPMUNK").Float()
}

// 
func (self *Physics) SetCHIPMUNK(member float64) {
    self.Set("CHIPMUNK", member)
}

// 
func (self *Physics) GetMATTERJS() float64{
    return self.Get("MATTERJS").Float()
}

// 
func (self *Physics) SetMATTERJS(member float64) {
    self.Set("MATTERJS", member)
}



// Parses the Physics Configuration object passed to the Game constructor and starts any physics systems specified within.
func (self *Physics) ParseConfigI(args ...interface{}) {
    self.Call("parseConfig", args)
}

// This will create an instance of the requested physics simulation.
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
    self.Call("startSystem", args)
}

// This will create a default physics body on the given game object or array of objects.
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
    self.Call("enable", args)
}

// preUpdate checks.
func (self *Physics) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Updates all running physics systems.
func (self *Physics) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Updates the physics bounds to match the world dimensions.
func (self *Physics) SetBoundsToWorldI(args ...interface{}) {
    self.Call("setBoundsToWorld", args)
}

// Clears down all active physics systems. This doesn't destroy them, it just clears them of objects and is called when the State changes.
func (self *Physics) ClearI(args ...interface{}) {
    self.Call("clear", args)
}

// Resets the active physics system. Called automatically on a Phaser.State swap.
func (self *Physics) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Destroys all active physics systems. Usually only called on a Game Shutdown, not on a State swap.
func (self *Physics) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
