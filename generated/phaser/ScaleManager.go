// Package phaser Automatic generation for Phaser.ScaleManager
// generated file ScaleManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ScaleManager The ScaleManager object handles the the scaling, resizing, and alignment of the
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

// NewScaleManager Create a new ScaleManager object - this is done automatically by {@link Phaser.Game}
// 
// The `width` and `height` constructor parameters can either be a number which represents pixels or a string that represents a percentage: e.g. `800` (for 800 pixels) or `"80%"` for 80%.
func NewScaleManager(game *Game, width interface{}, height interface{}) *ScaleManager {
    return &ScaleManager{js.Global.Get("Phaser").Get("ScaleManager").New(game, width, height)}
}
// NewScaleManagerI Create a new ScaleManager object - this is done automatically by {@link Phaser.Game}
// 
// The `width` and `height` constructor parameters can either be a number which represents pixels or a string that represents a percentage: e.g. `800` (for 800 pixels) or `"80%"` for 80%.
func NewScaleManagerI(args ...interface{}) *ScaleManager {
    return &ScaleManager{js.Global.Get("Phaser").Get("ScaleManager").New(args)}
}



// Game A reference to the currently running game.
func (self *ScaleManager) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *ScaleManager) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Dom Provides access to some cross-device DOM functions.
func (self *ScaleManager) Dom() *DOM{
    return &DOM{self.Object.Get("dom")}
}

// SetDomA Provides access to some cross-device DOM functions.
func (self *ScaleManager) SetDomA(member *DOM) {
    self.Object.Set("dom", member)
}

// Grid _EXPERIMENTAL:_ A responsive grid on which you can align game objects.
func (self *ScaleManager) Grid() *FlexGrid{
    return &FlexGrid{self.Object.Get("grid")}
}

// SetGridA _EXPERIMENTAL:_ A responsive grid on which you can align game objects.
func (self *ScaleManager) SetGridA(member *FlexGrid) {
    self.Object.Set("grid", member)
}

// Width Target width (in pixels) of the Display canvas.
func (self *ScaleManager) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA Target width (in pixels) of the Display canvas.
func (self *ScaleManager) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height Target height (in pixels) of the Display canvas.
func (self *ScaleManager) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA Target height (in pixels) of the Display canvas.
func (self *ScaleManager) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// MinWidth Minimum width the canvas should be scaled to (in pixels).
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) MinWidth() int{
    return self.Object.Get("minWidth").Int()
}

// SetMinWidthA Minimum width the canvas should be scaled to (in pixels).
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) SetMinWidthA(member int) {
    self.Object.Set("minWidth", member)
}

// MaxWidth Maximum width the canvas should be scaled to (in pixels).
// If null it will scale to whatever width the browser can handle.
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) MaxWidth() int{
    return self.Object.Get("maxWidth").Int()
}

// SetMaxWidthA Maximum width the canvas should be scaled to (in pixels).
// If null it will scale to whatever width the browser can handle.
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) SetMaxWidthA(member int) {
    self.Object.Set("maxWidth", member)
}

// MinHeight Minimum height the canvas should be scaled to (in pixels).
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) MinHeight() int{
    return self.Object.Get("minHeight").Int()
}

// SetMinHeightA Minimum height the canvas should be scaled to (in pixels).
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) SetMinHeightA(member int) {
    self.Object.Set("minHeight", member)
}

// MaxHeight Maximum height the canvas should be scaled to (in pixels).
// If null it will scale to whatever height the browser can handle.
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) MaxHeight() int{
    return self.Object.Get("maxHeight").Int()
}

// SetMaxHeightA Maximum height the canvas should be scaled to (in pixels).
// If null it will scale to whatever height the browser can handle.
// Change with {@link Phaser.ScaleManager#setMinMax setMinMax}.
func (self *ScaleManager) SetMaxHeightA(member int) {
    self.Object.Set("maxHeight", member)
}

// Offset The offset coordinates of the Display canvas from the top-left of the browser window.
// The is used internally by Phaser.Pointer (for Input) and possibly other types.
func (self *ScaleManager) Offset() *Point{
    return &Point{self.Object.Get("offset")}
}

// SetOffsetA The offset coordinates of the Display canvas from the top-left of the browser window.
// The is used internally by Phaser.Pointer (for Input) and possibly other types.
func (self *ScaleManager) SetOffsetA(member *Point) {
    self.Object.Set("offset", member)
}

// ForceLandscape If true, the game should only run in a landscape orientation.
// Change with {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
func (self *ScaleManager) ForceLandscape() bool{
    return self.Object.Get("forceLandscape").Bool()
}

// SetForceLandscapeA If true, the game should only run in a landscape orientation.
// Change with {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
func (self *ScaleManager) SetForceLandscapeA(member bool) {
    self.Object.Set("forceLandscape", member)
}

// ForcePortrait If true, the game should only run in a portrait 
// Change with {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
func (self *ScaleManager) ForcePortrait() bool{
    return self.Object.Get("forcePortrait").Bool()
}

// SetForcePortraitA If true, the game should only run in a portrait 
// Change with {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
func (self *ScaleManager) SetForcePortraitA(member bool) {
    self.Object.Set("forcePortrait", member)
}

// IncorrectOrientation True if {@link Phaser.ScaleManager#forceLandscape forceLandscape} or {@link Phaser.ScaleManager#forcePortrait forcePortrait} are set and do not agree with the browser orientation.
// 
// This value is not updated immediately.
func (self *ScaleManager) IncorrectOrientation() bool{
    return self.Object.Get("incorrectOrientation").Bool()
}

// SetIncorrectOrientationA True if {@link Phaser.ScaleManager#forceLandscape forceLandscape} or {@link Phaser.ScaleManager#forcePortrait forcePortrait} are set and do not agree with the browser orientation.
// 
// This value is not updated immediately.
func (self *ScaleManager) SetIncorrectOrientationA(member bool) {
    self.Object.Set("incorrectOrientation", member)
}

// OnOrientationChange This signal is dispatched when the orientation changes _or_ the validity of the current orientation changes.
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
func (self *ScaleManager) OnOrientationChange() *Signal{
    return &Signal{self.Object.Get("onOrientationChange")}
}

// SetOnOrientationChangeA This signal is dispatched when the orientation changes _or_ the validity of the current orientation changes.
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
func (self *ScaleManager) SetOnOrientationChangeA(member *Signal) {
    self.Object.Set("onOrientationChange", member)
}

// EnterIncorrectOrientation This signal is dispatched when the browser enters an incorrect orientation, as defined by {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) EnterIncorrectOrientation() *Signal{
    return &Signal{self.Object.Get("enterIncorrectOrientation")}
}

// SetEnterIncorrectOrientationA This signal is dispatched when the browser enters an incorrect orientation, as defined by {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) SetEnterIncorrectOrientationA(member *Signal) {
    self.Object.Set("enterIncorrectOrientation", member)
}

// LeaveIncorrectOrientation This signal is dispatched when the browser leaves an incorrect orientation, as defined by {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) LeaveIncorrectOrientation() *Signal{
    return &Signal{self.Object.Get("leaveIncorrectOrientation")}
}

// SetLeaveIncorrectOrientationA This signal is dispatched when the browser leaves an incorrect orientation, as defined by {@link Phaser.ScaleManager#forceOrientation forceOrientation}.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) SetLeaveIncorrectOrientationA(member *Signal) {
    self.Object.Set("leaveIncorrectOrientation", member)
}

// FullScreenTarget If specified, this is the DOM element on which the Fullscreen API enter request will be invoked.
// The target element must have the correct CSS styling and contain the Display canvas.
// 
// The elements style will be modified (ie. the width and height might be set to 100%)
// but it will not be added to, removed from, or repositioned within the DOM.
// An attempt is made to restore relevant style changes when fullscreen mode is left.
// 
// For pre-2.2.0 behavior set `game.scale.fullScreenTarget = game.canvas`.
func (self *ScaleManager) FullScreenTarget() DOMElement{
    return WrapDOMElement(self.Object.Get("fullScreenTarget"))
}

// SetFullScreenTargetA If specified, this is the DOM element on which the Fullscreen API enter request will be invoked.
// The target element must have the correct CSS styling and contain the Display canvas.
// 
// The elements style will be modified (ie. the width and height might be set to 100%)
// but it will not be added to, removed from, or repositioned within the DOM.
// An attempt is made to restore relevant style changes when fullscreen mode is left.
// 
// For pre-2.2.0 behavior set `game.scale.fullScreenTarget = game.canvas`.
func (self *ScaleManager) SetFullScreenTargetA(member DOMElement) {
    self.Object.Set("fullScreenTarget", member)
}

// OnFullScreenInit This signal is dispatched when fullscreen mode is ready to be initialized but
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
func (self *ScaleManager) OnFullScreenInit() *Signal{
    return &Signal{self.Object.Get("onFullScreenInit")}
}

// SetOnFullScreenInitA This signal is dispatched when fullscreen mode is ready to be initialized but
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
func (self *ScaleManager) SetOnFullScreenInitA(member *Signal) {
    self.Object.Set("onFullScreenInit", member)
}

// OnFullScreenChange This signal is dispatched when the browser enters or leaves fullscreen mode, if supported.
// 
// The signal is supplied with a single argument: `scale` (the ScaleManager). Use `scale.isFullScreen` to determine
// if currently running in Fullscreen mode.
func (self *ScaleManager) OnFullScreenChange() *Signal{
    return &Signal{self.Object.Get("onFullScreenChange")}
}

// SetOnFullScreenChangeA This signal is dispatched when the browser enters or leaves fullscreen mode, if supported.
// 
// The signal is supplied with a single argument: `scale` (the ScaleManager). Use `scale.isFullScreen` to determine
// if currently running in Fullscreen mode.
func (self *ScaleManager) SetOnFullScreenChangeA(member *Signal) {
    self.Object.Set("onFullScreenChange", member)
}

// OnFullScreenError This signal is dispatched when the browser fails to enter fullscreen mode;
// or if the device does not support fullscreen mode and `startFullScreen` is invoked.
// 
// The signal is supplied with a single argument: `scale` (the ScaleManager).
func (self *ScaleManager) OnFullScreenError() *Signal{
    return &Signal{self.Object.Get("onFullScreenError")}
}

// SetOnFullScreenErrorA This signal is dispatched when the browser fails to enter fullscreen mode;
// or if the device does not support fullscreen mode and `startFullScreen` is invoked.
// 
// The signal is supplied with a single argument: `scale` (the ScaleManager).
func (self *ScaleManager) SetOnFullScreenErrorA(member *Signal) {
    self.Object.Set("onFullScreenError", member)
}

// ScreenOrientation The _last known_ orientation of the screen, as defined in the Window Screen Web API.
// See {@link Phaser.DOM.getScreenOrientation} for possible values.
func (self *ScaleManager) ScreenOrientation() string{
    return self.Object.Get("screenOrientation").String()
}

// SetScreenOrientationA The _last known_ orientation of the screen, as defined in the Window Screen Web API.
// See {@link Phaser.DOM.getScreenOrientation} for possible values.
func (self *ScaleManager) SetScreenOrientationA(member string) {
    self.Object.Set("screenOrientation", member)
}

// ScaleFactor The _current_ scale factor based on the game dimensions vs. the scaled dimensions.
func (self *ScaleManager) ScaleFactor() *Point{
    return &Point{self.Object.Get("scaleFactor")}
}

// SetScaleFactorA The _current_ scale factor based on the game dimensions vs. the scaled dimensions.
func (self *ScaleManager) SetScaleFactorA(member *Point) {
    self.Object.Set("scaleFactor", member)
}

// ScaleFactorInversed The _current_ inversed scale factor. The displayed dimensions divided by the game dimensions.
func (self *ScaleManager) ScaleFactorInversed() *Point{
    return &Point{self.Object.Get("scaleFactorInversed")}
}

// SetScaleFactorInversedA The _current_ inversed scale factor. The displayed dimensions divided by the game dimensions.
func (self *ScaleManager) SetScaleFactorInversedA(member *Point) {
    self.Object.Set("scaleFactorInversed", member)
}

// Margin The Display canvas is aligned by adjusting the margins; the last margins are stored here.
func (self *ScaleManager) Margin() interface{}{
    return self.Object.Get("margin")
}

// SetMarginA The Display canvas is aligned by adjusting the margins; the last margins are stored here.
func (self *ScaleManager) SetMarginA(member interface{}) {
    self.Object.Set("margin", member)
}

// Bounds The bounds of the scaled game. The x/y will match the offset of the canvas element and the width/height the scaled width and height.
func (self *ScaleManager) Bounds() *Rectangle{
    return &Rectangle{self.Object.Get("bounds")}
}

// SetBoundsA The bounds of the scaled game. The x/y will match the offset of the canvas element and the width/height the scaled width and height.
func (self *ScaleManager) SetBoundsA(member *Rectangle) {
    self.Object.Set("bounds", member)
}

// AspectRatio The aspect ratio of the scaled Display canvas.
func (self *ScaleManager) AspectRatio() int{
    return self.Object.Get("aspectRatio").Int()
}

// SetAspectRatioA The aspect ratio of the scaled Display canvas.
func (self *ScaleManager) SetAspectRatioA(member int) {
    self.Object.Set("aspectRatio", member)
}

// SourceAspectRatio The aspect ratio of the original game dimensions.
func (self *ScaleManager) SourceAspectRatio() int{
    return self.Object.Get("sourceAspectRatio").Int()
}

// SetSourceAspectRatioA The aspect ratio of the original game dimensions.
func (self *ScaleManager) SetSourceAspectRatioA(member int) {
    self.Object.Set("sourceAspectRatio", member)
}

// WindowConstraints The edges on which to constrain the game Display/canvas in _addition_ to the restrictions of the parent container.
// 
// The properties are strings and can be '', 'visual', 'layout', or 'layout-soft'.
// - If 'visual', the edge will be constrained to the Window / displayed screen area
// - If 'layout', the edge will be constrained to the CSS Layout bounds
// - An invalid value is treated as 'visual'
func (self *ScaleManager) WindowConstraints() interface{}{
    return self.Object.Get("windowConstraints")
}

// SetWindowConstraintsA The edges on which to constrain the game Display/canvas in _addition_ to the restrictions of the parent container.
// 
// The properties are strings and can be '', 'visual', 'layout', or 'layout-soft'.
// - If 'visual', the edge will be constrained to the Window / displayed screen area
// - If 'layout', the edge will be constrained to the CSS Layout bounds
// - An invalid value is treated as 'visual'
func (self *ScaleManager) SetWindowConstraintsA(member interface{}) {
    self.Object.Set("windowConstraints", member)
}

// Compatibility Various compatibility settings.
// A value of "(auto)" indicates the setting is configured based on device and runtime information.
// 
// A {@link Phaser.ScaleManager#refresh refresh} may need to be performed after making changes.
func (self *ScaleManager) Compatibility() interface{}{
    return self.Object.Get("compatibility")
}

// SetCompatibilityA Various compatibility settings.
// A value of "(auto)" indicates the setting is configured based on device and runtime information.
// 
// A {@link Phaser.ScaleManager#refresh refresh} may need to be performed after making changes.
func (self *ScaleManager) SetCompatibilityA(member interface{}) {
    self.Object.Set("compatibility", member)
}

// ParentIsWindow If the parent container of the Game canvas is the browser window itself (i.e. document.body),
// rather than another div, this should set to `true`.
// 
// The {@link Phaser.ScaleManager#parentNode parentNode} property is generally ignored while this is in effect.
func (self *ScaleManager) ParentIsWindow() bool{
    return self.Object.Get("parentIsWindow").Bool()
}

// SetParentIsWindowA If the parent container of the Game canvas is the browser window itself (i.e. document.body),
// rather than another div, this should set to `true`.
// 
// The {@link Phaser.ScaleManager#parentNode parentNode} property is generally ignored while this is in effect.
func (self *ScaleManager) SetParentIsWindowA(member bool) {
    self.Object.Set("parentIsWindow", member)
}

// ParentNode The _original_ DOM element for the parent of the Display canvas.
// This may be different in fullscreen - see {@link Phaser.ScaleManager#createFullScreenTarget createFullScreenTarget}.
// 
// This should only be changed after moving the Game canvas to a different DOM parent.
func (self *ScaleManager) ParentNode() DOMElement{
    return WrapDOMElement(self.Object.Get("parentNode"))
}

// SetParentNodeA The _original_ DOM element for the parent of the Display canvas.
// This may be different in fullscreen - see {@link Phaser.ScaleManager#createFullScreenTarget createFullScreenTarget}.
// 
// This should only be changed after moving the Game canvas to a different DOM parent.
func (self *ScaleManager) SetParentNodeA(member DOMElement) {
    self.Object.Set("parentNode", member)
}

// ParentScaleFactor The scale of the game in relation to its parent container.
func (self *ScaleManager) ParentScaleFactor() *Point{
    return &Point{self.Object.Get("parentScaleFactor")}
}

// SetParentScaleFactorA The scale of the game in relation to its parent container.
func (self *ScaleManager) SetParentScaleFactorA(member *Point) {
    self.Object.Set("parentScaleFactor", member)
}

// TrackParentInterval The maximum time (in ms) between dimension update checks for the Canvas's parent element (or window).
// Update checks normally happen quicker in response to other events.
func (self *ScaleManager) TrackParentInterval() int{
    return self.Object.Get("trackParentInterval").Int()
}

// SetTrackParentIntervalA The maximum time (in ms) between dimension update checks for the Canvas's parent element (or window).
// Update checks normally happen quicker in response to other events.
func (self *ScaleManager) SetTrackParentIntervalA(member int) {
    self.Object.Set("trackParentInterval", member)
}

// OnSizeChange This signal is dispatched when the size of the Display canvas changes _or_ the size of the Game changes. 
// When invoked this is done _after_ the Canvas size/position have been updated.
// 
// This signal is _only_ called when a change occurs and a reflow may be required.
// For example, if the canvas does not change sizes because of CSS settings (such as min-width)
// then this signal will _not_ be triggered.
// 
// Use this to handle responsive game layout options.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) OnSizeChange() *Signal{
    return &Signal{self.Object.Get("onSizeChange")}
}

// SetOnSizeChangeA This signal is dispatched when the size of the Display canvas changes _or_ the size of the Game changes. 
// When invoked this is done _after_ the Canvas size/position have been updated.
// 
// This signal is _only_ called when a change occurs and a reflow may be required.
// For example, if the canvas does not change sizes because of CSS settings (such as min-width)
// then this signal will _not_ be triggered.
// 
// Use this to handle responsive game layout options.
// 
// This is signaled from `preUpdate` (or `pauseUpdate`) _even when_ the game is paused.
func (self *ScaleManager) SetOnSizeChangeA(member *Signal) {
    self.Object.Set("onSizeChange", member)
}

// EXACT_FIT A scale mode that stretches content to fill all available space - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) EXACT_FIT() int{
    return self.Object.Get("EXACT_FIT").Int()
}

// SetEXACT_FITA A scale mode that stretches content to fill all available space - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SetEXACT_FITA(member int) {
    self.Object.Set("EXACT_FIT", member)
}

// NO_SCALE A scale mode that prevents any scaling - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) NO_SCALE() int{
    return self.Object.Get("NO_SCALE").Int()
}

// SetNO_SCALEA A scale mode that prevents any scaling - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SetNO_SCALEA(member int) {
    self.Object.Set("NO_SCALE", member)
}

// SHOW_ALL A scale mode that shows the entire game while maintaining proportions - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SHOW_ALL() int{
    return self.Object.Get("SHOW_ALL").Int()
}

// SetSHOW_ALLA A scale mode that shows the entire game while maintaining proportions - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SetSHOW_ALLA(member int) {
    self.Object.Set("SHOW_ALL", member)
}

// RESIZE A scale mode that causes the Game size to change - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) RESIZE() int{
    return self.Object.Get("RESIZE").Int()
}

// SetRESIZEA A scale mode that causes the Game size to change - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SetRESIZEA(member int) {
    self.Object.Set("RESIZE", member)
}

// USER_SCALE A scale mode that allows a custom scale factor - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) USER_SCALE() int{
    return self.Object.Get("USER_SCALE").Int()
}

// SetUSER_SCALEA A scale mode that allows a custom scale factor - see {@link Phaser.ScaleManager#scaleMode scaleMode}.
func (self *ScaleManager) SetUSER_SCALEA(member int) {
    self.Object.Set("USER_SCALE", member)
}

// BoundingParent The DOM element that is considered the Parent bounding element, if any.
// 
// This `null` if {@link Phaser.ScaleManager#parentIsWindow parentIsWindow} is true or if fullscreen mode is entered and {@link Phaser.ScaleManager#fullScreenTarget fullScreenTarget} is specified.
// It will also be null if there is no game canvas or if the game canvas has no parent.
func (self *ScaleManager) BoundingParent() DOMElement{
    return WrapDOMElement(self.Object.Get("boundingParent"))
}

// SetBoundingParentA The DOM element that is considered the Parent bounding element, if any.
// 
// This `null` if {@link Phaser.ScaleManager#parentIsWindow parentIsWindow} is true or if fullscreen mode is entered and {@link Phaser.ScaleManager#fullScreenTarget fullScreenTarget} is specified.
// It will also be null if there is no game canvas or if the game canvas has no parent.
func (self *ScaleManager) SetBoundingParentA(member DOMElement) {
    self.Object.Set("boundingParent", member)
}

// ScaleMode The scaling method used by the ScaleManager when not in fullscreen.
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
func (self *ScaleManager) ScaleMode() int{
    return self.Object.Get("scaleMode").Int()
}

// SetScaleModeA The scaling method used by the ScaleManager when not in fullscreen.
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
func (self *ScaleManager) SetScaleModeA(member int) {
    self.Object.Set("scaleMode", member)
}

// FullScreenScaleMode The scaling method used by the ScaleManager when in fullscreen.
// 
// See {@link Phaser.ScaleManager#scaleMode scaleMode} for the different modes allowed.
func (self *ScaleManager) FullScreenScaleMode() int{
    return self.Object.Get("fullScreenScaleMode").Int()
}

// SetFullScreenScaleModeA The scaling method used by the ScaleManager when in fullscreen.
// 
// See {@link Phaser.ScaleManager#scaleMode scaleMode} for the different modes allowed.
func (self *ScaleManager) SetFullScreenScaleModeA(member int) {
    self.Object.Set("fullScreenScaleMode", member)
}

// CurrentScaleMode Returns the current scale mode - for normal or fullscreen operation.
// 
// See {@link Phaser.ScaleManager#scaleMode scaleMode} for the different modes allowed.
func (self *ScaleManager) CurrentScaleMode() int{
    return self.Object.Get("currentScaleMode").Int()
}

// SetCurrentScaleModeA Returns the current scale mode - for normal or fullscreen operation.
// 
// See {@link Phaser.ScaleManager#scaleMode scaleMode} for the different modes allowed.
func (self *ScaleManager) SetCurrentScaleModeA(member int) {
    self.Object.Set("currentScaleMode", member)
}

// PageAlignHorizontally When enabled the Display canvas will be horizontally-aligned _in the Parent container_ (or {@link Phaser.ScaleManager#parentIsWindow window}).
// 
// To align horizontally across the page the Display canvas should be added directly to page;
// or the parent container should itself be horizontally aligned.
// 
// Horizontal alignment is not applicable with the {@link Phaser.ScaleManager.RESIZE RESIZE} scaling mode.
func (self *ScaleManager) PageAlignHorizontally() bool{
    return self.Object.Get("pageAlignHorizontally").Bool()
}

// SetPageAlignHorizontallyA When enabled the Display canvas will be horizontally-aligned _in the Parent container_ (or {@link Phaser.ScaleManager#parentIsWindow window}).
// 
// To align horizontally across the page the Display canvas should be added directly to page;
// or the parent container should itself be horizontally aligned.
// 
// Horizontal alignment is not applicable with the {@link Phaser.ScaleManager.RESIZE RESIZE} scaling mode.
func (self *ScaleManager) SetPageAlignHorizontallyA(member bool) {
    self.Object.Set("pageAlignHorizontally", member)
}

// PageAlignVertically When enabled the Display canvas will be vertically-aligned _in the Parent container_ (or {@link Phaser.ScaleManager#parentIsWindow window}).
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
func (self *ScaleManager) PageAlignVertically() bool{
    return self.Object.Get("pageAlignVertically").Bool()
}

// SetPageAlignVerticallyA When enabled the Display canvas will be vertically-aligned _in the Parent container_ (or {@link Phaser.ScaleManager#parentIsWindow window}).
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
func (self *ScaleManager) SetPageAlignVerticallyA(member bool) {
    self.Object.Set("pageAlignVertically", member)
}

// IsFullScreen Returns true if the browser is in fullscreen mode, otherwise false.
func (self *ScaleManager) IsFullScreen() bool{
    return self.Object.Get("isFullScreen").Bool()
}

// SetIsFullScreenA Returns true if the browser is in fullscreen mode, otherwise false.
func (self *ScaleManager) SetIsFullScreenA(member bool) {
    self.Object.Set("isFullScreen", member)
}

// IsPortrait Returns true if the screen orientation is in portrait mode.
func (self *ScaleManager) IsPortrait() bool{
    return self.Object.Get("isPortrait").Bool()
}

// SetIsPortraitA Returns true if the screen orientation is in portrait mode.
func (self *ScaleManager) SetIsPortraitA(member bool) {
    self.Object.Set("isPortrait", member)
}

// IsLandscape Returns true if the screen orientation is in landscape mode.
func (self *ScaleManager) IsLandscape() bool{
    return self.Object.Get("isLandscape").Bool()
}

// SetIsLandscapeA Returns true if the screen orientation is in landscape mode.
func (self *ScaleManager) SetIsLandscapeA(member bool) {
    self.Object.Set("isLandscape", member)
}

// IsGamePortrait Returns true if the game dimensions are portrait (height > width).
// This is especially useful to check when using the RESIZE scale mode 
// but wanting to maintain game orientation on desktop browsers, 
// where typically the screen orientation will always be landscape regardless of the browser viewport.
func (self *ScaleManager) IsGamePortrait() bool{
    return self.Object.Get("isGamePortrait").Bool()
}

// SetIsGamePortraitA Returns true if the game dimensions are portrait (height > width).
// This is especially useful to check when using the RESIZE scale mode 
// but wanting to maintain game orientation on desktop browsers, 
// where typically the screen orientation will always be landscape regardless of the browser viewport.
func (self *ScaleManager) SetIsGamePortraitA(member bool) {
    self.Object.Set("isGamePortrait", member)
}

// IsGameLandscape Returns true if the game dimensions are landscape (width > height).
// This is especially useful to check when using the RESIZE scale mode 
// but wanting to maintain game orientation on desktop browsers, 
// where typically the screen orientation will always be landscape regardless of the browser viewport.
func (self *ScaleManager) IsGameLandscape() bool{
    return self.Object.Get("isGameLandscape").Bool()
}

// SetIsGameLandscapeA Returns true if the game dimensions are landscape (width > height).
// This is especially useful to check when using the RESIZE scale mode 
// but wanting to maintain game orientation on desktop browsers, 
// where typically the screen orientation will always be landscape regardless of the browser viewport.
func (self *ScaleManager) SetIsGameLandscapeA(member bool) {
    self.Object.Set("isGameLandscape", member)
}


// Boot Start the ScaleManager.
func (self *ScaleManager) Boot() {
    self.Object.Call("boot")
}

// BootI Start the ScaleManager.
func (self *ScaleManager) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// ParseConfig Load configuration settings.
func (self *ScaleManager) ParseConfig(config interface{}) {
    self.Object.Call("parseConfig", config)
}

// ParseConfigI Load configuration settings.
func (self *ScaleManager) ParseConfigI(args ...interface{}) {
    self.Object.Call("parseConfig", args)
}

// SetupScale Calculates and sets the game dimensions based on the given width and height.
// 
// This should _not_ be called when in fullscreen mode.
func (self *ScaleManager) SetupScale(width interface{}, height interface{}) {
    self.Object.Call("setupScale", width, height)
}

// SetupScaleI Calculates and sets the game dimensions based on the given width and height.
// 
// This should _not_ be called when in fullscreen mode.
func (self *ScaleManager) SetupScaleI(args ...interface{}) {
    self.Object.Call("setupScale", args)
}

// _gameResumed Invoked when the game is resumed.
func (self *ScaleManager) _gameResumed() {
    self.Object.Call("_gameResumed")
}

// _gameResumedI Invoked when the game is resumed.
func (self *ScaleManager) _gameResumedI(args ...interface{}) {
    self.Object.Call("_gameResumed", args)
}

// SetGameSize Set the actual Game size.
// Use this instead of directly changing `game.width` or `game.height`.
// 
// The actual physical display (Canvas element size) depends on various settings including
// - Scale mode
// - Scaling factor
// - Size of Canvas's parent element or CSS rules such as min-height/max-height;
// - The size of the Window
func (self *ScaleManager) SetGameSize(width int, height int) {
    self.Object.Call("setGameSize", width, height)
}

// SetGameSizeI Set the actual Game size.
// Use this instead of directly changing `game.width` or `game.height`.
// 
// The actual physical display (Canvas element size) depends on various settings including
// - Scale mode
// - Scaling factor
// - Size of Canvas's parent element or CSS rules such as min-height/max-height;
// - The size of the Window
func (self *ScaleManager) SetGameSizeI(args ...interface{}) {
    self.Object.Call("setGameSize", args)
}

// SetUserScale Set a User scaling factor used in the USER_SCALE scaling mode.
// 
// The target canvas size is computed by:
// 
//     canvas.width = (game.width * hScale) - hTrim
//     canvas.height = (game.height * vScale) - vTrim
// 
// This method can be used in the {@link Phaser.ScaleManager#setResizeCallback resize callback}.
func (self *ScaleManager) SetUserScale(hScale int, vScale float64) {
    self.Object.Call("setUserScale", hScale, vScale)
}

// SetUserScale1O Set a User scaling factor used in the USER_SCALE scaling mode.
// 
// The target canvas size is computed by:
// 
//     canvas.width = (game.width * hScale) - hTrim
//     canvas.height = (game.height * vScale) - vTrim
// 
// This method can be used in the {@link Phaser.ScaleManager#setResizeCallback resize callback}.
func (self *ScaleManager) SetUserScale1O(hScale int, vScale float64, hTrim int) {
    self.Object.Call("setUserScale", hScale, vScale, hTrim)
}

// SetUserScale2O Set a User scaling factor used in the USER_SCALE scaling mode.
// 
// The target canvas size is computed by:
// 
//     canvas.width = (game.width * hScale) - hTrim
//     canvas.height = (game.height * vScale) - vTrim
// 
// This method can be used in the {@link Phaser.ScaleManager#setResizeCallback resize callback}.
func (self *ScaleManager) SetUserScale2O(hScale int, vScale float64, hTrim int, vTrim int) {
    self.Object.Call("setUserScale", hScale, vScale, hTrim, vTrim)
}

// SetUserScaleI Set a User scaling factor used in the USER_SCALE scaling mode.
// 
// The target canvas size is computed by:
// 
//     canvas.width = (game.width * hScale) - hTrim
//     canvas.height = (game.height * vScale) - vTrim
// 
// This method can be used in the {@link Phaser.ScaleManager#setResizeCallback resize callback}.
func (self *ScaleManager) SetUserScaleI(args ...interface{}) {
    self.Object.Call("setUserScale", args)
}

// SetResizeCallback Sets the callback that will be invoked before sizing calculations.
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
func (self *ScaleManager) SetResizeCallback(callback interface{}, context interface{}) {
    self.Object.Call("setResizeCallback", callback, context)
}

// SetResizeCallbackI Sets the callback that will be invoked before sizing calculations.
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
    self.Object.Call("setResizeCallback", args)
}

// SignalSizeChange Signals a resize - IF the canvas or Game size differs from the last signal.
// 
// This also triggers updates on {@link Phaser.ScaleManager#grid grid} (FlexGrid) and, if in a RESIZE mode, `game.state` (StateManager).
func (self *ScaleManager) SignalSizeChange() {
    self.Object.Call("signalSizeChange")
}

// SignalSizeChangeI Signals a resize - IF the canvas or Game size differs from the last signal.
// 
// This also triggers updates on {@link Phaser.ScaleManager#grid grid} (FlexGrid) and, if in a RESIZE mode, `game.state` (StateManager).
func (self *ScaleManager) SignalSizeChangeI(args ...interface{}) {
    self.Object.Call("signalSizeChange", args)
}

// SetMinMax Set the min and max dimensions for the Display canvas.
// 
// _Note:_ The min/max dimensions are only applied in some cases
// - When the device is not in an incorrect orientation; or
// - The scale mode is EXACT_FIT when not in fullscreen
func (self *ScaleManager) SetMinMax(minWidth int, minHeight int) {
    self.Object.Call("setMinMax", minWidth, minHeight)
}

// SetMinMax1O Set the min and max dimensions for the Display canvas.
// 
// _Note:_ The min/max dimensions are only applied in some cases
// - When the device is not in an incorrect orientation; or
// - The scale mode is EXACT_FIT when not in fullscreen
func (self *ScaleManager) SetMinMax1O(minWidth int, minHeight int, maxWidth int) {
    self.Object.Call("setMinMax", minWidth, minHeight, maxWidth)
}

// SetMinMax2O Set the min and max dimensions for the Display canvas.
// 
// _Note:_ The min/max dimensions are only applied in some cases
// - When the device is not in an incorrect orientation; or
// - The scale mode is EXACT_FIT when not in fullscreen
func (self *ScaleManager) SetMinMax2O(minWidth int, minHeight int, maxWidth int, maxHeight int) {
    self.Object.Call("setMinMax", minWidth, minHeight, maxWidth, maxHeight)
}

// SetMinMaxI Set the min and max dimensions for the Display canvas.
// 
// _Note:_ The min/max dimensions are only applied in some cases
// - When the device is not in an incorrect orientation; or
// - The scale mode is EXACT_FIT when not in fullscreen
func (self *ScaleManager) SetMinMaxI(args ...interface{}) {
    self.Object.Call("setMinMax", args)
}

// PreUpdate The ScaleManager.preUpdate is called automatically by the core Game loop.
func (self *ScaleManager) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI The ScaleManager.preUpdate is called automatically by the core Game loop.
func (self *ScaleManager) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// PauseUpdate Update method while paused.
func (self *ScaleManager) PauseUpdate() {
    self.Object.Call("pauseUpdate")
}

// PauseUpdateI Update method while paused.
func (self *ScaleManager) PauseUpdateI(args ...interface{}) {
    self.Object.Call("pauseUpdate", args)
}

// UpdateDimensions Update the dimensions taking the parent scaling factor into account.
func (self *ScaleManager) UpdateDimensions(width int, height int, resize bool) {
    self.Object.Call("updateDimensions", width, height, resize)
}

// UpdateDimensionsI Update the dimensions taking the parent scaling factor into account.
func (self *ScaleManager) UpdateDimensionsI(args ...interface{}) {
    self.Object.Call("updateDimensions", args)
}

// UpdateScalingAndBounds Update relevant scaling values based on the ScaleManager dimension and game dimensions,
// which should already be set. This does not change {@link Phaser.ScaleManager#sourceAspectRatio sourceAspectRatio}.
func (self *ScaleManager) UpdateScalingAndBounds() {
    self.Object.Call("updateScalingAndBounds")
}

// UpdateScalingAndBoundsI Update relevant scaling values based on the ScaleManager dimension and game dimensions,
// which should already be set. This does not change {@link Phaser.ScaleManager#sourceAspectRatio sourceAspectRatio}.
func (self *ScaleManager) UpdateScalingAndBoundsI(args ...interface{}) {
    self.Object.Call("updateScalingAndBounds", args)
}

// ForceOrientation Force the game to run in only one orientation.
// 
// This enables generation of incorrect orientation signals and affects resizing but does not otherwise rotate or lock the orientation.
// 
// Orientation checks are performed via the Screen Orientation API, if available in browser. This means it will check your monitor
// orientation on desktop, or your device orientation on mobile, rather than comparing actual game dimensions. If you need to check the 
// viewport dimensions instead and bypass the Screen Orientation API then set: `ScaleManager.compatibility.orientationFallback = 'viewport'`
func (self *ScaleManager) ForceOrientation(forceLandscape bool) {
    self.Object.Call("forceOrientation", forceLandscape)
}

// ForceOrientation1O Force the game to run in only one orientation.
// 
// This enables generation of incorrect orientation signals and affects resizing but does not otherwise rotate or lock the orientation.
// 
// Orientation checks are performed via the Screen Orientation API, if available in browser. This means it will check your monitor
// orientation on desktop, or your device orientation on mobile, rather than comparing actual game dimensions. If you need to check the 
// viewport dimensions instead and bypass the Screen Orientation API then set: `ScaleManager.compatibility.orientationFallback = 'viewport'`
func (self *ScaleManager) ForceOrientation1O(forceLandscape bool, forcePortrait bool) {
    self.Object.Call("forceOrientation", forceLandscape, forcePortrait)
}

// ForceOrientationI Force the game to run in only one orientation.
// 
// This enables generation of incorrect orientation signals and affects resizing but does not otherwise rotate or lock the orientation.
// 
// Orientation checks are performed via the Screen Orientation API, if available in browser. This means it will check your monitor
// orientation on desktop, or your device orientation on mobile, rather than comparing actual game dimensions. If you need to check the 
// viewport dimensions instead and bypass the Screen Orientation API then set: `ScaleManager.compatibility.orientationFallback = 'viewport'`
func (self *ScaleManager) ForceOrientationI(args ...interface{}) {
    self.Object.Call("forceOrientation", args)
}

// ClassifyOrientation Classify the orientation, per `getScreenOrientation`.
func (self *ScaleManager) ClassifyOrientation(orientation string) string{
    return self.Object.Call("classifyOrientation", orientation).String()
}

// ClassifyOrientationI Classify the orientation, per `getScreenOrientation`.
func (self *ScaleManager) ClassifyOrientationI(args ...interface{}) string{
    return self.Object.Call("classifyOrientation", args).String()
}

// UpdateOrientationState Updates the current orientation and dispatches orientation change events.
func (self *ScaleManager) UpdateOrientationState() bool{
    return self.Object.Call("updateOrientationState").Bool()
}

// UpdateOrientationStateI Updates the current orientation and dispatches orientation change events.
func (self *ScaleManager) UpdateOrientationStateI(args ...interface{}) bool{
    return self.Object.Call("updateOrientationState", args).Bool()
}

// OrientationChange window.orientationchange event handler.
func (self *ScaleManager) OrientationChange(event *Event) {
    self.Object.Call("orientationChange", event)
}

// OrientationChangeI window.orientationchange event handler.
func (self *ScaleManager) OrientationChangeI(args ...interface{}) {
    self.Object.Call("orientationChange", args)
}

// WindowResize window.resize event handler.
func (self *ScaleManager) WindowResize(event *Event) {
    self.Object.Call("windowResize", event)
}

// WindowResizeI window.resize event handler.
func (self *ScaleManager) WindowResizeI(args ...interface{}) {
    self.Object.Call("windowResize", args)
}

// ScrollTop Scroll to the top - in some environments. See `compatibility.scrollTo`.
func (self *ScaleManager) ScrollTop() {
    self.Object.Call("scrollTop")
}

// ScrollTopI Scroll to the top - in some environments. See `compatibility.scrollTo`.
func (self *ScaleManager) ScrollTopI(args ...interface{}) {
    self.Object.Call("scrollTop", args)
}

// Refresh The "refresh" methods informs the ScaleManager that a layout refresh is required.
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
func (self *ScaleManager) Refresh() {
    self.Object.Call("refresh")
}

// RefreshI The "refresh" methods informs the ScaleManager that a layout refresh is required.
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
    self.Object.Call("refresh", args)
}

// UpdateLayout Updates the game / canvas position and size.
func (self *ScaleManager) UpdateLayout() {
    self.Object.Call("updateLayout")
}

// UpdateLayoutI Updates the game / canvas position and size.
func (self *ScaleManager) UpdateLayoutI(args ...interface{}) {
    self.Object.Call("updateLayout", args)
}

// GetParentBounds Returns the computed Parent size/bounds that the Display canvas is allowed/expected to fill.
// 
// If in fullscreen mode or without parent (see {@link Phaser.ScaleManager#parentIsWindow parentIsWindow}),
// this will be the bounds of the visual viewport itself.
// 
// This function takes the {@link Phaser.ScaleManager#windowConstraints windowConstraints} into consideration - if the parent is partially outside
// the viewport then this function may return a smaller than expected size.
// 
// Values are rounded to the nearest pixel.
func (self *ScaleManager) GetParentBounds() *Rectangle{
    return &Rectangle{self.Object.Call("getParentBounds")}
}

// GetParentBounds1O Returns the computed Parent size/bounds that the Display canvas is allowed/expected to fill.
// 
// If in fullscreen mode or without parent (see {@link Phaser.ScaleManager#parentIsWindow parentIsWindow}),
// this will be the bounds of the visual viewport itself.
// 
// This function takes the {@link Phaser.ScaleManager#windowConstraints windowConstraints} into consideration - if the parent is partially outside
// the viewport then this function may return a smaller than expected size.
// 
// Values are rounded to the nearest pixel.
func (self *ScaleManager) GetParentBounds1O(target *Rectangle) *Rectangle{
    return &Rectangle{self.Object.Call("getParentBounds", target)}
}

// GetParentBoundsI Returns the computed Parent size/bounds that the Display canvas is allowed/expected to fill.
// 
// If in fullscreen mode or without parent (see {@link Phaser.ScaleManager#parentIsWindow parentIsWindow}),
// this will be the bounds of the visual viewport itself.
// 
// This function takes the {@link Phaser.ScaleManager#windowConstraints windowConstraints} into consideration - if the parent is partially outside
// the viewport then this function may return a smaller than expected size.
// 
// Values are rounded to the nearest pixel.
func (self *ScaleManager) GetParentBoundsI(args ...interface{}) *Rectangle{
    return &Rectangle{self.Object.Call("getParentBounds", args)}
}

// AlignCanvas Update the canvas position/margins - for alignment within the parent container.
// 
// The canvas margins _must_ be reset/cleared prior to invoking this.
func (self *ScaleManager) AlignCanvas(horizontal bool, vertical bool) {
    self.Object.Call("alignCanvas", horizontal, vertical)
}

// AlignCanvasI Update the canvas position/margins - for alignment within the parent container.
// 
// The canvas margins _must_ be reset/cleared prior to invoking this.
func (self *ScaleManager) AlignCanvasI(args ...interface{}) {
    self.Object.Call("alignCanvas", args)
}

// ReflowGame Updates the Game state / size.
// 
// The canvas margins may always be adjusted, even if alignment is not in effect.
func (self *ScaleManager) ReflowGame() {
    self.Object.Call("reflowGame")
}

// ReflowGameI Updates the Game state / size.
// 
// The canvas margins may always be adjusted, even if alignment is not in effect.
func (self *ScaleManager) ReflowGameI(args ...interface{}) {
    self.Object.Call("reflowGame", args)
}

// ReflowCanvas Updates the Display canvas size.
// 
// The canvas margins may always be adjusted, even alignment is not in effect.
func (self *ScaleManager) ReflowCanvas() {
    self.Object.Call("reflowCanvas")
}

// ReflowCanvasI Updates the Display canvas size.
// 
// The canvas margins may always be adjusted, even alignment is not in effect.
func (self *ScaleManager) ReflowCanvasI(args ...interface{}) {
    self.Object.Call("reflowCanvas", args)
}

// ResetCanvas "Reset" the Display canvas and set the specified width/height.
func (self *ScaleManager) ResetCanvas() {
    self.Object.Call("resetCanvas")
}

// ResetCanvas1O "Reset" the Display canvas and set the specified width/height.
func (self *ScaleManager) ResetCanvas1O(cssWidth string) {
    self.Object.Call("resetCanvas", cssWidth)
}

// ResetCanvas2O "Reset" the Display canvas and set the specified width/height.
func (self *ScaleManager) ResetCanvas2O(cssWidth string, cssHeight string) {
    self.Object.Call("resetCanvas", cssWidth, cssHeight)
}

// ResetCanvasI "Reset" the Display canvas and set the specified width/height.
func (self *ScaleManager) ResetCanvasI(args ...interface{}) {
    self.Object.Call("resetCanvas", args)
}

// QueueUpdate Queues/marks a size/bounds check as needing to occur (from `preUpdate`).
func (self *ScaleManager) QueueUpdate(force bool) {
    self.Object.Call("queueUpdate", force)
}

// QueueUpdateI Queues/marks a size/bounds check as needing to occur (from `preUpdate`).
func (self *ScaleManager) QueueUpdateI(args ...interface{}) {
    self.Object.Call("queueUpdate", args)
}

// Reset Reset internal data/state.
func (self *ScaleManager) Reset() {
    self.Object.Call("reset")
}

// ResetI Reset internal data/state.
func (self *ScaleManager) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// SetMaximum Updates the width/height to that of the window.
func (self *ScaleManager) SetMaximum() {
    self.Object.Call("setMaximum")
}

// SetMaximumI Updates the width/height to that of the window.
func (self *ScaleManager) SetMaximumI(args ...interface{}) {
    self.Object.Call("setMaximum", args)
}

// SetShowAll Updates the width/height such that the game is scaled proportionally.
func (self *ScaleManager) SetShowAll(expanding bool) {
    self.Object.Call("setShowAll", expanding)
}

// SetShowAllI Updates the width/height such that the game is scaled proportionally.
func (self *ScaleManager) SetShowAllI(args ...interface{}) {
    self.Object.Call("setShowAll", args)
}

// SetExactFit Updates the width/height such that the game is stretched to the available size.
// Honors {@link Phaser.ScaleManager#maxWidth maxWidth} and {@link Phaser.ScaleManager#maxHeight maxHeight} when _not_ in fullscreen.
func (self *ScaleManager) SetExactFit() {
    self.Object.Call("setExactFit")
}

// SetExactFitI Updates the width/height such that the game is stretched to the available size.
// Honors {@link Phaser.ScaleManager#maxWidth maxWidth} and {@link Phaser.ScaleManager#maxHeight maxHeight} when _not_ in fullscreen.
func (self *ScaleManager) SetExactFitI(args ...interface{}) {
    self.Object.Call("setExactFit", args)
}

// CreateFullScreenTarget Creates a fullscreen target. This is called automatically as as needed when entering
// fullscreen mode and the resulting element is supplied to {@link Phaser.ScaleManager#onFullScreenInit onFullScreenInit}.
// 
// Use {@link Phaser.ScaleManager#onFullScreenInit onFullScreenInit} to customize the created object.
func (self *ScaleManager) CreateFullScreenTarget() {
    self.Object.Call("createFullScreenTarget")
}

// CreateFullScreenTargetI Creates a fullscreen target. This is called automatically as as needed when entering
// fullscreen mode and the resulting element is supplied to {@link Phaser.ScaleManager#onFullScreenInit onFullScreenInit}.
// 
// Use {@link Phaser.ScaleManager#onFullScreenInit onFullScreenInit} to customize the created object.
func (self *ScaleManager) CreateFullScreenTargetI(args ...interface{}) {
    self.Object.Call("createFullScreenTarget", args)
}

// StartFullScreen Start the browsers fullscreen mode - this _must_ be called from a user input Pointer or Mouse event.
// 
// The Fullscreen API must be supported by the browser for this to work - it is not the same as setting
// the game size to fill the browser window. See {@link Phaser.ScaleManager#compatibility compatibility.supportsFullScreen} to check if the current
// device is reported to support fullscreen mode.
// 
// The {@link Phaser.ScaleManager#fullScreenFailed fullScreenFailed} signal will be dispatched if the fullscreen change request failed or the game does not support the Fullscreen API.
func (self *ScaleManager) StartFullScreen() bool{
    return self.Object.Call("startFullScreen").Bool()
}

// StartFullScreen1O Start the browsers fullscreen mode - this _must_ be called from a user input Pointer or Mouse event.
// 
// The Fullscreen API must be supported by the browser for this to work - it is not the same as setting
// the game size to fill the browser window. See {@link Phaser.ScaleManager#compatibility compatibility.supportsFullScreen} to check if the current
// device is reported to support fullscreen mode.
// 
// The {@link Phaser.ScaleManager#fullScreenFailed fullScreenFailed} signal will be dispatched if the fullscreen change request failed or the game does not support the Fullscreen API.
func (self *ScaleManager) StartFullScreen1O(antialias bool) bool{
    return self.Object.Call("startFullScreen", antialias).Bool()
}

// StartFullScreen2O Start the browsers fullscreen mode - this _must_ be called from a user input Pointer or Mouse event.
// 
// The Fullscreen API must be supported by the browser for this to work - it is not the same as setting
// the game size to fill the browser window. See {@link Phaser.ScaleManager#compatibility compatibility.supportsFullScreen} to check if the current
// device is reported to support fullscreen mode.
// 
// The {@link Phaser.ScaleManager#fullScreenFailed fullScreenFailed} signal will be dispatched if the fullscreen change request failed or the game does not support the Fullscreen API.
func (self *ScaleManager) StartFullScreen2O(antialias bool, allowTrampoline bool) bool{
    return self.Object.Call("startFullScreen", antialias, allowTrampoline).Bool()
}

// StartFullScreenI Start the browsers fullscreen mode - this _must_ be called from a user input Pointer or Mouse event.
// 
// The Fullscreen API must be supported by the browser for this to work - it is not the same as setting
// the game size to fill the browser window. See {@link Phaser.ScaleManager#compatibility compatibility.supportsFullScreen} to check if the current
// device is reported to support fullscreen mode.
// 
// The {@link Phaser.ScaleManager#fullScreenFailed fullScreenFailed} signal will be dispatched if the fullscreen change request failed or the game does not support the Fullscreen API.
func (self *ScaleManager) StartFullScreenI(args ...interface{}) bool{
    return self.Object.Call("startFullScreen", args).Bool()
}

// StopFullScreen Stops / exits fullscreen mode, if active.
func (self *ScaleManager) StopFullScreen() bool{
    return self.Object.Call("stopFullScreen").Bool()
}

// StopFullScreenI Stops / exits fullscreen mode, if active.
func (self *ScaleManager) StopFullScreenI(args ...interface{}) bool{
    return self.Object.Call("stopFullScreen", args).Bool()
}

// CleanupCreatedTarget Cleans up the previous fullscreen target, if such was automatically created.
// This ensures the canvas is restored to its former parent, assuming the target didn't move.
func (self *ScaleManager) CleanupCreatedTarget() {
    self.Object.Call("cleanupCreatedTarget")
}

// CleanupCreatedTargetI Cleans up the previous fullscreen target, if such was automatically created.
// This ensures the canvas is restored to its former parent, assuming the target didn't move.
func (self *ScaleManager) CleanupCreatedTargetI(args ...interface{}) {
    self.Object.Call("cleanupCreatedTarget", args)
}

// PrepScreenMode Used to prepare/restore extra fullscreen mode settings.
// (This does move any elements within the DOM tree.)
func (self *ScaleManager) PrepScreenMode(enteringFullscreen bool) {
    self.Object.Call("prepScreenMode", enteringFullscreen)
}

// PrepScreenModeI Used to prepare/restore extra fullscreen mode settings.
// (This does move any elements within the DOM tree.)
func (self *ScaleManager) PrepScreenModeI(args ...interface{}) {
    self.Object.Call("prepScreenMode", args)
}

// FullScreenChange Called automatically when the browser enters of leaves fullscreen mode.
func (self *ScaleManager) FullScreenChange() {
    self.Object.Call("fullScreenChange")
}

// FullScreenChange1O Called automatically when the browser enters of leaves fullscreen mode.
func (self *ScaleManager) FullScreenChange1O(event *Event) {
    self.Object.Call("fullScreenChange", event)
}

// FullScreenChangeI Called automatically when the browser enters of leaves fullscreen mode.
func (self *ScaleManager) FullScreenChangeI(args ...interface{}) {
    self.Object.Call("fullScreenChange", args)
}

// FullScreenError Called automatically when the browser fullscreen request fails;
// or called when a fullscreen request is made on a device for which it is not supported.
func (self *ScaleManager) FullScreenError() {
    self.Object.Call("fullScreenError")
}

// FullScreenError1O Called automatically when the browser fullscreen request fails;
// or called when a fullscreen request is made on a device for which it is not supported.
func (self *ScaleManager) FullScreenError1O(event *Event) {
    self.Object.Call("fullScreenError", event)
}

// FullScreenErrorI Called automatically when the browser fullscreen request fails;
// or called when a fullscreen request is made on a device for which it is not supported.
func (self *ScaleManager) FullScreenErrorI(args ...interface{}) {
    self.Object.Call("fullScreenError", args)
}

// ScaleSprite Takes a Sprite or Image object and scales it to fit the given dimensions.
// Scaling happens proportionally without distortion to the sprites texture.
// The letterBox parameter controls if scaling will produce a letter-box effect or zoom the
// sprite until it fills the given values. Note that with letterBox set to false the scaled sprite may spill out over either
// the horizontal or vertical sides of the target dimensions. If you wish to stop this you can crop the Sprite.
func (self *ScaleManager) ScaleSprite(sprite interface{}) interface{}{
    return self.Object.Call("scaleSprite", sprite)
}

// ScaleSprite1O Takes a Sprite or Image object and scales it to fit the given dimensions.
// Scaling happens proportionally without distortion to the sprites texture.
// The letterBox parameter controls if scaling will produce a letter-box effect or zoom the
// sprite until it fills the given values. Note that with letterBox set to false the scaled sprite may spill out over either
// the horizontal or vertical sides of the target dimensions. If you wish to stop this you can crop the Sprite.
func (self *ScaleManager) ScaleSprite1O(sprite interface{}, width int) interface{}{
    return self.Object.Call("scaleSprite", sprite, width)
}

// ScaleSprite2O Takes a Sprite or Image object and scales it to fit the given dimensions.
// Scaling happens proportionally without distortion to the sprites texture.
// The letterBox parameter controls if scaling will produce a letter-box effect or zoom the
// sprite until it fills the given values. Note that with letterBox set to false the scaled sprite may spill out over either
// the horizontal or vertical sides of the target dimensions. If you wish to stop this you can crop the Sprite.
func (self *ScaleManager) ScaleSprite2O(sprite interface{}, width int, height int) interface{}{
    return self.Object.Call("scaleSprite", sprite, width, height)
}

// ScaleSprite3O Takes a Sprite or Image object and scales it to fit the given dimensions.
// Scaling happens proportionally without distortion to the sprites texture.
// The letterBox parameter controls if scaling will produce a letter-box effect or zoom the
// sprite until it fills the given values. Note that with letterBox set to false the scaled sprite may spill out over either
// the horizontal or vertical sides of the target dimensions. If you wish to stop this you can crop the Sprite.
func (self *ScaleManager) ScaleSprite3O(sprite interface{}, width int, height int, letterBox bool) interface{}{
    return self.Object.Call("scaleSprite", sprite, width, height, letterBox)
}

// ScaleSpriteI Takes a Sprite or Image object and scales it to fit the given dimensions.
// Scaling happens proportionally without distortion to the sprites texture.
// The letterBox parameter controls if scaling will produce a letter-box effect or zoom the
// sprite until it fills the given values. Note that with letterBox set to false the scaled sprite may spill out over either
// the horizontal or vertical sides of the target dimensions. If you wish to stop this you can crop the Sprite.
func (self *ScaleManager) ScaleSpriteI(args ...interface{}) interface{}{
    return self.Object.Call("scaleSprite", args)
}

// Destroy Destroys the ScaleManager and removes any event listeners.
// This should probably only be called when the game is destroyed.
func (self *ScaleManager) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys the ScaleManager and removes any event listeners.
// This should probably only be called when the game is destroyed.
func (self *ScaleManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

