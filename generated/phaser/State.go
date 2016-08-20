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


// This is a reference to the currently running Game.
func (self *State) GetGame() Game{
    return Game{self.Get("game")}
}

// This is a reference to the currently running Game.
func (self *State) SetGame(member Game) {
    self.Set("game", member)
}

// The string based identifier given to the State when added into the State Manager.
func (self *State) GetKey() string{
    return self.Get("key").String()
}

// The string based identifier given to the State when added into the State Manager.
func (self *State) SetKey(member string) {
    self.Set("key", member)
}

// A reference to the GameObjectFactory which can be used to add new objects to the World.
func (self *State) GetAdd() GameObjectFactory{
    return GameObjectFactory{self.Get("add")}
}

// A reference to the GameObjectFactory which can be used to add new objects to the World.
func (self *State) SetAdd(member GameObjectFactory) {
    self.Set("add", member)
}

// A reference to the GameObjectCreator which can be used to make new objects.
func (self *State) GetMake() GameObjectCreator{
    return GameObjectCreator{self.Get("make")}
}

// A reference to the GameObjectCreator which can be used to make new objects.
func (self *State) SetMake(member GameObjectCreator) {
    self.Set("make", member)
}

// A handy reference to World.camera.
func (self *State) GetCamera() Camera{
    return Camera{self.Get("camera")}
}

// A handy reference to World.camera.
func (self *State) SetCamera(member Camera) {
    self.Set("camera", member)
}

// A reference to the game cache which contains any loaded or generated assets, such as images, sound and more.
func (self *State) GetCache() Cache{
    return Cache{self.Get("cache")}
}

// A reference to the game cache which contains any loaded or generated assets, such as images, sound and more.
func (self *State) SetCache(member Cache) {
    self.Set("cache", member)
}

// A reference to the Input Manager.
func (self *State) GetInput() Input{
    return Input{self.Get("input")}
}

// A reference to the Input Manager.
func (self *State) SetInput(member Input) {
    self.Set("input", member)
}

// A reference to the Loader, which you mostly use in the preload method of your state to load external assets.
func (self *State) GetLoad() Loader{
    return Loader{self.Get("load")}
}

// A reference to the Loader, which you mostly use in the preload method of your state to load external assets.
func (self *State) SetLoad(member Loader) {
    self.Set("load", member)
}

// A reference to Math class with lots of helpful functions.
func (self *State) GetMath() Math{
    return Math{self.Get("math")}
}

// A reference to Math class with lots of helpful functions.
func (self *State) SetMath(member Math) {
    self.Set("math", member)
}

// A reference to the Sound Manager which can create, play and stop sounds, as well as adjust global volume.
func (self *State) GetSound() SoundManager{
    return SoundManager{self.Get("sound")}
}

// A reference to the Sound Manager which can create, play and stop sounds, as well as adjust global volume.
func (self *State) SetSound(member SoundManager) {
    self.Set("sound", member)
}

// A reference to the Scale Manager which controls the way the game scales on different displays.
func (self *State) GetScale() ScaleManager{
    return ScaleManager{self.Get("scale")}
}

// A reference to the Scale Manager which controls the way the game scales on different displays.
func (self *State) SetScale(member ScaleManager) {
    self.Set("scale", member)
}

// A reference to the Stage.
func (self *State) GetStage() Stage{
    return Stage{self.Get("stage")}
}

// A reference to the Stage.
func (self *State) SetStage(member Stage) {
    self.Set("stage", member)
}

// A reference to the State Manager, which controls state changes.
func (self *State) GetState() interface{}{
    return self.Get("state")
}

// A reference to the State Manager, which controls state changes.
func (self *State) SetState(member interface{}) {
    self.Set("state", member)
}

// A reference to the game clock and timed events system.
func (self *State) GetTime() Time{
    return Time{self.Get("time")}
}

// A reference to the game clock and timed events system.
func (self *State) SetTime(member Time) {
    self.Set("time", member)
}

// A reference to the tween manager.
func (self *State) GetTweens() TweenManager{
    return TweenManager{self.Get("tweens")}
}

// A reference to the tween manager.
func (self *State) SetTweens(member TweenManager) {
    self.Set("tweens", member)
}

// A reference to the game world. All objects live in the Game World and its size is not bound by the display resolution.
func (self *State) GetWorld() World{
    return World{self.Get("world")}
}

// A reference to the game world. All objects live in the Game World and its size is not bound by the display resolution.
func (self *State) SetWorld(member World) {
    self.Set("world", member)
}

// The Particle Manager. It is called during the core gameloop and updates any Particle Emitters it has created.
func (self *State) GetParticles() Particles{
    return Particles{self.Get("particles")}
}

// The Particle Manager. It is called during the core gameloop and updates any Particle Emitters it has created.
func (self *State) SetParticles(member Particles) {
    self.Set("particles", member)
}

// A reference to the physics manager which looks after the different physics systems available within Phaser.
func (self *State) GetPhysics() Physics{
    return Physics{self.Get("physics")}
}

// A reference to the physics manager which looks after the different physics systems available within Phaser.
func (self *State) SetPhysics(member Physics) {
    self.Set("physics", member)
}

// A reference to the seeded and repeatable random data generator.
func (self *State) GetRnd() RandomDataGenerator{
    return RandomDataGenerator{self.Get("rnd")}
}

// A reference to the seeded and repeatable random data generator.
func (self *State) SetRnd(member RandomDataGenerator) {
    self.Set("rnd", member)
}



// init is the very first function called when your State starts up. It's called before preload, create or anything else.
// If you need to route the game away to another State you could do so here, or if you need to prepare a set of variables
// or objects before the preloading starts.
func (self *State) InitI(args ...interface{}) {
    self.Call("init", args)
}

// preload is called first. Normally you'd use this to load your game assets (or those needed for the current State)
// You shouldn't create any objects in this method that require assets that you're also loading in this method, as
// they won't yet be available.
func (self *State) PreloadI(args ...interface{}) {
    self.Call("preload", args)
}

// loadUpdate is called during the Loader process. This only happens if you've set one or more assets to load in the preload method.
func (self *State) LoadUpdateI(args ...interface{}) {
    self.Call("loadUpdate", args)
}

// loadRender is called during the Loader process. This only happens if you've set one or more assets to load in the preload method.
// The difference between loadRender and render is that any objects you render in this method you must be sure their assets exist.
func (self *State) LoadRenderI(args ...interface{}) {
    self.Call("loadRender", args)
}

// create is called once preload has completed, this includes the loading of any assets from the Loader.
// If you don't have a preload method then create is the first method called in your State.
func (self *State) CreateI(args ...interface{}) {
    self.Call("create", args)
}

// The update method is left empty for your own use.
// It is called during the core game loop AFTER debug, physics, plugins and the Stage have had their preUpdate methods called.
// It is called BEFORE Stage, Tweens, Sounds, Input, Physics, Particles and Plugins have had their postUpdate methods called.
func (self *State) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// The preRender method is called after all Game Objects have been updated, but before any rendering takes place.
func (self *State) PreRenderI(args ...interface{}) {
    self.Call("preRender", args)
}

// Nearly all display objects in Phaser render automatically, you don't need to tell them to render.
// However the render method is called AFTER the game renderer and plugins have rendered, so you're able to do any
// final post-processing style effects here. Note that this happens before plugins postRender takes place.
func (self *State) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// If your game is set to Scalemode RESIZE then each time the browser resizes it will call this function, passing in the new width and height.
func (self *State) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// This method will be called if the core game loop is paused.
func (self *State) PausedI(args ...interface{}) {
    self.Call("paused", args)
}

// This method will be called when the core game loop resumes from a paused state.
func (self *State) ResumedI(args ...interface{}) {
    self.Call("resumed", args)
}

// pauseUpdate is called while the game is paused instead of preUpdate, update and postUpdate.
func (self *State) PauseUpdateI(args ...interface{}) {
    self.Call("pauseUpdate", args)
}

// This method will be called when the State is shutdown (i.e. you switch to another state from this one).
func (self *State) ShutdownI(args ...interface{}) {
    self.Call("shutdown", args)
}
