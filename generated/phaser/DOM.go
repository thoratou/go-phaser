// Automatic generation for Phaser.DOM
// generated file DOM.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// DOM utility class.
// 
// Provides a useful Window and Element functions as well as cross-browser compatibility buffer.
// 
// Some code originally derived from {@link https://github.com/ryanve/verge verge}.
// Some parts were inspired by the research of Ryan Van Etten, released under MIT License 2013.
type DOM struct {
    *js.Object
}


// DOM utility class.
// 
// Provides a useful Window and Element functions as well as cross-browser compatibility buffer.
// 
// Some code originally derived from {@link https://github.com/ryanve/verge verge}.
// Some parts were inspired by the research of Ryan Van Etten, released under MIT License 2013.
func NewDOM() *DOM {
    return &DOM{js.Global.Call("Phaser.DOM")}
}

// DOM utility class.
// 
// Provides a useful Window and Element functions as well as cross-browser compatibility buffer.
// 
// Some code originally derived from {@link https://github.com/ryanve/verge verge}.
// Some parts were inspired by the research of Ryan Van Etten, released under MIT License 2013.
func NewDOMI(args ...interface{}) *DOM {
    return &DOM{js.Global.Call("Phaser.DOM", args)}
}



// The bounds of the Visual viewport, as discussed in 
// {@link http://www.quirksmode.org/mobile/viewports.html A tale of two viewports — part one}
// with one difference: the viewport size _excludes_ scrollbars, as found on some desktop browsers.   
// 
// Supported mobile:
//   iOS/Safari, Android 4, IE10, Firefox OS (maybe not Firefox Android), Opera Mobile 16
// 
// The properties change dynamically.
func (self *DOM) GetVisualBoundsA() *Rectangle{
    return &Rectangle{self.Object.Get("visualBounds")}
}

// The bounds of the Visual viewport, as discussed in 
// {@link http://www.quirksmode.org/mobile/viewports.html A tale of two viewports — part one}
// with one difference: the viewport size _excludes_ scrollbars, as found on some desktop browsers.   
// 
// Supported mobile:
//   iOS/Safari, Android 4, IE10, Firefox OS (maybe not Firefox Android), Opera Mobile 16
// 
// The properties change dynamically.
func (self *DOM) SetVisualBoundsA(member *Rectangle) {
    self.Object.Set("visualBounds", member)
}

// The bounds of the Layout viewport, as discussed in 
// {@link http://www.quirksmode.org/mobile/viewports2.html A tale of two viewports — part two};
// but honoring the constraints as specified applicable viewport meta-tag.
// 
// The bounds returned are not guaranteed to be fully aligned with CSS media queries (see
// {@link http://www.matanich.com/2013/01/07/viewport-size/ What size is my viewport?}).
// 
// This is _not_ representative of the Visual bounds: in particular the non-primary axis will
// generally be significantly larger than the screen height on mobile devices when running with a
// constrained viewport.
// 
// The properties change dynamically.
func (self *DOM) GetLayoutBoundsA() *Rectangle{
    return &Rectangle{self.Object.Get("layoutBounds")}
}

// The bounds of the Layout viewport, as discussed in 
// {@link http://www.quirksmode.org/mobile/viewports2.html A tale of two viewports — part two};
// but honoring the constraints as specified applicable viewport meta-tag.
// 
// The bounds returned are not guaranteed to be fully aligned with CSS media queries (see
// {@link http://www.matanich.com/2013/01/07/viewport-size/ What size is my viewport?}).
// 
// This is _not_ representative of the Visual bounds: in particular the non-primary axis will
// generally be significantly larger than the screen height on mobile devices when running with a
// constrained viewport.
// 
// The properties change dynamically.
func (self *DOM) SetLayoutBoundsA(member *Rectangle) {
    self.Object.Set("layoutBounds", member)
}

// The size of the document / Layout viewport.
// 
// This incorrectly reports the dimensions in IE.
// 
// The properties change dynamically.
func (self *DOM) GetDocumentBoundsA() *Rectangle{
    return &Rectangle{self.Object.Get("documentBounds")}
}

// The size of the document / Layout viewport.
// 
// This incorrectly reports the dimensions in IE.
// 
// The properties change dynamically.
func (self *DOM) SetDocumentBoundsA(member *Rectangle) {
    self.Object.Set("documentBounds", member)
}

// A cross-browser window.scrollX.
func (self *DOM) GetScrollXA() int{
    return self.Object.Get("scrollX").Int()
}

// A cross-browser window.scrollX.
func (self *DOM) SetScrollXA(member int) {
    self.Object.Set("scrollX", member)
}

// A cross-browser window.scrollY.
func (self *DOM) GetScrollYA() int{
    return self.Object.Get("scrollY").Int()
}

// A cross-browser window.scrollY.
func (self *DOM) SetScrollYA(member int) {
    self.Object.Set("scrollY", member)
}



// Get the [absolute] position of the element relative to the Document.
// 
// The value may vary slightly as the page is scrolled due to rounding errors.
func (self *DOM) GetOffset(element *DOMElement) *Point{
    return &Point{self.Object.Call("getOffset", element)}
}

// Get the [absolute] position of the element relative to the Document.
// 
// The value may vary slightly as the page is scrolled due to rounding errors.
func (self *DOM) GetOffset1O(element *DOMElement, point *Point) *Point{
    return &Point{self.Object.Call("getOffset", element, point)}
}

// Get the [absolute] position of the element relative to the Document.
// 
// The value may vary slightly as the page is scrolled due to rounding errors.
func (self *DOM) GetOffsetI(args ...interface{}) *Point{
    return &Point{self.Object.Call("getOffset", args)}
}

// A cross-browser element.getBoundingClientRect method with optional cushion.
// 
// Returns a plain object containing the properties `top/bottom/left/right/width/height` with respect to the top-left corner of the current viewport.
// Its properties match the native rectangle.
// The cushion parameter is an amount of pixels (+/-) to cushion the element.
// It adjusts the measurements such that it is possible to detect when an element is near the viewport.
func (self *DOM) GetBounds(element interface{}) interface{}{
    return self.Object.Call("getBounds", element)
}

// A cross-browser element.getBoundingClientRect method with optional cushion.
// 
// Returns a plain object containing the properties `top/bottom/left/right/width/height` with respect to the top-left corner of the current viewport.
// Its properties match the native rectangle.
// The cushion parameter is an amount of pixels (+/-) to cushion the element.
// It adjusts the measurements such that it is possible to detect when an element is near the viewport.
func (self *DOM) GetBounds1O(element interface{}, cushion int) interface{}{
    return self.Object.Call("getBounds", element, cushion)
}

// A cross-browser element.getBoundingClientRect method with optional cushion.
// 
// Returns a plain object containing the properties `top/bottom/left/right/width/height` with respect to the top-left corner of the current viewport.
// Its properties match the native rectangle.
// The cushion parameter is an amount of pixels (+/-) to cushion the element.
// It adjusts the measurements such that it is possible to detect when an element is near the viewport.
func (self *DOM) GetBoundsI(args ...interface{}) interface{}{
    return self.Object.Call("getBounds", args)
}

// Calibrates element coordinates for `inLayoutViewport` checks.
func (self *DOM) Calibrate(coords interface{}) interface{}{
    return self.Object.Call("calibrate", coords)
}

// Calibrates element coordinates for `inLayoutViewport` checks.
func (self *DOM) Calibrate1O(coords interface{}, cushion int) interface{}{
    return self.Object.Call("calibrate", coords, cushion)
}

// Calibrates element coordinates for `inLayoutViewport` checks.
func (self *DOM) CalibrateI(args ...interface{}) interface{}{
    return self.Object.Call("calibrate", args)
}

// Get the Visual viewport aspect ratio (or the aspect ratio of an object or element)
func (self *DOM) GetAspectRatio() int{
    return self.Object.Call("getAspectRatio").Int()
}

// Get the Visual viewport aspect ratio (or the aspect ratio of an object or element)
func (self *DOM) GetAspectRatio1O(object interface{}) int{
    return self.Object.Call("getAspectRatio", object).Int()
}

// Get the Visual viewport aspect ratio (or the aspect ratio of an object or element)
func (self *DOM) GetAspectRatioI(args ...interface{}) int{
    return self.Object.Call("getAspectRatio", args).Int()
}

// Tests if the given DOM element is within the Layout viewport.
// 
// The optional cushion parameter allows you to specify a distance.
// 
// inLayoutViewport(element, 100) is `true` if the element is in the viewport or 100px near it.
// inLayoutViewport(element, -100) is `true` if the element is in the viewport or at least 100px near it.
func (self *DOM) InLayoutViewport(element interface{}) bool{
    return self.Object.Call("inLayoutViewport", element).Bool()
}

// Tests if the given DOM element is within the Layout viewport.
// 
// The optional cushion parameter allows you to specify a distance.
// 
// inLayoutViewport(element, 100) is `true` if the element is in the viewport or 100px near it.
// inLayoutViewport(element, -100) is `true` if the element is in the viewport or at least 100px near it.
func (self *DOM) InLayoutViewport1O(element interface{}, cushion int) bool{
    return self.Object.Call("inLayoutViewport", element, cushion).Bool()
}

// Tests if the given DOM element is within the Layout viewport.
// 
// The optional cushion parameter allows you to specify a distance.
// 
// inLayoutViewport(element, 100) is `true` if the element is in the viewport or 100px near it.
// inLayoutViewport(element, -100) is `true` if the element is in the viewport or at least 100px near it.
func (self *DOM) InLayoutViewportI(args ...interface{}) bool{
    return self.Object.Call("inLayoutViewport", args).Bool()
}

// Returns the device screen orientation.
// 
// Orientation values: 'portrait-primary', 'landscape-primary', 'portrait-secondary', 'landscape-secondary'.
// 
// Order of resolving:
// - Screen Orientation API, or variation of - Future track. Most desktop and mobile browsers.
// - Screen size ratio check - If fallback is 'screen', suited for desktops.
// - Viewport size ratio check - If fallback is 'viewport', suited for mobile.
// - window.orientation - If fallback is 'window.orientation', works iOS and probably most Android; non-recommended track.
// - Media query
// - Viewport size ratio check (probably only IE9 and legacy mobile gets here..)
// 
// See
// - https://w3c.github.io/screen-orientation/ (conflicts with mozOrientation/msOrientation)
// - https://developer.mozilla.org/en-US/docs/Web/API/Screen.orientation (mozOrientation)
// - http://msdn.microsoft.com/en-us/library/ie/dn342934(v=vs.85).aspx
// - https://developer.mozilla.org/en-US/docs/Web/Guide/CSS/Testing_media_queries
// - http://stackoverflow.com/questions/4917664/detect-viewport-orientation
// - http://www.matthewgifford.com/blog/2011/12/22/a-misconception-about-window-orientation
func (self *DOM) GetScreenOrientation() {
    self.Object.Call("getScreenOrientation")
}

// Returns the device screen orientation.
// 
// Orientation values: 'portrait-primary', 'landscape-primary', 'portrait-secondary', 'landscape-secondary'.
// 
// Order of resolving:
// - Screen Orientation API, or variation of - Future track. Most desktop and mobile browsers.
// - Screen size ratio check - If fallback is 'screen', suited for desktops.
// - Viewport size ratio check - If fallback is 'viewport', suited for mobile.
// - window.orientation - If fallback is 'window.orientation', works iOS and probably most Android; non-recommended track.
// - Media query
// - Viewport size ratio check (probably only IE9 and legacy mobile gets here..)
// 
// See
// - https://w3c.github.io/screen-orientation/ (conflicts with mozOrientation/msOrientation)
// - https://developer.mozilla.org/en-US/docs/Web/API/Screen.orientation (mozOrientation)
// - http://msdn.microsoft.com/en-us/library/ie/dn342934(v=vs.85).aspx
// - https://developer.mozilla.org/en-US/docs/Web/Guide/CSS/Testing_media_queries
// - http://stackoverflow.com/questions/4917664/detect-viewport-orientation
// - http://www.matthewgifford.com/blog/2011/12/22/a-misconception-about-window-orientation
func (self *DOM) GetScreenOrientation1O(primaryFallback string) {
    self.Object.Call("getScreenOrientation", primaryFallback)
}

// Returns the device screen orientation.
// 
// Orientation values: 'portrait-primary', 'landscape-primary', 'portrait-secondary', 'landscape-secondary'.
// 
// Order of resolving:
// - Screen Orientation API, or variation of - Future track. Most desktop and mobile browsers.
// - Screen size ratio check - If fallback is 'screen', suited for desktops.
// - Viewport size ratio check - If fallback is 'viewport', suited for mobile.
// - window.orientation - If fallback is 'window.orientation', works iOS and probably most Android; non-recommended track.
// - Media query
// - Viewport size ratio check (probably only IE9 and legacy mobile gets here..)
// 
// See
// - https://w3c.github.io/screen-orientation/ (conflicts with mozOrientation/msOrientation)
// - https://developer.mozilla.org/en-US/docs/Web/API/Screen.orientation (mozOrientation)
// - http://msdn.microsoft.com/en-us/library/ie/dn342934(v=vs.85).aspx
// - https://developer.mozilla.org/en-US/docs/Web/Guide/CSS/Testing_media_queries
// - http://stackoverflow.com/questions/4917664/detect-viewport-orientation
// - http://www.matthewgifford.com/blog/2011/12/22/a-misconception-about-window-orientation
func (self *DOM) GetScreenOrientationI(args ...interface{}) {
    self.Object.Call("getScreenOrientation", args)
}
