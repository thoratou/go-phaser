// Automatic generation for PIXI.CanvasRenderer
// generated file CanvasRenderer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// The CanvasRenderer draws the Stage and all its content onto a 2d canvas. This renderer should be used for browsers that do not support webGL.
// Don't forget to add the CanvasRenderer.view to your DOM or you will not see anything :)
type CanvasRenderer struct {
    *js.Object
}


// 
func (self *CanvasRenderer) GetGame() *PhaserGame{
    return &PhaserGame{self.Get("game")}
}

// 
func (self *CanvasRenderer) SetGame(member *PhaserGame) {
    self.Set("game", member)
}

// The renderer type.
func (self *CanvasRenderer) GetType() int{
    return self.Get("type").Int()
}

// The renderer type.
func (self *CanvasRenderer) SetType(member int) {
    self.Set("type", member)
}

// The resolution of the canvas.
func (self *CanvasRenderer) GetResolution() int{
    return self.Get("resolution").Int()
}

// The resolution of the canvas.
func (self *CanvasRenderer) SetResolution(member int) {
    self.Set("resolution", member)
}

// This sets if the CanvasRenderer will clear the canvas or not before the new render pass.
// If the Stage is NOT transparent Pixi will use a canvas sized fillRect operation every frame to set the canvas background color.
// If the Stage is transparent Pixi will use clearRect to clear the canvas every frame.
// Disable this by setting this to false. For example if your game has a canvas filling background image you often don't need this set.
func (self *CanvasRenderer) GetClearBeforeRender() bool{
    return self.Get("clearBeforeRender").Bool()
}

// This sets if the CanvasRenderer will clear the canvas or not before the new render pass.
// If the Stage is NOT transparent Pixi will use a canvas sized fillRect operation every frame to set the canvas background color.
// If the Stage is transparent Pixi will use clearRect to clear the canvas every frame.
// Disable this by setting this to false. For example if your game has a canvas filling background image you often don't need this set.
func (self *CanvasRenderer) SetClearBeforeRender(member bool) {
    self.Set("clearBeforeRender", member)
}

// Whether the render view is transparent
func (self *CanvasRenderer) GetTransparent() bool{
    return self.Get("transparent").Bool()
}

// Whether the render view is transparent
func (self *CanvasRenderer) SetTransparent(member bool) {
    self.Set("transparent", member)
}

// Whether the render view should be resized automatically
func (self *CanvasRenderer) GetAutoResize() bool{
    return self.Get("autoResize").Bool()
}

// Whether the render view should be resized automatically
func (self *CanvasRenderer) SetAutoResize(member bool) {
    self.Set("autoResize", member)
}

// The width of the canvas view
func (self *CanvasRenderer) GetWidth() int{
    return self.Get("width").Int()
}

// The width of the canvas view
func (self *CanvasRenderer) SetWidth(member int) {
    self.Set("width", member)
}

// The height of the canvas view
func (self *CanvasRenderer) GetHeight() int{
    return self.Get("height").Int()
}

// The height of the canvas view
func (self *CanvasRenderer) SetHeight(member int) {
    self.Set("height", member)
}

// The canvas element that everything is drawn to.
func (self *CanvasRenderer) GetView() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("view"))
}

// The canvas element that everything is drawn to.
func (self *CanvasRenderer) SetView(member dom.HTMLCanvasElement) {
    self.Set("view", member)
}

// The canvas 2d context that everything is drawn with
func (self *CanvasRenderer) GetContext() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Get("context"))
}

// The canvas 2d context that everything is drawn with
func (self *CanvasRenderer) SetContext(member dom.CanvasRenderingContext2D) {
    self.Set("context", member)
}

// Boolean flag controlling canvas refresh.
func (self *CanvasRenderer) GetRefresh() bool{
    return self.Get("refresh").Bool()
}

// Boolean flag controlling canvas refresh.
func (self *CanvasRenderer) SetRefresh(member bool) {
    self.Set("refresh", member)
}

// Internal var.
func (self *CanvasRenderer) GetCount() int{
    return self.Get("count").Int()
}

// Internal var.
func (self *CanvasRenderer) SetCount(member int) {
    self.Set("count", member)
}

// Instance of a PIXI.CanvasMaskManager, handles masking when using the canvas renderer
func (self *CanvasRenderer) GetCanvasMaskManager() *CanvasMaskManager{
    return &CanvasMaskManager{self.Get("CanvasMaskManager")}
}

// Instance of a PIXI.CanvasMaskManager, handles masking when using the canvas renderer
func (self *CanvasRenderer) SetCanvasMaskManager(member *CanvasMaskManager) {
    self.Set("CanvasMaskManager", member)
}

// The render session is just a bunch of parameter used for rendering
func (self *CanvasRenderer) GetRenderSession() interface{}{
    return self.Get("renderSession")
}

// The render session is just a bunch of parameter used for rendering
func (self *CanvasRenderer) SetRenderSession(member interface{}) {
    self.Set("renderSession", member)
}



// Renders the Stage to this canvas view
func (self *CanvasRenderer) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// Removes everything from the renderer and optionally removes the Canvas DOM element.
func (self *CanvasRenderer) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Resizes the canvas view to the specified width and height
func (self *CanvasRenderer) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// Renders a display object
func (self *CanvasRenderer) RenderDisplayObjectI(args ...interface{}) {
    self.Call("renderDisplayObject", args)
}

// Maps Pixi blend modes to canvas blend modes.
func (self *CanvasRenderer) MapBlendModesI(args ...interface{}) {
    self.Call("mapBlendModes", args)
}
