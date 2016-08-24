// Package phaser Automatic generation for Phaser.AudioSprite
// generated file AudioSprite.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// AudioSprite Audio Sprites are a combination of audio files and a JSON configuration.
// The JSON follows the format of that created by https://github.com/tonistiigi/audiosprite
type AudioSprite struct {
    *js.Object
}

// NewAudioSprite Audio Sprites are a combination of audio files and a JSON configuration.
// The JSON follows the format of that created by https://github.com/tonistiigi/audiosprite
func NewAudioSprite(game *Game, key string) *AudioSprite {
    return &AudioSprite{js.Global.Get("Phaser").Get("AudioSprite").New(game, key)}
}
// NewAudioSpriteI Audio Sprites are a combination of audio files and a JSON configuration.
// The JSON follows the format of that created by https://github.com/tonistiigi/audiosprite
func NewAudioSpriteI(args ...interface{}) *AudioSprite {
    return &AudioSprite{js.Global.Get("Phaser").Get("AudioSprite").New(args)}
}



// Game A reference to the currently running Game.
func (self *AudioSprite) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *AudioSprite) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Key Asset key for the Audio Sprite.
func (self *AudioSprite) Key() string{
    return self.Object.Get("key").String()
}

// SetKeyA Asset key for the Audio Sprite.
func (self *AudioSprite) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// Config JSON audio atlas object.
func (self *AudioSprite) Config() interface{}{
    return self.Object.Get("config")
}

// SetConfigA JSON audio atlas object.
func (self *AudioSprite) SetConfigA(member interface{}) {
    self.Object.Set("config", member)
}

// AutoplayKey If a sound is set to auto play, this holds the marker key of it.
func (self *AudioSprite) AutoplayKey() string{
    return self.Object.Get("autoplayKey").String()
}

// SetAutoplayKeyA If a sound is set to auto play, this holds the marker key of it.
func (self *AudioSprite) SetAutoplayKeyA(member string) {
    self.Object.Set("autoplayKey", member)
}

// Autoplay Is a sound set to autoplay or not?
func (self *AudioSprite) Autoplay() bool{
    return self.Object.Get("autoplay").Bool()
}

// SetAutoplayA Is a sound set to autoplay or not?
func (self *AudioSprite) SetAutoplayA(member bool) {
    self.Object.Set("autoplay", member)
}

// Sounds An object containing the Phaser.Sound objects for the Audio Sprite.
func (self *AudioSprite) Sounds() interface{}{
    return self.Object.Get("sounds")
}

// SetSoundsA An object containing the Phaser.Sound objects for the Audio Sprite.
func (self *AudioSprite) SetSoundsA(member interface{}) {
    self.Object.Set("sounds", member)
}


// Play Play a sound with the given name.
func (self *AudioSprite) Play() *Sound{
    return &Sound{self.Object.Call("play")}
}

// Play1O Play a sound with the given name.
func (self *AudioSprite) Play1O(marker string) *Sound{
    return &Sound{self.Object.Call("play", marker)}
}

// Play2O Play a sound with the given name.
func (self *AudioSprite) Play2O(marker string, volume int) *Sound{
    return &Sound{self.Object.Call("play", marker, volume)}
}

// PlayI Play a sound with the given name.
func (self *AudioSprite) PlayI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("play", args)}
}

// Stop Stop a sound with the given name.
func (self *AudioSprite) Stop() {
    self.Object.Call("stop")
}

// Stop1O Stop a sound with the given name.
func (self *AudioSprite) Stop1O(marker string) {
    self.Object.Call("stop", marker)
}

// StopI Stop a sound with the given name.
func (self *AudioSprite) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Get Get a sound with the given name.
func (self *AudioSprite) Get(marker string) *Sound{
    return &Sound{self.Object.Call("get", marker)}
}

// GetI Get a sound with the given name.
func (self *AudioSprite) GetI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("get", args)}
}

