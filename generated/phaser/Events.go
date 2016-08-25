// Package phaser Automatic generation for Phaser.Events
// generated file Events.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Events The Events component is a collection of events fired by the parent Game Object.
// 
// Phaser uses what are known as 'Signals' for all event handling. All of the events in
// this class are signals you can subscribe to, much in the same way you'd "listen" for
// an event.
// 
// For example to tell when a Sprite has been added to a new group, you can bind a function
// to the `onAddedToGroup` signal:
// 
// `sprite.events.onAddedToGroup.add(yourFunction, this);`
// 
// Where `yourFunction` is the function you want called when this event occurs.
// 
// For more details about how signals work please see the Phaser.Signal class.
// 
// The Input-related events will only be dispatched if the Sprite has had `inputEnabled` set to `true`
// and the Animation-related events only apply to game objects with animations like {@link Phaser.Sprite}.
type Events struct {
    *js.Object
}

// NewEvents The Events component is a collection of events fired by the parent Game Object.
// 
// Phaser uses what are known as 'Signals' for all event handling. All of the events in
// this class are signals you can subscribe to, much in the same way you'd "listen" for
// an event.
// 
// For example to tell when a Sprite has been added to a new group, you can bind a function
// to the `onAddedToGroup` signal:
// 
// `sprite.events.onAddedToGroup.add(yourFunction, this);`
// 
// Where `yourFunction` is the function you want called when this event occurs.
// 
// For more details about how signals work please see the Phaser.Signal class.
// 
// The Input-related events will only be dispatched if the Sprite has had `inputEnabled` set to `true`
// and the Animation-related events only apply to game objects with animations like {@link Phaser.Sprite}.
func NewEvents(sprite *Sprite) *Events {
    return &Events{js.Global.Get("Phaser").Get("Events").New(sprite)}
}
// NewEventsI The Events component is a collection of events fired by the parent Game Object.
// 
// Phaser uses what are known as 'Signals' for all event handling. All of the events in
// this class are signals you can subscribe to, much in the same way you'd "listen" for
// an event.
// 
// For example to tell when a Sprite has been added to a new group, you can bind a function
// to the `onAddedToGroup` signal:
// 
// `sprite.events.onAddedToGroup.add(yourFunction, this);`
// 
// Where `yourFunction` is the function you want called when this event occurs.
// 
// For more details about how signals work please see the Phaser.Signal class.
// 
// The Input-related events will only be dispatched if the Sprite has had `inputEnabled` set to `true`
// and the Animation-related events only apply to game objects with animations like {@link Phaser.Sprite}.
func NewEventsI(args ...interface{}) *Events {
    return &Events{js.Global.Get("Phaser").Get("Events").New(args)}
}



// Events Binding conversion method to Events point 
func ToEvents(jsStruct interface{}) *Events {
    if object, ok := jsStruct.(*js.Object); ok {
		return &Events{Object: object}
	}
	return nil
}



// Parent The Sprite that owns these events.
func (self *Events) Parent() *Sprite{
    return &Sprite{self.Object.Get("parent")}
}

// SetParentA The Sprite that owns these events.
func (self *Events) SetParentA(member *Sprite) {
    self.Object.Set("parent", member)
}

// OnAddedToGroup This signal is dispatched when this Game Object is added to a new Group.
// It is sent two arguments:
// {any} The Game Object that was added to the Group.
// {Phaser.Group} The Group it was added to.
func (self *Events) OnAddedToGroup() *Signal{
    return &Signal{self.Object.Get("onAddedToGroup")}
}

// SetOnAddedToGroupA This signal is dispatched when this Game Object is added to a new Group.
// It is sent two arguments:
// {any} The Game Object that was added to the Group.
// {Phaser.Group} The Group it was added to.
func (self *Events) SetOnAddedToGroupA(member *Signal) {
    self.Object.Set("onAddedToGroup", member)
}

// OnRemovedFromGroup This signal is dispatched when the Game Object is removed from a Group.
// It is sent two arguments:
// {any} The Game Object that was removed from the Group.
// {Phaser.Group} The Group it was removed from.
func (self *Events) OnRemovedFromGroup() *Signal{
    return &Signal{self.Object.Get("onRemovedFromGroup")}
}

// SetOnRemovedFromGroupA This signal is dispatched when the Game Object is removed from a Group.
// It is sent two arguments:
// {any} The Game Object that was removed from the Group.
// {Phaser.Group} The Group it was removed from.
func (self *Events) SetOnRemovedFromGroupA(member *Signal) {
    self.Object.Set("onRemovedFromGroup", member)
}

// OnRemovedFromWorld This Signal is never used internally by Phaser and is now deprecated.
func (self *Events) OnRemovedFromWorld() *Signal{
    return &Signal{self.Object.Get("onRemovedFromWorld")}
}

// SetOnRemovedFromWorldA This Signal is never used internally by Phaser and is now deprecated.
func (self *Events) SetOnRemovedFromWorldA(member *Signal) {
    self.Object.Set("onRemovedFromWorld", member)
}

// OnDestroy This signal is dispatched when the Game Object is destroyed.
// This happens when `Sprite.destroy()` is called, or `Group.destroy()` with `destroyChildren` set to true.
// It is sent one argument:
// {any} The Game Object that was destroyed.
func (self *Events) OnDestroy() *Signal{
    return &Signal{self.Object.Get("onDestroy")}
}

// SetOnDestroyA This signal is dispatched when the Game Object is destroyed.
// This happens when `Sprite.destroy()` is called, or `Group.destroy()` with `destroyChildren` set to true.
// It is sent one argument:
// {any} The Game Object that was destroyed.
func (self *Events) SetOnDestroyA(member *Signal) {
    self.Object.Set("onDestroy", member)
}

// OnKilled This signal is dispatched when the Game Object is killed.
// This happens when `Sprite.kill()` is called.
// Please understand the difference between `kill` and `destroy` by looking at their respective methods.
// It is sent one argument:
// {any} The Game Object that was killed.
func (self *Events) OnKilled() *Signal{
    return &Signal{self.Object.Get("onKilled")}
}

// SetOnKilledA This signal is dispatched when the Game Object is killed.
// This happens when `Sprite.kill()` is called.
// Please understand the difference between `kill` and `destroy` by looking at their respective methods.
// It is sent one argument:
// {any} The Game Object that was killed.
func (self *Events) SetOnKilledA(member *Signal) {
    self.Object.Set("onKilled", member)
}

// OnRevived This signal is dispatched when the Game Object is revived from a previously killed state.
// This happens when `Sprite.revive()` is called.
// It is sent one argument:
// {any} The Game Object that was revived.
func (self *Events) OnRevived() *Signal{
    return &Signal{self.Object.Get("onRevived")}
}

// SetOnRevivedA This signal is dispatched when the Game Object is revived from a previously killed state.
// This happens when `Sprite.revive()` is called.
// It is sent one argument:
// {any} The Game Object that was revived.
func (self *Events) SetOnRevivedA(member *Signal) {
    self.Object.Set("onRevived", member)
}

// OnOutOfBounds This signal is dispatched when the Game Object leaves the Phaser.World bounds.
// This signal is only if `Sprite.checkWorldBounds` is set to `true`.
// It is sent one argument:
// {any} The Game Object that left the World bounds.
func (self *Events) OnOutOfBounds() *Signal{
    return &Signal{self.Object.Get("onOutOfBounds")}
}

// SetOnOutOfBoundsA This signal is dispatched when the Game Object leaves the Phaser.World bounds.
// This signal is only if `Sprite.checkWorldBounds` is set to `true`.
// It is sent one argument:
// {any} The Game Object that left the World bounds.
func (self *Events) SetOnOutOfBoundsA(member *Signal) {
    self.Object.Set("onOutOfBounds", member)
}

// OnEnterBounds This signal is dispatched when the Game Object returns within the Phaser.World bounds, having previously been outside of them.
// This signal is only if `Sprite.checkWorldBounds` is set to `true`.
// It is sent one argument:
// {any} The Game Object that entered the World bounds.
func (self *Events) OnEnterBounds() *Signal{
    return &Signal{self.Object.Get("onEnterBounds")}
}

// SetOnEnterBoundsA This signal is dispatched when the Game Object returns within the Phaser.World bounds, having previously been outside of them.
// This signal is only if `Sprite.checkWorldBounds` is set to `true`.
// It is sent one argument:
// {any} The Game Object that entered the World bounds.
func (self *Events) SetOnEnterBoundsA(member *Signal) {
    self.Object.Set("onEnterBounds", member)
}

// OnInputOver This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an over event from a Phaser.Pointer.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) OnInputOver() *Signal{
    return &Signal{self.Object.Get("onInputOver")}
}

// SetOnInputOverA This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an over event from a Phaser.Pointer.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) SetOnInputOverA(member *Signal) {
    self.Object.Set("onInputOver", member)
}

// OnInputOut This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an out event from a Phaser.Pointer, which was previously over it.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) OnInputOut() *Signal{
    return &Signal{self.Object.Get("onInputOut")}
}

// SetOnInputOutA This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an out event from a Phaser.Pointer, which was previously over it.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) SetOnInputOutA(member *Signal) {
    self.Object.Set("onInputOut", member)
}

// OnInputDown This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives a down event from a Phaser.Pointer. This effectively means the Pointer has been
// pressed down (but not yet released) on the Game Object.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) OnInputDown() *Signal{
    return &Signal{self.Object.Get("onInputDown")}
}

// SetOnInputDownA This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives a down event from a Phaser.Pointer. This effectively means the Pointer has been
// pressed down (but not yet released) on the Game Object.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) SetOnInputDownA(member *Signal) {
    self.Object.Set("onInputDown", member)
}

// OnInputUp This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an up event from a Phaser.Pointer. This effectively means the Pointer had been
// pressed down, and was then released on the Game Object.
// It is sent three arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
// {boolean} isOver - Is the Pointer still over the Game Object?
func (self *Events) OnInputUp() *Signal{
    return &Signal{self.Object.Get("onInputUp")}
}

// SetOnInputUpA This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an up event from a Phaser.Pointer. This effectively means the Pointer had been
// pressed down, and was then released on the Game Object.
// It is sent three arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
// {boolean} isOver - Is the Pointer still over the Game Object?
func (self *Events) SetOnInputUpA(member *Signal) {
    self.Object.Set("onInputUp", member)
}

// OnDragStart This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
// It is sent when a Phaser.Pointer starts to drag the Game Object, taking into consideration the various
// drag limitations that may be set.
// It is sent four arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
// {number} The x coordinate that the drag started from.
// {number} The y coordinate that the drag started from.
func (self *Events) OnDragStart() *Signal{
    return &Signal{self.Object.Get("onDragStart")}
}

// SetOnDragStartA This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
// It is sent when a Phaser.Pointer starts to drag the Game Object, taking into consideration the various
// drag limitations that may be set.
// It is sent four arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
// {number} The x coordinate that the drag started from.
// {number} The y coordinate that the drag started from.
func (self *Events) SetOnDragStartA(member *Signal) {
    self.Object.Set("onDragStart", member)
}

// OnDragUpdate This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
// It is sent when a Phaser.Pointer is actively dragging the Game Object.
// Be warned: This is a high volume Signal. Be careful what you bind to it.
// It is sent six arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
// {number} The new x coordinate of the Game Object.
// {number} The new y coordinate of the Game Object.
// {Phaser.Point} A Point object that contains the point the Game Object was snapped to, if `snapOnDrag` has been enabled.
// {boolean} The `fromStart` boolean, indicates if this is the first update immediately after the drag has started.
func (self *Events) OnDragUpdate() *Signal{
    return &Signal{self.Object.Get("onDragUpdate")}
}

// SetOnDragUpdateA This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
// It is sent when a Phaser.Pointer is actively dragging the Game Object.
// Be warned: This is a high volume Signal. Be careful what you bind to it.
// It is sent six arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
// {number} The new x coordinate of the Game Object.
// {number} The new y coordinate of the Game Object.
// {Phaser.Point} A Point object that contains the point the Game Object was snapped to, if `snapOnDrag` has been enabled.
// {boolean} The `fromStart` boolean, indicates if this is the first update immediately after the drag has started.
func (self *Events) SetOnDragUpdateA(member *Signal) {
    self.Object.Set("onDragUpdate", member)
}

// OnDragStop This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
// It is sent when a Phaser.Pointer stops dragging the Game Object.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) OnDragStop() *Signal{
    return &Signal{self.Object.Get("onDragStop")}
}

// SetOnDragStopA This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
// It is sent when a Phaser.Pointer stops dragging the Game Object.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) SetOnDragStopA(member *Signal) {
    self.Object.Set("onDragStop", member)
}

// OnAnimationStart This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has been played.
// You can also listen to `Animation.onStart` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that was started.
func (self *Events) OnAnimationStart() *Signal{
    return &Signal{self.Object.Get("onAnimationStart")}
}

// SetOnAnimationStartA This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has been played.
// You can also listen to `Animation.onStart` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that was started.
func (self *Events) SetOnAnimationStartA(member *Signal) {
    self.Object.Set("onAnimationStart", member)
}

// OnAnimationComplete This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has been stopped (via `animation.stop()` and the `dispatchComplete` argument has been set.
// You can also listen to `Animation.onComplete` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that was stopped.
func (self *Events) OnAnimationComplete() *Signal{
    return &Signal{self.Object.Get("onAnimationComplete")}
}

// SetOnAnimationCompleteA This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has been stopped (via `animation.stop()` and the `dispatchComplete` argument has been set.
// You can also listen to `Animation.onComplete` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that was stopped.
func (self *Events) SetOnAnimationCompleteA(member *Signal) {
    self.Object.Set("onAnimationComplete", member)
}

// OnAnimationLoop This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has looped playback.
// You can also listen to `Animation.onLoop` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that looped.
func (self *Events) OnAnimationLoop() *Signal{
    return &Signal{self.Object.Get("onAnimationLoop")}
}

// SetOnAnimationLoopA This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has looped playback.
// You can also listen to `Animation.onLoop` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that looped.
func (self *Events) SetOnAnimationLoopA(member *Signal) {
    self.Object.Set("onAnimationLoop", member)
}


// Destroy Removes all events.
func (self *Events) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Removes all events.
func (self *Events) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

