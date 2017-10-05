// Package phaser Automatic generation for Phaser.BitmapData
// generated file BitmapData.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// BitmapData A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
type BitmapData struct {
    *js.Object
}

// NewBitmapData A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
func NewBitmapData(game *Game, key string) *BitmapData {
    return &BitmapData{js.Global.Get("Phaser").Get("BitmapData").New(game, key)}
}
// NewBitmapData1O A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
func NewBitmapData1O(game *Game, key string, width int) *BitmapData {
    return &BitmapData{js.Global.Get("Phaser").Get("BitmapData").New(game, key, width)}
}
// NewBitmapData2O A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
func NewBitmapData2O(game *Game, key string, width int, height int) *BitmapData {
    return &BitmapData{js.Global.Get("Phaser").Get("BitmapData").New(game, key, width, height)}
}
// NewBitmapData3O A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
func NewBitmapData3O(game *Game, key string, width int, height int, skipPool bool) *BitmapData {
    return &BitmapData{js.Global.Get("Phaser").Get("BitmapData").New(game, key, width, height, skipPool)}
}
// NewBitmapDataI A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
func NewBitmapDataI(args ...interface{}) *BitmapData {
    return &BitmapData{js.Global.Get("Phaser").Get("BitmapData").New(args)}
}



// BitmapData Binding conversion method to BitmapData point 
func ToBitmapData(jsStruct interface{}) *BitmapData {
    if object, ok := jsStruct.(*js.Object); ok {
		return &BitmapData{Object: object}
	}
	return nil
}



// Game A reference to the currently running game.
func (self *BitmapData) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *BitmapData) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Key The key of the BitmapData in the Cache, if stored there.
func (self *BitmapData) Key() string{
    return self.Object.Get("key").String()
}

// SetKeyA The key of the BitmapData in the Cache, if stored there.
func (self *BitmapData) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// Width The width of the BitmapData in pixels.
func (self *BitmapData) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the BitmapData in pixels.
func (self *BitmapData) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the BitmapData in pixels.
func (self *BitmapData) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the BitmapData in pixels.
func (self *BitmapData) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// Canvas The canvas to which this BitmapData draws.
func (self *BitmapData) Canvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("canvas"))
}

// SetCanvasA The canvas to which this BitmapData draws.
func (self *BitmapData) SetCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("canvas", member)
}

// Context The 2d context of the canvas.
func (self *BitmapData) Context() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("context"))
}

// SetContextA The 2d context of the canvas.
func (self *BitmapData) SetContextA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("context", member)
}

// Ctx A reference to BitmapData.context.
func (self *BitmapData) Ctx() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("ctx"))
}

// SetCtxA A reference to BitmapData.context.
func (self *BitmapData) SetCtxA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("ctx", member)
}

// SmoothProperty The context property needed for smoothing this Canvas.
func (self *BitmapData) SmoothProperty() string{
    return self.Object.Get("smoothProperty").String()
}

// SetSmoothPropertyA The context property needed for smoothing this Canvas.
func (self *BitmapData) SetSmoothPropertyA(member string) {
    self.Object.Set("smoothProperty", member)
}

// ImageData The context image data.
// Please note that a call to BitmapData.draw() or BitmapData.copy() does not update immediately this property for performance reason. Use BitmapData.update() to do so.
// This property is updated automatically after the first game loop, according to the dirty flag property.
func (self *BitmapData) ImageData() *ImageData{
    return &ImageData{self.Object.Get("imageData")}
}

// SetImageDataA The context image data.
// Please note that a call to BitmapData.draw() or BitmapData.copy() does not update immediately this property for performance reason. Use BitmapData.update() to do so.
// This property is updated automatically after the first game loop, according to the dirty flag property.
func (self *BitmapData) SetImageDataA(member *ImageData) {
    self.Object.Set("imageData", member)
}

// Data A Uint8ClampedArray view into BitmapData.buffer.
// Note that this is unavailable in some browsers (such as Epic Browser due to its security restrictions)
func (self *BitmapData) Data() *Uint8ClampedArray{
    return &Uint8ClampedArray{self.Object.Get("data")}
}

// SetDataA A Uint8ClampedArray view into BitmapData.buffer.
// Note that this is unavailable in some browsers (such as Epic Browser due to its security restrictions)
func (self *BitmapData) SetDataA(member *Uint8ClampedArray) {
    self.Object.Set("data", member)
}

// Pixels An Uint32Array view into BitmapData.buffer.
func (self *BitmapData) Pixels() *Uint32Array{
    return &Uint32Array{self.Object.Get("pixels")}
}

// SetPixelsA An Uint32Array view into BitmapData.buffer.
func (self *BitmapData) SetPixelsA(member *Uint32Array) {
    self.Object.Set("pixels", member)
}

// BaseTexture The PIXI.BaseTexture.
func (self *BitmapData) BaseTexture() *BaseTexture{
    return &BaseTexture{self.Object.Get("baseTexture")}
}

// SetBaseTextureA The PIXI.BaseTexture.
func (self *BitmapData) SetBaseTextureA(member *BaseTexture) {
    self.Object.Set("baseTexture", member)
}

// Texture The PIXI.Texture.
func (self *BitmapData) Texture() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// SetTextureA The PIXI.Texture.
func (self *BitmapData) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// FrameData The FrameData container this BitmapData uses for rendering.
func (self *BitmapData) FrameData() *FrameData{
    return &FrameData{self.Object.Get("frameData")}
}

// SetFrameDataA The FrameData container this BitmapData uses for rendering.
func (self *BitmapData) SetFrameDataA(member *FrameData) {
    self.Object.Set("frameData", member)
}

// TextureFrame The Frame this BitmapData uses for rendering.
func (self *BitmapData) TextureFrame() *Frame{
    return &Frame{self.Object.Get("textureFrame")}
}

// SetTextureFrameA The Frame this BitmapData uses for rendering.
func (self *BitmapData) SetTextureFrameA(member *Frame) {
    self.Object.Set("textureFrame", member)
}

// Type The const type of this object.
func (self *BitmapData) Type() int{
    return self.Object.Get("type").Int()
}

// SetTypeA The const type of this object.
func (self *BitmapData) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// DisableTextureUpload If disableTextureUpload is true this BitmapData will never send its image data to the GPU when its dirty flag is true.
func (self *BitmapData) DisableTextureUpload() bool{
    return self.Object.Get("disableTextureUpload").Bool()
}

// SetDisableTextureUploadA If disableTextureUpload is true this BitmapData will never send its image data to the GPU when its dirty flag is true.
func (self *BitmapData) SetDisableTextureUploadA(member bool) {
    self.Object.Set("disableTextureUpload", member)
}

// Dirty If dirty this BitmapData will be re-rendered.
func (self *BitmapData) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// SetDirtyA If dirty this BitmapData will be re-rendered.
func (self *BitmapData) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}


// Move Shifts the contents of this BitmapData by the distances given.
// 
// The image will wrap-around the edges on all sides if the wrap argument is true (the default).
func (self *BitmapData) Move(x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("move", x, y)}
}

// Move1O Shifts the contents of this BitmapData by the distances given.
// 
// The image will wrap-around the edges on all sides if the wrap argument is true (the default).
func (self *BitmapData) Move1O(x int, y int, wrap bool) *BitmapData{
    return &BitmapData{self.Object.Call("move", x, y, wrap)}
}

// MoveI Shifts the contents of this BitmapData by the distances given.
// 
// The image will wrap-around the edges on all sides if the wrap argument is true (the default).
func (self *BitmapData) MoveI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("move", args)}
}

// MoveH Shifts the contents of this BitmapData horizontally.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveH(distance int) *BitmapData{
    return &BitmapData{self.Object.Call("moveH", distance)}
}

// MoveH1O Shifts the contents of this BitmapData horizontally.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveH1O(distance int, wrap bool) *BitmapData{
    return &BitmapData{self.Object.Call("moveH", distance, wrap)}
}

// MoveHI Shifts the contents of this BitmapData horizontally.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveHI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("moveH", args)}
}

// MoveV Shifts the contents of this BitmapData vertically.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveV(distance int) *BitmapData{
    return &BitmapData{self.Object.Call("moveV", distance)}
}

// MoveV1O Shifts the contents of this BitmapData vertically.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveV1O(distance int, wrap bool) *BitmapData{
    return &BitmapData{self.Object.Call("moveV", distance, wrap)}
}

// MoveVI Shifts the contents of this BitmapData vertically.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveVI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("moveV", args)}
}

// Add Updates the given objects so that they use this BitmapData as their texture.
// This will replace any texture they will currently have set.
func (self *BitmapData) Add(object interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("add", object)}
}

// AddI Updates the given objects so that they use this BitmapData as their texture.
// This will replace any texture they will currently have set.
func (self *BitmapData) AddI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("add", args)}
}

// Load Takes the given Game Object, resizes this BitmapData to match it and then draws it into this BitmapDatas canvas, ready for further processing.
// The source game object is not modified by this operation.
// If the source object uses a texture as part of a Texture Atlas or Sprite Sheet, only the current frame will be used for sizing.
// If a string is given it will assume it's a cache key and look in Phaser.Cache for an image key matching the string.
func (self *BitmapData) Load(source interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("load", source)}
}

// LoadI Takes the given Game Object, resizes this BitmapData to match it and then draws it into this BitmapDatas canvas, ready for further processing.
// The source game object is not modified by this operation.
// If the source object uses a texture as part of a Texture Atlas or Sprite Sheet, only the current frame will be used for sizing.
// If a string is given it will assume it's a cache key and look in Phaser.Cache for an image key matching the string.
func (self *BitmapData) LoadI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("load", args)}
}

// Cls Clears the BitmapData context using a clearRect.
func (self *BitmapData) Cls() {
    self.Object.Call("cls")
}

// ClsI Clears the BitmapData context using a clearRect.
func (self *BitmapData) ClsI(args ...interface{}) {
    self.Object.Call("cls", args)
}

// Clear Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) Clear() *BitmapData{
    return &BitmapData{self.Object.Call("clear")}
}

// Clear1O Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) Clear1O(x int) *BitmapData{
    return &BitmapData{self.Object.Call("clear", x)}
}

// Clear2O Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) Clear2O(x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("clear", x, y)}
}

// Clear3O Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) Clear3O(x int, y int, width int) *BitmapData{
    return &BitmapData{self.Object.Call("clear", x, y, width)}
}

// Clear4O Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) Clear4O(x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("clear", x, y, width, height)}
}

// ClearI Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) ClearI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("clear", args)}
}

// Fill Fills the BitmapData with the given color.
func (self *BitmapData) Fill(r int, g int, b int) *BitmapData{
    return &BitmapData{self.Object.Call("fill", r, g, b)}
}

// Fill1O Fills the BitmapData with the given color.
func (self *BitmapData) Fill1O(r int, g int, b int, a int) *BitmapData{
    return &BitmapData{self.Object.Call("fill", r, g, b, a)}
}

// FillI Fills the BitmapData with the given color.
func (self *BitmapData) FillI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("fill", args)}
}

// GenerateTexture Creates a new Image element by converting this BitmapDatas canvas into a dataURL.
// 
// The image is then stored in the image Cache using the key given.
// 
// Finally a PIXI.Texture is created based on the image and returned.
// 
// You can apply the texture to a sprite or any other supporting object by using either the
// key or the texture. First call generateTexture:
// 
// `var texture = bitmapdata.generateTexture('ball');`
// 
// Then you can either apply the texture to a sprite:
// 
// `game.add.sprite(0, 0, texture);`
// 
// or by using the string based key:
// 
// `game.add.sprite(0, 0, 'ball');`
func (self *BitmapData) GenerateTexture(key string) *Texture{
    return &Texture{self.Object.Call("generateTexture", key)}
}

// GenerateTextureI Creates a new Image element by converting this BitmapDatas canvas into a dataURL.
// 
// The image is then stored in the image Cache using the key given.
// 
// Finally a PIXI.Texture is created based on the image and returned.
// 
// You can apply the texture to a sprite or any other supporting object by using either the
// key or the texture. First call generateTexture:
// 
// `var texture = bitmapdata.generateTexture('ball');`
// 
// Then you can either apply the texture to a sprite:
// 
// `game.add.sprite(0, 0, texture);`
// 
// or by using the string based key:
// 
// `game.add.sprite(0, 0, 'ball');`
func (self *BitmapData) GenerateTextureI(args ...interface{}) *Texture{
    return &Texture{self.Object.Call("generateTexture", args)}
}

// Resize Resizes the BitmapData. This changes the size of the underlying canvas and refreshes the buffer.
func (self *BitmapData) Resize(width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("resize", width, height)}
}

// ResizeI Resizes the BitmapData. This changes the size of the underlying canvas and refreshes the buffer.
func (self *BitmapData) ResizeI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("resize", args)}
}

// Update This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) Update() *BitmapData{
    return &BitmapData{self.Object.Call("update")}
}

// Update1O This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) Update1O(x int) *BitmapData{
    return &BitmapData{self.Object.Call("update", x)}
}

// Update2O This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) Update2O(x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("update", x, y)}
}

// Update3O This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) Update3O(x int, y int, width int) *BitmapData{
    return &BitmapData{self.Object.Call("update", x, y, width)}
}

// Update4O This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) Update4O(x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("update", x, y, width, height)}
}

// UpdateI This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) UpdateI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("update", args)}
}

// ProcessPixelRGB Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
// The callback will be sent a color object with 6 properties: `{ r: number, g: number, b: number, a: number, color: number, rgba: string }`.
// Where r, g, b and a are integers between 0 and 255 representing the color component values for red, green, blue and alpha.
// The `color` property is an Int32 of the full color. Note the endianess of this will change per system.
// The `rgba` property is a CSS style rgba() string which can be used with context.fillStyle calls, among others.
// The callback will also be sent the pixels x and y coordinates respectively.
// The callback must return either `false`, in which case no change will be made to the pixel, or a new color object.
// If a new color object is returned the pixel will be set to the r, g, b and a color values given within it.
func (self *BitmapData) ProcessPixelRGB(callback interface{}, callbackContext interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("processPixelRGB", callback, callbackContext)}
}

// ProcessPixelRGB1O Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
// The callback will be sent a color object with 6 properties: `{ r: number, g: number, b: number, a: number, color: number, rgba: string }`.
// Where r, g, b and a are integers between 0 and 255 representing the color component values for red, green, blue and alpha.
// The `color` property is an Int32 of the full color. Note the endianess of this will change per system.
// The `rgba` property is a CSS style rgba() string which can be used with context.fillStyle calls, among others.
// The callback will also be sent the pixels x and y coordinates respectively.
// The callback must return either `false`, in which case no change will be made to the pixel, or a new color object.
// If a new color object is returned the pixel will be set to the r, g, b and a color values given within it.
func (self *BitmapData) ProcessPixelRGB1O(callback interface{}, callbackContext interface{}, x int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixelRGB", callback, callbackContext, x)}
}

// ProcessPixelRGB2O Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
// The callback will be sent a color object with 6 properties: `{ r: number, g: number, b: number, a: number, color: number, rgba: string }`.
// Where r, g, b and a are integers between 0 and 255 representing the color component values for red, green, blue and alpha.
// The `color` property is an Int32 of the full color. Note the endianess of this will change per system.
// The `rgba` property is a CSS style rgba() string which can be used with context.fillStyle calls, among others.
// The callback will also be sent the pixels x and y coordinates respectively.
// The callback must return either `false`, in which case no change will be made to the pixel, or a new color object.
// If a new color object is returned the pixel will be set to the r, g, b and a color values given within it.
func (self *BitmapData) ProcessPixelRGB2O(callback interface{}, callbackContext interface{}, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixelRGB", callback, callbackContext, x, y)}
}

// ProcessPixelRGB3O Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
// The callback will be sent a color object with 6 properties: `{ r: number, g: number, b: number, a: number, color: number, rgba: string }`.
// Where r, g, b and a are integers between 0 and 255 representing the color component values for red, green, blue and alpha.
// The `color` property is an Int32 of the full color. Note the endianess of this will change per system.
// The `rgba` property is a CSS style rgba() string which can be used with context.fillStyle calls, among others.
// The callback will also be sent the pixels x and y coordinates respectively.
// The callback must return either `false`, in which case no change will be made to the pixel, or a new color object.
// If a new color object is returned the pixel will be set to the r, g, b and a color values given within it.
func (self *BitmapData) ProcessPixelRGB3O(callback interface{}, callbackContext interface{}, x int, y int, width int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixelRGB", callback, callbackContext, x, y, width)}
}

// ProcessPixelRGB4O Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
// The callback will be sent a color object with 6 properties: `{ r: number, g: number, b: number, a: number, color: number, rgba: string }`.
// Where r, g, b and a are integers between 0 and 255 representing the color component values for red, green, blue and alpha.
// The `color` property is an Int32 of the full color. Note the endianess of this will change per system.
// The `rgba` property is a CSS style rgba() string which can be used with context.fillStyle calls, among others.
// The callback will also be sent the pixels x and y coordinates respectively.
// The callback must return either `false`, in which case no change will be made to the pixel, or a new color object.
// If a new color object is returned the pixel will be set to the r, g, b and a color values given within it.
func (self *BitmapData) ProcessPixelRGB4O(callback interface{}, callbackContext interface{}, x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixelRGB", callback, callbackContext, x, y, width, height)}
}

// ProcessPixelRGBI Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
// The callback will be sent a color object with 6 properties: `{ r: number, g: number, b: number, a: number, color: number, rgba: string }`.
// Where r, g, b and a are integers between 0 and 255 representing the color component values for red, green, blue and alpha.
// The `color` property is an Int32 of the full color. Note the endianess of this will change per system.
// The `rgba` property is a CSS style rgba() string which can be used with context.fillStyle calls, among others.
// The callback will also be sent the pixels x and y coordinates respectively.
// The callback must return either `false`, in which case no change will be made to the pixel, or a new color object.
// If a new color object is returned the pixel will be set to the r, g, b and a color values given within it.
func (self *BitmapData) ProcessPixelRGBI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("processPixelRGB", args)}
}

// ProcessPixel Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixel(callback interface{}, callbackContext interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", callback, callbackContext)}
}

// ProcessPixel1O Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixel1O(callback interface{}, callbackContext interface{}, x int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", callback, callbackContext, x)}
}

// ProcessPixel2O Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixel2O(callback interface{}, callbackContext interface{}, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", callback, callbackContext, x, y)}
}

// ProcessPixel3O Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixel3O(callback interface{}, callbackContext interface{}, x int, y int, width int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", callback, callbackContext, x, y, width)}
}

// ProcessPixel4O Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixel4O(callback interface{}, callbackContext interface{}, x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", callback, callbackContext, x, y, width, height)}
}

// ProcessPixelI Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixelI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", args)}
}

// ReplaceRGB Replaces all pixels matching one color with another. The color values are given as two sets of RGBA values.
// An optional region parameter controls if the replacement happens in just a specific area of the BitmapData or the entire thing.
func (self *BitmapData) ReplaceRGB(r1 int, g1 int, b1 int, a1 int, r2 int, g2 int, b2 int, a2 int) *BitmapData{
    return &BitmapData{self.Object.Call("replaceRGB", r1, g1, b1, a1, r2, g2, b2, a2)}
}

// ReplaceRGB1O Replaces all pixels matching one color with another. The color values are given as two sets of RGBA values.
// An optional region parameter controls if the replacement happens in just a specific area of the BitmapData or the entire thing.
func (self *BitmapData) ReplaceRGB1O(r1 int, g1 int, b1 int, a1 int, r2 int, g2 int, b2 int, a2 int, region *Rectangle) *BitmapData{
    return &BitmapData{self.Object.Call("replaceRGB", r1, g1, b1, a1, r2, g2, b2, a2, region)}
}

// ReplaceRGBI Replaces all pixels matching one color with another. The color values are given as two sets of RGBA values.
// An optional region parameter controls if the replacement happens in just a specific area of the BitmapData or the entire thing.
func (self *BitmapData) ReplaceRGBI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("replaceRGB", args)}
}

// SetHSL Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSL() *BitmapData{
    return &BitmapData{self.Object.Call("setHSL")}
}

// SetHSL1O Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSL1O(h int) *BitmapData{
    return &BitmapData{self.Object.Call("setHSL", h)}
}

// SetHSL2O Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSL2O(h int, s int) *BitmapData{
    return &BitmapData{self.Object.Call("setHSL", h, s)}
}

// SetHSL3O Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSL3O(h int, s int, l int) *BitmapData{
    return &BitmapData{self.Object.Call("setHSL", h, s, l)}
}

// SetHSL4O Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSL4O(h int, s int, l int, region *Rectangle) *BitmapData{
    return &BitmapData{self.Object.Call("setHSL", h, s, l, region)}
}

// SetHSLI Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSLI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("setHSL", args)}
}

// ShiftHSL Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSL() *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL")}
}

// ShiftHSL1O Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSL1O(h int) *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL", h)}
}

// ShiftHSL2O Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSL2O(h int, s int) *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL", h, s)}
}

// ShiftHSL3O Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSL3O(h int, s int, l int) *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL", h, s, l)}
}

// ShiftHSL4O Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSL4O(h int, s int, l int, region *Rectangle) *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL", h, s, l, region)}
}

// ShiftHSLI Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSLI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL", args)}
}

// SetPixel32 Sets the color of the given pixel to the specified red, green, blue and alpha values.
func (self *BitmapData) SetPixel32(x int, y int, red int, green int, blue int, alpha int) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel32", x, y, red, green, blue, alpha)}
}

// SetPixel321O Sets the color of the given pixel to the specified red, green, blue and alpha values.
func (self *BitmapData) SetPixel321O(x int, y int, red int, green int, blue int, alpha int, immediate bool) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel32", x, y, red, green, blue, alpha, immediate)}
}

// SetPixel32I Sets the color of the given pixel to the specified red, green, blue and alpha values.
func (self *BitmapData) SetPixel32I(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel32", args)}
}

// SetPixel Sets the color of the given pixel to the specified red, green and blue values.
func (self *BitmapData) SetPixel(x int, y int, red int, green int, blue int) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel", x, y, red, green, blue)}
}

// SetPixel1O Sets the color of the given pixel to the specified red, green and blue values.
func (self *BitmapData) SetPixel1O(x int, y int, red int, green int, blue int, immediate bool) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel", x, y, red, green, blue, immediate)}
}

// SetPixelI Sets the color of the given pixel to the specified red, green and blue values.
func (self *BitmapData) SetPixelI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel", args)}
}

// GetPixel Get the color of a specific pixel in the context into a color object.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixel(x int, y int) interface{}{
    return self.Object.Call("getPixel", x, y)
}

// GetPixel1O Get the color of a specific pixel in the context into a color object.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixel1O(x int, y int, out interface{}) interface{}{
    return self.Object.Call("getPixel", x, y, out)
}

// GetPixelI Get the color of a specific pixel in the context into a color object.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelI(args ...interface{}) interface{}{
    return self.Object.Call("getPixel", args)
}

// GetPixel32 Get the color of a specific pixel including its alpha value.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
// Note that on little-endian systems the format is 0xAABBGGRR and on big-endian the format is 0xRRGGBBAA.
func (self *BitmapData) GetPixel32(x int, y int) int{
    return self.Object.Call("getPixel32", x, y).Int()
}

// GetPixel32I Get the color of a specific pixel including its alpha value.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
// Note that on little-endian systems the format is 0xAABBGGRR and on big-endian the format is 0xRRGGBBAA.
func (self *BitmapData) GetPixel32I(args ...interface{}) int{
    return self.Object.Call("getPixel32", args).Int()
}

// GetPixelRGB Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGB(x int, y int) interface{}{
    return self.Object.Call("getPixelRGB", x, y)
}

// GetPixelRGB1O Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGB1O(x int, y int, out interface{}) interface{}{
    return self.Object.Call("getPixelRGB", x, y, out)
}

// GetPixelRGB2O Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGB2O(x int, y int, out interface{}, hsl bool) interface{}{
    return self.Object.Call("getPixelRGB", x, y, out, hsl)
}

// GetPixelRGB3O Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGB3O(x int, y int, out interface{}, hsl bool, hsv bool) interface{}{
    return self.Object.Call("getPixelRGB", x, y, out, hsl, hsv)
}

// GetPixelRGBI Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGBI(args ...interface{}) interface{}{
    return self.Object.Call("getPixelRGB", args)
}

// GetPixels Gets all the pixels from the region specified by the given Rectangle object.
func (self *BitmapData) GetPixels(rect *Rectangle) *ImageData{
    return &ImageData{self.Object.Call("getPixels", rect)}
}

// GetPixelsI Gets all the pixels from the region specified by the given Rectangle object.
func (self *BitmapData) GetPixelsI(args ...interface{}) *ImageData{
    return &ImageData{self.Object.Call("getPixels", args)}
}

// GetFirstPixel Scans the BitmapData, pixel by pixel, until it encounters a pixel that isn't transparent (i.e. has an alpha value > 0).
// It then stops scanning and returns an object containing the color of the pixel in r, g and b properties and the location in the x and y properties.
// 
// The direction parameter controls from which direction it should start the scan:
// 
// 0 = top to bottom
// 1 = bottom to top
// 2 = left to right
// 3 = right to left
func (self *BitmapData) GetFirstPixel() interface{}{
    return self.Object.Call("getFirstPixel")
}

// GetFirstPixel1O Scans the BitmapData, pixel by pixel, until it encounters a pixel that isn't transparent (i.e. has an alpha value > 0).
// It then stops scanning and returns an object containing the color of the pixel in r, g and b properties and the location in the x and y properties.
// 
// The direction parameter controls from which direction it should start the scan:
// 
// 0 = top to bottom
// 1 = bottom to top
// 2 = left to right
// 3 = right to left
func (self *BitmapData) GetFirstPixel1O(direction int) interface{}{
    return self.Object.Call("getFirstPixel", direction)
}

// GetFirstPixelI Scans the BitmapData, pixel by pixel, until it encounters a pixel that isn't transparent (i.e. has an alpha value > 0).
// It then stops scanning and returns an object containing the color of the pixel in r, g and b properties and the location in the x and y properties.
// 
// The direction parameter controls from which direction it should start the scan:
// 
// 0 = top to bottom
// 1 = bottom to top
// 2 = left to right
// 3 = right to left
func (self *BitmapData) GetFirstPixelI(args ...interface{}) interface{}{
    return self.Object.Call("getFirstPixel", args)
}

// GetBounds Scans the BitmapData and calculates the bounds. This is a rectangle that defines the extent of all non-transparent pixels.
// The rectangle returned will extend from the top-left of the image to the bottom-right, excluding transparent pixels.
func (self *BitmapData) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// GetBounds1O Scans the BitmapData and calculates the bounds. This is a rectangle that defines the extent of all non-transparent pixels.
// The rectangle returned will extend from the top-left of the image to the bottom-right, excluding transparent pixels.
func (self *BitmapData) GetBounds1O(rect *Rectangle) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", rect)}
}

// GetBoundsI Scans the BitmapData and calculates the bounds. This is a rectangle that defines the extent of all non-transparent pixels.
// The rectangle returned will extend from the top-left of the image to the bottom-right, excluding transparent pixels.
func (self *BitmapData) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// AddToWorld Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld() *Image{
    return &Image{self.Object.Call("addToWorld")}
}

// AddToWorld1O Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld1O(x int) *Image{
    return &Image{self.Object.Call("addToWorld", x)}
}

// AddToWorld2O Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld2O(x int, y int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y)}
}

// AddToWorld3O Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld3O(x int, y int, anchorX int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, anchorX)}
}

// AddToWorld4O Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld4O(x int, y int, anchorX int, anchorY int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, anchorX, anchorY)}
}

// AddToWorld5O Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld5O(x int, y int, anchorX int, anchorY int, scaleX int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, anchorX, anchorY, scaleX)}
}

// AddToWorld6O Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld6O(x int, y int, anchorX int, anchorY int, scaleX int, scaleY int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, anchorX, anchorY, scaleX, scaleY)}
}

// AddToWorldI Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorldI(args ...interface{}) *Image{
    return &Image{self.Object.Call("addToWorld", args)}
}

// Copy Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy() *BitmapData{
    return &BitmapData{self.Object.Call("copy")}
}

// Copy1O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy1O(source interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source)}
}

// Copy2O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy2O(source interface{}, x int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x)}
}

// Copy3O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy3O(source interface{}, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y)}
}

// Copy4O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy4O(source interface{}, x int, y int, width int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width)}
}

// Copy5O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy5O(source interface{}, x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height)}
}

// Copy6O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy6O(source interface{}, x int, y int, width int, height int, tx int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx)}
}

// Copy7O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy7O(source interface{}, x int, y int, width int, height int, tx int, ty int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty)}
}

// Copy8O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy8O(source interface{}, x int, y int, width int, height int, tx int, ty int, newWidth int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty, newWidth)}
}

// Copy9O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy9O(source interface{}, x int, y int, width int, height int, tx int, ty int, newWidth int, newHeight int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty, newWidth, newHeight)}
}

// Copy10O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy10O(source interface{}, x int, y int, width int, height int, tx int, ty int, newWidth int, newHeight int, rotate int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty, newWidth, newHeight, rotate)}
}

// Copy11O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy11O(source interface{}, x int, y int, width int, height int, tx int, ty int, newWidth int, newHeight int, rotate int, anchorX int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty, newWidth, newHeight, rotate, anchorX)}
}

// Copy12O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy12O(source interface{}, x int, y int, width int, height int, tx int, ty int, newWidth int, newHeight int, rotate int, anchorX int, anchorY int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty, newWidth, newHeight, rotate, anchorX, anchorY)}
}

// Copy13O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy13O(source interface{}, x int, y int, width int, height int, tx int, ty int, newWidth int, newHeight int, rotate int, anchorX int, anchorY int, scaleX int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty, newWidth, newHeight, rotate, anchorX, anchorY, scaleX)}
}

// Copy14O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy14O(source interface{}, x int, y int, width int, height int, tx int, ty int, newWidth int, newHeight int, rotate int, anchorX int, anchorY int, scaleX int, scaleY int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty, newWidth, newHeight, rotate, anchorX, anchorY, scaleX, scaleY)}
}

// Copy15O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy15O(source interface{}, x int, y int, width int, height int, tx int, ty int, newWidth int, newHeight int, rotate int, anchorX int, anchorY int, scaleX int, scaleY int, alpha int) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty, newWidth, newHeight, rotate, anchorX, anchorY, scaleX, scaleY, alpha)}
}

// Copy16O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy16O(source interface{}, x int, y int, width int, height int, tx int, ty int, newWidth int, newHeight int, rotate int, anchorX int, anchorY int, scaleX int, scaleY int, alpha int, blendMode string) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty, newWidth, newHeight, rotate, anchorX, anchorY, scaleX, scaleY, alpha, blendMode)}
}

// Copy17O Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) Copy17O(source interface{}, x int, y int, width int, height int, tx int, ty int, newWidth int, newHeight int, rotate int, anchorX int, anchorY int, scaleX int, scaleY int, alpha int, blendMode string, roundPx bool) *BitmapData{
    return &BitmapData{self.Object.Call("copy", source, x, y, width, height, tx, ty, newWidth, newHeight, rotate, anchorX, anchorY, scaleX, scaleY, alpha, blendMode, roundPx)}
}

// CopyI Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
// 
// You can optionally resize, translate, rotate, scale, alpha or blend as it's drawn.
// 
// All rotation, scaling and drawing takes place around the regions center point by default, but can be changed with the anchor parameters.
// 
// Note that the source image can also be this BitmapData, which can create some interesting effects.
// 
// This method has a lot of parameters for maximum control.
// You can use the more friendly methods like `copyRect` and `draw` to avoid having to remember them all.
// 
// You may prefer to use `copyTransform` if you're simply trying to draw a Sprite to this BitmapData,
// and don't wish to translate, scale or rotate it from its original values.
func (self *BitmapData) CopyI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("copy", args)}
}

// CopyTransform Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransform() *BitmapData{
    return &BitmapData{self.Object.Call("copyTransform")}
}

// CopyTransform1O Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransform1O(source interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("copyTransform", source)}
}

// CopyTransform2O Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransform2O(source interface{}, blendMode string) *BitmapData{
    return &BitmapData{self.Object.Call("copyTransform", source, blendMode)}
}

// CopyTransform3O Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransform3O(source interface{}, blendMode string, roundPx bool) *BitmapData{
    return &BitmapData{self.Object.Call("copyTransform", source, blendMode, roundPx)}
}

// CopyTransformI Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransformI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("copyTransform", args)}
}

// CopyRect Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRect(source interface{}, area *Rectangle, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("copyRect", source, area, x, y)}
}

// CopyRect1O Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRect1O(source interface{}, area *Rectangle, x int, y int, alpha int) *BitmapData{
    return &BitmapData{self.Object.Call("copyRect", source, area, x, y, alpha)}
}

// CopyRect2O Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRect2O(source interface{}, area *Rectangle, x int, y int, alpha int, blendMode string) *BitmapData{
    return &BitmapData{self.Object.Call("copyRect", source, area, x, y, alpha, blendMode)}
}

// CopyRect3O Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRect3O(source interface{}, area *Rectangle, x int, y int, alpha int, blendMode string, roundPx bool) *BitmapData{
    return &BitmapData{self.Object.Call("copyRect", source, area, x, y, alpha, blendMode, roundPx)}
}

// CopyRectI Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRectI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("copyRect", args)}
}

// Draw Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
// You can use the optional width and height values to 'stretch' the sprite as it is drawn. This uses drawImage stretching, not scaling.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `draw`.
func (self *BitmapData) Draw(source interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("draw", source)}
}

// Draw1O Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
// You can use the optional width and height values to 'stretch' the sprite as it is drawn. This uses drawImage stretching, not scaling.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `draw`.
func (self *BitmapData) Draw1O(source interface{}, x int) *BitmapData{
    return &BitmapData{self.Object.Call("draw", source, x)}
}

// Draw2O Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
// You can use the optional width and height values to 'stretch' the sprite as it is drawn. This uses drawImage stretching, not scaling.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `draw`.
func (self *BitmapData) Draw2O(source interface{}, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("draw", source, x, y)}
}

// Draw3O Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
// You can use the optional width and height values to 'stretch' the sprite as it is drawn. This uses drawImage stretching, not scaling.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `draw`.
func (self *BitmapData) Draw3O(source interface{}, x int, y int, width int) *BitmapData{
    return &BitmapData{self.Object.Call("draw", source, x, y, width)}
}

// Draw4O Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
// You can use the optional width and height values to 'stretch' the sprite as it is drawn. This uses drawImage stretching, not scaling.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `draw`.
func (self *BitmapData) Draw4O(source interface{}, x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("draw", source, x, y, width, height)}
}

// Draw5O Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
// You can use the optional width and height values to 'stretch' the sprite as it is drawn. This uses drawImage stretching, not scaling.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `draw`.
func (self *BitmapData) Draw5O(source interface{}, x int, y int, width int, height int, blendMode string) *BitmapData{
    return &BitmapData{self.Object.Call("draw", source, x, y, width, height, blendMode)}
}

// Draw6O Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
// You can use the optional width and height values to 'stretch' the sprite as it is drawn. This uses drawImage stretching, not scaling.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `draw`.
func (self *BitmapData) Draw6O(source interface{}, x int, y int, width int, height int, blendMode string, roundPx bool) *BitmapData{
    return &BitmapData{self.Object.Call("draw", source, x, y, width, height, blendMode, roundPx)}
}

// DrawI Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
// You can use the optional width and height values to 'stretch' the sprite as it is drawn. This uses drawImage stretching, not scaling.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `draw`.
func (self *BitmapData) DrawI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("draw", args)}
}

// DrawGroup Draws the immediate children of a Phaser.Group to this BitmapData.
// 
// It's perfectly valid to pass in `game.world` as the Group, and it will iterate through the entire display list.
// 
// Children are drawn _only_ if they have their `exists` property set to `true`, and have image, or RenderTexture, based Textures.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `drawGroup`.
func (self *BitmapData) DrawGroup(group *Group) *BitmapData{
    return &BitmapData{self.Object.Call("drawGroup", group)}
}

// DrawGroup1O Draws the immediate children of a Phaser.Group to this BitmapData.
// 
// It's perfectly valid to pass in `game.world` as the Group, and it will iterate through the entire display list.
// 
// Children are drawn _only_ if they have their `exists` property set to `true`, and have image, or RenderTexture, based Textures.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `drawGroup`.
func (self *BitmapData) DrawGroup1O(group *Group, blendMode string) *BitmapData{
    return &BitmapData{self.Object.Call("drawGroup", group, blendMode)}
}

// DrawGroup2O Draws the immediate children of a Phaser.Group to this BitmapData.
// 
// It's perfectly valid to pass in `game.world` as the Group, and it will iterate through the entire display list.
// 
// Children are drawn _only_ if they have their `exists` property set to `true`, and have image, or RenderTexture, based Textures.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `drawGroup`.
func (self *BitmapData) DrawGroup2O(group *Group, blendMode string, roundPx bool) *BitmapData{
    return &BitmapData{self.Object.Call("drawGroup", group, blendMode, roundPx)}
}

// DrawGroupI Draws the immediate children of a Phaser.Group to this BitmapData.
// 
// It's perfectly valid to pass in `game.world` as the Group, and it will iterate through the entire display list.
// 
// Children are drawn _only_ if they have their `exists` property set to `true`, and have image, or RenderTexture, based Textures.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData they won't be visible.
// When drawing it will take into account the rotation, scale, scaleMode, alpha and tint values.
// 
// Note: You should ensure that at least 1 full update has taken place before calling this, 
// otherwise the objects are likely to render incorrectly, if at all.
// You can trigger an update yourself by calling `stage.updateTransform()` before calling `drawGroup`.
func (self *BitmapData) DrawGroupI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("drawGroup", args)}
}

// DrawGroupProxy A proxy for drawGroup that handles child iteration for more complex Game Objects.
func (self *BitmapData) DrawGroupProxy(child interface{}) {
    self.Object.Call("drawGroupProxy", child)
}

// DrawGroupProxy1O A proxy for drawGroup that handles child iteration for more complex Game Objects.
func (self *BitmapData) DrawGroupProxy1O(child interface{}, blendMode string) {
    self.Object.Call("drawGroupProxy", child, blendMode)
}

// DrawGroupProxy2O A proxy for drawGroup that handles child iteration for more complex Game Objects.
func (self *BitmapData) DrawGroupProxy2O(child interface{}, blendMode string, roundPx bool) {
    self.Object.Call("drawGroupProxy", child, blendMode, roundPx)
}

// DrawGroupProxyI A proxy for drawGroup that handles child iteration for more complex Game Objects.
func (self *BitmapData) DrawGroupProxyI(args ...interface{}) {
    self.Object.Call("drawGroupProxy", args)
}

// DrawFull Draws the Game Object or Group to this BitmapData and then recursively iterates through all of its children.
// 
// If a child has an `exists` property then it (and its children) will be only be drawn if exists is `true`.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData 
// they won't be drawn. Depending on your requirements you may need to resize the BitmapData in advance to match the 
// bounds of the top-level Game Object.
// 
// When drawing it will take into account the child's world rotation, scale and alpha values.
// 
// It's perfectly valid to pass in `game.world` as the parent object, and it will iterate through the entire display list.
// 
// Note: If you are trying to grab your entire game at the start of a State then you should ensure that at least 1 full update
// has taken place before doing so, otherwise all of the objects will render with incorrect positions and scales. You can 
// trigger an update yourself by calling `stage.updateTransform()` before calling `drawFull`.
func (self *BitmapData) DrawFull(parent interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("drawFull", parent)}
}

// DrawFull1O Draws the Game Object or Group to this BitmapData and then recursively iterates through all of its children.
// 
// If a child has an `exists` property then it (and its children) will be only be drawn if exists is `true`.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData 
// they won't be drawn. Depending on your requirements you may need to resize the BitmapData in advance to match the 
// bounds of the top-level Game Object.
// 
// When drawing it will take into account the child's world rotation, scale and alpha values.
// 
// It's perfectly valid to pass in `game.world` as the parent object, and it will iterate through the entire display list.
// 
// Note: If you are trying to grab your entire game at the start of a State then you should ensure that at least 1 full update
// has taken place before doing so, otherwise all of the objects will render with incorrect positions and scales. You can 
// trigger an update yourself by calling `stage.updateTransform()` before calling `drawFull`.
func (self *BitmapData) DrawFull1O(parent interface{}, blendMode string) *BitmapData{
    return &BitmapData{self.Object.Call("drawFull", parent, blendMode)}
}

// DrawFull2O Draws the Game Object or Group to this BitmapData and then recursively iterates through all of its children.
// 
// If a child has an `exists` property then it (and its children) will be only be drawn if exists is `true`.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData 
// they won't be drawn. Depending on your requirements you may need to resize the BitmapData in advance to match the 
// bounds of the top-level Game Object.
// 
// When drawing it will take into account the child's world rotation, scale and alpha values.
// 
// It's perfectly valid to pass in `game.world` as the parent object, and it will iterate through the entire display list.
// 
// Note: If you are trying to grab your entire game at the start of a State then you should ensure that at least 1 full update
// has taken place before doing so, otherwise all of the objects will render with incorrect positions and scales. You can 
// trigger an update yourself by calling `stage.updateTransform()` before calling `drawFull`.
func (self *BitmapData) DrawFull2O(parent interface{}, blendMode string, roundPx bool) *BitmapData{
    return &BitmapData{self.Object.Call("drawFull", parent, blendMode, roundPx)}
}

// DrawFullI Draws the Game Object or Group to this BitmapData and then recursively iterates through all of its children.
// 
// If a child has an `exists` property then it (and its children) will be only be drawn if exists is `true`.
// 
// The children will be drawn at their `x` and `y` world space coordinates. If this is outside the bounds of the BitmapData 
// they won't be drawn. Depending on your requirements you may need to resize the BitmapData in advance to match the 
// bounds of the top-level Game Object.
// 
// When drawing it will take into account the child's world rotation, scale and alpha values.
// 
// It's perfectly valid to pass in `game.world` as the parent object, and it will iterate through the entire display list.
// 
// Note: If you are trying to grab your entire game at the start of a State then you should ensure that at least 1 full update
// has taken place before doing so, otherwise all of the objects will render with incorrect positions and scales. You can 
// trigger an update yourself by calling `stage.updateTransform()` before calling `drawFull`.
func (self *BitmapData) DrawFullI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("drawFull", args)}
}

// Shadow Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) Shadow(color string) *BitmapData{
    return &BitmapData{self.Object.Call("shadow", color)}
}

// Shadow1O Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) Shadow1O(color string, blur int) *BitmapData{
    return &BitmapData{self.Object.Call("shadow", color, blur)}
}

// Shadow2O Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) Shadow2O(color string, blur int, x int) *BitmapData{
    return &BitmapData{self.Object.Call("shadow", color, blur, x)}
}

// Shadow3O Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) Shadow3O(color string, blur int, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("shadow", color, blur, x, y)}
}

// ShadowI Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) ShadowI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("shadow", args)}
}

// AlphaMask Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMask(source interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("alphaMask", source)}
}

// AlphaMask1O Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMask1O(source interface{}, mask interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("alphaMask", source, mask)}
}

// AlphaMask2O Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMask2O(source interface{}, mask interface{}, sourceRect *Rectangle) *BitmapData{
    return &BitmapData{self.Object.Call("alphaMask", source, mask, sourceRect)}
}

// AlphaMask3O Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMask3O(source interface{}, mask interface{}, sourceRect *Rectangle, maskRect *Rectangle) *BitmapData{
    return &BitmapData{self.Object.Call("alphaMask", source, mask, sourceRect, maskRect)}
}

// AlphaMaskI Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMaskI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("alphaMask", args)}
}

// Extract Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
// The original BitmapData remains unchanged.
// The destination BitmapData must be large enough to receive all of the pixels that are scanned unless the 'resize' parameter is true.
// Although the destination BitmapData is returned from this method, it's actually modified directly in place, meaning this call is perfectly valid:
// `picture.extract(mask, r, g, b)`
// You can specify optional r2, g2, b2 color values. If given the pixel written to the destination bitmap will be of the r2, g2, b2 color.
// If not given it will be written as the same color it was extracted. You can provide one or more alternative colors, allowing you to tint
// the color during extraction.
func (self *BitmapData) Extract(destination *BitmapData, r int, g int, b int) *BitmapData{
    return &BitmapData{self.Object.Call("extract", destination, r, g, b)}
}

// Extract1O Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
// The original BitmapData remains unchanged.
// The destination BitmapData must be large enough to receive all of the pixels that are scanned unless the 'resize' parameter is true.
// Although the destination BitmapData is returned from this method, it's actually modified directly in place, meaning this call is perfectly valid:
// `picture.extract(mask, r, g, b)`
// You can specify optional r2, g2, b2 color values. If given the pixel written to the destination bitmap will be of the r2, g2, b2 color.
// If not given it will be written as the same color it was extracted. You can provide one or more alternative colors, allowing you to tint
// the color during extraction.
func (self *BitmapData) Extract1O(destination *BitmapData, r int, g int, b int, a int) *BitmapData{
    return &BitmapData{self.Object.Call("extract", destination, r, g, b, a)}
}

// Extract2O Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
// The original BitmapData remains unchanged.
// The destination BitmapData must be large enough to receive all of the pixels that are scanned unless the 'resize' parameter is true.
// Although the destination BitmapData is returned from this method, it's actually modified directly in place, meaning this call is perfectly valid:
// `picture.extract(mask, r, g, b)`
// You can specify optional r2, g2, b2 color values. If given the pixel written to the destination bitmap will be of the r2, g2, b2 color.
// If not given it will be written as the same color it was extracted. You can provide one or more alternative colors, allowing you to tint
// the color during extraction.
func (self *BitmapData) Extract2O(destination *BitmapData, r int, g int, b int, a int, resize bool) *BitmapData{
    return &BitmapData{self.Object.Call("extract", destination, r, g, b, a, resize)}
}

// Extract3O Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
// The original BitmapData remains unchanged.
// The destination BitmapData must be large enough to receive all of the pixels that are scanned unless the 'resize' parameter is true.
// Although the destination BitmapData is returned from this method, it's actually modified directly in place, meaning this call is perfectly valid:
// `picture.extract(mask, r, g, b)`
// You can specify optional r2, g2, b2 color values. If given the pixel written to the destination bitmap will be of the r2, g2, b2 color.
// If not given it will be written as the same color it was extracted. You can provide one or more alternative colors, allowing you to tint
// the color during extraction.
func (self *BitmapData) Extract3O(destination *BitmapData, r int, g int, b int, a int, resize bool, r2 int) *BitmapData{
    return &BitmapData{self.Object.Call("extract", destination, r, g, b, a, resize, r2)}
}

// Extract4O Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
// The original BitmapData remains unchanged.
// The destination BitmapData must be large enough to receive all of the pixels that are scanned unless the 'resize' parameter is true.
// Although the destination BitmapData is returned from this method, it's actually modified directly in place, meaning this call is perfectly valid:
// `picture.extract(mask, r, g, b)`
// You can specify optional r2, g2, b2 color values. If given the pixel written to the destination bitmap will be of the r2, g2, b2 color.
// If not given it will be written as the same color it was extracted. You can provide one or more alternative colors, allowing you to tint
// the color during extraction.
func (self *BitmapData) Extract4O(destination *BitmapData, r int, g int, b int, a int, resize bool, r2 int, g2 int) *BitmapData{
    return &BitmapData{self.Object.Call("extract", destination, r, g, b, a, resize, r2, g2)}
}

// Extract5O Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
// The original BitmapData remains unchanged.
// The destination BitmapData must be large enough to receive all of the pixels that are scanned unless the 'resize' parameter is true.
// Although the destination BitmapData is returned from this method, it's actually modified directly in place, meaning this call is perfectly valid:
// `picture.extract(mask, r, g, b)`
// You can specify optional r2, g2, b2 color values. If given the pixel written to the destination bitmap will be of the r2, g2, b2 color.
// If not given it will be written as the same color it was extracted. You can provide one or more alternative colors, allowing you to tint
// the color during extraction.
func (self *BitmapData) Extract5O(destination *BitmapData, r int, g int, b int, a int, resize bool, r2 int, g2 int, b2 int) *BitmapData{
    return &BitmapData{self.Object.Call("extract", destination, r, g, b, a, resize, r2, g2, b2)}
}

// ExtractI Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
// The original BitmapData remains unchanged.
// The destination BitmapData must be large enough to receive all of the pixels that are scanned unless the 'resize' parameter is true.
// Although the destination BitmapData is returned from this method, it's actually modified directly in place, meaning this call is perfectly valid:
// `picture.extract(mask, r, g, b)`
// You can specify optional r2, g2, b2 color values. If given the pixel written to the destination bitmap will be of the r2, g2, b2 color.
// If not given it will be written as the same color it was extracted. You can provide one or more alternative colors, allowing you to tint
// the color during extraction.
func (self *BitmapData) ExtractI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("extract", args)}
}

// Rect Draws a filled Rectangle to the BitmapData at the given x, y coordinates and width / height in size.
func (self *BitmapData) Rect(x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("rect", x, y, width, height)}
}

// Rect1O Draws a filled Rectangle to the BitmapData at the given x, y coordinates and width / height in size.
func (self *BitmapData) Rect1O(x int, y int, width int, height int, fillStyle string) *BitmapData{
    return &BitmapData{self.Object.Call("rect", x, y, width, height, fillStyle)}
}

// RectI Draws a filled Rectangle to the BitmapData at the given x, y coordinates and width / height in size.
func (self *BitmapData) RectI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("rect", args)}
}

// Text Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) Text(text string, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("text", text, x, y)}
}

// Text1O Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) Text1O(text string, x int, y int, font string) *BitmapData{
    return &BitmapData{self.Object.Call("text", text, x, y, font)}
}

// Text2O Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) Text2O(text string, x int, y int, font string, color string) *BitmapData{
    return &BitmapData{self.Object.Call("text", text, x, y, font, color)}
}

// Text3O Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) Text3O(text string, x int, y int, font string, color string, shadow bool) *BitmapData{
    return &BitmapData{self.Object.Call("text", text, x, y, font, color, shadow)}
}

// TextI Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) TextI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("text", args)}
}

// Circle Draws a filled Circle to the BitmapData at the given x, y coordinates and radius in size.
func (self *BitmapData) Circle(x int, y int, radius int) *BitmapData{
    return &BitmapData{self.Object.Call("circle", x, y, radius)}
}

// Circle1O Draws a filled Circle to the BitmapData at the given x, y coordinates and radius in size.
func (self *BitmapData) Circle1O(x int, y int, radius int, fillStyle string) *BitmapData{
    return &BitmapData{self.Object.Call("circle", x, y, radius, fillStyle)}
}

// CircleI Draws a filled Circle to the BitmapData at the given x, y coordinates and radius in size.
func (self *BitmapData) CircleI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("circle", args)}
}

// Line Draws a line between the coordinates given in the color and thickness specified.
func (self *BitmapData) Line(x1 int, y1 int, x2 int, y2 int) *BitmapData{
    return &BitmapData{self.Object.Call("line", x1, y1, x2, y2)}
}

// Line1O Draws a line between the coordinates given in the color and thickness specified.
func (self *BitmapData) Line1O(x1 int, y1 int, x2 int, y2 int, color string) *BitmapData{
    return &BitmapData{self.Object.Call("line", x1, y1, x2, y2, color)}
}

// Line2O Draws a line between the coordinates given in the color and thickness specified.
func (self *BitmapData) Line2O(x1 int, y1 int, x2 int, y2 int, color string, width int) *BitmapData{
    return &BitmapData{self.Object.Call("line", x1, y1, x2, y2, color, width)}
}

// LineI Draws a line between the coordinates given in the color and thickness specified.
func (self *BitmapData) LineI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("line", args)}
}

// TextureLine Takes the given Line object and image and renders it to this BitmapData as a repeating texture line.
func (self *BitmapData) TextureLine(line *Line, image interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("textureLine", line, image)}
}

// TextureLine1O Takes the given Line object and image and renders it to this BitmapData as a repeating texture line.
func (self *BitmapData) TextureLine1O(line *Line, image interface{}, repeat string) *BitmapData{
    return &BitmapData{self.Object.Call("textureLine", line, image, repeat)}
}

// TextureLineI Takes the given Line object and image and renders it to this BitmapData as a repeating texture line.
func (self *BitmapData) TextureLineI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("textureLine", args)}
}

// Render If the game is running in WebGL this will push the texture up to the GPU if it's dirty.
// This is called automatically if the BitmapData is being used by a Sprite, otherwise you need to remember to call it in your render function.
// If you wish to suppress this functionality set BitmapData.disableTextureUpload to `true`.
func (self *BitmapData) Render() *BitmapData{
    return &BitmapData{self.Object.Call("render")}
}

// RenderI If the game is running in WebGL this will push the texture up to the GPU if it's dirty.
// This is called automatically if the BitmapData is being used by a Sprite, otherwise you need to remember to call it in your render function.
// If you wish to suppress this functionality set BitmapData.disableTextureUpload to `true`.
func (self *BitmapData) RenderI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("render", args)}
}

// Destroy Destroys this BitmapData and puts the canvas it was using back into the canvas pool for re-use.
func (self *BitmapData) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this BitmapData and puts the canvas it was using back into the canvas pool for re-use.
func (self *BitmapData) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// BlendReset Resets the blend mode (effectively sets it to 'source-over')
func (self *BitmapData) BlendReset() *BitmapData{
    return &BitmapData{self.Object.Call("blendReset")}
}

// BlendResetI Resets the blend mode (effectively sets it to 'source-over')
func (self *BitmapData) BlendResetI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendReset", args)}
}

// BlendSourceOver Sets the blend mode to 'source-over'
func (self *BitmapData) BlendSourceOver() *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceOver")}
}

// BlendSourceOverI Sets the blend mode to 'source-over'
func (self *BitmapData) BlendSourceOverI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceOver", args)}
}

// BlendSourceIn Sets the blend mode to 'source-in'
func (self *BitmapData) BlendSourceIn() *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceIn")}
}

// BlendSourceInI Sets the blend mode to 'source-in'
func (self *BitmapData) BlendSourceInI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceIn", args)}
}

// BlendSourceOut Sets the blend mode to 'source-out'
func (self *BitmapData) BlendSourceOut() *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceOut")}
}

// BlendSourceOutI Sets the blend mode to 'source-out'
func (self *BitmapData) BlendSourceOutI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceOut", args)}
}

// BlendSourceAtop Sets the blend mode to 'source-atop'
func (self *BitmapData) BlendSourceAtop() *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceAtop")}
}

// BlendSourceAtopI Sets the blend mode to 'source-atop'
func (self *BitmapData) BlendSourceAtopI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceAtop", args)}
}

// BlendDestinationOver Sets the blend mode to 'destination-over'
func (self *BitmapData) BlendDestinationOver() *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationOver")}
}

// BlendDestinationOverI Sets the blend mode to 'destination-over'
func (self *BitmapData) BlendDestinationOverI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationOver", args)}
}

// BlendDestinationIn Sets the blend mode to 'destination-in'
func (self *BitmapData) BlendDestinationIn() *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationIn")}
}

// BlendDestinationInI Sets the blend mode to 'destination-in'
func (self *BitmapData) BlendDestinationInI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationIn", args)}
}

// BlendDestinationOut Sets the blend mode to 'destination-out'
func (self *BitmapData) BlendDestinationOut() *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationOut")}
}

// BlendDestinationOutI Sets the blend mode to 'destination-out'
func (self *BitmapData) BlendDestinationOutI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationOut", args)}
}

// BlendDestinationAtop Sets the blend mode to 'destination-atop'
func (self *BitmapData) BlendDestinationAtop() *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationAtop")}
}

// BlendDestinationAtopI Sets the blend mode to 'destination-atop'
func (self *BitmapData) BlendDestinationAtopI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationAtop", args)}
}

// BlendXor Sets the blend mode to 'xor'
func (self *BitmapData) BlendXor() *BitmapData{
    return &BitmapData{self.Object.Call("blendXor")}
}

// BlendXorI Sets the blend mode to 'xor'
func (self *BitmapData) BlendXorI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendXor", args)}
}

// BlendAdd Sets the blend mode to 'lighter'
func (self *BitmapData) BlendAdd() *BitmapData{
    return &BitmapData{self.Object.Call("blendAdd")}
}

// BlendAddI Sets the blend mode to 'lighter'
func (self *BitmapData) BlendAddI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendAdd", args)}
}

// BlendMultiply Sets the blend mode to 'multiply'
func (self *BitmapData) BlendMultiply() *BitmapData{
    return &BitmapData{self.Object.Call("blendMultiply")}
}

// BlendMultiplyI Sets the blend mode to 'multiply'
func (self *BitmapData) BlendMultiplyI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendMultiply", args)}
}

// BlendScreen Sets the blend mode to 'screen'
func (self *BitmapData) BlendScreen() *BitmapData{
    return &BitmapData{self.Object.Call("blendScreen")}
}

// BlendScreenI Sets the blend mode to 'screen'
func (self *BitmapData) BlendScreenI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendScreen", args)}
}

// BlendOverlay Sets the blend mode to 'overlay'
func (self *BitmapData) BlendOverlay() *BitmapData{
    return &BitmapData{self.Object.Call("blendOverlay")}
}

// BlendOverlayI Sets the blend mode to 'overlay'
func (self *BitmapData) BlendOverlayI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendOverlay", args)}
}

// BlendDarken Sets the blend mode to 'darken'
func (self *BitmapData) BlendDarken() *BitmapData{
    return &BitmapData{self.Object.Call("blendDarken")}
}

// BlendDarkenI Sets the blend mode to 'darken'
func (self *BitmapData) BlendDarkenI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDarken", args)}
}

// BlendLighten Sets the blend mode to 'lighten'
func (self *BitmapData) BlendLighten() *BitmapData{
    return &BitmapData{self.Object.Call("blendLighten")}
}

// BlendLightenI Sets the blend mode to 'lighten'
func (self *BitmapData) BlendLightenI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendLighten", args)}
}

// BlendColorDodge Sets the blend mode to 'color-dodge'
func (self *BitmapData) BlendColorDodge() *BitmapData{
    return &BitmapData{self.Object.Call("blendColorDodge")}
}

// BlendColorDodgeI Sets the blend mode to 'color-dodge'
func (self *BitmapData) BlendColorDodgeI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendColorDodge", args)}
}

// BlendColorBurn Sets the blend mode to 'color-burn'
func (self *BitmapData) BlendColorBurn() *BitmapData{
    return &BitmapData{self.Object.Call("blendColorBurn")}
}

// BlendColorBurnI Sets the blend mode to 'color-burn'
func (self *BitmapData) BlendColorBurnI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendColorBurn", args)}
}

// BlendHardLight Sets the blend mode to 'hard-light'
func (self *BitmapData) BlendHardLight() *BitmapData{
    return &BitmapData{self.Object.Call("blendHardLight")}
}

// BlendHardLightI Sets the blend mode to 'hard-light'
func (self *BitmapData) BlendHardLightI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendHardLight", args)}
}

// BlendSoftLight Sets the blend mode to 'soft-light'
func (self *BitmapData) BlendSoftLight() *BitmapData{
    return &BitmapData{self.Object.Call("blendSoftLight")}
}

// BlendSoftLightI Sets the blend mode to 'soft-light'
func (self *BitmapData) BlendSoftLightI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSoftLight", args)}
}

// BlendDifference Sets the blend mode to 'difference'
func (self *BitmapData) BlendDifference() *BitmapData{
    return &BitmapData{self.Object.Call("blendDifference")}
}

// BlendDifferenceI Sets the blend mode to 'difference'
func (self *BitmapData) BlendDifferenceI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDifference", args)}
}

// BlendExclusion Sets the blend mode to 'exclusion'
func (self *BitmapData) BlendExclusion() *BitmapData{
    return &BitmapData{self.Object.Call("blendExclusion")}
}

// BlendExclusionI Sets the blend mode to 'exclusion'
func (self *BitmapData) BlendExclusionI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendExclusion", args)}
}

// BlendHue Sets the blend mode to 'hue'
func (self *BitmapData) BlendHue() *BitmapData{
    return &BitmapData{self.Object.Call("blendHue")}
}

// BlendHueI Sets the blend mode to 'hue'
func (self *BitmapData) BlendHueI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendHue", args)}
}

// BlendSaturation Sets the blend mode to 'saturation'
func (self *BitmapData) BlendSaturation() *BitmapData{
    return &BitmapData{self.Object.Call("blendSaturation")}
}

// BlendSaturationI Sets the blend mode to 'saturation'
func (self *BitmapData) BlendSaturationI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSaturation", args)}
}

// BlendColor Sets the blend mode to 'color'
func (self *BitmapData) BlendColor() *BitmapData{
    return &BitmapData{self.Object.Call("blendColor")}
}

// BlendColorI Sets the blend mode to 'color'
func (self *BitmapData) BlendColorI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendColor", args)}
}

// BlendLuminosity Sets the blend mode to 'luminosity'
func (self *BitmapData) BlendLuminosity() *BitmapData{
    return &BitmapData{self.Object.Call("blendLuminosity")}
}

// BlendLuminosityI Sets the blend mode to 'luminosity'
func (self *BitmapData) BlendLuminosityI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendLuminosity", args)}
}

// GetTransform Gets a JavaScript object that has 6 properties set that are used by BitmapData in a transform.
func (self *BitmapData) GetTransform(translateX int, translateY int, scaleX int, scaleY int, skewX int, skewY int) interface{}{
    return self.Object.Call("getTransform", translateX, translateY, scaleX, scaleY, skewX, skewY)
}

// GetTransformI Gets a JavaScript object that has 6 properties set that are used by BitmapData in a transform.
func (self *BitmapData) GetTransformI(args ...interface{}) interface{}{
    return self.Object.Call("getTransform", args)
}

