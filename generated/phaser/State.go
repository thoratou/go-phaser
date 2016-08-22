// Automatic generation for Phaser.State
// generated file State.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// This is a base State class which can be extended if you are creating your own game.
// It provides quick access to common functions such as the camera, cache, input, match, sound and more.
type State struct {
    *js.Object
}


// This is a base State class which can be extended if you are creating your own game.
// It provides quick access to common functions such as the camera, cache, input, match, sound and more.
func NewState() *State {
    return &State{js.Global.Call("Phaser.State")}
}

// This is a base State class which can be extended if you are creating your own game.
// It provides quick access to common functions such as the camera, cache, input, match, sound and more.
func NewStateI(args ...interface{}) *State {
    return &State{js.Global.Call("Phaser.State", args)}
}



// This is a reference to the currently running Game.
func (self *State) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// This is a reference to the currently running Game.
func (self *State) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The string based identifier given to the State when added into the State Manager.
func (self *State) GetKeyA() string{
    return self.Object.Get("key").String()
}

// The string based identifier given to the State when added into the State Manager.
func (self *State) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// A reference to the GameObjectFactory which can be used to add new objects to the World.
func (self *State) GetAddA() *GameObjectFactory{
    return &GameObjectFactory{self.Object.Get("add")}
}

// A reference to the GameObjectFactory which can be used to add new objects to the World.
func (self *State) SetAddA(member *GameObjectFactory) {
    self.Object.Set("add", member)
}

// A reference to the GameObjectCreator which can be used to make new objects.
func (self *State) GetMakeA() *GameObjectCreator{
    return &GameObjectCreator{self.Object.Get("make")}
}

// A reference to the GameObjectCreator which can be used to make new objects.
func (self *State) SetMakeA(member *GameObjectCreator) {
    self.Object.Set("make", member)
}

// A handy reference to World.camera.
func (self *State) GetCameraA() *Camera{
    return &Camera{self.Object.Get("camera")}
}

// A handy reference to World.camera.
func (self *State) SetCameraA(member *Camera) {
    self.Object.Set("camera", member)
}

// A reference to the game cache which contains any loaded or generated assets, such as images, sound and more.
func (self *State) GetCacheA() *Cache{
    return &Cache{self.Object.Get("cache")}
}

// A reference to the game cache which contains any loaded or generated assets, such as images, sound and more.
func (self *State) SetCacheA(member *Cache) {
    self.Object.Set("cache", member)
}

// A reference to the Input Manager.
func (self *State) GetInputA() *Input{
    return &Input{self.Object.Get("input")}
}

// A reference to the Input Manager.
func (self *State) SetInputA(member *Input) {
    self.Object.Set("input", member)
}

// A reference to the Loader, which you mostly use in the preload method of your state to load external assets.
func (self *State) GetLoadA() *Loader{
    return &Loader{self.Object.Get("load")}
}

// A reference to the Loader, which you mostly use in the preload method of your state to load external assets.
func (self *State) SetLoadA(member *Loader) {
    self.Object.Set("load", member)
}

// A reference to Math class with lots of helpful functions.
func (self *State) GetMathA() *Math{
    return &Math{self.Object.Get("math")}
}

// A reference to Math class with lots of helpful functions.
func (self *State) SetMathA(member *Math) {
    self.Object.Set("math", member)
}

// A reference to the Sound Manager which can create, play and stop sounds, as well as adjust global volume.
func (self *State) GetSoundA() *SoundManager{
    return &SoundManager{self.Object.Get("sound")}
}

// A reference to the Sound Manager which can create, play and stop sounds, as well as adjust global volume.
func (self *State) SetSoundA(member *SoundManager) {
    self.Object.Set("sound", member)
}

// A reference to the Scale Manager which controls the way the game scales on different displays.
func (self *State) GetScaleA() *ScaleManager{
    return &ScaleManager{self.Object.Get("scale")}
}

// A reference to the Scale Manager which controls the way the game scales on different displays.
func (self *State) SetScaleA(member *ScaleManager) {
    self.Object.Set("scale", member)
}

// A reference to the Stage.
func (self *State) GetStageA() *Stage{
    return &Stage{self.Object.Get("stage")}
}

// A reference to the Stage.
func (self *State) SetStageA(member *Stage) {
    self.Object.Set("stage", member)
}

// A reference to the State Manager, which controls state changes.
func (self *State) GetStateA() interface{}{
    return self.Object.Get("state")
}

// A reference to the State Manager, which controls state changes.
func (self *State) SetStateA(member interface{}) {
    self.Object.Set("state", member)
}

// A reference to the game clock and timed events system.
func (self *State) GetTimeA() *Time{
    return &Time{self.Object.Get("time")}
}

// A reference to the game clock and timed events system.
func (self *State) SetTimeA(member *Time) {
    self.Object.Set("time", member)
}

// A reference to the tween manager.
func (self *State) GetTweensA() *TweenManager{
    return &TweenManager{self.Object.Get("tweens")}
}

// A reference to the tween manager.
func (self *State) SetTweensA(member *TweenManager) {
    self.Object.Set("tweens", member)
}

// A reference to the game world. All objects live in the Game World and its size is not bound by the display resolution.
func (self *State) GetWorldA() *World{
    return &World{self.Object.Get("world")}
}

// A reference to the game world. All objects live in the Game World and its size is not bound by the display resolution.
func (self *State) SetWorldA(member *World) {
    self.Object.Set("world", member)
}

// The Particle Manager. It is called during the core gameloop and updates any Particle Emitters it has created.
func (self *State) GetParticlesA() *Particles{
    return &Particles{self.Object.Get("particles")}
}

// The Particle Manager. It is called during the core gameloop and updates any Particle Emitters it has created.
func (self *State) SetParticlesA(member *Particles) {
    self.Object.Set("particles", member)
}

// A reference to the physics manager which looks after the different physics systems available within Phaser.
func (self *State) GetPhysicsA() *Physics{
    return &Physics{self.Object.Get("physics")}
}

// A reference to the physics manager which looks after the different physics systems available within Phaser.
func (self *State) SetPhysicsA(member *Physics) {
    self.Object.Set("physics", member)
}

// A reference to the seeded and repeatable random data generator.
func (self *State) GetRndA() *RandomDataGenerator{
    return &RandomDataGenerator{self.Object.Get("rnd")}
}

// A reference to the seeded and repeatable random data generator.
func (self *State) SetRndA(member *RandomDataGenerator) {
    self.Object.Set("rnd", member)
}



// init is the very first function called when your State starts up. It's called before preload, create or anything else.
// If you need to route the game away to another State you could do so here, or if you need to prepare a set of variables
// or objects before the preloading starts.
func (self *State) Init() {
    self.Object.Call("init")
}

// init is the very first function called when your State starts up. It's called before preload, create or anything else.
// If you need to route the game away to another State you could do so here, or if you need to prepare a set of variables
// or objects before the preloading starts.
func (self *State) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// preload is called first. Normally you'd use this to load your game assets (or those needed for the current State)
// You shouldn't create any objects in this method that require assets that you're also loading in this method, as
// they won't yet be available.
func (self *State) Preload() {
    self.Object.Call("preload")
}

// preload is called first. Normally you'd use this to load your game assets (or those needed for the current State)
// You shouldn't create any objects in this method that require assets that you're also loading in this method, as
// they won't yet be available.
func (self *State) PreloadI(args ...interface{}) {
    self.Object.Call("preload", args)
}

// loadUpdate is called during the Loader process. This only happens if you've set one or more assets to load in the preload method.
func (self *State) LoadUpdate() {
    self.Object.Call("loadUpdate")
}

// loadUpdate is called during the Loader process. This only happens if you've set one or more assets to load in the preload method.
func (self *State) LoadUpdateI(args ...interface{}) {
    self.Object.Call("loadUpdate", args)
}

// loadRender is called during the Loader process. This only happens if you've set one or more assets to load in the preload method.
// The difference between loadRender and render is that any objects you render in this method you must be sure their assets exist.
func (self *State) LoadRender() {
    self.Object.Call("loadRender")
}

// loadRender is called during the Loader process. This only happens if you've set one or more assets to load in the preload method.
// The difference between loadRender and render is that any objects you render in this method you must be sure their assets exist.
func (self *State) LoadRenderI(args ...interface{}) {
    self.Object.Call("loadRender", args)
}

// create is called once preload has completed, this includes the loading of any assets from the Loader.
// If you don't have a preload method then create is the first method called in your State.
func (self *State) Create() {
    self.Object.Call("create")
}

// create is called once preload has completed, this includes the loading of any assets from the Loader.
// If you don't have a preload method then create is the first method called in your State.
func (self *State) CreateI(args ...interface{}) {
    self.Object.Call("create", args)
}

// The update method is left empty for your own use.
// It is called during the core game loop AFTER debug, physics, plugins and the Stage have had their preUpdate methods called.
// It is called BEFORE Stage, Tweens, Sounds, Input, Physics, Particles and Plugins have had their postUpdate methods called.
func (self *State) Update() {
    self.Object.Call("update")
}

// The update method is left empty for your own use.
// It is called during the core game loop AFTER debug, physics, plugins and the Stage have had their preUpdate methods called.
// It is called BEFORE Stage, Tweens, Sounds, Input, Physics, Particles and Plugins have had their postUpdate methods called.
func (self *State) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// The preRender method is called after all Game Objects have been updated, but before any rendering takes place.
func (self *State) PreRender() {
    self.Object.Call("preRender")
}

// The preRender method is called after all Game Objects have been updated, but before any rendering takes place.
func (self *State) PreRenderI(args ...interface{}) {
    self.Object.Call("preRender", args)
}

// Nearly all display objects in Phaser render automatically, you don't need to tell them to render.
// However the render method is called AFTER the game renderer and plugins have rendered, so you're able to do any
// final post-processing style effects here. Note that this happens before plugins postRender takes place.
func (self *State) Render() {
    self.Object.Call("render")
}

// Nearly all display objects in Phaser render automatically, you don't need to tell them to render.
// However the render method is called AFTER the game renderer and plugins have rendered, so you're able to do any
// final post-processing style effects here. Note that this happens before plugins postRender takes place.
func (self *State) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// If your game is set to Scalemode RESIZE then each time the browser resizes it will call this function, passing in the new width and height.
func (self *State) Resize() {
    self.Object.Call("resize")
}

// If your game is set to Scalemode RESIZE then each time the browser resizes it will call this function, passing in the new width and height.
func (self *State) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// This method will be called if the core game loop is paused.
func (self *State) Paused() {
    self.Object.Call("paused")
}

// This method will be called if the core game loop is paused.
func (self *State) PausedI(args ...interface{}) {
    self.Object.Call("paused", args)
}

// This method will be called when the core game loop resumes from a paused state.
func (self *State) Resumed() {
    self.Object.Call("resumed")
}

// This method will be called when the core game loop resumes from a paused state.
func (self *State) ResumedI(args ...interface{}) {
    self.Object.Call("resumed", args)
}

// pauseUpdate is called while the game is paused instead of preUpdate, update and postUpdate.
func (self *State) PauseUpdate() {
    self.Object.Call("pauseUpdate")
}

// pauseUpdate is called while the game is paused instead of preUpdate, update and postUpdate.
func (self *State) PauseUpdateI(args ...interface{}) {
    self.Object.Call("pauseUpdate", args)
}

// This method will be called when the State is shutdown (i.e. you switch to another state from this one).
func (self *State) Shutdown() {
    self.Object.Call("shutdown")
}

// This method will be called when the State is shutdown (i.e. you switch to another state from this one).
func (self *State) ShutdownI(args ...interface{}) {
    self.Object.Call("shutdown", args)
}
