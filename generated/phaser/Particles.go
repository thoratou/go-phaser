// Package phaser Automatic generation for Phaser.Particles
// generated file Particles.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Particles Phaser.Particles is the Particle Manager for the game. It is called during the game update loop and in turn updates any Emitters attached to it.
type Particles struct {
    *js.Object
}

// NewParticles Phaser.Particles is the Particle Manager for the game. It is called during the game update loop and in turn updates any Emitters attached to it.
func NewParticles(game *Game) *Particles {
    return &Particles{js.Global.Get("Phaser").Get("Particles").New(game)}
}
// NewParticlesI Phaser.Particles is the Particle Manager for the game. It is called during the game update loop and in turn updates any Emitters attached to it.
func NewParticlesI(args ...interface{}) *Particles {
    return &Particles{js.Global.Get("Phaser").Get("Particles").New(args)}
}



// Particles Binding conversion method to Particles point 
func ToParticles(jsStruct interface{}) *Particles {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Particles{Object: object}
	}
	return nil
}



// Game A reference to the currently running Game.
func (self *Particles) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Particles) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Emitters Internal emitters store.
func (self *Particles) Emitters() interface{}{
    return self.Object.Get("emitters")
}

// SetEmittersA Internal emitters store.
func (self *Particles) SetEmittersA(member interface{}) {
    self.Object.Set("emitters", member)
}

// ID -
func (self *Particles) ID() int{
    return self.Object.Get("ID").Int()
}

// SetIDA -
func (self *Particles) SetIDA(member int) {
    self.Object.Set("ID", member)
}


// Add Adds a new Particle Emitter to the Particle Manager.
func (self *Particles) Add(emitter *Emitter) *Emitter{
    return &Emitter{self.Object.Call("add", emitter)}
}

// AddI Adds a new Particle Emitter to the Particle Manager.
func (self *Particles) AddI(args ...interface{}) *Emitter{
    return &Emitter{self.Object.Call("add", args)}
}

// Remove Removes an existing Particle Emitter from the Particle Manager.
func (self *Particles) Remove(emitter *Emitter) {
    self.Object.Call("remove", emitter)
}

// RemoveI Removes an existing Particle Emitter from the Particle Manager.
func (self *Particles) RemoveI(args ...interface{}) {
    self.Object.Call("remove", args)
}

// Update Called by the core game loop. Updates all Emitters who have their exists value set to true.
func (self *Particles) Update() {
    self.Object.Call("update")
}

// UpdateI Called by the core game loop. Updates all Emitters who have their exists value set to true.
func (self *Particles) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

