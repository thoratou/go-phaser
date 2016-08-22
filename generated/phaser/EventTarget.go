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
func (self *EventTarget) Mixin(object interface{}) {
    self.Object.Call("mixin", object)
}

// Mixes in the properties of the EventTarget prototype onto another object
func (self *EventTarget) MixinI(args ...interface{}) {
    self.Object.Call("mixin", args)
}

// Return a list of assigned event listeners.
func (self *EventTarget) Listeners(eventName string) []interface{}{
	array00 := self.Object.Call("listeners", eventName)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Return a list of assigned event listeners.
func (self *EventTarget) ListenersI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("listeners", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Emit an event to all registered event listeners.
func (self *EventTarget) Emit(eventName string) bool{
    return self.Object.Call("emit", eventName).Bool()
}

// Emit an event to all registered event listeners.
func (self *EventTarget) EmitI(args ...interface{}) bool{
    return self.Object.Call("emit", args).Bool()
}

// Register a new EventListener for the given event.
func (self *EventTarget) On(eventName string, callback func(...interface{})) {
    self.Object.Call("on", eventName, callback)
}

// Register a new EventListener for the given event.
func (self *EventTarget) OnI(args ...interface{}) {
    self.Object.Call("on", args)
}

// Add an EventListener that's only called once.
func (self *EventTarget) Once(eventName string, callback func(...interface{})) {
    self.Object.Call("once", eventName, callback)
}

// Add an EventListener that's only called once.
func (self *EventTarget) OnceI(args ...interface{}) {
    self.Object.Call("once", args)
}

// Remove event listeners.
func (self *EventTarget) Off(eventName string, callback func(...interface{})) {
    self.Object.Call("off", eventName, callback)
}

// Remove event listeners.
func (self *EventTarget) OffI(args ...interface{}) {
    self.Object.Call("off", args)
}

// Remove all listeners or only the listeners for the specified event.
func (self *EventTarget) RemoveAllListeners(eventName string) {
    self.Object.Call("removeAllListeners", eventName)
}

// Remove all listeners or only the listeners for the specified event.
func (self *EventTarget) RemoveAllListenersI(args ...interface{}) {
    self.Object.Call("removeAllListeners", args)
}
