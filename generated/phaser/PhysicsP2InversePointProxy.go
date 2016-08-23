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


// A InversePointProxy is an internal class that allows for direct getter/setter style property access to Arrays and TypedArrays but inverses the values on set.
func NewPhysicsP2InversePointProxy(world *PhysicsP2, destination interface{}) *PhysicsP2InversePointProxy {
    return &PhysicsP2InversePointProxy{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("InversePointProxy").New(world, destination)}
}

// A InversePointProxy is an internal class that allows for direct getter/setter style property access to Arrays and TypedArrays but inverses the values on set.
func NewPhysicsP2InversePointProxyI(args ...interface{}) *PhysicsP2InversePointProxy {
    return &PhysicsP2InversePointProxy{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("InversePointProxy").New(args)}
}



// The x property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) X() int{
    return self.Object.Get("x").Int()
}

// The x property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) SetXA(member int) {
    self.Object.Set("x", member)
}

// The y property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) Y() int{
    return self.Object.Get("y").Int()
}

// The y property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) SetYA(member int) {
    self.Object.Set("y", member)
}

// The x property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) Mx() int{
    return self.Object.Get("mx").Int()
}

// The x property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) SetMxA(member int) {
    self.Object.Set("mx", member)
}

// The y property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) My() int{
    return self.Object.Get("my").Int()
}

// The y property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) SetMyA(member int) {
    self.Object.Set("my", member)
}


