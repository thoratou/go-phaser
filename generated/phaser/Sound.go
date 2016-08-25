// Package phaser Automatic generation for Phaser.Sound
// generated file Sound.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Sound The Sound class constructor.
type Sound struct {
    *js.Object
}

// NewSound The Sound class constructor.
func NewSound(game *Game, key string) *Sound {
    return &Sound{js.Global.Get("Phaser").Get("Sound").New(game, key)}
}
// NewSound1O The Sound class constructor.
func NewSound1O(game *Game, key string, volume int) *Sound {
    return &Sound{js.Global.Get("Phaser").Get("Sound").New(game, key, volume)}
}
// NewSound2O The Sound class constructor.
func NewSound2O(game *Game, key string, volume int, loop bool) *Sound {
    return &Sound{js.Global.Get("Phaser").Get("Sound").New(game, key, volume, loop)}
}
// NewSoundI The Sound class constructor.
func NewSoundI(args ...interface{}) *Sound {
    return &Sound{js.Global.Get("Phaser").Get("Sound").New(args)}
}



// Sound Binding conversion method to Sound point 
func ToSound(jsStruct interface{}) *Sound {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Sound{Object: object}
	}
	return nil
}



// Game A reference to the currently running Game.
func (self *Sound) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Sound) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Name Name of the sound.
func (self *Sound) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA Name of the sound.
func (self *Sound) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Key Asset key for the sound.
func (self *Sound) Key() string{
    return self.Object.Get("key").String()
}

// SetKeyA Asset key for the sound.
func (self *Sound) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// Loop Whether or not the sound or current sound marker will loop.
func (self *Sound) Loop() bool{
    return self.Object.Get("loop").Bool()
}

// SetLoopA Whether or not the sound or current sound marker will loop.
func (self *Sound) SetLoopA(member bool) {
    self.Object.Set("loop", member)
}

// Markers The sound markers.
func (self *Sound) Markers() interface{}{
    return self.Object.Get("markers")
}

// SetMarkersA The sound markers.
func (self *Sound) SetMarkersA(member interface{}) {
    self.Object.Set("markers", member)
}

// Context Reference to the AudioContext instance.
func (self *Sound) Context() *AudioContext{
    return &AudioContext{self.Object.Get("context")}
}

// SetContextA Reference to the AudioContext instance.
func (self *Sound) SetContextA(member *AudioContext) {
    self.Object.Set("context", member)
}

// Autoplay Boolean indicating whether the sound should start automatically.
func (self *Sound) Autoplay() bool{
    return self.Object.Get("autoplay").Bool()
}

// SetAutoplayA Boolean indicating whether the sound should start automatically.
func (self *Sound) SetAutoplayA(member bool) {
    self.Object.Set("autoplay", member)
}

// TotalDuration The total duration of the sound in seconds.
func (self *Sound) TotalDuration() int{
    return self.Object.Get("totalDuration").Int()
}

// SetTotalDurationA The total duration of the sound in seconds.
func (self *Sound) SetTotalDurationA(member int) {
    self.Object.Set("totalDuration", member)
}

// StartTime The time the Sound starts at (typically 0 unless starting from a marker)
func (self *Sound) StartTime() int{
    return self.Object.Get("startTime").Int()
}

// SetStartTimeA The time the Sound starts at (typically 0 unless starting from a marker)
func (self *Sound) SetStartTimeA(member int) {
    self.Object.Set("startTime", member)
}

// CurrentTime The current time the sound is at.
func (self *Sound) CurrentTime() int{
    return self.Object.Get("currentTime").Int()
}

// SetCurrentTimeA The current time the sound is at.
func (self *Sound) SetCurrentTimeA(member int) {
    self.Object.Set("currentTime", member)
}

// Duration The duration of the current sound marker in seconds.
func (self *Sound) Duration() int{
    return self.Object.Get("duration").Int()
}

// SetDurationA The duration of the current sound marker in seconds.
func (self *Sound) SetDurationA(member int) {
    self.Object.Set("duration", member)
}

// DurationMS The duration of the current sound marker in ms.
func (self *Sound) DurationMS() int{
    return self.Object.Get("durationMS").Int()
}

// SetDurationMSA The duration of the current sound marker in ms.
func (self *Sound) SetDurationMSA(member int) {
    self.Object.Set("durationMS", member)
}

// Position The position of the current sound marker.
func (self *Sound) Position() int{
    return self.Object.Get("position").Int()
}

// SetPositionA The position of the current sound marker.
func (self *Sound) SetPositionA(member int) {
    self.Object.Set("position", member)
}

// StopTime The time the sound stopped.
func (self *Sound) StopTime() int{
    return self.Object.Get("stopTime").Int()
}

// SetStopTimeA The time the sound stopped.
func (self *Sound) SetStopTimeA(member int) {
    self.Object.Set("stopTime", member)
}

// Paused true if the sound is paused, otherwise false.
func (self *Sound) Paused() bool{
    return self.Object.Get("paused").Bool()
}

// SetPausedA true if the sound is paused, otherwise false.
func (self *Sound) SetPausedA(member bool) {
    self.Object.Set("paused", member)
}

// PausedPosition The position the sound had reached when it was paused.
func (self *Sound) PausedPosition() int{
    return self.Object.Get("pausedPosition").Int()
}

// SetPausedPositionA The position the sound had reached when it was paused.
func (self *Sound) SetPausedPositionA(member int) {
    self.Object.Set("pausedPosition", member)
}

// PausedTime The game time at which the sound was paused.
func (self *Sound) PausedTime() int{
    return self.Object.Get("pausedTime").Int()
}

// SetPausedTimeA The game time at which the sound was paused.
func (self *Sound) SetPausedTimeA(member int) {
    self.Object.Set("pausedTime", member)
}

// IsPlaying true if the sound is currently playing, otherwise false.
func (self *Sound) IsPlaying() bool{
    return self.Object.Get("isPlaying").Bool()
}

// SetIsPlayingA true if the sound is currently playing, otherwise false.
func (self *Sound) SetIsPlayingA(member bool) {
    self.Object.Set("isPlaying", member)
}

// CurrentMarker The string ID of the currently playing marker, if any.
func (self *Sound) CurrentMarker() string{
    return self.Object.Get("currentMarker").String()
}

// SetCurrentMarkerA The string ID of the currently playing marker, if any.
func (self *Sound) SetCurrentMarkerA(member string) {
    self.Object.Set("currentMarker", member)
}

// FadeTween The tween that fades the audio, set via Sound.fadeIn and Sound.fadeOut.
func (self *Sound) FadeTween() *Tween{
    return &Tween{self.Object.Get("fadeTween")}
}

// SetFadeTweenA The tween that fades the audio, set via Sound.fadeIn and Sound.fadeOut.
func (self *Sound) SetFadeTweenA(member *Tween) {
    self.Object.Set("fadeTween", member)
}

// PendingPlayback true if the sound file is pending playback
func (self *Sound) PendingPlayback() bool{
    return self.Object.Get("pendingPlayback").Bool()
}

// SetPendingPlaybackA true if the sound file is pending playback
func (self *Sound) SetPendingPlaybackA(member bool) {
    self.Object.Set("pendingPlayback", member)
}

// Override if true when you play this sound it will always start from the beginning.
func (self *Sound) Override() bool{
    return self.Object.Get("override").Bool()
}

// SetOverrideA if true when you play this sound it will always start from the beginning.
func (self *Sound) SetOverrideA(member bool) {
    self.Object.Set("override", member)
}

// AllowMultiple This will allow you to have multiple instances of this Sound playing at once. This is only useful when running under Web Audio, and we recommend you implement a local pooling system to not flood the sound channels.
func (self *Sound) AllowMultiple() bool{
    return self.Object.Get("allowMultiple").Bool()
}

// SetAllowMultipleA This will allow you to have multiple instances of this Sound playing at once. This is only useful when running under Web Audio, and we recommend you implement a local pooling system to not flood the sound channels.
func (self *Sound) SetAllowMultipleA(member bool) {
    self.Object.Set("allowMultiple", member)
}

// UsingWebAudio true if this sound is being played with Web Audio.
func (self *Sound) UsingWebAudio() bool{
    return self.Object.Get("usingWebAudio").Bool()
}

// SetUsingWebAudioA true if this sound is being played with Web Audio.
func (self *Sound) SetUsingWebAudioA(member bool) {
    self.Object.Set("usingWebAudio", member)
}

// UsingAudioTag true if the sound is being played via the Audio tag.
func (self *Sound) UsingAudioTag() bool{
    return self.Object.Get("usingAudioTag").Bool()
}

// SetUsingAudioTagA true if the sound is being played via the Audio tag.
func (self *Sound) SetUsingAudioTagA(member bool) {
    self.Object.Set("usingAudioTag", member)
}

// ExternalNode If defined this Sound won't connect to the SoundManager master gain node, but will instead connect to externalNode.
func (self *Sound) ExternalNode() interface{}{
    return self.Object.Get("externalNode")
}

// SetExternalNodeA If defined this Sound won't connect to the SoundManager master gain node, but will instead connect to externalNode.
func (self *Sound) SetExternalNodeA(member interface{}) {
    self.Object.Set("externalNode", member)
}

// MasterGainNode The master gain node in a Web Audio system.
func (self *Sound) MasterGainNode() interface{}{
    return self.Object.Get("masterGainNode")
}

// SetMasterGainNodeA The master gain node in a Web Audio system.
func (self *Sound) SetMasterGainNodeA(member interface{}) {
    self.Object.Set("masterGainNode", member)
}

// GainNode The gain node in a Web Audio system.
func (self *Sound) GainNode() interface{}{
    return self.Object.Get("gainNode")
}

// SetGainNodeA The gain node in a Web Audio system.
func (self *Sound) SetGainNodeA(member interface{}) {
    self.Object.Set("gainNode", member)
}

// OnDecoded The onDecoded event is dispatched when the sound has finished decoding (typically for mp3 files)
func (self *Sound) OnDecoded() *Signal{
    return &Signal{self.Object.Get("onDecoded")}
}

// SetOnDecodedA The onDecoded event is dispatched when the sound has finished decoding (typically for mp3 files)
func (self *Sound) SetOnDecodedA(member *Signal) {
    self.Object.Set("onDecoded", member)
}

// OnPlay The onPlay event is dispatched each time this sound is played.
func (self *Sound) OnPlay() *Signal{
    return &Signal{self.Object.Get("onPlay")}
}

// SetOnPlayA The onPlay event is dispatched each time this sound is played.
func (self *Sound) SetOnPlayA(member *Signal) {
    self.Object.Set("onPlay", member)
}

// OnPause The onPause event is dispatched when this sound is paused.
func (self *Sound) OnPause() *Signal{
    return &Signal{self.Object.Get("onPause")}
}

// SetOnPauseA The onPause event is dispatched when this sound is paused.
func (self *Sound) SetOnPauseA(member *Signal) {
    self.Object.Set("onPause", member)
}

// OnResume The onResume event is dispatched when this sound is resumed from a paused state.
func (self *Sound) OnResume() *Signal{
    return &Signal{self.Object.Get("onResume")}
}

// SetOnResumeA The onResume event is dispatched when this sound is resumed from a paused state.
func (self *Sound) SetOnResumeA(member *Signal) {
    self.Object.Set("onResume", member)
}

// OnLoop The onLoop event is dispatched when this sound loops during playback.
func (self *Sound) OnLoop() *Signal{
    return &Signal{self.Object.Get("onLoop")}
}

// SetOnLoopA The onLoop event is dispatched when this sound loops during playback.
func (self *Sound) SetOnLoopA(member *Signal) {
    self.Object.Set("onLoop", member)
}

// OnStop The onStop event is dispatched when this sound stops playback.
func (self *Sound) OnStop() *Signal{
    return &Signal{self.Object.Get("onStop")}
}

// SetOnStopA The onStop event is dispatched when this sound stops playback.
func (self *Sound) SetOnStopA(member *Signal) {
    self.Object.Set("onStop", member)
}

// OnMute The onMute event is dispatched when this sound is muted.
func (self *Sound) OnMute() *Signal{
    return &Signal{self.Object.Get("onMute")}
}

// SetOnMuteA The onMute event is dispatched when this sound is muted.
func (self *Sound) SetOnMuteA(member *Signal) {
    self.Object.Set("onMute", member)
}

// OnMarkerComplete The onMarkerComplete event is dispatched when a marker within this sound completes playback.
func (self *Sound) OnMarkerComplete() *Signal{
    return &Signal{self.Object.Get("onMarkerComplete")}
}

// SetOnMarkerCompleteA The onMarkerComplete event is dispatched when a marker within this sound completes playback.
func (self *Sound) SetOnMarkerCompleteA(member *Signal) {
    self.Object.Set("onMarkerComplete", member)
}

// OnFadeComplete The onFadeComplete event is dispatched when this sound finishes fading either in or out.
func (self *Sound) OnFadeComplete() *Signal{
    return &Signal{self.Object.Get("onFadeComplete")}
}

// SetOnFadeCompleteA The onFadeComplete event is dispatched when this sound finishes fading either in or out.
func (self *Sound) SetOnFadeCompleteA(member *Signal) {
    self.Object.Set("onFadeComplete", member)
}

// IsDecoding Returns true if the sound file is still decoding.
func (self *Sound) IsDecoding() bool{
    return self.Object.Get("isDecoding").Bool()
}

// SetIsDecodingA Returns true if the sound file is still decoding.
func (self *Sound) SetIsDecodingA(member bool) {
    self.Object.Set("isDecoding", member)
}

// IsDecoded Returns true if the sound file has decoded.
func (self *Sound) IsDecoded() bool{
    return self.Object.Get("isDecoded").Bool()
}

// SetIsDecodedA Returns true if the sound file has decoded.
func (self *Sound) SetIsDecodedA(member bool) {
    self.Object.Set("isDecoded", member)
}

// Mute Gets or sets the muted state of this sound.
func (self *Sound) Mute() bool{
    return self.Object.Get("mute").Bool()
}

// SetMuteA Gets or sets the muted state of this sound.
func (self *Sound) SetMuteA(member bool) {
    self.Object.Set("mute", member)
}

// Volume Gets or sets the volume of this sound, a value between 0 and 1.
func (self *Sound) Volume() int{
    return self.Object.Get("volume").Int()
}

// SetVolumeA Gets or sets the volume of this sound, a value between 0 and 1.
func (self *Sound) SetVolumeA(member int) {
    self.Object.Set("volume", member)
}


// SoundHasUnlocked Called automatically when this sound is unlocked.
func (self *Sound) SoundHasUnlocked(key string) {
    self.Object.Call("soundHasUnlocked", key)
}

// SoundHasUnlockedI Called automatically when this sound is unlocked.
func (self *Sound) SoundHasUnlockedI(args ...interface{}) {
    self.Object.Call("soundHasUnlocked", args)
}

// AddMarker Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarker(name string, start int) {
    self.Object.Call("addMarker", name, start)
}

// AddMarker1O Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarker1O(name string, start int, duration int) {
    self.Object.Call("addMarker", name, start, duration)
}

// AddMarker2O Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarker2O(name string, start int, duration int, volume int) {
    self.Object.Call("addMarker", name, start, duration, volume)
}

// AddMarker3O Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarker3O(name string, start int, duration int, volume int, loop bool) {
    self.Object.Call("addMarker", name, start, duration, volume, loop)
}

// AddMarkerI Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarkerI(args ...interface{}) {
    self.Object.Call("addMarker", args)
}

// RemoveMarker Removes a marker from the sound.
func (self *Sound) RemoveMarker(name string) {
    self.Object.Call("removeMarker", name)
}

// RemoveMarkerI Removes a marker from the sound.
func (self *Sound) RemoveMarkerI(args ...interface{}) {
    self.Object.Call("removeMarker", args)
}

// OnEndedHandler Called automatically by the AudioContext when the sound stops playing.
// Doesn't get called if the sound is set to loop or is a section of an Audio Sprite.
func (self *Sound) OnEndedHandler() {
    self.Object.Call("onEndedHandler")
}

// OnEndedHandlerI Called automatically by the AudioContext when the sound stops playing.
// Doesn't get called if the sound is set to loop or is a section of an Audio Sprite.
func (self *Sound) OnEndedHandlerI(args ...interface{}) {
    self.Object.Call("onEndedHandler", args)
}

// Update Called automatically by Phaser.SoundManager.
func (self *Sound) Update() {
    self.Object.Call("update")
}

// UpdateI Called automatically by Phaser.SoundManager.
func (self *Sound) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// LoopFull Loops this entire sound. If you need to loop a section of it then use Sound.play and the marker and loop parameters.
func (self *Sound) LoopFull() *Sound{
    return &Sound{self.Object.Call("loopFull")}
}

// LoopFull1O Loops this entire sound. If you need to loop a section of it then use Sound.play and the marker and loop parameters.
func (self *Sound) LoopFull1O(volume int) *Sound{
    return &Sound{self.Object.Call("loopFull", volume)}
}

// LoopFullI Loops this entire sound. If you need to loop a section of it then use Sound.play and the marker and loop parameters.
func (self *Sound) LoopFullI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("loopFull", args)}
}

// Play Play this sound, or a marked section of it.
func (self *Sound) Play() *Sound{
    return &Sound{self.Object.Call("play")}
}

// Play1O Play this sound, or a marked section of it.
func (self *Sound) Play1O(marker string) *Sound{
    return &Sound{self.Object.Call("play", marker)}
}

// Play2O Play this sound, or a marked section of it.
func (self *Sound) Play2O(marker string, position int) *Sound{
    return &Sound{self.Object.Call("play", marker, position)}
}

// Play3O Play this sound, or a marked section of it.
func (self *Sound) Play3O(marker string, position int, volume int) *Sound{
    return &Sound{self.Object.Call("play", marker, position, volume)}
}

// Play4O Play this sound, or a marked section of it.
func (self *Sound) Play4O(marker string, position int, volume int, loop bool) *Sound{
    return &Sound{self.Object.Call("play", marker, position, volume, loop)}
}

// Play5O Play this sound, or a marked section of it.
func (self *Sound) Play5O(marker string, position int, volume int, loop bool, forceRestart bool) *Sound{
    return &Sound{self.Object.Call("play", marker, position, volume, loop, forceRestart)}
}

// PlayI Play this sound, or a marked section of it.
func (self *Sound) PlayI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("play", args)}
}

// Restart Restart the sound, or a marked section of it.
func (self *Sound) Restart() {
    self.Object.Call("restart")
}

// Restart1O Restart the sound, or a marked section of it.
func (self *Sound) Restart1O(marker string) {
    self.Object.Call("restart", marker)
}

// Restart2O Restart the sound, or a marked section of it.
func (self *Sound) Restart2O(marker string, position int) {
    self.Object.Call("restart", marker, position)
}

// Restart3O Restart the sound, or a marked section of it.
func (self *Sound) Restart3O(marker string, position int, volume int) {
    self.Object.Call("restart", marker, position, volume)
}

// Restart4O Restart the sound, or a marked section of it.
func (self *Sound) Restart4O(marker string, position int, volume int, loop bool) {
    self.Object.Call("restart", marker, position, volume, loop)
}

// RestartI Restart the sound, or a marked section of it.
func (self *Sound) RestartI(args ...interface{}) {
    self.Object.Call("restart", args)
}

// Pause Pauses the sound.
func (self *Sound) Pause() {
    self.Object.Call("pause")
}

// PauseI Pauses the sound.
func (self *Sound) PauseI(args ...interface{}) {
    self.Object.Call("pause", args)
}

// Resume Resumes the sound.
func (self *Sound) Resume() {
    self.Object.Call("resume")
}

// ResumeI Resumes the sound.
func (self *Sound) ResumeI(args ...interface{}) {
    self.Object.Call("resume", args)
}

// Stop Stop playing this sound.
func (self *Sound) Stop() {
    self.Object.Call("stop")
}

// StopI Stop playing this sound.
func (self *Sound) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// FadeIn Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeIn() {
    self.Object.Call("fadeIn")
}

// FadeIn1O Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeIn1O(duration int) {
    self.Object.Call("fadeIn", duration)
}

// FadeIn2O Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeIn2O(duration int, loop bool) {
    self.Object.Call("fadeIn", duration, loop)
}

// FadeIn3O Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeIn3O(duration int, loop bool, marker string) {
    self.Object.Call("fadeIn", duration, loop, marker)
}

// FadeInI Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeInI(args ...interface{}) {
    self.Object.Call("fadeIn", args)
}

// FadeOut Decreases the volume of this Sound from its current value to 0 over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (0) as the second parameter.
func (self *Sound) FadeOut() {
    self.Object.Call("fadeOut")
}

// FadeOut1O Decreases the volume of this Sound from its current value to 0 over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (0) as the second parameter.
func (self *Sound) FadeOut1O(duration int) {
    self.Object.Call("fadeOut", duration)
}

// FadeOutI Decreases the volume of this Sound from its current value to 0 over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (0) as the second parameter.
func (self *Sound) FadeOutI(args ...interface{}) {
    self.Object.Call("fadeOut", args)
}

// FadeTo Fades the volume of this Sound from its current value to the given volume over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter, 
// and the final volume (volume) as the second parameter.
func (self *Sound) FadeTo() {
    self.Object.Call("fadeTo")
}

// FadeTo1O Fades the volume of this Sound from its current value to the given volume over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter, 
// and the final volume (volume) as the second parameter.
func (self *Sound) FadeTo1O(duration int) {
    self.Object.Call("fadeTo", duration)
}

// FadeTo2O Fades the volume of this Sound from its current value to the given volume over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter, 
// and the final volume (volume) as the second parameter.
func (self *Sound) FadeTo2O(duration int, volume int) {
    self.Object.Call("fadeTo", duration, volume)
}

// FadeToI Fades the volume of this Sound from its current value to the given volume over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter, 
// and the final volume (volume) as the second parameter.
func (self *Sound) FadeToI(args ...interface{}) {
    self.Object.Call("fadeTo", args)
}

// FadeComplete Internal handler for Sound.fadeIn, Sound.fadeOut and Sound.fadeTo.
func (self *Sound) FadeComplete() {
    self.Object.Call("fadeComplete")
}

// FadeCompleteI Internal handler for Sound.fadeIn, Sound.fadeOut and Sound.fadeTo.
func (self *Sound) FadeCompleteI(args ...interface{}) {
    self.Object.Call("fadeComplete", args)
}

// UpdateGlobalVolume Called automatically by SoundManager.volume.
// 
// Sets the volume of AudioTag Sounds as a percentage of the Global Volume.
// 
// You should not normally call this directly.
func (self *Sound) UpdateGlobalVolume(globalVolume float64) {
    self.Object.Call("updateGlobalVolume", globalVolume)
}

// UpdateGlobalVolumeI Called automatically by SoundManager.volume.
// 
// Sets the volume of AudioTag Sounds as a percentage of the Global Volume.
// 
// You should not normally call this directly.
func (self *Sound) UpdateGlobalVolumeI(args ...interface{}) {
    self.Object.Call("updateGlobalVolume", args)
}

// Destroy Destroys this sound and all associated events and removes it from the SoundManager.
func (self *Sound) Destroy() {
    self.Object.Call("destroy")
}

// Destroy1O Destroys this sound and all associated events and removes it from the SoundManager.
func (self *Sound) Destroy1O(remove bool) {
    self.Object.Call("destroy", remove)
}

// DestroyI Destroys this sound and all associated events and removes it from the SoundManager.
func (self *Sound) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

