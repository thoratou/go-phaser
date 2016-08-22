// Automatic generation for Phaser.Tween
// generated file Tween.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A Tween allows you to alter one or more properties of a target object over a defined period of time.
// This can be used for things such as alpha fading Sprites, scaling them or motion.
// Use `Tween.to` or `Tween.from` to set-up the tween values. You can create multiple tweens on the same object
// by calling Tween.to multiple times on the same Tween. Additional tweens specified in this way become "child" tweens and
// are played through in sequence. You can use Tween.timeScale and Tween.reverse to control the playback of this Tween and all of its children.
type Tween struct {
    *js.Object
}


// A reference to the currently running Game.
func (self *Tween) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *Tween) SetGame(member *Game) {
    self.Set("game", member)
}

// The target object, such as a Phaser.Sprite or property like Phaser.Sprite.scale.
func (self *Tween) GetTarget() interface{}{
    return self.Get("target")
}

// The target object, such as a Phaser.Sprite or property like Phaser.Sprite.scale.
func (self *Tween) SetTarget(member interface{}) {
    self.Set("target", member)
}

// Reference to the TweenManager responsible for updating this Tween.
func (self *Tween) GetManager() *TweenManager{
    return &TweenManager{self.Get("manager")}
}

// Reference to the TweenManager responsible for updating this Tween.
func (self *Tween) SetManager(member *TweenManager) {
    self.Set("manager", member)
}

// An Array of TweenData objects that comprise the different parts of this Tween.
func (self *Tween) GetTimeline() []interface{}{
	array := self.Get("timeline")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// An Array of TweenData objects that comprise the different parts of this Tween.
func (self *Tween) SetTimeline(member []interface{}) {
    self.Set("timeline", member)
}

// If set to `true` the current tween will play in reverse.
// If the tween hasn't yet started this has no effect.
// If there are child tweens then all child tweens will play in reverse from the current point.
func (self *Tween) GetReverse() bool{
    return self.Get("reverse").Bool()
}

// If set to `true` the current tween will play in reverse.
// If the tween hasn't yet started this has no effect.
// If there are child tweens then all child tweens will play in reverse from the current point.
func (self *Tween) SetReverse(member bool) {
    self.Set("reverse", member)
}

// The speed at which the tweens will run. A value of 1 means it will match the game frame rate. 0.5 will run at half the frame rate. 2 at double the frame rate, etc.
// If a tweens duration is 1 second but timeScale is 0.5 then it will take 2 seconds to complete.
func (self *Tween) GetTimeScale() int{
    return self.Get("timeScale").Int()
}

// The speed at which the tweens will run. A value of 1 means it will match the game frame rate. 0.5 will run at half the frame rate. 2 at double the frame rate, etc.
// If a tweens duration is 1 second but timeScale is 0.5 then it will take 2 seconds to complete.
func (self *Tween) SetTimeScale(member int) {
    self.Set("timeScale", member)
}

// If the Tween and any child tweens are set to repeat this contains the current repeat count.
func (self *Tween) GetRepeatCounter() int{
    return self.Get("repeatCounter").Int()
}

// If the Tween and any child tweens are set to repeat this contains the current repeat count.
func (self *Tween) SetRepeatCounter(member int) {
    self.Set("repeatCounter", member)
}

// True if this Tween is ready to be deleted by the TweenManager.
func (self *Tween) GetPendingDelete() bool{
    return self.Get("pendingDelete").Bool()
}

// True if this Tween is ready to be deleted by the TweenManager.
func (self *Tween) SetPendingDelete(member bool) {
    self.Set("pendingDelete", member)
}

// The onStart event is fired when the Tween begins. If there is a delay before the tween starts then onStart fires after the delay is finished.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) GetOnStart() *Signal{
    return &Signal{self.Get("onStart")}
}

// The onStart event is fired when the Tween begins. If there is a delay before the tween starts then onStart fires after the delay is finished.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) SetOnStart(member *Signal) {
    self.Set("onStart", member)
}

// The onLoop event is fired if the Tween, or any child tweens loop.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) GetOnLoop() *Signal{
    return &Signal{self.Get("onLoop")}
}

// The onLoop event is fired if the Tween, or any child tweens loop.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) SetOnLoop(member *Signal) {
    self.Set("onLoop", member)
}

// The onRepeat event is fired if the Tween and all of its children repeats. If this tween has no children this will never be fired.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) GetOnRepeat() *Signal{
    return &Signal{self.Get("onRepeat")}
}

// The onRepeat event is fired if the Tween and all of its children repeats. If this tween has no children this will never be fired.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) SetOnRepeat(member *Signal) {
    self.Set("onRepeat", member)
}

// The onChildComplete event is fired when the Tween or any of its children completes.
// Fires every time a child completes unless a child is set to repeat forever.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) GetOnChildComplete() *Signal{
    return &Signal{self.Get("onChildComplete")}
}

// The onChildComplete event is fired when the Tween or any of its children completes.
// Fires every time a child completes unless a child is set to repeat forever.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) SetOnChildComplete(member *Signal) {
    self.Set("onChildComplete", member)
}

// The onComplete event is fired when the Tween and all of its children completes. Does not fire if the Tween is set to loop or repeatAll(-1).
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) GetOnComplete() *Signal{
    return &Signal{self.Get("onComplete")}
}

// The onComplete event is fired when the Tween and all of its children completes. Does not fire if the Tween is set to loop or repeatAll(-1).
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) SetOnComplete(member *Signal) {
    self.Set("onComplete", member)
}

// If the tween is running this is set to true, otherwise false. Tweens that are in a delayed state or waiting to start are considered as being running.
func (self *Tween) GetIsRunning() bool{
    return self.Get("isRunning").Bool()
}

// If the tween is running this is set to true, otherwise false. Tweens that are in a delayed state or waiting to start are considered as being running.
func (self *Tween) SetIsRunning(member bool) {
    self.Set("isRunning", member)
}

// The current Tween child being run.
func (self *Tween) GetCurrent() int{
    return self.Get("current").Int()
}

// The current Tween child being run.
func (self *Tween) SetCurrent(member int) {
    self.Set("current", member)
}

// Target property cache used when building the child data values.
func (self *Tween) GetProperties() interface{}{
    return self.Get("properties")
}

// Target property cache used when building the child data values.
func (self *Tween) SetProperties(member interface{}) {
    self.Set("properties", member)
}

// If this Tween is chained to another this holds a reference to it.
func (self *Tween) GetChainedTween() *Tween{
    return &Tween{self.Get("chainedTween")}
}

// If this Tween is chained to another this holds a reference to it.
func (self *Tween) SetChainedTween(member *Tween) {
    self.Set("chainedTween", member)
}

// Is this Tween paused or not?
func (self *Tween) GetIsPaused() bool{
    return self.Get("isPaused").Bool()
}

// Is this Tween paused or not?
func (self *Tween) SetIsPaused(member bool) {
    self.Set("isPaused", member)
}

// Is this Tween frame or time based? A frame based tween will use the physics elapsed timer when updating. This means
// it will retain the same consistent frame rate, regardless of the speed of the device. The duration value given should
// be given in frames.
// 
// If the Tween uses a time based update (which is the default) then the duration is given in milliseconds.
// In this situation a 2000ms tween will last exactly 2 seconds, regardless of the device and how many visual updates the tween
// has actually been through. For very short tweens you may wish to experiment with a frame based update instead.
// 
// The default value is whatever you've set in TweenManager.frameBased.
func (self *Tween) GetFrameBased() bool{
    return self.Get("frameBased").Bool()
}

// Is this Tween frame or time based? A frame based tween will use the physics elapsed timer when updating. This means
// it will retain the same consistent frame rate, regardless of the speed of the device. The duration value given should
// be given in frames.
// 
// If the Tween uses a time based update (which is the default) then the duration is given in milliseconds.
// In this situation a 2000ms tween will last exactly 2 seconds, regardless of the device and how many visual updates the tween
// has actually been through. For very short tweens you may wish to experiment with a frame based update instead.
// 
// The default value is whatever you've set in TweenManager.frameBased.
func (self *Tween) SetFrameBased(member bool) {
    self.Set("frameBased", member)
}

// Gets the total duration of this Tween, including all child tweens, in milliseconds.
func (self *Tween) GetTotalDuration() *TweenData{
    return &TweenData{self.Get("totalDuration")}
}

// Gets the total duration of this Tween, including all child tweens, in milliseconds.
func (self *Tween) SetTotalDuration(member *TweenData) {
    self.Set("totalDuration", member)
}



// Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) ToI(args ...interface{}) *Tween{
    return &Tween{self.Call("to", args)}
}

// Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) FromI(args ...interface{}) *Tween{
    return &Tween{self.Call("from", args)}
}

// Starts the tween running. Can also be called by the autoStart parameter of `Tween.to` or `Tween.from`.
// This sets the `Tween.isRunning` property to `true` and dispatches a `Tween.onStart` signal.
// If the Tween has a delay set then nothing will start tweening until the delay has expired.
func (self *Tween) StartI(args ...interface{}) *Tween{
    return &Tween{self.Call("start", args)}
}

// Stops the tween if running and flags it for deletion from the TweenManager.
// If called directly the `Tween.onComplete` signal is not dispatched and no chained tweens are started unless the complete parameter is set to `true`.
// If you just wish to pause a tween then use Tween.pause instead.
func (self *Tween) StopI(args ...interface{}) *Tween{
    return &Tween{self.Call("stop", args)}
}

// Updates either a single TweenData or all TweenData objects properties to the given value.
// Used internally by methods like Tween.delay, Tween.yoyo, etc. but can also be called directly if you know which property you want to tweak.
// The property is not checked, so if you pass an invalid one you'll generate a run-time error.
func (self *Tween) UpdateTweenDataI(args ...interface{}) *Tween{
    return &Tween{self.Call("updateTweenData", args)}
}

// Sets the delay in milliseconds before this tween will start. If there are child tweens it sets the delay before the first child starts.
// The delay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to delay.
// If you have child tweens and pass -1 as the index value it sets the delay across all of them.
func (self *Tween) DelayI(args ...interface{}) *Tween{
    return &Tween{self.Call("delay", args)}
}

// Sets the number of times this tween will repeat.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to repeat.
// If you have child tweens and pass -1 as the index value it sets the number of times they'll repeat across all of them.
// If you wish to define how many times this Tween and all children will repeat see Tween.repeatAll.
func (self *Tween) RepeatI(args ...interface{}) *Tween{
    return &Tween{self.Call("repeat", args)}
}

// Sets the delay in milliseconds before this tween will repeat itself.
// The repeatDelay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to set repeatDelay on.
// If you have child tweens and pass -1 as the index value it sets the repeatDelay across all of them.
func (self *Tween) RepeatDelayI(args ...interface{}) *Tween{
    return &Tween{self.Call("repeatDelay", args)}
}

// A Tween that has yoyo set to true will run through from its starting values to its end values and then play back in reverse from end to start.
// Used in combination with repeat you can create endless loops.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to yoyo.
// If you have child tweens and pass -1 as the index value it sets the yoyo property across all of them.
// If you wish to yoyo this Tween and all of its children then see Tween.yoyoAll.
func (self *Tween) YoyoI(args ...interface{}) *Tween{
    return &Tween{self.Call("yoyo", args)}
}

// Sets the delay in milliseconds before this tween will run a yoyo (only applies if yoyo is enabled).
// The repeatDelay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to set repeatDelay on.
// If you have child tweens and pass -1 as the index value it sets the repeatDelay across all of them.
func (self *Tween) YoyoDelayI(args ...interface{}) *Tween{
    return &Tween{self.Call("yoyoDelay", args)}
}

// Set easing function this tween will use, i.e. Phaser.Easing.Linear.None.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
// If you have child tweens and pass -1 as the index value it sets the easing function defined here across all of them.
func (self *Tween) EasingI(args ...interface{}) *Tween{
    return &Tween{self.Call("easing", args)}
}

// Sets the interpolation function the tween will use. By default it uses Phaser.Math.linearInterpolation.
// Also available: Phaser.Math.bezierInterpolation and Phaser.Math.catmullRomInterpolation.
// The interpolation function is only used if the target properties is an array.
// If you have child tweens and pass -1 as the index value and it will set the interpolation function across all of them.
func (self *Tween) InterpolationI(args ...interface{}) *Tween{
    return &Tween{self.Call("interpolation", args)}
}

// Set how many times this tween and all of its children will repeat.
// A tween (A) with 3 children (B,C,D) with a `repeatAll` value of 2 would play as: ABCDABCD before completing.
func (self *Tween) RepeatAllI(args ...interface{}) *Tween{
    return &Tween{self.Call("repeatAll", args)}
}

// This method allows you to chain tweens together. Any tween chained to this tween will have its `Tween.start` method called
// as soon as this tween completes. If this tween never completes (i.e. repeatAll or loop is set) then the chain will never progress.
// Note that `Tween.onComplete` will fire when *this* tween completes, not when the whole chain completes.
// For that you should listen to `onComplete` on the final tween in your chain.
// 
// If you pass multiple tweens to this method they will be joined into a single long chain.
// For example if this is Tween A and you pass in B, C and D then B will be chained to A, C will be chained to B and D will be chained to C.
// Any previously chained tweens that may have been set will be overwritten.
func (self *Tween) ChainI(args ...interface{}) *Tween{
    return &Tween{self.Call("chain", args)}
}

// Enables the looping of this tween. The tween will loop forever, and onComplete will never fire.
// 
// If `value` is `true` then this is the same as setting `Tween.repeatAll(-1)`.
// If `value` is `false` it is the same as setting `Tween.repeatAll(0)` and will reset the `repeatCounter` to zero.
// 
// Usage:
// game.add.tween(p).to({ x: 700 }, 1000, Phaser.Easing.Linear.None, true)
// .to({ y: 300 }, 1000, Phaser.Easing.Linear.None)
// .to({ x: 0 }, 1000, Phaser.Easing.Linear.None)
// .to({ y: 0 }, 1000, Phaser.Easing.Linear.None)
// .loop();
func (self *Tween) LoopI(args ...interface{}) *Tween{
    return &Tween{self.Call("loop", args)}
}

// Sets a callback to be fired each time this tween updates.
func (self *Tween) OnUpdateCallbackI(args ...interface{}) *Tween{
    return &Tween{self.Call("onUpdateCallback", args)}
}

// Pauses the tween. Resume playback with Tween.resume.
func (self *Tween) PauseI(args ...interface{}) {
    self.Call("pause", args)
}

// This is called by the core Game loop. Do not call it directly, instead use Tween.pause.
func (self *Tween) _pauseI(args ...interface{}) {
    self.Call("_pause", args)
}

// Resumes a paused tween.
func (self *Tween) ResumeI(args ...interface{}) {
    self.Call("resume", args)
}

// This is called by the core Game loop. Do not call it directly, instead use Tween.pause.
func (self *Tween) _resumeI(args ...interface{}) {
    self.Call("_resume", args)
}

// Core tween update function called by the TweenManager. Does not need to be invoked directly.
func (self *Tween) UpdateI(args ...interface{}) bool{
    return self.Call("update", args).Bool()
}

// This will generate an array populated with the tweened object values from start to end.
// It works by running the tween simulation at the given frame rate based on the values set-up in Tween.to and Tween.from.
// It ignores delay and repeat counts and any chained tweens, but does include child tweens.
// Just one play through of the tween data is returned, including yoyo if set.
func (self *Tween) GenerateDataI(args ...interface{}) []interface{}{
	array := self.Call("generateData", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}
