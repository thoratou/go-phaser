// Automatic generation for Phaser.Component.PhysicsBody
// generated file ComponentPhysicsBody.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The PhysicsBody component manages the Game Objects physics body and physics enabling.
// It also overrides the x and y properties, ensuring that any manual adjustment of them is reflected in the physics body itself.
type ComponentPhysicsBody struct {
    *js.Object
}


// `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
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
func (self *ComponentPhysicsBody) GetBody() interface{}{
    return self.Get("body")
}

// `body` is the Game Objects physics body. Once a Game Object is enabled for physics you access all associated 
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
func (self *ComponentPhysicsBody) SetBody(member interface{}) {
    self.Set("body", member)
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *ComponentPhysicsBody) GetX() int{
    return self.Get("x").Int()
}

// The position of the Game Object on the x axis relative to the local coordinates of the parent.
func (self *ComponentPhysicsBody) SetX(member int) {
    self.Set("x", member)
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *ComponentPhysicsBody) GetY() int{
    return self.Get("y").Int()
}

// The position of the Game Object on the y axis relative to the local coordinates of the parent.
func (self *ComponentPhysicsBody) SetY(member int) {
    self.Set("y", member)
}



// The PhysicsBody component preUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentPhysicsBody) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// The PhysicsBody component postUpdate handler.
// Called automatically by the Game Object.
func (self *ComponentPhysicsBody) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}
