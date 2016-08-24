// Package phaser Automatic generation for Phaser.Plugin
// generated file Plugin.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Plugin This is a base Plugin template to use for any Phaser plugin development.
type Plugin struct {
    *js.Object
}

// NewPlugin This is a base Plugin template to use for any Phaser plugin development.
func NewPlugin(game *Game, parent interface{}) *Plugin {
    return &Plugin{js.Global.Get("Phaser").Get("Plugin").New(game, parent)}
}
// NewPluginI This is a base Plugin template to use for any Phaser plugin development.
func NewPluginI(args ...interface{}) *Plugin {
    return &Plugin{js.Global.Get("Phaser").Get("Plugin").New(args)}
}



// Game A reference to the currently running game.
func (self *Plugin) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// SetGameA A reference to the currently running game.
func (self *Plugin) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// Parent The parent of this plugin. If added to the PluginManager the parent will be set to that, otherwise it will be null.
func (self *Plugin) Parent() interface{}{
    return self.Object.Get("parent")
}

// SetParentA The parent of this plugin. If added to the PluginManager the parent will be set to that, otherwise it will be null.
func (self *Plugin) SetParentA(member interface{}) {
    self.Object.Set("parent", member)
}

// Active A Plugin with active=true has its preUpdate and update methods called by the parent, otherwise they are skipped.
func (self *Plugin) Active() bool{
    return self.Object.Get("active").Bool()
}

// SetActiveA A Plugin with active=true has its preUpdate and update methods called by the parent, otherwise they are skipped.
func (self *Plugin) SetActiveA(member bool) {
    self.Object.Set("active", member)
}

// Visible A Plugin with visible=true has its render and postRender methods called by the parent, otherwise they are skipped.
func (self *Plugin) Visible() bool{
    return self.Object.Get("visible").Bool()
}

// SetVisibleA A Plugin with visible=true has its render and postRender methods called by the parent, otherwise they are skipped.
func (self *Plugin) SetVisibleA(member bool) {
    self.Object.Set("visible", member)
}

// HasPreUpdate A flag to indicate if this plugin has a preUpdate method.
func (self *Plugin) HasPreUpdate() bool{
    return self.Object.Get("hasPreUpdate").Bool()
}

// SetHasPreUpdateA A flag to indicate if this plugin has a preUpdate method.
func (self *Plugin) SetHasPreUpdateA(member bool) {
    self.Object.Set("hasPreUpdate", member)
}

// HasUpdate A flag to indicate if this plugin has an update method.
func (self *Plugin) HasUpdate() bool{
    return self.Object.Get("hasUpdate").Bool()
}

// SetHasUpdateA A flag to indicate if this plugin has an update method.
func (self *Plugin) SetHasUpdateA(member bool) {
    self.Object.Set("hasUpdate", member)
}

// HasPostUpdate A flag to indicate if this plugin has a postUpdate method.
func (self *Plugin) HasPostUpdate() bool{
    return self.Object.Get("hasPostUpdate").Bool()
}

// SetHasPostUpdateA A flag to indicate if this plugin has a postUpdate method.
func (self *Plugin) SetHasPostUpdateA(member bool) {
    self.Object.Set("hasPostUpdate", member)
}

// HasRender A flag to indicate if this plugin has a render method.
func (self *Plugin) HasRender() bool{
    return self.Object.Get("hasRender").Bool()
}

// SetHasRenderA A flag to indicate if this plugin has a render method.
func (self *Plugin) SetHasRenderA(member bool) {
    self.Object.Set("hasRender", member)
}

// HasPostRender A flag to indicate if this plugin has a postRender method.
func (self *Plugin) HasPostRender() bool{
    return self.Object.Get("hasPostRender").Bool()
}

// SetHasPostRenderA A flag to indicate if this plugin has a postRender method.
func (self *Plugin) SetHasPostRenderA(member bool) {
    self.Object.Set("hasPostRender", member)
}


// PreUpdate Pre-update is called at the very start of the update cycle, before any other subsystems have been updated (including Physics).
// It is only called if active is set to true.
func (self *Plugin) PreUpdate() {
    self.Object.Call("preUpdate")
}

// PreUpdateI Pre-update is called at the very start of the update cycle, before any other subsystems have been updated (including Physics).
// It is only called if active is set to true.
func (self *Plugin) PreUpdateI(args ...interface{}) {
    self.Object.Call("preUpdate", args)
}

// Update Update is called after all the core subsystems (Input, Tweens, Sound, etc) and the State have updated, but before the render.
// It is only called if active is set to true.
func (self *Plugin) Update() {
    self.Object.Call("update")
}

// UpdateI Update is called after all the core subsystems (Input, Tweens, Sound, etc) and the State have updated, but before the render.
// It is only called if active is set to true.
func (self *Plugin) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Render Render is called right after the Game Renderer completes, but before the State.render.
// It is only called if visible is set to true.
func (self *Plugin) Render() {
    self.Object.Call("render")
}

// RenderI Render is called right after the Game Renderer completes, but before the State.render.
// It is only called if visible is set to true.
func (self *Plugin) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// PostRender Post-render is called after the Game Renderer and State.render have run.
// It is only called if visible is set to true.
func (self *Plugin) PostRender() {
    self.Object.Call("postRender")
}

// PostRenderI Post-render is called after the Game Renderer and State.render have run.
// It is only called if visible is set to true.
func (self *Plugin) PostRenderI(args ...interface{}) {
    self.Object.Call("postRender", args)
}

// Destroy Clear down this Plugin and null out references
func (self *Plugin) Destroy() {
    self.Object.Call("destroy")
}

// DestroyI Clear down this Plugin and null out references
func (self *Plugin) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}

