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
func (self *PhysicsP2LockConstraint) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2LockConstraint) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2LockConstraint) GetWorldA() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2LockConstraint) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}


