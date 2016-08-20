// Automatic generation for Phaser.Tilemap
// generated file Tilemap.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Creates a new Phaser.Tilemap object. The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
// A Tile map is rendered to the display using a TilemapLayer. It is not added to the display list directly itself.
// A map may have multiple layers. You can perform operations on the map data such as copying, pasting, filling and shuffling the tiles around.
type Tilemap struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *Tilemap) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Tilemap) SetGame(member *Game) {
    self.Set("game", member)
}

// The key of this map data in the Phaser.Cache.
func (self *Tilemap) GetKey() string{
    return self.Get("key").String()
}

// The key of this map data in the Phaser.Cache.
func (self *Tilemap) SetKey(member string) {
    self.Set("key", member)
}

// The width of the map (in tiles).
func (self *Tilemap) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the map (in tiles).
func (self *Tilemap) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the map (in tiles).
func (self *Tilemap) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the map (in tiles).
func (self *Tilemap) SetHeight(member float64) {
    self.Set("height", member)
}

// The base width of the tiles in the map (in pixels).
func (self *Tilemap) GetTileWidth() float64{
    return self.Get("tileWidth").Float()
}

// The base width of the tiles in the map (in pixels).
func (self *Tilemap) SetTileWidth(member float64) {
    self.Set("tileWidth", member)
}

// The base height of the tiles in the map (in pixels).
func (self *Tilemap) GetTileHeight() float64{
    return self.Get("tileHeight").Float()
}

// The base height of the tiles in the map (in pixels).
func (self *Tilemap) SetTileHeight(member float64) {
    self.Set("tileHeight", member)
}

// The orientation of the map data (as specified in Tiled), usually 'orthogonal'.
func (self *Tilemap) GetOrientation() string{
    return self.Get("orientation").String()
}

// The orientation of the map data (as specified in Tiled), usually 'orthogonal'.
func (self *Tilemap) SetOrientation(member string) {
    self.Set("orientation", member)
}

// The format of the map data, either Phaser.Tilemap.CSV or Phaser.Tilemap.TILED_JSON.
func (self *Tilemap) GetFormat() float64{
    return self.Get("format").Float()
}

// The format of the map data, either Phaser.Tilemap.CSV or Phaser.Tilemap.TILED_JSON.
func (self *Tilemap) SetFormat(member float64) {
    self.Set("format", member)
}

// The version of the map data (as specified in Tiled, usually 1).
func (self *Tilemap) GetVersion() float64{
    return self.Get("version").Float()
}

// The version of the map data (as specified in Tiled, usually 1).
func (self *Tilemap) SetVersion(member float64) {
    self.Set("version", member)
}

// Map specific properties as specified in Tiled.
func (self *Tilemap) GetProperties() interface{}{
    return self.Get("properties")
}

// Map specific properties as specified in Tiled.
func (self *Tilemap) SetProperties(member interface{}) {
    self.Set("properties", member)
}

// The width of the map in pixels based on width * tileWidth.
func (self *Tilemap) GetWidthInPixels() float64{
    return self.Get("widthInPixels").Float()
}

// The width of the map in pixels based on width * tileWidth.
func (self *Tilemap) SetWidthInPixels(member float64) {
    self.Set("widthInPixels", member)
}

// The height of the map in pixels based on height * tileHeight.
func (self *Tilemap) GetHeightInPixels() float64{
    return self.Get("heightInPixels").Float()
}

// The height of the map in pixels based on height * tileHeight.
func (self *Tilemap) SetHeightInPixels(member float64) {
    self.Set("heightInPixels", member)
}

// An array of Tilemap layer data.
func (self *Tilemap) GetLayers() []interface{}{
	array := self.Get("layers")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of Tilemap layer data.
func (self *Tilemap) SetLayers(member []interface{}) {
    self.Set("layers", member)
}

// An array of Tilesets.
func (self *Tilemap) GetTilesets() []interface{}{
	array := self.Get("tilesets")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of Tilesets.
func (self *Tilemap) SetTilesets(member []interface{}) {
    self.Set("tilesets", member)
}

// An array of Image Collections.
func (self *Tilemap) GetImagecollections() []interface{}{
	array := self.Get("imagecollections")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of Image Collections.
func (self *Tilemap) SetImagecollections(member []interface{}) {
    self.Set("imagecollections", member)
}

// The super array of Tiles.
func (self *Tilemap) GetTiles() []interface{}{
	array := self.Get("tiles")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The super array of Tiles.
func (self *Tilemap) SetTiles(member []interface{}) {
    self.Set("tiles", member)
}

// An array of Tiled Object Layers.
func (self *Tilemap) GetObjects() []interface{}{
	array := self.Get("objects")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of Tiled Object Layers.
func (self *Tilemap) SetObjects(member []interface{}) {
    self.Set("objects", member)
}

// An array of tile indexes that collide.
func (self *Tilemap) GetCollideIndexes() []interface{}{
	array := self.Get("collideIndexes")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of tile indexes that collide.
func (self *Tilemap) SetCollideIndexes(member []interface{}) {
    self.Set("collideIndexes", member)
}

// An array of collision data (polylines, etc).
func (self *Tilemap) GetCollision() []interface{}{
	array := self.Get("collision")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of collision data (polylines, etc).
func (self *Tilemap) SetCollision(member []interface{}) {
    self.Set("collision", member)
}

// An array of Tiled Image Layers.
func (self *Tilemap) GetImages() []interface{}{
	array := self.Get("images")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of Tiled Image Layers.
func (self *Tilemap) SetImages(member []interface{}) {
    self.Set("images", member)
}

// The current layer.
func (self *Tilemap) GetCurrentLayer() float64{
    return self.Get("currentLayer").Float()
}

// The current layer.
func (self *Tilemap) SetCurrentLayer(member float64) {
    self.Set("currentLayer", member)
}

// Map data used for debug values only.
func (self *Tilemap) GetDebugMap() []interface{}{
	array := self.Get("debugMap")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Map data used for debug values only.
func (self *Tilemap) SetDebugMap(member []interface{}) {
    self.Set("debugMap", member)
}

// 
func (self *Tilemap) GetCSV() float64{
    return self.Get("CSV").Float()
}

// 
func (self *Tilemap) SetCSV(member float64) {
    self.Set("CSV", member)
}

// 
func (self *Tilemap) GetTILED_JSON() float64{
    return self.Get("TILED_JSON").Float()
}

// 
func (self *Tilemap) SetTILED_JSON(member float64) {
    self.Set("TILED_JSON", member)
}

// 
func (self *Tilemap) GetNORTH() float64{
    return self.Get("NORTH").Float()
}

// 
func (self *Tilemap) SetNORTH(member float64) {
    self.Set("NORTH", member)
}

// 
func (self *Tilemap) GetEAST() float64{
    return self.Get("EAST").Float()
}

// 
func (self *Tilemap) SetEAST(member float64) {
    self.Set("EAST", member)
}

// 
func (self *Tilemap) GetSOUTH() float64{
    return self.Get("SOUTH").Float()
}

// 
func (self *Tilemap) SetSOUTH(member float64) {
    self.Set("SOUTH", member)
}

// 
func (self *Tilemap) GetWEST() float64{
    return self.Get("WEST").Float()
}

// 
func (self *Tilemap) SetWEST(member float64) {
    self.Set("WEST", member)
}

// The current layer object.
func (self *Tilemap) GetLayer() interface{}{
    return self.Get("layer")
}

// The current layer object.
func (self *Tilemap) SetLayer(member interface{}) {
    self.Set("layer", member)
}



// Creates an empty map of the given dimensions and one blank layer. If layers already exist they are erased.
func (self *Tilemap) CreateI(args ...interface{}) *TilemapLayer{
    return &TilemapLayer{self.Call("create", args)}
}

// Sets the base tile size for the map.
func (self *Tilemap) SetTileSizeI(args ...interface{}) {
    self.Call("setTileSize", args)
}

// Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImageI(args ...interface{}) *Tileset{
    return &Tileset{self.Call("addTilesetImage", args)}
}

// Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjectsI(args ...interface{}) {
    self.Call("createFromObjects", args)
}

// Creates a Sprite for every object matching the given tile indexes in the map data.
// You can specify the group that the Sprite will be created in. If none is given it will be created in the World.
// You can optional specify if the tile will be replaced with another after the Sprite is created. This is useful if you want to lay down special 
// tiles in a level that are converted to Sprites, but want to replace the tile itself with a floor tile or similar once converted.
func (self *Tilemap) CreateFromTilesI(args ...interface{}) int{
    return self.Call("createFromTiles", args).Int()
}

// Creates a new TilemapLayer object. By default TilemapLayers are fixed to the camera.
// The `layer` parameter is important. If you've created your map in Tiled then you can get this by looking in Tiled and looking at the Layer name.
// Or you can open the JSON file it exports and look at the layers[].name value. Either way it must match.
// If you wish to create a blank layer to put your own tiles on then see Tilemap.createBlankLayer.
func (self *Tilemap) CreateLayerI(args ...interface{}) *TilemapLayer{
    return &TilemapLayer{self.Call("createLayer", args)}
}

// Creates a new and empty layer on this Tilemap. By default TilemapLayers are fixed to the camera.
func (self *Tilemap) CreateBlankLayerI(args ...interface{}) *TilemapLayer{
    return &TilemapLayer{self.Call("createBlankLayer", args)}
}

// Gets the layer index based on the layers name.
func (self *Tilemap) GetIndexI(args ...interface{}) float64{
    return self.Call("getIndex", args).Float()
}

// Gets the layer index based on its name.
func (self *Tilemap) GetLayerIndexI(args ...interface{}) float64{
    return self.Call("getLayerIndex", args).Float()
}

// Gets the tileset index based on its name.
func (self *Tilemap) GetTilesetIndexI(args ...interface{}) float64{
    return self.Call("getTilesetIndex", args).Float()
}

// Gets the image index based on its name.
func (self *Tilemap) GetImageIndexI(args ...interface{}) float64{
    return self.Call("getImageIndex", args).Float()
}

// Sets a global collision callback for the given tile index within the layer. This will affect all tiles on this layer that have the same index.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileIndexCallbackI(args ...interface{}) {
    self.Call("setTileIndexCallback", args)
}

// Sets a global collision callback for the given map location within the layer. This will affect all tiles on this layer found in the given area.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileLocationCallbackI(args ...interface{}) {
    self.Call("setTileLocationCallback", args)
}

// Sets collision the given tile or tiles. You can pass in either a single numeric index or an array of indexes: [ 2, 3, 15, 20].
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionI(args ...interface{}) {
    self.Call("setCollision", args)
}

// Sets collision on a range of tiles where the tile IDs increment sequentially.
// Calling this with a start value of 10 and a stop value of 14 would set collision for tiles 10, 11, 12, 13 and 14.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionBetweenI(args ...interface{}) {
    self.Call("setCollisionBetween", args)
}

// Sets collision on all tiles in the given layer, except for the IDs of those in the given array.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionByExclusionI(args ...interface{}) {
    self.Call("setCollisionByExclusion", args)
}

// Sets collision values on a tile in the set.
// You shouldn't usually call this method directly, instead use setCollision, setCollisionBetween or setCollisionByExclusion.
func (self *Tilemap) SetCollisionByIndexI(args ...interface{}) {
    self.Call("setCollisionByIndex", args)
}

// Gets the TilemapLayer index as used in the setCollision calls.
func (self *Tilemap) GetLayerI(args ...interface{}) float64{
    return self.Call("getLayer", args).Float()
}

// Turn off/on the recalculation of faces for tile or collision updates. 
// `setPreventRecalculate(true)` puts recalculation on hold while `setPreventRecalculate(false)` recalculates all the changed layers.
func (self *Tilemap) SetPreventRecalculateI(args ...interface{}) {
    self.Call("setPreventRecalculate", args)
}

// Internal function.
func (self *Tilemap) CalculateFacesI(args ...interface{}) {
    self.Call("calculateFaces", args)
}

// Gets the tile above the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileAboveI(args ...interface{}) {
    self.Call("getTileAbove", args)
}

// Gets the tile below the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileBelowI(args ...interface{}) {
    self.Call("getTileBelow", args)
}

// Gets the tile to the left of the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileLeftI(args ...interface{}) {
    self.Call("getTileLeft", args)
}

// Gets the tile to the right of the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileRightI(args ...interface{}) {
    self.Call("getTileRight", args)
}

// Sets the current layer to the given index.
func (self *Tilemap) SetLayerI(args ...interface{}) {
    self.Call("setLayer", args)
}

// Checks if there is a tile at the given location.
func (self *Tilemap) HasTileI(args ...interface{}) bool{
    return self.Call("hasTile", args).Bool()
}

// Removes the tile located at the given coordinates and updates the collision data.
func (self *Tilemap) RemoveTileI(args ...interface{}) *Tile{
    return &Tile{self.Call("removeTile", args)}
}

// Removes the tile located at the given coordinates and updates the collision data. The coordinates are given in pixel values.
func (self *Tilemap) RemoveTileWorldXYI(args ...interface{}) *Tile{
    return &Tile{self.Call("removeTileWorldXY", args)}
}

// Puts a tile of the given index value at the coordinate specified.
// If you pass `null` as the tile it will pass your call over to Tilemap.removeTile instead.
func (self *Tilemap) PutTileI(args ...interface{}) *Tile{
    return &Tile{self.Call("putTile", args)}
}

// Puts a tile into the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) PutTileWorldXYI(args ...interface{}) *Tile{
    return &Tile{self.Call("putTileWorldXY", args)}
}

// Searches the entire map layer for the first tile matching the given index, then returns that Phaser.Tile object.
// If no match is found it returns null.
// The search starts from the top-left tile and continues horizontally until it hits the end of the row, then it drops down to the next column.
// If the reverse boolean is true, it scans starting from the bottom-right corner traveling up to the top-left.
func (self *Tilemap) SearchTileIndexI(args ...interface{}) *Tile{
    return &Tile{self.Call("searchTileIndex", args)}
}

// Gets a tile from the Tilemap Layer. The coordinates are given in tile values.
func (self *Tilemap) GetTileI(args ...interface{}) *Tile{
    return &Tile{self.Call("getTile", args)}
}

// Gets a tile from the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) GetTileWorldXYI(args ...interface{}) *Tile{
    return &Tile{self.Call("getTileWorldXY", args)}
}

// Copies all of the tiles in the given rectangular block into the tilemap data buffer.
func (self *Tilemap) CopyI(args ...interface{}) []interface{}{
	array := self.Call("copy", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Pastes a previously copied block of tile data into the given x/y coordinates. Data should have been prepared with Tilemap.copy.
func (self *Tilemap) PasteI(args ...interface{}) {
    self.Call("paste", args)
}

// Scans the given area for tiles with an index matching tileA and swaps them with tileB.
func (self *Tilemap) SwapI(args ...interface{}) {
    self.Call("swap", args)
}

// Internal function that handles the swapping of tiles.
func (self *Tilemap) SwapHandlerI(args ...interface{}) {
    self.Call("swapHandler", args)
}

// For each tile in the given area defined by x/y and width/height run the given callback.
func (self *Tilemap) ForEachI(args ...interface{}) {
    self.Call("forEach", args)
}

// Scans the given area for tiles with an index matching `source` and updates their index to match `dest`.
func (self *Tilemap) ReplaceI(args ...interface{}) {
    self.Call("replace", args)
}

// Randomises a set of tiles in a given area.
func (self *Tilemap) RandomI(args ...interface{}) {
    self.Call("random", args)
}

// Shuffles a set of tiles in a given area. It will only randomise the tiles in that area, so if they're all the same nothing will appear to have changed!
func (self *Tilemap) ShuffleI(args ...interface{}) {
    self.Call("shuffle", args)
}

// Fills the given area with the specified tile.
func (self *Tilemap) FillI(args ...interface{}) {
    self.Call("fill", args)
}

// Removes all layers from this tile map.
func (self *Tilemap) RemoveAllLayersI(args ...interface{}) {
    self.Call("removeAllLayers", args)
}

// Dumps the tilemap data out to the console.
func (self *Tilemap) DumpI(args ...interface{}) {
    self.Call("dump", args)
}

// Removes all layer data from this tile map and nulls the game reference.
// Note: You are responsible for destroying any TilemapLayer objects you generated yourself, as Tilemap doesn't keep a reference to them.
func (self *Tilemap) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
