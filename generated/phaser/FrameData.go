// Automatic generation for Phaser.FrameData
// generated file FrameData.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// FrameData is a container for Frame objects, which are the internal representation of animation data in Phaser.
type FrameData struct {
    *js.Object
}


// The total number of frames in this FrameData set.
func (self *FrameData) GetTotal() float64{
    return self.Get("total").Float()
}

// The total number of frames in this FrameData set.
func (self *FrameData) SetTotal(member float64) {
    self.Set("total", member)
}



// Adds a new Frame to this FrameData collection. Typically called by the Animation.Parser and not directly.
func (self *FrameData) AddFrameI(args ...interface{}) *Frame{
    return &Frame{self.Call("addFrame", args)}
}

// Get a Frame by its numerical index.
func (self *FrameData) GetFrameI(args ...interface{}) *Frame{
    return &Frame{self.Call("getFrame", args)}
}

// Get a Frame by its frame name.
func (self *FrameData) GetFrameByNameI(args ...interface{}) *Frame{
    return &Frame{self.Call("getFrameByName", args)}
}

// Check if there is a Frame with the given name.
func (self *FrameData) CheckFrameNameI(args ...interface{}) bool{
    return self.Call("checkFrameName", args).Bool()
}

// Makes a copy of this FrameData including copies (not references) to all of the Frames it contains.
func (self *FrameData) CloneI(args ...interface{}) *FrameData{
    return &FrameData{self.Call("clone", args)}
}

// Returns a range of frames based on the given start and end frame indexes and returns them in an Array.
func (self *FrameData) GetFrameRangeI(args ...interface{}) []interface{}{
	array := self.Call("getFrameRange", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Returns all of the Frames in this FrameData set where the frame index is found in the input array.
// The frames are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFramesI(args ...interface{}) []interface{}{
	array := self.Call("getFrames", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Returns all of the Frame indexes in this FrameData set.
// The frames indexes are returned in the output array, or if none is provided in a new Array object.
func (self *FrameData) GetFrameIndexesI(args ...interface{}) []interface{}{
	array := self.Call("getFrameIndexes", args)
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Destroys this FrameData collection by nulling the _frames and _frameNames arrays.
func (self *FrameData) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
