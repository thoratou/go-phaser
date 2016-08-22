// Automatic generation for Phaser.Component.LifeSpan
// generated file ComponentLifeSpan.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// LifeSpan Component Features.
type ComponentLifeSpan struct {
    *js.Object
}


// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *ComponentLifeSpan) GetAliveA() bool{
    return self.Object.Get("alive").Bool()
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *ComponentLifeSpan) SetAliveA(member bool) {
    self.Object.Set("alive", member)
}

// The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *ComponentLifeSpan) GetLifespanA() int{
    return self.Object.Get("lifespan").Int()
}

// The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *ComponentLifeSpan) SetLifespanA(member int) {
    self.Object.Set("lifespan", member)
}



// The LifeSpan component preUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentLifeSpan) PreUpdate() {
    self.Object.Call("preUpdate")
}

// The LifeSpan component preUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentLifeSpan) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *ComponentLifeSpan) Revive() *DisplayObject{
    return &DisplayObject{self.Object.Call("revive")}
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *ComponentLifeSpan) Revive1O(health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", health)}
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *ComponentLifeSpan) ReviveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", args)}
}

// Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *ComponentLifeSpan) Kill() *DisplayObject{
    return &DisplayObject{self.Object.Call("kill")}
}

// Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *ComponentLifeSpan) KillI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("kill", args)}
}
