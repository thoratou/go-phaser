// Automatic generation for Phaser.Key
// generated file Key.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// If you need more fine-grained control over the handling of specific keys you can create and use Phaser.Key objects.
type Key struct {
    *js.Object
}


// A reference to the currently running game.
func (self *Key) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *Key) SetGame(member *Game) {
    self.Set("game", member)
}

// Stores the most recent DOM event.
func (self *Key) GetEvent() interface{}{
    return self.Get("event")
}

// Stores the most recent DOM event.
func (self *Key) SetEvent(member interface{}) {
    self.Set("event", member)
}

// The "down" state of the key. This will remain `true` for as long as the keyboard thinks this key is held down.
func (self *Key) GetIsDown() bool{
    return self.Get("isDown").Bool()
}

// The "down" state of the key. This will remain `true` for as long as the keyboard thinks this key is held down.
func (self *Key) SetIsDown(member bool) {
    self.Set("isDown", member)
}

// The "up" state of the key. This will remain `true` for as long as the keyboard thinks this key is up.
func (self *Key) GetIsUp() bool{
    return self.Get("isUp").Bool()
}

// The "up" state of the key. This will remain `true` for as long as the keyboard thinks this key is up.
func (self *Key) SetIsUp(member bool) {
    self.Set("isUp", member)
}

// The down state of the ALT key, if pressed at the same time as this key.
func (self *Key) GetAltKey() bool{
    return self.Get("altKey").Bool()
}

// The down state of the ALT key, if pressed at the same time as this key.
func (self *Key) SetAltKey(member bool) {
    self.Set("altKey", member)
}

// The down state of the CTRL key, if pressed at the same time as this key.
func (self *Key) GetCtrlKey() bool{
    return self.Get("ctrlKey").Bool()
}

// The down state of the CTRL key, if pressed at the same time as this key.
func (self *Key) SetCtrlKey(member bool) {
    self.Set("ctrlKey", member)
}

// The down state of the SHIFT key, if pressed at the same time as this key.
func (self *Key) GetShiftKey() bool{
    return self.Get("shiftKey").Bool()
}

// The down state of the SHIFT key, if pressed at the same time as this key.
func (self *Key) SetShiftKey(member bool) {
    self.Set("shiftKey", member)
}

// The timestamp when the key was last pressed down. This is based on Game.time.now.
func (self *Key) GetTimeDown() int{
    return self.Get("timeDown").Int()
}

// The timestamp when the key was last pressed down. This is based on Game.time.now.
func (self *Key) SetTimeDown(member int) {
    self.Set("timeDown", member)
}

// If the key is down this value holds the duration of that key press and is constantly updated.
// If the key is up it holds the duration of the previous down session. The number of milliseconds this key has been held down for.
func (self *Key) GetDuration() int{
    return self.Get("duration").Int()
}

// If the key is down this value holds the duration of that key press and is constantly updated.
// If the key is up it holds the duration of the previous down session. The number of milliseconds this key has been held down for.
func (self *Key) SetDuration(member int) {
    self.Set("duration", member)
}

// The timestamp when the key was last released. This is based on Game.time.now.
func (self *Key) GetTimeUp() int{
    return self.Get("timeUp").Int()
}

// The timestamp when the key was last released. This is based on Game.time.now.
func (self *Key) SetTimeUp(member int) {
    self.Set("timeUp", member)
}

// If a key is held down this holds down the number of times the key has 'repeated'.
func (self *Key) GetRepeats() int{
    return self.Get("repeats").Int()
}

// If a key is held down this holds down the number of times the key has 'repeated'.
func (self *Key) SetRepeats(member int) {
    self.Set("repeats", member)
}

// The keycode of this key.
func (self *Key) GetKeyCode() int{
    return self.Get("keyCode").Int()
}

// The keycode of this key.
func (self *Key) SetKeyCode(member int) {
    self.Set("keyCode", member)
}

// This Signal is dispatched every time this Key is pressed down. It is only dispatched once (until the key is released again).
func (self *Key) GetOnDown() *Signal{
    return &Signal{self.Get("onDown")}
}

// This Signal is dispatched every time this Key is pressed down. It is only dispatched once (until the key is released again).
func (self *Key) SetOnDown(member *Signal) {
    self.Set("onDown", member)
}

// A callback that is called while this Key is held down. Warning: Depending on refresh rate that could be 60+ times per second.
func (self *Key) SetOnHoldCallback(member func(...interface{})) {
    self.Set("onHoldCallback", member)
}

// The context under which the onHoldCallback will be called.
func (self *Key) GetOnHoldContext() interface{}{
    return self.Get("onHoldContext")
}

// The context under which the onHoldCallback will be called.
func (self *Key) SetOnHoldContext(member interface{}) {
    self.Set("onHoldContext", member)
}

// This Signal is dispatched every time this Key is released. It is only dispatched once (until the key is pressed and released again).
func (self *Key) GetOnUp() *Signal{
    return &Signal{self.Get("onUp")}
}

// This Signal is dispatched every time this Key is released. It is only dispatched once (until the key is pressed and released again).
func (self *Key) SetOnUp(member *Signal) {
    self.Set("onUp", member)
}



// Called automatically by Phaser.Keyboard.
func (self *Key) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Called automatically by Phaser.Keyboard.
func (self *Key) ProcessKeyDownI(args ...interface{}) {
    self.Call("processKeyDown", args)
}

// Called automatically by Phaser.Keyboard.
func (self *Key) ProcessKeyUpI(args ...interface{}) {
    self.Call("processKeyUp", args)
}

// Resets the state of this Key.
// 
// This sets isDown to false, isUp to true, resets the time to be the current time, and _enables_ the key.
// In addition, if it is a "hard reset", it clears clears any callbacks associated with the onDown and onUp events and removes the onHoldCallback.
func (self *Key) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) DownDurationI(args ...interface{}) bool{
    return self.Call("downDuration", args).Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) UpDurationI(args ...interface{}) bool{
    return self.Call("upDuration", args).Bool()
}
