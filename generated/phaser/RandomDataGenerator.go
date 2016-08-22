// Automatic generation for Phaser.RandomDataGenerator
// generated file RandomDataGenerator.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// An extremely useful repeatable random data generator.
// 
// Based on Nonsense by Josh Faul https://github.com/jocafa/Nonsense.
// 
// The random number genererator is based on the Alea PRNG, but is modified.
//  - https://github.com/coverslide/node-alea
//  - https://github.com/nquinlan/better-random-numbers-for-javascript-mirror
//  - http://baagoe.org/en/wiki/Better_random_numbers_for_javascript (original, perm. 404)
type RandomDataGenerator struct {
    *js.Object
}




// Private random helper.
func (self *RandomDataGenerator) RndI(args ...interface{}) int{
    return self.Call("rnd", args).Int()
}

// Reset the seed of the random data generator.
// 
// _Note_: the seed array is only processed up to the first `undefined` (or `null`) value, should such be present.
func (self *RandomDataGenerator) SowI(args ...interface{}) {
    self.Call("sow", args)
}

// Internal method that creates a seed hash.
func (self *RandomDataGenerator) HashI(args ...interface{}) int{
    return self.Call("hash", args).Int()
}

// Returns a random integer between 0 and 2^32.
func (self *RandomDataGenerator) IntegerI(args ...interface{}) int{
    return self.Call("integer", args).Int()
}

// Returns a random real number between 0 and 1.
func (self *RandomDataGenerator) FracI(args ...interface{}) int{
    return self.Call("frac", args).Int()
}

// Returns a random real number between 0 and 2^32.
func (self *RandomDataGenerator) RealI(args ...interface{}) int{
    return self.Call("real", args).Int()
}

// Returns a random integer between and including min and max.
func (self *RandomDataGenerator) IntegerInRangeI(args ...interface{}) int{
    return self.Call("integerInRange", args).Int()
}

// Returns a random integer between and including min and max.
// This method is an alias for RandomDataGenerator.integerInRange.
func (self *RandomDataGenerator) BetweenI(args ...interface{}) int{
    return self.Call("between", args).Int()
}

// Returns a random real number between min and max.
func (self *RandomDataGenerator) RealInRangeI(args ...interface{}) int{
    return self.Call("realInRange", args).Int()
}

// Returns a random real number between -1 and 1.
func (self *RandomDataGenerator) NormalI(args ...interface{}) int{
    return self.Call("normal", args).Int()
}

// Returns a valid RFC4122 version4 ID hex string from https://gist.github.com/1308368
func (self *RandomDataGenerator) UuidI(args ...interface{}) string{
    return self.Call("uuid", args).String()
}

// Returns a random member of `array`.
func (self *RandomDataGenerator) PickI(args ...interface{}) interface{}{
    return self.Call("pick", args)
}

// Returns a sign to be used with multiplication operator.
func (self *RandomDataGenerator) SignI(args ...interface{}) int{
    return self.Call("sign", args).Int()
}

// Returns a random member of `array`, favoring the earlier entries.
func (self *RandomDataGenerator) WeightedPickI(args ...interface{}) interface{}{
    return self.Call("weightedPick", args)
}

// Returns a random timestamp between min and max, or between the beginning of 2000 and the end of 2020 if min and max aren't specified.
func (self *RandomDataGenerator) TimestampI(args ...interface{}) int{
    return self.Call("timestamp", args).Int()
}

// Returns a random angle between -180 and 180.
func (self *RandomDataGenerator) AngleI(args ...interface{}) int{
    return self.Call("angle", args).Int()
}

// Gets or Sets the state of the generator. This allows you to retain the values
// that the generator is using between games, i.e. in a game save file.
// 
// To seed this generator with a previously saved state you can pass it as the 
// `seed` value in your game config, or call this method directly after Phaser has booted.
// 
// Call this method with no parameters to return the current state.
// 
// If providing a state it should match the same format that this method
// returns, which is a string with a header `!rnd` followed by the `c`,
// `s0`, `s1` and `s2` values respectively, each comma-delimited.
func (self *RandomDataGenerator) StateI(args ...interface{}) string{
    return self.Call("state", args).String()
}
