// Package phaser Automatic generation for Phaser.StateManager
// generated file StateManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// StateManager The State Manager is responsible for loading, setting up and switching game states.
type StateManager struct {
    *js.Object
}

// NewStateManager The State Manager is responsible for loading, setting up and switching game states.
func NewStateManager(game *Game) *StateManager {
    return &StateManager{js.Global.Get("Phaser").Get("StateManager").New(game)}
}
// NewStateManager1O The State Manager is responsible for loading, setting up and switching game states.
func NewStateManager1O(game *Game, pendingState interface{}) *StateManager {
    return &StateManager{js.Global.Get("Phaser").Get("StateManager").New(game, pendingState)}
}
// NewStateManagerI The State Manager is responsible for loading, setting up and switching game states.
func NewStateManagerI(args ...interface{}) *StateManager {
    return &StateManager{js.Global.Get("Phaser").Get("StateManager").New(args)}
}



// Game A reference to the currently running game.
func (self *StateManager) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *StateManager) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// States The object containing Phaser.States.
func (self *StateManager) States() interface{}{
    return self.Object.Get("states")
}

// SetStatesA The object containing Phaser.States.
func (self *StateManager) SetStatesA(member interface{}) {
    self.Object.Set("states", member)
}

// Current The current active State object.
func (self *StateManager) Current() string{
    return self.Object.Get("current").String()
}

// SetCurrentA The current active State object.
func (self *StateManager) SetCurrentA(member string) {
    self.Object.Set("current", member)
}

// OnStateChange onStateChange is a Phaser.Signal that is dispatched whenever the game changes state.
// 
// It is dispatched only when the new state is started, which isn't usually at the same time as StateManager.start
// is called because state swapping is done in sync with the game loop. It is dispatched *before* any of the new states
// methods (such as preload and create) are called, and *after* the previous states shutdown method has been run.
// 
// The callback you specify is sent two parameters: the string based key of the new state, 
// and the second parameter is the string based key of the old / previous state.
func (self *StateManager) OnStateChange() *Signal{
    return &Signal{self.Object.Get("onStateChange")}
}

// SetOnStateChangeA onStateChange is a Phaser.Signal that is dispatched whenever the game changes state.
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

// OnInitCallback This is called when the state is set as the active state.
func (self *StateManager) OnInitCallback() interface{}{
    return self.Object.Get("onInitCallback")
}

// SetOnInitCallbackA This is called when the state is set as the active state.
func (self *StateManager) SetOnInitCallbackA(member interface{}) {
    self.Object.Set("onInitCallback", member)
}

// OnPreloadCallback This is called when the state starts to load assets.
func (self *StateManager) OnPreloadCallback() interface{}{
    return self.Object.Get("onPreloadCallback")
}

// SetOnPreloadCallbackA This is called when the state starts to load assets.
func (self *StateManager) SetOnPreloadCallbackA(member interface{}) {
    self.Object.Set("onPreloadCallback", member)
}

// OnCreateCallback This is called when the state preload has finished and creation begins.
func (self *StateManager) OnCreateCallback() interface{}{
    return self.Object.Get("onCreateCallback")
}

// SetOnCreateCallbackA This is called when the state preload has finished and creation begins.
func (self *StateManager) SetOnCreateCallbackA(member interface{}) {
    self.Object.Set("onCreateCallback", member)
}

// OnUpdateCallback This is called when the state is updated, every game loop. It doesn't happen during preload (@see onLoadUpdateCallback).
func (self *StateManager) OnUpdateCallback() interface{}{
    return self.Object.Get("onUpdateCallback")
}

// SetOnUpdateCallbackA This is called when the state is updated, every game loop. It doesn't happen during preload (@see onLoadUpdateCallback).
func (self *StateManager) SetOnUpdateCallbackA(member interface{}) {
    self.Object.Set("onUpdateCallback", member)
}

// OnRenderCallback This is called post-render. It doesn't happen during preload (see onLoadRenderCallback).
func (self *StateManager) OnRenderCallback() interface{}{
    return self.Object.Get("onRenderCallback")
}

// SetOnRenderCallbackA This is called post-render. It doesn't happen during preload (see onLoadRenderCallback).
func (self *StateManager) SetOnRenderCallbackA(member interface{}) {
    self.Object.Set("onRenderCallback", member)
}

// OnResizeCallback This is called if ScaleManager.scalemode is RESIZE and a resize event occurs. It's passed the new width and height.
func (self *StateManager) OnResizeCallback() interface{}{
    return self.Object.Get("onResizeCallback")
}

// SetOnResizeCallbackA This is called if ScaleManager.scalemode is RESIZE and a resize event occurs. It's passed the new width and height.
func (self *StateManager) SetOnResizeCallbackA(member interface{}) {
    self.Object.Set("onResizeCallback", member)
}

// OnPreRenderCallback This is called before the state is rendered and before the stage is cleared but after all game objects have had their final properties adjusted.
func (self *StateManager) OnPreRenderCallback() interface{}{
    return self.Object.Get("onPreRenderCallback")
}

// SetOnPreRenderCallbackA This is called before the state is rendered and before the stage is cleared but after all game objects have had their final properties adjusted.
func (self *StateManager) SetOnPreRenderCallbackA(member interface{}) {
    self.Object.Set("onPreRenderCallback", member)
}

// OnLoadUpdateCallback This is called when the State is updated during the preload phase.
func (self *StateManager) OnLoadUpdateCallback() interface{}{
    return self.Object.Get("onLoadUpdateCallback")
}

// SetOnLoadUpdateCallbackA This is called when the State is updated during the preload phase.
func (self *StateManager) SetOnLoadUpdateCallbackA(member interface{}) {
    self.Object.Set("onLoadUpdateCallback", member)
}

// OnLoadRenderCallback This is called when the State is rendered during the preload phase.
func (self *StateManager) OnLoadRenderCallback() interface{}{
    return self.Object.Get("onLoadRenderCallback")
}

// SetOnLoadRenderCallbackA This is called when the State is rendered during the preload phase.
func (self *StateManager) SetOnLoadRenderCallbackA(member interface{}) {
    self.Object.Set("onLoadRenderCallback", member)
}

// OnPausedCallback This is called when the game is paused.
func (self *StateManager) OnPausedCallback() interface{}{
    return self.Object.Get("onPausedCallback")
}

// SetOnPausedCallbackA This is called when the game is paused.
func (self *StateManager) SetOnPausedCallbackA(member interface{}) {
    self.Object.Set("onPausedCallback", member)
}

// OnResumedCallback This is called when the game is resumed from a paused state.
func (self *StateManager) OnResumedCallback() interface{}{
    return self.Object.Get("onResumedCallback")
}

// SetOnResumedCallbackA This is called when the game is resumed from a paused state.
func (self *StateManager) SetOnResumedCallbackA(member interface{}) {
    self.Object.Set("onResumedCallback", member)
}

// OnPauseUpdateCallback This is called every frame while the game is paused.
func (self *StateManager) OnPauseUpdateCallback() interface{}{
    return self.Object.Get("onPauseUpdateCallback")
}

// SetOnPauseUpdateCallbackA This is called every frame while the game is paused.
func (self *StateManager) SetOnPauseUpdateCallbackA(member interface{}) {
    self.Object.Set("onPauseUpdateCallback", member)
}

// OnShutDownCallback This is called when the state is shut down (i.e. swapped to another state).
func (self *StateManager) OnShutDownCallback() interface{}{
    return self.Object.Get("onShutDownCallback")
}

// SetOnShutDownCallbackA This is called when the state is shut down (i.e. swapped to another state).
func (self *StateManager) SetOnShutDownCallbackA(member interface{}) {
    self.Object.Set("onShutDownCallback", member)
}

// Created True if the current state has had its `create` method run (if it has one, if not this is true by default).
func (self *StateManager) Created() bool{
    return self.Object.Get("created").Bool()
}

// SetCreatedA True if the current state has had its `create` method run (if it has one, if not this is true by default).
func (self *StateManager) SetCreatedA(member bool) {
    self.Object.Set("created", member)
}


// Boot The Boot handler is called by Phaser.Game when it first starts up.
func (self *StateManager) Boot() {
    self.Object.Call("boot")
}

// BootI The Boot handler is called by Phaser.Game when it first starts up.
func (self *StateManager) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// Add Adds a new State into the StateManager. You must give each State a unique key by which you'll identify it.
// The State can be either a Phaser.State object (or an object that extends it), a plain JavaScript object or a function.
// If a function is given a new state object will be created by calling it.
func (self *StateManager) Add(key string, state interface{}) {
    self.Object.Call("add", key, state)
}

// Add1O Adds a new State into the StateManager. You must give each State a unique key by which you'll identify it.
// The State can be either a Phaser.State object (or an object that extends it), a plain JavaScript object or a function.
// If a function is given a new state object will be created by calling it.
func (self *StateManager) Add1O(key string, state interface{}, autoStart bool) {
    self.Object.Call("add", key, state, autoStart)
}

// AddI Adds a new State into the StateManager. You must give each State a unique key by which you'll identify it.
// The State can be either a Phaser.State object (or an object that extends it), a plain JavaScript object or a function.
// If a function is given a new state object will be created by calling it.
func (self *StateManager) AddI(args ...interface{}) {
    self.Object.Call("add", args)
}

// Remove Delete the given state.
func (self *StateManager) Remove(key string) {
    self.Object.Call("remove", key)
}

// RemoveI Delete the given state.
func (self *StateManager) RemoveI(args ...interface{}) {
    self.Object.Call("remove", args)
}

// Start Start the given State. If a State is already running then State.shutDown will be called (if it exists) before switching to the new State.
func (self *StateManager) Start(key string, clearWorld bool, clearCache bool, parameter interface{}) {
    self.Object.Call("start", key, clearWorld, clearCache, parameter)
}

// StartI Start the given State. If a State is already running then State.shutDown will be called (if it exists) before switching to the new State.
func (self *StateManager) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Restart Restarts the current State. State.shutDown will be called (if it exists) before the State is restarted.
func (self *StateManager) Restart(clearWorld bool, clearCache bool, parameter interface{}) {
    self.Object.Call("restart", clearWorld, clearCache, parameter)
}

// RestartI Restarts the current State. State.shutDown will be called (if it exists) before the State is restarted.
func (self *StateManager) RestartI(args ...interface{}) {
    self.Object.Call("restart", args)
}

// Dummy Used by onInit and onShutdown when those functions don't exist on the state
func (self *StateManager) Dummy() {
    self.Object.Call("dummy")
}

// DummyI Used by onInit and onShutdown when those functions don't exist on the state
func (self *StateManager) DummyI(args ...interface{}) {
    self.Object.Call("dummy", args)
}

// PreUpdate preUpdate is called right at the start of the game loop. It is responsible for changing to a new state that was requested previously.
func (self *StateManager) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI preUpdate is called right at the start of the game loop. It is responsible for changing to a new state that was requested previously.
func (self *StateManager) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// ClearCurrentState This method clears the current State, calling its shutdown callback. The process also removes any active tweens,
// resets the camera, resets input, clears physics, removes timers and if set clears the world and cache too.
func (self *StateManager) ClearCurrentState() {
    self.Object.Call("clearCurrentState")
}

// ClearCurrentStateI This method clears the current State, calling its shutdown callback. The process also removes any active tweens,
// resets the camera, resets input, clears physics, removes timers and if set clears the world and cache too.
func (self *StateManager) ClearCurrentStateI(args ...interface{}) {
    self.Object.Call("clearCurrentState", args)
}

// CheckState Checks if a given phaser state is valid. A State is considered valid if it has at least one of the core functions: preload, create, update or render.
func (self *StateManager) CheckState(key string) bool{
    return self.Object.Call("checkState", key).Bool()
}

// CheckStateI Checks if a given phaser state is valid. A State is considered valid if it has at least one of the core functions: preload, create, update or render.
func (self *StateManager) CheckStateI(args ...interface{}) bool{
    return self.Object.Call("checkState", args).Bool()
}

// Link Links game properties to the State given by the key.
func (self *StateManager) Link(key string) {
    self.Object.Call("link", key)
}

// LinkI Links game properties to the State given by the key.
func (self *StateManager) LinkI(args ...interface{}) {
    self.Object.Call("link", args)
}

// Unlink Nulls all State level Phaser properties, including a reference to Game.
func (self *StateManager) Unlink(key string) {
    self.Object.Call("unlink", key)
}

// UnlinkI Nulls all State level Phaser properties, including a reference to Game.
func (self *StateManager) UnlinkI(args ...interface{}) {
    self.Object.Call("unlink", args)
}

// SetCurrentState Sets the current State. Should not be called directly (use StateManager.start)
func (self *StateManager) SetCurrentState(key string) {
    self.Object.Call("setCurrentState", key)
}

// SetCurrentStateI Sets the current State. Should not be called directly (use StateManager.start)
func (self *StateManager) SetCurrentStateI(args ...interface{}) {
    self.Object.Call("setCurrentState", args)
}

// GetCurrentState Gets the current State.
func (self *StateManager) GetCurrentState() *State{
    return &State{self.Object.Call("getCurrentState")}
}

// GetCurrentStateI Gets the current State.
func (self *StateManager) GetCurrentStateI(args ...interface{}) *State{
    return &State{self.Object.Call("getCurrentState", args)}
}

// LoadComplete empty description
func (self *StateManager) LoadComplete() {
    self.Object.Call("loadComplete")
}

// LoadCompleteI empty description
func (self *StateManager) LoadCompleteI(args ...interface{}) {
    self.Object.Call("loadComplete", args)
}

// Pause empty description
func (self *StateManager) Pause() {
    self.Object.Call("pause")
}

// PauseI empty description
func (self *StateManager) PauseI(args ...interface{}) {
    self.Object.Call("pause", args)
}

// Resume empty description
func (self *StateManager) Resume() {
    self.Object.Call("resume")
}

// ResumeI empty description
func (self *StateManager) ResumeI(args ...interface{}) {
    self.Object.Call("resume", args)
}

// Update empty description
func (self *StateManager) Update() {
    self.Object.Call("update")
}

// UpdateI empty description
func (self *StateManager) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// PauseUpdate empty description
func (self *StateManager) PauseUpdate() {
    self.Object.Call("pauseUpdate")
}

// PauseUpdateI empty description
func (self *StateManager) PauseUpdateI(args ...interface{}) {
    self.Object.Call("pauseUpdate", args)
}

// PreRender empty description
func (self *StateManager) PreRender(elapsedTime int) {
    self.Object.Call("preRender", elapsedTime)
}

// PreRenderI empty description
func (self *StateManager) PreRenderI(args ...interface{}) {
    self.Object.Call("preRender", args)
}

// Resize empty description
func (self *StateManager) Resize() {
    self.Object.Call("resize")
}

// ResizeI empty description
func (self *StateManager) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Render empty description
func (self *StateManager) Render() {
    self.Object.Call("render")
}

// RenderI empty description
func (self *StateManager) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// Destroy Removes all StateManager callback references to the State object, nulls the game reference and clears the States object.
// You don't recover from this without rebuilding the Phaser instance again.
func (self *StateManager) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Removes all StateManager callback references to the State object, nulls the game reference and clears the States object.
// You don't recover from this without rebuilding the Phaser instance again.
func (self *StateManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

