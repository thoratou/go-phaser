// Package phaser Automatic generation for Phaser.BitmapText
// generated file BitmapText.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// BitmapText BitmapText objects work by taking a texture file and an XML or JSON file that describes the font structure.
// It then generates a new Sprite object for each letter of the text, proportionally spaced out and aligned to 
// match the font structure.
// 
// BitmapText objects are less flexible than Text objects, in that they have less features such as shadows, fills and the ability 
// to use Web Fonts, however you trade this flexibility for rendering speed. You can also create visually compelling BitmapTexts by
// processing the font texture in an image editor, applying fills and any other effects required.
// 
// To create multi-line text insert \r, \n or \r\n escape codes into the text string.
// 
// If you are having performance issues due to the volume of sprites being rendered, and do not require the text to be constantly
// updating, you can use BitmapText.generateTexture to create a static texture from this BitmapText.
// 
// To create a BitmapText data files you can use:
// 
// BMFont (Windows, free): http://www.angelcode.com/products/bmfont/
// Glyph Designer (OS X, commercial): http://www.71squared.com/en/glyphdesigner
// Littera (Web-based, free): http://kvazars.com/littera/
// 
// For most use cases it is recommended to use XML. If you wish to use JSON, the formatting should be equal to the result of
// converting a valid XML file through the popular X2JS library. An online tool for conversion can be found here: http://codebeautify.org/xmltojson
// 
// If you were using an older version of Phaser (< 2.4) and using the DOMish parser hack, please remove this. It isn't required any longer.
type BitmapText struct {
    *js.Object
}

// NewBitmapText BitmapText objects work by taking a texture file and an XML or JSON file that describes the font structure.
// It then generates a new Sprite object for each letter of the text, proportionally spaced out and aligned to 
// match the font structure.
// 
// BitmapText objects are less flexible than Text objects, in that they have less features such as shadows, fills and the ability 
// to use Web Fonts, however you trade this flexibility for rendering speed. You can also create visually compelling BitmapTexts by
// processing the font texture in an image editor, applying fills and any other effects required.
// 
// To create multi-line text insert \r, \n or \r\n escape codes into the text string.
// 
// If you are having performance issues due to the volume of sprites being rendered, and do not require the text to be constantly
// updating, you can use BitmapText.generateTexture to create a static texture from this BitmapText.
// 
// To create a BitmapText data files you can use:
// 
// BMFont (Windows, free): http://www.angelcode.com/products/bmfont/
// Glyph Designer (OS X, commercial): http://www.71squared.com/en/glyphdesigner
// Littera (Web-based, free): http://kvazars.com/littera/
// 
// For most use cases it is recommended to use XML. If you wish to use JSON, the formatting should be equal to the result of
// converting a valid XML file through the popular X2JS library. An online tool for conversion can be found here: http://codebeautify.org/xmltojson
// 
// If you were using an older version of Phaser (< 2.4) and using the DOMish parser hack, please remove this. It isn't required any longer.
func NewBitmapText(game *Game, x int, y int, font string) *BitmapText {
    return &BitmapText{js.Global.Get("Phaser").Get("BitmapText").New(game, x, y, font)}
}
// NewBitmapText1O BitmapText objects work by taking a texture file and an XML or JSON file that describes the font structure.
// It then generates a new Sprite object for each letter of the text, proportionally spaced out and aligned to 
// match the font structure.
// 
// BitmapText objects are less flexible than Text objects, in that they have less features such as shadows, fills and the ability 
// to use Web Fonts, however you trade this flexibility for rendering speed. You can also create visually compelling BitmapTexts by
// processing the font texture in an image editor, applying fills and any other effects required.
// 
// To create multi-line text insert \r, \n or \r\n escape codes into the text string.
// 
// If you are having performance issues due to the volume of sprites being rendered, and do not require the text to be constantly
// updating, you can use BitmapText.generateTexture to create a static texture from this BitmapText.
// 
// To create a BitmapText data files you can use:
// 
// BMFont (Windows, free): http://www.angelcode.com/products/bmfont/
// Glyph Designer (OS X, commercial): http://www.71squared.com/en/glyphdesigner
// Littera (Web-based, free): http://kvazars.com/littera/
// 
// For most use cases it is recommended to use XML. If you wish to use JSON, the formatting should be equal to the result of
// converting a valid XML file through the popular X2JS library. An online tool for conversion can be found here: http://codebeautify.org/xmltojson
// 
// If you were using an older version of Phaser (< 2.4) and using the DOMish parser hack, please remove this. It isn't required any longer.
func NewBitmapText1O(game *Game, x int, y int, font string, text string) *BitmapText {
    return &BitmapText{js.Global.Get("Phaser").Get("BitmapText").New(game, x, y, font, text)}
}
// NewBitmapText2O BitmapText objects work by taking a texture file and an XML or JSON file that describes the font structure.
// It then generates a new Sprite object for each letter of the text, proportionally spaced out and aligned to 
// match the font structure.
// 
// BitmapText objects are less flexible than Text objects, in that they have less features such as shadows, fills and the ability 
// to use Web Fonts, however you trade this flexibility for rendering speed. You can also create visually compelling BitmapTexts by
// processing the font texture in an image editor, applying fills and any other effects required.
// 
// To create multi-line text insert \r, \n or \r\n escape codes into the text string.
// 
// If you are having performance issues due to the volume of sprites being rendered, and do not require the text to be constantly
// updating, you can use BitmapText.generateTexture to create a static texture from this BitmapText.
// 
// To create a BitmapText data files you can use:
// 
// BMFont (Windows, free): http://www.angelcode.com/products/bmfont/
// Glyph Designer (OS X, commercial): http://www.71squared.com/en/glyphdesigner
// Littera (Web-based, free): http://kvazars.com/littera/
// 
// For most use cases it is recommended to use XML. If you wish to use JSON, the formatting should be equal to the result of
// converting a valid XML file through the popular X2JS library. An online tool for conversion can be found here: http://codebeautify.org/xmltojson
// 
// If you were using an older version of Phaser (< 2.4) and using the DOMish parser hack, please remove this. It isn't required any longer.
func NewBitmapText2O(game *Game, x int, y int, font string, text string, size int) *BitmapText {
    return &BitmapText{js.Global.Get("Phaser").Get("BitmapText").New(game, x, y, font, text, size)}
}
// NewBitmapText3O BitmapText objects work by taking a texture file and an XML or JSON file that describes the font structure.
// It then generates a new Sprite object for each letter of the text, proportionally spaced out and aligned to 
// match the font structure.
// 
// BitmapText objects are less flexible than Text objects, in that they have less features such as shadows, fills and the ability 
// to use Web Fonts, however you trade this flexibility for rendering speed. You can also create visually compelling BitmapTexts by
// processing the font texture in an image editor, applying fills and any other effects required.
// 
// To create multi-line text insert \r, \n or \r\n escape codes into the text string.
// 
// If you are having performance issues due to the volume of sprites being rendered, and do not require the text to be constantly
// updating, you can use BitmapText.generateTexture to create a static texture from this BitmapText.
// 
// To create a BitmapText data files you can use:
// 
// BMFont (Windows, free): http://www.angelcode.com/products/bmfont/
// Glyph Designer (OS X, commercial): http://www.71squared.com/en/glyphdesigner
// Littera (Web-based, free): http://kvazars.com/littera/
// 
// For most use cases it is recommended to use XML. If you wish to use JSON, the formatting should be equal to the result of
// converting a valid XML file through the popular X2JS library. An online tool for conversion can be found here: http://codebeautify.org/xmltojson
// 
// If you were using an older version of Phaser (< 2.4) and using the DOMish parser hack, please remove this. It isn't required any longer.
func NewBitmapText3O(game *Game, x int, y int, font string, text string, size int, align string) *BitmapText {
    return &BitmapText{js.Global.Get("Phaser").Get("BitmapText").New(game, x, y, font, text, size, align)}
}
// NewBitmapTextI BitmapText objects work by taking a texture file and an XML or JSON file that describes the font structure.
// It then generates a new Sprite object for each letter of the text, proportionally spaced out and aligned to 
// match the font structure.
// 
// BitmapText objects are less flexible than Text objects, in that they have less features such as shadows, fills and the ability 
// to use Web Fonts, however you trade this flexibility for rendering speed. You can also create visually compelling BitmapTexts by
// processing the font texture in an image editor, applying fills and any other effects required.
// 
// To create multi-line text insert \r, \n or \r\n escape codes into the text string.
// 
// If you are having performance issues due to the volume of sprites being rendered, and do not require the text to be constantly
// updating, you can use BitmapText.generateTexture to create a static texture from this BitmapText.
// 
// To create a BitmapText data files you can use:
// 
// BMFont (Windows, free): http://www.angelcode.com/products/bmfont/
// Glyph Designer (OS X, commercial): http://www.71squared.com/en/glyphdesigner
// Littera (Web-based, free): http://kvazars.com/littera/
// 
// For most use cases it is recommended to use XML. If you wish to use JSON, the formatting should be equal to the result of
// converting a valid XML file through the popular X2JS library. An online tool for conversion can be found here: http://codebeautify.org/xmltojson
// 
// If you were using an older version of Phaser (< 2.4) and using the DOMish parser hack, please remove this. It isn't required any longer.
func NewBitmapTextI(args ...interface{}) *BitmapText {
    return &BitmapText{js.Global.Get("Phaser").Get("BitmapText").New(args)}
}



// BitmapText Binding conversion method to BitmapText point 
func ToBitmapText(jsStruct interface{}) *BitmapText {
    if object, ok := jsStruct.(*js.Object); ok {
		return &BitmapText{Object: object}
	}
	return nil
}



// Type The const type of this object.
func (self *BitmapText) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *BitmapText) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// PhysicsType The const physics body type of this object.
func (self *BitmapText) PhysicsType() int{
    return self.Object.Get("physicsType").Int()
}

// SetPhysicsTypeA The const physics body type of this object.
func (self *BitmapText) SetPhysicsTypeA(member int) {
    self.Object.Set("physicsType", member)
}

// TextWidth The width in pixels of the overall text area, taking into consideration multi-line text.
func (self *BitmapText) TextWidth() int{
    return self.Object.Get("textWidth").Int()
}

// SetTextWidthA The width in pixels of the overall text area, taking into consideration multi-line text.
func (self *BitmapText) SetTextWidthA(member int) {
    self.Object.Set("textWidth", member)
}

// TextHeight The height in pixels of the overall text area, taking into consideration multi-line text.
func (self *BitmapText) TextHeight() int{
    return self.Object.Get("textHeight").Int()
}

// SetTextHeightA The height in pixels of the overall text area, taking into consideration multi-line text.
func (self *BitmapText) SetTextHeightA(member int) {
    self.Object.Set("textHeight", member)
}

// Anchor The anchor value of this BitmapText.
func (self *BitmapText) Anchor() *Point{
    return &Point{self.Object.Get("anchor")}
}

// SetAnchorA The anchor value of this BitmapText.
func (self *BitmapText) SetAnchorA(member *Point) {
    self.Object.Set("anchor", member)
}

// Dirty The dirty state of this object.
func (self *BitmapText) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// SetDirtyA The dirty state of this object.
func (self *BitmapText) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// Align Alignment for multi-line text ('left', 'center' or 'right'), does not affect single lines of text.
func (self *BitmapText) Align() string{
    return self.Object.Get("align").String()
}

// SetAlignA Alignment for multi-line text ('left', 'center' or 'right'), does not affect single lines of text.
func (self *BitmapText) SetAlignA(member string) {
    self.Object.Set("align", member)
}

// Tint The tint applied to the BitmapText. This is a hex value. Set to white to disable (0xFFFFFF)
func (self *BitmapText) Tint() int{
    return self.Object.Get("tint").Int()
}

// SetTintA The tint applied to the BitmapText. This is a hex value. Set to white to disable (0xFFFFFF)
func (self *BitmapText) SetTintA(member int) {
    self.Object.Set("tint", member)
}

// Font The font the text will be rendered in, i.e. 'Arial'. Must be loaded in the browser before use.
func (self *BitmapText) Font() string{
    return self.Object.Get("font").String()
}

// SetFontA The font the text will be rendered in, i.e. 'Arial'. Must be loaded in the browser before use.
func (self *BitmapText) SetFontA(member string) {
    self.Object.Set("font", member)
}

// FontSize The size of the font in pixels.
func (self *BitmapText) FontSize() int{
    return self.Object.Get("fontSize").Int()
}

// SetFontSizeA The size of the font in pixels.
func (self *BitmapText) SetFontSizeA(member int) {
    self.Object.Set("fontSize", member)
}

// Text The text to be displayed by this BitmapText object.
func (self *BitmapText) Text() string{
    return self.Object.Get("text").String()
}

// SetTextA The text to be displayed by this BitmapText object.
func (self *BitmapText) SetTextA(member string) {
    self.Object.Set("text", member)
}

// MaxWidth The maximum display width of this BitmapText in pixels.
// 
// If BitmapText.text is longer than maxWidth then the lines will be automatically wrapped 
// based on the last whitespace character found in the line.
// 
// If no whitespace was found then no wrapping will take place and consequently the maxWidth value will not be honored.
// 
// Disable maxWidth by setting the value to 0. The maximum width of this BitmapText in pixels.
func (self *BitmapText) MaxWidth() int{
    return self.Object.Get("maxWidth").Int()
}

// SetMaxWidthA The maximum display width of this BitmapText in pixels.
// 
// If BitmapText.text is longer than maxWidth then the lines will be automatically wrapped 
// based on the last whitespace character found in the line.
// 
// If no whitespace was found then no wrapping will take place and consequently the maxWidth value will not be honored.
// 
// Disable maxWidth by setting the value to 0. The maximum width of this BitmapText in pixels.
func (self *BitmapText) SetMaxWidthA(member int) {
    self.Object.Set("maxWidth", member)
}

// Smoothed Enable or disable texture smoothing for this BitmapText.
// 
// The smoothing is applied to the BaseTexture of this font, which all letters of the text reference.
// 
// Smoothing is enabled by default.
func (self *BitmapText) Smoothed() bool{
    return self.Object.Get("smoothed").Bool()
}

// SetSmoothedA Enable or disable texture smoothing for this BitmapText.
// 
// The smoothing is applied to the BaseTexture of this font, which all letters of the text reference.
// 
// Smoothing is enabled by default.
func (self *BitmapText) SetSmoothedA(member bool) {
    self.Object.Set("smoothed", member)
}

// Children [read-only] The array of children of this container.
func (self *BitmapText) Children() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// SetChildrenA [read-only] The array of children of this container.
func (self *BitmapText) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// IgnoreChildInput If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *BitmapText) IgnoreChildInput() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// SetIgnoreChildInputA If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *BitmapText) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// Width The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *BitmapText) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *BitmapText) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *BitmapText) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *BitmapText) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Game A reference to the currently running Game.
func (self *BitmapText) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *BitmapText) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Name A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *BitmapText) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *BitmapText) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Data An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *BitmapText) Data() interface{}{
    return self.Object.Get("data")
}

// SetDataA An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *BitmapText) SetDataA(member interface{}) {
    self.Object.Set("data", member)
}

// Components The components this Game Object has installed.
func (self *BitmapText) Components() interface{}{
    return self.Object.Get("components")
}

// SetComponentsA The components this Game Object has installed.
func (self *BitmapText) SetComponentsA(member interface{}) {
    self.Object.Set("components", member)
}

// Z The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *BitmapText) Z() int{
    return self.Object.Get("z").Int()
}

// SetZA The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *BitmapText) SetZA(member int) {
    self.Object.Set("z", member)
}

// Events All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *BitmapText) Events() *Events{
    return &Events{self.Object.Get("events")}
}

// SetEventsA All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *BitmapText) SetEventsA(member *Events) {
    self.Object.Set("events", member)
}

// Animations If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *BitmapText) Animations() *AnimationManager{
    return &AnimationManager{self.Object.Get("animations")}
}

// SetAnimationsA If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *BitmapText) SetAnimationsA(member *AnimationManager) {
    self.Object.Set("animations", member)
}

// Key The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *BitmapText) Key() interface{}{
    return self.Object.Get("key")
}

// SetKeyA The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *BitmapText) SetKeyA(member interface{}) {
    self.Object.Set("key", member)
}

// World The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *BitmapText) World() *Point{
    return &Point{self.Object.Get("world")}
}

// SetWorldA The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *BitmapText) SetWorldA(member *Point) {
    self.Object.Set("world", member)
}

// Debug A debug flag designed for use with `Game.enableStep`.
func (self *BitmapText) Debug() bool{
    return self.Object.Get("debug").Bool()
}

// SetDebugA A debug flag designed for use with `Game.enableStep`.
func (self *BitmapText) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// PreviousPosition The position the Game Object was located in the previous frame.
func (self *BitmapText) PreviousPosition() *Point{
    return &Point{self.Object.Get("previousPosition")}
}

// SetPreviousPositionA The position the Game Object was located in the previous frame.
func (self *BitmapText) SetPreviousPositionA(member *Point) {
    self.Object.Set("previousPosition", member)
}

// PreviousRotation The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *BitmapText) PreviousRotation() int{
    return self.Object.Get("previousRotation").Int()
}

// SetPreviousRotationA The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *BitmapText) SetPreviousRotationA(member int) {
    self.Object.Set("previousRotation", member)
}

// RenderOrderID The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *BitmapText) RenderOrderID() int{
    return self.Object.Get("renderOrderID").Int()
}

// SetRenderOrderIDA The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *BitmapText) SetRenderOrderIDA(member int) {
    self.Object.Set("renderOrderID", member)
}

// Fresh A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *BitmapText) Fresh() bool{
    return self.Object.Get("fresh").Bool()
}

// SetFreshA A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *BitmapText) SetFreshA(member bool) {
    self.Object.Set("fresh", member)
}

// PendingDestroy A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *BitmapText) PendingDestroy() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// SetPendingDestroyA A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *BitmapText) SetPendingDestroyA(member bool) {
    self.Object.Set("pendingDestroy", member)
}

// Exists Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *BitmapText) Exists() bool{
    return self.Object.Get("exists").Bool()
}

// SetExistsA Controls if this Game Object is processed by the core game loop.
// If this Game Object has a physics body it also controls if its physics body is updated or not.
// When `exists` is set to `false` it will remove its physics body from the physics world if it has one.
// It also toggles the `visible` property to false as well.
// 
// Setting `exists` to true will add its physics body back in to the physics world, if it has one.
// It will also set the `visible` property to `true`.
func (self *BitmapText) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// Angle The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *BitmapText) Angle() int{
    return self.Object.Get("angle").Int()
}

// SetAngleA The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *BitmapText) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// AutoCull A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *BitmapText) AutoCull() bool{
    return self.Object.Get("autoCull").Bool()
}

// SetAutoCullA A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *BitmapText) SetAutoCullA(member bool) {
    self.Object.Set("autoCull", member)
}

// InCamera Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *BitmapText) InCamera() bool{
    return self.Object.Get("inCamera").Bool()
}

// SetInCameraA Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *BitmapText) SetInCameraA(member bool) {
    self.Object.Set("inCamera", member)
}

// OffsetX The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *BitmapText) OffsetX() int{
    return self.Object.Get("offsetX").Int()
}

// SetOffsetXA The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *BitmapText) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// OffsetY The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *BitmapText) OffsetY() int{
    return self.Object.Get("offsetY").Int()
}

// SetOffsetYA The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *BitmapText) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}

// CenterX The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *BitmapText) CenterX() int{
    return self.Object.Get("centerX").Int()
}

// SetCenterXA The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *BitmapText) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// CenterY The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *BitmapText) CenterY() int{
    return self.Object.Get("centerY").Int()
}

// SetCenterYA The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *BitmapText) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// Left The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *BitmapText) Left() int{
    return self.Object.Get("left").Int()
}

// SetLeftA The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *BitmapText) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// Right The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *BitmapText) Right() int{
    return self.Object.Get("right").Int()
}

// SetRightA The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *BitmapText) SetRightA(member int) {
    self.Object.Set("right", member)
}

// Top The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *BitmapText) Top() int{
    return self.Object.Get("top").Int()
}

// SetTopA The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *BitmapText) SetTopA(member int) {
    self.Object.Set("top", member)
}

// Bottom The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *BitmapText) Bottom() int{
    return self.Object.Get("bottom").Int()
}

// SetBottomA The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *BitmapText) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// DestroyPhase As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *BitmapText) DestroyPhase() bool{
    return self.Object.Get("destroyPhase").Bool()
}

// SetDestroyPhaseA As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *BitmapText) SetDestroyPhaseA(member bool) {
    self.Object.Set("destroyPhase", member)
}

// FixedToCamera A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
// 
// The values are adjusted at the rendering stage, overriding the Game Objects actual world position.
// 
// The end result is that the Game Object will appear to be 'fixed' to the camera, regardless of where in the game world
// the camera is viewing. This is useful if for example this Game Object is a UI item that you wish to be visible at all times 
// regardless where in the world the camera is.
// 
// The offsets are stored in the `cameraOffset` property.
// 
// Note that the `cameraOffset` values are in addition to any parent of this Game Object on the display list.
// 
// Be careful not to set `fixedToCamera` on Game Objects which are in Groups that already have `fixedToCamera` enabled on them.
func (self *BitmapText) FixedToCamera() bool{
    return self.Object.Get("fixedToCamera").Bool()
}

// SetFixedToCameraA A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
// 
// The values are adjusted at the rendering stage, overriding the Game Objects actual world position.
// 
// The end result is that the Game Object will appear to be 'fixed' to the camera, regardless of where in the game world
// the camera is viewing. This is useful if for example this Game Object is a UI item that you wish to be visible at all times 
// regardless where in the world the camera is.
// 
// The offsets are stored in the `cameraOffset` property.
// 
// Note that the `cameraOffset` values are in addition to any parent of this Game Object on the display list.
// 
// Be careful not to set `fixedToCamera` on Game Objects which are in Groups that already have `fixedToCamera` enabled on them.
func (self *BitmapText) SetFixedToCameraA(member bool) {
    self.Object.Set("fixedToCamera", member)
}

// CameraOffset The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *BitmapText) CameraOffset() *Point{
    return &Point{self.Object.Get("cameraOffset")}
}

// SetCameraOffsetA The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *BitmapText) SetCameraOffsetA(member *Point) {
    self.Object.Set("cameraOffset", member)
}

// Input The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *BitmapText) Input() interface{}{
    return self.Object.Get("input")
}

// SetInputA The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *BitmapText) SetInputA(member interface{}) {
    self.Object.Set("input", member)
}

// InputEnabled By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
// for this Game Object and it will then start to process click / touch events and more.
// 
// You can then access the Input Handler via `this.input`.
// 
// Note that Input related events are dispatched from `this.events`, i.e.: `events.onInputDown`.
// 
// If you set this property to false it will stop the Input Handler from processing any more input events.
// 
// If you want to _temporarily_ disable input for a Game Object, then it's better to set
// `input.enabled = false`, as it won't reset any of the Input Handlers internal properties.
// You can then toggle this back on as needed.
func (self *BitmapText) InputEnabled() bool{
    return self.Object.Get("inputEnabled").Bool()
}

// SetInputEnabledA By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
// for this Game Object and it will then start to process click / touch events and more.
// 
// You can then access the Input Handler via `this.input`.
// 
// Note that Input related events are dispatched from `this.events`, i.e.: `events.onInputDown`.
// 
// If you set this property to false it will stop the Input Handler from processing any more input events.
// 
// If you want to _temporarily_ disable input for a Game Object, then it's better to set
// `input.enabled = false`, as it won't reset any of the Input Handlers internal properties.
// You can then toggle this back on as needed.
func (self *BitmapText) SetInputEnabledA(member bool) {
    self.Object.Set("inputEnabled", member)
}

// CheckWorldBounds If this is set to `true` the Game Object checks if it is within the World bounds each frame. 
// 
// When it is no longer intersecting the world bounds it dispatches the `onOutOfBounds` event.
// 
// If it was *previously* out of bounds but is now intersecting the world bounds again it dispatches the `onEnterBounds` event.
// 
// It also optionally kills the Game Object if `outOfBoundsKill` is `true`.
// 
// When `checkWorldBounds` is enabled it forces the Game Object to calculate its full bounds every frame.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *BitmapText) CheckWorldBounds() bool{
    return self.Object.Get("checkWorldBounds").Bool()
}

// SetCheckWorldBoundsA If this is set to `true` the Game Object checks if it is within the World bounds each frame. 
// 
// When it is no longer intersecting the world bounds it dispatches the `onOutOfBounds` event.
// 
// If it was *previously* out of bounds but is now intersecting the world bounds again it dispatches the `onEnterBounds` event.
// 
// It also optionally kills the Game Object if `outOfBoundsKill` is `true`.
// 
// When `checkWorldBounds` is enabled it forces the Game Object to calculate its full bounds every frame.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *BitmapText) SetCheckWorldBoundsA(member bool) {
    self.Object.Set("checkWorldBounds", member)
}

// OutOfBoundsKill If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *BitmapText) OutOfBoundsKill() bool{
    return self.Object.Get("outOfBoundsKill").Bool()
}

// SetOutOfBoundsKillA If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *BitmapText) SetOutOfBoundsKillA(member bool) {
    self.Object.Set("outOfBoundsKill", member)
}

// OutOfCameraBoundsKill If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *BitmapText) OutOfCameraBoundsKill() bool{
    return self.Object.Get("outOfCameraBoundsKill").Bool()
}

// SetOutOfCameraBoundsKillA If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *BitmapText) SetOutOfCameraBoundsKillA(member bool) {
    self.Object.Set("outOfCameraBoundsKill", member)
}

// InWorld Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *BitmapText) InWorld() bool{
    return self.Object.Get("inWorld").Bool()
}

// SetInWorldA Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *BitmapText) SetInWorldA(member bool) {
    self.Object.Set("inWorld", member)
}

// Alive A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *BitmapText) Alive() bool{
    return self.Object.Get("alive").Bool()
}

// SetAliveA A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *BitmapText) SetAliveA(member bool) {
    self.Object.Set("alive", member)
}

// Lifespan The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *BitmapText) Lifespan() int{
    return self.Object.Get("lifespan").Int()
}

// SetLifespanA The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *BitmapText) SetLifespanA(member int) {
    self.Object.Set("lifespan", member)
}

// Body `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
// properties and methods via it.
// 
// By default Game Objects won't add themselves to any physics system and their `body` property will be `null`.
// 
// To enable this Game Object for physics you need to call `game.physics.enable(object, system)` where `object` is this object
// and `system` is the Physics system you are using. If none is given it defaults to `Phaser.Physics.Arcade`.
// 
// You can alternatively call `game.physics.arcade.enable(object)`, or add this Game Object to a physics enabled Group.
// 
// Important: Enabling a Game Object for P2 or Ninja physics will automatically set its `anchor` property to 0.5, 
// so the physics body is centered on the Game Object.
// 
// If you need a different result then adjust or re-create the Body shape offsets manually or reset the anchor after enabling physics.
func (self *BitmapText) Body() interface{}{
    return self.Object.Get("body")
}

// SetBodyA `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
// properties and methods via it.
// 
// By default Game Objects won't add themselves to any physics system and their `body` property will be `null`.
// 
// To enable this Game Object for physics you need to call `game.physics.enable(object, system)` where `object` is this object
// and `system` is the Physics system you are using. If none is given it defaults to `Phaser.Physics.Arcade`.
// 
// You can alternatively call `game.physics.arcade.enable(object)`, or add this Game Object to a physics enabled Group.
// 
// Important: Enabling a Game Object for P2 or Ninja physics will automatically set its `anchor` property to 0.5, 
// so the physics body is centered on the Game Object.
// 
// If you need a different result then adjust or re-create the Body shape offsets manually or reset the anchor after enabling physics.
func (self *BitmapText) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// X The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *BitmapText) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *BitmapText) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *BitmapText) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *BitmapText) SetYA(member int) {
    self.Object.Set("y", member)
}


// PreUpdate Automatically called by World.preUpdate.
func (self *BitmapText) PreUpdate() bool{
    return self.Object.Call("preUpdate").Bool()
}

// PreUpdateI Automatically called by World.preUpdate.
func (self *BitmapText) PreUpdateI(args ...interface{}) bool{
    return self.Object.Call("preUpdate", args).Bool()
}

// PostUpdate Automatically called by World.preUpdate.
func (self *BitmapText) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI Automatically called by World.preUpdate.
func (self *BitmapText) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// SetText The text to be displayed by this BitmapText object.
// 
// It's faster to use `BitmapText.text = string`, but this is kept for backwards compatibility.
func (self *BitmapText) SetText(text string) {
    self.Object.Call("setText", text)
}

// SetTextI The text to be displayed by this BitmapText object.
// 
// It's faster to use `BitmapText.text = string`, but this is kept for backwards compatibility.
func (self *BitmapText) SetTextI(args ...interface{}) {
    self.Object.Call("setText", args)
}

// ScanLine Given the input text this will scan the characters until either a newline is encountered, 
// or the line exceeds maxWidth, taking into account kerning, character widths and scaling.
func (self *BitmapText) ScanLine(data interface{}, scale float64, text string) interface{}{
    return self.Object.Call("scanLine", data, scale, text)
}

// ScanLineI Given the input text this will scan the characters until either a newline is encountered, 
// or the line exceeds maxWidth, taking into account kerning, character widths and scaling.
func (self *BitmapText) ScanLineI(args ...interface{}) interface{}{
    return self.Object.Call("scanLine", args)
}

// CleanText Given a text string this will scan each character in the string to ensure it exists
// in the BitmapText font data. If it doesn't the character is removed, or replaced with the `replace` argument.
// 
// If no font data has been loaded at all this returns an empty string, as nothing can be rendered.
func (self *BitmapText) CleanText(text string) string{
    return self.Object.Call("cleanText", text).String()
}

// CleanText1O Given a text string this will scan each character in the string to ensure it exists
// in the BitmapText font data. If it doesn't the character is removed, or replaced with the `replace` argument.
// 
// If no font data has been loaded at all this returns an empty string, as nothing can be rendered.
func (self *BitmapText) CleanText1O(text string, replace string) string{
    return self.Object.Call("cleanText", text, replace).String()
}

// CleanTextI Given a text string this will scan each character in the string to ensure it exists
// in the BitmapText font data. If it doesn't the character is removed, or replaced with the `replace` argument.
// 
// If no font data has been loaded at all this returns an empty string, as nothing can be rendered.
func (self *BitmapText) CleanTextI(args ...interface{}) string{
    return self.Object.Call("cleanText", args).String()
}

// UpdateText Renders text and updates it when needed.
func (self *BitmapText) UpdateText() {
    self.Object.Call("updateText")
}

// UpdateTextI Renders text and updates it when needed.
func (self *BitmapText) UpdateTextI(args ...interface{}) {
    self.Object.Call("updateText", args)
}

// PurgeGlyphs If a BitmapText changes from having a large number of characters to having very few characters it will cause lots of
// Sprites to be retained in the BitmapText._glyphs array. Although they are not attached to the display list they
// still take up memory while sat in the glyphs pool waiting to be re-used in the future.
// 
// If you know that the BitmapText will not grow any larger then you can purge out the excess glyphs from the pool 
// by calling this method.
// 
// Calling this doesn't prevent you from increasing the length of the text again in the future.
func (self *BitmapText) PurgeGlyphs() int{
    return self.Object.Call("purgeGlyphs").Int()
}

// PurgeGlyphsI If a BitmapText changes from having a large number of characters to having very few characters it will cause lots of
// Sprites to be retained in the BitmapText._glyphs array. Although they are not attached to the display list they
// still take up memory while sat in the glyphs pool waiting to be re-used in the future.
// 
// If you know that the BitmapText will not grow any larger then you can purge out the excess glyphs from the pool 
// by calling this method.
// 
// Calling this doesn't prevent you from increasing the length of the text again in the future.
func (self *BitmapText) PurgeGlyphsI(args ...interface{}) int{
    return self.Object.Call("purgeGlyphs", args).Int()
}

// UpdateTransform Updates the transform of this object.
func (self *BitmapText) UpdateTransform() {
    self.Object.Call("updateTransform")
}

// UpdateTransformI Updates the transform of this object.
func (self *BitmapText) UpdateTransformI(args ...interface{}) {
    self.Object.Call("updateTransform", args)
}

// AddChild Adds a child to the container.
func (self *BitmapText) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// AddChildI Adds a child to the container.
func (self *BitmapText) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// AddChildAt Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *BitmapText) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// AddChildAtI Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *BitmapText) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// SwapChildren Swaps the position of 2 Display Objects within this container.
func (self *BitmapText) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// SwapChildrenI Swaps the position of 2 Display Objects within this container.
func (self *BitmapText) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// GetChildIndex Returns the index position of a child DisplayObject instance
func (self *BitmapText) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// GetChildIndexI Returns the index position of a child DisplayObject instance
func (self *BitmapText) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// SetChildIndex Changes the position of an existing child in the display object container
func (self *BitmapText) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// SetChildIndexI Changes the position of an existing child in the display object container
func (self *BitmapText) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// GetChildAt Returns the child at the specified index
func (self *BitmapText) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// GetChildAtI Returns the child at the specified index
func (self *BitmapText) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// RemoveChild Removes a child from the container.
func (self *BitmapText) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// RemoveChildI Removes a child from the container.
func (self *BitmapText) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// RemoveChildAt Removes a child from the specified index position.
func (self *BitmapText) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// RemoveChildAtI Removes a child from the specified index position.
func (self *BitmapText) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// RemoveChildren Removes all children from this container that are within the begin and end indexes.
func (self *BitmapText) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// RemoveChildrenI Removes all children from this container that are within the begin and end indexes.
func (self *BitmapText) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// GetBounds Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *BitmapText) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// GetBoundsI Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *BitmapText) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// GetLocalBounds Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *BitmapText) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// GetLocalBoundsI Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *BitmapText) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// SetStageReference Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *BitmapText) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// SetStageReferenceI Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *BitmapText) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// RemoveStageReference Removes the current stage reference from the container and all of its children.
func (self *BitmapText) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// RemoveStageReferenceI Removes the current stage reference from the container and all of its children.
func (self *BitmapText) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// _renderWebGL Renders the object using the WebGL renderer
func (self *BitmapText) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// _renderWebGLI Renders the object using the WebGL renderer
func (self *BitmapText) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// _renderCanvas Renders the object using the Canvas renderer
func (self *BitmapText) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// _renderCanvasI Renders the object using the Canvas renderer
func (self *BitmapText) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

// Update Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *BitmapText) Update() {
    self.Object.Call("update")
}

// UpdateI Override this method in your own custom objects to handle any update requirements.
// It is called immediately after `preUpdate` and before `postUpdate`.
// Remember if this Game Object has any children you should call update on those too.
func (self *BitmapText) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// AlignIn Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *BitmapText) AlignIn(container interface{}) interface{}{
    return self.Object.Call("alignIn", container)
}

// AlignIn1O Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *BitmapText) AlignIn1O(container interface{}, position int) interface{}{
    return self.Object.Call("alignIn", container, position)
}

// AlignIn2O Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *BitmapText) AlignIn2O(container interface{}, position int, offsetX int) interface{}{
    return self.Object.Call("alignIn", container, position, offsetX)
}

// AlignIn3O Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *BitmapText) AlignIn3O(container interface{}, position int, offsetX int, offsetY int) interface{}{
    return self.Object.Call("alignIn", container, position, offsetX, offsetY)
}

// AlignInI Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *BitmapText) AlignInI(args ...interface{}) interface{}{
    return self.Object.Call("alignIn", args)
}

// AlignTo Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *BitmapText) AlignTo(parent interface{}) interface{}{
    return self.Object.Call("alignTo", parent)
}

// AlignTo1O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *BitmapText) AlignTo1O(parent interface{}, position int) interface{}{
    return self.Object.Call("alignTo", parent, position)
}

// AlignTo2O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *BitmapText) AlignTo2O(parent interface{}, position int, offsetX int) interface{}{
    return self.Object.Call("alignTo", parent, position, offsetX)
}

// AlignTo3O Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *BitmapText) AlignTo3O(parent interface{}, position int, offsetX int, offsetY int) interface{}{
    return self.Object.Call("alignTo", parent, position, offsetX, offsetY)
}

// AlignToI Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *BitmapText) AlignToI(args ...interface{}) interface{}{
    return self.Object.Call("alignTo", args)
}

// Destroy Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *BitmapText) Destroy() {
    self.Object.Call("destroy")
}

// Destroy1O Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *BitmapText) Destroy1O(destroyChildren bool) {
    self.Object.Call("destroy", destroyChildren)
}

// Destroy2O Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *BitmapText) Destroy2O(destroyChildren bool, destroyTexture bool) {
    self.Object.Call("destroy", destroyChildren, destroyTexture)
}

// DestroyI Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *BitmapText) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Revive Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *BitmapText) Revive() *DisplayObject{
    return &DisplayObject{self.Object.Call("revive")}
}

// Revive1O Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *BitmapText) Revive1O(health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", health)}
}

// ReviveI Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *BitmapText) ReviveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", args)}
}

// Kill Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *BitmapText) Kill() *DisplayObject{
    return &DisplayObject{self.Object.Call("kill")}
}

// KillI Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *BitmapText) KillI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("kill", args)}
}

// Reset Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *BitmapText) Reset(x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", x, y)}
}

// Reset1O Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *BitmapText) Reset1O(x int, y int, health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", x, y, health)}
}

// ResetI Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *BitmapText) ResetI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", args)}
}

