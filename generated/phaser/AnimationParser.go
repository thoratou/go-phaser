// Package phaser Automatic generation for Phaser.AnimationParser
// generated file AnimationParser.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// AnimationParser Responsible for parsing sprite sheet and JSON data into the internal FrameData format that Phaser uses for animations.
type AnimationParser struct {
    *js.Object
}

// NewAnimationParser Responsible for parsing sprite sheet and JSON data into the internal FrameData format that Phaser uses for animations.
func NewAnimationParser() *AnimationParser {
    return &AnimationParser{js.Global.Get("Phaser").Get("AnimationParser").New()}
}
// NewAnimationParserI Responsible for parsing sprite sheet and JSON data into the internal FrameData format that Phaser uses for animations.
func NewAnimationParserI(args ...interface{}) *AnimationParser {
    return &AnimationParser{js.Global.Get("Phaser").Get("AnimationParser").New(args)}
}




// SpriteSheet Parse a Sprite Sheet and extract the animation frame data from it.
func (self *AnimationParser) SpriteSheet(game *Game, key interface{}, frameWidth int, frameHeight int) *FrameData{
    return &FrameData{self.Object.Call("spriteSheet", game, key, frameWidth, frameHeight)}
}

// SpriteSheet1O Parse a Sprite Sheet and extract the animation frame data from it.
func (self *AnimationParser) SpriteSheet1O(game *Game, key interface{}, frameWidth int, frameHeight int, frameMax int) *FrameData{
    return &FrameData{self.Object.Call("spriteSheet", game, key, frameWidth, frameHeight, frameMax)}
}

// SpriteSheet2O Parse a Sprite Sheet and extract the animation frame data from it.
func (self *AnimationParser) SpriteSheet2O(game *Game, key interface{}, frameWidth int, frameHeight int, frameMax int, margin int) *FrameData{
    return &FrameData{self.Object.Call("spriteSheet", game, key, frameWidth, frameHeight, frameMax, margin)}
}

// SpriteSheet3O Parse a Sprite Sheet and extract the animation frame data from it.
func (self *AnimationParser) SpriteSheet3O(game *Game, key interface{}, frameWidth int, frameHeight int, frameMax int, margin int, spacing int) *FrameData{
    return &FrameData{self.Object.Call("spriteSheet", game, key, frameWidth, frameHeight, frameMax, margin, spacing)}
}

// SpriteSheetI Parse a Sprite Sheet and extract the animation frame data from it.
func (self *AnimationParser) SpriteSheetI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("spriteSheet", args)}
}

// JSONData Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONData(game *Game, json interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONData", game, json)}
}

// JSONDataI Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONData", args)}
}

// JSONDataPyxel Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataPyxel(game *Game, json interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONDataPyxel", game, json)}
}

// JSONDataPyxelI Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataPyxelI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONDataPyxel", args)}
}

// JSONDataHash Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataHash(game *Game, json interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONDataHash", game, json)}
}

// JSONDataHashI Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataHashI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONDataHash", args)}
}

// XMLData Parse the XML data and extract the animation frame data from it.
func (self *AnimationParser) XMLData(game *Game, xml interface{}) *FrameData{
    return &FrameData{self.Object.Call("XMLData", game, xml)}
}

// XMLDataI Parse the XML data and extract the animation frame data from it.
func (self *AnimationParser) XMLDataI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("XMLData", args)}
}

