// Automatic generation for Phaser.LoaderParser
// generated file LoaderParser.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Phaser.LoaderParser parses data objects from Phaser.Loader that need more preparation before they can be inserted into the Cache.
type LoaderParser struct {
    *js.Object
}




// Alias for xmlBitmapFont, for backwards compatibility.
func (self *LoaderParser) BitmapFontI(args ...interface{}) interface{}{
    return self.Call("bitmapFont", args)
}

// Parse a Bitmap Font from an XML file.
func (self *LoaderParser) XmlBitmapFontI(args ...interface{}) interface{}{
    return self.Call("xmlBitmapFont", args)
}

// Parse a Bitmap Font from a JSON file.
func (self *LoaderParser) JsonBitmapFontI(args ...interface{}) interface{}{
    return self.Call("jsonBitmapFont", args)
}

// Finalize Bitmap Font parsing.
func (self *LoaderParser) FinalizeBitmapFontI(args ...interface{}) interface{}{
    return self.Call("finalizeBitmapFont", args)
}
