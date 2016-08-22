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
func (self *RenderTexture) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *RenderTexture) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The key of the RenderTexture in the Cache, if stored there.
func (self *RenderTexture) GetKeyA() string{
    return self.Object.Get("key").String()
}

// The key of the RenderTexture in the Cache, if stored there.
func (self *RenderTexture) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// Base Phaser object type.
func (self *RenderTexture) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// Base Phaser object type.
func (self *RenderTexture) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The with of the render texture
func (self *RenderTexture) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The with of the render texture
func (self *RenderTexture) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the render texture
func (self *RenderTexture) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the render texture
func (self *RenderTexture) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The Resolution of the texture.
func (self *RenderTexture) GetResolutionA() int{
    return self.Object.Get("resolution").Int()
}

// The Resolution of the texture.
func (self *RenderTexture) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// The framing rectangle of the render texture
func (self *RenderTexture) GetFrameA() *Rectangle{
    return &Rectangle{self.Object.Get("frame")}
}

// The framing rectangle of the render texture
func (self *RenderTexture) SetFrameA(member *Rectangle) {
    self.Object.Set("frame", member)
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RenderTexture) GetCropA() *Rectangle{
    return &Rectangle{self.Object.Get("crop")}
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RenderTexture) SetCropA(member *Rectangle) {
    self.Object.Set("crop", member)
}

// The base texture object that this texture uses
func (self *RenderTexture) GetBaseTextureA() *BaseTexture{
    return &BaseTexture{self.Object.Get("baseTexture")}
}

// The base texture object that this texture uses
func (self *RenderTexture) SetBaseTextureA(member *BaseTexture) {
    self.Object.Set("baseTexture", member)
}

// The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RenderTexture) GetRendererA() interface{}{
    return self.Object.Get("renderer")
}

// The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RenderTexture) SetRendererA(member interface{}) {
    self.Object.Set("renderer", member)
}

// 
func (self *RenderTexture) GetValidA() bool{
    return self.Object.Get("valid").Bool()
}

// 
func (self *RenderTexture) SetValidA(member bool) {
    self.Object.Set("valid", member)
}

// Does this Texture have any frame data assigned to it?
func (self *RenderTexture) GetNoFrameA() bool{
    return self.Object.Get("noFrame").Bool()
}

// Does this Texture have any frame data assigned to it?
func (self *RenderTexture) SetNoFrameA(member bool) {
    self.Object.Set("noFrame", member)
}

// The texture trim data.
func (self *RenderTexture) GetTrimA() *Rectangle{
    return &Rectangle{self.Object.Get("trim")}
}

// The texture trim data.
func (self *RenderTexture) SetTrimA(member *Rectangle) {
    self.Object.Set("trim", member)
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RenderTexture) GetIsTilingA() bool{
    return self.Object.Get("isTiling").Bool()
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RenderTexture) SetIsTilingA(member bool) {
    self.Object.Set("isTiling", member)
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RenderTexture) GetRequiresUpdateA() bool{
    return self.Object.Get("requiresUpdate").Bool()
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RenderTexture) SetRequiresUpdateA(member bool) {
    self.Object.Set("requiresUpdate", member)
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *RenderTexture) GetRequiresReTintA() bool{
    return self.Object.Get("requiresReTint").Bool()
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *RenderTexture) SetRequiresReTintA(member bool) {
    self.Object.Set("requiresReTint", member)
}



// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RenderTexture) RenderXY(displayObject interface{}, x int, y int) {
    self.Object.Call("renderXY", displayObject, x, y)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RenderTexture) RenderXY1O(displayObject interface{}, x int, y int, clear bool) {
    self.Object.Call("renderXY", displayObject, x, y, clear)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RenderTexture) RenderXYI(args ...interface{}) {
    self.Object.Call("renderXY", args)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RenderTexture) RenderRawXY(displayObject interface{}, x int, y int) {
    self.Object.Call("renderRawXY", displayObject, x, y)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RenderTexture) RenderRawXY1O(displayObject interface{}, x int, y int, clear bool) {
    self.Object.Call("renderRawXY", displayObject, x, y, clear)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RenderTexture) RenderRawXYI(args ...interface{}) {
    self.Object.Call("renderRawXY", args)
}

// This function will draw the display object to the RenderTexture.
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

// This function will draw the display object to the RenderTexture.
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

// This function will draw the display object to the RenderTexture.
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

// This function will draw the display object to the RenderTexture.
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

// Resizes the RenderTexture.
func (self *RenderTexture) Resize(width int, height int, updateBase bool) {
    self.Object.Call("resize", width, height, updateBase)
}

// Resizes the RenderTexture.
func (self *RenderTexture) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Clears the RenderTexture.
func (self *RenderTexture) Clear() {
    self.Object.Call("clear")
}

// Clears the RenderTexture.
func (self *RenderTexture) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// This function will draw the display object to the texture.
func (self *RenderTexture) RenderWebGL(displayObject *DisplayObject) {
    self.Object.Call("renderWebGL", displayObject)
}

// This function will draw the display object to the texture.
func (self *RenderTexture) RenderWebGL1O(displayObject *DisplayObject, matrix *Matrix) {
    self.Object.Call("renderWebGL", displayObject, matrix)
}

// This function will draw the display object to the texture.
func (self *RenderTexture) RenderWebGL2O(displayObject *DisplayObject, matrix *Matrix, clear bool) {
    self.Object.Call("renderWebGL", displayObject, matrix, clear)
}

// This function will draw the display object to the texture.
func (self *RenderTexture) RenderWebGLI(args ...interface{}) {
    self.Object.Call("renderWebGL", args)
}

// This function will draw the display object to the texture.
func (self *RenderTexture) RenderCanvas(displayObject *DisplayObject) {
    self.Object.Call("renderCanvas", displayObject)
}

// This function will draw the display object to the texture.
func (self *RenderTexture) RenderCanvas1O(displayObject *DisplayObject, matrix *Matrix) {
    self.Object.Call("renderCanvas", displayObject, matrix)
}

// This function will draw the display object to the texture.
func (self *RenderTexture) RenderCanvas2O(displayObject *DisplayObject, matrix *Matrix, clear bool) {
    self.Object.Call("renderCanvas", displayObject, matrix, clear)
}

// This function will draw the display object to the texture.
func (self *RenderTexture) RenderCanvasI(args ...interface{}) {
    self.Object.Call("renderCanvas", args)
}

// Will return a HTML Image of the texture
func (self *RenderTexture) GetImage() *Image{
    return &Image{self.Object.Call("getImage")}
}

// Will return a HTML Image of the texture
func (self *RenderTexture) GetImageI(args ...interface{}) *Image{
    return &Image{self.Object.Call("getImage", args)}
}

// Will return a base64 encoded string of this texture. It works by calling RenderTexture.getCanvas and then running toDataURL on that.
func (self *RenderTexture) GetBase64() string{
    return self.Object.Call("getBase64").String()
}

// Will return a base64 encoded string of this texture. It works by calling RenderTexture.getCanvas and then running toDataURL on that.
func (self *RenderTexture) GetBase64I(args ...interface{}) string{
    return self.Object.Call("getBase64", args).String()
}

// Creates a Canvas element, renders this RenderTexture to it and then returns it.
func (self *RenderTexture) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("getCanvas"))
}

// Creates a Canvas element, renders this RenderTexture to it and then returns it.
func (self *RenderTexture) GetCanvasI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("getCanvas", args))
}

// Called when the base texture is loaded
func (self *RenderTexture) OnBaseTextureLoaded() {
    self.Object.Call("onBaseTextureLoaded")
}

// Called when the base texture is loaded
func (self *RenderTexture) OnBaseTextureLoadedI(args ...interface{}) {
    self.Object.Call("onBaseTextureLoaded", args)
}

// Destroys this texture
func (self *RenderTexture) Destroy(destroyBase bool) {
    self.Object.Call("destroy", destroyBase)
}

// Destroys this texture
func (self *RenderTexture) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Specifies the region of the baseTexture that this texture will use.
func (self *RenderTexture) SetFrame(frame *Rectangle) {
    self.Object.Call("setFrame", frame)
}

// Specifies the region of the baseTexture that this texture will use.
func (self *RenderTexture) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// Updates the internal WebGL UV cache.
func (self *RenderTexture) _updateUvs() {
    self.Object.Call("_updateUvs")
}

// Updates the internal WebGL UV cache.
func (self *RenderTexture) _updateUvsI(args ...interface{}) {
    self.Object.Call("_updateUvs", args)
}
