// Automatic generation for Phaser.Device
// generated file Device.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Detects device support capabilities and is responsible for device initialization - see {@link Phaser.Device.whenReady whenReady}.
// 
// This class represents a singleton object that can be accessed directly as `game.device`
// (or, as a fallback, `Phaser.Device` when a game instance is not available) without the need to instantiate it.
// 
// Unless otherwise noted the device capabilities are only guaranteed after initialization. Initialization
// occurs automatically and is guaranteed complete before {@link Phaser.Game} begins its "boot" phase.
// Feature detection can be modified in the {@link Phaser.Device.onInitialized onInitialized} signal.
// 
// When checking features using the exposed properties only the *truth-iness* of the value should be relied upon
// unless the documentation states otherwise: properties may return `false`, `''`, `null`, or even `undefined`
// when indicating the lack of a feature.
// 
// Uses elements from System.js by MrDoob and Modernizr
type Device struct {
    *js.Object
}


// The time the device became ready.
func (self *Device) GetDeviceReadyAt() int{
    return self.Get("deviceReadyAt").Int()
}

// The time the device became ready.
func (self *Device) SetDeviceReadyAt(member int) {
    self.Set("deviceReadyAt", member)
}

// The time as which initialization has completed.
func (self *Device) GetInitialized() bool{
    return self.Get("initialized").Bool()
}

// The time as which initialization has completed.
func (self *Device) SetInitialized(member bool) {
    self.Set("initialized", member)
}

// Is running on a desktop?
func (self *Device) GetDesktop() bool{
    return self.Get("desktop").Bool()
}

// Is running on a desktop?
func (self *Device) SetDesktop(member bool) {
    self.Set("desktop", member)
}

// Is running on iOS?
func (self *Device) GetIOS() bool{
    return self.Get("iOS").Bool()
}

// Is running on iOS?
func (self *Device) SetIOS(member bool) {
    self.Set("iOS", member)
}

// If running in iOS this will contain the major version number.
func (self *Device) GetIOSVersion() int{
    return self.Get("iOSVersion").Int()
}

// If running in iOS this will contain the major version number.
func (self *Device) SetIOSVersion(member int) {
    self.Set("iOSVersion", member)
}

// Is the game running under CocoonJS?
func (self *Device) GetCocoonJS() bool{
    return self.Get("cocoonJS").Bool()
}

// Is the game running under CocoonJS?
func (self *Device) SetCocoonJS(member bool) {
    self.Set("cocoonJS", member)
}

// Is this game running with CocoonJS.App?
func (self *Device) GetCocoonJSApp() bool{
    return self.Get("cocoonJSApp").Bool()
}

// Is this game running with CocoonJS.App?
func (self *Device) SetCocoonJSApp(member bool) {
    self.Set("cocoonJSApp", member)
}

// Is the game running under Apache Cordova?
func (self *Device) GetCordova() bool{
    return self.Get("cordova").Bool()
}

// Is the game running under Apache Cordova?
func (self *Device) SetCordova(member bool) {
    self.Set("cordova", member)
}

// Is the game running under Node.js?
func (self *Device) GetNode() bool{
    return self.Get("node").Bool()
}

// Is the game running under Node.js?
func (self *Device) SetNode(member bool) {
    self.Set("node", member)
}

// Is the game running under Node-Webkit?
func (self *Device) GetNodeWebkit() bool{
    return self.Get("nodeWebkit").Bool()
}

// Is the game running under Node-Webkit?
func (self *Device) SetNodeWebkit(member bool) {
    self.Set("nodeWebkit", member)
}

// Is the game running under GitHub Electron?
func (self *Device) GetElectron() bool{
    return self.Get("electron").Bool()
}

// Is the game running under GitHub Electron?
func (self *Device) SetElectron(member bool) {
    self.Set("electron", member)
}

// Is the game running under Ejecta?
func (self *Device) GetEjecta() bool{
    return self.Get("ejecta").Bool()
}

// Is the game running under Ejecta?
func (self *Device) SetEjecta(member bool) {
    self.Set("ejecta", member)
}

// Is the game running under the Intel Crosswalk XDK?
func (self *Device) GetCrosswalk() bool{
    return self.Get("crosswalk").Bool()
}

// Is the game running under the Intel Crosswalk XDK?
func (self *Device) SetCrosswalk(member bool) {
    self.Set("crosswalk", member)
}

// Is running on android?
func (self *Device) GetAndroid() bool{
    return self.Get("android").Bool()
}

// Is running on android?
func (self *Device) SetAndroid(member bool) {
    self.Set("android", member)
}

// Is running on chromeOS?
func (self *Device) GetChromeOS() bool{
    return self.Get("chromeOS").Bool()
}

// Is running on chromeOS?
func (self *Device) SetChromeOS(member bool) {
    self.Set("chromeOS", member)
}

// Is running on linux?
func (self *Device) GetLinux() bool{
    return self.Get("linux").Bool()
}

// Is running on linux?
func (self *Device) SetLinux(member bool) {
    self.Set("linux", member)
}

// Is running on macOS?
func (self *Device) GetMacOS() bool{
    return self.Get("macOS").Bool()
}

// Is running on macOS?
func (self *Device) SetMacOS(member bool) {
    self.Set("macOS", member)
}

// Is running on windows?
func (self *Device) GetWindows() bool{
    return self.Get("windows").Bool()
}

// Is running on windows?
func (self *Device) SetWindows(member bool) {
    self.Set("windows", member)
}

// Is running on a Windows Phone?
func (self *Device) GetWindowsPhone() bool{
    return self.Get("windowsPhone").Bool()
}

// Is running on a Windows Phone?
func (self *Device) SetWindowsPhone(member bool) {
    self.Set("windowsPhone", member)
}

// Is canvas available?
func (self *Device) GetCanvas() bool{
    return self.Get("canvas").Bool()
}

// Is canvas available?
func (self *Device) SetCanvas(member bool) {
    self.Set("canvas", member)
}

// True if canvas supports a 'copy' bitblt onto itself when the source and destination regions overlap.
func (self *Device) GetCanvasBitBltShift() bool{
    return self.Get("canvasBitBltShift").Bool()
}

// True if canvas supports a 'copy' bitblt onto itself when the source and destination regions overlap.
func (self *Device) SetCanvasBitBltShift(member bool) {
    self.Set("canvasBitBltShift", member)
}

// Is webGL available?
func (self *Device) GetWebGL() bool{
    return self.Get("webGL").Bool()
}

// Is webGL available?
func (self *Device) SetWebGL(member bool) {
    self.Set("webGL", member)
}

// Is file available?
func (self *Device) GetFile() bool{
    return self.Get("file").Bool()
}

// Is file available?
func (self *Device) SetFile(member bool) {
    self.Set("file", member)
}

// Is fileSystem available?
func (self *Device) GetFileSystem() bool{
    return self.Get("fileSystem").Bool()
}

// Is fileSystem available?
func (self *Device) SetFileSystem(member bool) {
    self.Set("fileSystem", member)
}

// Is localStorage available?
func (self *Device) GetLocalStorage() bool{
    return self.Get("localStorage").Bool()
}

// Is localStorage available?
func (self *Device) SetLocalStorage(member bool) {
    self.Set("localStorage", member)
}

// Is worker available?
func (self *Device) GetWorker() bool{
    return self.Get("worker").Bool()
}

// Is worker available?
func (self *Device) SetWorker(member bool) {
    self.Set("worker", member)
}

// Is css3D available?
func (self *Device) GetCss3D() bool{
    return self.Get("css3D").Bool()
}

// Is css3D available?
func (self *Device) SetCss3D(member bool) {
    self.Set("css3D", member)
}

// Is Pointer Lock available?
func (self *Device) GetPointerLock() bool{
    return self.Get("pointerLock").Bool()
}

// Is Pointer Lock available?
func (self *Device) SetPointerLock(member bool) {
    self.Set("pointerLock", member)
}

// Does the browser support TypedArrays?
func (self *Device) GetTypedArray() bool{
    return self.Get("typedArray").Bool()
}

// Does the browser support TypedArrays?
func (self *Device) SetTypedArray(member bool) {
    self.Set("typedArray", member)
}

// Does the device support the Vibration API?
func (self *Device) GetVibration() bool{
    return self.Get("vibration").Bool()
}

// Does the device support the Vibration API?
func (self *Device) SetVibration(member bool) {
    self.Set("vibration", member)
}

// Does the device support the getUserMedia API?
func (self *Device) GetGetUserMedia() bool{
    return self.Get("getUserMedia").Bool()
}

// Does the device support the getUserMedia API?
func (self *Device) SetGetUserMedia(member bool) {
    self.Set("getUserMedia", member)
}

// Is the browser running in strict mode (false) or quirks mode? (true)
func (self *Device) GetQuirksMode() bool{
    return self.Get("quirksMode").Bool()
}

// Is the browser running in strict mode (false) or quirks mode? (true)
func (self *Device) SetQuirksMode(member bool) {
    self.Set("quirksMode", member)
}

// Is touch available?
func (self *Device) GetTouch() bool{
    return self.Get("touch").Bool()
}

// Is touch available?
func (self *Device) SetTouch(member bool) {
    self.Set("touch", member)
}

// Is mspointer available?
func (self *Device) GetMspointer() bool{
    return self.Get("mspointer").Bool()
}

// Is mspointer available?
func (self *Device) SetMspointer(member bool) {
    self.Set("mspointer", member)
}

// The newest type of Wheel/Scroll event supported: 'wheel', 'mousewheel', 'DOMMouseScroll'
func (self *Device) GetWheelEvent() interface{}{
    return self.Get("wheelEvent")
}

// The newest type of Wheel/Scroll event supported: 'wheel', 'mousewheel', 'DOMMouseScroll'
func (self *Device) SetWheelEvent(member interface{}) {
    self.Set("wheelEvent", member)
}

// Set to true if running in Arora.
func (self *Device) GetArora() bool{
    return self.Get("arora").Bool()
}

// Set to true if running in Arora.
func (self *Device) SetArora(member bool) {
    self.Set("arora", member)
}

// Set to true if running in Chrome.
func (self *Device) GetChrome() bool{
    return self.Get("chrome").Bool()
}

// Set to true if running in Chrome.
func (self *Device) SetChrome(member bool) {
    self.Set("chrome", member)
}

// If running in Chrome this will contain the major version number.
func (self *Device) GetChromeVersion() int{
    return self.Get("chromeVersion").Int()
}

// If running in Chrome this will contain the major version number.
func (self *Device) SetChromeVersion(member int) {
    self.Set("chromeVersion", member)
}

// Set to true if running in Epiphany.
func (self *Device) GetEpiphany() bool{
    return self.Get("epiphany").Bool()
}

// Set to true if running in Epiphany.
func (self *Device) SetEpiphany(member bool) {
    self.Set("epiphany", member)
}

// Set to true if running in Firefox.
func (self *Device) GetFirefox() bool{
    return self.Get("firefox").Bool()
}

// Set to true if running in Firefox.
func (self *Device) SetFirefox(member bool) {
    self.Set("firefox", member)
}

// If running in Firefox this will contain the major version number.
func (self *Device) GetFirefoxVersion() int{
    return self.Get("firefoxVersion").Int()
}

// If running in Firefox this will contain the major version number.
func (self *Device) SetFirefoxVersion(member int) {
    self.Set("firefoxVersion", member)
}

// Set to true if running in Internet Explorer.
func (self *Device) GetIe() bool{
    return self.Get("ie").Bool()
}

// Set to true if running in Internet Explorer.
func (self *Device) SetIe(member bool) {
    self.Set("ie", member)
}

// If running in Internet Explorer this will contain the major version number. Beyond IE10 you should use Device.trident and Device.tridentVersion.
func (self *Device) GetIeVersion() int{
    return self.Get("ieVersion").Int()
}

// If running in Internet Explorer this will contain the major version number. Beyond IE10 you should use Device.trident and Device.tridentVersion.
func (self *Device) SetIeVersion(member int) {
    self.Set("ieVersion", member)
}

// Set to true if running a Trident version of Internet Explorer (IE11+)
func (self *Device) GetTrident() bool{
    return self.Get("trident").Bool()
}

// Set to true if running a Trident version of Internet Explorer (IE11+)
func (self *Device) SetTrident(member bool) {
    self.Set("trident", member)
}

// If running in Internet Explorer 11 this will contain the major version number. See {@link http://msdn.microsoft.com/en-us/library/ie/ms537503(v=vs.85).aspx}
func (self *Device) GetTridentVersion() int{
    return self.Get("tridentVersion").Int()
}

// If running in Internet Explorer 11 this will contain the major version number. See {@link http://msdn.microsoft.com/en-us/library/ie/ms537503(v=vs.85).aspx}
func (self *Device) SetTridentVersion(member int) {
    self.Set("tridentVersion", member)
}

// Set to true if running in Microsoft Edge browser.
func (self *Device) GetEdge() bool{
    return self.Get("edge").Bool()
}

// Set to true if running in Microsoft Edge browser.
func (self *Device) SetEdge(member bool) {
    self.Set("edge", member)
}

// Set to true if running in Mobile Safari.
func (self *Device) GetMobileSafari() bool{
    return self.Get("mobileSafari").Bool()
}

// Set to true if running in Mobile Safari.
func (self *Device) SetMobileSafari(member bool) {
    self.Set("mobileSafari", member)
}

// Set to true if running in Midori.
func (self *Device) GetMidori() bool{
    return self.Get("midori").Bool()
}

// Set to true if running in Midori.
func (self *Device) SetMidori(member bool) {
    self.Set("midori", member)
}

// Set to true if running in Opera.
func (self *Device) GetOpera() bool{
    return self.Get("opera").Bool()
}

// Set to true if running in Opera.
func (self *Device) SetOpera(member bool) {
    self.Set("opera", member)
}

// Set to true if running in Safari.
func (self *Device) GetSafari() bool{
    return self.Get("safari").Bool()
}

// Set to true if running in Safari.
func (self *Device) SetSafari(member bool) {
    self.Set("safari", member)
}

// If running in Safari this will contain the major version number.
func (self *Device) GetSafariVersion() int{
    return self.Get("safariVersion").Int()
}

// If running in Safari this will contain the major version number.
func (self *Device) SetSafariVersion(member int) {
    self.Set("safariVersion", member)
}

// Set to true if running as a WebApp, i.e. within a WebView
func (self *Device) GetWebApp() bool{
    return self.Get("webApp").Bool()
}

// Set to true if running as a WebApp, i.e. within a WebView
func (self *Device) SetWebApp(member bool) {
    self.Set("webApp", member)
}

// Set to true if running in the Silk browser (as used on the Amazon Kindle)
func (self *Device) GetSilk() bool{
    return self.Get("silk").Bool()
}

// Set to true if running in the Silk browser (as used on the Amazon Kindle)
func (self *Device) SetSilk(member bool) {
    self.Set("silk", member)
}

// Are Audio tags available?
func (self *Device) GetAudioData() bool{
    return self.Get("audioData").Bool()
}

// Are Audio tags available?
func (self *Device) SetAudioData(member bool) {
    self.Set("audioData", member)
}

// Is the WebAudio API available?
func (self *Device) GetWebAudio() bool{
    return self.Get("webAudio").Bool()
}

// Is the WebAudio API available?
func (self *Device) SetWebAudio(member bool) {
    self.Set("webAudio", member)
}

// Can this device play ogg files?
func (self *Device) GetOgg() bool{
    return self.Get("ogg").Bool()
}

// Can this device play ogg files?
func (self *Device) SetOgg(member bool) {
    self.Set("ogg", member)
}

// Can this device play opus files?
func (self *Device) GetOpus() bool{
    return self.Get("opus").Bool()
}

// Can this device play opus files?
func (self *Device) SetOpus(member bool) {
    self.Set("opus", member)
}

// Can this device play mp3 files?
func (self *Device) GetMp3() bool{
    return self.Get("mp3").Bool()
}

// Can this device play mp3 files?
func (self *Device) SetMp3(member bool) {
    self.Set("mp3", member)
}

// Can this device play wav files?
func (self *Device) GetWav() bool{
    return self.Get("wav").Bool()
}

// Can this device play wav files?
func (self *Device) SetWav(member bool) {
    self.Set("wav", member)
}

// Can this device play m4a files? True if this device can play m4a files.
func (self *Device) GetM4a() bool{
    return self.Get("m4a").Bool()
}

// Can this device play m4a files? True if this device can play m4a files.
func (self *Device) SetM4a(member bool) {
    self.Set("m4a", member)
}

// Can this device play webm files?
func (self *Device) GetWebm() bool{
    return self.Get("webm").Bool()
}

// Can this device play webm files?
func (self *Device) SetWebm(member bool) {
    self.Set("webm", member)
}

// Can this device play EC-3 Dolby Digital Plus files?
func (self *Device) GetDolby() bool{
    return self.Get("dolby").Bool()
}

// Can this device play EC-3 Dolby Digital Plus files?
func (self *Device) SetDolby(member bool) {
    self.Set("dolby", member)
}

// Can this device play ogg video files?
func (self *Device) GetOggVideo() bool{
    return self.Get("oggVideo").Bool()
}

// Can this device play ogg video files?
func (self *Device) SetOggVideo(member bool) {
    self.Set("oggVideo", member)
}

// Can this device play h264 mp4 video files?
func (self *Device) GetH264Video() bool{
    return self.Get("h264Video").Bool()
}

// Can this device play h264 mp4 video files?
func (self *Device) SetH264Video(member bool) {
    self.Set("h264Video", member)
}

// Can this device play h264 mp4 video files?
func (self *Device) GetMp4Video() bool{
    return self.Get("mp4Video").Bool()
}

// Can this device play h264 mp4 video files?
func (self *Device) SetMp4Video(member bool) {
    self.Set("mp4Video", member)
}

// Can this device play webm video files?
func (self *Device) GetWebmVideo() bool{
    return self.Get("webmVideo").Bool()
}

// Can this device play webm video files?
func (self *Device) SetWebmVideo(member bool) {
    self.Set("webmVideo", member)
}

// Can this device play vp9 video files?
func (self *Device) GetVp9Video() bool{
    return self.Get("vp9Video").Bool()
}

// Can this device play vp9 video files?
func (self *Device) SetVp9Video(member bool) {
    self.Set("vp9Video", member)
}

// Can this device play hls video files?
func (self *Device) GetHlsVideo() bool{
    return self.Get("hlsVideo").Bool()
}

// Can this device play hls video files?
func (self *Device) SetHlsVideo(member bool) {
    self.Set("hlsVideo", member)
}

// Is running on iPhone?
func (self *Device) GetIPhone() bool{
    return self.Get("iPhone").Bool()
}

// Is running on iPhone?
func (self *Device) SetIPhone(member bool) {
    self.Set("iPhone", member)
}

// Is running on iPhone4?
func (self *Device) GetIPhone4() bool{
    return self.Get("iPhone4").Bool()
}

// Is running on iPhone4?
func (self *Device) SetIPhone4(member bool) {
    self.Set("iPhone4", member)
}

// Is running on iPad?
func (self *Device) GetIPad() bool{
    return self.Get("iPad").Bool()
}

// Is running on iPad?
func (self *Device) SetIPad(member bool) {
    self.Set("iPad", member)
}

// PixelRatio of the host device?
func (self *Device) GetPixelRatio() int{
    return self.Get("pixelRatio").Int()
}

// PixelRatio of the host device?
func (self *Device) SetPixelRatio(member int) {
    self.Set("pixelRatio", member)
}

// Is the device big or little endian? (only detected if the browser supports TypedArrays)
func (self *Device) GetLittleEndian() bool{
    return self.Get("littleEndian").Bool()
}

// Is the device big or little endian? (only detected if the browser supports TypedArrays)
func (self *Device) SetLittleEndian(member bool) {
    self.Set("littleEndian", member)
}

// Same value as `littleEndian`.
func (self *Device) GetLITTLE_ENDIAN() bool{
    return self.Get("LITTLE_ENDIAN").Bool()
}

// Same value as `littleEndian`.
func (self *Device) SetLITTLE_ENDIAN(member bool) {
    self.Set("LITTLE_ENDIAN", member)
}

// Does the device context support 32bit pixel manipulation using array buffer views?
func (self *Device) GetSupport32bit() bool{
    return self.Get("support32bit").Bool()
}

// Does the device context support 32bit pixel manipulation using array buffer views?
func (self *Device) SetSupport32bit(member bool) {
    self.Set("support32bit", member)
}

// Does the browser support the Full Screen API?
func (self *Device) GetFullscreen() bool{
    return self.Get("fullscreen").Bool()
}

// Does the browser support the Full Screen API?
func (self *Device) SetFullscreen(member bool) {
    self.Set("fullscreen", member)
}

// If the browser supports the Full Screen API this holds the call you need to use to activate it.
func (self *Device) GetRequestFullscreen() string{
    return self.Get("requestFullscreen").String()
}

// If the browser supports the Full Screen API this holds the call you need to use to activate it.
func (self *Device) SetRequestFullscreen(member string) {
    self.Set("requestFullscreen", member)
}

// If the browser supports the Full Screen API this holds the call you need to use to cancel it.
func (self *Device) GetCancelFullscreen() string{
    return self.Get("cancelFullscreen").String()
}

// If the browser supports the Full Screen API this holds the call you need to use to cancel it.
func (self *Device) SetCancelFullscreen(member string) {
    self.Set("cancelFullscreen", member)
}

// Does the browser support access to the Keyboard during Full Screen mode?
func (self *Device) GetFullscreenKeyboard() bool{
    return self.Get("fullscreenKeyboard").Bool()
}

// Does the browser support access to the Keyboard during Full Screen mode?
func (self *Device) SetFullscreenKeyboard(member bool) {
    self.Set("fullscreenKeyboard", member)
}

// This signal is dispatched after device initialization occurs but before any of the ready
// callbacks (see {@link Phaser.Device.whenReady whenReady}) have been invoked.
// 
// Local "patching" for a particular device can/should be done in this event.
// 
// _Note_: This signal is removed after the device has been readied; if a handler has not been
// added _before_ `new Phaser.Game(..)` it is probably too late.
func (self *Device) GetOnInitialized() *Signal{
    return &Signal{self.Get("onInitialized")}
}

// This signal is dispatched after device initialization occurs but before any of the ready
// callbacks (see {@link Phaser.Device.whenReady whenReady}) have been invoked.
// 
// Local "patching" for a particular device can/should be done in this event.
// 
// _Note_: This signal is removed after the device has been readied; if a handler has not been
// added _before_ `new Phaser.Game(..)` it is probably too late.
func (self *Device) SetOnInitialized(member *Signal) {
    self.Set("onInitialized", member)
}



// Add a device-ready handler and ensure the device ready sequence is started.
// 
// Phaser.Device will _not_ activate or initialize until at least one `whenReady` handler is added,
// which is normally done automatically be calling `new Phaser.Game(..)`.
// 
// The handler is invoked when the device is considered "ready", which may be immediately
// if the device is already "ready". See {@link Phaser.Device#deviceReadyAt deviceReadyAt}.
func (self *Device) WhenReadyI(args ...interface{}) {
    self.Call("whenReady", args)
}

// Internal method used for checking when the device is ready.
// This function is removed from Phaser.Device when the device becomes ready.
func (self *Device) _readyCheckI(args ...interface{}) {
    self.Call("_readyCheck", args)
}

// Internal method to initialize the capability checks.
// This function is removed from Phaser.Device once the device is initialized.
func (self *Device) _initializeI(args ...interface{}) {
    self.Call("_initialize", args)
}

// Check whether the host environment can play audio.
func (self *Device) CanPlayAudioI(args ...interface{}) bool{
    return self.Call("canPlayAudio", args).Bool()
}

// Check whether the host environment can play video files.
func (self *Device) CanPlayVideoI(args ...interface{}) bool{
    return self.Call("canPlayVideo", args).Bool()
}

// Check whether the console is open.
// Note that this only works in Firefox with Firebug and earlier versions of Chrome.
// It used to work in Chrome, but then they removed the ability: {@link http://src.chromium.org/viewvc/blink?view=revision&revision=151136}
func (self *Device) IsConsoleOpenI(args ...interface{}) {
    self.Call("isConsoleOpen", args)
}

// Detect if the host is a an Android Stock browser.
// This is available before the device "ready" event.
// 
// Authors might want to scale down on effects and switch to the CANVAS rendering method on those devices.
func (self *Device) IsAndroidStockBrowserI(args ...interface{}) {
    self.Call("isAndroidStockBrowser", args)
}
