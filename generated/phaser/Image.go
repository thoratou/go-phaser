// Automatic generation for Phaser.Image
// generated file Image.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
type Image struct {
    *js.Object
}


// The const type of this object.
func (self *Image) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *Image) SetType(member float64) {
    self.Set("type", member)
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *Image) GetAnchor() *Point{
    return &Point{self.Get("anchor")}
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *Image) SetAnchor(member *Point) {
    self.Set("anchor", member)
}

// The texture that the sprite is using
func (self *Image) GetTexture() *Texture{
    return &Texture{self.Get("texture")}
}

// The texture that the sprite is using
func (self *Image) SetTexture(member *Texture) {
    self.Set("texture", member)
}

// The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *Image) GetTint() float64{
    return self.Get("tint").Float()
}

// The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *Image) SetTint(member float64) {
    self.Set("tint", member)
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *Image) GetTintedTexture() *Canvas{
    return &Canvas{self.Get("tintedTexture")}
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *Image) SetTintedTexture(member *Canvas) {
    self.Set("tintedTexture", member)
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *Image) GetBlendMode() float64{
    return self.Get("blendMode").Float()
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *Image) SetBlendMode(member float64) {
    self.Set("blendMode", member)
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *Image) GetShader() *AbstractFilter{
    return &AbstractFilter{self.Get("shader")}
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *Image) SetShader(member *AbstractFilter) {
    self.Set("shader", member)
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *Image) GetExists() bool{
    return self.Get("exists").Bool()
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *Image) SetExists(member bool) {
    self.Set("exists", member)
}

// The width of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Image) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Image) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Image) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the sprite, setting this will actually modify the scale to achieve the value set
func (self *Image) SetHeight(member float64) {
    self.Set("height", member)
}

// [read-only] The array of children of this container.
func (self *Image) GetChildren() []DisplayObject{
	array := self.Get("children")
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// [read-only] The array of children of this container.
func (self *Image) SetChildren(member []DisplayObject) {
    self.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Image) GetIgnoreChildInput() bool{
    return self.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Image) SetIgnoreChildInput(member bool) {
    self.Set("ignoreChildInput", member)
}

// A reference to the currently running Game.
func (self *Image) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Image) SetGame(member *Game) {
    self.Set("game", member)
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Image) GetName() string{
    return self.Get("name").String()
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Image) SetName(member string) {
    self.Set("name", member)
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Image) GetData() interface{}{
    return self.Get("data")
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Image) SetData(member interface{}) {
    self.Set("data", member)
}

// The components this Game Object has installed.
func (self *Image) GetComponents() interface{}{
    return self.Get("components")
}

// The components this Game Object has installed.
func (self *Image) SetComponents(member interface{}) {
    self.Set("components", member)
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Image) GetZ() float64{
    return self.Get("z").Float()
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Image) SetZ(member float64) {
    self.Set("z", member)
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Image) GetEvents() *Events{
    return &Events{self.Get("events")}
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Image) SetEvents(member *Events) {
    self.Set("events", member)
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Image) GetAnimations() *AnimationManager{
    return &AnimationManager{self.Get("animations")}
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Image) SetAnimations(member *AnimationManager) {
    self.Set("animations", member)
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Image) GetKey() interface{}{
    return self.Get("key")
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Image) SetKey(member interface{}) {
    self.Set("key", member)
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Image) GetWorld() *Point{
    return &Point{self.Get("world")}
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Image) SetWorld(member *Point) {
    self.Set("world", member)
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Image) GetDebug() bool{
    return self.Get("debug").Bool()
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Image) SetDebug(member bool) {
    self.Set("debug", member)
}

// The position the Game Object was located in the previous frame.
func (self *Image) GetPreviousPosition() *Point{
    return &Point{self.Get("previousPosition")}
}

// The position the Game Object was located in the previous frame.
func (self *Image) SetPreviousPosition(member *Point) {
    self.Set("previousPosition", member)
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Image) GetPreviousRotation() float64{
    return self.Get("previousRotation").Float()
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Image) SetPreviousRotation(member float64) {
    self.Set("previousRotation", member)
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Image) GetRenderOrderID() float64{
    return self.Get("renderOrderID").Float()
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Image) SetRenderOrderID(member float64) {
    self.Set("renderOrderID", member)
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Image) GetFresh() bool{
    return self.Get("fresh").Bool()
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Image) SetFresh(member bool) {
    self.Set("fresh", member)
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Image) GetPendingDestroy() bool{
    return self.Get("pendingDestroy").Bool()
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Image) SetPendingDestroy(member bool) {
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
func (self *Image) GetAngle() float64{
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
func (self *Image) SetAngle(member float64) {
    self.Set("angle", member)
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Image) GetAutoCull() bool{
    return self.Get("autoCull").Bool()
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Image) SetAutoCull(member bool) {
    self.Set("autoCull", member)
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Image) GetInCamera() bool{
    return self.Get("inCamera").Bool()
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Image) SetInCamera(member bool) {
    self.Set("inCamera", member)
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Image) GetOffsetX() float64{
    return self.Get("offsetX").Float()
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Image) SetOffsetX(member float64) {
    self.Set("offsetX", member)
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Image) GetOffsetY() float64{
    return self.Get("offsetY").Float()
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Image) SetOffsetY(member float64) {
    self.Set("offsetY", member)
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Image) GetCenterX() float64{
    return self.Get("centerX").Float()
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Image) SetCenterX(member float64) {
    self.Set("centerX", member)
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Image) GetCenterY() float64{
    return self.Get("centerY").Float()
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Image) SetCenterY(member float64) {
    self.Set("centerY", member)
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Image) GetLeft() float64{
    return self.Get("left").Float()
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Image) SetLeft(member float64) {
    self.Set("left", member)
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Image) GetRight() float64{
    return self.Get("right").Float()
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Image) SetRight(member float64) {
    self.Set("right", member)
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Image) GetTop() float64{
    return self.Get("top").Float()
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Image) SetTop(member float64) {
    self.Set("top", member)
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Image) GetBottom() float64{
    return self.Get("bottom").Float()
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Image) SetBottom(member float64) {
    self.Set("bottom", member)
}

// The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Image) GetCropRect() *Rectangle{
    return &Rectangle{self.Get("cropRect")}
}

// The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Image) SetCropRect(member *Rectangle) {
    self.Set("cropRect", member)
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Image) GetDestroyPhase() bool{
    return self.Get("destroyPhase").Bool()
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Image) SetDestroyPhase(member bool) {
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
func (self *Image) GetFixedToCamera() bool{
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
func (self *Image) SetFixedToCamera(member bool) {
    self.Set("fixedToCamera", member)
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Image) GetCameraOffset() *Point{
    return &Point{self.Get("cameraOffset")}
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Image) SetCameraOffset(member *Point) {
    self.Set("cameraOffset", member)
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Image) GetInput() interface{}{
    return self.Get("input")
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Image) SetInput(member interface{}) {
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
func (self *Image) GetInputEnabled() bool{
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
func (self *Image) SetInputEnabled(member bool) {
    self.Set("inputEnabled", member)
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Image) GetAlive() bool{
    return self.Get("alive").Bool()
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Image) SetAlive(member bool) {
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
func (self *Image) GetLifespan() float64{
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
func (self *Image) SetLifespan(member float64) {
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
func (self *Image) GetFrame() int{
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
func (self *Image) SetFrame(member int) {
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
func (self *Image) GetFrameName() string{
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
func (self *Image) SetFrameName(member string) {
    self.Set("frameName", member)
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *Image) GetSmoothed() bool{
    return self.Get("smoothed").Bool()
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *Image) SetSmoothed(member bool) {
    self.Set("smoothed", member)
}



// Automatically called by World.preUpdate.
func (self *Image) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Image) SetTextureI(args ...interface{}) {
    self.Call("setTexture", args)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *Image) OnTextureUpdateI(args ...interface{}) {
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
func (self *Image) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getBounds", args)}
}

// Renders the object using the WebGL renderer
func (self *Image) _renderWebGLI(args ...interface{}) {
    self.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *Image) _renderCanvasI(args ...interface{}) {
    self.Call("_renderCanvas", args)
}

// Adds a child to the container.
func (self *Image) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Image) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *Image) SwapChildrenI(args ...interface{}) {
    self.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *Image) GetChildIndexI(args ...interface{}) float64{
    return self.Call("getChildIndex", args).Float()
}

// Changes the position of an existing child in the display object container
func (self *Image) SetChildIndexI(args ...interface{}) {
    self.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *Image) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *Image) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *Image) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Image) RemoveChildrenI(args ...interface{}) {
    self.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Image) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Image) SetStageReferenceI(args ...interface{}) {
    self.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *Image) RemoveStageReferenceI(args ...interface{}) {
    self.Call("removeStageReference", args)
}

// Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *Image) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Internal method called by the World postUpdate cycle.
func (self *Image) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Image) PlayI(args ...interface{}) *Animation{
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
func (self *Image) AlignInI(args ...interface{}) interface{}{
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
func (self *Image) AlignToI(args ...interface{}) interface{}{
    return self.Call("alignTo", args)
}

// Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) BringToTopI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("bringToTop", args)}
}

// Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) SendToBackI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("sendToBack", args)}
}

// Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) MoveUpI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("moveUp", args)}
}

// Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Image) MoveDownI(args ...interface{}) *DisplayObject{
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
func (self *Image) CropI(args ...interface{}) {
    self.Call("crop", args)
}

// If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *Image) UpdateCropI(args ...interface{}) {
    self.Call("updateCrop", args)
}

// Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Image) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Image) ReviveI(args ...interface{}) *DisplayObject{
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
func (self *Image) KillI(args ...interface{}) *DisplayObject{
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
func (self *Image) LoadTextureI(args ...interface{}) {
    self.Call("loadTexture", args)
}

// Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *Image) SetFrameI(args ...interface{}) {
    self.Call("setFrame", args)
}

// Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *Image) ResizeFrameI(args ...interface{}) {
    self.Call("resizeFrame", args)
}

// Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *Image) ResetFrameI(args ...interface{}) {
    self.Call("resetFrame", args)
}

// Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *Image) OverlapI(args ...interface{}) bool{
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
func (self *Image) ResetI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("reset", args)}
}
