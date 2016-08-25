// Package phaser Automatic generation for PIXI.CanvasRenderer
// generated file CanvasRenderer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// CanvasRenderer The CanvasRenderer draws the Stage and all its content onto a 2d canvas. This renderer should be used for browsers that do not support webGL.
// Don't forget to add the CanvasRenderer.view to your DOM or you will not see anything :)
type CanvasRenderer struct {
    *js.Object
}

// NewCanvasRenderer The CanvasRenderer draws the Stage and all its content onto a 2d canvas. This renderer should be used for browsers that do not support webGL.
// Don't forget to add the CanvasRenderer.view to your DOM or you will not see anything :)
func NewCanvasRenderer(game *PhaserGame) *CanvasRenderer {
    return &CanvasRenderer{js.Global.Get("PIXI").Get("CanvasRenderer").New(game)}
}
// NewCanvasRendererI The CanvasRenderer draws the Stage and all its content onto a 2d canvas. This renderer should be used for browsers that do not support webGL.
// Don't forget to add the CanvasRenderer.view to your DOM or you will not see anything :)
func NewCanvasRendererI(args ...interface{}) *CanvasRenderer {
    return &CanvasRenderer{js.Global.Get("PIXI").Get("CanvasRenderer").New(args)}
}



// CanvasRenderer Binding conversion method to CanvasRenderer point 
func ToCanvasRenderer(jsStruct interface{}) *CanvasRenderer {
    if object, ok := jsStruct.(*js.Object); ok {
		return &CanvasRenderer{Object: object}
	}
	return nil
}



// Game empty description
func (self *CanvasRenderer) Game() *PhaserGame{
    return &PhaserGame{self.Object.Get("game")}
}

// SetGameA empty description
func (self *CanvasRenderer) SetGameA(member *PhaserGame) {
    self.Object.Set("game", member)
}

// Type The renderer type.
func (self *CanvasRenderer) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The renderer type.
func (self *CanvasRenderer) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Resolution The resolution of the canvas.
func (self *CanvasRenderer) Resolution() int{
    return self.Object.Get("resolution").Int()
}

// SetResolutionA The resolution of the canvas.
func (self *CanvasRenderer) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// ClearBeforeRender This sets if the CanvasRenderer will clear the canvas or not before the new render pass.
// If the Stage is NOT transparent Pixi will use a canvas sized fillRect operation every frame to set the canvas background color.
// If the Stage is transparent Pixi will use clearRect to clear the canvas every frame.
// Disable this by setting this to false. For example if your game has a canvas filling background image you often don't need this set.
func (self *CanvasRenderer) ClearBeforeRender() bool{
    return self.Object.Get("clearBeforeRender").Bool()
}

// SetClearBeforeRenderA This sets if the CanvasRenderer will clear the canvas or not before the new render pass.
// If the Stage is NOT transparent Pixi will use a canvas sized fillRect operation every frame to set the canvas background color.
// If the Stage is transparent Pixi will use clearRect to clear the canvas every frame.
// Disable this by setting this to false. For example if your game has a canvas filling background image you often don't need this set.
func (self *CanvasRenderer) SetClearBeforeRenderA(member bool) {
    self.Object.Set("clearBeforeRender", member)
}

// Transparent Whether the render view is transparent
func (self *CanvasRenderer) Transparent() bool{
    return self.Object.Get("transparent").Bool()
}

// SetTransparentA Whether the render view is transparent
func (self *CanvasRenderer) SetTransparentA(member bool) {
    self.Object.Set("transparent", member)
}

// AutoResize Whether the render view should be resized automatically
func (self *CanvasRenderer) AutoResize() bool{
    return self.Object.Get("autoResize").Bool()
}

// SetAutoResizeA Whether the render view should be resized automatically
func (self *CanvasRenderer) SetAutoResizeA(member bool) {
    self.Object.Set("autoResize", member)
}

// Width The width of the canvas view
func (self *CanvasRenderer) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the canvas view
func (self *CanvasRenderer) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the canvas view
func (self *CanvasRenderer) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the canvas view
func (self *CanvasRenderer) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// View The canvas element that everything is drawn to.
func (self *CanvasRenderer) View() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("view"))
}

// SetViewA The canvas element that everything is drawn to.
func (self *CanvasRenderer) SetViewA(member dom.HTMLCanvasElement) {
    self.Object.Set("view", member)
}

// Context The canvas 2d context that everything is drawn with
func (self *CanvasRenderer) Context() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("context"))
}

// SetContextA The canvas 2d context that everything is drawn with
func (self *CanvasRenderer) SetContextA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("context", member)
}

// Refresh Boolean flag controlling canvas refresh.
func (self *CanvasRenderer) Refresh() bool{
    return self.Object.Get("refresh").Bool()
}

// SetRefreshA Boolean flag controlling canvas refresh.
func (self *CanvasRenderer) SetRefreshA(member bool) {
    self.Object.Set("refresh", member)
}

// Count Internal var.
func (self *CanvasRenderer) Count() int{
    return self.Object.Get("count").Int()
}

// SetCountA Internal var.
func (self *CanvasRenderer) SetCountA(member int) {
    self.Object.Set("count", member)
}

// CanvasMaskManager Instance of a PIXI.CanvasMaskManager, handles masking when using the canvas renderer
func (self *CanvasRenderer) CanvasMaskManager() *CanvasMaskManager{
    return &CanvasMaskManager{self.Object.Get("CanvasMaskManager")}
}

// SetCanvasMaskManagerA Instance of a PIXI.CanvasMaskManager, handles masking when using the canvas renderer
func (self *CanvasRenderer) SetCanvasMaskManagerA(member *CanvasMaskManager) {
    self.Object.Set("CanvasMaskManager", member)
}

// RenderSession The render session is just a bunch of parameter used for rendering
func (self *CanvasRenderer) RenderSession() interface{}{
    return self.Object.Get("renderSession")
}

// SetRenderSessionA The render session is just a bunch of parameter used for rendering
func (self *CanvasRenderer) SetRenderSessionA(member interface{}) {
    self.Object.Set("renderSession", member)
}


// Render Renders the Stage to this canvas view
func (self *CanvasRenderer) Render(stage *Stage) {
    self.Object.Call("render", stage)
}

// RenderI Renders the Stage to this canvas view
func (self *CanvasRenderer) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// Destroy Removes everything from the renderer and optionally removes the Canvas DOM element.
func (self *CanvasRenderer) Destroy() {
    self.Object.Call("destroy")
}

// Destroy1O Removes everything from the renderer and optionally removes the Canvas DOM element.
func (self *CanvasRenderer) Destroy1O(removeView bool) {
    self.Object.Call("destroy", removeView)
}

// DestroyI Removes everything from the renderer and optionally removes the Canvas DOM element.
func (self *CanvasRenderer) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Resize Resizes the canvas view to the specified width and height
func (self *CanvasRenderer) Resize(width int, height int) {
    self.Object.Call("resize", width, height)
}

// ResizeI Resizes the canvas view to the specified width and height
func (self *CanvasRenderer) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// RenderDisplayObject Renders a display object
func (self *CanvasRenderer) RenderDisplayObject(displayObject *DisplayObject, context *dom.CanvasRenderingContext2D) {
    self.Object.Call("renderDisplayObject", displayObject, context)
}

// RenderDisplayObject1O Renders a display object
func (self *CanvasRenderer) RenderDisplayObject1O(displayObject *DisplayObject, context *dom.CanvasRenderingContext2D, matrix *Matrix) {
    self.Object.Call("renderDisplayObject", displayObject, context, matrix)
}

// RenderDisplayObjectI Renders a display object
func (self *CanvasRenderer) RenderDisplayObjectI(args ...interface{}) {
    self.Object.Call("renderDisplayObject", args)
}

// MapBlendModes Maps Pixi blend modes to canvas blend modes.
func (self *CanvasRenderer) MapBlendModes() {
    self.Object.Call("mapBlendModes")
}

// MapBlendModesI Maps Pixi blend modes to canvas blend modes.
func (self *CanvasRenderer) MapBlendModesI(args ...interface{}) {
    self.Object.Call("mapBlendModes", args)
}

