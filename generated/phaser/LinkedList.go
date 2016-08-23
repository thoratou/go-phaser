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


// A basic Linked List data structure.
// 
// This implementation _modifies_ the `prev` and `next` properties of each item added:
// - The `prev` and `next` properties must be writable and should not be used for any other purpose.
// - Items _cannot_ be added to multiple LinkedLists at the same time.
// - Only objects can be added.
func NewLinkedList() *LinkedList {
    return &LinkedList{js.Global.Get("Phaser").Get("LinkedList").New()}
}

// A basic Linked List data structure.
// 
// This implementation _modifies_ the `prev` and `next` properties of each item added:
// - The `prev` and `next` properties must be writable and should not be used for any other purpose.
// - Items _cannot_ be added to multiple LinkedLists at the same time.
// - Only objects can be added.
func NewLinkedListI(args ...interface{}) *LinkedList {
    return &LinkedList{js.Global.Get("Phaser").Get("LinkedList").New(args)}
}



// Next element in the list.
func (self *LinkedList) Next() interface{}{
    return self.Object.Get("next")
}

// Next element in the list.
func (self *LinkedList) SetNextA(member interface{}) {
    self.Object.Set("next", member)
}

// Previous element in the list.
func (self *LinkedList) Prev() interface{}{
    return self.Object.Get("prev")
}

// Previous element in the list.
func (self *LinkedList) SetPrevA(member interface{}) {
    self.Object.Set("prev", member)
}

// First element in the list.
func (self *LinkedList) First() interface{}{
    return self.Object.Get("first")
}

// First element in the list.
func (self *LinkedList) SetFirstA(member interface{}) {
    self.Object.Set("first", member)
}

// Last element in the list.
func (self *LinkedList) Last() interface{}{
    return self.Object.Get("last")
}

// Last element in the list.
func (self *LinkedList) SetLastA(member interface{}) {
    self.Object.Set("last", member)
}

// Number of elements in the list.
func (self *LinkedList) Total() int{
    return self.Object.Get("total").Int()
}

// Number of elements in the list.
func (self *LinkedList) SetTotalA(member int) {
    self.Object.Set("total", member)
}



// Adds a new element to this linked list.
func (self *LinkedList) Add(item interface{}) interface{}{
    return self.Object.Call("add", item)
}

// Adds a new element to this linked list.
func (self *LinkedList) AddI(args ...interface{}) interface{}{
    return self.Object.Call("add", args)
}

// Resets the first, last, next and previous node pointers in this list.
func (self *LinkedList) Reset() {
    self.Object.Call("reset")
}

// Resets the first, last, next and previous node pointers in this list.
func (self *LinkedList) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}

// Removes the given element from this linked list if it exists.
func (self *LinkedList) Remove(item interface{}) {
    self.Object.Call("remove", item)
}

// Removes the given element from this linked list if it exists.
func (self *LinkedList) RemoveI(args ...interface{}) {
    self.Object.Call("remove", args)
}

// Calls a function on all members of this list, using the member as the context for the callback.
// The function must exist on the member.
func (self *LinkedList) CallAll(callback interface{}) {
    self.Object.Call("callAll", callback)
}

// Calls a function on all members of this list, using the member as the context for the callback.
// The function must exist on the member.
func (self *LinkedList) CallAllI(args ...interface{}) {
    self.Object.Call("callAll", args)
}
