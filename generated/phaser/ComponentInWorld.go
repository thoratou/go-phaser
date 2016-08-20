// Automatic generation for Phaser.Component.InWorld
// generated file ComponentInWorld.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The InWorld component checks if a Game Object is within the Game World Bounds.
// An object is considered as being "in bounds" so long as its own bounds intersects at any point with the World bounds.
// If the AutoCull component is enabled on the Game Object then it will check the Game Object against the Camera bounds as well.
type ComponentInWorld struct {
    *js.Object
}


// If this is set to `true` the Game Object checks if it is within the World bounds each frame. 
// 
// When it is no longer intersecting the world bounds it dispatches the `onOutOfBounds` event.
// 
// If it was *previously* out of bounds but is now intersecting the world bounds again it dispatches the `onEnterBounds` event.
// 
// It also optionally kills the Game Object if `outOfBoundsKill` is `true`.
// 
// When `checkWorldBounds` is enabled it forces the Game Object to calculate its full bounds every frame.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *ComponentInWorld) GetCheckWorldBounds() bool{
    return self.Get("checkWorldBounds").Bool()
}

// If this is set to `true` the Game Object checks if it is within the World bounds each frame. 
// 
// When it is no longer intersecting the world bounds it dispatches the `onOutOfBounds` event.
// 
// If it was *previously* out of bounds but is now intersecting the world bounds again it dispatches the `onEnterBounds` event.
// 
// It also optionally kills the Game Object if `outOfBoundsKill` is `true`.
// 
// When `checkWorldBounds` is enabled it forces the Game Object to calculate its full bounds every frame.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *ComponentInWorld) SetCheckWorldBounds(member bool) {
    self.Set("checkWorldBounds", member)
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *ComponentInWorld) GetOutOfBoundsKill() bool{
    return self.Get("outOfBoundsKill").Bool()
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *ComponentInWorld) SetOutOfBoundsKill(member bool) {
    self.Set("outOfBoundsKill", member)
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *ComponentInWorld) GetOutOfCameraBoundsKill() bool{
    return self.Get("outOfCameraBoundsKill").Bool()
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *ComponentInWorld) SetOutOfCameraBoundsKill(member bool) {
    self.Set("outOfCameraBoundsKill", member)
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *ComponentInWorld) GetInWorld() bool{
    return self.Get("inWorld").Bool()
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *ComponentInWorld) SetInWorld(member bool) {
    self.Set("inWorld", member)
}



// The InWorld component preUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentInWorld) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}
