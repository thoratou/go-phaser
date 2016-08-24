// Package phaser Automatic generation for PIXI.CanvasTinter
// generated file CanvasTinter.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// CanvasTinter Utility methods for Sprite/Texture tinting.
type CanvasTinter struct {
    *js.Object
}

// NewCanvasTinter Utility methods for Sprite/Texture tinting.
func NewCanvasTinter() *CanvasTinter {
    return &CanvasTinter{js.Global.Get("PIXI").Get("CanvasTinter").New()}
}
// NewCanvasTinterI Utility methods for Sprite/Texture tinting.
func NewCanvasTinterI(args ...interface{}) *CanvasTinter {
    return &CanvasTinter{js.Global.Get("PIXI").Get("CanvasTinter").New(args)}
}



// CanHandleAlpha If the browser isn't capable of handling tinting with alpha this will be false.
// This property is only applicable if using tintWithPerPixel.
func (self *CanvasTinter) CanHandleAlpha() bool{
    return self.Object.Get("canHandleAlpha").Bool()
}

// SetCanHandleAlphaA If the browser isn't capable of handling tinting with alpha this will be false.
// This property is only applicable if using tintWithPerPixel.
func (self *CanvasTinter) SetCanHandleAlphaA(member bool) {
    self.Object.Set("canHandleAlpha", member)
}

// CanUseMultiply Whether or not the Canvas BlendModes are supported, consequently the ability to tint using the multiply method.
func (self *CanvasTinter) CanUseMultiply() bool{
    return self.Object.Get("canUseMultiply").Bool()
}

// SetCanUseMultiplyA Whether or not the Canvas BlendModes are supported, consequently the ability to tint using the multiply method.
func (self *CanvasTinter) SetCanUseMultiplyA(member bool) {
    self.Object.Set("canUseMultiply", member)
}


// GetTintedTexture Basically this method just needs a sprite and a color and tints the sprite with the given color.
func (self *CanvasTinter) GetTintedTexture(sprite *Sprite, color int) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("getTintedTexture", sprite, color))
}

// GetTintedTextureI Basically this method just needs a sprite and a color and tints the sprite with the given color.
func (self *CanvasTinter) GetTintedTextureI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("getTintedTexture", args))
}

// TintWithMultiply Tint a texture using the "multiply" operation.
func (self *CanvasTinter) TintWithMultiply(texture *Texture, color int, canvas *dom.HTMLCanvasElement) {
    self.Object.Call("tintWithMultiply", texture, color, canvas)
}

// TintWithMultiplyI Tint a texture using the "multiply" operation.
func (self *CanvasTinter) TintWithMultiplyI(args ...interface{}) {
    self.Object.Call("tintWithMultiply", args)
}

// TintPerPixel Tint a texture pixel per pixel.
func (self *CanvasTinter) TintPerPixel(texture *Texture, color int, canvas *dom.HTMLCanvasElement) {
    self.Object.Call("tintPerPixel", texture, color, canvas)
}

// TintPerPixelI Tint a texture pixel per pixel.
func (self *CanvasTinter) TintPerPixelI(args ...interface{}) {
    self.Object.Call("tintPerPixel", args)
}

// CheckInverseAlpha Checks if the browser correctly supports putImageData alpha channels.
func (self *CanvasTinter) CheckInverseAlpha() {
    self.Object.Call("checkInverseAlpha")
}

// CheckInverseAlphaI Checks if the browser correctly supports putImageData alpha channels.
func (self *CanvasTinter) CheckInverseAlphaI(args ...interface{}) {
    self.Object.Call("checkInverseAlpha", args)
}

// TintMethod The tinting method that will be used.
func (self *CanvasTinter) TintMethod() {
    self.Object.Call("tintMethod")
}

// TintMethodI The tinting method that will be used.
func (self *CanvasTinter) TintMethodI(args ...interface{}) {
    self.Object.Call("tintMethod", args)
}

