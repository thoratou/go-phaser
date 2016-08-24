// Package phaser Automatic generation for Phaser.Particles.Arcade
// generated file ParticlesArcade.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ParticlesArcade Arcade Particles is a Particle System integrated with Arcade Physics.
type ParticlesArcade struct {
    *js.Object
}

// NewParticlesArcade Arcade Particles is a Particle System integrated with Arcade Physics.
func NewParticlesArcade() *ParticlesArcade {
    return &ParticlesArcade{js.Global.Get("Phaser").Get("Particles").Get("Arcade").New()}
}
// NewParticlesArcadeI Arcade Particles is a Particle System integrated with Arcade Physics.
func NewParticlesArcadeI(args ...interface{}) *ParticlesArcade {
    return &ParticlesArcade{js.Global.Get("Phaser").Get("Particles").Get("Arcade").New(args)}
}




