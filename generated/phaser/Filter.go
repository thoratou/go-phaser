// Package phaser Automatic generation for Phaser.Filter
// generated file Filter.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Filter This is a base Filter class to use for any Phaser filter development.
// 
// The vast majority of filters (including all of those that ship with Phaser) use fragment shaders, and
// therefore only work in WebGL and are not supported by Canvas at all.
type Filter struct {
    *js.Object
}

// NewFilter This is a base Filter class to use for any Phaser filter development.
// 
// The vast majority of filters (including all of those that ship with Phaser) use fragment shaders, and
// therefore only work in WebGL and are not supported by Canvas at all.
func NewFilter(game *Game, uniforms interface{}, fragmentSrc interface{}) *Filter {
    return &Filter{js.Global.Get("Phaser").Get("Filter").New(game, uniforms, fragmentSrc)}
}
// NewFilterI This is a base Filter class to use for any Phaser filter development.
// 
// The vast majority of filters (including all of those that ship with Phaser) use fragment shaders, and
// therefore only work in WebGL and are not supported by Canvas at all.
func NewFilterI(args ...interface{}) *Filter {
    return &Filter{js.Global.Get("Phaser").Get("Filter").New(args)}
}



// Filter Binding conversion method to Filter point 
func ToFilter(jsStruct interface{}) *Filter {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Filter{Object: object}
	}
	return nil
}



// Game A reference to the currently running game.
func (self *Filter) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *Filter) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Type The const type of this object, either Phaser.WEBGL_FILTER or Phaser.CANVAS_FILTER.
func (self *Filter) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object, either Phaser.WEBGL_FILTER or Phaser.CANVAS_FILTER.
func (self *Filter) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Dirty Internal PIXI var.
func (self *Filter) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// SetDirtyA Internal PIXI var.
func (self *Filter) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// Padding Internal PIXI var.
func (self *Filter) Padding() int{
    return self.Object.Get("padding").Int()
}

// SetPaddingA Internal PIXI var.
func (self *Filter) SetPaddingA(member int) {
    self.Object.Set("padding", member)
}

// PrevPoint The previous position of the pointer (we don't update the uniform if the same)
func (self *Filter) PrevPoint() *Point{
    return &Point{self.Object.Get("prevPoint")}
}

// SetPrevPointA The previous position of the pointer (we don't update the uniform if the same)
func (self *Filter) SetPrevPointA(member *Point) {
    self.Object.Set("prevPoint", member)
}

// Uniforms Default uniform mappings. Compatible with ShaderToy and GLSLSandbox.
func (self *Filter) Uniforms() interface{}{
    return self.Object.Get("uniforms")
}

// SetUniformsA Default uniform mappings. Compatible with ShaderToy and GLSLSandbox.
func (self *Filter) SetUniformsA(member interface{}) {
    self.Object.Set("uniforms", member)
}

// FragmentSrc The fragment shader code.
func (self *Filter) FragmentSrc() interface{}{
    return self.Object.Get("fragmentSrc")
}

// SetFragmentSrcA The fragment shader code.
func (self *Filter) SetFragmentSrcA(member interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// Width The width (resolution uniform)
func (self *Filter) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width (resolution uniform)
func (self *Filter) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height (resolution uniform)
func (self *Filter) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height (resolution uniform)
func (self *Filter) SetHeightA(member int) {
    self.Object.Set("height", member)
}


// Init Should be over-ridden.
func (self *Filter) Init() {
    self.Object.Call("init")
}

// InitI Should be over-ridden.
func (self *Filter) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// SetResolution Set the resolution uniforms on the filter.
func (self *Filter) SetResolution(width int, height int) {
    self.Object.Call("setResolution", width, height)
}

// SetResolutionI Set the resolution uniforms on the filter.
func (self *Filter) SetResolutionI(args ...interface{}) {
    self.Object.Call("setResolution", args)
}

// Update Updates the filter.
func (self *Filter) Update() {
    self.Object.Call("update")
}

// Update1O Updates the filter.
func (self *Filter) Update1O(pointer *Pointer) {
    self.Object.Call("update", pointer)
}

// UpdateI Updates the filter.
func (self *Filter) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// AddToWorld Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld() *Image{
    return &Image{self.Object.Call("addToWorld")}
}

// AddToWorld1O Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld1O(x int) *Image{
    return &Image{self.Object.Call("addToWorld", x)}
}

// AddToWorld2O Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld2O(x int, y int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y)}
}

// AddToWorld3O Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld3O(x int, y int, width int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, width)}
}

// AddToWorld4O Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld4O(x int, y int, width int, height int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, width, height)}
}

// AddToWorld5O Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld5O(x int, y int, width int, height int, anchorX int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, width, height, anchorX)}
}

// AddToWorld6O Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld6O(x int, y int, width int, height int, anchorX int, anchorY int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, width, height, anchorX, anchorY)}
}

// AddToWorldI Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorldI(args ...interface{}) *Image{
    return &Image{self.Object.Call("addToWorld", args)}
}

// Destroy Clear down this Filter and null out references
func (self *Filter) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Clear down this Filter and null out references
func (self *Filter) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

