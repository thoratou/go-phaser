// Package phaser Automatic generation for PIXI.PIXI.DisplayObject
// generated file DisplayObject.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// DisplayObject The base class for all objects that are rendered. Contains properties for position, scaling,
// 
// rotation, masks and cache handling.
// 
// 
// 
// This is an abstract class and should not be used on its own, rather it should be extended.
// 
// 
// 
// It is used internally by the likes of PIXI.Sprite.
type DisplayObject struct {
    *js.Object
}

// NewDisplayObject The base class for all objects that are rendered. Contains properties for position, scaling,
// 
// rotation, masks and cache handling.
// 
// 
// 
// This is an abstract class and should not be used on its own, rather it should be extended.
// 
// 
// 
// It is used internally by the likes of PIXI.Sprite.
func NewDisplayObject() *DisplayObject {
    return &DisplayObject{js.Global.Get("PIXI").Get("PIXI").Get("DisplayObject").New()}
}
// NewDisplayObjectI The base class for all objects that are rendered. Contains properties for position, scaling,
// 
// rotation, masks and cache handling.
// 
// 
// 
// This is an abstract class and should not be used on its own, rather it should be extended.
// 
// 
// 
// It is used internally by the likes of PIXI.Sprite.
func NewDisplayObjectI(args ...interface{}) *DisplayObject {
    return &DisplayObject{js.Global.Get("PIXI").Get("PIXI").Get("DisplayObject").New(args)}
}



// DisplayObject Binding conversion method to DisplayObject point 
func ToDisplayObject(jsStruct interface{}) *DisplayObject {
    if object, ok := jsStruct.(*js.Object); ok {
		return &DisplayObject{Object: object}
	}
	return nil
}



// Position The coordinates, in pixels, of this DisplayObject, relative to its parent container.
// 
// 
// 
// The value of this property does not reflect any positioning happening further up the display list.
// 
// To obtain that value please see the `worldPosition` property.
func (self *DisplayObject) Position() *PIXIPoint{
    return &PIXIPoint{self.Object.Get("position")}
}

// SetPositionA The coordinates, in pixels, of this DisplayObject, relative to its parent container.
// 
// 
// 
// The value of this property does not reflect any positioning happening further up the display list.
// 
// To obtain that value please see the `worldPosition` property.
func (self *DisplayObject) SetPositionA(member *PIXIPoint) {
    self.Object.Set("position", member)
}

// Scale The scale of this DisplayObject. A scale of 1:1 represents the DisplayObject
// 
// at its default size. A value of 0.5 would scale this DisplayObject by half, and so on.
// 
// 
// 
// The value of this property does not reflect any scaling happening further up the display list.
// 
// To obtain that value please see the `worldScale` property.
func (self *DisplayObject) Scale() *PIXIPoint{
    return &PIXIPoint{self.Object.Get("scale")}
}

// SetScaleA The scale of this DisplayObject. A scale of 1:1 represents the DisplayObject
// 
// at its default size. A value of 0.5 would scale this DisplayObject by half, and so on.
// 
// 
// 
// The value of this property does not reflect any scaling happening further up the display list.
// 
// To obtain that value please see the `worldScale` property.
func (self *DisplayObject) SetScaleA(member *PIXIPoint) {
    self.Object.Set("scale", member)
}

// Pivot The pivot point of this DisplayObject that it rotates around. The values are expressed
// 
// in pixel values.
func (self *DisplayObject) Pivot() *PIXIPoint{
    return &PIXIPoint{self.Object.Get("pivot")}
}

// SetPivotA The pivot point of this DisplayObject that it rotates around. The values are expressed
// 
// in pixel values.
func (self *DisplayObject) SetPivotA(member *PIXIPoint) {
    self.Object.Set("pivot", member)
}

// Rotation The rotation of this DisplayObject. The value is given, and expressed, in radians, and is based on
// 
// a right-handed orientation.
// 
// 
// 
// The value of this property does not reflect any rotation happening further up the display list.
// 
// To obtain that value please see the `worldRotation` property.
func (self *DisplayObject) Rotation() int{
    return self.Object.Get("rotation").Int()
}

// SetRotationA The rotation of this DisplayObject. The value is given, and expressed, in radians, and is based on
// 
// a right-handed orientation.
// 
// 
// 
// The value of this property does not reflect any rotation happening further up the display list.
// 
// To obtain that value please see the `worldRotation` property.
func (self *DisplayObject) SetRotationA(member int) {
    self.Object.Set("rotation", member)
}

// Alpha The alpha value of this DisplayObject. A value of 1 is fully opaque. A value of 0 is transparent.
// 
// Please note that an object with an alpha value of 0 is skipped during the render pass.
// 
// 
// 
// The value of this property does not reflect any alpha values set further up the display list.
// 
// To obtain that value please see the `worldAlpha` property.
func (self *DisplayObject) Alpha() int{
    return self.Object.Get("alpha").Int()
}

// SetAlphaA The alpha value of this DisplayObject. A value of 1 is fully opaque. A value of 0 is transparent.
// 
// Please note that an object with an alpha value of 0 is skipped during the render pass.
// 
// 
// 
// The value of this property does not reflect any alpha values set further up the display list.
// 
// To obtain that value please see the `worldAlpha` property.
func (self *DisplayObject) SetAlphaA(member int) {
    self.Object.Set("alpha", member)
}

// Visible The visibility of this DisplayObject. A value of `false` makes the object invisible.
// 
// A value of `true` makes it visible. Please note that an object with a visible value of
// 
// `false` is skipped during the render pass. Equally a DisplayObject with visible false will
// 
// not render any of its children.
// 
// 
// 
// The value of this property does not reflect any visible values set further up the display list.
// 
// To obtain that value please see the `worldVisible` property.
func (self *DisplayObject) Visible() bool{
    return self.Object.Get("visible").Bool()
}

// SetVisibleA The visibility of this DisplayObject. A value of `false` makes the object invisible.
// 
// A value of `true` makes it visible. Please note that an object with a visible value of
// 
// `false` is skipped during the render pass. Equally a DisplayObject with visible false will
// 
// not render any of its children.
// 
// 
// 
// The value of this property does not reflect any visible values set further up the display list.
// 
// To obtain that value please see the `worldVisible` property.
func (self *DisplayObject) SetVisibleA(member bool) {
    self.Object.Set("visible", member)
}

// HitArea This is the defined area that will pick up mouse / touch events. It is null by default.
// 
// Setting it is a neat way of optimising the hitTest function that the interactionManager will use (as it will not need to hit test all the children)
func (self *DisplayObject) HitArea() interface{}{
    return self.Object.Get("hitArea")
}

// SetHitAreaA This is the defined area that will pick up mouse / touch events. It is null by default.
// 
// Setting it is a neat way of optimising the hitTest function that the interactionManager will use (as it will not need to hit test all the children)
func (self *DisplayObject) SetHitAreaA(member interface{}) {
    self.Object.Set("hitArea", member)
}

// Renderable Should this DisplayObject be rendered by the renderer? An object with a renderable value of
// 
// `false` is skipped during the render pass.
func (self *DisplayObject) Renderable() bool{
    return self.Object.Get("renderable").Bool()
}

// SetRenderableA Should this DisplayObject be rendered by the renderer? An object with a renderable value of
// 
// `false` is skipped during the render pass.
func (self *DisplayObject) SetRenderableA(member bool) {
    self.Object.Set("renderable", member)
}

// Parent The parent DisplayObjectContainer that this DisplayObject is a child of.
// 
// All DisplayObjects must belong to a parent in order to be rendered.
// 
// The root parent is the Stage object. This property is set automatically when the
// 
// DisplayObject is added to, or removed from, a DisplayObjectContainer.
func (self *DisplayObject) Parent() *PIXIDisplayObjectContainer{
    return &PIXIDisplayObjectContainer{self.Object.Get("parent")}
}

// SetParentA The parent DisplayObjectContainer that this DisplayObject is a child of.
// 
// All DisplayObjects must belong to a parent in order to be rendered.
// 
// The root parent is the Stage object. This property is set automatically when the
// 
// DisplayObject is added to, or removed from, a DisplayObjectContainer.
func (self *DisplayObject) SetParentA(member *PIXIDisplayObjectContainer) {
    self.Object.Set("parent", member)
}

// WorldAlpha The multiplied alpha value of this DisplayObject. A value of 1 is fully opaque. A value of 0 is transparent.
// 
// This value is the calculated total, based on the alpha values of all parents of this DisplayObjects 
// 
// in the display list.
// 
// 
// 
// To obtain, and set, the local alpha value, see the `alpha` property.
// 
// 
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// 
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) WorldAlpha() int{
    return self.Object.Get("worldAlpha").Int()
}

// SetWorldAlphaA The multiplied alpha value of this DisplayObject. A value of 1 is fully opaque. A value of 0 is transparent.
// 
// This value is the calculated total, based on the alpha values of all parents of this DisplayObjects 
// 
// in the display list.
// 
// 
// 
// To obtain, and set, the local alpha value, see the `alpha` property.
// 
// 
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// 
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldAlphaA(member int) {
    self.Object.Set("worldAlpha", member)
}

// WorldTransform The current transform of this DisplayObject.
// 
// 
// 
// This property contains the calculated total, based on the transforms of all parents of this 
// 
// DisplayObject in the display list.
// 
// 
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// 
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) WorldTransform() *PIXIMatrix{
    return &PIXIMatrix{self.Object.Get("worldTransform")}
}

// SetWorldTransformA The current transform of this DisplayObject.
// 
// 
// 
// This property contains the calculated total, based on the transforms of all parents of this 
// 
// DisplayObject in the display list.
// 
// 
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// 
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldTransformA(member *PIXIMatrix) {
    self.Object.Set("worldTransform", member)
}

// WorldPosition The coordinates, in pixels, of this DisplayObject within the world.
// 
// 
// 
// This property contains the calculated total, based on the positions of all parents of this 
// 
// DisplayObject in the display list.
// 
// 
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// 
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) WorldPosition() *PIXIPoint{
    return &PIXIPoint{self.Object.Get("worldPosition")}
}

// SetWorldPositionA The coordinates, in pixels, of this DisplayObject within the world.
// 
// 
// 
// This property contains the calculated total, based on the positions of all parents of this 
// 
// DisplayObject in the display list.
// 
// 
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// 
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldPositionA(member *PIXIPoint) {
    self.Object.Set("worldPosition", member)
}

// WorldScale The global scale of this DisplayObject.
// 
// 
// 
// This property contains the calculated total, based on the scales of all parents of this 
// 
// DisplayObject in the display list.
// 
// 
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// 
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) WorldScale() *PIXIPoint{
    return &PIXIPoint{self.Object.Get("worldScale")}
}

// SetWorldScaleA The global scale of this DisplayObject.
// 
// 
// 
// This property contains the calculated total, based on the scales of all parents of this 
// 
// DisplayObject in the display list.
// 
// 
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// 
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldScaleA(member *PIXIPoint) {
    self.Object.Set("worldScale", member)
}

// WorldRotation The rotation, in radians, of this DisplayObject.
// 
// 
// 
// This property contains the calculated total, based on the rotations of all parents of this 
// 
// DisplayObject in the display list.
// 
// 
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// 
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) WorldRotation() int{
    return self.Object.Get("worldRotation").Int()
}

// SetWorldRotationA The rotation, in radians, of this DisplayObject.
// 
// 
// 
// This property contains the calculated total, based on the rotations of all parents of this 
// 
// DisplayObject in the display list.
// 
// 
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// 
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldRotationA(member int) {
    self.Object.Set("worldRotation", member)
}

// FilterArea The rectangular area used by filters when rendering a shader for this DisplayObject.
func (self *DisplayObject) FilterArea() *Rectangle{
    return &Rectangle{self.Object.Get("filterArea")}
}

// SetFilterAreaA The rectangular area used by filters when rendering a shader for this DisplayObject.
func (self *DisplayObject) SetFilterAreaA(member *Rectangle) {
    self.Object.Set("filterArea", member)
}

// X The horizontal position of the DisplayObject, in pixels, relative to its parent.
// 
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) X() int{
    return self.Object.Get("x").Int()
}

// SetXA The horizontal position of the DisplayObject, in pixels, relative to its parent.
// 
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) SetXA(member int) {
    self.Object.Set("x", member)
}

// Y The vertical position of the DisplayObject, in pixels, relative to its parent.
// 
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) Y() int{
    return self.Object.Get("y").Int()
}

// SetYA The vertical position of the DisplayObject, in pixels, relative to its parent.
// 
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) SetYA(member int) {
    self.Object.Set("y", member)
}

// WorldVisible Indicates if this DisplayObject is visible, based on it, and all of its parents, `visible` property values.
func (self *DisplayObject) WorldVisible() bool{
    return self.Object.Get("worldVisible").Bool()
}

// SetWorldVisibleA Indicates if this DisplayObject is visible, based on it, and all of its parents, `visible` property values.
func (self *DisplayObject) SetWorldVisibleA(member bool) {
    self.Object.Set("worldVisible", member)
}

// Mask Sets a mask for this DisplayObject. A mask is an instance of a Graphics object.
// 
// When applied it limits the visible area of this DisplayObject to the shape of the mask.
// 
// Under a Canvas renderer it uses shape clipping. Under a WebGL renderer it uses a Stencil Buffer.
// 
// To remove a mask, set this property to `null`.
func (self *DisplayObject) Mask() *PIXIGraphics{
    return &PIXIGraphics{self.Object.Get("mask")}
}

// SetMaskA Sets a mask for this DisplayObject. A mask is an instance of a Graphics object.
// 
// When applied it limits the visible area of this DisplayObject to the shape of the mask.
// 
// Under a Canvas renderer it uses shape clipping. Under a WebGL renderer it uses a Stencil Buffer.
// 
// To remove a mask, set this property to `null`.
func (self *DisplayObject) SetMaskA(member *PIXIGraphics) {
    self.Object.Set("mask", member)
}

// Filters Sets the filters for this DisplayObject. This is a WebGL only feature, and is ignored by the Canvas
// 
// Renderer. A filter is a shader applied to this DisplayObject. You can modify the placement of the filter
// 
// using `DisplayObject.filterArea`.
// 
// 
// 
// To remove filters, set this property to `null`.
// 
// 
// 
// Note: You cannot have a filter set, and a MULTIPLY Blend Mode active, at the same time. Setting a 
// 
// filter will reset this DisplayObjects blend mode to NORMAL.
func (self *DisplayObject) Filters() []interface{}{
	array00 := self.Object.Get("filters")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetFiltersA Sets the filters for this DisplayObject. This is a WebGL only feature, and is ignored by the Canvas
// 
// Renderer. A filter is a shader applied to this DisplayObject. You can modify the placement of the filter
// 
// using `DisplayObject.filterArea`.
// 
// 
// 
// To remove filters, set this property to `null`.
// 
// 
// 
// Note: You cannot have a filter set, and a MULTIPLY Blend Mode active, at the same time. Setting a 
// 
// filter will reset this DisplayObjects blend mode to NORMAL.
func (self *DisplayObject) SetFiltersA(member []interface{}) {
    self.Object.Set("filters", member)
}

// CacheAsBitmap Sets if this DisplayObject should be cached as a bitmap.
// 
// 
// 
// When invoked it will take a snapshot of the DisplayObject, as it is at that moment, and store it 
// 
// in a RenderTexture. This is then used whenever this DisplayObject is rendered. It can provide a
// 
// performance benefit for complex, but static, DisplayObjects. I.e. those with lots of children.
// 
// 
// 
// Cached Bitmaps do not track their parents. If you update a property of this DisplayObject, it will not
// 
// re-generate the cached bitmap automatically. To do that you need to call `DisplayObject.updateCache`.
// 
// 
// 
// To remove a cached bitmap, set this property to `null`.
func (self *DisplayObject) CacheAsBitmap() bool{
    return self.Object.Get("cacheAsBitmap").Bool()
}

// SetCacheAsBitmapA Sets if this DisplayObject should be cached as a bitmap.
// 
// 
// 
// When invoked it will take a snapshot of the DisplayObject, as it is at that moment, and store it 
// 
// in a RenderTexture. This is then used whenever this DisplayObject is rendered. It can provide a
// 
// performance benefit for complex, but static, DisplayObjects. I.e. those with lots of children.
// 
// 
// 
// Cached Bitmaps do not track their parents. If you update a property of this DisplayObject, it will not
// 
// re-generate the cached bitmap automatically. To do that you need to call `DisplayObject.updateCache`.
// 
// 
// 
// To remove a cached bitmap, set this property to `null`.
func (self *DisplayObject) SetCacheAsBitmapA(member bool) {
    self.Object.Set("cacheAsBitmap", member)
}


