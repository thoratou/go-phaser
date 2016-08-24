// Package phaser Automatic generation for Phaser.InputHandler
// generated file InputHandler.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// InputHandler The Input Handler is bound to a specific Sprite and is responsible for managing all Input events on that Sprite.
type InputHandler struct {
    *js.Object
}

// NewInputHandler The Input Handler is bound to a specific Sprite and is responsible for managing all Input events on that Sprite.
func NewInputHandler(sprite *Sprite) *InputHandler {
    return &InputHandler{js.Global.Get("Phaser").Get("InputHandler").New(sprite)}
}
// NewInputHandlerI The Input Handler is bound to a specific Sprite and is responsible for managing all Input events on that Sprite.
func NewInputHandlerI(args ...interface{}) *InputHandler {
    return &InputHandler{js.Global.Get("Phaser").Get("InputHandler").New(args)}
}



// Sprite The Sprite object to which this Input Handler belongs.
func (self *InputHandler) Sprite() *Sprite{
    return &Sprite{self.Object.Get("sprite")}
}

// SetSpriteA The Sprite object to which this Input Handler belongs.
func (self *InputHandler) SetSpriteA(member *Sprite) {
    self.Object.Set("sprite", member)
}

// Game A reference to the currently running game.
func (self *InputHandler) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *InputHandler) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Enabled If enabled the Input Handler will process input requests and monitor pointer activity.
func (self *InputHandler) Enabled() bool{
    return self.Object.Get("enabled").Bool()
}

// SetEnabledA If enabled the Input Handler will process input requests and monitor pointer activity.
func (self *InputHandler) SetEnabledA(member bool) {
    self.Object.Set("enabled", member)
}

// Checked A disposable flag used by the Pointer class when performing priority checks.
func (self *InputHandler) Checked() bool{
    return self.Object.Get("checked").Bool()
}

// SetCheckedA A disposable flag used by the Pointer class when performing priority checks.
func (self *InputHandler) SetCheckedA(member bool) {
    self.Object.Set("checked", member)
}

// PriorityID The priorityID is used to determine which game objects should get priority when input events occur. For example if you have
// several Sprites that overlap, by default the one at the top of the display list is given priority for input events. You can
// stop this from happening by controlling the priorityID value. The higher the value, the more important they are considered to the Input events.
func (self *InputHandler) PriorityID() int{
    return self.Object.Get("priorityID").Int()
}

// SetPriorityIDA The priorityID is used to determine which game objects should get priority when input events occur. For example if you have
// several Sprites that overlap, by default the one at the top of the display list is given priority for input events. You can
// stop this from happening by controlling the priorityID value. The higher the value, the more important they are considered to the Input events.
func (self *InputHandler) SetPriorityIDA(member int) {
    self.Object.Set("priorityID", member)
}

// UseHandCursor On a desktop browser you can set the 'hand' cursor to appear when moving over the Sprite.
func (self *InputHandler) UseHandCursor() bool{
    return self.Object.Get("useHandCursor").Bool()
}

// SetUseHandCursorA On a desktop browser you can set the 'hand' cursor to appear when moving over the Sprite.
func (self *InputHandler) SetUseHandCursorA(member bool) {
    self.Object.Set("useHandCursor", member)
}

// IsDragged true if the Sprite is being currently dragged.
func (self *InputHandler) IsDragged() bool{
    return self.Object.Get("isDragged").Bool()
}

// SetIsDraggedA true if the Sprite is being currently dragged.
func (self *InputHandler) SetIsDraggedA(member bool) {
    self.Object.Set("isDragged", member)
}

// AllowHorizontalDrag Controls if the Sprite is allowed to be dragged horizontally.
func (self *InputHandler) AllowHorizontalDrag() bool{
    return self.Object.Get("allowHorizontalDrag").Bool()
}

// SetAllowHorizontalDragA Controls if the Sprite is allowed to be dragged horizontally.
func (self *InputHandler) SetAllowHorizontalDragA(member bool) {
    self.Object.Set("allowHorizontalDrag", member)
}

// AllowVerticalDrag Controls if the Sprite is allowed to be dragged vertically.
func (self *InputHandler) AllowVerticalDrag() bool{
    return self.Object.Get("allowVerticalDrag").Bool()
}

// SetAllowVerticalDragA Controls if the Sprite is allowed to be dragged vertically.
func (self *InputHandler) SetAllowVerticalDragA(member bool) {
    self.Object.Set("allowVerticalDrag", member)
}

// BringToTop If true when this Sprite is clicked or dragged it will automatically be bought to the top of the Group it is within.
func (self *InputHandler) BringToTop() bool{
    return self.Object.Get("bringToTop").Bool()
}

// SetBringToTopA If true when this Sprite is clicked or dragged it will automatically be bought to the top of the Group it is within.
func (self *InputHandler) SetBringToTopA(member bool) {
    self.Object.Set("bringToTop", member)
}

// SnapOffset A Point object that contains by how far the Sprite snap is offset.
func (self *InputHandler) SnapOffset() *Point{
    return &Point{self.Object.Get("snapOffset")}
}

// SetSnapOffsetA A Point object that contains by how far the Sprite snap is offset.
func (self *InputHandler) SetSnapOffsetA(member *Point) {
    self.Object.Set("snapOffset", member)
}

// SnapOnDrag When the Sprite is dragged this controls if the center of the Sprite will snap to the pointer on drag or not.
func (self *InputHandler) SnapOnDrag() bool{
    return self.Object.Get("snapOnDrag").Bool()
}

// SetSnapOnDragA When the Sprite is dragged this controls if the center of the Sprite will snap to the pointer on drag or not.
func (self *InputHandler) SetSnapOnDragA(member bool) {
    self.Object.Set("snapOnDrag", member)
}

// SnapOnRelease When the Sprite is dragged this controls if the Sprite will be snapped on release.
func (self *InputHandler) SnapOnRelease() bool{
    return self.Object.Get("snapOnRelease").Bool()
}

// SetSnapOnReleaseA When the Sprite is dragged this controls if the Sprite will be snapped on release.
func (self *InputHandler) SetSnapOnReleaseA(member bool) {
    self.Object.Set("snapOnRelease", member)
}

// SnapX When a Sprite has snapping enabled this holds the width of the snap grid.
func (self *InputHandler) SnapX() int{
    return self.Object.Get("snapX").Int()
}

// SetSnapXA When a Sprite has snapping enabled this holds the width of the snap grid.
func (self *InputHandler) SetSnapXA(member int) {
    self.Object.Set("snapX", member)
}

// SnapY When a Sprite has snapping enabled this holds the height of the snap grid.
func (self *InputHandler) SnapY() int{
    return self.Object.Get("snapY").Int()
}

// SetSnapYA When a Sprite has snapping enabled this holds the height of the snap grid.
func (self *InputHandler) SetSnapYA(member int) {
    self.Object.Set("snapY", member)
}

// SnapOffsetX This defines the top-left X coordinate of the snap grid.
func (self *InputHandler) SnapOffsetX() int{
    return self.Object.Get("snapOffsetX").Int()
}

// SetSnapOffsetXA This defines the top-left X coordinate of the snap grid.
func (self *InputHandler) SetSnapOffsetXA(member int) {
    self.Object.Set("snapOffsetX", member)
}

// SnapOffsetY This defines the top-left Y coordinate of the snap grid..
func (self *InputHandler) SnapOffsetY() int{
    return self.Object.Get("snapOffsetY").Int()
}

// SetSnapOffsetYA This defines the top-left Y coordinate of the snap grid..
func (self *InputHandler) SetSnapOffsetYA(member int) {
    self.Object.Set("snapOffsetY", member)
}

// PixelPerfectOver Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive, especially on mobile (where it's not even needed!) so only enable if required. Also see the less-expensive InputHandler.pixelPerfectClick. Use a pixel perfect check when testing for pointer over.
func (self *InputHandler) PixelPerfectOver() bool{
    return self.Object.Get("pixelPerfectOver").Bool()
}

// SetPixelPerfectOverA Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive, especially on mobile (where it's not even needed!) so only enable if required. Also see the less-expensive InputHandler.pixelPerfectClick. Use a pixel perfect check when testing for pointer over.
func (self *InputHandler) SetPixelPerfectOverA(member bool) {
    self.Object.Set("pixelPerfectOver", member)
}

// PixelPerfectClick Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite when it's clicked or touched.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive so only enable if you really need it. Use a pixel perfect check when testing for clicks or touches on the Sprite.
func (self *InputHandler) PixelPerfectClick() bool{
    return self.Object.Get("pixelPerfectClick").Bool()
}

// SetPixelPerfectClickA Set to true to use pixel perfect hit detection when checking if the pointer is over this Sprite when it's clicked or touched.
// The x/y coordinates of the pointer are tested against the image in combination with the InputHandler.pixelPerfectAlpha value.
// This feature only works for display objects with image based textures such as Sprites. It won't work on BitmapText or Rope.
// Warning: This is expensive so only enable if you really need it. Use a pixel perfect check when testing for clicks or touches on the Sprite.
func (self *InputHandler) SetPixelPerfectClickA(member bool) {
    self.Object.Set("pixelPerfectClick", member)
}

// PixelPerfectAlpha The alpha tolerance threshold. If the alpha value of the pixel matches or is above this value, it's considered a hit.
func (self *InputHandler) PixelPerfectAlpha() int{
    return self.Object.Get("pixelPerfectAlpha").Int()
}

// SetPixelPerfectAlphaA The alpha tolerance threshold. If the alpha value of the pixel matches or is above this value, it's considered a hit.
func (self *InputHandler) SetPixelPerfectAlphaA(member int) {
    self.Object.Set("pixelPerfectAlpha", member)
}

// Draggable Is this sprite allowed to be dragged by the mouse? true = yes, false = no
func (self *InputHandler) Draggable() bool{
    return self.Object.Get("draggable").Bool()
}

// SetDraggableA Is this sprite allowed to be dragged by the mouse? true = yes, false = no
func (self *InputHandler) SetDraggableA(member bool) {
    self.Object.Set("draggable", member)
}

// BoundsRect A region of the game world within which the sprite is restricted during drag.
func (self *InputHandler) BoundsRect() *Rectangle{
    return &Rectangle{self.Object.Get("boundsRect")}
}

// SetBoundsRectA A region of the game world within which the sprite is restricted during drag.
func (self *InputHandler) SetBoundsRectA(member *Rectangle) {
    self.Object.Set("boundsRect", member)
}

// BoundsSprite A Sprite the bounds of which this sprite is restricted during drag.
func (self *InputHandler) BoundsSprite() *Sprite{
    return &Sprite{self.Object.Get("boundsSprite")}
}

// SetBoundsSpriteA A Sprite the bounds of which this sprite is restricted during drag.
func (self *InputHandler) SetBoundsSpriteA(member *Sprite) {
    self.Object.Set("boundsSprite", member)
}

// ScaleLayer EXPERIMENTAL: Please do not use this property unless you know what it does. Likely to change in the future.
func (self *InputHandler) ScaleLayer() bool{
    return self.Object.Get("scaleLayer").Bool()
}

// SetScaleLayerA EXPERIMENTAL: Please do not use this property unless you know what it does. Likely to change in the future.
func (self *InputHandler) SetScaleLayerA(member bool) {
    self.Object.Set("scaleLayer", member)
}

// DragOffset The offset from the Sprites position that dragging takes place from.
func (self *InputHandler) DragOffset() *Point{
    return &Point{self.Object.Get("dragOffset")}
}

// SetDragOffsetA The offset from the Sprites position that dragging takes place from.
func (self *InputHandler) SetDragOffsetA(member *Point) {
    self.Object.Set("dragOffset", member)
}

// DragFromCenter Is the Sprite dragged from its center, or the point at which the Pointer was pressed down upon it?
func (self *InputHandler) DragFromCenter() bool{
    return self.Object.Get("dragFromCenter").Bool()
}

// SetDragFromCenterA Is the Sprite dragged from its center, or the point at which the Pointer was pressed down upon it?
func (self *InputHandler) SetDragFromCenterA(member bool) {
    self.Object.Set("dragFromCenter", member)
}

// DragStopBlocksInputUp If enabled, when the Sprite stops being dragged, it will only dispatch the `onDragStop` event, and not the `onInputUp` event. If set to `false` it will dispatch both events.
func (self *InputHandler) DragStopBlocksInputUp() bool{
    return self.Object.Get("dragStopBlocksInputUp").Bool()
}

// SetDragStopBlocksInputUpA If enabled, when the Sprite stops being dragged, it will only dispatch the `onDragStop` event, and not the `onInputUp` event. If set to `false` it will dispatch both events.
func (self *InputHandler) SetDragStopBlocksInputUpA(member bool) {
    self.Object.Set("dragStopBlocksInputUp", member)
}

// DragStartPoint The Point from which the most recent drag started from. Useful if you need to return an object to its starting position.
func (self *InputHandler) DragStartPoint() *Point{
    return &Point{self.Object.Get("dragStartPoint")}
}

// SetDragStartPointA The Point from which the most recent drag started from. Useful if you need to return an object to its starting position.
func (self *InputHandler) SetDragStartPointA(member *Point) {
    self.Object.Set("dragStartPoint", member)
}

// DragDistanceThreshold The distance, in pixels, the pointer has to move while being held down, before the Sprite thinks it is being dragged.
func (self *InputHandler) DragDistanceThreshold() int{
    return self.Object.Get("dragDistanceThreshold").Int()
}

// SetDragDistanceThresholdA The distance, in pixels, the pointer has to move while being held down, before the Sprite thinks it is being dragged.
func (self *InputHandler) SetDragDistanceThresholdA(member int) {
    self.Object.Set("dragDistanceThreshold", member)
}

// DragTimeThreshold The amount of time, in ms, the pointer has to be held down over the Sprite before it thinks it is being dragged.
func (self *InputHandler) DragTimeThreshold() int{
    return self.Object.Get("dragTimeThreshold").Int()
}

// SetDragTimeThresholdA The amount of time, in ms, the pointer has to be held down over the Sprite before it thinks it is being dragged.
func (self *InputHandler) SetDragTimeThresholdA(member int) {
    self.Object.Set("dragTimeThreshold", member)
}

// DownPoint A Point object containing the coordinates of the Pointer when it was first pressed down onto this Sprite.
func (self *InputHandler) DownPoint() *Point{
    return &Point{self.Object.Get("downPoint")}
}

// SetDownPointA A Point object containing the coordinates of the Pointer when it was first pressed down onto this Sprite.
func (self *InputHandler) SetDownPointA(member *Point) {
    self.Object.Set("downPoint", member)
}

// SnapPoint If the sprite is set to snap while dragging this holds the point of the most recent 'snap' event.
func (self *InputHandler) SnapPoint() *Point{
    return &Point{self.Object.Get("snapPoint")}
}

// SetSnapPointA If the sprite is set to snap while dragging this holds the point of the most recent 'snap' event.
func (self *InputHandler) SetSnapPointA(member *Point) {
    self.Object.Set("snapPoint", member)
}


// Start Starts the Input Handler running. This is called automatically when you enable input on a Sprite, or can be called directly if you need to set a specific priority.
func (self *InputHandler) Start() *Sprite{
    return &Sprite{self.Object.Call("start")}
}

// Start1O Starts the Input Handler running. This is called automatically when you enable input on a Sprite, or can be called directly if you need to set a specific priority.
func (self *InputHandler) Start1O(priority int) *Sprite{
    return &Sprite{self.Object.Call("start", priority)}
}

// Start2O Starts the Input Handler running. This is called automatically when you enable input on a Sprite, or can be called directly if you need to set a specific priority.
func (self *InputHandler) Start2O(priority int, useHandCursor bool) *Sprite{
    return &Sprite{self.Object.Call("start", priority, useHandCursor)}
}

// StartI Starts the Input Handler running. This is called automatically when you enable input on a Sprite, or can be called directly if you need to set a specific priority.
func (self *InputHandler) StartI(args ...interface{}) *Sprite{
    return &Sprite{self.Object.Call("start", args)}
}

// AddedToGroup Handles when the parent Sprite is added to a new Group.
func (self *InputHandler) AddedToGroup() {
    self.Object.Call("addedToGroup")
}

// AddedToGroupI Handles when the parent Sprite is added to a new Group.
func (self *InputHandler) AddedToGroupI(args ...interface{}) {
    self.Object.Call("addedToGroup", args)
}

// RemovedFromGroup Handles when the parent Sprite is removed from a Group.
func (self *InputHandler) RemovedFromGroup() {
    self.Object.Call("removedFromGroup")
}

// RemovedFromGroupI Handles when the parent Sprite is removed from a Group.
func (self *InputHandler) RemovedFromGroupI(args ...interface{}) {
    self.Object.Call("removedFromGroup", args)
}

// Reset Resets the Input Handler and disables it.
func (self *InputHandler) Reset() {
    self.Object.Call("reset")
}

// ResetI Resets the Input Handler and disables it.
func (self *InputHandler) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Stop Stops the Input Handler from running.
func (self *InputHandler) Stop() {
    self.Object.Call("stop")
}

// StopI Stops the Input Handler from running.
func (self *InputHandler) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Destroy Clean up memory.
func (self *InputHandler) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Clean up memory.
func (self *InputHandler) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// ValidForInput Checks if the object this InputHandler is bound to is valid for consideration in the Pointer move event.
// This is called by Phaser.Pointer and shouldn't typically be called directly.
func (self *InputHandler) ValidForInput(highestID int, highestRenderID int) bool{
    return self.Object.Call("validForInput", highestID, highestRenderID).Bool()
}

// ValidForInput1O Checks if the object this InputHandler is bound to is valid for consideration in the Pointer move event.
// This is called by Phaser.Pointer and shouldn't typically be called directly.
func (self *InputHandler) ValidForInput1O(highestID int, highestRenderID int, includePixelPerfect bool) bool{
    return self.Object.Call("validForInput", highestID, highestRenderID, includePixelPerfect).Bool()
}

// ValidForInputI Checks if the object this InputHandler is bound to is valid for consideration in the Pointer move event.
// This is called by Phaser.Pointer and shouldn't typically be called directly.
func (self *InputHandler) ValidForInputI(args ...interface{}) bool{
    return self.Object.Call("validForInput", args).Bool()
}

// IsPixelPerfect Is this object using pixel perfect checking?
func (self *InputHandler) IsPixelPerfect() bool{
    return self.Object.Call("isPixelPerfect").Bool()
}

// IsPixelPerfectI Is this object using pixel perfect checking?
func (self *InputHandler) IsPixelPerfectI(args ...interface{}) bool{
    return self.Object.Call("isPixelPerfect", args).Bool()
}

// PointerX The x coordinate of the Input pointer, relative to the top-left of the parent Sprite.
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerX() int{
    return self.Object.Call("pointerX").Int()
}

// PointerX1O The x coordinate of the Input pointer, relative to the top-left of the parent Sprite.
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerX1O(pointerId int) int{
    return self.Object.Call("pointerX", pointerId).Int()
}

// PointerXI The x coordinate of the Input pointer, relative to the top-left of the parent Sprite.
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerXI(args ...interface{}) int{
    return self.Object.Call("pointerX", args).Int()
}

// PointerY The y coordinate of the Input pointer, relative to the top-left of the parent Sprite
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerY() int{
    return self.Object.Call("pointerY").Int()
}

// PointerY1O The y coordinate of the Input pointer, relative to the top-left of the parent Sprite
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerY1O(pointerId int) int{
    return self.Object.Call("pointerY", pointerId).Int()
}

// PointerYI The y coordinate of the Input pointer, relative to the top-left of the parent Sprite
// This value is only set when the pointer is over this Sprite.
func (self *InputHandler) PointerYI(args ...interface{}) int{
    return self.Object.Call("pointerY", args).Int()
}

// PointerDown If the Pointer is down this returns true.
// This *only* checks if the Pointer is down, not if it's down over any specific Sprite.
func (self *InputHandler) PointerDown() bool{
    return self.Object.Call("pointerDown").Bool()
}

// PointerDown1O If the Pointer is down this returns true.
// This *only* checks if the Pointer is down, not if it's down over any specific Sprite.
func (self *InputHandler) PointerDown1O(pointerId int) bool{
    return self.Object.Call("pointerDown", pointerId).Bool()
}

// PointerDownI If the Pointer is down this returns true.
// This *only* checks if the Pointer is down, not if it's down over any specific Sprite.
func (self *InputHandler) PointerDownI(args ...interface{}) bool{
    return self.Object.Call("pointerDown", args).Bool()
}

// PointerUp If the Pointer is up this returns true.
// This *only* checks if the Pointer is up, not if it's up over any specific Sprite.
func (self *InputHandler) PointerUp() bool{
    return self.Object.Call("pointerUp").Bool()
}

// PointerUp1O If the Pointer is up this returns true.
// This *only* checks if the Pointer is up, not if it's up over any specific Sprite.
func (self *InputHandler) PointerUp1O(pointerId int) bool{
    return self.Object.Call("pointerUp", pointerId).Bool()
}

// PointerUpI If the Pointer is up this returns true.
// This *only* checks if the Pointer is up, not if it's up over any specific Sprite.
func (self *InputHandler) PointerUpI(args ...interface{}) bool{
    return self.Object.Call("pointerUp", args).Bool()
}

// PointerTimeDown A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeDown() int{
    return self.Object.Call("pointerTimeDown").Int()
}

// PointerTimeDown1O A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeDown1O(pointerId int) int{
    return self.Object.Call("pointerTimeDown", pointerId).Int()
}

// PointerTimeDownI A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeDownI(args ...interface{}) int{
    return self.Object.Call("pointerTimeDown", args).Int()
}

// PointerTimeUp A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeUp() int{
    return self.Object.Call("pointerTimeUp").Int()
}

// PointerTimeUp1O A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeUp1O(pointerId int) int{
    return self.Object.Call("pointerTimeUp", pointerId).Int()
}

// PointerTimeUpI A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeUpI(args ...interface{}) int{
    return self.Object.Call("pointerTimeUp", args).Int()
}

// PointerOver Is the Pointer over this Sprite?
func (self *InputHandler) PointerOver() bool{
    return self.Object.Call("pointerOver").Bool()
}

// PointerOver1O Is the Pointer over this Sprite?
func (self *InputHandler) PointerOver1O(pointerId int) bool{
    return self.Object.Call("pointerOver", pointerId).Bool()
}

// PointerOverI Is the Pointer over this Sprite?
func (self *InputHandler) PointerOverI(args ...interface{}) bool{
    return self.Object.Call("pointerOver", args).Bool()
}

// PointerOut Is the Pointer outside of this Sprite?
func (self *InputHandler) PointerOut() bool{
    return self.Object.Call("pointerOut").Bool()
}

// PointerOut1O Is the Pointer outside of this Sprite?
func (self *InputHandler) PointerOut1O(pointerId int) bool{
    return self.Object.Call("pointerOut", pointerId).Bool()
}

// PointerOutI Is the Pointer outside of this Sprite?
func (self *InputHandler) PointerOutI(args ...interface{}) bool{
    return self.Object.Call("pointerOut", args).Bool()
}

// PointerTimeOver A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeOver() int{
    return self.Object.Call("pointerTimeOver").Int()
}

// PointerTimeOver1O A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeOver1O(pointerId int) int{
    return self.Object.Call("pointerTimeOver", pointerId).Int()
}

// PointerTimeOverI A timestamp representing when the Pointer first touched the touchscreen.
func (self *InputHandler) PointerTimeOverI(args ...interface{}) int{
    return self.Object.Call("pointerTimeOver", args).Int()
}

// PointerTimeOut A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeOut() int{
    return self.Object.Call("pointerTimeOut").Int()
}

// PointerTimeOut1O A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeOut1O(pointerId int) int{
    return self.Object.Call("pointerTimeOut", pointerId).Int()
}

// PointerTimeOutI A timestamp representing when the Pointer left the touchscreen.
func (self *InputHandler) PointerTimeOutI(args ...interface{}) int{
    return self.Object.Call("pointerTimeOut", args).Int()
}

// PointerDragged Is this sprite being dragged by the mouse or not?
func (self *InputHandler) PointerDragged() bool{
    return self.Object.Call("pointerDragged").Bool()
}

// PointerDragged1O Is this sprite being dragged by the mouse or not?
func (self *InputHandler) PointerDragged1O(pointerId int) bool{
    return self.Object.Call("pointerDragged", pointerId).Bool()
}

// PointerDraggedI Is this sprite being dragged by the mouse or not?
func (self *InputHandler) PointerDraggedI(args ...interface{}) bool{
    return self.Object.Call("pointerDragged", args).Bool()
}

// CheckPointerDown Checks if the given pointer is both down and over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerDown(pointer *Pointer) bool{
    return self.Object.Call("checkPointerDown", pointer).Bool()
}

// CheckPointerDown1O Checks if the given pointer is both down and over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerDown1O(pointer *Pointer, fastTest bool) bool{
    return self.Object.Call("checkPointerDown", pointer, fastTest).Bool()
}

// CheckPointerDownI Checks if the given pointer is both down and over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerDownI(args ...interface{}) bool{
    return self.Object.Call("checkPointerDown", args).Bool()
}

// CheckPointerOver Checks if the given pointer is over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerOver(pointer *Pointer) bool{
    return self.Object.Call("checkPointerOver", pointer).Bool()
}

// CheckPointerOver1O Checks if the given pointer is over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerOver1O(pointer *Pointer, fastTest bool) bool{
    return self.Object.Call("checkPointerOver", pointer, fastTest).Bool()
}

// CheckPointerOverI Checks if the given pointer is over the Sprite this InputHandler belongs to.
// Use the `fastTest` flag is to quickly check just the bounding hit area even if `InputHandler.pixelPerfectOver` is `true`.
func (self *InputHandler) CheckPointerOverI(args ...interface{}) bool{
    return self.Object.Call("checkPointerOver", args).Bool()
}

// CheckPixel Runs a pixel perfect check against the given x/y coordinates of the Sprite this InputHandler is bound to.
// It compares the alpha value of the pixel and if >= InputHandler.pixelPerfectAlpha it returns true.
func (self *InputHandler) CheckPixel(x int, y int) bool{
    return self.Object.Call("checkPixel", x, y).Bool()
}

// CheckPixel1O Runs a pixel perfect check against the given x/y coordinates of the Sprite this InputHandler is bound to.
// It compares the alpha value of the pixel and if >= InputHandler.pixelPerfectAlpha it returns true.
func (self *InputHandler) CheckPixel1O(x int, y int, pointer *Pointer) bool{
    return self.Object.Call("checkPixel", x, y, pointer).Bool()
}

// CheckPixelI Runs a pixel perfect check against the given x/y coordinates of the Sprite this InputHandler is bound to.
// It compares the alpha value of the pixel and if >= InputHandler.pixelPerfectAlpha it returns true.
func (self *InputHandler) CheckPixelI(args ...interface{}) bool{
    return self.Object.Call("checkPixel", args).Bool()
}

// Update Internal Update method. This is called automatically and handles the Pointer
// and drag update loops.
func (self *InputHandler) Update(pointer *Pointer) bool{
    return self.Object.Call("update", pointer).Bool()
}

// UpdateI Internal Update method. This is called automatically and handles the Pointer
// and drag update loops.
func (self *InputHandler) UpdateI(args ...interface{}) bool{
    return self.Object.Call("update", args).Bool()
}

// _pointerOverHandler Internal method handling the pointer over event.
func (self *InputHandler) _pointerOverHandler(pointer *Pointer) {
    self.Object.Call("_pointerOverHandler", pointer)
}

// _pointerOverHandler1O Internal method handling the pointer over event.
func (self *InputHandler) _pointerOverHandler1O(pointer *Pointer, silent bool) {
    self.Object.Call("_pointerOverHandler", pointer, silent)
}

// _pointerOverHandlerI Internal method handling the pointer over event.
func (self *InputHandler) _pointerOverHandlerI(args ...interface{}) {
    self.Object.Call("_pointerOverHandler", args)
}

// _pointerOutHandler Internal method handling the pointer out event.
func (self *InputHandler) _pointerOutHandler(pointer *Pointer) {
    self.Object.Call("_pointerOutHandler", pointer)
}

// _pointerOutHandler1O Internal method handling the pointer out event.
func (self *InputHandler) _pointerOutHandler1O(pointer *Pointer, silent bool) {
    self.Object.Call("_pointerOutHandler", pointer, silent)
}

// _pointerOutHandlerI Internal method handling the pointer out event.
func (self *InputHandler) _pointerOutHandlerI(args ...interface{}) {
    self.Object.Call("_pointerOutHandler", args)
}

// _touchedHandler Internal method handling the touched / clicked event.
func (self *InputHandler) _touchedHandler(pointer *Pointer) {
    self.Object.Call("_touchedHandler", pointer)
}

// _touchedHandlerI Internal method handling the touched / clicked event.
func (self *InputHandler) _touchedHandlerI(args ...interface{}) {
    self.Object.Call("_touchedHandler", args)
}

// DragTimeElapsed Internal method handling the drag threshold timer.
func (self *InputHandler) DragTimeElapsed(pointer *Pointer) {
    self.Object.Call("dragTimeElapsed", pointer)
}

// DragTimeElapsedI Internal method handling the drag threshold timer.
func (self *InputHandler) DragTimeElapsedI(args ...interface{}) {
    self.Object.Call("dragTimeElapsed", args)
}

// _releasedHandler Internal method handling the pointer released event.
func (self *InputHandler) _releasedHandler(pointer *Pointer) {
    self.Object.Call("_releasedHandler", pointer)
}

// _releasedHandlerI Internal method handling the pointer released event.
func (self *InputHandler) _releasedHandlerI(args ...interface{}) {
    self.Object.Call("_releasedHandler", args)
}

// UpdateDrag Called as a Pointer actively drags this Game Object.
func (self *InputHandler) UpdateDrag(pointer *Pointer, fromStart bool) bool{
    return self.Object.Call("updateDrag", pointer, fromStart).Bool()
}

// UpdateDragI Called as a Pointer actively drags this Game Object.
func (self *InputHandler) UpdateDragI(args ...interface{}) bool{
    return self.Object.Call("updateDrag", args).Bool()
}

// JustOver Returns true if the pointer has entered the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustOver(pointerId int, delay int) bool{
    return self.Object.Call("justOver", pointerId, delay).Bool()
}

// JustOverI Returns true if the pointer has entered the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustOverI(args ...interface{}) bool{
    return self.Object.Call("justOver", args).Bool()
}

// JustOut Returns true if the pointer has left the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustOut(pointerId int, delay int) bool{
    return self.Object.Call("justOut", pointerId, delay).Bool()
}

// JustOutI Returns true if the pointer has left the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustOutI(args ...interface{}) bool{
    return self.Object.Call("justOut", args).Bool()
}

// JustPressed Returns true if the pointer has touched or clicked on the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustPressed(pointerId int, delay int) bool{
    return self.Object.Call("justPressed", pointerId, delay).Bool()
}

// JustPressedI Returns true if the pointer has touched or clicked on the Sprite within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustPressedI(args ...interface{}) bool{
    return self.Object.Call("justPressed", args).Bool()
}

// JustReleased Returns true if the pointer was touching this Sprite, but has been released within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustReleased(pointerId int, delay int) bool{
    return self.Object.Call("justReleased", pointerId, delay).Bool()
}

// JustReleasedI Returns true if the pointer was touching this Sprite, but has been released within the specified delay time (defaults to 500ms, half a second)
func (self *InputHandler) JustReleasedI(args ...interface{}) bool{
    return self.Object.Call("justReleased", args).Bool()
}

// OverDuration If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) OverDuration() int{
    return self.Object.Call("overDuration").Int()
}

// OverDuration1O If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) OverDuration1O(pointerId int) int{
    return self.Object.Call("overDuration", pointerId).Int()
}

// OverDurationI If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) OverDurationI(args ...interface{}) int{
    return self.Object.Call("overDuration", args).Int()
}

// DownDuration If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) DownDuration() int{
    return self.Object.Call("downDuration").Int()
}

// DownDuration1O If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) DownDuration1O(pointerId int) int{
    return self.Object.Call("downDuration", pointerId).Int()
}

// DownDurationI If the pointer is currently over this Sprite this returns how long it has been there for in milliseconds.
func (self *InputHandler) DownDurationI(args ...interface{}) int{
    return self.Object.Call("downDuration", args).Int()
}

// EnableDrag Allow this Sprite to be dragged by any valid pointer.
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
func (self *InputHandler) EnableDrag() {
    self.Object.Call("enableDrag")
}

// EnableDrag1O Allow this Sprite to be dragged by any valid pointer.
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
func (self *InputHandler) EnableDrag1O(lockCenter bool) {
    self.Object.Call("enableDrag", lockCenter)
}

// EnableDrag2O Allow this Sprite to be dragged by any valid pointer.
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
func (self *InputHandler) EnableDrag2O(lockCenter bool, bringToTop bool) {
    self.Object.Call("enableDrag", lockCenter, bringToTop)
}

// EnableDrag3O Allow this Sprite to be dragged by any valid pointer.
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
func (self *InputHandler) EnableDrag3O(lockCenter bool, bringToTop bool, pixelPerfect bool) {
    self.Object.Call("enableDrag", lockCenter, bringToTop, pixelPerfect)
}

// EnableDrag4O Allow this Sprite to be dragged by any valid pointer.
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
func (self *InputHandler) EnableDrag4O(lockCenter bool, bringToTop bool, pixelPerfect bool, alphaThreshold bool) {
    self.Object.Call("enableDrag", lockCenter, bringToTop, pixelPerfect, alphaThreshold)
}

// EnableDrag5O Allow this Sprite to be dragged by any valid pointer.
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
func (self *InputHandler) EnableDrag5O(lockCenter bool, bringToTop bool, pixelPerfect bool, alphaThreshold bool, boundsRect *Rectangle) {
    self.Object.Call("enableDrag", lockCenter, bringToTop, pixelPerfect, alphaThreshold, boundsRect)
}

// EnableDrag6O Allow this Sprite to be dragged by any valid pointer.
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
func (self *InputHandler) EnableDrag6O(lockCenter bool, bringToTop bool, pixelPerfect bool, alphaThreshold bool, boundsRect *Rectangle, boundsSprite *Sprite) {
    self.Object.Call("enableDrag", lockCenter, bringToTop, pixelPerfect, alphaThreshold, boundsRect, boundsSprite)
}

// EnableDragI Allow this Sprite to be dragged by any valid pointer.
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

// DisableDrag Stops this sprite from being able to be dragged.
// If it is currently the target of an active drag it will be stopped immediately; also disables any set callbacks.
func (self *InputHandler) DisableDrag() {
    self.Object.Call("disableDrag")
}

// DisableDragI Stops this sprite from being able to be dragged.
// If it is currently the target of an active drag it will be stopped immediately; also disables any set callbacks.
func (self *InputHandler) DisableDragI(args ...interface{}) {
    self.Object.Call("disableDrag", args)
}

// StartDrag Called by Pointer when drag starts on this Sprite. Should not usually be called directly.
func (self *InputHandler) StartDrag(pointer *Pointer) {
    self.Object.Call("startDrag", pointer)
}

// StartDragI Called by Pointer when drag starts on this Sprite. Should not usually be called directly.
func (self *InputHandler) StartDragI(args ...interface{}) {
    self.Object.Call("startDrag", args)
}

// GlobalToLocalX Warning: EXPERIMENTAL
func (self *InputHandler) GlobalToLocalX(x int) {
    self.Object.Call("globalToLocalX", x)
}

// GlobalToLocalXI Warning: EXPERIMENTAL
func (self *InputHandler) GlobalToLocalXI(args ...interface{}) {
    self.Object.Call("globalToLocalX", args)
}

// GlobalToLocalY Warning: EXPERIMENTAL
func (self *InputHandler) GlobalToLocalY(y int) {
    self.Object.Call("globalToLocalY", y)
}

// GlobalToLocalYI Warning: EXPERIMENTAL
func (self *InputHandler) GlobalToLocalYI(args ...interface{}) {
    self.Object.Call("globalToLocalY", args)
}

// StopDrag Called by Pointer when drag is stopped on this Sprite. Should not usually be called directly.
func (self *InputHandler) StopDrag(pointer *Pointer) {
    self.Object.Call("stopDrag", pointer)
}

// StopDragI Called by Pointer when drag is stopped on this Sprite. Should not usually be called directly.
func (self *InputHandler) StopDragI(args ...interface{}) {
    self.Object.Call("stopDrag", args)
}

// SetDragLock Restricts this sprite to drag movement only on the given axis. Note: If both are set to false the sprite will never move!
func (self *InputHandler) SetDragLock() {
    self.Object.Call("setDragLock")
}

// SetDragLock1O Restricts this sprite to drag movement only on the given axis. Note: If both are set to false the sprite will never move!
func (self *InputHandler) SetDragLock1O(allowHorizontal bool) {
    self.Object.Call("setDragLock", allowHorizontal)
}

// SetDragLock2O Restricts this sprite to drag movement only on the given axis. Note: If both are set to false the sprite will never move!
func (self *InputHandler) SetDragLock2O(allowHorizontal bool, allowVertical bool) {
    self.Object.Call("setDragLock", allowHorizontal, allowVertical)
}

// SetDragLockI Restricts this sprite to drag movement only on the given axis. Note: If both are set to false the sprite will never move!
func (self *InputHandler) SetDragLockI(args ...interface{}) {
    self.Object.Call("setDragLock", args)
}

// EnableSnap Make this Sprite snap to the given grid either during drag or when it's released.
// For example 16x16 as the snapX and snapY would make the sprite snap to every 16 pixels.
func (self *InputHandler) EnableSnap(snapX int, snapY int) {
    self.Object.Call("enableSnap", snapX, snapY)
}

// EnableSnap1O Make this Sprite snap to the given grid either during drag or when it's released.
// For example 16x16 as the snapX and snapY would make the sprite snap to every 16 pixels.
func (self *InputHandler) EnableSnap1O(snapX int, snapY int, onDrag bool) {
    self.Object.Call("enableSnap", snapX, snapY, onDrag)
}

// EnableSnap2O Make this Sprite snap to the given grid either during drag or when it's released.
// For example 16x16 as the snapX and snapY would make the sprite snap to every 16 pixels.
func (self *InputHandler) EnableSnap2O(snapX int, snapY int, onDrag bool, onRelease bool) {
    self.Object.Call("enableSnap", snapX, snapY, onDrag, onRelease)
}

// EnableSnap3O Make this Sprite snap to the given grid either during drag or when it's released.
// For example 16x16 as the snapX and snapY would make the sprite snap to every 16 pixels.
func (self *InputHandler) EnableSnap3O(snapX int, snapY int, onDrag bool, onRelease bool, snapOffsetX int) {
    self.Object.Call("enableSnap", snapX, snapY, onDrag, onRelease, snapOffsetX)
}

// EnableSnap4O Make this Sprite snap to the given grid either during drag or when it's released.
// For example 16x16 as the snapX and snapY would make the sprite snap to every 16 pixels.
func (self *InputHandler) EnableSnap4O(snapX int, snapY int, onDrag bool, onRelease bool, snapOffsetX int, snapOffsetY int) {
    self.Object.Call("enableSnap", snapX, snapY, onDrag, onRelease, snapOffsetX, snapOffsetY)
}

// EnableSnapI Make this Sprite snap to the given grid either during drag or when it's released.
// For example 16x16 as the snapX and snapY would make the sprite snap to every 16 pixels.
func (self *InputHandler) EnableSnapI(args ...interface{}) {
    self.Object.Call("enableSnap", args)
}

// DisableSnap Stops the sprite from snapping to a grid during drag or release.
func (self *InputHandler) DisableSnap() {
    self.Object.Call("disableSnap")
}

// DisableSnapI Stops the sprite from snapping to a grid during drag or release.
func (self *InputHandler) DisableSnapI(args ...interface{}) {
    self.Object.Call("disableSnap", args)
}

// CheckBoundsRect Bounds Rect check for the sprite drag
func (self *InputHandler) CheckBoundsRect() {
    self.Object.Call("checkBoundsRect")
}

// CheckBoundsRectI Bounds Rect check for the sprite drag
func (self *InputHandler) CheckBoundsRectI(args ...interface{}) {
    self.Object.Call("checkBoundsRect", args)
}

// CheckBoundsSprite Parent Sprite Bounds check for the sprite drag.
func (self *InputHandler) CheckBoundsSprite() {
    self.Object.Call("checkBoundsSprite")
}

// CheckBoundsSpriteI Parent Sprite Bounds check for the sprite drag.
func (self *InputHandler) CheckBoundsSpriteI(args ...interface{}) {
    self.Object.Call("checkBoundsSprite", args)
}

