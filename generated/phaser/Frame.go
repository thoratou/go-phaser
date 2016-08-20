// Automatic generation for Phaser.Frame
// generated file Frame.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A Frame is a single frame of an animation and is part of a FrameData collection.
type Frame struct {
    *js.Object
}


// The index of this Frame within the FrameData set it is being added to.
func (self *Frame) GetIndex() float64{
    return self.Get("index").Float()
}

// The index of this Frame within the FrameData set it is being added to.
func (self *Frame) SetIndex(member float64) {
    self.Set("index", member)
}

// X position within the image to cut from.
func (self *Frame) GetX() float64{
    return self.Get("x").Float()
}

// X position within the image to cut from.
func (self *Frame) SetX(member float64) {
    self.Set("x", member)
}

// Y position within the image to cut from.
func (self *Frame) GetY() float64{
    return self.Get("y").Float()
}

// Y position within the image to cut from.
func (self *Frame) SetY(member float64) {
    self.Set("y", member)
}

// Width of the frame.
func (self *Frame) GetWidth() float64{
    return self.Get("width").Float()
}

// Width of the frame.
func (self *Frame) SetWidth(member float64) {
    self.Set("width", member)
}

// Height of the frame.
func (self *Frame) GetHeight() float64{
    return self.Get("height").Float()
}

// Height of the frame.
func (self *Frame) SetHeight(member float64) {
    self.Set("height", member)
}

// Useful for Texture Atlas files (is set to the filename value).
func (self *Frame) GetName() string{
    return self.Get("name").String()
}

// Useful for Texture Atlas files (is set to the filename value).
func (self *Frame) SetName(member string) {
    self.Set("name", member)
}

// Center X position within the image to cut from.
func (self *Frame) GetCenterX() float64{
    return self.Get("centerX").Float()
}

// Center X position within the image to cut from.
func (self *Frame) SetCenterX(member float64) {
    self.Set("centerX", member)
}

// Center Y position within the image to cut from.
func (self *Frame) GetCenterY() float64{
    return self.Get("centerY").Float()
}

// Center Y position within the image to cut from.
func (self *Frame) SetCenterY(member float64) {
    self.Set("centerY", member)
}

// The distance from the top left to the bottom-right of this Frame.
func (self *Frame) GetDistance() float64{
    return self.Get("distance").Float()
}

// The distance from the top left to the bottom-right of this Frame.
func (self *Frame) SetDistance(member float64) {
    self.Set("distance", member)
}

// Rotated? (not yet implemented)
func (self *Frame) GetRotated() bool{
    return self.Get("rotated").Bool()
}

// Rotated? (not yet implemented)
func (self *Frame) SetRotated(member bool) {
    self.Set("rotated", member)
}

// Either 'cw' or 'ccw', rotation is always 90 degrees.
func (self *Frame) GetRotationDirection() string{
    return self.Get("rotationDirection").String()
}

// Either 'cw' or 'ccw', rotation is always 90 degrees.
func (self *Frame) SetRotationDirection(member string) {
    self.Set("rotationDirection", member)
}

// Was it trimmed when packed?
func (self *Frame) GetTrimmed() bool{
    return self.Get("trimmed").Bool()
}

// Was it trimmed when packed?
func (self *Frame) SetTrimmed(member bool) {
    self.Set("trimmed", member)
}

// Width of the original sprite before it was trimmed.
func (self *Frame) GetSourceSizeW() float64{
    return self.Get("sourceSizeW").Float()
}

// Width of the original sprite before it was trimmed.
func (self *Frame) SetSourceSizeW(member float64) {
    self.Set("sourceSizeW", member)
}

// Height of the original sprite before it was trimmed.
func (self *Frame) GetSourceSizeH() float64{
    return self.Get("sourceSizeH").Float()
}

// Height of the original sprite before it was trimmed.
func (self *Frame) SetSourceSizeH(member float64) {
    self.Set("sourceSizeH", member)
}

// X position of the trimmed sprite inside original sprite.
func (self *Frame) GetSpriteSourceSizeX() float64{
    return self.Get("spriteSourceSizeX").Float()
}

// X position of the trimmed sprite inside original sprite.
func (self *Frame) SetSpriteSourceSizeX(member float64) {
    self.Set("spriteSourceSizeX", member)
}

// Y position of the trimmed sprite inside original sprite.
func (self *Frame) GetSpriteSourceSizeY() float64{
    return self.Get("spriteSourceSizeY").Float()
}

// Y position of the trimmed sprite inside original sprite.
func (self *Frame) SetSpriteSourceSizeY(member float64) {
    self.Set("spriteSourceSizeY", member)
}

// Width of the trimmed sprite.
func (self *Frame) GetSpriteSourceSizeW() float64{
    return self.Get("spriteSourceSizeW").Float()
}

// Width of the trimmed sprite.
func (self *Frame) SetSpriteSourceSizeW(member float64) {
    self.Set("spriteSourceSizeW", member)
}

// Height of the trimmed sprite.
func (self *Frame) GetSpriteSourceSizeH() float64{
    return self.Get("spriteSourceSizeH").Float()
}

// Height of the trimmed sprite.
func (self *Frame) SetSpriteSourceSizeH(member float64) {
    self.Set("spriteSourceSizeH", member)
}

// The right of the Frame (x + width).
func (self *Frame) GetRight() float64{
    return self.Get("right").Float()
}

// The right of the Frame (x + width).
func (self *Frame) SetRight(member float64) {
    self.Set("right", member)
}

// The bottom of the frame (y + height).
func (self *Frame) GetBottom() float64{
    return self.Get("bottom").Float()
}

// The bottom of the frame (y + height).
func (self *Frame) SetBottom(member float64) {
    self.Set("bottom", member)
}



// Adjusts of all the Frame properties based on the given width and height values.
func (self *Frame) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// If the frame was trimmed when added to the Texture Atlas this records the trim and source data.
func (self *Frame) SetTrimI(args ...interface{}) {
    self.Call("setTrim", args)
}

// Clones this Frame into a new Phaser.Frame object and returns it.
// Note that all properties are cloned, including the name, index and UUID.
func (self *Frame) CloneI(args ...interface{}) Frame{
    return Frame{self.Call("clone", args)}
}

// Returns a Rectangle set to the dimensions of this Frame.
func (self *Frame) GetRectI(args ...interface{}) Rectangle{
    return Rectangle{self.Call("getRect", args)}
}
