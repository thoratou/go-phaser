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


// The CollisionGroup bitmask.
func (self *PhysicsP2CollisionGroup) GetMask() float64{
    return self.Get("mask").Float()
}

// The CollisionGroup bitmask.
func (self *PhysicsP2CollisionGroup) SetMask(member float64) {
    self.Set("mask", member)
}


