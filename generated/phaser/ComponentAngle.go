// Automatic generation for Phaser.Component.Angle
// generated file ComponentAngle.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Angle Component provides access to an `angle` property; the rotation of a Game Object in degrees.
type ComponentAngle struct {
    *js.Object
}


// The Angle Component provides access to an `angle` property; the rotation of a Game Object in degrees.
func NewComponentAngle() *ComponentAngle {
    return &ComponentAngle{js.Global.Call("Phaser.Component.Angle")}
}

// The Angle Component provides access to an `angle` property; the rotation of a Game Object in degrees.
func NewComponentAngleI(args ...interface{}) *ComponentAngle {
    return &ComponentAngle{js.Global.Call("Phaser.Component.Angle", args)}
}



// The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *ComponentAngle) GetAngleA() int{
    return self.Object.Get("angle").Int()
}

// The angle property is the rotation of the Game Object in *degrees* from its original orientation.
// 
// Values from 0 to 180 represent clockwise rotation; values from 0 to -180 represent counterclockwise rotation.
// 
// Values outside this range are added to or subtracted from 360 to obtain a value within the range. 
// For example, the statement player.angle = 450 is the same as player.angle = 90.
// 
// If you wish to work in radians instead of degrees you can use the property `rotation` instead. 
// Working in radians is slightly faster as it doesn't have to perform any calculations.
func (self *ComponentAngle) SetAngleA(member int) {
    self.Object.Set("angle", member)
}


