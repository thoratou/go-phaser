// Automatic generation for Phaser.BitmapData
// generated file BitmapData.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
type BitmapData struct {
    *js.Object
}


// A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
func NewBitmapData(game *Game, key string) *BitmapData {
    return &BitmapData{js.Global.Get("Phaser").Get("BitmapData").New(game, key)}
}

// A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
func NewBitmapData1O(game *Game, key string, width int) *BitmapData {
    return &BitmapData{js.Global.Get("Phaser").Get("BitmapData").New(game, key, width)}
}

// A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
func NewBitmapData2O(game *Game, key string, width int, height int) *BitmapData {
    return &BitmapData{js.Global.Get("Phaser").Get("BitmapData").New(game, key, width, height)}
}

// A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
func NewBitmapData3O(game *Game, key string, width int, height int, skipPool bool) *BitmapData {
    return &BitmapData{js.Global.Get("Phaser").Get("BitmapData").New(game, key, width, height, skipPool)}
}

// A BitmapData object contains a Canvas element to which you can draw anything you like via normal Canvas context operations.
// A single BitmapData can be used as the texture for one or many Images / Sprites. 
// So if you need to dynamically create a Sprite texture then they are a good choice.
// 
// Important note: Every BitmapData creates its own Canvas element. Because BitmapData's are now Game Objects themselves, and don't
// live on the display list, they are NOT automatically cleared when you change State. Therefore you _must_ call BitmapData.destroy
// in your State's shutdown method if you wish to free-up the resources the BitmapData used, it will not happen for you.
func NewBitmapDataI(args ...interface{}) *BitmapData {
    return &BitmapData{js.Global.Get("Phaser").Get("BitmapData").New(args)}
}



// A reference to the currently running game.
func (self *BitmapData) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *BitmapData) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The key of the BitmapData in the Cache, if stored there.
func (self *BitmapData) Key() string{
    return self.Object.Get("key").String()
}

// The key of the BitmapData in the Cache, if stored there.
func (self *BitmapData) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// The width of the BitmapData in pixels.
func (self *BitmapData) Width() int{
    return self.Object.Get("width").Int()
}

// The width of the BitmapData in pixels.
func (self *BitmapData) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the BitmapData in pixels.
func (self *BitmapData) Height() int{
    return self.Object.Get("height").Int()
}

// The height of the BitmapData in pixels.
func (self *BitmapData) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The canvas to which this BitmapData draws.
func (self *BitmapData) Canvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Get("canvas"))
}

// The canvas to which this BitmapData draws.
func (self *BitmapData) SetCanvasA(member dom.HTMLCanvasElement) {
    self.Object.Set("canvas", member)
}

// The 2d context of the canvas.
func (self *BitmapData) Context() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("context"))
}

// The 2d context of the canvas.
func (self *BitmapData) SetContextA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("context", member)
}

// A reference to BitmapData.context.
func (self *BitmapData) Ctx() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Get("ctx"))
}

// A reference to BitmapData.context.
func (self *BitmapData) SetCtxA(member dom.CanvasRenderingContext2D) {
    self.Object.Set("ctx", member)
}

// The context property needed for smoothing this Canvas.
func (self *BitmapData) SmoothProperty() string{
    return self.Object.Get("smoothProperty").String()
}

// The context property needed for smoothing this Canvas.
func (self *BitmapData) SetSmoothPropertyA(member string) {
    self.Object.Set("smoothProperty", member)
}

// The context image data.
// Please note that a call to BitmapData.draw() or BitmapData.copy() does not update immediately this property for performance reason. Use BitmapData.update() to do so.
// This property is updated automatically after the first game loop, according to the dirty flag property.
func (self *BitmapData) ImageData() *ImageData{
    return &ImageData{self.Object.Get("imageData")}
}

// The context image data.
// Please note that a call to BitmapData.draw() or BitmapData.copy() does not update immediately this property for performance reason. Use BitmapData.update() to do so.
// This property is updated automatically after the first game loop, according to the dirty flag property.
func (self *BitmapData) SetImageDataA(member *ImageData) {
    self.Object.Set("imageData", member)
}

// A Uint8ClampedArray view into BitmapData.buffer.
// Note that this is unavailable in some browsers (such as Epic Browser due to its security restrictions)
func (self *BitmapData) Data() *Uint8ClampedArray{
    return &Uint8ClampedArray{self.Object.Get("data")}
}

// A Uint8ClampedArray view into BitmapData.buffer.
// Note that this is unavailable in some browsers (such as Epic Browser due to its security restrictions)
func (self *BitmapData) SetDataA(member *Uint8ClampedArray) {
    self.Object.Set("data", member)
}

// An Uint32Array view into BitmapData.buffer.
func (self *BitmapData) Pixels() *Uint32Array{
    return &Uint32Array{self.Object.Get("pixels")}
}

// An Uint32Array view into BitmapData.buffer.
func (self *BitmapData) SetPixelsA(member *Uint32Array) {
    self.Object.Set("pixels", member)
}

// The PIXI.BaseTexture.
func (self *BitmapData) BaseTexture() *BaseTexture{
    return &BaseTexture{self.Object.Get("baseTexture")}
}

// The PIXI.BaseTexture.
func (self *BitmapData) SetBaseTextureA(member *BaseTexture) {
    self.Object.Set("baseTexture", member)
}

// The PIXI.Texture.
func (self *BitmapData) Texture() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// The PIXI.Texture.
func (self *BitmapData) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// The FrameData container this BitmapData uses for rendering.
func (self *BitmapData) FrameData() *FrameData{
    return &FrameData{self.Object.Get("frameData")}
}

// The FrameData container this BitmapData uses for rendering.
func (self *BitmapData) SetFrameDataA(member *FrameData) {
    self.Object.Set("frameData", member)
}

// The Frame this BitmapData uses for rendering.
func (self *BitmapData) TextureFrame() *Frame{
    return &Frame{self.Object.Get("textureFrame")}
}

// The Frame this BitmapData uses for rendering.
func (self *BitmapData) SetTextureFrameA(member *Frame) {
    self.Object.Set("textureFrame", member)
}

// The const type of this object.
func (self *BitmapData) Type() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *BitmapData) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// If disableTextureUpload is true this BitmapData will never send its image data to the GPU when its dirty flag is true.
func (self *BitmapData) DisableTextureUpload() bool{
    return self.Object.Get("disableTextureUpload").Bool()
}

// If disableTextureUpload is true this BitmapData will never send its image data to the GPU when its dirty flag is true.
func (self *BitmapData) SetDisableTextureUploadA(member bool) {
    self.Object.Set("disableTextureUpload", member)
}

// If dirty this BitmapData will be re-rendered.
func (self *BitmapData) Dirty() bool{
    return self.Object.Get("dirty").Bool()
}

// If dirty this BitmapData will be re-rendered.
func (self *BitmapData) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}



// Shifts the contents of this BitmapData by the distances given.
// 
// The image will wrap-around the edges on all sides if the wrap argument is true (the default).
func (self *BitmapData) Move(x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("move", x, y)}
}

// Shifts the contents of this BitmapData by the distances given.
// 
// The image will wrap-around the edges on all sides if the wrap argument is true (the default).
func (self *BitmapData) Move1O(x int, y int, wrap bool) *BitmapData{
    return &BitmapData{self.Object.Call("move", x, y, wrap)}
}

// Shifts the contents of this BitmapData by the distances given.
// 
// The image will wrap-around the edges on all sides if the wrap argument is true (the default).
func (self *BitmapData) MoveI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("move", args)}
}

// Shifts the contents of this BitmapData horizontally.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveH(distance int) *BitmapData{
    return &BitmapData{self.Object.Call("moveH", distance)}
}

// Shifts the contents of this BitmapData horizontally.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveH1O(distance int, wrap bool) *BitmapData{
    return &BitmapData{self.Object.Call("moveH", distance, wrap)}
}

// Shifts the contents of this BitmapData horizontally.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveHI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("moveH", args)}
}

// Shifts the contents of this BitmapData vertically.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveV(distance int) *BitmapData{
    return &BitmapData{self.Object.Call("moveV", distance)}
}

// Shifts the contents of this BitmapData vertically.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveV1O(distance int, wrap bool) *BitmapData{
    return &BitmapData{self.Object.Call("moveV", distance, wrap)}
}

// Shifts the contents of this BitmapData vertically.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveVI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("moveV", args)}
}

// Updates the given objects so that they use this BitmapData as their texture.
// This will replace any texture they will currently have set.
func (self *BitmapData) Add(object interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("add", object)}
}

// Updates the given objects so that they use this BitmapData as their texture.
// This will replace any texture they will currently have set.
func (self *BitmapData) AddI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("add", args)}
}

// Takes the given Game Object, resizes this BitmapData to match it and then draws it into this BitmapDatas canvas, ready for further processing.
// The source game object is not modified by this operation.
// If the source object uses a texture as part of a Texture Atlas or Sprite Sheet, only the current frame will be used for sizing.
// If a string is given it will assume it's a cache key and look in Phaser.Cache for an image key matching the string.
func (self *BitmapData) Load(source interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("load", source)}
}

// Takes the given Game Object, resizes this BitmapData to match it and then draws it into this BitmapDatas canvas, ready for further processing.
// The source game object is not modified by this operation.
// If the source object uses a texture as part of a Texture Atlas or Sprite Sheet, only the current frame will be used for sizing.
// If a string is given it will assume it's a cache key and look in Phaser.Cache for an image key matching the string.
func (self *BitmapData) LoadI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("load", args)}
}

// Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) Clear() *BitmapData{
    return &BitmapData{self.Object.Call("clear")}
}

// Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) Clear1O(x int) *BitmapData{
    return &BitmapData{self.Object.Call("clear", x)}
}

// Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) Clear2O(x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("clear", x, y)}
}

// Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) Clear3O(x int, y int, width int) *BitmapData{
    return &BitmapData{self.Object.Call("clear", x, y, width)}
}

// Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) Clear4O(x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("clear", x, y, width, height)}
}

// Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) ClearI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("clear", args)}
}

// Clears the BitmapData context using a clearRect.
func (self *BitmapData) Cls() {
    self.Object.Call("cls")
}

// Clears the BitmapData context using a clearRect.
func (self *BitmapData) ClsI(args ...interface{}) {
    self.Object.Call("cls", args)
}

// Fills the BitmapData with the given color.
func (self *BitmapData) Fill(r int, g int, b int) *BitmapData{
    return &BitmapData{self.Object.Call("fill", r, g, b)}
}

// Fills the BitmapData with the given color.
func (self *BitmapData) Fill1O(r int, g int, b int, a int) *BitmapData{
    return &BitmapData{self.Object.Call("fill", r, g, b, a)}
}

// Fills the BitmapData with the given color.
func (self *BitmapData) FillI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("fill", args)}
}

// Creates a new Image element by converting this BitmapDatas canvas into a dataURL.
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

// Creates a new Image element by converting this BitmapDatas canvas into a dataURL.
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

// Resizes the BitmapData. This changes the size of the underlying canvas and refreshes the buffer.
func (self *BitmapData) Resize(width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("resize", width, height)}
}

// Resizes the BitmapData. This changes the size of the underlying canvas and refreshes the buffer.
func (self *BitmapData) ResizeI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("resize", args)}
}

// This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) Update() *BitmapData{
    return &BitmapData{self.Object.Call("update")}
}

// This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) Update1O(x int) *BitmapData{
    return &BitmapData{self.Object.Call("update", x)}
}

// This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) Update2O(x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("update", x, y)}
}

// This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) Update3O(x int, y int, width int) *BitmapData{
    return &BitmapData{self.Object.Call("update", x, y, width)}
}

// This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) Update4O(x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("update", x, y, width, height)}
}

// This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) UpdateI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("update", args)}
}

// Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
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

// Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
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

// Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
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

// Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
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

// Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
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

// Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
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

// Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixel(callback interface{}, callbackContext interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", callback, callbackContext)}
}

// Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixel1O(callback interface{}, callbackContext interface{}, x int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", callback, callbackContext, x)}
}

// Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixel2O(callback interface{}, callbackContext interface{}, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", callback, callbackContext, x, y)}
}

// Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixel3O(callback interface{}, callbackContext interface{}, x int, y int, width int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", callback, callbackContext, x, y, width)}
}

// Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixel4O(callback interface{}, callbackContext interface{}, x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", callback, callbackContext, x, y, width, height)}
}

// Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixelI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("processPixel", args)}
}

// Replaces all pixels matching one color with another. The color values are given as two sets of RGBA values.
// An optional region parameter controls if the replacement happens in just a specific area of the BitmapData or the entire thing.
func (self *BitmapData) ReplaceRGB(r1 int, g1 int, b1 int, a1 int, r2 int, g2 int, b2 int, a2 int) *BitmapData{
    return &BitmapData{self.Object.Call("replaceRGB", r1, g1, b1, a1, r2, g2, b2, a2)}
}

// Replaces all pixels matching one color with another. The color values are given as two sets of RGBA values.
// An optional region parameter controls if the replacement happens in just a specific area of the BitmapData or the entire thing.
func (self *BitmapData) ReplaceRGB1O(r1 int, g1 int, b1 int, a1 int, r2 int, g2 int, b2 int, a2 int, region *Rectangle) *BitmapData{
    return &BitmapData{self.Object.Call("replaceRGB", r1, g1, b1, a1, r2, g2, b2, a2, region)}
}

// Replaces all pixels matching one color with another. The color values are given as two sets of RGBA values.
// An optional region parameter controls if the replacement happens in just a specific area of the BitmapData or the entire thing.
func (self *BitmapData) ReplaceRGBI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("replaceRGB", args)}
}

// Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSL() *BitmapData{
    return &BitmapData{self.Object.Call("setHSL")}
}

// Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSL1O(h int) *BitmapData{
    return &BitmapData{self.Object.Call("setHSL", h)}
}

// Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSL2O(h int, s int) *BitmapData{
    return &BitmapData{self.Object.Call("setHSL", h, s)}
}

// Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSL3O(h int, s int, l int) *BitmapData{
    return &BitmapData{self.Object.Call("setHSL", h, s, l)}
}

// Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSL4O(h int, s int, l int, region *Rectangle) *BitmapData{
    return &BitmapData{self.Object.Call("setHSL", h, s, l, region)}
}

// Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSLI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("setHSL", args)}
}

// Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSL() *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL")}
}

// Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSL1O(h int) *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL", h)}
}

// Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSL2O(h int, s int) *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL", h, s)}
}

// Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSL3O(h int, s int, l int) *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL", h, s, l)}
}

// Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSL4O(h int, s int, l int, region *Rectangle) *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL", h, s, l, region)}
}

// Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSLI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("shiftHSL", args)}
}

// Sets the color of the given pixel to the specified red, green, blue and alpha values.
func (self *BitmapData) SetPixel32(x int, y int, red int, green int, blue int, alpha int) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel32", x, y, red, green, blue, alpha)}
}

// Sets the color of the given pixel to the specified red, green, blue and alpha values.
func (self *BitmapData) SetPixel321O(x int, y int, red int, green int, blue int, alpha int, immediate bool) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel32", x, y, red, green, blue, alpha, immediate)}
}

// Sets the color of the given pixel to the specified red, green, blue and alpha values.
func (self *BitmapData) SetPixel32I(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel32", args)}
}

// Sets the color of the given pixel to the specified red, green and blue values.
func (self *BitmapData) SetPixel(x int, y int, red int, green int, blue int) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel", x, y, red, green, blue)}
}

// Sets the color of the given pixel to the specified red, green and blue values.
func (self *BitmapData) SetPixel1O(x int, y int, red int, green int, blue int, immediate bool) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel", x, y, red, green, blue, immediate)}
}

// Sets the color of the given pixel to the specified red, green and blue values.
func (self *BitmapData) SetPixelI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("setPixel", args)}
}

// Get the color of a specific pixel in the context into a color object.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixel(x int, y int) interface{}{
    return self.Object.Call("getPixel", x, y)
}

// Get the color of a specific pixel in the context into a color object.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixel1O(x int, y int, out interface{}) interface{}{
    return self.Object.Call("getPixel", x, y, out)
}

// Get the color of a specific pixel in the context into a color object.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelI(args ...interface{}) interface{}{
    return self.Object.Call("getPixel", args)
}

// Get the color of a specific pixel including its alpha value.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
// Note that on little-endian systems the format is 0xAABBGGRR and on big-endian the format is 0xRRGGBBAA.
func (self *BitmapData) GetPixel32(x int, y int) int{
    return self.Object.Call("getPixel32", x, y).Int()
}

// Get the color of a specific pixel including its alpha value.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
// Note that on little-endian systems the format is 0xAABBGGRR and on big-endian the format is 0xRRGGBBAA.
func (self *BitmapData) GetPixel32I(args ...interface{}) int{
    return self.Object.Call("getPixel32", args).Int()
}

// Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGB(x int, y int) interface{}{
    return self.Object.Call("getPixelRGB", x, y)
}

// Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGB1O(x int, y int, out interface{}) interface{}{
    return self.Object.Call("getPixelRGB", x, y, out)
}

// Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGB2O(x int, y int, out interface{}, hsl bool) interface{}{
    return self.Object.Call("getPixelRGB", x, y, out, hsl)
}

// Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGB3O(x int, y int, out interface{}, hsl bool, hsv bool) interface{}{
    return self.Object.Call("getPixelRGB", x, y, out, hsl, hsv)
}

// Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGBI(args ...interface{}) interface{}{
    return self.Object.Call("getPixelRGB", args)
}

// Gets all the pixels from the region specified by the given Rectangle object.
func (self *BitmapData) GetPixels(rect *Rectangle) *ImageData{
    return &ImageData{self.Object.Call("getPixels", rect)}
}

// Gets all the pixels from the region specified by the given Rectangle object.
func (self *BitmapData) GetPixelsI(args ...interface{}) *ImageData{
    return &ImageData{self.Object.Call("getPixels", args)}
}

// Scans the BitmapData, pixel by pixel, until it encounters a pixel that isn't transparent (i.e. has an alpha value > 0).
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

// Scans the BitmapData, pixel by pixel, until it encounters a pixel that isn't transparent (i.e. has an alpha value > 0).
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

// Scans the BitmapData, pixel by pixel, until it encounters a pixel that isn't transparent (i.e. has an alpha value > 0).
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

// Scans the BitmapData and calculates the bounds. This is a rectangle that defines the extent of all non-transparent pixels.
// The rectangle returned will extend from the top-left of the image to the bottom-right, excluding transparent pixels.
func (self *BitmapData) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// Scans the BitmapData and calculates the bounds. This is a rectangle that defines the extent of all non-transparent pixels.
// The rectangle returned will extend from the top-left of the image to the bottom-right, excluding transparent pixels.
func (self *BitmapData) GetBounds1O(rect *Rectangle) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", rect)}
}

// Scans the BitmapData and calculates the bounds. This is a rectangle that defines the extent of all non-transparent pixels.
// The rectangle returned will extend from the top-left of the image to the bottom-right, excluding transparent pixels.
func (self *BitmapData) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld() *Image{
    return &Image{self.Object.Call("addToWorld")}
}

// Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld1O(x int) *Image{
    return &Image{self.Object.Call("addToWorld", x)}
}

// Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld2O(x int, y int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y)}
}

// Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld3O(x int, y int, anchorX int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, anchorX)}
}

// Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld4O(x int, y int, anchorX int, anchorY int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, anchorX, anchorY)}
}

// Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld5O(x int, y int, anchorX int, anchorY int, scaleX int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, anchorX, anchorY, scaleX)}
}

// Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorld6O(x int, y int, anchorX int, anchorY int, scaleX int, scaleY int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, anchorX, anchorY, scaleX, scaleY)}
}

// Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorldI(args ...interface{}) *Image{
    return &Image{self.Object.Call("addToWorld", args)}
}

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Copies a rectangular area from the source object to this BitmapData. If you give `null` as the source it will copy from itself.
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

// Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransform() *BitmapData{
    return &BitmapData{self.Object.Call("copyTransform")}
}

// Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransform1O(source interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("copyTransform", source)}
}

// Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransform2O(source interface{}, blendMode string) *BitmapData{
    return &BitmapData{self.Object.Call("copyTransform", source, blendMode)}
}

// Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransform3O(source interface{}, blendMode string, roundPx bool) *BitmapData{
    return &BitmapData{self.Object.Call("copyTransform", source, blendMode, roundPx)}
}

// Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransformI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("copyTransform", args)}
}

// Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRect(source interface{}, area *Rectangle, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("copyRect", source, area, x, y)}
}

// Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRect1O(source interface{}, area *Rectangle, x int, y int, alpha int) *BitmapData{
    return &BitmapData{self.Object.Call("copyRect", source, area, x, y, alpha)}
}

// Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRect2O(source interface{}, area *Rectangle, x int, y int, alpha int, blendMode string) *BitmapData{
    return &BitmapData{self.Object.Call("copyRect", source, area, x, y, alpha, blendMode)}
}

// Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRect3O(source interface{}, area *Rectangle, x int, y int, alpha int, blendMode string, roundPx bool) *BitmapData{
    return &BitmapData{self.Object.Call("copyRect", source, area, x, y, alpha, blendMode, roundPx)}
}

// Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRectI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("copyRect", args)}
}

// Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
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

// Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
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

// Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
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

// Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
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

// Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
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

// Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
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

// Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
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

// Draws the given Phaser.Sprite, Phaser.Image or Phaser.Text to this BitmapData at the coordinates specified.
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

// Draws the immediate children of a Phaser.Group to this BitmapData.
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

// Draws the immediate children of a Phaser.Group to this BitmapData.
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

// Draws the immediate children of a Phaser.Group to this BitmapData.
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

// Draws the immediate children of a Phaser.Group to this BitmapData.
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

// A proxy for drawGroup that handles child iteration for more complex Game Objects.
func (self *BitmapData) DrawGroupProxy(child interface{}) {
    self.Object.Call("drawGroupProxy", child)
}

// A proxy for drawGroup that handles child iteration for more complex Game Objects.
func (self *BitmapData) DrawGroupProxy1O(child interface{}, blendMode string) {
    self.Object.Call("drawGroupProxy", child, blendMode)
}

// A proxy for drawGroup that handles child iteration for more complex Game Objects.
func (self *BitmapData) DrawGroupProxy2O(child interface{}, blendMode string, roundPx bool) {
    self.Object.Call("drawGroupProxy", child, blendMode, roundPx)
}

// A proxy for drawGroup that handles child iteration for more complex Game Objects.
func (self *BitmapData) DrawGroupProxyI(args ...interface{}) {
    self.Object.Call("drawGroupProxy", args)
}

// Draws the Game Object or Group to this BitmapData and then recursively iterates through all of its children.
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

// Draws the Game Object or Group to this BitmapData and then recursively iterates through all of its children.
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

// Draws the Game Object or Group to this BitmapData and then recursively iterates through all of its children.
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

// Draws the Game Object or Group to this BitmapData and then recursively iterates through all of its children.
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

// Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) Shadow(color string) *BitmapData{
    return &BitmapData{self.Object.Call("shadow", color)}
}

// Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) Shadow1O(color string, blur int) *BitmapData{
    return &BitmapData{self.Object.Call("shadow", color, blur)}
}

// Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) Shadow2O(color string, blur int, x int) *BitmapData{
    return &BitmapData{self.Object.Call("shadow", color, blur, x)}
}

// Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) Shadow3O(color string, blur int, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("shadow", color, blur, x, y)}
}

// Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) ShadowI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("shadow", args)}
}

// Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMask(source interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("alphaMask", source)}
}

// Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMask1O(source interface{}, mask interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("alphaMask", source, mask)}
}

// Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMask2O(source interface{}, mask interface{}, sourceRect *Rectangle) *BitmapData{
    return &BitmapData{self.Object.Call("alphaMask", source, mask, sourceRect)}
}

// Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMask3O(source interface{}, mask interface{}, sourceRect *Rectangle, maskRect *Rectangle) *BitmapData{
    return &BitmapData{self.Object.Call("alphaMask", source, mask, sourceRect, maskRect)}
}

// Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMaskI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("alphaMask", args)}
}

// Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
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

// Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
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

// Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
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

// Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
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

// Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
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

// Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
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

// Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
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

// Draws a filled Rectangle to the BitmapData at the given x, y coordinates and width / height in size.
func (self *BitmapData) Rect(x int, y int, width int, height int) *BitmapData{
    return &BitmapData{self.Object.Call("rect", x, y, width, height)}
}

// Draws a filled Rectangle to the BitmapData at the given x, y coordinates and width / height in size.
func (self *BitmapData) Rect1O(x int, y int, width int, height int, fillStyle string) *BitmapData{
    return &BitmapData{self.Object.Call("rect", x, y, width, height, fillStyle)}
}

// Draws a filled Rectangle to the BitmapData at the given x, y coordinates and width / height in size.
func (self *BitmapData) RectI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("rect", args)}
}

// Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) Text(text string, x int, y int) *BitmapData{
    return &BitmapData{self.Object.Call("text", text, x, y)}
}

// Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) Text1O(text string, x int, y int, font string) *BitmapData{
    return &BitmapData{self.Object.Call("text", text, x, y, font)}
}

// Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) Text2O(text string, x int, y int, font string, color string) *BitmapData{
    return &BitmapData{self.Object.Call("text", text, x, y, font, color)}
}

// Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) Text3O(text string, x int, y int, font string, color string, shadow bool) *BitmapData{
    return &BitmapData{self.Object.Call("text", text, x, y, font, color, shadow)}
}

// Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) TextI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("text", args)}
}

// Draws a filled Circle to the BitmapData at the given x, y coordinates and radius in size.
func (self *BitmapData) Circle(x int, y int, radius int) *BitmapData{
    return &BitmapData{self.Object.Call("circle", x, y, radius)}
}

// Draws a filled Circle to the BitmapData at the given x, y coordinates and radius in size.
func (self *BitmapData) Circle1O(x int, y int, radius int, fillStyle string) *BitmapData{
    return &BitmapData{self.Object.Call("circle", x, y, radius, fillStyle)}
}

// Draws a filled Circle to the BitmapData at the given x, y coordinates and radius in size.
func (self *BitmapData) CircleI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("circle", args)}
}

// Draws a line between the coordinates given in the color and thickness specified.
func (self *BitmapData) Line(x1 int, y1 int, x2 int, y2 int) *BitmapData{
    return &BitmapData{self.Object.Call("line", x1, y1, x2, y2)}
}

// Draws a line between the coordinates given in the color and thickness specified.
func (self *BitmapData) Line1O(x1 int, y1 int, x2 int, y2 int, color string) *BitmapData{
    return &BitmapData{self.Object.Call("line", x1, y1, x2, y2, color)}
}

// Draws a line between the coordinates given in the color and thickness specified.
func (self *BitmapData) Line2O(x1 int, y1 int, x2 int, y2 int, color string, width int) *BitmapData{
    return &BitmapData{self.Object.Call("line", x1, y1, x2, y2, color, width)}
}

// Draws a line between the coordinates given in the color and thickness specified.
func (self *BitmapData) LineI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("line", args)}
}

// Takes the given Line object and image and renders it to this BitmapData as a repeating texture line.
func (self *BitmapData) TextureLine(line *Line, image interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("textureLine", line, image)}
}

// Takes the given Line object and image and renders it to this BitmapData as a repeating texture line.
func (self *BitmapData) TextureLine1O(line *Line, image interface{}, repeat string) *BitmapData{
    return &BitmapData{self.Object.Call("textureLine", line, image, repeat)}
}

// Takes the given Line object and image and renders it to this BitmapData as a repeating texture line.
func (self *BitmapData) TextureLineI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("textureLine", args)}
}

// If the game is running in WebGL this will push the texture up to the GPU if it's dirty.
// This is called automatically if the BitmapData is being used by a Sprite, otherwise you need to remember to call it in your render function.
// If you wish to suppress this functionality set BitmapData.disableTextureUpload to `true`.
func (self *BitmapData) Render() *BitmapData{
    return &BitmapData{self.Object.Call("render")}
}

// If the game is running in WebGL this will push the texture up to the GPU if it's dirty.
// This is called automatically if the BitmapData is being used by a Sprite, otherwise you need to remember to call it in your render function.
// If you wish to suppress this functionality set BitmapData.disableTextureUpload to `true`.
func (self *BitmapData) RenderI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("render", args)}
}

// Destroys this BitmapData and puts the canvas it was using back into the canvas pool for re-use.
func (self *BitmapData) Destroy() {
    self.Object.Call("destroy")
}

// Destroys this BitmapData and puts the canvas it was using back into the canvas pool for re-use.
func (self *BitmapData) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Resets the blend mode (effectively sets it to 'source-over')
func (self *BitmapData) BlendReset() *BitmapData{
    return &BitmapData{self.Object.Call("blendReset")}
}

// Resets the blend mode (effectively sets it to 'source-over')
func (self *BitmapData) BlendResetI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendReset", args)}
}

// Sets the blend mode to 'source-over'
func (self *BitmapData) BlendSourceOver() *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceOver")}
}

// Sets the blend mode to 'source-over'
func (self *BitmapData) BlendSourceOverI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceOver", args)}
}

// Sets the blend mode to 'source-in'
func (self *BitmapData) BlendSourceIn() *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceIn")}
}

// Sets the blend mode to 'source-in'
func (self *BitmapData) BlendSourceInI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceIn", args)}
}

// Sets the blend mode to 'source-out'
func (self *BitmapData) BlendSourceOut() *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceOut")}
}

// Sets the blend mode to 'source-out'
func (self *BitmapData) BlendSourceOutI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceOut", args)}
}

// Sets the blend mode to 'source-atop'
func (self *BitmapData) BlendSourceAtop() *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceAtop")}
}

// Sets the blend mode to 'source-atop'
func (self *BitmapData) BlendSourceAtopI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSourceAtop", args)}
}

// Sets the blend mode to 'destination-over'
func (self *BitmapData) BlendDestinationOver() *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationOver")}
}

// Sets the blend mode to 'destination-over'
func (self *BitmapData) BlendDestinationOverI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationOver", args)}
}

// Sets the blend mode to 'destination-in'
func (self *BitmapData) BlendDestinationIn() *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationIn")}
}

// Sets the blend mode to 'destination-in'
func (self *BitmapData) BlendDestinationInI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationIn", args)}
}

// Sets the blend mode to 'destination-out'
func (self *BitmapData) BlendDestinationOut() *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationOut")}
}

// Sets the blend mode to 'destination-out'
func (self *BitmapData) BlendDestinationOutI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationOut", args)}
}

// Sets the blend mode to 'destination-atop'
func (self *BitmapData) BlendDestinationAtop() *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationAtop")}
}

// Sets the blend mode to 'destination-atop'
func (self *BitmapData) BlendDestinationAtopI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDestinationAtop", args)}
}

// Sets the blend mode to 'xor'
func (self *BitmapData) BlendXor() *BitmapData{
    return &BitmapData{self.Object.Call("blendXor")}
}

// Sets the blend mode to 'xor'
func (self *BitmapData) BlendXorI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendXor", args)}
}

// Sets the blend mode to 'lighter'
func (self *BitmapData) BlendAdd() *BitmapData{
    return &BitmapData{self.Object.Call("blendAdd")}
}

// Sets the blend mode to 'lighter'
func (self *BitmapData) BlendAddI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendAdd", args)}
}

// Sets the blend mode to 'multiply'
func (self *BitmapData) BlendMultiply() *BitmapData{
    return &BitmapData{self.Object.Call("blendMultiply")}
}

// Sets the blend mode to 'multiply'
func (self *BitmapData) BlendMultiplyI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendMultiply", args)}
}

// Sets the blend mode to 'screen'
func (self *BitmapData) BlendScreen() *BitmapData{
    return &BitmapData{self.Object.Call("blendScreen")}
}

// Sets the blend mode to 'screen'
func (self *BitmapData) BlendScreenI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendScreen", args)}
}

// Sets the blend mode to 'overlay'
func (self *BitmapData) BlendOverlay() *BitmapData{
    return &BitmapData{self.Object.Call("blendOverlay")}
}

// Sets the blend mode to 'overlay'
func (self *BitmapData) BlendOverlayI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendOverlay", args)}
}

// Sets the blend mode to 'darken'
func (self *BitmapData) BlendDarken() *BitmapData{
    return &BitmapData{self.Object.Call("blendDarken")}
}

// Sets the blend mode to 'darken'
func (self *BitmapData) BlendDarkenI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDarken", args)}
}

// Sets the blend mode to 'lighten'
func (self *BitmapData) BlendLighten() *BitmapData{
    return &BitmapData{self.Object.Call("blendLighten")}
}

// Sets the blend mode to 'lighten'
func (self *BitmapData) BlendLightenI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendLighten", args)}
}

// Sets the blend mode to 'color-dodge'
func (self *BitmapData) BlendColorDodge() *BitmapData{
    return &BitmapData{self.Object.Call("blendColorDodge")}
}

// Sets the blend mode to 'color-dodge'
func (self *BitmapData) BlendColorDodgeI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendColorDodge", args)}
}

// Sets the blend mode to 'color-burn'
func (self *BitmapData) BlendColorBurn() *BitmapData{
    return &BitmapData{self.Object.Call("blendColorBurn")}
}

// Sets the blend mode to 'color-burn'
func (self *BitmapData) BlendColorBurnI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendColorBurn", args)}
}

// Sets the blend mode to 'hard-light'
func (self *BitmapData) BlendHardLight() *BitmapData{
    return &BitmapData{self.Object.Call("blendHardLight")}
}

// Sets the blend mode to 'hard-light'
func (self *BitmapData) BlendHardLightI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendHardLight", args)}
}

// Sets the blend mode to 'soft-light'
func (self *BitmapData) BlendSoftLight() *BitmapData{
    return &BitmapData{self.Object.Call("blendSoftLight")}
}

// Sets the blend mode to 'soft-light'
func (self *BitmapData) BlendSoftLightI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSoftLight", args)}
}

// Sets the blend mode to 'difference'
func (self *BitmapData) BlendDifference() *BitmapData{
    return &BitmapData{self.Object.Call("blendDifference")}
}

// Sets the blend mode to 'difference'
func (self *BitmapData) BlendDifferenceI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendDifference", args)}
}

// Sets the blend mode to 'exclusion'
func (self *BitmapData) BlendExclusion() *BitmapData{
    return &BitmapData{self.Object.Call("blendExclusion")}
}

// Sets the blend mode to 'exclusion'
func (self *BitmapData) BlendExclusionI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendExclusion", args)}
}

// Sets the blend mode to 'hue'
func (self *BitmapData) BlendHue() *BitmapData{
    return &BitmapData{self.Object.Call("blendHue")}
}

// Sets the blend mode to 'hue'
func (self *BitmapData) BlendHueI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendHue", args)}
}

// Sets the blend mode to 'saturation'
func (self *BitmapData) BlendSaturation() *BitmapData{
    return &BitmapData{self.Object.Call("blendSaturation")}
}

// Sets the blend mode to 'saturation'
func (self *BitmapData) BlendSaturationI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendSaturation", args)}
}

// Sets the blend mode to 'color'
func (self *BitmapData) BlendColor() *BitmapData{
    return &BitmapData{self.Object.Call("blendColor")}
}

// Sets the blend mode to 'color'
func (self *BitmapData) BlendColorI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendColor", args)}
}

// Sets the blend mode to 'luminosity'
func (self *BitmapData) BlendLuminosity() *BitmapData{
    return &BitmapData{self.Object.Call("blendLuminosity")}
}

// Sets the blend mode to 'luminosity'
func (self *BitmapData) BlendLuminosityI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("blendLuminosity", args)}
}

// Gets a JavaScript object that has 6 properties set that are used by BitmapData in a transform.
func (self *BitmapData) GetTransform(translateX int, translateY int, scaleX int, scaleY int, skewX int, skewY int) interface{}{
    return self.Object.Call("getTransform", translateX, translateY, scaleX, scaleY, skewX, skewY)
}

// Gets a JavaScript object that has 6 properties set that are used by BitmapData in a transform.
func (self *BitmapData) GetTransformI(args ...interface{}) interface{}{
    return self.Object.Call("getTransform", args)
}
