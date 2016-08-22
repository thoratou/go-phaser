// Automatic generation for Phaser.Camera
// generated file Camera.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A Camera is your view into the game world. It has a position and size and renders only those objects within its field of view.
// The game automatically creates a single Stage sized camera on boot. Move the camera around the world with Phaser.Camera.x/y
type Camera struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *Camera) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Camera) SetGame(member *Game) {
    self.Set("game", member)
}

// A reference to the game world.
func (self *Camera) GetWorld() *World{
    return &World{self.Get("world")}
}

// A reference to the game world.
func (self *Camera) SetWorld(member *World) {
    self.Set("world", member)
}

// Reserved for future multiple camera set-ups.
func (self *Camera) GetId() int{
    return self.Get("id").Int()
}

// Reserved for future multiple camera set-ups.
func (self *Camera) SetId(member int) {
    self.Set("id", member)
}

// Camera view.
// The view into the world we wish to render (by default the game dimensions).
// The x/y values are in world coordinates, not screen coordinates, the width/height is how many pixels to render.
// Sprites outside of this view are not rendered if Sprite.autoCull is set to `true`. Otherwise they are always rendered.
func (self *Camera) GetView() *Rectangle{
    return &Rectangle{self.Get("view")}
}

// Camera view.
// The view into the world we wish to render (by default the game dimensions).
// The x/y values are in world coordinates, not screen coordinates, the width/height is how many pixels to render.
// Sprites outside of this view are not rendered if Sprite.autoCull is set to `true`. Otherwise they are always rendered.
func (self *Camera) SetView(member *Rectangle) {
    self.Set("view", member)
}

// The Camera is bound to this Rectangle and cannot move outside of it. By default it is enabled and set to the size of the World.
// The Rectangle can be located anywhere in the world and updated as often as you like. If you don't wish the Camera to be bound
// at all then set this to null. The values can be anything and are in World coordinates, with 0,0 being the top-left of the world. The Rectangle in which the Camera is bounded. Set to null to allow for movement anywhere.
func (self *Camera) GetBounds() *Rectangle{
    return &Rectangle{self.Get("bounds")}
}

// The Camera is bound to this Rectangle and cannot move outside of it. By default it is enabled and set to the size of the World.
// The Rectangle can be located anywhere in the world and updated as often as you like. If you don't wish the Camera to be bound
// at all then set this to null. The values can be anything and are in World coordinates, with 0,0 being the top-left of the world. The Rectangle in which the Camera is bounded. Set to null to allow for movement anywhere.
func (self *Camera) SetBounds(member *Rectangle) {
    self.Set("bounds", member)
}

// Moving inside this Rectangle will not cause the camera to move.
func (self *Camera) GetDeadzone() *Rectangle{
    return &Rectangle{self.Get("deadzone")}
}

// Moving inside this Rectangle will not cause the camera to move.
func (self *Camera) SetDeadzone(member *Rectangle) {
    self.Set("deadzone", member)
}

// Whether this camera is visible or not.
func (self *Camera) GetVisible() bool{
    return self.Get("visible").Bool()
}

// Whether this camera is visible or not.
func (self *Camera) SetVisible(member bool) {
    self.Set("visible", member)
}

// If a Camera has roundPx set to `true` it will call `view.floor` as part of its update loop, keeping its boundary to integer values. Set this to `false` to disable this from happening.
func (self *Camera) GetRoundPx() bool{
    return self.Get("roundPx").Bool()
}

// If a Camera has roundPx set to `true` it will call `view.floor` as part of its update loop, keeping its boundary to integer values. Set this to `false` to disable this from happening.
func (self *Camera) SetRoundPx(member bool) {
    self.Set("roundPx", member)
}

// Whether this camera is flush with the World Bounds or not.
func (self *Camera) GetAtLimit() bool{
    return self.Get("atLimit").Bool()
}

// Whether this camera is flush with the World Bounds or not.
func (self *Camera) SetAtLimit(member bool) {
    self.Set("atLimit", member)
}

// If the camera is tracking a Sprite, this is a reference to it, otherwise null.
func (self *Camera) GetTarget() *Sprite{
    return &Sprite{self.Get("target")}
}

// If the camera is tracking a Sprite, this is a reference to it, otherwise null.
func (self *Camera) SetTarget(member *Sprite) {
    self.Set("target", member)
}

// The display object to which all game objects are added. Set by World.boot.
func (self *Camera) GetDisplayObject() *DisplayObject{
    return &DisplayObject{self.Get("displayObject")}
}

// The display object to which all game objects are added. Set by World.boot.
func (self *Camera) SetDisplayObject(member *DisplayObject) {
    self.Set("displayObject", member)
}

// The scale of the display object to which all game objects are added. Set by World.boot.
func (self *Camera) GetScale() *Point{
    return &Point{self.Get("scale")}
}

// The scale of the display object to which all game objects are added. Set by World.boot.
func (self *Camera) SetScale(member *Point) {
    self.Set("scale", member)
}

// The total number of Sprites with `autoCull` set to `true` that are visible by this Camera.
func (self *Camera) GetTotalInView() int{
    return self.Get("totalInView").Int()
}

// The total number of Sprites with `autoCull` set to `true` that are visible by this Camera.
func (self *Camera) SetTotalInView(member int) {
    self.Set("totalInView", member)
}

// The linear interpolation value to use when following a target.
// The default values of 1 means the camera will instantly snap to the target coordinates.
// A lower value, such as 0.1 means the camera will more slowly track the target, giving
// a smooth transition. You can set the horizontal and vertical values independently, and also
// adjust this value in real-time during your game.
func (self *Camera) GetLerp() *Point{
    return &Point{self.Get("lerp")}
}

// The linear interpolation value to use when following a target.
// The default values of 1 means the camera will instantly snap to the target coordinates.
// A lower value, such as 0.1 means the camera will more slowly track the target, giving
// a smooth transition. You can set the horizontal and vertical values independently, and also
// adjust this value in real-time during your game.
func (self *Camera) SetLerp(member *Point) {
    self.Set("lerp", member)
}

// This signal is dispatched when the camera shake effect completes.
func (self *Camera) GetOnShakeComplete() *Signal{
    return &Signal{self.Get("onShakeComplete")}
}

// This signal is dispatched when the camera shake effect completes.
func (self *Camera) SetOnShakeComplete(member *Signal) {
    self.Set("onShakeComplete", member)
}

// This signal is dispatched when the camera flash effect completes.
func (self *Camera) GetOnFlashComplete() *Signal{
    return &Signal{self.Get("onFlashComplete")}
}

// This signal is dispatched when the camera flash effect completes.
func (self *Camera) SetOnFlashComplete(member *Signal) {
    self.Set("onFlashComplete", member)
}

// This signal is dispatched when the camera fade effect completes.
// When the fade effect completes you will be left with the screen black (or whatever
// color you faded to). In order to reset this call `Camera.resetFX`. This is called
// automatically when you change State.
func (self *Camera) GetOnFadeComplete() *Signal{
    return &Signal{self.Get("onFadeComplete")}
}

// This signal is dispatched when the camera fade effect completes.
// When the fade effect completes you will be left with the screen black (or whatever
// color you faded to). In order to reset this call `Camera.resetFX`. This is called
// automatically when you change State.
func (self *Camera) SetOnFadeComplete(member *Signal) {
    self.Set("onFadeComplete", member)
}

// The Graphics object used to handle camera fx such as fade and flash.
func (self *Camera) GetFx() *Graphics{
    return &Graphics{self.Get("fx")}
}

// The Graphics object used to handle camera fx such as fade and flash.
func (self *Camera) SetFx(member *Graphics) {
    self.Set("fx", member)
}

// 
func (self *Camera) GetFOLLOW_LOCKON() int{
    return self.Get("FOLLOW_LOCKON").Int()
}

// 
func (self *Camera) SetFOLLOW_LOCKON(member int) {
    self.Set("FOLLOW_LOCKON", member)
}

// 
func (self *Camera) GetFOLLOW_PLATFORMER() int{
    return self.Get("FOLLOW_PLATFORMER").Int()
}

// 
func (self *Camera) SetFOLLOW_PLATFORMER(member int) {
    self.Set("FOLLOW_PLATFORMER", member)
}

// 
func (self *Camera) GetFOLLOW_TOPDOWN() int{
    return self.Get("FOLLOW_TOPDOWN").Int()
}

// 
func (self *Camera) SetFOLLOW_TOPDOWN(member int) {
    self.Set("FOLLOW_TOPDOWN", member)
}

// 
func (self *Camera) GetFOLLOW_TOPDOWN_TIGHT() int{
    return self.Get("FOLLOW_TOPDOWN_TIGHT").Int()
}

// 
func (self *Camera) SetFOLLOW_TOPDOWN_TIGHT(member int) {
    self.Set("FOLLOW_TOPDOWN_TIGHT", member)
}

// 
func (self *Camera) GetSHAKE_BOTH() int{
    return self.Get("SHAKE_BOTH").Int()
}

// 
func (self *Camera) SetSHAKE_BOTH(member int) {
    self.Set("SHAKE_BOTH", member)
}

// 
func (self *Camera) GetSHAKE_HORIZONTAL() int{
    return self.Get("SHAKE_HORIZONTAL").Int()
}

// 
func (self *Camera) SetSHAKE_HORIZONTAL(member int) {
    self.Set("SHAKE_HORIZONTAL", member)
}

// 
func (self *Camera) GetSHAKE_VERTICAL() int{
    return self.Get("SHAKE_VERTICAL").Int()
}

// 
func (self *Camera) SetSHAKE_VERTICAL(member int) {
    self.Set("SHAKE_VERTICAL", member)
}

// 
func (self *Camera) GetENABLE_FX() bool{
    return self.Get("ENABLE_FX").Bool()
}

// 
func (self *Camera) SetENABLE_FX(member bool) {
    self.Set("ENABLE_FX", member)
}

// The Cameras x coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras x position.
func (self *Camera) GetX() int{
    return self.Get("x").Int()
}

// The Cameras x coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras x position.
func (self *Camera) SetX(member int) {
    self.Set("x", member)
}

// The Cameras y coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras y position.
func (self *Camera) GetY() int{
    return self.Get("y").Int()
}

// The Cameras y coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras y position.
func (self *Camera) SetY(member int) {
    self.Set("y", member)
}

// The Cameras position. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras xy position using Phaser.Point object.
func (self *Camera) GetPosition() *Point{
    return &Point{self.Get("position")}
}

// The Cameras position. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras xy position using Phaser.Point object.
func (self *Camera) SetPosition(member *Point) {
    self.Set("position", member)
}

// The Cameras width. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras width.
func (self *Camera) GetWidth() int{
    return self.Get("width").Int()
}

// The Cameras width. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras width.
func (self *Camera) SetWidth(member int) {
    self.Set("width", member)
}

// The Cameras height. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras height.
func (self *Camera) GetHeight() int{
    return self.Get("height").Int()
}

// The Cameras height. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras height.
func (self *Camera) SetHeight(member int) {
    self.Set("height", member)
}

// The Cameras shake intensity. Gets or sets the cameras shake intensity.
func (self *Camera) GetShakeIntensity() int{
    return self.Get("shakeIntensity").Int()
}

// The Cameras shake intensity. Gets or sets the cameras shake intensity.
func (self *Camera) SetShakeIntensity(member int) {
    self.Set("shakeIntensity", member)
}



// Called automatically by Phaser.World.
func (self *Camera) BootI(args ...interface{}) {
    self.Call("boot", args)
}

// Camera preUpdate. Sets the total view counter to zero.
func (self *Camera) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) FollowI(args ...interface{}) {
    self.Call("follow", args)
}

// Sets the Camera follow target to null, stopping it from following an object if it's doing so.
func (self *Camera) UnfollowI(args ...interface{}) {
    self.Call("unfollow", args)
}

// Move the camera focus on a display object instantly.
func (self *Camera) FocusOnI(args ...interface{}) {
    self.Call("focusOn", args)
}

// Move the camera focus on a location instantly.
func (self *Camera) FocusOnXYI(args ...interface{}) {
    self.Call("focusOnXY", args)
}

// This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) ShakeI(args ...interface{}) bool{
    return self.Call("shake", args).Bool()
}

// This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) FlashI(args ...interface{}) bool{
    return self.Call("flash", args).Bool()
}

// This creates a camera fade effect. It works by filling the game with the
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
    return self.Call("fade", args).Bool()
}

// The camera update loop. This is called automatically by the core game loop.
func (self *Camera) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Update the camera flash and fade effects.
func (self *Camera) UpdateFXI(args ...interface{}) {
    self.Call("updateFX", args)
}

// Update the camera shake effect.
func (self *Camera) UpdateShakeI(args ...interface{}) {
    self.Call("updateShake", args)
}

// Internal method that handles tracking a sprite.
func (self *Camera) UpdateTargetI(args ...interface{}) {
    self.Call("updateTarget", args)
}

// Update the Camera bounds to match the game world.
func (self *Camera) SetBoundsToWorldI(args ...interface{}) {
    self.Call("setBoundsToWorld", args)
}

// Method called to ensure the camera doesn't venture outside of the game world.
// Called automatically by Camera.update.
func (self *Camera) CheckBoundsI(args ...interface{}) {
    self.Call("checkBounds", args)
}

// A helper function to set both the X and Y properties of the camera at once
// without having to use game.camera.x and game.camera.y.
func (self *Camera) SetPositionI(args ...interface{}) {
    self.Call("setPosition", args)
}

// Sets the size of the view rectangle given the width and height in parameters.
func (self *Camera) SetSizeI(args ...interface{}) {
    self.Call("setSize", args)
}

// Resets the camera back to 0,0 and un-follows any object it may have been tracking.
// Also immediately resets any camera effects that may have been running such as
// shake, flash or fade.
func (self *Camera) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Resets any active FX, such as a fade or flash and immediately clears it.
// Useful to calling after a fade in order to remove the fade from the Stage.
func (self *Camera) ResetFXI(args ...interface{}) {
    self.Call("resetFX", args)
}
