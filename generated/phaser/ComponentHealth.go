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
func (self *ComponentHealth) GetHealth() float64{
    return self.Get("health").Float()
}

// The Game Objects health value. This is a handy property for setting and manipulating health on a Game Object.
// 
// It can be used in combination with the `damage` method or modified directly.
func (self *ComponentHealth) SetHealth(member float64) {
    self.Set("health", member)
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *ComponentHealth) GetMaxHealth() float64{
    return self.Get("maxHealth").Float()
}

// The Game Objects maximum health value. This works in combination with the `heal` method to ensure
// the health value never exceeds the maximum.
func (self *ComponentHealth) SetMaxHealth(member float64) {
    self.Set("maxHealth", member)
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *ComponentHealth) GetDamage() interface{}{
    return self.Get("damage")
}

// Damages the Game Object. This removes the given amount of health from the `health` property.
// 
// If health is taken below or is equal to zero then the `kill` method is called.
func (self *ComponentHealth) SetDamage(member interface{}) {
    self.Set("damage", member)
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *ComponentHealth) GetSetHealth() interface{}{
    return self.Get("setHealth")
}

// Sets the health property of the Game Object to the given amount.
// Will never exceed the `maxHealth` value.
func (self *ComponentHealth) SetSetHealth(member interface{}) {
    self.Set("setHealth", member)
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *ComponentHealth) GetHeal() interface{}{
    return self.Get("heal")
}

// Heal the Game Object. This adds the given amount of health to the `health` property.
func (self *ComponentHealth) SetHeal(member interface{}) {
    self.Set("heal", member)
}


