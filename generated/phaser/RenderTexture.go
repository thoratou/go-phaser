// Automatic generation for Phaser.RenderTexture
// generated file RenderTexture.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// A RenderTexture is a special texture that allows any displayObject to be rendered to it. It allows you to take many complex objects and
// render them down into a single quad (on WebGL) which can then be used to texture other display objects with. A way of generating textures at run-time.
type RenderTexture struct {
    *js.Object
}


// A reference to the currently running game.
func (self *RenderTexture) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *RenderTexture) SetGame(member *Game) {
    self.Set("game", member)
}

// The key of the RenderTexture in the Cache, if stored there.
func (self *RenderTexture) GetKey() string{
    return self.Get("key").String()
}

// The key of the RenderTexture in the Cache, if stored there.
func (self *RenderTexture) SetKey(member string) {
    self.Set("key", member)
}

// Base Phaser object type.
func (self *RenderTexture) GetType() int{
    return self.Get("type").Int()
}

// Base Phaser object type.
func (self *RenderTexture) SetType(member int) {
    self.Set("type", member)
}

// The with of the render texture
func (self *RenderTexture) GetWidth() int{
    return self.Get("width").Int()
}

// The with of the render texture
func (self *RenderTexture) SetWidth(member int) {
    self.Set("width", member)
}

// The height of the render texture
func (self *RenderTexture) GetHeight() int{
    return self.Get("height").Int()
}

// The height of the render texture
func (self *RenderTexture) SetHeight(member int) {
    self.Set("height", member)
}

// The Resolution of the texture.
func (self *RenderTexture) GetResolution() int{
    return self.Get("resolution").Int()
}

// The Resolution of the texture.
func (self *RenderTexture) SetResolution(member int) {
    self.Set("resolution", member)
}

// The framing rectangle of the render texture
func (self *RenderTexture) GetFrame() *Rectangle{
    return &Rectangle{self.Get("frame")}
}

// The framing rectangle of the render texture
func (self *RenderTexture) SetFrame(member *Rectangle) {
    self.Set("frame", member)
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RenderTexture) GetCrop() *Rectangle{
    return &Rectangle{self.Get("crop")}
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RenderTexture) SetCrop(member *Rectangle) {
    self.Set("crop", member)
}

// The base texture object that this texture uses
func (self *RenderTexture) GetBaseTexture() *BaseTexture{
    return &BaseTexture{self.Get("baseTexture")}
}

// The base texture object that this texture uses
func (self *RenderTexture) SetBaseTexture(member *BaseTexture) {
    self.Set("baseTexture", member)
}

// The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RenderTexture) GetRenderer() interface{}{
    return self.Get("renderer")
}

// The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RenderTexture) SetRenderer(member interface{}) {
    self.Set("renderer", member)
}

// 
func (self *RenderTexture) GetValid() bool{
    return self.Get("valid").Bool()
}

// 
func (self *RenderTexture) SetValid(member bool) {
    self.Set("valid", member)
}

// Does this Texture have any frame data assigned to it?
func (self *RenderTexture) GetNoFrame() bool{
    return self.Get("noFrame").Bool()
}

// Does this Texture have any frame data assigned to it?
func (self *RenderTexture) SetNoFrame(member bool) {
    self.Set("noFrame", member)
}

// The texture trim data.
func (self *RenderTexture) GetTrim() *Rectangle{
    return &Rectangle{self.Get("trim")}
}

// The texture trim data.
func (self *RenderTexture) SetTrim(member *Rectangle) {
    self.Set("trim", member)
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RenderTexture) GetIsTiling() bool{
    return self.Get("isTiling").Bool()
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RenderTexture) SetIsTiling(member bool) {
    self.Set("isTiling", member)
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RenderTexture) GetRequiresUpdate() bool{
    return self.Get("requiresUpdate").Bool()
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RenderTexture) SetRequiresUpdate(member bool) {
    self.Set("requiresUpdate", member)
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *RenderTexture) GetRequiresReTint() bool{
    return self.Get("requiresReTint").Bool()
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *RenderTexture) SetRequiresReTint(member bool) {
    self.Set("requiresReTint", member)
}



// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RenderTexture) RenderXYI(args ...interface{}) {
    self.Call("renderXY", args)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RenderTexture) RenderRawXYI(args ...interface{}) {
    self.Call("renderRawXY", args)
}

// This function will draw the display object to the RenderTexture.
// 
// In versions of Phaser prior to 2.4.0 the second parameter was a Phaser.Point object. 
// This is now a Matrix allowing you much more control over how the Display Object is rendered.
// If you need to replicate the earlier behavior please use Phaser.RenderTexture.renderXY instead.
// 
// If you wish for the displayObject to be rendered taking its current scale, rotation and translation into account then either
// pass `null`, leave it undefined or pass `displayObject.worldTransform` as the matrix value.
func (self *RenderTexture) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// Resizes the RenderTexture.
func (self *RenderTexture) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// Clears the RenderTexture.
func (self *RenderTexture) ClearI(args ...interface{}) {
    self.Call("clear", args)
}

// This function will draw the display object to the texture.
func (self *RenderTexture) RenderWebGLI(args ...interface{}) {
    self.Call("renderWebGL", args)
}

// This function will draw the display object to the texture.
func (self *RenderTexture) RenderCanvasI(args ...interface{}) {
    self.Call("renderCanvas", args)
}

// Will return a HTML Image of the texture
func (self *RenderTexture) GetImageI(args ...interface{}) *Image{
    return &Image{self.Call("getImage", args)}
}

// Will return a base64 encoded string of this texture. It works by calling RenderTexture.getCanvas and then running toDataURL on that.
func (self *RenderTexture) GetBase64I(args ...interface{}) string{
    return self.Call("getBase64", args).String()
}

// Creates a Canvas element, renders this RenderTexture to it and then returns it.
func (self *RenderTexture) GetCanvasI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("getCanvas", args))
}

// Called when the base texture is loaded
func (self *RenderTexture) OnBaseTextureLoadedI(args ...interface{}) {
    self.Call("onBaseTextureLoaded", args)
}

// Destroys this texture
func (self *RenderTexture) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Specifies the region of the baseTexture that this texture will use.
func (self *RenderTexture) SetFrameI(args ...interface{}) {
    self.Call("setFrame", args)
}

// Updates the internal WebGL UV cache.
func (self *RenderTexture) _updateUvsI(args ...interface{}) {
    self.Call("_updateUvs", args)
}
