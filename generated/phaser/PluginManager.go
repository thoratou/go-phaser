// Package phaser Automatic generation for Phaser.PluginManager
// generated file PluginManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// PluginManager The Plugin Manager is responsible for the loading, running and unloading of Phaser Plugins.
type PluginManager struct {
    *js.Object
}

// NewPluginManager The Plugin Manager is responsible for the loading, running and unloading of Phaser Plugins.
func NewPluginManager(game *Game) *PluginManager {
    return &PluginManager{js.Global.Get("Phaser").Get("PluginManager").New(game)}
}
// NewPluginManagerI The Plugin Manager is responsible for the loading, running and unloading of Phaser Plugins.
func NewPluginManagerI(args ...interface{}) *PluginManager {
    return &PluginManager{js.Global.Get("Phaser").Get("PluginManager").New(args)}
}



// PluginManager Binding conversion method to PluginManager point 
func ToPluginManager(jsStruct interface{}) *PluginManager {
    if object, ok := jsStruct.(*js.Object); ok {
		return &PluginManager{Object: object}
	}
	return nil
}



// Game A reference to the currently running game.
func (self *PluginManager) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *PluginManager) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Plugins An array of all the plugins being managed by this PluginManager.
func (self *PluginManager) Plugins() []Plugin{
	array00 := self.Object.Get("plugins")
	length00 := array00.Length()
	out00 := make([]Plugin, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = Plugin{array00.Index(i00)}
	}
	return out00
}

// SetPluginsA An array of all the plugins being managed by this PluginManager.
func (self *PluginManager) SetPluginsA(member []Plugin) {
    self.Object.Set("plugins", member)
}


// Add Add a new Plugin into the PluginManager.
// The Plugin must have 2 properties: game and parent. Plugin.game is set to the game reference the PluginManager uses, and parent is set to the PluginManager.
func (self *PluginManager) Add(plugin interface{}, parameter interface{}) *Plugin{
    return &Plugin{self.Object.Call("add", plugin, parameter)}
}

// AddI Add a new Plugin into the PluginManager.
// The Plugin must have 2 properties: game and parent. Plugin.game is set to the game reference the PluginManager uses, and parent is set to the PluginManager.
func (self *PluginManager) AddI(args ...interface{}) *Plugin{
    return &Plugin{self.Object.Call("add", args)}
}

// Remove Remove a Plugin from the PluginManager. It calls Plugin.destroy on the plugin before removing it from the manager.
func (self *PluginManager) Remove(plugin *Plugin) {
    self.Object.Call("remove", plugin)
}

// Remove1O Remove a Plugin from the PluginManager. It calls Plugin.destroy on the plugin before removing it from the manager.
func (self *PluginManager) Remove1O(plugin *Plugin, destroy bool) {
    self.Object.Call("remove", plugin, destroy)
}

// RemoveI Remove a Plugin from the PluginManager. It calls Plugin.destroy on the plugin before removing it from the manager.
func (self *PluginManager) RemoveI(args ...interface{}) {
    self.Object.Call("remove", args)
}

// RemoveAll Remove all Plugins from the PluginManager. It calls Plugin.destroy on every plugin before removing it from the manager.
func (self *PluginManager) RemoveAll() {
    self.Object.Call("removeAll")
}

// RemoveAllI Remove all Plugins from the PluginManager. It calls Plugin.destroy on every plugin before removing it from the manager.
func (self *PluginManager) RemoveAllI(args ...interface{}) {
    self.Object.Call("removeAll", args)
}

// PreUpdate Pre-update is called at the very start of the update cycle, before any other subsystems have been updated (including Physics).
// It only calls plugins who have active=true.
func (self *PluginManager) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI Pre-update is called at the very start of the update cycle, before any other subsystems have been updated (including Physics).
// It only calls plugins who have active=true.
func (self *PluginManager) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Update Update is called after all the core subsystems (Input, Tweens, Sound, etc) and the State have updated, but before the render.
// It only calls plugins who have active=true.
func (self *PluginManager) Update() {
    self.Object.Call("update")
}

// UpdateI Update is called after all the core subsystems (Input, Tweens, Sound, etc) and the State have updated, but before the render.
// It only calls plugins who have active=true.
func (self *PluginManager) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// PostUpdate PostUpdate is the last thing to be called before the world render.
// In particular, it is called after the world postUpdate, which means the camera has been adjusted.
// It only calls plugins who have active=true.
func (self *PluginManager) PostUpdate() {
    self.Object.Call("postUpdate")
}

// PostUpdateI PostUpdate is the last thing to be called before the world render.
// In particular, it is called after the world postUpdate, which means the camera has been adjusted.
// It only calls plugins who have active=true.
func (self *PluginManager) PostUpdateI(args ...interface{}) {
    self.Object.Call("postUpdate", args)
}

// Render Render is called right after the Game Renderer completes, but before the State.render.
// It only calls plugins who have visible=true.
func (self *PluginManager) Render() {
    self.Object.Call("render")
}

// RenderI Render is called right after the Game Renderer completes, but before the State.render.
// It only calls plugins who have visible=true.
func (self *PluginManager) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// PostRender Post-render is called after the Game Renderer and State.render have run.
// It only calls plugins who have visible=true.
func (self *PluginManager) PostRender() {
    self.Object.Call("postRender")
}

// PostRenderI Post-render is called after the Game Renderer and State.render have run.
// It only calls plugins who have visible=true.
func (self *PluginManager) PostRenderI(args ...interface{}) {
    self.Object.Call("postRender", args)
}

// Destroy Clear down this PluginManager, calls destroy on every plugin and nulls out references.
func (self *PluginManager) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Clear down this PluginManager, calls destroy on every plugin and nulls out references.
func (self *PluginManager) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

