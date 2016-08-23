// Automatic generation for Phaser.AudioSprite
// generated file AudioSprite.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Audio Sprites are a combination of audio files and a JSON configuration.
// The JSON follows the format of that created by https://github.com/tonistiigi/audiosprite
type AudioSprite struct {
    *js.Object
}


// Audio Sprites are a combination of audio files and a JSON configuration.
// The JSON follows the format of that created by https://github.com/tonistiigi/audiosprite
func NewAudioSprite(game *Game, key string) *AudioSprite {
    return &AudioSprite{js.Global.Get("Phaser").Get("AudioSprite").New(game, key)}
}

// Audio Sprites are a combination of audio files and a JSON configuration.
// The JSON follows the format of that created by https://github.com/tonistiigi/audiosprite
func NewAudioSpriteI(args ...interface{}) *AudioSprite {
    return &AudioSprite{js.Global.Get("Phaser").Get("AudioSprite").New(args)}
}



// A reference to the currently running Game.
func (self *AudioSprite) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *AudioSprite) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Asset key for the Audio Sprite.
func (self *AudioSprite) Key() string{
    return self.Object.Get("key").String()
}

// Asset key for the Audio Sprite.
func (self *AudioSprite) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// JSON audio atlas object.
func (self *AudioSprite) Config() interface{}{
    return self.Object.Get("config")
}

// JSON audio atlas object.
func (self *AudioSprite) SetConfigA(member interface{}) {
    self.Object.Set("config", member)
}

// If a sound is set to auto play, this holds the marker key of it.
func (self *AudioSprite) AutoplayKey() string{
    return self.Object.Get("autoplayKey").String()
}

// If a sound is set to auto play, this holds the marker key of it.
func (self *AudioSprite) SetAutoplayKeyA(member string) {
    self.Object.Set("autoplayKey", member)
}

// Is a sound set to autoplay or not?
func (self *AudioSprite) Autoplay() bool{
    return self.Object.Get("autoplay").Bool()
}

// Is a sound set to autoplay or not?
func (self *AudioSprite) SetAutoplayA(member bool) {
    self.Object.Set("autoplay", member)
}

// An object containing the Phaser.Sound objects for the Audio Sprite.
func (self *AudioSprite) Sounds() interface{}{
    return self.Object.Get("sounds")
}

// An object containing the Phaser.Sound objects for the Audio Sprite.
func (self *AudioSprite) SetSoundsA(member interface{}) {
    self.Object.Set("sounds", member)
}



// Play a sound with the given name.
func (self *AudioSprite) Play() *Sound{
    return &Sound{self.Object.Call("play")}
}

// Play a sound with the given name.
func (self *AudioSprite) Play1O(marker string) *Sound{
    return &Sound{self.Object.Call("play", marker)}
}

// Play a sound with the given name.
func (self *AudioSprite) Play2O(marker string, volume int) *Sound{
    return &Sound{self.Object.Call("play", marker, volume)}
}

// Play a sound with the given name.
func (self *AudioSprite) PlayI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("play", args)}
}

// Stop a sound with the given name.
func (self *AudioSprite) Stop() {
    self.Object.Call("stop")
}

// Stop a sound with the given name.
func (self *AudioSprite) Stop1O(marker string) {
    self.Object.Call("stop", marker)
}

// Stop a sound with the given name.
func (self *AudioSprite) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Get a sound with the given name.
func (self *AudioSprite) Get(marker string) *Sound{
    return &Sound{self.Object.Call("get", marker)}
}

// Get a sound with the given name.
func (self *AudioSprite) GetI(args ...interface{}) *Sound{
    return &Sound{self.Object.Call("get", args)}
}
