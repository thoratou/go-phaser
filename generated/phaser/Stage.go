// Package phaser Automatic generation for Phaser.Stage
// generated file Stage.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Stage The Stage controls root level display objects upon which everything is displayed.
// It also handles browser visibility handling and the pausing due to loss of focus.
type Stage struct {
    *js.Object
}

// NewStage The Stage controls root level display objects upon which everything is displayed.
// It also handles browser visibility handling and the pausing due to loss of focus.
func NewStage(game *Game) *Stage {
    return &Stage{js.Global.Get("Phaser").Get("Stage").New(game)}
}
// NewStageI The Stage controls root level display objects upon which everything is displayed.
// It also handles browser visibility handling and the pausing due to loss of focus.
func NewStageI(args ...interface{}) *Stage {
    return &Stage{js.Global.Get("Phaser").Get("Stage").New(args)}
}



// Stage Binding conversion method to Stage point 
func ToStage(jsStruct interface{}) *Stage {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Stage{Object: object}
	}
	return nil
}



// Game A reference to the currently running Game.
func (self *Stage) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Stage) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Name The name of this object.
func (self *Stage) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA The name of this object.
func (self *Stage) SetNameA(member string) {
    self.Object.Set("name", member)
}

// DisableVisibilityChange By default if the browser tab loses focus the game will pause.
// You can stop that behavior by setting this property to true.
// Note that the browser can still elect to pause your game if it wishes to do so,
// for example swapping to another browser tab. This will cause the RAF callback to halt,
// effectively pausing your game, even though no in-game pause event is triggered if you enable this property.
func (self *Stage) DisableVisibilityChange() bool{
    return self.Object.Get("disableVisibilityChange").Bool()
}

// SetDisableVisibilityChangeA By default if the browser tab loses focus the game will pause.
// You can stop that behavior by setting this property to true.
// Note that the browser can still elect to pause your game if it wishes to do so,
// for example swapping to another browser tab. This will cause the RAF callback to halt,
// effectively pausing your game, even though no in-game pause event is triggered if you enable this property.
func (self *Stage) SetDisableVisibilityChangeA(member bool) {
    self.Object.Set("disableVisibilityChange", member)
}

// Exists If exists is true the Stage and all children are updated, otherwise it is skipped.
func (self *Stage) Exists() bool{
    return self.Object.Get("exists").Bool()
}

// SetExistsA If exists is true the Stage and all children are updated, otherwise it is skipped.
func (self *Stage) SetExistsA(member bool) {
    self.Object.Set("exists", member)
}

// CurrentRenderOrderID Reset each frame, keeps a count of the total number of objects updated.
func (self *Stage) CurrentRenderOrderID() int{
    return self.Object.Get("currentRenderOrderID").Int()
}

// SetCurrentRenderOrderIDA Reset each frame, keeps a count of the total number of objects updated.
func (self *Stage) SetCurrentRenderOrderIDA(member int) {
    self.Object.Set("currentRenderOrderID", member)
}

// BackgroundColor Gets and sets the background color of the stage. The color can be given as a number: 0xff0000 or a hex string: '#ff0000'
func (self *Stage) BackgroundColor() interface{}{
    return self.Object.Get("backgroundColor")
}

// SetBackgroundColorA Gets and sets the background color of the stage. The color can be given as a number: 0xff0000 or a hex string: '#ff0000'
func (self *Stage) SetBackgroundColorA(member interface{}) {
    self.Object.Set("backgroundColor", member)
}

// Smoothed Enable or disable texture smoothing for all objects on this Stage. Only works for bitmap/image textures. Smoothing is enabled by default. Set to true to smooth all sprites rendered on this Stage, or false to disable smoothing (great for pixel art)
func (self *Stage) Smoothed() bool{
    return self.Object.Get("smoothed").Bool()
}

// SetSmoothedA Enable or disable texture smoothing for all objects on this Stage. Only works for bitmap/image textures. Smoothing is enabled by default. Set to true to smooth all sprites rendered on this Stage, or false to disable smoothing (great for pixel art)
func (self *Stage) SetSmoothedA(member bool) {
    self.Object.Set("smoothed", member)
}

// Children [read-only] The array of children of this container.
func (self *Stage) Children() []DisplayObject{
	array00 := self.Object.Get("children")
	length00 := array00.Length()
	out00 := make([]DisplayObject, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = DisplayObject{array00.Index(i00)}
	}
	return out00
}

// SetChildrenA [read-only] The array of children of this container.
func (self *Stage) SetChildrenA(member []DisplayObject) {
    self.Object.Set("children", member)
}

// IgnoreChildInput If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Stage) IgnoreChildInput() bool{
    return self.Object.Get("ignoreChildInput").Bool()
}

// SetIgnoreChildInputA If `ignoreChildInput`  is `false` it will allow this objects _children_ to be considered as valid for Input events.
// 
// If this property is `true` then the children will _not_ be considered as valid for Input events.
// 
// Note that this property isn't recursive: only immediate children are influenced, it doesn't scan further down.
func (self *Stage) SetIgnoreChildInputA(member bool) {
    self.Object.Set("ignoreChildInput", member)
}

// Width The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Stage) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The width of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Stage) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Stage) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The height of the displayObjectContainer, setting this will actually modify the scale to achieve the value set
func (self *Stage) SetHeightA(member int) {
    self.Object.Set("height", member)
}


// ParseConfig Parses a Game configuration object.
func (self *Stage) ParseConfig(config interface{}) {
    self.Object.Call("parseConfig", config)
}

// ParseConfigI Parses a Game configuration object.
func (self *Stage) ParseConfigI(args ...interface{}) {
    self.Object.Call("parseConfig", args)
}

// Boot Initialises the stage and adds the event listeners.
func (self *Stage) Boot() {
    self.Object.Call("boot")
}

// BootI Initialises the stage and adds the event listeners.
func (self *Stage) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// PreUpdate This is called automatically after the plugins preUpdate and before the State.update.
// Most objects have preUpdate methods and it's where initial movement and positioning is done.
func (self *Stage) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI This is called automatically after the plugins preUpdate and before the State.update.
// Most objects have preUpdate methods and it's where initial movement and positioning is done.
func (self *Stage) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Update This is called automatically after the State.update, but before particles or plugins update.
func (self *Stage) Update() {
    self.Object.Call("update")
}

// UpdateI This is called automatically after the State.update, but before particles or plugins update.
func (self *Stage) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// PostUpdate This is called automatically before the renderer runs and after the plugins have updated.
// In postUpdate this is where all the final physics calculations and object positioning happens.
// The objects are processed in the order of the display list.
func (self *Stage) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI This is called automatically before the renderer runs and after the plugins have updated.
// In postUpdate this is where all the final physics calculations and object positioning happens.
// The objects are processed in the order of the display list.
func (self *Stage) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// UpdateTransform Updates the transforms for all objects on the display list.
// This overrides the Pixi default as we don't need the interactionManager, but do need the game property check.
func (self *Stage) UpdateTransform() {
    self.Object.Call("updateTransform")
}

// UpdateTransformI Updates the transforms for all objects on the display list.
// This overrides the Pixi default as we don't need the interactionManager, but do need the game property check.
func (self *Stage) UpdateTransformI(args ...interface{}) {
    self.Object.Call("updateTransform", args)
}

// CheckVisibility Starts a page visibility event listener running, or window.onpagehide/onpageshow if not supported by the browser.
// Also listens for window.onblur and window.onfocus.
func (self *Stage) CheckVisibility() {
    self.Object.Call("checkVisibility")
}

// CheckVisibilityI Starts a page visibility event listener running, or window.onpagehide/onpageshow if not supported by the browser.
// Also listens for window.onblur and window.onfocus.
func (self *Stage) CheckVisibilityI(args ...interface{}) {
    self.Object.Call("checkVisibility", args)
}

// VisibilityChange This method is called when the document visibility is changed.
func (self *Stage) VisibilityChange(event *Event) {
    self.Object.Call("visibilityChange", event)
}

// VisibilityChangeI This method is called when the document visibility is changed.
func (self *Stage) VisibilityChangeI(args ...interface{}) {
    self.Object.Call("visibilityChange", args)
}

// SetBackgroundColor Sets the background color for the Stage.
// 
// The color can be given as a hex string (`'#RRGGBB'`), a CSS color string (`'rgb(r,g,b)'`), or a numeric value (`0xRRGGBB`).
// 
// An alpha channel is _not_ supported and will be ignored.
// 
// If you've set your game to be transparent then calls to setBackgroundColor are ignored.
func (self *Stage) SetBackgroundColor(color interface{}) {
    self.Object.Call("setBackgroundColor", color)
}

// SetBackgroundColorI Sets the background color for the Stage.
// 
// The color can be given as a hex string (`'#RRGGBB'`), a CSS color string (`'rgb(r,g,b)'`), or a numeric value (`0xRRGGBB`).
// 
// An alpha channel is _not_ supported and will be ignored.
// 
// If you've set your game to be transparent then calls to setBackgroundColor are ignored.
func (self *Stage) SetBackgroundColorI(args ...interface{}) {
    self.Object.Call("setBackgroundColor", args)
}

// Destroy Destroys the Stage and removes event listeners.
func (self *Stage) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the Stage and removes event listeners.
func (self *Stage) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// AddChild Adds a child to the container.
func (self *Stage) AddChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", child)}
}

// AddChildI Adds a child to the container.
func (self *Stage) AddChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChild", args)}
}

// AddChildAt Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Stage) AddChildAt(child *DisplayObject, index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", child, index)}
}

// AddChildAtI Adds a child to the container at a specified index. If the index is out of bounds an error will be thrown
func (self *Stage) AddChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("addChildAt", args)}
}

// SwapChildren Swaps the position of 2 Display Objects within this container.
func (self *Stage) SwapChildren(child *DisplayObject, child2 *DisplayObject) {
    self.Object.Call("swapChildren", child, child2)
}

// SwapChildrenI Swaps the position of 2 Display Objects within this container.
func (self *Stage) SwapChildrenI(args ...interface{}) {
    self.Object.Call("swapChildren", args)
}

// GetChildIndex Returns the index position of a child DisplayObject instance
func (self *Stage) GetChildIndex(child *DisplayObject) int{
    return self.Object.Call("getChildIndex", child).Int()
}

// GetChildIndexI Returns the index position of a child DisplayObject instance
func (self *Stage) GetChildIndexI(args ...interface{}) int{
    return self.Object.Call("getChildIndex", args).Int()
}

// SetChildIndex Changes the position of an existing child in the display object container
func (self *Stage) SetChildIndex(child *DisplayObject, index int) {
    self.Object.Call("setChildIndex", child, index)
}

// SetChildIndexI Changes the position of an existing child in the display object container
func (self *Stage) SetChildIndexI(args ...interface{}) {
    self.Object.Call("setChildIndex", args)
}

// GetChildAt Returns the child at the specified index
func (self *Stage) GetChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", index)}
}

// GetChildAtI Returns the child at the specified index
func (self *Stage) GetChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("getChildAt", args)}
}

// RemoveChild Removes a child from the container.
func (self *Stage) RemoveChild(child *DisplayObject) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", child)}
}

// RemoveChildI Removes a child from the container.
func (self *Stage) RemoveChildI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChild", args)}
}

// RemoveChildAt Removes a child from the specified index position.
func (self *Stage) RemoveChildAt(index int) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", index)}
}

// RemoveChildAtI Removes a child from the specified index position.
func (self *Stage) RemoveChildAtI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("removeChildAt", args)}
}

// RemoveChildren Removes all children from this container that are within the begin and end indexes.
func (self *Stage) RemoveChildren(beginIndex int, endIndex int) {
    self.Object.Call("removeChildren", beginIndex, endIndex)
}

// RemoveChildrenI Removes all children from this container that are within the begin and end indexes.
func (self *Stage) RemoveChildrenI(args ...interface{}) {
    self.Object.Call("removeChildren", args)
}

// GetBounds Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *Stage) GetBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getBounds")}
}

// GetBoundsI Retrieves the bounds of the displayObjectContainer as a rectangle. The bounds calculation takes all visible children into consideration.
func (self *Stage) GetBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getBounds", args)}
}

// GetLocalBounds Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Stage) GetLocalBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds")}
}

// GetLocalBoundsI Retrieves the non-global local bounds of the displayObjectContainer as a rectangle. The calculation takes all visible children into consideration.
func (self *Stage) GetLocalBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getLocalBounds", args)}
}

// SetStageReference Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Stage) SetStageReference(stage *Stage) {
    self.Object.Call("setStageReference", stage)
}

// SetStageReferenceI Sets the containers Stage reference. This is the Stage that this object, and all of its children, is connected to.
func (self *Stage) SetStageReferenceI(args ...interface{}) {
    self.Object.Call("setStageReference", args)
}

// RemoveStageReference Removes the current stage reference from the container and all of its children.
func (self *Stage) RemoveStageReference() {
    self.Object.Call("removeStageReference")
}

// RemoveStageReferenceI Removes the current stage reference from the container and all of its children.
func (self *Stage) RemoveStageReferenceI(args ...interface{}) {
    self.Object.Call("removeStageReference", args)
}

// _renderWebGL Renders the object using the WebGL renderer
func (self *Stage) _renderWebGL(renderSession *RenderSession) {
    self.Object.Call("_renderWebGL", renderSession)
}

// _renderWebGLI Renders the object using the WebGL renderer
func (self *Stage) _renderWebGLI(args ...interface{}) {
    self.Object.Call("_renderWebGL", args)
}

// _renderCanvas Renders the object using the Canvas renderer
func (self *Stage) _renderCanvas(renderSession *RenderSession) {
    self.Object.Call("_renderCanvas", renderSession)
}

// _renderCanvasI Renders the object using the Canvas renderer
func (self *Stage) _renderCanvasI(args ...interface{}) {
    self.Object.Call("_renderCanvas", args)
}

