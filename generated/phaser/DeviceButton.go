// Automatic generation for Phaser.DeviceButton
// generated file DeviceButton.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// DeviceButtons belong to both `Phaser.Pointer` and `Phaser.SinglePad` (Gamepad) instances.
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


// A reference to the Pointer or Gamepad that owns this button.
func (self *DeviceButton) GetParent() interface{}{
    return self.Get("parent")
}

// A reference to the Pointer or Gamepad that owns this button.
func (self *DeviceButton) SetParent(member interface{}) {
    self.Set("parent", member)
}

// A reference to the currently running game.
func (self *DeviceButton) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *DeviceButton) SetGame(member *Game) {
    self.Set("game", member)
}

// The DOM event that caused the change in button state.
func (self *DeviceButton) GetEvent() interface{}{
    return self.Get("event")
}

// The DOM event that caused the change in button state.
func (self *DeviceButton) SetEvent(member interface{}) {
    self.Set("event", member)
}

// The "down" state of the button.
func (self *DeviceButton) GetIsDown() bool{
    return self.Get("isDown").Bool()
}

// The "down" state of the button.
func (self *DeviceButton) SetIsDown(member bool) {
    self.Set("isDown", member)
}

// The "up" state of the button.
func (self *DeviceButton) GetIsUp() bool{
    return self.Get("isUp").Bool()
}

// The "up" state of the button.
func (self *DeviceButton) SetIsUp(member bool) {
    self.Set("isUp", member)
}

// The timestamp when the button was last pressed down.
func (self *DeviceButton) GetTimeDown() float64{
    return self.Get("timeDown").Float()
}

// The timestamp when the button was last pressed down.
func (self *DeviceButton) SetTimeDown(member float64) {
    self.Set("timeDown", member)
}

// The timestamp when the button was last released.
func (self *DeviceButton) GetTimeUp() float64{
    return self.Get("timeUp").Float()
}

// The timestamp when the button was last released.
func (self *DeviceButton) SetTimeUp(member float64) {
    self.Set("timeUp", member)
}

// Gamepad only.
// If a button is held down this holds down the number of times the button has 'repeated'.
func (self *DeviceButton) GetRepeats() float64{
    return self.Get("repeats").Float()
}

// Gamepad only.
// If a button is held down this holds down the number of times the button has 'repeated'.
func (self *DeviceButton) SetRepeats(member float64) {
    self.Set("repeats", member)
}

// True if the alt key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) GetAltKey() bool{
    return self.Get("altKey").Bool()
}

// True if the alt key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) SetAltKey(member bool) {
    self.Set("altKey", member)
}

// True if the shift key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) GetShiftKey() bool{
    return self.Get("shiftKey").Bool()
}

// True if the shift key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) SetShiftKey(member bool) {
    self.Set("shiftKey", member)
}

// True if the control key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) GetCtrlKey() bool{
    return self.Get("ctrlKey").Bool()
}

// True if the control key was held down when this button was last pressed or released.
// Not supported on Gamepads.
func (self *DeviceButton) SetCtrlKey(member bool) {
    self.Set("ctrlKey", member)
}

// Button value. Mainly useful for checking analog buttons (like shoulder triggers) on Gamepads.
func (self *DeviceButton) GetValue() float64{
    return self.Get("value").Float()
}

// Button value. Mainly useful for checking analog buttons (like shoulder triggers) on Gamepads.
func (self *DeviceButton) SetValue(member float64) {
    self.Set("value", member)
}

// The buttoncode of this button if a Gamepad, or the DOM button event value if a Pointer.
func (self *DeviceButton) GetButtonCode() float64{
    return self.Get("buttonCode").Float()
}

// The buttoncode of this button if a Gamepad, or the DOM button event value if a Pointer.
func (self *DeviceButton) SetButtonCode(member float64) {
    self.Set("buttonCode", member)
}

// This Signal is dispatched every time this DeviceButton is pressed down.
// It is only dispatched once (until the button is released again).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) GetOnDown() *Signal{
    return &Signal{self.Get("onDown")}
}

// This Signal is dispatched every time this DeviceButton is pressed down.
// It is only dispatched once (until the button is released again).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) SetOnDown(member *Signal) {
    self.Set("onDown", member)
}

// This Signal is dispatched every time this DeviceButton is released from a down state.
// It is only dispatched once (until the button is pressed again).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) GetOnUp() *Signal{
    return &Signal{self.Get("onUp")}
}

// This Signal is dispatched every time this DeviceButton is released from a down state.
// It is only dispatched once (until the button is pressed again).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) SetOnUp(member *Signal) {
    self.Set("onUp", member)
}

// Gamepad only.
// This Signal is dispatched every time this DeviceButton changes floating value (between, but not exactly, 0 and 1).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) GetOnFloat() *Signal{
    return &Signal{self.Get("onFloat")}
}

// Gamepad only.
// This Signal is dispatched every time this DeviceButton changes floating value (between, but not exactly, 0 and 1).
// When dispatched it sends 2 arguments: A reference to this DeviceButton and the value of the button.
func (self *DeviceButton) SetOnFloat(member *Signal) {
    self.Set("onFloat", member)
}

// How long the button has been held down for in milliseconds.
// If not currently down it returns -1.
func (self *DeviceButton) GetDuration() float64{
    return self.Get("duration").Float()
}

// How long the button has been held down for in milliseconds.
// If not currently down it returns -1.
func (self *DeviceButton) SetDuration(member float64) {
    self.Set("duration", member)
}



// Called automatically by Phaser.Pointer and Phaser.SinglePad.
// Handles the button down state.
func (self *DeviceButton) StartI(args ...interface{}) {
    self.Call("start", args)
}

// Called automatically by Phaser.Pointer and Phaser.SinglePad.
// Handles the button up state.
func (self *DeviceButton) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// Called automatically by Phaser.SinglePad.
func (self *DeviceButton) PadFloatI(args ...interface{}) {
    self.Call("padFloat", args)
}

// Returns the "just pressed" state of this button.
// Just pressed is considered true if the button was pressed down within the duration given (default 250ms).
func (self *DeviceButton) JustPressedI(args ...interface{}) bool{
    return self.Call("justPressed", args).Bool()
}

// Returns the "just released" state of this button.
// Just released is considered as being true if the button was released within the duration given (default 250ms).
func (self *DeviceButton) JustReleasedI(args ...interface{}) bool{
    return self.Call("justReleased", args).Bool()
}

// Resets this DeviceButton, changing it to an isUp state and resetting the duration and repeats counters.
func (self *DeviceButton) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Destroys this DeviceButton, this disposes of the onDown, onUp and onFloat signals 
// and clears the parent and game references.
func (self *DeviceButton) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
