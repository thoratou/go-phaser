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


// The Animation Component provides a `play` method, which is a proxy to the `AnimationManager.play` method.
func NewComponentAnimation() *ComponentAnimation {
    return &ComponentAnimation{js.Global.Call("Phaser.Component.Animation")}
}

// The Animation Component provides a `play` method, which is a proxy to the `AnimationManager.play` method.
func NewComponentAnimationI(args ...interface{}) *ComponentAnimation {
    return &ComponentAnimation{js.Global.Call("Phaser.Component.Animation", args)}
}





// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) Play(name string) *Animation{
    return &Animation{self.Object.Call("play", name)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) Play1O(name string, frameRate int) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) Play2O(name string, frameRate int, loop bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) Play3O(name string, frameRate int, loop bool, killOnComplete bool) *Animation{
    return &Animation{self.Object.Call("play", name, frameRate, loop, killOnComplete)}
}

// Plays an Animation.
// 
// The animation should have previously been created via `animations.add`.
// 
// If the animation is already playing calling this again won't do anything.
// If you need to reset an already running animation do so directly on the Animation object itself or via `AnimationManager.stop`.
func (self *ComponentAnimation) PlayI(args ...interface{}) *Animation{
    return &Animation{self.Object.Call("play", args)}
}
