// Automatic generation for Phaser.InputHandler
// generated file InputHandler.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Input Handler is bound to a specific Sprite and is responsible for managing all Input events on that Sprite.
type InputHandler struct {
    *js.Object
}


// The Sprite object to which this Input Handler belongs.
func (self *InputHandler) GetSprite() Sprite{
    return Sprite{self.Get("sprite")}
}

// The Sprite object to which this Input Handler belongs.
func (self *InputHandler) SetSprite(member Sprite) {
    self.Set("sprite", member)
}

// A reference to the currently running game.
func (self *InputHandler) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *InputHandler) SetGame(member Game) {
    self.Set("game", member)
}

// If enabled the Input Handler will process input requests and monitor pointer activity.
func (self *InputHandler) GetEnabled() bool{
    return self.Get("enabled").Bool()
}

// If enabled the Input Handler will process input requests and monitor pointer activity.
func (self *InputHandler) SetEnabled(member bool) {
    self.Set("enabled", member)
}

// A disposable flag used by the Pointer class when performing priority checks.
func (self *InputHandler) GetChecked() bool{
    return self.Get("checked").Bool()
}

// A disposable flag used by the Pointer class when performing priority checks.
func (self *InputHandler) SetChecked(member bool) {
    self.Set("checked", member)
}

// The priorityID is used to determine which game objects should get priority when input events occur. For example if you have
// several Sprites that overlap, by default the one at the top of the display list is given priority for input events. You can
// stop this from happening by controlling the priorityID value. The higher the value, the more important they are considered to the Input events.
func (self *InputHandler) GetPriorityID() float64{
    return self.Get("priorityID").Float()
}

// The priorityID is used to determine which game objects should get priority when input events occur. For example if you have
// several Sprites that overlap, by default the one at the top of the display list is given priority for input events. You can
// stop this from happening by controlling the priorityID value. The higher the value, the more important they are considered to the Input events.
func (self *InputHandler) SetPriorityID(member float64) {
    self.Set("priorityID", member)
}

// On a desktop browser you can set the 'hand' cursor to appear when moving over the Sprite.
func (self *InputHandler) GetUseHandCursor() bool{
    return self.Get("useHandCursor").Bool()
}

// On a desktop browser you can set the 'hand' cursor to appear when moving over the Sprite.
func (self *InputHandler) SetUseHandCursor(member bool) {
    self.Set("useHandCursor", member)
}

// true if the Sprite is being currently dragged.
func (self *InputHandler) GetIsDragged() bool{
    return self.Get("isDragged").Bool()
}

// true if the Sprite is being currently dragged.
func (self *InputHandler) SetIsDragged(member bool) {
    self.Set("isDragged", member)
}

// Controls if the Sprite is allowed to be dragged horizontally.
func (self *InputHandler) GetAllowHorizontalDrag() bool{
    return self.Get("allowHorizontalDrag").Bool()
}

// Controls if the Sprite is allowed to be dragged horizontally.
func (self *InputHandler) SetAllowHorizontalDrag(member bool) {
    self.Set("allowHorizontalDrag", member)
}

// Controls if the Sprite is allowed to be dragged vertically.
func (self *InputHandler) GetAllowVerticalDrag() bool{
    return self.Get("allowVerticalDrag").Bool()
}

// Controls if the Sprite is allowed to be dragged vertically.
func (self *InputHandler) SetAllowVerticalDrag(member bool) {
    self.Set("allowVerticalDrag", member)
}

// If true when this Sprite is clicked or dragged it will automatically be bought to the top of the Group it is within.
func (self *InputHandler) GetBringToTop() bool{
    return self.Get("bringToTop").Bool()
}

// If true when this Sprite is clicked or dragged it will automatically be bought to the top of the Group it is within.
func (self *InputHandler) SetBringToTop(member bool) {
    self.Set("bringToTop", member)
}

// A Point object that contains by how far the Sprite snap is offset.
func (self *InputHandler) GetSnapOffset() Point{
    return Point{self.Get("snapOffset")}
}

// A Point object that contains by how far the Sprite snap is offset.
func (self *InputHandler) SetSnapOffset(member Point) {
    self.Set("snapOffset", member)
}

// When the Sprite is dragged this controls if the center of the Sprite will snap to the pointer on drag or not.
func (self *InputHandler) GetSnapOnDrag() bool{
    return self.Get("snapOnDrag").Bool()
}

// When the Sprite is dragged this controls if the center of the Sprite will snap to the pointer on drag or not.
func (self *InputHandler) SetSnapOnDrag(member bool) {
    self.Set("snapOnDrag", member)
}

// When the Sprite is dragged this controls if the Sprite will be snapped on release.
func (self *InputHandler) GetSnapOnRelease() bool{
    return self.Get("snapOnRelease").Bool()
}

// When the Sprite is dragged this controls if the Sprite will be snapped on release.
func (self *InputHandler) SetSnapOnRelease(member bool) {
    self.Set("snapOnRelease", member)
}

// When a Sprite has snapping enabled this holds the width of the snap grid.
func (self *InputHandler) GetSnapX() float64{
    return self.Get("snapX").Float()
}

// When a Sprite has snapping enabled this holds the width of the snap grid.
func (self *InputHandler) SetSnapX(member float64) {
    self.Set("snapX", member)
}

// When a Sprite has snapping enabled this holds the height of the snap grid.
func (self *InputHandler) GetSnapY() float64{
    return self.Get("snapY").Float()
}

// When a Sprite has snapping enabled this holds the height of the snap grid.
func (self *InputHandler) SetSnapY(member float64) {
    self.Set("snapY", member)
}

// This defines the top-left X coordinate of the snap grid.
func (self *InputHandler) GetSnapOffsetX() float64{
    return self.Get("snapOffsetX").Float()
}

// This defines the top-left X coordinate of the snap grid.
func (self *InputHandler) SetSnapOffsetX(member float64) {
    self.Set("snapOffsetX", member)
}

// This defines the top-left Y coordinate of the snap grid..
func (self *InputHandler) GetSnapOffsetY() float64{
    return self.Get("snapOffsetY").Float()
}

// This defines the top-left Y coordinate of the snap grid..
func (self *InputHandler) SetSnapOffsetY(member float64) {
    self.Set("snapOffsetY", member)
}

// Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive, especially on mobile (where it's not even needed!) so only enable if required. Also see the less-expensive InputHandler.pixelPerfectClick. Use a pixel perfect check when testing for pointer over.
func (self *InputHandler) GetPixelPerfectOver() bool{
    return self.Get("pixelPerfectOver").Bool()
}

// Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive, especially on mobile (where it's not even needed!) so only enable if required. Also see the less-expensive InputHandler.pixelPerfectClick. Use a pixel perfect check when testing for pointer over.
func (self *InputHandler) SetPixelPerfectOver(member bool) {
    self.Set("pixelPerfectOver", member)
}

// Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite when it's clicked or touched.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive so only enable if you really need it. Use a pixel perfect check when testing for clicks or touches on the Sprite.
func (self *InputHandler) GetPixelPerfectClick() bool{
    return self.Get("pixelPerfectClick").Bool()
}

// Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite when it's clicked or touched.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive so only enable if you really need it. Use a pixel perfect check when testing for clicks or touches on the Sprite.
func (self *InputHandler) SetPixelPerfectClick(member bool) {
    self.Set("pixelPerfectClick", member)
}

// The alpha tolerance threshold. If the alpha value of the pixel matches or is above this value, it's considered a hit.
func (self *InputHandler) GetPixelPerfectAlpha() float64{
    return self.Get("pixelPerfectAlpha").Float()
}

// The alpha tolerance threshold. If the alpha value of the pixel matches or is above this value, it's considered a hit.
func (self *InputHandler) SetPixelPerfectAlpha(member float64) {
    self.Set("pixelPerfectAlpha", member)
}

// Is this sprite allowed to be dragged by the mouse? true = yes, false = no
func (self *InputHandler) GetDraggable() bool{
    return self.Get("draggable").Bool()
}

// Is this sprite allowed to be dragged by the mouse? true = yes, false = no
func (self *InputHandler) SetDraggable(member bool) {
    self.Set("draggable", member)
}

// A region of the game world within which the sprite is restricted during drag.
func (self *InputHandler) GetBoundsRect() Rectangle{
    return Rectangle{self.Get("boundsRect")}
}

// A region of the game world within which the sprite is restricted during drag.
func (self *InputHandler) SetBoundsRect(member Rectangle) {
    self.Set("boundsRect", member)
}

// A Sprite the bounds of which this sprite is restricted during drag.
func (self *InputHandler) GetBoundsSprite() Sprite{
    return Sprite{self.Get("boundsSprite")}
}

// A Sprite the bounds of which this sprite is restricted during drag.
func (self *InputHandler) SetBoundsSprite(member Sprite) {
    self.Set("boundsSprite", member)
}

// EXPERIMENTAL: Please do not use this property unless you know what it does. Likely to change in the future.
func (self *InputHandler) GetScaleLayer() bool{
    return self.Get("scaleLayer").Bool()
}

// EXPERIMENTAL: Please do not use this property unless you know what it does. Likely to change in the future.
func (self *InputHandler) SetScaleLayer(member bool) {
    self.Set("scaleLayer", member)
}

// The offset from the Sprites position that dragging takes place from.
func (self *InputHandler) GetDragOffset() Point{
    return Point{self.Get("dragOffset")}
}

// The offset from the Sprites position that dragging takes place from.
func (self *InputHandler) SetDragOffset(member Point) {
    self.Set("dragOffset", member)
}

// Is the Sprite dragged from its center, or the point at which the Pointer was pressed down upon it?
func (self *InputHandler) GetDragFromCenter() bool{
    return self.Get("dragFromCenter").Bool()
}

// Is the Sprite dragged from its center, or the point at which the Pointer was pressed down upon it?
func (self *InputHandler) SetDragFromCenter(member bool) {
    self.Set("dragFromCenter", member)
}

// If enabled, when the Sprite stops being dragged, it will only dispatch the `onDragStop` event, and not the `onInputUp` event. If set to `false` it will dispatch both events.
func (self *InputHandler) GetDragStopBlocksInputUp() bool{
    return self.Get("dragStopBlocksInputUp").Bool()
}

// If enabled, when the Sprite stops being dragged, it will only dispatch the `onDragStop` event, and not the `onInputUp` event. If set to `false` it will dispatch both events.
func (self *InputHandler) SetDragStopBlocksInputUp(member bool) {
    self.Set("dragStopBlocksInputUp", member)
}

// The Point from which the most recent drag started from. Useful if you need to return an object to its starting position.
func (self *InputHandler) GetDragStartPoint() Point{
    return Point{self.Get("dragStartPoint")}
}

// The Point from which the most recent drag started from. Useful if you need to return an object to its starting position.
func (self *InputHandler) SetDragStartPoint(member Point) {
    self.Set("dragStartPoint", member)
}

// The distance, in pixels, the pointer has to move while being held down, before the Sprite thinks it is being dragged.
func (self *InputHandler) GetDragDistanceThreshold() int{
    return self.Get("dragDistanceThreshold").Int()
}

// The distance, in pixels, the pointer has to move while being held down, before the Sprite thinks it is being dragged.
func (self *InputHandler) SetDragDistanceThreshold(member int) {
    self.Set("dragDistanceThreshold", member)
}

// The amount of time, in ms, the pointer has to be held down over the Sprite before it thinks it is being dragged.
func (self *InputHandler) GetDragTimeThreshold() int{
    return self.Get("dragTimeThreshold").Int()
}

// The amount of time, in ms, the pointer has to be held down over the Sprite before it thinks it is being dragged.
func (self *InputHandler) SetDragTimeThreshold(member int) {
    self.Set("dragTimeThreshold", member)
}

// A Point object containing the coordinates of the Pointer when it was first pressed down onto this Sprite.
func (self *InputHandler) GetDownPoint() Point{
    return Point{self.Get("downPoint")}
}

// A Point object containing the coordinates of the Pointer when it was first pressed down onto this Sprite.
func (self *InputHandler) SetDownPoint(member Point) {
    self.Set("downPoint", member)
}

// If the sprite is set to snap while dragging this holds the point of the most recent 'snap' event.
func (self *InputHandler) GetSnapPoint() Point{
    return Point{self.Get("snapPoint")}
}

// If the sprite is set to snap while dragging this holds the point of the most recent 'snap' event.
func (self *InputHandler) SetSnapPoint(member Point) {
    self.Set("snapPoint", member)
}



// Starts the Input Handler running. This is called automatically when you enable input on a Sprite, or can be called directly if you need to set a specific priority.
func (self *InputHandler) StartI(args ...interface{}) Sprite{
    return Sprite{self.Call("start", args)}
}

// Handles when the parent Sprite is added to a new Group.
func (self *InputHandler) AddedToGroupI(args ...interface{}) {
    self.Call("addedToGroup", args)
}

// Handles when the parent Sprite is removed from a Group.
func (self *InputHandler) RemovedFromGroupI(args ...interface{}) {
    self.Call("removedFromGroup", args)
}

// Resets the Input Handler and disables it.
func (self *InputHandler) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Stops the Input Handler from running.
func (self *InputHandler) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// Clean up memory.
func (self *InputHandler) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Checks if the object this InputHandler is bound to is valid for consideration in the Pointer move event.
// This is called by Phaser.Pointer and shouldn't typically be called directly.
func (self *InputHandler) ValidForInputI(args ...interface{}) bool{
    return self.Call("validForInput", args).Bool()
}

// Is this object using pixel perfect checking?
func (self *InputHandler) IsPixelPerfectI(args ...interface{}) bool{
    return self.Call("isPixelPerfect", args).Bool()
}

// The x coordinate of the Input pointer, relative to the top-left of the parent Sprite.
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerXI(args ...interface{}) float64{
    return self.Call("pointerX", args).Float()
}

// The y coordinate of the Input pointer, relative to the top-left of the parent Sprite
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerYI(args ...interface{}) float64{
    return self.Call("pointerY", args).Float()
}

// If the Pointer is down this returns true.
// This *only* checks if the Pointer is down, not if it's down over any specific Sprite.
func (self *InputHandler) PointerDownI(args ...interface{}) bool{
    return self.Call("pointerDown", args).Bool()
}

// If the Pointer is up this returns true.
// This *only* checks if the Pointer is up, not if it's up over any specific Sprite.
func (self *InputHandler) PointerUpI(args ...interface{}) bool{
    return self.Call("pointerUp", args).Bool()
}

// A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeDownI(args ...interface{}) float64{
    return self.Call("pointerTimeDown", args).Float()
}

// A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeUpI(args ...interface{}) float64{
    return self.Call("pointerTimeUp", args).Float()
}

// Is the Pointer over this Sprite?
func (self *InputHandler) PointerOverI(args ...interface{}) bool{
    return self.Call("pointerOver", args).Bool()
}

// Is the Pointer outside of this Sprite?
func (self *InputHandler) PointerOutI(args ...interface{}) bool{
    return self.Call("pointerOut", args).Bool()
}

// A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeOverI(args ...interface{}) float64{
    return self.Call("pointerTimeOver", args).Float()
}

// A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeOutI(args ...interface{}) float64{
    return self.Call("pointerTimeOut", args).Float()
}

// Is this sprite being dragged by the mouse or not?
func (self *InputHandler) PointerDraggedI(args ...interface{}) bool{
    return self.Call("pointerDragged", args).Bool()
}

// Checks if the given pointer is both down and over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerDownI(args ...interface{}) bool{
    return self.Call("checkPointerDown", args).Bool()
}

// Checks if the given pointer is over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerOverI(args ...interface{}) bool{
    return self.Call("checkPointerOver", args).Bool()
}

// Runs a pixel perfect check against the given x/y coordinates of the Sprite this InputHandler is bound to.
// It compares the alpha value of the pixel and if >= InputHandler.pixelPerfectAlpha it returns true.
func (self *InputHandler) CheckPixelI(args ...interface{}) bool{
    return self.Call("checkPixel", args).Bool()
}

// Internal Update method. This is called automatically and handles the Pointer
// and drag update loops.
func (self *InputHandler) UpdateI(args ...interface{}) bool{
    return self.Call("update", args).Bool()
}

// Internal method handling the pointer over event.
func (self *InputHandler) _pointerOverHandlerI(args ...interface{}) {
    self.Call("_pointerOverHandler", args)
}

// Internal method handling the pointer out event.
func (self *InputHandler) _pointerOutHandlerI(args ...interface{}) {
    self.Call("_pointerOutHandler", args)
}

// Internal method handling the touched / clicked event.
func (self *InputHandler) _touchedHandlerI(args ...interface{}) {
    self.Call("_touchedHandler", args)
}

// Internal method handling the drag threshold timer.
func (self *InputHandler) DragTimeElapsedI(args ...interface{}) {
    self.Call("dragTimeElapsed", args)
}

// Internal method handling the pointer released event.
func (self *InputHandler) _releasedHandlerI(args ...interface{}) {
    self.Call("_releasedHandler", args)
}

// Called as a Pointer actively drags this Game Object.
func (self *InputHandler) UpdateDragI(args ...interface{}) bool{
    return self.Call("updateDrag", args).Bool()
}

// Returns true if the pointer has entered the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustOverI(args ...interface{}) bool{
    return self.Call("justOver", args).Bool()
}

// Returns true if the pointer has left the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustOutI(args ...interface{}) bool{
    return self.Call("justOut", args).Bool()
}

// Returns true if the pointer has touched or clicked on the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustPressedI(args ...interface{}) bool{
    return self.Call("justPressed", args).Bool()
}

// Returns true if the pointer was touching this Sprite, but has been released within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustReleasedI(args ...interface{}) bool{
    return self.Call("justReleased", args).Bool()
}

// If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) OverDurationI(args ...interface{}) float64{
    return self.Call("overDuration", args).Float()
}

// If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) DownDurationI(args ...interface{}) float64{
    return self.Call("downDuration", args).Float()
}

// Allow this Sprite to be dragged by any valid pointer.
// 
// When the drag begins the Sprite.events.onDragStart event will be dispatched.
// 
// When the drag completes by way of the user letting go of the pointer that was dragging the sprite, the Sprite.events.onDragStop event is dispatched.
// 
// You can control the thresholds over when a drag starts via the properties:
// 
// `Pointer.dragDistanceThreshold` the distance, in pixels, that the pointer has to move
// before the drag will start.
// 
// `Pointer.dragTimeThreshold` the time, in ms, that the pointer must be held down on
// the Sprite before the drag will start.
// 
// You can set either (or both) of these properties after enabling a Sprite for drag.
// 
// For the duration of the drag the Sprite.events.onDragUpdate event is dispatched. This event is only dispatched when the pointer actually
// changes position and moves. The event sends 5 parameters: `sprite`, `pointer`, `dragX`, `dragY` and `snapPoint`.
func (self *InputHandler) EnableDragI(args ...interface{}) {
    self.Call("enableDrag", args)
}

// Stops this sprite from being able to be dragged.
// If it is currently the target of an active drag it will be stopped immediately; also disables any set callbacks.
func (self *InputHandler) DisableDragI(args ...interface{}) {
    self.Call("disableDrag", args)
}

// Called by Pointer when drag starts on this Sprite. Should not usually be called directly.
func (self *InputHandler) StartDragI(args ...interface{}) {
    self.Call("startDrag", args)
}

// Warning: EXPERIMENTAL
func (self *InputHandler) GlobalToLocalXI(args ...interface{}) {
    self.Call("globalToLocalX", args)
}

// Warning: EXPERIMENTAL
func (self *InputHandler) GlobalToLocalYI(args ...interface{}) {
    self.Call("globalToLocalY", args)
}

// Called by Pointer when drag is stopped on this Sprite. Should not usually be called directly.
func (self *InputHandler) StopDragI(args ...interface{}) {
    self.Call("stopDrag", args)
}

// Restricts this sprite to drag movement only on the given axis. Note: If both are set to false the sprite will never move!
func (self *InputHandler) SetDragLockI(args ...interface{}) {
    self.Call("setDragLock", args)
}

// Make this Sprite snap to the given grid either during drag or when it's released.
// For example 16x16 as the snapX and snapY would make the sprite snap to every 16 pixels.
func (self *InputHandler) EnableSnapI(args ...interface{}) {
    self.Call("enableSnap", args)
}

// Stops the sprite from snapping to a grid during drag or release.
func (self *InputHandler) DisableSnapI(args ...interface{}) {
    self.Call("disableSnap", args)
}

// Bounds Rect check for the sprite drag
func (self *InputHandler) CheckBoundsRectI(args ...interface{}) {
    self.Call("checkBoundsRect", args)
}

// Parent Sprite Bounds check for the sprite drag.
func (self *InputHandler) CheckBoundsSpriteI(args ...interface{}) {
    self.Call("checkBoundsSprite", args)
}
