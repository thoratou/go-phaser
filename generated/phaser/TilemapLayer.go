// Package phaser Automatic generation for Phaser.TilemapLayer
// generated file TilemapLayer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// TilemapLayer A TilemapLayer is a Phaser.Image/Sprite that renders a specific TileLayer of a Tilemap.
// 
// Since a TilemapLayer is a Sprite it can be moved around the display, added to other groups or display objects, etc.
// 
// By default TilemapLayers have fixedToCamera set to `true`. Changing this will break Camera follow and scrolling behavior.
type TilemapLayer struct {
    *js.Object
}

// NewTilemapLayer A TilemapLayer is a Phaser.Image/Sprite that renders a specific TileLayer of a Tilemap.
// 
// Since a TilemapLayer is a Sprite it can be moved around the display, added to other groups or display objects, etc.
// 
// By default TilemapLayers have fixedToCamera set to `true`. Changing this will break Camera follow and scrolling behavior.
func NewTilemapLayer(game *Game, tilemap *Tilemap, index int, width int, height int) *TilemapLayer {
    return &TilemapLayer{js.Global.Get("Phaser").Get("TilemapLayer").New(game, tilemap, index, width, height)}
}
// NewTilemapLayerI A TilemapLayer is a Phaser.Image/Sprite that renders a specific TileLayer of a Tilemap.
// 
// Since a TilemapLayer is a Sprite it can be moved around the display, added to other groups or display objects, etc.
// 
// By default TilemapLayers have fixedToCamera set to `true`. Changing this will break Camera follow and scrolling behavior.
func NewTilemapLayerI(args ...interface{}) *TilemapLayer {
    return &TilemapLayer{js.Global.Get("Phaser").Get("TilemapLayer").New(args)}
}



// TilemapLayer Binding conversion method to TilemapLayer point 
func ToTilemapLayer(jsStruct interface{}) *TilemapLayer {
    if object, ok := jsStruct.(*js.Object); ok {
		return &TilemapLayer{Object: object}
	}
	return nil
}



// Map The Tilemap to which this layer is bound.
func (self *TilemapLayer) Map() *Tilemap{
    return &Tilemap{self.Object.Get("map")}
}

// SetMapA The Tilemap to which this layer is bound.
func (self *TilemapLayer) SetMapA(member *Tilemap) {
    self.Object.Set("map", member)
}

// Index The index of this layer within the Tilemap.
func (self *TilemapLayer) Index() int{
    return self.Object.Get("index").Int()
}

// SetIndexA The index of this layer within the Tilemap.
func (self *TilemapLayer) SetIndexA(member int) {
    self.Object.Set("index", member)
}

// Layer The layer object within the Tilemap that this layer represents.
func (self *TilemapLayer) Layer() interface{}{
    return self.Object.Get("layer")
}

// SetLayerA The layer object within the Tilemap that this layer represents.
func (self *TilemapLayer) SetLayerA(member interface{}) {
    self.Object.Set("layer", member)
}

// Canvas The canvas to which this TilemapLayer draws.
func (self *TilemapLayer) Canvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("canvas"))
}

// SetCanvasA The canvas to which this TilemapLayer draws.
func (self *TilemapLayer) SetCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("canvas", member)
}

// Type The const type of this object.
func (self *TilemapLayer) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *TilemapLayer) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// PhysicsType The const physics body type of this object.
func (self *TilemapLayer) PhysicsType() int{
    return self.Object.Get("physicsType").Int()
}

// SetPhysicsTypeA The const physics body type of this object.
func (self *TilemapLayer) SetPhysicsTypeA(member int) {
    self.Object.Set("physicsType", member)
}

// RenderSettings Settings that control standard (non-diagnostic) rendering.
func (self *TilemapLayer) RenderSettings() interface{}{
    return self.Object.Get("renderSettings")
}

// SetRenderSettingsA Settings that control standard (non-diagnostic) rendering.
func (self *TilemapLayer) SetRenderSettingsA(member interface{}) {
    self.Object.Set("renderSettings", member)
}

// Debug Enable an additional "debug rendering" pass to display collision information.
func (self *TilemapLayer) Debug() bool{
    return self.Object.Get("debug").Bool()
}

// SetDebugA Enable an additional "debug rendering" pass to display collision information.
func (self *TilemapLayer) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// Exists Controls if the core game loop and physics update this game object or not.
func (self *TilemapLayer) Exists() bool{
    return self.Object.Get("exists").Bool()
}

// SetExistsA Controls if the core game loop and physics update this game object or not.
func (self *TilemapLayer) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// DebugSettings Settings used for debugging and diagnostics.
func (self *TilemapLayer) DebugSettings() interface{}{
    return self.Object.Get("debugSettings")
}

// SetDebugSettingsA Settings used for debugging and diagnostics.
func (self *TilemapLayer) SetDebugSettingsA(member interface{}) {
    self.Object.Set("debugSettings", member)
}

// ScrollFactorX Speed at which this layer scrolls horizontally, relative to the camera (e.g. scrollFactorX of 0.5 scrolls half as quickly as the 'normal' camera-locked layers do).
func (self *TilemapLayer) ScrollFactorX() int{
    return self.Object.Get("scrollFactorX").Int()
}

// SetScrollFactorXA Speed at which this layer scrolls horizontally, relative to the camera (e.g. scrollFactorX of 0.5 scrolls half as quickly as the 'normal' camera-locked layers do).
func (self *TilemapLayer) SetScrollFactorXA(member int) {
    self.Object.Set("scrollFactorX", member)
}

// ScrollFactorY Speed at which this layer scrolls vertically, relative to the camera (e.g. scrollFactorY of 0.5 scrolls half as quickly as the 'normal' camera-locked layers do)
func (self *TilemapLayer) ScrollFactorY() int{
    return self.Object.Get("scrollFactorY").Int()
}

// SetScrollFactorYA Speed at which this layer scrolls vertically, relative to the camera (e.g. scrollFactorY of 0.5 scrolls half as quickly as the 'normal' camera-locked layers do)
func (self *TilemapLayer) SetScrollFactorYA(member int) {
    self.Object.Set("scrollFactorY", member)
}

// Dirty If true tiles will be force rendered, even if such is not believed to be required.
func (self *TilemapLayer) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// SetDirtyA If true tiles will be force rendered, even if such is not believed to be required.
func (self *TilemapLayer) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// RayStepRate When ray-casting against tiles this is the number of steps it will jump. For larger tile sizes you can increase this to improve performance.
func (self *TilemapLayer) RayStepRate() int{
    return self.Object.Get("rayStepRate").Int()
}

// SetRayStepRateA When ray-casting against tiles this is the number of steps it will jump. For larger tile sizes you can increase this to improve performance.
func (self *TilemapLayer) SetRayStepRateA(member int) {
    self.Object.Set("rayStepRate", member)
}

// Anchor The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TilemapLayer) Anchor() *Point{
    return &Point{self.Object.Get("anchor")}
}

// SetAnchorA The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TilemapLayer) SetAnchorA(member *Point) {
    self.Object.Set("anchor", member)
}

// Texture The texture that the sprite is using
func (self *TilemapLayer) Texture() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// SetTextureA The texture that the sprite is using
func (self *TilemapLayer) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// Tint The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *TilemapLayer) Tint() int{
    return self.Object.Get("tint").Int()
}

// SetTintA The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *TilemapLayer) SetTintA(member int) {
    self.Object.Set("tint", member)
}

// TintedTexture A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TilemapLayer) TintedTexture() *Canvas{
    return &Canvas{self.Object.Get("tintedTexture")}
}

// SetTintedTextureA A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TilemapLayer) SetTintedTextureA(member *Canvas) {
    self.Object.Set("tintedTexture", member)
}

// BlendMode The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *TilemapLayer) BlendMode() int{
    return self.Object.Get("blendMode").Int()
}

// SetBlendModeA The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *TilemapLayer) SetBlendModeA(member int) {
    self.Object.Set("blendMode", member)
}

// Shader The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *TilemapLayer) Shader() *AbstractFilter{
    return &AbstractFilter{self.Object.Get("shader")}
}

// SetShaderA The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *TilemapLayer) SetShaderA(member *AbstractFilter) {
    self.Object.Set("shader", member)
}

// Width The width of the sprite, setting this will actually modify the scale to achieve the value set
func (self *TilemapLayer) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the sprite, setting this will actually modify the scale to achieve the value set
func (self *TilemapLayer) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the sprite, setting this will actually modify the scale to achieve the value set
func (self *TilemapLayer) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the sprite, setting this will actually modify the scale to achieve the value set
func (self *TilemapLayer) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Children [read-only] The array of children of this container.
func (self *TilemapLayer) Children() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// SetChildrenA [read-only] The array of children of this container.
func (self *TilemapLayer) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// IgnoreChildInput If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TilemapLayer) IgnoreChildInput() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// SetIgnoreChildInputA If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TilemapLayer) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// Game A reference to the currently running Game.
func (self *TilemapLayer) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *TilemapLayer) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Name A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *TilemapLayer) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *TilemapLayer) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Data An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *TilemapLayer) Data() interface{}{
    return self.Object.Get("data")
}

// SetDataA An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *TilemapLayer) SetDataA(member interface{}) {
    self.Object.Set("data", member)
}

// Components The components this Game Object has installed.
func (self *TilemapLayer) Components() interface{}{
    return self.Object.Get("components")
}

// SetComponentsA The components this Game Object has installed.
func (self *TilemapLayer) SetComponentsA(member interface{}) {
    self.Object.Set("components", member)
}

// Z The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *TilemapLayer) Z() int{
    return self.Object.Get("z").Int()
}

// SetZA The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *TilemapLayer) SetZA(member int) {
    self.Object.Set("z", member)
}

// Events All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *TilemapLayer) Events() *Events{
    return &Events{self.Object.Get("events")}
}

// SetEventsA All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *TilemapLayer) SetEventsA(member *Events) {
    self.Object.Set("events", member)
}

// Animations If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *TilemapLayer) Animations() *AnimationManager{
    return &AnimationManager{self.Object.Get("animations")}
}

// SetAnimationsA If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *TilemapLayer) SetAnimationsA(member *AnimationManager) {
    self.Object.Set("animations", member)
}

// Key The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *TilemapLayer) Key() interface{}{
    return self.Object.Get("key")
}

// SetKeyA The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *TilemapLayer) SetKeyA(member interface{}) {
    self.Object.Set("key", member)
}

// World The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *TilemapLayer) World() *Point{
    return &Point{self.Object.Get("world")}
}

// SetWorldA The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *TilemapLayer) SetWorldA(member *Point) {
    self.Object.Set("world", member)
}

// PreviousPosition The position the Game Object was located in the previous frame.
func (self *TilemapLayer) PreviousPosition() *Point{
    return &Point{self.Object.Get("previousPosition")}
}

// SetPreviousPositionA The position the Game Object was located in the previous frame.
func (self *TilemapLayer) SetPreviousPositionA(member *Point) {
    self.Object.Set("previousPosition", member)
}

// PreviousRotation The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *TilemapLayer) PreviousRotation() int{
    return self.Object.Get("previousRotation").Int()
}

// SetPreviousRotationA The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *TilemapLayer) SetPreviousRotationA(member int) {
    self.Object.Set("previousRotation", member)
}

// RenderOrderID The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *TilemapLayer) RenderOrderID() int{
    return self.Object.Get("renderOrderID").Int()
}

// SetRenderOrderIDA The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *TilemapLayer) SetRenderOrderIDA(member int) {
    self.Object.Set("renderOrderID", member)
}

// Fresh A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *TilemapLayer) Fresh() bool{
    return self.Object.Get("fresh").Bool()
}

// SetFreshA A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *TilemapLayer) SetFreshA(member bool) {
    self.Object.Set("fresh", member)
}

// PendingDestroy A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *TilemapLayer) PendingDestroy() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// SetPendingDestroyA A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *TilemapLayer) SetPendingDestroyA(member bool) {
    self.Object.Set("pendingDestroy", member)
}

// Angle The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *TilemapLayer) Angle() int{
    return self.Object.Get("angle").Int()
}

// SetAngleA The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *TilemapLayer) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// AutoCull A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *TilemapLayer) AutoCull() bool{
    return self.Object.Get("autoCull").Bool()
}

// SetAutoCullA A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *TilemapLayer) SetAutoCullA(member bool) {
    self.Object.Set("autoCull", member)
}

// InCamera Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *TilemapLayer) InCamera() bool{
    return self.Object.Get("inCamera").Bool()
}

// SetInCameraA Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *TilemapLayer) SetInCameraA(member bool) {
    self.Object.Set("inCamera", member)
}

// OffsetX The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *TilemapLayer) OffsetX() int{
    return self.Object.Get("offsetX").Int()
}

// SetOffsetXA The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *TilemapLayer) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// OffsetY The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *TilemapLayer) OffsetY() int{
    return self.Object.Get("offsetY").Int()
}

// SetOffsetYA The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *TilemapLayer) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}

// CenterX The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *TilemapLayer) CenterX() int{
    return self.Object.Get("centerX").Int()
}

// SetCenterXA The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *TilemapLayer) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// CenterY The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *TilemapLayer) CenterY() int{
    return self.Object.Get("centerY").Int()
}

// SetCenterYA The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *TilemapLayer) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// Left The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *TilemapLayer) Left() int{
    return self.Object.Get("left").Int()
}

// SetLeftA The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *TilemapLayer) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// Right The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *TilemapLayer) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *TilemapLayer) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Top The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *TilemapLayer) Top() int{
    return self.Object.Get("top").Int()
}

// SetTopA The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *TilemapLayer) SetTopA(member int) {
    self.Object.Set("top", member)
}

// Bottom The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *TilemapLayer) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *TilemapLayer) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// CropRect The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *TilemapLayer) CropRect() *Rectangle{
    return &Rectangle{self.Object.Get("cropRect")}
}

// SetCropRectA The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *TilemapLayer) SetCropRectA(member *Rectangle) {
    self.Object.Set("cropRect", member)
}

// DeltaX Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *TilemapLayer) DeltaX() int{
    return self.Object.Get("deltaX").Int()
}

// SetDeltaXA Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *TilemapLayer) SetDeltaXA(member int) {
    self.Object.Set("deltaX", member)
}

// DeltaY Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *TilemapLayer) DeltaY() int{
    return self.Object.Get("deltaY").Int()
}

// SetDeltaYA Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *TilemapLayer) SetDeltaYA(member int) {
    self.Object.Set("deltaY", member)
}

// DeltaZ Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *TilemapLayer) DeltaZ() int{
    return self.Object.Get("deltaZ").Int()
}

// SetDeltaZA Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *TilemapLayer) SetDeltaZA(member int) {
    self.Object.Set("deltaZ", member)
}

// DestroyPhase As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *TilemapLayer) DestroyPhase() bool{
    return self.Object.Get("destroyPhase").Bool()
}

// SetDestroyPhaseA As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *TilemapLayer) SetDestroyPhaseA(member bool) {
    self.Object.Set("destroyPhase", member)
}

// FixedToCamera A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
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
func (self *TilemapLayer) FixedToCamera() bool{
    return self.Object.Get("fixedToCamera").Bool()
}

// SetFixedToCameraA A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
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
func (self *TilemapLayer) SetFixedToCameraA(member bool) {
    self.Object.Set("fixedToCamera", member)
}

// CameraOffset The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *TilemapLayer) CameraOffset() *Point{
    return &Point{self.Object.Get("cameraOffset")}
}

// SetCameraOffsetA The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *TilemapLayer) SetCameraOffsetA(member *Point) {
    self.Object.Set("cameraOffset", member)
}

// Health The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *TilemapLayer) Health() int{
    return self.Object.Get("health").Int()
}

// SetHealthA The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *TilemapLayer) SetHealthA(member int) {
    self.Object.Set("health", member)
}

// MaxHealth The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *TilemapLayer) MaxHealth() int{
    return self.Object.Get("maxHealth").Int()
}

// SetMaxHealthA The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *TilemapLayer) SetMaxHealthA(member int) {
    self.Object.Set("maxHealth", member)
}

// Damage Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *TilemapLayer) Damage() interface{}{
    return self.Object.Get("damage")
}

// SetDamageA Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *TilemapLayer) SetDamageA(member interface{}) {
    self.Object.Set("damage", member)
}

// SetHealth Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *TilemapLayer) SetHealth() interface{}{
    return self.Object.Get("setHealth")
}

// SetSetHealthA Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *TilemapLayer) SetSetHealthA(member interface{}) {
    self.Object.Set("setHealth", member)
}

// Heal Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *TilemapLayer) Heal() interface{}{
    return self.Object.Get("heal")
}

// SetHealA Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *TilemapLayer) SetHealA(member interface{}) {
    self.Object.Set("heal", member)
}

// Input The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *TilemapLayer) Input() interface{}{
    return self.Object.Get("input")
}

// SetInputA The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *TilemapLayer) SetInputA(member interface{}) {
    self.Object.Set("input", member)
}

// InputEnabled By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
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
func (self *TilemapLayer) InputEnabled() bool{
    return self.Object.Get("inputEnabled").Bool()
}

// SetInputEnabledA By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
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
func (self *TilemapLayer) SetInputEnabledA(member bool) {
    self.Object.Set("inputEnabled", member)
}

// CheckWorldBounds If this is set to `true` the Game Object checks if it is within the World bounds each frame. 
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
func (self *TilemapLayer) CheckWorldBounds() bool{
    return self.Object.Get("checkWorldBounds").Bool()
}

// SetCheckWorldBoundsA If this is set to `true` the Game Object checks if it is within the World bounds each frame. 
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
func (self *TilemapLayer) SetCheckWorldBoundsA(member bool) {
    self.Object.Set("checkWorldBounds", member)
}

// OutOfBoundsKill If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *TilemapLayer) OutOfBoundsKill() bool{
    return self.Object.Get("outOfBoundsKill").Bool()
}

// SetOutOfBoundsKillA If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *TilemapLayer) SetOutOfBoundsKillA(member bool) {
    self.Object.Set("outOfBoundsKill", member)
}

// OutOfCameraBoundsKill If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *TilemapLayer) OutOfCameraBoundsKill() bool{
    return self.Object.Get("outOfCameraBoundsKill").Bool()
}

// SetOutOfCameraBoundsKillA If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *TilemapLayer) SetOutOfCameraBoundsKillA(member bool) {
    self.Object.Set("outOfCameraBoundsKill", member)
}

// InWorld Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *TilemapLayer) InWorld() bool{
    return self.Object.Get("inWorld").Bool()
}

// SetInWorldA Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *TilemapLayer) SetInWorldA(member bool) {
    self.Object.Set("inWorld", member)
}

// Alive A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *TilemapLayer) Alive() bool{
    return self.Object.Get("alive").Bool()
}

// SetAliveA A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *TilemapLayer) SetAliveA(member bool) {
    self.Object.Set("alive", member)
}

// Lifespan The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *TilemapLayer) Lifespan() int{
    return self.Object.Get("lifespan").Int()
}

// SetLifespanA The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *TilemapLayer) SetLifespanA(member int) {
    self.Object.Set("lifespan", member)
}

// Frame Gets or sets the current frame index of the texture being used to render this Game Object.
// 
// To change the frame set `frame` to the index of the new frame in the sprite sheet you wish this Game Object to use,
// for example: `player.frame = 4`.
// 
// If the frame index given doesn't exist it will revert to the first frame found in the texture.
// 
// If you are using a texture atlas then you should use the `frameName` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *TilemapLayer) Frame() int{
    return self.Object.Get("frame").Int()
}

// SetFrameA Gets or sets the current frame index of the texture being used to render this Game Object.
// 
// To change the frame set `frame` to the index of the new frame in the sprite sheet you wish this Game Object to use,
// for example: `player.frame = 4`.
// 
// If the frame index given doesn't exist it will revert to the first frame found in the texture.
// 
// If you are using a texture atlas then you should use the `frameName` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *TilemapLayer) SetFrameA(member int) {
    self.Object.Set("frame", member)
}

// FrameName Gets or sets the current frame name of the texture being used to render this Game Object.
// 
// To change the frame set `frameName` to the name of the new frame in the texture atlas you wish this Game Object to use, 
// for example: `player.frameName = "idle"`.
// 
// If the frame name given doesn't exist it will revert to the first frame found in the texture and throw a console warning.
// 
// If you are using a sprite sheet then you should use the `frame` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *TilemapLayer) FrameName() string{
    return self.Object.Get("frameName").String()
}

// SetFrameNameA Gets or sets the current frame name of the texture being used to render this Game Object.
// 
// To change the frame set `frameName` to the name of the new frame in the texture atlas you wish this Game Object to use, 
// for example: `player.frameName = "idle"`.
// 
// If the frame name given doesn't exist it will revert to the first frame found in the texture and throw a console warning.
// 
// If you are using a sprite sheet then you should use the `frame` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *TilemapLayer) SetFrameNameA(member string) {
    self.Object.Set("frameName", member)
}

// Body `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
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
func (self *TilemapLayer) Body() interface{}{
    return self.Object.Get("body")
}

// SetBodyA `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
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
func (self *TilemapLayer) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// X The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *TilemapLayer) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *TilemapLayer) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *TilemapLayer) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *TilemapLayer) SetYA(member int) {
    self.Object.Set("y", member)
}

// TransformCallback The callback that will apply any scale limiting to the worldTransform.
func (self *TilemapLayer) TransformCallback() interface{}{
    return self.Object.Get("transformCallback")
}

// SetTransformCallbackA The callback that will apply any scale limiting to the worldTransform.
func (self *TilemapLayer) SetTransformCallbackA(member interface{}) {
    self.Object.Set("transformCallback", member)
}

// TransformCallbackContext The context under which `transformCallback` is called.
func (self *TilemapLayer) TransformCallbackContext() interface{}{
    return self.Object.Get("transformCallbackContext")
}

// SetTransformCallbackContextA The context under which `transformCallback` is called.
func (self *TilemapLayer) SetTransformCallbackContextA(member interface{}) {
    self.Object.Set("transformCallbackContext", member)
}

// ScaleMin The minimum scale this Game Object will scale down to.
// 
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
// 
// Set it to `null` to remove the limit.
func (self *TilemapLayer) ScaleMin() *Point{
    return &Point{self.Object.Get("scaleMin")}
}

// SetScaleMinA The minimum scale this Game Object will scale down to.
// 
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
// 
// Set it to `null` to remove the limit.
func (self *TilemapLayer) SetScaleMinA(member *Point) {
    self.Object.Set("scaleMin", member)
}

// ScaleMax The maximum scale this Game Object will scale up to. 
// 
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
// 
// Set it to `null` to remove the limit.
func (self *TilemapLayer) ScaleMax() *Point{
    return &Point{self.Object.Get("scaleMax")}
}

// SetScaleMaxA The maximum scale this Game Object will scale up to. 
// 
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
// 
// Set it to `null` to remove the limit.
func (self *TilemapLayer) SetScaleMaxA(member *Point) {
    self.Object.Set("scaleMax", member)
}

// Smoothed Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *TilemapLayer) Smoothed() bool{
    return self.Object.Get("smoothed").Bool()
}

// SetSmoothedA Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *TilemapLayer) SetSmoothedA(member bool) {
    self.Object.Set("smoothed", member)
}


// EnsureSharedCopyCanvas Create if needed (and return) a shared copy canvas that is shared across all TilemapLayers.
// 
// Code that uses the canvas is responsible to ensure the dimensions and save/restore state as appropriate.
func (self *TilemapLayer) EnsureSharedCopyCanvas() {
    self.Object.Call("ensureSharedCopyCanvas")
}

// EnsureSharedCopyCanvasI Create if needed (and return) a shared copy canvas that is shared across all TilemapLayers.
// 
// Code that uses the canvas is responsible to ensure the dimensions and save/restore state as appropriate.
func (self *TilemapLayer) EnsureSharedCopyCanvasI(args ...interface{}) {
    self.Object.Call("ensureSharedCopyCanvas", args)
}

// PreUpdate Automatically called by World.preUpdate.
func (self *TilemapLayer) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI Automatically called by World.preUpdate.
func (self *TilemapLayer) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// PostUpdate Automatically called by World.postUpdate. Handles cache updates.
func (self *TilemapLayer) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI Automatically called by World.postUpdate. Handles cache updates.
func (self *TilemapLayer) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// _renderCanvas Automatically called by the Canvas Renderer.
// Overrides the Sprite._renderCanvas function.
func (self *TilemapLayer) _renderCanvas() {
    self.Object.Call("_renderCanvas")
}

// _renderCanvasI Automatically called by the Canvas Renderer.
// Overrides the Sprite._renderCanvas function.
func (self *TilemapLayer) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

// _renderWebGL Automatically called by the Canvas Renderer.
// Overrides the Sprite._renderWebGL function.
func (self *TilemapLayer) _renderWebGL() {
    self.Object.Call("_renderWebGL")
}

// _renderWebGLI Automatically called by the Canvas Renderer.
// Overrides the Sprite._renderWebGL function.
func (self *TilemapLayer) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// Destroy Destroys this TilemapLayer.
func (self *TilemapLayer) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this TilemapLayer.
func (self *TilemapLayer) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Resize Resizes the internal canvas and texture frame used by this TilemapLayer.
// 
// This is an expensive call, so don't bind it to a window resize event! But instead call it at carefully
// selected times.
// 
// Be aware that no validation of the new sizes takes place and the current map scroll coordinates are not
// modified either. You will have to handle both of these things from your game code if required.
func (self *TilemapLayer) Resize(width int, height int) {
    self.Object.Call("resize", width, height)
}

// ResizeI Resizes the internal canvas and texture frame used by this TilemapLayer.
// 
// This is an expensive call, so don't bind it to a window resize event! But instead call it at carefully
// selected times.
// 
// Be aware that no validation of the new sizes takes place and the current map scroll coordinates are not
// modified either. You will have to handle both of these things from your game code if required.
func (self *TilemapLayer) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// ResizeWorld Sets the world size to match the size of this layer.
func (self *TilemapLayer) ResizeWorld() {
    self.Object.Call("resizeWorld")
}

// ResizeWorldI Sets the world size to match the size of this layer.
func (self *TilemapLayer) ResizeWorldI(args ...interface{}) {
    self.Object.Call("resizeWorld", args)
}

// _fixX Take an x coordinate that doesn't account for scrollFactorX and 'fix' it into a scrolled local space.
func (self *TilemapLayer) _fixX(x int) int{
    return self.Object.Call("_fixX", x).Int()
}

// _fixXI Take an x coordinate that doesn't account for scrollFactorX and 'fix' it into a scrolled local space.
func (self *TilemapLayer) _fixXI(args ...interface{}) int{
    return self.Object.Call("_fixX", args).Int()
}

// _unfixX Take an x coordinate that _does_ account for scrollFactorX and 'unfix' it back to camera space.
func (self *TilemapLayer) _unfixX(x int) int{
    return self.Object.Call("_unfixX", x).Int()
}

// _unfixXI Take an x coordinate that _does_ account for scrollFactorX and 'unfix' it back to camera space.
func (self *TilemapLayer) _unfixXI(args ...interface{}) int{
    return self.Object.Call("_unfixX", args).Int()
}

// _fixY Take a y coordinate that doesn't account for scrollFactorY and 'fix' it into a scrolled local space.
func (self *TilemapLayer) _fixY(y int) int{
    return self.Object.Call("_fixY", y).Int()
}

// _fixYI Take a y coordinate that doesn't account for scrollFactorY and 'fix' it into a scrolled local space.
func (self *TilemapLayer) _fixYI(args ...interface{}) int{
    return self.Object.Call("_fixY", args).Int()
}

// _unfixY Take a y coordinate that _does_ account for scrollFactorY and 'unfix' it back to camera space.
func (self *TilemapLayer) _unfixY(y int) int{
    return self.Object.Call("_unfixY", y).Int()
}

// _unfixYI Take a y coordinate that _does_ account for scrollFactorY and 'unfix' it back to camera space.
func (self *TilemapLayer) _unfixYI(args ...interface{}) int{
    return self.Object.Call("_unfixY", args).Int()
}

// GetTileX Convert a pixel value to a tile coordinate.
func (self *TilemapLayer) GetTileX(x int) int{
    return self.Object.Call("getTileX", x).Int()
}

// GetTileXI Convert a pixel value to a tile coordinate.
func (self *TilemapLayer) GetTileXI(args ...interface{}) int{
    return self.Object.Call("getTileX", args).Int()
}

// GetTileY Convert a pixel value to a tile coordinate.
func (self *TilemapLayer) GetTileY(y int) int{
    return self.Object.Call("getTileY", y).Int()
}

// GetTileYI Convert a pixel value to a tile coordinate.
func (self *TilemapLayer) GetTileYI(args ...interface{}) int{
    return self.Object.Call("getTileY", args).Int()
}

// GetTileXY Convert a pixel coordinate to a tile coordinate.
func (self *TilemapLayer) GetTileXY(x int, y int, point interface{}) interface{}{
    return self.Object.Call("getTileXY", x, y, point)
}

// GetTileXYI Convert a pixel coordinate to a tile coordinate.
func (self *TilemapLayer) GetTileXYI(args ...interface{}) interface{}{
    return self.Object.Call("getTileXY", args)
}

// GetRayCastTiles Gets all tiles that intersect with the given line.
func (self *TilemapLayer) GetRayCastTiles(line *Line) []Tile{
	array00 := self.Object.Call("getRayCastTiles", line)
	length00 := array00.Length()
	out00 := make([]Tile, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tile{array00.Index(i00)}
	}
	return out00
}

// GetRayCastTiles1O Gets all tiles that intersect with the given line.
func (self *TilemapLayer) GetRayCastTiles1O(line *Line, stepRate int) []Tile{
	array00 := self.Object.Call("getRayCastTiles", line, stepRate)
	length00 := array00.Length()
	out00 := make([]Tile, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tile{array00.Index(i00)}
	}
	return out00
}

// GetRayCastTiles2O Gets all tiles that intersect with the given line.
func (self *TilemapLayer) GetRayCastTiles2O(line *Line, stepRate int, collides bool) []Tile{
	array00 := self.Object.Call("getRayCastTiles", line, stepRate, collides)
	length00 := array00.Length()
	out00 := make([]Tile, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tile{array00.Index(i00)}
	}
	return out00
}

// GetRayCastTiles3O Gets all tiles that intersect with the given line.
func (self *TilemapLayer) GetRayCastTiles3O(line *Line, stepRate int, collides bool, interestingFace bool) []Tile{
	array00 := self.Object.Call("getRayCastTiles", line, stepRate, collides, interestingFace)
	length00 := array00.Length()
	out00 := make([]Tile, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tile{array00.Index(i00)}
	}
	return out00
}

// GetRayCastTilesI Gets all tiles that intersect with the given line.
func (self *TilemapLayer) GetRayCastTilesI(args ...interface{}) []Tile{
	array00 := self.Object.Call("getRayCastTiles", args)
	length00 := array00.Length()
	out00 := make([]Tile, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tile{array00.Index(i00)}
	}
	return out00
}

// GetTiles Get all tiles that exist within the given area, defined by the top-left corner, width and height. Values given are in pixels, not tiles.
func (self *TilemapLayer) GetTiles(x int, y int, width int, height int) []Tile{
	array00 := self.Object.Call("getTiles", x, y, width, height)
	length00 := array00.Length()
	out00 := make([]Tile, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tile{array00.Index(i00)}
	}
	return out00
}

// GetTiles1O Get all tiles that exist within the given area, defined by the top-left corner, width and height. Values given are in pixels, not tiles.
func (self *TilemapLayer) GetTiles1O(x int, y int, width int, height int, collides bool) []Tile{
	array00 := self.Object.Call("getTiles", x, y, width, height, collides)
	length00 := array00.Length()
	out00 := make([]Tile, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tile{array00.Index(i00)}
	}
	return out00
}

// GetTiles2O Get all tiles that exist within the given area, defined by the top-left corner, width and height. Values given are in pixels, not tiles.
func (self *TilemapLayer) GetTiles2O(x int, y int, width int, height int, collides bool, interestingFace bool) []Tile{
	array00 := self.Object.Call("getTiles", x, y, width, height, collides, interestingFace)
	length00 := array00.Length()
	out00 := make([]Tile, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tile{array00.Index(i00)}
	}
	return out00
}

// GetTilesI Get all tiles that exist within the given area, defined by the top-left corner, width and height. Values given are in pixels, not tiles.
func (self *TilemapLayer) GetTilesI(args ...interface{}) []Tile{
	array00 := self.Object.Call("getTiles", args)
	length00 := array00.Length()
	out00 := make([]Tile, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tile{array00.Index(i00)}
	}
	return out00
}

// ResolveTileset Returns the appropriate tileset for the index, updating the internal cache as required.
// This should only be called if `tilesets[index]` evaluates to undefined.
func (self *TilemapLayer) ResolveTileset(Tile int) interface{}{
    return self.Object.Call("resolveTileset", Tile)
}

// ResolveTilesetI Returns the appropriate tileset for the index, updating the internal cache as required.
// This should only be called if `tilesets[index]` evaluates to undefined.
func (self *TilemapLayer) ResolveTilesetI(args ...interface{}) interface{}{
    return self.Object.Call("resolveTileset", args)
}

// ResetTilesetCache The TilemapLayer caches tileset look-ups.
// 
// Call this method of clear the cache if tilesets have been added or updated after the layer has been rendered.
func (self *TilemapLayer) ResetTilesetCache() {
    self.Object.Call("resetTilesetCache")
}

// ResetTilesetCacheI The TilemapLayer caches tileset look-ups.
// 
// Call this method of clear the cache if tilesets have been added or updated after the layer has been rendered.
func (self *TilemapLayer) ResetTilesetCacheI(args ...interface{}) {
    self.Object.Call("resetTilesetCache", args)
}

// SetScale This method will set the scale of the tilemap as well as update the underlying block data of this layer.
func (self *TilemapLayer) SetScale() {
    self.Object.Call("setScale")
}

// SetScale1O This method will set the scale of the tilemap as well as update the underlying block data of this layer.
func (self *TilemapLayer) SetScale1O(xScale int) {
    self.Object.Call("setScale", xScale)
}

// SetScale2O This method will set the scale of the tilemap as well as update the underlying block data of this layer.
func (self *TilemapLayer) SetScale2O(xScale int, yScale int) {
    self.Object.Call("setScale", xScale, yScale)
}

// SetScaleI This method will set the scale of the tilemap as well as update the underlying block data of this layer.
func (self *TilemapLayer) SetScaleI(args ...interface{}) {
    self.Object.Call("setScale", args)
}

// ShiftCanvas Shifts the contents of the canvas - does extra math so that different browsers agree on the result.
// 
// The specified (x/y) will be shifted to (0,0) after the copy and the newly exposed canvas area will need to be filled in.
func (self *TilemapLayer) ShiftCanvas(context *dom.CanvasRenderingContext2D, x int, y int) {
    self.Object.Call("shiftCanvas", context, x, y)
}

// ShiftCanvasI Shifts the contents of the canvas - does extra math so that different browsers agree on the result.
// 
// The specified (x/y) will be shifted to (0,0) after the copy and the newly exposed canvas area will need to be filled in.
func (self *TilemapLayer) ShiftCanvasI(args ...interface{}) {
    self.Object.Call("shiftCanvas", args)
}

// RenderRegion Render tiles in the given area given by the virtual tile coordinates biased by the given scroll factor.
// This will constrain the tile coordinates based on wrapping but not physical coordinates.
func (self *TilemapLayer) RenderRegion(scrollX int, scrollY int, left int, top int, right int, bottom int) {
    self.Object.Call("renderRegion", scrollX, scrollY, left, top, right, bottom)
}

// RenderRegionI Render tiles in the given area given by the virtual tile coordinates biased by the given scroll factor.
// This will constrain the tile coordinates based on wrapping but not physical coordinates.
func (self *TilemapLayer) RenderRegionI(args ...interface{}) {
    self.Object.Call("renderRegion", args)
}

// RenderDeltaScroll Shifts the canvas and render damaged edge tiles.
func (self *TilemapLayer) RenderDeltaScroll() {
    self.Object.Call("renderDeltaScroll")
}

// RenderDeltaScrollI Shifts the canvas and render damaged edge tiles.
func (self *TilemapLayer) RenderDeltaScrollI(args ...interface{}) {
    self.Object.Call("renderDeltaScroll", args)
}

// RenderFull Clear and render the entire canvas.
func (self *TilemapLayer) RenderFull() {
    self.Object.Call("renderFull")
}

// RenderFullI Clear and render the entire canvas.
func (self *TilemapLayer) RenderFullI(args ...interface{}) {
    self.Object.Call("renderFull", args)
}

// Render Renders the tiles to the layer canvas and pushes to the display.
func (self *TilemapLayer) Render() {
    self.Object.Call("render")
}

// RenderI Renders the tiles to the layer canvas and pushes to the display.
func (self *TilemapLayer) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// RenderDebug Renders a debug overlay on-top of the canvas. Called automatically by render when `debug` is true.
// 
// See `debugSettings` for assorted configuration options.
func (self *TilemapLayer) RenderDebug() {
    self.Object.Call("renderDebug")
}

// RenderDebugI Renders a debug overlay on-top of the canvas. Called automatically by render when `debug` is true.
// 
// See `debugSettings` for assorted configuration options.
func (self *TilemapLayer) RenderDebugI(args ...interface{}) {
    self.Object.Call("renderDebug", args)
}

// SetTexture Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *TilemapLayer) SetTexture(texture *Texture) {
    self.Object.Call("setTexture", texture)
}

// SetTexture1O Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *TilemapLayer) SetTexture1O(texture *Texture, destroy bool) {
    self.Object.Call("setTexture", texture, destroy)
}

// SetTextureI Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *TilemapLayer) SetTextureI(args ...interface{}) {
    self.Object.Call("setTexture", args)
}

// OnTextureUpdate When the texture is updated, this event will fire to update the scale and frame
func (self *TilemapLayer) OnTextureUpdate(event interface{}) {
    self.Object.Call("onTextureUpdate", event)
}

// OnTextureUpdateI When the texture is updated, this event will fire to update the scale and frame
func (self *TilemapLayer) OnTextureUpdateI(args ...interface{}) {
    self.Object.Call("onTextureUpdate", args)
}

// GetBounds Returns the bounds of the Sprite as a rectangle.
// The bounds calculation takes the worldTransform into account.
// 
// It is important to note that the transform is not updated when you call this method.
// So if this Sprite is the child of a Display Object which has had its transform
// updated since the last render pass, those changes will not yet have been applied
// to this Sprites worldTransform. If you need to ensure that all parent transforms
// are factored into this getBounds operation then you should call `updateTransform`
// on the root most object in this Sprites display list first.
func (self *TilemapLayer) GetBounds(matrix *Matrix) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", matrix)}
}

// GetBoundsI Returns the bounds of the Sprite as a rectangle.
// The bounds calculation takes the worldTransform into account.
// 
// It is important to note that the transform is not updated when you call this method.
// So if this Sprite is the child of a Display Object which has had its transform
// updated since the last render pass, those changes will not yet have been applied
// to this Sprites worldTransform. If you need to ensure that all parent transforms
// are factored into this getBounds operation then you should call `updateTransform`
// on the root most object in this Sprites display list first.
func (self *TilemapLayer) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// AddChild Adds a child to the container.
func (self *TilemapLayer) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// AddChildI Adds a child to the container.
func (self *TilemapLayer) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// AddChildAt Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *TilemapLayer) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// AddChildAtI Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *TilemapLayer) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// SwapChildren Swaps the position of 2 Display Objects within this container.
func (self *TilemapLayer) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// SwapChildrenI Swaps the position of 2 Display Objects within this container.
func (self *TilemapLayer) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// GetChildIndex Returns the index position of a child DisplayObject instance
func (self *TilemapLayer) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// GetChildIndexI Returns the index position of a child DisplayObject instance
func (self *TilemapLayer) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// SetChildIndex Changes the position of an existing child in the display object container
func (self *TilemapLayer) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// SetChildIndexI Changes the position of an existing child in the display object container
func (self *TilemapLayer) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// GetChildAt Returns the child at the specified index
func (self *TilemapLayer) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// GetChildAtI Returns the child at the specified index
func (self *TilemapLayer) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// RemoveChild Removes a child from the container.
func (self *TilemapLayer) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// RemoveChildI Removes a child from the container.
func (self *TilemapLayer) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// RemoveChildAt Removes a child from the specified index position.
func (self *TilemapLayer) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// RemoveChildAtI Removes a child from the specified index position.
func (self *TilemapLayer) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// RemoveChildren Removes all children from this container that are within the begin and end indexes.
func (self *TilemapLayer) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// RemoveChildrenI Removes all children from this container that are within the begin and end indexes.
func (self *TilemapLayer) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// GetLocalBounds Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *TilemapLayer) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// GetLocalBoundsI Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *TilemapLayer) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// SetStageReference Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *TilemapLayer) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// SetStageReferenceI Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *TilemapLayer) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// RemoveStageReference Removes the current stage reference from the container and all of its children.
func (self *TilemapLayer) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// RemoveStageReferenceI Removes the current stage reference from the container and all of its children.
func (self *TilemapLayer) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// Update Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *TilemapLayer) Update() {
    self.Object.Call("update")
}

// UpdateI Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *TilemapLayer) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Play Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TilemapLayer) Play(name string) *Animation{
    return &Animation{self.Object.Call("play", name)}
}

// Play1O Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TilemapLayer) Play1O(name string, frameRate int) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate)}
}

// Play2O Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TilemapLayer) Play2O(name string, frameRate int, loop bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop)}
}

// Play3O Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TilemapLayer) Play3O(name string, frameRate int, loop bool, killOnComplete bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop, killOnComplete)}
}

// PlayI Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TilemapLayer) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("play", args)}
}

// AlignIn Aligns this Game Object within another Game Object, or Rectangle, known as the
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
func (self *TilemapLayer) AlignIn(container interface{}) interface{}{
    return self.Object.Call("alignIn", container)
}

// AlignIn1O Aligns this Game Object within another Game Object, or Rectangle, known as the
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
func (self *TilemapLayer) AlignIn1O(container interface{}, position int) interface{}{
    return self.Object.Call("alignIn", container, position)
}

// AlignIn2O Aligns this Game Object within another Game Object, or Rectangle, known as the
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
func (self *TilemapLayer) AlignIn2O(container interface{}, position int, offsetX int) interface{}{
    return self.Object.Call("alignIn", container, position, offsetX)
}

// AlignIn3O Aligns this Game Object within another Game Object, or Rectangle, known as the
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
func (self *TilemapLayer) AlignIn3O(container interface{}, position int, offsetX int, offsetY int) interface{}{
    return self.Object.Call("alignIn", container, position, offsetX, offsetY)
}

// AlignInI Aligns this Game Object within another Game Object, or Rectangle, known as the
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
func (self *TilemapLayer) AlignInI(args ...interface{}) interface{}{
    return self.Object.Call("alignIn", args)
}

// AlignTo Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
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
func (self *TilemapLayer) AlignTo(parent interface{}) interface{}{
    return self.Object.Call("alignTo", parent)
}

// AlignTo1O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
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
func (self *TilemapLayer) AlignTo1O(parent interface{}, position int) interface{}{
    return self.Object.Call("alignTo", parent, position)
}

// AlignTo2O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
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
func (self *TilemapLayer) AlignTo2O(parent interface{}, position int, offsetX int) interface{}{
    return self.Object.Call("alignTo", parent, position, offsetX)
}

// AlignTo3O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
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
func (self *TilemapLayer) AlignTo3O(parent interface{}, position int, offsetX int, offsetY int) interface{}{
    return self.Object.Call("alignTo", parent, position, offsetX, offsetY)
}

// AlignToI Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
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
func (self *TilemapLayer) AlignToI(args ...interface{}) interface{}{
    return self.Object.Call("alignTo", args)
}

// BringToTop Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) BringToTop() *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop")}
}

// BringToTopI Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) BringToTopI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop", args)}
}

// SendToBack Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) SendToBack() *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack")}
}

// SendToBackI Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) SendToBackI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack", args)}
}

// MoveUp Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) MoveUp() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp")}
}

// MoveUpI Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) MoveUpI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp", args)}
}

// MoveDown Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) MoveDown() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown")}
}

// MoveDownI Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) MoveDownI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown", args)}
}

// Crop Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
// 
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
// 
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object 
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
// 
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`, 
// in which case the values are duplicated to a local object.
func (self *TilemapLayer) Crop(rect *Rectangle) {
    self.Object.Call("crop", rect)
}

// Crop1O Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
// 
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
// 
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object 
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
// 
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`, 
// in which case the values are duplicated to a local object.
func (self *TilemapLayer) Crop1O(rect *Rectangle, copy bool) {
    self.Object.Call("crop", rect, copy)
}

// CropI Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
// 
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
// 
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object 
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
// 
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`, 
// in which case the values are duplicated to a local object.
func (self *TilemapLayer) CropI(args ...interface{}) {
    self.Object.Call("crop", args)
}

// UpdateCrop If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *TilemapLayer) UpdateCrop() {
    self.Object.Call("updateCrop")
}

// UpdateCropI If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *TilemapLayer) UpdateCropI(args ...interface{}) {
    self.Object.Call("updateCrop", args)
}

// Revive Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *TilemapLayer) Revive() *DisplayObject{
    return &DisplayObject{self.Object.Call("revive")}
}

// Revive1O Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *TilemapLayer) Revive1O(health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", health)}
}

// ReviveI Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *TilemapLayer) ReviveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", args)}
}

// Kill Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *TilemapLayer) Kill() *DisplayObject{
    return &DisplayObject{self.Object.Call("kill")}
}

// KillI Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *TilemapLayer) KillI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("kill", args)}
}

// LoadTexture Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
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
func (self *TilemapLayer) LoadTexture(key interface{}) {
    self.Object.Call("loadTexture", key)
}

// LoadTexture1O Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
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
func (self *TilemapLayer) LoadTexture1O(key interface{}, frame interface{}) {
    self.Object.Call("loadTexture", key, frame)
}

// LoadTexture2O Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
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
func (self *TilemapLayer) LoadTexture2O(key interface{}, frame interface{}, stopAnimation bool) {
    self.Object.Call("loadTexture", key, frame, stopAnimation)
}

// LoadTextureI Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
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
func (self *TilemapLayer) LoadTextureI(args ...interface{}) {
    self.Object.Call("loadTexture", args)
}

// SetFrame Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *TilemapLayer) SetFrame(frame *Frame) {
    self.Object.Call("setFrame", frame)
}

// SetFrameI Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *TilemapLayer) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// ResizeFrame Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *TilemapLayer) ResizeFrame(parent interface{}, width int, height int) {
    self.Object.Call("resizeFrame", parent, width, height)
}

// ResizeFrameI Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *TilemapLayer) ResizeFrameI(args ...interface{}) {
    self.Object.Call("resizeFrame", args)
}

// ResetFrame Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *TilemapLayer) ResetFrame() {
    self.Object.Call("resetFrame")
}

// ResetFrameI Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *TilemapLayer) ResetFrameI(args ...interface{}) {
    self.Object.Call("resetFrame", args)
}

// Overlap Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *TilemapLayer) Overlap(displayObject interface{}) bool{
    return self.Object.Call("overlap", displayObject).Bool()
}

// OverlapI Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *TilemapLayer) OverlapI(args ...interface{}) bool{
    return self.Object.Call("overlap", args).Bool()
}

// Reset Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *TilemapLayer) Reset(x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", x, y)}
}

// Reset1O Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *TilemapLayer) Reset1O(x int, y int, health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", x, y, health)}
}

// ResetI Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *TilemapLayer) ResetI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", args)}
}

// CheckTransform Adjust scaling limits, if set, to this Game Object.
func (self *TilemapLayer) CheckTransform(wt *Matrix) {
    self.Object.Call("checkTransform", wt)
}

// CheckTransformI Adjust scaling limits, if set, to this Game Object.
func (self *TilemapLayer) CheckTransformI(args ...interface{}) {
    self.Object.Call("checkTransform", args)
}

// SetScaleMinMax Sets the scaleMin and scaleMax values. These values are used to limit how far this Game Object will scale based on its parent.
// 
// For example if this Game Object has a `minScale` value of 1 and its parent has a `scale` value of 0.5, the 0.5 will be ignored 
// and the scale value of 1 will be used, as the parents scale is lower than the minimum scale this Game Object should adhere to.
// 
// By setting these values you can carefully control how Game Objects deal with responsive scaling.
// 
// If only one parameter is given then that value will be used for both scaleMin and scaleMax:
// `setScaleMinMax(1)` = scaleMin.x, scaleMin.y, scaleMax.x and scaleMax.y all = 1
// 
// If only two parameters are given the first is set as scaleMin.x and y and the second as scaleMax.x and y:
// `setScaleMinMax(0.5, 2)` = scaleMin.x and y = 0.5 and scaleMax.x and y = 2
// 
// If you wish to set `scaleMin` with different values for x and y then either modify Game Object.scaleMin directly, 
// or pass `null` for the `maxX` and `maxY` parameters.
// 
// Call `setScaleMinMax(null)` to clear all previously set values.
func (self *TilemapLayer) SetScaleMinMax(minX interface{}, minY interface{}, maxX interface{}, maxY interface{}) {
    self.Object.Call("setScaleMinMax", minX, minY, maxX, maxY)
}

// SetScaleMinMaxI Sets the scaleMin and scaleMax values. These values are used to limit how far this Game Object will scale based on its parent.
// 
// For example if this Game Object has a `minScale` value of 1 and its parent has a `scale` value of 0.5, the 0.5 will be ignored 
// and the scale value of 1 will be used, as the parents scale is lower than the minimum scale this Game Object should adhere to.
// 
// By setting these values you can carefully control how Game Objects deal with responsive scaling.
// 
// If only one parameter is given then that value will be used for both scaleMin and scaleMax:
// `setScaleMinMax(1)` = scaleMin.x, scaleMin.y, scaleMax.x and scaleMax.y all = 1
// 
// If only two parameters are given the first is set as scaleMin.x and y and the second as scaleMax.x and y:
// `setScaleMinMax(0.5, 2)` = scaleMin.x and y = 0.5 and scaleMax.x and y = 2
// 
// If you wish to set `scaleMin` with different values for x and y then either modify Game Object.scaleMin directly, 
// or pass `null` for the `maxX` and `maxY` parameters.
// 
// Call `setScaleMinMax(null)` to clear all previously set values.
func (self *TilemapLayer) SetScaleMinMaxI(args ...interface{}) {
    self.Object.Call("setScaleMinMax", args)
}

