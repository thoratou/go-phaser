// Automatic generation for PIXI.CanvasTinter
// generated file CanvasTinter.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// Utility methods for Sprite/Texture tinting.
type CanvasTinter struct {
    *js.Object
}


// If the browser isn't capable of handling tinting with alpha this will be false.
// This property is only applicable if using tintWithPerPixel.
func (self *CanvasTinter) GetCanHandleAlpha() bool{
    return self.Get("canHandleAlpha").Bool()
}

// If the browser isn't capable of handling tinting with alpha this will be false.
// This property is only applicable if using tintWithPerPixel.
func (self *CanvasTinter) SetCanHandleAlpha(member bool) {
    self.Set("canHandleAlpha", member)
}

// Whether or not the Canvas BlendModes are supported, consequently the ability to tint using the multiply method.
func (self *CanvasTinter) GetCanUseMultiply() bool{
    return self.Get("canUseMultiply").Bool()
}

// Whether or not the Canvas BlendModes are supported, consequently the ability to tint using the multiply method.
func (self *CanvasTinter) SetCanUseMultiply(member bool) {
    self.Set("canUseMultiply", member)
}



// Basically this method just needs a sprite and a color and tints the sprite with the given color.
func (self *CanvasTinter) GetTintedTextureI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("getTintedTexture", args))
}

// Tint a texture using the "multiply" operation.
func (self *CanvasTinter) TintWithMultiplyI(args ...interface{}) {
    self.Call("tintWithMultiply", args)
}

// Tint a texture pixel per pixel.
func (self *CanvasTinter) TintPerPixelI(args ...interface{}) {
    self.Call("tintPerPixel", args)
}

// Checks if the browser correctly supports putImageData alpha channels.
func (self *CanvasTinter) CheckInverseAlphaI(args ...interface{}) {
    self.Call("checkInverseAlpha", args)
}

// The tinting method that will be used.
func (self *CanvasTinter) TintMethodI(args ...interface{}) {
    self.Call("tintMethod", args)
}
