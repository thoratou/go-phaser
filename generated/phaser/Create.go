// Automatic generation for Phaser.Create
// generated file Create.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// The Phaser.Create class is a collection of smaller helper methods that allow you to generate game content
// quickly and easily, without the need for any external files. You can create textures for sprites and in
// coming releases we'll add dynamic sound effect generation support as well (like sfxr).
// 
// Access this via `Game.create` (`this.game.create` from within a State object)
type Create struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *Create) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Create) SetGame(member *Game) {
    self.Set("game", member)
}

// The internal BitmapData Create uses to generate textures from.
func (self *Create) GetBmd() *BitmapData{
    return &BitmapData{self.Get("bmd")}
}

// The internal BitmapData Create uses to generate textures from.
func (self *Create) SetBmd(member *BitmapData) {
    self.Set("bmd", member)
}

// The canvas the BitmapData uses.
func (self *Create) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("canvas"))
}

// The canvas the BitmapData uses.
func (self *Create) SetCanvas(member dom.HTMLCanvasElement) {
    self.Set("canvas", member)
}

// The 2d context of the canvas.
func (self *Create) GetCtx() interface{}{
    return self.Get("ctx")
}

// The 2d context of the canvas.
func (self *Create) SetCtx(member interface{}) {
    self.Set("ctx", member)
}

// A range of 16 color palettes for use with sprite generation.
func (self *Create) GetPalettes() []interface{}{
	array := self.Get("palettes")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// A range of 16 color palettes for use with sprite generation.
func (self *Create) SetPalettes(member []interface{}) {
    self.Set("palettes", member)
}

// A 16 color palette by [Arne](http://androidarts.com/palette/16pal.htm)
func (self *Create) GetPALETTE_ARNE() float64{
    return self.Get("PALETTE_ARNE").Float()
}

// A 16 color palette by [Arne](http://androidarts.com/palette/16pal.htm)
func (self *Create) SetPALETTE_ARNE(member float64) {
    self.Set("PALETTE_ARNE", member)
}

// A 16 color JMP inspired palette.
func (self *Create) GetPALETTE_JMP() float64{
    return self.Get("PALETTE_JMP").Float()
}

// A 16 color JMP inspired palette.
func (self *Create) SetPALETTE_JMP(member float64) {
    self.Set("PALETTE_JMP", member)
}

// A 16 color CGA inspired palette.
func (self *Create) GetPALETTE_CGA() float64{
    return self.Get("PALETTE_CGA").Float()
}

// A 16 color CGA inspired palette.
func (self *Create) SetPALETTE_CGA(member float64) {
    self.Set("PALETTE_CGA", member)
}

// A 16 color C64 inspired palette.
func (self *Create) GetPALETTE_C64() float64{
    return self.Get("PALETTE_C64").Float()
}

// A 16 color C64 inspired palette.
func (self *Create) SetPALETTE_C64(member float64) {
    self.Set("PALETTE_C64", member)
}

// A 16 color palette inspired by Japanese computers like the MSX.
func (self *Create) GetPALETTE_JAPANESE_MACHINE() float64{
    return self.Get("PALETTE_JAPANESE_MACHINE").Float()
}

// A 16 color palette inspired by Japanese computers like the MSX.
func (self *Create) SetPALETTE_JAPANESE_MACHINE(member float64) {
    self.Set("PALETTE_JAPANESE_MACHINE", member)
}



// Generates a new PIXI.Texture from the given data, which can be applied to a Sprite.
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
    return &Texture{self.Call("texture", args)}
}

// Creates a grid texture based on the given dimensions.
func (self *Create) GridI(args ...interface{}) *Texture{
    return &Texture{self.Call("grid", args)}
}
