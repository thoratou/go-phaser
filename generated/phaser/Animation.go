// Package phaser Automatic generation for Phaser.Animation
// generated file Animation.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Animation An Animation instance contains a single animation and the controls to play it.
// 
// It is created by the AnimationManager, consists of Animation.Frame objects and belongs to a single Game Object such as a Sprite.
type Animation struct {
    *js.Object
}

// NewAnimation An Animation instance contains a single animation and the controls to play it.
// 
// It is created by the AnimationManager, consists of Animation.Frame objects and belongs to a single Game Object such as a Sprite.
func NewAnimation(game *Game, parent *Sprite, name string, frameData *FrameData, frames interface{}) *Animation {
    return &Animation{js.Global.Get("Phaser").Get("Animation").New(game, parent, name, frameData, frames)}
}
// NewAnimation1O An Animation instance contains a single animation and the controls to play it.
// 
// It is created by the AnimationManager, consists of Animation.Frame objects and belongs to a single Game Object such as a Sprite.
func NewAnimation1O(game *Game, parent *Sprite, name string, frameData *FrameData, frames interface{}, frameRate int) *Animation {
    return &Animation{js.Global.Get("Phaser").Get("Animation").New(game, parent, name, frameData, frames, frameRate)}
}
// NewAnimation2O An Animation instance contains a single animation and the controls to play it.
// 
// It is created by the AnimationManager, consists of Animation.Frame objects and belongs to a single Game Object such as a Sprite.
func NewAnimation2O(game *Game, parent *Sprite, name string, frameData *FrameData, frames interface{}, frameRate int, loop bool) *Animation {
    return &Animation{js.Global.Get("Phaser").Get("Animation").New(game, parent, name, frameData, frames, frameRate, loop)}
}
// NewAnimationI An Animation instance contains a single animation and the controls to play it.
// 
// It is created by the AnimationManager, consists of Animation.Frame objects and belongs to a single Game Object such as a Sprite.
func NewAnimationI(args ...interface{}) *Animation {
    return &Animation{js.Global.Get("Phaser").Get("Animation").New(args)}
}



// Game A reference to the currently running Game.
func (self *Animation) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Animation) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Name The user defined name given to this Animation.
func (self *Animation) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA The user defined name given to this Animation.
func (self *Animation) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Delay The delay in ms between each frame of the Animation, based on the given frameRate.
func (self *Animation) Delay() int{
    return self.Object.Get("delay").Int()
}

// SetDelayA The delay in ms between each frame of the Animation, based on the given frameRate.
func (self *Animation) SetDelayA(member int) {
    self.Object.Set("delay", member)
}

// Loop The loop state of the Animation.
func (self *Animation) Loop() bool{
    return self.Object.Get("loop").Bool()
}

// SetLoopA The loop state of the Animation.
func (self *Animation) SetLoopA(member bool) {
    self.Object.Set("loop", member)
}

// LoopCount The number of times the animation has looped since it was last started.
func (self *Animation) LoopCount() int{
    return self.Object.Get("loopCount").Int()
}

// SetLoopCountA The number of times the animation has looped since it was last started.
func (self *Animation) SetLoopCountA(member int) {
    self.Object.Set("loopCount", member)
}

// KillOnComplete Should the parent of this Animation be killed when the animation completes?
func (self *Animation) KillOnComplete() bool{
    return self.Object.Get("killOnComplete").Bool()
}

// SetKillOnCompleteA Should the parent of this Animation be killed when the animation completes?
func (self *Animation) SetKillOnCompleteA(member bool) {
    self.Object.Set("killOnComplete", member)
}

// IsFinished The finished state of the Animation. Set to true once playback completes, false during playback.
func (self *Animation) IsFinished() bool{
    return self.Object.Get("isFinished").Bool()
}

// SetIsFinishedA The finished state of the Animation. Set to true once playback completes, false during playback.
func (self *Animation) SetIsFinishedA(member bool) {
    self.Object.Set("isFinished", member)
}

// IsPlaying The playing state of the Animation. Set to false once playback completes, true during playback.
func (self *Animation) IsPlaying() bool{
    return self.Object.Get("isPlaying").Bool()
}

// SetIsPlayingA The playing state of the Animation. Set to false once playback completes, true during playback.
func (self *Animation) SetIsPlayingA(member bool) {
    self.Object.Set("isPlaying", member)
}

// IsPaused The paused state of the Animation.
func (self *Animation) IsPaused() bool{
    return self.Object.Get("isPaused").Bool()
}

// SetIsPausedA The paused state of the Animation.
func (self *Animation) SetIsPausedA(member bool) {
    self.Object.Set("isPaused", member)
}

// CurrentFrame The currently displayed frame of the Animation.
func (self *Animation) CurrentFrame() *Frame{
    return &Frame{self.Object.Get("currentFrame")}
}

// SetCurrentFrameA The currently displayed frame of the Animation.
func (self *Animation) SetCurrentFrameA(member *Frame) {
    self.Object.Set("currentFrame", member)
}

// OnStart This event is dispatched when this Animation starts playback.
func (self *Animation) OnStart() *Signal{
    return &Signal{self.Object.Get("onStart")}
}

// SetOnStartA This event is dispatched when this Animation starts playback.
func (self *Animation) SetOnStartA(member *Signal) {
    self.Object.Set("onStart", member)
}

// OnUpdate This event is dispatched when the Animation changes frame.
// By default this event is disabled due to its intensive nature. Enable it with: `Animation.enableUpdate = true`.
// Note that the event is only dispatched with the current frame. In a low-FPS environment Animations
// will automatically frame-skip to try and claw back time, so do not base your code on expecting to
// receive a perfectly sequential set of frames from this event.
func (self *Animation) OnUpdate() interface{}{
    return self.Object.Get("onUpdate")
}

// SetOnUpdateA This event is dispatched when the Animation changes frame.
// By default this event is disabled due to its intensive nature. Enable it with: `Animation.enableUpdate = true`.
// Note that the event is only dispatched with the current frame. In a low-FPS environment Animations
// will automatically frame-skip to try and claw back time, so do not base your code on expecting to
// receive a perfectly sequential set of frames from this event.
func (self *Animation) SetOnUpdateA(member interface{}) {
    self.Object.Set("onUpdate", member)
}

// OnComplete This event is dispatched when this Animation completes playback. If the animation is set to loop this is never fired, listen for onLoop instead.
func (self *Animation) OnComplete() *Signal{
    return &Signal{self.Object.Get("onComplete")}
}

// SetOnCompleteA This event is dispatched when this Animation completes playback. If the animation is set to loop this is never fired, listen for onLoop instead.
func (self *Animation) SetOnCompleteA(member *Signal) {
    self.Object.Set("onComplete", member)
}

// OnLoop This event is dispatched when this Animation loops.
func (self *Animation) OnLoop() *Signal{
    return &Signal{self.Object.Get("onLoop")}
}

// SetOnLoopA This event is dispatched when this Animation loops.
func (self *Animation) SetOnLoopA(member *Signal) {
    self.Object.Set("onLoop", member)
}

// IsReversed Indicates if the animation will play backwards.
func (self *Animation) IsReversed() bool{
    return self.Object.Get("isReversed").Bool()
}

// SetIsReversedA Indicates if the animation will play backwards.
func (self *Animation) SetIsReversedA(member bool) {
    self.Object.Set("isReversed", member)
}

// Paused Gets and sets the paused state of this Animation.
func (self *Animation) Paused() bool{
    return self.Object.Get("paused").Bool()
}

// SetPausedA Gets and sets the paused state of this Animation.
func (self *Animation) SetPausedA(member bool) {
    self.Object.Set("paused", member)
}

// Reversed Gets and sets the isReversed state of this Animation.
func (self *Animation) Reversed() bool{
    return self.Object.Get("reversed").Bool()
}

// SetReversedA Gets and sets the isReversed state of this Animation.
func (self *Animation) SetReversedA(member bool) {
    self.Object.Set("reversed", member)
}

// FrameTotal The total number of frames in the currently loaded FrameData, or -1 if no FrameData is loaded.
func (self *Animation) FrameTotal() int{
    return self.Object.Get("frameTotal").Int()
}

// SetFrameTotalA The total number of frames in the currently loaded FrameData, or -1 if no FrameData is loaded.
func (self *Animation) SetFrameTotalA(member int) {
    self.Object.Set("frameTotal", member)
}

// Frame Gets or sets the current frame index and updates the Texture Cache for display.
func (self *Animation) Frame() int{
    return self.Object.Get("frame").Int()
}

// SetFrameA Gets or sets the current frame index and updates the Texture Cache for display.
func (self *Animation) SetFrameA(member int) {
    self.Object.Set("frame", member)
}

// Speed Gets or sets the current speed of the animation in frames per second. Changing this in a playing animation will take effect from the next frame. Minimum value is 1.
func (self *Animation) Speed() int{
    return self.Object.Get("speed").Int()
}

// SetSpeedA Gets or sets the current speed of the animation in frames per second. Changing this in a playing animation will take effect from the next frame. Minimum value is 1.
func (self *Animation) SetSpeedA(member int) {
    self.Object.Set("speed", member)
}

// EnableUpdate Gets or sets if this animation will dispatch the onUpdate events upon changing frame.
func (self *Animation) EnableUpdate() bool{
    return self.Object.Get("enableUpdate").Bool()
}

// SetEnableUpdateA Gets or sets if this animation will dispatch the onUpdate events upon changing frame.
func (self *Animation) SetEnableUpdateA(member bool) {
    self.Object.Set("enableUpdate", member)
}


// Play Plays this animation.
func (self *Animation) Play() *Animation{
    return &Animation{self.Object.Call("play")}
}

// Play1O Plays this animation.
func (self *Animation) Play1O(frameRate int) *Animation{
    return &Animation{self.Object.Call("play", frameRate)}
}

// Play2O Plays this animation.
func (self *Animation) Play2O(frameRate int, loop bool) *Animation{
    return &Animation{self.Object.Call("play", frameRate, loop)}
}

// Play3O Plays this animation.
func (self *Animation) Play3O(frameRate int, loop bool, killOnComplete bool) *Animation{
    return &Animation{self.Object.Call("play", frameRate, loop, killOnComplete)}
}

// PlayI Plays this animation.
func (self *Animation) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("play", args)}
}

// Restart Sets this animation back to the first frame and restarts the animation.
func (self *Animation) Restart() {
    self.Object.Call("restart")
}

// RestartI Sets this animation back to the first frame and restarts the animation.
func (self *Animation) RestartI(args ...interface{}) {
    self.Object.Call("restart", args)
}

// Reverse Reverses the animation direction
func (self *Animation) Reverse() *Animation{
    return &Animation{self.Object.Call("reverse")}
}

// ReverseI Reverses the animation direction
func (self *Animation) ReverseI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("reverse", args)}
}

// ReverseOnce Reverses the animation direction for the current/next animation only
// Once the onComplete event is called this method will be called again and revert
// the reversed state.
func (self *Animation) ReverseOnce() *Animation{
    return &Animation{self.Object.Call("reverseOnce")}
}

// ReverseOnceI Reverses the animation direction for the current/next animation only
// Once the onComplete event is called this method will be called again and revert
// the reversed state.
func (self *Animation) ReverseOnceI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("reverseOnce", args)}
}

// SetFrame Sets this animations playback to a given frame with the given ID.
func (self *Animation) SetFrame() {
    self.Object.Call("setFrame")
}

// SetFrame1O Sets this animations playback to a given frame with the given ID.
func (self *Animation) SetFrame1O(frameId interface{}) {
    self.Object.Call("setFrame", frameId)
}

// SetFrame2O Sets this animations playback to a given frame with the given ID.
func (self *Animation) SetFrame2O(frameId interface{}, useLocalFrameIndex bool) {
    self.Object.Call("setFrame", frameId, useLocalFrameIndex)
}

// SetFrameI Sets this animations playback to a given frame with the given ID.
func (self *Animation) SetFrameI(args ...interface{}) {
    self.Object.Call("setFrame", args)
}

// Stop Stops playback of this animation and set it to a finished state. If a resetFrame is provided it will stop playback and set frame to the first in the animation.
// If `dispatchComplete` is true it will dispatch the complete events, otherwise they'll be ignored.
func (self *Animation) Stop() {
    self.Object.Call("stop")
}

// Stop1O Stops playback of this animation and set it to a finished state. If a resetFrame is provided it will stop playback and set frame to the first in the animation.
// If `dispatchComplete` is true it will dispatch the complete events, otherwise they'll be ignored.
func (self *Animation) Stop1O(resetFrame bool) {
    self.Object.Call("stop", resetFrame)
}

// Stop2O Stops playback of this animation and set it to a finished state. If a resetFrame is provided it will stop playback and set frame to the first in the animation.
// If `dispatchComplete` is true it will dispatch the complete events, otherwise they'll be ignored.
func (self *Animation) Stop2O(resetFrame bool, dispatchComplete bool) {
    self.Object.Call("stop", resetFrame, dispatchComplete)
}

// StopI Stops playback of this animation and set it to a finished state. If a resetFrame is provided it will stop playback and set frame to the first in the animation.
// If `dispatchComplete` is true it will dispatch the complete events, otherwise they'll be ignored.
func (self *Animation) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// OnPause Called when the Game enters a paused state.
func (self *Animation) OnPause() {
    self.Object.Call("onPause")
}

// OnPauseI Called when the Game enters a paused state.
func (self *Animation) OnPauseI(args ...interface{}) {
    self.Object.Call("onPause", args)
}

// OnResume Called when the Game resumes from a paused state.
func (self *Animation) OnResume() {
    self.Object.Call("onResume")
}

// OnResumeI Called when the Game resumes from a paused state.
func (self *Animation) OnResumeI(args ...interface{}) {
    self.Object.Call("onResume", args)
}

// Update Updates this animation. Called automatically by the AnimationManager.
func (self *Animation) Update() {
    self.Object.Call("update")
}

// UpdateI Updates this animation. Called automatically by the AnimationManager.
func (self *Animation) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// UpdateCurrentFrame Changes the currentFrame per the _frameIndex, updates the display state,
// and triggers the update signal.
// 
// Returns true if the current frame update was 'successful', false otherwise.
func (self *Animation) UpdateCurrentFrame(signalUpdate bool, fromPlay bool) bool{
    return self.Object.Call("updateCurrentFrame", signalUpdate, fromPlay).Bool()
}

// UpdateCurrentFrameI Changes the currentFrame per the _frameIndex, updates the display state,
// and triggers the update signal.
// 
// Returns true if the current frame update was 'successful', false otherwise.
func (self *Animation) UpdateCurrentFrameI(args ...interface{}) bool{
    return self.Object.Call("updateCurrentFrame", args).Bool()
}

// Next Advances by the given number of frames in the Animation, taking the loop value into consideration.
func (self *Animation) Next() {
    self.Object.Call("next")
}

// Next1O Advances by the given number of frames in the Animation, taking the loop value into consideration.
func (self *Animation) Next1O(quantity int) {
    self.Object.Call("next", quantity)
}

// NextI Advances by the given number of frames in the Animation, taking the loop value into consideration.
func (self *Animation) NextI(args ...interface{}) {
    self.Object.Call("next", args)
}

// Previous Moves backwards the given number of frames in the Animation, taking the loop value into consideration.
func (self *Animation) Previous() {
    self.Object.Call("previous")
}

// Previous1O Moves backwards the given number of frames in the Animation, taking the loop value into consideration.
func (self *Animation) Previous1O(quantity int) {
    self.Object.Call("previous", quantity)
}

// PreviousI Moves backwards the given number of frames in the Animation, taking the loop value into consideration.
func (self *Animation) PreviousI(args ...interface{}) {
    self.Object.Call("previous", args)
}

// UpdateFrameData Changes the FrameData object this Animation is using.
func (self *Animation) UpdateFrameData(frameData *FrameData) {
    self.Object.Call("updateFrameData", frameData)
}

// UpdateFrameDataI Changes the FrameData object this Animation is using.
func (self *Animation) UpdateFrameDataI(args ...interface{}) {
    self.Object.Call("updateFrameData", args)
}

// Destroy Cleans up this animation ready for deletion. Nulls all values and references.
func (self *Animation) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Cleans up this animation ready for deletion. Nulls all values and references.
func (self *Animation) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

// Complete Called internally when the animation finishes playback.
// Sets the isPlaying and isFinished states and dispatches the onAnimationComplete event if it exists on the parent and local onComplete event.
func (self *Animation) Complete() {
    self.Object.Call("complete")
}

// CompleteI Called internally when the animation finishes playback.
// Sets the isPlaying and isFinished states and dispatches the onAnimationComplete event if it exists on the parent and local onComplete event.
func (self *Animation) CompleteI(args ...interface{}) {
    self.Object.Call("complete", args)
}

// GenerateFrameNames Really handy function for when you are creating arrays of animation data but it's using frame names and not numbers.
// For example imagine you've got 30 frames named: 'explosion_0001-large' to 'explosion_0030-large'
// You could use this function to generate those by doing: Phaser.Animation.generateFrameNames('explosion_', 1, 30, '-large', 4);
func (self *Animation) GenerateFrameNames(prefix string, start int, stop int) []string{
	array00 := self.Object.Call("generateFrameNames", prefix, start, stop)
	length00 := array00.Length()
	out00 := make([]string, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).String()
		
	}
	return out00
}

// GenerateFrameNames1O Really handy function for when you are creating arrays of animation data but it's using frame names and not numbers.
// For example imagine you've got 30 frames named: 'explosion_0001-large' to 'explosion_0030-large'
// You could use this function to generate those by doing: Phaser.Animation.generateFrameNames('explosion_', 1, 30, '-large', 4);
func (self *Animation) GenerateFrameNames1O(prefix string, start int, stop int, suffix string) []string{
	array00 := self.Object.Call("generateFrameNames", prefix, start, stop, suffix)
	length00 := array00.Length()
	out00 := make([]string, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).String()
		
	}
	return out00
}

// GenerateFrameNames2O Really handy function for when you are creating arrays of animation data but it's using frame names and not numbers.
// For example imagine you've got 30 frames named: 'explosion_0001-large' to 'explosion_0030-large'
// You could use this function to generate those by doing: Phaser.Animation.generateFrameNames('explosion_', 1, 30, '-large', 4);
func (self *Animation) GenerateFrameNames2O(prefix string, start int, stop int, suffix string, zeroPad int) []string{
	array00 := self.Object.Call("generateFrameNames", prefix, start, stop, suffix, zeroPad)
	length00 := array00.Length()
	out00 := make([]string, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).String()
		
	}
	return out00
}

// GenerateFrameNamesI Really handy function for when you are creating arrays of animation data but it's using frame names and not numbers.
// For example imagine you've got 30 frames named: 'explosion_0001-large' to 'explosion_0030-large'
// You could use this function to generate those by doing: Phaser.Animation.generateFrameNames('explosion_', 1, 30, '-large', 4);
func (self *Animation) GenerateFrameNamesI(args ...interface{}) []string{
	array00 := self.Object.Call("generateFrameNames", args)
	length00 := array00.Length()
	out00 := make([]string, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).String()
		
	}
	return out00
}

