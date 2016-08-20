// Automatic generation for PIXI.BaseTexture
// generated file BaseTexture.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A texture stores the information that represents an image. All textures have a base texture.
type BaseTexture struct {
    *js.Object
}


// The Resolution of the texture.
func (self *BaseTexture) GetResolution() float64{
    return self.Get("resolution").Float()
}

// The Resolution of the texture.
func (self *BaseTexture) SetResolution(member float64) {
    self.Set("resolution", member)
}

// [read-only] The width of the base texture set when the image has loaded
func (self *BaseTexture) GetWidth() float64{
    return self.Get("width").Float()
}

// [read-only] The width of the base texture set when the image has loaded
func (self *BaseTexture) SetWidth(member float64) {
    self.Set("width", member)
}

// [read-only] The height of the base texture set when the image has loaded
func (self *BaseTexture) GetHeight() float64{
    return self.Get("height").Float()
}

// [read-only] The height of the base texture set when the image has loaded
func (self *BaseTexture) SetHeight(member float64) {
    self.Set("height", member)
}

// The scale mode to apply when scaling this texture
func (self *BaseTexture) GetScaleMode() float64{
    return self.Get("scaleMode").Float()
}

// The scale mode to apply when scaling this texture
func (self *BaseTexture) SetScaleMode(member float64) {
    self.Set("scaleMode", member)
}

// [read-only] Set to true once the base texture has loaded
func (self *BaseTexture) GetHasLoaded() bool{
    return self.Get("hasLoaded").Bool()
}

// [read-only] Set to true once the base texture has loaded
func (self *BaseTexture) SetHasLoaded(member bool) {
    self.Set("hasLoaded", member)
}

// The image source that is used to create the texture.
func (self *BaseTexture) GetSource() Image{
    return Image{self.Get("source")}
}

// The image source that is used to create the texture.
func (self *BaseTexture) SetSource(member Image) {
    self.Set("source", member)
}

// Controls if RGB channels should be pre-multiplied by Alpha  (WebGL only)
func (self *BaseTexture) GetPremultipliedAlpha() bool{
    return self.Get("premultipliedAlpha").Bool()
}

// Controls if RGB channels should be pre-multiplied by Alpha  (WebGL only)
func (self *BaseTexture) SetPremultipliedAlpha(member bool) {
    self.Set("premultipliedAlpha", member)
}

// Set this to true if a mipmap of this texture needs to be generated. This value needs to be set before the texture is used
// Also the texture must be a power of two size to work
func (self *BaseTexture) GetMipmap() bool{
    return self.Get("mipmap").Bool()
}

// Set this to true if a mipmap of this texture needs to be generated. This value needs to be set before the texture is used
// Also the texture must be a power of two size to work
func (self *BaseTexture) SetMipmap(member bool) {
    self.Set("mipmap", member)
}

// A BaseTexture can be set to skip the rendering phase in the WebGL Sprite Batch.
// 
// You may want to do this if you have a parent Sprite with no visible texture (i.e. uses the internal `__default` texture)
// that has children that you do want to render, without causing a batch flush in the process.
func (self *BaseTexture) GetSkipRender() bool{
    return self.Get("skipRender").Bool()
}

// A BaseTexture can be set to skip the rendering phase in the WebGL Sprite Batch.
// 
// You may want to do this if you have a parent Sprite with no visible texture (i.e. uses the internal `__default` texture)
// that has children that you do want to render, without causing a batch flush in the process.
func (self *BaseTexture) SetSkipRender(member bool) {
    self.Set("skipRender", member)
}



// Forces this BaseTexture to be set as loaded, with the given width and height.
// Then calls BaseTexture.dirty.
// Important for when you don't want to modify the source object by forcing in `complete` or dimension properties it may not have.
func (self *BaseTexture) ForceLoadedI(args ...interface{}) {
    self.Call("forceLoaded", args)
}

// Destroys this base texture
func (self *BaseTexture) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Changes the source image of the texture
func (self *BaseTexture) UpdateSourceImageI(args ...interface{}) {
    self.Call("updateSourceImage", args)
}

// Sets all glTextures to be dirty.
func (self *BaseTexture) DirtyI(args ...interface{}) {
    self.Call("dirty", args)
}

// Removes the base texture from the GPU, useful for managing resources on the GPU.
// Atexture is still 100% usable and will simply be reuploaded if there is a sprite on screen that is using it.
func (self *BaseTexture) UnloadFromGPUI(args ...interface{}) {
    self.Call("unloadFromGPU", args)
}

// Helper function that creates a base texture from the given canvas element.
func (self *BaseTexture) FromCanvasI(args ...interface{}) BaseTexture{
    return BaseTexture{self.Call("fromCanvas", args)}
}
