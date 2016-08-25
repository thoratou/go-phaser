// Package phaser Automatic generation for Phaser.Component.AutoCull
// generated file ComponentAutoCull.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ComponentAutoCull The AutoCull Component is responsible for providing methods that check if a Game Object is within the bounds of the World Camera.
// It is used by the InWorld component.
type ComponentAutoCull struct {
    *js.Object
}

// NewComponentAutoCull The AutoCull Component is responsible for providing methods that check if a Game Object is within the bounds of the World Camera.
// It is used by the InWorld component.
func NewComponentAutoCull() *ComponentAutoCull {
    return &ComponentAutoCull{js.Global.Get("Phaser").Get("Component").Get("AutoCull").New()}
}
// NewComponentAutoCullI The AutoCull Component is responsible for providing methods that check if a Game Object is within the bounds of the World Camera.
// It is used by the InWorld component.
func NewComponentAutoCullI(args ...interface{}) *ComponentAutoCull {
    return &ComponentAutoCull{js.Global.Get("Phaser").Get("Component").Get("AutoCull").New(args)}
}



// ComponentAutoCull Binding conversion method to ComponentAutoCull point 
func ToComponentAutoCull(jsStruct interface{}) *ComponentAutoCull {
    if object, ok := jsStruct.(*js.Object); ok {
		return &ComponentAutoCull{Object: object}
	}
	return nil
}



// AutoCull A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *ComponentAutoCull) AutoCull() bool{
    return self.Object.Get("autoCull").Bool()
}

// SetAutoCullA A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *ComponentAutoCull) SetAutoCullA(member bool) {
    self.Object.Set("autoCull", member)
}

// InCamera Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *ComponentAutoCull) InCamera() bool{
    return self.Object.Get("inCamera").Bool()
}

// SetInCameraA Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *ComponentAutoCull) SetInCameraA(member bool) {
    self.Object.Set("inCamera", member)
}


