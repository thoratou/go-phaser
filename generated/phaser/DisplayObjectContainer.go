// Automatic generation for PIXI.DisplayObjectContainer
// generated file DisplayObjectContainer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A DisplayObjectContainer represents a collection of display objects.
// It is the base class of all display objects that act as a container for other objects.
type DisplayObjectContainer struct {
    *js.Object
}


// [read-only] The array of children of this container.
func (self *DisplayObjectContainer) GetChildrenA() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// [read-only] The array of children of this container.
func (self *DisplayObjectContainer) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *DisplayObjectContainer) GetIgnoreChildInputA() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *DisplayObjectContainer) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *DisplayObjectContainer) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *DisplayObjectContainer) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *DisplayObjectContainer) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *DisplayObjectContainer) SetHeightA(member int) {
    self.Object.Set("height", member)
}



// Adds a child to the container.
func (self *DisplayObjectContainer) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// Adds a child to the container.
func (self *DisplayObjectContainer) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *DisplayObjectContainer) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *DisplayObjectContainer) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *DisplayObjectContainer) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// Swaps the position of 2 Display Objects within this container.
func (self *DisplayObjectContainer) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *DisplayObjectContainer) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// Returns the index position of a child DisplayObject instance
func (self *DisplayObjectContainer) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *DisplayObjectContainer) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// Changes the position of an existing child in the display object container
func (self *DisplayObjectContainer) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *DisplayObjectContainer) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// Returns the child at the specified index
func (self *DisplayObjectContainer) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *DisplayObjectContainer) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// Removes a child from the container.
func (self *DisplayObjectContainer) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *DisplayObjectContainer) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// Removes a child from the specified index position.
func (self *DisplayObjectContainer) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *DisplayObjectContainer) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// Removes all children from this container that are within the begin and end indexes.
func (self *DisplayObjectContainer) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *DisplayObjectContainer) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *DisplayObjectContainer) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *DisplayObjectContainer) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *DisplayObjectContainer) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *DisplayObjectContainer) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *DisplayObjectContainer) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *DisplayObjectContainer) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// Removes the current stage reference from the container and all of its children.
func (self *DisplayObjectContainer) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// Renders the object using the WebGL renderer
func (self *DisplayObjectContainer) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// Renders the object using the WebGL renderer
func (self *DisplayObjectContainer) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *DisplayObjectContainer) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// Renders the object using the Canvas renderer
func (self *DisplayObjectContainer) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}
