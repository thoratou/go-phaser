// Automatic generation for Phaser.Physics.Ninja
// generated file PhysicsNinja.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Ninja Physics. The Ninja Physics system was created in Flash by Metanet Software and ported to JavaScript by Richard Davey.
// 
// It allows for AABB and Circle to Tile collision. Tiles can be any of 34 different types, including slopes, convex and concave shapes.
// 
// It does what it does very well, but is ripe for expansion and optimisation. Here are some features that I'd love to see the community add:
// 
// * AABB to AABB collision
// * AABB to Circle collision
// * AABB and Circle 'immovable' property support
// * n-way collision, so an AABB/Circle could pass through a tile from below and land upon it.
// * QuadTree or spatial grid for faster Body vs. Tile Group look-ups.
// * Optimise the internal vector math and reduce the quantity of temporary vars created.
// * Expand Gravity and Bounce to allow for separate x/y axis values.
// * Support Bodies linked to Sprites that don't have anchor set to 0.5
// 
// Feel free to attempt any of the above and submit a Pull Request with your code! Be sure to include test cases proving they work.
type PhysicsNinja struct {
    *js.Object
}


// Ninja Physics. The Ninja Physics system was created in Flash by Metanet Software and ported to JavaScript by Richard Davey.
// 
// It allows for AABB and Circle to Tile collision. Tiles can be any of 34 different types, including slopes, convex and concave shapes.
// 
// It does what it does very well, but is ripe for expansion and optimisation. Here are some features that I'd love to see the community add:
// 
// * AABB to AABB collision
// * AABB to Circle collision
// * AABB and Circle 'immovable' property support
// * n-way collision, so an AABB/Circle could pass through a tile from below and land upon it.
// * QuadTree or spatial grid for faster Body vs. Tile Group look-ups.
// * Optimise the internal vector math and reduce the quantity of temporary vars created.
// * Expand Gravity and Bounce to allow for separate x/y axis values.
// * Support Bodies linked to Sprites that don't have anchor set to 0.5
// 
// Feel free to attempt any of the above and submit a Pull Request with your code! Be sure to include test cases proving they work.
func NewPhysicsNinja(game *Game) *PhysicsNinja {
    return &PhysicsNinja{js.Global.Get("Phaser").Get("Physics").Get("Ninja").New(game)}
}

// Ninja Physics. The Ninja Physics system was created in Flash by Metanet Software and ported to JavaScript by Richard Davey.
// 
// It allows for AABB and Circle to Tile collision. Tiles can be any of 34 different types, including slopes, convex and concave shapes.
// 
// It does what it does very well, but is ripe for expansion and optimisation. Here are some features that I'd love to see the community add:
// 
// * AABB to AABB collision
// * AABB to Circle collision
// * AABB and Circle 'immovable' property support
// * n-way collision, so an AABB/Circle could pass through a tile from below and land upon it.
// * QuadTree or spatial grid for faster Body vs. Tile Group look-ups.
// * Optimise the internal vector math and reduce the quantity of temporary vars created.
// * Expand Gravity and Bounce to allow for separate x/y axis values.
// * Support Bodies linked to Sprites that don't have anchor set to 0.5
// 
// Feel free to attempt any of the above and submit a Pull Request with your code! Be sure to include test cases proving they work.
func NewPhysicsNinjaI(args ...interface{}) *PhysicsNinja {
    return &PhysicsNinja{js.Global.Get("Phaser").Get("Physics").Get("Ninja").New(args)}
}



// Local reference to game.
func (self *PhysicsNinja) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *PhysicsNinja) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Local reference to game.time.
func (self *PhysicsNinja) Time() *Time{
    return &Time{self.Object.Get("time")}
}

// Local reference to game.time.
func (self *PhysicsNinja) SetTimeA(member *Time) {
    self.Object.Set("time", member)
}

// The World gravity setting.
func (self *PhysicsNinja) Gravity() int{
    return self.Object.Get("gravity").Int()
}

// The World gravity setting.
func (self *PhysicsNinja) SetGravityA(member int) {
    self.Object.Set("gravity", member)
}

// The bounds inside of which the physics world exists. Defaults to match the world bounds.
func (self *PhysicsNinja) Bounds() *Rectangle{
    return &Rectangle{self.Object.Get("bounds")}
}

// The bounds inside of which the physics world exists. Defaults to match the world bounds.
func (self *PhysicsNinja) SetBoundsA(member *Rectangle) {
    self.Object.Set("bounds", member)
}

// Used by the QuadTree to set the maximum number of objects per quad.
func (self *PhysicsNinja) MaxObjects() int{
    return self.Object.Get("maxObjects").Int()
}

// Used by the QuadTree to set the maximum number of objects per quad.
func (self *PhysicsNinja) SetMaxObjectsA(member int) {
    self.Object.Set("maxObjects", member)
}

// Used by the QuadTree to set the maximum number of iteration levels.
func (self *PhysicsNinja) MaxLevels() int{
    return self.Object.Get("maxLevels").Int()
}

// Used by the QuadTree to set the maximum number of iteration levels.
func (self *PhysicsNinja) SetMaxLevelsA(member int) {
    self.Object.Set("maxLevels", member)
}

// The world QuadTree.
func (self *PhysicsNinja) QuadTree() *QuadTree{
    return &QuadTree{self.Object.Get("quadTree")}
}

// The world QuadTree.
func (self *PhysicsNinja) SetQuadTreeA(member *QuadTree) {
    self.Object.Set("quadTree", member)
}



// This will create a Ninja Physics AABB body on the given game object. Its dimensions will match the width and height of the object at the point it is created.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableAABB(object interface{}) {
    self.Object.Call("enableAABB", object)
}

// This will create a Ninja Physics AABB body on the given game object. Its dimensions will match the width and height of the object at the point it is created.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableAABB1O(object interface{}, children bool) {
    self.Object.Call("enableAABB", object, children)
}

// This will create a Ninja Physics AABB body on the given game object. Its dimensions will match the width and height of the object at the point it is created.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableAABBI(args ...interface{}) {
    self.Object.Call("enableAABB", args)
}

// This will create a Ninja Physics Circle body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableCircle(object interface{}, radius int) {
    self.Object.Call("enableCircle", object, radius)
}

// This will create a Ninja Physics Circle body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableCircle1O(object interface{}, radius int, children bool) {
    self.Object.Call("enableCircle", object, radius, children)
}

// This will create a Ninja Physics Circle body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableCircleI(args ...interface{}) {
    self.Object.Call("enableCircle", args)
}

// This will create a Ninja Physics Tile body on the given game object. There are 34 different types of tile you can create, including 45 degree slopes,
// convex and concave circles and more. The id parameter controls which Tile type is created, but you can also change it at run-time.
// Note that for all degree based tile types they need to have an equal width and height. If the given object doesn't have equal width and height it will use the width.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableTile(object interface{}) {
    self.Object.Call("enableTile", object)
}

// This will create a Ninja Physics Tile body on the given game object. There are 34 different types of tile you can create, including 45 degree slopes,
// convex and concave circles and more. The id parameter controls which Tile type is created, but you can also change it at run-time.
// Note that for all degree based tile types they need to have an equal width and height. If the given object doesn't have equal width and height it will use the width.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableTile1O(object interface{}, id int) {
    self.Object.Call("enableTile", object, id)
}

// This will create a Ninja Physics Tile body on the given game object. There are 34 different types of tile you can create, including 45 degree slopes,
// convex and concave circles and more. The id parameter controls which Tile type is created, but you can also change it at run-time.
// Note that for all degree based tile types they need to have an equal width and height. If the given object doesn't have equal width and height it will use the width.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableTile2O(object interface{}, id int, children bool) {
    self.Object.Call("enableTile", object, id, children)
}

// This will create a Ninja Physics Tile body on the given game object. There are 34 different types of tile you can create, including 45 degree slopes,
// convex and concave circles and more. The id parameter controls which Tile type is created, but you can also change it at run-time.
// Note that for all degree based tile types they need to have an equal width and height. If the given object doesn't have equal width and height it will use the width.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableTileI(args ...interface{}) {
    self.Object.Call("enableTile", args)
}

// This will create a Ninja Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) Enable(object interface{}) {
    self.Object.Call("enable", object)
}

// This will create a Ninja Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) Enable1O(object interface{}, type_ int) {
    self.Object.Call("enable", object, type_)
}

// This will create a Ninja Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) Enable2O(object interface{}, type_ int, id int) {
    self.Object.Call("enable", object, type_, id)
}

// This will create a Ninja Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) Enable3O(object interface{}, type_ int, id int, radius int) {
    self.Object.Call("enable", object, type_, id, radius)
}

// This will create a Ninja Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) Enable4O(object interface{}, type_ int, id int, radius int, children bool) {
    self.Object.Call("enable", object, type_, id, radius, children)
}

// This will create a Ninja Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableI(args ...interface{}) {
    self.Object.Call("enable", args)
}

// Creates a Ninja Physics body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
func (self *PhysicsNinja) EnableBody(object interface{}) {
    self.Object.Call("enableBody", object)
}

// Creates a Ninja Physics body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
func (self *PhysicsNinja) EnableBodyI(args ...interface{}) {
    self.Object.Call("enableBody", args)
}

// Updates the size of this physics world.
func (self *PhysicsNinja) SetBounds(x int, y int, width int, height int) {
    self.Object.Call("setBounds", x, y, width, height)
}

// Updates the size of this physics world.
func (self *PhysicsNinja) SetBoundsI(args ...interface{}) {
    self.Object.Call("setBounds", args)
}

// Updates the size of this physics world to match the size of the game world.
func (self *PhysicsNinja) SetBoundsToWorld() {
    self.Object.Call("setBoundsToWorld")
}

// Updates the size of this physics world to match the size of the game world.
func (self *PhysicsNinja) SetBoundsToWorldI(args ...interface{}) {
    self.Object.Call("setBoundsToWorld", args)
}

// Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsNinja) ClearTilemapLayerBodies(map_ *Tilemap) {
    self.Object.Call("clearTilemapLayerBodies", map_)
}

// Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsNinja) ClearTilemapLayerBodies1O(map_ *Tilemap, layer interface{}) {
    self.Object.Call("clearTilemapLayerBodies", map_, layer)
}

// Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsNinja) ClearTilemapLayerBodiesI(args ...interface{}) {
    self.Object.Call("clearTilemapLayerBodies", args)
}

// Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics tiles.
// Only call this *after* you have specified all of the tiles you wish to collide with calls like Tilemap.setCollisionBetween, etc.
// Every time you call this method it will destroy any previously created bodies and remove them from the world.
// Therefore understand it's a very expensive operation and not to be done in a core game update loop.
// 
// In Ninja the Tiles have an ID from 0 to 33, where 0 is 'empty', 1 is a full tile, 2 is a 45-degree slope, etc. You can find the ID
// list either at the very bottom of `Tile.js`, or in a handy visual reference in the `resources/Ninja Physics Debug Tiles` folder in the repository.
// The slopeMap parameter is an array that controls how the indexes of the tiles in your tilemap data will map to the Ninja Tile IDs.
// For example if you had 6 tiles in your tileset: Imagine the first 4 should be converted into fully solid Tiles and the other 2 are 45-degree slopes.
// Your slopeMap array would look like this: `[ 1, 1, 1, 1, 2, 3 ]`.
// Where each element of the array is a tile in your tilemap and the resulting Ninja Tile it should create.
func (self *PhysicsNinja) ConvertTilemap(map_ *Tilemap, layer interface{}, slopeMap interface{}) []interface{}{
	array00 := self.Object.Call("convertTilemap", map_, layer, slopeMap)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Goes through all tiles in the given Tilemap and TilemapLayer and converts those set to collide into physics tiles.
// Only call this *after* you have specified all of the tiles you wish to collide with calls like Tilemap.setCollisionBetween, etc.
// Every time you call this method it will destroy any previously created bodies and remove them from the world.
// Therefore understand it's a very expensive operation and not to be done in a core game update loop.
// 
// In Ninja the Tiles have an ID from 0 to 33, where 0 is 'empty', 1 is a full tile, 2 is a 45-degree slope, etc. You can find the ID
// list either at the very bottom of `Tile.js`, or in a handy visual reference in the `resources/Ninja Physics Debug Tiles` folder in the repository.
// The slopeMap parameter is an array that controls how the indexes of the tiles in your tilemap data will map to the Ninja Tile IDs.
// For example if you had 6 tiles in your tileset: Imagine the first 4 should be converted into fully solid Tiles and the other 2 are 45-degree slopes.
// Your slopeMap array would look like this: `[ 1, 1, 1, 1, 2, 3 ]`.
// Where each element of the array is a tile in your tilemap and the resulting Ninja Tile it should create.
func (self *PhysicsNinja) ConvertTilemapI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("convertTilemap", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// The second parameter can be an array of objects, of differing types.
func (self *PhysicsNinja) Overlap(object1 interface{}, object2 interface{}) bool{
    return self.Object.Call("overlap", object1, object2).Bool()
}

// Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// The second parameter can be an array of objects, of differing types.
func (self *PhysicsNinja) Overlap1O(object1 interface{}, object2 interface{}, overlapCallback interface{}) bool{
    return self.Object.Call("overlap", object1, object2, overlapCallback).Bool()
}

// Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// The second parameter can be an array of objects, of differing types.
func (self *PhysicsNinja) Overlap2O(object1 interface{}, object2 interface{}, overlapCallback interface{}, processCallback interface{}) bool{
    return self.Object.Call("overlap", object1, object2, overlapCallback, processCallback).Bool()
}

// Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// The second parameter can be an array of objects, of differing types.
func (self *PhysicsNinja) Overlap3O(object1 interface{}, object2 interface{}, overlapCallback interface{}, processCallback interface{}, callbackContext interface{}) bool{
    return self.Object.Call("overlap", object1, object2, overlapCallback, processCallback, callbackContext).Bool()
}

// Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// The second parameter can be an array of objects, of differing types.
func (self *PhysicsNinja) OverlapI(args ...interface{}) bool{
    return self.Object.Call("overlap", args).Bool()
}

// Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// The second parameter can be an array of objects, of differing types.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
func (self *PhysicsNinja) Collide(object1 interface{}, object2 interface{}) bool{
    return self.Object.Call("collide", object1, object2).Bool()
}

// Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// The second parameter can be an array of objects, of differing types.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
func (self *PhysicsNinja) Collide1O(object1 interface{}, object2 interface{}, collideCallback interface{}) bool{
    return self.Object.Call("collide", object1, object2, collideCallback).Bool()
}

// Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// The second parameter can be an array of objects, of differing types.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
func (self *PhysicsNinja) Collide2O(object1 interface{}, object2 interface{}, collideCallback interface{}, processCallback interface{}) bool{
    return self.Object.Call("collide", object1, object2, collideCallback, processCallback).Bool()
}

// Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// The second parameter can be an array of objects, of differing types.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
func (self *PhysicsNinja) Collide3O(object1 interface{}, object2 interface{}, collideCallback interface{}, processCallback interface{}, callbackContext interface{}) bool{
    return self.Object.Call("collide", object1, object2, collideCallback, processCallback, callbackContext).Bool()
}

// Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// The second parameter can be an array of objects, of differing types.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
func (self *PhysicsNinja) CollideI(args ...interface{}) bool{
    return self.Object.Call("collide", args).Bool()
}

// Internal collision handler.
func (self *PhysicsNinja) CollideHandler(object1 interface{}, object2 interface{}, collideCallback interface{}, processCallback interface{}, callbackContext interface{}, overlapOnly bool) {
    self.Object.Call("collideHandler", object1, object2, collideCallback, processCallback, callbackContext, overlapOnly)
}

// Internal collision handler.
func (self *PhysicsNinja) CollideHandlerI(args ...interface{}) {
    self.Object.Call("collideHandler", args)
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideSpriteVsSprite() {
    self.Object.Call("collideSpriteVsSprite")
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideSpriteVsSpriteI(args ...interface{}) {
    self.Object.Call("collideSpriteVsSprite", args)
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideSpriteVsGroup() {
    self.Object.Call("collideSpriteVsGroup")
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideSpriteVsGroupI(args ...interface{}) {
    self.Object.Call("collideSpriteVsGroup", args)
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideGroupVsSelf() {
    self.Object.Call("collideGroupVsSelf")
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideGroupVsSelfI(args ...interface{}) {
    self.Object.Call("collideGroupVsSelf", args)
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideGroupVsGroup() {
    self.Object.Call("collideGroupVsGroup")
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideGroupVsGroupI(args ...interface{}) {
    self.Object.Call("collideGroupVsGroup", args)
}

// The core separation function to separate two physics bodies.
func (self *PhysicsNinja) Separate(body1 *PhysicsNinjaBody, body2 *PhysicsNinjaBody) bool{
    return self.Object.Call("separate", body1, body2).Bool()
}

// The core separation function to separate two physics bodies.
func (self *PhysicsNinja) SeparateI(args ...interface{}) bool{
    return self.Object.Call("separate", args).Bool()
}
