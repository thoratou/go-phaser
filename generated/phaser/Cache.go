// Automatic generation for Phaser.Cache
// generated file Cache.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// Phaser has one single cache in which it stores all assets.
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


// Phaser has one single cache in which it stores all assets.
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

// Phaser has one single cache in which it stores all assets.
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



// Local reference to game.
func (self *Cache) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *Cache) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Automatically resolve resource URLs to absolute paths for use with the Cache.getURL method.
func (self *Cache) AutoResolveURL() bool{
    return self.Object.Get("autoResolveURL").Bool()
}

// Automatically resolve resource URLs to absolute paths for use with the Cache.getURL method.
func (self *Cache) SetAutoResolveURLA(member bool) {
    self.Object.Set("autoResolveURL", member)
}

// This event is dispatched when the sound system is unlocked via a touch event on cellular devices.
func (self *Cache) OnSoundUnlock() *Signal{
    return &Signal{self.Object.Get("onSoundUnlock")}
}

// This event is dispatched when the sound system is unlocked via a touch event on cellular devices.
func (self *Cache) SetOnSoundUnlockA(member *Signal) {
    self.Object.Set("onSoundUnlock", member)
}

// 
func (self *Cache) CANVAS() int{
    return self.Object.Get("CANVAS").Int()
}

// 
func (self *Cache) SetCANVASA(member int) {
    self.Object.Set("CANVAS", member)
}

// 
func (self *Cache) IMAGE() int{
    return self.Object.Get("IMAGE").Int()
}

// 
func (self *Cache) SetIMAGEA(member int) {
    self.Object.Set("IMAGE", member)
}

// 
func (self *Cache) TEXTURE() int{
    return self.Object.Get("TEXTURE").Int()
}

// 
func (self *Cache) SetTEXTUREA(member int) {
    self.Object.Set("TEXTURE", member)
}

// 
func (self *Cache) SOUND() int{
    return self.Object.Get("SOUND").Int()
}

// 
func (self *Cache) SetSOUNDA(member int) {
    self.Object.Set("SOUND", member)
}

// 
func (self *Cache) TEXT() int{
    return self.Object.Get("TEXT").Int()
}

// 
func (self *Cache) SetTEXTA(member int) {
    self.Object.Set("TEXT", member)
}

// 
func (self *Cache) PHYSICS() int{
    return self.Object.Get("PHYSICS").Int()
}

// 
func (self *Cache) SetPHYSICSA(member int) {
    self.Object.Set("PHYSICS", member)
}

// 
func (self *Cache) TILEMAP() int{
    return self.Object.Get("TILEMAP").Int()
}

// 
func (self *Cache) SetTILEMAPA(member int) {
    self.Object.Set("TILEMAP", member)
}

// 
func (self *Cache) BINARY() int{
    return self.Object.Get("BINARY").Int()
}

// 
func (self *Cache) SetBINARYA(member int) {
    self.Object.Set("BINARY", member)
}

// 
func (self *Cache) BITMAPDATA() int{
    return self.Object.Get("BITMAPDATA").Int()
}

// 
func (self *Cache) SetBITMAPDATAA(member int) {
    self.Object.Set("BITMAPDATA", member)
}

// 
func (self *Cache) BITMAPFONT() int{
    return self.Object.Get("BITMAPFONT").Int()
}

// 
func (self *Cache) SetBITMAPFONTA(member int) {
    self.Object.Set("BITMAPFONT", member)
}

// 
func (self *Cache) JSON() int{
    return self.Object.Get("JSON").Int()
}

// 
func (self *Cache) SetJSONA(member int) {
    self.Object.Set("JSON", member)
}

// 
func (self *Cache) XML() int{
    return self.Object.Get("XML").Int()
}

// 
func (self *Cache) SetXMLA(member int) {
    self.Object.Set("XML", member)
}

// 
func (self *Cache) VIDEO() int{
    return self.Object.Get("VIDEO").Int()
}

// 
func (self *Cache) SetVIDEOA(member int) {
    self.Object.Set("VIDEO", member)
}

// 
func (self *Cache) SHADER() int{
    return self.Object.Get("SHADER").Int()
}

// 
func (self *Cache) SetSHADERA(member int) {
    self.Object.Set("SHADER", member)
}

// 
func (self *Cache) RENDER_TEXTURE() int{
    return self.Object.Get("RENDER_TEXTURE").Int()
}

// 
func (self *Cache) SetRENDER_TEXTUREA(member int) {
    self.Object.Set("RENDER_TEXTURE", member)
}

// The default image used for a texture when no other is specified.
func (self *Cache) DEFAULT() *Texture{
    return &Texture{self.Object.Get("DEFAULT")}
}

// The default image used for a texture when no other is specified.
func (self *Cache) SetDEFAULTA(member *Texture) {
    self.Object.Set("DEFAULT", member)
}

// The default image used for a texture when the source image is missing.
func (self *Cache) MISSING() *Texture{
    return &Texture{self.Object.Get("MISSING")}
}

// The default image used for a texture when the source image is missing.
func (self *Cache) SetMISSINGA(member *Texture) {
    self.Object.Set("MISSING", member)
}



// Add a new canvas object in to the cache.
func (self *Cache) AddCanvas(key string, canvas *dom.HTMLCanvasElement) {
    self.Object.Call("addCanvas", key, canvas)
}

// Add a new canvas object in to the cache.
func (self *Cache) AddCanvas1O(key string, canvas *dom.HTMLCanvasElement, context *dom.CanvasRenderingContext2D) {
    self.Object.Call("addCanvas", key, canvas, context)
}

// Add a new canvas object in to the cache.
func (self *Cache) AddCanvasI(args ...interface{}) {
    self.Object.Call("addCanvas", args)
}

// Adds an Image file into the Cache. The file must have already been loaded, typically via Phaser.Loader, but can also have been loaded into the DOM.
// If an image already exists in the cache with the same key then it is removed and destroyed, and the new image inserted in its place.
func (self *Cache) AddImage(key string, url string, data interface{}) interface{}{
    return self.Object.Call("addImage", key, url, data)
}

// Adds an Image file into the Cache. The file must have already been loaded, typically via Phaser.Loader, but can also have been loaded into the DOM.
// If an image already exists in the cache with the same key then it is removed and destroyed, and the new image inserted in its place.
func (self *Cache) AddImageI(args ...interface{}) interface{}{
    return self.Object.Call("addImage", args)
}

// Adds a default image to be used in special cases such as WebGL Filters.
// It uses the special reserved key of `__default`.
// This method is called automatically when the Cache is created.
// This image is skipped when `Cache.destroy` is called due to its internal requirements.
func (self *Cache) AddDefaultImage() {
    self.Object.Call("addDefaultImage")
}

// Adds a default image to be used in special cases such as WebGL Filters.
// It uses the special reserved key of `__default`.
// This method is called automatically when the Cache is created.
// This image is skipped when `Cache.destroy` is called due to its internal requirements.
func (self *Cache) AddDefaultImageI(args ...interface{}) {
    self.Object.Call("addDefaultImage", args)
}

// Adds an image to be used when a key is wrong / missing.
// It uses the special reserved key of `__missing`.
// This method is called automatically when the Cache is created.
// This image is skipped when `Cache.destroy` is called due to its internal requirements.
func (self *Cache) AddMissingImage() {
    self.Object.Call("addMissingImage")
}

// Adds an image to be used when a key is wrong / missing.
// It uses the special reserved key of `__missing`.
// This method is called automatically when the Cache is created.
// This image is skipped when `Cache.destroy` is called due to its internal requirements.
func (self *Cache) AddMissingImageI(args ...interface{}) {
    self.Object.Call("addMissingImage", args)
}

// Adds a Sound file into the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddSound(key string, url string, data interface{}, webAudio bool, audioTag bool) {
    self.Object.Call("addSound", key, url, data, webAudio, audioTag)
}

// Adds a Sound file into the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddSoundI(args ...interface{}) {
    self.Object.Call("addSound", args)
}

// Add a new text data.
func (self *Cache) AddText(key string, url string, data interface{}) {
    self.Object.Call("addText", key, url, data)
}

// Add a new text data.
func (self *Cache) AddTextI(args ...interface{}) {
    self.Object.Call("addText", args)
}

// Add a new physics data object to the Cache.
func (self *Cache) AddPhysicsData(key string, url string, JSONData interface{}, format int) {
    self.Object.Call("addPhysicsData", key, url, JSONData, format)
}

// Add a new physics data object to the Cache.
func (self *Cache) AddPhysicsDataI(args ...interface{}) {
    self.Object.Call("addPhysicsData", args)
}

// Add a new tilemap to the Cache.
func (self *Cache) AddTilemap(key string, url string, mapData interface{}, format int) {
    self.Object.Call("addTilemap", key, url, mapData, format)
}

// Add a new tilemap to the Cache.
func (self *Cache) AddTilemapI(args ...interface{}) {
    self.Object.Call("addTilemap", args)
}

// Add a binary object in to the cache.
func (self *Cache) AddBinary(key string, binaryData interface{}) {
    self.Object.Call("addBinary", key, binaryData)
}

// Add a binary object in to the cache.
func (self *Cache) AddBinaryI(args ...interface{}) {
    self.Object.Call("addBinary", args)
}

// Add a BitmapData object to the cache.
func (self *Cache) AddBitmapData(key string, bitmapData *BitmapData) *BitmapData{
    return &BitmapData{self.Object.Call("addBitmapData", key, bitmapData)}
}

// Add a BitmapData object to the cache.
func (self *Cache) AddBitmapData1O(key string, bitmapData *BitmapData, frameData interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("addBitmapData", key, bitmapData, frameData)}
}

// Add a BitmapData object to the cache.
func (self *Cache) AddBitmapDataI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("addBitmapData", args)}
}

// Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFont(key string, url string, data interface{}, atlasData interface{}) {
    self.Object.Call("addBitmapFont", key, url, data, atlasData)
}

// Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFont1O(key string, url string, data interface{}, atlasData interface{}, atlasType string) {
    self.Object.Call("addBitmapFont", key, url, data, atlasData, atlasType)
}

// Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFont2O(key string, url string, data interface{}, atlasData interface{}, atlasType string, xSpacing int) {
    self.Object.Call("addBitmapFont", key, url, data, atlasData, atlasType, xSpacing)
}

// Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFont3O(key string, url string, data interface{}, atlasData interface{}, atlasType string, xSpacing int, ySpacing int) {
    self.Object.Call("addBitmapFont", key, url, data, atlasData, atlasType, xSpacing, ySpacing)
}

// Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFontI(args ...interface{}) {
    self.Object.Call("addBitmapFont", args)
}

// Add a new json object into the cache.
func (self *Cache) AddJSON(key string, url string, data interface{}) {
    self.Object.Call("addJSON", key, url, data)
}

// Add a new json object into the cache.
func (self *Cache) AddJSONI(args ...interface{}) {
    self.Object.Call("addJSON", args)
}

// Add a new xml object into the cache.
func (self *Cache) AddXML(key string, url string, data interface{}) {
    self.Object.Call("addXML", key, url, data)
}

// Add a new xml object into the cache.
func (self *Cache) AddXMLI(args ...interface{}) {
    self.Object.Call("addXML", args)
}

// Adds a Video file into the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddVideo(key string, url string, data interface{}, isBlob bool) {
    self.Object.Call("addVideo", key, url, data, isBlob)
}

// Adds a Video file into the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddVideoI(args ...interface{}) {
    self.Object.Call("addVideo", args)
}

// Adds a Fragment Shader in to the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddShader(key string, url string, data interface{}) {
    self.Object.Call("addShader", key, url, data)
}

// Adds a Fragment Shader in to the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddShaderI(args ...interface{}) {
    self.Object.Call("addShader", args)
}

// Add a new Phaser.RenderTexture in to the cache.
func (self *Cache) AddRenderTexture(key string, texture *RenderTexture) {
    self.Object.Call("addRenderTexture", key, texture)
}

// Add a new Phaser.RenderTexture in to the cache.
func (self *Cache) AddRenderTextureI(args ...interface{}) {
    self.Object.Call("addRenderTexture", args)
}

// Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheet(key string, url string, data interface{}, frameWidth int, frameHeight int) {
    self.Object.Call("addSpriteSheet", key, url, data, frameWidth, frameHeight)
}

// Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheet1O(key string, url string, data interface{}, frameWidth int, frameHeight int, frameMax int) {
    self.Object.Call("addSpriteSheet", key, url, data, frameWidth, frameHeight, frameMax)
}

// Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheet2O(key string, url string, data interface{}, frameWidth int, frameHeight int, frameMax int, margin int) {
    self.Object.Call("addSpriteSheet", key, url, data, frameWidth, frameHeight, frameMax, margin)
}

// Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheet3O(key string, url string, data interface{}, frameWidth int, frameHeight int, frameMax int, margin int, spacing int) {
    self.Object.Call("addSpriteSheet", key, url, data, frameWidth, frameHeight, frameMax, margin, spacing)
}

// Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheetI(args ...interface{}) {
    self.Object.Call("addSpriteSheet", args)
}

// Add a new texture atlas to the Cache.
func (self *Cache) AddTextureAtlas(key string, url string, data interface{}, atlasData interface{}, format int) {
    self.Object.Call("addTextureAtlas", key, url, data, atlasData, format)
}

// Add a new texture atlas to the Cache.
func (self *Cache) AddTextureAtlasI(args ...interface{}) {
    self.Object.Call("addTextureAtlas", args)
}

// Reload a Sound file from the server.
func (self *Cache) ReloadSound(key string) {
    self.Object.Call("reloadSound", key)
}

// Reload a Sound file from the server.
func (self *Cache) ReloadSoundI(args ...interface{}) {
    self.Object.Call("reloadSound", args)
}

// Fires the onSoundUnlock event when the sound has completed reloading.
func (self *Cache) ReloadSoundComplete(key string) {
    self.Object.Call("reloadSoundComplete", key)
}

// Fires the onSoundUnlock event when the sound has completed reloading.
func (self *Cache) ReloadSoundCompleteI(args ...interface{}) {
    self.Object.Call("reloadSoundComplete", args)
}

// Updates the sound object in the cache.
func (self *Cache) UpdateSound(key string) {
    self.Object.Call("updateSound", key)
}

// Updates the sound object in the cache.
func (self *Cache) UpdateSoundI(args ...interface{}) {
    self.Object.Call("updateSound", args)
}

// Add a new decoded sound.
func (self *Cache) DecodedSound(key string, data interface{}) {
    self.Object.Call("decodedSound", key, data)
}

// Add a new decoded sound.
func (self *Cache) DecodedSoundI(args ...interface{}) {
    self.Object.Call("decodedSound", args)
}

// Check if the given sound has finished decoding.
func (self *Cache) IsSoundDecoded(key string) bool{
    return self.Object.Call("isSoundDecoded", key).Bool()
}

// Check if the given sound has finished decoding.
func (self *Cache) IsSoundDecodedI(args ...interface{}) bool{
    return self.Object.Call("isSoundDecoded", args).Bool()
}

// Check if the given sound is ready for playback.
// A sound is considered ready when it has finished decoding and the device is no longer touch locked.
func (self *Cache) IsSoundReady(key string) bool{
    return self.Object.Call("isSoundReady", key).Bool()
}

// Check if the given sound is ready for playback.
// A sound is considered ready when it has finished decoding and the device is no longer touch locked.
func (self *Cache) IsSoundReadyI(args ...interface{}) bool{
    return self.Object.Call("isSoundReady", args).Bool()
}

// Checks if a key for the given cache object type exists.
func (self *Cache) CheckKey(cache int, key string) bool{
    return self.Object.Call("checkKey", cache, key).Bool()
}

// Checks if a key for the given cache object type exists.
func (self *Cache) CheckKeyI(args ...interface{}) bool{
    return self.Object.Call("checkKey", args).Bool()
}

// Checks if the given URL has been loaded into the Cache.
// This method will only work if Cache.autoResolveURL was set to `true` before any preloading took place.
// The method will make a DOM src call to the URL given, so please be aware of this for certain file types, such as Sound files on Firefox
// which may cause double-load instances.
func (self *Cache) CheckURL(url string) bool{
    return self.Object.Call("checkURL", url).Bool()
}

// Checks if the given URL has been loaded into the Cache.
// This method will only work if Cache.autoResolveURL was set to `true` before any preloading took place.
// The method will make a DOM src call to the URL given, so please be aware of this for certain file types, such as Sound files on Firefox
// which may cause double-load instances.
func (self *Cache) CheckURLI(args ...interface{}) bool{
    return self.Object.Call("checkURL", args).Bool()
}

// Checks if the given key exists in the Canvas Cache.
func (self *Cache) CheckCanvasKey(key string) bool{
    return self.Object.Call("checkCanvasKey", key).Bool()
}

// Checks if the given key exists in the Canvas Cache.
func (self *Cache) CheckCanvasKeyI(args ...interface{}) bool{
    return self.Object.Call("checkCanvasKey", args).Bool()
}

// Checks if the given key exists in the Image Cache. Note that this also includes Texture Atlases, Sprite Sheets and Retro Fonts.
func (self *Cache) CheckImageKey(key string) bool{
    return self.Object.Call("checkImageKey", key).Bool()
}

// Checks if the given key exists in the Image Cache. Note that this also includes Texture Atlases, Sprite Sheets and Retro Fonts.
func (self *Cache) CheckImageKeyI(args ...interface{}) bool{
    return self.Object.Call("checkImageKey", args).Bool()
}

// Checks if the given key exists in the Texture Cache.
func (self *Cache) CheckTextureKey(key string) bool{
    return self.Object.Call("checkTextureKey", key).Bool()
}

// Checks if the given key exists in the Texture Cache.
func (self *Cache) CheckTextureKeyI(args ...interface{}) bool{
    return self.Object.Call("checkTextureKey", args).Bool()
}

// Checks if the given key exists in the Sound Cache.
func (self *Cache) CheckSoundKey(key string) bool{
    return self.Object.Call("checkSoundKey", key).Bool()
}

// Checks if the given key exists in the Sound Cache.
func (self *Cache) CheckSoundKeyI(args ...interface{}) bool{
    return self.Object.Call("checkSoundKey", args).Bool()
}

// Checks if the given key exists in the Text Cache.
func (self *Cache) CheckTextKey(key string) bool{
    return self.Object.Call("checkTextKey", key).Bool()
}

// Checks if the given key exists in the Text Cache.
func (self *Cache) CheckTextKeyI(args ...interface{}) bool{
    return self.Object.Call("checkTextKey", args).Bool()
}

// Checks if the given key exists in the Physics Cache.
func (self *Cache) CheckPhysicsKey(key string) bool{
    return self.Object.Call("checkPhysicsKey", key).Bool()
}

// Checks if the given key exists in the Physics Cache.
func (self *Cache) CheckPhysicsKeyI(args ...interface{}) bool{
    return self.Object.Call("checkPhysicsKey", args).Bool()
}

// Checks if the given key exists in the Tilemap Cache.
func (self *Cache) CheckTilemapKey(key string) bool{
    return self.Object.Call("checkTilemapKey", key).Bool()
}

// Checks if the given key exists in the Tilemap Cache.
func (self *Cache) CheckTilemapKeyI(args ...interface{}) bool{
    return self.Object.Call("checkTilemapKey", args).Bool()
}

// Checks if the given key exists in the Binary Cache.
func (self *Cache) CheckBinaryKey(key string) bool{
    return self.Object.Call("checkBinaryKey", key).Bool()
}

// Checks if the given key exists in the Binary Cache.
func (self *Cache) CheckBinaryKeyI(args ...interface{}) bool{
    return self.Object.Call("checkBinaryKey", args).Bool()
}

// Checks if the given key exists in the BitmapData Cache.
func (self *Cache) CheckBitmapDataKey(key string) bool{
    return self.Object.Call("checkBitmapDataKey", key).Bool()
}

// Checks if the given key exists in the BitmapData Cache.
func (self *Cache) CheckBitmapDataKeyI(args ...interface{}) bool{
    return self.Object.Call("checkBitmapDataKey", args).Bool()
}

// Checks if the given key exists in the BitmapFont Cache.
func (self *Cache) CheckBitmapFontKey(key string) bool{
    return self.Object.Call("checkBitmapFontKey", key).Bool()
}

// Checks if the given key exists in the BitmapFont Cache.
func (self *Cache) CheckBitmapFontKeyI(args ...interface{}) bool{
    return self.Object.Call("checkBitmapFontKey", args).Bool()
}

// Checks if the given key exists in the JSON Cache.
func (self *Cache) CheckJSONKey(key string) bool{
    return self.Object.Call("checkJSONKey", key).Bool()
}

// Checks if the given key exists in the JSON Cache.
func (self *Cache) CheckJSONKeyI(args ...interface{}) bool{
    return self.Object.Call("checkJSONKey", args).Bool()
}

// Checks if the given key exists in the XML Cache.
func (self *Cache) CheckXMLKey(key string) bool{
    return self.Object.Call("checkXMLKey", key).Bool()
}

// Checks if the given key exists in the XML Cache.
func (self *Cache) CheckXMLKeyI(args ...interface{}) bool{
    return self.Object.Call("checkXMLKey", args).Bool()
}

// Checks if the given key exists in the Video Cache.
func (self *Cache) CheckVideoKey(key string) bool{
    return self.Object.Call("checkVideoKey", key).Bool()
}

// Checks if the given key exists in the Video Cache.
func (self *Cache) CheckVideoKeyI(args ...interface{}) bool{
    return self.Object.Call("checkVideoKey", args).Bool()
}

// Checks if the given key exists in the Fragment Shader Cache.
func (self *Cache) CheckShaderKey(key string) bool{
    return self.Object.Call("checkShaderKey", key).Bool()
}

// Checks if the given key exists in the Fragment Shader Cache.
func (self *Cache) CheckShaderKeyI(args ...interface{}) bool{
    return self.Object.Call("checkShaderKey", args).Bool()
}

// Checks if the given key exists in the Render Texture Cache.
func (self *Cache) CheckRenderTextureKey(key string) bool{
    return self.Object.Call("checkRenderTextureKey", key).Bool()
}

// Checks if the given key exists in the Render Texture Cache.
func (self *Cache) CheckRenderTextureKeyI(args ...interface{}) bool{
    return self.Object.Call("checkRenderTextureKey", args).Bool()
}

// Get an item from a cache based on the given key and property.
// 
// This method is mostly used internally by other Cache methods such as `getImage` but is exposed
// publicly for your own use as well.
func (self *Cache) GetItem(key string, cache int) interface{}{
    return self.Object.Call("getItem", key, cache)
}

// Get an item from a cache based on the given key and property.
// 
// This method is mostly used internally by other Cache methods such as `getImage` but is exposed
// publicly for your own use as well.
func (self *Cache) GetItem1O(key string, cache int, method string) interface{}{
    return self.Object.Call("getItem", key, cache, method)
}

// Get an item from a cache based on the given key and property.
// 
// This method is mostly used internally by other Cache methods such as `getImage` but is exposed
// publicly for your own use as well.
func (self *Cache) GetItem2O(key string, cache int, method string, property string) interface{}{
    return self.Object.Call("getItem", key, cache, method, property)
}

// Get an item from a cache based on the given key and property.
// 
// This method is mostly used internally by other Cache methods such as `getImage` but is exposed
// publicly for your own use as well.
func (self *Cache) GetItemI(args ...interface{}) interface{}{
    return self.Object.Call("getItem", args)
}

// Gets a Canvas object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetCanvas(key string) interface{}{
    return self.Object.Call("getCanvas", key)
}

// Gets a Canvas object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetCanvasI(args ...interface{}) interface{}{
    return self.Object.Call("getCanvas", args)
}

// Gets a Image object from the cache. This returns a DOM Image object, not a Phaser.Image object.
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

// Gets a Image object from the cache. This returns a DOM Image object, not a Phaser.Image object.
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

// Gets a Image object from the cache. This returns a DOM Image object, not a Phaser.Image object.
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

// Gets a Image object from the cache. This returns a DOM Image object, not a Phaser.Image object.
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

// Get a single texture frame by key.
// 
// You'd only do this to get the default Frame created for a non-atlas / spritesheet image.
func (self *Cache) GetTextureFrame(key string) *Frame{
    return &Frame{self.Object.Call("getTextureFrame", key)}
}

// Get a single texture frame by key.
// 
// You'd only do this to get the default Frame created for a non-atlas / spritesheet image.
func (self *Cache) GetTextureFrameI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("getTextureFrame", args)}
}

// Gets a Phaser.Sound object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetSound(key string) *Sound{
    return &Sound{self.Object.Call("getSound", key)}
}

// Gets a Phaser.Sound object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetSoundI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("getSound", args)}
}

// Gets a raw Sound data object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetSoundData(key string) interface{}{
    return self.Object.Call("getSoundData", key)
}

// Gets a raw Sound data object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetSoundDataI(args ...interface{}) interface{}{
    return self.Object.Call("getSoundData", args)
}

// Gets a Text object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetText(key string) interface{}{
    return self.Object.Call("getText", key)
}

// Gets a Text object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetTextI(args ...interface{}) interface{}{
    return self.Object.Call("getText", args)
}

// Gets a Physics Data object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// You can get either the entire data set, a single object or a single fixture of an object from it.
func (self *Cache) GetPhysicsData(key string, object string, fixtureKey string) interface{}{
    return self.Object.Call("getPhysicsData", key, object, fixtureKey)
}

// Gets a Physics Data object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// You can get either the entire data set, a single object or a single fixture of an object from it.
func (self *Cache) GetPhysicsDataI(args ...interface{}) interface{}{
    return self.Object.Call("getPhysicsData", args)
}

// Gets a raw Tilemap data object from the cache. This will be in either CSV or JSON format.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetTilemapData(key string) interface{}{
    return self.Object.Call("getTilemapData", key)
}

// Gets a raw Tilemap data object from the cache. This will be in either CSV or JSON format.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetTilemapDataI(args ...interface{}) interface{}{
    return self.Object.Call("getTilemapData", args)
}

// Gets a binary object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBinary(key string) interface{}{
    return self.Object.Call("getBinary", key)
}

// Gets a binary object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBinaryI(args ...interface{}) interface{}{
    return self.Object.Call("getBinary", args)
}

// Gets a BitmapData object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBitmapData(key string) *BitmapData{
    return &BitmapData{self.Object.Call("getBitmapData", key)}
}

// Gets a BitmapData object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBitmapDataI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("getBitmapData", args)}
}

// Gets a Bitmap Font object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBitmapFont(key string) *BitmapFont{
    return &BitmapFont{self.Object.Call("getBitmapFont", key)}
}

// Gets a Bitmap Font object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBitmapFontI(args ...interface{}) *BitmapFont{
    return &BitmapFont{self.Object.Call("getBitmapFont", args)}
}

// Gets a JSON object from the cache.
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

// Gets a JSON object from the cache.
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

// Gets a JSON object from the cache.
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

// Gets an XML object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetXML(key string) interface{}{
    return self.Object.Call("getXML", key)
}

// Gets an XML object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetXMLI(args ...interface{}) interface{}{
    return self.Object.Call("getXML", args)
}

// Gets a Phaser.Video object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetVideo(key string) *Video{
    return &Video{self.Object.Call("getVideo", key)}
}

// Gets a Phaser.Video object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetVideoI(args ...interface{}) *Video{
    return &Video{self.Object.Call("getVideo", args)}
}

// Gets a fragment shader object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetShader(key string) string{
    return self.Object.Call("getShader", key).String()
}

// Gets a fragment shader object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetShaderI(args ...interface{}) string{
    return self.Object.Call("getShader", args).String()
}

// Gets a RenderTexture object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetRenderTexture(key string) interface{}{
    return self.Object.Call("getRenderTexture", key)
}

// Gets a RenderTexture object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetRenderTextureI(args ...interface{}) interface{}{
    return self.Object.Call("getRenderTexture", args)
}

// Gets a PIXI.BaseTexture by key from the given Cache.
func (self *Cache) GetBaseTexture(key string) *BaseTexture{
    return &BaseTexture{self.Object.Call("getBaseTexture", key)}
}

// Gets a PIXI.BaseTexture by key from the given Cache.
func (self *Cache) GetBaseTexture1O(key string, cache int) *BaseTexture{
    return &BaseTexture{self.Object.Call("getBaseTexture", key, cache)}
}

// Gets a PIXI.BaseTexture by key from the given Cache.
func (self *Cache) GetBaseTextureI(args ...interface{}) *BaseTexture{
    return &BaseTexture{self.Object.Call("getBaseTexture", args)}
}

// Get a single frame by key. You'd only do this to get the default Frame created for a non-atlas/spritesheet image.
func (self *Cache) GetFrame(key string) *Frame{
    return &Frame{self.Object.Call("getFrame", key)}
}

// Get a single frame by key. You'd only do this to get the default Frame created for a non-atlas/spritesheet image.
func (self *Cache) GetFrame1O(key string, cache int) *Frame{
    return &Frame{self.Object.Call("getFrame", key, cache)}
}

// Get a single frame by key. You'd only do this to get the default Frame created for a non-atlas/spritesheet image.
func (self *Cache) GetFrameI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("getFrame", args)}
}

// Get the total number of frames contained in the FrameData object specified by the given key.
func (self *Cache) GetFrameCount(key string) int{
    return self.Object.Call("getFrameCount", key).Int()
}

// Get the total number of frames contained in the FrameData object specified by the given key.
func (self *Cache) GetFrameCount1O(key string, cache int) int{
    return self.Object.Call("getFrameCount", key, cache).Int()
}

// Get the total number of frames contained in the FrameData object specified by the given key.
func (self *Cache) GetFrameCountI(args ...interface{}) int{
    return self.Object.Call("getFrameCount", args).Int()
}

// Gets a Phaser.FrameData object from the Image Cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetFrameData(key string) *FrameData{
    return &FrameData{self.Object.Call("getFrameData", key)}
}

// Gets a Phaser.FrameData object from the Image Cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetFrameData1O(key string, cache int) *FrameData{
    return &FrameData{self.Object.Call("getFrameData", key, cache)}
}

// Gets a Phaser.FrameData object from the Image Cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetFrameDataI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("getFrameData", args)}
}

// Check if the FrameData for the given key exists in the Image Cache.
func (self *Cache) HasFrameData(key string) bool{
    return self.Object.Call("hasFrameData", key).Bool()
}

// Check if the FrameData for the given key exists in the Image Cache.
func (self *Cache) HasFrameData1O(key string, cache int) bool{
    return self.Object.Call("hasFrameData", key, cache).Bool()
}

// Check if the FrameData for the given key exists in the Image Cache.
func (self *Cache) HasFrameDataI(args ...interface{}) bool{
    return self.Object.Call("hasFrameData", args).Bool()
}

// Replaces a set of frameData with a new Phaser.FrameData object.
func (self *Cache) UpdateFrameData(key string, frameData int) {
    self.Object.Call("updateFrameData", key, frameData)
}

// Replaces a set of frameData with a new Phaser.FrameData object.
func (self *Cache) UpdateFrameData1O(key string, frameData int, cache int) {
    self.Object.Call("updateFrameData", key, frameData, cache)
}

// Replaces a set of frameData with a new Phaser.FrameData object.
func (self *Cache) UpdateFrameDataI(args ...interface{}) {
    self.Object.Call("updateFrameData", args)
}

// Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByIndex(key string, index int) *Frame{
    return &Frame{self.Object.Call("getFrameByIndex", key, index)}
}

// Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByIndex1O(key string, index int, cache int) *Frame{
    return &Frame{self.Object.Call("getFrameByIndex", key, index, cache)}
}

// Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByIndexI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("getFrameByIndex", args)}
}

// Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByName(key string, name string) *Frame{
    return &Frame{self.Object.Call("getFrameByName", key, name)}
}

// Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByName1O(key string, name string, cache int) *Frame{
    return &Frame{self.Object.Call("getFrameByName", key, name, cache)}
}

// Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByNameI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("getFrameByName", args)}
}

// Get a cached object by the URL.
// This only returns a value if you set Cache.autoResolveURL to `true` *before* starting the preload of any assets.
// Be aware that every call to this function makes a DOM src query, so use carefully and double-check for implications in your target browsers/devices.
func (self *Cache) GetURL(url string) interface{}{
    return self.Object.Call("getURL", url)
}

// Get a cached object by the URL.
// This only returns a value if you set Cache.autoResolveURL to `true` *before* starting the preload of any assets.
// Be aware that every call to this function makes a DOM src query, so use carefully and double-check for implications in your target browsers/devices.
func (self *Cache) GetURLI(args ...interface{}) interface{}{
    return self.Object.Call("getURL", args)
}

// Gets all keys used in the requested Cache.
func (self *Cache) GetKeys() []interface{}{
	array00 := self.Object.Call("getKeys")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Gets all keys used in the requested Cache.
func (self *Cache) GetKeys1O(cache int) []interface{}{
	array00 := self.Object.Call("getKeys", cache)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Gets all keys used in the requested Cache.
func (self *Cache) GetKeysI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("getKeys", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Removes a canvas from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveCanvas(key string) {
    self.Object.Call("removeCanvas", key)
}

// Removes a canvas from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveCanvasI(args ...interface{}) {
    self.Object.Call("removeCanvas", args)
}

// Removes an image from the cache.
// 
// You can optionally elect to destroy it as well. This calls BaseTexture.destroy on it.
// 
// Note that this only removes it from the Phaser Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveImage(key string) {
    self.Object.Call("removeImage", key)
}

// Removes an image from the cache.
// 
// You can optionally elect to destroy it as well. This calls BaseTexture.destroy on it.
// 
// Note that this only removes it from the Phaser Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveImage1O(key string, destroyBaseTexture bool) {
    self.Object.Call("removeImage", key, destroyBaseTexture)
}

// Removes an image from the cache.
// 
// You can optionally elect to destroy it as well. This calls BaseTexture.destroy on it.
// 
// Note that this only removes it from the Phaser Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveImageI(args ...interface{}) {
    self.Object.Call("removeImage", args)
}

// Removes a sound from the cache.
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

// Removes a sound from the cache.
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

// Removes a text file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveText(key string) {
    self.Object.Call("removeText", key)
}

// Removes a text file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTextI(args ...interface{}) {
    self.Object.Call("removeText", args)
}

// Removes a physics data file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemovePhysics(key string) {
    self.Object.Call("removePhysics", key)
}

// Removes a physics data file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemovePhysicsI(args ...interface{}) {
    self.Object.Call("removePhysics", args)
}

// Removes a tilemap from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTilemap(key string) {
    self.Object.Call("removeTilemap", key)
}

// Removes a tilemap from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTilemapI(args ...interface{}) {
    self.Object.Call("removeTilemap", args)
}

// Removes a binary file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBinary(key string) {
    self.Object.Call("removeBinary", key)
}

// Removes a binary file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBinaryI(args ...interface{}) {
    self.Object.Call("removeBinary", args)
}

// Removes a bitmap data from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBitmapData(key string) {
    self.Object.Call("removeBitmapData", key)
}

// Removes a bitmap data from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBitmapDataI(args ...interface{}) {
    self.Object.Call("removeBitmapData", args)
}

// Removes a bitmap font from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBitmapFont(key string) {
    self.Object.Call("removeBitmapFont", key)
}

// Removes a bitmap font from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBitmapFontI(args ...interface{}) {
    self.Object.Call("removeBitmapFont", args)
}

// Removes a json object from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveJSON(key string) {
    self.Object.Call("removeJSON", key)
}

// Removes a json object from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveJSONI(args ...interface{}) {
    self.Object.Call("removeJSON", args)
}

// Removes a xml object from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveXML(key string) {
    self.Object.Call("removeXML", key)
}

// Removes a xml object from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveXMLI(args ...interface{}) {
    self.Object.Call("removeXML", args)
}

// Removes a video from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveVideo(key string) {
    self.Object.Call("removeVideo", key)
}

// Removes a video from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveVideoI(args ...interface{}) {
    self.Object.Call("removeVideo", args)
}

// Removes a shader from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveShader(key string) {
    self.Object.Call("removeShader", key)
}

// Removes a shader from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveShaderI(args ...interface{}) {
    self.Object.Call("removeShader", args)
}

// Removes a Render Texture from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveRenderTexture(key string) {
    self.Object.Call("removeRenderTexture", key)
}

// Removes a Render Texture from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveRenderTextureI(args ...interface{}) {
    self.Object.Call("removeRenderTexture", args)
}

// Removes a Sprite Sheet from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveSpriteSheet(key string) {
    self.Object.Call("removeSpriteSheet", key)
}

// Removes a Sprite Sheet from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveSpriteSheetI(args ...interface{}) {
    self.Object.Call("removeSpriteSheet", args)
}

// Removes a Texture Atlas from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTextureAtlas(key string) {
    self.Object.Call("removeTextureAtlas", key)
}

// Removes a Texture Atlas from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTextureAtlasI(args ...interface{}) {
    self.Object.Call("removeTextureAtlas", args)
}

// Empties out all of the GL Textures from Images stored in the cache.
// This is called automatically when the WebGL context is lost and then restored.
func (self *Cache) ClearGLTextures() {
    self.Object.Call("clearGLTextures")
}

// Empties out all of the GL Textures from Images stored in the cache.
// This is called automatically when the WebGL context is lost and then restored.
func (self *Cache) ClearGLTexturesI(args ...interface{}) {
    self.Object.Call("clearGLTextures", args)
}

// Resolves a URL to its absolute form and stores it in Cache._urlMap as long as Cache.autoResolveURL is set to `true`.
// This is then looked-up by the Cache.getURL and Cache.checkURL calls.
func (self *Cache) _resolveURL(url string) string{
    return self.Object.Call("_resolveURL", url).String()
}

// Resolves a URL to its absolute form and stores it in Cache._urlMap as long as Cache.autoResolveURL is set to `true`.
// This is then looked-up by the Cache.getURL and Cache.checkURL calls.
func (self *Cache) _resolveURL1O(url string, data interface{}) string{
    return self.Object.Call("_resolveURL", url, data).String()
}

// Resolves a URL to its absolute form and stores it in Cache._urlMap as long as Cache.autoResolveURL is set to `true`.
// This is then looked-up by the Cache.getURL and Cache.checkURL calls.
func (self *Cache) _resolveURLI(args ...interface{}) string{
    return self.Object.Call("_resolveURL", args).String()
}

// Clears the cache. Removes every local cache object reference.
// If an object in the cache has a `destroy` method it will also be called.
func (self *Cache) Destroy() {
    self.Object.Call("destroy")
}

// Clears the cache. Removes every local cache object reference.
// If an object in the cache has a `destroy` method it will also be called.
func (self *Cache) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
