// Automatic generation for Phaser.Physics.Arcade.TilemapCollision
// generated file PhysicsArcadeTilemapCollision.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Arcade Physics Tile map collision methods.
type PhysicsArcadeTilemapCollision struct {
    *js.Object
}


// A value added to the delta values during collision with tiles. Adjust this if you get tunneling.
func (self *PhysicsArcadeTilemapCollision) GetTILE_BIAS() int{
    return self.Get("TILE_BIAS").Int()
}

// A value added to the delta values during collision with tiles. Adjust this if you get tunneling.
func (self *PhysicsArcadeTilemapCollision) SetTILE_BIAS(member int) {
    self.Set("TILE_BIAS", member)
}


