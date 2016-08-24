// Package phaser Automatic generation for PIXI.Texture
// generated file Texture.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Texture A texture stores the information that represents an image or part of an image. It cannot be added
// to the display list directly. Instead use it as the texture for a PIXI.Sprite. If no frame is provided then the whole image is used.
type Texture struct {
    *js.Object
}

// NewTexture A texture stores the information that represents an image or part of an image. It cannot be added
// to the display list directly. Instead use it as the texture for a PIXI.Sprite. If no frame is provided then the whole image is used.
func NewTexture(baseTexture *BaseTexture, frame *Rectangle) *Texture {
    return &Texture{js.Global.Get("PIXI").Get("Texture").New(baseTexture, frame)}
}
// NewTexture1O A texture stores the information that represents an image or part of an image. It cannot be added
// to the display list directly. Instead use it as the texture for a PIXI.Sprite. If no frame is provided then the whole image is used.
func NewTexture1O(baseTexture *BaseTexture, frame *Rectangle, crop *Rectangle) *Texture {
    return &Texture{js.Global.Get("PIXI").Get("Texture").New(baseTexture, frame, crop)}
}
// NewTexture2O A texture stores the information that represents an image or part of an image. It cannot be added
// to the display list directly. Instead use it as the texture for a PIXI.Sprite. If no frame is provided then the whole image is used.
func NewTexture2O(baseTexture *BaseTexture, frame *Rectangle, crop *Rectangle, trim *Rectangle) *Texture {
    return &Texture{js.Global.Get("PIXI").Get("Texture").New(baseTexture, frame, crop, trim)}
}
// NewTextureI A texture stores the information that represents an image or part of an image. It cannot be added
// to the display list directly. Instead use it as the texture for a PIXI.Sprite. If no frame is provided then the whole image is used.
func NewTextureI(args ...interface{}) *Texture {
    return &Texture{js.Global.Get("PIXI").Get("Texture").New(args)}
}



// NoFrame Does this Texture have any frame data assigned to it?
func (self *Texture) NoFrame() bool{
    return self.Object.Get("noFrame").Bool()
}

// SetNoFrameA Does this Texture have any frame data assigned to it?
func (self *Texture) SetNoFrameA(member bool) {
    self.Object.Set("noFrame", member)
}

// BaseTexture The base texture that this texture uses.
func (self *Texture) BaseTexture() *BaseTexture{
    return &BaseTexture{self.Object.Get("baseTexture")}
}

// SetBaseTextureA The base texture that this texture uses.
func (self *Texture) SetBaseTextureA(member *BaseTexture) {
    self.Object.Set("baseTexture", member)
}

// Frame The frame specifies the region of the base texture that this texture uses
func (self *Texture) Frame() *Rectangle{
    return &Rectangle{self.Object.Get("frame")}
}

// SetFrameA The frame specifies the region of the base texture that this texture uses
func (self *Texture) SetFrameA(member *Rectangle) {
    self.Object.Set("frame", member)
}

// Trim The texture trim data.
func (self *Texture) Trim() *Rectangle{
    return &Rectangle{self.Object.Get("trim")}
}

// SetTrimA The texture trim data.
func (self *Texture) SetTrimA(member *Rectangle) {
    self.Object.Set("trim", member)
}

// Valid This will let the renderer know if the texture is valid. If it's not then it cannot be rendered.
func (self *Texture) Valid() bool{
    return self.Object.Get("valid").Bool()
}

// SetValidA This will let the renderer know if the texture is valid. If it's not then it cannot be rendered.
func (self *Texture) SetValidA(member bool) {
    self.Object.Set("valid", member)
}

// IsTiling Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *Texture) IsTiling() bool{
    return self.Object.Get("isTiling").Bool()
}

// SetIsTilingA Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *Texture) SetIsTilingA(member bool) {
    self.Object.Set("isTiling", member)
}

// RequiresUpdate This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *Texture) RequiresUpdate() bool{
    return self.Object.Get("requiresUpdate").Bool()
}

// SetRequiresUpdateA This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *Texture) SetRequiresUpdateA(member bool) {
    self.Object.Set("requiresUpdate", member)
}

// RequiresReTint This will let a renderer know that a tinted parent has updated its texture.
func (self *Texture) RequiresReTint() bool{
    return self.Object.Get("requiresReTint").Bool()
}

// SetRequiresReTintA This will let a renderer know that a tinted parent has updated its texture.
func (self *Texture) SetRequiresReTintA(member bool) {
    self.Object.Set("requiresReTint", member)
}

// Width The width of the Texture in pixels.
func (self *Texture) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the Texture in pixels.
func (self *Texture) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the Texture in pixels.
func (self *Texture) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the Texture in pixels.
func (self *Texture) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Crop This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *Texture) Crop() *Rectangle{
    return &Rectangle{self.Object.Get("crop")}
}

// SetCropA This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *Texture) SetCropA(member *Rectangle) {
    self.Object.Set("crop", member)
}


// OnBaseTextureLoaded Called when the base texture is loaded
func (self *Texture) OnBaseTextureLoaded() {
    self.Object.Call("onBaseTextureLoaded")
}

// OnBaseTextureLoadedI Called when the base texture is loaded
func (self *Texture) OnBaseTextureLoadedI(args ...interface{}) {
    self.Object.Call("onBaseTextureLoaded", args)
}

// Destroy Destroys this texture
func (self *Texture) Destroy(destroyBase bool) {
    self.Object.Call("destroy", destroyBase)
}

// DestroyI Destroys this texture
func (self *Texture) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// SetFrame Specifies the region of the baseTexture that this texture will use.
func (self *Texture) SetFrame(frame *Rectangle) {
    self.Object.Call("setFrame", frame)
}

// SetFrameI Specifies the region of the baseTexture that this texture will use.
func (self *Texture) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// _updateUvs Updates the internal WebGL UV cache.
func (self *Texture) _updateUvs() {
    self.Object.Call("_updateUvs")
}

// _updateUvsI Updates the internal WebGL UV cache.
func (self *Texture) _updateUvsI(args ...interface{}) {
    self.Object.Call("_updateUvs", args)
}

// FromCanvas Helper function that creates a new a Texture based on the given canvas element.
func (self *Texture) FromCanvas(canvas *Canvas, scaleMode int) *Texture{
    return &Texture{self.Object.Call("fromCanvas", canvas, scaleMode)}
}

// FromCanvasI Helper function that creates a new a Texture based on the given canvas element.
func (self *Texture) FromCanvasI(args ...interface{}) *Texture{
    return &Texture{self.Object.Call("fromCanvas", args)}
}

