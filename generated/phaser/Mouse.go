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
func (self *Mouse) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *Mouse) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// A reference to the Phaser Input Manager.
func (self *Mouse) GetInputA() *Input{
    return &Input{self.Object.Get("input")}
}

// A reference to the Phaser Input Manager.
func (self *Mouse) SetInputA(member *Input) {
    self.Object.Set("input", member)
}

// The context under which callbacks are called.
func (self *Mouse) GetCallbackContextA() interface{}{
    return self.Object.Get("callbackContext")
}

// The context under which callbacks are called.
func (self *Mouse) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// A callback that can be fired when the mouse is pressed down.
func (self *Mouse) SetMouseDownCallbackA(member func(...interface{})) {
    self.Object.Set("mouseDownCallback", member)
}

// A callback that can be fired when the mouse is released from a pressed down state.
func (self *Mouse) SetMouseUpCallbackA(member func(...interface{})) {
    self.Object.Set("mouseUpCallback", member)
}

// A callback that can be fired when the mouse is no longer over the game canvas.
func (self *Mouse) SetMouseOutCallbackA(member func(...interface{})) {
    self.Object.Set("mouseOutCallback", member)
}

// A callback that can be fired when the mouse enters the game canvas (usually after a mouseout).
func (self *Mouse) SetMouseOverCallbackA(member func(...interface{})) {
    self.Object.Set("mouseOverCallback", member)
}

// A callback that can be fired when the mousewheel is used.
func (self *Mouse) SetMouseWheelCallbackA(member func(...interface{})) {
    self.Object.Set("mouseWheelCallback", member)
}

// If true the DOM mouse events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *Mouse) GetCaptureA() bool{
    return self.Object.Get("capture").Bool()
}

// If true the DOM mouse events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *Mouse) SetCaptureA(member bool) {
    self.Object.Set("capture", member)
}

// This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *Mouse) GetButtonA() int{
    return self.Object.Get("button").Int()
}

// This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *Mouse) SetButtonA(member int) {
    self.Object.Set("button", member)
}

// The direction of the _last_ mousewheel usage 1 for up -1 for down.
func (self *Mouse) GetWheelDeltaA() int{
    return self.Object.Get("wheelDelta").Int()
}

// The direction of the _last_ mousewheel usage 1 for up -1 for down.
func (self *Mouse) SetWheelDeltaA(member int) {
    self.Object.Set("wheelDelta", member)
}

// Mouse input will only be processed if enabled.
func (self *Mouse) GetEnabledA() bool{
    return self.Object.Get("enabled").Bool()
}

// Mouse input will only be processed if enabled.
func (self *Mouse) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}

// If the mouse has been Pointer Locked successfully this will be set to true.
func (self *Mouse) GetLockedA() bool{
    return self.Object.Get("locked").Bool()
}

// If the mouse has been Pointer Locked successfully this will be set to true.
func (self *Mouse) SetLockedA(member bool) {
    self.Object.Set("locked", member)
}

// If true Pointer.stop will be called if the mouse leaves the game canvas.
func (self *Mouse) GetStopOnGameOutA() bool{
    return self.Object.Get("stopOnGameOut").Bool()
}

// If true Pointer.stop will be called if the mouse leaves the game canvas.
func (self *Mouse) SetStopOnGameOutA(member bool) {
    self.Object.Set("stopOnGameOut", member)
}

// This event is dispatched when the browser enters or leaves pointer lock state.
func (self *Mouse) GetPointerLockA() *Signal{
    return &Signal{self.Object.Get("pointerLock")}
}

// This event is dispatched when the browser enters or leaves pointer lock state.
func (self *Mouse) SetPointerLockA(member *Signal) {
    self.Object.Set("pointerLock", member)
}

// The browser mouse DOM event. Will be null if no mouse event has ever been received.
// Access this property only inside a Mouse event handler and do not keep references to it.
func (self *Mouse) GetEventA() interface{}{
    return self.Object.Get("event")
}

// The browser mouse DOM event. Will be null if no mouse event has ever been received.
// Access this property only inside a Mouse event handler and do not keep references to it.
func (self *Mouse) SetEventA(member interface{}) {
    self.Object.Set("event", member)
}

// 
func (self *Mouse) GetNO_BUTTONA() int{
    return self.Object.Get("NO_BUTTON").Int()
}

// 
func (self *Mouse) SetNO_BUTTONA(member int) {
    self.Object.Set("NO_BUTTON", member)
}

// 
func (self *Mouse) GetLEFT_BUTTONA() int{
    return self.Object.Get("LEFT_BUTTON").Int()
}

// 
func (self *Mouse) SetLEFT_BUTTONA(member int) {
    self.Object.Set("LEFT_BUTTON", member)
}

// 
func (self *Mouse) GetMIDDLE_BUTTONA() int{
    return self.Object.Get("MIDDLE_BUTTON").Int()
}

// 
func (self *Mouse) SetMIDDLE_BUTTONA(member int) {
    self.Object.Set("MIDDLE_BUTTON", member)
}

// 
func (self *Mouse) GetRIGHT_BUTTONA() int{
    return self.Object.Get("RIGHT_BUTTON").Int()
}

// 
func (self *Mouse) SetRIGHT_BUTTONA(member int) {
    self.Object.Set("RIGHT_BUTTON", member)
}

// 
func (self *Mouse) GetBACK_BUTTONA() int{
    return self.Object.Get("BACK_BUTTON").Int()
}

// 
func (self *Mouse) SetBACK_BUTTONA(member int) {
    self.Object.Set("BACK_BUTTON", member)
}

// 
func (self *Mouse) GetFORWARD_BUTTONA() int{
    return self.Object.Get("FORWARD_BUTTON").Int()
}

// 
func (self *Mouse) SetFORWARD_BUTTONA(member int) {
    self.Object.Set("FORWARD_BUTTON", member)
}

// 
func (self *Mouse) GetWHEEL_UPA() int{
    return self.Object.Get("WHEEL_UP").Int()
}

// 
func (self *Mouse) SetWHEEL_UPA(member int) {
    self.Object.Set("WHEEL_UP", member)
}

// 
func (self *Mouse) GetWHEEL_DOWNA() int{
    return self.Object.Get("WHEEL_DOWN").Int()
}

// 
func (self *Mouse) SetWHEEL_DOWNA(member int) {
    self.Object.Set("WHEEL_DOWN", member)
}



// Starts the event listeners running.
func (self *Mouse) Start() {
    self.Object.Call("start")
}

// Starts the event listeners running.
func (self *Mouse) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// The internal method that handles the mouse down event from the browser.
func (self *Mouse) OnMouseDown(event *MouseEvent) {
    self.Object.Call("onMouseDown", event)
}

// The internal method that handles the mouse down event from the browser.
func (self *Mouse) OnMouseDownI(args ...interface{}) {
    self.Object.Call("onMouseDown", args)
}

// The internal method that handles the mouse move event from the browser.
func (self *Mouse) OnMouseMove(event *MouseEvent) {
    self.Object.Call("onMouseMove", event)
}

// The internal method that handles the mouse move event from the browser.
func (self *Mouse) OnMouseMoveI(args ...interface{}) {
    self.Object.Call("onMouseMove", args)
}

// The internal method that handles the mouse up event from the browser.
func (self *Mouse) OnMouseUp(event *MouseEvent) {
    self.Object.Call("onMouseUp", event)
}

// The internal method that handles the mouse up event from the browser.
func (self *Mouse) OnMouseUpI(args ...interface{}) {
    self.Object.Call("onMouseUp", args)
}

// The internal method that handles the mouse up event from the window.
func (self *Mouse) OnMouseUpGlobal(event *MouseEvent) {
    self.Object.Call("onMouseUpGlobal", event)
}

// The internal method that handles the mouse up event from the window.
func (self *Mouse) OnMouseUpGlobalI(args ...interface{}) {
    self.Object.Call("onMouseUpGlobal", args)
}

// The internal method that handles the mouse out event from the window.
func (self *Mouse) OnMouseOutGlobal(event *MouseEvent) {
    self.Object.Call("onMouseOutGlobal", event)
}

// The internal method that handles the mouse out event from the window.
func (self *Mouse) OnMouseOutGlobalI(args ...interface{}) {
    self.Object.Call("onMouseOutGlobal", args)
}

// The internal method that handles the mouse out event from the browser.
func (self *Mouse) OnMouseOut(event *MouseEvent) {
    self.Object.Call("onMouseOut", event)
}

// The internal method that handles the mouse out event from the browser.
func (self *Mouse) OnMouseOutI(args ...interface{}) {
    self.Object.Call("onMouseOut", args)
}

// The internal method that handles the mouse over event from the browser.
func (self *Mouse) OnMouseOver(event *MouseEvent) {
    self.Object.Call("onMouseOver", event)
}

// The internal method that handles the mouse over event from the browser.
func (self *Mouse) OnMouseOverI(args ...interface{}) {
    self.Object.Call("onMouseOver", args)
}

// The internal method that handles the mouse wheel event from the browser.
func (self *Mouse) OnMouseWheel(event *MouseEvent) {
    self.Object.Call("onMouseWheel", event)
}

// The internal method that handles the mouse wheel event from the browser.
func (self *Mouse) OnMouseWheelI(args ...interface{}) {
    self.Object.Call("onMouseWheel", args)
}

// If the browser supports it you can request that the pointer be locked to the browser window.
// This is classically known as 'FPS controls', where the pointer can't leave the browser until the user presses an exit key.
// If the browser successfully enters a locked state the event Phaser.Mouse.pointerLock will be dispatched and the first parameter will be 'true'.
func (self *Mouse) RequestPointerLock() {
    self.Object.Call("requestPointerLock")
}

// If the browser supports it you can request that the pointer be locked to the browser window.
// This is classically known as 'FPS controls', where the pointer can't leave the browser until the user presses an exit key.
// If the browser successfully enters a locked state the event Phaser.Mouse.pointerLock will be dispatched and the first parameter will be 'true'.
func (self *Mouse) RequestPointerLockI(args ...interface{}) {
    self.Object.Call("requestPointerLock", args)
}

// Internal pointerLockChange handler.
func (self *Mouse) PointerLockChange(event *Event) {
    self.Object.Call("pointerLockChange", event)
}

// Internal pointerLockChange handler.
func (self *Mouse) PointerLockChangeI(args ...interface{}) {
    self.Object.Call("pointerLockChange", args)
}

// Internal release pointer lock handler.
func (self *Mouse) ReleasePointerLock() {
    self.Object.Call("releasePointerLock")
}

// Internal release pointer lock handler.
func (self *Mouse) ReleasePointerLockI(args ...interface{}) {
    self.Object.Call("releasePointerLock", args)
}

// Stop the event listeners.
func (self *Mouse) Stop() {
    self.Object.Call("stop")
}

// Stop the event listeners.
func (self *Mouse) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// A purely internal event support class to proxy 'wheelscroll' and 'DOMMouseWheel'
// events to 'wheel'-like events.
// 
// See https://developer.mozilla.org/en-US/docs/Web/Events/mousewheel for choosing a scale and delta mode.
func (self *Mouse) WheelEventProxy(scaleFactor int, deltaMode int) {
    self.Object.Call("WheelEventProxy", scaleFactor, deltaMode)
}

// A purely internal event support class to proxy 'wheelscroll' and 'DOMMouseWheel'
// events to 'wheel'-like events.
// 
// See https://developer.mozilla.org/en-US/docs/Web/Events/mousewheel for choosing a scale and delta mode.
func (self *Mouse) WheelEventProxyI(args ...interface{}) {
    self.Object.Call("WheelEventProxy", args)
}
