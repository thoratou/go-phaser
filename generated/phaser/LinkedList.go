// Automatic generation for Phaser.LinkedList
// generated file LinkedList.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// A basic Linked List data structure.
// 
// This implementation _modifies_ the `prev` and `next` properties of each item added:
// - The `prev` and `next` properties must be writable and should not be used for any other purpose.
// - Items _cannot_ be added to multiple LinkedLists at the same time.
// - Only objects can be added.
type LinkedList struct {
    *js.Object
}


// Next element in the list.
func (self *LinkedList) GetNext() interface{}{
    return self.Get("next")
}

// Next element in the list.
func (self *LinkedList) SetNext(member interface{}) {
    self.Set("next", member)
}

// Previous element in the list.
func (self *LinkedList) GetPrev() interface{}{
    return self.Get("prev")
}

// Previous element in the list.
func (self *LinkedList) SetPrev(member interface{}) {
    self.Set("prev", member)
}

// First element in the list.
func (self *LinkedList) GetFirst() interface{}{
    return self.Get("first")
}

// First element in the list.
func (self *LinkedList) SetFirst(member interface{}) {
    self.Set("first", member)
}

// Last element in the list.
func (self *LinkedList) GetLast() interface{}{
    return self.Get("last")
}

// Last element in the list.
func (self *LinkedList) SetLast(member interface{}) {
    self.Set("last", member)
}

// Number of elements in the list.
func (self *LinkedList) GetTotal() int{
    return self.Get("total").Int()
}

// Number of elements in the list.
func (self *LinkedList) SetTotal(member int) {
    self.Set("total", member)
}



// Adds a new element to this linked list.
func (self *LinkedList) AddI(args ...interface{}) interface{}{
    return self.Call("add", args)
}

// Resets the first, last, next and previous node pointers in this list.
func (self *LinkedList) ResetI(args ...interface{}) {
    self.Call("reset", args)
}

// Removes the given element from this linked list if it exists.
func (self *LinkedList) RemoveI(args ...interface{}) {
    self.Call("remove", args)
}

// Calls a function on all members of this list, using the member as the context for the callback.
// The function must exist on the member.
func (self *LinkedList) CallAllI(args ...interface{}) {
    self.Call("callAll", args)
}
