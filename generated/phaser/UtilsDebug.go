// Automatic generation for Phaser.Utils.Debug
// generated file UtilsDebug.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// A collection of methods for displaying debug information about game objects.
// 
// If your game is running in Canvas mode, then you should invoke all of the Debug methods from
// your games `render` function. This is because they are drawn directly onto the game canvas
// itself, so if you call any debug methods outside of `render` they are likely to be overwritten
// by the game itself.
// 
// If your game is running in WebGL then Debug will create a Sprite that is placed at the top of the Stage display list and bind a canvas texture
// to it, which must be uploaded every frame. Be advised: this is very expensive, especially in browsers like Firefox. So please only enable Debug
// in WebGL mode if you really need it (or your desktop can cope with it well) and disable it for production!
type UtilsDebug struct {
    *js.Object
}


// A collection of methods for displaying debug information about game objects.
// 
// If your game is running in Canvas mode, then you should invoke all of the Debug methods from
// your games `render` function. This is because they are drawn directly onto the game canvas
// itself, so if you call any debug methods outside of `render` they are likely to be overwritten
// by the game itself.
// 
// If your game is running in WebGL then Debug will create a Sprite that is placed at the top of the Stage display list and bind a canvas texture
// to it, which must be uploaded every frame. Be advised: this is very expensive, especially in browsers like Firefox. So please only enable Debug
// in WebGL mode if you really need it (or your desktop can cope with it well) and disable it for production!
func NewUtilsDebug(game *Game) *UtilsDebug {
    return &UtilsDebug{js.Global.Get("Phaser").Get("Utils").Get("Debug").New(game)}
}

// A collection of methods for displaying debug information about game objects.
// 
// If your game is running in Canvas mode, then you should invoke all of the Debug methods from
// your games `render` function. This is because they are drawn directly onto the game canvas
// itself, so if you call any debug methods outside of `render` they are likely to be overwritten
// by the game itself.
// 
// If your game is running in WebGL then Debug will create a Sprite that is placed at the top of the Stage display list and bind a canvas texture
// to it, which must be uploaded every frame. Be advised: this is very expensive, especially in browsers like Firefox. So please only enable Debug
// in WebGL mode if you really need it (or your desktop can cope with it well) and disable it for production!
func NewUtilsDebugI(args ...interface{}) *UtilsDebug {
    return &UtilsDebug{js.Global.Get("Phaser").Get("Utils").Get("Debug").New(args)}
}



// A reference to the currently running Game.
func (self *UtilsDebug) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *UtilsDebug) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// If debugging in WebGL mode we need this.
func (self *UtilsDebug) Sprite() *Image{
    return &Image{self.Object.Get("sprite")}
}

// If debugging in WebGL mode we need this.
func (self *UtilsDebug) SetSpriteA(member *Image) {
    self.Object.Set("sprite", member)
}

// In WebGL mode this BitmapData contains a copy of the debug canvas.
func (self *UtilsDebug) Bmd() *BitmapData{
    return &BitmapData{self.Object.Get("bmd")}
}

// In WebGL mode this BitmapData contains a copy of the debug canvas.
func (self *UtilsDebug) SetBmdA(member *BitmapData) {
    self.Object.Set("bmd", member)
}

// The canvas to which Debug calls draws.
func (self *UtilsDebug) Canvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("canvas"))
}

// The canvas to which Debug calls draws.
func (self *UtilsDebug) SetCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("canvas", member)
}

// The 2d context of the canvas.
func (self *UtilsDebug) Context() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("context"))
}

// The 2d context of the canvas.
func (self *UtilsDebug) SetContextA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("context", member)
}

// The font that the debug information is rendered in.
func (self *UtilsDebug) Font() string{
    return self.Object.Get("font").String()
}

// The font that the debug information is rendered in.
func (self *UtilsDebug) SetFontA(member string) {
    self.Object.Set("font", member)
}

// The spacing between columns.
func (self *UtilsDebug) ColumnWidth() int{
    return self.Object.Get("columnWidth").Int()
}

// The spacing between columns.
func (self *UtilsDebug) SetColumnWidthA(member int) {
    self.Object.Set("columnWidth", member)
}

// The line height between the debug text.
func (self *UtilsDebug) LineHeight() int{
    return self.Object.Get("lineHeight").Int()
}

// The line height between the debug text.
func (self *UtilsDebug) SetLineHeightA(member int) {
    self.Object.Set("lineHeight", member)
}

// Should the text be rendered with a slight shadow? Makes it easier to read on different types of background.
func (self *UtilsDebug) RenderShadow() bool{
    return self.Object.Get("renderShadow").Bool()
}

// Should the text be rendered with a slight shadow? Makes it easier to read on different types of background.
func (self *UtilsDebug) SetRenderShadowA(member bool) {
    self.Object.Set("renderShadow", member)
}

// The current X position the debug information will be rendered at.
func (self *UtilsDebug) CurrentX() int{
    return self.Object.Get("currentX").Int()
}

// The current X position the debug information will be rendered at.
func (self *UtilsDebug) SetCurrentXA(member int) {
    self.Object.Set("currentX", member)
}

// The current Y position the debug information will be rendered at.
func (self *UtilsDebug) CurrentY() int{
    return self.Object.Get("currentY").Int()
}

// The current Y position the debug information will be rendered at.
func (self *UtilsDebug) SetCurrentYA(member int) {
    self.Object.Set("currentY", member)
}

// The alpha of the Debug context, set before all debug information is rendered to it.
func (self *UtilsDebug) CurrentAlpha() int{
    return self.Object.Get("currentAlpha").Int()
}

// The alpha of the Debug context, set before all debug information is rendered to it.
func (self *UtilsDebug) SetCurrentAlphaA(member int) {
    self.Object.Set("currentAlpha", member)
}

// Does the canvas need re-rendering?
func (self *UtilsDebug) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// Does the canvas need re-rendering?
func (self *UtilsDebug) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}



// Internal method that boots the debug displayer.
func (self *UtilsDebug) Boot() {
    self.Object.Call("boot")
}

// Internal method that boots the debug displayer.
func (self *UtilsDebug) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// Internal method that resizes the BitmapData and Canvas.
// Called by ScaleManager.onSizeChange only in WebGL mode.
func (self *UtilsDebug) Resize(scaleManager *ScaleManager, width int, height int) {
    self.Object.Call("resize", scaleManager, width, height)
}

// Internal method that resizes the BitmapData and Canvas.
// Called by ScaleManager.onSizeChange only in WebGL mode.
func (self *UtilsDebug) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Internal method that clears the canvas (if a Sprite) ready for a new debug session.
func (self *UtilsDebug) PreUpdate() {
    self.Object.Call("preUpdate")
}

// Internal method that clears the canvas (if a Sprite) ready for a new debug session.
func (self *UtilsDebug) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Clears the Debug canvas.
func (self *UtilsDebug) Reset() {
    self.Object.Call("reset")
}

// Clears the Debug canvas.
func (self *UtilsDebug) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Internal method that resets and starts the debug output values.
func (self *UtilsDebug) Start() {
    self.Object.Call("start")
}

// Internal method that resets and starts the debug output values.
func (self *UtilsDebug) Start1O(x int) {
    self.Object.Call("start", x)
}

// Internal method that resets and starts the debug output values.
func (self *UtilsDebug) Start2O(x int, y int) {
    self.Object.Call("start", x, y)
}

// Internal method that resets and starts the debug output values.
func (self *UtilsDebug) Start3O(x int, y int, color string) {
    self.Object.Call("start", x, y, color)
}

// Internal method that resets and starts the debug output values.
func (self *UtilsDebug) Start4O(x int, y int, color string, columnWidth int) {
    self.Object.Call("start", x, y, color, columnWidth)
}

// Internal method that resets and starts the debug output values.
func (self *UtilsDebug) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Internal method that stops the debug output.
func (self *UtilsDebug) Stop() {
    self.Object.Call("stop")
}

// Internal method that stops the debug output.
func (self *UtilsDebug) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Internal method that outputs a single line of text split over as many columns as needed, one per parameter.
func (self *UtilsDebug) Line() {
    self.Object.Call("line")
}

// Internal method that outputs a single line of text split over as many columns as needed, one per parameter.
func (self *UtilsDebug) LineI(args ...interface{}) {
    self.Object.Call("line", args)
}

// Render Sound information, including decoded state, duration, volume and more.
func (self *UtilsDebug) SoundInfo(sound *Sound, x int, y int) {
    self.Object.Call("soundInfo", sound, x, y)
}

// Render Sound information, including decoded state, duration, volume and more.
func (self *UtilsDebug) SoundInfo1O(sound *Sound, x int, y int, color string) {
    self.Object.Call("soundInfo", sound, x, y, color)
}

// Render Sound information, including decoded state, duration, volume and more.
func (self *UtilsDebug) SoundInfoI(args ...interface{}) {
    self.Object.Call("soundInfo", args)
}

// Render camera information including dimensions and location.
func (self *UtilsDebug) CameraInfo(camera *Camera, x int, y int) {
    self.Object.Call("cameraInfo", camera, x, y)
}

// Render camera information including dimensions and location.
func (self *UtilsDebug) CameraInfo1O(camera *Camera, x int, y int, color string) {
    self.Object.Call("cameraInfo", camera, x, y, color)
}

// Render camera information including dimensions and location.
func (self *UtilsDebug) CameraInfoI(args ...interface{}) {
    self.Object.Call("cameraInfo", args)
}

// Render Timer information.
func (self *UtilsDebug) Timer(timer *Timer, x int, y int) {
    self.Object.Call("timer", timer, x, y)
}

// Render Timer information.
func (self *UtilsDebug) Timer1O(timer *Timer, x int, y int, color string) {
    self.Object.Call("timer", timer, x, y, color)
}

// Render Timer information.
func (self *UtilsDebug) TimerI(args ...interface{}) {
    self.Object.Call("timer", args)
}

// Renders the Pointer.circle object onto the stage in green if down or red if up along with debug text.
func (self *UtilsDebug) Pointer(pointer *Pointer) {
    self.Object.Call("pointer", pointer)
}

// Renders the Pointer.circle object onto the stage in green if down or red if up along with debug text.
func (self *UtilsDebug) Pointer1O(pointer *Pointer, hideIfUp bool) {
    self.Object.Call("pointer", pointer, hideIfUp)
}

// Renders the Pointer.circle object onto the stage in green if down or red if up along with debug text.
func (self *UtilsDebug) Pointer2O(pointer *Pointer, hideIfUp bool, downColor string) {
    self.Object.Call("pointer", pointer, hideIfUp, downColor)
}

// Renders the Pointer.circle object onto the stage in green if down or red if up along with debug text.
func (self *UtilsDebug) Pointer3O(pointer *Pointer, hideIfUp bool, downColor string, upColor string) {
    self.Object.Call("pointer", pointer, hideIfUp, downColor, upColor)
}

// Renders the Pointer.circle object onto the stage in green if down or red if up along with debug text.
func (self *UtilsDebug) Pointer4O(pointer *Pointer, hideIfUp bool, downColor string, upColor string, color string) {
    self.Object.Call("pointer", pointer, hideIfUp, downColor, upColor, color)
}

// Renders the Pointer.circle object onto the stage in green if down or red if up along with debug text.
func (self *UtilsDebug) PointerI(args ...interface{}) {
    self.Object.Call("pointer", args)
}

// Render Sprite Input Debug information.
func (self *UtilsDebug) SpriteInputInfo(sprite interface{}, x int, y int) {
    self.Object.Call("spriteInputInfo", sprite, x, y)
}

// Render Sprite Input Debug information.
func (self *UtilsDebug) SpriteInputInfo1O(sprite interface{}, x int, y int, color string) {
    self.Object.Call("spriteInputInfo", sprite, x, y, color)
}

// Render Sprite Input Debug information.
func (self *UtilsDebug) SpriteInputInfoI(args ...interface{}) {
    self.Object.Call("spriteInputInfo", args)
}

// Renders Phaser.Key object information.
func (self *UtilsDebug) Key(key *Key, x int, y int) {
    self.Object.Call("key", key, x, y)
}

// Renders Phaser.Key object information.
func (self *UtilsDebug) Key1O(key *Key, x int, y int, color string) {
    self.Object.Call("key", key, x, y, color)
}

// Renders Phaser.Key object information.
func (self *UtilsDebug) KeyI(args ...interface{}) {
    self.Object.Call("key", args)
}

// Render debug information about the Input object.
func (self *UtilsDebug) InputInfo(x int, y int) {
    self.Object.Call("inputInfo", x, y)
}

// Render debug information about the Input object.
func (self *UtilsDebug) InputInfo1O(x int, y int, color string) {
    self.Object.Call("inputInfo", x, y, color)
}

// Render debug information about the Input object.
func (self *UtilsDebug) InputInfoI(args ...interface{}) {
    self.Object.Call("inputInfo", args)
}

// Renders the Sprites bounds. Note: This is really expensive as it has to calculate the bounds every time you call it!
func (self *UtilsDebug) SpriteBounds(sprite interface{}) {
    self.Object.Call("spriteBounds", sprite)
}

// Renders the Sprites bounds. Note: This is really expensive as it has to calculate the bounds every time you call it!
func (self *UtilsDebug) SpriteBounds1O(sprite interface{}, color string) {
    self.Object.Call("spriteBounds", sprite, color)
}

// Renders the Sprites bounds. Note: This is really expensive as it has to calculate the bounds every time you call it!
func (self *UtilsDebug) SpriteBounds2O(sprite interface{}, color string, filled bool) {
    self.Object.Call("spriteBounds", sprite, color, filled)
}

// Renders the Sprites bounds. Note: This is really expensive as it has to calculate the bounds every time you call it!
func (self *UtilsDebug) SpriteBoundsI(args ...interface{}) {
    self.Object.Call("spriteBounds", args)
}

// Renders the Rope's segments. Note: This is really expensive as it has to calculate new segments every time you call it
func (self *UtilsDebug) RopeSegments(rope *Rope) {
    self.Object.Call("ropeSegments", rope)
}

// Renders the Rope's segments. Note: This is really expensive as it has to calculate new segments every time you call it
func (self *UtilsDebug) RopeSegments1O(rope *Rope, color string) {
    self.Object.Call("ropeSegments", rope, color)
}

// Renders the Rope's segments. Note: This is really expensive as it has to calculate new segments every time you call it
func (self *UtilsDebug) RopeSegments2O(rope *Rope, color string, filled bool) {
    self.Object.Call("ropeSegments", rope, color, filled)
}

// Renders the Rope's segments. Note: This is really expensive as it has to calculate new segments every time you call it
func (self *UtilsDebug) RopeSegmentsI(args ...interface{}) {
    self.Object.Call("ropeSegments", args)
}

// Render debug infos (including name, bounds info, position and some other properties) about the Sprite.
func (self *UtilsDebug) SpriteInfo(sprite *Sprite, x int, y int) {
    self.Object.Call("spriteInfo", sprite, x, y)
}

// Render debug infos (including name, bounds info, position and some other properties) about the Sprite.
func (self *UtilsDebug) SpriteInfo1O(sprite *Sprite, x int, y int, color string) {
    self.Object.Call("spriteInfo", sprite, x, y, color)
}

// Render debug infos (including name, bounds info, position and some other properties) about the Sprite.
func (self *UtilsDebug) SpriteInfoI(args ...interface{}) {
    self.Object.Call("spriteInfo", args)
}

// Renders the sprite coordinates in local, positional and world space.
func (self *UtilsDebug) SpriteCoords(sprite interface{}, x int, y int) {
    self.Object.Call("spriteCoords", sprite, x, y)
}

// Renders the sprite coordinates in local, positional and world space.
func (self *UtilsDebug) SpriteCoords1O(sprite interface{}, x int, y int, color string) {
    self.Object.Call("spriteCoords", sprite, x, y, color)
}

// Renders the sprite coordinates in local, positional and world space.
func (self *UtilsDebug) SpriteCoordsI(args ...interface{}) {
    self.Object.Call("spriteCoords", args)
}

// Renders Line information in the given color.
func (self *UtilsDebug) LineInfo(line *Line, x int, y int) {
    self.Object.Call("lineInfo", line, x, y)
}

// Renders Line information in the given color.
func (self *UtilsDebug) LineInfo1O(line *Line, x int, y int, color string) {
    self.Object.Call("lineInfo", line, x, y, color)
}

// Renders Line information in the given color.
func (self *UtilsDebug) LineInfoI(args ...interface{}) {
    self.Object.Call("lineInfo", args)
}

// Renders a single pixel at the given size.
func (self *UtilsDebug) Pixel(x int, y int) {
    self.Object.Call("pixel", x, y)
}

// Renders a single pixel at the given size.
func (self *UtilsDebug) Pixel1O(x int, y int, color string) {
    self.Object.Call("pixel", x, y, color)
}

// Renders a single pixel at the given size.
func (self *UtilsDebug) Pixel2O(x int, y int, color string, size int) {
    self.Object.Call("pixel", x, y, color, size)
}

// Renders a single pixel at the given size.
func (self *UtilsDebug) PixelI(args ...interface{}) {
    self.Object.Call("pixel", args)
}

// Renders a Phaser geometry object including Rectangle, Circle, Point or Line.
func (self *UtilsDebug) Geom(object interface{}) {
    self.Object.Call("geom", object)
}

// Renders a Phaser geometry object including Rectangle, Circle, Point or Line.
func (self *UtilsDebug) Geom1O(object interface{}, color string) {
    self.Object.Call("geom", object, color)
}

// Renders a Phaser geometry object including Rectangle, Circle, Point or Line.
func (self *UtilsDebug) Geom2O(object interface{}, color string, filled bool) {
    self.Object.Call("geom", object, color, filled)
}

// Renders a Phaser geometry object including Rectangle, Circle, Point or Line.
func (self *UtilsDebug) Geom3O(object interface{}, color string, filled bool, forceType int) {
    self.Object.Call("geom", object, color, filled, forceType)
}

// Renders a Phaser geometry object including Rectangle, Circle, Point or Line.
func (self *UtilsDebug) GeomI(args ...interface{}) {
    self.Object.Call("geom", args)
}

// Render a string of text.
func (self *UtilsDebug) Text(text string, x int, y int) {
    self.Object.Call("text", text, x, y)
}

// Render a string of text.
func (self *UtilsDebug) Text1O(text string, x int, y int, color string) {
    self.Object.Call("text", text, x, y, color)
}

// Render a string of text.
func (self *UtilsDebug) Text2O(text string, x int, y int, color string, font string) {
    self.Object.Call("text", text, x, y, color, font)
}

// Render a string of text.
func (self *UtilsDebug) TextI(args ...interface{}) {
    self.Object.Call("text", args)
}

// Visually renders a QuadTree to the display.
func (self *UtilsDebug) QuadTree(quadtree *QuadTree, color string) {
    self.Object.Call("quadTree", quadtree, color)
}

// Visually renders a QuadTree to the display.
func (self *UtilsDebug) QuadTreeI(args ...interface{}) {
    self.Object.Call("quadTree", args)
}

// Render a Sprites Physics body if it has one set. The body is rendered as a filled or stroked rectangle.
// This only works for Arcade Physics, Ninja Physics (AABB and Circle only) and Box2D Physics bodies.
// To display a P2 Physics body you should enable debug mode on the body when creating it.
func (self *UtilsDebug) Body(sprite *Sprite) {
    self.Object.Call("body", sprite)
}

// Render a Sprites Physics body if it has one set. The body is rendered as a filled or stroked rectangle.
// This only works for Arcade Physics, Ninja Physics (AABB and Circle only) and Box2D Physics bodies.
// To display a P2 Physics body you should enable debug mode on the body when creating it.
func (self *UtilsDebug) Body1O(sprite *Sprite, color string) {
    self.Object.Call("body", sprite, color)
}

// Render a Sprites Physics body if it has one set. The body is rendered as a filled or stroked rectangle.
// This only works for Arcade Physics, Ninja Physics (AABB and Circle only) and Box2D Physics bodies.
// To display a P2 Physics body you should enable debug mode on the body when creating it.
func (self *UtilsDebug) Body2O(sprite *Sprite, color string, filled bool) {
    self.Object.Call("body", sprite, color, filled)
}

// Render a Sprites Physics body if it has one set. The body is rendered as a filled or stroked rectangle.
// This only works for Arcade Physics, Ninja Physics (AABB and Circle only) and Box2D Physics bodies.
// To display a P2 Physics body you should enable debug mode on the body when creating it.
func (self *UtilsDebug) BodyI(args ...interface{}) {
    self.Object.Call("body", args)
}

// Render a Sprites Physic Body information.
func (self *UtilsDebug) BodyInfo(sprite *Sprite, x int, y int) {
    self.Object.Call("bodyInfo", sprite, x, y)
}

// Render a Sprites Physic Body information.
func (self *UtilsDebug) BodyInfo1O(sprite *Sprite, x int, y int, color string) {
    self.Object.Call("bodyInfo", sprite, x, y, color)
}

// Render a Sprites Physic Body information.
func (self *UtilsDebug) BodyInfoI(args ...interface{}) {
    self.Object.Call("bodyInfo", args)
}

// Renders 'debug draw' data for the Box2D world if it exists.
// This uses the standard debug drawing feature of Box2D, so colors will be decided by
// the Box2D engine.
func (self *UtilsDebug) Box2dWorld() {
    self.Object.Call("box2dWorld")
}

// Renders 'debug draw' data for the Box2D world if it exists.
// This uses the standard debug drawing feature of Box2D, so colors will be decided by
// the Box2D engine.
func (self *UtilsDebug) Box2dWorldI(args ...interface{}) {
    self.Object.Call("box2dWorld", args)
}

// Renders 'debug draw' data for the given Box2D body.
// This uses the standard debug drawing feature of Box2D, so colors will be decided by the Box2D engine.
func (self *UtilsDebug) Box2dBody(sprite *Sprite) {
    self.Object.Call("box2dBody", sprite)
}

// Renders 'debug draw' data for the given Box2D body.
// This uses the standard debug drawing feature of Box2D, so colors will be decided by the Box2D engine.
func (self *UtilsDebug) Box2dBody1O(sprite *Sprite, color string) {
    self.Object.Call("box2dBody", sprite, color)
}

// Renders 'debug draw' data for the given Box2D body.
// This uses the standard debug drawing feature of Box2D, so colors will be decided by the Box2D engine.
func (self *UtilsDebug) Box2dBodyI(args ...interface{}) {
    self.Object.Call("box2dBody", args)
}

// Destroy this object.
func (self *UtilsDebug) Destroy() {
    self.Object.Call("destroy")
}

// Destroy this object.
func (self *UtilsDebug) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
