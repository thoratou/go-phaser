// Automatic generation for Phaser.SinglePad
// generated file SinglePad.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A single Phaser Gamepad
type SinglePad struct {
    *js.Object
}


// Local reference to game.
func (self *SinglePad) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *SinglePad) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The gamepad index as per browsers data
func (self *SinglePad) GetIndexA() int{
    return self.Object.Get("index").Int()
}

// The gamepad index as per browsers data
func (self *SinglePad) SetIndexA(member int) {
    self.Object.Set("index", member)
}

// Whether or not this particular gamepad is connected or not.
func (self *SinglePad) GetConnectedA() bool{
    return self.Object.Get("connected").Bool()
}

// Whether or not this particular gamepad is connected or not.
func (self *SinglePad) SetConnectedA(member bool) {
    self.Object.Set("connected", member)
}

// The context under which the callbacks are run.
func (self *SinglePad) GetCallbackContextA() interface{}{
    return self.Object.Get("callbackContext")
}

// The context under which the callbacks are run.
func (self *SinglePad) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// This callback is invoked every time this gamepad is connected
func (self *SinglePad) SetOnConnectCallbackA(member func(...interface{})) {
    self.Object.Set("onConnectCallback", member)
}

// This callback is invoked every time this gamepad is disconnected
func (self *SinglePad) SetOnDisconnectCallbackA(member func(...interface{})) {
    self.Object.Set("onDisconnectCallback", member)
}

// This callback is invoked every time a button is pressed down.
func (self *SinglePad) SetOnDownCallbackA(member func(...interface{})) {
    self.Object.Set("onDownCallback", member)
}

// This callback is invoked every time a gamepad button is released.
func (self *SinglePad) SetOnUpCallbackA(member func(...interface{})) {
    self.Object.Set("onUpCallback", member)
}

// This callback is invoked every time an axis is changed.
func (self *SinglePad) SetOnAxisCallbackA(member func(...interface{})) {
    self.Object.Set("onAxisCallback", member)
}

// This callback is invoked every time a button is changed to a value where value > 0 and value < 1.
func (self *SinglePad) SetOnFloatCallbackA(member func(...interface{})) {
    self.Object.Set("onFloatCallback", member)
}

// Dead zone for axis feedback - within this value you won't trigger updates.
func (self *SinglePad) GetDeadZoneA() int{
    return self.Object.Get("deadZone").Int()
}

// Dead zone for axis feedback - within this value you won't trigger updates.
func (self *SinglePad) SetDeadZoneA(member int) {
    self.Object.Set("deadZone", member)
}



// Add callbacks to this Gamepad to handle connect / disconnect / button down / button up / axis change / float value buttons.
func (self *SinglePad) AddCallbacks(context interface{}, callbacks interface{}) {
    self.Object.Call("addCallbacks", context, callbacks)
}

// Add callbacks to this Gamepad to handle connect / disconnect / button down / button up / axis change / float value buttons.
func (self *SinglePad) AddCallbacksI(args ...interface{}) {
    self.Object.Call("addCallbacks", args)
}

// Gets a DeviceButton object from this controller to be stored and referenced locally.
// The DeviceButton object can then be polled, have events attached to it, etc.
func (self *SinglePad) GetButton(buttonCode int) *DeviceButton{
    return &DeviceButton{self.Object.Call("getButton", buttonCode)}
}

// Gets a DeviceButton object from this controller to be stored and referenced locally.
// The DeviceButton object can then be polled, have events attached to it, etc.
func (self *SinglePad) GetButtonI(args ...interface{}) *DeviceButton{
    return &DeviceButton{self.Object.Call("getButton", args)}
}

// Main update function called by Phaser.Gamepad.
func (self *SinglePad) PollStatus() {
    self.Object.Call("pollStatus")
}

// Main update function called by Phaser.Gamepad.
func (self *SinglePad) PollStatusI(args ...interface{}) {
    self.Object.Call("pollStatus", args)
}

// Gamepad connect function, should be called by Phaser.Gamepad.
func (self *SinglePad) Connect(rawPad interface{}) {
    self.Object.Call("connect", rawPad)
}

// Gamepad connect function, should be called by Phaser.Gamepad.
func (self *SinglePad) ConnectI(args ...interface{}) {
    self.Object.Call("connect", args)
}

// Gamepad disconnect function, should be called by Phaser.Gamepad.
func (self *SinglePad) Disconnect() {
    self.Object.Call("disconnect")
}

// Gamepad disconnect function, should be called by Phaser.Gamepad.
func (self *SinglePad) DisconnectI(args ...interface{}) {
    self.Object.Call("disconnect", args)
}

// Destroys this object and associated callback references.
func (self *SinglePad) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this object and associated callback references.
func (self *SinglePad) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Handles changes in axis.
func (self *SinglePad) ProcessAxisChange(axisState interface{}) {
    self.Object.Call("processAxisChange", axisState)
}

// Handles changes in axis.
func (self *SinglePad) ProcessAxisChangeI(args ...interface{}) {
    self.Object.Call("processAxisChange", args)
}

// Handles button down press.
func (self *SinglePad) ProcessButtonDown(buttonCode int, value interface{}) {
    self.Object.Call("processButtonDown", buttonCode, value)
}

// Handles button down press.
func (self *SinglePad) ProcessButtonDownI(args ...interface{}) {
    self.Object.Call("processButtonDown", args)
}

// Handles button release.
func (self *SinglePad) ProcessButtonUp(buttonCode int, value interface{}) {
    self.Object.Call("processButtonUp", buttonCode, value)
}

// Handles button release.
func (self *SinglePad) ProcessButtonUpI(args ...interface{}) {
    self.Object.Call("processButtonUp", args)
}

// Handles buttons with floating values (like analog buttons that acts almost like an axis but still registers like a button)
func (self *SinglePad) ProcessButtonFloat(buttonCode int, value interface{}) {
    self.Object.Call("processButtonFloat", buttonCode, value)
}

// Handles buttons with floating values (like analog buttons that acts almost like an axis but still registers like a button)
func (self *SinglePad) ProcessButtonFloatI(args ...interface{}) {
    self.Object.Call("processButtonFloat", args)
}

// Returns value of requested axis.
func (self *SinglePad) Axis(axisCode int) int{
    return self.Object.Call("axis", axisCode).Int()
}

// Returns value of requested axis.
func (self *SinglePad) AxisI(args ...interface{}) int{
    return self.Object.Call("axis", args).Int()
}

// Returns true if the button is pressed down.
func (self *SinglePad) IsDown(buttonCode int) bool{
    return self.Object.Call("isDown", buttonCode).Bool()
}

// Returns true if the button is pressed down.
func (self *SinglePad) IsDownI(args ...interface{}) bool{
    return self.Object.Call("isDown", args).Bool()
}

// Returns true if the button is not currently pressed.
func (self *SinglePad) IsUp(buttonCode int) bool{
    return self.Object.Call("isUp", buttonCode).Bool()
}

// Returns true if the button is not currently pressed.
func (self *SinglePad) IsUpI(args ...interface{}) bool{
    return self.Object.Call("isUp", args).Bool()
}

// Returns the "just released" state of a button from this gamepad. Just released is considered as being true if the button was released within the duration given (default 250ms).
func (self *SinglePad) JustReleased(buttonCode int, duration int) bool{
    return self.Object.Call("justReleased", buttonCode, duration).Bool()
}

// Returns the "just released" state of a button from this gamepad. Just released is considered as being true if the button was released within the duration given (default 250ms).
func (self *SinglePad) JustReleasedI(args ...interface{}) bool{
    return self.Object.Call("justReleased", args).Bool()
}

// Returns the "just pressed" state of a button from this gamepad. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *SinglePad) JustPressed(buttonCode int, duration int) bool{
    return self.Object.Call("justPressed", buttonCode, duration).Bool()
}

// Returns the "just pressed" state of a button from this gamepad. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *SinglePad) JustPressedI(args ...interface{}) bool{
    return self.Object.Call("justPressed", args).Bool()
}

// Returns the value of a gamepad button. Intended mainly for cases when you have floating button values, for example
// analog trigger buttons on the XBOX 360 controller.
func (self *SinglePad) ButtonValue(buttonCode int) int{
    return self.Object.Call("buttonValue", buttonCode).Int()
}

// Returns the value of a gamepad button. Intended mainly for cases when you have floating button values, for example
// analog trigger buttons on the XBOX 360 controller.
func (self *SinglePad) ButtonValueI(args ...interface{}) int{
    return self.Object.Call("buttonValue", args).Int()
}

// Reset all buttons/axes of this gamepad.
func (self *SinglePad) Reset() {
    self.Object.Call("reset")
}

// Reset all buttons/axes of this gamepad.
func (self *SinglePad) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}
