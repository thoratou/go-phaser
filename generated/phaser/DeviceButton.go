// Package phaser Automatic generation for Phaser.DeviceButton
// generated file DeviceButton.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// DeviceButton DeviceButtons belong to both `Phaser.Pointer` and `Phaser.SinglePad` (Gamepad) instances.
// 
// For Pointers they represent the various buttons that can exist on mice and pens, such as the left button, right button,
// middle button and advanced buttons like back and forward.
// 
// Access them via `Pointer.leftbutton`, `Pointer.rightButton` and so on.
// 
// On Gamepads they represent all buttons on the pad: from shoulder buttons to action buttons.
// 
// At the time of writing this there are device limitations you should be aware of:
// 
// - On Windows, if you install a mouse driver, and its utility software allows you to customize button actions 
//   (e.g., IntelliPoint and SetPoint), the middle (wheel) button, the 4th button, and the 5th button might not be set, 
//   even when they are pressed.
// - On Linux (GTK), the 4th button and the 5th button are not supported.
// - On Mac OS X 10.5 there is no platform API for implementing any advanced buttons.
type DeviceButton struct {
    *js.Object
}

// NewDeviceButton DeviceButtons belong to both `Phaser.Pointer` and `Phaser.SinglePad` (Gamepad) instances.
// 
// For Pointers they represent the various buttons that can exist on mice and pens, such as the left button, right button,
// middle button and advanced buttons like back and forward.
// 
// Access them via `Pointer.leftbutton`, `Pointer.rightButton` and so on.
// 
// On Gamepads they represent all buttons on the pad: from shoulder buttons to action buttons.
// 
// At the time of writing this there are device limitations you should be aware of:
// 
// - On Windows, if you install a mouse driver, and its utility software allows you to customize button actions 
//   (e.g., IntelliPoint and SetPoint), the middle (wheel) button, the 4th button, and the 5th button might not be set, 
//   even when they are pressed.
// - On Linux (GTK), the 4th button and the 5th button are not supported.
// - On Mac OS X 10.5 there is no platform API for implementing any advanced buttons.
func NewDeviceButton(parent interface{}, buttonCode int) *DeviceButton {
    return &DeviceButton{js.Global.Get("Phaser").Get("DeviceButton").New(parent, buttonCode)}
}
// NewDeviceButtonI DeviceButtons belong to both `Phaser.Pointer` and `Phaser.SinglePad` (Gamepad) instances.
// 
// For Pointers they represent the various buttons that can exist on mice and pens, such as the left button, right button,
// middle button and advanced buttons like back and forward.
// 
// Access them via `Pointer.leftbutton`, `Pointer.rightButton` and so on.
// 
// On Gamepads they represent all buttons on the pad: from shoulder buttons to action buttons.
// 
// At the time of writing this there are device limitations you should be aware of:
// 
// - On Windows, if you install a mouse driver, and its utility software allows you to customize button actions 
//   (e.g., IntelliPoint and SetPoint), the middle (wheel) button, the 4th button, and the 5th button might not be set, 
//   even when they are pressed.
// - On Linux (GTK), the 4th button and the 5th button are not supported.
// - On Mac OS X 10.5 there is no platform API for implementing any advanced buttons.
func NewDeviceButtonI(args ...interface{}) *DeviceButton {
    return &DeviceButton{js.Global.Get("Phaser").Get("DeviceButton").New(args)}
}



// Parent A reference to the Pointer or Gamepad that owns this button.
func (self *DeviceButton) Parent() interface{}{
    return self.Object.Get("parent")
}

// SetParentA A reference to the Pointer or Gamepad that owns this button.
func (self *DeviceButton) SetParentA(member interface{}) {
    self.Object.Set("parent", member)
}

// Game A reference to the currently running game.
func (self *DeviceButton) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *DeviceButton) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Event The DOM event that caused the change in button state.
func (self *DeviceButton) Event() interface{}{
    return self.Object.Get("event")
}

// SetEventA The DOM event that caused the change in button state.
func (self *DeviceButton) SetEventA(member interface{}) {
    self.Object.Set("event", member)
}

// IsDown The "down" state of the button.
func (self *DeviceButton) IsDown() bool{
    return self.Object.Get("isDown").Bool()
}

// SetIsDownA The "down" state of the button.
func (self *DeviceButton) SetIsDownA(member bool) {
    self.Object.Set("isDown", member)
}

// IsUp The "up" state of the button.
func (self *DeviceButton) IsUp() bool{
    return self.Object.Get("isUp").Bool()
}

// SetIsUpA The "up" state of the button.
func (self *DeviceButton) SetIsUpA(member bool) {
    self.Object.Set("isUp", member)
}

// TimeDown The timestamp when the button was last pressed down.
func (self *DeviceButton) TimeDown() int{
    return self.Object.Get("timeDown").Int()
}

// SetTimeDownA The timestamp when the button was last pressed down.
func (self *DeviceButton) SetTimeDownA(member int) {
    self.Object.Set("timeDown", member)
}

// TimeUp The timestamp when the button was last released.
func (self *DeviceButton) TimeUp() int{
    return self.Object.Get("timeUp").Int()
}

// SetTimeUpA The timestamp when the button was last released.
func (self *DeviceButton) SetTimeUpA(member int) {
    self.Object.Set("timeUp", member)
}

// Repeats Gamepad only.
// If a button is held down this holds down the number of times the button has 'repeated'.
func (self *DeviceButton) Repeats() int{
    return self.Object.Get("repeats").Int()
}

// SetRepeatsA Gamepad only.
// If a button is held down this holds down the number of times the button has 'repeated'.
func (self *DeviceButton) SetRepeatsA(member int) {
    self.Object.Set("repeats", member)
}

// AltKey True if the alt key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) AltKey() bool{
    return self.Object.Get("altKey").Bool()
}

// SetAltKeyA True if the alt key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) SetAltKeyA(member bool) {
    self.Object.Set("altKey", member)
}

// ShiftKey True if the shift key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) ShiftKey() bool{
    return self.Object.Get("shiftKey").Bool()
}

// SetShiftKeyA True if the shift key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) SetShiftKeyA(member bool) {
    self.Object.Set("shiftKey", member)
}

// CtrlKey True if the control key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) CtrlKey() bool{
    return self.Object.Get("ctrlKey").Bool()
}

// SetCtrlKeyA True if the control key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) SetCtrlKeyA(member bool) {
    self.Object.Set("ctrlKey", member)
}

// Value Button value. Mainly useful for checking analog buttons (like shoulder triggers) on Gamepads.
func (self *DeviceButton) Value() int{
    return self.Object.Get("value").Int()
}

// SetValueA Button value. Mainly useful for checking analog buttons (like shoulder triggers) on Gamepads.
func (self *DeviceButton) SetValueA(member int) {
    self.Object.Set("value", member)
}

// ButtonCode The buttoncode of this button if a Gamepad, or the DOM button event value if a Pointer.
func (self *DeviceButton) ButtonCode() int{
    return self.Object.Get("buttonCode").Int()
}

// SetButtonCodeA The buttoncode of this button if a Gamepad, or the DOM button event value if a Pointer.
func (self *DeviceButton) SetButtonCodeA(member int) {
    self.Object.Set("buttonCode", member)
}

// OnDown This Signal is dispatched every time this DeviceButton is pressed down.
// It is only dispatched once (until the button is released again).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) OnDown() *Signal{
    return &Signal{self.Object.Get("onDown")}
}

// SetOnDownA This Signal is dispatched every time this DeviceButton is pressed down.
// It is only dispatched once (until the button is released again).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) SetOnDownA(member *Signal) {
    self.Object.Set("onDown", member)
}

// OnUp This Signal is dispatched every time this DeviceButton is released from a down state.
// It is only dispatched once (until the button is pressed again).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) OnUp() *Signal{
    return &Signal{self.Object.Get("onUp")}
}

// SetOnUpA This Signal is dispatched every time this DeviceButton is released from a down state.
// It is only dispatched once (until the button is pressed again).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) SetOnUpA(member *Signal) {
    self.Object.Set("onUp", member)
}

// OnFloat Gamepad only.
// This Signal is dispatched every time this DeviceButton changes floating value (between, but not exactly, 0 and 1).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) OnFloat() *Signal{
    return &Signal{self.Object.Get("onFloat")}
}

// SetOnFloatA Gamepad only.
// This Signal is dispatched every time this DeviceButton changes floating value (between, but not exactly, 0 and 1).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) SetOnFloatA(member *Signal) {
    self.Object.Set("onFloat", member)
}

// Duration How long the button has been held down for in milliseconds.
// If not currently down it returns -1.
func (self *DeviceButton) Duration() int{
    return self.Object.Get("duration").Int()
}

// SetDurationA How long the button has been held down for in milliseconds.
// If not currently down it returns -1.
func (self *DeviceButton) SetDurationA(member int) {
    self.Object.Set("duration", member)
}


// Start Called automatically by Phaser.Pointer and Phaser.SinglePad.
// Handles the button down state.
func (self *DeviceButton) Start() {
    self.Object.Call("start")
}

// Start1O Called automatically by Phaser.Pointer and Phaser.SinglePad.
// Handles the button down state.
func (self *DeviceButton) Start1O(event interface{}) {
    self.Object.Call("start", event)
}

// Start2O Called automatically by Phaser.Pointer and Phaser.SinglePad.
// Handles the button down state.
func (self *DeviceButton) Start2O(event interface{}, value int) {
    self.Object.Call("start", event, value)
}

// StartI Called automatically by Phaser.Pointer and Phaser.SinglePad.
// Handles the button down state.
func (self *DeviceButton) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Stop Called automatically by Phaser.Pointer and Phaser.SinglePad.
// Handles the button up state.
func (self *DeviceButton) Stop() {
    self.Object.Call("stop")
}

// Stop1O Called automatically by Phaser.Pointer and Phaser.SinglePad.
// Handles the button up state.
func (self *DeviceButton) Stop1O(event interface{}) {
    self.Object.Call("stop", event)
}

// Stop2O Called automatically by Phaser.Pointer and Phaser.SinglePad.
// Handles the button up state.
func (self *DeviceButton) Stop2O(event interface{}, value int) {
    self.Object.Call("stop", event, value)
}

// StopI Called automatically by Phaser.Pointer and Phaser.SinglePad.
// Handles the button up state.
func (self *DeviceButton) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// PadFloat Called automatically by Phaser.SinglePad.
func (self *DeviceButton) PadFloat(value int) {
    self.Object.Call("padFloat", value)
}

// PadFloatI Called automatically by Phaser.SinglePad.
func (self *DeviceButton) PadFloatI(args ...interface{}) {
    self.Object.Call("padFloat", args)
}

// JustPressed Returns the "just pressed" state of this button.
// Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *DeviceButton) JustPressed() bool{
    return self.Object.Call("justPressed").Bool()
}

// JustPressed1O Returns the "just pressed" state of this button.
// Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *DeviceButton) JustPressed1O(duration int) bool{
    return self.Object.Call("justPressed", duration).Bool()
}

// JustPressedI Returns the "just pressed" state of this button.
// Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *DeviceButton) JustPressedI(args ...interface{}) bool{
    return self.Object.Call("justPressed", args).Bool()
}

// JustReleased Returns the "just released" state of this button.
// Just released is considered as being true if the button was released within the duration given (default 250ms).
func (self *DeviceButton) JustReleased() bool{
    return self.Object.Call("justReleased").Bool()
}

// JustReleased1O Returns the "just released" state of this button.
// Just released is considered as being true if the button was released within the duration given (default 250ms).
func (self *DeviceButton) JustReleased1O(duration int) bool{
    return self.Object.Call("justReleased", duration).Bool()
}

// JustReleasedI Returns the "just released" state of this button.
// Just released is considered as being true if the button was released within the duration given (default 250ms).
func (self *DeviceButton) JustReleasedI(args ...interface{}) bool{
    return self.Object.Call("justReleased", args).Bool()
}

// Reset Resets this DeviceButton, changing it to an isUp state and resetting the duration and repeats counters.
func (self *DeviceButton) Reset() {
    self.Object.Call("reset")
}

// ResetI Resets this DeviceButton, changing it to an isUp state and resetting the duration and repeats counters.
func (self *DeviceButton) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Destroy Destroys this DeviceButton, this disposes of the onDown, onUp and onFloat signals 
// and clears the parent and game references.
func (self *DeviceButton) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this DeviceButton, this disposes of the onDown, onUp and onFloat signals 
// and clears the parent and game references.
func (self *DeviceButton) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

