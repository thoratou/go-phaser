// Automatic generation for Phaser.Physics.P2.Material
// generated file PhysicsP2Material.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A P2 Material.
// 
// \o/ ~ "Because I'm a Material girl"
type PhysicsP2Material struct {
    *js.Object
}


// The user defined name given to this Material.
func (self *PhysicsP2Material) GetNameA() string{
    return self.Object.Get("name").String()
}

// The user defined name given to this Material.
func (self *PhysicsP2Material) SetNameA(member string) {
    self.Object.Set("name", member)
}


