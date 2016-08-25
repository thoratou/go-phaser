// Package phaser Automatic generation for Phaser.Tilemap
// generated file Tilemap.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Tilemap Creates a new Phaser.Tilemap object. The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
// A Tile map is rendered to the display using a TilemapLayer. It is not added to the display list directly itself.
// A map may have multiple layers. You can perform operations on the map data such as copying, pasting, filling and shuffling the tiles around.
type Tilemap struct {
    *js.Object
}

// NewTilemap Creates a new Phaser.Tilemap object. The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
// A Tile map is rendered to the display using a TilemapLayer. It is not added to the display list directly itself.
// A map may have multiple layers. You can perform operations on the map data such as copying, pasting, filling and shuffling the tiles around.
func NewTilemap(game *Game) *Tilemap {
    return &Tilemap{js.Global.Get("Phaser").Get("Tilemap").New(game)}
}
// NewTilemap1O Creates a new Phaser.Tilemap object. The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
// A Tile map is rendered to the display using a TilemapLayer. It is not added to the display list directly itself.
// A map may have multiple layers. You can perform operations on the map data such as copying, pasting, filling and shuffling the tiles around.
func NewTilemap1O(game *Game, key string) *Tilemap {
    return &Tilemap{js.Global.Get("Phaser").Get("Tilemap").New(game, key)}
}
// NewTilemap2O Creates a new Phaser.Tilemap object. The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
// A Tile map is rendered to the display using a TilemapLayer. It is not added to the display list directly itself.
// A map may have multiple layers. You can perform operations on the map data such as copying, pasting, filling and shuffling the tiles around.
func NewTilemap2O(game *Game, key string, tileWidth int) *Tilemap {
    return &Tilemap{js.Global.Get("Phaser").Get("Tilemap").New(game, key, tileWidth)}
}
// NewTilemap3O Creates a new Phaser.Tilemap object. The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
// A Tile map is rendered to the display using a TilemapLayer. It is not added to the display list directly itself.
// A map may have multiple layers. You can perform operations on the map data such as copying, pasting, filling and shuffling the tiles around.
func NewTilemap3O(game *Game, key string, tileWidth int, tileHeight int) *Tilemap {
    return &Tilemap{js.Global.Get("Phaser").Get("Tilemap").New(game, key, tileWidth, tileHeight)}
}
// NewTilemap4O Creates a new Phaser.Tilemap object. The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
// A Tile map is rendered to the display using a TilemapLayer. It is not added to the display list directly itself.
// A map may have multiple layers. You can perform operations on the map data such as copying, pasting, filling and shuffling the tiles around.
func NewTilemap4O(game *Game, key string, tileWidth int, tileHeight int, width int) *Tilemap {
    return &Tilemap{js.Global.Get("Phaser").Get("Tilemap").New(game, key, tileWidth, tileHeight, width)}
}
// NewTilemap5O Creates a new Phaser.Tilemap object. The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
// A Tile map is rendered to the display using a TilemapLayer. It is not added to the display list directly itself.
// A map may have multiple layers. You can perform operations on the map data such as copying, pasting, filling and shuffling the tiles around.
func NewTilemap5O(game *Game, key string, tileWidth int, tileHeight int, width int, height int) *Tilemap {
    return &Tilemap{js.Global.Get("Phaser").Get("Tilemap").New(game, key, tileWidth, tileHeight, width, height)}
}
// NewTilemapI Creates a new Phaser.Tilemap object. The map can either be populated with data from a Tiled JSON file or from a CSV file.
// To do this pass the Cache key as the first parameter. When using Tiled data you need only provide the key.
// When using CSV data you must provide the key and the tileWidth and tileHeight parameters.
// If creating a blank tilemap to be populated later, you can either specify no parameters at all and then use `Tilemap.create` or pass the map and tile dimensions here.
// Note that all Tilemaps use a base tile size to calculate dimensions from, but that a TilemapLayer may have its own unique tile size that overrides it.
// A Tile map is rendered to the display using a TilemapLayer. It is not added to the display list directly itself.
// A map may have multiple layers. You can perform operations on the map data such as copying, pasting, filling and shuffling the tiles around.
func NewTilemapI(args ...interface{}) *Tilemap {
    return &Tilemap{js.Global.Get("Phaser").Get("Tilemap").New(args)}
}



// Tilemap Binding conversion method to Tilemap point 
func ToTilemap(jsStruct interface{}) *Tilemap {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Tilemap{Object: object}
	}
	return nil
}



// Game A reference to the currently running Game.
func (self *Tilemap) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Tilemap) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Key The key of this map data in the Phaser.Cache.
func (self *Tilemap) Key() string{
    return self.Object.Get("key").String()
}

// SetKeyA The key of this map data in the Phaser.Cache.
func (self *Tilemap) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// Width The width of the map (in tiles).
func (self *Tilemap) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the map (in tiles).
func (self *Tilemap) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the map (in tiles).
func (self *Tilemap) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the map (in tiles).
func (self *Tilemap) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// TileWidth The base width of the tiles in the map (in pixels).
func (self *Tilemap) TileWidth() int{
    return self.Object.Get("tileWidth").Int()
}

// SetTileWidthA The base width of the tiles in the map (in pixels).
func (self *Tilemap) SetTileWidthA(member int) {
    self.Object.Set("tileWidth", member)
}

// TileHeight The base height of the tiles in the map (in pixels).
func (self *Tilemap) TileHeight() int{
    return self.Object.Get("tileHeight").Int()
}

// SetTileHeightA The base height of the tiles in the map (in pixels).
func (self *Tilemap) SetTileHeightA(member int) {
    self.Object.Set("tileHeight", member)
}

// Orientation The orientation of the map data (as specified in Tiled), usually 'orthogonal'.
func (self *Tilemap) Orientation() string{
    return self.Object.Get("orientation").String()
}

// SetOrientationA The orientation of the map data (as specified in Tiled), usually 'orthogonal'.
func (self *Tilemap) SetOrientationA(member string) {
    self.Object.Set("orientation", member)
}

// Format The format of the map data, either Phaser.Tilemap.CSV or Phaser.Tilemap.TILED_JSON.
func (self *Tilemap) Format() int{
    return self.Object.Get("format").Int()
}

// SetFormatA The format of the map data, either Phaser.Tilemap.CSV or Phaser.Tilemap.TILED_JSON.
func (self *Tilemap) SetFormatA(member int) {
    self.Object.Set("format", member)
}

// Version The version of the map data (as specified in Tiled, usually 1).
func (self *Tilemap) Version() int{
    return self.Object.Get("version").Int()
}

// SetVersionA The version of the map data (as specified in Tiled, usually 1).
func (self *Tilemap) SetVersionA(member int) {
    self.Object.Set("version", member)
}

// Properties Map specific properties as specified in Tiled.
func (self *Tilemap) Properties() interface{}{
    return self.Object.Get("properties")
}

// SetPropertiesA Map specific properties as specified in Tiled.
func (self *Tilemap) SetPropertiesA(member interface{}) {
    self.Object.Set("properties", member)
}

// WidthInPixels The width of the map in pixels based on width * tileWidth.
func (self *Tilemap) WidthInPixels() int{
    return self.Object.Get("widthInPixels").Int()
}

// SetWidthInPixelsA The width of the map in pixels based on width * tileWidth.
func (self *Tilemap) SetWidthInPixelsA(member int) {
    self.Object.Set("widthInPixels", member)
}

// HeightInPixels The height of the map in pixels based on height * tileHeight.
func (self *Tilemap) HeightInPixels() int{
    return self.Object.Get("heightInPixels").Int()
}

// SetHeightInPixelsA The height of the map in pixels based on height * tileHeight.
func (self *Tilemap) SetHeightInPixelsA(member int) {
    self.Object.Set("heightInPixels", member)
}

// Layers An array of Tilemap layer data.
func (self *Tilemap) Layers() []interface{}{
	array00 := self.Object.Get("layers")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetLayersA An array of Tilemap layer data.
func (self *Tilemap) SetLayersA(member []interface{}) {
    self.Object.Set("layers", member)
}

// Tilesets An array of Tilesets.
func (self *Tilemap) Tilesets() []interface{}{
	array00 := self.Object.Get("tilesets")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetTilesetsA An array of Tilesets.
func (self *Tilemap) SetTilesetsA(member []interface{}) {
    self.Object.Set("tilesets", member)
}

// Imagecollections An array of Image Collections.
func (self *Tilemap) Imagecollections() []interface{}{
	array00 := self.Object.Get("imagecollections")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetImagecollectionsA An array of Image Collections.
func (self *Tilemap) SetImagecollectionsA(member []interface{}) {
    self.Object.Set("imagecollections", member)
}

// Tiles The super array of Tiles.
func (self *Tilemap) Tiles() []interface{}{
	array00 := self.Object.Get("tiles")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetTilesA The super array of Tiles.
func (self *Tilemap) SetTilesA(member []interface{}) {
    self.Object.Set("tiles", member)
}

// Objects An array of Tiled Object Layers.
func (self *Tilemap) Objects() []interface{}{
	array00 := self.Object.Get("objects")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetObjectsA An array of Tiled Object Layers.
func (self *Tilemap) SetObjectsA(member []interface{}) {
    self.Object.Set("objects", member)
}

// CollideIndexes An array of tile indexes that collide.
func (self *Tilemap) CollideIndexes() []interface{}{
	array00 := self.Object.Get("collideIndexes")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetCollideIndexesA An array of tile indexes that collide.
func (self *Tilemap) SetCollideIndexesA(member []interface{}) {
    self.Object.Set("collideIndexes", member)
}

// Collision An array of collision data (polylines, etc).
func (self *Tilemap) Collision() []interface{}{
	array00 := self.Object.Get("collision")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetCollisionA An array of collision data (polylines, etc).
func (self *Tilemap) SetCollisionA(member []interface{}) {
    self.Object.Set("collision", member)
}

// Images An array of Tiled Image Layers.
func (self *Tilemap) Images() []interface{}{
	array00 := self.Object.Get("images")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetImagesA An array of Tiled Image Layers.
func (self *Tilemap) SetImagesA(member []interface{}) {
    self.Object.Set("images", member)
}

// CurrentLayer The current layer.
func (self *Tilemap) CurrentLayer() int{
    return self.Object.Get("currentLayer").Int()
}

// SetCurrentLayerA The current layer.
func (self *Tilemap) SetCurrentLayerA(member int) {
    self.Object.Set("currentLayer", member)
}

// DebugMap Map data used for debug values only.
func (self *Tilemap) DebugMap() []interface{}{
	array00 := self.Object.Get("debugMap")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetDebugMapA Map data used for debug values only.
func (self *Tilemap) SetDebugMapA(member []interface{}) {
    self.Object.Set("debugMap", member)
}

// CSV empty description
func (self *Tilemap) CSV() int{
    return self.Object.Get("CSV").Int()
}

// SetCSVA empty description
func (self *Tilemap) SetCSVA(member int) {
    self.Object.Set("CSV", member)
}

// TILED_JSON empty description
func (self *Tilemap) TILED_JSON() int{
    return self.Object.Get("TILED_JSON").Int()
}

// SetTILED_JSONA empty description
func (self *Tilemap) SetTILED_JSONA(member int) {
    self.Object.Set("TILED_JSON", member)
}

// NORTH empty description
func (self *Tilemap) NORTH() int{
    return self.Object.Get("NORTH").Int()
}

// SetNORTHA empty description
func (self *Tilemap) SetNORTHA(member int) {
    self.Object.Set("NORTH", member)
}

// EAST empty description
func (self *Tilemap) EAST() int{
    return self.Object.Get("EAST").Int()
}

// SetEASTA empty description
func (self *Tilemap) SetEASTA(member int) {
    self.Object.Set("EAST", member)
}

// SOUTH empty description
func (self *Tilemap) SOUTH() int{
    return self.Object.Get("SOUTH").Int()
}

// SetSOUTHA empty description
func (self *Tilemap) SetSOUTHA(member int) {
    self.Object.Set("SOUTH", member)
}

// WEST empty description
func (self *Tilemap) WEST() int{
    return self.Object.Get("WEST").Int()
}

// SetWESTA empty description
func (self *Tilemap) SetWESTA(member int) {
    self.Object.Set("WEST", member)
}

// Layer The current layer object.
func (self *Tilemap) Layer() interface{}{
    return self.Object.Get("layer")
}

// SetLayerA The current layer object.
func (self *Tilemap) SetLayerA(member interface{}) {
    self.Object.Set("layer", member)
}


// Create Creates an empty map of the given dimensions and one blank layer. If layers already exist they are erased.
func (self *Tilemap) Create(name string, width int, height int, tileWidth int, tileHeight int) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("create", name, width, height, tileWidth, tileHeight)}
}

// Create1O Creates an empty map of the given dimensions and one blank layer. If layers already exist they are erased.
func (self *Tilemap) Create1O(name string, width int, height int, tileWidth int, tileHeight int, group *Group) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("create", name, width, height, tileWidth, tileHeight, group)}
}

// CreateI Creates an empty map of the given dimensions and one blank layer. If layers already exist they are erased.
func (self *Tilemap) CreateI(args ...interface{}) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("create", args)}
}

// SetTileSize Sets the base tile size for the map.
func (self *Tilemap) SetTileSize(tileWidth int, tileHeight int) {
    self.Object.Call("setTileSize", tileWidth, tileHeight)
}

// SetTileSizeI Sets the base tile size for the map.
func (self *Tilemap) SetTileSizeI(args ...interface{}) {
    self.Object.Call("setTileSize", args)
}

// AddTilesetImage Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImage(tileset string) *Tileset{
    return &Tileset{self.Object.Call("addTilesetImage", tileset)}
}

// AddTilesetImage1O Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImage1O(tileset string, key interface{}) *Tileset{
    return &Tileset{self.Object.Call("addTilesetImage", tileset, key)}
}

// AddTilesetImage2O Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImage2O(tileset string, key interface{}, tileWidth int) *Tileset{
    return &Tileset{self.Object.Call("addTilesetImage", tileset, key, tileWidth)}
}

// AddTilesetImage3O Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImage3O(tileset string, key interface{}, tileWidth int, tileHeight int) *Tileset{
    return &Tileset{self.Object.Call("addTilesetImage", tileset, key, tileWidth, tileHeight)}
}

// AddTilesetImage4O Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImage4O(tileset string, key interface{}, tileWidth int, tileHeight int, tileMargin int) *Tileset{
    return &Tileset{self.Object.Call("addTilesetImage", tileset, key, tileWidth, tileHeight, tileMargin)}
}

// AddTilesetImage5O Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImage5O(tileset string, key interface{}, tileWidth int, tileHeight int, tileMargin int, tileSpacing int) *Tileset{
    return &Tileset{self.Object.Call("addTilesetImage", tileset, key, tileWidth, tileHeight, tileMargin, tileSpacing)}
}

// AddTilesetImage6O Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImage6O(tileset string, key interface{}, tileWidth int, tileHeight int, tileMargin int, tileSpacing int, gid int) *Tileset{
    return &Tileset{self.Object.Call("addTilesetImage", tileset, key, tileWidth, tileHeight, tileMargin, tileSpacing, gid)}
}

// AddTilesetImageI Adds an image to the map to be used as a tileset. A single map may use multiple tilesets.
// Note that the tileset name can be found in the JSON file exported from Tiled, or in the Tiled editor.
func (self *Tilemap) AddTilesetImageI(args ...interface{}) *Tileset{
    return &Tileset{self.Object.Call("addTilesetImage", args)}
}

// CreateFromObjects Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjects(name string, gid int, key string) {
    self.Object.Call("createFromObjects", name, gid, key)
}

// CreateFromObjects1O Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjects1O(name string, gid int, key string, frame interface{}) {
    self.Object.Call("createFromObjects", name, gid, key, frame)
}

// CreateFromObjects2O Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjects2O(name string, gid int, key string, frame interface{}, exists bool) {
    self.Object.Call("createFromObjects", name, gid, key, frame, exists)
}

// CreateFromObjects3O Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjects3O(name string, gid int, key string, frame interface{}, exists bool, autoCull bool) {
    self.Object.Call("createFromObjects", name, gid, key, frame, exists, autoCull)
}

// CreateFromObjects4O Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjects4O(name string, gid int, key string, frame interface{}, exists bool, autoCull bool, group *Group) {
    self.Object.Call("createFromObjects", name, gid, key, frame, exists, autoCull, group)
}

// CreateFromObjects5O Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjects5O(name string, gid int, key string, frame interface{}, exists bool, autoCull bool, group *Group, CustomClass interface{}) {
    self.Object.Call("createFromObjects", name, gid, key, frame, exists, autoCull, group, CustomClass)
}

// CreateFromObjects6O Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjects6O(name string, gid int, key string, frame interface{}, exists bool, autoCull bool, group *Group, CustomClass interface{}, adjustY bool) {
    self.Object.Call("createFromObjects", name, gid, key, frame, exists, autoCull, group, CustomClass, adjustY)
}

// CreateFromObjectsI Creates a Sprite for every object matching the given gid in the map data. You can optionally specify the group that the Sprite will be created in. If none is
// given it will be created in the World. All properties from the map data objectgroup are copied across to the Sprite, so you can use this as an easy way to
// configure Sprite properties from within the map editor. For example giving an object a property of alpha: 0.5 in the map editor will duplicate that when the
// Sprite is created. You could also give it a value like: body.velocity.x: 100 to set it moving automatically.
func (self *Tilemap) CreateFromObjectsI(args ...interface{}) {
    self.Object.Call("createFromObjects", args)
}

// CreateFromTiles Creates a Sprite for every object matching the given tile indexes in the map data.
// You can specify the group that the Sprite will be created in. If none is given it will be created in the World.
// You can optional specify if the tile will be replaced with another after the Sprite is created. This is useful if you want to lay down special 
// tiles in a level that are converted to Sprites, but want to replace the tile itself with a floor tile or similar once converted.
func (self *Tilemap) CreateFromTiles(tiles interface{}, replacements interface{}, key string) int{
    return self.Object.Call("createFromTiles", tiles, replacements, key).Int()
}

// CreateFromTiles1O Creates a Sprite for every object matching the given tile indexes in the map data.
// You can specify the group that the Sprite will be created in. If none is given it will be created in the World.
// You can optional specify if the tile will be replaced with another after the Sprite is created. This is useful if you want to lay down special 
// tiles in a level that are converted to Sprites, but want to replace the tile itself with a floor tile or similar once converted.
func (self *Tilemap) CreateFromTiles1O(tiles interface{}, replacements interface{}, key string, layer interface{}) int{
    return self.Object.Call("createFromTiles", tiles, replacements, key, layer).Int()
}

// CreateFromTiles2O Creates a Sprite for every object matching the given tile indexes in the map data.
// You can specify the group that the Sprite will be created in. If none is given it will be created in the World.
// You can optional specify if the tile will be replaced with another after the Sprite is created. This is useful if you want to lay down special 
// tiles in a level that are converted to Sprites, but want to replace the tile itself with a floor tile or similar once converted.
func (self *Tilemap) CreateFromTiles2O(tiles interface{}, replacements interface{}, key string, layer interface{}, group *Group) int{
    return self.Object.Call("createFromTiles", tiles, replacements, key, layer, group).Int()
}

// CreateFromTiles3O Creates a Sprite for every object matching the given tile indexes in the map data.
// You can specify the group that the Sprite will be created in. If none is given it will be created in the World.
// You can optional specify if the tile will be replaced with another after the Sprite is created. This is useful if you want to lay down special 
// tiles in a level that are converted to Sprites, but want to replace the tile itself with a floor tile or similar once converted.
func (self *Tilemap) CreateFromTiles3O(tiles interface{}, replacements interface{}, key string, layer interface{}, group *Group, properties interface{}) int{
    return self.Object.Call("createFromTiles", tiles, replacements, key, layer, group, properties).Int()
}

// CreateFromTilesI Creates a Sprite for every object matching the given tile indexes in the map data.
// You can specify the group that the Sprite will be created in. If none is given it will be created in the World.
// You can optional specify if the tile will be replaced with another after the Sprite is created. This is useful if you want to lay down special 
// tiles in a level that are converted to Sprites, but want to replace the tile itself with a floor tile or similar once converted.
func (self *Tilemap) CreateFromTilesI(args ...interface{}) int{
    return self.Object.Call("createFromTiles", args).Int()
}

// CreateLayer Creates a new TilemapLayer object. By default TilemapLayers are fixed to the camera.
// The `layer` parameter is important. If you've created your map in Tiled then you can get this by looking in Tiled and looking at the Layer name.
// Or you can open the JSON file it exports and look at the layers[].name value. Either way it must match.
// If you wish to create a blank layer to put your own tiles on then see Tilemap.createBlankLayer.
func (self *Tilemap) CreateLayer(layer interface{}) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createLayer", layer)}
}

// CreateLayer1O Creates a new TilemapLayer object. By default TilemapLayers are fixed to the camera.
// The `layer` parameter is important. If you've created your map in Tiled then you can get this by looking in Tiled and looking at the Layer name.
// Or you can open the JSON file it exports and look at the layers[].name value. Either way it must match.
// If you wish to create a blank layer to put your own tiles on then see Tilemap.createBlankLayer.
func (self *Tilemap) CreateLayer1O(layer interface{}, width int) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createLayer", layer, width)}
}

// CreateLayer2O Creates a new TilemapLayer object. By default TilemapLayers are fixed to the camera.
// The `layer` parameter is important. If you've created your map in Tiled then you can get this by looking in Tiled and looking at the Layer name.
// Or you can open the JSON file it exports and look at the layers[].name value. Either way it must match.
// If you wish to create a blank layer to put your own tiles on then see Tilemap.createBlankLayer.
func (self *Tilemap) CreateLayer2O(layer interface{}, width int, height int) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createLayer", layer, width, height)}
}

// CreateLayer3O Creates a new TilemapLayer object. By default TilemapLayers are fixed to the camera.
// The `layer` parameter is important. If you've created your map in Tiled then you can get this by looking in Tiled and looking at the Layer name.
// Or you can open the JSON file it exports and look at the layers[].name value. Either way it must match.
// If you wish to create a blank layer to put your own tiles on then see Tilemap.createBlankLayer.
func (self *Tilemap) CreateLayer3O(layer interface{}, width int, height int, group *Group) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createLayer", layer, width, height, group)}
}

// CreateLayer4O Creates a new TilemapLayer object. By default TilemapLayers are fixed to the camera.
// The `layer` parameter is important. If you've created your map in Tiled then you can get this by looking in Tiled and looking at the Layer name.
// Or you can open the JSON file it exports and look at the layers[].name value. Either way it must match.
// If you wish to create a blank layer to put your own tiles on then see Tilemap.createBlankLayer.
func (self *Tilemap) CreateLayer4O(layer interface{}, width int, height int, group *Group, pixiTest bool) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createLayer", layer, width, height, group, pixiTest)}
}

// CreateLayerI Creates a new TilemapLayer object. By default TilemapLayers are fixed to the camera.
// The `layer` parameter is important. If you've created your map in Tiled then you can get this by looking in Tiled and looking at the Layer name.
// Or you can open the JSON file it exports and look at the layers[].name value. Either way it must match.
// If you wish to create a blank layer to put your own tiles on then see Tilemap.createBlankLayer.
func (self *Tilemap) CreateLayerI(args ...interface{}) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createLayer", args)}
}

// CreateBlankLayer Creates a new and empty layer on this Tilemap. By default TilemapLayers are fixed to the camera.
func (self *Tilemap) CreateBlankLayer(name string, width int, height int, tileWidth int, tileHeight int) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createBlankLayer", name, width, height, tileWidth, tileHeight)}
}

// CreateBlankLayer1O Creates a new and empty layer on this Tilemap. By default TilemapLayers are fixed to the camera.
func (self *Tilemap) CreateBlankLayer1O(name string, width int, height int, tileWidth int, tileHeight int, group *Group) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createBlankLayer", name, width, height, tileWidth, tileHeight, group)}
}

// CreateBlankLayerI Creates a new and empty layer on this Tilemap. By default TilemapLayers are fixed to the camera.
func (self *Tilemap) CreateBlankLayerI(args ...interface{}) *TilemapLayer{
    return &TilemapLayer{self.Object.Call("createBlankLayer", args)}
}

// GetIndex Gets the layer index based on the layers name.
func (self *Tilemap) GetIndex(location []interface{}, name string) int{
    return self.Object.Call("getIndex", location, name).Int()
}

// GetIndexI Gets the layer index based on the layers name.
func (self *Tilemap) GetIndexI(args ...interface{}) int{
    return self.Object.Call("getIndex", args).Int()
}

// GetLayerIndex Gets the layer index based on its name.
func (self *Tilemap) GetLayerIndex(name string) int{
    return self.Object.Call("getLayerIndex", name).Int()
}

// GetLayerIndexI Gets the layer index based on its name.
func (self *Tilemap) GetLayerIndexI(args ...interface{}) int{
    return self.Object.Call("getLayerIndex", args).Int()
}

// GetTilesetIndex Gets the tileset index based on its name.
func (self *Tilemap) GetTilesetIndex(name string) int{
    return self.Object.Call("getTilesetIndex", name).Int()
}

// GetTilesetIndexI Gets the tileset index based on its name.
func (self *Tilemap) GetTilesetIndexI(args ...interface{}) int{
    return self.Object.Call("getTilesetIndex", args).Int()
}

// GetImageIndex Gets the image index based on its name.
func (self *Tilemap) GetImageIndex(name string) int{
    return self.Object.Call("getImageIndex", name).Int()
}

// GetImageIndexI Gets the image index based on its name.
func (self *Tilemap) GetImageIndexI(args ...interface{}) int{
    return self.Object.Call("getImageIndex", args).Int()
}

// SetTileIndexCallback Sets a global collision callback for the given tile index within the layer. This will affect all tiles on this layer that have the same index.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileIndexCallback(indexes interface{}, callback interface{}, callbackContext interface{}) {
    self.Object.Call("setTileIndexCallback", indexes, callback, callbackContext)
}

// SetTileIndexCallback1O Sets a global collision callback for the given tile index within the layer. This will affect all tiles on this layer that have the same index.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileIndexCallback1O(indexes interface{}, callback interface{}, callbackContext interface{}, layer interface{}) {
    self.Object.Call("setTileIndexCallback", indexes, callback, callbackContext, layer)
}

// SetTileIndexCallbackI Sets a global collision callback for the given tile index within the layer. This will affect all tiles on this layer that have the same index.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileIndexCallbackI(args ...interface{}) {
    self.Object.Call("setTileIndexCallback", args)
}

// SetTileLocationCallback Sets a global collision callback for the given map location within the layer. This will affect all tiles on this layer found in the given area.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileLocationCallback(x int, y int, width int, height int, callback interface{}, callbackContext interface{}) {
    self.Object.Call("setTileLocationCallback", x, y, width, height, callback, callbackContext)
}

// SetTileLocationCallback1O Sets a global collision callback for the given map location within the layer. This will affect all tiles on this layer found in the given area.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileLocationCallback1O(x int, y int, width int, height int, callback interface{}, callbackContext interface{}, layer interface{}) {
    self.Object.Call("setTileLocationCallback", x, y, width, height, callback, callbackContext, layer)
}

// SetTileLocationCallbackI Sets a global collision callback for the given map location within the layer. This will affect all tiles on this layer found in the given area.
// If a callback is already set for the tile index it will be replaced. Set the callback to null to remove it.
// If you want to set a callback for a tile at a specific location on the map then see setTileLocationCallback.
func (self *Tilemap) SetTileLocationCallbackI(args ...interface{}) {
    self.Object.Call("setTileLocationCallback", args)
}

// SetCollision Sets collision the given tile or tiles. You can pass in either a single numeric index or an array of indexes: [ 2, 3, 15, 20].
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollision(indexes interface{}) {
    self.Object.Call("setCollision", indexes)
}

// SetCollision1O Sets collision the given tile or tiles. You can pass in either a single numeric index or an array of indexes: [ 2, 3, 15, 20].
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollision1O(indexes interface{}, collides bool) {
    self.Object.Call("setCollision", indexes, collides)
}

// SetCollision2O Sets collision the given tile or tiles. You can pass in either a single numeric index or an array of indexes: [ 2, 3, 15, 20].
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollision2O(indexes interface{}, collides bool, layer interface{}) {
    self.Object.Call("setCollision", indexes, collides, layer)
}

// SetCollision3O Sets collision the given tile or tiles. You can pass in either a single numeric index or an array of indexes: [ 2, 3, 15, 20].
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollision3O(indexes interface{}, collides bool, layer interface{}, recalculate bool) {
    self.Object.Call("setCollision", indexes, collides, layer, recalculate)
}

// SetCollisionI Sets collision the given tile or tiles. You can pass in either a single numeric index or an array of indexes: [ 2, 3, 15, 20].
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionI(args ...interface{}) {
    self.Object.Call("setCollision", args)
}

// SetCollisionBetween Sets collision on a range of tiles where the tile IDs increment sequentially.
// Calling this with a start value of 10 and a stop value of 14 would set collision for tiles 10, 11, 12, 13 and 14.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionBetween(start int, stop int) {
    self.Object.Call("setCollisionBetween", start, stop)
}

// SetCollisionBetween1O Sets collision on a range of tiles where the tile IDs increment sequentially.
// Calling this with a start value of 10 and a stop value of 14 would set collision for tiles 10, 11, 12, 13 and 14.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionBetween1O(start int, stop int, collides bool) {
    self.Object.Call("setCollisionBetween", start, stop, collides)
}

// SetCollisionBetween2O Sets collision on a range of tiles where the tile IDs increment sequentially.
// Calling this with a start value of 10 and a stop value of 14 would set collision for tiles 10, 11, 12, 13 and 14.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionBetween2O(start int, stop int, collides bool, layer interface{}) {
    self.Object.Call("setCollisionBetween", start, stop, collides, layer)
}

// SetCollisionBetween3O Sets collision on a range of tiles where the tile IDs increment sequentially.
// Calling this with a start value of 10 and a stop value of 14 would set collision for tiles 10, 11, 12, 13 and 14.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionBetween3O(start int, stop int, collides bool, layer interface{}, recalculate bool) {
    self.Object.Call("setCollisionBetween", start, stop, collides, layer, recalculate)
}

// SetCollisionBetweenI Sets collision on a range of tiles where the tile IDs increment sequentially.
// Calling this with a start value of 10 and a stop value of 14 would set collision for tiles 10, 11, 12, 13 and 14.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionBetweenI(args ...interface{}) {
    self.Object.Call("setCollisionBetween", args)
}

// SetCollisionByExclusion Sets collision on all tiles in the given layer, except for the IDs of those in the given array.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionByExclusion(indexes []interface{}) {
    self.Object.Call("setCollisionByExclusion", indexes)
}

// SetCollisionByExclusion1O Sets collision on all tiles in the given layer, except for the IDs of those in the given array.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionByExclusion1O(indexes []interface{}, collides bool) {
    self.Object.Call("setCollisionByExclusion", indexes, collides)
}

// SetCollisionByExclusion2O Sets collision on all tiles in the given layer, except for the IDs of those in the given array.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionByExclusion2O(indexes []interface{}, collides bool, layer interface{}) {
    self.Object.Call("setCollisionByExclusion", indexes, collides, layer)
}

// SetCollisionByExclusion3O Sets collision on all tiles in the given layer, except for the IDs of those in the given array.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionByExclusion3O(indexes []interface{}, collides bool, layer interface{}, recalculate bool) {
    self.Object.Call("setCollisionByExclusion", indexes, collides, layer, recalculate)
}

// SetCollisionByExclusionI Sets collision on all tiles in the given layer, except for the IDs of those in the given array.
// The `collides` parameter controls if collision will be enabled (true) or disabled (false).
func (self *Tilemap) SetCollisionByExclusionI(args ...interface{}) {
    self.Object.Call("setCollisionByExclusion", args)
}

// SetCollisionByIndex Sets collision values on a tile in the set.
// You shouldn't usually call this method directly, instead use setCollision, setCollisionBetween or setCollisionByExclusion.
func (self *Tilemap) SetCollisionByIndex(index int) {
    self.Object.Call("setCollisionByIndex", index)
}

// SetCollisionByIndex1O Sets collision values on a tile in the set.
// You shouldn't usually call this method directly, instead use setCollision, setCollisionBetween or setCollisionByExclusion.
func (self *Tilemap) SetCollisionByIndex1O(index int, collides bool) {
    self.Object.Call("setCollisionByIndex", index, collides)
}

// SetCollisionByIndex2O Sets collision values on a tile in the set.
// You shouldn't usually call this method directly, instead use setCollision, setCollisionBetween or setCollisionByExclusion.
func (self *Tilemap) SetCollisionByIndex2O(index int, collides bool, layer int) {
    self.Object.Call("setCollisionByIndex", index, collides, layer)
}

// SetCollisionByIndex3O Sets collision values on a tile in the set.
// You shouldn't usually call this method directly, instead use setCollision, setCollisionBetween or setCollisionByExclusion.
func (self *Tilemap) SetCollisionByIndex3O(index int, collides bool, layer int, recalculate bool) {
    self.Object.Call("setCollisionByIndex", index, collides, layer, recalculate)
}

// SetCollisionByIndexI Sets collision values on a tile in the set.
// You shouldn't usually call this method directly, instead use setCollision, setCollisionBetween or setCollisionByExclusion.
func (self *Tilemap) SetCollisionByIndexI(args ...interface{}) {
    self.Object.Call("setCollisionByIndex", args)
}

// GetLayer Gets the TilemapLayer index as used in the setCollision calls.
func (self *Tilemap) GetLayer(layer interface{}) int{
    return self.Object.Call("getLayer", layer).Int()
}

// GetLayerI Gets the TilemapLayer index as used in the setCollision calls.
func (self *Tilemap) GetLayerI(args ...interface{}) int{
    return self.Object.Call("getLayer", args).Int()
}

// SetPreventRecalculate Turn off/on the recalculation of faces for tile or collision updates. 
// `setPreventRecalculate(true)` puts recalculation on hold while `setPreventRecalculate(false)` recalculates all the changed layers.
func (self *Tilemap) SetPreventRecalculate(value bool) {
    self.Object.Call("setPreventRecalculate", value)
}

// SetPreventRecalculateI Turn off/on the recalculation of faces for tile or collision updates. 
// `setPreventRecalculate(true)` puts recalculation on hold while `setPreventRecalculate(false)` recalculates all the changed layers.
func (self *Tilemap) SetPreventRecalculateI(args ...interface{}) {
    self.Object.Call("setPreventRecalculate", args)
}

// CalculateFaces Internal function.
func (self *Tilemap) CalculateFaces(layer int) {
    self.Object.Call("calculateFaces", layer)
}

// CalculateFacesI Internal function.
func (self *Tilemap) CalculateFacesI(args ...interface{}) {
    self.Object.Call("calculateFaces", args)
}

// GetTileAbove Gets the tile above the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileAbove(layer int, x int, y int) {
    self.Object.Call("getTileAbove", layer, x, y)
}

// GetTileAboveI Gets the tile above the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileAboveI(args ...interface{}) {
    self.Object.Call("getTileAbove", args)
}

// GetTileBelow Gets the tile below the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileBelow(layer int, x int, y int) {
    self.Object.Call("getTileBelow", layer, x, y)
}

// GetTileBelowI Gets the tile below the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileBelowI(args ...interface{}) {
    self.Object.Call("getTileBelow", args)
}

// GetTileLeft Gets the tile to the left of the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileLeft(layer int, x int, y int) {
    self.Object.Call("getTileLeft", layer, x, y)
}

// GetTileLeftI Gets the tile to the left of the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileLeftI(args ...interface{}) {
    self.Object.Call("getTileLeft", args)
}

// GetTileRight Gets the tile to the right of the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileRight(layer int, x int, y int) {
    self.Object.Call("getTileRight", layer, x, y)
}

// GetTileRightI Gets the tile to the right of the tile coordinates given.
// Mostly used as an internal function by calculateFaces.
func (self *Tilemap) GetTileRightI(args ...interface{}) {
    self.Object.Call("getTileRight", args)
}

// SetLayer Sets the current layer to the given index.
func (self *Tilemap) SetLayer(layer interface{}) {
    self.Object.Call("setLayer", layer)
}

// SetLayerI Sets the current layer to the given index.
func (self *Tilemap) SetLayerI(args ...interface{}) {
    self.Object.Call("setLayer", args)
}

// HasTile Checks if there is a tile at the given location.
func (self *Tilemap) HasTile(x int, y int, layer interface{}) bool{
    return self.Object.Call("hasTile", x, y, layer).Bool()
}

// HasTileI Checks if there is a tile at the given location.
func (self *Tilemap) HasTileI(args ...interface{}) bool{
    return self.Object.Call("hasTile", args).Bool()
}

// RemoveTile Removes the tile located at the given coordinates and updates the collision data.
func (self *Tilemap) RemoveTile(x int, y int) *Tile{
    return &Tile{self.Object.Call("removeTile", x, y)}
}

// RemoveTile1O Removes the tile located at the given coordinates and updates the collision data.
func (self *Tilemap) RemoveTile1O(x int, y int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("removeTile", x, y, layer)}
}

// RemoveTileI Removes the tile located at the given coordinates and updates the collision data.
func (self *Tilemap) RemoveTileI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("removeTile", args)}
}

// RemoveTileWorldXY Removes the tile located at the given coordinates and updates the collision data. The coordinates are given in pixel values.
func (self *Tilemap) RemoveTileWorldXY(x int, y int, tileWidth int, tileHeight int) *Tile{
    return &Tile{self.Object.Call("removeTileWorldXY", x, y, tileWidth, tileHeight)}
}

// RemoveTileWorldXY1O Removes the tile located at the given coordinates and updates the collision data. The coordinates are given in pixel values.
func (self *Tilemap) RemoveTileWorldXY1O(x int, y int, tileWidth int, tileHeight int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("removeTileWorldXY", x, y, tileWidth, tileHeight, layer)}
}

// RemoveTileWorldXYI Removes the tile located at the given coordinates and updates the collision data. The coordinates are given in pixel values.
func (self *Tilemap) RemoveTileWorldXYI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("removeTileWorldXY", args)}
}

// PutTile Puts a tile of the given index value at the coordinate specified.
// If you pass `null` as the tile it will pass your call over to Tilemap.removeTile instead.
func (self *Tilemap) PutTile(tile interface{}, x int, y int) *Tile{
    return &Tile{self.Object.Call("putTile", tile, x, y)}
}

// PutTile1O Puts a tile of the given index value at the coordinate specified.
// If you pass `null` as the tile it will pass your call over to Tilemap.removeTile instead.
func (self *Tilemap) PutTile1O(tile interface{}, x int, y int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("putTile", tile, x, y, layer)}
}

// PutTileI Puts a tile of the given index value at the coordinate specified.
// If you pass `null` as the tile it will pass your call over to Tilemap.removeTile instead.
func (self *Tilemap) PutTileI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("putTile", args)}
}

// PutTileWorldXY Puts a tile into the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) PutTileWorldXY(tile interface{}, x int, y int, tileWidth int, tileHeight int) *Tile{
    return &Tile{self.Object.Call("putTileWorldXY", tile, x, y, tileWidth, tileHeight)}
}

// PutTileWorldXY1O Puts a tile into the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) PutTileWorldXY1O(tile interface{}, x int, y int, tileWidth int, tileHeight int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("putTileWorldXY", tile, x, y, tileWidth, tileHeight, layer)}
}

// PutTileWorldXYI Puts a tile into the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) PutTileWorldXYI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("putTileWorldXY", args)}
}

// SearchTileIndex Searches the entire map layer for the first tile matching the given index, then returns that Phaser.Tile object.
// If no match is found it returns null.
// The search starts from the top-left tile and continues horizontally until it hits the end of the row, then it drops down to the next column.
// If the reverse boolean is true, it scans starting from the bottom-right corner traveling up to the top-left.
func (self *Tilemap) SearchTileIndex(index int) *Tile{
    return &Tile{self.Object.Call("searchTileIndex", index)}
}

// SearchTileIndex1O Searches the entire map layer for the first tile matching the given index, then returns that Phaser.Tile object.
// If no match is found it returns null.
// The search starts from the top-left tile and continues horizontally until it hits the end of the row, then it drops down to the next column.
// If the reverse boolean is true, it scans starting from the bottom-right corner traveling up to the top-left.
func (self *Tilemap) SearchTileIndex1O(index int, skip int) *Tile{
    return &Tile{self.Object.Call("searchTileIndex", index, skip)}
}

// SearchTileIndex2O Searches the entire map layer for the first tile matching the given index, then returns that Phaser.Tile object.
// If no match is found it returns null.
// The search starts from the top-left tile and continues horizontally until it hits the end of the row, then it drops down to the next column.
// If the reverse boolean is true, it scans starting from the bottom-right corner traveling up to the top-left.
func (self *Tilemap) SearchTileIndex2O(index int, skip int, reverse int) *Tile{
    return &Tile{self.Object.Call("searchTileIndex", index, skip, reverse)}
}

// SearchTileIndex3O Searches the entire map layer for the first tile matching the given index, then returns that Phaser.Tile object.
// If no match is found it returns null.
// The search starts from the top-left tile and continues horizontally until it hits the end of the row, then it drops down to the next column.
// If the reverse boolean is true, it scans starting from the bottom-right corner traveling up to the top-left.
func (self *Tilemap) SearchTileIndex3O(index int, skip int, reverse int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("searchTileIndex", index, skip, reverse, layer)}
}

// SearchTileIndexI Searches the entire map layer for the first tile matching the given index, then returns that Phaser.Tile object.
// If no match is found it returns null.
// The search starts from the top-left tile and continues horizontally until it hits the end of the row, then it drops down to the next column.
// If the reverse boolean is true, it scans starting from the bottom-right corner traveling up to the top-left.
func (self *Tilemap) SearchTileIndexI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("searchTileIndex", args)}
}

// GetTile Gets a tile from the Tilemap Layer. The coordinates are given in tile values.
func (self *Tilemap) GetTile(x int, y int) *Tile{
    return &Tile{self.Object.Call("getTile", x, y)}
}

// GetTile1O Gets a tile from the Tilemap Layer. The coordinates are given in tile values.
func (self *Tilemap) GetTile1O(x int, y int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("getTile", x, y, layer)}
}

// GetTile2O Gets a tile from the Tilemap Layer. The coordinates are given in tile values.
func (self *Tilemap) GetTile2O(x int, y int, layer interface{}, nonNull bool) *Tile{
    return &Tile{self.Object.Call("getTile", x, y, layer, nonNull)}
}

// GetTileI Gets a tile from the Tilemap Layer. The coordinates are given in tile values.
func (self *Tilemap) GetTileI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("getTile", args)}
}

// GetTileWorldXY Gets a tile from the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) GetTileWorldXY(x int, y int) *Tile{
    return &Tile{self.Object.Call("getTileWorldXY", x, y)}
}

// GetTileWorldXY1O Gets a tile from the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) GetTileWorldXY1O(x int, y int, tileWidth int) *Tile{
    return &Tile{self.Object.Call("getTileWorldXY", x, y, tileWidth)}
}

// GetTileWorldXY2O Gets a tile from the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) GetTileWorldXY2O(x int, y int, tileWidth int, tileHeight int) *Tile{
    return &Tile{self.Object.Call("getTileWorldXY", x, y, tileWidth, tileHeight)}
}

// GetTileWorldXY3O Gets a tile from the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) GetTileWorldXY3O(x int, y int, tileWidth int, tileHeight int, layer interface{}) *Tile{
    return &Tile{self.Object.Call("getTileWorldXY", x, y, tileWidth, tileHeight, layer)}
}

// GetTileWorldXY4O Gets a tile from the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) GetTileWorldXY4O(x int, y int, tileWidth int, tileHeight int, layer interface{}, nonNull bool) *Tile{
    return &Tile{self.Object.Call("getTileWorldXY", x, y, tileWidth, tileHeight, layer, nonNull)}
}

// GetTileWorldXYI Gets a tile from the Tilemap layer. The coordinates are given in pixel values.
func (self *Tilemap) GetTileWorldXYI(args ...interface{}) *Tile{
    return &Tile{self.Object.Call("getTileWorldXY", args)}
}

// Copy Copies all of the tiles in the given rectangular block into the tilemap data buffer.
func (self *Tilemap) Copy(x int, y int, width int, height int) []interface{}{
	array00 := self.Object.Call("copy", x, y, width, height)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Copy1O Copies all of the tiles in the given rectangular block into the tilemap data buffer.
func (self *Tilemap) Copy1O(x int, y int, width int, height int, layer interface{}) []interface{}{
	array00 := self.Object.Call("copy", x, y, width, height, layer)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// CopyI Copies all of the tiles in the given rectangular block into the tilemap data buffer.
func (self *Tilemap) CopyI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("copy", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Paste Pastes a previously copied block of tile data into the given x/y coordinates. Data should have been prepared with Tilemap.copy.
func (self *Tilemap) Paste(x int, y int, tileblock []interface{}) {
    self.Object.Call("paste", x, y, tileblock)
}

// Paste1O Pastes a previously copied block of tile data into the given x/y coordinates. Data should have been prepared with Tilemap.copy.
func (self *Tilemap) Paste1O(x int, y int, tileblock []interface{}, layer interface{}) {
    self.Object.Call("paste", x, y, tileblock, layer)
}

// PasteI Pastes a previously copied block of tile data into the given x/y coordinates. Data should have been prepared with Tilemap.copy.
func (self *Tilemap) PasteI(args ...interface{}) {
    self.Object.Call("paste", args)
}

// Swap Scans the given area for tiles with an index matching tileA and swaps them with tileB.
func (self *Tilemap) Swap(tileA int, tileB int, x int, y int, width int, height int) {
    self.Object.Call("swap", tileA, tileB, x, y, width, height)
}

// Swap1O Scans the given area for tiles with an index matching tileA and swaps them with tileB.
func (self *Tilemap) Swap1O(tileA int, tileB int, x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("swap", tileA, tileB, x, y, width, height, layer)
}

// SwapI Scans the given area for tiles with an index matching tileA and swaps them with tileB.
func (self *Tilemap) SwapI(args ...interface{}) {
    self.Object.Call("swap", args)
}

// SwapHandler Internal function that handles the swapping of tiles.
func (self *Tilemap) SwapHandler(value int) {
    self.Object.Call("swapHandler", value)
}

// SwapHandlerI Internal function that handles the swapping of tiles.
func (self *Tilemap) SwapHandlerI(args ...interface{}) {
    self.Object.Call("swapHandler", args)
}

// ForEach For each tile in the given area defined by x/y and width/height run the given callback.
func (self *Tilemap) ForEach(callback int, context int, x int, y int, width int, height int) {
    self.Object.Call("forEach", callback, context, x, y, width, height)
}

// ForEach1O For each tile in the given area defined by x/y and width/height run the given callback.
func (self *Tilemap) ForEach1O(callback int, context int, x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("forEach", callback, context, x, y, width, height, layer)
}

// ForEachI For each tile in the given area defined by x/y and width/height run the given callback.
func (self *Tilemap) ForEachI(args ...interface{}) {
    self.Object.Call("forEach", args)
}

// Replace Scans the given area for tiles with an index matching `source` and updates their index to match `dest`.
func (self *Tilemap) Replace(source int, dest int, x int, y int, width int, height int) {
    self.Object.Call("replace", source, dest, x, y, width, height)
}

// Replace1O Scans the given area for tiles with an index matching `source` and updates their index to match `dest`.
func (self *Tilemap) Replace1O(source int, dest int, x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("replace", source, dest, x, y, width, height, layer)
}

// ReplaceI Scans the given area for tiles with an index matching `source` and updates their index to match `dest`.
func (self *Tilemap) ReplaceI(args ...interface{}) {
    self.Object.Call("replace", args)
}

// Random Randomises a set of tiles in a given area.
func (self *Tilemap) Random(x int, y int, width int, height int) {
    self.Object.Call("random", x, y, width, height)
}

// Random1O Randomises a set of tiles in a given area.
func (self *Tilemap) Random1O(x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("random", x, y, width, height, layer)
}

// RandomI Randomises a set of tiles in a given area.
func (self *Tilemap) RandomI(args ...interface{}) {
    self.Object.Call("random", args)
}

// Shuffle Shuffles a set of tiles in a given area. It will only randomise the tiles in that area, so if they're all the same nothing will appear to have changed!
func (self *Tilemap) Shuffle(x int, y int, width int, height int) {
    self.Object.Call("shuffle", x, y, width, height)
}

// Shuffle1O Shuffles a set of tiles in a given area. It will only randomise the tiles in that area, so if they're all the same nothing will appear to have changed!
func (self *Tilemap) Shuffle1O(x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("shuffle", x, y, width, height, layer)
}

// ShuffleI Shuffles a set of tiles in a given area. It will only randomise the tiles in that area, so if they're all the same nothing will appear to have changed!
func (self *Tilemap) ShuffleI(args ...interface{}) {
    self.Object.Call("shuffle", args)
}

// Fill Fills the given area with the specified tile.
func (self *Tilemap) Fill(index int, x int, y int, width int, height int) {
    self.Object.Call("fill", index, x, y, width, height)
}

// Fill1O Fills the given area with the specified tile.
func (self *Tilemap) Fill1O(index int, x int, y int, width int, height int, layer interface{}) {
    self.Object.Call("fill", index, x, y, width, height, layer)
}

// FillI Fills the given area with the specified tile.
func (self *Tilemap) FillI(args ...interface{}) {
    self.Object.Call("fill", args)
}

// RemoveAllLayers Removes all layers from this tile map.
func (self *Tilemap) RemoveAllLayers() {
    self.Object.Call("removeAllLayers")
}

// RemoveAllLayersI Removes all layers from this tile map.
func (self *Tilemap) RemoveAllLayersI(args ...interface{}) {
    self.Object.Call("removeAllLayers", args)
}

// Dump Dumps the tilemap data out to the console.
func (self *Tilemap) Dump() {
    self.Object.Call("dump")
}

// DumpI Dumps the tilemap data out to the console.
func (self *Tilemap) DumpI(args ...interface{}) {
    self.Object.Call("dump", args)
}

// Destroy Removes all layer data from this tile map and nulls the game reference.
// Note: You are responsible for destroying any TilemapLayer objects you generated yourself, as Tilemap doesn't keep a reference to them.
func (self *Tilemap) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Removes all layer data from this tile map and nulls the game reference.
// Note: You are responsible for destroying any TilemapLayer objects you generated yourself, as Tilemap doesn't keep a reference to them.
func (self *Tilemap) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

