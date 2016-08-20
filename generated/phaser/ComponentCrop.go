// Automatic generation for Phaser.Component.Crop
// generated file ComponentCrop.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Crop component provides the ability to crop a texture based Game Object to a defined rectangle, 
// which can be updated in real-time.
type ComponentCrop struct {
    *js.Object
}


// The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *ComponentCrop) GetCropRect() *Rectangle{
    return &Rectangle{self.Get("cropRect")}
}

// The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *ComponentCrop) SetCropRect(member *Rectangle) {
    self.Set("cropRect", member)
}



// Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
// 
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
// 
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object 
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
// 
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`, 
// in which case the values are duplicated to a local object.
func (self *ComponentCrop) CropI(args ...interface{}) {
    self.Call("crop", args)
}

// If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *ComponentCrop) UpdateCropI(args ...interface{}) {
    self.Call("updateCrop", args)
}
