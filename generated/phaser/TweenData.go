// Package phaser Automatic generation for Phaser.TweenData
// generated file TweenData.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// TweenData A Phaser.Tween contains at least one TweenData object. It contains all of the tween data values, such as the
// starting and ending values, the ease function, interpolation and duration. The Tween acts as a timeline manager for
// TweenData objects and can contain multiple TweenData objects.
type TweenData struct {
    *js.Object
}

// NewTweenData A Phaser.Tween contains at least one TweenData object. It contains all of the tween data values, such as the
// starting and ending values, the ease function, interpolation and duration. The Tween acts as a timeline manager for
// TweenData objects and can contain multiple TweenData objects.
func NewTweenData(parent *Tween) *TweenData {
    return &TweenData{js.Global.Get("Phaser").Get("TweenData").New(parent)}
}
// NewTweenDataI A Phaser.Tween contains at least one TweenData object. It contains all of the tween data values, such as the
// starting and ending values, the ease function, interpolation and duration. The Tween acts as a timeline manager for
// TweenData objects and can contain multiple TweenData objects.
func NewTweenDataI(args ...interface{}) *TweenData {
    return &TweenData{js.Global.Get("Phaser").Get("TweenData").New(args)}
}



// TweenData Binding conversion method to TweenData point 
func ToTweenData(jsStruct interface{}) *TweenData {
    if object, ok := jsStruct.(*js.Object); ok {
		return &TweenData{Object: object}
	}
	return nil
}



// Parent The Tween which owns this TweenData.
func (self *TweenData) Parent() *Tween{
    return &Tween{self.Object.Get("parent")}
}

// SetParentA The Tween which owns this TweenData.
func (self *TweenData) SetParentA(member *Tween) {
    self.Object.Set("parent", member)
}

// Game A reference to the currently running Game.
func (self *TweenData) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running Game.
func (self *TweenData) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Duration The duration of the tween in ms.
func (self *TweenData) Duration() int{
    return self.Object.Get("duration").Int()
}

// SetDurationA The duration of the tween in ms.
func (self *TweenData) SetDurationA(member int) {
    self.Object.Set("duration", member)
}

// Percent A value between 0 and 1 that represents how far through the duration this tween is.
func (self *TweenData) Percent() int{
    return self.Object.Get("percent").Int()
}

// SetPercentA A value between 0 and 1 that represents how far through the duration this tween is.
func (self *TweenData) SetPercentA(member int) {
    self.Object.Set("percent", member)
}

// Value The current calculated value.
func (self *TweenData) Value() int{
    return self.Object.Get("value").Int()
}

// SetValueA The current calculated value.
func (self *TweenData) SetValueA(member int) {
    self.Object.Set("value", member)
}

// RepeatCounter If the Tween is set to repeat this contains the current repeat count.
func (self *TweenData) RepeatCounter() int{
    return self.Object.Get("repeatCounter").Int()
}

// SetRepeatCounterA If the Tween is set to repeat this contains the current repeat count.
func (self *TweenData) SetRepeatCounterA(member int) {
    self.Object.Set("repeatCounter", member)
}

// RepeatDelay The amount of time in ms between repeats of this tween.
func (self *TweenData) RepeatDelay() int{
    return self.Object.Get("repeatDelay").Int()
}

// SetRepeatDelayA The amount of time in ms between repeats of this tween.
func (self *TweenData) SetRepeatDelayA(member int) {
    self.Object.Set("repeatDelay", member)
}

// RepeatTotal The total number of times this Tween will repeat.
func (self *TweenData) RepeatTotal() int{
    return self.Object.Get("repeatTotal").Int()
}

// SetRepeatTotalA The total number of times this Tween will repeat.
func (self *TweenData) SetRepeatTotalA(member int) {
    self.Object.Set("repeatTotal", member)
}

// Interpolate True if the Tween will use interpolation (i.e. is an Array to Array tween)
func (self *TweenData) Interpolate() bool{
    return self.Object.Get("interpolate").Bool()
}

// SetInterpolateA True if the Tween will use interpolation (i.e. is an Array to Array tween)
func (self *TweenData) SetInterpolateA(member bool) {
    self.Object.Set("interpolate", member)
}

// Yoyo True if the Tween is set to yoyo, otherwise false.
func (self *TweenData) Yoyo() bool{
    return self.Object.Get("yoyo").Bool()
}

// SetYoyoA True if the Tween is set to yoyo, otherwise false.
func (self *TweenData) SetYoyoA(member bool) {
    self.Object.Set("yoyo", member)
}

// YoyoDelay The amount of time in ms between yoyos of this tween.
func (self *TweenData) YoyoDelay() int{
    return self.Object.Get("yoyoDelay").Int()
}

// SetYoyoDelayA The amount of time in ms between yoyos of this tween.
func (self *TweenData) SetYoyoDelayA(member int) {
    self.Object.Set("yoyoDelay", member)
}

// InReverse When a Tween is yoyoing this value holds if it's currently playing forwards (false) or in reverse (true).
func (self *TweenData) InReverse() bool{
    return self.Object.Get("inReverse").Bool()
}

// SetInReverseA When a Tween is yoyoing this value holds if it's currently playing forwards (false) or in reverse (true).
func (self *TweenData) SetInReverseA(member bool) {
    self.Object.Set("inReverse", member)
}

// Delay The amount to delay by until the Tween starts (in ms). Only applies to the start, use repeatDelay to handle repeats.
func (self *TweenData) Delay() int{
    return self.Object.Get("delay").Int()
}

// SetDelayA The amount to delay by until the Tween starts (in ms). Only applies to the start, use repeatDelay to handle repeats.
func (self *TweenData) SetDelayA(member int) {
    self.Object.Set("delay", member)
}

// Dt Current time value.
func (self *TweenData) Dt() int{
    return self.Object.Get("dt").Int()
}

// SetDtA Current time value.
func (self *TweenData) SetDtA(member int) {
    self.Object.Set("dt", member)
}

// StartTime The time the Tween started or null if it hasn't yet started.
func (self *TweenData) StartTime() int{
    return self.Object.Get("startTime").Int()
}

// SetStartTimeA The time the Tween started or null if it hasn't yet started.
func (self *TweenData) SetStartTimeA(member int) {
    self.Object.Set("startTime", member)
}

// EasingFunction The easing function used for the Tween.
func (self *TweenData) EasingFunction() interface{}{
    return self.Object.Get("easingFunction")
}

// SetEasingFunctionA The easing function used for the Tween.
func (self *TweenData) SetEasingFunctionA(member interface{}) {
    self.Object.Set("easingFunction", member)
}

// InterpolationFunction The interpolation function used for the Tween.
func (self *TweenData) InterpolationFunction() interface{}{
    return self.Object.Get("interpolationFunction")
}

// SetInterpolationFunctionA The interpolation function used for the Tween.
func (self *TweenData) SetInterpolationFunctionA(member interface{}) {
    self.Object.Set("interpolationFunction", member)
}

// InterpolationContext The interpolation function context used for the Tween.
func (self *TweenData) InterpolationContext() interface{}{
    return self.Object.Get("interpolationContext")
}

// SetInterpolationContextA The interpolation function context used for the Tween.
func (self *TweenData) SetInterpolationContextA(member interface{}) {
    self.Object.Set("interpolationContext", member)
}

// IsRunning If the tween is running this is set to `true`. Unless Phaser.Tween a TweenData that is waiting for a delay to expire is *not* considered as running.
func (self *TweenData) IsRunning() bool{
    return self.Object.Get("isRunning").Bool()
}

// SetIsRunningA If the tween is running this is set to `true`. Unless Phaser.Tween a TweenData that is waiting for a delay to expire is *not* considered as running.
func (self *TweenData) SetIsRunningA(member bool) {
    self.Object.Set("isRunning", member)
}

// IsFrom Is this a from tween or a to tween?
func (self *TweenData) IsFrom() bool{
    return self.Object.Get("isFrom").Bool()
}

// SetIsFromA Is this a from tween or a to tween?
func (self *TweenData) SetIsFromA(member bool) {
    self.Object.Set("isFrom", member)
}

// PENDING empty description
func (self *TweenData) PENDING() int{
    return self.Object.Get("PENDING").Int()
}

// SetPENDINGA empty description
func (self *TweenData) SetPENDINGA(member int) {
    self.Object.Set("PENDING", member)
}

// RUNNING empty description
func (self *TweenData) RUNNING() int{
    return self.Object.Get("RUNNING").Int()
}

// SetRUNNINGA empty description
func (self *TweenData) SetRUNNINGA(member int) {
    self.Object.Set("RUNNING", member)
}

// LOOPED empty description
func (self *TweenData) LOOPED() int{
    return self.Object.Get("LOOPED").Int()
}

// SetLOOPEDA empty description
func (self *TweenData) SetLOOPEDA(member int) {
    self.Object.Set("LOOPED", member)
}

// COMPLETE empty description
func (self *TweenData) COMPLETE() int{
    return self.Object.Get("COMPLETE").Int()
}

// SetCOMPLETEA empty description
func (self *TweenData) SetCOMPLETEA(member int) {
    self.Object.Set("COMPLETE", member)
}


// To Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
func (self *TweenData) To(properties interface{}) *TweenData{
    return &TweenData{self.Object.Call("to", properties)}
}

// To1O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
func (self *TweenData) To1O(properties interface{}, duration int) *TweenData{
    return &TweenData{self.Object.Call("to", properties, duration)}
}

// To2O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
func (self *TweenData) To2O(properties interface{}, duration int, ease interface{}) *TweenData{
    return &TweenData{self.Object.Call("to", properties, duration, ease)}
}

// To3O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
func (self *TweenData) To3O(properties interface{}, duration int, ease interface{}, delay int) *TweenData{
    return &TweenData{self.Object.Call("to", properties, duration, ease, delay)}
}

// To4O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
func (self *TweenData) To4O(properties interface{}, duration int, ease interface{}, delay int, repeat int) *TweenData{
    return &TweenData{self.Object.Call("to", properties, duration, ease, delay, repeat)}
}

// To5O Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
func (self *TweenData) To5O(properties interface{}, duration int, ease interface{}, delay int, repeat int, yoyo bool) *TweenData{
    return &TweenData{self.Object.Call("to", properties, duration, ease, delay, repeat, yoyo)}
}

// ToI Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
func (self *TweenData) ToI(args ...interface{}) *TweenData{
    return &TweenData{self.Object.Call("to", args)}
}

// From Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
func (self *TweenData) From(properties interface{}) *TweenData{
    return &TweenData{self.Object.Call("from", properties)}
}

// From1O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
func (self *TweenData) From1O(properties interface{}, duration int) *TweenData{
    return &TweenData{self.Object.Call("from", properties, duration)}
}

// From2O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
func (self *TweenData) From2O(properties interface{}, duration int, ease interface{}) *TweenData{
    return &TweenData{self.Object.Call("from", properties, duration, ease)}
}

// From3O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
func (self *TweenData) From3O(properties interface{}, duration int, ease interface{}, delay int) *TweenData{
    return &TweenData{self.Object.Call("from", properties, duration, ease, delay)}
}

// From4O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
func (self *TweenData) From4O(properties interface{}, duration int, ease interface{}, delay int, repeat int) *TweenData{
    return &TweenData{self.Object.Call("from", properties, duration, ease, delay, repeat)}
}

// From5O Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
func (self *TweenData) From5O(properties interface{}, duration int, ease interface{}, delay int, repeat int, yoyo bool) *TweenData{
    return &TweenData{self.Object.Call("from", properties, duration, ease, delay, repeat, yoyo)}
}

// FromI Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
func (self *TweenData) FromI(args ...interface{}) *TweenData{
    return &TweenData{self.Object.Call("from", args)}
}

// Start Starts the Tween running.
func (self *TweenData) Start() *TweenData{
    return &TweenData{self.Object.Call("start")}
}

// StartI Starts the Tween running.
func (self *TweenData) StartI(args ...interface{}) *TweenData{
    return &TweenData{self.Object.Call("start", args)}
}

// LoadValues Loads the values from the target object into this Tween.
func (self *TweenData) LoadValues() *TweenData{
    return &TweenData{self.Object.Call("loadValues")}
}

// LoadValuesI Loads the values from the target object into this Tween.
func (self *TweenData) LoadValuesI(args ...interface{}) *TweenData{
    return &TweenData{self.Object.Call("loadValues", args)}
}

// Update Updates this Tween. This is called automatically by Phaser.Tween.
func (self *TweenData) Update(time int) int{
    return self.Object.Call("update", time).Int()
}

// UpdateI Updates this Tween. This is called automatically by Phaser.Tween.
func (self *TweenData) UpdateI(args ...interface{}) int{
    return self.Object.Call("update", args).Int()
}

// GenerateData This will generate an array populated with the tweened object values from start to end.
// It works by running the tween simulation at the given frame rate based on the values set-up in Tween.to and Tween.from.
// Just one play through of the tween data is returned, including yoyo if set.
func (self *TweenData) GenerateData() []interface{}{
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
// Just one play through of the tween data is returned, including yoyo if set.
func (self *TweenData) GenerateData1O(frameRate int) []interface{}{
	array00 := self.Object.Call("generateData", frameRate)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GenerateDataI This will generate an array populated with the tweened object values from start to end.
// It works by running the tween simulation at the given frame rate based on the values set-up in Tween.to and Tween.from.
// Just one play through of the tween data is returned, including yoyo if set.
func (self *TweenData) GenerateDataI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("generateData", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Repeat Checks if this Tween is meant to repeat or yoyo and handles doing so.
func (self *TweenData) Repeat() int{
    return self.Object.Call("repeat").Int()
}

// RepeatI Checks if this Tween is meant to repeat or yoyo and handles doing so.
func (self *TweenData) RepeatI(args ...interface{}) int{
    return self.Object.Call("repeat", args).Int()
}

