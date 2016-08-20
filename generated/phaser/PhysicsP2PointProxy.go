// Automatic generation for Phaser.Physics.P2.PointProxy
// generated file PhysicsP2PointProxy.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A PointProxy is an internal class that allows for direct getter/setter style property access to Arrays and TypedArrays.
type PhysicsP2PointProxy struct {
    *js.Object
}


// The x property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) GetX() float64{
    return self.Get("x").Float()
}

// The x property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) SetX(member float64) {
    self.Set("x", member)
}

// The y property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) GetY() float64{
    return self.Get("y").Float()
}

// The y property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) SetY(member float64) {
    self.Set("y", member)
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) GetMx() float64{
    return self.Get("mx").Float()
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) SetMx(member float64) {
    self.Set("mx", member)
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) GetMy() float64{
    return self.Get("my").Float()
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) SetMy(member float64) {
    self.Set("my", member)
}


