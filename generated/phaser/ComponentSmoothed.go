// Automatic generation for Phaser.Component.Smoothed
// generated file ComponentSmoothed.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Smoothed component allows a Game Object to control anti-aliasing of an image based texture.
type ComponentSmoothed struct {
    *js.Object
}


// The Smoothed component allows a Game Object to control anti-aliasing of an image based texture.
func NewComponentSmoothed() *ComponentSmoothed {
    return &ComponentSmoothed{js.Global.Get("Phaser").Get("Component").Get("Smoothed").New()}
}

// The Smoothed component allows a Game Object to control anti-aliasing of an image based texture.
func NewComponentSmoothedI(args ...interface{}) *ComponentSmoothed {
    return &ComponentSmoothed{js.Global.Get("Phaser").Get("Component").Get("Smoothed").New(args)}
}



// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *ComponentSmoothed) Smoothed() bool{
    return self.Object.Get("smoothed").Bool()
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *ComponentSmoothed) SetSmoothedA(member bool) {
    self.Object.Set("smoothed", member)
}


