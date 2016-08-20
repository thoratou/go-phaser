// Automatic generation for Phaser.Touch
// generated file Touch.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Phaser.Touch handles touch events with your game. Note: Android 2.x only supports 1 touch event at once, no multi-touch.
// 
// You should not normally access this class directly, but instead use a Phaser.Pointer object which normalises all game input for you.
type Touch struct {
    *js.Object
}


// A reference to the currently running game.
func (self *Touch) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *Touch) SetGame(member Game) {
    self.Set("game", member)
}

// Touch events will only be processed if enabled.
func (self *Touch) GetEnabled() bool{
    return self.Get("enabled").Bool()
}

// Touch events will only be processed if enabled.
func (self *Touch) SetEnabled(member bool) {
    self.Set("enabled", member)
}

// An array of callbacks that will be fired every time a native touch start or touch end event is received from the browser.
// This is used internally to handle audio and video unlocking on mobile devices.
// To add a callback to this array please use `Touch.addTouchLockCallback`.
func (self *Touch) GetTouchLockCallbacks() []interface{}{
	array := self.Get("touchLockCallbacks")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of callbacks that will be fired every time a native touch start or touch end event is received from the browser.
// This is used internally to handle audio and video unlocking on mobile devices.
// To add a callback to this array please use `Touch.addTouchLockCallback`.
func (self *Touch) SetTouchLockCallbacks(member []interface{}) {
    self.Set("touchLockCallbacks", member)
}

// The context under which callbacks are called.
func (self *Touch) GetCallbackContext() interface{}{
    return self.Get("callbackContext")
}

// The context under which callbacks are called.
func (self *Touch) SetCallbackContext(member interface{}) {
    self.Set("callbackContext", member)
}

// A callback that can be fired on a touchStart event.
func (self *Touch) SetTouchStartCallback(member func(...interface{})) {
    self.Set("touchStartCallback", member)
}

// A callback that can be fired on a touchMove event.
func (self *Touch) SetTouchMoveCallback(member func(...interface{})) {
    self.Set("touchMoveCallback", member)
}

// A callback that can be fired on a touchEnd event.
func (self *Touch) SetTouchEndCallback(member func(...interface{})) {
    self.Set("touchEndCallback", member)
}

// A callback that can be fired on a touchEnter event.
func (self *Touch) SetTouchEnterCallback(member func(...interface{})) {
    self.Set("touchEnterCallback", member)
}

// A callback that can be fired on a touchLeave event.
func (self *Touch) SetTouchLeaveCallback(member func(...interface{})) {
    self.Set("touchLeaveCallback", member)
}

// A callback that can be fired on a touchCancel event.
func (self *Touch) SetTouchCancelCallback(member func(...interface{})) {
    self.Set("touchCancelCallback", member)
}

// If true the TouchEvent will have prevent.default called on it.
func (self *Touch) GetPreventDefault() bool{
    return self.Get("preventDefault").Bool()
}

// If true the TouchEvent will have prevent.default called on it.
func (self *Touch) SetPreventDefault(member bool) {
    self.Set("preventDefault", member)
}

// The browser touch DOM event. Will be set to null if no touch event has ever been received.
func (self *Touch) GetEvent() TouchEvent{
    return TouchEvent{self.Get("event")}
}

// The browser touch DOM event. Will be set to null if no touch event has ever been received.
func (self *Touch) SetEvent(member TouchEvent) {
    self.Set("event", member)
}



// Starts the event listeners running.
func (self *Touch) StartI(args ...interface{}) {
    self.Call("start", args)
}

// Consumes all touchmove events on the document (only enable this if you know you need it!).
func (self *Touch) ConsumeTouchMoveI(args ...interface{}) {
    self.Call("consumeTouchMove", args)
}

// Adds a callback that is fired when a browser touchstart or touchend event is received.
// 
// This is used internally to handle audio and video unlocking on mobile devices.
// 
// If the callback returns 'true' then the callback is automatically deleted once invoked.
// 
// The callback is added to the Phaser.Touch.touchLockCallbacks array and should be removed with Phaser.Touch.removeTouchLockCallback.
func (self *Touch) AddTouchLockCallbackI(args ...interface{}) {
    self.Call("addTouchLockCallback", args)
}

// Removes the callback at the defined index from the Phaser.Touch.touchLockCallbacks array
func (self *Touch) RemoveTouchLockCallbackI(args ...interface{}) bool{
    return self.Call("removeTouchLockCallback", args).Bool()
}

// The internal method that handles the touchstart event from the browser.
func (self *Touch) OnTouchStartI(args ...interface{}) {
    self.Call("onTouchStart", args)
}

// Touch cancel - touches that were disrupted (perhaps by moving into a plugin or browser chrome).
// Occurs for example on iOS when you put down 4 fingers and the app selector UI appears.
func (self *Touch) OnTouchCancelI(args ...interface{}) {
    self.Call("onTouchCancel", args)
}

// For touch enter and leave its a list of the touch points that have entered or left the target.
// Doesn't appear to be supported by most browsers on a canvas element yet.
func (self *Touch) OnTouchEnterI(args ...interface{}) {
    self.Call("onTouchEnter", args)
}

// For touch enter and leave its a list of the touch points that have entered or left the target.
// Doesn't appear to be supported by most browsers on a canvas element yet.
func (self *Touch) OnTouchLeaveI(args ...interface{}) {
    self.Call("onTouchLeave", args)
}

// The handler for the touchmove events.
func (self *Touch) OnTouchMoveI(args ...interface{}) {
    self.Call("onTouchMove", args)
}

// The handler for the touchend events.
func (self *Touch) OnTouchEndI(args ...interface{}) {
    self.Call("onTouchEnd", args)
}

// Stop the event listeners.
func (self *Touch) StopI(args ...interface{}) {
    self.Call("stop", args)
}
