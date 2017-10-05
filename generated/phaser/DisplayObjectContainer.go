// Package phaser Automatic generation for PIXI.DisplayObjectContainer
// generated file DisplayObjectContainer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// DisplayObjectContainer A DisplayObjectContainer represents a collection of display objects.
// 
// It is the base class of all display objects that act as a container for other objects.
type DisplayObjectContainer struct {
    *js.Object
}

// NewDisplayObjectContainer A DisplayObjectContainer represents a collection of display objects.
// 
// It is the base class of all display objects that act as a container for other objects.
func NewDisplayObjectContainer() *DisplayObjectContainer {
    return &DisplayObjectContainer{js.Global.Get("PIXI").Get("DisplayObjectContainer").New()}
}
// NewDisplayObjectContainerI A DisplayObjectContainer represents a collection of display objects.
// 
// It is the base class of all display objects that act as a container for other objects.
func NewDisplayObjectContainerI(args ...interface{}) *DisplayObjectContainer {
    return &DisplayObjectContainer{js.Global.Get("PIXI").Get("DisplayObjectContainer").New(args)}
}



// DisplayObjectContainer Binding conversion method to DisplayObjectContainer point 
func ToDisplayObjectContainer(jsStruct interface{}) *DisplayObjectContainer {
    if object, ok := jsStruct.(*js.Object); ok {
		return &DisplayObjectContainer{Object: object}
	}
	return nil
}



// Children [read-only] The array of children of this container.
func (self *DisplayObjectContainer) Children() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// SetChildrenA [read-only] The array of children of this container.
func (self *DisplayObjectContainer) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// IgnoreChildInput If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// 
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// 
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *DisplayObjectContainer) IgnoreChildInput() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// SetIgnoreChildInputA If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// 
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// 
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *DisplayObjectContainer) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// Width The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *DisplayObjectContainer) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *DisplayObjectContainer) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *DisplayObjectContainer) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *DisplayObjectContainer) SetHeightA(member int) {
    self.Object.Set("height", member)
}


// AddChild Adds a child to the container.
func (self *DisplayObjectContainer) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// AddChildI Adds a child to the container.
func (self *DisplayObjectContainer) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// AddChildAt Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *DisplayObjectContainer) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// AddChildAtI Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *DisplayObjectContainer) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// SwapChildren Swaps the position of 2 Display Objects within this container.
func (self *DisplayObjectContainer) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// SwapChildrenI Swaps the position of 2 Display Objects within this container.
func (self *DisplayObjectContainer) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// GetChildIndex Returns the index position of a child DisplayObject instance
func (self *DisplayObjectContainer) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// GetChildIndexI Returns the index position of a child DisplayObject instance
func (self *DisplayObjectContainer) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// SetChildIndex Changes the position of an existing child in the display object container
func (self *DisplayObjectContainer) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// SetChildIndexI Changes the position of an existing child in the display object container
func (self *DisplayObjectContainer) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// GetChildAt Returns the child at the specified index
func (self *DisplayObjectContainer) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// GetChildAtI Returns the child at the specified index
func (self *DisplayObjectContainer) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// RemoveChild Removes a child from the container.
func (self *DisplayObjectContainer) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// RemoveChildI Removes a child from the container.
func (self *DisplayObjectContainer) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// RemoveChildAt Removes a child from the specified index position.
func (self *DisplayObjectContainer) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// RemoveChildAtI Removes a child from the specified index position.
func (self *DisplayObjectContainer) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// RemoveChildren Removes all children from this container that are within the begin and end indexes.
func (self *DisplayObjectContainer) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// RemoveChildrenI Removes all children from this container that are within the begin and end indexes.
func (self *DisplayObjectContainer) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// GetBounds Retrieves the global bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *DisplayObjectContainer) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// GetBounds1O Retrieves the global bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *DisplayObjectContainer) GetBounds1O(targetCoordinateSpace interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", targetCoordinateSpace)}
}

// GetBoundsI Retrieves the global bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *DisplayObjectContainer) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// GetLocalBounds Retrieves the non-global local bounds of the displayObjectContainer as a rectangle without any transformations. The calculation takes all visible children into consideration.
func (self *DisplayObjectContainer) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// GetLocalBoundsI Retrieves the non-global local bounds of the displayObjectContainer as a rectangle without any transformations. The calculation takes all visible children into consideration.
func (self *DisplayObjectContainer) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// Contains Determines whether the specified display object is a child of the DisplayObjectContainer instance or the instance itself.
func (self *DisplayObjectContainer) Contains(child *DisplayObject) bool{
    return self.Object.Call("contains", child).Bool()
}

// ContainsI Determines whether the specified display object is a child of the DisplayObjectContainer instance or the instance itself.
func (self *DisplayObjectContainer) ContainsI(args ...interface{}) bool{
    return self.Object.Call("contains", args).Bool()
}

// _renderWebGL Renders the object using the WebGL renderer
func (self *DisplayObjectContainer) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// _renderWebGLI Renders the object using the WebGL renderer
func (self *DisplayObjectContainer) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// _renderCanvas Renders the object using the Canvas renderer
func (self *DisplayObjectContainer) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// _renderCanvasI Renders the object using the Canvas renderer
func (self *DisplayObjectContainer) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

