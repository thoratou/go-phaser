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
func (self *Texture) GetNoFrameA() bool{
    return self.Object.Get("noFrame").Bool()
}

// Does this Texture have any frame data assigned to it?
func (self *Texture) SetNoFrameA(member bool) {
    self.Object.Set("noFrame", member)
}

// The base texture that this texture uses.
func (self *Texture) GetBaseTextureA() *BaseTexture{
    return &BaseTexture{self.Object.Get("baseTexture")}
}

// The base texture that this texture uses.
func (self *Texture) SetBaseTextureA(member *BaseTexture) {
    self.Object.Set("baseTexture", member)
}

// The frame specifies the region of the base texture that this texture uses
func (self *Texture) GetFrameA() *Rectangle{
    return &Rectangle{self.Object.Get("frame")}
}

// The frame specifies the region of the base texture that this texture uses
func (self *Texture) SetFrameA(member *Rectangle) {
    self.Object.Set("frame", member)
}

// The texture trim data.
func (self *Texture) GetTrimA() *Rectangle{
    return &Rectangle{self.Object.Get("trim")}
}

// The texture trim data.
func (self *Texture) SetTrimA(member *Rectangle) {
    self.Object.Set("trim", member)
}

// This will let the renderer know if the texture is valid. If it's not then it cannot be rendered.
func (self *Texture) GetValidA() bool{
    return self.Object.Get("valid").Bool()
}

// This will let the renderer know if the texture is valid. If it's not then it cannot be rendered.
func (self *Texture) SetValidA(member bool) {
    self.Object.Set("valid", member)
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *Texture) GetIsTilingA() bool{
    return self.Object.Get("isTiling").Bool()
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *Texture) SetIsTilingA(member bool) {
    self.Object.Set("isTiling", member)
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *Texture) GetRequiresUpdateA() bool{
    return self.Object.Get("requiresUpdate").Bool()
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *Texture) SetRequiresUpdateA(member bool) {
    self.Object.Set("requiresUpdate", member)
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *Texture) GetRequiresReTintA() bool{
    return self.Object.Get("requiresReTint").Bool()
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *Texture) SetRequiresReTintA(member bool) {
    self.Object.Set("requiresReTint", member)
}

// The width of the Texture in pixels.
func (self *Texture) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the Texture in pixels.
func (self *Texture) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the Texture in pixels.
func (self *Texture) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the Texture in pixels.
func (self *Texture) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *Texture) GetCropA() *Rectangle{
    return &Rectangle{self.Object.Get("crop")}
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *Texture) SetCropA(member *Rectangle) {
    self.Object.Set("crop", member)
}



// Called when the base texture is loaded
func (self *Texture) OnBaseTextureLoaded() {
    self.Object.Call("onBaseTextureLoaded")
}

// Called when the base texture is loaded
func (self *Texture) OnBaseTextureLoadedI(args ...interface{}) {
    self.Object.Call("onBaseTextureLoaded", args)
}

// Destroys this texture
func (self *Texture) Destroy(destroyBase bool) {
    self.Object.Call("destroy", destroyBase)
}

// Destroys this texture
func (self *Texture) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Specifies the region of the baseTexture that this texture will use.
func (self *Texture) SetFrame(frame *Rectangle) {
    self.Object.Call("setFrame", frame)
}

// Specifies the region of the baseTexture that this texture will use.
func (self *Texture) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// Updates the internal WebGL UV cache.
func (self *Texture) _updateUvs() {
    self.Object.Call("_updateUvs")
}

// Updates the internal WebGL UV cache.
func (self *Texture) _updateUvsI(args ...interface{}) {
    self.Object.Call("_updateUvs", args)
}

// Helper function that creates a new a Texture based on the given canvas element.
func (self *Texture) FromCanvas(canvas *Canvas, scaleMode int) *Texture{
    return &Texture{self.Object.Call("fromCanvas", canvas, scaleMode)}
}

// Helper function that creates a new a Texture based on the given canvas element.
func (self *Texture) FromCanvasI(args ...interface{}) *Texture{
    return &Texture{self.Object.Call("fromCanvas", args)}
}
