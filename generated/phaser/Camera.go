// Package phaser Automatic generation for Phaser.Camera
// generated file Camera.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Camera A Camera is your view into the game world. It has a position and size and renders only those objects within its field of view.
// The game automatically creates a single Stage sized camera on boot. Move the camera around the world with Phaser.Camera.x/y
type Camera struct {
    *js.Object
}

// NewCamera A Camera is your view into the game world. It has a position and size and renders only those objects within its field of view.
// The game automatically creates a single Stage sized camera on boot. Move the camera around the world with Phaser.Camera.x/y
func NewCamera(game *Game, id int, x int, y int, width int, height int) *Camera {
    return &Camera{js.Global.Get("Phaser").Get("Camera").New(game, id, x, y, width, height)}
}
// NewCameraI A Camera is your view into the game world. It has a position and size and renders only those objects within its field of view.
// The game automatically creates a single Stage sized camera on boot. Move the camera around the world with Phaser.Camera.x/y
func NewCameraI(args ...interface{}) *Camera {
    return &Camera{js.Global.Get("Phaser").Get("Camera").New(args)}
}



// Game A reference to the currently running Game.
func (self *Camera) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Camera) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// World A reference to the game world.
func (self *Camera) World() *World{
    return &World{self.Object.Get("world")}
}

// SetWorldA A reference to the game world.
func (self *Camera) SetWorldA(member *World) {
    self.Object.Set("world", member)
}

// Id Reserved for future multiple camera set-ups.
func (self *Camera) Id() int{
    return self.Object.Get("id").Int()
}

// SetIdA Reserved for future multiple camera set-ups.
func (self *Camera) SetIdA(member int) {
    self.Object.Set("id", member)
}

// View Camera view.
// The view into the world we wish to render (by default the game dimensions).
// The x/y values are in world coordinates, not screen coordinates, the width/height is how many pixels to render.
// Sprites outside of this view are not rendered if Sprite.autoCull is set to `true`. Otherwise they are always rendered.
func (self *Camera) View() *Rectangle{
    return &Rectangle{self.Object.Get("view")}
}

// SetViewA Camera view.
// The view into the world we wish to render (by default the game dimensions).
// The x/y values are in world coordinates, not screen coordinates, the width/height is how many pixels to render.
// Sprites outside of this view are not rendered if Sprite.autoCull is set to `true`. Otherwise they are always rendered.
func (self *Camera) SetViewA(member *Rectangle) {
    self.Object.Set("view", member)
}

// Bounds The Camera is bound to this Rectangle and cannot move outside of it. By default it is enabled and set to the size of the World.
// The Rectangle can be located anywhere in the world and updated as often as you like. If you don't wish the Camera to be bound
// at all then set this to null. The values can be anything and are in World coordinates, with 0,0 being the top-left of the world. The Rectangle in which the Camera is bounded. Set to null to allow for movement anywhere.
func (self *Camera) Bounds() *Rectangle{
    return &Rectangle{self.Object.Get("bounds")}
}

// SetBoundsA The Camera is bound to this Rectangle and cannot move outside of it. By default it is enabled and set to the size of the World.
// The Rectangle can be located anywhere in the world and updated as often as you like. If you don't wish the Camera to be bound
// at all then set this to null. The values can be anything and are in World coordinates, with 0,0 being the top-left of the world. The Rectangle in which the Camera is bounded. Set to null to allow for movement anywhere.
func (self *Camera) SetBoundsA(member *Rectangle) {
    self.Object.Set("bounds", member)
}

// Deadzone Moving inside this Rectangle will not cause the camera to move.
func (self *Camera) Deadzone() *Rectangle{
    return &Rectangle{self.Object.Get("deadzone")}
}

// SetDeadzoneA Moving inside this Rectangle will not cause the camera to move.
func (self *Camera) SetDeadzoneA(member *Rectangle) {
    self.Object.Set("deadzone", member)
}

// Visible Whether this camera is visible or not.
func (self *Camera) Visible() bool{
    return self.Object.Get("visible").Bool()
}

// SetVisibleA Whether this camera is visible or not.
func (self *Camera) SetVisibleA(member bool) {
    self.Object.Set("visible", member)
}

// RoundPx If a Camera has roundPx set to `true` it will call `view.floor` as part of its update loop, keeping its boundary to integer values. Set this to `false` to disable this from happening.
func (self *Camera) RoundPx() bool{
    return self.Object.Get("roundPx").Bool()
}

// SetRoundPxA If a Camera has roundPx set to `true` it will call `view.floor` as part of its update loop, keeping its boundary to integer values. Set this to `false` to disable this from happening.
func (self *Camera) SetRoundPxA(member bool) {
    self.Object.Set("roundPx", member)
}

// AtLimit Whether this camera is flush with the World Bounds or not.
func (self *Camera) AtLimit() bool{
    return self.Object.Get("atLimit").Bool()
}

// SetAtLimitA Whether this camera is flush with the World Bounds or not.
func (self *Camera) SetAtLimitA(member bool) {
    self.Object.Set("atLimit", member)
}

// Target If the camera is tracking a Sprite, this is a reference to it, otherwise null.
func (self *Camera) Target() *Sprite{
    return &Sprite{self.Object.Get("target")}
}

// SetTargetA If the camera is tracking a Sprite, this is a reference to it, otherwise null.
func (self *Camera) SetTargetA(member *Sprite) {
    self.Object.Set("target", member)
}

// DisplayObject The display object to which all game objects are added. Set by World.boot.
func (self *Camera) DisplayObject() *DisplayObject{
    return &DisplayObject{self.Object.Get("displayObject")}
}

// SetDisplayObjectA The display object to which all game objects are added. Set by World.boot.
func (self *Camera) SetDisplayObjectA(member *DisplayObject) {
    self.Object.Set("displayObject", member)
}

// Scale The scale of the display object to which all game objects are added. Set by World.boot.
func (self *Camera) Scale() *Point{
    return &Point{self.Object.Get("scale")}
}

// SetScaleA The scale of the display object to which all game objects are added. Set by World.boot.
func (self *Camera) SetScaleA(member *Point) {
    self.Object.Set("scale", member)
}

// TotalInView The total number of Sprites with `autoCull` set to `true` that are visible by this Camera.
func (self *Camera) TotalInView() int{
    return self.Object.Get("totalInView").Int()
}

// SetTotalInViewA The total number of Sprites with `autoCull` set to `true` that are visible by this Camera.
func (self *Camera) SetTotalInViewA(member int) {
    self.Object.Set("totalInView", member)
}

// Lerp The linear interpolation value to use when following a target.
// The default values of 1 means the camera will instantly snap to the target coordinates.
// A lower value, such as 0.1 means the camera will more slowly track the target, giving
// a smooth transition. You can set the horizontal and vertical values independently, and also
// adjust this value in real-time during your game.
func (self *Camera) Lerp() *Point{
    return &Point{self.Object.Get("lerp")}
}

// SetLerpA The linear interpolation value to use when following a target.
// The default values of 1 means the camera will instantly snap to the target coordinates.
// A lower value, such as 0.1 means the camera will more slowly track the target, giving
// a smooth transition. You can set the horizontal and vertical values independently, and also
// adjust this value in real-time during your game.
func (self *Camera) SetLerpA(member *Point) {
    self.Object.Set("lerp", member)
}

// OnShakeComplete This signal is dispatched when the camera shake effect completes.
func (self *Camera) OnShakeComplete() *Signal{
    return &Signal{self.Object.Get("onShakeComplete")}
}

// SetOnShakeCompleteA This signal is dispatched when the camera shake effect completes.
func (self *Camera) SetOnShakeCompleteA(member *Signal) {
    self.Object.Set("onShakeComplete", member)
}

// OnFlashComplete This signal is dispatched when the camera flash effect completes.
func (self *Camera) OnFlashComplete() *Signal{
    return &Signal{self.Object.Get("onFlashComplete")}
}

// SetOnFlashCompleteA This signal is dispatched when the camera flash effect completes.
func (self *Camera) SetOnFlashCompleteA(member *Signal) {
    self.Object.Set("onFlashComplete", member)
}

// OnFadeComplete This signal is dispatched when the camera fade effect completes.
// When the fade effect completes you will be left with the screen black (or whatever
// color you faded to). In order to reset this call `Camera.resetFX`. This is called
// automatically when you change State.
func (self *Camera) OnFadeComplete() *Signal{
    return &Signal{self.Object.Get("onFadeComplete")}
}

// SetOnFadeCompleteA This signal is dispatched when the camera fade effect completes.
// When the fade effect completes you will be left with the screen black (or whatever
// color you faded to). In order to reset this call `Camera.resetFX`. This is called
// automatically when you change State.
func (self *Camera) SetOnFadeCompleteA(member *Signal) {
    self.Object.Set("onFadeComplete", member)
}

// Fx The Graphics object used to handle camera fx such as fade and flash.
func (self *Camera) Fx() *Graphics{
    return &Graphics{self.Object.Get("fx")}
}

// SetFxA The Graphics object used to handle camera fx such as fade and flash.
func (self *Camera) SetFxA(member *Graphics) {
    self.Object.Set("fx", member)
}

// FOLLOW_LOCKON empty description
func (self *Camera) FOLLOW_LOCKON() int{
    return self.Object.Get("FOLLOW_LOCKON").Int()
}

// SetFOLLOW_LOCKONA empty description
func (self *Camera) SetFOLLOW_LOCKONA(member int) {
    self.Object.Set("FOLLOW_LOCKON", member)
}

// FOLLOW_PLATFORMER empty description
func (self *Camera) FOLLOW_PLATFORMER() int{
    return self.Object.Get("FOLLOW_PLATFORMER").Int()
}

// SetFOLLOW_PLATFORMERA empty description
func (self *Camera) SetFOLLOW_PLATFORMERA(member int) {
    self.Object.Set("FOLLOW_PLATFORMER", member)
}

// FOLLOW_TOPDOWN empty description
func (self *Camera) FOLLOW_TOPDOWN() int{
    return self.Object.Get("FOLLOW_TOPDOWN").Int()
}

// SetFOLLOW_TOPDOWNA empty description
func (self *Camera) SetFOLLOW_TOPDOWNA(member int) {
    self.Object.Set("FOLLOW_TOPDOWN", member)
}

// FOLLOW_TOPDOWN_TIGHT empty description
func (self *Camera) FOLLOW_TOPDOWN_TIGHT() int{
    return self.Object.Get("FOLLOW_TOPDOWN_TIGHT").Int()
}

// SetFOLLOW_TOPDOWN_TIGHTA empty description
func (self *Camera) SetFOLLOW_TOPDOWN_TIGHTA(member int) {
    self.Object.Set("FOLLOW_TOPDOWN_TIGHT", member)
}

// SHAKE_BOTH empty description
func (self *Camera) SHAKE_BOTH() int{
    return self.Object.Get("SHAKE_BOTH").Int()
}

// SetSHAKE_BOTHA empty description
func (self *Camera) SetSHAKE_BOTHA(member int) {
    self.Object.Set("SHAKE_BOTH", member)
}

// SHAKE_HORIZONTAL empty description
func (self *Camera) SHAKE_HORIZONTAL() int{
    return self.Object.Get("SHAKE_HORIZONTAL").Int()
}

// SetSHAKE_HORIZONTALA empty description
func (self *Camera) SetSHAKE_HORIZONTALA(member int) {
    self.Object.Set("SHAKE_HORIZONTAL", member)
}

// SHAKE_VERTICAL empty description
func (self *Camera) SHAKE_VERTICAL() int{
    return self.Object.Get("SHAKE_VERTICAL").Int()
}

// SetSHAKE_VERTICALA empty description
func (self *Camera) SetSHAKE_VERTICALA(member int) {
    self.Object.Set("SHAKE_VERTICAL", member)
}

// ENABLE_FX empty description
func (self *Camera) ENABLE_FX() bool{
    return self.Object.Get("ENABLE_FX").Bool()
}

// SetENABLE_FXA empty description
func (self *Camera) SetENABLE_FXA(member bool) {
    self.Object.Set("ENABLE_FX", member)
}

// X The Cameras x coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras x position.
func (self *Camera) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The Cameras x coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras x position.
func (self *Camera) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The Cameras y coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras y position.
func (self *Camera) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The Cameras y coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras y position.
func (self *Camera) SetYA(member int) {
    self.Object.Set("y", member)
}

// Position The Cameras position. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras xy position using Phaser.Point object.
func (self *Camera) Position() *Point{
    return &Point{self.Object.Get("position")}
}

// SetPositionA The Cameras position. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras xy position using Phaser.Point object.
func (self *Camera) SetPositionA(member *Point) {
    self.Object.Set("position", member)
}

// Width The Cameras width. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras width.
func (self *Camera) Width() int{
    return self.Object.Get("width").Int()
}

// SetWidthA The Cameras width. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras width.
func (self *Camera) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// Height The Cameras height. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras height.
func (self *Camera) Height() int{
    return self.Object.Get("height").Int()
}

// SetHeightA The Cameras height. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras height.
func (self *Camera) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// ShakeIntensity The Cameras shake intensity. Gets or sets the cameras shake intensity.
func (self *Camera) ShakeIntensity() int{
    return self.Object.Get("shakeIntensity").Int()
}

// SetShakeIntensityA The Cameras shake intensity. Gets or sets the cameras shake intensity.
func (self *Camera) SetShakeIntensityA(member int) {
    self.Object.Set("shakeIntensity", member)
}


// Boot Called automatically by Phaser.World.
func (self *Camera) Boot() {
    self.Object.Call("boot")
}

// BootI Called automatically by Phaser.World.
func (self *Camera) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// PreUpdate Camera preUpdate. Sets the total view counter to zero.
func (self *Camera) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI Camera preUpdate. Sets the total view counter to zero.
func (self *Camera) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Follow Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) Follow(target interface{}) {
    self.Object.Call("follow", target)
}

// Follow1O Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) Follow1O(target interface{}, style int) {
    self.Object.Call("follow", target, style)
}

// Follow2O Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) Follow2O(target interface{}, style int, lerpX float64) {
    self.Object.Call("follow", target, style, lerpX)
}

// Follow3O Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) Follow3O(target interface{}, style int, lerpX float64, lerpY float64) {
    self.Object.Call("follow", target, style, lerpX, lerpY)
}

// FollowI Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) FollowI(args ...interface{}) {
    self.Object.Call("follow", args)
}

// Unfollow Sets the Camera follow target to null, stopping it from following an object if it's doing so.
func (self *Camera) Unfollow() {
    self.Object.Call("unfollow")
}

// UnfollowI Sets the Camera follow target to null, stopping it from following an object if it's doing so.
func (self *Camera) UnfollowI(args ...interface{}) {
    self.Object.Call("unfollow", args)
}

// FocusOn Move the camera focus on a display object instantly.
func (self *Camera) FocusOn(displayObject interface{}) {
    self.Object.Call("focusOn", displayObject)
}

// FocusOnI Move the camera focus on a display object instantly.
func (self *Camera) FocusOnI(args ...interface{}) {
    self.Object.Call("focusOn", args)
}

// FocusOnXY Move the camera focus on a location instantly.
func (self *Camera) FocusOnXY(x int, y int) {
    self.Object.Call("focusOnXY", x, y)
}

// FocusOnXYI Move the camera focus on a location instantly.
func (self *Camera) FocusOnXYI(args ...interface{}) {
    self.Object.Call("focusOnXY", args)
}

// Shake This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake() bool{
    return self.Object.Call("shake").Bool()
}

// Shake1O This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake1O(intensity float64) bool{
    return self.Object.Call("shake", intensity).Bool()
}

// Shake2O This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake2O(intensity float64, duration int) bool{
    return self.Object.Call("shake", intensity, duration).Bool()
}

// Shake3O This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake3O(intensity float64, duration int, force bool) bool{
    return self.Object.Call("shake", intensity, duration, force).Bool()
}

// Shake4O This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake4O(intensity float64, duration int, force bool, direction int) bool{
    return self.Object.Call("shake", intensity, duration, force, direction).Bool()
}

// Shake5O This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake5O(intensity float64, duration int, force bool, direction int, shakeBounds bool) bool{
    return self.Object.Call("shake", intensity, duration, force, direction, shakeBounds).Bool()
}

// ShakeI This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) ShakeI(args ...interface{}) bool{
    return self.Object.Call("shake", args).Bool()
}

// Flash This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) Flash() bool{
    return self.Object.Call("flash").Bool()
}

// Flash1O This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) Flash1O(color float64) bool{
    return self.Object.Call("flash", color).Bool()
}

// Flash2O This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) Flash2O(color float64, duration int) bool{
    return self.Object.Call("flash", color, duration).Bool()
}

// Flash3O This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) Flash3O(color float64, duration int, force bool) bool{
    return self.Object.Call("flash", color, duration, force).Bool()
}

// FlashI This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) FlashI(args ...interface{}) bool{
    return self.Object.Call("flash", args).Bool()
}

// Fade This creates a camera fade effect. It works by filling the game with the
// color specified, over the duration given, ending with a solid fill.
// 
// You can use this for things such as transitioning to a new scene.
// 
// The game will be left 'filled' at the end of this effect, likely obscuring
// everything. In order to reset it you can call `Camera.resetFX` and it will clear the
// fade. Or you can call `Camera.flash` with the same color as the fade, and it will
// reverse the process, bringing the game back into view again.
// 
// When the effect ends the signal Camera.onFadeComplete is dispatched.
func (self *Camera) Fade() bool{
    return self.Object.Call("fade").Bool()
}

// Fade1O This creates a camera fade effect. It works by filling the game with the
// color specified, over the duration given, ending with a solid fill.
// 
// You can use this for things such as transitioning to a new scene.
// 
// The game will be left 'filled' at the end of this effect, likely obscuring
// everything. In order to reset it you can call `Camera.resetFX` and it will clear the
// fade. Or you can call `Camera.flash` with the same color as the fade, and it will
// reverse the process, bringing the game back into view again.
// 
// When the effect ends the signal Camera.onFadeComplete is dispatched.
func (self *Camera) Fade1O(color float64) bool{
    return self.Object.Call("fade", color).Bool()
}

// Fade2O This creates a camera fade effect. It works by filling the game with the
// color specified, over the duration given, ending with a solid fill.
// 
// You can use this for things such as transitioning to a new scene.
// 
// The game will be left 'filled' at the end of this effect, likely obscuring
// everything. In order to reset it you can call `Camera.resetFX` and it will clear the
// fade. Or you can call `Camera.flash` with the same color as the fade, and it will
// reverse the process, bringing the game back into view again.
// 
// When the effect ends the signal Camera.onFadeComplete is dispatched.
func (self *Camera) Fade2O(color float64, duration int) bool{
    return self.Object.Call("fade", color, duration).Bool()
}

// Fade3O This creates a camera fade effect. It works by filling the game with the
// color specified, over the duration given, ending with a solid fill.
// 
// You can use this for things such as transitioning to a new scene.
// 
// The game will be left 'filled' at the end of this effect, likely obscuring
// everything. In order to reset it you can call `Camera.resetFX` and it will clear the
// fade. Or you can call `Camera.flash` with the same color as the fade, and it will
// reverse the process, bringing the game back into view again.
// 
// When the effect ends the signal Camera.onFadeComplete is dispatched.
func (self *Camera) Fade3O(color float64, duration int, force bool) bool{
    return self.Object.Call("fade", color, duration, force).Bool()
}

// FadeI This creates a camera fade effect. It works by filling the game with the
// color specified, over the duration given, ending with a solid fill.
// 
// You can use this for things such as transitioning to a new scene.
// 
// The game will be left 'filled' at the end of this effect, likely obscuring
// everything. In order to reset it you can call `Camera.resetFX` and it will clear the
// fade. Or you can call `Camera.flash` with the same color as the fade, and it will
// reverse the process, bringing the game back into view again.
// 
// When the effect ends the signal Camera.onFadeComplete is dispatched.
func (self *Camera) FadeI(args ...interface{}) bool{
    return self.Object.Call("fade", args).Bool()
}

// Update The camera update loop. This is called automatically by the core game loop.
func (self *Camera) Update() {
    self.Object.Call("update")
}

// UpdateI The camera update loop. This is called automatically by the core game loop.
func (self *Camera) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// UpdateFX Update the camera flash and fade effects.
func (self *Camera) UpdateFX() {
    self.Object.Call("updateFX")
}

// UpdateFXI Update the camera flash and fade effects.
func (self *Camera) UpdateFXI(args ...interface{}) {
    self.Object.Call("updateFX", args)
}

// UpdateShake Update the camera shake effect.
func (self *Camera) UpdateShake() {
    self.Object.Call("updateShake")
}

// UpdateShakeI Update the camera shake effect.
func (self *Camera) UpdateShakeI(args ...interface{}) {
    self.Object.Call("updateShake", args)
}

// UpdateTarget Internal method that handles tracking a sprite.
func (self *Camera) UpdateTarget() {
    self.Object.Call("updateTarget")
}

// UpdateTargetI Internal method that handles tracking a sprite.
func (self *Camera) UpdateTargetI(args ...interface{}) {
    self.Object.Call("updateTarget", args)
}

// SetBoundsToWorld Update the Camera bounds to match the game world.
func (self *Camera) SetBoundsToWorld() {
    self.Object.Call("setBoundsToWorld")
}

// SetBoundsToWorldI Update the Camera bounds to match the game world.
func (self *Camera) SetBoundsToWorldI(args ...interface{}) {
    self.Object.Call("setBoundsToWorld", args)
}

// CheckBounds Method called to ensure the camera doesn't venture outside of the game world.
// Called automatically by Camera.update.
func (self *Camera) CheckBounds() {
    self.Object.Call("checkBounds")
}

// CheckBoundsI Method called to ensure the camera doesn't venture outside of the game world.
// Called automatically by Camera.update.
func (self *Camera) CheckBoundsI(args ...interface{}) {
    self.Object.Call("checkBounds", args)
}

// SetPosition A helper function to set both the X and Y properties of the camera at once
// without having to use game.camera.x and game.camera.y.
func (self *Camera) SetPosition(x int, y int) {
    self.Object.Call("setPosition", x, y)
}

// SetPositionI A helper function to set both the X and Y properties of the camera at once
// without having to use game.camera.x and game.camera.y.
func (self *Camera) SetPositionI(args ...interface{}) {
    self.Object.Call("setPosition", args)
}

// SetSize Sets the size of the view rectangle given the width and height in parameters.
func (self *Camera) SetSize(width int, height int) {
    self.Object.Call("setSize", width, height)
}

// SetSizeI Sets the size of the view rectangle given the width and height in parameters.
func (self *Camera) SetSizeI(args ...interface{}) {
    self.Object.Call("setSize", args)
}

// Reset Resets the camera back to 0,0 and un-follows any object it may have been tracking.
// Also immediately resets any camera effects that may have been running such as
// shake, flash or fade.
func (self *Camera) Reset() {
    self.Object.Call("reset")
}

// ResetI Resets the camera back to 0,0 and un-follows any object it may have been tracking.
// Also immediately resets any camera effects that may have been running such as
// shake, flash or fade.
func (self *Camera) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// ResetFX Resets any active FX, such as a fade or flash and immediately clears it.
// Useful to calling after a fade in order to remove the fade from the Stage.
func (self *Camera) ResetFX() {
    self.Object.Call("resetFX")
}

// ResetFXI Resets any active FX, such as a fade or flash and immediately clears it.
// Useful to calling after a fade in order to remove the fade from the Stage.
func (self *Camera) ResetFXI(args ...interface{}) {
    self.Object.Call("resetFX", args)
}

