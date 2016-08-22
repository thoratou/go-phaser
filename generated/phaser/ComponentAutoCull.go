// Automatic generation for Phaser.Component.AutoCull
// generated file ComponentAutoCull.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The AutoCull Component is responsible for providing methods that check if a Game Object is within the bounds of the World Camera.
// It is used by the InWorld component.
type ComponentAutoCull struct {
    *js.Object
}


// The AutoCull Component is responsible for providing methods that check if a Game Object is within the bounds of the World Camera.
// It is used by the InWorld component.
func NewComponentAutoCull() *ComponentAutoCull {
    return &ComponentAutoCull{js.Global.Call("Phaser.Component.AutoCull")}
}

// The AutoCull Component is responsible for providing methods that check if a Game Object is within the bounds of the World Camera.
// It is used by the InWorld component.
func NewComponentAutoCullI(args ...interface{}) *ComponentAutoCull {
    return &ComponentAutoCull{js.Global.Call("Phaser.Component.AutoCull", args)}
}



// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *ComponentAutoCull) GetAutoCullA() bool{
    return self.Object.Get("autoCull").Bool()
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *ComponentAutoCull) SetAutoCullA(member bool) {
    self.Object.Set("autoCull", member)
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *ComponentAutoCull) GetInCameraA() bool{
    return self.Object.Get("inCamera").Bool()
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *ComponentAutoCull) SetInCameraA(member bool) {
    self.Object.Set("inCamera", member)
}


