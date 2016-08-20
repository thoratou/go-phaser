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


// Local reference to game.
func (self *Gamepad) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *Gamepad) SetGame(member *Game) {
    self.Set("game", member)
}

// Gamepad input will only be processed if enabled.
func (self *Gamepad) GetEnabled() bool{
    return self.Get("enabled").Bool()
}

// Gamepad input will only be processed if enabled.
func (self *Gamepad) SetEnabled(member bool) {
    self.Set("enabled", member)
}

// The context under which the callbacks are run.
func (self *Gamepad) GetCallbackContext() interface{}{
    return self.Get("callbackContext")
}

// The context under which the callbacks are run.
func (self *Gamepad) SetCallbackContext(member interface{}) {
    self.Set("callbackContext", member)
}

// This callback is invoked every time any gamepad is connected
func (self *Gamepad) SetOnConnectCallback(member func(...interface{})) {
    self.Set("onConnectCallback", member)
}

// This callback is invoked every time any gamepad is disconnected
func (self *Gamepad) SetOnDisconnectCallback(member func(...interface{})) {
    self.Set("onDisconnectCallback", member)
}

// This callback is invoked every time any gamepad button is pressed down.
func (self *Gamepad) SetOnDownCallback(member func(...interface{})) {
    self.Set("onDownCallback", member)
}

// This callback is invoked every time any gamepad button is released.
func (self *Gamepad) SetOnUpCallback(member func(...interface{})) {
    self.Set("onUpCallback", member)
}

// This callback is invoked every time any gamepad axis is changed.
func (self *Gamepad) SetOnAxisCallback(member func(...interface{})) {
    self.Set("onAxisCallback", member)
}

// This callback is invoked every time any gamepad button is changed to a value where value > 0 and value < 1.
func (self *Gamepad) SetOnFloatCallback(member func(...interface{})) {
    self.Set("onFloatCallback", member)
}

// If the gamepad input is active or not - if not active it should not be updated from Input.js
func (self *Gamepad) GetActive() bool{
    return self.Get("active").Bool()
}

// If the gamepad input is active or not - if not active it should not be updated from Input.js
func (self *Gamepad) SetActive(member bool) {
    self.Set("active", member)
}

// Whether or not gamepads are supported in current browser.
func (self *Gamepad) GetSupported() bool{
    return self.Get("supported").Bool()
}

// Whether or not gamepads are supported in current browser.
func (self *Gamepad) SetSupported(member bool) {
    self.Set("supported", member)
}

// How many live gamepads are currently connected.
func (self *Gamepad) GetPadsConnected() float64{
    return self.Get("padsConnected").Float()
}

// How many live gamepads are currently connected.
func (self *Gamepad) SetPadsConnected(member float64) {
    self.Set("padsConnected", member)
}

// Gamepad #1
func (self *Gamepad) GetPad1() *SinglePad{
    return &SinglePad{self.Get("pad1")}
}

// Gamepad #1
func (self *Gamepad) SetPad1(member *SinglePad) {
    self.Set("pad1", member)
}

// Gamepad #2
func (self *Gamepad) GetPad2() *SinglePad{
    return &SinglePad{self.Get("pad2")}
}

// Gamepad #2
func (self *Gamepad) SetPad2(member *SinglePad) {
    self.Set("pad2", member)
}

// Gamepad #3
func (self *Gamepad) GetPad3() *SinglePad{
    return &SinglePad{self.Get("pad3")}
}

// Gamepad #3
func (self *Gamepad) SetPad3(member *SinglePad) {
    self.Set("pad3", member)
}

// Gamepad #4
func (self *Gamepad) GetPad4() *SinglePad{
    return &SinglePad{self.Get("pad4")}
}

// Gamepad #4
func (self *Gamepad) SetPad4(member *SinglePad) {
    self.Set("pad4", member)
}



// Add callbacks to the main Gamepad handler to handle connect/disconnect/button down/button up/axis change/float value buttons.
func (self *Gamepad) AddCallbacksI(args ...interface{}) {
    self.Call("addCallbacks", args)
}

// Starts the Gamepad event handling.
// This MUST be called manually before Phaser will start polling the Gamepad API.
func (self *Gamepad) StartI(args ...interface{}) {
    self.Call("start", args)
}

// Main gamepad update loop. Should not be called manually.
func (self *Gamepad) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Updating connected gamepads (for Google Chrome). Should not be called manually.
func (self *Gamepad) _pollGamepadsI(args ...interface{}) {
    self.Call("_pollGamepads", args)
}

// Sets the deadZone variable for all four gamepads
func (self *Gamepad) SetDeadZonesI(args ...interface{}) {
    self.Call("setDeadZones", args)
}

// Stops the Gamepad event handling.
func (self *Gamepad) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// Reset all buttons/axes of all gamepads
func (self *Gamepad) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Returns the "just pressed" state of a button from ANY gamepad connected. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *Gamepad) JustPressedI(args ...interface{}) bool{
    return self.Call("justPressed", args).Bool()
}

// Returns true if the button is currently pressed down, on ANY gamepad.
func (self *Gamepad) IsDownI(args ...interface{}) bool{
    return self.Call("isDown", args).Bool()
}

// Destroys this object and the associated event listeners.
func (self *Gamepad) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
