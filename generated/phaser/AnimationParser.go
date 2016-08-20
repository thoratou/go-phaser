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
func (self *AnimationParser) SpriteSheetI(args ...interface{}) FrameData{
    return FrameData{self.Call("spriteSheet", args)}
}

// Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataI(args ...interface{}) FrameData{
    return FrameData{self.Call("JSONData", args)}
}

// Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataPyxelI(args ...interface{}) FrameData{
    return FrameData{self.Call("JSONDataPyxel", args)}
}

// Parse the JSON data and extract the animation frame data from it.
func (self *AnimationParser) JSONDataHashI(args ...interface{}) FrameData{
    return FrameData{self.Call("JSONDataHash", args)}
}

// Parse the XML data and extract the animation frame data from it.
func (self *AnimationParser) XMLDataI(args ...interface{}) FrameData{
    return FrameData{self.Call("XMLData", args)}
}
