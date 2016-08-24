// Package phaser Automatic generation for Phaser.Tween
// generated file Tween.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Tween A Tween allows you to alter one or more properties of a target object over a defined period of time.
// This can be used for things such as alpha fading Sprites, scaling them or motion.
// Use `Tween.to` or `Tween.from` to set-up the tween values. You can create multiple tweens on the same object
// by calling Tween.to multiple times on the same Tween. Additional tweens specified in this way become "child" tweens and
// are played through in sequence. You can use Tween.timeScale and Tween.reverse to control the playback of this Tween and all of its children.
type Tween struct {
    *js.Object
}

// NewTween A Tween allows you to alter one or more properties of a target object over a defined period of time.
// This can be used for things such as alpha fading Sprites, scaling them or motion.
// Use `Tween.to` or `Tween.from` to set-up the tween values. You can create multiple tweens on the same object
// by calling Tween.to multiple times on the same Tween. Additional tweens specified in this way become "child" tweens and
// are played through in sequence. You can use Tween.timeScale and Tween.reverse to control the playback of this Tween and all of its children.
func NewTween(target interface{}, game *Game, manager *TweenManager) *Tween {
    return &Tween{js.Global.Get("Phaser").Get("Tween").New(target, game, manager)}
}
// NewTweenI A Tween allows you to alter one or more properties of a target object over a defined period of time.
// This can be used for things such as alpha fading Sprites, scaling them or motion.
// Use `Tween.to` or `Tween.from` to set-up the tween values. You can create multiple tweens on the same object
// by calling Tween.to multiple times on the same Tween. Additional tweens specified in this way become "child" tweens and
// are played through in sequence. You can use Tween.timeScale and Tween.reverse to control the playback of this Tween and all of its children.
func NewTweenI(args ...interface{}) *Tween {
    return &Tween{js.Global.Get("Phaser").Get("Tween").New(args)}
}



// Game A reference to the currently running Game.
func (self *Tween) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *Tween) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Target The target object, such as a Phaser.Sprite or property like Phaser.Sprite.scale.
func (self *Tween) Target() interface{}{
    return self.Object.Get("target")
}

// SetTargetA The target object, such as a Phaser.Sprite or property like Phaser.Sprite.scale.
func (self *Tween) SetTargetA(member interface{}) {
    self.Object.Set("target", member)
}

// Manager Reference to the TweenManager responsible for updating this Tween.
func (self *Tween) Manager() *TweenManager{
    return &TweenManager{self.Object.Get("manager")}
}

// SetManagerA Reference to the TweenManager responsible for updating this Tween.
func (self *Tween) SetManagerA(member *TweenManager) {
    self.Object.Set("manager", member)
}

// Timeline An Array of TweenData objects that comprise the different parts of this Tween.
func (self *Tween) Timeline() []interface{}{
	array00 := self.Object.Get("timeline")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetTimelineA An Array of TweenData objects that comprise the different parts of this Tween.
func (self *Tween) SetTimelineA(member []interface{}) {
    self.Object.Set("timeline", member)
}

// Reverse If set to `true` the current tween will play in reverse.
// If the tween hasn't yet started this has no effect.
// If there are child tweens then all child tweens will play in reverse from the current point.
func (self *Tween) Reverse() bool{
    return self.Object.Get("reverse").Bool()
}

// SetReverseA If set to `true` the current tween will play in reverse.
// If the tween hasn't yet started this has no effect.
// If there are child tweens then all child tweens will play in reverse from the current point.
func (self *Tween) SetReverseA(member bool) {
    self.Object.Set("reverse", member)
}

// TimeScale The speed at which the tweens will run. A value of 1 means it will match the game frame rate. 0.5 will run at half the frame rate. 2 at double the frame rate, etc.
// If a tweens duration is 1 second but timeScale is 0.5 then it will take 2 seconds to complete.
func (self *Tween) TimeScale() int{
    return self.Object.Get("timeScale").Int()
}

// SetTimeScaleA The speed at which the tweens will run. A value of 1 means it will match the game frame rate. 0.5 will run at half the frame rate. 2 at double the frame rate, etc.
// If a tweens duration is 1 second but timeScale is 0.5 then it will take 2 seconds to complete.
func (self *Tween) SetTimeScaleA(member int) {
    self.Object.Set("timeScale", member)
}

// RepeatCounter If the Tween and any child tweens are set to repeat this contains the current repeat count.
func (self *Tween) RepeatCounter() int{
    return self.Object.Get("repeatCounter").Int()
}

// SetRepeatCounterA If the Tween and any child tweens are set to repeat this contains the current repeat count.
func (self *Tween) SetRepeatCounterA(member int) {
    self.Object.Set("repeatCounter", member)
}

// PendingDelete True if this Tween is ready to be deleted by the TweenManager.
func (self *Tween) PendingDelete() bool{
    return self.Object.Get("pendingDelete").Bool()
}

// SetPendingDeleteA True if this Tween is ready to be deleted by the TweenManager.
func (self *Tween) SetPendingDeleteA(member bool) {
    self.Object.Set("pendingDelete", member)
}

// OnStart The onStart event is fired when the Tween begins. If there is a delay before the tween starts then onStart fires after the delay is finished.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) OnStart() *Signal{
    return &Signal{self.Object.Get("onStart")}
}

// SetOnStartA The onStart event is fired when the Tween begins. If there is a delay before the tween starts then onStart fires after the delay is finished.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) SetOnStartA(member *Signal) {
    self.Object.Set("onStart", member)
}

// OnLoop The onLoop event is fired if the Tween, or any child tweens loop.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) OnLoop() *Signal{
    return &Signal{self.Object.Get("onLoop")}
}

// SetOnLoopA The onLoop event is fired if the Tween, or any child tweens loop.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) SetOnLoopA(member *Signal) {
    self.Object.Set("onLoop", member)
}

// OnRepeat The onRepeat event is fired if the Tween and all of its children repeats. If this tween has no children this will never be fired.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) OnRepeat() *Signal{
    return &Signal{self.Object.Get("onRepeat")}
}

// SetOnRepeatA The onRepeat event is fired if the Tween and all of its children repeats. If this tween has no children this will never be fired.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) SetOnRepeatA(member *Signal) {
    self.Object.Set("onRepeat", member)
}

// OnChildComplete The onChildComplete event is fired when the Tween or any of its children completes.
// Fires every time a child completes unless a child is set to repeat forever.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) OnChildComplete() *Signal{
    return &Signal{self.Object.Get("onChildComplete")}
}

// SetOnChildCompleteA The onChildComplete event is fired when the Tween or any of its children completes.
// Fires every time a child completes unless a child is set to repeat forever.
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) SetOnChildCompleteA(member *Signal) {
    self.Object.Set("onChildComplete", member)
}

// OnComplete The onComplete event is fired when the Tween and all of its children completes. Does not fire if the Tween is set to loop or repeatAll(-1).
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) OnComplete() *Signal{
    return &Signal{self.Object.Get("onComplete")}
}

// SetOnCompleteA The onComplete event is fired when the Tween and all of its children completes. Does not fire if the Tween is set to loop or repeatAll(-1).
// It will be sent 2 parameters: the target object and this tween.
func (self *Tween) SetOnCompleteA(member *Signal) {
    self.Object.Set("onComplete", member)
}

// IsRunning If the tween is running this is set to true, otherwise false. Tweens that are in a delayed state or waiting to start are considered as being running.
func (self *Tween) IsRunning() bool{
    return self.Object.Get("isRunning").Bool()
}

// SetIsRunningA If the tween is running this is set to true, otherwise false. Tweens that are in a delayed state or waiting to start are considered as being running.
func (self *Tween) SetIsRunningA(member bool) {
    self.Object.Set("isRunning", member)
}

// Current The current Tween child being run.
func (self *Tween) Current() int{
    return self.Object.Get("current").Int()
}

// SetCurrentA The current Tween child being run.
func (self *Tween) SetCurrentA(member int) {
    self.Object.Set("current", member)
}

// Properties Target property cache used when building the child data values.
func (self *Tween) Properties() interface{}{
    return self.Object.Get("properties")
}

// SetPropertiesA Target property cache used when building the child data values.
func (self *Tween) SetPropertiesA(member interface{}) {
    self.Object.Set("properties", member)
}

// ChainedTween If this Tween is chained to another this holds a reference to it.
func (self *Tween) ChainedTween() *Tween{
    return &Tween{self.Object.Get("chainedTween")}
}

// SetChainedTweenA If this Tween is chained to another this holds a reference to it.
func (self *Tween) SetChainedTweenA(member *Tween) {
    self.Object.Set("chainedTween", member)
}

// IsPaused Is this Tween paused or not?
func (self *Tween) IsPaused() bool{
    return self.Object.Get("isPaused").Bool()
}

// SetIsPausedA Is this Tween paused or not?
func (self *Tween) SetIsPausedA(member bool) {
    self.Object.Set("isPaused", member)
}

// FrameBased Is this Tween frame or time based? A frame based tween will use the physics elapsed timer when updating. This means
// it will retain the same consistent frame rate, regardless of the speed of the device. The duration value given should
// be given in frames.
// 
// If the Tween uses a time based update (which is the default) then the duration is given in milliseconds.
// In this situation a 2000ms tween will last exactly 2 seconds, regardless of the device and how many visual updates the tween
// has actually been through. For very short tweens you may wish to experiment with a frame based update instead.
// 
// The default value is whatever you've set in TweenManager.frameBased.
func (self *Tween) FrameBased() bool{
    return self.Object.Get("frameBased").Bool()
}

// SetFrameBasedA Is this Tween frame or time based? A frame based tween will use the physics elapsed timer when updating. This means
// it will retain the same consistent frame rate, regardless of the speed of the device. The duration value given should
// be given in frames.
// 
// If the Tween uses a time based update (which is the default) then the duration is given in milliseconds.
// In this situation a 2000ms tween will last exactly 2 seconds, regardless of the device and how many visual updates the tween
// has actually been through. For very short tweens you may wish to experiment with a frame based update instead.
// 
// The default value is whatever you've set in TweenManager.frameBased.
func (self *Tween) SetFrameBasedA(member bool) {
    self.Object.Set("frameBased", member)
}

// TotalDuration Gets the total duration of this Tween, including all child tweens, in milliseconds.
func (self *Tween) TotalDuration() *TweenData{
    return &TweenData{self.Object.Get("totalDuration")}
}

// SetTotalDurationA Gets the total duration of this Tween, including all child tweens, in milliseconds.
func (self *Tween) SetTotalDurationA(member *TweenData) {
    self.Object.Set("totalDuration", member)
}


// To Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) To(properties interface{}) *Tween{
    return &Tween{self.Object.Call("to", properties)}
}

// To1O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) To1O(properties interface{}, duration int) *Tween{
    return &Tween{self.Object.Call("to", properties, duration)}
}

// To2O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) To2O(properties interface{}, duration int, ease interface{}) *Tween{
    return &Tween{self.Object.Call("to", properties, duration, ease)}
}

// To3O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) To3O(properties interface{}, duration int, ease interface{}, autoStart bool) *Tween{
    return &Tween{self.Object.Call("to", properties, duration, ease, autoStart)}
}

// To4O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) To4O(properties interface{}, duration int, ease interface{}, autoStart bool, delay int) *Tween{
    return &Tween{self.Object.Call("to", properties, duration, ease, autoStart, delay)}
}

// To5O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) To5O(properties interface{}, duration int, ease interface{}, autoStart bool, delay int, repeat int) *Tween{
    return &Tween{self.Object.Call("to", properties, duration, ease, autoStart, delay, repeat)}
}

// To6O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) To6O(properties interface{}, duration int, ease interface{}, autoStart bool, delay int, repeat int, yoyo bool) *Tween{
    return &Tween{self.Object.Call("to", properties, duration, ease, autoStart, delay, repeat, yoyo)}
}

// ToI Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) ToI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("to", args)}
}

// From Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) From(properties interface{}) *Tween{
    return &Tween{self.Object.Call("from", properties)}
}

// From1O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) From1O(properties interface{}, duration int) *Tween{
    return &Tween{self.Object.Call("from", properties, duration)}
}

// From2O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) From2O(properties interface{}, duration int, ease interface{}) *Tween{
    return &Tween{self.Object.Call("from", properties, duration, ease)}
}

// From3O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) From3O(properties interface{}, duration int, ease interface{}, autoStart bool) *Tween{
    return &Tween{self.Object.Call("from", properties, duration, ease, autoStart)}
}

// From4O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) From4O(properties interface{}, duration int, ease interface{}, autoStart bool, delay int) *Tween{
    return &Tween{self.Object.Call("from", properties, duration, ease, autoStart, delay)}
}

// From5O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) From5O(properties interface{}, duration int, ease interface{}, autoStart bool, delay int, repeat int) *Tween{
    return &Tween{self.Object.Call("from", properties, duration, ease, autoStart, delay, repeat)}
}

// From6O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) From6O(properties interface{}, duration int, ease interface{}, autoStart bool, delay int, repeat int, yoyo bool) *Tween{
    return &Tween{self.Object.Call("from", properties, duration, ease, autoStart, delay, repeat, yoyo)}
}

// FromI Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
func (self *Tween) FromI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("from", args)}
}

// Start Starts the tween running. Can also be called by the autoStart parameter of `Tween.to` or `Tween.from`.
// This sets the `Tween.isRunning` property to `true` and dispatches a `Tween.onStart` signal.
// If the Tween has a delay set then nothing will start tweening until the delay has expired.
func (self *Tween) Start() *Tween{
    return &Tween{self.Object.Call("start")}
}

// Start1O Starts the tween running. Can also be called by the autoStart parameter of `Tween.to` or `Tween.from`.
// This sets the `Tween.isRunning` property to `true` and dispatches a `Tween.onStart` signal.
// If the Tween has a delay set then nothing will start tweening until the delay has expired.
func (self *Tween) Start1O(index int) *Tween{
    return &Tween{self.Object.Call("start", index)}
}

// StartI Starts the tween running. Can also be called by the autoStart parameter of `Tween.to` or `Tween.from`.
// This sets the `Tween.isRunning` property to `true` and dispatches a `Tween.onStart` signal.
// If the Tween has a delay set then nothing will start tweening until the delay has expired.
func (self *Tween) StartI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("start", args)}
}

// Stop Stops the tween if running and flags it for deletion from the TweenManager.
// If called directly the `Tween.onComplete` signal is not dispatched and no chained tweens are started unless the complete parameter is set to `true`.
// If you just wish to pause a tween then use Tween.pause instead.
func (self *Tween) Stop() *Tween{
    return &Tween{self.Object.Call("stop")}
}

// Stop1O Stops the tween if running and flags it for deletion from the TweenManager.
// If called directly the `Tween.onComplete` signal is not dispatched and no chained tweens are started unless the complete parameter is set to `true`.
// If you just wish to pause a tween then use Tween.pause instead.
func (self *Tween) Stop1O(complete bool) *Tween{
    return &Tween{self.Object.Call("stop", complete)}
}

// StopI Stops the tween if running and flags it for deletion from the TweenManager.
// If called directly the `Tween.onComplete` signal is not dispatched and no chained tweens are started unless the complete parameter is set to `true`.
// If you just wish to pause a tween then use Tween.pause instead.
func (self *Tween) StopI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("stop", args)}
}

// UpdateTweenData Updates either a single TweenData or all TweenData objects properties to the given value.
// Used internally by methods like Tween.delay, Tween.yoyo, etc. but can also be called directly if you know which property you want to tweak.
// The property is not checked, so if you pass an invalid one you'll generate a run-time error.
func (self *Tween) UpdateTweenData(property string, value interface{}) *Tween{
    return &Tween{self.Object.Call("updateTweenData", property, value)}
}

// UpdateTweenData1O Updates either a single TweenData or all TweenData objects properties to the given value.
// Used internally by methods like Tween.delay, Tween.yoyo, etc. but can also be called directly if you know which property you want to tweak.
// The property is not checked, so if you pass an invalid one you'll generate a run-time error.
func (self *Tween) UpdateTweenData1O(property string, value interface{}, index int) *Tween{
    return &Tween{self.Object.Call("updateTweenData", property, value, index)}
}

// UpdateTweenDataI Updates either a single TweenData or all TweenData objects properties to the given value.
// Used internally by methods like Tween.delay, Tween.yoyo, etc. but can also be called directly if you know which property you want to tweak.
// The property is not checked, so if you pass an invalid one you'll generate a run-time error.
func (self *Tween) UpdateTweenDataI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("updateTweenData", args)}
}

// Delay Sets the delay in milliseconds before this tween will start. If there are child tweens it sets the delay before the first child starts.
// The delay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to delay.
// If you have child tweens and pass -1 as the index value it sets the delay across all of them.
func (self *Tween) Delay(duration int) *Tween{
    return &Tween{self.Object.Call("delay", duration)}
}

// Delay1O Sets the delay in milliseconds before this tween will start. If there are child tweens it sets the delay before the first child starts.
// The delay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to delay.
// If you have child tweens and pass -1 as the index value it sets the delay across all of them.
func (self *Tween) Delay1O(duration int, index int) *Tween{
    return &Tween{self.Object.Call("delay", duration, index)}
}

// DelayI Sets the delay in milliseconds before this tween will start. If there are child tweens it sets the delay before the first child starts.
// The delay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to delay.
// If you have child tweens and pass -1 as the index value it sets the delay across all of them.
func (self *Tween) DelayI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("delay", args)}
}

// Repeat Sets the number of times this tween will repeat.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to repeat.
// If you have child tweens and pass -1 as the index value it sets the number of times they'll repeat across all of them.
// If you wish to define how many times this Tween and all children will repeat see Tween.repeatAll.
func (self *Tween) Repeat(total int) *Tween{
    return &Tween{self.Object.Call("repeat", total)}
}

// Repeat1O Sets the number of times this tween will repeat.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to repeat.
// If you have child tweens and pass -1 as the index value it sets the number of times they'll repeat across all of them.
// If you wish to define how many times this Tween and all children will repeat see Tween.repeatAll.
func (self *Tween) Repeat1O(total int, repeat int) *Tween{
    return &Tween{self.Object.Call("repeat", total, repeat)}
}

// Repeat2O Sets the number of times this tween will repeat.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to repeat.
// If you have child tweens and pass -1 as the index value it sets the number of times they'll repeat across all of them.
// If you wish to define how many times this Tween and all children will repeat see Tween.repeatAll.
func (self *Tween) Repeat2O(total int, repeat int, index int) *Tween{
    return &Tween{self.Object.Call("repeat", total, repeat, index)}
}

// RepeatI Sets the number of times this tween will repeat.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to repeat.
// If you have child tweens and pass -1 as the index value it sets the number of times they'll repeat across all of them.
// If you wish to define how many times this Tween and all children will repeat see Tween.repeatAll.
func (self *Tween) RepeatI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("repeat", args)}
}

// RepeatDelay Sets the delay in milliseconds before this tween will repeat itself.
// The repeatDelay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to set repeatDelay on.
// If you have child tweens and pass -1 as the index value it sets the repeatDelay across all of them.
func (self *Tween) RepeatDelay(duration int) *Tween{
    return &Tween{self.Object.Call("repeatDelay", duration)}
}

// RepeatDelay1O Sets the delay in milliseconds before this tween will repeat itself.
// The repeatDelay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to set repeatDelay on.
// If you have child tweens and pass -1 as the index value it sets the repeatDelay across all of them.
func (self *Tween) RepeatDelay1O(duration int, index int) *Tween{
    return &Tween{self.Object.Call("repeatDelay", duration, index)}
}

// RepeatDelayI Sets the delay in milliseconds before this tween will repeat itself.
// The repeatDelay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to set repeatDelay on.
// If you have child tweens and pass -1 as the index value it sets the repeatDelay across all of them.
func (self *Tween) RepeatDelayI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("repeatDelay", args)}
}

// Yoyo A Tween that has yoyo set to true will run through from its starting values to its end values and then play back in reverse from end to start.
// Used in combination with repeat you can create endless loops.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to yoyo.
// If you have child tweens and pass -1 as the index value it sets the yoyo property across all of them.
// If you wish to yoyo this Tween and all of its children then see Tween.yoyoAll.
func (self *Tween) Yoyo(enable bool) *Tween{
    return &Tween{self.Object.Call("yoyo", enable)}
}

// Yoyo1O A Tween that has yoyo set to true will run through from its starting values to its end values and then play back in reverse from end to start.
// Used in combination with repeat you can create endless loops.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to yoyo.
// If you have child tweens and pass -1 as the index value it sets the yoyo property across all of them.
// If you wish to yoyo this Tween and all of its children then see Tween.yoyoAll.
func (self *Tween) Yoyo1O(enable bool, yoyoDelay int) *Tween{
    return &Tween{self.Object.Call("yoyo", enable, yoyoDelay)}
}

// Yoyo2O A Tween that has yoyo set to true will run through from its starting values to its end values and then play back in reverse from end to start.
// Used in combination with repeat you can create endless loops.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to yoyo.
// If you have child tweens and pass -1 as the index value it sets the yoyo property across all of them.
// If you wish to yoyo this Tween and all of its children then see Tween.yoyoAll.
func (self *Tween) Yoyo2O(enable bool, yoyoDelay int, index int) *Tween{
    return &Tween{self.Object.Call("yoyo", enable, yoyoDelay, index)}
}

// YoyoI A Tween that has yoyo set to true will run through from its starting values to its end values and then play back in reverse from end to start.
// Used in combination with repeat you can create endless loops.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to yoyo.
// If you have child tweens and pass -1 as the index value it sets the yoyo property across all of them.
// If you wish to yoyo this Tween and all of its children then see Tween.yoyoAll.
func (self *Tween) YoyoI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("yoyo", args)}
}

// YoyoDelay Sets the delay in milliseconds before this tween will run a yoyo (only applies if yoyo is enabled).
// The repeatDelay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to set repeatDelay on.
// If you have child tweens and pass -1 as the index value it sets the repeatDelay across all of them.
func (self *Tween) YoyoDelay(duration int) *Tween{
    return &Tween{self.Object.Call("yoyoDelay", duration)}
}

// YoyoDelay1O Sets the delay in milliseconds before this tween will run a yoyo (only applies if yoyo is enabled).
// The repeatDelay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to set repeatDelay on.
// If you have child tweens and pass -1 as the index value it sets the repeatDelay across all of them.
func (self *Tween) YoyoDelay1O(duration int, index int) *Tween{
    return &Tween{self.Object.Call("yoyoDelay", duration, index)}
}

// YoyoDelayI Sets the delay in milliseconds before this tween will run a yoyo (only applies if yoyo is enabled).
// The repeatDelay is invoked as soon as you call `Tween.start`. If the tween is already running this method doesn't do anything for the current active tween.
// If you have not yet called `Tween.to` or `Tween.from` at least once then this method will do nothing, as there are no tweens to set repeatDelay on.
// If you have child tweens and pass -1 as the index value it sets the repeatDelay across all of them.
func (self *Tween) YoyoDelayI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("yoyoDelay", args)}
}

// Easing Set easing function this tween will use, i.e. Phaser.Easing.Linear.None.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
// If you have child tweens and pass -1 as the index value it sets the easing function defined here across all of them.
func (self *Tween) Easing(ease interface{}) *Tween{
    return &Tween{self.Object.Call("easing", ease)}
}

// Easing1O Set easing function this tween will use, i.e. Phaser.Easing.Linear.None.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
// If you have child tweens and pass -1 as the index value it sets the easing function defined here across all of them.
func (self *Tween) Easing1O(ease interface{}, index int) *Tween{
    return &Tween{self.Object.Call("easing", ease, index)}
}

// EasingI Set easing function this tween will use, i.e. Phaser.Easing.Linear.None.
// The ease function allows you define the rate of change. You can pass either a function such as Phaser.Easing.Circular.Out or a string such as "Circ".
// ".easeIn", ".easeOut" and "easeInOut" variants are all supported for all ease types.
// If you have child tweens and pass -1 as the index value it sets the easing function defined here across all of them.
func (self *Tween) EasingI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("easing", args)}
}

// Interpolation Sets the interpolation function the tween will use. By default it uses Phaser.Math.linearInterpolation.
// Also available: Phaser.Math.bezierInterpolation and Phaser.Math.catmullRomInterpolation.
// The interpolation function is only used if the target properties is an array.
// If you have child tweens and pass -1 as the index value and it will set the interpolation function across all of them.
func (self *Tween) Interpolation(interpolation interface{}) *Tween{
    return &Tween{self.Object.Call("interpolation", interpolation)}
}

// Interpolation1O Sets the interpolation function the tween will use. By default it uses Phaser.Math.linearInterpolation.
// Also available: Phaser.Math.bezierInterpolation and Phaser.Math.catmullRomInterpolation.
// The interpolation function is only used if the target properties is an array.
// If you have child tweens and pass -1 as the index value and it will set the interpolation function across all of them.
func (self *Tween) Interpolation1O(interpolation interface{}, context interface{}) *Tween{
    return &Tween{self.Object.Call("interpolation", interpolation, context)}
}

// Interpolation2O Sets the interpolation function the tween will use. By default it uses Phaser.Math.linearInterpolation.
// Also available: Phaser.Math.bezierInterpolation and Phaser.Math.catmullRomInterpolation.
// The interpolation function is only used if the target properties is an array.
// If you have child tweens and pass -1 as the index value and it will set the interpolation function across all of them.
func (self *Tween) Interpolation2O(interpolation interface{}, context interface{}, index int) *Tween{
    return &Tween{self.Object.Call("interpolation", interpolation, context, index)}
}

// InterpolationI Sets the interpolation function the tween will use. By default it uses Phaser.Math.linearInterpolation.
// Also available: Phaser.Math.bezierInterpolation and Phaser.Math.catmullRomInterpolation.
// The interpolation function is only used if the target properties is an array.
// If you have child tweens and pass -1 as the index value and it will set the interpolation function across all of them.
func (self *Tween) InterpolationI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("interpolation", args)}
}

// RepeatAll Set how many times this tween and all of its children will repeat.
// A tween (A) with 3 children (B,C,D) with a `repeatAll` value of 2 would play as: ABCDABCD before completing.
func (self *Tween) RepeatAll() *Tween{
    return &Tween{self.Object.Call("repeatAll")}
}

// RepeatAll1O Set how many times this tween and all of its children will repeat.
// A tween (A) with 3 children (B,C,D) with a `repeatAll` value of 2 would play as: ABCDABCD before completing.
func (self *Tween) RepeatAll1O(total int) *Tween{
    return &Tween{self.Object.Call("repeatAll", total)}
}

// RepeatAllI Set how many times this tween and all of its children will repeat.
// A tween (A) with 3 children (B,C,D) with a `repeatAll` value of 2 would play as: ABCDABCD before completing.
func (self *Tween) RepeatAllI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("repeatAll", args)}
}

// Chain This method allows you to chain tweens together. Any tween chained to this tween will have its `Tween.start` method called
// as soon as this tween completes. If this tween never completes (i.e. repeatAll or loop is set) then the chain will never progress.
// Note that `Tween.onComplete` will fire when *this* tween completes, not when the whole chain completes.
// For that you should listen to `onComplete` on the final tween in your chain.
// 
// If you pass multiple tweens to this method they will be joined into a single long chain.
// For example if this is Tween A and you pass in B, C and D then B will be chained to A, C will be chained to B and D will be chained to C.
// Any previously chained tweens that may have been set will be overwritten.
func (self *Tween) Chain(tweens *Tween) *Tween{
    return &Tween{self.Object.Call("chain", tweens)}
}

// ChainI This method allows you to chain tweens together. Any tween chained to this tween will have its `Tween.start` method called
// as soon as this tween completes. If this tween never completes (i.e. repeatAll or loop is set) then the chain will never progress.
// Note that `Tween.onComplete` will fire when *this* tween completes, not when the whole chain completes.
// For that you should listen to `onComplete` on the final tween in your chain.
// 
// If you pass multiple tweens to this method they will be joined into a single long chain.
// For example if this is Tween A and you pass in B, C and D then B will be chained to A, C will be chained to B and D will be chained to C.
// Any previously chained tweens that may have been set will be overwritten.
func (self *Tween) ChainI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("chain", args)}
}

// Loop Enables the looping of this tween. The tween will loop forever, and onComplete will never fire.
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
func (self *Tween) Loop() *Tween{
    return &Tween{self.Object.Call("loop")}
}

// Loop1O Enables the looping of this tween. The tween will loop forever, and onComplete will never fire.
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
func (self *Tween) Loop1O(value bool) *Tween{
    return &Tween{self.Object.Call("loop", value)}
}

// LoopI Enables the looping of this tween. The tween will loop forever, and onComplete will never fire.
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
    return &Tween{self.Object.Call("loop", args)}
}

// OnUpdateCallback Sets a callback to be fired each time this tween updates.
func (self *Tween) OnUpdateCallback(callback interface{}, callbackContext interface{}) *Tween{
    return &Tween{self.Object.Call("onUpdateCallback", callback, callbackContext)}
}

// OnUpdateCallbackI Sets a callback to be fired each time this tween updates.
func (self *Tween) OnUpdateCallbackI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("onUpdateCallback", args)}
}

// Pause Pauses the tween. Resume playback with Tween.resume.
func (self *Tween) Pause() {
    self.Object.Call("pause")
}

// PauseI Pauses the tween. Resume playback with Tween.resume.
func (self *Tween) PauseI(args ...interface{}) {
    self.Object.Call("pause", args)
}

// _pause This is called by the core Game loop. Do not call it directly, instead use Tween.pause.
func (self *Tween) _pause() {
    self.Object.Call("_pause")
}

// _pauseI This is called by the core Game loop. Do not call it directly, instead use Tween.pause.
func (self *Tween) _pauseI(args ...interface{}) {
    self.Object.Call("_pause", args)
}

// Resume Resumes a paused tween.
func (self *Tween) Resume() {
    self.Object.Call("resume")
}

// ResumeI Resumes a paused tween.
func (self *Tween) ResumeI(args ...interface{}) {
    self.Object.Call("resume", args)
}

// _resume This is called by the core Game loop. Do not call it directly, instead use Tween.pause.
func (self *Tween) _resume() {
    self.Object.Call("_resume")
}

// _resumeI This is called by the core Game loop. Do not call it directly, instead use Tween.pause.
func (self *Tween) _resumeI(args ...interface{}) {
    self.Object.Call("_resume", args)
}

// Update Core tween update function called by the TweenManager. Does not need to be invoked directly.
func (self *Tween) Update(time int) bool{
    return self.Object.Call("update", time).Bool()
}

// UpdateI Core tween update function called by the TweenManager. Does not need to be invoked directly.
func (self *Tween) UpdateI(args ...interface{}) bool{
    return self.Object.Call("update", args).Bool()
}

// GenerateData This will generate an array populated with the tweened object values from start to end.
// It works by running the tween simulation at the given frame rate based on the values set-up in Tween.to and Tween.from.
// It ignores delay and repeat counts and any chained tweens, but does include child tweens.
// Just one play through of the tween data is returned, including yoyo if set.
func (self *Tween) GenerateData() []interface{}{
	array00 := self.Object.Call("generateData")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GenerateData1O This will generate an array populated with the tweened object values from start to end.
// It works by running the tween simulation at the given frame rate based on the values set-up in Tween.to and Tween.from.
// It ignores delay and repeat counts and any chained tweens, but does include child tweens.
// Just one play through of the tween data is returned, including yoyo if set.
func (self *Tween) GenerateData1O(frameRate int) []interface{}{
	array00 := self.Object.Call("generateData", frameRate)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GenerateData2O This will generate an array populated with the tweened object values from start to end.
// It works by running the tween simulation at the given frame rate based on the values set-up in Tween.to and Tween.from.
// It ignores delay and repeat counts and any chained tweens, but does include child tweens.
// Just one play through of the tween data is returned, including yoyo if set.
func (self *Tween) GenerateData2O(frameRate int, data []interface{}) []interface{}{
	array00 := self.Object.Call("generateData", frameRate, data)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GenerateDataI This will generate an array populated with the tweened object values from start to end.
// It works by running the tween simulation at the given frame rate based on the values set-up in Tween.to and Tween.from.
// It ignores delay and repeat counts and any chained tweens, but does include child tweens.
// Just one play through of the tween data is returned, including yoyo if set.
func (self *Tween) GenerateDataI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("generateData", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

