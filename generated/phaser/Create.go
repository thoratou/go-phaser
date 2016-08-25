// Package phaser Automatic generation for Phaser.Create
// generated file Create.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// Create The Phaser.Create class is a collection of smaller helper methods that allow you to generate game content
// quickly and easily, without the need for any external files. You can create textures for sprites and in
// coming releases we'll add dynamic sound effect generation support as well (like sfxr).
// 
// Access this via `Game.create` (`this.game.create` from within a State object)
type Create struct {
    *js.Object
}

// NewCreate The Phaser.Create class is a collection of smaller helper methods that allow you to generate game content
// quickly and easily, without the need for any external files. You can create textures for sprites and in
// coming releases we'll add dynamic sound effect generation support as well (like sfxr).
// 
// Access this via `Game.create` (`this.game.create` from within a State object)
func NewCreate(game *Game) *Create {
    return &Create{js.Global.Get("Phaser").Get("Create").New(game)}
}
// NewCreateI The Phaser.Create class is a collection of smaller helper methods that allow you to generate game content
// quickly and easily, without the need for any external files. You can create textures for sprites and in
// coming releases we'll add dynamic sound effect generation support as well (like sfxr).
// 
// Access this via `Game.create` (`this.game.create` from within a State object)
func NewCreateI(args ...interface{}) *Create {
    return &Create{js.Global.Get("Phaser").Get("Create").New(args)}
}



// Create Binding conversion method to Create point 
func ToCreate(jsStruct interface{}) *Create {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Create{Object: object}
	}
	return nil
}



// Game A reference to the currently running Game.
func (self *Create) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Create) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Bmd The internal BitmapData Create uses to generate textures from.
func (self *Create) Bmd() *BitmapData{
    return &BitmapData{self.Object.Get("bmd")}
}

// SetBmdA The internal BitmapData Create uses to generate textures from.
func (self *Create) SetBmdA(member *BitmapData) {
    self.Object.Set("bmd", member)
}

// Canvas The canvas the BitmapData uses.
func (self *Create) Canvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("canvas"))
}

// SetCanvasA The canvas the BitmapData uses.
func (self *Create) SetCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("canvas", member)
}

// Ctx The 2d context of the canvas.
func (self *Create) Ctx() interface{}{
    return self.Object.Get("ctx")
}

// SetCtxA The 2d context of the canvas.
func (self *Create) SetCtxA(member interface{}) {
    self.Object.Set("ctx", member)
}

// Palettes A range of 16 color palettes for use with sprite generation.
func (self *Create) Palettes() []interface{}{
	array00 := self.Object.Get("palettes")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetPalettesA A range of 16 color palettes for use with sprite generation.
func (self *Create) SetPalettesA(member []interface{}) {
    self.Object.Set("palettes", member)
}

// PALETTE_ARNE A 16 color palette by [Arne](http://androidarts.com/palette/16pal.htm)
func (self *Create) PALETTE_ARNE() int{
    return self.Object.Get("PALETTE_ARNE").Int()
}

// SetPALETTE_ARNEA A 16 color palette by [Arne](http://androidarts.com/palette/16pal.htm)
func (self *Create) SetPALETTE_ARNEA(member int) {
    self.Object.Set("PALETTE_ARNE", member)
}

// PALETTE_JMP A 16 color JMP inspired palette.
func (self *Create) PALETTE_JMP() int{
    return self.Object.Get("PALETTE_JMP").Int()
}

// SetPALETTE_JMPA A 16 color JMP inspired palette.
func (self *Create) SetPALETTE_JMPA(member int) {
    self.Object.Set("PALETTE_JMP", member)
}

// PALETTE_CGA A 16 color CGA inspired palette.
func (self *Create) PALETTE_CGA() int{
    return self.Object.Get("PALETTE_CGA").Int()
}

// SetPALETTE_CGAA A 16 color CGA inspired palette.
func (self *Create) SetPALETTE_CGAA(member int) {
    self.Object.Set("PALETTE_CGA", member)
}

// PALETTE_C64 A 16 color C64 inspired palette.
func (self *Create) PALETTE_C64() int{
    return self.Object.Get("PALETTE_C64").Int()
}

// SetPALETTE_C64A A 16 color C64 inspired palette.
func (self *Create) SetPALETTE_C64A(member int) {
    self.Object.Set("PALETTE_C64", member)
}

// PALETTE_JAPANESE_MACHINE A 16 color palette inspired by Japanese computers like the MSX.
func (self *Create) PALETTE_JAPANESE_MACHINE() int{
    return self.Object.Get("PALETTE_JAPANESE_MACHINE").Int()
}

// SetPALETTE_JAPANESE_MACHINEA A 16 color palette inspired by Japanese computers like the MSX.
func (self *Create) SetPALETTE_JAPANESE_MACHINEA(member int) {
    self.Object.Set("PALETTE_JAPANESE_MACHINE", member)
}


// Texture Generates a new PIXI.Texture from the given data, which can be applied to a Sprite.
// 
// This allows you to create game graphics quickly and easily, with no external files but that use actual proper images
// rather than Phaser.Graphics objects, which are expensive to render and limited in scope.
// 
// Each element of the array is a string holding the pixel color values, as mapped to one of the Phaser.Create PALETTE consts.
// 
// For example:
// 
// `var data = [
//   ' 333 ',
//   ' 777 ',
//   'E333E',
//   ' 333 ',
//   ' 3 3 '
// ];`
// 
// `game.create.texture('bob', data);`
// 
// The above will create a new texture called `bob`, which will look like a little man wearing a hat. You can then use it
// for sprites the same way you use any other texture: `game.add.sprite(0, 0, 'bob');`
func (self *Create) Texture(key string, data []interface{}) *Texture{
    return &Texture{self.Object.Call("texture", key, data)}
}

// Texture1O Generates a new PIXI.Texture from the given data, which can be applied to a Sprite.
// 
// This allows you to create game graphics quickly and easily, with no external files but that use actual proper images
// rather than Phaser.Graphics objects, which are expensive to render and limited in scope.
// 
// Each element of the array is a string holding the pixel color values, as mapped to one of the Phaser.Create PALETTE consts.
// 
// For example:
// 
// `var data = [
//   ' 333 ',
//   ' 777 ',
//   'E333E',
//   ' 333 ',
//   ' 3 3 '
// ];`
// 
// `game.create.texture('bob', data);`
// 
// The above will create a new texture called `bob`, which will look like a little man wearing a hat. You can then use it
// for sprites the same way you use any other texture: `game.add.sprite(0, 0, 'bob');`
func (self *Create) Texture1O(key string, data []interface{}, pixelWidth int) *Texture{
    return &Texture{self.Object.Call("texture", key, data, pixelWidth)}
}

// Texture2O Generates a new PIXI.Texture from the given data, which can be applied to a Sprite.
// 
// This allows you to create game graphics quickly and easily, with no external files but that use actual proper images
// rather than Phaser.Graphics objects, which are expensive to render and limited in scope.
// 
// Each element of the array is a string holding the pixel color values, as mapped to one of the Phaser.Create PALETTE consts.
// 
// For example:
// 
// `var data = [
//   ' 333 ',
//   ' 777 ',
//   'E333E',
//   ' 333 ',
//   ' 3 3 '
// ];`
// 
// `game.create.texture('bob', data);`
// 
// The above will create a new texture called `bob`, which will look like a little man wearing a hat. You can then use it
// for sprites the same way you use any other texture: `game.add.sprite(0, 0, 'bob');`
func (self *Create) Texture2O(key string, data []interface{}, pixelWidth int, pixelHeight int) *Texture{
    return &Texture{self.Object.Call("texture", key, data, pixelWidth, pixelHeight)}
}

// Texture3O Generates a new PIXI.Texture from the given data, which can be applied to a Sprite.
// 
// This allows you to create game graphics quickly and easily, with no external files but that use actual proper images
// rather than Phaser.Graphics objects, which are expensive to render and limited in scope.
// 
// Each element of the array is a string holding the pixel color values, as mapped to one of the Phaser.Create PALETTE consts.
// 
// For example:
// 
// `var data = [
//   ' 333 ',
//   ' 777 ',
//   'E333E',
//   ' 333 ',
//   ' 3 3 '
// ];`
// 
// `game.create.texture('bob', data);`
// 
// The above will create a new texture called `bob`, which will look like a little man wearing a hat. You can then use it
// for sprites the same way you use any other texture: `game.add.sprite(0, 0, 'bob');`
func (self *Create) Texture3O(key string, data []interface{}, pixelWidth int, pixelHeight int, palette int) *Texture{
    return &Texture{self.Object.Call("texture", key, data, pixelWidth, pixelHeight, palette)}
}

// TextureI Generates a new PIXI.Texture from the given data, which can be applied to a Sprite.
// 
// This allows you to create game graphics quickly and easily, with no external files but that use actual proper images
// rather than Phaser.Graphics objects, which are expensive to render and limited in scope.
// 
// Each element of the array is a string holding the pixel color values, as mapped to one of the Phaser.Create PALETTE consts.
// 
// For example:
// 
// `var data = [
//   ' 333 ',
//   ' 777 ',
//   'E333E',
//   ' 333 ',
//   ' 3 3 '
// ];`
// 
// `game.create.texture('bob', data);`
// 
// The above will create a new texture called `bob`, which will look like a little man wearing a hat. You can then use it
// for sprites the same way you use any other texture: `game.add.sprite(0, 0, 'bob');`
func (self *Create) TextureI(args ...interface{}) *Texture{
    return &Texture{self.Object.Call("texture", args)}
}

// Grid Creates a grid texture based on the given dimensions.
func (self *Create) Grid(key string, width int, height int, cellWidth int, cellHeight int, color string) *Texture{
    return &Texture{self.Object.Call("grid", key, width, height, cellWidth, cellHeight, color)}
}

// GridI Creates a grid texture based on the given dimensions.
func (self *Create) GridI(args ...interface{}) *Texture{
    return &Texture{self.Object.Call("grid", args)}
}

