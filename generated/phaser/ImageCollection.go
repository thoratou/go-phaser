// Package phaser Automatic generation for Phaser.ImageCollection
// generated file ImageCollection.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// ImageCollection An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
type ImageCollection struct {
    *js.Object
}

// NewImageCollection An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection(name string, firstgid int) *ImageCollection {
    return &ImageCollection{js.Global.Get("Phaser").Get("ImageCollection").New(name, firstgid)}
}
// NewImageCollection1O An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection1O(name string, firstgid int, width int) *ImageCollection {
    return &ImageCollection{js.Global.Get("Phaser").Get("ImageCollection").New(name, firstgid, width)}
}
// NewImageCollection2O An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection2O(name string, firstgid int, width int, height int) *ImageCollection {
    return &ImageCollection{js.Global.Get("Phaser").Get("ImageCollection").New(name, firstgid, width, height)}
}
// NewImageCollection3O An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection3O(name string, firstgid int, width int, height int, margin int) *ImageCollection {
    return &ImageCollection{js.Global.Get("Phaser").Get("ImageCollection").New(name, firstgid, width, height, margin)}
}
// NewImageCollection4O An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection4O(name string, firstgid int, width int, height int, margin int, spacing int) *ImageCollection {
    return &ImageCollection{js.Global.Get("Phaser").Get("ImageCollection").New(name, firstgid, width, height, margin, spacing)}
}
// NewImageCollection5O An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollection5O(name string, firstgid int, width int, height int, margin int, spacing int, properties interface{}) *ImageCollection {
    return &ImageCollection{js.Global.Get("Phaser").Get("ImageCollection").New(name, firstgid, width, height, margin, spacing, properties)}
}
// NewImageCollectionI An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
func NewImageCollectionI(args ...interface{}) *ImageCollection {
    return &ImageCollection{js.Global.Get("Phaser").Get("ImageCollection").New(args)}
}



// Name The name of the Image Collection.
func (self *ImageCollection) Name() string{
    return self.Object.Get("name").String()
}

// SetNameA The name of the Image Collection.
func (self *ImageCollection) SetNameA(member string) {
    self.Object.Set("name", member)
}

// Firstgid The Tiled firstgid value.
// This is the starting index of the first image index this Image Collection contains.
func (self *ImageCollection) Firstgid() int{
    return self.Object.Get("firstgid").Int()
}

// SetFirstgidA The Tiled firstgid value.
// This is the starting index of the first image index this Image Collection contains.
func (self *ImageCollection) SetFirstgidA(member int) {
    self.Object.Set("firstgid", member)
}

// ImageWidth The width of the widest image (in pixels).
func (self *ImageCollection) ImageWidth() int{
    return self.Object.Get("imageWidth").Int()
}

// SetImageWidthA The width of the widest image (in pixels).
func (self *ImageCollection) SetImageWidthA(member int) {
    self.Object.Set("imageWidth", member)
}

// ImageHeight The height of the tallest image (in pixels).
func (self *ImageCollection) ImageHeight() int{
    return self.Object.Get("imageHeight").Int()
}

// SetImageHeightA The height of the tallest image (in pixels).
func (self *ImageCollection) SetImageHeightA(member int) {
    self.Object.Set("imageHeight", member)
}

// ImageMargin The margin around the images in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) ImageMargin() interface{}{
    return self.Object.Get("imageMargin")
}

// SetImageMarginA The margin around the images in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) SetImageMarginA(member interface{}) {
    self.Object.Set("imageMargin", member)
}

// ImageSpacing The spacing between each image in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) ImageSpacing() int{
    return self.Object.Get("imageSpacing").Int()
}

// SetImageSpacingA The spacing between each image in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) SetImageSpacingA(member int) {
    self.Object.Set("imageSpacing", member)
}

// Properties Image Collection-specific properties that are typically defined in the Tiled editor.
func (self *ImageCollection) Properties() interface{}{
    return self.Object.Get("properties")
}

// SetPropertiesA Image Collection-specific properties that are typically defined in the Tiled editor.
func (self *ImageCollection) SetPropertiesA(member interface{}) {
    self.Object.Set("properties", member)
}

// Images The cached images that are a part of this collection.
func (self *ImageCollection) Images() []interface{}{
	array00 := self.Object.Get("images")
	length00 := array00.Length()
	out00 := make([]interface{}, length00, length00)
	for i00 := 0; i00 < length00; i00++ {
		out00[i00] = array00.Index(i00)
	}
	return out00
}

// SetImagesA The cached images that are a part of this collection.
func (self *ImageCollection) SetImagesA(member []interface{}) {
    self.Object.Set("images", member)
}

// Total The total number of images in the image collection.
func (self *ImageCollection) Total() int{
    return self.Object.Get("total").Int()
}

// SetTotalA The total number of images in the image collection.
func (self *ImageCollection) SetTotalA(member int) {
    self.Object.Set("total", member)
}


// ContainsImageIndex Returns true if and only if this image collection contains the given image index.
func (self *ImageCollection) ContainsImageIndex(imageIndex int) bool{
    return self.Object.Call("containsImageIndex", imageIndex).Bool()
}

// ContainsImageIndexI Returns true if and only if this image collection contains the given image index.
func (self *ImageCollection) ContainsImageIndexI(args ...interface{}) bool{
    return self.Object.Call("containsImageIndex", args).Bool()
}

// AddImage Add an image to this Image Collection.
func (self *ImageCollection) AddImage(gid int, image string) {
    self.Object.Call("addImage", gid, image)
}

// AddImageI Add an image to this Image Collection.
func (self *ImageCollection) AddImageI(args ...interface{}) {
    self.Object.Call("addImage", args)
}

