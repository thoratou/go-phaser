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
func (self *Game) GetId() float64{
    return self.Get("id").Float()
}

// Phaser Game ID (for when Pixi supports multiple instances).
func (self *Game) SetId(member float64) {
    self.Set("id", member)
}

// The Phaser.Game configuration object.
func (self *Game) GetConfig() interface{}{
    return self.Get("config")
}

// The Phaser.Game configuration object.
func (self *Game) SetConfig(member interface{}) {
    self.Set("config", member)
}

// The Phaser.Physics.World configuration object.
func (self *Game) GetPhysicsConfig() interface{}{
    return self.Get("physicsConfig")
}

// The Phaser.Physics.World configuration object.
func (self *Game) SetPhysicsConfig(member interface{}) {
    self.Set("physicsConfig", member)
}

// The Games DOM parent.
func (self *Game) GetParent() interface{}{
    return self.Get("parent")
}

// The Games DOM parent.
func (self *Game) SetParent(member interface{}) {
    self.Set("parent", member)
}

// The current Game Width in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) GetWidth() int{
    return self.Get("width").Int()
}

// The current Game Width in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) SetWidth(member int) {
    self.Set("width", member)
}

// The current Game Height in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) GetHeight() int{
    return self.Get("height").Int()
}

// The current Game Height in pixels.
// 
// _Do not modify this property directly:_ use {@link Phaser.ScaleManager#setGameSize} - eg. `game.scale.setGameSize(width, height)` - instead.
func (self *Game) SetHeight(member int) {
    self.Set("height", member)
}

// The resolution of your game. This value is read only, but can be changed at start time it via a game configuration object.
func (self *Game) GetResolution() int{
    return self.Get("resolution").Int()
}

// The resolution of your game. This value is read only, but can be changed at start time it via a game configuration object.
func (self *Game) SetResolution(member int) {
    self.Set("resolution", member)
}

// Use a transparent canvas background or not.
func (self *Game) GetTransparent() bool{
    return self.Get("transparent").Bool()
}

// Use a transparent canvas background or not.
func (self *Game) SetTransparent(member bool) {
    self.Set("transparent", member)
}

// Anti-alias graphics. By default scaled images are smoothed in Canvas and WebGL, set anti-alias to false to disable this globally.
func (self *Game) GetAntialias() bool{
    return self.Get("antialias").Bool()
}

// Anti-alias graphics. By default scaled images are smoothed in Canvas and WebGL, set anti-alias to false to disable this globally.
func (self *Game) SetAntialias(member bool) {
    self.Set("antialias", member)
}

// The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
func (self *Game) GetPreserveDrawingBuffer() bool{
    return self.Get("preserveDrawingBuffer").Bool()
}

// The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
func (self *Game) SetPreserveDrawingBuffer(member bool) {
    self.Set("preserveDrawingBuffer", member)
}

// Clear the Canvas each frame before rendering the display list.
// You can set this to `false` to gain some performance if your game always contains a background that completely fills the display.
func (self *Game) GetClearBeforeRender() bool{
    return self.Get("clearBeforeRender").Bool()
}

// Clear the Canvas each frame before rendering the display list.
// You can set this to `false` to gain some performance if your game always contains a background that completely fills the display.
func (self *Game) SetClearBeforeRender(member bool) {
    self.Set("clearBeforeRender", member)
}

// The Pixi Renderer.
func (self *Game) GetRenderer() interface{}{
    return self.Get("renderer")
}

// The Pixi Renderer.
func (self *Game) SetRenderer(member interface{}) {
    self.Set("renderer", member)
}

// The Renderer this game will use. Either Phaser.AUTO, Phaser.CANVAS, Phaser.WEBGL, or Phaser.HEADLESS.
func (self *Game) GetRenderType() float64{
    return self.Get("renderType").Float()
}

// The Renderer this game will use. Either Phaser.AUTO, Phaser.CANVAS, Phaser.WEBGL, or Phaser.HEADLESS.
func (self *Game) SetRenderType(member float64) {
    self.Set("renderType", member)
}

// The StateManager.
func (self *Game) GetState() StateManager{
    return StateManager{self.Get("state")}
}

// The StateManager.
func (self *Game) SetState(member StateManager) {
    self.Set("state", member)
}

// Whether the game engine is booted, aka available.
func (self *Game) GetIsBooted() bool{
    return self.Get("isBooted").Bool()
}

// Whether the game engine is booted, aka available.
func (self *Game) SetIsBooted(member bool) {
    self.Set("isBooted", member)
}

// Is game running or paused?
func (self *Game) GetIsRunning() bool{
    return self.Get("isRunning").Bool()
}

// Is game running or paused?
func (self *Game) SetIsRunning(member bool) {
    self.Set("isRunning", member)
}

// Automatically handles the core game loop via requestAnimationFrame or setTimeout
func (self *Game) GetRaf() RequestAnimationFrame{
    return RequestAnimationFrame{self.Get("raf")}
}

// Automatically handles the core game loop via requestAnimationFrame or setTimeout
func (self *Game) SetRaf(member RequestAnimationFrame) {
    self.Set("raf", member)
}

// Reference to the Phaser.GameObjectFactory.
func (self *Game) GetAdd() GameObjectFactory{
    return GameObjectFactory{self.Get("add")}
}

// Reference to the Phaser.GameObjectFactory.
func (self *Game) SetAdd(member GameObjectFactory) {
    self.Set("add", member)
}

// Reference to the GameObject Creator.
func (self *Game) GetMake() GameObjectCreator{
    return GameObjectCreator{self.Get("make")}
}

// Reference to the GameObject Creator.
func (self *Game) SetMake(member GameObjectCreator) {
    self.Set("make", member)
}

// Reference to the assets cache.
func (self *Game) GetCache() Cache{
    return Cache{self.Get("cache")}
}

// Reference to the assets cache.
func (self *Game) SetCache(member Cache) {
    self.Set("cache", member)
}

// Reference to the input manager
func (self *Game) GetInput() Input{
    return Input{self.Get("input")}
}

// Reference to the input manager
func (self *Game) SetInput(member Input) {
    self.Set("input", member)
}

// Reference to the assets loader.
func (self *Game) GetLoad() Loader{
    return Loader{self.Get("load")}
}

// Reference to the assets loader.
func (self *Game) SetLoad(member Loader) {
    self.Set("load", member)
}

// Reference to the math helper.
func (self *Game) GetMath() Math{
    return Math{self.Get("math")}
}

// Reference to the math helper.
func (self *Game) SetMath(member Math) {
    self.Set("math", member)
}

// Reference to the network class.
func (self *Game) GetNet() Net{
    return Net{self.Get("net")}
}

// Reference to the network class.
func (self *Game) SetNet(member Net) {
    self.Set("net", member)
}

// The game scale manager.
func (self *Game) GetScale() ScaleManager{
    return ScaleManager{self.Get("scale")}
}

// The game scale manager.
func (self *Game) SetScale(member ScaleManager) {
    self.Set("scale", member)
}

// Reference to the sound manager.
func (self *Game) GetSound() SoundManager{
    return SoundManager{self.Get("sound")}
}

// Reference to the sound manager.
func (self *Game) SetSound(member SoundManager) {
    self.Set("sound", member)
}

// Reference to the stage.
func (self *Game) GetStage() Stage{
    return Stage{self.Get("stage")}
}

// Reference to the stage.
func (self *Game) SetStage(member Stage) {
    self.Set("stage", member)
}

// Reference to the core game clock.
func (self *Game) GetTime() Time{
    return Time{self.Get("time")}
}

// Reference to the core game clock.
func (self *Game) SetTime(member Time) {
    self.Set("time", member)
}

// Reference to the tween manager.
func (self *Game) GetTweens() TweenManager{
    return TweenManager{self.Get("tweens")}
}

// Reference to the tween manager.
func (self *Game) SetTweens(member TweenManager) {
    self.Set("tweens", member)
}

// Reference to the world.
func (self *Game) GetWorld() World{
    return World{self.Get("world")}
}

// Reference to the world.
func (self *Game) SetWorld(member World) {
    self.Set("world", member)
}

// Reference to the physics manager.
func (self *Game) GetPhysics() Physics{
    return Physics{self.Get("physics")}
}

// Reference to the physics manager.
func (self *Game) SetPhysics(member Physics) {
    self.Set("physics", member)
}

// Reference to the plugin manager.
func (self *Game) GetPlugins() PluginManager{
    return PluginManager{self.Get("plugins")}
}

// Reference to the plugin manager.
func (self *Game) SetPlugins(member PluginManager) {
    self.Set("plugins", member)
}

// Instance of repeatable random data generator helper.
func (self *Game) GetRnd() RandomDataGenerator{
    return RandomDataGenerator{self.Get("rnd")}
}

// Instance of repeatable random data generator helper.
func (self *Game) SetRnd(member RandomDataGenerator) {
    self.Set("rnd", member)
}

// Contains device information and capabilities.
func (self *Game) GetDevice() Device{
    return Device{self.Get("device")}
}

// Contains device information and capabilities.
func (self *Game) SetDevice(member Device) {
    self.Set("device", member)
}

// A handy reference to world.camera.
func (self *Game) GetCamera() Camera{
    return Camera{self.Get("camera")}
}

// A handy reference to world.camera.
func (self *Game) SetCamera(member Camera) {
    self.Set("camera", member)
}

// A handy reference to renderer.view, the canvas that the game is being rendered in to.
func (self *Game) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("canvas"))
}

// A handy reference to renderer.view, the canvas that the game is being rendered in to.
func (self *Game) SetCanvas(member dom.HTMLCanvasElement) {
    self.Set("canvas", member)
}

// A handy reference to renderer.context (only set for CANVAS games, not WebGL)
func (self *Game) GetContext() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Get("context"))
}

// A handy reference to renderer.context (only set for CANVAS games, not WebGL)
func (self *Game) SetContext(member dom.CanvasRenderingContext2D) {
    self.Set("context", member)
}

// A set of useful debug utilities.
func (self *Game) GetDebug() UtilsDebug{
    return UtilsDebug{self.Get("debug")}
}

// A set of useful debug utilities.
func (self *Game) SetDebug(member UtilsDebug) {
    self.Set("debug", member)
}

// The Particle Manager.
func (self *Game) GetParticles() Particles{
    return Particles{self.Get("particles")}
}

// The Particle Manager.
func (self *Game) SetParticles(member Particles) {
    self.Set("particles", member)
}

// The Asset Generator.
func (self *Game) GetCreate() Create{
    return Create{self.Get("create")}
}

// The Asset Generator.
func (self *Game) SetCreate(member Create) {
    self.Set("create", member)
}

// If `false` Phaser will automatically render the display list every update. If `true` the render loop will be skipped.
// You can toggle this value at run-time to gain exact control over when Phaser renders. This can be useful in certain types of game or application.
// Please note that if you don't render the display list then none of the game object transforms will be updated, so use this value carefully.
func (self *Game) GetLockRender() bool{
    return self.Get("lockRender").Bool()
}

// If `false` Phaser will automatically render the display list every update. If `true` the render loop will be skipped.
// You can toggle this value at run-time to gain exact control over when Phaser renders. This can be useful in certain types of game or application.
// Please note that if you don't render the display list then none of the game object transforms will be updated, so use this value carefully.
func (self *Game) SetLockRender(member bool) {
    self.Set("lockRender", member)
}

// Enable core loop stepping with Game.enableStep().
func (self *Game) GetStepping() bool{
    return self.Get("stepping").Bool()
}

// Enable core loop stepping with Game.enableStep().
func (self *Game) SetStepping(member bool) {
    self.Set("stepping", member)
}

// An internal property used by enableStep, but also useful to query from your own game objects.
func (self *Game) GetPendingStep() bool{
    return self.Get("pendingStep").Bool()
}

// An internal property used by enableStep, but also useful to query from your own game objects.
func (self *Game) SetPendingStep(member bool) {
    self.Set("pendingStep", member)
}

// When stepping is enabled this contains the current step cycle.
func (self *Game) GetStepCount() float64{
    return self.Get("stepCount").Float()
}

// When stepping is enabled this contains the current step cycle.
func (self *Game) SetStepCount(member float64) {
    self.Set("stepCount", member)
}

// This event is fired when the game pauses.
func (self *Game) GetOnPause() Signal{
    return Signal{self.Get("onPause")}
}

// This event is fired when the game pauses.
func (self *Game) SetOnPause(member Signal) {
    self.Set("onPause", member)
}

// This event is fired when the game resumes from a paused state.
func (self *Game) GetOnResume() Signal{
    return Signal{self.Get("onResume")}
}

// This event is fired when the game resumes from a paused state.
func (self *Game) SetOnResume(member Signal) {
    self.Set("onResume", member)
}

// This event is fired when the game no longer has focus (typically on page hide).
func (self *Game) GetOnBlur() Signal{
    return Signal{self.Get("onBlur")}
}

// This event is fired when the game no longer has focus (typically on page hide).
func (self *Game) SetOnBlur(member Signal) {
    self.Set("onBlur", member)
}

// This event is fired when the game has focus (typically on page show).
func (self *Game) GetOnFocus() Signal{
    return Signal{self.Get("onFocus")}
}

// This event is fired when the game has focus (typically on page show).
func (self *Game) SetOnFocus(member Signal) {
    self.Set("onFocus", member)
}

// The ID of the current/last logic update applied this render frame, starting from 0.
// The first update is `currentUpdateID === 0` and the last update is `currentUpdateID === updatesThisFrame.`
func (self *Game) GetCurrentUpdateID() int{
    return self.Get("currentUpdateID").Int()
}

// The ID of the current/last logic update applied this render frame, starting from 0.
// The first update is `currentUpdateID === 0` and the last update is `currentUpdateID === updatesThisFrame.`
func (self *Game) SetCurrentUpdateID(member int) {
    self.Set("currentUpdateID", member)
}

// Number of logic updates expected to occur this render frame; will be 1 unless there are catch-ups required (and allowed).
func (self *Game) GetUpdatesThisFrame() int{
    return self.Get("updatesThisFrame").Int()
}

// Number of logic updates expected to occur this render frame; will be 1 unless there are catch-ups required (and allowed).
func (self *Game) SetUpdatesThisFrame(member int) {
    self.Set("updatesThisFrame", member)
}

// If the game is struggling to maintain the desired FPS, this signal will be dispatched.
// The desired/chosen FPS should probably be closer to the {@link Phaser.Time#suggestedFps} value.
func (self *Game) GetFpsProblemNotifier() Signal{
    return Signal{self.Get("fpsProblemNotifier")}
}

// If the game is struggling to maintain the desired FPS, this signal will be dispatched.
// The desired/chosen FPS should probably be closer to the {@link Phaser.Time#suggestedFps} value.
func (self *Game) SetFpsProblemNotifier(member Signal) {
    self.Set("fpsProblemNotifier", member)
}

// Should the game loop force a logic update, regardless of the delta timer? Set to true if you know you need this. You can toggle it on the fly.
func (self *Game) GetForceSingleUpdate() bool{
    return self.Get("forceSingleUpdate").Bool()
}

// Should the game loop force a logic update, regardless of the delta timer? Set to true if you know you need this. You can toggle it on the fly.
func (self *Game) SetForceSingleUpdate(member bool) {
    self.Set("forceSingleUpdate", member)
}

// The paused state of the Game. A paused game doesn't update any of its subsystems.
// When a game is paused the onPause event is dispatched. When it is resumed the onResume event is dispatched. Gets and sets the paused state of the Game.
func (self *Game) GetPaused() bool{
    return self.Get("paused").Bool()
}

// The paused state of the Game. A paused game doesn't update any of its subsystems.
// When a game is paused the onPause event is dispatched. When it is resumed the onResume event is dispatched. Gets and sets the paused state of the Game.
func (self *Game) SetPaused(member bool) {
    self.Set("paused", member)
}



// Parses a Game configuration object.
func (self *Game) ParseConfigI(args ...interface{}) {
    self.Call("parseConfig", args)
}

// Initialize engine sub modules and start the game.
func (self *Game) BootI(args ...interface{}) {
    self.Call("boot", args)
}

// Displays a Phaser version debug header in the console.
func (self *Game) ShowDebugHeaderI(args ...interface{}) {
    self.Call("showDebugHeader", args)
}

// Checks if the device is capable of using the requested renderer and sets it up or an alternative if not.
func (self *Game) SetUpRendererI(args ...interface{}) {
    self.Call("setUpRenderer", args)
}

// Handles WebGL context loss.
func (self *Game) ContextLostI(args ...interface{}) {
    self.Call("contextLost", args)
}

// Handles WebGL context restoration.
func (self *Game) ContextRestoredI(args ...interface{}) {
    self.Call("contextRestored", args)
}

// The core game loop.
func (self *Game) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Updates all logic subsystems in Phaser. Called automatically by Game.update.
func (self *Game) UpdateLogicI(args ...interface{}) {
    self.Call("updateLogic", args)
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
    self.Call("updateRender", args)
}

// Enable core game loop stepping. When enabled you must call game.step() directly (perhaps via a DOM button?)
// Calling step will advance the game loop by one frame. This is extremely useful for hard to track down errors!
func (self *Game) EnableStepI(args ...interface{}) {
    self.Call("enableStep", args)
}

// Disables core game loop stepping.
func (self *Game) DisableStepI(args ...interface{}) {
    self.Call("disableStep", args)
}

// When stepping is enabled you must call this function directly (perhaps via a DOM button?) to advance the game loop by one frame.
// This is extremely useful to hard to track down errors! Use the internal stepCount property to monitor progress.
func (self *Game) StepI(args ...interface{}) {
    self.Call("step", args)
}

// Nukes the entire game from orbit.
// 
// Calls destroy on Game.state, Game.sound, Game.scale, Game.stage, Game.input, Game.physics and Game.plugins.
// 
// Then sets all of those local handlers to null, destroys the renderer, removes the canvas from the DOM
// and resets the PIXI default renderer.
func (self *Game) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Called by the Stage visibility handler.
func (self *Game) GamePausedI(args ...interface{}) {
    self.Call("gamePaused", args)
}

// Called by the Stage visibility handler.
func (self *Game) GameResumedI(args ...interface{}) {
    self.Call("gameResumed", args)
}

// Called by the Stage visibility handler.
func (self *Game) FocusLossI(args ...interface{}) {
    self.Call("focusLoss", args)
}

// Called by the Stage visibility handler.
func (self *Game) FocusGainI(args ...interface{}) {
    self.Call("focusGain", args)
}
