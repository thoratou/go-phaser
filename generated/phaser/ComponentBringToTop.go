// Package phaser Automatic generation for Phaser.Component.BringToTop
// generated file ComponentBringToTop.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ComponentBringToTop The BringToTop Component features quick access to Group sorting related methods.
type ComponentBringToTop struct {
    *js.Object
}

// NewComponentBringToTop The BringToTop Component features quick access to Group sorting related methods.
func NewComponentBringToTop() *ComponentBringToTop {
    return &ComponentBringToTop{js.Global.Get("Phaser").Get("Component").Get("BringToTop").New()}
}
// NewComponentBringToTopI The BringToTop Component features quick access to Group sorting related methods.
func NewComponentBringToTopI(args ...interface{}) *ComponentBringToTop {
    return &ComponentBringToTop{js.Global.Get("Phaser").Get("Component").Get("BringToTop").New(args)}
}



// ComponentBringToTop Binding conversion method to ComponentBringToTop point 
func ToComponentBringToTop(jsStruct interface{}) *ComponentBringToTop {
    if object, ok := jsStruct.(*js.Object); ok {
		return &ComponentBringToTop{Object: object}
	}
	return nil
}




// BringToTop Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *ComponentBringToTop) BringToTop() *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop")}
}

// BringToTopI Brings this Game Object to the top of its parents display list.
// Visually this means it will render over the top of any old child in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will bring it to the top of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *ComponentBringToTop) BringToTopI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("bringToTop", args)}
}

// SendToBack Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *ComponentBringToTop) SendToBack() *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack")}
}

// SendToBackI Sends this Game Object to the bottom of its parents display list.
// Visually this means it will render below all other children in the same Group.
// 
// If this Game Object hasn't been added to a custom Group then this method will send it to the bottom of the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *ComponentBringToTop) SendToBackI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("sendToBack", args)}
}

// MoveUp Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *ComponentBringToTop) MoveUp() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp")}
}

// MoveUpI Moves this Game Object up one place in its parents display list.
// This call has no effect if the Game Object is already at the top of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object up within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *ComponentBringToTop) MoveUpI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveUp", args)}
}

// MoveDown Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *ComponentBringToTop) MoveDown() *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown")}
}

// MoveDownI Moves this Game Object down one place in its parents display list.
// This call has no effect if the Game Object is already at the bottom of the display list.
// 
// If this Game Object hasn't been added to a custom Group then this method will move it one object down within the Game World, 
// because the World is the root Group from which all Game Objects descend.
func (self *ComponentBringToTop) MoveDownI(args ...interface{}) *DisplayObject{
    return &DisplayObject{self.Object.Call("moveDown", args)}
}

