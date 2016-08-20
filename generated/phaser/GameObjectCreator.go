// Automatic generation for Phaser.GameObjectCreator
// generated file GameObjectCreator.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The GameObjectCreator is a quick way to create common game objects _without_ adding them to the game world.
// The object creator can be accessed with {@linkcode Phaser.Game#make `game.make`}.
type GameObjectCreator struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *GameObjectCreator) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *GameObjectCreator) SetGame(member *Game) {
    self.Set("game", member)
}

// A reference to the game world.
func (self *GameObjectCreator) GetWorld() *World{
    return &World{self.Get("world")}
}

// A reference to the game world.
func (self *GameObjectCreator) SetWorld(member *World) {
    self.Set("world", member)
}



// Create a new Image object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectCreator) ImageI(args ...interface{}) *Image{
    return &Image{self.Call("image", args)}
}

// Create a new Sprite with specific position and sprite sheet key.
func (self *GameObjectCreator) SpriteI(args ...interface{}) *Sprite{
    return &Sprite{self.Call("sprite", args)}
}

// Create a tween object for a specific object.
// 
// The object can be any JavaScript object or Phaser object such as Sprite.
func (self *GameObjectCreator) TweenI(args ...interface{}) *Tween{
    return &Tween{self.Call("tween", args)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectCreator) GroupI(args ...interface{}) *Group{
    return &Group{self.Call("group", args)}
}

// Create a new SpriteBatch.
func (self *GameObjectCreator) SpriteBatchI(args ...interface{}) *SpriteBatch{
    return &SpriteBatch{self.Call("spriteBatch", args)}
}

// Creates a new Sound object.
func (self *GameObjectCreator) AudioI(args ...interface{}) *Sound{
    return &Sound{self.Call("audio", args)}
}

// Creates a new AudioSprite object.
func (self *GameObjectCreator) AudioSpriteI(args ...interface{}) *AudioSprite{
    return &AudioSprite{self.Call("audioSprite", args)}
}

// Creates a new Sound object.
func (self *GameObjectCreator) SoundI(args ...interface{}) *Sound{
    return &Sound{self.Call("sound", args)}
}

// Creates a new TileSprite object.
func (self *GameObjectCreator) TileSpriteI(args ...interface{}) *TileSprite{
    return &TileSprite{self.Call("tileSprite", args)}
}

// Creates a new Rope object.
func (self *GameObjectCreator) RopeI(args ...interface{}) *Rope{
    return &Rope{self.Call("rope", args)}
}

// Creates a new Text object.
func (self *GameObjectCreator) TextI(args ...interface{}) *Text{
    return &Text{self.Call("text", args)}
}

// Creates a new Button object.
func (self *GameObjectCreator) ButtonI(args ...interface{}) *Button{
    return &Button{self.Call("button", args)}
}

// Creates a new Graphics object.
func (self *GameObjectCreator) GraphicsI(args ...interface{}) *Graphics{
    return &Graphics{self.Call("graphics", args)}
}

// Creat a new Emitter.
// 
// An Emitter is a lightweight particle emitter. It can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accorindgly.
func (self *GameObjectCreator) EmitterI(args ...interface{}) *Emitter{
    return &Emitter{self.Call("emitter", args)}
}

// Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectCreator) RetroFontI(args ...interface{}) *RetroFont{
    return &RetroFont{self.Call("retroFont", args)}
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
func (self *GameObjectCreator) BitmapTextI(args ...interface{}) *BitmapText{
    return &BitmapText{self.Call("bitmapText", args)}
}

// Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectCreator) TilemapI(args ...interface{}) {
    self.Call("tilemap", args)
}

// A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectCreator) RenderTextureI(args ...interface{}) *RenderTexture{
    return &RenderTexture{self.Call("renderTexture", args)}
}

// Create a BitmpaData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectCreator) BitmapDataI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Call("bitmapData", args)}
}

// A WebGL shader/filter that can be applied to Sprites.
func (self *GameObjectCreator) FilterI(args ...interface{}) *Filter{
    return &Filter{self.Call("filter", args)}
}
