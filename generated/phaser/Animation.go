// Automatic generation for Phaser.Animation
// generated file Animation.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// An Animation instance contains a single animation and the controls to play it.
// 
// It is created by the AnimationManager, consists of Animation.Frame objects and belongs to a single Game Object such as a Sprite.
type Animation struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *Animation) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Animation) SetGame(member *Game) {
    self.Set("game", member)
}

// The user defined name given to this Animation.
func (self *Animation) GetName() string{
    return self.Get("name").String()
}

// The user defined name given to this Animation.
func (self *Animation) SetName(member string) {
    self.Set("name", member)
}

// The delay in ms between each frame of the Animation, based on the given frameRate.
func (self *Animation) GetDelay() int{
    return self.Get("delay").Int()
}

// The delay in ms between each frame of the Animation, based on the given frameRate.
func (self *Animation) SetDelay(member int) {
    self.Set("delay", member)
}

// The loop state of the Animation.
func (self *Animation) GetLoop() bool{
    return self.Get("loop").Bool()
}

// The loop state of the Animation.
func (self *Animation) SetLoop(member bool) {
    self.Set("loop", member)
}

// The number of times the animation has looped since it was last started.
func (self *Animation) GetLoopCount() int{
    return self.Get("loopCount").Int()
}

// The number of times the animation has looped since it was last started.
func (self *Animation) SetLoopCount(member int) {
    self.Set("loopCount", member)
}

// Should the parent of this Animation be killed when the animation completes?
func (self *Animation) GetKillOnComplete() bool{
    return self.Get("killOnComplete").Bool()
}

// Should the parent of this Animation be killed when the animation completes?
func (self *Animation) SetKillOnComplete(member bool) {
    self.Set("killOnComplete", member)
}

// The finished state of the Animation. Set to true once playback completes, false during playback.
func (self *Animation) GetIsFinished() bool{
    return self.Get("isFinished").Bool()
}

// The finished state of the Animation. Set to true once playback completes, false during playback.
func (self *Animation) SetIsFinished(member bool) {
    self.Set("isFinished", member)
}

// The playing state of the Animation. Set to false once playback completes, true during playback.
func (self *Animation) GetIsPlaying() bool{
    return self.Get("isPlaying").Bool()
}

// The playing state of the Animation. Set to false once playback completes, true during playback.
func (self *Animation) SetIsPlaying(member bool) {
    self.Set("isPlaying", member)
}

// The paused state of the Animation.
func (self *Animation) GetIsPaused() bool{
    return self.Get("isPaused").Bool()
}

// The paused state of the Animation.
func (self *Animation) SetIsPaused(member bool) {
    self.Set("isPaused", member)
}

// The currently displayed frame of the Animation.
func (self *Animation) GetCurrentFrame() *Frame{
    return &Frame{self.Get("currentFrame")}
}

// The currently displayed frame of the Animation.
func (self *Animation) SetCurrentFrame(member *Frame) {
    self.Set("currentFrame", member)
}

// This event is dispatched when this Animation starts playback.
func (self *Animation) GetOnStart() *Signal{
    return &Signal{self.Get("onStart")}
}

// This event is dispatched when this Animation starts playback.
func (self *Animation) SetOnStart(member *Signal) {
    self.Set("onStart", member)
}

// This event is dispatched when the Animation changes frame.
// By default this event is disabled due to its intensive nature. Enable it with: `Animation.enableUpdate = true`.
// Note that the event is only dispatched with the current frame. In a low-FPS environment Animations
// will automatically frame-skip to try and claw back time, so do not base your code on expecting to
// receive a perfectly sequential set of frames from this event.
func (self *Animation) GetOnUpdate() interface{}{
    return self.Get("onUpdate")
}

// This event is dispatched when the Animation changes frame.
// By default this event is disabled due to its intensive nature. Enable it with: `Animation.enableUpdate = true`.
// Note that the event is only dispatched with the current frame. In a low-FPS environment Animations
// will automatically frame-skip to try and claw back time, so do not base your code on expecting to
// receive a perfectly sequential set of frames from this event.
func (self *Animation) SetOnUpdate(member interface{}) {
    self.Set("onUpdate", member)
}

// This event is dispatched when this Animation completes playback. If the animation is set to loop this is never fired, listen for onLoop instead.
func (self *Animation) GetOnComplete() *Signal{
    return &Signal{self.Get("onComplete")}
}

// This event is dispatched when this Animation completes playback. If the animation is set to loop this is never fired, listen for onLoop instead.
func (self *Animation) SetOnComplete(member *Signal) {
    self.Set("onComplete", member)
}

// This event is dispatched when this Animation loops.
func (self *Animation) GetOnLoop() *Signal{
    return &Signal{self.Get("onLoop")}
}

// This event is dispatched when this Animation loops.
func (self *Animation) SetOnLoop(member *Signal) {
    self.Set("onLoop", member)
}

// Indicates if the animation will play backwards.
func (self *Animation) GetIsReversed() bool{
    return self.Get("isReversed").Bool()
}

// Indicates if the animation will play backwards.
func (self *Animation) SetIsReversed(member bool) {
    self.Set("isReversed", member)
}

// Gets and sets the paused state of this Animation.
func (self *Animation) GetPaused() bool{
    return self.Get("paused").Bool()
}

// Gets and sets the paused state of this Animation.
func (self *Animation) SetPaused(member bool) {
    self.Set("paused", member)
}

// Gets and sets the isReversed state of this Animation.
func (self *Animation) GetReversed() bool{
    return self.Get("reversed").Bool()
}

// Gets and sets the isReversed state of this Animation.
func (self *Animation) SetReversed(member bool) {
    self.Set("reversed", member)
}

// The total number of frames in the currently loaded FrameData, or -1 if no FrameData is loaded.
func (self *Animation) GetFrameTotal() int{
    return self.Get("frameTotal").Int()
}

// The total number of frames in the currently loaded FrameData, or -1 if no FrameData is loaded.
func (self *Animation) SetFrameTotal(member int) {
    self.Set("frameTotal", member)
}

// Gets or sets the current frame index and updates the Texture Cache for display.
func (self *Animation) GetFrame() int{
    return self.Get("frame").Int()
}

// Gets or sets the current frame index and updates the Texture Cache for display.
func (self *Animation) SetFrame(member int) {
    self.Set("frame", member)
}

// Gets or sets the current speed of the animation in frames per second. Changing this in a playing animation will take effect from the next frame. Minimum value is 1.
func (self *Animation) GetSpeed() int{
    return self.Get("speed").Int()
}

// Gets or sets the current speed of the animation in frames per second. Changing this in a playing animation will take effect from the next frame. Minimum value is 1.
func (self *Animation) SetSpeed(member int) {
    self.Set("speed", member)
}

// Gets or sets if this animation will dispatch the onUpdate events upon changing frame.
func (self *Animation) GetEnableUpdate() bool{
    return self.Get("enableUpdate").Bool()
}

// Gets or sets if this animation will dispatch the onUpdate events upon changing frame.
func (self *Animation) SetEnableUpdate(member bool) {
    self.Set("enableUpdate", member)
}



// Plays this animation.
func (self *Animation) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Call("play", args)}
}

// Sets this animation back to the first frame and restarts the animation.
func (self *Animation) RestartI(args ...interface{}) {
    self.Call("restart", args)
}

// Reverses the animation direction
func (self *Animation) ReverseI(args ...interface{}) *Animation{
    return &Animation{self.Call("reverse", args)}
}

// Reverses the animation direction for the current/next animation only
// Once the onComplete event is called this method will be called again and revert
// the reversed state.
func (self *Animation) ReverseOnceI(args ...interface{}) *Animation{
    return &Animation{self.Call("reverseOnce", args)}
}

// Sets this animations playback to a given frame with the given ID.
func (self *Animation) SetFrameI(args ...interface{}) {
    self.Call("setFrame", args)
}

// Stops playback of this animation and set it to a finished state. If a resetFrame is provided it will stop playback and set frame to the first in the animation.
// If `dispatchComplete` is true it will dispatch the complete events, otherwise they'll be ignored.
func (self *Animation) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// Called when the Game enters a paused state.
func (self *Animation) OnPauseI(args ...interface{}) {
    self.Call("onPause", args)
}

// Called when the Game resumes from a paused state.
func (self *Animation) OnResumeI(args ...interface{}) {
    self.Call("onResume", args)
}

// Updates this animation. Called automatically by the AnimationManager.
func (self *Animation) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Changes the currentFrame per the _frameIndex, updates the display state,
// and triggers the update signal.
// 
// Returns true if the current frame update was 'successful', false otherwise.
func (self *Animation) UpdateCurrentFrameI(args ...interface{}) bool{
    return self.Call("updateCurrentFrame", args).Bool()
}

// Advances by the given number of frames in the Animation, taking the loop value into consideration.
func (self *Animation) NextI(args ...interface{}) {
    self.Call("next", args)
}

// Moves backwards the given number of frames in the Animation, taking the loop value into consideration.
func (self *Animation) PreviousI(args ...interface{}) {
    self.Call("previous", args)
}

// Changes the FrameData object this Animation is using.
func (self *Animation) UpdateFrameDataI(args ...interface{}) {
    self.Call("updateFrameData", args)
}

// Cleans up this animation ready for deletion. Nulls all values and references.
func (self *Animation) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}

// Called internally when the animation finishes playback.
// Sets the isPlaying and isFinished states and dispatches the onAnimationComplete event if it exists on the parent and local onComplete event.
func (self *Animation) CompleteI(args ...interface{}) {
    self.Call("complete", args)
}

// Really handy function for when you are creating arrays of animation data but it's using frame names and not numbers.
// For example imagine you've got 30 frames named: 'explosion_0001-large' to 'explosion_0030-large'
// You could use this function to generate those by doing: Phaser.Animation.generateFrameNames('explosion_', 1, 30, '-large', 4);
func (self *Animation) GenerateFrameNamesI(args ...interface{}) []string{
	array := self.Call("generateFrameNames", args)
	length := array.Length()
	out := make([]string, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).String()
		
	}
	return out
}
