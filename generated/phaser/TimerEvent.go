// Automatic generation for Phaser.TimerEvent
// generated file TimerEvent.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A TimerEvent is a single event that is processed by a Phaser.Timer.
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


// The Timer object that this TimerEvent belongs to.
func (self *TimerEvent) GetTimerA() *Timer{
    return &Timer{self.Object.Get("timer")}
}

// The Timer object that this TimerEvent belongs to.
func (self *TimerEvent) SetTimerA(member *Timer) {
    self.Object.Set("timer", member)
}

// The delay in ms at which this TimerEvent fires.
func (self *TimerEvent) GetDelayA() int{
    return self.Object.Get("delay").Int()
}

// The delay in ms at which this TimerEvent fires.
func (self *TimerEvent) SetDelayA(member int) {
    self.Object.Set("delay", member)
}

// The tick is the next game clock time that this event will fire at.
func (self *TimerEvent) GetTickA() int{
    return self.Object.Get("tick").Int()
}

// The tick is the next game clock time that this event will fire at.
func (self *TimerEvent) SetTickA(member int) {
    self.Object.Set("tick", member)
}

// If this TimerEvent repeats it will do so this many times.
func (self *TimerEvent) GetRepeatCountA() int{
    return self.Object.Get("repeatCount").Int()
}

// If this TimerEvent repeats it will do so this many times.
func (self *TimerEvent) SetRepeatCountA(member int) {
    self.Object.Set("repeatCount", member)
}

// True if this TimerEvent loops, otherwise false.
func (self *TimerEvent) GetLoopA() bool{
    return self.Object.Get("loop").Bool()
}

// True if this TimerEvent loops, otherwise false.
func (self *TimerEvent) SetLoopA(member bool) {
    self.Object.Set("loop", member)
}

// The callback that will be called when the TimerEvent occurs.
func (self *TimerEvent) SetCallbackA(member func(...interface{})) {
    self.Object.Set("callback", member)
}

// The context in which the callback will be called.
func (self *TimerEvent) GetCallbackContextA() interface{}{
    return self.Object.Get("callbackContext")
}

// The context in which the callback will be called.
func (self *TimerEvent) SetCallbackContextA(member interface{}) {
    self.Object.Set("callbackContext", member)
}

// Additional arguments to be passed to the callback.
func (self *TimerEvent) GetArgsA() interface{}{
    return self.Object.Get("args")
}

// Additional arguments to be passed to the callback.
func (self *TimerEvent) SetArgsA(member interface{}) {
    self.Object.Set("args", member)
}

// A flag that controls if the TimerEvent is pending deletion.
func (self *TimerEvent) GetPendingDeleteA() bool{
    return self.Object.Get("pendingDelete").Bool()
}

// A flag that controls if the TimerEvent is pending deletion.
func (self *TimerEvent) SetPendingDeleteA(member bool) {
    self.Object.Set("pendingDelete", member)
}


