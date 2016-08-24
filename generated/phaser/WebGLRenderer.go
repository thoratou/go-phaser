// Package phaser Automatic generation for PIXI.WebGLRenderer
// generated file WebGLRenderer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// WebGLRenderer The WebGLRenderer draws the stage and all its content onto a webGL enabled canvas. This renderer
// should be used for browsers that support webGL. This Render works by automatically managing webGLBatchs.
// So no need for Sprite Batches or Sprite Clouds.
// Don't forget to add the view to your DOM or you will not see anything :)
type WebGLRenderer struct {
    *js.Object
}

// NewWebGLRenderer The WebGLRenderer draws the stage and all its content onto a webGL enabled canvas. This renderer
// should be used for browsers that support webGL. This Render works by automatically managing webGLBatchs.
// So no need for Sprite Batches or Sprite Clouds.
// Don't forget to add the view to your DOM or you will not see anything :)
func NewWebGLRenderer(game *PhaserGame) *WebGLRenderer {
    return &WebGLRenderer{js.Global.Get("PIXI").Get("WebGLRenderer").New(game)}
}
// NewWebGLRendererI The WebGLRenderer draws the stage and all its content onto a webGL enabled canvas. This renderer
// should be used for browsers that support webGL. This Render works by automatically managing webGLBatchs.
// So no need for Sprite Batches or Sprite Clouds.
// Don't forget to add the view to your DOM or you will not see anything :)
func NewWebGLRendererI(args ...interface{}) *WebGLRenderer {
    return &WebGLRenderer{js.Global.Get("PIXI").Get("WebGLRenderer").New(args)}
}



// Game empty description
func (self *WebGLRenderer) Game() *PhaserGame{
    return &PhaserGame{self.Object.Get("game")}
}

// SetGameA empty description
func (self *WebGLRenderer) SetGameA(member *PhaserGame) {
    self.Object.Set("game", member)
}

// Type empty description
func (self *WebGLRenderer) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA empty description
func (self *WebGLRenderer) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// Resolution The resolution of the renderer
func (self *WebGLRenderer) Resolution() int{
    return self.Object.Get("resolution").Int()
}

// SetResolutionA The resolution of the renderer
func (self *WebGLRenderer) SetResolutionA(member int) {
    self.Object.Set("resolution", member)
}

// Transparent Whether the render view is transparent
func (self *WebGLRenderer) Transparent() bool{
    return self.Object.Get("transparent").Bool()
}

// SetTransparentA Whether the render view is transparent
func (self *WebGLRenderer) SetTransparentA(member bool) {
    self.Object.Set("transparent", member)
}

// AutoResize Whether the render view should be resized automatically
func (self *WebGLRenderer) AutoResize() bool{
    return self.Object.Get("autoResize").Bool()
}

// SetAutoResizeA Whether the render view should be resized automatically
func (self *WebGLRenderer) SetAutoResizeA(member bool) {
    self.Object.Set("autoResize", member)
}

// PreserveDrawingBuffer The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
func (self *WebGLRenderer) PreserveDrawingBuffer() bool{
    return self.Object.Get("preserveDrawingBuffer").Bool()
}

// SetPreserveDrawingBufferA The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
func (self *WebGLRenderer) SetPreserveDrawingBufferA(member bool) {
    self.Object.Set("preserveDrawingBuffer", member)
}

// ClearBeforeRender This sets if the WebGLRenderer will clear the context texture or not before the new render pass. If true:
// If the Stage is NOT transparent, Pixi will clear to alpha (0, 0, 0, 0).
// If the Stage is transparent, Pixi will clear to the target Stage's background color.
// Disable this by setting this to false. For example: if your game has a canvas filling background image, you often don't need this set.
func (self *WebGLRenderer) ClearBeforeRender() bool{
    return self.Object.Get("clearBeforeRender").Bool()
}

// SetClearBeforeRenderA This sets if the WebGLRenderer will clear the context texture or not before the new render pass. If true:
// If the Stage is NOT transparent, Pixi will clear to alpha (0, 0, 0, 0).
// If the Stage is transparent, Pixi will clear to the target Stage's background color.
// Disable this by setting this to false. For example: if your game has a canvas filling background image, you often don't need this set.
func (self *WebGLRenderer) SetClearBeforeRenderA(member bool) {
    self.Object.Set("clearBeforeRender", member)
}

// Width The width of the canvas view
func (self *WebGLRenderer) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the canvas view
func (self *WebGLRenderer) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the canvas view
func (self *WebGLRenderer) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the canvas view
func (self *WebGLRenderer) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// View The canvas element that everything is drawn to
func (self *WebGLRenderer) View() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("view"))
}

// SetViewA The canvas element that everything is drawn to
func (self *WebGLRenderer) SetViewA(member dom.HTMLCanvasElement) {
    self.Object.Set("view", member)
}

// Projection empty description
func (self *WebGLRenderer) Projection() *Point{
    return &Point{self.Object.Get("projection")}
}

// SetProjectionA empty description
func (self *WebGLRenderer) SetProjectionA(member *Point) {
    self.Object.Set("projection", member)
}

// Offset empty description
func (self *WebGLRenderer) Offset() *Point{
    return &Point{self.Object.Get("offset")}
}

// SetOffsetA empty description
func (self *WebGLRenderer) SetOffsetA(member *Point) {
    self.Object.Set("offset", member)
}

// ShaderManager Deals with managing the shader programs and their attribs
func (self *WebGLRenderer) ShaderManager() *WebGLShaderManager{
    return &WebGLShaderManager{self.Object.Get("shaderManager")}
}

// SetShaderManagerA Deals with managing the shader programs and their attribs
func (self *WebGLRenderer) SetShaderManagerA(member *WebGLShaderManager) {
    self.Object.Set("shaderManager", member)
}

// SpriteBatch Manages the rendering of sprites
func (self *WebGLRenderer) SpriteBatch() *WebGLSpriteBatch{
    return &WebGLSpriteBatch{self.Object.Get("spriteBatch")}
}

// SetSpriteBatchA Manages the rendering of sprites
func (self *WebGLRenderer) SetSpriteBatchA(member *WebGLSpriteBatch) {
    self.Object.Set("spriteBatch", member)
}

// MaskManager Manages the masks using the stencil buffer
func (self *WebGLRenderer) MaskManager() *WebGLMaskManager{
    return &WebGLMaskManager{self.Object.Get("maskManager")}
}

// SetMaskManagerA Manages the masks using the stencil buffer
func (self *WebGLRenderer) SetMaskManagerA(member *WebGLMaskManager) {
    self.Object.Set("maskManager", member)
}

// FilterManager Manages the filters
func (self *WebGLRenderer) FilterManager() *WebGLFilterManager{
    return &WebGLFilterManager{self.Object.Get("filterManager")}
}

// SetFilterManagerA Manages the filters
func (self *WebGLRenderer) SetFilterManagerA(member *WebGLFilterManager) {
    self.Object.Set("filterManager", member)
}

// StencilManager Manages the stencil buffer
func (self *WebGLRenderer) StencilManager() *WebGLStencilManager{
    return &WebGLStencilManager{self.Object.Get("stencilManager")}
}

// SetStencilManagerA Manages the stencil buffer
func (self *WebGLRenderer) SetStencilManagerA(member *WebGLStencilManager) {
    self.Object.Set("stencilManager", member)
}

// BlendModeManager Manages the blendModes
func (self *WebGLRenderer) BlendModeManager() *WebGLBlendModeManager{
    return &WebGLBlendModeManager{self.Object.Get("blendModeManager")}
}

// SetBlendModeManagerA Manages the blendModes
func (self *WebGLRenderer) SetBlendModeManagerA(member *WebGLBlendModeManager) {
    self.Object.Set("blendModeManager", member)
}

// RenderSession empty description
func (self *WebGLRenderer) RenderSession() interface{}{
    return self.Object.Get("renderSession")
}

// SetRenderSessionA empty description
func (self *WebGLRenderer) SetRenderSessionA(member interface{}) {
    self.Object.Set("renderSession", member)
}


// InitContext empty description
func (self *WebGLRenderer) InitContext() {
    self.Object.Call("initContext")
}

// InitContextI empty description
func (self *WebGLRenderer) InitContextI(args ...interface{}) {
    self.Object.Call("initContext", args)
}

// Render Renders the stage to its webGL view
func (self *WebGLRenderer) Render(stage *Stage) {
    self.Object.Call("render", stage)
}

// RenderI Renders the stage to its webGL view
func (self *WebGLRenderer) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// RenderDisplayObject Renders a Display Object.
func (self *WebGLRenderer) RenderDisplayObject(displayObject *DisplayObject, projection *Point, buffer []interface{}) {
    self.Object.Call("renderDisplayObject", displayObject, projection, buffer)
}

// RenderDisplayObjectI Renders a Display Object.
func (self *WebGLRenderer) RenderDisplayObjectI(args ...interface{}) {
    self.Object.Call("renderDisplayObject", args)
}

// Resize Resizes the webGL view to the specified width and height.
func (self *WebGLRenderer) Resize(width int, height int) {
    self.Object.Call("resize", width, height)
}

// ResizeI Resizes the webGL view to the specified width and height.
func (self *WebGLRenderer) ResizeI(args ...interface{}) {
    self.Object.Call("resize", args)
}

// UpdateTexture Updates and Creates a WebGL texture for the renderers context.
func (self *WebGLRenderer) UpdateTexture(texture *Texture) bool{
    return self.Object.Call("updateTexture", texture).Bool()
}

// UpdateTextureI Updates and Creates a WebGL texture for the renderers context.
func (self *WebGLRenderer) UpdateTextureI(args ...interface{}) bool{
    return self.Object.Call("updateTexture", args).Bool()
}

// Destroy Removes everything from the renderer (event listeners, spritebatch, etc...)
func (self *WebGLRenderer) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Removes everything from the renderer (event listeners, spritebatch, etc...)
func (self *WebGLRenderer) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// MapBlendModes Maps Pixi blend modes to WebGL blend modes.
func (self *WebGLRenderer) MapBlendModes() {
    self.Object.Call("mapBlendModes")
}

// MapBlendModesI Maps Pixi blend modes to WebGL blend modes.
func (self *WebGLRenderer) MapBlendModesI(args ...interface{}) {
    self.Object.Call("mapBlendModes", args)
}

