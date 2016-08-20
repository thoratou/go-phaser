// Automatic generation for PIXI.PIXI
// generated file PIXI.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Namespace-class for [pixi.js](http://www.pixijs.com/).
// 
// Contains assorted static properties and enumerations.
type PIXI struct {
    *js.Object
}


// A reference to the Phaser Game instance that owns this Pixi renderer.
func (self *PIXI) GetGame() PhaserGame{
    return PhaserGame{self.Get("game")}
}

// A reference to the Phaser Game instance that owns this Pixi renderer.
func (self *PIXI) SetGame(member PhaserGame) {
    self.Set("game", member)
}

// 
func (self *PIXI) GetWEBGL_RENDERER() float64{
    return self.Get("WEBGL_RENDERER").Float()
}

// 
func (self *PIXI) SetWEBGL_RENDERER(member float64) {
    self.Set("WEBGL_RENDERER", member)
}

// 
func (self *PIXI) GetCANVAS_RENDERER() float64{
    return self.Get("CANVAS_RENDERER").Float()
}

// 
func (self *PIXI) SetCANVAS_RENDERER(member float64) {
    self.Set("CANVAS_RENDERER", member)
}

// Version of pixi that is loaded.
func (self *PIXI) GetVERSION() string{
    return self.Get("VERSION").String()
}

// Version of pixi that is loaded.
func (self *PIXI) SetVERSION(member string) {
    self.Set("VERSION", member)
}

// 
func (self *PIXI) GetPI_2() float64{
    return self.Get("PI_2").Float()
}

// 
func (self *PIXI) SetPI_2(member float64) {
    self.Set("PI_2", member)
}

// 
func (self *PIXI) GetRAD_TO_DEG() float64{
    return self.Get("RAD_TO_DEG").Float()
}

// 
func (self *PIXI) SetRAD_TO_DEG(member float64) {
    self.Set("RAD_TO_DEG", member)
}

// 
func (self *PIXI) GetDEG_TO_RAD() float64{
    return self.Get("DEG_TO_RAD").Float()
}

// 
func (self *PIXI) SetDEG_TO_RAD(member float64) {
    self.Set("DEG_TO_RAD", member)
}

// 
func (self *PIXI) GetRETINA_PREFIX() string{
    return self.Get("RETINA_PREFIX").String()
}

// 
func (self *PIXI) SetRETINA_PREFIX(member string) {
    self.Set("RETINA_PREFIX", member)
}

// The default render options if none are supplied to
// {{#crossLink "WebGLRenderer"}}{{/crossLink}} or {{#crossLink "CanvasRenderer"}}{{/crossLink}}.
func (self *PIXI) GetDefaultRenderOptions() interface{}{
    return self.Get("defaultRenderOptions")
}

// The default render options if none are supplied to
// {{#crossLink "WebGLRenderer"}}{{/crossLink}} or {{#crossLink "CanvasRenderer"}}{{/crossLink}}.
func (self *PIXI) SetDefaultRenderOptions(member interface{}) {
    self.Set("defaultRenderOptions", member)
}



// Converts a hex color number to an [R, G, B] array
func (self *PIXI) Hex2rgbI(args ...interface{}) {
    self.Call("hex2rgb", args)
}

// Converts a color as an [R, G, B] array to a hex number
func (self *PIXI) Rgb2hexI(args ...interface{}) {
    self.Call("rgb2hex", args)
}

// Checks whether the Canvas BlendModes are supported by the current browser for drawImage
func (self *PIXI) CanUseNewCanvasBlendModesI(args ...interface{}) bool{
    return self.Call("canUseNewCanvasBlendModes", args).Bool()
}

// Given a number, this function returns the closest number that is a power of two
// this function is taken from Starling Framework as its pretty neat ;)
func (self *PIXI) GetNextPowerOfTwoI(args ...interface{}) float64{
    return self.Call("getNextPowerOfTwo", args).Float()
}

// checks if the given width and height make a power of two texture
func (self *PIXI) IsPowerOfTwoI(args ...interface{}) bool{
    return self.Call("isPowerOfTwo", args).Bool()
}
