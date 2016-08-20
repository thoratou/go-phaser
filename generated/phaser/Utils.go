// Automatic generation for Phaser.Utils
// generated file Utils.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// 
type Utils struct {
    *js.Object
}




// Takes the given string and reverses it, returning the reversed string.
// For example if given the string `Atari 520ST` it would return `TS025 iratA`.
func (self *Utils) ReverseStringI(args ...interface{}) string{
    return self.Call("reverseString", args).String()
}

// Gets an objects property by string.
func (self *Utils) GetPropertyI(args ...interface{}) interface{}{
    return self.Call("getProperty", args)
}

// Sets an objects property by string.
func (self *Utils) SetPropertyI(args ...interface{}) interface{}{
    return self.Call("setProperty", args)
}

// Generate a random bool result based on the chance value.
// 
// Returns true or false based on the chance value (default 50%). For example if you wanted a player to have a 30% chance
// of getting a bonus, call chanceRoll(30) - true means the chance passed, false means it failed.
func (self *Utils) ChanceRollI(args ...interface{}) bool{
    return self.Call("chanceRoll", args).Bool()
}

// Choose between one of two values randomly.
func (self *Utils) RandomChoiceI(args ...interface{}) interface{}{
    return self.Call("randomChoice", args)
}

// Get a unit dimension from a string.
func (self *Utils) ParseDimensionI(args ...interface{}) float64{
    return self.Call("parseDimension", args).Float()
}

// Takes the given string and pads it out, to the length required, using the character
// specified. For example if you need a string to be 6 characters long, you can call:
// 
// `pad('bob', 6, '-', 2)`
// 
// This would return: `bob---` as it has padded it out to 6 characters, using the `-` on the right.
// 
// You can also use it to pad numbers (they are always returned as strings):
// 
// `pad(512, 6, '0', 1)`
// 
// Would return: `000512` with the string padded to the left.
// 
// If you don't specify a direction it'll pad to both sides:
// 
// `pad('c64', 7, '*')`
// 
// Would return: `**c64**`
func (self *Utils) PadI(args ...interface{}) string{
    return self.Call("pad", args).String()
}

// This is a slightly modified version of jQuery.isPlainObject.
// A plain object is an object whose internal class property is [object Object].
func (self *Utils) IsPlainObjectI(args ...interface{}) bool{
    return self.Call("isPlainObject", args).Bool()
}

// This is a slightly modified version of http://api.jquery.com/jQuery.extend/
func (self *Utils) ExtendI(args ...interface{}) interface{}{
    return self.Call("extend", args)
}

// Mixes in an existing mixin object with the target.
// 
// Values in the mixin that have either `get` or `set` functions are created as properties via `defineProperty`
// _except_ if they also define a `clone` method - if a clone method is defined that is called instead and
// the result is assigned directly.
func (self *Utils) MixinPrototypeI(args ...interface{}) {
    self.Call("mixinPrototype", args)
}

// Mixes the source object into the destination object, returning the newly modified destination object.
// Based on original code by @mudcube
func (self *Utils) MixinI(args ...interface{}) interface{}{
    return self.Call("mixin", args)
}