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


// A reference to the currently running game.
func (self *BitmapData) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *BitmapData) SetGame(member Game) {
    self.Set("game", member)
}

// The key of the BitmapData in the Cache, if stored there.
func (self *BitmapData) GetKey() string{
    return self.Get("key").String()
}

// The key of the BitmapData in the Cache, if stored there.
func (self *BitmapData) SetKey(member string) {
    self.Set("key", member)
}

// The width of the BitmapData in pixels.
func (self *BitmapData) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the BitmapData in pixels.
func (self *BitmapData) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the BitmapData in pixels.
func (self *BitmapData) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the BitmapData in pixels.
func (self *BitmapData) SetHeight(member float64) {
    self.Set("height", member)
}

// The canvas to which this BitmapData draws.
func (self *BitmapData) GetCanvas() dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Get("canvas"))
}

// The canvas to which this BitmapData draws.
func (self *BitmapData) SetCanvas(member dom.HTMLCanvasElement) {
    self.Set("canvas", member)
}

// The 2d context of the canvas.
func (self *BitmapData) GetContext() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Get("context"))
}

// The 2d context of the canvas.
func (self *BitmapData) SetContext(member dom.CanvasRenderingContext2D) {
    self.Set("context", member)
}

// A reference to BitmapData.context.
func (self *BitmapData) GetCtx() dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Get("ctx"))
}

// A reference to BitmapData.context.
func (self *BitmapData) SetCtx(member dom.CanvasRenderingContext2D) {
    self.Set("ctx", member)
}

// The context property needed for smoothing this Canvas.
func (self *BitmapData) GetSmoothProperty() string{
    return self.Get("smoothProperty").String()
}

// The context property needed for smoothing this Canvas.
func (self *BitmapData) SetSmoothProperty(member string) {
    self.Set("smoothProperty", member)
}

// The context image data.
// Please note that a call to BitmapData.draw() or BitmapData.copy() does not update immediately this property for performance reason. Use BitmapData.update() to do so.
// This property is updated automatically after the first game loop, according to the dirty flag property.
func (self *BitmapData) GetImageData() ImageData{
    return ImageData{self.Get("imageData")}
}

// The context image data.
// Please note that a call to BitmapData.draw() or BitmapData.copy() does not update immediately this property for performance reason. Use BitmapData.update() to do so.
// This property is updated automatically after the first game loop, according to the dirty flag property.
func (self *BitmapData) SetImageData(member ImageData) {
    self.Set("imageData", member)
}

// A Uint8ClampedArray view into BitmapData.buffer.
// Note that this is unavailable in some browsers (such as Epic Browser due to its security restrictions)
func (self *BitmapData) GetData() Uint8ClampedArray{
    return Uint8ClampedArray{self.Get("data")}
}

// A Uint8ClampedArray view into BitmapData.buffer.
// Note that this is unavailable in some browsers (such as Epic Browser due to its security restrictions)
func (self *BitmapData) SetData(member Uint8ClampedArray) {
    self.Set("data", member)
}

// An Uint32Array view into BitmapData.buffer.
func (self *BitmapData) GetPixels() Uint32Array{
    return Uint32Array{self.Get("pixels")}
}

// An Uint32Array view into BitmapData.buffer.
func (self *BitmapData) SetPixels(member Uint32Array) {
    self.Set("pixels", member)
}

// The PIXI.BaseTexture.
func (self *BitmapData) GetBaseTexture() BaseTexture{
    return BaseTexture{self.Get("baseTexture")}
}

// The PIXI.BaseTexture.
func (self *BitmapData) SetBaseTexture(member BaseTexture) {
    self.Set("baseTexture", member)
}

// The PIXI.Texture.
func (self *BitmapData) GetTexture() Texture{
    return Texture{self.Get("texture")}
}

// The PIXI.Texture.
func (self *BitmapData) SetTexture(member Texture) {
    self.Set("texture", member)
}

// The FrameData container this BitmapData uses for rendering.
func (self *BitmapData) GetFrameData() FrameData{
    return FrameData{self.Get("frameData")}
}

// The FrameData container this BitmapData uses for rendering.
func (self *BitmapData) SetFrameData(member FrameData) {
    self.Set("frameData", member)
}

// The Frame this BitmapData uses for rendering.
func (self *BitmapData) GetTextureFrame() Frame{
    return Frame{self.Get("textureFrame")}
}

// The Frame this BitmapData uses for rendering.
func (self *BitmapData) SetTextureFrame(member Frame) {
    self.Set("textureFrame", member)
}

// The const type of this object.
func (self *BitmapData) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *BitmapData) SetType(member float64) {
    self.Set("type", member)
}

// If disableTextureUpload is true this BitmapData will never send its image data to the GPU when its dirty flag is true.
func (self *BitmapData) GetDisableTextureUpload() bool{
    return self.Get("disableTextureUpload").Bool()
}

// If disableTextureUpload is true this BitmapData will never send its image data to the GPU when its dirty flag is true.
func (self *BitmapData) SetDisableTextureUpload(member bool) {
    self.Set("disableTextureUpload", member)
}

// If dirty this BitmapData will be re-rendered.
func (self *BitmapData) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// If dirty this BitmapData will be re-rendered.
func (self *BitmapData) SetDirty(member bool) {
    self.Set("dirty", member)
}



// Shifts the contents of this BitmapData by the distances given.
// 
// The image will wrap-around the edges on all sides if the wrap argument is true (the default).
func (self *BitmapData) MoveI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("move", args)}
}

// Shifts the contents of this BitmapData horizontally.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveHI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("moveH", args)}
}

// Shifts the contents of this BitmapData vertically.
// 
// The image will wrap-around the sides if the wrap argument is true (the default).
func (self *BitmapData) MoveVI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("moveV", args)}
}

// Updates the given objects so that they use this BitmapData as their texture.
// This will replace any texture they will currently have set.
func (self *BitmapData) AddI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("add", args)}
}

// Takes the given Game Object, resizes this BitmapData to match it and then draws it into this BitmapDatas canvas, ready for further processing.
// The source game object is not modified by this operation.
// If the source object uses a texture as part of a Texture Atlas or Sprite Sheet, only the current frame will be used for sizing.
// If a string is given it will assume it's a cache key and look in Phaser.Cache for an image key matching the string.
func (self *BitmapData) LoadI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("load", args)}
}

// Clears the BitmapData context using a clearRect.
// 
// You can optionally define the area to clear.
// If the arguments are left empty it will clear the entire canvas.
// 
// You may need to call BitmapData.update after this in order to clear out the pixel data, 
// but Phaser will not do this automatically for you.
func (self *BitmapData) ClearI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("clear", args)}
}

// Clears the BitmapData context using a clearRect.
func (self *BitmapData) ClsI(args ...interface{}) {
    self.Call("cls", args)
}

// Fills the BitmapData with the given color.
func (self *BitmapData) FillI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("fill", args)}
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
func (self *BitmapData) GenerateTextureI(args ...interface{}) Texture{
    return Texture{self.Call("generateTexture", args)}
}

// Resizes the BitmapData. This changes the size of the underlying canvas and refreshes the buffer.
func (self *BitmapData) ResizeI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("resize", args)}
}

// This re-creates the BitmapData.imageData from the current context.
// It then re-builds the ArrayBuffer, the data Uint8ClampedArray reference and the pixels Int32Array.
// If not given the dimensions defaults to the full size of the context.
// 
// Warning: This is a very expensive operation, so use it sparingly.
func (self *BitmapData) UpdateI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("update", args)}
}

// Scans through the area specified in this BitmapData and sends a color object for every pixel to the given callback.
// The callback will be sent a color object with 6 properties: `{ r: number, g: number, b: number, a: number, color: number, rgba: string }`.
// Where r, g, b and a are integers between 0 and 255 representing the color component values for red, green, blue and alpha.
// The `color` property is an Int32 of the full color. Note the endianess of this will change per system.
// The `rgba` property is a CSS style rgba() string which can be used with context.fillStyle calls, among others.
// The callback will also be sent the pixels x and y coordinates respectively.
// The callback must return either `false`, in which case no change will be made to the pixel, or a new color object.
// If a new color object is returned the pixel will be set to the r, g, b and a color values given within it.
func (self *BitmapData) ProcessPixelRGBI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("processPixelRGB", args)}
}

// Scans through the area specified in this BitmapData and sends the color for every pixel to the given callback along with its x and y coordinates.
// Whatever value the callback returns is set as the new color for that pixel, unless it returns the same color, in which case it's skipped.
// Note that the format of the color received will be different depending on if the system is big or little endian.
// It is expected that your callback will deal with endianess. If you'd rather Phaser did it then use processPixelRGB instead.
// The callback will also be sent the pixels x and y coordinates respectively.
func (self *BitmapData) ProcessPixelI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("processPixel", args)}
}

// Replaces all pixels matching one color with another. The color values are given as two sets of RGBA values.
// An optional region parameter controls if the replacement happens in just a specific area of the BitmapData or the entire thing.
func (self *BitmapData) ReplaceRGBI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("replaceRGB", args)}
}

// Sets the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
func (self *BitmapData) SetHSLI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("setHSL", args)}
}

// Shifts any or all of the hue, saturation and lightness values on every pixel in the given region, or the whole BitmapData if no region was specified.
// Shifting will add the given value onto the current h, s and l values, not replace them.
// The hue is wrapped to keep it within the range 0 to 1. Saturation and lightness are clamped to not exceed 1.
func (self *BitmapData) ShiftHSLI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("shiftHSL", args)}
}

// Sets the color of the given pixel to the specified red, green, blue and alpha values.
func (self *BitmapData) SetPixel32I(args ...interface{}) BitmapData{
    return BitmapData{self.Call("setPixel32", args)}
}

// Sets the color of the given pixel to the specified red, green and blue values.
func (self *BitmapData) SetPixelI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("setPixel", args)}
}

// Get the color of a specific pixel in the context into a color object.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelI(args ...interface{}) interface{}{
    return self.Call("getPixel", args)
}

// Get the color of a specific pixel including its alpha value.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
// Note that on little-endian systems the format is 0xAABBGGRR and on big-endian the format is 0xRRGGBBAA.
func (self *BitmapData) GetPixel32I(args ...interface{}) float64{
    return self.Call("getPixel32", args).Float()
}

// Get the color of a specific pixel including its alpha value as a color object containing r,g,b,a and rgba properties.
// If you have drawn anything to the BitmapData since it was created you must call BitmapData.update to refresh the array buffer,
// otherwise this may return out of date color values, or worse - throw a run-time error as it tries to access an array element that doesn't exist.
func (self *BitmapData) GetPixelRGBI(args ...interface{}) interface{}{
    return self.Call("getPixelRGB", args)
}

// Gets all the pixels from the region specified by the given Rectangle object.
func (self *BitmapData) GetPixelsI(args ...interface{}) ImageData{
    return ImageData{self.Call("getPixels", args)}
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
    return self.Call("getFirstPixel", args)
}

// Scans the BitmapData and calculates the bounds. This is a rectangle that defines the extent of all non-transparent pixels.
// The rectangle returned will extend from the top-left of the image to the bottom-right, excluding transparent pixels.
func (self *BitmapData) GetBoundsI(args ...interface{}) Rectangle{
    return Rectangle{self.Call("getBounds", args)}
}

// Creates a new Phaser.Image object, assigns this BitmapData to be its texture, adds it to the world then returns it.
func (self *BitmapData) AddToWorldI(args ...interface{}) Image{
    return Image{self.Call("addToWorld", args)}
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
func (self *BitmapData) CopyI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("copy", args)}
}

// Draws the given `source` Game Object to this BitmapData, using its `worldTransform` property to set the
// position, scale and rotation of where it is drawn. This function is used internally by `drawGroup`.
// It takes the objects tint and scale mode into consideration before drawing.
// 
// You can optionally specify Blend Mode and Round Pixels arguments.
func (self *BitmapData) CopyTransformI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("copyTransform", args)}
}

// Copies the area defined by the Rectangle parameter from the source image to this BitmapData at the given location.
func (self *BitmapData) CopyRectI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("copyRect", args)}
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
func (self *BitmapData) DrawI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("draw", args)}
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
func (self *BitmapData) DrawGroupI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("drawGroup", args)}
}

// A proxy for drawGroup that handles child iteration for more complex Game Objects.
func (self *BitmapData) DrawGroupProxyI(args ...interface{}) {
    self.Call("drawGroupProxy", args)
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
func (self *BitmapData) DrawFullI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("drawFull", args)}
}

// Sets the shadow properties of this BitmapDatas context which will affect all draw operations made to it.
// You can cancel an existing shadow by calling this method and passing no parameters.
// Note: At the time of writing (October 2014) Chrome still doesn't support shadowBlur used with drawImage.
func (self *BitmapData) ShadowI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("shadow", args)}
}

// Draws the image onto this BitmapData using an image as an alpha mask.
func (self *BitmapData) AlphaMaskI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("alphaMask", args)}
}

// Scans this BitmapData for all pixels matching the given r,g,b values and then draws them into the given destination BitmapData.
// The original BitmapData remains unchanged.
// The destination BitmapData must be large enough to receive all of the pixels that are scanned unless the 'resize' parameter is true.
// Although the destination BitmapData is returned from this method, it's actually modified directly in place, meaning this call is perfectly valid:
// `picture.extract(mask, r, g, b)`
// You can specify optional r2, g2, b2 color values. If given the pixel written to the destination bitmap will be of the r2, g2, b2 color.
// If not given it will be written as the same color it was extracted. You can provide one or more alternative colors, allowing you to tint
// the color during extraction.
func (self *BitmapData) ExtractI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("extract", args)}
}

// Draws a filled Rectangle to the BitmapData at the given x, y coordinates and width / height in size.
func (self *BitmapData) RectI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("rect", args)}
}

// Draws text to the BitmapData in the given font and color.
// The default font is 14px Courier, so useful for quickly drawing debug text.
// If you need to do a lot of font work to this BitmapData we'd recommend implementing your own text draw method.
func (self *BitmapData) TextI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("text", args)}
}

// Draws a filled Circle to the BitmapData at the given x, y coordinates and radius in size.
func (self *BitmapData) CircleI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("circle", args)}
}

// Draws a line between the coordinates given in the color and thickness specified.
func (self *BitmapData) LineI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("line", args)}
}

// Takes the given Line object and image and renders it to this BitmapData as a repeating texture line.
func (self *BitmapData) TextureLineI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("textureLine", args)}
}

// If the game is running in WebGL this will push the texture up to the GPU if it's dirty.
// This is called automatically if the BitmapData is being used by a Sprite, otherwise you need to remember to call it in your render function.
// If you wish to suppress this functionality set BitmapData.disableTextureUpload to `true`.
func (self *BitmapData) RenderI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("render", args)}
}

// Destroys this BitmapData and puts the canvas it was using back into the canvas pool for re-use.
func (self *BitmapData) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Resets the blend mode (effectively sets it to 'source-over')
func (self *BitmapData) BlendResetI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendReset", args)}
}

// Sets the blend mode to 'source-over'
func (self *BitmapData) BlendSourceOverI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendSourceOver", args)}
}

// Sets the blend mode to 'source-in'
func (self *BitmapData) BlendSourceInI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendSourceIn", args)}
}

// Sets the blend mode to 'source-out'
func (self *BitmapData) BlendSourceOutI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendSourceOut", args)}
}

// Sets the blend mode to 'source-atop'
func (self *BitmapData) BlendSourceAtopI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendSourceAtop", args)}
}

// Sets the blend mode to 'destination-over'
func (self *BitmapData) BlendDestinationOverI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendDestinationOver", args)}
}

// Sets the blend mode to 'destination-in'
func (self *BitmapData) BlendDestinationInI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendDestinationIn", args)}
}

// Sets the blend mode to 'destination-out'
func (self *BitmapData) BlendDestinationOutI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendDestinationOut", args)}
}

// Sets the blend mode to 'destination-atop'
func (self *BitmapData) BlendDestinationAtopI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendDestinationAtop", args)}
}

// Sets the blend mode to 'xor'
func (self *BitmapData) BlendXorI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendXor", args)}
}

// Sets the blend mode to 'lighter'
func (self *BitmapData) BlendAddI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendAdd", args)}
}

// Sets the blend mode to 'multiply'
func (self *BitmapData) BlendMultiplyI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendMultiply", args)}
}

// Sets the blend mode to 'screen'
func (self *BitmapData) BlendScreenI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendScreen", args)}
}

// Sets the blend mode to 'overlay'
func (self *BitmapData) BlendOverlayI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendOverlay", args)}
}

// Sets the blend mode to 'darken'
func (self *BitmapData) BlendDarkenI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendDarken", args)}
}

// Sets the blend mode to 'lighten'
func (self *BitmapData) BlendLightenI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendLighten", args)}
}

// Sets the blend mode to 'color-dodge'
func (self *BitmapData) BlendColorDodgeI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendColorDodge", args)}
}

// Sets the blend mode to 'color-burn'
func (self *BitmapData) BlendColorBurnI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendColorBurn", args)}
}

// Sets the blend mode to 'hard-light'
func (self *BitmapData) BlendHardLightI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendHardLight", args)}
}

// Sets the blend mode to 'soft-light'
func (self *BitmapData) BlendSoftLightI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendSoftLight", args)}
}

// Sets the blend mode to 'difference'
func (self *BitmapData) BlendDifferenceI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendDifference", args)}
}

// Sets the blend mode to 'exclusion'
func (self *BitmapData) BlendExclusionI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendExclusion", args)}
}

// Sets the blend mode to 'hue'
func (self *BitmapData) BlendHueI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendHue", args)}
}

// Sets the blend mode to 'saturation'
func (self *BitmapData) BlendSaturationI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendSaturation", args)}
}

// Sets the blend mode to 'color'
func (self *BitmapData) BlendColorI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendColor", args)}
}

// Sets the blend mode to 'luminosity'
func (self *BitmapData) BlendLuminosityI(args ...interface{}) BitmapData{
    return BitmapData{self.Call("blendLuminosity", args)}
}

// Gets a JavaScript object that has 6 properties set that are used by BitmapData in a transform.
func (self *BitmapData) GetTransformI(args ...interface{}) interface{}{
    return self.Call("getTransform", args)
}
