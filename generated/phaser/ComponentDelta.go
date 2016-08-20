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
func (self *ComponentDelta) GetDeltaX() float64{
    return self.Get("deltaX").Float()
}

// Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *ComponentDelta) SetDeltaX(member float64) {
    self.Set("deltaX", member)
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *ComponentDelta) GetDeltaY() float64{
    return self.Get("deltaY").Float()
}

// Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *ComponentDelta) SetDeltaY(member float64) {
    self.Set("deltaY", member)
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *ComponentDelta) GetDeltaZ() float64{
    return self.Get("deltaZ").Float()
}

// Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *ComponentDelta) SetDeltaZ(member float64) {
    self.Set("deltaZ", member)
}


