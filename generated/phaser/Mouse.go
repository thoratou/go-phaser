// Package phaser Automatic generation for Phaser.Mouse
// generated file Mouse.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Mouse The Mouse class is responsible for handling all aspects of mouse interaction with the browser.
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

// NewMouse The Mouse class is responsible for handling all aspects of mouse interaction with the browser.
// 
// It captures and processes mouse events that happen on the game canvas object.
// It also adds a single `mouseup` listener to `window` which is used to capture the mouse being released
// when not over the game.
// 
// You should not normally access this class directly, but instead use a Phaser.Pointer object
// which normalises all game input for you, including accurate button handling.
func NewMouse(game *Game) *Mouse {
    return &Mouse{js.Global.Get("Phaser").Get("Mouse").New(game)}
}
// NewMouseI The Mouse class is responsible for handling all aspects of mouse interaction with the browser.
// 
// It captures and processes mouse events that happen on the game canvas object.
// It also adds a single `mouseup` listener to `window` which is used to capture the mouse being released
// when not over the game.
// 
// You should not normally access this class directly, but instead use a Phaser.Pointer object
// which normalises all game input for you, including accurate button handling.
func NewMouseI(args ...interface{}) *Mouse {
    return &Mouse{js.Global.Get("Phaser").Get("Mouse").New(args)}
}



// Game A reference to the currently running game.
func (self *Mouse) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *Mouse) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Input A reference to the Phaser Input Manager.
func (self *Mouse) Input() *Input{
    return &Input{self.Object.Get("input")}
}

// SetInputA A reference to the Phaser Input Manager.
func (self *Mouse) SetInputA(member *Input) {
    self.Object.Set("input", member)
}

// CallbackContext The context under which callbacks are called.
func (self *Mouse) CallbackContext() interface{}{
    return self.Object.Get("callbackContext")
}

// SetCallbackContextA The context under which callbacks are called.
func (self *Mouse) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// MouseDownCallback A callback that can be fired when the mouse is pressed down.
func (self *Mouse) MouseDownCallback() interface{}{
    return self.Object.Get("mouseDownCallback")
}

// SetMouseDownCallbackA A callback that can be fired when the mouse is pressed down.
func (self *Mouse) SetMouseDownCallbackA(member interface{}) {
    self.Object.Set("mouseDownCallback", member)
}

// MouseUpCallback A callback that can be fired when the mouse is released from a pressed down state.
func (self *Mouse) MouseUpCallback() interface{}{
    return self.Object.Get("mouseUpCallback")
}

// SetMouseUpCallbackA A callback that can be fired when the mouse is released from a pressed down state.
func (self *Mouse) SetMouseUpCallbackA(member interface{}) {
    self.Object.Set("mouseUpCallback", member)
}

// MouseOutCallback A callback that can be fired when the mouse is no longer over the game canvas.
func (self *Mouse) MouseOutCallback() interface{}{
    return self.Object.Get("mouseOutCallback")
}

// SetMouseOutCallbackA A callback that can be fired when the mouse is no longer over the game canvas.
func (self *Mouse) SetMouseOutCallbackA(member interface{}) {
    self.Object.Set("mouseOutCallback", member)
}

// MouseOverCallback A callback that can be fired when the mouse enters the game canvas (usually after a mouseout).
func (self *Mouse) MouseOverCallback() interface{}{
    return self.Object.Get("mouseOverCallback")
}

// SetMouseOverCallbackA A callback that can be fired when the mouse enters the game canvas (usually after a mouseout).
func (self *Mouse) SetMouseOverCallbackA(member interface{}) {
    self.Object.Set("mouseOverCallback", member)
}

// MouseWheelCallback A callback that can be fired when the mousewheel is used.
func (self *Mouse) MouseWheelCallback() interface{}{
    return self.Object.Get("mouseWheelCallback")
}

// SetMouseWheelCallbackA A callback that can be fired when the mousewheel is used.
func (self *Mouse) SetMouseWheelCallbackA(member interface{}) {
    self.Object.Set("mouseWheelCallback", member)
}

// Capture If true the DOM mouse events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *Mouse) Capture() bool{
    return self.Object.Get("capture").Bool()
}

// SetCaptureA If true the DOM mouse events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *Mouse) SetCaptureA(member bool) {
    self.Object.Set("capture", member)
}

// Button This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *Mouse) Button() int{
    return self.Object.Get("button").Int()
}

// SetButtonA This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *Mouse) SetButtonA(member int) {
    self.Object.Set("button", member)
}

// WheelDelta The direction of the _last_ mousewheel usage 1 for up -1 for down.
func (self *Mouse) WheelDelta() int{
    return self.Object.Get("wheelDelta").Int()
}

// SetWheelDeltaA The direction of the _last_ mousewheel usage 1 for up -1 for down.
func (self *Mouse) SetWheelDeltaA(member int) {
    self.Object.Set("wheelDelta", member)
}

// Enabled Mouse input will only be processed if enabled.
func (self *Mouse) Enabled() bool{
    return self.Object.Get("enabled").Bool()
}

// SetEnabledA Mouse input will only be processed if enabled.
func (self *Mouse) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}

// Locked If the mouse has been Pointer Locked successfully this will be set to true.
func (self *Mouse) Locked() bool{
    return self.Object.Get("locked").Bool()
}

// SetLockedA If the mouse has been Pointer Locked successfully this will be set to true.
func (self *Mouse) SetLockedA(member bool) {
    self.Object.Set("locked", member)
}

// StopOnGameOut If true Pointer.stop will be called if the mouse leaves the game canvas.
func (self *Mouse) StopOnGameOut() bool{
    return self.Object.Get("stopOnGameOut").Bool()
}

// SetStopOnGameOutA If true Pointer.stop will be called if the mouse leaves the game canvas.
func (self *Mouse) SetStopOnGameOutA(member bool) {
    self.Object.Set("stopOnGameOut", member)
}

// PointerLock This event is dispatched when the browser enters or leaves pointer lock state.
func (self *Mouse) PointerLock() *Signal{
    return &Signal{self.Object.Get("pointerLock")}
}

// SetPointerLockA This event is dispatched when the browser enters or leaves pointer lock state.
func (self *Mouse) SetPointerLockA(member *Signal) {
    self.Object.Set("pointerLock", member)
}

// Event The browser mouse DOM event. Will be null if no mouse event has ever been received.
// Access this property only inside a Mouse event handler and do not keep references to it.
func (self *Mouse) Event() interface{}{
    return self.Object.Get("event")
}

// SetEventA The browser mouse DOM event. Will be null if no mouse event has ever been received.
// Access this property only inside a Mouse event handler and do not keep references to it.
func (self *Mouse) SetEventA(member interface{}) {
    self.Object.Set("event", member)
}

// NO_BUTTON empty description
func (self *Mouse) NO_BUTTON() int{
    return self.Object.Get("NO_BUTTON").Int()
}

// SetNO_BUTTONA empty description
func (self *Mouse) SetNO_BUTTONA(member int) {
    self.Object.Set("NO_BUTTON", member)
}

// LEFT_BUTTON empty description
func (self *Mouse) LEFT_BUTTON() int{
    return self.Object.Get("LEFT_BUTTON").Int()
}

// SetLEFT_BUTTONA empty description
func (self *Mouse) SetLEFT_BUTTONA(member int) {
    self.Object.Set("LEFT_BUTTON", member)
}

// MIDDLE_BUTTON empty description
func (self *Mouse) MIDDLE_BUTTON() int{
    return self.Object.Get("MIDDLE_BUTTON").Int()
}

// SetMIDDLE_BUTTONA empty description
func (self *Mouse) SetMIDDLE_BUTTONA(member int) {
    self.Object.Set("MIDDLE_BUTTON", member)
}

// RIGHT_BUTTON empty description
func (self *Mouse) RIGHT_BUTTON() int{
    return self.Object.Get("RIGHT_BUTTON").Int()
}

// SetRIGHT_BUTTONA empty description
func (self *Mouse) SetRIGHT_BUTTONA(member int) {
    self.Object.Set("RIGHT_BUTTON", member)
}

// BACK_BUTTON empty description
func (self *Mouse) BACK_BUTTON() int{
    return self.Object.Get("BACK_BUTTON").Int()
}

// SetBACK_BUTTONA empty description
func (self *Mouse) SetBACK_BUTTONA(member int) {
    self.Object.Set("BACK_BUTTON", member)
}

// FORWARD_BUTTON empty description
func (self *Mouse) FORWARD_BUTTON() int{
    return self.Object.Get("FORWARD_BUTTON").Int()
}

// SetFORWARD_BUTTONA empty description
func (self *Mouse) SetFORWARD_BUTTONA(member int) {
    self.Object.Set("FORWARD_BUTTON", member)
}

// WHEEL_UP empty description
func (self *Mouse) WHEEL_UP() int{
    return self.Object.Get("WHEEL_UP").Int()
}

// SetWHEEL_UPA empty description
func (self *Mouse) SetWHEEL_UPA(member int) {
    self.Object.Set("WHEEL_UP", member)
}

// WHEEL_DOWN empty description
func (self *Mouse) WHEEL_DOWN() int{
    return self.Object.Get("WHEEL_DOWN").Int()
}

// SetWHEEL_DOWNA empty description
func (self *Mouse) SetWHEEL_DOWNA(member int) {
    self.Object.Set("WHEEL_DOWN", member)
}


// Start Starts the event listeners running.
func (self *Mouse) Start() {
    self.Object.Call("start")
}

// StartI Starts the event listeners running.
func (self *Mouse) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// OnMouseDown The internal method that handles the mouse down event from the browser.
func (self *Mouse) OnMouseDown(event *MouseEvent) {
    self.Object.Call("onMouseDown", event)
}

// OnMouseDownI The internal method that handles the mouse down event from the browser.
func (self *Mouse) OnMouseDownI(args ...interface{}) {
    self.Object.Call("onMouseDown", args)
}

// OnMouseMove The internal method that handles the mouse move event from the browser.
func (self *Mouse) OnMouseMove(event *MouseEvent) {
    self.Object.Call("onMouseMove", event)
}

// OnMouseMoveI The internal method that handles the mouse move event from the browser.
func (self *Mouse) OnMouseMoveI(args ...interface{}) {
    self.Object.Call("onMouseMove", args)
}

// OnMouseUp The internal method that handles the mouse up event from the browser.
func (self *Mouse) OnMouseUp(event *MouseEvent) {
    self.Object.Call("onMouseUp", event)
}

// OnMouseUpI The internal method that handles the mouse up event from the browser.
func (self *Mouse) OnMouseUpI(args ...interface{}) {
    self.Object.Call("onMouseUp", args)
}

// OnMouseUpGlobal The internal method that handles the mouse up event from the window.
func (self *Mouse) OnMouseUpGlobal(event *MouseEvent) {
    self.Object.Call("onMouseUpGlobal", event)
}

// OnMouseUpGlobalI The internal method that handles the mouse up event from the window.
func (self *Mouse) OnMouseUpGlobalI(args ...interface{}) {
    self.Object.Call("onMouseUpGlobal", args)
}

// OnMouseOutGlobal The internal method that handles the mouse out event from the window.
func (self *Mouse) OnMouseOutGlobal(event *MouseEvent) {
    self.Object.Call("onMouseOutGlobal", event)
}

// OnMouseOutGlobalI The internal method that handles the mouse out event from the window.
func (self *Mouse) OnMouseOutGlobalI(args ...interface{}) {
    self.Object.Call("onMouseOutGlobal", args)
}

// OnMouseOut The internal method that handles the mouse out event from the browser.
func (self *Mouse) OnMouseOut(event *MouseEvent) {
    self.Object.Call("onMouseOut", event)
}

// OnMouseOutI The internal method that handles the mouse out event from the browser.
func (self *Mouse) OnMouseOutI(args ...interface{}) {
    self.Object.Call("onMouseOut", args)
}

// OnMouseOver The internal method that handles the mouse over event from the browser.
func (self *Mouse) OnMouseOver(event *MouseEvent) {
    self.Object.Call("onMouseOver", event)
}

// OnMouseOverI The internal method that handles the mouse over event from the browser.
func (self *Mouse) OnMouseOverI(args ...interface{}) {
    self.Object.Call("onMouseOver", args)
}

// OnMouseWheel The internal method that handles the mouse wheel event from the browser.
func (self *Mouse) OnMouseWheel(event *MouseEvent) {
    self.Object.Call("onMouseWheel", event)
}

// OnMouseWheelI The internal method that handles the mouse wheel event from the browser.
func (self *Mouse) OnMouseWheelI(args ...interface{}) {
    self.Object.Call("onMouseWheel", args)
}

// RequestPointerLock If the browser supports it you can request that the pointer be locked to the browser window.
// This is classically known as 'FPS controls', where the pointer can't leave the browser until the user presses an exit key.
// If the browser successfully enters a locked state the event Phaser.Mouse.pointerLock will be dispatched and the first parameter will be 'true'.
func (self *Mouse) RequestPointerLock() {
    self.Object.Call("requestPointerLock")
}

// RequestPointerLockI If the browser supports it you can request that the pointer be locked to the browser window.
// This is classically known as 'FPS controls', where the pointer can't leave the browser until the user presses an exit key.
// If the browser successfully enters a locked state the event Phaser.Mouse.pointerLock will be dispatched and the first parameter will be 'true'.
func (self *Mouse) RequestPointerLockI(args ...interface{}) {
    self.Object.Call("requestPointerLock", args)
}

// PointerLockChange Internal pointerLockChange handler.
func (self *Mouse) PointerLockChange(event *Event) {
    self.Object.Call("pointerLockChange", event)
}

// PointerLockChangeI Internal pointerLockChange handler.
func (self *Mouse) PointerLockChangeI(args ...interface{}) {
    self.Object.Call("pointerLockChange", args)
}

// ReleasePointerLock Internal release pointer lock handler.
func (self *Mouse) ReleasePointerLock() {
    self.Object.Call("releasePointerLock")
}

// ReleasePointerLockI Internal release pointer lock handler.
func (self *Mouse) ReleasePointerLockI(args ...interface{}) {
    self.Object.Call("releasePointerLock", args)
}

// Stop Stop the event listeners.
func (self *Mouse) Stop() {
    self.Object.Call("stop")
}

// StopI Stop the event listeners.
func (self *Mouse) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// WheelEventProxy A purely internal event support class to proxy 'wheelscroll' and 'DOMMouseWheel'
// events to 'wheel'-like events.
// 
// See https://developer.mozilla.org/en-US/docs/Web/Events/mousewheel for choosing a scale and delta mode.
func (self *Mouse) WheelEventProxy(scaleFactor int, deltaMode int) {
    self.Object.Call("WheelEventProxy", scaleFactor, deltaMode)
}

// WheelEventProxyI A purely internal event support class to proxy 'wheelscroll' and 'DOMMouseWheel'
// events to 'wheel'-like events.
// 
// See https://developer.mozilla.org/en-US/docs/Web/Events/mousewheel for choosing a scale and delta mode.
func (self *Mouse) WheelEventProxyI(args ...interface{}) {
    self.Object.Call("WheelEventProxy", args)
}

