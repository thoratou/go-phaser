// Automatic generation for Phaser.Particles.Arcade
// generated file ParticlesArcade.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Arcade Particles is a Particle System integrated with Arcade Physics.
type ParticlesArcade struct {
    *js.Object
}


// Arcade Particles is a Particle System integrated with Arcade Physics.
func NewParticlesArcade() *ParticlesArcade {
    return &ParticlesArcade{js.Global.Get("Phaser").Get("Particles").Get("Arcade").New()}
}

// Arcade Particles is a Particle System integrated with Arcade Physics.
func NewParticlesArcadeI(args ...interface{}) *ParticlesArcade {
    return &ParticlesArcade{js.Global.Get("Phaser").Get("Particles").Get("Arcade").New(args)}
}




