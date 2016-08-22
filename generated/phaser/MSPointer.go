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
func NewMSPointer(game *Game) *MSPointer {
    return &MSPointer{js.Global.Call("Phaser.MSPointer", game)}
}

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
func NewMSPointerI(args ...interface{}) *MSPointer {
    return &MSPointer{js.Global.Call("Phaser.MSPointer", args)}
}



// A reference to the currently running game.
func (self *MSPointer) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *MSPointer) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// A reference to the Phaser Input Manager.
func (self *MSPointer) GetInputA() *Input{
    return &Input{self.Object.Get("input")}
}

// A reference to the Phaser Input Manager.
func (self *MSPointer) SetInputA(member *Input) {
    self.Object.Set("input", member)
}

// The context under which callbacks are called (defaults to game).
func (self *MSPointer) GetCallbackContextA() interface{}{
    return self.Object.Get("callbackContext")
}

// The context under which callbacks are called (defaults to game).
func (self *MSPointer) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// A callback that can be fired on a MSPointerDown event.
func (self *MSPointer) SetPointerDownCallbackA(member func(...interface{})) {
    self.Object.Set("pointerDownCallback", member)
}

// A callback that can be fired on a MSPointerMove event.
func (self *MSPointer) SetPointerMoveCallbackA(member func(...interface{})) {
    self.Object.Set("pointerMoveCallback", member)
}

// A callback that can be fired on a MSPointerUp event.
func (self *MSPointer) SetPointerUpCallbackA(member func(...interface{})) {
    self.Object.Set("pointerUpCallback", member)
}

// If true the Pointer events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *MSPointer) GetCaptureA() bool{
    return self.Object.Get("capture").Bool()
}

// If true the Pointer events will have event.preventDefault applied to them, if false they will propagate fully.
func (self *MSPointer) SetCaptureA(member bool) {
    self.Object.Set("capture", member)
}

// This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *MSPointer) GetButtonA() int{
    return self.Object.Get("button").Int()
}

// This property was removed in Phaser 2.4 and should no longer be used.
// Instead please see the Pointer button properties such as `Pointer.leftButton`, `Pointer.rightButton` and so on.
// Or Pointer.button holds the DOM event button value if you require that.
func (self *MSPointer) SetButtonA(member int) {
    self.Object.Set("button", member)
}

// The browser MSPointer DOM event. Will be null if no event has ever been received.
// Access this property only inside a Pointer event handler and do not keep references to it.
func (self *MSPointer) GetEventA() interface{}{
    return self.Object.Get("event")
}

// The browser MSPointer DOM event. Will be null if no event has ever been received.
// Access this property only inside a Pointer event handler and do not keep references to it.
func (self *MSPointer) SetEventA(member interface{}) {
    self.Object.Set("event", member)
}

// MSPointer input will only be processed if enabled.
func (self *MSPointer) GetEnabledA() bool{
    return self.Object.Get("enabled").Bool()
}

// MSPointer input will only be processed if enabled.
func (self *MSPointer) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}



// Starts the event listeners running.
func (self *MSPointer) Start() {
    self.Object.Call("start")
}

// Starts the event listeners running.
func (self *MSPointer) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// The function that handles the PointerDown event.
func (self *MSPointer) OnPointerDown(event *PointerEvent) {
    self.Object.Call("onPointerDown", event)
}

// The function that handles the PointerDown event.
func (self *MSPointer) OnPointerDownI(args ...interface{}) {
    self.Object.Call("onPointerDown", args)
}

// The function that handles the PointerMove event.
func (self *MSPointer) OnPointerMove(event *PointerEvent) {
    self.Object.Call("onPointerMove", event)
}

// The function that handles the PointerMove event.
func (self *MSPointer) OnPointerMoveI(args ...interface{}) {
    self.Object.Call("onPointerMove", args)
}

// The function that handles the PointerUp event.
func (self *MSPointer) OnPointerUp(event *PointerEvent) {
    self.Object.Call("onPointerUp", event)
}

// The function that handles the PointerUp event.
func (self *MSPointer) OnPointerUpI(args ...interface{}) {
    self.Object.Call("onPointerUp", args)
}

// The internal method that handles the mouse up event from the window.
func (self *MSPointer) OnPointerUpGlobal(event *PointerEvent) {
    self.Object.Call("onPointerUpGlobal", event)
}

// The internal method that handles the mouse up event from the window.
func (self *MSPointer) OnPointerUpGlobalI(args ...interface{}) {
    self.Object.Call("onPointerUpGlobal", args)
}

// The internal method that handles the pointer out event from the browser.
func (self *MSPointer) OnPointerOut(event *PointerEvent) {
    self.Object.Call("onPointerOut", event)
}

// The internal method that handles the pointer out event from the browser.
func (self *MSPointer) OnPointerOutI(args ...interface{}) {
    self.Object.Call("onPointerOut", args)
}

// Stop the event listeners.
func (self *MSPointer) Stop() {
    self.Object.Call("stop")
}

// Stop the event listeners.
func (self *MSPointer) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}
