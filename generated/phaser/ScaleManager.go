// Automatic generation for Phaser.ScaleManager
// generated file ScaleManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The ScaleManager object handles the the scaling, resizing, and alignment of the
// Game size and the game Display canvas.
// 
// The Game size is the logical size of the game; the Display canvas has size as an HTML element.
// 
// The calculations of these are heavily influenced by the bounding Parent size which is the computed
// dimensions of the Display canvas's Parent container/element - the _effective CSS rules of the
// canvas's Parent element play an important role_ in the operation of the ScaleManager. 
// 
// The Display canvas - or Game size, depending {@link Phaser.ScaleManager#scaleMode scaleMode} - is updated to best utilize the Parent size.
// When in Fullscreen mode or with {@link Phaser.ScaleManager#parentIsWindow parentIsWindow} the Parent size is that of the visual viewport (see {@link Phaser.ScaleManager#getParentBounds getParentBounds}).
// 
// Parent and Display canvas containment guidelines:
// 
// - Style the Parent element (of the game canvas) to control the Parent size and
//   thus the Display canvas's size and layout.
// 
// - The Parent element's CSS styles should _effectively_ apply maximum (and minimum) bounding behavior.
// 
// - The Parent element should _not_ apply a padding as this is not accounted for.
//   If a padding is required apply it to the Parent's parent or apply a margin to the Parent.
//   If you need to add a border, margin or any other CSS around your game container, then use a parent element and
//   apply the CSS to this instead, otherwise you'll be constantly resizing the shape of the game container.
// 
// - The Display canvas layout CSS styles (i.e. margins, size) should not be altered/specified as
//   they may be updated by the ScaleManager.
type ScaleManager struct {
    *js.Object
}


// A reference to the currently running game.
func (self *ScaleManager) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *ScaleManager) SetGame(member Game) {
    self.Set("game", member)
}

// Provides access to some cross-device DOM functions.
func (self *ScaleManager) GetDom() DOM{
    return DOM{self.Get("dom")}
}

// Provides access to some cross-device DOM functions.
func (self *ScaleManager) SetDom(member DOM) {
    self.Set("dom", member)
}

// _EXPERIMENTAL:_ A responsive grid on which you can align game objects.
func (self *ScaleManager) GetGrid() FlexGrid{
    return FlexGrid{self.Get("grid")}
}

// _EXPERIMENTAL:_ A responsive grid on which you can align game objects.
func (self *ScaleManager) SetGrid(member FlexGrid) {
    self.Set("grid", member)
}

// Target width (in pixels) of the Display canvas.
func (self *ScaleManager) GetWidth() float64{
    return self.Get("width").Float()
}

// Target width (in pixels) of the Display canvas.
func (self *ScaleManager) SetWidth(member float64) {
    self.Set("width", member)
}

// Target height (in pixels) of the Display canvas.
func (self *ScaleManager) GetHeight() float64{
    return self.Get("height").Float()
}

// Target height (in pixels) of the Display canvas.
func (self *ScaleManager) SetHeight(member float64) {
    self.Set("height", member)
}

// Minimum width the canvas should be scaled to (in pixels).
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) GetMinWidth() float64{
    return self.Get("minWidth").Float()
}

// Minimum width the canvas should be scaled to (in pixels).
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) SetMinWidth(member float64) {
    self.Set("minWidth", member)
}

// Maximum width the canvas should be scaled to (in pixels).
// If null it will scale to whatever width the browser can handle.
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) GetMaxWidth() float64{
    return self.Get("maxWidth").Float()
}

// Maximum width the canvas should be scaled to (in pixels).
// If null it will scale to whatever width the browser can handle.
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) SetMaxWidth(member float64) {
    self.Set("maxWidth", member)
}

// Minimum height the canvas should be scaled to (in pixels).
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) GetMinHeight() float64{
    return self.Get("minHeight").Float()
}

// Minimum height the canvas should be scaled to (in pixels).
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) SetMinHeight(member float64) {
    self.Set("minHeight", member)
}

// Maximum height the canvas should be scaled to (in pixels).
// If null it will scale to whatever height the browser can handle.
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) GetMaxHeight() float64{
    return self.Get("maxHeight").Float()
}

// Maximum height the canvas should be scaled to (in pixels).
// If null it will scale to whatever height the browser can handle.
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) SetMaxHeight(member float64) {
    self.Set("maxHeight", member)
}

// The offset coordinates of the Display canvas from the top-left of the browser window.
// The is used internally by Phaser.Pointer (for Input) and possibly other types.
func (self *ScaleManager) GetOffset() Point{
    return Point{self.Get("offset")}
}

// The offset coordinates of the Display canvas from the top-left of the browser window.
// The is used internally by Phaser.Pointer (for Input) and possibly other types.
func (self *ScaleManager) SetOffset(member Point) {
    self.Set("offset", member)
}

// If true, the game should only run in a landscape orientation.
// Change with {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
func (self *ScaleManager) GetForceLandscape() bool{
    return self.Get("forceLandscape").Bool()
}

// If true, the game should only run in a landscape orientation.
// Change with {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
func (self *ScaleManager) SetForceLandscape(member bool) {
    self.Set("forceLandscape", member)
}

// If true, the game should only run in a portrait 
// Change with {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
func (self *ScaleManager) GetForcePortrait() bool{
    return self.Get("forcePortrait").Bool()
}

// If true, the game should only run in a portrait 
// Change with {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
func (self *ScaleManager) SetForcePortrait(member bool) {
    self.Set("forcePortrait", member)
}

// True if {@link Phaser.ScaleManager#forceLandscape forceLandscape} or {@link Phaser.ScaleManager#forcePortrait forcePortrait} are set and do not agree with the browser orientation.
// 
// This value is not updated immediately.
func (self *ScaleManager) GetIncorrectOrientation() bool{
    return self.Get("incorrectOrientation").Bool()
}

// True if {@link Phaser.ScaleManager#forceLandscape forceLandscape} or {@link Phaser.ScaleManager#forcePortrait forcePortrait} are set and do not agree with the browser orientation.
// 
// This value is not updated immediately.
func (self *ScaleManager) SetIncorrectOrientation(member bool) {
    self.Set("incorrectOrientation", member)
}

// This signal is dispatched when the orientation changes _or_ the validity of the current orientation changes.
// 
// The signal is supplied with the following arguments:
// - `scale` - the ScaleManager object
// - `prevOrientation`, a string - The previous orientation as per {@link Phaser.ScaleManager#screenOrientation screenOrientation}.
// - `wasIncorrect`, a boolean - True if the previous orientation was last determined to be incorrect.
// 
// Access the current orientation and validity with `scale.screenOrientation` and `scale.incorrectOrientation`.
// Thus the following tests can be done:
// 
//     // The orientation itself changed:
//     scale.screenOrientation !== prevOrientation
//     // The orientation just became incorrect:
//     scale.incorrectOrientation && !wasIncorrect
// 
// It is possible that this signal is triggered after {@link Phaser.ScaleManager#forceOrientation forceOrientation} so the orientation
// correctness changes even if the orientation itself does not change.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) GetOnOrientationChange() Signal{
    return Signal{self.Get("onOrientationChange")}
}

// This signal is dispatched when the orientation changes _or_ the validity of the current orientation changes.
// 
// The signal is supplied with the following arguments:
// - `scale` - the ScaleManager object
// - `prevOrientation`, a string - The previous orientation as per {@link Phaser.ScaleManager#screenOrientation screenOrientation}.
// - `wasIncorrect`, a boolean - True if the previous orientation was last determined to be incorrect.
// 
// Access the current orientation and validity with `scale.screenOrientation` and `scale.incorrectOrientation`.
// Thus the following tests can be done:
// 
//     // The orientation itself changed:
//     scale.screenOrientation !== prevOrientation
//     // The orientation just became incorrect:
//     scale.incorrectOrientation && !wasIncorrect
// 
// It is possible that this signal is triggered after {@link Phaser.ScaleManager#forceOrientation forceOrientation} so the orientation
// correctness changes even if the orientation itself does not change.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) SetOnOrientationChange(member Signal) {
    self.Set("onOrientationChange", member)
}

// This signal is dispatched when the browser enters an incorrect orientation, as defined by {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) GetEnterIncorrectOrientation() Signal{
    return Signal{self.Get("enterIncorrectOrientation")}
}

// This signal is dispatched when the browser enters an incorrect orientation, as defined by {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) SetEnterIncorrectOrientation(member Signal) {
    self.Set("enterIncorrectOrientation", member)
}

// This signal is dispatched when the browser leaves an incorrect orientation, as defined by {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) GetLeaveIncorrectOrientation() Signal{
    return Signal{self.Get("leaveIncorrectOrientation")}
}

// This signal is dispatched when the browser leaves an incorrect orientation, as defined by {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) SetLeaveIncorrectOrientation(member Signal) {
    self.Set("leaveIncorrectOrientation", member)
}

// If specified, this is the DOM element on which the Fullscreen API enter request will be invoked.
// The target element must have the correct CSS styling and contain the Display canvas.
// 
// The elements style will be modified (ie. the width and height might be set to 100%)
// but it will not be added to, removed from, or repositioned within the DOM.
// An attempt is made to restore relevant style changes when fullscreen mode is left.
// 
// For pre-2.2.0 behavior set `game.scale.fullScreenTarget = game.canvas`.
func (self *ScaleManager) GetFullScreenTarget() DOMElement{
    return WrapDOMElement(self.Get("fullScreenTarget"))
}

// If specified, this is the DOM element on which the Fullscreen API enter request will be invoked.
// The target element must have the correct CSS styling and contain the Display canvas.
// 
// The elements style will be modified (ie. the width and height might be set to 100%)
// but it will not be added to, removed from, or repositioned within the DOM.
// An attempt is made to restore relevant style changes when fullscreen mode is left.
// 
// For pre-2.2.0 behavior set `game.scale.fullScreenTarget = game.canvas`.
func (self *ScaleManager) SetFullScreenTarget(member DOMElement) {
    self.Set("fullScreenTarget", member)
}

// This signal is dispatched when fullscreen mode is ready to be initialized but
// before the fullscreen request.
// 
// The signal is passed two arguments: `scale` (the ScaleManager), and an object in the form `{targetElement: DOMElement}`.
// 
// The `targetElement` is the {@link Phaser.ScaleManager#fullScreenTarget fullScreenTarget} element,
// if such is assigned, or a new element created by {@link Phaser.ScaleManager#createFullScreenTarget createFullScreenTarget}.
// 
// Custom CSS styling or resets can be applied to `targetElement` as required.
// 
// If `targetElement` is _not_ the same element as {@link Phaser.ScaleManager#fullScreenTarget fullScreenTarget}:
// - After initialization the Display canvas is moved onto the `targetElement` for
//   the duration of the fullscreen mode, and restored to it's original DOM location when fullscreen is exited.
// - The `targetElement` is moved/re-parented within the DOM and may have its CSS styles updated.
// 
// The behavior of a pre-assigned target element is covered in {@link Phaser.ScaleManager#fullScreenTarget fullScreenTarget}.
func (self *ScaleManager) GetOnFullScreenInit() Signal{
    return Signal{self.Get("onFullScreenInit")}
}

// This signal is dispatched when fullscreen mode is ready to be initialized but
// before the fullscreen request.
// 
// The signal is passed two arguments: `scale` (the ScaleManager), and an object in the form `{targetElement: DOMElement}`.
// 
// The `targetElement` is the {@link Phaser.ScaleManager#fullScreenTarget fullScreenTarget} element,
// if such is assigned, or a new element created by {@link Phaser.ScaleManager#createFullScreenTarget createFullScreenTarget}.
// 
// Custom CSS styling or resets can be applied to `targetElement` as required.
// 
// If `targetElement` is _not_ the same element as {@link Phaser.ScaleManager#fullScreenTarget fullScreenTarget}:
// - After initialization the Display canvas is moved onto the `targetElement` for
//   the duration of the fullscreen mode, and restored to it's original DOM location when fullscreen is exited.
// - The `targetElement` is moved/re-parented within the DOM and may have its CSS styles updated.
// 
// The behavior of a pre-assigned target element is covered in {@link Phaser.ScaleManager#fullScreenTarget fullScreenTarget}.
func (self *ScaleManager) SetOnFullScreenInit(member Signal) {
    self.Set("onFullScreenInit", member)
}

// This signal is dispatched when the browser enters or leaves fullscreen mode, if supported.
// 
// The signal is supplied with a single argument: `scale` (the ScaleManager). Use `scale.isFullScreen` to determine
// if currently running in Fullscreen mode.
func (self *ScaleManager) GetOnFullScreenChange() Signal{
    return Signal{self.Get("onFullScreenChange")}
}

// This signal is dispatched when the browser enters or leaves fullscreen mode, if supported.
// 
// The signal is supplied with a single argument: `scale` (the ScaleManager). Use `scale.isFullScreen` to determine
// if currently running in Fullscreen mode.
func (self *ScaleManager) SetOnFullScreenChange(member Signal) {
    self.Set("onFullScreenChange", member)
}

// This signal is dispatched when the browser fails to enter fullscreen mode;
// or if the device does not support fullscreen mode and `startFullScreen` is invoked.
// 
// The signal is supplied with a single argument: `scale` (the ScaleManager).
func (self *ScaleManager) GetOnFullScreenError() Signal{
    return Signal{self.Get("onFullScreenError")}
}

// This signal is dispatched when the browser fails to enter fullscreen mode;
// or if the device does not support fullscreen mode and `startFullScreen` is invoked.
// 
// The signal is supplied with a single argument: `scale` (the ScaleManager).
func (self *ScaleManager) SetOnFullScreenError(member Signal) {
    self.Set("onFullScreenError", member)
}

// The _last known_ orientation of the screen, as defined in the Window Screen Web API.
// See {@link Phaser.DOM.getScreenOrientation} for possible values.
func (self *ScaleManager) GetScreenOrientation() string{
    return self.Get("screenOrientation").String()
}

// The _last known_ orientation of the screen, as defined in the Window Screen Web API.
// See {@link Phaser.DOM.getScreenOrientation} for possible values.
func (self *ScaleManager) SetScreenOrientation(member string) {
    self.Set("screenOrientation", member)
}

// The _current_ scale factor based on the game dimensions vs. the scaled dimensions.
func (self *ScaleManager) GetScaleFactor() Point{
    return Point{self.Get("scaleFactor")}
}

// The _current_ scale factor based on the game dimensions vs. the scaled dimensions.
func (self *ScaleManager) SetScaleFactor(member Point) {
    self.Set("scaleFactor", member)
}

// The _current_ inversed scale factor. The displayed dimensions divided by the game dimensions.
func (self *ScaleManager) GetScaleFactorInversed() Point{
    return Point{self.Get("scaleFactorInversed")}
}

// The _current_ inversed scale factor. The displayed dimensions divided by the game dimensions.
func (self *ScaleManager) SetScaleFactorInversed(member Point) {
    self.Set("scaleFactorInversed", member)
}

// The Display canvas is aligned by adjusting the margins; the last margins are stored here.
func (self *ScaleManager) GetMargin() float64{
    return self.Get("margin").Float()
}

// The Display canvas is aligned by adjusting the margins; the last margins are stored here.
func (self *ScaleManager) SetMargin(member float64) {
    self.Set("margin", member)
}

// The bounds of the scaled game. The x/y will match the offset of the canvas element and the width/height the scaled width and height.
func (self *ScaleManager) GetBounds() Rectangle{
    return Rectangle{self.Get("bounds")}
}

// The bounds of the scaled game. The x/y will match the offset of the canvas element and the width/height the scaled width and height.
func (self *ScaleManager) SetBounds(member Rectangle) {
    self.Set("bounds", member)
}

// The aspect ratio of the scaled Display canvas.
func (self *ScaleManager) GetAspectRatio() float64{
    return self.Get("aspectRatio").Float()
}

// The aspect ratio of the scaled Display canvas.
func (self *ScaleManager) SetAspectRatio(member float64) {
    self.Set("aspectRatio", member)
}

// The aspect ratio of the original game dimensions.
func (self *ScaleManager) GetSourceAspectRatio() float64{
    return self.Get("sourceAspectRatio").Float()
}

// The aspect ratio of the original game dimensions.
func (self *ScaleManager) SetSourceAspectRatio(member float64) {
    self.Set("sourceAspectRatio", member)
}

// The edges on which to constrain the game Display/canvas in _addition_ to the restrictions of the parent container.
// 
// The properties are strings and can be '', 'visual', 'layout', or 'layout-soft'.
// - If 'visual', the edge will be constrained to the Window / displayed screen area
// - If 'layout', the edge will be constrained to the CSS Layout bounds
// - An invalid value is treated as 'visual'
func (self *ScaleManager) GetWindowConstraints() interface{}{
    return self.Get("windowConstraints")
}

// The edges on which to constrain the game Display/canvas in _addition_ to the restrictions of the parent container.
// 
// The properties are strings and can be '', 'visual', 'layout', or 'layout-soft'.
// - If 'visual', the edge will be constrained to the Window / displayed screen area
// - If 'layout', the edge will be constrained to the CSS Layout bounds
// - An invalid value is treated as 'visual'
func (self *ScaleManager) SetWindowConstraints(member interface{}) {
    self.Set("windowConstraints", member)
}

// Various compatibility settings.
// A value of "(auto)" indicates the setting is configured based on device and runtime information.
// 
// A {@link Phaser.ScaleManager#refresh refresh} may need to be performed after making changes.
func (self *ScaleManager) GetCompatibility() interface{}{
    return self.Get("compatibility")
}

// Various compatibility settings.
// A value of "(auto)" indicates the setting is configured based on device and runtime information.
// 
// A {@link Phaser.ScaleManager#refresh refresh} may need to be performed after making changes.
func (self *ScaleManager) SetCompatibility(member interface{}) {
    self.Set("compatibility", member)
}

// If the parent container of the Game canvas is the browser window itself (i.e. document.body),
// rather than another div, this should set to `true`.
// 
// The {@link Phaser.ScaleManager#parentNode parentNode} property is generally ignored while this is in effect.
func (self *ScaleManager) GetParentIsWindow() bool{
    return self.Get("parentIsWindow").Bool()
}

// If the parent container of the Game canvas is the browser window itself (i.e. document.body),
// rather than another div, this should set to `true`.
// 
// The {@link Phaser.ScaleManager#parentNode parentNode} property is generally ignored while this is in effect.
func (self *ScaleManager) SetParentIsWindow(member bool) {
    self.Set("parentIsWindow", member)
}

// The _original_ DOM element for the parent of the Display canvas.
// This may be different in fullscreen - see {@link Phaser.ScaleManager#createFullScreenTarget createFullScreenTarget}.
// 
// This should only be changed after moving the Game canvas to a different DOM parent.
func (self *ScaleManager) GetParentNode() DOMElement{
    return WrapDOMElement(self.Get("parentNode"))
}

// The _original_ DOM element for the parent of the Display canvas.
// This may be different in fullscreen - see {@link Phaser.ScaleManager#createFullScreenTarget createFullScreenTarget}.
// 
// This should only be changed after moving the Game canvas to a different DOM parent.
func (self *ScaleManager) SetParentNode(member DOMElement) {
    self.Set("parentNode", member)
}

// The scale of the game in relation to its parent container.
func (self *ScaleManager) GetParentScaleFactor() Point{
    return Point{self.Get("parentScaleFactor")}
}

// The scale of the game in relation to its parent container.
func (self *ScaleManager) SetParentScaleFactor(member Point) {
    self.Set("parentScaleFactor", member)
}

// The maximum time (in ms) between dimension update checks for the Canvas's parent element (or window).
// Update checks normally happen quicker in response to other events.
func (self *ScaleManager) GetTrackParentInterval() int{
    return self.Get("trackParentInterval").Int()
}

// The maximum time (in ms) between dimension update checks for the Canvas's parent element (or window).
// Update checks normally happen quicker in response to other events.
func (self *ScaleManager) SetTrackParentInterval(member int) {
    self.Set("trackParentInterval", member)
}

// This signal is dispatched when the size of the Display canvas changes _or_ the size of the Game changes. 
// When invoked this is done _after_ the Canvas size/position have been updated.
// 
// This signal is _only_ called when a change occurs and a reflow may be required.
// For example, if the canvas does not change sizes because of CSS settings (such as min-width)
// then this signal will _not_ be triggered.
// 
// Use this to handle responsive game layout options.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) GetOnSizeChange() Signal{
    return Signal{self.Get("onSizeChange")}
}

// This signal is dispatched when the size of the Display canvas changes _or_ the size of the Game changes. 
// When invoked this is done _after_ the Canvas size/position have been updated.
// 
// This signal is _only_ called when a change occurs and a reflow may be required.
// For example, if the canvas does not change sizes because of CSS settings (such as min-width)
// then this signal will _not_ be triggered.
// 
// Use this to handle responsive game layout options.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) SetOnSizeChange(member Signal) {
    self.Set("onSizeChange", member)
}

// A scale mode that stretches content to fill all available space - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) GetEXACT_FIT() int{
    return self.Get("EXACT_FIT").Int()
}

// A scale mode that stretches content to fill all available space - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SetEXACT_FIT(member int) {
    self.Set("EXACT_FIT", member)
}

// A scale mode that prevents any scaling - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) GetNO_SCALE() int{
    return self.Get("NO_SCALE").Int()
}

// A scale mode that prevents any scaling - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SetNO_SCALE(member int) {
    self.Set("NO_SCALE", member)
}

// A scale mode that shows the entire game while maintaining proportions - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) GetSHOW_ALL() int{
    return self.Get("SHOW_ALL").Int()
}

// A scale mode that shows the entire game while maintaining proportions - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SetSHOW_ALL(member int) {
    self.Set("SHOW_ALL", member)
}

// A scale mode that causes the Game size to change - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) GetRESIZE() int{
    return self.Get("RESIZE").Int()
}

// A scale mode that causes the Game size to change - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SetRESIZE(member int) {
    self.Set("RESIZE", member)
}

// A scale mode that allows a custom scale factor - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) GetUSER_SCALE() int{
    return self.Get("USER_SCALE").Int()
}

// A scale mode that allows a custom scale factor - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SetUSER_SCALE(member int) {
    self.Set("USER_SCALE", member)
}

// The DOM element that is considered the Parent bounding element, if any.
// 
// This `null` if {@link Phaser.ScaleManager#parentIsWindow parentIsWindow} is true or if fullscreen mode is entered and {@link Phaser.ScaleManager#fullScreenTarget fullScreenTarget} is specified.
// It will also be null if there is no game canvas or if the game canvas has no parent.
func (self *ScaleManager) GetBoundingParent() DOMElement{
    return WrapDOMElement(self.Get("boundingParent"))
}

// The DOM element that is considered the Parent bounding element, if any.
// 
// This `null` if {@link Phaser.ScaleManager#parentIsWindow parentIsWindow} is true or if fullscreen mode is entered and {@link Phaser.ScaleManager#fullScreenTarget fullScreenTarget} is specified.
// It will also be null if there is no game canvas or if the game canvas has no parent.
func (self *ScaleManager) SetBoundingParent(member DOMElement) {
    self.Set("boundingParent", member)
}

// The scaling method used by the ScaleManager when not in fullscreen.
// 
// <dl>
//   <dt>{@link Phaser.ScaleManager.NO_SCALE}</dt>
//   <dd>
//       The Game display area will not be scaled - even if it is too large for the canvas/screen.
//       This mode _ignores_ any applied scaling factor and displays the canvas at the Game size.
//   </dd>
//   <dt>{@link Phaser.ScaleManager.EXACT_FIT}</dt>
//   <dd>
//       The Game display area will be _stretched_ to fill the entire size of the canvas's parent element and/or screen.
//       Proportions are not maintained.
//   </dd>
//   <dt>{@link Phaser.ScaleManager.SHOW_ALL}</dt>
//   <dd>
//       Show the entire game display area while _maintaining_ the original aspect ratio.
//   </dd>
//   <dt>{@link Phaser.ScaleManager.RESIZE}</dt>
//   <dd>
//       The dimensions of the game display area are changed to match the size of the parent container.
//       That is, this mode _changes the Game size_ to match the display size.
//       <p>
//       Any manually set Game size (see {@link Phaser.ScaleManager#setGameSize setGameSize}) is ignored while in effect.
//   </dd>
//   <dt>{@link Phaser.ScaleManager.USER_SCALE}</dt>
//   <dd>
//       The game Display is scaled according to the user-specified scale set by {@link Phaser.ScaleManager#setUserScale setUserScale}.
//       <p>
//       This scale can be adjusted in the {@link Phaser.ScaleManager#setResizeCallback resize callback}
//       for flexible custom-sizing needs.
//   </dd>
// </dl>
func (self *ScaleManager) GetScaleMode() int{
    return self.Get("scaleMode").Int()
}

// The scaling method used by the ScaleManager when not in fullscreen.
// 
// <dl>
//   <dt>{@link Phaser.ScaleManager.NO_SCALE}</dt>
//   <dd>
//       The Game display area will not be scaled - even if it is too large for the canvas/screen.
//       This mode _ignores_ any applied scaling factor and displays the canvas at the Game size.
//   </dd>
//   <dt>{@link Phaser.ScaleManager.EXACT_FIT}</dt>
//   <dd>
//       The Game display area will be _stretched_ to fill the entire size of the canvas's parent element and/or screen.
//       Proportions are not maintained.
//   </dd>
//   <dt>{@link Phaser.ScaleManager.SHOW_ALL}</dt>
//   <dd>
//       Show the entire game display area while _maintaining_ the original aspect ratio.
//   </dd>
//   <dt>{@link Phaser.ScaleManager.RESIZE}</dt>
//   <dd>
//       The dimensions of the game display area are changed to match the size of the parent container.
//       That is, this mode _changes the Game size_ to match the display size.
//       <p>
//       Any manually set Game size (see {@link Phaser.ScaleManager#setGameSize setGameSize}) is ignored while in effect.
//   </dd>
//   <dt>{@link Phaser.ScaleManager.USER_SCALE}</dt>
//   <dd>
//       The game Display is scaled according to the user-specified scale set by {@link Phaser.ScaleManager#setUserScale setUserScale}.
//       <p>
//       This scale can be adjusted in the {@link Phaser.ScaleManager#setResizeCallback resize callback}
//       for flexible custom-sizing needs.
//   </dd>
// </dl>
func (self *ScaleManager) SetScaleMode(member int) {
    self.Set("scaleMode", member)
}

// The scaling method used by the ScaleManager when in fullscreen.
// 
// See {@link Phaser.ScaleManager#scaleMode scaleMode} for the different modes allowed.
func (self *ScaleManager) GetFullScreenScaleMode() int{
    return self.Get("fullScreenScaleMode").Int()
}

// The scaling method used by the ScaleManager when in fullscreen.
// 
// See {@link Phaser.ScaleManager#scaleMode scaleMode} for the different modes allowed.
func (self *ScaleManager) SetFullScreenScaleMode(member int) {
    self.Set("fullScreenScaleMode", member)
}

// Returns the current scale mode - for normal or fullscreen operation.
// 
// See {@link Phaser.ScaleManager#scaleMode scaleMode} for the different modes allowed.
func (self *ScaleManager) GetCurrentScaleMode() float64{
    return self.Get("currentScaleMode").Float()
}

// Returns the current scale mode - for normal or fullscreen operation.
// 
// See {@link Phaser.ScaleManager#scaleMode scaleMode} for the different modes allowed.
func (self *ScaleManager) SetCurrentScaleMode(member float64) {
    self.Set("currentScaleMode", member)
}

// When enabled the Display canvas will be horizontally-aligned _in the Parent container_ (or {@link Phaser.ScaleManager#parentIsWindow window}).
// 
// To align horizontally across the page the Display canvas should be added directly to page;
// or the parent container should itself be horizontally aligned.
// 
// Horizontal alignment is not applicable with the {@link Phaser.ScaleManager.RESIZE RESIZE} scaling mode.
func (self *ScaleManager) GetPageAlignHorizontally() bool{
    return self.Get("pageAlignHorizontally").Bool()
}

// When enabled the Display canvas will be horizontally-aligned _in the Parent container_ (or {@link Phaser.ScaleManager#parentIsWindow window}).
// 
// To align horizontally across the page the Display canvas should be added directly to page;
// or the parent container should itself be horizontally aligned.
// 
// Horizontal alignment is not applicable with the {@link Phaser.ScaleManager.RESIZE RESIZE} scaling mode.
func (self *ScaleManager) SetPageAlignHorizontally(member bool) {
    self.Set("pageAlignHorizontally", member)
}

// When enabled the Display canvas will be vertically-aligned _in the Parent container_ (or {@link Phaser.ScaleManager#parentIsWindow window}).
// 
// To align vertically the Parent element should have a _non-collapsible_ height, such that it will maintain
// a height _larger_ than the height of the contained Game canvas - the game canvas will then be scaled vertically
// _within_ the remaining available height dictated by the Parent element.
// 
// One way to prevent the parent from collapsing is to add an absolute "min-height" CSS property to the parent element.
// If specifying a relative "min-height/height" or adjusting margins, the Parent height must still be non-collapsible (see note).
// 
// _Note_: In version 2.2 the minimum document height is _not_ automatically set to the viewport/window height.
// To automatically update the minimum document height set {@link Phaser.ScaleManager#compatibility compatibility.forceMinimumDocumentHeight} to true.
// 
// Vertical alignment is not applicable with the {@link Phaser.ScaleManager.RESIZE RESIZE} scaling mode.
func (self *ScaleManager) GetPageAlignVertically() bool{
    return self.Get("pageAlignVertically").Bool()
}

// When enabled the Display canvas will be vertically-aligned _in the Parent container_ (or {@link Phaser.ScaleManager#parentIsWindow window}).
// 
// To align vertically the Parent element should have a _non-collapsible_ height, such that it will maintain
// a height _larger_ than the height of the contained Game canvas - the game canvas will then be scaled vertically
// _within_ the remaining available height dictated by the Parent element.
// 
// One way to prevent the parent from collapsing is to add an absolute "min-height" CSS property to the parent element.
// If specifying a relative "min-height/height" or adjusting margins, the Parent height must still be non-collapsible (see note).
// 
// _Note_: In version 2.2 the minimum document height is _not_ automatically set to the viewport/window height.
// To automatically update the minimum document height set {@link Phaser.ScaleManager#compatibility compatibility.forceMinimumDocumentHeight} to true.
// 
// Vertical alignment is not applicable with the {@link Phaser.ScaleManager.RESIZE RESIZE} scaling mode.
func (self *ScaleManager) SetPageAlignVertically(member bool) {
    self.Set("pageAlignVertically", member)
}

// Returns true if the browser is in fullscreen mode, otherwise false.
func (self *ScaleManager) GetIsFullScreen() bool{
    return self.Get("isFullScreen").Bool()
}

// Returns true if the browser is in fullscreen mode, otherwise false.
func (self *ScaleManager) SetIsFullScreen(member bool) {
    self.Set("isFullScreen", member)
}

// Returns true if the screen orientation is in portrait mode.
func (self *ScaleManager) GetIsPortrait() bool{
    return self.Get("isPortrait").Bool()
}

// Returns true if the screen orientation is in portrait mode.
func (self *ScaleManager) SetIsPortrait(member bool) {
    self.Set("isPortrait", member)
}

// Returns true if the screen orientation is in landscape mode.
func (self *ScaleManager) GetIsLandscape() bool{
    return self.Get("isLandscape").Bool()
}

// Returns true if the screen orientation is in landscape mode.
func (self *ScaleManager) SetIsLandscape(member bool) {
    self.Set("isLandscape", member)
}

// Returns true if the game dimensions are portrait (height > width).
// This is especially useful to check when using the RESIZE scale mode 
// but wanting to maintain game orientation on desktop browsers, 
// where typically the screen orientation will always be landscape regardless of the browser viewport.
func (self *ScaleManager) GetIsGamePortrait() bool{
    return self.Get("isGamePortrait").Bool()
}

// Returns true if the game dimensions are portrait (height > width).
// This is especially useful to check when using the RESIZE scale mode 
// but wanting to maintain game orientation on desktop browsers, 
// where typically the screen orientation will always be landscape regardless of the browser viewport.
func (self *ScaleManager) SetIsGamePortrait(member bool) {
    self.Set("isGamePortrait", member)
}

// Returns true if the game dimensions are landscape (width > height).
// This is especially useful to check when using the RESIZE scale mode 
// but wanting to maintain game orientation on desktop browsers, 
// where typically the screen orientation will always be landscape regardless of the browser viewport.
func (self *ScaleManager) GetIsGameLandscape() bool{
    return self.Get("isGameLandscape").Bool()
}

// Returns true if the game dimensions are landscape (width > height).
// This is especially useful to check when using the RESIZE scale mode 
// but wanting to maintain game orientation on desktop browsers, 
// where typically the screen orientation will always be landscape regardless of the browser viewport.
func (self *ScaleManager) SetIsGameLandscape(member bool) {
    self.Set("isGameLandscape", member)
}



// Start the ScaleManager.
func (self *ScaleManager) BootI(args ...interface{}) {
    self.Call("boot", args)
}

// Load configuration settings.
func (self *ScaleManager) ParseConfigI(args ...interface{}) {
    self.Call("parseConfig", args)
}

// Calculates and sets the game dimensions based on the given width and height.
// 
// This should _not_ be called when in fullscreen mode.
func (self *ScaleManager) SetupScaleI(args ...interface{}) {
    self.Call("setupScale", args)
}

// Invoked when the game is resumed.
func (self *ScaleManager) _gameResumedI(args ...interface{}) {
    self.Call("_gameResumed", args)
}

// Set the actual Game size.
// Use this instead of directly changing `game.width` or `game.height`.
// 
// The actual physical display (Canvas element size) depends on various settings including
// - Scale mode
// - Scaling factor
// - Size of Canvas's parent element or CSS rules such as min-height/max-height;
// - The size of the Window
func (self *ScaleManager) SetGameSizeI(args ...interface{}) {
    self.Call("setGameSize", args)
}

// Set a User scaling factor used in the USER_SCALE scaling mode.
// 
// The target canvas size is computed by:
// 
//     canvas.width = (game.width * hScale) - hTrim
//     canvas.height = (game.height * vScale) - vTrim
// 
// This method can be used in the {@link Phaser.ScaleManager#setResizeCallback resize callback}.
func (self *ScaleManager) SetUserScaleI(args ...interface{}) {
    self.Call("setUserScale", args)
}

// Sets the callback that will be invoked before sizing calculations.
// 
// This is the appropriate place to call {@link Phaser.ScaleManager#setUserScale setUserScale} if needing custom dynamic scaling.
// 
// The callback is supplied with two arguments `scale` and `parentBounds` where `scale` is the ScaleManager
// and `parentBounds`, a Phaser.Rectangle, is the size of the Parent element.
// 
// This callback
// - May be invoked even though the parent container or canvas sizes have not changed
// - Unlike {@link Phaser.ScaleManager#onSizeChange onSizeChange}, it runs _before_ the canvas is guaranteed to be updated
// - Will be invoked from `preUpdate`, _even when_ the game is paused    
// 
// See {@link Phaser.ScaleManager#onSizeChange onSizeChange} for a better way of reacting to layout updates.
func (self *ScaleManager) SetResizeCallbackI(args ...interface{}) {
    self.Call("setResizeCallback", args)
}

// Signals a resize - IF the canvas or Game size differs from the last signal.
// 
// This also triggers updates on {@link Phaser.ScaleManager#grid grid} (FlexGrid) and, if in a RESIZE mode, `game.state` (StateManager).
func (self *ScaleManager) SignalSizeChangeI(args ...interface{}) {
    self.Call("signalSizeChange", args)
}

// Set the min and max dimensions for the Display canvas.
// 
// _Note:_ The min/max dimensions are only applied in some cases
// - When the device is not in an incorrect orientation; or
// - The scale mode is EXACT_FIT when not in fullscreen
func (self *ScaleManager) SetMinMaxI(args ...interface{}) {
    self.Call("setMinMax", args)
}

// The ScaleManager.preUpdate is called automatically by the core Game loop.
func (self *ScaleManager) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Update method while paused.
func (self *ScaleManager) PauseUpdateI(args ...interface{}) {
    self.Call("pauseUpdate", args)
}

// Update the dimensions taking the parent scaling factor into account.
func (self *ScaleManager) UpdateDimensionsI(args ...interface{}) {
    self.Call("updateDimensions", args)
}

// Update relevant scaling values based on the ScaleManager dimension and game dimensions,
// which should already be set. This does not change {@link Phaser.ScaleManager#sourceAspectRatio sourceAspectRatio}.
func (self *ScaleManager) UpdateScalingAndBoundsI(args ...interface{}) {
    self.Call("updateScalingAndBounds", args)
}

// Force the game to run in only one orientation.
// 
// This enables generation of incorrect orientation signals and affects resizing but does not otherwise rotate or lock the orientation.
// 
// Orientation checks are performed via the Screen Orientation API, if available in browser. This means it will check your monitor
// orientation on desktop, or your device orientation on mobile, rather than comparing actual game dimensions. If you need to check the 
// viewport dimensions instead and bypass the Screen Orientation API then set: `ScaleManager.compatibility.orientationFallback = 'viewport'`
func (self *ScaleManager) ForceOrientationI(args ...interface{}) {
    self.Call("forceOrientation", args)
}

// Classify the orientation, per `getScreenOrientation`.
func (self *ScaleManager) ClassifyOrientationI(args ...interface{}) string{
    return self.Call("classifyOrientation", args).String()
}

// Updates the current orientation and dispatches orientation change events.
func (self *ScaleManager) UpdateOrientationStateI(args ...interface{}) bool{
    return self.Call("updateOrientationState", args).Bool()
}

// window.orientationchange event handler.
func (self *ScaleManager) OrientationChangeI(args ...interface{}) {
    self.Call("orientationChange", args)
}

// window.resize event handler.
func (self *ScaleManager) WindowResizeI(args ...interface{}) {
    self.Call("windowResize", args)
}

// Scroll to the top - in some environments. See `compatibility.scrollTo`.
func (self *ScaleManager) ScrollTopI(args ...interface{}) {
    self.Call("scrollTop", args)
}

// The "refresh" methods informs the ScaleManager that a layout refresh is required.
// 
// The ScaleManager automatically queues a layout refresh (eg. updates the Game size or Display canvas layout)
// when the browser is resized, the orientation changes, or when there is a detected change
// of the Parent size. Refreshing is also done automatically when public properties,
// such as {@link Phaser.ScaleManager#scaleMode scaleMode}, are updated or state-changing methods are invoked.
// 
// The "refresh" method _may_ need to be used in a few (rare) situtations when
// 
// - a device change event is not correctly detected; or
// - the Parent size changes (and an immediate reflow is desired); or
// - the ScaleManager state is updated by non-standard means; or
// - certain {@link Phaser.ScaleManager#compatibility compatibility} properties are manually changed.
// 
// The queued layout refresh is not immediate but will run promptly in an upcoming `preRender`.
func (self *ScaleManager) RefreshI(args ...interface{}) {
    self.Call("refresh", args)
}

// Updates the game / canvas position and size.
func (self *ScaleManager) UpdateLayoutI(args ...interface{}) {
    self.Call("updateLayout", args)
}

// Returns the computed Parent size/bounds that the Display canvas is allowed/expected to fill.
// 
// If in fullscreen mode or without parent (see {@link Phaser.ScaleManager#parentIsWindow parentIsWindow}),
// this will be the bounds of the visual viewport itself.
// 
// This function takes the {@link Phaser.ScaleManager#windowConstraints windowConstraints} into consideration - if the parent is partially outside
// the viewport then this function may return a smaller than expected size.
// 
// Values are rounded to the nearest pixel.
func (self *ScaleManager) GetParentBoundsI(args ...interface{}) Rectangle{
    return Rectangle{self.Call("getParentBounds", args)}
}

// Update the canvas position/margins - for alignment within the parent container.
// 
// The canvas margins _must_ be reset/cleared prior to invoking this.
func (self *ScaleManager) AlignCanvasI(args ...interface{}) {
    self.Call("alignCanvas", args)
}

// Updates the Game state / size.
// 
// The canvas margins may always be adjusted, even if alignment is not in effect.
func (self *ScaleManager) ReflowGameI(args ...interface{}) {
    self.Call("reflowGame", args)
}

// Updates the Display canvas size.
// 
// The canvas margins may always be adjusted, even alignment is not in effect.
func (self *ScaleManager) ReflowCanvasI(args ...interface{}) {
    self.Call("reflowCanvas", args)
}

// "Reset" the Display canvas and set the specified width/height.
func (self *ScaleManager) ResetCanvasI(args ...interface{}) {
    self.Call("resetCanvas", args)
}

// Queues/marks a size/bounds check as needing to occur (from `preUpdate`).
func (self *ScaleManager) QueueUpdateI(args ...interface{}) {
    self.Call("queueUpdate", args)
}

// Reset internal data/state.
func (self *ScaleManager) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Updates the width/height to that of the window.
func (self *ScaleManager) SetMaximumI(args ...interface{}) {
    self.Call("setMaximum", args)
}

// Updates the width/height such that the game is scaled proportionally.
func (self *ScaleManager) SetShowAllI(args ...interface{}) {
    self.Call("setShowAll", args)
}

// Updates the width/height such that the game is stretched to the available size.
// Honors {@link Phaser.ScaleManager#maxWidth maxWidth} and {@link Phaser.ScaleManager#maxHeight maxHeight} when _not_ in fullscreen.
func (self *ScaleManager) SetExactFitI(args ...interface{}) {
    self.Call("setExactFit", args)
}

// Creates a fullscreen target. This is called automatically as as needed when entering
// fullscreen mode and the resulting element is supplied to {@link Phaser.ScaleManager#onFullScreenInit onFullScreenInit}.
// 
// Use {@link Phaser.ScaleManager#onFullScreenInit onFullScreenInit} to customize the created object.
func (self *ScaleManager) CreateFullScreenTargetI(args ...interface{}) {
    self.Call("createFullScreenTarget", args)
}

// Start the browsers fullscreen mode - this _must_ be called from a user input Pointer or Mouse event.
// 
// The Fullscreen API must be supported by the browser for this to work - it is not the same as setting
// the game size to fill the browser window. See {@link Phaser.ScaleManager#compatibility compatibility.supportsFullScreen} to check if the current
// device is reported to support fullscreen mode.
// 
// The {@link Phaser.ScaleManager#fullScreenFailed fullScreenFailed} signal will be dispatched if the fullscreen change request failed or the game does not support the Fullscreen API.
func (self *ScaleManager) StartFullScreenI(args ...interface{}) bool{
    return self.Call("startFullScreen", args).Bool()
}

// Stops / exits fullscreen mode, if active.
func (self *ScaleManager) StopFullScreenI(args ...interface{}) bool{
    return self.Call("stopFullScreen", args).Bool()
}

// Cleans up the previous fullscreen target, if such was automatically created.
// This ensures the canvas is restored to its former parent, assuming the target didn't move.
func (self *ScaleManager) CleanupCreatedTargetI(args ...interface{}) {
    self.Call("cleanupCreatedTarget", args)
}

// Used to prepare/restore extra fullscreen mode settings.
// (This does move any elements within the DOM tree.)
func (self *ScaleManager) PrepScreenModeI(args ...interface{}) {
    self.Call("prepScreenMode", args)
}

// Called automatically when the browser enters of leaves fullscreen mode.
func (self *ScaleManager) FullScreenChangeI(args ...interface{}) {
    self.Call("fullScreenChange", args)
}

// Called automatically when the browser fullscreen request fails;
// or called when a fullscreen request is made on a device for which it is not supported.
func (self *ScaleManager) FullScreenErrorI(args ...interface{}) {
    self.Call("fullScreenError", args)
}

// Takes a Sprite or Image object and scales it to fit the given dimensions.
// Scaling happens proportionally without distortion to the sprites texture.
// The letterBox parameter controls if scaling will produce a letter-box effect or zoom the
// sprite until it fills the given values. Note that with letterBox set to false the scaled sprite may spill out over either
// the horizontal or vertical sides of the target dimensions. If you wish to stop this you can crop the Sprite.
func (self *ScaleManager) ScaleSpriteI(args ...interface{}) interface{}{
    return self.Call("scaleSprite", args)
}

// Destroys the ScaleManager and removes any event listeners.
// This should probably only be called when the game is destroyed.
func (self *ScaleManager) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
