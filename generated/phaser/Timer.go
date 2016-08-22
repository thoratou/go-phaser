// Automatic generation for Phaser.Timer
// generated file Timer.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A Timer is a way to create and manage {@link Phaser.TimerEvent timer events} that wait for a specific duration and then run a callback.
// Many different timer events, with individual delays, can be added to the same Timer.
// 
// All Timer delays are in milliseconds (there are 1000 ms in 1 second); so a delay value of 250 represents a quarter of a second.
// 
// Timers are based on real life time, adjusted for game pause durations.
// That is, *timer events are based on elapsed {@link Phaser.Time game time}* and do *not* take physics time or slow motion into account.
type Timer struct {
    *js.Object
}


// Local reference to game.
func (self *Timer) GetGame() *Game{
    return &Game{self.Get("game")}
}

// Local reference to game.
func (self *Timer) SetGame(member *Game) {
    self.Set("game", member)
}

// True if the Timer is actively running.
// 
// Do not modify this boolean - use {@link Phaser.Timer#pause pause} (and {@link Phaser.Timer#resume resume}) to pause the timer.
func (self *Timer) GetRunning() bool{
    return self.Get("running").Bool()
}

// True if the Timer is actively running.
// 
// Do not modify this boolean - use {@link Phaser.Timer#pause pause} (and {@link Phaser.Timer#resume resume}) to pause the timer.
func (self *Timer) SetRunning(member bool) {
    self.Set("running", member)
}

// If true, the timer will automatically destroy itself after all the events have been dispatched (assuming no looping events).
func (self *Timer) GetAutoDestroy() bool{
    return self.Get("autoDestroy").Bool()
}

// If true, the timer will automatically destroy itself after all the events have been dispatched (assuming no looping events).
func (self *Timer) SetAutoDestroy(member bool) {
    self.Set("autoDestroy", member)
}

// An expired Timer is one in which all of its events have been dispatched and none are pending.
func (self *Timer) GetExpired() bool{
    return self.Get("expired").Bool()
}

// An expired Timer is one in which all of its events have been dispatched and none are pending.
func (self *Timer) SetExpired(member bool) {
    self.Set("expired", member)
}

// Elapsed time since the last frame (in ms).
func (self *Timer) GetElapsed() int{
    return self.Get("elapsed").Int()
}

// Elapsed time since the last frame (in ms).
func (self *Timer) SetElapsed(member int) {
    self.Set("elapsed", member)
}

// An array holding all of this timers Phaser.TimerEvent objects. Use the methods add, repeat and loop to populate it.
func (self *Timer) GetEvents() []TimerEvent{
	array := self.Get("events")
	length := array.Length()
	out := make([]TimerEvent, length, length)
	for i := 0; i < length; i++ {
		out[i] = TimerEvent{array.Index(i)}
	}
	return out
}

// An array holding all of this timers Phaser.TimerEvent objects. Use the methods add, repeat and loop to populate it.
func (self *Timer) SetEvents(member []TimerEvent) {
    self.Set("events", member)
}

// This signal will be dispatched when this Timer has completed which means that there are no more events in the queue.
// 
// The signal is supplied with one argument, `timer`, which is this Timer object.
func (self *Timer) GetOnComplete() *Signal{
    return &Signal{self.Get("onComplete")}
}

// This signal will be dispatched when this Timer has completed which means that there are no more events in the queue.
// 
// The signal is supplied with one argument, `timer`, which is this Timer object.
func (self *Timer) SetOnComplete(member *Signal) {
    self.Set("onComplete", member)
}

// The time the next tick will occur.
func (self *Timer) GetNextTick() int{
    return self.Get("nextTick").Int()
}

// The time the next tick will occur.
func (self *Timer) SetNextTick(member int) {
    self.Set("nextTick", member)
}

// If the difference in time between two frame updates exceeds this value, the event times are reset to avoid catch-up situations.
func (self *Timer) GetTimeCap() int{
    return self.Get("timeCap").Int()
}

// If the difference in time between two frame updates exceeds this value, the event times are reset to avoid catch-up situations.
func (self *Timer) SetTimeCap(member int) {
    self.Set("timeCap", member)
}

// The paused state of the Timer. You can pause the timer by calling Timer.pause() and Timer.resume() or by the game pausing.
func (self *Timer) GetPaused() bool{
    return self.Get("paused").Bool()
}

// The paused state of the Timer. You can pause the timer by calling Timer.pause() and Timer.resume() or by the game pausing.
func (self *Timer) SetPaused(member bool) {
    self.Set("paused", member)
}

// Number of milliseconds in a minute.
func (self *Timer) GetMINUTE() int{
    return self.Get("MINUTE").Int()
}

// Number of milliseconds in a minute.
func (self *Timer) SetMINUTE(member int) {
    self.Set("MINUTE", member)
}

// Number of milliseconds in a second.
func (self *Timer) GetSECOND() int{
    return self.Get("SECOND").Int()
}

// Number of milliseconds in a second.
func (self *Timer) SetSECOND(member int) {
    self.Set("SECOND", member)
}

// Number of milliseconds in half a second.
func (self *Timer) GetHALF() int{
    return self.Get("HALF").Int()
}

// Number of milliseconds in half a second.
func (self *Timer) SetHALF(member int) {
    self.Set("HALF", member)
}

// Number of milliseconds in a quarter of a second.
func (self *Timer) GetQUARTER() int{
    return self.Get("QUARTER").Int()
}

// Number of milliseconds in a quarter of a second.
func (self *Timer) SetQUARTER(member int) {
    self.Set("QUARTER", member)
}

// The time at which the next event will occur.
func (self *Timer) GetNext() int{
    return self.Get("next").Int()
}

// The time at which the next event will occur.
func (self *Timer) SetNext(member int) {
    self.Set("next", member)
}

// The duration in ms remaining until the next event will occur.
func (self *Timer) GetDuration() int{
    return self.Get("duration").Int()
}

// The duration in ms remaining until the next event will occur.
func (self *Timer) SetDuration(member int) {
    self.Set("duration", member)
}

// The number of pending events in the queue.
func (self *Timer) GetLength() int{
    return self.Get("length").Int()
}

// The number of pending events in the queue.
func (self *Timer) SetLength(member int) {
    self.Set("length", member)
}

// The duration in milliseconds that this Timer has been running for.
func (self *Timer) GetMs() int{
    return self.Get("ms").Int()
}

// The duration in milliseconds that this Timer has been running for.
func (self *Timer) SetMs(member int) {
    self.Set("ms", member)
}

// The duration in seconds that this Timer has been running for.
func (self *Timer) GetSeconds() int{
    return self.Get("seconds").Int()
}

// The duration in seconds that this Timer has been running for.
func (self *Timer) SetSeconds(member int) {
    self.Set("seconds", member)
}



// Creates a new TimerEvent on this Timer.
// 
// Use {@link Phaser.Timer#add}, {@link Phaser.Timer#repeat}, or {@link Phaser.Timer#loop} methods to create a new event.
func (self *Timer) CreateI(args ...interface{}) *TimerEvent{
    return &TimerEvent{self.Call("create", args)}
}

// Adds a new Event to this Timer.
// 
// The event will fire after the given amount of `delay` in milliseconds has passed, once the Timer has started running.
// The delay is in relation to when the Timer starts, not the time it was added. If the Timer is already running the delay will be calculated based on the timers current time.
// 
// Make sure to call {@link Phaser.Timer#start start} after adding all of the Events you require for this Timer.
func (self *Timer) AddI(args ...interface{}) *TimerEvent{
    return &TimerEvent{self.Call("add", args)}
}

// Adds a new TimerEvent that will always play through once and then repeat for the given number of iterations.
// 
// The event will fire after the given amount of `delay` in milliseconds has passed, once the Timer has started running.
// The delay is in relation to when the Timer starts, not the time it was added.
// If the Timer is already running the delay will be calculated based on the timers current time.
// 
// Make sure to call {@link Phaser.Timer#start start} after adding all of the Events you require for this Timer.
func (self *Timer) RepeatI(args ...interface{}) *TimerEvent{
    return &TimerEvent{self.Call("repeat", args)}
}

// Adds a new looped Event to this Timer that will repeat forever or until the Timer is stopped.
// 
// The event will fire after the given amount of `delay` in milliseconds has passed, once the Timer has started running.
// The delay is in relation to when the Timer starts, not the time it was added. If the Timer is already running the delay will be calculated based on the timers current time.
// 
// Make sure to call {@link Phaser.Timer#start start} after adding all of the Events you require for this Timer.
func (self *Timer) LoopI(args ...interface{}) *TimerEvent{
    return &TimerEvent{self.Call("loop", args)}
}

// Starts this Timer running.
func (self *Timer) StartI(args ...interface{}) {
    self.Call("start", args)
}

// Stops this Timer from running. Does not cause it to be destroyed if autoDestroy is set to true.
func (self *Timer) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// Removes a pending TimerEvent from the queue.
func (self *Timer) RemoveI(args ...interface{}) {
    self.Call("remove", args)
}

// Orders the events on this Timer so they are in tick order.
// This is called automatically when new events are created.
func (self *Timer) OrderI(args ...interface{}) {
    self.Call("order", args)
}

// Sort handler used by Phaser.Timer.order.
func (self *Timer) SortHandlerI(args ...interface{}) {
    self.Call("sortHandler", args)
}

// Clears any events from the Timer which have pendingDelete set to true and then resets the private _len and _i values.
func (self *Timer) ClearPendingEventsI(args ...interface{}) {
    self.Call("clearPendingEvents", args)
}

// The main Timer update event, called automatically by Phaser.Time.update.
func (self *Timer) UpdateI(args ...interface{}) bool{
    return self.Call("update", args).Bool()
}

// Pauses the Timer and all events in the queue.
func (self *Timer) PauseI(args ...interface{}) {
    self.Call("pause", args)
}

// Internal pause/resume control - user code should use Timer.pause instead.
func (self *Timer) _pauseI(args ...interface{}) {
    self.Call("_pause", args)
}

// Adjusts the time of all pending events and the nextTick by the given baseTime.
func (self *Timer) AdjustEventsI(args ...interface{}) {
    self.Call("adjustEvents", args)
}

// Resumes the Timer and updates all pending events.
func (self *Timer) ResumeI(args ...interface{}) {
    self.Call("resume", args)
}

// Internal pause/resume control - user code should use Timer.resume instead.
func (self *Timer) _resumeI(args ...interface{}) {
    self.Call("_resume", args)
}

// Removes all Events from this Timer and all callbacks linked to onComplete, but leaves the Timer running.    
// The onComplete callbacks won't be called.
func (self *Timer) RemoveAllI(args ...interface{}) {
    self.Call("removeAll", args)
}

// Destroys this Timer. Any pending Events are not dispatched.
// The onComplete callbacks won't be called.
func (self *Timer) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
