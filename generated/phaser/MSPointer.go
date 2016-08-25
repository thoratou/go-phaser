// Package phaser Automatic generation for Phaser.MSPointer
// generated file MSPointer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// MSPointer The MSPointer class handles Microsoft touch interactions with the game and the resulting Pointer objects.
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

// NewMSPointer The MSPointer class handles Microsoft touch interactions with the game and the resulting Pointer objects.
// 
// It will work only in Internet Explorer 10+ and Windows Store or Windows Phone 8 apps using JavaScript.
// http://msdn.microsoft.com/en-us/library/ie/hh673557(v=vs.85).aspx
// 
// You should not normally access this class directly, but instead use a Phaser.Pointer object which 
// normalises all game input for you including accurate button handling.
// 
// Please note that at the current time of writing Phaser does not yet support chorded button interactions:
// http://www.w3.org/TR/pointerevents/#chorded-button-interactions
func NewMSPointer(game *Game) *MSPointer {
    return &MSPointer{js.Global.Get("Phaser").Get("MSPointer").New(game)}
}
// NewMSPointerI The MSPointer class handles Microsoft touch interactions with the game and the resulting Pointer objects.
// 
// It will work only in Internet Explorer 10+ and Windows Store or Windows Phone 8 apps using JavaScript.
// http://msdn.microsoft.com/en-us/library/ie/hh673557(v=vs.85).aspx
// 
// You should not normally access this class directly, but instead use a Phaser.Pointer object which 
// normalises all game input for you including accurate button handling.
// 
// Please note that at the current time of writing Phaser does not yet support chorded button interactions:
// http://www.w3.org/TR/pointerevents/#chorded-button-interactions
func NewMSPointerI(args ...interface{}) *MSPointer {
    return &MSPointer{js.Global.Get("Phaser").Get("MSPointer").New(args)}
}



// MSPointer Binding conversion method to MSPointer point 
func ToMSPointer(jsStruct interface{}) *MSPointer {
    if object, ok := jsStruct.(*js.Object); ok {
		return &MSPointer{Object: object}
	}
	return nil
}



// Game A reference to the currently running game.
func (self *MSPointer) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *MSPointer) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Input A reference to the Phaser Input Manager.
func (self *MSPointer) Input() *Input{
    return &Input{self.Object.Get("input")}
}

// SetInputA A reference to the Phaser Input Manager.
func (self *MSPointer) SetInputA(member *Input) {
    self.Object.Set("input", member)
}

// CallbackContext The context under which callbacks are called (defaults to game).
func (self *MSPointer) CallbackContext() interface{}{
    return self.Object.Get("callbackContext")
}

// SetCallbackContextA The context under which callbacks are called (defaults to game).
func (self *MSPointer) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// PointerDownCallback A callback that can be fired on a MSPointerDown event.
func (self *MSPointer) PointerDownCallback() interface{}{
    return self.Object.Get("pointerDownCallback")
}

// SetPointerDownCallbackA A callback that can be fired on a MSPointerDown event.
func (self *MSPointer) SetPointerDownCallbackA(member interface{}) {
    self.Object.Set("pointerDownCallback", member)
}

// PointerMoveCallback A callback that can be fired on a MSPointerMove event.
func (self *MSPointer) PointerMoveCallback() interface{}{
    return self.Object.Get("pointerMoveCallback")
}

// SetPointerMoveCallbackA A callback that can be fired on a MSPointerMove event.
func (self *MSPointer) SetPointerMoveCallbackA(member interface{}) {
    self.Object.Set("pointerMoveCallback", member)
}

// PointerUpCallback A callback that can be fired on a MSPointerUp event.
func (self *MSPointer) PointerUpCallback() interface{}{
    return self.Object.Get("pointerUpCallback")
}

// SetPointerUpCallbackA A callback that can be fired on a MSPointerUp event.
func (self *MSPointer) SetPointerUpCallbackA(member interface{}) {
    self.Object.Set("pointerUpCallback", member)
}

// Capture If true the Pointer events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *MSPointer) Capture() bool{
    return self.Object.Get("capture").Bool()
}

// SetCaptureA If true the Pointer events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *MSPointer) SetCaptureA(member bool) {
    self.Object.Set("capture", member)
}

// Button This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *MSPointer) Button() int{
    return self.Object.Get("button").Int()
}

// SetButtonA This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *MSPointer) SetButtonA(member int) {
    self.Object.Set("button", member)
}

// Event The browser MSPointer DOM event. Will be null if no event has ever been received.
// Access this property only inside a Pointer event handler and do not keep references to it.
func (self *MSPointer) Event() interface{}{
    return self.Object.Get("event")
}

// SetEventA The browser MSPointer DOM event. Will be null if no event has ever been received.
// Access this property only inside a Pointer event handler and do not keep references to it.
func (self *MSPointer) SetEventA(member interface{}) {
    self.Object.Set("event", member)
}

// Enabled MSPointer input will only be processed if enabled.
func (self *MSPointer) Enabled() bool{
    return self.Object.Get("enabled").Bool()
}

// SetEnabledA MSPointer input will only be processed if enabled.
func (self *MSPointer) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}


// Start Starts the event listeners running.
func (self *MSPointer) Start() {
    self.Object.Call("start")
}

// StartI Starts the event listeners running.
func (self *MSPointer) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// OnPointerDown The function that handles the PointerDown event.
func (self *MSPointer) OnPointerDown(event *PointerEvent) {
    self.Object.Call("onPointerDown", event)
}

// OnPointerDownI The function that handles the PointerDown event.
func (self *MSPointer) OnPointerDownI(args ...interface{}) {
    self.Object.Call("onPointerDown", args)
}

// OnPointerMove The function that handles the PointerMove event.
func (self *MSPointer) OnPointerMove(event *PointerEvent) {
    self.Object.Call("onPointerMove", event)
}

// OnPointerMoveI The function that handles the PointerMove event.
func (self *MSPointer) OnPointerMoveI(args ...interface{}) {
    self.Object.Call("onPointerMove", args)
}

// OnPointerUp The function that handles the PointerUp event.
func (self *MSPointer) OnPointerUp(event *PointerEvent) {
    self.Object.Call("onPointerUp", event)
}

// OnPointerUpI The function that handles the PointerUp event.
func (self *MSPointer) OnPointerUpI(args ...interface{}) {
    self.Object.Call("onPointerUp", args)
}

// OnPointerUpGlobal The internal method that handles the mouse up event from the window.
func (self *MSPointer) OnPointerUpGlobal(event *PointerEvent) {
    self.Object.Call("onPointerUpGlobal", event)
}

// OnPointerUpGlobalI The internal method that handles the mouse up event from the window.
func (self *MSPointer) OnPointerUpGlobalI(args ...interface{}) {
    self.Object.Call("onPointerUpGlobal", args)
}

// OnPointerOut The internal method that handles the pointer out event from the browser.
func (self *MSPointer) OnPointerOut(event *PointerEvent) {
    self.Object.Call("onPointerOut", event)
}

// OnPointerOutI The internal method that handles the pointer out event from the browser.
func (self *MSPointer) OnPointerOutI(args ...interface{}) {
    self.Object.Call("onPointerOut", args)
}

// Stop Stop the event listeners.
func (self *MSPointer) Stop() {
    self.Object.Call("stop")
}

// StopI Stop the event listeners.
func (self *MSPointer) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

