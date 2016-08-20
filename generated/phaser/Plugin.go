// Automatic generation for Phaser.Plugin
// generated file Plugin.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// This is a base Plugin template to use for any Phaser plugin development.
type Plugin struct {
    *js.Object
}


// A reference to the currently running game.
func (self *Plugin) GetGame() Game{
    return Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *Plugin) SetGame(member Game) {
    self.Set("game", member)
}

// The parent of this plugin. If added to the PluginManager the parent will be set to that, otherwise it will be null.
func (self *Plugin) GetParent() interface{}{
    return self.Get("parent")
}

// The parent of this plugin. If added to the PluginManager the parent will be set to that, otherwise it will be null.
func (self *Plugin) SetParent(member interface{}) {
    self.Set("parent", member)
}

// A Plugin with active=true has its preUpdate and update methods called by the parent, otherwise they are skipped.
func (self *Plugin) GetActive() bool{
    return self.Get("active").Bool()
}

// A Plugin with active=true has its preUpdate and update methods called by the parent, otherwise they are skipped.
func (self *Plugin) SetActive(member bool) {
    self.Set("active", member)
}

// A Plugin with visible=true has its render and postRender methods called by the parent, otherwise they are skipped.
func (self *Plugin) GetVisible() bool{
    return self.Get("visible").Bool()
}

// A Plugin with visible=true has its render and postRender methods called by the parent, otherwise they are skipped.
func (self *Plugin) SetVisible(member bool) {
    self.Set("visible", member)
}

// A flag to indicate if this plugin has a preUpdate method.
func (self *Plugin) GetHasPreUpdate() bool{
    return self.Get("hasPreUpdate").Bool()
}

// A flag to indicate if this plugin has a preUpdate method.
func (self *Plugin) SetHasPreUpdate(member bool) {
    self.Set("hasPreUpdate", member)
}

// A flag to indicate if this plugin has an update method.
func (self *Plugin) GetHasUpdate() bool{
    return self.Get("hasUpdate").Bool()
}

// A flag to indicate if this plugin has an update method.
func (self *Plugin) SetHasUpdate(member bool) {
    self.Set("hasUpdate", member)
}

// A flag to indicate if this plugin has a postUpdate method.
func (self *Plugin) GetHasPostUpdate() bool{
    return self.Get("hasPostUpdate").Bool()
}

// A flag to indicate if this plugin has a postUpdate method.
func (self *Plugin) SetHasPostUpdate(member bool) {
    self.Set("hasPostUpdate", member)
}

// A flag to indicate if this plugin has a render method.
func (self *Plugin) GetHasRender() bool{
    return self.Get("hasRender").Bool()
}

// A flag to indicate if this plugin has a render method.
func (self *Plugin) SetHasRender(member bool) {
    self.Set("hasRender", member)
}

// A flag to indicate if this plugin has a postRender method.
func (self *Plugin) GetHasPostRender() bool{
    return self.Get("hasPostRender").Bool()
}

// A flag to indicate if this plugin has a postRender method.
func (self *Plugin) SetHasPostRender(member bool) {
    self.Set("hasPostRender", member)
}



// Pre-update is called at the very start of the update cycle, before any other subsystems have been updated (including Physics).
// It is only called if active is set to true.
func (self *Plugin) PreUpdateI(args ...interface{}) {
    self.Call("preUpdate", args)
}

// Update is called after all the core subsystems (Input, Tweens, Sound, etc) and the State have updated, but before the render.
// It is only called if active is set to true.
func (self *Plugin) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Render is called right after the Game Renderer completes, but before the State.render.
// It is only called if visible is set to true.
func (self *Plugin) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// Post-render is called after the Game Renderer and State.render have run.
// It is only called if visible is set to true.
func (self *Plugin) PostRenderI(args ...interface{}) {
    self.Call("postRender", args)
}

// Clear down this Plugin and null out references
func (self *Plugin) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
