// Automatic generation for Phaser.Physics.P2.LockConstraint
// generated file PhysicsP2LockConstraint.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Locks the relative position between two bodies.
type PhysicsP2LockConstraint struct {
    *js.Object
}


// Local reference to game.
func (self *PhysicsP2LockConstraint) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2LockConstraint) SetGame(member *Game) {
    self.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2LockConstraint) GetWorld() *PhysicsP2{
    return &PhysicsP2{self.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2LockConstraint) SetWorld(member *PhysicsP2) {
    self.Set("world", member)
}


