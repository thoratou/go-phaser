// Automatic generation for Phaser.World
// generated file World.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// "This world is but a canvas to our imagination." - Henry David Thoreau
// 
// A game has only one world. The world is an abstract place in which all game objects live. It is not bound
// by stage limits and can be any size. You look into the world via cameras. All game objects live within
// the world at world-based coordinates. By default a world is created the same size as your Stage.
type World struct {
    *js.Object
}


// The World has no fixed size, but it does have a bounds outside of which objects are no longer considered as being "in world" and you should use this to clean-up the display list and purge dead objects.
// By default we set the Bounds to be from 0,0 to Game.width,Game.height. I.e. it will match the size given to the game constructor with 0,0 representing the top-left of the display.
// However 0,0 is actually the center of the world, and if you rotate or scale the world all of that will happen from 0,0.
// So if you want to make a game in which the world itself will rotate you should adjust the bounds so that 0,0 is the center point, i.e. set them to -1000,-1000,2000,2000 for a 2000x2000 sized world centered around 0,0. Bound of this world that objects can not escape from.
func (self *World) GetBounds() *Rectangle{
    return &Rectangle{self.Get("bounds")}
}

// The World has no fixed size, but it does have a bounds outside of which objects are no longer considered as being "in world" and you should use this to clean-up the display list and purge dead objects.
// By default we set the Bounds to be from 0,0 to Game.width,Game.height. I.e. it will match the size given to the game constructor with 0,0 representing the top-left of the display.
// However 0,0 is actually the center of the world, and if you rotate or scale the world all of that will happen from 0,0.
// So if you want to make a game in which the world itself will rotate you should adjust the bounds so that 0,0 is the center point, i.e. set them to -1000,-1000,2000,2000 for a 2000x2000 sized world centered around 0,0. Bound of this world that objects can not escape from.
func (self *World) SetBounds(member *Rectangle) {
    self.Set("bounds", member)
}

// Camera instance.
func (self *World) GetCamera() *Camera{
    return &Camera{self.Get("camera")}
}

// Camera instance.
func (self *World) SetCamera(member *Camera) {
    self.Set("camera", member)
}

// True if the World has been given a specifically defined size (i.e. from a Tilemap or direct in code) or false if it's just matched to the Game dimensions.
func (self *World) Get_definedSize() bool{
    return self.Get("_definedSize").Bool()
}

// True if the World has been given a specifically defined size (i.e. from a Tilemap or direct in code) or false if it's just matched to the Game dimensions.
func (self *World) Set_definedSize(member bool) {
    self.Set("_definedSize", member)
}

// The defined width of the World. Sometimes the bounds needs to grow larger than this (if you resize the game) but this retains the original requested dimension.
func (self *World) Get_width() interface{}{
    return self.Get("_width")
}

// The defined width of the World. Sometimes the bounds needs to grow larger than this (if you resize the game) but this retains the original requested dimension.
func (self *World) Set_width(member interface{}) {
    self.Set("_width", member)
}

// The defined height of the World. Sometimes the bounds needs to grow larger than this (if you resize the game) but this retains the original requested dimension.
func (self *World) Get_height() interface{}{
    return self.Get("_height")
}

// The defined height of the World. Sometimes the bounds needs to grow larger than this (if you resize the game) but this retains the original requested dimension.
func (self *World) Set_height(member interface{}) {
    self.Set("_height", member)
}

// Gets or sets the current width of the game world. The world can never be smaller than the game (canvas) dimensions.
func (self *World) GetWidth() float64{
    return self.Get("width").Float()
}

// Gets or sets the current width of the game world. The world can never be smaller than the game (canvas) dimensions.
func (self *World) SetWidth(member float64) {
    self.Set("width", member)
}

// Gets or sets the current height of the game world. The world can never be smaller than the game (canvas) dimensions.
func (self *World) GetHeight() float64{
    return self.Get("height").Float()
}

// Gets or sets the current height of the game world. The world can never be smaller than the game (canvas) dimensions.
func (self *World) SetHeight(member float64) {
    self.Set("height", member)
}

// Gets the X position corresponding to the center point of the world.
func (self *World) GetCenterX() float64{
    return self.Get("centerX").Float()
}

// Gets the X position corresponding to the center point of the world.
func (self *World) SetCenterX(member float64) {
    self.Set("centerX", member)
}

// Gets the Y position corresponding to the center point of the world.
func (self *World) GetCenterY() float64{
    return self.Get("centerY").Float()
}

// Gets the Y position corresponding to the center point of the world.
func (self *World) SetCenterY(member float64) {
    self.Set("centerY", member)
}

// Gets a random integer which is lesser than or equal to the current width of the game world.
func (self *World) GetRandomX() float64{
    return self.Get("randomX").Float()
}

// Gets a random integer which is lesser than or equal to the current width of the game world.
func (self *World) SetRandomX(member float64) {
    self.Set("randomX", member)
}

// Gets a random integer which is lesser than or equal to the current height of the game world.
func (self *World) GetRandomY() float64{
    return self.Get("randomY").Float()
}

// Gets a random integer which is lesser than or equal to the current height of the game world.
func (self *World) SetRandomY(member float64) {
    self.Set("randomY", member)
}

// A reference to the currently running Game.
func (self *World) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *World) SetGame(member *Game) {
    self.Set("game", member)
}

// A name for this group. Not used internally but useful for debugging.
func (self *World) GetName() string{
    return self.Get("name").String()
}

// A name for this group. Not used internally but useful for debugging.
func (self *World) SetName(member string) {
    self.Set("name", member)
}

// The z-depth value of this object within its parent container/Group - the World is a Group as well.
// This value must be unique for each child in a Group.
func (self *World) GetZ() int{
    return self.Get("z").Int()
}

// The z-depth value of this object within its parent container/Group - the World is a Group as well.
// This value must be unique for each child in a Group.
func (self *World) SetZ(member int) {
    self.Set("z", member)
}

// Internal Phaser Type value.
func (self *World) GetType() int{
    return self.Get("type").Int()
}

// Internal Phaser Type value.
func (self *World) SetType(member int) {
    self.Set("type", member)
}

// The const physics body type of this object.
func (self *World) GetPhysicsType() float64{
    return self.Get("physicsType").Float()
}

// The const physics body type of this object.
func (self *World) SetPhysicsType(member float64) {
    self.Set("physicsType", member)
}

// The alive property is useful for Groups that are children of other Groups and need to be included/excluded in checks like forEachAlive.
func (self *World) GetAlive() bool{
    return self.Get("alive").Bool()
}

// The alive property is useful for Groups that are children of other Groups and need to be included/excluded in checks like forEachAlive.
func (self *World) SetAlive(member bool) {
    self.Set("alive", member)
}

// If exists is true the group is updated, otherwise it is skipped.
func (self *World) GetExists() bool{
    return self.Get("exists").Bool()
}

// If exists is true the group is updated, otherwise it is skipped.
func (self *World) SetExists(member bool) {
    self.Set("exists", member)
}

// A group with `ignoreDestroy` set to `true` ignores all calls to its `destroy` method.
func (self *World) GetIgnoreDestroy() bool{
    return self.Get("ignoreDestroy").Bool()
}

// A group with `ignoreDestroy` set to `true` ignores all calls to its `destroy` method.
func (self *World) SetIgnoreDestroy(member bool) {
    self.Set("ignoreDestroy", member)
}

// A Group is that has `pendingDestroy` set to `true` is flagged to have its destroy method
// called on the next logic update.
// You can set it directly to flag the Group to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy a Group from within one of its own callbacks
// or a callback of one of its children.
func (self *World) GetPendingDestroy() bool{
    return self.Get("pendingDestroy").Bool()
}

// A Group is that has `pendingDestroy` set to `true` is flagged to have its destroy method
// called on the next logic update.
// You can set it directly to flag the Group to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy a Group from within one of its own callbacks
// or a callback of one of its children.
func (self *World) SetPendingDestroy(member bool) {
    self.Set("pendingDestroy", member)
}

// The type of objects that will be created when using {@link Phaser.Group#create create} or {@link Phaser.Group#createMultiple createMultiple}.
// 
// Any object may be used but it should extend either Sprite or Image and accept the same constructor arguments:
// when a new object is created it is passed the following parameters to its constructor: `(game, x, y, key, frame)`.
func (self *World) GetClassType() interface{}{
    return self.Get("classType")
}

// The type of objects that will be created when using {@link Phaser.Group#create create} or {@link Phaser.Group#createMultiple createMultiple}.
// 
// Any object may be used but it should extend either Sprite or Image and accept the same constructor arguments:
// when a new object is created it is passed the following parameters to its constructor: `(game, x, y, key, frame)`.
func (self *World) SetClassType(member interface{}) {
    self.Set("classType", member)
}

// The current display object that the group cursor is pointing to, if any. (Can be set manually.)
// 
// The cursor is a way to iterate through the children in a Group using {@link Phaser.Group#next next} and {@link Phaser.Group#previous previous}.
func (self *World) GetCursor() *DisplayObject{
    return &DisplayObject{self.Get("cursor")}
}

// The current display object that the group cursor is pointing to, if any. (Can be set manually.)
// 
// The cursor is a way to iterate through the children in a Group using {@link Phaser.Group#next next} and {@link Phaser.Group#previous previous}.
func (self *World) SetCursor(member *DisplayObject) {
    self.Set("cursor", member)
}

// A Group with `inputEnableChildren` set to `true` will automatically call `inputEnabled = true` 
// on any children _added_ to, or _created by_, this Group.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
func (self *World) GetInputEnableChildren() bool{
    return self.Get("inputEnableChildren").Bool()
}

// A Group with `inputEnableChildren` set to `true` will automatically call `inputEnabled = true` 
// on any children _added_ to, or _created by_, this Group.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
func (self *World) SetInputEnableChildren(member bool) {
    self.Set("inputEnableChildren", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputDown signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) GetOnChildInputDown() *Signal{
    return &Signal{self.Get("onChildInputDown")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputDown signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) SetOnChildInputDown(member *Signal) {
    self.Set("onChildInputDown", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputUp signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 3 arguments: A reference to the Sprite that triggered the signal, 
// a reference to the Pointer that caused it, and a boolean value `isOver` that tells you if the Pointer
// is still over the Sprite or not.
func (self *World) GetOnChildInputUp() *Signal{
    return &Signal{self.Get("onChildInputUp")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputUp signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 3 arguments: A reference to the Sprite that triggered the signal, 
// a reference to the Pointer that caused it, and a boolean value `isOver` that tells you if the Pointer
// is still over the Sprite or not.
func (self *World) SetOnChildInputUp(member *Signal) {
    self.Set("onChildInputUp", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputOver signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) GetOnChildInputOver() *Signal{
    return &Signal{self.Get("onChildInputOver")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputOver signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) SetOnChildInputOver(member *Signal) {
    self.Set("onChildInputOver", member)
}

// This Signal is dispatched whenever a child of this Group emits an onInputOut signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) GetOnChildInputOut() *Signal{
    return &Signal{self.Get("onChildInputOut")}
}

// This Signal is dispatched whenever a child of this Group emits an onInputOut signal as a result
// of having been interacted with by a Pointer. You can bind functions to this Signal instead of to
// every child Sprite.
// 
// This Signal is sent 2 arguments: A reference to the Sprite that triggered the signal, and
// a reference to the Pointer that caused it.
func (self *World) SetOnChildInputOut(member *Signal) {
    self.Set("onChildInputOut", member)
}

// If true all Sprites created by, or added to this group, will have a physics body enabled on them.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
// 
// The default body type is controlled with {@link Phaser.Group#physicsBodyType physicsBodyType}.
func (self *World) GetEnableBody() bool{
    return self.Get("enableBody").Bool()
}

// If true all Sprites created by, or added to this group, will have a physics body enabled on them.
// 
// If there are children already in the Group at the time you set this property, they are not changed.
// 
// The default body type is controlled with {@link Phaser.Group#physicsBodyType physicsBodyType}.
func (self *World) SetEnableBody(member bool) {
    self.Set("enableBody", member)
}

// If true when a physics body is created (via {@link Phaser.Group#enableBody enableBody}) it will create a physics debug object as well.
// 
// This only works for P2 bodies.
func (self *World) GetEnableBodyDebug() bool{
    return self.Get("enableBodyDebug").Bool()
}

// If true when a physics body is created (via {@link Phaser.Group#enableBody enableBody}) it will create a physics debug object as well.
// 
// This only works for P2 bodies.
func (self *World) SetEnableBodyDebug(member bool) {
    self.Set("enableBodyDebug", member)
}

// If {@link Phaser.Group#enableBody enableBody} is true this is the type of physics body that is created on new Sprites.
// 
// The valid values are {@link Phaser.Physics.ARCADE}, {@link Phaser.Physics.P2JS}, {@link Phaser.Physics.NINJA}, etc.
func (self *World) GetPhysicsBodyType() int{
    return self.Get("physicsBodyType").Int()
}

// If {@link Phaser.Group#enableBody enableBody} is true this is the type of physics body that is created on new Sprites.
// 
// The valid values are {@link Phaser.Physics.ARCADE}, {@link Phaser.Physics.P2JS}, {@link Phaser.Physics.NINJA}, etc.
func (self *World) SetPhysicsBodyType(member int) {
    self.Set("physicsBodyType", member)
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
func (self *World) GetPhysicsSortDirection() int{
    return self.Get("physicsSortDirection").Int()
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
func (self *World) SetPhysicsSortDirection(member int) {
    self.Set("physicsSortDirection", member)
}

// This signal is dispatched when the group is destroyed.
func (self *World) GetOnDestroy() *Signal{
    return &Signal{self.Get("onDestroy")}
}

// This signal is dispatched when the group is destroyed.
func (self *World) SetOnDestroy(member *Signal) {
    self.Set("onDestroy", member)
}

// The current index of the Group cursor. Advance it with Group.next.
func (self *World) GetCursorIndex() int{
    return self.Get("cursorIndex").Int()
}

// The current index of the Group cursor. Advance it with Group.next.
func (self *World) SetCursorIndex(member int) {
    self.Set("cursorIndex", member)
}

// A Group that is fixed to the camera uses its x/y coordinates as offsets from the top left of the camera. These are stored in Group.cameraOffset.
// 
// Note that the cameraOffset values are in addition to any parent in the display list.
// So if this Group was in a Group that has x: 200, then this will be added to the cameraOffset.x
func (self *World) GetFixedToCamera() bool{
    return self.Get("fixedToCamera").Bool()
}

// A Group that is fixed to the camera uses its x/y coordinates as offsets from the top left of the camera. These are stored in Group.cameraOffset.
// 
// Note that the cameraOffset values are in addition to any parent in the display list.
// So if this Group was in a Group that has x: 200, then this will be added to the cameraOffset.x
func (self *World) SetFixedToCamera(member bool) {
    self.Set("fixedToCamera", member)
}

// If this object is {@link Phaser.Group#fixedToCamera fixedToCamera} then this stores the x/y position offset relative to the top-left of the camera view.
// If the parent of this Group is also `fixedToCamera` then the offset here is in addition to that and should typically be disabled.
func (self *World) GetCameraOffset() *Point{
    return &Point{self.Get("cameraOffset")}
}

// If this object is {@link Phaser.Group#fixedToCamera fixedToCamera} then this stores the x/y position offset relative to the top-left of the camera view.
// If the parent of this Group is also `fixedToCamera` then the offset here is in addition to that and should typically be disabled.
func (self *World) SetCameraOffset(member *Point) {
    self.Set("cameraOffset", member)
}

// The hash array is an array belonging to this Group into which you can add any of its children via Group.addToHash and Group.removeFromHash.
// 
// Only children of this Group can be added to and removed from the hash.
// 
// This hash is used automatically by Phaser Arcade Physics in order to perform non z-index based destructive sorting.
// However if you don't use Arcade Physics, or this isn't a physics enabled Group, then you can use the hash to perform your own
// sorting and filtering of Group children without touching their z-index (and therefore display draw order)
func (self *World) GetHash() []interface{}{
	array := self.Get("hash")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The hash array is an array belonging to this Group into which you can add any of its children via Group.addToHash and Group.removeFromHash.
// 
// Only children of this Group can be added to and removed from the hash.
// 
// This hash is used automatically by Phaser Arcade Physics in order to perform non z-index based destructive sorting.
// However if you don't use Arcade Physics, or this isn't a physics enabled Group, then you can use the hash to perform your own
// sorting and filtering of Group children without touching their z-index (and therefore display draw order)
func (self *World) SetHash(member []interface{}) {
    self.Set("hash", member)
}

// Total number of existing children in the group.
func (self *World) GetTotal() int{
    return self.Get("total").Int()
}

// Total number of existing children in the group.
func (self *World) SetTotal(member int) {
    self.Set("total", member)
}

// Total number of children in this group, regardless of exists/alive status.
func (self *World) GetLength() int{
    return self.Get("length").Int()
}

// Total number of children in this group, regardless of exists/alive status.
func (self *World) SetLength(member int) {
    self.Set("length", member)
}

// The angle of rotation of the group container, in degrees.
// 
// This adjusts the group itself by modifying its local rotation transform.
// 
// This has no impact on the rotation/angle properties of the children, but it will update their worldTransform
// and on-screen orientation and position.
func (self *World) GetAngle() float64{
    return self.Get("angle").Float()
}

// The angle of rotation of the group container, in degrees.
// 
// This adjusts the group itself by modifying its local rotation transform.
// 
// This has no impact on the rotation/angle properties of the children, but it will update their worldTransform
// and on-screen orientation and position.
func (self *World) SetAngle(member float64) {
    self.Set("angle", member)
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
func (self *World) GetLeft() float64{
    return self.Get("left").Float()
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
func (self *World) SetLeft(member float64) {
    self.Set("left", member)
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
func (self *World) GetRight() float64{
    return self.Get("right").Float()
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
func (self *World) SetRight(member float64) {
    self.Set("right", member)
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
func (self *World) GetTop() float64{
    return self.Get("top").Float()
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
func (self *World) SetTop(member float64) {
    self.Set("top", member)
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
func (self *World) GetBottom() float64{
    return self.Get("bottom").Float()
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
func (self *World) SetBottom(member float64) {
    self.Set("bottom", member)
}

// The x coordinate of the group container.
// 
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *World) GetX() float64{
    return self.Get("x").Float()
}

// The x coordinate of the group container.
// 
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *World) SetX(member float64) {
    self.Set("x", member)
}

// The y coordinate of the group container.
// 
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *World) GetY() float64{
    return self.Get("y").Float()
}

// The y coordinate of the group container.
// 
// You can adjust the group container itself by modifying its coordinates.
// This will have no impact on the x/y coordinates of its children, but it will update their worldTransform and on-screen position.
func (self *World) SetY(member float64) {
    self.Set("y", member)
}

// The angle of rotation of the group container, in radians.
// 
// This will adjust the group container itself by modifying its rotation.
// This will have no impact on the rotation value of its children, but it will update their worldTransform and on-screen position.
func (self *World) GetRotation() float64{
    return self.Get("rotation").Float()
}

// The angle of rotation of the group container, in radians.
// 
// This will adjust the group container itself by modifying its rotation.
// This will have no impact on the rotation value of its children, but it will update their worldTransform and on-screen position.
func (self *World) SetRotation(member float64) {
    self.Set("rotation", member)
}

// The visible state of the group. Non-visible Groups and all of their children are not rendered.
func (self *World) GetVisible() bool{
    return self.Get("visible").Bool()
}

// The visible state of the group. Non-visible Groups and all of their children are not rendered.
func (self *World) SetVisible(member bool) {
    self.Set("visible", member)
}

// The alpha value of the group container.
func (self *World) GetAlpha() float64{
    return self.Get("alpha").Float()
}

// The alpha value of the group container.
func (self *World) SetAlpha(member float64) {
    self.Set("alpha", member)
}

// [read-only] The array of children of this container.
func (self *World) GetChildren() []DisplayObject{
	array := self.Get("children")
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// [read-only] The array of children of this container.
func (self *World) SetChildren(member []DisplayObject) {
    self.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *World) GetIgnoreChildInput() bool{
    return self.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *World) SetIgnoreChildInput(member bool) {
    self.Set("ignoreChildInput", member)
}



// Initialises the game world.
func (self *World) BootI(args ...interface{}) {
    self.Call("boot", args)
}

// Called whenever the State changes or resets.
// 
// It resets the world.x and world.y coordinates back to zero,
// then resets the Camera.
func (self *World) StateChangeI(args ...interface{}) {
    self.Call("stateChange", args)
}

// Updates the size of this world and sets World.x/y to the given values
// The Camera bounds and Physics bounds (if set) are also updated to match the new World bounds.
func (self *World) SetBoundsI(args ...interface{}) {
    self.Call("setBounds", args)
}

// Updates the size of this world. Note that this doesn't modify the world x/y coordinates, just the width and height.
func (self *World) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// Destroyer of worlds.
func (self *World) ShutdownI(args ...interface{}) {
    self.Call("shutdown", args)
}

// This will take the given game object and check if its x/y coordinates fall outside of the world bounds.
// If they do it will reposition the object to the opposite side of the world, creating a wrap-around effect.
// If sprite has a P2 body then the body (sprite.body) should be passed as first parameter to the function.
// 
// Please understand there are limitations to this method. For example if you have scaled the World
// then objects won't always be re-positioned correctly, and you'll need to employ your own wrapping function.
func (self *World) WrapI(args ...interface{}) {
    self.Call("wrap", args)
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
func (self *World) AddI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("add", args)}
}

// Adds an existing object to this group.
// 
// The child is added to the group at the location specified by the index value, this allows you to control child ordering.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) AddAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addAt", args)}
}

// Adds a child of this Group into the hash array.
// This call will return false if the child is not a child of this Group, or is already in the hash.
func (self *World) AddToHashI(args ...interface{}) bool{
    return self.Call("addToHash", args).Bool()
}

// Removes a child of this Group from the hash array.
// This call will return false if the child is not in the hash.
func (self *World) RemoveFromHashI(args ...interface{}) bool{
    return self.Call("removeFromHash", args).Bool()
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
func (self *World) AddMultipleI(args ...interface{}) interface{}{
    return self.Call("addMultiple", args)
}

// Returns the child found at the given index within this group.
func (self *World) GetAtI(args ...interface{}) interface{}{
    return self.Call("getAt", args)
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
func (self *World) CreateI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("create", args)}
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
func (self *World) CreateMultipleI(args ...interface{}) []interface{}{
	array := self.Call("createMultiple", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Internal method that re-applies all of the children's Z values.
// 
// This must be called whenever children ordering is altered so that their `z` indices are correctly updated.
func (self *World) UpdateZI(args ...interface{}) {
    self.Call("updateZ", args)
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
func (self *World) AlignI(args ...interface{}) {
    self.Call("align", args)
}

// Sets the group cursor to the first child in the group.
// 
// If the optional index parameter is given it sets the cursor to the object at that index instead.
func (self *World) ResetCursorI(args ...interface{}) interface{}{
    return self.Call("resetCursor", args)
}

// Advances the group cursor to the next (higher) object in the group.
// 
// If the cursor is at the end of the group (top child) it is moved the start of the group (bottom child).
func (self *World) NextI(args ...interface{}) interface{}{
    return self.Call("next", args)
}

// Moves the group cursor to the previous (lower) child in the group.
// 
// If the cursor is at the start of the group (bottom child) it is moved to the end (top child).
func (self *World) PreviousI(args ...interface{}) interface{}{
    return self.Call("previous", args)
}

// Swaps the position of two children in this group.
// 
// Both children must be in this group, a child cannot be swapped with itself, and unparented children cannot be swapped.
func (self *World) SwapI(args ...interface{}) {
    self.Call("swap", args)
}

// Brings the given child to the top of this group so it renders above all other children.
func (self *World) BringToTopI(args ...interface{}) interface{}{
    return self.Call("bringToTop", args)
}

// Sends the given child to the bottom of this group so it renders below all other children.
func (self *World) SendToBackI(args ...interface{}) interface{}{
    return self.Call("sendToBack", args)
}

// Moves the given child up one place in this group unless it's already at the top.
func (self *World) MoveUpI(args ...interface{}) interface{}{
    return self.Call("moveUp", args)
}

// Moves the given child down one place in this group unless it's already at the bottom.
func (self *World) MoveDownI(args ...interface{}) interface{}{
    return self.Call("moveDown", args)
}

// Positions the child found at the given index within this group to the given x and y coordinates.
func (self *World) XyI(args ...interface{}) {
    self.Call("xy", args)
}

// Reverses all children in this group.
// 
// This operation applies only to immediate children and does not propagate to subgroups.
func (self *World) ReverseI(args ...interface{}) {
    self.Call("reverse", args)
}

// Get the index position of the given child in this group, which should match the child's `z` property.
func (self *World) GetIndexI(args ...interface{}) int{
    return self.Call("getIndex", args).Int()
}

// Searches the Group for the first instance of a child with the `name`
// property matching the given argument. Should more than one child have
// the same name only the first instance is returned.
func (self *World) GetByNameI(args ...interface{}) interface{}{
    return self.Call("getByName", args)
}

// Replaces a child of this Group with the given newChild. The newChild cannot be a member of this Group.
// 
// If `Group.enableBody` is set, then a physics body will be created on the object, so long as one does not already exist.
// 
// If `Group.inputEnableChildren` is set, then an Input Handler will be created on the object, so long as one does not already exist.
func (self *World) ReplaceI(args ...interface{}) interface{}{
    return self.Call("replace", args)
}

// Checks if the child has the given property.
// 
// Will scan up to 4 levels deep only.
func (self *World) HasPropertyI(args ...interface{}) bool{
    return self.Call("hasProperty", args).Bool()
}

// Sets a property to the given value on the child. The operation parameter controls how the value is set.
// 
// The operations are:
// - 0: set the existing value to the given value; if force is `true` a new property will be created if needed
// - 1: will add the given value to the value already present.
// - 2: will subtract the given value from the value already present.
// - 3: will multiply the value already present by the given value.
// - 4: will divide the value already present by the given value.
func (self *World) SetPropertyI(args ...interface{}) bool{
    return self.Call("setProperty", args).Bool()
}

// Checks a property for the given value on the child.
func (self *World) CheckPropertyI(args ...interface{}) bool{
    return self.Call("checkProperty", args).Bool()
}

// Quickly set a property on a single child of this group to a new value.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetI(args ...interface{}) bool{
    return self.Call("set", args).Bool()
}

// Quickly set the same property across all children of this group to a new value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be set on the group but not its children.
// If you need that ability please see `Group.setAllChildren`.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAllI(args ...interface{}) {
    self.Call("setAll", args)
}

// Quickly set the same property across all children of this group, and any child Groups, to a new value.
// 
// If this group contains other Groups then the same property is set across their children as well, iterating down until it reaches the bottom.
// Unlike with `setAll` the property is NOT set on child Groups itself.
// 
// The operation parameter controls how the new value is assigned to the property, from simple replacement to addition and multiplication.
func (self *World) SetAllChildrenI(args ...interface{}) {
    self.Call("setAllChildren", args)
}

// Quickly check that the same property across all children of this group is equal to the given value.
// 
// This call doesn't descend down children, so if you have a Group inside of this group, the property will be checked on the group but not its children.
func (self *World) CheckAllI(args ...interface{}) {
    self.Call("checkAll", args)
}

// Adds the amount to the given property on all children in this group.
// 
// `Group.addAll('x', 10)` will add 10 to the child.x value for each child.
func (self *World) AddAllI(args ...interface{}) {
    self.Call("addAll", args)
}

// Subtracts the amount from the given property on all children in this group.
// 
// `Group.subAll('x', 10)` will minus 10 from the child.x value for each child.
func (self *World) SubAllI(args ...interface{}) {
    self.Call("subAll", args)
}

// Multiplies the given property by the amount on all children in this group.
// 
// `Group.multiplyAll('x', 2)` will x2 the child.x value for each child.
func (self *World) MultiplyAllI(args ...interface{}) {
    self.Call("multiplyAll", args)
}

// Divides the given property by the amount on all children in this group.
// 
// `Group.divideAll('x', 2)` will half the child.x value for each child.
func (self *World) DivideAllI(args ...interface{}) {
    self.Call("divideAll", args)
}

// Calls a function, specified by name, on all children in the group who exist (or do not exist).
// 
// After the existsValue parameter you can add as many parameters as you like, which will all be passed to the child callback.
func (self *World) CallAllExistsI(args ...interface{}) {
    self.Call("callAllExists", args)
}

// Returns a reference to a function that exists on a child of the group based on the given callback array.
func (self *World) CallbackFromArrayI(args ...interface{}) {
    self.Call("callbackFromArray", args)
}

// Calls a function, specified by name, on all on children.
// 
// The function is called for all children regardless if they are dead or alive (see callAllExists for different options).
// After the method parameter and context you can add as many extra parameters as you like, which will all be passed to the child.
func (self *World) CallAllI(args ...interface{}) {
    self.Call("callAll", args)
}

// The core preUpdate - as called by World.
func (self *World) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// The core update - as called by World.
func (self *World) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// The core postUpdate - as called by World.
func (self *World) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
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
func (self *World) FilterI(args ...interface{}) *ArraySet{
    return &ArraySet{self.Call("filter", args)}
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
func (self *World) ForEachI(args ...interface{}) {
    self.Call("forEach", args)
}

// Call a function on each existing child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachExistsI(args ...interface{}) {
    self.Call("forEachExists", args)
}

// Call a function on each alive child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachAliveI(args ...interface{}) {
    self.Call("forEachAlive", args)
}

// Call a function on each dead child in this group.
// 
// See {@link Phaser.Group#forEach forEach} for details.
func (self *World) ForEachDeadI(args ...interface{}) {
    self.Call("forEachDead", args)
}

// Sort the children in the group according to a particular key and ordering.
// 
// Call this function to sort the group according to a particular key value and order.
// 
// For example to depth sort Sprites for Zelda-style game you might call `group.sort('y', Phaser.Group.SORT_ASCENDING)` at the bottom of your `State.update()`.
// 
// Internally this uses a standard JavaScript Array sort, so everything that applies there also applies here, including
// alphabetical sorting, mixing strings and numbers, and Unicode sorting. See MDN for more details.
func (self *World) SortI(args ...interface{}) {
    self.Call("sort", args)
}

// Sort the children in the group according to custom sort function.
// 
// The `sortHandler` is provided the two parameters: the two children involved in the comparison (a and b).
// It should return -1 if `a > b`, 1 if `a < b` or 0 if `a === b`.
func (self *World) CustomSortI(args ...interface{}) {
    self.Call("customSort", args)
}

// An internal helper function for the sort process.
func (self *World) AscendingSortHandlerI(args ...interface{}) {
    self.Call("ascendingSortHandler", args)
}

// An internal helper function for the sort process.
func (self *World) DescendingSortHandlerI(args ...interface{}) {
    self.Call("descendingSortHandler", args)
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
func (self *World) IterateI(args ...interface{}) interface{}{
    return self.Call("iterate", args)
}

// Get the first display object that exists, or doesn't exist.
// 
// You can use the optional argument `createIfNull` to create a new Game Object if none matching your exists argument were found in this Group.
// 
// It works by calling `Group.create` passing it the parameters given to this method, and returning the new child.
// 
// If a child *was* found , `createIfNull` is `false` and you provided the additional arguments then the child
// will be reset and/or have a new texture loaded on it. This is handled by `Group.resetChild`.
func (self *World) GetFirstExistsI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getFirstExists", args)}
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
func (self *World) GetFirstAliveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getFirstAlive", args)}
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
func (self *World) GetFirstDeadI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getFirstDead", args)}
}

// Takes a child and if the `x` and `y` arguments are given it calls `child.reset(x, y)` on it.
// 
// If the `key` and optionally the `frame` arguments are given, it calls `child.loadTexture(key, frame)` on it.
// 
// The two operations are separate. For example if you just wish to load a new texture then pass `null` as the x and y values.
func (self *World) ResetChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("resetChild", args)}
}

// Return the child at the top of this group.
// 
// The top child is the child displayed (rendered) above every other child.
func (self *World) GetTopI(args ...interface{}) interface{}{
    return self.Call("getTop", args)
}

// Returns the child at the bottom of this group.
// 
// The bottom child the child being displayed (rendered) below every other child.
func (self *World) GetBottomI(args ...interface{}) interface{}{
    return self.Call("getBottom", args)
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
func (self *World) GetClosestToI(args ...interface{}) interface{}{
    return self.Call("getClosestTo", args)
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
func (self *World) GetFurthestFromI(args ...interface{}) interface{}{
    return self.Call("getFurthestFrom", args)
}

// Get the number of living children in this group.
func (self *World) CountLivingI(args ...interface{}) int{
    return self.Call("countLiving", args).Int()
}

// Get the number of dead children in this group.
func (self *World) CountDeadI(args ...interface{}) int{
    return self.Call("countDead", args).Int()
}

// Returns a random child from the group.
func (self *World) GetRandomI(args ...interface{}) interface{}{
    return self.Call("getRandom", args)
}

// Removes the given child from this group.
// 
// This will dispatch an `onRemovedFromGroup` event from the child (if it has one), and optionally destroy the child.
// 
// If the group cursor was referring to the removed child it is updated to refer to the next child.
func (self *World) RemoveI(args ...interface{}) bool{
    return self.Call("remove", args).Bool()
}

// Moves all children from this Group to the Group given.
func (self *World) MoveAllI(args ...interface{}) *Group{
    return &Group{self.Call("moveAll", args)}
}

// Removes all children from this Group, but does not remove the group from its parent.
// 
// The children can be optionally destroyed as they are removed.
// 
// You can also optionally also destroy the BaseTexture the Child is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *World) RemoveAllI(args ...interface{}) {
    self.Call("removeAll", args)
}

// Removes all children from this group whose index falls beteen the given startIndex and endIndex values.
func (self *World) RemoveBetweenI(args ...interface{}) {
    self.Call("removeBetween", args)
}

// Destroys this group.
// 
// Removes all children, then removes this group from its parent and nulls references.
func (self *World) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
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
func (self *World) AlignInI(args ...interface{}) *Group{
    return &Group{self.Call("alignIn", args)}
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
func (self *World) AlignToI(args ...interface{}) *Group{
    return &Group{self.Call("alignTo", args)}
}

// Adds a child to the container.
func (self *World) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *World) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *World) SwapChildrenI(args ...interface{}) {
    self.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *World) GetChildIndexI(args ...interface{}) float64{
    return self.Call("getChildIndex", args).Float()
}

// Changes the position of an existing child in the display object container
func (self *World) SetChildIndexI(args ...interface{}) {
    self.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *World) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *World) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *World) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *World) RemoveChildrenI(args ...interface{}) {
    self.Call("removeChildren", args)
}

// Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *World) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getBounds", args)}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *World) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *World) SetStageReferenceI(args ...interface{}) {
    self.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *World) RemoveStageReferenceI(args ...interface{}) {
    self.Call("removeStageReference", args)
}

// Renders the object using the WebGL renderer
func (self *World) _renderWebGLI(args ...interface{}) {
    self.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *World) _renderCanvasI(args ...interface{}) {
    self.Call("_renderCanvas", args)
}
