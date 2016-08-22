// Automatic generation for Phaser.Mouse
// generated file Mouse.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Mouse class is responsible for handling all aspects of mouse interaction with the browser.
// 
// It captures and processes mouse events that happen on the game canvas object.
// It also adds a single `mouseup` listener to `window` which is used to capture the mouse being released
// when not over the game.
// 
// You should not normally access this class directly, but instead use a Phaser.Pointer object
// which normalises all game input for you, including accurate button handling.
type Mouse struct {
    *js.Object
}


// A reference to the currently running game.
func (self *Mouse) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *Mouse) SetGame(member *Game) {
    self.Set("game", member)
}

// A reference to the Phaser Input Manager.
func (self *Mouse) GetInput() *Input{
    return &Input{self.Get("input")}
}

// A reference to the Phaser Input Manager.
func (self *Mouse) SetInput(member *Input) {
    self.Set("input", member)
}

// The context under which callbacks are called.
func (self *Mouse) GetCallbackContext() interface{}{
    return self.Get("callbackContext")
}

// The context under which callbacks are called.
func (self *Mouse) SetCallbackContext(member interface{}) {
    self.Set("callbackContext", member)
}

// A callback that can be fired when the mouse is pressed down.
func (self *Mouse) SetMouseDownCallback(member func(...interface{})) {
    self.Set("mouseDownCallback", member)
}

// A callback that can be fired when the mouse is released from a pressed down state.
func (self *Mouse) SetMouseUpCallback(member func(...interface{})) {
    self.Set("mouseUpCallback", member)
}

// A callback that can be fired when the mouse is no longer over the game canvas.
func (self *Mouse) SetMouseOutCallback(member func(...interface{})) {
    self.Set("mouseOutCallback", member)
}

// A callback that can be fired when the mouse enters the game canvas (usually after a mouseout).
func (self *Mouse) SetMouseOverCallback(member func(...interface{})) {
    self.Set("mouseOverCallback", member)
}

// A callback that can be fired when the mousewheel is used.
func (self *Mouse) SetMouseWheelCallback(member func(...interface{})) {
    self.Set("mouseWheelCallback", member)
}

// If true the DOM mouse events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *Mouse) GetCapture() bool{
    return self.Get("capture").Bool()
}

// If true the DOM mouse events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *Mouse) SetCapture(member bool) {
    self.Set("capture", member)
}

// This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *Mouse) GetButton() int{
    return self.Get("button").Int()
}

// This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *Mouse) SetButton(member int) {
    self.Set("button", member)
}

// The direction of the _last_ mousewheel usage 1 for up -1 for down.
func (self *Mouse) GetWheelDelta() int{
    return self.Get("wheelDelta").Int()
}

// The direction of the _last_ mousewheel usage 1 for up -1 for down.
func (self *Mouse) SetWheelDelta(member int) {
    self.Set("wheelDelta", member)
}

// Mouse input will only be processed if enabled.
func (self *Mouse) GetEnabled() bool{
    return self.Get("enabled").Bool()
}

// Mouse input will only be processed if enabled.
func (self *Mouse) SetEnabled(member bool) {
    self.Set("enabled", member)
}

// If the mouse has been Pointer Locked successfully this will be set to true.
func (self *Mouse) GetLocked() bool{
    return self.Get("locked").Bool()
}

// If the mouse has been Pointer Locked successfully this will be set to true.
func (self *Mouse) SetLocked(member bool) {
    self.Set("locked", member)
}

// If true Pointer.stop will be called if the mouse leaves the game canvas.
func (self *Mouse) GetStopOnGameOut() bool{
    return self.Get("stopOnGameOut").Bool()
}

// If true Pointer.stop will be called if the mouse leaves the game canvas.
func (self *Mouse) SetStopOnGameOut(member bool) {
    self.Set("stopOnGameOut", member)
}

// This event is dispatched when the browser enters or leaves pointer lock state.
func (self *Mouse) GetPointerLock() *Signal{
    return &Signal{self.Get("pointerLock")}
}

// This event is dispatched when the browser enters or leaves pointer lock state.
func (self *Mouse) SetPointerLock(member *Signal) {
    self.Set("pointerLock", member)
}

// The browser mouse DOM event. Will be null if no mouse event has ever been received.
// Access this property only inside a Mouse event handler and do not keep references to it.
func (self *Mouse) GetEvent() interface{}{
    return self.Get("event")
}

// The browser mouse DOM event. Will be null if no mouse event has ever been received.
// Access this property only inside a Mouse event handler and do not keep references to it.
func (self *Mouse) SetEvent(member interface{}) {
    self.Set("event", member)
}

// 
func (self *Mouse) GetNO_BUTTON() int{
    return self.Get("NO_BUTTON").Int()
}

// 
func (self *Mouse) SetNO_BUTTON(member int) {
    self.Set("NO_BUTTON", member)
}

// 
func (self *Mouse) GetLEFT_BUTTON() int{
    return self.Get("LEFT_BUTTON").Int()
}

// 
func (self *Mouse) SetLEFT_BUTTON(member int) {
    self.Set("LEFT_BUTTON", member)
}

// 
func (self *Mouse) GetMIDDLE_BUTTON() int{
    return self.Get("MIDDLE_BUTTON").Int()
}

// 
func (self *Mouse) SetMIDDLE_BUTTON(member int) {
    self.Set("MIDDLE_BUTTON", member)
}

// 
func (self *Mouse) GetRIGHT_BUTTON() int{
    return self.Get("RIGHT_BUTTON").Int()
}

// 
func (self *Mouse) SetRIGHT_BUTTON(member int) {
    self.Set("RIGHT_BUTTON", member)
}

// 
func (self *Mouse) GetBACK_BUTTON() int{
    return self.Get("BACK_BUTTON").Int()
}

// 
func (self *Mouse) SetBACK_BUTTON(member int) {
    self.Set("BACK_BUTTON", member)
}

// 
func (self *Mouse) GetFORWARD_BUTTON() int{
    return self.Get("FORWARD_BUTTON").Int()
}

// 
func (self *Mouse) SetFORWARD_BUTTON(member int) {
    self.Set("FORWARD_BUTTON", member)
}

// 
func (self *Mouse) GetWHEEL_UP() int{
    return self.Get("WHEEL_UP").Int()
}

// 
func (self *Mouse) SetWHEEL_UP(member int) {
    self.Set("WHEEL_UP", member)
}

// 
func (self *Mouse) GetWHEEL_DOWN() int{
    return self.Get("WHEEL_DOWN").Int()
}

// 
func (self *Mouse) SetWHEEL_DOWN(member int) {
    self.Set("WHEEL_DOWN", member)
}



// Starts the event listeners running.
func (self *Mouse) StartI(args ...interface{}) {
    self.Call("start", args)
}

// The internal method that handles the mouse down event from the browser.
func (self *Mouse) OnMouseDownI(args ...interface{}) {
    self.Call("onMouseDown", args)
}

// The internal method that handles the mouse move event from the browser.
func (self *Mouse) OnMouseMoveI(args ...interface{}) {
    self.Call("onMouseMove", args)
}

// The internal method that handles the mouse up event from the browser.
func (self *Mouse) OnMouseUpI(args ...interface{}) {
    self.Call("onMouseUp", args)
}

// The internal method that handles the mouse up event from the window.
func (self *Mouse) OnMouseUpGlobalI(args ...interface{}) {
    self.Call("onMouseUpGlobal", args)
}

// The internal method that handles the mouse out event from the window.
func (self *Mouse) OnMouseOutGlobalI(args ...interface{}) {
    self.Call("onMouseOutGlobal", args)
}

// The internal method that handles the mouse out event from the browser.
func (self *Mouse) OnMouseOutI(args ...interface{}) {
    self.Call("onMouseOut", args)
}

// The internal method that handles the mouse over event from the browser.
func (self *Mouse) OnMouseOverI(args ...interface{}) {
    self.Call("onMouseOver", args)
}

// The internal method that handles the mouse wheel event from the browser.
func (self *Mouse) OnMouseWheelI(args ...interface{}) {
    self.Call("onMouseWheel", args)
}

// If the browser supports it you can request that the pointer be locked to the browser window.
// This is classically known as 'FPS controls', where the pointer can't leave the browser until the user presses an exit key.
// If the browser successfully enters a locked state the event Phaser.Mouse.pointerLock will be dispatched and the first parameter will be 'true'.
func (self *Mouse) RequestPointerLockI(args ...interface{}) {
    self.Call("requestPointerLock", args)
}

// Internal pointerLockChange handler.
func (self *Mouse) PointerLockChangeI(args ...interface{}) {
    self.Call("pointerLockChange", args)
}

// Internal release pointer lock handler.
func (self *Mouse) ReleasePointerLockI(args ...interface{}) {
    self.Call("releasePointerLock", args)
}

// Stop the event listeners.
func (self *Mouse) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// A purely internal event support class to proxy 'wheelscroll' and 'DOMMouseWheel'
// events to 'wheel'-like events.
// 
// See https://developer.mozilla.org/en-US/docs/Web/Events/mousewheel for choosing a scale and delta mode.
func (self *Mouse) WheelEventProxyI(args ...interface{}) {
    self.Call("WheelEventProxy", args)
}
