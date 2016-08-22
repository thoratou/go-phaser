// Automatic generation for PIXI.CanvasPool
// generated file CanvasPool.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// The CanvasPool is a global static object that allows Pixi and Phaser to pool canvas DOM elements.
type CanvasPool struct {
    *js.Object
}


// The pool into which the canvas dom elements are placed.
func (self *CanvasPool) GetPoolA() []interface{}{
	array00 := self.Object.Get("pool")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// The pool into which the canvas dom elements are placed.
func (self *CanvasPool) SetPoolA(member []interface{}) {
    self.Object.Set("pool", member)
}



// Creates a new Canvas DOM element, or pulls one from the pool if free.
func (self *CanvasPool) Create(parent interface{}, width int, height int) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("create", parent, width, height))
}

// Creates a new Canvas DOM element, or pulls one from the pool if free.
func (self *CanvasPool) CreateI(args ...interface{}) dom.HTMLCanvasElement{
    return WrapHTMLCanvasElement(self.Object.Call("create", args))
}

// Gets the first free canvas index from the pool.
func (self *CanvasPool) GetFirst() int{
    return self.Object.Call("getFirst").Int()
}

// Gets the first free canvas index from the pool.
func (self *CanvasPool) GetFirstI(args ...interface{}) int{
    return self.Object.Call("getFirst", args).Int()
}

// Removes the parent from a canvas element from the pool, freeing it up for re-use.
func (self *CanvasPool) Remove(parent interface{}) {
    self.Object.Call("remove", parent)
}

// Removes the parent from a canvas element from the pool, freeing it up for re-use.
func (self *CanvasPool) RemoveI(args ...interface{}) {
    self.Object.Call("remove", args)
}

// Removes the parent from a canvas element from the pool, freeing it up for re-use.
func (self *CanvasPool) RemoveByCanvas(canvas *dom.HTMLCanvasElement) {
    self.Object.Call("removeByCanvas", canvas)
}

// Removes the parent from a canvas element from the pool, freeing it up for re-use.
func (self *CanvasPool) RemoveByCanvasI(args ...interface{}) {
    self.Object.Call("removeByCanvas", args)
}

// Gets the total number of used canvas elements in the pool.
func (self *CanvasPool) GetTotal() int{
    return self.Object.Call("getTotal").Int()
}

// Gets the total number of used canvas elements in the pool.
func (self *CanvasPool) GetTotalI(args ...interface{}) int{
    return self.Object.Call("getTotal", args).Int()
}

// Gets the total number of free canvas elements in the pool.
func (self *CanvasPool) GetFree() int{
    return self.Object.Call("getFree").Int()
}

// Gets the total number of free canvas elements in the pool.
func (self *CanvasPool) GetFreeI(args ...interface{}) int{
    return self.Object.Call("getFree", args).Int()
}
