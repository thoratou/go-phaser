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
func (self *PhysicsP2InversePointProxy) GetX() float64{
    return self.Get("x").Float()
}

// The x property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) SetX(member float64) {
    self.Set("x", member)
}

// The y property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) GetY() float64{
    return self.Get("y").Float()
}

// The y property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) SetY(member float64) {
    self.Set("y", member)
}

// The x property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) GetMx() float64{
    return self.Get("mx").Float()
}

// The x property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) SetMx(member float64) {
    self.Set("mx", member)
}

// The y property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) GetMy() float64{
    return self.Get("my").Float()
}

// The y property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) SetMy(member float64) {
    self.Set("my", member)
}


