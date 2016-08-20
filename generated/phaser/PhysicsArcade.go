// Automatic generation for Phaser.Physics.Arcade
// generated file PhysicsArcade.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Arcade Physics world. Contains Arcade Physics related collision, overlap and motion methods.
type PhysicsArcade struct {
    *js.Object
}


// Local reference to game.
func (self *PhysicsArcade) GetGame() Game{
    return Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsArcade) SetGame(member Game) {
    self.Set("game", member)
}

// The World gravity setting. Defaults to x: 0, y: 0, or no gravity.
func (self *PhysicsArcade) GetGravity() Point{
    return Point{self.Get("gravity")}
}

// The World gravity setting. Defaults to x: 0, y: 0, or no gravity.
func (self *PhysicsArcade) SetGravity(member Point) {
    self.Set("gravity", member)
}

// The bounds inside of which the physics world exists. Defaults to match the world bounds.
func (self *PhysicsArcade) GetBounds() Rectangle{
    return Rectangle{self.Get("bounds")}
}

// The bounds inside of which the physics world exists. Defaults to match the world bounds.
func (self *PhysicsArcade) SetBounds(member Rectangle) {
    self.Set("bounds", member)
}

// Set the checkCollision properties to control for which bounds collision is processed.
// For example checkCollision.down = false means Bodies cannot collide with the World.bounds.bottom. An object containing allowed collision flags.
func (self *PhysicsArcade) GetCheckCollision() interface{}{
    return self.Get("checkCollision")
}

// Set the checkCollision properties to control for which bounds collision is processed.
// For example checkCollision.down = false means Bodies cannot collide with the World.bounds.bottom. An object containing allowed collision flags.
func (self *PhysicsArcade) SetCheckCollision(member interface{}) {
    self.Set("checkCollision", member)
}

// Used by the QuadTree to set the maximum number of objects per quad.
func (self *PhysicsArcade) GetMaxObjects() float64{
    return self.Get("maxObjects").Float()
}

// Used by the QuadTree to set the maximum number of objects per quad.
func (self *PhysicsArcade) SetMaxObjects(member float64) {
    self.Set("maxObjects", member)
}

// Used by the QuadTree to set the maximum number of iteration levels.
func (self *PhysicsArcade) GetMaxLevels() float64{
    return self.Get("maxLevels").Float()
}

// Used by the QuadTree to set the maximum number of iteration levels.
func (self *PhysicsArcade) SetMaxLevels(member float64) {
    self.Set("maxLevels", member)
}

// A value added to the delta values during collision checks.
func (self *PhysicsArcade) GetOVERLAP_BIAS() float64{
    return self.Get("OVERLAP_BIAS").Float()
}

// A value added to the delta values during collision checks.
func (self *PhysicsArcade) SetOVERLAP_BIAS(member float64) {
    self.Set("OVERLAP_BIAS", member)
}

// If true World.separate will always separate on the X axis before Y. Otherwise it will check gravity totals first.
func (self *PhysicsArcade) GetForceX() bool{
    return self.Get("forceX").Bool()
}

// If true World.separate will always separate on the X axis before Y. Otherwise it will check gravity totals first.
func (self *PhysicsArcade) SetForceX(member bool) {
    self.Set("forceX", member)
}

// Used when colliding a Sprite vs. a Group, or a Group vs. a Group, this defines the direction the sort is based on. Default is Phaser.Physics.Arcade.LEFT_RIGHT.
func (self *PhysicsArcade) GetSortDirection() float64{
    return self.Get("sortDirection").Float()
}

// Used when colliding a Sprite vs. a Group, or a Group vs. a Group, this defines the direction the sort is based on. Default is Phaser.Physics.Arcade.LEFT_RIGHT.
func (self *PhysicsArcade) SetSortDirection(member float64) {
    self.Set("sortDirection", member)
}

// If true the QuadTree will not be used for any collision. QuadTrees are great if objects are well spread out in your game, otherwise they are a performance hit. If you enable this you can disable on a per body basis via `Body.skipQuadTree`.
func (self *PhysicsArcade) GetSkipQuadTree() bool{
    return self.Get("skipQuadTree").Bool()
}

// If true the QuadTree will not be used for any collision. QuadTrees are great if objects are well spread out in your game, otherwise they are a performance hit. If you enable this you can disable on a per body basis via `Body.skipQuadTree`.
func (self *PhysicsArcade) SetSkipQuadTree(member bool) {
    self.Set("skipQuadTree", member)
}

// If `true` the `Body.preUpdate` method will be skipped, halting all motion for all bodies. Note that other methods such as `collide` will still work, so be careful not to call them on paused bodies.
func (self *PhysicsArcade) GetIsPaused() bool{
    return self.Get("isPaused").Bool()
}

// If `true` the `Body.preUpdate` method will be skipped, halting all motion for all bodies. Note that other methods such as `collide` will still work, so be careful not to call them on paused bodies.
func (self *PhysicsArcade) SetIsPaused(member bool) {
    self.Set("isPaused", member)
}

// The world QuadTree.
func (self *PhysicsArcade) GetQuadTree() QuadTree{
    return QuadTree{self.Get("quadTree")}
}

// The world QuadTree.
func (self *PhysicsArcade) SetQuadTree(member QuadTree) {
    self.Set("quadTree", member)
}

// A constant used for the sortDirection value.
// Use this if you don't wish to perform any pre-collision sorting at all, or will manually sort your Groups.
func (self *PhysicsArcade) GetSORT_NONE() float64{
    return self.Get("SORT_NONE").Float()
}

// A constant used for the sortDirection value.
// Use this if you don't wish to perform any pre-collision sorting at all, or will manually sort your Groups.
func (self *PhysicsArcade) SetSORT_NONE(member float64) {
    self.Set("SORT_NONE", member)
}

// A constant used for the sortDirection value.
// Use this if your game world is wide but short and scrolls from the left to the right (i.e. Mario)
func (self *PhysicsArcade) GetLEFT_RIGHT() float64{
    return self.Get("LEFT_RIGHT").Float()
}

// A constant used for the sortDirection value.
// Use this if your game world is wide but short and scrolls from the left to the right (i.e. Mario)
func (self *PhysicsArcade) SetLEFT_RIGHT(member float64) {
    self.Set("LEFT_RIGHT", member)
}

// A constant used for the sortDirection value.
// Use this if your game world is wide but short and scrolls from the right to the left (i.e. Mario backwards)
func (self *PhysicsArcade) GetRIGHT_LEFT() float64{
    return self.Get("RIGHT_LEFT").Float()
}

// A constant used for the sortDirection value.
// Use this if your game world is wide but short and scrolls from the right to the left (i.e. Mario backwards)
func (self *PhysicsArcade) SetRIGHT_LEFT(member float64) {
    self.Set("RIGHT_LEFT", member)
}

// A constant used for the sortDirection value.
// Use this if your game world is narrow but tall and scrolls from the top to the bottom (i.e. Dig Dug)
func (self *PhysicsArcade) GetTOP_BOTTOM() float64{
    return self.Get("TOP_BOTTOM").Float()
}

// A constant used for the sortDirection value.
// Use this if your game world is narrow but tall and scrolls from the top to the bottom (i.e. Dig Dug)
func (self *PhysicsArcade) SetTOP_BOTTOM(member float64) {
    self.Set("TOP_BOTTOM", member)
}

// A constant used for the sortDirection value.
// Use this if your game world is narrow but tall and scrolls from the bottom to the top (i.e. Commando or a vertically scrolling shoot-em-up)
func (self *PhysicsArcade) GetBOTTOM_TOP() float64{
    return self.Get("BOTTOM_TOP").Float()
}

// A constant used for the sortDirection value.
// Use this if your game world is narrow but tall and scrolls from the bottom to the top (i.e. Commando or a vertically scrolling shoot-em-up)
func (self *PhysicsArcade) SetBOTTOM_TOP(member float64) {
    self.Set("BOTTOM_TOP", member)
}



// Updates the size of this physics world.
func (self *PhysicsArcade) SetBoundsI(args ...interface{}) {
    self.Call("setBounds", args)
}

// Updates the size of this physics world to match the size of the game world.
func (self *PhysicsArcade) SetBoundsToWorldI(args ...interface{}) {
    self.Call("setBoundsToWorld", args)
}

// This will create an Arcade Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsArcade) EnableI(args ...interface{}) {
    self.Call("enable", args)
}

// Creates an Arcade Physics body on the given game object.
// 
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
// 
// When you add an Arcade Physics body to an object it will automatically add the object into its parent Groups hash array.
func (self *PhysicsArcade) EnableBodyI(args ...interface{}) {
    self.Call("enableBody", args)
}

// Called automatically by a Physics body, it updates all motion related values on the Body unless `World.isPaused` is `true`.
func (self *PhysicsArcade) UpdateMotionI(args ...interface{}) {
    self.Call("updateMotion", args)
}

// A tween-like function that takes a starting velocity and some other factors and returns an altered velocity.
// Based on a function in Flixel by @ADAMATOMIC
func (self *PhysicsArcade) ComputeVelocityI(args ...interface{}) float64{
    return self.Call("computeVelocity", args).Float()
}

// Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups within Groups).
func (self *PhysicsArcade) OverlapI(args ...interface{}) bool{
    return self.Call("overlap", args).Bool()
}

// Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// Both the first and second parameter can be arrays of objects, of differing types.
// If two arrays are passed, the contents of the first parameter will be tested against all contents of the 2nd parameter.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
// NOTE: This function is not recursive, and will not test against children of objects passed (i.e. Groups or Tilemaps within other Groups).
func (self *PhysicsArcade) CollideI(args ...interface{}) bool{
    return self.Call("collide", args).Bool()
}

// A Sort function for sorting two bodies based on a LEFT to RIGHT sort direction.
// 
// This is called automatically by World.sort
func (self *PhysicsArcade) SortLeftRightI(args ...interface{}) int{
    return self.Call("sortLeftRight", args).Int()
}

// A Sort function for sorting two bodies based on a RIGHT to LEFT sort direction.
// 
// This is called automatically by World.sort
func (self *PhysicsArcade) SortRightLeftI(args ...interface{}) int{
    return self.Call("sortRightLeft", args).Int()
}

// A Sort function for sorting two bodies based on a TOP to BOTTOM sort direction.
// 
// This is called automatically by World.sort
func (self *PhysicsArcade) SortTopBottomI(args ...interface{}) int{
    return self.Call("sortTopBottom", args).Int()
}

// A Sort function for sorting two bodies based on a BOTTOM to TOP sort direction.
// 
// This is called automatically by World.sort
func (self *PhysicsArcade) SortBottomTopI(args ...interface{}) int{
    return self.Call("sortBottomTop", args).Int()
}

// This method will sort a Groups hash array.
// 
// If the Group has `physicsSortDirection` set it will use the sort direction defined.
// 
// Otherwise if the sortDirection parameter is undefined, or Group.physicsSortDirection is null, it will use Phaser.Physics.Arcade.sortDirection.
// 
// By changing Group.physicsSortDirection you can customise each Group to sort in a different order.
func (self *PhysicsArcade) SortI(args ...interface{}) {
    self.Call("sort", args)
}

// Internal collision handler.
func (self *PhysicsArcade) CollideHandlerI(args ...interface{}) {
    self.Call("collideHandler", args)
}

// An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideSpriteVsSpriteI(args ...interface{}) bool{
    return self.Call("collideSpriteVsSprite", args).Bool()
}

// An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideSpriteVsGroupI(args ...interface{}) {
    self.Call("collideSpriteVsGroup", args)
}

// An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideGroupVsSelfI(args ...interface{}) bool{
    return self.Call("collideGroupVsSelf", args).Bool()
}

// An internal function. Use Phaser.Physics.Arcade.collide instead.
func (self *PhysicsArcade) CollideGroupVsGroupI(args ...interface{}) {
    self.Call("collideGroupVsGroup", args)
}

// The core separation function to separate two physics bodies.
func (self *PhysicsArcade) SeparateI(args ...interface{}) bool{
    return self.Call("separate", args).Bool()
}

// Check for intersection against two bodies.
func (self *PhysicsArcade) IntersectsI(args ...interface{}) bool{
    return self.Call("intersects", args).Bool()
}

// Checks to see if a circular Body intersects with a Rectangular Body.
func (self *PhysicsArcade) CircleBodyIntersectsI(args ...interface{}) bool{
    return self.Call("circleBodyIntersects", args).Bool()
}

// The core separation function to separate two circular physics bodies.
func (self *PhysicsArcade) SeparateCircleI(args ...interface{}) bool{
    return self.Call("separateCircle", args).Bool()
}

// Calculates the horizontal overlap between two Bodies and sets their properties accordingly, including:
// `touching.left`, `touching.right` and `overlapX`.
func (self *PhysicsArcade) GetOverlapXI(args ...interface{}) float64{
    return self.Call("getOverlapX", args).Float()
}

// Calculates the vertical overlap between two Bodies and sets their properties accordingly, including:
// `touching.up`, `touching.down` and `overlapY`.
func (self *PhysicsArcade) GetOverlapYI(args ...interface{}) float64{
    return self.Call("getOverlapY", args).Float()
}

// The core separation function to separate two physics bodies on the x axis.
func (self *PhysicsArcade) SeparateXI(args ...interface{}) bool{
    return self.Call("separateX", args).Bool()
}

// The core separation function to separate two physics bodies on the y axis.
func (self *PhysicsArcade) SeparateYI(args ...interface{}) bool{
    return self.Call("separateY", args).Bool()
}

// Given a Group and a Pointer this will check to see which Group children overlap with the Pointer coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the Pointer or not.
func (self *PhysicsArcade) GetObjectsUnderPointerI(args ...interface{}) []DisplayObject{
	array := self.Call("getObjectsUnderPointer", args)
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		
			out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// Given a Group and a location this will check to see which Group children overlap with the coordinates.
// Each child will be sent to the given callback for further processing.
// Note that the children are not checked for depth order, but simply if they overlap the coordinate or not.
func (self *PhysicsArcade) GetObjectsAtLocationI(args ...interface{}) []DisplayObject{
	array := self.Call("getObjectsAtLocation", args)
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		
			out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// Move the given display object towards the destination object at a steady velocity.
// If you specify a maxTime then it will adjust the speed (overwriting what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
// Note: Doesn't take into account acceleration, maxVelocity or drag (if you've set drag or acceleration too high this object may not move at all)
func (self *PhysicsArcade) MoveToObjectI(args ...interface{}) float64{
    return self.Call("moveToObject", args).Float()
}

// Move the given display object towards the pointer at a steady velocity. If no pointer is given it will use Phaser.Input.activePointer.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) MoveToPointerI(args ...interface{}) float64{
    return self.Call("moveToPointer", args).Float()
}

// Move the given display object towards the x/y coordinates at a steady velocity.
// If you specify a maxTime then it will adjust the speed (over-writing what you set) so it arrives at the destination in that number of seconds.
// Timings are approximate due to the way browser timers work. Allow for a variance of +- 50ms.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
// Note: Doesn't take into account acceleration, maxVelocity or drag (if you've set drag or acceleration too high this object may not move at all)
func (self *PhysicsArcade) MoveToXYI(args ...interface{}) float64{
    return self.Call("moveToXY", args).Float()
}

// Given the angle (in degrees) and speed calculate the velocity and return it as a Point object, or set it to the given point object.
// One way to use this is: velocityFromAngle(angle, 200, sprite.velocity) which will set the values directly to the sprites velocity and not create a new Point object.
func (self *PhysicsArcade) VelocityFromAngleI(args ...interface{}) Point{
    return Point{self.Call("velocityFromAngle", args)}
}

// Given the rotation (in radians) and speed calculate the velocity and return it as a Point object, or set it to the given point object.
// One way to use this is: velocityFromRotation(rotation, 200, sprite.velocity) which will set the values directly to the sprites velocity and not create a new Point object.
func (self *PhysicsArcade) VelocityFromRotationI(args ...interface{}) Point{
    return Point{self.Call("velocityFromRotation", args)}
}

// Given the rotation (in radians) and speed calculate the acceleration and return it as a Point object, or set it to the given point object.
// One way to use this is: accelerationFromRotation(rotation, 200, sprite.acceleration) which will set the values directly to the sprites acceleration and not create a new Point object.
func (self *PhysicsArcade) AccelerationFromRotationI(args ...interface{}) Point{
    return Point{self.Call("accelerationFromRotation", args)}
}

// Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToObjectI(args ...interface{}) float64{
    return self.Call("accelerateToObject", args).Float()
}

// Sets the acceleration.x/y property on the display object so it will move towards the target at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToPointerI(args ...interface{}) float64{
    return self.Call("accelerateToPointer", args).Float()
}

// Sets the acceleration.x/y property on the display object so it will move towards the x/y coordinates at the given speed (in pixels per second sq.)
// You must give a maximum speed value, beyond which the display object won't go any faster.
// Note: The display object does not continuously track the target. If the target changes location during transit the display object will not modify its course.
// Note: The display object doesn't stop moving once it reaches the destination coordinates.
func (self *PhysicsArcade) AccelerateToXYI(args ...interface{}) float64{
    return self.Call("accelerateToXY", args).Float()
}

// Find the distance between two display objects (like Sprites).
// 
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceBetweenI(args ...interface{}) float64{
    return self.Call("distanceBetween", args).Float()
}

// Find the distance between a display object (like a Sprite) and the given x/y coordinates.
// The calculation is made from the display objects x/y coordinate. This may be the top-left if its anchor hasn't been changed.
// If you need to calculate from the center of a display object instead use the method distanceBetweenCenters()
// 
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceToXYI(args ...interface{}) float64{
    return self.Call("distanceToXY", args).Float()
}

// Find the distance between a display object (like a Sprite) and a Pointer. If no Pointer is given the Input.activePointer is used.
// The calculation is made from the display objects x/y coordinate. This may be the top-left if its anchor hasn't been changed.
// If you need to calculate from the center of a display object instead use the method distanceBetweenCenters()
// 
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) DistanceToPointerI(args ...interface{}) float64{
    return self.Call("distanceToPointer", args).Float()
}

// Find the angle in radians between two display objects (like Sprites).
// 
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleBetweenI(args ...interface{}) float64{
    return self.Call("angleBetween", args).Float()
}

// Find the angle in radians between centers of two display objects (like Sprites).
func (self *PhysicsArcade) AngleBetweenCentersI(args ...interface{}) float64{
    return self.Call("angleBetweenCenters", args).Float()
}

// Find the angle in radians between a display object (like a Sprite) and the given x/y coordinate.
// 
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleToXYI(args ...interface{}) float64{
    return self.Call("angleToXY", args).Float()
}

// Find the angle in radians between a display object (like a Sprite) and a Pointer, taking their x/y and center into account.
// 
// The optional `world` argument allows you to return the result based on the Game Objects `world` property,
// instead of its `x` and `y` values. This is useful of the object has been nested inside an offset Group,
// or parent Game Object.
func (self *PhysicsArcade) AngleToPointerI(args ...interface{}) float64{
    return self.Call("angleToPointer", args).Float()
}

// Find the angle in radians between a display object (like a Sprite) and a Pointer, 
// taking their x/y and center into account relative to the world.
func (self *PhysicsArcade) WorldAngleToPointerI(args ...interface{}) float64{
    return self.Call("worldAngleToPointer", args).Float()
}
