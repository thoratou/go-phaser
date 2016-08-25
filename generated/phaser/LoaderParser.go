// Package phaser Automatic generation for Phaser.LoaderParser
// generated file LoaderParser.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// LoaderParser Phaser.LoaderParser parses data objects from Phaser.Loader that need more preparation before they can be inserted into the Cache.
type LoaderParser struct {
    *js.Object
}

// NewLoaderParser Phaser.LoaderParser parses data objects from Phaser.Loader that need more preparation before they can be inserted into the Cache.
func NewLoaderParser() *LoaderParser {
    return &LoaderParser{js.Global.Get("Phaser").Get("LoaderParser").New()}
}
// NewLoaderParserI Phaser.LoaderParser parses data objects from Phaser.Loader that need more preparation before they can be inserted into the Cache.
func NewLoaderParserI(args ...interface{}) *LoaderParser {
    return &LoaderParser{js.Global.Get("Phaser").Get("LoaderParser").New(args)}
}



// LoaderParser Binding conversion method to LoaderParser point 
func ToLoaderParser(jsStruct interface{}) *LoaderParser {
    if object, ok := jsStruct.(*js.Object); ok {
		return &LoaderParser{Object: object}
	}
	return nil
}




// BitmapFont Alias for xmlBitmapFont, for backwards compatibility.
func (self *LoaderParser) BitmapFont(xml interface{}, baseTexture *BaseTexture) interface{}{
    return self.Object.Call("bitmapFont", xml, baseTexture)
}

// BitmapFont1O Alias for xmlBitmapFont, for backwards compatibility.
func (self *LoaderParser) BitmapFont1O(xml interface{}, baseTexture *BaseTexture, xSpacing int) interface{}{
    return self.Object.Call("bitmapFont", xml, baseTexture, xSpacing)
}

// BitmapFont2O Alias for xmlBitmapFont, for backwards compatibility.
func (self *LoaderParser) BitmapFont2O(xml interface{}, baseTexture *BaseTexture, xSpacing int, ySpacing int) interface{}{
    return self.Object.Call("bitmapFont", xml, baseTexture, xSpacing, ySpacing)
}

// BitmapFontI Alias for xmlBitmapFont, for backwards compatibility.
func (self *LoaderParser) BitmapFontI(args ...interface{}) interface{}{
    return self.Object.Call("bitmapFont", args)
}

// XmlBitmapFont Parse a Bitmap Font from an XML file.
func (self *LoaderParser) XmlBitmapFont(xml interface{}, baseTexture *BaseTexture) interface{}{
    return self.Object.Call("xmlBitmapFont", xml, baseTexture)
}

// XmlBitmapFont1O Parse a Bitmap Font from an XML file.
func (self *LoaderParser) XmlBitmapFont1O(xml interface{}, baseTexture *BaseTexture, xSpacing int) interface{}{
    return self.Object.Call("xmlBitmapFont", xml, baseTexture, xSpacing)
}

// XmlBitmapFont2O Parse a Bitmap Font from an XML file.
func (self *LoaderParser) XmlBitmapFont2O(xml interface{}, baseTexture *BaseTexture, xSpacing int, ySpacing int) interface{}{
    return self.Object.Call("xmlBitmapFont", xml, baseTexture, xSpacing, ySpacing)
}

// XmlBitmapFontI Parse a Bitmap Font from an XML file.
func (self *LoaderParser) XmlBitmapFontI(args ...interface{}) interface{}{
    return self.Object.Call("xmlBitmapFont", args)
}

// JsonBitmapFont Parse a Bitmap Font from a JSON file.
func (self *LoaderParser) JsonBitmapFont(json interface{}, baseTexture *BaseTexture) interface{}{
    return self.Object.Call("jsonBitmapFont", json, baseTexture)
}

// JsonBitmapFont1O Parse a Bitmap Font from a JSON file.
func (self *LoaderParser) JsonBitmapFont1O(json interface{}, baseTexture *BaseTexture, xSpacing int) interface{}{
    return self.Object.Call("jsonBitmapFont", json, baseTexture, xSpacing)
}

// JsonBitmapFont2O Parse a Bitmap Font from a JSON file.
func (self *LoaderParser) JsonBitmapFont2O(json interface{}, baseTexture *BaseTexture, xSpacing int, ySpacing int) interface{}{
    return self.Object.Call("jsonBitmapFont", json, baseTexture, xSpacing, ySpacing)
}

// JsonBitmapFontI Parse a Bitmap Font from a JSON file.
func (self *LoaderParser) JsonBitmapFontI(args ...interface{}) interface{}{
    return self.Object.Call("jsonBitmapFont", args)
}

// FinalizeBitmapFont Finalize Bitmap Font parsing.
func (self *LoaderParser) FinalizeBitmapFont(baseTexture *BaseTexture, bitmapFontData interface{}) interface{}{
    return self.Object.Call("finalizeBitmapFont", baseTexture, bitmapFontData)
}

// FinalizeBitmapFontI Finalize Bitmap Font parsing.
func (self *LoaderParser) FinalizeBitmapFontI(args ...interface{}) interface{}{
    return self.Object.Call("finalizeBitmapFont", args)
}

