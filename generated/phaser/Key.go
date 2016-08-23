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


// If you need more fine-grained control over the handling of specific keys you can create and use Phaser.Key objects.
func NewKey(game *Game, keycode int) *Key {
    return &Key{js.Global.Get("Phaser").Get("Key").New(game, keycode)}
}

// If you need more fine-grained control over the handling of specific keys you can create and use Phaser.Key objects.
func NewKeyI(args ...interface{}) *Key {
    return &Key{js.Global.Get("Phaser").Get("Key").New(args)}
}



// A reference to the currently running game.
func (self *Key) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *Key) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Stores the most recent DOM event.
func (self *Key) Event() interface{}{
    return self.Object.Get("event")
}

// Stores the most recent DOM event.
func (self *Key) SetEventA(member interface{}) {
    self.Object.Set("event", member)
}

// The "down" state of the key. This will remain `true` for as long as the keyboard thinks this key is held down.
func (self *Key) IsDown() bool{
    return self.Object.Get("isDown").Bool()
}

// The "down" state of the key. This will remain `true` for as long as the keyboard thinks this key is held down.
func (self *Key) SetIsDownA(member bool) {
    self.Object.Set("isDown", member)
}

// The "up" state of the key. This will remain `true` for as long as the keyboard thinks this key is up.
func (self *Key) IsUp() bool{
    return self.Object.Get("isUp").Bool()
}

// The "up" state of the key. This will remain `true` for as long as the keyboard thinks this key is up.
func (self *Key) SetIsUpA(member bool) {
    self.Object.Set("isUp", member)
}

// The down state of the ALT key, if pressed at the same time as this key.
func (self *Key) AltKey() bool{
    return self.Object.Get("altKey").Bool()
}

// The down state of the ALT key, if pressed at the same time as this key.
func (self *Key) SetAltKeyA(member bool) {
    self.Object.Set("altKey", member)
}

// The down state of the CTRL key, if pressed at the same time as this key.
func (self *Key) CtrlKey() bool{
    return self.Object.Get("ctrlKey").Bool()
}

// The down state of the CTRL key, if pressed at the same time as this key.
func (self *Key) SetCtrlKeyA(member bool) {
    self.Object.Set("ctrlKey", member)
}

// The down state of the SHIFT key, if pressed at the same time as this key.
func (self *Key) ShiftKey() bool{
    return self.Object.Get("shiftKey").Bool()
}

// The down state of the SHIFT key, if pressed at the same time as this key.
func (self *Key) SetShiftKeyA(member bool) {
    self.Object.Set("shiftKey", member)
}

// The timestamp when the key was last pressed down. This is based on Game.time.now.
func (self *Key) TimeDown() int{
    return self.Object.Get("timeDown").Int()
}

// The timestamp when the key was last pressed down. This is based on Game.time.now.
func (self *Key) SetTimeDownA(member int) {
    self.Object.Set("timeDown", member)
}

// If the key is down this value holds the duration of that key press and is constantly updated.
// If the key is up it holds the duration of the previous down session. The number of milliseconds this key has been held down for.
func (self *Key) Duration() int{
    return self.Object.Get("duration").Int()
}

// If the key is down this value holds the duration of that key press and is constantly updated.
// If the key is up it holds the duration of the previous down session. The number of milliseconds this key has been held down for.
func (self *Key) SetDurationA(member int) {
    self.Object.Set("duration", member)
}

// The timestamp when the key was last released. This is based on Game.time.now.
func (self *Key) TimeUp() int{
    return self.Object.Get("timeUp").Int()
}

// The timestamp when the key was last released. This is based on Game.time.now.
func (self *Key) SetTimeUpA(member int) {
    self.Object.Set("timeUp", member)
}

// If a key is held down this holds down the number of times the key has 'repeated'.
func (self *Key) Repeats() int{
    return self.Object.Get("repeats").Int()
}

// If a key is held down this holds down the number of times the key has 'repeated'.
func (self *Key) SetRepeatsA(member int) {
    self.Object.Set("repeats", member)
}

// The keycode of this key.
func (self *Key) KeyCode() int{
    return self.Object.Get("keyCode").Int()
}

// The keycode of this key.
func (self *Key) SetKeyCodeA(member int) {
    self.Object.Set("keyCode", member)
}

// This Signal is dispatched every time this Key is pressed down. It is only dispatched once (until the key is released again).
func (self *Key) OnDown() *Signal{
    return &Signal{self.Object.Get("onDown")}
}

// This Signal is dispatched every time this Key is pressed down. It is only dispatched once (until the key is released again).
func (self *Key) SetOnDownA(member *Signal) {
    self.Object.Set("onDown", member)
}

// A callback that is called while this Key is held down. Warning: Depending on refresh rate that could be 60+ times per second.
func (self *Key) OnHoldCallback() interface{}{
    return self.Object.Get("onHoldCallback")
}

// A callback that is called while this Key is held down. Warning: Depending on refresh rate that could be 60+ times per second.
func (self *Key) SetOnHoldCallbackA(member interface{}) {
    self.Object.Set("onHoldCallback", member)
}

// The context under which the onHoldCallback will be called.
func (self *Key) OnHoldContext() interface{}{
    return self.Object.Get("onHoldContext")
}

// The context under which the onHoldCallback will be called.
func (self *Key) SetOnHoldContextA(member interface{}) {
    self.Object.Set("onHoldContext", member)
}

// This Signal is dispatched every time this Key is released. It is only dispatched once (until the key is pressed and released again).
func (self *Key) OnUp() *Signal{
    return &Signal{self.Object.Get("onUp")}
}

// This Signal is dispatched every time this Key is released. It is only dispatched once (until the key is pressed and released again).
func (self *Key) SetOnUpA(member *Signal) {
    self.Object.Set("onUp", member)
}



// Called automatically by Phaser.Keyboard.
func (self *Key) Update() {
    self.Object.Call("update")
}

// Called automatically by Phaser.Keyboard.
func (self *Key) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Called automatically by Phaser.Keyboard.
func (self *Key) ProcessKeyDown(event *KeyboardEvent) {
    self.Object.Call("processKeyDown", event)
}

// Called automatically by Phaser.Keyboard.
func (self *Key) ProcessKeyDownI(args ...interface{}) {
    self.Object.Call("processKeyDown", args)
}

// Called automatically by Phaser.Keyboard.
func (self *Key) ProcessKeyUp(event *KeyboardEvent) {
    self.Object.Call("processKeyUp", event)
}

// Called automatically by Phaser.Keyboard.
func (self *Key) ProcessKeyUpI(args ...interface{}) {
    self.Object.Call("processKeyUp", args)
}

// Resets the state of this Key.
// 
// This sets isDown to false, isUp to true, resets the time to be the current time, and _enables_ the key.
// In addition, if it is a "hard reset", it clears clears any callbacks associated with the onDown and onUp events and removes the onHoldCallback.
func (self *Key) Reset() {
    self.Object.Call("reset")
}

// Resets the state of this Key.
// 
// This sets isDown to false, isUp to true, resets the time to be the current time, and _enables_ the key.
// In addition, if it is a "hard reset", it clears clears any callbacks associated with the onDown and onUp events and removes the onHoldCallback.
func (self *Key) Reset1O(hard bool) {
    self.Object.Call("reset", hard)
}

// Resets the state of this Key.
// 
// This sets isDown to false, isUp to true, resets the time to be the current time, and _enables_ the key.
// In addition, if it is a "hard reset", it clears clears any callbacks associated with the onDown and onUp events and removes the onHoldCallback.
func (self *Key) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) DownDuration() bool{
    return self.Object.Call("downDuration").Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) DownDuration1O(duration int) bool{
    return self.Object.Call("downDuration", duration).Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) DownDurationI(args ...interface{}) bool{
    return self.Object.Call("downDuration", args).Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) UpDuration() bool{
    return self.Object.Call("upDuration").Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) UpDuration1O(duration int) bool{
    return self.Object.Call("upDuration", duration).Bool()
}

// Returns `true` if the Key was pressed down within the `duration` value given, or `false` if it either isn't down,
// or was pressed down longer ago than then given duration.
func (self *Key) UpDurationI(args ...interface{}) bool{
    return self.Object.Call("upDuration", args).Bool()
}
