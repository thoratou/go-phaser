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


// A reference to the currently running Game.
func (self *AudioSprite) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *AudioSprite) SetGame(member Game) {
    self.Set("game", member)
}

// Asset key for the Audio Sprite.
func (self *AudioSprite) GetKey() string{
    return self.Get("key").String()
}

// Asset key for the Audio Sprite.
func (self *AudioSprite) SetKey(member string) {
    self.Set("key", member)
}

// JSON audio atlas object.
func (self *AudioSprite) GetConfig() interface{}{
    return self.Get("config")
}

// JSON audio atlas object.
func (self *AudioSprite) SetConfig(member interface{}) {
    self.Set("config", member)
}

// If a sound is set to auto play, this holds the marker key of it.
func (self *AudioSprite) GetAutoplayKey() string{
    return self.Get("autoplayKey").String()
}

// If a sound is set to auto play, this holds the marker key of it.
func (self *AudioSprite) SetAutoplayKey(member string) {
    self.Set("autoplayKey", member)
}

// Is a sound set to autoplay or not?
func (self *AudioSprite) GetAutoplay() bool{
    return self.Get("autoplay").Bool()
}

// Is a sound set to autoplay or not?
func (self *AudioSprite) SetAutoplay(member bool) {
    self.Set("autoplay", member)
}

// An object containing the Phaser.Sound objects for the Audio Sprite.
func (self *AudioSprite) GetSounds() interface{}{
    return self.Get("sounds")
}

// An object containing the Phaser.Sound objects for the Audio Sprite.
func (self *AudioSprite) SetSounds(member interface{}) {
    self.Set("sounds", member)
}



// Play a sound with the given name.
func (self *AudioSprite) PlayI(args ...interface{}) Sound{
    return Sound{self.Call("play", args)}
}

// Stop a sound with the given name.
func (self *AudioSprite) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// Get a sound with the given name.
func (self *AudioSprite) GetI(args ...interface{}) Sound{
    return Sound{self.Call("get", args)}
}
