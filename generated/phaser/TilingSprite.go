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
func (self *TilingSprite) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the tiling sprite
func (self *TilingSprite) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the tiling sprite
func (self *TilingSprite) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the tiling sprite
func (self *TilingSprite) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The scaling of the image that is being tiled
func (self *TilingSprite) GetTileScaleA() *Point{
    return &Point{self.Object.Get("tileScale")}
}

// The scaling of the image that is being tiled
func (self *TilingSprite) SetTileScaleA(member *Point) {
    self.Object.Set("tileScale", member)
}

// A point that represents the scale of the texture object
func (self *TilingSprite) GetTileScaleOffsetA() *Point{
    return &Point{self.Object.Get("tileScaleOffset")}
}

// A point that represents the scale of the texture object
func (self *TilingSprite) SetTileScaleOffsetA(member *Point) {
    self.Object.Set("tileScaleOffset", member)
}

// The offset position of the image that is being tiled
func (self *TilingSprite) GetTilePositionA() *Point{
    return &Point{self.Object.Get("tilePosition")}
}

// The offset position of the image that is being tiled
func (self *TilingSprite) SetTilePositionA(member *Point) {
    self.Object.Set("tilePosition", member)
}

// Whether this sprite is renderable or not
func (self *TilingSprite) GetRenderableA() bool{
    return self.Object.Get("renderable").Bool()
}

// Whether this sprite is renderable or not
func (self *TilingSprite) SetRenderableA(member bool) {
    self.Object.Set("renderable", member)
}

// The tint applied to the sprite. This is a hex value
func (self *TilingSprite) GetTintA() int{
    return self.Object.Get("tint").Int()
}

// The tint applied to the sprite. This is a hex value
func (self *TilingSprite) SetTintA(member int) {
    self.Object.Set("tint", member)
}

// If enabled a green rectangle will be drawn behind the generated tiling texture, allowing you to visually
// debug the texture being used.
func (self *TilingSprite) GetTextureDebugA() bool{
    return self.Object.Get("textureDebug").Bool()
}

// If enabled a green rectangle will be drawn behind the generated tiling texture, allowing you to visually
// debug the texture being used.
func (self *TilingSprite) SetTextureDebugA(member bool) {
    self.Object.Set("textureDebug", member)
}

// The blend mode to be applied to the sprite
func (self *TilingSprite) GetBlendModeA() int{
    return self.Object.Get("blendMode").Int()
}

// The blend mode to be applied to the sprite
func (self *TilingSprite) SetBlendModeA(member int) {
    self.Object.Set("blendMode", member)
}

// The CanvasBuffer object that the tiled texture is drawn to.
func (self *TilingSprite) GetCanvasBufferA() *PIXICanvasBuffer{
    return &PIXICanvasBuffer{self.Object.Get("canvasBuffer")}
}

// The CanvasBuffer object that the tiled texture is drawn to.
func (self *TilingSprite) SetCanvasBufferA(member *PIXICanvasBuffer) {
    self.Object.Set("canvasBuffer", member)
}

// An internal Texture object that holds the tiling texture that was generated from TilingSprite.texture.
func (self *TilingSprite) GetTilingTextureA() *PIXITexture{
    return &PIXITexture{self.Object.Get("tilingTexture")}
}

// An internal Texture object that holds the tiling texture that was generated from TilingSprite.texture.
func (self *TilingSprite) SetTilingTextureA(member *PIXITexture) {
    self.Object.Set("tilingTexture", member)
}

// The Context fill pattern that is used to draw the TilingSprite in Canvas mode only (will be null in WebGL).
func (self *TilingSprite) GetTilePatternA() *PIXITexture{
    return &PIXITexture{self.Object.Get("tilePattern")}
}

// The Context fill pattern that is used to draw the TilingSprite in Canvas mode only (will be null in WebGL).
func (self *TilingSprite) SetTilePatternA(member *PIXITexture) {
    self.Object.Set("tilePattern", member)
}

// If true the TilingSprite will run generateTexture on its **next** render pass.
// This is set by the likes of Phaser.LoadTexture.setFrame.
func (self *TilingSprite) GetRefreshTextureA() bool{
    return self.Object.Get("refreshTexture").Bool()
}

// If true the TilingSprite will run generateTexture on its **next** render pass.
// This is set by the likes of Phaser.LoadTexture.setFrame.
func (self *TilingSprite) SetRefreshTextureA(member bool) {
    self.Object.Set("refreshTexture", member)
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TilingSprite) GetAnchorA() *Point{
    return &Point{self.Object.Get("anchor")}
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TilingSprite) SetAnchorA(member *Point) {
    self.Object.Set("anchor", member)
}

// The texture that the sprite is using
func (self *TilingSprite) GetTextureA() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// The texture that the sprite is using
func (self *TilingSprite) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TilingSprite) GetTintedTextureA() *Canvas{
    return &Canvas{self.Object.Get("tintedTexture")}
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TilingSprite) SetTintedTextureA(member *Canvas) {
    self.Object.Set("tintedTexture", member)
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *TilingSprite) GetShaderA() *AbstractFilter{
    return &AbstractFilter{self.Object.Get("shader")}
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *TilingSprite) SetShaderA(member *AbstractFilter) {
    self.Object.Set("shader", member)
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *TilingSprite) GetExistsA() bool{
    return self.Object.Get("exists").Bool()
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *TilingSprite) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// [read-only] The array of children of this container.
func (self *TilingSprite) GetChildrenA() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// [read-only] The array of children of this container.
func (self *TilingSprite) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TilingSprite) GetIgnoreChildInputA() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TilingSprite) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}



// Renders the object using the WebGL renderer
func (self *TilingSprite) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// Renders the object using the WebGL renderer
func (self *TilingSprite) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *TilingSprite) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// Renders the object using the Canvas renderer
func (self *TilingSprite) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *TilingSprite) OnTextureUpdate(event interface{}) {
    self.Object.Call("onTextureUpdate", event)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *TilingSprite) OnTextureUpdateI(args ...interface{}) {
    self.Object.Call("onTextureUpdate", args)
}

// 
func (self *TilingSprite) GenerateTilingTexture(forcePowerOfTwo bool, renderSession *RenderSession) {
    self.Object.Call("generateTilingTexture", forcePowerOfTwo, renderSession)
}

// 
func (self *TilingSprite) GenerateTilingTextureI(args ...interface{}) {
    self.Object.Call("generateTilingTexture", args)
}

// Returns the framing rectangle of the sprite as a PIXI.Rectangle object
func (self *TilingSprite) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// Returns the framing rectangle of the sprite as a PIXI.Rectangle object
func (self *TilingSprite) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *TilingSprite) SetTexture(texture *Texture, destroy bool) {
    self.Object.Call("setTexture", texture, destroy)
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *TilingSprite) SetTextureI(args ...interface{}) {
    self.Object.Call("setTexture", args)
}

// Adds a child to the container.
func (self *TilingSprite) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// Adds a child to the container.
func (self *TilingSprite) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *TilingSprite) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *TilingSprite) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *TilingSprite) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// Swaps the position of 2 Display Objects within this container.
func (self *TilingSprite) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *TilingSprite) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// Returns the index position of a child DisplayObject instance
func (self *TilingSprite) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *TilingSprite) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// Changes the position of an existing child in the display object container
func (self *TilingSprite) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *TilingSprite) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// Returns the child at the specified index
func (self *TilingSprite) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *TilingSprite) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// Removes a child from the container.
func (self *TilingSprite) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *TilingSprite) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// Removes a child from the specified index position.
func (self *TilingSprite) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *TilingSprite) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// Removes all children from this container that are within the begin and end indexes.
func (self *TilingSprite) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *TilingSprite) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *TilingSprite) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *TilingSprite) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *TilingSprite) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *TilingSprite) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// Removes the current stage reference from the container and all of its children.
func (self *TilingSprite) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}
