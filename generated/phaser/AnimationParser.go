// Automatic generation for Phaser.AnimationParser
// generated file AnimationParser.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Responsible for parsing sprite sheet and JSON data into the internal FrameData format that Phaser uses for animations.
type AnimationParser struct {
    *js.Object
}




// Parse a Sprite Sheet and extract the animation frame data from it.
func (self *AnimationParser) SpriteSheet(game *Game, key interface{}, frameWidth int, frameHeight int, frameMax int, margin int, spacing int) *FrameData{
    return &FrameData{self.Object.Call("spriteSheet", game, key, frameWidth, frameHeight, frameMax, margin, spacing)}
}

// Parse a Sprite Sheet and extract the animation frame data from it.
func (self *AnimationParser) SpriteSheetI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("spriteSheet", args)}
}

// Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONData(game *Game, json interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONData", game, json)}
}

// Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONData", args)}
}

// Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataPyxel(game *Game, json interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONDataPyxel", game, json)}
}

// Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataPyxelI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONDataPyxel", args)}
}

// Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataHash(game *Game, json interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONDataHash", game, json)}
}

// Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataHashI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("JSONDataHash", args)}
}

// Parse the XML data and extract the animation frame data from it.
func (self *AnimationParser) XMLData(game *Game, xml interface{}) *FrameData{
    return &FrameData{self.Object.Call("XMLData", game, xml)}
}

// Parse the XML data and extract the animation frame data from it.
func (self *AnimationParser) XMLDataI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("XMLData", args)}
}
