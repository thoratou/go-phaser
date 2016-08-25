// Package phaser Automatic generation for Phaser.Cache
// generated file Cache.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// Cache Phaser has one single cache in which it stores all assets.
// 
// The cache is split up into sections, such as images, sounds, video, json, etc. All assets are stored using
// a unique string-based key as their identifier. Assets stored in different areas of the cache can have the
// same key, for example 'playerWalking' could be used as the key for both a sprite sheet and an audio file,
// because they are unique data types.
// 
// The cache is automatically populated by the Phaser.Loader. When you use the loader to pull in external assets
// such as images they are automatically placed into their respective cache. Most common Game Objects, such as
// Sprites and Videos automatically query the cache to extract the assets they need on instantiation.
// 
// You can access the cache from within a State via `this.cache`. From here you can call any public method it has,
// including adding new entries to it, deleting them or querying them.
// 
// Understand that almost without exception when you get an item from the cache it will return a reference to the
// item stored in the cache, not a copy of it. Therefore if you retrieve an item and then modify it, the original
// object in the cache will also be updated, even if you don't put it back into the cache again.
// 
// By default when you change State the cache is _not_ cleared, although there is an option to clear it should
// your game require it. In a typical game set-up the cache is populated once after the main game has loaded and
// then used as an asset store.
type Cache struct {
    *js.Object
}

// NewCache Phaser has one single cache in which it stores all assets.
// 
// The cache is split up into sections, such as images, sounds, video, json, etc. All assets are stored using
// a unique string-based key as their identifier. Assets stored in different areas of the cache can have the
// same key, for example 'playerWalking' could be used as the key for both a sprite sheet and an audio file,
// because they are unique data types.
// 
// The cache is automatically populated by the Phaser.Loader. When you use the loader to pull in external assets
// such as images they are automatically placed into their respective cache. Most common Game Objects, such as
// Sprites and Videos automatically query the cache to extract the assets they need on instantiation.
// 
// You can access the cache from within a State via `this.cache`. From here you can call any public method it has,
// including adding new entries to it, deleting them or querying them.
// 
// Understand that almost without exception when you get an item from the cache it will return a reference to the
// item stored in the cache, not a copy of it. Therefore if you retrieve an item and then modify it, the original
// object in the cache will also be updated, even if you don't put it back into the cache again.
// 
// By default when you change State the cache is _not_ cleared, although there is an option to clear it should
// your game require it. In a typical game set-up the cache is populated once after the main game has loaded and
// then used as an asset store.
func NewCache(game *Game) *Cache {
    return &Cache{js.Global.Get("Phaser").Get("Cache").New(game)}
}
// NewCacheI Phaser has one single cache in which it stores all assets.
// 
// The cache is split up into sections, such as images, sounds, video, json, etc. All assets are stored using
// a unique string-based key as their identifier. Assets stored in different areas of the cache can have the
// same key, for example 'playerWalking' could be used as the key for both a sprite sheet and an audio file,
// because they are unique data types.
// 
// The cache is automatically populated by the Phaser.Loader. When you use the loader to pull in external assets
// such as images they are automatically placed into their respective cache. Most common Game Objects, such as
// Sprites and Videos automatically query the cache to extract the assets they need on instantiation.
// 
// You can access the cache from within a State via `this.cache`. From here you can call any public method it has,
// including adding new entries to it, deleting them or querying them.
// 
// Understand that almost without exception when you get an item from the cache it will return a reference to the
// item stored in the cache, not a copy of it. Therefore if you retrieve an item and then modify it, the original
// object in the cache will also be updated, even if you don't put it back into the cache again.
// 
// By default when you change State the cache is _not_ cleared, although there is an option to clear it should
// your game require it. In a typical game set-up the cache is populated once after the main game has loaded and
// then used as an asset store.
func NewCacheI(args ...interface{}) *Cache {
    return &Cache{js.Global.Get("Phaser").Get("Cache").New(args)}
}



// Cache Binding conversion method to Cache point 
func ToCache(jsStruct interface{}) *Cache {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Cache{Object: object}
	}
	return nil
}



// Game Local reference to game.
func (self *Cache) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *Cache) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// AutoResolveURL Automatically resolve resource URLs to absolute paths for use with the Cache.getURL method.
func (self *Cache) AutoResolveURL() bool{
    return self.Object.Get("autoResolveURL").Bool()
}

// SetAutoResolveURLA Automatically resolve resource URLs to absolute paths for use with the Cache.getURL method.
func (self *Cache) SetAutoResolveURLA(member bool) {
    self.Object.Set("autoResolveURL", member)
}

// OnSoundUnlock This event is dispatched when the sound system is unlocked via a touch event on cellular devices.
func (self *Cache) OnSoundUnlock() *Signal{
    return &Signal{self.Object.Get("onSoundUnlock")}
}

// SetOnSoundUnlockA This event is dispatched when the sound system is unlocked via a touch event on cellular devices.
func (self *Cache) SetOnSoundUnlockA(member *Signal) {
    self.Object.Set("onSoundUnlock", member)
}

// CANVAS empty description
func (self *Cache) CANVAS() int{
    return self.Object.Get("CANVAS").Int()
}

// SetCANVASA empty description
func (self *Cache) SetCANVASA(member int) {
    self.Object.Set("CANVAS", member)
}

// IMAGE empty description
func (self *Cache) IMAGE() int{
    return self.Object.Get("IMAGE").Int()
}

// SetIMAGEA empty description
func (self *Cache) SetIMAGEA(member int) {
    self.Object.Set("IMAGE", member)
}

// TEXTURE empty description
func (self *Cache) TEXTURE() int{
    return self.Object.Get("TEXTURE").Int()
}

// SetTEXTUREA empty description
func (self *Cache) SetTEXTUREA(member int) {
    self.Object.Set("TEXTURE", member)
}

// SOUND empty description
func (self *Cache) SOUND() int{
    return self.Object.Get("SOUND").Int()
}

// SetSOUNDA empty description
func (self *Cache) SetSOUNDA(member int) {
    self.Object.Set("SOUND", member)
}

// TEXT empty description
func (self *Cache) TEXT() int{
    return self.Object.Get("TEXT").Int()
}

// SetTEXTA empty description
func (self *Cache) SetTEXTA(member int) {
    self.Object.Set("TEXT", member)
}

// PHYSICS empty description
func (self *Cache) PHYSICS() int{
    return self.Object.Get("PHYSICS").Int()
}

// SetPHYSICSA empty description
func (self *Cache) SetPHYSICSA(member int) {
    self.Object.Set("PHYSICS", member)
}

// TILEMAP empty description
func (self *Cache) TILEMAP() int{
    return self.Object.Get("TILEMAP").Int()
}

// SetTILEMAPA empty description
func (self *Cache) SetTILEMAPA(member int) {
    self.Object.Set("TILEMAP", member)
}

// BINARY empty description
func (self *Cache) BINARY() int{
    return self.Object.Get("BINARY").Int()
}

// SetBINARYA empty description
func (self *Cache) SetBINARYA(member int) {
    self.Object.Set("BINARY", member)
}

// BITMAPDATA empty description
func (self *Cache) BITMAPDATA() int{
    return self.Object.Get("BITMAPDATA").Int()
}

// SetBITMAPDATAA empty description
func (self *Cache) SetBITMAPDATAA(member int) {
    self.Object.Set("BITMAPDATA", member)
}

// BITMAPFONT empty description
func (self *Cache) BITMAPFONT() int{
    return self.Object.Get("BITMAPFONT").Int()
}

// SetBITMAPFONTA empty description
func (self *Cache) SetBITMAPFONTA(member int) {
    self.Object.Set("BITMAPFONT", member)
}

// JSON empty description
func (self *Cache) JSON() int{
    return self.Object.Get("JSON").Int()
}

// SetJSONA empty description
func (self *Cache) SetJSONA(member int) {
    self.Object.Set("JSON", member)
}

// XML empty description
func (self *Cache) XML() int{
    return self.Object.Get("XML").Int()
}

// SetXMLA empty description
func (self *Cache) SetXMLA(member int) {
    self.Object.Set("XML", member)
}

// VIDEO empty description
func (self *Cache) VIDEO() int{
    return self.Object.Get("VIDEO").Int()
}

// SetVIDEOA empty description
func (self *Cache) SetVIDEOA(member int) {
    self.Object.Set("VIDEO", member)
}

// SHADER empty description
func (self *Cache) SHADER() int{
    return self.Object.Get("SHADER").Int()
}

// SetSHADERA empty description
func (self *Cache) SetSHADERA(member int) {
    self.Object.Set("SHADER", member)
}

// RENDER_TEXTURE empty description
func (self *Cache) RENDER_TEXTURE() int{
    return self.Object.Get("RENDER_TEXTURE").Int()
}

// SetRENDER_TEXTUREA empty description
func (self *Cache) SetRENDER_TEXTUREA(member int) {
    self.Object.Set("RENDER_TEXTURE", member)
}

// DEFAULT The default image used for a texture when no other is specified.
func (self *Cache) DEFAULT() *Texture{
    return &Texture{self.Object.Get("DEFAULT")}
}

// SetDEFAULTA The default image used for a texture when no other is specified.
func (self *Cache) SetDEFAULTA(member *Texture) {
    self.Object.Set("DEFAULT", member)
}

// MISSING The default image used for a texture when the source image is missing.
func (self *Cache) MISSING() *Texture{
    return &Texture{self.Object.Get("MISSING")}
}

// SetMISSINGA The default image used for a texture when the source image is missing.
func (self *Cache) SetMISSINGA(member *Texture) {
    self.Object.Set("MISSING", member)
}


// AddCanvas Add a new canvas object in to the cache.
func (self *Cache) AddCanvas(key string, canvas *dom.HTMLCanvasElement) {
    self.Object.Call("addCanvas", key, canvas)
}

// AddCanvas1O Add a new canvas object in to the cache.
func (self *Cache) AddCanvas1O(key string, canvas *dom.HTMLCanvasElement, context *dom.CanvasRenderingContext2D) {
    self.Object.Call("addCanvas", key, canvas, context)
}

// AddCanvasI Add a new canvas object in to the cache.
func (self *Cache) AddCanvasI(args ...interface{}) {
    self.Object.Call("addCanvas", args)
}

// AddImage Adds an Image file into the Cache. The file must have already been loaded, typically via Phaser.Loader, but can also have been loaded into the DOM.
// If an image already exists in the cache with the same key then it is removed and destroyed, and the new image inserted in its place.
func (self *Cache) AddImage(key string, url string, data interface{}) interface{}{
    return self.Object.Call("addImage", key, url, data)
}

// AddImageI Adds an Image file into the Cache. The file must have already been loaded, typically via Phaser.Loader, but can also have been loaded into the DOM.
// If an image already exists in the cache with the same key then it is removed and destroyed, and the new image inserted in its place.
func (self *Cache) AddImageI(args ...interface{}) interface{}{
    return self.Object.Call("addImage", args)
}

// AddDefaultImage Adds a default image to be used in special cases such as WebGL Filters.
// It uses the special reserved key of `__default`.
// This method is called automatically when the Cache is created.
// This image is skipped when `Cache.destroy` is called due to its internal requirements.
func (self *Cache) AddDefaultImage() {
    self.Object.Call("addDefaultImage")
}

// AddDefaultImageI Adds a default image to be used in special cases such as WebGL Filters.
// It uses the special reserved key of `__default`.
// This method is called automatically when the Cache is created.
// This image is skipped when `Cache.destroy` is called due to its internal requirements.
func (self *Cache) AddDefaultImageI(args ...interface{}) {
    self.Object.Call("addDefaultImage", args)
}

// AddMissingImage Adds an image to be used when a key is wrong / missing.
// It uses the special reserved key of `__missing`.
// This method is called automatically when the Cache is created.
// This image is skipped when `Cache.destroy` is called due to its internal requirements.
func (self *Cache) AddMissingImage() {
    self.Object.Call("addMissingImage")
}

// AddMissingImageI Adds an image to be used when a key is wrong / missing.
// It uses the special reserved key of `__missing`.
// This method is called automatically when the Cache is created.
// This image is skipped when `Cache.destroy` is called due to its internal requirements.
func (self *Cache) AddMissingImageI(args ...interface{}) {
    self.Object.Call("addMissingImage", args)
}

// AddSound Adds a Sound file into the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddSound(key string, url string, data interface{}, webAudio bool, audioTag bool) {
    self.Object.Call("addSound", key, url, data, webAudio, audioTag)
}

// AddSoundI Adds a Sound file into the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddSoundI(args ...interface{}) {
    self.Object.Call("addSound", args)
}

// AddText Add a new text data.
func (self *Cache) AddText(key string, url string, data interface{}) {
    self.Object.Call("addText", key, url, data)
}

// AddTextI Add a new text data.
func (self *Cache) AddTextI(args ...interface{}) {
    self.Object.Call("addText", args)
}

// AddPhysicsData Add a new physics data object to the Cache.
func (self *Cache) AddPhysicsData(key string, url string, JSONData interface{}, format int) {
    self.Object.Call("addPhysicsData", key, url, JSONData, format)
}

// AddPhysicsDataI Add a new physics data object to the Cache.
func (self *Cache) AddPhysicsDataI(args ...interface{}) {
    self.Object.Call("addPhysicsData", args)
}

// AddTilemap Add a new tilemap to the Cache.
func (self *Cache) AddTilemap(key string, url string, mapData interface{}, format int) {
    self.Object.Call("addTilemap", key, url, mapData, format)
}

// AddTilemapI Add a new tilemap to the Cache.
func (self *Cache) AddTilemapI(args ...interface{}) {
    self.Object.Call("addTilemap", args)
}

// AddBinary Add a binary object in to the cache.
func (self *Cache) AddBinary(key string, binaryData interface{}) {
    self.Object.Call("addBinary", key, binaryData)
}

// AddBinaryI Add a binary object in to the cache.
func (self *Cache) AddBinaryI(args ...interface{}) {
    self.Object.Call("addBinary", args)
}

// AddBitmapData Add a BitmapData object to the cache.
func (self *Cache) AddBitmapData(key string, bitmapData *BitmapData) *BitmapData{
    return &BitmapData{self.Object.Call("addBitmapData", key, bitmapData)}
}

// AddBitmapData1O Add a BitmapData object to the cache.
func (self *Cache) AddBitmapData1O(key string, bitmapData *BitmapData, frameData interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("addBitmapData", key, bitmapData, frameData)}
}

// AddBitmapDataI Add a BitmapData object to the cache.
func (self *Cache) AddBitmapDataI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("addBitmapData", args)}
}

// AddBitmapFont Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFont(key string, url string, data interface{}, atlasData interface{}) {
    self.Object.Call("addBitmapFont", key, url, data, atlasData)
}

// AddBitmapFont1O Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFont1O(key string, url string, data interface{}, atlasData interface{}, atlasType string) {
    self.Object.Call("addBitmapFont", key, url, data, atlasData, atlasType)
}

// AddBitmapFont2O Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFont2O(key string, url string, data interface{}, atlasData interface{}, atlasType string, xSpacing int) {
    self.Object.Call("addBitmapFont", key, url, data, atlasData, atlasType, xSpacing)
}

// AddBitmapFont3O Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFont3O(key string, url string, data interface{}, atlasData interface{}, atlasType string, xSpacing int, ySpacing int) {
    self.Object.Call("addBitmapFont", key, url, data, atlasData, atlasType, xSpacing, ySpacing)
}

// AddBitmapFontI Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFontI(args ...interface{}) {
    self.Object.Call("addBitmapFont", args)
}

// AddJSON Add a new json object into the cache.
func (self *Cache) AddJSON(key string, url string, data interface{}) {
    self.Object.Call("addJSON", key, url, data)
}

// AddJSONI Add a new json object into the cache.
func (self *Cache) AddJSONI(args ...interface{}) {
    self.Object.Call("addJSON", args)
}

// AddXML Add a new xml object into the cache.
func (self *Cache) AddXML(key string, url string, data interface{}) {
    self.Object.Call("addXML", key, url, data)
}

// AddXMLI Add a new xml object into the cache.
func (self *Cache) AddXMLI(args ...interface{}) {
    self.Object.Call("addXML", args)
}

// AddVideo Adds a Video file into the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddVideo(key string, url string, data interface{}, isBlob bool) {
    self.Object.Call("addVideo", key, url, data, isBlob)
}

// AddVideoI Adds a Video file into the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddVideoI(args ...interface{}) {
    self.Object.Call("addVideo", args)
}

// AddShader Adds a Fragment Shader in to the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddShader(key string, url string, data interface{}) {
    self.Object.Call("addShader", key, url, data)
}

// AddShaderI Adds a Fragment Shader in to the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddShaderI(args ...interface{}) {
    self.Object.Call("addShader", args)
}

// AddRenderTexture Add a new Phaser.RenderTexture in to the cache.
func (self *Cache) AddRenderTexture(key string, texture *RenderTexture) {
    self.Object.Call("addRenderTexture", key, texture)
}

// AddRenderTextureI Add a new Phaser.RenderTexture in to the cache.
func (self *Cache) AddRenderTextureI(args ...interface{}) {
    self.Object.Call("addRenderTexture", args)
}

// AddSpriteSheet Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheet(key string, url string, data interface{}, frameWidth int, frameHeight int) {
    self.Object.Call("addSpriteSheet", key, url, data, frameWidth, frameHeight)
}

// AddSpriteSheet1O Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheet1O(key string, url string, data interface{}, frameWidth int, frameHeight int, frameMax int) {
    self.Object.Call("addSpriteSheet", key, url, data, frameWidth, frameHeight, frameMax)
}

// AddSpriteSheet2O Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheet2O(key string, url string, data interface{}, frameWidth int, frameHeight int, frameMax int, margin int) {
    self.Object.Call("addSpriteSheet", key, url, data, frameWidth, frameHeight, frameMax, margin)
}

// AddSpriteSheet3O Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheet3O(key string, url string, data interface{}, frameWidth int, frameHeight int, frameMax int, margin int, spacing int) {
    self.Object.Call("addSpriteSheet", key, url, data, frameWidth, frameHeight, frameMax, margin, spacing)
}

// AddSpriteSheetI Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheetI(args ...interface{}) {
    self.Object.Call("addSpriteSheet", args)
}

// AddTextureAtlas Add a new texture atlas to the Cache.
func (self *Cache) AddTextureAtlas(key string, url string, data interface{}, atlasData interface{}, format int) {
    self.Object.Call("addTextureAtlas", key, url, data, atlasData, format)
}

// AddTextureAtlasI Add a new texture atlas to the Cache.
func (self *Cache) AddTextureAtlasI(args ...interface{}) {
    self.Object.Call("addTextureAtlas", args)
}

// ReloadSound Reload a Sound file from the server.
func (self *Cache) ReloadSound(key string) {
    self.Object.Call("reloadSound", key)
}

// ReloadSoundI Reload a Sound file from the server.
func (self *Cache) ReloadSoundI(args ...interface{}) {
    self.Object.Call("reloadSound", args)
}

// ReloadSoundComplete Fires the onSoundUnlock event when the sound has completed reloading.
func (self *Cache) ReloadSoundComplete(key string) {
    self.Object.Call("reloadSoundComplete", key)
}

// ReloadSoundCompleteI Fires the onSoundUnlock event when the sound has completed reloading.
func (self *Cache) ReloadSoundCompleteI(args ...interface{}) {
    self.Object.Call("reloadSoundComplete", args)
}

// UpdateSound Updates the sound object in the cache.
func (self *Cache) UpdateSound(key string) {
    self.Object.Call("updateSound", key)
}

// UpdateSoundI Updates the sound object in the cache.
func (self *Cache) UpdateSoundI(args ...interface{}) {
    self.Object.Call("updateSound", args)
}

// DecodedSound Add a new decoded sound.
func (self *Cache) DecodedSound(key string, data interface{}) {
    self.Object.Call("decodedSound", key, data)
}

// DecodedSoundI Add a new decoded sound.
func (self *Cache) DecodedSoundI(args ...interface{}) {
    self.Object.Call("decodedSound", args)
}

// IsSoundDecoded Check if the given sound has finished decoding.
func (self *Cache) IsSoundDecoded(key string) bool{
    return self.Object.Call("isSoundDecoded", key).Bool()
}

// IsSoundDecodedI Check if the given sound has finished decoding.
func (self *Cache) IsSoundDecodedI(args ...interface{}) bool{
    return self.Object.Call("isSoundDecoded", args).Bool()
}

// IsSoundReady Check if the given sound is ready for playback.
// A sound is considered ready when it has finished decoding and the device is no longer touch locked.
func (self *Cache) IsSoundReady(key string) bool{
    return self.Object.Call("isSoundReady", key).Bool()
}

// IsSoundReadyI Check if the given sound is ready for playback.
// A sound is considered ready when it has finished decoding and the device is no longer touch locked.
func (self *Cache) IsSoundReadyI(args ...interface{}) bool{
    return self.Object.Call("isSoundReady", args).Bool()
}

// CheckKey Checks if a key for the given cache object type exists.
func (self *Cache) CheckKey(cache int, key string) bool{
    return self.Object.Call("checkKey", cache, key).Bool()
}

// CheckKeyI Checks if a key for the given cache object type exists.
func (self *Cache) CheckKeyI(args ...interface{}) bool{
    return self.Object.Call("checkKey", args).Bool()
}

// CheckURL Checks if the given URL has been loaded into the Cache.
// This method will only work if Cache.autoResolveURL was set to `true` before any preloading took place.
// The method will make a DOM src call to the URL given, so please be aware of this for certain file types, such as Sound files on Firefox
// which may cause double-load instances.
func (self *Cache) CheckURL(url string) bool{
    return self.Object.Call("checkURL", url).Bool()
}

// CheckURLI Checks if the given URL has been loaded into the Cache.
// This method will only work if Cache.autoResolveURL was set to `true` before any preloading took place.
// The method will make a DOM src call to the URL given, so please be aware of this for certain file types, such as Sound files on Firefox
// which may cause double-load instances.
func (self *Cache) CheckURLI(args ...interface{}) bool{
    return self.Object.Call("checkURL", args).Bool()
}

// CheckCanvasKey Checks if the given key exists in the Canvas Cache.
func (self *Cache) CheckCanvasKey(key string) bool{
    return self.Object.Call("checkCanvasKey", key).Bool()
}

// CheckCanvasKeyI Checks if the given key exists in the Canvas Cache.
func (self *Cache) CheckCanvasKeyI(args ...interface{}) bool{
    return self.Object.Call("checkCanvasKey", args).Bool()
}

// CheckImageKey Checks if the given key exists in the Image Cache. Note that this also includes Texture Atlases, Sprite Sheets and Retro Fonts.
func (self *Cache) CheckImageKey(key string) bool{
    return self.Object.Call("checkImageKey", key).Bool()
}

// CheckImageKeyI Checks if the given key exists in the Image Cache. Note that this also includes Texture Atlases, Sprite Sheets and Retro Fonts.
func (self *Cache) CheckImageKeyI(args ...interface{}) bool{
    return self.Object.Call("checkImageKey", args).Bool()
}

// CheckTextureKey Checks if the given key exists in the Texture Cache.
func (self *Cache) CheckTextureKey(key string) bool{
    return self.Object.Call("checkTextureKey", key).Bool()
}

// CheckTextureKeyI Checks if the given key exists in the Texture Cache.
func (self *Cache) CheckTextureKeyI(args ...interface{}) bool{
    return self.Object.Call("checkTextureKey", args).Bool()
}

// CheckSoundKey Checks if the given key exists in the Sound Cache.
func (self *Cache) CheckSoundKey(key string) bool{
    return self.Object.Call("checkSoundKey", key).Bool()
}

// CheckSoundKeyI Checks if the given key exists in the Sound Cache.
func (self *Cache) CheckSoundKeyI(args ...interface{}) bool{
    return self.Object.Call("checkSoundKey", args).Bool()
}

// CheckTextKey Checks if the given key exists in the Text Cache.
func (self *Cache) CheckTextKey(key string) bool{
    return self.Object.Call("checkTextKey", key).Bool()
}

// CheckTextKeyI Checks if the given key exists in the Text Cache.
func (self *Cache) CheckTextKeyI(args ...interface{}) bool{
    return self.Object.Call("checkTextKey", args).Bool()
}

// CheckPhysicsKey Checks if the given key exists in the Physics Cache.
func (self *Cache) CheckPhysicsKey(key string) bool{
    return self.Object.Call("checkPhysicsKey", key).Bool()
}

// CheckPhysicsKeyI Checks if the given key exists in the Physics Cache.
func (self *Cache) CheckPhysicsKeyI(args ...interface{}) bool{
    return self.Object.Call("checkPhysicsKey", args).Bool()
}

// CheckTilemapKey Checks if the given key exists in the Tilemap Cache.
func (self *Cache) CheckTilemapKey(key string) bool{
    return self.Object.Call("checkTilemapKey", key).Bool()
}

// CheckTilemapKeyI Checks if the given key exists in the Tilemap Cache.
func (self *Cache) CheckTilemapKeyI(args ...interface{}) bool{
    return self.Object.Call("checkTilemapKey", args).Bool()
}

// CheckBinaryKey Checks if the given key exists in the Binary Cache.
func (self *Cache) CheckBinaryKey(key string) bool{
    return self.Object.Call("checkBinaryKey", key).Bool()
}

// CheckBinaryKeyI Checks if the given key exists in the Binary Cache.
func (self *Cache) CheckBinaryKeyI(args ...interface{}) bool{
    return self.Object.Call("checkBinaryKey", args).Bool()
}

// CheckBitmapDataKey Checks if the given key exists in the BitmapData Cache.
func (self *Cache) CheckBitmapDataKey(key string) bool{
    return self.Object.Call("checkBitmapDataKey", key).Bool()
}

// CheckBitmapDataKeyI Checks if the given key exists in the BitmapData Cache.
func (self *Cache) CheckBitmapDataKeyI(args ...interface{}) bool{
    return self.Object.Call("checkBitmapDataKey", args).Bool()
}

// CheckBitmapFontKey Checks if the given key exists in the BitmapFont Cache.
func (self *Cache) CheckBitmapFontKey(key string) bool{
    return self.Object.Call("checkBitmapFontKey", key).Bool()
}

// CheckBitmapFontKeyI Checks if the given key exists in the BitmapFont Cache.
func (self *Cache) CheckBitmapFontKeyI(args ...interface{}) bool{
    return self.Object.Call("checkBitmapFontKey", args).Bool()
}

// CheckJSONKey Checks if the given key exists in the JSON Cache.
func (self *Cache) CheckJSONKey(key string) bool{
    return self.Object.Call("checkJSONKey", key).Bool()
}

// CheckJSONKeyI Checks if the given key exists in the JSON Cache.
func (self *Cache) CheckJSONKeyI(args ...interface{}) bool{
    return self.Object.Call("checkJSONKey", args).Bool()
}

// CheckXMLKey Checks if the given key exists in the XML Cache.
func (self *Cache) CheckXMLKey(key string) bool{
    return self.Object.Call("checkXMLKey", key).Bool()
}

// CheckXMLKeyI Checks if the given key exists in the XML Cache.
func (self *Cache) CheckXMLKeyI(args ...interface{}) bool{
    return self.Object.Call("checkXMLKey", args).Bool()
}

// CheckVideoKey Checks if the given key exists in the Video Cache.
func (self *Cache) CheckVideoKey(key string) bool{
    return self.Object.Call("checkVideoKey", key).Bool()
}

// CheckVideoKeyI Checks if the given key exists in the Video Cache.
func (self *Cache) CheckVideoKeyI(args ...interface{}) bool{
    return self.Object.Call("checkVideoKey", args).Bool()
}

// CheckShaderKey Checks if the given key exists in the Fragment Shader Cache.
func (self *Cache) CheckShaderKey(key string) bool{
    return self.Object.Call("checkShaderKey", key).Bool()
}

// CheckShaderKeyI Checks if the given key exists in the Fragment Shader Cache.
func (self *Cache) CheckShaderKeyI(args ...interface{}) bool{
    return self.Object.Call("checkShaderKey", args).Bool()
}

// CheckRenderTextureKey Checks if the given key exists in the Render Texture Cache.
func (self *Cache) CheckRenderTextureKey(key string) bool{
    return self.Object.Call("checkRenderTextureKey", key).Bool()
}

// CheckRenderTextureKeyI Checks if the given key exists in the Render Texture Cache.
func (self *Cache) CheckRenderTextureKeyI(args ...interface{}) bool{
    return self.Object.Call("checkRenderTextureKey", args).Bool()
}

// GetItem Get an item from a cache based on the given key and property.
// 
// This method is mostly used internally by other Cache methods such as `getImage` but is exposed
// publicly for your own use as well.
func (self *Cache) GetItem(key string, cache int) interface{}{
    return self.Object.Call("getItem", key, cache)
}

// GetItem1O Get an item from a cache based on the given key and property.
// 
// This method is mostly used internally by other Cache methods such as `getImage` but is exposed
// publicly for your own use as well.
func (self *Cache) GetItem1O(key string, cache int, method string) interface{}{
    return self.Object.Call("getItem", key, cache, method)
}

// GetItem2O Get an item from a cache based on the given key and property.
// 
// This method is mostly used internally by other Cache methods such as `getImage` but is exposed
// publicly for your own use as well.
func (self *Cache) GetItem2O(key string, cache int, method string, property string) interface{}{
    return self.Object.Call("getItem", key, cache, method, property)
}

// GetItemI Get an item from a cache based on the given key and property.
// 
// This method is mostly used internally by other Cache methods such as `getImage` but is exposed
// publicly for your own use as well.
func (self *Cache) GetItemI(args ...interface{}) interface{}{
    return self.Object.Call("getItem", args)
}

// GetCanvas Gets a Canvas object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetCanvas(key string) interface{}{
    return self.Object.Call("getCanvas", key)
}

// GetCanvasI Gets a Canvas object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetCanvasI(args ...interface{}) interface{}{
    return self.Object.Call("getCanvas", args)
}

// GetImage Gets a Image object from the cache. This returns a DOM Image object, not a Phaser.Image object.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// Only the Image cache is searched, which covers images loaded via Loader.image, Sprite Sheets and Texture Atlases.
// 
// If you need the image used by a bitmap font or similar then please use those respective 'get' methods.
func (self *Cache) GetImage() *Image{
    return &Image{self.Object.Call("getImage")}
}

// GetImage1O Gets a Image object from the cache. This returns a DOM Image object, not a Phaser.Image object.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// Only the Image cache is searched, which covers images loaded via Loader.image, Sprite Sheets and Texture Atlases.
// 
// If you need the image used by a bitmap font or similar then please use those respective 'get' methods.
func (self *Cache) GetImage1O(key string) *Image{
    return &Image{self.Object.Call("getImage", key)}
}

// GetImage2O Gets a Image object from the cache. This returns a DOM Image object, not a Phaser.Image object.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// Only the Image cache is searched, which covers images loaded via Loader.image, Sprite Sheets and Texture Atlases.
// 
// If you need the image used by a bitmap font or similar then please use those respective 'get' methods.
func (self *Cache) GetImage2O(key string, full bool) *Image{
    return &Image{self.Object.Call("getImage", key, full)}
}

// GetImageI Gets a Image object from the cache. This returns a DOM Image object, not a Phaser.Image object.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// Only the Image cache is searched, which covers images loaded via Loader.image, Sprite Sheets and Texture Atlases.
// 
// If you need the image used by a bitmap font or similar then please use those respective 'get' methods.
func (self *Cache) GetImageI(args ...interface{}) *Image{
    return &Image{self.Object.Call("getImage", args)}
}

// GetTextureFrame Get a single texture frame by key.
// 
// You'd only do this to get the default Frame created for a non-atlas / spritesheet image.
func (self *Cache) GetTextureFrame(key string) *Frame{
    return &Frame{self.Object.Call("getTextureFrame", key)}
}

// GetTextureFrameI Get a single texture frame by key.
// 
// You'd only do this to get the default Frame created for a non-atlas / spritesheet image.
func (self *Cache) GetTextureFrameI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("getTextureFrame", args)}
}

// GetSound Gets a Phaser.Sound object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetSound(key string) *Sound{
    return &Sound{self.Object.Call("getSound", key)}
}

// GetSoundI Gets a Phaser.Sound object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetSoundI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("getSound", args)}
}

// GetSoundData Gets a raw Sound data object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetSoundData(key string) interface{}{
    return self.Object.Call("getSoundData", key)
}

// GetSoundDataI Gets a raw Sound data object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetSoundDataI(args ...interface{}) interface{}{
    return self.Object.Call("getSoundData", args)
}

// GetText Gets a Text object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetText(key string) interface{}{
    return self.Object.Call("getText", key)
}

// GetTextI Gets a Text object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetTextI(args ...interface{}) interface{}{
    return self.Object.Call("getText", args)
}

// GetPhysicsData Gets a Physics Data object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// You can get either the entire data set, a single object or a single fixture of an object from it.
func (self *Cache) GetPhysicsData(key string, object string, fixtureKey string) interface{}{
    return self.Object.Call("getPhysicsData", key, object, fixtureKey)
}

// GetPhysicsDataI Gets a Physics Data object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// You can get either the entire data set, a single object or a single fixture of an object from it.
func (self *Cache) GetPhysicsDataI(args ...interface{}) interface{}{
    return self.Object.Call("getPhysicsData", args)
}

// GetTilemapData Gets a raw Tilemap data object from the cache. This will be in either CSV or JSON format.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetTilemapData(key string) interface{}{
    return self.Object.Call("getTilemapData", key)
}

// GetTilemapDataI Gets a raw Tilemap data object from the cache. This will be in either CSV or JSON format.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetTilemapDataI(args ...interface{}) interface{}{
    return self.Object.Call("getTilemapData", args)
}

// GetBinary Gets a binary object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBinary(key string) interface{}{
    return self.Object.Call("getBinary", key)
}

// GetBinaryI Gets a binary object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBinaryI(args ...interface{}) interface{}{
    return self.Object.Call("getBinary", args)
}

// GetBitmapData Gets a BitmapData object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBitmapData(key string) *BitmapData{
    return &BitmapData{self.Object.Call("getBitmapData", key)}
}

// GetBitmapDataI Gets a BitmapData object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBitmapDataI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("getBitmapData", args)}
}

// GetBitmapFont Gets a Bitmap Font object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBitmapFont(key string) *BitmapFont{
    return &BitmapFont{self.Object.Call("getBitmapFont", key)}
}

// GetBitmapFontI Gets a Bitmap Font object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBitmapFontI(args ...interface{}) *BitmapFont{
    return &BitmapFont{self.Object.Call("getBitmapFont", args)}
}

// GetJSON Gets a JSON object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// You can either return the object by reference (the default), or return a clone
// of it by setting the `clone` argument to `true`.
func (self *Cache) GetJSON(key string) interface{}{
    return self.Object.Call("getJSON", key)
}

// GetJSON1O Gets a JSON object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// You can either return the object by reference (the default), or return a clone
// of it by setting the `clone` argument to `true`.
func (self *Cache) GetJSON1O(key string, clone bool) interface{}{
    return self.Object.Call("getJSON", key, clone)
}

// GetJSONI Gets a JSON object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// You can either return the object by reference (the default), or return a clone
// of it by setting the `clone` argument to `true`.
func (self *Cache) GetJSONI(args ...interface{}) interface{}{
    return self.Object.Call("getJSON", args)
}

// GetXML Gets an XML object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetXML(key string) interface{}{
    return self.Object.Call("getXML", key)
}

// GetXMLI Gets an XML object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetXMLI(args ...interface{}) interface{}{
    return self.Object.Call("getXML", args)
}

// GetVideo Gets a Phaser.Video object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetVideo(key string) *Video{
    return &Video{self.Object.Call("getVideo", key)}
}

// GetVideoI Gets a Phaser.Video object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetVideoI(args ...interface{}) *Video{
    return &Video{self.Object.Call("getVideo", args)}
}

// GetShader Gets a fragment shader object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetShader(key string) string{
    return self.Object.Call("getShader", key).String()
}

// GetShaderI Gets a fragment shader object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetShaderI(args ...interface{}) string{
    return self.Object.Call("getShader", args).String()
}

// GetRenderTexture Gets a RenderTexture object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetRenderTexture(key string) interface{}{
    return self.Object.Call("getRenderTexture", key)
}

// GetRenderTextureI Gets a RenderTexture object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetRenderTextureI(args ...interface{}) interface{}{
    return self.Object.Call("getRenderTexture", args)
}

// GetBaseTexture Gets a PIXI.BaseTexture by key from the given Cache.
func (self *Cache) GetBaseTexture(key string) *BaseTexture{
    return &BaseTexture{self.Object.Call("getBaseTexture", key)}
}

// GetBaseTexture1O Gets a PIXI.BaseTexture by key from the given Cache.
func (self *Cache) GetBaseTexture1O(key string, cache int) *BaseTexture{
    return &BaseTexture{self.Object.Call("getBaseTexture", key, cache)}
}

// GetBaseTextureI Gets a PIXI.BaseTexture by key from the given Cache.
func (self *Cache) GetBaseTextureI(args ...interface{}) *BaseTexture{
    return &BaseTexture{self.Object.Call("getBaseTexture", args)}
}

// GetFrame Get a single frame by key. You'd only do this to get the default Frame created for a non-atlas/spritesheet image.
func (self *Cache) GetFrame(key string) *Frame{
    return &Frame{self.Object.Call("getFrame", key)}
}

// GetFrame1O Get a single frame by key. You'd only do this to get the default Frame created for a non-atlas/spritesheet image.
func (self *Cache) GetFrame1O(key string, cache int) *Frame{
    return &Frame{self.Object.Call("getFrame", key, cache)}
}

// GetFrameI Get a single frame by key. You'd only do this to get the default Frame created for a non-atlas/spritesheet image.
func (self *Cache) GetFrameI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("getFrame", args)}
}

// GetFrameCount Get the total number of frames contained in the FrameData object specified by the given key.
func (self *Cache) GetFrameCount(key string) int{
    return self.Object.Call("getFrameCount", key).Int()
}

// GetFrameCount1O Get the total number of frames contained in the FrameData object specified by the given key.
func (self *Cache) GetFrameCount1O(key string, cache int) int{
    return self.Object.Call("getFrameCount", key, cache).Int()
}

// GetFrameCountI Get the total number of frames contained in the FrameData object specified by the given key.
func (self *Cache) GetFrameCountI(args ...interface{}) int{
    return self.Object.Call("getFrameCount", args).Int()
}

// GetFrameData Gets a Phaser.FrameData object from the Image Cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetFrameData(key string) *FrameData{
    return &FrameData{self.Object.Call("getFrameData", key)}
}

// GetFrameData1O Gets a Phaser.FrameData object from the Image Cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetFrameData1O(key string, cache int) *FrameData{
    return &FrameData{self.Object.Call("getFrameData", key, cache)}
}

// GetFrameDataI Gets a Phaser.FrameData object from the Image Cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetFrameDataI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("getFrameData", args)}
}

// HasFrameData Check if the FrameData for the given key exists in the Image Cache.
func (self *Cache) HasFrameData(key string) bool{
    return self.Object.Call("hasFrameData", key).Bool()
}

// HasFrameData1O Check if the FrameData for the given key exists in the Image Cache.
func (self *Cache) HasFrameData1O(key string, cache int) bool{
    return self.Object.Call("hasFrameData", key, cache).Bool()
}

// HasFrameDataI Check if the FrameData for the given key exists in the Image Cache.
func (self *Cache) HasFrameDataI(args ...interface{}) bool{
    return self.Object.Call("hasFrameData", args).Bool()
}

// UpdateFrameData Replaces a set of frameData with a new Phaser.FrameData object.
func (self *Cache) UpdateFrameData(key string, frameData int) {
    self.Object.Call("updateFrameData", key, frameData)
}

// UpdateFrameData1O Replaces a set of frameData with a new Phaser.FrameData object.
func (self *Cache) UpdateFrameData1O(key string, frameData int, cache int) {
    self.Object.Call("updateFrameData", key, frameData, cache)
}

// UpdateFrameDataI Replaces a set of frameData with a new Phaser.FrameData object.
func (self *Cache) UpdateFrameDataI(args ...interface{}) {
    self.Object.Call("updateFrameData", args)
}

// GetFrameByIndex Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByIndex(key string, index int) *Frame{
    return &Frame{self.Object.Call("getFrameByIndex", key, index)}
}

// GetFrameByIndex1O Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByIndex1O(key string, index int, cache int) *Frame{
    return &Frame{self.Object.Call("getFrameByIndex", key, index, cache)}
}

// GetFrameByIndexI Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByIndexI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("getFrameByIndex", args)}
}

// GetFrameByName Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByName(key string, name string) *Frame{
    return &Frame{self.Object.Call("getFrameByName", key, name)}
}

// GetFrameByName1O Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByName1O(key string, name string, cache int) *Frame{
    return &Frame{self.Object.Call("getFrameByName", key, name, cache)}
}

// GetFrameByNameI Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByNameI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("getFrameByName", args)}
}

// GetURL Get a cached object by the URL.
// This only returns a value if you set Cache.autoResolveURL to `true` *before* starting the preload of any assets.
// Be aware that every call to this function makes a DOM src query, so use carefully and double-check for implications in your target browsers/devices.
func (self *Cache) GetURL(url string) interface{}{
    return self.Object.Call("getURL", url)
}

// GetURLI Get a cached object by the URL.
// This only returns a value if you set Cache.autoResolveURL to `true` *before* starting the preload of any assets.
// Be aware that every call to this function makes a DOM src query, so use carefully and double-check for implications in your target browsers/devices.
func (self *Cache) GetURLI(args ...interface{}) interface{}{
    return self.Object.Call("getURL", args)
}

// GetKeys Gets all keys used in the requested Cache.
func (self *Cache) GetKeys() []interface{}{
	array00 := self.Object.Call("getKeys")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetKeys1O Gets all keys used in the requested Cache.
func (self *Cache) GetKeys1O(cache int) []interface{}{
	array00 := self.Object.Call("getKeys", cache)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetKeysI Gets all keys used in the requested Cache.
func (self *Cache) GetKeysI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("getKeys", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// RemoveCanvas Removes a canvas from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveCanvas(key string) {
    self.Object.Call("removeCanvas", key)
}

// RemoveCanvasI Removes a canvas from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveCanvasI(args ...interface{}) {
    self.Object.Call("removeCanvas", args)
}

// RemoveImage Removes an image from the cache.
// 
// You can optionally elect to destroy it as well. This calls BaseTexture.destroy on it.
// 
// Note that this only removes it from the Phaser Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveImage(key string) {
    self.Object.Call("removeImage", key)
}

// RemoveImage1O Removes an image from the cache.
// 
// You can optionally elect to destroy it as well. This calls BaseTexture.destroy on it.
// 
// Note that this only removes it from the Phaser Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveImage1O(key string, destroyBaseTexture bool) {
    self.Object.Call("removeImage", key, destroyBaseTexture)
}

// RemoveImageI Removes an image from the cache.
// 
// You can optionally elect to destroy it as well. This calls BaseTexture.destroy on it.
// 
// Note that this only removes it from the Phaser Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveImageI(args ...interface{}) {
    self.Object.Call("removeImage", args)
}

// RemoveSound Removes a sound from the cache.
// 
// If any `Phaser.Sound` objects use the audio file in the cache that you remove with this method, they will
// _automatically_ destroy themselves. If you wish to have full control over when Sounds are destroyed then
// you must finish your house-keeping and destroy them all yourself first, before calling this method.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveSound(key string) {
    self.Object.Call("removeSound", key)
}

// RemoveSoundI Removes a sound from the cache.
// 
// If any `Phaser.Sound` objects use the audio file in the cache that you remove with this method, they will
// _automatically_ destroy themselves. If you wish to have full control over when Sounds are destroyed then
// you must finish your house-keeping and destroy them all yourself first, before calling this method.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveSoundI(args ...interface{}) {
    self.Object.Call("removeSound", args)
}

// RemoveText Removes a text file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveText(key string) {
    self.Object.Call("removeText", key)
}

// RemoveTextI Removes a text file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTextI(args ...interface{}) {
    self.Object.Call("removeText", args)
}

// RemovePhysics Removes a physics data file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemovePhysics(key string) {
    self.Object.Call("removePhysics", key)
}

// RemovePhysicsI Removes a physics data file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemovePhysicsI(args ...interface{}) {
    self.Object.Call("removePhysics", args)
}

// RemoveTilemap Removes a tilemap from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTilemap(key string) {
    self.Object.Call("removeTilemap", key)
}

// RemoveTilemapI Removes a tilemap from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTilemapI(args ...interface{}) {
    self.Object.Call("removeTilemap", args)
}

// RemoveBinary Removes a binary file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBinary(key string) {
    self.Object.Call("removeBinary", key)
}

// RemoveBinaryI Removes a binary file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBinaryI(args ...interface{}) {
    self.Object.Call("removeBinary", args)
}

// RemoveBitmapData Removes a bitmap data from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBitmapData(key string) {
    self.Object.Call("removeBitmapData", key)
}

// RemoveBitmapDataI Removes a bitmap data from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBitmapDataI(args ...interface{}) {
    self.Object.Call("removeBitmapData", args)
}

// RemoveBitmapFont Removes a bitmap font from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBitmapFont(key string) {
    self.Object.Call("removeBitmapFont", key)
}

// RemoveBitmapFontI Removes a bitmap font from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBitmapFontI(args ...interface{}) {
    self.Object.Call("removeBitmapFont", args)
}

// RemoveJSON Removes a json object from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveJSON(key string) {
    self.Object.Call("removeJSON", key)
}

// RemoveJSONI Removes a json object from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveJSONI(args ...interface{}) {
    self.Object.Call("removeJSON", args)
}

// RemoveXML Removes a xml object from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveXML(key string) {
    self.Object.Call("removeXML", key)
}

// RemoveXMLI Removes a xml object from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveXMLI(args ...interface{}) {
    self.Object.Call("removeXML", args)
}

// RemoveVideo Removes a video from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveVideo(key string) {
    self.Object.Call("removeVideo", key)
}

// RemoveVideoI Removes a video from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveVideoI(args ...interface{}) {
    self.Object.Call("removeVideo", args)
}

// RemoveShader Removes a shader from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveShader(key string) {
    self.Object.Call("removeShader", key)
}

// RemoveShaderI Removes a shader from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveShaderI(args ...interface{}) {
    self.Object.Call("removeShader", args)
}

// RemoveRenderTexture Removes a Render Texture from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveRenderTexture(key string) {
    self.Object.Call("removeRenderTexture", key)
}

// RemoveRenderTextureI Removes a Render Texture from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveRenderTextureI(args ...interface{}) {
    self.Object.Call("removeRenderTexture", args)
}

// RemoveSpriteSheet Removes a Sprite Sheet from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveSpriteSheet(key string) {
    self.Object.Call("removeSpriteSheet", key)
}

// RemoveSpriteSheetI Removes a Sprite Sheet from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveSpriteSheetI(args ...interface{}) {
    self.Object.Call("removeSpriteSheet", args)
}

// RemoveTextureAtlas Removes a Texture Atlas from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTextureAtlas(key string) {
    self.Object.Call("removeTextureAtlas", key)
}

// RemoveTextureAtlasI Removes a Texture Atlas from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTextureAtlasI(args ...interface{}) {
    self.Object.Call("removeTextureAtlas", args)
}

// ClearGLTextures Empties out all of the GL Textures from Images stored in the cache.
// This is called automatically when the WebGL context is lost and then restored.
func (self *Cache) ClearGLTextures() {
    self.Object.Call("clearGLTextures")
}

// ClearGLTexturesI Empties out all of the GL Textures from Images stored in the cache.
// This is called automatically when the WebGL context is lost and then restored.
func (self *Cache) ClearGLTexturesI(args ...interface{}) {
    self.Object.Call("clearGLTextures", args)
}

// _resolveURL Resolves a URL to its absolute form and stores it in Cache._urlMap as long as Cache.autoResolveURL is set to `true`.
// This is then looked-up by the Cache.getURL and Cache.checkURL calls.
func (self *Cache) _resolveURL(url string) string{
    return self.Object.Call("_resolveURL", url).String()
}

// _resolveURL1O Resolves a URL to its absolute form and stores it in Cache._urlMap as long as Cache.autoResolveURL is set to `true`.
// This is then looked-up by the Cache.getURL and Cache.checkURL calls.
func (self *Cache) _resolveURL1O(url string, data interface{}) string{
    return self.Object.Call("_resolveURL", url, data).String()
}

// _resolveURLI Resolves a URL to its absolute form and stores it in Cache._urlMap as long as Cache.autoResolveURL is set to `true`.
// This is then looked-up by the Cache.getURL and Cache.checkURL calls.
func (self *Cache) _resolveURLI(args ...interface{}) string{
    return self.Object.Call("_resolveURL", args).String()
}

// Destroy Clears the cache. Removes every local cache object reference.
// If an object in the cache has a `destroy` method it will also be called.
func (self *Cache) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Clears the cache. Removes every local cache object reference.
// If an object in the cache has a `destroy` method it will also be called.
func (self *Cache) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

