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
func (self *PhysicsP2Spring) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2Spring) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2Spring) GetWorldA() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2Spring) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}

// The actual p2 spring object.
func (self *PhysicsP2Spring) GetDataA() *P2LinearSpring{
    return &P2LinearSpring{self.Object.Get("data")}
}

// The actual p2 spring object.
func (self *PhysicsP2Spring) SetDataA(member *P2LinearSpring) {
    self.Object.Set("data", member)
}


