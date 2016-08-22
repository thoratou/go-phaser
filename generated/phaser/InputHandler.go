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
func (self *InputHandler) GetSpriteA() *Sprite{
    return &Sprite{self.Object.Get("sprite")}
}

// The Sprite object to which this Input Handler belongs.
func (self *InputHandler) SetSpriteA(member *Sprite) {
    self.Object.Set("sprite", member)
}

// A reference to the currently running game.
func (self *InputHandler) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *InputHandler) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// If enabled the Input Handler will process input requests and monitor pointer activity.
func (self *InputHandler) GetEnabledA() bool{
    return self.Object.Get("enabled").Bool()
}

// If enabled the Input Handler will process input requests and monitor pointer activity.
func (self *InputHandler) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}

// A disposable flag used by the Pointer class when performing priority checks.
func (self *InputHandler) GetCheckedA() bool{
    return self.Object.Get("checked").Bool()
}

// A disposable flag used by the Pointer class when performing priority checks.
func (self *InputHandler) SetCheckedA(member bool) {
    self.Object.Set("checked", member)
}

// The priorityID is used to determine which game objects should get priority when input events occur. For example if you have
// several Sprites that overlap, by default the one at the top of the display list is given priority for input events. You can
// stop this from happening by controlling the priorityID value. The higher the value, the more important they are considered to the Input events.
func (self *InputHandler) GetPriorityIDA() int{
    return self.Object.Get("priorityID").Int()
}

// The priorityID is used to determine which game objects should get priority when input events occur. For example if you have
// several Sprites that overlap, by default the one at the top of the display list is given priority for input events. You can
// stop this from happening by controlling the priorityID value. The higher the value, the more important they are considered to the Input events.
func (self *InputHandler) SetPriorityIDA(member int) {
    self.Object.Set("priorityID", member)
}

// On a desktop browser you can set the 'hand' cursor to appear when moving over the Sprite.
func (self *InputHandler) GetUseHandCursorA() bool{
    return self.Object.Get("useHandCursor").Bool()
}

// On a desktop browser you can set the 'hand' cursor to appear when moving over the Sprite.
func (self *InputHandler) SetUseHandCursorA(member bool) {
    self.Object.Set("useHandCursor", member)
}

// true if the Sprite is being currently dragged.
func (self *InputHandler) GetIsDraggedA() bool{
    return self.Object.Get("isDragged").Bool()
}

// true if the Sprite is being currently dragged.
func (self *InputHandler) SetIsDraggedA(member bool) {
    self.Object.Set("isDragged", member)
}

// Controls if the Sprite is allowed to be dragged horizontally.
func (self *InputHandler) GetAllowHorizontalDragA() bool{
    return self.Object.Get("allowHorizontalDrag").Bool()
}

// Controls if the Sprite is allowed to be dragged horizontally.
func (self *InputHandler) SetAllowHorizontalDragA(member bool) {
    self.Object.Set("allowHorizontalDrag", member)
}

// Controls if the Sprite is allowed to be dragged vertically.
func (self *InputHandler) GetAllowVerticalDragA() bool{
    return self.Object.Get("allowVerticalDrag").Bool()
}

// Controls if the Sprite is allowed to be dragged vertically.
func (self *InputHandler) SetAllowVerticalDragA(member bool) {
    self.Object.Set("allowVerticalDrag", member)
}

// If true when this Sprite is clicked or dragged it will automatically be bought to the top of the Group it is within.
func (self *InputHandler) GetBringToTopA() bool{
    return self.Object.Get("bringToTop").Bool()
}

// If true when this Sprite is clicked or dragged it will automatically be bought to the top of the Group it is within.
func (self *InputHandler) SetBringToTopA(member bool) {
    self.Object.Set("bringToTop", member)
}

// A Point object that contains by how far the Sprite snap is offset.
func (self *InputHandler) GetSnapOffsetA() *Point{
    return &Point{self.Object.Get("snapOffset")}
}

// A Point object that contains by how far the Sprite snap is offset.
func (self *InputHandler) SetSnapOffsetA(member *Point) {
    self.Object.Set("snapOffset", member)
}

// When the Sprite is dragged this controls if the center of the Sprite will snap to the pointer on drag or not.
func (self *InputHandler) GetSnapOnDragA() bool{
    return self.Object.Get("snapOnDrag").Bool()
}

// When the Sprite is dragged this controls if the center of the Sprite will snap to the pointer on drag or not.
func (self *InputHandler) SetSnapOnDragA(member bool) {
    self.Object.Set("snapOnDrag", member)
}

// When the Sprite is dragged this controls if the Sprite will be snapped on release.
func (self *InputHandler) GetSnapOnReleaseA() bool{
    return self.Object.Get("snapOnRelease").Bool()
}

// When the Sprite is dragged this controls if the Sprite will be snapped on release.
func (self *InputHandler) SetSnapOnReleaseA(member bool) {
    self.Object.Set("snapOnRelease", member)
}

// When a Sprite has snapping enabled this holds the width of the snap grid.
func (self *InputHandler) GetSnapXA() int{
    return self.Object.Get("snapX").Int()
}

// When a Sprite has snapping enabled this holds the width of the snap grid.
func (self *InputHandler) SetSnapXA(member int) {
    self.Object.Set("snapX", member)
}

// When a Sprite has snapping enabled this holds the height of the snap grid.
func (self *InputHandler) GetSnapYA() int{
    return self.Object.Get("snapY").Int()
}

// When a Sprite has snapping enabled this holds the height of the snap grid.
func (self *InputHandler) SetSnapYA(member int) {
    self.Object.Set("snapY", member)
}

// This defines the top-left X coordinate of the snap grid.
func (self *InputHandler) GetSnapOffsetXA() int{
    return self.Object.Get("snapOffsetX").Int()
}

// This defines the top-left X coordinate of the snap grid.
func (self *InputHandler) SetSnapOffsetXA(member int) {
    self.Object.Set("snapOffsetX", member)
}

// This defines the top-left Y coordinate of the snap grid..
func (self *InputHandler) GetSnapOffsetYA() int{
    return self.Object.Get("snapOffsetY").Int()
}

// This defines the top-left Y coordinate of the snap grid..
func (self *InputHandler) SetSnapOffsetYA(member int) {
    self.Object.Set("snapOffsetY", member)
}

// Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive, especially on mobile (where it's not even needed!) so only enable if required. Also see the less-expensive InputHandler.pixelPerfectClick. Use a pixel perfect check when testing for pointer over.
func (self *InputHandler) GetPixelPerfectOverA() bool{
    return self.Object.Get("pixelPerfectOver").Bool()
}

// Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive, especially on mobile (where it's not even needed!) so only enable if required. Also see the less-expensive InputHandler.pixelPerfectClick. Use a pixel perfect check when testing for pointer over.
func (self *InputHandler) SetPixelPerfectOverA(member bool) {
    self.Object.Set("pixelPerfectOver", member)
}

// Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite when it's clicked or touched.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive so only enable if you really need it. Use a pixel perfect check when testing for clicks or touches on the Sprite.
func (self *InputHandler) GetPixelPerfectClickA() bool{
    return self.Object.Get("pixelPerfectClick").Bool()
}

// Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite when it's clicked or touched.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive so only enable if you really need it. Use a pixel perfect check when testing for clicks or touches on the Sprite.
func (self *InputHandler) SetPixelPerfectClickA(member bool) {
    self.Object.Set("pixelPerfectClick", member)
}

// The alpha tolerance threshold. If the alpha value of the pixel matches or is above this value, it's considered a hit.
func (self *InputHandler) GetPixelPerfectAlphaA() int{
    return self.Object.Get("pixelPerfectAlpha").Int()
}

// The alpha tolerance threshold. If the alpha value of the pixel matches or is above this value, it's considered a hit.
func (self *InputHandler) SetPixelPerfectAlphaA(member int) {
    self.Object.Set("pixelPerfectAlpha", member)
}

// Is this sprite allowed to be dragged by the mouse? true = yes, false = no
func (self *InputHandler) GetDraggableA() bool{
    return self.Object.Get("draggable").Bool()
}

// Is this sprite allowed to be dragged by the mouse? true = yes, false = no
func (self *InputHandler) SetDraggableA(member bool) {
    self.Object.Set("draggable", member)
}

// A region of the game world within which the sprite is restricted during drag.
func (self *InputHandler) GetBoundsRectA() *Rectangle{
    return &Rectangle{self.Object.Get("boundsRect")}
}

// A region of the game world within which the sprite is restricted during drag.
func (self *InputHandler) SetBoundsRectA(member *Rectangle) {
    self.Object.Set("boundsRect", member)
}

// A Sprite the bounds of which this sprite is restricted during drag.
func (self *InputHandler) GetBoundsSpriteA() *Sprite{
    return &Sprite{self.Object.Get("boundsSprite")}
}

// A Sprite the bounds of which this sprite is restricted during drag.
func (self *InputHandler) SetBoundsSpriteA(member *Sprite) {
    self.Object.Set("boundsSprite", member)
}

// EXPERIMENTAL: Please do not use this property unless you know what it does. Likely to change in the future.
func (self *InputHandler) GetScaleLayerA() bool{
    return self.Object.Get("scaleLayer").Bool()
}

// EXPERIMENTAL: Please do not use this property unless you know what it does. Likely to change in the future.
func (self *InputHandler) SetScaleLayerA(member bool) {
    self.Object.Set("scaleLayer", member)
}

// The offset from the Sprites position that dragging takes place from.
func (self *InputHandler) GetDragOffsetA() *Point{
    return &Point{self.Object.Get("dragOffset")}
}

// The offset from the Sprites position that dragging takes place from.
func (self *InputHandler) SetDragOffsetA(member *Point) {
    self.Object.Set("dragOffset", member)
}

// Is the Sprite dragged from its center, or the point at which the Pointer was pressed down upon it?
func (self *InputHandler) GetDragFromCenterA() bool{
    return self.Object.Get("dragFromCenter").Bool()
}

// Is the Sprite dragged from its center, or the point at which the Pointer was pressed down upon it?
func (self *InputHandler) SetDragFromCenterA(member bool) {
    self.Object.Set("dragFromCenter", member)
}

// If enabled, when the Sprite stops being dragged, it will only dispatch the `onDragStop` event, and not the `onInputUp` event. If set to `false` it will dispatch both events.
func (self *InputHandler) GetDragStopBlocksInputUpA() bool{
    return self.Object.Get("dragStopBlocksInputUp").Bool()
}

// If enabled, when the Sprite stops being dragged, it will only dispatch the `onDragStop` event, and not the `onInputUp` event. If set to `false` it will dispatch both events.
func (self *InputHandler) SetDragStopBlocksInputUpA(member bool) {
    self.Object.Set("dragStopBlocksInputUp", member)
}

// The Point from which the most recent drag started from. Useful if you need to return an object to its starting position.
func (self *InputHandler) GetDragStartPointA() *Point{
    return &Point{self.Object.Get("dragStartPoint")}
}

// The Point from which the most recent drag started from. Useful if you need to return an object to its starting position.
func (self *InputHandler) SetDragStartPointA(member *Point) {
    self.Object.Set("dragStartPoint", member)
}

// The distance, in pixels, the pointer has to move while being held down, before the Sprite thinks it is being dragged.
func (self *InputHandler) GetDragDistanceThresholdA() int{
    return self.Object.Get("dragDistanceThreshold").Int()
}

// The distance, in pixels, the pointer has to move while being held down, before the Sprite thinks it is being dragged.
func (self *InputHandler) SetDragDistanceThresholdA(member int) {
    self.Object.Set("dragDistanceThreshold", member)
}

// The amount of time, in ms, the pointer has to be held down over the Sprite before it thinks it is being dragged.
func (self *InputHandler) GetDragTimeThresholdA() int{
    return self.Object.Get("dragTimeThreshold").Int()
}

// The amount of time, in ms, the pointer has to be held down over the Sprite before it thinks it is being dragged.
func (self *InputHandler) SetDragTimeThresholdA(member int) {
    self.Object.Set("dragTimeThreshold", member)
}

// A Point object containing the coordinates of the Pointer when it was first pressed down onto this Sprite.
func (self *InputHandler) GetDownPointA() *Point{
    return &Point{self.Object.Get("downPoint")}
}

// A Point object containing the coordinates of the Pointer when it was first pressed down onto this Sprite.
func (self *InputHandler) SetDownPointA(member *Point) {
    self.Object.Set("downPoint", member)
}

// If the sprite is set to snap while dragging this holds the point of the most recent 'snap' event.
func (self *InputHandler) GetSnapPointA() *Point{
    return &Point{self.Object.Get("snapPoint")}
}

// If the sprite is set to snap while dragging this holds the point of the most recent 'snap' event.
func (self *InputHandler) SetSnapPointA(member *Point) {
    self.Object.Set("snapPoint", member)
}



// Starts the Input Handler running. This is called automatically when you enable input on a Sprite, or can be called directly if you need to set a specific priority.
func (self *InputHandler) Start(priority int, useHandCursor bool) *Sprite{
    return &Sprite{self.Object.Call("start", priority, useHandCursor)}
}

// Starts the Input Handler running. This is called automatically when you enable input on a Sprite, or can be called directly if you need to set a specific priority.
func (self *InputHandler) StartI(args ...interface{}) *Sprite{
    return &Sprite{self.Object.Call("start", args)}
}

// Handles when the parent Sprite is added to a new Group.
func (self *InputHandler) AddedToGroup() {
    self.Object.Call("addedToGroup")
}

// Handles when the parent Sprite is added to a new Group.
func (self *InputHandler) AddedToGroupI(args ...interface{}) {
    self.Object.Call("addedToGroup", args)
}

// Handles when the parent Sprite is removed from a Group.
func (self *InputHandler) RemovedFromGroup() {
    self.Object.Call("removedFromGroup")
}

// Handles when the parent Sprite is removed from a Group.
func (self *InputHandler) RemovedFromGroupI(args ...interface{}) {
    self.Object.Call("removedFromGroup", args)
}

// Resets the Input Handler and disables it.
func (self *InputHandler) Reset() {
    self.Object.Call("reset")
}

// Resets the Input Handler and disables it.
func (self *InputHandler) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Stops the Input Handler from running.
func (self *InputHandler) Stop() {
    self.Object.Call("stop")
}

// Stops the Input Handler from running.
func (self *InputHandler) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Clean up memory.
func (self *InputHandler) Destroy() {
    self.Object.Call("destroy")
}

// Clean up memory.
func (self *InputHandler) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Checks if the object this InputHandler is bound to is valid for consideration in the Pointer move event.
// This is called by Phaser.Pointer and shouldn't typically be called directly.
func (self *InputHandler) ValidForInput(highestID int, highestRenderID int, includePixelPerfect bool) bool{
    return self.Object.Call("validForInput", highestID, highestRenderID, includePixelPerfect).Bool()
}

// Checks if the object this InputHandler is bound to is valid for consideration in the Pointer move event.
// This is called by Phaser.Pointer and shouldn't typically be called directly.
func (self *InputHandler) ValidForInputI(args ...interface{}) bool{
    return self.Object.Call("validForInput", args).Bool()
}

// Is this object using pixel perfect checking?
func (self *InputHandler) IsPixelPerfect() bool{
    return self.Object.Call("isPixelPerfect").Bool()
}

// Is this object using pixel perfect checking?
func (self *InputHandler) IsPixelPerfectI(args ...interface{}) bool{
    return self.Object.Call("isPixelPerfect", args).Bool()
}

// The x coordinate of the Input pointer, relative to the top-left of the parent Sprite.
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerX(pointerId int) int{
    return self.Object.Call("pointerX", pointerId).Int()
}

// The x coordinate of the Input pointer, relative to the top-left of the parent Sprite.
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerXI(args ...interface{}) int{
    return self.Object.Call("pointerX", args).Int()
}

// The y coordinate of the Input pointer, relative to the top-left of the parent Sprite
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerY(pointerId int) int{
    return self.Object.Call("pointerY", pointerId).Int()
}

// The y coordinate of the Input pointer, relative to the top-left of the parent Sprite
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerYI(args ...interface{}) int{
    return self.Object.Call("pointerY", args).Int()
}

// If the Pointer is down this returns true.
// This *only* checks if the Pointer is down, not if it's down over any specific Sprite.
func (self *InputHandler) PointerDown(pointerId int) bool{
    return self.Object.Call("pointerDown", pointerId).Bool()
}

// If the Pointer is down this returns true.
// This *only* checks if the Pointer is down, not if it's down over any specific Sprite.
func (self *InputHandler) PointerDownI(args ...interface{}) bool{
    return self.Object.Call("pointerDown", args).Bool()
}

// If the Pointer is up this returns true.
// This *only* checks if the Pointer is up, not if it's up over any specific Sprite.
func (self *InputHandler) PointerUp(pointerId int) bool{
    return self.Object.Call("pointerUp", pointerId).Bool()
}

// If the Pointer is up this returns true.
// This *only* checks if the Pointer is up, not if it's up over any specific Sprite.
func (self *InputHandler) PointerUpI(args ...interface{}) bool{
    return self.Object.Call("pointerUp", args).Bool()
}

// A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeDown(pointerId int) int{
    return self.Object.Call("pointerTimeDown", pointerId).Int()
}

// A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeDownI(args ...interface{}) int{
    return self.Object.Call("pointerTimeDown", args).Int()
}

// A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeUp(pointerId int) int{
    return self.Object.Call("pointerTimeUp", pointerId).Int()
}

// A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeUpI(args ...interface{}) int{
    return self.Object.Call("pointerTimeUp", args).Int()
}

// Is the Pointer over this Sprite?
func (self *InputHandler) PointerOver(pointerId int) bool{
    return self.Object.Call("pointerOver", pointerId).Bool()
}

// Is the Pointer over this Sprite?
func (self *InputHandler) PointerOverI(args ...interface{}) bool{
    return self.Object.Call("pointerOver", args).Bool()
}

// Is the Pointer outside of this Sprite?
func (self *InputHandler) PointerOut(pointerId int) bool{
    return self.Object.Call("pointerOut", pointerId).Bool()
}

// Is the Pointer outside of this Sprite?
func (self *InputHandler) PointerOutI(args ...interface{}) bool{
    return self.Object.Call("pointerOut", args).Bool()
}

// A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeOver(pointerId int) int{
    return self.Object.Call("pointerTimeOver", pointerId).Int()
}

// A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeOverI(args ...interface{}) int{
    return self.Object.Call("pointerTimeOver", args).Int()
}

// A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeOut(pointerId int) int{
    return self.Object.Call("pointerTimeOut", pointerId).Int()
}

// A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeOutI(args ...interface{}) int{
    return self.Object.Call("pointerTimeOut", args).Int()
}

// Is this sprite being dragged by the mouse or not?
func (self *InputHandler) PointerDragged(pointerId int) bool{
    return self.Object.Call("pointerDragged", pointerId).Bool()
}

// Is this sprite being dragged by the mouse or not?
func (self *InputHandler) PointerDraggedI(args ...interface{}) bool{
    return self.Object.Call("pointerDragged", args).Bool()
}

// Checks if the given pointer is both down and over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerDown(pointer *Pointer, fastTest bool) bool{
    return self.Object.Call("checkPointerDown", pointer, fastTest).Bool()
}

// Checks if the given pointer is both down and over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerDownI(args ...interface{}) bool{
    return self.Object.Call("checkPointerDown", args).Bool()
}

// Checks if the given pointer is over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerOver(pointer *Pointer, fastTest bool) bool{
    return self.Object.Call("checkPointerOver", pointer, fastTest).Bool()
}

// Checks if the given pointer is over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerOverI(args ...interface{}) bool{
    return self.Object.Call("checkPointerOver", args).Bool()
}

// Runs a pixel perfect check against the given x/y coordinates of the Sprite this InputHandler is bound to.
// It compares the alpha value of the pixel and if >= InputHandler.pixelPerfectAlpha it returns true.
func (self *InputHandler) CheckPixel(x int, y int, pointer *Pointer) bool{
    return self.Object.Call("checkPixel", x, y, pointer).Bool()
}

// Runs a pixel perfect check against the given x/y coordinates of the Sprite this InputHandler is bound to.
// It compares the alpha value of the pixel and if >= InputHandler.pixelPerfectAlpha it returns true.
func (self *InputHandler) CheckPixelI(args ...interface{}) bool{
    return self.Object.Call("checkPixel", args).Bool()
}

// Internal Update method. This is called automatically and handles the Pointer
// and drag update loops.
func (self *InputHandler) Update(pointer *Pointer) bool{
    return self.Object.Call("update", pointer).Bool()
}

// Internal Update method. This is called automatically and handles the Pointer
// and drag update loops.
func (self *InputHandler) UpdateI(args ...interface{}) bool{
    return self.Object.Call("update", args).Bool()
}

// Internal method handling the pointer over event.
func (self *InputHandler) _pointerOverHandler(pointer *Pointer, silent bool) {
    self.Object.Call("_pointerOverHandler", pointer, silent)
}

// Internal method handling the pointer over event.
func (self *InputHandler) _pointerOverHandlerI(args ...interface{}) {
    self.Object.Call("_pointerOverHandler", args)
}

// Internal method handling the pointer out event.
func (self *InputHandler) _pointerOutHandler(pointer *Pointer, silent bool) {
    self.Object.Call("_pointerOutHandler", pointer, silent)
}

// Internal method handling the pointer out event.
func (self *InputHandler) _pointerOutHandlerI(args ...interface{}) {
    self.Object.Call("_pointerOutHandler", args)
}

// Internal method handling the touched / clicked event.
func (self *InputHandler) _touchedHandler(pointer *Pointer) {
    self.Object.Call("_touchedHandler", pointer)
}

// Internal method handling the touched / clicked event.
func (self *InputHandler) _touchedHandlerI(args ...interface{}) {
    self.Object.Call("_touchedHandler", args)
}

// Internal method handling the drag threshold timer.
func (self *InputHandler) DragTimeElapsed(pointer *Pointer) {
    self.Object.Call("dragTimeElapsed", pointer)
}

// Internal method handling the drag threshold timer.
func (self *InputHandler) DragTimeElapsedI(args ...interface{}) {
    self.Object.Call("dragTimeElapsed", args)
}

// Internal method handling the pointer released event.
func (self *InputHandler) _releasedHandler(pointer *Pointer) {
    self.Object.Call("_releasedHandler", pointer)
}

// Internal method handling the pointer released event.
func (self *InputHandler) _releasedHandlerI(args ...interface{}) {
    self.Object.Call("_releasedHandler", args)
}

// Called as a Pointer actively drags this Game Object.
func (self *InputHandler) UpdateDrag(pointer *Pointer, fromStart bool) bool{
    return self.Object.Call("updateDrag", pointer, fromStart).Bool()
}

// Called as a Pointer actively drags this Game Object.
func (self *InputHandler) UpdateDragI(args ...interface{}) bool{
    return self.Object.Call("updateDrag", args).Bool()
}

// Returns true if the pointer has entered the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustOver(pointerId int, delay int) bool{
    return self.Object.Call("justOver", pointerId, delay).Bool()
}

// Returns true if the pointer has entered the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustOverI(args ...interface{}) bool{
    return self.Object.Call("justOver", args).Bool()
}

// Returns true if the pointer has left the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustOut(pointerId int, delay int) bool{
    return self.Object.Call("justOut", pointerId, delay).Bool()
}

// Returns true if the pointer has left the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustOutI(args ...interface{}) bool{
    return self.Object.Call("justOut", args).Bool()
}

// Returns true if the pointer has touched or clicked on the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustPressed(pointerId int, delay int) bool{
    return self.Object.Call("justPressed", pointerId, delay).Bool()
}

// Returns true if the pointer has touched or clicked on the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustPressedI(args ...interface{}) bool{
    return self.Object.Call("justPressed", args).Bool()
}

// Returns true if the pointer was touching this Sprite, but has been released within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustReleased(pointerId int, delay int) bool{
    return self.Object.Call("justReleased", pointerId, delay).Bool()
}

// Returns true if the pointer was touching this Sprite, but has been released within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustReleasedI(args ...interface{}) bool{
    return self.Object.Call("justReleased", args).Bool()
}

// If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) OverDuration(pointerId int) int{
    return self.Object.Call("overDuration", pointerId).Int()
}

// If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) OverDurationI(args ...interface{}) int{
    return self.Object.Call("overDuration", args).Int()
}

// If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) DownDuration(pointerId int) int{
    return self.Object.Call("downDuration", pointerId).Int()
}

// If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) DownDurationI(args ...interface{}) int{
    return self.Object.Call("downDuration", args).Int()
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
func (self *InputHandler) EnableDrag(lockCenter bool, bringToTop bool, pixelPerfect bool, alphaThreshold bool, boundsRect *Rectangle, boundsSprite *Sprite) {
    self.Object.Call("enableDrag", lockCenter, bringToTop, pixelPerfect, alphaThreshold, boundsRect, boundsSprite)
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
    self.Object.Call("enableDrag", args)
}

// Stops this sprite from being able to be dragged.
// If it is currently the target of an active drag it will be stopped immediately; also disables any set callbacks.
func (self *InputHandler) DisableDrag() {
    self.Object.Call("disableDrag")
}

// Stops this sprite from being able to be dragged.
// If it is currently the target of an active drag it will be stopped immediately; also disables any set callbacks.
func (self *InputHandler) DisableDragI(args ...interface{}) {
    self.Object.Call("disableDrag", args)
}

// Called by Pointer when drag starts on this Sprite. Should not usually be called directly.
func (self *InputHandler) StartDrag(pointer *Pointer) {
    self.Object.Call("startDrag", pointer)
}

// Called by Pointer when drag starts on this Sprite. Should not usually be called directly.
func (self *InputHandler) StartDragI(args ...interface{}) {
    self.Object.Call("startDrag", args)
}

// Warning: EXPERIMENTAL
func (self *InputHandler) GlobalToLocalX(x int) {
    self.Object.Call("globalToLocalX", x)
}

// Warning: EXPERIMENTAL
func (self *InputHandler) GlobalToLocalXI(args ...interface{}) {
    self.Object.Call("globalToLocalX", args)
}

// Warning: EXPERIMENTAL
func (self *InputHandler) GlobalToLocalY(y int) {
    self.Object.Call("globalToLocalY", y)
}

// Warning: EXPERIMENTAL
func (self *InputHandler) GlobalToLocalYI(args ...interface{}) {
    self.Object.Call("globalToLocalY", args)
}

// Called by Pointer when drag is stopped on this Sprite. Should not usually be called directly.
func (self *InputHandler) StopDrag(pointer *Pointer) {
    self.Object.Call("stopDrag", pointer)
}

// Called by Pointer when drag is stopped on this Sprite. Should not usually be called directly.
func (self *InputHandler) StopDragI(args ...interface{}) {
    self.Object.Call("stopDrag", args)
}

// Restricts this sprite to drag movement only on the given axis. Note: If both are set to false the sprite will never move!
func (self *InputHandler) SetDragLock(allowHorizontal bool, allowVertical bool) {
    self.Object.Call("setDragLock", allowHorizontal, allowVertical)
}

// Restricts this sprite to drag movement only on the given axis. Note: If both are set to false the sprite will never move!
func (self *InputHandler) SetDragLockI(args ...interface{}) {
    self.Object.Call("setDragLock", args)
}

// Make this Sprite snap to the given grid either during drag or when it's released.
// For example 16x16 as the snapX and snapY would make the sprite snap to every 16 pixels.
func (self *InputHandler) EnableSnap(snapX int, snapY int, onDrag bool, onRelease bool, snapOffsetX int, snapOffsetY int) {
    self.Object.Call("enableSnap", snapX, snapY, onDrag, onRelease, snapOffsetX, snapOffsetY)
}

// Make this Sprite snap to the given grid either during drag or when it's released.
// For example 16x16 as the snapX and snapY would make the sprite snap to every 16 pixels.
func (self *InputHandler) EnableSnapI(args ...interface{}) {
    self.Object.Call("enableSnap", args)
}

// Stops the sprite from snapping to a grid during drag or release.
func (self *InputHandler) DisableSnap() {
    self.Object.Call("disableSnap")
}

// Stops the sprite from snapping to a grid during drag or release.
func (self *InputHandler) DisableSnapI(args ...interface{}) {
    self.Object.Call("disableSnap", args)
}

// Bounds Rect check for the sprite drag
func (self *InputHandler) CheckBoundsRect() {
    self.Object.Call("checkBoundsRect")
}

// Bounds Rect check for the sprite drag
func (self *InputHandler) CheckBoundsRectI(args ...interface{}) {
    self.Object.Call("checkBoundsRect", args)
}

// Parent Sprite Bounds check for the sprite drag.
func (self *InputHandler) CheckBoundsSprite() {
    self.Object.Call("checkBoundsSprite")
}

// Parent Sprite Bounds check for the sprite drag.
func (self *InputHandler) CheckBoundsSpriteI(args ...interface{}) {
    self.Object.Call("checkBoundsSprite", args)
}
