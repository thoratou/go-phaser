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
func (self *AnimationManager) GetSpriteA() *Sprite{
    return &Sprite{self.Object.Get("sprite")}
}

// A reference to the parent Sprite that owns this AnimationManager.
func (self *AnimationManager) SetSpriteA(member *Sprite) {
    self.Object.Set("sprite", member)
}

// A reference to the currently running Game.
func (self *AnimationManager) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running Game.
func (self *AnimationManager) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The currently displayed Frame of animation, if any.
// This property is only set once an Animation starts playing. Until that point it remains set as `null`.
func (self *AnimationManager) GetCurrentFrameA() *Frame{
    return &Frame{self.Object.Get("currentFrame")}
}

// The currently displayed Frame of animation, if any.
// This property is only set once an Animation starts playing. Until that point it remains set as `null`.
func (self *AnimationManager) SetCurrentFrameA(member *Frame) {
    self.Object.Set("currentFrame", member)
}

// The currently displayed animation, if any.
func (self *AnimationManager) GetCurrentAnimA() *Animation{
    return &Animation{self.Object.Get("currentAnim")}
}

// The currently displayed animation, if any.
func (self *AnimationManager) SetCurrentAnimA(member *Animation) {
    self.Object.Set("currentAnim", member)
}

// Should the animation data continue to update even if the Sprite.visible is set to false.
func (self *AnimationManager) GetUpdateIfVisibleA() bool{
    return self.Object.Get("updateIfVisible").Bool()
}

// Should the animation data continue to update even if the Sprite.visible is set to false.
func (self *AnimationManager) SetUpdateIfVisibleA(member bool) {
    self.Object.Set("updateIfVisible", member)
}

// Set to true once animation data has been loaded.
func (self *AnimationManager) GetIsLoadedA() bool{
    return self.Object.Get("isLoaded").Bool()
}

// Set to true once animation data has been loaded.
func (self *AnimationManager) SetIsLoadedA(member bool) {
    self.Object.Set("isLoaded", member)
}

// The current animations FrameData.
func (self *AnimationManager) GetFrameDataA() *FrameData{
    return &FrameData{self.Object.Get("frameData")}
}

// The current animations FrameData.
func (self *AnimationManager) SetFrameDataA(member *FrameData) {
    self.Object.Set("frameData", member)
}

// The total number of frames in the currently loaded FrameData, or -1 if no FrameData is loaded.
func (self *AnimationManager) GetFrameTotalA() int{
    return self.Object.Get("frameTotal").Int()
}

// The total number of frames in the currently loaded FrameData, or -1 if no FrameData is loaded.
func (self *AnimationManager) SetFrameTotalA(member int) {
    self.Object.Set("frameTotal", member)
}

// Gets and sets the paused state of the current animation.
func (self *AnimationManager) GetPausedA() bool{
    return self.Object.Get("paused").Bool()
}

// Gets and sets the paused state of the current animation.
func (self *AnimationManager) SetPausedA(member bool) {
    self.Object.Set("paused", member)
}

// Gets the current animation name, if set.
func (self *AnimationManager) GetNameA() string{
    return self.Object.Get("name").String()
}

// Gets the current animation name, if set.
func (self *AnimationManager) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Gets or sets the current frame index and updates the Texture Cache for display.
func (self *AnimationManager) GetFrameA() int{
    return self.Object.Get("frame").Int()
}

// Gets or sets the current frame index and updates the Texture Cache for display.
func (self *AnimationManager) SetFrameA(member int) {
    self.Object.Set("frame", member)
}

// Gets or sets the current frame name and updates the Texture Cache for display.
func (self *AnimationManager) GetFrameNameA() string{
    return self.Object.Get("frameName").String()
}

// Gets or sets the current frame name and updates the Texture Cache for display.
func (self *AnimationManager) SetFrameNameA(member string) {
    self.Object.Set("frameName", member)
}



// Loads FrameData into the internal temporary vars and resets the frame index to zero.
// This is called automatically when a new Sprite is created.
func (self *AnimationManager) LoadFrameData(frameData *FrameData, frame interface{}) bool{
    return self.Object.Call("loadFrameData", frameData, frame).Bool()
}

// Loads FrameData into the internal temporary vars and resets the frame index to zero.
// This is called automatically when a new Sprite is created.
func (self *AnimationManager) LoadFrameDataI(args ...interface{}) bool{
    return self.Object.Call("loadFrameData", args).Bool()
}

// Loads FrameData into the internal temporary vars and resets the frame index to zero.
// This is called automatically when a new Sprite is created.
func (self *AnimationManager) CopyFrameData(frameData *FrameData, frame interface{}) bool{
    return self.Object.Call("copyFrameData", frameData, frame).Bool()
}

// Loads FrameData into the internal temporary vars and resets the frame index to zero.
// This is called automatically when a new Sprite is created.
func (self *AnimationManager) CopyFrameDataI(args ...interface{}) bool{
    return self.Object.Call("copyFrameData", args).Bool()
}

// Adds a new animation under the given key. Optionally set the frames, frame rate and loop.
// Animations added in this way are played back with the play function.
func (self *AnimationManager) Add(name string) *Animation{
    return &Animation{self.Object.Call("add", name)}
}

// Adds a new animation under the given key. Optionally set the frames, frame rate and loop.
// Animations added in this way are played back with the play function.
func (self *AnimationManager) Add1O(name string, frames []interface{}) *Animation{
    return &Animation{self.Object.Call("add", name, frames)}
}

// Adds a new animation under the given key. Optionally set the frames, frame rate and loop.
// Animations added in this way are played back with the play function.
func (self *AnimationManager) Add2O(name string, frames []interface{}, frameRate int) *Animation{
    return &Animation{self.Object.Call("add", name, frames, frameRate)}
}

// Adds a new animation under the given key. Optionally set the frames, frame rate and loop.
// Animations added in this way are played back with the play function.
func (self *AnimationManager) Add3O(name string, frames []interface{}, frameRate int, loop bool) *Animation{
    return &Animation{self.Object.Call("add", name, frames, frameRate, loop)}
}

// Adds a new animation under the given key. Optionally set the frames, frame rate and loop.
// Animations added in this way are played back with the play function.
func (self *AnimationManager) Add4O(name string, frames []interface{}, frameRate int, loop bool, useNumericIndex bool) *Animation{
    return &Animation{self.Object.Call("add", name, frames, frameRate, loop, useNumericIndex)}
}

// Adds a new animation under the given key. Optionally set the frames, frame rate and loop.
// Animations added in this way are played back with the play function.
func (self *AnimationManager) AddI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("add", args)}
}

// Check whether the frames in the given array are valid and exist.
func (self *AnimationManager) ValidateFrames(frames []interface{}) bool{
    return self.Object.Call("validateFrames", frames).Bool()
}

// Check whether the frames in the given array are valid and exist.
func (self *AnimationManager) ValidateFrames1O(frames []interface{}, useNumericIndex bool) bool{
    return self.Object.Call("validateFrames", frames, useNumericIndex).Bool()
}

// Check whether the frames in the given array are valid and exist.
func (self *AnimationManager) ValidateFramesI(args ...interface{}) bool{
    return self.Object.Call("validateFrames", args).Bool()
}

// Play an animation based on the given key. The animation should previously have been added via `animations.add`
// 
// If the requested animation is already playing this request will be ignored. 
// If you need to reset an already running animation do so directly on the Animation object itself.
func (self *AnimationManager) Play(name string) *Animation{
    return &Animation{self.Object.Call("play", name)}
}

// Play an animation based on the given key. The animation should previously have been added via `animations.add`
// 
// If the requested animation is already playing this request will be ignored. 
// If you need to reset an already running animation do so directly on the Animation object itself.
func (self *AnimationManager) Play1O(name string, frameRate int) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate)}
}

// Play an animation based on the given key. The animation should previously have been added via `animations.add`
// 
// If the requested animation is already playing this request will be ignored. 
// If you need to reset an already running animation do so directly on the Animation object itself.
func (self *AnimationManager) Play2O(name string, frameRate int, loop bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop)}
}

// Play an animation based on the given key. The animation should previously have been added via `animations.add`
// 
// If the requested animation is already playing this request will be ignored. 
// If you need to reset an already running animation do so directly on the Animation object itself.
func (self *AnimationManager) Play3O(name string, frameRate int, loop bool, killOnComplete bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop, killOnComplete)}
}

// Play an animation based on the given key. The animation should previously have been added via `animations.add`
// 
// If the requested animation is already playing this request will be ignored. 
// If you need to reset an already running animation do so directly on the Animation object itself.
func (self *AnimationManager) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("play", args)}
}

// Stop playback of an animation. If a name is given that specific animation is stopped, otherwise the current animation is stopped.
// The currentAnim property of the AnimationManager is automatically set to the animation given.
func (self *AnimationManager) Stop() {
    self.Object.Call("stop")
}

// Stop playback of an animation. If a name is given that specific animation is stopped, otherwise the current animation is stopped.
// The currentAnim property of the AnimationManager is automatically set to the animation given.
func (self *AnimationManager) Stop1O(name string) {
    self.Object.Call("stop", name)
}

// Stop playback of an animation. If a name is given that specific animation is stopped, otherwise the current animation is stopped.
// The currentAnim property of the AnimationManager is automatically set to the animation given.
func (self *AnimationManager) Stop2O(name string, resetFrame bool) {
    self.Object.Call("stop", name, resetFrame)
}

// Stop playback of an animation. If a name is given that specific animation is stopped, otherwise the current animation is stopped.
// The currentAnim property of the AnimationManager is automatically set to the animation given.
func (self *AnimationManager) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// The main update function is called by the Sprites update loop. It's responsible for updating animation frames and firing related events.
func (self *AnimationManager) Update() bool{
    return self.Object.Call("update").Bool()
}

// The main update function is called by the Sprites update loop. It's responsible for updating animation frames and firing related events.
func (self *AnimationManager) UpdateI(args ...interface{}) bool{
    return self.Object.Call("update", args).Bool()
}

// Advances by the given number of frames in the current animation, taking the loop value into consideration.
func (self *AnimationManager) Next() {
    self.Object.Call("next")
}

// Advances by the given number of frames in the current animation, taking the loop value into consideration.
func (self *AnimationManager) Next1O(quantity int) {
    self.Object.Call("next", quantity)
}

// Advances by the given number of frames in the current animation, taking the loop value into consideration.
func (self *AnimationManager) NextI(args ...interface{}) {
    self.Object.Call("next", args)
}

// Moves backwards the given number of frames in the current animation, taking the loop value into consideration.
func (self *AnimationManager) Previous() {
    self.Object.Call("previous")
}

// Moves backwards the given number of frames in the current animation, taking the loop value into consideration.
func (self *AnimationManager) Previous1O(quantity int) {
    self.Object.Call("previous", quantity)
}

// Moves backwards the given number of frames in the current animation, taking the loop value into consideration.
func (self *AnimationManager) PreviousI(args ...interface{}) {
    self.Object.Call("previous", args)
}

// Returns an animation that was previously added by name.
func (self *AnimationManager) GetAnimation(name string) *Animation{
    return &Animation{self.Object.Call("getAnimation", name)}
}

// Returns an animation that was previously added by name.
func (self *AnimationManager) GetAnimationI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("getAnimation", args)}
}

// Refreshes the current frame data back to the parent Sprite and also resets the texture data.
func (self *AnimationManager) RefreshFrame() {
    self.Object.Call("refreshFrame")
}

// Refreshes the current frame data back to the parent Sprite and also resets the texture data.
func (self *AnimationManager) RefreshFrameI(args ...interface{}) {
    self.Object.Call("refreshFrame", args)
}

// Destroys all references this AnimationManager contains.
// Iterates through the list of animations stored in this manager and calls destroy on each of them.
func (self *AnimationManager) Destroy() {
    self.Object.Call("destroy")
}

// Destroys all references this AnimationManager contains.
// Iterates through the list of animations stored in this manager and calls destroy on each of them.
func (self *AnimationManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
