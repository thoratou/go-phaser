// Automatic generation for Phaser.Input
// generated file Input.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// Phaser.Input is the Input Manager for all types of Input across Phaser, including mouse, keyboard, touch and MSPointer.
// The Input manager is updated automatically by the core game loop.
type Input struct {
    *js.Object
}


// A reference to the currently running game.
func (self *Input) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *Input) SetGame(member Game) {
    self.Set("game", member)
}

// The canvas to which single pixels are drawn in order to perform pixel-perfect hit detection.
func (self *Input) GetHitCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("hitCanvas"))
}

// The canvas to which single pixels are drawn in order to perform pixel-perfect hit detection.
func (self *Input) SetHitCanvas(member dom.HTMLCanvasElement) {
    self.Set("hitCanvas", member)
}

// The context of the pixel perfect hit canvas.
func (self *Input) GetHitContext() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Get("hitContext"))
}

// The context of the pixel perfect hit canvas.
func (self *Input) SetHitContext(member dom.CanvasRenderingContext2D) {
    self.Set("hitContext", member)
}

// An array of callbacks that will be fired every time the activePointer receives a move event from the DOM.
// To add a callback to this array please use `Input.addMoveCallback`.
func (self *Input) GetMoveCallbacks() []interface{}{
	array := self.Get("moveCallbacks")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of callbacks that will be fired every time the activePointer receives a move event from the DOM.
// To add a callback to this array please use `Input.addMoveCallback`.
func (self *Input) SetMoveCallbacks(member []interface{}) {
    self.Set("moveCallbacks", member)
}

// How often should the input pointers be checked for updates? A value of 0 means every single frame (60fps); a value of 1 means every other frame (30fps) and so on.
func (self *Input) GetPollRate() float64{
    return self.Get("pollRate").Float()
}

// How often should the input pointers be checked for updates? A value of 0 means every single frame (60fps); a value of 1 means every other frame (30fps) and so on.
func (self *Input) SetPollRate(member float64) {
    self.Set("pollRate", member)
}

// When enabled, input (eg. Keyboard, Mouse, Touch) will be processed - as long as the individual sources are enabled themselves.
// 
// When not enabled, _all_ input sources are ignored. To disable just one type of input; for example, the Mouse, use `input.mouse.enabled = false`.
func (self *Input) GetEnabled() bool{
    return self.Get("enabled").Bool()
}

// When enabled, input (eg. Keyboard, Mouse, Touch) will be processed - as long as the individual sources are enabled themselves.
// 
// When not enabled, _all_ input sources are ignored. To disable just one type of input; for example, the Mouse, use `input.mouse.enabled = false`.
func (self *Input) SetEnabled(member bool) {
    self.Set("enabled", member)
}

// Controls the expected behavior when using a mouse and touch together on a multi-input device.
func (self *Input) GetMultiInputOverride() float64{
    return self.Get("multiInputOverride").Float()
}

// Controls the expected behavior when using a mouse and touch together on a multi-input device.
func (self *Input) SetMultiInputOverride(member float64) {
    self.Set("multiInputOverride", member)
}

// A point object representing the current position of the Pointer.
func (self *Input) GetPosition() Point{
    return Point{self.Get("position")}
}

// A point object representing the current position of the Pointer.
func (self *Input) SetPosition(member Point) {
    self.Set("position", member)
}

// A point object representing the speed of the Pointer. Only really useful in single Pointer games; otherwise see the Pointer objects directly.
func (self *Input) GetSpeed() Point{
    return Point{self.Get("speed")}
}

// A point object representing the speed of the Pointer. Only really useful in single Pointer games; otherwise see the Pointer objects directly.
func (self *Input) SetSpeed(member Point) {
    self.Set("speed", member)
}

// A Circle object centered on the x/y screen coordinates of the Input.
// Default size of 44px (Apples recommended "finger tip" size) but can be changed to anything.
func (self *Input) GetCircle() Circle{
    return Circle{self.Get("circle")}
}

// A Circle object centered on the x/y screen coordinates of the Input.
// Default size of 44px (Apples recommended "finger tip" size) but can be changed to anything.
func (self *Input) SetCircle(member Circle) {
    self.Set("circle", member)
}

// The scale by which all input coordinates are multiplied; calculated by the ScaleManager. In an un-scaled game the values will be x = 1 and y = 1.
func (self *Input) GetScale() Point{
    return Point{self.Get("scale")}
}

// The scale by which all input coordinates are multiplied; calculated by the ScaleManager. In an un-scaled game the values will be x = 1 and y = 1.
func (self *Input) SetScale(member Point) {
    self.Set("scale", member)
}

// The maximum number of Pointers allowed to be active at any one time. A value of -1 is only limited by the total number of pointers. For lots of games it's useful to set this to 1.
func (self *Input) GetMaxPointers() int{
    return self.Get("maxPointers").Int()
}

// The maximum number of Pointers allowed to be active at any one time. A value of -1 is only limited by the total number of pointers. For lots of games it's useful to set this to 1.
func (self *Input) SetMaxPointers(member int) {
    self.Set("maxPointers", member)
}

// The number of milliseconds that the Pointer has to be pressed down and then released to be considered a tap or click.
func (self *Input) GetTapRate() float64{
    return self.Get("tapRate").Float()
}

// The number of milliseconds that the Pointer has to be pressed down and then released to be considered a tap or click.
func (self *Input) SetTapRate(member float64) {
    self.Set("tapRate", member)
}

// The number of milliseconds between taps of the same Pointer for it to be considered a double tap / click.
func (self *Input) GetDoubleTapRate() float64{
    return self.Get("doubleTapRate").Float()
}

// The number of milliseconds between taps of the same Pointer for it to be considered a double tap / click.
func (self *Input) SetDoubleTapRate(member float64) {
    self.Set("doubleTapRate", member)
}

// The number of milliseconds that the Pointer has to be pressed down for it to fire a onHold event.
func (self *Input) GetHoldRate() float64{
    return self.Get("holdRate").Float()
}

// The number of milliseconds that the Pointer has to be pressed down for it to fire a onHold event.
func (self *Input) SetHoldRate(member float64) {
    self.Set("holdRate", member)
}

// The number of milliseconds below which the Pointer is considered justPressed.
func (self *Input) GetJustPressedRate() float64{
    return self.Get("justPressedRate").Float()
}

// The number of milliseconds below which the Pointer is considered justPressed.
func (self *Input) SetJustPressedRate(member float64) {
    self.Set("justPressedRate", member)
}

// The number of milliseconds below which the Pointer is considered justReleased .
func (self *Input) GetJustReleasedRate() float64{
    return self.Get("justReleasedRate").Float()
}

// The number of milliseconds below which the Pointer is considered justReleased .
func (self *Input) SetJustReleasedRate(member float64) {
    self.Set("justReleasedRate", member)
}

// Sets if the Pointer objects should record a history of x/y coordinates they have passed through.
// The history is cleared each time the Pointer is pressed down.
// The history is updated at the rate specified in Input.pollRate
func (self *Input) GetRecordPointerHistory() bool{
    return self.Get("recordPointerHistory").Bool()
}

// Sets if the Pointer objects should record a history of x/y coordinates they have passed through.
// The history is cleared each time the Pointer is pressed down.
// The history is updated at the rate specified in Input.pollRate
func (self *Input) SetRecordPointerHistory(member bool) {
    self.Set("recordPointerHistory", member)
}

// The rate in milliseconds at which the Pointer objects should update their tracking history.
func (self *Input) GetRecordRate() float64{
    return self.Get("recordRate").Float()
}

// The rate in milliseconds at which the Pointer objects should update their tracking history.
func (self *Input) SetRecordRate(member float64) {
    self.Set("recordRate", member)
}

// The total number of entries that can be recorded into the Pointer objects tracking history.
// If the Pointer is tracking one event every 100ms; then a trackLimit of 100 would store the last 10 seconds worth of history.
func (self *Input) GetRecordLimit() float64{
    return self.Get("recordLimit").Float()
}

// The total number of entries that can be recorded into the Pointer objects tracking history.
// If the Pointer is tracking one event every 100ms; then a trackLimit of 100 would store the last 10 seconds worth of history.
func (self *Input) SetRecordLimit(member float64) {
    self.Set("recordLimit", member)
}

// A Pointer object.
func (self *Input) GetPointer1() Pointer{
    return Pointer{self.Get("pointer1")}
}

// A Pointer object.
func (self *Input) SetPointer1(member Pointer) {
    self.Set("pointer1", member)
}

// A Pointer object.
func (self *Input) GetPointer2() Pointer{
    return Pointer{self.Get("pointer2")}
}

// A Pointer object.
func (self *Input) SetPointer2(member Pointer) {
    self.Set("pointer2", member)
}

// A Pointer object.
func (self *Input) GetPointer3() Pointer{
    return Pointer{self.Get("pointer3")}
}

// A Pointer object.
func (self *Input) SetPointer3(member Pointer) {
    self.Set("pointer3", member)
}

// A Pointer object.
func (self *Input) GetPointer4() Pointer{
    return Pointer{self.Get("pointer4")}
}

// A Pointer object.
func (self *Input) SetPointer4(member Pointer) {
    self.Set("pointer4", member)
}

// A Pointer object.
func (self *Input) GetPointer5() Pointer{
    return Pointer{self.Get("pointer5")}
}

// A Pointer object.
func (self *Input) SetPointer5(member Pointer) {
    self.Set("pointer5", member)
}

// A Pointer object.
func (self *Input) GetPointer6() Pointer{
    return Pointer{self.Get("pointer6")}
}

// A Pointer object.
func (self *Input) SetPointer6(member Pointer) {
    self.Set("pointer6", member)
}

// A Pointer object.
func (self *Input) GetPointer7() Pointer{
    return Pointer{self.Get("pointer7")}
}

// A Pointer object.
func (self *Input) SetPointer7(member Pointer) {
    self.Set("pointer7", member)
}

// A Pointer object.
func (self *Input) GetPointer8() Pointer{
    return Pointer{self.Get("pointer8")}
}

// A Pointer object.
func (self *Input) SetPointer8(member Pointer) {
    self.Set("pointer8", member)
}

// A Pointer object.
func (self *Input) GetPointer9() Pointer{
    return Pointer{self.Get("pointer9")}
}

// A Pointer object.
func (self *Input) SetPointer9(member Pointer) {
    self.Set("pointer9", member)
}

// A Pointer object.
func (self *Input) GetPointer10() Pointer{
    return Pointer{self.Get("pointer10")}
}

// A Pointer object.
func (self *Input) SetPointer10(member Pointer) {
    self.Set("pointer10", member)
}

// An array of non-mouse pointers that have been added to the game.
// The properties `pointer1..N` are aliases for `pointers[0..N-1]`.
func (self *Input) GetPointers() []Pointer{
	array := self.Get("pointers")
	length := array.Length()
	out := make([]Pointer, length, length)
	for i := 0; i < length; i++ {
		out[i] = Pointer{array.Index(i)}
	}
	return out
}

// An array of non-mouse pointers that have been added to the game.
// The properties `pointer1..N` are aliases for `pointers[0..N-1]`.
func (self *Input) SetPointers(member []Pointer) {
    self.Set("pointers", member)
}

// The most recently active Pointer object.
// 
// When you've limited max pointers to 1 this will accurately be either the first finger touched or mouse.
func (self *Input) GetActivePointer() Pointer{
    return Pointer{self.Get("activePointer")}
}

// The most recently active Pointer object.
// 
// When you've limited max pointers to 1 this will accurately be either the first finger touched or mouse.
func (self *Input) SetActivePointer(member Pointer) {
    self.Set("activePointer", member)
}

// The mouse has its own unique Phaser.Pointer object which you can use if making a desktop specific game.
func (self *Input) GetMousePointer() Pointer{
    return Pointer{self.Get("mousePointer")}
}

// The mouse has its own unique Phaser.Pointer object which you can use if making a desktop specific game.
func (self *Input) SetMousePointer(member Pointer) {
    self.Set("mousePointer", member)
}

// The Mouse Input manager.
// 
// You should not usually access this manager directly, but instead use Input.mousePointer or Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) GetMouse() Mouse{
    return Mouse{self.Get("mouse")}
}

// The Mouse Input manager.
// 
// You should not usually access this manager directly, but instead use Input.mousePointer or Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) SetMouse(member Mouse) {
    self.Set("mouse", member)
}

// The Keyboard Input manager.
func (self *Input) GetKeyboard() Keyboard{
    return Keyboard{self.Get("keyboard")}
}

// The Keyboard Input manager.
func (self *Input) SetKeyboard(member Keyboard) {
    self.Set("keyboard", member)
}

// The Touch Input manager.
// 
// You should not usually access this manager directly, but instead use Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) GetTouch() Touch{
    return Touch{self.Get("touch")}
}

// The Touch Input manager.
// 
// You should not usually access this manager directly, but instead use Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) SetTouch(member Touch) {
    self.Set("touch", member)
}

// The MSPointer Input manager.
// 
// You should not usually access this manager directly, but instead use Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) GetMspointer() MSPointer{
    return MSPointer{self.Get("mspointer")}
}

// The MSPointer Input manager.
// 
// You should not usually access this manager directly, but instead use Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) SetMspointer(member MSPointer) {
    self.Set("mspointer", member)
}

// The Gamepad Input manager.
func (self *Input) GetGamepad() Gamepad{
    return Gamepad{self.Get("gamepad")}
}

// The Gamepad Input manager.
func (self *Input) SetGamepad(member Gamepad) {
    self.Set("gamepad", member)
}

// If the Input Manager has been reset locked then all calls made to InputManager.reset, 
// such as from a State change, are ignored.
func (self *Input) GetResetLocked() bool{
    return self.Get("resetLocked").Bool()
}

// If the Input Manager has been reset locked then all calls made to InputManager.reset, 
// such as from a State change, are ignored.
func (self *Input) SetResetLocked(member bool) {
    self.Set("resetLocked", member)
}

// A Signal that is dispatched each time a pointer is pressed down.
func (self *Input) GetOnDown() Signal{
    return Signal{self.Get("onDown")}
}

// A Signal that is dispatched each time a pointer is pressed down.
func (self *Input) SetOnDown(member Signal) {
    self.Set("onDown", member)
}

// A Signal that is dispatched each time a pointer is released.
func (self *Input) GetOnUp() Signal{
    return Signal{self.Get("onUp")}
}

// A Signal that is dispatched each time a pointer is released.
func (self *Input) SetOnUp(member Signal) {
    self.Set("onUp", member)
}

// A Signal that is dispatched each time a pointer is tapped.
func (self *Input) GetOnTap() Signal{
    return Signal{self.Get("onTap")}
}

// A Signal that is dispatched each time a pointer is tapped.
func (self *Input) SetOnTap(member Signal) {
    self.Set("onTap", member)
}

// A Signal that is dispatched each time a pointer is held down.
func (self *Input) GetOnHold() Signal{
    return Signal{self.Get("onHold")}
}

// A Signal that is dispatched each time a pointer is held down.
func (self *Input) SetOnHold(member Signal) {
    self.Set("onHold", member)
}

// You can tell all Pointers to ignore any Game Object with a `priorityID` lower than this value.
// This is useful when stacking UI layers. Set to zero to disable.
func (self *Input) GetMinPriorityID() float64{
    return self.Get("minPriorityID").Float()
}

// You can tell all Pointers to ignore any Game Object with a `priorityID` lower than this value.
// This is useful when stacking UI layers. Set to zero to disable.
func (self *Input) SetMinPriorityID(member float64) {
    self.Set("minPriorityID", member)
}

// A list of interactive objects. The InputHandler components add and remove themselves from this list.
func (self *Input) GetInteractiveItems() ArraySet{
    return ArraySet{self.Get("interactiveItems")}
}

// A list of interactive objects. The InputHandler components add and remove themselves from this list.
func (self *Input) SetInteractiveItems(member ArraySet) {
    self.Set("interactiveItems", member)
}

// 
func (self *Input) GetMOUSE_OVERRIDES_TOUCH() float64{
    return self.Get("MOUSE_OVERRIDES_TOUCH").Float()
}

// 
func (self *Input) SetMOUSE_OVERRIDES_TOUCH(member float64) {
    self.Set("MOUSE_OVERRIDES_TOUCH", member)
}

// 
func (self *Input) GetTOUCH_OVERRIDES_MOUSE() float64{
    return self.Get("TOUCH_OVERRIDES_MOUSE").Float()
}

// 
func (self *Input) SetTOUCH_OVERRIDES_MOUSE(member float64) {
    self.Set("TOUCH_OVERRIDES_MOUSE", member)
}

// 
func (self *Input) GetMOUSE_TOUCH_COMBINE() float64{
    return self.Get("MOUSE_TOUCH_COMBINE").Float()
}

// 
func (self *Input) SetMOUSE_TOUCH_COMBINE(member float64) {
    self.Set("MOUSE_TOUCH_COMBINE", member)
}

// The maximum number of pointers that can be added. This excludes the mouse pointer.
func (self *Input) GetMAX_POINTERS() int{
    return self.Get("MAX_POINTERS").Int()
}

// The maximum number of pointers that can be added. This excludes the mouse pointer.
func (self *Input) SetMAX_POINTERS(member int) {
    self.Set("MAX_POINTERS", member)
}

// The X coordinate of the most recently active pointer.
// This value takes game scaling into account automatically. See Pointer.screenX/clientX for source values.
func (self *Input) GetX() float64{
    return self.Get("x").Float()
}

// The X coordinate of the most recently active pointer.
// This value takes game scaling into account automatically. See Pointer.screenX/clientX for source values.
func (self *Input) SetX(member float64) {
    self.Set("x", member)
}

// The Y coordinate of the most recently active pointer.
// This value takes game scaling into account automatically. See Pointer.screenY/clientY for source values.
func (self *Input) GetY() float64{
    return self.Get("y").Float()
}

// The Y coordinate of the most recently active pointer.
// This value takes game scaling into account automatically. See Pointer.screenY/clientY for source values.
func (self *Input) SetY(member float64) {
    self.Set("y", member)
}

// True if the Input is currently poll rate locked.
func (self *Input) GetPollLocked() bool{
    return self.Get("pollLocked").Bool()
}

// True if the Input is currently poll rate locked.
func (self *Input) SetPollLocked(member bool) {
    self.Set("pollLocked", member)
}

// The total number of inactive Pointers.
func (self *Input) GetTotalInactivePointers() float64{
    return self.Get("totalInactivePointers").Float()
}

// The total number of inactive Pointers.
func (self *Input) SetTotalInactivePointers(member float64) {
    self.Set("totalInactivePointers", member)
}

// The total number of active Pointers, not counting the mouse pointer.
func (self *Input) GetTotalActivePointers() int{
    return self.Get("totalActivePointers").Int()
}

// The total number of active Pointers, not counting the mouse pointer.
func (self *Input) SetTotalActivePointers(member int) {
    self.Set("totalActivePointers", member)
}

// The world X coordinate of the most recently active pointer.
func (self *Input) GetWorldX() float64{
    return self.Get("worldX").Float()
}

// The world X coordinate of the most recently active pointer.
func (self *Input) SetWorldX(member float64) {
    self.Set("worldX", member)
}

// The world Y coordinate of the most recently active pointer.
func (self *Input) GetWorldY() float64{
    return self.Get("worldY").Float()
}

// The world Y coordinate of the most recently active pointer.
func (self *Input) SetWorldY(member float64) {
    self.Set("worldY", member)
}



// Starts the Input Manager running.
func (self *Input) BootI(args ...interface{}) {
    self.Call("boot", args)
}

// Stops all of the Input Managers from running.
func (self *Input) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Adds a callback that is fired every time `Pointer.processInteractiveObjects` is called.
// The purpose of `processInteractiveObjects` is to work out which Game Object the Pointer is going to
// interact with. It works by polling all of the valid game objects, and then slowly discounting those
// that don't meet the criteria (i.e. they aren't under the Pointer, are disabled, invisible, etc).
// 
// Eventually a short-list of 'candidates' is created. These are all of the Game Objects which are valid
// for input and overlap with the Pointer. If you need fine-grained control over which of the items is
// selected then you can use this callback to do so.
// 
// The callback will be sent 3 parameters:
// 
// 1) A reference to the Phaser.Pointer object that is processing the Items.
// 2) An array containing all potential interactive candidates. This is an array of `InputHandler` objects, not Sprites.
// 3) The current 'favorite' candidate, based on its priorityID and position in the display list.
// 
// Your callback MUST return one of the candidates sent to it.
func (self *Input) SetInteractiveCandidateHandlerI(args ...interface{}) {
    self.Call("setInteractiveCandidateHandler", args)
}

// Adds a callback that is fired every time the activePointer receives a DOM move event such as a mousemove or touchmove.
// 
// The callback will be sent 4 parameters:
// 
// A reference to the Phaser.Pointer object that moved,
// The x position of the pointer,
// The y position,
// A boolean indicating if the movement was the result of a 'click' event (such as a mouse click or touch down).
// 
// It will be called every time the activePointer moves, which in a multi-touch game can be a lot of times, so this is best
// to only use if you've limited input to a single pointer (i.e. mouse or touch).
// 
// The callback is added to the Phaser.Input.moveCallbacks array and should be removed with Phaser.Input.deleteMoveCallback.
func (self *Input) AddMoveCallbackI(args ...interface{}) {
    self.Call("addMoveCallback", args)
}

// Removes the callback from the Phaser.Input.moveCallbacks array.
func (self *Input) DeleteMoveCallbackI(args ...interface{}) {
    self.Call("deleteMoveCallback", args)
}

// Add a new Pointer object to the Input Manager.
// By default Input creates 3 pointer objects: `mousePointer` (not include in part of general pointer pool), `pointer1` and `pointer2`.
// This method adds an additional pointer, up to a maximum of Phaser.Input.MAX_POINTERS (default of 10).
func (self *Input) AddPointerI(args ...interface{}) interface{}{
    return self.Call("addPointer", args)
}

// Updates the Input Manager. Called by the core Game loop.
func (self *Input) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Reset all of the Pointers and Input states.
// 
// The optional `hard` parameter will reset any events or callbacks that may be bound.
// Input.reset is called automatically during a State change or if a game loses focus / visibility.
// To control control the reset manually set {@link Phaser.InputManager.resetLocked} to `true`.
func (self *Input) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Resets the speed and old position properties.
func (self *Input) ResetSpeedI(args ...interface{}) {
    self.Call("resetSpeed", args)
}

// Find the first free Pointer object and start it, passing in the event data.
// This is called automatically by Phaser.Touch and Phaser.MSPointer.
func (self *Input) StartPointerI(args ...interface{}) Pointer{
    return Pointer{self.Call("startPointer", args)}
}

// Updates the matching Pointer object, passing in the event data.
// This is called automatically and should not normally need to be invoked.
func (self *Input) UpdatePointerI(args ...interface{}) Pointer{
    return Pointer{self.Call("updatePointer", args)}
}

// Stops the matching Pointer object, passing in the event data.
func (self *Input) StopPointerI(args ...interface{}) Pointer{
    return Pointer{self.Call("stopPointer", args)}
}

// Get the first Pointer with the given active state.
func (self *Input) GetPointerI(args ...interface{}) Pointer{
    return Pointer{self.Call("getPointer", args)}
}

// Get the Pointer object whos `identifier` property matches the given identifier value.
// 
// The identifier property is not set until the Pointer has been used at least once, as its populated by the DOM event.
// Also it can change every time you press the pointer down, and is not fixed once set.
// Note: Not all browsers set the identifier property and it's not part of the W3C spec, so you may need getPointerFromId instead.
func (self *Input) GetPointerFromIdentifierI(args ...interface{}) Pointer{
    return Pointer{self.Call("getPointerFromIdentifier", args)}
}

// Get the Pointer object whos `pointerId` property matches the given value.
// 
// The pointerId property is not set until the Pointer has been used at least once, as its populated by the DOM event.
// Also it can change every time you press the pointer down if the browser recycles it.
func (self *Input) GetPointerFromIdI(args ...interface{}) Pointer{
    return Pointer{self.Call("getPointerFromId", args)}
}

// This will return the local coordinates of the specified displayObject based on the given Pointer.
func (self *Input) GetLocalPositionI(args ...interface{}) Point{
    return Point{self.Call("getLocalPosition", args)}
}

// Tests if the pointer hits the given object.
func (self *Input) HitTestI(args ...interface{}) {
    self.Call("hitTest", args)
}

// Used for click trampolines. See {@link Phaser.Pointer.addClickTrampoline}.
func (self *Input) OnClickTrampolineI(args ...interface{}) {
    self.Call("onClickTrampoline", args)
}
