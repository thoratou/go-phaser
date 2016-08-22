// Automatic generation for Phaser.Cache
// generated file Cache.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

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


// Local reference to game.
func (self *Cache) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *Cache) SetGame(member *Game) {
    self.Set("game", member)
}

// Automatically resolve resource URLs to absolute paths for use with the Cache.getURL method.
func (self *Cache) GetAutoResolveURL() bool{
    return self.Get("autoResolveURL").Bool()
}

// Automatically resolve resource URLs to absolute paths for use with the Cache.getURL method.
func (self *Cache) SetAutoResolveURL(member bool) {
    self.Set("autoResolveURL", member)
}

// This event is dispatched when the sound system is unlocked via a touch event on cellular devices.
func (self *Cache) GetOnSoundUnlock() *Signal{
    return &Signal{self.Get("onSoundUnlock")}
}

// This event is dispatched when the sound system is unlocked via a touch event on cellular devices.
func (self *Cache) SetOnSoundUnlock(member *Signal) {
    self.Set("onSoundUnlock", member)
}

// 
func (self *Cache) GetCANVAS() int{
    return self.Get("CANVAS").Int()
}

// 
func (self *Cache) SetCANVAS(member int) {
    self.Set("CANVAS", member)
}

// 
func (self *Cache) GetIMAGE() int{
    return self.Get("IMAGE").Int()
}

// 
func (self *Cache) SetIMAGE(member int) {
    self.Set("IMAGE", member)
}

// 
func (self *Cache) GetTEXTURE() int{
    return self.Get("TEXTURE").Int()
}

// 
func (self *Cache) SetTEXTURE(member int) {
    self.Set("TEXTURE", member)
}

// 
func (self *Cache) GetSOUND() int{
    return self.Get("SOUND").Int()
}

// 
func (self *Cache) SetSOUND(member int) {
    self.Set("SOUND", member)
}

// 
func (self *Cache) GetTEXT() int{
    return self.Get("TEXT").Int()
}

// 
func (self *Cache) SetTEXT(member int) {
    self.Set("TEXT", member)
}

// 
func (self *Cache) GetPHYSICS() int{
    return self.Get("PHYSICS").Int()
}

// 
func (self *Cache) SetPHYSICS(member int) {
    self.Set("PHYSICS", member)
}

// 
func (self *Cache) GetTILEMAP() int{
    return self.Get("TILEMAP").Int()
}

// 
func (self *Cache) SetTILEMAP(member int) {
    self.Set("TILEMAP", member)
}

// 
func (self *Cache) GetBINARY() int{
    return self.Get("BINARY").Int()
}

// 
func (self *Cache) SetBINARY(member int) {
    self.Set("BINARY", member)
}

// 
func (self *Cache) GetBITMAPDATA() int{
    return self.Get("BITMAPDATA").Int()
}

// 
func (self *Cache) SetBITMAPDATA(member int) {
    self.Set("BITMAPDATA", member)
}

// 
func (self *Cache) GetBITMAPFONT() int{
    return self.Get("BITMAPFONT").Int()
}

// 
func (self *Cache) SetBITMAPFONT(member int) {
    self.Set("BITMAPFONT", member)
}

// 
func (self *Cache) GetJSON() int{
    return self.Get("JSON").Int()
}

// 
func (self *Cache) SetJSON(member int) {
    self.Set("JSON", member)
}

// 
func (self *Cache) GetXML() int{
    return self.Get("XML").Int()
}

// 
func (self *Cache) SetXML(member int) {
    self.Set("XML", member)
}

// 
func (self *Cache) GetVIDEO() int{
    return self.Get("VIDEO").Int()
}

// 
func (self *Cache) SetVIDEO(member int) {
    self.Set("VIDEO", member)
}

// 
func (self *Cache) GetSHADER() int{
    return self.Get("SHADER").Int()
}

// 
func (self *Cache) SetSHADER(member int) {
    self.Set("SHADER", member)
}

// 
func (self *Cache) GetRENDER_TEXTURE() int{
    return self.Get("RENDER_TEXTURE").Int()
}

// 
func (self *Cache) SetRENDER_TEXTURE(member int) {
    self.Set("RENDER_TEXTURE", member)
}

// The default image used for a texture when no other is specified.
func (self *Cache) GetDEFAULT() *Texture{
    return &Texture{self.Get("DEFAULT")}
}

// The default image used for a texture when no other is specified.
func (self *Cache) SetDEFAULT(member *Texture) {
    self.Set("DEFAULT", member)
}

// The default image used for a texture when the source image is missing.
func (self *Cache) GetMISSING() *Texture{
    return &Texture{self.Get("MISSING")}
}

// The default image used for a texture when the source image is missing.
func (self *Cache) SetMISSING(member *Texture) {
    self.Set("MISSING", member)
}



// Add a new canvas object in to the cache.
func (self *Cache) AddCanvasI(args ...interface{}) {
    self.Call("addCanvas", args)
}

// Adds an Image file into the Cache. The file must have already been loaded, typically via Phaser.Loader, but can also have been loaded into the DOM.
// If an image already exists in the cache with the same key then it is removed and destroyed, and the new image inserted in its place.
func (self *Cache) AddImageI(args ...interface{}) interface{}{
    return self.Call("addImage", args)
}

// Adds a default image to be used in special cases such as WebGL Filters.
// It uses the special reserved key of `__default`.
// This method is called automatically when the Cache is created.
// This image is skipped when `Cache.destroy` is called due to its internal requirements.
func (self *Cache) AddDefaultImageI(args ...interface{}) {
    self.Call("addDefaultImage", args)
}

// Adds an image to be used when a key is wrong / missing.
// It uses the special reserved key of `__missing`.
// This method is called automatically when the Cache is created.
// This image is skipped when `Cache.destroy` is called due to its internal requirements.
func (self *Cache) AddMissingImageI(args ...interface{}) {
    self.Call("addMissingImage", args)
}

// Adds a Sound file into the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddSoundI(args ...interface{}) {
    self.Call("addSound", args)
}

// Add a new text data.
func (self *Cache) AddTextI(args ...interface{}) {
    self.Call("addText", args)
}

// Add a new physics data object to the Cache.
func (self *Cache) AddPhysicsDataI(args ...interface{}) {
    self.Call("addPhysicsData", args)
}

// Add a new tilemap to the Cache.
func (self *Cache) AddTilemapI(args ...interface{}) {
    self.Call("addTilemap", args)
}

// Add a binary object in to the cache.
func (self *Cache) AddBinaryI(args ...interface{}) {
    self.Call("addBinary", args)
}

// Add a BitmapData object to the cache.
func (self *Cache) AddBitmapDataI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Call("addBitmapData", args)}
}

// Add a new Bitmap Font to the Cache.
func (self *Cache) AddBitmapFontI(args ...interface{}) {
    self.Call("addBitmapFont", args)
}

// Add a new json object into the cache.
func (self *Cache) AddJSONI(args ...interface{}) {
    self.Call("addJSON", args)
}

// Add a new xml object into the cache.
func (self *Cache) AddXMLI(args ...interface{}) {
    self.Call("addXML", args)
}

// Adds a Video file into the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddVideoI(args ...interface{}) {
    self.Call("addVideo", args)
}

// Adds a Fragment Shader in to the Cache. The file must have already been loaded, typically via Phaser.Loader.
func (self *Cache) AddShaderI(args ...interface{}) {
    self.Call("addShader", args)
}

// Add a new Phaser.RenderTexture in to the cache.
func (self *Cache) AddRenderTextureI(args ...interface{}) {
    self.Call("addRenderTexture", args)
}

// Add a new sprite sheet in to the cache.
func (self *Cache) AddSpriteSheetI(args ...interface{}) {
    self.Call("addSpriteSheet", args)
}

// Add a new texture atlas to the Cache.
func (self *Cache) AddTextureAtlasI(args ...interface{}) {
    self.Call("addTextureAtlas", args)
}

// Reload a Sound file from the server.
func (self *Cache) ReloadSoundI(args ...interface{}) {
    self.Call("reloadSound", args)
}

// Fires the onSoundUnlock event when the sound has completed reloading.
func (self *Cache) ReloadSoundCompleteI(args ...interface{}) {
    self.Call("reloadSoundComplete", args)
}

// Updates the sound object in the cache.
func (self *Cache) UpdateSoundI(args ...interface{}) {
    self.Call("updateSound", args)
}

// Add a new decoded sound.
func (self *Cache) DecodedSoundI(args ...interface{}) {
    self.Call("decodedSound", args)
}

// Check if the given sound has finished decoding.
func (self *Cache) IsSoundDecodedI(args ...interface{}) bool{
    return self.Call("isSoundDecoded", args).Bool()
}

// Check if the given sound is ready for playback.
// A sound is considered ready when it has finished decoding and the device is no longer touch locked.
func (self *Cache) IsSoundReadyI(args ...interface{}) bool{
    return self.Call("isSoundReady", args).Bool()
}

// Checks if a key for the given cache object type exists.
func (self *Cache) CheckKeyI(args ...interface{}) bool{
    return self.Call("checkKey", args).Bool()
}

// Checks if the given URL has been loaded into the Cache.
// This method will only work if Cache.autoResolveURL was set to `true` before any preloading took place.
// The method will make a DOM src call to the URL given, so please be aware of this for certain file types, such as Sound files on Firefox
// which may cause double-load instances.
func (self *Cache) CheckURLI(args ...interface{}) bool{
    return self.Call("checkURL", args).Bool()
}

// Checks if the given key exists in the Canvas Cache.
func (self *Cache) CheckCanvasKeyI(args ...interface{}) bool{
    return self.Call("checkCanvasKey", args).Bool()
}

// Checks if the given key exists in the Image Cache. Note that this also includes Texture Atlases, Sprite Sheets and Retro Fonts.
func (self *Cache) CheckImageKeyI(args ...interface{}) bool{
    return self.Call("checkImageKey", args).Bool()
}

// Checks if the given key exists in the Texture Cache.
func (self *Cache) CheckTextureKeyI(args ...interface{}) bool{
    return self.Call("checkTextureKey", args).Bool()
}

// Checks if the given key exists in the Sound Cache.
func (self *Cache) CheckSoundKeyI(args ...interface{}) bool{
    return self.Call("checkSoundKey", args).Bool()
}

// Checks if the given key exists in the Text Cache.
func (self *Cache) CheckTextKeyI(args ...interface{}) bool{
    return self.Call("checkTextKey", args).Bool()
}

// Checks if the given key exists in the Physics Cache.
func (self *Cache) CheckPhysicsKeyI(args ...interface{}) bool{
    return self.Call("checkPhysicsKey", args).Bool()
}

// Checks if the given key exists in the Tilemap Cache.
func (self *Cache) CheckTilemapKeyI(args ...interface{}) bool{
    return self.Call("checkTilemapKey", args).Bool()
}

// Checks if the given key exists in the Binary Cache.
func (self *Cache) CheckBinaryKeyI(args ...interface{}) bool{
    return self.Call("checkBinaryKey", args).Bool()
}

// Checks if the given key exists in the BitmapData Cache.
func (self *Cache) CheckBitmapDataKeyI(args ...interface{}) bool{
    return self.Call("checkBitmapDataKey", args).Bool()
}

// Checks if the given key exists in the BitmapFont Cache.
func (self *Cache) CheckBitmapFontKeyI(args ...interface{}) bool{
    return self.Call("checkBitmapFontKey", args).Bool()
}

// Checks if the given key exists in the JSON Cache.
func (self *Cache) CheckJSONKeyI(args ...interface{}) bool{
    return self.Call("checkJSONKey", args).Bool()
}

// Checks if the given key exists in the XML Cache.
func (self *Cache) CheckXMLKeyI(args ...interface{}) bool{
    return self.Call("checkXMLKey", args).Bool()
}

// Checks if the given key exists in the Video Cache.
func (self *Cache) CheckVideoKeyI(args ...interface{}) bool{
    return self.Call("checkVideoKey", args).Bool()
}

// Checks if the given key exists in the Fragment Shader Cache.
func (self *Cache) CheckShaderKeyI(args ...interface{}) bool{
    return self.Call("checkShaderKey", args).Bool()
}

// Checks if the given key exists in the Render Texture Cache.
func (self *Cache) CheckRenderTextureKeyI(args ...interface{}) bool{
    return self.Call("checkRenderTextureKey", args).Bool()
}

// Get an item from a cache based on the given key and property.
// 
// This method is mostly used internally by other Cache methods such as `getImage` but is exposed
// publicly for your own use as well.
func (self *Cache) GetItemI(args ...interface{}) interface{}{
    return self.Call("getItem", args)
}

// Gets a Canvas object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetCanvasI(args ...interface{}) interface{}{
    return self.Call("getCanvas", args)
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
    return &Image{self.Call("getImage", args)}
}

// Get a single texture frame by key.
// 
// You'd only do this to get the default Frame created for a non-atlas / spritesheet image.
func (self *Cache) GetTextureFrameI(args ...interface{}) *Frame{
    return &Frame{self.Call("getTextureFrame", args)}
}

// Gets a Phaser.Sound object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetSoundI(args ...interface{}) *Sound{
    return &Sound{self.Call("getSound", args)}
}

// Gets a raw Sound data object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetSoundDataI(args ...interface{}) interface{}{
    return self.Call("getSoundData", args)
}

// Gets a Text object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetTextI(args ...interface{}) interface{}{
    return self.Call("getText", args)
}

// Gets a Physics Data object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
// 
// You can get either the entire data set, a single object or a single fixture of an object from it.
func (self *Cache) GetPhysicsDataI(args ...interface{}) interface{}{
    return self.Call("getPhysicsData", args)
}

// Gets a raw Tilemap data object from the cache. This will be in either CSV or JSON format.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetTilemapDataI(args ...interface{}) interface{}{
    return self.Call("getTilemapData", args)
}

// Gets a binary object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBinaryI(args ...interface{}) interface{}{
    return self.Call("getBinary", args)
}

// Gets a BitmapData object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBitmapDataI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Call("getBitmapData", args)}
}

// Gets a Bitmap Font object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetBitmapFontI(args ...interface{}) *BitmapFont{
    return &BitmapFont{self.Call("getBitmapFont", args)}
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
    return self.Call("getJSON", args)
}

// Gets an XML object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetXMLI(args ...interface{}) interface{}{
    return self.Call("getXML", args)
}

// Gets a Phaser.Video object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetVideoI(args ...interface{}) *Video{
    return &Video{self.Call("getVideo", args)}
}

// Gets a fragment shader object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetShaderI(args ...interface{}) string{
    return self.Call("getShader", args).String()
}

// Gets a RenderTexture object from the cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetRenderTextureI(args ...interface{}) interface{}{
    return self.Call("getRenderTexture", args)
}

// Gets a PIXI.BaseTexture by key from the given Cache.
func (self *Cache) GetBaseTextureI(args ...interface{}) *BaseTexture{
    return &BaseTexture{self.Call("getBaseTexture", args)}
}

// Get a single frame by key. You'd only do this to get the default Frame created for a non-atlas/spritesheet image.
func (self *Cache) GetFrameI(args ...interface{}) *Frame{
    return &Frame{self.Call("getFrame", args)}
}

// Get the total number of frames contained in the FrameData object specified by the given key.
func (self *Cache) GetFrameCountI(args ...interface{}) int{
    return self.Call("getFrameCount", args).Int()
}

// Gets a Phaser.FrameData object from the Image Cache.
// 
// The object is looked-up based on the key given.
// 
// Note: If the object cannot be found a `console.warn` message is displayed.
func (self *Cache) GetFrameDataI(args ...interface{}) *FrameData{
    return &FrameData{self.Call("getFrameData", args)}
}

// Check if the FrameData for the given key exists in the Image Cache.
func (self *Cache) HasFrameDataI(args ...interface{}) bool{
    return self.Call("hasFrameData", args).Bool()
}

// Replaces a set of frameData with a new Phaser.FrameData object.
func (self *Cache) UpdateFrameDataI(args ...interface{}) {
    self.Call("updateFrameData", args)
}

// Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByIndexI(args ...interface{}) *Frame{
    return &Frame{self.Call("getFrameByIndex", args)}
}

// Get a single frame out of a frameData set by key.
func (self *Cache) GetFrameByNameI(args ...interface{}) *Frame{
    return &Frame{self.Call("getFrameByName", args)}
}

// Get a cached object by the URL.
// This only returns a value if you set Cache.autoResolveURL to `true` *before* starting the preload of any assets.
// Be aware that every call to this function makes a DOM src query, so use carefully and double-check for implications in your target browsers/devices.
func (self *Cache) GetURLI(args ...interface{}) interface{}{
    return self.Call("getURL", args)
}

// Gets all keys used in the requested Cache.
func (self *Cache) GetKeysI(args ...interface{}) []interface{}{
	array := self.Call("getKeys", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Removes a canvas from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveCanvasI(args ...interface{}) {
    self.Call("removeCanvas", args)
}

// Removes an image from the cache.
// 
// You can optionally elect to destroy it as well. This calls BaseTexture.destroy on it.
// 
// Note that this only removes it from the Phaser Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveImageI(args ...interface{}) {
    self.Call("removeImage", args)
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
    self.Call("removeSound", args)
}

// Removes a text file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTextI(args ...interface{}) {
    self.Call("removeText", args)
}

// Removes a physics data file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemovePhysicsI(args ...interface{}) {
    self.Call("removePhysics", args)
}

// Removes a tilemap from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTilemapI(args ...interface{}) {
    self.Call("removeTilemap", args)
}

// Removes a binary file from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBinaryI(args ...interface{}) {
    self.Call("removeBinary", args)
}

// Removes a bitmap data from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBitmapDataI(args ...interface{}) {
    self.Call("removeBitmapData", args)
}

// Removes a bitmap font from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveBitmapFontI(args ...interface{}) {
    self.Call("removeBitmapFont", args)
}

// Removes a json object from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveJSONI(args ...interface{}) {
    self.Call("removeJSON", args)
}

// Removes a xml object from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveXMLI(args ...interface{}) {
    self.Call("removeXML", args)
}

// Removes a video from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveVideoI(args ...interface{}) {
    self.Call("removeVideo", args)
}

// Removes a shader from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveShaderI(args ...interface{}) {
    self.Call("removeShader", args)
}

// Removes a Render Texture from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveRenderTextureI(args ...interface{}) {
    self.Call("removeRenderTexture", args)
}

// Removes a Sprite Sheet from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveSpriteSheetI(args ...interface{}) {
    self.Call("removeSpriteSheet", args)
}

// Removes a Texture Atlas from the cache.
// 
// Note that this only removes it from the Phaser.Cache. If you still have references to the data elsewhere
// then it will persist in memory.
func (self *Cache) RemoveTextureAtlasI(args ...interface{}) {
    self.Call("removeTextureAtlas", args)
}

// Empties out all of the GL Textures from Images stored in the cache.
// This is called automatically when the WebGL context is lost and then restored.
func (self *Cache) ClearGLTexturesI(args ...interface{}) {
    self.Call("clearGLTextures", args)
}

// Resolves a URL to its absolute form and stores it in Cache._urlMap as long as Cache.autoResolveURL is set to `true`.
// This is then looked-up by the Cache.getURL and Cache.checkURL calls.
func (self *Cache) _resolveURLI(args ...interface{}) string{
    return self.Call("_resolveURL", args).String()
}

// Clears the cache. Removes every local cache object reference.
// If an object in the cache has a `destroy` method it will also be called.
func (self *Cache) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
