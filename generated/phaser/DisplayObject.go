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


// The base class for all objects that are rendered. Contains properties for position, scaling,
// rotation, masks and cache handling.
// 
// This is an abstract class and should not be used on its own, rather it should be extended.
// 
// It is used internally by the likes of PIXI.Sprite.
func NewDisplayObject() *DisplayObject {
    return &DisplayObject{js.Global.Call("PIXI.PIXI.DisplayObject")}
}

// The base class for all objects that are rendered. Contains properties for position, scaling,
// rotation, masks and cache handling.
// 
// This is an abstract class and should not be used on its own, rather it should be extended.
// 
// It is used internally by the likes of PIXI.Sprite.
func NewDisplayObjectI(args ...interface{}) *DisplayObject {
    return &DisplayObject{js.Global.Call("PIXI.PIXI.DisplayObject", args)}
}



// The coordinates, in pixels, of this DisplayObject, relative to its parent container.
// 
// The value of this property does not reflect any positioning happening further up the display list.
// To obtain that value please see the `worldPosition` property.
func (self *DisplayObject) GetPositionA() *PIXIPoint{
    return &PIXIPoint{self.Object.Get("position")}
}

// The coordinates, in pixels, of this DisplayObject, relative to its parent container.
// 
// The value of this property does not reflect any positioning happening further up the display list.
// To obtain that value please see the `worldPosition` property.
func (self *DisplayObject) SetPositionA(member *PIXIPoint) {
    self.Object.Set("position", member)
}

// The scale of this DisplayObject. A scale of 1:1 represents the DisplayObject
// at its default size. A value of 0.5 would scale this DisplayObject by half, and so on.
// 
// The value of this property does not reflect any scaling happening further up the display list.
// To obtain that value please see the `worldScale` property.
func (self *DisplayObject) GetScaleA() *PIXIPoint{
    return &PIXIPoint{self.Object.Get("scale")}
}

// The scale of this DisplayObject. A scale of 1:1 represents the DisplayObject
// at its default size. A value of 0.5 would scale this DisplayObject by half, and so on.
// 
// The value of this property does not reflect any scaling happening further up the display list.
// To obtain that value please see the `worldScale` property.
func (self *DisplayObject) SetScaleA(member *PIXIPoint) {
    self.Object.Set("scale", member)
}

// The pivot point of this DisplayObject that it rotates around. The values are expressed
// in pixel values.
func (self *DisplayObject) GetPivotA() *PIXIPoint{
    return &PIXIPoint{self.Object.Get("pivot")}
}

// The pivot point of this DisplayObject that it rotates around. The values are expressed
// in pixel values.
func (self *DisplayObject) SetPivotA(member *PIXIPoint) {
    self.Object.Set("pivot", member)
}

// The rotation of this DisplayObject. The value is given, and expressed, in radians, and is based on
// a right-handed orientation.
// 
// The value of this property does not reflect any rotation happening further up the display list.
// To obtain that value please see the `worldRotation` property.
func (self *DisplayObject) GetRotationA() int{
    return self.Object.Get("rotation").Int()
}

// The rotation of this DisplayObject. The value is given, and expressed, in radians, and is based on
// a right-handed orientation.
// 
// The value of this property does not reflect any rotation happening further up the display list.
// To obtain that value please see the `worldRotation` property.
func (self *DisplayObject) SetRotationA(member int) {
    self.Object.Set("rotation", member)
}

// The alpha value of this DisplayObject. A value of 1 is fully opaque. A value of 0 is transparent.
// Please note that an object with an alpha value of 0 is skipped during the render pass.
// 
// The value of this property does not reflect any alpha values set further up the display list.
// To obtain that value please see the `worldAlpha` property.
func (self *DisplayObject) GetAlphaA() int{
    return self.Object.Get("alpha").Int()
}

// The alpha value of this DisplayObject. A value of 1 is fully opaque. A value of 0 is transparent.
// Please note that an object with an alpha value of 0 is skipped during the render pass.
// 
// The value of this property does not reflect any alpha values set further up the display list.
// To obtain that value please see the `worldAlpha` property.
func (self *DisplayObject) SetAlphaA(member int) {
    self.Object.Set("alpha", member)
}

// The visibility of this DisplayObject. A value of `false` makes the object invisible.
// A value of `true` makes it visible. Please note that an object with a visible value of
// `false` is skipped during the render pass. Equally a DisplayObject with visible false will
// not render any of its children.
// 
// The value of this property does not reflect any visible values set further up the display list.
// To obtain that value please see the `worldVisible` property.
func (self *DisplayObject) GetVisibleA() bool{
    return self.Object.Get("visible").Bool()
}

// The visibility of this DisplayObject. A value of `false` makes the object invisible.
// A value of `true` makes it visible. Please note that an object with a visible value of
// `false` is skipped during the render pass. Equally a DisplayObject with visible false will
// not render any of its children.
// 
// The value of this property does not reflect any visible values set further up the display list.
// To obtain that value please see the `worldVisible` property.
func (self *DisplayObject) SetVisibleA(member bool) {
    self.Object.Set("visible", member)
}

// This is the defined area that will pick up mouse / touch events. It is null by default.
// Setting it is a neat way of optimising the hitTest function that the interactionManager will use (as it will not need to hit test all the children)
func (self *DisplayObject) GetHitAreaA() interface{}{
    return self.Object.Get("hitArea")
}

// This is the defined area that will pick up mouse / touch events. It is null by default.
// Setting it is a neat way of optimising the hitTest function that the interactionManager will use (as it will not need to hit test all the children)
func (self *DisplayObject) SetHitAreaA(member interface{}) {
    self.Object.Set("hitArea", member)
}

// Should this DisplayObject be rendered by the renderer? An object with a renderable value of
// `false` is skipped during the render pass.
func (self *DisplayObject) GetRenderableA() bool{
    return self.Object.Get("renderable").Bool()
}

// Should this DisplayObject be rendered by the renderer? An object with a renderable value of
// `false` is skipped during the render pass.
func (self *DisplayObject) SetRenderableA(member bool) {
    self.Object.Set("renderable", member)
}

// The parent DisplayObjectContainer that this DisplayObject is a child of.
// All DisplayObjects must belong to a parent in order to be rendered.
// The root parent is the Stage object. This property is set automatically when the
// DisplayObject is added to, or removed from, a DisplayObjectContainer.
func (self *DisplayObject) GetParentA() *PIXIDisplayObjectContainer{
    return &PIXIDisplayObjectContainer{self.Object.Get("parent")}
}

// The parent DisplayObjectContainer that this DisplayObject is a child of.
// All DisplayObjects must belong to a parent in order to be rendered.
// The root parent is the Stage object. This property is set automatically when the
// DisplayObject is added to, or removed from, a DisplayObjectContainer.
func (self *DisplayObject) SetParentA(member *PIXIDisplayObjectContainer) {
    self.Object.Set("parent", member)
}

// The stage that this DisplayObject is connected to.
func (self *DisplayObject) GetStageA() *PIXIStage{
    return &PIXIStage{self.Object.Get("stage")}
}

// The stage that this DisplayObject is connected to.
func (self *DisplayObject) SetStageA(member *PIXIStage) {
    self.Object.Set("stage", member)
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
func (self *DisplayObject) GetWorldAlphaA() int{
    return self.Object.Get("worldAlpha").Int()
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
func (self *DisplayObject) SetWorldAlphaA(member int) {
    self.Object.Set("worldAlpha", member)
}

// The current transform of this DisplayObject.
// 
// This property contains the calculated total, based on the transforms of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) GetWorldTransformA() *PIXIMatrix{
    return &PIXIMatrix{self.Object.Get("worldTransform")}
}

// The current transform of this DisplayObject.
// 
// This property contains the calculated total, based on the transforms of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldTransformA(member *PIXIMatrix) {
    self.Object.Set("worldTransform", member)
}

// The coordinates, in pixels, of this DisplayObject within the world.
// 
// This property contains the calculated total, based on the positions of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) GetWorldPositionA() *PIXIPoint{
    return &PIXIPoint{self.Object.Get("worldPosition")}
}

// The coordinates, in pixels, of this DisplayObject within the world.
// 
// This property contains the calculated total, based on the positions of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldPositionA(member *PIXIPoint) {
    self.Object.Set("worldPosition", member)
}

// The global scale of this DisplayObject.
// 
// This property contains the calculated total, based on the scales of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) GetWorldScaleA() *PIXIPoint{
    return &PIXIPoint{self.Object.Get("worldScale")}
}

// The global scale of this DisplayObject.
// 
// This property contains the calculated total, based on the scales of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldScaleA(member *PIXIPoint) {
    self.Object.Set("worldScale", member)
}

// The rotation, in radians, of this DisplayObject.
// 
// This property contains the calculated total, based on the rotations of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) GetWorldRotationA() int{
    return self.Object.Get("worldRotation").Int()
}

// The rotation, in radians, of this DisplayObject.
// 
// This property contains the calculated total, based on the rotations of all parents of this 
// DisplayObject in the display list.
// 
// Note: This property is only updated at the end of the `updateTransform` call, once per render. Until 
// that happens this property will contain values based on the previous frame. Be mindful of this if
// accessing this property outside of the normal game flow, i.e. from an asynchronous event callback.
func (self *DisplayObject) SetWorldRotationA(member int) {
    self.Object.Set("worldRotation", member)
}

// The rectangular area used by filters when rendering a shader for this DisplayObject.
func (self *DisplayObject) GetFilterAreaA() *Rectangle{
    return &Rectangle{self.Object.Get("filterArea")}
}

// The rectangular area used by filters when rendering a shader for this DisplayObject.
func (self *DisplayObject) SetFilterAreaA(member *Rectangle) {
    self.Object.Set("filterArea", member)
}

// The horizontal position of the DisplayObject, in pixels, relative to its parent.
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) GetXA() int{
    return self.Object.Get("x").Int()
}

// The horizontal position of the DisplayObject, in pixels, relative to its parent.
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) SetXA(member int) {
    self.Object.Set("x", member)
}

// The vertical position of the DisplayObject, in pixels, relative to its parent.
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) GetYA() int{
    return self.Object.Get("y").Int()
}

// The vertical position of the DisplayObject, in pixels, relative to its parent.
// If you need the world position of the DisplayObject, use `DisplayObject.worldPosition` instead.
func (self *DisplayObject) SetYA(member int) {
    self.Object.Set("y", member)
}

// Indicates if this DisplayObject is visible, based on it, and all of its parents, `visible` property values.
func (self *DisplayObject) GetWorldVisibleA() bool{
    return self.Object.Get("worldVisible").Bool()
}

// Indicates if this DisplayObject is visible, based on it, and all of its parents, `visible` property values.
func (self *DisplayObject) SetWorldVisibleA(member bool) {
    self.Object.Set("worldVisible", member)
}

// Sets a mask for this DisplayObject. A mask is an instance of a Graphics object.
// When applied it limits the visible area of this DisplayObject to the shape of the mask.
// Under a Canvas renderer it uses shape clipping. Under a WebGL renderer it uses a Stencil Buffer.
// To remove a mask, set this property to `null`.
func (self *DisplayObject) GetMaskA() *PIXIGraphics{
    return &PIXIGraphics{self.Object.Get("mask")}
}

// Sets a mask for this DisplayObject. A mask is an instance of a Graphics object.
// When applied it limits the visible area of this DisplayObject to the shape of the mask.
// Under a Canvas renderer it uses shape clipping. Under a WebGL renderer it uses a Stencil Buffer.
// To remove a mask, set this property to `null`.
func (self *DisplayObject) SetMaskA(member *PIXIGraphics) {
    self.Object.Set("mask", member)
}

// Sets the filters for this DisplayObject. This is a WebGL only feature, and is ignored by the Canvas
// Renderer. A filter is a shader applied to this DisplayObject. You can modify the placement of the filter
// using `DisplayObject.filterArea`.
// 
// To remove filters, set this property to `null`.
// 
// Note: You cannot have a filter set, and a MULTIPLY Blend Mode active, at the same time. Setting a 
// filter will reset this DisplayObjects blend mode to NORMAL.
func (self *DisplayObject) GetFiltersA() []interface{}{
	array00 := self.Object.Get("filters")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00).Interface()
	}
	return out00
}

// Sets the filters for this DisplayObject. This is a WebGL only feature, and is ignored by the Canvas
// Renderer. A filter is a shader applied to this DisplayObject. You can modify the placement of the filter
// using `DisplayObject.filterArea`.
// 
// To remove filters, set this property to `null`.
// 
// Note: You cannot have a filter set, and a MULTIPLY Blend Mode active, at the same time. Setting a 
// filter will reset this DisplayObjects blend mode to NORMAL.
func (self *DisplayObject) SetFiltersA(member []interface{}) {
    self.Object.Set("filters", member)
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
func (self *DisplayObject) GetCacheAsBitmapA() bool{
    return self.Object.Get("cacheAsBitmap").Bool()
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
func (self *DisplayObject) SetCacheAsBitmapA(member bool) {
    self.Object.Set("cacheAsBitmap", member)
}


