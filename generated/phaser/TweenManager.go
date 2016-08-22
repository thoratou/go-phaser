// Automatic generation for Phaser.TweenManager
// generated file TweenManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Phaser.Game has a single instance of the TweenManager through which all Tween objects are created and updated.
// Tweens are hooked into the game clock and pause system, adjusting based on the game state.
// 
// TweenManager is based heavily on tween.js by http://soledadpenades.com.
// The difference being that tweens belong to a games instance of TweenManager, rather than to a global TWEEN object.
// It also has callbacks swapped for Signals and a few issues patched with regard to properties and completion errors.
// Please see https://github.com/sole/tween.js for a full list of contributors.
type TweenManager struct {
    *js.Object
}


// Local reference to game.
func (self *TweenManager) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *TweenManager) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Are all newly created Tweens frame or time based? A frame based tween will use the physics elapsed timer when updating. This means
// it will retain the same consistent frame rate, regardless of the speed of the device. The duration value given should
// be given in frames.
// 
// If the Tween uses a time based update (which is the default) then the duration is given in milliseconds.
// In this situation a 2000ms tween will last exactly 2 seconds, regardless of the device and how many visual updates the tween
// has actually been through. For very short tweens you may wish to experiment with a frame based update instead.
func (self *TweenManager) GetFrameBasedA() bool{
    return self.Object.Get("frameBased").Bool()
}

// Are all newly created Tweens frame or time based? A frame based tween will use the physics elapsed timer when updating. This means
// it will retain the same consistent frame rate, regardless of the speed of the device. The duration value given should
// be given in frames.
// 
// If the Tween uses a time based update (which is the default) then the duration is given in milliseconds.
// In this situation a 2000ms tween will last exactly 2 seconds, regardless of the device and how many visual updates the tween
// has actually been through. For very short tweens you may wish to experiment with a frame based update instead.
func (self *TweenManager) SetFrameBasedA(member bool) {
    self.Object.Set("frameBased", member)
}



// Get all the tween objects in an array.
func (self *TweenManager) GetAll() []Tween{
	array00 := self.Object.Call("getAll")
	length00 := array00.Length()
	out00 := make([]Tween, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tween{array00.Index(i00)}
	}
	return out00
}

// Get all the tween objects in an array.
func (self *TweenManager) GetAllI(args ...interface{}) []Tween{
	array00 := self.Object.Call("getAll", args)
	length00 := array00.Length()
	out00 := make([]Tween, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = Tween{array00.Index(i00)}
	}
	return out00
}

// Remove all tweens running and in the queue. Doesn't call any of the tween onComplete events.
func (self *TweenManager) RemoveAll() {
    self.Object.Call("removeAll")
}

// Remove all tweens running and in the queue. Doesn't call any of the tween onComplete events.
func (self *TweenManager) RemoveAllI(args ...interface{}) {
    self.Object.Call("removeAll", args)
}

// Remove all tweens from a specific object, array of objects or Group.
func (self *TweenManager) RemoveFrom(obj interface{}) {
    self.Object.Call("removeFrom", obj)
}

// Remove all tweens from a specific object, array of objects or Group.
func (self *TweenManager) RemoveFrom1O(obj interface{}, children bool) {
    self.Object.Call("removeFrom", obj, children)
}

// Remove all tweens from a specific object, array of objects or Group.
func (self *TweenManager) RemoveFromI(args ...interface{}) {
    self.Object.Call("removeFrom", args)
}

// Add a new tween into the TweenManager.
func (self *TweenManager) Add(tween *Tween) *Tween{
    return &Tween{self.Object.Call("add", tween)}
}

// Add a new tween into the TweenManager.
func (self *TweenManager) AddI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("add", args)}
}

// Create a tween object for a specific object. The object can be any JavaScript object or Phaser object such as Sprite.
func (self *TweenManager) Create(object interface{}) *Tween{
    return &Tween{self.Object.Call("create", object)}
}

// Create a tween object for a specific object. The object can be any JavaScript object or Phaser object such as Sprite.
func (self *TweenManager) CreateI(args ...interface{}) *Tween{
    return &Tween{self.Object.Call("create", args)}
}

// Remove a tween from this manager.
func (self *TweenManager) Remove(tween *Tween) {
    self.Object.Call("remove", tween)
}

// Remove a tween from this manager.
func (self *TweenManager) RemoveI(args ...interface{}) {
    self.Object.Call("remove", args)
}

// Update all the tween objects you added to this manager.
func (self *TweenManager) Update() bool{
    return self.Object.Call("update").Bool()
}

// Update all the tween objects you added to this manager.
func (self *TweenManager) UpdateI(args ...interface{}) bool{
    return self.Object.Call("update", args).Bool()
}

// Checks to see if a particular Sprite is currently being tweened.
func (self *TweenManager) IsTweening(object interface{}) bool{
    return self.Object.Call("isTweening", object).Bool()
}

// Checks to see if a particular Sprite is currently being tweened.
func (self *TweenManager) IsTweeningI(args ...interface{}) bool{
    return self.Object.Call("isTweening", args).Bool()
}

// Private. Called by game focus loss. Pauses all currently running tweens.
func (self *TweenManager) _pauseAll() {
    self.Object.Call("_pauseAll")
}

// Private. Called by game focus loss. Pauses all currently running tweens.
func (self *TweenManager) _pauseAllI(args ...interface{}) {
    self.Object.Call("_pauseAll", args)
}

// Private. Called by game focus loss. Resumes all currently paused tweens.
func (self *TweenManager) _resumeAll() {
    self.Object.Call("_resumeAll")
}

// Private. Called by game focus loss. Resumes all currently paused tweens.
func (self *TweenManager) _resumeAllI(args ...interface{}) {
    self.Object.Call("_resumeAll", args)
}

// Pauses all currently running tweens.
func (self *TweenManager) PauseAll() {
    self.Object.Call("pauseAll")
}

// Pauses all currently running tweens.
func (self *TweenManager) PauseAllI(args ...interface{}) {
    self.Object.Call("pauseAll", args)
}

// Resumes all currently paused tweens.
func (self *TweenManager) ResumeAll() {
    self.Object.Call("resumeAll")
}

// Resumes all currently paused tweens.
func (self *TweenManager) ResumeAllI(args ...interface{}) {
    self.Object.Call("resumeAll", args)
}
