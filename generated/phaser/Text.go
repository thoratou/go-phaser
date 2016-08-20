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
func (self *Text) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *Text) SetType(member float64) {
    self.Set("type", member)
}

// The const physics body type of this object.
func (self *Text) GetPhysicsType() float64{
    return self.Get("physicsType").Float()
}

// The const physics body type of this object.
func (self *Text) SetPhysicsType(member float64) {
    self.Set("physicsType", member)
}

// Specify a padding value which is added to the line width and height when calculating the Text size.
// ALlows you to add extra spacing if Phaser is unable to accurately determine the true font dimensions.
func (self *Text) GetPadding() *Point{
    return &Point{self.Get("padding")}
}

// Specify a padding value which is added to the line width and height when calculating the Text size.
// ALlows you to add extra spacing if Phaser is unable to accurately determine the true font dimensions.
func (self *Text) SetPadding(member *Point) {
    self.Set("padding", member)
}

// The textBounds property allows you to specify a rectangular region upon which text alignment is based.
// See `Text.setTextBounds` for more details.
func (self *Text) GetTextBounds() *Rectangle{
    return &Rectangle{self.Get("textBounds")}
}

// The textBounds property allows you to specify a rectangular region upon which text alignment is based.
// See `Text.setTextBounds` for more details.
func (self *Text) SetTextBounds(member *Rectangle) {
    self.Set("textBounds", member)
}

// The canvas element that the text is rendered.
func (self *Text) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("canvas"))
}

// The canvas element that the text is rendered.
func (self *Text) SetCanvas(member dom.HTMLCanvasElement) {
    self.Set("canvas", member)
}

// The context of the canvas element that the text is rendered to.
func (self *Text) GetContext() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("context"))
}

// The context of the canvas element that the text is rendered to.
func (self *Text) SetContext(member dom.HTMLCanvasElement) {
    self.Set("context", member)
}

// An array of the color values as specified by {@link Phaser.Text#addColor addColor}.
func (self *Text) GetColors() []interface{}{
	array := self.Get("colors")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of the color values as specified by {@link Phaser.Text#addColor addColor}.
func (self *Text) SetColors(member []interface{}) {
    self.Set("colors", member)
}

// An array of the stroke color values as specified by {@link Phaser.Text#addStrokeColor addStrokeColor}.
func (self *Text) GetStrokeColors() []interface{}{
	array := self.Get("strokeColors")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of the stroke color values as specified by {@link Phaser.Text#addStrokeColor addStrokeColor}.
func (self *Text) SetStrokeColors(member []interface{}) {
    self.Set("strokeColors", member)
}

// An array of the font styles values as specified by {@link Phaser.Text#addFontStyle addFontStyle}.
func (self *Text) GetFontStyles() []interface{}{
	array := self.Get("fontStyles")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of the font styles values as specified by {@link Phaser.Text#addFontStyle addFontStyle}.
func (self *Text) SetFontStyles(member []interface{}) {
    self.Set("fontStyles", member)
}

// An array of the font weights values as specified by {@link Phaser.Text#addFontWeight addFontWeight}.
func (self *Text) GetFontWeights() []interface{}{
	array := self.Get("fontWeights")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An array of the font weights values as specified by {@link Phaser.Text#addFontWeight addFontWeight}.
func (self *Text) SetFontWeights(member []interface{}) {
    self.Set("fontWeights", member)
}

// Should the linePositionX and Y values be automatically rounded before rendering the Text?
// You may wish to enable this if you want to remove the effect of sub-pixel aliasing from text.
func (self *Text) GetAutoRound() bool{
    return self.Get("autoRound").Bool()
}

// Should the linePositionX and Y values be automatically rounded before rendering the Text?
// You may wish to enable this if you want to remove the effect of sub-pixel aliasing from text.
func (self *Text) SetAutoRound(member bool) {
    self.Set("autoRound", member)
}

// Will this Text object use Basic or Advanced Word Wrapping?
// 
// Advanced wrapping breaks long words if they are the first of a line, and repeats the process as necessary.
// White space is condensed (e.g., consecutive spaces are replaced with one).
// Lines are trimmed of white space before processing.
// 
// It throws an error if wordWrapWidth is less than a single character.
func (self *Text) GetUseAdvancedWrap() bool{
    return self.Get("useAdvancedWrap").Bool()
}

// Will this Text object use Basic or Advanced Word Wrapping?
// 
// Advanced wrapping breaks long words if they are the first of a line, and repeats the process as necessary.
// White space is condensed (e.g., consecutive spaces are replaced with one).
// Lines are trimmed of white space before processing.
// 
// It throws an error if wordWrapWidth is less than a single character.
func (self *Text) SetUseAdvancedWrap(member bool) {
    self.Set("useAdvancedWrap", member)
}

// The text to be displayed by this Text object.
// Use a \n to insert a carriage return and split the text.
// The text will be rendered with any style currently set.
func (self *Text) GetText() string{
    return self.Get("text").String()
}

// The text to be displayed by this Text object.
// Use a \n to insert a carriage return and split the text.
// The text will be rendered with any style currently set.
func (self *Text) SetText(member string) {
    self.Set("text", member)
}

// Change the font used.
// 
// This is equivalent of the `font` property specified to {@link Phaser.Text#setStyle setStyle}, except
// that unlike using `setStyle` this will not change any current font fill/color settings.
// 
// The CSS font string can also be individually altered with the `font`, `fontSize`, `fontWeight`, `fontStyle`, and `fontVariant` properties.
func (self *Text) GetCssFont() string{
    return self.Get("cssFont").String()
}

// Change the font used.
// 
// This is equivalent of the `font` property specified to {@link Phaser.Text#setStyle setStyle}, except
// that unlike using `setStyle` this will not change any current font fill/color settings.
// 
// The CSS font string can also be individually altered with the `font`, `fontSize`, `fontWeight`, `fontStyle`, and `fontVariant` properties.
func (self *Text) SetCssFont(member string) {
    self.Set("cssFont", member)
}

// Change the font family that the text will be rendered in, such as 'Arial'.
// 
// Multiple CSS font families and generic fallbacks can be specified as long as
// {@link http://www.w3.org/TR/CSS2/fonts.html#propdef-font-family CSS font-family rules} are followed.
// 
// To change the entire font string use {@link Phaser.Text#cssFont cssFont} instead: eg. `text.cssFont = 'bold 20pt Arial'`.
func (self *Text) GetFont() string{
    return self.Get("font").String()
}

// Change the font family that the text will be rendered in, such as 'Arial'.
// 
// Multiple CSS font families and generic fallbacks can be specified as long as
// {@link http://www.w3.org/TR/CSS2/fonts.html#propdef-font-family CSS font-family rules} are followed.
// 
// To change the entire font string use {@link Phaser.Text#cssFont cssFont} instead: eg. `text.cssFont = 'bold 20pt Arial'`.
func (self *Text) SetFont(member string) {
    self.Set("font", member)
}

// The size of the font.
// 
// If the font size is specified in pixels (eg. `32` or `'32px`') then a number (ie. `32`) representing
// the font size in pixels is returned; otherwise the value with CSS unit is returned as a string (eg. `'12pt'`).
func (self *Text) GetFontSize() interface{}{
    return self.Get("fontSize")
}

// The size of the font.
// 
// If the font size is specified in pixels (eg. `32` or `'32px`') then a number (ie. `32`) representing
// the font size in pixels is returned; otherwise the value with CSS unit is returned as a string (eg. `'12pt'`).
func (self *Text) SetFontSize(member interface{}) {
    self.Set("fontSize", member)
}

// The weight of the font: 'normal', 'bold', or {@link http://www.w3.org/TR/CSS2/fonts.html#propdef-font-weight a valid CSS font weight}.
func (self *Text) GetFontWeight() string{
    return self.Get("fontWeight").String()
}

// The weight of the font: 'normal', 'bold', or {@link http://www.w3.org/TR/CSS2/fonts.html#propdef-font-weight a valid CSS font weight}.
func (self *Text) SetFontWeight(member string) {
    self.Set("fontWeight", member)
}

// The style of the font: 'normal', 'italic', 'oblique'
func (self *Text) GetFontStyle() string{
    return self.Get("fontStyle").String()
}

// The style of the font: 'normal', 'italic', 'oblique'
func (self *Text) SetFontStyle(member string) {
    self.Set("fontStyle", member)
}

// The variant the font: 'normal', 'small-caps'
func (self *Text) GetFontVariant() string{
    return self.Get("fontVariant").String()
}

// The variant the font: 'normal', 'small-caps'
func (self *Text) SetFontVariant(member string) {
    self.Set("fontVariant", member)
}

// A canvas fillstyle that will be used on the text eg 'red', '#00FF00'.
func (self *Text) GetFill() interface{}{
    return self.Get("fill")
}

// A canvas fillstyle that will be used on the text eg 'red', '#00FF00'.
func (self *Text) SetFill(member interface{}) {
    self.Set("fill", member)
}

// Controls the horizontal alignment for multiline text.
// Can be: 'left', 'center' or 'right'.
// Does not affect single lines of text. For that please see `setTextBounds`.
func (self *Text) GetAlign() string{
    return self.Get("align").String()
}

// Controls the horizontal alignment for multiline text.
// Can be: 'left', 'center' or 'right'.
// Does not affect single lines of text. For that please see `setTextBounds`.
func (self *Text) SetAlign(member string) {
    self.Set("align", member)
}

// The resolution of the canvas the text is rendered to.
// This defaults to match the resolution of the renderer, but can be changed on a per Text object basis.
func (self *Text) GetResolution() int{
    return self.Get("resolution").Int()
}

// The resolution of the canvas the text is rendered to.
// This defaults to match the resolution of the renderer, but can be changed on a per Text object basis.
func (self *Text) SetResolution(member int) {
    self.Set("resolution", member)
}

// The size (in pixels) of the tabs, for when text includes tab characters. 0 disables. 
// Can be an integer or an array of varying tab sizes, one tab per element.
// For example if you set tabs to 100 then when Text encounters a tab it will jump ahead 100 pixels.
// If you set tabs to be `[100,200]` then it will set the first tab at 100px and the second at 200px.
func (self *Text) GetTabs() interface{}{
    return self.Get("tabs")
}

// The size (in pixels) of the tabs, for when text includes tab characters. 0 disables. 
// Can be an integer or an array of varying tab sizes, one tab per element.
// For example if you set tabs to 100 then when Text encounters a tab it will jump ahead 100 pixels.
// If you set tabs to be `[100,200]` then it will set the first tab at 100px and the second at 200px.
func (self *Text) SetTabs(member interface{}) {
    self.Set("tabs", member)
}

// Horizontal alignment of the text within the `textBounds`. Can be: 'left', 'center' or 'right'.
func (self *Text) GetBoundsAlignH() string{
    return self.Get("boundsAlignH").String()
}

// Horizontal alignment of the text within the `textBounds`. Can be: 'left', 'center' or 'right'.
func (self *Text) SetBoundsAlignH(member string) {
    self.Set("boundsAlignH", member)
}

// Vertical alignment of the text within the `textBounds`. Can be: 'top', 'middle' or 'bottom'.
func (self *Text) GetBoundsAlignV() string{
    return self.Get("boundsAlignV").String()
}

// Vertical alignment of the text within the `textBounds`. Can be: 'top', 'middle' or 'bottom'.
func (self *Text) SetBoundsAlignV(member string) {
    self.Set("boundsAlignV", member)
}

// A canvas fillstyle that will be used on the text stroke eg 'blue', '#FCFF00'.
func (self *Text) GetStroke() string{
    return self.Get("stroke").String()
}

// A canvas fillstyle that will be used on the text stroke eg 'blue', '#FCFF00'.
func (self *Text) SetStroke(member string) {
    self.Set("stroke", member)
}

// A number that represents the thickness of the stroke. Default is 0 (no stroke)
func (self *Text) GetStrokeThickness() float64{
    return self.Get("strokeThickness").Float()
}

// A number that represents the thickness of the stroke. Default is 0 (no stroke)
func (self *Text) SetStrokeThickness(member float64) {
    self.Set("strokeThickness", member)
}

// Indicates if word wrap should be used.
func (self *Text) GetWordWrap() bool{
    return self.Get("wordWrap").Bool()
}

// Indicates if word wrap should be used.
func (self *Text) SetWordWrap(member bool) {
    self.Set("wordWrap", member)
}

// The width at which text will wrap.
func (self *Text) GetWordWrapWidth() float64{
    return self.Get("wordWrapWidth").Float()
}

// The width at which text will wrap.
func (self *Text) SetWordWrapWidth(member float64) {
    self.Set("wordWrapWidth", member)
}

// Additional spacing (in pixels) between each line of text if multi-line.
func (self *Text) GetLineSpacing() float64{
    return self.Get("lineSpacing").Float()
}

// Additional spacing (in pixels) between each line of text if multi-line.
func (self *Text) SetLineSpacing(member float64) {
    self.Set("lineSpacing", member)
}

// The shadowOffsetX value in pixels. This is how far offset horizontally the shadow effect will be.
func (self *Text) GetShadowOffsetX() float64{
    return self.Get("shadowOffsetX").Float()
}

// The shadowOffsetX value in pixels. This is how far offset horizontally the shadow effect will be.
func (self *Text) SetShadowOffsetX(member float64) {
    self.Set("shadowOffsetX", member)
}

// The shadowOffsetY value in pixels. This is how far offset vertically the shadow effect will be.
func (self *Text) GetShadowOffsetY() float64{
    return self.Get("shadowOffsetY").Float()
}

// The shadowOffsetY value in pixels. This is how far offset vertically the shadow effect will be.
func (self *Text) SetShadowOffsetY(member float64) {
    self.Set("shadowOffsetY", member)
}

// The color of the shadow, as given in CSS rgba format. Set the alpha component to 0 to disable the shadow.
func (self *Text) GetShadowColor() string{
    return self.Get("shadowColor").String()
}

// The color of the shadow, as given in CSS rgba format. Set the alpha component to 0 to disable the shadow.
func (self *Text) SetShadowColor(member string) {
    self.Set("shadowColor", member)
}

// The shadowBlur value. Make the shadow softer by applying a Gaussian blur to it. A number from 0 (no blur) up to approx. 10 (depending on scene).
func (self *Text) GetShadowBlur() float64{
    return self.Get("shadowBlur").Float()
}

// The shadowBlur value. Make the shadow softer by applying a Gaussian blur to it. A number from 0 (no blur) up to approx. 10 (depending on scene).
func (self *Text) SetShadowBlur(member float64) {
    self.Set("shadowBlur", member)
}

// Sets if the drop shadow is applied to the Text stroke.
func (self *Text) GetShadowStroke() bool{
    return self.Get("shadowStroke").Bool()
}

// Sets if the drop shadow is applied to the Text stroke.
func (self *Text) SetShadowStroke(member bool) {
    self.Set("shadowStroke", member)
}

// Sets if the drop shadow is applied to the Text fill.
func (self *Text) GetShadowFill() bool{
    return self.Get("shadowFill").Bool()
}

// Sets if the drop shadow is applied to the Text fill.
func (self *Text) SetShadowFill(member bool) {
    self.Set("shadowFill", member)
}

// The width of the Text. Setting this will modify the scale to achieve the value requested.
func (self *Text) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the Text. Setting this will modify the scale to achieve the value requested.
func (self *Text) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the Text. Setting this will modify the scale to achieve the value requested.
func (self *Text) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the Text. Setting this will modify the scale to achieve the value requested.
func (self *Text) SetHeight(member float64) {
    self.Set("height", member)
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *Text) GetAnchor() *Point{
    return &Point{self.Get("anchor")}
}

// The anchor sets the origin point of the texture.
// The default is 0,0 this means the texture's origin is the top left
// Setting than anchor to 0.5,0.5 means the textures origin is centered
// Setting the anchor to 1,1 would mean the textures origin points will be the bottom right corner
func (self *Text) SetAnchor(member *Point) {
    self.Set("anchor", member)
}

// The texture that the sprite is using
func (self *Text) GetTexture() *Texture{
    return &Texture{self.Get("texture")}
}

// The texture that the sprite is using
func (self *Text) SetTexture(member *Texture) {
    self.Set("texture", member)
}

// The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *Text) GetTint() float64{
    return self.Get("tint").Float()
}

// The tint applied to the sprite. This is a hex value. A value of 0xFFFFFF will remove any tint effect.
func (self *Text) SetTint(member float64) {
    self.Set("tint", member)
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *Text) GetTintedTexture() *Canvas{
    return &Canvas{self.Get("tintedTexture")}
}

// A canvas that contains the tinted version of the Sprite (in Canvas mode, WebGL doesn't populate this)
func (self *Text) SetTintedTexture(member *Canvas) {
    self.Set("tintedTexture", member)
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *Text) GetBlendMode() float64{
    return self.Get("blendMode").Float()
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
// 
// Warning: You cannot have a blend mode and a filter active on the same Sprite. Doing so will render the sprite invisible.
func (self *Text) SetBlendMode(member float64) {
    self.Set("blendMode", member)
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *Text) GetShader() *AbstractFilter{
    return &AbstractFilter{self.Get("shader")}
}

// The shader that will be used to render the texture to the stage. Set to null to remove a current shader.
func (self *Text) SetShader(member *AbstractFilter) {
    self.Set("shader", member)
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *Text) GetExists() bool{
    return self.Get("exists").Bool()
}

// Controls if this Sprite is processed by the core Phaser game loops and Group loops.
func (self *Text) SetExists(member bool) {
    self.Set("exists", member)
}

// [read-only] The array of children of this container.
func (self *Text) GetChildren() []DisplayObject{
	array := self.Get("children")
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// [read-only] The array of children of this container.
func (self *Text) SetChildren(member []DisplayObject) {
    self.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Text) GetIgnoreChildInput() bool{
    return self.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Text) SetIgnoreChildInput(member bool) {
    self.Set("ignoreChildInput", member)
}

// A reference to the currently running Game.
func (self *Text) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Text) SetGame(member *Game) {
    self.Set("game", member)
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Text) GetName() string{
    return self.Get("name").String()
}

// A user defined name given to this Game Object.
// This value isn't ever used internally by Phaser, it is meant as a game level property.
func (self *Text) SetName(member string) {
    self.Set("name", member)
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Text) GetData() interface{}{
    return self.Get("data")
}

// An empty Object that belongs to this Game Object.
// This value isn't ever used internally by Phaser, but may be used by your own code, or
// by Phaser Plugins, to store data that needs to be associated with the Game Object,
// without polluting the Game Object directly.
func (self *Text) SetData(member interface{}) {
    self.Set("data", member)
}

// The components this Game Object has installed.
func (self *Text) GetComponents() interface{}{
    return self.Get("components")
}

// The components this Game Object has installed.
func (self *Text) SetComponents(member interface{}) {
    self.Set("components", member)
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Text) GetZ() float64{
    return self.Get("z").Float()
}

// The z depth of this Game Object within its parent Group.
// No two objects in a Group can have the same z value.
// This value is adjusted automatically whenever the Group hierarchy changes.
// If you wish to re-order the layering of a Game Object then see methods like Group.moveUp or Group.bringToTop.
func (self *Text) SetZ(member float64) {
    self.Set("z", member)
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Text) GetEvents() *Events{
    return &Events{self.Get("events")}
}

// All Phaser Game Objects have an Events class which contains all of the events that are dispatched when certain things happen to this
// Game Object, or any of its components.
func (self *Text) SetEvents(member *Events) {
    self.Set("events", member)
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Text) GetAnimations() *AnimationManager{
    return &AnimationManager{self.Get("animations")}
}

// If the Game Object is enabled for animation (such as a Phaser.Sprite) this is a reference to its AnimationManager instance.
// Through it you can create, play, pause and stop animations.
func (self *Text) SetAnimations(member *AnimationManager) {
    self.Set("animations", member)
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Text) GetKey() interface{}{
    return self.Get("key")
}

// The key of the image or texture used by this Game Object during rendering.
// If it is a string it's the string used to retrieve the texture from the Phaser Image Cache.
// It can also be an instance of a RenderTexture, BitmapData, Video or PIXI.Texture.
// If a Game Object is created without a key it is automatically assigned the key `__default` which is a 32x32 transparent PNG stored within the Cache.
// If a Game Object is given a key which doesn't exist in the Image Cache it is re-assigned the key `__missing` which is a 32x32 PNG of a green box with a line through it.
func (self *Text) SetKey(member interface{}) {
    self.Set("key", member)
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Text) GetWorld() *Point{
    return &Point{self.Get("world")}
}

// The world coordinates of this Game Object in pixels.
// Depending on where in the display list this Game Object is placed this value can differ from `position`, 
// which contains the x/y coordinates relative to the Game Objects parent.
func (self *Text) SetWorld(member *Point) {
    self.Set("world", member)
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Text) GetDebug() bool{
    return self.Get("debug").Bool()
}

// A debug flag designed for use with `Game.enableStep`.
func (self *Text) SetDebug(member bool) {
    self.Set("debug", member)
}

// The position the Game Object was located in the previous frame.
func (self *Text) GetPreviousPosition() *Point{
    return &Point{self.Get("previousPosition")}
}

// The position the Game Object was located in the previous frame.
func (self *Text) SetPreviousPosition(member *Point) {
    self.Set("previousPosition", member)
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Text) GetPreviousRotation() float64{
    return self.Get("previousRotation").Float()
}

// The rotation the Game Object was in set to in the previous frame. Value is in radians.
func (self *Text) SetPreviousRotation(member float64) {
    self.Set("previousRotation", member)
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Text) GetRenderOrderID() float64{
    return self.Get("renderOrderID").Float()
}

// The render order ID is used internally by the renderer and Input Manager and should not be modified.
// This property is mostly used internally by the renderers, but is exposed for the use of plugins.
func (self *Text) SetRenderOrderID(member float64) {
    self.Set("renderOrderID", member)
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Text) GetFresh() bool{
    return self.Get("fresh").Bool()
}

// A Game Object is considered `fresh` if it has just been created or reset and is yet to receive a renderer transform update.
// This property is mostly used internally by the physics systems, but is exposed for the use of plugins.
func (self *Text) SetFresh(member bool) {
    self.Set("fresh", member)
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Text) GetPendingDestroy() bool{
    return self.Get("pendingDestroy").Bool()
}

// A Game Object is that is pendingDestroy is flagged to have its destroy method called on the next logic update.
// You can set it directly to allow you to flag an object to be destroyed on its next update.
// 
// This is extremely useful if you wish to destroy an object from within one of its own callbacks 
// such as with Buttons or other Input events.
func (self *Text) SetPendingDestroy(member bool) {
    self.Set("pendingDestroy", member)
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
func (self *Text) GetAngle() float64{
    return self.Get("angle").Float()
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
func (self *Text) SetAngle(member float64) {
    self.Set("angle", member)
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Text) GetAutoCull() bool{
    return self.Get("autoCull").Bool()
}

// A Game Object with `autoCull` set to true will check its bounds against the World Camera every frame.
// If it is not intersecting the Camera bounds at any point then it has its `renderable` property set to `false`.
// This keeps the Game Object alive and still processing updates, but forces it to skip the render step entirely.
// 
// This is a relatively expensive operation, especially if enabled on hundreds of Game Objects. So enable it only if you know it's required,
// or you have tested performance and find it acceptable.
func (self *Text) SetAutoCull(member bool) {
    self.Set("autoCull", member)
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Text) GetInCamera() bool{
    return self.Get("inCamera").Bool()
}

// Checks if the Game Objects bounds intersect with the Game Camera bounds.
// Returns `true` if they do, otherwise `false` if fully outside of the Cameras bounds.
func (self *Text) SetInCamera(member bool) {
    self.Set("inCamera", member)
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Text) GetOffsetX() float64{
    return self.Get("offsetX").Float()
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *Text) SetOffsetX(member float64) {
    self.Set("offsetX", member)
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Text) GetOffsetY() float64{
    return self.Get("offsetY").Float()
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *Text) SetOffsetY(member float64) {
    self.Set("offsetY", member)
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Text) GetCenterX() float64{
    return self.Get("centerX").Float()
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *Text) SetCenterX(member float64) {
    self.Set("centerX", member)
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Text) GetCenterY() float64{
    return self.Get("centerY").Float()
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *Text) SetCenterY(member float64) {
    self.Set("centerY", member)
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Text) GetLeft() float64{
    return self.Get("left").Float()
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *Text) SetLeft(member float64) {
    self.Set("left", member)
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Text) GetRight() float64{
    return self.Get("right").Float()
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *Text) SetRight(member float64) {
    self.Set("right", member)
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Text) GetTop() float64{
    return self.Get("top").Float()
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *Text) SetTop(member float64) {
    self.Set("top", member)
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Text) GetBottom() float64{
    return self.Get("bottom").Float()
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *Text) SetBottom(member float64) {
    self.Set("bottom", member)
}

// The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Text) GetCropRect() *Rectangle{
    return &Rectangle{self.Get("cropRect")}
}

// The Rectangle used to crop the texture this Game Object uses.
// Set this property via `crop`. 
// If you modify this property directly you must call `updateCrop` in order to have the change take effect.
func (self *Text) SetCropRect(member *Rectangle) {
    self.Set("cropRect", member)
}

// Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *Text) GetDeltaX() float64{
    return self.Get("deltaX").Float()
}

// Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *Text) SetDeltaX(member float64) {
    self.Set("deltaX", member)
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *Text) GetDeltaY() float64{
    return self.Get("deltaY").Float()
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *Text) SetDeltaY(member float64) {
    self.Set("deltaY", member)
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *Text) GetDeltaZ() float64{
    return self.Get("deltaZ").Float()
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *Text) SetDeltaZ(member float64) {
    self.Set("deltaZ", member)
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Text) GetDestroyPhase() bool{
    return self.Get("destroyPhase").Bool()
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *Text) SetDestroyPhase(member bool) {
    self.Set("destroyPhase", member)
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
func (self *Text) GetFixedToCamera() bool{
    return self.Get("fixedToCamera").Bool()
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
func (self *Text) SetFixedToCamera(member bool) {
    self.Set("fixedToCamera", member)
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Text) GetCameraOffset() *Point{
    return &Point{self.Get("cameraOffset")}
}

// The x/y coordinate offset applied to the top-left of the camera that this Game Object will be drawn at if `fixedToCamera` is true.
// 
// The values are relative to the top-left of the camera view and in addition to any parent of the Game Object on the display list.
func (self *Text) SetCameraOffset(member *Point) {
    self.Set("cameraOffset", member)
}

// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *Text) GetHealth() float64{
    return self.Get("health").Float()
}

// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *Text) SetHealth(member float64) {
    self.Set("health", member)
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *Text) GetMaxHealth() float64{
    return self.Get("maxHealth").Float()
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *Text) SetMaxHealth(member float64) {
    self.Set("maxHealth", member)
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *Text) GetDamage() interface{}{
    return self.Get("damage")
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *Text) SetDamage(member interface{}) {
    self.Set("damage", member)
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *Text) GetSetHealth() interface{}{
    return self.Get("setHealth")
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *Text) SetSetHealth(member interface{}) {
    self.Set("setHealth", member)
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *Text) GetHeal() interface{}{
    return self.Get("heal")
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *Text) SetHeal(member interface{}) {
    self.Set("heal", member)
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Text) GetInput() interface{}{
    return self.Get("input")
}

// The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *Text) SetInput(member interface{}) {
    self.Set("input", member)
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
func (self *Text) GetInputEnabled() bool{
    return self.Get("inputEnabled").Bool()
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
func (self *Text) SetInputEnabled(member bool) {
    self.Set("inputEnabled", member)
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
func (self *Text) GetCheckWorldBounds() bool{
    return self.Get("checkWorldBounds").Bool()
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
func (self *Text) SetCheckWorldBounds(member bool) {
    self.Set("checkWorldBounds", member)
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *Text) GetOutOfBoundsKill() bool{
    return self.Get("outOfBoundsKill").Bool()
}

// If this and the `checkWorldBounds` property are both set to `true` then the `kill` method is called as soon as `inWorld` returns false.
func (self *Text) SetOutOfBoundsKill(member bool) {
    self.Set("outOfBoundsKill", member)
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *Text) GetOutOfCameraBoundsKill() bool{
    return self.Get("outOfCameraBoundsKill").Bool()
}

// If this and the `autoCull` property are both set to `true`, then the `kill` method
// is called as soon as the Game Object leaves the camera bounds.
func (self *Text) SetOutOfCameraBoundsKill(member bool) {
    self.Set("outOfCameraBoundsKill", member)
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *Text) GetInWorld() bool{
    return self.Get("inWorld").Bool()
}

// Checks if the Game Objects bounds are within, or intersect at any point with the Game World bounds.
func (self *Text) SetInWorld(member bool) {
    self.Set("inWorld", member)
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Text) GetAlive() bool{
    return self.Get("alive").Bool()
}

// A useful flag to control if the Game Object is alive or dead.
// 
// This is set automatically by the Health components `damage` method should the object run out of health.
// Or you can toggle it via your game code.
// 
// This property is mostly just provided to be used by your game - it doesn't effect rendering or logic updates.
// However you can use `Group.getFirstAlive` in conjunction with this property for fast object pooling and recycling.
func (self *Text) SetAlive(member bool) {
    self.Set("alive", member)
}

// The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *Text) GetLifespan() float64{
    return self.Get("lifespan").Float()
}

// The lifespan allows you to give a Game Object a lifespan in milliseconds.
// 
// Once the Game Object is 'born' you can set this to a positive value.
// 
// It is automatically decremented by the millisecond equivalent of `game.time.physicsElapsed` each frame.
// When it reaches zero it will call the `kill` method.
// 
// Very handy for particles, bullets, collectibles, or any other short-lived entity.
func (self *Text) SetLifespan(member float64) {
    self.Set("lifespan", member)
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
func (self *Text) GetFrame() int{
    return self.Get("frame").Int()
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
func (self *Text) SetFrame(member int) {
    self.Set("frame", member)
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
func (self *Text) GetFrameName() string{
    return self.Get("frameName").String()
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
func (self *Text) SetFrameName(member string) {
    self.Set("frameName", member)
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
func (self *Text) GetBody() interface{}{
    return self.Get("body")
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
func (self *Text) SetBody(member interface{}) {
    self.Set("body", member)
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *Text) GetX() float64{
    return self.Get("x").Float()
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *Text) SetX(member float64) {
    self.Set("x", member)
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *Text) GetY() float64{
    return self.Get("y").Float()
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *Text) SetY(member float64) {
    self.Set("y", member)
}

// The callback that will apply any scale limiting to the worldTransform.
func (self *Text) SetTransformCallback(member func(...interface{})) {
    self.Set("transformCallback", member)
}

// The context under which `transformCallback` is called.
func (self *Text) GetTransformCallbackContext() interface{}{
    return self.Get("transformCallbackContext")
}

// The context under which `transformCallback` is called.
func (self *Text) SetTransformCallbackContext(member interface{}) {
    self.Set("transformCallbackContext", member)
}

// The minimum scale this Game Object will scale down to.
// 
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
// 
// Set it to `null` to remove the limit.
func (self *Text) GetScaleMin() *Point{
    return &Point{self.Get("scaleMin")}
}

// The minimum scale this Game Object will scale down to.
// 
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
// 
// Set it to `null` to remove the limit.
func (self *Text) SetScaleMin(member *Point) {
    self.Set("scaleMin", member)
}

// The maximum scale this Game Object will scale up to. 
// 
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
// 
// Set it to `null` to remove the limit.
func (self *Text) GetScaleMax() *Point{
    return &Point{self.Get("scaleMax")}
}

// The maximum scale this Game Object will scale up to. 
// 
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
// 
// Set it to `null` to remove the limit.
func (self *Text) SetScaleMax(member *Point) {
    self.Set("scaleMax", member)
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *Text) GetSmoothed() bool{
    return self.Get("smoothed").Bool()
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *Text) SetSmoothed(member bool) {
    self.Set("smoothed", member)
}



// Automatically called by World.preUpdate.
func (self *Text) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Override this function to handle any special update requirements.
func (self *Text) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Destroy this Text object, removing it from the group it belongs to.
func (self *Text) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Sets a drop shadow effect on the Text. You can specify the horizontal and vertical distance of the drop shadow with the `x` and `y` parameters.
// The color controls the shade of the shadow (default is black) and can be either an `rgba` or `hex` value.
// The blur is the strength of the shadow. A value of zero means a hard shadow, a value of 10 means a very soft shadow.
// To remove a shadow already in place you can call this method with no parameters set.
func (self *Text) SetShadowI(args ...interface{}) *Text{
    return &Text{self.Call("setShadow", args)}
}

// Set the style of the text by passing a single style object to it.
func (self *Text) SetStyleI(args ...interface{}) *Text{
    return &Text{self.Call("setStyle", args)}
}

// Renders text. This replaces the Pixi.Text.updateText function as we need a few extra bits in here.
func (self *Text) UpdateTextI(args ...interface{}) {
    self.Call("updateText", args)
}

// Renders a line of text that contains tab characters if Text.tab > 0.
// Called automatically by updateText.
func (self *Text) RenderTabLineI(args ...interface{}) {
    self.Call("renderTabLine", args)
}

// Sets the Shadow on the Text.context based on the Style settings, or disables it if not enabled.
// This is called automatically by Text.updateText.
func (self *Text) UpdateShadowI(args ...interface{}) {
    self.Call("updateShadow", args)
}

// Measures a line of text character by character taking into the account the specified character styles.
func (self *Text) MeasureLineI(args ...interface{}) int{
    return self.Call("measureLine", args).Int()
}

// Updates a line of text, applying fill and stroke per-character colors or style and weight per-character font if applicable.
func (self *Text) UpdateLineI(args ...interface{}) {
    self.Call("updateLine", args)
}

// Clears any text fill or stroke colors that were set by `addColor` or `addStrokeColor`.
func (self *Text) ClearColorsI(args ...interface{}) *Text{
    return &Text{self.Call("clearColors", args)}
}

// Clears any text styles or weights font that were set by `addFontStyle` or `addFontWeight`.
func (self *Text) ClearFontValuesI(args ...interface{}) *Text{
    return &Text{self.Call("clearFontValues", args)}
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
    return &Text{self.Call("addColor", args)}
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
    return &Text{self.Call("addStrokeColor", args)}
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
    return &Text{self.Call("addFontStyle", args)}
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
    return &Text{self.Call("addFontWeight", args)}
}

// Runs the given text through the Text.runWordWrap function and returns
// the results as an array, where each element of the array corresponds to a wrapped
// line of text.
// 
// Useful if you wish to control pagination on long pieces of content.
func (self *Text) PrecalculateWordWrapI(args ...interface{}) []interface{}{
	array := self.Call("precalculateWordWrap", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Greedy wrapping algorithm that will wrap words as the line grows longer than its horizontal bounds.
func (self *Text) RunWordWrapI(args ...interface{}) {
    self.Call("runWordWrap", args)
}

// Advanced wrapping algorithm that will wrap words as the line grows longer than its horizontal bounds.
// White space is condensed (e.g., consecutive spaces are replaced with one).
// Lines are trimmed of white space before processing.
// Throws an error if the user was smart enough to specify a wordWrapWidth less than a single character.
func (self *Text) AdvancedWordWrapI(args ...interface{}) {
    self.Call("advancedWordWrap", args)
}

// Greedy wrapping algorithm that will wrap words as the line grows longer than its horizontal bounds.
func (self *Text) BasicWordWrapI(args ...interface{}) {
    self.Call("basicWordWrap", args)
}

// Updates the internal `style.font` if it now differs according to generation from components.
func (self *Text) UpdateFontI(args ...interface{}) {
    self.Call("updateFont", args)
}

// Converting a short CSS-font string into the relevant components.
func (self *Text) FontToComponentsI(args ...interface{}) {
    self.Call("fontToComponents", args)
}

// Converts individual font components (see `fontToComponents`) to a short CSS font string.
func (self *Text) ComponentsToFontI(args ...interface{}) {
    self.Call("componentsToFont", args)
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
    return &Text{self.Call("setText", args)}
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
    return &Text{self.Call("parseList", args)}
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
    return &Text{self.Call("setTextBounds", args)}
}

// Updates the texture based on the canvas dimensions.
func (self *Text) UpdateTextureI(args ...interface{}) {
    self.Call("updateTexture", args)
}

// Renders the object using the WebGL renderer
func (self *Text) _renderWebGLI(args ...interface{}) {
    self.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer.
func (self *Text) _renderCanvasI(args ...interface{}) {
    self.Call("_renderCanvas", args)
}

// Calculates the ascent, descent and fontSize of a given font style.
func (self *Text) DetermineFontPropertiesI(args ...interface{}) {
    self.Call("determineFontProperties", args)
}

// Returns the bounds of the Text as a rectangle.
// The bounds calculation takes the worldTransform into account.
func (self *Text) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getBounds", args)}
}

// Sets the texture of the sprite. Be warned that this doesn't remove or destroy the previous
// texture this Sprite was using.
func (self *Text) SetTextureI(args ...interface{}) {
    self.Call("setTexture", args)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *Text) OnTextureUpdateI(args ...interface{}) {
    self.Call("onTextureUpdate", args)
}

// Adds a child to the container.
func (self *Text) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Text) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *Text) SwapChildrenI(args ...interface{}) {
    self.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *Text) GetChildIndexI(args ...interface{}) float64{
    return self.Call("getChildIndex", args).Float()
}

// Changes the position of an existing child in the display object container
func (self *Text) SetChildIndexI(args ...interface{}) {
    self.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *Text) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *Text) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *Text) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Text) RemoveChildrenI(args ...interface{}) {
    self.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Text) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Text) SetStageReferenceI(args ...interface{}) {
    self.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *Text) RemoveStageReferenceI(args ...interface{}) {
    self.Call("removeStageReference", args)
}

// Internal method called by the World postUpdate cycle.
func (self *Text) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *Text) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Call("play", args)}
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
    return self.Call("alignIn", args)
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
    return self.Call("alignTo", args)
}

// Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) BringToTopI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("bringToTop", args)}
}

// Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) SendToBackI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("sendToBack", args)}
}

// Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) MoveUpI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("moveUp", args)}
}

// Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *Text) MoveDownI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("moveDown", args)}
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
    self.Call("crop", args)
}

// If you have set a crop rectangle on this Game Object via `crop` and since modified the `cropRect` property,
// or the rectangle it references, then you need to update the crop frame by calling this method.
func (self *Text) UpdateCropI(args ...interface{}) {
    self.Call("updateCrop", args)
}

// Brings a 'dead' Game Object back to life, optionally resetting its health value in the process.
// 
// A resurrected Game Object has its `alive`, `exists` and `visible` properties all set to true.
// 
// It will dispatch the `onRevived` event. Listen to `events.onRevived` for the signal.
func (self *Text) ReviveI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("revive", args)}
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
    return &DisplayObject{self.Call("kill", args)}
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
    self.Call("loadTexture", args)
}

// Sets the texture frame the Game Object uses for rendering.
// 
// This is primarily an internal method used by `loadTexture`, but is exposed for the use of plugins and custom classes.
func (self *Text) SetFrameI(args ...interface{}) {
    self.Call("setFrame", args)
}

// Resizes the Frame dimensions that the Game Object uses for rendering.
// 
// You shouldn't normally need to ever call this, but in the case of special texture types such as Video or BitmapData
// it can be useful to adjust the dimensions directly in this way.
func (self *Text) ResizeFrameI(args ...interface{}) {
    self.Call("resizeFrame", args)
}

// Resets the texture frame dimensions that the Game Object uses for rendering.
func (self *Text) ResetFrameI(args ...interface{}) {
    self.Call("resetFrame", args)
}

// Checks to see if the bounds of this Game Object overlaps with the bounds of the given Display Object, 
// which can be a Sprite, Image, TileSprite or anything that extends those such as Button or provides a `getBounds` method and result.
// 
// This check ignores the `hitArea` property if set and runs a `getBounds` comparison on both objects to determine the result.
// 
// Therefore it's relatively expensive to use in large quantities, i.e. with lots of Sprites at a high frequency.
// It should be fine for low-volume testing where physics isn't required.
func (self *Text) OverlapI(args ...interface{}) bool{
    return self.Call("overlap", args).Bool()
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
    return &DisplayObject{self.Call("reset", args)}
}

// Adjust scaling limits, if set, to this Game Object.
func (self *Text) CheckTransformI(args ...interface{}) {
    self.Call("checkTransform", args)
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
    self.Call("setScaleMinMax", args)
}
