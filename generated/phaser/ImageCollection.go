// Automatic generation for Phaser.ImageCollection
// generated file ImageCollection.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// An Image Collection is a special tileset containing mulitple images, with no slicing into each image.
// 
// Image Collections are normally created automatically when Tiled data is loaded.
type ImageCollection struct {
    *js.Object
}


// The name of the Image Collection.
func (self *ImageCollection) GetName() string{
    return self.Get("name").String()
}

// The name of the Image Collection.
func (self *ImageCollection) SetName(member string) {
    self.Set("name", member)
}

// The Tiled firstgid value.
// This is the starting index of the first image index this Image Collection contains.
func (self *ImageCollection) GetFirstgid() int{
    return self.Get("firstgid").Int()
}

// The Tiled firstgid value.
// This is the starting index of the first image index this Image Collection contains.
func (self *ImageCollection) SetFirstgid(member int) {
    self.Set("firstgid", member)
}

// The width of the widest image (in pixels).
func (self *ImageCollection) GetImageWidth() int{
    return self.Get("imageWidth").Int()
}

// The width of the widest image (in pixels).
func (self *ImageCollection) SetImageWidth(member int) {
    self.Set("imageWidth", member)
}

// The height of the tallest image (in pixels).
func (self *ImageCollection) GetImageHeight() int{
    return self.Get("imageHeight").Int()
}

// The height of the tallest image (in pixels).
func (self *ImageCollection) SetImageHeight(member int) {
    self.Set("imageHeight", member)
}

// The margin around the images in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) GetImageMargin() interface{}{
    return self.Get("imageMargin")
}

// The margin around the images in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) SetImageMargin(member interface{}) {
    self.Set("imageMargin", member)
}

// The spacing between each image in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) GetImageSpacing() int{
    return self.Get("imageSpacing").Int()
}

// The spacing between each image in the collection (in pixels).
// Use `setSpacing` to change.
func (self *ImageCollection) SetImageSpacing(member int) {
    self.Set("imageSpacing", member)
}

// Image Collection-specific properties that are typically defined in the Tiled editor.
func (self *ImageCollection) GetProperties() interface{}{
    return self.Get("properties")
}

// Image Collection-specific properties that are typically defined in the Tiled editor.
func (self *ImageCollection) SetProperties(member interface{}) {
    self.Set("properties", member)
}

// The cached images that are a part of this collection.
func (self *ImageCollection) GetImages() []interface{}{
	array := self.Get("images")
	length := array.Length()
	out := make([]interface{}, length, length)
	for i := 0; i < length; i++ {
		out[i] = array.Index(i).Interface()
	}
	return out
}

// The cached images that are a part of this collection.
func (self *ImageCollection) SetImages(member []interface{}) {
    self.Set("images", member)
}

// The total number of images in the image collection.
func (self *ImageCollection) GetTotal() int{
    return self.Get("total").Int()
}

// The total number of images in the image collection.
func (self *ImageCollection) SetTotal(member int) {
    self.Set("total", member)
}



// Returns true if and only if this image collection contains the given image index.
func (self *ImageCollection) ContainsImageIndexI(args ...interface{}) bool{
    return self.Call("containsImageIndex", args).Bool()
}

// Add an image to this Image Collection.
func (self *ImageCollection) AddImageI(args ...interface{}) {
    self.Call("addImage", args)
}
