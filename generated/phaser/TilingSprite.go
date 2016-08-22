// Automatic generation for PIXI.TilingSprite
// generated file TilingSprite.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A tiling sprite is a fast way of rendering a tiling image
type TilingSprite struct {
    *js.Object
}


// The width of the tiling sprite
func (self *TilingSprite) GetWidth() int{
    return self.Get("width").Int()
}

// The width of the tiling sprite
func (self *TilingSprite) SetWidth(member int) {
    self.Set("width", member)
}

// The height of the tiling sprite
func (self *TilingSprite) GetHeight() int{
    return self.Get("height").Int()
}

// The height of the tiling sprite
func (self *TilingSprite) SetHeight(member int) {
    self.Set("height", member)
}

// The scaling of the image that is being tiled
func (self *TilingSprite) GetTileScale() *Point{
    return &Point{self.Get("tileScale")}
}

// The scaling of the image that is being tiled
func (self *TilingSprite) SetTileScale(member *Point) {
    self.Set("tileScale", member)
}

// A point that represents the scale of the texture object
func (self *TilingSprite) GetTileScaleOffset() *Point{
    return &Point{self.Get("tileScaleOffset")}
}

// A point that represents the scale of the texture object
func (self *TilingSprite) SetTileScaleOffset(member *Point) {
    self.Set("tileScaleOffset", member)
}

// The offset position of the image that is being tiled
func (self *TilingSprite) GetTilePosition() *Point{
    return &Point{self.Get("tilePosition")}
}

// The offset position of the image that is being tiled
func (self *TilingSprite) SetTilePosition(member *Point) {
    self.Set("tilePosition", member)
}

// Whether this sprite is renderable or not
func (self *TilingSprite) GetRenderable() bool{
    return self.Get("renderable").Bool()
}

// Whether this sprite is renderable or not
func (self *TilingSprite) SetRenderable(member bool) {
    self.Set("renderable", member)
}

// The tint applied to the sprite. This is a hex value
func (self *TilingSprite) GetTint() int{
    return self.Get("tint").Int()
}

// The tint applied to the sprite. This is a hex value
func (self *TilingSprite) SetTint(member int) {
    self.Set("tint", member)
}

// If enabled a green rectangle will be drawn behind the generated tiling texture, allowing you to visually
// debug the texture being used.
func (self *TilingSprite) GetTextureDebug() bool{
    return self.Get("textureDebug").Bool()
}

// If enabled a green rectangle will be drawn behind the generated tiling texture, allowing you to visually
// debug the texture being used.
func (self *TilingSprite) SetTextureDebug(member bool) {
    self.Set("textureDebug", member)
}

// The blend mode to be applied to the sprite
func (self *TilingSprite) GetBlendMode() int{
    return self.Get("blendMode").Int()
}

// The blend mode to be applied to the sprite
func (self *TilingSprite) SetBlendMode(member int) {
    self.Set("blendMode", member)
}

// The CanvasBuffer object that the tiled texture is drawn to.
func (self *TilingSprite) GetCanvasBuffer() *PIXICanvasBuffer{
    return &PIXICanvasBuffer{self.Get("canvasBuffer")}
}

// The CanvasBuffer object that the tiled texture is drawn to.
func (self *TilingSprite) SetCanvasBuffer(member *PIXICanvasBuffer) {
    self.Set("canvasBuffer", member)
}

// An internal Texture object that holds the tiling texture that was generated from TilingSprite.texture.
func (self *TilingSprite) GetTilingTexture() *PIXITexture{
    return &PIXITexture{self.Get("tilingTexture")}
}

// An internal Texture object that holds the tiling texture that was generated from TilingSprite.texture.
func (self *TilingSprite) SetTilingTexture(member *PIXITexture) {
    self.Set("tilingTexture", member)
}

// The Context fill pattern that is used to draw the TilingSprite in Canvas mode only (will be null in WebGL).
func (self *TilingSprite) GetTilePattern() *PIXITexture{
    return &PIXITexture{self.Get("tilePattern")}
}

// The Context fill pattern that is used to draw the TilingSprite in Canvas mode only (will be null in WebGL).
func (self *TilingSprite) SetTilePattern(member *PIXITexture) {
    self.Set("tilePattern", member)
}

// If true the TilingSprite will run generateTexture on its **next** render pass.
// This is set by the likes of Phaser.LoadTexture.setFrame.
func (self *TilingSprite) GetRefreshTexture() bool{
    return self.Get("refreshTexture").Bool()
}

// If true the TilingSprite will run generateTexture on its **next** render pass.
// This is set by the likes of Phaser.LoadTexture.setFrame.
func (self *TilingSprite) SetRefreshTexture(member bool) {
    self.Set("refreshTexture", member)
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TilingSprite) GetAnchor() *Point{
    return &Point{self.Get("anchor")}
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TilingSprite) SetAnchor(member *Point) {
    self.Set("anchor", member)
}

// The texture that the sprite is using
func (self *TilingSprite) GetTexture() *Texture{
    return &Texture{self.Get("texture")}
}

// The texture that the sprite is using
func (self *TilingSprite) SetTexture(member *Texture) {
    self.Set("texture", member)
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TilingSprite) GetTintedTexture() *Canvas{
    return &Canvas{self.Get("tintedTexture")}
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TilingSprite) SetTintedTexture(member *Canvas) {
    self.Set("tintedTexture", member)
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *TilingSprite) GetShader() *AbstractFilter{
    return &AbstractFilter{self.Get("shader")}
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *TilingSprite) SetShader(member *AbstractFilter) {
    self.Set("shader", member)
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *TilingSprite) GetExists() bool{
    return self.Get("exists").Bool()
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *TilingSprite) SetExists(member bool) {
    self.Set("exists", member)
}

// [read-only] The array of children of this container.
func (self *TilingSprite) GetChildren() []DisplayObject{
	array := self.Get("children")
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// [read-only] The array of children of this container.
func (self *TilingSprite) SetChildren(member []DisplayObject) {
    self.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TilingSprite) GetIgnoreChildInput() bool{
    return self.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TilingSprite) SetIgnoreChildInput(member bool) {
    self.Set("ignoreChildInput", member)
}



// Renders the object using the WebGL renderer
func (self *TilingSprite) _renderWebGLI(args ...interface{}) {
    self.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *TilingSprite) _renderCanvasI(args ...interface{}) {
    self.Call("_renderCanvas", args)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *TilingSprite) OnTextureUpdateI(args ...interface{}) {
    self.Call("onTextureUpdate", args)
}

// 
func (self *TilingSprite) GenerateTilingTextureI(args ...interface{}) {
    self.Call("generateTilingTexture", args)
}

// Returns the framing rectangle of the sprite as a PIXI.Rectangle object
func (self *TilingSprite) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getBounds", args)}
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *TilingSprite) SetTextureI(args ...interface{}) {
    self.Call("setTexture", args)
}

// Adds a child to the container.
func (self *TilingSprite) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *TilingSprite) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *TilingSprite) SwapChildrenI(args ...interface{}) {
    self.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *TilingSprite) GetChildIndexI(args ...interface{}) int{
    return self.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *TilingSprite) SetChildIndexI(args ...interface{}) {
    self.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *TilingSprite) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *TilingSprite) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *TilingSprite) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *TilingSprite) RemoveChildrenI(args ...interface{}) {
    self.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *TilingSprite) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *TilingSprite) SetStageReferenceI(args ...interface{}) {
    self.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *TilingSprite) RemoveStageReferenceI(args ...interface{}) {
    self.Call("removeStageReference", args)
}
