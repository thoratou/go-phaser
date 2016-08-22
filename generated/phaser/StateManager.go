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
func (self *StateManager) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *StateManager) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The object containing Phaser.States.
func (self *StateManager) GetStatesA() interface{}{
    return self.Object.Get("states")
}

// The object containing Phaser.States.
func (self *StateManager) SetStatesA(member interface{}) {
    self.Object.Set("states", member)
}

// The current active State object.
func (self *StateManager) GetCurrentA() string{
    return self.Object.Get("current").String()
}

// The current active State object.
func (self *StateManager) SetCurrentA(member string) {
    self.Object.Set("current", member)
}

// onStateChange is a Phaser.Signal that is dispatched whenever the game changes state.
// 
// It is dispatched only when the new state is started, which isn't usually at the same time as StateManager.start
// is called because state swapping is done in sync with the game loop. It is dispatched *before* any of the new states
// methods (such as preload and create) are called, and *after* the previous states shutdown method has been run.
// 
// The callback you specify is sent two parameters: the string based key of the new state, 
// and the second parameter is the string based key of the old / previous state.
func (self *StateManager) GetOnStateChangeA() *Signal{
    return &Signal{self.Object.Get("onStateChange")}
}

// onStateChange is a Phaser.Signal that is dispatched whenever the game changes state.
// 
// It is dispatched only when the new state is started, which isn't usually at the same time as StateManager.start
// is called because state swapping is done in sync with the game loop. It is dispatched *before* any of the new states
// methods (such as preload and create) are called, and *after* the previous states shutdown method has been run.
// 
// The callback you specify is sent two parameters: the string based key of the new state, 
// and the second parameter is the string based key of the old / previous state.
func (self *StateManager) SetOnStateChangeA(member *Signal) {
    self.Object.Set("onStateChange", member)
}

// This is called when the state is set as the active state.
func (self *StateManager) SetOnInitCallbackA(member func(...interface{})) {
    self.Object.Set("onInitCallback", member)
}

// This is called when the state starts to load assets.
func (self *StateManager) SetOnPreloadCallbackA(member func(...interface{})) {
    self.Object.Set("onPreloadCallback", member)
}

// This is called when the state preload has finished and creation begins.
func (self *StateManager) SetOnCreateCallbackA(member func(...interface{})) {
    self.Object.Set("onCreateCallback", member)
}

// This is called when the state is updated, every game loop. It doesn't happen during preload (@see onLoadUpdateCallback).
func (self *StateManager) SetOnUpdateCallbackA(member func(...interface{})) {
    self.Object.Set("onUpdateCallback", member)
}

// This is called post-render. It doesn't happen during preload (see onLoadRenderCallback).
func (self *StateManager) SetOnRenderCallbackA(member func(...interface{})) {
    self.Object.Set("onRenderCallback", member)
}

// This is called if ScaleManager.scalemode is RESIZE and a resize event occurs. It's passed the new width and height.
func (self *StateManager) SetOnResizeCallbackA(member func(...interface{})) {
    self.Object.Set("onResizeCallback", member)
}

// This is called before the state is rendered and before the stage is cleared but after all game objects have had their final properties adjusted.
func (self *StateManager) SetOnPreRenderCallbackA(member func(...interface{})) {
    self.Object.Set("onPreRenderCallback", member)
}

// This is called when the State is updated during the preload phase.
func (self *StateManager) SetOnLoadUpdateCallbackA(member func(...interface{})) {
    self.Object.Set("onLoadUpdateCallback", member)
}

// This is called when the State is rendered during the preload phase.
func (self *StateManager) SetOnLoadRenderCallbackA(member func(...interface{})) {
    self.Object.Set("onLoadRenderCallback", member)
}

// This is called when the game is paused.
func (self *StateManager) SetOnPausedCallbackA(member func(...interface{})) {
    self.Object.Set("onPausedCallback", member)
}

// This is called when the game is resumed from a paused state.
func (self *StateManager) SetOnResumedCallbackA(member func(...interface{})) {
    self.Object.Set("onResumedCallback", member)
}

// This is called every frame while the game is paused.
func (self *StateManager) SetOnPauseUpdateCallbackA(member func(...interface{})) {
    self.Object.Set("onPauseUpdateCallback", member)
}

// This is called when the state is shut down (i.e. swapped to another state).
func (self *StateManager) SetOnShutDownCallbackA(member func(...interface{})) {
    self.Object.Set("onShutDownCallback", member)
}

// True if the current state has had its `create` method run (if it has one, if not this is true by default).
func (self *StateManager) GetCreatedA() bool{
    return self.Object.Get("created").Bool()
}

// True if the current state has had its `create` method run (if it has one, if not this is true by default).
func (self *StateManager) SetCreatedA(member bool) {
    self.Object.Set("created", member)
}



// The Boot handler is called by Phaser.Game when it first starts up.
func (self *StateManager) Boot() {
    self.Object.Call("boot")
}

// The Boot handler is called by Phaser.Game when it first starts up.
func (self *StateManager) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// Adds a new State into the StateManager. You must give each State a unique key by which you'll identify it.
// The State can be either a Phaser.State object (or an object that extends it), a plain JavaScript object or a function.
// If a function is given a new state object will be created by calling it.
func (self *StateManager) Add(key string, state interface{}, autoStart bool) {
    self.Object.Call("add", key, state, autoStart)
}

// Adds a new State into the StateManager. You must give each State a unique key by which you'll identify it.
// The State can be either a Phaser.State object (or an object that extends it), a plain JavaScript object or a function.
// If a function is given a new state object will be created by calling it.
func (self *StateManager) AddI(args ...interface{}) {
    self.Object.Call("add", args)
}

// Delete the given state.
func (self *StateManager) Remove(key string) {
    self.Object.Call("remove", key)
}

// Delete the given state.
func (self *StateManager) RemoveI(args ...interface{}) {
    self.Object.Call("remove", args)
}

// Start the given State. If a State is already running then State.shutDown will be called (if it exists) before switching to the new State.
func (self *StateManager) Start(key string, clearWorld bool, clearCache bool, parameter interface{}) {
    self.Object.Call("start", key, clearWorld, clearCache, parameter)
}

// Start the given State. If a State is already running then State.shutDown will be called (if it exists) before switching to the new State.
func (self *StateManager) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Restarts the current State. State.shutDown will be called (if it exists) before the State is restarted.
func (self *StateManager) Restart(clearWorld bool, clearCache bool, parameter interface{}) {
    self.Object.Call("restart", clearWorld, clearCache, parameter)
}

// Restarts the current State. State.shutDown will be called (if it exists) before the State is restarted.
func (self *StateManager) RestartI(args ...interface{}) {
    self.Object.Call("restart", args)
}

// Used by onInit and onShutdown when those functions don't exist on the state
func (self *StateManager) Dummy() {
    self.Object.Call("dummy")
}

// Used by onInit and onShutdown when those functions don't exist on the state
func (self *StateManager) DummyI(args ...interface{}) {
    self.Object.Call("dummy", args)
}

// preUpdate is called right at the start of the game loop. It is responsible for changing to a new state that was requested previously.
func (self *StateManager) PreUpdate() {
    self.Object.Call("preUpdate")
}

// preUpdate is called right at the start of the game loop. It is responsible for changing to a new state that was requested previously.
func (self *StateManager) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// This method clears the current State, calling its shutdown callback. The process also removes any active tweens,
// resets the camera, resets input, clears physics, removes timers and if set clears the world and cache too.
func (self *StateManager) ClearCurrentState() {
    self.Object.Call("clearCurrentState")
}

// This method clears the current State, calling its shutdown callback. The process also removes any active tweens,
// resets the camera, resets input, clears physics, removes timers and if set clears the world and cache too.
func (self *StateManager) ClearCurrentStateI(args ...interface{}) {
    self.Object.Call("clearCurrentState", args)
}

// Checks if a given phaser state is valid. A State is considered valid if it has at least one of the core functions: preload, create, update or render.
func (self *StateManager) CheckState(key string) bool{
    return self.Object.Call("checkState", key).Bool()
}

// Checks if a given phaser state is valid. A State is considered valid if it has at least one of the core functions: preload, create, update or render.
func (self *StateManager) CheckStateI(args ...interface{}) bool{
    return self.Object.Call("checkState", args).Bool()
}

// Links game properties to the State given by the key.
func (self *StateManager) Link(key string) {
    self.Object.Call("link", key)
}

// Links game properties to the State given by the key.
func (self *StateManager) LinkI(args ...interface{}) {
    self.Object.Call("link", args)
}

// Nulls all State level Phaser properties, including a reference to Game.
func (self *StateManager) Unlink(key string) {
    self.Object.Call("unlink", key)
}

// Nulls all State level Phaser properties, including a reference to Game.
func (self *StateManager) UnlinkI(args ...interface{}) {
    self.Object.Call("unlink", args)
}

// Sets the current State. Should not be called directly (use StateManager.start)
func (self *StateManager) SetCurrentState(key string) {
    self.Object.Call("setCurrentState", key)
}

// Sets the current State. Should not be called directly (use StateManager.start)
func (self *StateManager) SetCurrentStateI(args ...interface{}) {
    self.Object.Call("setCurrentState", args)
}

// Gets the current State.
func (self *StateManager) GetCurrentState() *State{
    return &State{self.Object.Call("getCurrentState")}
}

// Gets the current State.
func (self *StateManager) GetCurrentStateI(args ...interface{}) *State{
    return &State{self.Object.Call("getCurrentState", args)}
}

// 
func (self *StateManager) LoadComplete() {
    self.Object.Call("loadComplete")
}

// 
func (self *StateManager) LoadCompleteI(args ...interface{}) {
    self.Object.Call("loadComplete", args)
}

// 
func (self *StateManager) Pause() {
    self.Object.Call("pause")
}

// 
func (self *StateManager) PauseI(args ...interface{}) {
    self.Object.Call("pause", args)
}

// 
func (self *StateManager) Resume() {
    self.Object.Call("resume")
}

// 
func (self *StateManager) ResumeI(args ...interface{}) {
    self.Object.Call("resume", args)
}

// 
func (self *StateManager) Update() {
    self.Object.Call("update")
}

// 
func (self *StateManager) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// 
func (self *StateManager) PauseUpdate() {
    self.Object.Call("pauseUpdate")
}

// 
func (self *StateManager) PauseUpdateI(args ...interface{}) {
    self.Object.Call("pauseUpdate", args)
}

// 
func (self *StateManager) PreRender(elapsedTime int) {
    self.Object.Call("preRender", elapsedTime)
}

// 
func (self *StateManager) PreRenderI(args ...interface{}) {
    self.Object.Call("preRender", args)
}

// 
func (self *StateManager) Resize() {
    self.Object.Call("resize")
}

// 
func (self *StateManager) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// 
func (self *StateManager) Render() {
    self.Object.Call("render")
}

// 
func (self *StateManager) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// Removes all StateManager callback references to the State object, nulls the game reference and clears the States object.
// You don't recover from this without rebuilding the Phaser instance again.
func (self *StateManager) Destroy() {
    self.Object.Call("destroy")
}

// Removes all StateManager callback references to the State object, nulls the game reference and clears the States object.
// You don't recover from this without rebuilding the Phaser instance again.
func (self *StateManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
