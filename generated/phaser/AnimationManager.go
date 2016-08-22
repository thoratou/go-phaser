// Automatic generation for Phaser.AnimationManager
// generated file AnimationManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Animation Manager is used to add, play and update Phaser Animations.
// Any Game Object such as Phaser.Sprite that supports animation contains a single AnimationManager instance.
type AnimationManager struct {
    *js.Object
}


// A reference to the parent Sprite that owns this AnimationManager.
func (self *AnimationManager) GetSprite() *Sprite{
    return &Sprite{self.Get("sprite")}
}

// A reference to the parent Sprite that owns this AnimationManager.
func (self *AnimationManager) SetSprite(member *Sprite) {
    self.Set("sprite", member)
}

// A reference to the currently running Game.
func (self *AnimationManager) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *AnimationManager) SetGame(member *Game) {
    self.Set("game", member)
}

// The currently displayed Frame of animation, if any.
// This property is only set once an Animation starts playing. Until that point it remains set as `null`.
func (self *AnimationManager) GetCurrentFrame() *Frame{
    return &Frame{self.Get("currentFrame")}
}

// The currently displayed Frame of animation, if any.
// This property is only set once an Animation starts playing. Until that point it remains set as `null`.
func (self *AnimationManager) SetCurrentFrame(member *Frame) {
    self.Set("currentFrame", member)
}

// The currently displayed animation, if any.
func (self *AnimationManager) GetCurrentAnim() *Animation{
    return &Animation{self.Get("currentAnim")}
}

// The currently displayed animation, if any.
func (self *AnimationManager) SetCurrentAnim(member *Animation) {
    self.Set("currentAnim", member)
}

// Should the animation data continue to update even if the Sprite.visible is set to false.
func (self *AnimationManager) GetUpdateIfVisible() bool{
    return self.Get("updateIfVisible").Bool()
}

// Should the animation data continue to update even if the Sprite.visible is set to false.
func (self *AnimationManager) SetUpdateIfVisible(member bool) {
    self.Set("updateIfVisible", member)
}

// Set to true once animation data has been loaded.
func (self *AnimationManager) GetIsLoaded() bool{
    return self.Get("isLoaded").Bool()
}

// Set to true once animation data has been loaded.
func (self *AnimationManager) SetIsLoaded(member bool) {
    self.Set("isLoaded", member)
}

// The current animations FrameData.
func (self *AnimationManager) GetFrameData() *FrameData{
    return &FrameData{self.Get("frameData")}
}

// The current animations FrameData.
func (self *AnimationManager) SetFrameData(member *FrameData) {
    self.Set("frameData", member)
}

// The total number of frames in the currently loaded FrameData, or -1 if no FrameData is loaded.
func (self *AnimationManager) GetFrameTotal() int{
    return self.Get("frameTotal").Int()
}

// The total number of frames in the currently loaded FrameData, or -1 if no FrameData is loaded.
func (self *AnimationManager) SetFrameTotal(member int) {
    self.Set("frameTotal", member)
}

// Gets and sets the paused state of the current animation.
func (self *AnimationManager) GetPaused() bool{
    return self.Get("paused").Bool()
}

// Gets and sets the paused state of the current animation.
func (self *AnimationManager) SetPaused(member bool) {
    self.Set("paused", member)
}

// Gets the current animation name, if set.
func (self *AnimationManager) GetName() string{
    return self.Get("name").String()
}

// Gets the current animation name, if set.
func (self *AnimationManager) SetName(member string) {
    self.Set("name", member)
}

// Gets or sets the current frame index and updates the Texture Cache for display.
func (self *AnimationManager) GetFrame() int{
    return self.Get("frame").Int()
}

// Gets or sets the current frame index and updates the Texture Cache for display.
func (self *AnimationManager) SetFrame(member int) {
    self.Set("frame", member)
}

// Gets or sets the current frame name and updates the Texture Cache for display.
func (self *AnimationManager) GetFrameName() string{
    return self.Get("frameName").String()
}

// Gets or sets the current frame name and updates the Texture Cache for display.
func (self *AnimationManager) SetFrameName(member string) {
    self.Set("frameName", member)
}



// Loads FrameData into the internal temporary vars and resets the frame index to zero.
// This is called automatically when a new Sprite is created.
func (self *AnimationManager) LoadFrameDataI(args ...interface{}) bool{
    return self.Call("loadFrameData", args).Bool()
}

// Loads FrameData into the internal temporary vars and resets the frame index to zero.
// This is called automatically when a new Sprite is created.
func (self *AnimationManager) CopyFrameDataI(args ...interface{}) bool{
    return self.Call("copyFrameData", args).Bool()
}

// Adds a new animation under the given key. Optionally set the frames, frame rate and loop.
// Animations added in this way are played back with the play function.
func (self *AnimationManager) AddI(args ...interface{}) *Animation{
    return &Animation{self.Call("add", args)}
}

// Check whether the frames in the given array are valid and exist.
func (self *AnimationManager) ValidateFramesI(args ...interface{}) bool{
    return self.Call("validateFrames", args).Bool()
}

// Play an animation based on the given key. The animation should previously have been added via `animations.add`
// 
// If the requested animation is already playing this request will be ignored. 
// If you need to reset an already running animation do so directly on the Animation object itself.
func (self *AnimationManager) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Call("play", args)}
}

// Stop playback of an animation. If a name is given that specific animation is stopped, otherwise the current animation is stopped.
// The currentAnim property of the AnimationManager is automatically set to the animation given.
func (self *AnimationManager) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// The main update function is called by the Sprites update loop. It's responsible for updating animation frames and firing related events.
func (self *AnimationManager) UpdateI(args ...interface{}) bool{
    return self.Call("update", args).Bool()
}

// Advances by the given number of frames in the current animation, taking the loop value into consideration.
func (self *AnimationManager) NextI(args ...interface{}) {
    self.Call("next", args)
}

// Moves backwards the given number of frames in the current animation, taking the loop value into consideration.
func (self *AnimationManager) PreviousI(args ...interface{}) {
    self.Call("previous", args)
}

// Returns an animation that was previously added by name.
func (self *AnimationManager) GetAnimationI(args ...interface{}) *Animation{
    return &Animation{self.Call("getAnimation", args)}
}

// Refreshes the current frame data back to the parent Sprite and also resets the texture data.
func (self *AnimationManager) RefreshFrameI(args ...interface{}) {
    self.Call("refreshFrame", args)
}

// Destroys all references this AnimationManager contains.
// Iterates through the list of animations stored in this manager and calls destroy on each of them.
func (self *AnimationManager) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
