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


// Namespace-class for [pixi.js](http://www.pixijs.com/).
// 
// Contains assorted static properties and enumerations.
func NewPIXI() *PIXI {
    return &PIXI{js.Global.Call("PIXI.PIXI")}
}

// Namespace-class for [pixi.js](http://www.pixijs.com/).
// 
// Contains assorted static properties and enumerations.
func NewPIXII(args ...interface{}) *PIXI {
    return &PIXI{js.Global.Call("PIXI.PIXI", args)}
}



// A reference to the Phaser Game instance that owns this Pixi renderer.
func (self *PIXI) GetGameA() *PhaserGame{
    return &PhaserGame{self.Object.Get("game")}
}

// A reference to the Phaser Game instance that owns this Pixi renderer.
func (self *PIXI) SetGameA(member *PhaserGame) {
    self.Object.Set("game", member)
}

// 
func (self *PIXI) GetWEBGL_RENDERERA() int{
    return self.Object.Get("WEBGL_RENDERER").Int()
}

// 
func (self *PIXI) SetWEBGL_RENDERERA(member int) {
    self.Object.Set("WEBGL_RENDERER", member)
}

// 
func (self *PIXI) GetCANVAS_RENDERERA() int{
    return self.Object.Get("CANVAS_RENDERER").Int()
}

// 
func (self *PIXI) SetCANVAS_RENDERERA(member int) {
    self.Object.Set("CANVAS_RENDERER", member)
}

// Version of pixi that is loaded.
func (self *PIXI) GetVERSIONA() string{
    return self.Object.Get("VERSION").String()
}

// Version of pixi that is loaded.
func (self *PIXI) SetVERSIONA(member string) {
    self.Object.Set("VERSION", member)
}

// 
func (self *PIXI) GetPI_2A() int{
    return self.Object.Get("PI_2").Int()
}

// 
func (self *PIXI) SetPI_2A(member int) {
    self.Object.Set("PI_2", member)
}

// 
func (self *PIXI) GetRAD_TO_DEGA() int{
    return self.Object.Get("RAD_TO_DEG").Int()
}

// 
func (self *PIXI) SetRAD_TO_DEGA(member int) {
    self.Object.Set("RAD_TO_DEG", member)
}

// 
func (self *PIXI) GetDEG_TO_RADA() int{
    return self.Object.Get("DEG_TO_RAD").Int()
}

// 
func (self *PIXI) SetDEG_TO_RADA(member int) {
    self.Object.Set("DEG_TO_RAD", member)
}

// 
func (self *PIXI) GetRETINA_PREFIXA() string{
    return self.Object.Get("RETINA_PREFIX").String()
}

// 
func (self *PIXI) SetRETINA_PREFIXA(member string) {
    self.Object.Set("RETINA_PREFIX", member)
}

// The default render options if none are supplied to
// {{#crossLink "WebGLRenderer"}}{{/crossLink}} or {{#crossLink "CanvasRenderer"}}{{/crossLink}}.
func (self *PIXI) GetDefaultRenderOptionsA() interface{}{
    return self.Object.Get("defaultRenderOptions")
}

// The default render options if none are supplied to
// {{#crossLink "WebGLRenderer"}}{{/crossLink}} or {{#crossLink "CanvasRenderer"}}{{/crossLink}}.
func (self *PIXI) SetDefaultRenderOptionsA(member interface{}) {
    self.Object.Set("defaultRenderOptions", member)
}



// Converts a hex color number to an [R, G, B] array
func (self *PIXI) Hex2rgb(hex int) {
    self.Object.Call("hex2rgb", hex)
}

// Converts a hex color number to an [R, G, B] array
func (self *PIXI) Hex2rgbI(args ...interface{}) {
    self.Object.Call("hex2rgb", args)
}

// Converts a color as an [R, G, B] array to a hex number
func (self *PIXI) Rgb2hex(rgb []interface{}) {
    self.Object.Call("rgb2hex", rgb)
}

// Converts a color as an [R, G, B] array to a hex number
func (self *PIXI) Rgb2hexI(args ...interface{}) {
    self.Object.Call("rgb2hex", args)
}

// Checks whether the Canvas BlendModes are supported by the current browser for drawImage
func (self *PIXI) CanUseNewCanvasBlendModes() bool{
    return self.Object.Call("canUseNewCanvasBlendModes").Bool()
}

// Checks whether the Canvas BlendModes are supported by the current browser for drawImage
func (self *PIXI) CanUseNewCanvasBlendModesI(args ...interface{}) bool{
    return self.Object.Call("canUseNewCanvasBlendModes", args).Bool()
}

// Given a number, this function returns the closest number that is a power of two
// this function is taken from Starling Framework as its pretty neat ;)
func (self *PIXI) GetNextPowerOfTwo(number int) int{
    return self.Object.Call("getNextPowerOfTwo", number).Int()
}

// Given a number, this function returns the closest number that is a power of two
// this function is taken from Starling Framework as its pretty neat ;)
func (self *PIXI) GetNextPowerOfTwoI(args ...interface{}) int{
    return self.Object.Call("getNextPowerOfTwo", args).Int()
}

// checks if the given width and height make a power of two texture
func (self *PIXI) IsPowerOfTwo(width int, height int) bool{
    return self.Object.Call("isPowerOfTwo", width, height).Bool()
}

// checks if the given width and height make a power of two texture
func (self *PIXI) IsPowerOfTwoI(args ...interface{}) bool{
    return self.Object.Call("isPowerOfTwo", args).Bool()
}
