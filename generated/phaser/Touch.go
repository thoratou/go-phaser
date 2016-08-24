// Package phaser Automatic generation for Phaser.Touch
// generated file Touch.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Touch Phaser.Touch handles touch events with your game. Note: Android 2.x only supports 1 touch event at once, no multi-touch.
// 
// You should not normally access this class directly, but instead use a Phaser.Pointer object which normalises all game input for you.
type Touch struct {
    *js.Object
}

// NewTouch Phaser.Touch handles touch events with your game. Note: Android 2.x only supports 1 touch event at once, no multi-touch.
// 
// You should not normally access this class directly, but instead use a Phaser.Pointer object which normalises all game input for you.
func NewTouch(game *Game) *Touch {
    return &Touch{js.Global.Get("Phaser").Get("Touch").New(game)}
}
// NewTouchI Phaser.Touch handles touch events with your game. Note: Android 2.x only supports 1 touch event at once, no multi-touch.
// 
// You should not normally access this class directly, but instead use a Phaser.Pointer object which normalises all game input for you.
func NewTouchI(args ...interface{}) *Touch {
    return &Touch{js.Global.Get("Phaser").Get("Touch").New(args)}
}



// Game A reference to the currently running game.
func (self *Touch) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *Touch) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Enabled Touch events will only be processed if enabled.
func (self *Touch) Enabled() bool{
    return self.Object.Get("enabled").Bool()
}

// SetEnabledA Touch events will only be processed if enabled.
func (self *Touch) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}

// TouchLockCallbacks An array of callbacks that will be fired every time a native touch start or touch end event is received from the browser.
// This is used internally to handle audio and video unlocking on mobile devices.
// To add a callback to this array please use `Touch.addTouchLockCallback`.
func (self *Touch) TouchLockCallbacks() []interface{}{
	array00 := self.Object.Get("touchLockCallbacks")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetTouchLockCallbacksA An array of callbacks that will be fired every time a native touch start or touch end event is received from the browser.
// This is used internally to handle audio and video unlocking on mobile devices.
// To add a callback to this array please use `Touch.addTouchLockCallback`.
func (self *Touch) SetTouchLockCallbacksA(member []interface{}) {
    self.Object.Set("touchLockCallbacks", member)
}

// CallbackContext The context under which callbacks are called.
func (self *Touch) CallbackContext() interface{}{
    return self.Object.Get("callbackContext")
}

// SetCallbackContextA The context under which callbacks are called.
func (self *Touch) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// TouchStartCallback A callback that can be fired on a touchStart event.
func (self *Touch) TouchStartCallback() interface{}{
    return self.Object.Get("touchStartCallback")
}

// SetTouchStartCallbackA A callback that can be fired on a touchStart event.
func (self *Touch) SetTouchStartCallbackA(member interface{}) {
    self.Object.Set("touchStartCallback", member)
}

// TouchMoveCallback A callback that can be fired on a touchMove event.
func (self *Touch) TouchMoveCallback() interface{}{
    return self.Object.Get("touchMoveCallback")
}

// SetTouchMoveCallbackA A callback that can be fired on a touchMove event.
func (self *Touch) SetTouchMoveCallbackA(member interface{}) {
    self.Object.Set("touchMoveCallback", member)
}

// TouchEndCallback A callback that can be fired on a touchEnd event.
func (self *Touch) TouchEndCallback() interface{}{
    return self.Object.Get("touchEndCallback")
}

// SetTouchEndCallbackA A callback that can be fired on a touchEnd event.
func (self *Touch) SetTouchEndCallbackA(member interface{}) {
    self.Object.Set("touchEndCallback", member)
}

// TouchEnterCallback A callback that can be fired on a touchEnter event.
func (self *Touch) TouchEnterCallback() interface{}{
    return self.Object.Get("touchEnterCallback")
}

// SetTouchEnterCallbackA A callback that can be fired on a touchEnter event.
func (self *Touch) SetTouchEnterCallbackA(member interface{}) {
    self.Object.Set("touchEnterCallback", member)
}

// TouchLeaveCallback A callback that can be fired on a touchLeave event.
func (self *Touch) TouchLeaveCallback() interface{}{
    return self.Object.Get("touchLeaveCallback")
}

// SetTouchLeaveCallbackA A callback that can be fired on a touchLeave event.
func (self *Touch) SetTouchLeaveCallbackA(member interface{}) {
    self.Object.Set("touchLeaveCallback", member)
}

// TouchCancelCallback A callback that can be fired on a touchCancel event.
func (self *Touch) TouchCancelCallback() interface{}{
    return self.Object.Get("touchCancelCallback")
}

// SetTouchCancelCallbackA A callback that can be fired on a touchCancel event.
func (self *Touch) SetTouchCancelCallbackA(member interface{}) {
    self.Object.Set("touchCancelCallback", member)
}

// PreventDefault If true the TouchEvent will have prevent.default called on it.
func (self *Touch) PreventDefault() bool{
    return self.Object.Get("preventDefault").Bool()
}

// SetPreventDefaultA If true the TouchEvent will have prevent.default called on it.
func (self *Touch) SetPreventDefaultA(member bool) {
    self.Object.Set("preventDefault", member)
}

// Event The browser touch DOM event. Will be set to null if no touch event has ever been received.
func (self *Touch) Event() *TouchEvent{
    return &TouchEvent{self.Object.Get("event")}
}

// SetEventA The browser touch DOM event. Will be set to null if no touch event has ever been received.
func (self *Touch) SetEventA(member *TouchEvent) {
    self.Object.Set("event", member)
}


// Start Starts the event listeners running.
func (self *Touch) Start() {
    self.Object.Call("start")
}

// StartI Starts the event listeners running.
func (self *Touch) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// ConsumeTouchMove Consumes all touchmove events on the document (only enable this if you know you need it!).
func (self *Touch) ConsumeTouchMove() {
    self.Object.Call("consumeTouchMove")
}

// ConsumeTouchMoveI Consumes all touchmove events on the document (only enable this if you know you need it!).
func (self *Touch) ConsumeTouchMoveI(args ...interface{}) {
    self.Object.Call("consumeTouchMove", args)
}

// AddTouchLockCallback Adds a callback that is fired when a browser touchstart or touchend event is received.
// 
// This is used internally to handle audio and video unlocking on mobile devices.
// 
// If the callback returns 'true' then the callback is automatically deleted once invoked.
// 
// The callback is added to the Phaser.Touch.touchLockCallbacks array and should be removed with Phaser.Touch.removeTouchLockCallback.
func (self *Touch) AddTouchLockCallback(callback interface{}, context interface{}) {
    self.Object.Call("addTouchLockCallback", callback, context)
}

// AddTouchLockCallback1O Adds a callback that is fired when a browser touchstart or touchend event is received.
// 
// This is used internally to handle audio and video unlocking on mobile devices.
// 
// If the callback returns 'true' then the callback is automatically deleted once invoked.
// 
// The callback is added to the Phaser.Touch.touchLockCallbacks array and should be removed with Phaser.Touch.removeTouchLockCallback.
func (self *Touch) AddTouchLockCallback1O(callback interface{}, context interface{}, onEnd bool) {
    self.Object.Call("addTouchLockCallback", callback, context, onEnd)
}

// AddTouchLockCallbackI Adds a callback that is fired when a browser touchstart or touchend event is received.
// 
// This is used internally to handle audio and video unlocking on mobile devices.
// 
// If the callback returns 'true' then the callback is automatically deleted once invoked.
// 
// The callback is added to the Phaser.Touch.touchLockCallbacks array and should be removed with Phaser.Touch.removeTouchLockCallback.
func (self *Touch) AddTouchLockCallbackI(args ...interface{}) {
    self.Object.Call("addTouchLockCallback", args)
}

// RemoveTouchLockCallback Removes the callback at the defined index from the Phaser.Touch.touchLockCallbacks array
func (self *Touch) RemoveTouchLockCallback(callback interface{}, context interface{}) bool{
    return self.Object.Call("removeTouchLockCallback", callback, context).Bool()
}

// RemoveTouchLockCallbackI Removes the callback at the defined index from the Phaser.Touch.touchLockCallbacks array
func (self *Touch) RemoveTouchLockCallbackI(args ...interface{}) bool{
    return self.Object.Call("removeTouchLockCallback", args).Bool()
}

// OnTouchStart The internal method that handles the touchstart event from the browser.
func (self *Touch) OnTouchStart(event *TouchEvent) {
    self.Object.Call("onTouchStart", event)
}

// OnTouchStartI The internal method that handles the touchstart event from the browser.
func (self *Touch) OnTouchStartI(args ...interface{}) {
    self.Object.Call("onTouchStart", args)
}

// OnTouchCancel Touch cancel - touches that were disrupted (perhaps by moving into a plugin or browser chrome).
// Occurs for example on iOS when you put down 4 fingers and the app selector UI appears.
func (self *Touch) OnTouchCancel(event *TouchEvent) {
    self.Object.Call("onTouchCancel", event)
}

// OnTouchCancelI Touch cancel - touches that were disrupted (perhaps by moving into a plugin or browser chrome).
// Occurs for example on iOS when you put down 4 fingers and the app selector UI appears.
func (self *Touch) OnTouchCancelI(args ...interface{}) {
    self.Object.Call("onTouchCancel", args)
}

// OnTouchEnter For touch enter and leave its a list of the touch points that have entered or left the target.
// Doesn't appear to be supported by most browsers on a canvas element yet.
func (self *Touch) OnTouchEnter(event *TouchEvent) {
    self.Object.Call("onTouchEnter", event)
}

// OnTouchEnterI For touch enter and leave its a list of the touch points that have entered or left the target.
// Doesn't appear to be supported by most browsers on a canvas element yet.
func (self *Touch) OnTouchEnterI(args ...interface{}) {
    self.Object.Call("onTouchEnter", args)
}

// OnTouchLeave For touch enter and leave its a list of the touch points that have entered or left the target.
// Doesn't appear to be supported by most browsers on a canvas element yet.
func (self *Touch) OnTouchLeave(event *TouchEvent) {
    self.Object.Call("onTouchLeave", event)
}

// OnTouchLeaveI For touch enter and leave its a list of the touch points that have entered or left the target.
// Doesn't appear to be supported by most browsers on a canvas element yet.
func (self *Touch) OnTouchLeaveI(args ...interface{}) {
    self.Object.Call("onTouchLeave", args)
}

// OnTouchMove The handler for the touchmove events.
func (self *Touch) OnTouchMove(event *TouchEvent) {
    self.Object.Call("onTouchMove", event)
}

// OnTouchMoveI The handler for the touchmove events.
func (self *Touch) OnTouchMoveI(args ...interface{}) {
    self.Object.Call("onTouchMove", args)
}

// OnTouchEnd The handler for the touchend events.
func (self *Touch) OnTouchEnd(event *TouchEvent) {
    self.Object.Call("onTouchEnd", event)
}

// OnTouchEndI The handler for the touchend events.
func (self *Touch) OnTouchEndI(args ...interface{}) {
    self.Object.Call("onTouchEnd", args)
}

// Stop Stop the event listeners.
func (self *Touch) Stop() {
    self.Object.Call("stop")
}

// StopI Stop the event listeners.
func (self *Touch) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

