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
func (self *PhysicsP2RotationalSpring) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2RotationalSpring) SetGame(member *Game) {
    self.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2RotationalSpring) GetWorld() *PhysicsP2{
    return &PhysicsP2{self.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2RotationalSpring) SetWorld(member *PhysicsP2) {
    self.Set("world", member)
}

// The actual p2 spring object.
func (self *PhysicsP2RotationalSpring) GetData() *P2RotationalSpring{
    return &P2RotationalSpring{self.Get("data")}
}

// The actual p2 spring object.
func (self *PhysicsP2RotationalSpring) SetData(member *P2RotationalSpring) {
    self.Set("data", member)
}


