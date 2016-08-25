// Package phaser Automatic generation for PIXI.PIXI
// generated file PIXI.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PIXI Namespace-class for [pixi.js](http://www.pixijs.com/).
// 
// Contains assorted static properties and enumerations.
type PIXI struct {
    *js.Object
}

// NewPIXI Namespace-class for [pixi.js](http://www.pixijs.com/).
// 
// Contains assorted static properties and enumerations.
func NewPIXI() *PIXI {
    return &PIXI{js.Global.Get("PIXI").Get("PIXI").New()}
}
// NewPIXII Namespace-class for [pixi.js](http://www.pixijs.com/).
// 
// Contains assorted static properties and enumerations.
func NewPIXII(args ...interface{}) *PIXI {
    return &PIXI{js.Global.Get("PIXI").Get("PIXI").New(args)}
}



// PIXI Binding conversion method to PIXI point 
func ToPIXI(jsStruct interface{}) *PIXI {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PIXI{Object: object}
	}
	return nil
}



// Game A reference to the Phaser Game instance that owns this Pixi renderer.
func (self *PIXI) Game() *PhaserGame{
    return &PhaserGame{self.Object.Get("game")}
}

// SetGameA A reference to the Phaser Game instance that owns this Pixi renderer.
func (self *PIXI) SetGameA(member *PhaserGame) {
    self.Object.Set("game", member)
}

// WEBGL_RENDERER empty description
func (self *PIXI) WEBGL_RENDERER() int{
    return self.Object.Get("WEBGL_RENDERER").Int()
}

// SetWEBGL_RENDERERA empty description
func (self *PIXI) SetWEBGL_RENDERERA(member int) {
    self.Object.Set("WEBGL_RENDERER", member)
}

// CANVAS_RENDERER empty description
func (self *PIXI) CANVAS_RENDERER() int{
    return self.Object.Get("CANVAS_RENDERER").Int()
}

// SetCANVAS_RENDERERA empty description
func (self *PIXI) SetCANVAS_RENDERERA(member int) {
    self.Object.Set("CANVAS_RENDERER", member)
}

// VERSION Version of pixi that is loaded.
func (self *PIXI) VERSION() string{
    return self.Object.Get("VERSION").String()
}

// SetVERSIONA Version of pixi that is loaded.
func (self *PIXI) SetVERSIONA(member string) {
    self.Object.Set("VERSION", member)
}

// PI_2 empty description
func (self *PIXI) PI_2() int{
    return self.Object.Get("PI_2").Int()
}

// SetPI_2A empty description
func (self *PIXI) SetPI_2A(member int) {
    self.Object.Set("PI_2", member)
}

// RAD_TO_DEG empty description
func (self *PIXI) RAD_TO_DEG() int{
    return self.Object.Get("RAD_TO_DEG").Int()
}

// SetRAD_TO_DEGA empty description
func (self *PIXI) SetRAD_TO_DEGA(member int) {
    self.Object.Set("RAD_TO_DEG", member)
}

// DEG_TO_RAD empty description
func (self *PIXI) DEG_TO_RAD() int{
    return self.Object.Get("DEG_TO_RAD").Int()
}

// SetDEG_TO_RADA empty description
func (self *PIXI) SetDEG_TO_RADA(member int) {
    self.Object.Set("DEG_TO_RAD", member)
}

// RETINA_PREFIX empty description
func (self *PIXI) RETINA_PREFIX() string{
    return self.Object.Get("RETINA_PREFIX").String()
}

// SetRETINA_PREFIXA empty description
func (self *PIXI) SetRETINA_PREFIXA(member string) {
    self.Object.Set("RETINA_PREFIX", member)
}

// DefaultRenderOptions The default render options if none are supplied to
// {{#crossLink "WebGLRenderer"}}{{/crossLink}} or {{#crossLink "CanvasRenderer"}}{{/crossLink}}.
func (self *PIXI) DefaultRenderOptions() interface{}{
    return self.Object.Get("defaultRenderOptions")
}

// SetDefaultRenderOptionsA The default render options if none are supplied to
// {{#crossLink "WebGLRenderer"}}{{/crossLink}} or {{#crossLink "CanvasRenderer"}}{{/crossLink}}.
func (self *PIXI) SetDefaultRenderOptionsA(member interface{}) {
    self.Object.Set("defaultRenderOptions", member)
}


// Hex2rgb Converts a hex color number to an [R, G, B] array
func (self *PIXI) Hex2rgb(hex int) {
    self.Object.Call("hex2rgb", hex)
}

// Hex2rgbI Converts a hex color number to an [R, G, B] array
func (self *PIXI) Hex2rgbI(args ...interface{}) {
    self.Object.Call("hex2rgb", args)
}

// Rgb2hex Converts a color as an [R, G, B] array to a hex number
func (self *PIXI) Rgb2hex(rgb []interface{}) {
    self.Object.Call("rgb2hex", rgb)
}

// Rgb2hexI Converts a color as an [R, G, B] array to a hex number
func (self *PIXI) Rgb2hexI(args ...interface{}) {
    self.Object.Call("rgb2hex", args)
}

// CanUseNewCanvasBlendModes Checks whether the Canvas BlendModes are supported by the current browser for drawImage
func (self *PIXI) CanUseNewCanvasBlendModes() bool{
    return self.Object.Call("canUseNewCanvasBlendModes").Bool()
}

// CanUseNewCanvasBlendModesI Checks whether the Canvas BlendModes are supported by the current browser for drawImage
func (self *PIXI) CanUseNewCanvasBlendModesI(args ...interface{}) bool{
    return self.Object.Call("canUseNewCanvasBlendModes", args).Bool()
}

// GetNextPowerOfTwo Given a number, this function returns the closest number that is a power of two
// this function is taken from Starling Framework as its pretty neat ;)
func (self *PIXI) GetNextPowerOfTwo(number int) int{
    return self.Object.Call("getNextPowerOfTwo", number).Int()
}

// GetNextPowerOfTwoI Given a number, this function returns the closest number that is a power of two
// this function is taken from Starling Framework as its pretty neat ;)
func (self *PIXI) GetNextPowerOfTwoI(args ...interface{}) int{
    return self.Object.Call("getNextPowerOfTwo", args).Int()
}

// IsPowerOfTwo checks if the given width and height make a power of two texture
func (self *PIXI) IsPowerOfTwo(width int, height int) bool{
    return self.Object.Call("isPowerOfTwo", width, height).Bool()
}

// IsPowerOfTwoI checks if the given width and height make a power of two texture
func (self *PIXI) IsPowerOfTwoI(args ...interface{}) bool{
    return self.Object.Call("isPowerOfTwo", args).Bool()
}

