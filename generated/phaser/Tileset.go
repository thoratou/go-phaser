// Package phaser Automatic generation for Phaser.Tileset
// generated file Tileset.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// Tileset A Tile set is a combination of an image containing the tiles and collision data per tile.
// 
// Tilesets are normally created automatically when Tiled data is loaded.
type Tileset struct {
    *js.Object
}

// NewTileset A Tile set is a combination of an image containing the tiles and collision data per tile.
// 
// Tilesets are normally created automatically when Tiled data is loaded.
func NewTileset(name string, firstgid int) *Tileset {
    return &Tileset{js.Global.Get("Phaser").Get("Tileset").New(name, firstgid)}
}
// NewTileset1O A Tile set is a combination of an image containing the tiles and collision data per tile.
// 
// Tilesets are normally created automatically when Tiled data is loaded.
func NewTileset1O(name string, firstgid int, width int) *Tileset {
    return &Tileset{js.Global.Get("Phaser").Get("Tileset").New(name, firstgid, width)}
}
// NewTileset2O A Tile set is a combination of an image containing the tiles and collision data per tile.
// 
// Tilesets are normally created automatically when Tiled data is loaded.
func NewTileset2O(name string, firstgid int, width int, height int) *Tileset {
    return &Tileset{js.Global.Get("Phaser").Get("Tileset").New(name, firstgid, width, height)}
}
// NewTileset3O A Tile set is a combination of an image containing the tiles and collision data per tile.
// 
// Tilesets are normally created automatically when Tiled data is loaded.
func NewTileset3O(name string, firstgid int, width int, height int, margin int) *Tileset {
    return &Tileset{js.Global.Get("Phaser").Get("Tileset").New(name, firstgid, width, height, margin)}
}
// NewTileset4O A Tile set is a combination of an image containing the tiles and collision data per tile.
// 
// Tilesets are normally created automatically when Tiled data is loaded.
func NewTileset4O(name string, firstgid int, width int, height int, margin int, spacing int) *Tileset {
    return &Tileset{js.Global.Get("Phaser").Get("Tileset").New(name, firstgid, width, height, margin, spacing)}
}
// NewTileset5O A Tile set is a combination of an image containing the tiles and collision data per tile.
// 
// Tilesets are normally created automatically when Tiled data is loaded.
func NewTileset5O(name string, firstgid int, width int, height int, margin int, spacing int, properties interface{}) *Tileset {
    return &Tileset{js.Global.Get("Phaser").Get("Tileset").New(name, firstgid, width, height, margin, spacing, properties)}
}
// NewTilesetI A Tile set is a combination of an image containing the tiles and collision data per tile.
// 
// Tilesets are normally created automatically when Tiled data is loaded.
func NewTilesetI(args ...interface{}) *Tileset {
    return &Tileset{js.Global.Get("Phaser").Get("Tileset").New(args)}
}



// Tileset Binding conversion method to Tileset point 
func ToTileset(jsStruct interface{}) *Tileset {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Tileset{Object: object}
	}
	return nil
}



// Name The name of the Tileset.
func (self *Tileset) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA The name of the Tileset.
func (self *Tileset) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Firstgid The Tiled firstgid value.
// This is the starting index of the first tile index this Tileset contains.
func (self *Tileset) Firstgid() int{
    return self.Object.Get("firstgid").Int()
}

// SetFirstgidA The Tiled firstgid value.
// This is the starting index of the first tile index this Tileset contains.
func (self *Tileset) SetFirstgidA(member int) {
    self.Object.Set("firstgid", member)
}

// TileWidth The width of each tile (in pixels).
func (self *Tileset) TileWidth() int{
    return self.Object.Get("tileWidth").Int()
}

// SetTileWidthA The width of each tile (in pixels).
func (self *Tileset) SetTileWidthA(member int) {
    self.Object.Set("tileWidth", member)
}

// TileHeight The height of each tile (in pixels).
func (self *Tileset) TileHeight() int{
    return self.Object.Get("tileHeight").Int()
}

// SetTileHeightA The height of each tile (in pixels).
func (self *Tileset) SetTileHeightA(member int) {
    self.Object.Set("tileHeight", member)
}

// TileMargin The margin around the tiles in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) TileMargin() interface{}{
    return self.Object.Get("tileMargin")
}

// SetTileMarginA The margin around the tiles in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) SetTileMarginA(member interface{}) {
    self.Object.Set("tileMargin", member)
}

// TileSpacing The spacing between each tile in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) TileSpacing() int{
    return self.Object.Get("tileSpacing").Int()
}

// SetTileSpacingA The spacing between each tile in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) SetTileSpacingA(member int) {
    self.Object.Set("tileSpacing", member)
}

// Properties Tileset-specific properties that are typically defined in the Tiled editor.
func (self *Tileset) Properties() interface{}{
    return self.Object.Get("properties")
}

// SetPropertiesA Tileset-specific properties that are typically defined in the Tiled editor.
func (self *Tileset) SetPropertiesA(member interface{}) {
    self.Object.Set("properties", member)
}

// Image The cached image that contains the individual tiles. Use {@link Phaser.Tileset.setImage setImage} to set.
func (self *Tileset) Image() interface{}{
    return self.Object.Get("image")
}

// SetImageA The cached image that contains the individual tiles. Use {@link Phaser.Tileset.setImage setImage} to set.
func (self *Tileset) SetImageA(member interface{}) {
    self.Object.Set("image", member)
}

// Rows The number of tile rows in the the tileset.
func (self *Tileset) Rows() interface{}{
    return self.Object.Get("rows")
}

// SetRowsA The number of tile rows in the the tileset.
func (self *Tileset) SetRowsA(member interface{}) {
    self.Object.Set("rows", member)
}

// Columns The number of tile columns in the tileset.
func (self *Tileset) Columns() int{
    return self.Object.Get("columns").Int()
}

// SetColumnsA The number of tile columns in the tileset.
func (self *Tileset) SetColumnsA(member int) {
    self.Object.Set("columns", member)
}

// Total The total number of tiles in the tileset.
func (self *Tileset) Total() int{
    return self.Object.Get("total").Int()
}

// SetTotalA The total number of tiles in the tileset.
func (self *Tileset) SetTotalA(member int) {
    self.Object.Set("total", member)
}


// Draw Draws a tile from this Tileset at the given coordinates on the context.
func (self *Tileset) Draw(context *dom.CanvasRenderingContext2D, x int, y int, index int) {
    self.Object.Call("draw", context, x, y, index)
}

// DrawI Draws a tile from this Tileset at the given coordinates on the context.
func (self *Tileset) DrawI(args ...interface{}) {
    self.Object.Call("draw", args)
}

// ContainsTileIndex Returns true if and only if this tileset contains the given tile index.
func (self *Tileset) ContainsTileIndex() bool{
    return self.Object.Call("containsTileIndex").Bool()
}

// ContainsTileIndexI Returns true if and only if this tileset contains the given tile index.
func (self *Tileset) ContainsTileIndexI(args ...interface{}) bool{
    return self.Object.Call("containsTileIndex", args).Bool()
}

// SetImage Set the image associated with this Tileset and update the tile data.
func (self *Tileset) SetImage(image *Image) {
    self.Object.Call("setImage", image)
}

// SetImageI Set the image associated with this Tileset and update the tile data.
func (self *Tileset) SetImageI(args ...interface{}) {
    self.Object.Call("setImage", args)
}

// SetSpacing Sets tile spacing and margins.
func (self *Tileset) SetSpacing() {
    self.Object.Call("setSpacing")
}

// SetSpacing1O Sets tile spacing and margins.
func (self *Tileset) SetSpacing1O(margin int) {
    self.Object.Call("setSpacing", margin)
}

// SetSpacing2O Sets tile spacing and margins.
func (self *Tileset) SetSpacing2O(margin int, spacing int) {
    self.Object.Call("setSpacing", margin, spacing)
}

// SetSpacingI Sets tile spacing and margins.
func (self *Tileset) SetSpacingI(args ...interface{}) {
    self.Object.Call("setSpacing", args)
}

// UpdateTileData Updates tile coordinates and tileset data.
func (self *Tileset) UpdateTileData(imageWidth int, imageHeight int) {
    self.Object.Call("updateTileData", imageWidth, imageHeight)
}

// UpdateTileDataI Updates tile coordinates and tileset data.
func (self *Tileset) UpdateTileDataI(args ...interface{}) {
    self.Object.Call("updateTileData", args)
}

