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
func (self *FlexGrid) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *FlexGrid) SetGame(member Game) {
    self.Set("game", member)
}

// A reference to the ScaleManager.
func (self *FlexGrid) GetManager() ScaleManager{
    return ScaleManager{self.Get("manager")}
}

// A reference to the ScaleManager.
func (self *FlexGrid) SetManager(member ScaleManager) {
    self.Set("manager", member)
}

// -
func (self *FlexGrid) GetPositionCustom() interface{}{
    return self.Get("positionCustom")
}

// -
func (self *FlexGrid) SetPositionCustom(member interface{}) {
    self.Set("positionCustom", member)
}

// The scale factor based on the game dimensions vs. the scaled dimensions.
func (self *FlexGrid) GetScaleCustom() interface{}{
    return self.Get("scaleCustom")
}

// The scale factor based on the game dimensions vs. the scaled dimensions.
func (self *FlexGrid) SetScaleCustom(member interface{}) {
    self.Set("scaleCustom", member)
}



// Sets the core game size. This resets the w/h parameters and bounds.
func (self *FlexGrid) SetSizeI(args ...interface{}) {
    self.Call("setSize", args)
}

// A custom layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateCustomLayerI(args ...interface{}) FlexLayer{
    return FlexLayer{self.Call("createCustomLayer", args)}
}

// A fluid layer is centered on the game and maintains its aspect ratio as it scales up and down.
func (self *FlexGrid) CreateFluidLayerI(args ...interface{}) FlexLayer{
    return FlexLayer{self.Call("createFluidLayer", args)}
}

// A full layer is placed at 0,0 and extends to the full size of the game. Children are scaled according to the fluid ratios.
func (self *FlexGrid) CreateFullLayerI(args ...interface{}) FlexLayer{
    return FlexLayer{self.Call("createFullLayer", args)}
}

// A fixed layer is centered on the game and is the size of the required dimensions and is never scaled.
func (self *FlexGrid) CreateFixedLayerI(args ...interface{}) FlexLayer{
    return FlexLayer{self.Call("createFixedLayer", args)}
}

// Resets the layer children references
func (self *FlexGrid) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Called when the game container changes dimensions.
func (self *FlexGrid) OnResizeI(args ...interface{}) {
    self.Call("onResize", args)
}

// Updates all internal vars such as the bounds and scale values.
func (self *FlexGrid) RefreshI(args ...interface{}) {
    self.Call("refresh", args)
}

// Fits a sprites width to the bounds.
func (self *FlexGrid) FitSpriteI(args ...interface{}) {
    self.Call("fitSprite", args)
}

// Call in the render function to output the bounds rects.
func (self *FlexGrid) DebugI(args ...interface{}) {
    self.Call("debug", args)
}
