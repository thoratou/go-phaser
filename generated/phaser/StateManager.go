// Automatic generation for Phaser.StateManager
// generated file StateManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The State Manager is responsible for loading, setting up and switching game states.
type StateManager struct {
    *js.Object
}


// A reference to the currently running game.
func (self *StateManager) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *StateManager) SetGame(member Game) {
    self.Set("game", member)
}

// The object containing Phaser.States.
func (self *StateManager) GetStates() interface{}{
    return self.Get("states")
}

// The object containing Phaser.States.
func (self *StateManager) SetStates(member interface{}) {
    self.Set("states", member)
}

// The current active State object.
func (self *StateManager) GetCurrent() string{
    return self.Get("current").String()
}

// The current active State object.
func (self *StateManager) SetCurrent(member string) {
    self.Set("current", member)
}

// onStateChange is a Phaser.Signal that is dispatched whenever the game changes state.
// 
// It is dispatched only when the new state is started, which isn't usually at the same time as StateManager.start
// is called because state swapping is done in sync with the game loop. It is dispatched *before* any of the new states
// methods (such as preload and create) are called, and *after* the previous states shutdown method has been run.
// 
// The callback you specify is sent two parameters: the string based key of the new state, 
// and the second parameter is the string based key of the old / previous state.
func (self *StateManager) GetOnStateChange() Signal{
    return Signal{self.Get("onStateChange")}
}

// onStateChange is a Phaser.Signal that is dispatched whenever the game changes state.
// 
// It is dispatched only when the new state is started, which isn't usually at the same time as StateManager.start
// is called because state swapping is done in sync with the game loop. It is dispatched *before* any of the new states
// methods (such as preload and create) are called, and *after* the previous states shutdown method has been run.
// 
// The callback you specify is sent two parameters: the string based key of the new state, 
// and the second parameter is the string based key of the old / previous state.
func (self *StateManager) SetOnStateChange(member Signal) {
    self.Set("onStateChange", member)
}

// This is called when the state is set as the active state.
func (self *StateManager) SetOnInitCallback(member func(...interface{})) {
    self.Set("onInitCallback", member)
}

// This is called when the state starts to load assets.
func (self *StateManager) SetOnPreloadCallback(member func(...interface{})) {
    self.Set("onPreloadCallback", member)
}

// This is called when the state preload has finished and creation begins.
func (self *StateManager) SetOnCreateCallback(member func(...interface{})) {
    self.Set("onCreateCallback", member)
}

// This is called when the state is updated, every game loop. It doesn't happen during preload (@see onLoadUpdateCallback).
func (self *StateManager) SetOnUpdateCallback(member func(...interface{})) {
    self.Set("onUpdateCallback", member)
}

// This is called post-render. It doesn't happen during preload (see onLoadRenderCallback).
func (self *StateManager) SetOnRenderCallback(member func(...interface{})) {
    self.Set("onRenderCallback", member)
}

// This is called if ScaleManager.scalemode is RESIZE and a resize event occurs. It's passed the new width and height.
func (self *StateManager) SetOnResizeCallback(member func(...interface{})) {
    self.Set("onResizeCallback", member)
}

// This is called before the state is rendered and before the stage is cleared but after all game objects have had their final properties adjusted.
func (self *StateManager) SetOnPreRenderCallback(member func(...interface{})) {
    self.Set("onPreRenderCallback", member)
}

// This is called when the State is updated during the preload phase.
func (self *StateManager) SetOnLoadUpdateCallback(member func(...interface{})) {
    self.Set("onLoadUpdateCallback", member)
}

// This is called when the State is rendered during the preload phase.
func (self *StateManager) SetOnLoadRenderCallback(member func(...interface{})) {
    self.Set("onLoadRenderCallback", member)
}

// This is called when the game is paused.
func (self *StateManager) SetOnPausedCallback(member func(...interface{})) {
    self.Set("onPausedCallback", member)
}

// This is called when the game is resumed from a paused state.
func (self *StateManager) SetOnResumedCallback(member func(...interface{})) {
    self.Set("onResumedCallback", member)
}

// This is called every frame while the game is paused.
func (self *StateManager) SetOnPauseUpdateCallback(member func(...interface{})) {
    self.Set("onPauseUpdateCallback", member)
}

// This is called when the state is shut down (i.e. swapped to another state).
func (self *StateManager) SetOnShutDownCallback(member func(...interface{})) {
    self.Set("onShutDownCallback", member)
}

// True if the current state has had its `create` method run (if it has one, if not this is true by default).
func (self *StateManager) GetCreated() bool{
    return self.Get("created").Bool()
}

// True if the current state has had its `create` method run (if it has one, if not this is true by default).
func (self *StateManager) SetCreated(member bool) {
    self.Set("created", member)
}



// The Boot handler is called by Phaser.Game when it first starts up.
func (self *StateManager) BootI(args ...interface{}) {
    self.Call("boot", args)
}

// Adds a new State into the StateManager. You must give each State a unique key by which you'll identify it.
// The State can be either a Phaser.State object (or an object that extends it), a plain JavaScript object or a function.
// If a function is given a new state object will be created by calling it.
func (self *StateManager) AddI(args ...interface{}) {
    self.Call("add", args)
}

// Delete the given state.
func (self *StateManager) RemoveI(args ...interface{}) {
    self.Call("remove", args)
}

// Start the given State. If a State is already running then State.shutDown will be called (if it exists) before switching to the new State.
func (self *StateManager) StartI(args ...interface{}) {
    self.Call("start", args)
}

// Restarts the current State. State.shutDown will be called (if it exists) before the State is restarted.
func (self *StateManager) RestartI(args ...interface{}) {
    self.Call("restart", args)
}

// Used by onInit and onShutdown when those functions don't exist on the state
func (self *StateManager) DummyI(args ...interface{}) {
    self.Call("dummy", args)
}

// preUpdate is called right at the start of the game loop. It is responsible for changing to a new state that was requested previously.
func (self *StateManager) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// This method clears the current State, calling its shutdown callback. The process also removes any active tweens,
// resets the camera, resets input, clears physics, removes timers and if set clears the world and cache too.
func (self *StateManager) ClearCurrentStateI(args ...interface{}) {
    self.Call("clearCurrentState", args)
}

// Checks if a given phaser state is valid. A State is considered valid if it has at least one of the core functions: preload, create, update or render.
func (self *StateManager) CheckStateI(args ...interface{}) bool{
    return self.Call("checkState", args).Bool()
}

// Links game properties to the State given by the key.
func (self *StateManager) LinkI(args ...interface{}) {
    self.Call("link", args)
}

// Nulls all State level Phaser properties, including a reference to Game.
func (self *StateManager) UnlinkI(args ...interface{}) {
    self.Call("unlink", args)
}

// Sets the current State. Should not be called directly (use StateManager.start)
func (self *StateManager) SetCurrentStateI(args ...interface{}) {
    self.Call("setCurrentState", args)
}

// Gets the current State.
func (self *StateManager) GetCurrentStateI(args ...interface{}) State{
    return State{self.Call("getCurrentState", args)}
}

// 
func (self *StateManager) LoadCompleteI(args ...interface{}) {
    self.Call("loadComplete", args)
}

// 
func (self *StateManager) PauseI(args ...interface{}) {
    self.Call("pause", args)
}

// 
func (self *StateManager) ResumeI(args ...interface{}) {
    self.Call("resume", args)
}

// 
func (self *StateManager) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// 
func (self *StateManager) PauseUpdateI(args ...interface{}) {
    self.Call("pauseUpdate", args)
}

// 
func (self *StateManager) PreRenderI(args ...interface{}) {
    self.Call("preRender", args)
}

// 
func (self *StateManager) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// 
func (self *StateManager) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// Removes all StateManager callback references to the State object, nulls the game reference and clears the States object.
// You don't recover from this without rebuilding the Phaser instance again.
func (self *StateManager) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
