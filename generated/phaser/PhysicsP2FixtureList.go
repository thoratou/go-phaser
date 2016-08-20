// Automatic generation for Phaser.Physics.P2.FixtureList
// generated file PhysicsP2FixtureList.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// Allow to access a list of created fixture (coming from Body#addPhaserPolygon)
// which itself parse the input from PhysicsEditor with the custom phaser exporter.
// You can access fixtures of a Body by a group index or even by providing a fixture Key.
// You can set the fixture key and also the group index for a fixture in PhysicsEditor.
// This gives you the power to create a complex body built of many fixtures and modify them
// during runtime (to remove parts, set masks, categories & sensor properties)
type PhysicsP2FixtureList struct {
    *js.Object
}




// 
func (self *PhysicsP2FixtureList) InitI(args ...interface{}) {
    self.Call("init", args)
}

// 
func (self *PhysicsP2FixtureList) SetCategoryI(args ...interface{}) {
    self.Call("setCategory", args)
}

// 
func (self *PhysicsP2FixtureList) SetMaskI(args ...interface{}) {
    self.Call("setMask", args)
}

// 
func (self *PhysicsP2FixtureList) SetSensorI(args ...interface{}) {
    self.Call("setSensor", args)
}

// 
func (self *PhysicsP2FixtureList) SetMaterialI(args ...interface{}) {
    self.Call("setMaterial", args)
}

// Accessor to get either a list of specified fixtures by key or the whole fixture list
func (self *PhysicsP2FixtureList) GetFixturesI(args ...interface{}) {
    self.Call("getFixtures", args)
}

// Accessor to get either a single fixture by its key.
func (self *PhysicsP2FixtureList) GetFixtureByKeyI(args ...interface{}) {
    self.Call("getFixtureByKey", args)
}

// Accessor to get a group of fixtures by its group index.
func (self *PhysicsP2FixtureList) GetGroupI(args ...interface{}) {
    self.Call("getGroup", args)
}

// Parser for the output of Phaser.Physics.P2.Body#addPhaserPolygon
func (self *PhysicsP2FixtureList) ParseI(args ...interface{}) {
    self.Call("parse", args)
}

// A helper to flatten arrays. This is very useful as the fixtures are nested from time to time due to the way P2 creates and splits polygons.
func (self *PhysicsP2FixtureList) FlattenI(args ...interface{}) {
    self.Call("flatten", args)
}
