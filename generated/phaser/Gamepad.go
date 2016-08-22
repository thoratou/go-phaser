// Automatic generation for Phaser.Gamepad
// generated file Gamepad.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Gamepad class handles gamepad input and dispatches gamepad events.
// 
// Remember to call `gamepad.start()`.
// 
// HTML5 GAMEPAD API SUPPORT IS AT AN EXPERIMENTAL STAGE!
// At moment of writing this (end of 2013) only Chrome supports parts of it out of the box. Firefox supports it
// via prefs flags (about:config, search gamepad). The browsers map the same controllers differently.
// This class has constants for Windows 7 Chrome mapping of XBOX 360 controller.
type Gamepad struct {
    *js.Object
}


// The Gamepad class handles gamepad input and dispatches gamepad events.
// 
// Remember to call `gamepad.start()`.
// 
// HTML5 GAMEPAD API SUPPORT IS AT AN EXPERIMENTAL STAGE!
// At moment of writing this (end of 2013) only Chrome supports parts of it out of the box. Firefox supports it
// via prefs flags (about:config, search gamepad). The browsers map the same controllers differently.
// This class has constants for Windows 7 Chrome mapping of XBOX 360 controller.
func NewGamepad(game *Game) *Gamepad {
    return &Gamepad{js.Global.Call("Phaser.Gamepad", game)}
}

// The Gamepad class handles gamepad input and dispatches gamepad events.
// 
// Remember to call `gamepad.start()`.
// 
// HTML5 GAMEPAD API SUPPORT IS AT AN EXPERIMENTAL STAGE!
// At moment of writing this (end of 2013) only Chrome supports parts of it out of the box. Firefox supports it
// via prefs flags (about:config, search gamepad). The browsers map the same controllers differently.
// This class has constants for Windows 7 Chrome mapping of XBOX 360 controller.
func NewGamepadI(args ...interface{}) *Gamepad {
    return &Gamepad{js.Global.Call("Phaser.Gamepad", args)}
}



// Local reference to game.
func (self *Gamepad) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *Gamepad) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Gamepad input will only be processed if enabled.
func (self *Gamepad) GetEnabledA() bool{
    return self.Object.Get("enabled").Bool()
}

// Gamepad input will only be processed if enabled.
func (self *Gamepad) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}

// The context under which the callbacks are run.
func (self *Gamepad) GetCallbackContextA() interface{}{
    return self.Object.Get("callbackContext")
}

// The context under which the callbacks are run.
func (self *Gamepad) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// This callback is invoked every time any gamepad is connected
func (self *Gamepad) SetOnConnectCallbackA(member func(...interface{})) {
    self.Object.Set("onConnectCallback", member)
}

// This callback is invoked every time any gamepad is disconnected
func (self *Gamepad) SetOnDisconnectCallbackA(member func(...interface{})) {
    self.Object.Set("onDisconnectCallback", member)
}

// This callback is invoked every time any gamepad button is pressed down.
func (self *Gamepad) SetOnDownCallbackA(member func(...interface{})) {
    self.Object.Set("onDownCallback", member)
}

// This callback is invoked every time any gamepad button is released.
func (self *Gamepad) SetOnUpCallbackA(member func(...interface{})) {
    self.Object.Set("onUpCallback", member)
}

// This callback is invoked every time any gamepad axis is changed.
func (self *Gamepad) SetOnAxisCallbackA(member func(...interface{})) {
    self.Object.Set("onAxisCallback", member)
}

// This callback is invoked every time any gamepad button is changed to a value where value > 0 and value < 1.
func (self *Gamepad) SetOnFloatCallbackA(member func(...interface{})) {
    self.Object.Set("onFloatCallback", member)
}

// If the gamepad input is active or not - if not active it should not be updated from Input.js
func (self *Gamepad) GetActiveA() bool{
    return self.Object.Get("active").Bool()
}

// If the gamepad input is active or not - if not active it should not be updated from Input.js
func (self *Gamepad) SetActiveA(member bool) {
    self.Object.Set("active", member)
}

// Whether or not gamepads are supported in current browser.
func (self *Gamepad) GetSupportedA() bool{
    return self.Object.Get("supported").Bool()
}

// Whether or not gamepads are supported in current browser.
func (self *Gamepad) SetSupportedA(member bool) {
    self.Object.Set("supported", member)
}

// How many live gamepads are currently connected.
func (self *Gamepad) GetPadsConnectedA() int{
    return self.Object.Get("padsConnected").Int()
}

// How many live gamepads are currently connected.
func (self *Gamepad) SetPadsConnectedA(member int) {
    self.Object.Set("padsConnected", member)
}

// Gamepad #1
func (self *Gamepad) GetPad1A() *SinglePad{
    return &SinglePad{self.Object.Get("pad1")}
}

// Gamepad #1
func (self *Gamepad) SetPad1A(member *SinglePad) {
    self.Object.Set("pad1", member)
}

// Gamepad #2
func (self *Gamepad) GetPad2A() *SinglePad{
    return &SinglePad{self.Object.Get("pad2")}
}

// Gamepad #2
func (self *Gamepad) SetPad2A(member *SinglePad) {
    self.Object.Set("pad2", member)
}

// Gamepad #3
func (self *Gamepad) GetPad3A() *SinglePad{
    return &SinglePad{self.Object.Get("pad3")}
}

// Gamepad #3
func (self *Gamepad) SetPad3A(member *SinglePad) {
    self.Object.Set("pad3", member)
}

// Gamepad #4
func (self *Gamepad) GetPad4A() *SinglePad{
    return &SinglePad{self.Object.Get("pad4")}
}

// Gamepad #4
func (self *Gamepad) SetPad4A(member *SinglePad) {
    self.Object.Set("pad4", member)
}



// Add callbacks to the main Gamepad handler to handle connect/disconnect/button down/button up/axis change/float value buttons.
func (self *Gamepad) AddCallbacks(context interface{}, callbacks interface{}) {
    self.Object.Call("addCallbacks", context, callbacks)
}

// Add callbacks to the main Gamepad handler to handle connect/disconnect/button down/button up/axis change/float value buttons.
func (self *Gamepad) AddCallbacksI(args ...interface{}) {
    self.Object.Call("addCallbacks", args)
}

// Starts the Gamepad event handling.
// This MUST be called manually before Phaser will start polling the Gamepad API.
func (self *Gamepad) Start() {
    self.Object.Call("start")
}

// Starts the Gamepad event handling.
// This MUST be called manually before Phaser will start polling the Gamepad API.
func (self *Gamepad) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Main gamepad update loop. Should not be called manually.
func (self *Gamepad) Update() {
    self.Object.Call("update")
}

// Main gamepad update loop. Should not be called manually.
func (self *Gamepad) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Updating connected gamepads (for Google Chrome). Should not be called manually.
func (self *Gamepad) _pollGamepads() {
    self.Object.Call("_pollGamepads")
}

// Updating connected gamepads (for Google Chrome). Should not be called manually.
func (self *Gamepad) _pollGamepadsI(args ...interface{}) {
    self.Object.Call("_pollGamepads", args)
}

// Sets the deadZone variable for all four gamepads
func (self *Gamepad) SetDeadZones() {
    self.Object.Call("setDeadZones")
}

// Sets the deadZone variable for all four gamepads
func (self *Gamepad) SetDeadZonesI(args ...interface{}) {
    self.Object.Call("setDeadZones", args)
}

// Stops the Gamepad event handling.
func (self *Gamepad) Stop() {
    self.Object.Call("stop")
}

// Stops the Gamepad event handling.
func (self *Gamepad) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Reset all buttons/axes of all gamepads
func (self *Gamepad) Reset() {
    self.Object.Call("reset")
}

// Reset all buttons/axes of all gamepads
func (self *Gamepad) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Returns the "just pressed" state of a button from ANY gamepad connected. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *Gamepad) JustPressed(buttonCode int) bool{
    return self.Object.Call("justPressed", buttonCode).Bool()
}

// Returns the "just pressed" state of a button from ANY gamepad connected. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *Gamepad) JustPressed1O(buttonCode int, duration int) bool{
    return self.Object.Call("justPressed", buttonCode, duration).Bool()
}

// Returns the "just pressed" state of a button from ANY gamepad connected. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *Gamepad) JustPressedI(args ...interface{}) bool{
    return self.Object.Call("justPressed", args).Bool()
}

// Returns true if the button is currently pressed down, on ANY gamepad.
func (self *Gamepad) IsDown(buttonCode int) bool{
    return self.Object.Call("isDown", buttonCode).Bool()
}

// Returns true if the button is currently pressed down, on ANY gamepad.
func (self *Gamepad) IsDownI(args ...interface{}) bool{
    return self.Object.Call("isDown", args).Bool()
}

// Destroys this object and the associated event listeners.
func (self *Gamepad) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this object and the associated event listeners.
func (self *Gamepad) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
