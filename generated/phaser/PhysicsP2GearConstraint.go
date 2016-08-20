// Automatic generation for Phaser.Physics.P2.GearConstraint
// generated file PhysicsP2GearConstraint.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
type PhysicsP2GearConstraint struct {
    *js.Object
}


// Local reference to game.
func (self *PhysicsP2GearConstraint) GetGame() Game{
    return Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2GearConstraint) SetGame(member Game) {
    self.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2GearConstraint) GetWorld() PhysicsP2{
    return PhysicsP2{self.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2GearConstraint) SetWorld(member PhysicsP2) {
    self.Set("world", member)
}


