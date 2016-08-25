// Package phaser Automatic generation for Phaser.Physics.P2.InversePointProxy
// generated file PhysicsP2InversePointProxy.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PhysicsP2InversePointProxy A InversePointProxy is an internal class that allows for direct getter/setter style property access to Arrays and TypedArrays but inverses the values on set.
type PhysicsP2InversePointProxy struct {
    *js.Object
}

// NewPhysicsP2InversePointProxy A InversePointProxy is an internal class that allows for direct getter/setter style property access to Arrays and TypedArrays but inverses the values on set.
func NewPhysicsP2InversePointProxy(world *PhysicsP2, destination interface{}) *PhysicsP2InversePointProxy {
    return &PhysicsP2InversePointProxy{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("InversePointProxy").New(world, destination)}
}
// NewPhysicsP2InversePointProxyI A InversePointProxy is an internal class that allows for direct getter/setter style property access to Arrays and TypedArrays but inverses the values on set.
func NewPhysicsP2InversePointProxyI(args ...interface{}) *PhysicsP2InversePointProxy {
    return &PhysicsP2InversePointProxy{js.Global.Get("Phaser").Get("Physics").Get("P2").Get("InversePointProxy").New(args)}
}



// PhysicsP2InversePointProxy Binding conversion method to PhysicsP2InversePointProxy point 
func ToPhysicsP2InversePointProxy(jsStruct interface{}) *PhysicsP2InversePointProxy {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PhysicsP2InversePointProxy{Object: object}
	}
	return nil
}



// X The x property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The x property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The y property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The y property of this InversePointProxy get and set in pixels.
func (self *PhysicsP2InversePointProxy) SetYA(member int) {
    self.Object.Set("y", member)
}

// Mx The x property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) Mx() int{
    return self.Object.Get("mx").Int()
}

// SetMxA The x property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) SetMxA(member int) {
    self.Object.Set("mx", member)
}

// My The y property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) My() int{
    return self.Object.Get("my").Int()
}

// SetMyA The y property of this InversePointProxy get and set in meters.
func (self *PhysicsP2InversePointProxy) SetMyA(member int) {
    self.Object.Set("my", member)
}


