// Automatic generation for Phaser.Loader
// generated file Loader.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Loader handles loading all external content such as Images, Sounds, Texture Atlases and data files.
// 
// The loader uses a combination of tag loading (eg. Image elements) and XHR and provides progress and completion callbacks.
// 
// Parallel loading (see {@link Phaser.Loader#enableParallel enableParallel}) is supported and enabled by default.
// Load-before behavior of parallel resources is controlled by synchronization points as discussed in {@link Phaser.Loader#withSyncPoint withSyncPoint}.
// 
// Texture Atlases can be created with tools such as [Texture Packer](https://www.codeandweb.com/texturepacker/phaser) and
// [Shoebox](http://renderhjs.net/shoebox/)
type Loader struct {
    *js.Object
}


// Local reference to game.
func (self *Loader) GetGame() Game{
    return Game{self.Get("game")}
}

// Local reference to game.
func (self *Loader) SetGame(member Game) {
    self.Set("game", member)
}

// Local reference to the Phaser.Cache.
func (self *Loader) GetCache() Cache{
    return Cache{self.Get("cache")}
}

// Local reference to the Phaser.Cache.
func (self *Loader) SetCache(member Cache) {
    self.Set("cache", member)
}

// If true all calls to Loader.reset will be ignored. Useful if you need to create a load queue before swapping to a preloader state.
func (self *Loader) GetResetLocked() bool{
    return self.Get("resetLocked").Bool()
}

// If true all calls to Loader.reset will be ignored. Useful if you need to create a load queue before swapping to a preloader state.
func (self *Loader) SetResetLocked(member bool) {
    self.Set("resetLocked", member)
}

// True if the Loader is in the process of loading the queue.
func (self *Loader) GetIsLoading() bool{
    return self.Get("isLoading").Bool()
}

// True if the Loader is in the process of loading the queue.
func (self *Loader) SetIsLoading(member bool) {
    self.Set("isLoading", member)
}

// True if all assets in the queue have finished loading.
func (self *Loader) GetHasLoaded() bool{
    return self.Get("hasLoaded").Bool()
}

// True if all assets in the queue have finished loading.
func (self *Loader) SetHasLoaded(member bool) {
    self.Set("hasLoaded", member)
}

// You can optionally link a progress sprite with {@link Phaser.Loader#setPreloadSprite setPreloadSprite}.
// 
// This property is an object containing: sprite, rect, direction, width and height
func (self *Loader) GetPreloadSprite() interface{}{
    return self.Get("preloadSprite")
}

// You can optionally link a progress sprite with {@link Phaser.Loader#setPreloadSprite setPreloadSprite}.
// 
// This property is an object containing: sprite, rect, direction, width and height
func (self *Loader) SetPreloadSprite(member interface{}) {
    self.Set("preloadSprite", member)
}

// The crossOrigin value applied to loaded images. Very often this needs to be set to 'anonymous'.
func (self *Loader) GetCrossOrigin() interface{}{
    return self.Get("crossOrigin")
}

// The crossOrigin value applied to loaded images. Very often this needs to be set to 'anonymous'.
func (self *Loader) SetCrossOrigin(member interface{}) {
    self.Set("crossOrigin", member)
}

// If you want to append a URL before the path of any asset you can set this here.
// Useful if allowing the asset base url to be configured outside of the game code.
// The string _must_ end with a "/".
func (self *Loader) GetBaseURL() string{
    return self.Get("baseURL").String()
}

// If you want to append a URL before the path of any asset you can set this here.
// Useful if allowing the asset base url to be configured outside of the game code.
// The string _must_ end with a "/".
func (self *Loader) SetBaseURL(member string) {
    self.Set("baseURL", member)
}

// The value of `path`, if set, is placed before any _relative_ file path given. For example:
// 
// `load.path = "images/sprites/";
// load.image("ball", "ball.png");
// load.image("tree", "level1/oaktree.png");
// load.image("boom", "http://server.com/explode.png");`
// 
// Would load the `ball` file from `images/sprites/ball.png` and the tree from
// `images/sprites/level1/oaktree.png` but the file `boom` would load from the URL
// given as it's an absolute URL.
// 
// Please note that the path is added before the filename but *after* the baseURL (if set.)
// 
// The string _must_ end with a "/".
func (self *Loader) GetPath() string{
    return self.Get("path").String()
}

// The value of `path`, if set, is placed before any _relative_ file path given. For example:
// 
// `load.path = "images/sprites/";
// load.image("ball", "ball.png");
// load.image("tree", "level1/oaktree.png");
// load.image("boom", "http://server.com/explode.png");`
// 
// Would load the `ball` file from `images/sprites/ball.png` and the tree from
// `images/sprites/level1/oaktree.png` but the file `boom` would load from the URL
// given as it's an absolute URL.
// 
// Please note that the path is added before the filename but *after* the baseURL (if set.)
// 
// The string _must_ end with a "/".
func (self *Loader) SetPath(member string) {
    self.Set("path", member)
}

// Used to map the application mime-types to to the Accept header in XHR requests.
// If you don't require these mappings, or they cause problems on your server, then
// remove them from the headers object and the XHR request will not try to use them.
func (self *Loader) GetHeaders() interface{}{
    return self.Get("headers")
}

// Used to map the application mime-types to to the Accept header in XHR requests.
// If you don't require these mappings, or they cause problems on your server, then
// remove them from the headers object and the XHR request will not try to use them.
func (self *Loader) SetHeaders(member interface{}) {
    self.Set("headers", member)
}

// This event is dispatched when the loading process starts: before the first file has been requested,
// but after all the initial packs have been loaded.
func (self *Loader) GetOnLoadStart() Signal{
    return Signal{self.Get("onLoadStart")}
}

// This event is dispatched when the loading process starts: before the first file has been requested,
// but after all the initial packs have been loaded.
func (self *Loader) SetOnLoadStart(member Signal) {
    self.Set("onLoadStart", member)
}

// This event is dispatched when the final file in the load queue has either loaded or failed.
func (self *Loader) GetOnLoadComplete() Signal{
    return Signal{self.Get("onLoadComplete")}
}

// This event is dispatched when the final file in the load queue has either loaded or failed.
func (self *Loader) SetOnLoadComplete(member Signal) {
    self.Set("onLoadComplete", member)
}

// This event is dispatched when an asset pack has either loaded or failed to load.
// 
// This is called when the asset pack manifest file has loaded and successfully added its contents to the loader queue.
// 
// Params: `(pack key, success?, total packs loaded, total packs)`
func (self *Loader) GetOnPackComplete() Signal{
    return Signal{self.Get("onPackComplete")}
}

// This event is dispatched when an asset pack has either loaded or failed to load.
// 
// This is called when the asset pack manifest file has loaded and successfully added its contents to the loader queue.
// 
// Params: `(pack key, success?, total packs loaded, total packs)`
func (self *Loader) SetOnPackComplete(member Signal) {
    self.Set("onPackComplete", member)
}

// This event is dispatched immediately before a file starts loading.
// It's possible the file may fail (eg. download error, invalid format) after this event is sent.
// 
// Params: `(progress, file key, file url)`
func (self *Loader) GetOnFileStart() Signal{
    return Signal{self.Get("onFileStart")}
}

// This event is dispatched immediately before a file starts loading.
// It's possible the file may fail (eg. download error, invalid format) after this event is sent.
// 
// Params: `(progress, file key, file url)`
func (self *Loader) SetOnFileStart(member Signal) {
    self.Set("onFileStart", member)
}

// This event is dispatched when a file has either loaded or failed to load.
// 
// Any function bound to this will receive the following parameters:
// 
// progress, file key, success?, total loaded files, total files
// 
// Where progress is a number between 1 and 100 (inclusive) representing the percentage of the load.
func (self *Loader) GetOnFileComplete() Signal{
    return Signal{self.Get("onFileComplete")}
}

// This event is dispatched when a file has either loaded or failed to load.
// 
// Any function bound to this will receive the following parameters:
// 
// progress, file key, success?, total loaded files, total files
// 
// Where progress is a number between 1 and 100 (inclusive) representing the percentage of the load.
func (self *Loader) SetOnFileComplete(member Signal) {
    self.Set("onFileComplete", member)
}

// This event is dispatched when a file (or pack) errors as a result of the load request.
// 
// For files it will be triggered before `onFileComplete`. For packs it will be triggered before `onPackComplete`.
// 
// Params: `(file key, file)`
func (self *Loader) GetOnFileError() Signal{
    return Signal{self.Get("onFileError")}
}

// This event is dispatched when a file (or pack) errors as a result of the load request.
// 
// For files it will be triggered before `onFileComplete`. For packs it will be triggered before `onPackComplete`.
// 
// Params: `(file key, file)`
func (self *Loader) SetOnFileError(member Signal) {
    self.Set("onFileError", member)
}

// If true and if the browser supports XDomainRequest, it will be used in preference for XHR.
// 
// This is only relevant for IE 9 and should _only_ be enabled for IE 9 clients when required by the server/CDN.
func (self *Loader) GetUseXDomainRequest() bool{
    return self.Get("useXDomainRequest").Bool()
}

// If true and if the browser supports XDomainRequest, it will be used in preference for XHR.
// 
// This is only relevant for IE 9 and should _only_ be enabled for IE 9 clients when required by the server/CDN.
func (self *Loader) SetUseXDomainRequest(member bool) {
    self.Set("useXDomainRequest", member)
}

// If true (the default) then parallel downloading will be enabled.
// 
// To disable all parallel downloads this must be set to false prior to any resource being loaded.
func (self *Loader) GetEnableParallel() bool{
    return self.Get("enableParallel").Bool()
}

// If true (the default) then parallel downloading will be enabled.
// 
// To disable all parallel downloads this must be set to false prior to any resource being loaded.
func (self *Loader) SetEnableParallel(member bool) {
    self.Set("enableParallel", member)
}

// The number of concurrent / parallel resources to try and fetch at once.
// 
// Many current browsers limit 6 requests per domain; this is slightly conservative.
func (self *Loader) GetMaxParallelDownloads() int{
    return self.Get("maxParallelDownloads").Int()
}

// The number of concurrent / parallel resources to try and fetch at once.
// 
// Many current browsers limit 6 requests per domain; this is slightly conservative.
func (self *Loader) SetMaxParallelDownloads(member int) {
    self.Set("maxParallelDownloads", member)
}

// A counter: if more than zero, files will be automatically added as a synchronization point.
func (self *Loader) Get_withSyncPointDepth() interface{}{
    return self.Get("_withSyncPointDepth")
}

// A counter: if more than zero, files will be automatically added as a synchronization point.
func (self *Loader) Set_withSyncPointDepth(member interface{}) {
    self.Set("_withSyncPointDepth", member)
}

// 
func (self *Loader) GetTEXTURE_ATLAS_JSON_ARRAY() float64{
    return self.Get("TEXTURE_ATLAS_JSON_ARRAY").Float()
}

// 
func (self *Loader) SetTEXTURE_ATLAS_JSON_ARRAY(member float64) {
    self.Set("TEXTURE_ATLAS_JSON_ARRAY", member)
}

// 
func (self *Loader) GetTEXTURE_ATLAS_JSON_HASH() float64{
    return self.Get("TEXTURE_ATLAS_JSON_HASH").Float()
}

// 
func (self *Loader) SetTEXTURE_ATLAS_JSON_HASH(member float64) {
    self.Set("TEXTURE_ATLAS_JSON_HASH", member)
}

// 
func (self *Loader) GetTEXTURE_ATLAS_XML_STARLING() float64{
    return self.Get("TEXTURE_ATLAS_XML_STARLING").Float()
}

// 
func (self *Loader) SetTEXTURE_ATLAS_XML_STARLING(member float64) {
    self.Set("TEXTURE_ATLAS_XML_STARLING", member)
}

// 
func (self *Loader) GetPHYSICS_LIME_CORONA_JSON() float64{
    return self.Get("PHYSICS_LIME_CORONA_JSON").Float()
}

// 
func (self *Loader) SetPHYSICS_LIME_CORONA_JSON(member float64) {
    self.Set("PHYSICS_LIME_CORONA_JSON", member)
}

// 
func (self *Loader) GetPHYSICS_PHASER_JSON() float64{
    return self.Get("PHYSICS_PHASER_JSON").Float()
}

// 
func (self *Loader) SetPHYSICS_PHASER_JSON(member float64) {
    self.Set("PHYSICS_PHASER_JSON", member)
}

// 
func (self *Loader) GetTEXTURE_ATLAS_JSON_PYXEL() float64{
    return self.Get("TEXTURE_ATLAS_JSON_PYXEL").Float()
}

// 
func (self *Loader) SetTEXTURE_ATLAS_JSON_PYXEL(member float64) {
    self.Set("TEXTURE_ATLAS_JSON_PYXEL", member)
}

// The non-rounded load progress value (from 0.0 to 100.0).
// 
// A general indicator of the progress.
// It is possible for the progress to decrease, after `onLoadStart`, if more files are dynamically added.
func (self *Loader) GetProgressFloat() interface{}{
    return self.Get("progressFloat")
}

// The non-rounded load progress value (from 0.0 to 100.0).
// 
// A general indicator of the progress.
// It is possible for the progress to decrease, after `onLoadStart`, if more files are dynamically added.
func (self *Loader) SetProgressFloat(member interface{}) {
    self.Set("progressFloat", member)
}

// The rounded load progress percentage value (from 0 to 100). See {@link Phaser.Loader#progressFloat}.
func (self *Loader) GetProgress() interface{}{
    return self.Get("progress")
}

// The rounded load progress percentage value (from 0 to 100). See {@link Phaser.Loader#progressFloat}.
func (self *Loader) SetProgress(member interface{}) {
    self.Set("progress", member)
}



// Set a Sprite to be a "preload" sprite by passing it to this method.
// 
// A "preload" sprite will have its width or height crop adjusted based on the percentage of the loader in real-time.
// This allows you to easily make loading bars for games.
// 
// The sprite will automatically be made visible when calling this.
func (self *Loader) SetPreloadSpriteI(args ...interface{}) {
    self.Call("setPreloadSprite", args)
}

// Called automatically by ScaleManager when the game resizes in RESIZE scalemode.
// 
// This can be used to adjust the preloading sprite size, eg.
func (self *Loader) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// Check whether a file/asset with a specific key is queued to be loaded.
// 
// To access a loaded asset use Phaser.Cache, eg. {@link Phaser.Cache#checkImageKey}
func (self *Loader) CheckKeyExistsI(args ...interface{}) bool{
    return self.Call("checkKeyExists", args).Bool()
}

// Get the queue-index of the file/asset with a specific key.
// 
// Only assets in the download file queue will be found.
func (self *Loader) GetAssetIndexI(args ...interface{}) float64{
    return self.Call("getAssetIndex", args).Float()
}

// Find a file/asset with a specific key.
// 
// Only assets in the download file queue will be found.
func (self *Loader) GetAssetI(args ...interface{}) interface{}{
    return self.Call("getAsset", args)
}

// Reset the loader and clear any queued assets. If `Loader.resetLocked` is true this operation will abort.
// 
// This will abort any loading and clear any queued assets.
// 
// Optionally you can clear any associated events.
func (self *Loader) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Internal function that adds a new entry to the file list. Do not call directly.
func (self *Loader) AddToFileListI(args ...interface{}) Loader{
    return Loader{self.Call("addToFileList", args)}
}

// Internal function that replaces an existing entry in the file list with a new one. Do not call directly.
func (self *Loader) ReplaceInFileListI(args ...interface{}) {
    self.Call("replaceInFileList", args)
}

// Add a JSON resource pack ('packfile') to the Loader.
// 
// A packfile is a JSON file that contains a list of assets to the be loaded.
// Please see the example 'loader/asset pack' in the Phaser Examples repository.
// 
// Packs are always put before the first non-pack file that is not loaded / loading.
// 
// This means that all packs added before any loading has started are added to the front
// of the file queue, in the order added.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// The URL of the packfile can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
func (self *Loader) PackI(args ...interface{}) Loader{
    return Loader{self.Call("pack", args)}
}

// Adds an Image to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// Phaser can load all common image types: png, jpg, gif and any other format the browser can natively handle.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the image via `Cache.getImage(key)`
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified the Loader will take the key and create a filename from that. For example if the key is "alien"
// and no URL is given then the Loader will set the URL to be "alien.png". It will always add `.png` as the extension.
// If you do not desire this action then provide a URL.
func (self *Loader) ImageI(args ...interface{}) Loader{
    return Loader{self.Call("image", args)}
}

// Adds an array of images to the current load queue.
// 
// It works by passing each element of the array to the Loader.image method.
// 
// The files are **not** loaded immediately after calling this method. The files are added to the queue ready to be loaded when the loader starts.
// 
// Phaser can load all common image types: png, jpg, gif and any other format the browser can natively handle.
// 
// The keys must be unique Strings. They are used to add the files to the Phaser.Cache upon successful load.
// 
// Retrieve the images via `Cache.getImage(key)`
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified the Loader will take the key and create a filename from that. For example if the key is "alien"
// and no URL is given then the Loader will set the URL to be "alien.png". It will always add `.png` as the extension.
// If you do not desire this action then provide a URL.
func (self *Loader) ImagesI(args ...interface{}) Loader{
    return Loader{self.Call("images", args)}
}

// Adds a Text file to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getText(key)`
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified the Loader will take the key and create a filename from that. For example if the key is "alien"
// and no URL is given then the Loader will set the URL to be "alien.txt". It will always add `.txt` as the extension.
// If you do not desire this action then provide a URL.
func (self *Loader) TextI(args ...interface{}) Loader{
    return Loader{self.Call("text", args)}
}

// Adds a JSON file to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getJSON(key)`. JSON files are automatically parsed upon load.
// If you need to control when the JSON is parsed then use `Loader.text` instead and parse the text file as needed.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified the Loader will take the key and create a filename from that. For example if the key is "alien"
// and no URL is given then the Loader will set the URL to be "alien.json". It will always add `.json` as the extension.
// If you do not desire this action then provide a URL.
func (self *Loader) JsonI(args ...interface{}) Loader{
    return Loader{self.Call("json", args)}
}

// Adds a fragment shader file to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getShader(key)`.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified the Loader will take the key and create a filename from that. For example if the key is "blur"
// and no URL is given then the Loader will set the URL to be "blur.frag". It will always add `.frag` as the extension.
// If you do not desire this action then provide a URL.
func (self *Loader) ShaderI(args ...interface{}) Loader{
    return Loader{self.Call("shader", args)}
}

// Adds an XML file to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getXML(key)`.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified the Loader will take the key and create a filename from that. For example if the key is "alien"
// and no URL is given then the Loader will set the URL to be "alien.xml". It will always add `.xml` as the extension.
// If you do not desire this action then provide a URL.
func (self *Loader) XmlI(args ...interface{}) Loader{
    return Loader{self.Call("xml", args)}
}

// Adds a JavaScript file to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// The key must be a unique String.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified the Loader will take the key and create a filename from that. For example if the key is "alien"
// and no URL is given then the Loader will set the URL to be "alien.js". It will always add `.js` as the extension.
// If you do not desire this action then provide a URL.
// 
// Upon successful load the JavaScript is automatically turned into a script tag and executed, so be careful what you load!
// 
// A callback, which will be invoked as the script tag has been created, can also be specified.
// The callback must return relevant `data`.
func (self *Loader) ScriptI(args ...interface{}) Loader{
    return Loader{self.Call("script", args)}
}

// Adds a binary file to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getBinary(key)`.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified the Loader will take the key and create a filename from that. For example if the key is "alien"
// and no URL is given then the Loader will set the URL to be "alien.bin". It will always add `.bin` as the extension.
// If you do not desire this action then provide a URL.
// 
// It will be loaded via xhr with a responseType of "arraybuffer". You can specify an optional callback to process the file after load.
// When the callback is called it will be passed 2 parameters: the key of the file and the file data.
// 
// WARNING: If a callback is specified the data will be set to whatever it returns. Always return the data object, even if you didn't modify it.
func (self *Loader) BinaryI(args ...interface{}) Loader{
    return Loader{self.Call("binary", args)}
}

// Adds a Sprite Sheet to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// To clarify the terminology that Phaser uses: A Sprite Sheet is an image containing frames, usually of an animation, that are all equal
// dimensions and often in sequence. For example if the frame size is 32x32 then every frame in the sprite sheet will be that size.
// Sometimes (outside of Phaser) the term "sprite sheet" is used to refer to a texture atlas.
// A Texture Atlas works by packing together images as best it can, using whatever frame sizes it likes, often with cropping and trimming
// the frames in the process. Software such as Texture Packer, Flash CC or Shoebox all generate texture atlases, not sprite sheets.
// If you've got an atlas then use `Loader.atlas` instead.
// 
// The key must be a unique String. It is used to add the image to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getImage(key)`. Sprite sheets, being image based, live in the same Cache as all other Images.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified the Loader will take the key and create a filename from that. For example if the key is "alien"
// and no URL is given then the Loader will set the URL to be "alien.png". It will always add `.png` as the extension.
// If you do not desire this action then provide a URL.
func (self *Loader) SpritesheetI(args ...interface{}) Loader{
    return Loader{self.Call("spritesheet", args)}
}

// Adds an audio file to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getSound(key)`.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// Mobile warning: There are some mobile devices (certain iPad 2 and iPad Mini revisions) that cannot play 48000 Hz audio.
// When they try to play the audio becomes extremely distorted and buzzes, eventually crashing the sound system.
// The solution is to use a lower encoding rate such as 44100 Hz.
func (self *Loader) AudioI(args ...interface{}) Loader{
    return Loader{self.Call("audio", args)}
}

// Adds an audio sprite file to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Audio Sprites are a combination of audio files and a JSON configuration.
// 
// The JSON follows the format of that created by https://github.com/tonistiigi/audiosprite
// 
// Retrieve the file via `Cache.getSoundData(key)`.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
func (self *Loader) AudiospriteI(args ...interface{}) Loader{
    return Loader{self.Call("audiosprite", args)}
}

// Adds a video file to the current load queue.
// 
// The file is **not** loaded immediately after calling this method. The file is added to the queue ready to be loaded when the loader starts.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getVideo(key)`.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// You don't need to preload a video in order to play it in your game. See `Video.createVideoFromURL` for details.
func (self *Loader) VideoI(args ...interface{}) Loader{
    return Loader{self.Call("video", args)}
}

// Adds a Tile Map data file to the current load queue.
// 
// You can choose to either load the data externally, by providing a URL to a json file.
// Or you can pass in a JSON object or String via the `data` parameter.
// If you pass a String the data is automatically run through `JSON.parse` and then immediately added to the Phaser.Cache.
// 
// If a URL is provided the file is **not** loaded immediately after calling this method, but is added to the load queue.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getTilemapData(key)`. JSON files are automatically parsed upon load.
// If you need to control when the JSON is parsed then use `Loader.text` instead and parse the text file as needed.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified and no data is given then the Loader will take the key and create a filename from that.
// For example if the key is "level1" and no URL or data is given then the Loader will set the URL to be "level1.json".
// If you set the format to be Tilemap.CSV it will set the URL to be "level1.csv" instead.
// 
// If you do not desire this action then provide a URL or data object.
func (self *Loader) TilemapI(args ...interface{}) Loader{
    return Loader{self.Call("tilemap", args)}
}

// Adds a physics data file to the current load queue.
// 
// The data must be in `Lime + Corona` JSON format. [Physics Editor](https://www.codeandweb.com) by code'n'web exports in this format natively.
// 
// You can choose to either load the data externally, by providing a URL to a json file.
// Or you can pass in a JSON object or String via the `data` parameter.
// If you pass a String the data is automatically run through `JSON.parse` and then immediately added to the Phaser.Cache.
// 
// If a URL is provided the file is **not** loaded immediately after calling this method, but is added to the load queue.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getJSON(key)`. JSON files are automatically parsed upon load.
// If you need to control when the JSON is parsed then use `Loader.text` instead and parse the text file as needed.
// 
// The URL can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the URL isn't specified and no data is given then the Loader will take the key and create a filename from that.
// For example if the key is "alien" and no URL or data is given then the Loader will set the URL to be "alien.json".
// It will always use `.json` as the extension.
// 
// If you do not desire this action then provide a URL or data object.
func (self *Loader) PhysicsI(args ...interface{}) Loader{
    return Loader{self.Call("physics", args)}
}

// Adds Bitmap Font files to the current load queue.
// 
// To create the Bitmap Font files you can use:
// 
// BMFont (Windows, free): http://www.angelcode.com/products/bmfont/
// Glyph Designer (OS X, commercial): http://www.71squared.com/en/glyphdesigner
// Littera (Web-based, free): http://kvazars.com/littera/
// 
// You can choose to either load the data externally, by providing a URL to an xml file.
// Or you can pass in an XML object or String via the `xmlData` parameter.
// If you pass a String the data is automatically run through `Loader.parseXML` and then immediately added to the Phaser.Cache.
// 
// If URLs are provided the files are **not** loaded immediately after calling this method, but are added to the load queue.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getBitmapFont(key)`. XML files are automatically parsed upon load.
// If you need to control when the XML is parsed then use `Loader.text` instead and parse the XML file as needed.
// 
// The URLs can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the textureURL isn't specified then the Loader will take the key and create a filename from that.
// For example if the key is "megaFont" and textureURL is null then the Loader will set the URL to be "megaFont.png".
// The same is true for the atlasURL. If atlasURL isn't specified and no atlasData has been provided then the Loader will
// set the atlasURL to be the key. For example if the key is "megaFont" the atlasURL will be set to "megaFont.xml".
// 
// If you do not desire this action then provide URLs and / or a data object.
func (self *Loader) BitmapFontI(args ...interface{}) Loader{
    return Loader{self.Call("bitmapFont", args)}
}

// Adds a Texture Atlas file to the current load queue.
// 
// Unlike `Loader.atlasJSONHash` this call expects the atlas data to be in a JSON Array format.
// 
// To create the Texture Atlas you can use tools such as:
// 
// [Texture Packer](https://www.codeandweb.com/texturepacker/phaser)
// [Shoebox](http://renderhjs.net/shoebox/)
// 
// If using Texture Packer we recommend you enable "Trim sprite names".
// If your atlas software has an option to "rotate" the resulting frames, you must disable it.
// 
// You can choose to either load the data externally, by providing a URL to a json file.
// Or you can pass in a JSON object or String via the `atlasData` parameter.
// If you pass a String the data is automatically run through `JSON.parse` and then immediately added to the Phaser.Cache.
// 
// If URLs are provided the files are **not** loaded immediately after calling this method, but are added to the load queue.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getImage(key)`. JSON files are automatically parsed upon load.
// If you need to control when the JSON is parsed then use `Loader.text` instead and parse the JSON file as needed.
// 
// The URLs can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the textureURL isn't specified then the Loader will take the key and create a filename from that.
// For example if the key is "player" and textureURL is null then the Loader will set the URL to be "player.png".
// The same is true for the atlasURL. If atlasURL isn't specified and no atlasData has been provided then the Loader will
// set the atlasURL to be the key. For example if the key is "player" the atlasURL will be set to "player.json".
// 
// If you do not desire this action then provide URLs and / or a data object.
func (self *Loader) AtlasJSONArrayI(args ...interface{}) Loader{
    return Loader{self.Call("atlasJSONArray", args)}
}

// Adds a Texture Atlas file to the current load queue.
// 
// Unlike `Loader.atlas` this call expects the atlas data to be in a JSON Hash format.
// 
// To create the Texture Atlas you can use tools such as:
// 
// [Texture Packer](https://www.codeandweb.com/texturepacker/phaser)
// [Shoebox](http://renderhjs.net/shoebox/)
// 
// If using Texture Packer we recommend you enable "Trim sprite names".
// If your atlas software has an option to "rotate" the resulting frames, you must disable it.
// 
// You can choose to either load the data externally, by providing a URL to a json file.
// Or you can pass in a JSON object or String via the `atlasData` parameter.
// If you pass a String the data is automatically run through `JSON.parse` and then immediately added to the Phaser.Cache.
// 
// If URLs are provided the files are **not** loaded immediately after calling this method, but are added to the load queue.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getImage(key)`. JSON files are automatically parsed upon load.
// If you need to control when the JSON is parsed then use `Loader.text` instead and parse the JSON file as needed.
// 
// The URLs can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the textureURL isn't specified then the Loader will take the key and create a filename from that.
// For example if the key is "player" and textureURL is null then the Loader will set the URL to be "player.png".
// The same is true for the atlasURL. If atlasURL isn't specified and no atlasData has been provided then the Loader will
// set the atlasURL to be the key. For example if the key is "player" the atlasURL will be set to "player.json".
// 
// If you do not desire this action then provide URLs and / or a data object.
func (self *Loader) AtlasJSONHashI(args ...interface{}) Loader{
    return Loader{self.Call("atlasJSONHash", args)}
}

// Adds a Texture Atlas file to the current load queue.
// 
// This call expects the atlas data to be in the Starling XML data format.
// 
// To create the Texture Atlas you can use tools such as:
// 
// [Texture Packer](https://www.codeandweb.com/texturepacker/phaser)
// [Shoebox](http://renderhjs.net/shoebox/)
// 
// If using Texture Packer we recommend you enable "Trim sprite names".
// If your atlas software has an option to "rotate" the resulting frames, you must disable it.
// 
// You can choose to either load the data externally, by providing a URL to an xml file.
// Or you can pass in an XML object or String via the `atlasData` parameter.
// If you pass a String the data is automatically run through `Loader.parseXML` and then immediately added to the Phaser.Cache.
// 
// If URLs are provided the files are **not** loaded immediately after calling this method, but are added to the load queue.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getImage(key)`. XML files are automatically parsed upon load.
// If you need to control when the XML is parsed then use `Loader.text` instead and parse the XML file as needed.
// 
// The URLs can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the textureURL isn't specified then the Loader will take the key and create a filename from that.
// For example if the key is "player" and textureURL is null then the Loader will set the URL to be "player.png".
// The same is true for the atlasURL. If atlasURL isn't specified and no atlasData has been provided then the Loader will
// set the atlasURL to be the key. For example if the key is "player" the atlasURL will be set to "player.xml".
// 
// If you do not desire this action then provide URLs and / or a data object.
func (self *Loader) AtlasXMLI(args ...interface{}) Loader{
    return Loader{self.Call("atlasXML", args)}
}

// Adds a Texture Atlas file to the current load queue.
// 
// To create the Texture Atlas you can use tools such as:
// 
// [Texture Packer](https://www.codeandweb.com/texturepacker/phaser)
// [Shoebox](http://renderhjs.net/shoebox/)
// 
// If using Texture Packer we recommend you enable "Trim sprite names".
// If your atlas software has an option to "rotate" the resulting frames, you must disable it.
// 
// You can choose to either load the data externally, by providing a URL to a json file.
// Or you can pass in a JSON object or String via the `atlasData` parameter.
// If you pass a String the data is automatically run through `JSON.parse` and then immediately added to the Phaser.Cache.
// 
// If URLs are provided the files are **not** loaded immediately after calling this method, but are added to the load queue.
// 
// The key must be a unique String. It is used to add the file to the Phaser.Cache upon successful load.
// 
// Retrieve the file via `Cache.getImage(key)`. JSON files are automatically parsed upon load.
// If you need to control when the JSON is parsed then use `Loader.text` instead and parse the JSON file as needed.
// 
// The URLs can be relative or absolute. If the URL is relative the `Loader.baseURL` and `Loader.path` values will be prepended to it.
// 
// If the textureURL isn't specified then the Loader will take the key and create a filename from that.
// For example if the key is "player" and textureURL is null then the Loader will set the URL to be "player.png".
// The same is true for the atlasURL. If atlasURL isn't specified and no atlasData has been provided then the Loader will
// set the atlasURL to be the key. For example if the key is "player" the atlasURL will be set to "player.json".
// 
// If you do not desire this action then provide URLs and / or a data object.
func (self *Loader) AtlasI(args ...interface{}) Loader{
    return Loader{self.Call("atlas", args)}
}

// Add a synchronization point to the assets/files added within the supplied callback.
// 
// A synchronization point denotes that an asset _must_ be completely loaded before
// subsequent assets can be loaded. An asset marked as a sync-point does not need to wait
// for previous assets to load (unless they are sync-points). Resources, such as packs, may still
// be downloaded around sync-points, as long as they do not finalize loading.
func (self *Loader) WithSyncPointsI(args ...interface{}) Loader{
    return Loader{self.Call("withSyncPoints", args)}
}

// Add a synchronization point to a specific file/asset in the load queue.
// 
// This has no effect on already loaded assets.
func (self *Loader) AddSyncPointI(args ...interface{}) Loader{
    return Loader{self.Call("addSyncPoint", args)}
}

// Remove a file/asset from the loading queue.
// 
// A file that is loaded or has started loading cannot be removed.
func (self *Loader) RemoveFileI(args ...interface{}) {
    self.Call("removeFile", args)
}

// Remove all file loading requests - this is _insufficient_ to stop current loading. Use `reset` instead.
func (self *Loader) RemoveAllI(args ...interface{}) {
    self.Call("removeAll", args)
}

// Start loading the assets. Normally you don't need to call this yourself as the StateManager will do so.
func (self *Loader) StartI(args ...interface{}) {
    self.Call("start", args)
}

// Process the next item(s) in the file/asset queue.
// 
// Process the queue and start loading enough items to fill up the inflight queue.
// 
// If a sync-file is encountered then subsequent asset processing is delayed until it completes.
// The exception to this rule is that packfiles can be downloaded (but not processed) even if
// there appear other sync files (ie. packs) - this enables multiple packfiles to be fetched in parallel.
// such as during the start phaser.
func (self *Loader) ProcessLoadQueueI(args ...interface{}) {
    self.Call("processLoadQueue", args)
}

// The loading is all finished.
func (self *Loader) FinishedLoadingI(args ...interface{}) {
    self.Call("finishedLoading", args)
}

// Informs the loader that the given file resource has been fetched and processed;
// or such a request has failed.
func (self *Loader) AsyncCompleteI(args ...interface{}) {
    self.Call("asyncComplete", args)
}

// Process pack data. This will usually modify the file list.
func (self *Loader) ProcessPackI(args ...interface{}) {
    self.Call("processPack", args)
}

// Transforms the asset URL.
// 
// The default implementation prepends the baseURL if the url doesn't begin with http or //
func (self *Loader) TransformUrlI(args ...interface{}) string{
    return self.Call("transformUrl", args).String()
}

// Start fetching a resource.
// 
// All code paths, async or otherwise, from this function must return to `asyncComplete`.
func (self *Loader) LoadFileI(args ...interface{}) {
    self.Call("loadFile", args)
}

// Continue async loading through an Image tag.
func (self *Loader) LoadImageTagI(args ...interface{}) {
    self.Call("loadImageTag", args)
}

// Continue async loading through a Video tag.
func (self *Loader) LoadVideoTagI(args ...interface{}) {
    self.Call("loadVideoTag", args)
}

// Continue async loading through an Audio tag.
func (self *Loader) LoadAudioTagI(args ...interface{}) {
    self.Call("loadAudioTag", args)
}

// Starts the xhr loader.
// 
// This is designed specifically to use with asset file processing.
func (self *Loader) XhrLoadI(args ...interface{}) {
    self.Call("xhrLoad", args)
}

// Give a bunch of URLs, return the first URL that has an extension this device thinks it can play.
// 
// It is assumed that the device can play "blob:" or "data:" URIs - There is no mime-type checking on data URIs.
func (self *Loader) GetVideoURLI(args ...interface{}) string{
    return self.Call("getVideoURL", args).String()
}

// Give a bunch of URLs, return the first URL that has an extension this device thinks it can play.
// 
// It is assumed that the device can play "blob:" or "data:" URIs - There is no mime-type checking on data URIs.
func (self *Loader) GetAudioURLI(args ...interface{}) string{
    return self.Call("getAudioURL", args).String()
}

// Error occurred when loading a file.
func (self *Loader) FileErrorI(args ...interface{}) {
    self.Call("fileError", args)
}

// Called when a file/resources had been downloaded and needs to be processed further.
func (self *Loader) FileCompleteI(args ...interface{}) {
    self.Call("fileComplete", args)
}

// Successfully loaded a JSON file - only used for certain types.
func (self *Loader) JsonLoadCompleteI(args ...interface{}) {
    self.Call("jsonLoadComplete", args)
}

// Successfully loaded a CSV file - only used for certain types.
func (self *Loader) CsvLoadCompleteI(args ...interface{}) {
    self.Call("csvLoadComplete", args)
}

// Successfully loaded an XML file - only used for certain types.
func (self *Loader) XmlLoadCompleteI(args ...interface{}) {
    self.Call("xmlLoadComplete", args)
}

// Parses string data as XML.
func (self *Loader) ParseXmlI(args ...interface{}) XMLDocument{
    return XMLDocument{self.Call("parseXml", args)}
}

// Update the loading sprite progress.
func (self *Loader) NextFileI(args ...interface{}) {
    self.Call("nextFile", args)
}

// Returns the number of files that have already been loaded, even if they errored.
func (self *Loader) TotalLoadedFilesI(args ...interface{}) float64{
    return self.Call("totalLoadedFiles", args).Float()
}

// Returns the number of files still waiting to be processed in the load queue. This value decreases as each file in the queue is loaded.
func (self *Loader) TotalQueuedFilesI(args ...interface{}) float64{
    return self.Call("totalQueuedFiles", args).Float()
}

// Returns the number of asset packs that have already been loaded, even if they errored.
func (self *Loader) TotalLoadedPacksI(args ...interface{}) float64{
    return self.Call("totalLoadedPacks", args).Float()
}

// Returns the number of asset packs still waiting to be processed in the load queue. This value decreases as each pack in the queue is loaded.
func (self *Loader) TotalQueuedPacksI(args ...interface{}) float64{
    return self.Call("totalQueuedPacks", args).Float()
}