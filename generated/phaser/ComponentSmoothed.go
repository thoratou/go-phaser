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


// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *ComponentSmoothed) GetSmoothed() bool{
    return self.Get("smoothed").Bool()
}

// Enable or disable texture smoothing for this Game Object.
// 
// It only takes effect if the Game Object is using an image based texture.
// 
// Smoothing is enabled by default.
func (self *ComponentSmoothed) SetSmoothed(member bool) {
    self.Set("smoothed", member)
}


