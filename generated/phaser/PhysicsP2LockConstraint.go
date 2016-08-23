// Automatic generation for Phaser.Physics.P2.LockConstraint
// generated file PhysicsP2LockConstraint.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Locks the relative position between two bodies.
type PhysicsP2LockConstraint struct {
    *js.Object
}


// Locks the relative position between two bodies.
func NewPhysicsP2LockConstraint(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body) *PhysicsP2LockConstraint {
    return &PhysicsP2LockConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("LockConstraint").New(world, bodyA, bodyB)}
}

// Locks the relative position between two bodies.
func NewPhysicsP2LockConstraint1O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, offset []interface{}) *PhysicsP2LockConstraint {
    return &PhysicsP2LockConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("LockConstraint").New(world, bodyA, bodyB, offset)}
}

// Locks the relative position between two bodies.
func NewPhysicsP2LockConstraint2O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, offset []interface{}, angle int) *PhysicsP2LockConstraint {
    return &PhysicsP2LockConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("LockConstraint").New(world, bodyA, bodyB, offset, angle)}
}

// Locks the relative position between two bodies.
func NewPhysicsP2LockConstraint3O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, offset []interface{}, angle int, maxForce int) *PhysicsP2LockConstraint {
    return &PhysicsP2LockConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("LockConstraint").New(world, bodyA, bodyB, offset, angle, maxForce)}
}

// Locks the relative position between two bodies.
func NewPhysicsP2LockConstraintI(args ...interface{}) *PhysicsP2LockConstraint {
    return &PhysicsP2LockConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("LockConstraint").New(args)}
}



// Local reference to game.
func (self *PhysicsP2LockConstraint) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2LockConstraint) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2LockConstraint) World() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2LockConstraint) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}


