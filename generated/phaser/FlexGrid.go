// Automatic generation for Phaser.FlexGrid
// generated file FlexGrid.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// WARNING: This is an EXPERIMENTAL class. The API will change significantly in the coming versions and is incomplete.
// Please try to avoid using in production games with a long time to build.
// This is also why the documentation is incomplete.
// 
// FlexGrid is a a responsive grid manager that works in conjunction with the ScaleManager RESIZE scaling mode and FlexLayers
// to provide for game object positioning in a responsive manner.
type FlexGrid struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *FlexGrid) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *FlexGrid) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// A reference to the ScaleManager.
func (self *FlexGrid) GetManagerA() *ScaleManager{
    return &ScaleManager{self.Object.Get("manager")}
}

// A reference to the ScaleManager.
func (self *FlexGrid) SetManagerA(member *ScaleManager) {
    self.Object.Set("manager", member)
}

// -
func (self *FlexGrid) GetPositionCustomA() interface{}{
    return self.Object.Get("positionCustom")
}

// -
func (self *FlexGrid) SetPositionCustomA(member interface{}) {
    self.Object.Set("positionCustom", member)
}

// The scale factor based on the game dimensions vs. the scaled dimensions.
func (self *FlexGrid) GetScaleCustomA() interface{}{
    return self.Object.Get("scaleCustom")
}

// The scale factor based on the game dimensions vs. the scaled dimensions.
func (self *FlexGrid) SetScaleCustomA(member interface{}) {
    self.Object.Set("scaleCustom", member)
}



// Sets the core game size. This resets the w/h parameters and bounds.
func (self *FlexGrid) SetSize(width int, height int) {
    self.Object.Call("setSize", width, height)
}

// Sets the core game size. This resets the w/h parameters and bounds.
func (self *FlexGrid) SetSizeI(args ...interface{}) {
    self.Object.Call("setSize", args)
}

// A custom layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateCustomLayer(width int, height int) *FlexLayer{
    return &FlexLayer{self.Object.Call("createCustomLayer", width, height)}
}

// A custom layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateCustomLayer1O(width int, height int, children []DisplayObject) *FlexLayer{
    return &FlexLayer{self.Object.Call("createCustomLayer", width, height, children)}
}

// A custom layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateCustomLayerI(args ...interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createCustomLayer", args)}
}

// A fluid layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateFluidLayer() *FlexLayer{
    return &FlexLayer{self.Object.Call("createFluidLayer")}
}

// A fluid layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateFluidLayer1O(children []interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFluidLayer", children)}
}

// A fluid layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateFluidLayerI(args ...interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFluidLayer", args)}
}

// A full layer is placed at 0,0 and extends to the full size of the game. Children are scaled according to the fluid ratios.
func (self *FlexGrid) CreateFullLayer() *FlexLayer{
    return &FlexLayer{self.Object.Call("createFullLayer")}
}

// A full layer is placed at 0,0 and extends to the full size of the game. Children are scaled according to the fluid ratios.
func (self *FlexGrid) CreateFullLayer1O(children []interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFullLayer", children)}
}

// A full layer is placed at 0,0 and extends to the full size of the game. Children are scaled according to the fluid ratios.
func (self *FlexGrid) CreateFullLayerI(args ...interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFullLayer", args)}
}

// A fixed layer is centered on the game and is the size of the required dimensions and is never scaled.
func (self *FlexGrid) CreateFixedLayer() *FlexLayer{
    return &FlexLayer{self.Object.Call("createFixedLayer")}
}

// A fixed layer is centered on the game and is the size of the required dimensions and is never scaled.
func (self *FlexGrid) CreateFixedLayer1O(children []DisplayObject) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFixedLayer", children)}
}

// A fixed layer is centered on the game and is the size of the required dimensions and is never scaled.
func (self *FlexGrid) CreateFixedLayerI(args ...interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFixedLayer", args)}
}

// Resets the layer children references
func (self *FlexGrid) Reset() {
    self.Object.Call("reset")
}

// Resets the layer children references
func (self *FlexGrid) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Called when the game container changes dimensions.
func (self *FlexGrid) OnResize(width int, height int) {
    self.Object.Call("onResize", width, height)
}

// Called when the game container changes dimensions.
func (self *FlexGrid) OnResizeI(args ...interface{}) {
    self.Object.Call("onResize", args)
}

// Updates all internal vars such as the bounds and scale values.
func (self *FlexGrid) Refresh() {
    self.Object.Call("refresh")
}

// Updates all internal vars such as the bounds and scale values.
func (self *FlexGrid) RefreshI(args ...interface{}) {
    self.Object.Call("refresh", args)
}

// Fits a sprites width to the bounds.
func (self *FlexGrid) FitSprite(sprite *Sprite) {
    self.Object.Call("fitSprite", sprite)
}

// Fits a sprites width to the bounds.
func (self *FlexGrid) FitSpriteI(args ...interface{}) {
    self.Object.Call("fitSprite", args)
}

// Call in the render function to output the bounds rects.
func (self *FlexGrid) Debug() {
    self.Object.Call("debug")
}

// Call in the render function to output the bounds rects.
func (self *FlexGrid) DebugI(args ...interface{}) {
    self.Object.Call("debug", args)
}
