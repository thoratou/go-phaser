// Automatic generation for PIXI.PIXI.DisplayObject
// generated file DisplayObject.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// The base class for all objects that are rendered. Contains properties for position, scaling,
// rotation, masks and cache handling.
// 
// This is an abstract class and should not be used on its own, rather it should be extended.
// 
// It is used internally by the likes of PIXI.Sprite.
type DisplayObject struct {
    *js.Object
}


// The coordinates, in pixels, of this DisplayObject, relative to its parent container.
// 
// The value of this property does not reflect any positioning happening further up the display list.
// To obtain that value please see the `worldPosition` property.
func (self *DisplayObject) GetPosition() *PIXIPoint{
    return &PIXIPoint{self.Get("position")}
}

// The coordinates, in pixels, of this DisplayObject, relative to its parent container.
// 
// The value of this property does not reflect any positioning happening further up the display list.
// To obtain that value please see the `worldPosition` property.
func (self *DisplayObject) SetPosition(member *PIXIPoint) {
    self.Set("position", member)
}

// The scale of this DisplayObject. A scale of 1:1 represents the DisplayObject
// at its default size. A value of 0.5 would scale this DisplayObject by half, and so on.
// 
// The value of this property does not reflect any scaling happening further up the display list.
// To obtain that value please see the `worldScale` property.
func (self *DisplayObject) GetScale() *PIXIPoint{
    return &PIXIPoint{self.Get("scale")}
}

// The scale of this DisplayObject. A scale of 1:1 represents the DisplayObject
// at its default size. A value of 0.5 would scale this DisplayObject by half, and so on.
// 
// The value of this property does not reflect any scaling happening further up the display list.
// To obtain that value please see the `worldScale` property.
func (self *DisplayObject) SetScale(member *PIXIPoint) {
    self.Set("scale", member)
}

// The pivot point of this DisplayObject that it rotates around. The values are expressed
// in pixel values.
func (self *DisplayObject) GetPivot() *PIXIPoint{
    return &PIXIPoint{self.Get("pivot")}
}

// The pivot point of this DisplayObject that it rotates around. The values are expressed
// in pixel values.
func (self *DisplayObject) SetPivot(member *PIXIPoint) {
    self.Set("pivot", member)
}

// The rotation of this DisplayObject. The value is given, and expressed, in radians, and is based on
// a right-handed orientation.
// 
// The value of this property does not reflect any rotation happening further up the display list.
// To obtain that value please see the `worldRotation` property.
func (self *DisplayObject) GetRotation() float64{
    return self.Get("rotation").Float()
}

// The rotation of this DisplayObject. The value is given, and expressed, in radians, and is based on
// a right-handed orientation.
// 
// The value of this property does not reflect any rotation happening further up the display list.
// To obtain that value please see the `worldRotation` property.
func (self *DisplayObject) SetRotation(member float64) {
    self.Set("rotation", member)
}

// The alpha value of this DisplayObject. A value of 1 is fully opaque. A value of 0 is transparent.
// Please note that an object with an alpha value of 0 is skipped during the render pass.
// 
// The value of this property does not reflect any alpha values set further up the display list.
// To obtain that value please see the `worldAlpha` property.
func (self *DisplayObject) GetAlpha() float64{
    return self.Get("alpha").Float()
}

// The alpha value of this DisplayObject. A value of 1 is fully opaque. A value of 0 is transparent.
// Please note that an object with an alpha value of 0 is skipped during the render pass.
// 
// The value of this property does not reflect any alpha values set further up the display list.
// To obtain that value please see the `worldAlpha` property.
func (self *DisplayObject) SetAlpha(member float64) {
    self.Set("alpha", member)
}

// The visibility of this DisplayObject. A value of `false` makes the object invisible.
// A value of `true` makes it visible. Please note that an object with a visible value of
// `false` is skipped during the render pass. Equally a DisplayObject with visible false will
// not render any of its children.
// 
// The value of this property does not reflect any visible values set further up the display list.
// To obtain that value please see the `worldVisible` property.
func (self *DisplayObject) GetVisible() bool{
    return self.Get("visible").Bool()
}

// The visibility of this DisplayObject. A value of `false` makes the object invisible.
// A value of `true` makes it visible. Please note that an object with a visible value of
// `false` is skipped during the render pass. Equally a DisplayObject with visible false will
// not render any of its children.
// 
// The value of this property does not reflect any visible values set further up the display list.
// To obtain that value please see the `worldVisible` property.
func (self *DisplayObject) SetVisible(member bool) {
    self.Set("visible", member)
}

// This is the defined area that will pick up mouse / touch events. It is null by default.
// Setting it is a neat way of optimising the hitTest function that the interactionManager will use (as it will not need to hit test all the children)
func (self *DisplayObject) GetHitArea() interface{}{
    return self.Get("hitArea")
}

// This is the defined area that will pick up mouse / touch events. It is null by default.
// Setting it is a neat way of optimising the hitTest function that the interactionManager will use (as it will not need to hit test all the children)
func (self *DisplayObject) SetHitArea(member interface{}) {
    self.Set("hitArea", member)
}

// Should this DisplayObject be rendered by the renderer? An object with a renderable value of
// `false` is skipped during the render pass.
func (self *DisplayObject) GetRenderable() bool{
    return self.Get("renderable").Bool()
}

// Should this DisplayObject be rendered by the renderer? An object with a renderable value of
// `false` is skipped during the render pass.
func (self *DisplayObject) SetRenderable(member bool) {
    self.Set("renderable", member)
}

// The parent DisplayObjectContainer that this DisplayObject is a child of.
// All DisplayObjects must belong to a parent in order to be rendered.
// The root parent is the Stage object. This property is set automatically when the
// DisplayObject is added to, or removed from, a DisplayObjectContainer.
func (self *DisplayObject) GetParent() *PIXIDisplayObjectContainer{
    return &PIXIDisplayObjectContainer{self.Get("parent")}
}

// The parent DisplayObjectContainer that this DisplayObject is a child of.
// All DisplayObjects must belong to a parent in order to be rendered.
// The root parent is the Stage object. This property is set automatically when the
// DisplayObject is added to, or removed from, a DisplayObjectContainer.
func (self *DisplayObject) SetParent(member *PIXIDisplayObjectContainer) {
    self.Set("parent", member)
}

// The stage that this DisplayObject is connected to.
func (self *DisplayObject) GetStage() *PIXIStage{
    return &PIXIStage{self.Get("stage")}
}

// The stage that this DisplayObject is connected to.
func (self *DisplayObject) SetStage(member *PIXIStage) {
    self.Set("stage", member)
}

// The multiplied alpha value of this DisplayObject. A value of 1 is fully opaque. A value of 0 is transparent.
// This value is the calculated total, based on the alpha values of all parents of this DisplayObjects 
// in the display list.
// 
// To obtain, and set, the local alpha value, see the `alpha` property.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) GetWorldAlpha() float64{
    return self.Get("worldAlpha").Float()
}

// The multiplied alpha value of this DisplayObject. A value of 1 is fully opaque. A value of 0 is transparent.
// This value is the calculated total, based on the alpha values of all parents of this DisplayObjects 
// in the display list.
// 
// To obtain, and set, the local alpha value, see the `alpha` property.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldAlpha(member float64) {
    self.Set("worldAlpha", member)
}

// The current transform of this DisplayObject.
// 
// This property contains the calculated total, based on the transforms of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) GetWorldTransform() *PIXIMatrix{
    return &PIXIMatrix{self.Get("worldTransform")}
}

// The current transform of this DisplayObject.
// 
// This property contains the calculated total, based on the transforms of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldTransform(member *PIXIMatrix) {
    self.Set("worldTransform", member)
}

// The coordinates, in pixels, of this DisplayObject within the world.
// 
// This property contains the calculated total, based on the positions of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) GetWorldPosition() *PIXIPoint{
    return &PIXIPoint{self.Get("worldPosition")}
}

// The coordinates, in pixels, of this DisplayObject within the world.
// 
// This property contains the calculated total, based on the positions of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldPosition(member *PIXIPoint) {
    self.Set("worldPosition", member)
}

// The global scale of this DisplayObject.
// 
// This property contains the calculated total, based on the scales of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) GetWorldScale() *PIXIPoint{
    return &PIXIPoint{self.Get("worldScale")}
}

// The global scale of this DisplayObject.
// 
// This property contains the calculated total, based on the scales of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldScale(member *PIXIPoint) {
    self.Set("worldScale", member)
}

// The rotation, in radians, of this DisplayObject.
// 
// This property contains the calculated total, based on the rotations of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) GetWorldRotation() float64{
    return self.Get("worldRotation").Float()
}

// The rotation, in radians, of this DisplayObject.
// 
// This property contains the calculated total, based on the rotations of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldRotation(member float64) {
    self.Set("worldRotation", member)
}

// The rectangular area used by filters when rendering a shader for this DisplayObject.
func (self *DisplayObject) GetFilterArea() *Rectangle{
    return &Rectangle{self.Get("filterArea")}
}

// The rectangular area used by filters when rendering a shader for this DisplayObject.
func (self *DisplayObject) SetFilterArea(member *Rectangle) {
    self.Set("filterArea", member)
}

// The horizontal position of the DisplayObject, in pixels, relative to its parent.
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) GetX() float64{
    return self.Get("x").Float()
}

// The horizontal position of the DisplayObject, in pixels, relative to its parent.
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) SetX(member float64) {
    self.Set("x", member)
}

// The vertical position of the DisplayObject, in pixels, relative to its parent.
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) GetY() float64{
    return self.Get("y").Float()
}

// The vertical position of the DisplayObject, in pixels, relative to its parent.
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) SetY(member float64) {
    self.Set("y", member)
}

// Indicates if this DisplayObject is visible, based on it, and all of its parents, `visible` property values.
func (self *DisplayObject) GetWorldVisible() bool{
    return self.Get("worldVisible").Bool()
}

// Indicates if this DisplayObject is visible, based on it, and all of its parents, `visible` property values.
func (self *DisplayObject) SetWorldVisible(member bool) {
    self.Set("worldVisible", member)
}

// Sets a mask for this DisplayObject. A mask is an instance of a Graphics object.
// When applied it limits the visible area of this DisplayObject to the shape of the mask.
// Under a Canvas renderer it uses shape clipping. Under a WebGL renderer it uses a Stencil Buffer.
// To remove a mask, set this property to `null`.
func (self *DisplayObject) GetMask() *PIXIGraphics{
    return &PIXIGraphics{self.Get("mask")}
}

// Sets a mask for this DisplayObject. A mask is an instance of a Graphics object.
// When applied it limits the visible area of this DisplayObject to the shape of the mask.
// Under a Canvas renderer it uses shape clipping. Under a WebGL renderer it uses a Stencil Buffer.
// To remove a mask, set this property to `null`.
func (self *DisplayObject) SetMask(member *PIXIGraphics) {
    self.Set("mask", member)
}

// Sets the filters for this DisplayObject. This is a WebGL only feature, and is ignored by the Canvas
// Renderer. A filter is a shader applied to this DisplayObject. You can modify the placement of the filter
// using `DisplayObject.filterArea`.
// 
// To remove filters, set this property to `null`.
// 
// Note: You cannot have a filter set, and a MULTIPLY Blend Mode active, at the same time. Setting a 
// filter will reset this DisplayObjects blend mode to NORMAL.
func (self *DisplayObject) GetFilters() []interface{}{
	array := self.Get("filters")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// Sets the filters for this DisplayObject. This is a WebGL only feature, and is ignored by the Canvas
// Renderer. A filter is a shader applied to this DisplayObject. You can modify the placement of the filter
// using `DisplayObject.filterArea`.
// 
// To remove filters, set this property to `null`.
// 
// Note: You cannot have a filter set, and a MULTIPLY Blend Mode active, at the same time. Setting a 
// filter will reset this DisplayObjects blend mode to NORMAL.
func (self *DisplayObject) SetFilters(member []interface{}) {
    self.Set("filters", member)
}

// Sets if this DisplayObject should be cached as a bitmap.
// 
// When invoked it will take a snapshot of the DisplayObject, as it is at that moment, and store it 
// in a RenderTexture. This is then used whenever this DisplayObject is rendered. It can provide a
// performance benefit for complex, but static, DisplayObjects. I.e. those with lots of children.
// 
// Cached Bitmaps do not track their parents. If you update a property of this DisplayObject, it will not
// re-generate the cached bitmap automatically. To do that you need to call `DisplayObject.updateCache`.
// 
// To remove a cached bitmap, set this property to `null`.
func (self *DisplayObject) GetCacheAsBitmap() bool{
    return self.Get("cacheAsBitmap").Bool()
}

// Sets if this DisplayObject should be cached as a bitmap.
// 
// When invoked it will take a snapshot of the DisplayObject, as it is at that moment, and store it 
// in a RenderTexture. This is then used whenever this DisplayObject is rendered. It can provide a
// performance benefit for complex, but static, DisplayObjects. I.e. those with lots of children.
// 
// Cached Bitmaps do not track their parents. If you update a property of this DisplayObject, it will not
// re-generate the cached bitmap automatically. To do that you need to call `DisplayObject.updateCache`.
// 
// To remove a cached bitmap, set this property to `null`.
func (self *DisplayObject) SetCacheAsBitmap(member bool) {
    self.Set("cacheAsBitmap", member)
}


