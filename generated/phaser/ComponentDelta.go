// Package phaser Automatic generation for Phaser.Component.Delta
// generated file ComponentDelta.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ComponentDelta The Delta component provides access to delta values between the Game Objects current and previous position.
type ComponentDelta struct {
    *js.Object
}

// NewComponentDelta The Delta component provides access to delta values between the Game Objects current and previous position.
func NewComponentDelta() *ComponentDelta {
    return &ComponentDelta{js.Global.Get("Phaser").Get("Component").Get("Delta").New()}
}
// NewComponentDeltaI The Delta component provides access to delta values between the Game Objects current and previous position.
func NewComponentDeltaI(args ...interface{}) *ComponentDelta {
    return &ComponentDelta{js.Global.Get("Phaser").Get("Component").Get("Delta").New(args)}
}



// ComponentDelta Binding conversion method to ComponentDelta point 
func ToComponentDelta(jsStruct interface{}) *ComponentDelta {
    if object, ok := jsStruct.(*js.Object); ok {
		return &ComponentDelta{Object: object}
	}
	return nil
}



// DeltaX Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *ComponentDelta) DeltaX() int{
    return self.Object.Get("deltaX").Int()
}

// SetDeltaXA Returns the delta x value. The difference between world.x now and in the previous frame.
// 
// The value will be positive if the Game Object has moved to the right or negative if to the left.
func (self *ComponentDelta) SetDeltaXA(member int) {
    self.Object.Set("deltaX", member)
}

// DeltaY Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *ComponentDelta) DeltaY() int{
    return self.Object.Get("deltaY").Int()
}

// SetDeltaYA Returns the delta y value. The difference between world.y now and in the previous frame.
// 
// The value will be positive if the Game Object has moved down or negative if up.
func (self *ComponentDelta) SetDeltaYA(member int) {
    self.Object.Set("deltaY", member)
}

// DeltaZ Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *ComponentDelta) DeltaZ() int{
    return self.Object.Get("deltaZ").Int()
}

// SetDeltaZA Returns the delta z value. The difference between rotation now and in the previous frame. The delta value.
func (self *ComponentDelta) SetDeltaZA(member int) {
    self.Object.Set("deltaZ", member)
}


