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


// 
func NewStrip(texture *Texture, width int, height int) *Strip {
    return &Strip{js.Global.Call("PIXI.Strip", texture, width, height)}
}

// 
func NewStripI(args ...interface{}) *Strip {
    return &Strip{js.Global.Call("PIXI.Strip", args)}
}



// The texture of the strip
func (self *Strip) GetTextureA() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// The texture of the strip
func (self *Strip) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// Whether the strip is dirty or not
func (self *Strip) GetDirtyA() bool{
    return self.Object.Get("dirty").Bool()
}

// Whether the strip is dirty or not
func (self *Strip) SetDirtyA(member bool) {
    self.Object.Set("dirty", member)
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
func (self *Strip) GetBlendModeA() int{
    return self.Object.Get("blendMode").Int()
}

// The blend mode to be applied to the sprite. Set to PIXI.blendModes.NORMAL to remove any blend mode.
func (self *Strip) SetBlendModeA(member int) {
    self.Object.Set("blendMode", member)
}

// Triangles in canvas mode are automatically antialiased, use this value to force triangles to overlap a bit with each other.
func (self *Strip) GetCanvasPaddingA() int{
    return self.Object.Get("canvasPadding").Int()
}

// Triangles in canvas mode are automatically antialiased, use this value to force triangles to overlap a bit with each other.
func (self *Strip) SetCanvasPaddingA(member int) {
    self.Object.Set("canvasPadding", member)
}

// [read-only] The array of children of this container.
func (self *Strip) GetChildrenA() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// [read-only] The array of children of this container.
func (self *Strip) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Strip) GetIgnoreChildInputA() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Strip) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Strip) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Strip) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Strip) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Strip) SetHeightA(member int) {
    self.Object.Set("height", member)
}



// Renders a flat strip
func (self *Strip) RenderStripFlat(strip *Strip) {
    self.Object.Call("renderStripFlat", strip)
}

// Renders a flat strip
func (self *Strip) RenderStripFlatI(args ...interface{}) {
    self.Object.Call("renderStripFlat", args)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *Strip) OnTextureUpdate(event interface{}) {
    self.Object.Call("onTextureUpdate", event)
}

// When the texture is updated, this event will fire to update the scale and frame
func (self *Strip) OnTextureUpdateI(args ...interface{}) {
    self.Object.Call("onTextureUpdate", args)
}

// Returns the bounds of the mesh as a rectangle. The bounds calculation takes the worldTransform into account.
func (self *Strip) GetBounds(matrix *Matrix) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", matrix)}
}

// Returns the bounds of the mesh as a rectangle. The bounds calculation takes the worldTransform into account.
func (self *Strip) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Adds a child to the container.
func (self *Strip) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// Adds a child to the container.
func (self *Strip) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Strip) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Strip) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *Strip) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// Swaps the position of 2 Display Objects within this container.
func (self *Strip) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *Strip) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// Returns the index position of a child DisplayObject instance
func (self *Strip) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *Strip) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// Changes the position of an existing child in the display object container
func (self *Strip) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *Strip) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// Returns the child at the specified index
func (self *Strip) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *Strip) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// Removes a child from the container.
func (self *Strip) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *Strip) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// Removes a child from the specified index position.
func (self *Strip) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Strip) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Strip) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Strip) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Strip) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Strip) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Strip) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *Strip) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// Removes the current stage reference from the container and all of its children.
func (self *Strip) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// Renders the object using the WebGL renderer
func (self *Strip) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// Renders the object using the WebGL renderer
func (self *Strip) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *Strip) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// Renders the object using the Canvas renderer
func (self *Strip) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}
