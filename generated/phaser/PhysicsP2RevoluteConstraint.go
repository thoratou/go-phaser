// Automatic generation for Phaser.Physics.P2.RevoluteConstraint
// generated file PhysicsP2RevoluteConstraint.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
// The pivot points are given in world (pixel) coordinates.
type PhysicsP2RevoluteConstraint struct {
    *js.Object
}


// Local reference to game.
func (self *PhysicsP2RevoluteConstraint) GetGame() Game{
    return Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2RevoluteConstraint) SetGame(member Game) {
    self.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2RevoluteConstraint) GetWorld() PhysicsP2{
    return PhysicsP2{self.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2RevoluteConstraint) SetWorld(member PhysicsP2) {
    self.Set("world", member)
}


