// Package phaser Automatic generation for Phaser.Physics.P2.Material
// generated file PhysicsP2Material.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsP2Material A P2 Material.
// 
// \o/ ~ "Because I'm a Material girl"
type PhysicsP2Material struct {
    *js.Object
}

// NewPhysicsP2Material A P2 Material.
// 
// \o/ ~ "Because I'm a Material girl"
func NewPhysicsP2Material(name string) *PhysicsP2Material {
    return &PhysicsP2Material{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Material").New(name)}
}
// NewPhysicsP2MaterialI A P2 Material.
// 
// \o/ ~ "Because I'm a Material girl"
func NewPhysicsP2MaterialI(args ...interface{}) *PhysicsP2Material {
    return &PhysicsP2Material{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("Material").New(args)}
}



// PhysicsP2Material Binding conversion method to PhysicsP2Material point 
func ToPhysicsP2Material(jsStruct interface{}) *PhysicsP2Material {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsP2Material{Object: object}
	}
	return nil
}



// Name The user defined name given to this Material.
func (self *PhysicsP2Material) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA The user defined name given to this Material.
func (self *PhysicsP2Material) SetNameA(member string) {
    self.Object.Set("name", member)
}


