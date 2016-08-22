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
func (self *PhysicsP2PrismaticConstraint) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2PrismaticConstraint) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2PrismaticConstraint) GetWorldA() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2PrismaticConstraint) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}


