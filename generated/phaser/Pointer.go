// Automatic generation for Phaser.Pointer
// generated file Pointer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A Pointer object is used by the Mouse, Touch and MSPoint managers and represents a single finger on the touch screen.
type Pointer struct {
    *js.Object
}


// A reference to the currently running game.
func (self *Pointer) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *Pointer) SetGame(member *Game) {
    self.Set("game", member)
}

// The ID of the Pointer object within the game. Each game can have up to 10 active pointers.
func (self *Pointer) GetId() int{
    return self.Get("id").Int()
}

// The ID of the Pointer object within the game. Each game can have up to 10 active pointers.
func (self *Pointer) SetId(member int) {
    self.Set("id", member)
}

// The const type of this object.
func (self *Pointer) GetType() int{
    return self.Get("type").Int()
}

// The const type of this object.
func (self *Pointer) SetType(member int) {
    self.Set("type", member)
}

// A Pointer object that exists is allowed to be checked for physics collisions and overlaps.
func (self *Pointer) GetExists() bool{
    return self.Get("exists").Bool()
}

// A Pointer object that exists is allowed to be checked for physics collisions and overlaps.
func (self *Pointer) SetExists(member bool) {
    self.Set("exists", member)
}

// The identifier property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) GetIdentifier() int{
    return self.Get("identifier").Int()
}

// The identifier property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) SetIdentifier(member int) {
    self.Set("identifier", member)
}

// The pointerId property of the Pointer as set by the DOM event when this Pointer is started. The browser can and will recycle this value.
func (self *Pointer) GetPointerId() int{
    return self.Get("pointerId").Int()
}

// The pointerId property of the Pointer as set by the DOM event when this Pointer is started. The browser can and will recycle this value.
func (self *Pointer) SetPointerId(member int) {
    self.Set("pointerId", member)
}

// The operational mode of this pointer.
func (self *Pointer) GetPointerMode() *PointerMode{
    return &PointerMode{self.Get("pointerMode")}
}

// The operational mode of this pointer.
func (self *Pointer) SetPointerMode(member *PointerMode) {
    self.Set("pointerMode", member)
}

// The target property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) GetTarget() interface{}{
    return self.Get("target")
}

// The target property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) SetTarget(member interface{}) {
    self.Set("target", member)
}

// The button property of the most recent DOM event when this Pointer is started.
// You should not rely on this value for accurate button detection, instead use the Pointer properties
// `leftButton`, `rightButton`, `middleButton` and so on.
func (self *Pointer) GetButton() interface{}{
    return self.Get("button")
}

// The button property of the most recent DOM event when this Pointer is started.
// You should not rely on this value for accurate button detection, instead use the Pointer properties
// `leftButton`, `rightButton`, `middleButton` and so on.
func (self *Pointer) SetButton(member interface{}) {
    self.Set("button", member)
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its left button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
func (self *Pointer) GetLeftButton() *DeviceButton{
    return &DeviceButton{self.Get("leftButton")}
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its left button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
func (self *Pointer) SetLeftButton(member *DeviceButton) {
    self.Set("leftButton", member)
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its middle button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) GetMiddleButton() *DeviceButton{
    return &DeviceButton{self.Get("middleButton")}
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its middle button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetMiddleButton(member *DeviceButton) {
    self.Set("middleButton", member)
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its right button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) GetRightButton() *DeviceButton{
    return &DeviceButton{self.Get("rightButton")}
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its right button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetRightButton(member *DeviceButton) {
    self.Set("rightButton", member)
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its X1 (back) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) GetBackButton() *DeviceButton{
    return &DeviceButton{self.Get("backButton")}
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its X1 (back) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetBackButton(member *DeviceButton) {
    self.Set("backButton", member)
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its X2 (forward) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) GetForwardButton() *DeviceButton{
    return &DeviceButton{self.Get("forwardButton")}
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its X2 (forward) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetForwardButton(member *DeviceButton) {
    self.Set("forwardButton", member)
}

// If this Pointer is a Pen / Stylus then you can access its eraser button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) GetEraserButton() *DeviceButton{
    return &DeviceButton{self.Get("eraserButton")}
}

// If this Pointer is a Pen / Stylus then you can access its eraser button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetEraserButton(member *DeviceButton) {
    self.Set("eraserButton", member)
}

// true if the Pointer is over the game canvas, otherwise false.
func (self *Pointer) GetWithinGame() bool{
    return self.Get("withinGame").Bool()
}

// true if the Pointer is over the game canvas, otherwise false.
func (self *Pointer) SetWithinGame(member bool) {
    self.Set("withinGame", member)
}

// The horizontal coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) GetClientX() int{
    return self.Get("clientX").Int()
}

// The horizontal coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) SetClientX(member int) {
    self.Set("clientX", member)
}

// The vertical coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) GetClientY() int{
    return self.Get("clientY").Int()
}

// The vertical coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) SetClientY(member int) {
    self.Set("clientY", member)
}

// The horizontal coordinate of the Pointer relative to whole document.
func (self *Pointer) GetPageX() int{
    return self.Get("pageX").Int()
}

// The horizontal coordinate of the Pointer relative to whole document.
func (self *Pointer) SetPageX(member int) {
    self.Set("pageX", member)
}

// The vertical coordinate of the Pointer relative to whole document.
func (self *Pointer) GetPageY() int{
    return self.Get("pageY").Int()
}

// The vertical coordinate of the Pointer relative to whole document.
func (self *Pointer) SetPageY(member int) {
    self.Set("pageY", member)
}

// The horizontal coordinate of the Pointer relative to the screen.
func (self *Pointer) GetScreenX() int{
    return self.Get("screenX").Int()
}

// The horizontal coordinate of the Pointer relative to the screen.
func (self *Pointer) SetScreenX(member int) {
    self.Set("screenX", member)
}

// The vertical coordinate of the Pointer relative to the screen.
func (self *Pointer) GetScreenY() int{
    return self.Get("screenY").Int()
}

// The vertical coordinate of the Pointer relative to the screen.
func (self *Pointer) SetScreenY(member int) {
    self.Set("screenY", member)
}

// The horizontal raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) GetRawMovementX() int{
    return self.Get("rawMovementX").Int()
}

// The horizontal raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetRawMovementX(member int) {
    self.Set("rawMovementX", member)
}

// The vertical raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) GetRawMovementY() int{
    return self.Get("rawMovementY").Int()
}

// The vertical raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetRawMovementY(member int) {
    self.Set("rawMovementY", member)
}

// The horizontal processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) GetMovementX() int{
    return self.Get("movementX").Int()
}

// The horizontal processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetMovementX(member int) {
    self.Set("movementX", member)
}

// The vertical processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) GetMovementY() int{
    return self.Get("movementY").Int()
}

// The vertical processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetMovementY(member int) {
    self.Set("movementY", member)
}

// The horizontal coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) GetX() int{
    return self.Get("x").Int()
}

// The horizontal coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) SetX(member int) {
    self.Set("x", member)
}

// The vertical coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) GetY() int{
    return self.Get("y").Int()
}

// The vertical coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) SetY(member int) {
    self.Set("y", member)
}

// If the Pointer is a mouse or pen / stylus this is true, otherwise false.
func (self *Pointer) GetIsMouse() bool{
    return self.Get("isMouse").Bool()
}

// If the Pointer is a mouse or pen / stylus this is true, otherwise false.
func (self *Pointer) SetIsMouse(member bool) {
    self.Set("isMouse", member)
}

// If the Pointer is touching the touchscreen, or *any* mouse or pen button is held down, isDown is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isDown.
func (self *Pointer) GetIsDown() bool{
    return self.Get("isDown").Bool()
}

// If the Pointer is touching the touchscreen, or *any* mouse or pen button is held down, isDown is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isDown.
func (self *Pointer) SetIsDown(member bool) {
    self.Set("isDown", member)
}

// If the Pointer is not touching the touchscreen, or *all* mouse or pen buttons are up, isUp is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isUp.
func (self *Pointer) GetIsUp() bool{
    return self.Get("isUp").Bool()
}

// If the Pointer is not touching the touchscreen, or *all* mouse or pen buttons are up, isUp is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isUp.
func (self *Pointer) SetIsUp(member bool) {
    self.Set("isUp", member)
}

// A timestamp representing when the Pointer first touched the touchscreen.
func (self *Pointer) GetTimeDown() int{
    return self.Get("timeDown").Int()
}

// A timestamp representing when the Pointer first touched the touchscreen.
func (self *Pointer) SetTimeDown(member int) {
    self.Set("timeDown", member)
}

// A timestamp representing when the Pointer left the touchscreen.
func (self *Pointer) GetTimeUp() int{
    return self.Get("timeUp").Int()
}

// A timestamp representing when the Pointer left the touchscreen.
func (self *Pointer) SetTimeUp(member int) {
    self.Set("timeUp", member)
}

// A timestamp representing when the Pointer was last tapped or clicked.
func (self *Pointer) GetPreviousTapTime() int{
    return self.Get("previousTapTime").Int()
}

// A timestamp representing when the Pointer was last tapped or clicked.
func (self *Pointer) SetPreviousTapTime(member int) {
    self.Set("previousTapTime", member)
}

// The total number of times this Pointer has been touched to the touchscreen.
func (self *Pointer) GetTotalTouches() int{
    return self.Get("totalTouches").Int()
}

// The total number of times this Pointer has been touched to the touchscreen.
func (self *Pointer) SetTotalTouches(member int) {
    self.Set("totalTouches", member)
}

// The number of milliseconds since the last click or touch event.
func (self *Pointer) GetMsSinceLastClick() int{
    return self.Get("msSinceLastClick").Int()
}

// The number of milliseconds since the last click or touch event.
func (self *Pointer) SetMsSinceLastClick(member int) {
    self.Set("msSinceLastClick", member)
}

// The Game Object this Pointer is currently over / touching / dragging.
func (self *Pointer) GetTargetObject() interface{}{
    return self.Get("targetObject")
}

// The Game Object this Pointer is currently over / touching / dragging.
func (self *Pointer) SetTargetObject(member interface{}) {
    self.Set("targetObject", member)
}

// This array is erased and re-populated every time this Pointer is updated. It contains references to all
// of the Game Objects that were considered as being valid for processing by this Pointer, this frame. To be
// valid they must have suitable a `priorityID`, be Input enabled, visible and actually have the Pointer over
// them. You can check the contents of this array in events such as `onInputDown`, but beware it is reset
// every frame.
func (self *Pointer) GetInteractiveCandidates() []interface{}{
	array := self.Get("interactiveCandidates")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// This array is erased and re-populated every time this Pointer is updated. It contains references to all
// of the Game Objects that were considered as being valid for processing by this Pointer, this frame. To be
// valid they must have suitable a `priorityID`, be Input enabled, visible and actually have the Pointer over
// them. You can check the contents of this array in events such as `onInputDown`, but beware it is reset
// every frame.
func (self *Pointer) SetInteractiveCandidates(member []interface{}) {
    self.Set("interactiveCandidates", member)
}

// An active pointer is one that is currently pressed down on the display. A Mouse is always active.
func (self *Pointer) GetActive() bool{
    return self.Get("active").Bool()
}

// An active pointer is one that is currently pressed down on the display. A Mouse is always active.
func (self *Pointer) SetActive(member bool) {
    self.Set("active", member)
}

// A dirty pointer needs to re-poll any interactive objects it may have been over, regardless if it has moved or not.
func (self *Pointer) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// A dirty pointer needs to re-poll any interactive objects it may have been over, regardless if it has moved or not.
func (self *Pointer) SetDirty(member bool) {
    self.Set("dirty", member)
}

// A Phaser.Point object containing the current x/y values of the pointer on the display.
func (self *Pointer) GetPosition() *Point{
    return &Point{self.Get("position")}
}

// A Phaser.Point object containing the current x/y values of the pointer on the display.
func (self *Pointer) SetPosition(member *Point) {
    self.Set("position", member)
}

// A Phaser.Point object containing the x/y values of the pointer when it was last in a down state on the display.
func (self *Pointer) GetPositionDown() *Point{
    return &Point{self.Get("positionDown")}
}

// A Phaser.Point object containing the x/y values of the pointer when it was last in a down state on the display.
func (self *Pointer) SetPositionDown(member *Point) {
    self.Set("positionDown", member)
}

// A Phaser.Point object containing the x/y values of the pointer when it was last released.
func (self *Pointer) GetPositionUp() *Point{
    return &Point{self.Get("positionUp")}
}

// A Phaser.Point object containing the x/y values of the pointer when it was last released.
func (self *Pointer) SetPositionUp(member *Point) {
    self.Set("positionUp", member)
}

// A Phaser.Circle that is centered on the x/y coordinates of this pointer, useful for hit detection.
// The Circle size is 44px (Apples recommended "finger tip" size).
func (self *Pointer) GetCircle() *Circle{
    return &Circle{self.Get("circle")}
}

// A Phaser.Circle that is centered on the x/y coordinates of this pointer, useful for hit detection.
// The Circle size is 44px (Apples recommended "finger tip" size).
func (self *Pointer) SetCircle(member *Circle) {
    self.Set("circle", member)
}

// No buttons at all.
func (self *Pointer) GetNO_BUTTON() int{
    return self.Get("NO_BUTTON").Int()
}

// No buttons at all.
func (self *Pointer) SetNO_BUTTON(member int) {
    self.Set("NO_BUTTON", member)
}

// The Left Mouse button, or in PointerEvent devices a Touch contact or Pen contact.
func (self *Pointer) GetLEFT_BUTTON() int{
    return self.Get("LEFT_BUTTON").Int()
}

// The Left Mouse button, or in PointerEvent devices a Touch contact or Pen contact.
func (self *Pointer) SetLEFT_BUTTON(member int) {
    self.Set("LEFT_BUTTON", member)
}

// The Right Mouse button, or in PointerEvent devices a Pen contact with a barrel button.
func (self *Pointer) GetRIGHT_BUTTON() int{
    return self.Get("RIGHT_BUTTON").Int()
}

// The Right Mouse button, or in PointerEvent devices a Pen contact with a barrel button.
func (self *Pointer) SetRIGHT_BUTTON(member int) {
    self.Set("RIGHT_BUTTON", member)
}

// The Middle Mouse button.
func (self *Pointer) GetMIDDLE_BUTTON() int{
    return self.Get("MIDDLE_BUTTON").Int()
}

// The Middle Mouse button.
func (self *Pointer) SetMIDDLE_BUTTON(member int) {
    self.Set("MIDDLE_BUTTON", member)
}

// The X1 button. This is typically the mouse Back button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) GetBACK_BUTTON() int{
    return self.Get("BACK_BUTTON").Int()
}

// The X1 button. This is typically the mouse Back button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) SetBACK_BUTTON(member int) {
    self.Set("BACK_BUTTON", member)
}

// The X2 button. This is typically the mouse Forward button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) GetFORWARD_BUTTON() int{
    return self.Get("FORWARD_BUTTON").Int()
}

// The X2 button. This is typically the mouse Forward button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) SetFORWARD_BUTTON(member int) {
    self.Set("FORWARD_BUTTON", member)
}

// The Eraser pen button on PointerEvent supported devices only.
func (self *Pointer) GetERASER_BUTTON() int{
    return self.Get("ERASER_BUTTON").Int()
}

// The Eraser pen button on PointerEvent supported devices only.
func (self *Pointer) SetERASER_BUTTON(member int) {
    self.Set("ERASER_BUTTON", member)
}

// How long the Pointer has been depressed on the touchscreen or *any* of the mouse buttons have been held down.
// If not currently down it returns -1.
// If you need to test a specific mouse or pen button then access the buttons directly, i.e. `Pointer.rightButton.duration`.
func (self *Pointer) GetDuration() int{
    return self.Get("duration").Int()
}

// How long the Pointer has been depressed on the touchscreen or *any* of the mouse buttons have been held down.
// If not currently down it returns -1.
// If you need to test a specific mouse or pen button then access the buttons directly, i.e. `Pointer.rightButton.duration`.
func (self *Pointer) SetDuration(member int) {
    self.Set("duration", member)
}

// Gets the X value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) GetWorldX() int{
    return self.Get("worldX").Int()
}

// Gets the X value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) SetWorldX(member int) {
    self.Set("worldX", member)
}

// Gets the Y value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) GetWorldY() int{
    return self.Get("worldY").Int()
}

// Gets the Y value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) SetWorldY(member int) {
    self.Set("worldY", member)
}



// Resets the states of all the button booleans.
func (self *Pointer) ResetButtonsI(args ...interface{}) {
    self.Call("resetButtons", args)
}

// Called by updateButtons.
func (self *Pointer) ProcessButtonsDownI(args ...interface{}) {
    self.Call("processButtonsDown", args)
}

// Called by updateButtons.
func (self *Pointer) ProcessButtonsUpI(args ...interface{}) {
    self.Call("processButtonsUp", args)
}

// Called when the event.buttons property changes from zero.
// Contains a button bitmask.
func (self *Pointer) UpdateButtonsI(args ...interface{}) {
    self.Call("updateButtons", args)
}

// Called when the Pointer is pressed onto the touchscreen.
func (self *Pointer) StartI(args ...interface{}) {
    self.Call("start", args)
}

// Called by the Input Manager.
func (self *Pointer) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Called when the Pointer is moved.
func (self *Pointer) MoveI(args ...interface{}) {
    self.Call("move", args)
}

// Process all interactive objects to find out which ones were updated in the recent Pointer move.
func (self *Pointer) ProcessInteractiveObjectsI(args ...interface{}) bool{
    return self.Call("processInteractiveObjects", args).Bool()
}

// This will change the `Pointer.targetObject` object to be the one provided.
// 
// This allows you to have fine-grained control over which object the Pointer is targeting.
// 
// Note that even if you set a new Target here, it is still able to be replaced by any other valid
// target during the next Pointer update.
func (self *Pointer) SwapTargetI(args ...interface{}) {
    self.Call("swapTarget", args)
}

// Called when the Pointer leaves the target area.
func (self *Pointer) LeaveI(args ...interface{}) {
    self.Call("leave", args)
}

// Called when the Pointer leaves the touchscreen.
func (self *Pointer) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// The Pointer is considered justPressed if the time it was pressed onto the touchscreen or clicked is less than justPressedRate.
// Note that calling justPressed doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was pressed down just once then see the Sprite.events.onInputDown event.
func (self *Pointer) JustPressedI(args ...interface{}) bool{
    return self.Call("justPressed", args).Bool()
}

// The Pointer is considered justReleased if the time it left the touchscreen is less than justReleasedRate.
// Note that calling justReleased doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was released just once then see the Sprite.events.onInputUp event.
func (self *Pointer) JustReleasedI(args ...interface{}) bool{
    return self.Call("justReleased", args).Bool()
}

// Add a click trampoline to this pointer.
// 
// A click trampoline is a callback that is run on the DOM 'click' event; this is primarily
// needed with certain browsers (ie. IE11) which restrict some actions like requestFullscreen
// to the DOM 'click' event and rejects it for 'pointer*' and 'mouse*' events.
// 
// This is used internally by the ScaleManager; click trampoline usage is uncommon.
// Click trampolines can only be added to pointers that are currently down.
func (self *Pointer) AddClickTrampolineI(args ...interface{}) {
    self.Call("addClickTrampoline", args)
}

// Fire all click trampolines for which the pointers are still referring to the registered object.
func (self *Pointer) ProcessClickTrampolinesI(args ...interface{}) {
    self.Call("processClickTrampolines", args)
}

// Resets the Pointer properties. Called by InputManager.reset when you perform a State change.
func (self *Pointer) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Resets the movementX and movementY properties. Use in your update handler after retrieving the values.
func (self *Pointer) ResetMovementI(args ...interface{}) {
    self.Call("resetMovement", args)
}
