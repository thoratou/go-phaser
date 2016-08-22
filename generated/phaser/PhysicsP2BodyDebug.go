// Automatic generation for Phaser.Physics.P2.BodyDebug
// generated file PhysicsP2BodyDebug.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Draws a P2 Body to a Graphics instance for visual debugging.
// Needless to say, for every body you enable debug drawing on, you are adding processor and graphical overhead.
// So use sparingly and rarely (if ever) in production code.
// 
// Also be aware that the Debug body is only updated when the Sprite it is connected to changes position. If you
// manipulate the sprite in any other way (such as moving it to another Group or bringToTop, etc) then you will
// need to manually adjust its BodyDebug as well.
type PhysicsP2BodyDebug struct {
    *js.Object
}


// Pixels per Length Unit.
func (self *PhysicsP2BodyDebug) GetPpuA() int{
    return self.Object.Get("ppu").Int()
}

// Pixels per Length Unit.
func (self *PhysicsP2BodyDebug) SetPpuA(member int) {
    self.Object.Set("ppu", member)
}

// The P2 Body to display debug data for.
func (self *PhysicsP2BodyDebug) GetBodyA() *PhysicsP2Body{
    return &PhysicsP2Body{self.Object.Get("body")}
}

// The P2 Body to display debug data for.
func (self *PhysicsP2BodyDebug) SetBodyA(member *PhysicsP2Body) {
    self.Object.Set("body", member)
}

// The canvas to render the debug info to.
func (self *PhysicsP2BodyDebug) GetCanvasA() *Graphics{
    return &Graphics{self.Object.Get("canvas")}
}

// The canvas to render the debug info to.
func (self *PhysicsP2BodyDebug) SetCanvasA(member *Graphics) {
    self.Object.Set("canvas", member)
}

// A reference to the currently running Game.
func (self *PhysicsP2BodyDebug) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *PhysicsP2BodyDebug) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// A name for this group. Not used internally but useful for debugging.
func (self *PhysicsP2BodyDebug) GetNameA() string{
    return self.Object.Get("name").String()
}

// A name for this group. Not used internally but useful for debugging.
func (self *PhysicsP2BodyDebug) SetNameA(member string) {
    self.Object.Set("name", member)
}

// The z-depth value of this object within its parent container/Group - the World is a Group as well.
// This value must be unique for each child in a Group.
func (self *PhysicsP2BodyDebug) GetZA() int{
    return self.Object.Get("z").Int()
}

// The z-depth value of this object within its parent container/Group - the World is a Group as well.
// This value must be unique for each child in a Group.
func (self *PhysicsP2BodyDebug) SetZA(member int) {
    self.Object.Set("z", member)
}

// Internal Phaser Type value.
func (self *PhysicsP2BodyDebug) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// Internal Phaser Type value.
func (self *PhysicsP2BodyDebug) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The const physics body type of this object.
func (self *PhysicsP2BodyDebug) GetPhysicsTypeA() int{
    return self.Object.Get("physicsType").Int()
}

// The const physics body type of this object.
func (self *PhysicsP2BodyDebug) SetPhysicsTypeA(member int) {
    self.Object.Set("physicsType", member)
}

// The alive property is useful for Groups that are children of other Groups and need to be included/excluded in checks like forEachAlive.
func (self *PhysicsP2BodyDebug) GetAliveA() bool{
    return self.Object.Get("alive").Bool()
}

// The alive property is useful for Groups that are children of other Groups and need to be included/excluded in checks like forEachAlive.
func (self *PhysicsP2BodyDebug) SetAliveA(member bool) {
    self.Object.Set("alive", member)
}

// If exists is true the group is updated, otherwise it is skipped.
func (self *PhysicsP2BodyDebug) GetExistsA() bool{
    return self.Object.Get("exists").Bool()
}

// If exists is true the group is updated, otherwise it is skipped.
func (self *PhysicsP2BodyDebug) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// A group with `ignoreDestroy` set to `true` ignores all calls to its `destroy` method.
func (self *PhysicsP2BodyDebug) GetIgnoreDestroyA() bool{
    return self.Object.Get("ignoreDestroy").Bool()
}

// A group with `ignoreDestroy` set to `true` ignores all calls to its `destroy` method.
func (self *PhysicsP2BodyDebug) SetIgnoreDestroyA(member bool) {
    self.Object.Set("ignoreDestroy", member)
}

// A Group is that has `pendingDestroy` set to `true` is flagged to have its destroy method
// called on the next logic update.
// You can set it directly to flag the Group to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy a Group from within one of its own callbacks
// or a callback of one of its children.
func (self *PhysicsP2BodyDebug) GetPendingDestroyA() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// A Group is that has `pendingDestroy` set to `true` is flagged to have its destroy method
// called on the next logic update.
// You can set it directly to flag the Group to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy a Group from within one of its own callbacks
// or a callback of one of its children.
func (self *PhysicsP2BodyDebug) SetPendingDestroyA(member bool) {
    self.Object.Set("pendingDestroy", member)
}

// The type of objects that will be created when using {@link Phaser.Group#create create} or {@link Phaser.Group#createMultiple createMultiple}.
// 
// Any object may be used but it should extend either Sprite or Image and accept the same constructor arguments:
// when a new object is created it is passed the following parameters to its constructor: `(game, x, y, key, frame)`.
func (self *PhysicsP2BodyDebug) GetClassTypeA() interface{}{
    return self.Object.Get("classType")
}

// The type of objects that will be created when using {@link Phaser.Group#create create} or {@link Phaser.Group#createMultiple createMultiple}.
// 
// Any object may be used but it should extend either Sprite or Image and accept the same constructor arguments:
// when a new object is created it is passed the following parameters to its constructor: `(game, x, y, key, frame)`.
func (self *PhysicsP2BodyDebug) SetClassTypeA(member interface{}) {
    self.Object.Set("classType", member)
}

// The current display object that the group cursor is pointing to, if any. (Can be set manually.)
// 
// The cursor is a way to iterate through the children in a Group using {@link Phaser.Group#next next} and {@link Phaser.Group#previous previous}.
func (self *PhysicsP2BodyDebug) GetCursorA() *DisplayObject{
    return &DisplayObject{self.Object.Get("cursor")}
}

// The current display object that the group cursor is pointing to, if any. (Can be set manually.)
// 
// The cursor is a way to iterate through the children in a Group using {@link Phaser.Group#next next} and {@link Phaser.Group#previous previous}.
func (self *PhysicsP2BodyDebug) SetCursorA(member *DisplayObject) {
    self.Object.Set("cursor", member)
}

// A Group with `inputEnableChildren` set to `true` will automatically call `inputEnabled = true` 
// on any children _added_ to, or _created by_, this Group.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
func (self *PhysicsP2BodyDebug) GetInputEnableChildrenA() bool{
    return self.Object.Get("inputEnableChildren").Bool()
}

// A Group with `inputEnableChildren` set to `true` will automatically call `inputEnabled = true` 
// on any children _added_ to, or _created by_, this Group.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
func (self *PhysicsP2BodyDebug) SetInputEnableChildrenA(member bool) {
    self.Object.Set("inputEnableChildren", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputDown signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *PhysicsP2BodyDebug) GetOnChildInputDownA() *Signal{
    return &Signal{self.Object.Get("onChildInputDown")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputDown signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *PhysicsP2BodyDebug) SetOnChildInputDownA(member *Signal) {
    self.Object.Set("onChildInputDown", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputUp signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 3 arguments: A reference to the Sprite that triggered the signal, 
// a reference to the Pointer that caused it, and a boolean value `isOver` that tells you if the Pointer
// is still over the Sprite or not.
func (self *PhysicsP2BodyDebug) GetOnChildInputUpA() *Signal{
    return &Signal{self.Object.Get("onChildInputUp")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputUp signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 3 arguments: A reference to the Sprite that triggered the signal, 
// a reference to the Pointer that caused it, and a boolean value `isOver` that tells you if the Pointer
// is still over the Sprite or not.
func (self *PhysicsP2BodyDebug) SetOnChildInputUpA(member *Signal) {
    self.Object.Set("onChildInputUp", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputOver signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *PhysicsP2BodyDebug) GetOnChildInputOverA() *Signal{
    return &Signal{self.Object.Get("onChildInputOver")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputOver signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *PhysicsP2BodyDebug) SetOnChildInputOverA(member *Signal) {
    self.Object.Set("onChildInputOver", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputOut signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *PhysicsP2BodyDebug) GetOnChildInputOutA() *Signal{
    return &Signal{self.Object.Get("onChildInputOut")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputOut signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *PhysicsP2BodyDebug) SetOnChildInputOutA(member *Signal) {
    self.Object.Set("onChildInputOut", member)
}

// If true all Sprites created by, or added to this group, will have a physics body enabled on them.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
// 
// The default body type is controlled with {@link Phaser.Group#physicsBodyType physicsBodyType}.
func (self *PhysicsP2BodyDebug) GetEnableBodyA() bool{
    return self.Object.Get("enableBody").Bool()
}

// If true all Sprites created by, or added to this group, will have a physics body enabled on them.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
// 
// The default body type is controlled with {@link Phaser.Group#physicsBodyType physicsBodyType}.
func (self *PhysicsP2BodyDebug) SetEnableBodyA(member bool) {
    self.Object.Set("enableBody", member)
}

// If true when a physics body is created (via {@link Phaser.Group#enableBody enableBody}) it will create a physics debug object as well.
// 
// This only works for P2 bodies.
func (self *PhysicsP2BodyDebug) GetEnableBodyDebugA() bool{
    return self.Object.Get("enableBodyDebug").Bool()
}

// If true when a physics body is created (via {@link Phaser.Group#enableBody enableBody}) it will create a physics debug object as well.
// 
// This only works for P2 bodies.
func (self *PhysicsP2BodyDebug) SetEnableBodyDebugA(member bool) {
    self.Object.Set("enableBodyDebug", member)
}

// If {@link Phaser.Group#enableBody enableBody} is true this is the type of physics body that is created on new Sprites.
// 
// The valid values are {@link Phaser.Physics.ARCADE}, {@link Phaser.Physics.P2JS}, {@link Phaser.Physics.NINJA}, etc.
func (self *PhysicsP2BodyDebug) GetPhysicsBodyTypeA() int{
    return self.Object.Get("physicsBodyType").Int()
}

// If {@link Phaser.Group#enableBody enableBody} is true this is the type of physics body that is created on new Sprites.
// 
// The valid values are {@link Phaser.Physics.ARCADE}, {@link Phaser.Physics.P2JS}, {@link Phaser.Physics.NINJA}, etc.
func (self *PhysicsP2BodyDebug) SetPhysicsBodyTypeA(member int) {
    self.Object.Set("physicsBodyType", member)
}

// If this Group contains Arcade Physics Sprites you can set a custom sort direction via this property.
// 
// It should be set to one of the Phaser.Physics.Arcade sort direction constants: 
// 
// Phaser.Physics.Arcade.SORT_NONE
// Phaser.Physics.Arcade.LEFT_RIGHT
// Phaser.Physics.Arcade.RIGHT_LEFT
// Phaser.Physics.Arcade.TOP_BOTTOM
// Phaser.Physics.Arcade.BOTTOM_TOP
// 
// If set to `null` the Group will use whatever Phaser.Physics.Arcade.sortDirection is set to. This is the default behavior.
func (self *PhysicsP2BodyDebug) GetPhysicsSortDirectionA() int{
    return self.Object.Get("physicsSortDirection").Int()
}

// If this Group contains Arcade Physics Sprites you can set a custom sort direction via this property.
// 
// It should be set to one of the Phaser.Physics.Arcade sort direction constants: 
// 
// Phaser.Physics.Arcade.SORT_NONE
// Phaser.Physics.Arcade.LEFT_RIGHT
// Phaser.Physics.Arcade.RIGHT_LEFT
// Phaser.Physics.Arcade.TOP_BOTTOM
// Phaser.Physics.Arcade.BOTTOM_TOP
// 
// If set to `null` the Group will use whatever Phaser.Physics.Arcade.sortDirection is set to. This is the default behavior.
func (self *PhysicsP2BodyDebug) SetPhysicsSortDirectionA(member int) {
    self.Object.Set("physicsSortDirection", member)
}

// This signal is dispatched when the group is destroyed.
func (self *PhysicsP2BodyDebug) GetOnDestroyA() *Signal{
    return &Signal{self.Object.Get("onDestroy")}
}

// This signal is dispatched when the group is destroyed.
func (self *PhysicsP2BodyDebug) SetOnDestroyA(member *Signal) {
    self.Object.Set("onDestroy", member)
}

// The current index of the Group cursor. Advance it with Group.next.
func (self *PhysicsP2BodyDebug) GetCursorIndexA() int{
    return self.Object.Get("cursorIndex").Int()
}

// The current index of the Group cursor. Advance it with Group.next.
func (self *PhysicsP2BodyDebug) SetCursorIndexA(member int) {
    self.Object.Set("cursorIndex", member)
}

// A Group that is fixed to the camera uses its x/y coordinates as offsets from the top left of the camera. These are stored in Group.cameraOffset.
// 
// Note that the cameraOffset values are in addition to any parent in the display list.
// So if this Group was in a Group that has x: 200, then this will be added to the cameraOffset.x
func (self *PhysicsP2BodyDebug) GetFixedToCameraA() bool{
    return self.Object.Get("fixedToCamera").Bool()
}

// A Group that is fixed to the camera uses its x/y coordinates as offsets from the top left of the camera. These are stored in Group.cameraOffset.
// 
// Note that the cameraOffset values are in addition to any parent in the display list.
// So if this Group was in a Group that has x: 200, then this will be added to the cameraOffset.x
func (self *PhysicsP2BodyDebug) SetFixedToCameraA(member bool) {
    self.Object.Set("fixedToCamera", member)
}

// If this object is {@link Phaser.Group#fixedToCamera fixedToCamera} then this stores the x/y position offset relative to the top-left of the camera view.
// If the parent of this Group is also `fixedToCamera` then the offset here is in addition to that and should typically be disabled.
func (self *PhysicsP2BodyDebug) GetCameraOffsetA() *Point{
    return &Point{self.Object.Get("cameraOffset")}
}

// If this object is {@link Phaser.Group#fixedToCamera fixedToCamera} then this stores the x/y position offset relative to the top-left of the camera view.
// If the parent of this Group is also `fixedToCamera` then the offset here is in addition to that and should typically be disabled.
func (self *PhysicsP2BodyDebug) SetCameraOffsetA(member *Point) {
    self.Object.Set("cameraOffset", member)
}

// The hash array is an array belonging to this Group into which you can add any of its children via Group.addToHash and Group.removeFromHash.
// 
// Only children of this Group can be added to and removed from the hash.
// 
// This hash is used automatically by Phaser Arcade Physics in order to perform non z-index based destructive sorting.
// However if you don't use Arcade Physics, or this isn't a physics enabled Group, then you can use the hash to perform your own
// sorting and filtering of Group children without touching their z-index (and therefore display draw order)
func (self *PhysicsP2BodyDebug) GetHashA() []interface{}{
	array00 := self.Object.Get("hash")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// The hash array is an array belonging to this Group into which you can add any of its children via Group.addToHash and Group.removeFromHash.
// 
// Only children of this Group can be added to and removed from the hash.
// 
// This hash is used automatically by Phaser Arcade Physics in order to perform non z-index based destructive sorting.
// However if you don't use Arcade Physics, or this isn't a physics enabled Group, then you can use the hash to perform your own
// sorting and filtering of Group children without touching their z-index (and therefore display draw order)
func (self *PhysicsP2BodyDebug) SetHashA(member []interface{}) {
    self.Object.Set("hash", member)
}

// Total number of existing children in the group.
func (self *PhysicsP2BodyDebug) GetTotalA() int{
    return self.Object.Get("total").Int()
}

// Total number of existing children in the group.
func (self *PhysicsP2BodyDebug) SetTotalA(member int) {
    self.Object.Set("total", member)
}

// Total number of children in this group, regardless of exists/alive status.
func (self *PhysicsP2BodyDebug) GetLengthA() int{
    return self.Object.Get("length").Int()
}

// Total number of children in this group, regardless of exists/alive status.
func (self *PhysicsP2BodyDebug) SetLengthA(member int) {
    self.Object.Set("length", member)
}

// The angle of rotation of the group container, in degrees.
// 
// This adjusts the group itself by modifying its local rotation transform.
// 
// This has no impact on the rotation/angle properties of the children, but it will update their worldTransform
// and on-screen orientation and position.
func (self *PhysicsP2BodyDebug) GetAngleA() int{
    return self.Object.Get("angle").Int()
}

// The angle of rotation of the group container, in degrees.
// 
// This adjusts the group itself by modifying its local rotation transform.
// 
// This has no impact on the rotation/angle properties of the children, but it will update their worldTransform
// and on-screen orientation and position.
func (self *PhysicsP2BodyDebug) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// The center x coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) GetCenterXA() int{
    return self.Object.Get("centerX").Int()
}

// The center x coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// The center y coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) GetCenterYA() int{
    return self.Object.Get("centerY").Int()
}

// The center y coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// The left coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) GetLeftA() int{
    return self.Object.Get("left").Int()
}

// The left coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// The right coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) GetRightA() int{
    return self.Object.Get("right").Int()
}

// The right coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) SetRightA(member int) {
    self.Object.Set("right", member)
}

// The top coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) GetTopA() int{
    return self.Object.Get("top").Int()
}

// The top coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) SetTopA(member int) {
    self.Object.Set("top", member)
}

// The bottom coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) GetBottomA() int{
    return self.Object.Get("bottom").Int()
}

// The bottom coordinate of this Group.
// 
// It is derived by calling `getBounds`, calculating the Groups dimensions based on its
// visible children.
// 
// Note that no ancestors are factored into the result, meaning that if this Group is 
// nested within another Group, with heavy transforms on it, the result of this property 
// is likely to be incorrect. It is safe to get and set this property if the Group is a
// top-level descendant of Phaser.World, or untransformed parents.
func (self *PhysicsP2BodyDebug) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// The x coordinate of the group container.
// 
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *PhysicsP2BodyDebug) GetXA() int{
    return self.Object.Get("x").Int()
}

// The x coordinate of the group container.
// 
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *PhysicsP2BodyDebug) SetXA(member int) {
    self.Object.Set("x", member)
}

// The y coordinate of the group container.
// 
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *PhysicsP2BodyDebug) GetYA() int{
    return self.Object.Get("y").Int()
}

// The y coordinate of the group container.
// 
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *PhysicsP2BodyDebug) SetYA(member int) {
    self.Object.Set("y", member)
}

// The angle of rotation of the group container, in radians.
// 
// This will adjust the group container itself by modifying its rotation.
// This will have no impact on the rotation value of its children, but it will update their worldTransform and on-screen position.
func (self *PhysicsP2BodyDebug) GetRotationA() int{
    return self.Object.Get("rotation").Int()
}

// The angle of rotation of the group container, in radians.
// 
// This will adjust the group container itself by modifying its rotation.
// This will have no impact on the rotation value of its children, but it will update their worldTransform and on-screen position.
func (self *PhysicsP2BodyDebug) SetRotationA(member int) {
    self.Object.Set("rotation", member)
}

// The visible state of the group. Non-visible Groups and all of their children are not rendered.
func (self *PhysicsP2BodyDebug) GetVisibleA() bool{
    return self.Object.Get("visible").Bool()
}

// The visible state of the group. Non-visible Groups and all of their children are not rendered.
func (self *PhysicsP2BodyDebug) SetVisibleA(member bool) {
    self.Object.Set("visible", member)
}

// The alpha value of the group container.
func (self *PhysicsP2BodyDebug) GetAlphaA() int{
    return self.Object.Get("alpha").Int()
}

// The alpha value of the group container.
func (self *PhysicsP2BodyDebug) SetAlphaA(member int) {
    self.Object.Set("alpha", member)
}

// [read-only] The array of children of this container.
func (self *PhysicsP2BodyDebug) GetChildrenA() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// [read-only] The array of children of this container.
func (self *PhysicsP2BodyDebug) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *PhysicsP2BodyDebug) GetIgnoreChildInputA() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *PhysicsP2BodyDebug) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *PhysicsP2BodyDebug) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *PhysicsP2BodyDebug) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *PhysicsP2BodyDebug) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *PhysicsP2BodyDebug) SetHeightA(member int) {
    self.Object.Set("height", member)
}



// Core update.
func (self *PhysicsP2BodyDebug) UpdateSpriteTransform() {
    self.Object.Call("updateSpriteTransform")
}

// Core update.
func (self *PhysicsP2BodyDebug) UpdateSpriteTransformI(args ...interface{}) {
    self.Object.Call("updateSpriteTransform", args)
}

// Draws the P2 shapes to the Graphics object.
func (self *PhysicsP2BodyDebug) Draw() {
    self.Object.Call("draw")
}

// Draws the P2 shapes to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawI(args ...interface{}) {
    self.Object.Call("draw", args)
}

// Draws a p2.Box to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawRectangle() {
    self.Object.Call("drawRectangle")
}

// Draws a p2.Box to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawRectangleI(args ...interface{}) {
    self.Object.Call("drawRectangle", args)
}

// Draws a p2.Circle to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawCircle() {
    self.Object.Call("drawCircle")
}

// Draws a p2.Circle to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawCircleI(args ...interface{}) {
    self.Object.Call("drawCircle", args)
}

// Draws a p2.Line to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawLine() {
    self.Object.Call("drawLine")
}

// Draws a p2.Line to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawLineI(args ...interface{}) {
    self.Object.Call("drawLine", args)
}

// Draws a p2.Convex to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawConvex() {
    self.Object.Call("drawConvex")
}

// Draws a p2.Convex to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawConvexI(args ...interface{}) {
    self.Object.Call("drawConvex", args)
}

// Draws a p2.Path to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawPath() {
    self.Object.Call("drawPath")
}

// Draws a p2.Path to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawPathI(args ...interface{}) {
    self.Object.Call("drawPath", args)
}

// Draws a p2.Plane to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawPlane() {
    self.Object.Call("drawPlane")
}

// Draws a p2.Plane to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawPlaneI(args ...interface{}) {
    self.Object.Call("drawPlane", args)
}

// Draws a p2.Capsule to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawCapsule() {
    self.Object.Call("drawCapsule")
}

// Draws a p2.Capsule to the Graphics object.
func (self *PhysicsP2BodyDebug) DrawCapsuleI(args ...interface{}) {
    self.Object.Call("drawCapsule", args)
}

// Picks a random pastel color.
func (self *PhysicsP2BodyDebug) RandomPastelHex() {
    self.Object.Call("randomPastelHex")
}

// Picks a random pastel color.
func (self *PhysicsP2BodyDebug) RandomPastelHexI(args ...interface{}) {
    self.Object.Call("randomPastelHex", args)
}

// Converts from RGB to Hex.
func (self *PhysicsP2BodyDebug) RgbToHex() {
    self.Object.Call("rgbToHex")
}

// Converts from RGB to Hex.
func (self *PhysicsP2BodyDebug) RgbToHexI(args ...interface{}) {
    self.Object.Call("rgbToHex", args)
}

// Component to hex conversion.
func (self *PhysicsP2BodyDebug) ComponentToHex() {
    self.Object.Call("componentToHex")
}

// Component to hex conversion.
func (self *PhysicsP2BodyDebug) ComponentToHexI(args ...interface{}) {
    self.Object.Call("componentToHex", args)
}

// Adds an existing object as the top child in this group.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If the child was already in this Group, it is simply returned, and nothing else happens to it.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
// 
// Use {@link Phaser.Group#addAt addAt} to control where a child is added. Use {@link Phaser.Group#create create} to create and add a new child.
func (self *PhysicsP2BodyDebug) Add(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("add", child)}
}

// Adds an existing object as the top child in this group.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If the child was already in this Group, it is simply returned, and nothing else happens to it.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
// 
// Use {@link Phaser.Group#addAt addAt} to control where a child is added. Use {@link Phaser.Group#create create} to create and add a new child.
func (self *PhysicsP2BodyDebug) Add1O(child *DisplayObject, silent bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("add", child, silent)}
}

// Adds an existing object as the top child in this group.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If the child was already in this Group, it is simply returned, and nothing else happens to it.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
// 
// Use {@link Phaser.Group#addAt addAt} to control where a child is added. Use {@link Phaser.Group#create create} to create and add a new child.
func (self *PhysicsP2BodyDebug) Add2O(child *DisplayObject, silent bool, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("add", child, silent, index)}
}

// Adds an existing object as the top child in this group.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If the child was already in this Group, it is simply returned, and nothing else happens to it.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
// 
// Use {@link Phaser.Group#addAt addAt} to control where a child is added. Use {@link Phaser.Group#create create} to create and add a new child.
func (self *PhysicsP2BodyDebug) AddI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("add", args)}
}

// Adds an existing object to this group.
// 
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) AddAt(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addAt", child)}
}

// Adds an existing object to this group.
// 
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) AddAt1O(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addAt", child, index)}
}

// Adds an existing object to this group.
// 
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) AddAt2O(child *DisplayObject, index int, silent bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("addAt", child, index, silent)}
}

// Adds an existing object to this group.
// 
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) AddAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addAt", args)}
}

// Adds a child of this Group into the hash array.
// This call will return false if the child is not a child of this Group, or is already in the hash.
func (self *PhysicsP2BodyDebug) AddToHash(child *DisplayObject) bool{
    return self.Object.Call("addToHash", child).Bool()
}

// Adds a child of this Group into the hash array.
// This call will return false if the child is not a child of this Group, or is already in the hash.
func (self *PhysicsP2BodyDebug) AddToHashI(args ...interface{}) bool{
    return self.Object.Call("addToHash", args).Bool()
}

// Removes a child of this Group from the hash array.
// This call will return false if the child is not in the hash.
func (self *PhysicsP2BodyDebug) RemoveFromHash(child *DisplayObject) bool{
    return self.Object.Call("removeFromHash", child).Bool()
}

// Removes a child of this Group from the hash array.
// This call will return false if the child is not in the hash.
func (self *PhysicsP2BodyDebug) RemoveFromHashI(args ...interface{}) bool{
    return self.Object.Call("removeFromHash", args).Bool()
}

// Adds an array of existing Display Objects to this Group.
// 
// The Display Objects are automatically added to the top of this Group, and will render on-top of everything already in this Group.
// 
// As well as an array you can also pass another Group as the first argument. In this case all of the children from that
// Group will be removed from it and added into this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) AddMultiple(children interface{}) interface{}{
    return self.Object.Call("addMultiple", children)
}

// Adds an array of existing Display Objects to this Group.
// 
// The Display Objects are automatically added to the top of this Group, and will render on-top of everything already in this Group.
// 
// As well as an array you can also pass another Group as the first argument. In this case all of the children from that
// Group will be removed from it and added into this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) AddMultiple1O(children interface{}, silent bool) interface{}{
    return self.Object.Call("addMultiple", children, silent)
}

// Adds an array of existing Display Objects to this Group.
// 
// The Display Objects are automatically added to the top of this Group, and will render on-top of everything already in this Group.
// 
// As well as an array you can also pass another Group as the first argument. In this case all of the children from that
// Group will be removed from it and added into this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) AddMultipleI(args ...interface{}) interface{}{
    return self.Object.Call("addMultiple", args)
}

// Returns the child found at the given index within this group.
func (self *PhysicsP2BodyDebug) GetAt(index int) interface{}{
    return self.Object.Call("getAt", index)
}

// Returns the child found at the given index within this group.
func (self *PhysicsP2BodyDebug) GetAtI(args ...interface{}) interface{}{
    return self.Object.Call("getAt", args)
}

// Creates a new Phaser.Sprite object and adds it to the top of this group.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) Create(x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", x, y)}
}

// Creates a new Phaser.Sprite object and adds it to the top of this group.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) Create1O(x int, y int, key interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", x, y, key)}
}

// Creates a new Phaser.Sprite object and adds it to the top of this group.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) Create2O(x int, y int, key interface{}, frame interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", x, y, key, frame)}
}

// Creates a new Phaser.Sprite object and adds it to the top of this group.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) Create3O(x int, y int, key interface{}, frame interface{}, exists bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", x, y, key, frame, exists)}
}

// Creates a new Phaser.Sprite object and adds it to the top of this group.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) Create4O(x int, y int, key interface{}, frame interface{}, exists bool, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", x, y, key, frame, exists, index)}
}

// Creates a new Phaser.Sprite object and adds it to the top of this group.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// The child is automatically added to the top of the group, and is displayed above every previous child.
// 
// Or if the _optional_ index is specified, the child is added at the location specified by the index value, 
// this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) CreateI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("create", args)}
}

// Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
// 
// This method is useful if you need to quickly generate a pool of sprites, such as bullets.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// You can provide an array as the `key` and / or `frame` arguments. When you do this
// it will create `quantity` Sprites for every key (and frame) in the arrays.
// 
// For example:
// 
// `createMultiple(25, ['ball', 'carrot'])`
// 
// In the above code there are 2 keys (ball and carrot) which means that 50 sprites will be
// created in total, 25 of each. You can also have the `frame` as an array:
// 
// `createMultiple(5, 'bricks', [0, 1, 2, 3])`
// 
// In the above there is one key (bricks), which is a sprite sheet. The frames array tells
// this method to use frames 0, 1, 2 and 3. So in total it will create 20 sprites, because
// the quantity was set to 5, so that is 5 brick sprites of frame 0, 5 brick sprites with
// frame 1, and so on.
// 
// If you set both the key and frame arguments to be arrays then understand it will create
// a total quantity of sprites equal to the size of both arrays times each other. I.e.:
// 
// `createMultiple(20, ['diamonds', 'balls'], [0, 1, 2])`
// 
// The above will create 20 'diamonds' of frame 0, 20 with frame 1 and 20 with frame 2.
// It will then create 20 'balls' of frame 0, 20 with frame 1 and 20 with frame 2.
// In total it will have created 120 sprites.
// 
// By default the Sprites will have their `exists` property set to `false`, and they will be 
// positioned at 0x0, relative to the `Group.x / y` values.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) CreateMultiple(quantity int, key interface{}) []interface{}{
	array00 := self.Object.Call("createMultiple", quantity, key)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
// 
// This method is useful if you need to quickly generate a pool of sprites, such as bullets.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// You can provide an array as the `key` and / or `frame` arguments. When you do this
// it will create `quantity` Sprites for every key (and frame) in the arrays.
// 
// For example:
// 
// `createMultiple(25, ['ball', 'carrot'])`
// 
// In the above code there are 2 keys (ball and carrot) which means that 50 sprites will be
// created in total, 25 of each. You can also have the `frame` as an array:
// 
// `createMultiple(5, 'bricks', [0, 1, 2, 3])`
// 
// In the above there is one key (bricks), which is a sprite sheet. The frames array tells
// this method to use frames 0, 1, 2 and 3. So in total it will create 20 sprites, because
// the quantity was set to 5, so that is 5 brick sprites of frame 0, 5 brick sprites with
// frame 1, and so on.
// 
// If you set both the key and frame arguments to be arrays then understand it will create
// a total quantity of sprites equal to the size of both arrays times each other. I.e.:
// 
// `createMultiple(20, ['diamonds', 'balls'], [0, 1, 2])`
// 
// The above will create 20 'diamonds' of frame 0, 20 with frame 1 and 20 with frame 2.
// It will then create 20 'balls' of frame 0, 20 with frame 1 and 20 with frame 2.
// In total it will have created 120 sprites.
// 
// By default the Sprites will have their `exists` property set to `false`, and they will be 
// positioned at 0x0, relative to the `Group.x / y` values.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) CreateMultiple1O(quantity int, key interface{}, frame interface{}) []interface{}{
	array00 := self.Object.Call("createMultiple", quantity, key, frame)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
// 
// This method is useful if you need to quickly generate a pool of sprites, such as bullets.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// You can provide an array as the `key` and / or `frame` arguments. When you do this
// it will create `quantity` Sprites for every key (and frame) in the arrays.
// 
// For example:
// 
// `createMultiple(25, ['ball', 'carrot'])`
// 
// In the above code there are 2 keys (ball and carrot) which means that 50 sprites will be
// created in total, 25 of each. You can also have the `frame` as an array:
// 
// `createMultiple(5, 'bricks', [0, 1, 2, 3])`
// 
// In the above there is one key (bricks), which is a sprite sheet. The frames array tells
// this method to use frames 0, 1, 2 and 3. So in total it will create 20 sprites, because
// the quantity was set to 5, so that is 5 brick sprites of frame 0, 5 brick sprites with
// frame 1, and so on.
// 
// If you set both the key and frame arguments to be arrays then understand it will create
// a total quantity of sprites equal to the size of both arrays times each other. I.e.:
// 
// `createMultiple(20, ['diamonds', 'balls'], [0, 1, 2])`
// 
// The above will create 20 'diamonds' of frame 0, 20 with frame 1 and 20 with frame 2.
// It will then create 20 'balls' of frame 0, 20 with frame 1 and 20 with frame 2.
// In total it will have created 120 sprites.
// 
// By default the Sprites will have their `exists` property set to `false`, and they will be 
// positioned at 0x0, relative to the `Group.x / y` values.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) CreateMultiple2O(quantity int, key interface{}, frame interface{}, exists bool) []interface{}{
	array00 := self.Object.Call("createMultiple", quantity, key, frame, exists)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Creates multiple Phaser.Sprite objects and adds them to the top of this Group.
// 
// This method is useful if you need to quickly generate a pool of sprites, such as bullets.
// 
// Use {@link Phaser.Group#classType classType} to change the type of object created.
// 
// You can provide an array as the `key` and / or `frame` arguments. When you do this
// it will create `quantity` Sprites for every key (and frame) in the arrays.
// 
// For example:
// 
// `createMultiple(25, ['ball', 'carrot'])`
// 
// In the above code there are 2 keys (ball and carrot) which means that 50 sprites will be
// created in total, 25 of each. You can also have the `frame` as an array:
// 
// `createMultiple(5, 'bricks', [0, 1, 2, 3])`
// 
// In the above there is one key (bricks), which is a sprite sheet. The frames array tells
// this method to use frames 0, 1, 2 and 3. So in total it will create 20 sprites, because
// the quantity was set to 5, so that is 5 brick sprites of frame 0, 5 brick sprites with
// frame 1, and so on.
// 
// If you set both the key and frame arguments to be arrays then understand it will create
// a total quantity of sprites equal to the size of both arrays times each other. I.e.:
// 
// `createMultiple(20, ['diamonds', 'balls'], [0, 1, 2])`
// 
// The above will create 20 'diamonds' of frame 0, 20 with frame 1 and 20 with frame 2.
// It will then create 20 'balls' of frame 0, 20 with frame 1 and 20 with frame 2.
// In total it will have created 120 sprites.
// 
// By default the Sprites will have their `exists` property set to `false`, and they will be 
// positioned at 0x0, relative to the `Group.x / y` values.
// 
// If `Group.enableBody` is set, then a physics body will be created on the objects, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the objects, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) CreateMultipleI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("createMultiple", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Internal method that re-applies all of the children's Z values.
// 
// This must be called whenever children ordering is altered so that their `z` indices are correctly updated.
func (self *PhysicsP2BodyDebug) UpdateZ() {
    self.Object.Call("updateZ")
}

// Internal method that re-applies all of the children's Z values.
// 
// This must be called whenever children ordering is altered so that their `z` indices are correctly updated.
func (self *PhysicsP2BodyDebug) UpdateZI(args ...interface{}) {
    self.Object.Call("updateZ", args)
}

// This method iterates through all children in the Group (regardless if they are visible or exist)
// and then changes their position so they are arranged in a Grid formation. Children must have
// the `alignTo` method in order to be positioned by this call. All default Phaser Game Objects have
// this.
// 
// The grid dimensions are determined by the first four arguments. The `rows` and `columns` arguments
// relate to the width and height of the grid respectively.
// 
// For example if the Group had 100 children in it:
// 
// `Group.align(10, 10, 32, 32)`
// 
// This will align all of the children into a grid formation of 10x10, using 32 pixels per
// grid cell. If you want a wider grid, you could do:
// 
// `Group.align(25, 4, 32, 32)`
// 
// This will align the children into a grid of 25x4, again using 32 pixels per grid cell.
// 
// You can choose to set _either_ the `rows` or `columns` value to -1. Doing so tells the method
// to keep on aligning children until there are no children left. For example if this Group had
// 48 children in it, the following:
// 
// `Group.align(-1, 8, 32, 32)`
// 
// ... will align the children so that there are 8 columns vertically (the second argument), 
// and each row will contain 6 sprites, except the last one, which will contain 5 (totaling 48)
// 
// You can also do:
// 
// `Group.align(10, -1, 32, 32)`
// 
// In this case it will create a grid 10 wide, and as tall as it needs to be in order to fit
// all of the children in.
// 
// The `position` property allows you to control where in each grid cell the child is positioned.
// This is a constant and can be one of `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, 
// `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, 
// `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` or `Phaser.BOTTOM_RIGHT`.
// 
// The final argument; `offset` lets you start the alignment from a specific child index.
func (self *PhysicsP2BodyDebug) Align(rows int, columns int, cellWidth int, cellHeight int) {
    self.Object.Call("align", rows, columns, cellWidth, cellHeight)
}

// This method iterates through all children in the Group (regardless if they are visible or exist)
// and then changes their position so they are arranged in a Grid formation. Children must have
// the `alignTo` method in order to be positioned by this call. All default Phaser Game Objects have
// this.
// 
// The grid dimensions are determined by the first four arguments. The `rows` and `columns` arguments
// relate to the width and height of the grid respectively.
// 
// For example if the Group had 100 children in it:
// 
// `Group.align(10, 10, 32, 32)`
// 
// This will align all of the children into a grid formation of 10x10, using 32 pixels per
// grid cell. If you want a wider grid, you could do:
// 
// `Group.align(25, 4, 32, 32)`
// 
// This will align the children into a grid of 25x4, again using 32 pixels per grid cell.
// 
// You can choose to set _either_ the `rows` or `columns` value to -1. Doing so tells the method
// to keep on aligning children until there are no children left. For example if this Group had
// 48 children in it, the following:
// 
// `Group.align(-1, 8, 32, 32)`
// 
// ... will align the children so that there are 8 columns vertically (the second argument), 
// and each row will contain 6 sprites, except the last one, which will contain 5 (totaling 48)
// 
// You can also do:
// 
// `Group.align(10, -1, 32, 32)`
// 
// In this case it will create a grid 10 wide, and as tall as it needs to be in order to fit
// all of the children in.
// 
// The `position` property allows you to control where in each grid cell the child is positioned.
// This is a constant and can be one of `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, 
// `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, 
// `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` or `Phaser.BOTTOM_RIGHT`.
// 
// The final argument; `offset` lets you start the alignment from a specific child index.
func (self *PhysicsP2BodyDebug) Align1O(rows int, columns int, cellWidth int, cellHeight int, position int) {
    self.Object.Call("align", rows, columns, cellWidth, cellHeight, position)
}

// This method iterates through all children in the Group (regardless if they are visible or exist)
// and then changes their position so they are arranged in a Grid formation. Children must have
// the `alignTo` method in order to be positioned by this call. All default Phaser Game Objects have
// this.
// 
// The grid dimensions are determined by the first four arguments. The `rows` and `columns` arguments
// relate to the width and height of the grid respectively.
// 
// For example if the Group had 100 children in it:
// 
// `Group.align(10, 10, 32, 32)`
// 
// This will align all of the children into a grid formation of 10x10, using 32 pixels per
// grid cell. If you want a wider grid, you could do:
// 
// `Group.align(25, 4, 32, 32)`
// 
// This will align the children into a grid of 25x4, again using 32 pixels per grid cell.
// 
// You can choose to set _either_ the `rows` or `columns` value to -1. Doing so tells the method
// to keep on aligning children until there are no children left. For example if this Group had
// 48 children in it, the following:
// 
// `Group.align(-1, 8, 32, 32)`
// 
// ... will align the children so that there are 8 columns vertically (the second argument), 
// and each row will contain 6 sprites, except the last one, which will contain 5 (totaling 48)
// 
// You can also do:
// 
// `Group.align(10, -1, 32, 32)`
// 
// In this case it will create a grid 10 wide, and as tall as it needs to be in order to fit
// all of the children in.
// 
// The `position` property allows you to control where in each grid cell the child is positioned.
// This is a constant and can be one of `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, 
// `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, 
// `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` or `Phaser.BOTTOM_RIGHT`.
// 
// The final argument; `offset` lets you start the alignment from a specific child index.
func (self *PhysicsP2BodyDebug) Align2O(rows int, columns int, cellWidth int, cellHeight int, position int, offset int) {
    self.Object.Call("align", rows, columns, cellWidth, cellHeight, position, offset)
}

// This method iterates through all children in the Group (regardless if they are visible or exist)
// and then changes their position so they are arranged in a Grid formation. Children must have
// the `alignTo` method in order to be positioned by this call. All default Phaser Game Objects have
// this.
// 
// The grid dimensions are determined by the first four arguments. The `rows` and `columns` arguments
// relate to the width and height of the grid respectively.
// 
// For example if the Group had 100 children in it:
// 
// `Group.align(10, 10, 32, 32)`
// 
// This will align all of the children into a grid formation of 10x10, using 32 pixels per
// grid cell. If you want a wider grid, you could do:
// 
// `Group.align(25, 4, 32, 32)`
// 
// This will align the children into a grid of 25x4, again using 32 pixels per grid cell.
// 
// You can choose to set _either_ the `rows` or `columns` value to -1. Doing so tells the method
// to keep on aligning children until there are no children left. For example if this Group had
// 48 children in it, the following:
// 
// `Group.align(-1, 8, 32, 32)`
// 
// ... will align the children so that there are 8 columns vertically (the second argument), 
// and each row will contain 6 sprites, except the last one, which will contain 5 (totaling 48)
// 
// You can also do:
// 
// `Group.align(10, -1, 32, 32)`
// 
// In this case it will create a grid 10 wide, and as tall as it needs to be in order to fit
// all of the children in.
// 
// The `position` property allows you to control where in each grid cell the child is positioned.
// This is a constant and can be one of `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, 
// `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, 
// `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` or `Phaser.BOTTOM_RIGHT`.
// 
// The final argument; `offset` lets you start the alignment from a specific child index.
func (self *PhysicsP2BodyDebug) AlignI(args ...interface{}) {
    self.Object.Call("align", args)
}

// Sets the group cursor to the first child in the group.
// 
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *PhysicsP2BodyDebug) ResetCursor() interface{}{
    return self.Object.Call("resetCursor")
}

// Sets the group cursor to the first child in the group.
// 
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *PhysicsP2BodyDebug) ResetCursor1O(index int) interface{}{
    return self.Object.Call("resetCursor", index)
}

// Sets the group cursor to the first child in the group.
// 
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *PhysicsP2BodyDebug) ResetCursorI(args ...interface{}) interface{}{
    return self.Object.Call("resetCursor", args)
}

// Advances the group cursor to the next (higher) object in the group.
// 
// If the cursor is at the end of the group (top child) it is moved the start of the group (bottom child).
func (self *PhysicsP2BodyDebug) Next() interface{}{
    return self.Object.Call("next")
}

// Advances the group cursor to the next (higher) object in the group.
// 
// If the cursor is at the end of the group (top child) it is moved the start of the group (bottom child).
func (self *PhysicsP2BodyDebug) NextI(args ...interface{}) interface{}{
    return self.Object.Call("next", args)
}

// Moves the group cursor to the previous (lower) child in the group.
// 
// If the cursor is at the start of the group (bottom child) it is moved to the end (top child).
func (self *PhysicsP2BodyDebug) Previous() interface{}{
    return self.Object.Call("previous")
}

// Moves the group cursor to the previous (lower) child in the group.
// 
// If the cursor is at the start of the group (bottom child) it is moved to the end (top child).
func (self *PhysicsP2BodyDebug) PreviousI(args ...interface{}) interface{}{
    return self.Object.Call("previous", args)
}

// Swaps the position of two children in this group.
// 
// Both children must be in this group, a child cannot be swapped with itself, and unparented children cannot be swapped.
func (self *PhysicsP2BodyDebug) Swap(child1 interface{}, child2 interface{}) {
    self.Object.Call("swap", child1, child2)
}

// Swaps the position of two children in this group.
// 
// Both children must be in this group, a child cannot be swapped with itself, and unparented children cannot be swapped.
func (self *PhysicsP2BodyDebug) SwapI(args ...interface{}) {
    self.Object.Call("swap", args)
}

// Brings the given child to the top of this group so it renders above all other children.
func (self *PhysicsP2BodyDebug) BringToTop(child interface{}) interface{}{
    return self.Object.Call("bringToTop", child)
}

// Brings the given child to the top of this group so it renders above all other children.
func (self *PhysicsP2BodyDebug) BringToTopI(args ...interface{}) interface{}{
    return self.Object.Call("bringToTop", args)
}

// Sends the given child to the bottom of this group so it renders below all other children.
func (self *PhysicsP2BodyDebug) SendToBack(child interface{}) interface{}{
    return self.Object.Call("sendToBack", child)
}

// Sends the given child to the bottom of this group so it renders below all other children.
func (self *PhysicsP2BodyDebug) SendToBackI(args ...interface{}) interface{}{
    return self.Object.Call("sendToBack", args)
}

// Moves the given child up one place in this group unless it's already at the top.
func (self *PhysicsP2BodyDebug) MoveUp(child interface{}) interface{}{
    return self.Object.Call("moveUp", child)
}

// Moves the given child up one place in this group unless it's already at the top.
func (self *PhysicsP2BodyDebug) MoveUpI(args ...interface{}) interface{}{
    return self.Object.Call("moveUp", args)
}

// Moves the given child down one place in this group unless it's already at the bottom.
func (self *PhysicsP2BodyDebug) MoveDown(child interface{}) interface{}{
    return self.Object.Call("moveDown", child)
}

// Moves the given child down one place in this group unless it's already at the bottom.
func (self *PhysicsP2BodyDebug) MoveDownI(args ...interface{}) interface{}{
    return self.Object.Call("moveDown", args)
}

// Positions the child found at the given index within this group to the given x and y coordinates.
func (self *PhysicsP2BodyDebug) Xy(index int, x int, y int) {
    self.Object.Call("xy", index, x, y)
}

// Positions the child found at the given index within this group to the given x and y coordinates.
func (self *PhysicsP2BodyDebug) XyI(args ...interface{}) {
    self.Object.Call("xy", args)
}

// Reverses all children in this group.
// 
// This operation applies only to immediate children and does not propagate to subgroups.
func (self *PhysicsP2BodyDebug) Reverse() {
    self.Object.Call("reverse")
}

// Reverses all children in this group.
// 
// This operation applies only to immediate children and does not propagate to subgroups.
func (self *PhysicsP2BodyDebug) ReverseI(args ...interface{}) {
    self.Object.Call("reverse", args)
}

// Get the index position of the given child in this group, which should match the child's `z` property.
func (self *PhysicsP2BodyDebug) GetIndex(child interface{}) int{
    return self.Object.Call("getIndex", child).Int()
}

// Get the index position of the given child in this group, which should match the child's `z` property.
func (self *PhysicsP2BodyDebug) GetIndexI(args ...interface{}) int{
    return self.Object.Call("getIndex", args).Int()
}

// Searches the Group for the first instance of a child with the `name`
// property matching the given argument. Should more than one child have
// the same name only the first instance is returned.
func (self *PhysicsP2BodyDebug) GetByName(name string) interface{}{
    return self.Object.Call("getByName", name)
}

// Searches the Group for the first instance of a child with the `name`
// property matching the given argument. Should more than one child have
// the same name only the first instance is returned.
func (self *PhysicsP2BodyDebug) GetByNameI(args ...interface{}) interface{}{
    return self.Object.Call("getByName", args)
}

// Replaces a child of this Group with the given newChild. The newChild cannot be a member of this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) Replace(oldChild interface{}, newChild interface{}) interface{}{
    return self.Object.Call("replace", oldChild, newChild)
}

// Replaces a child of this Group with the given newChild. The newChild cannot be a member of this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *PhysicsP2BodyDebug) ReplaceI(args ...interface{}) interface{}{
    return self.Object.Call("replace", args)
}

// Checks if the child has the given property.
// 
// Will scan up to 4 levels deep only.
func (self *PhysicsP2BodyDebug) HasProperty(child interface{}, key []string) bool{
    return self.Object.Call("hasProperty", child, key).Bool()
}

// Checks if the child has the given property.
// 
// Will scan up to 4 levels deep only.
func (self *PhysicsP2BodyDebug) HasPropertyI(args ...interface{}) bool{
    return self.Object.Call("hasProperty", args).Bool()
}

// Sets a property to the given value on the child. The operation parameter controls how the value is set.
// 
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *PhysicsP2BodyDebug) SetProperty(child interface{}, key []interface{}, value interface{}) bool{
    return self.Object.Call("setProperty", child, key, value).Bool()
}

// Sets a property to the given value on the child. The operation parameter controls how the value is set.
// 
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *PhysicsP2BodyDebug) SetProperty1O(child interface{}, key []interface{}, value interface{}, operation int) bool{
    return self.Object.Call("setProperty", child, key, value, operation).Bool()
}

// Sets a property to the given value on the child. The operation parameter controls how the value is set.
// 
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *PhysicsP2BodyDebug) SetProperty2O(child interface{}, key []interface{}, value interface{}, operation int, force bool) bool{
    return self.Object.Call("setProperty", child, key, value, operation, force).Bool()
}

// Sets a property to the given value on the child. The operation parameter controls how the value is set.
// 
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *PhysicsP2BodyDebug) SetPropertyI(args ...interface{}) bool{
    return self.Object.Call("setProperty", args).Bool()
}

// Checks a property for the given value on the child.
func (self *PhysicsP2BodyDebug) CheckProperty(child interface{}, key []interface{}, value interface{}) bool{
    return self.Object.Call("checkProperty", child, key, value).Bool()
}

// Checks a property for the given value on the child.
func (self *PhysicsP2BodyDebug) CheckProperty1O(child interface{}, key []interface{}, value interface{}, force bool) bool{
    return self.Object.Call("checkProperty", child, key, value, force).Bool()
}

// Checks a property for the given value on the child.
func (self *PhysicsP2BodyDebug) CheckPropertyI(args ...interface{}) bool{
    return self.Object.Call("checkProperty", args).Bool()
}

// Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) Set(child *Sprite, key string, value interface{}) bool{
    return self.Object.Call("set", child, key, value).Bool()
}

// Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) Set1O(child *Sprite, key string, value interface{}, checkAlive bool) bool{
    return self.Object.Call("set", child, key, value, checkAlive).Bool()
}

// Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) Set2O(child *Sprite, key string, value interface{}, checkAlive bool, checkVisible bool) bool{
    return self.Object.Call("set", child, key, value, checkAlive, checkVisible).Bool()
}

// Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) Set3O(child *Sprite, key string, value interface{}, checkAlive bool, checkVisible bool, operation int) bool{
    return self.Object.Call("set", child, key, value, checkAlive, checkVisible, operation).Bool()
}

// Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) Set4O(child *Sprite, key string, value interface{}, checkAlive bool, checkVisible bool, operation int, force bool) bool{
    return self.Object.Call("set", child, key, value, checkAlive, checkVisible, operation, force).Bool()
}

// Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetI(args ...interface{}) bool{
    return self.Object.Call("set", args).Bool()
}

// Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAll(key string, value interface{}) {
    self.Object.Call("setAll", key, value)
}

// Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAll1O(key string, value interface{}, checkAlive bool) {
    self.Object.Call("setAll", key, value, checkAlive)
}

// Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAll2O(key string, value interface{}, checkAlive bool, checkVisible bool) {
    self.Object.Call("setAll", key, value, checkAlive, checkVisible)
}

// Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAll3O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int) {
    self.Object.Call("setAll", key, value, checkAlive, checkVisible, operation)
}

// Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAll4O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int, force bool) {
    self.Object.Call("setAll", key, value, checkAlive, checkVisible, operation, force)
}

// Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAllI(args ...interface{}) {
    self.Object.Call("setAll", args)
}

// Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAllChildren(key string, value interface{}) {
    self.Object.Call("setAllChildren", key, value)
}

// Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAllChildren1O(key string, value interface{}, checkAlive bool) {
    self.Object.Call("setAllChildren", key, value, checkAlive)
}

// Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAllChildren2O(key string, value interface{}, checkAlive bool, checkVisible bool) {
    self.Object.Call("setAllChildren", key, value, checkAlive, checkVisible)
}

// Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAllChildren3O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int) {
    self.Object.Call("setAllChildren", key, value, checkAlive, checkVisible, operation)
}

// Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAllChildren4O(key string, value interface{}, checkAlive bool, checkVisible bool, operation int, force bool) {
    self.Object.Call("setAllChildren", key, value, checkAlive, checkVisible, operation, force)
}

// Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *PhysicsP2BodyDebug) SetAllChildrenI(args ...interface{}) {
    self.Object.Call("setAllChildren", args)
}

// Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *PhysicsP2BodyDebug) CheckAll(key string, value interface{}) {
    self.Object.Call("checkAll", key, value)
}

// Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *PhysicsP2BodyDebug) CheckAll1O(key string, value interface{}, checkAlive bool) {
    self.Object.Call("checkAll", key, value, checkAlive)
}

// Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *PhysicsP2BodyDebug) CheckAll2O(key string, value interface{}, checkAlive bool, checkVisible bool) {
    self.Object.Call("checkAll", key, value, checkAlive, checkVisible)
}

// Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *PhysicsP2BodyDebug) CheckAll3O(key string, value interface{}, checkAlive bool, checkVisible bool, force bool) {
    self.Object.Call("checkAll", key, value, checkAlive, checkVisible, force)
}

// Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *PhysicsP2BodyDebug) CheckAllI(args ...interface{}) {
    self.Object.Call("checkAll", args)
}

// Adds the amount to the given property on all children in this group.
// 
// `Group.addAll('x', 10)` will add 10 to the child.x value for each child.
func (self *PhysicsP2BodyDebug) AddAll(property string, amount int, checkAlive bool, checkVisible bool) {
    self.Object.Call("addAll", property, amount, checkAlive, checkVisible)
}

// Adds the amount to the given property on all children in this group.
// 
// `Group.addAll('x', 10)` will add 10 to the child.x value for each child.
func (self *PhysicsP2BodyDebug) AddAllI(args ...interface{}) {
    self.Object.Call("addAll", args)
}

// Subtracts the amount from the given property on all children in this group.
// 
// `Group.subAll('x', 10)` will minus 10 from the child.x value for each child.
func (self *PhysicsP2BodyDebug) SubAll(property string, amount int, checkAlive bool, checkVisible bool) {
    self.Object.Call("subAll", property, amount, checkAlive, checkVisible)
}

// Subtracts the amount from the given property on all children in this group.
// 
// `Group.subAll('x', 10)` will minus 10 from the child.x value for each child.
func (self *PhysicsP2BodyDebug) SubAllI(args ...interface{}) {
    self.Object.Call("subAll", args)
}

// Multiplies the given property by the amount on all children in this group.
// 
// `Group.multiplyAll('x', 2)` will x2 the child.x value for each child.
func (self *PhysicsP2BodyDebug) MultiplyAll(property string, amount int, checkAlive bool, checkVisible bool) {
    self.Object.Call("multiplyAll", property, amount, checkAlive, checkVisible)
}

// Multiplies the given property by the amount on all children in this group.
// 
// `Group.multiplyAll('x', 2)` will x2 the child.x value for each child.
func (self *PhysicsP2BodyDebug) MultiplyAllI(args ...interface{}) {
    self.Object.Call("multiplyAll", args)
}

// Divides the given property by the amount on all children in this group.
// 
// `Group.divideAll('x', 2)` will half the child.x value for each child.
func (self *PhysicsP2BodyDebug) DivideAll(property string, amount int, checkAlive bool, checkVisible bool) {
    self.Object.Call("divideAll", property, amount, checkAlive, checkVisible)
}

// Divides the given property by the amount on all children in this group.
// 
// `Group.divideAll('x', 2)` will half the child.x value for each child.
func (self *PhysicsP2BodyDebug) DivideAllI(args ...interface{}) {
    self.Object.Call("divideAll", args)
}

// Calls a function, specified by name, on all children in the group who exist (or do not exist).
// 
// After the existsValue parameter you can add as many parameters as you like, which will all be passed to the child callback.
func (self *PhysicsP2BodyDebug) CallAllExists(callback string, existsValue bool, parameter interface{}) {
    self.Object.Call("callAllExists", callback, existsValue, parameter)
}

// Calls a function, specified by name, on all children in the group who exist (or do not exist).
// 
// After the existsValue parameter you can add as many parameters as you like, which will all be passed to the child callback.
func (self *PhysicsP2BodyDebug) CallAllExistsI(args ...interface{}) {
    self.Object.Call("callAllExists", args)
}

// Returns a reference to a function that exists on a child of the group based on the given callback array.
func (self *PhysicsP2BodyDebug) CallbackFromArray(child interface{}, callback []interface{}, length int) {
    self.Object.Call("callbackFromArray", child, callback, length)
}

// Returns a reference to a function that exists on a child of the group based on the given callback array.
func (self *PhysicsP2BodyDebug) CallbackFromArrayI(args ...interface{}) {
    self.Object.Call("callbackFromArray", args)
}

// Calls a function, specified by name, on all on children.
// 
// The function is called for all children regardless if they are dead or alive (see callAllExists for different options).
// After the method parameter and context you can add as many extra parameters as you like, which will all be passed to the child.
func (self *PhysicsP2BodyDebug) CallAll(method string, context string, args interface{}) {
    self.Object.Call("callAll", method, context, args)
}

// Calls a function, specified by name, on all on children.
// 
// The function is called for all children regardless if they are dead or alive (see callAllExists for different options).
// After the method parameter and context you can add as many extra parameters as you like, which will all be passed to the child.
func (self *PhysicsP2BodyDebug) CallAllI(args ...interface{}) {
    self.Object.Call("callAll", args)
}

// The core preUpdate - as called by World.
func (self *PhysicsP2BodyDebug) PreUpdate() {
    self.Object.Call("preUpdate")
}

// The core preUpdate - as called by World.
func (self *PhysicsP2BodyDebug) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// The core update - as called by World.
func (self *PhysicsP2BodyDebug) Update() {
    self.Object.Call("update")
}

// The core update - as called by World.
func (self *PhysicsP2BodyDebug) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// The core postUpdate - as called by World.
func (self *PhysicsP2BodyDebug) PostUpdate() {
    self.Object.Call("postUpdate")
}

// The core postUpdate - as called by World.
func (self *PhysicsP2BodyDebug) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Find children matching a certain predicate.
// 
// For example:
// 
//     var healthyList = Group.filter(function(child, index, children) {
//         return child.health > 10 ? true : false;
//     }, true);
//     healthyList.callAll('attack');
// 
// Note: Currently this will skip any children which are Groups themselves.
func (self *PhysicsP2BodyDebug) Filter(predicate func(...interface{})) *ArraySet{
    return &ArraySet{self.Object.Call("filter", predicate)}
}

// Find children matching a certain predicate.
// 
// For example:
// 
//     var healthyList = Group.filter(function(child, index, children) {
//         return child.health > 10 ? true : false;
//     }, true);
//     healthyList.callAll('attack');
// 
// Note: Currently this will skip any children which are Groups themselves.
func (self *PhysicsP2BodyDebug) Filter1O(predicate func(...interface{}), checkExists bool) *ArraySet{
    return &ArraySet{self.Object.Call("filter", predicate, checkExists)}
}

// Find children matching a certain predicate.
// 
// For example:
// 
//     var healthyList = Group.filter(function(child, index, children) {
//         return child.health > 10 ? true : false;
//     }, true);
//     healthyList.callAll('attack');
// 
// Note: Currently this will skip any children which are Groups themselves.
func (self *PhysicsP2BodyDebug) FilterI(args ...interface{}) *ArraySet{
    return &ArraySet{self.Object.Call("filter", args)}
}

// Call a function on each child in this group.
// 
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
// 
//     Group.forEach(awardBonusGold, this, true, 100, 500)
// 
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
// 
// Note: This check will skip any children which are Groups themselves.
func (self *PhysicsP2BodyDebug) ForEach(callback func(...interface{}), callbackContext interface{}) {
    self.Object.Call("forEach", callback, callbackContext)
}

// Call a function on each child in this group.
// 
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
// 
//     Group.forEach(awardBonusGold, this, true, 100, 500)
// 
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
// 
// Note: This check will skip any children which are Groups themselves.
func (self *PhysicsP2BodyDebug) ForEach1O(callback func(...interface{}), callbackContext interface{}, checkExists bool) {
    self.Object.Call("forEach", callback, callbackContext, checkExists)
}

// Call a function on each child in this group.
// 
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
// 
//     Group.forEach(awardBonusGold, this, true, 100, 500)
// 
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
// 
// Note: This check will skip any children which are Groups themselves.
func (self *PhysicsP2BodyDebug) ForEach2O(callback func(...interface{}), callbackContext interface{}, checkExists bool, args interface{}) {
    self.Object.Call("forEach", callback, callbackContext, checkExists, args)
}

// Call a function on each child in this group.
// 
// Additional arguments for the callback can be specified after the `checkExists` parameter. For example,
// 
//     Group.forEach(awardBonusGold, this, true, 100, 500)
// 
// would invoke `awardBonusGold` function with the parameters `(child, 100, 500)`.
// 
// Note: This check will skip any children which are Groups themselves.
func (self *PhysicsP2BodyDebug) ForEachI(args ...interface{}) {
    self.Object.Call("forEach", args)
}

// Call a function on each existing child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *PhysicsP2BodyDebug) ForEachExists(callback func(...interface{}), callbackContext interface{}) {
    self.Object.Call("forEachExists", callback, callbackContext)
}

// Call a function on each existing child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *PhysicsP2BodyDebug) ForEachExists1O(callback func(...interface{}), callbackContext interface{}, args interface{}) {
    self.Object.Call("forEachExists", callback, callbackContext, args)
}

// Call a function on each existing child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *PhysicsP2BodyDebug) ForEachExistsI(args ...interface{}) {
    self.Object.Call("forEachExists", args)
}

// Call a function on each alive child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *PhysicsP2BodyDebug) ForEachAlive(callback func(...interface{}), callbackContext interface{}) {
    self.Object.Call("forEachAlive", callback, callbackContext)
}

// Call a function on each alive child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *PhysicsP2BodyDebug) ForEachAlive1O(callback func(...interface{}), callbackContext interface{}, args interface{}) {
    self.Object.Call("forEachAlive", callback, callbackContext, args)
}

// Call a function on each alive child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *PhysicsP2BodyDebug) ForEachAliveI(args ...interface{}) {
    self.Object.Call("forEachAlive", args)
}

// Call a function on each dead child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *PhysicsP2BodyDebug) ForEachDead(callback func(...interface{}), callbackContext interface{}) {
    self.Object.Call("forEachDead", callback, callbackContext)
}

// Call a function on each dead child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *PhysicsP2BodyDebug) ForEachDead1O(callback func(...interface{}), callbackContext interface{}, args interface{}) {
    self.Object.Call("forEachDead", callback, callbackContext, args)
}

// Call a function on each dead child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *PhysicsP2BodyDebug) ForEachDeadI(args ...interface{}) {
    self.Object.Call("forEachDead", args)
}

// Sort the children in the group according to a particular key and ordering.
// 
// Call this function to sort the group according to a particular key value and order.
// 
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
// 
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *PhysicsP2BodyDebug) Sort() {
    self.Object.Call("sort")
}

// Sort the children in the group according to a particular key and ordering.
// 
// Call this function to sort the group according to a particular key value and order.
// 
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
// 
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *PhysicsP2BodyDebug) Sort1O(key string) {
    self.Object.Call("sort", key)
}

// Sort the children in the group according to a particular key and ordering.
// 
// Call this function to sort the group according to a particular key value and order.
// 
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
// 
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *PhysicsP2BodyDebug) Sort2O(key string, order int) {
    self.Object.Call("sort", key, order)
}

// Sort the children in the group according to a particular key and ordering.
// 
// Call this function to sort the group according to a particular key value and order.
// 
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
// 
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *PhysicsP2BodyDebug) SortI(args ...interface{}) {
    self.Object.Call("sort", args)
}

// Sort the children in the group according to custom sort function.
// 
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *PhysicsP2BodyDebug) CustomSort(sortHandler func(...interface{})) {
    self.Object.Call("customSort", sortHandler)
}

// Sort the children in the group according to custom sort function.
// 
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *PhysicsP2BodyDebug) CustomSort1O(sortHandler func(...interface{}), context interface{}) {
    self.Object.Call("customSort", sortHandler, context)
}

// Sort the children in the group according to custom sort function.
// 
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *PhysicsP2BodyDebug) CustomSortI(args ...interface{}) {
    self.Object.Call("customSort", args)
}

// An internal helper function for the sort process.
func (self *PhysicsP2BodyDebug) AscendingSortHandler(a interface{}, b interface{}) {
    self.Object.Call("ascendingSortHandler", a, b)
}

// An internal helper function for the sort process.
func (self *PhysicsP2BodyDebug) AscendingSortHandlerI(args ...interface{}) {
    self.Object.Call("ascendingSortHandler", args)
}

// An internal helper function for the sort process.
func (self *PhysicsP2BodyDebug) DescendingSortHandler(a interface{}, b interface{}) {
    self.Object.Call("descendingSortHandler", a, b)
}

// An internal helper function for the sort process.
func (self *PhysicsP2BodyDebug) DescendingSortHandlerI(args ...interface{}) {
    self.Object.Call("descendingSortHandler", args)
}

// Iterates over the children of the group performing one of several actions for matched children.
// 
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
// 
// The result depends on the `returnType`:
// 
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
// 
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *PhysicsP2BodyDebug) Iterate(key string, value interface{}, returnType int) interface{}{
    return self.Object.Call("iterate", key, value, returnType)
}

// Iterates over the children of the group performing one of several actions for matched children.
// 
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
// 
// The result depends on the `returnType`:
// 
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
// 
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *PhysicsP2BodyDebug) Iterate1O(key string, value interface{}, returnType int, callback func(...interface{})) interface{}{
    return self.Object.Call("iterate", key, value, returnType, callback)
}

// Iterates over the children of the group performing one of several actions for matched children.
// 
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
// 
// The result depends on the `returnType`:
// 
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
// 
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *PhysicsP2BodyDebug) Iterate2O(key string, value interface{}, returnType int, callback func(...interface{}), callbackContext interface{}) interface{}{
    return self.Object.Call("iterate", key, value, returnType, callback, callbackContext)
}

// Iterates over the children of the group performing one of several actions for matched children.
// 
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
// 
// The result depends on the `returnType`:
// 
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
// 
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *PhysicsP2BodyDebug) Iterate3O(key string, value interface{}, returnType int, callback func(...interface{}), callbackContext interface{}, args []interface{}) interface{}{
    return self.Object.Call("iterate", key, value, returnType, callback, callbackContext, args)
}

// Iterates over the children of the group performing one of several actions for matched children.
// 
// A child is considered a match when it has a property, named `key`, whose value is equal to `value`
// according to a strict equality comparison.
// 
// The result depends on the `returnType`:
// 
// - {@link Phaser.Group.RETURN_TOTAL RETURN_TOTAL}:
//     The callback, if any, is applied to all matching children. The number of matched children is returned.
// - {@link Phaser.Group.RETURN_NONE RETURN_NONE}:
//     The callback, if any, is applied to all matching children. No value is returned.
// - {@link Phaser.Group.RETURN_CHILD RETURN_CHILD}:
//     The callback, if any, is applied to the *first* matching child and the *first* matched child is returned.
//     If there is no matching child then null is returned.
// 
// If `args` is specified it must be an array. The matched child will be assigned to the first
// element and the entire array will be applied to the callback function.
func (self *PhysicsP2BodyDebug) IterateI(args ...interface{}) interface{}{
    return self.Object.Call("iterate", args)
}

// Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstExists() *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists")}
}

// Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstExists1O(exists bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists)}
}

// Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstExists2O(exists bool, createIfNull bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull)}
}

// Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstExists3O(exists bool, createIfNull bool, x int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x)}
}

// Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstExists4O(exists bool, createIfNull bool, x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x, y)}
}

// Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstExists5O(exists bool, createIfNull bool, x int, y int, key interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x, y, key)}
}

// Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstExists6O(exists bool, createIfNull bool, x int, y int, key interface{}, frame interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", exists, createIfNull, x, y, key, frame)}
}

// Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstExistsI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstExists", args)}
}

// Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstAlive() *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive")}
}

// Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstAlive1O(createIfNull bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull)}
}

// Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstAlive2O(createIfNull bool, x int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x)}
}

// Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstAlive3O(createIfNull bool, x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x, y)}
}

// Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstAlive4O(createIfNull bool, x int, y int, key interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x, y, key)}
}

// Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstAlive5O(createIfNull bool, x int, y int, key interface{}, frame interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", createIfNull, x, y, key, frame)}
}

// Get the first child that is alive (`child.alive === true`).
// 
// This is handy for choosing a squad leader, etc.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no alive ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstAliveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstAlive", args)}
}

// Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstDead() *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead")}
}

// Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstDead1O(createIfNull bool) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", createIfNull)}
}

// Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstDead2O(createIfNull bool, x int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x)}
}

// Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstDead3O(createIfNull bool, x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x, y)}
}

// Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstDead4O(createIfNull bool, x int, y int, key interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x, y, key)}
}

// Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstDead5O(createIfNull bool, x int, y int, key interface{}, frame interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", createIfNull, x, y, key, frame)}
}

// Get the first child that is dead (`child.alive === false`).
// 
// This is handy for checking if everything has been wiped out and adding to the pool as needed.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if no dead ones were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *PhysicsP2BodyDebug) GetFirstDeadI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getFirstDead", args)}
}

// Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *PhysicsP2BodyDebug) ResetChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", child)}
}

// Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *PhysicsP2BodyDebug) ResetChild1O(child *DisplayObject, x int) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", child, x)}
}

// Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *PhysicsP2BodyDebug) ResetChild2O(child *DisplayObject, x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", child, x, y)}
}

// Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *PhysicsP2BodyDebug) ResetChild3O(child *DisplayObject, x int, y int, key interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", child, x, y, key)}
}

// Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *PhysicsP2BodyDebug) ResetChild4O(child *DisplayObject, x int, y int, key interface{}, frame interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", child, x, y, key, frame)}
}

// Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *PhysicsP2BodyDebug) ResetChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("resetChild", args)}
}

// Return the child at the top of this group.
// 
// The top child is the child displayed (rendered) above every other child.
func (self *PhysicsP2BodyDebug) GetTop() interface{}{
    return self.Object.Call("getTop")
}

// Return the child at the top of this group.
// 
// The top child is the child displayed (rendered) above every other child.
func (self *PhysicsP2BodyDebug) GetTopI(args ...interface{}) interface{}{
    return self.Object.Call("getTop", args)
}

// Returns the child at the bottom of this group.
// 
// The bottom child the child being displayed (rendered) below every other child.
func (self *PhysicsP2BodyDebug) GetBottom() interface{}{
    return self.Object.Call("getBottom")
}

// Returns the child at the bottom of this group.
// 
// The bottom child the child being displayed (rendered) below every other child.
func (self *PhysicsP2BodyDebug) GetBottomI(args ...interface{}) interface{}{
    return self.Object.Call("getBottom", args)
}

// Get the closest child to given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *PhysicsP2BodyDebug) GetClosestTo(object interface{}) interface{}{
    return self.Object.Call("getClosestTo", object)
}

// Get the closest child to given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *PhysicsP2BodyDebug) GetClosestTo1O(object interface{}, callback func(...interface{})) interface{}{
    return self.Object.Call("getClosestTo", object, callback)
}

// Get the closest child to given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *PhysicsP2BodyDebug) GetClosestTo2O(object interface{}, callback func(...interface{}), callbackContext interface{}) interface{}{
    return self.Object.Call("getClosestTo", object, callback, callbackContext)
}

// Get the closest child to given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'close' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *PhysicsP2BodyDebug) GetClosestToI(args ...interface{}) interface{}{
    return self.Object.Call("getClosestTo", args)
}

// Get the child furthest away from the given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *PhysicsP2BodyDebug) GetFurthestFrom(object interface{}) interface{}{
    return self.Object.Call("getFurthestFrom", object)
}

// Get the child furthest away from the given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *PhysicsP2BodyDebug) GetFurthestFrom1O(object interface{}, callback func(...interface{})) interface{}{
    return self.Object.Call("getFurthestFrom", object, callback)
}

// Get the child furthest away from the given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *PhysicsP2BodyDebug) GetFurthestFrom2O(object interface{}, callback func(...interface{}), callbackContext interface{}) interface{}{
    return self.Object.Call("getFurthestFrom", object, callback, callbackContext)
}

// Get the child furthest away from the given Object, with optional callback to filter children.
// 
// This can be a Sprite, Group, Image or any object with public x and y properties.
// 
// 'furthest away' is determined by the distance from the objects `x` and `y` properties compared to the childs `x` and `y` properties.
// 
// You can use the optional `callback` argument to apply your own filter to the distance checks.
// If the child is closer then the previous child, it will be sent to `callback` as the first argument,
// with the distance as the second. The callback should return `true` if it passes your 
// filtering criteria, otherwise it should return `false`.
func (self *PhysicsP2BodyDebug) GetFurthestFromI(args ...interface{}) interface{}{
    return self.Object.Call("getFurthestFrom", args)
}

// Get the number of living children in this group.
func (self *PhysicsP2BodyDebug) CountLiving() int{
    return self.Object.Call("countLiving").Int()
}

// Get the number of living children in this group.
func (self *PhysicsP2BodyDebug) CountLivingI(args ...interface{}) int{
    return self.Object.Call("countLiving", args).Int()
}

// Get the number of dead children in this group.
func (self *PhysicsP2BodyDebug) CountDead() int{
    return self.Object.Call("countDead").Int()
}

// Get the number of dead children in this group.
func (self *PhysicsP2BodyDebug) CountDeadI(args ...interface{}) int{
    return self.Object.Call("countDead", args).Int()
}

// Returns a random child from the group.
func (self *PhysicsP2BodyDebug) GetRandom() interface{}{
    return self.Object.Call("getRandom")
}

// Returns a random child from the group.
func (self *PhysicsP2BodyDebug) GetRandom1O(startIndex int) interface{}{
    return self.Object.Call("getRandom", startIndex)
}

// Returns a random child from the group.
func (self *PhysicsP2BodyDebug) GetRandom2O(startIndex int, length int) interface{}{
    return self.Object.Call("getRandom", startIndex, length)
}

// Returns a random child from the group.
func (self *PhysicsP2BodyDebug) GetRandomI(args ...interface{}) interface{}{
    return self.Object.Call("getRandom", args)
}

// Removes the given child from this group.
// 
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
// 
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *PhysicsP2BodyDebug) Remove(child interface{}) bool{
    return self.Object.Call("remove", child).Bool()
}

// Removes the given child from this group.
// 
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
// 
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *PhysicsP2BodyDebug) Remove1O(child interface{}, destroy bool) bool{
    return self.Object.Call("remove", child, destroy).Bool()
}

// Removes the given child from this group.
// 
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
// 
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *PhysicsP2BodyDebug) Remove2O(child interface{}, destroy bool, silent bool) bool{
    return self.Object.Call("remove", child, destroy, silent).Bool()
}

// Removes the given child from this group.
// 
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
// 
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *PhysicsP2BodyDebug) RemoveI(args ...interface{}) bool{
    return self.Object.Call("remove", args).Bool()
}

// Moves all children from this Group to the Group given.
func (self *PhysicsP2BodyDebug) MoveAll(group *Group) *Group{
    return &Group{self.Object.Call("moveAll", group)}
}

// Moves all children from this Group to the Group given.
func (self *PhysicsP2BodyDebug) MoveAll1O(group *Group, silent bool) *Group{
    return &Group{self.Object.Call("moveAll", group, silent)}
}

// Moves all children from this Group to the Group given.
func (self *PhysicsP2BodyDebug) MoveAllI(args ...interface{}) *Group{
    return &Group{self.Object.Call("moveAll", args)}
}

// Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *PhysicsP2BodyDebug) RemoveAll() {
    self.Object.Call("removeAll")
}

// Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *PhysicsP2BodyDebug) RemoveAll1O(destroy bool) {
    self.Object.Call("removeAll", destroy)
}

// Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *PhysicsP2BodyDebug) RemoveAll2O(destroy bool, silent bool) {
    self.Object.Call("removeAll", destroy, silent)
}

// Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *PhysicsP2BodyDebug) RemoveAll3O(destroy bool, silent bool, destroyTexture bool) {
    self.Object.Call("removeAll", destroy, silent, destroyTexture)
}

// Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *PhysicsP2BodyDebug) RemoveAllI(args ...interface{}) {
    self.Object.Call("removeAll", args)
}

// Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *PhysicsP2BodyDebug) RemoveBetween(startIndex int) {
    self.Object.Call("removeBetween", startIndex)
}

// Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *PhysicsP2BodyDebug) RemoveBetween1O(startIndex int, endIndex int) {
    self.Object.Call("removeBetween", startIndex, endIndex)
}

// Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *PhysicsP2BodyDebug) RemoveBetween2O(startIndex int, endIndex int, destroy bool) {
    self.Object.Call("removeBetween", startIndex, endIndex, destroy)
}

// Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *PhysicsP2BodyDebug) RemoveBetween3O(startIndex int, endIndex int, destroy bool, silent bool) {
    self.Object.Call("removeBetween", startIndex, endIndex, destroy, silent)
}

// Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *PhysicsP2BodyDebug) RemoveBetweenI(args ...interface{}) {
    self.Object.Call("removeBetween", args)
}

// Destroys this group.
// 
// Removes all children, then removes this group from its parent and nulls references.
func (self *PhysicsP2BodyDebug) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this group.
// 
// Removes all children, then removes this group from its parent and nulls references.
func (self *PhysicsP2BodyDebug) Destroy1O(destroyChildren bool) {
    self.Object.Call("destroy", destroyChildren)
}

// Destroys this group.
// 
// Removes all children, then removes this group from its parent and nulls references.
func (self *PhysicsP2BodyDebug) Destroy2O(destroyChildren bool, soft bool) {
    self.Object.Call("destroy", destroyChildren, soft)
}

// Destroys this group.
// 
// Removes all children, then removes this group from its parent and nulls references.
func (self *PhysicsP2BodyDebug) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *PhysicsP2BodyDebug) AlignIn(container interface{}) *Group{
    return &Group{self.Object.Call("alignIn", container)}
}

// Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *PhysicsP2BodyDebug) AlignIn1O(container interface{}, position int) *Group{
    return &Group{self.Object.Call("alignIn", container, position)}
}

// Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *PhysicsP2BodyDebug) AlignIn2O(container interface{}, position int, offsetX int) *Group{
    return &Group{self.Object.Call("alignIn", container, position, offsetX)}
}

// Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *PhysicsP2BodyDebug) AlignIn3O(container interface{}, position int, offsetX int, offsetY int) *Group{
    return &Group{self.Object.Call("alignIn", container, position, offsetX, offsetY)}
}

// Aligns this Group within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation and scale of its children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *PhysicsP2BodyDebug) AlignInI(args ...interface{}) *Group{
    return &Group{self.Object.Call("alignIn", args)}
}

// Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *PhysicsP2BodyDebug) AlignTo(parent interface{}) *Group{
    return &Group{self.Object.Call("alignTo", parent)}
}

// Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *PhysicsP2BodyDebug) AlignTo1O(parent interface{}, position int) *Group{
    return &Group{self.Object.Call("alignTo", parent, position)}
}

// Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *PhysicsP2BodyDebug) AlignTo2O(parent interface{}, position int, offsetX int) *Group{
    return &Group{self.Object.Call("alignTo", parent, position, offsetX)}
}

// Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *PhysicsP2BodyDebug) AlignTo3O(parent interface{}, position int, offsetX int, offsetY int) *Group{
    return &Group{self.Object.Call("alignTo", parent, position, offsetX, offsetY)}
}

// Aligns this Group to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Groups within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Group to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// Groups are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation and scale of the children.
// This allows you to neatly align Groups, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Group. For example:
// 
// `group.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `group` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *PhysicsP2BodyDebug) AlignToI(args ...interface{}) *Group{
    return &Group{self.Object.Call("alignTo", args)}
}

// Adds a child to the container.
func (self *PhysicsP2BodyDebug) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// Adds a child to the container.
func (self *PhysicsP2BodyDebug) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *PhysicsP2BodyDebug) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *PhysicsP2BodyDebug) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *PhysicsP2BodyDebug) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// Swaps the position of 2 Display Objects within this container.
func (self *PhysicsP2BodyDebug) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *PhysicsP2BodyDebug) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// Returns the index position of a child DisplayObject instance
func (self *PhysicsP2BodyDebug) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *PhysicsP2BodyDebug) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// Changes the position of an existing child in the display object container
func (self *PhysicsP2BodyDebug) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *PhysicsP2BodyDebug) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// Returns the child at the specified index
func (self *PhysicsP2BodyDebug) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *PhysicsP2BodyDebug) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// Removes a child from the container.
func (self *PhysicsP2BodyDebug) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *PhysicsP2BodyDebug) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// Removes a child from the specified index position.
func (self *PhysicsP2BodyDebug) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *PhysicsP2BodyDebug) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// Removes all children from this container that are within the begin and end indexes.
func (self *PhysicsP2BodyDebug) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *PhysicsP2BodyDebug) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *PhysicsP2BodyDebug) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *PhysicsP2BodyDebug) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *PhysicsP2BodyDebug) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *PhysicsP2BodyDebug) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *PhysicsP2BodyDebug) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *PhysicsP2BodyDebug) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// Removes the current stage reference from the container and all of its children.
func (self *PhysicsP2BodyDebug) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// Renders the object using the WebGL renderer
func (self *PhysicsP2BodyDebug) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// Renders the object using the WebGL renderer
func (self *PhysicsP2BodyDebug) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *PhysicsP2BodyDebug) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// Renders the object using the Canvas renderer
func (self *PhysicsP2BodyDebug) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}
