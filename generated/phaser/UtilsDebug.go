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


// A reference to the currently running Game.
func (self *UtilsDebug) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *UtilsDebug) SetGame(member Game) {
    self.Set("game", member)
}

// If debugging in WebGL mode we need this.
func (self *UtilsDebug) GetSprite() Image{
    return Image{self.Get("sprite")}
}

// If debugging in WebGL mode we need this.
func (self *UtilsDebug) SetSprite(member Image) {
    self.Set("sprite", member)
}

// In WebGL mode this BitmapData contains a copy of the debug canvas.
func (self *UtilsDebug) GetBmd() BitmapData{
    return BitmapData{self.Get("bmd")}
}

// In WebGL mode this BitmapData contains a copy of the debug canvas.
func (self *UtilsDebug) SetBmd(member BitmapData) {
    self.Set("bmd", member)
}

// The canvas to which Debug calls draws.
func (self *UtilsDebug) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("canvas"))
}

// The canvas to which Debug calls draws.
func (self *UtilsDebug) SetCanvas(member dom.HTMLCanvasElement) {
    self.Set("canvas", member)
}

// The 2d context of the canvas.
func (self *UtilsDebug) GetContext() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Get("context"))
}

// The 2d context of the canvas.
func (self *UtilsDebug) SetContext(member dom.CanvasRenderingContext2D) {
    self.Set("context", member)
}

// The font that the debug information is rendered in.
func (self *UtilsDebug) GetFont() string{
    return self.Get("font").String()
}

// The font that the debug information is rendered in.
func (self *UtilsDebug) SetFont(member string) {
    self.Set("font", member)
}

// The spacing between columns.
func (self *UtilsDebug) GetColumnWidth() float64{
    return self.Get("columnWidth").Float()
}

// The spacing between columns.
func (self *UtilsDebug) SetColumnWidth(member float64) {
    self.Set("columnWidth", member)
}

// The line height between the debug text.
func (self *UtilsDebug) GetLineHeight() float64{
    return self.Get("lineHeight").Float()
}

// The line height between the debug text.
func (self *UtilsDebug) SetLineHeight(member float64) {
    self.Set("lineHeight", member)
}

// Should the text be rendered with a slight shadow? Makes it easier to read on different types of background.
func (self *UtilsDebug) GetRenderShadow() bool{
    return self.Get("renderShadow").Bool()
}

// Should the text be rendered with a slight shadow? Makes it easier to read on different types of background.
func (self *UtilsDebug) SetRenderShadow(member bool) {
    self.Set("renderShadow", member)
}

// The current X position the debug information will be rendered at.
func (self *UtilsDebug) GetCurrentX() float64{
    return self.Get("currentX").Float()
}

// The current X position the debug information will be rendered at.
func (self *UtilsDebug) SetCurrentX(member float64) {
    self.Set("currentX", member)
}

// The current Y position the debug information will be rendered at.
func (self *UtilsDebug) GetCurrentY() float64{
    return self.Get("currentY").Float()
}

// The current Y position the debug information will be rendered at.
func (self *UtilsDebug) SetCurrentY(member float64) {
    self.Set("currentY", member)
}

// The alpha of the Debug context, set before all debug information is rendered to it.
func (self *UtilsDebug) GetCurrentAlpha() float64{
    return self.Get("currentAlpha").Float()
}

// The alpha of the Debug context, set before all debug information is rendered to it.
func (self *UtilsDebug) SetCurrentAlpha(member float64) {
    self.Set("currentAlpha", member)
}

// Does the canvas need re-rendering?
func (self *UtilsDebug) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// Does the canvas need re-rendering?
func (self *UtilsDebug) SetDirty(member bool) {
    self.Set("dirty", member)
}



// Internal method that boots the debug displayer.
func (self *UtilsDebug) BootI(args ...interface{}) {
    self.Call("boot", args)
}

// Internal method that resizes the BitmapData and Canvas.
// Called by ScaleManager.onSizeChange only in WebGL mode.
func (self *UtilsDebug) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// Internal method that clears the canvas (if a Sprite) ready for a new debug session.
func (self *UtilsDebug) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Clears the Debug canvas.
func (self *UtilsDebug) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Internal method that resets and starts the debug output values.
func (self *UtilsDebug) StartI(args ...interface{}) {
    self.Call("start", args)
}

// Internal method that stops the debug output.
func (self *UtilsDebug) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// Internal method that outputs a single line of text split over as many columns as needed, one per parameter.
func (self *UtilsDebug) LineI(args ...interface{}) {
    self.Call("line", args)
}

// Render Sound information, including decoded state, duration, volume and more.
func (self *UtilsDebug) SoundInfoI(args ...interface{}) {
    self.Call("soundInfo", args)
}

// Render camera information including dimensions and location.
func (self *UtilsDebug) CameraInfoI(args ...interface{}) {
    self.Call("cameraInfo", args)
}

// Render Timer information.
func (self *UtilsDebug) TimerI(args ...interface{}) {
    self.Call("timer", args)
}

// Renders the Pointer.circle object onto the stage in green if down or red if up along with debug text.
func (self *UtilsDebug) PointerI(args ...interface{}) {
    self.Call("pointer", args)
}

// Render Sprite Input Debug information.
func (self *UtilsDebug) SpriteInputInfoI(args ...interface{}) {
    self.Call("spriteInputInfo", args)
}

// Renders Phaser.Key object information.
func (self *UtilsDebug) KeyI(args ...interface{}) {
    self.Call("key", args)
}

// Render debug information about the Input object.
func (self *UtilsDebug) InputInfoI(args ...interface{}) {
    self.Call("inputInfo", args)
}

// Renders the Sprites bounds. Note: This is really expensive as it has to calculate the bounds every time you call it!
func (self *UtilsDebug) SpriteBoundsI(args ...interface{}) {
    self.Call("spriteBounds", args)
}

// Renders the Rope's segments. Note: This is really expensive as it has to calculate new segments every time you call it
func (self *UtilsDebug) RopeSegmentsI(args ...interface{}) {
    self.Call("ropeSegments", args)
}

// Render debug infos (including name, bounds info, position and some other properties) about the Sprite.
func (self *UtilsDebug) SpriteInfoI(args ...interface{}) {
    self.Call("spriteInfo", args)
}

// Renders the sprite coordinates in local, positional and world space.
func (self *UtilsDebug) SpriteCoordsI(args ...interface{}) {
    self.Call("spriteCoords", args)
}

// Renders Line information in the given color.
func (self *UtilsDebug) LineInfoI(args ...interface{}) {
    self.Call("lineInfo", args)
}

// Renders a single pixel at the given size.
func (self *UtilsDebug) PixelI(args ...interface{}) {
    self.Call("pixel", args)
}

// Renders a Phaser geometry object including Rectangle, Circle, Point or Line.
func (self *UtilsDebug) GeomI(args ...interface{}) {
    self.Call("geom", args)
}

// Render a string of text.
func (self *UtilsDebug) TextI(args ...interface{}) {
    self.Call("text", args)
}

// Visually renders a QuadTree to the display.
func (self *UtilsDebug) QuadTreeI(args ...interface{}) {
    self.Call("quadTree", args)
}

// Render a Sprites Physics body if it has one set. The body is rendered as a filled or stroked rectangle.
// This only works for Arcade Physics, Ninja Physics (AABB and Circle only) and Box2D Physics bodies.
// To display a P2 Physics body you should enable debug mode on the body when creating it.
func (self *UtilsDebug) BodyI(args ...interface{}) {
    self.Call("body", args)
}

// Render a Sprites Physic Body information.
func (self *UtilsDebug) BodyInfoI(args ...interface{}) {
    self.Call("bodyInfo", args)
}

// Renders 'debug draw' data for the Box2D world if it exists.
// This uses the standard debug drawing feature of Box2D, so colors will be decided by
// the Box2D engine.
func (self *UtilsDebug) Box2dWorldI(args ...interface{}) {
    self.Call("box2dWorld", args)
}

// Renders 'debug draw' data for the given Box2D body.
// This uses the standard debug drawing feature of Box2D, so colors will be decided by the Box2D engine.
func (self *UtilsDebug) Box2dBodyI(args ...interface{}) {
    self.Call("box2dBody", args)
}

// Destroy this object.
func (self *UtilsDebug) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
