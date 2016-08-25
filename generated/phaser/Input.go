// Package phaser Automatic generation for Phaser.Input
// generated file Input.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// Input Phaser.Input is the Input Manager for all types of Input across Phaser, including mouse, keyboard, touch and MSPointer.
// The Input manager is updated automatically by the core game loop.
type Input struct {
    *js.Object
}

// NewInput Phaser.Input is the Input Manager for all types of Input across Phaser, including mouse, keyboard, touch and MSPointer.
// The Input manager is updated automatically by the core game loop.
func NewInput(game *Game) *Input {
    return &Input{js.Global.Get("Phaser").Get("Input").New(game)}
}
// NewInputI Phaser.Input is the Input Manager for all types of Input across Phaser, including mouse, keyboard, touch and MSPointer.
// The Input manager is updated automatically by the core game loop.
func NewInputI(args ...interface{}) *Input {
    return &Input{js.Global.Get("Phaser").Get("Input").New(args)}
}



// Input Binding conversion method to Input point 
func ToInput(jsStruct interface{}) *Input {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Input{Object: object}
	}
	return nil
}



// Game A reference to the currently running game.
func (self *Input) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *Input) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// HitCanvas The canvas to which single pixels are drawn in order to perform pixel-perfect hit detection.
func (self *Input) HitCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("hitCanvas"))
}

// SetHitCanvasA The canvas to which single pixels are drawn in order to perform pixel-perfect hit detection.
func (self *Input) SetHitCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("hitCanvas", member)
}

// HitContext The context of the pixel perfect hit canvas.
func (self *Input) HitContext() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("hitContext"))
}

// SetHitContextA The context of the pixel perfect hit canvas.
func (self *Input) SetHitContextA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("hitContext", member)
}

// MoveCallbacks An array of callbacks that will be fired every time the activePointer receives a move event from the DOM.
// To add a callback to this array please use `Input.addMoveCallback`.
func (self *Input) MoveCallbacks() []interface{}{
	array00 := self.Object.Get("moveCallbacks")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetMoveCallbacksA An array of callbacks that will be fired every time the activePointer receives a move event from the DOM.
// To add a callback to this array please use `Input.addMoveCallback`.
func (self *Input) SetMoveCallbacksA(member []interface{}) {
    self.Object.Set("moveCallbacks", member)
}

// PollRate How often should the input pointers be checked for updates? A value of 0 means every single frame (60fps); a value of 1 means every other frame (30fps) and so on.
func (self *Input) PollRate() int{
    return self.Object.Get("pollRate").Int()
}

// SetPollRateA How often should the input pointers be checked for updates? A value of 0 means every single frame (60fps); a value of 1 means every other frame (30fps) and so on.
func (self *Input) SetPollRateA(member int) {
    self.Object.Set("pollRate", member)
}

// Enabled When enabled, input (eg. Keyboard, Mouse, Touch) will be processed - as long as the individual sources are enabled themselves.
// 
// When not enabled, _all_ input sources are ignored. To disable just one type of input; for example, the Mouse, use `input.mouse.enabled = false`.
func (self *Input) Enabled() bool{
    return self.Object.Get("enabled").Bool()
}

// SetEnabledA When enabled, input (eg. Keyboard, Mouse, Touch) will be processed - as long as the individual sources are enabled themselves.
// 
// When not enabled, _all_ input sources are ignored. To disable just one type of input; for example, the Mouse, use `input.mouse.enabled = false`.
func (self *Input) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}

// MultiInputOverride Controls the expected behavior when using a mouse and touch together on a multi-input device.
func (self *Input) MultiInputOverride() int{
    return self.Object.Get("multiInputOverride").Int()
}

// SetMultiInputOverrideA Controls the expected behavior when using a mouse and touch together on a multi-input device.
func (self *Input) SetMultiInputOverrideA(member int) {
    self.Object.Set("multiInputOverride", member)
}

// Position A point object representing the current position of the Pointer.
func (self *Input) Position() *Point{
    return &Point{self.Object.Get("position")}
}

// SetPositionA A point object representing the current position of the Pointer.
func (self *Input) SetPositionA(member *Point) {
    self.Object.Set("position", member)
}

// Speed A point object representing the speed of the Pointer. Only really useful in single Pointer games; otherwise see the Pointer objects directly.
func (self *Input) Speed() *Point{
    return &Point{self.Object.Get("speed")}
}

// SetSpeedA A point object representing the speed of the Pointer. Only really useful in single Pointer games; otherwise see the Pointer objects directly.
func (self *Input) SetSpeedA(member *Point) {
    self.Object.Set("speed", member)
}

// Circle A Circle object centered on the x/y screen coordinates of the Input.
// Default size of 44px (Apples recommended "finger tip" size) but can be changed to anything.
func (self *Input) Circle() *Circle{
    return &Circle{self.Object.Get("circle")}
}

// SetCircleA A Circle object centered on the x/y screen coordinates of the Input.
// Default size of 44px (Apples recommended "finger tip" size) but can be changed to anything.
func (self *Input) SetCircleA(member *Circle) {
    self.Object.Set("circle", member)
}

// Scale The scale by which all input coordinates are multiplied; calculated by the ScaleManager. In an un-scaled game the values will be x = 1 and y = 1.
func (self *Input) Scale() *Point{
    return &Point{self.Object.Get("scale")}
}

// SetScaleA The scale by which all input coordinates are multiplied; calculated by the ScaleManager. In an un-scaled game the values will be x = 1 and y = 1.
func (self *Input) SetScaleA(member *Point) {
    self.Object.Set("scale", member)
}

// MaxPointers The maximum number of Pointers allowed to be active at any one time. A value of -1 is only limited by the total number of pointers. For lots of games it's useful to set this to 1.
func (self *Input) MaxPointers() int{
    return self.Object.Get("maxPointers").Int()
}

// SetMaxPointersA The maximum number of Pointers allowed to be active at any one time. A value of -1 is only limited by the total number of pointers. For lots of games it's useful to set this to 1.
func (self *Input) SetMaxPointersA(member int) {
    self.Object.Set("maxPointers", member)
}

// TapRate The number of milliseconds that the Pointer has to be pressed down and then released to be considered a tap or click.
func (self *Input) TapRate() int{
    return self.Object.Get("tapRate").Int()
}

// SetTapRateA The number of milliseconds that the Pointer has to be pressed down and then released to be considered a tap or click.
func (self *Input) SetTapRateA(member int) {
    self.Object.Set("tapRate", member)
}

// DoubleTapRate The number of milliseconds between taps of the same Pointer for it to be considered a double tap / click.
func (self *Input) DoubleTapRate() int{
    return self.Object.Get("doubleTapRate").Int()
}

// SetDoubleTapRateA The number of milliseconds between taps of the same Pointer for it to be considered a double tap / click.
func (self *Input) SetDoubleTapRateA(member int) {
    self.Object.Set("doubleTapRate", member)
}

// HoldRate The number of milliseconds that the Pointer has to be pressed down for it to fire a onHold event.
func (self *Input) HoldRate() int{
    return self.Object.Get("holdRate").Int()
}

// SetHoldRateA The number of milliseconds that the Pointer has to be pressed down for it to fire a onHold event.
func (self *Input) SetHoldRateA(member int) {
    self.Object.Set("holdRate", member)
}

// JustPressedRate The number of milliseconds below which the Pointer is considered justPressed.
func (self *Input) JustPressedRate() int{
    return self.Object.Get("justPressedRate").Int()
}

// SetJustPressedRateA The number of milliseconds below which the Pointer is considered justPressed.
func (self *Input) SetJustPressedRateA(member int) {
    self.Object.Set("justPressedRate", member)
}

// JustReleasedRate The number of milliseconds below which the Pointer is considered justReleased .
func (self *Input) JustReleasedRate() int{
    return self.Object.Get("justReleasedRate").Int()
}

// SetJustReleasedRateA The number of milliseconds below which the Pointer is considered justReleased .
func (self *Input) SetJustReleasedRateA(member int) {
    self.Object.Set("justReleasedRate", member)
}

// RecordPointerHistory Sets if the Pointer objects should record a history of x/y coordinates they have passed through.
// The history is cleared each time the Pointer is pressed down.
// The history is updated at the rate specified in Input.pollRate
func (self *Input) RecordPointerHistory() bool{
    return self.Object.Get("recordPointerHistory").Bool()
}

// SetRecordPointerHistoryA Sets if the Pointer objects should record a history of x/y coordinates they have passed through.
// The history is cleared each time the Pointer is pressed down.
// The history is updated at the rate specified in Input.pollRate
func (self *Input) SetRecordPointerHistoryA(member bool) {
    self.Object.Set("recordPointerHistory", member)
}

// RecordRate The rate in milliseconds at which the Pointer objects should update their tracking history.
func (self *Input) RecordRate() int{
    return self.Object.Get("recordRate").Int()
}

// SetRecordRateA The rate in milliseconds at which the Pointer objects should update their tracking history.
func (self *Input) SetRecordRateA(member int) {
    self.Object.Set("recordRate", member)
}

// RecordLimit The total number of entries that can be recorded into the Pointer objects tracking history.
// If the Pointer is tracking one event every 100ms; then a trackLimit of 100 would store the last 10 seconds worth of history.
func (self *Input) RecordLimit() int{
    return self.Object.Get("recordLimit").Int()
}

// SetRecordLimitA The total number of entries that can be recorded into the Pointer objects tracking history.
// If the Pointer is tracking one event every 100ms; then a trackLimit of 100 would store the last 10 seconds worth of history.
func (self *Input) SetRecordLimitA(member int) {
    self.Object.Set("recordLimit", member)
}

// Pointer1 A Pointer object.
func (self *Input) Pointer1() *Pointer{
    return &Pointer{self.Object.Get("pointer1")}
}

// SetPointer1A A Pointer object.
func (self *Input) SetPointer1A(member *Pointer) {
    self.Object.Set("pointer1", member)
}

// Pointer2 A Pointer object.
func (self *Input) Pointer2() *Pointer{
    return &Pointer{self.Object.Get("pointer2")}
}

// SetPointer2A A Pointer object.
func (self *Input) SetPointer2A(member *Pointer) {
    self.Object.Set("pointer2", member)
}

// Pointer3 A Pointer object.
func (self *Input) Pointer3() *Pointer{
    return &Pointer{self.Object.Get("pointer3")}
}

// SetPointer3A A Pointer object.
func (self *Input) SetPointer3A(member *Pointer) {
    self.Object.Set("pointer3", member)
}

// Pointer4 A Pointer object.
func (self *Input) Pointer4() *Pointer{
    return &Pointer{self.Object.Get("pointer4")}
}

// SetPointer4A A Pointer object.
func (self *Input) SetPointer4A(member *Pointer) {
    self.Object.Set("pointer4", member)
}

// Pointer5 A Pointer object.
func (self *Input) Pointer5() *Pointer{
    return &Pointer{self.Object.Get("pointer5")}
}

// SetPointer5A A Pointer object.
func (self *Input) SetPointer5A(member *Pointer) {
    self.Object.Set("pointer5", member)
}

// Pointer6 A Pointer object.
func (self *Input) Pointer6() *Pointer{
    return &Pointer{self.Object.Get("pointer6")}
}

// SetPointer6A A Pointer object.
func (self *Input) SetPointer6A(member *Pointer) {
    self.Object.Set("pointer6", member)
}

// Pointer7 A Pointer object.
func (self *Input) Pointer7() *Pointer{
    return &Pointer{self.Object.Get("pointer7")}
}

// SetPointer7A A Pointer object.
func (self *Input) SetPointer7A(member *Pointer) {
    self.Object.Set("pointer7", member)
}

// Pointer8 A Pointer object.
func (self *Input) Pointer8() *Pointer{
    return &Pointer{self.Object.Get("pointer8")}
}

// SetPointer8A A Pointer object.
func (self *Input) SetPointer8A(member *Pointer) {
    self.Object.Set("pointer8", member)
}

// Pointer9 A Pointer object.
func (self *Input) Pointer9() *Pointer{
    return &Pointer{self.Object.Get("pointer9")}
}

// SetPointer9A A Pointer object.
func (self *Input) SetPointer9A(member *Pointer) {
    self.Object.Set("pointer9", member)
}

// Pointer10 A Pointer object.
func (self *Input) Pointer10() *Pointer{
    return &Pointer{self.Object.Get("pointer10")}
}

// SetPointer10A A Pointer object.
func (self *Input) SetPointer10A(member *Pointer) {
    self.Object.Set("pointer10", member)
}

// Pointers An array of non-mouse pointers that have been added to the game.
// The properties `pointer1..N` are aliases for `pointers[0..N-1]`.
func (self *Input) Pointers() []Pointer{
	array00 := self.Object.Get("pointers")
	length00 := array00.Length()
	out00 := make([]Pointer, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = Pointer{array00.Index(i00)}
	}
	return out00
}

// SetPointersA An array of non-mouse pointers that have been added to the game.
// The properties `pointer1..N` are aliases for `pointers[0..N-1]`.
func (self *Input) SetPointersA(member []Pointer) {
    self.Object.Set("pointers", member)
}

// ActivePointer The most recently active Pointer object.
// 
// When you've limited max pointers to 1 this will accurately be either the first finger touched or mouse.
func (self *Input) ActivePointer() *Pointer{
    return &Pointer{self.Object.Get("activePointer")}
}

// SetActivePointerA The most recently active Pointer object.
// 
// When you've limited max pointers to 1 this will accurately be either the first finger touched or mouse.
func (self *Input) SetActivePointerA(member *Pointer) {
    self.Object.Set("activePointer", member)
}

// MousePointer The mouse has its own unique Phaser.Pointer object which you can use if making a desktop specific game.
func (self *Input) MousePointer() *Pointer{
    return &Pointer{self.Object.Get("mousePointer")}
}

// SetMousePointerA The mouse has its own unique Phaser.Pointer object which you can use if making a desktop specific game.
func (self *Input) SetMousePointerA(member *Pointer) {
    self.Object.Set("mousePointer", member)
}

// Mouse The Mouse Input manager.
// 
// You should not usually access this manager directly, but instead use Input.mousePointer or Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) Mouse() *Mouse{
    return &Mouse{self.Object.Get("mouse")}
}

// SetMouseA The Mouse Input manager.
// 
// You should not usually access this manager directly, but instead use Input.mousePointer or Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) SetMouseA(member *Mouse) {
    self.Object.Set("mouse", member)
}

// Keyboard The Keyboard Input manager.
func (self *Input) Keyboard() *Keyboard{
    return &Keyboard{self.Object.Get("keyboard")}
}

// SetKeyboardA The Keyboard Input manager.
func (self *Input) SetKeyboardA(member *Keyboard) {
    self.Object.Set("keyboard", member)
}

// Touch The Touch Input manager.
// 
// You should not usually access this manager directly, but instead use Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) Touch() *Touch{
    return &Touch{self.Object.Get("touch")}
}

// SetTouchA The Touch Input manager.
// 
// You should not usually access this manager directly, but instead use Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) SetTouchA(member *Touch) {
    self.Object.Set("touch", member)
}

// Mspointer The MSPointer Input manager.
// 
// You should not usually access this manager directly, but instead use Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) Mspointer() *MSPointer{
    return &MSPointer{self.Object.Get("mspointer")}
}

// SetMspointerA The MSPointer Input manager.
// 
// You should not usually access this manager directly, but instead use Input.activePointer 
// which normalizes all the input values for you, regardless of browser.
func (self *Input) SetMspointerA(member *MSPointer) {
    self.Object.Set("mspointer", member)
}

// Gamepad The Gamepad Input manager.
func (self *Input) Gamepad() *Gamepad{
    return &Gamepad{self.Object.Get("gamepad")}
}

// SetGamepadA The Gamepad Input manager.
func (self *Input) SetGamepadA(member *Gamepad) {
    self.Object.Set("gamepad", member)
}

// ResetLocked If the Input Manager has been reset locked then all calls made to InputManager.reset, 
// such as from a State change, are ignored.
func (self *Input) ResetLocked() bool{
    return self.Object.Get("resetLocked").Bool()
}

// SetResetLockedA If the Input Manager has been reset locked then all calls made to InputManager.reset, 
// such as from a State change, are ignored.
func (self *Input) SetResetLockedA(member bool) {
    self.Object.Set("resetLocked", member)
}

// OnDown A Signal that is dispatched each time a pointer is pressed down.
func (self *Input) OnDown() *Signal{
    return &Signal{self.Object.Get("onDown")}
}

// SetOnDownA A Signal that is dispatched each time a pointer is pressed down.
func (self *Input) SetOnDownA(member *Signal) {
    self.Object.Set("onDown", member)
}

// OnUp A Signal that is dispatched each time a pointer is released.
func (self *Input) OnUp() *Signal{
    return &Signal{self.Object.Get("onUp")}
}

// SetOnUpA A Signal that is dispatched each time a pointer is released.
func (self *Input) SetOnUpA(member *Signal) {
    self.Object.Set("onUp", member)
}

// OnTap A Signal that is dispatched each time a pointer is tapped.
func (self *Input) OnTap() *Signal{
    return &Signal{self.Object.Get("onTap")}
}

// SetOnTapA A Signal that is dispatched each time a pointer is tapped.
func (self *Input) SetOnTapA(member *Signal) {
    self.Object.Set("onTap", member)
}

// OnHold A Signal that is dispatched each time a pointer is held down.
func (self *Input) OnHold() *Signal{
    return &Signal{self.Object.Get("onHold")}
}

// SetOnHoldA A Signal that is dispatched each time a pointer is held down.
func (self *Input) SetOnHoldA(member *Signal) {
    self.Object.Set("onHold", member)
}

// MinPriorityID You can tell all Pointers to ignore any Game Object with a `priorityID` lower than this value.
// This is useful when stacking UI layers. Set to zero to disable.
func (self *Input) MinPriorityID() int{
    return self.Object.Get("minPriorityID").Int()
}

// SetMinPriorityIDA You can tell all Pointers to ignore any Game Object with a `priorityID` lower than this value.
// This is useful when stacking UI layers. Set to zero to disable.
func (self *Input) SetMinPriorityIDA(member int) {
    self.Object.Set("minPriorityID", member)
}

// InteractiveItems A list of interactive objects. The InputHandler components add and remove themselves from this list.
func (self *Input) InteractiveItems() *ArraySet{
    return &ArraySet{self.Object.Get("interactiveItems")}
}

// SetInteractiveItemsA A list of interactive objects. The InputHandler components add and remove themselves from this list.
func (self *Input) SetInteractiveItemsA(member *ArraySet) {
    self.Object.Set("interactiveItems", member)
}

// MOUSE_OVERRIDES_TOUCH empty description
func (self *Input) MOUSE_OVERRIDES_TOUCH() int{
    return self.Object.Get("MOUSE_OVERRIDES_TOUCH").Int()
}

// SetMOUSE_OVERRIDES_TOUCHA empty description
func (self *Input) SetMOUSE_OVERRIDES_TOUCHA(member int) {
    self.Object.Set("MOUSE_OVERRIDES_TOUCH", member)
}

// TOUCH_OVERRIDES_MOUSE empty description
func (self *Input) TOUCH_OVERRIDES_MOUSE() int{
    return self.Object.Get("TOUCH_OVERRIDES_MOUSE").Int()
}

// SetTOUCH_OVERRIDES_MOUSEA empty description
func (self *Input) SetTOUCH_OVERRIDES_MOUSEA(member int) {
    self.Object.Set("TOUCH_OVERRIDES_MOUSE", member)
}

// MOUSE_TOUCH_COMBINE empty description
func (self *Input) MOUSE_TOUCH_COMBINE() int{
    return self.Object.Get("MOUSE_TOUCH_COMBINE").Int()
}

// SetMOUSE_TOUCH_COMBINEA empty description
func (self *Input) SetMOUSE_TOUCH_COMBINEA(member int) {
    self.Object.Set("MOUSE_TOUCH_COMBINE", member)
}

// MAX_POINTERS The maximum number of pointers that can be added. This excludes the mouse pointer.
func (self *Input) MAX_POINTERS() int{
    return self.Object.Get("MAX_POINTERS").Int()
}

// SetMAX_POINTERSA The maximum number of pointers that can be added. This excludes the mouse pointer.
func (self *Input) SetMAX_POINTERSA(member int) {
    self.Object.Set("MAX_POINTERS", member)
}

// X The X coordinate of the most recently active pointer.
// This value takes game scaling into account automatically. See Pointer.screenX/clientX for source values.
func (self *Input) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The X coordinate of the most recently active pointer.
// This value takes game scaling into account automatically. See Pointer.screenX/clientX for source values.
func (self *Input) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The Y coordinate of the most recently active pointer.
// This value takes game scaling into account automatically. See Pointer.screenY/clientY for source values.
func (self *Input) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The Y coordinate of the most recently active pointer.
// This value takes game scaling into account automatically. See Pointer.screenY/clientY for source values.
func (self *Input) SetYA(member int) {
    self.Object.Set("y", member)
}

// PollLocked True if the Input is currently poll rate locked.
func (self *Input) PollLocked() bool{
    return self.Object.Get("pollLocked").Bool()
}

// SetPollLockedA True if the Input is currently poll rate locked.
func (self *Input) SetPollLockedA(member bool) {
    self.Object.Set("pollLocked", member)
}

// TotalInactivePointers The total number of inactive Pointers.
func (self *Input) TotalInactivePointers() int{
    return self.Object.Get("totalInactivePointers").Int()
}

// SetTotalInactivePointersA The total number of inactive Pointers.
func (self *Input) SetTotalInactivePointersA(member int) {
    self.Object.Set("totalInactivePointers", member)
}

// TotalActivePointers The total number of active Pointers, not counting the mouse pointer.
func (self *Input) TotalActivePointers() int{
    return self.Object.Get("totalActivePointers").Int()
}

// SetTotalActivePointersA The total number of active Pointers, not counting the mouse pointer.
func (self *Input) SetTotalActivePointersA(member int) {
    self.Object.Set("totalActivePointers", member)
}

// WorldX The world X coordinate of the most recently active pointer.
func (self *Input) WorldX() int{
    return self.Object.Get("worldX").Int()
}

// SetWorldXA The world X coordinate of the most recently active pointer.
func (self *Input) SetWorldXA(member int) {
    self.Object.Set("worldX", member)
}

// WorldY The world Y coordinate of the most recently active pointer.
func (self *Input) WorldY() int{
    return self.Object.Get("worldY").Int()
}

// SetWorldYA The world Y coordinate of the most recently active pointer.
func (self *Input) SetWorldYA(member int) {
    self.Object.Set("worldY", member)
}


// Boot Starts the Input Manager running.
func (self *Input) Boot() {
    self.Object.Call("boot")
}

// BootI Starts the Input Manager running.
func (self *Input) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// Destroy Stops all of the Input Managers from running.
func (self *Input) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Stops all of the Input Managers from running.
func (self *Input) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// SetInteractiveCandidateHandler Adds a callback that is fired every time `Pointer.processInteractiveObjects` is called.
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
func (self *Input) SetInteractiveCandidateHandler(callback interface{}, context interface{}) {
    self.Object.Call("setInteractiveCandidateHandler", callback, context)
}

// SetInteractiveCandidateHandlerI Adds a callback that is fired every time `Pointer.processInteractiveObjects` is called.
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
    self.Object.Call("setInteractiveCandidateHandler", args)
}

// AddMoveCallback Adds a callback that is fired every time the activePointer receives a DOM move event such as a mousemove or touchmove.
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
func (self *Input) AddMoveCallback(callback interface{}, context interface{}) {
    self.Object.Call("addMoveCallback", callback, context)
}

// AddMoveCallbackI Adds a callback that is fired every time the activePointer receives a DOM move event such as a mousemove or touchmove.
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
    self.Object.Call("addMoveCallback", args)
}

// DeleteMoveCallback Removes the callback from the Phaser.Input.moveCallbacks array.
func (self *Input) DeleteMoveCallback(callback interface{}, context interface{}) {
    self.Object.Call("deleteMoveCallback", callback, context)
}

// DeleteMoveCallbackI Removes the callback from the Phaser.Input.moveCallbacks array.
func (self *Input) DeleteMoveCallbackI(args ...interface{}) {
    self.Object.Call("deleteMoveCallback", args)
}

// AddPointer Add a new Pointer object to the Input Manager.
// By default Input creates 3 pointer objects: `mousePointer` (not include in part of general pointer pool), `pointer1` and `pointer2`.
// This method adds an additional pointer, up to a maximum of Phaser.Input.MAX_POINTERS (default of 10).
func (self *Input) AddPointer() interface{}{
    return self.Object.Call("addPointer")
}

// AddPointerI Add a new Pointer object to the Input Manager.
// By default Input creates 3 pointer objects: `mousePointer` (not include in part of general pointer pool), `pointer1` and `pointer2`.
// This method adds an additional pointer, up to a maximum of Phaser.Input.MAX_POINTERS (default of 10).
func (self *Input) AddPointerI(args ...interface{}) interface{}{
    return self.Object.Call("addPointer", args)
}

// Update Updates the Input Manager. Called by the core Game loop.
func (self *Input) Update() {
    self.Object.Call("update")
}

// UpdateI Updates the Input Manager. Called by the core Game loop.
func (self *Input) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Reset Reset all of the Pointers and Input states.
// 
// The optional `hard` parameter will reset any events or callbacks that may be bound.
// Input.reset is called automatically during a State change or if a game loses focus / visibility.
// To control control the reset manually set {@link Phaser.InputManager.resetLocked} to `true`.
func (self *Input) Reset() {
    self.Object.Call("reset")
}

// Reset1O Reset all of the Pointers and Input states.
// 
// The optional `hard` parameter will reset any events or callbacks that may be bound.
// Input.reset is called automatically during a State change or if a game loses focus / visibility.
// To control control the reset manually set {@link Phaser.InputManager.resetLocked} to `true`.
func (self *Input) Reset1O(hard bool) {
    self.Object.Call("reset", hard)
}

// ResetI Reset all of the Pointers and Input states.
// 
// The optional `hard` parameter will reset any events or callbacks that may be bound.
// Input.reset is called automatically during a State change or if a game loses focus / visibility.
// To control control the reset manually set {@link Phaser.InputManager.resetLocked} to `true`.
func (self *Input) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// ResetSpeed Resets the speed and old position properties.
func (self *Input) ResetSpeed(x int, y int) {
    self.Object.Call("resetSpeed", x, y)
}

// ResetSpeedI Resets the speed and old position properties.
func (self *Input) ResetSpeedI(args ...interface{}) {
    self.Object.Call("resetSpeed", args)
}

// StartPointer Find the first free Pointer object and start it, passing in the event data.
// This is called automatically by Phaser.Touch and Phaser.MSPointer.
func (self *Input) StartPointer(event interface{}) *Pointer{
    return &Pointer{self.Object.Call("startPointer", event)}
}

// StartPointerI Find the first free Pointer object and start it, passing in the event data.
// This is called automatically by Phaser.Touch and Phaser.MSPointer.
func (self *Input) StartPointerI(args ...interface{}) *Pointer{
    return &Pointer{self.Object.Call("startPointer", args)}
}

// UpdatePointer Updates the matching Pointer object, passing in the event data.
// This is called automatically and should not normally need to be invoked.
func (self *Input) UpdatePointer(event interface{}) *Pointer{
    return &Pointer{self.Object.Call("updatePointer", event)}
}

// UpdatePointerI Updates the matching Pointer object, passing in the event data.
// This is called automatically and should not normally need to be invoked.
func (self *Input) UpdatePointerI(args ...interface{}) *Pointer{
    return &Pointer{self.Object.Call("updatePointer", args)}
}

// StopPointer Stops the matching Pointer object, passing in the event data.
func (self *Input) StopPointer(event interface{}) *Pointer{
    return &Pointer{self.Object.Call("stopPointer", event)}
}

// StopPointerI Stops the matching Pointer object, passing in the event data.
func (self *Input) StopPointerI(args ...interface{}) *Pointer{
    return &Pointer{self.Object.Call("stopPointer", args)}
}

// GetPointer Get the first Pointer with the given active state.
func (self *Input) GetPointer() *Pointer{
    return &Pointer{self.Object.Call("getPointer")}
}

// GetPointer1O Get the first Pointer with the given active state.
func (self *Input) GetPointer1O(isActive bool) *Pointer{
    return &Pointer{self.Object.Call("getPointer", isActive)}
}

// GetPointerI Get the first Pointer with the given active state.
func (self *Input) GetPointerI(args ...interface{}) *Pointer{
    return &Pointer{self.Object.Call("getPointer", args)}
}

// GetPointerFromIdentifier Get the Pointer object whos `identifier` property matches the given identifier value.
// 
// The identifier property is not set until the Pointer has been used at least once, as its populated by the DOM event.
// Also it can change every time you press the pointer down, and is not fixed once set.
// Note: Not all browsers set the identifier property and it's not part of the W3C spec, so you may need getPointerFromId instead.
func (self *Input) GetPointerFromIdentifier(identifier int) *Pointer{
    return &Pointer{self.Object.Call("getPointerFromIdentifier", identifier)}
}

// GetPointerFromIdentifierI Get the Pointer object whos `identifier` property matches the given identifier value.
// 
// The identifier property is not set until the Pointer has been used at least once, as its populated by the DOM event.
// Also it can change every time you press the pointer down, and is not fixed once set.
// Note: Not all browsers set the identifier property and it's not part of the W3C spec, so you may need getPointerFromId instead.
func (self *Input) GetPointerFromIdentifierI(args ...interface{}) *Pointer{
    return &Pointer{self.Object.Call("getPointerFromIdentifier", args)}
}

// GetPointerFromId Get the Pointer object whos `pointerId` property matches the given value.
// 
// The pointerId property is not set until the Pointer has been used at least once, as its populated by the DOM event.
// Also it can change every time you press the pointer down if the browser recycles it.
func (self *Input) GetPointerFromId(pointerId int) *Pointer{
    return &Pointer{self.Object.Call("getPointerFromId", pointerId)}
}

// GetPointerFromIdI Get the Pointer object whos `pointerId` property matches the given value.
// 
// The pointerId property is not set until the Pointer has been used at least once, as its populated by the DOM event.
// Also it can change every time you press the pointer down if the browser recycles it.
func (self *Input) GetPointerFromIdI(args ...interface{}) *Pointer{
    return &Pointer{self.Object.Call("getPointerFromId", args)}
}

// GetLocalPosition This will return the local coordinates of the specified displayObject based on the given Pointer.
func (self *Input) GetLocalPosition(displayObject interface{}, pointer *Pointer) *Point{
    return &Point{self.Object.Call("getLocalPosition", displayObject, pointer)}
}

// GetLocalPositionI This will return the local coordinates of the specified displayObject based on the given Pointer.
func (self *Input) GetLocalPositionI(args ...interface{}) *Point{
    return &Point{self.Object.Call("getLocalPosition", args)}
}

// HitTest Tests if the pointer hits the given object.
func (self *Input) HitTest(displayObject *DisplayObject, pointer *Pointer, localPoint *Point) {
    self.Object.Call("hitTest", displayObject, pointer, localPoint)
}

// HitTestI Tests if the pointer hits the given object.
func (self *Input) HitTestI(args ...interface{}) {
    self.Object.Call("hitTest", args)
}

// OnClickTrampoline Used for click trampolines. See {@link Phaser.Pointer.addClickTrampoline}.
func (self *Input) OnClickTrampoline() {
    self.Object.Call("onClickTrampoline")
}

// OnClickTrampolineI Used for click trampolines. See {@link Phaser.Pointer.addClickTrampoline}.
func (self *Input) OnClickTrampolineI(args ...interface{}) {
    self.Object.Call("onClickTrampoline", args)
}

