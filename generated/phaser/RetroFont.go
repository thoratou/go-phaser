// Package phaser Automatic generation for Phaser.RetroFont
// generated file RetroFont.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// RetroFont A Retro Font is similar to a BitmapFont, in that it uses a texture to render the text. However unlike a BitmapFont every character in a RetroFont
// is the same size. This makes it similar to a sprite sheet. You typically find font sheets like this from old 8/16-bit games and demos.
type RetroFont struct {
    *js.Object
}

// NewRetroFont A Retro Font is similar to a BitmapFont, in that it uses a texture to render the text. However unlike a BitmapFont every character in a RetroFont
// is the same size. This makes it similar to a sprite sheet. You typically find font sheets like this from old 8/16-bit games and demos.
func NewRetroFont(game *Game, key string, characterWidth int, characterHeight int, chars string) *RetroFont {
    return &RetroFont{js.Global.Get("Phaser").Get("RetroFont").New(game, key, characterWidth, characterHeight, chars)}
}
// NewRetroFont1O A Retro Font is similar to a BitmapFont, in that it uses a texture to render the text. However unlike a BitmapFont every character in a RetroFont
// is the same size. This makes it similar to a sprite sheet. You typically find font sheets like this from old 8/16-bit games and demos.
func NewRetroFont1O(game *Game, key string, characterWidth int, characterHeight int, chars string, charsPerRow int) *RetroFont {
    return &RetroFont{js.Global.Get("Phaser").Get("RetroFont").New(game, key, characterWidth, characterHeight, chars, charsPerRow)}
}
// NewRetroFont2O A Retro Font is similar to a BitmapFont, in that it uses a texture to render the text. However unlike a BitmapFont every character in a RetroFont
// is the same size. This makes it similar to a sprite sheet. You typically find font sheets like this from old 8/16-bit games and demos.
func NewRetroFont2O(game *Game, key string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int) *RetroFont {
    return &RetroFont{js.Global.Get("Phaser").Get("RetroFont").New(game, key, characterWidth, characterHeight, chars, charsPerRow, xSpacing)}
}
// NewRetroFont3O A Retro Font is similar to a BitmapFont, in that it uses a texture to render the text. However unlike a BitmapFont every character in a RetroFont
// is the same size. This makes it similar to a sprite sheet. You typically find font sheets like this from old 8/16-bit games and demos.
func NewRetroFont3O(game *Game, key string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int, ySpacing int) *RetroFont {
    return &RetroFont{js.Global.Get("Phaser").Get("RetroFont").New(game, key, characterWidth, characterHeight, chars, charsPerRow, xSpacing, ySpacing)}
}
// NewRetroFont4O A Retro Font is similar to a BitmapFont, in that it uses a texture to render the text. However unlike a BitmapFont every character in a RetroFont
// is the same size. This makes it similar to a sprite sheet. You typically find font sheets like this from old 8/16-bit games and demos.
func NewRetroFont4O(game *Game, key string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int, ySpacing int, xOffset int) *RetroFont {
    return &RetroFont{js.Global.Get("Phaser").Get("RetroFont").New(game, key, characterWidth, characterHeight, chars, charsPerRow, xSpacing, ySpacing, xOffset)}
}
// NewRetroFont5O A Retro Font is similar to a BitmapFont, in that it uses a texture to render the text. However unlike a BitmapFont every character in a RetroFont
// is the same size. This makes it similar to a sprite sheet. You typically find font sheets like this from old 8/16-bit games and demos.
func NewRetroFont5O(game *Game, key string, characterWidth int, characterHeight int, chars string, charsPerRow int, xSpacing int, ySpacing int, xOffset int, yOffset int) *RetroFont {
    return &RetroFont{js.Global.Get("Phaser").Get("RetroFont").New(game, key, characterWidth, characterHeight, chars, charsPerRow, xSpacing, ySpacing, xOffset, yOffset)}
}
// NewRetroFontI A Retro Font is similar to a BitmapFont, in that it uses a texture to render the text. However unlike a BitmapFont every character in a RetroFont
// is the same size. This makes it similar to a sprite sheet. You typically find font sheets like this from old 8/16-bit games and demos.
func NewRetroFontI(args ...interface{}) *RetroFont {
    return &RetroFont{js.Global.Get("Phaser").Get("RetroFont").New(args)}
}



// RetroFont Binding conversion method to RetroFont point 
func ToRetroFont(jsStruct interface{}) *RetroFont {
    if object, ok := jsStruct.(*js.Object); ok {
		return &RetroFont{Object: object}
	}
	return nil
}



// CharacterWidth The width of each character in the font set.
func (self *RetroFont) CharacterWidth() int{
    return self.Object.Get("characterWidth").Int()
}

// SetCharacterWidthA The width of each character in the font set.
func (self *RetroFont) SetCharacterWidthA(member int) {
    self.Object.Set("characterWidth", member)
}

// CharacterHeight The height of each character in the font set.
func (self *RetroFont) CharacterHeight() int{
    return self.Object.Get("characterHeight").Int()
}

// SetCharacterHeightA The height of each character in the font set.
func (self *RetroFont) SetCharacterHeightA(member int) {
    self.Object.Set("characterHeight", member)
}

// CharacterSpacingX If the characters in the font set have horizontal spacing between them set the required amount here.
func (self *RetroFont) CharacterSpacingX() int{
    return self.Object.Get("characterSpacingX").Int()
}

// SetCharacterSpacingXA If the characters in the font set have horizontal spacing between them set the required amount here.
func (self *RetroFont) SetCharacterSpacingXA(member int) {
    self.Object.Set("characterSpacingX", member)
}

// CharacterSpacingY If the characters in the font set have vertical spacing between them set the required amount here.
func (self *RetroFont) CharacterSpacingY() int{
    return self.Object.Get("characterSpacingY").Int()
}

// SetCharacterSpacingYA If the characters in the font set have vertical spacing between them set the required amount here.
func (self *RetroFont) SetCharacterSpacingYA(member int) {
    self.Object.Set("characterSpacingY", member)
}

// CharacterPerRow The number of characters per row in the font set.
func (self *RetroFont) CharacterPerRow() int{
    return self.Object.Get("characterPerRow").Int()
}

// SetCharacterPerRowA The number of characters per row in the font set.
func (self *RetroFont) SetCharacterPerRowA(member int) {
    self.Object.Set("characterPerRow", member)
}

// OffsetX If the font set doesn't start at the top left of the given image, specify the X coordinate offset here.
func (self *RetroFont) OffsetX() int{
    return self.Object.Get("offsetX").Int()
}

// SetOffsetXA If the font set doesn't start at the top left of the given image, specify the X coordinate offset here.
func (self *RetroFont) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// OffsetY If the font set doesn't start at the top left of the given image, specify the Y coordinate offset here.
func (self *RetroFont) OffsetY() int{
    return self.Object.Get("offsetY").Int()
}

// SetOffsetYA If the font set doesn't start at the top left of the given image, specify the Y coordinate offset here.
func (self *RetroFont) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}

// Align Alignment of the text when multiLine = true or a fixedWidth is set. Set to RetroFont.ALIGN_LEFT (default), RetroFont.ALIGN_RIGHT or RetroFont.ALIGN_CENTER.
func (self *RetroFont) Align() string{
    return self.Object.Get("align").String()
}

// SetAlignA Alignment of the text when multiLine = true or a fixedWidth is set. Set to RetroFont.ALIGN_LEFT (default), RetroFont.ALIGN_RIGHT or RetroFont.ALIGN_CENTER.
func (self *RetroFont) SetAlignA(member string) {
    self.Object.Set("align", member)
}

// MultiLine If set to true all carriage-returns in text will form new lines (see align). If false the font will only contain one single line of text (the default)
func (self *RetroFont) MultiLine() bool{
    return self.Object.Get("multiLine").Bool()
}

// SetMultiLineA If set to true all carriage-returns in text will form new lines (see align). If false the font will only contain one single line of text (the default)
func (self *RetroFont) SetMultiLineA(member bool) {
    self.Object.Set("multiLine", member)
}

// AutoUpperCase Automatically convert any text to upper case. Lots of old bitmap fonts only contain upper-case characters, so the default is true.
func (self *RetroFont) AutoUpperCase() bool{
    return self.Object.Get("autoUpperCase").Bool()
}

// SetAutoUpperCaseA Automatically convert any text to upper case. Lots of old bitmap fonts only contain upper-case characters, so the default is true.
func (self *RetroFont) SetAutoUpperCaseA(member bool) {
    self.Object.Set("autoUpperCase", member)
}

// CustomSpacingX Adds horizontal spacing between each character of the font, in pixels.
func (self *RetroFont) CustomSpacingX() int{
    return self.Object.Get("customSpacingX").Int()
}

// SetCustomSpacingXA Adds horizontal spacing between each character of the font, in pixels.
func (self *RetroFont) SetCustomSpacingXA(member int) {
    self.Object.Set("customSpacingX", member)
}

// CustomSpacingY Adds vertical spacing between each line of multi-line text, set in pixels.
func (self *RetroFont) CustomSpacingY() int{
    return self.Object.Get("customSpacingY").Int()
}

// SetCustomSpacingYA Adds vertical spacing between each line of multi-line text, set in pixels.
func (self *RetroFont) SetCustomSpacingYA(member int) {
    self.Object.Set("customSpacingY", member)
}

// FixedWidth If you need this RetroFont image to have a fixed width you can set the width in this value.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) FixedWidth() int{
    return self.Object.Get("fixedWidth").Int()
}

// SetFixedWidthA If you need this RetroFont image to have a fixed width you can set the width in this value.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) SetFixedWidthA(member int) {
    self.Object.Set("fixedWidth", member)
}

// FontSet A reference to the image stored in the Game.Cache that contains the font.
func (self *RetroFont) FontSet() *Image{
    return &Image{self.Object.Get("fontSet")}
}

// SetFontSetA A reference to the image stored in the Game.Cache that contains the font.
func (self *RetroFont) SetFontSetA(member *Image) {
    self.Object.Set("fontSet", member)
}

// FrameData The FrameData representing this Retro Font.
func (self *RetroFont) FrameData() *FrameData{
    return &FrameData{self.Object.Get("frameData")}
}

// SetFrameDataA The FrameData representing this Retro Font.
func (self *RetroFont) SetFrameDataA(member *FrameData) {
    self.Object.Set("frameData", member)
}

// Stamp The image that is stamped to the RenderTexture for each character in the font.
func (self *RetroFont) Stamp() *Image{
    return &Image{self.Object.Get("stamp")}
}

// SetStampA The image that is stamped to the RenderTexture for each character in the font.
func (self *RetroFont) SetStampA(member *Image) {
    self.Object.Set("stamp", member)
}

// Type Base Phaser object type.
func (self *RetroFont) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA Base Phaser object type.
func (self *RetroFont) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// ALIGN_LEFT Align each line of multi-line text to the left.
func (self *RetroFont) ALIGN_LEFT() string{
    return self.Object.Get("ALIGN_LEFT").String()
}

// SetALIGN_LEFTA Align each line of multi-line text to the left.
func (self *RetroFont) SetALIGN_LEFTA(member string) {
    self.Object.Set("ALIGN_LEFT", member)
}

// ALIGN_RIGHT Align each line of multi-line text to the right.
func (self *RetroFont) ALIGN_RIGHT() string{
    return self.Object.Get("ALIGN_RIGHT").String()
}

// SetALIGN_RIGHTA Align each line of multi-line text to the right.
func (self *RetroFont) SetALIGN_RIGHTA(member string) {
    self.Object.Set("ALIGN_RIGHT", member)
}

// ALIGN_CENTER Align each line of multi-line text in the center.
func (self *RetroFont) ALIGN_CENTER() string{
    return self.Object.Get("ALIGN_CENTER").String()
}

// SetALIGN_CENTERA Align each line of multi-line text in the center.
func (self *RetroFont) SetALIGN_CENTERA(member string) {
    self.Object.Set("ALIGN_CENTER", member)
}

// TEXT_SET1 Text Set 1 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
func (self *RetroFont) TEXT_SET1() string{
    return self.Object.Get("TEXT_SET1").String()
}

// SetTEXT_SET1A Text Set 1 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
func (self *RetroFont) SetTEXT_SET1A(member string) {
    self.Object.Set("TEXT_SET1", member)
}

// TEXT_SET2 Text Set 2 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) TEXT_SET2() string{
    return self.Object.Get("TEXT_SET2").String()
}

// SetTEXT_SET2A Text Set 2 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) SetTEXT_SET2A(member string) {
    self.Object.Set("TEXT_SET2", member)
}

// TEXT_SET3 Text Set 3 = ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
func (self *RetroFont) TEXT_SET3() string{
    return self.Object.Get("TEXT_SET3").String()
}

// SetTEXT_SET3A Text Set 3 = ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
func (self *RetroFont) SetTEXT_SET3A(member string) {
    self.Object.Set("TEXT_SET3", member)
}

// TEXT_SET4 Text Set 4 = ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789
func (self *RetroFont) TEXT_SET4() string{
    return self.Object.Get("TEXT_SET4").String()
}

// SetTEXT_SET4A Text Set 4 = ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789
func (self *RetroFont) SetTEXT_SET4A(member string) {
    self.Object.Set("TEXT_SET4", member)
}

// TEXT_SET5 Text Set 5 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,/() '!?-*:0123456789
func (self *RetroFont) TEXT_SET5() string{
    return self.Object.Get("TEXT_SET5").String()
}

// SetTEXT_SET5A Text Set 5 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,/() '!?-*:0123456789
func (self *RetroFont) SetTEXT_SET5A(member string) {
    self.Object.Set("TEXT_SET5", member)
}

// TEXT_SET6 Text Set 6 = ABCDEFGHIJKLMNOPQRSTUVWXYZ!?:;0123456789"(),-.'
func (self *RetroFont) TEXT_SET6() string{
    return self.Object.Get("TEXT_SET6").String()
}

// SetTEXT_SET6A Text Set 6 = ABCDEFGHIJKLMNOPQRSTUVWXYZ!?:;0123456789"(),-.'
func (self *RetroFont) SetTEXT_SET6A(member string) {
    self.Object.Set("TEXT_SET6", member)
}

// TEXT_SET7 Text Set 7 = AGMSY+:4BHNTZ!;5CIOU.?06DJPV,(17EKQW")28FLRX-'39
func (self *RetroFont) TEXT_SET7() string{
    return self.Object.Get("TEXT_SET7").String()
}

// SetTEXT_SET7A Text Set 7 = AGMSY+:4BHNTZ!;5CIOU.?06DJPV,(17EKQW")28FLRX-'39
func (self *RetroFont) SetTEXT_SET7A(member string) {
    self.Object.Set("TEXT_SET7", member)
}

// TEXT_SET8 Text Set 8 = 0123456789 .ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) TEXT_SET8() string{
    return self.Object.Get("TEXT_SET8").String()
}

// SetTEXT_SET8A Text Set 8 = 0123456789 .ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) SetTEXT_SET8A(member string) {
    self.Object.Set("TEXT_SET8", member)
}

// TEXT_SET9 Text Set 9 = ABCDEFGHIJKLMNOPQRSTUVWXYZ()-0123456789.:,'"?!
func (self *RetroFont) TEXT_SET9() string{
    return self.Object.Get("TEXT_SET9").String()
}

// SetTEXT_SET9A Text Set 9 = ABCDEFGHIJKLMNOPQRSTUVWXYZ()-0123456789.:,'"?!
func (self *RetroFont) SetTEXT_SET9A(member string) {
    self.Object.Set("TEXT_SET9", member)
}

// TEXT_SET10 Text Set 10 = ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) TEXT_SET10() string{
    return self.Object.Get("TEXT_SET10").String()
}

// SetTEXT_SET10A Text Set 10 = ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) SetTEXT_SET10A(member string) {
    self.Object.Set("TEXT_SET10", member)
}

// TEXT_SET11 Text Set 11 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,"-+!?()':;0123456789
func (self *RetroFont) TEXT_SET11() string{
    return self.Object.Get("TEXT_SET11").String()
}

// SetTEXT_SET11A Text Set 11 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,"-+!?()':;0123456789
func (self *RetroFont) SetTEXT_SET11A(member string) {
    self.Object.Set("TEXT_SET11", member)
}

// Text Set this value to update the text in this sprite. Carriage returns are automatically stripped out if multiLine is false. Text is converted to upper case if autoUpperCase is true.
func (self *RetroFont) Text() string{
    return self.Object.Get("text").String()
}

// SetTextA Set this value to update the text in this sprite. Carriage returns are automatically stripped out if multiLine is false. Text is converted to upper case if autoUpperCase is true.
func (self *RetroFont) SetTextA(member string) {
    self.Object.Set("text", member)
}

// Smoothed Sets if the stamp is smoothed or not.
func (self *RetroFont) Smoothed() bool{
    return self.Object.Get("smoothed").Bool()
}

// SetSmoothedA Sets if the stamp is smoothed or not.
func (self *RetroFont) SetSmoothedA(member bool) {
    self.Object.Set("smoothed", member)
}

// Game A reference to the currently running game.
func (self *RetroFont) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *RetroFont) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Key The key of the RenderTexture in the Cache, if stored there.
func (self *RetroFont) Key() string{
    return self.Object.Get("key").String()
}

// SetKeyA The key of the RenderTexture in the Cache, if stored there.
func (self *RetroFont) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// Width The with of the render texture
func (self *RetroFont) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The with of the render texture
func (self *RetroFont) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the render texture
func (self *RetroFont) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the render texture
func (self *RetroFont) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Resolution The Resolution of the texture.
func (self *RetroFont) Resolution() int{
    return self.Object.Get("resolution").Int()
}

// SetResolutionA The Resolution of the texture.
func (self *RetroFont) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// Frame The framing rectangle of the render texture
func (self *RetroFont) Frame() *Rectangle{
    return &Rectangle{self.Object.Get("frame")}
}

// SetFrameA The framing rectangle of the render texture
func (self *RetroFont) SetFrameA(member *Rectangle) {
    self.Object.Set("frame", member)
}

// Crop This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RetroFont) Crop() *Rectangle{
    return &Rectangle{self.Object.Get("crop")}
}

// SetCropA This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RetroFont) SetCropA(member *Rectangle) {
    self.Object.Set("crop", member)
}

// BaseTexture The base texture object that this texture uses
func (self *RetroFont) BaseTexture() *BaseTexture{
    return &BaseTexture{self.Object.Get("baseTexture")}
}

// SetBaseTextureA The base texture object that this texture uses
func (self *RetroFont) SetBaseTextureA(member *BaseTexture) {
    self.Object.Set("baseTexture", member)
}

// Renderer The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RetroFont) Renderer() interface{}{
    return self.Object.Get("renderer")
}

// SetRendererA The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RetroFont) SetRendererA(member interface{}) {
    self.Object.Set("renderer", member)
}

// Valid empty description
func (self *RetroFont) Valid() bool{
    return self.Object.Get("valid").Bool()
}

// SetValidA empty description
func (self *RetroFont) SetValidA(member bool) {
    self.Object.Set("valid", member)
}

// NoFrame Does this Texture have any frame data assigned to it?
func (self *RetroFont) NoFrame() bool{
    return self.Object.Get("noFrame").Bool()
}

// SetNoFrameA Does this Texture have any frame data assigned to it?
func (self *RetroFont) SetNoFrameA(member bool) {
    self.Object.Set("noFrame", member)
}

// Trim The texture trim data.
func (self *RetroFont) Trim() *Rectangle{
    return &Rectangle{self.Object.Get("trim")}
}

// SetTrimA The texture trim data.
func (self *RetroFont) SetTrimA(member *Rectangle) {
    self.Object.Set("trim", member)
}

// IsTiling Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RetroFont) IsTiling() bool{
    return self.Object.Get("isTiling").Bool()
}

// SetIsTilingA Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RetroFont) SetIsTilingA(member bool) {
    self.Object.Set("isTiling", member)
}

// RequiresUpdate This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RetroFont) RequiresUpdate() bool{
    return self.Object.Get("requiresUpdate").Bool()
}

// SetRequiresUpdateA This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RetroFont) SetRequiresUpdateA(member bool) {
    self.Object.Set("requiresUpdate", member)
}

// RequiresReTint This will let a renderer know that a tinted parent has updated its texture.
func (self *RetroFont) RequiresReTint() bool{
    return self.Object.Get("requiresReTint").Bool()
}

// SetRequiresReTintA This will let a renderer know that a tinted parent has updated its texture.
func (self *RetroFont) SetRequiresReTintA(member bool) {
    self.Object.Set("requiresReTint", member)
}


// SetFixedWidth If you need this RetroFont to have a fixed width and custom alignment you can set the width here.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) SetFixedWidth(width int) {
    self.Object.Call("setFixedWidth", width)
}

// SetFixedWidth1O If you need this RetroFont to have a fixed width and custom alignment you can set the width here.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) SetFixedWidth1O(width int, lineAlignment string) {
    self.Object.Call("setFixedWidth", width, lineAlignment)
}

// SetFixedWidthI If you need this RetroFont to have a fixed width and custom alignment you can set the width here.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) SetFixedWidthI(args ...interface{}) {
    self.Object.Call("setFixedWidth", args)
}

// SetText A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText(content string) {
    self.Object.Call("setText", content)
}

// SetText1O A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText1O(content string, multiLine bool) {
    self.Object.Call("setText", content, multiLine)
}

// SetText2O A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText2O(content string, multiLine bool, characterSpacing int) {
    self.Object.Call("setText", content, multiLine, characterSpacing)
}

// SetText3O A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText3O(content string, multiLine bool, characterSpacing int, lineSpacing int) {
    self.Object.Call("setText", content, multiLine, characterSpacing, lineSpacing)
}

// SetText4O A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText4O(content string, multiLine bool, characterSpacing int, lineSpacing int, lineAlignment string) {
    self.Object.Call("setText", content, multiLine, characterSpacing, lineSpacing, lineAlignment)
}

// SetText5O A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText5O(content string, multiLine bool, characterSpacing int, lineSpacing int, lineAlignment string, allowLowerCase bool) {
    self.Object.Call("setText", content, multiLine, characterSpacing, lineSpacing, lineAlignment, allowLowerCase)
}

// SetTextI A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetTextI(args ...interface{}) {
    self.Object.Call("setText", args)
}

// BuildRetroFontText Updates the texture with the new text.
func (self *RetroFont) BuildRetroFontText() {
    self.Object.Call("buildRetroFontText")
}

// BuildRetroFontTextI Updates the texture with the new text.
func (self *RetroFont) BuildRetroFontTextI(args ...interface{}) {
    self.Object.Call("buildRetroFontText", args)
}

// PasteLine Internal function that takes a single line of text (2nd parameter) and pastes it into the BitmapData at the given coordinates.
// Used by getLine and getMultiLine
func (self *RetroFont) PasteLine(line string, x int, y int, customSpacingX int) {
    self.Object.Call("pasteLine", line, x, y, customSpacingX)
}

// PasteLineI Internal function that takes a single line of text (2nd parameter) and pastes it into the BitmapData at the given coordinates.
// Used by getLine and getMultiLine
func (self *RetroFont) PasteLineI(args ...interface{}) {
    self.Object.Call("pasteLine", args)
}

// GetLongestLine Works out the longest line of text in _text and returns its length
func (self *RetroFont) GetLongestLine() int{
    return self.Object.Call("getLongestLine").Int()
}

// GetLongestLineI Works out the longest line of text in _text and returns its length
func (self *RetroFont) GetLongestLineI(args ...interface{}) int{
    return self.Object.Call("getLongestLine", args).Int()
}

// RemoveUnsupportedCharacters Internal helper function that removes all unsupported characters from the _text String, leaving only characters contained in the font set.
func (self *RetroFont) RemoveUnsupportedCharacters() string{
    return self.Object.Call("removeUnsupportedCharacters").String()
}

// RemoveUnsupportedCharacters1O Internal helper function that removes all unsupported characters from the _text String, leaving only characters contained in the font set.
func (self *RetroFont) RemoveUnsupportedCharacters1O(stripCR bool) string{
    return self.Object.Call("removeUnsupportedCharacters", stripCR).String()
}

// RemoveUnsupportedCharactersI Internal helper function that removes all unsupported characters from the _text String, leaving only characters contained in the font set.
func (self *RetroFont) RemoveUnsupportedCharactersI(args ...interface{}) string{
    return self.Object.Call("removeUnsupportedCharacters", args).String()
}

// UpdateOffset Updates the x and/or y offset that the font is rendered from. This updates all of the texture frames, so be careful how often it is called.
// Note that the values given for the x and y properties are either ADDED to or SUBTRACTED from (if negative) the existing offsetX/Y values of the characters.
// So if the current offsetY is 8 and you want it to start rendering from y16 you would call updateOffset(0, 8) to add 8 to the current y offset.
func (self *RetroFont) UpdateOffset() {
    self.Object.Call("updateOffset")
}

// UpdateOffset1O Updates the x and/or y offset that the font is rendered from. This updates all of the texture frames, so be careful how often it is called.
// Note that the values given for the x and y properties are either ADDED to or SUBTRACTED from (if negative) the existing offsetX/Y values of the characters.
// So if the current offsetY is 8 and you want it to start rendering from y16 you would call updateOffset(0, 8) to add 8 to the current y offset.
func (self *RetroFont) UpdateOffset1O(xOffset int) {
    self.Object.Call("updateOffset", xOffset)
}

// UpdateOffset2O Updates the x and/or y offset that the font is rendered from. This updates all of the texture frames, so be careful how often it is called.
// Note that the values given for the x and y properties are either ADDED to or SUBTRACTED from (if negative) the existing offsetX/Y values of the characters.
// So if the current offsetY is 8 and you want it to start rendering from y16 you would call updateOffset(0, 8) to add 8 to the current y offset.
func (self *RetroFont) UpdateOffset2O(xOffset int, yOffset int) {
    self.Object.Call("updateOffset", xOffset, yOffset)
}

// UpdateOffsetI Updates the x and/or y offset that the font is rendered from. This updates all of the texture frames, so be careful how often it is called.
// Note that the values given for the x and y properties are either ADDED to or SUBTRACTED from (if negative) the existing offsetX/Y values of the characters.
// So if the current offsetY is 8 and you want it to start rendering from y16 you would call updateOffset(0, 8) to add 8 to the current y offset.
func (self *RetroFont) UpdateOffsetI(args ...interface{}) {
    self.Object.Call("updateOffset", args)
}

// RenderXY This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RetroFont) RenderXY(displayObject interface{}, x int, y int) {
    self.Object.Call("renderXY", displayObject, x, y)
}

// RenderXY1O This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RetroFont) RenderXY1O(displayObject interface{}, x int, y int, clear bool) {
    self.Object.Call("renderXY", displayObject, x, y, clear)
}

// RenderXYI This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RetroFont) RenderXYI(args ...interface{}) {
    self.Object.Call("renderXY", args)
}

// RenderRawXY This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RetroFont) RenderRawXY(displayObject interface{}, x int, y int) {
    self.Object.Call("renderRawXY", displayObject, x, y)
}

// RenderRawXY1O This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RetroFont) RenderRawXY1O(displayObject interface{}, x int, y int, clear bool) {
    self.Object.Call("renderRawXY", displayObject, x, y, clear)
}

// RenderRawXYI This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RetroFont) RenderRawXYI(args ...interface{}) {
    self.Object.Call("renderRawXY", args)
}

// Render This function will draw the display object to the RenderTexture.
// 
// In versions of Phaser prior to 2.4.0 the second parameter was a Phaser.Point object. 
// This is now a Matrix allowing you much more control over how the Display Object is rendered.
// If you need to replicate the earlier behavior please use Phaser.RenderTexture.renderXY instead.
// 
// If you wish for the displayObject to be rendered taking its current scale, rotation and translation into account then either
// pass `null`, leave it undefined or pass `displayObject.worldTransform` as the matrix value.
func (self *RetroFont) Render(displayObject interface{}) {
    self.Object.Call("render", displayObject)
}

// Render1O This function will draw the display object to the RenderTexture.
// 
// In versions of Phaser prior to 2.4.0 the second parameter was a Phaser.Point object. 
// This is now a Matrix allowing you much more control over how the Display Object is rendered.
// If you need to replicate the earlier behavior please use Phaser.RenderTexture.renderXY instead.
// 
// If you wish for the displayObject to be rendered taking its current scale, rotation and translation into account then either
// pass `null`, leave it undefined or pass `displayObject.worldTransform` as the matrix value.
func (self *RetroFont) Render1O(displayObject interface{}, matrix *Matrix) {
    self.Object.Call("render", displayObject, matrix)
}

// Render2O This function will draw the display object to the RenderTexture.
// 
// In versions of Phaser prior to 2.4.0 the second parameter was a Phaser.Point object. 
// This is now a Matrix allowing you much more control over how the Display Object is rendered.
// If you need to replicate the earlier behavior please use Phaser.RenderTexture.renderXY instead.
// 
// If you wish for the displayObject to be rendered taking its current scale, rotation and translation into account then either
// pass `null`, leave it undefined or pass `displayObject.worldTransform` as the matrix value.
func (self *RetroFont) Render2O(displayObject interface{}, matrix *Matrix, clear bool) {
    self.Object.Call("render", displayObject, matrix, clear)
}

// RenderI This function will draw the display object to the RenderTexture.
// 
// In versions of Phaser prior to 2.4.0 the second parameter was a Phaser.Point object. 
// This is now a Matrix allowing you much more control over how the Display Object is rendered.
// If you need to replicate the earlier behavior please use Phaser.RenderTexture.renderXY instead.
// 
// If you wish for the displayObject to be rendered taking its current scale, rotation and translation into account then either
// pass `null`, leave it undefined or pass `displayObject.worldTransform` as the matrix value.
func (self *RetroFont) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// Resize Resizes the RenderTexture.
func (self *RetroFont) Resize(width int, height int, updateBase bool) {
    self.Object.Call("resize", width, height, updateBase)
}

// ResizeI Resizes the RenderTexture.
func (self *RetroFont) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Clear Clears the RenderTexture.
func (self *RetroFont) Clear() {
    self.Object.Call("clear")
}

// ClearI Clears the RenderTexture.
func (self *RetroFont) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// RenderWebGL This function will draw the display object to the texture.
func (self *RetroFont) RenderWebGL(displayObject *DisplayObject) {
    self.Object.Call("renderWebGL", displayObject)
}

// RenderWebGL1O This function will draw the display object to the texture.
func (self *RetroFont) RenderWebGL1O(displayObject *DisplayObject, matrix *Matrix) {
    self.Object.Call("renderWebGL", displayObject, matrix)
}

// RenderWebGL2O This function will draw the display object to the texture.
func (self *RetroFont) RenderWebGL2O(displayObject *DisplayObject, matrix *Matrix, clear bool) {
    self.Object.Call("renderWebGL", displayObject, matrix, clear)
}

// RenderWebGLI This function will draw the display object to the texture.
func (self *RetroFont) RenderWebGLI(args ...interface{}) {
    self.Object.Call("renderWebGL", args)
}

// RenderCanvas This function will draw the display object to the texture.
func (self *RetroFont) RenderCanvas(displayObject *DisplayObject) {
    self.Object.Call("renderCanvas", displayObject)
}

// RenderCanvas1O This function will draw the display object to the texture.
func (self *RetroFont) RenderCanvas1O(displayObject *DisplayObject, matrix *Matrix) {
    self.Object.Call("renderCanvas", displayObject, matrix)
}

// RenderCanvas2O This function will draw the display object to the texture.
func (self *RetroFont) RenderCanvas2O(displayObject *DisplayObject, matrix *Matrix, clear bool) {
    self.Object.Call("renderCanvas", displayObject, matrix, clear)
}

// RenderCanvasI This function will draw the display object to the texture.
func (self *RetroFont) RenderCanvasI(args ...interface{}) {
    self.Object.Call("renderCanvas", args)
}

// GetImage Will return a HTML Image of the texture
func (self *RetroFont) GetImage() *Image{
    return &Image{self.Object.Call("getImage")}
}

// GetImageI Will return a HTML Image of the texture
func (self *RetroFont) GetImageI(args ...interface{}) *Image{
    return &Image{self.Object.Call("getImage", args)}
}

// GetBase64 Will return a base64 encoded string of this texture. It works by calling RenderTexture.getCanvas and then running toDataURL on that.
func (self *RetroFont) GetBase64() string{
    return self.Object.Call("getBase64").String()
}

// GetBase64I Will return a base64 encoded string of this texture. It works by calling RenderTexture.getCanvas and then running toDataURL on that.
func (self *RetroFont) GetBase64I(args ...interface{}) string{
    return self.Object.Call("getBase64", args).String()
}

// GetCanvas Creates a Canvas element, renders this RenderTexture to it and then returns it.
func (self *RetroFont) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("getCanvas"))
}

// GetCanvasI Creates a Canvas element, renders this RenderTexture to it and then returns it.
func (self *RetroFont) GetCanvasI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("getCanvas", args))
}

// OnBaseTextureLoaded Called when the base texture is loaded
func (self *RetroFont) OnBaseTextureLoaded() {
    self.Object.Call("onBaseTextureLoaded")
}

// OnBaseTextureLoadedI Called when the base texture is loaded
func (self *RetroFont) OnBaseTextureLoadedI(args ...interface{}) {
    self.Object.Call("onBaseTextureLoaded", args)
}

// Destroy Destroys this texture
func (self *RetroFont) Destroy(destroyBase bool) {
    self.Object.Call("destroy", destroyBase)
}

// DestroyI Destroys this texture
func (self *RetroFont) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// SetFrame Specifies the region of the baseTexture that this texture will use.
func (self *RetroFont) SetFrame(frame *Rectangle) {
    self.Object.Call("setFrame", frame)
}

// SetFrameI Specifies the region of the baseTexture that this texture will use.
func (self *RetroFont) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// _updateUvs Updates the internal WebGL UV cache.
func (self *RetroFont) _updateUvs() {
    self.Object.Call("_updateUvs")
}

// _updateUvsI Updates the internal WebGL UV cache.
func (self *RetroFont) _updateUvsI(args ...interface{}) {
    self.Object.Call("_updateUvs", args)
}

