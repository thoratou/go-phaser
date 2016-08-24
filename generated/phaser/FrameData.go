// Package phaser Automatic generation for Phaser.FrameData
// generated file FrameData.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// FrameData FrameData is a container for Frame objects, which are the internal representation of animation data in Phaser.
type FrameData struct {
    *js.Object
}

// NewFrameData FrameData is a container for Frame objects, which are the internal representation of animation data in Phaser.
func NewFrameData() *FrameData {
    return &FrameData{js.Global.Get("Phaser").Get("FrameData").New()}
}
// NewFrameDataI FrameData is a container for Frame objects, which are the internal representation of animation data in Phaser.
func NewFrameDataI(args ...interface{}) *FrameData {
    return &FrameData{js.Global.Get("Phaser").Get("FrameData").New(args)}
}



// Total The total number of frames in this FrameData set.
func (self *FrameData) Total() int{
    return self.Object.Get("total").Int()
}

// SetTotalA The total number of frames in this FrameData set.
func (self *FrameData) SetTotalA(member int) {
    self.Object.Set("total", member)
}


// AddFrame Adds a new Frame to this FrameData collection. Typically called by the Animation.Parser and not directly.
func (self *FrameData) AddFrame(frame *Frame) *Frame{
    return &Frame{self.Object.Call("addFrame", frame)}
}

// AddFrameI Adds a new Frame to this FrameData collection. Typically called by the Animation.Parser and not directly.
func (self *FrameData) AddFrameI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("addFrame", args)}
}

// GetFrame Get a Frame by its numerical index.
func (self *FrameData) GetFrame(index int) *Frame{
    return &Frame{self.Object.Call("getFrame", index)}
}

// GetFrameI Get a Frame by its numerical index.
func (self *FrameData) GetFrameI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("getFrame", args)}
}

// GetFrameByName Get a Frame by its frame name.
func (self *FrameData) GetFrameByName(name string) *Frame{
    return &Frame{self.Object.Call("getFrameByName", name)}
}

// GetFrameByNameI Get a Frame by its frame name.
func (self *FrameData) GetFrameByNameI(args ...interface{}) *Frame{
    return &Frame{self.Object.Call("getFrameByName", args)}
}

// CheckFrameName Check if there is a Frame with the given name.
func (self *FrameData) CheckFrameName(name string) bool{
    return self.Object.Call("checkFrameName", name).Bool()
}

// CheckFrameNameI Check if there is a Frame with the given name.
func (self *FrameData) CheckFrameNameI(args ...interface{}) bool{
    return self.Object.Call("checkFrameName", args).Bool()
}

// Clone Makes a copy of this FrameData including copies (not references) to all of the Frames it contains.
func (self *FrameData) Clone() *FrameData{
    return &FrameData{self.Object.Call("clone")}
}

// CloneI Makes a copy of this FrameData including copies (not references) to all of the Frames it contains.
func (self *FrameData) CloneI(args ...interface{}) *FrameData{
    return &FrameData{self.Object.Call("clone", args)}
}

// GetFrameRange Returns a range of frames based on the given start and end frame indexes and returns them in an Array.
func (self *FrameData) GetFrameRange(start int, end int) []interface{}{
	array00 := self.Object.Call("getFrameRange", start, end)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrameRange1O Returns a range of frames based on the given start and end frame indexes and returns them in an Array.
func (self *FrameData) GetFrameRange1O(start int, end int, output []interface{}) []interface{}{
	array00 := self.Object.Call("getFrameRange", start, end, output)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrameRangeI Returns a range of frames based on the given start and end frame indexes and returns them in an Array.
func (self *FrameData) GetFrameRangeI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("getFrameRange", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrames Returns all of the Frames in this FrameData set where the frame index is found in the input array.
// The frames are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFrames() []interface{}{
	array00 := self.Object.Call("getFrames")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrames1O Returns all of the Frames in this FrameData set where the frame index is found in the input array.
// The frames are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFrames1O(frames []interface{}) []interface{}{
	array00 := self.Object.Call("getFrames", frames)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrames2O Returns all of the Frames in this FrameData set where the frame index is found in the input array.
// The frames are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFrames2O(frames []interface{}, useNumericIndex bool) []interface{}{
	array00 := self.Object.Call("getFrames", frames, useNumericIndex)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrames3O Returns all of the Frames in this FrameData set where the frame index is found in the input array.
// The frames are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFrames3O(frames []interface{}, useNumericIndex bool, output []interface{}) []interface{}{
	array00 := self.Object.Call("getFrames", frames, useNumericIndex, output)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFramesI Returns all of the Frames in this FrameData set where the frame index is found in the input array.
// The frames are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFramesI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("getFrames", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrameIndexes Returns all of the Frame indexes in this FrameData set.
// The frames indexes are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFrameIndexes() []interface{}{
	array00 := self.Object.Call("getFrameIndexes")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrameIndexes1O Returns all of the Frame indexes in this FrameData set.
// The frames indexes are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFrameIndexes1O(frames []interface{}) []interface{}{
	array00 := self.Object.Call("getFrameIndexes", frames)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrameIndexes2O Returns all of the Frame indexes in this FrameData set.
// The frames indexes are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFrameIndexes2O(frames []interface{}, useNumericIndex bool) []interface{}{
	array00 := self.Object.Call("getFrameIndexes", frames, useNumericIndex)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrameIndexes3O Returns all of the Frame indexes in this FrameData set.
// The frames indexes are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFrameIndexes3O(frames []interface{}, useNumericIndex bool, output []interface{}) []interface{}{
	array00 := self.Object.Call("getFrameIndexes", frames, useNumericIndex, output)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// GetFrameIndexesI Returns all of the Frame indexes in this FrameData set.
// The frames indexes are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFrameIndexesI(args ...interface{}) []interface{}{
	array00 := self.Object.Call("getFrameIndexes", args)
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// Destroy Destroys this FrameData collection by nulling the _frames and _frameNames arrays.
func (self *FrameData) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Destroys this FrameData collection by nulling the _frames and _frameNames arrays.
func (self *FrameData) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

