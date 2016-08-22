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


// Utility functions for dealing with Arrays.
func NewArrayUtils() *ArrayUtils {
    return &ArrayUtils{js.Global.Call("Phaser.ArrayUtils")}
}

// Utility functions for dealing with Arrays.
func NewArrayUtilsI(args ...interface{}) *ArrayUtils {
    return &ArrayUtils{js.Global.Call("Phaser.ArrayUtils", args)}
}





// Fetch a random entry from the given array.
// 
// Will return null if there are no array items that fall within the specified range
// or if there is no item for the randomly chosen index.
func (self *ArrayUtils) GetRandomItem(objects []interface{}, startIndex int, length int) interface{}{
    return self.Object.Call("getRandomItem", objects, startIndex, length)
}

// Fetch a random entry from the given array.
// 
// Will return null if there are no array items that fall within the specified range
// or if there is no item for the randomly chosen index.
func (self *ArrayUtils) GetRandomItemI(args ...interface{}) interface{}{
    return self.Object.Call("getRandomItem", args)
}

// Removes a random object from the given array and returns it.
// 
// Will return null if there are no array items that fall within the specified range
// or if there is no item for the randomly chosen index.
func (self *ArrayUtils) RemoveRandomItem(objects []interface{}, startIndex int, length int) interface{}{
    return self.Object.Call("removeRandomItem", objects, startIndex, length)
}

// Removes a random object from the given array and returns it.
// 
// Will return null if there are no array items that fall within the specified range
// or if there is no item for the randomly chosen index.
func (self *ArrayUtils) RemoveRandomItemI(args ...interface{}) interface{}{
    return self.Object.Call("removeRandomItem", args)
}

// A standard Fisher-Yates Array shuffle implementation which modifies the array in place.
func (self *ArrayUtils) Shuffle(array []interface{}) []interface{}{
	array00 := self.Object.Call("shuffle", array)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// A standard Fisher-Yates Array shuffle implementation which modifies the array in place.
func (self *ArrayUtils) ShuffleI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("shuffle", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Transposes the elements of the given matrix (array of arrays).
func (self *ArrayUtils) TransposeMatrix(array [][]interface{}) [][]interface{}{
	array00 := self.Object.Call("transposeMatrix", array)
	length00 := array00.Length()
	out00 := make([][]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = []interface{}{array00.Index(i00)}
	}
	return out00
}

// Transposes the elements of the given matrix (array of arrays).
func (self *ArrayUtils) TransposeMatrixI(args ...interface{}) [][]interface{}{
	array00 := self.Object.Call("transposeMatrix", args)
	length00 := array00.Length()
	out00 := make([][]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = []interface{}{array00.Index(i00)}
	}
	return out00
}

// Rotates the given matrix (array of arrays).
// 
// Based on the routine from {@link http://jsfiddle.net/MrPolywhirl/NH42z/}.
func (self *ArrayUtils) RotateMatrix(matrix [][]interface{}, direction interface{}) [][]interface{}{
	array00 := self.Object.Call("rotateMatrix", matrix, direction)
	length00 := array00.Length()
	out00 := make([][]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = []interface{}{array00.Index(i00)}
	}
	return out00
}

// Rotates the given matrix (array of arrays).
// 
// Based on the routine from {@link http://jsfiddle.net/MrPolywhirl/NH42z/}.
func (self *ArrayUtils) RotateMatrixI(args ...interface{}) [][]interface{}{
	array00 := self.Object.Call("rotateMatrix", args)
	length00 := array00.Length()
	out00 := make([][]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = []interface{}{array00.Index(i00)}
	}
	return out00
}

// Snaps a value to the nearest value in an array.
// The result will always be in the range `[first_value, last_value]`.
func (self *ArrayUtils) FindClosest(value int, arr []int) int{
    return self.Object.Call("findClosest", value, arr).Int()
}

// Snaps a value to the nearest value in an array.
// The result will always be in the range `[first_value, last_value]`.
func (self *ArrayUtils) FindClosestI(args ...interface{}) int{
    return self.Object.Call("findClosest", args).Int()
}

// Moves the element from the end of the array to the start, shifting all items in the process.
// The "rotation" happens to the right.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ F, A, B, C, D, E ]`
// 
// See also Phaser.ArrayUtils.rotateLeft.
func (self *ArrayUtils) RotateRight(array []interface{}) interface{}{
    return self.Object.Call("rotateRight", array)
}

// Moves the element from the end of the array to the start, shifting all items in the process.
// The "rotation" happens to the right.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ F, A, B, C, D, E ]`
// 
// See also Phaser.ArrayUtils.rotateLeft.
func (self *ArrayUtils) RotateRightI(args ...interface{}) interface{}{
    return self.Object.Call("rotateRight", args)
}

// Moves the element from the start of the array to the end, shifting all items in the process.
// The "rotation" happens to the left.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ B, C, D, E, F, A ]`
// 
// See also Phaser.ArrayUtils.rotateRight
func (self *ArrayUtils) RotateLeft(array []interface{}) interface{}{
    return self.Object.Call("rotateLeft", array)
}

// Moves the element from the start of the array to the end, shifting all items in the process.
// The "rotation" happens to the left.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ B, C, D, E, F, A ]`
// 
// See also Phaser.ArrayUtils.rotateRight
func (self *ArrayUtils) RotateLeftI(args ...interface{}) interface{}{
    return self.Object.Call("rotateLeft", args)
}

// Moves the element from the start of the array to the end, shifting all items in the process.
// The "rotation" happens to the left.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ B, C, D, E, F, A ]`
// 
// See also Phaser.ArrayUtils.rotateRight
func (self *ArrayUtils) Rotate(array []interface{}) interface{}{
    return self.Object.Call("rotate", array)
}

// Moves the element from the start of the array to the end, shifting all items in the process.
// The "rotation" happens to the left.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ B, C, D, E, F, A ]`
// 
// See also Phaser.ArrayUtils.rotateRight
func (self *ArrayUtils) RotateI(args ...interface{}) interface{}{
    return self.Object.Call("rotate", args)
}

// Create an array representing the inclusive range of numbers (usually integers) in `[start, end]`.
// This is equivalent to `numberArrayStep(start, end, 1)`.
func (self *ArrayUtils) NumberArray(start int, end int) []int{
	array00 := self.Object.Call("numberArray", start, end)
	length00 := array00.Length()
	out00 := make([]int, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Int()
		
	}
	return out00
}

// Create an array representing the inclusive range of numbers (usually integers) in `[start, end]`.
// This is equivalent to `numberArrayStep(start, end, 1)`.
func (self *ArrayUtils) NumberArrayI(args ...interface{}) []int{
	array00 := self.Object.Call("numberArray", args)
	length00 := array00.Length()
	out00 := make([]int, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Int()
		
	}
	return out00
}

// Create an array of numbers (positive and/or negative) progressing from `start`
// up to but not including `end` by advancing by `step`.
// 
// If `start` is less than `end` a zero-length range is created unless a negative `step` is specified.
// 
// Certain values for `start` and `end` (eg. NaN/undefined/null) are currently coerced to 0;
// for forward compatibility make sure to pass in actual numbers.
func (self *ArrayUtils) NumberArrayStep(start int) []interface{}{
	array00 := self.Object.Call("numberArrayStep", start)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Create an array of numbers (positive and/or negative) progressing from `start`
// up to but not including `end` by advancing by `step`.
// 
// If `start` is less than `end` a zero-length range is created unless a negative `step` is specified.
// 
// Certain values for `start` and `end` (eg. NaN/undefined/null) are currently coerced to 0;
// for forward compatibility make sure to pass in actual numbers.
func (self *ArrayUtils) NumberArrayStep1O(start int, end int) []interface{}{
	array00 := self.Object.Call("numberArrayStep", start, end)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Create an array of numbers (positive and/or negative) progressing from `start`
// up to but not including `end` by advancing by `step`.
// 
// If `start` is less than `end` a zero-length range is created unless a negative `step` is specified.
// 
// Certain values for `start` and `end` (eg. NaN/undefined/null) are currently coerced to 0;
// for forward compatibility make sure to pass in actual numbers.
func (self *ArrayUtils) NumberArrayStep2O(start int, end int, step int) []interface{}{
	array00 := self.Object.Call("numberArrayStep", start, end, step)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Create an array of numbers (positive and/or negative) progressing from `start`
// up to but not including `end` by advancing by `step`.
// 
// If `start` is less than `end` a zero-length range is created unless a negative `step` is specified.
// 
// Certain values for `start` and `end` (eg. NaN/undefined/null) are currently coerced to 0;
// for forward compatibility make sure to pass in actual numbers.
func (self *ArrayUtils) NumberArrayStepI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("numberArrayStep", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}
