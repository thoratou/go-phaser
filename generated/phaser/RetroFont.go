// Automatic generation for Phaser.RetroFont
// generated file RetroFont.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// A Retro Font is similar to a BitmapFont, in that it uses a texture to render the text. However unlike a BitmapFont every character in a RetroFont
// is the same size. This makes it similar to a sprite sheet. You typically find font sheets like this from old 8/16-bit games and demos.
type RetroFont struct {
    *js.Object
}


// The width of each character in the font set.
func (self *RetroFont) GetCharacterWidthA() int{
    return self.Object.Get("characterWidth").Int()
}

// The width of each character in the font set.
func (self *RetroFont) SetCharacterWidthA(member int) {
    self.Object.Set("characterWidth", member)
}

// The height of each character in the font set.
func (self *RetroFont) GetCharacterHeightA() int{
    return self.Object.Get("characterHeight").Int()
}

// The height of each character in the font set.
func (self *RetroFont) SetCharacterHeightA(member int) {
    self.Object.Set("characterHeight", member)
}

// If the characters in the font set have horizontal spacing between them set the required amount here.
func (self *RetroFont) GetCharacterSpacingXA() int{
    return self.Object.Get("characterSpacingX").Int()
}

// If the characters in the font set have horizontal spacing between them set the required amount here.
func (self *RetroFont) SetCharacterSpacingXA(member int) {
    self.Object.Set("characterSpacingX", member)
}

// If the characters in the font set have vertical spacing between them set the required amount here.
func (self *RetroFont) GetCharacterSpacingYA() int{
    return self.Object.Get("characterSpacingY").Int()
}

// If the characters in the font set have vertical spacing between them set the required amount here.
func (self *RetroFont) SetCharacterSpacingYA(member int) {
    self.Object.Set("characterSpacingY", member)
}

// The number of characters per row in the font set.
func (self *RetroFont) GetCharacterPerRowA() int{
    return self.Object.Get("characterPerRow").Int()
}

// The number of characters per row in the font set.
func (self *RetroFont) SetCharacterPerRowA(member int) {
    self.Object.Set("characterPerRow", member)
}

// If the font set doesn't start at the top left of the given image, specify the X coordinate offset here.
func (self *RetroFont) GetOffsetXA() int{
    return self.Object.Get("offsetX").Int()
}

// If the font set doesn't start at the top left of the given image, specify the X coordinate offset here.
func (self *RetroFont) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// If the font set doesn't start at the top left of the given image, specify the Y coordinate offset here.
func (self *RetroFont) GetOffsetYA() int{
    return self.Object.Get("offsetY").Int()
}

// If the font set doesn't start at the top left of the given image, specify the Y coordinate offset here.
func (self *RetroFont) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}

// Alignment of the text when multiLine = true or a fixedWidth is set. Set to RetroFont.ALIGN_LEFT (default), RetroFont.ALIGN_RIGHT or RetroFont.ALIGN_CENTER.
func (self *RetroFont) GetAlignA() string{
    return self.Object.Get("align").String()
}

// Alignment of the text when multiLine = true or a fixedWidth is set. Set to RetroFont.ALIGN_LEFT (default), RetroFont.ALIGN_RIGHT or RetroFont.ALIGN_CENTER.
func (self *RetroFont) SetAlignA(member string) {
    self.Object.Set("align", member)
}

// If set to true all carriage-returns in text will form new lines (see align). If false the font will only contain one single line of text (the default)
func (self *RetroFont) GetMultiLineA() bool{
    return self.Object.Get("multiLine").Bool()
}

// If set to true all carriage-returns in text will form new lines (see align). If false the font will only contain one single line of text (the default)
func (self *RetroFont) SetMultiLineA(member bool) {
    self.Object.Set("multiLine", member)
}

// Automatically convert any text to upper case. Lots of old bitmap fonts only contain upper-case characters, so the default is true.
func (self *RetroFont) GetAutoUpperCaseA() bool{
    return self.Object.Get("autoUpperCase").Bool()
}

// Automatically convert any text to upper case. Lots of old bitmap fonts only contain upper-case characters, so the default is true.
func (self *RetroFont) SetAutoUpperCaseA(member bool) {
    self.Object.Set("autoUpperCase", member)
}

// Adds horizontal spacing between each character of the font, in pixels.
func (self *RetroFont) GetCustomSpacingXA() int{
    return self.Object.Get("customSpacingX").Int()
}

// Adds horizontal spacing between each character of the font, in pixels.
func (self *RetroFont) SetCustomSpacingXA(member int) {
    self.Object.Set("customSpacingX", member)
}

// Adds vertical spacing between each line of multi-line text, set in pixels.
func (self *RetroFont) GetCustomSpacingYA() int{
    return self.Object.Get("customSpacingY").Int()
}

// Adds vertical spacing between each line of multi-line text, set in pixels.
func (self *RetroFont) SetCustomSpacingYA(member int) {
    self.Object.Set("customSpacingY", member)
}

// If you need this RetroFont image to have a fixed width you can set the width in this value.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) GetFixedWidthA() int{
    return self.Object.Get("fixedWidth").Int()
}

// If you need this RetroFont image to have a fixed width you can set the width in this value.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) SetFixedWidthA(member int) {
    self.Object.Set("fixedWidth", member)
}

// A reference to the image stored in the Game.Cache that contains the font.
func (self *RetroFont) GetFontSetA() *Image{
    return &Image{self.Object.Get("fontSet")}
}

// A reference to the image stored in the Game.Cache that contains the font.
func (self *RetroFont) SetFontSetA(member *Image) {
    self.Object.Set("fontSet", member)
}

// The FrameData representing this Retro Font.
func (self *RetroFont) GetFrameDataA() *FrameData{
    return &FrameData{self.Object.Get("frameData")}
}

// The FrameData representing this Retro Font.
func (self *RetroFont) SetFrameDataA(member *FrameData) {
    self.Object.Set("frameData", member)
}

// The image that is stamped to the RenderTexture for each character in the font.
func (self *RetroFont) GetStampA() *Image{
    return &Image{self.Object.Get("stamp")}
}

// The image that is stamped to the RenderTexture for each character in the font.
func (self *RetroFont) SetStampA(member *Image) {
    self.Object.Set("stamp", member)
}

// Base Phaser object type.
func (self *RetroFont) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// Base Phaser object type.
func (self *RetroFont) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Align each line of multi-line text to the left.
func (self *RetroFont) GetALIGN_LEFTA() string{
    return self.Object.Get("ALIGN_LEFT").String()
}

// Align each line of multi-line text to the left.
func (self *RetroFont) SetALIGN_LEFTA(member string) {
    self.Object.Set("ALIGN_LEFT", member)
}

// Align each line of multi-line text to the right.
func (self *RetroFont) GetALIGN_RIGHTA() string{
    return self.Object.Get("ALIGN_RIGHT").String()
}

// Align each line of multi-line text to the right.
func (self *RetroFont) SetALIGN_RIGHTA(member string) {
    self.Object.Set("ALIGN_RIGHT", member)
}

// Align each line of multi-line text in the center.
func (self *RetroFont) GetALIGN_CENTERA() string{
    return self.Object.Get("ALIGN_CENTER").String()
}

// Align each line of multi-line text in the center.
func (self *RetroFont) SetALIGN_CENTERA(member string) {
    self.Object.Set("ALIGN_CENTER", member)
}

// Text Set 1 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
func (self *RetroFont) GetTEXT_SET1A() string{
    return self.Object.Get("TEXT_SET1").String()
}

// Text Set 1 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
func (self *RetroFont) SetTEXT_SET1A(member string) {
    self.Object.Set("TEXT_SET1", member)
}

// Text Set 2 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) GetTEXT_SET2A() string{
    return self.Object.Get("TEXT_SET2").String()
}

// Text Set 2 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) SetTEXT_SET2A(member string) {
    self.Object.Set("TEXT_SET2", member)
}

// Text Set 3 = ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
func (self *RetroFont) GetTEXT_SET3A() string{
    return self.Object.Get("TEXT_SET3").String()
}

// Text Set 3 = ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
func (self *RetroFont) SetTEXT_SET3A(member string) {
    self.Object.Set("TEXT_SET3", member)
}

// Text Set 4 = ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789
func (self *RetroFont) GetTEXT_SET4A() string{
    return self.Object.Get("TEXT_SET4").String()
}

// Text Set 4 = ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789
func (self *RetroFont) SetTEXT_SET4A(member string) {
    self.Object.Set("TEXT_SET4", member)
}

// Text Set 5 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,/() '!?-*:0123456789
func (self *RetroFont) GetTEXT_SET5A() string{
    return self.Object.Get("TEXT_SET5").String()
}

// Text Set 5 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,/() '!?-*:0123456789
func (self *RetroFont) SetTEXT_SET5A(member string) {
    self.Object.Set("TEXT_SET5", member)
}

// Text Set 6 = ABCDEFGHIJKLMNOPQRSTUVWXYZ!?:;0123456789"(),-.'
func (self *RetroFont) GetTEXT_SET6A() string{
    return self.Object.Get("TEXT_SET6").String()
}

// Text Set 6 = ABCDEFGHIJKLMNOPQRSTUVWXYZ!?:;0123456789"(),-.'
func (self *RetroFont) SetTEXT_SET6A(member string) {
    self.Object.Set("TEXT_SET6", member)
}

// Text Set 7 = AGMSY+:4BHNTZ!;5CIOU.?06DJPV,(17EKQW")28FLRX-'39
func (self *RetroFont) GetTEXT_SET7A() string{
    return self.Object.Get("TEXT_SET7").String()
}

// Text Set 7 = AGMSY+:4BHNTZ!;5CIOU.?06DJPV,(17EKQW")28FLRX-'39
func (self *RetroFont) SetTEXT_SET7A(member string) {
    self.Object.Set("TEXT_SET7", member)
}

// Text Set 8 = 0123456789 .ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) GetTEXT_SET8A() string{
    return self.Object.Get("TEXT_SET8").String()
}

// Text Set 8 = 0123456789 .ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) SetTEXT_SET8A(member string) {
    self.Object.Set("TEXT_SET8", member)
}

// Text Set 9 = ABCDEFGHIJKLMNOPQRSTUVWXYZ()-0123456789.:,'"?!
func (self *RetroFont) GetTEXT_SET9A() string{
    return self.Object.Get("TEXT_SET9").String()
}

// Text Set 9 = ABCDEFGHIJKLMNOPQRSTUVWXYZ()-0123456789.:,'"?!
func (self *RetroFont) SetTEXT_SET9A(member string) {
    self.Object.Set("TEXT_SET9", member)
}

// Text Set 10 = ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) GetTEXT_SET10A() string{
    return self.Object.Get("TEXT_SET10").String()
}

// Text Set 10 = ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) SetTEXT_SET10A(member string) {
    self.Object.Set("TEXT_SET10", member)
}

// Text Set 11 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,"-+!?()':;0123456789
func (self *RetroFont) GetTEXT_SET11A() string{
    return self.Object.Get("TEXT_SET11").String()
}

// Text Set 11 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,"-+!?()':;0123456789
func (self *RetroFont) SetTEXT_SET11A(member string) {
    self.Object.Set("TEXT_SET11", member)
}

// Set this value to update the text in this sprite. Carriage returns are automatically stripped out if multiLine is false. Text is converted to upper case if autoUpperCase is true.
func (self *RetroFont) GetTextA() string{
    return self.Object.Get("text").String()
}

// Set this value to update the text in this sprite. Carriage returns are automatically stripped out if multiLine is false. Text is converted to upper case if autoUpperCase is true.
func (self *RetroFont) SetTextA(member string) {
    self.Object.Set("text", member)
}

// Sets if the stamp is smoothed or not.
func (self *RetroFont) GetSmoothedA() bool{
    return self.Object.Get("smoothed").Bool()
}

// Sets if the stamp is smoothed or not.
func (self *RetroFont) SetSmoothedA(member bool) {
    self.Object.Set("smoothed", member)
}

// A reference to the currently running game.
func (self *RetroFont) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *RetroFont) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The key of the RenderTexture in the Cache, if stored there.
func (self *RetroFont) GetKeyA() string{
    return self.Object.Get("key").String()
}

// The key of the RenderTexture in the Cache, if stored there.
func (self *RetroFont) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// The with of the render texture
func (self *RetroFont) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The with of the render texture
func (self *RetroFont) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the render texture
func (self *RetroFont) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the render texture
func (self *RetroFont) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The Resolution of the texture.
func (self *RetroFont) GetResolutionA() int{
    return self.Object.Get("resolution").Int()
}

// The Resolution of the texture.
func (self *RetroFont) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// The framing rectangle of the render texture
func (self *RetroFont) GetFrameA() *Rectangle{
    return &Rectangle{self.Object.Get("frame")}
}

// The framing rectangle of the render texture
func (self *RetroFont) SetFrameA(member *Rectangle) {
    self.Object.Set("frame", member)
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RetroFont) GetCropA() *Rectangle{
    return &Rectangle{self.Object.Get("crop")}
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RetroFont) SetCropA(member *Rectangle) {
    self.Object.Set("crop", member)
}

// The base texture object that this texture uses
func (self *RetroFont) GetBaseTextureA() *BaseTexture{
    return &BaseTexture{self.Object.Get("baseTexture")}
}

// The base texture object that this texture uses
func (self *RetroFont) SetBaseTextureA(member *BaseTexture) {
    self.Object.Set("baseTexture", member)
}

// The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RetroFont) GetRendererA() interface{}{
    return self.Object.Get("renderer")
}

// The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RetroFont) SetRendererA(member interface{}) {
    self.Object.Set("renderer", member)
}

// 
func (self *RetroFont) GetValidA() bool{
    return self.Object.Get("valid").Bool()
}

// 
func (self *RetroFont) SetValidA(member bool) {
    self.Object.Set("valid", member)
}

// Does this Texture have any frame data assigned to it?
func (self *RetroFont) GetNoFrameA() bool{
    return self.Object.Get("noFrame").Bool()
}

// Does this Texture have any frame data assigned to it?
func (self *RetroFont) SetNoFrameA(member bool) {
    self.Object.Set("noFrame", member)
}

// The texture trim data.
func (self *RetroFont) GetTrimA() *Rectangle{
    return &Rectangle{self.Object.Get("trim")}
}

// The texture trim data.
func (self *RetroFont) SetTrimA(member *Rectangle) {
    self.Object.Set("trim", member)
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RetroFont) GetIsTilingA() bool{
    return self.Object.Get("isTiling").Bool()
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RetroFont) SetIsTilingA(member bool) {
    self.Object.Set("isTiling", member)
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RetroFont) GetRequiresUpdateA() bool{
    return self.Object.Get("requiresUpdate").Bool()
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RetroFont) SetRequiresUpdateA(member bool) {
    self.Object.Set("requiresUpdate", member)
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *RetroFont) GetRequiresReTintA() bool{
    return self.Object.Get("requiresReTint").Bool()
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *RetroFont) SetRequiresReTintA(member bool) {
    self.Object.Set("requiresReTint", member)
}



// If you need this RetroFont to have a fixed width and custom alignment you can set the width here.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) SetFixedWidth(width int) {
    self.Object.Call("setFixedWidth", width)
}

// If you need this RetroFont to have a fixed width and custom alignment you can set the width here.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) SetFixedWidth1O(width int, lineAlignment string) {
    self.Object.Call("setFixedWidth", width, lineAlignment)
}

// If you need this RetroFont to have a fixed width and custom alignment you can set the width here.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) SetFixedWidthI(args ...interface{}) {
    self.Object.Call("setFixedWidth", args)
}

// A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText(content string) {
    self.Object.Call("setText", content)
}

// A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText1O(content string, multiLine bool) {
    self.Object.Call("setText", content, multiLine)
}

// A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText2O(content string, multiLine bool, characterSpacing int) {
    self.Object.Call("setText", content, multiLine, characterSpacing)
}

// A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText3O(content string, multiLine bool, characterSpacing int, lineSpacing int) {
    self.Object.Call("setText", content, multiLine, characterSpacing, lineSpacing)
}

// A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText4O(content string, multiLine bool, characterSpacing int, lineSpacing int, lineAlignment string) {
    self.Object.Call("setText", content, multiLine, characterSpacing, lineSpacing, lineAlignment)
}

// A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetText5O(content string, multiLine bool, characterSpacing int, lineSpacing int, lineAlignment string, allowLowerCase bool) {
    self.Object.Call("setText", content, multiLine, characterSpacing, lineSpacing, lineAlignment, allowLowerCase)
}

// A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetTextI(args ...interface{}) {
    self.Object.Call("setText", args)
}

// Updates the texture with the new text.
func (self *RetroFont) BuildRetroFontText() {
    self.Object.Call("buildRetroFontText")
}

// Updates the texture with the new text.
func (self *RetroFont) BuildRetroFontTextI(args ...interface{}) {
    self.Object.Call("buildRetroFontText", args)
}

// Internal function that takes a single line of text (2nd parameter) and pastes it into the BitmapData at the given coordinates.
// Used by getLine and getMultiLine
func (self *RetroFont) PasteLine(line string, x int, y int, customSpacingX int) {
    self.Object.Call("pasteLine", line, x, y, customSpacingX)
}

// Internal function that takes a single line of text (2nd parameter) and pastes it into the BitmapData at the given coordinates.
// Used by getLine and getMultiLine
func (self *RetroFont) PasteLineI(args ...interface{}) {
    self.Object.Call("pasteLine", args)
}

// Works out the longest line of text in _text and returns its length
func (self *RetroFont) GetLongestLine() int{
    return self.Object.Call("getLongestLine").Int()
}

// Works out the longest line of text in _text and returns its length
func (self *RetroFont) GetLongestLineI(args ...interface{}) int{
    return self.Object.Call("getLongestLine", args).Int()
}

// Internal helper function that removes all unsupported characters from the _text String, leaving only characters contained in the font set.
func (self *RetroFont) RemoveUnsupportedCharacters() string{
    return self.Object.Call("removeUnsupportedCharacters").String()
}

// Internal helper function that removes all unsupported characters from the _text String, leaving only characters contained in the font set.
func (self *RetroFont) RemoveUnsupportedCharacters1O(stripCR bool) string{
    return self.Object.Call("removeUnsupportedCharacters", stripCR).String()
}

// Internal helper function that removes all unsupported characters from the _text String, leaving only characters contained in the font set.
func (self *RetroFont) RemoveUnsupportedCharactersI(args ...interface{}) string{
    return self.Object.Call("removeUnsupportedCharacters", args).String()
}

// Updates the x and/or y offset that the font is rendered from. This updates all of the texture frames, so be careful how often it is called.
// Note that the values given for the x and y properties are either ADDED to or SUBTRACTED from (if negative) the existing offsetX/Y values of the characters.
// So if the current offsetY is 8 and you want it to start rendering from y16 you would call updateOffset(0, 8) to add 8 to the current y offset.
func (self *RetroFont) UpdateOffset() {
    self.Object.Call("updateOffset")
}

// Updates the x and/or y offset that the font is rendered from. This updates all of the texture frames, so be careful how often it is called.
// Note that the values given for the x and y properties are either ADDED to or SUBTRACTED from (if negative) the existing offsetX/Y values of the characters.
// So if the current offsetY is 8 and you want it to start rendering from y16 you would call updateOffset(0, 8) to add 8 to the current y offset.
func (self *RetroFont) UpdateOffset1O(xOffset int) {
    self.Object.Call("updateOffset", xOffset)
}

// Updates the x and/or y offset that the font is rendered from. This updates all of the texture frames, so be careful how often it is called.
// Note that the values given for the x and y properties are either ADDED to or SUBTRACTED from (if negative) the existing offsetX/Y values of the characters.
// So if the current offsetY is 8 and you want it to start rendering from y16 you would call updateOffset(0, 8) to add 8 to the current y offset.
func (self *RetroFont) UpdateOffset2O(xOffset int, yOffset int) {
    self.Object.Call("updateOffset", xOffset, yOffset)
}

// Updates the x and/or y offset that the font is rendered from. This updates all of the texture frames, so be careful how often it is called.
// Note that the values given for the x and y properties are either ADDED to or SUBTRACTED from (if negative) the existing offsetX/Y values of the characters.
// So if the current offsetY is 8 and you want it to start rendering from y16 you would call updateOffset(0, 8) to add 8 to the current y offset.
func (self *RetroFont) UpdateOffsetI(args ...interface{}) {
    self.Object.Call("updateOffset", args)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RetroFont) RenderXY(displayObject interface{}, x int, y int) {
    self.Object.Call("renderXY", displayObject, x, y)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RetroFont) RenderXY1O(displayObject interface{}, x int, y int, clear bool) {
    self.Object.Call("renderXY", displayObject, x, y, clear)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RetroFont) RenderXYI(args ...interface{}) {
    self.Object.Call("renderXY", args)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RetroFont) RenderRawXY(displayObject interface{}, x int, y int) {
    self.Object.Call("renderRawXY", displayObject, x, y)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RetroFont) RenderRawXY1O(displayObject interface{}, x int, y int, clear bool) {
    self.Object.Call("renderRawXY", displayObject, x, y, clear)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RetroFont) RenderRawXYI(args ...interface{}) {
    self.Object.Call("renderRawXY", args)
}

// This function will draw the display object to the RenderTexture.
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

// This function will draw the display object to the RenderTexture.
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

// This function will draw the display object to the RenderTexture.
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

// This function will draw the display object to the RenderTexture.
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

// Resizes the RenderTexture.
func (self *RetroFont) Resize(width int, height int, updateBase bool) {
    self.Object.Call("resize", width, height, updateBase)
}

// Resizes the RenderTexture.
func (self *RetroFont) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Clears the RenderTexture.
func (self *RetroFont) Clear() {
    self.Object.Call("clear")
}

// Clears the RenderTexture.
func (self *RetroFont) ClearI(args ...interface{}) {
    self.Object.Call("clear", args)
}

// This function will draw the display object to the texture.
func (self *RetroFont) RenderWebGL(displayObject *DisplayObject) {
    self.Object.Call("renderWebGL", displayObject)
}

// This function will draw the display object to the texture.
func (self *RetroFont) RenderWebGL1O(displayObject *DisplayObject, matrix *Matrix) {
    self.Object.Call("renderWebGL", displayObject, matrix)
}

// This function will draw the display object to the texture.
func (self *RetroFont) RenderWebGL2O(displayObject *DisplayObject, matrix *Matrix, clear bool) {
    self.Object.Call("renderWebGL", displayObject, matrix, clear)
}

// This function will draw the display object to the texture.
func (self *RetroFont) RenderWebGLI(args ...interface{}) {
    self.Object.Call("renderWebGL", args)
}

// This function will draw the display object to the texture.
func (self *RetroFont) RenderCanvas(displayObject *DisplayObject) {
    self.Object.Call("renderCanvas", displayObject)
}

// This function will draw the display object to the texture.
func (self *RetroFont) RenderCanvas1O(displayObject *DisplayObject, matrix *Matrix) {
    self.Object.Call("renderCanvas", displayObject, matrix)
}

// This function will draw the display object to the texture.
func (self *RetroFont) RenderCanvas2O(displayObject *DisplayObject, matrix *Matrix, clear bool) {
    self.Object.Call("renderCanvas", displayObject, matrix, clear)
}

// This function will draw the display object to the texture.
func (self *RetroFont) RenderCanvasI(args ...interface{}) {
    self.Object.Call("renderCanvas", args)
}

// Will return a HTML Image of the texture
func (self *RetroFont) GetImage() *Image{
    return &Image{self.Object.Call("getImage")}
}

// Will return a HTML Image of the texture
func (self *RetroFont) GetImageI(args ...interface{}) *Image{
    return &Image{self.Object.Call("getImage", args)}
}

// Will return a base64 encoded string of this texture. It works by calling RenderTexture.getCanvas and then running toDataURL on that.
func (self *RetroFont) GetBase64() string{
    return self.Object.Call("getBase64").String()
}

// Will return a base64 encoded string of this texture. It works by calling RenderTexture.getCanvas and then running toDataURL on that.
func (self *RetroFont) GetBase64I(args ...interface{}) string{
    return self.Object.Call("getBase64", args).String()
}

// Creates a Canvas element, renders this RenderTexture to it and then returns it.
func (self *RetroFont) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("getCanvas"))
}

// Creates a Canvas element, renders this RenderTexture to it and then returns it.
func (self *RetroFont) GetCanvasI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("getCanvas", args))
}

// Called when the base texture is loaded
func (self *RetroFont) OnBaseTextureLoaded() {
    self.Object.Call("onBaseTextureLoaded")
}

// Called when the base texture is loaded
func (self *RetroFont) OnBaseTextureLoadedI(args ...interface{}) {
    self.Object.Call("onBaseTextureLoaded", args)
}

// Destroys this texture
func (self *RetroFont) Destroy(destroyBase bool) {
    self.Object.Call("destroy", destroyBase)
}

// Destroys this texture
func (self *RetroFont) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Specifies the region of the baseTexture that this texture will use.
func (self *RetroFont) SetFrame(frame *Rectangle) {
    self.Object.Call("setFrame", frame)
}

// Specifies the region of the baseTexture that this texture will use.
func (self *RetroFont) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// Updates the internal WebGL UV cache.
func (self *RetroFont) _updateUvs() {
    self.Object.Call("_updateUvs")
}

// Updates the internal WebGL UV cache.
func (self *RetroFont) _updateUvsI(args ...interface{}) {
    self.Object.Call("_updateUvs", args)
}
