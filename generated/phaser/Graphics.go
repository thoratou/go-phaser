// Automatic generation for Phaser.Graphics
// generated file Graphics.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A Graphics object is a way to draw primitives to your game. Primitives include forms of geometry, such as Rectangles,
// Circles and Polygons. They also include lines, arcs and curves. When you initially create a Graphics object it will
// be empty. To 'draw' to it you first specify a lineStyle or fillStyle (or both), and then draw a shape. For example:
// 
// ```
// graphics.beginFill(0xff0000);
// graphics.drawCircle(50, 50, 100);
// graphics.endFill();
// ```
// 
// This will draw a circle shape to the Graphics object, with a diameter of 100, located at x: 50, y: 50.
// 
// When a Graphics object is rendered it will render differently based on if the game is running under Canvas or
// WebGL. Under Canvas it will use the HTML Canvas context drawing operations to draw the path. Under WebGL the
// graphics data is decomposed into polygons. Both of these are expensive processes, especially with complex shapes.
// 
// If your Graphics object doesn't change much (or at all) once you've drawn your shape to it, then you will help
// performance by calling `Graphics.generateTexture`. This will 'bake' the Graphics object into a Texture, and return it.
// You can then use this Texture for Sprites or other display objects. If your Graphics object updates frequently then
// you should avoid doing this, as it will constantly generate new textures, which will consume memory.
// 
// As you can tell, Graphics objects are a bit of a trade-off. While they are extremely useful, you need to be careful
// in their complexity and quantity of them in your game.
type Graphics struct {
    *js.Object
}


// The const type of this object.
func (self *Graphics) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *Graphics) SetType(member float64) {
    self.Set("type", member)
}

// The const physics body type of this object.
func (self *Graphics) GetPhysicsType() float64{
    return self.Get("physicsType").Float()
}

// The const physics body type of this object.
func (self *Graphics) SetPhysicsType(member float64) {
    self.Set("physicsType", member)
}

// The alpha value used when filling the Graphics object.
func (self *Graphics) GetFillAlpha() float64{
    return self.Get("fillAlpha").Float()
}

// The alpha value used when filling the Graphics object.
func (self *Graphics) SetFillAlpha(member float64) {
    self.Set("fillAlpha", member)
}

// The width (thickness) of any lines drawn.
func (self *Graphics) GetLineWidth() float64{
    return self.Get("lineWidth").Float()
}

// The width (thickness) of any lines drawn.
func (self *Graphics) SetLineWidth(member float64) {
    self.Set("lineWidth", member)
}

// The color of any lines drawn.
func (self *Graphics) GetLineColor() string{
    return self.Get("lineColor").String()
}

// The color of any lines drawn.
func (self *Graphics) SetLineColor(member string) {
    self.Set("lineColor", member)
}

// The tint applied to the graphic shape. This is a hex value. Apply a value of 0xFFFFFF to reset the tint.
func (self *Graphics) GetTint() float64{
    return self.Get("tint").Float()
}

// The tint applied to the graphic shape. This is a hex value. Apply a value of 0xFFFFFF to reset the tint.
func (self *Graphics) SetTint(member float64) {
    self.Set("tint", member)
}

// The blend mode to be applied to the graphic shape. Apply a value of PIXI.blendModes.NORMAL to reset the blend mode.
func (self *Graphics) GetBlendMode() float64{
    return self.Get("blendMode").Float()
}

// The blend mode to be applied to the graphic shape. Apply a value of PIXI.blendModes.NORMAL to reset the blend mode.
func (self *Graphics) SetBlendMode(member float64) {
    self.Set("blendMode", member)
}

// Whether this shape is being used as a mask.
func (self *Graphics) GetIsMask() bool{
    return self.Get("isMask").Bool()
}

// Whether this shape is being used as a mask.
func (self *Graphics) SetIsMask(member bool) {
    self.Set("isMask", member)
}

// The bounds' padding used for bounds calculation.
func (self *Graphics) GetBoundsPadding() float64{
    return self.Get("boundsPadding").Float()
}

// The bounds' padding used for bounds calculation.
func (self *Graphics) SetBoundsPadding(member float64) {
    self.Set("boundsPadding", member)
}

// [read-only] The array of children of this container.
func (self *Graphics) GetChildren() []DisplayObject{
	array := self.Get("children")
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// [read-only] The array of children of this container.
func (self *Graphics) SetChildren(member []DisplayObject) {
    self.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Graphics) GetIgnoreChildInput() bool{
    return self.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Graphics) SetIgnoreChildInput(member bool) {
    self.Set("ignoreChildInput", member)
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Graphics) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Graphics) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Graphics) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Graphics) SetHeight(member float64) {
    self.Set("height", member)
}

// A reference to the currently running Game.
func (self *Graphics) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Graphics) SetGame(member Game) {
    self.Set("game", member)
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Graphics) GetName() string{
    return self.Get("name").String()
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Graphics) SetName(member string) {
    self.Set("name", member)
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Graphics) GetData() interface{}{
    return self.Get("data")
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Graphics) SetData(member interface{}) {
    self.Set("data", member)
}

// The components this Game Object has installed.
func (self *Graphics) GetComponents() interface{}{
    return self.Get("components")
}

// The components this Game Object has installed.
func (self *Graphics) SetComponents(member interface{}) {
    self.Set("components", member)
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Graphics) GetZ() float64{
    return self.Get("z").Float()
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Graphics) SetZ(member float64) {
    self.Set("z", member)
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Graphics) GetEvents() Events{
    return Events{self.Get("events")}
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Graphics) SetEvents(member Events) {
    self.Set("events", member)
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Graphics) GetAnimations() AnimationManager{
    return AnimationManager{self.Get("animations")}
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Graphics) SetAnimations(member AnimationManager) {
    self.Set("animations", member)
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Graphics) GetKey() interface{}{
    return self.Get("key")
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Graphics) SetKey(member interface{}) {
    self.Set("key", member)
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Graphics) GetWorld() Point{
    return Point{self.Get("world")}
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Graphics) SetWorld(member Point) {
    self.Set("world", member)
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Graphics) GetDebug() bool{
    return self.Get("debug").Bool()
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Graphics) SetDebug(member bool) {
    self.Set("debug", member)
}

// The position the Game Object was located in the previous frame.
func (self *Graphics) GetPreviousPosition() Point{
    return Point{self.Get("previousPosition")}
}

// The position the Game Object was located in the previous frame.
func (self *Graphics) SetPreviousPosition(member Point) {
    self.Set("previousPosition", member)
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Graphics) GetPreviousRotation() float64{
    return self.Get("previousRotation").Float()
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Graphics) SetPreviousRotation(member float64) {
    self.Set("previousRotation", member)
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Graphics) GetRenderOrderID() float64{
    return self.Get("renderOrderID").Float()
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Graphics) SetRenderOrderID(member float64) {
    self.Set("renderOrderID", member)
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Graphics) GetFresh() bool{
    return self.Get("fresh").Bool()
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Graphics) SetFresh(member bool) {
    self.Set("fresh", member)
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Graphics) GetPendingDestroy() bool{
    return self.Get("pendingDestroy").Bool()
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Graphics) SetPendingDestroy(member bool) {
    self.Set("pendingDestroy", member)
}

// Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *Graphics) GetExists() bool{
    return self.Get("exists").Bool()
}

// Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *Graphics) SetExists(member bool) {
    self.Set("exists", member)
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
func (self *Graphics) GetAngle() float64{
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
func (self *Graphics) SetAngle(member float64) {
    self.Set("angle", member)
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Graphics) GetAutoCull() bool{
    return self.Get("autoCull").Bool()
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Graphics) SetAutoCull(member bool) {
    self.Set("autoCull", member)
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Graphics) GetInCamera() bool{
    return self.Get("inCamera").Bool()
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Graphics) SetInCamera(member bool) {
    self.Set("inCamera", member)
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Graphics) GetOffsetX() float64{
    return self.Get("offsetX").Float()
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Graphics) SetOffsetX(member float64) {
    self.Set("offsetX", member)
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Graphics) GetOffsetY() float64{
    return self.Get("offsetY").Float()
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Graphics) SetOffsetY(member float64) {
    self.Set("offsetY", member)
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Graphics) GetCenterX() float64{
    return self.Get("centerX").Float()
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Graphics) SetCenterX(member float64) {
    self.Set("centerX", member)
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Graphics) GetCenterY() float64{
    return self.Get("centerY").Float()
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Graphics) SetCenterY(member float64) {
    self.Set("centerY", member)
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Graphics) GetLeft() float64{
    return self.Get("left").Float()
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Graphics) SetLeft(member float64) {
    self.Set("left", member)
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Graphics) GetRight() float64{
    return self.Get("right").Float()
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Graphics) SetRight(member float64) {
    self.Set("right", member)
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Graphics) GetTop() float64{
    return self.Get("top").Float()
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Graphics) SetTop(member float64) {
    self.Set("top", member)
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Graphics) GetBottom() float64{
    return self.Get("bottom").Float()
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Graphics) SetBottom(member float64) {
    self.Set("bottom", member)
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Graphics) GetDestroyPhase() bool{
    return self.Get("destroyPhase").Bool()
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Graphics) SetDestroyPhase(member bool) {
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
func (self *Graphics) GetFixedToCamera() bool{
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
func (self *Graphics) SetFixedToCamera(member bool) {
    self.Set("fixedToCamera", member)
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Graphics) GetCameraOffset() Point{
    return Point{self.Get("cameraOffset")}
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Graphics) SetCameraOffset(member Point) {
    self.Set("cameraOffset", member)
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Graphics) GetInput() interface{}{
    return self.Get("input")
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Graphics) SetInput(member interface{}) {
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
func (self *Graphics) GetInputEnabled() bool{
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
func (self *Graphics) SetInputEnabled(member bool) {
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
func (self *Graphics) GetCheckWorldBounds() bool{
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
func (self *Graphics) SetCheckWorldBounds(member bool) {
    self.Set("checkWorldBounds", member)
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *Graphics) GetOutOfBoundsKill() bool{
    return self.Get("outOfBoundsKill").Bool()
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *Graphics) SetOutOfBoundsKill(member bool) {
    self.Set("outOfBoundsKill", member)
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *Graphics) GetOutOfCameraBoundsKill() bool{
    return self.Get("outOfCameraBoundsKill").Bool()
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *Graphics) SetOutOfCameraBoundsKill(member bool) {
    self.Set("outOfCameraBoundsKill", member)
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *Graphics) GetInWorld() bool{
    return self.Get("inWorld").Bool()
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *Graphics) SetInWorld(member bool) {
    self.Set("inWorld", member)
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Graphics) GetAlive() bool{
    return self.Get("alive").Bool()
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Graphics) SetAlive(member bool) {
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
func (self *Graphics) GetLifespan() float64{
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
func (self *Graphics) SetLifespan(member float64) {
    self.Set("lifespan", member)
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
func (self *Graphics) GetBody() interface{}{
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
func (self *Graphics) SetBody(member interface{}) {
    self.Set("body", member)
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *Graphics) GetX() float64{
    return self.Get("x").Float()
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *Graphics) SetX(member float64) {
    self.Set("x", member)
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *Graphics) GetY() float64{
    return self.Get("y").Float()
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *Graphics) SetY(member float64) {
    self.Set("y", member)
}



// Automatically called by World.preUpdate.
func (self *Graphics) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Destroy this Graphics instance.
func (self *Graphics) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Specifies the line style used for subsequent calls to Graphics methods such as the lineTo() method or the drawCircle() method.
func (self *Graphics) LineStyleI(args ...interface{}) Graphics{
    return Graphics{self.Call("lineStyle", args)}
}

// Moves the current drawing position to x, y.
func (self *Graphics) MoveToI(args ...interface{}) Graphics{
    return Graphics{self.Call("moveTo", args)}
}

// Draws a line using the current line style from the current drawing position to (x, y);
// The current drawing position is then set to (x, y).
func (self *Graphics) LineToI(args ...interface{}) Graphics{
    return Graphics{self.Call("lineTo", args)}
}

// Calculate the points for a quadratic bezier curve and then draws it.
// Based on: https://stackoverflow.com/questions/785097/how-do-i-implement-a-bezier-curve-in-c
func (self *Graphics) QuadraticCurveToI(args ...interface{}) Graphics{
    return Graphics{self.Call("quadraticCurveTo", args)}
}

// Calculate the points for a bezier curve and then draws it.
func (self *Graphics) BezierCurveToI(args ...interface{}) Graphics{
    return Graphics{self.Call("bezierCurveTo", args)}
}

// The arc method creates an arc/curve (used to create circles, or parts of circles).
func (self *Graphics) ArcI(args ...interface{}) Graphics{
    return Graphics{self.Call("arc", args)}
}

// Specifies a simple one-color fill that subsequent calls to other Graphics methods
// (such as lineTo() or drawCircle()) use when drawing.
func (self *Graphics) BeginFillI(args ...interface{}) Graphics{
    return Graphics{self.Call("beginFill", args)}
}

// Applies a fill to the lines and shapes that were added since the last call to the beginFill() method.
func (self *Graphics) EndFillI(args ...interface{}) Graphics{
    return Graphics{self.Call("endFill", args)}
}

// 
func (self *Graphics) DrawRectI(args ...interface{}) Graphics{
    return Graphics{self.Call("drawRect", args)}
}

// 
func (self *Graphics) DrawRoundedRectI(args ...interface{}) {
    self.Call("drawRoundedRect", args)
}

// Draws a circle.
func (self *Graphics) DrawCircleI(args ...interface{}) Graphics{
    return Graphics{self.Call("drawCircle", args)}
}

// Draws an ellipse.
func (self *Graphics) DrawEllipseI(args ...interface{}) Graphics{
    return Graphics{self.Call("drawEllipse", args)}
}

// Draws a polygon using the given path.
func (self *Graphics) DrawPolygonI(args ...interface{}) Graphics{
    return Graphics{self.Call("drawPolygon", args)}
}

// Clears the graphics that were drawn to this Graphics object, and resets fill and line style settings.
func (self *Graphics) ClearI(args ...interface{}) Graphics{
    return Graphics{self.Call("clear", args)}
}

// Useful function that returns a texture of the graphics object that can then be used to create sprites
// This can be quite useful if your geometry is complicated and needs to be reused multiple times.
func (self *Graphics) GenerateTextureI(args ...interface{}) Texture{
    return Texture{self.Call("generateTexture", args)}
}

// Renders the object using the WebGL renderer
func (self *Graphics) _renderWebGLI(args ...interface{}) {
    self.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *Graphics) _renderCanvasI(args ...interface{}) {
    self.Call("_renderCanvas", args)
}

// Retrieves the bounds of the graphic shape as a rectangle object
func (self *Graphics) GetBoundsI(args ...interface{}) Rectangle{
    return Rectangle{self.Call("getBounds", args)}
}

// Update the bounds of the object
func (self *Graphics) UpdateLocalBoundsI(args ...interface{}) {
    self.Call("updateLocalBounds", args)
}

// Generates the cached sprite when the sprite has cacheAsBitmap = true
func (self *Graphics) _generateCachedSpriteI(args ...interface{}) {
    self.Call("_generateCachedSprite", args)
}

// Updates texture size based on canvas size
func (self *Graphics) UpdateCachedSpriteTextureI(args ...interface{}) {
    self.Call("updateCachedSpriteTexture", args)
}

// Destroys a previous cached sprite.
func (self *Graphics) DestroyCachedSpriteI(args ...interface{}) {
    self.Call("destroyCachedSprite", args)
}

// Draws the given shape to this Graphics object. Can be any of Circle, Rectangle, Ellipse, Line or Polygon.
func (self *Graphics) DrawShapeI(args ...interface{}) GraphicsData{
    return GraphicsData{self.Call("drawShape", args)}
}

// Adds a child to the container.
func (self *Graphics) AddChildI(args ...interface{}) DisplayObject{
    return DisplayObject{self.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Graphics) AddChildAtI(args ...interface{}) DisplayObject{
    return DisplayObject{self.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *Graphics) SwapChildrenI(args ...interface{}) {
    self.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *Graphics) GetChildIndexI(args ...interface{}) float64{
    return self.Call("getChildIndex", args).Float()
}

// Changes the position of an existing child in the display object container
func (self *Graphics) SetChildIndexI(args ...interface{}) {
    self.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *Graphics) GetChildAtI(args ...interface{}) DisplayObject{
    return DisplayObject{self.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *Graphics) RemoveChildI(args ...interface{}) DisplayObject{
    return DisplayObject{self.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *Graphics) RemoveChildAtI(args ...interface{}) DisplayObject{
    return DisplayObject{self.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Graphics) RemoveChildrenI(args ...interface{}) {
    self.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Graphics) GetLocalBoundsI(args ...interface{}) Rectangle{
    return Rectangle{self.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Graphics) SetStageReferenceI(args ...interface{}) {
    self.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *Graphics) RemoveStageReferenceI(args ...interface{}) {
    self.Call("removeStageReference", args)
}

// Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *Graphics) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Internal method called by the World postUpdate cycle.
func (self *Graphics) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
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
func (self *Graphics) AlignInI(args ...interface{}) interface{}{
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
func (self *Graphics) AlignToI(args ...interface{}) interface{}{
    return self.Call("alignTo", args)
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Graphics) ReviveI(args ...interface{}) DisplayObject{
    return DisplayObject{self.Call("revive", args)}
}

// Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *Graphics) KillI(args ...interface{}) DisplayObject{
    return DisplayObject{self.Call("kill", args)}
}

// Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *Graphics) ResetI(args ...interface{}) DisplayObject{
    return DisplayObject{self.Call("reset", args)}
}
