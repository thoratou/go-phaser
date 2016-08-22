// Automatic generation for Phaser.SoundManager
// generated file SoundManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Sound Manager is responsible for playing back audio via either the Legacy HTML Audio tag or via Web Audio if the browser supports it.
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


// Local reference to game.
func (self *SoundManager) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *SoundManager) SetGame(member *Game) {
    self.Set("game", member)
}

// The event dispatched when a sound decodes (typically only for mp3 files)
func (self *SoundManager) GetOnSoundDecode() *Signal{
    return &Signal{self.Get("onSoundDecode")}
}

// The event dispatched when a sound decodes (typically only for mp3 files)
func (self *SoundManager) SetOnSoundDecode(member *Signal) {
    self.Set("onSoundDecode", member)
}

// This signal is dispatched whenever the global volume changes. The new volume is passed as the only parameter to your callback.
func (self *SoundManager) GetOnVolumeChange() *Signal{
    return &Signal{self.Get("onVolumeChange")}
}

// This signal is dispatched whenever the global volume changes. The new volume is passed as the only parameter to your callback.
func (self *SoundManager) SetOnVolumeChange(member *Signal) {
    self.Set("onVolumeChange", member)
}

// This signal is dispatched when the SoundManager is globally muted, either directly via game code or as a result of the game pausing.
func (self *SoundManager) GetOnMute() *Signal{
    return &Signal{self.Get("onMute")}
}

// This signal is dispatched when the SoundManager is globally muted, either directly via game code or as a result of the game pausing.
func (self *SoundManager) SetOnMute(member *Signal) {
    self.Set("onMute", member)
}

// This signal is dispatched when the SoundManager is globally un-muted, either directly via game code or as a result of the game resuming from a pause.
func (self *SoundManager) GetOnUnMute() *Signal{
    return &Signal{self.Get("onUnMute")}
}

// This signal is dispatched when the SoundManager is globally un-muted, either directly via game code or as a result of the game resuming from a pause.
func (self *SoundManager) SetOnUnMute(member *Signal) {
    self.Set("onUnMute", member)
}

// The AudioContext being used for playback.
func (self *SoundManager) GetContext() *AudioContext{
    return &AudioContext{self.Get("context")}
}

// The AudioContext being used for playback.
func (self *SoundManager) SetContext(member *AudioContext) {
    self.Set("context", member)
}

// True the SoundManager and device are both using Web Audio.
func (self *SoundManager) GetUsingWebAudio() bool{
    return self.Get("usingWebAudio").Bool()
}

// True the SoundManager and device are both using Web Audio.
func (self *SoundManager) SetUsingWebAudio(member bool) {
    self.Set("usingWebAudio", member)
}

// True the SoundManager and device are both using the Audio tag instead of Web Audio.
func (self *SoundManager) GetUsingAudioTag() bool{
    return self.Get("usingAudioTag").Bool()
}

// True the SoundManager and device are both using the Audio tag instead of Web Audio.
func (self *SoundManager) SetUsingAudioTag(member bool) {
    self.Set("usingAudioTag", member)
}

// True if audio been disabled via the PhaserGlobal (useful if you need to use a 3rd party audio library) or the device doesn't support any audio.
func (self *SoundManager) GetNoAudio() bool{
    return self.Get("noAudio").Bool()
}

// True if audio been disabled via the PhaserGlobal (useful if you need to use a 3rd party audio library) or the device doesn't support any audio.
func (self *SoundManager) SetNoAudio(member bool) {
    self.Set("noAudio", member)
}

// Used in conjunction with Sound.externalNode this allows you to stop a Sound node being connected to the SoundManager master gain node.
func (self *SoundManager) GetConnectToMaster() bool{
    return self.Get("connectToMaster").Bool()
}

// Used in conjunction with Sound.externalNode this allows you to stop a Sound node being connected to the SoundManager master gain node.
func (self *SoundManager) SetConnectToMaster(member bool) {
    self.Set("connectToMaster", member)
}

// true if the audio system is currently locked awaiting a touch event.
func (self *SoundManager) GetTouchLocked() bool{
    return self.Get("touchLocked").Bool()
}

// true if the audio system is currently locked awaiting a touch event.
func (self *SoundManager) SetTouchLocked(member bool) {
    self.Set("touchLocked", member)
}

// The number of audio channels to use in playback.
func (self *SoundManager) GetChannels() int{
    return self.Get("channels").Int()
}

// The number of audio channels to use in playback.
func (self *SoundManager) SetChannels(member int) {
    self.Set("channels", member)
}

// Set to true to have all sound muted when the Phaser game pauses (such as on loss of focus),
// or set to false to keep audio playing, regardless of the game pause state. You may need to
// do this should you wish to control audio muting via external DOM buttons or similar.
func (self *SoundManager) GetMuteOnPause() bool{
    return self.Get("muteOnPause").Bool()
}

// Set to true to have all sound muted when the Phaser game pauses (such as on loss of focus),
// or set to false to keep audio playing, regardless of the game pause state. You may need to
// do this should you wish to control audio muting via external DOM buttons or similar.
func (self *SoundManager) SetMuteOnPause(member bool) {
    self.Set("muteOnPause", member)
}

// Gets or sets the muted state of the SoundManager. This effects all sounds in the game.
func (self *SoundManager) GetMute() bool{
    return self.Get("mute").Bool()
}

// Gets or sets the muted state of the SoundManager. This effects all sounds in the game.
func (self *SoundManager) SetMute(member bool) {
    self.Set("mute", member)
}

// Gets or sets the global volume of the SoundManager, a value between 0 and 1. The value given is clamped to the range 0 to 1.
func (self *SoundManager) GetVolume() int{
    return self.Get("volume").Int()
}

// Gets or sets the global volume of the SoundManager, a value between 0 and 1. The value given is clamped to the range 0 to 1.
func (self *SoundManager) SetVolume(member int) {
    self.Set("volume", member)
}



// Initialises the sound manager.
func (self *SoundManager) BootI(args ...interface{}) {
    self.Call("boot", args)
}

// Sets the Input Manager touch callback to be SoundManager.unlock.
// Required for iOS audio device unlocking. Mostly just used internally.
func (self *SoundManager) SetTouchLockI(args ...interface{}) {
    self.Call("setTouchLock", args)
}

// Enables the audio, usually after the first touch.
func (self *SoundManager) UnlockI(args ...interface{}) bool{
    return self.Call("unlock", args).Bool()
}

// Stops all the sounds in the game.
func (self *SoundManager) StopAllI(args ...interface{}) {
    self.Call("stopAll", args)
}

// Pauses all the sounds in the game.
func (self *SoundManager) PauseAllI(args ...interface{}) {
    self.Call("pauseAll", args)
}

// Resumes every sound in the game.
func (self *SoundManager) ResumeAllI(args ...interface{}) {
    self.Call("resumeAll", args)
}

// Decode a sound by its asset key.
func (self *SoundManager) DecodeI(args ...interface{}) {
    self.Call("decode", args)
}

// This method allows you to give the SoundManager a list of Sound files, or keys, and a callback.
// Once all of the Sound files have finished decoding the callback will be invoked.
// The amount of time spent decoding depends on the codec used and file size.
// If all of the files given have already decoded the callback is triggered immediately.
func (self *SoundManager) SetDecodedCallbackI(args ...interface{}) {
    self.Call("setDecodedCallback", args)
}

// Updates every sound in the game, checks for audio unlock on mobile and monitors the decoding watch list.
func (self *SoundManager) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Adds a new Sound into the SoundManager.
func (self *SoundManager) AddI(args ...interface{}) *Sound{
    return &Sound{self.Call("add", args)}
}

// Adds a new AudioSprite into the SoundManager.
func (self *SoundManager) AddSpriteI(args ...interface{}) *AudioSprite{
    return &AudioSprite{self.Call("addSprite", args)}
}

// Removes a Sound from the SoundManager. The removed Sound is destroyed before removal.
func (self *SoundManager) RemoveI(args ...interface{}) bool{
    return self.Call("remove", args).Bool()
}

// Removes all Sounds from the SoundManager that have an asset key matching the given value.
// The removed Sounds are destroyed before removal.
func (self *SoundManager) RemoveByKeyI(args ...interface{}) int{
    return self.Call("removeByKey", args).Int()
}

// Adds a new Sound into the SoundManager and starts it playing.
func (self *SoundManager) PlayI(args ...interface{}) *Sound{
    return &Sound{self.Call("play", args)}
}

// Internal mute handler called automatically by the SoundManager.mute setter.
func (self *SoundManager) SetMuteI(args ...interface{}) {
    self.Call("setMute", args)
}

// Internal mute handler called automatically by the SoundManager.mute setter.
func (self *SoundManager) UnsetMuteI(args ...interface{}) {
    self.Call("unsetMute", args)
}

// Stops all the sounds in the game, then destroys them and finally clears up any callbacks.
func (self *SoundManager) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
