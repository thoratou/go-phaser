// Package phaser Automatic generation for Phaser.Physics.P2.ContactMaterial
// generated file PhysicsP2ContactMaterial.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsP2ContactMaterial Defines a physics material
type PhysicsP2ContactMaterial struct {
    *js.Object
}

// NewPhysicsP2ContactMaterial Defines a physics material
func NewPhysicsP2ContactMaterial(materialA *PhysicsP2Material, materialB *PhysicsP2Material) *PhysicsP2ContactMaterial {
    return &PhysicsP2ContactMaterial{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("ContactMaterial").New(materialA, materialB)}
}
// NewPhysicsP2ContactMaterial1O Defines a physics material
func NewPhysicsP2ContactMaterial1O(materialA *PhysicsP2Material, materialB *PhysicsP2Material, options interface{}) *PhysicsP2ContactMaterial {
    return &PhysicsP2ContactMaterial{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("ContactMaterial").New(materialA, materialB, options)}
}
// NewPhysicsP2ContactMaterialI Defines a physics material
func NewPhysicsP2ContactMaterialI(args ...interface{}) *PhysicsP2ContactMaterial {
    return &PhysicsP2ContactMaterial{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("ContactMaterial").New(args)}
}



// PhysicsP2ContactMaterial Binding conversion method to PhysicsP2ContactMaterial point 
func ToPhysicsP2ContactMaterial(jsStruct interface{}) *PhysicsP2ContactMaterial {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsP2ContactMaterial{Object: object}
	}
	return nil
}




