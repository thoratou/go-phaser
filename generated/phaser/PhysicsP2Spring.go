// Package phaser Automatic generation for Phaser.Physics.P2.Spring
// generated file PhysicsP2Spring.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsP2Spring Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
type PhysicsP2Spring struct {
    *js.Object
}

// NewPhysicsP2Spring Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2Spring(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body) *PhysicsP2Spring {
    return &PhysicsP2Spring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Spring").New(world, bodyA, bodyB)}
}
// NewPhysicsP2Spring1O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2Spring1O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, restLength int) *PhysicsP2Spring {
    return &PhysicsP2Spring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Spring").New(world, bodyA, bodyB, restLength)}
}
// NewPhysicsP2Spring2O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2Spring2O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, restLength int, stiffness int) *PhysicsP2Spring {
    return &PhysicsP2Spring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Spring").New(world, bodyA, bodyB, restLength, stiffness)}
}
// NewPhysicsP2Spring3O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2Spring3O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, restLength int, stiffness int, damping int) *PhysicsP2Spring {
    return &PhysicsP2Spring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Spring").New(world, bodyA, bodyB, restLength, stiffness, damping)}
}
// NewPhysicsP2Spring4O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2Spring4O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, restLength int, stiffness int, damping int, worldA []interface{}) *PhysicsP2Spring {
    return &PhysicsP2Spring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Spring").New(world, bodyA, bodyB, restLength, stiffness, damping, worldA)}
}
// NewPhysicsP2Spring5O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2Spring5O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, restLength int, stiffness int, damping int, worldA []interface{}, worldB []interface{}) *PhysicsP2Spring {
    return &PhysicsP2Spring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Spring").New(world, bodyA, bodyB, restLength, stiffness, damping, worldA, worldB)}
}
// NewPhysicsP2Spring6O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2Spring6O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, restLength int, stiffness int, damping int, worldA []interface{}, worldB []interface{}, localA []interface{}) *PhysicsP2Spring {
    return &PhysicsP2Spring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Spring").New(world, bodyA, bodyB, restLength, stiffness, damping, worldA, worldB, localA)}
}
// NewPhysicsP2Spring7O Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2Spring7O(world *PhysicsP2, bodyA *P2Body, bodyB *P2Body, restLength int, stiffness int, damping int, worldA []interface{}, worldB []interface{}, localA []interface{}, localB []interface{}) *PhysicsP2Spring {
    return &PhysicsP2Spring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Spring").New(world, bodyA, bodyB, restLength, stiffness, damping, worldA, worldB, localA, localB)}
}
// NewPhysicsP2SpringI Creates a linear spring, connecting two bodies. A spring can have a resting length, a stiffness and damping.
func NewPhysicsP2SpringI(args ...interface{}) *PhysicsP2Spring {
    return &PhysicsP2Spring{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Spring").New(args)}
}



// Game Local reference to game.
func (self *PhysicsP2Spring) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *PhysicsP2Spring) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// World Local reference to P2 World.
func (self *PhysicsP2Spring) World() *PhysicsP2{
    return &PhysicsP2{self.Object.Get("world")}
}

// SetWorldA Local reference to P2 World.
func (self *PhysicsP2Spring) SetWorldA(member *PhysicsP2) {
    self.Object.Set("world", member)
}

// Data The actual p2 spring object.
func (self *PhysicsP2Spring) Data() *P2LinearSpring{
    return &P2LinearSpring{self.Object.Get("data")}
}

// SetDataA The actual p2 spring object.
func (self *PhysicsP2Spring) SetDataA(member *P2LinearSpring) {
    self.Object.Set("data", member)
}


