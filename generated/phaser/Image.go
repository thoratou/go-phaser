// Package phaser Automatic generation for Phaser.Image
// generated file Image.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Image An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
type Image struct {
    *js.Object
}

// NewImage An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func NewImage(game *Game) *Image {
    return &Image{js.Global.Get("Phaser").Get("Image").New(game)}
}
// NewImage1O An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func NewImage1O(game *Game, x int) *Image {
    return &Image{js.Global.Get("Phaser").Get("Image").New(game, x)}
}
// NewImage2O An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func NewImage2O(game *Game, x int, y int) *Image {
    return &Image{js.Global.Get("Phaser").Get("Image").New(game, x, y)}
}
// NewImage3O An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func NewImage3O(game *Game, x int, y int, key interface{}) *Image {
    return &Image{js.Global.Get("Phaser").Get("Image").New(game, x, y, key)}
}
// NewImage4O An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func NewImage4O(game *Game, x int, y int, key interface{}, frame interface{}) *Image {
    return &Image{js.Global.Get("Phaser").Get("Image").New(game, x, y, key, frame)}
}
// NewImageI An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func NewImageI(args ...interface{}) *Image {
    return &Image{js.Global.Get("Phaser").Get("Image").New(args)}
}



// Image Binding conversion method to Image point 
func ToImage(jsStruct interface{}) *Image {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Image{Object: object}
	}
	return nil
}



// Type The const type of this object.
func (self *Image) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *Image) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Anchor The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *Image) Anchor() *Point{
    return &Point{self.Object.Get("anchor")}
}

// SetAnchorA The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *Image) SetAnchorA(member *Point) {
    self.Object.Set("anchor", member)
}

// Texture The texture that the sprite is using
func (self *Image) Texture() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// SetTextureA The texture that the sprite is using
func (self *Image) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// Tint The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *Image) Tint() int{
    return self.Object.Get("tint").Int()
}

// SetTintA The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *Image) SetTintA(member int) {
    self.Object.Set("tint", member)
}

// TintedTexture A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *Image) TintedTexture() *Canvas{
    return &Canvas{self.Object.Get("tintedTexture")}
}

// SetTintedTextureA A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *Image) SetTintedTextureA(member *Canvas) {
    self.Object.Set("tintedTexture", member)
}

// BlendMode The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *Image) BlendMode() int{
    return self.Object.Get("blendMode").Int()
}

// SetBlendModeA The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *Image) SetBlendModeA(member int) {
    self.Object.Set("blendMode", member)
}

// Shader The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *Image) Shader() *AbstractFilter{
    return &AbstractFilter{self.Object.Get("shader")}
}

// SetShaderA The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *Image) SetShaderA(member *AbstractFilter) {
    self.Object.Set("shader", member)
}

// Exists Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *Image) Exists() bool{
    return self.Object.Get("exists").Bool()
}

// SetExistsA Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *Image) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// Width The width of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Image) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Image) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Image) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Image) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Children [read-only] The array of children of this container.
func (self *Image) Children() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// SetChildrenA [read-only] The array of children of this container.
func (self *Image) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// IgnoreChildInput If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Image) IgnoreChildInput() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// SetIgnoreChildInputA If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Image) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// Game A reference to the currently running Game.
func (self *Image) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Image) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Name A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Image) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Image) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Data An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Image) Data() interface{}{
    return self.Object.Get("data")
}

// SetDataA An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Image) SetDataA(member interface{}) {
    self.Object.Set("data", member)
}

// Components The components this Game Object has installed.
func (self *Image) Components() interface{}{
    return self.Object.Get("components")
}

// SetComponentsA The components this Game Object has installed.
func (self *Image) SetComponentsA(member interface{}) {
    self.Object.Set("components", member)
}

// Z The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Image) Z() int{
    return self.Object.Get("z").Int()
}

// SetZA The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Image) SetZA(member int) {
    self.Object.Set("z", member)
}

// Events All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Image) Events() *Events{
    return &Events{self.Object.Get("events")}
}

// SetEventsA All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Image) SetEventsA(member *Events) {
    self.Object.Set("events", member)
}

// Animations If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Image) Animations() *AnimationManager{
    return &AnimationManager{self.Object.Get("animations")}
}

// SetAnimationsA If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Image) SetAnimationsA(member *AnimationManager) {
    self.Object.Set("animations", member)
}

// Key The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Image) Key() interface{}{
    return self.Object.Get("key")
}

// SetKeyA The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Image) SetKeyA(member interface{}) {
    self.Object.Set("key", member)
}

// World The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Image) World() *Point{
    return &Point{self.Object.Get("world")}
}

// SetWorldA The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Image) SetWorldA(member *Point) {
    self.Object.Set("world", member)
}

// Debug A debug flag designed for use with `Game.enableStep`.
func (self *Image) Debug() bool{
    return self.Object.Get("debug").Bool()
}

// SetDebugA A debug flag designed for use with `Game.enableStep`.
func (self *Image) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// PreviousPosition The position the Game Object was located in the previous frame.
func (self *Image) PreviousPosition() *Point{
    return &Point{self.Object.Get("previousPosition")}
}

// SetPreviousPositionA The position the Game Object was located in the previous frame.
func (self *Image) SetPreviousPositionA(member *Point) {
    self.Object.Set("previousPosition", member)
}

// PreviousRotation The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Image) PreviousRotation() int{
    return self.Object.Get("previousRotation").Int()
}

// SetPreviousRotationA The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Image) SetPreviousRotationA(member int) {
    self.Object.Set("previousRotation", member)
}

// RenderOrderID The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Image) RenderOrderID() int{
    return self.Object.Get("renderOrderID").Int()
}

// SetRenderOrderIDA The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Image) SetRenderOrderIDA(member int) {
    self.Object.Set("renderOrderID", member)
}

// Fresh A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Image) Fresh() bool{
    return self.Object.Get("fresh").Bool()
}

// SetFreshA A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Image) SetFreshA(member bool) {
    self.Object.Set("fresh", member)
}

// PendingDestroy A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Image) PendingDestroy() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// SetPendingDestroyA A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Image) SetPendingDestroyA(member bool) {
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
func (self *Image) Angle() int{
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
func (self *Image) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// AutoCull A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Image) AutoCull() bool{
    return self.Object.Get("autoCull").Bool()
}

// SetAutoCullA A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Image) SetAutoCullA(member bool) {
    self.Object.Set("autoCull", member)
}

// InCamera Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Image) InCamera() bool{
    return self.Object.Get("inCamera").Bool()
}

// SetInCameraA Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Image) SetInCameraA(member bool) {
    self.Object.Set("inCamera", member)
}

// OffsetX The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Image) OffsetX() int{
    return self.Object.Get("offsetX").Int()
}

// SetOffsetXA The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Image) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// OffsetY The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Image) OffsetY() int{
    return self.Object.Get("offsetY").Int()
}

// SetOffsetYA The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Image) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}

// CenterX The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Image) CenterX() int{
    return self.Object.Get("centerX").Int()
}

// SetCenterXA The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Image) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// CenterY The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Image) CenterY() int{
    return self.Object.Get("centerY").Int()
}

// SetCenterYA The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Image) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// Left The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Image) Left() int{
    return self.Object.Get("left").Int()
}

// SetLeftA The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Image) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// Right The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Image) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Image) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Top The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Image) Top() int{
    return self.Object.Get("top").Int()
}

// SetTopA The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Image) SetTopA(member int) {
    self.Object.Set("top", member)
}

// Bottom The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Image) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Image) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// CropRect The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Image) CropRect() *Rectangle{
    return &Rectangle{self.Object.Get("cropRect")}
}

// SetCropRectA The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Image) SetCropRectA(member *Rectangle) {
    self.Object.Set("cropRect", member)
}

// DestroyPhase As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Image) DestroyPhase() bool{
    return self.Object.Get("destroyPhase").Bool()
}

// SetDestroyPhaseA As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Image) SetDestroyPhaseA(member bool) {
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
func (self *Image) FixedToCamera() bool{
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
func (self *Image) SetFixedToCameraA(member bool) {
    self.Object.Set("fixedToCamera", member)
}

// CameraOffset The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Image) CameraOffset() *Point{
    return &Point{self.Object.Get("cameraOffset")}
}

// SetCameraOffsetA The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Image) SetCameraOffsetA(member *Point) {
    self.Object.Set("cameraOffset", member)
}

// Input The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Image) Input() interface{}{
    return self.Object.Get("input")
}

// SetInputA The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Image) SetInputA(member interface{}) {
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
func (self *Image) InputEnabled() bool{
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
func (self *Image) SetInputEnabledA(member bool) {
    self.Object.Set("inputEnabled", member)
}

// Alive A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Image) Alive() bool{
    return self.Object.Get("alive").Bool()
}

// SetAliveA A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Image) SetAliveA(member bool) {
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
func (self *Image) Lifespan() int{
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
func (self *Image) SetLifespanA(member int) {
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
func (self *Image) Frame() int{
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
func (self *Image) SetFrameA(member int) {
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
func (self *Image) FrameName() string{
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
func (self *Image) SetFrameNameA(member string) {
    self.Object.Set("frameName", member)
}

// Smoothed Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *Image) Smoothed() bool{
    return self.Object.Get("smoothed").Bool()
}

// SetSmoothedA Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *Image) SetSmoothedA(member bool) {
    self.Object.Set("smoothed", member)
}


// PreUpdate Automatically called by World.preUpdate.
func (self *Image) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI Automatically called by World.preUpdate.
func (self *Image) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// SetTexture Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Image) SetTexture(texture *Texture) {
    self.Object.Call("setTexture", texture)
}

// SetTexture1O Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Image) SetTexture1O(texture *Texture, destroy bool) {
    self.Object.Call("setTexture", texture, destroy)
}

// SetTextureI Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Image) SetTextureI(args ...interface{}) {
    self.Object.Call("setTexture", args)
}

// OnTextureUpdate When the texture is updated, this event will fire to update the scale and frame
func (self *Image) OnTextureUpdate(event interface{}) {
    self.Object.Call("onTextureUpdate", event)
}

// OnTextureUpdateI When the texture is updated, this event will fire to update the scale and frame
func (self *Image) OnTextureUpdateI(args ...interface{}) {
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
func (self *Image) GetBounds(matrix *Matrix) *Rectangle{
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
func (self *Image) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// _renderWebGL Renders the object using the WebGL renderer
func (self *Image) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// _renderWebGL1O Renders the object using the WebGL renderer
func (self *Image) _renderWebGL1O(renderSession *RenderSession, matrix *Matrix) {
    self.Object.Call("_renderWebGL", renderSession, matrix)
}

// _renderWebGLI Renders the object using the WebGL renderer
func (self *Image) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// _renderCanvas Renders the object using the Canvas renderer
func (self *Image) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// _renderCanvas1O Renders the object using the Canvas renderer
func (self *Image) _renderCanvas1O(renderSession *RenderSession, matrix *Matrix) {
    self.Object.Call("_renderCanvas", renderSession, matrix)
}

// _renderCanvasI Renders the object using the Canvas renderer
func (self *Image) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

// AddChild Adds a child to the container.
func (self *Image) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// AddChildI Adds a child to the container.
func (self *Image) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// AddChildAt Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Image) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// AddChildAtI Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Image) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// SwapChildren Swaps the position of 2 Display Objects within this container.
func (self *Image) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// SwapChildrenI Swaps the position of 2 Display Objects within this container.
func (self *Image) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// GetChildIndex Returns the index position of a child DisplayObject instance
func (self *Image) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// GetChildIndexI Returns the index position of a child DisplayObject instance
func (self *Image) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// SetChildIndex Changes the position of an existing child in the display object container
func (self *Image) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// SetChildIndexI Changes the position of an existing child in the display object container
func (self *Image) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// GetChildAt Returns the child at the specified index
func (self *Image) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// GetChildAtI Returns the child at the specified index
func (self *Image) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// RemoveChild Removes a child from the container.
func (self *Image) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// RemoveChildI Removes a child from the container.
func (self *Image) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// RemoveChildAt Removes a child from the specified index position.
func (self *Image) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// RemoveChildAtI Removes a child from the specified index position.
func (self *Image) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// RemoveChildren Removes all children from this container that are within the begin and end indexes.
func (self *Image) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// RemoveChildrenI Removes all children from this container that are within the begin and end indexes.
func (self *Image) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// GetLocalBounds Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Image) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// GetLocalBoundsI Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Image) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// SetStageReference Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Image) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// SetStageReferenceI Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Image) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// RemoveStageReference Removes the current stage reference from the container and all of its children.
func (self *Image) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// RemoveStageReferenceI Removes the current stage reference from the container and all of its children.
func (self *Image) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// Update Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *Image) Update() {
    self.Object.Call("update")
}

// UpdateI Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *Image) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// PostUpdate Internal method called by the World postUpdate cycle.
func (self *Image) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI Internal method called by the World postUpdate cycle.
func (self *Image) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Play Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Image) Play(name string) *Animation{
    return &Animation{self.Object.Call("play", name)}
}

// Play1O Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Image) Play1O(name string, frameRate int) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate)}
}

// Play2O Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Image) Play2O(name string, frameRate int, loop bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop)}
}

// Play3O Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Image) Play3O(name string, frameRate int, loop bool, killOnComplete bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop, killOnComplete)}
}

// PlayI Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Image) PlayI(args ...interface{}) *Animation{
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
func (self *Image) AlignIn(container interface{}) interface{}{
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
func (self *Image) AlignIn1O(container interface{}, position int) interface{}{
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
func (self *Image) AlignIn2O(container interface{}, position int, offsetX int) interface{}{
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
func (self *Image) AlignIn3O(container interface{}, position int, offsetX int, offsetY int) interface{}{
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
func (self *Image) AlignInI(args ...interface{}) interface{}{
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
func (self *Image) AlignTo(parent interface{}) interface{}{
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
func (self *Image) AlignTo1O(parent interface{}, position int) interface{}{
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
func (self *Image) AlignTo2O(parent interface{}, position int, offsetX int) interface{}{
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
func (self *Image) AlignTo3O(parent interface{}, position int, offsetX int, offsetY int) interface{}{
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
func (self *Image) AlignToI(args ...interface{}) interface{}{
    return self.Object.Call("alignTo", args)
}

// BringToTop Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) BringToTop() *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop")}
}

// BringToTopI Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) BringToTopI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop", args)}
}

// SendToBack Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) SendToBack() *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack")}
}

// SendToBackI Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) SendToBackI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack", args)}
}

// MoveUp Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) MoveUp() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp")}
}

// MoveUpI Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) MoveUpI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp", args)}
}

// MoveDown Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) MoveDown() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown")}
}

// MoveDownI Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) MoveDownI(args ...interface{}) *DisplayObject{
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
func (self *Image) Crop(rect *Rectangle) {
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
func (self *Image) Crop1O(rect *Rectangle, copy bool) {
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
func (self *Image) CropI(args ...interface{}) {
    self.Object.Call("crop", args)
}

// UpdateCrop If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *Image) UpdateCrop() {
    self.Object.Call("updateCrop")
}

// UpdateCropI If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *Image) UpdateCropI(args ...interface{}) {
    self.Object.Call("updateCrop", args)
}

// Destroy Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Image) Destroy() {
    self.Object.Call("destroy")
}

// Destroy1O Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Image) Destroy1O(destroyChildren bool) {
    self.Object.Call("destroy", destroyChildren)
}

// Destroy2O Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Image) Destroy2O(destroyChildren bool, destroyTexture bool) {
    self.Object.Call("destroy", destroyChildren, destroyTexture)
}

// DestroyI Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Image) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Revive Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Image) Revive() *DisplayObject{
    return &DisplayObject{self.Object.Call("revive")}
}

// Revive1O Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Image) Revive1O(health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", health)}
}

// ReviveI Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Image) ReviveI(args ...interface{}) *DisplayObject{
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
func (self *Image) Kill() *DisplayObject{
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
func (self *Image) KillI(args ...interface{}) *DisplayObject{
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
func (self *Image) LoadTexture(key interface{}) {
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
func (self *Image) LoadTexture1O(key interface{}, frame interface{}) {
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
func (self *Image) LoadTexture2O(key interface{}, frame interface{}, stopAnimation bool) {
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
func (self *Image) LoadTextureI(args ...interface{}) {
    self.Object.Call("loadTexture", args)
}

// SetFrame Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *Image) SetFrame(frame *Frame) {
    self.Object.Call("setFrame", frame)
}

// SetFrameI Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *Image) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// ResizeFrame Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *Image) ResizeFrame(parent interface{}, width int, height int) {
    self.Object.Call("resizeFrame", parent, width, height)
}

// ResizeFrameI Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *Image) ResizeFrameI(args ...interface{}) {
    self.Object.Call("resizeFrame", args)
}

// ResetFrame Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *Image) ResetFrame() {
    self.Object.Call("resetFrame")
}

// ResetFrameI Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *Image) ResetFrameI(args ...interface{}) {
    self.Object.Call("resetFrame", args)
}

// Overlap Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *Image) Overlap(displayObject interface{}) bool{
    return self.Object.Call("overlap", displayObject).Bool()
}

// OverlapI Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *Image) OverlapI(args ...interface{}) bool{
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
func (self *Image) Reset(x int, y int) *DisplayObject{
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
func (self *Image) Reset1O(x int, y int, health int) *DisplayObject{
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
func (self *Image) ResetI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", args)}
}

