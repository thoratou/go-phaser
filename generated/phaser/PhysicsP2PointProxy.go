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


// A PointProxy is an internal class that allows for direct getter/setter style property access to Arrays and TypedArrays.
func NewPhysicsP2PointProxy(world *PhysicsP2, destination interface{}) *PhysicsP2PointProxy {
    return &PhysicsP2PointProxy{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("PointProxy").New(world, destination)}
}

// A PointProxy is an internal class that allows for direct getter/setter style property access to Arrays and TypedArrays.
func NewPhysicsP2PointProxyI(args ...interface{}) *PhysicsP2PointProxy {
    return &PhysicsP2PointProxy{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("PointProxy").New(args)}
}



// The x property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) X() int{
    return self.Object.Get("x").Int()
}

// The x property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) SetXA(member int) {
    self.Object.Set("x", member)
}

// The y property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) Y() int{
    return self.Object.Get("y").Int()
}

// The y property of this PointProxy get and set in pixels.
func (self *PhysicsP2PointProxy) SetYA(member int) {
    self.Object.Set("y", member)
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) Mx() int{
    return self.Object.Get("mx").Int()
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) SetMxA(member int) {
    self.Object.Set("mx", member)
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) My() int{
    return self.Object.Get("my").Int()
}

// The x property of this PointProxy get and set in meters.
func (self *PhysicsP2PointProxy) SetMyA(member int) {
    self.Object.Set("my", member)
}


