// Automatic generation for Phaser.Stage
// generated file Stage.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Stage controls root level display objects upon which everything is displayed.
// It also handles browser visibility handling and the pausing due to loss of focus.
type Stage struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *Stage) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Stage) SetGame(member *Game) {
    self.Set("game", member)
}

// The name of this object.
func (self *Stage) GetName() string{
    return self.Get("name").String()
}

// The name of this object.
func (self *Stage) SetName(member string) {
    self.Set("name", member)
}

// By default if the browser tab loses focus the game will pause.
// You can stop that behavior by setting this property to true.
// Note that the browser can still elect to pause your game if it wishes to do so,
// for example swapping to another browser tab. This will cause the RAF callback to halt,
// effectively pausing your game, even though no in-game pause event is triggered if you enable this property.
func (self *Stage) GetDisableVisibilityChange() bool{
    return self.Get("disableVisibilityChange").Bool()
}

// By default if the browser tab loses focus the game will pause.
// You can stop that behavior by setting this property to true.
// Note that the browser can still elect to pause your game if it wishes to do so,
// for example swapping to another browser tab. This will cause the RAF callback to halt,
// effectively pausing your game, even though no in-game pause event is triggered if you enable this property.
func (self *Stage) SetDisableVisibilityChange(member bool) {
    self.Set("disableVisibilityChange", member)
}

// If exists is true the Stage and all children are updated, otherwise it is skipped.
func (self *Stage) GetExists() bool{
    return self.Get("exists").Bool()
}

// If exists is true the Stage and all children are updated, otherwise it is skipped.
func (self *Stage) SetExists(member bool) {
    self.Set("exists", member)
}

// Reset each frame, keeps a count of the total number of objects updated.
func (self *Stage) GetCurrentRenderOrderID() int{
    return self.Get("currentRenderOrderID").Int()
}

// Reset each frame, keeps a count of the total number of objects updated.
func (self *Stage) SetCurrentRenderOrderID(member int) {
    self.Set("currentRenderOrderID", member)
}

// Gets and sets the background color of the stage. The color can be given as a number: 0xff0000 or a hex string: '#ff0000'
func (self *Stage) GetBackgroundColor() interface{}{
    return self.Get("backgroundColor")
}

// Gets and sets the background color of the stage. The color can be given as a number: 0xff0000 or a hex string: '#ff0000'
func (self *Stage) SetBackgroundColor(member interface{}) {
    self.Set("backgroundColor", member)
}

// Enable or disable texture smoothing for all objects on this Stage. Only works for bitmap/image textures. Smoothing is enabled by default. Set to true to smooth all sprites rendered on this Stage, or false to disable smoothing (great for pixel art)
func (self *Stage) GetSmoothed() bool{
    return self.Get("smoothed").Bool()
}

// Enable or disable texture smoothing for all objects on this Stage. Only works for bitmap/image textures. Smoothing is enabled by default. Set to true to smooth all sprites rendered on this Stage, or false to disable smoothing (great for pixel art)
func (self *Stage) SetSmoothed(member bool) {
    self.Set("smoothed", member)
}

// [read-only] The array of children of this container.
func (self *Stage) GetChildren() []DisplayObject{
	array := self.Get("children")
	length := array.Length()
	out := make([]DisplayObject, length, length)
	for i := 0; i < length; i++ {
		out[i] = DisplayObject{array.Index(i)}
	}
	return out
}

// [read-only] The array of children of this container.
func (self *Stage) SetChildren(member []DisplayObject) {
    self.Set("children", member)
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Stage) GetIgnoreChildInput() bool{
    return self.Get("ignoreChildInput").Bool()
}

// If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Stage) SetIgnoreChildInput(member bool) {
    self.Set("ignoreChildInput", member)
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Stage) GetWidth() int{
    return self.Get("width").Int()
}

// The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Stage) SetWidth(member int) {
    self.Set("width", member)
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Stage) GetHeight() int{
    return self.Get("height").Int()
}

// The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Stage) SetHeight(member int) {
    self.Set("height", member)
}



// Parses a Game configuration object.
func (self *Stage) ParseConfigI(args ...interface{}) {
    self.Call("parseConfig", args)
}

// Initialises the stage and adds the event listeners.
func (self *Stage) BootI(args ...interface{}) {
    self.Call("boot", args)
}

// This is called automatically after the plugins preUpdate and before the State.update.
// Most objects have preUpdate methods and it's where initial movement and positioning is done.
func (self *Stage) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// This is called automatically after the State.update, but before particles or plugins update.
func (self *Stage) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// This is called automatically before the renderer runs and after the plugins have updated.
// In postUpdate this is where all the final physics calculations and object positioning happens.
// The objects are processed in the order of the display list.
func (self *Stage) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}

// Updates the transforms for all objects on the display list.
// This overrides the Pixi default as we don't need the interactionManager, but do need the game property check.
func (self *Stage) UpdateTransformI(args ...interface{}) {
    self.Call("updateTransform", args)
}

// Starts a page visibility event listener running, or window.onpagehide/onpageshow if not supported by the browser.
// Also listens for window.onblur and window.onfocus.
func (self *Stage) CheckVisibilityI(args ...interface{}) {
    self.Call("checkVisibility", args)
}

// This method is called when the document visibility is changed.
func (self *Stage) VisibilityChangeI(args ...interface{}) {
    self.Call("visibilityChange", args)
}

// Sets the background color for the Stage.
// 
// The color can be given as a hex string (`'#RRGGBB'`), a CSS color string (`'rgb(r,g,b)'`), or a numeric value (`0xRRGGBB`).
// 
// An alpha channel is _not_ supported and will be ignored.
// 
// If you've set your game to be transparent then calls to setBackgroundColor are ignored.
func (self *Stage) SetBackgroundColorI(args ...interface{}) {
    self.Call("setBackgroundColor", args)
}

// Destroys the Stage and removes event listeners.
func (self *Stage) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Adds a child to the container.
func (self *Stage) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChild", args)}
}

// Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Stage) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("addChildAt", args)}
}

// Swaps the position of 2 Display Objects within this container.
func (self *Stage) SwapChildrenI(args ...interface{}) {
    self.Call("swapChildren", args)
}

// Returns the index position of a child DisplayObject instance
func (self *Stage) GetChildIndexI(args ...interface{}) int{
    return self.Call("getChildIndex", args).Int()
}

// Changes the position of an existing child in the display object container
func (self *Stage) SetChildIndexI(args ...interface{}) {
    self.Call("setChildIndex", args)
}

// Returns the child at the specified index
func (self *Stage) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("getChildAt", args)}
}

// Removes a child from the container.
func (self *Stage) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChild", args)}
}

// Removes a child from the specified index position.
func (self *Stage) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Call("removeChildAt", args)}
}

// Removes all children from this container that are within the begin and end indexes.
func (self *Stage) RemoveChildrenI(args ...interface{}) {
    self.Call("removeChildren", args)
}

// Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *Stage) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getBounds", args)}
}

// Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Stage) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Call("getLocalBounds", args)}
}

// Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Stage) SetStageReferenceI(args ...interface{}) {
    self.Call("setStageReference", args)
}

// Removes the current stage reference from the container and all of its children.
func (self *Stage) RemoveStageReferenceI(args ...interface{}) {
    self.Call("removeStageReference", args)
}

// Renders the object using the WebGL renderer
func (self *Stage) _renderWebGLI(args ...interface{}) {
    self.Call("_renderWebGL", args)
}

// Renders the object using the Canvas renderer
func (self *Stage) _renderCanvasI(args ...interface{}) {
    self.Call("_renderCanvas", args)
}
