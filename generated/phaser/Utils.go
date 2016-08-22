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


// 
func NewUtils() *Utils {
    return &Utils{js.Global.Call("Phaser.Utils")}
}

// 
func NewUtilsI(args ...interface{}) *Utils {
    return &Utils{js.Global.Call("Phaser.Utils", args)}
}





// Takes the given string and reverses it, returning the reversed string.
// For example if given the string `Atari 520ST` it would return `TS025 iratA`.
func (self *Utils) ReverseString(string string) string{
    return self.Object.Call("reverseString", string).String()
}

// Takes the given string and reverses it, returning the reversed string.
// For example if given the string `Atari 520ST` it would return `TS025 iratA`.
func (self *Utils) ReverseStringI(args ...interface{}) string{
    return self.Object.Call("reverseString", args).String()
}

// Gets an objects property by string.
func (self *Utils) GetProperty(obj interface{}, prop string) interface{}{
    return self.Object.Call("getProperty", obj, prop)
}

// Gets an objects property by string.
func (self *Utils) GetPropertyI(args ...interface{}) interface{}{
    return self.Object.Call("getProperty", args)
}

// Sets an objects property by string.
func (self *Utils) SetProperty(obj interface{}, prop string) interface{}{
    return self.Object.Call("setProperty", obj, prop)
}

// Sets an objects property by string.
func (self *Utils) SetPropertyI(args ...interface{}) interface{}{
    return self.Object.Call("setProperty", args)
}

// Generate a random bool result based on the chance value.
// 
// Returns true or false based on the chance value (default 50%). For example if you wanted a player to have a 30% chance
// of getting a bonus, call chanceRoll(30) - true means the chance passed, false means it failed.
func (self *Utils) ChanceRoll(chance int) bool{
    return self.Object.Call("chanceRoll", chance).Bool()
}

// Generate a random bool result based on the chance value.
// 
// Returns true or false based on the chance value (default 50%). For example if you wanted a player to have a 30% chance
// of getting a bonus, call chanceRoll(30) - true means the chance passed, false means it failed.
func (self *Utils) ChanceRollI(args ...interface{}) bool{
    return self.Object.Call("chanceRoll", args).Bool()
}

// Choose between one of two values randomly.
func (self *Utils) RandomChoice(choice1 interface{}, choice2 interface{}) interface{}{
    return self.Object.Call("randomChoice", choice1, choice2)
}

// Choose between one of two values randomly.
func (self *Utils) RandomChoiceI(args ...interface{}) interface{}{
    return self.Object.Call("randomChoice", args)
}

// Get a unit dimension from a string.
func (self *Utils) ParseDimension(size interface{}, dimension int) int{
    return self.Object.Call("parseDimension", size, dimension).Int()
}

// Get a unit dimension from a string.
func (self *Utils) ParseDimensionI(args ...interface{}) int{
    return self.Object.Call("parseDimension", args).Int()
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
func (self *Utils) Pad(str string) string{
    return self.Object.Call("pad", str).String()
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
func (self *Utils) Pad1O(str string, len int) string{
    return self.Object.Call("pad", str, len).String()
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
func (self *Utils) Pad2O(str string, len int, pad string) string{
    return self.Object.Call("pad", str, len, pad).String()
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
func (self *Utils) Pad3O(str string, len int, pad string, dir int) string{
    return self.Object.Call("pad", str, len, pad, dir).String()
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
    return self.Object.Call("pad", args).String()
}

// This is a slightly modified version of jQuery.isPlainObject.
// A plain object is an object whose internal class property is [object Object].
func (self *Utils) IsPlainObject(obj interface{}) bool{
    return self.Object.Call("isPlainObject", obj).Bool()
}

// This is a slightly modified version of jQuery.isPlainObject.
// A plain object is an object whose internal class property is [object Object].
func (self *Utils) IsPlainObjectI(args ...interface{}) bool{
    return self.Object.Call("isPlainObject", args).Bool()
}

// This is a slightly modified version of http://api.jquery.com/jQuery.extend/
func (self *Utils) Extend(deep bool, target interface{}) interface{}{
    return self.Object.Call("extend", deep, target)
}

// This is a slightly modified version of http://api.jquery.com/jQuery.extend/
func (self *Utils) ExtendI(args ...interface{}) interface{}{
    return self.Object.Call("extend", args)
}

// Mixes in an existing mixin object with the target.
// 
// Values in the mixin that have either `get` or `set` functions are created as properties via `defineProperty`
// _except_ if they also define a `clone` method - if a clone method is defined that is called instead and
// the result is assigned directly.
func (self *Utils) MixinPrototype(target interface{}, mixin interface{}) {
    self.Object.Call("mixinPrototype", target, mixin)
}

// Mixes in an existing mixin object with the target.
// 
// Values in the mixin that have either `get` or `set` functions are created as properties via `defineProperty`
// _except_ if they also define a `clone` method - if a clone method is defined that is called instead and
// the result is assigned directly.
func (self *Utils) MixinPrototype1O(target interface{}, mixin interface{}, replace bool) {
    self.Object.Call("mixinPrototype", target, mixin, replace)
}

// Mixes in an existing mixin object with the target.
// 
// Values in the mixin that have either `get` or `set` functions are created as properties via `defineProperty`
// _except_ if they also define a `clone` method - if a clone method is defined that is called instead and
// the result is assigned directly.
func (self *Utils) MixinPrototypeI(args ...interface{}) {
    self.Object.Call("mixinPrototype", args)
}

// Mixes the source object into the destination object, returning the newly modified destination object.
// Based on original code by @mudcube
func (self *Utils) Mixin(from interface{}, to interface{}) interface{}{
    return self.Object.Call("mixin", from, to)
}

// Mixes the source object into the destination object, returning the newly modified destination object.
// Based on original code by @mudcube
func (self *Utils) MixinI(args ...interface{}) interface{}{
    return self.Object.Call("mixin", args)
}
