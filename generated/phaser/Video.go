// Automatic generation for Phaser.Video
// generated file Video.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// A Video object that takes a previously loaded Video from the Phaser Cache and handles playback of it.
// 
// Alternatively it takes a getUserMedia feed from an active webcam and streams the contents of that to
// the Video instead (see `startMediaStream` method)
// 
// The video can then be applied to a Sprite as a texture. If multiple Sprites share the same Video texture and playback
// changes (i.e. you pause the video, or seek to a new time) then this change will be seen across all Sprites simultaneously.
// 
// Due to a bug in IE11 you cannot play a video texture to a Sprite in WebGL. For IE11 force Canvas mode.
// 
// If you need each Sprite to be able to play a video fully independently then you will need one Video object per Sprite.
// Please understand the obvious performance implications of doing this, and the memory required to hold videos in RAM.
// 
// On some mobile browsers such as iOS Safari, you cannot play a video until the user has explicitly touched the screen.
// This works in the same way as audio unlocking. Phaser will handle the touch unlocking for you, however unlike with audio
// it's worth noting that every single Video needs to be touch unlocked, not just the first one. You can use the `changeSource`
// method to try and work around this limitation, but see the method help for details.
// 
// Small screen devices, especially iPod and iPhone will launch the video in its own native video player,
// outside of the Safari browser. There is no way to avoid this, it's a device imposed limitation.
// 
// Note: On iOS if you need to detect when the user presses the 'Done' button (before the video ends)
// then you need to add your own event listener
type Video struct {
    *js.Object
}


// A reference to the currently running game.
func (self *Video) GetGame() *Game{
    return &Game{self.Get("game")}
}

// A reference to the currently running game.
func (self *Video) SetGame(member *Game) {
    self.Set("game", member)
}

// The key of the Video in the Cache, if stored there. Will be `null` if this Video is using the webcam instead.
func (self *Video) GetKey() string{
    return self.Get("key").String()
}

// The key of the Video in the Cache, if stored there. Will be `null` if this Video is using the webcam instead.
func (self *Video) SetKey(member string) {
    self.Set("key", member)
}

// The width of the video in pixels.
func (self *Video) GetWidth() float64{
    return self.Get("width").Float()
}

// The width of the video in pixels.
func (self *Video) SetWidth(member float64) {
    self.Set("width", member)
}

// The height of the video in pixels.
func (self *Video) GetHeight() float64{
    return self.Get("height").Float()
}

// The height of the video in pixels.
func (self *Video) SetHeight(member float64) {
    self.Set("height", member)
}

// The const type of this object.
func (self *Video) GetType() float64{
    return self.Get("type").Float()
}

// The const type of this object.
func (self *Video) SetType(member float64) {
    self.Set("type", member)
}

// If true this video will never send its image data to the GPU when its dirty flag is true. This only applies in WebGL.
func (self *Video) GetDisableTextureUpload() bool{
    return self.Get("disableTextureUpload").Bool()
}

// If true this video will never send its image data to the GPU when its dirty flag is true. This only applies in WebGL.
func (self *Video) SetDisableTextureUpload(member bool) {
    self.Set("disableTextureUpload", member)
}

// true if this video is currently locked awaiting a touch event. This happens on some mobile devices, such as iOS.
func (self *Video) GetTouchLocked() bool{
    return self.Get("touchLocked").Bool()
}

// true if this video is currently locked awaiting a touch event. This happens on some mobile devices, such as iOS.
func (self *Video) SetTouchLocked(member bool) {
    self.Set("touchLocked", member)
}

// This signal is dispatched when the Video starts to play. It sends 3 parameters: a reference to the Video object, if the video is set to loop or not and the playback rate.
func (self *Video) GetOnPlay() *Signal{
    return &Signal{self.Get("onPlay")}
}

// This signal is dispatched when the Video starts to play. It sends 3 parameters: a reference to the Video object, if the video is set to loop or not and the playback rate.
func (self *Video) SetOnPlay(member *Signal) {
    self.Set("onPlay", member)
}

// This signal is dispatched if the Video source is changed. It sends 3 parameters: a reference to the Video object and the new width and height of the new video source.
func (self *Video) GetOnChangeSource() *Signal{
    return &Signal{self.Get("onChangeSource")}
}

// This signal is dispatched if the Video source is changed. It sends 3 parameters: a reference to the Video object and the new width and height of the new video source.
func (self *Video) SetOnChangeSource(member *Signal) {
    self.Set("onChangeSource", member)
}

// This signal is dispatched when the Video completes playback, i.e. enters an 'ended' state. On iOS specifically it also fires if the user hits the 'Done' button at any point during playback. Videos set to loop will never dispatch this signal.
func (self *Video) GetOnComplete() *Signal{
    return &Signal{self.Get("onComplete")}
}

// This signal is dispatched when the Video completes playback, i.e. enters an 'ended' state. On iOS specifically it also fires if the user hits the 'Done' button at any point during playback. Videos set to loop will never dispatch this signal.
func (self *Video) SetOnComplete(member *Signal) {
    self.Set("onComplete", member)
}

// This signal is dispatched if the user allows access to their webcam.
func (self *Video) GetOnAccess() *Signal{
    return &Signal{self.Get("onAccess")}
}

// This signal is dispatched if the user allows access to their webcam.
func (self *Video) SetOnAccess(member *Signal) {
    self.Set("onAccess", member)
}

// This signal is dispatched if an error occurs either getting permission to use the webcam (for a Video Stream) or when trying to play back a video file.
func (self *Video) GetOnError() *Signal{
    return &Signal{self.Get("onError")}
}

// This signal is dispatched if an error occurs either getting permission to use the webcam (for a Video Stream) or when trying to play back a video file.
func (self *Video) SetOnError(member *Signal) {
    self.Set("onError", member)
}

// This signal is dispatched if when asking for permission to use the webcam no response is given within a the Video.timeout limit.
// This may be because the user has picked `Not now` in the permissions window, or there is a delay in establishing the LocalMediaStream.
func (self *Video) GetOnTimeout() *Signal{
    return &Signal{self.Get("onTimeout")}
}

// This signal is dispatched if when asking for permission to use the webcam no response is given within a the Video.timeout limit.
// This may be because the user has picked `Not now` in the permissions window, or there is a delay in establishing the LocalMediaStream.
func (self *Video) SetOnTimeout(member *Signal) {
    self.Set("onTimeout", member)
}

// The amount of ms allowed to elapsed before the Video.onTimeout signal is dispatched while waiting for webcam access.
func (self *Video) GetTimeout() int{
    return self.Get("timeout").Int()
}

// The amount of ms allowed to elapsed before the Video.onTimeout signal is dispatched while waiting for webcam access.
func (self *Video) SetTimeout(member int) {
    self.Set("timeout", member)
}

// The HTML Video Element that is added to the document.
func (self *Video) GetVideo() dom.HTMLVideoElement{
    return WrapHTMLVideoElement(self.Get("video"))
}

// The HTML Video Element that is added to the document.
func (self *Video) SetVideo(member dom.HTMLVideoElement) {
    self.Set("video", member)
}

// The Video Stream data. Only set if this Video is streaming from the webcam via `startMediaStream`.
func (self *Video) GetVideoStream() *MediaStream{
    return &MediaStream{self.Get("videoStream")}
}

// The Video Stream data. Only set if this Video is streaming from the webcam via `startMediaStream`.
func (self *Video) SetVideoStream(member *MediaStream) {
    self.Set("videoStream", member)
}

// Is there a streaming video source? I.e. from a webcam.
func (self *Video) GetIsStreaming() bool{
    return self.Get("isStreaming").Bool()
}

// Is there a streaming video source? I.e. from a webcam.
func (self *Video) SetIsStreaming(member bool) {
    self.Set("isStreaming", member)
}

// When starting playback of a video Phaser will monitor its readyState using a setTimeout call.
// The setTimeout happens once every `Video.retryInterval` ms. It will carry on monitoring the video
// state in this manner until the `retryLimit` is reached and then abort.
func (self *Video) GetRetryLimit() int{
    return self.Get("retryLimit").Int()
}

// When starting playback of a video Phaser will monitor its readyState using a setTimeout call.
// The setTimeout happens once every `Video.retryInterval` ms. It will carry on monitoring the video
// state in this manner until the `retryLimit` is reached and then abort.
func (self *Video) SetRetryLimit(member int) {
    self.Set("retryLimit", member)
}

// The current retry attempt.
func (self *Video) GetRetry() int{
    return self.Get("retry").Int()
}

// The current retry attempt.
func (self *Video) SetRetry(member int) {
    self.Set("retry", member)
}

// The number of ms between each retry at monitoring the status of a downloading video.
func (self *Video) GetRetryInterval() int{
    return self.Get("retryInterval").Int()
}

// The number of ms between each retry at monitoring the status of a downloading video.
func (self *Video) SetRetryInterval(member int) {
    self.Set("retryInterval", member)
}

// The PIXI.Texture.
func (self *Video) GetTexture() *Texture{
    return &Texture{self.Get("texture")}
}

// The PIXI.Texture.
func (self *Video) SetTexture(member *Texture) {
    self.Set("texture", member)
}

// The Frame this video uses for rendering.
func (self *Video) GetTextureFrame() *Frame{
    return &Frame{self.Get("textureFrame")}
}

// The Frame this video uses for rendering.
func (self *Video) SetTextureFrame(member *Frame) {
    self.Set("textureFrame", member)
}

// A snapshot grabbed from the video. This is initially black. Populate it by calling Video.grab().
// When called the BitmapData is updated with a grab taken from the current video playing or active video stream.
// If Phaser has been compiled without BitmapData support this property will always be `null`.
func (self *Video) GetSnapshot() *BitmapData{
    return &BitmapData{self.Get("snapshot")}
}

// A snapshot grabbed from the video. This is initially black. Populate it by calling Video.grab().
// When called the BitmapData is updated with a grab taken from the current video playing or active video stream.
// If Phaser has been compiled without BitmapData support this property will always be `null`.
func (self *Video) SetSnapshot(member *BitmapData) {
    self.Set("snapshot", member)
}

// The current time of the video in seconds. If set the video will attempt to seek to that point in time.
func (self *Video) GetCurrentTime() float64{
    return self.Get("currentTime").Float()
}

// The current time of the video in seconds. If set the video will attempt to seek to that point in time.
func (self *Video) SetCurrentTime(member float64) {
    self.Set("currentTime", member)
}

// The duration of the video in seconds.
func (self *Video) GetDuration() float64{
    return self.Get("duration").Float()
}

// The duration of the video in seconds.
func (self *Video) SetDuration(member float64) {
    self.Set("duration", member)
}

// The progress of this video. This is a value between 0 and 1, where 0 is the start and 1 is the end of the video.
func (self *Video) GetProgress() float64{
    return self.Get("progress").Float()
}

// The progress of this video. This is a value between 0 and 1, where 0 is the start and 1 is the end of the video.
func (self *Video) SetProgress(member float64) {
    self.Set("progress", member)
}

// Gets or sets the muted state of the Video.
func (self *Video) GetMute() bool{
    return self.Get("mute").Bool()
}

// Gets or sets the muted state of the Video.
func (self *Video) SetMute(member bool) {
    self.Set("mute", member)
}

// Gets or sets the paused state of the Video.
// If the video is still touch locked (such as on iOS devices) this call has no effect.
func (self *Video) GetPaused() bool{
    return self.Get("paused").Bool()
}

// Gets or sets the paused state of the Video.
// If the video is still touch locked (such as on iOS devices) this call has no effect.
func (self *Video) SetPaused(member bool) {
    self.Set("paused", member)
}

// Gets or sets the volume of the Video, a value between 0 and 1. The value given is clamped to the range 0 to 1.
func (self *Video) GetVolume() float64{
    return self.Get("volume").Float()
}

// Gets or sets the volume of the Video, a value between 0 and 1. The value given is clamped to the range 0 to 1.
func (self *Video) SetVolume(member float64) {
    self.Set("volume", member)
}

// Gets or sets the playback rate of the Video. This is the speed at which the video is playing.
func (self *Video) GetPlaybackRate() float64{
    return self.Get("playbackRate").Float()
}

// Gets or sets the playback rate of the Video. This is the speed at which the video is playing.
func (self *Video) SetPlaybackRate(member float64) {
    self.Set("playbackRate", member)
}

// Gets or sets if the Video is set to loop.
// Please note that at present some browsers (i.e. Chrome) do not support *seamless* video looping.
// If the video isn't yet set this will always return false.
func (self *Video) GetLoop() bool{
    return self.Get("loop").Bool()
}

// Gets or sets if the Video is set to loop.
// Please note that at present some browsers (i.e. Chrome) do not support *seamless* video looping.
// If the video isn't yet set this will always return false.
func (self *Video) SetLoop(member bool) {
    self.Set("loop", member)
}

// True if the video is currently playing (and not paused or ended), otherwise false.
func (self *Video) GetPlaying() bool{
    return self.Get("playing").Bool()
}

// True if the video is currently playing (and not paused or ended), otherwise false.
func (self *Video) SetPlaying(member bool) {
    self.Set("playing", member)
}



// Connects to an external media stream for the webcam, rather than using a local one.
func (self *Video) ConnectToMediaStreamI(args ...interface{}) *Video{
    return &Video{self.Call("connectToMediaStream", args)}
}

// Instead of playing a video file this method allows you to stream video data from an attached webcam.
// 
// As soon as this method is called the user will be prompted by their browser to "Allow" access to the webcam.
// If they allow it the webcam feed is directed to this Video. Call `Video.play` to start the stream.
// 
// If they block the webcam the onError signal will be dispatched containing the NavigatorUserMediaError
// or MediaStreamError event.
// 
// You can optionally set a width and height for the stream. If set the input will be cropped to these dimensions.
// If not given then as soon as the stream has enough data the video dimensions will be changed to match the webcam device.
// You can listen for this with the onChangeSource signal.
func (self *Video) StartMediaStreamI(args ...interface{}) *Video{
    return &Video{self.Call("startMediaStream", args)}
}

// 
func (self *Video) GetUserMediaTimeoutI(args ...interface{}) {
    self.Call("getUserMediaTimeout", args)
}

// 
func (self *Video) GetUserMediaErrorI(args ...interface{}) {
    self.Call("getUserMediaError", args)
}

// 
func (self *Video) GetUserMediaSuccessI(args ...interface{}) {
    self.Call("getUserMediaSuccess", args)
}

// Creates a new Video element from the given Blob. The Blob must contain the video data in the correct encoded format.
// This method is typically called by the Phaser.Loader and Phaser.Cache for you, but is exposed publicly for convenience.
func (self *Video) CreateVideoFromBlobI(args ...interface{}) *Video{
    return &Video{self.Call("createVideoFromBlob", args)}
}

// Creates a new Video element from the given URL.
func (self *Video) CreateVideoFromURLI(args ...interface{}) *Video{
    return &Video{self.Call("createVideoFromURL", args)}
}

// Called automatically if the video source changes and updates the internal texture dimensions.
// Then dispatches the onChangeSource signal.
func (self *Video) UpdateTextureI(args ...interface{}) {
    self.Call("updateTexture", args)
}

// Called when the video completes playback (reaches and ended state).
// Dispatches the Video.onComplete signal.
func (self *Video) CompleteI(args ...interface{}) {
    self.Call("complete", args)
}

// Starts this video playing if it's not already doing so.
func (self *Video) PlayI(args ...interface{}) *Video{
    return &Video{self.Call("play", args)}
}

// Called when the video starts to play. Updates the texture.
func (self *Video) PlayHandlerI(args ...interface{}) {
    self.Call("playHandler", args)
}

// Stops the video playing.
// 
// This removes all locally set signals.
// 
// If you only wish to pause playback of the video, to resume at a later time, use `Video.paused = true` instead.
// If the video hasn't finished downloading calling `Video.stop` will not abort the download. To do that you need to
// call `Video.destroy` instead.
// 
// If you are using a video stream from a webcam then calling Stop will disconnect the MediaStream session and disable the webcam.
func (self *Video) StopI(args ...interface{}) *Video{
    return &Video{self.Call("stop", args)}
}

// Updates the given Display Objects so they use this Video as their texture.
// This will replace any texture they will currently have set.
func (self *Video) AddI(args ...interface{}) *Video{
    return &Video{self.Call("add", args)}
}

// Creates a new Phaser.Image object, assigns this Video to be its texture, adds it to the world then returns it.
func (self *Video) AddToWorldI(args ...interface{}) *Image{
    return &Image{self.Call("addToWorld", args)}
}

// If the game is running in WebGL this will push the texture up to the GPU if it's dirty.
// This is called automatically if the Video is being used by a Sprite, otherwise you need to remember to call it in your render function.
// If you wish to suppress this functionality set Video.disableTextureUpload to `true`.
func (self *Video) RenderI(args ...interface{}) {
    self.Call("render", args)
}

// Internal handler called automatically by the Video.mute setter.
func (self *Video) SetMuteI(args ...interface{}) {
    self.Call("setMute", args)
}

// Internal handler called automatically by the Video.mute setter.
func (self *Video) UnsetMuteI(args ...interface{}) {
    self.Call("unsetMute", args)
}

// Internal handler called automatically by the Video.paused setter.
func (self *Video) SetPauseI(args ...interface{}) {
    self.Call("setPause", args)
}

// Internal handler called automatically by the Video.paused setter.
func (self *Video) SetResumeI(args ...interface{}) {
    self.Call("setResume", args)
}

// On some mobile browsers you cannot play a video until the user has explicitly touched the video to allow it.
// Phaser handles this via the `setTouchLock` method. However if you have 3 different videos, maybe an "Intro", "Start" and "Game Over"
// split into three different Video objects, then you will need the user to touch-unlock every single one of them.
// 
// You can avoid this by using just one Video object and simply changing the video source. Once a Video element is unlocked it remains
// unlocked, even if the source changes. So you can use this to your benefit to avoid forcing the user to 'touch' the video yet again.
// 
// As you'd expect there are limitations. So far we've found that the videos need to be in the same encoding format and bitrate.
// This method will automatically handle a change in video dimensions, but if you try swapping to a different bitrate we've found it
// cannot render the new video on iOS (desktop browsers cope better).
// 
// When the video source is changed the video file is requested over the network. Listen for the `onChangeSource` signal to know
// when the new video has downloaded enough content to be able to be played. Previous settings such as the volume and loop state
// are adopted automatically by the new video.
func (self *Video) ChangeSourceI(args ...interface{}) *Video{
    return &Video{self.Call("changeSource", args)}
}

// Internal callback that monitors the download progress of a video after changing its source.
func (self *Video) CheckVideoProgressI(args ...interface{}) {
    self.Call("checkVideoProgress", args)
}

// Sets the Input Manager touch callback to be Video.unlock.
// Required for mobile video unlocking. Mostly just used internally.
func (self *Video) SetTouchLockI(args ...interface{}) {
    self.Call("setTouchLock", args)
}

// Enables the video on mobile devices, usually after the first touch.
// If the SoundManager hasn't been unlocked then this will automatically unlock that as well.
// Only one video can be pending unlock at any one time.
func (self *Video) UnlockI(args ...interface{}) {
    self.Call("unlock", args)
}

// Grabs the current frame from the Video or Video Stream and renders it to the Video.snapshot BitmapData.
// 
// You can optionally set if the BitmapData should be cleared or not, the alpha and the blend mode of the draw.
// 
// If you need more advanced control over the grabbing them call `Video.snapshot.copy` directly with the same parameters as BitmapData.copy.
func (self *Video) GrabI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Call("grab", args)}
}

// Removes the Video element from the DOM by calling parentNode.removeChild on itself.
// Also removes the autoplay and src attributes and nulls the reference.
func (self *Video) RemoveVideoElementI(args ...interface{}) {
    self.Call("removeVideoElement", args)
}

// Destroys the Video object. This calls `Video.stop` and then `Video.removeVideoElement`.
// If any Sprites are using this Video as their texture it is up to you to manage those.
func (self *Video) DestroyI(args ...interface{}) {
    self.Call("destroy", args)
}
