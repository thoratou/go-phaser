// Package phaser Automatic generation for Phaser.SoundManager
// generated file SoundManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// SoundManager The Sound Manager is responsible for playing back audio via either the Legacy HTML Audio tag or via Web Audio if the browser supports it.
// Note: On Firefox 25+ on Linux if you have media.gstreamer disabled in about:config then it cannot play back mp3 or m4a files.
// The audio file type and the encoding of those files are extremely important. Not all browsers can play all audio formats.
// There is a good guide to what's supported here: http://hpr.dogphilosophy.net/test/
// 
// If you are reloading a Phaser Game on a page that never properly refreshes (such as in an AngularJS project) then you will quickly run out
// of AudioContext nodes. If this is the case create a global var called PhaserGlobal on the window object before creating the game. The active
// AudioContext will then be saved to window.PhaserGlobal.audioContext when the Phaser game is destroyed, and re-used when it starts again.
// 
// Mobile warning: There are some mobile devices (certain iPad 2 and iPad Mini revisions) that cannot play 48000 Hz audio.
// When they try to play the audio becomes extremely distorted and buzzes, eventually crashing the sound system.
// The solution is to use a lower encoding rate such as 44100 Hz. Sometimes the audio context will
// be created with a sampleRate of 48000. If this happens and audio distorts you should re-create the context.
type SoundManager struct {
    *js.Object
}

// NewSoundManager The Sound Manager is responsible for playing back audio via either the Legacy HTML Audio tag or via Web Audio if the browser supports it.
// Note: On Firefox 25+ on Linux if you have media.gstreamer disabled in about:config then it cannot play back mp3 or m4a files.
// The audio file type and the encoding of those files are extremely important. Not all browsers can play all audio formats.
// There is a good guide to what's supported here: http://hpr.dogphilosophy.net/test/
// 
// If you are reloading a Phaser Game on a page that never properly refreshes (such as in an AngularJS project) then you will quickly run out
// of AudioContext nodes. If this is the case create a global var called PhaserGlobal on the window object before creating the game. The active
// AudioContext will then be saved to window.PhaserGlobal.audioContext when the Phaser game is destroyed, and re-used when it starts again.
// 
// Mobile warning: There are some mobile devices (certain iPad 2 and iPad Mini revisions) that cannot play 48000 Hz audio.
// When they try to play the audio becomes extremely distorted and buzzes, eventually crashing the sound system.
// The solution is to use a lower encoding rate such as 44100 Hz. Sometimes the audio context will
// be created with a sampleRate of 48000. If this happens and audio distorts you should re-create the context.
func NewSoundManager(game *Game) *SoundManager {
    return &SoundManager{js.Global.Get("Phaser").Get("SoundManager").New(game)}
}
// NewSoundManagerI The Sound Manager is responsible for playing back audio via either the Legacy HTML Audio tag or via Web Audio if the browser supports it.
// Note: On Firefox 25+ on Linux if you have media.gstreamer disabled in about:config then it cannot play back mp3 or m4a files.
// The audio file type and the encoding of those files are extremely important. Not all browsers can play all audio formats.
// There is a good guide to what's supported here: http://hpr.dogphilosophy.net/test/
// 
// If you are reloading a Phaser Game on a page that never properly refreshes (such as in an AngularJS project) then you will quickly run out
// of AudioContext nodes. If this is the case create a global var called PhaserGlobal on the window object before creating the game. The active
// AudioContext will then be saved to window.PhaserGlobal.audioContext when the Phaser game is destroyed, and re-used when it starts again.
// 
// Mobile warning: There are some mobile devices (certain iPad 2 and iPad Mini revisions) that cannot play 48000 Hz audio.
// When they try to play the audio becomes extremely distorted and buzzes, eventually crashing the sound system.
// The solution is to use a lower encoding rate such as 44100 Hz. Sometimes the audio context will
// be created with a sampleRate of 48000. If this happens and audio distorts you should re-create the context.
func NewSoundManagerI(args ...interface{}) *SoundManager {
    return &SoundManager{js.Global.Get("Phaser").Get("SoundManager").New(args)}
}



// SoundManager Binding conversion method to SoundManager point 
func ToSoundManager(jsStruct interface{}) *SoundManager {
    if object, ok := jsStruct.(*js.Object); ok {
		return &SoundManager{Object: object}
	}
	return nil
}



// Game Local reference to game.
func (self *SoundManager) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *SoundManager) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// OnSoundDecode The event dispatched when a sound decodes (typically only for mp3 files)
func (self *SoundManager) OnSoundDecode() *Signal{
    return &Signal{self.Object.Get("onSoundDecode")}
}

// SetOnSoundDecodeA The event dispatched when a sound decodes (typically only for mp3 files)
func (self *SoundManager) SetOnSoundDecodeA(member *Signal) {
    self.Object.Set("onSoundDecode", member)
}

// OnVolumeChange This signal is dispatched whenever the global volume changes. The new volume is passed as the only parameter to your callback.
func (self *SoundManager) OnVolumeChange() *Signal{
    return &Signal{self.Object.Get("onVolumeChange")}
}

// SetOnVolumeChangeA This signal is dispatched whenever the global volume changes. The new volume is passed as the only parameter to your callback.
func (self *SoundManager) SetOnVolumeChangeA(member *Signal) {
    self.Object.Set("onVolumeChange", member)
}

// OnMute This signal is dispatched when the SoundManager is globally muted, either directly via game code or as a result of the game pausing.
func (self *SoundManager) OnMute() *Signal{
    return &Signal{self.Object.Get("onMute")}
}

// SetOnMuteA This signal is dispatched when the SoundManager is globally muted, either directly via game code or as a result of the game pausing.
func (self *SoundManager) SetOnMuteA(member *Signal) {
    self.Object.Set("onMute", member)
}

// OnUnMute This signal is dispatched when the SoundManager is globally un-muted, either directly via game code or as a result of the game resuming from a pause.
func (self *SoundManager) OnUnMute() *Signal{
    return &Signal{self.Object.Get("onUnMute")}
}

// SetOnUnMuteA This signal is dispatched when the SoundManager is globally un-muted, either directly via game code or as a result of the game resuming from a pause.
func (self *SoundManager) SetOnUnMuteA(member *Signal) {
    self.Object.Set("onUnMute", member)
}

// Context The AudioContext being used for playback.
func (self *SoundManager) Context() *AudioContext{
    return &AudioContext{self.Object.Get("context")}
}

// SetContextA The AudioContext being used for playback.
func (self *SoundManager) SetContextA(member *AudioContext) {
    self.Object.Set("context", member)
}

// UsingWebAudio True the SoundManager and device are both using Web Audio.
func (self *SoundManager) UsingWebAudio() bool{
    return self.Object.Get("usingWebAudio").Bool()
}

// SetUsingWebAudioA True the SoundManager and device are both using Web Audio.
func (self *SoundManager) SetUsingWebAudioA(member bool) {
    self.Object.Set("usingWebAudio", member)
}

// UsingAudioTag True the SoundManager and device are both using the Audio tag instead of Web Audio.
func (self *SoundManager) UsingAudioTag() bool{
    return self.Object.Get("usingAudioTag").Bool()
}

// SetUsingAudioTagA True the SoundManager and device are both using the Audio tag instead of Web Audio.
func (self *SoundManager) SetUsingAudioTagA(member bool) {
    self.Object.Set("usingAudioTag", member)
}

// NoAudio True if audio been disabled via the PhaserGlobal (useful if you need to use a 3rd party audio library) or the device doesn't support any audio.
func (self *SoundManager) NoAudio() bool{
    return self.Object.Get("noAudio").Bool()
}

// SetNoAudioA True if audio been disabled via the PhaserGlobal (useful if you need to use a 3rd party audio library) or the device doesn't support any audio.
func (self *SoundManager) SetNoAudioA(member bool) {
    self.Object.Set("noAudio", member)
}

// ConnectToMaster Used in conjunction with Sound.externalNode this allows you to stop a Sound node being connected to the SoundManager master gain node.
func (self *SoundManager) ConnectToMaster() bool{
    return self.Object.Get("connectToMaster").Bool()
}

// SetConnectToMasterA Used in conjunction with Sound.externalNode this allows you to stop a Sound node being connected to the SoundManager master gain node.
func (self *SoundManager) SetConnectToMasterA(member bool) {
    self.Object.Set("connectToMaster", member)
}

// TouchLocked true if the audio system is currently locked awaiting a touch event.
func (self *SoundManager) TouchLocked() bool{
    return self.Object.Get("touchLocked").Bool()
}

// SetTouchLockedA true if the audio system is currently locked awaiting a touch event.
func (self *SoundManager) SetTouchLockedA(member bool) {
    self.Object.Set("touchLocked", member)
}

// Channels The number of audio channels to use in playback.
func (self *SoundManager) Channels() int{
    return self.Object.Get("channels").Int()
}

// SetChannelsA The number of audio channels to use in playback.
func (self *SoundManager) SetChannelsA(member int) {
    self.Object.Set("channels", member)
}

// MuteOnPause Set to true to have all sound muted when the Phaser game pauses (such as on loss of focus),
// or set to false to keep audio playing, regardless of the game pause state. You may need to
// do this should you wish to control audio muting via external DOM buttons or similar.
func (self *SoundManager) MuteOnPause() bool{
    return self.Object.Get("muteOnPause").Bool()
}

// SetMuteOnPauseA Set to true to have all sound muted when the Phaser game pauses (such as on loss of focus),
// or set to false to keep audio playing, regardless of the game pause state. You may need to
// do this should you wish to control audio muting via external DOM buttons or similar.
func (self *SoundManager) SetMuteOnPauseA(member bool) {
    self.Object.Set("muteOnPause", member)
}

// Mute Gets or sets the muted state of the SoundManager. This effects all sounds in the game.
func (self *SoundManager) Mute() bool{
    return self.Object.Get("mute").Bool()
}

// SetMuteA Gets or sets the muted state of the SoundManager. This effects all sounds in the game.
func (self *SoundManager) SetMuteA(member bool) {
    self.Object.Set("mute", member)
}

// Volume Gets or sets the global volume of the SoundManager, a value between 0 and 1. The value given is clamped to the range 0 to 1.
func (self *SoundManager) Volume() int{
    return self.Object.Get("volume").Int()
}

// SetVolumeA Gets or sets the global volume of the SoundManager, a value between 0 and 1. The value given is clamped to the range 0 to 1.
func (self *SoundManager) SetVolumeA(member int) {
    self.Object.Set("volume", member)
}


// Boot Initialises the sound manager.
func (self *SoundManager) Boot() {
    self.Object.Call("boot")
}

// BootI Initialises the sound manager.
func (self *SoundManager) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// SetTouchLock Sets the Input Manager touch callback to be SoundManager.unlock.
// Required for iOS audio device unlocking. Mostly just used internally.
func (self *SoundManager) SetTouchLock() {
    self.Object.Call("setTouchLock")
}

// SetTouchLockI Sets the Input Manager touch callback to be SoundManager.unlock.
// Required for iOS audio device unlocking. Mostly just used internally.
func (self *SoundManager) SetTouchLockI(args ...interface{}) {
    self.Object.Call("setTouchLock", args)
}

// Unlock Enables the audio, usually after the first touch.
func (self *SoundManager) Unlock() bool{
    return self.Object.Call("unlock").Bool()
}

// UnlockI Enables the audio, usually after the first touch.
func (self *SoundManager) UnlockI(args ...interface{}) bool{
    return self.Object.Call("unlock", args).Bool()
}

// StopAll Stops all the sounds in the game.
func (self *SoundManager) StopAll() {
    self.Object.Call("stopAll")
}

// StopAllI Stops all the sounds in the game.
func (self *SoundManager) StopAllI(args ...interface{}) {
    self.Object.Call("stopAll", args)
}

// PauseAll Pauses all the sounds in the game.
func (self *SoundManager) PauseAll() {
    self.Object.Call("pauseAll")
}

// PauseAllI Pauses all the sounds in the game.
func (self *SoundManager) PauseAllI(args ...interface{}) {
    self.Object.Call("pauseAll", args)
}

// ResumeAll Resumes every sound in the game.
func (self *SoundManager) ResumeAll() {
    self.Object.Call("resumeAll")
}

// ResumeAllI Resumes every sound in the game.
func (self *SoundManager) ResumeAllI(args ...interface{}) {
    self.Object.Call("resumeAll", args)
}

// Decode Decode a sound by its asset key.
func (self *SoundManager) Decode(key string) {
    self.Object.Call("decode", key)
}

// Decode1O Decode a sound by its asset key.
func (self *SoundManager) Decode1O(key string, sound *Sound) {
    self.Object.Call("decode", key, sound)
}

// DecodeI Decode a sound by its asset key.
func (self *SoundManager) DecodeI(args ...interface{}) {
    self.Object.Call("decode", args)
}

// SetDecodedCallback This method allows you to give the SoundManager a list of Sound files, or keys, and a callback.
// Once all of the Sound files have finished decoding the callback will be invoked.
// The amount of time spent decoding depends on the codec used and file size.
// If all of the files given have already decoded the callback is triggered immediately.
func (self *SoundManager) SetDecodedCallback(files interface{}, callback interface{}, callbackContext interface{}) {
    self.Object.Call("setDecodedCallback", files, callback, callbackContext)
}

// SetDecodedCallbackI This method allows you to give the SoundManager a list of Sound files, or keys, and a callback.
// Once all of the Sound files have finished decoding the callback will be invoked.
// The amount of time spent decoding depends on the codec used and file size.
// If all of the files given have already decoded the callback is triggered immediately.
func (self *SoundManager) SetDecodedCallbackI(args ...interface{}) {
    self.Object.Call("setDecodedCallback", args)
}

// Update Updates every sound in the game, checks for audio unlock on mobile and monitors the decoding watch list.
func (self *SoundManager) Update() {
    self.Object.Call("update")
}

// UpdateI Updates every sound in the game, checks for audio unlock on mobile and monitors the decoding watch list.
func (self *SoundManager) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Add Adds a new Sound into the SoundManager.
func (self *SoundManager) Add(key string) *Sound{
    return &Sound{self.Object.Call("add", key)}
}

// Add1O Adds a new Sound into the SoundManager.
func (self *SoundManager) Add1O(key string, volume int) *Sound{
    return &Sound{self.Object.Call("add", key, volume)}
}

// Add2O Adds a new Sound into the SoundManager.
func (self *SoundManager) Add2O(key string, volume int, loop bool) *Sound{
    return &Sound{self.Object.Call("add", key, volume, loop)}
}

// Add3O Adds a new Sound into the SoundManager.
func (self *SoundManager) Add3O(key string, volume int, loop bool, connect bool) *Sound{
    return &Sound{self.Object.Call("add", key, volume, loop, connect)}
}

// AddI Adds a new Sound into the SoundManager.
func (self *SoundManager) AddI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("add", args)}
}

// AddSprite Adds a new AudioSprite into the SoundManager.
func (self *SoundManager) AddSprite(key string) *AudioSprite{
    return &AudioSprite{self.Object.Call("addSprite", key)}
}

// AddSpriteI Adds a new AudioSprite into the SoundManager.
func (self *SoundManager) AddSpriteI(args ...interface{}) *AudioSprite{
    return &AudioSprite{self.Object.Call("addSprite", args)}
}

// Remove Removes a Sound from the SoundManager. The removed Sound is destroyed before removal.
func (self *SoundManager) Remove(sound *Sound) bool{
    return self.Object.Call("remove", sound).Bool()
}

// RemoveI Removes a Sound from the SoundManager. The removed Sound is destroyed before removal.
func (self *SoundManager) RemoveI(args ...interface{}) bool{
    return self.Object.Call("remove", args).Bool()
}

// RemoveByKey Removes all Sounds from the SoundManager that have an asset key matching the given value.
// The removed Sounds are destroyed before removal.
func (self *SoundManager) RemoveByKey(key string) int{
    return self.Object.Call("removeByKey", key).Int()
}

// RemoveByKeyI Removes all Sounds from the SoundManager that have an asset key matching the given value.
// The removed Sounds are destroyed before removal.
func (self *SoundManager) RemoveByKeyI(args ...interface{}) int{
    return self.Object.Call("removeByKey", args).Int()
}

// Play Adds a new Sound into the SoundManager and starts it playing.
func (self *SoundManager) Play(key string) *Sound{
    return &Sound{self.Object.Call("play", key)}
}

// Play1O Adds a new Sound into the SoundManager and starts it playing.
func (self *SoundManager) Play1O(key string, volume int) *Sound{
    return &Sound{self.Object.Call("play", key, volume)}
}

// Play2O Adds a new Sound into the SoundManager and starts it playing.
func (self *SoundManager) Play2O(key string, volume int, loop bool) *Sound{
    return &Sound{self.Object.Call("play", key, volume, loop)}
}

// PlayI Adds a new Sound into the SoundManager and starts it playing.
func (self *SoundManager) PlayI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("play", args)}
}

// SetMute Internal mute handler called automatically by the SoundManager.mute setter.
func (self *SoundManager) SetMute() {
    self.Object.Call("setMute")
}

// SetMuteI Internal mute handler called automatically by the SoundManager.mute setter.
func (self *SoundManager) SetMuteI(args ...interface{}) {
    self.Object.Call("setMute", args)
}

// UnsetMute Internal mute handler called automatically by the SoundManager.mute setter.
func (self *SoundManager) UnsetMute() {
    self.Object.Call("unsetMute")
}

// UnsetMuteI Internal mute handler called automatically by the SoundManager.mute setter.
func (self *SoundManager) UnsetMuteI(args ...interface{}) {
    self.Object.Call("unsetMute", args)
}

// Destroy Stops all the sounds in the game, then destroys them and finally clears up any callbacks.
func (self *SoundManager) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Stops all the sounds in the game, then destroys them and finally clears up any callbacks.
func (self *SoundManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

