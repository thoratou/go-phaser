// Automatic generation for Phaser.Canvas
// generated file Canvas.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// The Canvas class handles everything related to creating the `canvas` DOM tag that Phaser will use, 
// including styles, offset and aspect ratio.
type Canvas struct {
    *js.Object
}




// Creates a `canvas` DOM element. The element is not automatically added to the document.
func (self *Canvas) Create(parent interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("create", parent))
}

// Creates a `canvas` DOM element. The element is not automatically added to the document.
func (self *Canvas) Create1O(parent interface{}, width int) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("create", parent, width))
}

// Creates a `canvas` DOM element. The element is not automatically added to the document.
func (self *Canvas) Create2O(parent interface{}, width int, height int) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("create", parent, width, height))
}

// Creates a `canvas` DOM element. The element is not automatically added to the document.
func (self *Canvas) Create3O(parent interface{}, width int, height int, id string) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("create", parent, width, height, id))
}

// Creates a `canvas` DOM element. The element is not automatically added to the document.
func (self *Canvas) Create4O(parent interface{}, width int, height int, id string, skipPool bool) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("create", parent, width, height, id, skipPool))
}

// Creates a `canvas` DOM element. The element is not automatically added to the document.
func (self *Canvas) CreateI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("create", args))
}

// Sets the background color behind the canvas. This changes the canvas style property.
func (self *Canvas) SetBackgroundColor(canvas *dom.HTMLCanvasElement) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setBackgroundColor", canvas))
}

// Sets the background color behind the canvas. This changes the canvas style property.
func (self *Canvas) SetBackgroundColor1O(canvas *dom.HTMLCanvasElement, color string) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setBackgroundColor", canvas, color))
}

// Sets the background color behind the canvas. This changes the canvas style property.
func (self *Canvas) SetBackgroundColorI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setBackgroundColor", args))
}

// Sets the touch-action property on the canvas style. Can be used to disable default browser touch actions.
func (self *Canvas) SetTouchAction(canvas *dom.HTMLCanvasElement) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setTouchAction", canvas))
}

// Sets the touch-action property on the canvas style. Can be used to disable default browser touch actions.
func (self *Canvas) SetTouchAction1O(canvas *dom.HTMLCanvasElement, value string) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setTouchAction", canvas, value))
}

// Sets the touch-action property on the canvas style. Can be used to disable default browser touch actions.
func (self *Canvas) SetTouchActionI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setTouchAction", args))
}

// Sets the user-select property on the canvas style. Can be used to disable default browser selection actions.
func (self *Canvas) SetUserSelect(canvas *dom.HTMLCanvasElement) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setUserSelect", canvas))
}

// Sets the user-select property on the canvas style. Can be used to disable default browser selection actions.
func (self *Canvas) SetUserSelect1O(canvas *dom.HTMLCanvasElement, value string) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setUserSelect", canvas, value))
}

// Sets the user-select property on the canvas style. Can be used to disable default browser selection actions.
func (self *Canvas) SetUserSelectI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setUserSelect", args))
}

// Adds the given canvas element to the DOM. The canvas will be added as a child of the given parent.
// If no parent is given it will be added as a child of the document.body.
func (self *Canvas) AddToDOM(canvas *dom.HTMLCanvasElement, parent interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("addToDOM", canvas, parent))
}

// Adds the given canvas element to the DOM. The canvas will be added as a child of the given parent.
// If no parent is given it will be added as a child of the document.body.
func (self *Canvas) AddToDOM1O(canvas *dom.HTMLCanvasElement, parent interface{}, overflowHidden bool) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("addToDOM", canvas, parent, overflowHidden))
}

// Adds the given canvas element to the DOM. The canvas will be added as a child of the given parent.
// If no parent is given it will be added as a child of the document.body.
func (self *Canvas) AddToDOMI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("addToDOM", args))
}

// Removes the given canvas element from the DOM.
func (self *Canvas) RemoveFromDOM(canvas *dom.HTMLCanvasElement) {
    self.Object.Call("removeFromDOM", canvas)
}

// Removes the given canvas element from the DOM.
func (self *Canvas) RemoveFromDOMI(args ...interface{}) {
    self.Object.Call("removeFromDOM", args)
}

// Sets the transform of the given canvas to the matrix values provided.
func (self *Canvas) SetTransform(context *dom.CanvasRenderingContext2D, translateX int, translateY int, scaleX int, scaleY int, skewX int, skewY int) dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Call("setTransform", context, translateX, translateY, scaleX, scaleY, skewX, skewY))
}

// Sets the transform of the given canvas to the matrix values provided.
func (self *Canvas) SetTransformI(args ...interface{}) dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Call("setTransform", args))
}

// Sets the Image Smoothing property on the given context. Set to false to disable image smoothing.
// By default browsers have image smoothing enabled, which isn't always what you visually want, especially
// when using pixel art in a game. Note that this sets the property on the context itself, so that any image
// drawn to the context will be affected. This sets the property across all current browsers but support is
// patchy on earlier browsers, especially on mobile.
func (self *Canvas) SetSmoothingEnabled(context *dom.CanvasRenderingContext2D, value bool) dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Call("setSmoothingEnabled", context, value))
}

// Sets the Image Smoothing property on the given context. Set to false to disable image smoothing.
// By default browsers have image smoothing enabled, which isn't always what you visually want, especially
// when using pixel art in a game. Note that this sets the property on the context itself, so that any image
// drawn to the context will be affected. This sets the property across all current browsers but support is
// patchy on earlier browsers, especially on mobile.
func (self *Canvas) SetSmoothingEnabledI(args ...interface{}) dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Object.Call("setSmoothingEnabled", args))
}

// Gets the Smoothing Enabled vendor prefix being used on the given context, or null if not set.
func (self *Canvas) GetSmoothingPrefix(context *dom.CanvasRenderingContext2D) interface{}{
    return self.Object.Call("getSmoothingPrefix", context)
}

// Gets the Smoothing Enabled vendor prefix being used on the given context, or null if not set.
func (self *Canvas) GetSmoothingPrefixI(args ...interface{}) interface{}{
    return self.Object.Call("getSmoothingPrefix", args)
}

// Returns `true` if the given context has image smoothing enabled, otherwise returns `false`.
func (self *Canvas) GetSmoothingEnabled(context *dom.CanvasRenderingContext2D) bool{
    return self.Object.Call("getSmoothingEnabled", context).Bool()
}

// Returns `true` if the given context has image smoothing enabled, otherwise returns `false`.
func (self *Canvas) GetSmoothingEnabledI(args ...interface{}) bool{
    return self.Object.Call("getSmoothingEnabled", args).Bool()
}

// Sets the CSS image-rendering property on the given canvas to be 'crisp' (aka 'optimize contrast' on webkit).
// Note that if this doesn't given the desired result then see the setSmoothingEnabled.
func (self *Canvas) SetImageRenderingCrisp(canvas *dom.HTMLCanvasElement) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setImageRenderingCrisp", canvas))
}

// Sets the CSS image-rendering property on the given canvas to be 'crisp' (aka 'optimize contrast' on webkit).
// Note that if this doesn't given the desired result then see the setSmoothingEnabled.
func (self *Canvas) SetImageRenderingCrispI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setImageRenderingCrisp", args))
}

// Sets the CSS image-rendering property on the given canvas to be 'bicubic' (aka 'auto').
// Note that if this doesn't given the desired result then see the CanvasUtils.setSmoothingEnabled method.
func (self *Canvas) SetImageRenderingBicubic(canvas *dom.HTMLCanvasElement) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setImageRenderingBicubic", canvas))
}

// Sets the CSS image-rendering property on the given canvas to be 'bicubic' (aka 'auto').
// Note that if this doesn't given the desired result then see the CanvasUtils.setSmoothingEnabled method.
func (self *Canvas) SetImageRenderingBicubicI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("setImageRenderingBicubic", args))
}
