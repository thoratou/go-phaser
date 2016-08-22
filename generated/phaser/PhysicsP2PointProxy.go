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
func (self *PhysicsP2PointProxy) GetX() int{
    return self.Get("x").Int()
}

// The x property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) SetX(member int) {
    self.Set("x", member)
}

// The y property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) GetY() int{
    return self.Get("y").Int()
}

// The y property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) SetY(member int) {
    self.Set("y", member)
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) GetMx() int{
    return self.Get("mx").Int()
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) SetMx(member int) {
    self.Set("mx", member)
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) GetMy() int{
    return self.Get("my").Int()
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) SetMy(member int) {
    self.Set("my", member)
}


