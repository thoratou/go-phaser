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


// The Sound class constructor.
func NewSound(game *Game, key string) *Sound {
    return &Sound{js.Global.Get("Phaser").Get("Sound").New(game, key)}
}

// The Sound class constructor.
func NewSound1O(game *Game, key string, volume int) *Sound {
    return &Sound{js.Global.Get("Phaser").Get("Sound").New(game, key, volume)}
}

// The Sound class constructor.
func NewSound2O(game *Game, key string, volume int, loop bool) *Sound {
    return &Sound{js.Global.Get("Phaser").Get("Sound").New(game, key, volume, loop)}
}

// The Sound class constructor.
func NewSoundI(args ...interface{}) *Sound {
    return &Sound{js.Global.Get("Phaser").Get("Sound").New(args)}
}



// A reference to the currently running Game.
func (self *Sound) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *Sound) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Name of the sound.
func (self *Sound) Name() string{
    return self.Object.Get("name").String()
}

// Name of the sound.
func (self *Sound) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Asset key for the sound.
func (self *Sound) Key() string{
    return self.Object.Get("key").String()
}

// Asset key for the sound.
func (self *Sound) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// Whether or not the sound or current sound marker will loop.
func (self *Sound) Loop() bool{
    return self.Object.Get("loop").Bool()
}

// Whether or not the sound or current sound marker will loop.
func (self *Sound) SetLoopA(member bool) {
    self.Object.Set("loop", member)
}

// The sound markers.
func (self *Sound) Markers() interface{}{
    return self.Object.Get("markers")
}

// The sound markers.
func (self *Sound) SetMarkersA(member interface{}) {
    self.Object.Set("markers", member)
}

// Reference to the AudioContext instance.
func (self *Sound) Context() *AudioContext{
    return &AudioContext{self.Object.Get("context")}
}

// Reference to the AudioContext instance.
func (self *Sound) SetContextA(member *AudioContext) {
    self.Object.Set("context", member)
}

// Boolean indicating whether the sound should start automatically.
func (self *Sound) Autoplay() bool{
    return self.Object.Get("autoplay").Bool()
}

// Boolean indicating whether the sound should start automatically.
func (self *Sound) SetAutoplayA(member bool) {
    self.Object.Set("autoplay", member)
}

// The total duration of the sound in seconds.
func (self *Sound) TotalDuration() int{
    return self.Object.Get("totalDuration").Int()
}

// The total duration of the sound in seconds.
func (self *Sound) SetTotalDurationA(member int) {
    self.Object.Set("totalDuration", member)
}

// The time the Sound starts at (typically 0 unless starting from a marker)
func (self *Sound) StartTime() int{
    return self.Object.Get("startTime").Int()
}

// The time the Sound starts at (typically 0 unless starting from a marker)
func (self *Sound) SetStartTimeA(member int) {
    self.Object.Set("startTime", member)
}

// The current time the sound is at.
func (self *Sound) CurrentTime() int{
    return self.Object.Get("currentTime").Int()
}

// The current time the sound is at.
func (self *Sound) SetCurrentTimeA(member int) {
    self.Object.Set("currentTime", member)
}

// The duration of the current sound marker in seconds.
func (self *Sound) Duration() int{
    return self.Object.Get("duration").Int()
}

// The duration of the current sound marker in seconds.
func (self *Sound) SetDurationA(member int) {
    self.Object.Set("duration", member)
}

// The duration of the current sound marker in ms.
func (self *Sound) DurationMS() int{
    return self.Object.Get("durationMS").Int()
}

// The duration of the current sound marker in ms.
func (self *Sound) SetDurationMSA(member int) {
    self.Object.Set("durationMS", member)
}

// The position of the current sound marker.
func (self *Sound) Position() int{
    return self.Object.Get("position").Int()
}

// The position of the current sound marker.
func (self *Sound) SetPositionA(member int) {
    self.Object.Set("position", member)
}

// The time the sound stopped.
func (self *Sound) StopTime() int{
    return self.Object.Get("stopTime").Int()
}

// The time the sound stopped.
func (self *Sound) SetStopTimeA(member int) {
    self.Object.Set("stopTime", member)
}

// true if the sound is paused, otherwise false.
func (self *Sound) Paused() bool{
    return self.Object.Get("paused").Bool()
}

// true if the sound is paused, otherwise false.
func (self *Sound) SetPausedA(member bool) {
    self.Object.Set("paused", member)
}

// The position the sound had reached when it was paused.
func (self *Sound) PausedPosition() int{
    return self.Object.Get("pausedPosition").Int()
}

// The position the sound had reached when it was paused.
func (self *Sound) SetPausedPositionA(member int) {
    self.Object.Set("pausedPosition", member)
}

// The game time at which the sound was paused.
func (self *Sound) PausedTime() int{
    return self.Object.Get("pausedTime").Int()
}

// The game time at which the sound was paused.
func (self *Sound) SetPausedTimeA(member int) {
    self.Object.Set("pausedTime", member)
}

// true if the sound is currently playing, otherwise false.
func (self *Sound) IsPlaying() bool{
    return self.Object.Get("isPlaying").Bool()
}

// true if the sound is currently playing, otherwise false.
func (self *Sound) SetIsPlayingA(member bool) {
    self.Object.Set("isPlaying", member)
}

// The string ID of the currently playing marker, if any.
func (self *Sound) CurrentMarker() string{
    return self.Object.Get("currentMarker").String()
}

// The string ID of the currently playing marker, if any.
func (self *Sound) SetCurrentMarkerA(member string) {
    self.Object.Set("currentMarker", member)
}

// The tween that fades the audio, set via Sound.fadeIn and Sound.fadeOut.
func (self *Sound) FadeTween() *Tween{
    return &Tween{self.Object.Get("fadeTween")}
}

// The tween that fades the audio, set via Sound.fadeIn and Sound.fadeOut.
func (self *Sound) SetFadeTweenA(member *Tween) {
    self.Object.Set("fadeTween", member)
}

// true if the sound file is pending playback
func (self *Sound) PendingPlayback() bool{
    return self.Object.Get("pendingPlayback").Bool()
}

// true if the sound file is pending playback
func (self *Sound) SetPendingPlaybackA(member bool) {
    self.Object.Set("pendingPlayback", member)
}

// if true when you play this sound it will always start from the beginning.
func (self *Sound) Override() bool{
    return self.Object.Get("override").Bool()
}

// if true when you play this sound it will always start from the beginning.
func (self *Sound) SetOverrideA(member bool) {
    self.Object.Set("override", member)
}

// This will allow you to have multiple instances of this Sound playing at once. This is only useful when running under Web Audio, and we recommend you implement a local pooling system to not flood the sound channels.
func (self *Sound) AllowMultiple() bool{
    return self.Object.Get("allowMultiple").Bool()
}

// This will allow you to have multiple instances of this Sound playing at once. This is only useful when running under Web Audio, and we recommend you implement a local pooling system to not flood the sound channels.
func (self *Sound) SetAllowMultipleA(member bool) {
    self.Object.Set("allowMultiple", member)
}

// true if this sound is being played with Web Audio.
func (self *Sound) UsingWebAudio() bool{
    return self.Object.Get("usingWebAudio").Bool()
}

// true if this sound is being played with Web Audio.
func (self *Sound) SetUsingWebAudioA(member bool) {
    self.Object.Set("usingWebAudio", member)
}

// true if the sound is being played via the Audio tag.
func (self *Sound) UsingAudioTag() bool{
    return self.Object.Get("usingAudioTag").Bool()
}

// true if the sound is being played via the Audio tag.
func (self *Sound) SetUsingAudioTagA(member bool) {
    self.Object.Set("usingAudioTag", member)
}

// If defined this Sound won't connect to the SoundManager master gain node, but will instead connect to externalNode.
func (self *Sound) ExternalNode() interface{}{
    return self.Object.Get("externalNode")
}

// If defined this Sound won't connect to the SoundManager master gain node, but will instead connect to externalNode.
func (self *Sound) SetExternalNodeA(member interface{}) {
    self.Object.Set("externalNode", member)
}

// The master gain node in a Web Audio system.
func (self *Sound) MasterGainNode() interface{}{
    return self.Object.Get("masterGainNode")
}

// The master gain node in a Web Audio system.
func (self *Sound) SetMasterGainNodeA(member interface{}) {
    self.Object.Set("masterGainNode", member)
}

// The gain node in a Web Audio system.
func (self *Sound) GainNode() interface{}{
    return self.Object.Get("gainNode")
}

// The gain node in a Web Audio system.
func (self *Sound) SetGainNodeA(member interface{}) {
    self.Object.Set("gainNode", member)
}

// The onDecoded event is dispatched when the sound has finished decoding (typically for mp3 files)
func (self *Sound) OnDecoded() *Signal{
    return &Signal{self.Object.Get("onDecoded")}
}

// The onDecoded event is dispatched when the sound has finished decoding (typically for mp3 files)
func (self *Sound) SetOnDecodedA(member *Signal) {
    self.Object.Set("onDecoded", member)
}

// The onPlay event is dispatched each time this sound is played.
func (self *Sound) OnPlay() *Signal{
    return &Signal{self.Object.Get("onPlay")}
}

// The onPlay event is dispatched each time this sound is played.
func (self *Sound) SetOnPlayA(member *Signal) {
    self.Object.Set("onPlay", member)
}

// The onPause event is dispatched when this sound is paused.
func (self *Sound) OnPause() *Signal{
    return &Signal{self.Object.Get("onPause")}
}

// The onPause event is dispatched when this sound is paused.
func (self *Sound) SetOnPauseA(member *Signal) {
    self.Object.Set("onPause", member)
}

// The onResume event is dispatched when this sound is resumed from a paused state.
func (self *Sound) OnResume() *Signal{
    return &Signal{self.Object.Get("onResume")}
}

// The onResume event is dispatched when this sound is resumed from a paused state.
func (self *Sound) SetOnResumeA(member *Signal) {
    self.Object.Set("onResume", member)
}

// The onLoop event is dispatched when this sound loops during playback.
func (self *Sound) OnLoop() *Signal{
    return &Signal{self.Object.Get("onLoop")}
}

// The onLoop event is dispatched when this sound loops during playback.
func (self *Sound) SetOnLoopA(member *Signal) {
    self.Object.Set("onLoop", member)
}

// The onStop event is dispatched when this sound stops playback.
func (self *Sound) OnStop() *Signal{
    return &Signal{self.Object.Get("onStop")}
}

// The onStop event is dispatched when this sound stops playback.
func (self *Sound) SetOnStopA(member *Signal) {
    self.Object.Set("onStop", member)
}

// The onMute event is dispatched when this sound is muted.
func (self *Sound) OnMute() *Signal{
    return &Signal{self.Object.Get("onMute")}
}

// The onMute event is dispatched when this sound is muted.
func (self *Sound) SetOnMuteA(member *Signal) {
    self.Object.Set("onMute", member)
}

// The onMarkerComplete event is dispatched when a marker within this sound completes playback.
func (self *Sound) OnMarkerComplete() *Signal{
    return &Signal{self.Object.Get("onMarkerComplete")}
}

// The onMarkerComplete event is dispatched when a marker within this sound completes playback.
func (self *Sound) SetOnMarkerCompleteA(member *Signal) {
    self.Object.Set("onMarkerComplete", member)
}

// The onFadeComplete event is dispatched when this sound finishes fading either in or out.
func (self *Sound) OnFadeComplete() *Signal{
    return &Signal{self.Object.Get("onFadeComplete")}
}

// The onFadeComplete event is dispatched when this sound finishes fading either in or out.
func (self *Sound) SetOnFadeCompleteA(member *Signal) {
    self.Object.Set("onFadeComplete", member)
}

// Returns true if the sound file is still decoding.
func (self *Sound) IsDecoding() bool{
    return self.Object.Get("isDecoding").Bool()
}

// Returns true if the sound file is still decoding.
func (self *Sound) SetIsDecodingA(member bool) {
    self.Object.Set("isDecoding", member)
}

// Returns true if the sound file has decoded.
func (self *Sound) IsDecoded() bool{
    return self.Object.Get("isDecoded").Bool()
}

// Returns true if the sound file has decoded.
func (self *Sound) SetIsDecodedA(member bool) {
    self.Object.Set("isDecoded", member)
}

// Gets or sets the muted state of this sound.
func (self *Sound) Mute() bool{
    return self.Object.Get("mute").Bool()
}

// Gets or sets the muted state of this sound.
func (self *Sound) SetMuteA(member bool) {
    self.Object.Set("mute", member)
}

// Gets or sets the volume of this sound, a value between 0 and 1.
func (self *Sound) Volume() int{
    return self.Object.Get("volume").Int()
}

// Gets or sets the volume of this sound, a value between 0 and 1.
func (self *Sound) SetVolumeA(member int) {
    self.Object.Set("volume", member)
}



// Called automatically when this sound is unlocked.
func (self *Sound) SoundHasUnlocked(key string) {
    self.Object.Call("soundHasUnlocked", key)
}

// Called automatically when this sound is unlocked.
func (self *Sound) SoundHasUnlockedI(args ...interface{}) {
    self.Object.Call("soundHasUnlocked", args)
}

// Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarker(name string, start int) {
    self.Object.Call("addMarker", name, start)
}

// Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarker1O(name string, start int, duration int) {
    self.Object.Call("addMarker", name, start, duration)
}

// Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarker2O(name string, start int, duration int, volume int) {
    self.Object.Call("addMarker", name, start, duration, volume)
}

// Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarker3O(name string, start int, duration int, volume int, loop bool) {
    self.Object.Call("addMarker", name, start, duration, volume, loop)
}

// Adds a marker into the current Sound. A marker is represented by a unique key and a start time and duration.
// This allows you to bundle multiple sounds together into a single audio file and use markers to jump between them for playback.
func (self *Sound) AddMarkerI(args ...interface{}) {
    self.Object.Call("addMarker", args)
}

// Removes a marker from the sound.
func (self *Sound) RemoveMarker(name string) {
    self.Object.Call("removeMarker", name)
}

// Removes a marker from the sound.
func (self *Sound) RemoveMarkerI(args ...interface{}) {
    self.Object.Call("removeMarker", args)
}

// Called automatically by the AudioContext when the sound stops playing.
// Doesn't get called if the sound is set to loop or is a section of an Audio Sprite.
func (self *Sound) OnEndedHandler() {
    self.Object.Call("onEndedHandler")
}

// Called automatically by the AudioContext when the sound stops playing.
// Doesn't get called if the sound is set to loop or is a section of an Audio Sprite.
func (self *Sound) OnEndedHandlerI(args ...interface{}) {
    self.Object.Call("onEndedHandler", args)
}

// Called automatically by Phaser.SoundManager.
func (self *Sound) Update() {
    self.Object.Call("update")
}

// Called automatically by Phaser.SoundManager.
func (self *Sound) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Loops this entire sound. If you need to loop a section of it then use Sound.play and the marker and loop parameters.
func (self *Sound) LoopFull() *Sound{
    return &Sound{self.Object.Call("loopFull")}
}

// Loops this entire sound. If you need to loop a section of it then use Sound.play and the marker and loop parameters.
func (self *Sound) LoopFull1O(volume int) *Sound{
    return &Sound{self.Object.Call("loopFull", volume)}
}

// Loops this entire sound. If you need to loop a section of it then use Sound.play and the marker and loop parameters.
func (self *Sound) LoopFullI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("loopFull", args)}
}

// Play this sound, or a marked section of it.
func (self *Sound) Play() *Sound{
    return &Sound{self.Object.Call("play")}
}

// Play this sound, or a marked section of it.
func (self *Sound) Play1O(marker string) *Sound{
    return &Sound{self.Object.Call("play", marker)}
}

// Play this sound, or a marked section of it.
func (self *Sound) Play2O(marker string, position int) *Sound{
    return &Sound{self.Object.Call("play", marker, position)}
}

// Play this sound, or a marked section of it.
func (self *Sound) Play3O(marker string, position int, volume int) *Sound{
    return &Sound{self.Object.Call("play", marker, position, volume)}
}

// Play this sound, or a marked section of it.
func (self *Sound) Play4O(marker string, position int, volume int, loop bool) *Sound{
    return &Sound{self.Object.Call("play", marker, position, volume, loop)}
}

// Play this sound, or a marked section of it.
func (self *Sound) Play5O(marker string, position int, volume int, loop bool, forceRestart bool) *Sound{
    return &Sound{self.Object.Call("play", marker, position, volume, loop, forceRestart)}
}

// Play this sound, or a marked section of it.
func (self *Sound) PlayI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("play", args)}
}

// Restart the sound, or a marked section of it.
func (self *Sound) Restart() {
    self.Object.Call("restart")
}

// Restart the sound, or a marked section of it.
func (self *Sound) Restart1O(marker string) {
    self.Object.Call("restart", marker)
}

// Restart the sound, or a marked section of it.
func (self *Sound) Restart2O(marker string, position int) {
    self.Object.Call("restart", marker, position)
}

// Restart the sound, or a marked section of it.
func (self *Sound) Restart3O(marker string, position int, volume int) {
    self.Object.Call("restart", marker, position, volume)
}

// Restart the sound, or a marked section of it.
func (self *Sound) Restart4O(marker string, position int, volume int, loop bool) {
    self.Object.Call("restart", marker, position, volume, loop)
}

// Restart the sound, or a marked section of it.
func (self *Sound) RestartI(args ...interface{}) {
    self.Object.Call("restart", args)
}

// Pauses the sound.
func (self *Sound) Pause() {
    self.Object.Call("pause")
}

// Pauses the sound.
func (self *Sound) PauseI(args ...interface{}) {
    self.Object.Call("pause", args)
}

// Resumes the sound.
func (self *Sound) Resume() {
    self.Object.Call("resume")
}

// Resumes the sound.
func (self *Sound) ResumeI(args ...interface{}) {
    self.Object.Call("resume", args)
}

// Stop playing this sound.
func (self *Sound) Stop() {
    self.Object.Call("stop")
}

// Stop playing this sound.
func (self *Sound) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeIn() {
    self.Object.Call("fadeIn")
}

// Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeIn1O(duration int) {
    self.Object.Call("fadeIn", duration)
}

// Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeIn2O(duration int, loop bool) {
    self.Object.Call("fadeIn", duration, loop)
}

// Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeIn3O(duration int, loop bool, marker string) {
    self.Object.Call("fadeIn", duration, loop, marker)
}

// Starts this sound playing (or restarts it if already doing so) and sets the volume to zero.
// Then increases the volume from 0 to 1 over the duration specified.
// 
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (1) as the second parameter.
func (self *Sound) FadeInI(args ...interface{}) {
    self.Object.Call("fadeIn", args)
}

// Decreases the volume of this Sound from its current value to 0 over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (0) as the second parameter.
func (self *Sound) FadeOut() {
    self.Object.Call("fadeOut")
}

// Decreases the volume of this Sound from its current value to 0 over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (0) as the second parameter.
func (self *Sound) FadeOut1O(duration int) {
    self.Object.Call("fadeOut", duration)
}

// Decreases the volume of this Sound from its current value to 0 over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter,
// and the final volume (0) as the second parameter.
func (self *Sound) FadeOutI(args ...interface{}) {
    self.Object.Call("fadeOut", args)
}

// Fades the volume of this Sound from its current value to the given volume over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter, 
// and the final volume (volume) as the second parameter.
func (self *Sound) FadeTo() {
    self.Object.Call("fadeTo")
}

// Fades the volume of this Sound from its current value to the given volume over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter, 
// and the final volume (volume) as the second parameter.
func (self *Sound) FadeTo1O(duration int) {
    self.Object.Call("fadeTo", duration)
}

// Fades the volume of this Sound from its current value to the given volume over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter, 
// and the final volume (volume) as the second parameter.
func (self *Sound) FadeTo2O(duration int, volume int) {
    self.Object.Call("fadeTo", duration, volume)
}

// Fades the volume of this Sound from its current value to the given volume over the duration specified.
// At the end of the fade Sound.onFadeComplete is dispatched with this Sound object as the first parameter, 
// and the final volume (volume) as the second parameter.
func (self *Sound) FadeToI(args ...interface{}) {
    self.Object.Call("fadeTo", args)
}

// Internal handler for Sound.fadeIn, Sound.fadeOut and Sound.fadeTo.
func (self *Sound) FadeComplete() {
    self.Object.Call("fadeComplete")
}

// Internal handler for Sound.fadeIn, Sound.fadeOut and Sound.fadeTo.
func (self *Sound) FadeCompleteI(args ...interface{}) {
    self.Object.Call("fadeComplete", args)
}

// Called automatically by SoundManager.volume.
// 
// Sets the volume of AudioTag Sounds as a percentage of the Global Volume.
// 
// You should not normally call this directly.
func (self *Sound) UpdateGlobalVolume(globalVolume float64) {
    self.Object.Call("updateGlobalVolume", globalVolume)
}

// Called automatically by SoundManager.volume.
// 
// Sets the volume of AudioTag Sounds as a percentage of the Global Volume.
// 
// You should not normally call this directly.
func (self *Sound) UpdateGlobalVolumeI(args ...interface{}) {
    self.Object.Call("updateGlobalVolume", args)
}

// Destroys this sound and all associated events and removes it from the SoundManager.
func (self *Sound) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this sound and all associated events and removes it from the SoundManager.
func (self *Sound) Destroy1O(remove bool) {
    self.Object.Call("destroy", remove)
}

// Destroys this sound and all associated events and removes it from the SoundManager.
func (self *Sound) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
