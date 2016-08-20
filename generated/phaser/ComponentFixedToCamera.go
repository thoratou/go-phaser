// Automatic generation for Phaser.Component.FixedToCamera
// generated file ComponentFixedToCamera.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The FixedToCamera component enables a Game Object to be rendered relative to the game camera coordinates, regardless 
// of where in the world the camera is. This is used for things like sticking game UI to the camera that scrolls as it moves around the world.
type ComponentFixedToCamera struct {
    *js.Object
}


// A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
// 
// The values are adjusted at the rendering stage, overriding the Game Objects actual world position.
// 
// The end result is that the Game Object will appear to be 'fixed' to the camera, regardless of where in the game world
// the camera is viewing. This is useful if for example this Game Object is a UI item that you wish to be visible at all times 
// regardless where in the world the camera is.
// 
// The offsets are stored in the `cameraOffset` property.
// 
// Note that the `cameraOffset` values are in addition to any parent of this Game Object on the display list.
// 
// Be careful not to set `fixedToCamera` on Game Objects which are in Groups that already have `fixedToCamera` enabled on them.
func (self *ComponentFixedToCamera) GetFixedToCamera() bool{
    return self.Get("fixedToCamera").Bool()
}

// A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
// 
// The values are adjusted at the rendering stage, overriding the Game Objects actual world position.
// 
// The end result is that the Game Object will appear to be 'fixed' to the camera, regardless of where in the game world
// the camera is viewing. This is useful if for example this Game Object is a UI item that you wish to be visible at all times 
// regardless where in the world the camera is.
// 
// The offsets are stored in the `cameraOffset` property.
// 
// Note that the `cameraOffset` values are in addition to any parent of this Game Object on the display list.
// 
// Be careful not to set `fixedToCamera` on Game Objects which are in Groups that already have `fixedToCamera` enabled on them.
func (self *ComponentFixedToCamera) SetFixedToCamera(member bool) {
    self.Set("fixedToCamera", member)
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *ComponentFixedToCamera) GetCameraOffset() *Point{
    return &Point{self.Get("cameraOffset")}
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *ComponentFixedToCamera) SetCameraOffset(member *Point) {
    self.Set("cameraOffset", member)
}



// The FixedToCamera component postUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentFixedToCamera) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}
