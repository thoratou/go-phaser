// Package phaser Automatic generation for Phaser.Component.Core
// generated file ComponentCore.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ComponentCore Core Component Features.
type ComponentCore struct {
    *js.Object
}

// NewComponentCore Core Component Features.
func NewComponentCore() *ComponentCore {
    return &ComponentCore{js.Global.Get("Phaser").Get("Component").Get("Core").New()}
}
// NewComponentCoreI Core Component Features.
func NewComponentCoreI(args ...interface{}) *ComponentCore {
    return &ComponentCore{js.Global.Get("Phaser").Get("Component").Get("Core").New(args)}
}



// ComponentCore Binding conversion method to ComponentCore point 
func ToComponentCore(jsStruct interface{}) *ComponentCore {
    if object, ok := jsStruct.(*js.Object); ok {
		return &ComponentCore{Object: object}
	}
	return nil
}



// Game A reference to the currently running Game.
func (self *ComponentCore) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *ComponentCore) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Name A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *ComponentCore) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *ComponentCore) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Data An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *ComponentCore) Data() interface{}{
    return self.Object.Get("data")
}

// SetDataA An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *ComponentCore) SetDataA(member interface{}) {
    self.Object.Set("data", member)
}

// Components The components this Game Object has installed.
func (self *ComponentCore) Components() interface{}{
    return self.Object.Get("components")
}

// SetComponentsA The components this Game Object has installed.
func (self *ComponentCore) SetComponentsA(member interface{}) {
    self.Object.Set("components", member)
}

// Z The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *ComponentCore) Z() int{
    return self.Object.Get("z").Int()
}

// SetZA The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *ComponentCore) SetZA(member int) {
    self.Object.Set("z", member)
}

// Events All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *ComponentCore) Events() *Events{
    return &Events{self.Object.Get("events")}
}

// SetEventsA All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *ComponentCore) SetEventsA(member *Events) {
    self.Object.Set("events", member)
}

// Animations If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *ComponentCore) Animations() *AnimationManager{
    return &AnimationManager{self.Object.Get("animations")}
}

// SetAnimationsA If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *ComponentCore) SetAnimationsA(member *AnimationManager) {
    self.Object.Set("animations", member)
}

// Key The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *ComponentCore) Key() interface{}{
    return self.Object.Get("key")
}

// SetKeyA The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *ComponentCore) SetKeyA(member interface{}) {
    self.Object.Set("key", member)
}

// World The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *ComponentCore) World() *Point{
    return &Point{self.Object.Get("world")}
}

// SetWorldA The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *ComponentCore) SetWorldA(member *Point) {
    self.Object.Set("world", member)
}

// Debug A debug flag designed for use with `Game.enableStep`.
func (self *ComponentCore) Debug() bool{
    return self.Object.Get("debug").Bool()
}

// SetDebugA A debug flag designed for use with `Game.enableStep`.
func (self *ComponentCore) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// PreviousPosition The position the Game Object was located in the previous frame.
func (self *ComponentCore) PreviousPosition() *Point{
    return &Point{self.Object.Get("previousPosition")}
}

// SetPreviousPositionA The position the Game Object was located in the previous frame.
func (self *ComponentCore) SetPreviousPositionA(member *Point) {
    self.Object.Set("previousPosition", member)
}

// PreviousRotation The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *ComponentCore) PreviousRotation() int{
    return self.Object.Get("previousRotation").Int()
}

// SetPreviousRotationA The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *ComponentCore) SetPreviousRotationA(member int) {
    self.Object.Set("previousRotation", member)
}

// RenderOrderID The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *ComponentCore) RenderOrderID() int{
    return self.Object.Get("renderOrderID").Int()
}

// SetRenderOrderIDA The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *ComponentCore) SetRenderOrderIDA(member int) {
    self.Object.Set("renderOrderID", member)
}

// Fresh A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *ComponentCore) Fresh() bool{
    return self.Object.Get("fresh").Bool()
}

// SetFreshA A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *ComponentCore) SetFreshA(member bool) {
    self.Object.Set("fresh", member)
}

// PendingDestroy A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *ComponentCore) PendingDestroy() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// SetPendingDestroyA A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *ComponentCore) SetPendingDestroyA(member bool) {
    self.Object.Set("pendingDestroy", member)
}

// Exists Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *ComponentCore) Exists() bool{
    return self.Object.Get("exists").Bool()
}

// SetExistsA Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *ComponentCore) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}


// Install Installs / registers mixin components.
// 
// The `this` context should be that of the applicable object instance or prototype.
func (self *ComponentCore) Install() {
    self.Object.Call("install")
}

// InstallI Installs / registers mixin components.
// 
// The `this` context should be that of the applicable object instance or prototype.
func (self *ComponentCore) InstallI(args ...interface{}) {
    self.Object.Call("install", args)
}

// Init Initializes the mixin components.
// 
// The `this` context should be an instance of the component mixin target.
func (self *ComponentCore) Init() {
    self.Object.Call("init")
}

// InitI Initializes the mixin components.
// 
// The `this` context should be an instance of the component mixin target.
func (self *ComponentCore) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Update Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *ComponentCore) Update() {
    self.Object.Call("update")
}

// UpdateI Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *ComponentCore) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// PostUpdate Internal method called by the World postUpdate cycle.
func (self *ComponentCore) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI Internal method called by the World postUpdate cycle.
func (self *ComponentCore) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

