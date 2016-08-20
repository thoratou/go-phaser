// Automatic generation for Phaser.TweenData
// generated file TweenData.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A Phaser.Tween contains at least one TweenData object. It contains all of the tween data values, such as the
// starting and ending values, the ease function, interpolation and duration. The Tween acts as a timeline manager for
// TweenData objects and can contain multiple TweenData objects.
type TweenData struct {
    *js.Object
}


// The Tween which owns this TweenData.
func (self *TweenData) GetParent() *Tween{
    return &Tween{self.Get("parent")}
}

// The Tween which owns this TweenData.
func (self *TweenData) SetParent(member *Tween) {
    self.Set("parent", member)
}

// A reference to the currently running Game.
func (self *TweenData) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running Game.
func (self *TweenData) SetGame(member *Game) {
    self.Set("game", member)
}

// The duration of the tween in ms.
func (self *TweenData) GetDuration() float64{
    return self.Get("duration").Float()
}

// The duration of the tween in ms.
func (self *TweenData) SetDuration(member float64) {
    self.Set("duration", member)
}

// A value between 0 and 1 that represents how far through the duration this tween is.
func (self *TweenData) GetPercent() float64{
    return self.Get("percent").Float()
}

// A value between 0 and 1 that represents how far through the duration this tween is.
func (self *TweenData) SetPercent(member float64) {
    self.Set("percent", member)
}

// The current calculated value.
func (self *TweenData) GetValue() float64{
    return self.Get("value").Float()
}

// The current calculated value.
func (self *TweenData) SetValue(member float64) {
    self.Set("value", member)
}

// If the Tween is set to repeat this contains the current repeat count.
func (self *TweenData) GetRepeatCounter() float64{
    return self.Get("repeatCounter").Float()
}

// If the Tween is set to repeat this contains the current repeat count.
func (self *TweenData) SetRepeatCounter(member float64) {
    self.Set("repeatCounter", member)
}

// The amount of time in ms between repeats of this tween.
func (self *TweenData) GetRepeatDelay() float64{
    return self.Get("repeatDelay").Float()
}

// The amount of time in ms between repeats of this tween.
func (self *TweenData) SetRepeatDelay(member float64) {
    self.Set("repeatDelay", member)
}

// The total number of times this Tween will repeat.
func (self *TweenData) GetRepeatTotal() float64{
    return self.Get("repeatTotal").Float()
}

// The total number of times this Tween will repeat.
func (self *TweenData) SetRepeatTotal(member float64) {
    self.Set("repeatTotal", member)
}

// True if the Tween will use interpolation (i.e. is an Array to Array tween)
func (self *TweenData) GetInterpolate() bool{
    return self.Get("interpolate").Bool()
}

// True if the Tween will use interpolation (i.e. is an Array to Array tween)
func (self *TweenData) SetInterpolate(member bool) {
    self.Set("interpolate", member)
}

// True if the Tween is set to yoyo, otherwise false.
func (self *TweenData) GetYoyo() bool{
    return self.Get("yoyo").Bool()
}

// True if the Tween is set to yoyo, otherwise false.
func (self *TweenData) SetYoyo(member bool) {
    self.Set("yoyo", member)
}

// The amount of time in ms between yoyos of this tween.
func (self *TweenData) GetYoyoDelay() float64{
    return self.Get("yoyoDelay").Float()
}

// The amount of time in ms between yoyos of this tween.
func (self *TweenData) SetYoyoDelay(member float64) {
    self.Set("yoyoDelay", member)
}

// When a Tween is yoyoing this value holds if it's currently playing forwards (false) or in reverse (true).
func (self *TweenData) GetInReverse() bool{
    return self.Get("inReverse").Bool()
}

// When a Tween is yoyoing this value holds if it's currently playing forwards (false) or in reverse (true).
func (self *TweenData) SetInReverse(member bool) {
    self.Set("inReverse", member)
}

// The amount to delay by until the Tween starts (in ms). Only applies to the start, use repeatDelay to handle repeats.
func (self *TweenData) GetDelay() float64{
    return self.Get("delay").Float()
}

// The amount to delay by until the Tween starts (in ms). Only applies to the start, use repeatDelay to handle repeats.
func (self *TweenData) SetDelay(member float64) {
    self.Set("delay", member)
}

// Current time value.
func (self *TweenData) GetDt() float64{
    return self.Get("dt").Float()
}

// Current time value.
func (self *TweenData) SetDt(member float64) {
    self.Set("dt", member)
}

// The time the Tween started or null if it hasn't yet started.
func (self *TweenData) GetStartTime() float64{
    return self.Get("startTime").Float()
}

// The time the Tween started or null if it hasn't yet started.
func (self *TweenData) SetStartTime(member float64) {
    self.Set("startTime", member)
}

// The easing function used for the Tween.
func (self *TweenData) SetEasingFunction(member func(...interface{})) {
    self.Set("easingFunction", member)
}

// The interpolation function used for the Tween.
func (self *TweenData) SetInterpolationFunction(member func(...interface{})) {
    self.Set("interpolationFunction", member)
}

// The interpolation function context used for the Tween.
func (self *TweenData) GetInterpolationContext() interface{}{
    return self.Get("interpolationContext")
}

// The interpolation function context used for the Tween.
func (self *TweenData) SetInterpolationContext(member interface{}) {
    self.Set("interpolationContext", member)
}

// If the tween is running this is set to `true`. Unless Phaser.Tween a TweenData that is waiting for a delay to expire is *not* considered as running.
func (self *TweenData) GetIsRunning() bool{
    return self.Get("isRunning").Bool()
}

// If the tween is running this is set to `true`. Unless Phaser.Tween a TweenData that is waiting for a delay to expire is *not* considered as running.
func (self *TweenData) SetIsRunning(member bool) {
    self.Set("isRunning", member)
}

// Is this a from tween or a to tween?
func (self *TweenData) GetIsFrom() bool{
    return self.Get("isFrom").Bool()
}

// Is this a from tween or a to tween?
func (self *TweenData) SetIsFrom(member bool) {
    self.Set("isFrom", member)
}

// 
func (self *TweenData) GetPENDING() float64{
    return self.Get("PENDING").Float()
}

// 
func (self *TweenData) SetPENDING(member float64) {
    self.Set("PENDING", member)
}

// 
func (self *TweenData) GetRUNNING() float64{
    return self.Get("RUNNING").Float()
}

// 
func (self *TweenData) SetRUNNING(member float64) {
    self.Set("RUNNING", member)
}

// 
func (self *TweenData) GetLOOPED() float64{
    return self.Get("LOOPED").Float()
}

// 
func (self *TweenData) SetLOOPED(member float64) {
    self.Set("LOOPED", member)
}

// 
func (self *TweenData) GetCOMPLETE() float64{
    return self.Get("COMPLETE").Float()
}

// 
func (self *TweenData) SetCOMPLETE(member float64) {
    self.Set("COMPLETE", member)
}



// Sets this tween to be a `to` tween on the properties given. A `to` tween starts at the current value and tweens to the destination value given.
// For example a Sprite with an `x` coordinate of 100 could be tweened to `x` 200 by giving a properties object of `{ x: 200 }`.
func (self *TweenData) ToI(args ...interface{}) *TweenData{
    return &TweenData{self.Call("to", args)}
}

// Sets this tween to be a `from` tween on the properties given. A `from` tween sets the target to the destination value and tweens to its current value.
// For example a Sprite with an `x` coordinate of 100 tweened from `x` 500 would be set to `x` 500 and then tweened to `x` 100 by giving a properties object of `{ x: 500 }`.
func (self *TweenData) FromI(args ...interface{}) *TweenData{
    return &TweenData{self.Call("from", args)}
}

// Starts the Tween running.
func (self *TweenData) StartI(args ...interface{}) *TweenData{
    return &TweenData{self.Call("start", args)}
}

// Loads the values from the target object into this Tween.
func (self *TweenData) LoadValuesI(args ...interface{}) *TweenData{
    return &TweenData{self.Call("loadValues", args)}
}

// Updates this Tween. This is called automatically by Phaser.Tween.
func (self *TweenData) UpdateI(args ...interface{}) float64{
    return self.Call("update", args).Float()
}

// This will generate an array populated with the tweened object values from start to end.
// It works by running the tween simulation at the given frame rate based on the values set-up in Tween.to and Tween.from.
// Just one play through of the tween data is returned, including yoyo if set.
func (self *TweenData) GenerateDataI(args ...interface{}) []interface{}{
	array := self.Call("generateData", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Checks if this Tween is meant to repeat or yoyo and handles doing so.
func (self *TweenData) RepeatI(args ...interface{}) float64{
    return self.Call("repeat", args).Float()
}
