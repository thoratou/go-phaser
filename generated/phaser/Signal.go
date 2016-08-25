// Package phaser Automatic generation for Phaser.Signal
// generated file Signal.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Signal Signals are what Phaser uses to handle events and event dispatching.
// You can listen for a Signal by binding a callback / function to it.
// This is done by using either `Signal.add` or `Signal.addOnce`.
// 
// For example you can listen for a touch or click event from the Input Manager 
// by using its `onDown` Signal:
// 
// `game.input.onDown.add(function() { ... });`
// 
// Rather than inline your function, you can pass a reference:
// 
// `game.input.onDown.add(clicked, this);`
// `function clicked () { ... }`
// 
// In this case the second argument (`this`) is the context in which your function should be called.
// 
// Now every time the InputManager dispatches the `onDown` signal (or event), your function
// will be called.
// 
// Very often a Signal will send arguments to your function.
// This is specific to the Signal itself.
// If you're unsure then check the documentation, or failing that simply do:
// 
// `Signal.add(function() { console.log(arguments); })`
// 
// and it will log all of the arguments your function received from the Signal.
// 
// Sprites have lots of default signals you can listen to in their Events class, such as:
// 
// `sprite.events.onKilled`
// 
// Which is called automatically whenever the Sprite is killed.
// There are lots of other events, see the Events component for a list.
// 
// As well as listening to pre-defined Signals you can also create your own:
// 
// `var mySignal = new Phaser.Signal();`
// 
// This creates a new Signal. You can bind a callback to it:
// 
// `mySignal.add(myCallback, this);`
// 
// and then finally when ready you can dispatch the Signal:
// 
// `mySignal.dispatch(your arguments);`
// 
// And your callback will be invoked. See the dispatch method for more details.
type Signal struct {
    *js.Object
}

// NewSignal Signals are what Phaser uses to handle events and event dispatching.
// You can listen for a Signal by binding a callback / function to it.
// This is done by using either `Signal.add` or `Signal.addOnce`.
// 
// For example you can listen for a touch or click event from the Input Manager 
// by using its `onDown` Signal:
// 
// `game.input.onDown.add(function() { ... });`
// 
// Rather than inline your function, you can pass a reference:
// 
// `game.input.onDown.add(clicked, this);`
// `function clicked () { ... }`
// 
// In this case the second argument (`this`) is the context in which your function should be called.
// 
// Now every time the InputManager dispatches the `onDown` signal (or event), your function
// will be called.
// 
// Very often a Signal will send arguments to your function.
// This is specific to the Signal itself.
// If you're unsure then check the documentation, or failing that simply do:
// 
// `Signal.add(function() { console.log(arguments); })`
// 
// and it will log all of the arguments your function received from the Signal.
// 
// Sprites have lots of default signals you can listen to in their Events class, such as:
// 
// `sprite.events.onKilled`
// 
// Which is called automatically whenever the Sprite is killed.
// There are lots of other events, see the Events component for a list.
// 
// As well as listening to pre-defined Signals you can also create your own:
// 
// `var mySignal = new Phaser.Signal();`
// 
// This creates a new Signal. You can bind a callback to it:
// 
// `mySignal.add(myCallback, this);`
// 
// and then finally when ready you can dispatch the Signal:
// 
// `mySignal.dispatch(your arguments);`
// 
// And your callback will be invoked. See the dispatch method for more details.
func NewSignal() *Signal {
    return &Signal{js.Global.Get("Phaser").Get("Signal").New()}
}
// NewSignalI Signals are what Phaser uses to handle events and event dispatching.
// You can listen for a Signal by binding a callback / function to it.
// This is done by using either `Signal.add` or `Signal.addOnce`.
// 
// For example you can listen for a touch or click event from the Input Manager 
// by using its `onDown` Signal:
// 
// `game.input.onDown.add(function() { ... });`
// 
// Rather than inline your function, you can pass a reference:
// 
// `game.input.onDown.add(clicked, this);`
// `function clicked () { ... }`
// 
// In this case the second argument (`this`) is the context in which your function should be called.
// 
// Now every time the InputManager dispatches the `onDown` signal (or event), your function
// will be called.
// 
// Very often a Signal will send arguments to your function.
// This is specific to the Signal itself.
// If you're unsure then check the documentation, or failing that simply do:
// 
// `Signal.add(function() { console.log(arguments); })`
// 
// and it will log all of the arguments your function received from the Signal.
// 
// Sprites have lots of default signals you can listen to in their Events class, such as:
// 
// `sprite.events.onKilled`
// 
// Which is called automatically whenever the Sprite is killed.
// There are lots of other events, see the Events component for a list.
// 
// As well as listening to pre-defined Signals you can also create your own:
// 
// `var mySignal = new Phaser.Signal();`
// 
// This creates a new Signal. You can bind a callback to it:
// 
// `mySignal.add(myCallback, this);`
// 
// and then finally when ready you can dispatch the Signal:
// 
// `mySignal.dispatch(your arguments);`
// 
// And your callback will be invoked. See the dispatch method for more details.
func NewSignalI(args ...interface{}) *Signal {
    return &Signal{js.Global.Get("Phaser").Get("Signal").New(args)}
}



// Signal Binding conversion method to Signal point 
func ToSignal(jsStruct interface{}) *Signal {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Signal{Object: object}
	}
	return nil
}



// Memorize Memorize the previously dispatched event?
// 
// If an event has been memorized it is automatically dispatched when a new listener is added with {@link Phaser.Signal#add add} or {@link Phaser.Signal#addOnce addOnce}.
// Use {@link Phaser.Signal#forget forget} to clear any currently memorized event.
func (self *Signal) Memorize() bool{
    return self.Object.Get("memorize").Bool()
}

// SetMemorizeA Memorize the previously dispatched event?
// 
// If an event has been memorized it is automatically dispatched when a new listener is added with {@link Phaser.Signal#add add} or {@link Phaser.Signal#addOnce addOnce}.
// Use {@link Phaser.Signal#forget forget} to clear any currently memorized event.
func (self *Signal) SetMemorizeA(member bool) {
    self.Object.Set("memorize", member)
}

// Active Is the Signal active? Only active signals will broadcast dispatched events.
// 
// Setting this property during a dispatch will only affect the next dispatch. To stop the propagation of a signal from a listener use {@link Phaser.Signal#halt halt}.
func (self *Signal) Active() bool{
    return self.Object.Get("active").Bool()
}

// SetActiveA Is the Signal active? Only active signals will broadcast dispatched events.
// 
// Setting this property during a dispatch will only affect the next dispatch. To stop the propagation of a signal from a listener use {@link Phaser.Signal#halt halt}.
func (self *Signal) SetActiveA(member bool) {
    self.Object.Set("active", member)
}


// ValidateListener empty description
func (self *Signal) ValidateListener(listener interface{}, fnName string) {
    self.Object.Call("validateListener", listener, fnName)
}

// ValidateListenerI empty description
func (self *Signal) ValidateListenerI(args ...interface{}) {
    self.Object.Call("validateListener", args)
}

// _registerListener empty description
func (self *Signal) _registerListener(listener interface{}, isOnce bool) *SignalBinding{
    return &SignalBinding{self.Object.Call("_registerListener", listener, isOnce)}
}

// _registerListener1O empty description
func (self *Signal) _registerListener1O(listener interface{}, isOnce bool, listenerContext interface{}) *SignalBinding{
    return &SignalBinding{self.Object.Call("_registerListener", listener, isOnce, listenerContext)}
}

// _registerListener2O empty description
func (self *Signal) _registerListener2O(listener interface{}, isOnce bool, listenerContext interface{}, priority int) *SignalBinding{
    return &SignalBinding{self.Object.Call("_registerListener", listener, isOnce, listenerContext, priority)}
}

// _registerListenerI empty description
func (self *Signal) _registerListenerI(args ...interface{}) *SignalBinding{
    return &SignalBinding{self.Object.Call("_registerListener", args)}
}

// _addBinding empty description
func (self *Signal) _addBinding(binding *SignalBinding) {
    self.Object.Call("_addBinding", binding)
}

// _addBindingI empty description
func (self *Signal) _addBindingI(args ...interface{}) {
    self.Object.Call("_addBinding", args)
}

// _indexOfListener empty description
func (self *Signal) _indexOfListener(listener interface{}) int{
    return self.Object.Call("_indexOfListener", listener).Int()
}

// _indexOfListener1O empty description
func (self *Signal) _indexOfListener1O(listener interface{}, context interface{}) int{
    return self.Object.Call("_indexOfListener", listener, context).Int()
}

// _indexOfListenerI empty description
func (self *Signal) _indexOfListenerI(args ...interface{}) int{
    return self.Object.Call("_indexOfListener", args).Int()
}

// Has Check if a specific listener is attached.
func (self *Signal) Has(listener interface{}) bool{
    return self.Object.Call("has", listener).Bool()
}

// Has1O Check if a specific listener is attached.
func (self *Signal) Has1O(listener interface{}, context interface{}) bool{
    return self.Object.Call("has", listener, context).Bool()
}

// HasI Check if a specific listener is attached.
func (self *Signal) HasI(args ...interface{}) bool{
    return self.Object.Call("has", args).Bool()
}

// Add Add an event listener for this signal.
// 
// An event listener is a callback with a related context and priority.
// 
// You can optionally provide extra arguments which will be passed to the callback after any internal parameters.
// 
// For example: `Phaser.Key.onDown` when dispatched will send the Phaser.Key object that caused the signal as the first parameter.
// Any arguments you've specified after `priority` will be sent as well:
// 
// `fireButton.onDown.add(shoot, this, 0, 'lazer', 100);`
// 
// When onDown dispatches it will call the `shoot` callback passing it: `Phaser.Key, 'lazer', 100`.
// 
// Where the first parameter is the one that Key.onDown dispatches internally and 'lazer', 
// and the value 100 were the custom arguments given in the call to 'add'.
func (self *Signal) Add(listener interface{}) *SignalBinding{
    return &SignalBinding{self.Object.Call("add", listener)}
}

// Add1O Add an event listener for this signal.
// 
// An event listener is a callback with a related context and priority.
// 
// You can optionally provide extra arguments which will be passed to the callback after any internal parameters.
// 
// For example: `Phaser.Key.onDown` when dispatched will send the Phaser.Key object that caused the signal as the first parameter.
// Any arguments you've specified after `priority` will be sent as well:
// 
// `fireButton.onDown.add(shoot, this, 0, 'lazer', 100);`
// 
// When onDown dispatches it will call the `shoot` callback passing it: `Phaser.Key, 'lazer', 100`.
// 
// Where the first parameter is the one that Key.onDown dispatches internally and 'lazer', 
// and the value 100 were the custom arguments given in the call to 'add'.
func (self *Signal) Add1O(listener interface{}, listenerContext interface{}) *SignalBinding{
    return &SignalBinding{self.Object.Call("add", listener, listenerContext)}
}

// Add2O Add an event listener for this signal.
// 
// An event listener is a callback with a related context and priority.
// 
// You can optionally provide extra arguments which will be passed to the callback after any internal parameters.
// 
// For example: `Phaser.Key.onDown` when dispatched will send the Phaser.Key object that caused the signal as the first parameter.
// Any arguments you've specified after `priority` will be sent as well:
// 
// `fireButton.onDown.add(shoot, this, 0, 'lazer', 100);`
// 
// When onDown dispatches it will call the `shoot` callback passing it: `Phaser.Key, 'lazer', 100`.
// 
// Where the first parameter is the one that Key.onDown dispatches internally and 'lazer', 
// and the value 100 were the custom arguments given in the call to 'add'.
func (self *Signal) Add2O(listener interface{}, listenerContext interface{}, priority int) *SignalBinding{
    return &SignalBinding{self.Object.Call("add", listener, listenerContext, priority)}
}

// Add3O Add an event listener for this signal.
// 
// An event listener is a callback with a related context and priority.
// 
// You can optionally provide extra arguments which will be passed to the callback after any internal parameters.
// 
// For example: `Phaser.Key.onDown` when dispatched will send the Phaser.Key object that caused the signal as the first parameter.
// Any arguments you've specified after `priority` will be sent as well:
// 
// `fireButton.onDown.add(shoot, this, 0, 'lazer', 100);`
// 
// When onDown dispatches it will call the `shoot` callback passing it: `Phaser.Key, 'lazer', 100`.
// 
// Where the first parameter is the one that Key.onDown dispatches internally and 'lazer', 
// and the value 100 were the custom arguments given in the call to 'add'.
func (self *Signal) Add3O(listener interface{}, listenerContext interface{}, priority int, args interface{}) *SignalBinding{
    return &SignalBinding{self.Object.Call("add", listener, listenerContext, priority, args)}
}

// AddI Add an event listener for this signal.
// 
// An event listener is a callback with a related context and priority.
// 
// You can optionally provide extra arguments which will be passed to the callback after any internal parameters.
// 
// For example: `Phaser.Key.onDown` when dispatched will send the Phaser.Key object that caused the signal as the first parameter.
// Any arguments you've specified after `priority` will be sent as well:
// 
// `fireButton.onDown.add(shoot, this, 0, 'lazer', 100);`
// 
// When onDown dispatches it will call the `shoot` callback passing it: `Phaser.Key, 'lazer', 100`.
// 
// Where the first parameter is the one that Key.onDown dispatches internally and 'lazer', 
// and the value 100 were the custom arguments given in the call to 'add'.
func (self *Signal) AddI(args ...interface{}) *SignalBinding{
    return &SignalBinding{self.Object.Call("add", args)}
}

// AddOnce Add a one-time listener - the listener is automatically removed after the first execution.
// 
// If there is as {@link Phaser.Signal#memorize memorized} event then it will be dispatched and
// the listener will be removed immediately.
func (self *Signal) AddOnce(listener interface{}) *SignalBinding{
    return &SignalBinding{self.Object.Call("addOnce", listener)}
}

// AddOnce1O Add a one-time listener - the listener is automatically removed after the first execution.
// 
// If there is as {@link Phaser.Signal#memorize memorized} event then it will be dispatched and
// the listener will be removed immediately.
func (self *Signal) AddOnce1O(listener interface{}, listenerContext interface{}) *SignalBinding{
    return &SignalBinding{self.Object.Call("addOnce", listener, listenerContext)}
}

// AddOnce2O Add a one-time listener - the listener is automatically removed after the first execution.
// 
// If there is as {@link Phaser.Signal#memorize memorized} event then it will be dispatched and
// the listener will be removed immediately.
func (self *Signal) AddOnce2O(listener interface{}, listenerContext interface{}, priority int) *SignalBinding{
    return &SignalBinding{self.Object.Call("addOnce", listener, listenerContext, priority)}
}

// AddOnce3O Add a one-time listener - the listener is automatically removed after the first execution.
// 
// If there is as {@link Phaser.Signal#memorize memorized} event then it will be dispatched and
// the listener will be removed immediately.
func (self *Signal) AddOnce3O(listener interface{}, listenerContext interface{}, priority int, args interface{}) *SignalBinding{
    return &SignalBinding{self.Object.Call("addOnce", listener, listenerContext, priority, args)}
}

// AddOnceI Add a one-time listener - the listener is automatically removed after the first execution.
// 
// If there is as {@link Phaser.Signal#memorize memorized} event then it will be dispatched and
// the listener will be removed immediately.
func (self *Signal) AddOnceI(args ...interface{}) *SignalBinding{
    return &SignalBinding{self.Object.Call("addOnce", args)}
}

// Remove Remove a single event listener.
func (self *Signal) Remove(listener interface{}) interface{}{
    return self.Object.Call("remove", listener)
}

// Remove1O Remove a single event listener.
func (self *Signal) Remove1O(listener interface{}, context interface{}) interface{}{
    return self.Object.Call("remove", listener, context)
}

// RemoveI Remove a single event listener.
func (self *Signal) RemoveI(args ...interface{}) interface{}{
    return self.Object.Call("remove", args)
}

// RemoveAll Remove all event listeners.
func (self *Signal) RemoveAll() {
    self.Object.Call("removeAll")
}

// RemoveAll1O Remove all event listeners.
func (self *Signal) RemoveAll1O(context interface{}) {
    self.Object.Call("removeAll", context)
}

// RemoveAllI Remove all event listeners.
func (self *Signal) RemoveAllI(args ...interface{}) {
    self.Object.Call("removeAll", args)
}

// GetNumListeners Gets the total number of listeners attached to this Signal.
func (self *Signal) GetNumListeners() int{
    return self.Object.Call("getNumListeners").Int()
}

// GetNumListenersI Gets the total number of listeners attached to this Signal.
func (self *Signal) GetNumListenersI(args ...interface{}) int{
    return self.Object.Call("getNumListeners", args).Int()
}

// Halt Stop propagation of the event, blocking the dispatch to next listener on the queue.
// 
// This should be called only during event dispatch as calling it before/after dispatch won't affect another broadcast.
// See {@link Phaser.Signal#active active} to enable/disable the signal entirely.
func (self *Signal) Halt() {
    self.Object.Call("halt")
}

// HaltI Stop propagation of the event, blocking the dispatch to next listener on the queue.
// 
// This should be called only during event dispatch as calling it before/after dispatch won't affect another broadcast.
// See {@link Phaser.Signal#active active} to enable/disable the signal entirely.
func (self *Signal) HaltI(args ...interface{}) {
    self.Object.Call("halt", args)
}

// Dispatch Dispatch / broadcast the event to all listeners.
// 
// To create an instance-bound dispatch for this Signal, use {@link Phaser.Signal#boundDispatch boundDispatch}.
func (self *Signal) Dispatch() {
    self.Object.Call("dispatch")
}

// Dispatch1O Dispatch / broadcast the event to all listeners.
// 
// To create an instance-bound dispatch for this Signal, use {@link Phaser.Signal#boundDispatch boundDispatch}.
func (self *Signal) Dispatch1O(params interface{}) {
    self.Object.Call("dispatch", params)
}

// DispatchI Dispatch / broadcast the event to all listeners.
// 
// To create an instance-bound dispatch for this Signal, use {@link Phaser.Signal#boundDispatch boundDispatch}.
func (self *Signal) DispatchI(args ...interface{}) {
    self.Object.Call("dispatch", args)
}

// Forget Forget the currently {@link Phaser.Signal#memorize memorized} event, if any.
func (self *Signal) Forget() {
    self.Object.Call("forget")
}

// ForgetI Forget the currently {@link Phaser.Signal#memorize memorized} event, if any.
func (self *Signal) ForgetI(args ...interface{}) {
    self.Object.Call("forget", args)
}

// Dispose Dispose the signal - no more events can be dispatched.
// 
// This removes all event listeners and clears references to external objects.
// Calling methods on a disposed objects results in undefined behavior.
func (self *Signal) Dispose() {
    self.Object.Call("dispose")
}

// DisposeI Dispose the signal - no more events can be dispatched.
// 
// This removes all event listeners and clears references to external objects.
// Calling methods on a disposed objects results in undefined behavior.
func (self *Signal) DisposeI(args ...interface{}) {
    self.Object.Call("dispose", args)
}

// ToString A string representation of the object.
func (self *Signal) ToString() string{
    return self.Object.Call("toString").String()
}

// ToStringI A string representation of the object.
func (self *Signal) ToStringI(args ...interface{}) string{
    return self.Object.Call("toString", args).String()
}

