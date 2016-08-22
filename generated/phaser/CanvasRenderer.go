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


// The CanvasRenderer draws the Stage and all its content onto a 2d canvas. This renderer should be used for browsers that do not support webGL.
// Don't forget to add the CanvasRenderer.view to your DOM or you will not see anything :)
func NewCanvasRenderer(game *PhaserGame) *CanvasRenderer {
    return &CanvasRenderer{js.Global.Call("PIXI.CanvasRenderer", game)}
}

// The CanvasRenderer draws the Stage and all its content onto a 2d canvas. This renderer should be used for browsers that do not support webGL.
// Don't forget to add the CanvasRenderer.view to your DOM or you will not see anything :)
func NewCanvasRendererI(args ...interface{}) *CanvasRenderer {
    return &CanvasRenderer{js.Global.Call("PIXI.CanvasRenderer", args)}
}



// 
func (self *CanvasRenderer) GetGameA() *PhaserGame{
    return &PhaserGame{self.Object.Get("game")}
}

// 
func (self *CanvasRenderer) SetGameA(member *PhaserGame) {
    self.Object.Set("game", member)
}

// The renderer type.
func (self *CanvasRenderer) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The renderer type.
func (self *CanvasRenderer) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// The resolution of the canvas.
func (self *CanvasRenderer) GetResolutionA() int{
    return self.Object.Get("resolution").Int()
}

// The resolution of the canvas.
func (self *CanvasRenderer) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// This sets if the CanvasRenderer will clear the canvas or not before the new render pass.
// If the Stage is NOT transparent Pixi will use a canvas sized fillRect operation every frame to set the canvas background color.
// If the Stage is transparent Pixi will use clearRect to clear the canvas every frame.
// Disable this by setting this to false. For example if your game has a canvas filling background image you often don't need this set.
func (self *CanvasRenderer) GetClearBeforeRenderA() bool{
    return self.Object.Get("clearBeforeRender").Bool()
}

// This sets if the CanvasRenderer will clear the canvas or not before the new render pass.
// If the Stage is NOT transparent Pixi will use a canvas sized fillRect operation every frame to set the canvas background color.
// If the Stage is transparent Pixi will use clearRect to clear the canvas every frame.
// Disable this by setting this to false. For example if your game has a canvas filling background image you often don't need this set.
func (self *CanvasRenderer) SetClearBeforeRenderA(member bool) {
    self.Object.Set("clearBeforeRender", member)
}

// Whether the render view is transparent
func (self *CanvasRenderer) GetTransparentA() bool{
    return self.Object.Get("transparent").Bool()
}

// Whether the render view is transparent
func (self *CanvasRenderer) SetTransparentA(member bool) {
    self.Object.Set("transparent", member)
}

// Whether the render view should be resized automatically
func (self *CanvasRenderer) GetAutoResizeA() bool{
    return self.Object.Get("autoResize").Bool()
}

// Whether the render view should be resized automatically
func (self *CanvasRenderer) SetAutoResizeA(member bool) {
    self.Object.Set("autoResize", member)
}

// The width of the canvas view
func (self *CanvasRenderer) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the canvas view
func (self *CanvasRenderer) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the canvas view
func (self *CanvasRenderer) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the canvas view
func (self *CanvasRenderer) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The canvas element that everything is drawn to.
func (self *CanvasRenderer) GetViewA() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("view"))
}

// The canvas element that everything is drawn to.
func (self *CanvasRenderer) SetViewA(member dom.HTMLCanvasElement) {
    self.Object.Set("view", member)
}

// The canvas 2d context that everything is drawn with
func (self *CanvasRenderer) GetContextA() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("context"))
}

// The canvas 2d context that everything is drawn with
func (self *CanvasRenderer) SetContextA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("context", member)
}

// Boolean flag controlling canvas refresh.
func (self *CanvasRenderer) GetRefreshA() bool{
    return self.Object.Get("refresh").Bool()
}

// Boolean flag controlling canvas refresh.
func (self *CanvasRenderer) SetRefreshA(member bool) {
    self.Object.Set("refresh", member)
}

// Internal var.
func (self *CanvasRenderer) GetCountA() int{
    return self.Object.Get("count").Int()
}

// Internal var.
func (self *CanvasRenderer) SetCountA(member int) {
    self.Object.Set("count", member)
}

// Instance of a PIXI.CanvasMaskManager, handles masking when using the canvas renderer
func (self *CanvasRenderer) GetCanvasMaskManagerA() *CanvasMaskManager{
    return &CanvasMaskManager{self.Object.Get("CanvasMaskManager")}
}

// Instance of a PIXI.CanvasMaskManager, handles masking when using the canvas renderer
func (self *CanvasRenderer) SetCanvasMaskManagerA(member *CanvasMaskManager) {
    self.Object.Set("CanvasMaskManager", member)
}

// The render session is just a bunch of parameter used for rendering
func (self *CanvasRenderer) GetRenderSessionA() interface{}{
    return self.Object.Get("renderSession")
}

// The render session is just a bunch of parameter used for rendering
func (self *CanvasRenderer) SetRenderSessionA(member interface{}) {
    self.Object.Set("renderSession", member)
}



// Renders the Stage to this canvas view
func (self *CanvasRenderer) Render(stage *Stage) {
    self.Object.Call("render", stage)
}

// Renders the Stage to this canvas view
func (self *CanvasRenderer) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// Removes everything from the renderer and optionally removes the Canvas DOM element.
func (self *CanvasRenderer) Destroy() {
    self.Object.Call("destroy")
}

// Removes everything from the renderer and optionally removes the Canvas DOM element.
func (self *CanvasRenderer) Destroy1O(removeView bool) {
    self.Object.Call("destroy", removeView)
}

// Removes everything from the renderer and optionally removes the Canvas DOM element.
func (self *CanvasRenderer) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Resizes the canvas view to the specified width and height
func (self *CanvasRenderer) Resize(width int, height int) {
    self.Object.Call("resize", width, height)
}

// Resizes the canvas view to the specified width and height
func (self *CanvasRenderer) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// Renders a display object
func (self *CanvasRenderer) RenderDisplayObject(displayObject *DisplayObject, context *dom.CanvasRenderingContext2D) {
    self.Object.Call("renderDisplayObject", displayObject, context)
}

// Renders a display object
func (self *CanvasRenderer) RenderDisplayObject1O(displayObject *DisplayObject, context *dom.CanvasRenderingContext2D, matrix *Matrix) {
    self.Object.Call("renderDisplayObject", displayObject, context, matrix)
}

// Renders a display object
func (self *CanvasRenderer) RenderDisplayObjectI(args ...interface{}) {
    self.Object.Call("renderDisplayObject", args)
}

// Maps Pixi blend modes to canvas blend modes.
func (self *CanvasRenderer) MapBlendModes() {
    self.Object.Call("mapBlendModes")
}

// Maps Pixi blend modes to canvas blend modes.
func (self *CanvasRenderer) MapBlendModesI(args ...interface{}) {
    self.Object.Call("mapBlendModes", args)
}
