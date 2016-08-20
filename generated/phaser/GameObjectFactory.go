// Automatic generation for Phaser.GameObjectFactory
// generated file GameObjectFactory.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The GameObjectFactory is a quick way to create many common game objects
// using {@linkcode Phaser.Game#add `game.add`}.
// 
// Created objects are _automatically added_ to the appropriate Manager, World, or manually specified parent Group.
type GameObjectFactory struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *GameObjectFactory) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *GameObjectFactory) SetGame(member Game) {
    self.Set("game", member)
}

// A reference to the game world.
func (self *GameObjectFactory) GetWorld() World{
    return World{self.Get("world")}
}

// A reference to the game world.
func (self *GameObjectFactory) SetWorld(member World) {
    self.Set("world", member)
}



// Adds an existing display object to the game world.
func (self *GameObjectFactory) ExistingI(args ...interface{}) interface{}{
    return self.Call("existing", args)
}

// Weapons provide the ability to easily create a bullet pool and manager.
// 
// Weapons fire Phaser.Bullet objects, which are essentially Sprites with a few extra properties.
// The Bullets are enabled for Arcade Physics. They do not currently work with P2 Physics.
// 
// The Bullets are created inside of `Weapon.bullets`, which is a Phaser.Group instance. Anything you
// can usually do with a Group, such as move it around the display list, iterate it, etc can be done
// to the bullets Group too.
// 
// Bullets can have textures and even animations. You can control the speed at which they are fired,
// the firing rate, the firing angle, and even set things like gravity for them.
func (self *GameObjectFactory) WeaponI(args ...interface{}) Weapon{
    return Weapon{self.Call("weapon", args)}
}

// Create a new `Image` object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// 
// It can still rotate, scale, crop and receive input events. 
// This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectFactory) ImageI(args ...interface{}) Image{
    return Image{self.Call("image", args)}
}

// Create a new Sprite with specific position and sprite sheet key.
// 
// At its most basic a Sprite consists of a set of coordinates and a texture that is used when rendered.
// They also contain additional properties allowing for physics motion (via Sprite.body), input handling (via Sprite.input),
// events (via Sprite.events), animation (via Sprite.animations), camera culling and more. Please see the Examples for use cases.
func (self *GameObjectFactory) SpriteI(args ...interface{}) Sprite{
    return Sprite{self.Call("sprite", args)}
}

// Create a new Creature Animation object.
// 
// Creature is a custom Game Object used in conjunction with the Creature Runtime libraries by Kestrel Moon Studios.
// 
// It allows you to display animated Game Objects that were created with the [Creature Automated Animation Tool](http://www.kestrelmoon.com/creature/).
// 
// Note 1: You can only use Phaser.Creature objects in WebGL enabled games. They do not work in Canvas mode games.
// 
// Note 2: You must use a build of Phaser that includes the CreatureMeshBone.js runtime and gl-matrix.js, or have them
// loaded before your Phaser game boots.
// 
// See the Phaser custom build process for more details.
func (self *GameObjectFactory) CreatureI(args ...interface{}) Creature{
    return Creature{self.Call("creature", args)}
}

// Create a tween on a specific object.
// 
// The object can be any JavaScript object or Phaser object such as Sprite.
func (self *GameObjectFactory) TweenI(args ...interface{}) Tween{
    return Tween{self.Call("tween", args)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectFactory) GroupI(args ...interface{}) Group{
    return Group{self.Call("group", args)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
// 
// A Physics Group is the same as an ordinary Group except that is has enableBody turned on by default, so any Sprites it creates
// are automatically given a physics body.
func (self *GameObjectFactory) PhysicsGroupI(args ...interface{}) Group{
    return Group{self.Call("physicsGroup", args)}
}

// A SpriteBatch is a really fast version of a Phaser Group built solely for speed.
// Use when you need a lot of sprites or particles all sharing the same texture.
// The speed gains are specifically for WebGL. In Canvas mode you won't see any real difference.
func (self *GameObjectFactory) SpriteBatchI(args ...interface{}) SpriteBatch{
    return SpriteBatch{self.Call("spriteBatch", args)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) AudioI(args ...interface{}) Sound{
    return Sound{self.Call("audio", args)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) SoundI(args ...interface{}) Sound{
    return Sound{self.Call("sound", args)}
}

// Creates a new AudioSprite object.
func (self *GameObjectFactory) AudioSpriteI(args ...interface{}) AudioSprite{
    return AudioSprite{self.Call("audioSprite", args)}
}

// Creates a new TileSprite object.
func (self *GameObjectFactory) TileSpriteI(args ...interface{}) TileSprite{
    return TileSprite{self.Call("tileSprite", args)}
}

// Creates a new Rope object.
// 
// Example usage: https://github.com/codevinsky/phaser-rope-demo/blob/master/dist/demo.js
func (self *GameObjectFactory) RopeI(args ...interface{}) Rope{
    return Rope{self.Call("rope", args)}
}

// Creates a new Text object.
func (self *GameObjectFactory) TextI(args ...interface{}) Text{
    return Text{self.Call("text", args)}
}

// Creates a new Button object.
func (self *GameObjectFactory) ButtonI(args ...interface{}) Button{
    return Button{self.Call("button", args)}
}

// Creates a new Graphics object.
func (self *GameObjectFactory) GraphicsI(args ...interface{}) Graphics{
    return Graphics{self.Call("graphics", args)}
}

// Create a new Emitter.
// 
// A particle emitter can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accordingly.
func (self *GameObjectFactory) EmitterI(args ...interface{}) ParticlesArcadeEmitter{
    return ParticlesArcadeEmitter{self.Call("emitter", args)}
}

// Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectFactory) RetroFontI(args ...interface{}) RetroFont{
    return RetroFont{self.Call("retroFont", args)}
}

// Create a new BitmapText object.
// 
// BitmapText objects work by taking a texture file and an XML file that describes the font structure.
// It then generates a new Sprite object for each letter of the text, proportionally spaced out and aligned to 
// match the font structure.
// 
// BitmapText objects are less flexible than Text objects, in that they have less features such as shadows, fills and the ability 
// to use Web Fonts. However you trade this flexibility for pure rendering speed. You can also create visually compelling BitmapTexts by 
// processing the font texture in an image editor first, applying fills and any other effects required.
// 
// To create multi-line text insert \r, \n or \r\n escape codes into the text string.
// 
// To create a BitmapText data files you can use:
// 
// BMFont (Windows, free): http://www.angelcode.com/products/bmfont/
// Glyph Designer (OS X, commercial): http://www.71squared.com/en/glyphdesigner
// Littera (Web-based, free): http://kvazars.com/littera/
func (self *GameObjectFactory) BitmapTextI(args ...interface{}) BitmapText{
    return BitmapText{self.Call("bitmapText", args)}
}

// Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectFactory) TilemapI(args ...interface{}) Tilemap{
    return Tilemap{self.Call("tilemap", args)}
}

// A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectFactory) RenderTextureI(args ...interface{}) RenderTexture{
    return RenderTexture{self.Call("renderTexture", args)}
}

// Create a Video object.
// 
// This will return a Phaser.Video object which you can pass to a Sprite to be used as a texture.
func (self *GameObjectFactory) VideoI(args ...interface{}) Video{
    return Video{self.Call("video", args)}
}

// Create a BitmapData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectFactory) BitmapDataI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("bitmapData", args)}
}

// A WebGL shader/filter that can be applied to Sprites.
func (self *GameObjectFactory) FilterI(args ...interface{}) Filter{
    return Filter{self.Call("filter", args)}
}

// Add a new Plugin into the PluginManager.
// 
// The Plugin must have 2 properties: `game` and `parent`. Plugin.game is set to the game reference the PluginManager uses, and parent is set to the PluginManager.
func (self *GameObjectFactory) PluginI(args ...interface{}) Plugin{
    return Plugin{self.Call("plugin", args)}
}
