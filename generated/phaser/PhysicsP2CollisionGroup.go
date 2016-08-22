// Automatic generation for Phaser.Physics.P2.CollisionGroup
// generated file PhysicsP2CollisionGroup.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Collision Group
type PhysicsP2CollisionGroup struct {
    *js.Object
}


// Collision Group
func NewPhysicsP2CollisionGroup(bitmask int) *PhysicsP2CollisionGroup {
    return &PhysicsP2CollisionGroup{js.Global.Call("Phaser.Physics.P2.CollisionGroup", bitmask)}
}

// Collision Group
func NewPhysicsP2CollisionGroupI(args ...interface{}) *PhysicsP2CollisionGroup {
    return &PhysicsP2CollisionGroup{js.Global.Call("Phaser.Physics.P2.CollisionGroup", args)}
}



// The CollisionGroup bitmask.
func (self *PhysicsP2CollisionGroup) GetMaskA() int{
    return self.Object.Get("mask").Int()
}

// The CollisionGroup bitmask.
func (self *PhysicsP2CollisionGroup) SetMaskA(member int) {
    self.Object.Set("mask", member)
}


