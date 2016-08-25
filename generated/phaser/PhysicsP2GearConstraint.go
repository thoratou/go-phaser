// Package phaser Automatic generation for Phaser.Physics.P2.GearConstraint
// generated file PhysicsP2GearConstraint.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsP2GearConstraint Connects two bodies at given offset points, letting them rotate relative to each other around this point.
type PhysicsP2GearConstraint struct {
    *js.Object
}

// NewPhysicsP2GearConstraint Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2GearConstraint(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body) *PhysicsP2GearConstraint {
    return &PhysicsP2GearConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("GearConstraint").New(world, bodyA, bodyB)}
}
// NewPhysicsP2GearConstraint1O Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2GearConstraint1O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, angle int) *PhysicsP2GearConstraint {
    return &PhysicsP2GearConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("GearConstraint").New(world, bodyA, bodyB, angle)}
}
// NewPhysicsP2GearConstraint2O Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2GearConstraint2O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, angle int, ratio int) *PhysicsP2GearConstraint {
    return &PhysicsP2GearConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("GearConstraint").New(world, bodyA, bodyB, angle, ratio)}
}
// NewPhysicsP2GearConstraintI Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2GearConstraintI(args ...interface{}) *PhysicsP2GearConstraint {
    return &PhysicsP2GearConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("GearConstraint").New(args)}
}



// PhysicsP2GearConstraint Binding conversion method to PhysicsP2GearConstraint point 
func ToPhysicsP2GearConstraint(jsStruct interface{}) *PhysicsP2GearConstraint {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsP2GearConstraint{Object: object}
	}
	return nil
}



// Game Local reference to game.
func (self *PhysicsP2GearConstraint) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *PhysicsP2GearConstraint) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// World Local reference to P2 World.
func (self *PhysicsP2GearConstraint) World() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// SetWorldA Local reference to P2 World.
func (self *PhysicsP2GearConstraint) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}


