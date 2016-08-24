// Package phaser Automatic generation for Phaser.Component.PhysicsBody
// generated file ComponentPhysicsBody.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ComponentPhysicsBody The PhysicsBody component manages the Game Objects physics body and physics enabling.
// It also overrides the x and y properties, ensuring that any manual adjustment of them is reflected in the physics body itself.
type ComponentPhysicsBody struct {
    *js.Object
}

// NewComponentPhysicsBody The PhysicsBody component manages the Game Objects physics body and physics enabling.
// It also overrides the x and y properties, ensuring that any manual adjustment of them is reflected in the physics body itself.
func NewComponentPhysicsBody() *ComponentPhysicsBody {
    return &ComponentPhysicsBody{js.Global.Get("Phaser").Get("Component").Get("PhysicsBody").New()}
}
// NewComponentPhysicsBodyI The PhysicsBody component manages the Game Objects physics body and physics enabling.
// It also overrides the x and y properties, ensuring that any manual adjustment of them is reflected in the physics body itself.
func NewComponentPhysicsBodyI(args ...interface{}) *ComponentPhysicsBody {
    return &ComponentPhysicsBody{js.Global.Get("Phaser").Get("Component").Get("PhysicsBody").New(args)}
}



// Body `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
// properties and methods via it.
// 
// By default Game Objects won't add themselves to any physics system and their `body` property will be `null`.
// 
// To enable this Game Object for physics you need to call `game.physics.enable(object, system)` where `object` is this object
// and `system` is the Physics system you are using. If none is given it defaults to `Phaser.Physics.Arcade`.
// 
// You can alternatively call `game.physics.arcade.enable(object)`, or add this Game Object to a physics enabled Group.
// 
// Important: Enabling a Game Object for P2 or Ninja physics will automatically set its `anchor` property to 0.5, 
// so the physics body is centered on the Game Object.
// 
// If you need a different result then adjust or re-create the Body shape offsets manually or reset the anchor after enabling physics.
func (self *ComponentPhysicsBody) Body() interface{}{
    return self.Object.Get("body")
}

// SetBodyA `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
// properties and methods via it.
// 
// By default Game Objects won't add themselves to any physics system and their `body` property will be `null`.
// 
// To enable this Game Object for physics you need to call `game.physics.enable(object, system)` where `object` is this object
// and `system` is the Physics system you are using. If none is given it defaults to `Phaser.Physics.Arcade`.
// 
// You can alternatively call `game.physics.arcade.enable(object)`, or add this Game Object to a physics enabled Group.
// 
// Important: Enabling a Game Object for P2 or Ninja physics will automatically set its `anchor` property to 0.5, 
// so the physics body is centered on the Game Object.
// 
// If you need a different result then adjust or re-create the Body shape offsets manually or reset the anchor after enabling physics.
func (self *ComponentPhysicsBody) SetBodyA(member interface{}) {
    self.Object.Set("body", member)
}

// X The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *ComponentPhysicsBody) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *ComponentPhysicsBody) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *ComponentPhysicsBody) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *ComponentPhysicsBody) SetYA(member int) {
    self.Object.Set("y", member)
}


// PreUpdate The PhysicsBody component preUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentPhysicsBody) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI The PhysicsBody component preUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentPhysicsBody) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// PostUpdate The PhysicsBody component postUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentPhysicsBody) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI The PhysicsBody component postUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentPhysicsBody) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

