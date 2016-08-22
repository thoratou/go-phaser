// Automatic generation for Phaser.Filter
// generated file Filter.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// This is a base Filter class to use for any Phaser filter development.
// 
// The vast majority of filters (including all of those that ship with Phaser) use fragment shaders, and
// therefore only work in WebGL and are not supported by Canvas at all.
type Filter struct {
    *js.Object
}


// A reference to the currently running game.
func (self *Filter) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *Filter) SetGame(member *Game) {
    self.Set("game", member)
}

// The const type of this object, either Phaser.WEBGL_FILTER or Phaser.CANVAS_FILTER.
func (self *Filter) GetType() int{
    return self.Get("type").Int()
}

// The const type of this object, either Phaser.WEBGL_FILTER or Phaser.CANVAS_FILTER.
func (self *Filter) SetType(member int) {
    self.Set("type", member)
}

// Internal PIXI var.
func (self *Filter) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// Internal PIXI var.
func (self *Filter) SetDirty(member bool) {
    self.Set("dirty", member)
}

// Internal PIXI var.
func (self *Filter) GetPadding() int{
    return self.Get("padding").Int()
}

// Internal PIXI var.
func (self *Filter) SetPadding(member int) {
    self.Set("padding", member)
}

// The previous position of the pointer (we don't update the uniform if the same)
func (self *Filter) GetPrevPoint() *Point{
    return &Point{self.Get("prevPoint")}
}

// The previous position of the pointer (we don't update the uniform if the same)
func (self *Filter) SetPrevPoint(member *Point) {
    self.Set("prevPoint", member)
}

// Default uniform mappings. Compatible with ShaderToy and GLSLSandbox.
func (self *Filter) GetUniforms() interface{}{
    return self.Get("uniforms")
}

// Default uniform mappings. Compatible with ShaderToy and GLSLSandbox.
func (self *Filter) SetUniforms(member interface{}) {
    self.Set("uniforms", member)
}

// The fragment shader code.
func (self *Filter) GetFragmentSrc() interface{}{
    return self.Get("fragmentSrc")
}

// The fragment shader code.
func (self *Filter) SetFragmentSrc(member interface{}) {
    self.Set("fragmentSrc", member)
}

// The width (resolution uniform)
func (self *Filter) GetWidth() int{
    return self.Get("width").Int()
}

// The width (resolution uniform)
func (self *Filter) SetWidth(member int) {
    self.Set("width", member)
}

// The height (resolution uniform)
func (self *Filter) GetHeight() int{
    return self.Get("height").Int()
}

// The height (resolution uniform)
func (self *Filter) SetHeight(member int) {
    self.Set("height", member)
}



// Should be over-ridden.
func (self *Filter) InitI(args ...interface{}) {
    self.Call("init", args)
}

// Set the resolution uniforms on the filter.
func (self *Filter) SetResolutionI(args ...interface{}) {
    self.Call("setResolution", args)
}

// Updates the filter.
func (self *Filter) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorldI(args ...interface{}) *Image{
    return &Image{self.Call("addToWorld", args)}
}

// Clear down this Filter and null out references
func (self *Filter) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
