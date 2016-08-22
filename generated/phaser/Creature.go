// Automatic generation for Phaser.Creature
// generated file Creature.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Creature is a custom Game Object used in conjunction with the Creature Runtime libraries by Kestrel Moon Studios.
// 
// It allows you to display animated Game Objects that were created with the [Creature Automated Animation Tool](http://www.kestrelmoon.com/creature/).
// 
// Note 1: You can only use Phaser.Creature objects in WebGL enabled games. They do not work in Canvas mode games.
// 
// Note 2: You must use a build of Phaser that includes the CreatureMeshBone.js runtime and gl-matrix.js, or have them
// loaded before your Phaser game boots.
// 
// See the Phaser custom build process for more details.
// 
// By default the Creature runtimes are NOT included in any pre-configured version of Phaser.
// 
// So you'll need to do `grunt custom` to create a build that includes them.
type Creature struct {
    *js.Object
}


// Creature is a custom Game Object used in conjunction with the Creature Runtime libraries by Kestrel Moon Studios.
// 
// It allows you to display animated Game Objects that were created with the [Creature Automated Animation Tool](http://www.kestrelmoon.com/creature/).
// 
// Note 1: You can only use Phaser.Creature objects in WebGL enabled games. They do not work in Canvas mode games.
// 
// Note 2: You must use a build of Phaser that includes the CreatureMeshBone.js runtime and gl-matrix.js, or have them
// loaded before your Phaser game boots.
// 
// See the Phaser custom build process for more details.
// 
// By default the Creature runtimes are NOT included in any pre-configured version of Phaser.
// 
// So you'll need to do `grunt custom` to create a build that includes them.
func NewCreature(game *Game, x int, y int, key interface{}, mesh string) *Creature {
    return &Creature{js.Global.Call("Phaser.Creature", game, x, y, key, mesh)}
}

// Creature is a custom Game Object used in conjunction with the Creature Runtime libraries by Kestrel Moon Studios.
// 
// It allows you to display animated Game Objects that were created with the [Creature Automated Animation Tool](http://www.kestrelmoon.com/creature/).
// 
// Note 1: You can only use Phaser.Creature objects in WebGL enabled games. They do not work in Canvas mode games.
// 
// Note 2: You must use a build of Phaser that includes the CreatureMeshBone.js runtime and gl-matrix.js, or have them
// loaded before your Phaser game boots.
// 
// See the Phaser custom build process for more details.
// 
// By default the Creature runtimes are NOT included in any pre-configured version of Phaser.
// 
// So you'll need to do `grunt custom` to create a build that includes them.
func NewCreature1O(game *Game, x int, y int, key interface{}, mesh string, animation string) *Creature {
    return &Creature{js.Global.Call("Phaser.Creature", game, x, y, key, mesh, animation)}
}

// Creature is a custom Game Object used in conjunction with the Creature Runtime libraries by Kestrel Moon Studios.
// 
// It allows you to display animated Game Objects that were created with the [Creature Automated Animation Tool](http://www.kestrelmoon.com/creature/).
// 
// Note 1: You can only use Phaser.Creature objects in WebGL enabled games. They do not work in Canvas mode games.
// 
// Note 2: You must use a build of Phaser that includes the CreatureMeshBone.js runtime and gl-matrix.js, or have them
// loaded before your Phaser game boots.
// 
// See the Phaser custom build process for more details.
// 
// By default the Creature runtimes are NOT included in any pre-configured version of Phaser.
// 
// So you'll need to do `grunt custom` to create a build that includes them.
func NewCreatureI(args ...interface{}) *Creature {
    return &Creature{js.Global.Call("Phaser.Creature", args)}
}



// The const type of this object.
func (self *Creature) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *Creature) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The CreatureAnimation instance.
func (self *Creature) GetAnimationA() *CreatureAnimation{
    return &CreatureAnimation{self.Object.Get("animation")}
}

// The CreatureAnimation instance.
func (self *Creature) SetAnimationA(member *CreatureAnimation) {
    self.Object.Set("animation", member)
}

// The CreatureManager instance for this object.
func (self *Creature) GetManagerA() *CreatureManager{
    return &CreatureManager{self.Object.Get("manager")}
}

// The CreatureManager instance for this object.
func (self *Creature) SetManagerA(member *CreatureManager) {
    self.Object.Set("manager", member)
}

// How quickly the animation advances.
func (self *Creature) GetTimeDeltaA() int{
    return self.Object.Get("timeDelta").Int()
}

// How quickly the animation advances.
func (self *Creature) SetTimeDeltaA(member int) {
    self.Object.Set("timeDelta", member)
}

// The texture the animation is using.
func (self *Creature) GetTextureA() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// The texture the animation is using.
func (self *Creature) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// The minimum bounds point.
func (self *Creature) GetCreatureBoundsMinA() *Point{
    return &Point{self.Object.Get("creatureBoundsMin")}
}

// The minimum bounds point.
func (self *Creature) SetCreatureBoundsMinA(member *Point) {
    self.Object.Set("creatureBoundsMin", member)
}

// The maximum bounds point.
func (self *Creature) GetCreatureBoundsMaxA() *Point{
    return &Point{self.Object.Get("creatureBoundsMax")}
}

// The maximum bounds point.
func (self *Creature) SetCreatureBoundsMaxA(member *Point) {
    self.Object.Set("creatureBoundsMax", member)
}

// The vertices data.
func (self *Creature) GetVerticesA() *Float32Array{
    return &Float32Array{self.Object.Get("vertices")}
}

// The vertices data.
func (self *Creature) SetVerticesA(member *Float32Array) {
    self.Object.Set("vertices", member)
}

// The UV data.
func (self *Creature) GetUvsA() *Float32Array{
    return &Float32Array{self.Object.Get("uvs")}
}

// The UV data.
func (self *Creature) SetUvsA(member *Float32Array) {
    self.Object.Set("uvs", member)
}

// 
func (self *Creature) GetIndicesA() *Uint16Array{
    return &Uint16Array{self.Object.Get("indices")}
}

// 
func (self *Creature) SetIndicesA(member *Uint16Array) {
    self.Object.Set("indices", member)
}

// The vertices colors
func (self *Creature) GetColorsA() *Uint16Array{
    return &Uint16Array{self.Object.Get("colors")}
}

// The vertices colors
func (self *Creature) SetColorsA(member *Uint16Array) {
    self.Object.Set("colors", member)
}

// Is the _current_ animation playing?
func (self *Creature) GetIsPlayingA() bool{
    return self.Object.Get("isPlaying").Bool()
}

// Is the _current_ animation playing?
func (self *Creature) SetIsPlayingA(member bool) {
    self.Object.Set("isPlaying", member)
}

// Should the _current_ animation loop or not?
func (self *Creature) GetLoopA() bool{
    return self.Object.Get("loop").Bool()
}

// Should the _current_ animation loop or not?
func (self *Creature) SetLoopA(member bool) {
    self.Object.Set("loop", member)
}

// [read-only] The array of children of this container.
func (self *Creature) GetChildrenA() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// [read-only] The array of children of this container.
func (self *Creature) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Creature) GetIgnoreChildInputA() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Creature) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Creature) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Creature) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Creature) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Creature) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// A reference to the currently running Game.
func (self *Creature) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *Creature) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Creature) GetNameA() string{
    return self.Object.Get("name").String()
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Creature) SetNameA(member string) {
    self.Object.Set("name", member)
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Creature) GetDataA() interface{}{
    return self.Object.Get("data")
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Creature) SetDataA(member interface{}) {
    self.Object.Set("data", member)
}

// The components this Game Object has installed.
func (self *Creature) GetComponentsA() interface{}{
    return self.Object.Get("components")
}

// The components this Game Object has installed.
func (self *Creature) SetComponentsA(member interface{}) {
    self.Object.Set("components", member)
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Creature) GetZA() int{
    return self.Object.Get("z").Int()
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Creature) SetZA(member int) {
    self.Object.Set("z", member)
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Creature) GetEventsA() *Events{
    return &Events{self.Object.Get("events")}
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Creature) SetEventsA(member *Events) {
    self.Object.Set("events", member)
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Creature) GetAnimationsA() *AnimationManager{
    return &AnimationManager{self.Object.Get("animations")}
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Creature) SetAnimationsA(member *AnimationManager) {
    self.Object.Set("animations", member)
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Creature) GetKeyA() interface{}{
    return self.Object.Get("key")
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Creature) SetKeyA(member interface{}) {
    self.Object.Set("key", member)
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Creature) GetWorldA() *Point{
    return &Point{self.Object.Get("world")}
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Creature) SetWorldA(member *Point) {
    self.Object.Set("world", member)
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Creature) GetDebugA() bool{
    return self.Object.Get("debug").Bool()
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Creature) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// The position the Game Object was located in the previous frame.
func (self *Creature) GetPreviousPositionA() *Point{
    return &Point{self.Object.Get("previousPosition")}
}

// The position the Game Object was located in the previous frame.
func (self *Creature) SetPreviousPositionA(member *Point) {
    self.Object.Set("previousPosition", member)
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Creature) GetPreviousRotationA() int{
    return self.Object.Get("previousRotation").Int()
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Creature) SetPreviousRotationA(member int) {
    self.Object.Set("previousRotation", member)
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Creature) GetRenderOrderIDA() int{
    return self.Object.Get("renderOrderID").Int()
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Creature) SetRenderOrderIDA(member int) {
    self.Object.Set("renderOrderID", member)
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Creature) GetFreshA() bool{
    return self.Object.Get("fresh").Bool()
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Creature) SetFreshA(member bool) {
    self.Object.Set("fresh", member)
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Creature) GetPendingDestroyA() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Creature) SetPendingDestroyA(member bool) {
    self.Object.Set("pendingDestroy", member)
}

// Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *Creature) GetExistsA() bool{
    return self.Object.Get("exists").Bool()
}

// Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *Creature) SetExistsA(member bool) {
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
func (self *Creature) GetAngleA() int{
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
func (self *Creature) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Creature) GetAutoCullA() bool{
    return self.Object.Get("autoCull").Bool()
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Creature) SetAutoCullA(member bool) {
    self.Object.Set("autoCull", member)
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Creature) GetInCameraA() bool{
    return self.Object.Get("inCamera").Bool()
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Creature) SetInCameraA(member bool) {
    self.Object.Set("inCamera", member)
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Creature) GetDestroyPhaseA() bool{
    return self.Object.Get("destroyPhase").Bool()
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Creature) SetDestroyPhaseA(member bool) {
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
func (self *Creature) GetFixedToCameraA() bool{
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
func (self *Creature) SetFixedToCameraA(member bool) {
    self.Object.Set("fixedToCamera", member)
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Creature) GetCameraOffsetA() *Point{
    return &Point{self.Object.Get("cameraOffset")}
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Creature) SetCameraOffsetA(member *Point) {
    self.Object.Set("cameraOffset", member)
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Creature) GetAliveA() bool{
    return self.Object.Get("alive").Bool()
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Creature) SetAliveA(member bool) {
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
func (self *Creature) GetLifespanA() int{
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
func (self *Creature) SetLifespanA(member int) {
    self.Object.Set("lifespan", member)
}



// Automatically called by World.preUpdate.
func (self *Creature) PreUpdate() {
    self.Object.Call("preUpdate")
}

// Automatically called by World.preUpdate.
func (self *Creature) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// 
func (self *Creature) _initWebGL() {
    self.Object.Call("_initWebGL")
}

// 
func (self *Creature) _initWebGLI(args ...interface{}) {
    self.Object.Call("_initWebGL", args)
}

// 
func (self *Creature) _renderWebGL() {
    self.Object.Call("_renderWebGL")
}

// 
func (self *Creature) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// 
func (self *Creature) _renderCreature() {
    self.Object.Call("_renderCreature")
}

// 
func (self *Creature) _renderCreatureI(args ...interface{}) {
    self.Object.Call("_renderCreature", args)
}

// 
func (self *Creature) UpdateCreatureBounds() {
    self.Object.Call("updateCreatureBounds")
}

// 
func (self *Creature) UpdateCreatureBoundsI(args ...interface{}) {
    self.Object.Call("updateCreatureBounds", args)
}

// 
func (self *Creature) UpdateData() {
    self.Object.Call("updateData")
}

// 
func (self *Creature) UpdateDataI(args ...interface{}) {
    self.Object.Call("updateData", args)
}

// 
func (self *Creature) UpdateRenderData() {
    self.Object.Call("updateRenderData")
}

// 
func (self *Creature) UpdateRenderDataI(args ...interface{}) {
    self.Object.Call("updateRenderData", args)
}

// Sets the Animation this Creature object will play, as defined in the mesh data.
func (self *Creature) SetAnimation(key string) {
    self.Object.Call("setAnimation", key)
}

// Sets the Animation this Creature object will play, as defined in the mesh data.
func (self *Creature) SetAnimationI(args ...interface{}) {
    self.Object.Call("setAnimation", args)
}

// Plays the currently set animation.
func (self *Creature) Play() {
    self.Object.Call("play")
}

// Plays the currently set animation.
func (self *Creature) Play1O(loop bool) {
    self.Object.Call("play", loop)
}

// Plays the currently set animation.
func (self *Creature) PlayI(args ...interface{}) {
    self.Object.Call("play", args)
}

// Stops the currently playing animation.
func (self *Creature) Stop() {
    self.Object.Call("stop")
}

// Stops the currently playing animation.
func (self *Creature) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Adds a child to the container.
func (self *Creature) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// Adds a child to the container.
func (self *Creature) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Creature) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Creature) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *Creature) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// Swaps the position of 2 Display Objects within this container.
func (self *Creature) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *Creature) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// Returns the index position of a child DisplayObject instance
func (self *Creature) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *Creature) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// Changes the position of an existing child in the display object container
func (self *Creature) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *Creature) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// Returns the child at the specified index
func (self *Creature) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *Creature) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// Removes a child from the container.
func (self *Creature) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *Creature) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// Removes a child from the specified index position.
func (self *Creature) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Creature) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Creature) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *Creature) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *Creature) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Creature) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Creature) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Creature) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Creature) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *Creature) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// Removes the current stage reference from the container and all of its children.
func (self *Creature) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// Renders the object using the Canvas renderer
func (self *Creature) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// Renders the object using the Canvas renderer
func (self *Creature) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

// Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *Creature) Update() {
    self.Object.Call("update")
}

// Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *Creature) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Internal method called by the World postUpdate cycle.
func (self *Creature) PostUpdate() {
    self.Object.Call("postUpdate")
}

// Internal method called by the World postUpdate cycle.
func (self *Creature) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Creature) BringToTop() *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop")}
}

// Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Creature) BringToTopI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop", args)}
}

// Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Creature) SendToBack() *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack")}
}

// Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Creature) SendToBackI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack", args)}
}

// Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Creature) MoveUp() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp")}
}

// Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Creature) MoveUpI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp", args)}
}

// Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Creature) MoveDown() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown")}
}

// Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Creature) MoveDownI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown", args)}
}

// Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Creature) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Creature) Destroy1O(destroyChildren bool) {
    self.Object.Call("destroy", destroyChildren)
}

// Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Creature) Destroy2O(destroyChildren bool, destroyTexture bool) {
    self.Object.Call("destroy", destroyChildren, destroyTexture)
}

// Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *Creature) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Creature) Revive() *DisplayObject{
    return &DisplayObject{self.Object.Call("revive")}
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Creature) Revive1O(health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", health)}
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Creature) ReviveI(args ...interface{}) *DisplayObject{
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
func (self *Creature) Kill() *DisplayObject{
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
func (self *Creature) KillI(args ...interface{}) *DisplayObject{
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
func (self *Creature) Reset(x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", x, y)}
}

// Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *Creature) Reset1O(x int, y int, health int) *DisplayObject{
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
func (self *Creature) ResetI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", args)}
}
