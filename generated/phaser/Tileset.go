// Automatic generation for Phaser.Tileset
// generated file Tileset.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A Tile set is a combination of an image containing the tiles and collision data per tile.
// 
// Tilesets are normally created automatically when Tiled data is loaded.
type Tileset struct {
    *js.Object
}


// The name of the Tileset.
func (self *Tileset) GetName() string{
    return self.Get("name").String()
}

// The name of the Tileset.
func (self *Tileset) SetName(member string) {
    self.Set("name", member)
}

// The Tiled firstgid value.
// This is the starting index of the first tile index this Tileset contains.
func (self *Tileset) GetFirstgid() int{
    return self.Get("firstgid").Int()
}

// The Tiled firstgid value.
// This is the starting index of the first tile index this Tileset contains.
func (self *Tileset) SetFirstgid(member int) {
    self.Set("firstgid", member)
}

// The width of each tile (in pixels).
func (self *Tileset) GetTileWidth() int{
    return self.Get("tileWidth").Int()
}

// The width of each tile (in pixels).
func (self *Tileset) SetTileWidth(member int) {
    self.Set("tileWidth", member)
}

// The height of each tile (in pixels).
func (self *Tileset) GetTileHeight() int{
    return self.Get("tileHeight").Int()
}

// The height of each tile (in pixels).
func (self *Tileset) SetTileHeight(member int) {
    self.Set("tileHeight", member)
}

// The margin around the tiles in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) GetTileMargin() interface{}{
    return self.Get("tileMargin")
}

// The margin around the tiles in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) SetTileMargin(member interface{}) {
    self.Set("tileMargin", member)
}

// The spacing between each tile in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) GetTileSpacing() int{
    return self.Get("tileSpacing").Int()
}

// The spacing between each tile in the sheet (in pixels).
// Use `setSpacing` to change.
func (self *Tileset) SetTileSpacing(member int) {
    self.Set("tileSpacing", member)
}

// Tileset-specific properties that are typically defined in the Tiled editor.
func (self *Tileset) GetProperties() interface{}{
    return self.Get("properties")
}

// Tileset-specific properties that are typically defined in the Tiled editor.
func (self *Tileset) SetProperties(member interface{}) {
    self.Set("properties", member)
}

// The cached image that contains the individual tiles. Use {@link Phaser.Tileset.setImage setImage} to set.
func (self *Tileset) GetImage() interface{}{
    return self.Get("image")
}

// The cached image that contains the individual tiles. Use {@link Phaser.Tileset.setImage setImage} to set.
func (self *Tileset) SetImage(member interface{}) {
    self.Set("image", member)
}

// The number of tile rows in the the tileset.
func (self *Tileset) GetRows() interface{}{
    return self.Get("rows")
}

// The number of tile rows in the the tileset.
func (self *Tileset) SetRows(member interface{}) {
    self.Set("rows", member)
}

// The number of tile columns in the tileset.
func (self *Tileset) GetColumns() int{
    return self.Get("columns").Int()
}

// The number of tile columns in the tileset.
func (self *Tileset) SetColumns(member int) {
    self.Set("columns", member)
}

// The total number of tiles in the tileset.
func (self *Tileset) GetTotal() int{
    return self.Get("total").Int()
}

// The total number of tiles in the tileset.
func (self *Tileset) SetTotal(member int) {
    self.Set("total", member)
}



// Draws a tile from this Tileset at the given coordinates on the context.
func (self *Tileset) DrawI(args ...interface{}) {
    self.Call("draw", args)
}

// Returns true if and only if this tileset contains the given tile index.
func (self *Tileset) ContainsTileIndexI(args ...interface{}) bool{
    return self.Call("containsTileIndex", args).Bool()
}

// Set the image associated with this Tileset and update the tile data.
func (self *Tileset) SetImageI(args ...interface{}) {
    self.Call("setImage", args)
}

// Sets tile spacing and margins.
func (self *Tileset) SetSpacingI(args ...interface{}) {
    self.Call("setSpacing", args)
}

// Updates tile coordinates and tileset data.
func (self *Tileset) UpdateTileDataI(args ...interface{}) {
    self.Call("updateTileData", args)
}
