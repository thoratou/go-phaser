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


// Local reference to game.
func (self *PhysicsNinja) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *PhysicsNinja) SetGame(member *Game) {
    self.Set("game", member)
}

// Local reference to game.time.
func (self *PhysicsNinja) GetTime() *Time{
    return &Time{self.Get("time")}
}

// Local reference to game.time.
func (self *PhysicsNinja) SetTime(member *Time) {
    self.Set("time", member)
}

// The World gravity setting.
func (self *PhysicsNinja) GetGravity() float64{
    return self.Get("gravity").Float()
}

// The World gravity setting.
func (self *PhysicsNinja) SetGravity(member float64) {
    self.Set("gravity", member)
}

// The bounds inside of which the physics world exists. Defaults to match the world bounds.
func (self *PhysicsNinja) GetBounds() *Rectangle{
    return &Rectangle{self.Get("bounds")}
}

// The bounds inside of which the physics world exists. Defaults to match the world bounds.
func (self *PhysicsNinja) SetBounds(member *Rectangle) {
    self.Set("bounds", member)
}

// Used by the QuadTree to set the maximum number of objects per quad.
func (self *PhysicsNinja) GetMaxObjects() float64{
    return self.Get("maxObjects").Float()
}

// Used by the QuadTree to set the maximum number of objects per quad.
func (self *PhysicsNinja) SetMaxObjects(member float64) {
    self.Set("maxObjects", member)
}

// Used by the QuadTree to set the maximum number of iteration levels.
func (self *PhysicsNinja) GetMaxLevels() float64{
    return self.Get("maxLevels").Float()
}

// Used by the QuadTree to set the maximum number of iteration levels.
func (self *PhysicsNinja) SetMaxLevels(member float64) {
    self.Set("maxLevels", member)
}

// The world QuadTree.
func (self *PhysicsNinja) GetQuadTree() *QuadTree{
    return &QuadTree{self.Get("quadTree")}
}

// The world QuadTree.
func (self *PhysicsNinja) SetQuadTree(member *QuadTree) {
    self.Set("quadTree", member)
}



// This will create a Ninja Physics AABB body on the given game object. Its dimensions will match the width and height of the object at the point it is created.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableAABBI(args ...interface{}) {
    self.Call("enableAABB", args)
}

// This will create a Ninja Physics Circle body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableCircleI(args ...interface{}) {
    self.Call("enableCircle", args)
}

// This will create a Ninja Physics Tile body on the given game object. There are 34 different types of tile you can create, including 45 degree slopes,
// convex and concave circles and more. The id parameter controls which Tile type is created, but you can also change it at run-time.
// Note that for all degree based tile types they need to have an equal width and height. If the given object doesn't have equal width and height it will use the width.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableTileI(args ...interface{}) {
    self.Call("enableTile", args)
}

// This will create a Ninja Physics body on the given game object or array of game objects.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the object is destroyed.
func (self *PhysicsNinja) EnableI(args ...interface{}) {
    self.Call("enable", args)
}

// Creates a Ninja Physics body on the given game object.
// A game object can only have 1 physics body active at any one time, and it can't be changed until the body is nulled.
func (self *PhysicsNinja) EnableBodyI(args ...interface{}) {
    self.Call("enableBody", args)
}

// Updates the size of this physics world.
func (self *PhysicsNinja) SetBoundsI(args ...interface{}) {
    self.Call("setBounds", args)
}

// Updates the size of this physics world to match the size of the game world.
func (self *PhysicsNinja) SetBoundsToWorldI(args ...interface{}) {
    self.Call("setBoundsToWorld", args)
}

// Clears all physics bodies from the given TilemapLayer that were created with `World.convertTilemap`.
func (self *PhysicsNinja) ClearTilemapLayerBodiesI(args ...interface{}) {
    self.Call("clearTilemapLayerBodies", args)
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
	array := self.Call("convertTilemap", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Checks for overlaps between two game objects. The objects can be Sprites, Groups or Emitters.
// You can perform Sprite vs. Sprite, Sprite vs. Group and Group vs. Group overlap checks.
// Unlike collide the objects are NOT automatically separated or have any physics applied, they merely test for overlap results.
// The second parameter can be an array of objects, of differing types.
func (self *PhysicsNinja) OverlapI(args ...interface{}) bool{
    return self.Call("overlap", args).Bool()
}

// Checks for collision between two game objects. You can perform Sprite vs. Sprite, Sprite vs. Group, Group vs. Group, Sprite vs. Tilemap Layer or Group vs. Tilemap Layer collisions.
// The second parameter can be an array of objects, of differing types.
// The objects are also automatically separated. If you don't require separation then use ArcadePhysics.overlap instead.
// An optional processCallback can be provided. If given this function will be called when two sprites are found to be colliding. It is called before any separation takes place,
// giving you the chance to perform additional checks. If the function returns true then the collision and separation is carried out. If it returns false it is skipped.
// The collideCallback is an optional function that is only called if two sprites collide. If a processCallback has been set then it needs to return true for collideCallback to be called.
func (self *PhysicsNinja) CollideI(args ...interface{}) bool{
    return self.Call("collide", args).Bool()
}

// Internal collision handler.
func (self *PhysicsNinja) CollideHandlerI(args ...interface{}) {
    self.Call("collideHandler", args)
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideSpriteVsSpriteI(args ...interface{}) {
    self.Call("collideSpriteVsSprite", args)
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideSpriteVsGroupI(args ...interface{}) {
    self.Call("collideSpriteVsGroup", args)
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideGroupVsSelfI(args ...interface{}) {
    self.Call("collideGroupVsSelf", args)
}

// An internal function. Use Phaser.Physics.Ninja.collide instead.
func (self *PhysicsNinja) CollideGroupVsGroupI(args ...interface{}) {
    self.Call("collideGroupVsGroup", args)
}

// The core separation function to separate two physics bodies.
func (self *PhysicsNinja) SeparateI(args ...interface{}) bool{
    return self.Call("separate", args).Bool()
}
