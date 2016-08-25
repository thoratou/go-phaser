// Package phaser Automatic generation for Phaser.Physics.P2.CollisionGroup
// generated file PhysicsP2CollisionGroup.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsP2CollisionGroup Collision Group
type PhysicsP2CollisionGroup struct {
    *js.Object
}

// NewPhysicsP2CollisionGroup Collision Group
func NewPhysicsP2CollisionGroup(bitmask int) *PhysicsP2CollisionGroup {
    return &PhysicsP2CollisionGroup{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("CollisionGroup").New(bitmask)}
}
// NewPhysicsP2CollisionGroupI Collision Group
func NewPhysicsP2CollisionGroupI(args ...interface{}) *PhysicsP2CollisionGroup {
    return &PhysicsP2CollisionGroup{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("CollisionGroup").New(args)}
}



// PhysicsP2CollisionGroup Binding conversion method to PhysicsP2CollisionGroup point 
func ToPhysicsP2CollisionGroup(jsStruct interface{}) *PhysicsP2CollisionGroup {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsP2CollisionGroup{Object: object}
	}
	return nil
}



// Mask The CollisionGroup bitmask.
func (self *PhysicsP2CollisionGroup) Mask() int{
    return self.Object.Get("mask").Int()
}

// SetMaskA The CollisionGroup bitmask.
func (self *PhysicsP2CollisionGroup) SetMaskA(member int) {
    self.Object.Set("mask", member)
}


