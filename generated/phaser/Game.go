// Package phaser Automatic generation for Phaser.Game
// generated file Game.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// Game This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
type Game struct {
    *js.Object
}

// NewGame This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
func NewGame() *Game {
    return &Game{js.Global.Get("Phaser").Get("Game").New()}
}
// NewGame1O This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
func NewGame1O(width interface{}) *Game {
    return &Game{js.Global.Get("Phaser").Get("Game").New(width)}
}
// NewGame2O This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
func NewGame2O(width interface{}, height interface{}) *Game {
    return &Game{js.Global.Get("Phaser").Get("Game").New(width, height)}
}
// NewGame3O This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
func NewGame3O(width interface{}, height interface{}, renderer int) *Game {
    return &Game{js.Global.Get("Phaser").Get("Game").New(width, height, renderer)}
}
// NewGame4O This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
func NewGame4O(width interface{}, height interface{}, renderer int, parent interface{}) *Game {
    return &Game{js.Global.Get("Phaser").Get("Game").New(width, height, renderer, parent)}
}
// NewGame5O This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
func NewGame5O(width interface{}, height interface{}, renderer int, parent interface{}, state interface{}) *Game {
    return &Game{js.Global.Get("Phaser").Get("Game").New(width, height, renderer, parent, state)}
}
// NewGame6O This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
func NewGame6O(width interface{}, height interface{}, renderer int, parent interface{}, state interface{}, transparent bool) *Game {
    return &Game{js.Global.Get("Phaser").Get("Game").New(width, height, renderer, parent, state, transparent)}
}
// NewGame7O This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
func NewGame7O(width interface{}, height interface{}, renderer int, parent interface{}, state interface{}, transparent bool, antialias bool) *Game {
    return &Game{js.Global.Get("Phaser").Get("Game").New(width, height, renderer, parent, state, transparent, antialias)}
}
// NewGame8O This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
func NewGame8O(width interface{}, height interface{}, renderer int, parent interface{}, state interface{}, transparent bool, antialias bool, physicsConfig interface{}) *Game {
    return &Game{js.Global.Get("Phaser").Get("Game").New(width, height, renderer, parent, state, transparent, antialias, physicsConfig)}
}
// NewGameI This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
func NewGameI(args ...interface{}) *Game {
    return &Game{js.Global.Get("Phaser").Get("Game").New(args)}
}



// Id Phaser Game ID (for when Pixi supports multiple instances).
func (self *Game) Id() int{
    return self.Object.Get("id").Int()
}

// SetIdA Phaser Game ID (for when Pixi supports multiple instances).
func (self *Game) SetIdA(member int) {
    self.Object.Set("id", member)
}

// Config The Phaser.Game configuration object.
func (self *Game) Config() interface{}{
    return self.Object.Get("config")
}

// SetConfigA The Phaser.Game configuration object.
func (self *Game) SetConfigA(member interface{}) {
    self.Object.Set("config", member)
}

// PhysicsConfig The Phaser.Physics.World configuration object.
func (self *Game) PhysicsConfig() interface{}{
    return self.Object.Get("physicsConfig")
}

// SetPhysicsConfigA The Phaser.Physics.World configuration object.
func (self *Game) SetPhysicsConfigA(member interface{}) {
    self.Object.Set("physicsConfig", member)
}

// Parent The Games DOM parent.
func (self *Game) Parent() interface{}{
    return self.Object.Get("parent")
}

// SetParentA The Games DOM parent.
func (self *Game) SetParentA(member interface{}) {
    self.Object.Set("parent", member)
}

// Width The current Game Width in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The current Game Width in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The current Game Height in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The current Game Height in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Resolution The resolution of your game. This value is read only, but can be changed at start time it via a game configuration object.
func (self *Game) Resolution() int{
    return self.Object.Get("resolution").Int()
}

// SetResolutionA The resolution of your game. This value is read only, but can be changed at start time it via a game configuration object.
func (self *Game) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// Transparent Use a transparent canvas background or not.
func (self *Game) Transparent() bool{
    return self.Object.Get("transparent").Bool()
}

// SetTransparentA Use a transparent canvas background or not.
func (self *Game) SetTransparentA(member bool) {
    self.Object.Set("transparent", member)
}

// Antialias Anti-alias graphics. By default scaled images are smoothed in Canvas and WebGL, set anti-alias to false to disable this globally.
func (self *Game) Antialias() bool{
    return self.Object.Get("antialias").Bool()
}

// SetAntialiasA Anti-alias graphics. By default scaled images are smoothed in Canvas and WebGL, set anti-alias to false to disable this globally.
func (self *Game) SetAntialiasA(member bool) {
    self.Object.Set("antialias", member)
}

// PreserveDrawingBuffer The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
func (self *Game) PreserveDrawingBuffer() bool{
    return self.Object.Get("preserveDrawingBuffer").Bool()
}

// SetPreserveDrawingBufferA The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
func (self *Game) SetPreserveDrawingBufferA(member bool) {
    self.Object.Set("preserveDrawingBuffer", member)
}

// ClearBeforeRender Clear the Canvas each frame before rendering the display list.
// You can set this to `false` to gain some performance if your game always contains a background that completely fills the display.
func (self *Game) ClearBeforeRender() bool{
    return self.Object.Get("clearBeforeRender").Bool()
}

// SetClearBeforeRenderA Clear the Canvas each frame before rendering the display list.
// You can set this to `false` to gain some performance if your game always contains a background that completely fills the display.
func (self *Game) SetClearBeforeRenderA(member bool) {
    self.Object.Set("clearBeforeRender", member)
}

// Renderer The Pixi Renderer.
func (self *Game) Renderer() interface{}{
    return self.Object.Get("renderer")
}

// SetRendererA The Pixi Renderer.
func (self *Game) SetRendererA(member interface{}) {
    self.Object.Set("renderer", member)
}

// RenderType The Renderer this game will use. Either Phaser.AUTO, Phaser.CANVAS, Phaser.WEBGL, or Phaser.HEADLESS.
func (self *Game) RenderType() int{
    return self.Object.Get("renderType").Int()
}

// SetRenderTypeA The Renderer this game will use. Either Phaser.AUTO, Phaser.CANVAS, Phaser.WEBGL, or Phaser.HEADLESS.
func (self *Game) SetRenderTypeA(member int) {
    self.Object.Set("renderType", member)
}

// State The StateManager.
func (self *Game) State() *StateManager{
    return &StateManager{self.Object.Get("state")}
}

// SetStateA The StateManager.
func (self *Game) SetStateA(member *StateManager) {
    self.Object.Set("state", member)
}

// IsBooted Whether the game engine is booted, aka available.
func (self *Game) IsBooted() bool{
    return self.Object.Get("isBooted").Bool()
}

// SetIsBootedA Whether the game engine is booted, aka available.
func (self *Game) SetIsBootedA(member bool) {
    self.Object.Set("isBooted", member)
}

// IsRunning Is game running or paused?
func (self *Game) IsRunning() bool{
    return self.Object.Get("isRunning").Bool()
}

// SetIsRunningA Is game running or paused?
func (self *Game) SetIsRunningA(member bool) {
    self.Object.Set("isRunning", member)
}

// Raf Automatically handles the core game loop via requestAnimationFrame or setTimeout
func (self *Game) Raf() *RequestAnimationFrame{
    return &RequestAnimationFrame{self.Object.Get("raf")}
}

// SetRafA Automatically handles the core game loop via requestAnimationFrame or setTimeout
func (self *Game) SetRafA(member *RequestAnimationFrame) {
    self.Object.Set("raf", member)
}

// Add Reference to the Phaser.GameObjectFactory.
func (self *Game) Add() *GameObjectFactory{
    return &GameObjectFactory{self.Object.Get("add")}
}

// SetAddA Reference to the Phaser.GameObjectFactory.
func (self *Game) SetAddA(member *GameObjectFactory) {
    self.Object.Set("add", member)
}

// Make Reference to the GameObject Creator.
func (self *Game) Make() *GameObjectCreator{
    return &GameObjectCreator{self.Object.Get("make")}
}

// SetMakeA Reference to the GameObject Creator.
func (self *Game) SetMakeA(member *GameObjectCreator) {
    self.Object.Set("make", member)
}

// Cache Reference to the assets cache.
func (self *Game) Cache() *Cache{
    return &Cache{self.Object.Get("cache")}
}

// SetCacheA Reference to the assets cache.
func (self *Game) SetCacheA(member *Cache) {
    self.Object.Set("cache", member)
}

// Input Reference to the input manager
func (self *Game) Input() *Input{
    return &Input{self.Object.Get("input")}
}

// SetInputA Reference to the input manager
func (self *Game) SetInputA(member *Input) {
    self.Object.Set("input", member)
}

// Load Reference to the assets loader.
func (self *Game) Load() *Loader{
    return &Loader{self.Object.Get("load")}
}

// SetLoadA Reference to the assets loader.
func (self *Game) SetLoadA(member *Loader) {
    self.Object.Set("load", member)
}

// Math Reference to the math helper.
func (self *Game) Math() *Math{
    return &Math{self.Object.Get("math")}
}

// SetMathA Reference to the math helper.
func (self *Game) SetMathA(member *Math) {
    self.Object.Set("math", member)
}

// Net Reference to the network class.
func (self *Game) Net() *Net{
    return &Net{self.Object.Get("net")}
}

// SetNetA Reference to the network class.
func (self *Game) SetNetA(member *Net) {
    self.Object.Set("net", member)
}

// Scale The game scale manager.
func (self *Game) Scale() *ScaleManager{
    return &ScaleManager{self.Object.Get("scale")}
}

// SetScaleA The game scale manager.
func (self *Game) SetScaleA(member *ScaleManager) {
    self.Object.Set("scale", member)
}

// Sound Reference to the sound manager.
func (self *Game) Sound() *SoundManager{
    return &SoundManager{self.Object.Get("sound")}
}

// SetSoundA Reference to the sound manager.
func (self *Game) SetSoundA(member *SoundManager) {
    self.Object.Set("sound", member)
}

// Stage Reference to the stage.
func (self *Game) Stage() *Stage{
    return &Stage{self.Object.Get("stage")}
}

// SetStageA Reference to the stage.
func (self *Game) SetStageA(member *Stage) {
    self.Object.Set("stage", member)
}

// Time Reference to the core game clock.
func (self *Game) Time() *Time{
    return &Time{self.Object.Get("time")}
}

// SetTimeA Reference to the core game clock.
func (self *Game) SetTimeA(member *Time) {
    self.Object.Set("time", member)
}

// Tweens Reference to the tween manager.
func (self *Game) Tweens() *TweenManager{
    return &TweenManager{self.Object.Get("tweens")}
}

// SetTweensA Reference to the tween manager.
func (self *Game) SetTweensA(member *TweenManager) {
    self.Object.Set("tweens", member)
}

// World Reference to the world.
func (self *Game) World() *World{
    return &World{self.Object.Get("world")}
}

// SetWorldA Reference to the world.
func (self *Game) SetWorldA(member *World) {
    self.Object.Set("world", member)
}

// Physics Reference to the physics manager.
func (self *Game) Physics() *Physics{
    return &Physics{self.Object.Get("physics")}
}

// SetPhysicsA Reference to the physics manager.
func (self *Game) SetPhysicsA(member *Physics) {
    self.Object.Set("physics", member)
}

// Plugins Reference to the plugin manager.
func (self *Game) Plugins() *PluginManager{
    return &PluginManager{self.Object.Get("plugins")}
}

// SetPluginsA Reference to the plugin manager.
func (self *Game) SetPluginsA(member *PluginManager) {
    self.Object.Set("plugins", member)
}

// Rnd Instance of repeatable random data generator helper.
func (self *Game) Rnd() *RandomDataGenerator{
    return &RandomDataGenerator{self.Object.Get("rnd")}
}

// SetRndA Instance of repeatable random data generator helper.
func (self *Game) SetRndA(member *RandomDataGenerator) {
    self.Object.Set("rnd", member)
}

// Device Contains device information and capabilities.
func (self *Game) Device() *Device{
    return &Device{self.Object.Get("device")}
}

// SetDeviceA Contains device information and capabilities.
func (self *Game) SetDeviceA(member *Device) {
    self.Object.Set("device", member)
}

// Camera A handy reference to world.camera.
func (self *Game) Camera() *Camera{
    return &Camera{self.Object.Get("camera")}
}

// SetCameraA A handy reference to world.camera.
func (self *Game) SetCameraA(member *Camera) {
    self.Object.Set("camera", member)
}

// Canvas A handy reference to renderer.view, the canvas that the game is being rendered in to.
func (self *Game) Canvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("canvas"))
}

// SetCanvasA A handy reference to renderer.view, the canvas that the game is being rendered in to.
func (self *Game) SetCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("canvas", member)
}

// Context A handy reference to renderer.context (only set for CANVAS games, not WebGL)
func (self *Game) Context() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("context"))
}

// SetContextA A handy reference to renderer.context (only set for CANVAS games, not WebGL)
func (self *Game) SetContextA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("context", member)
}

// Debug A set of useful debug utilities.
func (self *Game) Debug() *UtilsDebug{
    return &UtilsDebug{self.Object.Get("debug")}
}

// SetDebugA A set of useful debug utilities.
func (self *Game) SetDebugA(member *UtilsDebug) {
    self.Object.Set("debug", member)
}

// Particles The Particle Manager.
func (self *Game) Particles() *Particles{
    return &Particles{self.Object.Get("particles")}
}

// SetParticlesA The Particle Manager.
func (self *Game) SetParticlesA(member *Particles) {
    self.Object.Set("particles", member)
}

// Create The Asset Generator.
func (self *Game) Create() *Create{
    return &Create{self.Object.Get("create")}
}

// SetCreateA The Asset Generator.
func (self *Game) SetCreateA(member *Create) {
    self.Object.Set("create", member)
}

// LockRender If `false` Phaser will automatically render the display list every update. If `true` the render loop will be skipped.
// You can toggle this value at run-time to gain exact control over when Phaser renders. This can be useful in certain types of game or application.
// Please note that if you don't render the display list then none of the game object transforms will be updated, so use this value carefully.
func (self *Game) LockRender() bool{
    return self.Object.Get("lockRender").Bool()
}

// SetLockRenderA If `false` Phaser will automatically render the display list every update. If `true` the render loop will be skipped.
// You can toggle this value at run-time to gain exact control over when Phaser renders. This can be useful in certain types of game or application.
// Please note that if you don't render the display list then none of the game object transforms will be updated, so use this value carefully.
func (self *Game) SetLockRenderA(member bool) {
    self.Object.Set("lockRender", member)
}

// Stepping Enable core loop stepping with Game.enableStep().
func (self *Game) Stepping() bool{
    return self.Object.Get("stepping").Bool()
}

// SetSteppingA Enable core loop stepping with Game.enableStep().
func (self *Game) SetSteppingA(member bool) {
    self.Object.Set("stepping", member)
}

// PendingStep An internal property used by enableStep, but also useful to query from your own game objects.
func (self *Game) PendingStep() bool{
    return self.Object.Get("pendingStep").Bool()
}

// SetPendingStepA An internal property used by enableStep, but also useful to query from your own game objects.
func (self *Game) SetPendingStepA(member bool) {
    self.Object.Set("pendingStep", member)
}

// StepCount When stepping is enabled this contains the current step cycle.
func (self *Game) StepCount() int{
    return self.Object.Get("stepCount").Int()
}

// SetStepCountA When stepping is enabled this contains the current step cycle.
func (self *Game) SetStepCountA(member int) {
    self.Object.Set("stepCount", member)
}

// OnPause This event is fired when the game pauses.
func (self *Game) OnPause() *Signal{
    return &Signal{self.Object.Get("onPause")}
}

// SetOnPauseA This event is fired when the game pauses.
func (self *Game) SetOnPauseA(member *Signal) {
    self.Object.Set("onPause", member)
}

// OnResume This event is fired when the game resumes from a paused state.
func (self *Game) OnResume() *Signal{
    return &Signal{self.Object.Get("onResume")}
}

// SetOnResumeA This event is fired when the game resumes from a paused state.
func (self *Game) SetOnResumeA(member *Signal) {
    self.Object.Set("onResume", member)
}

// OnBlur This event is fired when the game no longer has focus (typically on page hide).
func (self *Game) OnBlur() *Signal{
    return &Signal{self.Object.Get("onBlur")}
}

// SetOnBlurA This event is fired when the game no longer has focus (typically on page hide).
func (self *Game) SetOnBlurA(member *Signal) {
    self.Object.Set("onBlur", member)
}

// OnFocus This event is fired when the game has focus (typically on page show).
func (self *Game) OnFocus() *Signal{
    return &Signal{self.Object.Get("onFocus")}
}

// SetOnFocusA This event is fired when the game has focus (typically on page show).
func (self *Game) SetOnFocusA(member *Signal) {
    self.Object.Set("onFocus", member)
}

// CurrentUpdateID The ID of the current/last logic update applied this render frame, starting from 0.
// The first update is `currentUpdateID === 0` and the last update is `currentUpdateID === updatesThisFrame.`
func (self *Game) CurrentUpdateID() int{
    return self.Object.Get("currentUpdateID").Int()
}

// SetCurrentUpdateIDA The ID of the current/last logic update applied this render frame, starting from 0.
// The first update is `currentUpdateID === 0` and the last update is `currentUpdateID === updatesThisFrame.`
func (self *Game) SetCurrentUpdateIDA(member int) {
    self.Object.Set("currentUpdateID", member)
}

// UpdatesThisFrame Number of logic updates expected to occur this render frame; will be 1 unless there are catch-ups required (and allowed).
func (self *Game) UpdatesThisFrame() int{
    return self.Object.Get("updatesThisFrame").Int()
}

// SetUpdatesThisFrameA Number of logic updates expected to occur this render frame; will be 1 unless there are catch-ups required (and allowed).
func (self *Game) SetUpdatesThisFrameA(member int) {
    self.Object.Set("updatesThisFrame", member)
}

// FpsProblemNotifier If the game is struggling to maintain the desired FPS, this signal will be dispatched.
// The desired/chosen FPS should probably be closer to the {@link Phaser.Time#suggestedFps} value.
func (self *Game) FpsProblemNotifier() *Signal{
    return &Signal{self.Object.Get("fpsProblemNotifier")}
}

// SetFpsProblemNotifierA If the game is struggling to maintain the desired FPS, this signal will be dispatched.
// The desired/chosen FPS should probably be closer to the {@link Phaser.Time#suggestedFps} value.
func (self *Game) SetFpsProblemNotifierA(member *Signal) {
    self.Object.Set("fpsProblemNotifier", member)
}

// ForceSingleUpdate Should the game loop force a logic update, regardless of the delta timer? Set to true if you know you need this. You can toggle it on the fly.
func (self *Game) ForceSingleUpdate() bool{
    return self.Object.Get("forceSingleUpdate").Bool()
}

// SetForceSingleUpdateA Should the game loop force a logic update, regardless of the delta timer? Set to true if you know you need this. You can toggle it on the fly.
func (self *Game) SetForceSingleUpdateA(member bool) {
    self.Object.Set("forceSingleUpdate", member)
}

// Paused The paused state of the Game. A paused game doesn't update any of its subsystems.
// When a game is paused the onPause event is dispatched. When it is resumed the onResume event is dispatched. Gets and sets the paused state of the Game.
func (self *Game) Paused() bool{
    return self.Object.Get("paused").Bool()
}

// SetPausedA The paused state of the Game. A paused game doesn't update any of its subsystems.
// When a game is paused the onPause event is dispatched. When it is resumed the onResume event is dispatched. Gets and sets the paused state of the Game.
func (self *Game) SetPausedA(member bool) {
    self.Object.Set("paused", member)
}


// ParseConfig Parses a Game configuration object.
func (self *Game) ParseConfig() {
    self.Object.Call("parseConfig")
}

// ParseConfigI Parses a Game configuration object.
func (self *Game) ParseConfigI(args ...interface{}) {
    self.Object.Call("parseConfig", args)
}

// Boot Initialize engine sub modules and start the game.
func (self *Game) Boot() {
    self.Object.Call("boot")
}

// BootI Initialize engine sub modules and start the game.
func (self *Game) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// ShowDebugHeader Displays a Phaser version debug header in the console.
func (self *Game) ShowDebugHeader() {
    self.Object.Call("showDebugHeader")
}

// ShowDebugHeaderI Displays a Phaser version debug header in the console.
func (self *Game) ShowDebugHeaderI(args ...interface{}) {
    self.Object.Call("showDebugHeader", args)
}

// SetUpRenderer Checks if the device is capable of using the requested renderer and sets it up or an alternative if not.
func (self *Game) SetUpRenderer() {
    self.Object.Call("setUpRenderer")
}

// SetUpRendererI Checks if the device is capable of using the requested renderer and sets it up or an alternative if not.
func (self *Game) SetUpRendererI(args ...interface{}) {
    self.Object.Call("setUpRenderer", args)
}

// ContextLost Handles WebGL context loss.
func (self *Game) ContextLost(event *Event) {
    self.Object.Call("contextLost", event)
}

// ContextLostI Handles WebGL context loss.
func (self *Game) ContextLostI(args ...interface{}) {
    self.Object.Call("contextLost", args)
}

// ContextRestored Handles WebGL context restoration.
func (self *Game) ContextRestored() {
    self.Object.Call("contextRestored")
}

// ContextRestoredI Handles WebGL context restoration.
func (self *Game) ContextRestoredI(args ...interface{}) {
    self.Object.Call("contextRestored", args)
}

// Update The core game loop.
func (self *Game) Update(time int) {
    self.Object.Call("update", time)
}

// UpdateI The core game loop.
func (self *Game) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// UpdateLogic Updates all logic subsystems in Phaser. Called automatically by Game.update.
func (self *Game) UpdateLogic(timeStep int) {
    self.Object.Call("updateLogic", timeStep)
}

// UpdateLogicI Updates all logic subsystems in Phaser. Called automatically by Game.update.
func (self *Game) UpdateLogicI(args ...interface{}) {
    self.Object.Call("updateLogic", args)
}

// UpdateRender Runs the Render cycle.
// It starts by calling State.preRender. In here you can do any last minute adjustments of display objects as required.
// It then calls the renderer, which renders the entire display list, starting from the Stage object and working down.
// It then calls plugin.render on any loaded plugins, in the order in which they were enabled.
// After this State.render is called. Any rendering that happens here will take place on-top of the display list.
// Finally plugin.postRender is called on any loaded plugins, in the order in which they were enabled.
// This method is called automatically by Game.update, you don't need to call it directly.
// Should you wish to have fine-grained control over when Phaser renders then use the `Game.lockRender` boolean.
// Phaser will only render when this boolean is `false`.
func (self *Game) UpdateRender(elapsedTime int) {
    self.Object.Call("updateRender", elapsedTime)
}

// UpdateRenderI Runs the Render cycle.
// It starts by calling State.preRender. In here you can do any last minute adjustments of display objects as required.
// It then calls the renderer, which renders the entire display list, starting from the Stage object and working down.
// It then calls plugin.render on any loaded plugins, in the order in which they were enabled.
// After this State.render is called. Any rendering that happens here will take place on-top of the display list.
// Finally plugin.postRender is called on any loaded plugins, in the order in which they were enabled.
// This method is called automatically by Game.update, you don't need to call it directly.
// Should you wish to have fine-grained control over when Phaser renders then use the `Game.lockRender` boolean.
// Phaser will only render when this boolean is `false`.
func (self *Game) UpdateRenderI(args ...interface{}) {
    self.Object.Call("updateRender", args)
}

// EnableStep Enable core game loop stepping. When enabled you must call game.step() directly (perhaps via a DOM button?)
// Calling step will advance the game loop by one frame. This is extremely useful for hard to track down errors!
func (self *Game) EnableStep() {
    self.Object.Call("enableStep")
}

// EnableStepI Enable core game loop stepping. When enabled you must call game.step() directly (perhaps via a DOM button?)
// Calling step will advance the game loop by one frame. This is extremely useful for hard to track down errors!
func (self *Game) EnableStepI(args ...interface{}) {
    self.Object.Call("enableStep", args)
}

// DisableStep Disables core game loop stepping.
func (self *Game) DisableStep() {
    self.Object.Call("disableStep")
}

// DisableStepI Disables core game loop stepping.
func (self *Game) DisableStepI(args ...interface{}) {
    self.Object.Call("disableStep", args)
}

// Step When stepping is enabled you must call this function directly (perhaps via a DOM button?) to advance the game loop by one frame.
// This is extremely useful to hard to track down errors! Use the internal stepCount property to monitor progress.
func (self *Game) Step() {
    self.Object.Call("step")
}

// StepI When stepping is enabled you must call this function directly (perhaps via a DOM button?) to advance the game loop by one frame.
// This is extremely useful to hard to track down errors! Use the internal stepCount property to monitor progress.
func (self *Game) StepI(args ...interface{}) {
    self.Object.Call("step", args)
}

// Destroy Nukes the entire game from orbit.
// 
// Calls destroy on Game.state, Game.sound, Game.scale, Game.stage, Game.input, Game.physics and Game.plugins.
// 
// Then sets all of those local handlers to null, destroys the renderer, removes the canvas from the DOM
// and resets the PIXI default renderer.
func (self *Game) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Nukes the entire game from orbit.
// 
// Calls destroy on Game.state, Game.sound, Game.scale, Game.stage, Game.input, Game.physics and Game.plugins.
// 
// Then sets all of those local handlers to null, destroys the renderer, removes the canvas from the DOM
// and resets the PIXI default renderer.
func (self *Game) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// GamePaused Called by the Stage visibility handler.
func (self *Game) GamePaused(event interface{}) {
    self.Object.Call("gamePaused", event)
}

// GamePausedI Called by the Stage visibility handler.
func (self *Game) GamePausedI(args ...interface{}) {
    self.Object.Call("gamePaused", args)
}

// GameResumed Called by the Stage visibility handler.
func (self *Game) GameResumed(event interface{}) {
    self.Object.Call("gameResumed", event)
}

// GameResumedI Called by the Stage visibility handler.
func (self *Game) GameResumedI(args ...interface{}) {
    self.Object.Call("gameResumed", args)
}

// FocusLoss Called by the Stage visibility handler.
func (self *Game) FocusLoss(event interface{}) {
    self.Object.Call("focusLoss", event)
}

// FocusLossI Called by the Stage visibility handler.
func (self *Game) FocusLossI(args ...interface{}) {
    self.Object.Call("focusLoss", args)
}

// FocusGain Called by the Stage visibility handler.
func (self *Game) FocusGain(event interface{}) {
    self.Object.Call("focusGain", event)
}

// FocusGainI Called by the Stage visibility handler.
func (self *Game) FocusGainI(args ...interface{}) {
    self.Object.Call("focusGain", args)
}

