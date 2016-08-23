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


// The InWorld component checks if a Game Object is within the Game World Bounds.
// An object is considered as being "in bounds" so long as its own bounds intersects at any point with the World bounds.
// If the AutoCull component is enabled on the Game Object then it will check the Game Object against the Camera bounds as well.
func NewComponentInWorld() *ComponentInWorld {
    return &ComponentInWorld{js.Global.Get("Phaser").Get("Component").Get("InWorld").New()}
}

// The InWorld component checks if a Game Object is within the Game World Bounds.
// An object is considered as being "in bounds" so long as its own bounds intersects at any point with the World bounds.
// If the AutoCull component is enabled on the Game Object then it will check the Game Object against the Camera bounds as well.
func NewComponentInWorldI(args ...interface{}) *ComponentInWorld {
    return &ComponentInWorld{js.Global.Get("Phaser").Get("Component").Get("InWorld").New(args)}
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
func (self *ComponentInWorld) CheckWorldBounds() bool{
    return self.Object.Get("checkWorldBounds").Bool()
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
func (self *ComponentInWorld) SetCheckWorldBoundsA(member bool) {
    self.Object.Set("checkWorldBounds", member)
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *ComponentInWorld) OutOfBoundsKill() bool{
    return self.Object.Get("outOfBoundsKill").Bool()
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *ComponentInWorld) SetOutOfBoundsKillA(member bool) {
    self.Object.Set("outOfBoundsKill", member)
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *ComponentInWorld) OutOfCameraBoundsKill() bool{
    return self.Object.Get("outOfCameraBoundsKill").Bool()
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *ComponentInWorld) SetOutOfCameraBoundsKillA(member bool) {
    self.Object.Set("outOfCameraBoundsKill", member)
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *ComponentInWorld) InWorld() bool{
    return self.Object.Get("inWorld").Bool()
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *ComponentInWorld) SetInWorldA(member bool) {
    self.Object.Set("inWorld", member)
}



// The InWorld component preUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentInWorld) PreUpdate() {
    self.Object.Call("preUpdate")
}

// The InWorld component preUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentInWorld) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}
