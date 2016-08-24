// Package phaser Automatic generation for PIXI.Event
// generated file Event.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Event Creates an homogenous object for tracking events so users can know what to expect.
type Event struct {
    *js.Object
}

// NewEvent Creates an homogenous object for tracking events so users can know what to expect.
func NewEvent(target interface{}, name string, data interface{}) *Event {
    return &Event{js.Global.Get("PIXI").Get("Event").New(target, name, data)}
}
// NewEventI Creates an homogenous object for tracking events so users can know what to expect.
func NewEventI(args ...interface{}) *Event {
    return &Event{js.Global.Get("PIXI").Get("Event").New(args)}
}



// Target The original target the event triggered on.
func (self *Event) Target() interface{}{
    return self.Object.Get("target")
}

// SetTargetA The original target the event triggered on.
func (self *Event) SetTargetA(member interface{}) {
    self.Object.Set("target", member)
}

// Type The string name of the event that this represents.
func (self *Event) Type() string{
    return self.Object.Get("type").String()
}

// SetTypeA The string name of the event that this represents.
func (self *Event) SetTypeA(member string) {
    self.Object.Set("type", member)
}

// Data The data that was passed in with this event.
func (self *Event) Data() interface{}{
    return self.Object.Get("data")
}

// SetDataA The data that was passed in with this event.
func (self *Event) SetDataA(member interface{}) {
    self.Object.Set("data", member)
}

// TimeStamp The timestamp when the event occurred.
func (self *Event) TimeStamp() int{
    return self.Object.Get("timeStamp").Int()
}

// SetTimeStampA The timestamp when the event occurred.
func (self *Event) SetTimeStampA(member int) {
    self.Object.Set("timeStamp", member)
}


// StopPropagation Stops the propagation of events up the scene graph (prevents bubbling).
func (self *Event) StopPropagation() {
    self.Object.Call("stopPropagation")
}

// StopPropagationI Stops the propagation of events up the scene graph (prevents bubbling).
func (self *Event) StopPropagationI(args ...interface{}) {
    self.Object.Call("stopPropagation", args)
}

// StopImmediatePropagation Stops the propagation of events to sibling listeners (no longer calls any listeners).
func (self *Event) StopImmediatePropagation() {
    self.Object.Call("stopImmediatePropagation")
}

// StopImmediatePropagationI Stops the propagation of events to sibling listeners (no longer calls any listeners).
func (self *Event) StopImmediatePropagationI(args ...interface{}) {
    self.Object.Call("stopImmediatePropagation", args)
}

