// Package phaser Automatic generation for Phaser.Key
// generated file Key.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Key If you need more fine-grained control over the handling of specific keys you can create and use Phaser.Key objects.
type Key struct {
    *js.Object
}

// NewKey If you need more fine-grained control over the handling of specific keys you can create and use Phaser.Key objects.
func NewKey(game *Game, keycode int) *Key {
    return &Key{js.Global.Get("Phaser").Get("Key").New(game, keycode)}
}
// NewKeyI If you need more fine-grained control over the handling of specific keys you can create and use Phaser.Key objects.
func NewKeyI(args ...interface{}) *Key {
    return &Key{js.Global.Get("Phaser").Get("Key").New(args)}
}



// Game A reference to the currently running game.
func (self *Key) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *Key) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Event Stores the most recent DOM event.
func (self *Key) Event() interface{}{
    return self.Object.Get("event")
}

// SetEventA Stores the most recent DOM event.
func (self *Key) SetEventA(member interface{}) {
    self.Object.Set("event", member)
}

// IsDown The "down" state of the key. This will remain `true` for as long as the keyboard thinks this key is held down.
func (self *Key) IsDown() bool{
    return self.Object.Get("isDown").Bool()
}

// SetIsDownA The "down" state of the key. This will remain `true` for as long as the keyboard thinks this key is held down.
func (self *Key) SetIsDownA(member bool) {
    self.Object.Set("isDown", member)
}

// IsUp The "up" state of the key. This will remain `true` for as long as the keyboard thinks this key is up.
func (self *Key) IsUp() bool{
    return self.Object.Get("isUp").Bool()
}

// SetIsUpA The "up" state of the key. This will remain `true` for as long as the keyboard thinks this key is up.
func (self *Key) SetIsUpA(member bool) {
    self.Object.Set("isUp", member)
}

// AltKey The down state of the ALT key, if pressed at the same time as this key.
func (self *Key) AltKey() bool{
    return self.Object.Get("altKey").Bool()
}

// SetAltKeyA The down state of the ALT key, if pressed at the same time as this key.
func (self *Key) SetAltKeyA(member bool) {
    self.Object.Set("altKey", member)
}

// CtrlKey The down state of the CTRL key, if pressed at the same time as this key.
func (self *Key) CtrlKey() bool{
    return self.Object.Get("ctrlKey").Bool()
}

// SetCtrlKeyA The down state of the CTRL key, if pressed at the same time as this key.
func (self *Key) SetCtrlKeyA(member bool) {
    self.Object.Set("ctrlKey", member)
}

// ShiftKey The down state of the SHIFT key, if pressed at the same time as this key.
func (self *Key) ShiftKey() bool{
    return self.Object.Get("shiftKey").Bool()
}

// SetShiftKeyA The down state of the SHIFT key, if pressed at the same time as this key.
func (self *Key) SetShiftKeyA(member bool) {
    self.Object.Set("shiftKey", member)
}

// TimeDown The timestamp when the key was last pressed down. This is based on Game.time.now.
func (self *Key) TimeDown() int{
    return self.Object.Get("timeDown").Int()
}

// SetTimeDownA The timestamp when the key was last pressed down. This is based on Game.time.now.
func (self *Key) SetTimeDownA(member int) {
    self.Object.Set("timeDown", member)
}

// Duration If the key is down this value holds the duration of that key press and is constantly updated.
// If the key is up it holds the duration of the previous down session. The number of milliseconds this key has been held down for.
func (self *Key) Duration() int{
    return self.Object.Get("duration").Int()
}

// SetDurationA If the key is down this value holds the duration of that key press and is constantly updated.
// If the key is up it holds the duration of the previous down session. The number of milliseconds this key has been held down for.
func (self *Key) SetDurationA(member int) {
    self.Object.Set("duration", member)
}

// TimeUp The timestamp when the key was last released. This is based on Game.time.now.
func (self *Key) TimeUp() int{
    return self.Object.Get("timeUp").Int()
}

// SetTimeUpA The timestamp when the key was last released. This is based on Game.time.now.
func (self *Key) SetTimeUpA(member int) {
    self.Object.Set("timeUp", member)
}

// Repeats If a key is held down this holds down the number of times the key has 'repeated'.
func (self *Key) Repeats() int{
    return self.Object.Get("repeats").Int()
}

// SetRepeatsA If a key is held down this holds down the number of times the key has 'repeated'.
func (self *Key) SetRepeatsA(member int) {
    self.Object.Set("repeats", member)
}

// KeyCode The keycode of this key.
func (self *Key) KeyCode() int{
    return self.Object.Get("keyCode").Int()
}

// SetKeyCodeA The keycode of this key.
func (self *Key) SetKeyCodeA(member int) {
    self.Object.Set("keyCode", member)
}

// OnDown This Signal is dispatched every time this Key is pressed down. It is only dispatched once (until the key is released again).
func (self *Key) OnDown() *Signal{
    return &Signal{self.Object.Get("onDown")}
}

// SetOnDownA This Signal is dispatched every time this Key is pressed down. It is only dispatched once (until the key is released again).
func (self *Key) SetOnDownA(member *Signal) {
    self.Object.Set("onDown", member)
}

// OnHoldCallback A callback that is called while this Key is held down. Warning: Depending on refresh rate that could be 60+ times per second.
func (self *Key) OnHoldCallback() interface{}{
    return self.Object.Get("onHoldCallback")
}

// SetOnHoldCallbackA A callback that is called while this Key is held down. Warning: Depending on refresh rate that could be 60+ times per second.
func (self *Key) SetOnHoldCallbackA(member interface{}) {
    self.Object.Set("onHoldCallback", member)
}

// OnHoldContext The context under which the onHoldCallback will be called.
func (self *Key) OnHoldContext() interface{}{
    return self.Object.Get("onHoldContext")
}

// SetOnHoldContextA The context under which the onHoldCallback will be called.
func (self *Key) SetOnHoldContextA(member interface{}) {
    self.Object.Set("onHoldContext", member)
}

// OnUp This Signal is dispatched every time this Key is released. It is only dispatched once (until the key is pressed and released again).
func (self *Key) OnUp() *Signal{
    return &Signal{self.Object.Get("onUp")}
}

// SetOnUpA This Signal is dispatched every time this Key is released. It is only dispatched once (until the key is pressed and released again).
func (self *Key) SetOnUpA(member *Signal) {
    self.Object.Set("onUp", member)
}


// Update Called automatically by Phaser.Keyboard.
func (self *Key) Update() {
    self.Object.Call("update")
}

// UpdateI Called automatically by Phaser.Keyboard.
func (self *Key) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// ProcessKeyDown Called automatically by Phaser.Keyboard.
func (self *Key) ProcessKeyDown(event *KeyboardEvent) {
    self.Object.Call("processKeyDown", event)
}

// ProcessKeyDownI Called automatically by Phaser.Keyboard.
func (self *Key) ProcessKeyDownI(args ...interface{}) {
    self.Object.Call("processKeyDown", args)
}

// ProcessKeyUp Called automatically by Phaser.Keyboard.
func (self *Key) ProcessKeyUp(event *KeyboardEvent) {
    self.Object.Call("processKeyUp", event)
}

// ProcessKeyUpI Called automatically by Phaser.Keyboard.
func (self *Key) ProcessKeyUpI(args ...interface{}) {
    self.Object.Call("processKeyUp", args)
}

// Reset Resets the state of this Key.
// 
// This sets isDown to false, isUp to true, resets the time to be the current time, and _enables_ the key.
// In addition, if it is a "hard reset", it clears clears any callbacks associated with the onDown and onUp events and removes the onHoldCallback.
func (self *Key) Reset() {
    self.Object.Call("reset")
}

// Reset1O Resets the state of this Key.
// 
// This sets isDown to false, isUp to true, resets the time to be the current time, and _enables_ the key.
// In addition, if it is a "hard reset", it clears clears any callbacks associated with the onDown and onUp events and removes the onHoldCallback.
func (self *Key) Reset1O(hard bool) {
    self.Object.Call("reset", hard)
}

// ResetI Resets the state of this Key.
// 
// This sets isDown to false, isUp to true, resets the time to be the current time, and _enables_ the key.
// In addition, if it is a "hard reset", it clears clears any callbacks associated with the onDown and onUp events and removes the onHoldCallback.
func (self *Key) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// DownDuration Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) DownDuration() bool{
    return self.Object.Call("downDuration").Bool()
}

// DownDuration1O Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) DownDuration1O(duration int) bool{
    return self.Object.Call("downDuration", duration).Bool()
}

// DownDurationI Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) DownDurationI(args ...interface{}) bool{
    return self.Object.Call("downDuration", args).Bool()
}

// UpDuration Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) UpDuration() bool{
    return self.Object.Call("upDuration").Bool()
}

// UpDuration1O Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) UpDuration1O(duration int) bool{
    return self.Object.Call("upDuration", duration).Bool()
}

// UpDurationI Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) UpDurationI(args ...interface{}) bool{
    return self.Object.Call("upDuration", args).Bool()
}

