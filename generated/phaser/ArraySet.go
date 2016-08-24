// Package phaser Automatic generation for Phaser.ArraySet
// generated file ArraySet.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ArraySet ArraySet is a Set data structure (items must be unique within the set) that also maintains order.
// This allows specific items to be easily added or removed from the Set.
// 
// Item equality (and uniqueness) is determined by the behavior of `Array.indexOf`.
// 
// This used primarily by the Input subsystem.
type ArraySet struct {
    *js.Object
}

// NewArraySet ArraySet is a Set data structure (items must be unique within the set) that also maintains order.
// This allows specific items to be easily added or removed from the Set.
// 
// Item equality (and uniqueness) is determined by the behavior of `Array.indexOf`.
// 
// This used primarily by the Input subsystem.
func NewArraySet() *ArraySet {
    return &ArraySet{js.Global.Get("Phaser").Get("ArraySet").New()}
}
// NewArraySet1O ArraySet is a Set data structure (items must be unique within the set) that also maintains order.
// This allows specific items to be easily added or removed from the Set.
// 
// Item equality (and uniqueness) is determined by the behavior of `Array.indexOf`.
// 
// This used primarily by the Input subsystem.
func NewArraySet1O(list []interface{}) *ArraySet {
    return &ArraySet{js.Global.Get("Phaser").Get("ArraySet").New(list)}
}
// NewArraySetI ArraySet is a Set data structure (items must be unique within the set) that also maintains order.
// This allows specific items to be easily added or removed from the Set.
// 
// Item equality (and uniqueness) is determined by the behavior of `Array.indexOf`.
// 
// This used primarily by the Input subsystem.
func NewArraySetI(args ...interface{}) *ArraySet {
    return &ArraySet{js.Global.Get("Phaser").Get("ArraySet").New(args)}
}



// Position Current cursor position as established by `first` and `next`.
func (self *ArraySet) Position() int{
    return self.Object.Get("position").Int()
}

// SetPositionA Current cursor position as established by `first` and `next`.
func (self *ArraySet) SetPositionA(member int) {
    self.Object.Set("position", member)
}

// List The backing array.
func (self *ArraySet) List() []interface{}{
	array00 := self.Object.Get("list")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetListA The backing array.
func (self *ArraySet) SetListA(member []interface{}) {
    self.Object.Set("list", member)
}

// Total Number of items in the ArraySet. Same as `list.length`.
func (self *ArraySet) Total() int{
    return self.Object.Get("total").Int()
}

// SetTotalA Number of items in the ArraySet. Same as `list.length`.
func (self *ArraySet) SetTotalA(member int) {
    self.Object.Set("total", member)
}

// First Returns the first item and resets the cursor to the start.
func (self *ArraySet) First() interface{}{
    return self.Object.Get("first")
}

// SetFirstA Returns the first item and resets the cursor to the start.
func (self *ArraySet) SetFirstA(member interface{}) {
    self.Object.Set("first", member)
}

// Next Returns the the next item (based on the cursor) and advances the cursor.
func (self *ArraySet) Next() interface{}{
    return self.Object.Get("next")
}

// SetNextA Returns the the next item (based on the cursor) and advances the cursor.
func (self *ArraySet) SetNextA(member interface{}) {
    self.Object.Set("next", member)
}


// Add Adds a new element to the end of the list.
// If the item already exists in the list it is not moved.
func (self *ArraySet) Add(item interface{}) interface{}{
    return self.Object.Call("add", item)
}

// AddI Adds a new element to the end of the list.
// If the item already exists in the list it is not moved.
func (self *ArraySet) AddI(args ...interface{}) interface{}{
    return self.Object.Call("add", args)
}

// GetIndex Gets the index of the item in the list, or -1 if it isn't in the list.
func (self *ArraySet) GetIndex(item interface{}) int{
    return self.Object.Call("getIndex", item).Int()
}

// GetIndexI Gets the index of the item in the list, or -1 if it isn't in the list.
func (self *ArraySet) GetIndexI(args ...interface{}) int{
    return self.Object.Call("getIndex", args).Int()
}

// GetByKey Gets an item from the set based on the property strictly equaling the value given.
// Returns null if not found.
func (self *ArraySet) GetByKey(property string, value interface{}) interface{}{
    return self.Object.Call("getByKey", property, value)
}

// GetByKeyI Gets an item from the set based on the property strictly equaling the value given.
// Returns null if not found.
func (self *ArraySet) GetByKeyI(args ...interface{}) interface{}{
    return self.Object.Call("getByKey", args)
}

// Exists Checks for the item within this list.
func (self *ArraySet) Exists(item interface{}) bool{
    return self.Object.Call("exists", item).Bool()
}

// ExistsI Checks for the item within this list.
func (self *ArraySet) ExistsI(args ...interface{}) bool{
    return self.Object.Call("exists", args).Bool()
}

// Reset Removes all the items.
func (self *ArraySet) Reset() {
    self.Object.Call("reset")
}

// ResetI Removes all the items.
func (self *ArraySet) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Remove Removes the given element from this list if it exists.
func (self *ArraySet) Remove(item interface{}) interface{}{
    return self.Object.Call("remove", item)
}

// RemoveI Removes the given element from this list if it exists.
func (self *ArraySet) RemoveI(args ...interface{}) interface{}{
    return self.Object.Call("remove", args)
}

// SetAll Sets the property `key` to the given value on all members of this list.
func (self *ArraySet) SetAll(key interface{}, value interface{}) {
    self.Object.Call("setAll", key, value)
}

// SetAllI Sets the property `key` to the given value on all members of this list.
func (self *ArraySet) SetAllI(args ...interface{}) {
    self.Object.Call("setAll", args)
}

// CallAll Calls a function on all members of this list, using the member as the context for the callback.
// 
// If the `key` property is present it must be a function.
// The function is invoked using the item as the context.
func (self *ArraySet) CallAll(key string, parameter interface{}) {
    self.Object.Call("callAll", key, parameter)
}

// CallAllI Calls a function on all members of this list, using the member as the context for the callback.
// 
// If the `key` property is present it must be a function.
// The function is invoked using the item as the context.
func (self *ArraySet) CallAllI(args ...interface{}) {
    self.Object.Call("callAll", args)
}

// RemoveAll Removes every member from this ArraySet and optionally destroys it.
func (self *ArraySet) RemoveAll() {
    self.Object.Call("removeAll")
}

// RemoveAll1O Removes every member from this ArraySet and optionally destroys it.
func (self *ArraySet) RemoveAll1O(destroy bool) {
    self.Object.Call("removeAll", destroy)
}

// RemoveAllI Removes every member from this ArraySet and optionally destroys it.
func (self *ArraySet) RemoveAllI(args ...interface{}) {
    self.Object.Call("removeAll", args)
}

