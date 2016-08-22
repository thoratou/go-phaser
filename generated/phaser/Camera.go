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
func (self *Camera) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *Camera) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// A reference to the game world.
func (self *Camera) GetWorldA() *World{
    return &World{self.Object.Get("world")}
}

// A reference to the game world.
func (self *Camera) SetWorldA(member *World) {
    self.Object.Set("world", member)
}

// Reserved for future multiple camera set-ups.
func (self *Camera) GetIdA() int{
    return self.Object.Get("id").Int()
}

// Reserved for future multiple camera set-ups.
func (self *Camera) SetIdA(member int) {
    self.Object.Set("id", member)
}

// Camera view.
// The view into the world we wish to render (by default the game dimensions).
// The x/y values are in world coordinates, not screen coordinates, the width/height is how many pixels to render.
// Sprites outside of this view are not rendered if Sprite.autoCull is set to `true`. Otherwise they are always rendered.
func (self *Camera) GetViewA() *Rectangle{
    return &Rectangle{self.Object.Get("view")}
}

// Camera view.
// The view into the world we wish to render (by default the game dimensions).
// The x/y values are in world coordinates, not screen coordinates, the width/height is how many pixels to render.
// Sprites outside of this view are not rendered if Sprite.autoCull is set to `true`. Otherwise they are always rendered.
func (self *Camera) SetViewA(member *Rectangle) {
    self.Object.Set("view", member)
}

// The Camera is bound to this Rectangle and cannot move outside of it. By default it is enabled and set to the size of the World.
// The Rectangle can be located anywhere in the world and updated as often as you like. If you don't wish the Camera to be bound
// at all then set this to null. The values can be anything and are in World coordinates, with 0,0 being the top-left of the world. The Rectangle in which the Camera is bounded. Set to null to allow for movement anywhere.
func (self *Camera) GetBoundsA() *Rectangle{
    return &Rectangle{self.Object.Get("bounds")}
}

// The Camera is bound to this Rectangle and cannot move outside of it. By default it is enabled and set to the size of the World.
// The Rectangle can be located anywhere in the world and updated as often as you like. If you don't wish the Camera to be bound
// at all then set this to null. The values can be anything and are in World coordinates, with 0,0 being the top-left of the world. The Rectangle in which the Camera is bounded. Set to null to allow for movement anywhere.
func (self *Camera) SetBoundsA(member *Rectangle) {
    self.Object.Set("bounds", member)
}

// Moving inside this Rectangle will not cause the camera to move.
func (self *Camera) GetDeadzoneA() *Rectangle{
    return &Rectangle{self.Object.Get("deadzone")}
}

// Moving inside this Rectangle will not cause the camera to move.
func (self *Camera) SetDeadzoneA(member *Rectangle) {
    self.Object.Set("deadzone", member)
}

// Whether this camera is visible or not.
func (self *Camera) GetVisibleA() bool{
    return self.Object.Get("visible").Bool()
}

// Whether this camera is visible or not.
func (self *Camera) SetVisibleA(member bool) {
    self.Object.Set("visible", member)
}

// If a Camera has roundPx set to `true` it will call `view.floor` as part of its update loop, keeping its boundary to integer values. Set this to `false` to disable this from happening.
func (self *Camera) GetRoundPxA() bool{
    return self.Object.Get("roundPx").Bool()
}

// If a Camera has roundPx set to `true` it will call `view.floor` as part of its update loop, keeping its boundary to integer values. Set this to `false` to disable this from happening.
func (self *Camera) SetRoundPxA(member bool) {
    self.Object.Set("roundPx", member)
}

// Whether this camera is flush with the World Bounds or not.
func (self *Camera) GetAtLimitA() bool{
    return self.Object.Get("atLimit").Bool()
}

// Whether this camera is flush with the World Bounds or not.
func (self *Camera) SetAtLimitA(member bool) {
    self.Object.Set("atLimit", member)
}

// If the camera is tracking a Sprite, this is a reference to it, otherwise null.
func (self *Camera) GetTargetA() *Sprite{
    return &Sprite{self.Object.Get("target")}
}

// If the camera is tracking a Sprite, this is a reference to it, otherwise null.
func (self *Camera) SetTargetA(member *Sprite) {
    self.Object.Set("target", member)
}

// The display object to which all game objects are added. Set by World.boot.
func (self *Camera) GetDisplayObjectA() *DisplayObject{
    return &DisplayObject{self.Object.Get("displayObject")}
}

// The display object to which all game objects are added. Set by World.boot.
func (self *Camera) SetDisplayObjectA(member *DisplayObject) {
    self.Object.Set("displayObject", member)
}

// The scale of the display object to which all game objects are added. Set by World.boot.
func (self *Camera) GetScaleA() *Point{
    return &Point{self.Object.Get("scale")}
}

// The scale of the display object to which all game objects are added. Set by World.boot.
func (self *Camera) SetScaleA(member *Point) {
    self.Object.Set("scale", member)
}

// The total number of Sprites with `autoCull` set to `true` that are visible by this Camera.
func (self *Camera) GetTotalInViewA() int{
    return self.Object.Get("totalInView").Int()
}

// The total number of Sprites with `autoCull` set to `true` that are visible by this Camera.
func (self *Camera) SetTotalInViewA(member int) {
    self.Object.Set("totalInView", member)
}

// The linear interpolation value to use when following a target.
// The default values of 1 means the camera will instantly snap to the target coordinates.
// A lower value, such as 0.1 means the camera will more slowly track the target, giving
// a smooth transition. You can set the horizontal and vertical values independently, and also
// adjust this value in real-time during your game.
func (self *Camera) GetLerpA() *Point{
    return &Point{self.Object.Get("lerp")}
}

// The linear interpolation value to use when following a target.
// The default values of 1 means the camera will instantly snap to the target coordinates.
// A lower value, such as 0.1 means the camera will more slowly track the target, giving
// a smooth transition. You can set the horizontal and vertical values independently, and also
// adjust this value in real-time during your game.
func (self *Camera) SetLerpA(member *Point) {
    self.Object.Set("lerp", member)
}

// This signal is dispatched when the camera shake effect completes.
func (self *Camera) GetOnShakeCompleteA() *Signal{
    return &Signal{self.Object.Get("onShakeComplete")}
}

// This signal is dispatched when the camera shake effect completes.
func (self *Camera) SetOnShakeCompleteA(member *Signal) {
    self.Object.Set("onShakeComplete", member)
}

// This signal is dispatched when the camera flash effect completes.
func (self *Camera) GetOnFlashCompleteA() *Signal{
    return &Signal{self.Object.Get("onFlashComplete")}
}

// This signal is dispatched when the camera flash effect completes.
func (self *Camera) SetOnFlashCompleteA(member *Signal) {
    self.Object.Set("onFlashComplete", member)
}

// This signal is dispatched when the camera fade effect completes.
// When the fade effect completes you will be left with the screen black (or whatever
// color you faded to). In order to reset this call `Camera.resetFX`. This is called
// automatically when you change State.
func (self *Camera) GetOnFadeCompleteA() *Signal{
    return &Signal{self.Object.Get("onFadeComplete")}
}

// This signal is dispatched when the camera fade effect completes.
// When the fade effect completes you will be left with the screen black (or whatever
// color you faded to). In order to reset this call `Camera.resetFX`. This is called
// automatically when you change State.
func (self *Camera) SetOnFadeCompleteA(member *Signal) {
    self.Object.Set("onFadeComplete", member)
}

// The Graphics object used to handle camera fx such as fade and flash.
func (self *Camera) GetFxA() *Graphics{
    return &Graphics{self.Object.Get("fx")}
}

// The Graphics object used to handle camera fx such as fade and flash.
func (self *Camera) SetFxA(member *Graphics) {
    self.Object.Set("fx", member)
}

// 
func (self *Camera) GetFOLLOW_LOCKONA() int{
    return self.Object.Get("FOLLOW_LOCKON").Int()
}

// 
func (self *Camera) SetFOLLOW_LOCKONA(member int) {
    self.Object.Set("FOLLOW_LOCKON", member)
}

// 
func (self *Camera) GetFOLLOW_PLATFORMERA() int{
    return self.Object.Get("FOLLOW_PLATFORMER").Int()
}

// 
func (self *Camera) SetFOLLOW_PLATFORMERA(member int) {
    self.Object.Set("FOLLOW_PLATFORMER", member)
}

// 
func (self *Camera) GetFOLLOW_TOPDOWNA() int{
    return self.Object.Get("FOLLOW_TOPDOWN").Int()
}

// 
func (self *Camera) SetFOLLOW_TOPDOWNA(member int) {
    self.Object.Set("FOLLOW_TOPDOWN", member)
}

// 
func (self *Camera) GetFOLLOW_TOPDOWN_TIGHTA() int{
    return self.Object.Get("FOLLOW_TOPDOWN_TIGHT").Int()
}

// 
func (self *Camera) SetFOLLOW_TOPDOWN_TIGHTA(member int) {
    self.Object.Set("FOLLOW_TOPDOWN_TIGHT", member)
}

// 
func (self *Camera) GetSHAKE_BOTHA() int{
    return self.Object.Get("SHAKE_BOTH").Int()
}

// 
func (self *Camera) SetSHAKE_BOTHA(member int) {
    self.Object.Set("SHAKE_BOTH", member)
}

// 
func (self *Camera) GetSHAKE_HORIZONTALA() int{
    return self.Object.Get("SHAKE_HORIZONTAL").Int()
}

// 
func (self *Camera) SetSHAKE_HORIZONTALA(member int) {
    self.Object.Set("SHAKE_HORIZONTAL", member)
}

// 
func (self *Camera) GetSHAKE_VERTICALA() int{
    return self.Object.Get("SHAKE_VERTICAL").Int()
}

// 
func (self *Camera) SetSHAKE_VERTICALA(member int) {
    self.Object.Set("SHAKE_VERTICAL", member)
}

// 
func (self *Camera) GetENABLE_FXA() bool{
    return self.Object.Get("ENABLE_FX").Bool()
}

// 
func (self *Camera) SetENABLE_FXA(member bool) {
    self.Object.Set("ENABLE_FX", member)
}

// The Cameras x coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras x position.
func (self *Camera) GetXA() int{
    return self.Object.Get("x").Int()
}

// The Cameras x coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras x position.
func (self *Camera) SetXA(member int) {
    self.Object.Set("x", member)
}

// The Cameras y coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras y position.
func (self *Camera) GetYA() int{
    return self.Object.Get("y").Int()
}

// The Cameras y coordinate. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras y position.
func (self *Camera) SetYA(member int) {
    self.Object.Set("y", member)
}

// The Cameras position. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras xy position using Phaser.Point object.
func (self *Camera) GetPositionA() *Point{
    return &Point{self.Object.Get("position")}
}

// The Cameras position. This value is automatically clamped if it falls outside of the World bounds. Gets or sets the cameras xy position using Phaser.Point object.
func (self *Camera) SetPositionA(member *Point) {
    self.Object.Set("position", member)
}

// The Cameras width. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras width.
func (self *Camera) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The Cameras width. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras width.
func (self *Camera) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The Cameras height. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras height.
func (self *Camera) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The Cameras height. By default this is the same as the Game size and should not be adjusted for now. Gets or sets the cameras height.
func (self *Camera) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The Cameras shake intensity. Gets or sets the cameras shake intensity.
func (self *Camera) GetShakeIntensityA() int{
    return self.Object.Get("shakeIntensity").Int()
}

// The Cameras shake intensity. Gets or sets the cameras shake intensity.
func (self *Camera) SetShakeIntensityA(member int) {
    self.Object.Set("shakeIntensity", member)
}



// Called automatically by Phaser.World.
func (self *Camera) Boot() {
    self.Object.Call("boot")
}

// Called automatically by Phaser.World.
func (self *Camera) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// Camera preUpdate. Sets the total view counter to zero.
func (self *Camera) PreUpdate() {
    self.Object.Call("preUpdate")
}

// Camera preUpdate. Sets the total view counter to zero.
func (self *Camera) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) Follow(target interface{}) {
    self.Object.Call("follow", target)
}

// Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) Follow1O(target interface{}, style int) {
    self.Object.Call("follow", target, style)
}

// Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) Follow2O(target interface{}, style int, lerpX float64) {
    self.Object.Call("follow", target, style, lerpX)
}

// Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) Follow3O(target interface{}, style int, lerpX float64, lerpY float64) {
    self.Object.Call("follow", target, style, lerpX, lerpY)
}

// Tell the camera which sprite to follow.
// 
// You can set the follow type and a linear interpolation value.
// Use low lerp values (such as 0.1) to automatically smooth the camera motion.
// 
// If you find you're getting a slight "jitter" effect when following a Sprite it's probably to do with sub-pixel rendering of the Sprite position.
// This can be disabled by setting `game.renderer.renderSession.roundPixels = true` to force full pixel rendering.
func (self *Camera) FollowI(args ...interface{}) {
    self.Object.Call("follow", args)
}

// Sets the Camera follow target to null, stopping it from following an object if it's doing so.
func (self *Camera) Unfollow() {
    self.Object.Call("unfollow")
}

// Sets the Camera follow target to null, stopping it from following an object if it's doing so.
func (self *Camera) UnfollowI(args ...interface{}) {
    self.Object.Call("unfollow", args)
}

// Move the camera focus on a display object instantly.
func (self *Camera) FocusOn(displayObject interface{}) {
    self.Object.Call("focusOn", displayObject)
}

// Move the camera focus on a display object instantly.
func (self *Camera) FocusOnI(args ...interface{}) {
    self.Object.Call("focusOn", args)
}

// Move the camera focus on a location instantly.
func (self *Camera) FocusOnXY(x int, y int) {
    self.Object.Call("focusOnXY", x, y)
}

// Move the camera focus on a location instantly.
func (self *Camera) FocusOnXYI(args ...interface{}) {
    self.Object.Call("focusOnXY", args)
}

// This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake() bool{
    return self.Object.Call("shake").Bool()
}

// This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake1O(intensity float64) bool{
    return self.Object.Call("shake", intensity).Bool()
}

// This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake2O(intensity float64, duration int) bool{
    return self.Object.Call("shake", intensity, duration).Bool()
}

// This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake3O(intensity float64, duration int, force bool) bool{
    return self.Object.Call("shake", intensity, duration, force).Bool()
}

// This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake4O(intensity float64, duration int, force bool, direction int) bool{
    return self.Object.Call("shake", intensity, duration, force, direction).Bool()
}

// This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) Shake5O(intensity float64, duration int, force bool, direction int, shakeBounds bool) bool{
    return self.Object.Call("shake", intensity, duration, force, direction, shakeBounds).Bool()
}

// This creates a camera shake effect. It works by applying a random amount of additional
// spacing on the x and y axis each frame. You can control the intensity and duration
// of the effect, and if it should effect both axis or just one.
// 
// When the shake effect ends the signal Camera.onShakeComplete is dispatched.
func (self *Camera) ShakeI(args ...interface{}) bool{
    return self.Object.Call("shake", args).Bool()
}

// This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) Flash() bool{
    return self.Object.Call("flash").Bool()
}

// This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) Flash1O(color float64) bool{
    return self.Object.Call("flash", color).Bool()
}

// This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) Flash2O(color float64, duration int) bool{
    return self.Object.Call("flash", color, duration).Bool()
}

// This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) Flash3O(color float64, duration int, force bool) bool{
    return self.Object.Call("flash", color, duration, force).Bool()
}

// This creates a camera flash effect. It works by filling the game with the solid fill
// color specified, and then fading it away to alpha 0 over the duration given.
// 
// You can use this for things such as hit feedback effects.
// 
// When the effect ends the signal Camera.onFlashComplete is dispatched.
func (self *Camera) FlashI(args ...interface{}) bool{
    return self.Object.Call("flash", args).Bool()
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
func (self *Camera) Fade() bool{
    return self.Object.Call("fade").Bool()
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
func (self *Camera) Fade1O(color float64) bool{
    return self.Object.Call("fade", color).Bool()
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
func (self *Camera) Fade2O(color float64, duration int) bool{
    return self.Object.Call("fade", color, duration).Bool()
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
func (self *Camera) Fade3O(color float64, duration int, force bool) bool{
    return self.Object.Call("fade", color, duration, force).Bool()
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
    return self.Object.Call("fade", args).Bool()
}

// The camera update loop. This is called automatically by the core game loop.
func (self *Camera) Update() {
    self.Object.Call("update")
}

// The camera update loop. This is called automatically by the core game loop.
func (self *Camera) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Update the camera flash and fade effects.
func (self *Camera) UpdateFX() {
    self.Object.Call("updateFX")
}

// Update the camera flash and fade effects.
func (self *Camera) UpdateFXI(args ...interface{}) {
    self.Object.Call("updateFX", args)
}

// Update the camera shake effect.
func (self *Camera) UpdateShake() {
    self.Object.Call("updateShake")
}

// Update the camera shake effect.
func (self *Camera) UpdateShakeI(args ...interface{}) {
    self.Object.Call("updateShake", args)
}

// Internal method that handles tracking a sprite.
func (self *Camera) UpdateTarget() {
    self.Object.Call("updateTarget")
}

// Internal method that handles tracking a sprite.
func (self *Camera) UpdateTargetI(args ...interface{}) {
    self.Object.Call("updateTarget", args)
}

// Update the Camera bounds to match the game world.
func (self *Camera) SetBoundsToWorld() {
    self.Object.Call("setBoundsToWorld")
}

// Update the Camera bounds to match the game world.
func (self *Camera) SetBoundsToWorldI(args ...interface{}) {
    self.Object.Call("setBoundsToWorld", args)
}

// Method called to ensure the camera doesn't venture outside of the game world.
// Called automatically by Camera.update.
func (self *Camera) CheckBounds() {
    self.Object.Call("checkBounds")
}

// Method called to ensure the camera doesn't venture outside of the game world.
// Called automatically by Camera.update.
func (self *Camera) CheckBoundsI(args ...interface{}) {
    self.Object.Call("checkBounds", args)
}

// A helper function to set both the X and Y properties of the camera at once
// without having to use game.camera.x and game.camera.y.
func (self *Camera) SetPosition(x int, y int) {
    self.Object.Call("setPosition", x, y)
}

// A helper function to set both the X and Y properties of the camera at once
// without having to use game.camera.x and game.camera.y.
func (self *Camera) SetPositionI(args ...interface{}) {
    self.Object.Call("setPosition", args)
}

// Sets the size of the view rectangle given the width and height in parameters.
func (self *Camera) SetSize(width int, height int) {
    self.Object.Call("setSize", width, height)
}

// Sets the size of the view rectangle given the width and height in parameters.
func (self *Camera) SetSizeI(args ...interface{}) {
    self.Object.Call("setSize", args)
}

// Resets the camera back to 0,0 and un-follows any object it may have been tracking.
// Also immediately resets any camera effects that may have been running such as
// shake, flash or fade.
func (self *Camera) Reset() {
    self.Object.Call("reset")
}

// Resets the camera back to 0,0 and un-follows any object it may have been tracking.
// Also immediately resets any camera effects that may have been running such as
// shake, flash or fade.
func (self *Camera) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Resets any active FX, such as a fade or flash and immediately clears it.
// Useful to calling after a fade in order to remove the fade from the Stage.
func (self *Camera) ResetFX() {
    self.Object.Call("resetFX")
}

// Resets any active FX, such as a fade or flash and immediately clears it.
// Useful to calling after a fade in order to remove the fade from the Stage.
func (self *Camera) ResetFXI(args ...interface{}) {
    self.Object.Call("resetFX", args)
}
