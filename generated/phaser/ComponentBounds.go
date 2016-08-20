// Automatic generation for Phaser.Component.Bounds
// generated file ComponentBounds.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The Bounds component contains properties related to the bounds of the Game Object.
type ComponentBounds struct {
    *js.Object
}


// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *ComponentBounds) GetOffsetX() float64{
    return self.Get("offsetX").Float()
}

// The amount the Game Object is visually offset from its x coordinate.
// This is the same as `width * anchor.x`.
// It will only be > 0 if anchor.x is not equal to zero.
func (self *ComponentBounds) SetOffsetX(member float64) {
    self.Set("offsetX", member)
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *ComponentBounds) GetOffsetY() float64{
    return self.Get("offsetY").Float()
}

// The amount the Game Object is visually offset from its y coordinate.
// This is the same as `height * anchor.y`.
// It will only be > 0 if anchor.y is not equal to zero.
func (self *ComponentBounds) SetOffsetY(member float64) {
    self.Set("offsetY", member)
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *ComponentBounds) GetCenterX() float64{
    return self.Get("centerX").Float()
}

// The center x coordinate of the Game Object.
// This is the same as `(x - offsetX) + (width / 2)`.
func (self *ComponentBounds) SetCenterX(member float64) {
    self.Set("centerX", member)
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *ComponentBounds) GetCenterY() float64{
    return self.Get("centerY").Float()
}

// The center y coordinate of the Game Object.
// This is the same as `(y - offsetY) + (height / 2)`.
func (self *ComponentBounds) SetCenterY(member float64) {
    self.Set("centerY", member)
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *ComponentBounds) GetLeft() float64{
    return self.Get("left").Float()
}

// The left coordinate of the Game Object.
// This is the same as `x - offsetX`.
func (self *ComponentBounds) SetLeft(member float64) {
    self.Set("left", member)
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *ComponentBounds) GetRight() float64{
    return self.Get("right").Float()
}

// The right coordinate of the Game Object.
// This is the same as `x + width - offsetX`.
func (self *ComponentBounds) SetRight(member float64) {
    self.Set("right", member)
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *ComponentBounds) GetTop() float64{
    return self.Get("top").Float()
}

// The y coordinate of the Game Object.
// This is the same as `y - offsetY`.
func (self *ComponentBounds) SetTop(member float64) {
    self.Set("top", member)
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *ComponentBounds) GetBottom() float64{
    return self.Get("bottom").Float()
}

// The sum of the y and height properties.
// This is the same as `y + height - offsetY`.
func (self *ComponentBounds) SetBottom(member float64) {
    self.Set("bottom", member)
}



// Aligns this Game Object within another Game Object, or Rectangle, known as the
// 'container', to one of 9 possible positions.
// 
// The container must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the container. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT`, `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_CENTER`, 
// `Phaser.CENTER`, `Phaser.RIGHT_CENTER`, `Phaser.BOTTOM_LEFT`, 
// `Phaser.BOTTOM_CENTER` and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// container, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignIn(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the containers bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the container bounds by that amount, and providing a positive
// one expands it.
func (self *ComponentBounds) AlignInI(args ...interface{}) interface{}{
    return self.Call("alignIn", args)
}

// Aligns this Game Object to the side of another Game Object, or Rectangle, known as the
// 'parent', in one of 11 possible positions.
// 
// The parent must be a Game Object, or Phaser.Rectangle object. This can include properties
// such as `World.bounds` or `Camera.view`, for aligning Game Objects within the world 
// and camera bounds. Or it can include other Sprites, Images, Text objects, BitmapText,
// TileSprites or Buttons.
// 
// Please note that aligning a Sprite to another Game Object does **not** make it a child of
// the parent. It simply modifies its position coordinates so it aligns with it.
// 
// The position constants you can use are:
// 
// `Phaser.TOP_LEFT` (default), `Phaser.TOP_CENTER`, `Phaser.TOP_RIGHT`, `Phaser.LEFT_TOP`, 
// `Phaser.LEFT_CENTER`, `Phaser.LEFT_BOTTOM`, `Phaser.RIGHT_TOP`, `Phaser.RIGHT_CENTER`, 
// `Phaser.RIGHT_BOTTOM`, `Phaser.BOTTOM_LEFT`, `Phaser.BOTTOM_CENTER` 
// and `Phaser.BOTTOM_RIGHT`.
// 
// The Game Objects are placed in such a way that their _bounds_ align with the
// parent, taking into consideration rotation, scale and the anchor property.
// This allows you to neatly align Game Objects, irrespective of their position value.
// 
// The optional `offsetX` and `offsetY` arguments allow you to apply extra spacing to the final
// aligned position of the Game Object. For example:
// 
// `sprite.alignTo(background, Phaser.BOTTOM_RIGHT, -20, -20)`
// 
// Would align the `sprite` to the bottom-right, but moved 20 pixels in from the corner.
// Think of the offsets as applying an adjustment to the parents bounds before the alignment takes place.
// So providing a negative offset will 'shrink' the parent bounds by that amount, and providing a positive
// one expands it.
func (self *ComponentBounds) AlignToI(args ...interface{}) interface{}{
    return self.Call("alignTo", args)
}
