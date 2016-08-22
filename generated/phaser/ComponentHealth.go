// Automatic generation for Phaser.Component.Health
// generated file ComponentHealth.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Health component provides the ability for Game Objects to have a `health` property 
// that can be damaged and reset through game code.
// Requires the LifeSpan component.
type ComponentHealth struct {
    *js.Object
}


// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *ComponentHealth) GetHealthA() int{
    return self.Object.Get("health").Int()
}

// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *ComponentHealth) SetHealthA(member int) {
    self.Object.Set("health", member)
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *ComponentHealth) GetMaxHealthA() int{
    return self.Object.Get("maxHealth").Int()
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *ComponentHealth) SetMaxHealthA(member int) {
    self.Object.Set("maxHealth", member)
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *ComponentHealth) GetDamageA() interface{}{
    return self.Object.Get("damage")
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *ComponentHealth) SetDamageA(member interface{}) {
    self.Object.Set("damage", member)
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *ComponentHealth) GetSetHealthA() interface{}{
    return self.Object.Get("setHealth")
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *ComponentHealth) SetSetHealthA(member interface{}) {
    self.Object.Set("setHealth", member)
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *ComponentHealth) GetHealA() interface{}{
    return self.Object.Get("heal")
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *ComponentHealth) SetHealA(member interface{}) {
    self.Object.Set("heal", member)
}


