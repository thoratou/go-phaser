// Package phaser Automatic generation for PIXI.EventTarget
// generated file EventTarget.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// EventTarget Mixins event emitter functionality to a class
type EventTarget struct {
    *js.Object
}

// NewEventTarget Mixins event emitter functionality to a class
func NewEventTarget() *EventTarget {
    return &EventTarget{js.Global.Get("PIXI").Get("EventTarget").New()}
}
// NewEventTargetI Mixins event emitter functionality to a class
func NewEventTargetI(args ...interface{}) *EventTarget {
    return &EventTarget{js.Global.Get("PIXI").Get("EventTarget").New(args)}
}



// EventTarget Binding conversion method to EventTarget point 
func ToEventTarget(jsStruct interface{}) *EventTarget {
    if object, ok := jsStruct.(*js.Object); ok {
		return &EventTarget{Object: object}
	}
	return nil
}




// Mixin Mixes in the properties of the EventTarget prototype onto another object
func (self *EventTarget) Mixin(object interface{}) {
    self.Object.Call("mixin", object)
}

// MixinI Mixes in the properties of the EventTarget prototype onto another object
func (self *EventTarget) MixinI(args ...interface{}) {
    self.Object.Call("mixin", args)
}

// Listeners Return a list of assigned event listeners.
func (self *EventTarget) Listeners(eventName string) []interface{}{
	array00 := self.Object.Call("listeners", eventName)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ListenersI Return a list of assigned event listeners.
func (self *EventTarget) ListenersI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("listeners", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Emit Emit an event to all registered event listeners.
func (self *EventTarget) Emit(eventName string) bool{
    return self.Object.Call("emit", eventName).Bool()
}

// EmitI Emit an event to all registered event listeners.
func (self *EventTarget) EmitI(args ...interface{}) bool{
    return self.Object.Call("emit", args).Bool()
}

// On Register a new EventListener for the given event.
func (self *EventTarget) On(eventName string, callback interface{}) {
    self.Object.Call("on", eventName, callback)
}

// OnI Register a new EventListener for the given event.
func (self *EventTarget) OnI(args ...interface{}) {
    self.Object.Call("on", args)
}

// Once Add an EventListener that's only called once.
func (self *EventTarget) Once(eventName string, callback interface{}) {
    self.Object.Call("once", eventName, callback)
}

// OnceI Add an EventListener that's only called once.
func (self *EventTarget) OnceI(args ...interface{}) {
    self.Object.Call("once", args)
}

// Off Remove event listeners.
func (self *EventTarget) Off(eventName string, callback interface{}) {
    self.Object.Call("off", eventName, callback)
}

// OffI Remove event listeners.
func (self *EventTarget) OffI(args ...interface{}) {
    self.Object.Call("off", args)
}

// RemoveAllListeners Remove all listeners or only the listeners for the specified event.
func (self *EventTarget) RemoveAllListeners(eventName string) {
    self.Object.Call("removeAllListeners", eventName)
}

// RemoveAllListenersI Remove all listeners or only the listeners for the specified event.
func (self *EventTarget) RemoveAllListenersI(args ...interface{}) {
    self.Object.Call("removeAllListeners", args)
}

