// Package phaser Automatic generation for Phaser.Pointer
// generated file Pointer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Pointer A Pointer object is used by the Mouse, Touch and MSPoint managers and represents a single finger on the touch screen.
type Pointer struct {
    *js.Object
}

// NewPointer A Pointer object is used by the Mouse, Touch and MSPoint managers and represents a single finger on the touch screen.
func NewPointer(game *Game, id int, pointerMode *PointerMode) *Pointer {
    return &Pointer{js.Global.Get("Phaser").Get("Pointer").New(game, id, pointerMode)}
}
// NewPointerI A Pointer object is used by the Mouse, Touch and MSPoint managers and represents a single finger on the touch screen.
func NewPointerI(args ...interface{}) *Pointer {
    return &Pointer{js.Global.Get("Phaser").Get("Pointer").New(args)}
}



// Pointer Binding conversion method to Pointer point 
func ToPointer(jsStruct interface{}) *Pointer {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Pointer{Object: object}
	}
	return nil
}



// Game A reference to the currently running game.
func (self *Pointer) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *Pointer) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Id The ID of the Pointer object within the game. Each game can have up to 10 active pointers.
func (self *Pointer) Id() int{
    return self.Object.Get("id").Int()
}

// SetIdA The ID of the Pointer object within the game. Each game can have up to 10 active pointers.
func (self *Pointer) SetIdA(member int) {
    self.Object.Set("id", member)
}

// Type The const type of this object.
func (self *Pointer) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *Pointer) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Exists A Pointer object that exists is allowed to be checked for physics collisions and overlaps.
func (self *Pointer) Exists() bool{
    return self.Object.Get("exists").Bool()
}

// SetExistsA A Pointer object that exists is allowed to be checked for physics collisions and overlaps.
func (self *Pointer) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// Identifier The identifier property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) Identifier() int{
    return self.Object.Get("identifier").Int()
}

// SetIdentifierA The identifier property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) SetIdentifierA(member int) {
    self.Object.Set("identifier", member)
}

// PointerId The pointerId property of the Pointer as set by the DOM event when this Pointer is started. The browser can and will recycle this value.
func (self *Pointer) PointerId() int{
    return self.Object.Get("pointerId").Int()
}

// SetPointerIdA The pointerId property of the Pointer as set by the DOM event when this Pointer is started. The browser can and will recycle this value.
func (self *Pointer) SetPointerIdA(member int) {
    self.Object.Set("pointerId", member)
}

// PointerMode The operational mode of this pointer.
func (self *Pointer) PointerMode() *PointerMode{
    return &PointerMode{self.Object.Get("pointerMode")}
}

// SetPointerModeA The operational mode of this pointer.
func (self *Pointer) SetPointerModeA(member *PointerMode) {
    self.Object.Set("pointerMode", member)
}

// Target The target property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) Target() interface{}{
    return self.Object.Get("target")
}

// SetTargetA The target property of the Pointer as set by the DOM event when this Pointer is started.
func (self *Pointer) SetTargetA(member interface{}) {
    self.Object.Set("target", member)
}

// Button The button property of the most recent DOM event when this Pointer is started.
// You should not rely on this value for accurate button detection, instead use the Pointer properties
// `leftButton`, `rightButton`, `middleButton` and so on.
func (self *Pointer) Button() interface{}{
    return self.Object.Get("button")
}

// SetButtonA The button property of the most recent DOM event when this Pointer is started.
// You should not rely on this value for accurate button detection, instead use the Pointer properties
// `leftButton`, `rightButton`, `middleButton` and so on.
func (self *Pointer) SetButtonA(member interface{}) {
    self.Object.Set("button", member)
}

// LeftButton If this Pointer is a Mouse or Pen / Stylus then you can access its left button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
func (self *Pointer) LeftButton() *DeviceButton{
    return &DeviceButton{self.Object.Get("leftButton")}
}

// SetLeftButtonA If this Pointer is a Mouse or Pen / Stylus then you can access its left button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
func (self *Pointer) SetLeftButtonA(member *DeviceButton) {
    self.Object.Set("leftButton", member)
}

// MiddleButton If this Pointer is a Mouse or Pen / Stylus then you can access its middle button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) MiddleButton() *DeviceButton{
    return &DeviceButton{self.Object.Get("middleButton")}
}

// SetMiddleButtonA If this Pointer is a Mouse or Pen / Stylus then you can access its middle button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetMiddleButtonA(member *DeviceButton) {
    self.Object.Set("middleButton", member)
}

// RightButton If this Pointer is a Mouse or Pen / Stylus then you can access its right button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) RightButton() *DeviceButton{
    return &DeviceButton{self.Object.Get("rightButton")}
}

// SetRightButtonA If this Pointer is a Mouse or Pen / Stylus then you can access its right button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetRightButtonA(member *DeviceButton) {
    self.Object.Set("rightButton", member)
}

// BackButton If this Pointer is a Mouse or Pen / Stylus then you can access its X1 (back) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) BackButton() *DeviceButton{
    return &DeviceButton{self.Object.Get("backButton")}
}

// SetBackButtonA If this Pointer is a Mouse or Pen / Stylus then you can access its X1 (back) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetBackButtonA(member *DeviceButton) {
    self.Object.Set("backButton", member)
}

// ForwardButton If this Pointer is a Mouse or Pen / Stylus then you can access its X2 (forward) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) ForwardButton() *DeviceButton{
    return &DeviceButton{self.Object.Get("forwardButton")}
}

// SetForwardButtonA If this Pointer is a Mouse or Pen / Stylus then you can access its X2 (forward) button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetForwardButtonA(member *DeviceButton) {
    self.Object.Set("forwardButton", member)
}

// EraserButton If this Pointer is a Pen / Stylus then you can access its eraser button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) EraserButton() *DeviceButton{
    return &DeviceButton{self.Object.Get("eraserButton")}
}

// SetEraserButtonA If this Pointer is a Pen / Stylus then you can access its eraser button directly through this property.
// 
// The DeviceButton has its own properties such as `isDown`, `duration` and methods like `justReleased` for more fine-grained
// button control.
// 
// Please see the DeviceButton docs for details on browser button limitations.
func (self *Pointer) SetEraserButtonA(member *DeviceButton) {
    self.Object.Set("eraserButton", member)
}

// WithinGame true if the Pointer is over the game canvas, otherwise false.
func (self *Pointer) WithinGame() bool{
    return self.Object.Get("withinGame").Bool()
}

// SetWithinGameA true if the Pointer is over the game canvas, otherwise false.
func (self *Pointer) SetWithinGameA(member bool) {
    self.Object.Set("withinGame", member)
}

// ClientX The horizontal coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) ClientX() int{
    return self.Object.Get("clientX").Int()
}

// SetClientXA The horizontal coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) SetClientXA(member int) {
    self.Object.Set("clientX", member)
}

// ClientY The vertical coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) ClientY() int{
    return self.Object.Get("clientY").Int()
}

// SetClientYA The vertical coordinate of the Pointer within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (self *Pointer) SetClientYA(member int) {
    self.Object.Set("clientY", member)
}

// PageX The horizontal coordinate of the Pointer relative to whole document.
func (self *Pointer) PageX() int{
    return self.Object.Get("pageX").Int()
}

// SetPageXA The horizontal coordinate of the Pointer relative to whole document.
func (self *Pointer) SetPageXA(member int) {
    self.Object.Set("pageX", member)
}

// PageY The vertical coordinate of the Pointer relative to whole document.
func (self *Pointer) PageY() int{
    return self.Object.Get("pageY").Int()
}

// SetPageYA The vertical coordinate of the Pointer relative to whole document.
func (self *Pointer) SetPageYA(member int) {
    self.Object.Set("pageY", member)
}

// ScreenX The horizontal coordinate of the Pointer relative to the screen.
func (self *Pointer) ScreenX() int{
    return self.Object.Get("screenX").Int()
}

// SetScreenXA The horizontal coordinate of the Pointer relative to the screen.
func (self *Pointer) SetScreenXA(member int) {
    self.Object.Set("screenX", member)
}

// ScreenY The vertical coordinate of the Pointer relative to the screen.
func (self *Pointer) ScreenY() int{
    return self.Object.Get("screenY").Int()
}

// SetScreenYA The vertical coordinate of the Pointer relative to the screen.
func (self *Pointer) SetScreenYA(member int) {
    self.Object.Set("screenY", member)
}

// RawMovementX The horizontal raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) RawMovementX() int{
    return self.Object.Get("rawMovementX").Int()
}

// SetRawMovementXA The horizontal raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetRawMovementXA(member int) {
    self.Object.Set("rawMovementX", member)
}

// RawMovementY The vertical raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) RawMovementY() int{
    return self.Object.Get("rawMovementY").Int()
}

// SetRawMovementYA The vertical raw relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetRawMovementYA(member int) {
    self.Object.Set("rawMovementY", member)
}

// MovementX The horizontal processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) MovementX() int{
    return self.Object.Get("movementX").Int()
}

// SetMovementXA The horizontal processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetMovementXA(member int) {
    self.Object.Set("movementX", member)
}

// MovementY The vertical processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) MovementY() int{
    return self.Object.Get("movementY").Int()
}

// SetMovementYA The vertical processed relative movement of the Pointer in pixels since last event.
func (self *Pointer) SetMovementYA(member int) {
    self.Object.Set("movementY", member)
}

// X The horizontal coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The horizontal coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The vertical coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The vertical coordinate of the Pointer. This value is automatically scaled based on the game scale.
func (self *Pointer) SetYA(member int) {
    self.Object.Set("y", member)
}

// IsMouse If the Pointer is a mouse or pen / stylus this is true, otherwise false.
func (self *Pointer) IsMouse() bool{
    return self.Object.Get("isMouse").Bool()
}

// SetIsMouseA If the Pointer is a mouse or pen / stylus this is true, otherwise false.
func (self *Pointer) SetIsMouseA(member bool) {
    self.Object.Set("isMouse", member)
}

// IsDown If the Pointer is touching the touchscreen, or *any* mouse or pen button is held down, isDown is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isDown.
func (self *Pointer) IsDown() bool{
    return self.Object.Get("isDown").Bool()
}

// SetIsDownA If the Pointer is touching the touchscreen, or *any* mouse or pen button is held down, isDown is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isDown.
func (self *Pointer) SetIsDownA(member bool) {
    self.Object.Set("isDown", member)
}

// IsUp If the Pointer is not touching the touchscreen, or *all* mouse or pen buttons are up, isUp is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isUp.
func (self *Pointer) IsUp() bool{
    return self.Object.Get("isUp").Bool()
}

// SetIsUpA If the Pointer is not touching the touchscreen, or *all* mouse or pen buttons are up, isUp is set to true.
// If you need to check a specific mouse or pen button then use the button properties, i.e. Pointer.rightButton.isUp.
func (self *Pointer) SetIsUpA(member bool) {
    self.Object.Set("isUp", member)
}

// TimeDown A timestamp representing when the Pointer first touched the touchscreen.
func (self *Pointer) TimeDown() int{
    return self.Object.Get("timeDown").Int()
}

// SetTimeDownA A timestamp representing when the Pointer first touched the touchscreen.
func (self *Pointer) SetTimeDownA(member int) {
    self.Object.Set("timeDown", member)
}

// TimeUp A timestamp representing when the Pointer left the touchscreen.
func (self *Pointer) TimeUp() int{
    return self.Object.Get("timeUp").Int()
}

// SetTimeUpA A timestamp representing when the Pointer left the touchscreen.
func (self *Pointer) SetTimeUpA(member int) {
    self.Object.Set("timeUp", member)
}

// PreviousTapTime A timestamp representing when the Pointer was last tapped or clicked.
func (self *Pointer) PreviousTapTime() int{
    return self.Object.Get("previousTapTime").Int()
}

// SetPreviousTapTimeA A timestamp representing when the Pointer was last tapped or clicked.
func (self *Pointer) SetPreviousTapTimeA(member int) {
    self.Object.Set("previousTapTime", member)
}

// TotalTouches The total number of times this Pointer has been touched to the touchscreen.
func (self *Pointer) TotalTouches() int{
    return self.Object.Get("totalTouches").Int()
}

// SetTotalTouchesA The total number of times this Pointer has been touched to the touchscreen.
func (self *Pointer) SetTotalTouchesA(member int) {
    self.Object.Set("totalTouches", member)
}

// MsSinceLastClick The number of milliseconds since the last click or touch event.
func (self *Pointer) MsSinceLastClick() int{
    return self.Object.Get("msSinceLastClick").Int()
}

// SetMsSinceLastClickA The number of milliseconds since the last click or touch event.
func (self *Pointer) SetMsSinceLastClickA(member int) {
    self.Object.Set("msSinceLastClick", member)
}

// TargetObject The Game Object this Pointer is currently over / touching / dragging.
func (self *Pointer) TargetObject() interface{}{
    return self.Object.Get("targetObject")
}

// SetTargetObjectA The Game Object this Pointer is currently over / touching / dragging.
func (self *Pointer) SetTargetObjectA(member interface{}) {
    self.Object.Set("targetObject", member)
}

// InteractiveCandidates This array is erased and re-populated every time this Pointer is updated. It contains references to all
// of the Game Objects that were considered as being valid for processing by this Pointer, this frame. To be
// valid they must have suitable a `priorityID`, be Input enabled, visible and actually have the Pointer over
// them. You can check the contents of this array in events such as `onInputDown`, but beware it is reset
// every frame.
func (self *Pointer) InteractiveCandidates() []interface{}{
	array00 := self.Object.Get("interactiveCandidates")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetInteractiveCandidatesA This array is erased and re-populated every time this Pointer is updated. It contains references to all
// of the Game Objects that were considered as being valid for processing by this Pointer, this frame. To be
// valid they must have suitable a `priorityID`, be Input enabled, visible and actually have the Pointer over
// them. You can check the contents of this array in events such as `onInputDown`, but beware it is reset
// every frame.
func (self *Pointer) SetInteractiveCandidatesA(member []interface{}) {
    self.Object.Set("interactiveCandidates", member)
}

// Active An active pointer is one that is currently pressed down on the display. A Mouse is always active.
func (self *Pointer) Active() bool{
    return self.Object.Get("active").Bool()
}

// SetActiveA An active pointer is one that is currently pressed down on the display. A Mouse is always active.
func (self *Pointer) SetActiveA(member bool) {
    self.Object.Set("active", member)
}

// Dirty A dirty pointer needs to re-poll any interactive objects it may have been over, regardless if it has moved or not.
func (self *Pointer) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// SetDirtyA A dirty pointer needs to re-poll any interactive objects it may have been over, regardless if it has moved or not.
func (self *Pointer) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// Position A Phaser.Point object containing the current x/y values of the pointer on the display.
func (self *Pointer) Position() *Point{
    return &Point{self.Object.Get("position")}
}

// SetPositionA A Phaser.Point object containing the current x/y values of the pointer on the display.
func (self *Pointer) SetPositionA(member *Point) {
    self.Object.Set("position", member)
}

// PositionDown A Phaser.Point object containing the x/y values of the pointer when it was last in a down state on the display.
func (self *Pointer) PositionDown() *Point{
    return &Point{self.Object.Get("positionDown")}
}

// SetPositionDownA A Phaser.Point object containing the x/y values of the pointer when it was last in a down state on the display.
func (self *Pointer) SetPositionDownA(member *Point) {
    self.Object.Set("positionDown", member)
}

// PositionUp A Phaser.Point object containing the x/y values of the pointer when it was last released.
func (self *Pointer) PositionUp() *Point{
    return &Point{self.Object.Get("positionUp")}
}

// SetPositionUpA A Phaser.Point object containing the x/y values of the pointer when it was last released.
func (self *Pointer) SetPositionUpA(member *Point) {
    self.Object.Set("positionUp", member)
}

// Circle A Phaser.Circle that is centered on the x/y coordinates of this pointer, useful for hit detection.
// The Circle size is 44px (Apples recommended "finger tip" size).
func (self *Pointer) Circle() *Circle{
    return &Circle{self.Object.Get("circle")}
}

// SetCircleA A Phaser.Circle that is centered on the x/y coordinates of this pointer, useful for hit detection.
// The Circle size is 44px (Apples recommended "finger tip" size).
func (self *Pointer) SetCircleA(member *Circle) {
    self.Object.Set("circle", member)
}

// NO_BUTTON No buttons at all.
func (self *Pointer) NO_BUTTON() int{
    return self.Object.Get("NO_BUTTON").Int()
}

// SetNO_BUTTONA No buttons at all.
func (self *Pointer) SetNO_BUTTONA(member int) {
    self.Object.Set("NO_BUTTON", member)
}

// LEFT_BUTTON The Left Mouse button, or in PointerEvent devices a Touch contact or Pen contact.
func (self *Pointer) LEFT_BUTTON() int{
    return self.Object.Get("LEFT_BUTTON").Int()
}

// SetLEFT_BUTTONA The Left Mouse button, or in PointerEvent devices a Touch contact or Pen contact.
func (self *Pointer) SetLEFT_BUTTONA(member int) {
    self.Object.Set("LEFT_BUTTON", member)
}

// RIGHT_BUTTON The Right Mouse button, or in PointerEvent devices a Pen contact with a barrel button.
func (self *Pointer) RIGHT_BUTTON() int{
    return self.Object.Get("RIGHT_BUTTON").Int()
}

// SetRIGHT_BUTTONA The Right Mouse button, or in PointerEvent devices a Pen contact with a barrel button.
func (self *Pointer) SetRIGHT_BUTTONA(member int) {
    self.Object.Set("RIGHT_BUTTON", member)
}

// MIDDLE_BUTTON The Middle Mouse button.
func (self *Pointer) MIDDLE_BUTTON() int{
    return self.Object.Get("MIDDLE_BUTTON").Int()
}

// SetMIDDLE_BUTTONA The Middle Mouse button.
func (self *Pointer) SetMIDDLE_BUTTONA(member int) {
    self.Object.Set("MIDDLE_BUTTON", member)
}

// BACK_BUTTON The X1 button. This is typically the mouse Back button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) BACK_BUTTON() int{
    return self.Object.Get("BACK_BUTTON").Int()
}

// SetBACK_BUTTONA The X1 button. This is typically the mouse Back button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) SetBACK_BUTTONA(member int) {
    self.Object.Set("BACK_BUTTON", member)
}

// FORWARD_BUTTON The X2 button. This is typically the mouse Forward button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) FORWARD_BUTTON() int{
    return self.Object.Get("FORWARD_BUTTON").Int()
}

// SetFORWARD_BUTTONA The X2 button. This is typically the mouse Forward button, but is often reconfigured.
// On Linux (GTK) this is unsupported. On Windows if advanced pointer software (such as IntelliPoint) is installed this doesn't register.
func (self *Pointer) SetFORWARD_BUTTONA(member int) {
    self.Object.Set("FORWARD_BUTTON", member)
}

// ERASER_BUTTON The Eraser pen button on PointerEvent supported devices only.
func (self *Pointer) ERASER_BUTTON() int{
    return self.Object.Get("ERASER_BUTTON").Int()
}

// SetERASER_BUTTONA The Eraser pen button on PointerEvent supported devices only.
func (self *Pointer) SetERASER_BUTTONA(member int) {
    self.Object.Set("ERASER_BUTTON", member)
}

// Duration How long the Pointer has been depressed on the touchscreen or *any* of the mouse buttons have been held down.
// If not currently down it returns -1.
// If you need to test a specific mouse or pen button then access the buttons directly, i.e. `Pointer.rightButton.duration`.
func (self *Pointer) Duration() int{
    return self.Object.Get("duration").Int()
}

// SetDurationA How long the Pointer has been depressed on the touchscreen or *any* of the mouse buttons have been held down.
// If not currently down it returns -1.
// If you need to test a specific mouse or pen button then access the buttons directly, i.e. `Pointer.rightButton.duration`.
func (self *Pointer) SetDurationA(member int) {
    self.Object.Set("duration", member)
}

// WorldX Gets the X value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) WorldX() int{
    return self.Object.Get("worldX").Int()
}

// SetWorldXA Gets the X value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) SetWorldXA(member int) {
    self.Object.Set("worldX", member)
}

// WorldY Gets the Y value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) WorldY() int{
    return self.Object.Get("worldY").Int()
}

// SetWorldYA Gets the Y value of this Pointer in world coordinates based on the world camera.
func (self *Pointer) SetWorldYA(member int) {
    self.Object.Set("worldY", member)
}


// ResetButtons Resets the states of all the button booleans.
func (self *Pointer) ResetButtons() {
    self.Object.Call("resetButtons")
}

// ResetButtonsI Resets the states of all the button booleans.
func (self *Pointer) ResetButtonsI(args ...interface{}) {
    self.Object.Call("resetButtons", args)
}

// ProcessButtonsDown Called by updateButtons.
func (self *Pointer) ProcessButtonsDown(buttons int, event *MouseEvent) {
    self.Object.Call("processButtonsDown", buttons, event)
}

// ProcessButtonsDownI Called by updateButtons.
func (self *Pointer) ProcessButtonsDownI(args ...interface{}) {
    self.Object.Call("processButtonsDown", args)
}

// ProcessButtonsUp Called by updateButtons.
func (self *Pointer) ProcessButtonsUp(buttons int, event *MouseEvent) {
    self.Object.Call("processButtonsUp", buttons, event)
}

// ProcessButtonsUpI Called by updateButtons.
func (self *Pointer) ProcessButtonsUpI(args ...interface{}) {
    self.Object.Call("processButtonsUp", args)
}

// UpdateButtons Called when the event.buttons property changes from zero.
// Contains a button bitmask.
func (self *Pointer) UpdateButtons(event *MouseEvent) {
    self.Object.Call("updateButtons", event)
}

// UpdateButtonsI Called when the event.buttons property changes from zero.
// Contains a button bitmask.
func (self *Pointer) UpdateButtonsI(args ...interface{}) {
    self.Object.Call("updateButtons", args)
}

// Start Called when the Pointer is pressed onto the touchscreen.
func (self *Pointer) Start(event interface{}) {
    self.Object.Call("start", event)
}

// StartI Called when the Pointer is pressed onto the touchscreen.
func (self *Pointer) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Update Called by the Input Manager.
func (self *Pointer) Update() {
    self.Object.Call("update")
}

// UpdateI Called by the Input Manager.
func (self *Pointer) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Move Called when the Pointer is moved.
func (self *Pointer) Move(event interface{}) {
    self.Object.Call("move", event)
}

// Move1O Called when the Pointer is moved.
func (self *Pointer) Move1O(event interface{}, fromClick bool) {
    self.Object.Call("move", event, fromClick)
}

// MoveI Called when the Pointer is moved.
func (self *Pointer) MoveI(args ...interface{}) {
    self.Object.Call("move", args)
}

// ProcessInteractiveObjects Process all interactive objects to find out which ones were updated in the recent Pointer move.
func (self *Pointer) ProcessInteractiveObjects() bool{
    return self.Object.Call("processInteractiveObjects").Bool()
}

// ProcessInteractiveObjects1O Process all interactive objects to find out which ones were updated in the recent Pointer move.
func (self *Pointer) ProcessInteractiveObjects1O(fromClick bool) bool{
    return self.Object.Call("processInteractiveObjects", fromClick).Bool()
}

// ProcessInteractiveObjectsI Process all interactive objects to find out which ones were updated in the recent Pointer move.
func (self *Pointer) ProcessInteractiveObjectsI(args ...interface{}) bool{
    return self.Object.Call("processInteractiveObjects", args).Bool()
}

// SwapTarget This will change the `Pointer.targetObject` object to be the one provided.
// 
// This allows you to have fine-grained control over which object the Pointer is targeting.
// 
// Note that even if you set a new Target here, it is still able to be replaced by any other valid
// target during the next Pointer update.
func (self *Pointer) SwapTarget(newTarget *InputHandler) {
    self.Object.Call("swapTarget", newTarget)
}

// SwapTarget1O This will change the `Pointer.targetObject` object to be the one provided.
// 
// This allows you to have fine-grained control over which object the Pointer is targeting.
// 
// Note that even if you set a new Target here, it is still able to be replaced by any other valid
// target during the next Pointer update.
func (self *Pointer) SwapTarget1O(newTarget *InputHandler, silent bool) {
    self.Object.Call("swapTarget", newTarget, silent)
}

// SwapTargetI This will change the `Pointer.targetObject` object to be the one provided.
// 
// This allows you to have fine-grained control over which object the Pointer is targeting.
// 
// Note that even if you set a new Target here, it is still able to be replaced by any other valid
// target during the next Pointer update.
func (self *Pointer) SwapTargetI(args ...interface{}) {
    self.Object.Call("swapTarget", args)
}

// Leave Called when the Pointer leaves the target area.
func (self *Pointer) Leave(event interface{}) {
    self.Object.Call("leave", event)
}

// LeaveI Called when the Pointer leaves the target area.
func (self *Pointer) LeaveI(args ...interface{}) {
    self.Object.Call("leave", args)
}

// Stop Called when the Pointer leaves the touchscreen.
func (self *Pointer) Stop(event interface{}) {
    self.Object.Call("stop", event)
}

// StopI Called when the Pointer leaves the touchscreen.
func (self *Pointer) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// JustPressed The Pointer is considered justPressed if the time it was pressed onto the touchscreen or clicked is less than justPressedRate.
// Note that calling justPressed doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was pressed down just once then see the Sprite.events.onInputDown event.
func (self *Pointer) JustPressed() bool{
    return self.Object.Call("justPressed").Bool()
}

// JustPressed1O The Pointer is considered justPressed if the time it was pressed onto the touchscreen or clicked is less than justPressedRate.
// Note that calling justPressed doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was pressed down just once then see the Sprite.events.onInputDown event.
func (self *Pointer) JustPressed1O(duration int) bool{
    return self.Object.Call("justPressed", duration).Bool()
}

// JustPressedI The Pointer is considered justPressed if the time it was pressed onto the touchscreen or clicked is less than justPressedRate.
// Note that calling justPressed doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was pressed down just once then see the Sprite.events.onInputDown event.
func (self *Pointer) JustPressedI(args ...interface{}) bool{
    return self.Object.Call("justPressed", args).Bool()
}

// JustReleased The Pointer is considered justReleased if the time it left the touchscreen is less than justReleasedRate.
// Note that calling justReleased doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was released just once then see the Sprite.events.onInputUp event.
func (self *Pointer) JustReleased() bool{
    return self.Object.Call("justReleased").Bool()
}

// JustReleased1O The Pointer is considered justReleased if the time it left the touchscreen is less than justReleasedRate.
// Note that calling justReleased doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was released just once then see the Sprite.events.onInputUp event.
func (self *Pointer) JustReleased1O(duration int) bool{
    return self.Object.Call("justReleased", duration).Bool()
}

// JustReleasedI The Pointer is considered justReleased if the time it left the touchscreen is less than justReleasedRate.
// Note that calling justReleased doesn't reset the pressed status of the Pointer, it will return `true` for as long as the duration is valid.
// If you wish to check if the Pointer was released just once then see the Sprite.events.onInputUp event.
func (self *Pointer) JustReleasedI(args ...interface{}) bool{
    return self.Object.Call("justReleased", args).Bool()
}

// AddClickTrampoline Add a click trampoline to this pointer.
// 
// A click trampoline is a callback that is run on the DOM 'click' event; this is primarily
// needed with certain browsers (ie. IE11) which restrict some actions like requestFullscreen
// to the DOM 'click' event and rejects it for 'pointer*' and 'mouse*' events.
// 
// This is used internally by the ScaleManager; click trampoline usage is uncommon.
// Click trampolines can only be added to pointers that are currently down.
func (self *Pointer) AddClickTrampoline(name string, callback interface{}, callbackContext interface{}, callbackArgs interface{}) {
    self.Object.Call("addClickTrampoline", name, callback, callbackContext, callbackArgs)
}

// AddClickTrampolineI Add a click trampoline to this pointer.
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

// ProcessClickTrampolines Fire all click trampolines for which the pointers are still referring to the registered object.
func (self *Pointer) ProcessClickTrampolines() {
    self.Object.Call("processClickTrampolines")
}

// ProcessClickTrampolinesI Fire all click trampolines for which the pointers are still referring to the registered object.
func (self *Pointer) ProcessClickTrampolinesI(args ...interface{}) {
    self.Object.Call("processClickTrampolines", args)
}

// Reset Resets the Pointer properties. Called by InputManager.reset when you perform a State change.
func (self *Pointer) Reset() {
    self.Object.Call("reset")
}

// ResetI Resets the Pointer properties. Called by InputManager.reset when you perform a State change.
func (self *Pointer) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// ResetMovement Resets the movementX and movementY properties. Use in your update handler after retrieving the values.
func (self *Pointer) ResetMovement() {
    self.Object.Call("resetMovement")
}

// ResetMovementI Resets the movementX and movementY properties. Use in your update handler after retrieving the values.
func (self *Pointer) ResetMovementI(args ...interface{}) {
    self.Object.Call("resetMovement", args)
}

