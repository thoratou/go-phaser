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
func (self *Canvas) CreateI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("create", args))
}

// Sets the background color behind the canvas. This changes the canvas style property.
func (self *Canvas) SetBackgroundColorI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("setBackgroundColor", args))
}

// Sets the touch-action property on the canvas style. Can be used to disable default browser touch actions.
func (self *Canvas) SetTouchActionI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("setTouchAction", args))
}

// Sets the user-select property on the canvas style. Can be used to disable default browser selection actions.
func (self *Canvas) SetUserSelectI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("setUserSelect", args))
}

// Adds the given canvas element to the DOM. The canvas will be added as a child of the given parent.
// If no parent is given it will be added as a child of the document.body.
func (self *Canvas) AddToDOMI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("addToDOM", args))
}

// Removes the given canvas element from the DOM.
func (self *Canvas) RemoveFromDOMI(args ...interface{}) {
    self.Call("removeFromDOM", args)
}

// Sets the transform of the given canvas to the matrix values provided.
func (self *Canvas) SetTransformI(args ...interface{}) dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Call("setTransform", args))
}

// Sets the Image Smoothing property on the given context. Set to false to disable image smoothing.
// By default browsers have image smoothing enabled, which isn't always what you visually want, especially
// when using pixel art in a game. Note that this sets the property on the context itself, so that any image
// drawn to the context will be affected. This sets the property across all current browsers but support is
// patchy on earlier browsers, especially on mobile.
func (self *Canvas) SetSmoothingEnabledI(args ...interface{}) dom.CanvasRenderingContext2D{
    return WrapCanvasRenderingContext2D(self.Call("setSmoothingEnabled", args))
}

// Gets the Smoothing Enabled vendor prefix being used on the given context, or null if not set.
func (self *Canvas) GetSmoothingPrefixI(args ...interface{}) interface{}{
    return self.Call("getSmoothingPrefix", args)
}

// Returns `true` if the given context has image smoothing enabled, otherwise returns `false`.
func (self *Canvas) GetSmoothingEnabledI(args ...interface{}) bool{
    return self.Call("getSmoothingEnabled", args).Bool()
}

// Sets the CSS image-rendering property on the given canvas to be 'crisp' (aka 'optimize contrast' on webkit).
// Note that if this doesn't given the desired result then see the setSmoothingEnabled.
func (self *Canvas) SetImageRenderingCrispI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("setImageRenderingCrisp", args))
}

// Sets the CSS image-rendering property on the given canvas to be 'bicubic' (aka 'auto').
// Note that if this doesn't given the desired result then see the CanvasUtils.setSmoothingEnabled method.
func (self *Canvas) SetImageRenderingBicubicI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Call("setImageRenderingBicubic", args))
}
