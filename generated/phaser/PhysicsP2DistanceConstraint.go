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
func (self *PhysicsP2DistanceConstraint) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2DistanceConstraint) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2DistanceConstraint) GetWorldA() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2DistanceConstraint) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}


