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
func (self *RequestAnimationFrame) GetGame() *Game{
    return &Game{self.Get("game")}
}

// The currently running game.
func (self *RequestAnimationFrame) SetGame(member *Game) {
    self.Set("game", member)
}

// true if RequestAnimationFrame is running, otherwise false.
func (self *RequestAnimationFrame) GetIsRunning() bool{
    return self.Get("isRunning").Bool()
}

// true if RequestAnimationFrame is running, otherwise false.
func (self *RequestAnimationFrame) SetIsRunning(member bool) {
    self.Set("isRunning", member)
}

// Tell Phaser to use setTimeOut even if raf is available.
func (self *RequestAnimationFrame) GetForceSetTimeOut() bool{
    return self.Get("forceSetTimeOut").Bool()
}

// Tell Phaser to use setTimeOut even if raf is available.
func (self *RequestAnimationFrame) SetForceSetTimeOut(member bool) {
    self.Set("forceSetTimeOut", member)
}



// Starts the requestAnimationFrame running or setTimeout if unavailable in browser
func (self *RequestAnimationFrame) StartI(args ...interface{}) {
    self.Call("start", args)
}

// The update method for the requestAnimationFrame
func (self *RequestAnimationFrame) UpdateRAFI(args ...interface{}) {
    self.Call("updateRAF", args)
}

// The update method for the setTimeout.
func (self *RequestAnimationFrame) UpdateSetTimeoutI(args ...interface{}) {
    self.Call("updateSetTimeout", args)
}

// Stops the requestAnimationFrame from running.
func (self *RequestAnimationFrame) StopI(args ...interface{}) {
    self.Call("stop", args)
}

// Is the browser using setTimeout?
func (self *RequestAnimationFrame) IsSetTimeOutI(args ...interface{}) bool{
    return self.Call("isSetTimeOut", args).Bool()
}

// Is the browser using requestAnimationFrame?
func (self *RequestAnimationFrame) IsRAFI(args ...interface{}) bool{
    return self.Call("isRAF", args).Bool()
}
