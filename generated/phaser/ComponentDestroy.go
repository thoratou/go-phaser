// Automatic generation for Phaser.Component.Destroy
// generated file ComponentDestroy.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Destroy component is responsible for destroying a Game Object.
type ComponentDestroy struct {
    *js.Object
}


// The Destroy component is responsible for destroying a Game Object.
func NewComponentDestroy() *ComponentDestroy {
    return &ComponentDestroy{js.Global.Get("Phaser").Get("Component").Get("Destroy").New()}
}

// The Destroy component is responsible for destroying a Game Object.
func NewComponentDestroyI(args ...interface{}) *ComponentDestroy {
    return &ComponentDestroy{js.Global.Get("Phaser").Get("Component").Get("Destroy").New(args)}
}



// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *ComponentDestroy) DestroyPhase() bool{
    return self.Object.Get("destroyPhase").Bool()
}

// As a Game Object runs through its destroy method this flag is set to true, 
// and can be checked in any sub-systems or plugins it is being destroyed from.
func (self *ComponentDestroy) SetDestroyPhaseA(member bool) {
    self.Object.Set("destroyPhase", member)
}



// Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *ComponentDestroy) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *ComponentDestroy) Destroy1O(destroyChildren bool) {
    self.Object.Call("destroy", destroyChildren)
}

// Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *ComponentDestroy) Destroy2O(destroyChildren bool, destroyTexture bool) {
    self.Object.Call("destroy", destroyChildren, destroyTexture)
}

// Destroys the Game Object. This removes it from its parent group, destroys the input, event and animation handlers if present
// and nulls its reference to `game`, freeing it up for garbage collection.
// 
// If this Game Object has the Events component it will also dispatch the `onDestroy` event.
// 
// You can optionally also destroy the BaseTexture this Game Object is using. Be careful if you've
// more than one Game Object sharing the same BaseTexture.
func (self *ComponentDestroy) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
