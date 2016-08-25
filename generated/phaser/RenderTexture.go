// Package phaser Automatic generation for Phaser.RenderTexture
// generated file RenderTexture.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// RenderTexture A RenderTexture is a special texture that allows any displayObject to be rendered to it. It allows you to take many complex objects and
// render them down into a single quad (on WebGL) which can then be used to texture other display objects with. A way of generating textures at run-time.
type RenderTexture struct {
    *js.Object
}

// NewRenderTexture A RenderTexture is a special texture that allows any displayObject to be rendered to it. It allows you to take many complex objects and
// render them down into a single quad (on WebGL) which can then be used to texture other display objects with. A way of generating textures at run-time.
func NewRenderTexture(game *Game) *RenderTexture {
    return &RenderTexture{js.Global.Get("Phaser").Get("RenderTexture").New(game)}
}
// NewRenderTexture1O A RenderTexture is a special texture that allows any displayObject to be rendered to it. It allows you to take many complex objects and
// render them down into a single quad (on WebGL) which can then be used to texture other display objects with. A way of generating textures at run-time.
func NewRenderTexture1O(game *Game, width int) *RenderTexture {
    return &RenderTexture{js.Global.Get("Phaser").Get("RenderTexture").New(game, width)}
}
// NewRenderTexture2O A RenderTexture is a special texture that allows any displayObject to be rendered to it. It allows you to take many complex objects and
// render them down into a single quad (on WebGL) which can then be used to texture other display objects with. A way of generating textures at run-time.
func NewRenderTexture2O(game *Game, width int, height int) *RenderTexture {
    return &RenderTexture{js.Global.Get("Phaser").Get("RenderTexture").New(game, width, height)}
}
// NewRenderTexture3O A RenderTexture is a special texture that allows any displayObject to be rendered to it. It allows you to take many complex objects and
// render them down into a single quad (on WebGL) which can then be used to texture other display objects with. A way of generating textures at run-time.
func NewRenderTexture3O(game *Game, width int, height int, key string) *RenderTexture {
    return &RenderTexture{js.Global.Get("Phaser").Get("RenderTexture").New(game, width, height, key)}
}
// NewRenderTexture4O A RenderTexture is a special texture that allows any displayObject to be rendered to it. It allows you to take many complex objects and
// render them down into a single quad (on WebGL) which can then be used to texture other display objects with. A way of generating textures at run-time.
func NewRenderTexture4O(game *Game, width int, height int, key string, scaleMode int) *RenderTexture {
    return &RenderTexture{js.Global.Get("Phaser").Get("RenderTexture").New(game, width, height, key, scaleMode)}
}
// NewRenderTexture5O A RenderTexture is a special texture that allows any displayObject to be rendered to it. It allows you to take many complex objects and
// render them down into a single quad (on WebGL) which can then be used to texture other display objects with. A way of generating textures at run-time.
func NewRenderTexture5O(game *Game, width int, height int, key string, scaleMode int, resolution int) *RenderTexture {
    return &RenderTexture{js.Global.Get("Phaser").Get("RenderTexture").New(game, width, height, key, scaleMode, resolution)}
}
// NewRenderTextureI A RenderTexture is a special texture that allows any displayObject to be rendered to it. It allows you to take many complex objects and
// render them down into a single quad (on WebGL) which can then be used to texture other display objects with. A way of generating textures at run-time.
func NewRenderTextureI(args ...interface{}) *RenderTexture {
    return &RenderTexture{js.Global.Get("Phaser").Get("RenderTexture").New(args)}
}



// RenderTexture Binding conversion method to RenderTexture point 
func ToRenderTexture(jsStruct interface{}) *RenderTexture {
    if object, ok := jsStruct.(*js.Object); ok {
		return &RenderTexture{Object: object}
	}
	return nil
}



// Game A reference to the currently running game.
func (self *RenderTexture) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *RenderTexture) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Key The key of the RenderTexture in the Cache, if stored there.
func (self *RenderTexture) Key() string{
    return self.Object.Get("key").String()
}

// SetKeyA The key of the RenderTexture in the Cache, if stored there.
func (self *RenderTexture) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// Type Base Phaser object type.
func (self *RenderTexture) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA Base Phaser object type.
func (self *RenderTexture) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Width The with of the render texture
func (self *RenderTexture) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The with of the render texture
func (self *RenderTexture) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the render texture
func (self *RenderTexture) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the render texture
func (self *RenderTexture) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Resolution The Resolution of the texture.
func (self *RenderTexture) Resolution() int{
    return self.Object.Get("resolution").Int()
}

// SetResolutionA The Resolution of the texture.
func (self *RenderTexture) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// Frame The framing rectangle of the render texture
func (self *RenderTexture) Frame() *Rectangle{
    return &Rectangle{self.Object.Get("frame")}
}

// SetFrameA The framing rectangle of the render texture
func (self *RenderTexture) SetFrameA(member *Rectangle) {
    self.Object.Set("frame", member)
}

// Crop This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RenderTexture) Crop() *Rectangle{
    return &Rectangle{self.Object.Get("crop")}
}

// SetCropA This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RenderTexture) SetCropA(member *Rectangle) {
    self.Object.Set("crop", member)
}

// BaseTexture The base texture object that this texture uses
func (self *RenderTexture) BaseTexture() *BaseTexture{
    return &BaseTexture{self.Object.Get("baseTexture")}
}

// SetBaseTextureA The base texture object that this texture uses
func (self *RenderTexture) SetBaseTextureA(member *BaseTexture) {
    self.Object.Set("baseTexture", member)
}

// Renderer The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RenderTexture) Renderer() interface{}{
    return self.Object.Get("renderer")
}

// SetRendererA The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RenderTexture) SetRendererA(member interface{}) {
    self.Object.Set("renderer", member)
}

// Valid empty description
func (self *RenderTexture) Valid() bool{
    return self.Object.Get("valid").Bool()
}

// SetValidA empty description
func (self *RenderTexture) SetValidA(member bool) {
    self.Object.Set("valid", member)
}

// NoFrame Does this Texture have any frame data assigned to it?
func (self *RenderTexture) NoFrame() bool{
    return self.Object.Get("noFrame").Bool()
}

// SetNoFrameA Does this Texture have any frame data assigned to it?
func (self *RenderTexture) SetNoFrameA(member bool) {
    self.Object.Set("noFrame", member)
}

// Trim The texture trim data.
func (self *RenderTexture) Trim() *Rectangle{
    return &Rectangle{self.Object.Get("trim")}
}

// SetTrimA The texture trim data.
func (self *RenderTexture) SetTrimA(member *Rectangle) {
    self.Object.Set("trim", member)
}

// IsTiling Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RenderTexture) IsTiling() bool{
    return self.Object.Get("isTiling").Bool()
}

// SetIsTilingA Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RenderTexture) SetIsTilingA(member bool) {
    self.Object.Set("isTiling", member)
}

// RequiresUpdate This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RenderTexture) RequiresUpdate() bool{
    return self.Object.Get("requiresUpdate").Bool()
}

// SetRequiresUpdateA This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RenderTexture) SetRequiresUpdateA(member bool) {
    self.Object.Set("requiresUpdate", member)
}

// RequiresReTint This will let a renderer know that a tinted parent has updated its texture.
func (self *RenderTexture) RequiresReTint() bool{
    return self.Object.Get("requiresReTint").Bool()
}

// SetRequiresReTintA This will let a renderer know that a tinted parent has updated its texture.
func (self *RenderTexture) SetRequiresReTintA(member bool) {
    self.Object.Set("requiresReTint", member)
}


// RenderXY This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RenderTexture) RenderXY(displayObject interface{}, x int, y int) {
    self.Object.Call("renderXY", displayObject, x, y)
}

// RenderXY1O This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RenderTexture) RenderXY1O(displayObject interface{}, x int, y int, clear bool) {
    self.Object.Call("renderXY", displayObject, x, y, clear)
}

// RenderXYI This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RenderTexture) RenderXYI(args ...interface{}) {
    self.Object.Call("renderXY", args)
}

// RenderRawXY This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RenderTexture) RenderRawXY(displayObject interface{}, x int, y int) {
    self.Object.Call("renderRawXY", displayObject, x, y)
}

// RenderRawXY1O This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RenderTexture) RenderRawXY1O(displayObject interface{}, x int, y int, clear bool) {
    self.Object.Call("renderRawXY", displayObject, x, y, clear)
}

// RenderRawXYI This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RenderTexture) RenderRawXYI(args ...interface{}) {
    self.Object.Call("renderRawXY", args)
}

// Render This function will draw the display object to the RenderTexture.
// 
// In versions of Phaser prior to 2.4.0 the second parameter was a Phaser.Point object. 
// This is now a Matrix allowing you much more control over how the Display Object is rendered.
// If you need to replicate the earlier behavior please use Phaser.RenderTexture.renderXY instead.
// 
// If you wish for the displayObject to be rendered taking its current scale, rotation and translation into account then either
// pass `null`, leave it undefined or pass `displayObject.worldTransform` as the matrix value.
func (self *RenderTexture) Render(displayObject interface{}) {
    self.Object.Call("render", displayObject)
}

// Render1O This function will draw the display object to the RenderTexture.
// 
// In versions of Phaser prior to 2.4.0 the second parameter was a Phaser.Point object. 
// This is now a Matrix allowing you much more control over how the Display Object is rendered.
// If you need to replicate the earlier behavior please use Phaser.RenderTexture.renderXY instead.
// 
// If you wish for the displayObject to be rendered taking its current scale, rotation and translation into account then either
// pass `null`, leave it undefined or pass `displayObject.worldTransform` as the matrix value.
func (self *RenderTexture) Render1O(displayObject interface{}, matrix *Matrix) {
    self.Object.Call("render", displayObject, matrix)
}

// Render2O This function will draw the display object to the RenderTexture.
// 
// In versions of Phaser prior to 2.4.0 the second parameter was a Phaser.Point object. 
// This is now a Matrix allowing you much more control over how the Display Object is rendered.
// If you need to replicate the earlier behavior please use Phaser.RenderTexture.renderXY instead.
// 
// If you wish for the displayObject to be rendered taking its current scale, rotation and translation into account then either
// pass `null`, leave it undefined or pass `displayObject.worldTransform` as the matrix value.
func (self *RenderTexture) Render2O(displayObject interface{}, matrix *Matrix, clear bool) {
    self.Object.Call("render", displayObject, matrix, clear)
}

// RenderI This function will draw the display object to the RenderTexture.
// 
// In versions of Phaser prior to 2.4.0 the second parameter was a Phaser.Point object. 
// This is now a Matrix allowing you much more control over how the Display Object is rendered.
// If you need to replicate the earlier behavior please use Phaser.RenderTexture.renderXY instead.
// 
// If you wish for the displayObject to be rendered taking its current scale, rotation and translation into account then either
// pass `null`, leave it undefined or pass `displayObject.worldTransform` as the matrix value.
func (self *RenderTexture) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// Resize Resizes the RenderTexture.
func (self *RenderTexture) Resize(width int, height int, updateBase bool) {
    self.Object.Call("resize", width, height, updateBase)
}

// ResizeI Resizes the RenderTexture.
func (self *RenderTexture) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Clear Clears the RenderTexture.
func (self *RenderTexture) Clear() {
    self.Object.Call("clear")
}

// ClearI Clears the RenderTexture.
func (self *RenderTexture) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// RenderWebGL This function will draw the display object to the texture.
func (self *RenderTexture) RenderWebGL(displayObject *DisplayObject) {
    self.Object.Call("renderWebGL", displayObject)
}

// RenderWebGL1O This function will draw the display object to the texture.
func (self *RenderTexture) RenderWebGL1O(displayObject *DisplayObject, matrix *Matrix) {
    self.Object.Call("renderWebGL", displayObject, matrix)
}

// RenderWebGL2O This function will draw the display object to the texture.
func (self *RenderTexture) RenderWebGL2O(displayObject *DisplayObject, matrix *Matrix, clear bool) {
    self.Object.Call("renderWebGL", displayObject, matrix, clear)
}

// RenderWebGLI This function will draw the display object to the texture.
func (self *RenderTexture) RenderWebGLI(args ...interface{}) {
    self.Object.Call("renderWebGL", args)
}

// RenderCanvas This function will draw the display object to the texture.
func (self *RenderTexture) RenderCanvas(displayObject *DisplayObject) {
    self.Object.Call("renderCanvas", displayObject)
}

// RenderCanvas1O This function will draw the display object to the texture.
func (self *RenderTexture) RenderCanvas1O(displayObject *DisplayObject, matrix *Matrix) {
    self.Object.Call("renderCanvas", displayObject, matrix)
}

// RenderCanvas2O This function will draw the display object to the texture.
func (self *RenderTexture) RenderCanvas2O(displayObject *DisplayObject, matrix *Matrix, clear bool) {
    self.Object.Call("renderCanvas", displayObject, matrix, clear)
}

// RenderCanvasI This function will draw the display object to the texture.
func (self *RenderTexture) RenderCanvasI(args ...interface{}) {
    self.Object.Call("renderCanvas", args)
}

// GetImage Will return a HTML Image of the texture
func (self *RenderTexture) GetImage() *Image{
    return &Image{self.Object.Call("getImage")}
}

// GetImageI Will return a HTML Image of the texture
func (self *RenderTexture) GetImageI(args ...interface{}) *Image{
    return &Image{self.Object.Call("getImage", args)}
}

// GetBase64 Will return a base64 encoded string of this texture. It works by calling RenderTexture.getCanvas and then running toDataURL on that.
func (self *RenderTexture) GetBase64() string{
    return self.Object.Call("getBase64").String()
}

// GetBase64I Will return a base64 encoded string of this texture. It works by calling RenderTexture.getCanvas and then running toDataURL on that.
func (self *RenderTexture) GetBase64I(args ...interface{}) string{
    return self.Object.Call("getBase64", args).String()
}

// GetCanvas Creates a Canvas element, renders this RenderTexture to it and then returns it.
func (self *RenderTexture) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("getCanvas"))
}

// GetCanvasI Creates a Canvas element, renders this RenderTexture to it and then returns it.
func (self *RenderTexture) GetCanvasI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("getCanvas", args))
}

// OnBaseTextureLoaded Called when the base texture is loaded
func (self *RenderTexture) OnBaseTextureLoaded() {
    self.Object.Call("onBaseTextureLoaded")
}

// OnBaseTextureLoadedI Called when the base texture is loaded
func (self *RenderTexture) OnBaseTextureLoadedI(args ...interface{}) {
    self.Object.Call("onBaseTextureLoaded", args)
}

// Destroy Destroys this texture
func (self *RenderTexture) Destroy(destroyBase bool) {
    self.Object.Call("destroy", destroyBase)
}

// DestroyI Destroys this texture
func (self *RenderTexture) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// SetFrame Specifies the region of the baseTexture that this texture will use.
func (self *RenderTexture) SetFrame(frame *Rectangle) {
    self.Object.Call("setFrame", frame)
}

// SetFrameI Specifies the region of the baseTexture that this texture will use.
func (self *RenderTexture) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// _updateUvs Updates the internal WebGL UV cache.
func (self *RenderTexture) _updateUvs() {
    self.Object.Call("_updateUvs")
}

// _updateUvsI Updates the internal WebGL UV cache.
func (self *RenderTexture) _updateUvsI(args ...interface{}) {
    self.Object.Call("_updateUvs", args)
}

