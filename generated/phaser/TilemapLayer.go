// Automatic generation for Phaser.TilemapLayer
// generated file TilemapLayer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// A TilemapLayer is a Phaser.Image/Sprite that renders a specific TileLayer of a Tilemap.
// 
// Since a TilemapLayer is a Sprite it can be moved around the display, added to other groups or display objects, etc.
// 
// By default TilemapLayers have fixedToCamera set to `true`. Changing this will break Camera follow and scrolling behavior.
type TilemapLayer struct {
    *js.Object
}


// The Tilemap to which this layer is bound.
func (self *TilemapLayer) GetMap() *Tilemap{
    return &Tilemap{self.Get("map")}
}

// The Tilemap to which this layer is bound.
func (self *TilemapLayer) SetMap(member *Tilemap) {
    self.Set("map", member)
}

// The index of this layer within the Tilemap.
func (self *TilemapLayer) GetIndex() float64{
    return self.Get("index").Float()
}

// The index of this layer within the Tilemap.
func (self *TilemapLayer) SetIndex(member float64) {
    self.Set("index", member)
}

// The layer object within the Tilemap that this layer represents.
func (self *TilemapLayer) GetLayer() interface{}{
    return self.Get("layer")
}

// The layer object within the Tilemap that this layer represents.
func (self *TilemapLayer) SetLayer(member interface{}) {
    self.Set("layer", member)
}

// The canvas to which this TilemapLayer draws.
func (self *TilemapLayer) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("canvas"))
}

// The canvas to which this TilemapLayer draws.
func (self *TilemapLayer) SetCanvas(member dom.HTMLCanvasElement) {
    self.Set("canvas", member)
}

// The const type of this object.
func (self *TilemapLayer) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *TilemapLayer) SetType(member float64) {
    self.Set("type", member)
}

// The const physics body type of this object.
func (self *TilemapLayer) GetPhysicsType() float64{
    return self.Get("physicsType").Float()
}

// The const physics body type of this object.
func (self *TilemapLayer) SetPhysicsType(member float64) {
    self.Set("physicsType", member)
}

// Settings that control standard (non-diagnostic) rendering.
func (self *TilemapLayer) GetRenderSettings() interface{}{
    return self.Get("renderSettings")
}

// Settings that control standard (non-diagnostic) rendering.
func (self *TilemapLayer) SetRenderSettings(member interface{}) {
    self.Set("renderSettings", member)
}

// Enable an additional "debug rendering" pass to display collision information.
func (self *TilemapLayer) GetDebug() bool{
    return self.Get("debug").Bool()
}

// Enable an additional "debug rendering" pass to display collision information.
func (self *TilemapLayer) SetDebug(member bool) {
    self.Set("debug", member)
}

// Controls if the core game loop and physics update this game object or not.
func (self *TilemapLayer) GetExists() bool{
    return self.Get("exists").Bool()
}

// Controls if the core game loop and physics update this game object or not.
func (self *TilemapLayer) SetExists(member bool) {
    self.Set("exists", member)
}

// Settings used for debugging and diagnostics.
func (self *TilemapLayer) GetDebugSettings() interface{}{
    return self.Get("debugSettings")
}

// Settings used for debugging and diagnostics.
func (self *TilemapLayer) SetDebugSettings(member interface{}) {
    self.Set("debugSettings", member)
}

// Speed at which this layer scrolls horizontally, relative to the camera (e.g. scrollFactorX of 0.5 scrolls half as quickly as the 'normal' camera-locked layers do).
func (self *TilemapLayer) GetScrollFactorX() float64{
    return self.Get("scrollFactorX").Float()
}

// Speed at which this layer scrolls horizontally, relative to the camera (e.g. scrollFactorX of 0.5 scrolls half as quickly as the 'normal' camera-locked layers do).
func (self *TilemapLayer) SetScrollFactorX(member float64) {
    self.Set("scrollFactorX", member)
}

// Speed at which this layer scrolls vertically, relative to the camera (e.g. scrollFactorY of 0.5 scrolls half as quickly as the 'normal' camera-locked layers do)
func (self *TilemapLayer) GetScrollFactorY() float64{
    return self.Get("scrollFactorY").Float()
}

// Speed at which this layer scrolls vertically, relative to the camera (e.g. scrollFactorY of 0.5 scrolls half as quickly as the 'normal' camera-locked layers do)
func (self *TilemapLayer) SetScrollFactorY(member float64) {
    self.Set("scrollFactorY", member)
}

// If true tiles will be force rendered, even if such is not believed to be required.
func (self *TilemapLayer) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// If true tiles will be force rendered, even if such is not believed to be required.
func (self *TilemapLayer) SetDirty(member bool) {
    self.Set("dirty", member)
}

// When ray-casting against tiles this is the number of steps it will jump. For larger tile sizes you can increase this to improve performance.
func (self *TilemapLayer) GetRayStepRate() int{
    return self.Get("rayStepRate").Int()
}

// When ray-casting against tiles this is the number of steps it will jump. For larger tile sizes you can increase this to improve performance.
func (self *TilemapLayer) SetRayStepRate(member int) {
    self.Set("rayStepRate", member)
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TilemapLayer) GetAnchor() *Point{
    return &Point{self.Get("anchor")}
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *TilemapLayer) SetAnchor(member *Point) {
    self.Set("anchor", member)
}

// The texture that the sprite is using
func (self *TilemapLayer) GetTexture() *Texture{
    return &Texture{self.Get("texture")}
}

// The texture that the sprite is using
func (self *TilemapLayer) SetTexture(member *Texture) {
    self.Set("texture", member)
}

// The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *TilemapLayer) GetTint() float64{
    return self.Get("tint").Float()
}

// The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *TilemapLayer) SetTint(member float64) {
    self.Set("tint", member)
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TilemapLayer) GetTintedTexture() *Canvas{
    return &Canvas{self.Get("tintedTexture")}
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *TilemapLayer) SetTintedTexture(member *Canvas) {
    self.Set("tintedTexture", member)
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *TilemapLayer) GetBlendMode() float64{
    return self.Get("blendMode").Float()
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *TilemapLayer) SetBlendMode(member float64) {
    self.Set("blendMode", member)
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *TilemapLayer) GetShader() *AbstractFilter{
    return &AbstractFilter{self.Get("shader")}
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *TilemapLayer) SetShader(member *AbstractFilter) {
    self.Set("shader", member)
}

// The width of the sprite, setting this will actually modify the scale to achieve the value set
func (self *TilemapLayer) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the sprite, setting this will actually modify the scale to achieve the value set
func (self *TilemapLayer) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the sprite, setting this will actually modify the scale to achieve the value set
func (self *TilemapLayer) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the sprite, setting this will actually modify the scale to achieve the value set
func (self *TilemapLayer) SetHeight(member float64) {
    self.Set("height", member)
}

// [read-only] The array of children of this container.
func (self *TilemapLayer) GetChildren() []DisplayObject{
	array := self.Get("children")
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// [read-only] The array of children of this container.
func (self *TilemapLayer) SetChildren(member []DisplayObject) {
    self.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TilemapLayer) GetIgnoreChildInput() bool{
    return self.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *TilemapLayer) SetIgnoreChildInput(member bool) {
    self.Set("ignoreChildInput", member)
}

// A reference to the currently running Game.
func (self *TilemapLayer) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *TilemapLayer) SetGame(member *Game) {
    self.Set("game", member)
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *TilemapLayer) GetName() string{
    return self.Get("name").String()
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *TilemapLayer) SetName(member string) {
    self.Set("name", member)
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *TilemapLayer) GetData() interface{}{
    return self.Get("data")
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *TilemapLayer) SetData(member interface{}) {
    self.Set("data", member)
}

// The components this Game Object has installed.
func (self *TilemapLayer) GetComponents() interface{}{
    return self.Get("components")
}

// The components this Game Object has installed.
func (self *TilemapLayer) SetComponents(member interface{}) {
    self.Set("components", member)
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *TilemapLayer) GetZ() float64{
    return self.Get("z").Float()
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *TilemapLayer) SetZ(member float64) {
    self.Set("z", member)
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *TilemapLayer) GetEvents() *Events{
    return &Events{self.Get("events")}
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *TilemapLayer) SetEvents(member *Events) {
    self.Set("events", member)
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *TilemapLayer) GetAnimations() *AnimationManager{
    return &AnimationManager{self.Get("animations")}
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *TilemapLayer) SetAnimations(member *AnimationManager) {
    self.Set("animations", member)
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *TilemapLayer) GetKey() interface{}{
    return self.Get("key")
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *TilemapLayer) SetKey(member interface{}) {
    self.Set("key", member)
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *TilemapLayer) GetWorld() *Point{
    return &Point{self.Get("world")}
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *TilemapLayer) SetWorld(member *Point) {
    self.Set("world", member)
}

// The position the Game Object was located in the previous frame.
func (self *TilemapLayer) GetPreviousPosition() *Point{
    return &Point{self.Get("previousPosition")}
}

// The position the Game Object was located in the previous frame.
func (self *TilemapLayer) SetPreviousPosition(member *Point) {
    self.Set("previousPosition", member)
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *TilemapLayer) GetPreviousRotation() float64{
    return self.Get("previousRotation").Float()
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *TilemapLayer) SetPreviousRotation(member float64) {
    self.Set("previousRotation", member)
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *TilemapLayer) GetRenderOrderID() float64{
    return self.Get("renderOrderID").Float()
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *TilemapLayer) SetRenderOrderID(member float64) {
    self.Set("renderOrderID", member)
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *TilemapLayer) GetFresh() bool{
    return self.Get("fresh").Bool()
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *TilemapLayer) SetFresh(member bool) {
    self.Set("fresh", member)
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *TilemapLayer) GetPendingDestroy() bool{
    return self.Get("pendingDestroy").Bool()
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *TilemapLayer) SetPendingDestroy(member bool) {
    self.Set("pendingDestroy", member)
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
func (self *TilemapLayer) GetAngle() float64{
    return self.Get("angle").Float()
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
func (self *TilemapLayer) SetAngle(member float64) {
    self.Set("angle", member)
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *TilemapLayer) GetAutoCull() bool{
    return self.Get("autoCull").Bool()
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *TilemapLayer) SetAutoCull(member bool) {
    self.Set("autoCull", member)
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *TilemapLayer) GetInCamera() bool{
    return self.Get("inCamera").Bool()
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *TilemapLayer) SetInCamera(member bool) {
    self.Set("inCamera", member)
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *TilemapLayer) GetOffsetX() float64{
    return self.Get("offsetX").Float()
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *TilemapLayer) SetOffsetX(member float64) {
    self.Set("offsetX", member)
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *TilemapLayer) GetOffsetY() float64{
    return self.Get("offsetY").Float()
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *TilemapLayer) SetOffsetY(member float64) {
    self.Set("offsetY", member)
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *TilemapLayer) GetCenterX() float64{
    return self.Get("centerX").Float()
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *TilemapLayer) SetCenterX(member float64) {
    self.Set("centerX", member)
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *TilemapLayer) GetCenterY() float64{
    return self.Get("centerY").Float()
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *TilemapLayer) SetCenterY(member float64) {
    self.Set("centerY", member)
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *TilemapLayer) GetLeft() float64{
    return self.Get("left").Float()
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *TilemapLayer) SetLeft(member float64) {
    self.Set("left", member)
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *TilemapLayer) GetRight() float64{
    return self.Get("right").Float()
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *TilemapLayer) SetRight(member float64) {
    self.Set("right", member)
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *TilemapLayer) GetTop() float64{
    return self.Get("top").Float()
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *TilemapLayer) SetTop(member float64) {
    self.Set("top", member)
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *TilemapLayer) GetBottom() float64{
    return self.Get("bottom").Float()
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *TilemapLayer) SetBottom(member float64) {
    self.Set("bottom", member)
}

// The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *TilemapLayer) GetCropRect() *Rectangle{
    return &Rectangle{self.Get("cropRect")}
}

// The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *TilemapLayer) SetCropRect(member *Rectangle) {
    self.Set("cropRect", member)
}

// Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *TilemapLayer) GetDeltaX() float64{
    return self.Get("deltaX").Float()
}

// Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *TilemapLayer) SetDeltaX(member float64) {
    self.Set("deltaX", member)
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *TilemapLayer) GetDeltaY() float64{
    return self.Get("deltaY").Float()
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *TilemapLayer) SetDeltaY(member float64) {
    self.Set("deltaY", member)
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *TilemapLayer) GetDeltaZ() float64{
    return self.Get("deltaZ").Float()
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *TilemapLayer) SetDeltaZ(member float64) {
    self.Set("deltaZ", member)
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *TilemapLayer) GetDestroyPhase() bool{
    return self.Get("destroyPhase").Bool()
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *TilemapLayer) SetDestroyPhase(member bool) {
    self.Set("destroyPhase", member)
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
func (self *TilemapLayer) GetFixedToCamera() bool{
    return self.Get("fixedToCamera").Bool()
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
func (self *TilemapLayer) SetFixedToCamera(member bool) {
    self.Set("fixedToCamera", member)
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *TilemapLayer) GetCameraOffset() *Point{
    return &Point{self.Get("cameraOffset")}
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *TilemapLayer) SetCameraOffset(member *Point) {
    self.Set("cameraOffset", member)
}

// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *TilemapLayer) GetHealth() float64{
    return self.Get("health").Float()
}

// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *TilemapLayer) SetHealth(member float64) {
    self.Set("health", member)
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *TilemapLayer) GetMaxHealth() float64{
    return self.Get("maxHealth").Float()
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *TilemapLayer) SetMaxHealth(member float64) {
    self.Set("maxHealth", member)
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *TilemapLayer) GetDamage() interface{}{
    return self.Get("damage")
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *TilemapLayer) SetDamage(member interface{}) {
    self.Set("damage", member)
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *TilemapLayer) GetSetHealth() interface{}{
    return self.Get("setHealth")
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *TilemapLayer) SetSetHealth(member interface{}) {
    self.Set("setHealth", member)
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *TilemapLayer) GetHeal() interface{}{
    return self.Get("heal")
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *TilemapLayer) SetHeal(member interface{}) {
    self.Set("heal", member)
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *TilemapLayer) GetInput() interface{}{
    return self.Get("input")
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *TilemapLayer) SetInput(member interface{}) {
    self.Set("input", member)
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
func (self *TilemapLayer) GetInputEnabled() bool{
    return self.Get("inputEnabled").Bool()
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
func (self *TilemapLayer) SetInputEnabled(member bool) {
    self.Set("inputEnabled", member)
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
func (self *TilemapLayer) GetCheckWorldBounds() bool{
    return self.Get("checkWorldBounds").Bool()
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
func (self *TilemapLayer) SetCheckWorldBounds(member bool) {
    self.Set("checkWorldBounds", member)
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *TilemapLayer) GetOutOfBoundsKill() bool{
    return self.Get("outOfBoundsKill").Bool()
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *TilemapLayer) SetOutOfBoundsKill(member bool) {
    self.Set("outOfBoundsKill", member)
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *TilemapLayer) GetOutOfCameraBoundsKill() bool{
    return self.Get("outOfCameraBoundsKill").Bool()
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *TilemapLayer) SetOutOfCameraBoundsKill(member bool) {
    self.Set("outOfCameraBoundsKill", member)
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *TilemapLayer) GetInWorld() bool{
    return self.Get("inWorld").Bool()
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *TilemapLayer) SetInWorld(member bool) {
    self.Set("inWorld", member)
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *TilemapLayer) GetAlive() bool{
    return self.Get("alive").Bool()
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *TilemapLayer) SetAlive(member bool) {
    self.Set("alive", member)
}

// The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *TilemapLayer) GetLifespan() float64{
    return self.Get("lifespan").Float()
}

// The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *TilemapLayer) SetLifespan(member float64) {
    self.Set("lifespan", member)
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
func (self *TilemapLayer) GetFrame() int{
    return self.Get("frame").Int()
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
func (self *TilemapLayer) SetFrame(member int) {
    self.Set("frame", member)
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
func (self *TilemapLayer) GetFrameName() string{
    return self.Get("frameName").String()
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
func (self *TilemapLayer) SetFrameName(member string) {
    self.Set("frameName", member)
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
func (self *TilemapLayer) GetBody() interface{}{
    return self.Get("body")
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
func (self *TilemapLayer) SetBody(member interface{}) {
    self.Set("body", member)
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *TilemapLayer) GetX() float64{
    return self.Get("x").Float()
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *TilemapLayer) SetX(member float64) {
    self.Set("x", member)
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *TilemapLayer) GetY() float64{
    return self.Get("y").Float()
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *TilemapLayer) SetY(member float64) {
    self.Set("y", member)
}

// The callback that will apply any scale limiting to the worldTransform.
func (self *TilemapLayer) SetTransformCallback(member func(...interface{})) {
    self.Set("transformCallback", member)
}

// The context under which `transformCallback` is called.
func (self *TilemapLayer) GetTransformCallbackContext() interface{}{
    return self.Get("transformCallbackContext")
}

// The context under which `transformCallback` is called.
func (self *TilemapLayer) SetTransformCallbackContext(member interface{}) {
    self.Set("transformCallbackContext", member)
}

// The minimum scale this Game Object will scale down to.
// 
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
// 
// Set it to `null` to remove the limit.
func (self *TilemapLayer) GetScaleMin() *Point{
    return &Point{self.Get("scaleMin")}
}

// The minimum scale this Game Object will scale down to.
// 
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
// 
// Set it to `null` to remove the limit.
func (self *TilemapLayer) SetScaleMin(member *Point) {
    self.Set("scaleMin", member)
}

// The maximum scale this Game Object will scale up to. 
// 
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
// 
// Set it to `null` to remove the limit.
func (self *TilemapLayer) GetScaleMax() *Point{
    return &Point{self.Get("scaleMax")}
}

// The maximum scale this Game Object will scale up to. 
// 
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
// 
// Set it to `null` to remove the limit.
func (self *TilemapLayer) SetScaleMax(member *Point) {
    self.Set("scaleMax", member)
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *TilemapLayer) GetSmoothed() bool{
    return self.Get("smoothed").Bool()
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *TilemapLayer) SetSmoothed(member bool) {
    self.Set("smoothed", member)
}



// Create if needed (and return) a shared copy canvas that is shared across all TilemapLayers.
// 
// Code that uses the canvas is responsible to ensure the dimensions and save/restore state as appropriate.
func (self *TilemapLayer) EnsureSharedCopyCanvasI(args ...interface{}) {
    self.Call("ensureSharedCopyCanvas", args)
}

// Automatically called by World.preUpdate.
func (self *TilemapLayer) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Automatically called by World.postUpdate. Handles cache updates.
func (self *TilemapLayer) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}

// Automatically called by the Canvas Renderer.
// Overrides the Sprite._renderCanvas function.
func (self *TilemapLayer) _renderCanvasI(args ...interface{}) {
    self.Call("_renderCanvas", args)
}

// Automatically called by the Canvas Renderer.
// Overrides the Sprite._renderWebGL function.
func (self *TilemapLayer) _renderWebGLI(args ...interface{}) {
    self.Call("_renderWebGL", args)
}

// Destroys this TilemapLayer.
func (self *TilemapLayer) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Resizes the internal canvas and texture frame used by this TilemapLayer.
// 
// This is an expensive call, so don't bind it to a window resize event! But instead call it at carefully
// selected times.
// 
// Be aware that no validation of the new sizes takes place and the current map scroll coordinates are not
// modified either. You will have to handle both of these things from your game code if required.
func (self *TilemapLayer) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// Sets the world size to match the size of this layer.
func (self *TilemapLayer) ResizeWorldI(args ...interface{}) {
    self.Call("resizeWorld", args)
}

// Take an x coordinate that doesn't account for scrollFactorX and 'fix' it into a scrolled local space.
func (self *TilemapLayer) _fixXI(args ...interface{}) float64{
    return self.Call("_fixX", args).Float()
}

// Take an x coordinate that _does_ account for scrollFactorX and 'unfix' it back to camera space.
func (self *TilemapLayer) _unfixXI(args ...interface{}) float64{
    return self.Call("_unfixX", args).Float()
}

// Take a y coordinate that doesn't account for scrollFactorY and 'fix' it into a scrolled local space.
func (self *TilemapLayer) _fixYI(args ...interface{}) float64{
    return self.Call("_fixY", args).Float()
}

// Take a y coordinate that _does_ account for scrollFactorY and 'unfix' it back to camera space.
func (self *TilemapLayer) _unfixYI(args ...interface{}) float64{
    return self.Call("_unfixY", args).Float()
}

// Convert a pixel value to a tile coordinate.
func (self *TilemapLayer) GetTileXI(args ...interface{}) int{
    return self.Call("getTileX", args).Int()
}

// Convert a pixel value to a tile coordinate.
func (self *TilemapLayer) GetTileYI(args ...interface{}) int{
    return self.Call("getTileY", args).Int()
}

// Convert a pixel coordinate to a tile coordinate.
func (self *TilemapLayer) GetTileXYI(args ...interface{}) interface{}{
    return self.Call("getTileXY", args)
}

// Gets all tiles that intersect with the given line.
func (self *TilemapLayer) GetRayCastTilesI(args ...interface{}) []Tile{
	array := self.Call("getRayCastTiles", args)
	length := array.Length()
	out := make([]Tile, length, length)
	for i := 0; i < length; i++ {
		
			out[i] = Tile{array.Index(i)}
	}
	return out
}

// Get all tiles that exist within the given area, defined by the top-left corner, width and height. Values given are in pixels, not tiles.
func (self *TilemapLayer) GetTilesI(args ...interface{}) []Tile{
	array := self.Call("getTiles", args)
	length := array.Length()
	out := make([]Tile, length, length)
	for i := 0; i < length; i++ {
		
			out[i] = Tile{array.Index(i)}
	}
	return out
}

// Returns the appropriate tileset for the index, updating the internal cache as required.
// This should only be called if `tilesets[index]` evaluates to undefined.
func (self *TilemapLayer) ResolveTilesetI(args ...interface{}) interface{}{
    return self.Call("resolveTileset", args)
}

// The TilemapLayer caches tileset look-ups.
// 
// Call this method of clear the cache if tilesets have been added or updated after the layer has been rendered.
func (self *TilemapLayer) ResetTilesetCacheI(args ...interface{}) {
    self.Call("resetTilesetCache", args)
}

// This method will set the scale of the tilemap as well as update the underlying block data of this layer.
func (self *TilemapLayer) SetScaleI(args ...interface{}) {
    self.Call("setScale", args)
}

// Shifts the contents of the canvas - does extra math so that different browsers agree on the result.
// 
// The specified (x/y) will be shifted to (0,0) after the copy and the newly exposed canvas area will need to be filled in.
func (self *TilemapLayer) ShiftCanvasI(args ...interface{}) {
    self.Call("shiftCanvas", args)
}

// Render tiles in the given area given by the virtual tile coordinates biased by the given scroll factor.
// This will constrain the tile coordinates based on wrapping but not physical coordinates.
func (self *TilemapLayer) RenderRegionI(args ...interface{}) {
    self.Call("renderRegion", args)
}

// Shifts the canvas and render damaged edge tiles.
func (self *TilemapLayer) RenderDeltaScrollI(args ...interface{}) {
    self.Call("renderDeltaScroll", args)
}

// Clear and render the entire canvas.
func (self *TilemapLayer) RenderFullI(args ...interface{}) {
    self.Call("renderFull", args)
}

// Renders the tiles to the layer canvas and pushes to the display.
func (self *TilemapLayer) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// Renders a debug overlay on-top of the canvas. Called automatically by render when `debug` is true.
// 
// See `debugSettings` for assorted configuration options.
func (self *TilemapLayer) RenderDebugI(args ...interface{}) {
    self.Call("renderDebug", args)
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *TilemapLayer) SetTextureI(args ...interface{}) {
    self.Call("setTexture", args)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *TilemapLayer) OnTextureUpdateI(args ...interface{}) {
    self.Call("onTextureUpdate", args)
}

// Returns the bounds of the Sprite as a rectangle.
// The bounds calculation takes the worldTransform into account.
// 
// It is important to note that the transform is not updated when you call this method.
// So if this Sprite is the child of a Display Object which has had its transform
// updated since the last render pass, those changes will not yet have been applied
// to this Sprites worldTransform. If you need to ensure that all parent transforms
// are factored into this getBounds operation then you should call `updateTransform`
// on the root most object in this Sprites display list first.
func (self *TilemapLayer) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getBounds", args)}
}

// Adds a child to the container.
func (self *TilemapLayer) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *TilemapLayer) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *TilemapLayer) SwapChildrenI(args ...interface{}) {
    self.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *TilemapLayer) GetChildIndexI(args ...interface{}) float64{
    return self.Call("getChildIndex", args).Float()
}

// Changes the position of an existing child in the display object container
func (self *TilemapLayer) SetChildIndexI(args ...interface{}) {
    self.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *TilemapLayer) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *TilemapLayer) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *TilemapLayer) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *TilemapLayer) RemoveChildrenI(args ...interface{}) {
    self.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *TilemapLayer) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *TilemapLayer) SetStageReferenceI(args ...interface{}) {
    self.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *TilemapLayer) RemoveStageReferenceI(args ...interface{}) {
    self.Call("removeStageReference", args)
}

// Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *TilemapLayer) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *TilemapLayer) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Call("play", args)}
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
func (self *TilemapLayer) AlignInI(args ...interface{}) interface{}{
    return self.Call("alignIn", args)
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
func (self *TilemapLayer) AlignToI(args ...interface{}) interface{}{
    return self.Call("alignTo", args)
}

// Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) BringToTopI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("bringToTop", args)}
}

// Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) SendToBackI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("sendToBack", args)}
}

// Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) MoveUpI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("moveUp", args)}
}

// Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *TilemapLayer) MoveDownI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("moveDown", args)}
}

// Crop allows you to crop the texture being used to display this Game Object.
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
    self.Call("crop", args)
}

// If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *TilemapLayer) UpdateCropI(args ...interface{}) {
    self.Call("updateCrop", args)
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *TilemapLayer) ReviveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("revive", args)}
}

// Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *TilemapLayer) KillI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("kill", args)}
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
func (self *TilemapLayer) LoadTextureI(args ...interface{}) {
    self.Call("loadTexture", args)
}

// Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *TilemapLayer) SetFrameI(args ...interface{}) {
    self.Call("setFrame", args)
}

// Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *TilemapLayer) ResizeFrameI(args ...interface{}) {
    self.Call("resizeFrame", args)
}

// Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *TilemapLayer) ResetFrameI(args ...interface{}) {
    self.Call("resetFrame", args)
}

// Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *TilemapLayer) OverlapI(args ...interface{}) bool{
    return self.Call("overlap", args).Bool()
}

// Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *TilemapLayer) ResetI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("reset", args)}
}

// Adjust scaling limits, if set, to this Game Object.
func (self *TilemapLayer) CheckTransformI(args ...interface{}) {
    self.Call("checkTransform", args)
}

// Sets the scaleMin and scaleMax values. These values are used to limit how far this Game Object will scale based on its parent.
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
    self.Call("setScaleMinMax", args)
}
