// Automatic generation for Phaser.Game
// generated file Game.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// This is where the magic happens. The Game object is the heart of your game,
// providing quick access to common functions and handling the boot process.
// 
// "Hell, there are no rules here - we're trying to accomplish something."
//                                                       Thomas A. Edison
type Game struct {
    *js.Object
}


// Phaser Game ID (for when Pixi supports multiple instances).
func (self *Game) GetIdA() int{
    return self.Object.Get("id").Int()
}

// Phaser Game ID (for when Pixi supports multiple instances).
func (self *Game) SetIdA(member int) {
    self.Object.Set("id", member)
}

// The Phaser.Game configuration object.
func (self *Game) GetConfigA() interface{}{
    return self.Object.Get("config")
}

// The Phaser.Game configuration object.
func (self *Game) SetConfigA(member interface{}) {
    self.Object.Set("config", member)
}

// The Phaser.Physics.World configuration object.
func (self *Game) GetPhysicsConfigA() interface{}{
    return self.Object.Get("physicsConfig")
}

// The Phaser.Physics.World configuration object.
func (self *Game) SetPhysicsConfigA(member interface{}) {
    self.Object.Set("physicsConfig", member)
}

// The Games DOM parent.
func (self *Game) GetParentA() interface{}{
    return self.Object.Get("parent")
}

// The Games DOM parent.
func (self *Game) SetParentA(member interface{}) {
    self.Object.Set("parent", member)
}

// The current Game Width in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The current Game Width in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The current Game Height in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The current Game Height in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The resolution of your game. This value is read only, but can be changed at start time it via a game configuration object.
func (self *Game) GetResolutionA() int{
    return self.Object.Get("resolution").Int()
}

// The resolution of your game. This value is read only, but can be changed at start time it via a game configuration object.
func (self *Game) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// Use a transparent canvas background or not.
func (self *Game) GetTransparentA() bool{
    return self.Object.Get("transparent").Bool()
}

// Use a transparent canvas background or not.
func (self *Game) SetTransparentA(member bool) {
    self.Object.Set("transparent", member)
}

// Anti-alias graphics. By default scaled images are smoothed in Canvas and WebGL, set anti-alias to false to disable this globally.
func (self *Game) GetAntialiasA() bool{
    return self.Object.Get("antialias").Bool()
}

// Anti-alias graphics. By default scaled images are smoothed in Canvas and WebGL, set anti-alias to false to disable this globally.
func (self *Game) SetAntialiasA(member bool) {
    self.Object.Set("antialias", member)
}

// The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
func (self *Game) GetPreserveDrawingBufferA() bool{
    return self.Object.Get("preserveDrawingBuffer").Bool()
}

// The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
func (self *Game) SetPreserveDrawingBufferA(member bool) {
    self.Object.Set("preserveDrawingBuffer", member)
}

// Clear the Canvas each frame before rendering the display list.
// You can set this to `false` to gain some performance if your game always contains a background that completely fills the display.
func (self *Game) GetClearBeforeRenderA() bool{
    return self.Object.Get("clearBeforeRender").Bool()
}

// Clear the Canvas each frame before rendering the display list.
// You can set this to `false` to gain some performance if your game always contains a background that completely fills the display.
func (self *Game) SetClearBeforeRenderA(member bool) {
    self.Object.Set("clearBeforeRender", member)
}

// The Pixi Renderer.
func (self *Game) GetRendererA() interface{}{
    return self.Object.Get("renderer")
}

// The Pixi Renderer.
func (self *Game) SetRendererA(member interface{}) {
    self.Object.Set("renderer", member)
}

// The Renderer this game will use. Either Phaser.AUTO, Phaser.CANVAS, Phaser.WEBGL, or Phaser.HEADLESS.
func (self *Game) GetRenderTypeA() int{
    return self.Object.Get("renderType").Int()
}

// The Renderer this game will use. Either Phaser.AUTO, Phaser.CANVAS, Phaser.WEBGL, or Phaser.HEADLESS.
func (self *Game) SetRenderTypeA(member int) {
    self.Object.Set("renderType", member)
}

// The StateManager.
func (self *Game) GetStateA() *StateManager{
    return &StateManager{self.Object.Get("state")}
}

// The StateManager.
func (self *Game) SetStateA(member *StateManager) {
    self.Object.Set("state", member)
}

// Whether the game engine is booted, aka available.
func (self *Game) GetIsBootedA() bool{
    return self.Object.Get("isBooted").Bool()
}

// Whether the game engine is booted, aka available.
func (self *Game) SetIsBootedA(member bool) {
    self.Object.Set("isBooted", member)
}

// Is game running or paused?
func (self *Game) GetIsRunningA() bool{
    return self.Object.Get("isRunning").Bool()
}

// Is game running or paused?
func (self *Game) SetIsRunningA(member bool) {
    self.Object.Set("isRunning", member)
}

// Automatically handles the core game loop via requestAnimationFrame or setTimeout
func (self *Game) GetRafA() *RequestAnimationFrame{
    return &RequestAnimationFrame{self.Object.Get("raf")}
}

// Automatically handles the core game loop via requestAnimationFrame or setTimeout
func (self *Game) SetRafA(member *RequestAnimationFrame) {
    self.Object.Set("raf", member)
}

// Reference to the Phaser.GameObjectFactory.
func (self *Game) GetAddA() *GameObjectFactory{
    return &GameObjectFactory{self.Object.Get("add")}
}

// Reference to the Phaser.GameObjectFactory.
func (self *Game) SetAddA(member *GameObjectFactory) {
    self.Object.Set("add", member)
}

// Reference to the GameObject Creator.
func (self *Game) GetMakeA() *GameObjectCreator{
    return &GameObjectCreator{self.Object.Get("make")}
}

// Reference to the GameObject Creator.
func (self *Game) SetMakeA(member *GameObjectCreator) {
    self.Object.Set("make", member)
}

// Reference to the assets cache.
func (self *Game) GetCacheA() *Cache{
    return &Cache{self.Object.Get("cache")}
}

// Reference to the assets cache.
func (self *Game) SetCacheA(member *Cache) {
    self.Object.Set("cache", member)
}

// Reference to the input manager
func (self *Game) GetInputA() *Input{
    return &Input{self.Object.Get("input")}
}

// Reference to the input manager
func (self *Game) SetInputA(member *Input) {
    self.Object.Set("input", member)
}

// Reference to the assets loader.
func (self *Game) GetLoadA() *Loader{
    return &Loader{self.Object.Get("load")}
}

// Reference to the assets loader.
func (self *Game) SetLoadA(member *Loader) {
    self.Object.Set("load", member)
}

// Reference to the math helper.
func (self *Game) GetMathA() *Math{
    return &Math{self.Object.Get("math")}
}

// Reference to the math helper.
func (self *Game) SetMathA(member *Math) {
    self.Object.Set("math", member)
}

// Reference to the network class.
func (self *Game) GetNetA() *Net{
    return &Net{self.Object.Get("net")}
}

// Reference to the network class.
func (self *Game) SetNetA(member *Net) {
    self.Object.Set("net", member)
}

// The game scale manager.
func (self *Game) GetScaleA() *ScaleManager{
    return &ScaleManager{self.Object.Get("scale")}
}

// The game scale manager.
func (self *Game) SetScaleA(member *ScaleManager) {
    self.Object.Set("scale", member)
}

// Reference to the sound manager.
func (self *Game) GetSoundA() *SoundManager{
    return &SoundManager{self.Object.Get("sound")}
}

// Reference to the sound manager.
func (self *Game) SetSoundA(member *SoundManager) {
    self.Object.Set("sound", member)
}

// Reference to the stage.
func (self *Game) GetStageA() *Stage{
    return &Stage{self.Object.Get("stage")}
}

// Reference to the stage.
func (self *Game) SetStageA(member *Stage) {
    self.Object.Set("stage", member)
}

// Reference to the core game clock.
func (self *Game) GetTimeA() *Time{
    return &Time{self.Object.Get("time")}
}

// Reference to the core game clock.
func (self *Game) SetTimeA(member *Time) {
    self.Object.Set("time", member)
}

// Reference to the tween manager.
func (self *Game) GetTweensA() *TweenManager{
    return &TweenManager{self.Object.Get("tweens")}
}

// Reference to the tween manager.
func (self *Game) SetTweensA(member *TweenManager) {
    self.Object.Set("tweens", member)
}

// Reference to the world.
func (self *Game) GetWorldA() *World{
    return &World{self.Object.Get("world")}
}

// Reference to the world.
func (self *Game) SetWorldA(member *World) {
    self.Object.Set("world", member)
}

// Reference to the physics manager.
func (self *Game) GetPhysicsA() *Physics{
    return &Physics{self.Object.Get("physics")}
}

// Reference to the physics manager.
func (self *Game) SetPhysicsA(member *Physics) {
    self.Object.Set("physics", member)
}

// Reference to the plugin manager.
func (self *Game) GetPluginsA() *PluginManager{
    return &PluginManager{self.Object.Get("plugins")}
}

// Reference to the plugin manager.
func (self *Game) SetPluginsA(member *PluginManager) {
    self.Object.Set("plugins", member)
}

// Instance of repeatable random data generator helper.
func (self *Game) GetRndA() *RandomDataGenerator{
    return &RandomDataGenerator{self.Object.Get("rnd")}
}

// Instance of repeatable random data generator helper.
func (self *Game) SetRndA(member *RandomDataGenerator) {
    self.Object.Set("rnd", member)
}

// Contains device information and capabilities.
func (self *Game) GetDeviceA() *Device{
    return &Device{self.Object.Get("device")}
}

// Contains device information and capabilities.
func (self *Game) SetDeviceA(member *Device) {
    self.Object.Set("device", member)
}

// A handy reference to world.camera.
func (self *Game) GetCameraA() *Camera{
    return &Camera{self.Object.Get("camera")}
}

// A handy reference to world.camera.
func (self *Game) SetCameraA(member *Camera) {
    self.Object.Set("camera", member)
}

// A handy reference to renderer.view, the canvas that the game is being rendered in to.
func (self *Game) GetCanvasA() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("canvas"))
}

// A handy reference to renderer.view, the canvas that the game is being rendered in to.
func (self *Game) SetCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("canvas", member)
}

// A handy reference to renderer.context (only set for CANVAS games, not WebGL)
func (self *Game) GetContextA() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("context"))
}

// A handy reference to renderer.context (only set for CANVAS games, not WebGL)
func (self *Game) SetContextA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("context", member)
}

// A set of useful debug utilities.
func (self *Game) GetDebugA() *UtilsDebug{
    return &UtilsDebug{self.Object.Get("debug")}
}

// A set of useful debug utilities.
func (self *Game) SetDebugA(member *UtilsDebug) {
    self.Object.Set("debug", member)
}

// The Particle Manager.
func (self *Game) GetParticlesA() *Particles{
    return &Particles{self.Object.Get("particles")}
}

// The Particle Manager.
func (self *Game) SetParticlesA(member *Particles) {
    self.Object.Set("particles", member)
}

// The Asset Generator.
func (self *Game) GetCreateA() *Create{
    return &Create{self.Object.Get("create")}
}

// The Asset Generator.
func (self *Game) SetCreateA(member *Create) {
    self.Object.Set("create", member)
}

// If `false` Phaser will automatically render the display list every update. If `true` the render loop will be skipped.
// You can toggle this value at run-time to gain exact control over when Phaser renders. This can be useful in certain types of game or application.
// Please note that if you don't render the display list then none of the game object transforms will be updated, so use this value carefully.
func (self *Game) GetLockRenderA() bool{
    return self.Object.Get("lockRender").Bool()
}

// If `false` Phaser will automatically render the display list every update. If `true` the render loop will be skipped.
// You can toggle this value at run-time to gain exact control over when Phaser renders. This can be useful in certain types of game or application.
// Please note that if you don't render the display list then none of the game object transforms will be updated, so use this value carefully.
func (self *Game) SetLockRenderA(member bool) {
    self.Object.Set("lockRender", member)
}

// Enable core loop stepping with Game.enableStep().
func (self *Game) GetSteppingA() bool{
    return self.Object.Get("stepping").Bool()
}

// Enable core loop stepping with Game.enableStep().
func (self *Game) SetSteppingA(member bool) {
    self.Object.Set("stepping", member)
}

// An internal property used by enableStep, but also useful to query from your own game objects.
func (self *Game) GetPendingStepA() bool{
    return self.Object.Get("pendingStep").Bool()
}

// An internal property used by enableStep, but also useful to query from your own game objects.
func (self *Game) SetPendingStepA(member bool) {
    self.Object.Set("pendingStep", member)
}

// When stepping is enabled this contains the current step cycle.
func (self *Game) GetStepCountA() int{
    return self.Object.Get("stepCount").Int()
}

// When stepping is enabled this contains the current step cycle.
func (self *Game) SetStepCountA(member int) {
    self.Object.Set("stepCount", member)
}

// This event is fired when the game pauses.
func (self *Game) GetOnPauseA() *Signal{
    return &Signal{self.Object.Get("onPause")}
}

// This event is fired when the game pauses.
func (self *Game) SetOnPauseA(member *Signal) {
    self.Object.Set("onPause", member)
}

// This event is fired when the game resumes from a paused state.
func (self *Game) GetOnResumeA() *Signal{
    return &Signal{self.Object.Get("onResume")}
}

// This event is fired when the game resumes from a paused state.
func (self *Game) SetOnResumeA(member *Signal) {
    self.Object.Set("onResume", member)
}

// This event is fired when the game no longer has focus (typically on page hide).
func (self *Game) GetOnBlurA() *Signal{
    return &Signal{self.Object.Get("onBlur")}
}

// This event is fired when the game no longer has focus (typically on page hide).
func (self *Game) SetOnBlurA(member *Signal) {
    self.Object.Set("onBlur", member)
}

// This event is fired when the game has focus (typically on page show).
func (self *Game) GetOnFocusA() *Signal{
    return &Signal{self.Object.Get("onFocus")}
}

// This event is fired when the game has focus (typically on page show).
func (self *Game) SetOnFocusA(member *Signal) {
    self.Object.Set("onFocus", member)
}

// The ID of the current/last logic update applied this render frame, starting from 0.
// The first update is `currentUpdateID === 0` and the last update is `currentUpdateID === updatesThisFrame.`
func (self *Game) GetCurrentUpdateIDA() int{
    return self.Object.Get("currentUpdateID").Int()
}

// The ID of the current/last logic update applied this render frame, starting from 0.
// The first update is `currentUpdateID === 0` and the last update is `currentUpdateID === updatesThisFrame.`
func (self *Game) SetCurrentUpdateIDA(member int) {
    self.Object.Set("currentUpdateID", member)
}

// Number of logic updates expected to occur this render frame; will be 1 unless there are catch-ups required (and allowed).
func (self *Game) GetUpdatesThisFrameA() int{
    return self.Object.Get("updatesThisFrame").Int()
}

// Number of logic updates expected to occur this render frame; will be 1 unless there are catch-ups required (and allowed).
func (self *Game) SetUpdatesThisFrameA(member int) {
    self.Object.Set("updatesThisFrame", member)
}

// If the game is struggling to maintain the desired FPS, this signal will be dispatched.
// The desired/chosen FPS should probably be closer to the {@link Phaser.Time#suggestedFps} value.
func (self *Game) GetFpsProblemNotifierA() *Signal{
    return &Signal{self.Object.Get("fpsProblemNotifier")}
}

// If the game is struggling to maintain the desired FPS, this signal will be dispatched.
// The desired/chosen FPS should probably be closer to the {@link Phaser.Time#suggestedFps} value.
func (self *Game) SetFpsProblemNotifierA(member *Signal) {
    self.Object.Set("fpsProblemNotifier", member)
}

// Should the game loop force a logic update, regardless of the delta timer? Set to true if you know you need this. You can toggle it on the fly.
func (self *Game) GetForceSingleUpdateA() bool{
    return self.Object.Get("forceSingleUpdate").Bool()
}

// Should the game loop force a logic update, regardless of the delta timer? Set to true if you know you need this. You can toggle it on the fly.
func (self *Game) SetForceSingleUpdateA(member bool) {
    self.Object.Set("forceSingleUpdate", member)
}

// The paused state of the Game. A paused game doesn't update any of its subsystems.
// When a game is paused the onPause event is dispatched. When it is resumed the onResume event is dispatched. Gets and sets the paused state of the Game.
func (self *Game) GetPausedA() bool{
    return self.Object.Get("paused").Bool()
}

// The paused state of the Game. A paused game doesn't update any of its subsystems.
// When a game is paused the onPause event is dispatched. When it is resumed the onResume event is dispatched. Gets and sets the paused state of the Game.
func (self *Game) SetPausedA(member bool) {
    self.Object.Set("paused", member)
}



// Parses a Game configuration object.
func (self *Game) ParseConfig() {
    self.Object.Call("parseConfig")
}

// Parses a Game configuration object.
func (self *Game) ParseConfigI(args ...interface{}) {
    self.Object.Call("parseConfig", args)
}

// Initialize engine sub modules and start the game.
func (self *Game) Boot() {
    self.Object.Call("boot")
}

// Initialize engine sub modules and start the game.
func (self *Game) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// Displays a Phaser version debug header in the console.
func (self *Game) ShowDebugHeader() {
    self.Object.Call("showDebugHeader")
}

// Displays a Phaser version debug header in the console.
func (self *Game) ShowDebugHeaderI(args ...interface{}) {
    self.Object.Call("showDebugHeader", args)
}

// Checks if the device is capable of using the requested renderer and sets it up or an alternative if not.
func (self *Game) SetUpRenderer() {
    self.Object.Call("setUpRenderer")
}

// Checks if the device is capable of using the requested renderer and sets it up or an alternative if not.
func (self *Game) SetUpRendererI(args ...interface{}) {
    self.Object.Call("setUpRenderer", args)
}

// Handles WebGL context loss.
func (self *Game) ContextLost(event *Event) {
    self.Object.Call("contextLost", event)
}

// Handles WebGL context loss.
func (self *Game) ContextLostI(args ...interface{}) {
    self.Object.Call("contextLost", args)
}

// Handles WebGL context restoration.
func (self *Game) ContextRestored() {
    self.Object.Call("contextRestored")
}

// Handles WebGL context restoration.
func (self *Game) ContextRestoredI(args ...interface{}) {
    self.Object.Call("contextRestored", args)
}

// The core game loop.
func (self *Game) Update(time int) {
    self.Object.Call("update", time)
}

// The core game loop.
func (self *Game) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Updates all logic subsystems in Phaser. Called automatically by Game.update.
func (self *Game) UpdateLogic(timeStep int) {
    self.Object.Call("updateLogic", timeStep)
}

// Updates all logic subsystems in Phaser. Called automatically by Game.update.
func (self *Game) UpdateLogicI(args ...interface{}) {
    self.Object.Call("updateLogic", args)
}

// Runs the Render cycle.
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

// Runs the Render cycle.
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

// Enable core game loop stepping. When enabled you must call game.step() directly (perhaps via a DOM button?)
// Calling step will advance the game loop by one frame. This is extremely useful for hard to track down errors!
func (self *Game) EnableStep() {
    self.Object.Call("enableStep")
}

// Enable core game loop stepping. When enabled you must call game.step() directly (perhaps via a DOM button?)
// Calling step will advance the game loop by one frame. This is extremely useful for hard to track down errors!
func (self *Game) EnableStepI(args ...interface{}) {
    self.Object.Call("enableStep", args)
}

// Disables core game loop stepping.
func (self *Game) DisableStep() {
    self.Object.Call("disableStep")
}

// Disables core game loop stepping.
func (self *Game) DisableStepI(args ...interface{}) {
    self.Object.Call("disableStep", args)
}

// When stepping is enabled you must call this function directly (perhaps via a DOM button?) to advance the game loop by one frame.
// This is extremely useful to hard to track down errors! Use the internal stepCount property to monitor progress.
func (self *Game) Step() {
    self.Object.Call("step")
}

// When stepping is enabled you must call this function directly (perhaps via a DOM button?) to advance the game loop by one frame.
// This is extremely useful to hard to track down errors! Use the internal stepCount property to monitor progress.
func (self *Game) StepI(args ...interface{}) {
    self.Object.Call("step", args)
}

// Nukes the entire game from orbit.
// 
// Calls destroy on Game.state, Game.sound, Game.scale, Game.stage, Game.input, Game.physics and Game.plugins.
// 
// Then sets all of those local handlers to null, destroys the renderer, removes the canvas from the DOM
// and resets the PIXI default renderer.
func (self *Game) Destroy() {
    self.Object.Call("destroy")
}

// Nukes the entire game from orbit.
// 
// Calls destroy on Game.state, Game.sound, Game.scale, Game.stage, Game.input, Game.physics and Game.plugins.
// 
// Then sets all of those local handlers to null, destroys the renderer, removes the canvas from the DOM
// and resets the PIXI default renderer.
func (self *Game) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Called by the Stage visibility handler.
func (self *Game) GamePaused(event interface{}) {
    self.Object.Call("gamePaused", event)
}

// Called by the Stage visibility handler.
func (self *Game) GamePausedI(args ...interface{}) {
    self.Object.Call("gamePaused", args)
}

// Called by the Stage visibility handler.
func (self *Game) GameResumed(event interface{}) {
    self.Object.Call("gameResumed", event)
}

// Called by the Stage visibility handler.
func (self *Game) GameResumedI(args ...interface{}) {
    self.Object.Call("gameResumed", args)
}

// Called by the Stage visibility handler.
func (self *Game) FocusLoss(event interface{}) {
    self.Object.Call("focusLoss", event)
}

// Called by the Stage visibility handler.
func (self *Game) FocusLossI(args ...interface{}) {
    self.Object.Call("focusLoss", args)
}

// Called by the Stage visibility handler.
func (self *Game) FocusGain(event interface{}) {
    self.Object.Call("focusGain", event)
}

// Called by the Stage visibility handler.
func (self *Game) FocusGainI(args ...interface{}) {
    self.Object.Call("focusGain", args)
}
