// Automatic generation for Phaser.TileSprite
// generated file TileSprite.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A TileSprite is a Sprite that has a repeating texture. The texture can be scrolled and scaled independently of the TileSprite itself.
// Textures will automatically wrap and are designed so that you can create game backdrops using seamless textures as a source.
// 
// TileSprites have no input handler or physics bodies by default, both need enabling in the same way as for normal Sprites.
// 
// You shouldn't ever create a TileSprite any larger than your actual screen size. If you want to create a large repeating background
// that scrolls across the whole map of your game, then you create a TileSprite that fits the screen size and then use the `tilePosition`
// property to scroll the texture as the player moves. If you create a TileSprite that is thousands of pixels in size then it will 
// consume huge amounts of memory and cause performance issues. Remember: use `tilePosition` to scroll your texture and `tileScale` to
// adjust the scale of the texture - don't resize the sprite itself or make it larger than it needs.
// 
// An important note about texture dimensions:
// 
// When running under Canvas a TileSprite can use any texture size without issue. When running under WebGL the texture should ideally be
// a power of two in size (i.e. 4, 8, 16, 32, 64, 128, 256, 512, etch pixels width by height). If the texture isn't a power of two
// it will be rendered to a blank canvas that is the correct size, which means you may have 'blank' areas appearing to the right and
// bottom of your frame. To avoid this ensure your textures are perfect powers of two.
// 
// TileSprites support animations in the same way that Sprites do. You add and play animations using the AnimationManager. However
// if your game is running under WebGL please note that each frame of the animation must be a power of two in size, or it will receive
// additional padding to enforce it to be so.
type TileSprite struct {
    *js.Object
}


// The const type of this object.
func (self *TileSprite) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *TileSprite) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The const physics body type of this object.
func (self *TileSprite) GetPhysicsTypeA() int{
    return self.Object.Get("physicsType").Int()
}

// The const physics body type of this object.
func (self *TileSprite) SetPhysicsTypeA(member int) {
    self.Object.Set("physicsType", member)
}

// The width of the tiling sprite
func (self *TileSprite) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the tiling sprite
func (self *TileSprite) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the tiling sprite
func (self *TileSprite) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the tiling sprite
func (self *TileSprite) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The scaling of the image that is being tiled
func (self *TileSprite) GetTileScaleA() *Point{
    return &Point{self.Object.Get("tileScale")}
}

// The scaling of the image that is being tiled
func (self *TileSprite) SetTileScaleA(member *Point) {
    self.Object.Set("tileScale", member)
}

// A point that represents the scale of the texture object
func (self *TileSprite) GetTileScaleOffsetA() *Point{
    return &Point{self.Object.Get("tileScaleOffset")}
}

// A point that represents the scale of the texture object
func (self *TileSprite) SetTileScaleOffsetA(member *Point) {
    self.Object.Set("tileScaleOffset", member)
}

// The offset position of the image that is being tiled
func (self *TileSprite) GetTilePositionA() *Point{
    return &Point{self.Object.Get("tilePosition")}
}

// The offset position of the image that is being tiled
func (self *TileSprite) SetTilePositionA(member *Point) {
    self.Object.Set("tilePosition", member)
}

// Whether this sprite is renderable or not
func (self *TileSprite) GetRenderableA() bool{
    return self.Object.Get("renderable").Bool()
}

// Whether this sprite is renderable or not
func (self *TileSprite) SetRenderableA(member bool) {
    self.Object.Set("renderable", member)
}

// The tint applied to the sprite. This is a hex value
func (self *TileSprite) GetTintA() int{
    return self.Object.Get("tint").Int()
}

// The tint applied to the sprite. This is a hex value
func (self *TileSprite) SetTintA(member int) {
    self.Object.Set("tint", member)
}

// If enabled a green rectangle will be drawn behind the generated tiling texture, allowing you to visually
// debug the texture being used.
func (self *TileSprite) GetTextureDebugA() bool{
    return self.Object.Get("textureDebug").Bool()
}

// If enabled a green rectangle will be drawn behind the generated tiling texture, allowing you to visually
// debug the texture being used.
func (self *TileSprite) SetTextureDebugA(member bool) {
    self.Object.Set("textureDebug", member)
}

// The blend mode to be applied to the sprite
func (self *TileSprite) GetBlendModeA() int{
    return self.Object.Get("blendMode").Int()
}

// The blend mode to be applied to the sprite
func (self *TileSprite) SetBlendModeA(member int) {
    self.Object.Set("blendMode", member)
}

// The CanvasBuffer object that the tiled texture is drawn to.
func (self *TileSprite) GetCanvasBufferA() *PIXICanvasBuffer{
    return &PIXICanvasBuffer{self.Object.Get("canvasBuffer")}
}

// The CanvasBuffer object that the tiled texture is drawn to.
func (self *TileSprite) SetCanvasBufferA(member *PIXICanvasBuffer) {
    self.Object.Set("canvasBuffer", member)
}

// An internal Texture object that holds the tiling texture that was generated from TilingSprite.texture.
func (self *TileSprite) GetTilingTextureA() *PIXITexture{
    return &PIXITexture{self.Object.Get("tilingTexture")}
}

// An internal Texture object that holds the tiling texture that was generated from TilingSprite.texture.
func (self *TileSprite) SetTilingTextureA(member *PIXITexture) {
    self.Object.Set("tilingTexture", member)
}

// The Context fill pattern that is used to draw the TilingSprite in Canvas mode only (will be null in WebGL).
func (self *TileSprite) GetTilePatternA() *PIXITexture{
    return &PIXITexture{self.Object.Get("tilePattern")}
}

// The Context fill pattern that is used to draw the TilingSprite in Canvas mode only (will be null in WebGL).
func (self *TileSprite) SetTilePatternA(member *PIXITexture) {
    self.Object.Set("tilePattern", member)
}

// If true the TilingSprite will run generateTexture on its **next** render pass.
// This is set by the likes of Phaser.LoadTexture.setFrame.
func (self *TileSprite) GetRefreshTextureA() bool{
    return self.Object.Get("refreshTexture").Bool()
}

// If true the TilingSprite will run generateTexture on its **next** render pass.
// This is set by the likes of Phaser.LoadTexture.setFrame.
func (self *TileSprite) SetRefreshTextureA(member bool) {
    self.Object.Set("refreshTexture", member)
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TileSprite) GetAnchorA() *Point{
    return &Point{self.Object.Get("anchor")}
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TileSprite) SetAnchorA(member *Point) {
    self.Object.Set("anchor", member)
}

// The texture that the sprite is using
func (self *TileSprite) GetTextureA() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// The texture that the sprite is using
func (self *TileSprite) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TileSprite) GetTintedTextureA() *Canvas{
    return &Canvas{self.Object.Get("tintedTexture")}
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TileSprite) SetTintedTextureA(member *Canvas) {
    self.Object.Set("tintedTexture", member)
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *TileSprite) GetShaderA() *AbstractFilter{
    return &AbstractFilter{self.Object.Get("shader")}
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *TileSprite) SetShaderA(member *AbstractFilter) {
    self.Object.Set("shader", member)
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *TileSprite) GetExistsA() bool{
    return self.Object.Get("exists").Bool()
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *TileSprite) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// [read-only] The array of children of this container.
func (self *TileSprite) GetChildrenA() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// [read-only] The array of children of this container.
func (self *TileSprite) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TileSprite) GetIgnoreChildInputA() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TileSprite) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// A reference to the currently running Game.
func (self *TileSprite) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *TileSprite) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *TileSprite) GetNameA() string{
    return self.Object.Get("name").String()
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *TileSprite) SetNameA(member string) {
    self.Object.Set("name", member)
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *TileSprite) GetDataA() interface{}{
    return self.Object.Get("data")
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *TileSprite) SetDataA(member interface{}) {
    self.Object.Set("data", member)
}

// The components this Game Object has installed.
func (self *TileSprite) GetComponentsA() interface{}{
    return self.Object.Get("components")
}

// The components this Game Object has installed.
func (self *TileSprite) SetComponentsA(member interface{}) {
    self.Object.Set("components", member)
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *TileSprite) GetZA() int{
    return self.Object.Get("z").Int()
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *TileSprite) SetZA(member int) {
    self.Object.Set("z", member)
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *TileSprite) GetEventsA() *Events{
    return &Events{self.Object.Get("events")}
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *TileSprite) SetEventsA(member *Events) {
    self.Object.Set("events", member)
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *TileSprite) GetAnimationsA() *AnimationManager{
    return &AnimationManager{self.Object.Get("animations")}
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *TileSprite) SetAnimationsA(member *AnimationManager) {
    self.Object.Set("animations", member)
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *TileSprite) GetKeyA() interface{}{
    return self.Object.Get("key")
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *TileSprite) SetKeyA(member interface{}) {
    self.Object.Set("key", member)
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *TileSprite) GetWorldA() *Point{
    return &Point{self.Object.Get("world")}
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *TileSprite) SetWorldA(member *Point) {
    self.Object.Set("world", member)
}

// A debug flag designed for use with `Game.enableStep`.
func (self *TileSprite) GetDebugA() bool{
    return self.Object.Get("debug").Bool()
}

// A debug flag designed for use with `Game.enableStep`.
func (self *TileSprite) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// The position the Game Object was located in the previous frame.
func (self *TileSprite) GetPreviousPositionA() *Point{
    return &Point{self.Object.Get("previousPosition")}
}

// The position the Game Object was located in the previous frame.
func (self *TileSprite) SetPreviousPositionA(member *Point) {
    self.Object.Set("previousPosition", member)
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *TileSprite) GetPreviousRotationA() int{
    return self.Object.Get("previousRotation").Int()
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *TileSprite) SetPreviousRotationA(member int) {
    self.Object.Set("previousRotation", member)
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *TileSprite) GetRenderOrderIDA() int{
    return self.Object.Get("renderOrderID").Int()
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *TileSprite) SetRenderOrderIDA(member int) {
    self.Object.Set("renderOrderID", member)
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *TileSprite) GetFreshA() bool{
    return self.Object.Get("fresh").Bool()
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *TileSprite) SetFreshA(member bool) {
    self.Object.Set("fresh", member)
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *TileSprite) GetPendingDestroyA() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *TileSprite) SetPendingDestroyA(member bool) {
    self.Object.Set("pendingDestroy", member)
}

// The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *TileSprite) GetAngleA() int{
    return self.Object.Get("angle").Int()
}

// The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *TileSprite) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *TileSprite) GetAutoCullA() bool{
    return self.Object.Get("autoCull").Bool()
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *TileSprite) SetAutoCullA(member bool) {
    self.Object.Set("autoCull", member)
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *TileSprite) GetInCameraA() bool{
    return self.Object.Get("inCamera").Bool()
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *TileSprite) SetInCameraA(member bool) {
    self.Object.Set("inCamera", member)
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *TileSprite) GetOffsetXA() int{
    return self.Object.Get("offsetX").Int()
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *TileSprite) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *TileSprite) GetOffsetYA() int{
    return self.Object.Get("offsetY").Int()
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *TileSprite) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *TileSprite) GetCenterXA() int{
    return self.Object.Get("centerX").Int()
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *TileSprite) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *TileSprite) GetCenterYA() int{
    return self.Object.Get("centerY").Int()
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *TileSprite) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *TileSprite) GetLeftA() int{
    return self.Object.Get("left").Int()
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *TileSprite) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *TileSprite) GetRightA() int{
    return self.Object.Get("right").Int()
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *TileSprite) SetRightA(member int) {
    self.Object.Set("right", member)
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *TileSprite) GetTopA() int{
    return self.Object.Get("top").Int()
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *TileSprite) SetTopA(member int) {
    self.Object.Set("top", member)
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *TileSprite) GetBottomA() int{
    return self.Object.Get("bottom").Int()
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *TileSprite) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *TileSprite) GetDestroyPhaseA() bool{
    return self.Object.Get("destroyPhase").Bool()
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *TileSprite) SetDestroyPhaseA(member bool) {
    self.Object.Set("destroyPhase", member)
}

// A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
// 
// The values are adjusted at the rendering stage, overriding the Game Objects actual world position.
// 
// The end result is that the Game Object will appear to be 'fixed' to the camera, regardless of where in the game world
// the camera is viewing. This is useful if for example this Game Object is a UI item that you wish to be visible at all times 
// regardless where in the world the camera is.
// 
// The offsets are stored in the `cameraOffset` property.
// 
// Note that the `cameraOffset` values are in addition to any parent of this Game Object on the display list.
// 
// Be careful not to set `fixedToCamera` on Game Objects which are in Groups that already have `fixedToCamera` enabled on them.
func (self *TileSprite) GetFixedToCameraA() bool{
    return self.Object.Get("fixedToCamera").Bool()
}

// A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
// 
// The values are adjusted at the rendering stage, overriding the Game Objects actual world position.
// 
// The end result is that the Game Object will appear to be 'fixed' to the camera, regardless of where in the game world
// the camera is viewing. This is useful if for example this Game Object is a UI item that you wish to be visible at all times 
// regardless where in the world the camera is.
// 
// The offsets are stored in the `cameraOffset` property.
// 
// Note that the `cameraOffset` values are in addition to any parent of this Game Object on the display list.
// 
// Be careful not to set `fixedToCamera` on Game Objects which are in Groups that already have `fixedToCamera` enabled on them.
func (self *TileSprite) SetFixedToCameraA(member bool) {
    self.Object.Set("fixedToCamera", member)
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *TileSprite) GetCameraOffsetA() *Point{
    return &Point{self.Object.Get("cameraOffset")}
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *TileSprite) SetCameraOffsetA(member *Point) {
    self.Object.Set("cameraOffset", member)
}

// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *TileSprite) GetHealthA() int{
    return self.Object.Get("health").Int()
}

// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *TileSprite) SetHealthA(member int) {
    self.Object.Set("health", member)
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *TileSprite) GetMaxHealthA() int{
    return self.Object.Get("maxHealth").Int()
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *TileSprite) SetMaxHealthA(member int) {
    self.Object.Set("maxHealth", member)
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *TileSprite) GetDamageA() interface{}{
    return self.Object.Get("damage")
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *TileSprite) SetDamageA(member interface{}) {
    self.Object.Set("damage", member)
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *TileSprite) GetSetHealthA() interface{}{
    return self.Object.Get("setHealth")
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *TileSprite) SetSetHealthA(member interface{}) {
    self.Object.Set("setHealth", member)
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *TileSprite) GetHealA() interface{}{
    return self.Object.Get("heal")
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *TileSprite) SetHealA(member interface{}) {
    self.Object.Set("heal", member)
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *TileSprite) GetInputA() interface{}{
    return self.Object.Get("input")
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *TileSprite) SetInputA(member interface{}) {
    self.Object.Set("input", member)
}

// By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
// for this Game Object and it will then start to process click / touch events and more.
// 
// You can then access the Input Handler via `this.input`.
// 
// Note that Input related events are dispatched from `this.events`, i.e.: `events.onInputDown`.
// 
// If you set this property to false it will stop the Input Handler from processing any more input events.
// 
// If you want to _temporarily_ disable input for a Game Object, then it's better to set
// `input.enabled = false`, as it won't reset any of the Input Handlers internal properties.
// You can then toggle this back on as needed.
func (self *TileSprite) GetInputEnabledA() bool{
    return self.Object.Get("inputEnabled").Bool()
}

// By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
// for this Game Object and it will then start to process click / touch events and more.
// 
// You can then access the Input Handler via `this.input`.
// 
// Note that Input related events are dispatched from `this.events`, i.e.: `events.onInputDown`.
// 
// If you set this property to false it will stop the Input Handler from processing any more input events.
// 
// If you want to _temporarily_ disable input for a Game Object, then it's better to set
// `input.enabled = false`, as it won't reset any of the Input Handlers internal properties.
// You can then toggle this back on as needed.
func (self *TileSprite) SetInputEnabledA(member bool) {
    self.Object.Set("inputEnabled", member)
}

// If this is set to `true` the Game Object checks if it is within the World bounds each frame. 
// 
// When it is no longer intersecting the world bounds it dispatches the `onOutOfBounds` event.
// 
// If it was *previously* out of bounds but is now intersecting the world bounds again it dispatches the `onEnterBounds` event.
// 
// It also optionally kills the Game Object if `outOfBoundsKill` is `true`.
// 
// When `checkWorldBounds` is enabled it forces the Game Object to calculate its full bounds every frame.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *TileSprite) GetCheckWorldBoundsA() bool{
    return self.Object.Get("checkWorldBounds").Bool()
}

// If this is set to `true` the Game Object checks if it is within the World bounds each frame. 
// 
// When it is no longer intersecting the world bounds it dispatches the `onOutOfBounds` event.
// 
// If it was *previously* out of bounds but is now intersecting the world bounds again it dispatches the `onEnterBounds` event.
// 
// It also optionally kills the Game Object if `outOfBoundsKill` is `true`.
// 
// When `checkWorldBounds` is enabled it forces the Game Object to calculate its full bounds every frame.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *TileSprite) SetCheckWorldBoundsA(member bool) {
    self.Object.Set("checkWorldBounds", member)
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *TileSprite) GetOutOfBoundsKillA() bool{
    return self.Object.Get("outOfBoundsKill").Bool()
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *TileSprite) SetOutOfBoundsKillA(member bool) {
    self.Object.Set("outOfBoundsKill", member)
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *TileSprite) GetOutOfCameraBoundsKillA() bool{
    return self.Object.Get("outOfCameraBoundsKill").Bool()
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *TileSprite) SetOutOfCameraBoundsKillA(member bool) {
    self.Object.Set("outOfCameraBoundsKill", member)
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *TileSprite) GetInWorldA() bool{
    return self.Object.Get("inWorld").Bool()
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *TileSprite) SetInWorldA(member bool) {
    self.Object.Set("inWorld", member)
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *TileSprite) GetAliveA() bool{
    return self.Object.Get("alive").Bool()
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *TileSprite) SetAliveA(member bool) {
    self.Object.Set("alive", member)
}

// The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *TileSprite) GetLifespanA() int{
    return self.Object.Get("lifespan").Int()
}

// The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *TileSprite) SetLifespanA(member int) {
    self.Object.Set("lifespan", member)
}

// Gets or sets the current frame index of the texture being used to render this Game Object.
// 
// To change the frame set `frame` to the index of the new frame in the sprite sheet you wish this Game Object to use,
// for example: `player.frame = 4`.
// 
// If the frame index given doesn't exist it will revert to the first frame found in the texture.
// 
// If you are using a texture atlas then you should use the `frameName` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *TileSprite) GetFrameA() int{
    return self.Object.Get("frame").Int()
}

// Gets or sets the current frame index of the texture being used to render this Game Object.
// 
// To change the frame set `frame` to the index of the new frame in the sprite sheet you wish this Game Object to use,
// for example: `player.frame = 4`.
// 
// If the frame index given doesn't exist it will revert to the first frame found in the texture.
// 
// If you are using a texture atlas then you should use the `frameName` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *TileSprite) SetFrameA(member int) {
    self.Object.Set("frame", member)
}

// Gets or sets the current frame name of the texture being used to render this Game Object.
// 
// To change the frame set `frameName` to the name of the new frame in the texture atlas you wish this Game Object to use, 
// for example: `player.frameName = "idle"`.
// 
// If the frame name given doesn't exist it will revert to the first frame found in the texture and throw a console warning.
// 
// If you are using a sprite sheet then you should use the `frame` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *TileSprite) GetFrameNameA() string{
    return self.Object.Get("frameName").String()
}

// Gets or sets the current frame name of the texture being used to render this Game Object.
// 
// To change the frame set `frameName` to the name of the new frame in the texture atlas you wish this Game Object to use, 
// for example: `player.frameName = "idle"`.
// 
// If the frame name given doesn't exist it will revert to the first frame found in the texture and throw a console warning.
// 
// If you are using a sprite sheet then you should use the `frame` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *TileSprite) SetFrameNameA(member string) {
    self.Object.Set("frameName", member)
}

// `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
// properties and methods via it.
// 
// By default Game Objects won't add themselves to any physics system and their `body` property will be `null`.
// 
// To enable this Game Object for physics you need to call `game.physics.enable(object, system)` where `object` is this object
// and `system` is the Physics system you are using. If none is given it defaults to `Phaser.Physics.Arcade`.
// 
// You can alternatively call `game.physics.arcade.enable(object)`, or add this Game Object to a physics enabled Group.
// 
// Important: Enabling a Game Object for P2 or Ninja physics will automatically set its `anchor` property to 0.5, 
// so the physics body is centered on the Game Object.
// 
// If you need a different result then adjust or re-create the Body shape offsets manually or reset the anchor after enabling physics.
func (self *TileSprite) GetBodyA() interface{}{
    return self.Object.Get("body")
}

// `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
// properties and methods via it.
// 
// By default Game Objects won't add themselves to any physics system and their `body` property will be `null`.
// 
// To enable this Game Object for physics you need to call `game.physics.enable(object, system)` where `object` is this object
// and `system` is the Physics system you are using. If none is given it defaults to `Phaser.Physics.Arcade`.
// 
// You can alternatively call `game.physics.arcade.enable(object)`, or add this Game Object to a physics enabled Group.
// 
// Important: Enabling a Game Object for P2 or Ninja physics will automatically set its `anchor` property to 0.5, 
// so the physics body is centered on the Game Object.
// 
// If you need a different result then adjust or re-create the Body shape offsets manually or reset the anchor after enabling physics.
func (self *TileSprite) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *TileSprite) GetXA() int{
    return self.Object.Get("x").Int()
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *TileSprite) SetXA(member int) {
    self.Object.Set("x", member)
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *TileSprite) GetYA() int{
    return self.Object.Get("y").Int()
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *TileSprite) SetYA(member int) {
    self.Object.Set("y", member)
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *TileSprite) GetSmoothedA() bool{
    return self.Object.Get("smoothed").Bool()
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *TileSprite) SetSmoothedA(member bool) {
    self.Object.Set("smoothed", member)
}



// Automatically called by World.preUpdate.
func (self *TileSprite) PreUpdate() {
    self.Object.Call("preUpdate")
}

// Automatically called by World.preUpdate.
func (self *TileSprite) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Sets this TileSprite to automatically scroll in the given direction until stopped via TileSprite.stopScroll().
// The scroll speed is specified in pixels per second.
// A negative x value will scroll to the left. A positive x value will scroll to the right.
// A negative y value will scroll up. A positive y value will scroll down.
func (self *TileSprite) AutoScroll(x int, y int) {
    self.Object.Call("autoScroll", x, y)
}

// Sets this TileSprite to automatically scroll in the given direction until stopped via TileSprite.stopScroll().
// The scroll speed is specified in pixels per second.
// A negative x value will scroll to the left. A positive x value will scroll to the right.
// A negative y value will scroll up. A positive y value will scroll down.
func (self *TileSprite) AutoScrollI(args ...interface{}) {
    self.Object.Call("autoScroll", args)
}

// Stops an automatically scrolling TileSprite.
func (self *TileSprite) StopScroll() {
    self.Object.Call("stopScroll")
}

// Stops an automatically scrolling TileSprite.
func (self *TileSprite) StopScrollI(args ...interface{}) {
    self.Object.Call("stopScroll", args)
}

// Destroys the TileSprite. This removes it from its parent group, destroys the event and animation handlers if present
// and nulls its reference to game, freeing it up for garbage collection.
func (self *TileSprite) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the TileSprite. This removes it from its parent group, destroys the event and animation handlers if present
// and nulls its reference to game, freeing it up for garbage collection.
func (self *TileSprite) Destroy1O(destroyChildren bool) {
    self.Object.Call("destroy", destroyChildren)
}

// Destroys the TileSprite. This removes it from its parent group, destroys the event and animation handlers if present
// and nulls its reference to game, freeing it up for garbage collection.
func (self *TileSprite) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Resets the TileSprite. This places the TileSprite at the given x/y world coordinates, resets the tilePosition and then
// sets alive, exists, visible and renderable all to true. Also resets the outOfBounds state.
// If the TileSprite has a physics body that too is reset.
func (self *TileSprite) Reset(x int, y int) *TileSprite{
    return &TileSprite{self.Object.Call("reset", x, y)}
}

// Resets the TileSprite. This places the TileSprite at the given x/y world coordinates, resets the tilePosition and then
// sets alive, exists, visible and renderable all to true. Also resets the outOfBounds state.
// If the TileSprite has a physics body that too is reset.
func (self *TileSprite) ResetI(args ...interface{}) *TileSprite{
    return &TileSprite{self.Object.Call("reset", args)}
}

// Renders the object using the WebGL renderer
func (self *TileSprite) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// Renders the object using the WebGL renderer
func (self *TileSprite) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *TileSprite) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// Renders the object using the Canvas renderer
func (self *TileSprite) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *TileSprite) OnTextureUpdate(event interface{}) {
    self.Object.Call("onTextureUpdate", event)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *TileSprite) OnTextureUpdateI(args ...interface{}) {
    self.Object.Call("onTextureUpdate", args)
}

// 
func (self *TileSprite) GenerateTilingTexture(forcePowerOfTwo bool, renderSession *RenderSession) {
    self.Object.Call("generateTilingTexture", forcePowerOfTwo, renderSession)
}

// 
func (self *TileSprite) GenerateTilingTextureI(args ...interface{}) {
    self.Object.Call("generateTilingTexture", args)
}

// Returns the framing rectangle of the sprite as a PIXI.Rectangle object
func (self *TileSprite) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// Returns the framing rectangle of the sprite as a PIXI.Rectangle object
func (self *TileSprite) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *TileSprite) SetTexture(texture *Texture) {
    self.Object.Call("setTexture", texture)
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *TileSprite) SetTexture1O(texture *Texture, destroy bool) {
    self.Object.Call("setTexture", texture, destroy)
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *TileSprite) SetTextureI(args ...interface{}) {
    self.Object.Call("setTexture", args)
}

// Adds a child to the container.
func (self *TileSprite) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// Adds a child to the container.
func (self *TileSprite) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *TileSprite) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *TileSprite) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *TileSprite) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// Swaps the position of 2 Display Objects within this container.
func (self *TileSprite) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *TileSprite) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// Returns the index position of a child DisplayObject instance
func (self *TileSprite) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *TileSprite) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// Changes the position of an existing child in the display object container
func (self *TileSprite) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *TileSprite) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// Returns the child at the specified index
func (self *TileSprite) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *TileSprite) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// Removes a child from the container.
func (self *TileSprite) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *TileSprite) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// Removes a child from the specified index position.
func (self *TileSprite) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *TileSprite) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// Removes all children from this container that are within the begin and end indexes.
func (self *TileSprite) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *TileSprite) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *TileSprite) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *TileSprite) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *TileSprite) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *TileSprite) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// Removes the current stage reference from the container and all of its children.
func (self *TileSprite) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *TileSprite) Update() {
    self.Object.Call("update")
}

// Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *TileSprite) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Internal method called by the World postUpdate cycle.
func (self *TileSprite) PostUpdate() {
    self.Object.Call("postUpdate")
}

// Internal method called by the World postUpdate cycle.
func (self *TileSprite) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TileSprite) Play(name string) *Animation{
    return &Animation{self.Object.Call("play", name)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TileSprite) Play1O(name string, frameRate int) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TileSprite) Play2O(name string, frameRate int, loop bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TileSprite) Play3O(name string, frameRate int, loop bool, killOnComplete bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop, killOnComplete)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TileSprite) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("play", args)}
}

// Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *TileSprite) AlignIn(container interface{}) interface{}{
    return self.Object.Call("alignIn", container)
}

// Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *TileSprite) AlignIn1O(container interface{}, position int) interface{}{
    return self.Object.Call("alignIn", container, position)
}

// Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *TileSprite) AlignIn2O(container interface{}, position int, offsetX int) interface{}{
    return self.Object.Call("alignIn", container, position, offsetX)
}

// Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *TileSprite) AlignIn3O(container interface{}, position int, offsetX int, offsetY int) interface{}{
    return self.Object.Call("alignIn", container, position, offsetX, offsetY)
}

// Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *TileSprite) AlignInI(args ...interface{}) interface{}{
    return self.Object.Call("alignIn", args)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *TileSprite) AlignTo(parent interface{}) interface{}{
    return self.Object.Call("alignTo", parent)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *TileSprite) AlignTo1O(parent interface{}, position int) interface{}{
    return self.Object.Call("alignTo", parent, position)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *TileSprite) AlignTo2O(parent interface{}, position int, offsetX int) interface{}{
    return self.Object.Call("alignTo", parent, position, offsetX)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *TileSprite) AlignTo3O(parent interface{}, position int, offsetX int, offsetY int) interface{}{
    return self.Object.Call("alignTo", parent, position, offsetX, offsetY)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *TileSprite) AlignToI(args ...interface{}) interface{}{
    return self.Object.Call("alignTo", args)
}

// Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TileSprite) BringToTop() *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop")}
}

// Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TileSprite) BringToTopI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop", args)}
}

// Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TileSprite) SendToBack() *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack")}
}

// Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TileSprite) SendToBackI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack", args)}
}

// Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TileSprite) MoveUp() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp")}
}

// Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TileSprite) MoveUpI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp", args)}
}

// Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TileSprite) MoveDown() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown")}
}

// Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TileSprite) MoveDownI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown", args)}
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *TileSprite) Revive() *DisplayObject{
    return &DisplayObject{self.Object.Call("revive")}
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *TileSprite) Revive1O(health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", health)}
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *TileSprite) ReviveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", args)}
}

// Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *TileSprite) Kill() *DisplayObject{
    return &DisplayObject{self.Object.Call("kill")}
}

// Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *TileSprite) KillI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("kill", args)}
}

// Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *TileSprite) LoadTexture(key interface{}) {
    self.Object.Call("loadTexture", key)
}

// Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *TileSprite) LoadTexture1O(key interface{}, frame interface{}) {
    self.Object.Call("loadTexture", key, frame)
}

// Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *TileSprite) LoadTexture2O(key interface{}, frame interface{}, stopAnimation bool) {
    self.Object.Call("loadTexture", key, frame, stopAnimation)
}

// Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *TileSprite) LoadTextureI(args ...interface{}) {
    self.Object.Call("loadTexture", args)
}

// Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *TileSprite) SetFrame(frame *Frame) {
    self.Object.Call("setFrame", frame)
}

// Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *TileSprite) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *TileSprite) ResizeFrame(parent interface{}, width int, height int) {
    self.Object.Call("resizeFrame", parent, width, height)
}

// Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *TileSprite) ResizeFrameI(args ...interface{}) {
    self.Object.Call("resizeFrame", args)
}

// Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *TileSprite) ResetFrame() {
    self.Object.Call("resetFrame")
}

// Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *TileSprite) ResetFrameI(args ...interface{}) {
    self.Object.Call("resetFrame", args)
}

// Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *TileSprite) Overlap(displayObject interface{}) bool{
    return self.Object.Call("overlap", displayObject).Bool()
}

// Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *TileSprite) OverlapI(args ...interface{}) bool{
    return self.Object.Call("overlap", args).Bool()
}
