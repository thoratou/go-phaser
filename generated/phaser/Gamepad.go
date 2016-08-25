// Package phaser Automatic generation for Phaser.Gamepad
// generated file Gamepad.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Gamepad The Gamepad class handles gamepad input and dispatches gamepad events.
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

// NewGamepad The Gamepad class handles gamepad input and dispatches gamepad events.
// 
// Remember to call `gamepad.start()`.
// 
// HTML5 GAMEPAD API SUPPORT IS AT AN EXPERIMENTAL STAGE!
// At moment of writing this (end of 2013) only Chrome supports parts of it out of the box. Firefox supports it
// via prefs flags (about:config, search gamepad). The browsers map the same controllers differently.
// This class has constants for Windows 7 Chrome mapping of XBOX 360 controller.
func NewGamepad(game *Game) *Gamepad {
    return &Gamepad{js.Global.Get("Phaser").Get("Gamepad").New(game)}
}
// NewGamepadI The Gamepad class handles gamepad input and dispatches gamepad events.
// 
// Remember to call `gamepad.start()`.
// 
// HTML5 GAMEPAD API SUPPORT IS AT AN EXPERIMENTAL STAGE!
// At moment of writing this (end of 2013) only Chrome supports parts of it out of the box. Firefox supports it
// via prefs flags (about:config, search gamepad). The browsers map the same controllers differently.
// This class has constants for Windows 7 Chrome mapping of XBOX 360 controller.
func NewGamepadI(args ...interface{}) *Gamepad {
    return &Gamepad{js.Global.Get("Phaser").Get("Gamepad").New(args)}
}



// Gamepad Binding conversion method to Gamepad point 
func ToGamepad(jsStruct interface{}) *Gamepad {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Gamepad{Object: object}
	}
	return nil
}



// Game Local reference to game.
func (self *Gamepad) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *Gamepad) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Enabled Gamepad input will only be processed if enabled.
func (self *Gamepad) Enabled() bool{
    return self.Object.Get("enabled").Bool()
}

// SetEnabledA Gamepad input will only be processed if enabled.
func (self *Gamepad) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}

// CallbackContext The context under which the callbacks are run.
func (self *Gamepad) CallbackContext() interface{}{
    return self.Object.Get("callbackContext")
}

// SetCallbackContextA The context under which the callbacks are run.
func (self *Gamepad) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// OnConnectCallback This callback is invoked every time any gamepad is connected
func (self *Gamepad) OnConnectCallback() interface{}{
    return self.Object.Get("onConnectCallback")
}

// SetOnConnectCallbackA This callback is invoked every time any gamepad is connected
func (self *Gamepad) SetOnConnectCallbackA(member interface{}) {
    self.Object.Set("onConnectCallback", member)
}

// OnDisconnectCallback This callback is invoked every time any gamepad is disconnected
func (self *Gamepad) OnDisconnectCallback() interface{}{
    return self.Object.Get("onDisconnectCallback")
}

// SetOnDisconnectCallbackA This callback is invoked every time any gamepad is disconnected
func (self *Gamepad) SetOnDisconnectCallbackA(member interface{}) {
    self.Object.Set("onDisconnectCallback", member)
}

// OnDownCallback This callback is invoked every time any gamepad button is pressed down.
func (self *Gamepad) OnDownCallback() interface{}{
    return self.Object.Get("onDownCallback")
}

// SetOnDownCallbackA This callback is invoked every time any gamepad button is pressed down.
func (self *Gamepad) SetOnDownCallbackA(member interface{}) {
    self.Object.Set("onDownCallback", member)
}

// OnUpCallback This callback is invoked every time any gamepad button is released.
func (self *Gamepad) OnUpCallback() interface{}{
    return self.Object.Get("onUpCallback")
}

// SetOnUpCallbackA This callback is invoked every time any gamepad button is released.
func (self *Gamepad) SetOnUpCallbackA(member interface{}) {
    self.Object.Set("onUpCallback", member)
}

// OnAxisCallback This callback is invoked every time any gamepad axis is changed.
func (self *Gamepad) OnAxisCallback() interface{}{
    return self.Object.Get("onAxisCallback")
}

// SetOnAxisCallbackA This callback is invoked every time any gamepad axis is changed.
func (self *Gamepad) SetOnAxisCallbackA(member interface{}) {
    self.Object.Set("onAxisCallback", member)
}

// OnFloatCallback This callback is invoked every time any gamepad button is changed to a value where value > 0 and value < 1.
func (self *Gamepad) OnFloatCallback() interface{}{
    return self.Object.Get("onFloatCallback")
}

// SetOnFloatCallbackA This callback is invoked every time any gamepad button is changed to a value where value > 0 and value < 1.
func (self *Gamepad) SetOnFloatCallbackA(member interface{}) {
    self.Object.Set("onFloatCallback", member)
}

// Active If the gamepad input is active or not - if not active it should not be updated from Input.js
func (self *Gamepad) Active() bool{
    return self.Object.Get("active").Bool()
}

// SetActiveA If the gamepad input is active or not - if not active it should not be updated from Input.js
func (self *Gamepad) SetActiveA(member bool) {
    self.Object.Set("active", member)
}

// Supported Whether or not gamepads are supported in current browser.
func (self *Gamepad) Supported() bool{
    return self.Object.Get("supported").Bool()
}

// SetSupportedA Whether or not gamepads are supported in current browser.
func (self *Gamepad) SetSupportedA(member bool) {
    self.Object.Set("supported", member)
}

// PadsConnected How many live gamepads are currently connected.
func (self *Gamepad) PadsConnected() int{
    return self.Object.Get("padsConnected").Int()
}

// SetPadsConnectedA How many live gamepads are currently connected.
func (self *Gamepad) SetPadsConnectedA(member int) {
    self.Object.Set("padsConnected", member)
}

// Pad1 Gamepad #1
func (self *Gamepad) Pad1() *SinglePad{
    return &SinglePad{self.Object.Get("pad1")}
}

// SetPad1A Gamepad #1
func (self *Gamepad) SetPad1A(member *SinglePad) {
    self.Object.Set("pad1", member)
}

// Pad2 Gamepad #2
func (self *Gamepad) Pad2() *SinglePad{
    return &SinglePad{self.Object.Get("pad2")}
}

// SetPad2A Gamepad #2
func (self *Gamepad) SetPad2A(member *SinglePad) {
    self.Object.Set("pad2", member)
}

// Pad3 Gamepad #3
func (self *Gamepad) Pad3() *SinglePad{
    return &SinglePad{self.Object.Get("pad3")}
}

// SetPad3A Gamepad #3
func (self *Gamepad) SetPad3A(member *SinglePad) {
    self.Object.Set("pad3", member)
}

// Pad4 Gamepad #4
func (self *Gamepad) Pad4() *SinglePad{
    return &SinglePad{self.Object.Get("pad4")}
}

// SetPad4A Gamepad #4
func (self *Gamepad) SetPad4A(member *SinglePad) {
    self.Object.Set("pad4", member)
}


// AddCallbacks Add callbacks to the main Gamepad handler to handle connect/disconnect/button down/button up/axis change/float value buttons.
func (self *Gamepad) AddCallbacks(context interface{}, callbacks interface{}) {
    self.Object.Call("addCallbacks", context, callbacks)
}

// AddCallbacksI Add callbacks to the main Gamepad handler to handle connect/disconnect/button down/button up/axis change/float value buttons.
func (self *Gamepad) AddCallbacksI(args ...interface{}) {
    self.Object.Call("addCallbacks", args)
}

// Start Starts the Gamepad event handling.
// This MUST be called manually before Phaser will start polling the Gamepad API.
func (self *Gamepad) Start() {
    self.Object.Call("start")
}

// StartI Starts the Gamepad event handling.
// This MUST be called manually before Phaser will start polling the Gamepad API.
func (self *Gamepad) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Update Main gamepad update loop. Should not be called manually.
func (self *Gamepad) Update() {
    self.Object.Call("update")
}

// UpdateI Main gamepad update loop. Should not be called manually.
func (self *Gamepad) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// _pollGamepads Updating connected gamepads (for Google Chrome). Should not be called manually.
func (self *Gamepad) _pollGamepads() {
    self.Object.Call("_pollGamepads")
}

// _pollGamepadsI Updating connected gamepads (for Google Chrome). Should not be called manually.
func (self *Gamepad) _pollGamepadsI(args ...interface{}) {
    self.Object.Call("_pollGamepads", args)
}

// SetDeadZones Sets the deadZone variable for all four gamepads
func (self *Gamepad) SetDeadZones() {
    self.Object.Call("setDeadZones")
}

// SetDeadZonesI Sets the deadZone variable for all four gamepads
func (self *Gamepad) SetDeadZonesI(args ...interface{}) {
    self.Object.Call("setDeadZones", args)
}

// Stop Stops the Gamepad event handling.
func (self *Gamepad) Stop() {
    self.Object.Call("stop")
}

// StopI Stops the Gamepad event handling.
func (self *Gamepad) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Reset Reset all buttons/axes of all gamepads
func (self *Gamepad) Reset() {
    self.Object.Call("reset")
}

// ResetI Reset all buttons/axes of all gamepads
func (self *Gamepad) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// JustPressed Returns the "just pressed" state of a button from ANY gamepad connected. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *Gamepad) JustPressed(buttonCode int) bool{
    return self.Object.Call("justPressed", buttonCode).Bool()
}

// JustPressed1O Returns the "just pressed" state of a button from ANY gamepad connected. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *Gamepad) JustPressed1O(buttonCode int, duration int) bool{
    return self.Object.Call("justPressed", buttonCode, duration).Bool()
}

// JustPressedI Returns the "just pressed" state of a button from ANY gamepad connected. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *Gamepad) JustPressedI(args ...interface{}) bool{
    return self.Object.Call("justPressed", args).Bool()
}

// IsDown Returns true if the button is currently pressed down, on ANY gamepad.
func (self *Gamepad) IsDown(buttonCode int) bool{
    return self.Object.Call("isDown", buttonCode).Bool()
}

// IsDownI Returns true if the button is currently pressed down, on ANY gamepad.
func (self *Gamepad) IsDownI(args ...interface{}) bool{
    return self.Object.Call("isDown", args).Bool()
}

// Destroy Destroys this object and the associated event listeners.
func (self *Gamepad) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this object and the associated event listeners.
func (self *Gamepad) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

