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
func (self *SinglePad) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *SinglePad) SetGame(member *Game) {
    self.Set("game", member)
}

// The gamepad index as per browsers data
func (self *SinglePad) GetIndex() int{
    return self.Get("index").Int()
}

// The gamepad index as per browsers data
func (self *SinglePad) SetIndex(member int) {
    self.Set("index", member)
}

// Whether or not this particular gamepad is connected or not.
func (self *SinglePad) GetConnected() bool{
    return self.Get("connected").Bool()
}

// Whether or not this particular gamepad is connected or not.
func (self *SinglePad) SetConnected(member bool) {
    self.Set("connected", member)
}

// The context under which the callbacks are run.
func (self *SinglePad) GetCallbackContext() interface{}{
    return self.Get("callbackContext")
}

// The context under which the callbacks are run.
func (self *SinglePad) SetCallbackContext(member interface{}) {
    self.Set("callbackContext", member)
}

// This callback is invoked every time this gamepad is connected
func (self *SinglePad) SetOnConnectCallback(member func(...interface{})) {
    self.Set("onConnectCallback", member)
}

// This callback is invoked every time this gamepad is disconnected
func (self *SinglePad) SetOnDisconnectCallback(member func(...interface{})) {
    self.Set("onDisconnectCallback", member)
}

// This callback is invoked every time a button is pressed down.
func (self *SinglePad) SetOnDownCallback(member func(...interface{})) {
    self.Set("onDownCallback", member)
}

// This callback is invoked every time a gamepad button is released.
func (self *SinglePad) SetOnUpCallback(member func(...interface{})) {
    self.Set("onUpCallback", member)
}

// This callback is invoked every time an axis is changed.
func (self *SinglePad) SetOnAxisCallback(member func(...interface{})) {
    self.Set("onAxisCallback", member)
}

// This callback is invoked every time a button is changed to a value where value > 0 and value < 1.
func (self *SinglePad) SetOnFloatCallback(member func(...interface{})) {
    self.Set("onFloatCallback", member)
}

// Dead zone for axis feedback - within this value you won't trigger updates.
func (self *SinglePad) GetDeadZone() int{
    return self.Get("deadZone").Int()
}

// Dead zone for axis feedback - within this value you won't trigger updates.
func (self *SinglePad) SetDeadZone(member int) {
    self.Set("deadZone", member)
}



// Add callbacks to this Gamepad to handle connect / disconnect / button down / button up / axis change / float value buttons.
func (self *SinglePad) AddCallbacksI(args ...interface{}) {
    self.Call("addCallbacks", args)
}

// Gets a DeviceButton object from this controller to be stored and referenced locally.
// The DeviceButton object can then be polled, have events attached to it, etc.
func (self *SinglePad) GetButtonI(args ...interface{}) *DeviceButton{
    return &DeviceButton{self.Call("getButton", args)}
}

// Main update function called by Phaser.Gamepad.
func (self *SinglePad) PollStatusI(args ...interface{}) {
    self.Call("pollStatus", args)
}

// Gamepad connect function, should be called by Phaser.Gamepad.
func (self *SinglePad) ConnectI(args ...interface{}) {
    self.Call("connect", args)
}

// Gamepad disconnect function, should be called by Phaser.Gamepad.
func (self *SinglePad) DisconnectI(args ...interface{}) {
    self.Call("disconnect", args)
}

// Destroys this object and associated callback references.
func (self *SinglePad) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Handles changes in axis.
func (self *SinglePad) ProcessAxisChangeI(args ...interface{}) {
    self.Call("processAxisChange", args)
}

// Handles button down press.
func (self *SinglePad) ProcessButtonDownI(args ...interface{}) {
    self.Call("processButtonDown", args)
}

// Handles button release.
func (self *SinglePad) ProcessButtonUpI(args ...interface{}) {
    self.Call("processButtonUp", args)
}

// Handles buttons with floating values (like analog buttons that acts almost like an axis but still registers like a button)
func (self *SinglePad) ProcessButtonFloatI(args ...interface{}) {
    self.Call("processButtonFloat", args)
}

// Returns value of requested axis.
func (self *SinglePad) AxisI(args ...interface{}) int{
    return self.Call("axis", args).Int()
}

// Returns true if the button is pressed down.
func (self *SinglePad) IsDownI(args ...interface{}) bool{
    return self.Call("isDown", args).Bool()
}

// Returns true if the button is not currently pressed.
func (self *SinglePad) IsUpI(args ...interface{}) bool{
    return self.Call("isUp", args).Bool()
}

// Returns the "just released" state of a button from this gamepad. Just released is considered as being true if the button was released within the duration given (default 250ms).
func (self *SinglePad) JustReleasedI(args ...interface{}) bool{
    return self.Call("justReleased", args).Bool()
}

// Returns the "just pressed" state of a button from this gamepad. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *SinglePad) JustPressedI(args ...interface{}) bool{
    return self.Call("justPressed", args).Bool()
}

// Returns the value of a gamepad button. Intended mainly for cases when you have floating button values, for example
// analog trigger buttons on the XBOX 360 controller.
func (self *SinglePad) ButtonValueI(args ...interface{}) int{
    return self.Call("buttonValue", args).Int()
}

// Reset all buttons/axes of this gamepad.
func (self *SinglePad) ResetI(args ...interface{}) {
    self.Call("reset", args)
}
