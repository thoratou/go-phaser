// Package phaser Automatic generation for PIXI.BaseTexture
// generated file BaseTexture.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// BaseTexture A texture stores the information that represents an image. All textures have a base texture.
type BaseTexture struct {
    *js.Object
}

// NewBaseTexture A texture stores the information that represents an image. All textures have a base texture.
func NewBaseTexture(source interface{}, scaleMode int) *BaseTexture {
    return &BaseTexture{js.Global.Get("PIXI").Get("BaseTexture").New(source, scaleMode)}
}
// NewBaseTextureI A texture stores the information that represents an image. All textures have a base texture.
func NewBaseTextureI(args ...interface{}) *BaseTexture {
    return &BaseTexture{js.Global.Get("PIXI").Get("BaseTexture").New(args)}
}



// Resolution The Resolution of the texture.
func (self *BaseTexture) Resolution() int{
    return self.Object.Get("resolution").Int()
}

// SetResolutionA The Resolution of the texture.
func (self *BaseTexture) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// Width [read-only] The width of the base texture set when the image has loaded
func (self *BaseTexture) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA [read-only] The width of the base texture set when the image has loaded
func (self *BaseTexture) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height [read-only] The height of the base texture set when the image has loaded
func (self *BaseTexture) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA [read-only] The height of the base texture set when the image has loaded
func (self *BaseTexture) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// ScaleMode The scale mode to apply when scaling this texture
func (self *BaseTexture) ScaleMode() int{
    return self.Object.Get("scaleMode").Int()
}

// SetScaleModeA The scale mode to apply when scaling this texture
func (self *BaseTexture) SetScaleModeA(member int) {
    self.Object.Set("scaleMode", member)
}

// HasLoaded [read-only] Set to true once the base texture has loaded
func (self *BaseTexture) HasLoaded() bool{
    return self.Object.Get("hasLoaded").Bool()
}

// SetHasLoadedA [read-only] Set to true once the base texture has loaded
func (self *BaseTexture) SetHasLoadedA(member bool) {
    self.Object.Set("hasLoaded", member)
}

// Source The image source that is used to create the texture.
func (self *BaseTexture) Source() *Image{
    return &Image{self.Object.Get("source")}
}

// SetSourceA The image source that is used to create the texture.
func (self *BaseTexture) SetSourceA(member *Image) {
    self.Object.Set("source", member)
}

// PremultipliedAlpha Controls if RGB channels should be pre-multiplied by Alpha  (WebGL only)
func (self *BaseTexture) PremultipliedAlpha() bool{
    return self.Object.Get("premultipliedAlpha").Bool()
}

// SetPremultipliedAlphaA Controls if RGB channels should be pre-multiplied by Alpha  (WebGL only)
func (self *BaseTexture) SetPremultipliedAlphaA(member bool) {
    self.Object.Set("premultipliedAlpha", member)
}

// Mipmap Set this to true if a mipmap of this texture needs to be generated. This value needs to be set before the texture is used
// Also the texture must be a power of two size to work
func (self *BaseTexture) Mipmap() bool{
    return self.Object.Get("mipmap").Bool()
}

// SetMipmapA Set this to true if a mipmap of this texture needs to be generated. This value needs to be set before the texture is used
// Also the texture must be a power of two size to work
func (self *BaseTexture) SetMipmapA(member bool) {
    self.Object.Set("mipmap", member)
}

// SkipRender A BaseTexture can be set to skip the rendering phase in the WebGL Sprite Batch.
// 
// You may want to do this if you have a parent Sprite with no visible texture (i.e. uses the internal `__default` texture)
// that has children that you do want to render, without causing a batch flush in the process.
func (self *BaseTexture) SkipRender() bool{
    return self.Object.Get("skipRender").Bool()
}

// SetSkipRenderA A BaseTexture can be set to skip the rendering phase in the WebGL Sprite Batch.
// 
// You may want to do this if you have a parent Sprite with no visible texture (i.e. uses the internal `__default` texture)
// that has children that you do want to render, without causing a batch flush in the process.
func (self *BaseTexture) SetSkipRenderA(member bool) {
    self.Object.Set("skipRender", member)
}


// ForceLoaded Forces this BaseTexture to be set as loaded, with the given width and height.
// Then calls BaseTexture.dirty.
// Important for when you don't want to modify the source object by forcing in `complete` or dimension properties it may not have.
func (self *BaseTexture) ForceLoaded(width int, height int) {
    self.Object.Call("forceLoaded", width, height)
}

// ForceLoadedI Forces this BaseTexture to be set as loaded, with the given width and height.
// Then calls BaseTexture.dirty.
// Important for when you don't want to modify the source object by forcing in `complete` or dimension properties it may not have.
func (self *BaseTexture) ForceLoadedI(args ...interface{}) {
    self.Object.Call("forceLoaded", args)
}

// Destroy Destroys this base texture
func (self *BaseTexture) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this base texture
func (self *BaseTexture) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// UpdateSourceImage Changes the source image of the texture
func (self *BaseTexture) UpdateSourceImage(newSrc string) {
    self.Object.Call("updateSourceImage", newSrc)
}

// UpdateSourceImageI Changes the source image of the texture
func (self *BaseTexture) UpdateSourceImageI(args ...interface{}) {
    self.Object.Call("updateSourceImage", args)
}

// Dirty Sets all glTextures to be dirty.
func (self *BaseTexture) Dirty() {
    self.Object.Call("dirty")
}

// DirtyI Sets all glTextures to be dirty.
func (self *BaseTexture) DirtyI(args ...interface{}) {
    self.Object.Call("dirty", args)
}

// UnloadFromGPU Removes the base texture from the GPU, useful for managing resources on the GPU.
// Atexture is still 100% usable and will simply be reuploaded if there is a sprite on screen that is using it.
func (self *BaseTexture) UnloadFromGPU() {
    self.Object.Call("unloadFromGPU")
}

// UnloadFromGPUI Removes the base texture from the GPU, useful for managing resources on the GPU.
// Atexture is still 100% usable and will simply be reuploaded if there is a sprite on screen that is using it.
func (self *BaseTexture) UnloadFromGPUI(args ...interface{}) {
    self.Object.Call("unloadFromGPU", args)
}

// FromCanvas Helper function that creates a base texture from the given canvas element.
func (self *BaseTexture) FromCanvas(canvas *Canvas, scaleMode int) *BaseTexture{
    return &BaseTexture{self.Object.Call("fromCanvas", canvas, scaleMode)}
}

// FromCanvasI Helper function that creates a base texture from the given canvas element.
func (self *BaseTexture) FromCanvasI(args ...interface{}) *BaseTexture{
    return &BaseTexture{self.Object.Call("fromCanvas", args)}
}

