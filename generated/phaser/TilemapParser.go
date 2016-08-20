// Automatic generation for Phaser.TilemapParser
// generated file TilemapParser.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Phaser.TilemapParser parses data objects from Phaser.Loader that need more preparation before they can be inserted into a Tilemap.
type TilemapParser struct {
    *js.Object
}


// When scanning the Tiled map data the TilemapParser can either insert a null value (true) or
// a Phaser.Tile instance with an index of -1 (false, the default). Depending on your game type
// depends how this should be configured. If you've a large sparsely populated map and the tile
// data doesn't need to change then setting this value to `true` will help with memory consumption.
// However if your map is small, or you need to update the tiles (perhaps the map dynamically changes
// during the game) then leave the default value set.
func (self *TilemapParser) GetINSERT_NULL() bool{
    return self.Get("INSERT_NULL").Bool()
}

// When scanning the Tiled map data the TilemapParser can either insert a null value (true) or
// a Phaser.Tile instance with an index of -1 (false, the default). Depending on your game type
// depends how this should be configured. If you've a large sparsely populated map and the tile
// data doesn't need to change then setting this value to `true` will help with memory consumption.
// However if your map is small, or you need to update the tiles (perhaps the map dynamically changes
// during the game) then leave the default value set.
func (self *TilemapParser) SetINSERT_NULL(member bool) {
    self.Set("INSERT_NULL", member)
}

// A tiled flag that resides within the 32 bit of the object gid and
// indicates whether the tiled/object is flipped horizontally.
func (self *TilemapParser) GetFLIPPED_HORIZONTALLY_FLAG() float64{
    return self.Get("FLIPPED_HORIZONTALLY_FLAG").Float()
}

// A tiled flag that resides within the 32 bit of the object gid and
// indicates whether the tiled/object is flipped horizontally.
func (self *TilemapParser) SetFLIPPED_HORIZONTALLY_FLAG(member float64) {
    self.Set("FLIPPED_HORIZONTALLY_FLAG", member)
}

// A tiled flag that resides within the 31 bit of the object gid and
// indicates whether the tiled/object is flipped vertically.
func (self *TilemapParser) GetFLIPPED_VERTICALLY_FLAG() float64{
    return self.Get("FLIPPED_VERTICALLY_FLAG").Float()
}

// A tiled flag that resides within the 31 bit of the object gid and
// indicates whether the tiled/object is flipped vertically.
func (self *TilemapParser) SetFLIPPED_VERTICALLY_FLAG(member float64) {
    self.Set("FLIPPED_VERTICALLY_FLAG", member)
}

// A tiled flag that resides within the 30 bit of the object gid and
// indicates whether the tiled/object is flipped diagonally.
func (self *TilemapParser) GetFLIPPED_DIAGONALLY_FLAG() float64{
    return self.Get("FLIPPED_DIAGONALLY_FLAG").Float()
}

// A tiled flag that resides within the 30 bit of the object gid and
// indicates whether the tiled/object is flipped diagonally.
func (self *TilemapParser) SetFLIPPED_DIAGONALLY_FLAG(member float64) {
    self.Set("FLIPPED_DIAGONALLY_FLAG", member)
}



// Parse tilemap data from the cache and creates a Tilemap object.
func (self *TilemapParser) ParseI(args ...interface{}) interface{}{
    return self.Call("parse", args)
}

// Parses a CSV file into valid map data.
func (self *TilemapParser) ParseCSVI(args ...interface{}) interface{}{
    return self.Call("parseCSV", args)
}

// Returns an empty map data object.
func (self *TilemapParser) GetEmptyDataI(args ...interface{}) interface{}{
    return self.Call("getEmptyData", args)
}

// Parses a Tiled JSON file into valid map data.
func (self *TilemapParser) ParseJSONI(args ...interface{}) interface{}{
    return self.Call("parseJSON", args)
}
