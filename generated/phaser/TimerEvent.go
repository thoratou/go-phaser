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
func (self *TimerEvent) GetTimer() Timer{
    return Timer{self.Get("timer")}
}

// The Timer object that this TimerEvent belongs to.
func (self *TimerEvent) SetTimer(member Timer) {
    self.Set("timer", member)
}

// The delay in ms at which this TimerEvent fires.
func (self *TimerEvent) GetDelay() float64{
    return self.Get("delay").Float()
}

// The delay in ms at which this TimerEvent fires.
func (self *TimerEvent) SetDelay(member float64) {
    self.Set("delay", member)
}

// The tick is the next game clock time that this event will fire at.
func (self *TimerEvent) GetTick() float64{
    return self.Get("tick").Float()
}

// The tick is the next game clock time that this event will fire at.
func (self *TimerEvent) SetTick(member float64) {
    self.Set("tick", member)
}

// If this TimerEvent repeats it will do so this many times.
func (self *TimerEvent) GetRepeatCount() float64{
    return self.Get("repeatCount").Float()
}

// If this TimerEvent repeats it will do so this many times.
func (self *TimerEvent) SetRepeatCount(member float64) {
    self.Set("repeatCount", member)
}

// True if this TimerEvent loops, otherwise false.
func (self *TimerEvent) GetLoop() bool{
    return self.Get("loop").Bool()
}

// True if this TimerEvent loops, otherwise false.
func (self *TimerEvent) SetLoop(member bool) {
    self.Set("loop", member)
}

// The callback that will be called when the TimerEvent occurs.
func (self *TimerEvent) SetCallback(member func(...interface{})) {
    self.Set("callback", member)
}

// The context in which the callback will be called.
func (self *TimerEvent) GetCallbackContext() interface{}{
    return self.Get("callbackContext")
}

// The context in which the callback will be called.
func (self *TimerEvent) SetCallbackContext(member interface{}) {
    self.Set("callbackContext", member)
}

// Additional arguments to be passed to the callback.
func (self *TimerEvent) GetArgs() interface{}{
    return self.Get("args")
}

// Additional arguments to be passed to the callback.
func (self *TimerEvent) SetArgs(member interface{}) {
    self.Set("args", member)
}

// A flag that controls if the TimerEvent is pending deletion.
func (self *TimerEvent) GetPendingDelete() bool{
    return self.Get("pendingDelete").Bool()
}

// A flag that controls if the TimerEvent is pending deletion.
func (self *TimerEvent) SetPendingDelete(member bool) {
    self.Set("pendingDelete", member)
}


