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
func (self *ComponentDelta) GetDeltaX() int{
    return self.Get("deltaX").Int()
}

// Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *ComponentDelta) SetDeltaX(member int) {
    self.Set("deltaX", member)
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *ComponentDelta) GetDeltaY() int{
    return self.Get("deltaY").Int()
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *ComponentDelta) SetDeltaY(member int) {
    self.Set("deltaY", member)
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *ComponentDelta) GetDeltaZ() int{
    return self.Get("deltaZ").Int()
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *ComponentDelta) SetDeltaZ(member int) {
    self.Set("deltaZ", member)
}


