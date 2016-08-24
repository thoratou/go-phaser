// Package phaser Automatic generation for Phaser.Physics.P2.PrismaticConstraint
// generated file PhysicsP2PrismaticConstraint.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsP2PrismaticConstraint Connects two bodies at given offset points, letting them rotate relative to each other around this point.
type PhysicsP2PrismaticConstraint struct {
    *js.Object
}

// NewPhysicsP2PrismaticConstraint Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("PrismaticConstraint").New(world, bodyA, bodyB)}
}
// NewPhysicsP2PrismaticConstraint1O Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint1O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, lockRotation bool) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("PrismaticConstraint").New(world, bodyA, bodyB, lockRotation)}
}
// NewPhysicsP2PrismaticConstraint2O Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint2O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, lockRotation bool, anchorA []interface{}) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("PrismaticConstraint").New(world, bodyA, bodyB, lockRotation, anchorA)}
}
// NewPhysicsP2PrismaticConstraint3O Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint3O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, lockRotation bool, anchorA []interface{}, anchorB []interface{}) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("PrismaticConstraint").New(world, bodyA, bodyB, lockRotation, anchorA, anchorB)}
}
// NewPhysicsP2PrismaticConstraint4O Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint4O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, lockRotation bool, anchorA []interface{}, anchorB []interface{}, axis []interface{}) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("PrismaticConstraint").New(world, bodyA, bodyB, lockRotation, anchorA, anchorB, axis)}
}
// NewPhysicsP2PrismaticConstraint5O Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint5O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, lockRotation bool, anchorA []interface{}, anchorB []interface{}, axis []interface{}, maxForce int) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("PrismaticConstraint").New(world, bodyA, bodyB, lockRotation, anchorA, anchorB, axis, maxForce)}
}
// NewPhysicsP2PrismaticConstraintI Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraintI(args ...interface{}) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("PrismaticConstraint").New(args)}
}



// Game Local reference to game.
func (self *PhysicsP2PrismaticConstraint) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *PhysicsP2PrismaticConstraint) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// World Local reference to P2 World.
func (self *PhysicsP2PrismaticConstraint) World() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// SetWorldA Local reference to P2 World.
func (self *PhysicsP2PrismaticConstraint) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}


