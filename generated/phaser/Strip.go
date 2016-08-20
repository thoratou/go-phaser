// Automatic generation for PIXI.Strip
// generated file Strip.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type Strip struct {
    *js.Object
}


// The texture of the strip
func (self *Strip) GetTexture() *Texture{
    return &Texture{self.Get("texture")}
}

// The texture of the strip
func (self *Strip) SetTexture(member *Texture) {
    self.Set("texture", member)
}

// Whether the strip is dirty or not
func (self *Strip) GetDirty() bool{
    return self.Get("dirty").Bool()
}

// Whether the strip is dirty or not
func (self *Strip) SetDirty(member bool) {
    self.Set("dirty", member)
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
func (self *Strip) GetBlendMode() float64{
    return self.Get("blendMode").Float()
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
func (self *Strip) SetBlendMode(member float64) {
    self.Set("blendMode", member)
}

// Triangles in canvas mode are automatically antialiased, use this value to force triangles to overlap a bit with each other.
func (self *Strip) GetCanvasPadding() float64{
    return self.Get("canvasPadding").Float()
}

// Triangles in canvas mode are automatically antialiased, use this value to force triangles to overlap a bit with each other.
func (self *Strip) SetCanvasPadding(member float64) {
    self.Set("canvasPadding", member)
}

// [read-only] The array of children of this container.
func (self *Strip) GetChildren() []DisplayObject{
	array := self.Get("children")
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// [read-only] The array of children of this container.
func (self *Strip) SetChildren(member []DisplayObject) {
    self.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Strip) GetIgnoreChildInput() bool{
    return self.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Strip) SetIgnoreChildInput(member bool) {
    self.Set("ignoreChildInput", member)
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Strip) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Strip) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Strip) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Strip) SetHeight(member float64) {
    self.Set("height", member)
}



// Renders a flat strip
func (self *Strip) RenderStripFlatI(args ...interface{}) {
    self.Call("renderStripFlat", args)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *Strip) OnTextureUpdateI(args ...interface{}) {
    self.Call("onTextureUpdate", args)
}

// Returns the bounds of the mesh as a rectangle. The bounds calculation takes the worldTransform into account.
func (self *Strip) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getBounds", args)}
}

// Adds a child to the container.
func (self *Strip) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Strip) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *Strip) SwapChildrenI(args ...interface{}) {
    self.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *Strip) GetChildIndexI(args ...interface{}) float64{
    return self.Call("getChildIndex", args).Float()
}

// Changes the position of an existing child in the display object container
func (self *Strip) SetChildIndexI(args ...interface{}) {
    self.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *Strip) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *Strip) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *Strip) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Strip) RemoveChildrenI(args ...interface{}) {
    self.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Strip) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Strip) SetStageReferenceI(args ...interface{}) {
    self.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *Strip) RemoveStageReferenceI(args ...interface{}) {
    self.Call("removeStageReference", args)
}

// Renders the object using the WebGL renderer
func (self *Strip) _renderWebGLI(args ...interface{}) {
    self.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *Strip) _renderCanvasI(args ...interface{}) {
    self.Call("_renderCanvas", args)
}
