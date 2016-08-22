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


// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2GearConstraint(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body) *PhysicsP2GearConstraint {
    return &PhysicsP2GearConstraint{js.Global.Call("Phaser.Physics.P2.GearConstraint", world, bodyA, bodyB)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2GearConstraint1O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, angle int) *PhysicsP2GearConstraint {
    return &PhysicsP2GearConstraint{js.Global.Call("Phaser.Physics.P2.GearConstraint", world, bodyA, bodyB, angle)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2GearConstraint2O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, angle int, ratio int) *PhysicsP2GearConstraint {
    return &PhysicsP2GearConstraint{js.Global.Call("Phaser.Physics.P2.GearConstraint", world, bodyA, bodyB, angle, ratio)}
}

// Connects two bodies at given offset points, letting them rotate relative to each other around this point.
func NewPhysicsP2GearConstraintI(args ...interface{}) *PhysicsP2GearConstraint {
    return &PhysicsP2GearConstraint{js.Global.Call("Phaser.Physics.P2.GearConstraint", args)}
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


