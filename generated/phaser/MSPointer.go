// Automatic generation for Phaser.MSPointer
// generated file MSPointer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The MSPointer class handles Microsoft touch interactions with the game and the resulting Pointer objects.
// 
// It will work only in Internet Explorer 10+ and Windows Store or Windows Phone 8 apps using JavaScript.
// http://msdn.microsoft.com/en-us/library/ie/hh673557(v=vs.85).aspx
// 
// You should not normally access this class directly, but instead use a Phaser.Pointer object which 
// normalises all game input for you including accurate button handling.
// 
// Please note that at the current time of writing Phaser does not yet support chorded button interactions:
// http://www.w3.org/TR/pointerevents/#chorded-button-interactions
type MSPointer struct {
    *js.Object
}


// A reference to the currently running game.
func (self *MSPointer) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *MSPointer) SetGame(member *Game) {
    self.Set("game", member)
}

// A reference to the Phaser Input Manager.
func (self *MSPointer) GetInput() *Input{
    return &Input{self.Get("input")}
}

// A reference to the Phaser Input Manager.
func (self *MSPointer) SetInput(member *Input) {
    self.Set("input", member)
}

// The context under which callbacks are called (defaults to game).
func (self *MSPointer) GetCallbackContext() interface{}{
    return self.Get("callbackContext")
}

// The context under which callbacks are called (defaults to game).
func (self *MSPointer) SetCallbackContext(member interface{}) {
    self.Set("callbackContext", member)
}

// A callback that can be fired on a MSPointerDown event.
func (self *MSPointer) SetPointerDownCallback(member func(...interface{})) {
    self.Set("pointerDownCallback", member)
}

// A callback that can be fired on a MSPointerMove event.
func (self *MSPointer) SetPointerMoveCallback(member func(...interface{})) {
    self.Set("pointerMoveCallback", member)
}

// A callback that can be fired on a MSPointerUp event.
func (self *MSPointer) SetPointerUpCallback(member func(...interface{})) {
    self.Set("pointerUpCallback", member)
}

// If true the Pointer events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *MSPointer) GetCapture() bool{
    return self.Get("capture").Bool()
}

// If true the Pointer events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *MSPointer) SetCapture(member bool) {
    self.Set("capture", member)
}

// This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *MSPointer) GetButton() int{
    return self.Get("button").Int()
}

// This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *MSPointer) SetButton(member int) {
    self.Set("button", member)
}

// The browser MSPointer DOM event. Will be null if no event has ever been received.
// Access this property only inside a Pointer event handler and do not keep references to it.
func (self *MSPointer) GetEvent() interface{}{
    return self.Get("event")
}

// The browser MSPointer DOM event. Will be null if no event has ever been received.
// Access this property only inside a Pointer event handler and do not keep references to it.
func (self *MSPointer) SetEvent(member interface{}) {
    self.Set("event", member)
}

// MSPointer input will only be processed if enabled.
func (self *MSPointer) GetEnabled() bool{
    return self.Get("enabled").Bool()
}

// MSPointer input will only be processed if enabled.
func (self *MSPointer) SetEnabled(member bool) {
    self.Set("enabled", member)
}



// Starts the event listeners running.
func (self *MSPointer) StartI(args ...interface{}) {
    self.Call("start", args)
}

// The function that handles the PointerDown event.
func (self *MSPointer) OnPointerDownI(args ...interface{}) {
    self.Call("onPointerDown", args)
}

// The function that handles the PointerMove event.
func (self *MSPointer) OnPointerMoveI(args ...interface{}) {
    self.Call("onPointerMove", args)
}

// The function that handles the PointerUp event.
func (self *MSPointer) OnPointerUpI(args ...interface{}) {
    self.Call("onPointerUp", args)
}

// The internal method that handles the mouse up event from the window.
func (self *MSPointer) OnPointerUpGlobalI(args ...interface{}) {
    self.Call("onPointerUpGlobal", args)
}

// The internal method that handles the pointer out event from the browser.
func (self *MSPointer) OnPointerOutI(args ...interface{}) {
    self.Call("onPointerOut", args)
}

// Stop the event listeners.
func (self *MSPointer) StopI(args ...interface{}) {
    self.Call("stop", args)
}
