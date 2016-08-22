// Automatic generation for Phaser.Component.ScaleMinMax
// generated file ComponentScaleMinMax.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The ScaleMinMax component allows a Game Object to limit how far it can be scaled by its parent.
type ComponentScaleMinMax struct {
    *js.Object
}


// The ScaleMinMax component allows a Game Object to limit how far it can be scaled by its parent.
func NewComponentScaleMinMax() *ComponentScaleMinMax {
    return &ComponentScaleMinMax{js.Global.Call("Phaser.Component.ScaleMinMax")}
}

// The ScaleMinMax component allows a Game Object to limit how far it can be scaled by its parent.
func NewComponentScaleMinMaxI(args ...interface{}) *ComponentScaleMinMax {
    return &ComponentScaleMinMax{js.Global.Call("Phaser.Component.ScaleMinMax", args)}
}



// The callback that will apply any scale limiting to the worldTransform.
func (self *ComponentScaleMinMax) SetTransformCallbackA(member func(...interface{})) {
    self.Object.Set("transformCallback", member)
}

// The context under which `transformCallback` is called.
func (self *ComponentScaleMinMax) GetTransformCallbackContextA() interface{}{
    return self.Object.Get("transformCallbackContext")
}

// The context under which `transformCallback` is called.
func (self *ComponentScaleMinMax) SetTransformCallbackContextA(member interface{}) {
    self.Object.Set("transformCallbackContext", member)
}

// The minimum scale this Game Object will scale down to.
// 
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
// 
// Set it to `null` to remove the limit.
func (self *ComponentScaleMinMax) GetScaleMinA() *Point{
    return &Point{self.Object.Get("scaleMin")}
}

// The minimum scale this Game Object will scale down to.
// 
// It allows you to prevent a parent from scaling this Game Object lower than the given value.
// 
// Set it to `null` to remove the limit.
func (self *ComponentScaleMinMax) SetScaleMinA(member *Point) {
    self.Object.Set("scaleMin", member)
}

// The maximum scale this Game Object will scale up to. 
// 
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
// 
// Set it to `null` to remove the limit.
func (self *ComponentScaleMinMax) GetScaleMaxA() *Point{
    return &Point{self.Object.Get("scaleMax")}
}

// The maximum scale this Game Object will scale up to. 
// 
// It allows you to prevent a parent from scaling this Game Object higher than the given value.
// 
// Set it to `null` to remove the limit.
func (self *ComponentScaleMinMax) SetScaleMaxA(member *Point) {
    self.Object.Set("scaleMax", member)
}



// Adjust scaling limits, if set, to this Game Object.
func (self *ComponentScaleMinMax) CheckTransform(wt *Matrix) {
    self.Object.Call("checkTransform", wt)
}

// Adjust scaling limits, if set, to this Game Object.
func (self *ComponentScaleMinMax) CheckTransformI(args ...interface{}) {
    self.Object.Call("checkTransform", args)
}

// Sets the scaleMin and scaleMax values. These values are used to limit how far this Game Object will scale based on its parent.
// 
// For example if this Game Object has a `minScale` value of 1 and its parent has a `scale` value of 0.5, the 0.5 will be ignored 
// and the scale value of 1 will be used, as the parents scale is lower than the minimum scale this Game Object should adhere to.
// 
// By setting these values you can carefully control how Game Objects deal with responsive scaling.
// 
// If only one parameter is given then that value will be used for both scaleMin and scaleMax:
// `setScaleMinMax(1)` = scaleMin.x, scaleMin.y, scaleMax.x and scaleMax.y all = 1
// 
// If only two parameters are given the first is set as scaleMin.x and y and the second as scaleMax.x and y:
// `setScaleMinMax(0.5, 2)` = scaleMin.x and y = 0.5 and scaleMax.x and y = 2
// 
// If you wish to set `scaleMin` with different values for x and y then either modify Game Object.scaleMin directly, 
// or pass `null` for the `maxX` and `maxY` parameters.
// 
// Call `setScaleMinMax(null)` to clear all previously set values.
func (self *ComponentScaleMinMax) SetScaleMinMax(minX interface{}, minY interface{}, maxX interface{}, maxY interface{}) {
    self.Object.Call("setScaleMinMax", minX, minY, maxX, maxY)
}

// Sets the scaleMin and scaleMax values. These values are used to limit how far this Game Object will scale based on its parent.
// 
// For example if this Game Object has a `minScale` value of 1 and its parent has a `scale` value of 0.5, the 0.5 will be ignored 
// and the scale value of 1 will be used, as the parents scale is lower than the minimum scale this Game Object should adhere to.
// 
// By setting these values you can carefully control how Game Objects deal with responsive scaling.
// 
// If only one parameter is given then that value will be used for both scaleMin and scaleMax:
// `setScaleMinMax(1)` = scaleMin.x, scaleMin.y, scaleMax.x and scaleMax.y all = 1
// 
// If only two parameters are given the first is set as scaleMin.x and y and the second as scaleMax.x and y:
// `setScaleMinMax(0.5, 2)` = scaleMin.x and y = 0.5 and scaleMax.x and y = 2
// 
// If you wish to set `scaleMin` with different values for x and y then either modify Game Object.scaleMin directly, 
// or pass `null` for the `maxX` and `maxY` parameters.
// 
// Call `setScaleMinMax(null)` to clear all previously set values.
func (self *ComponentScaleMinMax) SetScaleMinMaxI(args ...interface{}) {
    self.Object.Call("setScaleMinMax", args)
}
