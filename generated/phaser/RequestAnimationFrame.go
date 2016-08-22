// Automatic generation for Phaser.RequestAnimationFrame
// generated file RequestAnimationFrame.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Abstracts away the use of RAF or setTimeOut for the core game update loop.
type RequestAnimationFrame struct {
    *js.Object
}


// The currently running game.
func (self *RequestAnimationFrame) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// The currently running game.
func (self *RequestAnimationFrame) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// true if RequestAnimationFrame is running, otherwise false.
func (self *RequestAnimationFrame) GetIsRunningA() bool{
    return self.Object.Get("isRunning").Bool()
}

// true if RequestAnimationFrame is running, otherwise false.
func (self *RequestAnimationFrame) SetIsRunningA(member bool) {
    self.Object.Set("isRunning", member)
}

// Tell Phaser to use setTimeOut even if raf is available.
func (self *RequestAnimationFrame) GetForceSetTimeOutA() bool{
    return self.Object.Get("forceSetTimeOut").Bool()
}

// Tell Phaser to use setTimeOut even if raf is available.
func (self *RequestAnimationFrame) SetForceSetTimeOutA(member bool) {
    self.Object.Set("forceSetTimeOut", member)
}



// Starts the requestAnimationFrame running or setTimeout if unavailable in browser
func (self *RequestAnimationFrame) Start() {
    self.Object.Call("start")
}

// Starts the requestAnimationFrame running or setTimeout if unavailable in browser
func (self *RequestAnimationFrame) StartI(args ...interface{}) {
    self.Object.Call("start", args)
}

// The update method for the requestAnimationFrame
func (self *RequestAnimationFrame) UpdateRAF() {
    self.Object.Call("updateRAF")
}

// The update method for the requestAnimationFrame
func (self *RequestAnimationFrame) UpdateRAFI(args ...interface{}) {
    self.Object.Call("updateRAF", args)
}

// The update method for the setTimeout.
func (self *RequestAnimationFrame) UpdateSetTimeout() {
    self.Object.Call("updateSetTimeout")
}

// The update method for the setTimeout.
func (self *RequestAnimationFrame) UpdateSetTimeoutI(args ...interface{}) {
    self.Object.Call("updateSetTimeout", args)
}

// Stops the requestAnimationFrame from running.
func (self *RequestAnimationFrame) Stop() {
    self.Object.Call("stop")
}

// Stops the requestAnimationFrame from running.
func (self *RequestAnimationFrame) StopI(args ...interface{}) {
    self.Object.Call("stop", args)
}

// Is the browser using setTimeout?
func (self *RequestAnimationFrame) IsSetTimeOut() bool{
    return self.Object.Call("isSetTimeOut").Bool()
}

// Is the browser using setTimeout?
func (self *RequestAnimationFrame) IsSetTimeOutI(args ...interface{}) bool{
    return self.Object.Call("isSetTimeOut", args).Bool()
}

// Is the browser using requestAnimationFrame?
func (self *RequestAnimationFrame) IsRAF() bool{
    return self.Object.Call("isRAF").Bool()
}

// Is the browser using requestAnimationFrame?
func (self *RequestAnimationFrame) IsRAFI(args ...interface{}) bool{
    return self.Object.Call("isRAF", args).Bool()
}
