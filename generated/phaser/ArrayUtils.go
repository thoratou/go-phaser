// Automatic generation for Phaser.ArrayUtils
// generated file ArrayUtils.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Utility functions for dealing with Arrays.
type ArrayUtils struct {
    *js.Object
}




// Fetch a random entry from the given array.
// 
// Will return null if there are no array items that fall within the specified range
// or if there is no item for the randomly chosen index.
func (self *ArrayUtils) GetRandomItemI(args ...interface{}) interface{}{
    return self.Call("getRandomItem", args)
}

// Removes a random object from the given array and returns it.
// 
// Will return null if there are no array items that fall within the specified range
// or if there is no item for the randomly chosen index.
func (self *ArrayUtils) RemoveRandomItemI(args ...interface{}) interface{}{
    return self.Call("removeRandomItem", args)
}

// A standard Fisher-Yates Array shuffle implementation which modifies the array in place.
func (self *ArrayUtils) ShuffleI(args ...interface{}) []interface{}{
	array := self.Call("shuffle", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Transposes the elements of the given matrix (array of arrays).
func (self *ArrayUtils) TransposeMatrixI(args ...interface{}) [][]interface{}{
	array := self.Call("transposeMatrix", args)
	length := array.Length()
	out := make([][]interface{}, length, length)
	for i := 0; i < length; i++ {
		
			out[i] = []interface{}{array.Index(i)}
	}
	return out
}

// Rotates the given matrix (array of arrays).
// 
// Based on the routine from {@link http://jsfiddle.net/MrPolywhirl/NH42z/}.
func (self *ArrayUtils) RotateMatrixI(args ...interface{}) [][]interface{}{
	array := self.Call("rotateMatrix", args)
	length := array.Length()
	out := make([][]interface{}, length, length)
	for i := 0; i < length; i++ {
		
			out[i] = []interface{}{array.Index(i)}
	}
	return out
}

// Snaps a value to the nearest value in an array.
// The result will always be in the range `[first_value, last_value]`.
func (self *ArrayUtils) FindClosestI(args ...interface{}) float64{
    return self.Call("findClosest", args).Float()
}

// Moves the element from the end of the array to the start, shifting all items in the process.
// The "rotation" happens to the right.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ F, A, B, C, D, E ]`
// 
// See also Phaser.ArrayUtils.rotateLeft.
func (self *ArrayUtils) RotateRightI(args ...interface{}) interface{}{
    return self.Call("rotateRight", args)
}

// Moves the element from the start of the array to the end, shifting all items in the process.
// The "rotation" happens to the left.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ B, C, D, E, F, A ]`
// 
// See also Phaser.ArrayUtils.rotateRight
func (self *ArrayUtils) RotateLeftI(args ...interface{}) interface{}{
    return self.Call("rotateLeft", args)
}

// Moves the element from the start of the array to the end, shifting all items in the process.
// The "rotation" happens to the left.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ B, C, D, E, F, A ]`
// 
// See also Phaser.ArrayUtils.rotateRight
func (self *ArrayUtils) RotateI(args ...interface{}) interface{}{
    return self.Call("rotate", args)
}

// Create an array representing the inclusive range of numbers (usually integers) in `[start, end]`.
// This is equivalent to `numberArrayStep(start, end, 1)`.
func (self *ArrayUtils) NumberArrayI(args ...interface{}) []float64{
	array := self.Call("numberArray", args)
	length := array.Length()
	out := make([]float64, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Float()
		
	}
	return out
}

// Create an array of numbers (positive and/or negative) progressing from `start`
// up to but not including `end` by advancing by `step`.
// 
// If `start` is less than `end` a zero-length range is created unless a negative `step` is specified.
// 
// Certain values for `start` and `end` (eg. NaN/undefined/null) are currently coerced to 0;
// for forward compatibility make sure to pass in actual numbers.
func (self *ArrayUtils) NumberArrayStepI(args ...interface{}) []interface{}{
	array := self.Call("numberArrayStep", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}
