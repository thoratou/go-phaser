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


// A Frame is a single frame of an animation and is part of a FrameData collection.
func NewFrame(index int, x int, y int, width int, height int, name string) *Frame {
    return &Frame{js.Global.Get("Phaser").Get("Frame").New(index, x, y, width, height, name)}
}

// A Frame is a single frame of an animation and is part of a FrameData collection.
func NewFrameI(args ...interface{}) *Frame {
    return &Frame{js.Global.Get("Phaser").Get("Frame").New(args)}
}



// The index of this Frame within the FrameData set it is being added to.
func (self *Frame) Index() int{
    return self.Object.Get("index").Int()
}

// The index of this Frame within the FrameData set it is being added to.
func (self *Frame) SetIndexA(member int) {
    self.Object.Set("index", member)
}

// X position within the image to cut from.
func (self *Frame) X() int{
    return self.Object.Get("x").Int()
}

// X position within the image to cut from.
func (self *Frame) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y position within the image to cut from.
func (self *Frame) Y() int{
    return self.Object.Get("y").Int()
}

// Y position within the image to cut from.
func (self *Frame) SetYA(member int) {
    self.Object.Set("y", member)
}

// Width of the frame.
func (self *Frame) Width() int{
    return self.Object.Get("width").Int()
}

// Width of the frame.
func (self *Frame) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height of the frame.
func (self *Frame) Height() int{
    return self.Object.Get("height").Int()
}

// Height of the frame.
func (self *Frame) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Useful for Texture Atlas files (is set to the filename value).
func (self *Frame) Name() string{
    return self.Object.Get("name").String()
}

// Useful for Texture Atlas files (is set to the filename value).
func (self *Frame) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Center X position within the image to cut from.
func (self *Frame) CenterX() int{
    return self.Object.Get("centerX").Int()
}

// Center X position within the image to cut from.
func (self *Frame) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// Center Y position within the image to cut from.
func (self *Frame) CenterY() int{
    return self.Object.Get("centerY").Int()
}

// Center Y position within the image to cut from.
func (self *Frame) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// The distance from the top left to the bottom-right of this Frame.
func (self *Frame) Distance() int{
    return self.Object.Get("distance").Int()
}

// The distance from the top left to the bottom-right of this Frame.
func (self *Frame) SetDistanceA(member int) {
    self.Object.Set("distance", member)
}

// Rotated? (not yet implemented)
func (self *Frame) Rotated() bool{
    return self.Object.Get("rotated").Bool()
}

// Rotated? (not yet implemented)
func (self *Frame) SetRotatedA(member bool) {
    self.Object.Set("rotated", member)
}

// Either 'cw' or 'ccw', rotation is always 90 degrees.
func (self *Frame) RotationDirection() string{
    return self.Object.Get("rotationDirection").String()
}

// Either 'cw' or 'ccw', rotation is always 90 degrees.
func (self *Frame) SetRotationDirectionA(member string) {
    self.Object.Set("rotationDirection", member)
}

// Was it trimmed when packed?
func (self *Frame) Trimmed() bool{
    return self.Object.Get("trimmed").Bool()
}

// Was it trimmed when packed?
func (self *Frame) SetTrimmedA(member bool) {
    self.Object.Set("trimmed", member)
}

// Width of the original sprite before it was trimmed.
func (self *Frame) SourceSizeW() int{
    return self.Object.Get("sourceSizeW").Int()
}

// Width of the original sprite before it was trimmed.
func (self *Frame) SetSourceSizeWA(member int) {
    self.Object.Set("sourceSizeW", member)
}

// Height of the original sprite before it was trimmed.
func (self *Frame) SourceSizeH() int{
    return self.Object.Get("sourceSizeH").Int()
}

// Height of the original sprite before it was trimmed.
func (self *Frame) SetSourceSizeHA(member int) {
    self.Object.Set("sourceSizeH", member)
}

// X position of the trimmed sprite inside original sprite.
func (self *Frame) SpriteSourceSizeX() int{
    return self.Object.Get("spriteSourceSizeX").Int()
}

// X position of the trimmed sprite inside original sprite.
func (self *Frame) SetSpriteSourceSizeXA(member int) {
    self.Object.Set("spriteSourceSizeX", member)
}

// Y position of the trimmed sprite inside original sprite.
func (self *Frame) SpriteSourceSizeY() int{
    return self.Object.Get("spriteSourceSizeY").Int()
}

// Y position of the trimmed sprite inside original sprite.
func (self *Frame) SetSpriteSourceSizeYA(member int) {
    self.Object.Set("spriteSourceSizeY", member)
}

// Width of the trimmed sprite.
func (self *Frame) SpriteSourceSizeW() int{
    return self.Object.Get("spriteSourceSizeW").Int()
}

// Width of the trimmed sprite.
func (self *Frame) SetSpriteSourceSizeWA(member int) {
    self.Object.Set("spriteSourceSizeW", member)
}

// Height of the trimmed sprite.
func (self *Frame) SpriteSourceSizeH() int{
    return self.Object.Get("spriteSourceSizeH").Int()
}

// Height of the trimmed sprite.
func (self *Frame) SetSpriteSourceSizeHA(member int) {
    self.Object.Set("spriteSourceSizeH", member)
}

// The right of the Frame (x + width).
func (self *Frame) Right() int{
    return self.Object.Get("right").Int()
}

// The right of the Frame (x + width).
func (self *Frame) SetRightA(member int) {
    self.Object.Set("right", member)
}

// The bottom of the frame (y + height).
func (self *Frame) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// The bottom of the frame (y + height).
func (self *Frame) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}



// Adjusts of all the Frame properties based on the given width and height values.
func (self *Frame) Resize(width int, height int) {
    self.Object.Call("resize", width, height)
}

// Adjusts of all the Frame properties based on the given width and height values.
func (self *Frame) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// If the frame was trimmed when added to the Texture Atlas this records the trim and source data.
func (self *Frame) SetTrim(trimmed bool, actualWidth int, actualHeight int, destX int, destY int, destWidth int, destHeight int) {
    self.Object.Call("setTrim", trimmed, actualWidth, actualHeight, destX, destY, destWidth, destHeight)
}

// If the frame was trimmed when added to the Texture Atlas this records the trim and source data.
func (self *Frame) SetTrimI(args ...interface{}) {
    self.Object.Call("setTrim", args)
}

// Clones this Frame into a new Phaser.Frame object and returns it.
// Note that all properties are cloned, including the name, index and UUID.
func (self *Frame) Clone() *Frame{
    return &Frame{self.Object.Call("clone")}
}

// Clones this Frame into a new Phaser.Frame object and returns it.
// Note that all properties are cloned, including the name, index and UUID.
func (self *Frame) CloneI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("clone", args)}
}

// Returns a Rectangle set to the dimensions of this Frame.
func (self *Frame) GetRect() *Rectangle{
    return &Rectangle{self.Object.Call("getRect")}
}

// Returns a Rectangle set to the dimensions of this Frame.
func (self *Frame) GetRect1O(out *Rectangle) *Rectangle{
    return &Rectangle{self.Object.Call("getRect", out)}
}

// Returns a Rectangle set to the dimensions of this Frame.
func (self *Frame) GetRectI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getRect", args)}
}
