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


// A constraint that tries to keep the distance between two bodies constant.
func NewPhysicsP2DistanceConstraint(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, distance int) *PhysicsP2DistanceConstraint {
    return &PhysicsP2DistanceConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("DistanceConstraint").New(world, bodyA, bodyB, distance)}
}

// A constraint that tries to keep the distance between two bodies constant.
func NewPhysicsP2DistanceConstraint1O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, distance int, localAnchorA []interface{}) *PhysicsP2DistanceConstraint {
    return &PhysicsP2DistanceConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("DistanceConstraint").New(world, bodyA, bodyB, distance, localAnchorA)}
}

// A constraint that tries to keep the distance between two bodies constant.
func NewPhysicsP2DistanceConstraint2O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, distance int, localAnchorA []interface{}, localAnchorB []interface{}) *PhysicsP2DistanceConstraint {
    return &PhysicsP2DistanceConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("DistanceConstraint").New(world, bodyA, bodyB, distance, localAnchorA, localAnchorB)}
}

// A constraint that tries to keep the distance between two bodies constant.
func NewPhysicsP2DistanceConstraint3O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, distance int, localAnchorA []interface{}, localAnchorB []interface{}, maxForce interface{}) *PhysicsP2DistanceConstraint {
    return &PhysicsP2DistanceConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("DistanceConstraint").New(world, bodyA, bodyB, distance, localAnchorA, localAnchorB, maxForce)}
}

// A constraint that tries to keep the distance between two bodies constant.
func NewPhysicsP2DistanceConstraintI(args ...interface{}) *PhysicsP2DistanceConstraint {
    return &PhysicsP2DistanceConstraint{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("DistanceConstraint").New(args)}
}



// Local reference to game.
func (self *PhysicsP2DistanceConstraint) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsP2DistanceConstraint) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to P2 World.
func (self *PhysicsP2DistanceConstraint) World() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// Local reference to P2 World.
func (self *PhysicsP2DistanceConstraint) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}


