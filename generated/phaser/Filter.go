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


// This is a base Filter class to use for any Phaser filter development.
// 
// The vast majority of filters (including all of those that ship with Phaser) use fragment shaders, and
// therefore only work in WebGL and are not supported by Canvas at all.
func NewFilter(game *Game, uniforms interface{}, fragmentSrc interface{}) *Filter {
    return &Filter{js.Global.Call("Phaser.Filter", game, uniforms, fragmentSrc)}
}

// This is a base Filter class to use for any Phaser filter development.
// 
// The vast majority of filters (including all of those that ship with Phaser) use fragment shaders, and
// therefore only work in WebGL and are not supported by Canvas at all.
func NewFilterI(args ...interface{}) *Filter {
    return &Filter{js.Global.Call("Phaser.Filter", args)}
}



// A reference to the currently running game.
func (self *Filter) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *Filter) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The const type of this object, either Phaser.WEBGL_FILTER or Phaser.CANVAS_FILTER.
func (self *Filter) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object, either Phaser.WEBGL_FILTER or Phaser.CANVAS_FILTER.
func (self *Filter) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Internal PIXI var.
func (self *Filter) GetDirtyA() bool{
    return self.Object.Get("dirty").Bool()
}

// Internal PIXI var.
func (self *Filter) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// Internal PIXI var.
func (self *Filter) GetPaddingA() int{
    return self.Object.Get("padding").Int()
}

// Internal PIXI var.
func (self *Filter) SetPaddingA(member int) {
    self.Object.Set("padding", member)
}

// The previous position of the pointer (we don't update the uniform if the same)
func (self *Filter) GetPrevPointA() *Point{
    return &Point{self.Object.Get("prevPoint")}
}

// The previous position of the pointer (we don't update the uniform if the same)
func (self *Filter) SetPrevPointA(member *Point) {
    self.Object.Set("prevPoint", member)
}

// Default uniform mappings. Compatible with ShaderToy and GLSLSandbox.
func (self *Filter) GetUniformsA() interface{}{
    return self.Object.Get("uniforms")
}

// Default uniform mappings. Compatible with ShaderToy and GLSLSandbox.
func (self *Filter) SetUniformsA(member interface{}) {
    self.Object.Set("uniforms", member)
}

// The fragment shader code.
func (self *Filter) GetFragmentSrcA() interface{}{
    return self.Object.Get("fragmentSrc")
}

// The fragment shader code.
func (self *Filter) SetFragmentSrcA(member interface{}) {
    self.Object.Set("fragmentSrc", member)
}

// The width (resolution uniform)
func (self *Filter) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width (resolution uniform)
func (self *Filter) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height (resolution uniform)
func (self *Filter) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height (resolution uniform)
func (self *Filter) SetHeightA(member int) {
    self.Object.Set("height", member)
}



// Should be over-ridden.
func (self *Filter) Init() {
    self.Object.Call("init")
}

// Should be over-ridden.
func (self *Filter) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Set the resolution uniforms on the filter.
func (self *Filter) SetResolution(width int, height int) {
    self.Object.Call("setResolution", width, height)
}

// Set the resolution uniforms on the filter.
func (self *Filter) SetResolutionI(args ...interface{}) {
    self.Object.Call("setResolution", args)
}

// Updates the filter.
func (self *Filter) Update() {
    self.Object.Call("update")
}

// Updates the filter.
func (self *Filter) Update1O(pointer *Pointer) {
    self.Object.Call("update", pointer)
}

// Updates the filter.
func (self *Filter) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld() *Image{
    return &Image{self.Object.Call("addToWorld")}
}

// Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld1O(x int) *Image{
    return &Image{self.Object.Call("addToWorld", x)}
}

// Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld2O(x int, y int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y)}
}

// Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld3O(x int, y int, width int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, width)}
}

// Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld4O(x int, y int, width int, height int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, width, height)}
}

// Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld5O(x int, y int, width int, height int, anchorX int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, width, height, anchorX)}
}

// Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorld6O(x int, y int, width int, height int, anchorX int, anchorY int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, width, height, anchorX, anchorY)}
}

// Creates a new Phaser.Image object using a blank texture and assigns 
// this Filter to it. The image is then added to the world.
// 
// If you don't provide width and height values then Filter.width and Filter.height are used.
// 
// If you do provide width and height values then this filter will be resized to match those
// values.
func (self *Filter) AddToWorldI(args ...interface{}) *Image{
    return &Image{self.Object.Call("addToWorld", args)}
}

// Clear down this Filter and null out references
func (self *Filter) Destroy() {
    self.Object.Call("destroy")
}

// Clear down this Filter and null out references
func (self *Filter) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
