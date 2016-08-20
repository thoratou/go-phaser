// Automatic generation for PIXI.WebGLRenderer
// generated file WebGLRenderer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// The WebGLRenderer draws the stage and all its content onto a webGL enabled canvas. This renderer
// should be used for browsers that support webGL. This Render works by automatically managing webGLBatchs.
// So no need for Sprite Batches or Sprite Clouds.
// Don't forget to add the view to your DOM or you will not see anything :)
type WebGLRenderer struct {
    *js.Object
}


// 
func (self *WebGLRenderer) GetGame() PhaserGame{
    return PhaserGame{self.Get("game")}
}

// 
func (self *WebGLRenderer) SetGame(member PhaserGame) {
    self.Set("game", member)
}

// 
func (self *WebGLRenderer) GetType() float64{
    return self.Get("type").Float()
}

// 
func (self *WebGLRenderer) SetType(member float64) {
    self.Set("type", member)
}

// The resolution of the renderer
func (self *WebGLRenderer) GetResolution() float64{
    return self.Get("resolution").Float()
}

// The resolution of the renderer
func (self *WebGLRenderer) SetResolution(member float64) {
    self.Set("resolution", member)
}

// Whether the render view is transparent
func (self *WebGLRenderer) GetTransparent() bool{
    return self.Get("transparent").Bool()
}

// Whether the render view is transparent
func (self *WebGLRenderer) SetTransparent(member bool) {
    self.Set("transparent", member)
}

// Whether the render view should be resized automatically
func (self *WebGLRenderer) GetAutoResize() bool{
    return self.Get("autoResize").Bool()
}

// Whether the render view should be resized automatically
func (self *WebGLRenderer) SetAutoResize(member bool) {
    self.Set("autoResize", member)
}

// The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
func (self *WebGLRenderer) GetPreserveDrawingBuffer() bool{
    return self.Get("preserveDrawingBuffer").Bool()
}

// The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
func (self *WebGLRenderer) SetPreserveDrawingBuffer(member bool) {
    self.Set("preserveDrawingBuffer", member)
}

// This sets if the WebGLRenderer will clear the context texture or not before the new render pass. If true:
// If the Stage is NOT transparent, Pixi will clear to alpha (0, 0, 0, 0).
// If the Stage is transparent, Pixi will clear to the target Stage's background color.
// Disable this by setting this to false. For example: if your game has a canvas filling background image, you often don't need this set.
func (self *WebGLRenderer) GetClearBeforeRender() bool{
    return self.Get("clearBeforeRender").Bool()
}

// This sets if the WebGLRenderer will clear the context texture or not before the new render pass. If true:
// If the Stage is NOT transparent, Pixi will clear to alpha (0, 0, 0, 0).
// If the Stage is transparent, Pixi will clear to the target Stage's background color.
// Disable this by setting this to false. For example: if your game has a canvas filling background image, you often don't need this set.
func (self *WebGLRenderer) SetClearBeforeRender(member bool) {
    self.Set("clearBeforeRender", member)
}

// The width of the canvas view
func (self *WebGLRenderer) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the canvas view
func (self *WebGLRenderer) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the canvas view
func (self *WebGLRenderer) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the canvas view
func (self *WebGLRenderer) SetHeight(member float64) {
    self.Set("height", member)
}

// The canvas element that everything is drawn to
func (self *WebGLRenderer) GetView() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("view"))
}

// The canvas element that everything is drawn to
func (self *WebGLRenderer) SetView(member dom.HTMLCanvasElement) {
    self.Set("view", member)
}

// 
func (self *WebGLRenderer) GetProjection() Point{
    return Point{self.Get("projection")}
}

// 
func (self *WebGLRenderer) SetProjection(member Point) {
    self.Set("projection", member)
}

// 
func (self *WebGLRenderer) GetOffset() Point{
    return Point{self.Get("offset")}
}

// 
func (self *WebGLRenderer) SetOffset(member Point) {
    self.Set("offset", member)
}

// Deals with managing the shader programs and their attribs
func (self *WebGLRenderer) GetShaderManager() WebGLShaderManager{
    return WebGLShaderManager{self.Get("shaderManager")}
}

// Deals with managing the shader programs and their attribs
func (self *WebGLRenderer) SetShaderManager(member WebGLShaderManager) {
    self.Set("shaderManager", member)
}

// Manages the rendering of sprites
func (self *WebGLRenderer) GetSpriteBatch() WebGLSpriteBatch{
    return WebGLSpriteBatch{self.Get("spriteBatch")}
}

// Manages the rendering of sprites
func (self *WebGLRenderer) SetSpriteBatch(member WebGLSpriteBatch) {
    self.Set("spriteBatch", member)
}

// Manages the masks using the stencil buffer
func (self *WebGLRenderer) GetMaskManager() WebGLMaskManager{
    return WebGLMaskManager{self.Get("maskManager")}
}

// Manages the masks using the stencil buffer
func (self *WebGLRenderer) SetMaskManager(member WebGLMaskManager) {
    self.Set("maskManager", member)
}

// Manages the filters
func (self *WebGLRenderer) GetFilterManager() WebGLFilterManager{
    return WebGLFilterManager{self.Get("filterManager")}
}

// Manages the filters
func (self *WebGLRenderer) SetFilterManager(member WebGLFilterManager) {
    self.Set("filterManager", member)
}

// Manages the stencil buffer
func (self *WebGLRenderer) GetStencilManager() WebGLStencilManager{
    return WebGLStencilManager{self.Get("stencilManager")}
}

// Manages the stencil buffer
func (self *WebGLRenderer) SetStencilManager(member WebGLStencilManager) {
    self.Set("stencilManager", member)
}

// Manages the blendModes
func (self *WebGLRenderer) GetBlendModeManager() WebGLBlendModeManager{
    return WebGLBlendModeManager{self.Get("blendModeManager")}
}

// Manages the blendModes
func (self *WebGLRenderer) SetBlendModeManager(member WebGLBlendModeManager) {
    self.Set("blendModeManager", member)
}

// 
func (self *WebGLRenderer) GetRenderSession() interface{}{
    return self.Get("renderSession")
}

// 
func (self *WebGLRenderer) SetRenderSession(member interface{}) {
    self.Set("renderSession", member)
}



// 
func (self *WebGLRenderer) InitContextI(args ...interface{}) {
    self.Call("initContext", args)
}

// Renders the stage to its webGL view
func (self *WebGLRenderer) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// Renders a Display Object.
func (self *WebGLRenderer) RenderDisplayObjectI(args ...interface{}) {
    self.Call("renderDisplayObject", args)
}

// Resizes the webGL view to the specified width and height.
func (self *WebGLRenderer) ResizeI(args ...interface{}) {
    self.Call("resize", args)
}

// Updates and Creates a WebGL texture for the renderers context.
func (self *WebGLRenderer) UpdateTextureI(args ...interface{}) bool{
    return self.Call("updateTexture", args).Bool()
}

// Removes everything from the renderer (event listeners, spritebatch, etc...)
func (self *WebGLRenderer) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Maps Pixi blend modes to WebGL blend modes.
func (self *WebGLRenderer) MapBlendModesI(args ...interface{}) {
    self.Call("mapBlendModes", args)
}
