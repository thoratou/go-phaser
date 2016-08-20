// Automatic generation for Phaser.Component.InCamera
// generated file ComponentInCamera.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The InCamera component checks if the Game Object intersects with the Game Camera.
type ComponentInCamera struct {
    *js.Object
}


// Checks if this Game Objects bounds intersects with the Game Cameras bounds.
// 
// It will be `true` if they intersect, or `false` if the Game Object is fully outside of the Cameras bounds.
// 
// An object outside the bounds can be considered for camera culling if it has the AutoCull component.
func (self *ComponentInCamera) GetInCamera() bool{
    return self.Get("inCamera").Bool()
}

// Checks if this Game Objects bounds intersects with the Game Cameras bounds.
// 
// It will be `true` if they intersect, or `false` if the Game Object is fully outside of the Cameras bounds.
// 
// An object outside the bounds can be considered for camera culling if it has the AutoCull component.
func (self *ComponentInCamera) SetInCamera(member bool) {
    self.Set("inCamera", member)
}


