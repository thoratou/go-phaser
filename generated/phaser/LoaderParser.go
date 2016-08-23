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


// Phaser.LoaderParser parses data objects from Phaser.Loader that need more preparation before they can be inserted into the Cache.
func NewLoaderParser() *LoaderParser {
    return &LoaderParser{js.Global.Get("Phaser").Get("LoaderParser").New()}
}

// Phaser.LoaderParser parses data objects from Phaser.Loader that need more preparation before they can be inserted into the Cache.
func NewLoaderParserI(args ...interface{}) *LoaderParser {
    return &LoaderParser{js.Global.Get("Phaser").Get("LoaderParser").New(args)}
}





// Alias for xmlBitmapFont, for backwards compatibility.
func (self *LoaderParser) BitmapFont(xml interface{}, baseTexture *BaseTexture) interface{}{
    return self.Object.Call("bitmapFont", xml, baseTexture)
}

// Alias for xmlBitmapFont, for backwards compatibility.
func (self *LoaderParser) BitmapFont1O(xml interface{}, baseTexture *BaseTexture, xSpacing int) interface{}{
    return self.Object.Call("bitmapFont", xml, baseTexture, xSpacing)
}

// Alias for xmlBitmapFont, for backwards compatibility.
func (self *LoaderParser) BitmapFont2O(xml interface{}, baseTexture *BaseTexture, xSpacing int, ySpacing int) interface{}{
    return self.Object.Call("bitmapFont", xml, baseTexture, xSpacing, ySpacing)
}

// Alias for xmlBitmapFont, for backwards compatibility.
func (self *LoaderParser) BitmapFontI(args ...interface{}) interface{}{
    return self.Object.Call("bitmapFont", args)
}

// Parse a Bitmap Font from an XML file.
func (self *LoaderParser) XmlBitmapFont(xml interface{}, baseTexture *BaseTexture) interface{}{
    return self.Object.Call("xmlBitmapFont", xml, baseTexture)
}

// Parse a Bitmap Font from an XML file.
func (self *LoaderParser) XmlBitmapFont1O(xml interface{}, baseTexture *BaseTexture, xSpacing int) interface{}{
    return self.Object.Call("xmlBitmapFont", xml, baseTexture, xSpacing)
}

// Parse a Bitmap Font from an XML file.
func (self *LoaderParser) XmlBitmapFont2O(xml interface{}, baseTexture *BaseTexture, xSpacing int, ySpacing int) interface{}{
    return self.Object.Call("xmlBitmapFont", xml, baseTexture, xSpacing, ySpacing)
}

// Parse a Bitmap Font from an XML file.
func (self *LoaderParser) XmlBitmapFontI(args ...interface{}) interface{}{
    return self.Object.Call("xmlBitmapFont", args)
}

// Parse a Bitmap Font from a JSON file.
func (self *LoaderParser) JsonBitmapFont(json interface{}, baseTexture *BaseTexture) interface{}{
    return self.Object.Call("jsonBitmapFont", json, baseTexture)
}

// Parse a Bitmap Font from a JSON file.
func (self *LoaderParser) JsonBitmapFont1O(json interface{}, baseTexture *BaseTexture, xSpacing int) interface{}{
    return self.Object.Call("jsonBitmapFont", json, baseTexture, xSpacing)
}

// Parse a Bitmap Font from a JSON file.
func (self *LoaderParser) JsonBitmapFont2O(json interface{}, baseTexture *BaseTexture, xSpacing int, ySpacing int) interface{}{
    return self.Object.Call("jsonBitmapFont", json, baseTexture, xSpacing, ySpacing)
}

// Parse a Bitmap Font from a JSON file.
func (self *LoaderParser) JsonBitmapFontI(args ...interface{}) interface{}{
    return self.Object.Call("jsonBitmapFont", args)
}

// Finalize Bitmap Font parsing.
func (self *LoaderParser) FinalizeBitmapFont(baseTexture *BaseTexture, bitmapFontData interface{}) interface{}{
    return self.Object.Call("finalizeBitmapFont", baseTexture, bitmapFontData)
}

// Finalize Bitmap Font parsing.
func (self *LoaderParser) FinalizeBitmapFontI(args ...interface{}) interface{}{
    return self.Object.Call("finalizeBitmapFont", args)
}
