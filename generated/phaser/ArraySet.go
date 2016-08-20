// Automatic generation for Phaser.ArraySet
// generated file ArraySet.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ArraySet is a Set data structure (items must be unique within the set) that also maintains order.
// This allows specific items to be easily added or removed from the Set.
// 
// Item equality (and uniqueness) is determined by the behavior of `Array.indexOf`.
// 
// This used primarily by the Input subsystem.
type ArraySet struct {
    *js.Object
}


// Current cursor position as established by `first` and `next`.
func (self *ArraySet) GetPosition() int{
    return self.Get("position").Int()
}

// Current cursor position as established by `first` and `next`.
func (self *ArraySet) SetPosition(member int) {
    self.Set("position", member)
}

// The backing array.
func (self *ArraySet) GetList() []interface{}{
	array := self.Get("list")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The backing array.
func (self *ArraySet) SetList(member []interface{}) {
    self.Set("list", member)
}

// Number of items in the ArraySet. Same as `list.length`.
func (self *ArraySet) GetTotal() int{
    return self.Get("total").Int()
}

// Number of items in the ArraySet. Same as `list.length`.
func (self *ArraySet) SetTotal(member int) {
    self.Set("total", member)
}

// Returns the first item and resets the cursor to the start.
func (self *ArraySet) GetFirst() interface{}{
    return self.Get("first")
}

// Returns the first item and resets the cursor to the start.
func (self *ArraySet) SetFirst(member interface{}) {
    self.Set("first", member)
}

// Returns the the next item (based on the cursor) and advances the cursor.
func (self *ArraySet) GetNext() interface{}{
    return self.Get("next")
}

// Returns the the next item (based on the cursor) and advances the cursor.
func (self *ArraySet) SetNext(member interface{}) {
    self.Set("next", member)
}



// Adds a new element to the end of the list.
// If the item already exists in the list it is not moved.
func (self *ArraySet) AddI(args ...interface{}) interface{}{
    return self.Call("add", args)
}

// Gets the index of the item in the list, or -1 if it isn't in the list.
func (self *ArraySet) GetIndexI(args ...interface{}) int{
    return self.Call("getIndex", args).Int()
}

// Gets an item from the set based on the property strictly equaling the value given.
// Returns null if not found.
func (self *ArraySet) GetByKeyI(args ...interface{}) interface{}{
    return self.Call("getByKey", args)
}

// Checks for the item within this list.
func (self *ArraySet) ExistsI(args ...interface{}) bool{
    return self.Call("exists", args).Bool()
}

// Removes all the items.
func (self *ArraySet) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Removes the given element from this list if it exists.
func (self *ArraySet) RemoveI(args ...interface{}) interface{}{
    return self.Call("remove", args)
}

// Sets the property `key` to the given value on all members of this list.
func (self *ArraySet) SetAllI(args ...interface{}) {
    self.Call("setAll", args)
}

// Calls a function on all members of this list, using the member as the context for the callback.
// 
// If the `key` property is present it must be a function.
// The function is invoked using the item as the context.
func (self *ArraySet) CallAllI(args ...interface{}) {
    self.Call("callAll", args)
}

// Removes every member from this ArraySet and optionally destroys it.
func (self *ArraySet) RemoveAllI(args ...interface{}) {
    self.Call("removeAll", args)
}
