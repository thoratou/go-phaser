// Automatic generation for PIXI.Texture
// generated file Texture.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A texture stores the information that represents an image or part of an image. It cannot be added
// to the display list directly. Instead use it as the texture for a PIXI.Sprite. If no frame is provided then the whole image is used.
type Texture struct {
    *js.Object
}


// Does this Texture have any frame data assigned to it?
func (self *Texture) GetNoFrame() bool{
    return self.Get("noFrame").Bool()
}

// Does this Texture have any frame data assigned to it?
func (self *Texture) SetNoFrame(member bool) {
    self.Set("noFrame", member)
}

// The base texture that this texture uses.
func (self *Texture) GetBaseTexture() *BaseTexture{
    return &BaseTexture{self.Get("baseTexture")}
}

// The base texture that this texture uses.
func (self *Texture) SetBaseTexture(member *BaseTexture) {
    self.Set("baseTexture", member)
}

// The frame specifies the region of the base texture that this texture uses
func (self *Texture) GetFrame() *Rectangle{
    return &Rectangle{self.Get("frame")}
}

// The frame specifies the region of the base texture that this texture uses
func (self *Texture) SetFrame(member *Rectangle) {
    self.Set("frame", member)
}

// The texture trim data.
func (self *Texture) GetTrim() *Rectangle{
    return &Rectangle{self.Get("trim")}
}

// The texture trim data.
func (self *Texture) SetTrim(member *Rectangle) {
    self.Set("trim", member)
}

// This will let the renderer know if the texture is valid. If it's not then it cannot be rendered.
func (self *Texture) GetValid() bool{
    return self.Get("valid").Bool()
}

// This will let the renderer know if the texture is valid. If it's not then it cannot be rendered.
func (self *Texture) SetValid(member bool) {
    self.Set("valid", member)
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *Texture) GetIsTiling() bool{
    return self.Get("isTiling").Bool()
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *Texture) SetIsTiling(member bool) {
    self.Set("isTiling", member)
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *Texture) GetRequiresUpdate() bool{
    return self.Get("requiresUpdate").Bool()
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *Texture) SetRequiresUpdate(member bool) {
    self.Set("requiresUpdate", member)
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *Texture) GetRequiresReTint() bool{
    return self.Get("requiresReTint").Bool()
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *Texture) SetRequiresReTint(member bool) {
    self.Set("requiresReTint", member)
}

// The width of the Texture in pixels.
func (self *Texture) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the Texture in pixels.
func (self *Texture) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the Texture in pixels.
func (self *Texture) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the Texture in pixels.
func (self *Texture) SetHeight(member float64) {
    self.Set("height", member)
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *Texture) GetCrop() *Rectangle{
    return &Rectangle{self.Get("crop")}
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *Texture) SetCrop(member *Rectangle) {
    self.Set("crop", member)
}



// Called when the base texture is loaded
func (self *Texture) OnBaseTextureLoadedI(args ...interface{}) {
    self.Call("onBaseTextureLoaded", args)
}

// Destroys this texture
func (self *Texture) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Specifies the region of the baseTexture that this texture will use.
func (self *Texture) SetFrameI(args ...interface{}) {
    self.Call("setFrame", args)
}

// Updates the internal WebGL UV cache.
func (self *Texture) _updateUvsI(args ...interface{}) {
    self.Call("_updateUvs", args)
}

// Helper function that creates a new a Texture based on the given canvas element.
func (self *Texture) FromCanvasI(args ...interface{}) *Texture{
    return &Texture{self.Call("fromCanvas", args)}
}
