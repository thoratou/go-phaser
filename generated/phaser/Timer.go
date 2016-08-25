// Package phaser Automatic generation for Phaser.Timer
// generated file Timer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Timer A Timer is a way to create and manage {@link Phaser.TimerEvent timer events} that wait for a specific duration and then run a callback.
// Many different timer events, with individual delays, can be added to the same Timer.
// 
// All Timer delays are in milliseconds (there are 1000 ms in 1 second); so a delay value of 250 represents a quarter of a second.
// 
// Timers are based on real life time, adjusted for game pause durations.
// That is, *timer events are based on elapsed {@link Phaser.Time game time}* and do *not* take physics time or slow motion into account.
type Timer struct {
    *js.Object
}

// NewTimer A Timer is a way to create and manage {@link Phaser.TimerEvent timer events} that wait for a specific duration and then run a callback.
// Many different timer events, with individual delays, can be added to the same Timer.
// 
// All Timer delays are in milliseconds (there are 1000 ms in 1 second); so a delay value of 250 represents a quarter of a second.
// 
// Timers are based on real life time, adjusted for game pause durations.
// That is, *timer events are based on elapsed {@link Phaser.Time game time}* and do *not* take physics time or slow motion into account.
func NewTimer(game *Game) *Timer {
    return &Timer{js.Global.Get("Phaser").Get("Timer").New(game)}
}
// NewTimer1O A Timer is a way to create and manage {@link Phaser.TimerEvent timer events} that wait for a specific duration and then run a callback.
// Many different timer events, with individual delays, can be added to the same Timer.
// 
// All Timer delays are in milliseconds (there are 1000 ms in 1 second); so a delay value of 250 represents a quarter of a second.
// 
// Timers are based on real life time, adjusted for game pause durations.
// That is, *timer events are based on elapsed {@link Phaser.Time game time}* and do *not* take physics time or slow motion into account.
func NewTimer1O(game *Game, autoDestroy bool) *Timer {
    return &Timer{js.Global.Get("Phaser").Get("Timer").New(game, autoDestroy)}
}
// NewTimerI A Timer is a way to create and manage {@link Phaser.TimerEvent timer events} that wait for a specific duration and then run a callback.
// Many different timer events, with individual delays, can be added to the same Timer.
// 
// All Timer delays are in milliseconds (there are 1000 ms in 1 second); so a delay value of 250 represents a quarter of a second.
// 
// Timers are based on real life time, adjusted for game pause durations.
// That is, *timer events are based on elapsed {@link Phaser.Time game time}* and do *not* take physics time or slow motion into account.
func NewTimerI(args ...interface{}) *Timer {
    return &Timer{js.Global.Get("Phaser").Get("Timer").New(args)}
}



// Timer Binding conversion method to Timer point 
func ToTimer(jsStruct interface{}) *Timer {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Timer{Object: object}
	}
	return nil
}



// Game Local reference to game.
func (self *Timer) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA Local reference to game.
func (self *Timer) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Running True if the Timer is actively running.
// 
// Do not modify this boolean - use {@link Phaser.Timer#pause pause} (and {@link Phaser.Timer#resume resume}) to pause the timer.
func (self *Timer) Running() bool{
    return self.Object.Get("running").Bool()
}

// SetRunningA True if the Timer is actively running.
// 
// Do not modify this boolean - use {@link Phaser.Timer#pause pause} (and {@link Phaser.Timer#resume resume}) to pause the timer.
func (self *Timer) SetRunningA(member bool) {
    self.Object.Set("running", member)
}

// AutoDestroy If true, the timer will automatically destroy itself after all the events have been dispatched (assuming no looping events).
func (self *Timer) AutoDestroy() bool{
    return self.Object.Get("autoDestroy").Bool()
}

// SetAutoDestroyA If true, the timer will automatically destroy itself after all the events have been dispatched (assuming no looping events).
func (self *Timer) SetAutoDestroyA(member bool) {
    self.Object.Set("autoDestroy", member)
}

// Expired An expired Timer is one in which all of its events have been dispatched and none are pending.
func (self *Timer) Expired() bool{
    return self.Object.Get("expired").Bool()
}

// SetExpiredA An expired Timer is one in which all of its events have been dispatched and none are pending.
func (self *Timer) SetExpiredA(member bool) {
    self.Object.Set("expired", member)
}

// Elapsed Elapsed time since the last frame (in ms).
func (self *Timer) Elapsed() int{
    return self.Object.Get("elapsed").Int()
}

// SetElapsedA Elapsed time since the last frame (in ms).
func (self *Timer) SetElapsedA(member int) {
    self.Object.Set("elapsed", member)
}

// Events An array holding all of this timers Phaser.TimerEvent objects. Use the methods add, repeat and loop to populate it.
func (self *Timer) Events() []TimerEvent{
	array00 := self.Object.Get("events")
	length00 := array00.Length()
	out00 := make([]TimerEvent, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = TimerEvent{array00.Index(i00)}
	}
	return out00
}

// SetEventsA An array holding all of this timers Phaser.TimerEvent objects. Use the methods add, repeat and loop to populate it.
func (self *Timer) SetEventsA(member []TimerEvent) {
    self.Object.Set("events", member)
}

// OnComplete This signal will be dispatched when this Timer has completed which means that there are no more events in the queue.
// 
// The signal is supplied with one argument, `timer`, which is this Timer object.
func (self *Timer) OnComplete() *Signal{
    return &Signal{self.Object.Get("onComplete")}
}

// SetOnCompleteA This signal will be dispatched when this Timer has completed which means that there are no more events in the queue.
// 
// The signal is supplied with one argument, `timer`, which is this Timer object.
func (self *Timer) SetOnCompleteA(member *Signal) {
    self.Object.Set("onComplete", member)
}

// NextTick The time the next tick will occur.
func (self *Timer) NextTick() int{
    return self.Object.Get("nextTick").Int()
}

// SetNextTickA The time the next tick will occur.
func (self *Timer) SetNextTickA(member int) {
    self.Object.Set("nextTick", member)
}

// TimeCap If the difference in time between two frame updates exceeds this value, the event times are reset to avoid catch-up situations.
func (self *Timer) TimeCap() int{
    return self.Object.Get("timeCap").Int()
}

// SetTimeCapA If the difference in time between two frame updates exceeds this value, the event times are reset to avoid catch-up situations.
func (self *Timer) SetTimeCapA(member int) {
    self.Object.Set("timeCap", member)
}

// Paused The paused state of the Timer. You can pause the timer by calling Timer.pause() and Timer.resume() or by the game pausing.
func (self *Timer) Paused() bool{
    return self.Object.Get("paused").Bool()
}

// SetPausedA The paused state of the Timer. You can pause the timer by calling Timer.pause() and Timer.resume() or by the game pausing.
func (self *Timer) SetPausedA(member bool) {
    self.Object.Set("paused", member)
}

// MINUTE Number of milliseconds in a minute.
func (self *Timer) MINUTE() int{
    return self.Object.Get("MINUTE").Int()
}

// SetMINUTEA Number of milliseconds in a minute.
func (self *Timer) SetMINUTEA(member int) {
    self.Object.Set("MINUTE", member)
}

// SECOND Number of milliseconds in a second.
func (self *Timer) SECOND() int{
    return self.Object.Get("SECOND").Int()
}

// SetSECONDA Number of milliseconds in a second.
func (self *Timer) SetSECONDA(member int) {
    self.Object.Set("SECOND", member)
}

// HALF Number of milliseconds in half a second.
func (self *Timer) HALF() int{
    return self.Object.Get("HALF").Int()
}

// SetHALFA Number of milliseconds in half a second.
func (self *Timer) SetHALFA(member int) {
    self.Object.Set("HALF", member)
}

// QUARTER Number of milliseconds in a quarter of a second.
func (self *Timer) QUARTER() int{
    return self.Object.Get("QUARTER").Int()
}

// SetQUARTERA Number of milliseconds in a quarter of a second.
func (self *Timer) SetQUARTERA(member int) {
    self.Object.Set("QUARTER", member)
}

// Next The time at which the next event will occur.
func (self *Timer) Next() int{
    return self.Object.Get("next").Int()
}

// SetNextA The time at which the next event will occur.
func (self *Timer) SetNextA(member int) {
    self.Object.Set("next", member)
}

// Duration The duration in ms remaining until the next event will occur.
func (self *Timer) Duration() int{
    return self.Object.Get("duration").Int()
}

// SetDurationA The duration in ms remaining until the next event will occur.
func (self *Timer) SetDurationA(member int) {
    self.Object.Set("duration", member)
}

// Length The number of pending events in the queue.
func (self *Timer) Length() int{
    return self.Object.Get("length").Int()
}

// SetLengthA The number of pending events in the queue.
func (self *Timer) SetLengthA(member int) {
    self.Object.Set("length", member)
}

// Ms The duration in milliseconds that this Timer has been running for.
func (self *Timer) Ms() int{
    return self.Object.Get("ms").Int()
}

// SetMsA The duration in milliseconds that this Timer has been running for.
func (self *Timer) SetMsA(member int) {
    self.Object.Set("ms", member)
}

// Seconds The duration in seconds that this Timer has been running for.
func (self *Timer) Seconds() int{
    return self.Object.Get("seconds").Int()
}

// SetSecondsA The duration in seconds that this Timer has been running for.
func (self *Timer) SetSecondsA(member int) {
    self.Object.Set("seconds", member)
}


// Create Creates a new TimerEvent on this Timer.
// 
// Use {@link Phaser.Timer#add}, {@link Phaser.Timer#repeat}, or {@link Phaser.Timer#loop} methods to create a new event.
func (self *Timer) Create(delay int, loop bool, repeatCount int, callback interface{}, callbackContext interface{}, arguments []interface{}) *TimerEvent{
    return &TimerEvent{self.Object.Call("create", delay, loop, repeatCount, callback, callbackContext, arguments)}
}

// CreateI Creates a new TimerEvent on this Timer.
// 
// Use {@link Phaser.Timer#add}, {@link Phaser.Timer#repeat}, or {@link Phaser.Timer#loop} methods to create a new event.
func (self *Timer) CreateI(args ...interface{}) *TimerEvent{
    return &TimerEvent{self.Object.Call("create", args)}
}

// Add Adds a new Event to this Timer.
// 
// The event will fire after the given amount of `delay` in milliseconds has passed, once the Timer has started running.
// The delay is in relation to when the Timer starts, not the time it was added. If the Timer is already running the delay will be calculated based on the timers current time.
// 
// Make sure to call {@link Phaser.Timer#start start} after adding all of the Events you require for this Timer.
func (self *Timer) Add(delay int, callback interface{}, callbackContext interface{}, arguments interface{}) *TimerEvent{
    return &TimerEvent{self.Object.Call("add", delay, callback, callbackContext, arguments)}
}

// AddI Adds a new Event to this Timer.
// 
// The event will fire after the given amount of `delay` in milliseconds has passed, once the Timer has started running.
// The delay is in relation to when the Timer starts, not the time it was added. If the Timer is already running the delay will be calculated based on the timers current time.
// 
// Make sure to call {@link Phaser.Timer#start start} after adding all of the Events you require for this Timer.
func (self *Timer) AddI(args ...interface{}) *TimerEvent{
    return &TimerEvent{self.Object.Call("add", args)}
}

// Repeat Adds a new TimerEvent that will always play through once and then repeat for the given number of iterations.
// 
// The event will fire after the given amount of `delay` in milliseconds has passed, once the Timer has started running.
// The delay is in relation to when the Timer starts, not the time it was added.
// If the Timer is already running the delay will be calculated based on the timers current time.
// 
// Make sure to call {@link Phaser.Timer#start start} after adding all of the Events you require for this Timer.
func (self *Timer) Repeat(delay int, repeatCount int, callback interface{}, callbackContext interface{}, arguments interface{}) *TimerEvent{
    return &TimerEvent{self.Object.Call("repeat", delay, repeatCount, callback, callbackContext, arguments)}
}

// RepeatI Adds a new TimerEvent that will always play through once and then repeat for the given number of iterations.
// 
// The event will fire after the given amount of `delay` in milliseconds has passed, once the Timer has started running.
// The delay is in relation to when the Timer starts, not the time it was added.
// If the Timer is already running the delay will be calculated based on the timers current time.
// 
// Make sure to call {@link Phaser.Timer#start start} after adding all of the Events you require for this Timer.
func (self *Timer) RepeatI(args ...interface{}) *TimerEvent{
    return &TimerEvent{self.Object.Call("repeat", args)}
}

// Loop Adds a new looped Event to this Timer that will repeat forever or until the Timer is stopped.
// 
// The event will fire after the given amount of `delay` in milliseconds has passed, once the Timer has started running.
// The delay is in relation to when the Timer starts, not the time it was added. If the Timer is already running the delay will be calculated based on the timers current time.
// 
// Make sure to call {@link Phaser.Timer#start start} after adding all of the Events you require for this Timer.
func (self *Timer) Loop(delay int, callback interface{}, callbackContext interface{}, arguments interface{}) *TimerEvent{
    return &TimerEvent{self.Object.Call("loop", delay, callback, callbackContext, arguments)}
}

// LoopI Adds a new looped Event to this Timer that will repeat forever or until the Timer is stopped.
// 
// The event will fire after the given amount of `delay` in milliseconds has passed, once the Timer has started running.
// The delay is in relation to when the Timer starts, not the time it was added. If the Timer is already running the delay will be calculated based on the timers current time.
// 
// Make sure to call {@link Phaser.Timer#start start} after adding all of the Events you require for this Timer.
func (self *Timer) LoopI(args ...interface{}) *TimerEvent{
    return &TimerEvent{self.Object.Call("loop", args)}
}

// Start Starts this Timer running.
func (self *Timer) Start() {
    self.Object.Call("start")
}

// Start1O Starts this Timer running.
func (self *Timer) Start1O(delay int) {
    self.Object.Call("start", delay)
}

// StartI Starts this Timer running.
func (self *Timer) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// Stop Stops this Timer from running. Does not cause it to be destroyed if autoDestroy is set to true.
func (self *Timer) Stop() {
    self.Object.Call("stop")
}

// Stop1O Stops this Timer from running. Does not cause it to be destroyed if autoDestroy is set to true.
func (self *Timer) Stop1O(clearEvents bool) {
    self.Object.Call("stop", clearEvents)
}

// StopI Stops this Timer from running. Does not cause it to be destroyed if autoDestroy is set to true.
func (self *Timer) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Remove Removes a pending TimerEvent from the queue.
func (self *Timer) Remove(event *TimerEvent) {
    self.Object.Call("remove", event)
}

// RemoveI Removes a pending TimerEvent from the queue.
func (self *Timer) RemoveI(args ...interface{}) {
    self.Object.Call("remove", args)
}

// Order Orders the events on this Timer so they are in tick order.
// This is called automatically when new events are created.
func (self *Timer) Order() {
    self.Object.Call("order")
}

// OrderI Orders the events on this Timer so they are in tick order.
// This is called automatically when new events are created.
func (self *Timer) OrderI(args ...interface{}) {
    self.Object.Call("order", args)
}

// SortHandler Sort handler used by Phaser.Timer.order.
func (self *Timer) SortHandler() {
    self.Object.Call("sortHandler")
}

// SortHandlerI Sort handler used by Phaser.Timer.order.
func (self *Timer) SortHandlerI(args ...interface{}) {
    self.Object.Call("sortHandler", args)
}

// ClearPendingEvents Clears any events from the Timer which have pendingDelete set to true and then resets the private _len and _i values.
func (self *Timer) ClearPendingEvents() {
    self.Object.Call("clearPendingEvents")
}

// ClearPendingEventsI Clears any events from the Timer which have pendingDelete set to true and then resets the private _len and _i values.
func (self *Timer) ClearPendingEventsI(args ...interface{}) {
    self.Object.Call("clearPendingEvents", args)
}

// Update The main Timer update event, called automatically by Phaser.Time.update.
func (self *Timer) Update(time int) bool{
    return self.Object.Call("update", time).Bool()
}

// UpdateI The main Timer update event, called automatically by Phaser.Time.update.
func (self *Timer) UpdateI(args ...interface{}) bool{
    return self.Object.Call("update", args).Bool()
}

// Pause Pauses the Timer and all events in the queue.
func (self *Timer) Pause() {
    self.Object.Call("pause")
}

// PauseI Pauses the Timer and all events in the queue.
func (self *Timer) PauseI(args ...interface{}) {
    self.Object.Call("pause", args)
}

// _pause Internal pause/resume control - user code should use Timer.pause instead.
func (self *Timer) _pause() {
    self.Object.Call("_pause")
}

// _pauseI Internal pause/resume control - user code should use Timer.pause instead.
func (self *Timer) _pauseI(args ...interface{}) {
    self.Object.Call("_pause", args)
}

// AdjustEvents Adjusts the time of all pending events and the nextTick by the given baseTime.
func (self *Timer) AdjustEvents() {
    self.Object.Call("adjustEvents")
}

// AdjustEventsI Adjusts the time of all pending events and the nextTick by the given baseTime.
func (self *Timer) AdjustEventsI(args ...interface{}) {
    self.Object.Call("adjustEvents", args)
}

// Resume Resumes the Timer and updates all pending events.
func (self *Timer) Resume() {
    self.Object.Call("resume")
}

// ResumeI Resumes the Timer and updates all pending events.
func (self *Timer) ResumeI(args ...interface{}) {
    self.Object.Call("resume", args)
}

// _resume Internal pause/resume control - user code should use Timer.resume instead.
func (self *Timer) _resume() {
    self.Object.Call("_resume")
}

// _resumeI Internal pause/resume control - user code should use Timer.resume instead.
func (self *Timer) _resumeI(args ...interface{}) {
    self.Object.Call("_resume", args)
}

// RemoveAll Removes all Events from this Timer and all callbacks linked to onComplete, but leaves the Timer running.    
// The onComplete callbacks won't be called.
func (self *Timer) RemoveAll() {
    self.Object.Call("removeAll")
}

// RemoveAllI Removes all Events from this Timer and all callbacks linked to onComplete, but leaves the Timer running.    
// The onComplete callbacks won't be called.
func (self *Timer) RemoveAllI(args ...interface{}) {
    self.Object.Call("removeAll", args)
}

// Destroy Destroys this Timer. Any pending Events are not dispatched.
// The onComplete callbacks won't be called.
func (self *Timer) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this Timer. Any pending Events are not dispatched.
// The onComplete callbacks won't be called.
func (self *Timer) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

