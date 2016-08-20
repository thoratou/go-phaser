// Automatic generation for Phaser.Physics.P2.Spring
// generated file PhysicsP2Spring.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
type PhysicsP2Spring struct {
    *js.Object
}


// Local reference to game.
func (self *PhysicsP2Spring) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2Spring) SetGame(member *Game) {
    self.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2Spring) GetWorld() *PhysicsP2{
    return &PhysicsP2{self.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2Spring) SetWorld(member *PhysicsP2) {
    self.Set("world", member)
}

// The actual p2 spring object.
func (self *PhysicsP2Spring) GetData() *P2LinearSpring{
    return &P2LinearSpring{self.Get("data")}
}

// The actual p2 spring object.
func (self *PhysicsP2Spring) SetData(member *P2LinearSpring) {
    self.Set("data", member)
}


