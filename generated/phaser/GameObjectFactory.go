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


// The GameObjectFactory is a quick way to create many common game objects
// using {@linkcode Phaser.Game#add `game.add`}.
// 
// Created objects are _automatically added_ to the appropriate Manager, World, or manually specified parent Group.
func NewGameObjectFactory(game *Game) *GameObjectFactory {
    return &GameObjectFactory{js.Global.Call("Phaser.GameObjectFactory", game)}
}

// The GameObjectFactory is a quick way to create many common game objects
// using {@linkcode Phaser.Game#add `game.add`}.
// 
// Created objects are _automatically added_ to the appropriate Manager, World, or manually specified parent Group.
func NewGameObjectFactoryI(args ...interface{}) *GameObjectFactory {
    return &GameObjectFactory{js.Global.Call("Phaser.GameObjectFactory", args)}
}



// A reference to the currently running Game.
func (self *GameObjectFactory) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *GameObjectFactory) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// A reference to the game world.
func (self *GameObjectFactory) GetWorldA() *World{
    return &World{self.Object.Get("world")}
}

// A reference to the game world.
func (self *GameObjectFactory) SetWorldA(member *World) {
    self.Object.Set("world", member)
}



// Adds an existing display object to the game world.
func (self *GameObjectFactory) Existing(object interface{}) interface{}{
    return self.Object.Call("existing", object)
}

// Adds an existing display object to the game world.
func (self *GameObjectFactory) ExistingI(args ...interface{}) interface{}{
    return self.Object.Call("existing", args)
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
func (self *GameObjectFactory) Weapon() *Weapon{
    return &Weapon{self.Object.Call("weapon")}
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
func (self *GameObjectFactory) Weapon1O(quantity int) *Weapon{
    return &Weapon{self.Object.Call("weapon", quantity)}
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
func (self *GameObjectFactory) Weapon2O(quantity int, key interface{}) *Weapon{
    return &Weapon{self.Object.Call("weapon", quantity, key)}
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
func (self *GameObjectFactory) Weapon3O(quantity int, key interface{}, frame interface{}) *Weapon{
    return &Weapon{self.Object.Call("weapon", quantity, key, frame)}
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
func (self *GameObjectFactory) Weapon4O(quantity int, key interface{}, frame interface{}, group *Group) *Weapon{
    return &Weapon{self.Object.Call("weapon", quantity, key, frame, group)}
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
func (self *GameObjectFactory) WeaponI(args ...interface{}) *Weapon{
    return &Weapon{self.Object.Call("weapon", args)}
}

// Create a new `Image` object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// 
// It can still rotate, scale, crop and receive input events. 
// This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectFactory) Image() *Image{
    return &Image{self.Object.Call("image")}
}

// Create a new `Image` object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// 
// It can still rotate, scale, crop and receive input events. 
// This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectFactory) Image1O(x int) *Image{
    return &Image{self.Object.Call("image", x)}
}

// Create a new `Image` object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// 
// It can still rotate, scale, crop and receive input events. 
// This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectFactory) Image2O(x int, y int) *Image{
    return &Image{self.Object.Call("image", x, y)}
}

// Create a new `Image` object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// 
// It can still rotate, scale, crop and receive input events. 
// This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectFactory) Image3O(x int, y int, key interface{}) *Image{
    return &Image{self.Object.Call("image", x, y, key)}
}

// Create a new `Image` object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// 
// It can still rotate, scale, crop and receive input events. 
// This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectFactory) Image4O(x int, y int, key interface{}, frame interface{}) *Image{
    return &Image{self.Object.Call("image", x, y, key, frame)}
}

// Create a new `Image` object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// 
// It can still rotate, scale, crop and receive input events. 
// This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectFactory) Image5O(x int, y int, key interface{}, frame interface{}, group *Group) *Image{
    return &Image{self.Object.Call("image", x, y, key, frame, group)}
}

// Create a new `Image` object.
// 
// An Image is a light-weight object you can use to display anything that doesn't need physics or animation.
// 
// It can still rotate, scale, crop and receive input events. 
// This makes it perfect for logos, backgrounds, simple buttons and other non-Sprite graphics.
func (self *GameObjectFactory) ImageI(args ...interface{}) *Image{
    return &Image{self.Object.Call("image", args)}
}

// Create a new Sprite with specific position and sprite sheet key.
// 
// At its most basic a Sprite consists of a set of coordinates and a texture that is used when rendered.
// They also contain additional properties allowing for physics motion (via Sprite.body), input handling (via Sprite.input),
// events (via Sprite.events), animation (via Sprite.animations), camera culling and more. Please see the Examples for use cases.
func (self *GameObjectFactory) Sprite() *Sprite{
    return &Sprite{self.Object.Call("sprite")}
}

// Create a new Sprite with specific position and sprite sheet key.
// 
// At its most basic a Sprite consists of a set of coordinates and a texture that is used when rendered.
// They also contain additional properties allowing for physics motion (via Sprite.body), input handling (via Sprite.input),
// events (via Sprite.events), animation (via Sprite.animations), camera culling and more. Please see the Examples for use cases.
func (self *GameObjectFactory) Sprite1O(x int) *Sprite{
    return &Sprite{self.Object.Call("sprite", x)}
}

// Create a new Sprite with specific position and sprite sheet key.
// 
// At its most basic a Sprite consists of a set of coordinates and a texture that is used when rendered.
// They also contain additional properties allowing for physics motion (via Sprite.body), input handling (via Sprite.input),
// events (via Sprite.events), animation (via Sprite.animations), camera culling and more. Please see the Examples for use cases.
func (self *GameObjectFactory) Sprite2O(x int, y int) *Sprite{
    return &Sprite{self.Object.Call("sprite", x, y)}
}

// Create a new Sprite with specific position and sprite sheet key.
// 
// At its most basic a Sprite consists of a set of coordinates and a texture that is used when rendered.
// They also contain additional properties allowing for physics motion (via Sprite.body), input handling (via Sprite.input),
// events (via Sprite.events), animation (via Sprite.animations), camera culling and more. Please see the Examples for use cases.
func (self *GameObjectFactory) Sprite3O(x int, y int, key interface{}) *Sprite{
    return &Sprite{self.Object.Call("sprite", x, y, key)}
}

// Create a new Sprite with specific position and sprite sheet key.
// 
// At its most basic a Sprite consists of a set of coordinates and a texture that is used when rendered.
// They also contain additional properties allowing for physics motion (via Sprite.body), input handling (via Sprite.input),
// events (via Sprite.events), animation (via Sprite.animations), camera culling and more. Please see the Examples for use cases.
func (self *GameObjectFactory) Sprite4O(x int, y int, key interface{}, frame interface{}) *Sprite{
    return &Sprite{self.Object.Call("sprite", x, y, key, frame)}
}

// Create a new Sprite with specific position and sprite sheet key.
// 
// At its most basic a Sprite consists of a set of coordinates and a texture that is used when rendered.
// They also contain additional properties allowing for physics motion (via Sprite.body), input handling (via Sprite.input),
// events (via Sprite.events), animation (via Sprite.animations), camera culling and more. Please see the Examples for use cases.
func (self *GameObjectFactory) Sprite5O(x int, y int, key interface{}, frame interface{}, group *Group) *Sprite{
    return &Sprite{self.Object.Call("sprite", x, y, key, frame, group)}
}

// Create a new Sprite with specific position and sprite sheet key.
// 
// At its most basic a Sprite consists of a set of coordinates and a texture that is used when rendered.
// They also contain additional properties allowing for physics motion (via Sprite.body), input handling (via Sprite.input),
// events (via Sprite.events), animation (via Sprite.animations), camera culling and more. Please see the Examples for use cases.
func (self *GameObjectFactory) SpriteI(args ...interface{}) *Sprite{
    return &Sprite{self.Object.Call("sprite", args)}
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
func (self *GameObjectFactory) Creature() *Creature{
    return &Creature{self.Object.Call("creature")}
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
func (self *GameObjectFactory) Creature1O(x int) *Creature{
    return &Creature{self.Object.Call("creature", x)}
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
func (self *GameObjectFactory) Creature2O(x int, y int) *Creature{
    return &Creature{self.Object.Call("creature", x, y)}
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
func (self *GameObjectFactory) Creature3O(x int, y int, key interface{}) *Creature{
    return &Creature{self.Object.Call("creature", x, y, key)}
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
func (self *GameObjectFactory) Creature4O(x int, y int, key interface{}, group *Group) *Creature{
    return &Creature{self.Object.Call("creature", x, y, key, group)}
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
func (self *GameObjectFactory) CreatureI(args ...interface{}) *Creature{
    return &Creature{self.Object.Call("creature", args)}
}

// Create a tween on a specific object.
// 
// The object can be any JavaScript object or Phaser object such as Sprite.
func (self *GameObjectFactory) Tween(object interface{}) *Tween{
    return &Tween{self.Object.Call("tween", object)}
}

// Create a tween on a specific object.
// 
// The object can be any JavaScript object or Phaser object such as Sprite.
func (self *GameObjectFactory) TweenI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("tween", args)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectFactory) Group() *Group{
    return &Group{self.Object.Call("group")}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectFactory) Group1O(parent interface{}) *Group{
    return &Group{self.Object.Call("group", parent)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectFactory) Group2O(parent interface{}, name string) *Group{
    return &Group{self.Object.Call("group", parent, name)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectFactory) Group3O(parent interface{}, name string, addToStage bool) *Group{
    return &Group{self.Object.Call("group", parent, name, addToStage)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectFactory) Group4O(parent interface{}, name string, addToStage bool, enableBody bool) *Group{
    return &Group{self.Object.Call("group", parent, name, addToStage, enableBody)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectFactory) Group5O(parent interface{}, name string, addToStage bool, enableBody bool, physicsBodyType int) *Group{
    return &Group{self.Object.Call("group", parent, name, addToStage, enableBody, physicsBodyType)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
func (self *GameObjectFactory) GroupI(args ...interface{}) *Group{
    return &Group{self.Object.Call("group", args)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
// 
// A Physics Group is the same as an ordinary Group except that is has enableBody turned on by default, so any Sprites it creates
// are automatically given a physics body.
func (self *GameObjectFactory) PhysicsGroup() *Group{
    return &Group{self.Object.Call("physicsGroup")}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
// 
// A Physics Group is the same as an ordinary Group except that is has enableBody turned on by default, so any Sprites it creates
// are automatically given a physics body.
func (self *GameObjectFactory) PhysicsGroup1O(physicsBodyType int) *Group{
    return &Group{self.Object.Call("physicsGroup", physicsBodyType)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
// 
// A Physics Group is the same as an ordinary Group except that is has enableBody turned on by default, so any Sprites it creates
// are automatically given a physics body.
func (self *GameObjectFactory) PhysicsGroup2O(physicsBodyType int, parent interface{}) *Group{
    return &Group{self.Object.Call("physicsGroup", physicsBodyType, parent)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
// 
// A Physics Group is the same as an ordinary Group except that is has enableBody turned on by default, so any Sprites it creates
// are automatically given a physics body.
func (self *GameObjectFactory) PhysicsGroup3O(physicsBodyType int, parent interface{}, name string) *Group{
    return &Group{self.Object.Call("physicsGroup", physicsBodyType, parent, name)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
// 
// A Physics Group is the same as an ordinary Group except that is has enableBody turned on by default, so any Sprites it creates
// are automatically given a physics body.
func (self *GameObjectFactory) PhysicsGroup4O(physicsBodyType int, parent interface{}, name string, addToStage bool) *Group{
    return &Group{self.Object.Call("physicsGroup", physicsBodyType, parent, name, addToStage)}
}

// A Group is a container for display objects that allows for fast pooling, recycling and collision checks.
// 
// A Physics Group is the same as an ordinary Group except that is has enableBody turned on by default, so any Sprites it creates
// are automatically given a physics body.
func (self *GameObjectFactory) PhysicsGroupI(args ...interface{}) *Group{
    return &Group{self.Object.Call("physicsGroup", args)}
}

// A SpriteBatch is a really fast version of a Phaser Group built solely for speed.
// Use when you need a lot of sprites or particles all sharing the same texture.
// The speed gains are specifically for WebGL. In Canvas mode you won't see any real difference.
func (self *GameObjectFactory) SpriteBatch(parent interface{}) *SpriteBatch{
    return &SpriteBatch{self.Object.Call("spriteBatch", parent)}
}

// A SpriteBatch is a really fast version of a Phaser Group built solely for speed.
// Use when you need a lot of sprites or particles all sharing the same texture.
// The speed gains are specifically for WebGL. In Canvas mode you won't see any real difference.
func (self *GameObjectFactory) SpriteBatch1O(parent interface{}, name string) *SpriteBatch{
    return &SpriteBatch{self.Object.Call("spriteBatch", parent, name)}
}

// A SpriteBatch is a really fast version of a Phaser Group built solely for speed.
// Use when you need a lot of sprites or particles all sharing the same texture.
// The speed gains are specifically for WebGL. In Canvas mode you won't see any real difference.
func (self *GameObjectFactory) SpriteBatch2O(parent interface{}, name string, addToStage bool) *SpriteBatch{
    return &SpriteBatch{self.Object.Call("spriteBatch", parent, name, addToStage)}
}

// A SpriteBatch is a really fast version of a Phaser Group built solely for speed.
// Use when you need a lot of sprites or particles all sharing the same texture.
// The speed gains are specifically for WebGL. In Canvas mode you won't see any real difference.
func (self *GameObjectFactory) SpriteBatchI(args ...interface{}) *SpriteBatch{
    return &SpriteBatch{self.Object.Call("spriteBatch", args)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) Audio(key string) *Sound{
    return &Sound{self.Object.Call("audio", key)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) Audio1O(key string, volume int) *Sound{
    return &Sound{self.Object.Call("audio", key, volume)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) Audio2O(key string, volume int, loop bool) *Sound{
    return &Sound{self.Object.Call("audio", key, volume, loop)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) Audio3O(key string, volume int, loop bool, connect bool) *Sound{
    return &Sound{self.Object.Call("audio", key, volume, loop, connect)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) AudioI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("audio", args)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) Sound(key string) *Sound{
    return &Sound{self.Object.Call("sound", key)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) Sound1O(key string, volume int) *Sound{
    return &Sound{self.Object.Call("sound", key, volume)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) Sound2O(key string, volume int, loop bool) *Sound{
    return &Sound{self.Object.Call("sound", key, volume, loop)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) Sound3O(key string, volume int, loop bool, connect bool) *Sound{
    return &Sound{self.Object.Call("sound", key, volume, loop, connect)}
}

// Creates a new Sound object.
func (self *GameObjectFactory) SoundI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("sound", args)}
}

// Creates a new AudioSprite object.
func (self *GameObjectFactory) AudioSprite(key string) *AudioSprite{
    return &AudioSprite{self.Object.Call("audioSprite", key)}
}

// Creates a new AudioSprite object.
func (self *GameObjectFactory) AudioSpriteI(args ...interface{}) *AudioSprite{
    return &AudioSprite{self.Object.Call("audioSprite", args)}
}

// Creates a new TileSprite object.
func (self *GameObjectFactory) TileSprite(x int, y int, width int, height int, key interface{}) *TileSprite{
    return &TileSprite{self.Object.Call("tileSprite", x, y, width, height, key)}
}

// Creates a new TileSprite object.
func (self *GameObjectFactory) TileSprite1O(x int, y int, width int, height int, key interface{}, frame interface{}) *TileSprite{
    return &TileSprite{self.Object.Call("tileSprite", x, y, width, height, key, frame)}
}

// Creates a new TileSprite object.
func (self *GameObjectFactory) TileSprite2O(x int, y int, width int, height int, key interface{}, frame interface{}, group *Group) *TileSprite{
    return &TileSprite{self.Object.Call("tileSprite", x, y, width, height, key, frame, group)}
}

// Creates a new TileSprite object.
func (self *GameObjectFactory) TileSpriteI(args ...interface{}) *TileSprite{
    return &TileSprite{self.Object.Call("tileSprite", args)}
}

// Creates a new Rope object.
// 
// Example usage: https://github.com/codevinsky/phaser-rope-demo/blob/master/dist/demo.js
func (self *GameObjectFactory) Rope(x int, y int, key interface{}, frame interface{}, points []interface{}) *Rope{
    return &Rope{self.Object.Call("rope", x, y, key, frame, points)}
}

// Creates a new Rope object.
// 
// Example usage: https://github.com/codevinsky/phaser-rope-demo/blob/master/dist/demo.js
func (self *GameObjectFactory) Rope1O(x int, y int, key interface{}, frame interface{}, points []interface{}, group *Group) *Rope{
    return &Rope{self.Object.Call("rope", x, y, key, frame, points, group)}
}

// Creates a new Rope object.
// 
// Example usage: https://github.com/codevinsky/phaser-rope-demo/blob/master/dist/demo.js
func (self *GameObjectFactory) RopeI(args ...interface{}) *Rope{
    return &Rope{self.Object.Call("rope", args)}
}

// Creates a new Text object.
func (self *GameObjectFactory) Text() *Text{
    return &Text{self.Object.Call("text")}
}

// Creates a new Text object.
func (self *GameObjectFactory) Text1O(x int) *Text{
    return &Text{self.Object.Call("text", x)}
}

// Creates a new Text object.
func (self *GameObjectFactory) Text2O(x int, y int) *Text{
    return &Text{self.Object.Call("text", x, y)}
}

// Creates a new Text object.
func (self *GameObjectFactory) Text3O(x int, y int, text string) *Text{
    return &Text{self.Object.Call("text", x, y, text)}
}

// Creates a new Text object.
func (self *GameObjectFactory) Text4O(x int, y int, text string, style interface{}) *Text{
    return &Text{self.Object.Call("text", x, y, text, style)}
}

// Creates a new Text object.
func (self *GameObjectFactory) Text5O(x int, y int, text string, style interface{}, group *Group) *Text{
    return &Text{self.Object.Call("text", x, y, text, style, group)}
}

// Creates a new Text object.
func (self *GameObjectFactory) TextI(args ...interface{}) *Text{
    return &Text{self.Object.Call("text", args)}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button() *Button{
    return &Button{self.Object.Call("button")}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button1O(x int) *Button{
    return &Button{self.Object.Call("button", x)}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button2O(x int, y int) *Button{
    return &Button{self.Object.Call("button", x, y)}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button3O(x int, y int, key string) *Button{
    return &Button{self.Object.Call("button", x, y, key)}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button4O(x int, y int, key string, callback func(...interface{})) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback)}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button5O(x int, y int, key string, callback func(...interface{}), callbackContext interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext)}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button6O(x int, y int, key string, callback func(...interface{}), callbackContext interface{}, overFrame interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext, overFrame)}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button7O(x int, y int, key string, callback func(...interface{}), callbackContext interface{}, overFrame interface{}, outFrame interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext, overFrame, outFrame)}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button8O(x int, y int, key string, callback func(...interface{}), callbackContext interface{}, overFrame interface{}, outFrame interface{}, downFrame interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext, overFrame, outFrame, downFrame)}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button9O(x int, y int, key string, callback func(...interface{}), callbackContext interface{}, overFrame interface{}, outFrame interface{}, downFrame interface{}, upFrame interface{}) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext, overFrame, outFrame, downFrame, upFrame)}
}

// Creates a new Button object.
func (self *GameObjectFactory) Button10O(x int, y int, key string, callback func(...interface{}), callbackContext interface{}, overFrame interface{}, outFrame interface{}, downFrame interface{}, upFrame interface{}, group *Group) *Button{
    return &Button{self.Object.Call("button", x, y, key, callback, callbackContext, overFrame, outFrame, downFrame, upFrame, group)}
}

// Creates a new Button object.
func (self *GameObjectFactory) ButtonI(args ...interface{}) *Button{
    return &Button{self.Object.Call("button", args)}
}

// Creates a new Graphics object.
func (self *GameObjectFactory) Graphics() *Graphics{
    return &Graphics{self.Object.Call("graphics")}
}

// Creates a new Graphics object.
func (self *GameObjectFactory) Graphics1O(x int) *Graphics{
    return &Graphics{self.Object.Call("graphics", x)}
}

// Creates a new Graphics object.
func (self *GameObjectFactory) Graphics2O(x int, y int) *Graphics{
    return &Graphics{self.Object.Call("graphics", x, y)}
}

// Creates a new Graphics object.
func (self *GameObjectFactory) Graphics3O(x int, y int, group *Group) *Graphics{
    return &Graphics{self.Object.Call("graphics", x, y, group)}
}

// Creates a new Graphics object.
func (self *GameObjectFactory) GraphicsI(args ...interface{}) *Graphics{
    return &Graphics{self.Object.Call("graphics", args)}
}

// Create a new Emitter.
// 
// A particle emitter can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accordingly.
func (self *GameObjectFactory) Emitter() *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("emitter")}
}

// Create a new Emitter.
// 
// A particle emitter can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accordingly.
func (self *GameObjectFactory) Emitter1O(x int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("emitter", x)}
}

// Create a new Emitter.
// 
// A particle emitter can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accordingly.
func (self *GameObjectFactory) Emitter2O(x int, y int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("emitter", x, y)}
}

// Create a new Emitter.
// 
// A particle emitter can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accordingly.
func (self *GameObjectFactory) Emitter3O(x int, y int, maxParticles int) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("emitter", x, y, maxParticles)}
}

// Create a new Emitter.
// 
// A particle emitter can be used for one-time explosions or for
// continuous effects like rain and fire. All it really does is launch Particle objects out
// at set intervals, and fixes their positions and velocities accordingly.
func (self *GameObjectFactory) EmitterI(args ...interface{}) *ParticlesArcadeEmitter{
    return &ParticlesArcadeEmitter{self.Object.Call("emitter", args)}
}

// Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectFactory) RetroFont(font string, characterWidth int, characterHeight int, chars string, charsPerRow int) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", font, characterWidth, characterHeight, chars, charsPerRow)}
}

// Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectFactory) RetroFont1O(font string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", font, characterWidth, characterHeight, chars, charsPerRow, xSpacing)}
}

// Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectFactory) RetroFont2O(font string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int, ySpacing int) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", font, characterWidth, characterHeight, chars, charsPerRow, xSpacing, ySpacing)}
}

// Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectFactory) RetroFont3O(font string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int, ySpacing int, xOffset int) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", font, characterWidth, characterHeight, chars, charsPerRow, xSpacing, ySpacing, xOffset)}
}

// Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectFactory) RetroFont4O(font string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int, ySpacing int, xOffset int, yOffset int) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", font, characterWidth, characterHeight, chars, charsPerRow, xSpacing, ySpacing, xOffset, yOffset)}
}

// Create a new RetroFont object.
// 
// A RetroFont can be used as a texture for an Image or Sprite and optionally add it to the Cache.
// A RetroFont uses a bitmap which contains fixed with characters for the font set. You use character spacing to define the set.
// If you need variable width character support then use a BitmapText object instead. The main difference between a RetroFont and a BitmapText
// is that a RetroFont creates a single texture that you can apply to a game object, where-as a BitmapText creates one Sprite object per letter of text.
// The texture can be asssigned or one or multiple images/sprites, but note that the text the RetroFont uses will be shared across them all,
// i.e. if you need each Image to have different text in it, then you need to create multiple RetroFont objects.
func (self *GameObjectFactory) RetroFontI(args ...interface{}) *RetroFont{
    return &RetroFont{self.Object.Call("retroFont", args)}
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
func (self *GameObjectFactory) BitmapText(x int, y int, font string) *BitmapText{
    return &BitmapText{self.Object.Call("bitmapText", x, y, font)}
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
func (self *GameObjectFactory) BitmapText1O(x int, y int, font string, text string) *BitmapText{
    return &BitmapText{self.Object.Call("bitmapText", x, y, font, text)}
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
func (self *GameObjectFactory) BitmapText2O(x int, y int, font string, text string, size int) *BitmapText{
    return &BitmapText{self.Object.Call("bitmapText", x, y, font, text, size)}
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
func (self *GameObjectFactory) BitmapText3O(x int, y int, font string, text string, size int, group *Group) *BitmapText{
    return &BitmapText{self.Object.Call("bitmapText", x, y, font, text, size, group)}
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
func (self *GameObjectFactory) BitmapTextI(args ...interface{}) *BitmapText{
    return &BitmapText{self.Object.Call("bitmapText", args)}
}

// Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectFactory) Tilemap() *Tilemap{
    return &Tilemap{self.Object.Call("tilemap")}
}

// Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectFactory) Tilemap1O(key string) *Tilemap{
    return &Tilemap{self.Object.Call("tilemap", key)}
}

// Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectFactory) Tilemap2O(key string, tileWidth int) *Tilemap{
    return &Tilemap{self.Object.Call("tilemap", key, tileWidth)}
}

// Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectFactory) Tilemap3O(key string, tileWidth int, tileHeight int) *Tilemap{
    return &Tilemap{self.Object.Call("tilemap", key, tileWidth, tileHeight)}
}

// Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectFactory) Tilemap4O(key string, tileWidth int, tileHeight int, width int) *Tilemap{
    return &Tilemap{self.Object.Call("tilemap", key, tileWidth, tileHeight, width)}
}

// Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectFactory) Tilemap5O(key string, tileWidth int, tileHeight int, width int, height int) *Tilemap{
    return &Tilemap{self.Object.Call("tilemap", key, tileWidth, tileHeight, width, height)}
}

// Creates a new Phaser.Tilemap object.
// 
// The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
func (self *GameObjectFactory) TilemapI(args ...interface{}) *Tilemap{
    return &Tilemap{self.Object.Call("tilemap", args)}
}

// A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectFactory) RenderTexture() *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture")}
}

// A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectFactory) RenderTexture1O(width int) *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture", width)}
}

// A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectFactory) RenderTexture2O(width int, height int) *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture", width, height)}
}

// A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectFactory) RenderTexture3O(width int, height int, key string) *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture", width, height, key)}
}

// A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectFactory) RenderTexture4O(width int, height int, key string, addToCache bool) *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture", width, height, key, addToCache)}
}

// A dynamic initially blank canvas to which images can be drawn.
func (self *GameObjectFactory) RenderTextureI(args ...interface{}) *RenderTexture{
    return &RenderTexture{self.Object.Call("renderTexture", args)}
}

// Create a Video object.
// 
// This will return a Phaser.Video object which you can pass to a Sprite to be used as a texture.
func (self *GameObjectFactory) Video() *Video{
    return &Video{self.Object.Call("video")}
}

// Create a Video object.
// 
// This will return a Phaser.Video object which you can pass to a Sprite to be used as a texture.
func (self *GameObjectFactory) Video1O(key interface{}) *Video{
    return &Video{self.Object.Call("video", key)}
}

// Create a Video object.
// 
// This will return a Phaser.Video object which you can pass to a Sprite to be used as a texture.
func (self *GameObjectFactory) Video2O(key interface{}, url interface{}) *Video{
    return &Video{self.Object.Call("video", key, url)}
}

// Create a Video object.
// 
// This will return a Phaser.Video object which you can pass to a Sprite to be used as a texture.
func (self *GameObjectFactory) VideoI(args ...interface{}) *Video{
    return &Video{self.Object.Call("video", args)}
}

// Create a BitmapData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectFactory) BitmapData() *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData")}
}

// Create a BitmapData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectFactory) BitmapData1O(width int) *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData", width)}
}

// Create a BitmapData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectFactory) BitmapData2O(width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData", width, height)}
}

// Create a BitmapData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectFactory) BitmapData3O(width int, height int, key string) *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData", width, height, key)}
}

// Create a BitmapData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectFactory) BitmapData4O(width int, height int, key string, addToCache bool) *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData", width, height, key, addToCache)}
}

// Create a BitmapData object.
// 
// A BitmapData object can be manipulated and drawn to like a traditional Canvas object and used to texture Sprites.
func (self *GameObjectFactory) BitmapDataI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("bitmapData", args)}
}

// A WebGL shader/filter that can be applied to Sprites.
func (self *GameObjectFactory) Filter(filter string, args interface{}) *Filter{
    return &Filter{self.Object.Call("filter", filter, args)}
}

// A WebGL shader/filter that can be applied to Sprites.
func (self *GameObjectFactory) FilterI(args ...interface{}) *Filter{
    return &Filter{self.Object.Call("filter", args)}
}

// Add a new Plugin into the PluginManager.
// 
// The Plugin must have 2 properties: `game` and `parent`. Plugin.game is set to the game reference the PluginManager uses, and parent is set to the PluginManager.
func (self *GameObjectFactory) Plugin(plugin interface{}, parameter interface{}) *Plugin{
    return &Plugin{self.Object.Call("plugin", plugin, parameter)}
}

// Add a new Plugin into the PluginManager.
// 
// The Plugin must have 2 properties: `game` and `parent`. Plugin.game is set to the game reference the PluginManager uses, and parent is set to the PluginManager.
func (self *GameObjectFactory) PluginI(args ...interface{}) *Plugin{
    return &Plugin{self.Object.Call("plugin", args)}
}
