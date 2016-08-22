// Automatic generation for Phaser.Tileset
// generated file Tileset.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// A Tile set is a combination of an image containing the tiles and collision data per tile.
// 
// Tilesets are normally created automatically when Tiled data is loaded.
type Tileset struct {
    *js.Object
}


// The name of the Tileset.
func (self *Tileset) GetNameA() string{
    return self.Object.Get("name").String()
}

// The name of the Tileset.
func (self *Tileset) SetNameA(member string) {
    self.Object.Set("name", member)
}

// The Tiled firstgid value.
// This is the starting index of the first tile index this Tileset contains.
func (self *Tileset) GetFirstgidA() int{
    return self.Object.Get("firstgid").Int()
}

// The Tiled firstgid value.
// This is the starting index of the first tile index this Tileset contains.
func (self *Tileset) SetFirstgidA(member int) {
    self.Object.Set("firstgid", member)
}

// The width of each tile (in pixels).
func (self *Tileset) GetTileWidthA() int{
    return self.Object.Get("tileWidth").Int()
}

// The width of each tile (in pixels).
func (self *Tileset) SetTileWidthA(member int) {
    self.Object.Set("tileWidth", member)
}

// The height of each tile (in pixels).
func (self *Tileset) GetTileHeightA() int{
    return self.Object.Get("tileHeight").Int()
}

// The height of each tile (in pixels).
func (self *Tileset) SetTileHeightA(member int) {
    self.Object.Set("tileHeight", member)
}

// The margin around the tiles in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) GetTileMarginA() interface{}{
    return self.Object.Get("tileMargin")
}

// The margin around the tiles in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) SetTileMarginA(member interface{}) {
    self.Object.Set("tileMargin", member)
}

// The spacing between each tile in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) GetTileSpacingA() int{
    return self.Object.Get("tileSpacing").Int()
}

// The spacing between each tile in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) SetTileSpacingA(member int) {
    self.Object.Set("tileSpacing", member)
}

// Tileset-specific properties that are typically defined in the Tiled editor.
func (self *Tileset) GetPropertiesA() interface{}{
    return self.Object.Get("properties")
}

// Tileset-specific properties that are typically defined in the Tiled editor.
func (self *Tileset) SetPropertiesA(member interface{}) {
    self.Object.Set("properties", member)
}

// The cached image that contains the individual tiles. Use {@link Phaser.Tileset.setImage setImage} to set.
func (self *Tileset) GetImageA() interface{}{
    return self.Object.Get("image")
}

// The cached image that contains the individual tiles. Use {@link Phaser.Tileset.setImage setImage} to set.
func (self *Tileset) SetImageA(member interface{}) {
    self.Object.Set("image", member)
}

// The number of tile rows in the the tileset.
func (self *Tileset) GetRowsA() interface{}{
    return self.Object.Get("rows")
}

// The number of tile rows in the the tileset.
func (self *Tileset) SetRowsA(member interface{}) {
    self.Object.Set("rows", member)
}

// The number of tile columns in the tileset.
func (self *Tileset) GetColumnsA() int{
    return self.Object.Get("columns").Int()
}

// The number of tile columns in the tileset.
func (self *Tileset) SetColumnsA(member int) {
    self.Object.Set("columns", member)
}

// The total number of tiles in the tileset.
func (self *Tileset) GetTotalA() int{
    return self.Object.Get("total").Int()
}

// The total number of tiles in the tileset.
func (self *Tileset) SetTotalA(member int) {
    self.Object.Set("total", member)
}



// Draws a tile from this Tileset at the given coordinates on the context.
func (self *Tileset) Draw(context *dom.CanvasRenderingContext2D, x int, y int, index int) {
    self.Object.Call("draw", context, x, y, index)
}

// Draws a tile from this Tileset at the given coordinates on the context.
func (self *Tileset) DrawI(args ...interface{}) {
    self.Object.Call("draw", args)
}

// Returns true if and only if this tileset contains the given tile index.
func (self *Tileset) ContainsTileIndex() bool{
    return self.Object.Call("containsTileIndex").Bool()
}

// Returns true if and only if this tileset contains the given tile index.
func (self *Tileset) ContainsTileIndexI(args ...interface{}) bool{
    return self.Object.Call("containsTileIndex", args).Bool()
}

// Set the image associated with this Tileset and update the tile data.
func (self *Tileset) SetImage(image *Image) {
    self.Object.Call("setImage", image)
}

// Set the image associated with this Tileset and update the tile data.
func (self *Tileset) SetImageI(args ...interface{}) {
    self.Object.Call("setImage", args)
}

// Sets tile spacing and margins.
func (self *Tileset) SetSpacing() {
    self.Object.Call("setSpacing")
}

// Sets tile spacing and margins.
func (self *Tileset) SetSpacing1O(margin int) {
    self.Object.Call("setSpacing", margin)
}

// Sets tile spacing and margins.
func (self *Tileset) SetSpacing2O(margin int, spacing int) {
    self.Object.Call("setSpacing", margin, spacing)
}

// Sets tile spacing and margins.
func (self *Tileset) SetSpacingI(args ...interface{}) {
    self.Object.Call("setSpacing", args)
}

// Updates tile coordinates and tileset data.
func (self *Tileset) UpdateTileData(imageWidth int, imageHeight int) {
    self.Object.Call("updateTileData", imageWidth, imageHeight)
}

// Updates tile coordinates and tileset data.
func (self *Tileset) UpdateTileDataI(args ...interface{}) {
    self.Object.Call("updateTileData", args)
}
