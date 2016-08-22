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
func (self *Pointer) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *Pointer) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The ID of the Pointer object within the game. Each game can have up to 10 active pointers.
func (self *Pointer) GetIdA() int{
    return self.Object.Get("id").Int()
}

// The ID of the Pointer object within the game. Each game can have up to 10 active pointers.
func (self *Pointer) SetIdA(member int) {
    self.Object.Set("id", member)
}

// The const type of this object.
func (self *Pointer) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *Pointer) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// A Pointer object that exists is allowed to be checked for physics collisions and overlaps.
func (self *Pointer) GetExistsA() bool{
    return self.Object.Get("exists").Bool()
}

// A Pointer object that exists is allowed to be checked for physics collisions and overlaps.
func (self *Pointer) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// The identifier property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) GetIdentifierA() int{
    return self.Object.Get("identifier").Int()
}

// The identifier property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) SetIdentifierA(member int) {
    self.Object.Set("identifier", member)
}

// The pointerId property of the Pointer as set by the DOM event when this Pointer is started. The browser can and will recycle this value.
func (self *Pointer) GetPointerIdA() int{
    return self.Object.Get("pointerId").Int()
}

// The pointerId property of the Pointer as set by the DOM event when this Pointer is started. The browser can and will recycle this value.
func (self *Pointer) SetPointerIdA(member int) {
    self.Object.Set("pointerId", member)
}

// The operational mode of this pointer.
func (self *Pointer) GetPointerModeA() *PointerMode{
    return &PointerMode{self.Object.Get("pointerMode")}
}

// The operational mode of this pointer.
func (self *Pointer) SetPointerModeA(member *PointerMode) {
    self.Object.Set("pointerMode", member)
}

// The target property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) GetTargetA() interface{}{
    return self.Object.Get("target")
}

// The target property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) SetTargetA(member interface{}) {
    self.Object.Set("target", member)
}

// The button property of the most recent DOM event when this Pointer is started.
// You should not rely on this value for accurate button detection, instead use the Pointer properties
// `leftButton`, `rightButton`, `middleButton` and so on.
func (self *Pointer) GetButtonA() interface{}{
    return self.Object.Get("button")
}

// The button property of the most recent DOM event when this Pointer is started.
// You should not rely on this value for accurate button detection, instead use the Pointer properties
// `leftButton`, `rightButton`, `middleButton` and so on.
func (self *Pointer) SetButtonA(member interface{}) {
    self.Object.Set("button", member)
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its left button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
func (self *Pointer) GetLeftButtonA() *DeviceButton{
    return &DeviceButton{self.Object.Get("leftButton")}
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its left button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
func (self *Pointer) SetLeftButtonA(member *DeviceButton) {
    self.Object.Set("leftButton", member)
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its middle button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) GetMiddleButtonA() *DeviceButton{
    return &DeviceButton{self.Object.Get("middleButton")}
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its middle button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetMiddleButtonA(member *DeviceButton) {
    self.Object.Set("middleButton", member)
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its right button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) GetRightButtonA() *DeviceButton{
    return &DeviceButton{self.Object.Get("rightButton")}
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its right button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetRightButtonA(member *DeviceButton) {
    self.Object.Set("rightButton", member)
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its X1 (back) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) GetBackButtonA() *DeviceButton{
    return &DeviceButton{self.Object.Get("backButton")}
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its X1 (back) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetBackButtonA(member *DeviceButton) {
    self.Object.Set("backButton", member)
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its X2 (forward) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) GetForwardButtonA() *DeviceButton{
    return &DeviceButton{self.Object.Get("forwardButton")}
}

// If this Pointer is a Mouse or Pen / Stylus then you can access its X2 (forward) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetForwardButtonA(member *DeviceButton) {
    self.Object.Set("forwardButton", member)
}

// If this Pointer is a Pen / Stylus then you can access its eraser button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) GetEraserButtonA() *DeviceButton{
    return &DeviceButton{self.Object.Get("eraserButton")}
}

// If this Pointer is a Pen / Stylus then you can access its eraser button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetEraserButtonA(member *DeviceButton) {
    self.Object.Set("eraserButton", member)
}

// true if the Pointer is over the game canvas, otherwise false.
func (self *Pointer) GetWithinGameA() bool{
    return self.Object.Get("withinGame").Bool()
}

// true if the Pointer is over the game canvas, otherwise false.
func (self *Pointer) SetWithinGameA(member bool) {
    self.Object.Set("withinGame", member)
}

// The horizontal coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) GetClientXA() int{
    return self.Object.Get("clientX").Int()
}

// The horizontal coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) SetClientXA(member int) {
    self.Object.Set("clientX", member)
}

// The vertical coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) GetClientYA() int{
    return self.Object.Get("clientY").Int()
}

// The vertical coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) SetClientYA(member int) {
    self.Object.Set("clientY", member)
}

// The horizontal coordinate of the Pointer relative to whole document.
func (self *Pointer) GetPageXA() int{
    return self.Object.Get("pageX").Int()
}

// The horizontal coordinate of the Pointer relative to whole document.
func (self *Pointer) SetPageXA(member int) {
    self.Object.Set("pageX", member)
}

// The vertical coordinate of the Pointer relative to whole document.
func (self *Pointer) GetPageYA() int{
    return self.Object.Get("pageY").Int()
}

// The vertical coordinate of the Pointer relative to whole document.
func (self *Pointer) SetPageYA(member int) {
    self.Object.Set("pageY", member)
}

// The horizontal coordinate of the Pointer relative to the screen.
func (self *Pointer) GetScreenXA() int{
    return self.Object.Get("screenX").Int()
}

// The horizontal coordinate of the Pointer relative to the screen.
func (self *Pointer) SetScreenXA(member int) {
    self.Object.Set("screenX", member)
}

// The vertical coordinate of the Pointer relative to the screen.
func (self *Pointer) GetScreenYA() int{
    return self.Object.Get("screenY").Int()
}

// The vertical coordinate of the Pointer relative to the screen.
func (self *Pointer) SetScreenYA(member int) {
    self.Object.Set("screenY", member)
}

// The horizontal raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) GetRawMovementXA() int{
    return self.Object.Get("rawMovementX").Int()
}

// The horizontal raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetRawMovementXA(member int) {
    self.Object.Set("rawMovementX", member)
}

// The vertical raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) GetRawMovementYA() int{
    return self.Object.Get("rawMovementY").Int()
}

// The vertical raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetRawMovementYA(member int) {
    self.Object.Set("rawMovementY", member)
}

// The horizontal processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) GetMovementXA() int{
    return self.Object.Get("movementX").Int()
}

// The horizontal processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetMovementXA(member int) {
    self.Object.Set("movementX", member)
}

// The vertical processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) GetMovementYA() int{
    return self.Object.Get("movementY").Int()
}

// The vertical processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetMovementYA(member int) {
    self.Object.Set("movementY", member)
}

// The horizontal coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) GetXA() int{
    return self.Object.Get("x").Int()
}

// The horizontal coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) SetXA(member int) {
    self.Object.Set("x", member)
}

// The vertical coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) GetYA() int{
    return self.Object.Get("y").Int()
}

// The vertical coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) SetYA(member int) {
    self.Object.Set("y", member)
}

// If the Pointer is a mouse or pen / stylus this is true, otherwise false.
func (self *Pointer) GetIsMouseA() bool{
    return self.Object.Get("isMouse").Bool()
}

// If the Pointer is a mouse or pen / stylus this is true, otherwise false.
func (self *Pointer) SetIsMouseA(member bool) {
    self.Object.Set("isMouse", member)
}

// If the Pointer is touching the touchscreen, or *any* mouse or pen button is held down, isDown is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isDown.
func (self *Pointer) GetIsDownA() bool{
    return self.Object.Get("isDown").Bool()
}

// If the Pointer is touching the touchscreen, or *any* mouse or pen button is held down, isDown is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isDown.
func (self *Pointer) SetIsDownA(member bool) {
    self.Object.Set("isDown", member)
}

// If the Pointer is not touching the touchscreen, or *all* mouse or pen buttons are up, isUp is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isUp.
func (self *Pointer) GetIsUpA() bool{
    return self.Object.Get("isUp").Bool()
}

// If the Pointer is not touching the touchscreen, or *all* mouse or pen buttons are up, isUp is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isUp.
func (self *Pointer) SetIsUpA(member bool) {
    self.Object.Set("isUp", member)
}

// A timestamp representing when the Pointer first touched the touchscreen.
func (self *Pointer) GetTimeDownA() int{
    return self.Object.Get("timeDown").Int()
}

// A timestamp representing when the Pointer first touched the touchscreen.
func (self *Pointer) SetTimeDownA(member int) {
    self.Object.Set("timeDown", member)
}

// A timestamp representing when the Pointer left the touchscreen.
func (self *Pointer) GetTimeUpA() int{
    return self.Object.Get("timeUp").Int()
}

// A timestamp representing when the Pointer left the touchscreen.
func (self *Pointer) SetTimeUpA(member int) {
    self.Object.Set("timeUp", member)
}

// A timestamp representing when the Pointer was last tapped or clicked.
func (self *Pointer) GetPreviousTapTimeA() int{
    return self.Object.Get("previousTapTime").Int()
}

// A timestamp representing when the Pointer was last tapped or clicked.
func (self *Pointer) SetPreviousTapTimeA(member int) {
    self.Object.Set("previousTapTime", member)
}

// The total number of times this Pointer has been touched to the touchscreen.
func (self *Pointer) GetTotalTouchesA() int{
    return self.Object.Get("totalTouches").Int()
}

// The total number of times this Pointer has been touched to the touchscreen.
func (self *Pointer) SetTotalTouchesA(member int) {
    self.Object.Set("totalTouches", member)
}

// The number of milliseconds since the last click or touch event.
func (self *Pointer) GetMsSinceLastClickA() int{
    return self.Object.Get("msSinceLastClick").Int()
}

// The number of milliseconds since the last click or touch event.
func (self *Pointer) SetMsSinceLastClickA(member int) {
    self.Object.Set("msSinceLastClick", member)
}

// The Game Object this Pointer is currently over / touching / dragging.
func (self *Pointer) GetTargetObjectA() interface{}{
    return self.Object.Get("targetObject")
}

// The Game Object this Pointer is currently over / touching / dragging.
func (self *Pointer) SetTargetObjectA(member interface{}) {
    self.Object.Set("targetObject", member)
}

// This array is erased and re-populated every time this Pointer is updated. It contains references to all
// of the Game Objects that were considered as being valid for processing by this Pointer, this frame. To be
// valid they must have suitable a `priorityID`, be Input enabled, visible and actually have the Pointer over
// them. You can check the contents of this array in events such as `onInputDown`, but beware it is reset
// every frame.
func (self *Pointer) GetInteractiveCandidatesA() []interface{}{
	array00 := self.Object.Get("interactiveCandidates")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// This array is erased and re-populated every time this Pointer is updated. It contains references to all
// of the Game Objects that were considered as being valid for processing by this Pointer, this frame. To be
// valid they must have suitable a `priorityID`, be Input enabled, visible and actually have the Pointer over
// them. You can check the contents of this array in events such as `onInputDown`, but beware it is reset
// every frame.
func (self *Pointer) SetInteractiveCandidatesA(member []interface{}) {
    self.Object.Set("interactiveCandidates", member)
}

// An active pointer is one that is currently pressed down on the display. A Mouse is always active.
func (self *Pointer) GetActiveA() bool{
    return self.Object.Get("active").Bool()
}

// An active pointer is one that is currently pressed down on the display. A Mouse is always active.
func (self *Pointer) SetActiveA(member bool) {
    self.Object.Set("active", member)
}

// A dirty pointer needs to re-poll any interactive objects it may have been over, regardless if it has moved or not.
func (self *Pointer) GetDirtyA() bool{
    return self.Object.Get("dirty").Bool()
}

// A dirty pointer needs to re-poll any interactive objects it may have been over, regardless if it has moved or not.
func (self *Pointer) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// A Phaser.Point object containing the current x/y values of the pointer on the display.
func (self *Pointer) GetPositionA() *Point{
    return &Point{self.Object.Get("position")}
}

// A Phaser.Point object containing the current x/y values of the pointer on the display.
func (self *Pointer) SetPositionA(member *Point) {
    self.Object.Set("position", member)
}

// A Phaser.Point object containing the x/y values of the pointer when it was last in a down state on the display.
func (self *Pointer) GetPositionDownA() *Point{
    return &Point{self.Object.Get("positionDown")}
}

// A Phaser.Point object containing the x/y values of the pointer when it was last in a down state on the display.
func (self *Pointer) SetPositionDownA(member *Point) {
    self.Object.Set("positionDown", member)
}

// A Phaser.Point object containing the x/y values of the pointer when it was last released.
func (self *Pointer) GetPositionUpA() *Point{
    return &Point{self.Object.Get("positionUp")}
}

// A Phaser.Point object containing the x/y values of the pointer when it was last released.
func (self *Pointer) SetPositionUpA(member *Point) {
    self.Object.Set("positionUp", member)
}

// A Phaser.Circle that is centered on the x/y coordinates of this pointer, useful for hit detection.
// The Circle size is 44px (Apples recommended "finger tip" size).
func (self *Pointer) GetCircleA() *Circle{
    return &Circle{self.Object.Get("circle")}
}

// A Phaser.Circle that is centered on the x/y coordinates of this pointer, useful for hit detection.
// The Circle size is 44px (Apples recommended "finger tip" size).
func (self *Pointer) SetCircleA(member *Circle) {
    self.Object.Set("circle", member)
}

// No buttons at all.
func (self *Pointer) GetNO_BUTTONA() int{
    return self.Object.Get("NO_BUTTON").Int()
}

// No buttons at all.
func (self *Pointer) SetNO_BUTTONA(member int) {
    self.Object.Set("NO_BUTTON", member)
}

// The Left Mouse button, or in PointerEvent devices a Touch contact or Pen contact.
func (self *Pointer) GetLEFT_BUTTONA() int{
    return self.Object.Get("LEFT_BUTTON").Int()
}

// The Left Mouse button, or in PointerEvent devices a Touch contact or Pen contact.
func (self *Pointer) SetLEFT_BUTTONA(member int) {
    self.Object.Set("LEFT_BUTTON", member)
}

// The Right Mouse button, or in PointerEvent devices a Pen contact with a barrel button.
func (self *Pointer) GetRIGHT_BUTTONA() int{
    return self.Object.Get("RIGHT_BUTTON").Int()
}

// The Right Mouse button, or in PointerEvent devices a Pen contact with a barrel button.
func (self *Pointer) SetRIGHT_BUTTONA(member int) {
    self.Object.Set("RIGHT_BUTTON", member)
}

// The Middle Mouse button.
func (self *Pointer) GetMIDDLE_BUTTONA() int{
    return self.Object.Get("MIDDLE_BUTTON").Int()
}

// The Middle Mouse button.
func (self *Pointer) SetMIDDLE_BUTTONA(member int) {
    self.Object.Set("MIDDLE_BUTTON", member)
}

// The X1 button. This is typically the mouse Back button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) GetBACK_BUTTONA() int{
    return self.Object.Get("BACK_BUTTON").Int()
}

// The X1 button. This is typically the mouse Back button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) SetBACK_BUTTONA(member int) {
    self.Object.Set("BACK_BUTTON", member)
}

// The X2 button. This is typically the mouse Forward button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) GetFORWARD_BUTTONA() int{
    return self.Object.Get("FORWARD_BUTTON").Int()
}

// The X2 button. This is typically the mouse Forward button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) SetFORWARD_BUTTONA(member int) {
    self.Object.Set("FORWARD_BUTTON", member)
}

// The Eraser pen button on PointerEvent supported devices only.
func (self *Pointer) GetERASER_BUTTONA() int{
    return self.Object.Get("ERASER_BUTTON").Int()
}

// The Eraser pen button on PointerEvent supported devices only.
func (self *Pointer) SetERASER_BUTTONA(member int) {
    self.Object.Set("ERASER_BUTTON", member)
}

// How long the Pointer has been depressed on the touchscreen or *any* of the mouse buttons have been held down.
// If not currently down it returns -1.
// If you need to test a specific mouse or pen button then access the buttons directly, i.e. `Pointer.rightButton.duration`.
func (self *Pointer) GetDurationA() int{
    return self.Object.Get("duration").Int()
}

// How long the Pointer has been depressed on the touchscreen or *any* of the mouse buttons have been held down.
// If not currently down it returns -1.
// If you need to test a specific mouse or pen button then access the buttons directly, i.e. `Pointer.rightButton.duration`.
func (self *Pointer) SetDurationA(member int) {
    self.Object.Set("duration", member)
}

// Gets the X value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) GetWorldXA() int{
    return self.Object.Get("worldX").Int()
}

// Gets the X value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) SetWorldXA(member int) {
    self.Object.Set("worldX", member)
}

// Gets the Y value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) GetWorldYA() int{
    return self.Object.Get("worldY").Int()
}

// Gets the Y value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) SetWorldYA(member int) {
    self.Object.Set("worldY", member)
}



// Resets the states of all the button booleans.
func (self *Pointer) ResetButtons() {
    self.Object.Call("resetButtons")
}

// Resets the states of all the button booleans.
func (self *Pointer) ResetButtonsI(args ...interface{}) {
    self.Object.Call("resetButtons", args)
}

// Called by updateButtons.
func (self *Pointer) ProcessButtonsDown(buttons int, event *MouseEvent) {
    self.Object.Call("processButtonsDown", buttons, event)
}

// Called by updateButtons.
func (self *Pointer) ProcessButtonsDownI(args ...interface{}) {
    self.Object.Call("processButtonsDown", args)
}

// Called by updateButtons.
func (self *Pointer) ProcessButtonsUp(buttons int, event *MouseEvent) {
    self.Object.Call("processButtonsUp", buttons, event)
}

// Called by updateButtons.
func (self *Pointer) ProcessButtonsUpI(args ...interface{}) {
    self.Object.Call("processButtonsUp", args)
}

// Called when the event.buttons property changes from zero.
// Contains a button bitmask.
func (self *Pointer) UpdateButtons(event *MouseEvent) {
    self.Object.Call("updateButtons", event)
}

// Called when the event.buttons property changes from zero.
// Contains a button bitmask.
func (self *Pointer) UpdateButtonsI(args ...interface{}) {
    self.Object.Call("updateButtons", args)
}

// Called when the Pointer is pressed onto the touchscreen.
func (self *Pointer) Start(event interface{}) {
    self.Object.Call("start", event)
}

// Called when the Pointer is pressed onto the touchscreen.
func (self *Pointer) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Called by the Input Manager.
func (self *Pointer) Update() {
    self.Object.Call("update")
}

// Called by the Input Manager.
func (self *Pointer) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Called when the Pointer is moved.
func (self *Pointer) Move(event interface{}) {
    self.Object.Call("move", event)
}

// Called when the Pointer is moved.
func (self *Pointer) Move1O(event interface{}, fromClick bool) {
    self.Object.Call("move", event, fromClick)
}

// Called when the Pointer is moved.
func (self *Pointer) MoveI(args ...interface{}) {
    self.Object.Call("move", args)
}

// Process all interactive objects to find out which ones were updated in the recent Pointer move.
func (self *Pointer) ProcessInteractiveObjects() bool{
    return self.Object.Call("processInteractiveObjects").Bool()
}

// Process all interactive objects to find out which ones were updated in the recent Pointer move.
func (self *Pointer) ProcessInteractiveObjects1O(fromClick bool) bool{
    return self.Object.Call("processInteractiveObjects", fromClick).Bool()
}

// Process all interactive objects to find out which ones were updated in the recent Pointer move.
func (self *Pointer) ProcessInteractiveObjectsI(args ...interface{}) bool{
    return self.Object.Call("processInteractiveObjects", args).Bool()
}

// This will change the `Pointer.targetObject` object to be the one provided.
// 
// This allows you to have fine-grained control over which object the Pointer is targeting.
// 
// Note that even if you set a new Target here, it is still able to be replaced by any other valid
// target during the next Pointer update.
func (self *Pointer) SwapTarget(newTarget *InputHandler) {
    self.Object.Call("swapTarget", newTarget)
}

// This will change the `Pointer.targetObject` object to be the one provided.
// 
// This allows you to have fine-grained control over which object the Pointer is targeting.
// 
// Note that even if you set a new Target here, it is still able to be replaced by any other valid
// target during the next Pointer update.
func (self *Pointer) SwapTarget1O(newTarget *InputHandler, silent bool) {
    self.Object.Call("swapTarget", newTarget, silent)
}

// This will change the `Pointer.targetObject` object to be the one provided.
// 
// This allows you to have fine-grained control over which object the Pointer is targeting.
// 
// Note that even if you set a new Target here, it is still able to be replaced by any other valid
// target during the next Pointer update.
func (self *Pointer) SwapTargetI(args ...interface{}) {
    self.Object.Call("swapTarget", args)
}

// Called when the Pointer leaves the target area.
func (self *Pointer) Leave(event interface{}) {
    self.Object.Call("leave", event)
}

// Called when the Pointer leaves the target area.
func (self *Pointer) LeaveI(args ...interface{}) {
    self.Object.Call("leave", args)
}

// Called when the Pointer leaves the touchscreen.
func (self *Pointer) Stop(event interface{}) {
    self.Object.Call("stop", event)
}

// Called when the Pointer leaves the touchscreen.
func (self *Pointer) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// The Pointer is considered justPressed if the time it was pressed onto the touchscreen or clicked is less than justPressedRate.
// Note that calling justPressed doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was pressed down just once then see the Sprite.events.onInputDown event.
func (self *Pointer) JustPressed() bool{
    return self.Object.Call("justPressed").Bool()
}

// The Pointer is considered justPressed if the time it was pressed onto the touchscreen or clicked is less than justPressedRate.
// Note that calling justPressed doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was pressed down just once then see the Sprite.events.onInputDown event.
func (self *Pointer) JustPressed1O(duration int) bool{
    return self.Object.Call("justPressed", duration).Bool()
}

// The Pointer is considered justPressed if the time it was pressed onto the touchscreen or clicked is less than justPressedRate.
// Note that calling justPressed doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was pressed down just once then see the Sprite.events.onInputDown event.
func (self *Pointer) JustPressedI(args ...interface{}) bool{
    return self.Object.Call("justPressed", args).Bool()
}

// The Pointer is considered justReleased if the time it left the touchscreen is less than justReleasedRate.
// Note that calling justReleased doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was released just once then see the Sprite.events.onInputUp event.
func (self *Pointer) JustReleased() bool{
    return self.Object.Call("justReleased").Bool()
}

// The Pointer is considered justReleased if the time it left the touchscreen is less than justReleasedRate.
// Note that calling justReleased doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was released just once then see the Sprite.events.onInputUp event.
func (self *Pointer) JustReleased1O(duration int) bool{
    return self.Object.Call("justReleased", duration).Bool()
}

// The Pointer is considered justReleased if the time it left the touchscreen is less than justReleasedRate.
// Note that calling justReleased doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was released just once then see the Sprite.events.onInputUp event.
func (self *Pointer) JustReleasedI(args ...interface{}) bool{
    return self.Object.Call("justReleased", args).Bool()
}

// Add a click trampoline to this pointer.
// 
// A click trampoline is a callback that is run on the DOM 'click' event; this is primarily
// needed with certain browsers (ie. IE11) which restrict some actions like requestFullscreen
// to the DOM 'click' event and rejects it for 'pointer*' and 'mouse*' events.
// 
// This is used internally by the ScaleManager; click trampoline usage is uncommon.
// Click trampolines can only be added to pointers that are currently down.
func (self *Pointer) AddClickTrampoline(name string, callback func(...interface{}), callbackContext interface{}, callbackArgs interface{}) {
    self.Object.Call("addClickTrampoline", name, callback, callbackContext, callbackArgs)
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
    self.Object.Call("addClickTrampoline", args)
}

// Fire all click trampolines for which the pointers are still referring to the registered object.
func (self *Pointer) ProcessClickTrampolines() {
    self.Object.Call("processClickTrampolines")
}

// Fire all click trampolines for which the pointers are still referring to the registered object.
func (self *Pointer) ProcessClickTrampolinesI(args ...interface{}) {
    self.Object.Call("processClickTrampolines", args)
}

// Resets the Pointer properties. Called by InputManager.reset when you perform a State change.
func (self *Pointer) Reset() {
    self.Object.Call("reset")
}

// Resets the Pointer properties. Called by InputManager.reset when you perform a State change.
func (self *Pointer) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Resets the movementX and movementY properties. Use in your update handler after retrieving the values.
func (self *Pointer) ResetMovement() {
    self.Object.Call("resetMovement")
}

// Resets the movementX and movementY properties. Use in your update handler after retrieving the values.
func (self *Pointer) ResetMovementI(args ...interface{}) {
    self.Object.Call("resetMovement", args)
}
