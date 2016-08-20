// Automatic generation for Phaser.Component.Animation
// generated file ComponentAnimation.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Animation Component provides a `play` method, which is a proxy to the `AnimationManager.play` method.
type ComponentAnimation struct {
    *js.Object
}




// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) PlayI(args ...interface{}) Animation{
    return Animation{self.Call("play", args)}
}
