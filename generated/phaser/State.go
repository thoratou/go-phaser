// Package phaser Automatic generation for Phaser.State
// generated file State.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// State This is a base State class which can be extended if you are creating your own game.
// It provides quick access to common functions such as the camera, cache, input, match, sound and more.
type State struct {
    *js.Object
}

// NewState This is a base State class which can be extended if you are creating your own game.
// It provides quick access to common functions such as the camera, cache, input, match, sound and more.
func NewState() *State {
    return &State{js.Global.Get("Phaser").Get("State").New()}
}
// NewStateI This is a base State class which can be extended if you are creating your own game.
// It provides quick access to common functions such as the camera, cache, input, match, sound and more.
func NewStateI(args ...interface{}) *State {
    return &State{js.Global.Get("Phaser").Get("State").New(args)}
}



// Game This is a reference to the currently running Game.
func (self *State) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA This is a reference to the currently running Game.
func (self *State) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Key The string based identifier given to the State when added into the State Manager.
func (self *State) Key() string{
    return self.Object.Get("key").String()
}

// SetKeyA The string based identifier given to the State when added into the State Manager.
func (self *State) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// Add A reference to the GameObjectFactory which can be used to add new objects to the World.
func (self *State) Add() *GameObjectFactory{
    return &GameObjectFactory{self.Object.Get("add")}
}

// SetAddA A reference to the GameObjectFactory which can be used to add new objects to the World.
func (self *State) SetAddA(member *GameObjectFactory) {
    self.Object.Set("add", member)
}

// Make A reference to the GameObjectCreator which can be used to make new objects.
func (self *State) Make() *GameObjectCreator{
    return &GameObjectCreator{self.Object.Get("make")}
}

// SetMakeA A reference to the GameObjectCreator which can be used to make new objects.
func (self *State) SetMakeA(member *GameObjectCreator) {
    self.Object.Set("make", member)
}

// Camera A handy reference to World.camera.
func (self *State) Camera() *Camera{
    return &Camera{self.Object.Get("camera")}
}

// SetCameraA A handy reference to World.camera.
func (self *State) SetCameraA(member *Camera) {
    self.Object.Set("camera", member)
}

// Cache A reference to the game cache which contains any loaded or generated assets, such as images, sound and more.
func (self *State) Cache() *Cache{
    return &Cache{self.Object.Get("cache")}
}

// SetCacheA A reference to the game cache which contains any loaded or generated assets, such as images, sound and more.
func (self *State) SetCacheA(member *Cache) {
    self.Object.Set("cache", member)
}

// Input A reference to the Input Manager.
func (self *State) Input() *Input{
    return &Input{self.Object.Get("input")}
}

// SetInputA A reference to the Input Manager.
func (self *State) SetInputA(member *Input) {
    self.Object.Set("input", member)
}

// Load A reference to the Loader, which you mostly use in the preload method of your state to load external assets.
func (self *State) Load() *Loader{
    return &Loader{self.Object.Get("load")}
}

// SetLoadA A reference to the Loader, which you mostly use in the preload method of your state to load external assets.
func (self *State) SetLoadA(member *Loader) {
    self.Object.Set("load", member)
}

// Math A reference to Math class with lots of helpful functions.
func (self *State) Math() *Math{
    return &Math{self.Object.Get("math")}
}

// SetMathA A reference to Math class with lots of helpful functions.
func (self *State) SetMathA(member *Math) {
    self.Object.Set("math", member)
}

// Sound A reference to the Sound Manager which can create, play and stop sounds, as well as adjust global volume.
func (self *State) Sound() *SoundManager{
    return &SoundManager{self.Object.Get("sound")}
}

// SetSoundA A reference to the Sound Manager which can create, play and stop sounds, as well as adjust global volume.
func (self *State) SetSoundA(member *SoundManager) {
    self.Object.Set("sound", member)
}

// Scale A reference to the Scale Manager which controls the way the game scales on different displays.
func (self *State) Scale() *ScaleManager{
    return &ScaleManager{self.Object.Get("scale")}
}

// SetScaleA A reference to the Scale Manager which controls the way the game scales on different displays.
func (self *State) SetScaleA(member *ScaleManager) {
    self.Object.Set("scale", member)
}

// Stage A reference to the Stage.
func (self *State) Stage() *Stage{
    return &Stage{self.Object.Get("stage")}
}

// SetStageA A reference to the Stage.
func (self *State) SetStageA(member *Stage) {
    self.Object.Set("stage", member)
}

// State A reference to the State Manager, which controls state changes.
func (self *State) State() interface{}{
    return self.Object.Get("state")
}

// SetStateA A reference to the State Manager, which controls state changes.
func (self *State) SetStateA(member interface{}) {
    self.Object.Set("state", member)
}

// Time A reference to the game clock and timed events system.
func (self *State) Time() *Time{
    return &Time{self.Object.Get("time")}
}

// SetTimeA A reference to the game clock and timed events system.
func (self *State) SetTimeA(member *Time) {
    self.Object.Set("time", member)
}

// Tweens A reference to the tween manager.
func (self *State) Tweens() *TweenManager{
    return &TweenManager{self.Object.Get("tweens")}
}

// SetTweensA A reference to the tween manager.
func (self *State) SetTweensA(member *TweenManager) {
    self.Object.Set("tweens", member)
}

// World A reference to the game world. All objects live in the Game World and its size is not bound by the display resolution.
func (self *State) World() *World{
    return &World{self.Object.Get("world")}
}

// SetWorldA A reference to the game world. All objects live in the Game World and its size is not bound by the display resolution.
func (self *State) SetWorldA(member *World) {
    self.Object.Set("world", member)
}

// Particles The Particle Manager. It is called during the core gameloop and updates any Particle Emitters it has created.
func (self *State) Particles() *Particles{
    return &Particles{self.Object.Get("particles")}
}

// SetParticlesA The Particle Manager. It is called during the core gameloop and updates any Particle Emitters it has created.
func (self *State) SetParticlesA(member *Particles) {
    self.Object.Set("particles", member)
}

// Physics A reference to the physics manager which looks after the different physics systems available within Phaser.
func (self *State) Physics() *Physics{
    return &Physics{self.Object.Get("physics")}
}

// SetPhysicsA A reference to the physics manager which looks after the different physics systems available within Phaser.
func (self *State) SetPhysicsA(member *Physics) {
    self.Object.Set("physics", member)
}

// Rnd A reference to the seeded and repeatable random data generator.
func (self *State) Rnd() *RandomDataGenerator{
    return &RandomDataGenerator{self.Object.Get("rnd")}
}

// SetRndA A reference to the seeded and repeatable random data generator.
func (self *State) SetRndA(member *RandomDataGenerator) {
    self.Object.Set("rnd", member)
}


// Init init is the very first function called when your State starts up. It's called before preload, create or anything else.
// If you need to route the game away to another State you could do so here, or if you need to prepare a set of variables
// or objects before the preloading starts.
func (self *State) Init() {
    self.Object.Call("init")
}

// InitI init is the very first function called when your State starts up. It's called before preload, create or anything else.
// If you need to route the game away to another State you could do so here, or if you need to prepare a set of variables
// or objects before the preloading starts.
func (self *State) InitI(args ...interface{}) {
    self.Object.Call("init", args)
}

// Preload preload is called first. Normally you'd use this to load your game assets (or those needed for the current State)
// You shouldn't create any objects in this method that require assets that you're also loading in this method, as
// they won't yet be available.
func (self *State) Preload() {
    self.Object.Call("preload")
}

// PreloadI preload is called first. Normally you'd use this to load your game assets (or those needed for the current State)
// You shouldn't create any objects in this method that require assets that you're also loading in this method, as
// they won't yet be available.
func (self *State) PreloadI(args ...interface{}) {
    self.Object.Call("preload", args)
}

// LoadUpdate loadUpdate is called during the Loader process. This only happens if you've set one or more assets to load in the preload method.
func (self *State) LoadUpdate() {
    self.Object.Call("loadUpdate")
}

// LoadUpdateI loadUpdate is called during the Loader process. This only happens if you've set one or more assets to load in the preload method.
func (self *State) LoadUpdateI(args ...interface{}) {
    self.Object.Call("loadUpdate", args)
}

// LoadRender loadRender is called during the Loader process. This only happens if you've set one or more assets to load in the preload method.
// The difference between loadRender and render is that any objects you render in this method you must be sure their assets exist.
func (self *State) LoadRender() {
    self.Object.Call("loadRender")
}

// LoadRenderI loadRender is called during the Loader process. This only happens if you've set one or more assets to load in the preload method.
// The difference between loadRender and render is that any objects you render in this method you must be sure their assets exist.
func (self *State) LoadRenderI(args ...interface{}) {
    self.Object.Call("loadRender", args)
}

// Create create is called once preload has completed, this includes the loading of any assets from the Loader.
// If you don't have a preload method then create is the first method called in your State.
func (self *State) Create() {
    self.Object.Call("create")
}

// CreateI create is called once preload has completed, this includes the loading of any assets from the Loader.
// If you don't have a preload method then create is the first method called in your State.
func (self *State) CreateI(args ...interface{}) {
    self.Object.Call("create", args)
}

// Update The update method is left empty for your own use.
// It is called during the core game loop AFTER debug, physics, plugins and the Stage have had their preUpdate methods called.
// It is called BEFORE Stage, Tweens, Sounds, Input, Physics, Particles and Plugins have had their postUpdate methods called.
func (self *State) Update() {
    self.Object.Call("update")
}

// UpdateI The update method is left empty for your own use.
// It is called during the core game loop AFTER debug, physics, plugins and the Stage have had their preUpdate methods called.
// It is called BEFORE Stage, Tweens, Sounds, Input, Physics, Particles and Plugins have had their postUpdate methods called.
func (self *State) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// PreRender The preRender method is called after all Game Objects have been updated, but before any rendering takes place.
func (self *State) PreRender() {
    self.Object.Call("preRender")
}

// PreRenderI The preRender method is called after all Game Objects have been updated, but before any rendering takes place.
func (self *State) PreRenderI(args ...interface{}) {
    self.Object.Call("preRender", args)
}

// Render Nearly all display objects in Phaser render automatically, you don't need to tell them to render.
// However the render method is called AFTER the game renderer and plugins have rendered, so you're able to do any
// final post-processing style effects here. Note that this happens before plugins postRender takes place.
func (self *State) Render() {
    self.Object.Call("render")
}

// RenderI Nearly all display objects in Phaser render automatically, you don't need to tell them to render.
// However the render method is called AFTER the game renderer and plugins have rendered, so you're able to do any
// final post-processing style effects here. Note that this happens before plugins postRender takes place.
func (self *State) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// Resize If your game is set to Scalemode RESIZE then each time the browser resizes it will call this function, passing in the new width and height.
func (self *State) Resize() {
    self.Object.Call("resize")
}

// ResizeI If your game is set to Scalemode RESIZE then each time the browser resizes it will call this function, passing in the new width and height.
func (self *State) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Paused This method will be called if the core game loop is paused.
func (self *State) Paused() {
    self.Object.Call("paused")
}

// PausedI This method will be called if the core game loop is paused.
func (self *State) PausedI(args ...interface{}) {
    self.Object.Call("paused", args)
}

// Resumed This method will be called when the core game loop resumes from a paused state.
func (self *State) Resumed() {
    self.Object.Call("resumed")
}

// ResumedI This method will be called when the core game loop resumes from a paused state.
func (self *State) ResumedI(args ...interface{}) {
    self.Object.Call("resumed", args)
}

// PauseUpdate pauseUpdate is called while the game is paused instead of preUpdate, update and postUpdate.
func (self *State) PauseUpdate() {
    self.Object.Call("pauseUpdate")
}

// PauseUpdateI pauseUpdate is called while the game is paused instead of preUpdate, update and postUpdate.
func (self *State) PauseUpdateI(args ...interface{}) {
    self.Object.Call("pauseUpdate", args)
}

// Shutdown This method will be called when the State is shutdown (i.e. you switch to another state from this one).
func (self *State) Shutdown() {
    self.Object.Call("shutdown")
}

// ShutdownI This method will be called when the State is shutdown (i.e. you switch to another state from this one).
func (self *State) ShutdownI(args ...interface{}) {
    self.Object.Call("shutdown", args)
}

