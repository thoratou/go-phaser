// Automatic generation for Phaser.Physics.P2.ContactMaterial
// generated file PhysicsP2ContactMaterial.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Defines a physics material
type PhysicsP2ContactMaterial struct {
    *js.Object
}


// Defines a physics material
func NewPhysicsP2ContactMaterial(materialA *PhysicsP2Material, materialB *PhysicsP2Material) *PhysicsP2ContactMaterial {
    return &PhysicsP2ContactMaterial{js.Global.Call("Phaser.Physics.P2.ContactMaterial", materialA, materialB)}
}

// Defines a physics material
func NewPhysicsP2ContactMaterial1O(materialA *PhysicsP2Material, materialB *PhysicsP2Material, options interface{}) *PhysicsP2ContactMaterial {
    return &PhysicsP2ContactMaterial{js.Global.Call("Phaser.Physics.P2.ContactMaterial", materialA, materialB, options)}
}

// Defines a physics material
func NewPhysicsP2ContactMaterialI(args ...interface{}) *PhysicsP2ContactMaterial {
    return &PhysicsP2ContactMaterial{js.Global.Call("Phaser.Physics.P2.ContactMaterial", args)}
}




