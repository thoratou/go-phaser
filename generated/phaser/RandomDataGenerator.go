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


// An extremely useful repeatable random data generator.
// 
// Based on Nonsense by Josh Faul https://github.com/jocafa/Nonsense.
// 
// The random number genererator is based on the Alea PRNG, but is modified.
//  - https://github.com/coverslide/node-alea
//  - https://github.com/nquinlan/better-random-numbers-for-javascript-mirror
//  - http://baagoe.org/en/wiki/Better_random_numbers_for_javascript (original, perm. 404)
func NewRandomDataGenerator() *RandomDataGenerator {
    return &RandomDataGenerator{js.Global.Get("Phaser").Get("RandomDataGenerator").New()}
}

// An extremely useful repeatable random data generator.
// 
// Based on Nonsense by Josh Faul https://github.com/jocafa/Nonsense.
// 
// The random number genererator is based on the Alea PRNG, but is modified.
//  - https://github.com/coverslide/node-alea
//  - https://github.com/nquinlan/better-random-numbers-for-javascript-mirror
//  - http://baagoe.org/en/wiki/Better_random_numbers_for_javascript (original, perm. 404)
func NewRandomDataGenerator1O(seeds interface{}) *RandomDataGenerator {
    return &RandomDataGenerator{js.Global.Get("Phaser").Get("RandomDataGenerator").New(seeds)}
}

// An extremely useful repeatable random data generator.
// 
// Based on Nonsense by Josh Faul https://github.com/jocafa/Nonsense.
// 
// The random number genererator is based on the Alea PRNG, but is modified.
//  - https://github.com/coverslide/node-alea
//  - https://github.com/nquinlan/better-random-numbers-for-javascript-mirror
//  - http://baagoe.org/en/wiki/Better_random_numbers_for_javascript (original, perm. 404)
func NewRandomDataGeneratorI(args ...interface{}) *RandomDataGenerator {
    return &RandomDataGenerator{js.Global.Get("Phaser").Get("RandomDataGenerator").New(args)}
}





// Private random helper.
func (self *RandomDataGenerator) Rnd() int{
    return self.Object.Call("rnd").Int()
}

// Private random helper.
func (self *RandomDataGenerator) RndI(args ...interface{}) int{
    return self.Object.Call("rnd", args).Int()
}

// Reset the seed of the random data generator.
// 
// _Note_: the seed array is only processed up to the first `undefined` (or `null`) value, should such be present.
func (self *RandomDataGenerator) Sow(seeds []interface{}) {
    self.Object.Call("sow", seeds)
}

// Reset the seed of the random data generator.
// 
// _Note_: the seed array is only processed up to the first `undefined` (or `null`) value, should such be present.
func (self *RandomDataGenerator) SowI(args ...interface{}) {
    self.Object.Call("sow", args)
}

// Internal method that creates a seed hash.
func (self *RandomDataGenerator) Hash(data interface{}) int{
    return self.Object.Call("hash", data).Int()
}

// Internal method that creates a seed hash.
func (self *RandomDataGenerator) HashI(args ...interface{}) int{
    return self.Object.Call("hash", args).Int()
}

// Returns a random integer between 0 and 2^32.
func (self *RandomDataGenerator) Integer() int{
    return self.Object.Call("integer").Int()
}

// Returns a random integer between 0 and 2^32.
func (self *RandomDataGenerator) IntegerI(args ...interface{}) int{
    return self.Object.Call("integer", args).Int()
}

// Returns a random real number between 0 and 1.
func (self *RandomDataGenerator) Frac() int{
    return self.Object.Call("frac").Int()
}

// Returns a random real number between 0 and 1.
func (self *RandomDataGenerator) FracI(args ...interface{}) int{
    return self.Object.Call("frac", args).Int()
}

// Returns a random real number between 0 and 2^32.
func (self *RandomDataGenerator) Real() int{
    return self.Object.Call("real").Int()
}

// Returns a random real number between 0 and 2^32.
func (self *RandomDataGenerator) RealI(args ...interface{}) int{
    return self.Object.Call("real", args).Int()
}

// Returns a random integer between and including min and max.
func (self *RandomDataGenerator) IntegerInRange(min int, max int) int{
    return self.Object.Call("integerInRange", min, max).Int()
}

// Returns a random integer between and including min and max.
func (self *RandomDataGenerator) IntegerInRangeI(args ...interface{}) int{
    return self.Object.Call("integerInRange", args).Int()
}

// Returns a random integer between and including min and max.
// This method is an alias for RandomDataGenerator.integerInRange.
func (self *RandomDataGenerator) Between(min int, max int) int{
    return self.Object.Call("between", min, max).Int()
}

// Returns a random integer between and including min and max.
// This method is an alias for RandomDataGenerator.integerInRange.
func (self *RandomDataGenerator) BetweenI(args ...interface{}) int{
    return self.Object.Call("between", args).Int()
}

// Returns a random real number between min and max.
func (self *RandomDataGenerator) RealInRange(min int, max int) int{
    return self.Object.Call("realInRange", min, max).Int()
}

// Returns a random real number between min and max.
func (self *RandomDataGenerator) RealInRangeI(args ...interface{}) int{
    return self.Object.Call("realInRange", args).Int()
}

// Returns a random real number between -1 and 1.
func (self *RandomDataGenerator) Normal() int{
    return self.Object.Call("normal").Int()
}

// Returns a random real number between -1 and 1.
func (self *RandomDataGenerator) NormalI(args ...interface{}) int{
    return self.Object.Call("normal", args).Int()
}

// Returns a valid RFC4122 version4 ID hex string from https://gist.github.com/1308368
func (self *RandomDataGenerator) Uuid() string{
    return self.Object.Call("uuid").String()
}

// Returns a valid RFC4122 version4 ID hex string from https://gist.github.com/1308368
func (self *RandomDataGenerator) UuidI(args ...interface{}) string{
    return self.Object.Call("uuid", args).String()
}

// Returns a random member of `array`.
func (self *RandomDataGenerator) Pick(ary []interface{}) interface{}{
    return self.Object.Call("pick", ary)
}

// Returns a random member of `array`.
func (self *RandomDataGenerator) PickI(args ...interface{}) interface{}{
    return self.Object.Call("pick", args)
}

// Returns a sign to be used with multiplication operator.
func (self *RandomDataGenerator) Sign() int{
    return self.Object.Call("sign").Int()
}

// Returns a sign to be used with multiplication operator.
func (self *RandomDataGenerator) SignI(args ...interface{}) int{
    return self.Object.Call("sign", args).Int()
}

// Returns a random member of `array`, favoring the earlier entries.
func (self *RandomDataGenerator) WeightedPick(ary []interface{}) interface{}{
    return self.Object.Call("weightedPick", ary)
}

// Returns a random member of `array`, favoring the earlier entries.
func (self *RandomDataGenerator) WeightedPickI(args ...interface{}) interface{}{
    return self.Object.Call("weightedPick", args)
}

// Returns a random timestamp between min and max, or between the beginning of 2000 and the end of 2020 if min and max aren't specified.
func (self *RandomDataGenerator) Timestamp(min int, max int) int{
    return self.Object.Call("timestamp", min, max).Int()
}

// Returns a random timestamp between min and max, or between the beginning of 2000 and the end of 2020 if min and max aren't specified.
func (self *RandomDataGenerator) TimestampI(args ...interface{}) int{
    return self.Object.Call("timestamp", args).Int()
}

// Returns a random angle between -180 and 180.
func (self *RandomDataGenerator) Angle() int{
    return self.Object.Call("angle").Int()
}

// Returns a random angle between -180 and 180.
func (self *RandomDataGenerator) AngleI(args ...interface{}) int{
    return self.Object.Call("angle", args).Int()
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
func (self *RandomDataGenerator) State() string{
    return self.Object.Call("state").String()
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
func (self *RandomDataGenerator) State1O(state string) string{
    return self.Object.Call("state", state).String()
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
    return self.Object.Call("state", args).String()
}
