// Automatic generation for Phaser.SignalBinding
// generated file SignalBinding.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Object that represents a binding between a Signal and a listener function.
// This is an internal constructor and shouldn't be created directly.
// Inspired by Joa Ebert AS3 SignalBinding and Robert Penner's Slot classes.
type SignalBinding struct {
    *js.Object
}


// Context on which listener will be executed (object that should represent the `this` variable inside listener function).
func (self *SignalBinding) GetContext() interface{}{
    return self.Get("context")
}

// Context on which listener will be executed (object that should represent the `this` variable inside listener function).
func (self *SignalBinding) SetContext(member interface{}) {
    self.Set("context", member)
}

// The number of times the handler function has been called.
func (self *SignalBinding) GetCallCount() float64{
    return self.Get("callCount").Float()
}

// The number of times the handler function has been called.
func (self *SignalBinding) SetCallCount(member float64) {
    self.Set("callCount", member)
}

// If binding is active and should be executed.
func (self *SignalBinding) GetActive() bool{
    return self.Get("active").Bool()
}

// If binding is active and should be executed.
func (self *SignalBinding) SetActive(member bool) {
    self.Set("active", member)
}

// Default parameters passed to listener during `Signal.dispatch` and `SignalBinding.execute` (curried parameters).
func (self *SignalBinding) GetParams() interface{}{
    return self.Get("params")
}

// Default parameters passed to listener during `Signal.dispatch` and `SignalBinding.execute` (curried parameters).
func (self *SignalBinding) SetParams(member interface{}) {
    self.Set("params", member)
}



// Call listener passing arbitrary parameters.
// If binding was added using `Signal.addOnce()` it will be automatically removed from signal dispatch queue, this method is used internally for the signal dispatch.
func (self *SignalBinding) ExecuteI(args ...interface{}) interface{}{
    return self.Call("execute", args)
}

// Detach binding from signal.
// alias to: @see mySignal.remove(myBinding.getListener());
func (self *SignalBinding) DetachI(args ...interface{}) interface{}{
    return self.Call("detach", args)
}

// 
func (self *SignalBinding) IsBoundI(args ...interface{}) bool{
    return self.Call("isBound", args).Bool()
}

// 
func (self *SignalBinding) IsOnceI(args ...interface{}) bool{
    return self.Call("isOnce", args).Bool()
}

// 
func (self *SignalBinding) GetListenerI(args ...interface{}) func(...interface{}){
    return func(...interface{}){self.Call("getListener", args)}
}

// 
func (self *SignalBinding) GetSignalI(args ...interface{}) Signal{
    return Signal{self.Call("getSignal", args)}
}

// Delete instance properties
func (self *SignalBinding) _destroyI(args ...interface{}) {
    self.Call("_destroy", args)
}

// 
func (self *SignalBinding) ToStringI(args ...interface{}) string{
    return self.Call("toString", args).String()
}
