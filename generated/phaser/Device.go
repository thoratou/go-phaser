// Package phaser Automatic generation for Phaser.Device
// generated file Device.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Device Detects device support capabilities and is responsible for device initialization - see {@link Phaser.Device.whenReady whenReady}.
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

// NewDevice It is not possible to instantiate the Device class manually.
func NewDevice() *Device {
    return &Device{js.Global.Get("Phaser").Get("Device").New()}
}
// NewDeviceI It is not possible to instantiate the Device class manually.
func NewDeviceI(args ...interface{}) *Device {
    return &Device{js.Global.Get("Phaser").Get("Device").New(args)}
}



// Device Binding conversion method to Device point 
func ToDevice(jsStruct interface{}) *Device {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Device{Object: object}
	}
	return nil
}



// DeviceReadyAt The time the device became ready.
func (self *Device) DeviceReadyAt() int{
    return self.Object.Get("deviceReadyAt").Int()
}

// SetDeviceReadyAtA The time the device became ready.
func (self *Device) SetDeviceReadyAtA(member int) {
    self.Object.Set("deviceReadyAt", member)
}

// Initialized The time as which initialization has completed.
func (self *Device) Initialized() bool{
    return self.Object.Get("initialized").Bool()
}

// SetInitializedA The time as which initialization has completed.
func (self *Device) SetInitializedA(member bool) {
    self.Object.Set("initialized", member)
}

// Desktop Is running on a desktop?
func (self *Device) Desktop() bool{
    return self.Object.Get("desktop").Bool()
}

// SetDesktopA Is running on a desktop?
func (self *Device) SetDesktopA(member bool) {
    self.Object.Set("desktop", member)
}

// IOS Is running on iOS?
func (self *Device) IOS() bool{
    return self.Object.Get("iOS").Bool()
}

// SetIOSA Is running on iOS?
func (self *Device) SetIOSA(member bool) {
    self.Object.Set("iOS", member)
}

// IOSVersion If running in iOS this will contain the major version number.
func (self *Device) IOSVersion() int{
    return self.Object.Get("iOSVersion").Int()
}

// SetIOSVersionA If running in iOS this will contain the major version number.
func (self *Device) SetIOSVersionA(member int) {
    self.Object.Set("iOSVersion", member)
}

// CocoonJS Is the game running under CocoonJS?
func (self *Device) CocoonJS() bool{
    return self.Object.Get("cocoonJS").Bool()
}

// SetCocoonJSA Is the game running under CocoonJS?
func (self *Device) SetCocoonJSA(member bool) {
    self.Object.Set("cocoonJS", member)
}

// CocoonJSApp Is this game running with CocoonJS.App?
func (self *Device) CocoonJSApp() bool{
    return self.Object.Get("cocoonJSApp").Bool()
}

// SetCocoonJSAppA Is this game running with CocoonJS.App?
func (self *Device) SetCocoonJSAppA(member bool) {
    self.Object.Set("cocoonJSApp", member)
}

// Cordova Is the game running under Apache Cordova?
func (self *Device) Cordova() bool{
    return self.Object.Get("cordova").Bool()
}

// SetCordovaA Is the game running under Apache Cordova?
func (self *Device) SetCordovaA(member bool) {
    self.Object.Set("cordova", member)
}

// Node Is the game running under Node.js?
func (self *Device) Node() bool{
    return self.Object.Get("node").Bool()
}

// SetNodeA Is the game running under Node.js?
func (self *Device) SetNodeA(member bool) {
    self.Object.Set("node", member)
}

// NodeWebkit Is the game running under Node-Webkit?
func (self *Device) NodeWebkit() bool{
    return self.Object.Get("nodeWebkit").Bool()
}

// SetNodeWebkitA Is the game running under Node-Webkit?
func (self *Device) SetNodeWebkitA(member bool) {
    self.Object.Set("nodeWebkit", member)
}

// Electron Is the game running under GitHub Electron?
func (self *Device) Electron() bool{
    return self.Object.Get("electron").Bool()
}

// SetElectronA Is the game running under GitHub Electron?
func (self *Device) SetElectronA(member bool) {
    self.Object.Set("electron", member)
}

// Ejecta Is the game running under Ejecta?
func (self *Device) Ejecta() bool{
    return self.Object.Get("ejecta").Bool()
}

// SetEjectaA Is the game running under Ejecta?
func (self *Device) SetEjectaA(member bool) {
    self.Object.Set("ejecta", member)
}

// Crosswalk Is the game running under the Intel Crosswalk XDK?
func (self *Device) Crosswalk() bool{
    return self.Object.Get("crosswalk").Bool()
}

// SetCrosswalkA Is the game running under the Intel Crosswalk XDK?
func (self *Device) SetCrosswalkA(member bool) {
    self.Object.Set("crosswalk", member)
}

// Android Is running on android?
func (self *Device) Android() bool{
    return self.Object.Get("android").Bool()
}

// SetAndroidA Is running on android?
func (self *Device) SetAndroidA(member bool) {
    self.Object.Set("android", member)
}

// ChromeOS Is running on chromeOS?
func (self *Device) ChromeOS() bool{
    return self.Object.Get("chromeOS").Bool()
}

// SetChromeOSA Is running on chromeOS?
func (self *Device) SetChromeOSA(member bool) {
    self.Object.Set("chromeOS", member)
}

// Linux Is running on linux?
func (self *Device) Linux() bool{
    return self.Object.Get("linux").Bool()
}

// SetLinuxA Is running on linux?
func (self *Device) SetLinuxA(member bool) {
    self.Object.Set("linux", member)
}

// MacOS Is running on macOS?
func (self *Device) MacOS() bool{
    return self.Object.Get("macOS").Bool()
}

// SetMacOSA Is running on macOS?
func (self *Device) SetMacOSA(member bool) {
    self.Object.Set("macOS", member)
}

// Windows Is running on windows?
func (self *Device) Windows() bool{
    return self.Object.Get("windows").Bool()
}

// SetWindowsA Is running on windows?
func (self *Device) SetWindowsA(member bool) {
    self.Object.Set("windows", member)
}

// WindowsPhone Is running on a Windows Phone?
func (self *Device) WindowsPhone() bool{
    return self.Object.Get("windowsPhone").Bool()
}

// SetWindowsPhoneA Is running on a Windows Phone?
func (self *Device) SetWindowsPhoneA(member bool) {
    self.Object.Set("windowsPhone", member)
}

// Canvas Is canvas available?
func (self *Device) Canvas() bool{
    return self.Object.Get("canvas").Bool()
}

// SetCanvasA Is canvas available?
func (self *Device) SetCanvasA(member bool) {
    self.Object.Set("canvas", member)
}

// CanvasBitBltShift True if canvas supports a 'copy' bitblt onto itself when the source and destination regions overlap.
func (self *Device) CanvasBitBltShift() bool{
    return self.Object.Get("canvasBitBltShift").Bool()
}

// SetCanvasBitBltShiftA True if canvas supports a 'copy' bitblt onto itself when the source and destination regions overlap.
func (self *Device) SetCanvasBitBltShiftA(member bool) {
    self.Object.Set("canvasBitBltShift", member)
}

// WebGL Is webGL available?
func (self *Device) WebGL() bool{
    return self.Object.Get("webGL").Bool()
}

// SetWebGLA Is webGL available?
func (self *Device) SetWebGLA(member bool) {
    self.Object.Set("webGL", member)
}

// File Is file available?
func (self *Device) File() bool{
    return self.Object.Get("file").Bool()
}

// SetFileA Is file available?
func (self *Device) SetFileA(member bool) {
    self.Object.Set("file", member)
}

// FileSystem Is fileSystem available?
func (self *Device) FileSystem() bool{
    return self.Object.Get("fileSystem").Bool()
}

// SetFileSystemA Is fileSystem available?
func (self *Device) SetFileSystemA(member bool) {
    self.Object.Set("fileSystem", member)
}

// LocalStorage Is localStorage available?
func (self *Device) LocalStorage() bool{
    return self.Object.Get("localStorage").Bool()
}

// SetLocalStorageA Is localStorage available?
func (self *Device) SetLocalStorageA(member bool) {
    self.Object.Set("localStorage", member)
}

// Worker Is worker available?
func (self *Device) Worker() bool{
    return self.Object.Get("worker").Bool()
}

// SetWorkerA Is worker available?
func (self *Device) SetWorkerA(member bool) {
    self.Object.Set("worker", member)
}

// Css3D Is css3D available?
func (self *Device) Css3D() bool{
    return self.Object.Get("css3D").Bool()
}

// SetCss3DA Is css3D available?
func (self *Device) SetCss3DA(member bool) {
    self.Object.Set("css3D", member)
}

// PointerLock Is Pointer Lock available?
func (self *Device) PointerLock() bool{
    return self.Object.Get("pointerLock").Bool()
}

// SetPointerLockA Is Pointer Lock available?
func (self *Device) SetPointerLockA(member bool) {
    self.Object.Set("pointerLock", member)
}

// TypedArray Does the browser support TypedArrays?
func (self *Device) TypedArray() bool{
    return self.Object.Get("typedArray").Bool()
}

// SetTypedArrayA Does the browser support TypedArrays?
func (self *Device) SetTypedArrayA(member bool) {
    self.Object.Set("typedArray", member)
}

// Vibration Does the device support the Vibration API?
func (self *Device) Vibration() bool{
    return self.Object.Get("vibration").Bool()
}

// SetVibrationA Does the device support the Vibration API?
func (self *Device) SetVibrationA(member bool) {
    self.Object.Set("vibration", member)
}

// GetUserMedia Does the device support the getUserMedia API?
func (self *Device) GetUserMedia() bool{
    return self.Object.Get("getUserMedia").Bool()
}

// SetGetUserMediaA Does the device support the getUserMedia API?
func (self *Device) SetGetUserMediaA(member bool) {
    self.Object.Set("getUserMedia", member)
}

// QuirksMode Is the browser running in strict mode (false) or quirks mode? (true)
func (self *Device) QuirksMode() bool{
    return self.Object.Get("quirksMode").Bool()
}

// SetQuirksModeA Is the browser running in strict mode (false) or quirks mode? (true)
func (self *Device) SetQuirksModeA(member bool) {
    self.Object.Set("quirksMode", member)
}

// Touch Is touch available?
func (self *Device) Touch() bool{
    return self.Object.Get("touch").Bool()
}

// SetTouchA Is touch available?
func (self *Device) SetTouchA(member bool) {
    self.Object.Set("touch", member)
}

// Mspointer Is mspointer available?
func (self *Device) Mspointer() bool{
    return self.Object.Get("mspointer").Bool()
}

// SetMspointerA Is mspointer available?
func (self *Device) SetMspointerA(member bool) {
    self.Object.Set("mspointer", member)
}

// WheelEvent The newest type of Wheel/Scroll event supported: 'wheel', 'mousewheel', 'DOMMouseScroll'
func (self *Device) WheelEvent() interface{}{
    return self.Object.Get("wheelEvent")
}

// SetWheelEventA The newest type of Wheel/Scroll event supported: 'wheel', 'mousewheel', 'DOMMouseScroll'
func (self *Device) SetWheelEventA(member interface{}) {
    self.Object.Set("wheelEvent", member)
}

// Arora Set to true if running in Arora.
func (self *Device) Arora() bool{
    return self.Object.Get("arora").Bool()
}

// SetAroraA Set to true if running in Arora.
func (self *Device) SetAroraA(member bool) {
    self.Object.Set("arora", member)
}

// Chrome Set to true if running in Chrome.
func (self *Device) Chrome() bool{
    return self.Object.Get("chrome").Bool()
}

// SetChromeA Set to true if running in Chrome.
func (self *Device) SetChromeA(member bool) {
    self.Object.Set("chrome", member)
}

// ChromeVersion If running in Chrome this will contain the major version number.
func (self *Device) ChromeVersion() int{
    return self.Object.Get("chromeVersion").Int()
}

// SetChromeVersionA If running in Chrome this will contain the major version number.
func (self *Device) SetChromeVersionA(member int) {
    self.Object.Set("chromeVersion", member)
}

// Epiphany Set to true if running in Epiphany.
func (self *Device) Epiphany() bool{
    return self.Object.Get("epiphany").Bool()
}

// SetEpiphanyA Set to true if running in Epiphany.
func (self *Device) SetEpiphanyA(member bool) {
    self.Object.Set("epiphany", member)
}

// Firefox Set to true if running in Firefox.
func (self *Device) Firefox() bool{
    return self.Object.Get("firefox").Bool()
}

// SetFirefoxA Set to true if running in Firefox.
func (self *Device) SetFirefoxA(member bool) {
    self.Object.Set("firefox", member)
}

// FirefoxVersion If running in Firefox this will contain the major version number.
func (self *Device) FirefoxVersion() int{
    return self.Object.Get("firefoxVersion").Int()
}

// SetFirefoxVersionA If running in Firefox this will contain the major version number.
func (self *Device) SetFirefoxVersionA(member int) {
    self.Object.Set("firefoxVersion", member)
}

// Ie Set to true if running in Internet Explorer.
func (self *Device) Ie() bool{
    return self.Object.Get("ie").Bool()
}

// SetIeA Set to true if running in Internet Explorer.
func (self *Device) SetIeA(member bool) {
    self.Object.Set("ie", member)
}

// IeVersion If running in Internet Explorer this will contain the major version number. Beyond IE10 you should use Device.trident and Device.tridentVersion.
func (self *Device) IeVersion() int{
    return self.Object.Get("ieVersion").Int()
}

// SetIeVersionA If running in Internet Explorer this will contain the major version number. Beyond IE10 you should use Device.trident and Device.tridentVersion.
func (self *Device) SetIeVersionA(member int) {
    self.Object.Set("ieVersion", member)
}

// Trident Set to true if running a Trident version of Internet Explorer (IE11+)
func (self *Device) Trident() bool{
    return self.Object.Get("trident").Bool()
}

// SetTridentA Set to true if running a Trident version of Internet Explorer (IE11+)
func (self *Device) SetTridentA(member bool) {
    self.Object.Set("trident", member)
}

// TridentVersion If running in Internet Explorer 11 this will contain the major version number. See {@link http://msdn.microsoft.com/en-us/library/ie/ms537503(v=vs.85).aspx}
func (self *Device) TridentVersion() int{
    return self.Object.Get("tridentVersion").Int()
}

// SetTridentVersionA If running in Internet Explorer 11 this will contain the major version number. See {@link http://msdn.microsoft.com/en-us/library/ie/ms537503(v=vs.85).aspx}
func (self *Device) SetTridentVersionA(member int) {
    self.Object.Set("tridentVersion", member)
}

// Edge Set to true if running in Microsoft Edge browser.
func (self *Device) Edge() bool{
    return self.Object.Get("edge").Bool()
}

// SetEdgeA Set to true if running in Microsoft Edge browser.
func (self *Device) SetEdgeA(member bool) {
    self.Object.Set("edge", member)
}

// MobileSafari Set to true if running in Mobile Safari.
func (self *Device) MobileSafari() bool{
    return self.Object.Get("mobileSafari").Bool()
}

// SetMobileSafariA Set to true if running in Mobile Safari.
func (self *Device) SetMobileSafariA(member bool) {
    self.Object.Set("mobileSafari", member)
}

// Midori Set to true if running in Midori.
func (self *Device) Midori() bool{
    return self.Object.Get("midori").Bool()
}

// SetMidoriA Set to true if running in Midori.
func (self *Device) SetMidoriA(member bool) {
    self.Object.Set("midori", member)
}

// Opera Set to true if running in Opera.
func (self *Device) Opera() bool{
    return self.Object.Get("opera").Bool()
}

// SetOperaA Set to true if running in Opera.
func (self *Device) SetOperaA(member bool) {
    self.Object.Set("opera", member)
}

// Safari Set to true if running in Safari.
func (self *Device) Safari() bool{
    return self.Object.Get("safari").Bool()
}

// SetSafariA Set to true if running in Safari.
func (self *Device) SetSafariA(member bool) {
    self.Object.Set("safari", member)
}

// SafariVersion If running in Safari this will contain the major version number.
func (self *Device) SafariVersion() int{
    return self.Object.Get("safariVersion").Int()
}

// SetSafariVersionA If running in Safari this will contain the major version number.
func (self *Device) SetSafariVersionA(member int) {
    self.Object.Set("safariVersion", member)
}

// WebApp Set to true if running as a WebApp, i.e. within a WebView
func (self *Device) WebApp() bool{
    return self.Object.Get("webApp").Bool()
}

// SetWebAppA Set to true if running as a WebApp, i.e. within a WebView
func (self *Device) SetWebAppA(member bool) {
    self.Object.Set("webApp", member)
}

// Silk Set to true if running in the Silk browser (as used on the Amazon Kindle)
func (self *Device) Silk() bool{
    return self.Object.Get("silk").Bool()
}

// SetSilkA Set to true if running in the Silk browser (as used on the Amazon Kindle)
func (self *Device) SetSilkA(member bool) {
    self.Object.Set("silk", member)
}

// AudioData Are Audio tags available?
func (self *Device) AudioData() bool{
    return self.Object.Get("audioData").Bool()
}

// SetAudioDataA Are Audio tags available?
func (self *Device) SetAudioDataA(member bool) {
    self.Object.Set("audioData", member)
}

// WebAudio Is the WebAudio API available?
func (self *Device) WebAudio() bool{
    return self.Object.Get("webAudio").Bool()
}

// SetWebAudioA Is the WebAudio API available?
func (self *Device) SetWebAudioA(member bool) {
    self.Object.Set("webAudio", member)
}

// Ogg Can this device play ogg files?
func (self *Device) Ogg() bool{
    return self.Object.Get("ogg").Bool()
}

// SetOggA Can this device play ogg files?
func (self *Device) SetOggA(member bool) {
    self.Object.Set("ogg", member)
}

// Opus Can this device play opus files?
func (self *Device) Opus() bool{
    return self.Object.Get("opus").Bool()
}

// SetOpusA Can this device play opus files?
func (self *Device) SetOpusA(member bool) {
    self.Object.Set("opus", member)
}

// Mp3 Can this device play mp3 files?
func (self *Device) Mp3() bool{
    return self.Object.Get("mp3").Bool()
}

// SetMp3A Can this device play mp3 files?
func (self *Device) SetMp3A(member bool) {
    self.Object.Set("mp3", member)
}

// Wav Can this device play wav files?
func (self *Device) Wav() bool{
    return self.Object.Get("wav").Bool()
}

// SetWavA Can this device play wav files?
func (self *Device) SetWavA(member bool) {
    self.Object.Set("wav", member)
}

// M4a Can this device play m4a files? True if this device can play m4a files.
func (self *Device) M4a() bool{
    return self.Object.Get("m4a").Bool()
}

// SetM4aA Can this device play m4a files? True if this device can play m4a files.
func (self *Device) SetM4aA(member bool) {
    self.Object.Set("m4a", member)
}

// Webm Can this device play webm files?
func (self *Device) Webm() bool{
    return self.Object.Get("webm").Bool()
}

// SetWebmA Can this device play webm files?
func (self *Device) SetWebmA(member bool) {
    self.Object.Set("webm", member)
}

// Dolby Can this device play EC-3 Dolby Digital Plus files?
func (self *Device) Dolby() bool{
    return self.Object.Get("dolby").Bool()
}

// SetDolbyA Can this device play EC-3 Dolby Digital Plus files?
func (self *Device) SetDolbyA(member bool) {
    self.Object.Set("dolby", member)
}

// OggVideo Can this device play ogg video files?
func (self *Device) OggVideo() bool{
    return self.Object.Get("oggVideo").Bool()
}

// SetOggVideoA Can this device play ogg video files?
func (self *Device) SetOggVideoA(member bool) {
    self.Object.Set("oggVideo", member)
}

// H264Video Can this device play h264 mp4 video files?
func (self *Device) H264Video() bool{
    return self.Object.Get("h264Video").Bool()
}

// SetH264VideoA Can this device play h264 mp4 video files?
func (self *Device) SetH264VideoA(member bool) {
    self.Object.Set("h264Video", member)
}

// Mp4Video Can this device play h264 mp4 video files?
func (self *Device) Mp4Video() bool{
    return self.Object.Get("mp4Video").Bool()
}

// SetMp4VideoA Can this device play h264 mp4 video files?
func (self *Device) SetMp4VideoA(member bool) {
    self.Object.Set("mp4Video", member)
}

// WebmVideo Can this device play webm video files?
func (self *Device) WebmVideo() bool{
    return self.Object.Get("webmVideo").Bool()
}

// SetWebmVideoA Can this device play webm video files?
func (self *Device) SetWebmVideoA(member bool) {
    self.Object.Set("webmVideo", member)
}

// Vp9Video Can this device play vp9 video files?
func (self *Device) Vp9Video() bool{
    return self.Object.Get("vp9Video").Bool()
}

// SetVp9VideoA Can this device play vp9 video files?
func (self *Device) SetVp9VideoA(member bool) {
    self.Object.Set("vp9Video", member)
}

// HlsVideo Can this device play hls video files?
func (self *Device) HlsVideo() bool{
    return self.Object.Get("hlsVideo").Bool()
}

// SetHlsVideoA Can this device play hls video files?
func (self *Device) SetHlsVideoA(member bool) {
    self.Object.Set("hlsVideo", member)
}

// IPhone Is running on iPhone?
func (self *Device) IPhone() bool{
    return self.Object.Get("iPhone").Bool()
}

// SetIPhoneA Is running on iPhone?
func (self *Device) SetIPhoneA(member bool) {
    self.Object.Set("iPhone", member)
}

// IPhone4 Is running on iPhone4?
func (self *Device) IPhone4() bool{
    return self.Object.Get("iPhone4").Bool()
}

// SetIPhone4A Is running on iPhone4?
func (self *Device) SetIPhone4A(member bool) {
    self.Object.Set("iPhone4", member)
}

// IPad Is running on iPad?
func (self *Device) IPad() bool{
    return self.Object.Get("iPad").Bool()
}

// SetIPadA Is running on iPad?
func (self *Device) SetIPadA(member bool) {
    self.Object.Set("iPad", member)
}

// PixelRatio PixelRatio of the host device?
func (self *Device) PixelRatio() int{
    return self.Object.Get("pixelRatio").Int()
}

// SetPixelRatioA PixelRatio of the host device?
func (self *Device) SetPixelRatioA(member int) {
    self.Object.Set("pixelRatio", member)
}

// LittleEndian Is the device big or little endian? (only detected if the browser supports TypedArrays)
func (self *Device) LittleEndian() bool{
    return self.Object.Get("littleEndian").Bool()
}

// SetLittleEndianA Is the device big or little endian? (only detected if the browser supports TypedArrays)
func (self *Device) SetLittleEndianA(member bool) {
    self.Object.Set("littleEndian", member)
}

// LITTLE_ENDIAN Same value as `littleEndian`.
func (self *Device) LITTLE_ENDIAN() bool{
    return self.Object.Get("LITTLE_ENDIAN").Bool()
}

// SetLITTLE_ENDIANA Same value as `littleEndian`.
func (self *Device) SetLITTLE_ENDIANA(member bool) {
    self.Object.Set("LITTLE_ENDIAN", member)
}

// Support32bit Does the device context support 32bit pixel manipulation using array buffer views?
func (self *Device) Support32bit() bool{
    return self.Object.Get("support32bit").Bool()
}

// SetSupport32bitA Does the device context support 32bit pixel manipulation using array buffer views?
func (self *Device) SetSupport32bitA(member bool) {
    self.Object.Set("support32bit", member)
}

// Fullscreen Does the browser support the Full Screen API?
func (self *Device) Fullscreen() bool{
    return self.Object.Get("fullscreen").Bool()
}

// SetFullscreenA Does the browser support the Full Screen API?
func (self *Device) SetFullscreenA(member bool) {
    self.Object.Set("fullscreen", member)
}

// RequestFullscreen If the browser supports the Full Screen API this holds the call you need to use to activate it.
func (self *Device) RequestFullscreen() string{
    return self.Object.Get("requestFullscreen").String()
}

// SetRequestFullscreenA If the browser supports the Full Screen API this holds the call you need to use to activate it.
func (self *Device) SetRequestFullscreenA(member string) {
    self.Object.Set("requestFullscreen", member)
}

// CancelFullscreen If the browser supports the Full Screen API this holds the call you need to use to cancel it.
func (self *Device) CancelFullscreen() string{
    return self.Object.Get("cancelFullscreen").String()
}

// SetCancelFullscreenA If the browser supports the Full Screen API this holds the call you need to use to cancel it.
func (self *Device) SetCancelFullscreenA(member string) {
    self.Object.Set("cancelFullscreen", member)
}

// FullscreenKeyboard Does the browser support access to the Keyboard during Full Screen mode?
func (self *Device) FullscreenKeyboard() bool{
    return self.Object.Get("fullscreenKeyboard").Bool()
}

// SetFullscreenKeyboardA Does the browser support access to the Keyboard during Full Screen mode?
func (self *Device) SetFullscreenKeyboardA(member bool) {
    self.Object.Set("fullscreenKeyboard", member)
}

// OnInitialized This signal is dispatched after device initialization occurs but before any of the ready
// callbacks (see {@link Phaser.Device.whenReady whenReady}) have been invoked.
// 
// Local "patching" for a particular device can/should be done in this event.
// 
// _Note_: This signal is removed after the device has been readied; if a handler has not been
// added _before_ `new Phaser.Game(..)` it is probably too late.
func (self *Device) OnInitialized() *Signal{
    return &Signal{self.Object.Get("onInitialized")}
}

// SetOnInitializedA This signal is dispatched after device initialization occurs but before any of the ready
// callbacks (see {@link Phaser.Device.whenReady whenReady}) have been invoked.
// 
// Local "patching" for a particular device can/should be done in this event.
// 
// _Note_: This signal is removed after the device has been readied; if a handler has not been
// added _before_ `new Phaser.Game(..)` it is probably too late.
func (self *Device) SetOnInitializedA(member *Signal) {
    self.Object.Set("onInitialized", member)
}


// WhenReady Add a device-ready handler and ensure the device ready sequence is started.
// 
// Phaser.Device will _not_ activate or initialize until at least one `whenReady` handler is added,
// which is normally done automatically be calling `new Phaser.Game(..)`.
// 
// The handler is invoked when the device is considered "ready", which may be immediately
// if the device is already "ready". See {@link Phaser.Device#deviceReadyAt deviceReadyAt}.
func (self *Device) WhenReady(handler interface{}) {
    self.Object.Call("whenReady", handler)
}

// WhenReady1O Add a device-ready handler and ensure the device ready sequence is started.
// 
// Phaser.Device will _not_ activate or initialize until at least one `whenReady` handler is added,
// which is normally done automatically be calling `new Phaser.Game(..)`.
// 
// The handler is invoked when the device is considered "ready", which may be immediately
// if the device is already "ready". See {@link Phaser.Device#deviceReadyAt deviceReadyAt}.
func (self *Device) WhenReady1O(handler interface{}, context interface{}) {
    self.Object.Call("whenReady", handler, context)
}

// WhenReady2O Add a device-ready handler and ensure the device ready sequence is started.
// 
// Phaser.Device will _not_ activate or initialize until at least one `whenReady` handler is added,
// which is normally done automatically be calling `new Phaser.Game(..)`.
// 
// The handler is invoked when the device is considered "ready", which may be immediately
// if the device is already "ready". See {@link Phaser.Device#deviceReadyAt deviceReadyAt}.
func (self *Device) WhenReady2O(handler interface{}, context interface{}, nonPrimer bool) {
    self.Object.Call("whenReady", handler, context, nonPrimer)
}

// WhenReadyI Add a device-ready handler and ensure the device ready sequence is started.
// 
// Phaser.Device will _not_ activate or initialize until at least one `whenReady` handler is added,
// which is normally done automatically be calling `new Phaser.Game(..)`.
// 
// The handler is invoked when the device is considered "ready", which may be immediately
// if the device is already "ready". See {@link Phaser.Device#deviceReadyAt deviceReadyAt}.
func (self *Device) WhenReadyI(args ...interface{}) {
    self.Object.Call("whenReady", args)
}

// _readyCheck Internal method used for checking when the device is ready.
// This function is removed from Phaser.Device when the device becomes ready.
func (self *Device) _readyCheck() {
    self.Object.Call("_readyCheck")
}

// _readyCheckI Internal method used for checking when the device is ready.
// This function is removed from Phaser.Device when the device becomes ready.
func (self *Device) _readyCheckI(args ...interface{}) {
    self.Object.Call("_readyCheck", args)
}

// _initialize Internal method to initialize the capability checks.
// This function is removed from Phaser.Device once the device is initialized.
func (self *Device) _initialize() {
    self.Object.Call("_initialize")
}

// _initializeI Internal method to initialize the capability checks.
// This function is removed from Phaser.Device once the device is initialized.
func (self *Device) _initializeI(args ...interface{}) {
    self.Object.Call("_initialize", args)
}

// CanPlayAudio Check whether the host environment can play audio.
func (self *Device) CanPlayAudio(type_ string) bool{
    return self.Object.Call("canPlayAudio", type_).Bool()
}

// CanPlayAudioI Check whether the host environment can play audio.
func (self *Device) CanPlayAudioI(args ...interface{}) bool{
    return self.Object.Call("canPlayAudio", args).Bool()
}

// CanPlayVideo Check whether the host environment can play video files.
func (self *Device) CanPlayVideo(type_ string) bool{
    return self.Object.Call("canPlayVideo", type_).Bool()
}

// CanPlayVideoI Check whether the host environment can play video files.
func (self *Device) CanPlayVideoI(args ...interface{}) bool{
    return self.Object.Call("canPlayVideo", args).Bool()
}

// IsConsoleOpen Check whether the console is open.
// Note that this only works in Firefox with Firebug and earlier versions of Chrome.
// It used to work in Chrome, but then they removed the ability: {@link http://src.chromium.org/viewvc/blink?view=revision&revision=151136}
func (self *Device) IsConsoleOpen() {
    self.Object.Call("isConsoleOpen")
}

// IsConsoleOpenI Check whether the console is open.
// Note that this only works in Firefox with Firebug and earlier versions of Chrome.
// It used to work in Chrome, but then they removed the ability: {@link http://src.chromium.org/viewvc/blink?view=revision&revision=151136}
func (self *Device) IsConsoleOpenI(args ...interface{}) {
    self.Object.Call("isConsoleOpen", args)
}

// IsAndroidStockBrowser Detect if the host is a an Android Stock browser.
// This is available before the device "ready" event.
// 
// Authors might want to scale down on effects and switch to the CANVAS rendering method on those devices.
func (self *Device) IsAndroidStockBrowser() {
    self.Object.Call("isAndroidStockBrowser")
}

// IsAndroidStockBrowserI Detect if the host is a an Android Stock browser.
// This is available before the device "ready" event.
// 
// Authors might want to scale down on effects and switch to the CANVAS rendering method on those devices.
func (self *Device) IsAndroidStockBrowserI(args ...interface{}) {
    self.Object.Call("isAndroidStockBrowser", args)
}

