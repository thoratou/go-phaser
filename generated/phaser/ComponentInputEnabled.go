// Package phaser Automatic generation for Phaser.Component.InputEnabled
// generated file ComponentInputEnabled.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ComponentInputEnabled The InputEnabled component allows a Game Object to have its own InputHandler and process input related events.
type ComponentInputEnabled struct {
    *js.Object
}

// NewComponentInputEnabled The InputEnabled component allows a Game Object to have its own InputHandler and process input related events.
func NewComponentInputEnabled() *ComponentInputEnabled {
    return &ComponentInputEnabled{js.Global.Get("Phaser").Get("Component").Get("InputEnabled").New()}
}
// NewComponentInputEnabledI The InputEnabled component allows a Game Object to have its own InputHandler and process input related events.
func NewComponentInputEnabledI(args ...interface{}) *ComponentInputEnabled {
    return &ComponentInputEnabled{js.Global.Get("Phaser").Get("Component").Get("InputEnabled").New(args)}
}



// ComponentInputEnabled Binding conversion method to ComponentInputEnabled point 
func ToComponentInputEnabled(jsStruct interface{}) *ComponentInputEnabled {
    if object, ok := jsStruct.(*js.Object); ok {
		return &ComponentInputEnabled{Object: object}
	}
	return nil
}



// Input The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *ComponentInputEnabled) Input() interface{}{
    return self.Object.Get("input")
}

// SetInputA The Input Handler for this Game Object.
// 
// By default it is disabled. If you wish this Game Object to process input events you should enable it with: `inputEnabled = true`.
// 
// After you have done this, this property will be a reference to the Phaser InputHandler.
func (self *ComponentInputEnabled) SetInputA(member interface{}) {
    self.Object.Set("input", member)
}

// InputEnabled By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
// for this Game Object and it will then start to process click / touch events and more.
// 
// You can then access the Input Handler via `this.input`.
// 
// Note that Input related events are dispatched from `this.events`, i.e.: `events.onInputDown`.
// 
// If you set this property to false it will stop the Input Handler from processing any more input events.
// 
// If you want to _temporarily_ disable input for a Game Object, then it's better to set
// `input.enabled = false`, as it won't reset any of the Input Handlers internal properties.
// You can then toggle this back on as needed.
func (self *ComponentInputEnabled) InputEnabled() bool{
    return self.Object.Get("inputEnabled").Bool()
}

// SetInputEnabledA By default a Game Object won't process any input events. By setting `inputEnabled` to true a Phaser.InputHandler is created
// for this Game Object and it will then start to process click / touch events and more.
// 
// You can then access the Input Handler via `this.input`.
// 
// Note that Input related events are dispatched from `this.events`, i.e.: `events.onInputDown`.
// 
// If you set this property to false it will stop the Input Handler from processing any more input events.
// 
// If you want to _temporarily_ disable input for a Game Object, then it's better to set
// `input.enabled = false`, as it won't reset any of the Input Handlers internal properties.
// You can then toggle this back on as needed.
func (self *ComponentInputEnabled) SetInputEnabledA(member bool) {
    self.Object.Set("inputEnabled", member)
}


