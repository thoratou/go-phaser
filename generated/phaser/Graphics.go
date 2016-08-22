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
func (self *Graphics) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *Graphics) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The const physics body type of this object.
func (self *Graphics) GetPhysicsTypeA() int{
    return self.Object.Get("physicsType").Int()
}

// The const physics body type of this object.
func (self *Graphics) SetPhysicsTypeA(member int) {
    self.Object.Set("physicsType", member)
}

// The alpha value used when filling the Graphics object.
func (self *Graphics) GetFillAlphaA() int{
    return self.Object.Get("fillAlpha").Int()
}

// The alpha value used when filling the Graphics object.
func (self *Graphics) SetFillAlphaA(member int) {
    self.Object.Set("fillAlpha", member)
}

// The width (thickness) of any lines drawn.
func (self *Graphics) GetLineWidthA() int{
    return self.Object.Get("lineWidth").Int()
}

// The width (thickness) of any lines drawn.
func (self *Graphics) SetLineWidthA(member int) {
    self.Object.Set("lineWidth", member)
}

// The color of any lines drawn.
func (self *Graphics) GetLineColorA() string{
    return self.Object.Get("lineColor").String()
}

// The color of any lines drawn.
func (self *Graphics) SetLineColorA(member string) {
    self.Object.Set("lineColor", member)
}

// The tint applied to the graphic shape. This is a hex value. Apply a value of 0xFFFFFF to reset the tint.
func (self *Graphics) GetTintA() int{
    return self.Object.Get("tint").Int()
}

// The tint applied to the graphic shape. This is a hex value. Apply a value of 0xFFFFFF to reset the tint.
func (self *Graphics) SetTintA(member int) {
    self.Object.Set("tint", member)
}

// The blend mode to be applied to the graphic shape. Apply a value of PIXI.blendModes.NORMAL to reset the blend mode.
func (self *Graphics) GetBlendModeA() int{
    return self.Object.Get("blendMode").Int()
}

// The blend mode to be applied to the graphic shape. Apply a value of PIXI.blendModes.NORMAL to reset the blend mode.
func (self *Graphics) SetBlendModeA(member int) {
    self.Object.Set("blendMode", member)
}

// Whether this shape is being used as a mask.
func (self *Graphics) GetIsMaskA() bool{
    return self.Object.Get("isMask").Bool()
}

// Whether this shape is being used as a mask.
func (self *Graphics) SetIsMaskA(member bool) {
    self.Object.Set("isMask", member)
}

// The bounds' padding used for bounds calculation.
func (self *Graphics) GetBoundsPaddingA() int{
    return self.Object.Get("boundsPadding").Int()
}

// The bounds' padding used for bounds calculation.
func (self *Graphics) SetBoundsPaddingA(member int) {
    self.Object.Set("boundsPadding", member)
}

// [read-only] The array of children of this container.
func (self *Graphics) GetChildrenA() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// [read-only] The array of children of this container.
func (self *Graphics) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Graphics) GetIgnoreChildInputA() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Graphics) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Graphics) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Graphics) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Graphics) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Graphics) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// A reference to the currently running Game.
func (self *Graphics) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *Graphics) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Graphics) GetNameA() string{
    return self.Object.Get("name").String()
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Graphics) SetNameA(member string) {
    self.Object.Set("name", member)
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Graphics) GetDataA() interface{}{
    return self.Object.Get("data")
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Graphics) SetDataA(member interface{}) {
    self.Object.Set("data", member)
}

// The components this Game Object has installed.
func (self *Graphics) GetComponentsA() interface{}{
    return self.Object.Get("components")
}

// The components this Game Object has installed.
func (self *Graphics) SetComponentsA(member interface{}) {
    self.Object.Set("components", member)
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Graphics) GetZA() int{
    return self.Object.Get("z").Int()
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Graphics) SetZA(member int) {
    self.Object.Set("z", member)
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Graphics) GetEventsA() *Events{
    return &Events{self.Object.Get("events")}
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Graphics) SetEventsA(member *Events) {
    self.Object.Set("events", member)
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Graphics) GetAnimationsA() *AnimationManager{
    return &AnimationManager{self.Object.Get("animations")}
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Graphics) SetAnimationsA(member *AnimationManager) {
    self.Object.Set("animations", member)
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Graphics) GetKeyA() interface{}{
    return self.Object.Get("key")
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Graphics) SetKeyA(member interface{}) {
    self.Object.Set("key", member)
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Graphics) GetWorldA() *Point{
    return &Point{self.Object.Get("world")}
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Graphics) SetWorldA(member *Point) {
    self.Object.Set("world", member)
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Graphics) GetDebugA() bool{
    return self.Object.Get("debug").Bool()
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Graphics) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// The position the Game Object was located in the previous frame.
func (self *Graphics) GetPreviousPositionA() *Point{
    return &Point{self.Object.Get("previousPosition")}
}

// The position the Game Object was located in the previous frame.
func (self *Graphics) SetPreviousPositionA(member *Point) {
    self.Object.Set("previousPosition", member)
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Graphics) GetPreviousRotationA() int{
    return self.Object.Get("previousRotation").Int()
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Graphics) SetPreviousRotationA(member int) {
    self.Object.Set("previousRotation", member)
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Graphics) GetRenderOrderIDA() int{
    return self.Object.Get("renderOrderID").Int()
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Graphics) SetRenderOrderIDA(member int) {
    self.Object.Set("renderOrderID", member)
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Graphics) GetFreshA() bool{
    return self.Object.Get("fresh").Bool()
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Graphics) SetFreshA(member bool) {
    self.Object.Set("fresh", member)
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Graphics) GetPendingDestroyA() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Graphics) SetPendingDestroyA(member bool) {
    self.Object.Set("pendingDestroy", member)
}

// Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *Graphics) GetExistsA() bool{
    return self.Object.Get("exists").Bool()
}

// Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *Graphics) SetExistsA(member bool) {
    self.Object.Set("exists", member)
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
func (self *Graphics) GetAngleA() int{
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
func (self *Graphics) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Graphics) GetAutoCullA() bool{
    return self.Object.Get("autoCull").Bool()
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Graphics) SetAutoCullA(member bool) {
    self.Object.Set("autoCull", member)
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Graphics) GetInCameraA() bool{
    return self.Object.Get("inCamera").Bool()
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Graphics) SetInCameraA(member bool) {
    self.Object.Set("inCamera", member)
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Graphics) GetOffsetXA() int{
    return self.Object.Get("offsetX").Int()
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Graphics) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Graphics) GetOffsetYA() int{
    return self.Object.Get("offsetY").Int()
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Graphics) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Graphics) GetCenterXA() int{
    return self.Object.Get("centerX").Int()
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Graphics) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Graphics) GetCenterYA() int{
    return self.Object.Get("centerY").Int()
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Graphics) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Graphics) GetLeftA() int{
    return self.Object.Get("left").Int()
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Graphics) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Graphics) GetRightA() int{
    return self.Object.Get("right").Int()
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Graphics) SetRightA(member int) {
    self.Object.Set("right", member)
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Graphics) GetTopA() int{
    return self.Object.Get("top").Int()
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Graphics) SetTopA(member int) {
    self.Object.Set("top", member)
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Graphics) GetBottomA() int{
    return self.Object.Get("bottom").Int()
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Graphics) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Graphics) GetDestroyPhaseA() bool{
    return self.Object.Get("destroyPhase").Bool()
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Graphics) SetDestroyPhaseA(member bool) {
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
func (self *Graphics) GetFixedToCameraA() bool{
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
func (self *Graphics) SetFixedToCameraA(member bool) {
    self.Object.Set("fixedToCamera", member)
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Graphics) GetCameraOffsetA() *Point{
    return &Point{self.Object.Get("cameraOffset")}
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Graphics) SetCameraOffsetA(member *Point) {
    self.Object.Set("cameraOffset", member)
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Graphics) GetInputA() interface{}{
    return self.Object.Get("input")
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Graphics) SetInputA(member interface{}) {
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
func (self *Graphics) GetInputEnabledA() bool{
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
func (self *Graphics) SetInputEnabledA(member bool) {
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
func (self *Graphics) GetCheckWorldBoundsA() bool{
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
func (self *Graphics) SetCheckWorldBoundsA(member bool) {
    self.Object.Set("checkWorldBounds", member)
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *Graphics) GetOutOfBoundsKillA() bool{
    return self.Object.Get("outOfBoundsKill").Bool()
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *Graphics) SetOutOfBoundsKillA(member bool) {
    self.Object.Set("outOfBoundsKill", member)
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *Graphics) GetOutOfCameraBoundsKillA() bool{
    return self.Object.Get("outOfCameraBoundsKill").Bool()
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *Graphics) SetOutOfCameraBoundsKillA(member bool) {
    self.Object.Set("outOfCameraBoundsKill", member)
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *Graphics) GetInWorldA() bool{
    return self.Object.Get("inWorld").Bool()
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *Graphics) SetInWorldA(member bool) {
    self.Object.Set("inWorld", member)
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Graphics) GetAliveA() bool{
    return self.Object.Get("alive").Bool()
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Graphics) SetAliveA(member bool) {
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
func (self *Graphics) GetLifespanA() int{
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
func (self *Graphics) SetLifespanA(member int) {
    self.Object.Set("lifespan", member)
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
func (self *Graphics) GetBodyA() interface{}{
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
func (self *Graphics) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *Graphics) GetXA() int{
    return self.Object.Get("x").Int()
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *Graphics) SetXA(member int) {
    self.Object.Set("x", member)
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *Graphics) GetYA() int{
    return self.Object.Get("y").Int()
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *Graphics) SetYA(member int) {
    self.Object.Set("y", member)
}



// Automatically called by World.preUpdate.
func (self *Graphics) PreUpdate() {
    self.Object.Call("preUpdate")
}

// Automatically called by World.preUpdate.
func (self *Graphics) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Destroy this Graphics instance.
func (self *Graphics) Destroy(destroyChildren bool) {
    self.Object.Call("destroy", destroyChildren)
}

// Destroy this Graphics instance.
func (self *Graphics) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Specifies the line style used for subsequent calls to Graphics methods such as the lineTo() method or the drawCircle() method.
func (self *Graphics) LineStyle(lineWidth int, color int, alpha int) *Graphics{
    return &Graphics{self.Object.Call("lineStyle", lineWidth, color, alpha)}
}

// Specifies the line style used for subsequent calls to Graphics methods such as the lineTo() method or the drawCircle() method.
func (self *Graphics) LineStyleI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("lineStyle", args)}
}

// Moves the current drawing position to x, y.
func (self *Graphics) MoveTo(x int, y int) *Graphics{
    return &Graphics{self.Object.Call("moveTo", x, y)}
}

// Moves the current drawing position to x, y.
func (self *Graphics) MoveToI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("moveTo", args)}
}

// Draws a line using the current line style from the current drawing position to (x, y);
// The current drawing position is then set to (x, y).
func (self *Graphics) LineTo(x int, y int) *Graphics{
    return &Graphics{self.Object.Call("lineTo", x, y)}
}

// Draws a line using the current line style from the current drawing position to (x, y);
// The current drawing position is then set to (x, y).
func (self *Graphics) LineToI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("lineTo", args)}
}

// Calculate the points for a quadratic bezier curve and then draws it.
// Based on: https://stackoverflow.com/questions/785097/how-do-i-implement-a-bezier-curve-in-c
func (self *Graphics) QuadraticCurveTo(cpX int, cpY int, toX int, toY int) *Graphics{
    return &Graphics{self.Object.Call("quadraticCurveTo", cpX, cpY, toX, toY)}
}

// Calculate the points for a quadratic bezier curve and then draws it.
// Based on: https://stackoverflow.com/questions/785097/how-do-i-implement-a-bezier-curve-in-c
func (self *Graphics) QuadraticCurveToI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("quadraticCurveTo", args)}
}

// Calculate the points for a bezier curve and then draws it.
func (self *Graphics) BezierCurveTo(cpX int, cpY int, cpX2 int, cpY2 int, toX int, toY int) *Graphics{
    return &Graphics{self.Object.Call("bezierCurveTo", cpX, cpY, cpX2, cpY2, toX, toY)}
}

// Calculate the points for a bezier curve and then draws it.
func (self *Graphics) BezierCurveToI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("bezierCurveTo", args)}
}

// The arc method creates an arc/curve (used to create circles, or parts of circles).
func (self *Graphics) Arc(cx int, cy int, radius int, startAngle int, endAngle int, anticlockwise bool, segments int) *Graphics{
    return &Graphics{self.Object.Call("arc", cx, cy, radius, startAngle, endAngle, anticlockwise, segments)}
}

// The arc method creates an arc/curve (used to create circles, or parts of circles).
func (self *Graphics) ArcI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("arc", args)}
}

// Specifies a simple one-color fill that subsequent calls to other Graphics methods
// (such as lineTo() or drawCircle()) use when drawing.
func (self *Graphics) BeginFill(color int, alpha int) *Graphics{
    return &Graphics{self.Object.Call("beginFill", color, alpha)}
}

// Specifies a simple one-color fill that subsequent calls to other Graphics methods
// (such as lineTo() or drawCircle()) use when drawing.
func (self *Graphics) BeginFillI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("beginFill", args)}
}

// Applies a fill to the lines and shapes that were added since the last call to the beginFill() method.
func (self *Graphics) EndFill() *Graphics{
    return &Graphics{self.Object.Call("endFill")}
}

// Applies a fill to the lines and shapes that were added since the last call to the beginFill() method.
func (self *Graphics) EndFillI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("endFill", args)}
}

// 
func (self *Graphics) DrawRect(x int, y int, width int, height int) *Graphics{
    return &Graphics{self.Object.Call("drawRect", x, y, width, height)}
}

// 
func (self *Graphics) DrawRectI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("drawRect", args)}
}

// 
func (self *Graphics) DrawRoundedRect(x int, y int, width int, height int, radius int) {
    self.Object.Call("drawRoundedRect", x, y, width, height, radius)
}

// 
func (self *Graphics) DrawRoundedRectI(args ...interface{}) {
    self.Object.Call("drawRoundedRect", args)
}

// Draws a circle.
func (self *Graphics) DrawCircle(x int, y int, diameter int) *Graphics{
    return &Graphics{self.Object.Call("drawCircle", x, y, diameter)}
}

// Draws a circle.
func (self *Graphics) DrawCircleI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("drawCircle", args)}
}

// Draws an ellipse.
func (self *Graphics) DrawEllipse(x int, y int, width int, height int) *Graphics{
    return &Graphics{self.Object.Call("drawEllipse", x, y, width, height)}
}

// Draws an ellipse.
func (self *Graphics) DrawEllipseI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("drawEllipse", args)}
}

// Draws a polygon using the given path.
func (self *Graphics) DrawPolygon(path interface{}) *Graphics{
    return &Graphics{self.Object.Call("drawPolygon", path)}
}

// Draws a polygon using the given path.
func (self *Graphics) DrawPolygonI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("drawPolygon", args)}
}

// Clears the graphics that were drawn to this Graphics object, and resets fill and line style settings.
func (self *Graphics) Clear() *Graphics{
    return &Graphics{self.Object.Call("clear")}
}

// Clears the graphics that were drawn to this Graphics object, and resets fill and line style settings.
func (self *Graphics) ClearI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("clear", args)}
}

// Useful function that returns a texture of the graphics object that can then be used to create sprites
// This can be quite useful if your geometry is complicated and needs to be reused multiple times.
func (self *Graphics) GenerateTexture(resolution int, scaleMode int, padding int) *Texture{
    return &Texture{self.Object.Call("generateTexture", resolution, scaleMode, padding)}
}

// Useful function that returns a texture of the graphics object that can then be used to create sprites
// This can be quite useful if your geometry is complicated and needs to be reused multiple times.
func (self *Graphics) GenerateTextureI(args ...interface{}) *Texture{
    return &Texture{self.Object.Call("generateTexture", args)}
}

// Renders the object using the WebGL renderer
func (self *Graphics) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// Renders the object using the WebGL renderer
func (self *Graphics) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *Graphics) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// Renders the object using the Canvas renderer
func (self *Graphics) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

// Retrieves the bounds of the graphic shape as a rectangle object
func (self *Graphics) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// Retrieves the bounds of the graphic shape as a rectangle object
func (self *Graphics) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Update the bounds of the object
func (self *Graphics) UpdateLocalBounds() {
    self.Object.Call("updateLocalBounds")
}

// Update the bounds of the object
func (self *Graphics) UpdateLocalBoundsI(args ...interface{}) {
    self.Object.Call("updateLocalBounds", args)
}

// Generates the cached sprite when the sprite has cacheAsBitmap = true
func (self *Graphics) _generateCachedSprite() {
    self.Object.Call("_generateCachedSprite")
}

// Generates the cached sprite when the sprite has cacheAsBitmap = true
func (self *Graphics) _generateCachedSpriteI(args ...interface{}) {
    self.Object.Call("_generateCachedSprite", args)
}

// Updates texture size based on canvas size
func (self *Graphics) UpdateCachedSpriteTexture() {
    self.Object.Call("updateCachedSpriteTexture")
}

// Updates texture size based on canvas size
func (self *Graphics) UpdateCachedSpriteTextureI(args ...interface{}) {
    self.Object.Call("updateCachedSpriteTexture", args)
}

// Destroys a previous cached sprite.
func (self *Graphics) DestroyCachedSprite() {
    self.Object.Call("destroyCachedSprite")
}

// Destroys a previous cached sprite.
func (self *Graphics) DestroyCachedSpriteI(args ...interface{}) {
    self.Object.Call("destroyCachedSprite", args)
}

// Draws the given shape to this Graphics object. Can be any of Circle, Rectangle, Ellipse, Line or Polygon.
func (self *Graphics) DrawShape(shape interface{}) *GraphicsData{
    return &GraphicsData{self.Object.Call("drawShape", shape)}
}

// Draws the given shape to this Graphics object. Can be any of Circle, Rectangle, Ellipse, Line or Polygon.
func (self *Graphics) DrawShapeI(args ...interface{}) *GraphicsData{
    return &GraphicsData{self.Object.Call("drawShape", args)}
}

// Adds a child to the container.
func (self *Graphics) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// Adds a child to the container.
func (self *Graphics) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Graphics) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Graphics) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *Graphics) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// Swaps the position of 2 Display Objects within this container.
func (self *Graphics) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *Graphics) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// Returns the index position of a child DisplayObject instance
func (self *Graphics) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *Graphics) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// Changes the position of an existing child in the display object container
func (self *Graphics) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *Graphics) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// Returns the child at the specified index
func (self *Graphics) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *Graphics) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// Removes a child from the container.
func (self *Graphics) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *Graphics) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// Removes a child from the specified index position.
func (self *Graphics) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Graphics) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Graphics) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Graphics) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Graphics) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Graphics) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Graphics) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *Graphics) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// Removes the current stage reference from the container and all of its children.
func (self *Graphics) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *Graphics) Update() {
    self.Object.Call("update")
}

// Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *Graphics) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Internal method called by the World postUpdate cycle.
func (self *Graphics) PostUpdate() {
    self.Object.Call("postUpdate")
}

// Internal method called by the World postUpdate cycle.
func (self *Graphics) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
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
func (self *Graphics) AlignIn(container interface{}, position int, offsetX int, offsetY int) interface{}{
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
func (self *Graphics) AlignInI(args ...interface{}) interface{}{
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
func (self *Graphics) AlignTo(parent interface{}, position int, offsetX int, offsetY int) interface{}{
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
func (self *Graphics) AlignToI(args ...interface{}) interface{}{
    return self.Object.Call("alignTo", args)
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Graphics) Revive(health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", health)}
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Graphics) ReviveI(args ...interface{}) *DisplayObject{
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
func (self *Graphics) Kill() *DisplayObject{
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
func (self *Graphics) KillI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("kill", args)}
}

// Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *Graphics) Reset(x int, y int, health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", x, y, health)}
}

// Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *Graphics) ResetI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", args)}
}
