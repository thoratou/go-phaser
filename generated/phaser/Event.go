// Automatic generation for PIXI.Event
// generated file Event.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Creates an homogenous object for tracking events so users can know what to expect.
type Event struct {
    *js.Object
}


// The original target the event triggered on.
func (self *Event) GetTarget() interface{}{
    return self.Get("target")
}

// The original target the event triggered on.
func (self *Event) SetTarget(member interface{}) {
    self.Set("target", member)
}

// The string name of the event that this represents.
func (self *Event) GetType() string{
    return self.Get("type").String()
}

// The string name of the event that this represents.
func (self *Event) SetType(member string) {
    self.Set("type", member)
}

// The data that was passed in with this event.
func (self *Event) GetData() interface{}{
    return self.Get("data")
}

// The data that was passed in with this event.
func (self *Event) SetData(member interface{}) {
    self.Set("data", member)
}

// The timestamp when the event occurred.
func (self *Event) GetTimeStamp() int{
    return self.Get("timeStamp").Int()
}

// The timestamp when the event occurred.
func (self *Event) SetTimeStamp(member int) {
    self.Set("timeStamp", member)
}



// Stops the propagation of events up the scene graph (prevents bubbling).
func (self *Event) StopPropagationI(args ...interface{}) {
    self.Call("stopPropagation", args)
}

// Stops the propagation of events to sibling listeners (no longer calls any listeners).
func (self *Event) StopImmediatePropagationI(args ...interface{}) {
    self.Call("stopImmediatePropagation", args)
}
