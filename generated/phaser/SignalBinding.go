// Package phaser Automatic generation for Phaser.SignalBinding
// generated file SignalBinding.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// SignalBinding Object that represents a binding between a Signal and a listener function.
// This is an internal constructor and shouldn't be created directly.
// Inspired by Joa Ebert AS3 SignalBinding and Robert Penner's Slot classes.
type SignalBinding struct {
    *js.Object
}

// NewSignalBinding Object that represents a binding between a Signal and a listener function.
// This is an internal constructor and shouldn't be created directly.
// Inspired by Joa Ebert AS3 SignalBinding and Robert Penner's Slot classes.
func NewSignalBinding(signal *Signal, listener interface{}, isOnce bool) *SignalBinding {
    return &SignalBinding{js.Global.Get("Phaser").Get("SignalBinding").New(signal, listener, isOnce)}
}
// NewSignalBinding1O Object that represents a binding between a Signal and a listener function.
// This is an internal constructor and shouldn't be created directly.
// Inspired by Joa Ebert AS3 SignalBinding and Robert Penner's Slot classes.
func NewSignalBinding1O(signal *Signal, listener interface{}, isOnce bool, listenerContext interface{}) *SignalBinding {
    return &SignalBinding{js.Global.Get("Phaser").Get("SignalBinding").New(signal, listener, isOnce, listenerContext)}
}
// NewSignalBinding2O Object that represents a binding between a Signal and a listener function.
// This is an internal constructor and shouldn't be created directly.
// Inspired by Joa Ebert AS3 SignalBinding and Robert Penner's Slot classes.
func NewSignalBinding2O(signal *Signal, listener interface{}, isOnce bool, listenerContext interface{}, priority int) *SignalBinding {
    return &SignalBinding{js.Global.Get("Phaser").Get("SignalBinding").New(signal, listener, isOnce, listenerContext, priority)}
}
// NewSignalBinding3O Object that represents a binding between a Signal and a listener function.
// This is an internal constructor and shouldn't be created directly.
// Inspired by Joa Ebert AS3 SignalBinding and Robert Penner's Slot classes.
func NewSignalBinding3O(signal *Signal, listener interface{}, isOnce bool, listenerContext interface{}, priority int, args interface{}) *SignalBinding {
    return &SignalBinding{js.Global.Get("Phaser").Get("SignalBinding").New(signal, listener, isOnce, listenerContext, priority, args)}
}
// NewSignalBindingI Object that represents a binding between a Signal and a listener function.
// This is an internal constructor and shouldn't be created directly.
// Inspired by Joa Ebert AS3 SignalBinding and Robert Penner's Slot classes.
func NewSignalBindingI(args ...interface{}) *SignalBinding {
    return &SignalBinding{js.Global.Get("Phaser").Get("SignalBinding").New(args)}
}



// SignalBinding Binding conversion method to SignalBinding point 
func ToSignalBinding(jsStruct interface{}) *SignalBinding {
    if object, ok := jsStruct.(*js.Object); ok {
		return &SignalBinding{Object: object}
	}
	return nil
}



// Context Context on which listener will be executed (object that should represent the `this` variable inside listener function).
func (self *SignalBinding) Context() interface{}{
    return self.Object.Get("context")
}

// SetContextA Context on which listener will be executed (object that should represent the `this` variable inside listener function).
func (self *SignalBinding) SetContextA(member interface{}) {
    self.Object.Set("context", member)
}

// CallCount The number of times the handler function has been called.
func (self *SignalBinding) CallCount() int{
    return self.Object.Get("callCount").Int()
}

// SetCallCountA The number of times the handler function has been called.
func (self *SignalBinding) SetCallCountA(member int) {
    self.Object.Set("callCount", member)
}

// Active If binding is active and should be executed.
func (self *SignalBinding) Active() bool{
    return self.Object.Get("active").Bool()
}

// SetActiveA If binding is active and should be executed.
func (self *SignalBinding) SetActiveA(member bool) {
    self.Object.Set("active", member)
}

// Params Default parameters passed to listener during `Signal.dispatch` and `SignalBinding.execute` (curried parameters).
func (self *SignalBinding) Params() interface{}{
    return self.Object.Get("params")
}

// SetParamsA Default parameters passed to listener during `Signal.dispatch` and `SignalBinding.execute` (curried parameters).
func (self *SignalBinding) SetParamsA(member interface{}) {
    self.Object.Set("params", member)
}


// Execute Call listener passing arbitrary parameters.
// If binding was added using `Signal.addOnce()` it will be automatically removed from signal dispatch queue, this method is used internally for the signal dispatch.
func (self *SignalBinding) Execute() interface{}{
    return self.Object.Call("execute")
}

// Execute1O Call listener passing arbitrary parameters.
// If binding was added using `Signal.addOnce()` it will be automatically removed from signal dispatch queue, this method is used internally for the signal dispatch.
func (self *SignalBinding) Execute1O(paramsArr []interface{}) interface{}{
    return self.Object.Call("execute", paramsArr)
}

// ExecuteI Call listener passing arbitrary parameters.
// If binding was added using `Signal.addOnce()` it will be automatically removed from signal dispatch queue, this method is used internally for the signal dispatch.
func (self *SignalBinding) ExecuteI(args ...interface{}) interface{}{
    return self.Object.Call("execute", args)
}

// Detach Detach binding from signal.
// alias to: @see mySignal.remove(myBinding.getListener());
func (self *SignalBinding) Detach() interface{}{
    return self.Object.Call("detach")
}

// DetachI Detach binding from signal.
// alias to: @see mySignal.remove(myBinding.getListener());
func (self *SignalBinding) DetachI(args ...interface{}) interface{}{
    return self.Object.Call("detach", args)
}

// IsBound empty description
func (self *SignalBinding) IsBound() bool{
    return self.Object.Call("isBound").Bool()
}

// IsBoundI empty description
func (self *SignalBinding) IsBoundI(args ...interface{}) bool{
    return self.Object.Call("isBound", args).Bool()
}

// IsOnce empty description
func (self *SignalBinding) IsOnce() bool{
    return self.Object.Call("isOnce").Bool()
}

// IsOnceI empty description
func (self *SignalBinding) IsOnceI(args ...interface{}) bool{
    return self.Object.Call("isOnce", args).Bool()
}

// GetListener empty description
func (self *SignalBinding) GetListener() interface{}{
    return self.Object.Call("getListener")
}

// GetListenerI empty description
func (self *SignalBinding) GetListenerI(args ...interface{}) interface{}{
    return self.Object.Call("getListener", args)
}

// GetSignal empty description
func (self *SignalBinding) GetSignal() *Signal{
    return &Signal{self.Object.Call("getSignal")}
}

// GetSignalI empty description
func (self *SignalBinding) GetSignalI(args ...interface{}) *Signal{
    return &Signal{self.Object.Call("getSignal", args)}
}

// _destroy Delete instance properties
func (self *SignalBinding) _destroy() {
    self.Object.Call("_destroy")
}

// _destroyI Delete instance properties
func (self *SignalBinding) _destroyI(args ...interface{}) {
    self.Object.Call("_destroy", args)
}

// ToString empty description
func (self *SignalBinding) ToString() string{
    return self.Object.Call("toString").String()
}

// ToStringI empty description
func (self *SignalBinding) ToStringI(args ...interface{}) string{
    return self.Object.Call("toString", args).String()
}

