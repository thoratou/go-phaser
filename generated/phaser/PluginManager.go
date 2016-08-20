// Automatic generation for Phaser.PluginManager
// generated file PluginManager.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Plugin Manager is responsible for the loading, running and unloading of Phaser Plugins.
type PluginManager struct {
    *js.Object
}


// A reference to the currently running game.
func (self *PluginManager) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *PluginManager) SetGame(member Game) {
    self.Set("game", member)
}

// An array of all the plugins being managed by this PluginManager.
func (self *PluginManager) GetPlugins() []Plugin{
	array := self.Get("plugins")
	length := array.Length()
	out := make([]Plugin, length, length)
	for i := 0; i < length; i++ {
		out[i] = Plugin{array.Index(i)}
	}
	return out
}

// An array of all the plugins being managed by this PluginManager.
func (self *PluginManager) SetPlugins(member []Plugin) {
    self.Set("plugins", member)
}



// Add a new Plugin into the PluginManager.
// The Plugin must have 2 properties: game and parent. Plugin.game is set to the game reference the PluginManager uses, and parent is set to the PluginManager.
func (self *PluginManager) AddI(args ...interface{}) Plugin{
    return Plugin{self.Call("add", args)}
}

// Remove a Plugin from the PluginManager. It calls Plugin.destroy on the plugin before removing it from the manager.
func (self *PluginManager) RemoveI(args ...interface{}) {
    self.Call("remove", args)
}

// Remove all Plugins from the PluginManager. It calls Plugin.destroy on every plugin before removing it from the manager.
func (self *PluginManager) RemoveAllI(args ...interface{}) {
    self.Call("removeAll", args)
}

// Pre-update is called at the very start of the update cycle, before any other subsystems have been updated (including Physics).
// It only calls plugins who have active=true.
func (self *PluginManager) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Update is called after all the core subsystems (Input, Tweens, Sound, etc) and the State have updated, but before the render.
// It only calls plugins who have active=true.
func (self *PluginManager) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// PostUpdate is the last thing to be called before the world render.
// In particular, it is called after the world postUpdate, which means the camera has been adjusted.
// It only calls plugins who have active=true.
func (self *PluginManager) PostUpdateI(args ...interface{}) {
    self.Call("postUpdate", args)
}

// Render is called right after the Game Renderer completes, but before the State.render.
// It only calls plugins who have visible=true.
func (self *PluginManager) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// Post-render is called after the Game Renderer and State.render have run.
// It only calls plugins who have visible=true.
func (self *PluginManager) PostRenderI(args ...interface{}) {
    self.Call("postRender", args)
}

// Clear down this PluginManager, calls destroy on every plugin and nulls out references.
func (self *PluginManager) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
