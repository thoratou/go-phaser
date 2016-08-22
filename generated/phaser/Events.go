// Automatic generation for Phaser.Events
// generated file Events.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Events component is a collection of events fired by the parent Game Object.
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


// The Sprite that owns these events.
func (self *Events) GetParentA() *Sprite{
    return &Sprite{self.Object.Get("parent")}
}

// The Sprite that owns these events.
func (self *Events) SetParentA(member *Sprite) {
    self.Object.Set("parent", member)
}

// This signal is dispatched when this Game Object is added to a new Group.
// It is sent two arguments:
// {any} The Game Object that was added to the Group.
// {Phaser.Group} The Group it was added to.
func (self *Events) GetOnAddedToGroupA() *Signal{
    return &Signal{self.Object.Get("onAddedToGroup")}
}

// This signal is dispatched when this Game Object is added to a new Group.
// It is sent two arguments:
// {any} The Game Object that was added to the Group.
// {Phaser.Group} The Group it was added to.
func (self *Events) SetOnAddedToGroupA(member *Signal) {
    self.Object.Set("onAddedToGroup", member)
}

// This signal is dispatched when the Game Object is removed from a Group.
// It is sent two arguments:
// {any} The Game Object that was removed from the Group.
// {Phaser.Group} The Group it was removed from.
func (self *Events) GetOnRemovedFromGroupA() *Signal{
    return &Signal{self.Object.Get("onRemovedFromGroup")}
}

// This signal is dispatched when the Game Object is removed from a Group.
// It is sent two arguments:
// {any} The Game Object that was removed from the Group.
// {Phaser.Group} The Group it was removed from.
func (self *Events) SetOnRemovedFromGroupA(member *Signal) {
    self.Object.Set("onRemovedFromGroup", member)
}

// This Signal is never used internally by Phaser and is now deprecated.
func (self *Events) GetOnRemovedFromWorldA() *Signal{
    return &Signal{self.Object.Get("onRemovedFromWorld")}
}

// This Signal is never used internally by Phaser and is now deprecated.
func (self *Events) SetOnRemovedFromWorldA(member *Signal) {
    self.Object.Set("onRemovedFromWorld", member)
}

// This signal is dispatched when the Game Object is destroyed.
// This happens when `Sprite.destroy()` is called, or `Group.destroy()` with `destroyChildren` set to true.
// It is sent one argument:
// {any} The Game Object that was destroyed.
func (self *Events) GetOnDestroyA() *Signal{
    return &Signal{self.Object.Get("onDestroy")}
}

// This signal is dispatched when the Game Object is destroyed.
// This happens when `Sprite.destroy()` is called, or `Group.destroy()` with `destroyChildren` set to true.
// It is sent one argument:
// {any} The Game Object that was destroyed.
func (self *Events) SetOnDestroyA(member *Signal) {
    self.Object.Set("onDestroy", member)
}

// This signal is dispatched when the Game Object is killed.
// This happens when `Sprite.kill()` is called.
// Please understand the difference between `kill` and `destroy` by looking at their respective methods.
// It is sent one argument:
// {any} The Game Object that was killed.
func (self *Events) GetOnKilledA() *Signal{
    return &Signal{self.Object.Get("onKilled")}
}

// This signal is dispatched when the Game Object is killed.
// This happens when `Sprite.kill()` is called.
// Please understand the difference between `kill` and `destroy` by looking at their respective methods.
// It is sent one argument:
// {any} The Game Object that was killed.
func (self *Events) SetOnKilledA(member *Signal) {
    self.Object.Set("onKilled", member)
}

// This signal is dispatched when the Game Object is revived from a previously killed state.
// This happens when `Sprite.revive()` is called.
// It is sent one argument:
// {any} The Game Object that was revived.
func (self *Events) GetOnRevivedA() *Signal{
    return &Signal{self.Object.Get("onRevived")}
}

// This signal is dispatched when the Game Object is revived from a previously killed state.
// This happens when `Sprite.revive()` is called.
// It is sent one argument:
// {any} The Game Object that was revived.
func (self *Events) SetOnRevivedA(member *Signal) {
    self.Object.Set("onRevived", member)
}

// This signal is dispatched when the Game Object leaves the Phaser.World bounds.
// This signal is only if `Sprite.checkWorldBounds` is set to `true`.
// It is sent one argument:
// {any} The Game Object that left the World bounds.
func (self *Events) GetOnOutOfBoundsA() *Signal{
    return &Signal{self.Object.Get("onOutOfBounds")}
}

// This signal is dispatched when the Game Object leaves the Phaser.World bounds.
// This signal is only if `Sprite.checkWorldBounds` is set to `true`.
// It is sent one argument:
// {any} The Game Object that left the World bounds.
func (self *Events) SetOnOutOfBoundsA(member *Signal) {
    self.Object.Set("onOutOfBounds", member)
}

// This signal is dispatched when the Game Object returns within the Phaser.World bounds, having previously been outside of them.
// This signal is only if `Sprite.checkWorldBounds` is set to `true`.
// It is sent one argument:
// {any} The Game Object that entered the World bounds.
func (self *Events) GetOnEnterBoundsA() *Signal{
    return &Signal{self.Object.Get("onEnterBounds")}
}

// This signal is dispatched when the Game Object returns within the Phaser.World bounds, having previously been outside of them.
// This signal is only if `Sprite.checkWorldBounds` is set to `true`.
// It is sent one argument:
// {any} The Game Object that entered the World bounds.
func (self *Events) SetOnEnterBoundsA(member *Signal) {
    self.Object.Set("onEnterBounds", member)
}

// This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an over event from a Phaser.Pointer.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) GetOnInputOverA() *Signal{
    return &Signal{self.Object.Get("onInputOver")}
}

// This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an over event from a Phaser.Pointer.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) SetOnInputOverA(member *Signal) {
    self.Object.Set("onInputOver", member)
}

// This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an out event from a Phaser.Pointer, which was previously over it.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) GetOnInputOutA() *Signal{
    return &Signal{self.Object.Get("onInputOut")}
}

// This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an out event from a Phaser.Pointer, which was previously over it.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) SetOnInputOutA(member *Signal) {
    self.Object.Set("onInputOut", member)
}

// This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives a down event from a Phaser.Pointer. This effectively means the Pointer has been
// pressed down (but not yet released) on the Game Object.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) GetOnInputDownA() *Signal{
    return &Signal{self.Object.Get("onInputDown")}
}

// This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives a down event from a Phaser.Pointer. This effectively means the Pointer has been
// pressed down (but not yet released) on the Game Object.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) SetOnInputDownA(member *Signal) {
    self.Object.Set("onInputDown", member)
}

// This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an up event from a Phaser.Pointer. This effectively means the Pointer had been
// pressed down, and was then released on the Game Object.
// It is sent three arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
// {boolean} isOver - Is the Pointer still over the Game Object?
func (self *Events) GetOnInputUpA() *Signal{
    return &Signal{self.Object.Get("onInputUp")}
}

// This signal is dispatched if the Game Object has `inputEnabled` set to `true`, 
// and receives an up event from a Phaser.Pointer. This effectively means the Pointer had been
// pressed down, and was then released on the Game Object.
// It is sent three arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
// {boolean} isOver - Is the Pointer still over the Game Object?
func (self *Events) SetOnInputUpA(member *Signal) {
    self.Object.Set("onInputUp", member)
}

// This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
// It is sent when a Phaser.Pointer starts to drag the Game Object, taking into consideration the various
// drag limitations that may be set.
// It is sent four arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
// {number} The x coordinate that the drag started from.
// {number} The y coordinate that the drag started from.
func (self *Events) GetOnDragStartA() *Signal{
    return &Signal{self.Object.Get("onDragStart")}
}

// This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
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

// This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
// It is sent when a Phaser.Pointer is actively dragging the Game Object.
// Be warned: This is a high volume Signal. Be careful what you bind to it.
// It is sent six arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
// {number} The new x coordinate of the Game Object.
// {number} The new y coordinate of the Game Object.
// {Phaser.Point} A Point object that contains the point the Game Object was snapped to, if `snapOnDrag` has been enabled.
// {boolean} The `fromStart` boolean, indicates if this is the first update immediately after the drag has started.
func (self *Events) GetOnDragUpdateA() *Signal{
    return &Signal{self.Object.Get("onDragUpdate")}
}

// This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
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

// This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
// It is sent when a Phaser.Pointer stops dragging the Game Object.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) GetOnDragStopA() *Signal{
    return &Signal{self.Object.Get("onDragStop")}
}

// This signal is dispatched if the Game Object has been `inputEnabled` and `enableDrag` has been set.
// It is sent when a Phaser.Pointer stops dragging the Game Object.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Pointer} The Phaser.Pointer object that caused the event.
func (self *Events) SetOnDragStopA(member *Signal) {
    self.Object.Set("onDragStop", member)
}

// This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has been played.
// You can also listen to `Animation.onStart` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that was started.
func (self *Events) GetOnAnimationStartA() *Signal{
    return &Signal{self.Object.Get("onAnimationStart")}
}

// This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has been played.
// You can also listen to `Animation.onStart` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that was started.
func (self *Events) SetOnAnimationStartA(member *Signal) {
    self.Object.Set("onAnimationStart", member)
}

// This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has been stopped (via `animation.stop()` and the `dispatchComplete` argument has been set.
// You can also listen to `Animation.onComplete` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that was stopped.
func (self *Events) GetOnAnimationCompleteA() *Signal{
    return &Signal{self.Object.Get("onAnimationComplete")}
}

// This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has been stopped (via `animation.stop()` and the `dispatchComplete` argument has been set.
// You can also listen to `Animation.onComplete` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that was stopped.
func (self *Events) SetOnAnimationCompleteA(member *Signal) {
    self.Object.Set("onAnimationComplete", member)
}

// This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has looped playback.
// You can also listen to `Animation.onLoop` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that looped.
func (self *Events) GetOnAnimationLoopA() *Signal{
    return &Signal{self.Object.Get("onAnimationLoop")}
}

// This signal is dispatched if the Game Object has the AnimationManager component, 
// and an Animation has looped playback.
// You can also listen to `Animation.onLoop` rather than via the Game Objects events.
// It is sent two arguments:
// {any} The Game Object that received the event.
// {Phaser.Animation} The Phaser.Animation that looped.
func (self *Events) SetOnAnimationLoopA(member *Signal) {
    self.Object.Set("onAnimationLoop", member)
}



// Removes all events.
func (self *Events) Destroy() {
    self.Object.Call("destroy")
}

// Removes all events.
func (self *Events) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
