// Automatic generation for Phaser.Physics.P2.DistanceConstraint
// generated file PhysicsP2DistanceConstraint.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A constraint that tries to keep the distance between two bodies constant.
type PhysicsP2DistanceConstraint struct {
    *js.Object
}


// Local reference to game.
func (self *PhysicsP2DistanceConstraint) GetGame() Game{
    return Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2DistanceConstraint) SetGame(member Game) {
    self.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2DistanceConstraint) GetWorld() PhysicsP2{
    return PhysicsP2{self.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2DistanceConstraint) SetWorld(member PhysicsP2) {
    self.Set("world", member)
}


