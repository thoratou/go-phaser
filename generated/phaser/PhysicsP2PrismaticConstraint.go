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


// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Call("Phaser.Physics.P2.PrismaticConstraint", world, bodyA, bodyB)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint1O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, lockRotation bool) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Call("Phaser.Physics.P2.PrismaticConstraint", world, bodyA, bodyB, lockRotation)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint2O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, lockRotation bool, anchorA []interface{}) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Call("Phaser.Physics.P2.PrismaticConstraint", world, bodyA, bodyB, lockRotation, anchorA)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint3O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, lockRotation bool, anchorA []interface{}, anchorB []interface{}) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Call("Phaser.Physics.P2.PrismaticConstraint", world, bodyA, bodyB, lockRotation, anchorA, anchorB)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint4O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, lockRotation bool, anchorA []interface{}, anchorB []interface{}, axis []interface{}) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Call("Phaser.Physics.P2.PrismaticConstraint", world, bodyA, bodyB, lockRotation, anchorA, anchorB, axis)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraint5O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, lockRotation bool, anchorA []interface{}, anchorB []interface{}, axis []interface{}, maxForce int) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Call("Phaser.Physics.P2.PrismaticConstraint", world, bodyA, bodyB, lockRotation, anchorA, anchorB, axis, maxForce)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2PrismaticConstraintI(args ...interface{}) *PhysicsP2PrismaticConstraint {
    return &PhysicsP2PrismaticConstraint{js.Global.Call("Phaser.Physics.P2.PrismaticConstraint", args)}
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


