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
func (self *Frame) GetIndex() int{
    return self.Get("index").Int()
}

// The index of this Frame within the FrameData set it is being added to.
func (self *Frame) SetIndex(member int) {
    self.Set("index", member)
}

// X position within the image to cut from.
func (self *Frame) GetX() int{
    return self.Get("x").Int()
}

// X position within the image to cut from.
func (self *Frame) SetX(member int) {
    self.Set("x", member)
}

// Y position within the image to cut from.
func (self *Frame) GetY() int{
    return self.Get("y").Int()
}

// Y position within the image to cut from.
func (self *Frame) SetY(member int) {
    self.Set("y", member)
}

// Width of the frame.
func (self *Frame) GetWidth() int{
    return self.Get("width").Int()
}

// Width of the frame.
func (self *Frame) SetWidth(member int) {
    self.Set("width", member)
}

// Height of the frame.
func (self *Frame) GetHeight() int{
    return self.Get("height").Int()
}

// Height of the frame.
func (self *Frame) SetHeight(member int) {
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
func (self *Frame) GetCenterX() int{
    return self.Get("centerX").Int()
}

// Center X position within the image to cut from.
func (self *Frame) SetCenterX(member int) {
    self.Set("centerX", member)
}

// Center Y position within the image to cut from.
func (self *Frame) GetCenterY() int{
    return self.Get("centerY").Int()
}

// Center Y position within the image to cut from.
func (self *Frame) SetCenterY(member int) {
    self.Set("centerY", member)
}

// The distance from the top left to the bottom-right of this Frame.
func (self *Frame) GetDistance() int{
    return self.Get("distance").Int()
}

// The distance from the top left to the bottom-right of this Frame.
func (self *Frame) SetDistance(member int) {
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
func (self *Frame) GetSourceSizeW() int{
    return self.Get("sourceSizeW").Int()
}

// Width of the original sprite before it was trimmed.
func (self *Frame) SetSourceSizeW(member int) {
    self.Set("sourceSizeW", member)
}

// Height of the original sprite before it was trimmed.
func (self *Frame) GetSourceSizeH() int{
    return self.Get("sourceSizeH").Int()
}

// Height of the original sprite before it was trimmed.
func (self *Frame) SetSourceSizeH(member int) {
    self.Set("sourceSizeH", member)
}

// X position of the trimmed sprite inside original sprite.
func (self *Frame) GetSpriteSourceSizeX() int{
    return self.Get("spriteSourceSizeX").Int()
}

// X position of the trimmed sprite inside original sprite.
func (self *Frame) SetSpriteSourceSizeX(member int) {
    self.Set("spriteSourceSizeX", member)
}

// Y position of the trimmed sprite inside original sprite.
func (self *Frame) GetSpriteSourceSizeY() int{
    return self.Get("spriteSourceSizeY").Int()
}

// Y position of the trimmed sprite inside original sprite.
func (self *Frame) SetSpriteSourceSizeY(member int) {
    self.Set("spriteSourceSizeY", member)
}

// Width of the trimmed sprite.
func (self *Frame) GetSpriteSourceSizeW() int{
    return self.Get("spriteSourceSizeW").Int()
}

// Width of the trimmed sprite.
func (self *Frame) SetSpriteSourceSizeW(member int) {
    self.Set("spriteSourceSizeW", member)
}

// Height of the trimmed sprite.
func (self *Frame) GetSpriteSourceSizeH() int{
    return self.Get("spriteSourceSizeH").Int()
}

// Height of the trimmed sprite.
func (self *Frame) SetSpriteSourceSizeH(member int) {
    self.Set("spriteSourceSizeH", member)
}

// The right of the Frame (x + width).
func (self *Frame) GetRight() int{
    return self.Get("right").Int()
}

// The right of the Frame (x + width).
func (self *Frame) SetRight(member int) {
    self.Set("right", member)
}

// The bottom of the frame (y + height).
func (self *Frame) GetBottom() int{
    return self.Get("bottom").Int()
}

// The bottom of the frame (y + height).
func (self *Frame) SetBottom(member int) {
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
func (self *Frame) CloneI(args ...interface{}) *Frame{
    return &Frame{self.Call("clone", args)}
}

// Returns a Rectangle set to the dimensions of this Frame.
func (self *Frame) GetRectI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getRect", args)}
}
