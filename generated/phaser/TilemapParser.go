// Package phaser Automatic generation for Phaser.TilemapParser
// generated file TilemapParser.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// TilemapParser Phaser.TilemapParser parses data objects from Phaser.Loader that need more preparation before they can be inserted into a Tilemap.
type TilemapParser struct {
    *js.Object
}

// NewTilemapParser Phaser.TilemapParser parses data objects from Phaser.Loader that need more preparation before they can be inserted into a Tilemap.
func NewTilemapParser() *TilemapParser {
    return &TilemapParser{js.Global.Get("Phaser").Get("TilemapParser").New()}
}
// NewTilemapParserI Phaser.TilemapParser parses data objects from Phaser.Loader that need more preparation before they can be inserted into a Tilemap.
func NewTilemapParserI(args ...interface{}) *TilemapParser {
    return &TilemapParser{js.Global.Get("Phaser").Get("TilemapParser").New(args)}
}



// TilemapParser Binding conversion method to TilemapParser point 
func ToTilemapParser(jsStruct interface{}) *TilemapParser {
    if object, ok := jsStruct.(*js.Object); ok {
		return &TilemapParser{Object: object}
	}
	return nil
}



// INSERT_NULL When scanning the Tiled map data the TilemapParser can either insert a null value (true) or
// a Phaser.Tile instance with an index of -1 (false, the default). Depending on your game type
// depends how this should be configured. If you've a large sparsely populated map and the tile
// data doesn't need to change then setting this value to `true` will help with memory consumption.
// However if your map is small, or you need to update the tiles (perhaps the map dynamically changes
// during the game) then leave the default value set.
func (self *TilemapParser) INSERT_NULL() bool{
    return self.Object.Get("INSERT_NULL").Bool()
}

// SetINSERT_NULLA When scanning the Tiled map data the TilemapParser can either insert a null value (true) or
// a Phaser.Tile instance with an index of -1 (false, the default). Depending on your game type
// depends how this should be configured. If you've a large sparsely populated map and the tile
// data doesn't need to change then setting this value to `true` will help with memory consumption.
// However if your map is small, or you need to update the tiles (perhaps the map dynamically changes
// during the game) then leave the default value set.
func (self *TilemapParser) SetINSERT_NULLA(member bool) {
    self.Object.Set("INSERT_NULL", member)
}

// FLIPPED_HORIZONTALLY_FLAG A tiled flag that resides within the 32 bit of the object gid and
// indicates whether the tiled/object is flipped horizontally.
func (self *TilemapParser) FLIPPED_HORIZONTALLY_FLAG() int{
    return self.Object.Get("FLIPPED_HORIZONTALLY_FLAG").Int()
}

// SetFLIPPED_HORIZONTALLY_FLAGA A tiled flag that resides within the 32 bit of the object gid and
// indicates whether the tiled/object is flipped horizontally.
func (self *TilemapParser) SetFLIPPED_HORIZONTALLY_FLAGA(member int) {
    self.Object.Set("FLIPPED_HORIZONTALLY_FLAG", member)
}

// FLIPPED_VERTICALLY_FLAG A tiled flag that resides within the 31 bit of the object gid and
// indicates whether the tiled/object is flipped vertically.
func (self *TilemapParser) FLIPPED_VERTICALLY_FLAG() int{
    return self.Object.Get("FLIPPED_VERTICALLY_FLAG").Int()
}

// SetFLIPPED_VERTICALLY_FLAGA A tiled flag that resides within the 31 bit of the object gid and
// indicates whether the tiled/object is flipped vertically.
func (self *TilemapParser) SetFLIPPED_VERTICALLY_FLAGA(member int) {
    self.Object.Set("FLIPPED_VERTICALLY_FLAG", member)
}

// FLIPPED_DIAGONALLY_FLAG A tiled flag that resides within the 30 bit of the object gid and
// indicates whether the tiled/object is flipped diagonally.
func (self *TilemapParser) FLIPPED_DIAGONALLY_FLAG() int{
    return self.Object.Get("FLIPPED_DIAGONALLY_FLAG").Int()
}

// SetFLIPPED_DIAGONALLY_FLAGA A tiled flag that resides within the 30 bit of the object gid and
// indicates whether the tiled/object is flipped diagonally.
func (self *TilemapParser) SetFLIPPED_DIAGONALLY_FLAGA(member int) {
    self.Object.Set("FLIPPED_DIAGONALLY_FLAG", member)
}


// Parse Parse tilemap data from the cache and creates a Tilemap object.
func (self *TilemapParser) Parse(game *Game, key string) interface{}{
    return self.Object.Call("parse", game, key)
}

// Parse1O Parse tilemap data from the cache and creates a Tilemap object.
func (self *TilemapParser) Parse1O(game *Game, key string, tileWidth int) interface{}{
    return self.Object.Call("parse", game, key, tileWidth)
}

// Parse2O Parse tilemap data from the cache and creates a Tilemap object.
func (self *TilemapParser) Parse2O(game *Game, key string, tileWidth int, tileHeight int) interface{}{
    return self.Object.Call("parse", game, key, tileWidth, tileHeight)
}

// Parse3O Parse tilemap data from the cache and creates a Tilemap object.
func (self *TilemapParser) Parse3O(game *Game, key string, tileWidth int, tileHeight int, width int) interface{}{
    return self.Object.Call("parse", game, key, tileWidth, tileHeight, width)
}

// Parse4O Parse tilemap data from the cache and creates a Tilemap object.
func (self *TilemapParser) Parse4O(game *Game, key string, tileWidth int, tileHeight int, width int, height int) interface{}{
    return self.Object.Call("parse", game, key, tileWidth, tileHeight, width, height)
}

// ParseI Parse tilemap data from the cache and creates a Tilemap object.
func (self *TilemapParser) ParseI(args ...interface{}) interface{}{
    return self.Object.Call("parse", args)
}

// ParseCSV Parses a CSV file into valid map data.
func (self *TilemapParser) ParseCSV(key string, data string) interface{}{
    return self.Object.Call("parseCSV", key, data)
}

// ParseCSV1O Parses a CSV file into valid map data.
func (self *TilemapParser) ParseCSV1O(key string, data string, tileWidth int) interface{}{
    return self.Object.Call("parseCSV", key, data, tileWidth)
}

// ParseCSV2O Parses a CSV file into valid map data.
func (self *TilemapParser) ParseCSV2O(key string, data string, tileWidth int, tileHeight int) interface{}{
    return self.Object.Call("parseCSV", key, data, tileWidth, tileHeight)
}

// ParseCSVI Parses a CSV file into valid map data.
func (self *TilemapParser) ParseCSVI(args ...interface{}) interface{}{
    return self.Object.Call("parseCSV", args)
}

// GetEmptyData Returns an empty map data object.
func (self *TilemapParser) GetEmptyData() interface{}{
    return self.Object.Call("getEmptyData")
}

// GetEmptyDataI Returns an empty map data object.
func (self *TilemapParser) GetEmptyDataI(args ...interface{}) interface{}{
    return self.Object.Call("getEmptyData", args)
}

// ParseJSON Parses a Tiled JSON file into valid map data.
func (self *TilemapParser) ParseJSON(json interface{}) interface{}{
    return self.Object.Call("parseJSON", json)
}

// ParseJSONI Parses a Tiled JSON file into valid map data.
func (self *TilemapParser) ParseJSONI(args ...interface{}) interface{}{
    return self.Object.Call("parseJSON", args)
}

