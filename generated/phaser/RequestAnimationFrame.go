// Package phaser Automatic generation for Phaser.RequestAnimationFrame
// generated file RequestAnimationFrame.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// RequestAnimationFrame Abstracts away the use of RAF or setTimeOut for the core game update loop.
type RequestAnimationFrame struct {
    *js.Object
}

// NewRequestAnimationFrame Abstracts away the use of RAF or setTimeOut for the core game update loop.
func NewRequestAnimationFrame(game *Game) *RequestAnimationFrame {
    return &RequestAnimationFrame{js.Global.Get("Phaser").Get("RequestAnimationFrame").New(game)}
}
// NewRequestAnimationFrame1O Abstracts away the use of RAF or setTimeOut for the core game update loop.
func NewRequestAnimationFrame1O(game *Game, forceSetTimeOut bool) *RequestAnimationFrame {
    return &RequestAnimationFrame{js.Global.Get("Phaser").Get("RequestAnimationFrame").New(game, forceSetTimeOut)}
}
// NewRequestAnimationFrameI Abstracts away the use of RAF or setTimeOut for the core game update loop.
func NewRequestAnimationFrameI(args ...interface{}) *RequestAnimationFrame {
    return &RequestAnimationFrame{js.Global.Get("Phaser").Get("RequestAnimationFrame").New(args)}
}



// RequestAnimationFrame Binding conversion method to RequestAnimationFrame point 
func ToRequestAnimationFrame(jsStruct interface{}) *RequestAnimationFrame {
    if object, ok := jsStruct.(*js.Object); ok {
		return &RequestAnimationFrame{Object: object}
	}
	return nil
}



// Game The currently running game.
func (self *RequestAnimationFrame) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA The currently running game.
func (self *RequestAnimationFrame) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// IsRunning true if RequestAnimationFrame is running, otherwise false.
func (self *RequestAnimationFrame) IsRunning() bool{
    return self.Object.Get("isRunning").Bool()
}

// SetIsRunningA true if RequestAnimationFrame is running, otherwise false.
func (self *RequestAnimationFrame) SetIsRunningA(member bool) {
    self.Object.Set("isRunning", member)
}

// ForceSetTimeOut Tell Phaser to use setTimeOut even if raf is available.
func (self *RequestAnimationFrame) ForceSetTimeOut() bool{
    return self.Object.Get("forceSetTimeOut").Bool()
}

// SetForceSetTimeOutA Tell Phaser to use setTimeOut even if raf is available.
func (self *RequestAnimationFrame) SetForceSetTimeOutA(member bool) {
    self.Object.Set("forceSetTimeOut", member)
}


// Start Starts the requestAnimationFrame running or setTimeout if unavailable in browser
func (self *RequestAnimationFrame) Start() {
    self.Object.Call("start")
}

// StartI Starts the requestAnimationFrame running or setTimeout if unavailable in browser
func (self *RequestAnimationFrame) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// UpdateRAF The update method for the requestAnimationFrame
func (self *RequestAnimationFrame) UpdateRAF() {
    self.Object.Call("updateRAF")
}

// UpdateRAFI The update method for the requestAnimationFrame
func (self *RequestAnimationFrame) UpdateRAFI(args ...interface{}) {
    self.Object.Call("updateRAF", args)
}

// UpdateSetTimeout The update method for the setTimeout.
func (self *RequestAnimationFrame) UpdateSetTimeout() {
    self.Object.Call("updateSetTimeout")
}

// UpdateSetTimeoutI The update method for the setTimeout.
func (self *RequestAnimationFrame) UpdateSetTimeoutI(args ...interface{}) {
    self.Object.Call("updateSetTimeout", args)
}

// Stop Stops the requestAnimationFrame from running.
func (self *RequestAnimationFrame) Stop() {
    self.Object.Call("stop")
}

// StopI Stops the requestAnimationFrame from running.
func (self *RequestAnimationFrame) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// IsSetTimeOut Is the browser using setTimeout?
func (self *RequestAnimationFrame) IsSetTimeOut() bool{
    return self.Object.Call("isSetTimeOut").Bool()
}

// IsSetTimeOutI Is the browser using setTimeout?
func (self *RequestAnimationFrame) IsSetTimeOutI(args ...interface{}) bool{
    return self.Object.Call("isSetTimeOut", args).Bool()
}

// IsRAF Is the browser using requestAnimationFrame?
func (self *RequestAnimationFrame) IsRAF() bool{
    return self.Object.Call("isRAF").Bool()
}

// IsRAFI Is the browser using requestAnimationFrame?
func (self *RequestAnimationFrame) IsRAFI(args ...interface{}) bool{
    return self.Object.Call("isRAF", args).Bool()
}

