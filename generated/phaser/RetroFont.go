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
func (self *RetroFont) GetCharacterWidth() float64{
    return self.Get("characterWidth").Float()
}

// The width of each character in the font set.
func (self *RetroFont) SetCharacterWidth(member float64) {
    self.Set("characterWidth", member)
}

// The height of each character in the font set.
func (self *RetroFont) GetCharacterHeight() float64{
    return self.Get("characterHeight").Float()
}

// The height of each character in the font set.
func (self *RetroFont) SetCharacterHeight(member float64) {
    self.Set("characterHeight", member)
}

// If the characters in the font set have horizontal spacing between them set the required amount here.
func (self *RetroFont) GetCharacterSpacingX() float64{
    return self.Get("characterSpacingX").Float()
}

// If the characters in the font set have horizontal spacing between them set the required amount here.
func (self *RetroFont) SetCharacterSpacingX(member float64) {
    self.Set("characterSpacingX", member)
}

// If the characters in the font set have vertical spacing between them set the required amount here.
func (self *RetroFont) GetCharacterSpacingY() float64{
    return self.Get("characterSpacingY").Float()
}

// If the characters in the font set have vertical spacing between them set the required amount here.
func (self *RetroFont) SetCharacterSpacingY(member float64) {
    self.Set("characterSpacingY", member)
}

// The number of characters per row in the font set.
func (self *RetroFont) GetCharacterPerRow() float64{
    return self.Get("characterPerRow").Float()
}

// The number of characters per row in the font set.
func (self *RetroFont) SetCharacterPerRow(member float64) {
    self.Set("characterPerRow", member)
}

// If the font set doesn't start at the top left of the given image, specify the X coordinate offset here.
func (self *RetroFont) GetOffsetX() float64{
    return self.Get("offsetX").Float()
}

// If the font set doesn't start at the top left of the given image, specify the X coordinate offset here.
func (self *RetroFont) SetOffsetX(member float64) {
    self.Set("offsetX", member)
}

// If the font set doesn't start at the top left of the given image, specify the Y coordinate offset here.
func (self *RetroFont) GetOffsetY() float64{
    return self.Get("offsetY").Float()
}

// If the font set doesn't start at the top left of the given image, specify the Y coordinate offset here.
func (self *RetroFont) SetOffsetY(member float64) {
    self.Set("offsetY", member)
}

// Alignment of the text when multiLine = true or a fixedWidth is set. Set to RetroFont.ALIGN_LEFT (default), RetroFont.ALIGN_RIGHT or RetroFont.ALIGN_CENTER.
func (self *RetroFont) GetAlign() string{
    return self.Get("align").String()
}

// Alignment of the text when multiLine = true or a fixedWidth is set. Set to RetroFont.ALIGN_LEFT (default), RetroFont.ALIGN_RIGHT or RetroFont.ALIGN_CENTER.
func (self *RetroFont) SetAlign(member string) {
    self.Set("align", member)
}

// If set to true all carriage-returns in text will form new lines (see align). If false the font will only contain one single line of text (the default)
func (self *RetroFont) GetMultiLine() bool{
    return self.Get("multiLine").Bool()
}

// If set to true all carriage-returns in text will form new lines (see align). If false the font will only contain one single line of text (the default)
func (self *RetroFont) SetMultiLine(member bool) {
    self.Set("multiLine", member)
}

// Automatically convert any text to upper case. Lots of old bitmap fonts only contain upper-case characters, so the default is true.
func (self *RetroFont) GetAutoUpperCase() bool{
    return self.Get("autoUpperCase").Bool()
}

// Automatically convert any text to upper case. Lots of old bitmap fonts only contain upper-case characters, so the default is true.
func (self *RetroFont) SetAutoUpperCase(member bool) {
    self.Set("autoUpperCase", member)
}

// Adds horizontal spacing between each character of the font, in pixels.
func (self *RetroFont) GetCustomSpacingX() float64{
    return self.Get("customSpacingX").Float()
}

// Adds horizontal spacing between each character of the font, in pixels.
func (self *RetroFont) SetCustomSpacingX(member float64) {
    self.Set("customSpacingX", member)
}

// Adds vertical spacing between each line of multi-line text, set in pixels.
func (self *RetroFont) GetCustomSpacingY() float64{
    return self.Get("customSpacingY").Float()
}

// Adds vertical spacing between each line of multi-line text, set in pixels.
func (self *RetroFont) SetCustomSpacingY(member float64) {
    self.Set("customSpacingY", member)
}

// If you need this RetroFont image to have a fixed width you can set the width in this value.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) GetFixedWidth() float64{
    return self.Get("fixedWidth").Float()
}

// If you need this RetroFont image to have a fixed width you can set the width in this value.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) SetFixedWidth(member float64) {
    self.Set("fixedWidth", member)
}

// A reference to the image stored in the Game.Cache that contains the font.
func (self *RetroFont) GetFontSet() *Image{
    return &Image{self.Get("fontSet")}
}

// A reference to the image stored in the Game.Cache that contains the font.
func (self *RetroFont) SetFontSet(member *Image) {
    self.Set("fontSet", member)
}

// The FrameData representing this Retro Font.
func (self *RetroFont) GetFrameData() *FrameData{
    return &FrameData{self.Get("frameData")}
}

// The FrameData representing this Retro Font.
func (self *RetroFont) SetFrameData(member *FrameData) {
    self.Set("frameData", member)
}

// The image that is stamped to the RenderTexture for each character in the font.
func (self *RetroFont) GetStamp() *Image{
    return &Image{self.Get("stamp")}
}

// The image that is stamped to the RenderTexture for each character in the font.
func (self *RetroFont) SetStamp(member *Image) {
    self.Set("stamp", member)
}

// Base Phaser object type.
func (self *RetroFont) GetType() float64{
    return self.Get("type").Float()
}

// Base Phaser object type.
func (self *RetroFont) SetType(member float64) {
    self.Set("type", member)
}

// Align each line of multi-line text to the left.
func (self *RetroFont) GetALIGN_LEFT() string{
    return self.Get("ALIGN_LEFT").String()
}

// Align each line of multi-line text to the left.
func (self *RetroFont) SetALIGN_LEFT(member string) {
    self.Set("ALIGN_LEFT", member)
}

// Align each line of multi-line text to the right.
func (self *RetroFont) GetALIGN_RIGHT() string{
    return self.Get("ALIGN_RIGHT").String()
}

// Align each line of multi-line text to the right.
func (self *RetroFont) SetALIGN_RIGHT(member string) {
    self.Set("ALIGN_RIGHT", member)
}

// Align each line of multi-line text in the center.
func (self *RetroFont) GetALIGN_CENTER() string{
    return self.Get("ALIGN_CENTER").String()
}

// Align each line of multi-line text in the center.
func (self *RetroFont) SetALIGN_CENTER(member string) {
    self.Set("ALIGN_CENTER", member)
}

// Text Set 1 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
func (self *RetroFont) GetTEXT_SET1() string{
    return self.Get("TEXT_SET1").String()
}

// Text Set 1 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
func (self *RetroFont) SetTEXT_SET1(member string) {
    self.Set("TEXT_SET1", member)
}

// Text Set 2 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) GetTEXT_SET2() string{
    return self.Get("TEXT_SET2").String()
}

// Text Set 2 =  !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) SetTEXT_SET2(member string) {
    self.Set("TEXT_SET2", member)
}

// Text Set 3 = ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
func (self *RetroFont) GetTEXT_SET3() string{
    return self.Get("TEXT_SET3").String()
}

// Text Set 3 = ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
func (self *RetroFont) SetTEXT_SET3(member string) {
    self.Set("TEXT_SET3", member)
}

// Text Set 4 = ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789
func (self *RetroFont) GetTEXT_SET4() string{
    return self.Get("TEXT_SET4").String()
}

// Text Set 4 = ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789
func (self *RetroFont) SetTEXT_SET4(member string) {
    self.Set("TEXT_SET4", member)
}

// Text Set 5 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,/() '!?-*:0123456789
func (self *RetroFont) GetTEXT_SET5() string{
    return self.Get("TEXT_SET5").String()
}

// Text Set 5 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,/() '!?-*:0123456789
func (self *RetroFont) SetTEXT_SET5(member string) {
    self.Set("TEXT_SET5", member)
}

// Text Set 6 = ABCDEFGHIJKLMNOPQRSTUVWXYZ!?:;0123456789"(),-.'
func (self *RetroFont) GetTEXT_SET6() string{
    return self.Get("TEXT_SET6").String()
}

// Text Set 6 = ABCDEFGHIJKLMNOPQRSTUVWXYZ!?:;0123456789"(),-.'
func (self *RetroFont) SetTEXT_SET6(member string) {
    self.Set("TEXT_SET6", member)
}

// Text Set 7 = AGMSY+:4BHNTZ!;5CIOU.?06DJPV,(17EKQW")28FLRX-'39
func (self *RetroFont) GetTEXT_SET7() string{
    return self.Get("TEXT_SET7").String()
}

// Text Set 7 = AGMSY+:4BHNTZ!;5CIOU.?06DJPV,(17EKQW")28FLRX-'39
func (self *RetroFont) SetTEXT_SET7(member string) {
    self.Set("TEXT_SET7", member)
}

// Text Set 8 = 0123456789 .ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) GetTEXT_SET8() string{
    return self.Get("TEXT_SET8").String()
}

// Text Set 8 = 0123456789 .ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) SetTEXT_SET8(member string) {
    self.Set("TEXT_SET8", member)
}

// Text Set 9 = ABCDEFGHIJKLMNOPQRSTUVWXYZ()-0123456789.:,'"?!
func (self *RetroFont) GetTEXT_SET9() string{
    return self.Get("TEXT_SET9").String()
}

// Text Set 9 = ABCDEFGHIJKLMNOPQRSTUVWXYZ()-0123456789.:,'"?!
func (self *RetroFont) SetTEXT_SET9(member string) {
    self.Set("TEXT_SET9", member)
}

// Text Set 10 = ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) GetTEXT_SET10() string{
    return self.Get("TEXT_SET10").String()
}

// Text Set 10 = ABCDEFGHIJKLMNOPQRSTUVWXYZ
func (self *RetroFont) SetTEXT_SET10(member string) {
    self.Set("TEXT_SET10", member)
}

// Text Set 11 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,"-+!?()':;0123456789
func (self *RetroFont) GetTEXT_SET11() string{
    return self.Get("TEXT_SET11").String()
}

// Text Set 11 = ABCDEFGHIJKLMNOPQRSTUVWXYZ.,"-+!?()':;0123456789
func (self *RetroFont) SetTEXT_SET11(member string) {
    self.Set("TEXT_SET11", member)
}

// Set this value to update the text in this sprite. Carriage returns are automatically stripped out if multiLine is false. Text is converted to upper case if autoUpperCase is true.
func (self *RetroFont) GetText() string{
    return self.Get("text").String()
}

// Set this value to update the text in this sprite. Carriage returns are automatically stripped out if multiLine is false. Text is converted to upper case if autoUpperCase is true.
func (self *RetroFont) SetText(member string) {
    self.Set("text", member)
}

// Sets if the stamp is smoothed or not.
func (self *RetroFont) GetSmoothed() bool{
    return self.Get("smoothed").Bool()
}

// Sets if the stamp is smoothed or not.
func (self *RetroFont) SetSmoothed(member bool) {
    self.Set("smoothed", member)
}

// A reference to the currently running game.
func (self *RetroFont) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *RetroFont) SetGame(member *Game) {
    self.Set("game", member)
}

// The key of the RenderTexture in the Cache, if stored there.
func (self *RetroFont) GetKey() string{
    return self.Get("key").String()
}

// The key of the RenderTexture in the Cache, if stored there.
func (self *RetroFont) SetKey(member string) {
    self.Set("key", member)
}

// The with of the render texture
func (self *RetroFont) GetWidth() float64{
    return self.Get("width").Float()
}

// The with of the render texture
func (self *RetroFont) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the render texture
func (self *RetroFont) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the render texture
func (self *RetroFont) SetHeight(member float64) {
    self.Set("height", member)
}

// The Resolution of the texture.
func (self *RetroFont) GetResolution() float64{
    return self.Get("resolution").Float()
}

// The Resolution of the texture.
func (self *RetroFont) SetResolution(member float64) {
    self.Set("resolution", member)
}

// The framing rectangle of the render texture
func (self *RetroFont) GetFrame() *Rectangle{
    return &Rectangle{self.Get("frame")}
}

// The framing rectangle of the render texture
func (self *RetroFont) SetFrame(member *Rectangle) {
    self.Set("frame", member)
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RetroFont) GetCrop() *Rectangle{
    return &Rectangle{self.Get("crop")}
}

// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering,
// irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
func (self *RetroFont) SetCrop(member *Rectangle) {
    self.Set("crop", member)
}

// The base texture object that this texture uses
func (self *RetroFont) GetBaseTexture() *BaseTexture{
    return &BaseTexture{self.Get("baseTexture")}
}

// The base texture object that this texture uses
func (self *RetroFont) SetBaseTexture(member *BaseTexture) {
    self.Set("baseTexture", member)
}

// The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RetroFont) GetRenderer() interface{}{
    return self.Get("renderer")
}

// The renderer this RenderTexture uses. A RenderTexture can only belong to one renderer at the moment if its webGL.
func (self *RetroFont) SetRenderer(member interface{}) {
    self.Set("renderer", member)
}

// 
func (self *RetroFont) GetValid() bool{
    return self.Get("valid").Bool()
}

// 
func (self *RetroFont) SetValid(member bool) {
    self.Set("valid", member)
}

// Does this Texture have any frame data assigned to it?
func (self *RetroFont) GetNoFrame() bool{
    return self.Get("noFrame").Bool()
}

// Does this Texture have any frame data assigned to it?
func (self *RetroFont) SetNoFrame(member bool) {
    self.Set("noFrame", member)
}

// The texture trim data.
func (self *RetroFont) GetTrim() *Rectangle{
    return &Rectangle{self.Get("trim")}
}

// The texture trim data.
func (self *RetroFont) SetTrim(member *Rectangle) {
    self.Set("trim", member)
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RetroFont) GetIsTiling() bool{
    return self.Get("isTiling").Bool()
}

// Is this a tiling texture? As used by the likes of a TilingSprite.
func (self *RetroFont) SetIsTiling(member bool) {
    self.Set("isTiling", member)
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RetroFont) GetRequiresUpdate() bool{
    return self.Get("requiresUpdate").Bool()
}

// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
func (self *RetroFont) SetRequiresUpdate(member bool) {
    self.Set("requiresUpdate", member)
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *RetroFont) GetRequiresReTint() bool{
    return self.Get("requiresReTint").Bool()
}

// This will let a renderer know that a tinted parent has updated its texture.
func (self *RetroFont) SetRequiresReTint(member bool) {
    self.Set("requiresReTint", member)
}



// If you need this RetroFont to have a fixed width and custom alignment you can set the width here.
// If text is wider than the width specified it will be cropped off.
func (self *RetroFont) SetFixedWidthI(args ...interface{}) {
    self.Call("setFixedWidth", args)
}

// A helper function that quickly sets lots of variables at once, and then updates the text.
func (self *RetroFont) SetTextI(args ...interface{}) {
    self.Call("setText", args)
}

// Updates the texture with the new text.
func (self *RetroFont) BuildRetroFontTextI(args ...interface{}) {
    self.Call("buildRetroFontText", args)
}

// Internal function that takes a single line of text (2nd parameter) and pastes it into the BitmapData at the given coordinates.
// Used by getLine and getMultiLine
func (self *RetroFont) PasteLineI(args ...interface{}) {
    self.Call("pasteLine", args)
}

// Works out the longest line of text in _text and returns its length
func (self *RetroFont) GetLongestLineI(args ...interface{}) float64{
    return self.Call("getLongestLine", args).Float()
}

// Internal helper function that removes all unsupported characters from the _text String, leaving only characters contained in the font set.
func (self *RetroFont) RemoveUnsupportedCharactersI(args ...interface{}) string{
    return self.Call("removeUnsupportedCharacters", args).String()
}

// Updates the x and/or y offset that the font is rendered from. This updates all of the texture frames, so be careful how often it is called.
// Note that the values given for the x and y properties are either ADDED to or SUBTRACTED from (if negative) the existing offsetX/Y values of the characters.
// So if the current offsetY is 8 and you want it to start rendering from y16 you would call updateOffset(0, 8) to add 8 to the current y offset.
func (self *RetroFont) UpdateOffsetI(args ...interface{}) {
    self.Call("updateOffset", args)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it takes into account scale and rotation.
// 
// If you don't want those then use RenderTexture.renderRawXY instead.
func (self *RetroFont) RenderXYI(args ...interface{}) {
    self.Call("renderXY", args)
}

// This function will draw the display object to the RenderTexture at the given coordinates.
// 
// When the display object is drawn it doesn't take into account scale, rotation or translation.
// 
// If you need those then use RenderTexture.renderXY instead.
func (self *RetroFont) RenderRawXYI(args ...interface{}) {
    self.Call("renderRawXY", args)
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
    self.Call("render", args)
}

// Resizes the RenderTexture.
func (self *RetroFont) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// Clears the RenderTexture.
func (self *RetroFont) ClearI(args ...interface{}) {
    self.Call("clear", args)
}

// This function will draw the display object to the texture.
func (self *RetroFont) RenderWebGLI(args ...interface{}) {
    self.Call("renderWebGL", args)
}

// This function will draw the display object to the texture.
func (self *RetroFont) RenderCanvasI(args ...interface{}) {
    self.Call("renderCanvas", args)
}

// Will return a HTML Image of the texture
func (self *RetroFont) GetImageI(args ...interface{}) *Image{
    return &Image{self.Call("getImage", args)}
}

// Will return a base64 encoded string of this texture. It works by calling RenderTexture.getCanvas and then running toDataURL on that.
func (self *RetroFont) GetBase64I(args ...interface{}) string{
    return self.Call("getBase64", args).String()
}

// Creates a Canvas element, renders this RenderTexture to it and then returns it.
func (self *RetroFont) GetCanvasI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("getCanvas", args))
}

// Called when the base texture is loaded
func (self *RetroFont) OnBaseTextureLoadedI(args ...interface{}) {
    self.Call("onBaseTextureLoaded", args)
}

// Destroys this texture
func (self *RetroFont) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Specifies the region of the baseTexture that this texture will use.
func (self *RetroFont) SetFrameI(args ...interface{}) {
    self.Call("setFrame", args)
}

// Updates the internal WebGL UV cache.
func (self *RetroFont) _updateUvsI(args ...interface{}) {
    self.Call("_updateUvs", args)
}
