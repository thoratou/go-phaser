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
func (self *Device) GetDeviceReadyAtA() int{
    return self.Object.Get("deviceReadyAt").Int()
}

// The time the device became ready.
func (self *Device) SetDeviceReadyAtA(member int) {
    self.Object.Set("deviceReadyAt", member)
}

// The time as which initialization has completed.
func (self *Device) GetInitializedA() bool{
    return self.Object.Get("initialized").Bool()
}

// The time as which initialization has completed.
func (self *Device) SetInitializedA(member bool) {
    self.Object.Set("initialized", member)
}

// Is running on a desktop?
func (self *Device) GetDesktopA() bool{
    return self.Object.Get("desktop").Bool()
}

// Is running on a desktop?
func (self *Device) SetDesktopA(member bool) {
    self.Object.Set("desktop", member)
}

// Is running on iOS?
func (self *Device) GetIOSA() bool{
    return self.Object.Get("iOS").Bool()
}

// Is running on iOS?
func (self *Device) SetIOSA(member bool) {
    self.Object.Set("iOS", member)
}

// If running in iOS this will contain the major version number.
func (self *Device) GetIOSVersionA() int{
    return self.Object.Get("iOSVersion").Int()
}

// If running in iOS this will contain the major version number.
func (self *Device) SetIOSVersionA(member int) {
    self.Object.Set("iOSVersion", member)
}

// Is the game running under CocoonJS?
func (self *Device) GetCocoonJSA() bool{
    return self.Object.Get("cocoonJS").Bool()
}

// Is the game running under CocoonJS?
func (self *Device) SetCocoonJSA(member bool) {
    self.Object.Set("cocoonJS", member)
}

// Is this game running with CocoonJS.App?
func (self *Device) GetCocoonJSAppA() bool{
    return self.Object.Get("cocoonJSApp").Bool()
}

// Is this game running with CocoonJS.App?
func (self *Device) SetCocoonJSAppA(member bool) {
    self.Object.Set("cocoonJSApp", member)
}

// Is the game running under Apache Cordova?
func (self *Device) GetCordovaA() bool{
    return self.Object.Get("cordova").Bool()
}

// Is the game running under Apache Cordova?
func (self *Device) SetCordovaA(member bool) {
    self.Object.Set("cordova", member)
}

// Is the game running under Node.js?
func (self *Device) GetNodeA() bool{
    return self.Object.Get("node").Bool()
}

// Is the game running under Node.js?
func (self *Device) SetNodeA(member bool) {
    self.Object.Set("node", member)
}

// Is the game running under Node-Webkit?
func (self *Device) GetNodeWebkitA() bool{
    return self.Object.Get("nodeWebkit").Bool()
}

// Is the game running under Node-Webkit?
func (self *Device) SetNodeWebkitA(member bool) {
    self.Object.Set("nodeWebkit", member)
}

// Is the game running under GitHub Electron?
func (self *Device) GetElectronA() bool{
    return self.Object.Get("electron").Bool()
}

// Is the game running under GitHub Electron?
func (self *Device) SetElectronA(member bool) {
    self.Object.Set("electron", member)
}

// Is the game running under Ejecta?
func (self *Device) GetEjectaA() bool{
    return self.Object.Get("ejecta").Bool()
}

// Is the game running under Ejecta?
func (self *Device) SetEjectaA(member bool) {
    self.Object.Set("ejecta", member)
}

// Is the game running under the Intel Crosswalk XDK?
func (self *Device) GetCrosswalkA() bool{
    return self.Object.Get("crosswalk").Bool()
}

// Is the game running under the Intel Crosswalk XDK?
func (self *Device) SetCrosswalkA(member bool) {
    self.Object.Set("crosswalk", member)
}

// Is running on android?
func (self *Device) GetAndroidA() bool{
    return self.Object.Get("android").Bool()
}

// Is running on android?
func (self *Device) SetAndroidA(member bool) {
    self.Object.Set("android", member)
}

// Is running on chromeOS?
func (self *Device) GetChromeOSA() bool{
    return self.Object.Get("chromeOS").Bool()
}

// Is running on chromeOS?
func (self *Device) SetChromeOSA(member bool) {
    self.Object.Set("chromeOS", member)
}

// Is running on linux?
func (self *Device) GetLinuxA() bool{
    return self.Object.Get("linux").Bool()
}

// Is running on linux?
func (self *Device) SetLinuxA(member bool) {
    self.Object.Set("linux", member)
}

// Is running on macOS?
func (self *Device) GetMacOSA() bool{
    return self.Object.Get("macOS").Bool()
}

// Is running on macOS?
func (self *Device) SetMacOSA(member bool) {
    self.Object.Set("macOS", member)
}

// Is running on windows?
func (self *Device) GetWindowsA() bool{
    return self.Object.Get("windows").Bool()
}

// Is running on windows?
func (self *Device) SetWindowsA(member bool) {
    self.Object.Set("windows", member)
}

// Is running on a Windows Phone?
func (self *Device) GetWindowsPhoneA() bool{
    return self.Object.Get("windowsPhone").Bool()
}

// Is running on a Windows Phone?
func (self *Device) SetWindowsPhoneA(member bool) {
    self.Object.Set("windowsPhone", member)
}

// Is canvas available?
func (self *Device) GetCanvasA() bool{
    return self.Object.Get("canvas").Bool()
}

// Is canvas available?
func (self *Device) SetCanvasA(member bool) {
    self.Object.Set("canvas", member)
}

// True if canvas supports a 'copy' bitblt onto itself when the source and destination regions overlap.
func (self *Device) GetCanvasBitBltShiftA() bool{
    return self.Object.Get("canvasBitBltShift").Bool()
}

// True if canvas supports a 'copy' bitblt onto itself when the source and destination regions overlap.
func (self *Device) SetCanvasBitBltShiftA(member bool) {
    self.Object.Set("canvasBitBltShift", member)
}

// Is webGL available?
func (self *Device) GetWebGLA() bool{
    return self.Object.Get("webGL").Bool()
}

// Is webGL available?
func (self *Device) SetWebGLA(member bool) {
    self.Object.Set("webGL", member)
}

// Is file available?
func (self *Device) GetFileA() bool{
    return self.Object.Get("file").Bool()
}

// Is file available?
func (self *Device) SetFileA(member bool) {
    self.Object.Set("file", member)
}

// Is fileSystem available?
func (self *Device) GetFileSystemA() bool{
    return self.Object.Get("fileSystem").Bool()
}

// Is fileSystem available?
func (self *Device) SetFileSystemA(member bool) {
    self.Object.Set("fileSystem", member)
}

// Is localStorage available?
func (self *Device) GetLocalStorageA() bool{
    return self.Object.Get("localStorage").Bool()
}

// Is localStorage available?
func (self *Device) SetLocalStorageA(member bool) {
    self.Object.Set("localStorage", member)
}

// Is worker available?
func (self *Device) GetWorkerA() bool{
    return self.Object.Get("worker").Bool()
}

// Is worker available?
func (self *Device) SetWorkerA(member bool) {
    self.Object.Set("worker", member)
}

// Is css3D available?
func (self *Device) GetCss3DA() bool{
    return self.Object.Get("css3D").Bool()
}

// Is css3D available?
func (self *Device) SetCss3DA(member bool) {
    self.Object.Set("css3D", member)
}

// Is Pointer Lock available?
func (self *Device) GetPointerLockA() bool{
    return self.Object.Get("pointerLock").Bool()
}

// Is Pointer Lock available?
func (self *Device) SetPointerLockA(member bool) {
    self.Object.Set("pointerLock", member)
}

// Does the browser support TypedArrays?
func (self *Device) GetTypedArrayA() bool{
    return self.Object.Get("typedArray").Bool()
}

// Does the browser support TypedArrays?
func (self *Device) SetTypedArrayA(member bool) {
    self.Object.Set("typedArray", member)
}

// Does the device support the Vibration API?
func (self *Device) GetVibrationA() bool{
    return self.Object.Get("vibration").Bool()
}

// Does the device support the Vibration API?
func (self *Device) SetVibrationA(member bool) {
    self.Object.Set("vibration", member)
}

// Does the device support the getUserMedia API?
func (self *Device) GetGetUserMediaA() bool{
    return self.Object.Get("getUserMedia").Bool()
}

// Does the device support the getUserMedia API?
func (self *Device) SetGetUserMediaA(member bool) {
    self.Object.Set("getUserMedia", member)
}

// Is the browser running in strict mode (false) or quirks mode? (true)
func (self *Device) GetQuirksModeA() bool{
    return self.Object.Get("quirksMode").Bool()
}

// Is the browser running in strict mode (false) or quirks mode? (true)
func (self *Device) SetQuirksModeA(member bool) {
    self.Object.Set("quirksMode", member)
}

// Is touch available?
func (self *Device) GetTouchA() bool{
    return self.Object.Get("touch").Bool()
}

// Is touch available?
func (self *Device) SetTouchA(member bool) {
    self.Object.Set("touch", member)
}

// Is mspointer available?
func (self *Device) GetMspointerA() bool{
    return self.Object.Get("mspointer").Bool()
}

// Is mspointer available?
func (self *Device) SetMspointerA(member bool) {
    self.Object.Set("mspointer", member)
}

// The newest type of Wheel/Scroll event supported: 'wheel', 'mousewheel', 'DOMMouseScroll'
func (self *Device) GetWheelEventA() interface{}{
    return self.Object.Get("wheelEvent")
}

// The newest type of Wheel/Scroll event supported: 'wheel', 'mousewheel', 'DOMMouseScroll'
func (self *Device) SetWheelEventA(member interface{}) {
    self.Object.Set("wheelEvent", member)
}

// Set to true if running in Arora.
func (self *Device) GetAroraA() bool{
    return self.Object.Get("arora").Bool()
}

// Set to true if running in Arora.
func (self *Device) SetAroraA(member bool) {
    self.Object.Set("arora", member)
}

// Set to true if running in Chrome.
func (self *Device) GetChromeA() bool{
    return self.Object.Get("chrome").Bool()
}

// Set to true if running in Chrome.
func (self *Device) SetChromeA(member bool) {
    self.Object.Set("chrome", member)
}

// If running in Chrome this will contain the major version number.
func (self *Device) GetChromeVersionA() int{
    return self.Object.Get("chromeVersion").Int()
}

// If running in Chrome this will contain the major version number.
func (self *Device) SetChromeVersionA(member int) {
    self.Object.Set("chromeVersion", member)
}

// Set to true if running in Epiphany.
func (self *Device) GetEpiphanyA() bool{
    return self.Object.Get("epiphany").Bool()
}

// Set to true if running in Epiphany.
func (self *Device) SetEpiphanyA(member bool) {
    self.Object.Set("epiphany", member)
}

// Set to true if running in Firefox.
func (self *Device) GetFirefoxA() bool{
    return self.Object.Get("firefox").Bool()
}

// Set to true if running in Firefox.
func (self *Device) SetFirefoxA(member bool) {
    self.Object.Set("firefox", member)
}

// If running in Firefox this will contain the major version number.
func (self *Device) GetFirefoxVersionA() int{
    return self.Object.Get("firefoxVersion").Int()
}

// If running in Firefox this will contain the major version number.
func (self *Device) SetFirefoxVersionA(member int) {
    self.Object.Set("firefoxVersion", member)
}

// Set to true if running in Internet Explorer.
func (self *Device) GetIeA() bool{
    return self.Object.Get("ie").Bool()
}

// Set to true if running in Internet Explorer.
func (self *Device) SetIeA(member bool) {
    self.Object.Set("ie", member)
}

// If running in Internet Explorer this will contain the major version number. Beyond IE10 you should use Device.trident and Device.tridentVersion.
func (self *Device) GetIeVersionA() int{
    return self.Object.Get("ieVersion").Int()
}

// If running in Internet Explorer this will contain the major version number. Beyond IE10 you should use Device.trident and Device.tridentVersion.
func (self *Device) SetIeVersionA(member int) {
    self.Object.Set("ieVersion", member)
}

// Set to true if running a Trident version of Internet Explorer (IE11+)
func (self *Device) GetTridentA() bool{
    return self.Object.Get("trident").Bool()
}

// Set to true if running a Trident version of Internet Explorer (IE11+)
func (self *Device) SetTridentA(member bool) {
    self.Object.Set("trident", member)
}

// If running in Internet Explorer 11 this will contain the major version number. See {@link http://msdn.microsoft.com/en-us/library/ie/ms537503(v=vs.85).aspx}
func (self *Device) GetTridentVersionA() int{
    return self.Object.Get("tridentVersion").Int()
}

// If running in Internet Explorer 11 this will contain the major version number. See {@link http://msdn.microsoft.com/en-us/library/ie/ms537503(v=vs.85).aspx}
func (self *Device) SetTridentVersionA(member int) {
    self.Object.Set("tridentVersion", member)
}

// Set to true if running in Microsoft Edge browser.
func (self *Device) GetEdgeA() bool{
    return self.Object.Get("edge").Bool()
}

// Set to true if running in Microsoft Edge browser.
func (self *Device) SetEdgeA(member bool) {
    self.Object.Set("edge", member)
}

// Set to true if running in Mobile Safari.
func (self *Device) GetMobileSafariA() bool{
    return self.Object.Get("mobileSafari").Bool()
}

// Set to true if running in Mobile Safari.
func (self *Device) SetMobileSafariA(member bool) {
    self.Object.Set("mobileSafari", member)
}

// Set to true if running in Midori.
func (self *Device) GetMidoriA() bool{
    return self.Object.Get("midori").Bool()
}

// Set to true if running in Midori.
func (self *Device) SetMidoriA(member bool) {
    self.Object.Set("midori", member)
}

// Set to true if running in Opera.
func (self *Device) GetOperaA() bool{
    return self.Object.Get("opera").Bool()
}

// Set to true if running in Opera.
func (self *Device) SetOperaA(member bool) {
    self.Object.Set("opera", member)
}

// Set to true if running in Safari.
func (self *Device) GetSafariA() bool{
    return self.Object.Get("safari").Bool()
}

// Set to true if running in Safari.
func (self *Device) SetSafariA(member bool) {
    self.Object.Set("safari", member)
}

// If running in Safari this will contain the major version number.
func (self *Device) GetSafariVersionA() int{
    return self.Object.Get("safariVersion").Int()
}

// If running in Safari this will contain the major version number.
func (self *Device) SetSafariVersionA(member int) {
    self.Object.Set("safariVersion", member)
}

// Set to true if running as a WebApp, i.e. within a WebView
func (self *Device) GetWebAppA() bool{
    return self.Object.Get("webApp").Bool()
}

// Set to true if running as a WebApp, i.e. within a WebView
func (self *Device) SetWebAppA(member bool) {
    self.Object.Set("webApp", member)
}

// Set to true if running in the Silk browser (as used on the Amazon Kindle)
func (self *Device) GetSilkA() bool{
    return self.Object.Get("silk").Bool()
}

// Set to true if running in the Silk browser (as used on the Amazon Kindle)
func (self *Device) SetSilkA(member bool) {
    self.Object.Set("silk", member)
}

// Are Audio tags available?
func (self *Device) GetAudioDataA() bool{
    return self.Object.Get("audioData").Bool()
}

// Are Audio tags available?
func (self *Device) SetAudioDataA(member bool) {
    self.Object.Set("audioData", member)
}

// Is the WebAudio API available?
func (self *Device) GetWebAudioA() bool{
    return self.Object.Get("webAudio").Bool()
}

// Is the WebAudio API available?
func (self *Device) SetWebAudioA(member bool) {
    self.Object.Set("webAudio", member)
}

// Can this device play ogg files?
func (self *Device) GetOggA() bool{
    return self.Object.Get("ogg").Bool()
}

// Can this device play ogg files?
func (self *Device) SetOggA(member bool) {
    self.Object.Set("ogg", member)
}

// Can this device play opus files?
func (self *Device) GetOpusA() bool{
    return self.Object.Get("opus").Bool()
}

// Can this device play opus files?
func (self *Device) SetOpusA(member bool) {
    self.Object.Set("opus", member)
}

// Can this device play mp3 files?
func (self *Device) GetMp3A() bool{
    return self.Object.Get("mp3").Bool()
}

// Can this device play mp3 files?
func (self *Device) SetMp3A(member bool) {
    self.Object.Set("mp3", member)
}

// Can this device play wav files?
func (self *Device) GetWavA() bool{
    return self.Object.Get("wav").Bool()
}

// Can this device play wav files?
func (self *Device) SetWavA(member bool) {
    self.Object.Set("wav", member)
}

// Can this device play m4a files? True if this device can play m4a files.
func (self *Device) GetM4aA() bool{
    return self.Object.Get("m4a").Bool()
}

// Can this device play m4a files? True if this device can play m4a files.
func (self *Device) SetM4aA(member bool) {
    self.Object.Set("m4a", member)
}

// Can this device play webm files?
func (self *Device) GetWebmA() bool{
    return self.Object.Get("webm").Bool()
}

// Can this device play webm files?
func (self *Device) SetWebmA(member bool) {
    self.Object.Set("webm", member)
}

// Can this device play EC-3 Dolby Digital Plus files?
func (self *Device) GetDolbyA() bool{
    return self.Object.Get("dolby").Bool()
}

// Can this device play EC-3 Dolby Digital Plus files?
func (self *Device) SetDolbyA(member bool) {
    self.Object.Set("dolby", member)
}

// Can this device play ogg video files?
func (self *Device) GetOggVideoA() bool{
    return self.Object.Get("oggVideo").Bool()
}

// Can this device play ogg video files?
func (self *Device) SetOggVideoA(member bool) {
    self.Object.Set("oggVideo", member)
}

// Can this device play h264 mp4 video files?
func (self *Device) GetH264VideoA() bool{
    return self.Object.Get("h264Video").Bool()
}

// Can this device play h264 mp4 video files?
func (self *Device) SetH264VideoA(member bool) {
    self.Object.Set("h264Video", member)
}

// Can this device play h264 mp4 video files?
func (self *Device) GetMp4VideoA() bool{
    return self.Object.Get("mp4Video").Bool()
}

// Can this device play h264 mp4 video files?
func (self *Device) SetMp4VideoA(member bool) {
    self.Object.Set("mp4Video", member)
}

// Can this device play webm video files?
func (self *Device) GetWebmVideoA() bool{
    return self.Object.Get("webmVideo").Bool()
}

// Can this device play webm video files?
func (self *Device) SetWebmVideoA(member bool) {
    self.Object.Set("webmVideo", member)
}

// Can this device play vp9 video files?
func (self *Device) GetVp9VideoA() bool{
    return self.Object.Get("vp9Video").Bool()
}

// Can this device play vp9 video files?
func (self *Device) SetVp9VideoA(member bool) {
    self.Object.Set("vp9Video", member)
}

// Can this device play hls video files?
func (self *Device) GetHlsVideoA() bool{
    return self.Object.Get("hlsVideo").Bool()
}

// Can this device play hls video files?
func (self *Device) SetHlsVideoA(member bool) {
    self.Object.Set("hlsVideo", member)
}

// Is running on iPhone?
func (self *Device) GetIPhoneA() bool{
    return self.Object.Get("iPhone").Bool()
}

// Is running on iPhone?
func (self *Device) SetIPhoneA(member bool) {
    self.Object.Set("iPhone", member)
}

// Is running on iPhone4?
func (self *Device) GetIPhone4A() bool{
    return self.Object.Get("iPhone4").Bool()
}

// Is running on iPhone4?
func (self *Device) SetIPhone4A(member bool) {
    self.Object.Set("iPhone4", member)
}

// Is running on iPad?
func (self *Device) GetIPadA() bool{
    return self.Object.Get("iPad").Bool()
}

// Is running on iPad?
func (self *Device) SetIPadA(member bool) {
    self.Object.Set("iPad", member)
}

// PixelRatio of the host device?
func (self *Device) GetPixelRatioA() int{
    return self.Object.Get("pixelRatio").Int()
}

// PixelRatio of the host device?
func (self *Device) SetPixelRatioA(member int) {
    self.Object.Set("pixelRatio", member)
}

// Is the device big or little endian? (only detected if the browser supports TypedArrays)
func (self *Device) GetLittleEndianA() bool{
    return self.Object.Get("littleEndian").Bool()
}

// Is the device big or little endian? (only detected if the browser supports TypedArrays)
func (self *Device) SetLittleEndianA(member bool) {
    self.Object.Set("littleEndian", member)
}

// Same value as `littleEndian`.
func (self *Device) GetLITTLE_ENDIANA() bool{
    return self.Object.Get("LITTLE_ENDIAN").Bool()
}

// Same value as `littleEndian`.
func (self *Device) SetLITTLE_ENDIANA(member bool) {
    self.Object.Set("LITTLE_ENDIAN", member)
}

// Does the device context support 32bit pixel manipulation using array buffer views?
func (self *Device) GetSupport32bitA() bool{
    return self.Object.Get("support32bit").Bool()
}

// Does the device context support 32bit pixel manipulation using array buffer views?
func (self *Device) SetSupport32bitA(member bool) {
    self.Object.Set("support32bit", member)
}

// Does the browser support the Full Screen API?
func (self *Device) GetFullscreenA() bool{
    return self.Object.Get("fullscreen").Bool()
}

// Does the browser support the Full Screen API?
func (self *Device) SetFullscreenA(member bool) {
    self.Object.Set("fullscreen", member)
}

// If the browser supports the Full Screen API this holds the call you need to use to activate it.
func (self *Device) GetRequestFullscreenA() string{
    return self.Object.Get("requestFullscreen").String()
}

// If the browser supports the Full Screen API this holds the call you need to use to activate it.
func (self *Device) SetRequestFullscreenA(member string) {
    self.Object.Set("requestFullscreen", member)
}

// If the browser supports the Full Screen API this holds the call you need to use to cancel it.
func (self *Device) GetCancelFullscreenA() string{
    return self.Object.Get("cancelFullscreen").String()
}

// If the browser supports the Full Screen API this holds the call you need to use to cancel it.
func (self *Device) SetCancelFullscreenA(member string) {
    self.Object.Set("cancelFullscreen", member)
}

// Does the browser support access to the Keyboard during Full Screen mode?
func (self *Device) GetFullscreenKeyboardA() bool{
    return self.Object.Get("fullscreenKeyboard").Bool()
}

// Does the browser support access to the Keyboard during Full Screen mode?
func (self *Device) SetFullscreenKeyboardA(member bool) {
    self.Object.Set("fullscreenKeyboard", member)
}

// This signal is dispatched after device initialization occurs but before any of the ready
// callbacks (see {@link Phaser.Device.whenReady whenReady}) have been invoked.
// 
// Local "patching" for a particular device can/should be done in this event.
// 
// _Note_: This signal is removed after the device has been readied; if a handler has not been
// added _before_ `new Phaser.Game(..)` it is probably too late.
func (self *Device) GetOnInitializedA() *Signal{
    return &Signal{self.Object.Get("onInitialized")}
}

// This signal is dispatched after device initialization occurs but before any of the ready
// callbacks (see {@link Phaser.Device.whenReady whenReady}) have been invoked.
// 
// Local "patching" for a particular device can/should be done in this event.
// 
// _Note_: This signal is removed after the device has been readied; if a handler has not been
// added _before_ `new Phaser.Game(..)` it is probably too late.
func (self *Device) SetOnInitializedA(member *Signal) {
    self.Object.Set("onInitialized", member)
}



// Add a device-ready handler and ensure the device ready sequence is started.
// 
// Phaser.Device will _not_ activate or initialize until at least one `whenReady` handler is added,
// which is normally done automatically be calling `new Phaser.Game(..)`.
// 
// The handler is invoked when the device is considered "ready", which may be immediately
// if the device is already "ready". See {@link Phaser.Device#deviceReadyAt deviceReadyAt}.
func (self *Device) WhenReady(handler func(...interface{})) {
    self.Object.Call("whenReady", handler)
}

// Add a device-ready handler and ensure the device ready sequence is started.
// 
// Phaser.Device will _not_ activate or initialize until at least one `whenReady` handler is added,
// which is normally done automatically be calling `new Phaser.Game(..)`.
// 
// The handler is invoked when the device is considered "ready", which may be immediately
// if the device is already "ready". See {@link Phaser.Device#deviceReadyAt deviceReadyAt}.
func (self *Device) WhenReady1O(handler func(...interface{}), context interface{}) {
    self.Object.Call("whenReady", handler, context)
}

// Add a device-ready handler and ensure the device ready sequence is started.
// 
// Phaser.Device will _not_ activate or initialize until at least one `whenReady` handler is added,
// which is normally done automatically be calling `new Phaser.Game(..)`.
// 
// The handler is invoked when the device is considered "ready", which may be immediately
// if the device is already "ready". See {@link Phaser.Device#deviceReadyAt deviceReadyAt}.
func (self *Device) WhenReady2O(handler func(...interface{}), context interface{}, nonPrimer bool) {
    self.Object.Call("whenReady", handler, context, nonPrimer)
}

// Add a device-ready handler and ensure the device ready sequence is started.
// 
// Phaser.Device will _not_ activate or initialize until at least one `whenReady` handler is added,
// which is normally done automatically be calling `new Phaser.Game(..)`.
// 
// The handler is invoked when the device is considered "ready", which may be immediately
// if the device is already "ready". See {@link Phaser.Device#deviceReadyAt deviceReadyAt}.
func (self *Device) WhenReadyI(args ...interface{}) {
    self.Object.Call("whenReady", args)
}

// Internal method used for checking when the device is ready.
// This function is removed from Phaser.Device when the device becomes ready.
func (self *Device) _readyCheck() {
    self.Object.Call("_readyCheck")
}

// Internal method used for checking when the device is ready.
// This function is removed from Phaser.Device when the device becomes ready.
func (self *Device) _readyCheckI(args ...interface{}) {
    self.Object.Call("_readyCheck", args)
}

// Internal method to initialize the capability checks.
// This function is removed from Phaser.Device once the device is initialized.
func (self *Device) _initialize() {
    self.Object.Call("_initialize")
}

// Internal method to initialize the capability checks.
// This function is removed from Phaser.Device once the device is initialized.
func (self *Device) _initializeI(args ...interface{}) {
    self.Object.Call("_initialize", args)
}

// Check whether the host environment can play audio.
func (self *Device) CanPlayAudio(type_ string) bool{
    return self.Object.Call("canPlayAudio", type_).Bool()
}

// Check whether the host environment can play audio.
func (self *Device) CanPlayAudioI(args ...interface{}) bool{
    return self.Object.Call("canPlayAudio", args).Bool()
}

// Check whether the host environment can play video files.
func (self *Device) CanPlayVideo(type_ string) bool{
    return self.Object.Call("canPlayVideo", type_).Bool()
}

// Check whether the host environment can play video files.
func (self *Device) CanPlayVideoI(args ...interface{}) bool{
    return self.Object.Call("canPlayVideo", args).Bool()
}

// Check whether the console is open.
// Note that this only works in Firefox with Firebug and earlier versions of Chrome.
// It used to work in Chrome, but then they removed the ability: {@link http://src.chromium.org/viewvc/blink?view=revision&revision=151136}
func (self *Device) IsConsoleOpen() {
    self.Object.Call("isConsoleOpen")
}

// Check whether the console is open.
// Note that this only works in Firefox with Firebug and earlier versions of Chrome.
// It used to work in Chrome, but then they removed the ability: {@link http://src.chromium.org/viewvc/blink?view=revision&revision=151136}
func (self *Device) IsConsoleOpenI(args ...interface{}) {
    self.Object.Call("isConsoleOpen", args)
}

// Detect if the host is a an Android Stock browser.
// This is available before the device "ready" event.
// 
// Authors might want to scale down on effects and switch to the CANVAS rendering method on those devices.
func (self *Device) IsAndroidStockBrowser() {
    self.Object.Call("isAndroidStockBrowser")
}

// Detect if the host is a an Android Stock browser.
// This is available before the device "ready" event.
// 
// Authors might want to scale down on effects and switch to the CANVAS rendering method on those devices.
func (self *Device) IsAndroidStockBrowserI(args ...interface{}) {
    self.Object.Call("isAndroidStockBrowser", args)
}
