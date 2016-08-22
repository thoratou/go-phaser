// Automatic generation for Phaser.Particles
// generated file Particles.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Phaser.Particles is the Particle Manager for the game. It is called during the game update loop and in turn updates any Emitters attached to it.
type Particles struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *Particles) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Particles) SetGame(member *Game) {
    self.Set("game", member)
}

// Internal emitters store.
func (self *Particles) GetEmitters() interface{}{
    return self.Get("emitters")
}

// Internal emitters store.
func (self *Particles) SetEmitters(member interface{}) {
    self.Set("emitters", member)
}

// -
func (self *Particles) GetID() int{
    return self.Get("ID").Int()
}

// -
func (self *Particles) SetID(member int) {
    self.Set("ID", member)
}



// Adds a new Particle Emitter to the Particle Manager.
func (self *Particles) AddI(args ...interface{}) *Emitter{
    return &Emitter{self.Call("add", args)}
}

// Removes an existing Particle Emitter from the Particle Manager.
func (self *Particles) RemoveI(args ...interface{}) {
    self.Call("remove", args)
}

// Called by the core game loop. Updates all Emitters who have their exists value set to true.
func (self *Particles) UpdateI(args ...interface{}) {
    self.Call("update", args)
}
