// Package phaser Automatic generation for Phaser.TimerEvent
// generated file TimerEvent.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// TimerEvent A TimerEvent is a single event that is processed by a Phaser.Timer.
// 
// It consists of a delay, which is a value in milliseconds after which the event will fire.
// When the event fires it calls a specific callback with the specified arguments.
// 
// TimerEvents are removed by their parent timer once finished firing or repeating.
// 
// Use {@link Phaser.Timer#add}, {@link Phaser.Timer#repeat}, or {@link Phaser.Timer#loop} methods to create a new event.
type TimerEvent struct {
    *js.Object
}

// NewTimerEvent A TimerEvent is a single event that is processed by a Phaser.Timer.
// 
// It consists of a delay, which is a value in milliseconds after which the event will fire.
// When the event fires it calls a specific callback with the specified arguments.
// 
// TimerEvents are removed by their parent timer once finished firing or repeating.
// 
// Use {@link Phaser.Timer#add}, {@link Phaser.Timer#repeat}, or {@link Phaser.Timer#loop} methods to create a new event.
func NewTimerEvent(timer *Timer, delay int, tick int, repeatCount int, loop bool, callback interface{}, callbackContext interface{}, arguments []interface{}) *TimerEvent {
    return &TimerEvent{js.Global.Get("Phaser").Get("TimerEvent").New(timer, delay, tick, repeatCount, loop, callback, callbackContext, arguments)}
}
// NewTimerEventI A TimerEvent is a single event that is processed by a Phaser.Timer.
// 
// It consists of a delay, which is a value in milliseconds after which the event will fire.
// When the event fires it calls a specific callback with the specified arguments.
// 
// TimerEvents are removed by their parent timer once finished firing or repeating.
// 
// Use {@link Phaser.Timer#add}, {@link Phaser.Timer#repeat}, or {@link Phaser.Timer#loop} methods to create a new event.
func NewTimerEventI(args ...interface{}) *TimerEvent {
    return &TimerEvent{js.Global.Get("Phaser").Get("TimerEvent").New(args)}
}



// Timer The Timer object that this TimerEvent belongs to.
func (self *TimerEvent) Timer() *Timer{
    return &Timer{self.Object.Get("timer")}
}

// SetTimerA The Timer object that this TimerEvent belongs to.
func (self *TimerEvent) SetTimerA(member *Timer) {
    self.Object.Set("timer", member)
}

// Delay The delay in ms at which this TimerEvent fires.
func (self *TimerEvent) Delay() int{
    return self.Object.Get("delay").Int()
}

// SetDelayA The delay in ms at which this TimerEvent fires.
func (self *TimerEvent) SetDelayA(member int) {
    self.Object.Set("delay", member)
}

// Tick The tick is the next game clock time that this event will fire at.
func (self *TimerEvent) Tick() int{
    return self.Object.Get("tick").Int()
}

// SetTickA The tick is the next game clock time that this event will fire at.
func (self *TimerEvent) SetTickA(member int) {
    self.Object.Set("tick", member)
}

// RepeatCount If this TimerEvent repeats it will do so this many times.
func (self *TimerEvent) RepeatCount() int{
    return self.Object.Get("repeatCount").Int()
}

// SetRepeatCountA If this TimerEvent repeats it will do so this many times.
func (self *TimerEvent) SetRepeatCountA(member int) {
    self.Object.Set("repeatCount", member)
}

// Loop True if this TimerEvent loops, otherwise false.
func (self *TimerEvent) Loop() bool{
    return self.Object.Get("loop").Bool()
}

// SetLoopA True if this TimerEvent loops, otherwise false.
func (self *TimerEvent) SetLoopA(member bool) {
    self.Object.Set("loop", member)
}

// Callback The callback that will be called when the TimerEvent occurs.
func (self *TimerEvent) Callback() interface{}{
    return self.Object.Get("callback")
}

// SetCallbackA The callback that will be called when the TimerEvent occurs.
func (self *TimerEvent) SetCallbackA(member interface{}) {
    self.Object.Set("callback", member)
}

// CallbackContext The context in which the callback will be called.
func (self *TimerEvent) CallbackContext() interface{}{
    return self.Object.Get("callbackContext")
}

// SetCallbackContextA The context in which the callback will be called.
func (self *TimerEvent) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// Args Additional arguments to be passed to the callback.
func (self *TimerEvent) Args() interface{}{
    return self.Object.Get("args")
}

// SetArgsA Additional arguments to be passed to the callback.
func (self *TimerEvent) SetArgsA(member interface{}) {
    self.Object.Set("args", member)
}

// PendingDelete A flag that controls if the TimerEvent is pending deletion.
func (self *TimerEvent) PendingDelete() bool{
    return self.Object.Get("pendingDelete").Bool()
}

// SetPendingDeleteA A flag that controls if the TimerEvent is pending deletion.
func (self *TimerEvent) SetPendingDeleteA(member bool) {
    self.Object.Set("pendingDelete", member)
}


