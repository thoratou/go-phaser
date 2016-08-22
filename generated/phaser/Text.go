// Automatic generation for Phaser.Text
// generated file Text.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// Create a new game object for displaying Text.
// 
// This uses a local hidden Canvas object and renders the type into it. It then makes a texture from this for rendering to the view.
// Because of this you can only display fonts that are currently loaded and available to the browser: fonts must be pre-loaded.
// 
// See {@link http://www.jordanm.co.uk/tinytype this compatibility table} for the available default fonts across mobile browsers.
type Text struct {
    *js.Object
}


// The const type of this object.
func (self *Text) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *Text) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The const physics body type of this object.
func (self *Text) GetPhysicsTypeA() int{
    return self.Object.Get("physicsType").Int()
}

// The const physics body type of this object.
func (self *Text) SetPhysicsTypeA(member int) {
    self.Object.Set("physicsType", member)
}

// Specify a padding value which is added to the line width and height when calculating the Text size.
// ALlows you to add extra spacing if Phaser is unable to accurately determine the true font dimensions.
func (self *Text) GetPaddingA() *Point{
    return &Point{self.Object.Get("padding")}
}

// Specify a padding value which is added to the line width and height when calculating the Text size.
// ALlows you to add extra spacing if Phaser is unable to accurately determine the true font dimensions.
func (self *Text) SetPaddingA(member *Point) {
    self.Object.Set("padding", member)
}

// The textBounds property allows you to specify a rectangular region upon which text alignment is based.
// See `Text.setTextBounds` for more details.
func (self *Text) GetTextBoundsA() *Rectangle{
    return &Rectangle{self.Object.Get("textBounds")}
}

// The textBounds property allows you to specify a rectangular region upon which text alignment is based.
// See `Text.setTextBounds` for more details.
func (self *Text) SetTextBoundsA(member *Rectangle) {
    self.Object.Set("textBounds", member)
}

// The canvas element that the text is rendered.
func (self *Text) GetCanvasA() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("canvas"))
}

// The canvas element that the text is rendered.
func (self *Text) SetCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("canvas", member)
}

// The context of the canvas element that the text is rendered to.
func (self *Text) GetContextA() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("context"))
}

// The context of the canvas element that the text is rendered to.
func (self *Text) SetContextA(member dom.HTMLCanvasElement) {
    self.Object.Set("context", member)
}

// An array of the color values as specified by {@link Phaser.Text#addColor addColor}.
func (self *Text) GetColorsA() []interface{}{
	array00 := self.Object.Get("colors")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of the color values as specified by {@link Phaser.Text#addColor addColor}.
func (self *Text) SetColorsA(member []interface{}) {
    self.Object.Set("colors", member)
}

// An array of the stroke color values as specified by {@link Phaser.Text#addStrokeColor addStrokeColor}.
func (self *Text) GetStrokeColorsA() []interface{}{
	array00 := self.Object.Get("strokeColors")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of the stroke color values as specified by {@link Phaser.Text#addStrokeColor addStrokeColor}.
func (self *Text) SetStrokeColorsA(member []interface{}) {
    self.Object.Set("strokeColors", member)
}

// An array of the font styles values as specified by {@link Phaser.Text#addFontStyle addFontStyle}.
func (self *Text) GetFontStylesA() []interface{}{
	array00 := self.Object.Get("fontStyles")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of the font styles values as specified by {@link Phaser.Text#addFontStyle addFontStyle}.
func (self *Text) SetFontStylesA(member []interface{}) {
    self.Object.Set("fontStyles", member)
}

// An array of the font weights values as specified by {@link Phaser.Text#addFontWeight addFontWeight}.
func (self *Text) GetFontWeightsA() []interface{}{
	array00 := self.Object.Get("fontWeights")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// An array of the font weights values as specified by {@link Phaser.Text#addFontWeight addFontWeight}.
func (self *Text) SetFontWeightsA(member []interface{}) {
    self.Object.Set("fontWeights", member)
}

// Should the linePositionX and Y values be automatically rounded before rendering the Text?
// You may wish to enable this if you want to remove the effect of sub-pixel aliasing from text.
func (self *Text) GetAutoRoundA() bool{
    return self.Object.Get("autoRound").Bool()
}

// Should the linePositionX and Y values be automatically rounded before rendering the Text?
// You may wish to enable this if you want to remove the effect of sub-pixel aliasing from text.
func (self *Text) SetAutoRoundA(member bool) {
    self.Object.Set("autoRound", member)
}

// Will this Text object use Basic or Advanced Word Wrapping?
// 
// Advanced wrapping breaks long words if they are the first of a line, and repeats the process as necessary.
// White space is condensed (e.g., consecutive spaces are replaced with one).
// Lines are trimmed of white space before processing.
// 
// It throws an error if wordWrapWidth is less than a single character.
func (self *Text) GetUseAdvancedWrapA() bool{
    return self.Object.Get("useAdvancedWrap").Bool()
}

// Will this Text object use Basic or Advanced Word Wrapping?
// 
// Advanced wrapping breaks long words if they are the first of a line, and repeats the process as necessary.
// White space is condensed (e.g., consecutive spaces are replaced with one).
// Lines are trimmed of white space before processing.
// 
// It throws an error if wordWrapWidth is less than a single character.
func (self *Text) SetUseAdvancedWrapA(member bool) {
    self.Object.Set("useAdvancedWrap", member)
}

// The text to be displayed by this Text object.
// Use a \n to insert a carriage return and split the text.
// The text will be rendered with any style currently set.
func (self *Text) GetTextA() string{
    return self.Object.Get("text").String()
}

// The text to be displayed by this Text object.
// Use a \n to insert a carriage return and split the text.
// The text will be rendered with any style currently set.
func (self *Text) SetTextA(member string) {
    self.Object.Set("text", member)
}

// Change the font used.
// 
// This is equivalent of the `font` property specified to {@link Phaser.Text#setStyle setStyle}, except
// that unlike using `setStyle` this will not change any current font fill/color settings.
// 
// The CSS font string can also be individually altered with the `font`, `fontSize`, `fontWeight`, `fontStyle`, and `fontVariant` properties.
func (self *Text) GetCssFontA() string{
    return self.Object.Get("cssFont").String()
}

// Change the font used.
// 
// This is equivalent of the `font` property specified to {@link Phaser.Text#setStyle setStyle}, except
// that unlike using `setStyle` this will not change any current font fill/color settings.
// 
// The CSS font string can also be individually altered with the `font`, `fontSize`, `fontWeight`, `fontStyle`, and `fontVariant` properties.
func (self *Text) SetCssFontA(member string) {
    self.Object.Set("cssFont", member)
}

// Change the font family that the text will be rendered in, such as 'Arial'.
// 
// Multiple CSS font families and generic fallbacks can be specified as long as
// {@link http://www.w3.org/TR/CSS2/fonts.html#propdef-font-family CSS font-family rules} are followed.
// 
// To change the entire font string use {@link Phaser.Text#cssFont cssFont} instead: eg. `text.cssFont = 'bold 20pt Arial'`.
func (self *Text) GetFontA() string{
    return self.Object.Get("font").String()
}

// Change the font family that the text will be rendered in, such as 'Arial'.
// 
// Multiple CSS font families and generic fallbacks can be specified as long as
// {@link http://www.w3.org/TR/CSS2/fonts.html#propdef-font-family CSS font-family rules} are followed.
// 
// To change the entire font string use {@link Phaser.Text#cssFont cssFont} instead: eg. `text.cssFont = 'bold 20pt Arial'`.
func (self *Text) SetFontA(member string) {
    self.Object.Set("font", member)
}

// The size of the font.
// 
// If the font size is specified in pixels (eg. `32` or `'32px`') then a number (ie. `32`) representing
// the font size in pixels is returned; otherwise the value with CSS unit is returned as a string (eg. `'12pt'`).
func (self *Text) GetFontSizeA() interface{}{
    return self.Object.Get("fontSize")
}

// The size of the font.
// 
// If the font size is specified in pixels (eg. `32` or `'32px`') then a number (ie. `32`) representing
// the font size in pixels is returned; otherwise the value with CSS unit is returned as a string (eg. `'12pt'`).
func (self *Text) SetFontSizeA(member interface{}) {
    self.Object.Set("fontSize", member)
}

// The weight of the font: 'normal', 'bold', or {@link http://www.w3.org/TR/CSS2/fonts.html#propdef-font-weight a valid CSS font weight}.
func (self *Text) GetFontWeightA() string{
    return self.Object.Get("fontWeight").String()
}

// The weight of the font: 'normal', 'bold', or {@link http://www.w3.org/TR/CSS2/fonts.html#propdef-font-weight a valid CSS font weight}.
func (self *Text) SetFontWeightA(member string) {
    self.Object.Set("fontWeight", member)
}

// The style of the font: 'normal', 'italic', 'oblique'
func (self *Text) GetFontStyleA() string{
    return self.Object.Get("fontStyle").String()
}

// The style of the font: 'normal', 'italic', 'oblique'
func (self *Text) SetFontStyleA(member string) {
    self.Object.Set("fontStyle", member)
}

// The variant the font: 'normal', 'small-caps'
func (self *Text) GetFontVariantA() string{
    return self.Object.Get("fontVariant").String()
}

// The variant the font: 'normal', 'small-caps'
func (self *Text) SetFontVariantA(member string) {
    self.Object.Set("fontVariant", member)
}

// A canvas fillstyle that will be used on the text eg 'red', '#00FF00'.
func (self *Text) GetFillA() interface{}{
    return self.Object.Get("fill")
}

// A canvas fillstyle that will be used on the text eg 'red', '#00FF00'.
func (self *Text) SetFillA(member interface{}) {
    self.Object.Set("fill", member)
}

// Controls the horizontal alignment for multiline text.
// Can be: 'left', 'center' or 'right'.
// Does not affect single lines of text. For that please see `setTextBounds`.
func (self *Text) GetAlignA() string{
    return self.Object.Get("align").String()
}

// Controls the horizontal alignment for multiline text.
// Can be: 'left', 'center' or 'right'.
// Does not affect single lines of text. For that please see `setTextBounds`.
func (self *Text) SetAlignA(member string) {
    self.Object.Set("align", member)
}

// The resolution of the canvas the text is rendered to.
// This defaults to match the resolution of the renderer, but can be changed on a per Text object basis.
func (self *Text) GetResolutionA() int{
    return self.Object.Get("resolution").Int()
}

// The resolution of the canvas the text is rendered to.
// This defaults to match the resolution of the renderer, but can be changed on a per Text object basis.
func (self *Text) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// The size (in pixels) of the tabs, for when text includes tab characters. 0 disables. 
// Can be an integer or an array of varying tab sizes, one tab per element.
// For example if you set tabs to 100 then when Text encounters a tab it will jump ahead 100 pixels.
// If you set tabs to be `[100,200]` then it will set the first tab at 100px and the second at 200px.
func (self *Text) GetTabsA() interface{}{
    return self.Object.Get("tabs")
}

// The size (in pixels) of the tabs, for when text includes tab characters. 0 disables. 
// Can be an integer or an array of varying tab sizes, one tab per element.
// For example if you set tabs to 100 then when Text encounters a tab it will jump ahead 100 pixels.
// If you set tabs to be `[100,200]` then it will set the first tab at 100px and the second at 200px.
func (self *Text) SetTabsA(member interface{}) {
    self.Object.Set("tabs", member)
}

// Horizontal alignment of the text within the `textBounds`. Can be: 'left', 'center' or 'right'.
func (self *Text) GetBoundsAlignHA() string{
    return self.Object.Get("boundsAlignH").String()
}

// Horizontal alignment of the text within the `textBounds`. Can be: 'left', 'center' or 'right'.
func (self *Text) SetBoundsAlignHA(member string) {
    self.Object.Set("boundsAlignH", member)
}

// Vertical alignment of the text within the `textBounds`. Can be: 'top', 'middle' or 'bottom'.
func (self *Text) GetBoundsAlignVA() string{
    return self.Object.Get("boundsAlignV").String()
}

// Vertical alignment of the text within the `textBounds`. Can be: 'top', 'middle' or 'bottom'.
func (self *Text) SetBoundsAlignVA(member string) {
    self.Object.Set("boundsAlignV", member)
}

// A canvas fillstyle that will be used on the text stroke eg 'blue', '#FCFF00'.
func (self *Text) GetStrokeA() string{
    return self.Object.Get("stroke").String()
}

// A canvas fillstyle that will be used on the text stroke eg 'blue', '#FCFF00'.
func (self *Text) SetStrokeA(member string) {
    self.Object.Set("stroke", member)
}

// A number that represents the thickness of the stroke. Default is 0 (no stroke)
func (self *Text) GetStrokeThicknessA() int{
    return self.Object.Get("strokeThickness").Int()
}

// A number that represents the thickness of the stroke. Default is 0 (no stroke)
func (self *Text) SetStrokeThicknessA(member int) {
    self.Object.Set("strokeThickness", member)
}

// Indicates if word wrap should be used.
func (self *Text) GetWordWrapA() bool{
    return self.Object.Get("wordWrap").Bool()
}

// Indicates if word wrap should be used.
func (self *Text) SetWordWrapA(member bool) {
    self.Object.Set("wordWrap", member)
}

// The width at which text will wrap.
func (self *Text) GetWordWrapWidthA() int{
    return self.Object.Get("wordWrapWidth").Int()
}

// The width at which text will wrap.
func (self *Text) SetWordWrapWidthA(member int) {
    self.Object.Set("wordWrapWidth", member)
}

// Additional spacing (in pixels) between each line of text if multi-line.
func (self *Text) GetLineSpacingA() int{
    return self.Object.Get("lineSpacing").Int()
}

// Additional spacing (in pixels) between each line of text if multi-line.
func (self *Text) SetLineSpacingA(member int) {
    self.Object.Set("lineSpacing", member)
}

// The shadowOffsetX value in pixels. This is how far offset horizontally the shadow effect will be.
func (self *Text) GetShadowOffsetXA() int{
    return self.Object.Get("shadowOffsetX").Int()
}

// The shadowOffsetX value in pixels. This is how far offset horizontally the shadow effect will be.
func (self *Text) SetShadowOffsetXA(member int) {
    self.Object.Set("shadowOffsetX", member)
}

// The shadowOffsetY value in pixels. This is how far offset vertically the shadow effect will be.
func (self *Text) GetShadowOffsetYA() int{
    return self.Object.Get("shadowOffsetY").Int()
}

// The shadowOffsetY value in pixels. This is how far offset vertically the shadow effect will be.
func (self *Text) SetShadowOffsetYA(member int) {
    self.Object.Set("shadowOffsetY", member)
}

// The color of the shadow, as given in CSS rgba format. Set the alpha component to 0 to disable the shadow.
func (self *Text) GetShadowColorA() string{
    return self.Object.Get("shadowColor").String()
}

// The color of the shadow, as given in CSS rgba format. Set the alpha component to 0 to disable the shadow.
func (self *Text) SetShadowColorA(member string) {
    self.Object.Set("shadowColor", member)
}

// The shadowBlur value. Make the shadow softer by applying a Gaussian blur to it. A number from 0 (no blur) up to approx. 10 (depending on scene).
func (self *Text) GetShadowBlurA() int{
    return self.Object.Get("shadowBlur").Int()
}

// The shadowBlur value. Make the shadow softer by applying a Gaussian blur to it. A number from 0 (no blur) up to approx. 10 (depending on scene).
func (self *Text) SetShadowBlurA(member int) {
    self.Object.Set("shadowBlur", member)
}

// Sets if the drop shadow is applied to the Text stroke.
func (self *Text) GetShadowStrokeA() bool{
    return self.Object.Get("shadowStroke").Bool()
}

// Sets if the drop shadow is applied to the Text stroke.
func (self *Text) SetShadowStrokeA(member bool) {
    self.Object.Set("shadowStroke", member)
}

// Sets if the drop shadow is applied to the Text fill.
func (self *Text) GetShadowFillA() bool{
    return self.Object.Get("shadowFill").Bool()
}

// Sets if the drop shadow is applied to the Text fill.
func (self *Text) SetShadowFillA(member bool) {
    self.Object.Set("shadowFill", member)
}

// The width of the Text. Setting this will modify the scale to achieve the value requested.
func (self *Text) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the Text. Setting this will modify the scale to achieve the value requested.
func (self *Text) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the Text. Setting this will modify the scale to achieve the value requested.
func (self *Text) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the Text. Setting this will modify the scale to achieve the value requested.
func (self *Text) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *Text) GetAnchorA() *Point{
    return &Point{self.Object.Get("anchor")}
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *Text) SetAnchorA(member *Point) {
    self.Object.Set("anchor", member)
}

// The texture that the sprite is using
func (self *Text) GetTextureA() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// The texture that the sprite is using
func (self *Text) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *Text) GetTintA() int{
    return self.Object.Get("tint").Int()
}

// The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *Text) SetTintA(member int) {
    self.Object.Set("tint", member)
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *Text) GetTintedTextureA() *Canvas{
    return &Canvas{self.Object.Get("tintedTexture")}
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *Text) SetTintedTextureA(member *Canvas) {
    self.Object.Set("tintedTexture", member)
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *Text) GetBlendModeA() int{
    return self.Object.Get("blendMode").Int()
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *Text) SetBlendModeA(member int) {
    self.Object.Set("blendMode", member)
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *Text) GetShaderA() *AbstractFilter{
    return &AbstractFilter{self.Object.Get("shader")}
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *Text) SetShaderA(member *AbstractFilter) {
    self.Object.Set("shader", member)
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *Text) GetExistsA() bool{
    return self.Object.Get("exists").Bool()
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *Text) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// [read-only] The array of children of this container.
func (self *Text) GetChildrenA() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// [read-only] The array of children of this container.
func (self *Text) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Text) GetIgnoreChildInputA() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Text) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// A reference to the currently running Game.
func (self *Text) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *Text) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Text) GetNameA() string{
    return self.Object.Get("name").String()
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Text) SetNameA(member string) {
    self.Object.Set("name", member)
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Text) GetDataA() interface{}{
    return self.Object.Get("data")
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Text) SetDataA(member interface{}) {
    self.Object.Set("data", member)
}

// The components this Game Object has installed.
func (self *Text) GetComponentsA() interface{}{
    return self.Object.Get("components")
}

// The components this Game Object has installed.
func (self *Text) SetComponentsA(member interface{}) {
    self.Object.Set("components", member)
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Text) GetZA() int{
    return self.Object.Get("z").Int()
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Text) SetZA(member int) {
    self.Object.Set("z", member)
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Text) GetEventsA() *Events{
    return &Events{self.Object.Get("events")}
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Text) SetEventsA(member *Events) {
    self.Object.Set("events", member)
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Text) GetAnimationsA() *AnimationManager{
    return &AnimationManager{self.Object.Get("animations")}
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Text) SetAnimationsA(member *AnimationManager) {
    self.Object.Set("animations", member)
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Text) GetKeyA() interface{}{
    return self.Object.Get("key")
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Text) SetKeyA(member interface{}) {
    self.Object.Set("key", member)
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Text) GetWorldA() *Point{
    return &Point{self.Object.Get("world")}
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Text) SetWorldA(member *Point) {
    self.Object.Set("world", member)
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Text) GetDebugA() bool{
    return self.Object.Get("debug").Bool()
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Text) SetDebugA(member bool) {
    self.Object.Set("debug", member)
}

// The position the Game Object was located in the previous frame.
func (self *Text) GetPreviousPositionA() *Point{
    return &Point{self.Object.Get("previousPosition")}
}

// The position the Game Object was located in the previous frame.
func (self *Text) SetPreviousPositionA(member *Point) {
    self.Object.Set("previousPosition", member)
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Text) GetPreviousRotationA() int{
    return self.Object.Get("previousRotation").Int()
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Text) SetPreviousRotationA(member int) {
    self.Object.Set("previousRotation", member)
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Text) GetRenderOrderIDA() int{
    return self.Object.Get("renderOrderID").Int()
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Text) SetRenderOrderIDA(member int) {
    self.Object.Set("renderOrderID", member)
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Text) GetFreshA() bool{
    return self.Object.Get("fresh").Bool()
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Text) SetFreshA(member bool) {
    self.Object.Set("fresh", member)
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Text) GetPendingDestroyA() bool{
    return self.Object.Get("pendingDestroy").Bool()
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Text) SetPendingDestroyA(member bool) {
    self.Object.Set("pendingDestroy", member)
}

// The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *Text) GetAngleA() int{
    return self.Object.Get("angle").Int()
}

// The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *Text) SetAngleA(member int) {
    self.Object.Set("angle", member)
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Text) GetAutoCullA() bool{
    return self.Object.Get("autoCull").Bool()
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Text) SetAutoCullA(member bool) {
    self.Object.Set("autoCull", member)
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Text) GetInCameraA() bool{
    return self.Object.Get("inCamera").Bool()
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Text) SetInCameraA(member bool) {
    self.Object.Set("inCamera", member)
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Text) GetOffsetXA() int{
    return self.Object.Get("offsetX").Int()
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Text) SetOffsetXA(member int) {
    self.Object.Set("offsetX", member)
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Text) GetOffsetYA() int{
    return self.Object.Get("offsetY").Int()
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Text) SetOffsetYA(member int) {
    self.Object.Set("offsetY", member)
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Text) GetCenterXA() int{
    return self.Object.Get("centerX").Int()
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Text) SetCenterXA(member int) {
    self.Object.Set("centerX", member)
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Text) GetCenterYA() int{
    return self.Object.Get("centerY").Int()
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Text) SetCenterYA(member int) {
    self.Object.Set("centerY", member)
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Text) GetLeftA() int{
    return self.Object.Get("left").Int()
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Text) SetLeftA(member int) {
    self.Object.Set("left", member)
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Text) GetRightA() int{
    return self.Object.Get("right").Int()
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Text) SetRightA(member int) {
    self.Object.Set("right", member)
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Text) GetTopA() int{
    return self.Object.Get("top").Int()
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Text) SetTopA(member int) {
    self.Object.Set("top", member)
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Text) GetBottomA() int{
    return self.Object.Get("bottom").Int()
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Text) SetBottomA(member int) {
    self.Object.Set("bottom", member)
}

// The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Text) GetCropRectA() *Rectangle{
    return &Rectangle{self.Object.Get("cropRect")}
}

// The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Text) SetCropRectA(member *Rectangle) {
    self.Object.Set("cropRect", member)
}

// Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *Text) GetDeltaXA() int{
    return self.Object.Get("deltaX").Int()
}

// Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *Text) SetDeltaXA(member int) {
    self.Object.Set("deltaX", member)
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *Text) GetDeltaYA() int{
    return self.Object.Get("deltaY").Int()
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *Text) SetDeltaYA(member int) {
    self.Object.Set("deltaY", member)
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *Text) GetDeltaZA() int{
    return self.Object.Get("deltaZ").Int()
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *Text) SetDeltaZA(member int) {
    self.Object.Set("deltaZ", member)
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Text) GetDestroyPhaseA() bool{
    return self.Object.Get("destroyPhase").Bool()
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Text) SetDestroyPhaseA(member bool) {
    self.Object.Set("destroyPhase", member)
}

// A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
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
func (self *Text) GetFixedToCameraA() bool{
    return self.Object.Get("fixedToCamera").Bool()
}

// A Game Object that is "fixed" to the camera uses its x/y coordinates as offsets from the top left of the camera during rendering.
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
func (self *Text) SetFixedToCameraA(member bool) {
    self.Object.Set("fixedToCamera", member)
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Text) GetCameraOffsetA() *Point{
    return &Point{self.Object.Get("cameraOffset")}
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Text) SetCameraOffsetA(member *Point) {
    self.Object.Set("cameraOffset", member)
}

// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *Text) GetHealthA() int{
    return self.Object.Get("health").Int()
}

// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *Text) SetHealthA(member int) {
    self.Object.Set("health", member)
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *Text) GetMaxHealthA() int{
    return self.Object.Get("maxHealth").Int()
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *Text) SetMaxHealthA(member int) {
    self.Object.Set("maxHealth", member)
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *Text) GetDamageA() interface{}{
    return self.Object.Get("damage")
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *Text) SetDamageA(member interface{}) {
    self.Object.Set("damage", member)
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *Text) GetSetHealthA() interface{}{
    return self.Object.Get("setHealth")
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *Text) SetSetHealthA(member interface{}) {
    self.Object.Set("setHealth", member)
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *Text) GetHealA() interface{}{
    return self.Object.Get("heal")
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *Text) SetHealA(member interface{}) {
    self.Object.Set("heal", member)
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Text) GetInputA() interface{}{
    return self.Object.Get("input")
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Text) SetInputA(member interface{}) {
    self.Object.Set("input", member)
}

// By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
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
func (self *Text) GetInputEnabledA() bool{
    return self.Object.Get("inputEnabled").Bool()
}

// By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
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
func (self *Text) SetInputEnabledA(member bool) {
    self.Object.Set("inputEnabled", member)
}

// If this is set to `true` the Game Object checks if it is within the World bounds each frame. 
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
func (self *Text) GetCheckWorldBoundsA() bool{
    return self.Object.Get("checkWorldBounds").Bool()
}

// If this is set to `true` the Game Object checks if it is within the World bounds each frame. 
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
func (self *Text) SetCheckWorldBoundsA(member bool) {
    self.Object.Set("checkWorldBounds", member)
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *Text) GetOutOfBoundsKillA() bool{
    return self.Object.Get("outOfBoundsKill").Bool()
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *Text) SetOutOfBoundsKillA(member bool) {
    self.Object.Set("outOfBoundsKill", member)
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *Text) GetOutOfCameraBoundsKillA() bool{
    return self.Object.Get("outOfCameraBoundsKill").Bool()
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *Text) SetOutOfCameraBoundsKillA(member bool) {
    self.Object.Set("outOfCameraBoundsKill", member)
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *Text) GetInWorldA() bool{
    return self.Object.Get("inWorld").Bool()
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *Text) SetInWorldA(member bool) {
    self.Object.Set("inWorld", member)
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Text) GetAliveA() bool{
    return self.Object.Get("alive").Bool()
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Text) SetAliveA(member bool) {
    self.Object.Set("alive", member)
}

// The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *Text) GetLifespanA() int{
    return self.Object.Get("lifespan").Int()
}

// The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *Text) SetLifespanA(member int) {
    self.Object.Set("lifespan", member)
}

// Gets or sets the current frame index of the texture being used to render this Game Object.
// 
// To change the frame set `frame` to the index of the new frame in the sprite sheet you wish this Game Object to use,
// for example: `player.frame = 4`.
// 
// If the frame index given doesn't exist it will revert to the first frame found in the texture.
// 
// If you are using a texture atlas then you should use the `frameName` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *Text) GetFrameA() int{
    return self.Object.Get("frame").Int()
}

// Gets or sets the current frame index of the texture being used to render this Game Object.
// 
// To change the frame set `frame` to the index of the new frame in the sprite sheet you wish this Game Object to use,
// for example: `player.frame = 4`.
// 
// If the frame index given doesn't exist it will revert to the first frame found in the texture.
// 
// If you are using a texture atlas then you should use the `frameName` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *Text) SetFrameA(member int) {
    self.Object.Set("frame", member)
}

// Gets or sets the current frame name of the texture being used to render this Game Object.
// 
// To change the frame set `frameName` to the name of the new frame in the texture atlas you wish this Game Object to use, 
// for example: `player.frameName = "idle"`.
// 
// If the frame name given doesn't exist it will revert to the first frame found in the texture and throw a console warning.
// 
// If you are using a sprite sheet then you should use the `frame` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *Text) GetFrameNameA() string{
    return self.Object.Get("frameName").String()
}

// Gets or sets the current frame name of the texture being used to render this Game Object.
// 
// To change the frame set `frameName` to the name of the new frame in the texture atlas you wish this Game Object to use, 
// for example: `player.frameName = "idle"`.
// 
// If the frame name given doesn't exist it will revert to the first frame found in the texture and throw a console warning.
// 
// If you are using a sprite sheet then you should use the `frame` property instead.
// 
// If you wish to fully replace the texture being used see `loadTexture`.
func (self *Text) SetFrameNameA(member string) {
    self.Object.Set("frameName", member)
}

// `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
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
func (self *Text) GetBodyA() interface{}{
    return self.Object.Get("body")
}

// `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
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
func (self *Text) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *Text) GetXA() int{
    return self.Object.Get("x").Int()
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *Text) SetXA(member int) {
    self.Object.Set("x", member)
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *Text) GetYA() int{
    return self.Object.Get("y").Int()
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *Text) SetYA(member int) {
    self.Object.Set("y", member)
}

// The callback that will apply any scale limiting to the worldTransform.
func (self *Text) SetTransformCallbackA(member func(...interface{})) {
    self.Object.Set("transformCallback", member)
}

// The context under which `transformCallback` is called.
func (self *Text) GetTransformCallbackContextA() interface{}{
    return self.Object.Get("transformCallbackContext")
}

// The context under which `transformCallback` is called.
func (self *Text) SetTransformCallbackContextA(member interface{}) {
    self.Object.Set("transformCallbackContext", member)
}

// The minimum scale this Game Object will scale down to.
// 
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
// 
// Set it to `null` to remove the limit.
func (self *Text) GetScaleMinA() *Point{
    return &Point{self.Object.Get("scaleMin")}
}

// The minimum scale this Game Object will scale down to.
// 
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
// 
// Set it to `null` to remove the limit.
func (self *Text) SetScaleMinA(member *Point) {
    self.Object.Set("scaleMin", member)
}

// The maximum scale this Game Object will scale up to. 
// 
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
// 
// Set it to `null` to remove the limit.
func (self *Text) GetScaleMaxA() *Point{
    return &Point{self.Object.Get("scaleMax")}
}

// The maximum scale this Game Object will scale up to. 
// 
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
// 
// Set it to `null` to remove the limit.
func (self *Text) SetScaleMaxA(member *Point) {
    self.Object.Set("scaleMax", member)
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *Text) GetSmoothedA() bool{
    return self.Object.Get("smoothed").Bool()
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *Text) SetSmoothedA(member bool) {
    self.Object.Set("smoothed", member)
}



// Automatically called by World.preUpdate.
func (self *Text) PreUpdate() {
    self.Object.Call("preUpdate")
}

// Automatically called by World.preUpdate.
func (self *Text) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Override this function to handle any special update requirements.
func (self *Text) Update() {
    self.Object.Call("update")
}

// Override this function to handle any special update requirements.
func (self *Text) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Destroy this Text object, removing it from the group it belongs to.
func (self *Text) Destroy() {
    self.Object.Call("destroy")
}

// Destroy this Text object, removing it from the group it belongs to.
func (self *Text) Destroy1O(destroyChildren bool) {
    self.Object.Call("destroy", destroyChildren)
}

// Destroy this Text object, removing it from the group it belongs to.
func (self *Text) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Sets a drop shadow effect on the Text. You can specify the horizontal and vertical distance of the drop shadow with the `x` and `y` parameters.
// The color controls the shade of the shadow (default is black) and can be either an `rgba` or `hex` value.
// The blur is the strength of the shadow. A value of zero means a hard shadow, a value of 10 means a very soft shadow.
// To remove a shadow already in place you can call this method with no parameters set.
func (self *Text) SetShadow() *Text{
    return &Text{self.Object.Call("setShadow")}
}

// Sets a drop shadow effect on the Text. You can specify the horizontal and vertical distance of the drop shadow with the `x` and `y` parameters.
// The color controls the shade of the shadow (default is black) and can be either an `rgba` or `hex` value.
// The blur is the strength of the shadow. A value of zero means a hard shadow, a value of 10 means a very soft shadow.
// To remove a shadow already in place you can call this method with no parameters set.
func (self *Text) SetShadow1O(x int) *Text{
    return &Text{self.Object.Call("setShadow", x)}
}

// Sets a drop shadow effect on the Text. You can specify the horizontal and vertical distance of the drop shadow with the `x` and `y` parameters.
// The color controls the shade of the shadow (default is black) and can be either an `rgba` or `hex` value.
// The blur is the strength of the shadow. A value of zero means a hard shadow, a value of 10 means a very soft shadow.
// To remove a shadow already in place you can call this method with no parameters set.
func (self *Text) SetShadow2O(x int, y int) *Text{
    return &Text{self.Object.Call("setShadow", x, y)}
}

// Sets a drop shadow effect on the Text. You can specify the horizontal and vertical distance of the drop shadow with the `x` and `y` parameters.
// The color controls the shade of the shadow (default is black) and can be either an `rgba` or `hex` value.
// The blur is the strength of the shadow. A value of zero means a hard shadow, a value of 10 means a very soft shadow.
// To remove a shadow already in place you can call this method with no parameters set.
func (self *Text) SetShadow3O(x int, y int, color string) *Text{
    return &Text{self.Object.Call("setShadow", x, y, color)}
}

// Sets a drop shadow effect on the Text. You can specify the horizontal and vertical distance of the drop shadow with the `x` and `y` parameters.
// The color controls the shade of the shadow (default is black) and can be either an `rgba` or `hex` value.
// The blur is the strength of the shadow. A value of zero means a hard shadow, a value of 10 means a very soft shadow.
// To remove a shadow already in place you can call this method with no parameters set.
func (self *Text) SetShadow4O(x int, y int, color string, blur int) *Text{
    return &Text{self.Object.Call("setShadow", x, y, color, blur)}
}

// Sets a drop shadow effect on the Text. You can specify the horizontal and vertical distance of the drop shadow with the `x` and `y` parameters.
// The color controls the shade of the shadow (default is black) and can be either an `rgba` or `hex` value.
// The blur is the strength of the shadow. A value of zero means a hard shadow, a value of 10 means a very soft shadow.
// To remove a shadow already in place you can call this method with no parameters set.
func (self *Text) SetShadow5O(x int, y int, color string, blur int, shadowStroke bool) *Text{
    return &Text{self.Object.Call("setShadow", x, y, color, blur, shadowStroke)}
}

// Sets a drop shadow effect on the Text. You can specify the horizontal and vertical distance of the drop shadow with the `x` and `y` parameters.
// The color controls the shade of the shadow (default is black) and can be either an `rgba` or `hex` value.
// The blur is the strength of the shadow. A value of zero means a hard shadow, a value of 10 means a very soft shadow.
// To remove a shadow already in place you can call this method with no parameters set.
func (self *Text) SetShadow6O(x int, y int, color string, blur int, shadowStroke bool, shadowFill bool) *Text{
    return &Text{self.Object.Call("setShadow", x, y, color, blur, shadowStroke, shadowFill)}
}

// Sets a drop shadow effect on the Text. You can specify the horizontal and vertical distance of the drop shadow with the `x` and `y` parameters.
// The color controls the shade of the shadow (default is black) and can be either an `rgba` or `hex` value.
// The blur is the strength of the shadow. A value of zero means a hard shadow, a value of 10 means a very soft shadow.
// To remove a shadow already in place you can call this method with no parameters set.
func (self *Text) SetShadowI(args ...interface{}) *Text{
    return &Text{self.Object.Call("setShadow", args)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle() *Text{
    return &Text{self.Object.Call("setStyle")}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle1O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle2O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle3O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle4O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle5O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle6O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle7O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle8O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle9O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle10O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle11O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle12O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle13O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle14O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle15O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle16O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle17O(style interface{}) *Text{
    return &Text{self.Object.Call("setStyle", style)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyle18O(style interface{}, update bool) *Text{
    return &Text{self.Object.Call("setStyle", style, update)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyleI(args ...interface{}) *Text{
    return &Text{self.Object.Call("setStyle", args)}
}

// Renders text. This replaces the Pixi.Text.updateText function as we need a few extra bits in here.
func (self *Text) UpdateText() {
    self.Object.Call("updateText")
}

// Renders text. This replaces the Pixi.Text.updateText function as we need a few extra bits in here.
func (self *Text) UpdateTextI(args ...interface{}) {
    self.Object.Call("updateText", args)
}

// Renders a line of text that contains tab characters if Text.tab > 0.
// Called automatically by updateText.
func (self *Text) RenderTabLine(line string, x int, y int, fill bool) {
    self.Object.Call("renderTabLine", line, x, y, fill)
}

// Renders a line of text that contains tab characters if Text.tab > 0.
// Called automatically by updateText.
func (self *Text) RenderTabLineI(args ...interface{}) {
    self.Object.Call("renderTabLine", args)
}

// Sets the Shadow on the Text.context based on the Style settings, or disables it if not enabled.
// This is called automatically by Text.updateText.
func (self *Text) UpdateShadow(state bool) {
    self.Object.Call("updateShadow", state)
}

// Sets the Shadow on the Text.context based on the Style settings, or disables it if not enabled.
// This is called automatically by Text.updateText.
func (self *Text) UpdateShadowI(args ...interface{}) {
    self.Object.Call("updateShadow", args)
}

// Measures a line of text character by character taking into the account the specified character styles.
func (self *Text) MeasureLine(line string) int{
    return self.Object.Call("measureLine", line).Int()
}

// Measures a line of text character by character taking into the account the specified character styles.
func (self *Text) MeasureLineI(args ...interface{}) int{
    return self.Object.Call("measureLine", args).Int()
}

// Updates a line of text, applying fill and stroke per-character colors or style and weight per-character font if applicable.
func (self *Text) UpdateLine() {
    self.Object.Call("updateLine")
}

// Updates a line of text, applying fill and stroke per-character colors or style and weight per-character font if applicable.
func (self *Text) UpdateLineI(args ...interface{}) {
    self.Object.Call("updateLine", args)
}

// Clears any text fill or stroke colors that were set by `addColor` or `addStrokeColor`.
func (self *Text) ClearColors() *Text{
    return &Text{self.Object.Call("clearColors")}
}

// Clears any text fill or stroke colors that were set by `addColor` or `addStrokeColor`.
func (self *Text) ClearColorsI(args ...interface{}) *Text{
    return &Text{self.Object.Call("clearColors", args)}
}

// Clears any text styles or weights font that were set by `addFontStyle` or `addFontWeight`.
func (self *Text) ClearFontValues() *Text{
    return &Text{self.Object.Call("clearFontValues")}
}

// Clears any text styles or weights font that were set by `addFontStyle` or `addFontWeight`.
func (self *Text) ClearFontValuesI(args ...interface{}) *Text{
    return &Text{self.Object.Call("clearFontValues", args)}
}

// Set specific colors for certain characters within the Text.
// 
// It works by taking a color value, which is a typical HTML string such as `#ff0000` or `rgb(255,0,0)` and a position.
// The position value is the index of the character in the Text string to start applying this color to.
// Once set the color remains in use until either another color or the end of the string is encountered.
// For example if the Text was `Photon Storm` and you did `Text.addColor('#ffff00', 6)` it would color in the word `Storm` in yellow.
// 
// If you wish to change the stroke color see addStrokeColor instead.
func (self *Text) AddColor(color string, position int) *Text{
    return &Text{self.Object.Call("addColor", color, position)}
}

// Set specific colors for certain characters within the Text.
// 
// It works by taking a color value, which is a typical HTML string such as `#ff0000` or `rgb(255,0,0)` and a position.
// The position value is the index of the character in the Text string to start applying this color to.
// Once set the color remains in use until either another color or the end of the string is encountered.
// For example if the Text was `Photon Storm` and you did `Text.addColor('#ffff00', 6)` it would color in the word `Storm` in yellow.
// 
// If you wish to change the stroke color see addStrokeColor instead.
func (self *Text) AddColorI(args ...interface{}) *Text{
    return &Text{self.Object.Call("addColor", args)}
}

// Set specific stroke colors for certain characters within the Text.
// 
// It works by taking a color value, which is a typical HTML string such as `#ff0000` or `rgb(255,0,0)` and a position.
// The position value is the index of the character in the Text string to start applying this color to.
// Once set the color remains in use until either another color or the end of the string is encountered.
// For example if the Text was `Photon Storm` and you did `Text.addColor('#ffff00', 6)` it would color in the word `Storm` in yellow.
// 
// This has no effect if stroke is disabled or has a thickness of 0.
// 
// If you wish to change the text fill color see addColor instead.
func (self *Text) AddStrokeColor(color string, position int) *Text{
    return &Text{self.Object.Call("addStrokeColor", color, position)}
}

// Set specific stroke colors for certain characters within the Text.
// 
// It works by taking a color value, which is a typical HTML string such as `#ff0000` or `rgb(255,0,0)` and a position.
// The position value is the index of the character in the Text string to start applying this color to.
// Once set the color remains in use until either another color or the end of the string is encountered.
// For example if the Text was `Photon Storm` and you did `Text.addColor('#ffff00', 6)` it would color in the word `Storm` in yellow.
// 
// This has no effect if stroke is disabled or has a thickness of 0.
// 
// If you wish to change the text fill color see addColor instead.
func (self *Text) AddStrokeColorI(args ...interface{}) *Text{
    return &Text{self.Object.Call("addStrokeColor", args)}
}

// Set specific font styles for certain characters within the Text.
// 
// It works by taking a font style value, which is a typical string such as `normal`, `italic` or `oblique`.
// The position value is the index of the character in the Text string to start applying this font style to.
// Once set the font style remains in use until either another font style or the end of the string is encountered.
// For example if the Text was `Photon Storm` and you did `Text.addFontStyle('italic', 6)` it would font style in the word `Storm` in italic.
// 
// If you wish to change the text font weight see addFontWeight instead.
func (self *Text) AddFontStyle(style string, position int) *Text{
    return &Text{self.Object.Call("addFontStyle", style, position)}
}

// Set specific font styles for certain characters within the Text.
// 
// It works by taking a font style value, which is a typical string such as `normal`, `italic` or `oblique`.
// The position value is the index of the character in the Text string to start applying this font style to.
// Once set the font style remains in use until either another font style or the end of the string is encountered.
// For example if the Text was `Photon Storm` and you did `Text.addFontStyle('italic', 6)` it would font style in the word `Storm` in italic.
// 
// If you wish to change the text font weight see addFontWeight instead.
func (self *Text) AddFontStyleI(args ...interface{}) *Text{
    return &Text{self.Object.Call("addFontStyle", args)}
}

// Set specific font weights for certain characters within the Text.
// 
// It works by taking a font weight value, which is a typical string such as `normal`, `bold`, `bolder`, etc.
// The position value is the index of the character in the Text string to start applying this font weight to.
// Once set the font weight remains in use until either another font weight or the end of the string is encountered.
// For example if the Text was `Photon Storm` and you did `Text.addFontWeight('bold', 6)` it would font weight in the word `Storm` in bold.
// 
// If you wish to change the text font style see addFontStyle instead.
func (self *Text) AddFontWeight(style string, position int) *Text{
    return &Text{self.Object.Call("addFontWeight", style, position)}
}

// Set specific font weights for certain characters within the Text.
// 
// It works by taking a font weight value, which is a typical string such as `normal`, `bold`, `bolder`, etc.
// The position value is the index of the character in the Text string to start applying this font weight to.
// Once set the font weight remains in use until either another font weight or the end of the string is encountered.
// For example if the Text was `Photon Storm` and you did `Text.addFontWeight('bold', 6)` it would font weight in the word `Storm` in bold.
// 
// If you wish to change the text font style see addFontStyle instead.
func (self *Text) AddFontWeightI(args ...interface{}) *Text{
    return &Text{self.Object.Call("addFontWeight", args)}
}

// Runs the given text through the Text.runWordWrap function and returns
// the results as an array, where each element of the array corresponds to a wrapped
// line of text.
// 
// Useful if you wish to control pagination on long pieces of content.
func (self *Text) PrecalculateWordWrap(text string) []interface{}{
	array00 := self.Object.Call("precalculateWordWrap", text)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Runs the given text through the Text.runWordWrap function and returns
// the results as an array, where each element of the array corresponds to a wrapped
// line of text.
// 
// Useful if you wish to control pagination on long pieces of content.
func (self *Text) PrecalculateWordWrapI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("precalculateWordWrap", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Greedy wrapping algorithm that will wrap words as the line grows longer than its horizontal bounds.
func (self *Text) RunWordWrap(text string) {
    self.Object.Call("runWordWrap", text)
}

// Greedy wrapping algorithm that will wrap words as the line grows longer than its horizontal bounds.
func (self *Text) RunWordWrapI(args ...interface{}) {
    self.Object.Call("runWordWrap", args)
}

// Advanced wrapping algorithm that will wrap words as the line grows longer than its horizontal bounds.
// White space is condensed (e.g., consecutive spaces are replaced with one).
// Lines are trimmed of white space before processing.
// Throws an error if the user was smart enough to specify a wordWrapWidth less than a single character.
func (self *Text) AdvancedWordWrap(text string) {
    self.Object.Call("advancedWordWrap", text)
}

// Advanced wrapping algorithm that will wrap words as the line grows longer than its horizontal bounds.
// White space is condensed (e.g., consecutive spaces are replaced with one).
// Lines are trimmed of white space before processing.
// Throws an error if the user was smart enough to specify a wordWrapWidth less than a single character.
func (self *Text) AdvancedWordWrapI(args ...interface{}) {
    self.Object.Call("advancedWordWrap", args)
}

// Greedy wrapping algorithm that will wrap words as the line grows longer than its horizontal bounds.
func (self *Text) BasicWordWrap(text string) {
    self.Object.Call("basicWordWrap", text)
}

// Greedy wrapping algorithm that will wrap words as the line grows longer than its horizontal bounds.
func (self *Text) BasicWordWrapI(args ...interface{}) {
    self.Object.Call("basicWordWrap", args)
}

// Updates the internal `style.font` if it now differs according to generation from components.
func (self *Text) UpdateFont(components interface{}) {
    self.Object.Call("updateFont", components)
}

// Updates the internal `style.font` if it now differs according to generation from components.
func (self *Text) UpdateFontI(args ...interface{}) {
    self.Object.Call("updateFont", args)
}

// Converting a short CSS-font string into the relevant components.
func (self *Text) FontToComponents(font string) {
    self.Object.Call("fontToComponents", font)
}

// Converting a short CSS-font string into the relevant components.
func (self *Text) FontToComponentsI(args ...interface{}) {
    self.Object.Call("fontToComponents", args)
}

// Converts individual font components (see `fontToComponents`) to a short CSS font string.
func (self *Text) ComponentsToFont(components interface{}) {
    self.Object.Call("componentsToFont", components)
}

// Converts individual font components (see `fontToComponents`) to a short CSS font string.
func (self *Text) ComponentsToFontI(args ...interface{}) {
    self.Object.Call("componentsToFont", args)
}

// The text to be displayed by this Text object.
// Use a \n to insert a carriage return and split the text.
// The text will be rendered with any style currently set.
// 
// Use the optional `immediate` argument if you need the Text display to update immediately.
// 
// If not it will re-create the texture of this Text object during the next time the render
// loop is called.
func (self *Text) SetText() *Text{
    return &Text{self.Object.Call("setText")}
}

// The text to be displayed by this Text object.
// Use a \n to insert a carriage return and split the text.
// The text will be rendered with any style currently set.
// 
// Use the optional `immediate` argument if you need the Text display to update immediately.
// 
// If not it will re-create the texture of this Text object during the next time the render
// loop is called.
func (self *Text) SetText1O(text string) *Text{
    return &Text{self.Object.Call("setText", text)}
}

// The text to be displayed by this Text object.
// Use a \n to insert a carriage return and split the text.
// The text will be rendered with any style currently set.
// 
// Use the optional `immediate` argument if you need the Text display to update immediately.
// 
// If not it will re-create the texture of this Text object during the next time the render
// loop is called.
func (self *Text) SetText2O(text string, immediate bool) *Text{
    return &Text{self.Object.Call("setText", text, immediate)}
}

// The text to be displayed by this Text object.
// Use a \n to insert a carriage return and split the text.
// The text will be rendered with any style currently set.
// 
// Use the optional `immediate` argument if you need the Text display to update immediately.
// 
// If not it will re-create the texture of this Text object during the next time the render
// loop is called.
func (self *Text) SetTextI(args ...interface{}) *Text{
    return &Text{self.Object.Call("setText", args)}
}

// Converts the given array into a tab delimited string and then updates this Text object.
// This is mostly used when you want to display external data using tab stops.
// 
// The array can be either single or multi dimensional depending on the result you need:
// 
// `[ 'a', 'b', 'c' ]` would convert in to `"a\tb\tc"`.
// 
// Where as:
// 
// `[
//      [ 'a', 'b', 'c' ],
//      [ 'd', 'e', 'f']
// ]`
// 
// would convert in to: `"a\tb\tc\nd\te\tf"`
func (self *Text) ParseList(list []interface{}) *Text{
    return &Text{self.Object.Call("parseList", list)}
}

// Converts the given array into a tab delimited string and then updates this Text object.
// This is mostly used when you want to display external data using tab stops.
// 
// The array can be either single or multi dimensional depending on the result you need:
// 
// `[ 'a', 'b', 'c' ]` would convert in to `"a\tb\tc"`.
// 
// Where as:
// 
// `[
//      [ 'a', 'b', 'c' ],
//      [ 'd', 'e', 'f']
// ]`
// 
// would convert in to: `"a\tb\tc\nd\te\tf"`
func (self *Text) ParseListI(args ...interface{}) *Text{
    return &Text{self.Object.Call("parseList", args)}
}

// The Text Bounds is a rectangular region that you control the dimensions of into which the Text object itself is positioned,
// regardless of the number of lines in the text, the font size or any other attribute.
// 
// Alignment is controlled via the properties `boundsAlignH` and `boundsAlignV` within the Text.style object, or can be directly
// set through the setters `Text.boundsAlignH` and `Text.boundsAlignV`. Bounds alignment is independent of text alignment.
// 
// For example: If your game is 800x600 in size and you set the text bounds to be 0,0,800,600 then by setting boundsAlignH to
// 'center' and boundsAlignV to 'bottom' the text will render in the center and at the bottom of your game window, regardless of
// how many lines of text there may be. Even if you adjust the text content or change the style it will remain at the bottom center
// of the text bounds.
// 
// This is especially powerful when you need to align text against specific coordinates in your game, but the actual text dimensions
// may vary based on font (say for multi-lingual games).
// 
// If `Text.wordWrapWidth` is greater than the width of the text bounds it is clamped to match the bounds width.
// 
// Call this method with no arguments given to reset an existing textBounds.
// 
// It works by calculating the final position based on the Text.canvas size, which is modified as the text is updated. Some fonts
// have additional padding around them which you can mitigate by tweaking the Text.padding property. It then adjusts the `pivot`
// property based on the given bounds and canvas size. This means if you need to set the pivot property directly in your game then
// you either cannot use `setTextBounds` or you must place the Text object inside another DisplayObject on which you set the pivot.
func (self *Text) SetTextBounds() *Text{
    return &Text{self.Object.Call("setTextBounds")}
}

// The Text Bounds is a rectangular region that you control the dimensions of into which the Text object itself is positioned,
// regardless of the number of lines in the text, the font size or any other attribute.
// 
// Alignment is controlled via the properties `boundsAlignH` and `boundsAlignV` within the Text.style object, or can be directly
// set through the setters `Text.boundsAlignH` and `Text.boundsAlignV`. Bounds alignment is independent of text alignment.
// 
// For example: If your game is 800x600 in size and you set the text bounds to be 0,0,800,600 then by setting boundsAlignH to
// 'center' and boundsAlignV to 'bottom' the text will render in the center and at the bottom of your game window, regardless of
// how many lines of text there may be. Even if you adjust the text content or change the style it will remain at the bottom center
// of the text bounds.
// 
// This is especially powerful when you need to align text against specific coordinates in your game, but the actual text dimensions
// may vary based on font (say for multi-lingual games).
// 
// If `Text.wordWrapWidth` is greater than the width of the text bounds it is clamped to match the bounds width.
// 
// Call this method with no arguments given to reset an existing textBounds.
// 
// It works by calculating the final position based on the Text.canvas size, which is modified as the text is updated. Some fonts
// have additional padding around them which you can mitigate by tweaking the Text.padding property. It then adjusts the `pivot`
// property based on the given bounds and canvas size. This means if you need to set the pivot property directly in your game then
// you either cannot use `setTextBounds` or you must place the Text object inside another DisplayObject on which you set the pivot.
func (self *Text) SetTextBounds1O(x int) *Text{
    return &Text{self.Object.Call("setTextBounds", x)}
}

// The Text Bounds is a rectangular region that you control the dimensions of into which the Text object itself is positioned,
// regardless of the number of lines in the text, the font size or any other attribute.
// 
// Alignment is controlled via the properties `boundsAlignH` and `boundsAlignV` within the Text.style object, or can be directly
// set through the setters `Text.boundsAlignH` and `Text.boundsAlignV`. Bounds alignment is independent of text alignment.
// 
// For example: If your game is 800x600 in size and you set the text bounds to be 0,0,800,600 then by setting boundsAlignH to
// 'center' and boundsAlignV to 'bottom' the text will render in the center and at the bottom of your game window, regardless of
// how many lines of text there may be. Even if you adjust the text content or change the style it will remain at the bottom center
// of the text bounds.
// 
// This is especially powerful when you need to align text against specific coordinates in your game, but the actual text dimensions
// may vary based on font (say for multi-lingual games).
// 
// If `Text.wordWrapWidth` is greater than the width of the text bounds it is clamped to match the bounds width.
// 
// Call this method with no arguments given to reset an existing textBounds.
// 
// It works by calculating the final position based on the Text.canvas size, which is modified as the text is updated. Some fonts
// have additional padding around them which you can mitigate by tweaking the Text.padding property. It then adjusts the `pivot`
// property based on the given bounds and canvas size. This means if you need to set the pivot property directly in your game then
// you either cannot use `setTextBounds` or you must place the Text object inside another DisplayObject on which you set the pivot.
func (self *Text) SetTextBounds2O(x int, y int) *Text{
    return &Text{self.Object.Call("setTextBounds", x, y)}
}

// The Text Bounds is a rectangular region that you control the dimensions of into which the Text object itself is positioned,
// regardless of the number of lines in the text, the font size or any other attribute.
// 
// Alignment is controlled via the properties `boundsAlignH` and `boundsAlignV` within the Text.style object, or can be directly
// set through the setters `Text.boundsAlignH` and `Text.boundsAlignV`. Bounds alignment is independent of text alignment.
// 
// For example: If your game is 800x600 in size and you set the text bounds to be 0,0,800,600 then by setting boundsAlignH to
// 'center' and boundsAlignV to 'bottom' the text will render in the center and at the bottom of your game window, regardless of
// how many lines of text there may be. Even if you adjust the text content or change the style it will remain at the bottom center
// of the text bounds.
// 
// This is especially powerful when you need to align text against specific coordinates in your game, but the actual text dimensions
// may vary based on font (say for multi-lingual games).
// 
// If `Text.wordWrapWidth` is greater than the width of the text bounds it is clamped to match the bounds width.
// 
// Call this method with no arguments given to reset an existing textBounds.
// 
// It works by calculating the final position based on the Text.canvas size, which is modified as the text is updated. Some fonts
// have additional padding around them which you can mitigate by tweaking the Text.padding property. It then adjusts the `pivot`
// property based on the given bounds and canvas size. This means if you need to set the pivot property directly in your game then
// you either cannot use `setTextBounds` or you must place the Text object inside another DisplayObject on which you set the pivot.
func (self *Text) SetTextBounds3O(x int, y int, width int) *Text{
    return &Text{self.Object.Call("setTextBounds", x, y, width)}
}

// The Text Bounds is a rectangular region that you control the dimensions of into which the Text object itself is positioned,
// regardless of the number of lines in the text, the font size or any other attribute.
// 
// Alignment is controlled via the properties `boundsAlignH` and `boundsAlignV` within the Text.style object, or can be directly
// set through the setters `Text.boundsAlignH` and `Text.boundsAlignV`. Bounds alignment is independent of text alignment.
// 
// For example: If your game is 800x600 in size and you set the text bounds to be 0,0,800,600 then by setting boundsAlignH to
// 'center' and boundsAlignV to 'bottom' the text will render in the center and at the bottom of your game window, regardless of
// how many lines of text there may be. Even if you adjust the text content or change the style it will remain at the bottom center
// of the text bounds.
// 
// This is especially powerful when you need to align text against specific coordinates in your game, but the actual text dimensions
// may vary based on font (say for multi-lingual games).
// 
// If `Text.wordWrapWidth` is greater than the width of the text bounds it is clamped to match the bounds width.
// 
// Call this method with no arguments given to reset an existing textBounds.
// 
// It works by calculating the final position based on the Text.canvas size, which is modified as the text is updated. Some fonts
// have additional padding around them which you can mitigate by tweaking the Text.padding property. It then adjusts the `pivot`
// property based on the given bounds and canvas size. This means if you need to set the pivot property directly in your game then
// you either cannot use `setTextBounds` or you must place the Text object inside another DisplayObject on which you set the pivot.
func (self *Text) SetTextBounds4O(x int, y int, width int, height int) *Text{
    return &Text{self.Object.Call("setTextBounds", x, y, width, height)}
}

// The Text Bounds is a rectangular region that you control the dimensions of into which the Text object itself is positioned,
// regardless of the number of lines in the text, the font size or any other attribute.
// 
// Alignment is controlled via the properties `boundsAlignH` and `boundsAlignV` within the Text.style object, or can be directly
// set through the setters `Text.boundsAlignH` and `Text.boundsAlignV`. Bounds alignment is independent of text alignment.
// 
// For example: If your game is 800x600 in size and you set the text bounds to be 0,0,800,600 then by setting boundsAlignH to
// 'center' and boundsAlignV to 'bottom' the text will render in the center and at the bottom of your game window, regardless of
// how many lines of text there may be. Even if you adjust the text content or change the style it will remain at the bottom center
// of the text bounds.
// 
// This is especially powerful when you need to align text against specific coordinates in your game, but the actual text dimensions
// may vary based on font (say for multi-lingual games).
// 
// If `Text.wordWrapWidth` is greater than the width of the text bounds it is clamped to match the bounds width.
// 
// Call this method with no arguments given to reset an existing textBounds.
// 
// It works by calculating the final position based on the Text.canvas size, which is modified as the text is updated. Some fonts
// have additional padding around them which you can mitigate by tweaking the Text.padding property. It then adjusts the `pivot`
// property based on the given bounds and canvas size. This means if you need to set the pivot property directly in your game then
// you either cannot use `setTextBounds` or you must place the Text object inside another DisplayObject on which you set the pivot.
func (self *Text) SetTextBoundsI(args ...interface{}) *Text{
    return &Text{self.Object.Call("setTextBounds", args)}
}

// Updates the texture based on the canvas dimensions.
func (self *Text) UpdateTexture() {
    self.Object.Call("updateTexture")
}

// Updates the texture based on the canvas dimensions.
func (self *Text) UpdateTextureI(args ...interface{}) {
    self.Object.Call("updateTexture", args)
}

// Renders the object using the WebGL renderer
func (self *Text) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// Renders the object using the WebGL renderer
func (self *Text) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer.
func (self *Text) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// Renders the object using the Canvas renderer.
func (self *Text) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

// Calculates the ascent, descent and fontSize of a given font style.
func (self *Text) DetermineFontProperties(fontStyle interface{}) {
    self.Object.Call("determineFontProperties", fontStyle)
}

// Calculates the ascent, descent and fontSize of a given font style.
func (self *Text) DetermineFontPropertiesI(args ...interface{}) {
    self.Object.Call("determineFontProperties", args)
}

// Returns the bounds of the Text as a rectangle.
// The bounds calculation takes the worldTransform into account.
func (self *Text) GetBounds(matrix *Matrix) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", matrix)}
}

// Returns the bounds of the Text as a rectangle.
// The bounds calculation takes the worldTransform into account.
func (self *Text) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Text) SetTexture(texture *Texture) {
    self.Object.Call("setTexture", texture)
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Text) SetTexture1O(texture *Texture, destroy bool) {
    self.Object.Call("setTexture", texture, destroy)
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Text) SetTextureI(args ...interface{}) {
    self.Object.Call("setTexture", args)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *Text) OnTextureUpdate(event interface{}) {
    self.Object.Call("onTextureUpdate", event)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *Text) OnTextureUpdateI(args ...interface{}) {
    self.Object.Call("onTextureUpdate", args)
}

// Adds a child to the container.
func (self *Text) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// Adds a child to the container.
func (self *Text) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Text) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Text) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *Text) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// Swaps the position of 2 Display Objects within this container.
func (self *Text) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *Text) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// Returns the index position of a child DisplayObject instance
func (self *Text) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *Text) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// Changes the position of an existing child in the display object container
func (self *Text) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *Text) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// Returns the child at the specified index
func (self *Text) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *Text) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// Removes a child from the container.
func (self *Text) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *Text) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// Removes a child from the specified index position.
func (self *Text) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Text) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Text) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Text) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Text) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Text) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Text) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *Text) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// Removes the current stage reference from the container and all of its children.
func (self *Text) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// Internal method called by the World postUpdate cycle.
func (self *Text) PostUpdate() {
    self.Object.Call("postUpdate")
}

// Internal method called by the World postUpdate cycle.
func (self *Text) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Text) Play(name string) *Animation{
    return &Animation{self.Object.Call("play", name)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Text) Play1O(name string, frameRate int) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Text) Play2O(name string, frameRate int, loop bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Text) Play3O(name string, frameRate int, loop bool, killOnComplete bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop, killOnComplete)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Text) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("play", args)}
}

// Aligns this Game Object within another Game Object, or Rectangle, known as the
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
func (self *Text) AlignIn(container interface{}) interface{}{
    return self.Object.Call("alignIn", container)
}

// Aligns this Game Object within another Game Object, or Rectangle, known as the
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
func (self *Text) AlignIn1O(container interface{}, position int) interface{}{
    return self.Object.Call("alignIn", container, position)
}

// Aligns this Game Object within another Game Object, or Rectangle, known as the
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
func (self *Text) AlignIn2O(container interface{}, position int, offsetX int) interface{}{
    return self.Object.Call("alignIn", container, position, offsetX)
}

// Aligns this Game Object within another Game Object, or Rectangle, known as the
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
func (self *Text) AlignIn3O(container interface{}, position int, offsetX int, offsetY int) interface{}{
    return self.Object.Call("alignIn", container, position, offsetX, offsetY)
}

// Aligns this Game Object within another Game Object, or Rectangle, known as the
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
func (self *Text) AlignInI(args ...interface{}) interface{}{
    return self.Object.Call("alignIn", args)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
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
func (self *Text) AlignTo(parent interface{}) interface{}{
    return self.Object.Call("alignTo", parent)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
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
func (self *Text) AlignTo1O(parent interface{}, position int) interface{}{
    return self.Object.Call("alignTo", parent, position)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
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
func (self *Text) AlignTo2O(parent interface{}, position int, offsetX int) interface{}{
    return self.Object.Call("alignTo", parent, position, offsetX)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
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
func (self *Text) AlignTo3O(parent interface{}, position int, offsetX int, offsetY int) interface{}{
    return self.Object.Call("alignTo", parent, position, offsetX, offsetY)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
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
func (self *Text) AlignToI(args ...interface{}) interface{}{
    return self.Object.Call("alignTo", args)
}

// Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) BringToTop() *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop")}
}

// Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) BringToTopI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop", args)}
}

// Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) SendToBack() *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack")}
}

// Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) SendToBackI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack", args)}
}

// Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) MoveUp() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp")}
}

// Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) MoveUpI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp", args)}
}

// Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) MoveDown() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown")}
}

// Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) MoveDownI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown", args)}
}

// Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
// 
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
// 
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object 
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
// 
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`, 
// in which case the values are duplicated to a local object.
func (self *Text) Crop(rect *Rectangle) {
    self.Object.Call("crop", rect)
}

// Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
// 
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
// 
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object 
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
// 
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`, 
// in which case the values are duplicated to a local object.
func (self *Text) Crop1O(rect *Rectangle, copy bool) {
    self.Object.Call("crop", rect, copy)
}

// Crop allows you to crop the texture being used to display this Game Object.
// Setting a crop rectangle modifies the core texture frame. The Game Object width and height properties will be adjusted accordingly.
// 
// Cropping takes place from the top-left and can be modified in real-time either by providing an updated rectangle object to this method,
// or by modifying `cropRect` property directly and then calling `updateCrop`.
// 
// The rectangle object given to this method can be either a `Phaser.Rectangle` or any other object 
// so long as it has public `x`, `y`, `width`, `height`, `right` and `bottom` properties.
// 
// A reference to the rectangle is stored in `cropRect` unless the `copy` parameter is `true`, 
// in which case the values are duplicated to a local object.
func (self *Text) CropI(args ...interface{}) {
    self.Object.Call("crop", args)
}

// If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *Text) UpdateCrop() {
    self.Object.Call("updateCrop")
}

// If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *Text) UpdateCropI(args ...interface{}) {
    self.Object.Call("updateCrop", args)
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Text) Revive() *DisplayObject{
    return &DisplayObject{self.Object.Call("revive")}
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Text) Revive1O(health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", health)}
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Text) ReviveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("revive", args)}
}

// Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *Text) Kill() *DisplayObject{
    return &DisplayObject{self.Object.Call("kill")}
}

// Kills a Game Object. A killed Game Object has its `alive`, `exists` and `visible` properties all set to false.
// 
// It will dispatch the `onKilled` event. You can listen to `events.onKilled` for the signal.
// 
// Note that killing a Game Object is a way for you to quickly recycle it in an object pool,
// it doesn't destroy the object or free it up from memory.
// 
// If you don't need this Game Object any more you should call `destroy` instead.
func (self *Text) KillI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("kill", args)}
}

// Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *Text) LoadTexture(key interface{}) {
    self.Object.Call("loadTexture", key)
}

// Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *Text) LoadTexture1O(key interface{}, frame interface{}) {
    self.Object.Call("loadTexture", key, frame)
}

// Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *Text) LoadTexture2O(key interface{}, frame interface{}, stopAnimation bool) {
    self.Object.Call("loadTexture", key, frame, stopAnimation)
}

// Changes the base texture the Game Object is using. The old texture is removed and the new one is referenced or fetched from the Cache.
// 
// If your Game Object is using a frame from a texture atlas and you just wish to change to another frame, then see the `frame` or `frameName` properties instead.
// 
// You should only use `loadTexture` if you want to replace the base texture entirely.
// 
// Calling this method causes a WebGL texture update, so use sparingly or in low-intensity portions of your game, or if you know the new texture is already on the GPU.
// 
// You can use the new const `Phaser.PENDING_ATLAS` as the texture key for any sprite. 
// Doing this then sets the key to be the `frame` argument (the frame is set to zero). 
// 
// This allows you to create sprites using `load.image` during development, and then change them 
// to use a Texture Atlas later in development by simply searching your code for 'PENDING_ATLAS' 
// and swapping it to be the key of the atlas data.
// 
// Note: You cannot use a RenderTexture as a texture for a TileSprite.
func (self *Text) LoadTextureI(args ...interface{}) {
    self.Object.Call("loadTexture", args)
}

// Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *Text) SetFrame(frame *Frame) {
    self.Object.Call("setFrame", frame)
}

// Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *Text) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *Text) ResizeFrame(parent interface{}, width int, height int) {
    self.Object.Call("resizeFrame", parent, width, height)
}

// Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *Text) ResizeFrameI(args ...interface{}) {
    self.Object.Call("resizeFrame", args)
}

// Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *Text) ResetFrame() {
    self.Object.Call("resetFrame")
}

// Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *Text) ResetFrameI(args ...interface{}) {
    self.Object.Call("resetFrame", args)
}

// Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *Text) Overlap(displayObject interface{}) bool{
    return self.Object.Call("overlap", displayObject).Bool()
}

// Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *Text) OverlapI(args ...interface{}) bool{
    return self.Object.Call("overlap", args).Bool()
}

// Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *Text) Reset(x int, y int) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", x, y)}
}

// Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *Text) Reset1O(x int, y int, health int) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", x, y, health)}
}

// Resets the Game Object.
// 
// This moves the Game Object to the given x/y world coordinates and sets `fresh`, `exists`, 
// `visible` and `renderable` to true.
// 
// If this Game Object has the LifeSpan component it will also set `alive` to true and `health` to the given value.
// 
// If this Game Object has a Physics Body it will reset the Body.
func (self *Text) ResetI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("reset", args)}
}

// Adjust scaling limits, if set, to this Game Object.
func (self *Text) CheckTransform(wt *Matrix) {
    self.Object.Call("checkTransform", wt)
}

// Adjust scaling limits, if set, to this Game Object.
func (self *Text) CheckTransformI(args ...interface{}) {
    self.Object.Call("checkTransform", args)
}

// Sets the scaleMin and scaleMax values. These values are used to limit how far this Game Object will scale based on its parent.
// 
// For example if this Game Object has a `minScale` value of 1 and its parent has a `scale` value of 0.5, the 0.5 will be ignored 
// and the scale value of 1 will be used, as the parents scale is lower than the minimum scale this Game Object should adhere to.
// 
// By setting these values you can carefully control how Game Objects deal with responsive scaling.
// 
// If only one parameter is given then that value will be used for both scaleMin and scaleMax:
// `setScaleMinMax(1)` = scaleMin.x, scaleMin.y, scaleMax.x and scaleMax.y all = 1
// 
// If only two parameters are given the first is set as scaleMin.x and y and the second as scaleMax.x and y:
// `setScaleMinMax(0.5, 2)` = scaleMin.x and y = 0.5 and scaleMax.x and y = 2
// 
// If you wish to set `scaleMin` with different values for x and y then either modify Game Object.scaleMin directly, 
// or pass `null` for the `maxX` and `maxY` parameters.
// 
// Call `setScaleMinMax(null)` to clear all previously set values.
func (self *Text) SetScaleMinMax(minX interface{}, minY interface{}, maxX interface{}, maxY interface{}) {
    self.Object.Call("setScaleMinMax", minX, minY, maxX, maxY)
}

// Sets the scaleMin and scaleMax values. These values are used to limit how far this Game Object will scale based on its parent.
// 
// For example if this Game Object has a `minScale` value of 1 and its parent has a `scale` value of 0.5, the 0.5 will be ignored 
// and the scale value of 1 will be used, as the parents scale is lower than the minimum scale this Game Object should adhere to.
// 
// By setting these values you can carefully control how Game Objects deal with responsive scaling.
// 
// If only one parameter is given then that value will be used for both scaleMin and scaleMax:
// `setScaleMinMax(1)` = scaleMin.x, scaleMin.y, scaleMax.x and scaleMax.y all = 1
// 
// If only two parameters are given the first is set as scaleMin.x and y and the second as scaleMax.x and y:
// `setScaleMinMax(0.5, 2)` = scaleMin.x and y = 0.5 and scaleMax.x and y = 2
// 
// If you wish to set `scaleMin` with different values for x and y then either modify Game Object.scaleMin directly, 
// or pass `null` for the `maxX` and `maxY` parameters.
// 
// Call `setScaleMinMax(null)` to clear all previously set values.
func (self *Text) SetScaleMinMaxI(args ...interface{}) {
    self.Object.Call("setScaleMinMax", args)
}
