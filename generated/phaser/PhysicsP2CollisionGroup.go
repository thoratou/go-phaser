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
func (self *PhysicsP2CollisionGroup) GetMask() int{
    return self.Get("mask").Int()
}

// The CollisionGroup bitmask.
func (self *PhysicsP2CollisionGroup) SetMask(member int) {
    self.Set("mask", member)
}


