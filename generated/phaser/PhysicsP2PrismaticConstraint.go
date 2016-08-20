// Automatic generation for Phaser.Physics.P2.PrismaticConstraint
// generated file PhysicsP2PrismaticConstraint.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
type PhysicsP2PrismaticConstraint struct {
    *js.Object
}


// Local reference to game.
func (self *PhysicsP2PrismaticConstraint) GetGame() Game{
    return Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2PrismaticConstraint) SetGame(member Game) {
    self.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2PrismaticConstraint) GetWorld() PhysicsP2{
    return PhysicsP2{self.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2PrismaticConstraint) SetWorld(member PhysicsP2) {
    self.Set("world", member)
}


