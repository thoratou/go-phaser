// Package phaser Automatic generation for Phaser.ArrayUtils
// generated file ArrayUtils.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ArrayUtils Utility functions for dealing with Arrays.
type ArrayUtils struct {
    *js.Object
}

// NewArrayUtils Utility functions for dealing with Arrays.
func NewArrayUtils() *ArrayUtils {
    return &ArrayUtils{js.Global.Get("Phaser").Get("ArrayUtils").New()}
}
// NewArrayUtilsI Utility functions for dealing with Arrays.
func NewArrayUtilsI(args ...interface{}) *ArrayUtils {
    return &ArrayUtils{js.Global.Get("Phaser").Get("ArrayUtils").New(args)}
}



// ArrayUtils Binding conversion method to ArrayUtils point 
func ToArrayUtils(jsStruct interface{}) *ArrayUtils {
    if object, ok := jsStruct.(*js.Object); ok {
		return &ArrayUtils{Object: object}
	}
	return nil
}




// GetRandomItem Fetch a random entry from the given array.
// 
// Will return null if there are no array items that fall within the specified range
// or if there is no item for the randomly chosen index.
func (self *ArrayUtils) GetRandomItem(objects []interface{}, startIndex int, length int) interface{}{
    return self.Object.Call("getRandomItem", objects, startIndex, length)
}

// GetRandomItemI Fetch a random entry from the given array.
// 
// Will return null if there are no array items that fall within the specified range
// or if there is no item for the randomly chosen index.
func (self *ArrayUtils) GetRandomItemI(args ...interface{}) interface{}{
    return self.Object.Call("getRandomItem", args)
}

// RemoveRandomItem Removes a random object from the given array and returns it.
// 
// Will return null if there are no array items that fall within the specified range
// or if there is no item for the randomly chosen index.
func (self *ArrayUtils) RemoveRandomItem(objects []interface{}, startIndex int, length int) interface{}{
    return self.Object.Call("removeRandomItem", objects, startIndex, length)
}

// RemoveRandomItemI Removes a random object from the given array and returns it.
// 
// Will return null if there are no array items that fall within the specified range
// or if there is no item for the randomly chosen index.
func (self *ArrayUtils) RemoveRandomItemI(args ...interface{}) interface{}{
    return self.Object.Call("removeRandomItem", args)
}

// Shuffle A standard Fisher-Yates Array shuffle implementation which modifies the array in place.
func (self *ArrayUtils) Shuffle(array []interface{}) []interface{}{
	array00 := self.Object.Call("shuffle", array)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// ShuffleI A standard Fisher-Yates Array shuffle implementation which modifies the array in place.
func (self *ArrayUtils) ShuffleI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("shuffle", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// TransposeMatrix Transposes the elements of the given matrix (array of arrays).
func (self *ArrayUtils) TransposeMatrix(array [][]interface{}) [][]interface{}{
	array00 := self.Object.Call("transposeMatrix", array)
	length00 := array00.Length()
	out00 := make([][]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = []interface{}{array00.Index(i00)}
	}
	return out00
}

// TransposeMatrixI Transposes the elements of the given matrix (array of arrays).
func (self *ArrayUtils) TransposeMatrixI(args ...interface{}) [][]interface{}{
	array00 := self.Object.Call("transposeMatrix", args)
	length00 := array00.Length()
	out00 := make([][]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		
			out00[i00] = []interface{}{array00.Index(i00)}
	}
	return out00
}

// RotateMatrix Rotates the given matrix (array of arrays).
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

// RotateMatrixI Rotates the given matrix (array of arrays).
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

// FindClosest Snaps a value to the nearest value in an array.
// The result will always be in the range `[first_value, last_value]`.
func (self *ArrayUtils) FindClosest(value int, arr []int) int{
    return self.Object.Call("findClosest", value, arr).Int()
}

// FindClosestI Snaps a value to the nearest value in an array.
// The result will always be in the range `[first_value, last_value]`.
func (self *ArrayUtils) FindClosestI(args ...interface{}) int{
    return self.Object.Call("findClosest", args).Int()
}

// RotateRight Moves the element from the end of the array to the start, shifting all items in the process.
// The "rotation" happens to the right.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ F, A, B, C, D, E ]`
// 
// See also Phaser.ArrayUtils.rotateLeft.
func (self *ArrayUtils) RotateRight(array []interface{}) interface{}{
    return self.Object.Call("rotateRight", array)
}

// RotateRightI Moves the element from the end of the array to the start, shifting all items in the process.
// The "rotation" happens to the right.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ F, A, B, C, D, E ]`
// 
// See also Phaser.ArrayUtils.rotateLeft.
func (self *ArrayUtils) RotateRightI(args ...interface{}) interface{}{
    return self.Object.Call("rotateRight", args)
}

// RotateLeft Moves the element from the start of the array to the end, shifting all items in the process.
// The "rotation" happens to the left.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ B, C, D, E, F, A ]`
// 
// See also Phaser.ArrayUtils.rotateRight
func (self *ArrayUtils) RotateLeft(array []interface{}) interface{}{
    return self.Object.Call("rotateLeft", array)
}

// RotateLeftI Moves the element from the start of the array to the end, shifting all items in the process.
// The "rotation" happens to the left.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ B, C, D, E, F, A ]`
// 
// See also Phaser.ArrayUtils.rotateRight
func (self *ArrayUtils) RotateLeftI(args ...interface{}) interface{}{
    return self.Object.Call("rotateLeft", args)
}

// Rotate Moves the element from the start of the array to the end, shifting all items in the process.
// The "rotation" happens to the left.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ B, C, D, E, F, A ]`
// 
// See also Phaser.ArrayUtils.rotateRight
func (self *ArrayUtils) Rotate(array []interface{}) interface{}{
    return self.Object.Call("rotate", array)
}

// RotateI Moves the element from the start of the array to the end, shifting all items in the process.
// The "rotation" happens to the left.
// 
// Before: `[ A, B, C, D, E, F ]`
// After: `[ B, C, D, E, F, A ]`
// 
// See also Phaser.ArrayUtils.rotateRight
func (self *ArrayUtils) RotateI(args ...interface{}) interface{}{
    return self.Object.Call("rotate", args)
}

// NumberArray Create an array representing the inclusive range of numbers (usually integers) in `[start, end]`.
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

// NumberArrayI Create an array representing the inclusive range of numbers (usually integers) in `[start, end]`.
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

// NumberArrayStep Create an array of numbers (positive and/or negative) progressing from `start`
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
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// NumberArrayStep1O Create an array of numbers (positive and/or negative) progressing from `start`
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
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// NumberArrayStep2O Create an array of numbers (positive and/or negative) progressing from `start`
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
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// NumberArrayStepI Create an array of numbers (positive and/or negative) progressing from `start`
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
		out00[i00] = array00.Index(i00)
	}
	return out00
}

