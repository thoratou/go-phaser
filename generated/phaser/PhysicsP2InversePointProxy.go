// Automatic generation for Phaser.Physics.P2.InversePointProxy
// generated file PhysicsP2InversePointProxy.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A InversePointProxy is an internal class that allows for direct getter/setter style property access to Arrays and TypedArrays but inverses the values on set.
type PhysicsP2InversePointProxy struct {
    *js.Object
}


// The x property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) GetX() int{
    return self.Get("x").Int()
}

// The x property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) SetX(member int) {
    self.Set("x", member)
}

// The y property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) GetY() int{
    return self.Get("y").Int()
}

// The y property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) SetY(member int) {
    self.Set("y", member)
}

// The x property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) GetMx() int{
    return self.Get("mx").Int()
}

// The x property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) SetMx(member int) {
    self.Set("mx", member)
}

// The y property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) GetMy() int{
    return self.Get("my").Int()
}

// The y property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) SetMy(member int) {
    self.Set("my", member)
}


