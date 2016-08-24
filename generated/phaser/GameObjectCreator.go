// Package phaser Automatic generation for Phaser.GameObjectCreator
// generated file GameObjectCreator.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// GameObjectCreator The GameObjectCreator is a quick way to create common game objects _without_ adding them to the game world.
// The object creator can be accessed with {@linkcode Phaser.Game#make `game.make`}.
type GameObjectCreator struct {
    *js.Object
}

// NewGameObjectCreator The GameObjectCreator is a quick way to create common game objects _without_ adding them to the game world.
// The object creator can be accessed with {@linkcode Phaser.Game#make `game.make`}.
func NewGameObjectCreator(game *Game) *GameObjectCreator {
    return &GameObjectCreator{js.Global.Get("Phaser").Get("GameObjectCreator").New(game)}
}
// NewGameObjectCreatorI The GameObjectCreator is a quick way to create common game objects _without_ adding them to the game world.
// The object creator can be accessed with {@linkcode Phaser.Game#make `game.make`}.
func NewGameObjectCreatorI(args ...interface{}) *GameObjectCreator {
    return &GameObjectCreator{js.Global.Get("Phaser").Get("GameObjectCreator").New(args)}
}



// Game A reference to the currently running Game.
func (self *GameObjectCreator) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *GameObjectCreator) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// World A reference to the game world.
func (self *GameObjectCreator) World() *World{
    return &World{self.Object.Get("world")}
}

// SetWorldA A reference to the game world.
func (self *GameObjectCreator) SetWorldA(member *World) {
    self.Object.Set("world", member)
}


// Image Create a new Image object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectCreator) Image(x int, y int, key interface{}) *Image{
    return &Image{self.Object.Call("image", x, y, key)}
}

// Image1O Create a new Image object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectCreator) Image1O(x int, y int, key interface{}, frame interface{}) *Image{
    return &Image{self.Object.Call("image", x, y, key, frame)}
}

// ImageI Create a new Image object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// It can still rotate, scale, crop and receive input events. This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectCreator) ImageI(args ...interface{}) *Image{
    return &Image{self.Object.Call("image", args)}
}

// Sprite Create a new Sprite with specific position and sprite sheet key.
func (self *GameObjectCreator) Sprite(x int, y int, key interface{}) *Sprite{
    return &Sprite{self.Object.Call("sprite", x, y, key)}
}

// Sprite1O Create a new Sprite with specific position and sprite sheet key.
func (self *GameObjectCreator) Sprite1O(x int, y int, key interface{}, frame interface{}) *Sprite{
    return &Sprite{self.Object.Call("sprite", x, y, key, frame)}
}

// SpriteI Create a new Sprite with specific position and sprite sheet key.
func (self *GameObjectCreator) SpriteI(args ...interface{}) *Sprite{
    return &Sprite{self.Object.Call("sprite", args)}
}

// Tween Create a tween object for a specific object.
// 
// The object can be any JavaScript object or Phaser object such as Sprite.
func (self *GameObjectCreator) Tween(obj interface{}) *Tween{
    return &Tween{self.Object.Call("tween", obj)}
}

// TweenI Create a tween object for a specific object.
// 
// The object can be any JavaScript object or Phaser object such as Sprite.
func (self *GameObjectCreator) TweenI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("tween", args)}
}

// Group A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectCreator) Group(parent interface{}) *Group{
    return &Group{self.Object.Call("group", parent)}
}

// Group1O A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectCreator) Group1O(parent interface{}, name string) *Group{
    return &Group{self.Object.Call("group", parent, name)}
}

// Group2O A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectCreator) Group2O(parent interface{}, name string, addToStage bool) *Group{
    return &Group{self.Object.Call("group", parent, name, addToStage)}
}

// Group3O A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectCreator) Group3O(parent interface{}, name string, addToStage bool, enableBody bool) *Group{
    return &Group{self.Object.Call("group", parent, name, addToStage, enableBody)}
}

// Group4O A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectCreator) Group4O(parent interface{}, name string, addToStage bool, enableBody bool, physicsBodyType int) *Group{
    return &Group{self.Object.Call("group", parent, name, addToStage, enableBody, physicsBodyType)}
}

// GroupI A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectCreator) GroupI(args ...interface{}) *Group{
    return &Group{self.Object.Call("group", args)}
}

// SpriteBatch Create a new SpriteBatch.
func (self *GameObjectCreator) SpriteBatch(parent interface{}) *SpriteBatch{
    return &SpriteBatch{self.Object.Call("spriteBatch", parent)}
}

// SpriteBatch1O Create a new SpriteBatch.
func (self *GameObjectCreator) SpriteBatch1O(parent interface{}, name string) *SpriteBatch{
    return &SpriteBatch{self.Object.Call("spriteBatch", parent, name)}
}

// SpriteBatch2O Create a new SpriteBatch.
func (self *GameObjectCreator) SpriteBatch2O(parent interface{}, name string, addToStage bool) *SpriteBatch{
    return &SpriteBatch{self.Object.Call("spriteBatch", parent, name, addToStage)}
}

// SpriteBatchI Create a new SpriteBatch.
func (self *GameObjectCreator) SpriteBatchI(args ...interface{}) *SpriteBatch{
    return &SpriteBatch{self.Object.Call("spriteBatch", args)}
}

// Audio Creates a new Sound object.
func (self *GameObjectCreator) Audio(key string) *Sound{
    return &Sound{self.Object.Call("audio", key)}
}

// Audio1O Creates a new Sound object.
func (self *GameObjectCreator) Audio1O(key string, volume int) *Sound{
    return &Sound{self.Object.Call("audio", key, volume)}
}

// Audio2O Creates a new Sound object.
func (self *GameObjectCreator) Audio2O(key string, volume int, loop bool) *Sound{
    return &Sound{self.Object.Call("audio", key, volume, loop)}
}

// Audio3O Creates a new Sound object.
func (self *GameObjectCreator) Audio3O(key string, volume int, loop bool, connect bool) *Sound{
    return &Sound{self.Object.Call("audio", key, volume, loop, connect)}
}

// AudioI Creates a new Sound object.
func (self *GameObjectCreator) AudioI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("audio", args)}
}

// AudioSprite Creates a new AudioSprite object.
func (self *GameObjectCreator) AudioSprite(key string) *AudioSprite{
    return &AudioSprite{self.Object.Call("audioSprite", key)}
}

// AudioSpriteI Creates a new AudioSprite object.
func (self *GameObjectCreator) AudioSpriteI(args ...interface{}) *AudioSprite{
    return &AudioSprite{self.Object.Call("audioSprite", args)}
}

// Sound Creates a new Sound object.
func (self *GameObjectCreator) Sound(key string) *Sound{
    return &Sound{self.Object.Call("sound", key)}
}

// Sound1O Creates a new Sound object.
func (self *GameObjectCreator) Sound1O(key string, volume int) *Sound{
    return &Sound{self.Object.Call("sound", key, volume)}
}

// Sound2O Creates a new Sound object.
func (self *GameObjectCreator) Sound2O(key string, volume int, loop bool) *Sound{
    return &Sound{self.Object.Call("sound", key, volume, loop)}
}

// Sound3O Creates a new Sound object.
func (self *GameObjectCreator) Sound3O(key string, volume int, loop bool, connect bool) *Sound{
    return &Sound{self.Object.Call("sound", key, volume, loop, connect)}
}

// SoundI Creates a new Sound object.
func (self *GameObjectCreator) SoundI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("sound", args)}
}

// TileSprite Creates a new TileSprite object.
func (self *GameObjectCreator) TileSprite(x int, y int, width int, height int, key interface{}, frame interface{}) *TileSprite{
    return &TileSprite{self.Object.Call("tileSprite", x, y, width, height, key, frame)}
}

// TileSpriteI Creates a new TileSprite object.
func (self *GameObjectCreator) TileSpriteI(args ...interface{}) *TileSprite{
    return &TileSprite{self.Object.Call("tileSprite", args)}
}

// Rope Creates a new Rope object.
func (self *GameObjectCreator) Rope(x int, y int, width int, height int, key interface{}, frame interface{}) *Rope{
    return &Rope{self.Object.Call("rope", x, y, width, height, key, frame)}
}

// RopeI Creates a new Rope object.
func (self *GameObjectCreator) RopeI(args ...interface{}) *Rope{
    return &Rope{self.Object.Call("rope", args)}
}

// Text Creates a new Text object.
func (self *GameObjectCreator) Text(x int, y int, text string, style interface{}) *Text{
    return &Text{self.Object.Call("text", x, y, text, style)}
}

// TextI Creates a new Text object.
func (self *GameObjectCreator) TextI(args ...interface{}) *Text{
    return &Text{self.Object.Call("text", args)}
}

// Button Creates a new Button object.
func (self *GameObjectCreator) Button() *Button{
    return &Button{self.Object.Call("button")}
}

// Button1O Creates a new Button object.
func (self *GameObjectCreator) Button1O(x int) *Button{
    return &Button{self.Object.Call("button", x)}
}

// Button2O Creates a new Button object.
func (self *GameObjectCreator) Button2O(x int, y int) *Button{
    return &Button{self.Object.Call("button", x, y)}
}

// Button3O Creates a new Button object.
func (self *GameObjectCreator) Button3O(x int, y int, key string) *Button{
    return &Button{self.Object.Call("button", x, y, key)}
}

// Button4O Creates a new Button object.
func (self *GameObjectCreator) Button4O(x int, y int, key string, callback interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback)}
}

// Button5O Creates a new Button object.
func (self *GameObjectCreator) Button5O(x int, y int, key string, callback interface{}, callbackContext interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext)}
}

// Button6O Creates a new Button object.
func (self *GameObjectCreator) Button6O(x int, y int, key string, callback interface{}, callbackContext interface{}, overFrame interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext, overFrame)}
}

// Button7O Creates a new Button object.
func (self *GameObjectCreator) Button7O(x int, y int, key string, callback interface{}, callbackContext interface{}, overFrame interface{}, outFrame interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext, overFrame, outFrame)}
}

// Button8O Creates a new Button object.
func (self *GameObjectCreator) Button8O(x int, y int, key string, callback interface{}, callbackContext interface{}, overFrame interface{}, outFrame interface{}, downFrame interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext, overFrame, outFrame, downFrame)}
}

// Button9O Creates a new Button object.
func (self *GameObjectCreator) Button9O(x int, y int, key string, callback interface{}, callbackContext interface{}, overFrame interface{}, outFrame interface{}, downFrame interface{}, upFrame interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext, overFrame, outFrame, downFrame, upFrame)}
}

// ButtonI Creates a new Button object.
func (self *GameObjectCreator) ButtonI(args ...interface{}) *Button{
    return &Button{self.Object.Call("button", args)}
}

// Graphics Creates a new Graphics object.
func (self *GameObjectCreator) Graphics() *Graphics{
    return &Graphics{self.Object.Call("graphics")}
}

// Graphics1O Creates a new Graphics object.
func (self *GameObjectCreator) Graphics1O(x int) *Graphics{
    return &Graphics{self.Object.Call("graphics", x)}
}

// Graphics2O Creates a new Graphics object.
func (self *GameObjectCreator) Graphics2O(x int, y int) *Graphics{
    return &Graphics{self.Object.Call("graphics", x, y)}
}

// GraphicsI Creates a new Graphics object.
func (self *GameObjectCreator) GraphicsI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("graphics", args)}
}

// Emitter Creat a new Emitter.
// 
// An Emitter is a lightweight particle emitter. It can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accorindgly.
func (self *GameObjectCreator) Emitter() *Emitter{
    return &Emitter{self.Object.Call("emitter")}
}

// Emitter1O Creat a new Emitter.
// 
// An Emitter is a lightweight particle emitter. It can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accorindgly.
func (self *GameObjectCreator) Emitter1O(x int) *Emitter{
    return &Emitter{self.Object.Call("emitter", x)}
}

// Emitter2O Creat a new Emitter.
// 
// An Emitter is a lightweight particle emitter. It can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accorindgly.
func (self *GameObjectCreator) Emitter2O(x int, y int) *Emitter{
    return &Emitter{self.Object.Call("emitter", x, y)}
}

// Emitter3O Creat a new Emitter.
// 
// An Emitter is a lightweight particle emitter. It can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accorindgly.
func (self *GameObjectCreator) Emitter3O(x int, y int, maxParticles int) *Emitter{
    return &Emitter{self.Object.Call("emitter", x, y, maxParticles)}
}

// EmitterI Creat a new Emitter.
// 
// An Emitter is a lightweight particle emitter. It can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accorindgly.
func (self *GameObjectCreator) EmitterI(args ...interface{}) *Emitter{
    return &Emitter{self.Object.Call("emitter", args)}
}

// RetroFont Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectCreator) RetroFont(font string, characterWidth int, characterHeight int, chars string, charsPerRow int) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", font, characterWidth, characterHeight, chars, charsPerRow)}
}

// RetroFont1O Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectCreator) RetroFont1O(font string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", font, characterWidth, characterHeight, chars, charsPerRow, xSpacing)}
}

// RetroFont2O Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectCreator) RetroFont2O(font string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int, ySpacing int) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", font, characterWidth, characterHeight, chars, charsPerRow, xSpacing, ySpacing)}
}

// RetroFont3O Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectCreator) RetroFont3O(font string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int, ySpacing int, xOffset int) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", font, characterWidth, characterHeight, chars, charsPerRow, xSpacing, ySpacing, xOffset)}
}

// RetroFont4O Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectCreator) RetroFont4O(font string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int, ySpacing int, xOffset int, yOffset int) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", font, characterWidth, characterHeight, chars, charsPerRow, xSpacing, ySpacing, xOffset, yOffset)}
}

// RetroFontI Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectCreator) RetroFontI(args ...interface{}) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", args)}
}

// BitmapText Create a new BitmapText object.
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
func (self *GameObjectCreator) BitmapText(x int, y int, font string) *BitmapText{
    return &BitmapText{self.Object.Call("bitmapText", x, y, font)}
}

// BitmapText1O Create a new BitmapText object.
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
func (self *GameObjectCreator) BitmapText1O(x int, y int, font string, text string) *BitmapText{
    return &BitmapText{self.Object.Call("bitmapText", x, y, font, text)}
}

// BitmapText2O Create a new BitmapText object.
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
func (self *GameObjectCreator) BitmapText2O(x int, y int, font string, text string, size int) *BitmapText{
    return &BitmapText{self.Object.Call("bitmapText", x, y, font, text, size)}
}

// BitmapText3O Create a new BitmapText object.
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
func (self *GameObjectCreator) BitmapText3O(x int, y int, font string, text string, size int, align string) *BitmapText{
    return &BitmapText{self.Object.Call("bitmapText", x, y, font, text, size, align)}
}

// BitmapTextI Create a new BitmapText object.
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
    return &BitmapText{self.Object.Call("bitmapText", args)}
}

// Tilemap Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectCreator) Tilemap() {
    self.Object.Call("tilemap")
}

// Tilemap1O Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectCreator) Tilemap1O(key string) {
    self.Object.Call("tilemap", key)
}

// Tilemap2O Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectCreator) Tilemap2O(key string, tileWidth int) {
    self.Object.Call("tilemap", key, tileWidth)
}

// Tilemap3O Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectCreator) Tilemap3O(key string, tileWidth int, tileHeight int) {
    self.Object.Call("tilemap", key, tileWidth, tileHeight)
}

// Tilemap4O Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectCreator) Tilemap4O(key string, tileWidth int, tileHeight int, width int) {
    self.Object.Call("tilemap", key, tileWidth, tileHeight, width)
}

// Tilemap5O Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectCreator) Tilemap5O(key string, tileWidth int, tileHeight int, width int, height int) {
    self.Object.Call("tilemap", key, tileWidth, tileHeight, width, height)
}

// TilemapI Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectCreator) TilemapI(args ...interface{}) {
    self.Object.Call("tilemap", args)
}

// RenderTexture A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectCreator) RenderTexture() *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture")}
}

// RenderTexture1O A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectCreator) RenderTexture1O(width int) *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture", width)}
}

// RenderTexture2O A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectCreator) RenderTexture2O(width int, height int) *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture", width, height)}
}

// RenderTexture3O A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectCreator) RenderTexture3O(width int, height int, key string) *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture", width, height, key)}
}

// RenderTexture4O A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectCreator) RenderTexture4O(width int, height int, key string, addToCache bool) *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture", width, height, key, addToCache)}
}

// RenderTextureI A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectCreator) RenderTextureI(args ...interface{}) *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture", args)}
}

// BitmapData Create a BitmpaData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectCreator) BitmapData() *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData")}
}

// BitmapData1O Create a BitmpaData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectCreator) BitmapData1O(width int) *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData", width)}
}

// BitmapData2O Create a BitmpaData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectCreator) BitmapData2O(width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData", width, height)}
}

// BitmapData3O Create a BitmpaData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectCreator) BitmapData3O(width int, height int, key string) *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData", width, height, key)}
}

// BitmapData4O Create a BitmpaData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectCreator) BitmapData4O(width int, height int, key string, addToCache bool) *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData", width, height, key, addToCache)}
}

// BitmapDataI Create a BitmpaData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectCreator) BitmapDataI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData", args)}
}

// Filter A WebGL shader/filter that can be applied to Sprites.
func (self *GameObjectCreator) Filter(filter string, args interface{}) *Filter{
    return &Filter{self.Object.Call("filter", filter, args)}
}

// FilterI A WebGL shader/filter that can be applied to Sprites.
func (self *GameObjectCreator) FilterI(args ...interface{}) *Filter{
    return &Filter{self.Object.Call("filter", args)}
}

