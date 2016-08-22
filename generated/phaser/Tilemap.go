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
func (self *Tilemap) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *Tilemap) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The key of this map data in the Phaser.Cache.
func (self *Tilemap) GetKeyA() string{
    return self.Object.Get("key").String()
}

// The key of this map data in the Phaser.Cache.
func (self *Tilemap) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// The width of the map (in tiles).
func (self *Tilemap) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the map (in tiles).
func (self *Tilemap) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the map (in tiles).
func (self *Tilemap) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the map (in tiles).
func (self *Tilemap) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The base width of the tiles in the map (in pixels).
func (self *Tilemap) GetTileWidthA() int{
    return self.Object.Get("tileWidth").Int()
}

// The base width of the tiles in the map (in pixels).
func (self *Tilemap) SetTileWidthA(member int) {
    self.Object.Set("tileWidth", member)
}

// The base height of the tiles in the map (in pixels).
func (self *Tilemap) GetTileHeightA() int{
    return self.Object.Get("tileHeight").Int()
}

// The base height of the tiles in the map (in pixels).
func (self *Tilemap) SetTileHeightA(member int) {
    self.Object.Set("tileHeight", member)
}

// The orientation of the map data (as specified in Tiled), usually 'orthogonal'.
func (self *Tilemap) GetOrientationA() string{
    return self.Object.Get("orientation").String()
}

// The orientation of the map data (as specified in Tiled), usually 'orthogonal'.
func (self *Tilemap) SetOrientationA(member string) {
    self.Object.Set("orientation", member)
}

// The format of the map data, either Phaser.Tilemap.CSV or Phaser.Tilemap.TILED_JSON.
func (self *Tilemap) GetFormatA() int{
    return self.Object.Get("format").Int()
}

// The format of the map data, either Phaser.Tilemap.CSV or Phaser.Tilemap.TILED_JSON.
func (self *Tilemap) SetFormatA(member int) {
    self.Object.Set("format", member)
}

// The version of the map data (as specified in Tiled, usually 1).
func (self *Tilemap) GetVersionA() int{
    return self.Object.Get("version").Int()
}

// The version of the map data (as specified in Tiled, usually 1).
func (self *Tilemap) SetVersionA(member int) {
    self.Object.Set("version", member)
}

// Map specific properties as specified in Tiled.
func (self *Tilemap) GetPropertiesA() interface{}{
    return self.Object.Get("properties")
}

// Map specific properties as specified in Tiled.
func (self *Tilemap) SetPropertiesA(member interface{}) {
    self.Object.Set("properties", member)
}

// The width of the map in pixels based on width * tileWidth.
func (self *Tilemap) GetWidthInPixelsA() int{
    return self.Object.Get("widthInPixels").Int()
}

// The width of the map in pixels based on width * tileWidth.
func (self *Tilemap) SetWidthInPixelsA(member int) {
    self.Object.Set("widthInPixels", member)
}

// The height of the map in pixels based on height * tileHeight.
func (self *Tilemap) GetHeightInPixelsA() int{
    return self.Object.Get("heightInPixels").Int()
}

// The height of the map in pixels based on height * tileHeight.
func (self *Tilemap) SetHeightInPixelsA(member int) {
    self.Object.Set("heightInPixels", member)
}

// An array of Tilemap layer data.
func (self *Tilemap) GetLayersA() []interface{}{
	array00 := self.Object.Get("layers")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of Tilemap layer data.
func (self *Tilemap) SetLayersA(member []interface{}) {
    self.Object.Set("layers", member)
}

// An array of Tilesets.
func (self *Tilemap) GetTilesetsA() []interface{}{
	array00 := self.Object.Get("tilesets")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of Tilesets.
func (self *Tilemap) SetTilesetsA(member []interface{}) {
    self.Object.Set("tilesets", member)
}

// An array of Image Collections.
func (self *Tilemap) GetImagecollectionsA() []interface{}{
	array00 := self.Object.Get("imagecollections")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of Image Collections.
func (self *Tilemap) SetImagecollectionsA(member []interface{}) {
    self.Object.Set("imagecollections", member)
}

// The super array of Tiles.
func (self *Tilemap) GetTilesA() []interface{}{
	array00 := self.Object.Get("tiles")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// The super array of Tiles.
func (self *Tilemap) SetTilesA(member []interface{}) {
    self.Object.Set("tiles", member)
}

// An array of Tiled Object Layers.
func (self *Tilemap) GetObjectsA() []interface{}{
	array00 := self.Object.Get("objects")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of Tiled Object Layers.
func (self *Tilemap) SetObjectsA(member []interface{}) {
    self.Object.Set("objects", member)
}

// An array of tile indexes that collide.
func (self *Tilemap) GetCollideIndexesA() []interface{}{
	array00 := self.Object.Get("collideIndexes")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of tile indexes that collide.
func (self *Tilemap) SetCollideIndexesA(member []interface{}) {
    self.Object.Set("collideIndexes", member)
}

// An array of collision data (polylines, etc).
func (self *Tilemap) GetCollisionA() []interface{}{
	array00 := self.Object.Get("collision")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of collision data (polylines, etc).
func (self *Tilemap) SetCollisionA(member []interface{}) {
    self.Object.Set("collision", member)
}

// An array of Tiled Image Layers.
func (self *Tilemap) GetImagesA() []interface{}{
	array00 := self.Object.Get("images")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of Tiled Image Layers.
func (self *Tilemap) SetImagesA(member []interface{}) {
    self.Object.Set("images", member)
}

// The current layer.
func (self *Tilemap) GetCurrentLayerA() int{
    return self.Object.Get("currentLayer").Int()
}

// The current layer.
func (self *Tilemap) SetCurrentLayerA(member int) {
    self.Object.Set("currentLayer", member)
}

// Map data used for debug values only.
func (self *Tilemap) GetDebugMapA() []interface{}{
	array00 := self.Object.Get("debugMap")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Map data used for debug values only.
func (self *Tilemap) SetDebugMapA(member []interface{}) {
    self.Object.Set("debugMap", member)
}

// 
func (self *Tilemap) GetCSVA() int{
    return self.Object.Get("CSV").Int()
}

// 
func (self *Tilemap) SetCSVA(member int) {
    self.Object.Set("CSV", member)
}

// 
func (self *Tilemap) GetTILED_JSONA() int{
    return self.Object.Get("TILED_JSON").Int()
}

// 
func (self *Tilemap) SetTILED_JSONA(member int) {
    self.Object.Set("TILED_JSON", member)
}

// 
func (self *Tilemap) GetNORTHA() int{
    return self.Object.Get("NORTH").Int()
}

// 
func (self *Tilemap) SetNORTHA(member int) {
    self.Object.Set("NORTH", member)
}

// 
func (self *Tilemap) GetEASTA() int{
    return self.Object.Get("EAST").Int()
}

// 
func (self *Tilemap) SetEASTA(member int) {
    self.Object.Set("EAST", member)
}

// 
func (self *Tilemap) GetSOUTHA() int{
    return self.Object.Get("SOUTH").Int()
}

// 
func (self *Tilemap) SetSOUTHA(member int) {
    self.Object.Set("SOUTH", member)
}

// 
func (self *Tilemap) GetWESTA() int{
    return self.Object.Get("WEST").Int()
}

// 
func (self *Tilemap) SetWESTA(member int) {
    self.Object.Set("WEST", member)
}

// The current layer object.
func (self *Tilemap) GetLayerA() interface{}{
    return self.Object.Get("layer")
}

// The current layer object.
func (self *Tilemap) SetLayerA(member interface{}) {
    self.Object.Set("layer", member)
}



// Creates an empty map of the given dimensions and one blank layer. If layers already exist they are erased.
func (self *Tilemap) Create(name string, width int, height int, tileWidth int, tileHeight int, group *Group) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("create", name, width, height, tileWidth, tileHeight, group)}
}

// Creates an empty map of the given dimensions and one blank layer. If layers already exist they are erased.
func (self *Tilemap) CreateI(args ...interface{}) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("create", args)}
}

// Sets the base tile size for the map.
func (self *Tilemap) SetTileSize(tileWidth int, tileHeight int) {
    self.Object.Call("setTileSize", tileWidth, tileHeight)
}

// Sets the base tile size for the map.
func (self *Tilemap) SetTileSizeI(args ...interface{}) {
    self.Object.Call("setTileSize", args)
}

// Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImage(tileset string, key interface{}, tileWidth int, tileHeight int, tileMargin int, tileSpacing int, gid int) *Tileset{
    return &Tileset{self.Object.Call("addTilesetImage", tileset, key, tileWidth, tileHeight, tileMargin, tileSpacing, gid)}
}

// Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImageI(args ...interface{}) *Tileset{
    return &Tileset{self.Object.Call("addTilesetImage", args)}
}

// Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjects(name string, gid int, key string, frame interface{}, exists bool, autoCull bool, group *Group, CustomClass interface{}, adjustY bool) {
    self.Object.Call("createFromObjects", name, gid, key, frame, exists, autoCull, group, CustomClass, adjustY)
}

// Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjectsI(args ...interface{}) {
    self.Object.Call("createFromObjects", args)
}

// Creates a Sprite for every object matching the given tile indexes in the map data.
// You can specify the group that the Sprite will be created in. If none is given it will be created in the World.
// You can optional specify if the tile will be replaced with another after the Sprite is created. This is useful if you want to lay down special 
// tiles in a level that are converted to Sprites, but want to replace the tile itself with a floor tile or similar once converted.
func (self *Tilemap) CreateFromTiles(tiles interface{}, replacements interface{}, key string, layer interface{}, group *Group, properties interface{}) int{
    return self.Object.Call("createFromTiles", tiles, replacements, key, layer, group, properties).Int()
}

// Creates a Sprite for every object matching the given tile indexes in the map data.
// You can specify the group that the Sprite will be created in. If none is given it will be created in the World.
// You can optional specify if the tile will be replaced with another after the Sprite is created. This is useful if you want to lay down special 
// tiles in a level that are converted to Sprites, but want to replace the tile itself with a floor tile or similar once converted.
func (self *Tilemap) CreateFromTilesI(args ...interface{}) int{
    return self.Object.Call("createFromTiles", args).Int()
}

// Creates a new TilemapLayer object. By default TilemapLayers are fixed to the camera.
// The `layer` parameter is important. If you've created your map in Tiled then you can get this by looking in Tiled and looking at the Layer name.
// Or you can open the JSON file it exports and look at the layers[].name value. Either way it must match.
// If you wish to create a blank layer to put your own tiles on then see Tilemap.createBlankLayer.
func (self *Tilemap) CreateLayer(layer interface{}, width int, height int, group *Group, pixiTest bool) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createLayer", layer, width, height, group, pixiTest)}
}

// Creates a new TilemapLayer object. By default TilemapLayers are fixed to the camera.
// The `layer` parameter is important. If you've created your map in Tiled then you can get this by looking in Tiled and looking at the Layer name.
// Or you can open the JSON file it exports and look at the layers[].name value. Either way it must match.
// If you wish to create a blank layer to put your own tiles on then see Tilemap.createBlankLayer.
func (self *Tilemap) CreateLayerI(args ...interface{}) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createLayer", args)}
}

// Creates a new and empty layer on this Tilemap. By default TilemapLayers are fixed to the camera.
func (self *Tilemap) CreateBlankLayer(name string, width int, height int, tileWidth int, tileHeight int, group *Group) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createBlankLayer", name, width, height, tileWidth, tileHeight, group)}
}

// Creates a new and empty layer on this Tilemap. By default TilemapLayers are fixed to the camera.
func (self *Tilemap) CreateBlankLayerI(args ...interface{}) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createBlankLayer", args)}
}

// Gets the layer index based on the layers name.
func (self *Tilemap) GetIndex(location []interface{}, name string) int{
    return self.Object.Call("getIndex", location, name).Int()
}

// Gets the layer index based on the layers name.
func (self *Tilemap) GetIndexI(args ...interface{}) int{
    return self.Object.Call("getIndex", args).Int()
}

// Gets the layer index based on its name.
func (self *Tilemap) GetLayerIndex(name string) int{
    return self.Object.Call("getLayerIndex", name).Int()
}

// Gets the layer index based on its name.
func (self *Tilemap) GetLayerIndexI(args ...interface{}) int{
    return self.Object.Call("getLayerIndex", args).Int()
}

// Gets the tileset index based on its name.
func (self *Tilemap) GetTilesetIndex(name string) int{
    return self.Object.Call("getTilesetIndex", name).Int()
}

// Gets the tileset index based on its name.
func (self *Tilemap) GetTilesetIndexI(args ...interface{}) int{
    return self.Object.Call("getTilesetIndex", args).Int()
}

// Gets the image index based on its name.
func (self *Tilemap) GetImageIndex(name string) int{
    return self.Object.Call("getImageIndex", name).Int()
}

// Gets the image index based on its name.
func (self *Tilemap) GetImageIndexI(args ...interface{}) int{
    return self.Object.Call("getImageIndex", args).Int()
}

// Sets a global collision callback for the given tile index within the layer. This will affect all tiles on this layer that have the same index.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileIndexCallback(indexes interface{}, callback func(...interface{}), callbackContext interface{}, layer interface{}) {
    self.Object.Call("setTileIndexCallback", indexes, callback, callbackContext, layer)
}

// Sets a global collision callback for the given tile index within the layer. This will affect all tiles on this layer that have the same index.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileIndexCallbackI(args ...interface{}) {
    self.Object.Call("setTileIndexCallback", args)
}

// Sets a global collision callback for the given map location within the layer. This will affect all tiles on this layer found in the given area.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileLocationCallback(x int, y int, width int, height int, callback func(...interface{}), callbackContext interface{}, layer interface{}) {
    self.Object.Call("setTileLocationCallback", x, y, width, height, callback, callbackContext, layer)
}

// Sets a global collision callback for the given map location within the layer. This will affect all tiles on this layer found in the given area.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileLocationCallbackI(args ...interface{}) {
    self.Object.Call("setTileLocationCallback", args)
}

// Sets collision the given tile or tiles. You can pass in either a single numeric index or an array of indexes: [ 2, 3, 15, 20].
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollision(indexes interface{}, collides bool, layer interface{}, recalculate bool) {
    self.Object.Call("setCollision", indexes, collides, layer, recalculate)
}

// Sets collision the given tile or tiles. You can pass in either a single numeric index or an array of indexes: [ 2, 3, 15, 20].
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionI(args ...interface{}) {
    self.Object.Call("setCollision", args)
}

// Sets collision on a range of tiles where the tile IDs increment sequentially.
// Calling this with a start value of 10 and a stop value of 14 would set collision for tiles 10, 11, 12, 13 and 14.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionBetween(start int, stop int, collides bool, layer interface{}, recalculate bool) {
    self.Object.Call("setCollisionBetween", start, stop, collides, layer, recalculate)
}

// Sets collision on a range of tiles where the tile IDs increment sequentially.
// Calling this with a start value of 10 and a stop value of 14 would set collision for tiles 10, 11, 12, 13 and 14.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionBetweenI(args ...interface{}) {
    self.Object.Call("setCollisionBetween", args)
}

// Sets collision on all tiles in the given layer, except for the IDs of those in the given array.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionByExclusion(indexes []interface{}, collides bool, layer interface{}, recalculate bool) {
    self.Object.Call("setCollisionByExclusion", indexes, collides, layer, recalculate)
}

// Sets collision on all tiles in the given layer, except for the IDs of those in the given array.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionByExclusionI(args ...interface{}) {
    self.Object.Call("setCollisionByExclusion", args)
}

// Sets collision values on a tile in the set.
// You shouldn't usually call this method directly, instead use setCollision, setCollisionBetween or setCollisionByExclusion.
func (self *Tilemap) SetCollisionByIndex(index int, collides bool, layer int, recalculate bool) {
    self.Object.Call("setCollisionByIndex", index, collides, layer, recalculate)
}

// Sets collision values on a tile in the set.
// You shouldn't usually call this method directly, instead use setCollision, setCollisionBetween or setCollisionByExclusion.
func (self *Tilemap) SetCollisionByIndexI(args ...interface{}) {
    self.Object.Call("setCollisionByIndex", args)
}

// Gets the TilemapLayer index as used in the setCollision calls.
func (self *Tilemap) GetLayer(layer interface{}) int{
    return self.Object.Call("getLayer", layer).Int()
}

// Gets the TilemapLayer index as used in the setCollision calls.
func (self *Tilemap) GetLayerI(args ...interface{}) int{
    return self.Object.Call("getLayer", args).Int()
}

// Turn off/on the recalculation of faces for tile or collision updates. 
// `setPreventRecalculate(true)` puts recalculation on hold while `setPreventRecalculate(false)` recalculates all the changed layers.
func (self *Tilemap) SetPreventRecalculate(value bool) {
    self.Object.Call("setPreventRecalculate", value)
}

// Turn off/on the recalculation of faces for tile or collision updates. 
// `setPreventRecalculate(true)` puts recalculation on hold while `setPreventRecalculate(false)` recalculates all the changed layers.
func (self *Tilemap) SetPreventRecalculateI(args ...interface{}) {
    self.Object.Call("setPreventRecalculate", args)
}

// Internal function.
func (self *Tilemap) CalculateFaces(layer int) {
    self.Object.Call("calculateFaces", layer)
}

// Internal function.
func (self *Tilemap) CalculateFacesI(args ...interface{}) {
    self.Object.Call("calculateFaces", args)
}

// Gets the tile above the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileAbove(layer int, x int, y int) {
    self.Object.Call("getTileAbove", layer, x, y)
}

// Gets the tile above the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileAboveI(args ...interface{}) {
    self.Object.Call("getTileAbove", args)
}

// Gets the tile below the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileBelow(layer int, x int, y int) {
    self.Object.Call("getTileBelow", layer, x, y)
}

// Gets the tile below the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileBelowI(args ...interface{}) {
    self.Object.Call("getTileBelow", args)
}

// Gets the tile to the left of the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileLeft(layer int, x int, y int) {
    self.Object.Call("getTileLeft", layer, x, y)
}

// Gets the tile to the left of the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileLeftI(args ...interface{}) {
    self.Object.Call("getTileLeft", args)
}

// Gets the tile to the right of the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileRight(layer int, x int, y int) {
    self.Object.Call("getTileRight", layer, x, y)
}

// Gets the tile to the right of the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileRightI(args ...interface{}) {
    self.Object.Call("getTileRight", args)
}

// Sets the current layer to the given index.
func (self *Tilemap) SetLayer(layer interface{}) {
    self.Object.Call("setLayer", layer)
}

// Sets the current layer to the given index.
func (self *Tilemap) SetLayerI(args ...interface{}) {
    self.Object.Call("setLayer", args)
}

// Checks if there is a tile at the given location.
func (self *Tilemap) HasTile(x int, y int, layer interface{}) bool{
    return self.Object.Call("hasTile", x, y, layer).Bool()
}

// Checks if there is a tile at the given location.
func (self *Tilemap) HasTileI(args ...interface{}) bool{
    return self.Object.Call("hasTile", args).Bool()
}

// Removes the tile located at the given coordinates and updates the collision data.
func (self *Tilemap) RemoveTile(x int, y int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("removeTile", x, y, layer)}
}

// Removes the tile located at the given coordinates and updates the collision data.
func (self *Tilemap) RemoveTileI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("removeTile", args)}
}

// Removes the tile located at the given coordinates and updates the collision data. The coordinates are given in pixel values.
func (self *Tilemap) RemoveTileWorldXY(x int, y int, tileWidth int, tileHeight int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("removeTileWorldXY", x, y, tileWidth, tileHeight, layer)}
}

// Removes the tile located at the given coordinates and updates the collision data. The coordinates are given in pixel values.
func (self *Tilemap) RemoveTileWorldXYI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("removeTileWorldXY", args)}
}

// Puts a tile of the given index value at the coordinate specified.
// If you pass `null` as the tile it will pass your call over to Tilemap.removeTile instead.
func (self *Tilemap) PutTile(tile interface{}, x int, y int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("putTile", tile, x, y, layer)}
}

// Puts a tile of the given index value at the coordinate specified.
// If you pass `null` as the tile it will pass your call over to Tilemap.removeTile instead.
func (self *Tilemap) PutTileI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("putTile", args)}
}

// Puts a tile into the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) PutTileWorldXY(tile interface{}, x int, y int, tileWidth int, tileHeight int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("putTileWorldXY", tile, x, y, tileWidth, tileHeight, layer)}
}

// Puts a tile into the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) PutTileWorldXYI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("putTileWorldXY", args)}
}

// Searches the entire map layer for the first tile matching the given index, then returns that Phaser.Tile object.
// If no match is found it returns null.
// The search starts from the top-left tile and continues horizontally until it hits the end of the row, then it drops down to the next column.
// If the reverse boolean is true, it scans starting from the bottom-right corner traveling up to the top-left.
func (self *Tilemap) SearchTileIndex(index int, skip int, reverse int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("searchTileIndex", index, skip, reverse, layer)}
}

// Searches the entire map layer for the first tile matching the given index, then returns that Phaser.Tile object.
// If no match is found it returns null.
// The search starts from the top-left tile and continues horizontally until it hits the end of the row, then it drops down to the next column.
// If the reverse boolean is true, it scans starting from the bottom-right corner traveling up to the top-left.
func (self *Tilemap) SearchTileIndexI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("searchTileIndex", args)}
}

// Gets a tile from the Tilemap Layer. The coordinates are given in tile values.
func (self *Tilemap) GetTile(x int, y int, layer interface{}, nonNull bool) *Tile{
    return &Tile{self.Object.Call("getTile", x, y, layer, nonNull)}
}

// Gets a tile from the Tilemap Layer. The coordinates are given in tile values.
func (self *Tilemap) GetTileI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("getTile", args)}
}

// Gets a tile from the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) GetTileWorldXY(x int, y int, tileWidth int, tileHeight int, layer interface{}, nonNull bool) *Tile{
    return &Tile{self.Object.Call("getTileWorldXY", x, y, tileWidth, tileHeight, layer, nonNull)}
}

// Gets a tile from the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) GetTileWorldXYI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("getTileWorldXY", args)}
}

// Copies all of the tiles in the given rectangular block into the tilemap data buffer.
func (self *Tilemap) Copy(x int, y int, width int, height int, layer interface{}) []interface{}{
	array00 := self.Object.Call("copy", x, y, width, height, layer)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Copies all of the tiles in the given rectangular block into the tilemap data buffer.
func (self *Tilemap) CopyI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("copy", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Pastes a previously copied block of tile data into the given x/y coordinates. Data should have been prepared with Tilemap.copy.
func (self *Tilemap) Paste(x int, y int, tileblock []interface{}, layer interface{}) {
    self.Object.Call("paste", x, y, tileblock, layer)
}

// Pastes a previously copied block of tile data into the given x/y coordinates. Data should have been prepared with Tilemap.copy.
func (self *Tilemap) PasteI(args ...interface{}) {
    self.Object.Call("paste", args)
}

// Scans the given area for tiles with an index matching tileA and swaps them with tileB.
func (self *Tilemap) Swap(tileA int, tileB int, x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("swap", tileA, tileB, x, y, width, height, layer)
}

// Scans the given area for tiles with an index matching tileA and swaps them with tileB.
func (self *Tilemap) SwapI(args ...interface{}) {
    self.Object.Call("swap", args)
}

// Internal function that handles the swapping of tiles.
func (self *Tilemap) SwapHandler(value int) {
    self.Object.Call("swapHandler", value)
}

// Internal function that handles the swapping of tiles.
func (self *Tilemap) SwapHandlerI(args ...interface{}) {
    self.Object.Call("swapHandler", args)
}

// For each tile in the given area defined by x/y and width/height run the given callback.
func (self *Tilemap) ForEach(callback int, context int, x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("forEach", callback, context, x, y, width, height, layer)
}

// For each tile in the given area defined by x/y and width/height run the given callback.
func (self *Tilemap) ForEachI(args ...interface{}) {
    self.Object.Call("forEach", args)
}

// Scans the given area for tiles with an index matching `source` and updates their index to match `dest`.
func (self *Tilemap) Replace(source int, dest int, x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("replace", source, dest, x, y, width, height, layer)
}

// Scans the given area for tiles with an index matching `source` and updates their index to match `dest`.
func (self *Tilemap) ReplaceI(args ...interface{}) {
    self.Object.Call("replace", args)
}

// Randomises a set of tiles in a given area.
func (self *Tilemap) Random(x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("random", x, y, width, height, layer)
}

// Randomises a set of tiles in a given area.
func (self *Tilemap) RandomI(args ...interface{}) {
    self.Object.Call("random", args)
}

// Shuffles a set of tiles in a given area. It will only randomise the tiles in that area, so if they're all the same nothing will appear to have changed!
func (self *Tilemap) Shuffle(x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("shuffle", x, y, width, height, layer)
}

// Shuffles a set of tiles in a given area. It will only randomise the tiles in that area, so if they're all the same nothing will appear to have changed!
func (self *Tilemap) ShuffleI(args ...interface{}) {
    self.Object.Call("shuffle", args)
}

// Fills the given area with the specified tile.
func (self *Tilemap) Fill(index int, x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("fill", index, x, y, width, height, layer)
}

// Fills the given area with the specified tile.
func (self *Tilemap) FillI(args ...interface{}) {
    self.Object.Call("fill", args)
}

// Removes all layers from this tile map.
func (self *Tilemap) RemoveAllLayers() {
    self.Object.Call("removeAllLayers")
}

// Removes all layers from this tile map.
func (self *Tilemap) RemoveAllLayersI(args ...interface{}) {
    self.Object.Call("removeAllLayers", args)
}

// Dumps the tilemap data out to the console.
func (self *Tilemap) Dump() {
    self.Object.Call("dump")
}

// Dumps the tilemap data out to the console.
func (self *Tilemap) DumpI(args ...interface{}) {
    self.Object.Call("dump", args)
}

// Removes all layer data from this tile map and nulls the game reference.
// Note: You are responsible for destroying any TilemapLayer objects you generated yourself, as Tilemap doesn't keep a reference to them.
func (self *Tilemap) Destroy() {
    self.Object.Call("destroy")
}

// Removes all layer data from this tile map and nulls the game reference.
// Note: You are responsible for destroying any TilemapLayer objects you generated yourself, as Tilemap doesn't keep a reference to them.
func (self *Tilemap) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
