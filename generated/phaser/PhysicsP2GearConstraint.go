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
func (self *PhysicsP2GearConstraint) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2GearConstraint) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2GearConstraint) GetWorldA() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2GearConstraint) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}


