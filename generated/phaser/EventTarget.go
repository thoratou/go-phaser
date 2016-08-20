// Automatic generation for PIXI.EventTarget
// generated file EventTarget.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Mixins event emitter functionality to a class
type EventTarget struct {
    *js.Object
}




// Mixes in the properties of the EventTarget prototype onto another object
func (self *EventTarget) MixinI(args ...interface{}) {
    self.Call("mixin", args)
}

// Return a list of assigned event listeners.
func (self *EventTarget) ListenersI(args ...interface{}) []interface{}{
	array := self.Call("listeners", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Emit an event to all registered event listeners.
func (self *EventTarget) EmitI(args ...interface{}) bool{
    return self.Call("emit", args).Bool()
}

// Register a new EventListener for the given event.
func (self *EventTarget) OnI(args ...interface{}) {
    self.Call("on", args)
}

// Add an EventListener that's only called once.
func (self *EventTarget) OnceI(args ...interface{}) {
    self.Call("once", args)
}

// Remove event listeners.
func (self *EventTarget) OffI(args ...interface{}) {
    self.Call("off", args)
}

// Remove all listeners or only the listeners for the specified event.
func (self *EventTarget) RemoveAllListenersI(args ...interface{}) {
    self.Call("removeAllListeners", args)
}
