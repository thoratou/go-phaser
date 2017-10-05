// Package phaser Automatic generation for PIXI.TilingSprite
// generated file TilingSprite.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// TilingSprite A tiling sprite is a fast way of rendering a tiling image
type TilingSprite struct {
    *js.Object
}

// NewTilingSprite A tiling sprite is a fast way of rendering a tiling image
func NewTilingSprite(texture *Texture, width int, height int) *TilingSprite {
    return &TilingSprite{js.Global.Get("PIXI").Get("TilingSprite").New(texture, width, height)}
}
// NewTilingSpriteI A tiling sprite is a fast way of rendering a tiling image
func NewTilingSpriteI(args ...interface{}) *TilingSprite {
    return &TilingSprite{js.Global.Get("PIXI").Get("TilingSprite").New(args)}
}



// TilingSprite Binding conversion method to TilingSprite point 
func ToTilingSprite(jsStruct interface{}) *TilingSprite {
    if object, ok := jsStruct.(*js.Object); ok {
		return &TilingSprite{Object: object}
	}
	return nil
}



// Width The width of the tiling sprite
func (self *TilingSprite) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the tiling sprite
func (self *TilingSprite) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the tiling sprite
func (self *TilingSprite) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the tiling sprite
func (self *TilingSprite) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// TileScale The scaling of the image that is being tiled
func (self *TilingSprite) TileScale() *Point{
    return &Point{self.Object.Get("tileScale")}
}

// SetTileScaleA The scaling of the image that is being tiled
func (self *TilingSprite) SetTileScaleA(member *Point) {
    self.Object.Set("tileScale", member)
}

// TileScaleOffset A point that represents the scale of the texture object
func (self *TilingSprite) TileScaleOffset() *Point{
    return &Point{self.Object.Get("tileScaleOffset")}
}

// SetTileScaleOffsetA A point that represents the scale of the texture object
func (self *TilingSprite) SetTileScaleOffsetA(member *Point) {
    self.Object.Set("tileScaleOffset", member)
}

// TilePosition The offset position of the image that is being tiled
func (self *TilingSprite) TilePosition() *Point{
    return &Point{self.Object.Get("tilePosition")}
}

// SetTilePositionA The offset position of the image that is being tiled
func (self *TilingSprite) SetTilePositionA(member *Point) {
    self.Object.Set("tilePosition", member)
}

// Renderable Whether this sprite is renderable or not
func (self *TilingSprite) Renderable() bool{
    return self.Object.Get("renderable").Bool()
}

// SetRenderableA Whether this sprite is renderable or not
func (self *TilingSprite) SetRenderableA(member bool) {
    self.Object.Set("renderable", member)
}

// Tint The tint applied to the sprite. This is a hex value
func (self *TilingSprite) Tint() int{
    return self.Object.Get("tint").Int()
}

// SetTintA The tint applied to the sprite. This is a hex value
func (self *TilingSprite) SetTintA(member int) {
    self.Object.Set("tint", member)
}

// TextureDebug If enabled a green rectangle will be drawn behind the generated tiling texture, allowing you to visually
// 
// debug the texture being used.
func (self *TilingSprite) TextureDebug() bool{
    return self.Object.Get("textureDebug").Bool()
}

// SetTextureDebugA If enabled a green rectangle will be drawn behind the generated tiling texture, allowing you to visually
// 
// debug the texture being used.
func (self *TilingSprite) SetTextureDebugA(member bool) {
    self.Object.Set("textureDebug", member)
}

// BlendMode The blend mode to be applied to the sprite
func (self *TilingSprite) BlendMode() int{
    return self.Object.Get("blendMode").Int()
}

// SetBlendModeA The blend mode to be applied to the sprite
func (self *TilingSprite) SetBlendModeA(member int) {
    self.Object.Set("blendMode", member)
}

// CanvasBuffer The CanvasBuffer object that the tiled texture is drawn to.
func (self *TilingSprite) CanvasBuffer() *PIXICanvasBuffer{
    return &PIXICanvasBuffer{self.Object.Get("canvasBuffer")}
}

// SetCanvasBufferA The CanvasBuffer object that the tiled texture is drawn to.
func (self *TilingSprite) SetCanvasBufferA(member *PIXICanvasBuffer) {
    self.Object.Set("canvasBuffer", member)
}

// TilingTexture An internal Texture object that holds the tiling texture that was generated from TilingSprite.texture.
func (self *TilingSprite) TilingTexture() *PIXITexture{
    return &PIXITexture{self.Object.Get("tilingTexture")}
}

// SetTilingTextureA An internal Texture object that holds the tiling texture that was generated from TilingSprite.texture.
func (self *TilingSprite) SetTilingTextureA(member *PIXITexture) {
    self.Object.Set("tilingTexture", member)
}

// TilePattern The Context fill pattern that is used to draw the TilingSprite in Canvas mode only (will be null in WebGL).
func (self *TilingSprite) TilePattern() *PIXITexture{
    return &PIXITexture{self.Object.Get("tilePattern")}
}

// SetTilePatternA The Context fill pattern that is used to draw the TilingSprite in Canvas mode only (will be null in WebGL).
func (self *TilingSprite) SetTilePatternA(member *PIXITexture) {
    self.Object.Set("tilePattern", member)
}

// RefreshTexture If true the TilingSprite will run generateTexture on its **next** render pass.
// 
// This is set by the likes of Phaser.LoadTexture.setFrame.
func (self *TilingSprite) RefreshTexture() bool{
    return self.Object.Get("refreshTexture").Bool()
}

// SetRefreshTextureA If true the TilingSprite will run generateTexture on its **next** render pass.
// 
// This is set by the likes of Phaser.LoadTexture.setFrame.
func (self *TilingSprite) SetRefreshTextureA(member bool) {
    self.Object.Set("refreshTexture", member)
}

// Anchor The anchor sets the origin point of the texture.
// 
// The default is 0,0 this means the texture's origin is the top left
// 
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// 
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TilingSprite) Anchor() *Point{
    return &Point{self.Object.Get("anchor")}
}

// SetAnchorA The anchor sets the origin point of the texture.
// 
// The default is 0,0 this means the texture's origin is the top left
// 
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// 
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TilingSprite) SetAnchorA(member *Point) {
    self.Object.Set("anchor", member)
}

// Texture The texture that the sprite is using
func (self *TilingSprite) Texture() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// SetTextureA The texture that the sprite is using
func (self *TilingSprite) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// TintedTexture A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TilingSprite) TintedTexture() *Canvas{
    return &Canvas{self.Object.Get("tintedTexture")}
}

// SetTintedTextureA A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TilingSprite) SetTintedTextureA(member *Canvas) {
    self.Object.Set("tintedTexture", member)
}

// Shader The shader that will be used to render this Sprite.
// 
// Set to null to remove a current shader.
func (self *TilingSprite) Shader() *AbstractFilter{
    return &AbstractFilter{self.Object.Get("shader")}
}

// SetShaderA The shader that will be used to render this Sprite.
// 
// Set to null to remove a current shader.
func (self *TilingSprite) SetShaderA(member *AbstractFilter) {
    self.Object.Set("shader", member)
}

// Exists Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *TilingSprite) Exists() bool{
    return self.Object.Get("exists").Bool()
}

// SetExistsA Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *TilingSprite) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// Children [read-only] The array of children of this container.
func (self *TilingSprite) Children() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// SetChildrenA [read-only] The array of children of this container.
func (self *TilingSprite) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// IgnoreChildInput If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// 
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// 
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TilingSprite) IgnoreChildInput() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// SetIgnoreChildInputA If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// 
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// 
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TilingSprite) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}


// _renderWebGL Renders the object using the WebGL renderer
func (self *TilingSprite) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// _renderWebGLI Renders the object using the WebGL renderer
func (self *TilingSprite) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// _renderCanvas Renders the object using the Canvas renderer
func (self *TilingSprite) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// _renderCanvasI Renders the object using the Canvas renderer
func (self *TilingSprite) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

// OnTextureUpdate When the texture is updated, this event will fire to update the scale and frame
func (self *TilingSprite) OnTextureUpdate(event interface{}) {
    self.Object.Call("onTextureUpdate", event)
}

// OnTextureUpdateI When the texture is updated, this event will fire to update the scale and frame
func (self *TilingSprite) OnTextureUpdateI(args ...interface{}) {
    self.Object.Call("onTextureUpdate", args)
}

// GenerateTilingTexture empty description
func (self *TilingSprite) GenerateTilingTexture(forcePowerOfTwo bool, renderSession *RenderSession) {
    self.Object.Call("generateTilingTexture", forcePowerOfTwo, renderSession)
}

// GenerateTilingTextureI empty description
func (self *TilingSprite) GenerateTilingTextureI(args ...interface{}) {
    self.Object.Call("generateTilingTexture", args)
}

// GetBounds Returns the framing rectangle of the sprite as a PIXI.Rectangle object
func (self *TilingSprite) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// GetBoundsI Returns the framing rectangle of the sprite as a PIXI.Rectangle object
func (self *TilingSprite) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// SetTexture Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// 
// texture this Sprite was using.
func (self *TilingSprite) SetTexture(texture *Texture) {
    self.Object.Call("setTexture", texture)
}

// SetTexture1O Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// 
// texture this Sprite was using.
func (self *TilingSprite) SetTexture1O(texture *Texture, destroy bool) {
    self.Object.Call("setTexture", texture, destroy)
}

// SetTextureI Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// 
// texture this Sprite was using.
func (self *TilingSprite) SetTextureI(args ...interface{}) {
    self.Object.Call("setTexture", args)
}

// GetLocalBounds Retrieves the non-global local bounds of the Sprite as a rectangle. The calculation takes all visible children into consideration.
func (self *TilingSprite) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// GetLocalBoundsI Retrieves the non-global local bounds of the Sprite as a rectangle. The calculation takes all visible children into consideration.
func (self *TilingSprite) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// AddChild Adds a child to the container.
func (self *TilingSprite) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// AddChildI Adds a child to the container.
func (self *TilingSprite) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// AddChildAt Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *TilingSprite) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// AddChildAtI Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *TilingSprite) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// SwapChildren Swaps the position of 2 Display Objects within this container.
func (self *TilingSprite) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// SwapChildrenI Swaps the position of 2 Display Objects within this container.
func (self *TilingSprite) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// GetChildIndex Returns the index position of a child DisplayObject instance
func (self *TilingSprite) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// GetChildIndexI Returns the index position of a child DisplayObject instance
func (self *TilingSprite) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// SetChildIndex Changes the position of an existing child in the display object container
func (self *TilingSprite) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// SetChildIndexI Changes the position of an existing child in the display object container
func (self *TilingSprite) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// GetChildAt Returns the child at the specified index
func (self *TilingSprite) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// GetChildAtI Returns the child at the specified index
func (self *TilingSprite) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// RemoveChild Removes a child from the container.
func (self *TilingSprite) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// RemoveChildI Removes a child from the container.
func (self *TilingSprite) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// RemoveChildAt Removes a child from the specified index position.
func (self *TilingSprite) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// RemoveChildAtI Removes a child from the specified index position.
func (self *TilingSprite) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// RemoveChildren Removes all children from this container that are within the begin and end indexes.
func (self *TilingSprite) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// RemoveChildrenI Removes all children from this container that are within the begin and end indexes.
func (self *TilingSprite) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// Contains Determines whether the specified display object is a child of the DisplayObjectContainer instance or the instance itself.
func (self *TilingSprite) Contains(child *DisplayObject) bool{
    return self.Object.Call("contains", child).Bool()
}

// ContainsI Determines whether the specified display object is a child of the DisplayObjectContainer instance or the instance itself.
func (self *TilingSprite) ContainsI(args ...interface{}) bool{
    return self.Object.Call("contains", args).Bool()
}

