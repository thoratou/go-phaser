// Automatic generation for Phaser.Component.Core
// generated file ComponentCore.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Core Component Features.
type ComponentCore struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *ComponentCore) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *ComponentCore) SetGame(member *Game) {
    self.Set("game", member)
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *ComponentCore) GetName() string{
    return self.Get("name").String()
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *ComponentCore) SetName(member string) {
    self.Set("name", member)
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *ComponentCore) GetData() interface{}{
    return self.Get("data")
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *ComponentCore) SetData(member interface{}) {
    self.Set("data", member)
}

// The components this Game Object has installed.
func (self *ComponentCore) GetComponents() interface{}{
    return self.Get("components")
}

// The components this Game Object has installed.
func (self *ComponentCore) SetComponents(member interface{}) {
    self.Set("components", member)
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *ComponentCore) GetZ() float64{
    return self.Get("z").Float()
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *ComponentCore) SetZ(member float64) {
    self.Set("z", member)
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *ComponentCore) GetEvents() *Events{
    return &Events{self.Get("events")}
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *ComponentCore) SetEvents(member *Events) {
    self.Set("events", member)
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *ComponentCore) GetAnimations() *AnimationManager{
    return &AnimationManager{self.Get("animations")}
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *ComponentCore) SetAnimations(member *AnimationManager) {
    self.Set("animations", member)
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *ComponentCore) GetKey() interface{}{
    return self.Get("key")
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *ComponentCore) SetKey(member interface{}) {
    self.Set("key", member)
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *ComponentCore) GetWorld() *Point{
    return &Point{self.Get("world")}
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *ComponentCore) SetWorld(member *Point) {
    self.Set("world", member)
}

// A debug flag designed for use with `Game.enableStep`.
func (self *ComponentCore) GetDebug() bool{
    return self.Get("debug").Bool()
}

// A debug flag designed for use with `Game.enableStep`.
func (self *ComponentCore) SetDebug(member bool) {
    self.Set("debug", member)
}

// The position the Game Object was located in the previous frame.
func (self *ComponentCore) GetPreviousPosition() *Point{
    return &Point{self.Get("previousPosition")}
}

// The position the Game Object was located in the previous frame.
func (self *ComponentCore) SetPreviousPosition(member *Point) {
    self.Set("previousPosition", member)
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *ComponentCore) GetPreviousRotation() float64{
    return self.Get("previousRotation").Float()
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *ComponentCore) SetPreviousRotation(member float64) {
    self.Set("previousRotation", member)
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *ComponentCore) GetRenderOrderID() float64{
    return self.Get("renderOrderID").Float()
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *ComponentCore) SetRenderOrderID(member float64) {
    self.Set("renderOrderID", member)
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *ComponentCore) GetFresh() bool{
    return self.Get("fresh").Bool()
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *ComponentCore) SetFresh(member bool) {
    self.Set("fresh", member)
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *ComponentCore) GetPendingDestroy() bool{
    return self.Get("pendingDestroy").Bool()
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *ComponentCore) SetPendingDestroy(member bool) {
    self.Set("pendingDestroy", member)
}

// Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *ComponentCore) GetExists() bool{
    return self.Get("exists").Bool()
}

// Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *ComponentCore) SetExists(member bool) {
    self.Set("exists", member)
}



// Installs / registers mixin components.
// 
// The `this` context should be that of the applicable object instance or prototype.
func (self *ComponentCore) InstallI(args ...interface{}) {
    self.Call("install", args)
}

// Initializes the mixin components.
// 
// The `this` context should be an instance of the component mixin target.
func (self *ComponentCore) InitI(args ...interface{}) {
    self.Call("init", args)
}

// Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *ComponentCore) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Internal method called by the World postUpdate cycle.
func (self *ComponentCore) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}
