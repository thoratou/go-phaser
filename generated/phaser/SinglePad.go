// Package phaser Automatic generation for Phaser.SinglePad
// generated file SinglePad.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// SinglePad A single Phaser Gamepad
type SinglePad struct {
    *js.Object
}

// NewSinglePad A single Phaser Gamepad
func NewSinglePad(game *Game, padParent interface{}) *SinglePad {
    return &SinglePad{js.Global.Get("Phaser").Get("SinglePad").New(game, padParent)}
}
// NewSinglePadI A single Phaser Gamepad
func NewSinglePadI(args ...interface{}) *SinglePad {
    return &SinglePad{js.Global.Get("Phaser").Get("SinglePad").New(args)}
}



// SinglePad Binding conversion method to SinglePad point 
func ToSinglePad(jsStruct interface{}) *SinglePad {
    if object, ok := jsStruct.(*js.Object); ok {
		return &SinglePad{Object: object}
	}
	return nil
}



// Game Local reference to game.
func (self *SinglePad) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *SinglePad) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Index The gamepad index as per browsers data
func (self *SinglePad) Index() int{
    return self.Object.Get("index").Int()
}

// SetIndexA The gamepad index as per browsers data
func (self *SinglePad) SetIndexA(member int) {
    self.Object.Set("index", member)
}

// Connected Whether or not this particular gamepad is connected or not.
func (self *SinglePad) Connected() bool{
    return self.Object.Get("connected").Bool()
}

// SetConnectedA Whether or not this particular gamepad is connected or not.
func (self *SinglePad) SetConnectedA(member bool) {
    self.Object.Set("connected", member)
}

// CallbackContext The context under which the callbacks are run.
func (self *SinglePad) CallbackContext() interface{}{
    return self.Object.Get("callbackContext")
}

// SetCallbackContextA The context under which the callbacks are run.
func (self *SinglePad) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// OnConnectCallback This callback is invoked every time this gamepad is connected
func (self *SinglePad) OnConnectCallback() interface{}{
    return self.Object.Get("onConnectCallback")
}

// SetOnConnectCallbackA This callback is invoked every time this gamepad is connected
func (self *SinglePad) SetOnConnectCallbackA(member interface{}) {
    self.Object.Set("onConnectCallback", member)
}

// OnDisconnectCallback This callback is invoked every time this gamepad is disconnected
func (self *SinglePad) OnDisconnectCallback() interface{}{
    return self.Object.Get("onDisconnectCallback")
}

// SetOnDisconnectCallbackA This callback is invoked every time this gamepad is disconnected
func (self *SinglePad) SetOnDisconnectCallbackA(member interface{}) {
    self.Object.Set("onDisconnectCallback", member)
}

// OnDownCallback This callback is invoked every time a button is pressed down.
func (self *SinglePad) OnDownCallback() interface{}{
    return self.Object.Get("onDownCallback")
}

// SetOnDownCallbackA This callback is invoked every time a button is pressed down.
func (self *SinglePad) SetOnDownCallbackA(member interface{}) {
    self.Object.Set("onDownCallback", member)
}

// OnUpCallback This callback is invoked every time a gamepad button is released.
func (self *SinglePad) OnUpCallback() interface{}{
    return self.Object.Get("onUpCallback")
}

// SetOnUpCallbackA This callback is invoked every time a gamepad button is released.
func (self *SinglePad) SetOnUpCallbackA(member interface{}) {
    self.Object.Set("onUpCallback", member)
}

// OnAxisCallback This callback is invoked every time an axis is changed.
func (self *SinglePad) OnAxisCallback() interface{}{
    return self.Object.Get("onAxisCallback")
}

// SetOnAxisCallbackA This callback is invoked every time an axis is changed.
func (self *SinglePad) SetOnAxisCallbackA(member interface{}) {
    self.Object.Set("onAxisCallback", member)
}

// OnFloatCallback This callback is invoked every time a button is changed to a value where value > 0 and value < 1.
func (self *SinglePad) OnFloatCallback() interface{}{
    return self.Object.Get("onFloatCallback")
}

// SetOnFloatCallbackA This callback is invoked every time a button is changed to a value where value > 0 and value < 1.
func (self *SinglePad) SetOnFloatCallbackA(member interface{}) {
    self.Object.Set("onFloatCallback", member)
}

// DeadZone Dead zone for axis feedback - within this value you won't trigger updates.
func (self *SinglePad) DeadZone() int{
    return self.Object.Get("deadZone").Int()
}

// SetDeadZoneA Dead zone for axis feedback - within this value you won't trigger updates.
func (self *SinglePad) SetDeadZoneA(member int) {
    self.Object.Set("deadZone", member)
}


// AddCallbacks Add callbacks to this Gamepad to handle connect / disconnect / button down / button up / axis change / float value buttons.
func (self *SinglePad) AddCallbacks(context interface{}, callbacks interface{}) {
    self.Object.Call("addCallbacks", context, callbacks)
}

// AddCallbacksI Add callbacks to this Gamepad to handle connect / disconnect / button down / button up / axis change / float value buttons.
func (self *SinglePad) AddCallbacksI(args ...interface{}) {
    self.Object.Call("addCallbacks", args)
}

// GetButton Gets a DeviceButton object from this controller to be stored and referenced locally.
// The DeviceButton object can then be polled, have events attached to it, etc.
func (self *SinglePad) GetButton(buttonCode int) *DeviceButton{
    return &DeviceButton{self.Object.Call("getButton", buttonCode)}
}

// GetButtonI Gets a DeviceButton object from this controller to be stored and referenced locally.
// The DeviceButton object can then be polled, have events attached to it, etc.
func (self *SinglePad) GetButtonI(args ...interface{}) *DeviceButton{
    return &DeviceButton{self.Object.Call("getButton", args)}
}

// PollStatus Main update function called by Phaser.Gamepad.
func (self *SinglePad) PollStatus() {
    self.Object.Call("pollStatus")
}

// PollStatusI Main update function called by Phaser.Gamepad.
func (self *SinglePad) PollStatusI(args ...interface{}) {
    self.Object.Call("pollStatus", args)
}

// Connect Gamepad connect function, should be called by Phaser.Gamepad.
func (self *SinglePad) Connect(rawPad interface{}) {
    self.Object.Call("connect", rawPad)
}

// ConnectI Gamepad connect function, should be called by Phaser.Gamepad.
func (self *SinglePad) ConnectI(args ...interface{}) {
    self.Object.Call("connect", args)
}

// Disconnect Gamepad disconnect function, should be called by Phaser.Gamepad.
func (self *SinglePad) Disconnect() {
    self.Object.Call("disconnect")
}

// DisconnectI Gamepad disconnect function, should be called by Phaser.Gamepad.
func (self *SinglePad) DisconnectI(args ...interface{}) {
    self.Object.Call("disconnect", args)
}

// Destroy Destroys this object and associated callback references.
func (self *SinglePad) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this object and associated callback references.
func (self *SinglePad) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// ProcessAxisChange Handles changes in axis.
func (self *SinglePad) ProcessAxisChange(axisState interface{}) {
    self.Object.Call("processAxisChange", axisState)
}

// ProcessAxisChangeI Handles changes in axis.
func (self *SinglePad) ProcessAxisChangeI(args ...interface{}) {
    self.Object.Call("processAxisChange", args)
}

// ProcessButtonDown Handles button down press.
func (self *SinglePad) ProcessButtonDown(buttonCode int, value interface{}) {
    self.Object.Call("processButtonDown", buttonCode, value)
}

// ProcessButtonDownI Handles button down press.
func (self *SinglePad) ProcessButtonDownI(args ...interface{}) {
    self.Object.Call("processButtonDown", args)
}

// ProcessButtonUp Handles button release.
func (self *SinglePad) ProcessButtonUp(buttonCode int, value interface{}) {
    self.Object.Call("processButtonUp", buttonCode, value)
}

// ProcessButtonUpI Handles button release.
func (self *SinglePad) ProcessButtonUpI(args ...interface{}) {
    self.Object.Call("processButtonUp", args)
}

// ProcessButtonFloat Handles buttons with floating values (like analog buttons that acts almost like an axis but still registers like a button)
func (self *SinglePad) ProcessButtonFloat(buttonCode int, value interface{}) {
    self.Object.Call("processButtonFloat", buttonCode, value)
}

// ProcessButtonFloatI Handles buttons with floating values (like analog buttons that acts almost like an axis but still registers like a button)
func (self *SinglePad) ProcessButtonFloatI(args ...interface{}) {
    self.Object.Call("processButtonFloat", args)
}

// Axis Returns value of requested axis.
func (self *SinglePad) Axis(axisCode int) int{
    return self.Object.Call("axis", axisCode).Int()
}

// AxisI Returns value of requested axis.
func (self *SinglePad) AxisI(args ...interface{}) int{
    return self.Object.Call("axis", args).Int()
}

// IsDown Returns true if the button is pressed down.
func (self *SinglePad) IsDown(buttonCode int) bool{
    return self.Object.Call("isDown", buttonCode).Bool()
}

// IsDownI Returns true if the button is pressed down.
func (self *SinglePad) IsDownI(args ...interface{}) bool{
    return self.Object.Call("isDown", args).Bool()
}

// IsUp Returns true if the button is not currently pressed.
func (self *SinglePad) IsUp(buttonCode int) bool{
    return self.Object.Call("isUp", buttonCode).Bool()
}

// IsUpI Returns true if the button is not currently pressed.
func (self *SinglePad) IsUpI(args ...interface{}) bool{
    return self.Object.Call("isUp", args).Bool()
}

// JustReleased Returns the "just released" state of a button from this gamepad. Just released is considered as being true if the button was released within the duration given (default 250ms).
func (self *SinglePad) JustReleased(buttonCode int) bool{
    return self.Object.Call("justReleased", buttonCode).Bool()
}

// JustReleased1O Returns the "just released" state of a button from this gamepad. Just released is considered as being true if the button was released within the duration given (default 250ms).
func (self *SinglePad) JustReleased1O(buttonCode int, duration int) bool{
    return self.Object.Call("justReleased", buttonCode, duration).Bool()
}

// JustReleasedI Returns the "just released" state of a button from this gamepad. Just released is considered as being true if the button was released within the duration given (default 250ms).
func (self *SinglePad) JustReleasedI(args ...interface{}) bool{
    return self.Object.Call("justReleased", args).Bool()
}

// JustPressed Returns the "just pressed" state of a button from this gamepad. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *SinglePad) JustPressed(buttonCode int) bool{
    return self.Object.Call("justPressed", buttonCode).Bool()
}

// JustPressed1O Returns the "just pressed" state of a button from this gamepad. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *SinglePad) JustPressed1O(buttonCode int, duration int) bool{
    return self.Object.Call("justPressed", buttonCode, duration).Bool()
}

// JustPressedI Returns the "just pressed" state of a button from this gamepad. Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *SinglePad) JustPressedI(args ...interface{}) bool{
    return self.Object.Call("justPressed", args).Bool()
}

// ButtonValue Returns the value of a gamepad button. Intended mainly for cases when you have floating button values, for example
// analog trigger buttons on the XBOX 360 controller.
func (self *SinglePad) ButtonValue(buttonCode int) int{
    return self.Object.Call("buttonValue", buttonCode).Int()
}

// ButtonValueI Returns the value of a gamepad button. Intended mainly for cases when you have floating button values, for example
// analog trigger buttons on the XBOX 360 controller.
func (self *SinglePad) ButtonValueI(args ...interface{}) int{
    return self.Object.Call("buttonValue", args).Int()
}

// Reset Reset all buttons/axes of this gamepad.
func (self *SinglePad) Reset() {
    self.Object.Call("reset")
}

// ResetI Reset all buttons/axes of this gamepad.
func (self *SinglePad) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

