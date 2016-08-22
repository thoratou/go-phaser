// Automatic generation for Phaser.Physics.P2.RotationalSpring
// generated file PhysicsP2RotationalSpring.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Creates a rotational spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
type PhysicsP2RotationalSpring struct {
    *js.Object
}


// Local reference to game.
func (self *PhysicsP2RotationalSpring) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2RotationalSpring) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2RotationalSpring) GetWorldA() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2RotationalSpring) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}

// The actual p2 spring object.
func (self *PhysicsP2RotationalSpring) GetDataA() *P2RotationalSpring{
    return &P2RotationalSpring{self.Object.Get("data")}
}

// The actual p2 spring object.
func (self *PhysicsP2RotationalSpring) SetDataA(member *P2RotationalSpring) {
    self.Object.Set("data", member)
}


