// Automatic generation for Phaser.Component.Delta
// generated file ComponentDelta.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Delta component provides access to delta values between the Game Objects current and previous position.
type ComponentDelta struct {
    *js.Object
}


// Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *ComponentDelta) GetDeltaXA() int{
    return self.Object.Get("deltaX").Int()
}

// Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *ComponentDelta) SetDeltaXA(member int) {
    self.Object.Set("deltaX", member)
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *ComponentDelta) GetDeltaYA() int{
    return self.Object.Get("deltaY").Int()
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *ComponentDelta) SetDeltaYA(member int) {
    self.Object.Set("deltaY", member)
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *ComponentDelta) GetDeltaZA() int{
    return self.Object.Get("deltaZ").Int()
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *ComponentDelta) SetDeltaZA(member int) {
    self.Object.Set("deltaZ", member)
}


