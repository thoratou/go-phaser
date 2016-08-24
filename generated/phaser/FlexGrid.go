// Package phaser Automatic generation for Phaser.FlexGrid
// generated file FlexGrid.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// FlexGrid WARNING: This is an EXPERIMENTAL class. The API will change significantly in the coming versions and is incomplete.
// Please try to avoid using in production games with a long time to build.
// This is also why the documentation is incomplete.
// 
// FlexGrid is a a responsive grid manager that works in conjunction with the ScaleManager RESIZE scaling mode and FlexLayers
// to provide for game object positioning in a responsive manner.
type FlexGrid struct {
    *js.Object
}

// NewFlexGrid WARNING: This is an EXPERIMENTAL class. The API will change significantly in the coming versions and is incomplete.
// Please try to avoid using in production games with a long time to build.
// This is also why the documentation is incomplete.
// 
// FlexGrid is a a responsive grid manager that works in conjunction with the ScaleManager RESIZE scaling mode and FlexLayers
// to provide for game object positioning in a responsive manner.
func NewFlexGrid(manager *ScaleManager, width int, height int) *FlexGrid {
    return &FlexGrid{js.Global.Get("Phaser").Get("FlexGrid").New(manager, width, height)}
}
// NewFlexGridI WARNING: This is an EXPERIMENTAL class. The API will change significantly in the coming versions and is incomplete.
// Please try to avoid using in production games with a long time to build.
// This is also why the documentation is incomplete.
// 
// FlexGrid is a a responsive grid manager that works in conjunction with the ScaleManager RESIZE scaling mode and FlexLayers
// to provide for game object positioning in a responsive manner.
func NewFlexGridI(args ...interface{}) *FlexGrid {
    return &FlexGrid{js.Global.Get("Phaser").Get("FlexGrid").New(args)}
}



// Game A reference to the currently running Game.
func (self *FlexGrid) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *FlexGrid) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Manager A reference to the ScaleManager.
func (self *FlexGrid) Manager() *ScaleManager{
    return &ScaleManager{self.Object.Get("manager")}
}

// SetManagerA A reference to the ScaleManager.
func (self *FlexGrid) SetManagerA(member *ScaleManager) {
    self.Object.Set("manager", member)
}

// PositionCustom -
func (self *FlexGrid) PositionCustom() interface{}{
    return self.Object.Get("positionCustom")
}

// SetPositionCustomA -
func (self *FlexGrid) SetPositionCustomA(member interface{}) {
    self.Object.Set("positionCustom", member)
}

// ScaleCustom The scale factor based on the game dimensions vs. the scaled dimensions.
func (self *FlexGrid) ScaleCustom() interface{}{
    return self.Object.Get("scaleCustom")
}

// SetScaleCustomA The scale factor based on the game dimensions vs. the scaled dimensions.
func (self *FlexGrid) SetScaleCustomA(member interface{}) {
    self.Object.Set("scaleCustom", member)
}


// SetSize Sets the core game size. This resets the w/h parameters and bounds.
func (self *FlexGrid) SetSize(width int, height int) {
    self.Object.Call("setSize", width, height)
}

// SetSizeI Sets the core game size. This resets the w/h parameters and bounds.
func (self *FlexGrid) SetSizeI(args ...interface{}) {
    self.Object.Call("setSize", args)
}

// CreateCustomLayer A custom layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateCustomLayer(width int, height int) *FlexLayer{
    return &FlexLayer{self.Object.Call("createCustomLayer", width, height)}
}

// CreateCustomLayer1O A custom layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateCustomLayer1O(width int, height int, children []DisplayObject) *FlexLayer{
    return &FlexLayer{self.Object.Call("createCustomLayer", width, height, children)}
}

// CreateCustomLayerI A custom layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateCustomLayerI(args ...interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createCustomLayer", args)}
}

// CreateFluidLayer A fluid layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateFluidLayer() *FlexLayer{
    return &FlexLayer{self.Object.Call("createFluidLayer")}
}

// CreateFluidLayer1O A fluid layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateFluidLayer1O(children []interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFluidLayer", children)}
}

// CreateFluidLayerI A fluid layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateFluidLayerI(args ...interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFluidLayer", args)}
}

// CreateFullLayer A full layer is placed at 0,0 and extends to the full size of the game. Children are scaled according to the fluid ratios.
func (self *FlexGrid) CreateFullLayer() *FlexLayer{
    return &FlexLayer{self.Object.Call("createFullLayer")}
}

// CreateFullLayer1O A full layer is placed at 0,0 and extends to the full size of the game. Children are scaled according to the fluid ratios.
func (self *FlexGrid) CreateFullLayer1O(children []interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFullLayer", children)}
}

// CreateFullLayerI A full layer is placed at 0,0 and extends to the full size of the game. Children are scaled according to the fluid ratios.
func (self *FlexGrid) CreateFullLayerI(args ...interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFullLayer", args)}
}

// CreateFixedLayer A fixed layer is centered on the game and is the size of the required dimensions and is never scaled.
func (self *FlexGrid) CreateFixedLayer() *FlexLayer{
    return &FlexLayer{self.Object.Call("createFixedLayer")}
}

// CreateFixedLayer1O A fixed layer is centered on the game and is the size of the required dimensions and is never scaled.
func (self *FlexGrid) CreateFixedLayer1O(children []DisplayObject) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFixedLayer", children)}
}

// CreateFixedLayerI A fixed layer is centered on the game and is the size of the required dimensions and is never scaled.
func (self *FlexGrid) CreateFixedLayerI(args ...interface{}) *FlexLayer{
    return &FlexLayer{self.Object.Call("createFixedLayer", args)}
}

// Reset Resets the layer children references
func (self *FlexGrid) Reset() {
    self.Object.Call("reset")
}

// ResetI Resets the layer children references
func (self *FlexGrid) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// OnResize Called when the game container changes dimensions.
func (self *FlexGrid) OnResize(width int, height int) {
    self.Object.Call("onResize", width, height)
}

// OnResizeI Called when the game container changes dimensions.
func (self *FlexGrid) OnResizeI(args ...interface{}) {
    self.Object.Call("onResize", args)
}

// Refresh Updates all internal vars such as the bounds and scale values.
func (self *FlexGrid) Refresh() {
    self.Object.Call("refresh")
}

// RefreshI Updates all internal vars such as the bounds and scale values.
func (self *FlexGrid) RefreshI(args ...interface{}) {
    self.Object.Call("refresh", args)
}

// FitSprite Fits a sprites width to the bounds.
func (self *FlexGrid) FitSprite(sprite *Sprite) {
    self.Object.Call("fitSprite", sprite)
}

// FitSpriteI Fits a sprites width to the bounds.
func (self *FlexGrid) FitSpriteI(args ...interface{}) {
    self.Object.Call("fitSprite", args)
}

// Debug Call in the render function to output the bounds rects.
func (self *FlexGrid) Debug() {
    self.Object.Call("debug")
}

// DebugI Call in the render function to output the bounds rects.
func (self *FlexGrid) DebugI(args ...interface{}) {
    self.Object.Call("debug", args)
}

