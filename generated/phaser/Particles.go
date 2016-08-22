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


// Phaser.Particles is the Particle Manager for the game. It is called during the game update loop and in turn updates any Emitters attached to it.
func NewParticles(game *Game) *Particles {
    return &Particles{js.Global.Call("Phaser.Particles", game)}
}

// Phaser.Particles is the Particle Manager for the game. It is called during the game update loop and in turn updates any Emitters attached to it.
func NewParticlesI(args ...interface{}) *Particles {
    return &Particles{js.Global.Call("Phaser.Particles", args)}
}



// A reference to the currently running Game.
func (self *Particles) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *Particles) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Internal emitters store.
func (self *Particles) GetEmittersA() interface{}{
    return self.Object.Get("emitters")
}

// Internal emitters store.
func (self *Particles) SetEmittersA(member interface{}) {
    self.Object.Set("emitters", member)
}

// -
func (self *Particles) GetIDA() int{
    return self.Object.Get("ID").Int()
}

// -
func (self *Particles) SetIDA(member int) {
    self.Object.Set("ID", member)
}



// Adds a new Particle Emitter to the Particle Manager.
func (self *Particles) Add(emitter *Emitter) *Emitter{
    return &Emitter{self.Object.Call("add", emitter)}
}

// Adds a new Particle Emitter to the Particle Manager.
func (self *Particles) AddI(args ...interface{}) *Emitter{
    return &Emitter{self.Object.Call("add", args)}
}

// Removes an existing Particle Emitter from the Particle Manager.
func (self *Particles) Remove(emitter *Emitter) {
    self.Object.Call("remove", emitter)
}

// Removes an existing Particle Emitter from the Particle Manager.
func (self *Particles) RemoveI(args ...interface{}) {
    self.Object.Call("remove", args)
}

// Called by the core game loop. Updates all Emitters who have their exists value set to true.
func (self *Particles) Update() {
    self.Object.Call("update")
}

// Called by the core game loop. Updates all Emitters who have their exists value set to true.
func (self *Particles) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}
