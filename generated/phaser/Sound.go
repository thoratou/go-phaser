// Automatic generation for Phaser.Sound
// generated file Sound.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Sound class constructor.
type Sound struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *Sound) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Sound) SetGame(member Game) {
    self.Set("game", member)
}

// Name of the sound.
func (self *Sound) GetName() string{
    return self.Get("name").String()
}

// Name of the sound.
func (self *Sound) SetName(member string) {
    self.Set("name", member)
}

// Asset key for the sound.
func (self *Sound) GetKey() string{
    return self.Get("key").String()
}

// Asset key for the sound.
func (self *Sound) SetKey(member string) {
    self.Set("key", member)
}

// Whether or not the sound or current sound marker will loop.
func (self *Sound) GetLoop() bool{
    return self.Get("loop").Bool()
}

// Whether or not the sound or current sound marker will loop.
func (self *Sound) SetLoop(member bool) {
    self.Set("loop", member)
}

// The sound markers.
func (self *Sound) GetMarkers() interface{}{
    return self.Get("markers")
}

// The sound markers.
func (self *Sound) SetMarkers(member interface{}) {
    self.Set("markers", member)
}

// Reference to the AudioContext instance.
func (self *Sound) GetContext() AudioContext{
    return AudioContext{self.Get("context")}
}

// Reference to the AudioContext instance.
func (self *Sound) SetContext(member AudioContext) {
    self.Set("context", member)
}

// Boolean indicating whether the sound should start automatically.
func (self *Sound) GetAutoplay() bool{
    return self.Get("autoplay").Bool()
}

// Boolean indicating whether the sound should start automatically.
func (self *Sound) SetAutoplay(member bool) {
    self.Set("autoplay", member)
}

// The total duration of the sound in seconds.
func (self *Sound) GetTotalDuration() float64{
    return self.Get("totalDuration").Float()
}

// The total duration of the sound in seconds.
func (self *Sound) SetTotalDuration(member float64) {
    self.Set("totalDuration", member)
}

// The time the Sound starts at (typically 0 unless starting from a marker)
func (self *Sound) GetStartTime() float64{
    return self.Get("startTime").Float()
}

// The time the Sound starts at (typically 0 unless starting from a marker)
func (self *Sound) SetStartTime(member float64) {
    self.Set("startTime", member)
}

// The current time the sound is at.
func (self *Sound) GetCurrentTime() float64{
    return self.Get("currentTime").Float()
}

// The current time the sound is at.
func (self *Sound) SetCurrentTime(member float64) {
    self.Set("currentTime", member)
}

// The duration of the current sound marker in seconds.
func (self *Sound) GetDuration() float64{
    return self.Get("duration").Float()
}

// The duration of the current sound marker in seconds.
func (self *Sound) SetDuration(member float64) {
    self.Set("duration", member)
}

// The duration of the current sound marker in ms.
func (self *Sound) GetDurationMS() float64{
    return self.Get("durationMS").Float()
}

// The duration of the current sound marker in ms.
func (self *Sound) SetDurationMS(member float64) {
    self.Set("durationMS", member)
}

// The position of the current sound marker.
func (self *Sound) GetPosition() float64{
    return self.Get("position").Float()
}

// The position of the current sound marker.
func (self *Sound) SetPosition(member float64) {
    self.Set("position", member)
}

// The time the sound stopped.
func (self *Sound) GetStopTime() float64{
    return self.Get("stopTime").Float()
}

// The time the sound stopped.
func (self *Sound) SetStopTime(member float64) {
    self.Set("stopTime", member)
}

// true if the sound is paused, otherwise false.
func (self *Sound) GetPaused() bool{
    return self.Get("paused").Bool()
}

// true if the sound is paused, otherwise false.
func (self *Sound) SetPaused(member bool) {
    self.Set("paused", member)
}

// The position the sound had reached when it was paused.
func (self *Sound) GetPausedPosition() float64{
    return self.Get("pausedPosition").Float()
}

// The position the sound had reached when it was paused.
func (self *Sound) SetPausedPosition(member float64) {
    self.Set("pausedPosition", member)
}

// The game time at which the sound was paused.
func (self *Sound) GetPausedTime() float64{
    return self.Get("pausedTime").Float()
}

// The game time at which the sound was paused.
func (self *Sound) SetPausedTime(member float64) {
    self.Set("pausedTime", member)
}

// true if the sound is currently playing, otherwise false.
func (self *Sound) GetIsPlaying() bool{
    return self.Get("isPlaying").Bool()
}

// true if the sound is currently playing, otherwise false.
func (self *Sound) SetIsPlaying(member bool) {
    self.Set("isPlaying", member)
}

// The string ID of the currently playing marker, if any.
func (self *Sound) GetCurrentMarker() string{
    return self.Get("currentMarker").String()
}

// The string ID of the currently playing marker, if any.
func (self *Sound) SetCurrentMarker(member string) {
    self.Set("currentMarker", member)
}

// The tween that fades the audio, set via Sound.fadeIn and Sound.fadeOut.
func (self *Sound) GetFadeTween() Tween{
    return Tween{self.Get("fadeTween")}
}

// The tween that fades the audio, set via Sound.fadeIn and Sound.fadeOut.
func (self *Sound) SetFadeTween(member Tween) {
    self.Set("fadeTween", member)
}

// true if the sound file is pending playback
func (self *Sound) GetPendingPlayback() bool{
    return self.Get("pendingPlayback").Bool()
}

// true if the sound file is pending playback
func (self *Sound) SetPendingPlayback(member bool) {
    self.Set("pendingPlayback", member)
}

// if true when you play this sound it will always start from the beginning.
func (self *Sound) GetOverride() bool{
    return self.Get("override").Bool()
}

// if true when you play this sound it will always start from the beginning.
func (self *Sound) SetOverride(member bool) {
    self.Set("override", member)
}

// This will allow you to have multiple instances of this Sound playing at once. This is only useful when running under Web Audio, and we recommend you implement a local pooling system to not flood the sound channels.
func (self *Sound) GetAllowMultiple() bool{
    return self.Get("allowMultiple").Bool()
}

// This will allow you to have multiple instances of this Sound playing at once. This is only useful when running under Web Audio, and we recommend you implement a local pooling system to not flood the sound channels.
func (self *Sound) SetAllowMultiple(member bool) {
    self.Set("allowMultiple", member)
}

// true if this sound is being played with Web Audio.
func (self *Sound) GetUsingWebAudio() bool{
    return self.Get("usingWebAudio").Bool()
}

// true if this sound is being played with Web Audio.
func (self *Sound) SetUsingWebAudio(member bool) {
    self.Set("usingWebAudio", member)
}

// true if the sound is being played via the Audio tag.
func (self *Sound) GetUsingAudioTag() bool{
    return self.Get("usingAudioTag").Bool()
}

// true if the sound is being played via the Audio tag.
func (self *Sound) SetUsingAudioTag(member bool) {
    self.Set("usingAudioTag", member)
}

// If defined this Sound won't connect to the SoundManager master gain node, but will instead connect to externalNode.
func (self *Sound) GetExternalNode() interface{}{
    return self.Get("externalNode")
}

// If defined this Sound won't connect to the SoundManager master gain node, but will instead connect to externalNode.
func (self *Sound) SetExternalNode(member interface{}) {
    self.Set("externalNode", member)
}

// The master gain node in a Web Audio system.
func (self *Sound) GetMasterGainNode() interface{}{
    return self.Get("masterGainNode")
}

// The master gain node in a Web Audio system.
func (self *Sound) SetMasterGainNode(member interface{}) {
    self.Set("masterGainNode", member)
}

// The gain node in a Web Audio system.
func (self *Sound) GetGainNode() interface{}{
    return self.Get("gainNode")
}

// The gain node in a Web Audio system.
func (self *Sound) SetGainNode(member interface{}) {
    self.Set("gainNode", member)
}

// The onDecoded event is dispatched when the sound has finished decoding (typically for mp3 files)
func (self *Sound) GetOnDecoded() Signal{
    return Signal{self.Get("onDecoded")}
}

// The onDecoded event is dispatched when the sound has finished decoding (typically for mp3 files)
func (self *Sound) SetOnDecoded(member Signal) {
    self.Set("onDecoded", member)
}

// The onPlay event is dispatched each time this sound is played.
func (self *Sound) GetOnPlay() Signal{
    return Signal{self.Get("onPlay")}
}

// The onPlay event is dispatched each time this sound is played.
func (self *Sound) SetOnPlay(member Signal) {
    self.Set("onPlay", member)
}

// The onPause event is dispatched when this sound is paused.
func (self *Sound) GetOnPause() Signal{
    return Signal{self.Get("onPause")}
}

// The onPause event is dispatched when this sound is paused.
func (self *Sound) SetOnPause(member Signal) {
    self.Set("onPause", member)
}

// The onResume event is dispatched when this sound is resumed from a paused state.
func (self *Sound) GetOnResume() Signal{
    return Signal{self.Get("onResume")}
}

// The onResume event is dispatched when this sound is resumed from a paused state.
func (self *Sound) SetOnResume(member Signal) {
    self.Set("onResume", member)
}

// The onLoop event is dispatched when this sound loops during playback.
func (self *Sound) GetOnLoop() Signal{
    return Signal{self.Get("onLoop")}
}

// The onLoop event is dispatched when this sound loops during playback.
func (self *Sound) SetOnLoop(member Signal) {
    self.Set("onLoop", member)
}

// The onStop event is dispatched when this sound stops playback.
func (self *Sound) GetOnStop() Signal{
    return Signal{self.Get("onStop")}
}

// The onStop event is dispatched when this sound stops playback.
func (self *Sound) SetOnStop(member Signal) {
    self.Set("onStop", member)
}

// The onMute event is dispatched when this sound is muted.
func (self *Sound) GetOnMute() Signal{
    return Signal{self.Get("onMute")}
}

// The onMute event is dispatched when this sound is muted.
func (self *Sound) SetOnMute(member Signal) {
    self.Set("onMute", member)
}

// The onMarkerComplete event is dispatched when a marker within this sound completes playback.
func (self *Sound) GetOnMarkerComplete() Signal{
    return Signal{self.Get("onMarkerComplete")}
}

// The onMarkerComplete event is dispatched when a marker within this sound completes playback.
func (self *Sound) SetOnMarkerComplete(member Signal) {
    self.Set("onMarkerComplete", member)
}

// The onFadeComplete event is dispatched when this sound finishes fading either in or out.
func (self *Sound) GetOnFadeComplete() Signal{
    return Signal{self.Get("onFadeComplete")}
}

// The onFadeComplete event is dispatched when this sound finishes fading either in or out.
func (self *Sound) SetOnFadeComplete(member Signal) {
    self.Set("onFadeComplete", member)
}

// Returns true if the sound file is still decoding.
func (self *Sound) GetIsDecoding() bool{
    return self.Get("isDecoding").Bool()
}

// Returns true if the sound file is still decoding.
func (self *Sound) SetIsDecoding(member bool) {
    self.Set("isDecoding", member)
}

// Returns true if the sound file has decoded.
func (self *Sound) GetIsDecoded() bool{
    return self.Get("isDecoded").Bool()
}

// Returns true if the sound file has decoded.
func (self *Sound) SetIsDecoded(member bool) {
    self.Set("isDecoded", member)
}

// Gets or sets the muted state of this sound.
func (self *Sound) GetMute() bool{
    return self.Get("mute").Bool()
}

// Gets or sets the muted state of this sound.
func (self *Sound) SetMute(member bool) {
    self.Set("mute", member)
}

// Gets or sets the volume of this sound, a value between 0 and 1.
func (self *Sound) GetVolume() float64{
    return self.Get("volume").Float()
}

// Gets or sets the volume of this sound, a value between 0 and 1.
func (self *Sound) SetVolume(member float64) {
    self.Set("volume", member)
}



// Called automatically when this sound is unlocked.
func (self *Sound) SoundHasUnlockedI(args ...interface{}) {
    self.Call("soundHasUnlocked", args)
}

// Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarkerI(args ...interface{}) {
    self.Call("addMarker", args)
}

// Removes a marker from the sound.
func (self *Sound) RemoveMarkerI(args ...interface{}) {
    self.Call("removeMarker", args)
}

// Called automatically by the AudioContext when the sound stops playing.
// Doesn't get called if the sound is set to loop or is a section of an Audio Sprite.
func (self *Sound) OnEndedHandlerI(args ...interface{}) {
    self.Call("onEndedHandler", args)
}

// Called automatically by Phaser.SoundManager.
func (self *Sound) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Loops this entire sound. If you need to loop a section of it then use Sound.play and the marker and loop parameters.
func (self *Sound) LoopFullI(args ...interface{}) Sound{
    return Sound{self.Call("loopFull", args)}
}

// Play this sound, or a marked section of it.
func (self *Sound) PlayI(args ...interface{}) Sound{
    return Sound{self.Call("play", args)}
}

// Restart the sound, or a marked section of it.
func (self *Sound) RestartI(args ...interface{}) {
    self.Call("restart", args)
}

// Pauses the sound.
func (self *Sound) PauseI(args ...interface{}) {
    self.Call("pause", args)
}

// Resumes the sound.
func (self *Sound) ResumeI(args ...interface{}) {
    self.Call("resume", args)
}

// Stop playing this sound.
func (self *Sound) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeInI(args ...interface{}) {
    self.Call("fadeIn", args)
}

// Decreases the volume of this Sound from its current value to 0 over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (0) as the second parameter.
func (self *Sound) FadeOutI(args ...interface{}) {
    self.Call("fadeOut", args)
}

// Fades the volume of this Sound from its current value to the given volume over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter, 
// and the final volume (volume) as the second parameter.
func (self *Sound) FadeToI(args ...interface{}) {
    self.Call("fadeTo", args)
}

// Internal handler for Sound.fadeIn, Sound.fadeOut and Sound.fadeTo.
func (self *Sound) FadeCompleteI(args ...interface{}) {
    self.Call("fadeComplete", args)
}

// Called automatically by SoundManager.volume.
// 
// Sets the volume of AudioTag Sounds as a percentage of the Global Volume.
// 
// You should not normally call this directly.
func (self *Sound) UpdateGlobalVolumeI(args ...interface{}) {
    self.Call("updateGlobalVolume", args)
}

// Destroys this sound and all associated events and removes it from the SoundManager.
func (self *Sound) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
